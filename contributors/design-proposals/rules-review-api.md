# "What can I do?" API

Author: Eric Chiang (eric.chiang@coreos.com)

## Overview

Currently, to determine if a user is authorized to perform a set of actions, that user has to query each action individually through a `SelfSubjectAccessReview`.

Beyond making the authorization layer hard to reason about, it means web interfaces such as the OpenShift Web Console, Tectonic Console, and Kubernetes Dashboard, have to perform individual calls for _every resource_ a page displays. There's no way for a user, or an application acting on behalf of a user, to ask for all the permissions a user can make in bulk. This makes its hard to build pages that are proactive about what's displayed or grayed out based on the end user's permissions. UIs can only handle 403 responses after a user has already performed a forbidden action.

This is a proposal to add authorization APIs that allow a client to determine what actions they can make within a namespace. We expect this API to be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. This API should NOT be used by external systems to drive their own authorization decisions, as this raises confused deputy, cache lifetime/revocation, and correctness concerns. The `*AccessReview` APIs remain the correct way to defer authorization decisions to the API server.

OpenShift adopted a [`RulesReview` API][openshift-rules-review] to accomplish this same goal, and this proposal is largely a port of that implementation.

[kubernetes/kubernetes#48051](https://github.com/kubernetes/kubernetes/pull/48051) implements most of this proposal.

## API additions

Add a top level type to the `authorization.k8s.io` API group called `SelfSubjectRulesReview`. This mirrors the existing `SelfSubjectAccessReview`.

```
type SelfSubjectRulesReview struct {
	metav1.TypeMeta

	Spec SelfSubjectRulesReviewSpec

	// Status is filled in by the server and represents the set of actions a user can perform.
	Status SubjectRulesReviewStatus
}

type SelfSubjectRulesReviewSpec struct {
	// Namespace to evaluate rules for. Required.
	Namespace string
}

type SubjectRulesReviewStatus struct {
	// ResourceRules is the list of actions the subject is allowed to perform on resources.
	// The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	ResourceRules []ResourceRule
	// NonResourceRules is the list of actions the subject is allowed to perform on non-resources.
	// The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
	NonResourceRules []NonResourceRule
	// EvaluationError can appear in combination with Rules. It indicates an error occurred during
	// rule evaluation, such as an authorizer that doesn't support rule evaluation, and that
	// ResourceRules and/or NonResourceRules may be incomplete.
	EvaluationError string
    // Incomplete indicates that the returned list is known to be incomplete.
    Incomplete bool
}
```

The `ResourceRules` and `NonResourceRules` rules are similar to the types use by RBAC and the internal authorization system.

```
# docstrings omitted for brevity. 
type ResourceRule struct {
	Verbs []string
	APIGroups []string
	Resources []string
	ResourceNames []string
}

type NonResourceRule struct {
	Verbs []string
	NonResourceURLs []string
}
```

All of these fields can include the string `*` to indicate all values are allowed.

### Differences from OpenShift: user extras vs. scopes

OpenShift `SelfSubjectRulesReviewSpec` takes a set of [`Scopes`][openshift-scopes]. This lets OpenShift clients use the API for queries such as _"what could I do if I provide this scope to limit my credentials?"_

 In core kube, scopes are replaced by "user extras" field, a map of opaque strings that can be used for implementation specific user data. Unlike OpenShift, where scopes are always used to restrict credential powers, user extras are commonly used to expand powers. For example, the proposed [Keystone authentiator][keystone-authn] used them to include additional roles and project fields.

Since user extras can be used to expand the power of users, instead of only restricting, this proposal argues that `SelfSubjectRulesReview` shouldn't let a client specify them like `Scopes`. It wouldn't be within the spirit of a `SelfSubject` resource to let a user determine information about other projects or roles.

This could hopefully be solved by introducing a `SubjectRulesReview` API to query the rules for any user. An aggregated API server could use the `SubjectRulesReview` to back an API resource that let a user provide restrictive user extras, such as scopes.

## Webhook authorizers

Some authorizers live external to Kubernetes through an API server webhook and wouldn't immediately support a rules review query.

To communicate with external authorizers, the following types will be defined to query the rules for an arbitrary user. This proposal does NOT propose adding these types to the API immediately, since clients can use user impersonation and a `SelfSubjectRulesReview` to accomplish something similar.

```
type SubjectRulesReview struct {
	metav1.TypeMeta

	Spec SubjectRulesReviewSpec

	// Status is filled in by the server and indicates the set of actions a user can perform.
	Status SubjectRulesReviewStatus
}

type SubjectRulesReviewSpec struct {
	// Namespace to evalue rules for. Required.
	Namespace string

        // User to be evaluated for.
	UID string
	User string
	Groups []string
	Extras map[string][]string
}
```

Currently, external authorizers are configured through the following API server flag and which POSTs a `SubjectAccessReview` to determine a user's access:

```
--authorization-webhook-config-file
```

The config file uses the kubeconfig format.

There are a few options to support a second kind of query.

* Add another webhook flag with a second config file.
* Introduce a [kubeconfig extension][kubeconfig-extension] that indicates the server can handle either a `SubjectRulesReview` or a `SubjectAccessReview`
* Introduce a second context in the kubeconfig for the `SubjectRulesReview`. Have some way of indicating which context for `SubjectRulesReview` and which is for `SubjectAccessReview`, for example by well-known context names for each.

The doc proposed adding a second webhook config for `RulesReview`, and not overloading the existing config passed to `--authorization-webhook-config-file`.

[openshift-rules-review]: https://github.com/openshift/origin/blob/v3.6.0/pkg/authorization/apis/authorization/types.go#L152
[openshift-scopes]: https://github.com/openshift/origin/blob/v3.6.0/pkg/authorization/apis/authorization/types.go#L164-L168
[keystone-authn]: https://github.com/kubernetes/kubernetes/pull/25624/files#diff-897f0cab87e784d9fc6813f04f128f62R40
[kubeconfig-extension]: https://github.com/kubernetes/client-go/blob/v3.0.0/tools/clientcmd/api/v1/types.go#L51
