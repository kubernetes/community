# Introducing Priority to Kubelet Memory Eviction

**Author**: David Ashpole (@dashpole)

**Last Updated**: 8/07/2017

**Status**: Proposal

This document explores various schemes to include priority in kubelet memory evictions

## Introduction

### Definitions
["Priority"](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/pod-priority-api.md) is an integer, somehow set by users when running their pods.  Assumed to be intentionally set, and controlled properly by mechanisms outside of this proposal.  
["Quality of Service"](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-qos.md), or QoS, is the performance SLO kubernetes provides based on their resource requests and limits.
 - Guaranteed: Requests == Limits  
 - Burstable: Requests < Limits  
 - Besteffort: No Requests  

["Memory Eviction"](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/kubelet-eviction.md): is the process of removing a pod from the kubelet when under memory pressure in order to free up resources.  Eviction decisions are made at the node level by the kubelet.  
"Preemption": is the process of deleting one pod from a node in order to make room for another pod, which is deemed by the scheduler to be more important to run.  Preemption decisions are made at the cluster level by the scheduler.  

### Background and Motivation
Prior to kubernetes v1.6, the [critical pod](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/rescheduling-for-critical-pods.md) annotation was introduced to prevent the permanent eviction of "critical" system pods.  For static pods, we never evict, as an evicted static pod is never re-run.  For non-static critical pods, we guarantee that they are rescheduled.  The [Kubernetes Priority](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/pod-priority-api.md) proposal introduced the concept of priority to kubernetes v1.7 as a long-term solution to ensuring that more important pods are run ahead of those that are less important.  Even though critical pods can no longer be permanently evicted, evictions can still lead to disruption in critical cluster functionality.  Integrating priority into the existing eviction algorithm can improve cluster stability by decreasing the likelyhood of evictions for critical pods in most cases.  This proposal explores possible implementations for integrating priority with the kubelet process of eviction.

The current method for ranking pods for eviction is by QoS, then usage over requests.  This currently holds the invariant that a pod that does not exceed its requests is not evicted, since the sum of the requests on the node cannot exceed the allocatable memory on the node.

If the kubelet is unable to respond to memory pressure in time, an OOM Kill may be triggered.  In this case, processes are killed based on their OOM Score.  See the [OOM Score configuration docs](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-qos.md#oom-score-configuration-at-the-nodes) for more details.

### Goals
 - Transparency.  Users should be able to understand why their pods are evicted, and what, if anything, they can do to prevent future evictions.
 - Low Abuse.  High priority pods should not be able to intentionally, or unintentionally disrupt large numbers of well-behaved pods.
 - Respect Priority.  Pods that have higher priority should be less likely to be evicted.

The goal of this proposal is to build consensus on the general design of how priority is integrated with memory eviction.  This proposal itself will not be merged, but will result in a set of changes to the [kubelet eviction documentation](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/kubelet-eviction.md) after a broad design is settled on.

### Non Goals
The implementation of priority itself is outside the scope of this proposal, and is covered in the [Priority](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/pod-priority-api.md) proposal.  This includes mechanisms to control which pods are allowed to be given which priority levels.

The scope of this design is restricted to the node, and does not make any proposals regarding cluster-level controllers.

## Proposed Implementation

### Only evict pods where usage > requests.  Then sort by Function(priority, usage - requests)
This solution provides users with a clear path to avoiding evictions. It prevents abuse by high-priority pods by not allowing them to disrupt other pods that are below, or near their requests.  Using a function provides a more nuanced approach to allowing pods to consume "unused" (not requested) memory on the node.  Power users, or cluster administrators can determine how unused memory is allocated to pods by choosing priority levels for pods that are closer for more equal sharing of extra memory, or further apart to give better availability to higher priority pods.  For pods that have equal priority, the function is equivalent to usage - requests, so that clusters that do not have priority enabled maintain behavior that is similar (though not exactly the same) as today's behavior.

## Alternatives

### By QoS, priority, then usage over requests
This solution is closest to the current behavior since it only makes a small modification by considering priority before usage over requests.  Since this is a small change from the current implementation, it should be an easy transition for cluster admins and users to make.  High-priority pods are only able to able to disrupt pods in their QoS tier or below, which lowers the potential for abuse since users can run their workloads as guaranteed if they need to avoid evictions.  However, this means that burstable pods consuming less than their requests could be evicted if a different high priority burstable pod bursts.  High priority pods are given increased availability depending on the QoS of pods that they share the node with.  If the node has many guaranteed pods, it is still possible that the high priority pod could be evicted.  If the node does not have many guaranteed pods, then the high-priority pod is able to consume all memory not consumed by Guaranteed pods.

### By priority, then usage over requests, but only evicting pods where usage > requests
This solution is similar in practice to the "By QoS, priority, then usage over requests" proposal, but preserves the current invariant that pods are guaranteed to be able to consume their requests without facing eviction.  Like the "By QoS, priority, then usage over requests" proposal, it exempts guaranteed pods from eviction, and thus lowers the potential for abuse by high priority pods.  Users can easily understand that their pods are evicted because they exceed their requests.  This solution provides additional availability for high-priority pods by giving priority access to remaining allocatable memory on the node and memory other pods request, but do not use. 

### By Priority, QoS, then usage over requests
This solution allows high-priority pods to consume up to their limits with very little chance of eviction, unless other high-priority pods are also present on the same node.  This solution would require using Quota controls on pod limits, rather than requests, as high-priority pods' requests make only a minor difference in their chance for eviction, and limits are a better indicator of what they can consume.  Evictions are easy for users to understand, as they can reason that a higher priority pod bursted, but does not provide a course of action to prevent evictions other than raising the priority of their own pod.

### Function(priority, usage over requests)
This solution specifies a mapping between priority, and usage - request to an eviction score.  For example, a possible implementation could score pods based on priority * requests / usage, or priority - (usage - requests).  It has the potential to be a solution that can balance prioritizing high-priority pods, and preventing abuse from high-priority pods.  However, it would require cluster administrators to understand this mapping in order to correctly specify priority levels, and would be prone to configuration errors.  Users would have little insight into why their pods are evicted.

## Implementation Timeline:
The integration of priority and evictions is targeted for kubernetes v1.8.  For clusters that have priority disabled, behavior will be as if all pods had equal priority.
