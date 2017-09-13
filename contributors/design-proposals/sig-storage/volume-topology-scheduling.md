# Volume Topology-aware Scheduling

Authors: @msau42

This document presents a detailed design for making the default Kubernetes
scheduler aware of volume topology constraints, and making the
PersistentVolumeClaim (PVC) binding aware of scheduling decisions.


## Goals
* Allow a Pod to request one or more topology-constrained Persistent
Volumes (PV) that are compatible with the Pod's other scheduling
constraints, such as resource requirements and affinity/anti-affinity
policies.
* Support arbitrary PV topology constraints (i.e. node,
rack, zone, foo, bar).
* Support topology constraints for statically created PVs and dynamically
provisioned PVs.
* No scheduling latency performance regression for Pods that do not use
topology-constrained PVs.


## Non Goals
* Fitting a pod after the initial PVC binding has been completed.
    * The more constraints you add to your pod, the less flexible it becomes
in terms of placement.  Because of this, tightly constrained storage, such as
local storage, is only recommended for specific use cases, and the pods should
have high priority.
    * Pod priorities could handle this scenario.  PV taints/tolerations could
also handle this case.
* Binding decision considering scheduling constraints from two or more pods
sharing the same PVC.
    * The scheduler itself only handles one pod at a time.  It’s possible the
two pods may not run at the same time either, so there’s no guarantee that you
will know both pod’s requirements at once.
    * For two+ pods simultaneously sharing a PVC, this scenario may require an
operator to schedule them together.
    * For two+ pods non-simultaneously sharing a PVC, this scenario could be
handled by pod priorities.


## Problem
Volumes can have topology constraints that restrict the set of nodes that the
volume is accessible from.  For example, a GCE PD can only be accessed from a
single zone, and a local disk can only be accessible from a single node.  In the
future, there could be other topology constraints, such as rack or region.

A pod that uses such a volume must be scheduled to a node that fits within the
volume’s topology constraints.  In addition, a pod can have further constraints
and limitations, such as the pod’s resource requests (cpu, memory, etc), and
pod/node affinity and anti-affinity policies.

Currently, the process of binding and provisioning volumes are done before a pod
is scheduled.  Therefore, it cannot take into account any of the pod’s other
scheduling constraints.  This makes it possible for the PV controller to bind a
PVC to a PV or provision a PV with constraints that can make a pod unschedulable.

### Examples
* In multizone clusters, the PV controller has a hardcoded heuristic to provision
PVCs for StatefulSets spread across zones.  If that zone does not have enough
cpu/memory capacity to fit the pod, then the pod is stuck in pending state because
its volume is bound to that zone.
* Local storage exasperates this issue.  The chance of a node not having enough
cpu/memory is higher than the chance of a zone not having enough cpu/memory.
* Local storage PVC binding does not have any node spreading logic.  So local PV
binding will very likely conflict with any pod anti-affinity policies if there is
more than one local PV on a node.
* A pod may need multiple PVCs.  As an example, one PVC can point to a local SSD for
fast data access, and another PVC can point to a local HDD for logging.  Since PVC
binding happens without considering if multiple PVCs are related, it is very likely
for the two PVCs to be bound to local disks on different nodes, making the pod
unschedulable.
* For multizone clusters and deployments requesting multiple dynamically provisioned
zonal PVs, each PVC Is provisioned independently, and is likely to provision each PV
In different zones, making the pod unschedulable.

To solve the issue of initial volume binding and provisioning causing an impossible
pod placement, volume binding and provisioning should be more tightly coupled with
pod scheduling.


## Background
In 1.7, we added alpha support for local PVs (TODO: link) with node affinity.
You can specify a PV object with node affinity, and if a pod is using such a PV,
the scheduler will evaluate the PV node affinity in addition to the other
scheduling predicates.  So far, the PV node affinity only influences pod
scheduling once the PVC is already bound.  The initial PVC binding decision was
unchanged.  This proposal addresses the initial PVC binding decision.


