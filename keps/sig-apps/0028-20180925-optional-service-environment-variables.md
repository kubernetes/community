---
kep-number: 28
title: Optional Service Environment Variables
authors:
  - "@bradhoekstra"
  - "@kongslund"
owning-sig: sig-apps
participating-sigs:
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-09-25
last-updated: 2018-09-25
status: provisional
see-also:
  - https://github.com/kubernetes/community/pull/1249
replaces:
superseded-by:
---

# Optional Service Environment Variables

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

This enhancement allows application developers to choose whether their Pods will receive [environment variables](https://kubernetes.io/docs/concepts/services-networking/service/#environment-variables) from services in their namespace. They can choose to disable them via the new `enableServiceLinks` field in `PodSpec`. The current behaviour will continue to be the default behaviour, but the developer may choose to disable these environment variables for certain workloads for reasons such as incompatibilities with other expected environment variables or scalability issues.

## Motivation

Today, a list of all services that were running when a pod's containers are created is automatically injected to those containers as environment variables matching the syntax of Docker links. There is no way to disable this.

Docker links have long been considered as a [deprecated legacy feature](https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/) of Docker since the introduction of networks and DNS. Likewise, in Kubernetes, DNS is to be preferred over service links.

Possible issues with injected service links are:

* Accidental coupling.
* Incompatibilities with container images that no longer utilize service links and explicitly fail at startup time if certain service links are defined.
* Performance penalty in starting up pods [for namespaces with many services](https://github.com/kubernetes/kubernetes/issues/1768#issuecomment-330778184)

### Goals

* Allow users to choose whether to inject service environment variables in their Pods.
* Do this in a backwards-compatible, non-breaking way. Default to the current behaviour.

### Non-Goals

N/A

## Proposal

### User Stories

* As an application developer, I want to be able to disable service link injection since the injected environment variables interfere with a Docker image that I am trying to run on Kubernetes.
* As an application developer, I want to be able to disable service link injection since I don't need it and it takes increasingly longer time to start pods as services are added to the namespace.
* As an application developer, I want to be able to disable service link injection since pods can fail to start if the environment variable list becomes too long. This can happen when there are >5,000 services in the same namespace.

### Implementation Details/Notes/Constraints

`PodSpec` is extended with an additional field, `enableServiceLinks`. The field should be a pointer to a boolean and default to true if nil.

In `kubelet_pods.go`, the value of that field is passed along to the function `getServiceEnvVarMap` where it is used to decide which services will be propogated into environment variables. In case `enableServiceLinks` is false then only the `kubernetes` service in the `kl.masterServiceNamespace` should be injected. The latter is needed in order to preserve Kubernetes variables such as `KUBERNETES_SERVICE_HOST` since a lot of code depends on it.

### Risks and Mitigations

The current behaviour is being kept as the default as much existing code and documentation depends on these environment variables.

## Graduation Criteria

N/A

## Implementation History

- 2017-10-21: First draft of original proposal [PR](https://github.com/kubernetes/community/pull/1249)
- 2018-02-22: First draft of implementation [PR](https://github.com/kubernetes/kubernetes/pull/60206)
- 2018-08-31: General consensus of implementation plan
- 2018-09-17: First draft of new implementation [PR](https://github.com/kubernetes/kubernetes/pull/68754)
- 2018-09-24: Implementation merged into master
- 2018-09-25: Converting proposal into this KEP
