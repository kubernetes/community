# Topology aware routing of services

Author: @m1093782566

## Objective

Figure out a generic way to implement the "local service" route, say "topology aware routing of service". 

Locality is defined by user, it can be any topology-related thing. "Local" means the "same topology level", e.g. same node, same rack, same failure zone, same failure region, same cloud provider etc.

## GOAL

A generic way to support topology aware routing of services in arbitrary topological domains, e.g. node, rack, zone, region, etc. whatever.

## Non-goal

Scheduler spreading to implement this sort of topology guarantee.

## Use cases

* Logging agents such as fluentd. Deploy fluentd as DaemonSet and applications only need to communicate with the fluentd in the same node.
* For a sharded service that keeps per-node local information in each shard.
* Authenticating proxies such as [aws-es-proxy](https://github.com/kopeio/aws-es-proxy).
* In container identity wg, being able to give daemonset pods a unique identity per host is on the 2018 plan, and ensuring local pods can communicate to local node services securely is a key goal there. -- from @smarterclayton
* Regional data costs in multi-AZ setup - for instance, in AWS, with a multi-AZ setup, half of the traffic will switch AZ, incurring regional data Transfer costs, whereas if something was local, it wouldn't hit the network.
* Performance benefit (node local/rack local) is lower latency/higher bandwidth.

## Background

It's a pain point for multi-zone clusters deployment since cross-zone network traffic being charged, while in-zone is not. In addition, cross-node traffic may carry sensitive metadata from other nodes. Therefore, users always prefer the service backends that close to them, e.g. same zone, rack and host etc. for security, performance and cost concerns.

Kubernetes scheduler can constraining a pod to only be able to run on particular nodes/zones. However, Kubernetes service proxy just randomly picks an available backend for service routing and this one can be very far from the user, so we need a topology-aware service routing solution in Kubernetes. Basically, to find the nearest service backend. In other words, allowing people to configure if ALWAY reach a to local service backend. In this way, they can reduce network latency, improve security, save money and so on. However, because topology is arbitrary, zone, region, rack, generator, whatever, who knows? We should allow arbitrary locality.

`ExternalTrafficPolicy` was added in v1.4, but only for NodePort and external LB traffic. NodeName was added to `EndpointAddress` to allow kube-proxy to filter local endpoints for various future purposes.

Based on our experience of advanced routing setup and recent demo of enabling this feature in Kubernetes, this document would like to introduce a more generic way to support arbitrary service topology.

## Proposal

This proposal builds off of earlier requests to [use local pods only for kube-proxy loadbalancing](https://github.com/kubernetes/kubernetes/issues/7433) and [node-local service proposal](https://github.com/kubernetes/kubernetes/pull/28637). But, this document proposes that not only the particular "node-local" user case should be taken care, but also a more generic way should be figured out.

Locality is an "user-defined" thing. When we set topology key "hostname" for service, we expect node carries different node labels on the key "hostname".

Users can control the level of topology. For example, if someone run logging agent as a daemonset, he can set the "hard" topology requirement for same-host. If "hard" is not met, then just return "service not available". 

And if someone set a "soft" topology requirement for same-host, say he "preferred" same-host endpoints and can accept other hosts when for some reasons local service's backend is not available on some host.

If multiple endpoints satisfy the "hard" or "soft" topology requirement, we will randomly pick one by default. Routing decision is expected to be implemented in L3/4 VIP level such as kube proxy.

## Implementation details

### API changes

#### New type ServicePolicy

The user need a way to declare which service is "local service" and what is the "topology key" of "local service".

This will be accomplished through a new type object `ServicePolicy`.
Endpoint(s) with specify label will be selected by label selector in
`ServicePolicy`, and `ServicePolicy` will declare the topology policy for those endpoints.

Cluster administrators can configure what services are "local" and what topological they prefer via `ServicePolicy`. `ServicePolicy` is a namespace-scope resource and is strict optional. We can configure policies other than topological reference in `ServicePolicy`, but this proposal will not cover them.

```go
type ServicePolicy struct {
  TypeMeta
  ObjectMeta

  // specification of the topology policy of this ServicePolicy
  Spec TopologyPolicySpec
}

type TopologyPolicySpec struct {
  // ServiceSelector select the service to which this TopologyPolicy object applies.
  // One service only can be selected by single ServicePolicy, in this case, the topology rules are combined additively.
  // This field is NOT optional an empty ServiceSelector will result in err.
  ServiceSelector metav1.LabelSelector `json:"endPointSelector" protobuf:"bytes,1,opt,name=podSelector"`

  // topology is used to achieve "local" service in a given topology level.
  // User can control what ever topology level they want.
  // +optional
  Topology ServiceTopology `json:"topology" protobuf:"bytes,1,opt,name=topology"`
}

// Defines a service topolgoy information.
type ServiceTopology struct {
  // Valid values for mode are "ignored", "required", "preferred".
  // "ignored" is the default value and the associated topology key will have no effect.
  // "required" is the "hard" requirement for topology key and an example would be  “only visit service backends in the same zone”.
  // If the topology requirements specified by this field are not met, the LB, such as kube-proxy will not pick endpoints for the service.
  // "preferred" is the "soft" requirement for topology key and an example would be
  // "prefer to visit service backends in the same rack, but OK to other racks if none match"
  // +optional
  Mode ServicetopologyMode `json:"mode" protobuf:"bytes,1,opt,name=mode"`

  // key is the key for the node label that the system uses to denote
  // such a topology domain. There are some built-in topology keys, e.g.
  // kubernetes.io/hostname, failure-domain.beta.kubernetes.io/zone and failure-domain.beta.kubernetes.io/region etc.
  // The built-in topology keys can be good examples and we recommend users switch to a similar mode for portability, but it's NOT enforced.
  // Users can define whatever topolgoy key they like since topology is arbitrary.
  // +optional
  Key string `json:"key" protobuf:"bytes,2,opt,name=key"`
}
```

An example of `ServicePolicy`:

```yaml
kind: ServicePolicy
metadata:
  name: service-policy-example
  namespace: test
spec:
  serviceSelector:
    matchLabels:
      app: test
  topology:
    key: kubernetes.io/hostname
    mode: required
```

<<<<<<< HEAD
In our example, services in namespace `foo` with label `app=bar` will be chosen. Requests to these services will be routed only to backends on nodes with the same value for `kubernetes.io/hostname` as the originating pod's node. If we want the "same host", probably every host should carry unique `kubernetes.io/hostname` labels.

We can configure multiple `ServicePolicy` targeting a single service. In this case, the service will carry multiple topology requirements and the relationships of all the requirements are logical `AND`, for example,

```yaml
kind: ServicePolicy
metadata:
  name: service-policy-example-1
  namespace: foo
spec:
  serviceSelector:
    matchLabels:
      app: bar
  topology:
    key: kubernetes.io/region
    mode: required    
---
kind: ServicePolicy
metadata:
  name: service-policy-example-2
  namespace: foo
spec:
  serviceSelector:
    matchLabels:
      app: bar
  topology:
    key: kubernetes.io/switch
    mode: required   
```

In our example, services in namespace `foo` with label `app=bar` will be dominated by both `service-policy-example-1` and `service-policy-example-2`. Requests to these services will be routed only to backends that satisfy both same region and same switch as kube-proxy.


#### Endpoints API changes

Although `NodeName` was already added to `EndpointAddress`, we want `Endpoints` to carry more node's topology informations so that allowing more topology levels other than hostname.

So, create a new `Topology` field in `Endpoints.Subsets.Addresses` for identifying what topology domain the endpoints pod exists, e.g. what host, rack, zone, region etc. In other words, copy the topology-related labels of node hosting the endpoint to `EndpointAddress.Topology`.

```go
type EndpointAddress struct {
  // labels of node hosting the endpoint
  Topology map[string]string
}
```

## Endpoints Controller changes

Endpoint Controller will populate the `Topology` for each `EndpointAddress`. We want `EndpointAddress.Topology` to tell the LB, such as kube-proxy what topological domain(e.g. host, rack, zone, region etc.) the endpoints is in.

Endpoints controller will need to watch two extra resources: ServicePolicy and Nodes. Watching ServicePolicy for knowing what services have topological preferences. Watching Nodes for knowing labels of node hosting the endpoint and copy the node labels referenced in the service spec's topology constraints to EndpointAddress.

Endpoints Controller will maintain two extra caches: `NodeToPodsCache` and `ServiceToPoliciesCache`. 
`NodeToPodsCache` maps the node's name to the pods running on it. Node's add, delete and labels' change will trigger `NodeToPodsCache` reindex.

`ServiceToPoliciesCache` maps the Service's namespaced name to all of its ServicePolicys.

So, the new logic of endpoint controller might like:

```go
go watch Node, ServicePolicy
// In each sync loop, given a service, sync its endpoints
for i, pod := range service backends; do
  servicePolicys := ServiceToPoliciesCache[service.Name]
  node := nodeCache[pod.Spec.NodeName]
  // endpointAddress := &v1.EndpointAddress {}
  // Copy all topology-related labels of node hosting endpoint to endpoint
  // We can only include node labels referenced in the service spec's topology constraints
  for _, servicePolicy := range servicePolicys; do
    topoKey := servicePolicy.Topology.Key
    endpointAddress.Topology[topoKey] = node.Labels[topoKey]
  done
  endpoints.Subsets[i].Addresses = endpointAddress
done
```

## Kube-proxy changes

Kube-proxy will respect topology keys for each service, so kube-proxy on different nodes may create different proxy rules.

Kube-proxy will watch its own node and will find the endpoints that are in the same topology domain as the node if `service.Topology.Mode != ignored`.

The new logic of kube-proxy might like:

```go
go watch node with its nodename
switch service.Topology.Mode {
  case "ignored":
    route request to an endpoint randomly
  case "required":
    endpointsMeetRequirement := make([]endpointInfo, 0)
    topologyKey := service.Topology.Key
    // filter out endpoints that does not meet the "hard" topology requirements
    for i := range service's endpoints.Subsets; do
      ss := endpoints.Subsets[i]
      for j := range ss.Addresses; do
        // check if endpoint are in the same topology domain as the node running kube-proxy
        if ss.Addresses[j].Topology[topologyKey] == node.Labels[topologyKey]; then
          endpointsMeetHardRequirement = append(endpointsMeetHardRequirement, endpoint)
        fi
      done
    done
    // If multiple endpoints match, randomly select one
    if len(endpointsMeetHardRequirement) != 0; then 
      route request to an endpoint in the endpointsMeetHardRequirement randomly
    fi
  case "preferred":
    // Try to find endpoints that meet the "soft" topology requirements firstly,
    // If no one match, kube-proxy tell the kernel all available endpoints and ask it to to route each request randomly to one of them.
}
```
