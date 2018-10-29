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
      * [Considerations](#considerations)
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

It is important for overall scalability that kube-proxy does not watch for service/endpoint changes that it is not going to affect. This can save a lot of load on the apiserver, networking, and kube-proxy itself by never requesting the updates in the first place. As such, annotating the services directly is considered insufficient as the kube-proxy would still have to watch for changed to the service.

The proposal is to make this feature available at the namespace level. We will support a new label for namespaces: `networking.k8s.io/service-proxy=disabled`

When this label is set, kube-proxy will behave as if services in that namespace do not exist. None of the functionality that kube-proxy provides will be available for services in that namespace.

It is expected that this feature will mainly be used on large clusters with lots (>1000) of services. Any use of this feature in a smaller cluster will have negligible impact.

The envisioned cluster that will make use of this feature looks something like the following:
* Most/all traffic from outside the cluster is handled by gateways, such that each service in the cluster does not need a nodePort
* These small number of entry points into the cluster are a part of the service mesh
* There are many micro-services in the cluster, all a part of the service mesh, that are only accessed from inside the service mesh
  * These services are in a separate namespace from the gateways

#### Design

Currently, when ProxyServer starts up it creates informers for all Service (ServiceConfig) and Endpoints (EndpointsConfig) objects using a single shared informer factory. The new design will make these previous objects be per-namespace, and only listen on namespaces that are not 'disabled'.

The ProxyServer type will be updated with the following new methods:
* func (s *ProxyServer) StartWatchingNamespace(ns string)
  * Check if namespace is currently watched, if it is then return
  * Create a shared informer factory configured with the namespace
  * Create a ServiceConfig and EndpointsConfig object using the shared informer factory
* func (s *ProxyServer) StopWatchingNamespace(ns string)
  * Check if namespace is currently watched, if it is not then return
  * Stop the ServiceConfig and EndpointsConfig for that namespace
  * Send deletion events for all objects those configs knew about
  * Delete the config objects

At startup time, ProxyServer will create an informer for all Namespace objects.
* When a namespace objects is created or updated:
  * Check for the above label, and if it is not set or is not 'disabled':
    * StartWatchingNamespace()
  * Else:
    * StopWatchingNamespace()
* When a namespace object is deleted:
  * StopWatchingNamespace()

#### Considerations

kube-proxy has logic in it right now to not sync rules until the config objects have been synced. Care should be taken to make sure this logic still works, and that the data is only considered synced when the Namespace informer and all ServiceConfig and EndpointsConfig objects are synced.

#### Testing

The following cases should be tested. In each case, make sure that services are added/removed from iptables (or other) as expected:
* Adding/removing services from namespaces with and without the above label
* Adding/removing the above label from namespaces with existing services
* Deleting a namespace with services with and without the above label
* Having a label value other than 'disabled', which should behave as if the label is not set

### Risks and Mitigations

We will keep the existing behaviour enabled by default, and only disable it when the cluster operator specifically asks to do so.

## Graduation Criteria

N/A

## Implementation History

- 2018-10-17 - This KEP is created
- 2018-10-28 - KEP updated
