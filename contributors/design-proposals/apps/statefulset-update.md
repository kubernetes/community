# StatefulSet Updates

**Author**: kow3ns@

**Status**: Proposal

## Abstract
Currently (as of Kubernetes 1.6), `.Spec.Replicas` and 
`.Spec.Template.Containers` are the only mutable fields of the 
StatefulSet API object. Updating `.Spec.Replicas` will scale the number of Pods 
in the StatefulSet. Updating `.Spec.Template.Containers` causes all subsequently 
created Pods to have the specified containers. In order to cause the 
StatefulSet controller to apply its updated `.Spec`, users must manually delete 
each Pod. This manual method of applying updates is error prone. The 
implementation of this proposal will add the capability to perform ordered,
automated, sequential updates. 

## Affected Components
1. API Server
1. Kubectl
1. StatefulSet Controller
1. StatefulSetSpec API object
1. StatefulSetStatus API object

## Use Cases
Upon implementation, this design will support the following in scope use cases, 
and it will not rule out the future implementation of the out of scope use 
cases.

### In Scope
- As the administrator of a stateful application, in order to vertically scale 
my application, I want to update resource limits or requested resources.
- As the administrator of a stateful application, in order to deploy critical 
security updates, break fix patches, and feature releases, I want to update 
container images.
- As the administrator of a stateful application, in order to update my 
application's configuration, I want to update environment variables, container 
entry point commands or parameters, or configuration files.
- As the administrator of the logging and monitoring infrastructure for my 
organization, in order to add logging and monitoring side cars, I want to patch
a Pods' containers to add images.

### Out of Scope
- As the administrator of a stateful application, in order to increase the 
applications storage capacity, I want to update PersistentVolumes.
- As the administrator of a stateful application, in order to update the 
network configuration of the application, I want to update Services and 
container ports in a consistent way.
- As the administrator of a stateful application, when I scale my application 
horizontally, I want associated PodDisruptionBudgets to be adjusted to 
compensate for the application's scaling.

