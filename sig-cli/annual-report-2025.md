# 2025 Annual Report: SIG CLI

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

   - KEP 3104 (kuberc) introduced as Alpha in v1.33, graduated from Alpha to Beta in the v1.36 release of kubectl, with some quality of life improvements from user feedback, including view/set commands, and the credential plugin allowlist
   - KEP 1441 (debug) began process for deprecating the legacy profile by making "general" the default profile and adding warnings if people are using legacy. Planned deprecation is release v1.39
   - KEP 859 (KUBECTL_COMMAND_HEADERS) promoted to stable in in release v1.35
   - KEP 5295 (kyaml) introduced in v1.34 as alpha and promoted to beta in v1.35
   - KEP 2590 (subresource) promoted to stable in v1.33
   - Deprecated kubeconfig's preferences to make way for KEP 3104 (kuberc)
   - kui and krm-functions retired due to lack of support

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

SIG CLI is seeking help to remove the requirement for the tar binary for kubectl cp. Currently the command requires that the tar binary be available on both the pod and the machine the command is run from. This will likely require help from SIG Node and the various CRI maintainers.

Kustomize only has one active maintainer.

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

   - [KubeCon EU 2025 - What's New With Kubectl and Kustomize … and How You Can Help!](https://www.youtube.com/watch?v=KQBz7nwWxUE)
   - [KubeCon NA 2025 - What's New With Kubectl and Kustomize … and How You Can Help!](https://www.youtube.com/watch?v=2EGxV-3fTwI)

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Beta
    - [3104 - Introduce kuberc](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3104-introduce-kuberc) - 1.34
    - [5295 - KYAML](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/5295-kyaml) - v1.35

  - Stable
    - [2590 - Kubectl Subresource Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2590-kubectl-subresource) - v1.33
    - [859 - Kubectl Commands In Headers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/859-kubectl-headers) - v1.35

## [Subprojects](https://git.k8s.io/community/sig-cli#subprojects)


**Retired in 2025:**
  - kui
  - krm-functions
**Continuing:**
  - cli-experimental
  - cli-sdk
  - cli-utils
  - krew
  - krew-index
  - kubectl
  - kubectl-validate
  - kustomize

## [Working groups](https://git.k8s.io/community/sig-cli#working-groups)

**New in 2025:**
 - AI Integration
 - Node Lifecycle

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cli/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cli/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
