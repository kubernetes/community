# SIG Apps Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Apps covers developing, deploying, and operating applications on Kubernetes with a focus on the application developer and application operator experience.

### In scope

#### Code, Binaries and Services

- APIs used for running applications (e.g., Workloads API)
- Tools and documentation to aid in ecosystem tool interoperability around apps (e.g., Application CRD/Controller)
- Grandfathered in tools used to aid in development of and management of workloads (e.g., Kompose)

#### Cross-cutting and Externally Facing Processes

- A discussion platform for solving app development and management problems
- Represent the needs and persona of application developers and operators

### Out of scope

- Code ownership of ecosystem tools. Discussion of the tools is in scope but ownership of them is outside the scope of Kubernetes aside from legacy situations
- Do not recommend one way to do things (e.g., picking a template language)
- Do not endorse one particular ecosystem tool

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Report the SIG status at events and community meetings wherever possible
- Actively promote diversity and inclusion in the SIG
- Uphold the Kubernetes Code of Conduct especially in terms of personal behavior and responsibility
- Chairs oversee the subproject creation process

### Deviations from [sig-governance]

- Generic technical leads are not appropriate for this SIG because sub-projects maintain their processes
- Chairs follow the Technical Leads process in the subproject creation process
- Proposing and making decisions MAY be done without the use of KEPS so long as the decision is documented in a linkable medium.

### Subproject Creation

SIG Chairs following Technical Leads process defined in [sig-governance]

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
