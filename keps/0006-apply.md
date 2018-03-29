---
kep-number: 6
title: Apply
authors:
  - "@lavalamp"
owning-sig: sig-api-machinery
participating-sigs:
  - sig-api-machinery
  - sig-cli
reviewers:
  - "@pwittrock"
  - "@erictune"
approvers:
  - "@bgrant0607"
editor: TBD
creation-date: 2018-03-28
last-updated: 2018-03-28
status: provisional
see-also:
  - n/a
replaces:
  - n/a
superseded-by:
  - n/a
---

# Apply

## Table of Contents

- [Apply](#apply)
   - [Table of Contents](#table-of-contents)
   - [Summary](#summary)
   - [Motivation](#motivation)
      - [Goals](#goals)
      - [Non-Goals](#non-goals)
   - [Proposal](#proposal)
      - [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
      - [Risks and Mitigations](#risks-and-mitigations)
   - [Graduation Criteria](#graduation-criteria)
   - [Implementation History](#implementation-history)
   - [Drawbacks](#drawbacks)
   - [Alternatives](#alternatives)

## Summary

`kubectl apply` is a core part of the Kubernetes config workflow, but it is
buggy and hard to fix. This functionality will be regularized and moved to the
control plane.

## Motivation

Example problems today:

* User does POST, then changes something and applies: surprise!
* User does an apply, then `kubectl edit`, then applies again: surprise!
* User does GET, edits locally, then apply: surprise!
* User tweaks some annotations, then applies: surprise!
* Alice applies something, then Bob applies something: surprise!

Why can't a smaller change fix the problems? Why hasn't it already been fixed?

* Too many components need to change to deliver a fix
* Organic evolution and lack of systematic approach
  * It is hard to make fixes that cohere instead of interfere without a clear model of the feature
* Lack of API support meant client-side implementation
  * The client sends a PATCH to the server, which necessitated strategic merge patch--as no patch format conveniently captures the data type that is actually needed.
  * Tactical errors: SMP was not easy to version, fixing anything required client and server changes and a 2 release deprecation period.
* The implications of our schema were not understood, leading to bugs.
  * e.g., non-positional lists, sets, undiscriminated unions, implicit context
  * Complex and confusing defaulting behavior (e.g., Always pull policy from :latest)
  * Non-declarative-friendly API behavior (e.g., selector updates)

### Goals

"Apply" is intended to allow users and systems to cooperatively determine the
desired state of an object. The resulting system should:

* Be robust to changes made by other users, systems, defaulters (including mutating admission control webhooks), and object schema evolution.
* Be agnostic about prior steps in a CI/CD system (and not require such a system).
* Have low cognitive burden:
  * For integrators: a single API concept supports all object types; integrators
    have to learn one thing total, not one thing per operation per api object.
    Client side logic should be kept to a minimum; CURL should be sufficient to
    use the apply feature.
  * For users: looking at a config change, it should be intuitive what the
    system will do. The “magic” is easy to understand and invoke.
  * Error messages should--to the extent possible--tell users why they had a
    conflict, not just what the conflict was.
  * Error messages should be delivered at the earliest possible point of
    intervention.

Goal: The control plane delivers a comprehensive solution.

Goal: Apply can be called by non-go languages and non-kubectl clients. (e.g.,
via CURL.)

### Non-Goals

* Multi-object apply will not be changed: it remains client side for now
* Providing an API for just performing merges (without affecting state in the
  cluster) is left as future work.
* Some sources of user confusion will not be addressed:
  * Changing the name field makes a new object rather than renaming an existing object
  * Changing fields that can’t really be changed (e.g., Service type).

## Proposal

Some highlights of things we intend to change:

* Apply will be moved to the control plane: [overall design](goo.gl/UbCRuf).
  * It will be invoked by sending a certain Content-Type with the verb PATCH.
* The last-applied annotation will be promoted to a first-class citizen under
  metadata. Multiple appliers will be allowed.
* Apply will have user-targeted and controller-targeted variants.
* The Go IDL will be fixed: [design](goo.gl/EBGu2V). OpenAPI data models will be fixed. Result: 2-way and
  3-way merges can be implemented correctly.
* 2-way and 3-way merges will be implemented correctly: [design](goo.gl/nRZVWL).
* Dry-run will be implemented on control plane verbs (POST and PUT).
  * Admission webhooks will have their API appended accordingly.
* The defaulting and conversion stack will be solidified to allow converting
  partially specified objects.
* An upgrade path will be implemented so that version skew between kubectl and
  the control plane will not have disastrous results.
* Strategic Merge Patch and the existing merge key annotations will be
  deprecated. Development on these will stop, but they will not be removed until
  the v1 API goes away (i.e., likely 3+ years).

The linked documents should be read for a more complete picture.

### Implementation Details/Notes/Constraints [optional]

What are the caveats to the implementation?
What are some important details that didn't come across above.
Go in to as much detail as necessary here.
This might be a good place to talk about core concepts and how they releate.

### Risks and Mitigations

There are many things that will need to change. We are considering using a
feature branch.

## Graduation Criteria

This can be promoted to beta when it is a drop-in replacement for the existing
kubectl apply, which has no regressions (which aren't bug fixes).

This will be promoted to GA once it's gone a sufficient amount of time as beta
with no changes.

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks

Why should this KEP _not_ be implemented: many bugs in kubectl apply will go
away. Users might be depending on the bugs.

## Alternatives

It's our belief that all routes to fixing the user pain involve
centralizing this functionality in the control plane.
