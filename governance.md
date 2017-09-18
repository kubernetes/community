This is a Work in Progress, documenting approximately how we have been operating up to this point.

# Principles

The Kubernetes community adheres to the following principles:
* Open: Kubernetes is open source. See repository guidelines and CLA, below.
* Welcoming and respectful: See Code of Conduct, below.
* Transparent and accessible: Work and collaboration should be done in public. See SIG governance, below.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with project objectives, [scope](http://kubernetes.io/docs/whatisk8s/), and [design principles](contributors/design-proposals/architecture/principles.md).

# Code of Conduct

The Kubernetes community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

As a member of the Kubernetes project, you represent the project and your fellow contributors. 
We value our community tremendously and we'd like to keep cultivating a friendly and collaborative
environment for our contributors and users. We want everyone in the community to have 
[positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).

# Community membership

See [community membership]

# Community groups

The project has 3 main types of groups:
1. Special Interest Groups, SIGs
2. Working Groups, WGs
3. Committees

Note that the project is also in the process of forming a Steering
Committee, details of which will be documented soon.

## SIGs

The Kubernetes project is organized primarily into Special Interest
Groups, or SIGs. Each SIG is comprised of members from multiple
companies and organizations, with a common purpose of advancing the
project with respect to a specific topic, such as Networking or
Documentation. Our goal is to enable a distributed decision structure
and code ownership, as well as providing focused forums for getting
work done, making decisions, and onboarding new contributors. Every
identifiable subpart of the project (e.g., github org, repository,
subdirectory, API, test, issue, PR) is intended to be owned by some
SIG. 

Areas covered by SIGs may be vertically focused on particular
components or functions, cross-cutting/horizontal, spanning many/all
functional areas of the project, or in support of the project
itself. Examples:
* Vertical: Network, Storage, Node, Scheduling, Big Data
* Horizontal: Scalability, Architecture
* Project: Testing, Release, Docs, PM, Contributor Experience

SIGs must have at least one and ideally two SIG leads at any given
time. SIG leads are intended to be organizers and facilitators,
responsible for the operation of the SIG and for communication and
coordination with the other SIGs, the Steering Committee, and the
broader community.

We still have work to do to unify people organization via SIGs and
code organization using [OWNERS](contributors/devel/owners.md), and to
find homes for all subparts of the project, such as the build
system. We also need to create a charter for each SIG to more clearly
specify its scope (topics, subsystems, code repos and directories),
responsibilities, areas of authority, how members and roles of
authority/leadership are selected/granted, how decisions are made, and
how conflicts are resolved. A template for intra-SIG governance should
be developed in order to simplify SIG creation, but SIGs should be
relatively free to customize or change how they operate, within some
broad guidelines and constraints imposed by cross-SIG processes (e.g.,
the release process) and assets (e.g., the kubernetes repo).

A primary reason that SIGs exist is as forums for collaboration. 
Much work in a SIG should stay local within that SIG. However, SIGs
must communicate in the open, ensure other SIGs and community members
can find notes of meetings, discussions, designs, and decisions, and
periodically communicate a high-level summary of the SIG's work to the
community.

See [sig governance] for more details about current SIG operating
mechanics, such as mailing lists, meeting times, etc.

## Working Groups

We need community rallying points to facilitate discussions/work
regarding topics that are too young/short-lived, or narrow/small, or
decoupled from specific efforts to be tied to ownership as SIGs or
Committees are, this is the purpose of Working Groups (WG). The intent
is to make Working Groups relatively easy to create and to deprecate,
once inactive.

## Committees

Some topics, such as Security or Code of Conduct, require
discretion. Whereas SIGs are voluntary groups which operate in the
open and anyone can join, Committees do not have open membership and do
not always operate in the open.  The steering committee can form
committees as needed, for bounded or unbounded duration.  Membership
of a committee is decided by the steering committee.  Like a SIG, a
committee has a charter and a lead, and will report to the steering
committee periodically, and to the community as makes sense, given the
charter.

## Cross-project Communication and Coordination

While most work shouldn’t require expensive coordination with other
SIGs, there will be efforts (features, refactoring, etc.) that cross
SIG boundaries.  In this case, it is expected that the SIGs coordinate
with each other and come to mutually agreed solutions. In some cases,
it may make sense to form a Working Group for joint work.  Cross-SIG
coordination will naturally require more time and implies a certain
amount of overhead.  This is intentional to encourage changes to be
well encapsulated whenever possible.

On the other hand, several SIGs do have project-wide impact, for
example Release, Testing, and API Machinery. Even those that do not
may sometimes need to make changes or impose new processes or
conventions that affect other SIGs. In these cases, project-wide
communication processes will need to be followed. For example,
proposals with project-wide impact will need to be announced more
broadly, with the opportunity for members of other SIGs to provide
feedback and guidance. However, the SIG that owns the area, according
to its charter, will own the decision. In the case of extended debate
or deadlock, decisions may be escalated to the Steering Committee,
which is expected to be uncommon.

The exact processes and guidelines for such cross-project
communication have yet to be formalized, but when in doubt, use
kubernetes-dev@googlegroups.com and make an announcement at the
community meeting.

# Repository guidelines

All repositories under Kubernetes github orgs, such as kubernetes and kubernetes-incubator,
should follow the procedures outlined in the [incubator document](incubator.md). All code projects
use the [Apache Licence version 2.0](LICENSE). Documentation repositories should use the
[Creative Commons License version 4.0](https://github.com/kubernetes/kubernetes.github.io/blob/master/LICENSE).

# Incubator process

See [incubator process]

# CLA

All contributors must sign the CNCF CLA, as described [here](CLA.md).

[community membership]: /community-membership.md
[sig governance]: /sig-governance.md
[incubator process]: /incubator.md

[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/governance.md?pixel)]()
