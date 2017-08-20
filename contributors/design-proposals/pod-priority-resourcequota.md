# Priority in ResourceQuota

Authors:

@resouer @bsalamat @derekwaynecarr 

Sep 2017

  * [Objective](#objective)
    * [Non-Goals](#non-goals)
  * [Background](#background)
  * [Overview](#overview)
  * [Detailed Design](#detailed-design)
    * [Expected behavior of ResourceQuota](#resourcequota)
    * [Expected behavior of ResourceQuota admission controller](#resourcequota-admission-controller)


## Objective

This feature is designed to make ResourceQuota become priority aware, 3 sub-tasks are included.

* Add a field to ResourceQuota that specifies priority class names.
* Incorporate priority in quota checks.
* ResourceQuota admission controller to check ResourceQuota priority class name and make corresponding decision.

### Non-Goals 

* Add priority in Pod spec (this has been implemented in: #45610)

## Background 

Since we already have [priority field in Pod spec](https://github.com/kubernetes/kubernetes/pull/45610), 
Pods can now be classified into different priority classes. We would like to be able to create quota for various priority classes in order to manage cluster resources better and limit abuse scenarios. In order to implement this, we need to include priority class name field to ResourceQuota definition.

## Overview 

This design doc introduces a new field in ResourceQuota to specify a group of priority classes for the quota to match with, and explains how quota enforcement logic is changed to apply the quota to pods with the given priority classes.

## Detailed Design 

### Expected behavior of ResourceQuota

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
  // A list of PriorityClassName with which this quota is expected to match.
  // If specified, the quota's hard limits are restricted to only account resources oriented around pod which matched with given priority classes.
  // If not specified, the quota's hard limits are restricted to only account resources oriented around pods without priority class name.
  // +optional
  PriorityClassNameList []string
}
```
Any `PriorityClassName` in `PriorityClassNameList` of ResourceQuota will be checked against `PriorityClassName` field in the Pod, in detail:

* If `PriorityClassNameList` is present, for any given `PriorityClassName` in the list:
    * This quota's hard limits matches to all pods with the matched priority class.
    * If any `PriorityClassName` is set to *, it means this quota matches all priority classes.
* If `PriorityClassNameList` is not present, it indicates this `ResourceQuota` matches to pods without priority class name.
* The absence of `ResourceQuota` for a particular `PriorityClassName` means no quota at that priority.
* If multiple `ResourceQuota` apply to a Pod, the pod must satisfy all of them.

It is worth noting that quota is assigned per namespace, not per user, in Kubernetes. Adding priority class names to ResourceQuota will not change anything with respect to the scope of ResourceQuota. In other words, when priority class names are defined for a ResourceQuota object, the quota is applied to all pods at those priority classes in the specified namespace.

Unlike pod spec, ResourceQuota does not have Priority `*int32` in its spec. As explained above, `ResourceQuota` with `PriorityClassNameList` matches against pods with the same `PriorityClassName` in list, regardless of the integer value of priority.

### Expected behavior of ResourceQuota admission controller

* `ResourceQuota` admission controller **should not reject** creation or update of `ResourceQuota` object if new `ResourceQuota` names a `PriorityClass` that does not exist, and empty `PriorityClassNameList` field is also allowed. For now, we don't want to introduce extra order by this design.
* `ResourceQuota` admission controller **should not reject** creation of a `ResourceQuota` object if the `PriorityClassNameList` field of an existing `ResourceQuota` object has one or more of the same `PriorityClassNames`. Overlapping quota is allowed in Kubernetes.
* `ResourceQuota` admission controller **should reject** pod with priority set but has no matched ResourceQuota found.
* `ResourceQuota` admission controller **should reject** update of `ResourceQuota` object. Users should create a new quota and delete the old one if the quota needs to be updated.
* `ResourceQuota` admission controller **should not reject** deletion of `ResourceQuota` objects with reference. In other words, a `ResourceQuota` object can be deleted even if there are Pods that the `ResourceQuota` object applies to them.
