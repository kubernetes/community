# ControllerRef proposal

* Authors: gmarek, enisoc
* Last edit: 2017-01-25
* Status: partially implemented

Approvers:
* [ ] briangrant
* [ ] dbsmith

**Table of Contents**

* [Goals](#goals)
* [Non-goals](#non-goals)
* [API](#api)
* [Behavior](#behavior)
* [Upgrading](#upgrading)
* [Implementation](#implementation)
* [Alternatives](#alternatives)

# Goals

* The main goal of ControllerRef (controller reference) is to solve the problem
  of controllers that fight over controlled objects due to overlapping selectors
  (e.g. a ReplicaSet fighting with a ReplicationController over Pods because
  both controllers have label selectors that match those Pods).
  Fighting controllers can [destabilize the apiserver](https://github.com/kubernetes/kubernetes/issues/24433),
  [thrash objects back-and-forth](https://github.com/kubernetes/kubernetes/issues/24152),
  or [cause controller operations to hang](https://github.com/kubernetes/kubernetes/issues/8598).

  We don't want to have just an in-memory solution because we don't want a
  Controller Manager crash to cause a massive reshuffling of controlled objects.
  Therefore, we need to persist the mapping from each object to its controller.

* A secondary goal of ControllerRef is to provide back-links from a given object
  to the controller that manages it, which can be used for:
  * Efficient object->controller lookup, without having to list all controllers.
  * Generic object grouping (e.g. in a UI), without having to know about all
    third-party controller types in advance.
  * Replacing many if not all uses of the `kubernetes.io/created-by` annotation.

# Non-goals

* Overlapping selectors will continue to be considered user error.

  ControllerRef will prevent this user error from destabilizing the cluster or
  causing endless back-and-forth fighting between controllers, but it will not
  make it completely safe to create controllers with overlapping selectors.

  In particular, this proposal does not address cases such as Deployment or
  StatefulSet, in which "families" of orphans may exist that ought to be adopted
  as indivisible units.
  Since multiple controllers may race to adopt orphans, the user must ensure
  selectors do not overlap to avoid breaking up families.
  Breaking up families of orphans could result in corruption or loss of
  Deployment rollout state and history, and possibly also corruption or loss of
  StatefulSet application data.

* ControllerRefs will not be used for cascading deletion.

  Although ControllerRef will extend OwnerReference and rely on its machinery,
  the [Garbage Collector](garbage-collection.md) will continue to implement
  cascading deletion as before, without considering ControllerRefs explicitly.

# API

There will be a new API field in OwnerReference that marks whether a given owner
is a managing controller:

```go
type OwnerReference struct {
    …
    // If true, this reference points to the managing controller.
    // +optional
    Controller *bool
}
```

A ControllerRef is thus defined as an OwnerReference with `Controller=true`.
Each object may have at most one ControllerRef in its list of OwnerReferences.
The validator for OwnerReferences lists will fail any update that would violate
this invariant.

# Behavior

This section summarizes the new behavior for existing controllers.
It can also serve as a guide for respecting ControllerRef when writing new
controllers.

## The Three Laws of Controllers

All controllers that manage collections of objects should obey the following
rules.

1. **Take ownership**

   A controller must claim *ownership* of any objects it creates by adding a
   ControllerRef, and may also claim ownership of an object it didn't create,
   as long as the object has no existing ControllerRef (i.e. it is an *orphan*).

1. **Don't interfere**

   A controller may not take any action (e.g. edit/scale/delete) on an object it
   does not own, except to [*adopt*](#adoption) the object if allowed by the
   First Law.

1. **Don't share**

   A controller should not count an object it does not own toward satisfying its
   desired state (e.g. a certain number of replicas), although it may include
   the object in plans to achieve its desired state (e.g. through adoption)
   as long as such plans do not conflict with the First or Second Laws.

## Adoption

If a controller finds an orphaned object that matches its selector, it should
try to adopt the object by adding a ControllerRef.

Multiple controllers may race to adopt a given object, but only one can win
by being the first to add a ControllerRef to the object's OwnerReferences list.
The losers will see their adoptions fail due to a validation error as explained
[above](#api).

If a controller has a non-nil `DeletionTimestamp`, it must not attempt adoption
or take any other actions except updating its `Status`.
This prevents readoption of objects orphaned by the [orphan finalizer](garbage-collection.md#part-ii-the-orphan-finalizer)
during deletion of the controller.

## Orphaning

When a controller is deleted, the objects it owns will either be orphaned or
deleted according to the normal [Garbage Collection](garbage-collection.md)
behavior, based on OwnerReferences.

In addition, if a controller finds that it owns an object that no longer matches
its selector, it should orphan the object by removing itself from the object's
OwnerReferences list. Since ControllerRef is just a special type of
OwnerReference, this also means the ControllerRef is removed.

## Watches

Many controller managers use watches to *sync* each controller (prompting it to
reconcile desired and actual state) as soon as a relevant event occurs for one
of its controlled objects, as well as to let controllers wait for asynchronous
operations to complete on those objects.
The controller manager subscribes to a stream of events about controlled objects
and routes each event to a particular controller.

Previously, the controller manager used only label selectors to decide which
controller to route an event to. If multiple controllers had overlapping
selectors, events might be misrouted, causing the wrong controllers to sync,
or causing some controllers to freeze because they keep waiting for an event
that already happened.

Some controller managers introduced a workaround that used the controller
creation timestamps and names to break ties. However, that did not prevent
misrouting if the overlapping controllers were of different types.
It also only worked while controllers themselves assigned ownership over objects
using the same tie-break rules.

Now that controller ownership is defined in terms of ControllerRef,
controller managers should use the following guidelines for responding to watch
events:

* If the object has a ControllerRef:
  * Sync only the referenced controller.
  * Update `expectations` counters for the referenced controller.
  * If an *Update* event removes the ControllerRef, sync any controllers whose
    selectors match to give each one a chance to adopt the object.
* If the object is an orphan:
  * *Add* event
    * Sync any controllers whose selectors match to give each one a chance to
      adopt the object.
    * Do *not* update counters on `expectations`.
      Controllers should never be waiting for creation of an orphan because
      anything they create should have a ControllerRef.
  * *Delete* event
    * Do *not* sync any controllers.
      Controllers should never care about orphans disappearing.
    * Do *not* update counters on `expectations`.
      Controllers should never be waiting for deletion of an orphan because they
      are not allowed to delete objects they don't own.
  * *Update* event
    * If labels changed, sync any controllers whose selectors match to give each
      one a chance to adopt the object.

## Controller-specific behavior

This section lists considerations specific to a given controller.

* **StatefulSet**

  * A StatefulSet will not adopt any Pod whose name does not match the template
    it uses to create new Pods: `{statefulset name}-{ordinal}`.
    This is because Pods in a given StatefulSet form a "family" that may use pod
    names (via their generated DNS entries) to coordinate among themselves.
    Adopting Pods with the wrong names would violate StatefulSet's semantics.

    Adoption is allowed when Pod names match, so it remains possible to orphan a
    family of Pods (by deleting their StatefulSet without cascading) and then
    create a new StatefulSet with the same name and selector to adopt them.

* **CronJob**

  * CronJob does not use [watches](#watches), so that section doesn't apply.
    Instead, all CronJobs are processed together upon every "sync".
  * CronJob applies a `created-by` annotation to link Jobs to the CronJob that
    created them.
    If a ControllerRef is found, it should be used instead to determine this
    link.

# Upgrading

In the absence of controllers with overlapping selectors, upgrading or
downgrading the master to or from a version that introduces ControllerRef
should have no user-visible effects.
If no one is fighting, adoption should always succeed eventually, so ultimately
only the selectors matter on either side of the transition.

If there are controllers with overlapping selectors at the time of an *upgrade*:

* Back-and-forth thrashing should stop after the upgrade.
* The ownership of existing objects might change due to races during
  [adoption](#adoption). As mentioned in the [non-goals](#non-goals) section,
  this can include breaking up families of objects that should have stayed
  together.
* Controllers might create additional objects due to the new
  ["Don't share"](#behavior) rule.

If there are controllers with overlapping selectors at the time of a
*downgrade*:

* Controllers may begin to fight and thrash objects.
* The ownership of existing objects might change due to ignoring ControllerRef.
* Controllers might delete objects due to rollback of the
  ["Don't share"](#behavior) rule.

# Implementation

Checked items had been completed at the time of the last edit of this proposal
(2017-01-25).

* [x] Add API field for `Controller` to the `OwnerReference` type.
* [x] Add validator that prevents an object from having multiple ControllerRefs.
* [x] Add `ControllerRefManager` types to encapsulate ControllerRef manipulation
  logic.
* [ ] Update all affected controllers to respect ControllerRef.
  * [ ] ReplicationController
    * [x] Include ControllerRef on all created objects.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] ReplicaSet
    * [x] Include ControllerRef on all created objects.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] StatefulSet
    * [ ] Include ControllerRef on all created objects.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] DaemonSet
    * [ ] Include ControllerRef on all created objects.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] Deployment
    * [x] Include ControllerRef on all created objects.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] Job
    * [ ] Include ControllerRef on all created objects.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] CronJob
    * [ ] Include ControllerRef on all created objects.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map Jobs to their parent CronJobs.
* [ ] Tests
  * [ ] Update existing controller tests to use ControllerRef.
  * [ ] Add test for overlapping controllers of different types.
* [ ] Replace or augment uses of `CreatedByAnnotation` with ControllerRef.
  * [ ] `kubectl describe` list of controllers for an object.
  * [ ] `kubectl drain` Pod filtering.
  * [ ] Classifying failed Pods in e2e test framework.

# Alternatives

The following alternatives were considered:

* Centralized "ReferenceController" component that manages adoption/orphaning.

  Not chosen because:
  * Hard to make it work for all imaginable 3rd party objects.
  * Adding hooks to framework makes it possible for users to write their own
    logic.

* Separate API field for `ControllerRef` in the ObjectMeta.

  Not chosen because:
  * Complicated relationship between `ControllerRef` and `OwnerReference`
    when it comes to deletion/adoption.


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/controller-ref.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
