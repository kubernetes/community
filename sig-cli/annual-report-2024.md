# 2024 Annual Report: SIG CLI

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

- Custom profiling support in kubectl debug is promoted to GA
- Interactive delete in kubectl delete is promoted to GA
- Transition from SPDY to Websockets is promoted to beta
- kuberc alpha phase implementation is started and reviewed extensively
- Removing kustomize from kubectl is discussed and rejected
- Krew has 286 plugins
- Kustomize new releases (latest one v5.6.0 as of now) 

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- Kui project seems to be not actively maintained

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

- [KubeCon Europe 2024](https://www.youtube.com/watch?v=LjXZjt_yOJ8)
- [KubeCon North America 2024](https://www.youtube.com/watch?v=EL2mx5Ukho8)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Stable
    - [1441 - kubectl debug](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/1441-kubectl-debug) - v1.30
    - [3638 - Improve kubectl plugin resolution for non-shadowing subcommands](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3638-kubectl-plugin-subcommands) - v1.30
    - [3805 - Kubectl Server-Side Apply by default](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3805-ssa-default) - v1.32
    - [3895 - Interactive(-i) flag to kubectl delete for user confirmation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3895-kubectl-delete-interactivity) - v1.30
    - [4292 - Custom profiling support in kubectl debug command](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/4292-kubectl-debug-custom-profile) - 1.32

## [Subprojects](https://git.k8s.io/community/sig-cli#subprojects)


**Continuing:**
  - cli-experimental
  - cli-sdk
  - cli-utils
  - krew
  - krew-index
  - krm-functions
  - kubectl
  - kubectl-validate
  - kui
  - kustomize

## [Working groups](https://git.k8s.io/community/sig-cli#working-groups)


## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cli/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cli/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
