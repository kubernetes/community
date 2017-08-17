# Priority in ResourceQuota

@resouer @bsalamat

Aug 2017
  * [Objective](#objective)
    * [Non-Goals](#non-goals)
  * [Background](#background)
  * [Overview](#overview)
  * [Detailed Design](#detailed-design)
    * [Expected behavior of ResourceQuota apply to pod](#resourcequota-apply-to-pod)
    * [Expected behavior of ResourceQuota admission controller](#resourcequota-admission-controller)


## Objective

This feature is designed to make ResourceQuota become priority aware, three sub-tasks are included.

* Add a field to ResourceQuota that specifies a priority class name.
* Incorporate priority in quota checks.
* ResourceQuota admission controller to check ResourceQuota priority class and resolve the priority class to its integer value.

### Non-Goals 

* Add priority in Pod spec (this is implemented in: #45610)
* Incorporate priority in pod scheduling logic.

## Background 

Since we already have [priority filed in Pod spec](https://github.com/kubernetes/kubernetes/pull/45610), 
Pods can now be classified into different priority classes. So it is nature to make ResourceQuota apply to Pods based on its priority class. In order to implement this, we need to add a priority filed to ResourceQuota definition.

## Overview 

This design doc introduces the motivation of adding priority filed in ResourceQuota and
how this priority impacts ResourceQuta applying to Pod. 

## Detailed Design 

### Expected behavior of ResourceQuota apply to pod 

The design will be like:

```go
// ResourceQuotaSpec defines the desired hard limits to enforce for Quota
type ResourceQuotaSpec struct {
  // Hard is the set of desired hard limits for each named resource
  // +optional
  Hard ResourceList
  // A collection of filters that must match each object tracked by a quota.
  // If not specified, the quota matches all objects.
  // +optional
  Scopes []ResourceQuotaScope
  // If specified, indicates this ResourceQuota applies to that priority class
  // If not specified, this ResourceQuota applies to all priority classes.
  // +optional
  PriorityClassName string
  // The priority value. Various system components use this field to find the
  // priority of the ResourceQuota. When Priority Admission Controller is enabled, it
  // prevents users from setting this field. The admission controller populates
  // this field from PriorityClassName.
  // +optional
  Priority *int32
}
```

* One could expect when `PriorityClassName` field is set, it indicates that this quota applies to Pod with same priority class. 
* If this field is not present, this quota applies to all priority classes.

It is worth noting that quota is assigned per namespace, not per user, in Kubernetes. Adding priority to ResourceQuota will not change anything with respect to the scope of ResourceQuota. In other words, when priority is defined for a ResourceQuota object, the quota is applied to all pods at that priority in the specified namespace.

Just like pod spec, we also have `Priority *int32` in `ResourceQuta` spec. So the class don't need to be resolved every time it is being used.

### Expected behavior of ResourceQuota admission controller

This part is about how to make admission controller work with priority aware `ResourceQuota`.

* ResourceQuota admission controller should reject creation or update of ResourceQuota object if new ResourceQuota names a PriorityClass that does not exist, i.e. is not a pre-defined or user-specified PriorityClass (but empty is allowed).
* ResourceQuota admission controller should reject creation of ResourceQuota object if another ResourceQuota object with same PriorityClass already exists, or if a ResourceQuota object with no PriorityClass exists.
* ResourceQuota admission controller should reject update of ResourceQuota object if the new (updated) PriorityClass matches that of another ResourceQuota object, or if another ResourceQuota object has no PriorityClass.
* priority admission controller should reject deletion of a PriorityClass object if it is named in any ResourceQuota object in any namespace.
