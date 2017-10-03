# Containerized mounter using volume utilities in pods

## Goal
Kubernetes should be able to run all utilities that are needed to provision/attach/mount/unmount/detach/delete volumes in *pods* instead of running them on *the host*. The host can be a minimal Linux distribution without tools to create e.g. Ceph RBD or mount GlusterFS volumes.

## Secondary objectives
These are not requirements per se, just things to consider before drawing the final design.
* CNCF designs Container Storage Interface (CSI). So far, this CSI expects that "volume plugins" on each host are long-running processes with a fixed (gRPC?) API. We should aim the same direction, using exec instead of gRPC, hoping to switch to CSI when it's ready. In other words, there should be one long-running container for a volume plugin that serves all volumes of given type on a host.
* We should try to avoid complicated configuration. The system should work out of the box or with very limited configuration.

## Terminology

**Mount utilities** for a volume pluigin are all tools that are necessary to use a volume plugin. This includes not only utilities needed to *mount* the filesystem (e.g. `mount.glusterfs` for Gluster), but also utilities needed to attach, detach, provision or delete the volume, such as `/usr/bin/rbd` for Ceph RBD.

## User story
Admin wants to run Kubernetes on a distro that does not ship `mount.glusterfs` that's needed for GlusterFS volumes.
1. Admin installs Kubernetes in any way.
2. Admin runs Kubernetes as usual. There are new command line options described below, but they will have sane defaults so no configuration is necessary in most cases.
  * During alpha incubation, kubelet command line option `--experimental-mount-namespace=kube-mount` **must be used** to enable this feature and to tell Kubernetes where to looks for pods with mount utilities. This option will default to `kube-mount` after alpha.
3. Admin deploys a DaemonSet that runs a pod with `mount.glusterfs` on each node in namespace `kube-mount`. In future, this could be done by installer.
4. User creates a pod that uses a GlusterFS volume. Kubelet finds a pod with mount utilities on the node and uses it to mount the volume instead of expecting that `mount.glusterfs` is available on the host.

- User does not need to configure anything and sees the pod Running as usual.
- Admin needs to deploy the DaemonSet and configure Kubernetes a bit.
- It's quite hard to update the DaemonSet, see below.

## Alternatives
### Sidecar containers
We considered this user story:
* Admin installs Kubernetes.
* Admin configures Kubernetes to use sidecar container with template XXX for glusterfs mount/unmount operations and pod with template YYY for glusterfs provision/attach/detach/delete operations. These templates would be yaml files stored somewhere.
* User creates a pod that uses a GlusterFS volume. Kubelet find a sidecar template for gluster, injects it into the pod and runs it before any mount operation. It then uses `docker exec mount <what> <where>` to mount Gluster volumes for the pod. After that, it starts init containers and the "real" pod containers.
* User deletes the pod. Kubelet kills all "real" containers in the pod and uses the sidecar container to unmount gluster volumes. Finally, it kills the sidecar container.

-> User does not need to configure anything and sees the pod Running as usual.
-> Admin needs to set up the templates.

Similarly, when attaching/detaching a volume, attach/detach controller would spawn a pod on a random node and the controller would then use `kubectl exec <the pod> <any attach/detach utility>` to attach/detach the volume. E.g. Ceph RBD volume plugin needs to execute things during attach/detach. After the volume is attached, the controller would kill the pod.

Advantages:
* It's probably easier to update the templates than update the DaemonSet.

Drawbacks:
* Admin needs to store the templates somewhere. Where?
* Short-living processes instead of long-running ones that would mimic CSI (so we could catch bugs early or even redesign CSI).
* Needs some refactoring in kubelet - now kubelet mounts everything and then starts containers. We would need kubelet to start some container(s) first, then mount, then run the rest. This is probably possible, but needs better analysis (and I got lost in kubelet...)

### Infrastructure containers

Mount utilities could be also part of infrastructure container that holds network namespace (when using Docker). Now it's typically simple `pause` container that does not do anything, it could hold mount utilities too.

Advantages:
* Easy to set up
* No extra container running

Disadvantages:
* One container for all mount utilities. Admin needs to make a single container that holds utilities for e.g. both gluster and nfs and whatnot.
* Needs some refactoring in kubelet - now kubelet mounts everything and then starts containers. We would need kubelet to start some container(s) first, then mount, then run the rest. This is probably possible, but needs better analysis (and I got lost in kubelet...)
* Short-living processes instead of long-running ones that would mimic CSI (so we could catch bugs early or even redesign CSI).
* Infrastructure container is implementation detail and CRI does not even allow executing binaries in it.

**We've decided to go with long running DaemonSet pod as described below.**

## Prerequisites
[HostPath volume propagation](propagation.md) must be implemented first.

## Requirements on DaemonSets with mount utilities
These are rules that need to be followed by DaemonSet authors:
* One DaemonSet can serve mount utilities for one or more volume plugins. We expect that one volume plugin per DaemonSet will be the most popular choice.
* One DaemonSet must provide *all* utilities that are needed to provision, attach, mount, unmount, detach and delete a volume for one volume plugin, including `mkfs` and `fsck` utilities if they're needed.
    * E.g. `mkfs.ext4` is likely to be available on all hosts, but a pod with mount utilities should not depend on that nor use it.
    * The only exception are kernel modules. They are not portable across distros and they should be on the host.
