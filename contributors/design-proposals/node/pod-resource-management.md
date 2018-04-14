# Kubelet pod level resource management

**Authors**:

1. Buddha Prakash (@dubstack)
1. Vishnu Kannan (@vishh)
1. Derek Carr (@derekwaynecarr)

**Last Updated**: 02/21/2017

**Status**: Implementation planned for Kubernetes 1.6

This document proposes a design for introducing pod level resource accounting
to Kubernetes. It outlines the implementation and associated rollout plan.

## Introduction

Kubernetes supports container level isolation by allowing users
to specify [compute resource requirements](/contributors/design-proposals/scheduling/resources.md) via requests and
limits on individual containers.  The `kubelet` delegates creation of a
cgroup sandbox for each container to its associated container runtime.

Each pod has an associated [Quality of Service (QoS)](resource-qos.md)
class based on the aggregate resource requirements made by individual
containers in the pod.  The `kubelet` has the ability to
[evict pods](kubelet-eviction.md) when compute resources are scarce. It evicts
pods with the lowest QoS class in order to attempt to maintain stability of the
node.

The `kubelet` has no associated cgroup sandbox for individual QoS classes or
individual pods.  This inhibits the ability to perform proper resource
accounting on the node, and introduces a number of code complexities when
trying to build features around QoS.

This design introduces a new cgroup hierarchy to enable the following:

1. Enforce QoS classes on the node. 
1. Simplify resource accounting at the pod level.
1. Allow containers in a pod to share slack resources within its pod cgroup.
For example, a Burstable pod has two containers, where one container makes a
CPU request and the other container does not.  The latter container should
get CPU time not used by the former container.  Today, it must compete for
scare resources at the node level across all BestEffort containers.
1. Ability to charge per container overhead to the pod instead of the node.
This overhead is container runtime specific.  For example, `docker` has
an associated `containerd-shim` process that is created for each container
which should be charged to the pod.
1. Ability to charge any memory usage of memory-backed volumes to the pod when
an individual container exits instead of the node.

## Enabling QoS and Pod level cgroups

To enable the new cgroup hierarchy, the operator must enable the
`--cgroups-per-qos` flag.  Once enabled, the `kubelet` will start managing
inner nodes of the described cgroup hierarchy.

The `--cgroup-root` flag if not specified when the `--cgroups-per-qos` flag
is enabled will default to `/`.  The `kubelet` will parent any cgroups
it creates below that specified value per the
[node allocatable](node-allocatable.md) design.

## Configuring a cgroup driver

The `kubelet` will support manipulation of the cgroup hierarchy on
the host using a cgroup driver. The driver is configured via the
`--cgroup-driver` flag.

The supported values are the following:

* `cgroupfs` is the default driver that performs direct manipulation of the
cgroup filesystem on the host in order to manage cgroup sandboxes.
* `systemd` is an alternative driver that manages cgroup sandboxes using
transient slices for resources that are supported by that init system.

Depending on the configuration of the associated container runtime,
operators may have to choose a particular cgroup driver to ensure
proper system behavior.  For example, if operators use the `systemd`
cgroup driver provided by the `docker` runtime, the `kubelet` must
be configured to use the `systemd` cgroup driver.

Implementation of either driver will delegate to the libcontainer library
in opencontainers/runc.

### Conversion of cgroupfs to systemd naming conventions

Internally, the `kubelet` maintains both an abstract and a concrete name
for its associated cgroup sandboxes.  The abstract name follows the traditional
`cgroupfs` style syntax.  The concrete name is the name for how the cgroup
sandbox actually appears on the host filesystem after any conversions performed
based on the cgroup driver.

If the `systemd` cgroup driver is used, the `kubelet` converts the `cgroupfs`
style syntax into transient slices, and as a result, it must follow `systemd`
conventions for path encoding.

For example, the cgroup name `/burstable/pod123-456` is translated to a
transient slice with the name `burstable-pod123_456.slice`.  Given how
systemd manages the cgroup filesystem, the concrete name for the cgroup
sandbox becomes `/burstable.slice/burstable-pod123_456.slice`.

## Integration with container runtimes

The `kubelet` when integrating with container runtimes always provides the
concrete cgroup filesystem name for the pod sandbox.

## Conversion of CPU millicores to cgroup configuration

