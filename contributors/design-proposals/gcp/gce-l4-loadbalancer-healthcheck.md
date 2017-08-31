# GCE L4 load-balancers' health checks for nodes

## Goal
Set up health checks for GCE L4 load-balancer to ensure it is only
targeting healthy nodes.

## Motivation
On cloud providers which support external load balancers, setting the
type field to "LoadBalancer" will provision a L4 load-balancer for the
service ([doc](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer)),
which load-balances traffic to k8s nodes. As of k8s 1.6, we don't
create health check for L4 load-balancer by default, which means all
traffic will be forwarded to any one of the nodes blindly.

This is undesired in cases:
- k8s components including kubelet dead on nodes. Nodes will be flipped
to unhealthy after a long propagation (~40s), even if we remove nodes
from target pool at that point it is too slow.
- kube-proxy dead on nodes while kubelet is still alive. Requests will
be continually forwarded to nodes that may not be able to properly route
traffic.

For now, the only case health check will be created is for
[OnlyLocal Service](https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer).
We should have a node-level health check for load balancers that are used
by non-OnlyLocal services.

## Design
Healthchecking the kube-proxys seems to be the best choice:
- kube-proxy runs on every nodes and it is the pivot for service traffic
routing.
- Port 10249 on nodes is currently used for both kube-proxy's healthz and
pprof.
- We already have a similar mechanism for healthchecking OnlyLocal services
in kube-proxy.

The plan is to enable health check on all LoadBalancer services (if use GCP
as cloud provider).

## Implementation
kube-proxy
- Separate healthz from pprof (/metrics) to use a different port and bind it
to 0.0.0.0. As we will only allow traffic from load-balancer source IPs, this
wouldn't be a big security concern.
- Make healthz check timestamp in iptables mode while always returns "ok" in
other modes.

GCE cloud provider (through kube-controller-manager)
- Manage `k8s-l4-healthcheck` firewall and healthcheck resources.
These two resources should be shared among all non-OnlyLocal LoadBalancer
services.
- Add a new flag to pipe in the healthz port num as it is configurable on
kube-proxy.

Version skew:
- Running higher version master (with L4 health check feature enabled) with
lower version nodes (without kube-proxy exposing healthz port) should fall
back to the original behavior (no health check).
- Rollback shouldn't be a big issue. Even if health check is left on Network
load-balancer, it will fail on all nodes and fall back to blindly forwarding
traffic.
