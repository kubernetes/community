# WG Kubernetes AI Conformance Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [wg-governance].

## Scope

The goal of this group is to define a standardized set of capabilities, APIs, and configurations that a Kubernetes cluster must offer to reliably and efficiently run AI/ML workloads. This initiative aims to simplify AI/ML operations on Kubernetes, accelerate adoption, guarantee interoperability and portability for AI workloads, and enable ecosystem growth on an industry-standard foundation.

### In scope

#### Code, Binaries and Services

- The primary artifact will be the (working title for now) "CNCF Kubernetes AI Conformance" specification and support for the creation of a suite of tests to demonstrate conformance.

#### Cross-cutting and Externally Facing Processes

- The Working Group will consider its primary problem-solving objective complete upon the successful definition and adoption of a stable (working title for now) "CNCF Kubernetes AI Conformance" specification.
- The first Version will not include any tests in the Kubernetes Codebase and will be a self assessment questionnaire
- Once the initial conformance is established and widely recognized, the ongoing maintenance and evolution of the conformance will be evaluated. Including determining a suitable SIG for Code Organization, Ownership and Architecture and support to create the initial suite of tests.
After the Ownership is clarified and handed over, the WG dissolves.

### Out of scope

This WG is not responsible for maintenance and evolution of the conformance program, which will be run by the CNCF. Some aspects may be owned by existing SIGs (e.g. Architecture, Testing) alongside the CNCF, such as Kubernetes specific definitions, testing and tooling, and the pre-existing Kubernetes conformance definition and tests.

## Roles and Organization Management

This WG follows adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

### Additional responsibilities of Chairs

- The Working Group will have designated Chair(s) responsible for guiding discussions and ensuring progress.
- A note-taker will be assigned for each meeting, and active participation from all contributors will be encouraged.
- Agendas and meeting notes will be publicly accessible.

### Stakeholder SIGs

- SIG Architecture (Sponsoring SIG)
- SIG Testing

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
