# 2021 Annual Report: SIG Scalability

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - We helped validating scalability and reliability impact of many features across whole year.
   - We have started testing scalability of large services (1000+ pods).

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We've improved testing framework and infrastructure, for example:
     - Added support for modules in our tests
     - Started measuring availability of api-server
     - Added support for measuring cilium propagation delay, dns latency

3. KEP work in 2021:
   - Stable
     - None
   - Beta
     - [1040 - Priority and Fairness for API Server Requests](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1040-priority-and-fairness/README.md) - 1.23
   - Alpha
     - [647 - APIServer Tracing](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/647-apiserver-tracing/README.md) - 1.22
     - [1669 - Proxy Terminating Endpoints](https://github.com/kubernetes/enhancements/blob/master/keps/sig-network/1669-proxy-terminating-endpoints/README.md) - 1.22
     - [2464 - Kubetest2 CI Migration](https://github.com/kubernetes/enhancements/blob/master/keps/sig-testing/2464-kubetest2-ci-migration/README.md) - 1.21
   - Pre-alpha
     - None

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Each subproject could benefit from additional hands. However, the Scalability Test Frameworks and Scalability and Performance tests and validation are the ones where we can grow contributors - the other require both very deep and wide understanding of Kubernetes before making reasonable contributions.

2. What metrics/community health stats does your group care about and/or measure?

   - We care mostly about metrics for our main repository [perf-tests]. We measure them through [devstats]. We keep track of number of reviewers, opened issues and PRs merged.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Yes, [CONTRIBUTING.md] points to maintained list of issues with label `good first issue`.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We don't have any special requirements for reviewers/approvers, but most of core areas of sig-scalability required deep and wide understanding of Kubernetes.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, we have contributors from [multiple companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Scalability&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1609477200000&to=1640926800000).

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - End users usually come with questions about scalability, they rarely want to contribute.

## Membership

- Primary slack channel member count: 2029
- Primary mailing list member count: 236
- Primary meeting attendee count (estimated, if needed): 5
- Primary meeting participant count (estimated, if needed): 4
- Unique reviewers for SIG-owned packages: 6
- Unique approvers for SIG-owned packages: 5

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- None

Retired in 2021:
- None

Continuing:
- [kubernetes-scalability-and-performance-tests-and-validation](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-and-performance-tests-and-validation)
- [kubernetes-scalability-bottlenecks-detection](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-bottlenecks-detection)
- [kubernetes-scalability-definition](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-definition)
- [kubernetes-scalability-governance](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-governance)
- [kubernetes-scalability-test-frameworks](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-test-frameworks)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- None

Retired in 2021:
- None

Continuing:
- [WG Reliability](https://git.k8s.io/community/wg-reliability/) ([2021 report](https://github.com/kubernetes/community/blob/master/wg-reliability/annual-report-2021.md))

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
      - [Kubecon NA 2021](https://www.youtube.com/watch?v=sEdSoWslQ6A)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scalability/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scalability/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
[perf-tests]: https://github.com/kubernetes/perf-tests
[devstats]: k8s.devstats.cncf.io/
