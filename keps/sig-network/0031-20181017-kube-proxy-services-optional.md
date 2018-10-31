---
kep-number: 0
title: Make kube-proxy service abstraction optional
authors:
  - "@bradhoekstra"
owning-sig: sig-network
participating-sigs:
reviewers:
  - TBD
approvers:
  - TBD
editor: "@bradhoekstra"
creation-date: 2018-10-17
last-updated: 2018-10-17
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
* The apiserver sending all services and endpoints to all kube-proxy pods
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

It is important for overall scalability that kube-proxy does not receive data for Service/Endpoints objects that it is not going to affect. This can reduce load on the apiserver, networking, and kube-proxy itself by never receiving the updates in the first place.

The proposal is to make this feature available by annotating the Service object with this label: `kube-proxy.kubernetes.io/disabled=true`. The associated Endpoints object will automatically inherit that label from the Service object as well.

When this label is set, kube-proxy will behave as if that service does not exist. None of the functionality that kube-proxy provides will be available for that service.

It is expected that this feature will mainly be used on large clusters with lots (>1000) of services. Any use of this feature in a smaller cluster will have negligible impact.

The envisioned cluster that will make use of this feature looks something like the following:
* Most/all traffic from outside the cluster is handled by gateways, such that each service in the cluster does not need a nodePort
* These small number of entry points into the cluster are a part of the service mesh
* There are many micro-services in the cluster, all a part of the service mesh, that are only accessed from inside the service mesh

#### Design

Currently, when ProxyServer starts up it creates informers for all Service (ServiceConfig) and Endpoints (EndpointsConfig) objects using a single shared informer factory.

The new design will simply add a LabelSelector filter to the shared informer factory, such that objects with the above label are filtered out by the API server:
```diff
-       informerFactory := informers.NewSharedInformerFactory(s.Client, s.ConfigSyncPeriod)
+       informerFactory := informers.NewSharedInformerFactoryWithOptions(s.Client, s.ConfigSyncPeriod,
+               informers.WithTweakListOptions(func(options *v1meta.ListOptions) {
+                       options.LabelSelector = "kube-proxy.kubernetes.io/disabled!=true"
+               }))
```

#### Testing

The following cases should be tested. In each case, make sure that services are added/removed from iptables (or other) as expected:
* Adding/removing services/endpoints with and without the above label
* Adding/removing the above label from existing services/endpoints
* Having a label value other than 'true', which should behave as if the label is not set

### Risks and Mitigations

We will keep the existing behaviour enabled by default, and only disable it when the cluster operator specifically asks to do so.

## Graduation Criteria

N/A

## Implementation History

- 2018-10-17 - This KEP is created
- 2018-10-28 - KEP updated
