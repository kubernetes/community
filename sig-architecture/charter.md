# SIG Architecture Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The Architecture SIG maintains and evolves the design principles of Kubernetes, and provides a consistent body of expertise necessary to ensure architectural consistency over time.

### In scope

#### Code, Binaries and Services

- *Conformance test definitions*
- *API definitions*

#### Cross-cutting and Externally Facing Processes

- API review process
- Conformance test review and management
- The Kubernetes Enhancement Proposal (KEP) process and KEP reviews
- Design documentation tracking
- Deprecation policies
- Architectural initiative backlog management

### Out of scope

- KEPs that do not have architectural implications
- The release enhancement delivery process

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Manage and curate the project boards associated with all sub-projects ahead of every SIG meeting so they may be discussed
- Ensure the agenda is populated 24 hours in advance of the meeting, or the meeting is then cancelled
- Report the SIG status at events and community meetings wherever possible
- Actively promote diversity and inclusion in the SIG
- Uphold the Kubernetes Code of Conduct especially in terms of personal behavior and responsibility

### Deviations from [sig-governance]

- Generic technical leads are not appropriate for this SIG because sub-projects maintain their own documented, accountable processes

### Subproject Creation

Federation of Subprojects as defined in [sig-governance]

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-architecture/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
