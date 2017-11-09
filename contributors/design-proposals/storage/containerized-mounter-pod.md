# Containerized mounter using volume utilities in pods

## Goal
Kubernetes should be able to run all utilities that are needed to provision/attach/mount/unmount/detach/delete volumes in *pods* instead of running them on *the host*. The host can be a minimal Linux distribution without tools to create e.g. Ceph RBD or mount GlusterFS volumes.

## Secondary objectives
These are not requirements per se, just things to consider before drawing the final design.
* CNCF designs Container Storage Interface (CSI). So far, this CSI expects that "volume plugins" on each host are long-running processes with a fixed gRPC API. We should aim the same direction, hoping to switch to CSI when it's ready. In other words, there should be one long-running container for a volume plugin that serves all volumes of given type on a host.
* We should try to avoid complicated configuration. The system should work out of the box or with very limited configuration.

## Terminology

**Mount utilities** for a volume plugin are all tools that are necessary to use a volume plugin. This includes not only utilities needed to *mount* the filesystem (e.g. `mount.glusterfs` for Gluster), but also utilities needed to attach, detach, provision or delete the volume, such as `/usr/bin/rbd` for Ceph RBD.

## User story
Admin wants to run Kubernetes on a distro that does not ship `mount.glusterfs` that's needed for GlusterFS volumes.
1. Admin installs and runs Kubernetes in any way.
1. Admin deploys a DaemonSet that runs a pod with `mount.glusterfs` on each node. In future, this could be done by installer.
1. User creates a pod that uses a GlusterFS volume. Kubelet finds a pod with mount utilities on the node and uses it to mount the volume instead of expecting that `mount.glusterfs` is available on the host.

- User does not need to configure anything and sees the pod Running as usual.
- Admin just needs to deploy the DaemonSet.
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

## Design

* Pod with mount utilities puts a registration JSON file into `/var/lib/kubelet/plugin-containers/<plugin name>.json` on the host with name of the container where mount utilities should be executed:
  ```json
  {
      "podNamespace": "kubernetes-storage",
      "podName": "gluster-daemon-set-xtzwv",
      "podUID": "5d1942bd-7358-40e8-9547-a04345c85be9",
      "containerName": "gluster"
  }
  ```
  * Pod UID is used to avoid situation when a pod with mount utilities is terminated and leaves its registration file on the host. Kubelet should not assume that newly started pod with the same namespace+name has the same mount utilities.
  * All slashes in `<plugin name>` must be replaced with tilde, e.g. `kubernetes.io~glusterfs.json`.
  * Creating the file must be atomic so kubelet cannot accidentally read partly written file.

 * All volume plugins use `VolumeHost.GetExec` to get the right exec interface when running their utilities.

 * Kubelet's implementation of `VolumeHost.GetExec` looks at `/var/lib/kubelet/plugin-containers/<plugin name>.json` if it has a container for given volume plugin.
   * If the file exists and referred container is running, it returns `Exec` interface implementation that leads to CRI's `ExecSync` into the container (i.e. `docker exec <container> ...`)
   * If the file does not exist or referred container is not running, it returns `Exec` interface implementation that leads to `os.Exec`. This way, pods do not need to remove the registration file when they're terminated.
   * Kubelet does not cache content of `plugin-containers/`, one extra `open()`/`read()` with each exec won't harm and it makes Kubelet more robust to changes in the directory.

* In future, this registration of volume plugin pods should be replaced by a gRPC interface based on Device Plugin registration.

## Requirements on DaemonSets with mount utilities
These are rules that need to be followed by DaemonSet authors:
* One DaemonSet can serve mount utilities for one or more volume plugins. We expect that one volume plugin per DaemonSet will be the most popular choice.
* One DaemonSet must provide *all* utilities that are needed to provision, attach, mount, unmount, detach and delete a volume for a volume plugin, including `mkfs` and `fsck` utilities if they're needed.
    * E.g. `mkfs.ext4` is likely to be available on all hosts, but a pod with mount utilities should not depend on that nor use it.
    * Kernel modules should be available in the pod with mount utilities too. "Available" does not imply that they need to be shipped in a container, we expect that binding `/lib/modules` from host to `/lib/modules` in the pod will be enough for all modules that are needed by Kubernetes internal volume plugins (all distros I checked incl. the "minimal" ones ship scsi.ko, rbd.ko, nfs.ko and fuse). This will allow future flex volumes ship vendor-specific kernel modules. It's up to the vendor to ensure that any kernel module matches the kernel on the host.
    * The only exception is udev (or similar device manager). Only one udev can run on a system, therefore it should run on the host. If a volume plugin needs to talk to udev (e.g. by calling `udevadm trigger`), they must do it on the host and not in a container with mount utilities.
