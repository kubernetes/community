
## Background

The extensible admission control
[proposal](admission_control_extension.md)
proposed making admission control extensible. In the proposal, the `initializer
admission controller` and the `generic webhook admission controller` are the two
controllers that set default initializers and external admission hooks for
resources newly created. These two admission controllers are in the same binary
as the apiserver. This
[section](admission_control_extension.md#dynamic-configuration)
gave a preliminary design of the dynamic configuration of the list of the
default admission controls. This document hashes out the implementation details.

## Goals

* Admin is able to predict what initializers/webhooks will be applied to newly
  created objects.

* Admin needs to be able to ensure initializers/webhooks config will be applied within some bound

* As a fallback, admin can always restart an apiserver and guarantee it sees the latest config

* Do not block the entire cluster if the initializers/webhooks are not ready
  after registration.

## Specification

We assume initializers could be "fail open". We need to update the extensible
admission control
[proposal](admission_control_extension.md)
if this is accepted.

The schema is evolved from the prototype in
[#132](https://github.com/kubernetes/community/pull/132).

```golang
// InitializerConfiguration describes the configuration of initializers.
type InitializerConfiguration struct {
    metav1.TypeMeta

    v1.ObjectMeta

    // Initializers is a list of resources and their default initializers
    // Order-sensitive.
    // When merging multiple InitializerConfigurations, we sort the initializers
    // from different InitializerConfigurations by the name of the
    // InitializerConfigurations; the order of the initializers from the same
    // InitializerConfiguration is preserved.
    // +optional
    Initializers []Initializer `json:"initializers,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
}

// Initializer describes the name and the failure policy of an initializer, and
// what resources it applies to.
type Initializer struct {
    // Name is the identifier of the initializer. It will be added to the
    // object that needs to be initialized.
    // Name should be fully qualified, e.g., alwayspullimages.kubernetes.io, where
    // "alwayspullimages" is the name of the webhook, and kubernetes.io is the name
    // of the organization.
    // Required
    Name string `json:"name"`

    // Rules describes what resources/subresources the initializer cares about.
    // The initializer cares about an operation if it matches _any_ Rule.
    Rules []Rule `json:"rules,omitempty"`

    // FailurePolicy defines what happens if the responsible initializer controller
    // fails to takes action. Allowed values are Ignore, or Fail. If "Ignore" is
    // set, initializer is removed from the initializers list of an object if
    // the timeout is reached; If "Fail" is set, apiserver returns timeout error
    // if the timeout is reached. The default timeout for each initializer is
    // 5s.
    FailurePolicy *FailurePolicyType `json:"failurePolicy,omitempty"`
}

// Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended 
// to make sure that all the tuple expansions are valid.
type Rule struct {
    // APIGroups is the API groups the resources belong to. '*' is all groups.
    // If '*' is present, the length of the slice must be one.
    // Required.
    APIGroups []string `json:"apiGroups,omitempty"`

    // APIVersions is the API versions the resources belong to. '*' is all versions.
    // If '*' is present, the length of the slice must be one.
    // Required.
    APIVersions []string `json:"apiVersions,omitempty"`

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
    Resources []string `json:"resources,omitempty"`
}

type FailurePolicyType string

const (
    // Ignore means the initializer is removed from the initializers list of an
    // object if the initializer is timed out.
    Ignore FailurePolicyType = "Ignore"
    // For 1.7, only "Ignore" is allowed. "Fail" will be allowed when the
    // extensible admission feature is beta.
    Fail FailurePolicyType = "Fail"
)

// ExternalAdmissionHookConfiguration describes the configuration of initializers.
type ExternalAdmissionHookConfiguration struct {
    metav1.TypeMeta

    v1.ObjectMeta
    // ExternalAdmissionHooks is a list of external admission webhooks and the
    // affected resources and operations.
    // +optional
    ExternalAdmissionHooks []ExternalAdmissionHook `json:"externalAdmissionHooks,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
}

// ExternalAdmissionHook describes an external admission webhook and the
// resources and operations it applies to.
type ExternalAdmissionHook struct {
    // The name of the external admission webhook.
    // Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where
    // "imagepolicy" is the name of the webhook, and kubernetes.io is the name
    // of the organization.
    // Required.
    Name string `json:"name"`

    // ClientConfig defines how to communicate with the hook.
    // Required
    ClientConfig AdmissionHookClientConfig `json:"clientConfig"`

    // Rules describes what operations on what resources/subresources the webhook cares about.
    // The webhook cares about an operation if it matches _any_ Rule.
    Rules []RuleWithVerbs `json:"rules,omitempty"`

    // FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
    // allowed values are Ignore or Fail. Defaults to Ignore.
    // +optional
    FailurePolicy *FailurePolicyType
}

// RuleWithVerbs is a tuple of Verbs and Resources. It is recommended to make 
// sure that all the tuple expansions are valid.
type RuleWithVerbs struct {
    // Verbs is the verbs the admission hook cares about - CREATE, UPDATE, or *
    // for all verbs.
    // If '*' is present, the length of the slice must be one.
    // Required.
    Verbs []OperationType `json:"verbs,omitempty"`
    // Rule is embedded, it describes other criteria of the rule, like
    // APIGroups, APIVersions, Resources, etc. 
    Rule `json:",inline"`
}

type OperationType string

const (
    VerbAll OperationType = "*"
    Create  OperationType = "CREATE"
    Update  OperationType = "UPDATE"
    Delete  OperationType = "DELETE"
    Connect OperationType = "CONNECT"
)

// AdmissionHookClientConfig contains the information to make a TLS
// connection with the webhook
type AdmissionHookClientConfig struct {
    // Service is a reference to the service for this webhook. If there is only
    // one port open for the service, that port will be used. If there are multiple
    // ports open, port 443 will be used if it is open, otherwise it is an error.
    // Required
    Service ServiceReference `json:"service"`
    // CABundle is a PEM encoded CA bundle which will be used to validate webhook's server certificate.
    // Required
    CABundle []byte `json:"caBundle"`
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
    // Namespace is the namespace of the service
    // Required
    Namespace string `json:"namespace"`
    // Name is the name of the service
    // Required
    Name string `json:"name"`
}
```

Notes:
* There could be multiple InitializerConfiguration and
  ExternalAdmissionHookConfiguration. Every service provider can define their
  own.

* This schema asserts a global order of initializers, that is, initializers are
  applied to different resources in the *same* order, if they opt-in for the
  resources.

* The API will be placed at k8s.io/apiserver for 1.7.

* We will figure out a more flexible way to represent the order of initializers
  in the beta version.

* We excluded `Retry` as a FailurePolicy, because we want to expose the
  flakiness of an admission controller; and admission controllers like the quota
  controller are not idempotent.

* There are multiple ways to compose `Rules []Rule` to achieve the same effect.
  It is recommended to compact to as few Rules as possible, but make sure all
  expansions of the `<Verbs, APIGroups, APIVersions, Resource>` tuple in each
  Rule are valid. We need to document the best practice.

## Synchronization of admission control configurations

If the `initializer admission controller` and the `generic webhook admission
controller` watch the admission control configurations and act upon deltas, their
cached version of the configuration might be arbitrarily delayed. This makes it
impossible to predict what initializer/hooks will be applied to newly created
objects.

To make the behavior of `initializer admission controller` and the `generic
webhook admission controller` predictable, we let them do a consistent read (a
"LIST") of the InitializerConfiguration and ExternalAdmissionHookConfiguration
every 1s. If there isn't any successful read in the last 5s, the two admission
controllers block all incoming request. One consistent read per second isn't
going to cause performance issues.

In the HA setup, apiservers must be configured with --etcd-quorum-read=true.

See [Considered but REJECTED alternatives](#considered-but-rejected-alternatives) for considered alternatives.

## Handling initializers/webhooks that are not ready but registered

We only allow initializers/webhooks to be created as "fail open". This could be
enforced via validation. They can upgrade themselves to "fail closed" via the
normal Update operation. A human can also update them to "fail closed" later. 

See [Considered but REJECTED alternatives](#considered-but-rejected-alternatives) for considered alternatives.

## Handling fail-open initializers

The original [proposal](admission_control_extension.md) assumed initializers always failed closed. It is dangerous since crashed 
initializers can block the whole cluster. We propose to allow initializers to 
fail open, and in 1.7, let all initializers fail open.

#### Implementation of fail open initializers.

In the initializer prototype
[PR](https://github.com/kubernetes/kubernetes/pull/36721), the apiserver that
handles the CREATE request
[watches](https://github.com/kubernetes/kubernetes/pull/36721/files#diff-2c081fad5c858e67c96f75adac185093R349)
the uninitialized object. We can add a timer there and let the apiserver remove
the timed out initializer.

If the apiserver crashes, then we fall back to a `read repair` mechanism. When
handling a GET request, the apiserver checks the objectMeta.CreationTimestamp of
the object, if a global initializer timeout (e.g., 10 mins) has reached, the
apiserver removes the first initializer in the object.

In the HA setup, apiserver needs to take the clock drift into account as well.

Note that the fallback is only invoked when the initializer and the apiserver
crashes, so it is rare.

See [Considered but REJECTED alternatives](#considered-but-rejected-alternatives) for considered alternatives.

## Future work

1. Figuring out a better schema to represent the order among
   initializers/webhooks, e.g., adding fields like lists of initializers that
   must execute before/after the current one.

2. #1 will allow parallel initializers as well.

3. implement the fail closed initializers according to
   [proposal](admission_control_extension.md#initializers).

4. more efficient check of AdmissionControlConfiguration changes. Currently we
   do periodic consistent read every second.

5. block incoming requests if the `initializer admission controller` and the
   `generic webhook admission controller` haven't acknowledged a recent change
   to AdmissionControlConfiguration. Currently we only guarantee a change
   becomes effective in 1s.

## Considered but REJECTED alternatives:

### synchronization mechanism

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

### Handling not ready initializers/webhook

#### Rejected 1. 

add readiness check to initializer and webhooks, `initializer admission
controller` and `generic webhook admission controller` only apply those have
passed readiness check. Specifically, we add `readiness` fields to
`AdmissionControllerConfiguration`; then we either create yet another controller
to probe for the readiness and update the `AdmissionControllerConfiguration`, or
ask each initializer/webhook to update their readiness in the
`AdmissionControllerConfigure`. The former is complex.  The latter is
essentially the same as the first approach, except that we need to introduce the
additional concept of "readiness".

### Handling fail-open initializers

#### Rejected 1. use a controller

A `fail-open initializers controller` will remove the timed out fail-open
initializers from objects' initializers list. The controller uses shared
informers to track uninitialized objects. Every 30s, the controller 

* makes a snapshot of the uninitialized objects in the informers.
* indexes the objects by the name of the first initializer in the objectMeta.Initializers
* compares with the snapshot 30s ago, finds objects whose first initializers haven't changed
* does a consistent read of AdmissionControllerConfiguration, finds which initializers are fail-open
* spawns goroutines to send patches to remove fail-open initializers
