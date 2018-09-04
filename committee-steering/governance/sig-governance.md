# SIG Roles and Organizational Governance

This charter adheres to the conventions described in the [Kubernetes Charter README].

This document will be updated as needed to meet the current needs of the Kubernetes project.

## Roles

### Notes on Roles

Unless otherwise stated, individuals are expected to be responsive and active within
their roles.  Within this section "member" refers to a member of a Chair, Tech Lead or
Subproject Owner Role.  (this different from a SIG or Organization Member).

- Initial members are defined at the founding of the SIG or Subproject as part of the acceptance
  of that SIG or Subproject.
- Members *SHOULD* remain active and responsive in their Roles.
- Members taking an extended leave of 1 or more months *SHOULD*
  coordinate with other members to ensure the
  role is adequately staffed during the leave.
- Members going on leave for 1-3 months *MAY* work with other
  members to identify a temporary replacement.
- Members of a role *SHOULD* remove any other members that have not communicated a 
  leave of absence and either cannot be reached for more than 1 month or are not
  fulfilling their documented responsibilities for more than 1 month.
  This may be done through a [super-majority] vote of members, or if there are not
  enough *active* members to get a super-majority of votes cast, then removal may occur
  through a [super-majority] vote between Chairs, Tech Leads and Subproject Owners.
- Membership disagreements may be escalated to the SIG Chairs.  SIG Chair membership
  disagreements may be escalated to the Steering Committee.
- Members *MAY* decide to step down at anytime and propose a replacement.  Use lazy consensus amongst
  other members with fallback on majority vote to accept proposal.  The candidate *SHOULD* be supported by a
  majority of SIG Members or Subproject Contributors (as applicable).
- Members *MAY* select additional members through a [super-majority] vote amongst members. This
  *SHOULD* be supported by a majority of SIG Members or Subproject Contributors (as applicable).

### Chair

- Chair
  - Run operations and processes governing the SIG
  - Number: 2-3
  - Membership tracked in [sigs.yaml]

### Tech Lead

- *Optional Role*: SIG Technical Leads
  - Establish new subprojects
  - Decommission existing subprojects
  - Resolve X-Subproject technical issues and decisions
  - Number: 2-3
  - Membership tracked in [sigs.yaml]
  
### Subproject Owner

- Subproject Owners
  - Scoped to a subproject defined in [sigs.yaml]
  - Seed members established at subproject founding
  - *SHOULD* be an escalation point for technical discussions and decisions in the subproject
  - *SHOULD* set milestone priorities or delegate this responsibility
  - Number: 2-3
  - Membership tracked in [sigs.yaml]

### Member

- Members
  - *SHOULD* maintain health of at least one subproject or the health of the SIG
  - *SHOULD* show sustained contributions to at least one subproject or to the SIG
  - *SHOULD* hold some documented role or responsibility in the SIG and / or at least one subproject
    (e.g. reviewer, approver, etc)
  - *MAY* build new functionality for subprojects
  - *MAY* participate in decision making for the subprojects they hold roles in
  - Includes all reviewers and approvers in [OWNERS] files for subprojects

### Security Contact

- Security Contact
  - *MUST* be a contact point for the Product Security Team to reach out to for
    triaging and handling of incoming issues
  - *MUST* accept the [Embargo Policy]
  - Defined in `SECURITY_CONTACTS` files, this is only relevant to the root file in
    the repository. Template [SECURITY_CONTACTS]

## Organizational Management

- SIG meets bi-weekly on zoom with agenda in meeting notes
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for Kubecon
  - *SHOULD* be organized by chairs unless delegated to specific Members
- SIG updates to Kubernetes community meeting on a regular basis
  - *SHOULD* be presented by chairs unless delegated to specific Members
- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project Management

#### Subproject Creation

---

Option 1: by SIG Technical Leads

- Subprojects may be created by [KEP] proposal and accepted by [lazy-consensus] with fallback on majority vote of
  SIG Technical Leads.  The result *SHOULD* be supported by the majority of SIG members.
  - KEP *MUST* establish subproject owners
  - [sigs.yaml] *MUST* be updated to include subproject information and [OWNERS] files with subproject owners
  - Where subprojects processes differ from the SIG governance, they must document how
    - e.g. if subprojects release separately - they must document how release and planning is performed

Option 2: by Federation of Subprojects

- Subprojects may be created by [KEP] proposal and accepted by [lazy-consensus] with fallback on majority vote of
  subproject owners in the SIG.  The result *SHOULD* be supported by the majority of members.
  - KEP *MUST* establish subproject owners
  - [sigs.yaml] *MUST* be updated to include subproject information and [OWNERS] files with subproject owners
  - Where subprojects processes differ from the SIG governance, they must document how
    - e.g. if subprojects release separately - they must document how release and planning is performed
    
Subprojects may create repos under *github.com/kubernetes-sigs* through [lazy-consensus] of subproject owners.

---

- Subprojects must define how releases are performed and milestones are set.  Example:

> - Release milestones
>   - Follows the kubernetes/kubernetes release milestones and schedule
>   - Priorities for upcoming release are discussed during the SIG meeting following the preceding release and
>     shared through a PR.  Priorities are finalized before feature freeze.
> - Code and artifacts are published as part of the kubernetes/kubernetes release

### Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly following alternatives
they have defined.

- Proposing and making decisions
  - Proposals sent as [KEP] PRs and published to googlegroup as announcement
  - Follow [KEP] decision making process

- Test health
  - Canonical health of code published to <link to dashboard>
  - Consistently broken tests automatically send an alert to <link to google group>
  - SIG members are responsible for responding to broken tests alert.  PRs that break tests should be rolled back
    if not fixed within 24 hours (business hours).
  - Test dashboard checked and reviewed at start of each SIG meeting.  Owners assigned for any broken tests.
    and followed up during the next SIG meeting.

Issues impacting multiple subprojects in the SIG should be resolved by either:

- Option 1: SIG Technical Leads
- Option 2: Federation of Subproject Owners

### SIG Retirement

- In the event that the SIG is unable to regularly establish consistent quorum
  or otherwise fulfill its Organizational Management responsibilities
  - after 3 or more months it *SHOULD* be retired
  - after 6 or more months it *MUST* be retired

[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454
[OWNERS]: contributors/devel/owners.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Embargo Policy]: https://github.com/kubernetes/sig-release/blob/master/security-release-process-documentation/security-release-process.md#embargo-policy
[SECURITY_CONTACTS]: https://github.com/kubernetes/kubernetes-template-project/blob/master/SECURITY_CONTACTS
