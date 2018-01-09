# HostPath Volume Propagation

## Abstract

A proposal to add support for propagation mode in HostPath volume, which allows
mounts within containers to visible outside the container and mounts after pods
creation visible to containers. Propagation [modes] (https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) contains "shared", "slave", "private",
"unbindable". Out of them, docker supports "shared" / "slave" / "private".

Several existing issues and PRs were already created regarding that particular
subject:
* Capability to specify mount propagation mode of per volume with docker [#20698] (https://github.com/kubernetes/kubernetes/pull/20698)
* Set propagation to "shared" for hostPath volume [#31504] (https://github.com/kubernetes/kubernetes/pull/31504)

## Use Cases

1. (From @Kaffa-MY) Our team attempts to containerize flocker with zfs as back-end
storage, and launch them in DaemonSet. Containers in the same flocker node need
to read/write and share the same mounted volume. Currently the volume mount
propagation mode cannot be specified between the host and the container, and then
the volume mount of each container would be isolated from each other.
This use case is also referenced by Containerized Volume Client Drivers - Design
Proposal [#22216] (https://github.com/kubernetes/kubernetes/pull/22216)

1. (From @majewsky) I'm currently putting the [OpenStack Swift object storage] (https://github.com/openstack/swift) into
k8s on CoreOS. Swift's storage services expect storage drives to be mounted at
/srv/node/{drive-id} (where {drive-id} is defined by the cluster's ring, the topology
description data structure which is shared between all cluster members). Because
there are several such services on each node (about a dozen, actually), I assemble
/srv/node in the host mount namespace, and pass it into the containers as a hostPath
volume.
Swift is designed such that drives can be mounted and unmounted at any time (most
importantly to hot-swap failed drives) and the services can keep running, but if
the services run in a private mount namespace, they won't see the mounts/unmounts
performed on the host mount namespace until the containers are restarted.
The slave mount namespace is the correct solution for this AFAICS. Until this
becomes available in k8s, we will have to have operations restart containers manually
based on monitoring alerts.

1. (From @victorgp) When using CoreOS Container Linux that does not provides external fuse systems
like, in our case, GlusterFS, and you need a container to do the mounts. The only
way to see those mounts in the host, hence also visible by other containers, is by
sharing the mount propagation.

1. (From @YorikSar) For OpenStack project, Neutron, we need network namespaces
created by it to persist across reboot of pods with Neutron agents. Without it
we have unnecessary data plane downtime during rolling update of these agents.
Neutron L3 agent creates interfaces and iptables rules for each virtual router
in a separate network namespace. For managing them it uses ip netns command that
creates persistent network namespaces by calling unshare(CLONE_NEWNET) and then
bind-mounting new network namespace's inode from /proc/self/ns/net to file with
specified name in /run/netns dir. These bind mounts are the only references to
these namespaces that remain.
When we restart the pod, its mount namespace is destroyed with all these bind
mounts, so all network namespaces created by the agent are gone. For them to
survive we need to bind mount a dir from host mount namespace to container one
with shared flag, so that all bind mounts are propagated across mount namespaces
and references to network namespaces persist.

1. (From https://github.com/kubernetes/kubernetes/issues/46643) I expect the
   container to start and any fuse mounts it creates in a volume that exists on
   other containers in the pod (that are using :slave) are available to those
   other containers.

   In other words, two containers in the same pod share an EmptyDir. One
   container mounts something in it and the other one can see it. The first
   container must have (r)shared mount propagation to the EmptyDir, the second
   one can have (r)slave.


## Implementation Alternatives

### Add an option in VolumeMount API

The new `VolumeMount` will look like:

```go
type MountPropagationMode string

const (
	// MountPropagationHostToContainer means that the volume in a container will
	// receive new mounts from the host or other containers, but filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rslave" in Linux terminology).
	MountPropagationHostToContainer  MountPropagationMode = "HostToContainer"
	// MountPropagationBidirectional means that the volume in a container will
	// receive new mounts from the host or other containers, and its own mounts
	// will be propagated from the container to the host or other containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rshared" in Linux terminology).
	MountPropagationBidirectional MountPropagationMode = "Bidirectional"
)

type VolumeMount struct {
	// Required: This must match the Name of a Volume [above].
	Name string `json:"name"`
	// Optional: Defaults to false (read-write).
	ReadOnly bool `json:"readOnly,omitempty"`
	// Required.
	MountPath string `json:"mountPath"`
	// mountPropagation is the mode how are mounts in the volume propagated from
	// the host to the container and from the container to the host.
	// When not set, MountPropagationHostToContainer is used.
	// This field is alpha in 1.8 and can be reworked or removed in a future
	// release.
	// Optional.
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty"`
}
```

Default would be `HostToContainer`, i.e. `rslave`, which should not break
backward compatibility, `Bidirectional` must be explicitly requested.
Using enum instead of simple `PropagateMounts bool` allows us to extend the
modes to `private` or non-recursive `shared` and `slave` if we need so in
future.

Only privileged containers are allowed to use `Bidirectional` for their volumes.
This will be enforced during validation.

Opinion against this:

1. This will affect all volumes, while only HostPath need this. It could be
checked during validation and any non-HostPath volumes with non-default
propagation could be rejected.

1. This need API change, which is discouraged.

### Add an option in HostPathVolumeSource

The new `HostPathVolumeSource` will look like:

```go
type MountPropagationMode string

const (
	// MountPropagationHostToContainer means that the volume in a container will
	// receive new mounts from the host or other containers, but filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rslave" in Linux terminology).
	MountPropagationHostToContainer  MountPropagationMode = "HostToContainer"
	// MountPropagationBidirectional means that the volume in a container will
	// receive new mounts from the host or other containers, and its own mounts
	// will be propagated from the container to the host or other containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rshared" in Linux terminology).
	MountPropagationBidirectional MountPropagationMode = "Bidirectional"
)

type HostPathVolumeSource struct {
	Path string `json:"path"`
	// mountPropagation is the mode how are mounts in the volume propagated from
	// the host to the container and from the container to the host.
	// When not set, MountPropagationHostToContainer is used.
	// This field is alpha in 1.8 and can be reworked or removed in a future
	// release.
	// Optional.
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty"`
}
```

Default would be `HostToContainer`, i.e. `rslave`, which should not break
backward compatibility, `Bidirectional` must be explicitly requested.
Using enum instead of simple `PropagateMounts bool` allows us to extend the
modes to `private` or non-recursive `shared` and `slave` if we need so in
future.

Only privileged containers can use HostPath with `Bidirectional` mount
propagation - kubelet silently downgrades the propagation to `HostToContainer`
when running `Bidirectional` HostPath in a non-privileged container. This allows
us to use the same `HostPathVolumeSource` in a pod with two containers, one
non-privileged with `HostToContainer` propagation and second privileged with
`Bidirectional` that mounts stuff for the first one.

Opinion against this:

1. This need API change, which is discouraged.

1. All containers use this volume will share the same propagation mode.

1. Silent downgrade from `Bidirectional` to `HostToContainer` for non-privileged
   containers.

1. (From @jonboulle) May cause cross-runtime compatibility issue.

1. It's not possible to validate a pod + mount propagation. Mount propagation
   is stored in a HostPath PersistentVolume object, while privileged mode is
   stored in Pod object. Validator sees only one object and we don't do
   cross-object validation and can't reject non-privileged pod that uses a PV
   with shared mount propagation.

### Make HostPath shared for privileged containers, slave for non-privileged.

Given only HostPath needs this feature, and CAP_SYS_ADMIN access is needed when
making mounts inside container, we can bind propagation mode with existing option
privileged, or we can introduce a new option in SecurityContext to control this.

The propagation mode could be determined by the following logic:

```go
// Environment check to ensure "shared" is supported.
if !dockerNewerThanV110 || !mountPathIsShared {
	return ""
}
if container.SecurityContext.Privileged {
	return "shared"
} else {
	return "slave"
}
```

Opinion against this:

1. This changes the behavior of existing config.

1. (From @euank) "shared" is not correctly supported by some kernels, we need
runtime support matrix and when that will be addressed.

1. This may cause silently fail and be a debuggability nightmare on many
distros.

1. (From @euank) Changing those mountflags may make docker even less stable,
this may lock up kernel accidentally or potentially leak mounts.

1. (From @jsafrane) Typical container that needs to mount something needs to
see host's `/dev` and `/sys` as HostPath volumes. This would make them shared
without any way to opt-out. Docker creates a new `/dev/shm` in the
container, which gets propagated to the host, shadowing host's `/dev/shm`.
Similarly, systemd running in a container is very picky about `/sys/fs/cgroup`
and something prevents it from starting if `/sys` is shared.

## Decision

* We will take 'Add an option in VolumeMount API'
  * With an alpha feature gate in 1.8.
  * Only privileged containers can use `rshared` (`Bidirectional`) mount
    propagation (with a validator).

* During alpha, all the behavior above must be explicitly enabled by
  `kubelet --feature-gates=MountPropagation=true`
  It will be used only for testing of volume plugins in e2e tests and
  Mount propagation may be redesigned or even removed in any future release.

  When the feature is enabled:

  * The default mount propagation of **all** volumes (incl. GCE, AWS, Cinder,
	Gluster, Flex, ...) will be `slave`, which is different to current
	`private`. Extensive testing is needed! We may restrict it to HostPath +
	EmptyDir in Beta.

  * **Any** volume in a privileged container can be `Bidirectional`. We may
  restrict it to HostPath + EmptyDir in Beta.

  * Kubelet's Docker shim layer will check that it is able to run a container
    with shared mount propagation on `/var/lib/kubelet` during startup and log
    a warning otherwise. This ensures that both Docker and kubelet see the same
    `/var/lib/kubelet` and it can be shared into containers.
    E.g. Google COS-58 runs Docker in a separate mount namespace with slave
    propagation and thus can't run a container with shared propagation on
    anything.

    This will be done via simple docker version check (1.13 is required) when
    the feature gate is enabled.

  * Node conformance suite will check that mount propagation in /var/lib/kubelet
    works.

  * When running on a distro with `private` as default mount propagation
    (probably anything that does not run systemd, such as Debian Wheezy),
	Kubelet will make `/var/lib/kubelet` share-able into containers and it will
	refuse to start if it's unsuccessful.

	It sounds complicated, but it's simple
	`mount --bind --rshared /var/lib/kubelet /var/lib/kubelet`. See
	kubernetes/kubernetes#45724


## Extra Concerns

@lucab and @euank has some extra concerns about pod isolation when propagation
modes are changed, listed below:

1. how to clean such pod resources (as mounts are now crossing pod boundaries,
thus they can be kept busy indefinitely by processes outside of the pod)

1. side-effects on restarts (possibly piling up layers of full-propagation mounts)

1. how does this interacts with other mount features (nested volumeMounts may or
may not propagate back to the host, depending of ordering of mount operations)

1. limitations this imposes on runtimes (RO-remounting may now affects the host,
is it on purpose or a dangerous side-effect?)

1. A shared mount target imposes some constraints on its parent subtree (generally,
it has to be shared as well), which in turn prevents some mount operations when
preparing a pod (eg. MS_MOVE).

1. The "on-by-default" nature means existing hostpath mounts, which used to be
harmless, could begin consuming kernel resources and cause a node to crash. Even
if a pod does not create any new mountpoints under its hostpath bindmount, it's
not hard to reach multiplicative explosions with shared bindmounts and so the
change in default + no cleanup could result in existing workloads knocking the
node over.

These concerns are valid and we decide to limit the propagation mode to HostPath
volume only, in HostPath, we expect any runtime should NOT perform any additional
actions (such as clean up). This behavior is also consistent with current HostPath
logic: kube does not take care of the content in HostPath either.