## Design
The design can be broken up into a few areas:
* User-facing API to invoke new behavior
* Integrating PV binding with pod scheduling
* Binding multiple PVCs as a single transaction
* Recovery from kubelet rejection of pod
* Making dynamic provisioning topology-aware

For the alpha phase, only the user-facing API and PV binding and scheduler
integration are necessary.  The remaining areas can be handled in beta and GA
phases.

### User-facing API
This new binding behavior will apply to all unbound PVCs and can be controlled
through install-time configuration passed to the scheduler and controller-manager.

In alpha, this configuration can be controlled by a feature gate,
`VolumeTopologyScheduling`.

#### Alternative
An alternative approach is to only trigger the new behavior only for volumes that
have topology constraints.  A new annotation can be added to the StorageClass to
indicate that its volumes have constraints and will need to use the new
topology-aware scheduling logic.

```
// Value is “true” or “false”.  Default is false.
const AlphaVolumeTopologySchedulingAnnotation = “volume.alpha.kubernetes.io/topology-scheduling”
```

While this approach can let us introduce the new behavior gradually, it has a
few downsides:
* StorageClass will be required to use this new logic, even if dynamic
provisioning is not used (in the case of local storage).
* We have to maintain two different paths for volume binding.
* We will be depending on the storage admin to correctly configure the
StorageClasses.

For those reasons, we prefer to switch the behavior for all unbound PVCs and
clearly communicate these changes to the community before changing the behavior
by default in GA.

### Integrating binding with scheduling
For the alpha phase, the focus is on static provisioning of PVs to support
persistent local storage.

TODO: flow chart

The proposed new workflow for volume binding is:
1. Admin statically creates PVs and/or StorageClasses.
2. User creates unbound PVC and there are no prebound PVs for it.
3. **NEW:** PVC binding and provisioning is delayed until a pod is created that
references it.
4. User creates a pod that uses the PVC.
5. Pod starts to get processed by the scheduler.
6. **NEW:** A new predicate function, called MatchUnboundPVCs, will look at all of
a Pod’s unbound PVCs, and try to find matching PVs for that node based on the
PV topology.  If there are no matching PVs, then it checks if dynamic
provisioning is possible for that node.
7. **NEW:** The scheduler continues to evaluate priorities.  A new priority
function, called PrioritizeUnboundPVCs, will get the PV matches per PVC per
node, and compute a priority score based on various factors.
8. **NEW:** After evaluating all the existing predicates and priorities, the
scheduler will pick a node, and call a new assume function, AssumePVCs,
passing in the Node.  The assume function will check if any binding or
provisioning operations need to be done.  If so, it will update the PV cache to
mark the PVs with the chosen PVCs.
9. **NEW:** If PVC binding or provisioning is required, we do NOT AssumePod.
Instead, a new bind function, BindPVCs, will be called asynchronously, passing
in the selected node.  The bind function will prebind the PV to the PVC, or
trigger dynamic provisioning, and then wait until the PVCs are bound
successfully or encounter failure.  Then, it always sends the Pod through the
scheduler again for reasons explained later.
10. When a Pod makes a successful scheduler pass once all PVCs are bound, the
scheduler assumes and binds the Pod to a Node.
11. Kubelet starts the Pod.

This new workflow will have the scheduler handle unbound PVCs by choosing PVs
and prebinding them to the PVCs.  The PV controller completes the binding
transaction, handling it as a prebound PV scenario.

Prebound PVCs and PVs will still immediately be bound by the PV controller.

One important point to note is that for the alpha phase, manual recovery is
required in following error conditions:
* A Pod has multiple PVCs, and only a subset of them successfully bind.
* The scheduler chose a PV and prebound it, but the PVC could not be bound.

In both of these scenarios, the primary cause for these errors is if a user
prebinds other PVCs to a PV that the scheduler chose.  Some workarounds to
avoid these error conditions are to:
* Prebind the PV instead.
* Separate out volumes that the user prebinds from the volumes that are
available for the system to choose from by StorageClass.

In addition, this new design will cause functional regression for the
following use cases:
* While a Pod is running, a user deletes the PVC and PV, and then recreates them.
    * With the current behavior, the PVC will rebind to the same or different PV.
    * With this new behavior, the PVC will not rebind because binding would only
