---
kep-number: 30
title: NodeLocal DNS Cache
authors:
  - "@prameshj"
owning-sig: sig-network
participating-sigs:
  - sig-network
reviewers:
  - "@thockin"
  - "@bowei"
  - "@johnbelamaric"
  - "@sdodson"
approvers:
  - "@thockin"
  - "@bowei"
editor: TBD
creation-date: 2018-10-05
last-updated: 2018-10-30
status: provisional
---

# NodeLocal DNS Cache

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Rollout Plan](#rollout-plan)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

This proposal aims to improve DNS performance by running a dns caching agent on cluster nodes as a Daemonset. In today's architecture, pods in ClusterFirst DNS mode reach out to a kube-dns serviceIP for DNS queries. This is translated to a kube-dns endpoint via iptables rules added by kube-proxy. With this new architecture, pods will reach out to the dns caching agent running on the same node, thereby avoiding iptables DNAT rules and connection tracking. The local caching agent will query kube-dns for cache misses of cluster hostnames(cluster.local suffix by default).


## Motivation

* With the current DNS achitecture, it is possible that pods with the highest DNS QPS have to reach out to a different node, if there is no local kube-dns instance.  
Having a local cache will help improve the latency in such scenarios. 

* Skipping iptables DNAT and connection tracking will help reduce [conntrack races](https://github.com/kubernetes/kubernetes/issues/56903) and avoid UDP DNS entries filling up conntrack table.

* Connections from local caching agent to kube-dns can be upgraded to TCP. TCP conntrack entries will be removed on connection close in contrast with UDP entries that have to timeout ([default](https://www.kernel.org/doc/Documentation/networking/nf_conntrack-sysctl.txt) `nf_conntrack_udp_timeout` is 30 seconds)

* Upgrading DNS queries from UDP to TCP would reduce tail latency attributed to dropped UDP packets and DNS timeouts usually up to 30s (3 retries + 10s timeout). Since the nodelocal cache listens for UDP DNS queries, applications don't need to be changed.

* Metrics & visibility into dns requests at a node level.

* Neg caching can be re-enabled, thereby reducing number of queries to kube-dns.

* There are several open github issues proposing a local DNS Cache daemonset and scripts to run it:
	* [https://github.com/kubernetes/kubernetes/issues/7470#issuecomment-248912603](https://github.com/kubernetes/kubernetes/issues/7470#issuecomment-248912603)

	* [https://github.com/kubernetes/kubernetes/issues/32749](https://github.com/kubernetes/kubernetes/issues/32749)

	* [https://github.com/kubernetes/kubernetes/issues/45363](https://github.com/kubernetes/kubernetes/issues/45363)


This shows that there is interest in the wider Kubernetes community for a solution similar to the proposal here. 


### Goals

Being able to run a dns caching agent as a Daemonset and get pods to use the local instance. Having visibility into cache stats and other metrics.

### Non-Goals

* Providing a replacement for kube-dns/CoreDNS.
* Changing the underlying protocol for DNS (e.g. to gRPC)

## Proposal

A nodeLocal dns cache runs on all cluster nodes. This is managed as an add-on, runs as a Daemonset. All pods using clusterDNS will now talk to the nodeLocal cache, which will query kube-dns in case of cache misses in cluster's configured DNS suffix and for all reverse lookups(in-addr.arpa and ip6.arpa). User-configured stubDomains will be passed on to this local agent.  
The node's resolv.conf will be used by this local agent for all other cache misses. One benefit of doing the non-cluster lookups on the nodes from which they are happening, rather than the kube-dns instances, is better use of per-node DNS resources in cloud. For instance, in a 10-node cluster with 3 kube-dns instances, the 3 nodes running kube-dns will end up resolving all external hostnames and can exhaust QPS quota. Spreading the queries across the 10 nodes will help alleviate this.

#### Daemonset and Listen Interface for caching agent

The caching agent daemonset runs in hostNetwork mode in kube-system namespace with a Priority Class of “system-node-critical”. It listens for dns requests on a dummy interface created on the host. A separate ip address is assigned to this dummy interface, so that requests to kube-dns or any other custom service are not incorrectly intercepted by the caching agent. This will be a link-local ip address selected by the user. Each cluster node will have this dummy interface. This ip address will be passed on to kubelet via the --cluster-dns flag, if the feature is enabled.

The selected link-local IP will be handled specially because of the NOTRACK rules described in the section below.

#### iptables NOTRACK

NOTRACK rules are added for connections to and from the nodelocal dns ip. Additional rules in FILTER table to whitelist these connections, since the INPUT and OUTPUT chains have a default DROP policy.

The nodelocal cache process will create the dummy interface and iptables rules . It gets the nodelocal dns ip as a parameter, performs setup and listens for dns requests. The Daemonset runs in privileged securityContext since it needs to create this dummy interface and add iptables rules.
 The cache process will also periodically ensure that the dummy interface and iptables rules are present, in the background. Rules need to be checked in the raw table and filter table. Rules in these tables do not grow with number of valid services. Services with no endpoints will have rules added in filter table to drop packets destined to these ip. The resource usage for periodic iptables check was measured by creating 2k services with no endpoints and running the nodelocal caching agent. Peak memory usage was 20Mi for the caching agent when it was responding to queries along with the periodic checks. This was measured using `kubectl top` command. More details on the testing are in the following section.

[Proposal presentation](https://docs.google.com/presentation/d/1c43cZqbVhGAlw3dSNQIOGuvQmDfKaA2yiAPRoYpa6iY), also shared at the sig-networking meeting on 2018-10-04

Slide 5 has a diagram showing how the new dns cache fits in.

#### Choice of caching agent

The current plan is to run CoreDNS by default. Benchmark [ tests](https://github.com/kubernetes/perf-tests/tree/master/dns) were run using [Unbound dns server](https://www.nlnetlabs.nl/projects/unbound/about/) and CoreDNS. 2 more tests were added to query for 20 different services and to query several external hostnames.

Tests were run on a 1.9.7 cluster with 2 nodes on GCE, using Unbound 1.7.3 and CoreDNS 1.2.3.
Resource limits for nodelocaldns daemonset was CPU - 50m, Memory 25Mi

Resource usage and QPS were measured with a nanny process for Unbound/CoreDNS plugin adding iptables rules and ensuring that the rules exist, every minute.

Caching was minimized in Unbound by setting:
msg-cache-size: 0
rrset-cache-size: 0
msg-cache-slabs:1
rrset-cache-slabs:1
Previous tests did not set the last 2 and there were quite a few unexpected cache hits.

Caching was disabled in CoreDNS by skipping the cache plugin from Corefile.

These are the results when dnsperf test was run with no QPS limit. In this mode, the tool  sends queries until they start timing out.

| Test Type             | Program | Caching | QPS  |
|-----------------------|---------|---------|------|
| Multiple services(20) | CoreDNS | Yes     | 860  |
| Multiple services(20) | Unbound | Yes     | 3030 |
|                       |         |         |      |
| External queries      | CoreDNS | Yes     | 213  |
| External queries      | Unbound | Yes     | 115  |
|                       |         |         |      |
| Single Service        | CoreDNS | Yes     | 834  |
| Single Service        | Unbound | Yes     | 3287 |
|                       |         |         |      |
| Single NXDomain       | CoreDNS | Yes     | 816  |
| Single NXDomain       | Unbound | Yes     | 3136 |
|                       |         |         |      |
| Multiple services(20) | CoreDNS | No      | 859  |
| Multiple services(20) | Unbound | No      | 1463 |
|                       |         |         |      |
| External queries      | CoreDNS | No      | 180  |
| External queries      | Unbound | No      | 108  |
|                       |         |         |      |
| Single Service        | CoreDNS | No      | 818  |
| Single Service        | Unbound | No      | 2992 |
|                       |         |         |      |
| Single NXDomain       | CoreDNS | No      | 827  |
| Single NXDomain       | Unbound | No      | 2986 |


Peak memory usage was ~20 Mi for both Unbound and CoreDNS.

For the single service and single NXDomain query, Unbound still had cache hits since caching could not be completely disabled.

CoreDNS QPS was twice as much as Unbound for external queries. They were mostly unique hostnames from this file - [ftp://ftp.nominum.com/pub/nominum/dnsperf/data/queryfile-example-current.gz](ftp://ftp.nominum.com/pub/nominum/dnsperf/data/queryfile-example-current.gz)

When multiple cluster services were queried with cache misses, Unbound was better(1463 vs 859), but not by a large factor.

Unbound performs much better when all requests are cache hits.

CoreDNS will be the local cache agent in the first release, after considering these reasons:

*  Better QPS numbers for external hostname queries
*  Single process, no need for a separate nanny process
*  Prometheus metrics already available, also we can get per zone stats. Unbound gives consolidated stats.
*  Easier to make changes to the source code

 It is possible to run any program as caching agent by modifying the daemonset and configmap spec. Publishing an image with Unbound DNS can be added as a follow up.

Based on the prototype/test results, these are the recommended defaults: 
CPU request: 50m
Memory Limit : 25m  

CPU request can be dropped to a smaller value if QPS needs are lower.

#### Metrics

Per-zone metrics will be available via the metrics/prometheus plugin in CoreDNS.


### Risks and Mitigations

Having the pods query the nodelocal cache introduces a single point of failure.

* This is mitigated by having a livenessProbe to periodically ensure DNS is working. In case of upgrades, the recommendation is to drain the node before starting to upgrade the local instance. The user can also configure [customPodDNS](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-config) pointing to clusterDNS ip for pods that cannot handle DNS disruption during upgrade.

* The Daemonset is assigned a PriorityClass of "system-node-critical", to ensure it is not evicted.

* Populating both the nodelocal cache ip address and kube-dns ip address in resolv.conf is not a reliable option. Depending on underlying implementation, this can result in kube-dns being queried only if cache ip does not repond, or both queried simultaneously.


## Graduation Criteria
TODO

## Rollout Plan
This feature will be launched with Alpha support in the first release. Master versions v1.13 and above will deploy the new add-on. Node versions v1.13 and above will have kubelet code to modify pods' resolv.conf. Nodes running older versions will run the nodelocal daemonset, but it will not be used. The user can specify a custom dnsConfig to use this local cache dns server.

## Implementation History

* 2018-10-05 - Creation of the KEP
* 2018-10-30 - Follow up comments and choice of cache agent

## Drawbacks [optional]

Additional resource consumption for the Daemonset might not be necessary for clusters with low DNS QPS needs. 


## Alternatives [optional]

* The listen ip address for the dns cache could be a service ip. This ip address is obtained by creating a nodelocaldns service, with same endpoints as the clusterDNS service. Using the same endpoints as clusterDNS helps reduce DNS downtime in case of upgrades/restart. When no other special handling is provided, queries to the nodelocaldns ip will be served by kube-dns/CoreDNS pods. Kubelet takes the service name as an argument `--cluster-dns-svc=<namespace>/<svc name>`, looks up the ip address and populates pods' resolv.conf with this value instead of clusterDNS.
This approach works only for iptables mode of kube-proxy. This is because kube-proxy creates a dummy interface bound to all service IPs in ipvs mode and ipvs rules are added to load-balance between endpoints. The packet seems to get dropped if there are no endpoints. If there are endpoints, adding iptables rules does not bypass the ipvs loadbalancing rules.

* A nodelocaldns service can be created with a hard requirement of same-node endpoint, once we have [this](https://github.com/kubernetes/community/pull/2846) supported. All the pods in the nodelocaldns daemonset will be endpoints, the one running locally will be selected. iptables rules to NOTRACK connections can still be added, in order to skip DNAT in the iptables kube-proxy implementation.

* Instead of just a dns-cache, a full-fledged kube-dns instance can be run on all nodes. This will consume much more resources since each instance will also watch Services and Endpoints.
