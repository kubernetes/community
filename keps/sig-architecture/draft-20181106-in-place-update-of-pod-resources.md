---
kep-number: draft-20181106
title: In-place Update of Pod Resources
authors:
  - "@kgolab"
  - "@bskiba"
  - "@schylek"
owning-sig: sig-architecture
participating-sigs:
  - sig-autoscaling
  - sig-node
  - sig-scheduling
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-11-06
last-updated: 2018-11-06
status: provisional
see-also:
replaces:
superseded-by:
---

# In-place Update of Pod Resources

## Table of Contents

   * [In-place Update of Pod Resources](#in-place-update-of-pod-resources)
      * [Table of Contents](#table-of-contents)
      * [Summary](#summary)
      * [Motivation](#motivation)
         * [Goals](#goals)
         * [Non-Goals](#non-goals)
      * [Proposal](#proposal)
         * [API Changes](#api-changes)
         * [Flow Control](#flow-control)
         * [Notes](#notes)
         * [Risks and Mitigations](#risks-and-mitigations)
      * [Graduation Criteria](#graduation-criteria)
      * [Implementation History](#implementation-history)
      * [Alternatives](#alternatives)

## Summary

This proposal aims at allowing Pod resource requests & limits to be updated
in-place, without a need to restart the Pod or its Containers.

The **core idea** behind the proposal is to make PodSpec mutable with regards to
Resources, denoting **desired** resources.
Additionally PodStatus is extended to provide information about **actual**
resource allocation.

This document builds upon [proposal for live and in-place vertical scaling][] and
[Vertical Resources Scaling in Kubernetes][].

[proposal for live and in-place vertical scaling]: https://github.com/kubernetes/community/pull/1719
[Vertical Resources Scaling in Kubernetes]: https://docs.google.com/document/d/18K-bl1EVsmJ04xeRq9o_vfY2GDgek6B6wmLjXw-kos4/edit?ts=5b96bf40

## Motivation

Resources allocated to a Pod's Container can require a change for various reasons:
* load handled by the Pod has increased significantly and current resources are
  not enough to handle it,
* load has decreased significantly and currently allocated resources are unused
  and thus wasted,
* Resources have simply been set improperly.

Currently changing Resources allocation requires the Pod to be recreated since
the PodSpec is immutable.

While many stateless workloads are designed to withstand such a disruption, some
are more sensitive, especially when using low number of Pod replicas.

Moreover, for stateful or batch workloads, a Pod restart is a serious
disruption, resulting in lower availability or higher cost of running.

Allowing Resources to be changed without recreating a Pod nor restarting a
Container addresses this issue directly.

### Goals

* Primary: allow to change Pod resource requests & limits without restarting its
  Containers.
* Secondary: allow actors (users, VPA, StatefulSet, JobController) to decide
  how to proceed if in-place resource update is not available.
* Secondary: allow users to specify which Pods and Containers can be updated
  without a restart.

### Non-Goals

The explicit non-goal of this KEP is to avoid controlling full life-cycle of a
Pod which failed an in-place resource update. These cases should be handled by
actors which initiated the update.

Other identified non-goals are:
* allow to change Pod QoS class without a restart,
* to change resources of Init Containers without a restart.

## Proposal

### API Changes

PodSpec becomes mutable with regards to resources and limits.
Additionally, PodSpec becomes a Pod subresource to allow fine-grained access control.

PodStatus is extended with information about actually allocated resources.

Thanks to the above:
* PodSpec.Container.ResourceRequirements becomes purely a declaration,
  denoting **desired** state of the Pod,
* PodStatus.ContainerStatus.ResourceAllocated (new object) denotes **actual**
  state of the Pod resources.

To distinguish between possible states of the Pod resources,
a new PodCondition InPlaceResize is added, with the following states:
* (empty) - the default value; resource update awaits reconciliation
  (if ResourceRequirements differs from ResourceAllocated),
* Awaiting - awaiting resources to be freed (e.g. via pre-emption),
* Failed - resource update could not have been performed in-place
  but might be possible if some conditions change,
* Rejected - resource update was rejected by any of the components involved.

To provide some fine-grained control to the user,
PodSpec.Container.ResourceRequirements is extended with ResizingPolicy flag,
available per each resource request (CPU, memory) :
* InPlace - the default value; allow in-place resize of the Container,
* RestartContainer - restart the Container to apply new resource values
  (e.g. Java process needs to change its Xmx flag),
* RestartPod - restart whole Pod to apply new resource values
  (e.g. Pod requires its Init Containers to re-run).

By using the ResizingPolicy flag the user can mark Containers or Pods as safe
(or unsafe) for in-place resources update.

This flag **may** be used by the actors starting the process to decide if
the process should be started at all (for example VPA might decide to
evict Pod with RestartPod policy).
This flag **must** be used by Kubelet to verify the actions needed.

Setting the flag to separately control CPU & memory is due to an observation
that usually CPU can be added/removed without much problems whereas
changes to available memory are more probable to require restarts.

### Flow Control

TODO

### Notes

TODO

### Risks and Mitigations

TODO

## Graduation Criteria

TODO

## Implementation History

- 2018-11-06 - initial KEP draft created

## Alternatives

TODO

