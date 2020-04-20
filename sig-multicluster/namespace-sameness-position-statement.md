# Namespace Sameness - SIG Multicluster Position Statement

Author: Jeremy Olmsted-Thompson (**[@jeremyot](https://github.com/jeremyot)**), Google  
Last Edit: 2020/04/20  
Status: RELEASED

## Goal
To establish a normative statement for multi-cluster namespace semantics and
governance as a building block for further development which will require
specifying behaviors across clusters.

## Context
Users are reaching for multi-cluster deployments for a
[variety of reasons](http://bit.ly/k8s-multicluster-conversation-starter-doc).
However, Kubernetes treats the cluster boundary as the edge of the universe.
There are currently no standard practices for how to extend the Kubernetes
resource model across multiple clusters. Without common patterns we can’t build
portable tooling to facilitate multi-cluster capabilities and know that behavior
will be consistent for each user.

## Scope
A single organization may need multiple, disjoint sets of clusters. They may,
for example, represent different phases of the development lifecycle (dev,
staging, prod) or support unrelated projects. Each organization governs their
own clusters in isolation, so the scope of a namespace can only reasonably be
declared within the organizational boundary. The scope of namespace identity is
defined as the union of clusters, governed by a single authority, that are
expected to work together.  An authority is a company, organization, team,
individual, or other entity which is entrusted to manage these clusters and, in
particular, to create namespaces in them.

## Position
**For a set of related clusters governed by a single authority, all namespaces of
a given name are considered to be the same namespace. A single namespace should
have a consistent owner across the set of clusters.**

## Appendix

This position statement is intentionally very abstract, with the goal of not
being overly prescriptive.  The following examples are HYPOTHETICAL, but
hopefully make the statement clearer.

### Example 1: Namespace creation

Consider an organization that runs two clusters, "A" and "B".  They have a
cluster-ops team which is responsible for keeping those clusters alive and
healthy.  They also have app-teams, "foo" and "bar" which use those clusters.

Cluster-ops owns the creation of namespaces in both clusters. They can choose
to delegate the ability to create namespaces to foo-team and bar-team, but any
namespace name that is allocated in "reserved" in all clusters.  How the
delegation works is an area for innovation, but some possible examples
include:
  * A self-service portal which allocate a name in a global database and
    creates the namespace on the user's behalf
  * Tooling that allocates a name in a global database and issues a credential
    to the user which is checked in an admission controller
  * An admission controller which requires namespaces be prefixed by the team
    name

However they choose to delegate (or not), once foo-team requests a namespace
called "database" in cluster A, no other team may request a namespace called
"database" in cluster B. That name is taken.

### Example 2: RBAC sync

Consider the same organization from example 1.  As with many large-sized
organizations, they already have a central LDAP server which stores policies
about who is supposed to be able to access what systems.  They enforce this by
converting those policies into Kubernetes RBAC rules and pushing them down into
their clusters.

Cluster-ops runs a metrics service in each cluster.  The RBAC for this service
should be the same in each cluster (i.e. the same set of people administer it,
regardless of which cluster).  The LDAP-to-RBAC sync process can assume that
the "metrics" namespace in each cluster should get the same RBAC rules.

If there are special RBAC rules that need to be applied in some clusters (e.g.
clusters in EU have more limited access), the LDAP-to-RBAC sync implementation
can apply specializations based on whatever criteria it understands.

### Example 3: Multi-cluster services

Consider the same organization from previous examples.  They have enabled an
implementation of multi-cluster Services, which lets them access backends which
are running on other clusters in the group.

As in example 2, cluster-ops run a per-cluster metrics service,  It does not
make sense for clients in cluster A to access the metrics server of cluster B.
Even though this runs in the same "metrics" namespace in both clusters (and
thus namespce sameness applies), the implementation of multi-cluster services
probably should not be on-by-default.  The details of how multi-cluster
services will work are an area for innovation, but ideas include:
  * Opt-in: services must be "exported" to be merged across clusters
  * Opt-out: services or namespace can be opted out of service merging
  * Different discovery: merged services and "raw" services use different names
    or other discovery mechanisms

However the implementation works, the metrics service is not automatically
merged across clusters, though the LDAP-to-RBAC sync from example 2 can still
apply consistent policies.
