---
kep-number: 1
title: Pod Ready++
authors:
  - "freehan@"
owning-sig: sig-network
participating-sigs:
  - sig-node
  - sig-cli
reviewers:
  - thockin@
  - dchen1107@
approvers:
  - thockin@
  - dchen1107@
editor: freehan@
creation-date: 2018-04-01
last-updated: 2018-04-01
status: provisional

---

# Pod Ready++


## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Alternatives](#alternatives-optional)


## Summary

This proposal aims to add extensibility to pod readiness. Besides container readiness, external feedback can be injected into PodStatus and influence pod readiness. Thus, achieving pod “ready++”. 

## Motivation

Pod readiness indicates whether the pod is ready to serve traffic. Pod readiness is dictated by kubelet with user specified readiness probe. On the other hand, pod readiness determines whether pod address shows up on the address list on related endpoints object. K8s primitives that manage pods, such as Deployment, only takes pod status into account for decision making, such as advancement during rolling update. 

For example, during deployment rolling update, a new pod becomes ready. On the other hand, service, network policy and load-balancer are not yet ready for the new pod due to whatever reason (e.g. slowness in api machinery, endpoints controller, kube-proxy, iptables or infrastructure programming). This may cause service disruption or lost of backend capacity. In extreme cases, if rolling update completes before any new replacement pod actually start serving traffic, this will cause service outage. 


### Goals

- Allow extra signals for pod readiness.

### Non-Goals

- Provide generic framework to solve all transition problems in k8s (e.g. blue green deployment).

## Proposal

[K8s Proposal: Pod Ready++](https://docs.google.com/document/d/1VFZbc_IqPf_Msd-jul7LKTmGjvQ5qRldYOFV0lGqxf8/edit#)

### PodSpec
Introduce an extra field called ReadinessGates in PodSpec. The field stores a list of ReadinessGate structure as follows: 
```yaml
type ReadinessGate struct {
	conditionType string	
}
```
The ReadinessGate struct has only one string field called ConditionType. ConditionType refers to a condition in the PodCondition list in PodStatus. And the status of conditions specified in the ReadinessGates will be evaluated for pod readiness. If the condition does not exist in the PodCondition list, its status will be default to false. 

#### Constraints:
- ReadinessGates can only be specified at pod creation. 
- No Update allowed on ReadinessGates.
- ConditionType must conform to the naming convention of custom pod condition.

### Pod Readiness
Change the pod readiness definition to as follows:
```
Pod is ready == containers are ready AND conditions in ReadinessGates are True
```
Kubelet will evaluate conditions specified in ReadinessGates and update the pod “Ready” status. For example, in the following pod spec, two readinessGates are specified. The status of “www.example.com/feature-1” is false, hence the pod is not ready. 

```yaml
Kind: Pod 
… 
spec: 
  readinessGates:
  - conditionType: www.example.com/feature-1
  - conditionType: www.example.com/feature-2
… 
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2018-01-01T00:00:00Z
    status: "False"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2018-01-01T00:00:00Z
    status: "False"
    type: www.example.com/feature-1
  - lastProbeTime: null
    lastTransitionTime: 2018-01-01T00:00:00Z
    status: "True"
    type: www.example.com/feature-2
  containerStatuses:
  - containerID: docker://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    ready : true
… 
```

Another pod condition `ContainerReady` will be introduced to capture the old pod `Ready` condition. 
```
ContainerReady is true == containers are ready
```

### Custom Pod Condition
Custom pod condition can be injected thru PATCH action using KubeClient. Please be noted that “kubectl patch” does not support patching object status. Need to use client-go or other KubeClient implementations. 

Naming Convention:
The type of custom pod condition must comply with k8s label key format. For example, “www.example.com/feature-1”.


### Implementation Details/Notes/Constraints

##### Workloads
To conform with this proposals, workload controllers MUST take pod “Ready” condition as the final signal to proceed during transitions.  

For the workloads that take pod readiness as a critical signal for its decision making, they will automatically comply with this proposal without any change. Majority, if not all, of the workloads satisfy this condition. 

##### Kubelet
- Use PATCH instead of PUT to update PodStatus fields that are dictated by kubelet. 
- Only compare the fields that managed by kubelet for PodStatus reconciliation .
- Watch PodStatus changes and evaluate ReadinessGates for pod readiness.

### Feature Integration
In this section, we will discuss how to make ReadinessGates transparent to K8s API user. In order words, a K8s API user does not need to specify ReadinessGates to use specific features. This allows existing manifests to just work with features that require ReadinessGate.
Each feature will bear the burden of injecting ReadinessGate and keep its custom pod condition in sync. ReadinessGate can be injected using mutating webhook at pod creation time. After pod creation, each feature is responsible for keeping its custom pod condition in sync as long as its ReadinessGate exists in the PodSpec. This can be achieved by running k8s controller to sync conditions on relevant pods. This is to ensure that PodStatus is observable and recoverable even when catastrophic failure (e.g. loss of data) occurs at API server. 



### Risks and Mitigations

Risks:
- Features that utilize the extension point from this proposal may abuse the API.  
- User confusion on pod ready++

Mitigations:
- Better specification and API validation.
- Better CLI/UI/UX
  

## Graduation Criteria

- Kubelet changes should not have any impact on kubelet reliability. 
- Feature integration with the pod ready++ extension.


## Implementation History

TBD


## Alternatives 

##### Why not fix the workloads?

There are a lot of workloads including core workloads such as deployment and 3rd party workloads such as spark operator. Most if not all of them take pod readiness as a critical signal for decision making, while ignoring higher level abstractions (e.g. service, network policy and ingress). To complicate the problem more, label selector makes membership relationship implicit and dynamic. Solving this problem in all workload controllers would require much bigger change than this proposal. 

##### Why not extend container readiness?

Container readiness is tied to low level constructs such as runtime. This inherently implies that the kubelet and underlying system has full knowledge of container status. Injecting external feedback into container status would complicate the abstraction and control flow. Meanwhile, higher level abstractions (e.g. service) generally takes pod as the atom instead of container. 