Kubernetes measures CPU requests and limits in millicores.

The following formula is used to convert CPU in millicores to cgroup values:

* cpu.shares = (cpu in millicores * 1024) / 1000
* cpu.cfs_period_us = 100000 (i.e. 100ms)
* cpu.cfs_quota_us = quota = (cpu in millicores * 100000) / 1000

## Pod level cgroups

The `kubelet` will create a cgroup sandbox for each pod.

The naming convention for the cgroup sandbox is `pod<pod.UID>`.  It enables
the `kubelet` to associate a particular cgroup on the host filesystem
with a corresponding pod without managing any additional state.  This is useful
when the `kubelet` restarts and needs to verify the cgroup filesystem.

A pod can belong to one of the following 3 QoS classes in decreasing priority:

1. Guaranteed
1. Burstable
1. BestEffort

The resource configuration for the cgroup sandbox is dependent upon the
pod's associated QoS class.

### Guaranteed QoS

A pod in this QoS class has its cgroup sandbox configured as follows:

```
pod<UID>/cpu.shares = sum(pod.spec.containers.resources.requests[cpu])
pod<UID>/cpu.cfs_quota_us = sum(pod.spec.containers.resources.limits[cpu])
pod<UID>/memory.limit_in_bytes = sum(pod.spec.containers.resources.limits[memory])
```

### Burstable QoS

A pod in this QoS class has its cgroup sandbox configured as follows:

```
pod<UID>/cpu.shares = sum(pod.spec.containers.resources.requests[cpu])
```

If all containers in the pod specify a cpu limit:

```
pod<UID>/cpu.cfs_quota_us = sum(pod.spec.containers.resources.limits[cpu])
```

Finally, if all containers in the pod specify a memory limit:

```
pod<UID>/memory.limit_in_bytes = sum(pod.spec.containers.resources.limits[memory])
```

### BestEffort QoS

A pod in this QoS class has its cgroup sandbox configured as follows:

```
pod<UID>/cpu.shares = 2
```

## QoS level cgroups

The `kubelet` defines a `--cgroup-root` flag that is used to specify the `ROOT`
node in the cgroup hierarchy below which the `kubelet` should manage individual
cgroup sandboxes.  It is strongly recommended that users keep the default
value for `--cgroup-root` as `/` in order to avoid deep cgroup hierarchies.  The
`kubelet` creates a cgroup sandbox under the specified path `ROOT/kubepods` per
[node allocatable](node-allocatable.md) to parent pods.  For simplicity, we will
refer to `ROOT/kubepods` as `ROOT` in this document.

The `ROOT` cgroup sandbox is used to parent all pod sandboxes that are in
the Guaranteed QoS class.  By definition, pods in this class have cpu and 
memory limits specified that are equivalent to their requests so the pod 
level cgroup sandbox confines resource consumption without the need of an
additional cgroup sandbox for the tier.

When the `kubelet` launches, it will ensure a `Burstable` cgroup sandbox
and a `BestEffort` cgroup sandbox exist as children of `ROOT`.  These cgroup
sandboxes will parent pod level cgroups in those associated QoS classes.

The `kubelet` highly prioritizes resource utilization, and thus
allows BestEffort and Burstable pods to potentially consume as many
resources that are presently available on the node.

For compressible resources like CPU, the `kubelet` attempts to mitigate
the issue via its use of CPU CFS shares.  CPU time is proportioned
dynamically when there is contention using CFS shares that attempts to
ensure minimum requests are satisfied.

For incompressible resources, this prioritization scheme can inhibit the
ability of a pod to have its requests satisfied.  For example, a Guaranteed
pods memory request may not be satisfied if there are active BestEffort
pods consuming all available memory.

As a node operator, I may want to satisfy the following use cases:

1. I want to prioritize access to compressible resources for my system
and/or kubernetes daemons over end-user pods.
1. I want to prioritize access to compressible resources for my Guaranteed
workloads over my Burstable workloads.
1. I want to prioritize access to compressible resources for my Burstable
workloads over my BestEffort workloads.

