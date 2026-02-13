# 2025 Annual Report: WG AI Conformance

## Current initiatives and Project Health


1. What work did the WG do this year that should be highlighted?

   *   **Formation of the Working Group:** The WG AI Conformance was officially established in 2025 to define a standardized set of capabilities, APIs, and configurations for running AI/ML workloads on Kubernetes. The charter was ratified and the group began bi-weekly meetings.

   *   **Launch of the Kubernetes AI Conformance Program:** The group launched the official AI conformance program. This included the creation of the `cncf/k8s-ai-conformance` repository to host the program's resources and certification process, and an announcement at KubeCon + CloudNativeCon North America 2025.

   *   **Release of Conformance Checklists:** The WG defined and released conformance checklists for multiple Kubernetes versions, providing a clear baseline for AI-ready clusters:
        *   `AI Conformance-1.33.yaml`
        *   `AI Conformance-1.34.yaml`
        *   `AI Conformance-1.35.yaml`

   *   **Established Requirement & Test Processes:**
        *   Transitioned requirement tracking from documents to a structured GitHub Project: [WG AI Conformance Requirements](https://github.com/orgs/kubernetes-sigs/projects/114).
        *   Defined the "Kubernetes AI Conformance Requirement (KAR)" process, mirroring the KEP process, to manage the lifecycle of requirements (SHOULD/MUST).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   *   **Automated Test Implementation:** While the conformance checklists are established, the group is transitioning towards automated verification. We need contributors to help design and implement these automated tests in 2026 to replace the current self-assessment model.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025
    - [SIG Architecture](https://git.k8s.io/community/sig-architecture/)
      - [18 Sep 2025](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit?tab=t.0#bookmark=kix.kq4658qk0xqr)
      - [30 Oct 2025](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit?tab=t.0#bookmark=id.yo05n0lgug01)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-ai-conformance/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
