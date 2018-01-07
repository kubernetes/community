# Support IPVS Load Balancing Mode in Kubernetes

**Authors**

* dujun (@m1093782566)
* haibin xie (@haibinxie)

## Abstract

This proposal summarizes what’s expected in alpha version IPVS load balancing support, includes Kubernetes user behavior changes, build and deployment changes, and the test validation planned. 

## Motivation

There’s been a lot discussion and voice of enabling IPVS as in-cluster service load balancing mode.  IPVS performs better than iptables, meanwhile supports more sophisticated load balancing algorithms than iptables (least load, least connections, locality, weighted) as well as other useful features (e.g. health checking, retries etc).

IPVS is an alternative of iptables as load balancer, it’s assumed reader of this proposal is familiar with IPTables load balancer mode. 

##  Kubernetes behavior change

### Changes to kube-proxy startup parameter

**1. Proxy mode**

In addition to existing userspace and iptables mode, ipvs mode is configured via `--proxy-mode=ipvs`.

**2. Expose different  load balancing  algorithms**

A new kube-proxy parameter `--ipvs-scheduler` will be added to specify IPVS load balancing algorithm. Below is list of supported values, if it’s not configured `rr` is default value, if it’s incorrectly configured kube-proxy will exit with error message.

- rr: round-robin

- lc: least connection

- dh: destination hashing

- sh: source hashing

- sed: shortest expected delay

- nq: never queue

