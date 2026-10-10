# WG LTS Charter
This charter adheres to the [wg-governance] guidance, as well as
adheres to the general conventions described in the [Kubernetes
Charter README] and the Roles and Organization Management outlined
in [sig-governance], where applicable to a Working Group.

## Scope
The Long Term Support Working Group (WG LTS) is organized with the
goal of provide a cross SIG location for focused collection of
stakeholder feedback regarding the support stance of Kubernetes
project in order to inform proposals for support improvements.

"WG LTS" is simply shorter than "WG To LTS Or Not To LTS" or "WG
What Are We Releasing And Why And How Is It Best Integrated, Validate,
And Supported", but should NOT be read in that shortness to imply
establishing a traditional LTS scheme (multi-year support; upgrade
from LTS N to N+1, skipping intermediate versions) is the foregone
conclusion of the WG.

### In Scope
* What is a Kubernetes release?
* What is a supported release?  Ie: Stable branches to which critical fixes are backported.
 * Number of community supported branches.
 * Duration of community support per supported branch.
 * Upgrade path considerations.
* Costs of Kubernetes releases in terms of:
 * Infrastructure
 * People
* Process: how/where code changes are backported to community supported branches.

### Out of Scope
* The lifecycle of projects outside of the Kubernetes org.
* Technical and end-user support:  The WG will make recommendations
  around support to those responsible for relevant code and responsible
  for the release engineering operations and automation, but does not
  own code itself.
* Code, test case, automation implementations:  this is a working
  group, no code implementation is the responsibility of this Working
  Group.

### More Details
For even more details on scope and conflicting tradeoffs and implications of
different support schemes, please see [WG LTS Proposal Presentation],
the [WG LTS meeting minutes], and [WG LTS YouTube Channel].

## Stakeholders
Stakeholders in this workgroup span multiple personas and SIGs.

There are developers of Kubernetes internals.  Arguably all feature
code delivering SIGs are stakeholders here in as much as proposals
for changes in support will impact the way the community develops,
ships and maintains its code.  From the perspective of technical
component interoparability across version skews, developers in:
* SIG API Machinery
* SIG CLI
* SIG Node
will be particularly impacted if wider version skew support to be
attempted.

There is another set of developers of Kubernetes internals who
are involved in the integration of Kubernetes with a hosting
platform or infrastructure providers and delivering Kubernetes
components to run on cluster nodes. This brings in additional dimensions
of version skew and does so with more external components and support
matrix complexity.  These are developers in, for example:
* SIG Storage (CSI)
* SIG Network (CNI)
* SIG Node (CRI)
* SIG Cloud Provider
  * SIG AWS
  * SIG Azure
  * SIG GCP
  * SIG IBMCloud
  * SIG OpenStack
  * SIG VMware
* SIG Cluster Lifecycle

There is yet another set of developers of Kubernetes internals who are
those involved in meta-topics:
* SIG Release: production of supported release artifacts
* SIG Testing: how we can most effectively test Kubernetes
* Product Security Committee (PSC): security vulnerability handling
* SIG Architecture: maintains and evolves the design principles of Kubernetes, and provides a consistent body of expertise necessary to ensure architectural consistency over time.  Also defines conformance testing.
* Steering Committee: scope includes deciding how and when official releases of Kubernetes artifacts are made and what they include

Then of course there is also a variety of types of users of Kubernetes:
* Kubernetes end users
* Kubernetes cluster operators
* Kubernetes vendors, distributions, and hosting providers
These have their own operational complexities and desires around support
duration and supported upgrade version skews.

## Deliverables
The specific deliverable is somewhat To-Be-Determined based on
initial survey and analysis by the WG.

The initial goal is to enumerate the many and often conflicting
desires around support, version skew, and upgrade path, with an eye
for possible areas for improvement.  Today these topics are covered
in a disjoint set of documents across multiple repositories and the
official Kubernetes website, so even this enumeration is complicated.

On the one hand the WG's surveys might lead to something larger and
strategic like a community agreed and implementable KEP returned
to SIG Release to operationalize.  This will come from discussion
and analysis of the release frequency and support duration relative
to adoption and adoption blocking issues.

On the other hand it might lead to more tactical project improvements
that result in a culture of and releases characterized by higher
quality and stability, such as:
* Backported patches more rigorously enforced to only contain critical fixes
* Critical APIs are stable (vN, not vNbetaX)
* kubernetes/kubernetes repo master branch kept in a releasable state
* More clear testing signal
* Higher test coverage
* Meaningful version-skew, upgrade, and downgrade tests that are immediately fixed, or patches reverted, when broken
* Reliable API schema upgrades and downgrades
* Greater supported version skew for kubectl and client-go

## Timeline and Disbanding
Working Groups are intended to be time limited endeavors.  We expect
to survey the state of support in late 2018, especially using KubeCon
China and KubeCon North America in November and December 2018
respectively as discussion forums.  In early 2019 we hope this
coalesces toward concrete proposals, with implementation coming by
later 2019.

We will evaluate our mid-year progress and next steps at KubeCon
EU 2019 and look to establish a wrap up plan at that point.

[wg-governance]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://git.k8s.io/community/committee-steering/governance/README.md
[sig-governance]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[WG LTS Proposal Presentation]: https://docs.google.com/presentation/d/1-Z-mUNIs3mUi7AdP1KwoAVNviwKrCoo3lxMb5wzCWbk/edit?usp=sharing
[WG LTS meeting minutes]: https://docs.google.com/document/d/1J2CJ-q9WlvCnIVkoEo9tAo19h08kOgUJAS3HxaSMsLA/edit?ts=5bda357d
[WG LTS YouTube Channel]: https://www.youtube.com/playlist?list=PL69nYSiGNLP13_zDqYfUjfLZ2Lu9a3pv-
