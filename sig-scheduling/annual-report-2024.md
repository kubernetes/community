# 2024 Annual Report: SIG Scheduling

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- **The performance improvement**
  - Have put a lot of effort into QueueingHint ([KEP-4247](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4247-queueinghint)), which enhances the scheduler's retrying efficiency.
  - Introduced the asynchronous preemption ([KEP-4832](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4832-async-preemption)), which improves the scheduling throughput with the preemption.
  - Improved [the internal scheduling performance test (scheduler-perf)](https://github.com/kubernetes/kubernetes/tree/master/test/integration/scheduler_perf) to cover more scenarios and alert us when the degradation is introduced.
- **DRA**: Improved the DRA scheduling with the structured parameters ([KEP-4381](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4381-dra-structured-parameters#kube-scheduler)).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- [scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- [Kubecon EU 2024: SIG_Scheduling Intro & Deep Dive](https://sched.co/1YhjR)
- [Kubecon NA 2024: SIG_Scheduling Intro & Updates](https://sched.co/1hovV)

4. KEP work in 2024 (v1.30, v1.31, v1.32):

  - Alpha
    - [4816 - DRA Prioritized List](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4816-dra-prioritized-list) - v1.32
    - [4832 - Asynchronous Preemption](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4832-async-preemption) - v1.32

  - Beta
    - [3633 - Introduce MatchLabelKeys and MismatchLabelKeys to PodAffinity and PodAntiAffinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3633-matchlabelkeys-to-podaffinity) - v1.31
    - [4247 - Per-plugin callback functions for efficient requeueing in the scheduling queue](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/4247-queueinghint/README.md) - v1.32

  - Stable
    - [3022 - Tuning the number of domains in PodTopologySpread](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3022-min-domains-in-pod-topology-spread) - v1.30
    - [3521 - Pod Scheduling Readiness](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3521-pod-scheduling-readiness) - v1.30
    - [3838 - Pod Mutable Scheduling Directives](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3838-pod-mutable-scheduling-directives) - v1.30 

## [Subprojects](https://git.k8s.io/community/sig-scheduling#subprojects)


**Continuing:**
  - cluster-capacity
  - descheduler
  - kube-scheduler-simulator
  - kube-scheduler-wasm-extension
  - kueue
  - kwok
  - scheduler
  - scheduler-plugins

## [Working groups](https://git.k8s.io/community/sig-scheduling#working-groups)

**New in 2024:**
 - Device Management
 - Serving
**Continuing:**
 - Batch
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scheduling/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scheduling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
