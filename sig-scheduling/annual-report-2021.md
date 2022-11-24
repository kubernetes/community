# 2021 Annual Report: SIG Scheduling

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Efficient re-queueing of pods, significantly cutting the number of failed scheduling cycles
   - Improvements to preemption performance
   - Simplified plugin configuration in component config
   - Scheduler simulator: https://github.com/kubernetes-sigs/kube-scheduler-simulator

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Performance improvements and benchmarking
   - Code refactorings and cleanups
   - Enhancements to node resource -based scoring (see [101946](https://github.com/kubernetes/kubernetes/pull/101946) and [101822](https://github.com/kubernetes/kubernetes/pull/101822))


3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [2249 - Multi-scheduling Profiles](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1451-multi-scheduling-profiles) - 1.22
     - [1845 - Prioritization on Volume Capacity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1845-prioritization-on-volume-capacity) - 1.22
   - Beta
     - [2249 - Namespace Selector for Pod Affinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2249-pod-affinity-namespace-selector) - 1.22
     - [1923 - Prefer Nominated Node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1923-prefer-nominated-node) - 1.22   
     - [2458 - Resource Fit Scoring Strategy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2458-node-resource-score-strategy) - 1.22
     - [2891 - Simplified Scheduler Config](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/2891-simplified-config/kep.yaml) - 1.22
     - [785 - Scheduler Component Config API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/785-scheduler-component-config-api) - 1.23
     - [2926 - Job Mutable Scheduling Directives](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2926-job-mutable-scheduling-directives) - 1.23
   - Alpha
     - None
   - Pre-alpha
     - None

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - https://github.com/kubernetes-sigs/kube-scheduler-simulator is short of reviewers/owners.

2. What metrics/community health stats does your group care about and/or measure?

   - Diversity
   - Number of contributors
   - Meetings attendance
   

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - It looks up-to-date 

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Nothing special.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count: 2529
- Primary mailing list member count: 586
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 5
- Unique reviewers for SIG-owned packages: 6 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 4 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [kube scheduler simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator)
-

Retired in 2021:
- None
-

Continuing:
- [Descheduler](https://github.com/kubernetes-sigs/descheduler)
- [Scheduler Plugins](https://github.com/kubernetes-sigs/scheduler-plugins)
- [Cluster Capacity](https://github.com/kubernetes-sigs/cluster-capacity)
- [Poseidon](https://github.com/kubernetes-sigs/poseidon)
- [KubeBatch](https://github.com/kubernetes-sigs/kube-batch)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [Structured Logging](https://github.com/kubernetes/community/tree/master/wg-structured-logging)


Retired in 2021:
- None
-

Continuing:
- [Policy](https://github.com/kubernetes/community/tree/master/wg-policy)
- [Multitenancy](https://github.com/kubernetes/community/tree/master/wg-multitenancy)

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [2021 KubeCon EU: SIG-Scheduling Intro & Deep Dive](https://sched.co/iE7P)
      - [2021 KubeCon NA: SIG-Scheduling Intro & Deep Dive](https://sched.co/lV8m)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scheduling/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scheduling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

