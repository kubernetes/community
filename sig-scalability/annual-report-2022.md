# 2022 Annual Report: SIG Scalability

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - We helped validating scalability and reliability impact of many features across whole year.
   - We extended scalability testing framework and built a number of new optional tests to
     cover other usecases (e.g. batch workloads)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Localized scalability/performance improvements across the codebase, e.g.:
      - Improved LIST calls performance for rare selectors
      - Optimizing kube-apiserver compression
      - Optimizing watch by reducing number of allocations
      - Improved kube-apiserver graceful shutdown
   - All the improvements to scalability test frameworks in [perf-tests]

3. KEP work in 2022 (v1.24, v1.25, v1.26):

   - SIG Scalability doesn't own non-test code, so all below are officially
     tracked by other SIGs, but are driven (or co-driven) by SIG Scalability.
   - Stable
     - [1164 - Deprecate and Remove SelfLink](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1164-remove-selflink) - 1.24
     - [1904 - Efficient Watch Resumption](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption) - 1.24
   - Beta
     - [1040 - Priority and Fairness for API Server Requests](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1040-priority-and-fairness/README.md) - 1.24, 1.25, 1.26
     - [1669 - Proxy Terminating Endpoints](https://github.com/kubernetes/enhancements/blob/master/keps/sig-network/1669-proxy-terminating-endpoints/README.md) - 1.26
   - Alpha
     - [3453 - Minimize IPtables Restore](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3453-minimize-iptables-restore)
   - Pre-alpha
     - [3157 - Watch List](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list) - 1.26


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Each SIG Scalability subproject could benefit from additional hands.
     However, the Scalability Test Frameworks and Scalability and Performance tests and validation
     are the ones with the lower barier to entry - see [perf-tests] repository.
     Other areas require both very deep and wide understanding of Kubernetes before making reasonable contributions.

2. What metrics/community health stats does your group care about and/or measure?

   - We care mostly about metrics for our main repository [perf-tests]. We measure them through [devstats]. We keep track of number of reviewers, opened issues and PRs merged.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Yes, [CONTRIBUTING.md] points to maintained list of issues with label `good first issue`.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We don't have any special requirements for reviewers/approvers, but most of core areas of sig-scalability require
     very deep and wide understanding of the whole Kubernetes. It is much easier to start with our testing infrastructure.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, we have contributors from [multiple companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Scalability&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1641013200000&to=1672462800000).

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - End users usually come with questions about scalability, they rarely want to contribute.
   - End user could also help us by providing and documenting their scalability needs to help us
     better define our goals.

## Membership

- Primary slack channel member count: 2299
- Primary mailing list member count: 260
- Primary meeting attendee count (estimated, if needed): 7
- Primary meeting participant count (estimated, if needed): 5
- Unique reviewers for SIG-owned packages: 6
- Unique approvers for SIG-owned packages: 5

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-scalability#subprojects)

**New in 2022:**
- None

**Retired in 2022:**
- None

**Continuing:**
- [kubernetes-scalability-and-performance-tests-and-validation](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-and-performance-tests-and-validation)
- [kubernetes-scalability-bottlenecks-detection](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-bottlenecks-detection)
- [kubernetes-scalability-definition](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-definition)
- [kubernetes-scalability-governance](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-governance)
- [kubernetes-scalability-test-frameworks](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-test-frameworks)


## [Working groups](https://git.k8s.io/community/sig-scalability#working-groups)


**Continuing:**

 - Reliability

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
  - [Kubecon Valencia](https://kccnceu2022.sched.com/event/ytsB/intro-deep-dive-sig-scalability-marcel-zieba-wojciech-tyczynski-google)
  - [Kubecon Detroit](https://kccncna2022.sched.com/event/182Pr/intro-deep-dive-sig-scalability-marcel-zieba-google)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scalability/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scalability/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
[devstats]: k8s.devstats.cncf.io/
[perf-tests]: https://github.com/kubernetes/perf-tests/
