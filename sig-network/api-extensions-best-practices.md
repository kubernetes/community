# API Extensions - Best Practices

It is sometimes desireable to build new networking features as [CRDs] rather
than [in-tree]. Reasons may include (but aren't limited to):

* A desire to be free of lockstep with the [KEP] process
* Experimentation and highly iterative development
* The feature is useful to multiple parties, but is niche

Whatever the case may be, we have multiple SIG Network sub-projects that have
taken this approach for various reasons:

* [Gateway API]
* [Network Policy]
* [Multi Network]

This document intends to share guidelines and best practices for how to create
and manage these kinds of "official, but add-on and optional" features, which
throughout this document we'll refer to as **Official CRDs**.

> **Note**: **This is not a standard**, but **projects should try and follow
> this guidance** unless they have some good reasons not to, as the guidance
> here comes from the experience of those who have walked this road before.

> **Note**: This is **not a comprehensive guide**.

[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
[in-tree]:https://github.com/kubernetes/kubernetes/
[KEP]:https://github.com/kubernetes/enhancements/blob/master/keps/README.md
[Gateway API]:https://github.com/kubernetes-sigs/gateway-api
[Network Policy]:https://github.com/kubernetes-sigs/network-policy-api
[Multi Network]:https://github.com/kubernetes-sigs/multi-network

## Objectives

At a high level, SIG Network sub-projects that deploy as [CRDs] will need to
deal with (at least) the following dimensions:

* Enhancement Proposals
* Developing CRDs
* Delivering CRDs
* Conformance Levels
* Conformance Tests & Reports
* CRD Lifecycle Management
* Documentation
* Community & Evangelism

In the following section we'll talk about each of these in-depth, including
some lessons learned by existing or historical sub-projects.

[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/

### Enhancement Proposals

Traditionally Kubernetes enhancements go through the [KEP] process to progress.
While it may be OK for a sub-project to use the KEP process for an extension
API, it may also be reasonable for that project to have their own exclusive
enhancement proposal process specific to the project.

You can see examples of this with some of the current sub-projects:

* [Gateway Enhancement Proposals (GEPs)][geps]
* [Network Policy Enhancement Proposals (NPEPs)][npeps]

Key characteristics of this pattern are:

* an easy to find directory at the root of repositories, such as `geps/`, with
  references provided from the `README.md`.
* a readme for the enhancement process (we recommend using GEP as a template,
  and we recommend a motivation focused proposal process (see more below)).
* a project board which organizes these GEPs and helps provide visibility into
  the project's priorities and roadmap
* trackable issues and PRs for every proposal that go through phases (e.g.
  (`provisional`, `implementable`, `experimental`, e.t.c.)

We generally recommend this pattern unless there's a strong reason not to.

> **Note**: When starting new projects, we advise having enhancement proposals
> preceed any code whatsoever, for posterity and to help ensure inclusion.

> **Note**: Taking this approach doesn't mean complete autonomy. Any APIs
> (extension, or otherwise) **MUST** go through the API review process with the
> SIG Network leads for any content considered stable/standard/GA. Experimental
> features, alpha, beta, e.t.c. can be developed without that additional
> oversight until they are promoted.

[KEP]:https://github.com/kubernetes/enhancements/blob/master/keps/README.md
[geps]:https://github.com/kubernetes-sigs/gateway-api/blob/main/geps/overview.md
[npeps]:https://github.com/kubernetes-sigs/network-policy-api/tree/main/npeps

#### Motivation Focused Enhancement Proposals

It can be very easy to bring a new proposal to the table, and start with a huge
GEP and maybe a ton of upfront code. In our experience however having too much
up front, and in particular coming right to the table with an implementation
can cause a proposal to get gridlocked very quickly.

**We strongly recommend a "What, why and who before how" approach**. The first
PR which adds an enhancement proposal should establish:

* **What?** - What do we want to do, at a high level? A summary and/or specific
  high-level goals to achieve.
* **Why?** - The motivation, the user stories, the reason we want to do this.
* **Who?** - Every proposal stands a better chance of succeeding if there are
  multiple aligned parties interested in it, and collaborating on it upfront.

Without any indication of _how_ we accomplish it. This, plus a process of steps
for graduation (e.g. `provisional` -> `implementable` -> `experimental` ->
`standard`) allows the community to work more iteratively, and work on alignment
with the principles before spending any time in the implementation details.

### Developing CRDs

TODO

### Delivering CRDs

TODO

### Conformance Levels

TODO - from Gateway API conformance levels

### Conformance Tests & Reports

TODO - framework for testing and reporting

### CRD Lifecycle Management

TODO - installing, and managing on real systems

### Documentation

TODO

### Community & Evangelism

TODO - blogging, conference talks, e.t.c.

## Important Notes

> **Note**: At the time of writing, we have no extension APIs that have
> transferred themselves in-tree after doing development iterations
> out-of-tree. This is uncharted territory at present, though we are not
> opposed to it happening if it seems appropriate.