* It is expected that these daemon sets will run privileged pods that will see host's `/proc`, `/dev`, `/sys`, `/var/lib/kubelet` and such. Especially `/var/lib/kubelet` must be mounted with shared mount propagation so kubelet can see mounts created by the pods.
* The pods with mount utilities should run some simple init as PID 1 that reaps zombies of potential fuse daemons.
* To allow Kubernetes to discover these pods with mount utilities:
    * All DaemonSets for all chosen volume plugins must run in one dedicated namespace.
    * All pods with mount utilities for a volume plugin `kubernetes.io/foo` must have label `mount.kubernetes.io/foo=true`. 
       * All pods with mount utilities for a flex volume with driver `bar` must have label `mounter.kubernetes.io/flexvolume/bar=true` so there can be different DaemonSets for different flex drivers instead of one monolithic DaemonSet with drivers for all flex volumes.

To sum it up, it's just a daemon set that spawns privileged pods with some labels, running a simple init and waiting for Kubernetes to do `kubectl exec <the pod> <some utility> <args>`.

## Design

### Configuration of the host OS
With the [HostPath volume propagation](propagation.md) implemented, we must ensure that `/var/lib/kubelet` is share-able into containers so a pod with mount utilities can mount something there in its container and this mount will be visible to kubelet and docker (or other container engine) on the host.
We propose:
*  During startup. kubelet checks that `/var/lib/kubelet` on the host is on a mount with shared mount propagation.
    *  If not, kubelet makes `/var/lib/kubelet` shareable:
       ```shell
       mount --bind /var/lib/kubelet /var/lib/kubelet
       mount --make-shared /var/lib/kubelet
       ```
       Note that these commands are executed on the host in case kubelet runs in a container.
       If this startup adjustment fails, kubelet refuses to start with a clear error message.
* Kubelet does not do any other special check. We assume that kubelet either runs on the host and in the same mount namespace as docker and the host, or it runs in a container that sees /var/lib/kubelet as it is on the host and with shared mount propagation.

This ensures that kubelet runs out of the box on any distro without any configuration done by the cluster admin.

### Volume plugins
* All volume plugins need to be updated to use a new `VolumeExec` interface to call external utilities like `mount`, `mkfs`, `rbd lock` and such. Implementation of the interface will be provided by caller and will lead either to `exec` on the host or `kubectl exec` or `docker exec` in a remote or local pod with utilities for appropriate volume plugin (or docker-exec-like command if another container engine is used).

### Controller
* There will be new parameter to kube-controller-manager and kubelet:
    * `--experimental-mount-namespace`, which specifies a dedicated namespace where all pods with mount utilities reside. It would default to `kube-mount`.
* Whenever PV or attach/detach controller needs to call a volume plugin, it looks for *any* running pod in the specified namespace with label `mount.kubernetes.io/foo=true` (or `mount.kubernetes.io/flexvolume/foo=true` for flex volumes) and calls the volume plugin so it all mount utilities are executed as `kubectl exec <pod> xxx` (of course, we'll use clientset interface instead of executing `kubectl`).
* If such pod does not exist, it executes the mount utilities on the host as usual.
* During alpha, no controller-manager changes will be done. That means that Ceph RBD provisioner will still require `/usr/bin/rbd` installed on the master. All other volume plugins will work without any problem, as they don't execute any utility when attaching/detaching/provisioning/deleting a volume.

### Kubelet
* kubelet will get the same parameters as described above, `--experimental-mount-namespace`.
* When kubelet talks to a volume plugin *foo*, it finds a pod in the dedicated namespace running on the node with label `mount.kubernetes.io/foo=true` (or `mount.kubernetes.io/flexvolume/foo=true` for flex volumes) and calls the volume plugin with `VolumeExec` pointing to the pod. All utilities that are executed by the volume plugin for mount/unmount/waitForAttach are executed in the pod running on the node.
* In such pod does not exist, it executes the mount utilities on the host as usual.

As consequence, kubelet will try to run mount utilities on the host when it starts and has not received pods with mount utilities yet. This is likely to fails with a cryptic error:
```
mount: wrong fs type, bad option, bad superblock on 192.168.0.1:/test_vol,
       missing codepage or helper program, or other error
```

Kubelet will periodically retry mounting the volume and it will eventually succeed when pod with mount utilities is scheduled and running on the node.

### VolumePluginMgr
Volume plugin manager runs in attach/detach controller, PV controller and in kubelet and holds a list of all volume plugins. This list of volume plugins is discovered during process startup. Especially for flex volumes, the list is read from `/usr/libexec/kubernetes/...` and it is never updated. We need to update VolumePluginMgr to add flex volumes from running pods.

### Upgrade
Upgrade of the DaemonSet with pods with mount utilities needs to be done node by node and with extra care. The pods may run fuse daemons and killing such pod with glusterfs fuse daemon would kill all pods that use glusterfs on the same node.

In order to update the DaemonSet, admin must do for every node:
* Mark the node as tainted. Only the pod with mount utilities can tolerate the taint, all other pods are evicted. As result, all volumes are unmounted and detached.
* Update the pod.
* Remove the taint.

Is there a way how to make it with DaemonSet rolling update? Is there any other way how to do this upgrade better?


## Implementation notes

* During alpha, only kubelet will be updated and all volume plugins except flex will be updated.
* During alpha, `kubelet --experimental-mount-namespace=<ns>` must be used to enable this feature so it does not break anything accidentally if this feature is buggy. In beta and GA, this feature will be enabled by default and `--experimental-mount-namespace=` could be used to explicitly disable this feature or change the namespace.

Consequences:

* Ceph RBD dynamic provisioning will still need `/usr/bin/rbd` installed on master(s). All other volume plugins will work without any problem, as they don't execute any utility when attaching/detaching/provisioning/deleting a volume.
* Flex still needs `/usr/libexec` scripts deployed to all nodes and master(s).

