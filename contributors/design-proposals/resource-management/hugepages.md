# HugePages support in Kubernetes

**Authors**
* Derek Carr (@derekwaynecarr)
* Seth Jennings (@sjenning)
* Piotr Prokop (@PiotrProkop)

**Status**: In progress

## Abstract

A proposal to enable applications running in a Kubernetes cluster to use huge
pages.

A pod may request a number of huge pages.  The `scheduler` is able to place the
pod on a node that can satisfy that request.  The `kubelet` advertises an
allocatable number of huge pages to support scheduling decisions. A pod may
consume hugepages via `hugetlbfs` or `shmget`.  Huge pages are not
overcommitted.

## Motivation

Memory is managed in blocks known as pages.  On most systems, a page is 4Ki. 1Mi
of memory is equal to 256 pages; 1Gi of memory is 256,000 pages, etc. CPUs have
a built-in memory management unit that manages a list of these pages in
hardware. The Translation Lookaside Buffer (TLB) is a small hardware cache of
virtual-to-physical page mappings.  If the virtual address passed in a hardware
instruction can be found in the TLB, the mapping can be determined quickly.  If
not, a TLB miss occurs, and the system falls back to slower, software based
address translation.  This results in performance issues.  Since the size of the
TLB is fixed, the only way to reduce the chance of a TLB miss is to increase the
page size.

A huge page is a memory page that is larger than 4Ki.  On x86_64 architectures,
there are two common huge page sizes: 2Mi and 1Gi.  Sizes vary on other
architectures, but the idea is the same.  In order to use huge pages,
application must write code that is aware of them.  Transparent Huge Pages (THP)
attempts to automate the management of huge pages without application knowledge,
but they have limitations.  In particular, they are limited to 2Mi page sizes.
THP might lead to performance degradation on nodes with high memory utilization
or fragmentation due to defragmenting efforts of THP, which can lock memory
pages. For this reason, some applications may be designed to (or recommend)
usage of pre-allocated huge pages instead of THP.

Managing memory is hard, and unfortunately, there is no one-size fits all
solution for all applications.

## Scope

This proposal only includes pre-allocated huge pages configured on the node by
the administrator at boot time or by manual dynamic allocation.  It does not
discuss how the cluster could dynamically attempt to allocate huge pages in an
attempt to find a fit for a pod pending scheduling.  It is anticipated that
operators may use a variety of strategies to allocate huge pages, but we do not
anticipate the kubelet itself doing the allocation.  Allocation of huge pages
ideally happens soon after boot time.

This proposal defers issues relating to NUMA.

## Use Cases

The class of applications that benefit from huge pages typically have
- A large memory working set
- A sensitivity to memory access latency

Example applications include:
- database management systems (MySQL, PostgreSQL, MongoDB, Oracle, etc.)
- Java applications can back the heap with huge pages using the
  `-XX:+UseLargePages` and `-XX:LagePageSizeInBytes` options.
- packet processing systems (DPDK)

Applications can generally use huge pages by calling
- `mmap()` with `MAP_ANONYMOUS | MAP_HUGETLB` and use it as anonymous memory
- `mmap()` a file backed by `hugetlbfs`
- `shmget()` with `SHM_HUGETLB` and use it as a shared memory segment (see Known
  Issues).

1. A pod can use huge pages with any of the prior described methods.
1. A pod can request huge pages.
1. A scheduler can bind pods to nodes that have available huge pages.
1. A quota may limit usage of huge pages.
1. A limit range may constrain min and max huge page requests.

## Feature Gate

The proposal introduces huge pages as an Alpha feature.

It must be enabled via the `--feature-gates=HugePages=true` flag on pertinent
components pending graduation to Beta.

## Node Specification

Huge pages cannot be overcommitted on a node.

A system may support multiple huge page sizes.  It is assumed that most nodes
will be configured to primarily use the default huge page size as returned via
`grep Hugepagesize /proc/meminfo`.  This defaults to 2Mi on most Linux systems
unless overridden by `default_hugepagesz=1g` in kernel boot parameters.