Almost all operators are encouraged to support the first use case by enforcing
[node allocatable](node-allocatable.md) via `--system-reserved` and `--kube-reserved`
flags.  It is understood that not all operators may feel the need to extend
that level of reservation to Guaranteed and Burstable workloads if they choose
to prioritize utilization.  That said, many users in the community deploy
cluster services as Guaranteed or Burstable workloads via a `DaemonSet` and would like a similar
resource reservation model as is provided via [node allocatable](node-allocatable)
for system and kubernetes daemons.

For operators that have this concern, the `kubelet` with opt-in configuration
will attempt to limit the ability for a pod in a lower QoS tier to burst utilization
of a compressible resource that was requested by a pod in a higher QoS tier.

The `kubelet` will support a flag `experimental-qos-reserved` that
takes a set of percentages per incompressible resource that controls how the
QoS cgroup sandbox attempts to reserve resources for its tier.  It attempts
to reserve requested resources to exclude pods from lower QoS classes from
using resources requested by higher QoS classes. The flag will accept values
in a range from 0-100%, where a value of `0%` instructs the `kubelet` to attempt
no reservation, and a value of `100%` will instruct the `kubelet` to attempt to
reserve the sum of requested resource across all pods on the node.  The `kubelet`
initially will only support `memory`.  The default value per incompressible
resource if not specified is for no reservation to occur for the incompressible
resource.

Prior to starting a pod, the `kubelet` will attempt to update the
QoS cgroup sandbox associated with the lower QoS tier(s) in order
to prevent consumption of the requested resource by the new pod.
For example, prior to starting a Guaranteed pod, the Burstable
and BestEffort QoS cgroup sandboxes are adjusted.  For resource
specific details, and concerns, see the sections per resource that
follow.

The `kubelet` will allocate resources to the QoS level cgroup
dynamically in response to the following events:

1. kubelet startup/recovery
1. prior to creation of the pod level cgroup
1. after deletion of the pod level cgroup
1. at periodic intervals to reach `experimental-qos-reserved`
heurisitc that converge to a desired state.

All writes to the QoS level cgroup sandboxes are protected via a
common lock in the kubelet to ensure we do not have multiple concurrent
writers to this tier in the hierarchy.

### QoS level CPU allocation

The `BestEffort` cgroup sandbox is statically configured as follows:

```
ROOT/besteffort/cpu.shares = 2
```

This ensures that allocation of CPU time to pods in this QoS class
is given the lowest priority.

The `Burstable` cgroup sandbox CPU share allocation is dynamic based
on the set of pods currently scheduled to the node.  

```
ROOT/burstable/cpu.shares = max(sum(Burstable pods cpu requests, 2)
```

The Burstable cgroup sandbox is updated dynamically in the exit
points described in the previous section.  Given the compressible
nature of CPU, and the fact that cpu.shares are evaluated via relative
priority, the risk of an update being incorrect is minimized as the `kubelet`
converges to a desired state.  Failure to set `cpu.shares` at the QoS level
cgroup would result in `500m` of cpu for a Guaranteed pod to have different
meaning than `500m` of cpu for a Burstable pod in the current hierarchy.  This
is because the default `cpu.shares` value if unspecified is `1024` and `cpu.shares`
are evaluated relative to sibling nodes in the cgroup hierarchy.  As a consequence,
all of the Burstable pods under contention would have a relative priority of 1 cpu
unless updated dynamically to capture the sum of requests.  For this reason,
we will always set `cpu.shares` for the QoS level sandboxes
by default as part of roll-out for this feature.

### QoS level memory allocation

By default, no memory limits are applied to the BestEffort
and Burstable QoS level cgroups unless a `--qos-reserve-requests` value
is specified for memory.

The heuristic that is applied is as follows for each QoS level sandbox:

```
ROOT/burstable/memory.limit_in_bytes = 
    Node.Allocatable - {(summation of memory requests of `Guaranteed` pods)*(reservePercent / 100)}
ROOT/besteffort/memory.limit_in_bytes = 
    Node.Allocatable - {(summation of memory requests of all `Guaranteed` and `Burstable` pods)*(reservePercent / 100)}
```

A value of `--experimental-qos-reserved=memory=100%` will cause the
`kubelet` to adjust the Burstable and BestEffort cgroups from consuming memory
that was requested by a higher QoS class. This increases the risk
of inducing OOM on BestEffort and Burstable workloads in favor of increasing
memory resource guarantees for Guaranteed and Burstable workloads.  A value of
`--experimental-qos-reserved=memory=0%` will allow a Burstable
and BestEffort QoS sandbox to consume up to the full node allocatable amount if
available, but increases the risk that a Guaranteed workload will not have
access to requested memory.

