# predicates ordering



Status: proposal

Author: yastij
Approvers: 
* gmarek
* bsalamat
* k82cn




## Abstract

This document describes how and why reordering predicates helps to achieve performance for the kubernetes scheduler.
We will expose the motivations behind this proposal, The two steps/solution we see to tackle this problem and the timeline decided to implement these.


## Motivation

While working on a [Pull request](https://github.com/kubernetes/kubernetes/pull/50185) related to a proposal, we saw that the order of running predicates isn’t defined. 

This makes the scheduler perform extra-computation that isn’t needed, As an example we [outlined](https://github.com/kubernetes/kubernetes/pull/50185) that the kubernetes scheduler runs predicates against nodes even if marked “unschedulable”.

Reordering predicates allows us to avoid this problem, by computing the most restrictive predicates first. To do so, we propose two reordering types.



## Static ordering

This ordering will be the default ordering. If a policy config is provided with a subset of predicates, only those predicates will be invoked using the static ordering. 




|Position                  | Predicate                        | comments (note, justification...)              |
 ----------------- | ---------------------------- | ------------------
| 1 | `CheckNodeConditionPredicate`  | we really don’t want to check predicates against unschedulable nodes. |
| 2           | `PodFitsHost`            | we check the pod.spec.nodeName. |
| 3           | `PodFitsHostPorts` | we check ports asked on the spec. |
| 4 | `PodMatchNodeSelector`            | check node label after narrowing search. |
| 5           | `PodFitsResources `            | this one comes here since it’s not restrictive enough as we do not try to match values but ranges. |
| 6           | `NoDiskConflict` | Following the resource predicate, we check disk |
| 7 | `PodToleratesNodeTaints '`            | check toleration here, as node might have toleration |
| 8          | `PodToleratesNodeNoExecuteTaints`            | check toleration here, as node might have toleration |
| 9           | `CheckNodeLabelPresence ` | labels are easy to check, so this one goes before |
| 10 | `checkServiceAffinity `            | - |
| 11           | `MaxPDVolumeCountPredicate `            | - |
| 12           | `VolumeNodePredicate ` | - |
| 13 | `VolumeZonePredicate `            | - |
| 14           | `CheckNodeMemoryPressurePredicate`            | doesn’t happen often |
| 15           | `CheckNodeDiskPressurePredicate` | doesn’t happen often |
| 16 | `InterPodAffinityMatches`            | Most expensive predicate to compute |


## End-user ordering

Using scheduling policy file, the cluster admin can override the default static ordering. This gives administrator the maximum flexibility regarding scheduler behaviour and enables scheduler to adapt to cluster usage. 
Please note that the order must be a positive integer, also, when providing equal ordering for many predicates, scheduler will determine the order and won't guarantee that the order will remain the same between them.
Finally updating the scheduling policy file will require a scheduler restart.

as an example the following is scheduler policy file using an end-user ordering:

``` json
{
"kind" : "Policy",
"apiVersion" : "v1",
"predicates" : [
	{"name" : "PodFitsHostPorts", "order": 2},
	{"name" : "PodFitsResources", "order": 3},
	{"name" : "NoDiskConflict", "order": 5},
	{"name" : "PodToleratesNodeTaints", "order": 4},
	{"name" : "MatchNodeSelector", "order": 6},
	{"name" : "PodFitsHost", "order": 1}
	],
"priorities" : [
	{"name" : "LeastRequestedPriority", "weight" : 1},
	{"name" : "BalancedResourceAllocation", "weight" : 1},
	{"name" : "ServiceSpreadingPriority", "weight" : 1},
	{"name" : "EqualPriority", "weight" : 1}
	],
"hardPodAffinitySymmetricWeight" : 10
}
```


## Timeline

* static ordering: GA in 1.9
* dynamic ordering: TBD based on customer feedback
