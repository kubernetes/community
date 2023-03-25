<<<<<<< HEAD
# 2022 Annual Report: SIG Docs

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   -
   -
   -

2. What initiatives are you working on that aren't being tracked in KEPs?

   -
   -
   -



3. KEP work in 2022 (v1.24, v1.25, v1.26):


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   -
   -
   -

2. What metrics/community health stats does your group care about and/or measure?

   -
   -
   -

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   -

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   -

5. Does the group have contributors from multiple companies/affiliations?

   -

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count:
- Primary mailing list member count:
- Primary meeting attendee count (estimated, if needed):
- Primary meeting participant count (estimated, if needed):
- Unique reviewers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-docs#subprojects)



**New in 2022:**

  - localization

**Continuing:**

  - kubernetes-blog
  - reference-docs
  - website


## [Working groups](https://git.k8s.io/community/sig-docs#working-groups)


## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-docs/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-docs/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
=======
# **2022 Annual Report: SIG Docs**

***Current initiatives***

What work did the SIG do this year that should be highlighted?

* Contributors across many SIGs, companies, and foundations worked together with SIG Docs to finalize the [dockershim removal](https://kubernetes.io/blog/2022/02/17/dockershim-faq/) from the Kubernetes website
* Our Localization subproject was officially documented in our community resources, with two subproject leads appointed and the `area/localization` label created
* The Hindi localization, our first localized set of documentation using the Devanagari script, officially launched
* A new [Issue Wrangler](https://github.com/kubernetes/website/discussions/38861) role, as a compliment to the [SIG Docs PR Wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) role, has been put together to tackle issue triage and feature requests more efficiently
* (add more)

**What initiatives are you working on that aren't being tracked in KEPs?**

* Katacoda tutorial [shutdown notice](https://github.com/kubernetes/website/issues/33936) and replacement work (see Katacoda Shutdown information on the Kubernetes blog [here](https://kubernetes.io/blog/2023/02/14/kubernetes-katacoda-tutorials-stop-from-2023-03-31/))
* [Diagram implementation guide](https://kubernetes.io/docs/contribute/style/diagram-guide/) for Kubernetes documentation, using the Mermaid JavaScript library
* (did i miss something?)

**KEP work in 2022 (v1.24, v1.25, v1.26):**

* N/A

***Project health***

**What areas and/or subprojects does your group need the most help with? Any areas with 2 or fewer OWNERs? (link to more details)**

* Our localizations have two owners as a minimum default, thus, SIG Docs doesn't count this as a project health issue. Our Localization subproject had one of its leads move to Emeritus in early 2023, however, we're in the process of mentoring a second lead throughout 2023
* Our blog subproject is still short on resources, however, we have been able to add an additional reviewer over the last year. We're still looking to add to the very small pool of active editors, which are our most critical resource for article publication
* We're looking for active contributors to take on the Issue Wrangling role, which would require some involvement in community meetings and Slack discussion to better triage issues and bugs
* Project management is an area that our leads have been trying to cover, but SIG Docs overall could use more help in this space. Help with coordinating large-scale efforts such as the Katacoda tutorial replacement, or organizing the various attempts to debug Netlify issues, would be a huge benefit
* (where else do we need help?)

**What metrics/community health stats does your group care about and/or measure?**

* SIG Docs has a dashboard available with [site analytics](https://lookerstudio.google.com/u/0/reporting/fede2672-b2fd-402a-91d2-7473bdb10f04/page/567IC). Some highlights include:
    * 2022 Pages views: 112,180,650
    * Top pages for 2022 (excluding the home page):
        * https://kubernetes.io/docs/reference/kubectl/cheatsheet/
        * https://kubernetes.io/docs/concepts/services-networking/service/
        * https://kubernetes.io/docs/concepts/services-networking/ingress/
    * It is worth noting that our top pages in 2022 mirror our 2021 results
* PR velocity and open PR age is tracked in [Devstats](https://k8s.devstats.cncf.io/d/25/open-pr-age-by-repository-group?orgId=1&var-period=q&var-repogroup_name=SIG%20Docs&var-kind_name=All&from=1640991600000&to=1672527599000). We aim to have < 150 open PRs for the English localization, and will take steps as needed if we see the figure climbing much above that. It's worth noting that during every documentation review period for a release, our PR number in the English localization will climb significantly

**Does your CONTRIBUTING.md help new contributors engage with your group specifically by pointing to activities or programs that provide useful context or allow easy participation?**

* It is [updated](https://kubernetes.io/docs/contribute/) to the best of our knowledge (this isn't true, we need to make some tweaks)

**If your group has special training, requirements for reviewers/approvers, or processes beyond the general contributor guide, does your CONTRIBUTING.md document those to help existing contributors grow throughout the contributor ladder?**

* We have [a well-documented guide](https://kubernetes.io/docs/contribute/) that details how folks can get started and the various ways in which they can scale the contributor ladder
* We actively assign good-first-issue labels to issues that we think early-lifecycle contributors could begin with, with other contributors sharing these issues actively in the #sig-docs Slack channel. We've also created introduction presentations and mentoring sessions at various KubeCon events to try and funnel contributors into our SIG. Some of them are listed below:
    * [Intro to Kubernetes Docs](https://www.youtube.com/watch?v=pprMgmNzDcw) – presented at KubeCon/CloudNativeCon NA, 2020
    * [Kubernetes SIG Docs: A Deep Dive](https://www.youtube.com/watch?v=GDfcBF5et3Q) – presented at KubeCon/CloudNativeCon NA, 2021


**Does the group have contributors from multiple companies/affiliations?**

* Yes. We have a contributor base spread across [98 companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Docs&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1640991600000&to=1670886000000) with the top 50% of contributions coming from IBM, DaoCloud, Google, NEC Corporation, VMWare, and Red Hat. We'd like to give a shout-out to all of the independent contributors who feature third on this ranking in terms of number of contributions
* Our group also strives to be a positive example when it comes to Diversity, Equity, and Inclusion, where we ensure we have SIG leads from some of our biggest membership bases (India, China), alongside women in technical leadership roles (co-chairs, tech leads)

**Are there ways end users/companies can contribute that they currently are not? If one of those ways is more full time support, what would they work on and why?**

* (to add)

***Membership***

* Primary Slack channel member count: 2,264
* Primary mailing list member count: 576
* Primary meeting attendee count (estimated, if needed): 12
* Primary meeting participant count (estimated, if needed): 6
* Unique reviewers for SIG-owned packages: xx
* Unique approvers for SIG-owned packages: xx

***Subprojects***

New in 2022:

    localization

Continuing:

    kubernetes-blog
    reference-docs
    website

***Operational***

Operational tasks in sig-governance.md:

- [ ] README.md reviewed for accuracy and updated if needed
- [ ] CONTRIBUTING.md reviewed for accuracy and updated if needed (or created if missing and your contributor steps and experience are different or more in-depth than the documentation listed in the general contributor guide and devel folder.)
- [ ] Subprojects list and linked OWNERS files in sigs.yaml reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in sigs.yaml are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2022 are linked from README.md and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2022 (e.g. community meetings, KubeCon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
    * [Localization Subproject Launch to SIG Docs](https://groups.google.com/g/kubernetes-sig-docs/c/DUFVEi-9tWc/m/ji2vPwBEBQAJ)
    * (I cannot find our KubeCon NA 2022 talk but lets please list it here)
>>>>>>> 4ad06693 (add initial draft of sig docs 2022 annual report)