be triggered when the Pod has to go through scheduling.  This could cause
problems with Pod volume cleanup because the PVC will no longer be bound.
    * But deleting the PVC and PV (especially both) while a Pod is still using it
is also dangerous and should be prevented in the first place.  There could be
serious problems if you recreate the PVC and it rebinds to a different PV.  If
the PV had a reclaim policy of Delete, then deleting the PVC could cause the
underlying disk to get deleted while the Pod is still using it.


#### PV Controller Changes
When the feature gate is enabled, the PV controller needs to skip binding and
provisioning all unbound PVCs that don’t have prebound PVs, and let it come
through the scheduler path.  This is the block in `syncUnboundClaim` where
`claim.Spec.VolumeName == “”` and there is no PV that is prebound to the PVC.

No other state machine changes are required.  The PV controller continues to
handle the remaining scenarios without any change.

The methods to find matching PVs for a claim, prebind PVs, and to invoke and
handle dynamic provisioning need to be refactored for use by the new scheduler
functions.

#### Scheduler Changes

##### Predicate
A new predicate function checks all of a Pod's unbound PVCs can be satisfied
by existing PVs or dynamically provisioned PVs that are
topologically-constrained to the Node.
```
MatchUnboundPVCs(pod *v1.Pod, node *v1.Node) (canBeBound bool, err error)
```
1. If all the Pod’s PVCs are bound, return true.
2. Otherwise try to find matching PVs for all of the unbound PVCs in order of
decreasing requested capacity.
3. Walk through all the PVs.
4. Find best matching PV for the PVC where PV topology is satisfied by the Node.
5. Temporarily cache this PV in the PVC object, keyed by Node, for fast
processing later in the priority and bind functions.
6. Return true if all PVCs are matched.
7. If there are still unmatched PVCs, check if dynamic provisioning is possible.
For this alpha phase, the provisioner is not topology aware, so the predicate
will just return true if there is a provisioner specified (internal or external).
8. Otherwise return false.

TODO: caching format and details

##### Priority
After all the predicates run, there is a reduced set of Nodes that can fit a
Pod. A new priority function will rank the remaining nodes based on the
unbound PVCs and their matching PVs.
```
PrioritizeUnboundPVCs(pod *v1.Pod, filteredNodes HostPriorityList) (rankedNodes HostPriorityList, err error) 
```
1. For each Node, get the cached PV matches for the Pod’s PVCs.
2. Compute a priority score for the Node using the following factors:
    1. How close the PVC’s requested capacity and PV’s capacity are.
    2. Matching static PVs is preferred over dynamic provisioning.

TODO (beta): figure out weights and exact calculation

##### Assume
Once all the predicates and priorities have run, then the scheduler picks a
Node.  Then we can bind or provision PVCs for that Node.  For better scheduler
performance, we’ll assume that the binding will likely succeed, and update the
PV cache first.  Then the actual binding API update will be made
asynchronously, and the scheduler can continue processing other Pods.

For the alpha phase, the AssumePVCs function will be directly called by the
scheduler.  We’ll consider creating a generic scheduler interface in a
subsequent phase.

```
AssumePVCs(pod *v1.Pod, node *v1.Node) (pvcBindingRequired bool, err error)
```
1. If all the Pod’s PVCs are bound, return false.
2. For static PV binding:
    1. Get the cached matching PVs for the PVCs on that Node.
    2. Validate the actual PV state.
    3. Mark PV.ClaimRef in the PV cache.
3. For in-tree and external dynamic provisioning:
    1. Nothing.
4. Return true.


##### Bind
If AssumePVCs returns pvcBindingRequired, then the BindPVCs function is called
as a go routine.  Otherwise, we can continue with assuming and binding the Pod
to the Node.

For the alpha phase, the BindUnboundPVCs function will be directly called by the
scheduler.  We’ll consider creating a generic scheduler interface in a subsequent
phase.

