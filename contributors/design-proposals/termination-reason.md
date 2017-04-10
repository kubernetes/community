# Pod Termination Reason

**Author**: kow3ns@ yguo0905@

**Status**: Proposal

## Abstract

Currently, as of Kubernetes 1.6, when a Pod is
 [terminated](https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods), 
 the Pod's containers will have their 
 [pre-stop handlers](https://kubernetes.io/docs/concepts/containerscontainer-lifecycle-hooks/#hook-detail) 
invoked prior to receiving a `TERM` signal. Kubelet will wait until the 
specified termination grace period has expired, and then send a `KILL` signal. 
While this grace period allows a container's entry point to perform any 
necessary cleanup, the cleanup tasks performed are often dependent on the reason 
for termination. 
In this proposal, we design a mechanism to deliver a signal to a Pod's 
containers indicating why the Pod has been terminated. Applications can use 
this additional contextual information to determine what actions they should 
take during their termination grace period.

## Motivation
Consider a stateful, replicated storage application where the application's data 
are non-uniformly replicated across servers. Such replication schemes are common 
for BASE storage applications (e.g. Cassandra, Riak) and for sharded, 
replicated, strongly consistent storage applications (e.g Couchbase, MongoDB, 
HBase). For many such applications, it is not atypical for a single server to be 
responsible for serving TiBs of data.

When a Pod containing such an application is terminated because it will be 
decommissioned, terminated without the intent of resurrection (e.g. when the 
user horizontally scales the application down), during its termination grace 
period, we expect the application, or a process managing the application, 
to do the following.

1. Drain all open client connections.
1. Flush any in memory data structures to the file system.
1. Synchronize the file system with storage media.
1. Trigger a global redistribution of data so that the application's data are 
not left under replicated.
1. Modify the application membership to remove the decommissioned instance.

However, in contrast to the above, when a Pod is terminated with the intention
that it should be rescheduled (e.g. The Pod's controller is performing a 
configuration, image, or resource update, the Pod is being moved due to a 
taint, or the Pod is being evicted due to a managed node upgrade),
we expect the application, or its managing process, to do the following.

1. Drain all open client connections.
1. Flush any in memory data structures to the file system.
1. Synchronize the file system with storage media.
1. Gracefully terminated the application without triggering a data 
 redistribution.

This will leave the application temporarily under replicated, and, we claim
that, for a distributed storage application this temporary under replication is 
desirable. Consider that, such applications are designed to remain available 
during some number of temporary process failures, and PodDisruptionBudgets 
provide a way for application administrators to communicate the number of 
planned temporary process failures (disruptions) that an application 
can tolerate. 

When we expect a terminated Pod to be rescheduled immediately, or nearly so, 
redistributing its data will only increase the expected time to recovery. This 
is only exacerbated by the fact that when the terminated Pod is resurrected, 
the application will have to once again redistribute its data to 
make use of the restored capacity. For applications that may require TiBs of 
data to be redistributed for a single Pod update, the network, storage, and 
time costs associated with redistribution are prohibitive when there is a 
very high probability that we can reschedule the Pod immediately. 

If the application's Persistent Volumes are backed by local storage media (note 
that this is not currently possible but we expect this feature to land in the 
near future), we expect a decommissioned instance to behave identically to an 
instance whose PVs are backed by remote storage media. However, as a local PV is 
intrinsically tied to the lifetime of the Node, leaving the application 
under replicated during destructive updates to the Pod is inherently more 
risky. We still believe that, under these circumstances, temporary 
under replication is a viable strategy, and users can mitigate their risk by 
increasing the number of replicas of the application's data. 

## Affected Components
1. API Server
1. API Machinery
1. Controllers (that wish to utilize this feature)
1. Kubelet
1. Kubectl

## Use Cases

1. As the implementor of a controller or operator that manages a stateful 
workload, I want to send a signal to the Pods created by my controller or 
operator to control the application's replication behavior during Pod 
termination.
1. As the implementor of tenant applications on a Kubernetes cluster, I want 
my applications to receive a signal when they are evicted from a node due to 
intolerable taints.
1. As the administrator of a Kubernetes cluster, I want to provide a signal to 
tenant applications that they are being terminated in violation of their 
declared disruption budgets or termination grace periods due to a superseding, 
global, organization policy (e.g. My organization requires that an application's 
termination grace period and disruption budget allow managed upgrades to make 
progress.)
1. As a user I want to be able to send a termination reason to my Pods when I 
use the Kubernetes API or kubectl to gracefully delete a Pod.
1. As a Reliability Engineer for an application running on a Kubernetes cluster,
I want to configure Pods to communicate the reason that they were terminated via 
a HTTP GET request so that I can configure monitoring and alerting based on 
historic trends.

