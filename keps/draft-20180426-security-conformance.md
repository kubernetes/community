---
kep-number: draft-20180426
title: Security Conformance
authors:
  - "@easeway"
owning-sig: sig-auth
participating-sigs:
  - sig-testing
  - sig-api-machinery
  - sig-network
  - sig-node
  - sig-cluster-lifecycle
reviewers:
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
  - "@davidopp"
approvers:
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
editor: "@easeway"
creation-date: 2018-04-26
status: provisional
---

# Security Conformance

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [List of Criteria](#list-of-criteria)
* [Graduation Criteria](#graduation-criteria)

## Summary

Security Conformance defines citeria to certify a Kubernetes cluster to be minimally secure.

## Motivation

Configuring a Kubernetes cluster is complicated,
and it's very easy to mis-configure a cluster to be insecure.
The current Kubernetes conformance test suite barely covers security aspects.

### Goals

- Define the initial list (to evolve later) of criteria for minimal security requirements
  - Define the role of the actor (unprivileged user, namespace admin, cluster admin, etc.)
  - Define pre-requisites to run test for verifying the criteria
  - Define test behaviors
  - Define test expectations

### Non-Goals

- A full list of criteria at the beginning
  - The list of criteria is to be curated by the community
  - This proposal starts with minimal criteria to be evolved quickly
- Mechanisms to implement the test
- Mechanisms to setup environment for the tests
- Host setup and control plane configurations (covered by CIS benchmark)
- Non-security related

## Proposal

### List of Criteria

#### API privileges of default service accounts

##### Rationale

Every namespace has a service account named `default`.
This is used by Pods created without specifying a particular service account.
When the process running in the Pod invokes Kubernetes API, 
this `default` service account is used as the identity for authn/authz.

##### Pre-requisites

- A Kubernetes cluster should be up and running
- A `kubeconfig` to allow the test to
  - Create a Pod in a namespace
  - It doesn't matter what the additional permissions the user have
- A namespace created for test
  - This can be `default` namespace, but not `kube-system`
  - The `default` service account in the namespace should not be granted extra permissions
- The DNS name of API server inside the cluster, if it's not default

##### Test Behavior

- Create a Pod in the test namespace without specifying a service account (using `default`)
- The Pod contains test logic which invokes Kubernetes API to
  - `get/list/watch` namespaces
  - `create` another Pod in the namespace
  - `create` a new namespace
  - `update/patch` labels/annotations of current Pod
  - `delete` current Pod

##### Test Expectation

- All attempts using Kubernetes API should fail (Forbidden).

#### Kubelet r/w API Access

##### Rationale

Kubelet read/write API (listening on port 10250) is designed to be consumed by API server.
It's a security concern if other components, 
including the workload running on the cluster is able to access the API.

##### Pre-requisites

- A Kubernetes cluster should be up and running
- A `kubeconfig` to allow the test to
  - Create a Pod in a namespace
- A namespace for test Pod
- The DNS name of API server inside the cluster, if it's not default
- The DNS name/IP of Kubelet to be tested

##### Test Behavior

- Create a Pod in the test namespace running the test code
- The test code in the Pod attempts to access R/W API on Kubelet directly
  - `GET /exec/{namespace}/{currentPodID}/{currentContainerName}`

##### Test Expectation

- The API invocation should fail.

## Graduation Criteria

- The list of criteria is agreed by community
- The tests are integrated into Kubernetes conformance test suite
