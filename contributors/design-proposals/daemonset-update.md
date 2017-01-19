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

- Users can trigger a rolling update of DaemonSet at a controlled speed, which
  is achieved by:
  - Only a certain number of DaemonSet pods can be down at the same time during
    an update
  - A DaemonSet pod needs to be ready for a specific amount of time before it's
    considered up
- Users can monitor the status of a DaemonSet update (e.g. the number of pods
  that are updated and healthy)
- A broken DaemonSet update should not continue, but one can still update the
  DaemonSet again to fix it
- Users should be able to update a DaemonSet even during an ongoing DaemonSet
  upgrade -- in other words, rollover (e.g. update the DaemonSet to fix a broken
  DaemonSet update)
- Users should be able to view the history of previous DaemonSet updates 
- Users can figure out the revision of a DaemonSet's pod (e.g. which version is
  this DaemonSet pod?)

Here are some potential requirements that haven't been covered by this proposal:

- DaemonSet should provide at-most-one guarantee per node (i.e. at most one pod
  from a DaemonSet can exist on a node at any time)
- Uptime is critical for each pod of a DaemonSet during an upgrade (e.g. the time
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
	// The maximum number of DaemonSet pods that can be unavailable during
	// the update. Value can be an absolute number (ex: 5) or a percentage of total
	// number of DaemonSet pods at the start of the update (ex: 10%). Absolute
	// number is calculated from percentage by rounding up.
	// This must be greater than 0.
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
	// Note: Existing fields, including Selector and Template are ommitted in
	// this proposal.  

	// Update strategy to replace existing DaemonSet pods with new pods.
	UpdateStrategy DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`

	// Minimum number of seconds for which a newly created DaemonSet pod should
	// be ready without any of its container crashing, for it to be considered
	// available. Defaults to 0 (pod will be considered available as soon as it
	// is ready).
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// The number of old PodTemplates to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 3. 
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
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
DaemonSets on the apiserver. 

For each pending DaemonSet updates, it will:

1. Find all pods whose labels are matched by DaemonSet `.spec.selector`. 
   - If `OwnerReference` is implemented for DaemonSets, filter out pods that
     aren't controlled by this DaemonSet too
1. Find all nodes that should run these pods created by this DaemonSet.
1. Find existing PodTemplates whose labels are matched by DaemonSet
   `.spec.selector`
   - Sort those PodTemplates by creation timestamp and only retain at most
     `.spec.revisionHistoryLimit` latest PodTemplates (remove the rest)
   - Find the PodTemplate whose `.template` is the same as DaemonSet
     `.spec.template`. If not found, create a new PodTemplate from DaemonSet
     `.spec.template`:
     - The name will be `<DaemonSet-Name>-<Hash-of-pod-template>`
     - PodTemplate `.metadata.labels` will have a "pod-template-hash" label,
       value be the hash of PodTemplate `.template` (note: don't include the
       "pod-template-hash" label when calculating hash)
     - PodTemplate `.metadata.annotations` will be copied from DaemonSet 
       `.metadata.annotations`
1. Create daemon pods on nodes when they should have those pods running but not
   yet, with the same "pod-template-hash" labels in daemon pods' `Metadata.Labels`.
   Otherwise, delete running daemon pods that shouldn't be running on nodes. 
1. Check `DaemonSetUpdateStrategy`:
   - If `OnDelete`: do nothing
   - If `RollingUpdate`:
     - From all nodes that should run daemon pods, check the daemon pod's
       "pod-template-hash" label. If the label value doesn't equal to the hash
       of `DaemonSetSpec.Template.Spec` and if `MaxUnavailable` isn't reached,
       kill the pod. Here, we kill unavailable pods first. 
       - `MaxUnavailable` >= the total number of DaemonSet pods that have not
          become `Ready` for `MinReadySeconds`
1. Cleanup, update DaemonSet status  

If DaemonSet Controller crashes during an update, it can still recover. 

### kubectl 

#### kubectl rollout 

Users can use `kubectl rollout` to monitor or manage DaemonSet updates, just
like Deployment rollouts. For example, 

- `kubectl rollout history daemonset/<DaemonSet-Name>`: to view the history of
  DaemonSet updates. We use `PodTemplate` created by DaemonSets to store update
  history. 
- `kubectl rollout status daemonset/<DaemonSet-Name>`: to see the DaemonSet
  upgrade status 

## Updating DaemonSets mid-way

Users can update an updated DaemonSet before its rollout completes.
In this case, the existing daemon pods will not continue rolling out and the new
one will begin rolling out.


## Deleting DaemonSets

Deleting a DaemonSet (with cascading) will delete all its pods and PodTemplates. 


## DaemonSet Strategies

DaemonSetStrategy specifies how the new daemon pods should replace existing ones.
To begin with, we will support 2 types:

* On delete: Do nothing, until existing daemon pods are killed (for backward
  compatibility).
  - Other alternative names: No-op, External
* Rolling update: We gradually kill existing ones while creating the new one.


## Tests

- Updating a RollingUpdate DaemonSet will trigger updates to its daemon pods. 
- Updating an OnDelete DaemonSet will not trigger updates, until the pods are
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


## Future Plans

In the future, we may:

- Implement at-most-one and/or at-least-one guarantees for DaemonSets (i.e. at
  most/at least one pod from a DaemonSet can exist on a node at any time)
  - At-most-one would use a deterministic name for the pod (e.g. use node name
    as daemon pod name suffix)
- Support use cases where uptime is critical for each pod of a DaemonSet during
  an upgrade
  - One approach is to use dummy pods to pre-pull images to reduce down time 
- Support use cases that each DaemonSet pod can still fit on the node after
  being updated (unless it becomes larger than the node). Some possible
  approaches include:
  - Make DaemonSet pods (daemons) have higher priority than non-daemons, and
    kubelet will evict pods with lower priority to make room for higher priority
    ones 
  - The DaemonSet controller will evict pods when daemons can't fit on the node
  - The DaemonSet controller will cordon the node before upgrading the daemon on
    it, and uncordon the node once it's done
- Support use cases that require the node to be drained before the daemons on it 
  can updated (e.g. logging daemons)
  - The DaemonSet controller will drain the node before upgrading the daemon on
    it, and uncordon the node once it's done 
- Make DaemonSets admin-only resources (admon = people who manage nodes). Some
  possible approaches include:
  - Remove namespace from DaemonSets (DaemonSets become node-level resources)
  - Modify RBAC bootstrap policy to make DaemonSets admin-only 
  - Delegation or impersonation 
- Make PodTemplate, which is used to store DaemonSet history, an admin-only or
  read only resource.
- Support more DaemonSet update strategies 
- Support rolling back DaemonSets 
  - `kubectl rollout undo daemonset/<DaemonSet-Name>` to roll back 
- Implement a subresource for DaemonSet history (e.g. `daemonsets/foo/history`)
  that summarize the information in the history 
- Allow user-defined DaemonSet unique label key
- Support pausing DaemonSet rolling update
- Support auto-rollback DaemonSets
