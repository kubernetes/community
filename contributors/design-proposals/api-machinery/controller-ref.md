# ControllerRef proposal

* Authors: gmarek, enisoc
* Last edit: [2017-02-06](#history)
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
* [History](#history)

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
  We also want to expose the mapping so that controllers can be in multiple
  processes (e.g. for HA of kube-controller-manager) and separate binaries
  (e.g. for controllers that are API extensions).
  Therefore, we will persist the mapping from each object to its controller in
  the API object itself.

* A secondary goal of ControllerRef is to provide back-links from a given object
  to the controller that manages it, which can be used for:
  * Efficient object->controller lookup, without having to list all controllers.
  * Generic object grouping (e.g. in a UI), without having to know about all
    third-party controller types in advance.
  * Replacing certain uses of the `kubernetes.io/created-by` annotation,
    and potentially enabling eventual deprecation of that annotation.
    However, deprecation is not being proposed at this time, so any uses that
    remain will be unaffected.

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

* ControllerRef is not intended to replace [selector generation](selector-generation.md),
  used by some controllers like Job to ensure all selectors are unique
  and prevent overlapping selectors from occurring in the first place.

  However, ControllerRef will still provide extra protection and consistent
  cross-controller semantics for controllers that already use selector
  generation. For example, selector generation can be manually overridden,
  which leaves open the possibility of overlapping selectors due to user error.

* This proposal does not change how cascading deletion works.

  Although ControllerRef will extend OwnerReference and rely on its machinery,
  the [Garbage Collector](garbage-collection.md) will continue to implement
  cascading deletion as before.
  That is, the GC will look at all OwnerReferences without caring whether a
  given OwnerReference happens to be a ControllerRef or not.

# API

The `Controller` API field in OwnerReference marks whether a given owner is a
managing controller:

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

This section summarizes the intended behavior for existing controllers.
It can also serve as a guide for respecting ControllerRef when writing new
controllers.

## The Three Laws of Controllers

All controllers that manage collections of objects should obey the following
rules.

1. **Take ownership**

   A controller should claim *ownership* of any objects it creates by adding a
   ControllerRef, and may also claim ownership of an object it didn't create,
   as long as the object has no existing ControllerRef (i.e. it is an *orphan*).

1. **Don't interfere**

   A controller should not take any action (e.g. edit/scale/delete) on an object
   it does not own, except to [*adopt*](#adoption) the object if allowed by the
   First Law.

1. **Don't share**

   A controller should not count an object it does not own toward satisfying its
   desired state (e.g. a certain number of replicas), although it may include
   the object in plans to achieve its desired state (e.g. through adoption)
   as long as such plans do not conflict with the First or Second Laws.

## Adoption

If a controller finds an orphaned object (an object with no ControllerRef) that
matches its selector, it may try to adopt the object by adding a ControllerRef.
Note that whether or not the controller *should* try to adopt the object depends
on the particular controller and object.

Multiple controllers can race to adopt a given object, but only one can win
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

Many controllers use watches to *sync* each controller instance (prompting it to
reconcile desired and actual state) as soon as a relevant event occurs for one
of its controlled objects, as well as to let controllers wait for asynchronous
operations to complete on those objects.
The controller subscribes to a stream of events about controlled objects
and routes each event to a particular controller instance.

Previously, the controller used only label selectors to decide which
controller to route an event to. If multiple controllers had overlapping
selectors, events might be misrouted, causing the wrong controllers to sync.
Controllers could also freeze because they keep waiting for an event that
already came but was misrouted, manifesting as `kubectl` commands that hang.

Some controllers introduced a workaround to break ties. For example, they would
sort all controller instances with matching selectors, first by creation
timestamp and then by name, and always route the event to the first controller
in this list. However, that did not prevent misrouting if the overlapping
controllers were of different types. It also only worked while controllers
themselves assigned ownership over objects using the same tie-break rules.

Now that controller ownership is defined in terms of ControllerRef,
controllers should use the following guidelines for responding to watch events:

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

## Default garbage collection policy

Controllers that used to rely on client-side cascading deletion should set a
[`DefaultGarbageCollectionPolicy`](https://github.com/kubernetes/kubernetes/blob/dd22743b54f280f41e68f206449a13ca949aca4e/pkg/genericapiserver/registry/rest/delete.go#L43)
of `rest.OrphanDependents` when they are updated to implement ControllerRef.

This ensures that deleting only the controller, without specifying the optional
`DeleteOptions.OrphanDependents` flag, remains a non-cascading delete.
Otherwise, the behavior would change to server-side cascading deletion by
default as soon as the controller manager is upgraded to a version that performs
adoption by setting ControllerRefs.

Example from [ReplicationController](https://github.com/kubernetes/kubernetes/blob/9ae2dfacf196ca7dbee798ee9c3e1663a5f39473/pkg/registry/core/replicationcontroller/strategy.go#L49):

```go
// DefaultGarbageCollectionPolicy returns Orphan because that was the default
// behavior before the server-side garbage collection was implemented.
func (rcStrategy) DefaultGarbageCollectionPolicy() rest.GarbageCollectionPolicy {
	return rest.OrphanDependents
}
```

New controllers that don't have legacy behavior to preserve can omit this
controller-specific default to use the [global default](https://github.com/kubernetes/kubernetes/blob/2bb1e7581544b9bd059eafe6ac29775332e5a1d6/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go#L543),
which is to enable server-side cascading deletion.

## Controller-specific behavior

This section lists considerations specific to a given controller.

* **ReplicaSet/ReplicationController**

  * These controllers currently only enable ControllerRef behavior when the
    Garbage Collector is enabled. When ControllerRef was first added to these
    controllers, the main purpose was to enable server-side cascading deletion
    via the Garbage Collector, so it made sense to gate it behind the same flag.

    However, in order to achieve the [goals](#goals) of this proposal, it is
    necessary to set ControllerRefs and perform adoption/orphaning regardless of
    whether server-side cascading deletion (the Garbage Collector) is enabled.
    For example, turning off the GC should not cause controllers to start
    fighting again. Therefore, these controllers will be updated to always
    enable ControllerRef.

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

  * CronJob [does not use watches](https://github.com/kubernetes/kubernetes/blob/9ae2dfacf196ca7dbee798ee9c3e1663a5f39473/pkg/controller/cronjob/cronjob_controller.go#L20),
    so [that section](#watches) doesn't apply.
    Instead, all CronJobs are processed together upon every "sync".
  * CronJob applies a `created-by` annotation to link Jobs to the CronJob that
    created them.
    If a ControllerRef is found, it should be used instead to determine this
    link.

## Created-by annotation

Aside from the change to CronJob mentioned above, several other uses of the
`kubernetes.io/created-by` annotation have been identified that would be better
served by ControllerRef because it tracks who *currently* controls an object,
not just who originally created it.

As a first step, the specific uses identified in the [Implementation](#implementation)
section will be augmented to prefer ControllerRef if one is found.
If no ControllerRef is found, they will fall back to looking at `created-by`.

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
* Controllers might create additional objects because they start to respect the
  ["Don't share"](#behavior) rule.

If there are controllers with overlapping selectors at the time of a
*downgrade*:

* Controllers may begin to fight and thrash objects.
* The ownership of existing objects might change due to ignoring ControllerRef.
* Controllers might delete objects because they stop respecting the
  ["Don't share"](#behavior) rule.

# Implementation

Checked items had been completed at the time of the [last edit](#history) of
this proposal.

* [x] Add API field for `Controller` to the `OwnerReference` type.
* [x] Add validator that prevents an object from having multiple ControllerRefs.
* [x] Add `ControllerRefManager` types to encapsulate ControllerRef manipulation
  logic.
* [ ] Update all affected controllers to respect ControllerRef.
  * [ ] ReplicationController
    * [ ] Don't touch controlled objects if DeletionTimestamp is set.
      * [x] Don't adopt/manage objects.
      * [ ] Don't orphan objects.
    * [x] Include ControllerRef on all created objects.
    * [x] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Enable ControllerRef regardless of `--enable-garbage-collector` flag.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] ReplicaSet
    * [ ] Don't touch controlled objects if DeletionTimestamp is set.
      * [x] Don't adopt/manage objects.
      * [ ] Don't orphan objects.
    * [x] Include ControllerRef on all created objects.
    * [x] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Enable ControllerRef regardless of `--enable-garbage-collector` flag.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] StatefulSet
    * [ ] Don't touch controlled objects if DeletionTimestamp is set.
    * [ ] Include ControllerRef on all created objects.
    * [ ] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] DaemonSet
    * [x] Don't touch controlled objects if DeletionTimestamp is set.
    * [ ] Include ControllerRef on all created objects.
    * [ ] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] Deployment
    * [x] Don't touch controlled objects if DeletionTimestamp is set.
    * [x] Include ControllerRef on all created objects.
    * [x] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [x] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] Job
    * [x] Don't touch controlled objects if DeletionTimestamp is set.
    * [ ] Include ControllerRef on all created objects.
    * [ ] Set DefaultGarbageCollectionPolicy to OrphanDependents.
    * [ ] Use ControllerRefManager to adopt and orphan.
    * [ ] Use ControllerRef to map watch events to controllers.
  * [ ] CronJob
    * [ ] Don't touch controlled objects if DeletionTimestamp is set.
    * [ ] Include ControllerRef on all created objects.
    * [ ] Set DefaultGarbageCollectionPolicy to OrphanDependents.
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

# History

Summary of significant revisions to this document:

* 2017-02-06 (enisoc)
  * [Controller-specific behavior](#controller-specific-behavior)
    * Enable ControllerRef regardless of whether GC is enabled.
  * [Implementation](#implementation)
    * Audit whether existing controllers respect DeletionTimestamp.
* 2017-02-01 (enisoc)
  * Clarify existing specifications and add details not previously specified.
  * [Non-goals](#non-goals)
    * Make explicit that overlapping selectors are still user error.
  * [Behavior](#behavior)
    * Summarize fundamental rules that all new controllers should follow.
    * Explain how the validator prevents multiple ControllerRefs on an object.
    * Specify how ControllerRef should affect the use of watches/expectations.
    * Specify important controller-specific behavior for existing controllers.
    * Specify necessary changes to default GC policy when adding ControllerRef.
    * Propose changing certain uses of `created-by` annotation to ControllerRef.
  * [Upgrading](#upgrading)
    * Specify ControllerRef-related behavior changes upon upgrade/downgrade.
  * [Implementation](#implementation)
    * List all work to be done and mark items already completed as of this edit.
