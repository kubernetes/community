# Known Scalability Issues with Kubernetes Services

_by Shyam JVS, Google Inc (with inputs from Brian Grant, Wojciech Tyczynski & Dan Winship)_

**June 2018**

This document serves as a catalog of issues we've known/discovered with kubernetes services as of June 2018, focusing on their scalability/performance. The purpose of the document is to make the information common knowledge for the community, so we can work together towards improving it. Listing them below in no particular order.


## Iptables packet processing performance when large number of service ports exist

#### Issue

Iptables can be slow in packet processing when a large number of services exist. As the number of service ports increases, the KUBE-SERVICES chain gets longer and since it’s also evaluated very frequently that can bring down the performance. There have been some recent improvements in this area ([#56164](https://github.com/kubernetes/kubernetes/pull/56164), [#57461](https://github.com/kubernetes/kubernetes/pull/57461), [#60306](https://github.com/kubernetes/kubernetes/pull/60306)) - but we still need to measure and come up with safe bounds for number of services that can be 'decently' supported (10k seems to be a reasonable estimate from [past discussions](https://github.com/kubernetes/kubernetes/issues/48938#issue-243000172)). Also, further improvements might be required based on those results.

#### Relevant issues/links

- [#48938](https://github.com/kubernetes/kubernetes/issues/48938)

#### Possible Solution(s)

- Moving from iptables to IPVS may help here if that works stably. We have the IPVS alternative [implemented](https://github.com/kubernetes/kubernetes/pull/46580), but it still hasn’t gone to GA.
- For a long time the official kernel upstream answer for this issue was that nftables was going to solve all iptables-related scalability problems (both with packet processing and with rule changes), and there's now an out-of-tree nftables kube-proxy backend too ([#62720](https://github.com/kubernetes/kubernetes/issues/62720)).
- There's also an alternative plan to fix the packet processing speed problems by rewriting iptables to use eBPF inside the kernel (see - https://cilium.io/blog/2018/04/17/why-is-the-kernel-community-replacing-iptables ).


## Slow/failing iptables-restore operations when large number of rules exist

#### Issue

When we’re running services worth a large number of backends (> 100k) on our large-cluster scalability tests, we’re noticing that kube-proxy is timing out while trying to do iptables-restore due to failing to acquire lock over iptables. There are at least two parts to this problem:

- If a process tried to do a very large iptables-restore, and other processes on the system were also doing iptables-related things, the iptables-restore might get restarted multiple times (inside the kernel) before completing, which could then cause another iptables operation running in another goroutine to time out before it ever got the xtables lock.
- The `iptables-restore` implementation is such that it grabs the lock before it parses its input, and simply parsing tens of thousands of iptables rules takes a noticeable amount of time. So if two iptables commands start at the same time, the first might grab the lock, and then start parsing its input, and burn through half of the other iptables command's `--wait` time before it even gets to the point of passing the rules off to the kernel.

TODO: Find if there's any issue arising here from k8s side (relevant bug linked below).

#### Relevant issues/links

- [#48107](https://github.com/kubernetes/kubernetes/issues/48107)

#### Possible Solution(s)

For the first problem:

- There's [a claim](https://github.com/kubernetes/kubernetes/issues/48107#issuecomment-398081930) that upgrading to linux kernel version 4.15 may fix this. Though we haven't verified this (current COS versions using <= 4.14 kernel version).

For the second problem:

- It's not easy to fix it because of the way that `iptables-restore` was written to share code with the main `iptables` binary, and no one is working on fixing it because the official plan is to move to nftables instead. If the world ends up moving to iptables-over-eBPF rather than nftables then this will probably need to be fixed at some point.

In general:

- The issue requires more debugging, but it may be possible that we are able to (at least partially) mitigate the problem with some tweaks here and there on kubernetes side.
- Decide to support a smaller number of backends for kubernetes officially if it can't be fixed.


## Apiserver spending too much cpu/memory in duplicate Endpoints serializations

#### Issue

Currently while serving watches, the apiserver is deep-copying deserialized endpoints objects from etcd and serializing it once for each kube-proxy watch it serves (which is 5k watches in our largest clusters).

#### Relevant issues/links

- [#55779](https://github.com/kubernetes/kubernetes/issues/55779#issuecomment-354452477)
- [#48938](https://github.com/kubernetes/kubernetes/issues/48938)
- [#58050](https://github.com/kubernetes/kubernetes/issues/58050)

#### Possible Solution(s)

- Avoid those duplicate serializations if/when possible ([#60067](https://github.com/kubernetes/kubernetes/pull/60067) was an attempt to do this, but it didn't seem to help too much).
- Moving from 'endpoints' -> 'endpoint' API change (discussed in detail in below issues) would mitigate this problem as the serialization load would decrease with smaller objects.


## Endpoints traffic is quadratic in the number of endpoints

#### Issue

Endpoints object for a service contains all the individual endpoints of that service. As a result, whenever even a single pod in a service is added/updated/deleted, the whole endpoints object (which includes even the other endpoints that didn't change) is re-computed, written to storage and sent to all readers. Not being able to efficiently read/update individual endpoint changes can lead to (for e.g during rolling upgrade of a service) endpoints operations that are quadratic in the number of its elements. If you consider watches in the picture (there's one from each kube-proxy), the situation becomes even worse as the quadratic traffic gets multiplied further with number of watches (usually equal to #nodes in the cluster).

Overall, this is a serious performance drawback affecting multiple components in the control-plane (apiserver, etcd, endpoints-controller, kube-proxy). The current endpoints API (which was designed at a very early stage of kubernetes when people weren't really thinking too much about scalability/performance) makes it hard to solve this problem without introducing breaking changes.

#### Relevant issues/links

- [#47787](https://github.com/kubernetes/kubernetes/issues/47787)
- [#24552](https://github.com/kubernetes/kubernetes/issues/24552)
- [#24553](https://github.com/kubernetes/kubernetes/issues/24553)
- [#8190](https://github.com/kubernetes/kubernetes/issues/8190)

#### Possible Solution(s)

- Redesign endpoints API with (at least) the following properties:
  - Lives in an API group
  - Not coupled to service resource lifecycle
  - Can track both pod and non-pod endpoints
  - Separate API object for each individual endpoint
  - Has a 'Ready' property as part of the object
  - Port ranges and other features we’ve put off for a long time


## Define/measure performance SLIs for endpoints propagation

#### Issue

We need to measure the e2e endpoints propagation latency, i.e the time from when a 
service is created until its endpoints are populated, given that all of the associated pods 
are already running and ready. We also need to come up with reasonable SLOs for the 
same and verify that they’re satisfied at scale.

#### Relevant issues/links

- [#10436](https://github.com/kubernetes/kubernetes/issues/10436)
- This [doc on networking SLIs](https://docs.google.com/document/d/1kudI8uILO13ySkpE5HgNHvIGn-yXVQatpyM3YTVVUis) proposes one way of formulating the SLI

#### Possible Solution(s)

- Measure the SLI defined in the above doc and come up with a reasonable SLO


## Master election shouldn’t consume unnecessary watch bandwidth

#### Issue

Currently we store the leader lock for master components like the scheduler in the corresponding endpoints object. This has the undesirable side effect of sending a notification down the kube-proxy <-> master (and kube-dns <-> master) watch every second, which is not actually required.

#### Relevant issues/links

- [#34627](https://github.com/kubernetes/kubernetes/issues/34627)

#### Possible Solutions(s)

- Avoid overloading endpoints as a lock API. We instead want to move to using the [new Lease API](https://github.com/kubernetes/kubernetes/pull/64246) (ETA 1.12). One potential challenge here is performing upgrades of HA clusters that are using the old endpoints-based locking approach to this new API approach without breaking them.
