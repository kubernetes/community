# SIG Platform Dev Charter

## Mission Statement

SIG Platform Development (SIG-PD) is a [cross-cutting SIG][sigs] whose
mission is to design, build, and maintain a cohesive, first-party
Platform Development Kit (PDK) for developers who want to extend the
Kubernetes platform by adding new APIs or enhancing existing APIs,
as well as to advocate on behalf of such developers by driving
feedback into other SIGs.

The tools in this kit must work well together and must fit into a single,
cohesive vision (to be defined as an activity of the SIG) for how extension API
developers can on-board and grow on the Kubernetes platform.

Although the kit should be opinionated to encourage best practices based on our
experience building core APIs, the kit and its tools must remain optional.
That is, SIG-PD must never own end-user-visible features that are required for
a cluster to pass [conformance], although it may collaborate with other SIGs
that own such features.

[sigs]: https://github.com/kubernetes/community/blob/master/governance.md#sigs
[conformance]: https://github.com/cncf/k8s-conformance

### Motivation

Today, there exists no organization around the ecosystem for extending Kubernetes.
We can guide developers towards best practices by providing tools and guides that
help them work with the underlying Kubernetes primitives in well-understood ways.
By collaborating upstream on these projects, we can begin to understand the
shortcomings of Kubernetes extension primitives and more effectively reason
about future requirements.

Examples of best practices we’ve seen ignored in Kubernetes extensions include but are not limited to:

* Least-privilege service accounts
  * Many examples grant permission to install arbitrary CRDs.
* Level-triggered reconciliation
  * Many examples encourage brittle, purely event-triggered logic
    (hooking behavior directly to add/update/delete watch events).
* API Conventions
  * Many examples give no guidance on the use of Kubernetes API conventions like
    spec/status, owner/controller references, orphan/adopt semantics, deletion
    propagation policies, backward/forward compatibility, etc.
* Scalability
  * Many examples encourage separate watch streams per controller,
    which won't scale as more people start writing/installing custom controllers.
  * Many examples give no guidance on error handling, retries,
    exponential backoff, optimistic concurrency, etc.
* Idiomatic use of API groups and versions
  * Many examples try to independently version multiple Kinds within a single Group.
  * Many examples conflate [format versioning][] (a Kind with multiple Versions)
    and entity versioning (must be separate Kinds in Kubernetes).

Addressing these problems requires expertise from across many SIGs
(api-machinery, apps, auth, cli, docs, release, scalability, testing)
because the task of creating and distributing API extensions is effectively
a microcosm of developing and releasing Kubernetes itself.

Many vendors have been working on this problem
(CoreOS, Giant Swarm, Rook, Google, Red Hat),
but they have all at some point proposed doing so as a community effort
within the Kubernetes project.
This indicates broad agreement that the foundation for building
extensions to the Kubernetes API should not be left up to competition
within the ecosystem.

By analogy, we intend to build something like the first-party
[Windows Driver Kit] (*not* the [Windows SDK]),
which helps developers make use of interfaces provided by the
Windows Kernel to extend its capabilities.
Our Kubernetes PDK will help developers make use of interfaces
provided by the Kubernetes API Server to extend its capabilities.

[format versioning]: https://cloudplatform.googleblog.com/2018/03/API-design-which-version-of-versioning-is-right-for-you.html
[Windows Driver Kit]: https://en.wikipedia.org/wiki/Windows_Driver_Kit
[Windows SDK]: https://en.wikipedia.org/wiki/Microsoft_Windows_SDK

## Scope

### Tools, Examples, Guides, and Experiments

The primary scope of SIG-PD is to design, build, and maintain a cohesive,
first-party toolkit for developers who want to extend the Kubernetes platform
by adding new APIs or enhancing existing APIs.

Components of this toolkit should include but are not limited to:

* Tools and libraries to help people use API extension mechanisms, including authoring types and coping with versions and upgrades.
* Tools and libraries to help people write custom controllers.
* Tools and libraries to help people test, document, package, and release platform extensions (APIs/controllers).
* Tools and libraries to help people create a CLI experience for platform extensions.
* Tools to help people consume platform extensions.
* Documentation and examples for using API extension mechanisms. (API object serving)
* Documentation and examples for writing custom controllers. (API behavior)
* Documentation and examples of putting together various parts of the toolkit
  to achieve high-level goals.

