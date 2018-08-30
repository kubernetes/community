---
kep-number: 27
title: Resource Quota based on Node Labels
authors:
  - "@vishh"
  - "@bsalamat"
owning-sig: sig-scheduling
participating-sigs: sig-architecture
reviewers:
  - "@derekwaynecarr"
  - "@davidopp"
approvers:
  - TBD
editor: TBD
creation-date: 2018-08-23
status: provisional
---

# Resource Quota based on Node Labels

## Summary

Allowing Resource Quota to be applied on pods based on their node selector configuration opens up a flexible interface for addressing some immediate and potential future use cases.

## Motivation

As a kubernetes cluster administrator, I'd like to,

1. Restrict namespaces to specific HW types they can consume. Nodes are expected to be homogeneous wrt. to specific types of HW and HW type will be exposed as node labels.
   * A concrete example - An intern should only use the cheapest GPU available in my cluster, while researchers can consume the latest or most expensive GPUs.
2. Restrict compute resources consumed by namespaces on different zones or dedicated node pools.
3. Restrict compute resources consumed by namespaces based on policy (FIPS, HIPAA, etc) compliance on individual nodes.

This proposal presents flexible solution(s) for addressing these use cases without introducing much additional complexity to core kubernetes.

## Potential solutions

This proposal currently identifies two possible solutions, with the first one being the _preferred_ solution.

### Solution A - Extend Resource Quota Scopes

Resource Quota already includes a built in extension mechanism called [Resource Scopes](https://github.com/kubernetes/api/blob/master/core/v1/types.go#L4746).
It is possible to add a new Resource Scope called “NodeAffinityKey” (or something similar) that will allow for Resource Quota limits to apply to node selector and/or affinity fields specified in the pod spec.

Here’s an illustration of a sample object with these new fields:

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: hipaa-nodes
  namespace: team-1
spec:
  hard:
    cpu: 1000
    memory: 100Gi
  scopeSelector:
   scopeName: NodeAffinityKey
   operator: In
   values: [“hipaa-compliant: true”] 
```

``` yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: nvidia-tesla-v100-quota
  namespace: team-1
spec:
  hard:
   - nvidia.com/gpu: 128
  scopeSelector:
   scopeName: NodeAffinityKey
   operator: In
   values: [“nvidia.com/gpu-type:nvidia-tesla-v100”]
```

It is possible for quotas to overlap with this feature as is the case today.
All quotas have to be satisfied for the pod to be admitted.

[Quota configuration object](https://github.com/kubernetes/kubernetes/blob/7f23a743e8c23ac6489340bbb34fa6f1d392db9d/plugin/pkg/admission/resourcequota/apis/resourcequota/types.go#L32) will also support the new scope to allow for preventing pods from running on nodes that match a label selector unless a corresponding quota object has been created.

#### Pros

- Support arbitrary properties to be consumed as part of quota as long as they are exposed as node labels.
- Little added cognitive burden - follows existing API paradigms.
- Implementation is straightforward.
- Doesn’t compromise portability - Quota remains an administrator burden.

#### Cons

- Requires property labels to become standardized if portability is desired. This is required anyways irrespective of how they are exposed outside of the node for scheduling portability.
- Label keys and values are concatenated. Given that most selector use cases for quota will be deterministic (one -> one), the proposed API schema might be adequate.

### Solution B - Extend Resource Quota to include an explicit Node Selector field

This solution is similar to the previous one with changes to the API where instead of re-using scopes we can add an explicit Node Selector field to the Resource Quota object.

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: hipaa-nodes
  namespace: team-1
spec:
  hard:
   cpu: 1000
   memory: 100Gi
  podNodeSelector:
   matchExpressions:
    - key: hipaa-compliant
      operator: In
      values: ["true"]
```

Users should already be familiar with the Node Selector spec illustrated here as it is used in pod and volume topology specifications.
However this solution introduces a field that is only applicable to a few types of resources that Resource Quota can be used to control.

### Solution C - CRD for expressing Resource Quota for extended resources

The idea behind this solution is to let individual kubernetes vendors create additional CRDs that will allow for expressing quota per namespace for their resource and have a controller that will use mutating webhooks to quota pods on creation & deletion.
The controller can also keep track of “in use” quota for the resource it owns similar to the built in resource quota object.
The schema for quota is controlled by the resource vendor and the onus of maintaining compatibility and portability is on them.

#### Pros

- Maximum flexibility
  - Use arbitrary specifications associated with a pod to define quota policies
  - The spec for quota itself can be arbitrarily complex
- Develop and maintain outside of upstream

#### Cons

- Added administrator burden. An admin needs to identify multiple types of quota objects based on the HW they consume.
- It is not trivial to develop an external CRD given the lack of some critical validation, versioning, and lifecycle primitives.
- Tracking quota is non trivial - perhaps a canonical (example) quota controller might help ease the pain.
- Hard to generate available and in-use quota reports for users - existing quota support in ecosystem components will not support this new quota object (kubectl for example).
