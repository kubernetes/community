# StatefulSet Updates

**Author**: kow3ns@

**Status**: Proposal

## Abstract
Currently (as of Kubernetes 1.6), `.Spec.Replicas`, and 
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
containers to add images.

### Out of Scope
- As the administrator of a stateful application, in order to increase the 
applications storage capacity, I want to update PersistentVolumes.
- As the administrator of a stateful application, in order to update the 
network configuration of the application, I want to update Services and 
container ports in a consistent way.
- As the administrator of a stateful application, when I scale my application 
horizontally, I want associated PodDistruptionBudgets to be adjusted to 
compensate for the application's scaling.

## Assumptions
 - StatefulSet update must support singleton StatefulSets. However, an update in
 this case will cause a temporary outage. This is acceptable as a single 
 process application is, by definition, not highly available.
 - Disruption in Kubernetes is controlled by PodDistruptionBugets. As 
 StatefulSet updates progress one Pod at a time, and only occur when all 
 other Pods have a Status of Running and a Ready Condition, they can not 
 violate reasonable PodDisrutptionBugdets.
 - Without priority and preemption, there is no guarantee that an update will 
 not block due to a loss of capacity or due to the scheduling of another Pod
 between Pod termination and Pod creation. This is mitigated by blocking the 
 update when a Pod fails to schedule. Remediation will require operator 
 intervention. This implementation is no worse than the current behavior with 
 respect to eviction.
 - We will eventually implement a signal that is delivered to Pods to indicate 
 the 
 [reason for termination](https://github.com/kubernetes/kubernetes/issues/1462).
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
  version. StatefulSet update should work well when rolling out an update, 
  or performing a rollback, between two specific revisions of the StatefulSet.

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
- Users must be able to rollback an update.
- Users must be able to roll forward to fix a failing/failed update.
- Users must be able to view the status of an update.
- Users should be able to view a bounded history of the updates that have been 
applied to the StatefulSet.

## API Object

The following modifications will be made to the StatefulSetStatus API object.

```go
 type StatefulSetStatus struct {
    // ObservedGeneration and Replicas fields are ommitted for brevity.

 	// TemplateRevision, if not nil, is the revision of the PodTemplate that was 
 	// used to create Pods with ordinals in the sequence
 	// [0,CurrentReplicas).
 	TemplateRevision  *int64 `json:"templateRevision,omitempty"`
 	
 	// TargetTemplateRevision, if not nil, is the revision of the PodTemplate 
 	// that was used to create Pods with ordinals in the sequence 
    // [Replicas - UpdatedReplicas, Replicas).
 	TargetTemplateRevision *int64 `json:"targetTemplateRevision,omitempty"`
 	
 	// ReadyReplicas is the current number of Pods, created by the StatefulSet
    // controller, that have a Status for Running and a Ready Condition.
 	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
 	
 	// CurrentRevisionReplicas is the number of Pods created by the StatefulSet 
    // controller from the PodTemplateSpec indicated by CurrentTemplateRevision.
 	CurrentReplicas int32 `json:"currentReplicas,omitempty"`
 	
 	// UpdatedReplicas is the number of Pods created by the StatefulSet
    // controller from the PodTemplateSpec indicated by TargetTemplateRevision.
 	UpdatedReplicas int32 `json:"taretReplicas,omitempty"`
}
```

The following modifications will be made to the StatefulSetSpec API object.

```go
type StatefulSetSpec struct {
    // Replicas, Selector, Template, VolumeClaimsTemplate, and ServiceName 
    // ommitted for brevity.
	v1.PodTemplateSpec `json:"template"`

	// TemplateRevision is a monotonically increasing, 64 bit, integer used to 
	// indicate the version of the of the PodTemplateSpec. If nil, the 
	// StatefulSetController has not initialized its revision history,
	// change tracking is not enabled, and all Pods will be created from 
	// Template.
	TemplateRevision  *int64 `json:"templateRevision"`

	// RevisionPartition partitions the Pods in the StatefulSet by ordinal such 
    // that all Pods with a lower ordinal will be created from the PodTemplate that
    // represents the current revision of the StatefulSet's revision history and 
    // all Pods with an a greater or equal ordinal will be created from the 
    // PodTemplate that represents the target revision of the StatefulSet's 
    // revision history.
	RevisionPartition *int32 `json:"revisionPartition,omitempty`

	// RevisionHistoryLimit is the maximum number of PodTemplates that will 
	// be maintained in the StatefulSet's revision history. It must be at 
	// least two.
	RevisionHisotryLimit int32 `json:historyRevisionDepth,omitempty`
}
```

Additionally, we introduce the following constants.

```go 
// StatefulSetPodTemplateLabel is the label applied to a PodTemplate to allow
// the StatefulSet controller to select the PodTemplates in its revision 
// history.
const StatefulSetPodTemplateLabel = "created-by-statefulset"

