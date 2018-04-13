---
kep-number: 8
title: Topology-aware Workload Controllers
authors:
  - "@janetkuo"
owning-sig: sig-apps
participating-sigs:
  - sig-scheduling
reviewers:
  - "@kow3ns"
  - "@bsalamat"
approvers:
  - "@erictune"
  - "@davidopp"
editor: "@janetkuo"
creation-date: 2018-04-13
last-updated: 2018-04-13
status: provisional
see-also:
  - n/a
replaces:
  - n/a
superseded-by:
  - n/a
---

# Topology-aware Workload Controllers

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Summary](#summary)
- [Motivation](#motivation)
  * [Goals](#goals)
  * [Non-Goals](#non-goals)
- [Background](#background)
  * [Spreading](#spreading)
  * [Update Patterns](#update-patterns)
- [Proposal](#proposal)
  * [User Stories](#user-stories)
    + [Story 1: Update By Failure Domain](#story-1-update-by-failure-domain)
    + [Story 2: Update Across Failure Domain (PDB Per Failure Domain)](#story-2-update-across-failure-domain-pdb-per-failure-domain)
    + [Story 3: Update Across Failure Domain (Overall PDB)](#story-3-update-across-failure-domain-overall-pdb)
    + [Story 4: Scale Down](#story-4-scale-down)
  * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
  * [Risks and Mitigations](#risks-and-mitigations)
- [Graduation Criteria](#graduation-criteria)
- [Implementation History](#implementation-history)
- [Advantages](#advantages)
- [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
  * [New Update Strategy](#new-update-strategy)
    + [Advantages](#advantages-1)
    + [Drawbacks](#drawbacks-1)
  * [Multiple Workloads Controllers (One For Each Failure Domain)](#multiple-workloads-controllers-one-for-each-failure-domain)
    + [Advantages](#advantages-2)
    + [Drawbacks](#drawbacks-2)
  * [A New Workloads Controller That Manages Multiple Failure Domain Controllers](#a-new-workloads-controller-that-manages-multiple-failure-domain-controllers)
    + [Advantages](#advantages-3)
    + [Drawbacks](#drawbacks-3)

*TODO: regenerate TOC before merging (https://github.com/ekalinin/github-markdown-toc)*

## Summary

This KEP answers two questions:

* How do we extend topology requests to specify update patterns, not just pod
  placement? 
* How do we ensure that topology constraints and preferences are maintained
  across updates, controller scale up/down?

## Motivation

Topology is arbitrary (see [#41442][]). Kubernetes users today can perform a
rolling update on their workloads. However, none of current update patterns
takes topology requests into account. When controllers are scaled down, topology
requests are ignored, too (see [#4301][]).

[#41442]: https://github.com/kubernetes/kubernetes/issues/41442
[#4301]: https://github.com/kubernetes/kubernetes/issues/4301

### Goals

* Extend topology requests to update patterns of all workload controllers,
  including customized ones.
* Make sure topology constraints and preferences are maintained across workload
  controllers updates, scale up/down.

### Non-Goals

Other topology related design discussions, such as services and volumes, are out
of scope. 

## Background

Here is how pods placement (scheduling) and deletion is managed today:

| Controller  | Creation time                   | Scale up            | Scale down                          | Rolling update                               |
|-------------|---------------------------------|---------------------|-------------------------------------|----------------------------------------------|
| DaemonSet   | DaemonSet or scheduler          | -                   | -                                   | DaemonSet (kills pods by health)             |
| StatefulSet | Scheduler                       | Scheduler           | StatefulSet (kills pods by ordinal) | StatefulSet (kills pods by ordinal)          |
| ReplicaSet  | Scheduler                       | Scheduler           | ReplicaSet (kills pod by health)    | -                                            |
| Deployment  | ReplicaSet creation or scale up | ReplicaSet scale up | Replicasets scale down              | ReplicaSet scale up + ReplicaSets scale down |

### Spreading 

First, DaemonSet is special, because it is the only workload controller whose
pods don’t go through default scheduler. In the future, DaemonSet may create
pods that get scheduled by the default scheduler. Either way, DaemonSet pods
spread equally, because a DaemonSet schedules one pod per node. 

For other workload controllers, during pods placement (pods creation), the
default scheduler offers best-effort spreading (`SelectorSpreadPriority`) and
explicit spreading rule (pod anti-affinity). 

However, when scaling down, the controller decides which pods to kill first,
without going through default scheduler. Therefore, scaling down ignores
scheduling policy. 

Controller rolling update is done by killing old replicas and creating new ones.
Similar to scaling down cases, when old replicas are brought down, the decision
doesn’t go through the default scheduler, either. On the other hand, the
creation of new replicas goes through the default scheduler, so the spreading
only works after the rolling update is done. 

### Update Patterns 

The workload controllers that support rolling update feature offer customizable
update strategies. Users can specify the number of old replicas to bring down or
the number of new replicas to create during a rollout. 

Existing update patterns are not topology-aware. For example, users cannot
specify things like: 

* Update by failure domain: Update one failure domain at a time (staged domain)
* Update across failure domains: Only allow X number of replicas to be down per
  failure domain 

## Proposal

We can think of updates as planned disruptions. When a workload controller
performs a rolling update, it will evict (instead of deleting) old pods and
creates new pods. Users will be able to use [`PodDisruptionBudget`][](PDB) to
control all sorts of planned disruptions, including updates. 

To achieve this, PDB should allow users to specify allowed pod disruptions for
a collection of nodes (such as the nodes in one failure domain), using a new PDB
field `.spec.nodeSelectorTerm`. 

This requires changes to workloads controllers code (change pod deletions to
evictions), PDB API, and PDB controller code.

[`PodDisruptionBudget`]: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/

### User Stories

#### Story 1: Update By Failure Domain 

When a user wants to update one failure domain at a time, they can create a PDB
to disallow any disruptions in failure domains other than the first domain to
update:

```yaml
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: app-pdb
spec:
  maxUnavailable: 0
  selector:
    matchLabels:
      app: foo
  nodeSelectorTerm:
      # Only allow evicting/updating pods in zone “bar”
      matchExpressions:
      - key: failure-domain.beta.kubernetes.io/zone
        operator: NotIn
        values: 
        - bar
      - key: failure-domain.beta.kubernetes.io/zone
        operator: Exists
```

#### Story 2: Update Across Failure Domain (PDB Per Failure Domain)

When a user wants to update across failure domains, and only allow one pod per
failure domain to be unavailable at a time, they can specify:

```yaml
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: app-pdb
spec:
  maxUnavailable: 1
  selector:
    matchLabels:
      app: foo
  nodeSelectorTerm:
      # Only allow 1 pod to be evicted/updated per zone
      matchExpressions:
      - key: failure-domain.beta.kubernetes.io/zone
        operator: Exists
```

#### Story 3: Update Across Failure Domain (Overall PDB)

When a user wants to update across failure domains and only allow one pod to be
unavailable at a time in all failure domains, they can specify:

```yaml
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: app-pdb
spec:
  maxUnavailable: 1
  selector:
    matchLabels:
      app: foo
```

#### Story 4: Scale Down 

When a user wants to scale down a workload controller, they can also use PDB to
limit overall disruptions or in specific failure domains, similar to Story 3. 

### Implementation Details/Notes/Constraints

* We don't support more than 1 PDB rule covering the same pod today. This needs
  to be updated. 
* The topology-aware PDB's maxUnavailable and minAvaiable is calculated based on
  (number of replicas / number of domains).
* Update one failure domain at a time is manual. It might be fine because it’s
  an advanced use case.

### Risks and Mitigations

Controllers will be changed to evict pods instead of delete pods.

Risks:

* If a pod eviction fails because of PDB, the controller may enter eviction
  hotloop.
* The controller may create more API calls than it should, because the
  controller isn’t aware of which pods to evict first (the controller won’t
  look at PDB before deciding which pods to evict).

Mitigations:

* We need to make sure the controllers won't enter eviction hotloop with enough
  tests, and we cannot do things like, always evict the same pods.
* Need to minimize and benchmark additional API calls.

## Graduation Criteria

This can be promoted to beta when all Kubernetes workload controllers can switch
to use pod evictions and has no regressions.

This will be promoted to GA once it's gone a sufficient amount of time as beta
with no changes.

## Implementation History

TBD

## Advantages

* Work with all workloads controllers (including custom controller) that manage
  pods directly or indirectly. There’s no need to implement topology update
  logic in controllers. 
* No need to add complex logic to Deployments and ReplicaSets to add
  topology-awareness to Deployment’s update strategy.

## Drawbacks

* User experience: users need to create PDBs for controllers update strategy,
  it might be confusing.

## Alternatives

### New Update Strategy 

Add a new update strategy for each workload controller. Each workload controller
manages its own new update pattern. 

A drawback of this solution is that all controllers need to implement this
logic themselves. 

Moreover, this new update strategy is not trivial to implement for Deployments,
because Deployment doesn’t manage pods directly. Deployment performs a rolling
update by scaling up/down its ReplicaSets. To implement this feature, Deployment
needs to be able to specify failure domains when scaling up/down ReplicaSets.
Therefore, we’ll need to add a new topology field to ReplicaSet, for the
Deployment to tell the ReplicaSet in which failure domain it should create or
delete pods. ReplicaSet controller also needs to have the logic for creating and
deleting pods of specific failure domains. 

#### Advantages

* User experience: users set the update strategy in each controller. It’s used
  the same way as existing update strategy. It’s easier to learn.

#### Drawbacks

* It’s complicated to implement for Deployments. 
* Need to implement this feature for each controller, including custom
  controllers. 

### Multiple Workloads Controllers (One For Each Failure Domain)

Users can create one workload controller for each failure domain, and update one
failure domain at a time by updating one controller at a time. 

#### Advantages

* Works today. Don’t need to change anything. 

#### Drawbacks

* Users need to manage one workloads for each failure domain instead of one for
  all. 
* Users need to update failure domain controllers when topology changes.

### A New Workloads Controller That Manages Multiple Failure Domain Controllers

This is a follow up on the previous alternative (multiple workloads controllers,
one for each failure domain). User can use a single, top-level controller to
manage multiple controllers based on failure domain. 

Because workloads controllers that support rolling update feature also support
rollout history and rollback, we need to implement history and rollback for this
top-level controller. This top-level controller will create its own controller
history and children controllers for specific failure domains. Children
controllers aren’t expected to be controlled (updated or rolled back) by the
users, and any changes will be reverted/reconciled by the top-level controller
based on the top-level controller’s spec. Autoscaler should not target children
controllers directly, either. 

#### Advantages

* Automate the management of multiple controllers with the same template but
  different failure domain. 

#### Drawbacks

* Need to add another set of resources and controllers in Kubernetes core. 
* Each existing workloads controller needs to have a new top-level controller
  that manages it. Doesn’t work automatically for existing controllers.
