---
kep-number: 31
title: Make kube-proxy service abstraction optional
authors:
  - "@bradhoekstra"
owning-sig: sig-network
participating-sigs:
reviewers:
  - "@freehan"
approvers:
  - "@thockin"
editor: "@bradhoekstra"
creation-date: 2018-10-17
last-updated: 2018-11-12
status: provisional
see-also:
replaces:
superseded-by:
---

# Make kube-proxy service abstraction optional

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
      * [Story 1](#story-1)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
      * [Design](#design)
      * [Testing](#testing)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)

## Summary

In a cluster that has a service mesh a lot of the work being done by kube-proxy is redundant and wasted.
Specifically, services that are only reached via other services in the mesh will never use the service abstraction implemented by kube-proxy in iptables (or ipvs).
By informing the kube-proxy of this, we can lighten the work it is doing and the burden on its proxy backend.

## Motivation

The motivation for the enhancement is to allow higher scalability in large clusters with lots of services that are making use of a service mesh.

### Goals

The goal is to reduce the load on:
* The kube-proxy having to deserialize and process all services and endpoints
* The backend system (e.g. iptables) for whichever proxy mode kube-proxy is using

### Non-Goals

* Making sure the service is still routable via the service mesh
* Preserving any kube-proxy functionality for any intentionally disabled Service, including but not limited to: externalIPs, external LB routing, nodePorts, externalTrafficPolicy, healthCheckNodePort, UDP, SCTP

## Proposal

### User Stories

#### Story 1

As a cluster operator, operating a cluster using a service mesh I want to be able to disable the kube-proxy service implementation for services in that mesh to reduce overall load on the whole cluster

### Implementation Details/Notes/Constraints

#### Overview

In a cluster where a service is only accessed via other applications in the service mesh the work that kube-proxy does to program the proxy (e.g. iptables) for that service is duplicated and unused. The service mesh itself handles load balancing for the service VIP. This case is often true in the standard service mesh setup of utilizing ingress/egress gateways, such that services are not directly exposed outside the cluster. In this setup, application services rarely make use of other Service features such as externalIPs, external LB routing, nodePorts, externalTrafficPolicy, healthCheckNodePort, UDP, SCTP. We can optimize this cluster by giving kube-proxy a way to not have to perform the duplicate work for these services.

It is important for overall scalability that kube-proxy does not receive data for Service/Endpoints objects that it is not going to affect. This can reduce load on the kube-proxy and the network by never receiving the updates in the first place.

The proposal is to make this feature available by annotating the Service object with this label: `service.kubernetes.io/service-proxy-name`. If this label key is set, with any value, the associated Endpoints object will automatically inherit that label from the Service object as well.

When this label is set, kube-proxy will behave as if that service does not exist. None of the functionality that kube-proxy provides will be available for that service.

kube-proxy will properly implement this label both at object creation and on dynamic addition/removal/updates of this label, either providing functionality or not for the service based on the latest version on the object.

It is optional for other service proxy implementations (besides kube-proxy) to implement this feature. They may ignore this value and still remain conformant with kubernetes services.

It is expected that this feature will mainly be used on large clusters with lots (>1000) of services. Any use of this feature in a smaller cluster will have negligible impact.

The envisioned cluster that will make use of this feature looks something like the following:
* Most/all traffic from outside the cluster is handled by gateways, such that each service in the cluster does not need a nodePort
* These small number of entry points into the cluster are a part of the service mesh
* There are many micro-services in the cluster, all a part of the service mesh, that are only accessed from inside the service mesh

Higher level frameworks built on top of service meshes, such as [Knative](https://github.com/knative/docs), will be able to enable this feature by default due to having a more controlled application/service model and being reliant on the service mesh.

#### Design

Currently, when ProxyServer starts up it creates informers for all Service (ServiceConfig) and Endpoints (EndpointsConfig) objects using a single shared informer factory.

The new design will simply add a LabelSelector filter to the shared informer factory, such that objects with the above label are filtered out by the API server:
```diff
-       informerFactory := informers.NewSharedInformerFactory(s.Client, s.ConfigSyncPeriod)
+       informerFactory := informers.NewSharedInformerFactoryWithOptions(s.Client, s.ConfigSyncPeriod,
+               informers.WithTweakListOptions(func(options *v1meta.ListOptions) {
+                       options.LabelSelector = "!service.kubernetes.io/service-proxy-name"
+               }))
```

This code will also handle the dynamic label update case. When the label selector is matched (service is enabled) an 'add' event will be generated by the informer. When the label selector is not matched (service is disabled) a 'delete' event will be generated by the informer.

#### Testing

The following cases should be tested. In each case, make sure that services are added/removed from iptables (or other) as expected:
* Adding/removing services/endpoints with and without the above label
* Adding/removing the above label from existing services/endpoints

### Risks and Mitigations

We will keep the existing behaviour enabled by default, and only disable the kube-proxy service proxy when the service contains this new label.

This will have no effect on alternate service proxy implementations since they will not handle this label.

## Graduation Criteria

N/A

## Implementation History

- 2018-10-17 - This KEP is created
- 2018-11-12 - KEP updated, including approver/reviewer
