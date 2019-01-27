# Graduate Admission Webhooks to GA

Authors: @mbohlool

Reviewers: @liggit @deads2k @sttts @caesarxuchao

Sig: api-machinery

## Summary

Admission webhooks are a way to extend kubernetes by putting hooks over object creation/modification. Admission webhooks can mutate or validate the object. This feature has been Beta since Kubernetes 1.9. This document outline required steps to graduate it to GA.

## New Features

Based on users feedback, these are the changes that are proposed to graduate Admission Webhooks to GA.

### Object selector

Currently Admission Webhook supports namespace selectors, but that may not be enough for some cases that admission webhook need to be skipped on some objects. For example if the Admission Webhook is running inside cluster and its rules includes wildcards which match required API objects for its own execution. An object selector would be useful exclude those objects. Also in case of an optional webhooks, an object selector gives the end user to include or exclude an object without having access to admission webhook configuration which is normally restricted to cluster admins.
Note that namespace objects must match both object selector (if specified) and namespace selector to be sent to the Webhook except for cluster-scoped resources which namespace selector will not applied to. The propsoe change is to add an ObjectSelector to the webhook API both in v1 and v1beta1.

```golang
type Webhook struct {
    ...
     // ObjectSelector decides whether to run the webhook on an object based
     // on whether the object.metadata.labels matches the selector. An object is
     // matches both NamespaceSelector and ObjectSelector.
     //
     // Default to the empty LabelSelector, which matches everything.
     // +optional
     ObjectSelector *metav1.LabelSelector `json:"objectSelector,omitempty" protobuf:"bytes,7,opt,name=objectSelector"`
}
```

### Scope

Current webhook Rules applies to objects of all scopes. That means a Rule that uses wildcards to target an object type will be applied to both cluster scoped and namespaced objects. The proposal is to add a scope field to Admission Webhook configuration to limit webhook target on namespaced object or cluster scoped objects. The field will be added to both v1 and v1beta1. The field is optional and empty value means no scope restriction.

```golang
type ScopeType string

const (
     // ClusterScope means that scope is limited to cluster objects.
     ClusterScope ScopeType = "Cluster"
     // NamespacedScope means that scope is limited to namespaced objects.
     NamespacedScope ScopeType = "Namespaced"
)

type Rule struct {
    ...

     // Scope specifies the scope of this rule. If unespecified, the scope is
     // not limited.
     //
     // +optional
     Scope ScopeType `json:"scope,omitempty" protobuf:"bytes,3,opt,name=scope"`
}
```

### http for localhost

Admission Webhook requires https for all connections. This is not necessary when the webhook in running on the same machine as apiserver as a side-car or static pod. By relaxing this restriction, it would be possible to run side-cards or static pods on the same machine without cert management. The API change for this feature is only documentation.

### timeout configuration

Admission Webhook has a default timeout of 30 seconds today but long running webhooks would result in a poor performance. By adding a timeout field to the configuration, the webhook author can limit the running time of the webhook that either result in failing the API call earlier or ignore the webhook call based on the failure policy. This feature, however, would not let the timeout to be extended more than todays 30 seconds default and the timeout would be defaulted to 10 seconds for v1 API while stays 30 second for v1beta API to keep backward compatibility.

```golang
type Webhook struct {
    ...
     // TimeoutSeconds specifies the timeout for this webhook. After the timeout passes,
     // the webhook call will be ignored or the API call will fail based on the
     // failure policy.
     // The timeout value should be between 1 and 30 seconds.
     // Default to 10 seconds for v1 API and 30 seconds for v1beta1 API.
     // +optional
     TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,8,opt,name=timeoutSeconds"`
}
```

### Port configuration

Today Admission Webhook port is always expected to be 443 on service reference. But this limitation was arbitrary and there are cases that Admission Webhook cannot use this port. This feature will add a port field to service based webhooks and allows specifying a port other than 443 for service based webhooks. Specifying port should already be available for URL based webhooks.

```golang
type ServiceReference struct {
    ...

     // If specified, the port on the service that hosting webhook.
     // Default to 443 for backward compatibility.
     // +optional
     Port *int32 `json:"port,omitempty" protobuf:"varint,4,opt,name=port"`
}
```

### AdmissionReview v1

The payload API server sends to Admission webhooks is called AdmissionReview which is `v1beta1` today. The proposal is to promote the API to v1 with no change to the API object definition. Because of different versions of Admission Webhooks, there should be a way for the webhook developer to specify which version of AdmissionReview they support. The version should be an ordered list which reflects the webhooks preference of the versions.
A version list will be added to webhook configuration that would be defaulted to `['v1beta1']` in v1beta1 API and will be a required field in v1.

V1 API looks like this:

```golang
type Webhook struct {
    ...

     // AdmissionReviewVersions is an ordered list of preferred `AdmissionReview`
     // versions the Webhook expects. API server will try to use first version in
     // the list which is supports. If none of the versions specified in this list
     // supported by API server, validation will fail for this object.
     // This field is required and cannot be empty.
     AdmissionReviewVersions []string `json:"admissionReviewVersions" protobuf:"bytes,9,rep,name=admissionReviewVersions"`
}
```

V1beta1 API looks like this:

```golang
type Webhook struct {
    ...

     // AdmissionReviewVersions is an ordered list of preferred `AdmissionReview`
     // versions the Webhook expects. API server will try to use first version in
     // the list which is supports. If none of the versions specified in this list
     // supported by API server, validation will fail for this object.
     // Default to `['v1beta1']`.
     // +optional
     AdmissionReviewVersions []string `json:"admissionReviewVersions,omitempty" protobuf:"bytes,9,rep,name=admissionReviewVersions"`
}
```

## V1 API

The full v1 API will be look like this:

```golang
package v1

