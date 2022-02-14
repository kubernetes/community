# 2021 Annual Report: SIG Cluster Lifecycle

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - We started preparation for the Dockershim removal in 1.24. Minikube and kubeadm
   maintainers helped with documentation updates, implementing changes and talking
   to the cri-dockerd maintainers.
   - Cluster API continues on its road to maturity and released its first MAJOR release
   (v1.0) of the GitHub repository. The API itself graduated to v1beta1.
   - kubeadm continues with some of the efforts to solve technical debt issues
   and further stabilize the project. A new API version v1beta3 was released.
   - Support for pure IPv6 was added to kOps. The contributions to this area of kOps
   resulted in collaboration with the wider ecosystem.
   - For more highlights check the "Current initiatives" section of our
   [yearly subproject survey](https://forms.gle/xZn8DXww4XxPsXvCA).


2. What initiatives are you working on that aren't being tracked in KEPs?

   - SIG Cluster Lifecyle has not added any new KEPs that affect the whole group or the wider Kubernetes.
   New KEPs and KEP updates were only done for [kubeadm](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm),
   which is the only subproject that is part of the Kubernetes release artifacts.
   Most of the activity in this group happens outside of the KEP process.
   Individual projects have implemented their own proposal tracking means, such
   as the [Cluster API CAEP process](https://sigs.k8s.io/cluster-api/docs/proposals).
   - A general theme across subprojects is driving them to maturity.
   With the exception of projects like kubeadm, minikube, kOps and kubespray most of
   our other subprojects are fairly new in Kubernetes years.

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

As noted above, only KEP work for kubeadm was done in 2021.
Some KEPs were retroactively updated to GA, since the work there was done without KEP updates.

   - Stable
     - [2500 - kubeadm: "join control-plane" workflow](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2500-kubeadm-join-control-plane-workflow)
     - [2501 - kubeadm: phase CLI support](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2501-kubeadm-phases-to-beta)
     - [2502 - kubeadm: "copy certs" on join](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2502-Certificates-copy-for-join-control-plane)
     - [2506 - kubeadm: remove "ClusterStatus"](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2506-Remove-ClusterStatus-from-kubeadm-config)
   - Beta
     - [1739 - kubeadm: customization with patches](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/1739-customization-with-patches)
     - [970 - kubeadm: config API to v1beta3](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/970-kubeadm-config)
   - Alpha
     - [2568 - kubeadm: non-root control plane](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/2568-kubeadm-non-root-control-plane)
   - Pre-alpha
     - NONE

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   In our survey results we found a few common areas where more help is needed across subprojects:
   - Code review
   - CI / e2e test monitoring and integration
   - Roadmap planning
   - Docs authoring and review

   All subprojects would appreciate more contributors, but here are a few highlights:
   - [etcdadm](https://github.com/kubernetes-etcdadm)
   - [cluster-addons](https://github.com/kubernetes-sigs/cluster-addons)
   - [cluster-api-provider-gcp](https://github.com/kubernetes-sigs/cluster-api-provider-gcp)
   - [kubeadm](https://github.com/kubernetes/kubeadm)

2. What metrics/community health stats does your group care about and/or measure?

   - We started collecting metrics from all of our subprojects using an
     [yearly survey](https://forms.gle/xZn8DXww4XxPsXvCA). It contains a number of questions
     related to project health, OWNERS files, contributor onboarding, etc. What we saw for this annual
     report is that our OWNERS files and SIG README.md are mostly up-to-date and that subprojects are doing
     what they can to onboard new contributors. We are considering preparing some actions to better educate
     our subprojects about onboarding contributors and graduating more OWNERS.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - As per the survey, we are seeing a good number of projects to keep their contributing
     documentation up-to-date. We are drafting action items for the leads to improve the understanding
     around [CONTRIBUTING.md] management. We have a good number of projects that have participated in LFX and GSoC.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Our subprojects have implemented their own contributing process and they have their own criteria
     for contributors to climb the ladders. We have not seen complains from new or existing
     contributors related to contributor guides. The SIG leads are open to help subproject leads with tips
     on this topic.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, as confirmed by our survey.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   In our survey results we found a few common areas where more help is needed across subprojects:
   - Code review
   - CI / e2e test monitoring and integration
   - Roadmap planning
   - Docs authoring and review

   These areas seems suitable for both full time and part time contributors.
   Replying to user questions on Slack and other communication channels is something
   that can be considered as full time support.
   Users/companies can reach out to the subproject leads if certain details
   are missing in the [CONTRIBUTING.md] file of a subproject.

## Membership

- Primary slack channel member count: 2868
- Primary mailing list member count: 1157
- Primary meeting attendee count (estimated, if needed): 5-10
- Primary meeting participant count (estimated, if needed): 5-10
- Unique reviewers for SIG-owned packages: 30+ <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 30+ <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

We do not count the overall group members because we have
[20 suprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects).
The responsibility of measuring membership is delegated to subproject leads.

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [cluster-api-provider-kubevirt](https://github.com/kubernetes/community/commit/ebeab03999e0406362670624fff5af5b1fcb08a4)
- [cluster-api-operator](https://github.com/kubernetes/community/commit/09ad92d62474a02da27381fefdea3c7acd78e244)

Retired in 2021:
- [cluster-api-provider-docker](https://github.com/kubernetes/community/commit/9b38820fa993a11afe8a90bbc7ea4268f85c5df2)

Continuing:
- [We have 20 suprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- NONE

Retired in 2021:
- [WG Component Standard](https://github.com/kubernetes/community/commit/a8fb89db5534f659e62e5c04528445b933d8e434)

Continuing:
- [WG Reliability](https://git.k8s.io/community/wg-reliability)
 ([2021 report](https://git.k8s.io/community/wg-reliability/annual-report-2021.md))

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Vince Prignano became a SIG chair, Timothy St. Clair moved to emeritus](https://groups.google.com/g/kubernetes-sig-cluster-lifecycle/c/LDF5udJnrzI/m/d5THJ-lsAAAJ)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cluster-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
