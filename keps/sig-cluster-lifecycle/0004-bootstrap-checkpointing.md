---
kep-number: 4
title: Kubernetes Bootstrap Checkpointing Proposal
status: implemented
authors:
  - "@timothysc"
owning-sig: sig-cluster-lifecycle
participating-sigs:
  - sig-node
reviewers:
  - "@yujuhong"
  - "@luxas"
  - "@roberthbailey"
approvers:
  - "@yujuhong"
  - "@roberthbailey"
editor:
  name: @timothysc
creation-date: 2017-10-20
last-updated: 2018-01-23
---

# Kubernetes Bootstrap Checkpointing Proposal

## Table of Contents

* [Summary](#summary)
* [Objectives](#objectives)
  * [Goals](#goals)
  * [Non-Goals](#non-goals)
* [Proposal](#proposal)
  * [User Stories](#user-stories)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Unresolved Questions](#unresolved-questions)

## Summary

There are several methods to deploy a kubernetes cluster, one method that
offers some unique advantages is self hosting.  The purpose of this proposal
is to outline a method to checkpoint specific annotated pods, namely the
control plane components, for the purpose of enabling self hosting.

The details of self hosting are beyond the scope of this proposal, and are
outlined in the references listed below:

  - [Self Hosted Kubernetes][0]
  - [Kubeadm Upgrades][1]

Extra details on this proposal, and its history, can be found in the links
below:

  - [Bootstrap Checkpointing Draft 1][2]
  - [Bootstrap Checkpointing Draft 2][3]
  - [WIP Implementation][4]

## Objectives

The scope of this proposal is **bounded**, but has the potential for broader
reuse in the future.  The reader should be mindful of the explicitly stated
[Non-Goals](#non-goals) that are listed below.

### Goals

 - Provide a basic framework for recording annotated *Pods* to the filesystem.
 - Ensure that a restart of the kubelet checks for existence of these files
 and loads them on startup.

### Non-Goals

- This is not a generic checkpointing mechanism for arbitrary resources.
(e.g. Secrets)  Such changes require wider discussions.
- This will not checkpoint internal kubelet state.
- This proposal does not cover self hosted kubelet(s).  It is beyond the
scope of this proposal, and comes with it's own unique set of challenges.

## Proposal
The enablement of this feature is gated by a single command line flag that
is passed to the kubelet on startup, ```--bootstrap-checkpoint-path``` ,
and will be denoted that it is ```[Alpha]```.

### User Stories

#### Pod Submission to Running
- On submission of a Pod, via kubeadm or an operator, an annotation
```node.kubernetes.io/bootstrap-checkpoint=true``` is added to that Pod, which
indicates that it should be checkpointed by the kubelet.  When the kubelet
receives a notification from the apiserver that a new pod is to run, it will
inspect the ```--bootstrap-checkpoint-path``` flag to determine if
checkpointing is enabled.  Finally, the kubelet will perform an atomic
write of a ```Pod_UID.yaml``` file when the afore mentioned annotation exists.
The scope of this annotation is bounded and will not be promoted to a field.

#### Pod Deletion
- On detected deletion of a Pod, the kubelet will remove the associated
checkpoint from the filesystem.  Any failure to remove a pod, or file, will
result in an error notification in the kubelet logs.

#### Cold Start
- On a cold start, the kubelet will check the value of
```--bootstrap-checkpoint-path```.  If the value is specified, it will read in
the contents of the that directory and startup the appropriate Pod.  Lastly,
the kubelet will then pull the list of pods from the api-server and rectify
what is supposed to be running according to what is bound, and will go through
its normal startup procedure.

### Implementation Constraints
Due to its opt-in behavior, administrators will need to take the same precautions
necessary in segregating master nodes, when enabling the bootstrap annotation.

Please see [WIP Implementation][4] for more details.

## Graduation Criteria

Graduating this feature is a responsibility of sig-cluster-lifecycle and
sig-node to determine over the course of the 1.10 and 1.11 releases.  History
has taught us that initial implementations often have a tendency overlook use
cases and require refinement.  It is the goal of this proposal to have an
initial alpha implementation of bootstrap checkpoining in the 1.9 cycle,
and further refinement will occur after we have validated it across several
deployments.

## Testing
Testing of this feature will occur in three parts.
- Unit testing of standard code behavior
- Simple node-e2e test to ensure restart recovery
- (TODO) E2E test w/kubeadm self hosted master restart recovery of an apiserver.

## Implementation History

- 20171020 - 1.9 draft proposal
- 20171101 - 1.9 accepted proposal
- 20171114 - 1.9 alpha implementation code complete

## Unresolved Questions

* None at this time.

[0]: /contributors/design-proposals/cluster-lifecycle/self-hosted-kubernetes.md
[1]: https://github.com/kubernetes/community/pull/825
[2]: https://docs.google.com/document/d/1hhrCa_nv0Sg4O_zJYOnelE8a5ClieyewEsQM6c7-5-o/edit?ts=5988fba8#
[3]: https://docs.google.com/document/d/1qmK0Iq4fqxnd8COBFZHpip27fT-qSPkOgy1x2QqjYaQ/edit?ts=599b797c#
[4]: https://github.com/kubernetes/kubernetes/pull/50984
