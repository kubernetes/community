# 2021 Annual Report: SIG Node

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - The CI testing subgroup crossed its two year threshold and managed to bring stability to the Kubelet Serial tests (and majority of other suites).
   - Deprecation of Dockershim
   -

2. What initiatives are you working on that aren't being tracked in KEPs?

   - CI/Testability improvements
   - Overall testability and stability improvements
   -

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

**TODO** (help-wanted, if anyone can generate these before I get to it)

   - Stable
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.stable
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.stable
   - Beta
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.beta
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.beta
   - Alpha
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.alpha
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.alpha
   - Pre-alpha
     - [$kep-number - $title](https://git.k8s.io/community/$link/README.md)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - **TODO** Feedback wanted

2. What metrics/community health stats does your group care about and/or measure?

   - Incoming vs Completed PRs get a readout during every community meeting with
     approximated tracking of trends.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - The Node CONTRIBUTING.md has some useful documentation for new contributors, such
     as the [Getting Started](https://github.com/kubernetes/community/blob/master/sig-node/CONTRIBUTING.md#getting-started)
     section. We need to make some improvements to talk about how we work and provide
     more guidance. That work is tracked in [kubernetes/community/6611](https://github.com/kubernetes/community/issues/6611). 

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We have additional requirements that are not currently documented publicly.
     The work to move those to community documentation is tracked in [kubernetes/community/6612](https://github.com/kubernetes/community/issues/6612).

5. Does the group have contributors from multiple companies/affiliations?

   - We have contributors from a wide range of companies including Google,
     RedHat, VMware, Intel, and Nvidia. Sig Leads are from Google and RedHat.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - sig-node today struggles to accept new features or changes in a way that
     allows us to ship with confidence. The sig needs contributors to help pay
     down technical debt in tests, test infrastucture, and overall
     maintainability of the codebase.
   - If a full time maintainer joined the sig, there would be an array of
     projects that in practice take full time commitment to do well:
       - Ensuring that test-grid is green and following up when tests fail or
         flakiness changes.
       - Paying down technical debt throughout the kubelet through adding or
         improving tests, and aiding with refactors and rewrites.
       - Introducing new forms of testing to find and fix latent bugs in the
         kubelet and improving our confidence in changes.

## Membership

- Primary slack channel member count: 3141
- Primary mailing list member count: 725
- Primary meeting attendee count (estimated, if needed): ~15-30
- Primary meeting participant count (estimated, if needed): ~10
- Unique reviewers for SIG-owned packages: 
- Unique approvers for SIG-owned packages: 7 top level approvers

Include any other ways you measure group membership

## Subprojects

New in 2021: None

Retired in 2021: None

Continuing:

- [cri-api](https://git.k8s.io/community/sig-node#cri-api)
- [cri-tools](https://git.k8s.io/community/sig-node#cri-tools)
- [ci-testing](https://git.k8s.io/community/sig-node#ci-testing)
- [kubelet](https://git.k8s.io/community/sig-node#kubelet)
- [node-api](https://git.k8s.io/community/sig-node#node-api)
- [node-feature-discovery](https://git.k8s.io/community/sig-node#node-feature-discovery)
- [node-problem-detector](https://git.k8s.io/community/sig-node#node-problem-detector)
-
[noderesourcetopology-api](https://git.k8s.io/community/sig-node#$noderesourcetopology-api)
- [security-profiles-operator](https://git.k8s.io/community/sig-node#security-profiles-operator)

## Working groups

New in 2021:
- [wg-policy](https://git.k8s.io/community/wg-policy/) ([2021 report](https://git.k8s.io/community/wg-policy/annual-report-2021.md))
- [wg-structured-logging](https://git.k8s.io/community/wg-structured-logging/) ([2021 report](https://git.k8s.io/community/wg-structured-logging/annual-report-2021.md))

Retired in 2021: None

Continuing:
- [wg-multitenancy](https://git.k8s.io/community/wg-multitenancy/) ([2021 report](https://git.k8s.io/community/wg-multitenancy/annual-report-2021.md))
-

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-node/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-node/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

