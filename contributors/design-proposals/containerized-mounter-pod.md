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

## Prerequisites
[HostPath volume propagation](propagation.md) must be implemented first.

## Requirements on DaemonSets with mount utilities
These are rules that need to be followed by DaemonSet authors:
* One DaemonSet can serve mount utilities for one or more volume plugins. We expect that one volume plugin per DaemonSet will be the most popular choice.
* One DaemonSet must provide *all* utilities that are needed to provision, attach, mount, unmount, detach and delete a volume for a volume plugin, including `mkfs` and `fsck` utilities if they're needed.
    * E.g. `mkfs.ext4` is likely to be available on all hosts, but a pod with mount utilities should not depend on that nor use it.
    * The only exception are kernel modules. They are not portable across distros and they *should* be on the host.
* It is expected that these daemon sets will run privileged pods that will see host's `/proc`, `/dev`, `/sys`, `/var/lib/kubelet` and such. Especially `/var/lib/kubelet` must be mounted with shared mount propagation so kubelet can see mounts created by the pods.
* The pods with mount utilities should run some simple init as PID 1 that reaps zombies of potential fuse daemons.
* The pods with mount utilities run a daemon with gRPC server that implements `ExecService` defined below.
  * Upon starting, this daemon puts a UNIX domain socket into `/var/lib/kubelet/plugin-sockets/` directory on the host. This way, kubelet is able to discover all pods with mount utilities on a node.
  * Kubernetes will ship implementation of this daemon that creates the socket on the right place and simply executes anything what kubelet asks for.

To sum it up, it's just a daemon set that spawns privileged pods, running a simple init + a daemon that executes mount utilities as requested by kubelet via gRPC.

## Design

### Volume plugins
* All volume plugins need to be updated to use a new `mount.Exec` interface to call external utilities like `mount`, `mkfs`, `rbd lock` and such. Implementation of the interface will be provided by caller and will lead either to simple `os.exec` on the host or a gRPC call to a socket in `/var/lib/kubelet/plugin-sockets/` directory.

### Controllers
TODO after alpha: how will controller-manager talk to a remote pod? It's relatively easy to do something like `kubectl exec <mount pod>` from controller-manager, however it's harder to *discover* the right pod. See Open items below for possible solution(s).

### Kubelet
* When kubelet talks to a volume plugin, it looks for a socket named `/var/lib/kubelet/plugin-sockets/<plugin-name>`. This allows for easier discovery of flex volume drivers - probe in https://github.com/kubernetes/community/pull/833 needs to scan `/var/lib/kubelet/plugin-sockets/` too and find sockets in any new subdirectories.
  * If the socket does not exist, kubelet gives the volume plugin plain `os.Exec` and all mount utilities are executed on the host.
  * If the socket exists, kubelet gives the volume plugin `GRPCExec` and all mount utilities are executed via gRPC on the socket which presumably leads to a pod with mount utilities running a gRPC server.

As consequence, kubelet may try to run mount utilities on the host shortly after startup - it has not received pods with mount utilities yet and thus `/var/lib/kubelet/plugin-sockets/` is empty. This is likely to fails, sometimes with a cryptic error like this:
```
mount: wrong fs type, bad option, bad superblock on 192.168.0.1:/test_vol,
       missing codepage or helper program, or other error
```

Kubelet will periodically retry mounting the volume and it will eventually succeed when pod with mount utilities is scheduled and running on the node.

### gRPC API

`ExecService` is a simple gRPC service defined in [CRI gRPC proto](https://github.com/kubernetes/kubernetes/blob/a1c0510d006ccff9be8478f86635c86658c9bf73/pkg/kubelet/apis/cri/v1alpha1/runtime/api.proto) that allows to execute anything via gRPC:

```protobuf
service ExecService {
    // ExecSync runs a command in a container synchronously.
    rpc ExecSync(ExecSyncRequest) returns (ExecSyncResponse) {}
}
```

* Both `ExecSyncRequest` and `ExecSyncResponse` is copied from [RuntimeService](https://github.com/kubernetes/kubernetes/blob/a1c0510d006ccff9be8478f86635c86658c9bf73/pkg/kubelet/apis/cri/v1alpha1/runtime/api.proto#L65). So far, mount utilities don't need any stdin and stdout+stderr are typically short. Therefore there is no streaming of these file descriptors.

* No authentication / authorization is done on the server side, anyone who connects to the socket can execute anything. It is expected that only root has access to `/var/lib/kubelet/plugin-sockets/`.

* Kubernetes will ship a daemon with server implementation of this API in `cmd/volume-exec`. This implementation simply calls `os.Exec` for each `ExecRequest` it gets and returns the right response.

  * Authors of container images with mount utilities can then add this `volume-exec` daemon to their image, they don't need to care about anything else.

### Upgrade
Upgrade of the DaemonSet with pods with mount utilities needs to be done node by node and with extra care. The pods may run fuse daemons and killing such pod with glusterfs fuse daemon would kill all pods that use glusterfs on the same node.

In order to update the DaemonSet, admin must do for every node:
* Mark the node as tainted. Only the pod with mount utilities can tolerate the taint, all other pods are evicted. As result, all volumes are unmounted and detached.
* Update the pod.
* Remove the taint.

Is there a way how to make it with DaemonSet rolling update? Is there any other way how to do this upgrade better?


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

## Implementation notes

* During alpha, only kubelet will be updated
* Depending on flex dynamic probing in https://github.com/kubernetes/community/pull/833, flex may or may not be supported during alpha.

Consequences:

* Ceph RBD dynamic provisioning will still need `/usr/bin/rbd` installed on master(s). All other volume plugins will work without any problem, as they don't execute any utility when attaching/detaching/provisioning/deleting a volume.
* Flex still needs `/usr/libexec` scripts deployed to master(s) and maybe to nodes.