Since memory is an incompressible resource, it is possible that a QoS
level cgroup sandbox may not be able to reduce memory usage below the
value specified in the heuristic described earlier during pod admission
and pod termination.

As a result, the `kubelet` runs a periodic thread to attempt to converge
to this desired state from the above heuristic.  If unreclaimable memory
usage has exceeded the desired limit for the sandbox, the `kubelet` will
attempt to set the effective limit near the current usage to put pressure
on the QoS cgroup sandbox and prevent further consumption.

The `kubelet` will not wait for the QoS cgroup memory limit to converge
to the desired state prior to execution of the pod, but it will always
attempt to cap the existing usage of QoS cgroup sandboxes in lower tiers.
This does mean that the new pod could induce an OOM event at the `ROOT`
cgroup, but ideally per our QoS design, the oom_killer targets a pod
in a lower QoS class, or eviction evicts a lower QoS pod.  The periodic
task is then able to converge to the steady desired state so any future
pods in a lower QoS class do not impact the pod at a higher QoS class.

Adjusting the memory limits for the QoS level cgroup sandbox carries
greater risk given the incompressible nature of memory.  As a result,
we are not enabling this function by default, but would like operators
that want to value resource priority over resource utilization to gather
real-world feedback on its utility.

As a best practice, operators that want to provide a similar resource
reservation model for Guaranteed pods as we offer via enforcement of
node allocatable are encouraged to schedule their Guaranteed pods first
as it will ensure the Burstable and BestEffort tiers have had their QoS
memory limits appropriately adjusted before taking unbounded workload on
node.

## Memory backed volumes

The pod level cgroup ensures that any writes to a memory backed volume
are correctly charged to the pod sandbox even when a container process
in the pod restarts.

All memory backed volumes are removed when a pod reaches a terminal state.

The `kubelet` verifies that a pod's cgroup is deleted from the
host before deleting a pod from the API server as part of the graceful
deletion process.

## Log basic cgroup management

The `kubelet` will log and collect metrics associated with cgroup manipulation.

It will log metrics for cgroup create, update, and delete actions.

## Rollout Plan

### Kubernetes 1.5

The support for the described cgroup hierarchy is experimental.

### Kubernetes 1.6+

The feature will be enabled by default.

As a result, we will recommend that users drain their nodes prior
to upgrade of the `kubelet`.  If users do not drain their nodes, the
`kubelet` will act as follows:

1. If a pod has a `RestartPolicy=Never`, then mark the pod
as `Failed` and terminate its workload.
1. All other pods that are not parented by a pod-level cgroup
will be restarted.

The `cgroups-per-qos` flag will be enabled by default, but user's
may choose to opt-out.  We may deprecate this opt-out mechanism
in Kubernetes 1.7, and remove the flag entirely in Kubernetes 1.8.

#### Risk Assessment

The impact of the unified cgroup hierarchy is restricted to the `kubelet`.

Potential issues:

1. Bugs
1. Performance and/or reliability issues for `BestEffort` pods.  This is
most likely to appear on E2E test runs that mix/match pods across different
QoS tiers.
1. User misconfiguration; most notably the `--cgroup-driver` needs to match
the expected behavior of the container runtime.  We provide clear errors
in `kubelet` logs for container runtimes that we include in tree.

#### Proposed Timeline

* 01/31/2017 - Discuss the rollout plan in sig-node meeting
* 02/14/2017 - Flip the switch to enable pod level cgroups by default
 * enable existing experimental behavior by default
* 02/21/2017 - Assess impacts based on enablement
* 02/27/2017 - Kubernetes Feature complete (i.e. code freeze)
 * opt-in behavior surrounding the feature (`experimental-qos-reserved` support) completed.
* 03/01/2017 - Send an announcement to kubernetes-dev@ about the rollout and potential impact
* 03/22/2017 - Kubernetes 1.6 release
* TBD (1.7?) - Eliminate the option to not use the new cgroup hierarchy.

This is based on the tentative timeline of kubernetes 1.6 release. Need to work out the timeline with the 1.6 release czar.

