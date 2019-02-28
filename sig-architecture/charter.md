# SIG Architecture Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The Architecture SIG maintains and evolves the design principles of Kubernetes, and provides a consistent body of expertise necessary to ensure architectural consistency over time.

### In scope

#### Code, Binaries, Docs, and Services

- *Conformance test definitions*
- *API definitions*
- *Architectural renderings*
- *API conventions*
- *Design principles*
- *Deprecation policy*

#### Cross-cutting and Externally Facing Processes

- API review process
- Conformance test review and management
- Design documentation management
- Deprecation policy management
- Architectural initiative backlog management

### Out of scope

- KEPs that do not have architectural implications or impact are managed by their respective sponsoring SIG(s)
- The release enhancement delivery [process] that is part of the SIG-Release Release Team [subproject]
- The KEP process itself is now managed by SIG-PM and architecture subprojects are stakeholders

## Roles and Organization Management

This sig follows and adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Manage and curate the project boards associated with all sub-projects ahead of every SIG meeting so they may be discussed
- Ensure the agenda is populated 24 hours in advance of the meeting, or the meeting is then cancelled
- Report the SIG status at events and community meetings wherever possible
- Actively promote diversity and inclusion in the SIG
- Uphold the Kubernetes Code of Conduct especially in terms of personal behavior and responsibility

### Deviations from [sig-governance]

### Subproject Creation

Federation of Subprojects as defined in [sig-governance]

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-architecture/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[here]: https://docs.google.com/document/d/1TTcfvf8T_tBhGDm-wjgg31WrWjYg8IZEmo3b1mpUXh0/edit?usp=sharing
[conflicts]: https://github.com/kubernetes/community/pull/2074#discussion_r184466503
[process]: https://github.com/kubernetes/sig-release/blob/master/release-team/role-handbooks/enhancements/README.md
[subproject]: https://github.com/kubernetes/sig-release/blob/master/release-team/README.md