// StatefulSetTemplateRevisionLabel is the label applied to a PodTemplate or 
// Pod to indicate the position of the object's Template in the revision 
// history of a StatefulSet.
const StatefulSetTemplateRevisionLabel = "statefulset-template-revision"
```

## StatefulSet Controller
The StatefulSet controller will watch for modifications to StatefulSet and Pod 
API objects. When a StatefulSet is created or updated, or when one 
of the Pods in a StatefulSet is updated or deleted, the StatefulSet
controller will attempt to create, update, or delete Pods to conform the 
current state of the system to the user declared target state. 
The user declared target state of the system, with respect to an individual 
StatefulSet, is determined as below.

### Target State
The declared target state of a StatefulSet requires that all Pods in the 
StatefulSet conform to exactly one or two PodTemplates in the StatefulSet's 
revision history. If the declared target state references two PodTemplates, as 
is the case when a user wants to perform a canary update or a phased roll out, 
they are partitioned around an ordinal such that all Pods with a lower ordinal 
conform to one PodTemplate and all Pods with a greater or equal ordinal 
conform to the other. The conditions that define this state in terms of the 
StatefulSet's StatefulSetSpec and StatefulSetStatus are below.

1. The StatefulSet contains exactly `[0,.Spec.Replicas)` Pods.
1. If StatefulSet's `.Spec.RevisionPartition` is nil, then the following is true.
    1. The StatefulSet's `.Status.TemplateRevision` is equal to its 
    `.Status.TargetRevision`.
    1. All Pods in the StatefulSet have been generated from the PodTemplate 
    labeled with a `StatefulSetTemplateRevisionLabel` equal to its 
    `.Status.TemplateRevision`.
1. If the StatefulSet's `.Spec.RevisionPartition` is not nil, then the following 
is true.
    1. All Pods with ordinals is the sequence `[0,.Spec.RevisionPartition)` have 
     been generated from the PodTemplate in the StatefulSet's revision history
     that is labeled with a `StatefulSetTemplateRevisionLabel` equal to 
     `.Status.TemplateRevision`.
    1. All Pods with ordinals in the sequence 
     `[Spec.RevisionParition,.Spec.Replicas)` have been created with the 
     PodTemplate in the StatefulSet's revision history that is labeled with a 
     `StatefulSetTemplateRevisionLabel` equal to 
     `.Status.TargetTemplateRevision`.

### Revised Controller Algorithm
The StatefulSet controller will use the following algorithm to continue to 
make progress toward the user declared [target state](#target-state) while 
respecting the controller's 
[identity](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#pod-identity), 
[deployment, and scaling](https://kubernetes.io/docs/concepts/abstractions/controllers/statefulsets/#deployment-and-scaling-guarantee) 
guarantees.

1. The controller will 
[reconstruct the revision history](#history-reconstruction) of the StatefulSet.
1. The controller will process any [template updates](#template-updates) to 
ensure that the StatefulSet's revision history is consistent with the user 
declared desired state.
1. The controller will select all Pods in the StatefulSet, filter any Pods not 
owned by the StatefulSet, and sort the remaining Pods in ordinal order.
1. If any Pods with ordinals in the sequence `[0,.Spec.Replicas)` have not been 
created, for the Pod corresponding to the lowest such ordinal, the controller 
will [select the PodTemplate](#podtemplate-selection) from the StatefulSet's 
revision history corresponding to the Pod's ordinal and create the Pod.
1. If all Pods in the sequence `[0,.Spec.Replicas)` have been created, but if any 
have a Status other than Running or do not have a Ready Condition, the 
StatefulSet controller will wait for these Pods to either become Running and 
Ready, or to be completely deleted. 
1. If all Pods in the sequence `[0,.Spec.Replicas)` have a Status of Running and 
a Condition indicating Ready, and if `.Spec.Replicas` is less than 
`.Status.Replicas`, the controller will delete the Pod corresponding to the 
largest ordinal. This implies that scaling takes precedence over Pod updates.
1. If all Pods in the range `[0,.Spec.Replicas)` have a Status of Running and 
a Ready Condition, if `.Spec.Replicas` is equal to `.Status.Replicas`, and if 
there are Pods that do not match the 
[declared desired PodTemplate](#podtemplate-selection), the Pod corresponding to 
the largest ordinal will be deleted.
1. If the StatefulSet controller has achieved the 
[declared target state](#target-state), and if that state has a 
`.Spec.ParitionOrdinal` equal to `0`, the StatefulSet controller will 
[complete any in progress updates](#update-completion).
1. The controller will [report its status](#status-reporting).
1. The controller will perform any necessary
[maintenance of its revision history](#history-maintenance).

### StatefulSet Revision History
The StatefulSet controller will use labeled, versioned PodTemplates to keep a 
history of updates preformed on a StatefulSet. The number of stored PodTemplates 
is considered to be the limit of the StatefulSet's revision history. The 
maximum revision history limit for a StatefulSet must be at least two, but it 
may be greater.

#### PodTemplate Creation
When the StatefulSet controller creates a PodTemplate for a StatefulSet, it will 
do the following.

1. The controller will set the PodTemplate's `.PodTemplateSpec` field to the 
StatefulSet's `.Spec.Template` field.
1. The controller will create a ControllerRef object in the PodTemplate's 
`.OwnerReferences` list to mediate selector overlapping.
1. The controller will label the PodTemplate with a 
`StatefulSetPodTemplateLabel` set to the StatefulSet's `.Name` to allow for 
selection of the PodTemplates that comprise the StatefulSet's revision history.
1. The controller will label the PodTemplate with a 
`StaefulSetTemplateRevisionLabel` set to the StatefulSet's 
`.Spec.TemplateRevision`.
1. The controller will set the Name of the PodTemplate to a concatenation of the 
`.Name` of the StatefulSet and the `.Spec.TemplateRevision`.
1. The controller will then create the PodTemplate.

#### PodTemplate Deletion
When the `StatefulSet` controller deletes a PodTemplate in the revision 
history of a StatefulSet it will do the following.

1. If the PodTemplate's ControllerRef does not match the StatefulSet, the 
controller will not delete the PodTemplate. In this way, we prevent selector 
overlap from causing the deletion of PodTemplates that are part of another 
object's revision history. In practice, these PodTemplates will be filtered out
prior to history maintenance.
1. If the PodTemplate's ControllerRef matches the StatefulSet, the 
StatefulSet controller orphan the PodTemplate by removing its ControllerRef, 
and it will allow the PodTemplate to be deleted via garbage collection.

#### History Reconstruction
In order to reconstruct the history of revisions to a StatefulSet, the 
StatefulSet controller will do the following.

1. If the StatefulSet's `.Spec.TemplateRevision` is nil, the StatefulSet 
has never been updated, and its history has never been initialized. This is 
the state the object will be in when a cluster is first upgraded from a version
that does not support StatefulSet update to a version that does. In this case,
the controller will not enforce PodTemplate revisions. When creating Pods, 
it will always use the StatefulSet's `.Spec.Template`. Otherwise, the controller 
will continue as below.
1. The controller will select all PodTemplates with a 
`StatefulSetPodTemplateLabel` matching the `.Name` field of the StatefulSet.
1. The controller will filter out all PodTemplates that do not contain a 
ControllerRef matching the the StatefulSet. If the controller selects 
PodTemplates that it does not own, it will report an error, but it will continue
reconstructing the StatefulSet's history.
1. The controller will filter out all PodTemplates that do not have a 
`StatefulSetTemplateRevisionLabel` mapped to a valid revision. This can only 
occur if the user purposefully deletes the label. In this case, the 
controller will report an error, but it will continue reconstructing the 
StatefulSet's revision history.
1. For all the remaining PodTemplates, the controller will sort them in 
ascending order by the value mapped to their `StatefulSetTemplateRevisionLabel`. 
This will reconstruct a list of PodTemplates from oldest to newest. Note that, 
as the revision is monotonically increasing for an individual StatefulSet, and
as we use ControllerRef to mitigate selector overlap, the StatefulSet's history
 is a strictly ordered set.

#### History Maintenance 
In order to prevent the revision history of the StatefulSet from exceeding 
memory or storage limits, the StatefulSet controller will periodically prune 
the oldest PodTemplates from the StatefulSet's revision history. 

1. The StatefulSet controller will 
[reconstruct the revision history](#history-reconstruction) 
of the StatefulSet.
1. If the number of PodTemplates in the StatefulSet's revision history is 
greater than the StatefulSet's `.Spec.RevisionHistoryLimit, the 
StatefulSet controller will delete PodTemplates, starting with the head of 
the revision history, until the limit of the revision history is equal to 
the StatefulSet's `.Spec.RevisionHistoryLimit`.
1. As a StatefulSet's `.Spec.RevisionHistoryLimit` is always at least two, and 
as the PodTemplates corresponding to `.Status.TemplateRevision` 
or `.Status.TargetTemplateRevision` are always the most recent PodTemplates 
in the revision history, the StatefulSet controller will not delete any 
`PodTemplates` that represent the current or target revisions.

