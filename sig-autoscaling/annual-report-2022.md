# 2022 Annual Report: SIG Autoscaling

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Continued productionisation of VPA

2. What initiatives are you working on that aren't being tracked in KEPs?

   - [Multidimensional Pod Autoscaler](https://github.com/kubernetes/autoscaler/pull/5342)
   - [Balancer](https://github.com/kubernetes/autoscaler/blob/master/balancer/proposals/balancer.md)
   - Improving SIG Processes
      - [Increasing documentation of CA cloudprovider implementation expectations](https://github.com/kubernetes/autoscaler/pull/5198)
      - [Moving to a regular release schedule for CA](https://github.com/kubernetes/autoscaler/pull/5589)


3. KEP work in 2022 (v1.24, v1.25, v1.26):

- None


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [Horizontal Pod Autoscaler](https://github.com/kubernetes/kubernetes/pull/117178) - down to one approver and reviewer currently.
   - Cluster Autoscaler issue triage continues to be a challenge.
   -

2. What metrics/community health stats does your group care about and/or measure?

   - SIG Call engagement/attendance
   - Numbers of open PRs
   -

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A

5. Does the group have contributors from multiple companies/affiliations?

   - Yes - though the SIG's owners files continue to be heavily weighted towards a single company/

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - HPA - expanding the owners/reviewers and improving the reliability of tests in this area

## Membership

- Primary slack channel member count: 4470
- Primary mailing list member count: 291
- Primary meeting attendee count (estimated, if needed): ~10
- Primary meeting participant count (estimated, if needed): ~5
- Unique reviewers for SIG-owned packages: 9 (have aggregated this across all SIG owned subprojects)
- Unique approvers for SIG-owned packages: 9 (have aggregated this across all SIG owned subprojects)
- Cluster Autoscaler including cloud provider implementations: ~40

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-autoscaling#subprojects)



**Continuing:**

  - addon-resizer
  - cluster-autoscaler
  - horizontal-pod-autoscaler
  - vertical-pod-autoscaler


## [Working groups](https://git.k8s.io/community/sig-autoscaling#working-groups)


**New in 2022:**

 - [Balancer](https://github.com/kubernetes/autoscaler/blob/master/balancer/proposals/balancer.md)

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
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-autoscaling/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-autoscaling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
