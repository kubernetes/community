# 2022 Annual Report: SIG Release

## Current initiatives

SIG Release has a [roadmap](https://github.com/kubernetes/sig-release/blob/master/roadmap.md) that captures high level initiatives that the SIG is working toward, with specific information captured in our [project board](https://github.com/orgs/kubernetes/projects/30). 

1. What work did the SIG do this year that should be highlighted?
   - [Migrated most of deb/rpm package building into release process to reduce Google Build Admin involvement in releases](https://github.com/kubernetes/release/issues/2737)
   - Proof-of-concept of using OpenSUSE Build Service to build and publish packages using community infrastructure. Reflected in updates to [1731 - Publishing Kubernetes packages on community infrastructure](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/1731-publishing-packages). See [biweekly meeting from 18 October 2022](https://youtu.be/8l8X3vSAJAw?t=787) for a good overview.
   - [Signing of Release Artifacts](https://github.com/kubernetes/enhancements/issues/3031).

2. What initiatives are you working on that aren't being tracked in KEPs?

   - [Donation of a new project for SLSA Attestation](https://github.com/kubernetes-sigs/tejolote)
   - We begun work on a new process for onboarding release manager associates and a ladder for becoming a full release manager.

3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - beta:
    - [3000 - Artifact Distribution Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3000-artifact-distribution) - v1.25, v1.26
    - [3027 - SLSA Level 3 Compliance](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3027-slsa-compliance) - v1.25
    - [3031 - Signing release artifacts](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3031-signing-release-artifacts) - v1.25, v1.26

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [kubernetes-sigs/bom](https://github.com/kubernetes-sigs/bom/blob/main/OWNERS)
   - [kubernetes-sigs/downloadkubernetes](https://github.com/kubernetes-sigs/downloadkubernetes/blob/master/OWNERS)
   - [kubernetes-sigs/mdtoc](https://github.com/kubernetes-sigs/mdtoc/blob/master/OWNERS)
   - [kubernetes-sigs/release-notes](https://github.com/kubernetes-sigs/release-notes/blob/master/OWNERS)
   - [kubernetes-sigs/release-team-shadow-stats](https://github.com/kubernetes-sigs/release-team-shadow-stats/blob/master/OWNERS)
   - [kubernetes-sigs/tejolote](https://github.com/kubernetes-sigs/tejolote/blob/main/OWNERS)
   - [kubernetes-sigs/zeitgeist](https://github.com/kubernetes-sigs/zeitgeist/blob/master/OWNERS)
   
2. What metrics/community health stats does your group care about and/or measure?

   Some data tracking efforts that SIG Release performs include monitoring release team applications,
   release manager activities and code commits to ensure timely release cuts in our repos.

   In support of better understanding the diversity of the release team, [kubernetes-sigs/release-team-shadow-stats](https://github.com/kubernetes-sigs/release-team-shadow-stats) was begun to provide better reporting and visibility on release team metrics. For our KubeCon EU SIG update, we also presented a historical breakdown of the location of release team members and the geographic distribution of the release teams. 

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - The [sig-release CONTRIBUTING.md](https://github.com/kubernetes/sig-release/blob/master/CONTRIBUTING.md) could be updated to provivde more specific information regarding how to participate in both subprojects. An [issue](https://github.com/kubernetes/sig-release/issues/2200) was opened in [kuberentes/sig-release](https://github.com/kubernetes/sig-release) and an [issue](https://github.com/kubernetes/release/issues/2980) was created  in [kuberentes/release](https://github.com/kubernetes/release) to update this to collect information and make it more discoverable. 

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   We have a lot of information, including detailed role handbooks for both subprojects:

   - [Release Team](https://github.com/kubernetes/sig-release/tree/master/release-team)
   - [Release Engineering](https://github.com/kubernetes/sig-release/tree/master/release-engineering)

   However, these are not linked directly from CONTRIBUTING.md. We will use the [issue above](https://github.com/kubernetes/sig-release/issues/2200) to improve the discoverability of these. 


5. Does the group have contributors from multiple companies/affiliations?

   - Yes. Based on the last year of data from [dev stats](https://k8s.devstats.cncf.io/d/55/company-prs-in-repository-groups?orgId=1&var-period_name=Last%20year&var-repogroups=SIG%20Release&var-repos=All&var-companies=All&var-countries=All) we have had contributions from the following companies over the last year:
      - Red Hat Inc,
      - Chainguard Inc,
      - Intel Corporation
      - Liquid Reply
      - Kubermatic GmbH
      - Google LLC
      - Microsoft Corporation
      - Cisco
      - Amazon
      - VMware Inc
      - SUSE LLC
      - International Business Machines
      - Jetstack LTD
      - Mesosphere
      - Mastercard International Incorporated
      - DaoCloud Network Technology Co. Ltd.
      - Oracle America Inc.
      - Rackspace
      - NEC Corporation

   This data reflects company information in terms of PRs. We also have had contributions from individuals that have no company affiliation and several individuals from the CNCF.
   In addition to code contributions, the release teams during 2022 were staffed by individuals from a wide range of corporations, as well as students and other independent individuals. These contributions are not all reflected by the devstats query above, but are important to recognize.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - The release tooling is currently supported by the core of the Release Engineering team. There are opportunities for end users and Kubernetes distributors to support the maintenance of the tooling within our repositories, including the `bom` tool. Additionally, vendor companies that build and distribute their own Kubernetes releases could provide more support to SIG Release, specifically Release Engineering, in order to grow contributors that could help with important tasks like Go version updates. 
   

## Membership

The following stats are accurate as of March 15th, 2023. Numbers were pulled from Slack, the mailing list, and [kubernetes/release](https://github.com/kubernetes/release), the primary repository for Release Engineering tooling. 

- Primary slack channel member count: 2980
- Primary mailing list member count: 605
- Primary meeting attendee count (estimated, if needed): 15 (SIG Release BiWeekly), 20 (Release Team)
- Primary meeting participant count (estimated, if needed): 10 (SIG Release BiWeekly), 20 (Release Team)
- Unique reviewers for SIG-owned packages: 11
- Unique approvers for SIG-owned packages: 9


## [Subprojects](https://git.k8s.io/community/sig-release#subprojects)

**Retired in 2022:**

  - kubernetes/repo-infra

**Continuing:**

  - Release Engineering
  - Release Team
  - SIG Release Process Documentation


## [Working groups](https://git.k8s.io/community/sig-release#working-groups)


**Continuing:**

 - Reliability

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Releasing Kubernetes Less Often and More Secure - The SIG Release Update, KubeCon EU 2022](https://www.youtube.com/watch?v=qhQYu077zZU)
      - [How SIG Release Cooks Trustworthy Artifacts From Raw Source Code, KubeCon NA 2022](https://www.youtube.com/watch?v=F9Mvt4jm4uM)
      - [SIG Release Meeting at Contributor Summit NA](https://www.youtube.com/watch?v=dyXY5XoQnBM&list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&index=41&t=55s).

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-release/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-release/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
