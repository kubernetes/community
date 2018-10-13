---
kep-number: 31
title: Enabling clients to tell if resource endpoints serve the same set of objects
authors:
  - "@xuchao"
owning-sig: sig-api-machinery
reviewers:
  - "@deads2k"
  - "@lavalamp"
approvers:
  - "@deads2k"
  - "@lavalamp"
creation-date: 2018-10-12
last-updated: 2018-xx-xx
status: provisional
---

# Enabling clients to tell if resource endpoints server the same set of objects

## Table of Contents

      * [Summary](#summary)
      * [Motivation](#motivation)
            * [Breaking forward compatibility](#breaking-forward-compatibility)
            * [Inaccurate resource quota counting](#inaccurate-resource-quota-counting)
            * [Inefficiency of API-discovery based dynamic clients](#inefficiency-of-api-discovery-based-dynamic-clients)
      * [Goals](#goals)
      * [Proposal](#proposal)
         * [API changes to the discovery API](#api-changes-to-the-discovery-api)
         * [Implementation details](#implementation-details)
         * [Risks and Mitigations](#risks-and-mitigations)
      * [Graduation Criteria](#graduation-criteria)
      * [Alternatives](#alternatives)

## Summary

We propose to expand the discovery API to enable clients to tell if resource
endpoints (e.g., `extensions/v1beta1/replicaset` and `apps/v1/replicaset`) are
referring the same set of objects.

## Motivation

The inability of telling if resource endpoints refer the same set of objects
leads to a few issues.

#### Breaking forward compatibility

For example, if we configure an admission webhook to validate all replicasets,
one would set the webhook configuration this way:
```yaml
rules:
  - apiGroups:
    - "extensions,apps"
    apiVersions:
    - "*"
    resources:
    - replicasets
```

If in a future Kubernetes release, replicasets are accessible via
`fancy-apps/v1` as well, requests sent to `fancy-apps/v1` will be left
unchecked. Note that the admins of cluster upgrades are not necessarily the
admins of the admission webhooks, so it is not always possible to coordinate
cluster upgrades with webhook configuration upgrades.

#### Inaccurate resource quota counting

Although `extensions/v1beta1/replicasets` and `apps/v1/replicasets` refer to the
same set of objects, they have separate resource quotas. User can bypass the 
resource quota of one endpoint by using the other endpoint.

#### Inefficiency of API-discovery based dynamic clients

For example, the Kubernetes [garbage collector][] watches all discoverable
resource endpoints and thus wasting bandwidth. As another example, the [storage
migrator][] migrates the same objects multiple times if they are served via
multiple endpoints.

[garbage collector]:https://github.com/kubernetes/kubernetes/tree/master/pkg/controller/garbagecollector
[storage migrator]:https://github.com/kubernetes-sigs/kube-storage-version-migrator


## Goals

The successful mechanism should
* enable clients to tell if two resource endpoints refer to the same objects.
* prevent clients from relying on the implementation details of the mechanism.
* work for all resources, including built-in resources, custom resources, and
  aggregated resources.

## Proposal

### API changes to the discovery API

We add a new field `InternalID` to the [APIResource][] type. We carefully avoid
mentioning any implementation details in the field name or in the comment.

[APIResource]:https://github.com/kubernetes/kubernetes/blob/f22334f14d92565ec3ff9d4ff2b995eae9af622a/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L881-L905

```golang
type APIResource struct {
        // A unique ID of the set of objects this resource endpoint represents.
        // Resources have the same internalID if and only if they refer to the
        // same set of objects.
        InternalID string
	// These are the existing fields.
	Name string
	SingularName string
	Namespaced bool
	Group string
	Version string
	Kind string
	Verbs Verbs
	ShortNames []string
	Categories []string
}
```

### Implementation details

For built-in resources, their `internalID`s are set to `SHA256(<etcd key
prefix>)`. For example, for both `extensions/v1beta1/replicasets` and
`apps/v1/replicasets`, the etcd key prefix is `/registry/replicasets`. Hashing
the prefix is to encourage the clients to only test the equality of the hashed
values, instead of relying on the absolute value.

For custom resources, their `internalID`s are also set to `SHA256(<etcd key
prefix>)`. In the current implementation, the etcd key prefix for a custom
resource is `/registry/<crd.spec.group>/<crd.spec.names.plural>`.

For aggregated resources, because their discovery doc is fully controlled
by the aggregated apiserver, the kube-apiserver has no means to validate their
`internalID`. If the server is implemented with the generic apiserver library,
the `internalID` will be `SHA256(<etcd key prefix>)`.

For subresources, the `internalID` field is left empty. If the main resource can
be accessed via another resource endpoint, then so does the subresource.

For non-persistent resources like `tokenReviews` or `subjectAccessReviews`,
though the objects are not persisted, the [forward compatibility][] motivation
still applies, e.g., admins might configure the admission webhooks to intercept
requests sent to all endpoints. Thus, the `internalID` cannot be left empty, it
will be set to `SHA256(<the would-be etcd key prefix>)`, e.g., for
`tokenReviews`, it's `SHA256(/registry/tokenreviews)`. 

[forward compatibility]:#breaking-forwards-compatibility

### Risks and Mitigations

In the future, the "etcd key prefix" might not be sufficient to uniquely
identify a set of objects. We can always add more factors to the `internalID` to
ensure their uniqueness. It does not break backwards compatibility because the
`internalID` is opaque.

Another risk is that an aggregated apiserver accidentally reports `internalID`
that's identical to the built-in or the custom resources, this will confuse
clients. Because the kube-apiserver has zero control over the discovery doc of
aggregated resources, it cannot do any validation to prevent this kind of error.
It will be aggregated apiserver provider's responsibility to prevent such errors.

## Graduation Criteria

Because the discovery API is a GA feature, to add the `internalID` field, it
needs to be of GA quality.

## Alternatives
1. Adding to the discovery API a reference to the canonical endpoint. For
   example, in the discovery API, `extensions/v1beta1/replicasets` reports
   `apps/v1/replicasets` as the canonical endpoint. This approach is similar to
   `internalID` proposal, but because the resource names are explicitly exposed,
   clients might use the information in unintended ways.

2. Serving a list of all sets of aliasing resources via a new API. Aggregated
   apiservers make such a design complex. For example, we will need to design how
   the aggregated apiserver registers its resource aliases. 
