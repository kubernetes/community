# Node Allocatable Resources

### Authors: timstclair@, vishh@

## Overview

Kubernetes nodes typically run many OS system daemons in addition to kubernetes daemons like kubelet, runtime, etc. and user pods.
Kubernetes assumes that all the compute resources available, referred to as `Capacity`, in a node are available for user pods.
In reality, system daemons use non-trivial amoutn of resources and their availability is critical for the stability of the system.
To address this issue, this proposal introduces the concept of `Allocatable` which identifies the amount of compute resources available to user pods.
Specifically, the kubelet will provide a few knobs to reserve resources for OS system daemons and kubernetes daemons.

By explicitly reserving compute resources, the intention is to avoid overcommiting the node and not have system daemons compete with user pods.
The resources available to system daemons and user pods will be capped based on user specified reservations.

If `Allocatable` is available, the scheduler use that instead of `Capacity`, thereby not overcommiting the node.

## Design

### Definitions

1. **Node Capacity** - Already provided as
   [`NodeStatus.Capacity`](https://htmlpreview.github.io/?https://github.com/kubernetes/kubernetes/blob/HEAD/docs/api-reference/v1/definitions.html#_v1_nodestatus),
   this is total capacity read from the node instance, and assumed to be constant.
2. **System-Reserved** (proposed) - Compute resources reserved for processes which are not managed by
   Kubernetes. Currently this covers all the processes lumped together in the `/system` raw
   container.
3. **Kubelet Allocatable** - Compute resources available for scheduling (including scheduled &
   unscheduled resources). This value is the focus of this proposal. See [below](#api-changes) for
   more details.
4. **Kube-Reserved** (proposed) - Compute resources reserved for Kubernetes components such as the
   docker daemon, kubelet, kube proxy, etc.

### API changes

#### Allocatable

Add `Allocatable` (4) to
[`NodeStatus`](https://htmlpreview.github.io/?https://github.com/kubernetes/kubernetes/blob/HEAD/docs/api-reference/v1/definitions.html#_v1_nodestatus):

```
type NodeStatus struct {
  ...
  // Allocatable represents schedulable resources of a node.
  Allocatable ResourceList `json:"allocatable,omitempty"`
  ...
}
```

Allocatable will be computed by the Kubelet and reported to the API server. It is defined to be:

```
   [Allocatable] = [Node Capacity] - [Kube-Reserved] - [System-Reserved]
```

The scheduler will use `Allocatable` in place of `Capacity` when scheduling pods, and the Kubelet
will use it when performing admission checks.

*Note: Since kernel usage can fluctuate and is out of kubernetes control, it will be reported as a
 separate value (probably via the metrics API). Reporting kernel usage is out-of-scope for this
 proposal.*

#### Kube-Reserved

`KubeReserved` is the parameter specifying resources reserved for kubernetes components (4). It is
provided as a command-line flag to the Kubelet at startup, and therefore cannot be changed during
normal Kubelet operation (this may change in the [future](#future-work)).

The flag will be specified as a serialized `ResourceList`, with resources defined by the API
`ResourceName` and values specified in `resource.Quantity` format, e.g.:

```
--kube-reserved=cpu=500m,memory=5Mi
```

Initially we will only support CPU and memory, but will eventually support more resources like [local storage](#phase-3) and io proportional weights to improve node reliability.

#### System-Reserved

In the initial implementation, `SystemReserved` will be functionally equivalent to
[`KubeReserved`](#kube-reserved), but with a different semantic meaning. While KubeReserved
designates resources set aside for kubernetes components, SystemReserved designates resources set
aside for non-kubernetes components (currently this is reported as all the processes lumped
together in the `/system` raw container on non-systemd nodes).

## Recommended Cgroups Setup

Following is the recommended cgroup configuration for Kubernetes nodes.
All OS system daemons are expected to be placed under a top level `SystemReserved` cgroup.
`Kubelet` and `Container Runtime` are expected to be placed under `KubeReserved` cgroup.
The reason for recommending placing the `Container Runtime` under `KubeReserved` is as follows:

1. A container runtime on Kubernetes nodes is not expected to be used outside of the Kubelet.
1. It's resource consumption is tied to the number of pods running on a node.

Note that the hierarchy below recommends having dedicated cgroups for kubelet and the runtime to individally track their usage.

/ (Cgroup Root)
.
+..systemreserved or system.slice (`SystemReserved` enforced here *optionally* by kubelet)
.        .    .tasks(sshd,udev,etc)
.
.
+..kubereserved or kube.slice (`KubeReserved` enforced here *optionally* by kubelet)
.	 .
.	 +..kubelet
.	 .   .tasks(kubelet)
.        .
.	 +..runtime
.	     .tasks(docker-engine, containerd)
.	 
.
+..kubepods or kubepods.slice (Node Allocatable enforced here by Kubelet)
.	 .
.	 +..PodGuaranteed
.	 .	  .
.	 .	  +..Container1
.	 .	  .        .tasks(container processes)
.	 .	  .
.	 .        +..PodOverhead
.	 .        .        .tasks(per-pod processes)
.	 .        ...
.	 .
.	 +..Burstable
.	 .	  .
.	 .	  +..PodBurstable
.	 .	  .	    .
.	 .	  .	    +..Container1
.	 .	  .	    .         .tasks(container processes)
.	 .	  .	    +..Container2
.	 .	  .	    .         .tasks(container processes)
.	 .	  .	    .
.      	 .     	  .    	    ...
.	 .	  .
.	 .	  ...
.	 .
. 	 .
.	 +..Besteffort
.	 .	  .
.	 .	  +..PodBesteffort
.	 .	  .    	    .
.	 .	  .	    +..Container1
.	 .	  .	    .         .tasks(container processes)
.	 .	  .	    +..Container2
.	 .	  .	    .         .tasks(container processes)
.	 .	  .	    .
.      	 .     	  .    	    ...
.	 .	  .
. 	 .	  ...

`systemreserved` & `kubereserved` cgroups are expected to be created by users. If Kubelet is creating cgroups for itself and docker daemon, it will create the `kubereserved` cgroup automatically.

`kubepods` cgroups will be created by kubelet automatically if it is not already there. If the cgroup driver is set to `systemd` then Kubelet will create a `kubepods.slice` via systemd.
By default, Kubelet will `mkdir` `/kubepods` cgroup directly via cgroupfs.

#### Containerizing Kubelet

If Kubelet is managed using a container runtime, have the runtime create cgroups for kubelet under `kubereserved`.

### Metrics

Kubelet identifies it's own cgroup and exposes it's usage metrics via the Summary metrics API (/stats/summary)
With docker runtime, kubelet identifies docker runtime's cgroups too and exposes metrics for it via the Summary metrics API.
To provide a complete overview of a node, Kubelet will expose metrics from cgroups enforcing `SystemReserved`, `KubeReserved` & `Allocatable` too.

## Relationship with Kubelet Evictions

To improve the reliability of nodes, kubelet evicts pods whenever the node runs out of memory or local storage.
Together, evictions and node allocatable help improve node stability.

As of v1.5, evictions are based on `Capacity` (overall node usage). Kubelet evicts pods based on QoS and user configured eviction thresholds.

From v1.6, if `Allocatable` is enforced by default across all pods on a node using cgroups, pods are not expected to exceed `Allocatable`.
Memory and CPU limits are enforced using cgroups, but there exists no easy means to enforce storage limits though. Enforcing storage limits using Linux Quota is not possible since it's not hierarchical. 
Once storage is supported as a resource for `Allocatable`, Kubelet has to perform evictions based on `Allocatable` in addition to `Capacity`.
More details on [evictions here](./kubelet-eviction.md#enforce-node-allocatable).

Note that even if `KubeReserved` & `SystemReserved` is enforced, kernel memory will still not be restricted and so kubelet will have to continue performing evictions based on overall node usage.

## Implementation Phases

### Phase 1 - Introduce Allocatable to the system without enforcement

**Status**: Implemented v1.2

In this phase, Kubelet will support specifying `KubeReserved` & `SystemReserved` resource reservations via kubelet flags.
The defaults for these flags will be `""`, meaning zero cpu or memory reservations.
Kubelet will compute `Allocatable` and update `Node.Status` to include it.
The scheduler will use `Allocatable` instead of `Capacity` if it is available.

### Phase 2 - Enforce Allocatable on Pods

**Status**: Targetted for v1.6

In this phase, Kubelet will automatically create a top level cgroup to enforce Node Allocatable on all user pods.
Kubelet will support specifying the top level cgroups for `KubeReserved` and `SystemReserved` and support *optionally* placing resource restrictions on these top level cgroups.
Users are expected to specify `KubeReserved` and `SystemReserved` based on their deployment requirements.

Resource requirements for Kubelet and the runtime is typically proportional to the number of pods running on a node.
Once a user identified the maximum pod density for each of their nodes, they will be able to compute `KubeReserved` using [this performance dashboard](http://node-perf-dash.k8s.io/#/builds).
[This blog post](http://blog.kubernetes.io/2016/11/visualize-kubelet-performance-with-node-dashboard.html) explains how the dashboard has to be interpreted.
Note that this dashboard provides usage metrics for docker runtime only as of now.

New flags introduced in this phase are as follows:

1. `--enforce-node-allocatable=[pods][,][kube-reserved],[system-reserved]`
  * This flag will default to `pods` in v1.6.
  * Nodes have to be drained prior to upgrading to v1.6.
  * kubelet's behavior if a node has not been drained prior to upgrades is as follows
    * If a pod has a `RestartPolicy=Never`, then mark the pod as `Failed` and terminate its workload.
    * All other pods that are not parented by new Allocatable top level cgroup will be restarted.
  * Users intending to turn off this feature can set this flag to `""`.
  * `--kube-reserved` and `--system-reserved` flags are expected to be set by users for this flag to have any meaningful effect on node stability.
  * By including `kube-reserved` or `system-reserved` in this flag's value, and by specifying the following two flags, Kubelet will attempt to enforce the reservations specified via `--kube-reserved` & `system-reserved` respectively.

2. `--kube-reserved-cgroup=<absolute path to a cgroup>`
   * This flag helps kubelet identify the control group managing all kube components like Kubelet & container runtime that fall under the `KubeReserved` reservation.

3. `--system-reserved-cgroup=<absolute path to a cgroup>`
   * This flag helps kubelet identify the control group managing all OS specific system daemons that fall under the `SystemReserved` reservation.

#### Rollout details

This phase is expected to improve Kubernetes node stability. However it requires users to specify non-default values for `--kube-reserved` & `--system-reserved` flags though.

The rollout of this phase has been long due and hence we are attempting to include it in v1.6

Since `KubeReserved` and `SystemReserved` continue to have `""` as defaults, the node's `Allocatable` does not change automatically.
Since this phase requires node drains (or pod restarts/terminations), it is considered disruptive to users.

To rollback this phase, set `--enforce-node-allocatable` flag to `""`.

This phase might cause the following symptoms to occur if `--kube-reserved` and/or `--system-reserved` flags are also specified.

1. OOM kills of containers and/or evictions of pods. This can happen primarily to `Burstable` and `BestEffort` pods since they can no longer use up all the resource available on the node.

##### Proposed Timeline

02/14/2017 - Discuss the rollout plan in sig-node meeting
02/15/2017 - Flip the switch to enable pod level cgroups by default
enable existing experimental behavior by default
02/21/2017 - Assess impacts based on enablement
02/27/2017 - Kubernetes Feature complete (i.e. code freeze)
03/01/2017 - Send an announcement to kubernetes-dev@ about this rollout along with rollback options and potential issues. Recommend users to set kube and system reserved.
03/22/2017 - Kubernetes 1.6 release

### Phase 3 - Metrics & support for Storage

*Status*: Targetted for v1.7

In this phase, Kubelet will expose usage metrics for `KubeReserved`, `SystemReserved` and `Allocatable` top level cgroups via Summary metrics API.
`Storage` will also be introduced as a reservable resource in this phase.
Support for evictions based on Allocatable will be introduced in this phase.

## Known Issues

### Kubernetes reservation is smaller than kubernetes component usage

**Solution**: Initially, do nothing (best effort). Let the kubernetes daemons overflow the reserved
resources and hope for the best. If the node usage is less than Allocatable, there will be some room
for overflow and the node should continue to function. If the node has been scheduled to `allocatable`
(worst-case scenario) it may enter an unstable state, which is the current behavior in this
situation.

A recommended alternative is to enforce KubeReserved once Kubelet supports it (Phase 2).
In the [future](#future-work) we may set a parent cgroup for kubernetes components, with limits set
according to `KubeReserved`.

### 3rd party schedulers

The community should be notified that an update to schedulers is recommended, but if a scheduler is
not updated it falls under the above case of "scheduler is not allocatable-resources aware".



<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/node-allocatable.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