### Template Updates
The StatefulSet controller will create PodTemplates upon mutation of the 
`.Spec.Template` of a StatefulSet.

1. When the StafefulSet controller observes a mutation to a StatefulSet's 
 `.Spec.Template` it will compare the `.Spec.TemplateRevision` to the 
 `.Status.TargetTemplateRevision`.
1. If the `.Spec.TemplateRevision` is equivalent to the 
`.Status.TargetTemplateRevision`, no update has occurred. Note that, in the 
event that both are nil, they are considered to be equivalent, and we expect 
this to occur after an initial upgrade to a version of Kubernetes that supports 
StatefulSet update form one that does not.
1. If the `.Status.TemplateRevision` field is nil, and the 
`.Spec.TemplateRevision` is not nil, then the StatefulSet has no revision 
history. To initialize its revision history, the StatefulSet controller will 
set both `.Status.TemplateRevision` and `.Status.TargetTemplateRevision` 
to `.Spec.TemplateRevision` and 
[create a new PodTemplate](#podtemplate-creation). 
1. If the `.Status.TemplateRevision` is not nil, and if the 
`.Spec.TemplateRevision` is not equal to the `.Status.TargetTemplateRevision`, 
the StatefulSet controller will do the following.
    1. The controller will 
    [reconstruct the revision history](#history-reconsturction) of the 
    StatefulSet.
    1. If the revision history of the StatefulSet contains a PodTemplate 
    whose `.PodTemplateSpec` is semantically, deeply equivalent to the 
    StatefulSet's `.Spec.Template`, the youngest such PodTemplate will be used 
    as the target PodTemplate. 
    1. If no such PodTemplate exists, the StatefulSet controller will 
    [create a new PodTemplate](#podtemplate-creation) from the StatefulSet's 
    `.Spec.Template`, and it will use this as the target PodTemplate.
    1. The controller will update the StatefulSet's `.Status.TargetTemplate` 
    based on the selection made above.

### PodTemplate Selection
When the StatefulSet controller creates the Pods in a StatefulSet, it will use 
the following criteria to select the PodTemplateSpec used to create a 
Pod. These criteria allow the controller to continue to make progress toward 
its target state, while respecting its guarantees and allowing for rolling 
updates back and forward.

1. If the StatefulSet's `.Spec.TemplateRevision` is nil, then the cluster 
has been upgraded from a version that does not support StatefulSet update to 
a version that does. 
    1. In this case the `.Spec.Template` is the current revision, 
    and no Pods in the StatefulSet should be labeled with a 
    `StatefulSetPodTemplateRevision` label. 
    1. The StatefulSet will initialize its revision history on the first 
    update to its `.Spec.Template`.
1. If the StatefulSet's `.Spec.TemplateRevision` is equal to its 
`.Status.TemplateRevision`, then there is no update in progress and all 
Pods will be created from the PodTemplate matching this revision.
1. If the Pod's ordinal is in the sequence `[0,.Status.CurrentReplicas)`, 
then it was previously created from the PodTemplate matching the 
StatefulSet's `.Status.TemplateRevision`, and it will be recreated 
from this PodTemplate.
1. If the Pod's ordinal is in the sequence
 `[.Spec.Replicas-.Status.UpdatedReplicas,.Spec.Replicas)`, then it was
 previously created from the PodTemplate matching the StatefulSet's, 
 `.Status.TargetTemplateRevision`, and it will be recreated from this 
 PodTemplate.
1. If the ordinal does not meet either of the prior two conditions, and 
if ordinal is in the sequence `[0, .Spec.RevisionPartition)`, it will be created 
from the PodTemplate matching the StatefulSet's 
`.Status.TemplateRevision`.
1. Otherwise, the Pod is created from the PodTemplate matching the 
StatefulSet's `.Status.TargetTemplateRevision`. 

### Update Completion
A StatefulSet update is complete when the following conditions are met.

1. All Pods with ordinals in the sequence `[0,.Spec.Replicas)` have a Status of 
Running and a Ready Condition.
1. The StatefulSet's `.Spec.RevisionPartition` is equal to `0`. 
1. All Pods in the StatefulSet are labeled with a 
`StatefulSetTemplateRevisionLabel` equal to the StatefulSet's 
`.Status.TargetTemplateRevision` (This implies they have been created from 
the PodTemplate at that revision).

When a StatefulSet update is complete, the controller will signal completion by 
doing the following.

1. The controller will set the StatefulSet's `.Status.TemplateRevision` to its 
`.Status.TargetTemplateRevision`. 
1. The controller will set the StatefulSet's `Status.CurrentReplicas` to its 
`Status.UpdatedReplicas`.
1. The controller will set the StatefulSet's `Status.UpdatedReplicas` to 0.

### Status Reporting
After processing the creation, update, or deletion of a StatefulSet or Pod, 
the StatefulSet controller will record its status by persisting the 
a StatefulSetStatus object. This has two purposes.

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
Pods that have a Status of Running and a ReadyCondition.
1. The controller will set the `.Status.TemplateRevision` and 
`.Status.TargetTemplateRevision` 
in accordance with [maintaining its revision history](#history-maintenance) 
and the status of any [complete updates](#update-completion).
1. The controller will set the `.Status.CurrentReplicas` to the number of 
Pods that it has created from the PodTemplate that corresponds to the 
current revision of the StatefulSet.
1. The controller will set the `.Status.UpdatedReplicas` to the number of Pods 
that it has created from the PodTemplate that corresponds to the target 
revision of the StatefulSet.
1. The controller will then persist the StatefulSetStatus make it durable and 
communicate it to observers.

## API Server
The API Server will perform validation for StatefulSet updates and ensure that 
a StatefulSet's `.Spec.TemplateRevision` is a generator for a strictly 
monotonically increasing sequence.

### StatefulSet Validation
As is currently implemented, the API Server will not allow mutation to any 
fields of the StatefulSet object other than `.Spec.Replicas` and 
`.Spec.Template.Containers`. This design imposes the following, additional 
constraints.

1. The `.Spec.RevisionHistoryDepty` must be greater than or equal to `2`. 
1. The `.Spec.PositionOrdinal` must be in the sequence `[0,.Spec.Replicas)`. 

### TemplateRevision Maintenance
It will be the responsibility of the API Server to enforce that updates to 
StatefulSet's `.Spec.Template` atomically increment the 
`.Spec.TemplateRevision` counter. There is no need for the value to be 
strictly sequential, but it must be strictly, monotonically increasing.
As validation will not allow mutation to any field other than the 
`.Spec.Template.Containers` field, the API Server need not track all fields of 
StatefulSet's `.Spec` for modifications, but it must trigger an update to the 
revision when the current and previous `.Spec.Template` versions fail a test for
deep semantic equality.

## Kubectl
Kubectl will  use the `rollout` command to control and provide the status of 
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
Users can create a StatefulSet using `kubectl create`.

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
        image: gcr.io/google_containers/nginx-slim:0.8
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
kubectl create -f web.yaml
```

The only difference between the proposed and current implementation is that 
the proposed implementation will initialize the StatefulSet's revision history 
upon initial creation.

### Rolling out an Update
Users can create a rolling update using `kubectl apply`. If a user creates a 
StatefulSet [as above](#initial-deployment), the user can trigger a rolling 
update by updating image (as in the manifest as below).

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
        image: gcr.io/google_containers/nginx-slim:0.9
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
 `.Spec.RevisionPartition` is set to `.Spec.Replicas - 1`.
 
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
        image: gcr.io/google_containers/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
      revisionPartition: 2
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
      containers:
      - name: nginx
        image: gcr.io/google_containers/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
      partitionOrdinal: 3
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

### Staged Roll Outs
Users can create a canary using `kubectl apply`. The only difference between a
 [canary](#canaries) and a staged roll out is that the `.Spec.RevisionPartition` 
 is set to value less than `.Spec.Replicas - 1`.
 
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
        image: gcr.io/google_containers/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
      revisionParition: 2
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

Staged roll outs can be used to roll out a configuration, image, or resource 
update to some portion of the fleet maintained by the StatefulSet prior to 
updating the entire fleet. It is useful to support linear, geometric, and 
exponential roll out of an update. Users can modify the 
`.Spec.RevisionPartition` to allow the roll out to progress.

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
        image: gcr.io/google_containers/nginx-slim:0.9
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
      revisionPartition: 1
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
for stateful applications (e.g. A users has a minor configuration error but has 
already modified the storage format for the application). Users can use 
sequential `kubectl apply`'s to update the `.Status.TargetRevision` of a 
StatefulSet. This will respect the `.Spec.RevisionPartition` with respect to the 
target state, and it therefor interacts well with canaries and staged roll outs.
Note that, while users can update the target template revision, they can not 
update the current template revision. The only way to advance the current 
template revision is to successfully complete an update.

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
a StatefulSet, as proposed [here](https://github.com/kubernetes/kubernetes/issues/1462),
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
intermittent compaction as a form of garbage collection. Application that use 
log structured merge trees with size tiered compaction (e.g Cassandra) or append 
only B(+/*) Trees (e.g Couchbase) can temporarily double their storage usage when 
compacting their on disk storage. If there is insufficient space for compaction 
to progress, these applications will either fail or degrade  until 
additional capacity is added. While, if the user is using AWS EBS or GCE PD, 
there are valid manual workarounds to expand the size of a PD, it would be 
useful to automate the resize via updates to the StatefulSet's 
VolumeClaimsTemplate.

### In Place Updates
Currently configuration, images, and resource request/limits updates are all 
performed destructively. Without a [termination reason](https://github.com/kubernetes/kubernetes/issues/1462)
implementation, there is little value to implementing in place image updates, 
and configuration and resource request/limit updates are not possible.
When [termination reason](#https://github.com/kubernetes/kubernetes/issues/1462) 
is implemented we may modify the behavior of StatefulSet update to only update, 
rather than delete and create, Pods when the only mutated value is the container
 image, and if resizable resource request/limits is implemented, we may extend 
 the above to allow for updates to Pod resources.
