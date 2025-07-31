# Namespace Sameness - SIG Multicluster Position Statement

Last Edit: 2023/04/22
Status: RELEASED

## Goal

To establish a normative statement for multi-cluster namespace semantics and
governance as a building block for further development which will require
specifying behaviors across clusters.

## Context

Users are reaching for multi-cluster deployments for a
[variety of reasons](http://bit.ly/k8s-multicluster-conversation-starter-doc).
However, Kubernetes has historically treated the cluster boundary as the edge
of the universe.  Without common patterns we can’t build portable tooling to
facilitate multi-cluster capabilities and know that behavior will be consistent
for each user.

## Scope

A single organization may need multiple, disjoint sets of clusters. They may,
for example, represent different phases of the development lifecycle (dev,
staging, prod) or support unrelated projects. Each organization governs their
own clusters in isolation, so the scope of a namespace can only reasonably be
declared within the organizational boundary. The scope of namespace identity is
defined as the set of all clusters, governed by a single authority, that are
expected to work together.  An authority is a company, organization, team,
individual, or other entity which is entrusted to manage these clusters and, in
particular, to create namespaces in them.

## Position

**Within a set of related clusters which are governed by a single authority (a
"clusterset"), samely named namespaces in all clusters are considered to be
related, to have the same owner(s), and to represent the same intention,
unless otherwise dictated by some overriding policy, including, but not limited
to, implementation-specific configuration.**

## Appendix

This position statement is intentionally very abstract, with the goal of not
being overly prescriptive.  The following examples are HYPOTHETICAL, but
hopefully make the statement clearer.

### Example 1: Namespace creation

Consider an organization that runs two clusters, "A" and "B".  They have a
cluster-ops team which is responsible for keeping those clusters alive and
healthy.  They also have app-teams, "foo" and "bar" which use those clusters.

Cluster-ops owns the creation of namespaces in all clusters. They can choose
to delegate the ability to create namespaces to foo-team and bar-team, but any
namespace name that is allocated is "reserved" in all clusters.  How the
delegation works is an area for innovation, but some possible examples
include:
  * A self-service portal which allocates a name in a global database and
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

### Example 3: Multi-cluster Services

Consider the same organization from previous examples.  They have enabled an
implementation of multi-cluster Services, which lets them access backends which
are running on other clusters in the group.

As in example 2, cluster-ops run a per-cluster metrics service,  It does not
make sense for clients in cluster A to access the metrics server of cluster B.
Even though this runs in the same "metrics" namespace in both clusters (and
thus namespace sameness applies), the implementation of multi-cluster services
probably should not be on-by-default.  The details of how multi-cluster
services work are out of scope for this document, but valid approaches could
include:
  * Opt-in: services must be explicitly "exported" to be merged across clusters
  * Opt-out: services or namespaces can be opted out of service merging
  * Different discovery: merged services and "raw" services use different names
    or other discovery mechanisms

However the implementation works, the metrics service is not _necessarily_
merged across clusters, even if the LDAP-to-RBAC sync from example 2 still
applies consistent RBAC policies.  Those are independent capabilities.

### Example 4: Authority sameness policy

Consider the same organization from previous examples, but with more clusters:
"A", "B", "C", and "D".  Business rules dictate that the foo-team and the
bar-team have a common security posture and may share clusters, but the
billing-team has stricter requirements and must not share clusters with other
teams. Cluster-ops wants to assign foo-team's and bar-team's namespaces to
clusters A and B, and billing-team's namespaces to clusters C and D.  They also
want to ensure that the opposite is never true - foo-team's and bar-team's
namespaces cannot be used from clusters C or D, and billing-team's namespaces
cannot be used from clusters A or B.

What exactly it means to "be used" depends on each specific multi-cluster
capability.  For multi-cluster Services, the billing-team's endpoints from
clusters C and D would be eligible for merging, but billing endpoints from
clusters A and B would not.  This can be used by implementations to add another
layer of safety - even if cluster A were compromised and the billing-team's
namespaces was created, the MCS implementation can use the policy described
above to know that it should NOT merge billing endpoints from cluster A.

As an implementation choice, the authority may offer controls which govern
which cluster-namespaces are considered for sameness and which are not.
The details of how these policies might be expressed could include:
  * In-cluster namespace labels or annotations
  * In-cluster custom-resources
  * An external control-plane API
