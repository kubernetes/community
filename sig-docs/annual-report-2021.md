# 2021 Annual Report: SIG Docs

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - SIG Docs put meaningful effort into growing its contributor and reviewer base in 2021, introducing [a shadow program for PR Wrangling](https://github.com/kubernetes/website/issues/31956) as well as dedicating more time to being active via our Slack community channel. This is an ongoing effort to grow our contributor base to a stable number, alongside easing the burden on our small group of reviewers and approvers.
   - Alongside growing our contributor base, SIG Docs also worked on a leadership transition strategy to bring community members into leadership roles. Via a specialized six month mentorship program expertly led by Steering Committee member Paris Pittman, SIG Docs was able to grow its leadership cohort for the main SIG, as well as some of its subgroups, adding new co-chairs and tech leads.
      - [SIG Docs google group](https://groups.google.com/g/kubernetes-sig-docs/)
      - [Call for help sent to dev@kubernetes.io, kubernetes-sig-leads, kubernetes-sig-docs](https://groups.google.com/g/kubernetes-sig-docs/c/hspG6mzgkrs)
      - [Announcement of new roles and leadership nominations](https://groups.google.com/g/kubernetes-sig-docs/c/cgrAyDLxydk)


2. What initiatives are you working on that aren't being tracked in KEPs?

   - Localization Subproject: SIG Docs is working on formalizing the localization work that has been ongoing for some time, with appointed leads of this initiative as well as recognizing the contributions of various community members across the different languages the Kubernetes website has been translated into. This subproject will be finalized by Q1 2022, with all active localizations informed and updated. The issue tracking the formalizing process can be viewed [here](https://github.com/kubernetes/website/issues/31955).
   - New Contributor Ambassador Program: As a continuation of our push to grow the SIG Docs contributor base, we're working on a specalized role that aims to support new and would-be contributors get up to speed with our processes and workflows. This role would be capped at six months for it to be shared amongst the community, with this feeding into a possible reviewer funnel as contributors get more comfortable with providing feedback to others. The issue tracking the formalization and documentation of this role can be viewed [here](https://github.com/kubernetes/website/issues/31946).

3. KEP work in 2021 (continuous and does not target any release):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - [1326 - Doc policies for third party content](https://git.k8s.io/enhancements/keps/sig-docs/1326-third-party-content-in-docs/README.md)


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Require more contributors to form a stable pool for the blog subproject, in particular and SIG Docs, generally.
   - Require guidelines on how we could ramp up folks nominated to the New Contributor Ambassador role specific to SIG Docs. 

2. What metrics/community health stats does your group care about and/or measure?


   - SIG Docs has a [dashboard](https://datastudio.google.com/u/0/reporting/fede2672-b2fd-402a-91d2-7473bdb10f04/page/567IC) available with site analytics. Some highlights include: 
       - 2021 Pages views: 111,565,437
       - Top pages for 2021 (excluding the home page): 
           - https://kubernetes.io/docs/reference/kubectl/cheatsheet/
           - https://kubernetes.io/docs/concepts/services-networking/service/
           - https://kubernetes.io/docs/concepts/services-networking/ingress/
   - PR velocity and open PR age is tracked in [Devstat](https://k8s.devstats.cncf.io/d/25/open-pr-age-by-repository-group?orgId=1&var-period=q&var-repogroup_name=SIG%20Docs&var-kind_name=All). We aim to have < 100 open PRs for the English localization, and will take steps as needed if we see the figure climbing much above that. For example, based on 12/31/2020 data, we have: Average number of opened PRs 368 with the median opened PRs age 5 days 21 hours 26 minutes 42 seconds

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - [It's updated](https://kubernetes.io/docs/contribute/) to the best of our knowledge.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We have a [well-documented guide](https://kubernetes.io/docs/contribute/) that details how folks can get started and the various ways in which they can scale the contributor ladder.
   - We actively assign good-first-issue labels to issues that we think early-lifecycle contributors could begin with. During KubeCons we also do introduction presentations and mentoring sessions to try and funnel contributors in to our SIG. Some of them are listed below,
      - [Intro to Kubernetes Docs](https://www.youtube.com/watch?v=pprMgmNzDcw), presented at KubeCon NA, 2020
      - [Kubernetes SIG Docs: A Deep Dive](https://www.youtube.com/watch?v=GDfcBF5et3Q), presented at KubeCon NA, 2021

5. Does the group have contributors from multiple companies/affiliations?

   - We have a contributor base spread across [98 companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Docs&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1609455600000&to=1639350000000) with top 50% contributions coming our way from IBM, Google, VMWare, RedHat, and NEC Corporation. We are al

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - With a contributor base that is spread globally, SIG Docs identified a potential opportunity for diversifying our documentation via multiple language support. Given the wide audience, this helps in making it more inclusive and user friendly for non-native English users of Kubernetes. Towards formalizing this ongoing initiative, SIG Docs will be finalizing it as a [localization subproject](https://github.com/kubernetes/website/issues/31955). We aim for the subproject to, potentially, provide more avenues for contribution. More details around leadership and the formalization process can be found in [this message](https://groups.google.com/a/kubernetes.io/g/dev/c/SP6weMvx3wg/m/l8LAL-OFCQAJ) sent to dev@kubernetes.io 

## Membership

- Primary slack channel member count: 1876
- Primary mailing list member count: 453
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: 76 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 93<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
- [kubernetes-blog](https://git.k8s.io/community/sig-docs#kubernetes-blog)
- [reference-docs](https://git.k8s.io/community/sig-docs#reference-docs)
- [website](https://git.k8s.io/community/sig-docs#website)

In 2021, SIG Docs started to formalize the localization subgroup into an official subproject.

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Kubernetes SIG Docs: A Deep Dive](https://www.youtube.com/watch?v=GDfcBF5et3Q), presented at KubeCon NA, 2021
      - [SIG Docs needs your help! to dev@kubernetes.io](https://groups.google.com/g/kubernetes-sig-docs/c/hspG6mzgkrs)
      - [SIG Docs Co-Chair Nomination: Divya Mohan to dev@kubernetes.io](https://groups.google.com/g/kubernetes-sig-docs/c/_1R7sh-_iiQ)
      - [SIG Docs: New roles & leadership nominations to dev@kubernetes.io](https://groups.google.com/a/kubernetes.io/g/dev/c/SP6weMvx3wg/m/l8LAL-OFCQAJ)
      - [SIG Docs APAC Meeting has changed to 1 AM IST/5.30 AM UTC on the last Wednesday of every month to kubernetes-sig-docs@googlegroups.com](https://groups.google.com/g/kubernetes-sig-docs/c/P7iLejmEIFA/m/-dPsBOpoDAAJ)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-docs/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-docs/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

