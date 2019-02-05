# Deployment

Authors:
- Brian Grant (@bgrant0607)
- Clayton Coleman (@smarterclayton)
- Dan Mace (@ironcladlou)
- David Oppenheimer (@davidopp)
- Janet Kuo (@janetkuo)
- Michail Kargakis (@kargakis)
- Nikhil Jindal (@nikhiljindal)

## Abstract

A proposal for implementing a new resource - Deployment - which will enable
declarative config updates for ReplicaSets. Users will be able to create a
Deployment, which will spin up a ReplicaSet to bring up the desired Pods.
Users can also target the Deployment to an existing ReplicaSet either by
rolling back an existing Deployment or creating a new Deployment that can
adopt an existing ReplicaSet. The exact mechanics of replacement depends on
the DeploymentStrategy chosen by the user. DeploymentStrategies are explained
in detail in a later section.

## Implementation

### API Object

The `Deployment` API object will have the following structure:

```go
type Deployment struct {
  TypeMeta
  ObjectMeta

  // Specification of the desired behavior of the Deployment.
  Spec DeploymentSpec

  // Most recently observed status of the Deployment.
  Status DeploymentStatus
}

type DeploymentSpec struct {
  // Number of desired pods. This is a pointer to distinguish between explicit
  // zero and not specified. Defaults to 1.
  Replicas *int32

  // Label selector for pods. Existing ReplicaSets whose pods are
  // selected by this will be scaled down. New ReplicaSets will be
  // created with this selector, with a unique label `pod-template-hash`.
  // If Selector is empty, it is defaulted to the labels present on the Pod template.
  Selector map[string]string

  // Describes the pods that will be created.
  Template *PodTemplateSpec

  // The deployment strategy to use to replace existing pods with new ones.
  Strategy DeploymentStrategy

  // Minimum number of seconds for which a newly created pod should be ready
  // without any of its container crashing, for it to be considered available.
  // Defaults to 0 (pod will be considered available as soon as it is ready)
  MinReadySeconds int32
}

type DeploymentStrategy struct {
  // Type of deployment. Can be "Recreate" or "RollingUpdate".
  Type DeploymentStrategyType

  // Rolling update config params. Present only if DeploymentStrategyType =
  // RollingUpdate.
  RollingUpdate *RollingUpdateDeploymentStrategy
}

type DeploymentStrategyType string

const (
  // Kill all existing pods before creating new ones.
  RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

  // Replace the old ReplicaSets by new one using rolling update i.e gradually scale
  // down the old ReplicaSets and scale up the new one.
  RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

// Spec to control the desired behavior of rolling update.
type RollingUpdateDeploymentStrategy struct {
  // The maximum number of pods that can be unavailable during the update.
  // Value can be an absolute number (ex: 5) or a percentage of total pods at the start of update (ex: 10%).
  // Absolute number is calculated from percentage by rounding up.
  // This can not be 0 if MaxSurge is 0.
  // By default, a fixed value of 1 is used.
  // Example: when this is set to 30%, the old RC can be scaled down by 30%
  // immediately when the rolling update starts. Once new pods are ready, old RC
  // can be scaled down further, followed by scaling up the new RC, ensuring
  // that at least 70% of original number of pods are available at all times
  // during the update.
  MaxUnavailable IntOrString

  // The maximum number of pods that can be scheduled above the original number of
  // pods.
  // Value can be an absolute number (ex: 5) or a percentage of total pods at
  // the start of the update (ex: 10%). This can not be 0 if MaxUnavailable is 0.
  // Absolute number is calculated from percentage by rounding up.
  // By default, a value of 1 is used.
  // Example: when this is set to 30%, the new RC can be scaled up by 30%
  // immediately when the rolling update starts. Once old pods have been killed,
  // new RC can be scaled up further, ensuring that total number of pods running
  // at any time during the update is atmost 130% of original pods.
  MaxSurge IntOrString
}

type DeploymentStatus struct {
  // Total number of ready pods targeted by this deployment (this
  // includes both the old and new pods).
  Replicas int32

  // Total number of new ready pods with the desired template spec.
  UpdatedReplicas int32

  // Count of hash collisions for the Deployment. The Deployment controller uses this
  // field as a collision avoidance mechanism when it needs to create the name for the
  // newest ReplicaSet.
  CollisionCount *int64
}

```

### Controller

#### Deployment Controller

The DeploymentController will process Deployments and crud ReplicaSets.
For each creation or update for a Deployment, it will:

1. Find all RSs (ReplicaSets) whose label selector is a superset of DeploymentSpec.Selector.
   - For now, we will do this in the client - list all RSs and then filter out the
     ones we want. Eventually, we want to expose this in the API.
2. The new RS can have the same selector as the old RS and hence we add a unique
   selector to all these RSs (and the corresponding label to their pods) to ensure
   that they do not select the newly created pods (or old pods get selected by the
   new RS).
   - The label key will be "pod-template-hash".
   - The label value will be the hash of {podTemplateSpec+collisionCount} where podTemplateSpec
     is the one that the new RS uses and collisionCount is a counter in the DeploymentStatus
     that increments every time a [hash collision](#hashing-collisions) happens (hash
     collisions should be rare with fnv).
   - If the RSs and pods don't already have this label and selector:
     - We will first add this to RS.PodTemplateSpec.Metadata.Labels for all RSs to
       ensure that all new pods that they create will have this label.
     - Then we will add this label to their existing pods
     - Eventually we flip the RS selector to use the new label.
     This process potentially can be abstracted to a new endpoint for controllers [1].
