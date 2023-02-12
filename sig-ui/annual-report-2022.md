# 2022 Annual Report: SIG UI

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Kubernetes Dashboard v2.5.0 - v2.7.0 released in 2022, and latest release is v2.7.0 in Sep 2022. Each releases support latest release of Kubernetes at the time.
   - Not only Kubernetes modules, but also almost all dependencies are updated to the latest versions.
   - The biggest highlight is that new archtecture was introduced. This splited containers for frontend and backend into each and this makes it possible to tune the each performance.
   - We are now preparing new release that includes the new architecture.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We are aiming to adapt tools to new architecture.
      - [Helm Chart](https://github.com/kubernetes/dashboard/pull/7544)
      - [Development environment](https://github.com/kubernetes/dashboard/pull/7602)

3. KEP work in 2022 (v1.24, v1.25, v1.26):

   - None

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Recruiting new contributors even if in partial area. To be able to traverse/contribute, a new contributor needs to have an understanding of Angular, Golang, and a reasonably knowledge of Kubernetes client-go package.
   - More contributors needed for fixing functional gaps between dashboard and Kubernetes API or kubectl.
   - To support for RTL languages, we needs who can develop dashbboard and review RTL languages.
   - More approvers/reviewers are needed for following translation sub teams:
      - [French translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/modules/web/i18n/fr/OWNERS)
      - [Korean translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/modules/web/i18n/ko/OWNERS)
      - [Spanish translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/modules/web/i18n/es/OWNERS)

2. What metrics/community health stats does your group care about and/or measure?

   - There are two primary sources of PRs: Automated dependency PRs (Dependabot) and contributor-created PRs. Dependabot PRs are merged or closed within 24h, while contributor-created PRs are typically reviewed within 3-4 days.
   - Currently there are no real metrics that we measure regarding Issue/PR turnaround.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - We have well documented guide for [getting started](https://github.com/kubernetes/dashboard/blob/master/docs/developer/getting-started.md) and our [CONTRIBUTING.md](https://github.com/kubernetes/dashboard/blob/master/CONTRIBUTING.md) leads to it.
   - Also, we attempt to maintain Good First Issue labels on things that we think early-lifecycle contributors could begin with. During KubeCons we also do introduction presentations and mentoring sessions to try and funnel contributors in to our SIG.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We have translation sub teams in Kubernetes Dashboard that have responsibility to manage translation files for each languages. And we have [special guide](https://github.com/kubernetes/dashboard/blob/master/docs/developer/internationalization.md) to organize and to manage the translation sub teams.
   - Also, the development environment container is available to facilitate translation work.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, current chairs are spread between Plural and NEC. And contributors are also spread across multiple companies including Wiremind, PWC and Tencent. See more [contributors in 2021](https://github.com/kubernetes/dashboard/graphs/contributors?from=2022-01-01&to=2022-12-31&type=c).

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - [Internationalization](https://github.com/kubernetes/dashboard/blob/master/docs/developer/internationalization.md) is easy for contributing.

## Membership

- Primary slack channel member count: 1245
- Primary mailing list member count: 223
- Primary meeting attendee count (estimated, if needed): N/A
- Primary meeting participant count (estimated, if needed): N/A
- Unique reviewers for SIG-owned packages: 3<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 3<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-ui#subprojects)

**Continuing:**

- [kubernetes/dashboard](https://github.com/kubernetes/dashboard)
- [kubernetes-sigs/dashboard-metrics-scraper](https://github.com/kubernetes-sigs/dashboard-metrics-scraper)

## [Working groups](https://git.k8s.io/community/sig-ui#working-groups)

- None

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
      - None
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Kubernetes SIG UI Introduction and Updates at KubeCon+CloudNativeCon Europe 2022](https://sched.co/ytpx)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-ui/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-ui/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
