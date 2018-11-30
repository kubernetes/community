# KEP: Portable Service Definitions

---
kep-number: 31
title: Portable Service Definitions
authors:
  - "@mattfarina"
owning-sig: sig-apps
participating-sigs:
  - sig-service-catalog
reviewers:
  - "@carolynvs"
  - "@kibbles-n-bytes"
  - "@duglin"
  - "@jboyd01"
  - "@prydonius"
  - "@kow3ns"
approvers:
  - "@mattfarina"
  - "@prydonius"
  - "@kow3ns"
editor: TBD
creation-date: 2018-11-13
last-updated: 2018-11-19
status: provisional
see-also:
replaces:
superseded-by:

---

# Portable Service Definitions

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

## Summary

The goal of this feature is to enable an application to be deployed into multiple environments while relying on external services that are not part of the application and using the same objects in all environments. That includes service instances that may or may not be managed by Kubernetes. For example, take a WordPress application that relies on MySQL. If the application is running in GKE or AKS it may leverage a Google or Azure MySQL as a Service. If it is running in Minikube or in a bare metal cluster it may use MySQL managed by the Kubernetes cluster. In all of these cases, the same resource is declared asking for the service (e.g., MySQL) and the credentials to use the service are returned in a common manner (e.g., a secret with the same type and schema). This enables portability from one environment to another because working with services in all environments leverages the same Kubernetes objects with the same schemas.

## Motivation

Workload portability is commonly cited as a goal for those deploying workloads. The Kubernetes API can provide a common API for deploying workloads across varying environments enabling some level of portability. For example, a deployment can be run in a cluster running on GKE, AKS, EKS, or clusters running elsewhere.

But, many applications rely on software as a service (SaaS). The reason for this is to push the operational details on to someone else who specializes in that particular service so the application developers and operators can focus on their application and business logic.

The problem is that one cannot deploy the same application in two different environments by two different providers, if the applications leverages services, with the same set of resources. This includes cases where the service being leveraged is common (e.g., MySQL as a Service). This problem limits application portability and sharing (e.g., in open source).

This KEP is looking to solve this problem by providing Kubernetes compatible objects, via CRDs and Secrets, that can be used in many environments by many providers to make working with common services easier. This can be used for services like database (e.g., MySQL, PostgreSQL), DNS, SMTP, and many others.

### Goals

* Provide a common way to request common services (e.g., MySQL)
* Provide a common means to obtain credentials to use the service
* Provide a common method to detect which services are available
* Provide a system that can be implemented for the major public clouds, on-premise clusters, and local cluster (e.g., Docker for Mac)

### Non-Goals

* Provide an out of the box solution for every bespoke service provided by everyone
* Replace Kubernetes Service Catalog

## Proposal

### User Stories

#### Story 1

As a user of Kubernetes, I can query the services I can declaratively request using a Kubernetes native API. For example, using the command `kubectl get crds` where I can see services alongside other resources that can be created.

#### Story 2

As a user of Kubernetes, I can declaratively request an instance of a service using a custom resource. When the service is provided the means to use that service (e.g., credentials in a secret) are provided in a common and consistent manner. The same resource and secret can be used in clusters running in different locations and the way the service is provided may be different.

#### Story 3

As a cluster operator or application operator, I can discover controllers implementing the CRDs and secrets to support the application portability in my cluster.

#### Story 4

As a cluster operator or application operator, I can set default values and provider custom settings for a service. 

### Implementation Details/Notes/Constraints

To solve the two user stories there are two types of Kubernetes resources that can be leveraged.

1. Custom resource definitions (CRDs) can be used to describe a service. The CRDs can be implemented by controllers for different environments and the list of installed CRDs can be queried to see what is supported in a cluster
2. Secrets with a specific type and schema can be used to handle credentials and other relevant information for services that have them (e.g., a database). Not all services will require a secret (e.g., DNS)

This subproject will list and document the resources and how controllers can implement them. This provides for interoperability including that for controllers and other tools, like validators, and a canonical listing held by a vendor neutral party.

In addition to the resources, this subproject will also provide a controller implementing the defined services to support testing, providing an example implementation, and to support other Kubernetes subprojects (e.g., Minikube). Controllers produced by this project are _not_ meant to be used in production.

3rd party controllers implementing the CRDs and secrets can use a variety of methods to implement the service handling. This is where the Kubernetes Service Catalog can be an option. This subproject will not host or support 3rd party controllers but will list them to aide in users discovering them. This is in support of the 3rd user story.

To support custom settings for services by a service provider and to add the ability to add default settings (user story 4) we are considering a pattern of using CRDs for a controller with configuration on a cluster wide and namespace based level. An example of this in existence today is the cert manager issuer and cluster issuer resources. How to support this pattern will be worked out as part of the process in the next step of building a working system. This pattern is tentative until worked out in practice.

The next step is to work out the details on a common service and an initial process by which future services can be added. To work out the details we will start with MySQL and go through the process to make it work as managed by Kubernetes and each of the top three public clouds as defined by adoption. Public clouds and other cloud platforms following the top three are welcome to be involved in the process but are not required in the process.

Before this KEP can move to being implementable at least 2 services need to go through the process of being implemented to prove out the process elements of this system.

### Risks and Mitigations

Two known risks include:

1. The access controls and relationships between accounts and services. How will proper user and tenancy information be passed to clouds that require this form of information?
2. Details required, when requesting a service, can vary between cloud providers that will implement this as a SaaS. How can that information be commonized or otherwise handled?

## Graduation Criteria

The following are the graduation criteria:

- 5 organizations have adopted using the portable service definitions
- Service definitions for at least 3 services have been created and are in use
- Documentation exists explaining how a controller implementer can use the CRDs and secrets to create a controller of their own
- Service consumer documentation exists explaining how to use portable service definitions within their applications
- A documented process for bringing a new service from suggestion to an implementable solution

## Implementation History

- the _Summary_, _Motivation_, _Proposal_ sections developed

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives [optional]

An alternative is to modify the service catalog to leverage CRDs and return common secret credentials. Drawbacks to this are that it would be a complicated re-write to the service catalog, according to the service catalog team, and the solution would still require the open service broker (OSB) to be implemented in all environments (e.g., Minikube) even where simpler models (e.g., a controller) could be used instead. The solution proposed here could work with the service catalog in environments it makes sense and use other models in other environments. The focus here is more on the application operator experience working with services than all of the implementations required to power it.

## Infrastructure Needed [optional]

The following infrastructure elements are needed:

- A new subproject under SIG Apps for organizational purposes
- A git repository, in the `kubernetes-sigs` organization, to host the CRD and Secrets schemas along with the Kubernetes provided controller
- Testing infrastructure to continuously test the codebase
