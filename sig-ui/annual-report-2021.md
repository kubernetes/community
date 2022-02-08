# 2021 Annual Report: SIG UI

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Kubernetes Dashboard v2.2.0-v2.4.0 released in 2021, and latest release is v2.5.0 in Feb 2022. Each releases support latest release of K8s at the time.
   - Almost all dependencies are updated to the latest versions. As a result, Nodejs has been upgraded to v16.x, Angular to v13.x, and Golang to v17.x.
   - In addition to the ongoing effort of maintenance, we are improving support for [various resource objects](https://github.com/kubernetes/dashboard/issues/5232) and adding the ability for installations to have custom themes.
   - Another major area to highlight is that we are continually onboarding new language translations for the Dashboard. Spanish support added recently by [#6587].
   - Furthermore, by migrating the task runner from Gulp to Make and updating the configs and scripts, various improvements that were previously blocked have become possible.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - There have been no KEPs opened as part of SIG-UI.

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - None

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Recruiting new contributors even if in partial area. To be able to traverse/contribute, a new contributor needs to have an understanding of Angular, Golang, and a reasonably knowledge of Kubernetes client-go package.
   - More contributors needed for fixing functional gaps between dashboard and Kubernetes API or kubectl.
   - [Support for RTL languages](https://github.com/kubernetes/dashboard/pull/6305) needs who can develop dashbboard and review such languages.
   - [French translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/i18n/fr/OWNERS)
   - [Japanese translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/i18n/ja/OWNERS)
   - [Korean translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/i18n/ko/OWNERS)
   - [Spanish translation sub team for Kubernetes Dashboard](https://github.com/kubernetes/dashboard/blob/master/i18n/es/OWNERS)

2. What metrics/community health stats does your group care about and/or measure?

   - There are two primary sources of PRs: Automated dependency PRs (Dependabot) and contributor-created PRs. Dependabot PRs are merged or closed within 24h, while contributor-created PRs are typically reviewed within 3-4 days.
   - Currently there are no real metrics that we measure regarding Issue/PR turnaround.

3. Does your [CONTRIBUTING.md](https://github.com/kubernetes/dashboard/tree/master/CONTRIBUTING.md) help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - We have well documented guide for [getting started](https://github.com/kubernetes/dashboard/blob/master/docs/developer/getting-started.md).
   - Also, we attempt to maintain Good First Issue labels on things that we think early-lifecycle contributors could begin with. During KubeCons we also do introduction presentations and mentoring sessions to try and funnel contributors in to our SIG.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md](https://github.com/kubernetes/dashboard/tree/master/CONTRIBUTING.md) document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We have translation sub teams in Kubernetes Dashboard that have responsibility to manage translation files for each languages. And we have [special guide](https://github.com/kubernetes/dashboard/blob/master/docs/developer/internationalization.md) to organize and to manage the translation sub teams.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, current chairs are spread between Kubermatic and NEC. And contributors are also spread across multiple companies including Wiremind, PWC and Tencent. See more [contributors in 2021](https://github.com/kubernetes/dashboard/graphs/contributors?from=2021-01-01&to=2021-12-31&type=c).


6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - [Internationalization](https://github.com/kubernetes/dashboard/blob/master/docs/developer/internationalization.md) is easy for contributing.

## Membership

- Primary slack channel member count: 1164
- Primary mailing list member count: 223
- Primary meeting attendee count (estimated, if needed): N/A
- Primary meeting participant count (estimated, if needed): N/A
- Unique reviewers for SIG-owned packages: 4<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 4<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership
- [Contributors for SIG UI in last year by devstats](https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=SIG%20UI&var-repo_name=kubernetes%2Fkubernetes&var-country_name=All): 109

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
- [kubernetes/dashboard](https://github.com/kubernetes/dashboard)
- [kubernetes-sigs/dashboard-metrics-scraper](https://github.com/kubernetes-sigs/dashboard-metrics-scraper)

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md](./README.md) reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md](https://github.com/kubernetes/dashboard/CONTRIBUTING.md) reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
      - None
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - None

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-ui/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-ui/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
