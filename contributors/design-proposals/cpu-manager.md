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
   and another is not (e.g. auxillary side-car containers responsible for
   log forwarding, metrics collection and the like.)
1. Do not cap CPU quota for guaranteed containers that are granted
   exclusive cores, since that would be antithetical to (1) above.
1. Take physical processor topology into account in the CPU affinity policy.

### Related issues

* Feature: [Further differentiate performance characteristics associated
  with pod level QoS](https://github.com/kubernetes/features/issues/276)

## Proposed changes

### CPU Manager component

The *CPU Manager* is a new software component in Kubelet responsible for
assigning pod containers to sets of CPUs on the local node. In the
future, it may be expanded to control shared processor resources like
caches.

The CPU manager interacts directly with the kuberuntime. The CPU Manager
is notified when containers come and go, before delegating container
creation via the container runtime interface and after the container's
destruction respectively. The CPU Manager emits CPU settings for
containers in response.

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

##### Options for discovering topology

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
type CPUManagerPolicy interface {
  Init(driver CPUDriver, topo CPUTopo)
  Add(c v1.Container, qos QoS) error
  Remove(c v1.Container, qos QoS) error
}

type CPUDriver {
  GetPods() []v1.Pod
  GetCPUs(containerID string) CPUList
  SetCPUs(containerID string, clist CPUList) error
  // Future: RDT L3 and L2 cache masks, etc.
}

type CPUTopo TBD

type CPUList string

func (c CPUList) Size() int {}

// Returns a CPU list with size n and the remainder or
// an error if the request cannot be satisfied, taking
// into account the supplied topology.
//
// @post: c = set_union(taken, remaining),
//        empty_set = set_intersection(taken, remainder)
func (c CPUList) Take(n int, topo CPUTopo) (taken CPUList,
                                            remainder CPUList,
                                            err error) {}

// Returns a CPU list that includes all CPUs in c and d and no others.
//
// @post: result = set_union(c, d)
func (c CPUList) Add(d CPUList) (result CPUList) {}
```

Kubernetes will ship with three CPU manager policies. Only one policy is
active at a time on a given node, chosen by the operator via Kubelet
configuration. The three policies are **no-op**, **static** and **dynamic**.
Each policy is described below.

#### Policy 1: "no-op" cpuset control [default]

This policy preserves the existing Kubelet behavior of doing nothing
with the cgroup `cpuset.cpus` and `cpuset.mems` controls. This “no-op”
policy would become the default CPU Manager policy until the effects of
the other policies are better understood.

#### Policy 2: "static" cpuset control

The "static" policy allocates exclusive CPUs for containers if they are
included in a pod of "Guaranteed" [QoS class][qos] and the container's
resource limit for the CPU resource is an integer greater than or
equal to one.

When exclusive CPUs are allocated for a container, those CPUs are
removed from the allowed CPUs of every other container running on the
node. Once allocated at pod admission time, an exclusive CPU remains
assigned to a single container for the lifetime of the pod (until it
becomes terminal.)

##### Implementation sketch

```go
// Implements CPUManagerPolicy
type staticManager struct {
  driver CPUDriver
  topo   CPUTopo
  // CPU list assigned to non-exclusive containers.
  shared CPUList
}

func (m *staticManager) Init(driver CPUDriver, topo CPUTopo) {
  m.driver = driver
  m.topo = topo
}

func (m *staticManager) Add(c v1.Container, qos QoS) error {
  if p.QoS == GUARANTEED && numExclusive(c) > 0 {
    excl, err := allocate(numExclusive(c))
    if err != nil {
      return err
    }
    m.driver.SetCPUs(c.ID, excl)
    return nil
  }

  // Default case: assign the shared set.
  m.driver.SetCPUs(c.ID, m.shared)
  return nil
}

func (m *staticManager) Remove(c v1.Container, qos QoS) error {
  m.free(m.driver.GetCPUs(c.ID))
}

func (m *staticManager) allocate(n int) (CPUList, err) {
  excl, remaining, err := m.shared.Take(n, m.topo)
  if err != nil {
    return "", err
  }
  m.setShared(remaining)
  return excl, nil
}

func (m *staticManager) free(c CPUList) {
  m.setShared(m.shared.add(c))
}

func (m *staticManager) setShared(c CPUList) {
  prev := m.shared
  m.shared = c
  for _, pod := range m.driver.GetPods() {
    for _, container := range p.Containers {
      if driver.GetCPUs(container.ID) == prev {
        driver.SetCPUs(m.shared)
      }
    }
  }
}

// @pre: container_qos = guaranteed
func numExclusive(c v1.Container) int {
  if c.resources.requests["cpu"] % 1000 == 0 {
    return c.resources.requests["cpu"] / 1000
  }
  return 0
}
```

##### Example pod specs and interpretation

| Pod                                        | Interpretation                 |
| ------------------------------------------ | ------------------------------ |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 0.5 | Container **A** is assigned to the shared cpuset. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 2.0 | Container **A** is assigned two sibling threads on the same physical core (HT) or two physical cores on the same socket (no HT.)<br /><br /> The shared cpuset is shrunk to  make room for the exclusively allocated CPUs. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 1.0<br />&emsp;A:<br />&emsp;&emsp;cpu: 0.5 | Container **A** is assigned one exclusive CPU and container **B** is assigned to the shared cpuset. |
| Pod [Guaranteed]:<br />&emsp;A:<br />&emsp;&emsp;cpu: 1.5<br />&emsp;A:<br />&emsp;&emsp;cpu: 0.5 | Both containers **A** and **B** are assigned to the shared cpuset. |
| Pod [Burstable] | All containers are assigned to the shared cpuset. |
| Pod [BestEffort] | All containers are assigned to the shared cpuset. |

#### Policy 3: "dynamic" cpuset control

_TODO: Describe the policy._

##### Implementation sketch

```go
// Implements CPUManagerPolicy.
type dynamicManager struct {}

func (m *dynamicManager) Init(driver CPUDriver, topo CPUTopo) {
  // TODO
}

func (m *dynamicManager) Add(c v1.Container, qos QoS) error {
  // TODO
}

func (m *dynamicManager) Remove(c v1.Container, qos QoS) error {
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
* Read effective CPU assinments at runtime for alerting. This could be
  satisfied by the checkpointing requirement.
* Configuration
  * How does the CPU Manager coexist with existing kube-reserved
    settings?
  * How does the CPU Manager coexist with related Linux kernel
    configuration (e.g. `isolcpus`.) The operator may want to specify a
    low-water-mark for the size of the shared cpuset. The operator may
    want to correlate exclusive cores with the isolated CPUs, in which
    case the strategy outlined above where allocations are taken
    directly from the shared pool is too simplistic. We could allow an
    explicit pool of cores that may be exclusively allocated and default
    this to the shared pool (leaving at least one core fro the shared
    cpuset to be used for OS, infra and non-exclusive containers.

## Practical challenges

1. Synchronizing CPU Manager state with the container runtime via the
   CRI. Runc/libcontainer allows container cgroup settings to be updtaed
   after creation, but neither the Kubelet docker shim nor the CRI
   implement a similar interface.
    1. Mitigation: [PR 46105](https://github.com/kubernetes/kubernetes/pull/46105)

## Implementation roadmap

### Phase 1

* Internal API exists to allocate CPUs to containers
  ([PR 46105](https://github.com/kubernetes/kubernetes/pull/46105))
* Kubelet configuration includes a CPU manager policy (initially only no-op)
* No-op policy is implemented.
* All existing unit and e2e tests pass.
* Initial unit tests pass.

### Phase 2

* Kubelet can discover "basic" CPU topology (HT-to-physical-core map)
* Static policy is implemented.
* Unit tests for static policy pass.
* e2e tests for static policy pass.
* Performance metrics for one or more plausible synthetic workloads show
  benefit over no-op policy.

### Phase 3

* Dynamic policy is implemented.
* Unit tests for dynamic policy pass.
* e2e tests for dynamic policy pass.
* Performance metrics for one or more plausible synthetic workloads show
  benefit over no-op policy.

### Phase 4

* Kubelet can discover "advanced" CPU topology (NUMA).

## Appendix A: cpuset pitfalls

1. `cpuset.sched_relax_domain_level`
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

[ht]: http://www.intel.com/content/www/us/en/architecture-and-technology/hyper-threading/hyper-threading-technology.html
[hwloc]: https://www.open-mpi.org/projects/hwloc
[procfs]: http://man7.org/linux/man-pages/man5/proc.5.html
[qos]: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-qos.md
[topo]:
http://github.com/intelsdi-x/swan/tree/master/pkg/isolation/topo