* It is expected that these daemon sets will run privileged pods that will see host's `/proc`, `/dev`, `/sys`, `/var/lib/kubelet` and such. Especially `/var/lib/kubelet` must be mounted with shared mount propagation so kubelet can see mounts created by the pods.
* The pods with mount utilities should run some simple init as PID 1 that reaps zombies of potential fuse daemons.
* The pods with mount utilities must put a file into `/var/lib/kubelet/plugin-containers/<plugin name>.json` for each volume plugin it supports. It should overwrite any existing file - it's probably leftover from older pod.
   * Admin is responsible to run only one pod with utilities for one volume plugin on a single host. When two pods for say GlusterFS are scheduled on the same node they will overwrite the registration file of each other.
   * Downward API can be used to get pod's name and namespace.
   * Root privileges (or CAP_DAC_OVERRIDE) are needed to write to `/var/lib/kubelet/plugin-containers/`.

To sum it up, it's just a daemon set that spawns privileged pods, running a simple init and registering itself into Kubernetes by placing a file into well-known location.

**Note**: It may be quite difficult to create a pod that see's host's `/dev` and `/sys`, contains necessary kernel modules, does the initialization right and reaps zombies. We're going to provide a template with all this.

### Upgrade
Upgrade of DaemonSets with pods with fuse-based mount utilities needs to be done node by node and with extra care. Killing a pod with fuse daemon(s) inside will un-mount all volumes that are used by other pods on the host and may result in data loss.

In order to update the fuse-based DaemonSet (=GlusterFS or CephFS), admin must do for every node:
* Mark the node as tainted. Only the pod with mount utilities can tolerate the taint, all other pods are evicted. As result, all volumes are unmounted and detached.
* Update the pod.
* Remove the taint.

Is there a way how to make it with DaemonSet rolling update? Is there any other way how to do this upgrade better?

### Containerized kubelet

Kubelet should behave the same when it runs inside a container:
* Use `os.Exec` to run mount utilities inside its own container when no pod with mount utilities is registered. This is current behavior, `mkfs.ext4`, `lsblk`, `rbd` and such are executed in context of the kubelet's container now.
* Use `nsenter <host> mount` to mount things when no pod with mount utilities is registered. Again, this is current behavior.
* Use CRI's `ExecSync` to execute both utilities and the final `mount` when a pod with mount utilities is registered so everything is executed in this pod.

## Open items

* How will controller-manager talk to pods with mount utilities?

  1. Mount pods expose a gRPC service.
      * controller-manager must be configured with the service namespace + name.
      * Some authentication must be implemented (=additional configuration of certificates and whatnot).
      * -> seems to be complicated.

  2. Mount pods run in a dedicated namespace and have labels that tell which volume plugins they can handle.
      * controller manager scans a namespace with a labelselector and does `kubectl exec <pod>` to execute anything in the pod.
      * Needs configuration of the namespace.
      * Admin must make sure that nothing else can run in the namespace (e.g. rogue pods that would steal volumes).
      * Admin/installer must configure access to the namespace so only pv-controller and attach-detach-controller can do `exec` there.

  3. We allow pods to run on hosts that run controller-manager.

      * Usual socket in `/var/lib/kubelet/plugin-sockets` will work.
      * Can it work on GKE?

We do not implement any of these approaches, as we expect that most volume plugins are going to be moved to CSI soon-ish. The only affected volume plugins are:

* Ceph dynamic provisioning - we can use external provisioner during tests.
* Flex - it has its own dynamic registration of flex drivers.

## Implementation notes
As we expect that most volume plugins are going to be moved to CSI soon, all implementation of this proposal will be guarded by alpha feature gate "MountContainers" which is never going leave alpha. Whole implementation of this proposal is going to be removed when the plugins are fully moved to CSI.

Corresponding e2e tests for internal volume plugins will initially run only when with the feature gate is enabled and they will continue running when we move the volume plugins to CSI to ensure we won't introduce regressions.
