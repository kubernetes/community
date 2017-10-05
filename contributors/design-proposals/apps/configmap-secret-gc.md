# Garbage Collection of ConfigMaps and Secrets

**Author**: Janet Kuo (@janetkuo)

**Status**: Proposal

- [Abstract](#abstract)
- [Requirements](#requirements)
- [API Changes](#api-changes)
- [Controllers](#controllers)
  * [Core Controllers that Support ConfigMaps/Secrets Garbage Collection](#core-controllers-that-support-configmapssecrets-garbage-collection)
  * [Update ConfigMap/Secret's Owner List](#update-configmapsecrets-owner-list)
    + [On Controller API Resource Creation](#on-controller-api-resource-creation)
    + [On Controller API Resource Template Updates](#on-controller-api-resource-template-updates)
- [Alternatives](#alternatives)
  * [Embed ConfigMaps/Secrets in Controller Spec](#embed-configmapssecrets-in-controller-spec)
  * [Specify ConfigMaps/Secrets Garbage Collection in Controller Spec](#specify-configmapssecrets-garbage-collection-in-controller-spec)
  * [Specify ConfigMaps/Secrets garbage collection in PodSpec](#specify-configmapssecrets-garbage-collection-in-podspec)
  * [Automatic ConfigMap/Secret References Updates](#automatic-configmapsecret-references-updates)
- [Caveats](#caveats)
- [Testing Plan](#testing-plan)


## Abstract

ConfigMaps allow users to decouple configuration artifacts from image content to
keep containerized applications portable. The data stored in ConfigMaps can be
consumed in pods or provide the configurations for system components. Secrets
are similar to ConfigMaps, but are intended to hold sensitive information. 

To do a rolling update of a ConfigMap/Secret, users need to create a new
ConfigMap, update its references in workloads resources (such as Deployment),
and delete the old ConfigMap. To facilitate ConfigMap/Secrets rollouts and
management ([#22368](https://github.com/kubernetes/kubernetes/issues/22368)), we
propose a mechanism that supports garbage collection of unused ConfigMaps and
Secrets. 


## Requirements

1.  Garbage collection of a ConfigMap/Secret should be easy to configure,
    manage, and understand. 
1.  If a user wants a ConfigMap/Secret to be garbage collected, the intent
    should be easily discoverable by other users who consume the same
    ConfigMap/Secret. No surprises. 
1.  The design of garbage collection of ConfigMaps/Secrets should cover
    Deployments, StatefulSets, DaemonSets, ReplicaSets, and
    ReplicationControllers. 
1.  For controllers that support history and rollback feature:
    ConfigMaps/Secrets should not be garbage collected until the referencing
    controller history object is removed, so that rollback still works. 
1.  Garbage collection of ConfigMaps/Secrets should work after a controller's
    pod template is updated.
1.  Users can only enable or disable the garbage collection of a
    ConfigMap/Secret when creating the ConfigMap/Secret.


## API Changes

A new *immutable* API field is introduced to ConfigMap and Secret API. 


```go
type ConfigMap struct {
  // If true, the supported controllers (Deployments, StatefulSets, DaemonSets,
  // ReplicaSets, and ReplicationControllers) that reference this ConfigMap will
  // become its owner. 
  // Note that garbage collector will delete this ConfigMap once its owners are 
  // gone, even if this ConfigMap is still referenced by other API resources
  // (such as Jobs).
  // This can only be specified at creation time.
  // Default to false.
  ControllersAsOwners *bool
}
```


```go
type Secret struct {
  // If true, the supported controllers (Deployments, StatefulSets, DaemonSets,
  // ReplicaSets, and ReplicationControllers) that reference this Secret will
  // become its owner. 
  // Note that garbage collector will delete this ConfigMap once its owners are 
  // gone, even if this ConfigMap is still referenced by other API resources
  // (such as Jobs).
  // This can only be specified at creation time.
  // Default to false.
  ControllersAsOwners *bool
}
```

ConfigMap and Secret garbage collection are disabled by default for backward
compatibility. 


## Controllers

This section is presented as a generalization of how an arbitrary controller can
use this new API field to garbage collect the ConfigMaps/Secrets it references. 

A controller can have the ConfigMaps/Secrets it references garbage collected by
declaring ownership of those ConfigMaps/Secrets, i.e. set the ConfigMaps/Secrets
(`.metadata.ownerReferences`).
The [garbage collector](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)
will then delete the ConfigMaps/Secrets when the owners are all gone. Note that
garbage collector only deletes objects that once have owners but no longer do.
ConfigMaps/Secrets won't be garbage collected if they never have any owner. 

Most workload controllers support rolling update, revision history and rollback
features. To make sure rollback still works, for a controller with revision
history feature, the ConfigMaps/Secrets ownership should belong to the
controller's history instead of the controller itself. For example, a Deployment
should not become the owner of any referenced ConfigMaps/Secrets, but the
Deployment's history objects (ReplicaSets) will own those ConfigMaps/Secrets.

If a controller does not support rolling update, revision history, or rollback
(such as ReplicaSets and ReplicationControllers), it should mark itself as the
owner of the ConfigMaps/Secrets it references in order to have them garbage
collected when the controller is gone. 


### Core Controllers that Support ConfigMaps/Secrets Garbage Collection

*   Controllers with history: Deployments, StatefulSets, DaemonSets
*   Controllers without history: ReplicaSets, ReplicationControllers 


### Update ConfigMap/Secret's Owner List

There are 2 possible scenarios when a referenced ConfigMap/Secret's owner list
(`.metadata.ownerReferences`) needs to be updated:

1.  After a controller API resource is created.
1.  After a controller API resource's pod template is updated, and: 
    1.  The controller's current history object is created, or 
    1.  The ConfigMap/Secret references in the pod template are changed. 

Note that Deployment is a special case here. Because a Deployment's revision
history (ReplicaSets) can declare ConfigMaps/Secrets ownership directly, we
don't need to implement it in Deployment controller code. Deployment controller
delegates the work to ReplicaSet controller. 

#### On Controller API Resource Creation

When a controller API object (with or without history) exists, the controller
will do the following:

1.  Find all referenced ConfigMaps/Secrets of the following:
    1.  Controller with history (e.g. StatefulSets and DaemonSets): find the
        ConfigMaps/Secrets referenced in the pod templates of current revision
        history. 
    1.  Controller without history (e.g. ReplicaSets and ReplicationControlelrs):
        find the ConfigMaps/Secrets referenced in its own pod template.
1.  Check the value of `.controllersAsOwner` field of each such ConfigMap/Secret.
1.  If the field is set to `true`, add the following to the ConfigMap/Secret's
    owner list `.metadata.ownerReferences` (if it's not already in the list). If
    it is set to `false`, remove instead. (We don't need to remove if the field
    is immutable.)
    1.  Controller with history: add/remove the controller's current revision
        history to/from the owner list.
    1.  Controller without history: add the controller itself to the above owner
        list (no need to remove, because the newly created controller won't be
        already in the owner list).


#### On Controller API Resource Template Updates

When a controller API resource's pod template is updated, the controller will do
the following:

1.  Check if ConfigMap/Secret referenced in the pod template is changed or not. 
    1.  Controller with history: if a new history object is created, add the
        controller's current revision history to the owner list.
    1.  Controller without history: if the previous pod template referenced
        different sets of ConfigMaps/Secrets, add the controller itself to the
        owner list of current referenced ConfigMaps/Secrets, and remove from the
        previously referenced but no longer referenced ConfigMaps/Secrets. 


## Alternatives

Other alternatives considered are listed below. 


### Embed ConfigMaps/Secrets in Controller Spec

Add `.configMapTemplate` and `.secretTemplate` fields to a controller's `.spec`. 

**Cons:**

*   Adds complexity to the API, tooling dealing with ConfigMaps/Secrets, and the
    implementation of controllers.
*   Controller spec becomes fat.


### Specify ConfigMaps/Secrets Garbage Collection in Controller Spec

Add a `bool` field in controller `.spec` to specify whether garbage collection
of its ConfigMaps/Secrets is enabled. For more granularity, add another field
allowing users to specify a list of ConfigMaps/Secrets that should not be
garbage collected even if the aforementioned field is enabled. 

**Cons:**

*   User intent conflict. For example, user A wants a ConfigMap to be garbage
    collected with a Deployment, and therefore the ConfigMap will be deleted
    after the Deployment is gone. User B cannot stop the ConfigMap from being
    deleted even if B disables garbage collection of that ConfigMap in B's
    controller spec. 
*   To enable or disable garbage collection of a ConfigMap/Secret, users need to
    update spec of all controllers that reference it. 


### Specify ConfigMaps/Secrets garbage collection in PodSpec

Add a `bool` field everywhere in `PodSpec` that ConfigMaps/Secrets are
referenced (e.g. `.containers.envFrom.configMapRef`) to specify whether garbage
collection of the ConfigMap/Secret is enabled. 

**Cons:**

*   Verbosity. if a ConfigMap is referenced in multiple places in a pod template,
    users need to specify it everywhere. 
*   User intent conflict.
*   To enable or disable garbage collection of a ConfigMap/Secret, users need to
    update pod template of all controllers that reference it. For controllers
    with history, this means triggering rollouts.


### Automatic ConfigMap/Secret References Updates

Every time a ConfigMap/Secret is updated, the controller automatically creates a
new ConfigMap/Secret (the new name can be generated by appending a hash), and
updates references in pod templates in controller history objects. Every
ConfigMap/Secret update triggers a rollout.

For example, make Deployment controller create an ConfigMap "abc-<hash>" from
ConfigMap "abc" referenced in its pod template, and then reference "abc-<hash>"
in ReplicaSet pod template. Or something like
[#31701](https://github.com/kubernetes/kubernetes/pull/31701).

This is not strictly about garbage collection of ConfigMaps/Secrets, but is
relevant to the garbage collection design. 

**Pros:**

*   ConfigMap/Secret references are updated automatically.

**Cons:**

*   It becomes tricky for controllers (e.g. Deployments) to compare its pod
    template against its history (e.g. ReplicaSets) to find current history
    (ReplicaSet). 
*   Rollbacks trigger updates to original ConfigMaps. All consumers of a
    ConfigMap/Secret are then forced to do rollbacks together. 


## Caveats

1.  CronJobs and Jobs are not supported. 
1.  Controllers that don't support this feature should not reference
    ConfigMaps/Secrets with the fields be set to `true`. Otherwise,
    ConfigMaps/Secrets may be removed even if the controllers are still
    referencing them. 


## Testing Plan

1.  ConfigMap/Secret `.metadata.ownerReferences` are correct after Deployment,
    StatefulSet, DaemonSet, ReplicaSet and ReplicationController creation and
    template update.
