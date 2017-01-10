# DaemonSet Updates

**Author**: @madhusudancs, @lukaszo, @janetkuo

**Status**: Proposal

## Abstract

A proposal for adding the update feature to `DaemonSet`.

Users already can update a `DaemonSet`, which will not cause changes to its
subsequent pods, until those pods are killed. In this proposal, we plan to add
a "RollingUpdate" strategy which allows DaemonSet to downstream its changes to
pods. 

## Implementation 

### API Object 

To enable DaemonSet upgrades, the `DaemonSet` API object will have the following
structure:

```go 
type DaemonSetUpdateStrategy struct {
	// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". 
	// Default is OnDelete.
	Type DaemonSetUpdateStrategyType

	// Rolling update config params. Present only if DaemonSetUpdateStrategy =
	// RollingUpdate.
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be. Same as DeploymentStrategy.RollingUpdate.
	RollingUpdate *RollingUpdateDaemonSet
}

type DaemonSetUpdateStrategyType string

const (
	// Replace the old daemons by new ones using rolling update i.e replace them on each node one after the other.
	RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"
        
	// Replace the old daemons only when it's killed
	OnDeleteDaemonSetStrategyType DaemonSetUpdateStrategyType = "OnDelete"
)

// Spec to control the desired behavior of daemon set rolling update.
type RollingUpdateDaemonSet struct {
	// The maximum number of DaemonSet pods that can be unavailable during the
	// update. Value can be an absolute number (ex: 5) or a percentage of total
	// number of DaemonSet pods at the start of the update (ex: 10%). Absolute
	// number is calculated from percentage by rounding up.
	// This cannot be 0.
	// Default value is 1.
	// Example: when this is set to 30%, 30% of the currently running DaemonSet
	// pods can be stopped for an update at any given time. The update starts
	// by stopping at most 30% of the currently running DaemonSet pods and then
	// brings up new DaemonSet pods in their place. Once the new pods are ready,
	// it then proceeds onto other DaemonSet pods, thus ensuring that at least
	// 70% of original number of DaemonSet pods are available at all times
	// during the update.
	MaxUnavailable intstr.IntOrString
}

// DaemonSetSpec is the specification of a daemon set.
type DaemonSetSpec struct {
	// Selector is a label query over pods that are managed by the daemon set.
	// Must match in order to be controlled.
	// If empty, defaulted to labels on Pod template.
	// More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,1,opt,name=selector"`

	// Template is the object that describes the pod that will be created.
	// The DaemonSet will create exactly one copy of this pod on every node
	// that matches the template's node selector (or on every node if no node
	// selector is specified).
	// More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template
	Template v1.PodTemplateSpec `json:"template" protobuf:"bytes,2,opt,name=template"`

	// Update strategy to replace existing DaemonSet pods with new pods.
	UpdateStrategy DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
}

const (
	// DefaultDaemonSetUniqueLabelKey is the default key of the labels that is added
	// to daemon set pods to distinguish between old and new pod templates during
	// DaemonSet update. See DaemonSetSpec's UniqueLabelKey field for more information.
	DefaultDaemonSetUniqueLabelKey string = "daemonset.kubernetes.io/podTemplateHash"
)

// DaemonSetStatus represents the current status of a daemon set.
type DaemonSetStatus struct {
	// CurrentNumberScheduled is the number of nodes that are running at least 1
	// daemon pod and are supposed to run the daemon pod.
	CurrentNumberScheduled int32

	// NumberMisscheduled is the number of nodes that are running the daemon pod, but are
	// not supposed to run the daemon pod.
	NumberMisscheduled int32

	// DesiredNumberScheduled is the total number of nodes that should be running the daemon
	// pod (including nodes correctly running the daemon pod).
	DesiredNumberScheduled int32

	// NumberReady is the number of nodes that should be running the daemon pod and have one
	// or more of the daemon pod running and ready.
	NumberReady int32

	// ObservedGeneration is the most recent generation observed by the daemon set controller.
	ObservedGeneration int64

	// UpdatedNumberScheduled is the total number of nodes that are running updated
	// daemon pod
	UpdatedNumberScheduled int32 `json:"updatedNumberScheduled"`

}
```

### Controller 

#### DaemonSet Controller 

The DaemonSet Controller will make DaemonSet updates happen. It will watch
DaemonSets in etcd. 

For each pending DaemonSet updates, it will:

1. Find all pods whose label is matched by `DaemonSetSpec.Selector`. 
1. Find all nodes that should run these pods created by this DaemonSet.
1. Find an existing PodTemplate whose `Template` is the same as
   `DaemonSetSpec.Template`
   - If not found, create a new PodTemplate from `DaemonSetSpec.Template`
   - The name will be `<DaemonSet-Name>-<Hash-of-pod-template>`
   - Add "pod-template-hash" labels to `PodTemplate.Metadata.Labels`, value be
     the hash of `PodTemplate.Spec`
1. Create daemon pods on nodes when they should have those pods running but not
   yet, and add the same "pod-template-hash" labels to their `Metadata.Labels`.
   Otherwise, delete running daemon pods that shouldn't be running on nodes. 
1. Check `DaemonSetUpdateStrategy`:
   - If `OnDelete`: do nothing
   - If `RollingUpdate`:
     - From all nodes that should run daemon pods, check the daemon pod's
       "pod-template-hash" label, and kill it if it doesn't equal to the hash of
       `DaemonSetSpec.Template.Spec` and if it won't violate `MaxUnavailable`.
     - Go back to step 1 
1. Cleanup 

If DaemonSet Controller crashes during an update, it can still recover. 


## Updating DaemonSets mid-way

Users can update an updated DaemonSet before its rollout completes.
In this case, the existing daemon pods will not continue rolling out and the new
one will begin rolling out.


## Deleting DaemonSets

Deleting a DaemonSet (with cascading) will delete all its pods and podtemplates. 


## DaemonSet Strategies

DaemonSetStrategy specifies how the new daemon pods should replace existing ones.
To begin with, we will support 2 types:

* On delete: Do nothing, until existing daemon pods are killed (for backward
  compatibility).
* Rolling update: We gradually kill existing ones while creating the new one.


## Tests

- Updating a RollingUpdate DaemonSet will trigger updates to its daemon pods. 
- Updating a OnDelete DaemonSet will not trigger updates, until the pods are
  killed. 
- Users can use node labels to choose which nodes this DaemonSet should target.
  DaemonSet updates only affect pods on those nodes.
  - For example, some nodes may be running manifest pods, and other nodes will
    be running daemon pods 
- DaemonSet should support at most one daemon pod per node guarantee.
  - Adding or deleting nodes won't break that.
- Users should be able to specify acceptable downtime of their daemon pods, and
  DaemonSet updates should respect that. 


## Future

In the future, we may:

- Support more DaemonSet update types
- Support history and rollback
- Add clean up policy for history
- Add minReadySeconds and make DaemonSet update strategy respect it
- Support pausing DaemonSets