## Assumptions
 - StatefulSet update must support singleton StatefulSets. However, an update in
 this case will cause a temporary outage. This is acceptable as a single 
 process application is, by definition, not highly available.
 - Disruption in Kubernetes is controlled by PodDisruptionBudgets. As 
 StatefulSet updates progress one Pod at a time, and only occur when all 
 other Pods have a Status of Running and a Ready Condition, they can not 
 violate reasonable PodDisruptionBudgets.
 - Without priority and preemption, there is no guarantee that an update will 
 not block due to a loss of capacity or due to the scheduling of another Pod
 between Pod termination and Pod creation. This is mitigated by blocking the 
 update when a Pod fails to schedule. Remediation will require operator 
 intervention. This implementation is no worse than the current behavior with 
 respect to eviction.
 - We will eventually implement a signal that is delivered to Pods to indicate 
 the 
 [reason for termination](https://github.com/kubernetes/community/pull/541).
 - StatefulSet updates will use the methodology outlined in the 
 [controller history](https://github.com/kubernetes/community/pull/594) proposal 
 for version tracking, update detection, and rollback detection.
 This will be a general implementation, usable for any Pod in a Kubernetes 
 cluster. It is, therefore, out of scope to design such a mechanism here.
 - Kubelet does not support resizing a container's resources without terminating 
 the Pod. In place resource reallocation is out of scope for this design. 
 Vertical scaling must be performed destructively.
 - The primary means of configuration update will be configuration files, 
 command line flags, environment variables, or ConfigMaps consumed as the one 
 of the former. 
 - In place configuration update via SIGHUP is not universally 
 supported, and Kubelet provides no mechanism to perform this currently. Pod 
 reconfiguration will be performed destructively.
 - Stateful applications are likely to evolve wire protocols and storage formats
  between versions. In most cases, when updating the application's Pod's 
  containers, it will not be safe to roll back or forward to an arbitrary 
  version. Controller based Pod update should work well when rolling out an 
  update, or performing a rollback, between two specific revisions of the 
  controlled API object. This is how Deployment functions, and this property is,
  perhaps, even more critical for stateful applications.

## Requirements
This design is based on the following requirements.
- Users must be able to update the containers of a StatefulSet's Pods.
  - Updates to container commands, images, resources and configuration must be 
  supported.
- The update must progress in a sequential, deterministic order and respect the 
  StatefulSet
  [identity](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity), 
  [deployment, and scaling](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#deployment-and-scaling-guarantee) 
  guarantees.
- A failed update must halt.
- Users must be able to roll back an update.
- Users must be able to roll forward to fix a failing/failed update.
- Users must be able to view the status of an update.
- Users should be able to view a bounded history of the updates that have been 
applied to the StatefulSet.

## API Objects

The following modifications will be made to the StatefulSetSpec API object.

```go
// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet 
// controller will use to perform updates. It includes any additional parameters 
// necessary to preform the update for the indicated strategy.
type StatefulSetUpdateStrategy struct {
    // Type indicates the type of the StatefulSetUpdateStrategy.
    Type StatefulSetUpdateStrategyType
    // Partition is used to communicate the ordinal at which to partition 
    // the StatefulSet when Type is PartitionStatefulSetStrategyType. This 
    // value must be set when Type is PartitionStatefulSetStrategyType, 
    // and it must be nil otherwise.
    Partition *PartitionStatefulSetStrategy

// StatefulSetUpdateStrategyType is a string enumeration type that enumerates 
// all possible update strategies for the StatefulSet controller.
type StatefulSetUpdateStrategyType string

const (
    // PartitionStatefulSetStrategyType indicates that updates will only be 
    // applied to a partition of the StatefulSet. This is useful for canaries 
    // and phased roll outs. When a scale operation is performed with this 
    // strategy, new Pods will be created from the updated specification.
    PartitionStatefulSetStrategyType StatefulSetUpdateStrategyType = "Partition"
    // RollingUpdateStatefulSetStrategyType indicates that update will be 
    // applied to all Pods in the StatefulSet with respect to the StatefulSet 
    // ordering constraints. When a scale operation is performed with this 
    // strategy, new Pods will be created from the updated specification.
    RollingUpdateStatefulSetStrategyType = "RollingUpdate"
    // OnDeleteStatefulSetStrategyType triggers the legacy behavior. Version 
    // tracking and ordered rolling restarts are disabled. Pods are recreated 
    // from the StatefulSetSpec when they are manually deleted. When a scale 
    // operation is performed with this strategy, new Pods will be created 
    // from the current specification.
    OnDeleteStatefulSetStrategyType = "OnDelete"
)

// PartitionStatefulSetStrategy contains the parameters used with the 
// PartitionStatefulSetStrategyType.
type PartitionStatefulSetStrategy struct {
    // Ordinal indicates the ordinal at which the StatefulSet should be 
    // partitioned.
    Ordinal int32
}

type StatefulSetSpec struct {
    // Replicas, Selector, Template, VolumeClaimsTemplate, and ServiceName 
    // omitted for brevity.
    
    // UpdateStrategy indicates the StatefulSetUpdateStrategy that will be 
    // employed to update Pods in the StatefulSet when a revision is made to 
    // Template or VolumeClaimsTemplate.
    UpdateStrategy StatefulSetUpdateStrategy `json:"updateStrategy,omitempty`
    
    // RevisionHistoryLimit is the maximum number of revisions that will 
    // be maintained in the StatefulSet's revision history. The revision history
    // consists of all revisions not represented by a currently applied 
    // StatefulSetSpec version. The default value is 2.
    RevisionHistoryLimit *int32 `json:revisionHistoryLimit,omitempty`
}
```

The following modifications will be made to the StatefulSetStatus API object.

```go
 type StatefulSetStatus struct {
    // ObservedGeneration and Replicas fields are omitted for brevity.
    
    // CurrentRevision, if not empty, indicates the version of PodSpecTemplate, 
    // VolumeClaimsTemplate tuple used to generate Pods in the sequence
    // [0,CurrentReplicas).
    CurrentRevision string `json:"currentRevision,omitempty"`
    
    // UpdateRevision, if not empty, indicates the version of PodSpecTemplate, 
    // VolumeClaimsTemplate tuple used to generate Pods in the sequence
    // [Replicas-UpdatedReplicas,Replicas)
    UpdateRevision string `json:"updateRevision,omitempty"`
    
    // ReadyReplicas is the current number of Pods, created by the StatefulSet
    // controller, that have a Status of Running and a Ready Condition.
    ReadyReplicas int32 `json:"readyReplicas,omitempty"`
    
    // CurrentReplicas is the number of Pods created by the StatefulSet 
    // controller from the PodTemplateSpec, VolumeClaimsTemplate tuple indicated 
    // by CurrentRevision.
    CurrentReplicas int32 `json:"currentReplicas,omitempty"`
    
    // UpdatedReplicas is the number of Pods created by the StatefulSet
    // controller from the PodTemplateSpec, VolumeClaimsTemplate tuple indicated 
    // by UpdateRevision.
    UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`
}
```

Additionally we introduce the following constant.

```go
// StatefulSetRevisionLabel is the label used by StatefulSet controller to track
// which version of StatefulSet's StatefulSetSpec was used generate a Pod.
const StatefulSetRevisionLabel = "statefulset.kubernetes.io/revision"

```
## StatefulSet Controller
The StatefulSet controller will watch for modifications to StatefulSet and Pod 
API objects. When a StatefulSet is created or updated, or when one 
of the Pods in a StatefulSet is updated or deleted, the StatefulSet
controller will attempt to create, update, or delete Pods to conform the 
current state of the system to the user declared [target state](#target-state). 

### Revised Controller Algorithm
The StatefulSet controller will use the following algorithm to continue to 
make progress toward the user declared [target state](#target-state) while 
respecting the controller's 
[identity](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity), 
[deployment, and scaling](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#deployment-and-scaling-guarantee) 
guarantees. The StatefulSet controller will use the technique proposed in 
[Controller History](https://github.com/kubernetes/community/pull/594) to 
snapshot and version its [target Object state](#target-pod-state).

1. The controller will reconstruct the 
[revision history](#history-reconstruction) of the StatefulSet.
1. The controller will 
[process any updates to its StatefulSetSpec](#specification-updates) to 
ensure that the StatefulSet's revision history is consistent with the user 
declared desired state.
1. The controller will select all Pods in the StatefulSet, filter any Pods not 
owned by the StatefulSet, and sort the remaining Pods in ordinal order.
1. For all created Pods, the controller will perform any necessary
[non-destructive state reconciliation](#pod-state-reconciliation).
1. If any Pods with ordinals in the sequence `[0,.Spec.Replicas)` have not been 
created, for the Pod corresponding to the lowest such ordinal, the controller 
will create the Pod with declared [target Pod state](#target-pod-state).
1. If all Pods in the sequence `[0,.Spec.Replicas)` have been created, but if any 
do not have a Ready Condition, the StatefulSet controller will wait for these 
Pods to either become Ready, or to be completely deleted. 
1. If all Pods in the sequence `[0,.Spec.Replicas)` have a Ready Condition, and 
if `.Spec.Replicas` is less than `.Status.Replicas`, the controller will delete 
the Pod corresponding to the largest ordinal. This implies that scaling takes 
precedence over Pod updates.
1. If all Pods in the sequence `[0,.Spec.Replicas)` have a Status of Running and 
a Ready Condition, if `.Spec.Replicas` is equal to `.Status.Replicas`, and if 
there are Pods that do not match their [target Pod state](#target-pod-state), 
the Pod with the largest ordinal in that set will be deleted.
1. If the StatefulSet controller has achieved the 
[declared target state](#target-state) the StatefulSet controller will 
[complete any in progress updates](#update-completion).
1. The controller will [report its status](#status-reporting).
1. The controller will perform any necessary
[maintenance of its revision history](#history-maintenance).

### Target State
The target state of the StatefulSet controller with respect to an individual 
StatefulSet is defined as follows. 

1. The StatefulSet contains exactly `[0,.Spec.Replicas)` Pods.
1. All Pods in the StatefulSet have the correct 
[target Pod state](#target-pod-state).

### Target Pod State
As in the [Controller History](https://github.com/kubernetes/community/pull/594) 
proposal we define the target Object state of StatefulSetSpec specification type 
object to be the `.Template` and `.VolumeClaimsTemplate`. The latter is currently 
immutable, but we will version it as one day this constraint may be lifted. This 
state provides enough information to generate a Pod and its associated 
PersistentVolumeClaims. The target Pod State for a Pod in a StatefulSet is as 
follows.
1. The Pods PersistentVolumeClaims have been created.
   - Note that we do not currently delete PersistentVolumeClaims.
1. If the Pod's ordinal is in the sequence `[0,.Spec.Replicas)` the Pod should 
have a Ready Condition. This implies the Pod is Running.
1. If Pod's ordinal is greater than or equal to `.Spec.Replicas`, the Pod 
should be completely terminated and deleted.
1. If the StatefulSet's `Spec.UpdateStrategy.Type` is equal to 
`OnDeleteStatefulSetStrategyType`, no version tracking is performed, Pods 
can be at an arbitrary version, and they will be recreated from the current 
`.Spec.Template` and `.Spec.VolumeClaimsTemplate` when the are deleted.
1. If StatefulSet's `Spec.UpdateStrategy.Type` is equal to 
`RollingUpdateStatefulSetStrategyType` then the version of the Pod should be 
as follows.
    1. If the Pod's ordinal is in the sequence `[0,.Status.CurrentReplicas)`, 
    the Pod should be consistent with version indicated by `Status.CurrentRevision`.
    1. If the Pod's ordinal is in the sequence 
    `[.Status.Replicas - .Status.UpdatedReplicas, .Status.Replicas)`
    the Pod should be consistent with the version indicated by 
    `Status.UpdateRevision`.
1. If the StatefulSet's `.Spec.UpdateStrategy.Type` is equal to 
`PartitionStatefulSetStrategyType` then the version of the Pod should be 
as follows.
    1. If the Pod's ordinal is in the sequence `[0,.Status.CurrentReplicas)`, 
    the Pod should be consistent with version indicated by `Status.CurrentRevision`.
    1. If the Pod's ordinal is in the sequence 
    `[.Status.Replicas - .Status.UpdatedReplicas, .Status.Replicas)` the Pod 
    should be consistent with the version indicated by `Status.UpdateRevision`.
    1. If the Pod does not meet either of the prior two conditions, and if 
    ordinal is in the sequence `[0, .Spec.UpdateStrategy.Partition.Ordinal)`, 
    it should be consistent with the version indicated by 
    `Status.CurrentRevision`.
    1. Otherwise, the Pod should be consistent with the version indicated 
    by `Status.UpdateRevision`.

### Pod State Reconciliation
In order to reconcile a Pod with declared desired 
[target state](#target-pod-state) the StatefulSet controller will do the 
following.

1. If the Pod is already consistent with its target state the controller will do 
nothing.
1. If the Pod is labeled with a `StatefulSetRevisionLabel` that indicates 
the Pod was generated from a version of the StatefulSetSpec that is semantically 
equivalent to, but not equal to, the [target version](#target-pod-state), the 
StatefulSet controller will update the Pod with a `StatefulSetRevisionLabel` 
indicating the new semantically equivalent version. This form of reconciliation 
is non-destructive.
1. If the Pod was not created from the target version, the Pod will be deleted 
and recreated from that version. This form of reconciliation is destructive.

### Specification Updates
The StatefulSet controller will [snapshot](#snapshot-creation) its target 
Object state when mutations are made to its `.Spec.Template` or 
`.Spec.VolumeClaimsTemplate` (Note that the latter is currently immutable).

1. When the StatefulSet controller observes a mutation to a StatefulSet's 
 `.Spec.Template` it will snapshot its target Object state and compare 
the snapshot with the version indicated by its `.Status.UpdateRevision`.
1. If the current state is equivalent to the version indicated by 
`.Status.UpdateRevision` no update has occurred. 
1. If the `Status.CurrentRevision` field is empty, then the StatefulSet has no 
revision history. To initialize its revision history, the StatefulSet controller 
will set both `.Status.CurrentRevision` and `.Status.UpdateRevision` to the 
version of the current snapshot. 
1. If the `.Status.CurrentRevision` is not empty, and if the 
`.Status.UpdateRevision` is not equal to the version of the current snapshot, 
the StatefulSet controller will set the `.Status.UpdateRevision` to the version 
indicated by the current snapshot.

### StatefulSet Revision History
The StatefulSet controller will use the technique proposed in 
[Controller History](https://github.com/kubernetes/community/pull/594) to 
snapshot and version its target Object state.

#### Snapshot Creation
In order to snapshot a version of its target Object state, it will 
serialize and store the `.Spec.Template` and `.Spec.VolumesClaimsTemplate` 
along with the `.Generation` in each snapshot. Each snapshot will be labeled
with the StatefulSet's `.Selector`.

#### History Reconstruction
As proposed in 
[Controller History](https://github.com/kubernetes/community/pull/594), in 
order to reconstruct the revision history of a StatefulSet, the StatefulSet 
controller will select all snapshots based on its `Spec.Selector` and sort them 
by the contained `.Generation`. This will produce an ordered set of 
revisions to the StatefulSet's target Object state.

#### History Maintenance 
In order to prevent the revision history of the StatefulSet from exceeding 
memory or storage limits, the StatefulSet controller will periodically prune 
its revision history so that no more that `.Spec.RevisionHistoryLimit` non-live 
versions of target Object state are preserved.

### Update Completion
The criteria for update completion is as follows.

1. If the StatefulSet's `.Spec.UpdateStrategy.Type` is equal to 
`OnDeleteStatefulSetStrategyType` then no version tracking is performed. In
this case, an update can never be in progress.
1. If the StatefulSet's `.Spec.UpdateStrategy.Type` is equal to 
`PartitionStatefulSetStrategyType` updates can not complete. The version 
indicated `.Status.UpdateRevision` will only be applied to Pods with ordinals 
in the sequence `(.Spec.UpdateStrategy.Partition.Ordinal,.Spec.Replicas)`.
1. If the StatefulSet's `.Spec.UpdateStrategy.Type` is equal to 
`RollingUpdateStatefulSetStrategyType`, then an update is complete when the 
StatefulSet is at its [target state](#target-state). The StatefulSet controller 
will signal update completion as follows.
    1. The controller will set `.Status.CurrentRevision` to the value of 
    `.Status.UpdateRevision`.
    1. The controller will set `.Status.CurrentReplicas` to 
    `.Status.UpdatedReplicas`. Note that this value will be equal to 
    `.Status.Replicas`.
    1. The controller will set `.Status.UpdatedReplicas` to 0.

### Status Reporting
After processing the creation, update, or deletion of a StatefulSet or Pod, 
the StatefulSet controller will record its status by persisting a 
StatefulSetStatus object. This has two purposes.

1. It allows the StatefulSet controller to recreate the exact StatefulSet 
membership in the event of a hard restart of the entire system.
1. It communicates the current state of the StatefulSet to clients. Using the 
`.Status.ObserverGeneration`, clients can construct a linearizable view of 
the operations performed by the controller.

When the StatefulSet controller records the status of a StatefulSet it will 
do the following.

1. The controller will increment the `.Status.ObservedGeneration` to communicate 
the `.Generation` of the StatefulSet object that was observed.
1. The controller will set the `.Status.Replicas` to the current number of 
created Pods.
1. The controller will set the `.Status.ReadyReplicas` to the current number of 
Pods that have a Ready Condition.
1. The controller will set the `.Status.CurrentRevision` and 
`.Status.UpdateRevision` in accordance with StatefulSet's 
[revision history](#statefulset-revision-history) and 
any [complete updates](#update-completion).
1. The controller will set the `.Status.CurrentReplicas` to the number of 
Pods that it has created from the version indicated by 
`.Status.CurrentRevision`.
1. The controller will set the `.Status.UpdatedReplicas` to the number of Pods 
that it has created from the version indicated by `.Status.UpdateRevision`.
1. The controller will then persist the StatefulSetStatus make it durable and 
communicate it to observers.

## API Server
The API Server will perform validation for StatefulSet creation and updates.

### StatefulSet Validation
As is currently implemented, the API Server will not allow mutation to any 
fields of the StatefulSet object other than `.Spec.Replicas` and 
`.Spec.Template.Containers`. This design imposes the following, additional 
constraints.

1. If the `.Spec.UpdateStrategy.Type` is equal to 
`PartitionStatefulSetStrategyType`, the API Server should fail validation 
if any of the following conditions are true.
   1. `.Spec.UpdateStrategy.Partition` is nil.
   1. `.Spec.UpdateStrategy.Partition` is not nil, and 
   `.Spec.UpdateStrategy.Partition.Ordinal` not in the sequence 
   `(0,.Spec.Replicas)`.
1. The API Server will fail validation on any update to a StatefulSetStatus
object if any of the following conditions are true.
    1. `.Status.Replicas` is negative.
    1. `.Status.ReadyReplicas` is negative or greater than `.Status.Replicas`.
    1. `.Status.CurrentReplicas` is negative or greater than `.Status.Replicas`.
    1. `.Status.UpdateReplicas` is negative or greater than `.Status.Replicas`.
   
## Kubectl
Kubectl will use the `rollout` command to control and provide the status of 
StatefulSet updates.

 - `kubectl rollout status statefulset <StatefulSet-Name>`: displays the status 
 of a StatefulSet update.
 - `kubectl rollout undo statefulset <StatefulSet-Name>`: triggers a rollback 
 of the current update.
 - `kubectl rollout history statefulset <StatefulSet-Name>`: displays a the 
 StatefulSets revision history.

## Usage
This section demonstrates how the design functions in typical usage scenarios.

### Initial Deployment
Users can create a StatefulSet using `kubectl apply`.

Given the following manifest `web.yaml`

```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

Users can use the following command to create the StatefulSet.

```shell
kubectl apply -f web.yaml
```

The only difference between the proposed and current implementation is that 
the proposed implementation will initialize the StatefulSet's revision history 
upon initial creation.

### Rolling out an Update
Users can create a rolling update using `kubectl apply`. If a user creates a 
StatefulSet [as above](#initial-deployment), the user can trigger a rolling 
update by updating the image (as in the manifest as below).

```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      updateStrategy: 
        type: RollingUpdate
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```


Users can use the following command to trigger a rolling update.

```shell
kubectl apply -f web.yaml
```

### Canaries
Users can create a canary using `kubectl apply`. The only difference between a
 [rolling update](#rolling-out-an-update) and a canary is that the 
 `.Spec.UpdateStrategy.Type` is set to `PartitionStatefulSetStrategyType` and 
 the `.Spec.UpdateStrategy.Partition.Ordinal` is set to `.Spec.Replicas-1`.
 
 
```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      updateStrategy: 
        type: Partition
        partition: 
          ordinal: 2
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
     
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

Users can also simultaneously scale up and add a canary. This reduces risk 
for some deployment scenarios by adding additional capacity for the canary. 
For example, in the manifest below, `.Spec.Replicas` is increased to `4` while 
`.Spec.UpdateStrategy.Partition.Ordinal` is set to `.Spec.Replicas-1`.

```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 4
  template:
    metadata:
      labels:
        app: nginx
    spec:
      updateStrategy: 
        type: Partition
        partition: 
          ordinal: 3
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

### Phased Roll Outs
Users can create a canary using `kubectl apply`. The only difference between a
 [canary](#canaries) and a phased roll out is that the 
 `.Spec.UpdateStrategy.Partition.Ordinal` is set to a value less than 
 `.Spec.Replicas-1`.
 
```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 4
  template:
    metadata:
      labels:
        app: nginx
    spec:
      updateStrategy: 
        type: Partition
        partition: 
          ordinal: 2
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

Phased roll outs can be used to roll out a configuration, image, or resource 
update to some portion of the fleet maintained by the StatefulSet prior to 
updating the entire fleet. It is useful to support linear, geometric, and 
exponential roll out of an update. Users can modify the 
`.Spec.UpdateStrategy.Partition.Ordinal`  to allow the roll out to progress.

```yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      updateStrategy: 
        type: Partition
        partition: 
          ordinal: 1
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.alpha.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

### Rollbacks
To rollback an update, users can use the `kubectl rollout` command.

The command below will roll back the `web` StatefulSet to the previous revision in 
its history. If a roll out is in progress, it will stop deploying the target 
revision, and roll back to the current revision.

```shell
kubectl rollout undo statefulset web
```

### Rolling Forward
Rolling back is usually the safest, and often the fastest, strategy to mitigate
deployment failure, but rolling forward is sometimes the only practical solution 
for stateful applications (e.g. A user has a minor configuration error but has 
already modified the storage format for the application). Users can use 
sequential `kubectl apply`'s to update the StatefulSet's current 
[target state](#target-state). The StatefulSet's `.Spec.GenerationPartition` 
will be respected, and it therefore interacts well with canaries and phased roll
 outs.

## Tests
- Updating a StatefulSet's containers will trigger updates to the StatefulSet's 
Pods respecting the 
[identity](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity) 
and [deployment, and scaling](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#deployment-and-scaling-guarantee) 
guarantees.
- A StatefulSet update will block on failure.
- A StatefulSet update can be rolled back.
- A StatefulSet update can be rolled forward by applying another update. 
- A StatefulSet update's status can be retrieved.
- A StatefulSet's revision history contains all updates with respect to the 
configured revision history limit.
- A StatefulSet update can create a canary.
- A StatefulSet update can be performed in stages.

## Future Work
In the future, we may implement the following features to enhance StatefulSet 
updates.

### Termination Reason
Without communicating a signal indicating the reason for termination to a Pod in 
a StatefulSet, as proposed [here](https://github.com/kubernetes/community/pull/541),
the tenant application has no way to determine if it is being terminated due to 
a scale down operation or due to an update. 

Consider a BASE distributed storage application like Cassandra, where 2 TiB of 
persistent data is not atypical, and the data distribution is not identical on 
every server. We want to enable two distinct behaviors based on the reason for 
termination.

- If the termination is due to scale down, during the configured termination 
grace period, the entry point of the Pod should cause the application to drain 
its client connections, replicate its persisted data (so that the cluster is not 
left under replicated) and decommission the application to remove it from the 
cluster.
- If the termination is due to a temporary capacity loss (e.g. an update or an 
image upgrade), the application should drain all of its client connections, 
flush any in memory data structures to the file system, and synchronize the 
file system with storage media. It should not redistribute its data.

If the application implements the strategy of always redistributing its data, 
we unnecessarily decrease recovery time during an update and incur the 
additional network and storage cost of two full data redistributions for every 
updated node.
It should be noted that this is already an issue for Node cordon and Pod eviction 
(due to drain or taints), and applications can use the same mitigation as they 
would for these events for StatefulSet update.

### VolumeTemplatesSpec Updates
While this proposal does not address 
[VolumeTemplateSpec updates](https://github.com/kubernetes/kubernetes/issues/41015), 
this would be a valuable feature for production users of storage systems that use
intermittent compaction as a form of garbage collection. Applications that use 
log structured merge trees with size tiered compaction (e.g Cassandra) or append 
only B(+/*) Trees (e.g Couchbase) can temporarily double their storage requirement 
during compaction. If there is insufficient space for compaction 
to progress, these applications will either fail or degrade until 
additional capacity is added. While, if the user is using AWS EBS or GCE PD, 
there are valid manual workarounds to expand the size of a PD, it would be 
useful to automate the resize via updates to the StatefulSet's 
VolumeClaimsTemplate.

### In Place Updates
Currently configuration, images, and resource request/limits updates are all 
performed destructively. Without a [termination reason](https://github.com/kubernetes/community/pull/541)
implementation, there is little value to implementing in place image updates, 
and configuration and resource request/limit updates are not possible.
When [termination reason](#https://github.com/kubernetes/kubernetes/issues/1462) 
is implemented we may modify the behavior of StatefulSet update to only update, 
rather than delete and create, Pods when the only mutated value is the container
 image, and if resizable resource request/limits is implemented, we may extend 
 the above to allow for updates to Pod resources.