For each supported huge page size, the node will advertise a resource of the
form `hugepages-<hugepagesize>`.  On Linux, supported huge page sizes are
determined by parsing the `/sys/kernel/mm/hugepages/hugepages-{size}kB`
directory on the host. Kubernetes will expose a `hugepages-<hugepagesize>`
resource using binary notation form. It will convert `<hugepagesize>` into the
most compact binary notation using integer values.  For example, if a node
supports `hugepages-2048kB`, a resource `hugepages-2Mi` will be shown in node
capacity and allocatable values. Operators may set aside pre-allocated huge
pages that are not available for user pods similar to normal memory via the
`--system-reserved` flag.

There are a variety of huge page sizes supported across different hardware
architectures.  It is preferred to have a resource per size in order to better
support quota.  For example, 1 huge page with size 2Mi is orders of magnitude
different than 1 huge page with size 1Gi.  We assume gigantic pages are even
more precious resources than huge pages.

Pre-allocated huge pages reduce the amount of allocatable memory on a node. The
node will treat pre-allocated huge pages similar to other system reservations
and reduce the amount of `memory` it reports using the following formula:

```
[Allocatable] = [Node Capacity] - 
 [Kube-Reserved] - 
 [System-Reserved] - 
 [Pre-Allocated-HugePages * HugePageSize] -
 [Hard-Eviction-Threshold]
```

The following represents a machine with 10Gi of memory.  1Gi of memory has been
reserved as 512 pre-allocated huge pages sized 2Mi.  As you can see, the
allocatable memory has been reduced to account for the amount of huge pages
reserved.

```
apiVersion: v1
kind: Node
metadata:
  name: node1
...
status:
  capacity:
    memory: 10Gi
    hugepages-2Mi: 1Gi
  allocatable:
    memory: 9Gi
    hugepages-2Mi: 1Gi
...  
```

## Pod Specification

A pod must make a request to consume pre-allocated huge pages using the resource
`hugepages-<hugepagesize>` whose quantity is a positive amount of memory in
bytes.  The specified amount must align with the `<hugepagesize>`; otherwise,
the pod will fail validation.  For example, it would be valid to request
`hugepages-2Mi: 4Mi`, but invalid to request `hugepages-2Mi: 3Mi`.

The request and limit for `hugepages-<hugepagesize>` must match.  Similar to
memory, an application that requests `hugepages-<hugepagesize>` resource is at
minimum in the `Burstable` QoS class.

If a pod consumes huge pages via `shmget`, it must run with a supplemental group
that matches `/proc/sys/vm/hugetlb_shm_group` on the node.  Configuration of
this group is outside the scope of this specification.

Initially, a pod may not consume multiple huge page sizes in a single pod spec.
Attempting to use `hugepages-2Mi` and `hugepages-1Gi` in the same pod spec will
fail validation.  We believe it is rare for applications to attempt to use
multiple huge page sizes. This restriction may be lifted in the future with
community presented use cases.  Introducing the feature with this restriction
limits the exposure of API changes needed when consuming huge pages via volumes.

In order to consume huge pages backed by the `hugetlbfs` filesystem inside the
specified container in the pod, it is helpful to understand the set of mount
options used with `hugetlbfs`.  For more details, see "Using Huge Pages" here:
https://www.kernel.org/doc/Documentation/vm/hugetlbpage.txt

```
mount -t hugetlbfs \
	-o uid=<value>,gid=<value>,mode=<value>,pagesize=<value>,size=<value>,\
	min_size=<value>,nr_inodes=<value> none /mnt/huge
```

