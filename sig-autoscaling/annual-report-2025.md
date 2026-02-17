# 2025 Annual Report: SIG Autoscaling

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

- 3 minor releases of VPA: [1.3.0](https://github.com/kubernetes/autoscaler/releases/tag/vertical-pod-autoscaler-1.3.0), [1.4.0](https://github.com/kubernetes/autoscaler/releases/tag/vertical-pod-autoscaler-1.4.0), [1.5.0](https://github.com/kubernetes/autoscaler/releases/tag/vertical-pod-autoscaler-1.5.0)
  - In Place Pod Resizing promoted to Beta (`InPlaceOrRecreate` feature) (@adrianmoisey)
- VPA Enhancement Proposals
  - AEP-7862: [AEP-7862: CPU Startup Boost](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/7862-cpu-startup-boost) (@kamarabbas99)
  - AEP-8026 [AEP-8026: Allow per-VPA component configuration parameters](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/8026-per-vpa-component-configuration) (@omerap12)
  - AEP-8818: [InPlace Update Mode work begun](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/8818-in-place-only) (@omerap12)
- 2 minor releases of Cluster Autoscaler: [1.33.0](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.33.0), [1.34.0](https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.33.0)
- Cluster Autoscaler Enhancement Proposals
  - [Granular Resource Limits (CapacityQuota API)](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/granular-resource-limits.md) @norbertcyran
  - Spare Capacity (CapacityBuffer API)[https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/buffers.md] @jbtk
- 7 minor release of Karpenter: [1.2.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.2.0), [1.3.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.3.0), [1.4.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.4.0), [1.5.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.5.0), [1.6.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.6.0), [1.7.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.7.0), [1.8.0](https://github.com/kubernetes-sigs/karpenter/releases/tag/v1.8.0)
- Karpenter RFCs
  - [Static Capacity](https://github.com/kubernetes-sigs/karpenter/blob/main/designs/static-capacity.md) @sumukha-radhakrishna
  - [NodeRegistrationHealthy Status Condition](https://github.com/kubernetes-sigs/karpenter/pull/1910) @jigisha620
  - [Node Overlay](https://github.com/kubernetes-sigs/karpenter/blob/main/designs/node-overlay.md) @engedaam
  - [Karpenter Integration Test Migration](https://github.com/kubernetes-sigs/karpenter/blob/main/designs/karpenter-integration-testing.md) @engedaam
  - [DRA KWOK Driver](https://github.com/kubernetes-sigs/karpenter/blob/main/designs/karpenter-dra-kwok-driver.md) @alimaazamat
  - [GTE and LTE Operators for Requirements](https://github.com/kubernetes-sigs/karpenter/blob/main/designs/gte-lte-operators.md) @ellistarn
- Healthy Ecosystem
  - Karpenter provider maturity
    - [IBM Cloud Provider](https://github.com/kubernetes-sigs/karpenter-provider-ibm-cloud)
    - [Cluster API Provider](https://github.com/kubernetes-sigs/karpenter-provider-cluster-api)
  - Cluster Autoscaler provider growth
    - New [CoreWeave Provider](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/cloudprovider/coreweave)
    - New [Utho Provider](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/cloudprovider/utho)
- Significant Leadership Updates
  - Maciek Pytel steps down as Chair (Thank you @maciekpytel !)
  - Kuba Tużnik (@towca) promoted to Chair
  - Jack Francis (@jackfrancis) promoted to Lead
  - Ray Wainman served as Lead, now Emeritus (Thank you @raywainman !)
  - Adrian Moisey @adrianmoisey promoted to Lead

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- Omer Aplatony has stepped up considerably to reinforce HPA (thank you @omerap12), in addition to his work maintaining VPA. At least one more regular maintainer is still needed. We are tracking progress here:
  - https://github.com/kubernetes/kubernetes/issues/128948


3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

- [Kubecon NA 2025 SIG Update](https://www.youtube.com/watch?v=aflZ5ccrgnw)

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

- Alpha
  - [5030 - Integrate CSI Volume attach limits with cluster autoscaler](https://github.com/kubernetes/enhancements/tree/master/keps/sig-autoscaling/5030-attach-limit-autoscaler) - v1.35
- Beta
  - [4951 - Configurable tolerance for HPA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-autoscaling/4951-configurable-hpa-tolerance) - v1.35

## [Subprojects](https://git.k8s.io/community/sig-autoscaling#subprojects)


**Continuing:**
  - cluster-autoscaler
  - karpenter
  - horizontal-pod-autoscaler
  - vertical-pod-autoscaler
  - addon-resizer

## [Working groups](https://git.k8s.io/community/sig-autoscaling#working-groups)

**New in 2025:**
 - Node Lifecycle
**Continuing:**
 - Batch
 - Device Management
 - Serving (Retiring as of Feb 2026)

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-autoscaling/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-autoscaling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
