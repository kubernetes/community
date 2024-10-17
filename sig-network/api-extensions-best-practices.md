# API Extensions - Best Practices

It is sometimes desireable to build new networking features as [CRDs] rather
than [in-tree]. Reasons may include (but aren't limited to):

* A desire to be free of lockstep with the [KEP] process
* Experimentation and highly iterative development
* The feature is useful to multiple parties, but is niche
* the API doesn't or can't have a default implementation in core

Whatever the case may be, we have multiple SIG Network sub-projects that have
taken this approach for various reasons:

* [Gateway API]
* [Network Policy]
* [Multi Network]

This document intends to share guidelines and best practices for how to create
and manage these kinds of "official, but add-on and optional" features, which
throughout this document we'll refer to as **Official CRDs**.

> **Note**: **This is not a standard**, but **SIG Network projects should try
> and follow this guidance** unless they have some good reasons not to, as the
> guidance here comes from the experience of those who have walked this road
> before. This guidance might be applicable to some other SIGs, but as of the
> time of writing none except SIG Network have reviewed this.

> **Note**: This is **not a comprehensive guide**.

[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
[in-tree]:https://github.com/kubernetes/api
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
* CRD Lifecycle Management
* Conformance Tests & Reports
* Documentation
* Community & Evangelism

In the following section we'll talk about each of these in-depth, including
some lessons learned by existing or historical sub-projects.

[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/

### Enhancement Proposals

Projects which produce optional [Custom Resource Defintion (CRD)][CRD]-based
APIs should provide a project-specific enhancement proposal system. You can see
examples of using project-specific enhancement proposals with some of the
current sub-projects:

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

We strongly recommend this pattern (unless there's some very compelling reason
reason not to use it for any particular project).

> **Note**: When starting new projects, we advise having enhancement proposals
> preceed any code whatsoever, for posterity and to ensure time for all
> stakeholders in the community are represented.

> **Note**: Taking this approach doesn't mean complete autonomy. Any GA APIs
> (version is v1 or greater) with a `*.k8s.io` group **MUST** go through the API
> review process with the SIG Network leads for any content considered
> stable/standard/GA. The use of any non-`*.k8s.io` group for APIs published by
> official SIG Network sub-projects or working groups needs to be ratified with
> SIG Network leads first. Experimental features, alpha, beta, e.t.c. can be
> developed without that additional oversight until they are promoted to GA.

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

### CRD Lifecycle Management

TODO - installing, and managing on real systems

### Conformance Tests & Reports

TODO - framework for testing and reporting

### Documentation

TODO

## Important Notes

> **Note**: At the time of writing, we have no extension APIs that have
> transferred themselves in-tree after doing development iterations
> out-of-tree. This is uncharted territory at present, though we are not
> opposed to it happening if it seems appropriate.
