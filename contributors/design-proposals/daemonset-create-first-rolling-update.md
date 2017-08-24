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

This feature is a blocker for supporting self-hosted Kubernetes in `kubeadm`.

### Implementation plan

A new `SurgingRollingUpdate` strategy will be added to the `DaemonSet` API
object. Then an additional hook for this strategy will be added to the part of
the reconciliation loop that runs when pod updates are needed. That hook will:

1. Determine how many new pods can be created, according to the number of extra
   pods currently running and the `maxSurge` parameter.
2. Determine how many previous-generation pods can be retired once their
   replacement pods are scheduled and running.
3. Instruct the controller-manager to add and delete pods based on the results
   of (1) and (2).

This is mechanically almost identical to the current `RollingUpdate` strategy,
with the key differences being the direct creation of new pods (the
`RollingUpdate` strategy only deletes pods up to the `maxUnavailable` threshold
and then waits for the main reconciler create new pods to replace them).

One minor complication is that the controller-manager's reconciliation loop
actively enforces the one-pod-per-node invariant. This invariant will need to
be relaxed when the `SurgingRollingUpdate` strategy temporarily creates extra
pods.

### Alternatives considered

The `maxSurge` parameter could have been added to the existing `RollingUpdate`
strategy (as is the case for Deployments). However, this would break backwards
compatibility, since the `maxUnavailable` parameter currently requires
`maxUnavailable` to be > 0, but this is not required with
`SurgingRollingUpdate`.

There are plans to create a `v1beta2` iteration of the extensions API, so it
may be possible to add the `maxSurge` strategy to `RollingUpdate` and deprecate
`SurgingRollingUpdate` then.

For more background see https://github.com/kubernetes/kubernetes/issues/48841.

### Considerations / questions

1. How are `hostPort`s handled? 

They are not handled as part of this proposal. We can either:

  1. Attempt to determine a mechanism to handle this (unclear)
  2. Note in the documentation that this may cause port conflicts that prevent
    new pods from scheduling successfully (and therefore updates from completing
    successfully)
  3. Add a validation check that rejects the combination of `hostPort` and the
    `SurgingRollingUpdate` strategy.

2. How will the scheduler handle hostPort collisions?

TODO(diegs): investigate this.

3. How are other shared resources (e.g. local devices, GPUs, specific core
   types) handled?

This is a generalization of the `hostPort` question. I think that in this case
(especially for now) we should make a note in the API documentation at best
that running two pods that require scarce resources at once on a node may
(naturally) cause conflicts.
