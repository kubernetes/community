# Support specifying NodePort IP range

Author: @m1093782566

# Objective

This document proposes creating a option for kube-proxy to specify NodePort IP range.

# Background

NodePort type service gives developers the freedom to set up their own load balancers, to expose one or more nodes’ IPs directly. The service will be visible as the nodes's IPs. For now, the NodePort addresses are the IPs from all available interfaces.

With iptables magic, all the IPs whose `ADDRTYPE` matches `dst-type LOCAL` will be taken as the address of NodePort, which might look like,

```shell
Chain KUBE-SERVICES (2 references)
target     prot opt source               destination         
KUBE-NODEPORTS  all  --  0.0.0.0/0            0.0.0.0/0            /* kubernetes service nodeports; NOTE: this must be the last rule in this chain */ ADDRTYPE match dst-type LOCAL
```
By default, kube-proxy accepts everything from NodePort without any filter. It can be a problem for nodes which has both public and private NICs, and people only want to provide a service in private network and avoid exposing any internal service on the public IPs. 

# Proposal

This proposal builds off of earlier requests to [[proxy] Listening on a specific IP for nodePort ](https://github.com/kubernetes/kubernetes/issues/21070), but proposes that we should find a way to tell kube-proxy what the NodePort IP blocks are instead of a single IP.

## Create new kube-proxy configuration option

There should be an admin option to kube-proxy for specifying which IP to NodePort. The option is a list of IP blocks, say `--nodeport-addresses`. These IP blocks as a parameter to select the interfaces where nodeport works. In case someone would like to expose a service on localhost for local visit and some other interfaces for particular purpose, an array of IP blocks would do that. People can populate it from their private subnets the same on every node.

The `--nodeport-addresses` is defaulted to `0.0.0.0/0`, which means select all available interfaces and is compliance with current NodePort behaviour.

If people set the `--nodeport-addresses` option to "127.0.0.0/8", kube-proxy will only select the loopback interface for NodePort.

If people set the `--nodeport-addresses` option to "default-route", kube-proxy will select the "who has the default route" interfaces. It's the same heuristic we use for `--advertise-address` in kube-apiserver and others. 

If people provide a non-zero IP block for `--nodeport-addresses`, kube-proxy will filter that down to just the IPs that applied to the node. 

So, the following values for `--nodeport-addresses` are all valid:

```
0.0.0.0/0
127.0.0.0/8
default-route
127.0.0.1/32,default-route
127.0.0.0/8,192.168.0.0/16
```


And an empty string for `--nodeport-addresses` is considered as invalid.

> NOTE: There is already a option `--bind-address`, but it has nothing to do with nodeport and  we need IP blocks instead of single IP.

kube-proxy will periodically refresh proxy rules based on the list of IP blocks specified by `--nodeport-addresses`, in case of something like DHCP. 

For example, if IP address of `eth0` changes from `172.10.1.2` to `172.10.2.100` and user specifies `172.10.0.0/16` for `--advertise-address`. Kube-proxy will make sure proxy rules `-d 172.10.0.0/16` exist. 

However, if IP address of `eth0` changes from `172.10.1.2` to `192.168.3.4` and user only specifies `172.10.0.0/16` for `--advertise-address`. Kube-proxy will NOT create proxy rules for `192.168.3.4` unless `eth0` has the default route. 

When refer to DHCP user case, network administrator usually reserves a RANGE of IP addresses for the DHCP server. So, IP address change will always fall in an IP range in DHCP scenario. That's to say an IP address of a interface will not change from `172.10.1.2` to `192.168.3.4` in our example.

## Kube-proxy implementation support

The implementation is simple.

### iptables

iptables support specify CIDR in the destination parameter(`-d`), e.g. `-d 192.168.0.0/16`.

For the special `default-route` case, we should use `-i` option in iptables command, e.g. `-i eth0`.

### Linux userspace

Same as iptables.

### ipvs

Create IPVS virtual services one by one according to provided node IPs, which is almost same as current behaviour(fetch all IPs from host).

### Window userspace

Create multiple goroutines, each goroutine listens on a specific node IP to serve NodePort.

### winkernel 

Need to specify node IPs [here](https://github.com/kubernetes/kubernetes/blob/master/pkg/proxy/winkernel/proxier.go#L1053) - current behaviour is leave the VIP to be empty to automatically select the node IP.
