---
kep-number: 0018
title: Make CPU Manager respect "isolcpus"
authors:
  - "@Levovar"
owning-sig: sig-node
participating-sigs:
  - sig-node
reviewers:
  - "@jeremyeder"
  - "@ConnorDoyle"
  - "@bgrant0607"
  - "@dchen1107" 
approvers:
  - TBD
editor: TBD
creation-date: 2018-07-30
last-updated: 2018-07-31
status: provisional
see-also:
  - N/A
  - N/A
replaces:
  - N/A
superseded-by:
  - N/A
---

# Make CPU Manager respect "isolcpus"

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

"Isolcpus" is a boot-time Linux kernel parameter, which can be used to isolate CPU cores from the generic Linux scheduler.
This kernel setting is routinely used within the Linux community to manually isolate, and then assign CPUs to specialized workloads.
The CPU Manager implemented within kubelet currently ignores this kernel setting when creating cpusets for Pods.
This KEP proposes that CPU Manager should respects this kernel setting when assigning Pods to cpusets, through whichever supported CPU management policy.

## Motivation

The CPU Manager always assumes that it is the alpha and omega on a node, when it comes to managing the CPU resources of the host.
However, in certain infrastructures this might not always be the case.
While it is already possible to effectively take-away CPU cores from the CPU manager via the kube-reserved and system-reserved kubelet flags, this implicit way of expressing isolation needs is not dynamic enough to cover all use-cases.

Therefore, the need arises to enhance existing CPU manager with a method of explicitly defining a discontinuous pool of CPUs it can manage.
Making kubelet respect the isolcpus kernel setting fulfills exactly that need, while also doing it in a de-facto standard way.
 
If Kubernetes' CPU manager would support this more granular node configuration, it would enable infrastructure administrators to make multiple "CPU managers" seamlessly inter-work on the same node.
For example:
- outsourcing the management of a subset of specialized, or optimized CPUs to an external CPU manager without any (other) change in Kubelet's CPU manager
- ensure proper resource accounting and separation within a hybrid infrastructure (e.g. Openstack + Kubernetes running on the same node)

### Goals

The goal is to make any and all Kubernetes supported CPU management policies restrictable to a subset of a nodes' capacity.
The goal is to make Kubernetes respect an already existing node-level configuration option, which already means exactly that in the Linux community.

### Non-Goals

It is outside the scope of this KEP to restrict any other Kubernetes resource manager to a subset of a resource group (like memory, devices, etc.).
It is also outside the scope of this KEP to enhance the CPU manager itself with more fine-grained management policies, or introduce topology awareness into the CPU manager.
The aim of this KEP is to continue to let Kubernetes manage some CPU cores however it sees fit, but also let room for "other" managers running on the same host.

## Proposal

This is where we get down to the nitty gritty of what the proposal actually is.

### User Stories [optional]

Detail the things that people will be able to do if this KEP is implemented.
Include as much detail as possible so that people can understand the "how" of the system.
The goal here is to make this feel real for users without getting bogged down.

#### Story 1

#### Story 2

### Implementation Details/Notes/Constraints [optional]

What are the caveats to the implementation?
What are some important details that didn't come across above.
Go in to as much detail as necessary here.
This might be a good place to talk about core concepts and how they releate.

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives [optional]

Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other possible approaches to delivering the value proposed by a KEP.