The proposal recommends extending the existing `EmptyDirVolumeSource` to satisfy
this use case.  A new `medium=HugePages` option would be supported.  To write
into this volume, the pod must make a request for huge pages. The `pagesize`
argument is inferred from the `hugepages-<hugepagesize>` from the resource
request.  If in the future, multiple huge page sizes are supported in a single
pod spec, we may modify the `EmptyDirVolumeSource` to provide an optional page
size.  The existing `sizeLimit` option for `emptyDir` would restrict usage to
the minimum value specified between `sizeLimit` and the sum of huge page limits
of all containers in a pod. This keeps the behavior consistent with memory
backed `emptyDir` volumes whose usage is ultimately constrained by the pod
cgroup sandbox memory settings.  The `min_size` option is omitted as its not
necessary.  The `nr_inodes` mount option is omitted at this time in the same
manner it is omitted with `medium=Memory` when using `tmpfs`.

The following is a sample pod that is limited to 1Gi huge pages of size 2Mi. It
can consume those pages using `shmget()` or via `mmap()` with the specified
volume.

```
apiVersion: v1
kind: Pod
metadata:
  name: example
spec:
  containers:
...
    volumeMounts:
    - mountPath: /hugepages
      name: hugepage
    resources:
      requests:
        hugepages-2Mi: 1Gi
      limits:
        hugepages-2Mi: 1Gi
  volumes:
  - name: hugepage
    emptyDir:
      medium: HugePages
```

## CRI Updates

The `LinuxContainerResources` message should be extended to support specifying
huge page limits per size.  The specification for huge pages should align with
opencontainers/runtime-spec.

see:
https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#huge-page-limits

The CRI changes are required before promoting this feature to Beta.

## Cgroup Enforcement

To use this feature, the `--cgroups-per-qos` must be enabled.  In addition, the
`hugetlb` cgroup must be mounted.

The `kubepods` cgroup is bounded by the `Allocatable` value.

The QoS level cgroups are left unbounded across all huge page pool sizes.

The pod level cgroup sandbox is configured as follows, where `hugepagesize` is
the system supported huge page size(s).  If no request is made for huge pages of
a particular size, the limit is set to 0 for all supported types on the node.

```
pod<UID>/hugetlb.<hugepagesize>.limit_in_bytes = sum(pod.spec.containers.resources.limits[hugepages-<hugepagesize>])
```

If the container runtime supports specification of huge page limits, the
container cgroup sandbox will be configured with the specified limit.

The `kubelet` will ensure the `hugetlb` has no usage charged to the pod level
cgroup sandbox prior to deleting the pod to ensure all resources are reclaimed.

## Limits and Quota

The `ResourceQuota` resource will be extended to support accounting for
`hugepages-<hugepagesize>` similar to `cpu` and `memory`.  The `LimitRange`
resource will be extended to define min and max constraints for `hugepages`
similar to `cpu` and `memory`.

## Scheduler changes

The scheduler will need to ensure any huge page request defined in the pod spec
can be fulfilled by a candidate node.

## cAdvisor changes

cAdvisor will need to be modified to return the number of pre-allocated huge
pages per page size on the node.  It will be used to determine capacity and
calculate allocatable values on the node.

## Roadmap

### Version 1.8

Initial alpha support for huge pages usage by pods.

### Version 1.9

Resource Quota support. Limit Range support. Beta support for huge pages
(pending community feedback)

## Known Issues

### Huge pages as shared memory

For the Java use case, the JVM maps the huge pages as a shared memory segment
and memlocks them to prevent the system from moving or swapping them out.

There are several issues here:
- The user running the Java app must be a member of the gid set in the
  `vm.huge_tlb_shm_group` sysctl
- sysctl `kernel.shmmax` must allow the size of the shared memory segment
- The user's memlock ulimits must allow the size of the shared memory segment
- `vm.huge_tlb_shm_group` is not namespaced.

### NUMA

NUMA is complicated.  To support NUMA, the node must support cpu pinning,
devices, and memory locality.  Extending that requirement to huge pages is not
much different.  It is anticipated that the `kubelet` will provide future NUMA
locality guarantees as a feature of QoS.  In particular, pods in the
`Guaranteed` QoS class are expected to have NUMA locality preferences.