Such components will be managed by subprojects within the SIG.

In addition to components of the actual toolkit, the SIG may also use
subprojects to collaborate on experimental tools and patterns that may
one day be added to the toolkit or otherwise inform its development.

### Out of Scope

Examples of things that are *out of scope* for SIG-PD include:

* Any features of Kubernetes that are both visible to end users and required for conformance,
  even if they relate to providing extension points. (e.g. CRD)
* Shared libraries that are meant for purposes beyond platform extension,
  even if platform extension is among their potential uses. (e.g. client-go)
* Tools that are meant for purposes beyond platform extension,
  even if platform extension is among their potential uses. (e.g. FaaS frameworks)
* Tools for developing, deploying, and/or managing "applications" (not platform extensions)
  on Kubernetes.
* Tools for interfacing with extension points contained within specific APIs
  (e.g. CSI, an extension point within the storage volume API), as opposed to
  extension points for creating APIs themselves (e.g. CRD/Aggregation).
* Tools that incidentally *use* platform extension, but are not targeted primarily at
  helping people create their own platform extensions.
* Tools for platform extension that neither fit into the current design philosophy
  of the toolkit, nor represent a potential future direction for the toolkit according
  to the judgement of the SIG.

Unlike [SIG API Machinery][] (SIG-AM), SIG-PD does not and must never own code that
implements an end-user-visible feature that is required to pass [conformance].
Other SIGs that maintain core APIs that are required for conformance may
one day choose to use our tools, but that fact must be invisible to end users
of the API, and conformance tests must not require the API to be implemented
on top of our tools.

For example, although we develop first-party tools, docs, and best practices for
*using* features like [CRD], we do not own the code that *implements* that feature
(in [apiextensions-apiserver]) because CRD is a required feature for a
conforming Kubernetes cluster.
Whenever changes are needed in such features to improve the user experience,
SIG-PD will give feedback to the appropriate SIG and work with them to define
the requirements by advocating on behalf of platform extension developers.

Similarly, although the examples and tools we develop will make heavy use of [client-go],
we do not own that library because it's a low-level facility for talking to the
API server *for any purpose*. For example, client-go must be a common denominator
for use by developer tools like kubectl or Helm (deploying apps by posting manifests),
and as such it is outside the scope of SIG-PD.

To put it another way, *from the perspective of extension developers*:

*The problem with client-go is not that it's bad at what it does (it's actually quite good);
the problem is that client-go is the wrong tool for the job.
Building Operators with client-go is like writing shell scripts in assembly.*

This is why people focused on extension developers have built higher-level abstractions
on top of client-go (operator-kit, kubebuilder, metacontroller).
SIG-PD will focus on building the right tools for the job of extending the
Kubernetes platform.

To the extent that improvements to client-go are needed specifically to better
serve developers of platform extensions, SIG-PD will give feedback to and
collaborate with SIG API Machinery, who will continue to own the code.

[SIG API Machinery]: https://github.com/kubernetes/community/tree/master/sig-api-machinery
[CRD]: https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/
[apiextensions-apiserver]: https://github.com/kubernetes/apiextensions-apiserver
[client-go]: https://github.com/kubernetes/client-go/

### Inaugural Subprojects

These are subprojects that we propose to start at the inception of the SIG.
Upon approval of the SIG charter, we will begin recruiting members and call for
proposals to staff and flesh out these subprojects through the [KEP] process.

The seed TLs will approve these initial subproject KEPs to bootstrap the set of
subproject owners, after which subproject owners will form the primary
decision-making body.

Note that some of these are existing subprojects of other SIGs,
and will only be transferred with consent of the owning SIG.

* PDK Roadmap
  * Define and maintain the future vision for a cohesive, first-party Platform Development Kit (PDK).
* API Samples
  * Create and maintain high-quality sample code for using Kubernetes API extension points.
  * Includes sample-apiserver and sample-controller from [server-sdk][am-subprojects] in SIG-AM.
