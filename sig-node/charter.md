# SIG Node Governance

## Scope

SIG Node is responsible for the components that support the controlled
interactions between pods and host resources.  We manage the lifecycle of pods
that are scheduled to a node.  We focus on enabling a broad set of workload
types, including hardware and performance sensitive workloads.  We maintain
isolation boundaries between pods on a node, as well as the pod and the host.  We
aim to continuously improve node reliability.

### In scope

The following topics fall under ownership of this SIG.

* Kubelet related features (e.g. Pod lifecycle)
* Node level performance and scalability (with [sig-scalability](../sig-scalability))
* Node reliability (prooblem detection and remediation)
* Node lifecycle management (with [sig-cluster-lifecycle](../sig-cluster-lifecycle))
* Container runtimes
* Device management
* Images, package management
* Host resource management (with [sig-scheduling](../sig-scheduling))
* Hardware discovery
* Issues related to node, pod, container monitoring (with [sig-instrumentation](../sig-instrumentation))
* Node level security and Pod isolation (with [sig-auth](../sig-auth))
* Host OS and/or kernel interactions (to a limited extent)

### Out of scope

The following topics are out of the scope of this SIG

* network management
* persistent storage management

## Roles

Membership for roles tracked in: [OWNERS]

- **Chairs**
  - Run operations and processes governing the SIG
  - *MUST* remain active in the role and are  removed from the position if they
    are unresponsive for > 3 months and *MAY* be removed if not proactively
    working with other chairs to fulfill responsibilities.  To remove an
    inactive chair, a majority vote from existing chairs should occur. If a
    [super-majority] is not possible, the [steering-committee] may sponsor
    removal.
  - *MAY* declare an intention to take sabbatical for no more than 3 months in a
    calendar year.  Chairs should notify existing chairs of their intent, and
    absent at least 2 remaining chairs, *MUST* nominate a temporary replacement
    from the existing set of technical leads during this period.
  - *SHOULD* represent a diverse set of organizations.  This is predicated on
    interest in the role from multiple organizations with qualified candidates.
    Qualification is measured by existing chairs and *SHOULD* be supported by a
    majority of SIG Members under [lazy-consensus].
  - *MAY* decide to step down at any time and propose a replacement.  Use lazy
    consensus amongst chairs with fallback on majority vote to accept proposal.
    This *SHOULD* be supported by a majority of SIG Members.
  - *MAY* select additional chairs through a [super-majority] vote amongst
    chairs.  This *SHOULD* be supported by a majority of SIG Members.
  - Number: 1-3
  - Defined in [sigs.yaml]

- **Technical Leads**
  - Technical leads seeded by legacy SIG chairs from subproject owners
  - Establish new subprojects
  - Decommission existing subprojects  
  - Sponsor working groups
  - End of life working groups
  - *MUST* ensure that all areas of the SIG (including those that fall outside
    subprojects) have active owners.
  - *MAY* set SIG milestone priorities.
  - *SHOULD* act as an escalation point for technical disputes in subprojects.
  - *MUST* resolve issues impacting multiple subprojects in the SIG.
  - *SHOULD* meet or communicate regularly enough to ensure they are aligned.
    Decisions made by technical leads should be consistent across the SIG. If
    technical leads are providing inconsistent direction, then they *MUST* align
    themselves.
  - *MAY* select additional tech leads through a [super-majority] vote amongst
    tech leads and chairs.  This *SHOULD* be supported by a majority of SIG
    Members under [lazy-consensus].
  - *SHOULD* represent a diverse set of organizations.  This is predicated on
    interest in the role from multiple organizations with qualified candidates.
  - *MUST* remain active in the role and are automatically removed from the
    position if they are unresponsive for > 3 months and *MAY* be removed if not
    proactively working with other technical leads to fulfill responsibilities.
    Removal of a technical lead requires [super-majority] vote amongst tech
    leads and chairs.
  - *MAY* declare an intention to take sabbatical for no more than 3 months in a
    calendar year.  Chairs should notify existing chairs of their intent, and
    absent at least 2 remaining chairs, *MUST* nominate a temporary replacement
    from the existing set of technical leads during this period.    
  - Technical Leads *SHOULD* represent a diverse set of organizations.  This is
    predicated on interest in the role from multiple organizations with
    qualified candidates.
  - *SHOULD* have a record of consistent solid technical judgement and
    leadership over a sustained period especially over areas for which they are
    the de-facto lead
  - Number: 3-7
  - Defined in [sigs.yaml] and [OWNERS] files

