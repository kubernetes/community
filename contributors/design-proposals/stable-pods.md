# Introduce Stable Pods

**Author**: @kargakis

**Status**: Proposal

## Abstract
Today, Deployments, ReplicaSets, and DaemonSets use MinReadySeconds in order to include
an additional delay on top of readiness checks and facilitate more robust rollouts. The
ReplicaSet controller decides how many Ready Pods a ReplicaSet runs for at least
MinReadySeconds and that number is reflected in the ReplicaSetStatus as AvailableReplicas.
The Deployment controller when scaling down old ReplicaSets, will not proceed if the
minimum number of Pods that are required to run are not Ready for at least MinReadySeconds.
If MinReadySeconds is not specified then the Pods are included in AvailableReplicas as
soon as they are Ready. 

Problems that have been identified so far with the current state of things:
1. A Pod is marked Ready by the kubelet as soon as it passes its readiness check. The
ReplicaSet controller runs as part of master and estimates when a Pod is stable by
comparing the time the Pod became Ready (as seen by the kubelet) with MinReadySeconds.
[Clock-skew between master and nodes will affect these availability checks](https://github.com/kubernetes/kubernetes/issues/29229). 
2. PodDisruptionBudget is working with Ready Pods and has no notion of MinReadySeconds
when used by workload controllers.
3. The workload controllers cannot take advantage of readiness checks when checking the
state of their underlying Pods.
4. It's very hard to reason about which pods are not stable just by looking in the status
of a controller, because it only reports aggregated data. Imagine a ReplicaSet with 200
replicas, and only 3 of them are unstable.

## Design
All problems above can be solved by moving MinReadySeconds in the PodSpec. Once a kubelet
observes that a Pod has been Ready for at least MinReadySeconds without any of its
containers crashing and passes its readiness check successfully for another time, it will
update the PodStatus with an Stable Condition set to Status=True. Higher-level
orchestrators running on different machines such as the ReplicaSet or the
PodDisruptionBudget controller will merely need to look at the Stable condition that is set
in the status of a Pod.

Services should not be affected by the implementation of Stable Pods. Endpoints should
continue to be added or removed when a Pod transitions to or from Ready.

### API changes

A new field is proposed in the PodSpec:
```go
	// Minimum number of seconds for which a newly created Pod should be Ready
	// without any of its container crashing, for it to be considered Stable.
	// Defaults to 0 (pod will be considered Stable as soon as it is Ready)
	// +optional
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
```
and a new PodConditionType:
```go
	// PodStable is added in a Pod after it has been Ready for at least MinReadySeconds.
	// The Pod should already have an Endpoint and serve requests, this condition lets
	// higher-level orchestrators know that the Pod is stable after some amount of time
	// without having any of its containers crashed.
	PodStable PodConditionType = "Stable"
```

The reason for adding a separate Condition instead of extending the existing Ready
condition is mainly for allowing traffic in Pods. Users can have robust rollouts if their
applications expose metrics which are then fed into readiness checks.

Additionally:
* Deployments/ReplicaSets/DaemonSets already use MinReadySeconds in their Spec so we
should probably deprecate those fields in favor of the field in the PodTemplateSpec and
remove them in a future version. 
* All the aforementioned controllers will not propagate MinReadySeconds from their Spec
down to their PodTemplateSpec because that may lead in differences in the PodTemplateSpec
between a Deployment and a ReplicaSet resulting in new rollouts. If MinReadySeconds is
specified both in the Spec and PodTemplateSpec for a controller, the field in the
PodTemplateSpec will stomp the field in the Spec. If it is specified only in the Spec,
Pods can be created using it without updating the controller's PodTemplateSpec.

### kubelet changes
For a Pod that specifies MinReadySeconds, kubelet will need to rerun the readiness check
after MinReadySeconds and also check if any of its containers has crashed in the meantime.
If everything is ok, the kubelet will switch the Pod's Stable condition to Status=True.
Pods that don't specify MinReadySeconds will be considered Stable as soon as they are
Ready.

### Controller manager changes
The ReplicaSet and DaemonSet controllers will create new Pods by setting MinReadySeconds in
the PodSpec if it is specified either in their Spec or PodTemplateSpec. The value in the
PodTemplateSpec should stomp the value in the Spec. For Pods that do not specify
MinReadySeconds (those running on old kubelets), the controllers will continue to use the
current approach for estimating stability. Eventually, we should switch all controllers that
work with Pods to use virtual clocks to avoid clock-skew with older kubelets.

The PDB controller will need to be extended to recognize the Stable Condition in Pods.
It may also need to get into the business of estimating stability for ReplicaSets that
already use MinReadySeconds because some of their underlying Pods may run on old kubelets.
More discussion about the PDB controller can be found [here](https://github.com/kubernetes/kubernetes/issues/34776).

## Future work
Since we will move forward with Stable as the name of the Condition for a Pod that is Ready
for at least MinReadySeconds, we should also rename the status fields in ReplicaSet,
Deployment, and DaemonSet from AvailableReplicas to StableReplicas. Unfortunately this
cannot be done for v1 due to backwards-compatibility.
