# Equivalence class based scheduling in Kubernetes

**Authors**:

@resouer @wojtek-t @davidopp


# Guideline

- [Objectives](#objectives)
   - [Goals](#goals)
   - [Non-Goals](#non-goals)
- [Background](#background)
   - [Terminology](#terminology)
- [Overview](#overview)
- [Detailed Design](#detailed-design)
   - [Define equivalence class](#define-equivalence-class)
   - [Equivalence class in predicate phase](#equivalence-class-in-predicate)
   - [Keep equivalence class cache up-to-date](#keep-equivalence--class-cache-up-to-date)
- [Notes for scheduler developers](#notes-for-scheduler-developer)
- [References](#references)

# Objectives

## Goals

-  Define the equivalence class for pods during predicate phase in Kubernetes.
-  Define how to use equivalence class to speed up predicate process.
-  Define how to ensure information cached in equivalence class is up-to-date.

## Non-Goals

-  Apply equivalence class to priorities. We have refactored priorities to a Map-Reduce style process, we need to re-evaluate whether equivalence design can or can not apply to this new model.

# Background

Pods in Kubernetes cluster usually have identical requirements and constraints, just think about a Deployment with a number of replications. So rather than determining feasibility for every pending pod on every node, we can only do predicates one pod per equivalence class – a group of tasks with identical requirements, and reuse the predicate results for other equivalent pods. 

We hope to use this mechanism to help to improve scheduler's scalability, especially in cases like Replication Controller with huge number of instances, or eliminate pressure caused by complex predicate functions.

The concept of equivalence class in scheduling is a proven feature used originally in [Google Borg] [1].

## Terminology

Equivalence class: a group of pods which has identical requirements and constraints. 

Equivalence class based scheduling: the scheduler will do predicate for only one pod per equivalence class, and reuse this result for all other equivalent pods.

# Overview

This document describes what is equivalence class, and how to do equivalence based scheduling in Kubernetes. The basic idea is when you apply the predicate functions to a pod, cache the results (namely, for each machine, whether the pod is feasible on that machine).

Scheduler watches for API objects change like bindings and unbindings and node changes, and marks a cached value as invalid whenever there is a change that invalidates a cached value. (For example, if the labels on a node change, or a new pod gets bound to a machine, then all cached values related to that machine are invalidated.) In the future when we have in-place updates, some updates to pods running on the machine would also cause the node to be marked invalid. This is how we keep equivalence class cache up-to-date.

When scheduling a new pod, check to see if the predicate result for an equivalent pod is already cached. If so, re-evaluate the predicate functions just for the "invalid" values (i.e. not for all nodes and predicates), and update the cache.


# Detailed Design

## 1. Define equivalence class

There are two options were proposed.

Option 1: use the attributes of Pod API object to decide if given pods are equivalent, the attributes include labels, some annotations, affinity, resource limit etc.

Option 2: use controller reference, i.e. simply consider pods belonging to same controller reference
to be equivalent.

Regarding first option - The biggest concern in this approach is that if someone will add dependency on some new field at some point, we don't have good way to test it and ensure that equivalence pod will be updated at that point too.

Regarding second option - In detail, using the "ControllerRef" which is defined as "OwnerReference (from ObjectMeta) with the "Controller" field set to true as the "equivalence class". In this approach, we would have all RC, RS, Job etc handled by exactly the same mechanism. Also, this would be faster to compute it. 

For example, two pods created by the same `ReplicaSets` will be considered as equivalent since they will have exactly the same resource requirements from one pod template. On the other hand, two pods created by two `ReplicaSets` will not be considered as equivalent regardless of whether they have same resource requirements or not.

**Conclusion:**

Choose option 2. And we will calculate a unique `uint64` hash for pods belonging to same equivalence class which known as `equivalenceHash`.

## 2. Equivalence class in predicate phase

Predicate is the first phase in scheduler to filter out nodes which are feasible to run the workload. In detail:

1. Predicates functions are registered in scheduler
2. The predicates will be checked by `scheduler.findNodesThatFit(pod, nodes, predicateFuncs ...)`. 
3. The check process `scheduler.podFitsOnNode(pod, node, predicateFuncs ...)` is executed in parallel for every node. 

### 2.1 Design an equivalence class cache

The step 3 is where registered predicate functions will be called against given pod and node. This step includes:

1. Check if given pod has equivalence class.
2. If yes, use equivalence class cache to do predicate.

In detail, we need to have an equivalence class cache to store all predicates results per node. The data structure is a 3 level map with keys of the levels being: `nodeName`, `predicateKey` and `equivalenceHash`.

```go
predicateMap := algorithmCache[nodeName].predicatesCache.Get(predicateKey)
hostPredicate := predicateMap[equivalenceHash]
```
For example: the cached `GeneralPredicates` result for equivalence class `1000392826` on node `node_1` is:

```go
algorithmCache["node_1"].predicatesCache.Get("GeneralPredicates")[1000392826]
```

This will return a `HostPredicate` struct:

```go
type HostPredicate struct {
   Fit         bool
   FailReasons []algorithm.PredicateFailureReason
}

```

Please note we use predicate name as key in `predicatesCache`, so the number of entries in the cache is less or equal to the total number of registered predicates in scheduler. The cache size is limited.

### 2.2 Use cached predicate result to do predicate

The pseudo code of predicate process with equivalence class will be like:

```go
func (ec *EquivalenceCache) PredicateWithECache(
   podName, nodeName, predicateKey string,
   equivalenceHash uint64,
) (bool, []algorithm.PredicateFailureReason, bool) {
   if algorithmCache, exist := ec.algorithmCache[nodeName]; exist {
      if predicateMap, exist := algorithmCache.predicatesCache.Get(predicateKey); exist {
         if hostPredicate, ok := predicateMap[equivalenceHash]; ok {
            // fit
            if hostPredicate.Fit {
               return true, []algorithm.PredicateFailureReason{}, false
            } else {
               // unfit
               return false, hostPredicate.FailReasons, false
            }
         } else {
            // cached result is invalid
            return false, []algorithm.PredicateFailureReason{}, true
         }
      }
   }
   return false, []algorithm.PredicateFailureReason{}, true
}
```

One thing to note is, if the `hostPredicate` is not present in the logic above, it will be considered as `invalid`. That means although this pod has equivalence class, it does not have cached predicate result yet, or the cached data is not valid. It needs to go through normal predicate process and write the result into equivalence class cache.

### 2.3 What if no equivalence class is found for pod?

If no equivalence class is found for given pod, normal predicate process will be executed.

## 3. Keep equivalence class cache up-to-date

The key of this equivalence class based scheduling is how to keep the equivalence cache up-to-date. Since even one single pod been scheduled to a node will make the cached result not stand as the available resource on this node has changed.

One approach is that we can invalidate the cached predicate result for this node. But in a heavy load cluster state change happens frequently and makes the design less meaningful.

So in this design, we proposed the ability to invalidate cached result for specific predicate. For example, when a new pod is scheduled to a node, the cached result for `PodFitsResources` should be invalidated on this node while others can still be re-used. That's also another reason we use predicate name as key for the cached value.

During the implementation, we need to consider all the cases which may affect the effectiveness of cached predicate result. The logic includes three dimensions:

- **Operation**: 
    - what operation will cause this cache invalid.
- **Invalid predicates**: 
    - what predicate should be invalidated.
- **Scope**: 
    - the cache of which node should be invalidated, or all nodes.

Please note with the change of predicates in subsequent development, this doc will become out-of-date, while you can always check the latest e-class cache update process in `pkg/scheduler/factory/factory.go`.

### 3.1 Persistent Volume

- **Operation:**
    - ADD, DELETE

- **Invalid predicates**:

    - `MaxEBSVolumeCount`, `MaxGCEPDVolumeCount`, `MaxAzureDiskVolumeCount` (only if the added/deleted PV is one of them)

- **Scope**:

    - All nodes (we don't know which node this PV will be attached to)


### 3.2 Persistent Volume Claim

- **Operation:**
    - ADD, DELETE

- **Invalid predicates:**

    - `MaxPDVolumeCountPredicate` (only if the added/deleted PVC as a bound volume so it drops to the PV change case, otherwise it should not affect scheduler).

- **Scope:**
    - All nodes (we don't know which node this PV will be attached to).


### 3.3 Service

- **Operation:**
    - ADD, DELETE

- **Invalid predicates:** 

    - `ServiceAffinity`

- **Scope:**
    - All nodes (`serviceAffinity` is a cluster scope predicate).



- **Operation:**
    - UPDATE

- **Invalid predicates:**

    - `ServiceAffinity` (only if the `spec.Selector` filed is updated)

- **Scope:**
    - All nodes (`serviceAffinity` is a cluster scope predicate),.


### 3.4 Pod

- **Operation:**
    - ADD

- **Invalid predicates:**
    - `GeneralPredicates`. This invalidate should be done during `scheduler.assume(...)` because binding can be asynchronous. So we just optimistically invalidate predicate cached result there, and if later this pod failed to bind, the following pods will go through normal predicate functions and nothing breaks.

    - No `MatchInterPodAffinity`: the scheduler will make sure newly bound pod will not break the existing inter pod affinity. So we do not need to invalidate MatchInterPodAffinity when pod added. But when a pod is deleted, existing inter pod affinity may become invalid. (e.g. this pod was preferred by some else, or vice versa).

        - NOTE: assumptions above **will not** stand when we implemented features like `RequiredDuringSchedulingRequiredDuringExecution`.

    - No `NoDiskConflict`: the newly scheduled pod fits to existing pods on this node, it will also fits to equivalence class of existing pods.

- **Scope:** 
    - The node where the pod is bound.



- **Operation:**
    - UPDATE

- **Invalid predicates:**

    - Only if `pod.NodeName` did not change (otherwise it drops to add/delete case)

    - `GeneralPredicates` if the pod's resource requests are updated.

    - `MatchInterPodAffinity` if the pod's labels are updated.

- **Scope:**
    - The node where the pod is bound.



- **Operation:**
    - DELETE

- **Invalid predicates:**
    - `MatchInterPodAffinity` if the pod's labels are updated.

- **Scope:**
    - All nodes in the same failure domain

- **Invalid predicates:**

    - `NoDiskConflict` if the pod has special volume like `RBD`, `ISCSI`, `GCEPersistentDisk` etc.

- **Scope:**
    - The node where the pod is bound.


### 3.5 Node


- **Operation:**
    - UPDATE

- **Invalid predicates:**

    - `GeneralPredicates`, if `node.Status.Allocatable` or node labels changed.

    - `ServiceAffinity`, if node labels changed, since selector result may change.

    - `MatchInterPodAffinity`, if value of label changed, since any node label can be topology key of pod.

    - `NoVolumeZoneConflict`, if zone related label change.

    - `PodToleratesNodeTaints`, if node taints changed.

    - `CheckNodeMemoryPressure`, `CheckNodeDiskPressure`, `CheckNodeCondition`, if related node condition changed.

- **Scope:**
    - The updated node.

- **Operation:**
    - DELETE

- **Invalid predicates:**
    - All predicates

- **Scope:**
    - The deleted node


# Notes for scheduler developers

1. When implementing a new predicate, developers are expected to check how related API object changes (add/delete/update) affect the result of their new predicate function and invalidate cached results of the predicate function if necessary, in scheduler/factory/factory.go.

2. When updating an existing predicate, developers should consider whether their changes introduce new dependency on attributes of any API objects like Pod, Node, Service, etc. If so, developer should consider invalidating caches results of this predicate in scheduler/factory/factory.go.


# References

Main implementation PRs: 

- https://github.com/kubernetes/kubernetes/pull/31605
- https://github.com/kubernetes/kubernetes/pull/34685
- https://github.com/kubernetes/kubernetes/pull/36238
- https://github.com/kubernetes/kubernetes/pull/41541


[1]: http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/43438.pdf "Google Borg paper"
