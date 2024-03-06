# 2023 Annual Report: SIG Scheduling

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [3280 - Guarantee PodDisruptionBudget When Preemption Happens](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3280-guarantee-pdb-when-preemption-happens) - v1.27
    - [3633 - Introduce MatchLabelKeys and MismatchLabelKeys to PodAffinity and PodAntiAffinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3633-matchlabelkeys-to-podaffinity) - v1.29

  - Beta
    - [3521 - Pod Scheduling Readiness](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3521-pod-scheduling-readiness) - v1.27
    - [3838 - Pod Mutable Scheduling Directives](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3838-pod-mutable-scheduling-directives) - 1.27
    - [3902 - Decouple TaintManager from NodeLifeCycleController](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3902-decoupled-taint-manager) - v1.29
    - [4247 - Per-plugin callback functions for efficient requeueing in the scheduling queue](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4247-queueinghint) - v1.28

  - Stable
    - [2926 - Mutable Node Scheduling Directives for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2926-job-mutable-scheduling-directives) - v1.27
    - [3243 - Respect PodTopologySpread after rolling upgrades](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3243-respect-pod-topology-spread-after-rolling-upgrades) - v1.29

## [Subprojects](https://git.k8s.io/community/sig-scheduling#subprojects)


**New in 2023:**
  - [kube-scheduler-wasm-extension](https://git.k8s.io/community/<no value>#kube-scheduler-wasm-extension)
**Retired in 2023:**
  - kube-batch
**Continuing:**
  - cluster-capacity
  - descheduler
  - kube-scheduler-simulator
  - kueue
  - kwok
  - scheduler
  - scheduler-plugins

## [Working groups](https://git.k8s.io/community/sig-scheduling#working-groups)

**Retired in 2023:**
 - Multitenancy
**Continuing:**
 - Batch
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scheduling/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scheduling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
