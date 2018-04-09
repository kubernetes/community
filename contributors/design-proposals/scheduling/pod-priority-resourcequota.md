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

This design doc introduces how to define a group of priority class scopes for the quota to match with and explains how quota enforcement logic is changed to apply the quota to pods with the given priority classes.

## Detailed Design

### Changes in ResourceQuota

ResourceQuotaSpec contains an array of filters, `Scopes`, that if mentioned, must match each object tracked by a ResourceQuota.
```go
// ResourceQuotaSpec defines the desired hard limits to enforce for Quota
type ResourceQuotaSpec struct {
        ...

        // A collection of filters that must match each object tracked by a quota.
        // If not specified, the quota matches all objects.
        // +optional
        Scopes []ResourceQuotaScope
}
```
Four new `ResourceQuotaScope` will be defined for matching pods based on priority class names.
```go
// A ResourceQuotaScope defines a filter that must match each object tracked by a quota
type ResourceQuotaScope string

const (
        ...
        ResourceQuotaScopePriorityClassNameExists ResourceQuotaScope = "PriorityClassNameExists"
        // Match all pod objects that do not have any priority class mentioned
        ResourceQuotaScopePriorityClassNameNotExists ResourceQuotaScope = "PriorityClassNameNotExists"
        // Match all pod objects that have priority class from the set
        ResourceQuotaScopePriorityClassNameIn ResourceQuotaScope = "PriorityClassNameIn"
        // Match all pod objects that do not have priority class from the set
        ResourceQuotaScopePriorityClassNameNotIn ResourceQuotaScope = "PriorityClassNameNotIn"
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

        // MatchScopes is a collection of filters based on priority classes.
        // If the object in the intercepted request matches these rules,
        // quota system will ensure that corresponding quota MUST have
        // priority based Scopes matching the object in request. 
        //
        // If MatchScopes has matched on an object, request for the resource will be denied 
        // if there is no quota with matching Scopes. In this case, matching priority class based Scopes
        // will be an additional requirement for any quota to qualified as covering quota.
        // +optional
        MatchScopes []string `json:"matchScopes,omitempty"`
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
      - "ResourceQuotaScopePriorityClassNameIn:cluster-services"
```

2. Admin will then create a corresponding resource quota object in `kube-system` namespace:

    `$ kubectl create quota critical --hard=count/pods=10 --scopes=ResourceQuotaScopePriorityClassNameIn:cluster-services -n kube-system`

In this case, a pod creation will be allowed if:
1. Pod has no priority class and created in any namespace.
2. Pod has priority class other than `cluster-service` and created in any namespace.
3. Pod has priority class `cluster-service` and created in `kube-system` namespace, and passed resource quota check.

Pod creation will be rejected if pod has priority class `cluster-service` and created in namespace other than `kube-system`


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
      - "ResourceQuotaScopePriorityClassNameExists"
```

2. Create resource quota to match all pods where there is priority set

    `$ kubectl create quota example --hard=count/pods=10 --scopes=ResourceQuotaScopePriorityClassNameExists`