type ScopeType string

const (
     // ClusterScope means that scope is limited to cluster objects.
     ClusterScope ScopeType = "Cluster"
     // NamespacedScope means that scope is limited to namespaced objects.
     NamespacedScope ScopeType = "Namespaced"
)

// Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended
// to make sure that all the tuple expansions are valid.
type Rule struct {
     // APIGroups is the API groups the resources belong to. '*' is all groups.
     // If '*' is present, the length of the slice must be one.
     // Required.
     APIGroups []string `json:"apiGroups,omitempty" protobuf:"bytes,1,rep,name=apiGroups"`

     // APIVersions is the API versions the resources belong to. '*' is all versions.
     // If '*' is present, the length of the slice must be one.
     // Required.
     APIVersions []string `json:"apiVersions,omitempty" protobuf:"bytes,2,rep,name=apiVersions"`

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
     // If wildcard is present, the validation rule will ensure resources do not
     // overlap with each other.
     //
     // Depending on the enclosing object, subresources might not be allowed.
     // Required.
     Resources []string `json:"resources,omitempty" protobuf:"bytes,3,rep,name=resources"`

     // Scope specifies the scope of this rule. If unespecified, the scope is
     // not limited.
     //
     // +optional
     Scope ScopeType `json:"scope,omitempty" protobuf:"bytes,3,opt,name=scope"`
}

type FailurePolicyType string

const (
     // Ignore means that an error calling the webhook is ignored.
     Ignore FailurePolicyType = "Ignore"
     // Fail means that an error calling the webhook causes the admission to fail.
     Fail FailurePolicyType = "Fail"
)

type SideEffectClass string

