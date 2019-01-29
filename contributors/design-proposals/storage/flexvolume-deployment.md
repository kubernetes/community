# **Dynamic Flexvolume Plugin Discovery**

## **Objective**

Kubelet and controller-manager do not need to be restarted manually in order for new Flexvolume plugins to be recognized.

## **Background**

Beginning in version 1.8, the Kubernetes Storage SIG is putting a stop to accepting in-tree volume plugins and advises all storage providers to implement out-of-tree plugins. Currently, there are two recommended implementations: Container Storage Interface (CSI) and Flexvolume.

[CSI](https://github.com/container-storage-interface/spec/blob/master/spec.md) provides a single interface that storage vendors can implement in order for their storage solutions to work across many different container orchestrators, and volume plugins are out-of-tree by design. This is a large effort, the full implementation of CSI is several quarters away, and there is a need for an immediate solution for storage vendors to continue adding volume plugins.

[Flexvolume] is an in-tree plugin that has the ability to run any storage solution by executing volume commands against a user-provided driver on the Kubernetes host, and this currently exists today. However, the process of setting up Flexvolume is very manual, pushing it out of consideration for many users. Problems include having to copy the driver to a specific location in each node, manually restarting kubelet, and user's limited access to machines.

An automated deployment technique is discussed in [Recommended Driver Deployment Method](#recommended-driver-deployment-method). The crucial change required to enable this method is allowing kubelet and controller manager to dynamically discover plugin changes.


## **Overview**

When there is a modification of the driver directory, a notification is sent to the filesystem watch from kubelet or controller manager. When kubelet or controller-manager searches for plugins (such as when a volume needs to be mounted), if there is a signal from the watch, it probes the driver directory and loads currently installed drivers as volume plugins.

The modification can be a driver install (addition), upgrade/downgrade (update), or uninstall (deletion). If a volume depends on an existing driver, it can be *updated* but not *deleted*.

## **Detailed Design**

In the volume plugin code, introduce a `PluginStub` interface containing a single method `Init()`, and have `VolumePlugin` extend it. Create a `PluginProber` type which extends `PluginStub` and includes methods `Init()` and `Probe()`. Change the type of plugins inside the volume plugin manager's plugin list to `PluginStub`.

`Init()` initializes fsnotify, creates a watch on the driver directory as well as its subdirectories (if any), and spawn a goroutine listening to the signal. When the goroutine receives signal that a new directory is created, create a watch for the directory so that driver changes can be seen.

`Probe()` scans the driver directory only when the goroutine sets a flag. If the flag is set, return true (indicating that new plugins are available) and the list of plugins. Otherwise, return false and nil. After the scan, the watch is refreshed to include the new list of subdirectories. The goroutine should only record a signal if there has been a 1-second delay since the last signal (see [Security Considerations](#security-considerations)). Because inotify (used by fsnotify) can only be used to watch an existing directory, the goroutine needs to maintain the invariant that the driver directory always exists.

Iterating through the list of plugins inside `InitPlugins()` from `volume/plugins.go`, if the plugin is an instance of `PluginProber`, only call its `Init()` and nothing else. Add an additional field, `flexVolumePluginList`, in `VolumePluginMgr` as a cache. For every iteration of the plugin list, call `Probe()` and update `flexVolumePluginList` if true is returned, and iterate through the new plugin list. If the return value is false, iterate through the existing `flexVolumePluginList`. If `Probe()` fails, use the cached plugin instead. However, if the plugin fails to initialize, log the error but do not use the cached version. The user needs to be aware that their driver implementation has a problem initializing, so the system should not silently use an older version.

Because Flexvolume has two separate plugin instantiations (attachable and non-attachable), it's worth considering the case when a driver that implements attach/detach is replaced with a driver that does not, or vice versa. This does not cause an issue because plugins are recreated every time the driver directory is changed.

There is a possibility that a Flexvolume command execution occurs at the same time as the driver is updated, which leads to a bad execution. This cannot be solved within the Kubernetes system without an overhaul. Instead, this is discussed in [Atomic Driver Installation](#atomic-driver-installation) as part of the deployment mechanism. As part of the solution, the Prober will **ignore all files that begins with "."** in the driver directory.

Word of caution about symlinks in the Flexvolume plugin directory: as a result of the recursive filesystem watch implementation, if a symlink links to a directory, unless the directory is visible to the prober (i.e. it's inside the Flexvolume plugin directory and does not start with '.'), the directory's files and subdirectories are not added to filesystem watch, thus their change will not trigger a probe.


## **Alternative Designs**

1) Make `PluginProber` a separate component, and pass it around as a dependency.

Pros: Avoids the common `PluginStub` interface. There isn't much shared functionality between `VolumePlugin` and `PluginProber`. The only purpose this shared abstraction serves is for `PluginProber` to reuse the existing machinery of plugins list.

Cons: Would have to increase dependency surface area, notably `KubeletDeps`.

I'm currently undecided whether to use this design or the `PluginStub` design.

2) Use a polling model instead of a watch for probing for driver changes.

Pros: Simpler to implement.

Cons: Kubelet or controller manager iterates through the plugin list many times, so Probe() is called very frequently. Using this model would increase unnecessary disk usage. This issue is mitigated if we guarantee that `PluginProber` is the last `PluginStub` in the iteration, and only `Probe()` if no other plugin is matched, but this logic adds additional complexity.

3) Use a polling model + cache. Poll every x seconds/minutes.

Pros: Mostly mitigates issues with the previous approach.

Cons: Depending on the polling period, either it's needlessly frequent, or it's too infrequent to pick up driver updates quickly.

4) Have the `flexVolumePluginList` cache live in `PluginProber` instead of `VolumePluginMgr`.

Pros: `VolumePluginMgr` doesn't need to treat Flexvolume plugins any differently from other plugins.

Cons: `PluginProber` doesn't have the function to validate a plugin. This function lives in `VolumePluginMgr`. Alternatively, the function can be passed into `PluginProber`.


## **Security Considerations**

The Flexvolume driver directory can be continuously modified (accidentally or maliciously), making every` Probe()` call trigger a disk read, and `Probe()` calls could happen every couple of milliseconds and in bursts (i.e. lots of calls at first and then silence for some time). This may decrease kubelet's or controller manager's disk IO usage, impacting the performance of other system operations.

As a safety measure, add a 1-second minimum delay between the processing of filesystem watch signals.


## **Testing Plan**

Add new unit tests in `plugin_tests.go` to cover new probing functionality and the heterogeneous plugin types in the plugins list.

Add e2e tests that follow the user story. Write one for initial driver installation, one for an update for the same driver, one for adding another driver, and one for removing a driver.

## **Recommended Driver Deployment Method**

This section describes one possible method to automatically deploy Flexvolume drivers. The goal is that drivers must be deployed on nodes (and master when attach is required) without having to manually access any machine instance.

Driver Installation:

*   Alice is a storage plugin author and would like to deploy a Flexvolume driver on all node instances. She creates an image by copying her driver and the [deployment script](#driver-deployment-script) to a busybox base image, and makes her image available Bob, a cluster admin.
*   Bob modifies the existing deployment DaemonSet spec with the name of the given image, and creates the DaemonSet.
*   Charlie, an end user, creates volumes using the installed plugin.

The user story for driver update is similar: Alice creates a new image with her new drivers, and Bob deploys it using the DaemonSet spec.

### Driver Deployment Script

This script assumes that only a *single driver file* is necessary, and is located at `/$DRIVER` on the deployment image.

``` bash
#!/bin/sh

set -o errexit
set -o pipefail

VENDOR=k8s.io
DRIVER=nfs

# Assuming the single driver file is located at /$DRIVER inside the DaemonSet image.

driver_dir=$VENDOR${VENDOR:+"~"}${DRIVER}
if [ ! -d "/flexmnt/$driver_dir" ]; then
  mkdir "/flexmnt/$driver_dir"
fi

cp "/$DRIVER" "/flexmnt/$driver_dir/.$DRIVER"
mv -f "/flexmnt/$driver_dir/.$DRIVER" "/flexmnt/$driver_dir/$DRIVER"

while : ; do
  sleep 3600
done
```

### Deployment DaemonSet
``` yaml
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: flex-set
spec:
  template:
    metadata:
      name: flex-deploy
      labels:
        app: flex-deploy
    spec:
      containers:
        - image: <deployment_image>
          name: flex-deploy
          securityContext:
              privileged: true
          volumeMounts:
            - mountPath: /flexmnt
              name: flexvolume-mount
      volumes:
        - name: flexvolume-mount
          hostPath:
            path: <host_driver_directory>
```

### Atomic Driver Installation
Regular file copy is not an atomic file operation, so if it were used to install the driver, it's possible that kubelet or controller manager executes the driver when it's partially installed, or the driver gets modified while it's being executed. Care must be taken to ensure the installation operation is atomic.

The deployment script provided above uses renaming, which is atomic, to ensure that from the perspective of kubelet or controller manager, the driver file is completely written to disk in a single operation. The file is first installed with a name prefixed with '.', which the prober ignores.

### Alternatives

* Using Jobs instead of DaemonSets to deploy.

Pros: Designed for containers that eventually terminate. No need to have the container go into an infinite loop.

Cons: Does not guarantee every node has a pod running. Pod anti-affinity can be used to ensure no more than one pod runs on the same node, but since the Job spec requests a constant number of pods to run to completion, Jobs cannot ensure that pods are scheduled on new nodes.

## **Open Questions**

* How does this system work with containerized kubelet?
* Are there any SELinux implications?

[Flexvolume]: /contributors/devel/sig-storage/flexvolume.md