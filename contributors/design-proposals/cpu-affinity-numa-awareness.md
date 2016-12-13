**Author**: Derek Carr

**Last** **Updated**: 12/12/2016

**Status**: Pending Review

# CPU Affinity and NUMA Topology Awareness

This proposal describes enhancements to the Kubernetes API to improve the
utilization of compute resources for containers that have a need to avoid
cross NUMA node memory access by containers.

This proposal describes the set of changes recommended to support NUMA
awareness in the `Node` and `Pod` API to support this feature.  It does not
yet prescribe specifically the actor that performs NUMA aware scheduling
decisions, instead it focuses purely on the underlying primitives required
to support building more advanced scheduling capabilities.

## What is NUMA?

Non-uniform memory architecture (NUMA) describes multi-socket machines that
subdivide memory into nodes where each node is associated with a list of
CPU cores.  This architecture is the norm for modern machines.

An interconnect bus provides connections between nodes so each CPU can
access all memory.  The interconnect can be overwhelmed by concurrent
cross-node traffic, and as a result, processes that need to access memory
on a different node can experience increased latency.

As a result, many applications see a performance benefit when the workload
is affined to a particular NUMA node and CPU core(s).

## NUMA topology

In order to support NUMA affined workloads, the `Node` must make its
NUMA topology available for introspection by other agents that schedule
pods.

This proposal recommends that the `NodeStatus` is augmented as follows:

```
// NodeStatus is information about the current status of a node.
type NodeStatus struct {
    ...
	// Topology represents the NUMA topology of a node to aid NUMA aware scheduling.
	// +optional
	Topology NUMATopology
}

// NUMATopology describes the NUMA topology of a node.
type NUMATopology struct {
	// NUMANodes represents the list of NUMA nodes in the topology.
	NUMANodes []NUMANode
}

// NUMANode describes a single NUMA node.
type NUMANode struct {
	// Identifies a NUMA node on a single host.
	NUMANodeID string
	// Capacity represents the total resources associated to the NUMA node.
	// cpu: 4 <number of cores>
	// memory: <amount of memory in normal page size>
	// hugepages: <amount of memory in huge page size>
	Capacity ResourceList
	// Allocatable represents the resources of a NUMA node that are available for scheduling.
	// +optional
	Allocatable ResourceList
	// CPUSet represents the physical numbers of the CPU cores
	// associated with this node.
	// Example: 0-3 or 0,2,4,6
	// The values are expressed in the List Format syntax specified
	// here: http://man7.org/linux/man-pages/man7/cpuset.7.html
	CPUSet string
}
```

## Node Configuration

### Isolating host processes

By default, load balancing is done across all CPUs, except those marked isolated
using the kernel boot time `isolcpus=` argument.  When configuring a node to support
CPU and NUMA affinity, many operators may wish to isolate host processes to particular
cores.

It is recommended that operators set a CPU value for `--system-reserved`
in whole cores that aligns with the set of cpus that are made available to the default
kernel scheduling algorithm.  If an operator is on a `systemd` managed platform, they
may choose instead to set the `CPUAffinity` value for the root slice to the set of CPU
cores that are reserved for the host processes.

**TODO**

1. how should `kubelet` discover the reserved `cpu-set` value?
1. in a numa system, `kubelet` reservation for memory needs to be removed from
a particular numa node capacity so numa node allocatable is as expected. 

### Configuring Taints

The following `Taint` keys are defined to enable CPU pinning and NUMA awareness.

#### CPUAffinity

* Effect: `NoScheduleNoAdmitNoExecute`
* Potential values:
 * `dedicated`

If `dedicated`, all pods that match this taint will require dedicated compute resources. Each
pod bound to this node must request CPU in whole cores.  The CPU limit must equal the request.

#### NUMACPUAffinity

* Effect: `NoScheduleNoAdmitNoExecute`
* Potential values:
 * `strict`

If `strict`, all pods that match this taint must request CPU (whole or fractional cores) that
fit a single NUMA node `cpu` allocatable.

#### NUMAMemoryPolicy

* Effect: `NoScheduleNoAdmitNoExecute`
* Potential values:
 * `strict`
 * `preferred`

If `strict`, all pods that match this taint must request `memory` that fits it's assigned
NUMA node `memory` allocatable.

If `preferred`, all pods that match this taint are not required to have their `memory` request
fit it's assigned NUMA node `memory` allocatable.

## Pod Specification

### API changes

The following API changes are proposed to the `PodSpec` to allow CPU and NUMA affinity to be defined.

```
// PodSpec is a description of a pod
type PodSpec struct {
...
	// NodeName is a request to schedule this pod onto a specific node.  If it is non-empty,
	// the scheduler simply schedules this pod onto that node, assuming that it fits resource
	// requirements.
	// +optional
	NodeName string
	// Identifies a NUMA node that affines the pod.  If it is non-empty, the value must
	// correspond to a particular NUMA node on the same node that the pod is scheduled against.
	// This value is only set if either the `CPUAffinity` or `NUMACPUAffinity` tolerations
	// are present on the pod.
	// +optional
	NUMANodeID string
	// CPUAffinity controls the CPU affinity of the executed pod.
	// If it is non-empty, the value must correspond to a particular set
	// of CPU cores in the matching NUMA node on the machine that the pod is scheduled against.
	// This value is only set if either the `CPUAffinity` or `NUMACPUAffinity` tolerations
	// are present on the pod.
	// The values are expressed in the List Format syntax specified here:
	// here: http://man7.org/linux/man-pages/man7/cpuset.7.html
	// +optional
	CPUAffinity string	
```

### REST API changes

The `/pod/<pod-name>/bind` operation will allow updating the NUMA and CPU
affinity values.  The same permissions required to schedule a pod to a
node in the cluster will be required to bind a pod to a particular NUMA node
and CPU set.

### Tolerations

Pods that require CPU and NUMA affinity prior to execution must set the
appropriate `Tolerations` for the associated taints.

### Multiple containers

If a pod has multiple containers, the set of containers must all fit
a specific NUMA node, and the set of affined CPUs are shared among containers.

Pod level cgroups are used to actually affine the container to the specified
CPU set.

## Resource Quota changes

Operators must be able to limit the consumption of dedicated CPU cores via quota.

## Kubelet changes

The `kubelet` will enforce the presence of the required pod tolerations assigned to the node.

The `kubelet` will pend the execution of any pod that is assigned to the node, but has
not populated the required fields for a particular toleration.

* If the toleration `CPUAffinity` is present on a `Pod`, the pod will not start
any associated container until the `Pod.Spec.CPUAffinity` is populated.
* If the toleration `NUMAAffinity` is present on a `Pod`, the pod will not start
any associated container until the `Pod.Spec.NUMANodeID` is populated.

The delayed execution of the pod enables both a single and dual-phase scheduler to
place pods on a particular NUMA node and set of CPU cores.

## Required work:

1. pod level cgroup support roll-out
1. implement support for `NoScheduleNoAdmitNoExecute` taint effect
1. expose NUMA topology in cAdvisor
1. expose NUMA topology in node status
1. pod level cgroup support for enabling cpu set

## Future considerations

1. Author `NUMATopologyPredicate` in scheduler to enable NUMA aware scheduling.
1. Restrict vertical autoscaling of CPU and NUMA affined workloads.
 