3. Find if there exists an RS for which value of "pod-template-hash" label
   is same as hash of DeploymentSpec.PodTemplateSpec. If it exists already, then
   this is the RS that will be ramped up. If there is no such RS, then we create
   a new one using DeploymentSpec and then add a "pod-template-hash" label
   to it. The size of the new RS depends on the used DeploymentStrategyType.
4. Scale up the new RS and scale down the olds ones as per the DeploymentStrategy.
   Raise events appropriately (both in case of failure or success).
5. Go back to step 1 unless the new RS has been ramped up to desired replicas
   and the old RSs have been ramped down to 0.
6. Cleanup old RSs as per revisionHistoryLimit.

DeploymentController is stateless so that it can recover in case it crashes during a deployment.

[1] See https://github.com/kubernetes/kubernetes/issues/36897

### MinReadySeconds

We will implement MinReadySeconds using the Ready condition in Pod. We will add
a LastTransitionTime to PodCondition and update kubelet to set Ready to false,
each time any container crashes. Kubelet will set Ready condition back to true once
all containers are ready. For containers without a readiness probe, we will
assume that they are ready as soon as they are up.
https://github.com/kubernetes/kubernetes/issues/11234 tracks updating kubelet
and https://github.com/kubernetes/kubernetes/issues/12615 tracks adding
LastTransitionTime to PodCondition.

## Changing Deployment mid-way

### Updating

Users can update an ongoing Deployment before it is completed.
In this case, the existing rollout will be stalled and the new one will
begin.
For example, consider the following case:
- User updates a Deployment to rolling-update 10 pods with image:v1 to
  pods with image:v2.
- User then updates this Deployment to create pods with image:v3,
  when the image:v2 RS had been ramped up to 5 pods and the image:v1 RS
  had been ramped down to 5 pods.
- When Deployment Controller observes the new update, it will create
  a new RS for creating pods with image:v3. It will then start ramping up this
  new RS to 10 pods and will ramp down both the existing RSs to 0.

### Deleting

Users can pause/cancel a rollout by doing a non-cascading deletion of the Deployment
before it is complete. Recreating the same Deployment will resume it.
For example, consider the following case:
- User creates a Deployment to perform a rolling-update for 10 pods from image:v1 to
 image:v2.
- User then deletes the Deployment while the old and new RSs are at 5 replicas each.
  User will end up with 2 RSs with 5 replicas each.
User can then re-create the same Deployment again in which case, DeploymentController will
notice that the second RS exists already which it can ramp up while ramping down
the first one.

### Rollback

We want to allow the user to rollback a Deployment. To rollback a completed (or
ongoing) Deployment, users can simply use `kubectl rollout undo` or update the
Deployment directly by using its spec.rollbackTo.revision field and specify the
revision they want to rollback to or no revision which means that the Deployment
will be rolled back to its previous revision.

## Deployment Strategies

DeploymentStrategy specifies how the new RS should replace existing RSs.
To begin with, we will support 2 types of Deployment:
* Recreate: We kill all existing RSs and then bring up the new one. This results
  in quick Deployment but there is a downtime when old pods are down but
  the new ones have not come up yet.
* Rolling update: We gradually scale down old RSs while scaling up the new one.
  This results in a slower Deployment, but there can be no downtime. Depending on
  the strategy parameters, it is possible to have at all times during the rollout
  available pods (old or new). The number of available pods and when is a pod
  considered "available" can be configured using RollingUpdateDeploymentStrategy.

## Hashing collisions

Hashing collisions are a real thing with the existing hashing algorithm[1]. We
need to switch to a more stable algorithm like fnv. Preliminary benchmarks[2]
show that while fnv is a bit slower than adler, it is much more stable. Also,
hashing an API object is subject to API changes which means that the name
for a ReplicaSet may differ between minor Kubernetes versions.

For both of the aforementioned cases, we will use a field in the DeploymentStatus,
called collisionCount, to create a unique hash value when a hash collision happens.
The Deployment controller will compute the hash value of {template+collisionCount},
and will use the resulting hash in the ReplicaSet names and selectors. One side
effect of this hash collision avoidance mechanism is that we don't need to
migrate ReplicaSets that were created with adler.

[1] https://github.com/kubernetes/kubernetes/issues/29735

[2] https://github.com/kubernetes/kubernetes/pull/39527

## Future

Apart from the above, we want to add support for the following:
* Running the deployment process in a pod: In future, we can run the deployment process in a pod. Then users can define their own custom deployments and we can run it using the image name.
* More DeploymentStrategyTypes: https://github.com/openshift/origin/blob/master/examples/deployment/README.md#deployment-types lists most commonly used ones.
* Triggers: Deployment will have a trigger field to identify what triggered the deployment. Options are: Manual/UserTriggered, Autoscaler, NewImage.
* Automatic rollback on error: We want to support automatic rollback on error or timeout.

## References

- https://github.com/kubernetes/kubernetes/issues/1743 has most of the
  discussion that resulted in this proposal.
