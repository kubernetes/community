---
kep-number: TBD
title: IPVS Load Balancing Mode in Kubernetes
status: implemented
authors:
    - "@rramkumar1"
owning-sig: sig-network
reviewers:
  - "@thockin"
  - "@m1093782566"
approvers:
  - "@thockin"
  - "@m1093782566"
editor:
  - "@thockin"
  - "@m1093782566"
creation-date: 2018-03-21
---

# IPVS Load Balancing Mode in Kubernetes

**Note: This is a retroactive KEP. Credit goes to @m1093782566, @haibinxie, and @quinton-hoole for all information & design in this KEP.**

**Important References: https://github.com/kubernetes/community/pull/692/files**

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non\-goals](#non-goals)
* [Proposal](#proposal)
  * [Kube-Proxy Parameter Changes](#kube-proxy-parameter-changes)
  * [Build Changes](#build-changes)
  * [Deployment Changes](#deployment-changes)
  * [Design Considerations](#design-considerations)
    * [IPVS service network topology](#ipvs-service-network-topology)
    * [Port remapping](#port-remapping)
    * [Falling back to iptables](#falling-back-to-iptables)
    * [Supporting NodePort service](#supporting-nodeport-service)
    * [Supporting ClusterIP service](#supporting-clusterip-service)
    * [Supporting LoadBalancer service](#supporting-loadbalancer-service)
    * [Session Affinity](#session-affinity)
    * [Cleaning up inactive rules](#cleaning-up-inactive-rules)
    * [Sync loop pseudo code](#sync-loop-pseudo-code)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks](#drawbacks)
* [Alternatives](#alternatives)

## Summary

We are building a new implementation of kube proxy built on top of IPVS (IP Virtual Server).

## Motivation

As Kubernetes grows in usage, the scalability of its resources becomes more and more
important. In particular, the scalability of services is paramount to the adoption of Kubernetes
by developers/companies running large workloads. Kube Proxy, the building block of service routing
has relied on the battle-hardened iptables to implement the core supported service types such as
ClusterIP and NodePort. However, iptables struggles to scale to tens of thousands of services because
it is designed purely for firewalling purposes and is based on in-kernel rule chains. On the
other hand, IPVS is specifically designed for load balancing and uses more efficient data structures
under the hood. For more information on the performance benefits of IPVS vs. iptables, take a look
at these [slides](https://docs.google.com/presentation/d/1BaIAywY2qqeHtyGZtlyAp89JIZs59MZLKcFLxKE6LyM/edit?usp=sharing).

### Goals

* Improve the performance of services

### Non-goals

None

### Challenges and Open Questions [optional]

None


## Proposal

### Kube-Proxy Parameter Changes

***Parameter: --proxy-mode***
In addition to existing userspace and iptables modes, IPVS mode is configured via --proxy-mode=ipvs. In the initial implementation, it implicitly uses IPVS [NAT](http://www.linuxvirtualserver.org/VS-NAT.html) mode.

***Parameter: --ipvs-scheduler***
A new kube-proxy parameter will be added to specify the IPVS load balancing algorithm, with the parameter being --ipvs-scheduler. If it’s not configured, then round-robin (rr) is default value. If it’s incorrectly configured, then kube-proxy will exit with error message.
  * rr: round-robin
  * lc: least connection
  * dh: destination hashing
  * sh: source hashing
  * sed: shortest expected delay
  * nq: never queue
For more details, refer to http://kb.linuxvirtualserver.org/wiki/Ipvsadm

In future, we can implement service specific scheduler (potentially via annotation), which has higher priority and overwrites the value.

***Parameter: --cleanup-ipvs***
Similar to the --cleanup-iptables parameter, if true, cleanup IPVS configuration and IPTables rules that are created in IPVS mode.

***Parameter: --ipvs-sync-period***
Maximum interval of how often IPVS rules are refreshed (e.g. '5s', '1m'). Must be greater than 0.

***Parameter: --ipvs-min-sync-period***
Minimum interval of how often the IPVS rules are refreshed (e.g. '5s', '1m'). Must be greater than 0.


### Build Changes

No changes at all. The IPVS implementation is built on [docker/libnetwork](https://godoc.org/github.com/docker/libnetwork/ipvs) IPVS library, which is a pure-golang implementation and talks to kernel via socket communication.

### Deployment Changes

IPVS kernel module installation is beyond Kubernetes. It’s assumed that IPVS kernel modules are installed on the node before running kube-proxy. When kube-proxy starts, if the proxy mode is IPVS, kube-proxy would validate if IPVS modules are installed on the node. If it’s not installed, then kube-proxy will fall back to the iptables proxy mode.

### Design Considerations

#### IPVS service network topology

We will create a dummy interface and assign all kubernetes service ClusterIP's to the dummy interface (default name is `kube-ipvs0`). For example,

```shell
# ip link add kube-ipvs0 type dummy
# ip addr
...
73: kube-ipvs0: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN qlen 1000
    link/ether 26:1f:cc:f8:cd:0f brd ff:ff:ff:ff:ff:ff

#### Assume 10.102.128.4 is service Cluster IP
# ip addr add 10.102.128.4/32 dev kube-ipvs0
...
73: kube-ipvs0: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN qlen 1000
    link/ether 1a:ce:f5:5f:c1:4d brd ff:ff:ff:ff:ff:ff
    inet 10.102.128.4/32 scope global kube-ipvs0
       valid_lft forever preferred_lft forever
```

Note that the relationship between a Kubernetes service and an IPVS service is `1:N`. Consider a Kubernetes service that has more than one access IP. For example, an External IP type service has 2 access IP's (ClusterIP and External IP). Then the IPVS proxier will create 2 IPVS services - one for Cluster IP and the other one for External IP.

The relationship between a Kubernetes endpoint and an IPVS destination is `1:1`.
For instance, deletion of a Kubernetes service will trigger deletion of the corresponding IPVS service and address bound to dummy interface.


#### Port remapping

There are 3 proxy modes in ipvs - NAT (masq), IPIP and DR. Only NAT mode supports port remapping. We will use IPVS NAT mode in order to supporting port remapping. The following example shows ipvs mapping service port `3080` to container port `8080`.

```shell
# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0         
  -> 10.244.1.237:8080            Masq    1      0          0     

```

#### Falling back to iptables

IPVS proxier will employ iptables in doing packet filtering, SNAT and supporting NodePort type service. Specifically, ipvs proxier will fall back on iptables in the following 4 scenarios.

* kube-proxy start with --masquerade-all=true
* Specify cluster CIDR in kube-proxy startup
* Load Balancer Source Ranges is specified for LB type service
* Support NodePort type service

And, IPVS proxier will maintain 5 kubernetes-specific chains in nat table

- KUBE-POSTROUTING
- KUBE-MARK-MASQ
- KUBE-MARK-DROP

`KUBE-POSTROUTING`, `KUBE-MARK-MASQ`, ` KUBE-MARK-DROP` are maintained by kubelet and ipvs proxier won't create them. IPVS proxier will make sure chains `KUBE-MARK-SERVICES` and `KUBE-NODEPORTS` exist in its sync loop.

**1. kube-proxy start with --masquerade-all=true**

If kube-proxy starts with `--masquerade-all=true`, the IPVS proxier will masquerade all traffic accessing service ClusterIP, which behaves same as what iptables proxier does.
Suppose there is a serivice with Cluster IP `10.244.5.1` and port `8080`:

```shell
# iptables -t nat -nL

Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-POSTROUTING  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */

Chain KUBE-POSTROUTING (1 references)
target     prot opt source               destination         
MASQUERADE  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service traffic requiring SNAT */ mark match 0x4000/0x4000

Chain KUBE-MARK-DROP (0 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x8000

Chain KUBE-MARK-MASQ (6 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x4000

Chain KUBE-SERVICES (2 references)
target     prot opt source               destination         
KUBE-MARK-MASQ  tcp  -- 0.0.0.0/0        10.244.5.1            /* default/foo:http cluster IP */ tcp dpt:8080
```

**2. Specify cluster CIDR in kube-proxy startup**

If kube-proxy starts with `--cluster-cidr=<cidr>`, the IPVS proxier will masquerade off-cluster traffic accessing service ClusterIP, which behaves same as what iptables proxier does.
Suppose kube-proxy is provided with the cluster cidr `10.244.16.0/24`, and service Cluster IP is `10.244.5.1` and port is `8080`:

```shell
# iptables -t nat -nL

Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-POSTROUTING  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */

Chain KUBE-POSTROUTING (1 references)
target     prot opt source               destination         
MASQUERADE  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service traffic requiring SNAT */ mark match 0x4000/0x4000

Chain KUBE-MARK-DROP (0 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x8000

Chain KUBE-MARK-MASQ (6 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x4000

Chain KUBE-SERVICES (2 references)
target     prot opt source               destination         
KUBE-MARK-MASQ  tcp  -- !10.244.16.0/24        10.244.5.1            /* default/foo:http cluster IP */ tcp dpt:8080
```

**3. Load Balancer Source Ranges is specified for LB type service**

When service's `LoadBalancerStatus.ingress.IP` is not empty and service's `LoadBalancerSourceRanges` is specified, IPVS proxier will install iptables rules which looks like what is shown below.

Suppose service's `LoadBalancerStatus.ingress.IP` is `10.96.1.2` and service's `LoadBalancerSourceRanges` is `10.120.2.0/24`:

```shell
# iptables -t nat -nL

Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-POSTROUTING  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */

Chain KUBE-POSTROUTING (1 references)
target     prot opt source               destination         
MASQUERADE  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service traffic requiring SNAT */ mark match 0x4000/0x4000

Chain KUBE-MARK-DROP (0 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x8000

Chain KUBE-MARK-MASQ (6 references)
target     prot opt source               destination         
MARK       all  --  0.0.0.0/0            0.0.0.0/0            MARK or 0x4000

Chain KUBE-SERVICES (2 references)
target     prot opt source       destination         
ACCEPT  tcp  -- 10.120.2.0/24    10.96.1.2       /* default/foo:http loadbalancer IP */ tcp dpt:8080
DROP    tcp  -- 0.0.0.0/0        10.96.1.2       /* default/foo:http loadbalancer IP */ tcp dpt:8080
```

**4. Support NodePort type service**

Please check the section below.

#### Supporting NodePort service

For supporting NodePort type service, iptables will recruit the existing implementation in the iptables proxier. For example,

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
Type:			NodePort
IP:			    10.101.28.148
Port:			http	3080/TCP
NodePort:		http	31604/TCP
Endpoints:		172.17.0.2:80
Session Affinity:	None

# iptables -t nat -nL

[root@100-106-179-225 ~]# iptables -t nat -nL
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
KUBE-SERVICES  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */

Chain KUBE-SERVICES (2 references)
target     prot opt source               destination         
KUBE-MARK-MASQ  tcp  -- !172.16.0.0/16        10.101.28.148        /* default/nginx-service:http cluster IP */ tcp dpt:3080
KUBE-SVC-6IM33IEVEEV7U3GP  tcp  --  0.0.0.0/0            10.101.28.148        /* default/nginx-service:http cluster IP */ tcp dpt:3080
KUBE-NODEPORTS  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service nodeports; NOTE: this must be the last rule in this chain */ ADDRTYPE match dst-type LOCAL

Chain KUBE-NODEPORTS (1 references)
target     prot opt source               destination         
KUBE-MARK-MASQ  tcp  --  0.0.0.0/0            0.0.0.0/0            /* default/nginx-service:http */ tcp dpt:31604
KUBE-SVC-6IM33IEVEEV7U3GP  tcp  --  0.0.0.0/0            0.0.0.0/0            /* default/nginx-service:http */ tcp dpt:31604

Chain KUBE-SVC-6IM33IEVEEV7U3GP (2 references)
target     prot opt source               destination
KUBE-SEP-Q3UCPZ54E6Q2R4UT  all  --  0.0.0.0/0            0.0.0.0/0            /* default/nginx-service:http */
Chain KUBE-SEP-Q3UCPZ54E6Q2R4UT (1 references)
target     prot opt source               destination         
KUBE-MARK-MASQ  all  --  172.17.0.2           0.0.0.0/0            /* default/nginx-service:http */
DNAT  
```

#### Supporting ClusterIP service

When creating a ClusterIP type service, IPVS proxier will do 3 things:

* make sure dummy interface exists in the node
* bind service cluster IP to the dummy interface
* create an IPVS service whose address corresponds to the Kubernetes service Cluster IP.

For example,

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
Type:			ClusterIP
IP:			    10.102.128.4
Port:			http	3080/TCP
Endpoints:		10.244.0.235:8080,10.244.1.237:8080
Session Affinity:	None

# ip addr
...
73: kube-ipvs0: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN qlen 1000
    link/ether 1a:ce:f5:5f:c1:4d brd ff:ff:ff:ff:ff:ff
    inet 10.102.128.4/32 scope global kube-ipvs0
       valid_lft forever preferred_lft forever

# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0         
  -> 10.244.1.237:8080            Masq    1      0          0   
```

### Support LoadBalancer service

IPVS proxier will NOT bind LB's ingress IP to the dummy interface. When creating a LoadBalancer type service, ipvs proxier will do 4 things:

- Make sure dummy interface exists in the node
- Bind service cluster IP to the dummy interface
- Create an ipvs service whose address corresponding to kubernetes service Cluster IP
- Iterate LB's ingress IPs, create an ipvs service whose address corresponding LB's ingress IP

For example,

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
IP:			    10.102.128.4
Port:			http	3080/TCP
Endpoints:		10.244.0.235:8080
Session Affinity:	None

#### Only bind Cluter IP to dummy interface
# ip addr
...
73: kube-ipvs0: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN qlen 1000
    link/ether 1a:ce:f5:5f:c1:4d brd ff:ff:ff:ff:ff:ff
    inet 10.102.128.4/32 scope global kube-ipvs0
       valid_lft forever preferred_lft forever

#### Suppose LB's ingress IPs {10.96.1.2, 10.93.1.3}. IPVS proxier will create 1 ipvs service for cluster IP and 2 ipvs services for LB's ingree IP. Each ipvs service has its destination.
# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0           
TCP  10.96.1.2:3080 rr  
  -> 10.244.0.235:8080            Masq    1      0          0   
TCP  10.96.1.3:3080 rr  
  -> 10.244.0.235:8080            Masq    1      0          0   
```

Since there is a need of supporting access control for `LB.ingress.IP`. IPVS proxier will fall back on iptables. Iptables will drop any packet which is not from `LB.LoadBalancerSourceRanges`. For example,

```shell
# iptables -A KUBE-SERVICES -d {ingress.IP} --dport {service.Port} -s {LB.LoadBalancerSourceRanges} -j ACCEPT
```

When the packet reach the end of chain, ipvs proxier will drop it.

```shell
# iptables -A KUBE-SERVICES -d {ingress.IP} --dport {service.Port} -j KUBE-MARK-DROP
```

### Support Only NodeLocal Endpoints

Similar to iptables proxier, when a service has the "Only NodeLocal Endpoints" annotation,  ipvs proxier will only proxy traffic to endpoints in the local node.

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
IP:			    10.102.128.4
Port:			http	3080/TCP
Endpoints:		10.244.0.235:8080, 10.244.1.235:8080
Session Affinity:	None

#### Assume only endpoint 10.244.0.235:8080 is in the same host with kube-proxy

#### There should be 1 destination for ipvs service.
[root@SHA1000130405 home]# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0               
```

#### Session affinity

IPVS support client IP session affinity (persistent connection). When a service specifies session affinity, the IPVS proxier will set a timeout value (180min=10800s by default) in the IPVS service. For example,

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
IP:			    10.102.128.4
Port:			http	3080/TCP
Session Affinity:	ClientIP

# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.102.128.4:3080 rr persistent 10800
```

#### Cleaning up inactive rules

It seems difficult to distinguish if an IPVS service is created by the IPVS proxier or other processes. Currently we assume IPVS rules will be created only by the IPVS proxier on a node, so we can clear all IPVSrules on a node. We should add warnings in documentation and flag comments.

#### Sync loop pseudo code

Similar to the iptables proxier, the IPVS proxier will do a full sync loop in a configured period. Also, each update on a Kubernetes service or endpoint will trigger an IPVS service or destination update. For example,

* Creating a Kubernetes service will trigger creating a new IPVS service.
* Updating a Kubernetes service(for instance, change session affinity) will trigger updating an existing IPVS service.
* Deleting a Kubernetes service will trigger deleting an IPVS service.
* Adding an endpoint for a Kubernetes service will trigger adding a destination for an existing IPVS service.
* Updating an endpoint for a Kubernetes service will trigger updating a destination for an existing IPVS service.
* Deleting an endpoint for a Kubernetes service will trigger deleting a destination for an existing IPVS service.

Any IPVS service or destination updates will send an update command to kernel via socket communication, which won't take a service down.

The sync loop pseudo code is shown below:

```go
func (proxier *Proxier) syncProxyRules() {
	When service or endpoint update, begin sync ipvs rules and iptables rules if needed.
    ensure dummy interface exists, if not, create one.
    for svcName, svcInfo := range proxier.serviceMap {
      // Capture the clusterIP.
      construct ipvs service from svcInfo
      Set session affinity flag and timeout value for ipvs service if specified session affinity
      bind Cluster IP to dummy interface
      call libnetwork API to create ipvs service and destinations

      // Capture externalIPs.
      if externalIP is local then hold the svcInfo.Port so that can install ipvs rules on it
      construct ipvs service from svcInfo
      Set session affinity flag and timeout value for ipvs service if specified session affinity
      call libnetwork API to create ipvs service and destinations

      // Capture load-balancer ingress.
	    for _, ingress := range svcInfo.LoadBalancerStatus.Ingress {
		    if ingress.IP != "" {
          if len(svcInfo.LoadBalancerSourceRanges) != 0 {
            install specific iptables
          }
          construct ipvs service from svcInfo
          Set session affinity flag and timeout value for ipvs service if specified session affinity
          call libnetwork API to create ipvs service and destinations
        }
      }

      // Capture nodeports.
      if svcInfo.NodePort != 0 {
		    fall back on iptables, recruit existing iptables proxier implementation
      }

      call libnetwork API to clean up legacy ipvs services which is inactive any longer
      unbind service address from dummy interface
      clean up legacy iptables chains and rules
    }
}
```

## Graduation Criteria

### Beta -> GA

The following requirements should be met before moving from Beta to GA. It is
suggested to file an issue which tracks all the action items.

- [ ] Testing
    - [ ] 48 hours of green e2e tests. 
    - [ ] Flakes must be identified and filed as issues.
    - [ ] Integrate with scale tests and. Failures should be filed as issues.
- [ ] Development work
    - [ ] Identify all pending changes/refactors. Release blockers must be prioritized and fixed.
    - [ ] Identify all bugs. Release blocking bugs must be identified and fixed.
- [ ] Docs 
    - [ ] All user-facing documentation must be updated.

### GA -> Future

__TODO__

## Implementation History

**In chronological order**

1. https://github.com/kubernetes/kubernetes/pull/46580

2. https://github.com/kubernetes/kubernetes/pull/52528

3. https://github.com/kubernetes/kubernetes/pull/54219

4. https://github.com/kubernetes/kubernetes/pull/57268

5. https://github.com/kubernetes/kubernetes/pull/58052


## Drawbacks [optional]

None

## Alternatives [optional]

None