const (
     // SideEffectClassUnknown means that no information is known about the side effects of calling the webhook.
     // If a request with the dry-run attribute would trigger a call to this webhook, the request will instead fail.
     SideEffectClassUnknown SideEffectClass = "Unknown"
     // SideEffectClassNone means that calling the webhook will have no side effects.
     SideEffectClassNone SideEffectClass = "None"
     // SideEffectClassSome means that calling the webhook will possibly have side effects.
     // If a request with the dry-run attribute would trigger a call to this webhook, the request will instead fail.
     SideEffectClassSome SideEffectClass = "Some"
     // SideEffectClassNoneOnDryRun means that calling the webhook will possibly have side effects, but if the
     // request being reviewed has the dry-run attribute, the side effects will be suppressed.
     SideEffectClassNoneOnDryRun SideEffectClass = "NoneOnDryRun"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
type ValidatingWebhookConfiguration struct {
     metav1.TypeMeta `json:",inline"`
     // Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
     // +optional
     metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
     // Webhooks is a list of webhooks and the affected resources and operations.
     // +optional
     // +patchMergeKey=name
     // +patchStrategy=merge
     Webhooks []Webhook `json:"webhooks,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=Webhooks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.
type ValidatingWebhookConfigurationList struct {
     metav1.TypeMeta `json:",inline"`
     // Standard list metadata.
     // More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
     // +optional
     metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
     // List of ValidatingWebhookConfiguration.
     Items []ValidatingWebhookConfiguration `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.
type MutatingWebhookConfiguration struct {
     metav1.TypeMeta `json:",inline"`
     // Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
     // +optional
     metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
     // Webhooks is a list of webhooks and the affected resources and operations.
     // +optional
     // +patchMergeKey=name
     // +patchStrategy=merge
     Webhooks []Webhook `json:"webhooks,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=Webhooks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
type MutatingWebhookConfigurationList struct {
     metav1.TypeMeta `json:",inline"`
     // Standard list metadata.
     // More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
     // +optional
     metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
     // List of MutatingWebhookConfiguration.
     Items []MutatingWebhookConfiguration `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// Webhook describes an admission webhook and the resources and operations it applies to.
type Webhook struct {
     // The name of the admission webhook.
     // Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where
     // "imagepolicy" is the name of the webhook, and kubernetes.io is the name
     // of the organization.
     // Required.
     Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

     // ClientConfig defines how to communicate with the hook.
     // Required
     ClientConfig WebhookClientConfig `json:"clientConfig" protobuf:"bytes,2,opt,name=clientConfig"`

     // Rules describes what operations on what resources/subresources the webhook cares about.
     // The webhook cares about an operation if it matches _any_ Rule.
     // However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks
     // from putting the cluster in a state which cannot be recovered from without completely
     // disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called
     // on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
     Rules []RuleWithOperations `json:"rules,omitempty" protobuf:"bytes,3,rep,name=rules"`

     // FailurePolicy defines how unrecognized errors from the admission endpoint are handled -
     // allowed values are Ignore or Fail. Defaults to Ignore.
     // +optional
     FailurePolicy *FailurePolicyType `json:"failurePolicy,omitempty" protobuf:"bytes,4,opt,name=failurePolicy,casttype=FailurePolicyType"`

     // NamespaceSelector decides whether to run the webhook on an object based
     // on whether the namespace for that object matches the selector. If the
     // object itself is a namespace, the matching is performed on
     // object.metadata.labels. If the object is another cluster scoped resource,
     // it never skips the webhook.
     //
     // For example, to run the webhook on any objects whose namespace is not
     // associated with "runlevel" of "0" or "1";  you will set the selector as
     // follows:
     // "namespaceSelector": {
     //   "matchExpressions": [
     //     {
     //       "key": "runlevel",
     //       "operator": "NotIn",
     //       "values": [
     //      "0",
     //      "1"
     //       ]
     //     }
     //   ]
     // }
     //
     // If instead you want to only run the webhook on any objects whose
     // namespace is associated with the "environment" of "prod" or "staging";
     // you will set the selector as follows:
     // "namespaceSelector": {
     //   "matchExpressions": [
     //     {
     //       "key": "environment",
     //       "operator": "In",
     //       "values": [
     //      "prod",
     //      "staging"
     //       ]
     //     }
     //   ]
     // }
     //
     // See
     // https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
     // for more examples of label selectors.
     //
     // Default to the empty LabelSelector, which matches everything.
     // +optional
     NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty" protobuf:"bytes,5,opt,name=namespaceSelector"`

     // SideEffects states whether this webhookk has side effects.
     // Acceptable values are: Unknown, None, Some, NoneOnDryRun
     // Webhooks with side effects MUST implement a reconciliation system, since a request may be
     // rejected by a future step in the admission change and the side effects therefore need to be undone.
     // Requests with the dryRun attribute will be auto-rejected if they match a webhook with
     // sideEffects == Unknown or Some. Defaults to Unknown.
     // +optional
     SideEffects *SideEffectClass `json:"sideEffects,omitempty" protobuf:"bytes,6,opt,name=sideEffects,casttype=SideEffectClass"`

     // ObjectSelector decides whether to run the webhook on an object based
     // on whether the object.metadata.labels matches the selector. An object is
     // matches both NamespaceSelector and ObjectSelector.
     //
     // See
     // https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
     // for more examples of label selectors.
     //
     // Default to the empty LabelSelector, which matches everything.
     // +optional
     ObjectSelector *metav1.LabelSelector `json:"objectSelector,omitempty" protobuf:"bytes,7,opt,name=objectSelector"`

     // TimeoutSeconds specifies the timeout for this webhook. After the timeout passes,
     // the webhook call will be ignored or the API call will fail based on the
     // failure policy.
     // The timeout value should be between 1 and 30 seconds.
     // Default to 10 seconds.
     // +optional
     TimeoutSeconds *int32 `json:"timeout,omitempty" protobuf:"varint,8,opt,name=timeout"`

     // AdmissionReviewVersions is an ordered list of preferred `AdmissionReview`
     // versions the Webhook expects. API server will try to use first version in
     // the list which is supports. If none of the versions specified in this list
     // supported by API server, validation will fail for this object.
     // This field is required and cannot be empty.
     AdmissionReviewVersions []string `json:"admissionReviewVersions" protobuf:"bytes,9,rep,name=admissionReviewVersions"`
}

// RuleWithOperations is a tuple of Operations and Resources. It is recommended to make
// sure that all the tuple expansions are valid.
type RuleWithOperations struct {
     // Operations is the operations the admission hook cares about - CREATE, UPDATE, or *
     // for all operations.
     // If '*' is present, the length of the slice must be one.
     // Required.
     Operations []OperationType `json:"operations,omitempty" protobuf:"bytes,1,rep,name=operations,casttype=OperationType"`
     // Rule is embedded, it describes other criteria of the rule, like
     // APIGroups, APIVersions, Resources, etc.
     Rule `json:",inline" protobuf:"bytes,2,opt,name=rule"`
}

type OperationType string

// The constants should be kept in sync with those defined in k8s.io/kubernetes/pkg/admission/interface.go.
const (
    OperationAll OperationType = "*"
     Create       OperationType = "CREATE"
     Update       OperationType = "UPDATE"
     Delete       OperationType = "DELETE"
     Connect      OperationType = "CONNECT"
)

// WebhookClientConfig contains the information to make a TLS
// connection with the webhook
type WebhookClientConfig struct {
     // `url` gives the location of the webhook, in standard URL form
     // (`scheme://host:port/path`). Exactly one of `url` or `service`
     // must be specified.
     //
     // The `host` should not refer to a service running in the cluster; use
     // the `service` field instead. The host might be resolved via external
     // DNS in some apiservers (e.g., `kube-apiserver` cannot resolve
     // in-cluster DNS as that would be a layering violation). `host` may
     // also be an IP address.
     //
     // Please note that using `localhost` or `127.0.0.1` as a `host` is
     // risky unless you take great care to run this webhook on all hosts
     // which run an apiserver which might need to make calls to this
     // webhook. Such installs are likely to be non-portable, i.e., not easy
     // to turn up in a new cluster.
     //
     // The scheme must be "https" and the URL must begin with "https://"
     // except for localhost in any form (e.g. localhost, 127.0.0.1, ::1).
     //
     // A path is optional, and if present may be any string permissible in
     // a URL. You may use the path to pass an arbitrary string to the
     // webhook, for example, a cluster identifier.
     //
     // Attempting to use a user or basic auth e.g. "user:password@" is not
     // allowed. Fragments ("#...") and query parameters ("?...") are not
     // allowed, either.
     //
     // +optional
     URL *string `json:"url,omitempty" protobuf:"bytes,3,opt,name=url"`

     // `service` is a reference to the service for this webhook. Either
     // `service` or `url` must be specified.
     //
     // If the webhook is running within the cluster, then you should use `service`.
     //
     // +optional
     Service *ServiceReference `json:"service,omitempty" protobuf:"bytes,1,opt,name=service"`

     // `caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate.
     // If unspecified, system trust roots on the apiserver are used.
     // +optional
     CABundle []byte `json:"caBundle,omitempty" protobuf:"bytes,2,opt,name=caBundle"`
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
     // `namespace` is the namespace of the service.
     // Required
     Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
     // `name` is the name of the service.
     // Required
     Name string `json:"name" protobuf:"bytes,2,opt,name=name"`

     // `path` is an optional URL path which will be sent in any request to
     // this service.
     // +optional
     Path *string `json:"path,omitempty" protobuf:"bytes,3,opt,name=path"`

     // If specified, the port on the service that hosting webhook.
     // Default to 443 for backward compatibility.
     // +optional
     Port *int32 `json:"port,omitempty" protobuf:"varint,4,opt,name=port"`
}
```

## V1beta1 changes

All of the proposed changes will be added to v1beta1 for backward compatibility and roundtrip-ability between different versions of Webhook Configuration API objects. The only difference would be:

- Default Value for `timeoutSeconds` field will be 30 seconds for `v1beta1`.
- `AdmissionReviewVersions` list is optional in v1beta1 and defaulted to `['v1beta1']` while required in `v1`.

## Validations

These set of new validation will be applied to both v1 and v1beta1:

- `Scope` field can only have `Cluster` or `Namespaced` values

- `https` validation is relaxed for localhost (and similar urls like `127.0.0.1`)

- `Timeout` field should be between 1 and 30 seconds.

- `AdmissionReviewVersions` list should have at least one version supported by the API Server serving it. Note that for downgrade compatibility, Webhook authors should always support as many `AdmissionReview` versions as possible.
