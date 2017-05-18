# Dynamic admission control configuration

## Background

The extensible admission control
[proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/admission_control_extension.md)
proposed making admission control extensible. In the proposal, the `initializer
admission controller` and the `generic webhook admission controller` are the two
controllers that set default initializers and external admission hooks for
resources newly created. These two admission controllers are in the same binary
as the apiserver. This
[section](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/admission_control_extension.md#dynamic-configuration)
gave a preliminary design of the dynamic configuration of the list of the
default admission controls. This document hashes out the implementation details.

## Goals

* Admin is able to predict what initializers/webhooks will be applied to newly
  created objects.

* Admin needs to be able to ensure initializers/webhooks config will be applied within some bound

* As a fallback, admin can always restart an apiserver and guarantee it sees the latest config

* Do not block the entire cluster if the intializers/webhooks are not ready
  after registration.

## Specification

We assume initializers could be "fail open". We need to update the extensible
admission control
[proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/admission_control_extension.md)
if this is accepted.

The schema is copied from 
[#132](https://github.com/kubernetes/community/pull/132) with a few
modifications.

```golang
type AdmissionControlConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ResourceInitializers is a list of resources and their default initializers
	// +optional
	ResourceInitializers []ResouceInitializer `json:"resourceInitializers,omitempty" protobuf:"bytes,1,rep,name=resourceInitializers"`

	// ExternalAdmissionHooks is a list of external admission webhooks and the
	// affected resources and operations.
	// +optional
	ExternalAdmissionHooks []ExternalAdmissionHook `json:"externalAdmissionHooks,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=externalAdmissionHooks"`
}

// ResouceInitializer describes the default initializers that will be
// applied to a resource. The order of initializers is sensitive.
type ResouceInitializer struct {
	// APIGroup is the API group of the resource
	// Required.
	APIGroup string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup"`

	// APIVersions is the API Versions of the resource
	// '*' means all API Versions.
	// If '*' is present, the length of the slice must be one.
	// Required.
	APIVersions []string `json:"apiVersions,omitempty" protobuf:"bytes,2,rep,name=apiVersions"`

	// Resource is resource to be initialized
	// Required.
	Resource string `json:"resource" protobuf:"bytes,3,opt,name=resource"`

	// Initializers is a list of initializers that will be applied to the
	// resource by default. It is order-sensitive.
	Initializers []Initializer `json:"initializers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,3,rep,name=initializers"`
}

// Initializer describes the name and the failure policy of an initializer.
type Initializer struct {
	// Name is the identifier of the initializer. It will be added to the
	// object that needs to be initialized.
	// Required
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// FailurePolicy defines what happens if the responsible initializer controller
	// fails to takes action. Allowed values are Ignore, or Fail. If "Ignore" is
	// set, initializer is removed from the initializers list of an object if
	// the timeout is reached; If "Fail" is set, apiserver returns timeout error
	// if the timeout is reached.
	FailurePolicy *FailurePolicyType `json:"failurePolicy,omitempty" protobuf:"bytes,2,opt,name=failurePolicy"`
}

type FailurePolicyType string

const (
	// Ignore means the initilizer is removed from the initializers list of an
	// object if the initializer is timed out.
	Ignore FailurePolicyType = "Ignore"
	// For 1.7, only "Ignore" is allowed. "Fail" will be allowed when the
	// extensible admission feature is beta.
	Fail FailurePolicyType = "Fail"
)

// ExternalAdmissionHook describes an external admission webhook and the
// resources and operations it applies to.
type ExternalAdmissionHook struct {
	// The name of the external admission webhook.
	// Required.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// ClientConfig defines how to communicate with the hook.
	// Required
	ClientConfig AdmissionHookClientConfig `json:"clientConfig" protobuf:"bytes,2,opt,name=clientConfig"`

	// Rules describes what operations on what resources/subresources the webhook cares about.
	// The webhook cares about an operation if it matches _any_ Rule.
	Rules []Rule `json:"rules,omitempty" protobuf:"bytes,3,rep,name=rules"`

	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
	// allowed values are Ignore or Fail. Defaults to Ignore.
	// +optional
	FailurePolicy *FailurePolicyType
}

// Rule describes the Verbs and Resources an admission hook cares about. Each
// Rule is a tuple of Verbs and Resources.It is recommended to make sure all
// the tuple expansions are valid.
type Rule struct {
	// Verbs is the verbs the admission hook cares about - CREATE, UPDATE, or *
	// for all verbs.
	// If '*' is present, the length of the slice must be one.
	// Required.
	Verbs []OperationType `json:"verbs,omitempty" protobuf:"bytes,1,rep,name=verbs"`

	// APIGroups is the API groups the resources belong to. '*' is all groups.
	// If '*' is present, the length of the slice must be one.
	// Required.
	APIGroups []string `json:"apiGroups,omitempty" protobuf:"bytes,2,rep,name=apiGroups"`

	// APIVersions is the API versions the resources belong to. '*' is all versions.
	// If '*' is present, the length of the slice must be one.
	// Required.
	APIVersions []string `json:"apiVersions,omitempty" protobuf:"bytes,3,rep,name=apiVersions"`

	// Resources is a list of resources this rule applies to.
	//
	// For example:
	// 'pods' means pods.
	// 'pods/log' means the log subresource of pods.
	// '*' means all resources, but not subresources.
	// 'pods/*' means all subresources of pods.
	// '*/scale' means all scale subresources.
	// '*/*' means all resources and their subresources.
	//
	// If '*' or '*/*' is present, the length of the slice must be one.
	// Required.
	Resources []string `json:"resources,omitempty" protobuf:"bytes,4,rep,name=resources"`
}

type OperationType string

const (
	VerbAll OperationType = "*"
	Create  OperationType = "CREATE"
	Update  OperationType = "UPDATE"
)

// AdmissionHookClientConfig contains the information to make a TLS
// connection with the webhook
type AdmissionHookClientConfig struct {
	// Service is a reference to the service for this webhook. If there is only
	// one port open for the service, that port will be used. If there are multiple
	// ports open, port 443 will be used if it is open, otherwise it is an error.
	// Required
	Service ServiceReference `json:"service" protobuf:"bytes,1,opt,name=service"`
	// CABundle is a PEM encoded CA bundle which will be used to validate webhook's server certificate.
	// Required
	CABundle []byte `json:"caBundle" protobuf:"bytes,2,rep,name=caBundle"`
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// Namespace is the namespace of the service
	// Required
	Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
	// Name is the name of the service
	// Required
	Name string `json:"name" protobuf:"bytes,2,opt,name=name"`
}
```

## Synchronization of AdmissionControlConfiguration 

If the `initializer admission controller` and the `generic webhook admission
controller` watch the `AdmissionControlConfiguration` and act upon deltas, their
cached version of the configuration might be arbitrarily delayed. This makes it
impossible to predict what initializer/hooks will be applied to newly created
objects.

We propose the following way to make the behavior of the `initializer admission
controller` and the `generic webhook admission controller` predictable.

#### 1. Do consistent read of AdmissionControlConfiguration periodically

The `initializer admission controller` and the `generic webhook admission
controller` do a consistent read of the AdmissionControlConfiguration every 1s.
If there isn't any successful read in the last 5s, the two admission controllers 
block all incoming request.
One consistent read per second isn't going to cause performance issues.

## What if an initializer controller/webhook is not ready after registered?

This will block the entire cluster. We have a few options:

1. only allow initializers/webhooks to be created as "fail open". This could be
   enforced via validation. They can upgrade themselves to "fail closed" via the
   normal Update operation. A human can also update them to "fail closed" later. 

2. less preferred: add readiness check to initializer and webhooks, `initializer
   admission controller` and `generic webhook admission controller` only apply
   those have passed readiness check. Specifically, we add `readiness` fields to
   `AdmissionControllerConfiguration`; then we either create yet another
   controller to probe for the readiness and update the
   `AdmissionControllerConfiguration`, or ask each initializer/webhook to update
   their readiness in the `AdmissionControllerConfigure`. The former is complex.
   The latter is essentially the same as the first approach, except that we need
   to introduce the additional concept of "readiness".

## Handling fail-open initializers

The original [proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/admission_control_extension.md) assumed initializers always failed closed. It is dangerous since crashed 
initializers can block the whole cluster. We propose to allow initializers to 
fail open, and in 1.7, let all initializers fail open. We considered the two
approaches to implement the fail open initializers.

#### 1. apiserver + read repair

In the initializer prototype
[PR](https://github.com/kubernetes/kubernetes/pull/36721), the apiserver that
handles the CREATE request
[watches](https://github.com/kubernetes/kubernetes/pull/36721/files#diff-2c081fad5c858e67c96f75adac185093R349)
the uninitialized object. We can add a timer there and let the apiserver remove
the timed out initializer.

If the apiserver crashes, then we fall back to a `read repair` mechanism. When
handling a GET request, the apiserver checks the objectMeta.CreationTimestamp of
the object, if a global intializer timeout (e.g., 10 mins) has reached, the
apiserver removes the first initializer in the object.

In HA setup, apiserver needs to take the clock drift into account as well.

Note that the fallback is only invoked when the initializer and the apiserver
crashes, so it is rare.

#### 2. use a controller

A `fail-open initializers controller` will remove the timed out fail-open
initializers from objects' initializers list. The controller uses shared
informers to track uninitialized objects. Every 30s, the controller 

* makes a snapshot of the uninitialized objects in the informers.
* indexes the objects by the name of the first initialilzer in the objectMeta.Initializers
* compares with the snapshot 30s ago, finds objects whose first initializers haven't changed
* does a consistent read of AdmissionControllerConfiguration, finds which initializers are fail-open
* spawns goroutines to send patches to remove fail-open initializers

## Future work

1. allow the user to POST to individual initializer/webhook, expressing partial
   order among initializers/webhooks, and let a controller assembles the
   ordered list of initializers/webhooks.

2. #1 will allow parallel initializers as well.

3. make the AdmissionControllerConfiguration more flexible in expressing the
   combination of verbs and resources, if needed.

4. implement the fail closed initializers according to
   [proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/admission_control_extension.md#initializers).

5. more efficient check of AdmissionControlConfiguration changes. Currently we
   do periodic consistent read every second.

6. block incoming requests if the `initializer admission controller` and the
   `generic webhook admission controller` haven't acknowledged a recent change
   to AdmissionControlConfiguration. Currently we only guarantee a change
   becomes effective in 1s.

## Considered but REJECTED synchronization mechinism:

#### Rejected 1. Always do consistent read

Rejected because of inefficiency.

The `initializer admission controller` and the `generic webhook admission
controller` always do consistent read of the `AdmissionControlConfiguration`
before applying the configuration to the incoming objects. This adds latency to
every CREATE request. Because the two admission controllers are in the same
process as the apiserver, the latency mainly consists of the consistent read
latency of the backend storage (etcd), and the proto unmarshalling.


#### Rejected 2. Don't synchronize, but report what is the cached version

Rejected because it violates Goal 2 on the time bound.

The main goal is *NOT* to always apply the latest
`AdmissionControlConfiguration`, but to make it predictable what
initializers/hooks will be applied. If we introduce the
`generation/observedGeneration` concept to the `AdmissionControlConfiguration`,
then a human (e.g., a cluster admin) can compare the generation with the
observedGeneration and predict if all the initializer/hooks listed in the
`AdmissionControlConfiguration` will be applied. 

In the HA setup, the `observedGeneration` reported by of every apiserver's
`initializer admission controller` and `generic webhook admission controller`
are different, so the API needs to record multiple `observedGeneration`.

#### Rejected 3. Always do a consistent read of a smaller object

Rejected because of the complexity.

A consistent read of the AdmissionControlConfiguration object is expensive, we
cannot do it for every incoming request.

Alternatively, we record the resource version of the AdmissionControlConfiguration
in a configmap. The apiserver that handles an update of the AdmissionControlConfiguration
updates the configmap with the updated resource version. In the HA setup, there
are multiple apiservers that update this configmap, they should only
update if the recorded resource version is lower than the local one.

The `initializer admission controller` and the `generic webhook admission
controller` do a consistent read of the configmap *everytime* before applying
the configuration to an incoming request. If the configmap has changed, then
they do a consistent read of the `AdmissionControlConfiguration`.
