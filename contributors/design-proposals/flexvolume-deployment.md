# **Automated Flexvolume Deployment**

## **Objective**

Automate the deployment of Flexvolume drivers with the following goals:
*   Drivers must be deployed on nodes (and master when attach is required) without having to manually access any machine instance.
*   Kubelet and controller-manager do not need to be restarted manually in order for the new plugin to be recognized.

## **Background**

Beginning in version 1.8, the Kubernetes Storage SIG is putting a stop to accepting in-tree volume plugins and advises all storage providers to implement out-of-tree plugins. Currently, there are two recommended implementations: Container Storage Interface (CSI) and Flexvolume.

[CSI](https://github.com/container-storage-interface/spec/blob/master/spec.md) provides a single interface that storage vendors can implement in order for their storage solutions to work across many different container orchestrators, and volume plugins are out-of-tree by design. This is a large effort, the full implementation of CSI is several quarters away, and there is a need for an immediate solution for storage vendors to continue adding volume plugins.

[Flexvolume](https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md) is an in-tree plugin that has the ability to run any storage solution by executing volume commands against a user-provided driver on the Kubernetes host, and this currently exists today. However, the process of setting up Flexvolume is very manual, pushing it out of consideration for many users. Problems include having to copy the driver to a specific location in each node, manually restarting kubelet, and user's limited access to machines.


## **Overview**


### User Story

Driver Installation:

*   Alice is a storage plugin author and would like to deploy a Flexvolume driver on all node instances. She
    1. prepares her Flexvolume driver directory, with driver names in `[vendor~]driver/driver` format (e.g. `flexvolume/k8s~nfs/nfs`, see [Flexvolume documentation](https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md#prerequisites)).
    2. creates an image by copying her driver to the Flexvolume deployment base image at `/flexvolume`.
    3. makes her image available Bob, a cluster admin.
*   Bob modifies the existing deployment DaemonSet spec with the name of the given image, and creates the DaemonSet.
*   Charlie, an end user, creates volumes using the installed plugin.

The user story for driver update is similar: Alice creates a new image with her new drivers, and Bob deploys it using the DaemonSet spec.

Note that the `/flexvolume` directory must look exactly like what is desired in the Flexvolume directory on the host (as described in the [Flexvolume documentation](https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md#prerequisites)). The deployment will replace the existing driver directory on the host with contents in `/flexvolume`. Thus, in order to add a new driver without removing existing ones, existing drivers must also appear in `/flexvolume`.


### High Level Design

The DaemonSet mounts a hostpath volume exposing the host's Flexvolume driver directory onto every pod. The base deployment image contains a script that copies drivers in the image to the hostpath. A notification is then sent to the filesystem watch from kubelet or controller manager. During volume creation, if there is a signal from the watch, kubelet or controller manager probes the driver directory and loads currently installed drivers as volume plugins.


## **Detailed Design**

### Public Deployment Base Image
Composed of the busybox image and the deployment script described in [Driver Installation Script](#driver-installation-script).

### Copying Driver File to Image

Using the deployment base image, the plugin author copies the Flexvolume driver directory (e.g. `flexvolume/k8s~nfs/nfs`) to `/flexvolume` and makes the image available to the cluster admin.


### Driver Installation Script

The script will copy the existing content of `/flexvolume` on the host to a location in `/tmp`, and then attempt to copy user-provided drivers to that directory. If the copy fails, the original drivers are restored. This script will not perform any driver validation.

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
          volumeMounts:
            - mountPath: /flexmnt
              name: flexvolume-mount
      volumes:
        - name: flexvolume-mount
          hostPath:
            path: <host_driver_directory>
```

### Dynamic Plugin Discovery

In the volume plugin code, introduce a `PluginStub` interface containing a single method `Init()`, and have `VolumePlugin` extend it. Create a `PluginProber` type which extends `PluginStub` and includes methods `Init()` and `Probe()`.

`Init()` initializes fsnotify, creates a watch on the driver directory as well as its subdirectories (if any), and spawn a goroutine listening to the signal. When the goroutine receives signal that a new directory is created, create a watch for the directory so that driver changes can be seen.

`Probe()` scans the driver directory only when the goroutine sets a flag. If the flag is set, return true (indicating that new plugins are available) and the list of plugins. Otherwise, return false and nil. After the scan, the watch is refreshed to include the new list of subdirectories. The goroutine should only record a signal if there has been a 1-second delay since the last signal (see [Security Considerations](#security-considerations)). Because inotify (used by fsnotify) can only be used to watch an existing directory, the goroutine needs to maintain the invariant that the driver directory always exists.

Inside `InitPlugins()` from `volume/plugins.go`, if the `PluginStub` is an instance of `PluginProber`, only call its `Init()` and nothing else. Add an additional field, `flexVolumePluginList`, in `VolumePluginMgr` as a cache. For every iteration of the plugin list, call `Probe()` and update `flexVolumePluginList` if true is returned, and iterate through the new plugin list. If the return value is false, iterate through the existing `flexVolumePluginList`.

Because Flexvolume has two separate plugin instantiations (attachable and non-attachable), it's worth considering the case when a driver that implements attach/detach is replaced with a driver that does not, or vice versa. This does not cause an issue because plugins are recreated every time the driver directory is changed.

There is a possibility that a probe occurs at the same time the DaemonSet updates the driver, so the prober's view of drivers is inconsistent. However, this is very rare and when it does occur, the next `Probe()`call, which occurs shortly after, will be consistent.


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

4) Using Jobs instead of DaemonSets to deploy.

Pros: Designed for containers that eventually terminate. No need to have the container go into an infinite loop.

Cons: Does not guarantee every node has a pod running. Pod anti-affinity can be used to ensure no more than one pod runs on the same node, but since the Job spec requests a constant number of pods to run to completion, Jobs cannot ensure that pods are scheduled on new nodes.

5) Have the `flexVolumePluginList` cache live in `PluginProber` instead of `VolumePluginMgr`.

Pros: `VolumePluginMgr` doesn't need to treat Flexvolume plugins any differently from other plugins.

Cons: `PluginProber` doesn't have the function to validate a plugin. This function lives in `VolumePluginMgr`. Alternatively, the function can be passed into `PluginProber`.


## **Security Considerations**

The Flexvolume driver directory can be continuously modified (accidentally or maliciously), making every` Probe()` call trigger a disk read, and `Probe()` calls could happen every couple of milliseconds and in bursts (i.e. lots of calls at first and then silence for some time). This may decrease kubelet's or controller manager's disk IO usage, impacting the performance of other system operations.

As a safety measure, add a 1-second minimum delay between the processing of filesystem watch signals.


## **Testing Plan**

Add new unit tests in `plugin_tests.go` to cover new probing functionality and the heterogeneous plugin types in the plugins list.

Add e2e tests that follow the user story. Write one for initial driver installation, one for an update for the same driver, one for adding another driver, and one for removing a driver.

## **Open Questions**

* How does this system work with containerized kubelet?
* What if drivers are updated while Flexvolume plugin is executing commands?
* If DaemonSet deployment fails, how are errors shown to the user?
