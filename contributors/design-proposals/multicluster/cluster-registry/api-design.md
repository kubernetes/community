# Cluster Registry API

@perotinus, @madhusudancs

Original draft: 08/16/2017

**Reviewed** in SIG multi-cluster meeting on 8/29

*This doc is a Markdown conversion of the original Cluster Registry API
[Google doc](https://docs.google.com/document/d/1Oi9EO3Jwtp69obakl-9YpLkP764GZzsz95XJlX1a960/edit).
That doc is deprecated, and this one is canonical; however, the old doc will be
preserved so as not to lose comment and revision history that it contains.*
## Table of Contents

-   [Purpose](#purpose)
-   [Motivating use cases](#motivating-use-cases)
-   [API](#api)
-   [Authorization-based filtering of the list of clusters](#authorization-based-filtering-of-the-list-of-clusters)
-   [Status](#status)
-   [Auth](#auth)
-   [Key differences vs existing Federation API `Cluster` object](#key-differences-vs-existing-federation-api-cluster-object)
-   [Open questions](#open-questions)

## Purpose

The cluster registry API is intended to provide a common abstraction for other
tools that will perform operations on multiple clusters. It provides an
interface to a list of objects that will store metadata about clusters that can
be used by other tools. The cluster registry implementation is meant to remain
simple: we believe there is benefit in defining a common layer that can be used
by many different tools to solve different kinds of multi-cluster problems.

It may be helpful to consider this API as an extension of the `kubeconfig` file.
The `kubeconfig` file contains a list of clusters with the auth data necessary
for kubectl to access them; the cluster registry API intends to provide this
data, plus some additional useful metadata, from a remote location instead of
from the user's local machine.

## Motivating use cases

These were presented at the SIG-Federation F2F meeting on 8/4/17
([Atlassian](https://docs.google.com/document/d/1PH859COCWSkRxILrQd6wDdYLGJaBtWQkSN3I-Lnam3g/edit#heading=h.suxgoa67n1aw),
[CoreOS](https://docs.google.com/presentation/d/1InJagQNOxqA0ftK0peJLzyEFU2IZEXrJprDN6fcleMg/edit#slide=id.p),
[Google](https://docs.google.com/presentation/d/1Php_HnHI-Sy20ieyd_jBgr7XTs0fKT0Cq9z6dC4zOMc/),
[RedHat](https://docs.google.com/presentation/d/1dExjeSQTXI8_k00nqXRkSIFPTkzAzUTFtETU4Trg5yw/edit#slide=id.p)).
Each of the use cases presented assumes the ability to access a registry of
clusters, and so all are valid motivating use cases for the cluster registry
API. Note that these use cases will generally require more tooling than the
cluster registry itself. The cluster registry API will support what these other
tools will need in order to operate, but will not intrinsically support these
use cases.

-   Consistent configuration across clusters/replication of resources
-   Federated Ingress: load balancing across multiple clusters, potentially
    geo-aware
-   Multi-cluster application distribution, with policy/balancing
-   Disaster recovery/failover
-   Human- and tool- parseable interaction with a list of clusters
-   Monitoring/health checking/status reporting for clusters, potentially with
    dashboards
-   Policy-based and jurisdictional placement of workloads

## API

This document defines the cluster registry API. It is an evolution of the
[current Federation cluster API](https://git.k8s.io/federation/apis/federation/types.go#L99),
and is designed more specifically for the "cluster registry" use case in
contrast to the Federation `Cluster` object, which was made for the
active-control-plane Federation.

The API is a Kubernetes-style REST API that supports the following operations:

1.  `POST` - to create new objects.
1.  `GET` - to retrieve both lists and individual objects.
1.  `PUT` - to update or create an object.
1.  `DELETE` - to delete an object.
1.  `PATCH` - to modify the fields of an object.

Optional API operations:

1.  `WATCH` - to receive a stream of changes made to a given object. As `WATCH`
    is not a standard HTTP method, this operation will be implemented as `GET
    /<resource>&watch=true`. We believe that it's not always necessary to
    support WATCH for this API. Implementations can choose to support or not
    support this operation. An implementation that does not support the
    operation should return HTTP error 405, StatusMethodNotAllowed, per the
    [relevant Kubernetes API conventions](/contributors/devel/sig-architecture/api-conventions.md#error-codes).

We also intend to support a use case where the server returns a file that can be
stored for later use. We expect this to be doable with the standard API
machinery; and if the API is implemented not using the Kubernetes API machinery,
that the returned file must be interoperable with the response from a Kubernetes
API server.

[The API](https://git.k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1/types.go)
is defined in the cluster registry repo, and is not replicated here in order to
avoid mismatches.

All top-level objects that define resources in Kubernetes embed a
`meta.ObjectMeta` that in-turn contains a number of fields. All the fields in
that struct are potentially useful with the exception of the `ClusterName` and
the `Namespace` fields. Having a `ClusterName` field alongside a `Name` field in
the cluster registry API will be confusing to our users. Therefore, in the
initial API implementation, we will add validation logic that rejects `Cluster`
objects that contain a value for the `ClusterName` field. The `Cluster` object's
`Namespace` field will be disabled by making the object be root scoped instead
of namespace scoped.

The `Cluster` object will have `Spec` and `Status` fields, following the
[Kubernetes API conventions](/contributors/devel/sig-architecture/api-conventions.md#spec-and-status).
There was argument in favor of a `State` field instead of `Spec` and `Status`
fields, since the `Cluster` in the registry does not necessarily hold a user's
intent about the cluster being represented, but instead may hold descriptive
information about the cluster and information about the status of the cluster;
and because the cluster registry provides no controller that performs
reconciliation on `Cluster` objects. However, after
[discussion with SIG-arch](https://groups.google.com/forum/#!topic/kubernetes-sig-architecture/ptK2mVtha38),
the decision was made in favor of spec and status.

## Authorization-based filtering of the list of clusters

The initial version of the cluster registry supports a cluster list API that
does not take authorization rules into account. It returns a list of clusters
similar to how other Kubernetes List APIs list the objects in the presence of
RBAC rules. A future version of this API will take authorization rules into
account and only return the subset of clusters a user is authorized to access in
the registry.

## Status

There are use cases for the cluster registry that call for storing status that
is provided by more active controllers, e.g. health checks and cluster capacity.
At this point, these use cases are not as well-defined as the use cases that
require a data store, and so we do not intend to propose a complete definition
for the `ClusterStatus` type. We recognize the value of conventions, so as these
use cases become more clearly defined, the API of the `ClusterStatus` will be
extended appropriately.

## Auth

The cluster registry API will not provide strongly-typed objects for returning
auth info. Instead, it will provide a generic type that clients can use as they
see fit. This is intended to mirror what `kubectl` does with its
[AuthProviderConfig](https://git.k8s.io/client-go/tools/clientcmd/api/types.go#L144).
As open standards are developed for cluster auth, the API can be extended to
provide first-class support for these. We want to avoid baking non-open
standards into the API, and so having to support potentially a multiplicity of
them as they change. The cluster registry itself is not intended to be a
credential store, but instead to provide "pointers" that will provide the
information needed by callers to authenticate to a cluster. There is some more
context
[here](https://docs.google.com/a/google.com/document/d/1cxKV4Faywsn_to49csN0S0TZLYuHgExusEsmgKQWc28/edit?usp=sharing).

## Key differences vs existing Federation API `Cluster` object

-   Active controller is not required; the registry can be used without any
    controllers
-   `WATCH` support is not required

## Open questions

All open questions have been
[migrated](https://github.com/kubernetes/cluster-registry/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20%22Migrated%20from%20the%20Cluster%22)
to issues in the cluster registry repo.