```
BindUnboundPVCs(pod *v1.Pod, node *v1.Node) (err error)
```
1. For static PV binding:
    1. Prebind the PV by updating the `PersistentVolume.ClaimRef` field.
    2. If the prebind fails, revert the cache updates.
    3. Otherwise, wait for the PVCs to be bound, PVC/PV object is deleted, or
PV.ClaimRef field is cleared
2. For in-tree dynamic provisioning:
    1. Make the provision call, which will create a new PV object that is
prebound to the PVC.
    2. If provisioning is successful, wait for the PVCs to be bound, PVC/PV
object is deleted, or PV.ClaimRef field is cleared
3. For external provisioning:
    1. Set the annotation for external provisioners.
4. Send Pod back through scheduling, regardless of success or failure.
    1. In the case of success, we need one more pass through the scheduler in
order to evaluate other volume predicates that require the PVC to be bound, as
described below.
    2. In the case of failure, we want to retry binding/provisioning.

Note that for external provisioning, we do not wait for the PVCs to be bound, so
the Pod will be sent through scheduling repeatedly until the PVCs are bound.
This is because there is no function call for external provisioning, so if we
did wait, we could be waiting forever for the PVC to bind.  It’s possible that
in the meantime, a user could create a PV that satisfies the PVC and doesn’t
need the external provisioner anymore.

TODO: pv controller has a high resync frequency, do we need something similar
for the scheduler too

##### Pod preemption considerations
The MatchUnboundPVs predicate does not need to be re-evaluated for pod
preemption.  Preempting a pod that uses a PV will not free up capacity on that
node because the PV lifecycle is independent of the Pod’s lifecycle.

##### Other scheduler predicates
Currently, there are a few existing scheduler predicates that require the PVC
to be bound.  The bound assumption needs to be changed in order to work with
this new workflow.

TODO: how to handle race condition of PVCs becoming bound in the middle of
running predicates?  Do we need to impose ordering on the predicates?

###### Max PD Volume Count Predicate
This predicate checks the maximum number of PDs per node is not exceeded.  It
needs to be integrated into the binding decision so that we don’t bind or
provision a PV if it’s going to cause the node to exceed the max PD limit.  But
until it is integrated, we need to make one more pass in the scheduler after all
the PVCs are bound.  The current copy of the predicate in the default scheduler
has to remain to account for the already-bound volumes.

###### Volume Zone Predicate
This predicate makes sure that the zone label on a PV matches the zone label of
the node.  If the volume is not bound, this predicate can be ignored, as the
binding logic will take into account zone constraints on the PV.

However, this assumes that zonal PVs like GCE PDs and AWS EBS have been updated
to use the new PV topology specification, which is not the case as of 1.8.  So
until those plugins are updated, the binding and provisioning decisions will be
topology-unaware, and we need to make one more pass in the scheduler after all
the PVCs are bound.

This predicate needs to remain in the default scheduler to handle the
already-bound volumes using the old zonal labeling.  It can be removed once that
mechanism is deprecated and unsupported.

###### Volume Node Predicate
This is a new predicate added in 1.7 to handle the new PV node affinity.  It
evaluates the node affinity against the node’s labels to determine if the pod
can be scheduled on that node.  If the volume is not bound, this predicate can
be ignored, as the binding logic will take into account the PV node affinity.

#### Performance and Optimizations
Let:
* N = number of nodes
* V = number of all PVs
* C = number of claims in a pod

C is expected to be very small (< 5) so shouldn’t factor in.

The current PV binding mechanism just walks through all the PVs once, so its
running time O(V).

Without any optimizations, the new PV binding mechanism has to run through all
PVs for every node, so its running time is O(NV).

A few optimizations can be made to improve the performance:

1. Optimizing for PVs that don’t use node affinity (to prevent performance
regression):
    1. Index the PVs by StorageClass and only search the PV list with matching
StorageClass.
    2. Keep temporary state in the PVC cache if we previously succeeded or
failed to match PVs, and if none of the PVs have node affinity.  Then we can
skip PV matching on subsequent nodes, and just return the result of the first
attempt.
2. Optimizing for PVs that have node affinity:
    1. When a static PV is created, if node affinity is present, evaluate it
