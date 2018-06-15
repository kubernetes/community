# Priority in ResourceQuota

Authors:

Harry Zhang [@resouer](https://github.com/resouer)
Vikas Choudhary [@vikaschoudhary16](https://github.com/vikaschoudhary16)

Main Reviewers:

Bobby [@bsalamat](https://github.com/bsalamat)
Derek [@derekwaynecarr](https://github.com/derekwaynecarr)

Dec 2017

  * [Objective](#objective)
    * [Non-Goals](#non-goals)
  * [Background](#background)
  * [Overview](#overview)
  * [Detailed Design](#detailed-design)
    * [Changes in ResourceQuota](#changes-in-resourceQuota)
    * [Changes in Admission Controller configuration](#changes-in-admission-controller-configuration)
    * [Expected behavior of ResourceQuota admission controller and Quota system](#expected-behavior-of-resourcequota-admission-controller-and-resourcequota-system)
      * [Backward Compatibility](#backward-compatibility)
  * [Sample user story 1](#sample-user-story-1)
  * [Sample user story 2](#sample-user-story-2)


## Objective

This feature is designed to make `ResourceQuota` become priority aware, several sub-tasks are included.

1. Expand `Scopes` in `ResourceQuotaSpec` to represent priority class names and corresponding behavior.
2. Incorporate corresponding behavior in quota checking process.
3. Update the `ResourceQuota` admission controller to check priority class name and perform expected admission.

### Non-Goals

* Add priority in Pod spec (this is implemented separately in: [45610](https://github.com/kubernetes/kubernetes/pull/45610))

## Background

Since we already have [priority field in Pod spec](https://github.com/kubernetes/kubernetes/pull/45610),
Pods can now be classified into different priority classes. We would like to be able to create quota for various priority classes in order to manage cluster resources better and limit abuse scenarios.

One approach to implement this is by adding priority class name field to `ResourceQuota` API definition. While this arbitrary field of API object will introduce inflexibility to potential change in future and also not adequate to express all semantics.

Thus, we decide to reuse the existing `Scopes` of `ResourceQuotaSpec` to provide a richer semantics for quota to cooperate with priority classes.

## Overview

This design doc introduces how to define a priority class scope and scope selectors for the quota to match with and explains how quota enforcement logic is changed to apply the quota to pods with the given priority classes.

## Detailed Design

### Changes in ResourceQuota

ResourceQuotaSpec contains an array of filters, `Scopes`, that if mentioned, must match each object tracked by a ResourceQuota.

A new field `scopeSelector` will be introduced.
```go
// ResourceQuotaSpec defines the desired hard limits to enforce for Quota
type ResourceQuotaSpec struct {
        ...

        // A collection of filters that must match each object tracked by a quota.
        // If not specified, the quota matches all objects.
        // +optional
        Scopes []ResourceQuotaScope
        // ScopeSelector is also a collection of filters like Scopes that must match each object tracked by a quota
        // but expressed using ScopeSelectorOperator in combination with possible values.
        // +optional
        ScopeSelector *ScopeSelector
}

// A scope selector represents the AND of the selectors represented
// by the scoped-resource selector terms.
type ScopeSelector struct {
        // A list of scope selector requirements by scope of the resources.
        // +optional
        MatchExpressions []ScopedResourceSelectorRequirement
}

// A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator
// that relates the scope name and values.
type ScopedResourceSelectorRequirement struct {
        // The name of the scope that the selector applies to.
        ScopeName ResourceQuotaScope
        // Represents a scope's relationship to a set of values.
        // Valid operators are In, NotIn, Exists, DoesNotExist.
        Operator ScopeSelectorOperator
        // An array of string values. If the operator is In or NotIn,
        // the values array must be non-empty. If the operator is Exists or DoesNotExist,
        // the values array must be empty.
        // This array is replaced during a strategic merge patch.
        // +optional
        Values []string
}

// A scope selector operator is the set of operators that can be used in
// a scope selector requirement.
type ScopeSelectorOperator string

const (
        ScopeSelectorOpIn           ScopeSelectorOperator = "In"
        ScopeSelectorOpNotIn        ScopeSelectorOperator = "NotIn"
        ScopeSelectorOpExists       ScopeSelectorOperator = "Exists"
        ScopeSelectorOpDoesNotExist ScopeSelectorOperator = "DoesNotExist"
)
```
A new `ResourceQuotaScope` will be defined for matching pods based on priority class names.

```go
// A ResourceQuotaScope defines a filter that must match each object tracked by a quota
type ResourceQuotaScope string

const (
        ...
        ResourceQuotaScopePriorityClass ResourceQuotaScope = "PriorityClass"
)
```

### Changes in Admission Controller Configuration

A new field `MatchScopes` will be added to `Configuration.LimitedResource`. `MatchScopes` will be a collection of one or more of the four newly added priority class based `Scopes` that are explained in above section.

```go
// Configuration provides configuration for the ResourceQuota admission controller.
type Configuration struct {
        ...
        LimitedResources []LimitedResource
}

// LimitedResource matches a resource whose consumption is limited by default.
// To consume the resource, there must exist an associated quota that limits
// its consumption.
type LimitedResource struct {
        ...
        // For each intercepted request, the quota system will figure out if the input object
        // satisfies a scope which is present in this listing, then
        // quota system will ensure that there is a covering quota.  In the
        // absence of a covering quota, the quota system will deny the request.
        // For example, if an administrator wants to globally enforce that
        // a quota must exist to create a pod with "cluster-services" priorityclass
        // the list would include "scopeName=PriorityClass, Operator=In, Value=cluster-services"
        // +optional
        MatchScopes []v1.ScopedResourceSelectorRequirement `json:"matchScopes,omitempty"`
}
```

### Expected Behavior of ResourceQuota Admission Controller and ResourceQuota System
`MatchScopes` will be configured in admission controller configuration to apply quota based on priority class names. If `MatchScopes` matches/selects an incoming pod request, request will be **denied if a Covering Quota is missing**. The meaning of Covering Quota is: any quota which has priority class based `Scopes` that matches/selects the pod in the request. 

Please note that this priority class based criteria will be an **additional** criteria that must be satisfied by covering quota.

For more details, please refer to the `Sample user story` sections at the end of this doc.

#### Backward Compatibility

If a Pod's requested resources are not matched by any of the filters in admission controller configuration's `MatchScopes`, overall behavior for the pod will be same as it is today where `ResourceQuota` has no awareness of priority. In such a case, request will be allowed if no covering `ResourceQuota` is found.

Couple of other noteworthy details:
1. If multiple `ResourceQuota` apply to a Pod, the pod must satisfy all of them.
2. We do not enforce referential integrity across objects. i.e. Creation or updating of ResourceQuota object, scopes of which names a PriorityClass that does not exist, are allowed.

This design also tries to enable flexibility for its configuration. Here are several sample user stories.

#### Sample User Story 1
**As a cluster admin, I want `cluster-services` priority only apply to `kube-system` namespace , so that I can ensure those critical daemons on each node while normal user's workloads will not disrupt that ability.**

To enforce above policy:
1. Admin will create admission controller configuration as below:
```yaml
apiVersion: apiserver.k8s.io/v1alpha1
kind: AdmissionConfiguration
plugins:
- name: "ResourceQuota"
  configuration:
    apiVersion: resourcequota.admission.k8s.io/v1alpha1
    kind: Configuration
    limitedResources:
    - resource: pods
      matchScopes:
      - scopeName: PriorityClass
        operator: In
        values: ["cluster-services"]
```

2. Admin will then create a corresponding resource quota object in `kube-system` namespace:

```shell
$ cat ./quota.yml
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-cluster-services
  spec:
    hard:
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["cluster-services"]

$ kubectl create -f ./quota.yml -n kube-system`
```

In this case, a pod creation will be allowed if:
1. Pod has no priority class and created in any namespace.
2. Pod has priority class other than `cluster-services` and created in any namespace.
3. Pod has priority class `cluster-services` and created in `kube-system` namespace, and passed resource quota check.

Pod creation will be rejected if pod has priority class `cluster-services` and created in namespace other than `kube-system`


#### Sample User Story 2
**As a cluster admin, I want a specific resource quota apply to any pod which has priority been set**

To enforce above policy:
1. Create admission controller configuration:
```yaml
apiVersion: apiserver.k8s.io/v1alpha1
kind: AdmissionConfiguration
plugins:
- name: "ResourceQuota"
  configuration:
    apiVersion: resourcequota.admission.k8s.io/v1alpha1
    kind: Configuration
    limitedResources:
    - resource: pods
      matchScopes:
      - operator : Exists
        scopeName: PriorityClass
```

2. Create resource quota to match all pods where there is priority set

```shell
$ cat ./quota.yml
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-cluster-services
  spec:
    hard:
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["cluster-services"]

$ kubectl create -f ./quota.yml -n kube-system`
```
