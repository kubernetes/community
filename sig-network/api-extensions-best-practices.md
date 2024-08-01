# API Extensions - Best Practices

> **Warning**: SIG API Machinery is responsible for API development. This
> document provides some guidance based on SIG Network's experience, however
> when in doubt discuss with API Machinery as this document is not an
> authoritative or complete guide on API development as a whole.

> **Warning**: This document is a work in progress: see sections below which
> are noted with `TODO`.

This guide provides help for those who would want to provide new upstream
networking functionality via an extension API rather than [in-tree], as it is
sometimes desirable to build new networking features as [CRDs] rather than
in-tree. Reasons including (but not limited to):

* A need to move at a different pace than the core project
* Experimentation and/or highly iterative development
* The feature is useful to multiple parties, but is niche
* the API doesn't or can't have a default implementation in core

Whatever the case may be, we have multiple SIG Network sub-projects that have
taken this approach for various reasons:

* [Gateway API]
* [Network Policy]
* [Multi Network]

Here you'll find guidelines and best practices for how to create and manage
these kinds of "official, but add-on and optional" features, which throughout
this document we'll refer to as **Official CRDs**.

> **Note**: **This is not a standard**, but **SIG Network projects should try
> and follow this guidance** unless they have some good reasons not to, as the
> guidance here comes from the experience of those who have walked this road
> before. This guidance might be applicable to some other SIGs, but as of the
> time of writing none except SIG Network have reviewed this.

> **Note**: This is **not a comprehensive guide**.

[in-tree]:https://github.com/kubernetes/api
[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
[KEP]:https://github.com/kubernetes/enhancements/blob/master/keps/README.md
[Gateway API]:https://github.com/kubernetes-sigs/gateway-api
[Network Policy]:https://github.com/kubernetes-sigs/network-policy-api
[Multi Network]:https://github.com/kubernetes-sigs/multi-network

## Objectives

At a high level, SIG Network sub-projects that deploy as [CRDs] will need to
deal with (at least) the following dimensions:

* Enhancement Proposals
* Delivering CRDs
* CRD Lifecycle Management
* Conformance Tests & Reports
* Documentation
* Community & Evangelism

In the following section we'll talk about each of these in-depth, including
some lessons learned by existing or historical sub-projects.

[CRDs]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/

### Enhancement Proposals

Projects which produce optional [Custom Resource Definition (CRD)][CRD]-based
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
> precede any code whatsoever, for posterity and to ensure time for all
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

#### Known Challenges

Developing with CRDs can be challenging and we've gained a lot of experience
with some of the trickier bits over the years. In this section we'll describe in
detail some subtle or particularly nebulous challenges we've faced, and provide
guidance on how to navigate them.

##### When a feature can be supported by many but not ALL implementations

When developing an API in a community (especially if there are many
implementations) you may find there are features that are not common across all
implementations, but a subset of them can implement them.

**In an ideal situation no user is provided fields in an API that they can't
actually use**. It can be tempting to take some tradeoffs on this to try to find
a middle-ground so that you can provide more overall features. The [Gateway API]
project (in particular) ran into this problem early on and decided to try an
approach called [Conformance Support Levels].

With the "support levels approach", all the ubiquitous features are covered by
a "core" level of support. Features which are supported by at least 3
implementations are covered by an "extended" level of support. Each of these has
conformance tests and are included in reports, but ultimately when the user is
defining their [HTTPRoute] (or other resource) **it can be very unclear whether
the fields marked as "extended" will actually be functional given any particular
implementation, and between different implementations**. Morever (at the time
of writing) Gateway API doesn't have a complete solution for how feedback to a
user should work in this regard (but there are some [ongoing proposals]).

It is due to the problems stated above and the lack of implementation
complete-ness so far that **we do not advise employing an approach like
conformance levels at this time**. In general, we advise that all fields
available in your APIs be functional regardless of implementation.

> **Note**: This guidance may be updated in the future if we come to more
> definitive conclusions on the value and efficacy of support levels for
> end-users.

[Gateway API]:https://github.com/kubernetes-sigs/gateway-api/
[Conformance Support Levels]:https://gateway-api.sigs.k8s.io/concepts/conformance/#2-support-levels
[HTTPRoute]:https://gateway-api.sigs.k8s.io/api-types/httproute/
[ongoing proposals]:https://gateway-api.sigs.k8s.io/geps/gep-2162/

### Delivering CRDs

TODO

### CRD Lifecycle Management

TODO - installing, and managing CRDs. How this should work on platforms, who's
       responsible (and trying to reduce responsibility on the cluster operator)
       Also need to talk about interop between multiple dependent projects.

### Conformance Tests & Reports

TODO - framework for testing and reporting

## Important Notes

> **Note**: At the time of writing, we have no extension APIs that have
> transferred themselves in-tree after doing development iterations
> out-of-tree. This is uncharted territory at present, though we are not
> opposed to it happening if it seems appropriate.

## TODO

- [ ] Delivering CRDs
- [ ] Conformance Tests and Conformance Reports
- [ ] Documentation
- [ ] Community & Evangelism
