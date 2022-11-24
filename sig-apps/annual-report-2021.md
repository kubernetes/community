# 2021 Annual Report: SIG Apps

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Major improvements in the Job API (see below job-related KEPs).
   - Stability and availability improvements across several controllers (see below KEPs for DaemonSets and StatefulSets).
   - Leadership changes: Maciej joins leadership, Matt and Adnan moving to emeritus.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Conformance testing promotions.

3. KEP work in 2021 (1.21, 1.22, 1.23):

   - Stable
     - [19 - CronJob to Stable](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/19-Graduate-CronJob-to-Stable/README.md) - 1.21
     - [85 - PodDisruptionBudget to GA](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/85-Graduate-PDB-to-Stable/README.md) - 1.22
     - [592 - TTL After Finished](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/592-ttl-after-finish/README.md) - 1.23
   - Beta
     - [2185 - Random Pod Selection on ReplicaSet Downscale](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2185-random-pod-select-on-replicaset-downscale/README.md) - 1.22
     - [1591 - Allow DaemonSets to surge during update like Deployments](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/1591-daemonset-surge/README.md) - 1.22
     - [2214 - Indexed Job](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2214-indexed-job/README.md) - 1.22
     - [2232 - Suspend Job](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2232-suspend-jobs/README.md) - 1.22
     - [2255 - ReplicaSet Pod Deletion Cost](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2255-pod-cost/README.md) - 1.22
     - [2307 - Job tracking without lingering Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2926-job-mutable-scheduling-directives/README.md) - 1.23
     - [2599 - minReadySeconds for StatefulSets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2599-minreadyseconds-for-statefulsets/README.md) - 1.23
     - [2926 - Mutable Node Scheduling Directives for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2599-minreadyseconds-for-statefulsets/README.md) - 1.23
   - Alpha
     - [2185 - Random Pod Selection on ReplicaSet Downscale](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2185-random-pod-select-on-replicaset-downscale/README.md) - 1.21
     - [1591 - Allow DaemonSets to surge during update like Deployments](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/1591-daemonset-surge/README.md) - 1.21
     - [2214 - Indexed Job](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2214-indexed-job/README.md) - 1.21
     - [2232 - Suspend Job](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2232-suspend-jobs/README.md) - 1.21
     - [2255 - ReplicaSet Pod Deletion Cost](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2255-pod-cost/README.md) - 1.21
     - [2307 - Job tracking without lingering Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2307-job-tracking-without-lingering-pods/README.md) - 1.22
     - [2599 - minReadySeconds for StatefulSets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2599-minreadyseconds-for-statefulsets/README.md) - 1.22
     - [1847 - Auto delete PVCs created by StatefulSet](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/1847-autoremove-statefulset-pvcs/README.md) - 1.23
     - [2879 - Track ready Pods in Job status](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/2879-ready-pods-job-status/README.md) - 1.23

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Growing additional reviewers and approvers, also see no. 4 below.

2. What metrics/community health stats does your group care about and/or measure?

   - Open untriaged issues and PRs

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - It's up-to-date.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We’re working with SIG-CLI on starting a monthly review club, details to be announced soon.

5. Does the group have contributors from multiple companies/affiliations?

   - [11 companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Apps&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1609455600000&to=1639350000000) contributed code in 2021.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Increase the test coverage for all controllers.

## Membership

- Primary slack channel member count: 2900
- Primary mailing list member count: 674
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: 23
- Unique approvers for SIG-owned packages: 31

Include any other ways you measure group membership

## Subprojects

New in 2021:
- None


Retired in 2021:
- None

Continuing:
- [application](https://git.k8s.io/community/sig-apps#application)
- [examples](https://git.k8s.io/community/sig-apps#examples)
- [execution-hook](https://git.k8s.io/community/sig-apps#execution-hook)
- [kompose](https://git.k8s.io/community/sig-apps#kompose)

## Working groups

New in 2021:
- None

Retired in 2021:
- None

Continuing:
- [Data Protection](https://git.k8s.io/community/wg-data-protection/) ([2021 report](https://github.com/kubernetes/community/blob/master/wg-data-protection/annual-report-2021.md))

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon NA 2021 Sig Apps updates](https://youtu.be/ZvFYvYiMeTs)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-apps/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-apps/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

