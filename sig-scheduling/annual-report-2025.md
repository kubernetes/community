# 2025 Annual Report: SIG Scheduling

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- Started Workload Aware Scheduling (aka WAS) initiative
  - First implementation of gang scheduling
- Active support for DRA
- Improvements to scheduling performance
- New Chairs and Tech leads

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- [scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins/)
- [kwok](https://github.com/kubernetes-sigs/kwok)
- [kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator/)
- [kube-scheduler-wasm-extension](https://github.com/kubernetes-sigs/kube-scheduler-wasm-extension)

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

- [Kubecon NA maintainer session](https://sched.co/27NlU)
- [Kubecon EU maintainer session](https://sched.co/1td1i)


4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Alpha
    - [4671 - Gang Scheduling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4671-gang-scheduling) - v1.35
    - [4815 - DRA Partitionable Devices](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4815-dra-partitionable-devices) - v1.33
    - [5004 - DRA Extended Resource](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5004-dra-extended-resource) - v1.34
    - [5007 - DRA Device Binding Conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5007-device-attach-before-pod-scheduled) - v1.34
    - [5055 - DRA: device taints and tolerations](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5055-dra-device-taints-and-tolerations) - v1.33
    - [5075 - DRA Consumable Capacity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5075-dra-consumable-capacity) - v1.34
    - [5471 - Extended Toleration Operators for Threshold-Based Placement](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5471-enable-sla-based-scheduling) - v1.35

  - Beta
    - [4816 - DRA Prioritized List](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4816-dra-prioritized-list) - v1.34
    - [4832 - Asynchronous Preemption](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4832-async-preemption) - v1.33
    - [5142 - Pop pod from backoffQ when activeQ is empty](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5142-pop-backoffq-when-activeq-empty) - v1.33
    - [5229 - Asynchronous API calls during scheduling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5229-asynchronous-api-calls-during-scheduling) - v1.34
    - [5234 - DRA ResourceSlice Mixins](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5234-dra-resourceslice-mixins) - v1.35
    - [5278 - Nominated node name for an expected pod placement](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5278-nominated-node-name-for-expectation) - v1.35
    - [5501 - Reflect PreEnqueue rejections in Pod status](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5501-reflect-preenqueue-rejections-in-pod-status) - v1.35
    - [5598 - Opportunistic batching](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/5598-opportunistic-batching) - v1.35

  - Stable
    - [3094 - Take taints/tolerations into consideration when calculating PodTopologySpread skew](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3094-pod-topology-spread-considering-taints) - 1.33
    - [3633 - Introduce MatchLabelKeys and MismatchLabelKeys to PodAffinity and PodAntiAffinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3633-matchlabelkeys-to-podaffinity) - v1.33
    - [3902 - Decouple TaintManager from NodeLifeCycleController](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3902-decoupled-taint-manager) - v1.34
    - [4247 - Per-plugin callback functions for efficient requeueing in the scheduling queue](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/4247-queueinghint) - v1.34

## [Subprojects](https://git.k8s.io/community/sig-scheduling#subprojects)


**New in 2025:**
  - dra-driver-topology
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

**New in 2025:**
 - Checkpoint Restore
 - Node Lifecycle
**Retired in 2025:**
 - Policy
**Continuing:**
 - Batch
 - Device Management
 - Serving
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scheduling/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scheduling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
