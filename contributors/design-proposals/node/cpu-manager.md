# CPU Manager

_Authors:_

* @ConnorDoyle - Connor Doyle &lt;connor.p.doyle@intel.com&gt;
* @flyingcougar - Szymon Scharmach &lt;szymon.scharmach@intel.com&gt;
* @sjenning - Seth Jennings &lt;sjenning@redhat.com&gt;

**Contents:**

* [Overview](#overview)
* [Proposed changes](#proposed-changes)
* [Operations and observability](#operations-and-observability)
* [Practical challenges](#practical-challenges)
* [Implementation roadmap](#implementation-roadmap)
* [Appendix A: cpuset pitfalls](#appendix-a-cpuset-pitfalls)

## Overview

_Problems to solve:_

1. Poor or unpredictable performance observed compared to virtual machine
   based orchestration systems. Application latency and lower CPU
   throughput compared to VMs due to cpu quota being fulfilled across all
   cores, rather than exclusive cores, which results in fewer context
   switches and higher cache affinity.
1. Unacceptable latency attributed to the OS process scheduler, especially
   for “fast” virtual network functions (want to approach line rate on
   modern server NICs.)

_Solution requirements:_

1. Provide an API-driven contract from the system to a user: "if you are a
   Guaranteed pod with 1 or more cores of cpu, the system will try to make
   sure that the pod gets its cpu quota primarily from reserved core(s),
   resulting in fewer context switches and higher cache affinity".
1. Support the case where in a given pod, one container is latency-critical
   and another is not (e.g. auxiliary side-car containers responsible for
   log forwarding, metrics collection and the like.)
1. Do not cap CPU quota for guaranteed containers that are granted
   exclusive cores, since that would be antithetical to (1) above.
1. Take physical processor topology into account in the CPU affinity policy.

### Related issues

* Feature: [Further differentiate performance characteristics associated
  with pod level QoS](https://github.com/kubernetes/features/issues/276)
* Feature: [Add CPU Manager for pod cpuset
  assignment](https://github.com/kubernetes/features/issues/375)

## Proposed changes

### CPU Manager component

The *CPU Manager* is a new software component in Kubelet responsible for
assigning pod containers to sets of CPUs on the local node. In later
phases, the scope will expand to include caches, a critical shared
processor resource.

The kuberuntime notifies the CPU manager when containers come and
go. The first such notification occurs in between the container runtime
interface calls to create and start the container. The second notification
occurs after the container is stopped by the container runtime. The CPU
Manager writes CPU settings for containers using a new CRI method named
[`UpdateContainerResources`](https://github.com/kubernetes/kubernetes/pull/46105).
This new method is invoked from two places in the CPU manager: during each
call to `AddContainer` and also periodically from a separate
reconciliation loop.

![cpu-manager-block-diagram](https://user-images.githubusercontent.com/379372/30137651-2352f4f0-9319-11e7-8be7-0aaeb6ce593a.png)

_CPU Manager block diagram. `Policy`, `State`, and `Topology` types are
factored out of the CPU Manager to promote reuse and to make it easier
to build and test new policies. The shared state abstraction allows
other Kubelet components to be agnostic of the CPU manager policy for
observability and checkpointing extensions._

#### Discovering CPU topology

The CPU Manager must understand basic topology. First of all, it must
determine the number of logical CPUs (hardware threads) available for
allocation. On architectures that support [hyper-threading][ht], sibling
threads share a number of hardware resources including the cache
hierarchy. On multi-socket systems, logical CPUs co-resident on a socket
share L3 cache. Although there may be some programs that benefit from
disjoint caches, the policies described in this proposal assume cache
affinity will yield better application and overall system performance for
most cases. In all scenarios described below, we prefer to acquire logical
CPUs topologically. For example, allocating two CPUs on a system that has
hyper-threading turned on yields both sibling threads on the same
physical core. Likewise, allocating two CPUs on a non-hyper-threaded
system yields two cores on the same socket.

**Decision:** Initially the CPU Manager will re-use the existing discovery
mechanism in cAdvisor.

Alternate options considered for discovering topology:

1. Read and parse the virtual file [`/proc/cpuinfo`][procfs] and construct a
   convenient data structure.
1. Execute a simple program like `lscpu -p` in a subprocess and construct a
   convenient data structure based on the output. Here is an example of
   [data structure to represent CPU topology][topo] in go. The linked package
   contains code to build a ThreadSet from the output of `lscpu -p`.
1. Execute a mature external topology program like [`mpi-hwloc`][hwloc] --
   potentially adding support for the hwloc file format to the Kubelet.

#### CPU Manager interfaces (sketch)

```go
type State interface {
  GetCPUSet(containerID string) (cpuset.CPUSet, bool)
  GetDefaultCPUSet() cpuset.CPUSet
  GetCPUSetOrDefault(containerID string) cpuset.CPUSet
  SetCPUSet(containerID string, cpuset CPUSet)
  SetDefaultCPUSet(cpuset CPUSet)
  Delete(containerID string)
}

type Manager interface {
  Start(ActivePodsFunc, status.PodStatusProvider, runtimeService)
  AddContainer(p *Pod, c *Container, containerID string) error
  RemoveContainer(containerID string) error
  State() state.Reader
}

type Policy interface {
  Name() string
  Start(s state.State)
  AddContainer(s State, pod *Pod, container *Container, containerID string) error
  RemoveContainer(s State, containerID string) error
}

type CPUSet map[int]struct{} // set operations and parsing/formatting helpers

type CPUTopology // convenient type for querying and filtering CPUs
```

#### Configuring the CPU Manager

Kubernetes will ship with three CPU manager policies. Only one policy is
active at a time on a given node, chosen by the operator via Kubelet
configuration. The three policies are **none**, **static** and **dynamic**.

The active CPU manager policy is set through a new Kubelet
configuration value `--cpu-manager-policy`. The default value is `none`.

The CPU manager periodically writes resource updates through the CRI in
order to reconcile in-memory cpuset assignments with cgroupfs. The
reconcile frequency is set through a new Kubelet configuration value
`--cpu-manager-reconcile-period`. If not specified, it defaults to the
same duration as `--node-status-update-frequency` (which itself defaults
to 10 seconds at time of writing.)

Each policy is described below.

#### Policy 1: "none" cpuset control [default]

This policy preserves the existing Kubelet behavior of doing nothing
with the cgroup `cpuset.cpus` and `cpuset.mems` controls. This "none"
policy would become the default CPU Manager policy until the effects of
the other policies are better understood.

#### Policy 2: "static" cpuset control

The "static" policy allocates exclusive CPUs for containers if they are
included in a pod of "Guaranteed" [QoS class][qos] and the container's
resource limit for the CPU resource is an integer greater than or
equal to one. All other containers share a set of CPUs.

When exclusive CPUs are allocated for a container, those CPUs are
removed from the allowed CPUs of every other container running on the
node. Once allocated at pod admission time, an exclusive CPU remains
assigned to a single container for the lifetime of the pod (until it
becomes terminal.)

The Kubelet requires the total CPU reservation from `--kube-reserved`
and `--system-reserved` to be greater than zero when the static policy is
enabled. This is because zero CPU reservation would allow the shared pool to
become empty. The set of reserved CPUs is taken in order of ascending
physical core ID. Operator documentation will be updated to explain how to
configure the system to use the low-numbered physical cores for kube-reserved
and system-reserved cgroups.

Workloads that need to know their own CPU mask, e.g. for managing
thread-level affinity, can read it from the virtual file `/proc/self/status`:

```
$ grep -i cpus /proc/self/status
Cpus_allowed:   77
Cpus_allowed_list:      0-2,4-6
```

Note that containers running in the shared cpuset should not attempt any
application-level CPU affinity of their own, as those settings may be
overwritten without notice (whenever exclusive cores are
allocated or deallocated.)

##### Implementation sketch

The static policy maintains the following sets of logical CPUs:

- **SHARED:** Burstable, BestEffort, and non-integral Guaranteed containers
  run here. Initially this contains all CPU IDs on the system. As
  exclusive allocations are created and destroyed, this CPU set shrinks
  and grows, accordingly. This is stored in the state as the default
  CPU set.

- **RESERVED:** A subset of the shared pool which is not exclusively
  allocatable. The membership of this pool is static for the lifetime of
  the Kubelet. The size of the reserved pool is the ceiling of the total
  CPU reservation from `--kube-reserved` and `--system-reserved`.
  Reserved CPUs are taken topologically starting with lowest-indexed
  physical core, as reported by cAdvisor.

- **ASSIGNABLE:** Equal to `SHARED - RESERVED`. Exclusive CPUs are allocated
  from this pool.

- **EXCLUSIVE ALLOCATIONS:** CPU sets assigned exclusively to one container.
  These are stored as explicit assignments in the state.

When an exclusive allocation is made, the static policy also updates the
default cpuset in the state abstraction. The CPU manager's periodic
reconcile loop takes care of updating the cpuset in cgroupfs for any
containers that may be running in the shared pool. For this reason,
applications running within exclusively-allocated containers must tolerate
potentially sharing their allocated CPUs for up to the CPU manager
reconcile period.

```go
func (p *staticPolicy) Start(s State) {
	fullCpuset := cpuset.NewCPUSet()
	for cpuid := 0; cpuid < p.topology.NumCPUs; cpuid++ {
		fullCpuset.Add(cpuid)
	}
	// Figure out which cores shall not be used in shared pool
	reserved, _ := takeByTopology(p.topology, fullCpuset, p.topology.NumReservedCores)
	s.SetDefaultCPUSet(fullCpuset.Difference(reserved))
}

func (p *staticPolicy) AddContainer(s State, pod *Pod, container *Container, containerID string) error {
  if numCPUs := numGuaranteedCPUs(pod, container); numCPUs != 0 {
    // container should get some exclusively allocated CPUs
    cpuset, err := p.allocateCPUs(s, numCPUs)
    if err != nil {
      return err
    }
    s.SetCPUSet(containerID, cpuset)
  }
  // container belongs in the shared pool (nothing to do; use default cpuset)
  return nil
}

func (p *staticPolicy) RemoveContainer(s State, containerID string) error {
  if toRelease, ok := s.GetCPUSet(containerID); ok {
    s.Delete(containerID)
    s.SetDefaultCPUSet(s.GetDefaultCPUSet().Union(toRelease))
  }
  return nil
}
```

##### Example pod specs and interpretation

| Pod                                        | Interpretation                 |
| ------------------------------------------ | ------------------------------ |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 0.5 | Container **A** is assigned to the shared cpuset. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 2.0 | Container **A** is assigned two sibling threads on the same physical core (HT) or two physical cores on the same socket (no HT.)<br /><br /> The shared cpuset is shrunk to  make room for the exclusively allocated CPUs. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 1.0<br />&emsp;B:<br />&emsp;&emsp;cpu: 0.5 | Container **A** is assigned one exclusive CPU and container **B** is assigned to the shared cpuset. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 1.5<br />&emsp;B:<br />&emsp;&emsp;cpu: 0.5 | Both containers **A** and **B** are assigned to the shared cpuset. |
| Pod [Burstable] | All containers are assigned to the shared cpuset. |
| Pod [BestEffort] | All containers are assigned to the shared cpuset. |

##### Example scenarios and interactions

1. _A container arrives that requires exclusive cores._
    1. Kuberuntime calls the CRI delegate to create the container.
    1. Kuberuntime adds the container with the CPU manager.
    1. CPU manager adds the container to the static policy.
    1. Static policy acquires CPUs from the default pool, by
       topological-best-fit.
    1. Static policy updates the state, adding an assignment for the new
       container and removing those CPUs from the default pool.
    1. CPU manager reads container assignment from the state.
    1. CPU manager updates the container resources via the CRI.
    1. Kuberuntime calls the CRI delegate to start the container.

1. _A container that was assigned exclusive cores terminates._
    1. Kuberuntime removes the container with the CPU manager.
    1. CPU manager removes the container with the static policy.
    1. Static policy adds the container's assigned CPUs back to the default
       pool.
    1. Kuberuntime calls the CRI delegate to remove the container.
    1. Asynchronously, the CPU manager's reconcile loop updates the
       cpuset for all containers running in the shared pool.

1. _The shared pool becomes empty._
    1. This cannot happen. The size of the shared pool is greater than
       the number of exclusively allocatable CPUs. The Kubelet requires the
       total CPU reservation from `--kube-reserved` and `--system-reserved`
       to be greater than zero when the static policy is enabled. The number
       of exclusively allocatable CPUs is
       `floor(capacity.cpu - allocatable.cpu)` and the shared pool initially
       contains all CPUs in the system.

#### Policy 3: "dynamic" cpuset control

_TODO: Describe the policy._

Capturing discussions from resource management meetings and proposal comments:

Unlike the static policy, when the dynamic policy allocates exclusive CPUs to
a container, the cpuset may change during the container's lifetime. If deemed
necessary, we discussed providing a signal in the following way. We could
project (a subset of) the CPU manager state into a volume visible to selected
containers. User workloads could subscribe to update events in a normal Linux
manner (e.g. inotify.)

##### Implementation sketch

```go
func (p *dynamicPolicy) Start(s State) {
	// TODO
}

func (p *dynamicPolicy) AddContainer(s State, pod *Pod, container *Container, containerID string) error {
	// TODO
}

func (p *dynamicPolicy) RemoveContainer(s State, containerID string) error {
	// TODO
}
```

##### Example pod specs and interpretation

| Pod                                        | Interpretation                 |
| ------------------------------------------ | ------------------------------ |
|                                            |                                |
|                                            |                                |

## Operations and observability

* Checkpointing assignments
  * The CPU Manager must be able to pick up where it left off in case the
    Kubelet restarts for any reason.
* Read effective CPU assignments at runtime for alerting. This could be
  satisfied by the checkpointing requirement.

## Practical challenges

1. Synchronizing CPU Manager state with the container runtime via the
   CRI. Runc/libcontainer allows container cgroup settings to be updated
   after creation, but neither the Kubelet docker shim nor the CRI
   implement a similar interface.
    1. Mitigation: [PR 46105](https://github.com/kubernetes/kubernetes/pull/46105)
1. Compatibility with the `isolcpus` Linux kernel boot parameter. The operator
   may want to correlate exclusive cores with the isolated CPUs, in which
   case the static policy outlined above, where allocations are taken
   directly from the shared pool, is too simplistic.
    1. Mitigation: defer supporting this until a new policy tailored for
       use with `isolcpus` can be added.

## Implementation roadmap

### Phase 1: None policy [TARGET: Kubernetes v1.8]

* Internal API exists to allocate CPUs to containers
  ([PR 46105](https://github.com/kubernetes/kubernetes/pull/46105))
* Kubelet configuration includes a CPU manager policy (initially only none)
* None policy is implemented.
* All existing unit and e2e tests pass.
* Initial unit tests pass.

### Phase 2: Static policy [TARGET: Kubernetes v1.8]

* Kubelet can discover "basic" CPU topology (HT-to-physical-core map)
* Static policy is implemented.
* Unit tests for static policy pass.
* e2e tests for static policy pass.
* Performance metrics for one or more plausible synthetic workloads show
  benefit over none policy.

### Phase 3: Beta support [TARGET: Kubernetes v1.9]

* Container CPU assignments are durable across Kubelet restarts.
* Expanded user and operator docs and tutorials.

### Later phases [TARGET: After Kubernetes v1.9]

* Static policy also manages [cache allocation][cat] on supported platforms.
* Dynamic policy is implemented.
* Unit tests for dynamic policy pass.
* e2e tests for dynamic policy pass.
* Performance metrics for one or more plausible synthetic workloads show
  benefit over none policy.
* Kubelet can discover "advanced" topology (NUMA).
* Node-level coordination for NUMA-dependent resource allocations, for example
  devices, CPUs, memory-backed volumes including hugepages.

## Appendix A: cpuset pitfalls

1. [`cpuset.sched_relax_domain_level`][cpuset-files]. "controls the width of
   the range of CPUs over  which  the kernel scheduler performs immediate
   rebalancing of runnable tasks across CPUs."
1. Child cpusets must be subsets of their parents. If B is a child of A,
   then B must be a subset of A. Attempting to shrink A such that B
   would contain allowed CPUs not in A is not allowed (the write will
   fail.) Nested cpusets must be shrunk bottom-up. By the same rationale,
   nested cpusets must be expanded top-down.
1. Dynamically changing cpusets by directly writing to the sysfs would
   create inconsistencies with container runtimes.
1. The `exclusive` flag. This will not be used. We will achieve
   exclusivity for a CPU by removing it from all other assigned cpusets.
1. Tricky semantics when cpusets are combined with CFS shares and quota.

[cat]: http://www.intel.com/content/www/us/en/communications/cache-monitoring-cache-allocation-technologies.html
[cpuset-files]: http://man7.org/linux/man-pages/man7/cpuset.7.html#FILES
[ht]: http://www.intel.com/content/www/us/en/architecture-and-technology/hyper-threading/hyper-threading-technology.html
[hwloc]: https://www.open-mpi.org/projects/hwloc
[node-allocatable]: /contributors/design-proposals/node/node-allocatable.md#phase-2---enforce-allocatable-on-pods
[procfs]: http://man7.org/linux/man-pages/man5/proc.5.html
[qos]: /contributors/design-proposals/node/resource-qos.md
[topo]: http://github.com/intelsdi-x/swan/tree/master/pkg/isolation/topo
