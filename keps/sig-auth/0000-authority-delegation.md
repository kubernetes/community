---
kep-number: 0
title: Authority Delegation
authors:
  - "@mikedanese"
owning-sig: sig-auth
participating-sigs:
  - sig-auth
  - sig-api-machinery
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: provisional
---

# Authority Delegation

This document proposes a design to support coherent, consistent authorization
delegation in Kubernetes.

## Motivation

Kubernetes has no mechanism to propagate authority through the stack.

* For use in authorization: All controllers act as deputies but we only have
  quality control over the core controllers. 3rd party controllers may turn out
  to be more easily confused and tricked into using their authority incorrectly,
  putting our users at unnecessary risk. Ideally the authority of a controller
  would be limited to (a subset of) the authority of the users actively using
  the controller.
* For use in audit: Lack of provenance of authority greatly restricts what we
  can implement as far as audit logging and access transparency. Tracing the
  authority of a low level action to the source is painful, imprecise,
  dependent on k8s version, even when each actor in the path is cooperative.

The proposed authorization model seeks to solve these issues.

<!---

## Table of Contents

## Summary


### Goals

### Non-Goals

-->


## Proposal

### Selected Authority

Today Kubernetes supports a kind of ambient authority. However support for
selected authority is a generally useful model to support.

This is particularly useful in a deputy scenario to prevent confused deputies.
When a controller acts on behalf of an end principal, the authority by which it
acts should be explicit. This prevents a controller from misusing permission
granted ambiently to the controller.

Authority is selected per request by passing an “Authority-Selector” header
with the request:

```
Authority-Selector: {“authorityReference”:{“kind”:”ClusterRole”,”name”:”viewer”}}
```
or:
```
Authority-Selector: {“authorityReference”:{“kind”:”AuthorityDelegation”,”uid”:”0000-0...”}}
```

The value of the selector is a `SelectedAuthority` object in the
`authorization.k8s.io` API group:

```go
type SelectedAuthority struct {
  metav1.TypeMeta
  AuthorityReference AuthorityReference
}

type AuthorityReference struct {
  metav1.TypeMeta
  Name string
  UID types.UID
}
```

When a request is received by an API, the API would normally iterate over the
set of authorizers and each authorizer would check all policies that would
permit the action, searching for a policy that permits the action and is bound
to the user.

When a request is received by the API with a selected authority, the API would
check the specific authorizer that supports the AuthorityReference kind, and
the authorizer would check the request against the specific policy referenced
in by the SelectedAuthority.

### Persistent Authority Delegation

#### AuthorityDelegation

A new non-namespaced persistent object will be added to the
`authorization.k8s.io` API group to store instances of authority delegation.

```go
type AuthoritiyDelegation struct {
  metav1.TypeMeta
  metav1.ObjectMeta

  // The user that created the delegation.
  User Info
  // The principal that is delegating authority.
  Subject SubjectRef
  // The principals that are delegated authority.
  Delegates []SubjectRef
  // defined above
  Authority AuthorityReference
  // If this delegation is created selecting a previously delegated authority,
  // the delegation chain is recorded here.
  DelegationHistory []AuthorityDelegationReference
}

type SubjectRef struct {
  Kind SubjectKind
  Name string
  UID types.UID
}

const (
  GroupKind SubjectKind = "Group"
  UserKind SubjectKind = "User"
)

type AuthrityDelegationReference struct {
  UID types.UID
}
```

#### AuthorityDelegationPolicy

A new policy will be introduced to implictly create `AuthorityDelegations` in
reaction to API actions. The `AuthorityDelegationPolicy` will be a
non-namespaced object in the `authorization.k8s.io` API group.

```go
type AuthorityDelegationPolicy struct {
  // The resource attributes of an action to implicitly create a delegation for.
  ResourceAttributes *ResourceAttributes
  // The principals that are delegated authority.
  Delegates []SubjectRef
  // The permissions that are delegated.
  Authority AuthorityReference
}
```

When an API action matching an `AuthorityDelegationPolicy` is processed, an
`AuthorityDelegation` is implicitly created before any objects are persisted.

#### Example of a Cascading Delegation

Suppose a situation where:

1) End user Foobar creates a deployment.
2) Deployment controller creates a replicaset in response.
3) ReplicaSet controller creates a pod in response.

In a delegated authority model, the delegation graph looks like:

      Foobar -> Deployment controller -> ReplicaSet controller

Each edge of this graph represents a single authority delegation. Controllers
use the authority of the delegation represented by the incoming edge.

#### Semantics of selecting a DelegatedAuthority

Delegations encodes a permission delegation at a point in time. Revoking
permission of the end user does not revoke permission of historic delegations
made by the end user. If team member Foobar leaves the team, all the production
deployments Foobar created while on the team should continue to function.
However, actions taken using authority that was historically delegated by Foobar
are linked to Foobar in audit via the delegation history. Similarly, persistent
authority delegations made by Foobar while they were on the team are queryable
by querying the list of authority delegations.

Right now, users can create replica sets without necessarily being able to
create pods. In that case, the end user lacks the authority that the replicaset
controller would need to operate.

### Ephemeral Authority Delegation

Authority delegation via other mechanisms should have equivalent representation
in audit logs.

* Via ServiceAccount token vending
* Via Impersonation

### Audit: Access Transparency

Audit logs should encode the full delegation history of an action.

## Extra

### Replacing controller RBAC policy

RBAC policy to authorize contollers should be migrated to
`AuthorizationDelegationPolicy` once all controllers are selecting
`AuthorityDelegations`.

### ABAC: Conditional Authority

It would be useful to restrict `AuthorityReferences` in an `AuthorityDelegation`
to resource attributes, e.g.:

* `PodTemplate` hash must be foo to use a delegated authority.
* Modifications must be made to only the replicas field to use this delegated
  authority.

<!---

### Authorization Tokens

### SubjectAccessReview

### Expiring persisted authority delgations

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

-->
