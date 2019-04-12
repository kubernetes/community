# SIG Roles and Organizational Governance

This charter adheres to the conventions described in the [Kubernetes Charter README]. It will be updated as needed to meet the current needs of the Kubernetes project.

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow these guidelines:

- Create a charter and have it approved according to the [SIG charter process]
- Meet regularly, at least for 30 minutes every 3 weeks, except November and December
- Keep up-to-date meeting notes, linked from the SIG's page in the community repo
- Record meetings and make them publicly available on a YouTube playlist
- Report activity in the weekly community meeting at least once every quarter
- Participate in release planning meetings and retrospectives, and burndown meetings, as needed
- Ensure related work happens in a project-owned github org and repository, with code and tests explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response, bug fixes, etc.
- Use the [forums provided] as the primary means of working, communicating, and collaborating, as opposed to private emails and meetings

The process for setting up a SIG or Working Group (WG) is listed in the [sig-wg-lifecycle] document.
## Roles

### Notes on Roles

Within this section "member" refers to a member of a Chair, Tech Lead or
Subproject Owner Role.  (this different from a SIG or Organization Member).  

- Initial members are defined at the founding of the SIG or Subproject as part of the acceptance
  of that SIG or Subproject.
- Members *SHOULD* remain active and responsive in their Roles.
- Members *MUST* be [community members] to be eligible to hold a leadership role
  within a SIG.
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
  - *MUST* be a contact point for the Product Security Committee to reach out to for
    triaging and handling of incoming issues
  - *MUST* accept the [Embargo Policy]
  - Defined in `SECURITY_CONTACTS` files, this is only relevant to the root file in
    the repository. Template [SECURITY_CONTACTS]

## Organizational Management

- SIG meets bi-weekly on zoom with agenda in meeting notes
  - *SHOULD* be facilitated by chairs unless delegated to specific Members
- SIG overview and deep-dive sessions organized for KubeCon/CloudNativeCon
  - *SHOULD* be organized by chairs unless delegated to specific Members
- SIG updates to Kubernetes community meeting on a regular basis
  - *SHOULD* be presented by chairs unless delegated to specific Members
- Contributing instructions defined in the SIG CONTRIBUTING.md

### Project Management
In addition, SIGs have the following responsibilities to SIG PM:
- identify SIG annual roadmap
- identify all SIG features in the current release
- actively track / maintain SIG features within [k/enhancements]
- attend [SIG PM] meetings, as needed / requested

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

[SIG PM]: https://github.com/kubernetes/community/tree/master/sig-pm
[k/enhancements]: https://github.com/kubernetes/enhancements
[forums provided]: /communication/README.md

[lazy-consensus]: http://en.osswiki.info/concepts/lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://git.k8s.io/enhancements/keps/YYYYMMDD-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454
[OWNERS]: contributors/devel/owners.md
[SIG Charter process]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Embargo Policy]: https://git.k8s.io/security/private-distributors-list.md#embargo-policy
[SECURITY_CONTACTS]: https://github.com/kubernetes/kubernetes-template-project/blob/master/SECURITY_CONTACTS
[sig-wg-lifecycle]: /sig-wg-lifecycle.md
[community members]: /community-membership.md
