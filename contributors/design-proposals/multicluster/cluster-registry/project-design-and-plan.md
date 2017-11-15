# Cluster registry design and plan

@perotinus

Updated: 11/2/17

*REVIEWED* in SIG-multicluster meeting on 10/24

*This doc is a Markdown conversion of the original Cluster registry design and
plan
[Google doc](https://docs.google.com/document/d/1bVvq9lDIbE-Glyr6GkSGWYkLb2cCNk9bR8LL7Wm-L6g/edit).
That doc is deprecated, and this one is canonical; however, the old doc will be
preserved so as not to lose comment and revision history that it contains.*

## Table of Contents

-   [Background](#background)
-   [Goal](#goal)
-   [Technical requirements](#technical-requirements)
    -   [Alpha](#alpha)
    -   [Beta](#beta)
    -   [Later](#later)
-   [Implementation design](#implementation-design)
    -   [Alternatives](#alternatives)
    -   [Using a CRD](#using-a-crd)
-   [Tooling design](#tooling-design)
    -   [User tooling](#user-tooling)
-   [Repository process](#repository-process)
-   [Release strategy](#release-strategy)
    -   [Version skew](#version-skew)
-   [Test strategy](#test-strategy)
-   [Milestones and timelines](#milestones-and-timelines)
    -   [Alpha (targeting late Q4 '17)](#alpha-targeting-late-q4-'17)
    -   [Beta (targeting mid Q1 '18)](#beta-targeting-mid-q1-'18)
    -   [Stable (targeting mid Q2 '18)](#stable-targeting-mid-q2-'18)
    -   [Later](#later)

## Background

SIG-multicluster has identified a cluster registry as being a key enabling
component for multi-cluster use cases. The goal of the SIG in this project is
that the API defined for the cluster registry become a standard for
multi-cluster tools. The API design is being discussed in
[a separate doc](https://drive.google.com/a/google.com/open?id=1Oi9EO3Jwtp69obakl-9YpLkP764GZzsz95XJlX1a960).
A working prototype of the cluster registry has been assembled in
[a new repository](https://github.com/kubernetes/cluster-registry).

## Goal

This document intends to lay out an initial plan for moving the cluster registry
from the prototype state through alpha, beta and eventually to a stable release.

## Technical requirements

These requirements are derived mainly from the output of the
[August 9th multi-cluster SIG meeting](https://docs.google.com/document/d/11cB3HK67BZUb7aNOCpQK8JsA8Na8F-F_6bFu77KrKT4/edit#heading=h.cuvqls7pl9qc).
However, they also derive (at least indirectly) from the results of the
[SIG F2F meeting ](https://docs.google.com/document/d/1HkVBSm9L9UJC2f3wfs_8zt1PJmv6iepdtJ2fmkCOHys/edit)
that took place earlier, as well as the use cases presented by various parties
at that meeting.

### Alpha

-   [API] Provides an HTTP server with a Kubernetes-style API for CRUD
    operations on a registry of clusters
    -   Provides a way to filter across clusters by label
    -   Supports annotations
-   [API] Provides information about each cluster's API server's
    location/endpoint
-   [API] Provides the ability to assign user-friendly names for clusters
-   [API] Provides pointers to authentication information for clusters' API
    servers
    -   This implies that it does not support storing credentials directly
-   [Implementation] Supports both independent and aggregated deployment models
    -   Supports delegated authn/authz in aggregated mode
    -   Supports integration into Federation via aggregated deployment model

### Beta

-   Supports providing a flat file of clusters for storage and later use
    -   This may be provided by the ecosystem rather than the registry API
        implementation directly
-   Supports independent authz for reading/writing clusters
-   `kubectl` integration with the cluster registry API server is first-class
    and on par with kubectl integration with core Kubernetes APIs
-   Supports grouping of clusters
-   Supports specifying and enforcing read and/or write authorization for groups
    of clusters in the registry
-   Working Federation integration
-   Supports status from active controllers

### Later

-   Supports an HA deployment strategy
-   Supports guarantees around immutability/identity of clusters in list
-   Version skew between various components is understood and supported skews
    are defined and tested

## Implementation design

The cluster registry will be implemented using the
[Kubernetes API machinery](https://github.com/kubernetes/apimachinery). The
cluster registry API server will be a fork and rework of the existing Federation
API server, scaled down and simplified to match the simpler set of requirements
for the cluster registry. It will use the
[apiserver](https://github.com/kubernetes/apiserver) library, plus some code
copied from the core Kubernetes repo that provides scaffolding for certain
features in the API server. This is currently implemented in a prototype form
in the [cluster-registry repo](https://github.com/kubernetes/cluster-registry).

The API will be implemented using the Kubernetes API machinery, as a new API
with two objects, `Cluster` and `ClusterList`. Other APIs may be added in the
future to support future use cases, but the intention is that the cluster
registry API server remain minimal and only provide the APIs that users of a
cluster registry would want for the cluster registry.

The cluster registry will not be responsible for storing secrets. It will
contain pointers to other secret stores which will need to be implemented
independently. The cluster registry API will not provide proxy access to
clusters, and will not interact with clusters on a user's behalf. Storing secret
information in the cluster registry will be heavily discouraged by its
documentation, and will be considered a misuse of the registry. This allows us
to sidestep the complexity of implementing a secure credential storage.

The expectation is that Federation and other programs will use the
cluster-registry as an aggregated API server rather than via direct code
integration. Therefore, the cluster registry will explicitly support being
deployed as an API server that can be aggregated by other API servers.

### Alternatives

#### Using a CRD

The cluster registry could be implemented as a CRD that is registered with a
Kubernetes API server. This implementation is lighter weight than running a full
API server. If desired, the administrator could then disable the majority of the
Kubernetes APIs for non-admin users and so make it appear as if the API server
only supports cluster objects. It should be possible for a user to migrate
without much effort from a CRD-based to an API-server-based implementation of
the cluster registry, but the cluster-registry project is not currently planning
to spend time supporting this use case. CRDs do not support (and may never
support) versioning, which is very desirable for the cluster registry API. Users
who wish to use a CRD implementation will have to design and maintain it
themselves.

## Tooling design

### User tooling

In the alpha stage, the cluster registry will repurpose the kubefed tool from
Federation and use it to initialize a cluster registry. In early stages, this
will only create a deployment with one replica, running the API server and etcd
in a `Pod`. As the HA requirements for the cluster registry are fleshed out,
this tool may need to be updated or replaced to support deploying a cluster
registry in multiple clusters and with multiple etcd replicas.

Since the cluster registry is a Kubernetes API server that serves a custom
resource type, it will be usable by `kubectl`. We expect that the kubectl
experience for custom APIs will soon be on par with that of core Kubernetes
APIs, since there has been a significant investment in making `kubectl` provide
very detailed output from its describe subcommand; and `kubectl` 1.9 is expected
to introduce API server-originated columns. Therefore, we will not initially
implement any special tooling for interacting with the registry, and will tell
users to use `kubectl` or generated client libraries.

## Repository process

The cluster registry is a top-level Kubernetes repository, and thus it requires
some process to ensure stability for dependent projects and accountability.
Since the Federation project wants to use the cluster registry instead of its
current cluster API, there is a requirement to establish process even though the
project is young and does not yet have a lot of contributors.

The standard Kubernetes Prow bots have been enabled on the repo. There is
currently some functionality around managing PR approval and reviewer assignment
and such that does not yet live in Prow, but given the limited number of
contributors at this point it seems reasonable to wait for sig-testing to
implement this functionality in Prow rather than enabling the deprecated tools.
In most cases where a process is necessary, we will defer to the spirit of the
Kubernetes process (if not the letter) though we will modify it as necessary for
the scope and scale of the cluster registry project. There is not yet a merge
queue, and until there is a clear need for one we do not intend to add one.

The code in the repository will use bazel as its build system. This is in-line
with what the Kubernetes project is attempting to move towards, and since the
cluster registry has a similar but more limited set of needs than Kubernetes, we
expect that bazel will support our needs adequately. The structure that bazel
uses is meant to be compatible with go tooling, so if necessary, we can migrate
away from bazel in the future without having to entirely revamp the repository
structure.

There is not currently a good process in Kubernetes for keeping vendored
dependencies up-to-date. The strategy we expect to take with the cluster
registry is to update only when necessary, and to use the same dependency
versions that Kubernetes uses unless there is some particular incompatibility.

## Release strategy

For early versions, the cluster registry release will consist of a container
image and a client tool. The container image will contain the cluster registry
API server, and the tool will be used to bootstrap this image plus an etcd image
into a running cluster. The container will be published to GCR, and the client
tool releases will be stored in a GCS bucket, following a pattern used by other
k/ projects.

For now, the release process will be managed mostly manually by repository
maintainers. We will create a release branch that will be used for releases, and
use GitHub releases along with some additional per-release documentation
(`CHANGELOG`, etc). `CHANGELOG`s and release notes will be collected manually
until the volume of work becomes too great. We do not intend to create multiple
release branches until the project is more stable. Releases will undergo a
more-detailed set of tests that will ensure compatibility with recent released
versions of `kubectl` (for the registry) and Kubernetes (for the cluster
registry as an aggregated API server). Having a well-defined release process
will be a stable release requirement, and by that point we expect to have gained
some practical experience that will make it easier to codify the requirements
around doing releases. The cluster registry will use semantic versioning, but
its versions will not map to Kubernetes versions. Cluster registry releases will
not follow the Kubernetes release cycle, though Kubernetes releases may trigger
cluster registry releases if there are compatibility issues that need to be
fixed.

Projects that want to vendor the cluster registry will be able to do so. We
expect that these projects will vendor from the release branch if they want a
stable version, or from a desired SHA if they are comfortable using a version
that has not necessarily been fully vetted.

As the project matures, we expect the tool to evolve (or be replaced) in order
to support deployment against an existing etcd instance (potentially provided by
an etcd operator), and to provide a HA story for hosting a cluster registry.
This is considered future work and will not be addressed directly in this
document.

Cross-compilation support in bazel is still a work in progress, so the cluster
registry will not be able to easily provide binary releases for every platform
until this is supported by bazel. If it becomes necessary to provide
cross-platform binaries before bazel cross-compilation is available, the
repository is setup to support common go tooling, so we should be able to devise
a process for doing so.

### Version skew

There are several participants in the cluster registry ecosystem whose versions
will be conceptually able to float relative to each other:

-   The cluster registry API server
-   The host cluster's API server
-   `kubectl`

We will need to define the version skew restraints between these components and
ensure that our testing validates key skews that we care about.

## Test strategy

The cluster registry is a simple component, and so should be able to be tested
extensively without too much difficulty. The API machinery code is already
tested by its owning teams, and since the cluster registry is a straightforward
API server, it should not require much testing of its own. The bulk of the tests
written will be integration tests, to ensure that it runs correctly in
aggregated and independent modes in a Kubernetes cluster; to verify that various
versions of kubectl can interact with it; and to verify that it can be upgraded
safely. A full testing strategy is a requirement for a GA launch; we expect
development of a test suite to be an ongoing effort in the early stages of
development.

The command-line tool will require testing. It should be E2E tested against
recent versions of Kubernetes, to ensure that a simple cluster registry can be
created in a Kubernetes cluster. The multiplicity of configuration options it
provides cannot conveniently be tested in E2E tests, and so will be validated in
unit tests.

## Milestones and timelines

### Alpha (targeting late Q4 '17)

-   Test suite is running on each PR, with reasonable unit test coverage and
    minimal integration/E2E testing
-   Repository processes (OWNERship, who can merge, review lag standards,
    project planning, issue triaging, etc.) established
-   Contributor documentation written
-   User documentation drafted
-   Cluster registry API also alpha
-   All Alpha technical requirements met

### Beta (targeting mid Q1 '18)

-   Full suite of integration/E2E tests running on each PR
-   API is beta
-   Preparatory tasks for GA
-   All Beta technical requirements met
-   User documentation complete and proofed, content-wise
-   Enough feedback solicited from users, or inferred from download
    statistics/repository issues

### Stable (targeting mid Q2 '18)

-   Fully fleshed-out user documentation
-   User documentation is published in a finalized location
-   First version of API is GA
-   Documented upgrade test procedure, with appropriate test tooling implemented
-   Plan for and documentation about Kubernetes version support (e.g., which
    versions of Kubernetes a cluster registry can be a delegated API server
    with)
-   Releases for all platforms
-   Well-defined and documented release process

### Later

-   Documented approach to doing a HA deployment
-   Work on Later technical requirements

## Questions

-   How does RBAC work in this case?
-   Generated code. How does it work? How do we make it clean? Do we check it
    in?
-   Do we need/want an example client that uses the cluster registry for some
    basic operations? A demo, as it were?
-   Labels vs. fields. Versioning? Do we graduate? Do we define a policy for
    label names and have the server validate it?
-   Where should the user documentation for the cluster registry live?
    kubernetes.io doesn't quite seem appropriate, but perhaps that's the right
    place?
-   Is there a reasonable way to support a CRD-based implementation? Should this
    project support it directly, or work to not prevent it from working, or
    ignore it?

## History

| Date  | Details  |
|--|--|
| 10/9/17  | Initial draft |
| 10/16/17 | Minor edits based on comments |
| 10/19/17 | Added section specifically about version skew; change P0/P1/P2 to alpha/beta/stable; added some milestone requirements |
| 11/2/17  | Resolved most comments and added minor tweaks to text in order to do so |