## Assumptions

1. Kubelet will always invoke the pre-stop handler prior to sending a `TERM` 
signal to the entry point of a container.
1. pre-stop handlers will not contain complicated or long running business logic. 
The logic associated with container cleanup will be executed by the container's
entry point during the container's configured termination grace period.

## Requirements

1. Termination reason must be available to a container at the time its pre-stop 
handler is invoked. In this way, applications can perform any actions
necessary to configure the container's entry point prior to the entry point 
receiving the `TERM` signal.
1. The delivery mechanism for termination reason must be consumable by, and 
 compatible with, existing HTTP and command based pre-stop handlers.
1. Any headers, variables, or parameters used to communicate a termination 
reason must not be hard coded, and the receiving container must be able to 
override the defaults.

## API Objects

Termination reason is implemented by a string type.

```golang
// TerminationReason is a string type used to indicate the reason for 
// termination to the preStop lifecycle hook of a Pod's containers.
type TerminationReason string

const (
	// ReasonEviction is the default reason used to communicate that a Pod has 
 	// been terminated due to an eviction.
	ReasonEviction TerminationReason = "Eviction"
	// ReasonIntolerableTaint is the default reason used to communicate that a Pod 
	// has been terminated due to a taint for which it has no declared toleration.
	ReasonIntolerableTaint = "IntolerableTaint"
	// ReasonDecommissioned is the default reason sent by controllers to 
	// indicate that a Pod has been terminated due to a horizontal scaling 
	// event. Pods receiving this reason should not expect to be rescheduled.
	ReasonDecommissioned = "Decommissioned"
	// ReasonUpdate is the reason sent by a controller or User to indicate that 
	// a Pod has been terminated for the purposes of a destructive update. Pods 
	// receiving this signal should expect to be rescheduled immediately. Note
	// that this DOES NOT imply that scheduling will succeed.
	ReasonUpdate = "Update"
)

```

DeleteOptions is modified to carry the termination reason.

```golang
type DeleteOptions struct {
	// Other fields omitted for brevity.
	
	// Reason indicates the reason for the Pod's termination. This field may 
	// be supplied by a user or a controller.
	Reason *TerminationReason `json:"reason,omitempty"`
}
```

The ObjectMeta struct is modified to carry the termination reason via 
a Pod's Metadata.

```golang
type ObjectMeta {
	// Other fields omitted for brevity.
	
        // TerminationReason indicates the reason for the Pod's termination.
        // This field may be supplied by a user or a controller.
	TerminationReason *TerminationReason `json:"reason,omitempty"`
}
```

TerminationReasonDelivery provides configuration delivery method of Lifecycle's 
`PreStop` Handler.

```golang
// TerminationReasonDelivery is used to configure the delivery method for
// termination reasons. It is a union type, and exactly one of the fields may be
// non-nil. The Env field is compatible with command preStop lifecycle hooks,
// and the Header field is compatible with HTTP GET hooks.
type TerminationReasonDelivery struct {
    // Env is the name of the environment variable that will be set with the 
    // termination reason. Env must only be set when used with a command 
    // Action, and it must be a valid environment variable name.
    Env *string `json:"env,ommitempty"`
    // Header is the name of the header that will be set to the termination 
    // reason. It must only be set when used with a HTTP GET Action.
    Header *string `json:"header,ommitempty"`
}

const ( 
    // DefaultTerminationReasonEnv is the default environment variable that 
    // is used to communicate a termination reason to a command Action.
    DefaultTerminationReasonEnv string = "KUBE_POD_TERM_REASON"
    // DefaultTerminationReasonHeader is the default header used to communicate
    // a termination reason to a HTTP GET Action.
    DefaultTerminationReasonHeader string = "KUBE-POD-TERM-REASON"
)
```

PreStopHandler is introduced to aggregate Handler and TerminationReasonDelivery.

```golang
// PreStopHandler aggregates Handler and a TerminationReasonDelivery to allow 
// for configuration of the delivery method for a termination reason consumed 
// by the hook.
type PreStopHandler struct {
    Handler
    // ReasonDelivery provides configuration for the delivery of a termination 
    // reason to the PreStopHandler. If nil, the termination reason will be 
    // delivered to the preStop lifecycle hook by setting the
    // KUBE_POD_TERM_REASON environment variable to the value of the termination
    // reason.
    ReasonDelivery *TerminationReasonDelivery `json:"reasonDelivery,ommitempty"`
}
```

The Lifecycle struct is modified such that `PreStop` is a PreStopHandler.