against all the nodes.  For each node, keep an in-memory map of all its PVs
keyed by StorageClass.  When finding matching PVs for a particular node, try to
match against the PVs in the node’s PV map instead of the cluster-wide PV list.

For the alpha phase, the optimizations are not required.  However, they should
be required for beta and GA.

#### Packaging
The new bind logic that is invoked by the scheduler can be packaged in a few
ways:
* As a library to be directly called in the default scheduler
* As a scheduler extender

We propose taking the library approach, as this method is simplest to release
and deploy.  Some downsides are:
* The binding logic will be executed using two different caches, one in the
scheduler process, and one in the PV controller process.  There is the potential
for more race conditions due to the caches being out of sync.
* Refactoring the binding logic into a common library is more challenging
because the scheduler’s cache and PV controller’s cache have different interfaces
and private methods.

##### Extender cons
However, the cons of the extender approach outweighs the cons of the library
approach.

With an extender approach, the PV controller could implement the scheduler
extender HTTP endpoint, and the advantage is the binding logic triggered by the
scheduler can share the same caches and state as the PV controller.

However, deployment of this scheduler extender in a master HA configuration is
extremely complex.  The scheduler has to be configured with the hostname or IP of
the PV controller.  In a HA setup, the active scheduler and active PV controller
could run on the same, or different node, and the node can change at any time.
Exporting a network endpoint in the controller manager process is unprecedented
and there would be many additional features required, such as adding a mechanism
to get a stable network name, adding authorization and access control, and
dealing with DDOS attacks and other potential security issues.  Adding to those
challenges is the fact that there are countless ways for users to deploy
Kubernetes.

With all this complexity, the library approach is the most feasible in a single
release time frame, and aligns better with the current Kubernetes architecture.

#### Other Alternatives Considered

##### One scheduler function
An alternative design considered was to do the predicate, priority and bind
functions all in one function at the end right before Pod binding, in order to
reduce the number of passes we have to make over all the PVs.  However, this
design does not work well with pod preemption.  Pod preemption needs to be able
to evaluate if evicting a lower priority Pod will make a higher priority Pod
schedulable, and it does this by re-evaluating predicates without the lower
priority Pod.

If we had put the MatchUnboundPVCs predicate at the end, then pod preemption
wouldn’t have an accurate filtered nodes list, and could end up preempting pods
on a Node that the higher priority pod still cannot run on due to PVC
requirements.  For that reason, the PVC binding decision needs to be have its
predicate function separated out and evaluated with the rest of the predicates.

##### Pull entire PVC binding into the scheduler
The proposed design only has the scheduler initiating the binding transaction
by prebinding the PV.  An alternative is to pull the whole two-way binding
transaction into the scheduler, but there are some complex scenarios that
scheduler’s Pod sync loop cannot handle:
* PVC and PV getting unexpectedly unbound or lost
* PVC and PV state getting partially updated
* PVC and PV deletion and cleanup

Handling these scenarios in the scheduler’s Pod sync loop is not possible, so
they have to remain in the PV controller.

##### Keep all PVC binding in the PV controller
Add a new “node” annotation to the PVC object and have the scheduler update it
when it schedules the pod to a node.  A new scheduling predicate is still
needed to filter and match the PVs (but not actually bind).

Major downsides of this approach include:
* Requires PV controller to change its sync loop to operate on pods, in order
to handle the multiple PVCs in a pod scenario.  This is a potentially big change
that would be hard to keep separate and feature-gated from the current PV logic.
* Needs new coordination between PV controller and kubelet to reject a pod if
binding fails to send the pod back to scheduler.
* Initial filtering (matching) and binding are done asynchronously so there is a
higher chance of binding failures.  As an example, we could match PV1 for both
pod1 and pod2, but only pod1 will be able to bind successfully, and pod2 has to
be sent back to scheduling.
* More race conditions and higher chance of binding failure because of 3
different components operating independently (scheduler, PV controller, kubelet).

