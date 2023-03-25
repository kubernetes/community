# 2022 Annual Report: SIG Multicluster

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Significant progress on KEP-1645, the [MCS API](https://sigs.k8s.io/mcs-api) now with conformance tests
   - Developed the About API for ClusterSets

2. What initiatives are you working on that aren't being tracked in KEPs?

   - SIG-MC is building a website ([multicluster.sigs.k8s.io](https://multicluster.sigs.k8s.io/)) to track and communicate multicluster related initiatives, patterns, and best practices. ([contribute](https://github.com/kubernetes-sigs/sig-multicluster-site))
   - We are working with the Gateway API project to formalize interaction with multi-cluster services
   - The SIG is exploring cross-cluster workload management - what might it take to support multicluster controllers? and the [work API](https://github.com/kubernetes-sigs/work-api)

3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2149 - ClusterID for ClusterSet Identification](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid) - v1.24
  - progress to beta:
    - [1645 - MCS API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/1645-multi-cluster-services-api)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - There has been continued and repeated interest in multicluster controller development, but we need more people to come and share specific use cases, previous attempts at implementations, and/or concerns before we can hope to propose any generally useful solutions.
   - See this [doc](https://docs.google.com/document/d/1VMEv8ivK0ovLN8PTv9fGsTqUAErWkv6sE98bdju13oM/edit#heading=h.u7jfy9wqpd2b) for areas the SIG has been exploring. Help trying things out, refining use cases, and prototyping would be greatly appreciated.

2. What metrics/community health stats does your group care about and/or measure?

   - We tend to focus most on meetup attendees and presenters, and contributors.
     - In H2 2022 we averaged 15 attendees/week, now up to ~20 in 2023-03
     - 2-3 presenters each week in 2023, up from ~1 in H2 2022
   - We have regular discussion and repo contributions from 5-6 companies.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Yes

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A

5. Does the group have contributors from multiple companies/affiliations?

   - Yes - contributions from 5-6 companies. Most bi-weekly SIG meetings have contributors from 4+ companies.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - See [multicluster.sigs.k8s.io/contributing](https://multicluster.sigs.k8s.io/contributing/) to get started.
   - Check our [bi-weekly discussion notes](https://goo.gl/PhyWjj) for recent topics and areas that could use attention.
   - Bring your own ideas or pain points.
   - The community we have is great, we just need more of it!

## Membership

- Primary slack channel member count: 3335
- Primary mailing list member count: 841
- Primary meeting attendee count (estimated, if needed): ~20
- Primary meeting participant count (estimated, if needed): ~8
- Unique reviewers for SIG-owned packages: 6
- Unique approvers for SIG-owned packages: 6

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-multicluster#subprojects)



**Retired in 2022:**

  - kubemci
  - Kubefed

**Continuing:**

  - about-api
  - mcs-api
  - work-api


<!-- ## [Working groups](https://git.k8s.io/community/sig-multicluster#working-groups)

NEED CLEANUP

**Continuing:**

 - IoT Edge
 - Policy -->

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
    - Bi-weekly meeting notes [Meeting notes and Agenda](https://goo.gl/PhyWjj)
    - [Kubecon NA Update](https://kccncna2022.sched.com/event/182P2/sig-multicluster-intro-and-deep-dive-jeremy-olmsted-thompson-laura-lorenz-google-paul-morie-apple)
    - [Kubecon EU Update](https://kccnceu2022.sched.com/event/ytq6/sig-multicluster-intro-and-deep-dive-jeremy-olmsted-thompson-laura-lorenz-google-paul-morie-apple)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
