# Dynamic admission control configuration

## Background

[#132](https://github.com/kubernetes/community/pull/132) proposed making
admission control extensible. In the proposal, the `initializer admission
controller` and the `generic webhook admission controller` are the two
controllers that set default initializers and external admission hooks for
resources newly created. These two admission controllers are in the same binary
as the apiserver. This [section](https://github.com/smarterclayton/community/blob/be132e88f7597ab3927b788a3de6d5ab6de673d2/contributors/design-proposals/admission_control_extension.md#dynamic-configuration)
of #132 gave a preliminary design of the dynamic configuration of the list of
the default admission controls. This document hashes out the implementation
details.

## Goals

* Admin is able to predict what initializers/webhooks will be applied to newly
  created objects.

* Admin needs to be able to ensure initializers/webhooks config will be applied within some bound

* As a fallback, admin can always restart an apiserver and guarantee it sees the latest config

* Do not block the entire cluster if the intializers/webhooks are not ready
  after registration.

## Specification

We assume initializers could be "fail open". We need to update #132 if this is
accepted.

The schema is copied from
[#132](https://github.com/kubernetes/community/pull/132) with a few
modifications.

```golang
type AdmissionControlConfiguration struct {
    TypeMeta // although this object could simply be serialized like ComponentConfig

    // ResourceInitializers is a list of resources and their default initializers
    ResourceInitializers []ResourceDefaultInitializer

    ExternalAdmissionHooks []ExternalAdmissionHook
}

// Because the order of initializers matters, and each resource might need
// differnt order, the ResourceDefaultInitializers are indexed by Resource.
type ResourceDefaultInitializer struct {
    // Resource identifies the type of resource to be initialized that should be
    // initialized
    Resource GroupResource
    // Initializers are the default names that will be registered to this resource
    Initializers []Initializer
}

type Initializer struct {
    // Name is the string that will be registered to the resource that needs
    // initialization.
    Name string

    // **Optional for alpha implement**
    // FailurePolicy defines what happens if there is no initializer controller
    // takes action. Allowed values are Ignore, or Fail. If "Ignore" is set, 
    // apiserver removes initilizer from the initializers list of the resource
    // if the timeout is reached; If "Fail" is set, apiserver returns timeout
    // error if the timeout is reached.
    FailurePolicy FailurePolicyType

    // **Optional for alpha implement**
    // If timeout is reached, the intializer is removed from the resource's
    // initializer list by the apiserver.
    // Default to XXX seconds.
    Timeout *int64
}

type FailurePolicyType string

const (
    Ignore FailurePolicyType = "Ignore"
    Fail FailurePolicyType = "Fail"
)

type ExternalAdmissionHook struct {
    // Operations is the list of operations this hook will be invoked on - Create, Update, or *
    // for all operations. Defaults to '*'.
    Operations []OperationType
    // Resources are the resources this hook should be invoked on. '*' is all resources.
    Resources []Resource
    // Subresources is list of subresources. If non-empty, this hook should be invoked on 
    // all combinations of Resources and Subresources. '*' is all subresources.
    Subresources []string

    // ClientConfig defines how to talk to the hook.
    ClientConfig AdmissionHookClientConfig

    // FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
    // allowed values are Ignore, Fail. Default value is Fail
    FailurePolicy FailurePolicyType
}

type Resource struct {
    // Group is the API group the resource belongs to.
    Group string
    // Resource is the name of the resource.
    Resource string
}

type OperationType string

const (
    All OperationType = "*"
    Create OperationType= "Create"
    Update OperationType= "Update"
)

// AdmissionHookClientConfig contains the information to make a TLS
// connection with the webhook
type AdmissionHookClientConfig struct {
    // Service is a reference to the service for this webhook. It must communicate
	// on port 443
	Service ServiceReference
	// CABundle is a PEM encoded CA bundle which will be used to validate webhook's server certificate.
	CABundle []byte
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// Namespace is the namespace of the service
	Namespace string
	// Name is the name of the service
	Name string
}
```

## Synchronization of AdmissionControlConfiguration 

If the `initializer admission controller` and the `generic webhook admission
controller` watch the `AdmissionControlConfiguration` and act upon deltas, their
cached version of the configuration might be arbitrarily delayed. This makes it
impossible to predict what initializer/hooks will be applied to newly created
objects.

We propose two ways to make the behavior of the `initializer admission
controller` and the `generic webhook admission controller` predictable.

#### 1. Do consistent read of AdmissionControlConfiguration periodically

The `initializer admission controller` and the `generic webhook admission
controller` do a consistent read of the AdmissionControlConfiguration either 30s
after the last read, or when there is a request that needs the two controllers to
apply the configuration, whichever comes later.

If the read fails, the two admission controllers block all incoming request.

#### 2. Always do a consistent read of a smaller object

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

## What if an initializer controller/webhook is not ready after registered? (**optional for alpha implement**)

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


## Considered bug REJECTED synchronization mechinism:

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
