# WG Kubernetes AI Conformance Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The goal of this group is to define a standardized set of capabilities, APIs, and configurations that a Kubernetes cluster must offer to reliably and efficiently run AI/ML workloads. This initiative aims to simplify AI/ML operations on Kubernetes, accelerate adoption, guarantee interoperability and portability for AI workloads, and enable ecosystem growth on an industry-standard foundation.

### In scope

#### Code, Binaries and Services

- The primary artifact will be the (working title for now) "CNCF Kubernetes AI Conformance" specification and a suite of tests to demonstrate conformance.

#### Cross-cutting and Externally Facing Processes

- The Working Group will consider its primary problem-solving objective complete upon the successful definition and initial adoption of a stable (working title for now) "CNCF Kubernetes AI Conformance" specification.
- Once the foundational conformance is established and widely recognized, the ongoing maintenance and evolution of the conformance would be evaluated, and could ideally transition to a Special Interest Group (SIG) with a long-term charter, at which point the Working Group would dissolve.

### Out of scope

This WG is not responsible for maintenance and evolution of the conformance program, which will be run by the CNCF. Some aspects may be owned by existing SIGs (Architecture, Testing) alongside the CNCF, such as Kubernetes specific testing and tooling, and the pre-existing Kubernetes conformance definition and tests.

## Roles and Organization Management

This WG follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- The Working Group will have designated Chair(s) responsible for guiding discussions and ensuring progress.
- A note-taker will be assigned for each meeting, and active participation from all contributors will be encouraged.
- Agendas and meeting notes will be publicly accessible.

### Stakeholder SIGs

- SIG Architecture (Sponsoring SIG)
- SIG Testing


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