* API Generators
  * Create and maintain code generators for building Kubernetes APIs.
  * Includes gengo and code-generator from [idl-schema-client-pipeline][am-subprojects] in SIG-AM.
* [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder)
  * An existing subproject of SIG-AM for creating standalone or aggregated API servers.

[am-subprojects]: https://github.com/kubernetes/community/tree/master/sig-api-machinery#subprojects

### Experimental Subprojects

These are subprojects that we believe represent experiments that are within the scope
of the SIG, and that the SIG should facilitate and collaborate on as a way of gathering
insight on what should go into the PDK.
Upon approval of the SIG charter, after the inaugural subprojects are staffed with voting
subproject owners, we will propose these subprojects through the [KEP] process.

Note that some of these are existing subprojects of other SIGs,
and will only be transferred with consent of the owning SIG.
Others are currently external to Kubernetes, and due diligence will be exercised to
ensure a proper transfer of project management if the KEP is approved by the SIG.

Also note that even if the SIG accepts these subprojects,
that does not make them part of the PDK.

* [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)
  * A framework (generators and libraries) for building extension APIs and controllers.
* [metacontroller](https://github.com/GoogleCloudPlatform/metacontroller)
  * A hook-based wrapper to simplify adoption of core controller code, patterns, and best practices.

## Roles

### SIG Chair

The role of SIG Chair is purely organizational in this SIG.
Chairs do not necessarily have influence over SIG-wide technical decisions,
unless they also happen to be subproject owners and/or technical leads.

- Run operations and processes governing the SIG.
- Seed members established at SIG founding.
- Chairs *MAY* decide to step down at anytime and propose a replacement.  Use lazy consensus amongst
  chairs with fallback on majority vote to accept proposal.  This *SHOULD* be supported by a majority of
  SIG Members.
- Chairs *MAY* select additional chairs through a [super-majority] vote amongst chairs.  This
  *SHOULD* be supported by a majority of SIG Members.
- Chairs *MUST* remain active in the role and are automatically removed from the position if they are
  unresponsive for > 3 months and *MAY* be removed if not proactively working with other chairs to fulfill
  responsibilities.
- Number: 2+
  - There is no upper bound on this number because this role is purely
    one of service to the SIG, not influence over it.
- Defined in [sigs.yaml].

### SIG Technical Leads

- Resolve cross-subproject technical issues and decisions.
  - Issues within a subproject should be resolved by that subproject's owners, not the SIG TLs.
  - Issues that require cross-subproject agreement should be resolved first by lazy consensus
    or majority vote among subproject owners, but may be escalated to SIG TLs if agreement can't be reached.
    - Examples of cross-subproject issues:
      - Agreement on a cohesive vision for the PDK.
      - Agreement on interfaces and integrations between subprojects, or on common building blocks.
      - Changes requested in one or multiple subprojects to better align them to each other or consolidate code/effort.
    - If subproject owners or other SIG members are not satisfied with the resolution,
      they should escalate to SIG Architecture or the steering committee as appropriate.
- Seed members established at SIG founding.
- TLs *MAY* decide to step down at anytime and propose a replacement.
  Use lazy consensus amongst subproject owners with fallback on majority vote to accept proposal.
  This *SHOULD* be supported by a majority of SIG Members.
- TLs *MAY* select additional TLs through a [super-majority] vote amongst TLs.
  This *SHOULD* be supported by a majority of SIG Members.
- TLs *MUST* remain active in the role and are automatically removed from the position if they are
  unresponsive for > 3 months and *MAY* be removed if not proactively working with other TLs to fulfill
  responsibilities.
- Number: 3-5
  - This number is bounded to ensure efficient resolution of escalations within the SIG.
  - There must never be more than 2 TLs from the same company.
- Defined in [sigs.yaml].

### Subproject Owners

- Scoped to a subproject defined in [sigs.yaml].
- Resolve issues within their subproject when consensus cannot be reached among subproject members.
- Collectively decide on cross-subproject issues by lazy consensus with fallback on majority vote,
  and escalation to SIG TLs only if necessary.
  - Each person who is an owner of any subproject of the SIG has one vote,
    regardless of whether they are an owner of more than one subproject.
- Seed members established at subproject founding
- *MUST* be an escalation point for technical discussions and decisions in the subproject
- *MUST* set milestone priorities or delegate this responsibility
- *MUST* remain active in the role and are automatically removed from the position if they are unresponsive
  for > 3 months.
- *MAY* be removed if not proactively working with other Subproject Owners to fulfill responsibilities.
- *MAY* decide to step down at anytime and propose a replacement.  Use [lazy-consensus] amongst subproject owners
  with fallback on majority vote to accept proposal.  This *SHOULD* be supported by a majority of subproject
  contributors (those having some role in the subproject).
- *MAY* select additional subproject owners through a [super-majority] vote amongst subproject owners.  This
  *SHOULD* be supported by a majority of subproject contributors (through [lazy-consensus] with fallback on voting).
- Number: 3
  - This number must be fixed to prevent individual subprojects from
    artificially inflating their influence on votes among subproject owners.
  - However, approvers for parts of the subproject can be added without making them subproject owners.
- Only the approvers in the top-level subproject [OWNERS] file listed in [sigs.yaml]
  are considered subproject owners.

### Members

- *MUST* maintain health of at least one subproject or the health of the SIG
- *MUST* show sustained contributions to at least one subproject or to the SIG
- *SHOULD* hold some documented role or responsibility in the SIG and / or at least one subproject
  (e.g. reviewer, approver, etc)
- *MAY* build new functionality for subprojects
- *MAY* participate in decision making for the subprojects they hold roles in
- Includes all reviewers and approvers in [OWNERS] files for subprojects

## Organizational management

- SIG meets bi-weekly on zoom with agenda in meeting notes
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for Kubecon
  - *SHOULD* be organized by chairs unless delegated to specific Members

- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project management

Subprojects are used both for building and maintaining components of the first-party toolkit,
as well as for conducting and facilitating experiments that the SIG believes will inform
future development of the toolkit.

The PDK Roadmap subproject is responsible for defining what components are in the official toolkit.
All subprojects should be considered experimental by default unless they are in the PDK Roadmap.

#### Subproject creation

Subprojects may be proposed through the [KEP] process. To be accepted, the subproject KEP must achieve all of:

* [Lazy consensus][lazy-consensus] approval among subproject owners, with fallback on majority vote.

  *and*

* Majority approval among SIG TLs.
  * All SIG TLs *MUST* explicitly vote Yes or No on every subproject proposal
    as part of "remaining active in the role".
  * Along with their No vote, SIG TLs *SHOULD* provide rationale such as:
    * Not in scope for the SIG.
    * Should be part of another subproject.
    * Should be reconsidered when it's more mature.

The KEP *MUST* establish subproject owners, and [sigs.yaml] *MUST* be updated
to include subproject information and [OWNERS] files with subproject owners.
Where subproject processes differ from the SIG governance, they must document how.

Subprojects must define how releases are performed and milestones are set,
as well as how releases will be qualified through testing.

#### Subproject retirement

Subprojects may be retired (removed from the SIG) by [super-majority] vote among
subproject owners.

Part of the SIG's scope is to conduct and facilitate experimentation (through subprojects)
around tools and patterns that may or may not actually make it into the first-party,
Kubernetes-branded toolkit.

It is therefore expected that some subprojects will be retired, either because
they have stopped making progress, they have merged with other subprojects,
or the SIG no longer believes the subproject represents a direction that might
someday be useful for the first-party toolkit.

Note that when a subproject is retired, its owners lose future SIG voting rights,
unless they are also tech leads or owners of another remaining subproject.

### Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly following alternatives
they have defined.

- Proposing and making decisions
  - Proposals sent as [KEP] PRs and announced to a subproject mailing list.
  - Follow [KEP] decision making process.

- Test health
  - Subprojects must define a test suite that ensures the health of the master branch.
  - PRs that break tests should be rolled back if not fixed within 1 business day.

[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[OWNERS]: https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md

