# Aggregated API Servers

## Abstract

We want to divide the single monolithic API server into multiple aggregated
servers. Anyone should be able to write their own aggregated API server to expose APIs they want.
Cluster admins should be able to expose new APIs at runtime by bringing up new
aggregated servers.

## Motivation

* Extensibility: We want to allow community members to write their own API
  servers to expose APIs they want. Cluster admins should be able to use these
  servers without having to require any change in the core kubernetes
  repository.
* Unblock new APIs from core kubernetes team review: A lot of new API proposals
  are currently blocked on review from the core kubernetes team. By allowing
  developers to expose their APIs as a separate server and enabling the cluster
  admin to use it without any change to the core kubernetes repository, we
  unblock these APIs.
* Place for staging experimental APIs: New APIs can be developed in separate
  aggregated servers, and installed only by those willing to take the risk of
  installing an experimental API. Once they are stable, it is then easy to
  package them up for installation in other clusters.
* Ensure that new APIs follow kubernetes conventions: Without the mechanism
  proposed here, community members might be forced to roll their own thing which
  may or may not follow kubernetes conventions.

## Goal

* Developers should be able to write their own API server and cluster admins
  should be able to add them to their cluster, exposing new APIs at runtime. All
  of this should not require any change to the core kubernetes API server.
* These new APIs should be seamless extensions of the core kubernetes APIs (ex:
  they should be operated upon via kubectl).

## Non Goals

The following are related but are not the goals of this specific proposal:
* Make it easy to write a kubernetes API server.

## High Level Architecture

There will be a new component, `kube-aggregator`, which has these responsibilities:
* Provide an API for registering API servers.
* Summarize discovery information from all the servers.
* Proxy client requests to individual servers.

The reverse proxy is provided for convenience. Clients can discover server URLs
using the summarized discovery information and contact them directly. Simple
clients can always use the proxy and don't need to know that under the hood
multiple apiservers are running.

Wording note: When we say "API servers" we really mean groups of apiservers,
since any individual apiserver is horizontally replicable. Similarly,
kube-aggregator itself is horizontally replicable.

## Operational configurations

