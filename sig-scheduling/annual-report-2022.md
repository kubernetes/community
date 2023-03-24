# 2022 Annual Report: SIG Scheduling

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Improvements to PodTopologySpread API: adding more knobs to control spreading behavior (introducing minDomains and matchLabelKey).
   - Adding more hooks to simplify integrations with external schedulers (mutable pod scheduling directives, pod scheduling readiness)


2. What initiatives are you working on that aren't being tracked in KEPs?

   - Performance improvements
   - Code refactorings and cleanups




3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - beta:
    - [3022 - Tuning the number of domains in PodTopologySpread](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3022-min-domains-in-pod-topology-spread) - v1.25
    - [3094 - Take taints/tolerations into consideration when calculating PodTopologySpread skew](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/3094-pod-topology-spread-considering-taints) - v1.25
  - stable:
    - [1258 - Default Pod Topology Spread](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1258-default-pod-topology-spread) - v1.24
    - [1923 - Prefer Nominated Node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1923-prefer-nominated-node) - v1.24
    - [2249 - Namespace Selector for Pod Affinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2249-pod-affinity-namespace-selector) - v1.24
    - [785 - Scheduler Component Config API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/785-scheduler-component-config-api) - v1.25
    - [902 - Add NonPreempting Option For PriorityClasses](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/902-non-preempting-priorityclass) - v1.24


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   -
   -
   -

2. What metrics/community health stats does your group care about and/or measure?

   - Diversity 
   - Number of contributors
   - Meetings attendance 

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - It looks up-to-date 

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Nothing special

5. Does the group have contributors from multiple companies/affiliations?

   - Yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count: 3050
- Primary mailing list member count: 655
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 5
- Unique reviewers for SIG-owned packages: 6 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 4 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-scheduling#subprojects)



**New in 2022:**

  - kueue
  - kwok

**Retired in 2022:**

  - poseidon

**Continuing:**

  - cluster-capacity
  - descheduler
  - kube-batch
  - kube-scheduler-simulator
  - scheduler
  - scheduler-plugins


## [Working groups](https://git.k8s.io/community/sig-scheduling#working-groups)


**New in 2022:**

 - Batch

**Continuing:**

 - Multitenancy
 - Policy
 - Structured Logging

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
      - [2022 KubeCon NA: SIG-Scheduling Intro & Deep Dive](https://www.youtube.com/watch?v=1GpTE9L9oBM) 
      - [2022 KubeCon EU: SIG-Scheduling Intro & Deep Dive](https://www.youtube.com/watch?v=R2CpmLfHUYk)
      - 

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scheduling/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scheduling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
