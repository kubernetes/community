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
last-updated: 2018-10-05
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

#### Daemonset, Service, Listen Interface for caching agent

The caching agent daemonset runs in hostNetwork mode in kube-system namespace with a Priority Class of “system-node-critical”. It listens for dns requests on a dummy interface created on the host. A separate ip address is assigned to this dummy interface, so that requests to kube-dns or any other custom service are not incorrectly intercepted by the caching agent. This ip address is obtained by creating a nodelocaldns service, with no endpoints. Kubelet takes the service name as an argument `--localDNS=<namespace>/<svc name>`, looks up the ip address and populates pods' resolv.conf with this value instead of clusterDNS. Each cluster node will have a dummy interface with this service ip assigned to it. This IP will be handled specially because of the NOTRACK rules described in the section below.

#### iptables NOTRACK

NOTRACK rules are added for connections to and from the nodelocal dns service ip. Additional rules in FILTER table to whitelist these connections, since the INPUT and OUTPUT chains have a default DROP policy.

A dnscache nanny, similar to the [dnsmasq nanny](https://github.com/kubernetes/dns/tree/master/pkg/dnsmasq) will create the dummy interface and iptables rules. The nanny gets the nodelocal dns service ip address by querying kube-dns for the service name. The Daemonset runs in privileged securityContext since the nanny container needs to create this dummy interface and add iptables rules.
 The nanny will also periodically ensure that the iptables rules are present. The resource usage for periodic iptables check needs to be measured. We can reduce the frequency of the checks by having the nanny query the dns cache and checking/adding rules only if the query fails.


[Proposal presentation](https://docs.google.com/presentation/d/1c43cZqbVhGAlw3dSNQIOGuvQmDfKaA2yiAPRoYpa6iY), also shared at the sig-networking meeting on 2018-10-04

Slide 5 has a diagram showing how the new dns cache fits in.

#### Choice of caching agent

The current plan is to run [Unbound dns server](https://www.nlnetlabs.nl/projects/unbound/about/) by default, based on these benchmark[ tests](https://github.com/kubernetes/perf-tests/tree/master/dns).

Will rerun benchmark tests with coreDNS after SO_REUSEPORT support and measure QPS.

Tests were run on a 1.9.7 cluster with 2 nodes, using Unbound 1.7.3 and coreDNS 1.2. Unbound QPS was 3600(caching)/3400(no caching), coreDNS 950(caching)/500(no caching)


Based on the prototype/test results, these are the recommended defaults: 
CPU request: 50m
Memory Limit : 25m  
Resource usage and QPS to be recomputed after metrics collection has been added.

#### Metrics

Unbound metrics can be extracted from the program unbound-control. This can be run from the nanny container.
Unbound-control-setup is required to setup keys/certs for SSL.


### Risks and Mitigations

* Having the pods query the nodelocal cache introduces a single point of failure.   
This is mitigated by assigning a PriorityClass of "system-node-critical", so it is always running. The nanny process will periodically check if the dns agent is running.
Populating both the nodelocal cache ip address and kube-dns ip address in resolv.conf is not a reliable option. Depending on underlying implementation, this can result in kube-dns being queried only if cache ip does not repond, or both queried simultaneously.


## Graduation Criteria
TODO

## Rollout Plan
This feature will be launched with Alpha support in the first release. Master versions v1.13 and above will deploy the new add-on. Node versions v1.13 and above will have kubelet code to modify pods' resolv.conf. Nodes running older versions will run the nodelocal daemonset, but it will not be used. The user can specify a custom dnsConfig to use this local cache dns server.

## Implementation History

* 2018-10-05 - Creation of the KEP

## Drawbacks [optional]

Additional resource consumption for the Daemonset might not be necessary for clusters with low DNS QPS needs. 


## Alternatives [optional]

The listen ip address for the dns cache could be a link-local ip address, if there is a reliable way to reserve a link-local subnet. This will still require iptables rules for connectivity, in order to skip conntrack.


Instead of just a dns-cache, a full-fledged kube-dns instance can be run on all nodes. This will consume much more resources since each instance will also watch Services and Endpoints.
