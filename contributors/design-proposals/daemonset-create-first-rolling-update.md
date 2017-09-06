## DaemonSet create-first rolling update strategy

DaemonSets currently support a "delete first, then add" rolling update
strategy.

This document presents a design for a new "add first, then delete" rolling
update strategy, called `SurgingRollingUpdate`. This strategy will create new
pods first, cleaning up the old pods once the new ones are running. As with
Deployments, users will use a `maxSurge` parameter to control how many nodes
are allowed to have >1 pod running during a rolling update.

### Example use case

For self-hosted Kubernetes DaemonSets are a natural fit for running the
scheduler and controller-manager on master nodes. However, since DaemonSets are
managed by the controller-manager, if `maxUnavailable` >= the number of masters
then a rolling update of the controller-manager would delete all the old pods
before the new pods were scheduled, leading to a control plane outage.

This feature is a blocker for single-master self-hosted Kubernetes, e.g. in
`kubeadm`, and we also believe it would be a generally useful feature.

### Implementation plan

A new `SurgingRollingUpdate` strategy will be added to the `DaemonSet` API
object. The data structures in `extensions/v1beta1` will be modified as
follows:

```go
type DaemonSetUpdateStrategy struct {
	// Type of daemon set update. Can be "RollingUpdate", "SurgingUpdate", or "OnDelete".
	// Default is OnDelete.
	// +optional
	Type DaemonSetUpdateStrategyType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`

	// Rolling update config params. Present only if type = "RollingUpdate".
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be. Same as Deployment `strategy.rollingUpdate`.
	// See https://github.com/kubernetes/kubernetes/issues/35345
	// +optional
	RollingUpdate *RollingUpdateDaemonSet `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate"`

	// Surging rolling update config params. Present only if type = "SurgingRollingUpdate".
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be. Same as Deployment `strategy.rollingUpdate`.
	// See https://github.com/kubernetes/kubernetes/issues/35345
	// +optional
	SurgingRollingUpdate *SurgingRollingUpdateDaemonSet `json:"surgingRollingUpdate,omitempty" protobuf:"bytes,3,opt,name=surgingRollingUpdate"`
}

...

const (
	// Replace the old daemons by new ones using rolling update i.e replace them on each node one after the other,
	// killing the old pod before starting the new one.
	RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"

	// Replace the old daemons by new ones using rolling update i.e replace them on each node one
	// after the other, creating the new pod and then killing the old one.
	SurgingRollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "SurgingRollingUpdate"

	// Replace the old daemons only when it's killed
	OnDeleteDaemonSetStrategyType DaemonSetUpdateStrategyType = "OnDelete"
)

...

// Spec to control the desired behavior of a daemon set surging rolling update.
type SurgingRollingUpdateDaemonSet struct {
	// The maximum number of DaemonSet pods that can be scheduled above the desired number of pods
	// during the update. Value can be an absolute number (ex: 5) or a percentage of the total number
	// of DaemonSet pods at the start of the update (ex: 10%). The absolute number is calculated from
	// the percentage by rounding up. This cannot be 0. The default value is 1. Example: when this is
	// set to 30%, at most 30% of the total number of nodes that should be running the daemon pod
	// (i.e. status.desiredNumberScheduled) can have 2 pods running at any given time. The update
	// starts by starting replacements for at most 30% of those DaemonSet pods. Once the new pods are
	// available it then stops the existing pods before proceeding onto other DaemonSet pods, thus
	// ensuring that at most 130% of the desired final number of DaemonSet  pods are running at all
	// times during the update.
	// +optional
	MaxSurge *intstr.IntOrString `json:"maxSurge,omitempty" protobuf:"bytes,2,opt,name=maxSurge"`
}
```

Then an additional hook for this strategy will be added to the part of the
reconciliation loop that runs when pod updates are needed. That hook will:

1. Determine how many nodes can have new pods scheduled, according to the
   number of extra nodes currently running more than one pod and the `maxSurge`
   parameter.
2. Determine how many nodes can have their previous-generation pods retired
   once the replacement pods are scheduled and running.
3. Instruct the controller-manager to add and delete pods on nodes based on the
   results of (1) and (2).

This is mechanically almost identical to the current `RollingUpdate` strategy,
with the key differences being the direct creation of new pods on nodes (the
`RollingUpdate` strategy only deletes pods up to the `maxUnavailable` threshold
and then waits for the main reconciler create new pods to replace them).

One minor complication is that the controller-manager's reconciliation loop
actively tries to enforce the one-pod-per-node invariant, and work is underway
to fix cases in which the invariant is not held
[kubernetes/kubernetes#50477](https://github.com/kubernetes/kubernetes/issues/50477).
However, this invariant will need to be relaxed when the `SurgingRollingUpdate`
strategy temporarily creates extra pods.

### Alternatives considered

The `maxSurge` parameter could have been added to the existing `RollingUpdate`
strategy (as is the case for Deployments). However, this would break validation
backwards compatibility, since the `maxUnavailable` parameter currently
requires `maxUnavailable` to be > 0, but this is not required with
`SurgingRollingUpdate`.

There are plans to create a `v1beta2` iteration of the apps API (enabled by
default in 1.9+), so it may be possible to add the `maxSurge` strategy to
`RollingUpdate` and deprecate `SurgingRollingUpdate` then, or make the change
now since the `v1beta2` API is incoming. However, to be conservative it was
decided to implement this as a separate strategy.

For more background see https://github.com/kubernetes/kubernetes/issues/48841.

### Considerations / questions

1. How are `hostPort`s handled?

They are not handled as part of this proposal. We can either:

   1. Attempt to determine a mechanism to handle this (unclear, perhaps best
      left as future work, i.e. prior to GA)
   2. Note in the documentation that this may cause port conflicts that prevent
      new pods from scheduling successfully (and therefore updates from
      completing successfully)
   3. Add a validation check that rejects the combination of `hostPort` and the
      `SurgingRollingUpdate` strategy.

2. How will the scheduler handle hostPort collisions?

The pods will not be scheduled. This is also true of Deployments (the new pods
can get stuck in pending). To verify, create a new Deployment that uses a
hostPort, has replicas = the number of workers (or use a selector), and set
`maxUnavailable` to 0. If you attempt to update this Deployment the update will
get stuck since the scheduler cannot find any nodes upon which to place the
pods.

It's noted that DaemonSets use `hostPort`s more frequently than Deployments,
and Deployment replica counts are ideally << the number of nodes, so this
issue will likely affect DaemonSets much more than Deployments.

3. How are other shared resources (e.g. local devices, GPUs, specific core
   types) handled?

This is a generalization of the `hostPort` question. I think that in this case
(especially for now) we should make a note in the API documentation at best
that running two pods that require scarce resources at once on a node may
(naturally) cause conflicts.

4. When can this strategy be used?

This strategy will be selectable by users at both DaemonSet creation time
and modification time. If a rolling update is occurring while the strategy
is changed the new strategy will take effect and reconciliation will occur.

In the case of `RollingUpdate` => `SurgingRollingUpdate`:

- New pods will be created up to `maxSurge` (see below).
- The controller will stop deleting the old pods *before* new pods are running,
  instead waiting for the surging pods to become ready.

In the case of `SurgingRollingUpdate` => `RollingUpdate`:

- Surging pods will be killed by the controller (the current behavior in the
  controller is to retain the oldest pod per node when multiple are present).
- The controller will begin deleting the old pods up to `maxUnavailable` and
  then let the controller reschedule them.

5. Is this strategy enabled by default? Feature gated? Alpha? Beta?

#TODO(luxas)