- **Subproject Owners**
  - Scoped to a subproject defined in [sigs.yaml]
  - Seed subproject owners *MAY* include initial proposal author(s) and/or Tech
    Leads
  - *MUST* be an escalation point for technical decisions and failing or flaky
    tests in the subproject
  - *MUST* proactively maintain the subproject health and status or delegate
    this responsibility. For actively developed projects, this *SHOULD* include
    setting milestone priorities, tracking release bits, and ensure adequate
    test coverage. For maintenance mode projects, this *SHOULD* include
    maintaining test health and ensuring compatibility with new features.
  - *MUST* remain active in the role and are automatically removed from the
    position if they are unresponsive for > 3 months.
  - *MAY* declare an intention to take sabbatical for no more than 3 months in a
    calendar year.  Should notify existing owners of their intent, and *MAY*
    nominate a temporary replacement from the existing set of owners during this
    period.
  - *SHOULD* have a record of consistent solid technical judgement over a
    sustained period, especially over areas for which they are the owner
  - *MAY* be removed if not proactively working with other Subproject Owners to
    fulfill responsibilities.
  - *MAY* decide to step down at any time and propose a replacement.  Use
    [lazy-consensus] amongst subproject owners with fallback on majority vote to
    accept proposal.  This *SHOULD* be supported by a majority of subproject
    contributors (those having some role in the subproject).
  - *MAY* select additional subproject owners through a [super-majority] vote
    amongst subproject owners.  This *SHOULD* be supported by a majority of
    subproject contributors (through [lazy-consensus] with fallback on voting).
  - *SHOULD* work to ensure technical excellence within area
    - actively work to reduce complexity of the area
    - define general acceptance criteria for work as needed (code, test,
      documentation guidelines, etc)
    - consider cross-project interactions in technical decisions
    - prioritize urgent issues appropriately (e.g. addressing critical bugs,
      security vulnerabilities, or test failures)
  - Number: 1-7 per subproject (depending on size and scope)
  - Defined in [sigs.yaml] [OWNERS] files

- Members
  - *MUST* maintain health of at least one subproject or the health of the SIG
  - *MUST* show sustained contributions to at least one subproject or to the SIG
  - *SHOULD* hold some documented role or responsibility in the SIG and / or at
    least one subproject (e.g. reviewer, approver, etc)
  - *MAY* build new functionality for subprojects
  - *MAY* participate in decision making for the subprojects they hold roles in
  - *SHOULD* include all reviewers and approvers in [OWNERS] files for
    subprojects
  - *MUST* remain active in the role and may be removed from the position if
    they are unresponsive for > 1 year.
  - *MAY* be removed if not proactively working with other Subproject Owners to
    fulfill responsibilities.

## Organizational management

- SIG meets weekly on [zoom] with [agenda in meeting notes]
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for Kubecon
  - *SHOULD* be organized by chairs unless delegated to specific Members

- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project management

#### Subproject creation

- Subprojects *MAY* be created by [KEP] proposal and *MUST* be approved by at
  least 1 tech lead, and accepted by [lazy-consensus] with fallback on majority
  vote of SIG Technical Leads.  The result *SHOULD* be supported by the majority
  of SIG members.
  - KEP *MUST* establish subproject owners
  - [sigs.yaml] *MUST* be updated to include subproject information and [OWNERS]
    files with subproject owners
  - Where subprojects processes differ from the SIG, they must document how
    - e.g. if subprojects release separately - they must document how release
      and planning is performed
  - Subprojects may not be seeded from existing repositories pending CNCF
    process definition

### Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly
following alternatives they have defined.

- Proposing and making decisions
  - Proposals *SHOULD* be sent as [KEP] PRs and published to
    [kubernetes-sig-node] as an announcement
  - Proposals *MUST* be presented and discussed in at least 1 SIG node meeting
    - If the authors are unable to attend, another community member can be found
      to champion the proposal.
  - Proposals *SHOULD* follow [KEP] decision making process
    - For proposed extensions to a subproject, the subproject *MUST* be
      identified in the KEP, and approvers *MUST* include subproject owners

- Test health
  - Canonical health of code published to [testgrid]
  - Consistently broken tests automatically send an alert to
    [kubernetes-sig-node-test-failures]
  - SIG members are responsible for responding to broken tests alert.  PRs that
    break tests should be rolled back if not fixed within 1 business day.
  - Test dashboard checked and reviewed at start of each SIG meeting.  Owners
    assigned for any broken tests and followed up during the next SIG meeting.

[steering-committee]: https://github.com/kubernetes/community/blob/master/committee-steering/OWNERS
[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[OWNERS]: https://github.com/kubernetes/community/blob/master/sig-auth/OWNERS
[agenda and meeting notes]: https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit
[zoom]: https://zoom.us/j/4799874685
[testgrid]: https://k8s-testgrid.appspot.com/sig-node#Summary
[kubernetes-sig-node]: https://groups.google.com/forum/#!forum/kubernetes-sig-node
[kubernetes-sig-node-test-failures]: https://groups.google.com/forum/#!forum/kubernetes-sig-node-test-failures