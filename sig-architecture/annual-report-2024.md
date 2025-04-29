# 2024 Annual Report: SIG Architecture

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->
- At Kubecon NA, we merged the last outstanding conformance test to reach "100% conformance tested" milestone (see [issue](https://github.com/kubernetes/enhancements/issues/4945)). With this we fully paid down the debt for older APIs which did not have conformance tests.
- Go workspaces for k/k (see [issue](https://github.com/kubernetes/enhancements/issues/4402))
- Launched WG Device Management (see [PR](https://github.com/kubernetes/community/pull/7805))
- Launched WG Serving (see [PR](https://github.com/kubernetes/community/pull/7823))
- Add apisnoop subproject (see [Issue](https://github.com/kubernetes/org/issues/4705))

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

The team needs new contributors. While we handle regular release tasks well, we could improve areas like simplifying the KEP template. Fresh perspectives in sig-architecture meetings would help challenge the status quo. We also need long-term contributors for future leadership transitions. Most subprojects (e.g., api-reviews, PRRs) are self-sustaining, but we haven't tackled new conformance challenges (e.g., mandatory RBAC). Maintaining the release cadence remains our priority.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

Yes we had SIG-Architecture related talks in:
- [kubecon EU 2024](https://www.youtube.com/watch?v=8YaKHvoVZy4)
- [kubecon NA 2024](https://www.youtube.com/watch?v=8YaKHvoVZy4)

Thanks to John Belamaric and David Eads!

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

- stable
  - [Go workspaces for k/k](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/4402-go-workspaces)

<!-- 

  - Beta
    - [4330 - Compatibility Versions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/4330-compatibility-versions) - v1.32

  - Stable
    - [4402 - Go workspaces for k/k](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/4402-go-workspaces) - v1.30 -->

## [Subprojects](https://git.k8s.io/community/sig-architecture#subprojects)


**New in 2024:**
  - apisnoop
  - wg-device-management
  - wg-serving
**Continuing:**
  - architecture-and-api-governance
  - code-organization
  - conformance-definition
  - enhancements
  - production-readiness

## [Working groups](https://git.k8s.io/community/sig-architecture#working-groups)

**New in 2024:**
 - Device Management
 - Serving
**Retired in 2024:**
 - API Expression
**Continuing:**
 - LTS
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-architecture/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-architecture/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
