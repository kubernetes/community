# SIG Apps Charter

## Scope

SIG Apps Covers developing, deploying, and operating applications on Kubernetes with a focus on the application developer and application operator experience.

### Goals

* Discussion of how to define, develop, and run applications in Kubernetes. This is in order better understand needs, improve on Kubernetes solutions, and share lessons to grow the number of applications running on Kubernetes
* Foster an ecosystem of tools to aid in developing, deploying, and operating applications on Kubernetes
* Be the voice of the people running applications in Kubernetes. This includes providing input on Kuberentes features, suggesting new features, and sharing the perspective of application developers and operators in the development of Kubernetes
* Enable varying workloads (e.g., databases, web services, machine learning) to run on Kubernetes through development of the Workloads API
* Develop subprojects that foster interoperability between tools used to develop, deploy, and operate applications
* Create documentation to enable application developers, application operators, and supporting tool developers to leverage Kubernetes
* Develop and maintain Kubernetes subprojects that aid in the development, deployment, and operation of applications on Kubernetes

### Non-goals

* Do not endorse one particular ecosystem tool
* Do not pick which apps to run on top of Kubernetes
* Do not recommend one way to do things (e.g., picking a template language)

## Roles

Membership for roles tracked in: [sigs.yaml]

- Chair
  - Run operations and processes governing the SIG
  - Seed members established at SIG founding
  - Chairs *MAY* decide to step down at anytime and propose a replacement.  Use lazy consensus amongst
    chairs with fallback on majority vote to accept proposal.  This *SHOULD* be supported by a majority of
    SIG Members.
  - Chairs *MAY* select additional chairs through a [super-majority] vote amongst chairs.  This
    *SHOULD* be supported by a majority of SIG Members.
  - Chairs *MUST* remain active in the role and are automatically removed from the position if they are
    unresponsive for > 3 months and *MAY* be removed if not proactively working with other chairs to fulfill
    responsibilities.
  - Number: 2-3
  - Defined in [sigs.yaml]

- Subproject Owners
  - Scoped to a subproject defined in [sigs.yaml]
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
  - Number: Minimum of 3
  - Defined in [sigs.yaml] [OWNERS] files

- Members
  - *MUST* maintain health of at least one subproject or the health of the SIG
  - *MUST* show sustained contributions to at least one subproject or to the SIG
  - *SHOULD* hold some documented role or responsibility in the SIG and / or at least one subproject
    (e.g. reviewer, approver, etc)
  - *MAY* build new functionality for subprojects
  - *MAY* participate in decision making for the subprojects they hold roles in
  - Includes all reviewers and approvers in [OWNERS] files for subprojects

## Organizational management

- SIG meets weekly on zoom with agenda in meeting notes
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for Kubecon
  - *SHOULD* be organized by chairs unless delegated to specific Members

- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project management

#### Subproject creation

Subprojects may be created by [KEP] proposal and accepted by [lazy-consensus] with fallback on majority vote of
SIG Chairs.  The result *SHOULD* be supported by the majority of SIG members.
  - KEP *MUST* establish subproject owners
  - [sigs.yaml] *MUST* be updated to include subproject information and [OWNERS] files with subproject owners
  - Where subprojects processes differ from the SIG governance, they must document how
    - e.g. if subprojects release separately - they must document how release and planning is performed

Subprojects must define how releases are performed and milestones are set.  Example:

> - Release milestones
>   - Follows the kubernetes/kubernetes release milestones and schedule
>   - Priorities for upcoming release are discussed during the SIG meeting following the preceding release and
>     shared through a PR.  Priorities are finalized before feature freeze.
> - Code and artifacts are published as part of the kubernetes/kubernetes release

#### Subproject retirement

Subprojects may be retired, where they are archived to the GitHub kubernetes-retired organization, when they are
no longer supported based on the following criteria.

- A subproject is no longer supported when there are no active owners with activity on the project for the following time:
  - Subprojects with no known users can be retired after being unsupported for > 3 months
  - Subprojects with known users may be retired after providing at least 6 months notification of retirement
- Use [lazy-consensus] amongst chairs with fallback on majority vote to decide to retire.  This *SHOULD* be
  supported by a majority of SIG Members.

### Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly following alternatives
they have defined.

- Proposing and making decisions
  - Proposals against sub-projects, not part of core Kubernetes, will have issues filed against their repositories
  - When issues are used for Proposals sub-project will have their own decision making process
  - Proposals against core Kubernetes sent as [KEP] PRs and published to kubernetes-sig-apps as announcement
  - When KEPs are used follow [KEP] decision making process

- Test health
  - Canonical health of code published to a dashboard.
  - Consistently broken tests automatically send an alert to a mailing list for the subproject maintainers.
  - SIG members are responsible for responding to broken tests alert.  PRs that break tests should be rolled back
    if not fixed within 24 hours (business hours).

Issues impacting multiple subprojects in the SIG should be resolved by SIG Chairs

[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454
[OWNERS]: contributors/devel/owners.md



