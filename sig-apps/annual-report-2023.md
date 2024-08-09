# 2023 Annual Report: SIG Apps

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  In the past year SIG Apps focused on stabilizing the currently supported workload
  APIs, finishing the work that has already started, and providing review bandwidth
  for the work driven by the WG Batch.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  None.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

  Yes. SIG-Apps provided the following updates:
  - [KubeCon EU 2023](https://youtu.be/OC9egbi8IQw?si=wKvIcvAkGkoeBLya)
  - [KubeCon NA 2023](https://youtu.be/7LNOTuo-bdE?si=REB0-pck2zk23dSx)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Beta
    - [3017 - Pod Healthy Policy for PDB](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3017-pod-healthy-policy-for-pdb) - v1.27
    - [3335 - StatefulSet Slice](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3335-statefulset-slice) - v1.27
    - [3715 - Elastic Indexed Job](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3715-elastic-indexed-job) - v1.27
    - [3850 - Backoff Limits Per Index For Indexed Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3850-backoff-limits-per-index-for-indexed-jobs) - v1.29
    - [3939 - Allow Replacement of Pods in a Job when fully terminating](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3939-allow-replacement-when-fully-terminated) - v1.29
    - [4017 - Pod Index Label](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4017-pod-index-label) - v1.28
    - [4026 - Add job creation timestamp to job annotations](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4026-crojob-scheduled-timestamp-annotation) - v1.28

  - Stable
    - [1847 - Auto delete PVCs created by StatefulSet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/1847-autoremove-statefulset-pvcs) - v1.28
    - [2804 - Consolidate Workload controllers life cycle status](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2804-consolidate-workload-controllers-status) - v1.27
    - [2879 - Track ready Pods in Job status](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2879-ready-pods-job-status) - v1.29
    - [3140 - TimeZone support in CronJob](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3140-TimeZone-support-in-CronJob) - v1.27

## [Subprojects](https://git.k8s.io/community/sig-apps#subprojects)

**Continuing:**
  - application
  - examples
  - execution-hook
  - kompose
  - workloads-api

## [Working groups](https://git.k8s.io/community/sig-apps#working-groups)

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
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-apps/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-apps/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
