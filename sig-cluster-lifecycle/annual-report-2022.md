# 2022 Annual Report: SIG Cluster Lifecycle

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - We migrated the default image repository in kubeadm and some of our other k8s deployment
   tools from k8s.gcr.io to registry.k8s.io.
   - kubeadm initiated work on migrating its etcd bootstrapping to use ["learner mode"](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/3614-etcd-learner-mode)
   - We finilized migrating subprojects away from the legacy ["master" label/taint](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-taint)
   - We completed the setup of the ["kubernetes-sigs/cluster-api-addon-provider-helm"](https://github.com/kubernetes-sigs/cluster-api-addon-provider-helm/blob/main/OWNERS) subproject providing an option for Cluster API users to manage addons, and most specifically CPI being moved out of tree.
   - We created the ["logical-cluster"](https://github.com/kubernetes-sigs/logical-cluster/blob/main/OWNERS) subproject to help in cases of CAPI cluster fleets.
   - We worked tighly with test-infra supporting the migration to registry.k8s.io as well as reviewing test footprint of Cluster API.
   - We started a Cluster API release team modeled from the Kubernetes release team, and this allowed the introduction of a [quarterly release cadence](https://github.com/kubernetes-sigs/cluster-api/tree/main/docs/release) for the entire 2023.
  
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

3. KEP work in 2022 (v1.24, v1.25, v1.26):

As noted above, only kubeadm relies on the KEP process; work completed in 2022 includes.

   - [KEP-2915](https://features.k8s.io/2915): Replace usage of the kubelet-config-x.y naming in kubeadm, using UnversionedKubeletConfigMap alpha:v1.23; beta: v1.24; GA: v1.25; Removed: v1.26. This is not included probably because it started in 2021.
   - [KEP-1739](https://features.k8s.io/1739): kubeadm customization with patches: kubeadm: add support for patching a "kubeletconfiguration" target #110405 v1.25. This is similar that the KEP started before 2022, and the only kubeletconfiguration patching was supported in 2022.
   - Some KEPs were retroactively updated to GA, since the work there was done without KEP updates.

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [cluster-api-addon-provider-helm](https://github.com/kubernetes-sigs/cluster-api-addon-provider-helm/blob/main/OWNERS)
   - [etcdadm](https://github.com/kubernetes-sigs/etcdadm)
   - [cluster-addons](https://github.com/kubernetes-sigs/cluster-addons)

2. What metrics/community health stats does your group care about and/or measure?

   - We started collecting metrics from all of our subprojects using an
     [yearly survey](https://forms.gle/xZn8DXww4XxPsXvCA). It contains a number of questions
     related to project health, OWNERS files, contributor onboarding, etc. What we saw for this annual
     report is that our OWNERS files and SIG README.md are mostly up-to-date and that subprojects are doing
     what they can to onboard new contributors. We are considering preparing some actions to better educate
     our subprojects about onboarding contributors and graduating more OWNERS. We did not do a survey
     for 2022.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - As per the 2021 [survey](https://forms.gle/xZn8DXww4XxPsXvCA), we are seeing a good number of projects to keep their contributing
     documentation up-to-date. We are drafting action items for the leads to improve the understanding
     around [CONTRIBUTING.md] management. We have a few projects that have participated in LFX and GSoC
     (kOps, cluster-addons, kubeadm).

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Our subprojects have implemented their own contributing process and they have their own criteria
     for contributors to climb the ladders. We have not seen complains from new or existing
     contributors related to contributor guides. The SIG leads are open to help subproject leads with tips
     on this topic.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, as confirmed by our [survey](https://forms.gle/xZn8DXww4XxPsXvCA); also following
     dashboard are confirming above data:
     - https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Cluster%20Lifecycle&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&viewPanel=1&from=1641016800000&to=1672466400000
     - https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Cluster%20Lifecycle%20(Cluster%20API)&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&viewPanel=1&from=1641016800000&to=1672466400000

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   In our [survey](https://forms.gle/xZn8DXww4XxPsXvCA) results we found a few common areas where more help is needed across subprojects:
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

- Primary slack channel member count: 3116
- Primary mailing list member count: 1100+
- Primary meeting attendee count (estimated, if needed): 5-10
- Primary meeting participant count (estimated, if needed): 5-10
- Unique reviewers for SIG-owned packages: 30+ <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 30+ <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects)



**New in 2022:**

  - cluster-api-operator
  - cluster-api-provider-cloudstack
  - cluster-api-addon-provider-helm
  - kOps

**Retired in 2022:**

  - kops
  - kube-up

**Continuing:**

  - cluster-addons
  - cluster-api
  - cluster-api-provider-aws
  - cluster-api-provider-azure
  - cluster-api-provider-digitalocean
  - cluster-api-provider-gcp
  - cluster-api-provider-ibmcloud
  - cluster-api-provider-kubemark
  - cluster-api-provider-kubevirt
  - cluster-api-provider-nested
  - cluster-api-provider-openstack
  - cluster-api-provider-packet
  - cluster-api-provider-vsphere
  - etcdadm
  - image-builder
  - kubeadm
  - kubespray
  - minikube


## [Working groups](https://git.k8s.io/community/sig-cluster-lifecycle#working-groups)


**Continuing:**

 - Reliability

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - We had both [KubeCon EU](https://www.youtube.com/watch?v=9H8flXm_lKk) and [KubeCon NA](https://www.youtube.com/watch?v=0Zo0cWYU0fM) sessions in 2022

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cluster-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
