---
kep-number: 12
title: Support deployment controller in-place update
authors:
  - "@jian-he"
owning-sig: sig-apps
reviewers:
  - @janetkuo, @kargakis
approvers:
  - TBD
editor: "@jian-he"
creation-date: "2018-05-23"
last-updated: "2018-05-31"
status: provisional
---

# Support deployment controller in-place update

Table of Contents
=================

   * [Support deployment controller in-place update](#support-deployment-controller-in-place-update)
   * [Table of Contents](#table-of-contents)
      * [Summary](#summary)
      * [Motivation](#motivation)
      * [Goals](#goals)
      * [Proposal](#proposal)
         * [API Change](#api-change)
         * [Deployment Controller](#deployment-controller)
            * [Update](#update)
            * [Rollback](#rollback)
         * [Alternatives](#alternatives)
            
## Summary
We propose a new update strategy for deployment controller to do rolling update without tearing down the pods, i.e. in-place update.


## Motivation
Today, Deployment controller, while doing rolling update, will teardown all the pods and then recreate all the pods.
This approach works in most cases but has several limitations:
* The re-created pod may be allocated onto different nodes, but certain pods depend on local state on the node. 

    For example: some pods have persistent state in local storage. 
    Especially when we do high pressure with very large scale testing, we may go through several rounds of upgrades. If the pods randomly run on any nodes, 
    the network environment is not consistent, it becomes hard to guarantee a stable result.
* The new pod may not be allocated due to resources taken by other pods. Preemption can mitigate this situation, but this 
introduces an extra scheduling loop that can be avoided.
* The IP address may be changed due to the pod is re-created, but we want fixed IP across container upgrades. 

    It is true that pods in Kubernetes don't have sticky IPs by default. In the scenario where the node fails, 
    we implemented a CNI plugin to ensure that the new pod allocated still get the same ip address.


## Goals
Enable deployment controller to update the containers inside the pods (such as container image, cmd etc.) without tearing down the pods

## Proposal

### API Change
Augment deployment controller to have a new update strategy type: InPlaceUpdate.

```

type DeploymentStrategy struct {
	// Type of deployment. Can be "Recreate" or "RollingUpdate" or "InPlaceUpdate". Default is RollingUpdate.
	// +optional
	Type DeploymentStrategyType

	// Rolling update config params. Present only if DeploymentStrategyType =
	// RollingUpdate.
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be.
	// +optional
	RollingUpdate *RollingUpdateDeployment

+	// InPlace update config params. Present only if DeploymentStrategyType =
+	// InPlaceUpdate.
+	// +optional
+	InPlaceUpdate *InPlaceUpdateDeployment
}

const(
	// Kill all existing pods before creating new ones.
	RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

        // Replace the old RCs by new one using rolling update i.e gradually scale down the old RCs and scale up the new one.
        RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
    
        // Update containers inside the pods without tearing down the pods
+	InPlaceUpdateDeploymentStrategyType DeploymentStrategyType = "InPlaceUpdate"
)

// Spec to control the desired behavior of in-place update
+ type InPlaceUpdateDeployment struct {
+	 // The number of pods to be updated in each batch.
+	 // Value can be an absolute number (ex: 5) or a percentage of total pods at the start of update (ex: 10%).
+	 // Absolute number is calculated from percentage by rounding down.
+	 BatchSize intstr.IntOrString
+ }
```

User specifies the new update strategy type in the spec and issue the request

### Deployment Controller

#### Update
Below is the proposed approach for in-place update.
Wait until expected number of pods created, when deployment controller receives the in-place update request

```
    newReplicaSet created with 0 replica
    ...

    // Find all old pods and update the unhealthy(failed, pending, not ready) pods first with the new spec from new ReplicaSet and then the running pods,
    // BatchSize is the number of pods to update at each round.
    podsToUpdate = getPodsToUpdate(oldPods, BatchSize)
    for pod range(podsToUpdate)
        update(pod, newSpec)

    // swap oldReplicaSet and newReplicaSet information such as spec, revision, annotation etc.
    copyInfo = oldReplicaSet // write into etcd, in case the controller failed that may cause oldReplicaSet lost
    oldReplicaSet = newReplicaSet
    newReplicaSet = copyInfo
```
In essence, this approach replaces the current ReplicaSet with the new ReplicaSet info. The pods are always managed by 
the same ReplicaSet object, but just its content gets swapped.

We propose not updating the creationTimeStamp of the ReplicaSet, since pods are not re-created and creation-timestamp also stays the same.

#### Rollback
Rollback will be supported in a similar fashion as the existing rollback strategy. It first checkout the old ReplicaSet and then do update.
To do the update,  it'll use the same steps as above to update the pods in-place, instead of deleting old pods and creating new pods.

### Alternatives

```
    newReplicaSet created with 0 replica
    ...

    // Find all old pods and update the unhealthy(failed, pending, not ready) pods first with the new spec from new ReplicaSet and then the running pods,
    podsToUpdate = getPodsToUpdate(oldPods, BatchSize)
    for pod range(podsToUpdate)
        update(pod, newSpec)
    
    // Pause the old replica set. This is required for this scenario: if we reduce oldReplicaSet immediately, there's a chance
    // that oldReplicaSet may kill the pods prematurely.
    pause(oldReplicaSet)
    
    // Reduce oldReplicaSet by the BatchSize
    scaleDown(oldReplicaSet, replicas - BatchSize)
    
    // update pod controllerRef to the new ReplicaSet
    for pod range(podsToUpdate)
        update(pod, newReplicaSet-controllerRef)

    // Increase newReplicaSet by BatchSize
    scaleUp(newReplicaSet, replicas + BatchSize)

```
This approach requires some changes in ReplicaSet controller to support pause semantics. Another way is to pass certain information into ReplicaSet and prevent it from prematurely killing
the pods while updating the pods' controllerRef and ReplicaSet's replicas. Or make Kubernetes support transactions which is another big effort.
