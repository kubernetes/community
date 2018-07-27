# SIG Auth Governance

## Roles

Membership for roles tracked in: [OWNERS]

- **Chairs**
  - Run operations and processes governing the SIG
  - Chairs *MUST* remain active in the role and are automatically removed from the position if they
    are unresponsive for > 3 months and *MAY* be removed if not proactively working with other
    chairs to fulfill responsibilities. Chairs *MAY* arrange a short-term (<6 months) leave of
    absence with coordinated coverage of responsibilities.
  - Chairs *SHOULD* represent a diverse set of organizations.  This is predicated on interest in the
    role from multiple organizations with qualified candidates.
  - Chairs *MAY* decide to step down at any time and propose a replacement.  Use lazy consensus
    amongst chairs with fallback on majority vote to accept proposal.  This *SHOULD* be supported by
    a majority of SIG Members.
  - Chairs *MAY* select additional chairs through a [super-majority] vote amongst chairs.  This
    *SHOULD* be supported by a majority of SIG Members.
  - Number: 3
  - Defined in [sigs.yaml]

- **Technical Leads**
  - Establish new subprojects
  - Decommission existing subprojects
  - Seed technical leads nominated by legacy SIG leads from subproject owners
  - *MUST* ensure that all areas of the SIG (including those that fall outside subprojects) have
    active owners.
  - *MAY* set SIG milestone priorities.
  - *SHOULD* act as an escalation point for technical disputes in subprojects.
  - *MUST* resolve issues impacting multiple subprojects in the SIG.
  - *SHOULD* meet or communicate regularly enough to ensure they are aligned. Decisions made by
    technical leads should be consistent across the SIG. If technical leads are providing
    inconsistent direction, then they *MUST* align themselves.
  - Technical Leads *MAY* select additional tech leads through a [super-majority] vote amongst tech
    leads and chairs.  This *SHOULD* be supported by a majority of SIG Members under
    [lazy-consensus].
  - Technical Leads *MUST* remain active in the role and are automatically removed from the position
    if they are unresponsive for > 3 months and *MAY* be removed if not proactively working with
    other technical leads to fulfill responsibilities. Technical leads *MAY* arrange a short-term
    (<6 months) leave of absence with coordinated coverage of responsibilities.
  - Technical Leads *SHOULD* represent a diverse set of organizations.  This is predicated on
    interest in the role from multiple organizations with qualified candidates.
  - *SHOULD* have a record of consistent solid technical judgement and leadership over a sustained
    period especially over areas for which they are the de facto lead
  - Number: 3-7
  - Defined in [sigs.yaml] and [OWNERS] files



- **Subproject Owners**
  - Scoped to a subproject defined in [sigs.yaml]
  - Seed subproject owners *MAY* include initial proposal author(s) and/or Tech Leads
  - *MUST* be an escalation point for technical decisions and failing or flaky tests in the
    subproject
  - *MUST* proactively maintain the subproject health and status or delegate this
    responsibility. For actively developed projects, this *SHOULD* include setting milestone
    priorities, tracking release bits, and ensure adequate test coverage. For maintenance mode
    projects, this *SHOULD* include maintaining test health and ensuring compatibility with new
    features.
  - *MUST* remain active in the role and are automatically removed from the position if they are
    unresponsive for > 3 months. Subproject owners *MAY* arrange a short-term (<6 months) leave of
    absence with coordinated coverage of responsibilities.
  - *SHOULD* have a record of consistent solid technical judgement over a sustained period,
    especially over areas for which they are the owner
  - *MAY* be removed if not proactively working with other Subproject Owners to fulfill
    responsibilities.
  - *MAY* decide to step down at any time and propose a replacement.  Use [lazy-consensus] amongst
    subproject owners with fallback on majority vote to accept proposal.  This *SHOULD* be supported
    by a majority of subproject contributors (those having some role in the subproject).
  - *MAY* select additional subproject owners through a [super-majority] vote amongst subproject
    owners.  This *SHOULD* be supported by a majority of subproject contributors (through
    [lazy-consensus] with fallback on voting).
  - *SHOULD* work to ensure technical excellence within area
    - actively work to reduce complexity of the area
    - define general acceptance criteria for work as needed (code, test, documentation guidelines,
      etc)
    - consider cross-project interactions in technical decisions
    - prioritize urgent issues appropriately (e.g. addressing critical bugs, security
      vulnerabilities, or test failures)
  - Number: 1-3 per subproject
  - Defined in [sigs.yaml] [OWNERS] files

- Members
  - *MUST* maintain health of at least one subproject or the health of the SIG
  - *MUST* show sustained contributions to at least one subproject or to the SIG
  - *SHOULD* hold some documented role or responsibility in the SIG and / or at least one subproject
    (e.g. reviewer, approver, etc)
  - *MAY* build new functionality for subprojects
  - *MAY* participate in decision making for the subprojects they hold roles in
  - *SHOULD* include all reviewers and approvers in [OWNERS] files for subprojects
  - *MUST* remain active in the role and are automatically removed from the position if they are
    unresponsive for > 1 year.
  - *MAY* be removed if not proactively working with other Subproject Owners to fulfill
    responsibilities.

## Organizational management

- SIG meets bi-weekly on [zoom] with [agenda in meeting notes]
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for Kubecon
  - *SHOULD* be organized by chairs unless delegated to specific Members

- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project management

#### Subproject creation

- Subprojects *MAY* be created by [KEP] proposal and *MUST* be approved by at least 1 tech lead, and
  accepted by [lazy-consensus] with fallback on majority vote of SIG Technical Leads.  The result
  *SHOULD* be supported by the majority of SIG members.
  - KEP *MUST* establish subproject owners
  - [sigs.yaml] *MUST* be updated to include subproject information and [OWNERS] files with
    subproject owners
  - Where subprojects processes differ from the SIG, they must document how
    - e.g. if subprojects release separately - they must document how release and planning is
      performed

### Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly following alternatives
they have defined.

- Proposing and making decisions
  - Proposals *SHOULD* be sent as [KEP] PRs and published to [kubernetes-sig-auth] as an
    announcement
  - Proposals *MUST* be presented and discussed in at least 1 SIG auth meeting
    - If the authors are unable to attend, another community member can be found to champion the
      project.
  - Proposals *SHOULD* follow [KEP] decision making process
    - For proposed extensions to a subproject, the subproject *MUST* be identified in the KEP, and
      approvers *MUST* include subproject owners

- Test health
  - Canonical health of code published to [testgrid]
  - Consistently broken tests automatically send an alert to [kubernetes-sig-auth-test-failures]
  - SIG members are responsible for responding to broken tests alert.  PRs that break tests should
    be rolled back if not fixed within 1 business day.
  - Test dashboard checked and reviewed at start of each SIG meeting.  Owners assigned for any
    broken tests and followed up during the next SIG meeting.


[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[OWNERS]: https://github.com/kubernetes/community/blob/master/sig-auth/OWNERS
[agenda and meeting notes]: https://docs.google.com/document/d/1woLGRoONE3EBVx-wTb4pvp4CI7tmLZ6lS26VTbosLKM/view
[zoom]: https://zoom.us/my/k8s.sig.auth
[testgrid]: https://k8s-testgrid.appspot.com/sig-auth#Summary
[kubernetes-sig-auth]: https://groups.google.com/forum/#!forum/kubernetes-sig-auth
[kubernetes-sig-auth-test-failures]: https://groups.google.com/forum/#!forum/kubernetes-sig-auth-test-failures
