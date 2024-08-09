# 2023 Annual Report: SIG Windows

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  - Windows Operational Readiness [v0.1.0](https://github.com/kubernetes-sigs/windows-operational-readiness/releases/tag/0.1.0)
    was released
  - [Windows support for InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/112599)
  - [CRI only metrics](https://github.com/kubernetes/kubernetes/pull/116968)
  - Windows-service-proxy was highlighted as a working example of building your own kube-proxy out of tree using KPNG
  - Kicked off work on [Image Pull Per Runtime Class](https://github.com/kubernetes/enhancements/pull/4217/) feature
    which spans kubelet and CRI
  - Leadership changes
    - Aravindh Puthiyaparambil became co-chair while Mark Rossetti also took on the mantle of tech lead
    - Jay Vyas stepped down as tech lead and was replaced by Amim Knabben

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   No

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

  - [What's new with SIG Windows, Kubecon EU](https://sched.co/1HyTs)
  - [What's new with SIG Windows, Kubecon NA](https://sched.co/1R2mL)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [2258 - Node log query](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2258-node-log-query) - v1.27
    - [4216: Image pull per runtime class](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/4216-image-pull-per-runtime-class/) - v1.29

## [Subprojects](https://git.k8s.io/community/sig-windows#subprojects)

**New in 2023:**
  - [windows-service-proxy](https://git.k8s.io/community/<no value>#windows-service-proxy)

**Continuing:**
  - windows-gmsa
  - windows-operational-readiness
  - windows-samples
  - windows-testing
  - windows-tools

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
