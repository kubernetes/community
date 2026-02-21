# 2025 Annual Report: SIG Apps

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  - Modernized and revitalized the [examples](https://github.com/kubernetes/examples/)
    repository by archiving legacy/unmaintained content and introducing high-impact AI/ML reference manifests.
  - Started [agent-sandbox](https://github.com/kubernetes-sigs/agent-sandbox) subproject.
  - Sponsored the creation of several new [working groups](#working-groups).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  None.

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

  Yes. SIG-Apps provided the following updates:
  - [KubeCon EU 2025](https://www.youtube.com/watch?v=KlsxQMfdKLw)
  - [KubeCon NA 2025](https://www.youtube.com/watch?v=9xpp3wvoSDg)

4. KEP work in 2025 (v1.33, v1.34, v1.35):

  - Alpha
    - [5440 - Mutable Pod Resources for Suspended Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/5440-mutable-job-pod-resource-updates) - 1.35

  - Beta
    - [961 - Implement maxUnavailable for StatefulSets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/961-maxunavailable-for-statefulset) - v1.35
    - [3973 - Consider Terminating Pods in Deployments](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3973-consider-terminating-pods-deployment) - v1.35

  - Stable
    - [3850 - Backoff Limits Per Index For Indexed Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3850-backoff-limits-per-index-for-indexed-jobs) - v1.33
    - [3939 - Allow Replacement of Pods in a Job when fully terminating](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3939-allow-replacement-when-fully-terminated) - v1.34
    - [3998 - Job success/completion policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3998-job-success-completion-policy) - v1.33
    - [4368 - Job API managed-by label](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/4368-support-managed-by-for-batch-jobs) - v1.35

## [Subprojects](https://git.k8s.io/community/sig-apps#subprojects)

**New in 2025:**
  - agent-sandbox

**Continuing:**
  - application
  - examples
  - execution-hook
  - kjob
  - kompose
  - workloads-api

## [Working groups](https://git.k8s.io/community/sig-apps#working-groups)

**New in 2025:**
 - AI Integration
 - Checkpoint Restore
 - Node Lifecycle

**Continuing:**
 - Batch
 - Data Protection
 - Serving

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-apps/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-apps/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
