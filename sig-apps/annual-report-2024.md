# 2024 Annual Report: SIG Apps

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  - Mailing list has been migrated to kubernetes managed account.
  - Sponsored the creation of [WG Serving](/wg-serving).
  - We continued cleaning up the backlog of started enhancements.


2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

None.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

  Yes. SIG-Apps provided the following updates:
  - [KubeCon EU 2024](https://www.youtube.com/watch?v=bE8XpaJwq-Q)
  - [KubeCon NA 2024](https://www.youtube.com/watch?v=NWZhAs69heA)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
  - Alpha
    - [4443 - More granular Job failure reasons for PodFailurePolicy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4443-configurable-pod-failure-policy-reasons) - v1.31

  - Beta
    - [3998 - Job success/completion policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3998-job-success-completion-policy) - v1.31
    - [4368 - Job API managed-by label](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4368-support-managed-by-for-batch-jobs) - v1.32

  - Stable
    - [1847 - Auto delete PVCs created by StatefulSet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/1847-autoremove-statefulset-pvcs) - v1.32
    - [2185 - Random Pod Selection on ReplicaSet Downscale](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2185-random-pod-select-on-replicaset-downscale) - v1.31
    - [3017 - Pod Healthy Policy for PDB](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3017-pod-healthy-policy-for-pdb) - v1.31
    - [3329 - Retriable and non-retriable Pod failures for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3329-retriable-and-non-retriable-failures) - v1.31
    - [3335 - StatefulSet Slice](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3335-statefulset-slice) - v1.31
    - [3715 - Elastic Indexed Job](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3715-elastic-indexed-job) - v1.31
    - [4017 - Pod Index Label](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4017-pod-index-label) - v1.32
    - [4026 - Add job creation timestamp to job annotations](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4026-crojob-scheduled-timestamp-annotation) - v1.32

## [Subprojects](https://git.k8s.io/community/sig-apps#subprojects)

**New in 2024:**
  - kjob
**Continuing:**
  - application
  - examples
  - execution-hook
  - kompose
  - workloads-api

## [Working groups](https://git.k8s.io/community/sig-apps#working-groups)

**New in 2024:**
 - Serving
**Continuing:**
 - Batch
 - Data Protection

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-apps/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-apps/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
