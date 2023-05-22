# 2022 Annual Report: SIG Testing

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Ginkgov2 migration
     - https://github.com/kubernetes/test-infra/pull/26250
     - https://github.com/kubernetes/kubernetes/pull/109111
   - Utilizing context provided by Ginkgo
     - https://github.com/kubernetes/kubernetes/pull/112923
   - Test failure description improvements
     - https://github.com/kubernetes/kubernetes/pull/113538
   - Launched prow docs site
   - TestGrid API now available at testgrid-data.k8s.io
   - Prow API service ("Gangway" component) merged, enabling Prow installations to programmatically trigger Prow jobs

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Continue to improve testing and enforcing best practices on software development
   - Cost efficiency:
     - Clean up orphan jobs
     - Optimize existing tooling: artifact storage, log size, …
     - Extend the CI to other Cloud Providers
   - Development of a new TestGrid UI


3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - beta:
    - [3041 - NodeConformance and NodeFeature labels cleanup](https://github.com/kubernetes/enhancements/tree/master/keps/sig-testing/3041-node-conformance-and-features) - v1.26


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Additional reviewers (and eventually approvers) on [Prow](https://github.com/kubernetes/test-infra/tree/master/prow)
   - Many less-active tools in test-infra, such as:
     - [Boskos](https://github.com/kubernetes/test-infra/tree/master/boskos)
     - [Kettle](https://github.com/kubernetes/test-infra/tree/master/kettle)
     - [Kubetest2](https://github.com/kubernetes-sigs/kubetest2)

2. What metrics/community health stats does your group care about and/or measure?

   - Reviewers and approvers

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Yes! See [k8s/sig-testing: Contribution](https://github.com/kubernetes/sig-testing#contribution) and [k8s/test-infra: Issue Triage](https://github.com/kubernetes/test-infra/blob/master/CONTRIBUTING.md#issue-triage).

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A: no special training, requirements, or processes.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - (_more reviewers, promoting active reviewers to approvers?_)
   - (_to be filled_)

## Membership

- Primary slack channel member count: 2,452
- Primary mailing list member count: 407
- Primary meeting attendee count (estimated, if needed): 7-8 on average
- Primary meeting participant count (estimated, if needed): 5-6, estimate
- Unique reviewers for SIG-owned packages: 24
- Unique approvers for SIG-owned packages: 26

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-testing#subprojects)



**Retired in 2022:**

  - k8s-gsm-tools

**Continuing:**

  - boskos
  - e2e-framework
  - kind
  - kubetest2
  - prow
  - sig-testing
  - test-infra
  - testing-commons


## [Working groups](https://git.k8s.io/community/sig-testing#working-groups)


**Continuing:**

 - Reliability

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - KubeCon NA 2022: [SIG Testing: Intro And Updates](https://www.youtube.com/watch?v=CdKBl6CncHg)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-testing/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-testing/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