## Future enhancements

### Add Pod level metrics to Kubelet's metrics provider

Update the `kubelet` metrics provider to include pod level metrics.

### Evaluate supporting evictions local to QoS cgroup sandboxes

Rather than induce eviction at `/` or `/kubepods`, evaluate supporting
eviction decisions for the unbounded QoS tiers (Burstable, BestEffort).

## Examples

The following describes the cgroup representation of a node with pods
across multiple QoS classes.

### Cgroup Hierarchy

The following identifies a sample hierarchy based on the described design.

It assumes the flag `--experimental-qos-reserved` is not enabled for clarity.

```
$ROOT
  |
  +- Pod1
  |   |
  |   +- Container1
  |   +- Container2
  |   ...
  +- Pod2
  |   +- Container3
  |   ...
  +- ...
  |
  +- burstable
  |   |
  |   +- Pod3
  |   |   |
  |   |   +- Container4
  |   |   ...
  |   +- Pod4
  |   |   +- Container5
  |   |   ...
  |   +- ...
  |
  +- besteffort
  |   |
  |   +- Pod5
  |   |   |
  |   |   +- Container6
  |   |   +- Container7
  |   |   ...
  |   +- ...
```

### Guaranteed Pods

We have two pods Pod1 and Pod2 having Pod Spec given below

```yaml
kind: Pod
metadata:
    name: Pod1
spec:
    containers:
        name: foo
            resources:
                limits:
                    cpu: 10m
                    memory: 1Gi
        name: bar
            resources:
                limits:
                    cpu: 100m
                    memory: 2Gi
```

```yaml
kind: Pod
metadata:
    name: Pod2
spec:
    containers:
        name: foo
            resources:
                limits:
                    cpu: 20m
                    memory: 2Gii
```

Pod1 and Pod2 are both classified as Guaranteed and are nested under the `ROOT` cgroup.

```
/ROOT/Pod1/cpu.quota = 110m  
/ROOT/Pod1/cpu.shares = 110m  
/ROOT/Pod1/memory.limit_in_bytes = 3Gi  
/ROOT/Pod2/cpu.quota = 20m  
/ROOT/Pod2/cpu.shares = 20m  
/ROOT/Pod2/memory.limit_in_bytes = 2Gi
```

#### Burstable Pods

We have two pods Pod3 and Pod4 having Pod Spec given below:

```yaml
kind: Pod
metadata:
    name: Pod3
spec:
    containers:
        name: foo
            resources:
                limits:
                    cpu: 50m
                    memory: 2Gi
                requests:
                    cpu: 20m
                    memory: 1Gi
        name: bar
            resources:
                limits:
                    cpu: 100m
                    memory: 1Gi
```

```yaml
kind: Pod
metadata:
    name: Pod4
spec:
    containers:
        name: foo
            resources:
                limits:
                    cpu: 20m
                    memory: 2Gi
                requests:
                    cpu: 10m
                    memory: 1Gi  
```

Pod3 and Pod4 are both classified as Burstable and are hence nested under
the Burstable cgroup.

```
/ROOT/burstable/cpu.shares = 130m
/ROOT/burstable/memory.limit_in_bytes = Allocatable - 5Gi
/ROOT/burstable/Pod3/cpu.quota = 150m
/ROOT/burstable/Pod3/cpu.shares = 120m
/ROOT/burstable/Pod3/memory.limit_in_bytes = 3Gi
/ROOT/burstable/Pod4/cpu.quota = 20m
/ROOT/burstable/Pod4/cpu.shares = 10m
/ROOT/burstable/Pod4/memory.limit_in_bytes = 2Gi
```

#### Best Effort pods

We have a pod, Pod5, having Pod Spec given below:

```yaml
kind: Pod
metadata:
    name: Pod5
spec:
    containers:
        name: foo
            resources:
        name: bar
            resources:
```

Pod5 is classified as BestEffort and is hence nested under the BestEffort cgroup

```
/ROOT/besteffort/cpu.shares = 2
/ROOT/besteffort/cpu.quota= not set
/ROOT/besteffort/memory.limit_in_bytes = Allocatable - 7Gi
/ROOT/besteffort/Pod5/memory.limit_in_bytes = no limit
```