For more details about it, refer to [http://kb.linuxvirtualserver.org/wiki/Ipvsadm](http://kb.linuxvirtualserver.org/wiki/Ipvsadm)

**3. clean up inactive rules**

In order to cleaning up inactive rules(includes iptables rules and ipvs rules), we will introduce a new  kube-proxy parameter `--cleanup-proxyrules ` and mark the older `--cleanup-iptables` deprecated. 

**4. ipvs maximum sync period**

Introduce a new parameter ` --ipvs-sync-period` to specify the maximum interval of how often ipvs rules are refreshed (e.g. '5s', '1m', '2h22m').  Must be greater than 0.

**5. ipvs minimum sync period**

Introduce a new parameter ` --ipvs-min-sync-period`  to the minimum interval of how often the ipvs rules can be refreshed as endpoints and services change (e.g. '5s', '1m', '2h22m').

### Changes to build

No changes at all. The IPVS implementation is built on [docker/libnetwork](https://godoc.org/github.com/docker/libnetwork/ipvs) ipvs library, which is a pure-golang implementation and talks to kernel via socket communication.

### Changes to deployment

IPVS kernel module installation is beyond kubernetes, it’s assumed IPVS kernel modules are installed on the node before running kube-proxy. When kube-proxy starts, if proxy mode is IPVS kube-proxy would validate if IPVS modules are installed on the node, if it’s not installed kube-proxy will fall back to userspace proxy mode.

## Other design considerations

### Clean up inactive rules

It seems difficult to distinguish if an ipvs service is created by ipvs proxier or other processes. Currently we assume ipvs rules will be created only by ipvs proxier in a kubernetes node, so we can clear all ipvs rules in a kubernetes node. Probably we should add warnings in document or flag comment. 

### When fall back on iptables

IPVS proxier will employ iptables in doing packet filtering, SNAT and supporting NodePort type service. Specifically, ipvs proxier will fall back on iptables in the following 4 scenarios.

* kube-proxy start with --masquerade-all=true
* Specify cluster CIDR in kube-proxy startup
* Load Balancer Source Ranges is specified for LB type service
* Support NodePort type service

And, IPVS proxier will maintain 5 kubernetes-specific chains in nat table

- KUBE-POSTROUTING 
- KUBE-MARK-MASQ
- KUBE-MARK-DROP
- KUBE-MARK-SERVICES
- KUBE-NODEPORTS

`KUBE-POSTROUTING`, `KUBE-MARK-MASQ`, ` KUBE-MARK-DROP` are maintained by kubelet and ipvs proxier won't create them. IPVS proxier will make sure chains `KUBE-MARK-SERVICES` and `KUBE-NODEPORTS` exist in its sync loop.

**1. kube-proxy start with --masquerade-all=true**

If kube-proxy start with `--masquerade-all=true`, ipvs proxier will masquerade all traffic accessing service Cluster IP, which behaves same as what iptables proxier does. Suppose there is a serivice with Cluster IP `10.244.5.1` and port `8080`, then the iptables installed by ipvs proxier should be like what is shown below.

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

If kube-proxy start with `--cluster-cidr=<cidr>`, ipvs proxier will masquerade off-cluster traffic accessing service Cluster IP, which behaves same as what iptables proxier does. Suppose kube-proxy is provided with the cluster cidr `10.244.16.0/24`, and service Cluster IP is `10.244.5.1` and port is `8080`, then the iptables installed by ipvs proxier should be like what is shown below.

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

When service's `LoadBalancerStatus.ingress.IP` is not empty and service's `LoadBalancerSourceRanges` is specified, ipvs proxier will install iptables which looks like what is shown below. 

Suppose service's `LoadBalancerStatus.ingress.IP` is `10.96.1.2` and service's `LoadBalancerSourceRanges` is `10.120.2.0/24`.

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

### Support NodePort type service

For supporting NodePort type service, iptables will recruit the exsiting implementation in iptables proxier. For example, 

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
DNAT       tcp  --  0.0.0.0/0            0.0.0.0/0            /* default/nginx-service:http */ tcp to:172.17.0.2:80
```

### IPVS Service Network topology

We will create a dummy interface and assign all kubernetes service Cluster IPs to the dummy interface (default name is `kube-ipvs0`). For example,

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

Note that the relationship between kubernetes service and ipvs service is `1:N`。The address of ipvs service corresponding to service's access IP, such as Cluster IP, external IP and LB.ingress.IP. If a kubernetes service has more than one access IP, for example, an External IP type service has 2 access IP(ClusterIP and External IP), then ipvs proxier will create 2 ipvs serivices - one for Cluster IP and the other one for External IP.

The relationship between kubernetes endpoint and ipvs destination is `1:1`. For instance,

Delete a kubernetes service will trigger deletion of corresponding ipvs service and address bound to dummy interface.

### Port remapping 

There are 3 proxy modes in ipvs - NAT(masq), IPIP and DR. Only NAT mode support port remapping. In alpha version, it implicitly uses IPVS NAT mode in order to supporting port remapping. The following example shows ipvs mapping service port `3080` to container port `8080`.

```shell
# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0         
  -> 10.244.1.237:8080            Masq    1      0          0     
```

### Support Cluster IP type service

When creating a Cluster IP type service, ipvs proxier will do 3 things:

* make sure dummy interface exists in the node


* bind service cluster IP to the dummy interface

* create an ipvs service whose address corresponding to kubernetes service Cluster IP.

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

### Session affinity

IPVS support client IP session affinity, it's called persistent connection. When a service specify session affinity, ipvs proxier will set a timeout value(180min=10800s by default) in the ipvs service. For example,

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


### Support service External IP

IPVS proxier will NOT bind service External IP to the dummy interface. When creating an External IP type service, ipvs proxier will do 4 things:

- make sure dummy interface exists in the node


- bind service cluster IP to the dummy interface
- create an ipvs service whose address corresponding to kubernetes service Cluster IP
- create an ipvs service whose address corresponding to kubernetes service External IP

For example,

```shell
# kubectl describe svc nginx-service
Name:			nginx-service
...
IP:			    10.102.128.4
External IPs:   100.106.89.164
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

#### There should be 2 ipvs services, one for Cluter IP and the other one for External IP. Each ipvs service has its destination.
# ipvsadm -ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn     
TCP  10.102.128.4:3080 rr
  -> 10.244.0.235:8080            Masq    1      0          0           
TCP  100.106.89.164:3080 rr  
  -> 10.244.0.235:8080            Masq    1      0          0            
```

### Support LoadBalancer type service

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
  -> 10.244.0.235:8080            Masq    1      0          0   
TCP  10.96.1.3:3080 rr  
  -> 10.244.0.235:8080            Masq    1      0          0   
```

Since there is a need of supporting access control for `LB.ingress.IP`. IPVS proxier will fall back on iptables. Iptables will drop any packet which is not from `LB.LoadBalancerSourceRanges`.  For example,

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

### Sync loop pseudo code

Similar to iptables proxier, ipvs proxier will do full sync loop in a configured period. Besides, each update on kubernetes service or endpoint will trigger an ipvs service or destination update. For example,

* Creating a kubernetes service will trigger creating a new ipvs service.


* Updating a kubernetes service(for instance, change session affinity) will trigger updating an existing ipvs service.
* Deleting a kubernetes service will trigger deleting an ipvs service.
* Adding an endpoint for a kubernetes service will trigger adding a destination for an existing ipvs service.
* Updating an endpoint for a kubernetes service will trigger updating a destination for an existing ipvs service.
* Deleting an endpoint for a kubernetes service will trigger deleting a destination for an existing ipvs service.



Any ipvs service or destination updates will send update command to kernel via socket communication, which won't take a service down.

The sync loop pseudo code is shown below.

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

### Clean up legacy rules when switch proxy mode

when userspace proxy start up, it will clean iptables rules and ipvs rules.

when iptables proxy start up, it will clean userspace rules and ipvs rules.

when ipvs proxy start up, it will clean userspace rules and iptables rules.

## Test validation

**Functionality tests, all below traffic should be reachable**
```
pod -> pod, same VM
pop -> pod, other VM
pod -> own VM, own hostPort
pod -> own VM, other hostPort
pod -> other VM, other hostPort

pod -> own VM
pod -> other VM
pod -> internet
pod -> http://metadata

VM -> pod, same VM
VM -> pod, other VM
VM -> same VM hostPort
VM -> other VM hostPort

pod -> own clusterIP, hairpin
pod -> own clusterIP, same VM, other pod, no port remap
pod -> own clusterIP, same VM, other pod, port remap
pod -> own clusterIP, other VM, other pod, no port remap
pod -> own clusterIP, other VM, other pod, port remap
pod -> other clusterIP, same VM, no port remap
pod -> other clusterIP, same VM, port remap
pod -> other clusterIP, other VM, no port remap
pod -> other clusterIP, other VM, port remap
pod -> own node, own nodePort, hairpin
pod -> own node, own nodePort, policy=local
pod -> own node, own nodePort, same VM
pod -> own node, own nodePort, other VM
pod -> own node, other nodePort, policy=local
pod -> own node, other nodePort, same VM
pod -> own node, other nodePort, other VM
pod -> other node, own nodeport, policy=local
pod -> other node, own nodeport, same VM
pod -> other node, own nodeport, other VM
pod -> other node, other nodeport, policy=local
pod -> other node, other nodeport, same VM
pod -> other node, other nodeport, other VM
pod -> own external LB, no remap, policy=local
pod -> own external LB, no remap, same VM
pod -> own external LB, no remap, other VM
pod -> own external LB, remap, policy=local
pod -> own external LB, remap, same VM
pod -> own external LB, remap, other VM

VM -> same VM nodePort, policy=local
VM -> same VM nbodePort, same VM
VM -> same VM nbodePort, other VM
VM -> other VM nodePort, policy=local
VM -> other VM nbodePort, same VM
VM -> other VM nbodePort, other VM

VM -> external LB

public -> nodeport, policy=local
public -> nodeport, policy=global
public -> external LB, no remap, policy=local
public -> external LB, no remap, policy=global
public -> external LB, remap, policy=local
public -> external LB, remap, policy=global

public -> nodeport, manual backend
public -> external LB, manual backend
```
