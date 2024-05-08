# 2023 Annual Report: SIG CLI

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

    - Interactive delete in kubectl delete
    - Custom profiling support in kubectl debug
    - Transition from SPDY to Websockets
    - Plugin resolution for non-shadowing subcommands in kubectl promoted to beta
    - New pruning design for kubectl apply --prune, aka ApplySet
    - Aggregated Discovery promoted to GA
    - kubectl debug promoted to GA
    - New minor release for kui
    - Krew has 242 plugins

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

    - Lost 2 Kustomize primary maintainers and looking for help in 2024

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

  - [KubeCon Europe 2023](https://www.youtube.com/watch?v=X-XDr8XhHHU)
  - [KubeCon North America 2023](https://www.youtube.com/watch?v=RggqaCSdOGA)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [3659 - KEP Template](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3659-kubectl-apply-prune) - v1.27
    - [3805 - Kubectl Server-Side Apply by default](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3805-ssa-default) - v1.27
    - [4006 - Transition from SPDY to Websockets](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/4006-transition-spdy-to-websockets/kep.yaml) - 1.28

  - Beta
    - [2590 - Kubectl Subresource Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2590-kubectl-subresource) - v1.27
    - [3638 - Improve kubectl plugin resolution for non-shadowing subcommands](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3638-kubectl-plugin-subcommands) - v1.29
    - [3895 - Interactive(-i) flag to kubectl delete for user confirmation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3895-kubectl-delete-interactivity) - v1.29

  - Stable
    - [1440 - Kubectl events](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/1440-kubectl-events) - v1.28
    - [2227 - kubectl default container](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2227-kubectl-default-container) - v1.27
    - [2906 - Kustomize Function Catalog](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2906-kustomize-function-catalog) - v1.27
    - [3104 - Introduce kuberc](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3104-introduce-kuberc) - v1.28
    - [3515 - Kubectl Explain OpenAPIv3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3515-kubectl-explain-openapiv3) - 1.29

## [Subprojects](https://git.k8s.io/community/sig-cli#subprojects)


**New in 2023:**
  - [kubectl-validate](https://git.k8s.io/community/sig-cli#kubectl-validate)

**Continuing:**
  - cli-experimental
  - cli-sdk
  - cli-utils
  - krew
  - krew-index
  - krm-functions
  - kubectl
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
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cli/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cli/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
