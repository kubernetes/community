# Priority in ResourceQuota

Authors:

Harry Zhang @resouer 

Main Reviewers:

Bobby @bsalamat Derek @derekwaynecarr 

Dec 2017

  * [Objective](#objective)
    * [Non-Goals](#non-goals)
  * [Background](#background)
  * [Overview](#overview)
  * [Detailed Design](#detailed-design)
    * [Expected behavior of ResourceQuota](#resourcequota)
    * [Expected behavior of ResourceQuota admission controller](#resourcequota-admission-controller)


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

One approach to implement this is adding priority class name field to `ResourceQuota` API definition. While this arbitrary field of API object will introduce inflexibility to potential change in future and also, not adequate to express all semantics.

Thus, we decide to reuse the existing `Scopes` of `ResourceQuotaSpec` to provide a richer semantics for quota to cooperate with priority classes. 

## Overview 

This design doc introduces how to define a group of priority class scopes for the quota to match with and explains how quota enforcement logic is changed to apply the quota to pods with the given priority classes.

## Detailed Design 

### Expected behavior of ResourceQuota

Four new values of `scopes` will be defined for `ResourceQuota` to describe its relationship with priority class name:

| Scope | Description |
| ----- | ----------- |
| PriorityClassNameIn:priorityName1:priorityName2 | Match `kind=Pod` where `spec.priorityClassName in (priorityName1, priorityName2)` |
| PriorityClassNameNotIn:priorityName1:priorityName2 | Match `kind=Pod` where `spec.priorityClassName notin (priorityName1, priorityName2)` |
| PriorityClassNameExists |  Match `kind=Pod` where `spec.priorityClassName` |
| PriorityClassNameNotExists |  Match `kind=Pod` where `!spec.priorityClassName` |

*Example: Match all pods independent of priority*

`$ kubectl create quota example --hard=count/pods=10`

*Example: Match all pods where priority is foo*

`$ kubectl create quota example --hard=count/pods=10 --scopes=PriorityClassNameIn:foo`

*Example: Match all pods where there is no priority*

`$ kubectl create quota example --hard=count/pods=10 --scopes=PriorityClassNameNotExists`

To restrict usage of a particular priority without explicit quota, we can do the following by a new field named `matchScopes` introduced to `ResourceQuota` configuration.:

*Example: Any pod with priority cluster-services or other-important-stuff must require explicit quota*

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
    - "PriorityClassNameIn:cluster-services:other-important-stuff"
```

*Example: Any pod with priority set must require explicit quota*

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
    - "PriorityClassNameExists"
```

There are some details to note:

1. The absence of a quota for a particular priority class name means the quota is unbounded (unless configured otherwise by operator in admission control configuration).
2. If multiple `ResourceQuota` apply to a Pod, the pod must satisfy all of them.


### Expected behavior of ResourceQuota admission controller

* We do not enforce referential integrity across objects. i.e. Creation or updating of ResourceQuota object, scopes of which names a PriorityClass that does not exist, are allowed.
* `ResourceQuota` admission controller **should not reject** deletion of `ResourceQuota` objects with reference.

#### Changes in AdmissionConfiguration

1. `MatchScopes`, which has been explained in the section of `Expected behavior of ResourceQuota`.
2. For backward compatibility consideration, we hope it is configurable for `ResourceQuota` admission controller to reject or not to reject the pod with priority unbounded `ResourceQuota`. In order to support this scenario, we will rely on whether the value of `MatchScopes` is set as a flag.
    1. If `MatchScopes` is set, it indicates a pod with priority set but `ResourceQuota` is unbounded will be rejected by admission controller. The subsequent pods will be checked against all `ResourceQuota` and also respect this "reject when unbounded" rule.
    2. If `MatchScopes` is not set, it indicates a pod with priority set but `ResourceQuota` is unbounded will be accepted by admission. This is equivalent to the today's default behavior when `ResourceQuota` has no awareness of priority. The subsequent pods will also be checked against all `ResourceQuota` but still be accepted if no matched `ResourceQuota` found.

Those changes will be added into `LimitedResource` struct of `AdmissionConfiguration`:

```go
// LimitedResource matches a resource whose consumption is limited by default.
// To consume the resource, there must exist an associated quota that limits
// its consumption.
type LimitedResource struct {
  APIGroup string `json:"apiGroup,omitempty"`

  Resource string `json:"resource"`

  MatchContains []string

  // For each intercepted request, the quota system will evaluate
  // its resource usage. Also, it will iterate through each Kind=Pod and
  // check if its priority class fills in any critiria in this listing. User
  // can define certen types of `MatchScopes` to describe the expected priority
  // class this quota check should apply to.
  // For example,
  // PriorityClassNameIn:priorityName1:priorityName2
  // Means: this quota check apply to `kind=Pod` where `spec.priorityClassName in (priorityName1, priorityName2)`
  // Also, if MatchScopes is set, a pod with priority set but ResourceQuota is unbounded will be rejected.
  // Otherwise, it will be accepted even when ResourceQuota is unbounded.
  MatchScopes []string
}
```

*Example usage in `AdmissionConfiguration`:*
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
          defaultunbound: false
          matchContains:
          - pods
          - requests.cpu
          matchScopes:
          - PriorityClassNameIn:priorityName1:priorityName2
```

