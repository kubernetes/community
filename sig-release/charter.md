# SIG Release Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

- Production of Kubernetes releases on a reliable schedule
- Ensure there is a consistent group of community members in place to support the release process across time
- Provide guidance and tooling to facilitate the production of automated releases
- Serve as a tightly integrated partner with other SIGs to empower SIGs to integrate their repositories into the release process

### In scope

- Ensuring quality Kubernetes releases
  - Defining and staffing release roles to manage the resolution of release blocking criteria
  - Defining and driving development processes (e.g. merge queues, cherrypicks) and release processes
    (e.g. burndown meetings, cutting beta releases) with the intent of meeting the release schedule
  - Managing the creation of release specific artifacts, including:
    - Code branches
    - Binary artifacts
    - Release notes
- Continually improving release and development processes
  - Working closely with SIG Contributor Experience to define and build tools to facilitate release process (e.g. dashboards)
  - Working closely with SIG Testing to determine and implement tests, automation, and labeling required for stable releases
  - Working with downstream communities responsible for packaging Kubernetes releases
  - Working with other SIGs to agree upon the responsibilities of their SIG with respect to the release
  - Defining and collecting metrics related to the release in order to measure progress over each release
  - Facilitating release retrospectives
- Collaborating with downstream communities which build artifacts from Kubernetes releases

### Out of scope

#### Support

SIG Release itself is not responsible for end user support or creation of patches for support streams. There are support forums for end users to ask questions and report bugs, subject matter experts in other SIGs triage and address issues and when necessary mark bug fixes for inclusion in a patch release.

## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance] and opts-in to updates and modifications to [sig-governance].

Specifically, the common guidelines (see: [sig-governance]) for continuity of membership within roles in the SIG are followed.

### Deviations from [sig-governance]

- SIG Release subprojects have subproject chairs
- SIG Release does not have top-level SIG Technical Leads. With few exceptions, technical decisions should be handled within the scope of the relevant SIG Release subproject.

#### SIG Membership

SIG Release has a concept of membership. SIG members can be occasionally called on to assist with decision making, especially as it relates to gathering historical context around existing policies and enacting new policy.

SIG Release membership is represented by membership in the `sig-release` GitHub team.

SIG Release membership is defined as the set of Kubernetes contributors included in:
- All SIG Release top-level subproject OWNERS files
- All documented former Release Team members holding Lead roles e.g., Enhancements Lead, Patch Release Team

Subproject `approvers` and incoming Release Team Leads should ensure that new members are added to the `sig-release` GitHub team.

SIG Release Chairs will periodically review the `sig-release` GitHub team to ensure it remains accurate and up-to-date.

All SIG Release roles will be filled by SIG Release members, except where explicitly defined in other policy. A notable exception to this would be Release Team Shadows.

It may be implied, given the criteria for SIG membership, but to be explicit:
- SIG Release membership is representative of people who actively contribute to the health of the SIG. Given that, those members should also be enabled to help drive SIG Release policy.
- SIG Chairs should represent the sentiment of and facilitate the decision making by SIG Members.

### Subproject Creation

Subprojects must be created by [KEP] proposal and accepted by [lazy-consensus].

In the event that lazy consensus cannot be reached:
- Fallback to a majority decision by SIG Chairs
- SIG Release Members may override the majority decision of SIG Chairs by a supermajority (60%)

Additional requirements:
- KEP must establish subproject chairs
- [sigs.yaml] must be updated to include subproject information and OWNERS files with subproject chairs


[KEP]: https://git.k8s.io/enhancements/keps/YYYYMMDD-kep-template.md
[Kubernetes Charter README]: /committee-steering/governance/README.md
[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[sig-governance]: /committee-steering/governance/sig-governance.md
[sigs.yaml]: /sigs.yaml
