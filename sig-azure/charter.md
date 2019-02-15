# SIG Azure Charter

_The following is a charter for the Kubernetes Special Interest Group for Azure. It delineates the roles of SIG leadership, SIG members, as well as the organizational processes for the SIG, both as they relate to project management and technical processes for SIG subprojects._

## Roles

### SIG Chairs

- Run operations and processes governing the SIG 
- Seed members established at SIG founding 
- Chairs MAY decide to step down at anytime and propose a replacement. Use lazy consensus amongst chairs with fallback on majority vote to accept proposal. This SHOULD be supported by a majority of SIG Members.
- Chairs MAY select additional chairs through a [super-majority] vote amongst chairs. This SHOULD be supported by a majority of SIG Members.
- Chairs MUST remain active in the role and are automatically removed from the position if they are unresponsive for &gt; 3 months and MAY be removed if not proactively working with other chairs to fulfill responsibilities. Coordinated leaves of absence serve as exception to this requirement.
- Number: 2 - 3 
- Defined in [sigs.yaml]

### SIG Technical Leads

- Establish new subprojects 
- Decommission existing subprojects 
- Resolve X-Subproject technical issues and decisions 
- Technical Leads MUST remain active in the role and are automatically removed from the position if they are unresponsive for &gt; 3 months and MAY be removed if not proactively working with other chairs to fulfill responsibilities. Coordinated leaves of absence serve as exception to this requirement. 
- Number: 2 - 3 
- Defined in [sigs.yaml]

### Subproject Owners

- Scoped to a subproject defined in [sigs.yaml]
- Seed members established at subproject founding
- MUST be an escalation point for technical discussions and decisions in the subproject
- MUST set milestone priorities or delegate this responsibility 
- MUST remain active in the role and are automatically removed from the position if they are unresponsive for &gt; 3 months. Coordinated leaves of absence serve as exception to this requirement.
- MAY be removed if not proactively working with other Subproject Owners to fulfill responsibilities.
- MAY decide to step down at anytime and propose a replacement. Use [lazy-consensus] amongst subproject owners with fallback on majority vote to accept proposal. This SHOULD be supported by a majority of subproject contributors (those having some role in the subproject).
- MAY select additional subproject owners through a [super-majority] vote amongst subproject owners. This SHOULD be supported by a majority of subproject contributors (through [lazy-consensus] with fallback on voting).
- Number: 3 - 5
- Defined in [sigs.yaml] [OWNERS] files

**IMPORTANT**

_With regards to leadership roles i.e., Chairs, Technical Leads, and Subproject Owners, we MUST, as a SIG, ensure that positions are held by a committee of members across a diverse set of companies. This allows for thoughtful discussion and structural management that can serve the needs of every consumer of Kubernetes on Azure._

### Members

- MUST maintain health of at least one subproject or the health of the SIG 
- MUST show sustained contributions to at least one subproject or to the SIG 
- SHOULD hold some documented role or responsibility in the SIG and / or at least one subproject (e.g. reviewer, approver, etc) 
- MAY build new functionality for subprojects 
- MAY participate in decision making for the subprojects they hold roles in 
- Includes all reviewers and approvers in [OWNERS] files for subprojects

## Organizational management

- SIG meets bi-weekly on zoom with agenda in meeting notes 
  - SHOULD be facilitated by chairs unless delegated to specific Members 
- SIG overview and deep-dive sessions organized for KubeCon/CloudNativeCon 
  - SHOULD be organized by chairs unless delegated to specific Members 
- Contributing instructions defined in the SIG CONTRIBUTING.md 

### Project management

#### Subproject creation

Subprojects
may be created by [KEP] proposal and accepted by [lazy-consensus] with fallback on majority vote of SIG Technical Leads. The result SHOULD be supported by the majority of SIG members.

- KEP MUST establish subproject owners
- [sigs.yaml] MUST be updated to include subproject information and [OWNERS] files with subproject owners
- Where subprojects processes differ from the SIG governance, they MUST document how 
  - e.g., if subprojects release separately - they must document how release and planning is performed 

Subprojects must define how releases are performed and milestones are set. 

Example:
- Release milestones 
  - Follows the kubernetes/kubernetes release milestones and schedule 
  - Priorities for upcoming release are discussed during the SIG meeting following the preceding release and shared through a PR. Priorities are finalized before feature freeze. 
- Code and artifacts are published as part of the kubernetes/kubernetes release 

### Technical processes

Subprojects of the SIG MUST use the following processes unless explicitly following alternatives they have defined.

- Proposing and making decisions 
  - Proposals sent as [KEP] PRs and published to Google group as announcement 
  - Follow [KEP] decision making process 

- Test health 
  - Canonical health of code published to
  - Consistently broken tests automatically send an alert to
  - SIG members are responsible for responding to broken tests alert. PRs that break tests should be rolled back if not fixed within 24 hours (business hours). 
  - Test dashboard checked and reviewed at start of each SIG meeting. Owners assigned for any broken tests. and followed up during the next SIG meeting. 

Issues impacting multiple subprojects in the SIG should be resolved by SIG Technical Leads, with fallback to consensus of subproject owners.

[lazy-consensus]: http://en.osswiki.info/concepts/lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://git.k8s.io/enhancements/keps/YYYYMMDD-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454
[OWNERS]: /contributors/guide/owners.md
