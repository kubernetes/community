---
kep-number: 0024
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
last-updated: 2018-08-07
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
* [Alternatives [optional]](#alternatives-optional)

## Summary

"Isolcpus" is a boot-time Linux kernel parameter, which can be used to isolate CPU cores from the generic Linux scheduler.
This kernel setting is routinely used within the Linux community to manually isolate, and then assign CPUs to specialized workloads.
The CPU Manager implemented within kubelet currently ignores this kernel setting when creating cpusets for Pods.
This KEP proposes that CPU Manager should respect the aformentioned kernel setting when assigning Pods to cpusets. The manager should behave the same irrespective of its configured management policy.
Inter-working with the isolcpus kernel parameter should be a node-wide, policy-agnostic setting.

## Motivation

Kubelet's in-built CPU Manager always assumes that it is the primary software component managing the CPU cores of the host.
However, in certain infrastructures this might not always be the case.
While it is already possible to effectively take-away CPU cores from the Kubernetes managed workloads via the kube-reserved and system-reserved kubelet flags, this implicit way of declaring a Kubernetes managed CPU pool is not flexible enough to cover all use-cases.

Therefore, the need arises to enhance existing CPU manager with a method of explicitly defining a discontinuous pool of CPUs it can manage.
Making kubelet respect the isolcpus kernel setting fulfills exactly that need, while also doing it in a de-facto standard way.

If Kubernetes' CPU manager would support this more granular node configuration, then infrastructure administrators could make multiple "CPU managers" seamlessly inter-work on the same node.
Such feature could come in handy if one would like to:
- outsource the management of a subset of specialized, or optimized cores (e.g. real-time enabled CPUs, CPUs with different HT configuration etc.) to an external CPU manager without any (other) change in Kubelet's CPU manager
- ensure proper resource accounting and separation within a hybrid infrastructure (e.g. Openstack + Kubernetes running on the same node)

### Goals

The goal is to make any and all Kubernetes supported CPU management policies restrictable to a subset of a nodes' capacity.
The goal is to make Kubernetes respect an already existing node-level Linux kernel parameter, which carries this exact meaning within the Linux community.

### Non-Goals

It is outside the scope of this KEP to restrict any other Kubernetes resource manager to a subset of another resource group (like memory, devices, etc.).
It is also outside the scope of this KEP to enhance kubelet's CPU manager itself with more fine-grained management policies, or introduce topology awareness into the CPU manager as an additional policy.
The aim of this KEP is to continue to let Kubernetes manage some CPU cores however it sees fit, but at the same time also leave the supervision of truly isolated resources to "other" resource managers.
Lastly, while it would be an interesting research topic of how different CPU managers (one of them being kubelet) could inter-work with each other in run-time to dynamically re-partition the CPU sets they manage, it is unfortunately also outside the scope of this simple KEP.
What this enhancement is trying to achieve first and foremost is isolation. Alignment of the isolated resources is left to the cloud infrastructure operators at this stage of the feature.

## Proposal

### User Stories

#### User Story 1 - As an infrastructure operator, I would like to exclusively dedicate some discontinuously numbered CPU cores to services not (entirely) supervised by Kubernetes

As stated in the Motivation section, Kubernetes might not be the only CPU manager running on a node in certain infrastructures.
A very specific example is an infrastructure which hosts real-time, very performance sensitive applications such as e.g. mobile network radio equipments.

Even this specific example can be broken down to multiple sub user-stories:
- a whole workload, or just some very sensitive parts of it continue to run directly on bare metal, while the rest of its communication partners are managed by Kubernetes
- everything is ran by Kubernetes, but some Pods require the services of a specialized CPU manager for optimal performance

In both cases the end result is effectively the same: the infrastructure operator manually dedicates a subset of a host's CPU capacity to a specialized controller, betting on that the specialized controller can serve the exact needs of the operator better.
The only difference between the sub user-stories is whether the operator also needs to somehow make the specialized controller inter-work with Kubernetes (for example by making the separated, and probably optimized CPUs available for consumption as "Devices"), or just simply work in isolation from its CPU manager.

In any case, the CPU cores used by such specialized controllers are routinely isolated from the operating system via the isolcpus parameter. Besides isolating these cores, operators usually also:
- manually optimize these cores (e.g. HTing, real-time patches, removal of kernel threads etc.)
- align the NUMA socket ID of these cores to other devices consumed by the sensitive applications (e.g. network devices)

Considering the above, it would make sense to re-use the same parameter to isolate these resources from Kubernetes too. Later on, when the specialized external resource controller actually starts dealing out these CPUs to workloads, it is usually done via the same mechanisms also employed by kubelet: either via the creation of CPU sets, or by manually setting the CPU affinity of other processes.

#### User Story 2 - As an infrastructure operator, I would like to run multiple cloud infrastructures in the same edge cloud

This user-story is actually very similar to the previous one, but less abstract. Imagine that an operator would like to run Openstack, VMware or any other popular cloud infrastructures next to Kubernetes, but without the need to physically separate these infrastructure.

Sometimes an operator simply does not have the possibility to separate her infrastructures on the host level, because simply there are not enough nodes available on the site. Typical use-case is an edge cloud, where usually multiple, high-available, NAS-including cloud infrastructures need to be brought-up on only a handful of nodes (3-10).

But, it can also happen that an operator simply would not wish to dedicate very powerful -e.g. OCP standard- servers in her central data centre just to host an under-utilized, "minority" cloud installation next to her "major" one.

In both cases, the resource manager components of both infrastructures will inevitably contest for the same resources. It should be noted that all different infrastructures need to also dedicate some CPUs to their management components too, in order to guarantee certain SLAs.

The different managers of more mature cloud infrastructures -for example Openstack- can already be configured to manage only a subset of a nodes' resource; isolated from all other process via the isolcpus kernel parameter.
If Kubernetes would also support the same feature, operators would be able to 1: isolate the common compute CPU pool from the operation system, and 2: manually divide the pool between the infrastructures however they see fit. 

### Implementation Details/Notes/Constraints

The pure implementation of the feature described in this document would be a fairly simple one. Kubernetes already contains code to remove a couple of CPU cores from the domain of its CPU management policies. The only enhancement needed to be done is to:
- interrogate the setting of the isolcpus kernel parameter in a programmatic manner during kubelet startup (even in the worst-case scenario it could be done via the os package)
- remove the listed CPU cores from the list of the Node's allocatable CPU pool

The really tricky part is how to control when the aforementioned functionality should be done. As the current CPU Manager does not take into account the isolcpus kernel setting when determining a Nodes allocatable CPU capacity, suddenly changing this in GA would be a backward incompatible change. 
On the other hand, this setting should be a Node-level setting, rather than be tied to any CPU management policy.
Reason is that CPU manager already contains two policies, which again should not be changed in a backward incompatible manner.
Therefore, if respecting isolcpus would be done via the introduction of new CPU management policies, it would require two new variants already at Day1: one for each existing policy (default, static), but respecting the isolcpus kernel setting.
This complexity would only increase with every newly introduced policy, unnecessarily cluttering kubelet's already sizeable configuration catalogue.

Instead, the proposal is to introduce one, new alpha-level feature gate to the kubelet binary, called "RespectIsolCpus". The type of the newly introduced flag should be boolean.
If the flag is defined and also set to true, the Node's allocatable CPU pool is decreased as described above, irrespective of which CPU management policy is configured for kubelet.
If the flag is not defined, or it is explicitly set to false; the configured CPU management policy will continue to work without any changes in its functionality.

### Risks and Mitigations

As the outlined implementation concept is entirely backward compatible, no special risks are foreseen with the introduction of this functionality.

The feature itself could be seen as some kind of mitigation of a larger, more complex issue. If CPU manager would support sub-node level, explicit CPU pooling; this feature might not even be needed.
This idea was discussed multiple times, but was always put on hold by the community due to the many risks it would have raised on the Kubernetes ecosystem.

By making kubelet configurable to respect isolcpus kernel parameter cloud infrastructure operators would still be able to achieve their functional requirements, but without any of the drawbacks on Kubernetes core.

## Graduation Criteria

This feature is imagined to be a configurable feature even after graduation.
What is described in the implementation design section could be considered as the first phase of the feature.
Nevertheless, multiple optional enhancements can be imagined if the community is open to them:
- graduating the alpha feature gate to a GA kubelet configuration flag
- explicitly configuring the pool of CPU cores kubelet can manage, rather than subtracting the ones listed in isolcpus from the total capacity of the node
- dynamically adjusting the pool of CPUs kubelet can manage by searching for the presence of a variety of other OS settings, kernel settings, systemd settings, Openstack component configurations etc. on the same node

## Implementation History

N/A

## Alternatives

Some alternatives were already mentioned throughout the document together with their drawbacks, namely:
- enhancing kubelet's CPU manager with topology information, and CPU pool management
- implementing a new isolcpus-respecting variant for each currently supported CPU management policy

Other alternatives were not considered at the time of writing this document.
