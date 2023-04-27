# 2022 Annual Report: SIG Apps

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - We're working closely with [WG Batch](https://git.k8s.io/community/wg-batch/README.md) towards improving batch workloads for use in HPC, AI/ML and data analytics workflows.
   - [Mentoring cohort](https://github.com/kubernetes/community/issues/6665) which allowed us to grow our reviewers' list.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - None.

3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2804 - Consolidate Workload controllers life cycle status](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2804-consolidate-workload-controllers-status) - v1.24
    - [961 - Implement maxUnavailable for StatefulSets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/961-maxunavailable-for-statefulset) - v1.24
    - [3329 - Retriable and non-retriable Pod failures for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3329-retriable-and-non-retriable-failures) - v1.25
    - [3017 - Pod Healthy Policy for PDB](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3017-pod-healthy-policy-for-pdb) - v1.26
  - beta:
    - [2879 - Track ready Pods in Job status](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2879-ready-pods-job-status) - v1.24
    - [3140 - TimeZone support in CronJob](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3140-TimeZone-support-in-CronJob) - v1.25
    - [2307 - Job tracking without lingering Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2307-job-tracking-without-lingering-pods) - v1.26
    - [3329 - Retriable and non-retriable Pod failures for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3329-retriable-and-non-retriable-failures) - v1.26
  - stable:
    - [2214 - Indexed Job](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2214-indexed-job) - v1.24
    - [2232 - Suspend Job](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2232-suspend-jobs) - v1.24
    - [1591 - Allow DaemonSets to surge during update like Deployments](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/1591-daemonset-surge/kep.yaml) - v1.25
    - [2599 - minReadySeconds for StatefulSets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2599-minreadyseconds-for-statefulsets) - v1.25

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - The [mentoring cohort](https://github.com/kubernetes/community/issues/6665) we've run in 2022 attracted a few new members, but we still need more.

2. What metrics/community health stats does your group care about and/or measure?

   - [Open untriaged issues and PRs](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Aneeds-triage+label%3Asig%2Fapps)

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - It's up-to-date.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - No additional training than the recent mentoring cohort is planned, yet.

5. Does the group have contributors from multiple companies/affiliations?

   - [13 companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Apps&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1609455600000&to=1639350000000) contributed code in 2022.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Increase the test coverage for all controllers, is probably the best way to help and start contributing to sig-apps owned controllers.
     This also nicely fits into the on-going effort of [increasing the project reliability](https://groups.google.com/g/kubernetes-sig-release/c/P5gFtnjXDqI/m/lVVZ40w2AAAJ).

## Membership

- Primary slack channel member count: 3224
- Primary mailing list member count: 710
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: 71
- Unique approvers for SIG-owned packages: 35

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-apps#subprojects)

**Continuing:**

  - application
  - examples
  - execution-hook
  - kompose
  - workloads-api

## [Working groups](https://git.k8s.io/community/sig-apps#working-groups)

**New in 2022:**

 - Batch

**Continuing:**

 - Data Protection

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon EU 2022 SIG Apps Updates](https://www.youtube.com/watch?v=JAUIUNhYZWg)
      - [KubeCon NA 2022 SIG Apps Updates](https://www.youtube.com/watch?v=UliDcWor_d0)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-apps/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-apps/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
