# Processor Cache Management 

_Authors:_

* Malini Bhandaru &lt;malini.k.bhandaru@intel.com&gt;
* Connor Doyle &lt;connor.p.doyle@intel.com&gt;
* Dakshina Ilangovan &lt;dakshina.ilangovan@intel.com&gt;
* Lin Yang &lt;lin.a.yang@intel.com&gt;

**Contents:**

* [Overview](#overview)
* [Proposed changes](#proposed-changes)
* [Example Scenarios](#example-scenarios)
* [Evaluation](#evaluation)
* [Implementation roadmap](#implementation-roadmap)
* [Background Details](#background-details)

## Overview

On shared platforms “noisy neighbors” affect application performance. It is a
serious consideration for latency and jitter sensitive workloads such as
Network Virtual Functions (NFV), Real Time and Interactive applications.
Contention for processor cache affects performance in the same vein as
resources such as CPU and RAM. **We introduce cache allocation to the arsenal
of core pinning, RAM allocation, and job priority to reduce interference and
ensure more predictable performance.** See [whitepaper](https://builders.intel.com/docs/networkbuilders/deterministic_network_functions_virtualization_with_Intel_Resource_Director_Technology.pdf) on the benefits that
cache allocation provides NFV workloads.

Cache is particularly important to hide memory access latency given the
disparity in speed of the CPU core (typically measured in GigaHertz) and memory
access that could be 100s of nanoseconds when there is a last level cache miss.
Processor cache is typically allocated to an application on a first come, first
served basis, with cache lines being evicted based on various [cache
replacement policies](https://en.wikipedia.org/wiki/Cache_replacement_policies) such as Least Recently Used and Least Frequently Used
among others. Different applications have different [working set](https://en.wikipedia.org/wiki/Working_set) sizes, and some known as 
streaming applications, constantly pollute the cache with new lines. Consider
for example a video player or a cryptography application. Without any dedicated
allocation of cache, streaming applications or those with large working sets,
for example App[0] running on Core 0 in the figure below, the majority of the
cache lines may be unavailable to other workloads such as App[1] running on
Core 1. Further, new workloads may come and go, further contributing to
performance variability.

![noisy neighbor](https://user-images.githubusercontent.com/3691428/35775410-e2907a88-093c-11e8-8c7f-6cec5f7f3b6f.png)

A “noisy neighbor” on core zero over-utilizes the shared platform cache
resources causing performance inversion (though the priority app on core one is
higher priority, it runs slower than expected).

In this initial enablement, given few users have a sense of cache significance
(other than deployers of  Network Virtual Functions), we are not seeking to
establish cache as a first class resource, that is a resource that is
explicitly requested at container launch time. But seeking to control cache
allocation to a container based on it QoS.

## Proposed Changes

### Extending CPU Manager

The original [cpu manager design proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/cpu-manager.md) alluded to extending its scope
to control processor cache allocation to obtain more predictable performance. 
The CPU Manager intercepts all container related actions, in particular add,
remove, start, and state update. Below we discuss changes for cache allocation. 

### Pod QoS Class

Kubernetes supports a notion of [pod QoS](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/), the higher the value, the higher
the priority. We treat the Guaranteed QOS pod as deserving of more cache
resources.

### Configuring CPU Manager Cache Resource

It further comprehends three policies, namely “none”, “static”, and “dynamic”,
and policy is selected through kubelet configuration. 

 - **None** The cache resource allocation feature is still experimental, and
   thus the default policy will stay “none” and have  no effect on cache
   reservation. 
 - **Static** While for CPUs the Static policy has logical CPU pools, in
   particular “SHARED”, “RESERVED”, ASSIGNABLE, and “EXCLUSIVE ALLOCATIONS”,
   we shall define for Cache just two COS, namely **“SHARED” and “RESERVED”**.
   The shared pool could be half the available cache ways or any other
   fraction, alternately we could define multiple cache pools with different
   degrees of overlap.  Without loss of generality, for the remainder of this
   discussion, let us consider just two classes of service, namely SHARED and
   RESERVED. Given the fixed assignments, we have the option as a first pass to
   not modify the “state update” code. Alternately, we could modify the update
   code, possibly in an out of band manner, to indicate the number of
   containers assigned to the configured COS. This would be useful in
   implementing back pressure, that is, to decline  launch requests for high
   QOS containers on a heavily subscribed high QOS COS host.
 - **Dynamic**  True dynamic policy will not be implemented at this time. 

### Alternative Cache Configuration

Instead of just two (or more) overlapping COS, one might ask why not partition
the cache and assign each exclusively to a CPU core, along the lines of CPU
pinning. The underlying cache management library provides such support, but in
this experimental  release, we opt for overlapping pools because of the greater
sharing they provide (leading to better best-effort performance) yet providing
less contention for workloads that have a higher QOS requirement. Further, when
workloads are pinned to a CPU core, they already have exclusive access to the
core’s L1 and L2 cache.

### Pre-defined COS
Out of the box, there will be a default set of COS with their associated cache
ways. Documentation shall be provided to indicate how to alter them.

1. SHARED - includes only the lower half of the cache ways
1. RESERVED that includes all the cache ways

In the presence of cache allocation functionality on a platform, Kube
initialization code that runs when a new host is added to a cluster will
register the default or otherwise configured COS.  All Kube and System daemon
workloads could be assigned to either RESERVED or SHARED, and is configurable. 

## Example Scenarios

1. A Container arrives with Guaranteed QoS and other cpu and RAM specific 
   requirements
   1. Kuberuntime calls the CRI delegate to create the container
   1. Kuberuntime adds the containers with the CPU Manager
   1. CPU Manager adds the container to the static policy. The container is
      added to the RESERVED (in terms of number of cache ways) COS.
   1. No change flows related to all container state-reads and resource updates
      that today use the CRI.
1. A Container arrives with Non-Guraranteed QoS and other cpu and RAM specific
   requirements
   Everything proceeds as in Case 1 above except for step ( c ) where the
   container is assigned to the SHARED by virtue of the number of cache ways it
   includes.
1. A container exits
   Everything proceeds as with other container handling by the CPU manager with
   an additional step, the container to COS mapping is cleared and updates via
   CRI reflect the same.

## Evaluation

The following checklist has to be evaluated.

- No perceptible performance difference when the policy is “None” in container
  performance
- No perceptible performance difference when the policy is “Static” or
  “Dynamic” and all container workloads submitted have a priority equal to high.
- Perceptible performance improvement when a latency/jitter sensitive container
  workload is assigned a priority equal to high and all other workloads are
  assigned priority low with either “Static” or “Dynamic” policy.
- Perceptible difference in performance, all things being equal with respect to
  launched containers except for a streaming application such as a video player
  running first as  low priority and next as high priority. In the former
  scenarios, cache disruption is limited to the bottom half while in the latter
  it has full use of the cache and thus affects adversely more co-resident
  container workloads. 

## Implementation Roadmap

### Phase 1
V 1.14 to include the above experimental cache allocation.

### Phase 2
Change the default policy to be Static.

### Future
Dynamically allocate cache, to provide even more resource control and
predictable performance for high value workloads. 

## Background Details

### Hardware Support for Managing Processor Last Level Cache 

Clouds, shared environments, are becoming ubiquitous. Towards achieving more
predictable performance Intel is working on a family of technologies, called 
[IntelⓇ Resource Directory Technologies](https://www.intel.com/content/www/us/en/architecture-and-technology/resource-director-technology.html)(RDT), which allows monitoring resource
usage and enforcing limits for cache, memory and network bandwidth. We shall
leverage Cache Allocation Technology (CAT)  on [processors where the feature](https://github.com/intel/intel-cmt-cat/blob/master/README)
is available to provide the functionality.

CAT allows defining Cache classes-of-service (COS). In the figure below we
illustrate a processor with eight cache ways. Note how COS may overlap in all
the cache ways they include or there may be no sharing. A COS mask must contain
a set of contiguous ways. The number of cache ways available and the number of
COS that may be defined are both processor generation/SKU specific, and
discoverable via the CPUID. The default bitmask essentially provides no cache
reservation, whereas the overlapped bitmask, particularly COS0 provides
assigned processes access to the most amount of cache. All processes will
contend for the M0 cache resource pool. Should we assign all processes to COS0,
it is effectively the same as using the Default Bitmask COSX, providing no
differential cache allocation.

![CAT COS](https://user-images.githubusercontent.com/3691428/35775518-b1ac4530-093e-11e8-8f41-2fcff7fc5ade.png)

### Software Support to Manage Processor Last Level Cache 

Intel has been working on [Linux support for Intel RDT](https://01.org/intel-rdt-linux/blogs/fyu1/2017/resource-allocation-intel%C2%AE-resource-director-technology) using the [resctrl file
system](https://github.com/intel/intel-cmt-cat/wiki/resctrl). We propose cache management leveraging modules from a [new
open source resource management project](https://github.com/intel/rmd/releases) written in GO, Intel initiated to 
provide a higher level abstraction to RDT. 

The library currently handles only cache as a resource. It uses CPUID to detect
the amount of cache available, hides register mask details (which morphs with
each processor generation), facilitates defining cache **classes of service
(COS)**, and supports pre-configuring COS based on whether the anticipated
workloads are storage, network virtual functions, machine learning training or
other traditional workloads based on lab studies.

A container could be assigned to a COS that is  dynamically defined with cache
ways exclusively assigned to it. When the  container exists, the allocated
cache ways are recouped. But, as earlier mentioned due to the limitation on the
mask needing to be a contiguous stream of 1s, it introduces fragmentation,
noticed when seeing to build a larger bundle of cache resources. Either we need
to coalesce cache ways by redefining a COS or till such time as they are
released decline requests for larger resource groups. Redefining a COS just to
include a different set of cache ways carries with it a penalty, it requires
cache re-warming and introduces latency and jitter, the very issue we were
seeking to mitigate. Supporting dynamic COS creation is worth exploring
particularly when Cache becomes a first-class resource. In this initial release
we hope that users gain a sense of the benefits cache resource control brings,
particularly for latency/jitter sensitive workloads such as Network Function
Virtualization.
