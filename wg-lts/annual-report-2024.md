# 2024 Annual Report: WG LTS

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

   * Kubernetes Upgrade Survey
     * [Kubernetes Upgrade Survey responses](https://docs.google.com/spreadsheets/d/1VHkIDvqu6OT05sAFbTeCvVK42mFIQFJ-xdAyO7dl4mU/edit?gid=1310802220#gid=1310802220)
     * [Analysis](https://docs.google.com/presentation/d/1HeuZ_3R_U2FmwhMTp_vnAf4UFXvSoMhHZfcUKGyiz2E/edit)
   * Kubecon presentations / discussions
     * Kubernetes Contributor Summit EU 2024 - [Viability of Kubernetes community LTS](https://www.youtube.com/watch?v=ktOszIqEJJQ&list=PL69nYSiGNLP1TJ5uTeTRtjY3SBoNoEpz7&index=10)
   * [Success stories / cautionary tales from upgrades](https://docs.google.com/document/d/1HbNDKtl3LdcJsCuPHGsSCvg3e7ZDzu-ca1PSm16sZWU/edit?tab=t.0#heading=h.if8xiqpwaclb)
   * Compatibility / Emulation version
     * [KEP-4330](https://github.com/kubernetes/enhancements/issues/4330)
   * Kubernetes regressions / backports
     * [Strengthened Kubernetes backport requirements](https://github.com/kubernetes/community/issues/7634)
     * Tracked data on [Kubernetes regressions](https://docs.google.com/spreadsheets/d/1LbGKBC4D2sLkcmzY9qDx9u-1D9TKC_ZrM8iA1eHW4Hs/edit#gid=1283859152)
   * Ecosystem support / compatibility
     * Monitored and engaged in discussions around containerd 2.0 / 1.7 support
       * https://github.com/containerd/containerd/pull/9833
       * https://github.com/containerd/containerd/pull/9879 
     * Engaged with the Go project on runtime compatibility
       * 1.21: added Go compatibility version support for 2 years
         * https://go.dev/doc/godebug
         * GODEBUG settings added for compatibility will be maintained for a minimum of two years (four Go releases).
       * 1.23: separated compatibility defaults from language version
         * https://github.com/golang/go/issues/65573

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Not currently

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed
- [ ] Updates provided to sponsoring SIGs in 2024
     - [SIG Architecture](https://git.k8s.io/community/sig-architecture/)
        - [Success stories / cautionary tales](https://groups.google.com/g/kubernetes-sig-architecture/c/j1O3qy1iFI0) broadcast
        - [Compatibility mode](https://groups.google.com/g/kubernetes-sig-architecture/c/TVecQbtYA-s) broadcast
     - [SIG Cluster Lifecycle](https://git.k8s.io/community/sig-cluster-lifecycle/)
        - no specific updates
     - [SIG K8s Infra](https://git.k8s.io/community/sig-k8s-infra/)
        - no specific updates
     - [SIG Release](https://git.k8s.io/community/sig-release/)
        - no specific updates
     - [SIG Security](https://git.k8s.io/community/sig-security/)
        - no specific updates
     - [SIG Testing](https://git.k8s.io/community/sig-testing/)
        - no specific updates

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-lts/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
