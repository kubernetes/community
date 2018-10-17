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
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)

## Summary

In a cluster that has a service mesh a lot of the work being done by kube-proxy is redundant and wasted.
Specifically, services that are only reached via other services in the mesh will never use the service abstaction implemented by kube-proxy in iptables (or ipvs).
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

## Proposal

### User Stories

#### Story 1

As a cluster operator, operating a cluster using a service mesh I want to be able to disable the kube-proxy service implementation for services in that mesh to reduce overall load on the whole cluster

### Implementation Details/Notes/Constraints

It is important for overall scalability that kube-proxy does not watch for service/endpoint changes that it is not going to affect. This can save a lot of load on the apiserver, networking, and kube-proxy itself by never requesting the updates in the first place. As such, annotating the services directly is considered insufficient as the kube-proxy would still have to watch for changed to the service.

The proposal is to make this feature available at the namespace level:

We will support a new label for namespaces: networking.k8s.io/kube-proxy=disabled

kube-proxy will be modified to watch all namespaces and stop watching for services/endpoints in namespaces with the above label.

The following cases should be tested. In each case, make sure that services are added/removed from iptables (or other) as expected:
* Adding/removing services from namespaces with and without the above label
* Adding/removing the above label from namespaces with existing services

### Risks and Mitigations

We will keep kube-proxy enabled by default, and only disable it when the cluster operator specifically asks to do so.

## Graduation Criteria

N/A

## Implementation History

- 2018-10-17 - This KEP is created
