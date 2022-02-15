# 2021 Annual Report: SIG Release

## Current initiatives

### 1. What work did the SIG do this year that should be highlighted?

#### Release Engineering

After finalizing the rewrite of the release process from bash into golang,
the release engineering team has been focusing its efforts on two main
areas:

   1. Improving the release automation on two fronts:
      1. Adding new features, tests and checks to the release process which
         were missing from the original anago (binary verification, CVE
         disclosure, building from custom branches and repositories).
      1. Consolidating the codebases of new repositories which SIG Release
         brought under its responsibility. The range of new repositories we
         are consolidating go from critical projects (like the image promoter)
         to less important repositories (like downloadkubernetes.com) 
   1. Hardening the Kubernetes Supply Chain via key efforts:
      1. SBOM Generation
      1. SLSA 3 compliance
      1. Artifact signing 

#### Release Team

### 2. What initiatives are you working on that aren't being tracked in KEPs?

The most important change currently under development not tracked in a KEP is
the new automated branch forward. Tests are currently underway and we aim to
have automated forward of the release branch during code freeze by the 1.25 cycle.
[A recent announcement sent to the dev mailing list](https://groups.google.com/a/kubernetes.io/g/dev/c/qbHPJjUF3s8)
has more details about the plan.

### 3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Alpha
     - [KEP-2853 - Kubernetes repository branch rename](https://github.com/kubernetes/enhancements/blob/master/keps/sig-release/2853-k-core-branch-rename/README.md) - $milestone.stable
     - [KEP-3027 - SLSA Level 3 Compliance in the Kubernetes Release Process](https://github.com/kubernetes/enhancements/blob/master/keps/sig-release/3027-slsa-compliance/README.md) - $milestone.stable
     - [KEP-3031: Signing release artifacts](https://github.com/kubernetes/enhancements/blob/master/keps/sig-release/3031-signing-release-artifacts/README.mdhttps://git.k8s.io/community/$link/README.md) - $milestone.beta
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.beta

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   All of the following areas are reviewed by the Release Engineering
   subproject, but we could always use more help here:

   - [kubernetes-sigs/bom](https://github.com/kubernetes-sigs/bom/blob/main/OWNERS)
   - [kubernetes-sigs/downloadkubernetes](https://github.com/kubernetes-sigs/downloadkubernetes/blob/master/OWNERS)
   - [kubernetes-sigs/mdtoc](https://github.com/kubernetes-sigs/mdtoc/blob/master/OWNERS)
   - [kubernetes-sigs/release-notes](https://github.com/kubernetes-sigs/release-notes/blob/master/OWNERS)
   - [kubernetes-sigs/zeitgeist](https://github.com/kubernetes-sigs/zeitgeist/blob/master/OWNERS)
   - [kubernetes/repo-infra](https://github.com/kubernetes/repo-infra/blob/master/OWNERS)

2. What metrics/community health stats does your group care about and/or measure?

   -
   -
   -

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   -

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - [Release Team](https://github.com/kubernetes/sig-release/tree/master/release-team)
   - [Release Engineering](https://github.com/kubernetes/sig-release/tree/master/release-engineering)

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, over the past two years, we've had contributors from the following companies (non-exhaustive, gathered from [here](https://k8s.devstats.cncf.io/d/55/company-prs-in-repository-groups?orgId=1&var-period_name=Last%202%20years&var-repogroups=SIG%20Release&var-repos=All&var-companies=All&var-countries=All)):
     - Red Hat
     - Cisco
     - Chainguard
     - Mattermost
     - Apple
     - SUSE
     - VMware
     - Upbound
     - Google
     - Jetstack
     - Kubermatic
     - IBM
     - HashiCorp
     - SAP
     - HSBC
     - Huawei
     - Intel
     - Autodesk

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

Accurate of 2022-02-14.
Stats are primarily pulled from kubernetes/release, the primary repository for
Release Engineering tooling/work, which serves as a reasonable representation
of reviewers/approvers across SIG Release repositories.

- Primary Slack channel member count: 2458
- Primary mailing list member count: 501
- Primary meeting attendee count (estimated, if needed): 20
- Primary meeting participant count (estimated, if needed): 10
- Unique reviewers for SIG-owned packages (from kubernetes/release): 24
- Unique approvers for SIG-owned packages (from kubernetes/release): 7

Include any other ways you measure group membership

## Subprojects

Retired in 2021:

- [Licensing](https://git.k8s.io/community/sig-release#licensing)

Continuing:

- [Release Engineering](https://git.k8s.io/community/sig-release#release-engineering)
- [Release Team](https://git.k8s.io/community/sig-release#release-team)

## Working groups

New in 2021:

- [WG Reliability](https://git.k8s.io/community/wg-reliability/) ([2021 report](https://git.k8s.io/community/wg-reliability/annual-report-2021.md))

Retired in 2021:

WG K8s Infra was converted into [SIG K8s Infra](https://git.k8s.io/community/sig-k8s-infra) in 2021.

Continuing:

- [WG Reliability](https://git.k8s.io/community/wg-reliability/) ([2021 report](https://git.k8s.io/community/wg-reliability/annual-report-2021.md))

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

      - [Hardening the Kubernetes Software Supply Chain Through Better Transparency](https://www.youtube.com/watch?v=W6hUXv66rRc) KubeCon + CloudNativeCon NA 2021

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-release/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-release/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
<!-- BEGIN CUSTOM CONTENT -->

<!-- END CUSTOM CONTENT -->