```golang
type Lifecycle struct {
    // Other fields and existing comments omitted for brevity.
 
    PreStop *PreStopHandler
}
```

## API Server 
During its graceful delete processing, When a termination reason is specified 
via DeleteOptions, the API server will include the reason in the 
Pod's metadata when setting the Pod's `DeletionTimestamp`.

### Validation
In addition to the existing validation performed for the Lifecycle struct, the 
API Server should fail validation if `PreStop` is not nil, and if one of the 
following are true.

1. `PreStop` indicates a command action and `.ReasonDelivery.Header` is not nil.
1. `PreStop` indicates a HTTP GET action and `.ReasonDelivery.Env` is not nil.
1. Both `.ReasonDelivery.Header` and `.ReasonDelivery.Env` are not nil.
1. `.ReasonDelivery.Env` is not nil and it points to string that is not a 
valid environment variable.

### Pod Deletion
When the API Server performs its graceful delete processing, in addition to 
setting the `DeletionTimestamp` of the subject ObjectMeta, if the DeleteOptions 
contains a `Reason`, the `Reason` should be copied to the `TerminationReason` 
field of the ObjectMeta.

## Kubelet
Kubelet will supply a termination reason via a configured pre-stop lifecycle 
hook under the following conditions.

1. The Pod has been explicitly deleted with a supplied `TerminationReason`.
1. The Pod is the target of an eviction.
1. The Pod is being terminated due to an intolerable taint.

In all cases Kubelet will deliver the termination reason to a container 
with a declared `PreStop` handler in accordance with the configuration of its 
`ReasonDelivery`.

### Termination Reason Delivery

When Kublet processes the `PreStop` Handlers of a Pod's containers, prior to 
sending a `TERM` signal and starting the termination grace timer, if a 
termination reason is to be delivered, Kubelet will do the following.

1. If the `PreStop` handler indicates a command action, Kubelet will supply the 
termination reason to the container based on the following criteria.
   1. If the `ReasonDelivery` is nil Kubelet will set the 
   `DefaultTerminationReasonEnv` to the value of the termination reason. This 
   is the default method of delivery for a command action pre-stop handler.
   1. If the `Env` field of the `TerminationReasonDelivery` is not nil, Kubelet 
   will set the environment variable indicated by this field to the value of 
   the termination reason.
   1. If none of the above criteria are met, then the pre-stop handler is 
   malformed, and API Server [validation](#validation) has failed to reject 
   its creation or update. In this case, Kubelet will log an error and deliver 
   the termination reason via the default method specified above.
1. If the `PreStop` handler indicates a HTTP GET action, Kublet will supply the 
termination reason to the specified endpoint based on the following criteria.
   1. If the `ReasonDelivery` is nil Kubelet will set the header indicated by 
   `DefaultTerminationReasonHeader` to the value of the termination reason. This 
   is the default method of delivery for HTTP GET pre-stop handlers.
   1. If `ReasonDelivery` is not nil, and if its `Header` field 
   is not nil, Kubelet will add a header, whose name is indicated by the value 
   of this field and whose value is the termination reason, to the HTTP GET 
   request, prior to sending the request.
   1. If none of the above criteria are met, then the pre-stop handler is 
   malformed, and API Server [validation](#validation) has failed. In this case,
   Kubelet will log an error and deliver the termination reason via the default
   header specified above.

### Pod Deletion 

When Kubelet finds that a Pod's `DeletionTimestamp` is set, during its 
termination processing, if a `TerminationReason` has been set, it will 
[deliver the termination reason](#termination-reason-delivery).

### Pod Eviction

When Kubelet targets a Pod for eviction, it will 
[deliver a termination reason](#termination-reason-delivery) of 
`ReasonEviction`.

### Intolerable Taints

When Kubelet evicts a Pod due to an intolerable taint, it will 
[delivery a termination reason](#termination-reason-delivery) of 
`ReasonIntolerableTaint`.

## Kubectl
Kubectl will use the `--reason` parameter to allow users to pass a arbitrary 
string as the termination reason as shown below.

```shell
 > kubectl delete po my-pod --reason="resolve issue 354961"
```

Kubectl will simply populate the `Reason` field of the DeleteOptions for the 
DELETE request with the supplied reason.

## Tests
- A termination reason can be delivered to a command action via an environment 
variable
- A termination reason can be delivered to a command action via a flag.
- A termination reason can be delivered to a HTTP GET action via a header.
- A termination reason is delivered during Pod eviction.
- A termination reason is delivered when a Pod is evicted due to an intolerable 
taint.
- A termination reason is delivered when provided by kubectl.