There are two configurations in which it makes sense to run `kube-aggregator`.
 1. In **test mode**/**single-user mode**. An individual developer who wants to test
    their own apiserver could run their own private copy of `kube-aggregator`,
    configured such that only they can interact with it. This allows for testing
    both `kube-aggregator` and any custom apiservers without the potential for
    causing any collateral damage in the rest of the cluster. Unfortunately, in
    this configuration, `kube-aggregator`'s built in proxy will lack the client
    cert that allows it to perform authentication that the rest of the cluster
    will trust, so its functionality will be somewhat limited.
 2. In **gateway mode**. The `kube-apiserver` will embed the `kube-aggregator` component
    and it will function as the official gateway to the cluster, where it aggregates
    all of the apiservers the cluster administer wishes to provide.

### Constraints

* Unique API group versions across servers: Each API server (and groups of servers, in HA)
  should expose unique API group versions. So, for example, you can serve
  `api.mycompany.com/v1` from one apiserver and the replacement
  `api.mycompany.com/v2` from another apiserver while you update clients. But
  you can't serve `api.mycompany.com/v1/frobbers` and
  `api.mycompany.com/v1/grobinators` from different apiservers. This restriction
  allows us to limit the scope of `kube-aggregator` to a manageable level.
* Follow API conventions: APIs exposed by every API server should adhere to [kubernetes API
  conventions](/contributors/devel/sig-architecture/api-conventions.md).
* Support discovery API: Each API server should support the kubernetes discovery API
  (list the supported groupVersions at `/apis` and list the supported resources
  at `/apis/<groupVersion>/`)
* No bootstrap problem: The core kubernetes apiserver must not depend on any
  other aggregated server to come up. Non-core apiservers may use other non-core
  apiservers, but must not fail in their absence.

## Component Dependency Order

`kube-aggregator` is not part of the core `kube-apiserver`.
The dependency order (for the cluster gateway configuration) looks like this:
 1. `etcd`
 2. `kube-apiserver`
 3. core scheduler, kubelet, service proxy (enough stuff to create a pod, run it on a node, and find it via service)
 4. `kubernetes-aggregator` as a pod/service - default summarizer and proxy
 5. controllers
 6. other API servers and their controllers
 7. clients, web consoles, etc

Nothing below the `kubernetes-aggregator` can rely on the aggregator or proxy
being present.  `kubernetes-aggregator` should be runnable as a pod backing a
service in a well-known location.  Something like `api.kube-public.svc` or
similar seems appropriate since we'll want to allow network traffic to it from
every other namespace in the cluster.  We recommend using a dedicated namespace,
since compromise of that namespace will expose the entire cluster: the
proxy has the power to act as any user against any API server.

## Implementation Details

### Summarizing discovery information

We can have a very simple Go program to summarize discovery information from all
servers. Cluster admins will register each aggregated API server (its baseURL and swagger
spec path) with the proxy. The proxy will summarize the list of all group versions
exposed by all registered API servers with their individual URLs at `/apis`.

### Reverse proxy

We can use any standard reverse proxy server like nginx or extend the same Go program that
summarizes discovery information to act as reverse proxy for all aggregated servers.

Cluster admins are also free to use any of the multiple open source API management tools
(for example, there is [Kong](https://getkong.org/), which is written in lua and there is
[Tyk](https://tyk.io/), which is written in Go). These API management tools
provide a lot more functionality like: rate-limiting, caching, logging,
transformations and authentication.
In future, we can also use ingress. That will give cluster admins the flexibility to
easily swap out the ingress controller by a Go reverse proxy, nginx, haproxy
or any other solution they might want.

`kubernetes-aggregator` uses a simple proxy implementation alongside its discovery information
which supports connection upgrade (for `exec`, `attach`, etc) and runs with delegated
authentication and authorization against the core `kube-apiserver`.  As a proxy, it adds
complete user information, including user, groups, and "extra" for backing API servers.

### Storage

Each API server is responsible for storing their resources. They can have their
own etcd or can use kubernetes server's etcd using [third party
resources](extending-api.md#adding-custom-resources-to-the-kubernetes-api-server).

### Health check

Kubernetes server's `/api/v1/componentstatuses` will continue to report status
of master components that it depends on (scheduler and various controllers).
Since clients have access to server URLs, they can use that to do
health check of individual servers.
In future, if a global health check is required, we can expose a health check
endpoint in the proxy that will report the status of all aggregated api servers
in the cluster.

### Auth

Since the actual server which serves client's request can be opaque to the client,
all API servers need to have homogeneous authentication and authorisation mechanisms.
All API servers will handle authn and authz for their resources themselves.
The current authentication infrastructure allows token authentication delegation to the
core `kube-apiserver` and trust of an authentication proxy, which can be fulfilled by
`kubernetes-aggregator`.

#### Server Role Bootstrapping

External API servers will often have to provide roles for the resources they
provide to other API servers in the cluster.  This will usually be RBAC
clusterroles, RBAC clusterrolebindings, and apiaggregation types to describe
their API server.  The external API server should *never* attempt to
self-register these since the power to mutate those resources provides the
power to destroy the cluster.  Instead, there are two paths:
 1. the easy path - In this flow, the API server supports a `/bootstrap/<group>` endpoint
    which provides the resources that can be piped to a `kubectl create -f` command a cluster-admin
    can use those endpoints to prime other servers.
 2. the reliable path - In a production cluster, you generally want to know, audit, and
    track the resources required to make your cluster work.  In these scenarios, you want
    to have the API resource list ahead of time.  API server authors can provide a template.

Nothing stops an external API server from supporting both.

### kubectl

kubectl will talk to `kube-aggregator`'s discovery endpoint and use the discovery API to
figure out the operations and resources supported in the cluster.
We will need to make kubectl truly generic. Right now, a lot of operations
(like get, describe) are hardcoded in the binary for all resources. A future
proposal will provide details on moving those operations to server.

Note that it is possible for kubectl to talk to individual servers directly in
which case proxy will not be required at all, but this requires a bit more logic
in kubectl. We can do this in future, if desired.

### Handling global policies

Now that we have resources spread across multiple API servers, we need to
be careful to ensure that global policies (limit ranges, resource quotas, etc) are enforced.
Future proposals will improve how this is done across the cluster.

#### Namespaces

When a namespaced resource is created in any of the aggregated server, that
server first needs to check with the kubernetes server that:

* The namespace exists.
* User has authorization to create resources in that namespace.
* Resource quota for the namespace is not exceeded.

To prevent race conditions, the kubernetes server might need to expose an atomic
API for all these operations.

While deleting a namespace, kubernetes server needs to ensure that resources in
that namespace maintained by other servers are deleted as well. We can do this
using resource [finalizers](/contributors/design-proposals/architecture/namespaces.md#finalizers). Each server
will add themselves in the set of finalizers before they create a resource in
the corresponding namespace and delete all their resources in that namespace,
whenever it is to be deleted (kubernetes API server already has this code, we
will refactor it into a library to enable reuse).

Future proposal will talk about this in more detail and provide a better
mechanism.

#### Limit ranges and resource quotas

kubernetes server maintains [resource quotas](/contributors/design-proposals/resource-management/admission_control_resource_quota.md) and
[limit ranges](/contributors/design-proposals/resource-management/admission_control_limit_range.md) for all resources.
Aggregated servers will need to check with the kubernetes server before creating any
resource.

## Methods for running on hosted kubernetes clusters

Where "hosted" means the cluster users have very limited or no permissions to
change the control plane installation, for example on GKE, where it is managed
by Google. There are three ways of running on such a cluster:.

 1. `kube-aggregator` will run in the single-user / test configuration on any
    installation of Kubernetes, even if the user starting it only has permissions
    in one namespace.
 2. Just like 1 above, if all of the users can agree on a location, then they
    can make a public namespace and run a copy of `kube-aggregator` in that
    namespace for everyone. The downside of running like this is that none of the
    cluster components (controllers, nodes, etc) would be going through this
    kube-aggregator.
 3. The hosted cluster provider can integrate `kube-aggregator` into the
    cluster. This is the best configuration, but it may take a quarter or two after
    `kube-aggregator` is ready to go for providers to complete this integration.

## Alternatives

There were other alternatives that we had discussed.

* Instead of adding a proxy in front, let the core kubernetes server provide an
  API for other servers to register themselves. It can also provide a discovery
  API which the clients can use to discover other servers and then talk to them
  directly. But this would have required another server API a lot of client logic as well.
* Validating aggregated servers: We can validate new servers when they are registered
  with the proxy, or keep validating them at regular intervals, or validate
  them only when explicitly requested, or not validate at all.
  We decided that the proxy will just assume that all the servers are valid
  (conform to our api conventions). In future, we can provide conformance tests.

## Future Work

* Validate servers: We should have some conformance tests that validate that the
  servers follow kubernetes api-conventions.
* Provide centralised auth service: It is very hard to ensure homogeneous auth
  across multiple aggregated servers, especially in case of hosted clusters
  (where different people control the different servers). We can fix it by
  providing a centralised authentication and authorization service which all of
  the servers can use.

