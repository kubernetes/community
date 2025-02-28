# 2024 Annual Report: SIG Testing

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?
- (Together with SIG K8s Infra) Migration of the Prow control plane (and related resources) from Google-owned infra to community-owned infra, including:
  - Notifying the community about the migration (and work needed beforehand)
  - Identifying and migrating jobs blocking the migration
  - Adding new logs/image locations and updating references to them
  - Updating tools (Kettle, TestGrid, etc.) to use the new Prow instance
- Moved Prow code from k8s/test-infra into its own repository at [k8s-sigs/prow](https://github.com/kubernetes-sigs/prow)
<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?
Yup! SIG Testing covers a large variety of subprojects and is generally in need of contributors and maintainers.
- Important:
  - Prow: We still need more contributors and maintainers for Prow, especially folks who can focus on the Kubernetes project's use of Prow in particular.
  - Hydrophone: Feedback and help enabling adoption of new conformance testing tools for end users.
- Useful:
  - TestGrid
  - Boskos
  - Kettle

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?
- KubeCon NA 2024:
  - "Achieving and Maintaining a Healthy CI with Zero Test Flakes"
    - Event: https://sched.co/1hoxc
    - Recording: https://youtu.be/hl3jjCTTL50 
- Contributor Summit NA 2024:
  - "Unified framework for unit integration and E2E testing"
    - Event: https://sched.co/1nSjo
    - Recording: https://youtu.be/VCG559w9gzo
  - TestGrid: Visualizing Test Results
    - Event: https://sched.co/1nShv
    - Slides:
      https://docs.google.com/presentation/d/13yD2AixxJEscl0Fxe9Sw79V5WTJh7QCi/edit?usp=sharing

4. KEP work in 2024 (v1.30, v1.31, v1.32): N/A

## [Subprojects](https://git.k8s.io/community/sig-testing#subprojects)

**New in 2024:**
  - hydrophone
  - testgrid

**Continuing:**
  - Cloud Provider for KIND
  - boskos
  - e2e-framework
  - kind
  - kubetest2
  - prow
  - sig-testing
  - test-infra
  - testing-commons

## [Working groups](https://git.k8s.io/community/sig-testing#working-groups)

**Continuing:**
 - LTS

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-testing/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-testing/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
