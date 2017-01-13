# DaemonSet Updates

**Author**: @madhusudancs, @lukaszo, @janetkuo

**Status**: Proposal

## Abstract

A proposal for adding the update feature to `DaemonSet`. This feature will be
implemented on server side (in `DaemonSet` API). 

Users already can update a `DaemonSet` today (Kubernetes release 1.5), which will
not cause changes to its subsequent pods, until those pods are killed. In this
proposal, we plan to add a "RollingUpdate" strategy which allows DaemonSet to
downstream its changes to pods. 

## Requirements

In this proposal, we design DaemonSet updates based on the following requirements:

- Users can have a controlled number of in-flight updates of a DaemonSet (e.g.
  DaemonSet only updates one pod at a time)
- Users can monitor the status of a DaemonSet update (e.g. the number of pods
  that are updated and healthy)
- A broken DaemonSet update should not continue 
- Users should be able to update a DaemonSet, even during an ongoing DaemonSet
  upgrade (e.g. update the DaemonSet to fix a broken DaemonSet update)
- Users should be able to view the history of previous DaemonSet updates 
- Users can figure out the revision of a DaemonSet's pod (e.g. which version is
  this DaemonSet pod?)

Here are some potential requirements that haven't been covered by this proposal:

- Users can have rate-limited DaemonSet updates (e.g. only upgrade 10% per hour)
- DaemonSet should provide at-most-one guarantee per node (i.e. at most one pod
  from a DaemonSet can exist on a node at any time)
- Uptime is critical for each pod of a DaemonSet during upgrade (e.g. the time
  from a DaemonSet pods being killed to recreated and healthy should be < 5s)
- Each DaemonSet pod can still fit on the node after being updated
- Some DaemonSets require the node to be drained before the DeamonSet's pod on it 
  is updated (e.g. logging daemons)
- DaemonSet's pods are implicitly given higher priority than non-daemons
- DaemonSets can only be operated by admins (i.e. people who manage nodes)
  - This is required if we allow DaemonSet controllers to drain, cordon,
    uncordon nodes, evict pods, or allow DaemonSet pods to have higher priority

## Implementation 

### API Object 

To enable DaemonSet upgrades, `DaemonSet` related API object will have the following
changes:

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
	// See https://github.com/kubernetes/kubernetes/issues/35345
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
	// The maximum number of DaemonSet pods that can be updated at a time.
	// Value can be an absolute number (ex: 5) or a percentage of total
	// number of DaemonSet pods at the start of the update (ex: 10%). Absolute
	// number is calculated from percentage by rounding up.
	// This cannot be 0.
	// Default value is 1.
	// Example: when this is set to 30%, 30% of the currently running DaemonSet
	// pods can be updated at any given time. The update starts
	// by stopping at most 30% of the currently running DaemonSet pods and then
	// brings up new DaemonSet pods in their place. Once the new pods are ready,
	// it then proceeds onto other DaemonSet pods, thus ensuring that at least
	// 70% of original number of DaemonSet pods are available at all times
	// during the update.
	MaxInFlight intstr.IntOrString
}

// DaemonSetSpec is the specification of a daemon set.
type DaemonSetSpec struct {
	// Note: Existing fields, including Selector and Template are ommitted in
	// this proposal.  

	// Update strategy to replace existing DaemonSet pods with new pods.
	UpdateStrategy DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
}

const (
	// DefaultDaemonSetUniqueLabelKey is the default key of the labels that is added
	// to daemon set pods to distinguish between old and new pod templates during
	// DaemonSet update.
	DefaultDaemonSetUniqueLabelKey string = "pod-template-hash"
)

// DaemonSetStatus represents the current status of a daemon set.
type DaemonSetStatus struct {
	// Note: Existing fields, including CurrentNumberScheduled, NumberMissscheduled,
	// DesiredNumberScheduled, NumberReady, and ObservedGeneration are ommitted in
	// this proposal.

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
   - If `OwnerReference` is implemented for DaemonSets, filter out pods that
     aren't controlled by this DaemonSet too
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
       "pod-template-hash" label. If the label value doesn't equal to the hash
       of `DaemonSetSpec.Template.Spec` and if `MaxInFlight` isn't reached, kill
       the pod and create one with new pod template.
       - `MaxInFlight` = the number of DaemonSet pods with new pod template (i.e.
         the same as `DaemonSetSpec.Template`) that have not become `Ready`
         (still being updated)
1. Cleanup, update DaemonSet status  

If DaemonSet Controller crashes during an update, it can still recover. 

### kubectl 

#### kubectl rollout 

Users can use `kubectl rollout` to monitor or manage DaemonSet updates, just
like Deployment rollouts. For example, 

- `kubectl rollout history`: to view history of DaemonSet updates. We use
  `PodTemplate` created by DaemonSets to store update history. 
- `kubectl rollout status`: to see the DaemonSet upgrade status 

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
  - Other alternatives: No-op, External
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
- DaemonSets can be updated while already being updated (i.e. rollover updates)
- Broken rollout can be rolled back (by applying old config)


## Future

In the future, we may:

- Support more DaemonSet update types
- Support rollback
- Allow user-defined unique label key 
- Add clean up policy for history
- Add minReadySeconds and make DaemonSet update strategy respect it
- Support pausing DaemonSets
- Support auto-rollback