##### Move PVC binding to kubelet
Looking into the future, with the potential for NUMA-aware scheduling, you could
have a sub-scheduler on each node to handle the pod scheduling within a node.  It
could make sense to have the volume binding as part of this sub-scheduler, to make
sure that the volume selected will have NUMA affinity with the rest of the
resources that the pod requested.

However, there are potential security concerns because kubelet would need to see
unbound PVs in order to bind them.  For local storage, the PVs could be restricted
to just that node, but for zonal storage, it could see all the PVs in that zone.

In addition, the sub-scheduler is just a thought at this point, and there are no
concrete proposals in this area yet.

### Binding multiple PVCs in one transaction
For the alpha phase, this is not required.  Since the scheduler is serialized, a
partial binding failure should be a rare occurrence.  Manual recovery will need
to be done for those rare failures.

### Recovering from kubelet rejection of pod
Again, for alpha phase, manual intervention will be required.  For beta, we’ll
have to consider some sort of rollback mechanism for more robustness.  Perhaps
this rollback mechanism can also be used for binding multiple PVCs in one
transaction.

### Making dynamic provisioning topology aware
TODO (beta): Design details

For alpha, we are not focusing on this use case.  But it should be able to
follow the new workflow closely with some modifications.
* The FindUnboundPVCs predicate function needs to get provisionable capacity per
topology dimension from the provisioner somehow.
* The PrioritizeUnboundPVCs priority function can add a new priority score factor
based on available capacity per node.
* The BindUnboundPVCs bind function needs to pass in the node to the provisioner.
The internal and external provisioning APIs need to be updated to take in a node
parameter.


## Testing

### E2E tests
* StatefulSet, replicas=3, specifying pod anti-affinity
    * Positive: Local PVs on each of the nodes
    * Negative: Local PVs only on 2 out of the 3 nodes
* StatefulSet specifying pod affinity
    * Positive: Multiple local PVs on a node
    * Negative: Only one local PV available per node
* Multiple PVCs specified in a pod
    * Positive: Enough local PVs available on a single node
    * Negative: Not enough local PVs available on a single node
* Fallback to dynamic provisioning if unsuitable static PVs
* Existing PV tests need to be modified, if alpha, to check for PVC bound after
pod has started.
* Existing dynamic provisioning tests may need to be disabled for alpha if they
don’t launch pods.  Or extended to launch pods.
* Existing prebinding tests should still work.

### Unit tests
* All PVCs found a match on first node.  Verify match is best suited based on
capacity.
* All PVCs found a match on second node.  Verify match is best suited based on
capacity.
* Only 2 out of 3 PVCs have a match.
* Priority scoring doesn’t change the given priorityList order.
* Priority scoring changes the priorityList order.
* Don’t match PVs that are prebound


## Implementation Plan

### Alpha
* New feature gate for volume topology scheduling
* Refactor PV controller methods into a common library
* PV controller: Disable binding unbound PVCs if feature gate is set
* Predicate: Filter nodes and find matching PVs
* Predicate: Check if provisioner exists for dynamic provisioning
* Update existing predicates to skip unbound PVC
* Bind: Trigger PV binding
* Bind: Trigger dynamic provisioning
* Tests: Refactor all the tests that expect a PVC to be bound before scheduling
a Pod (only if alpha is enabled)

### Beta
* Scheduler cache: Optimizations for no PV node affinity
* Priority: capacity match score
* Bind: Handle partial binding failure
* Plugins: Convert all zonal volume plugins to use new PV node affinity (GCE PD,
AWS EBS, what else?)
* Make dynamic provisioning topology aware

### GA
* Predicate: Handle max PD per node limit
* Scheduler cache: Optimizations for PV node affinity
* Handle kubelet rejection of pod


## Open Issues
* Can generic device resource API be leveraged at all?  Probably not, because:
    * It will only work for local storage (node specific devices), and not zonal
storage.
    * Storage already has its own first class resources in K8s (PVC/PV) with an
independent lifecycle.  The current resource API proposal does not have an a way to
specify identity/persistence for devices.
* Will this be able to work with the node sub-scheduler design for NUMA-aware
scheduling?
    * It’s still in a very early discussion phase.
