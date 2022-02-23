# 2021 Annual Report: SIG Cloud Provider

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Made progress toward removing cloud provider related code from the core Kubernetes code base.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Individual clouds have made significant progress on their cloud-controller-managers.
     - https://github.com/kubernetes/cloud-provider-aws
     - https://github.com/kubernetes/cloud-provider-gcp
     - https://github.com/kubernetes/cloud-provider-openstack
     - https://github.com/kubernetes/cloud-provider-vsphere
     - https://github.com/kubernetes/cloud-provider-alibaba-cloud
     - https://github.com/kubernetes-sigs/cloud-provider-azure

3. KEP work in 2021 (1.21, 1.22, 1.23):

   - Leader Migration for Controller Managers #2436
     - https://github.com/kubernetes/enhancements/issues/2436
   - Kubelet Credential Provider #2133
     - https://github.com/kubernetes/enhancements/issues/2133
   - KEP for adding webhook hosting capability to the CCM framework #2699
     - https://github.com/kubernetes/enhancements/issues/2699

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
   - Beta
     - [2436 - Leader Migration for Controller Managers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2436-controller-manager-leader-migration) - 1.22
   - Alpha
     - [2133 - Kubelet Credential Provider](https://git.k8s.io/community/$link/README.md) - 1.20
   - Pre-alpha
     - [2699 - KEP for adding webhook hosting capability to the CCM framework](https://github.com/cheftako/enhancements/blob/master/keps/sig-cloud-provider/2699-add-webhook-hosting-to-ccm/README.md)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   -

2. What metrics/community health stats does your group care about and/or
   measure?

   - We don't currently measure any statistics regarding community health.
   - In the future, we'd like to understand better what areas our community
     needs help in, i.e.  areas that are under-documented, areas with lots of
     bugs, etc.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group
specifically by pointing to activities or programs that provide useful context
or allow easy participation?

   - We are adding a new CONTRIBUTING.md in the same PR as this report.
   - As a follow-up, we need to reference CONTRIBUTING.md in our subprojects.

4. If your group has special training, requirements for reviewers/approvers, or
processes beyond the general [contributor guide], does your [CONTRIBUTING.md]
document those to help **existing** contributors grow throughout the
[contributor ladder]?

   - N/A (we don't have special requirements beyond the general [contributor
     guide].

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, we have contributors from Amazon, Google, VMware, and others.

6. Are there ways end users/companies can contribute that they currently are
not?  If one of those ways is more full time support, what would they work on
and why?

   - Yes, we could use help on the extraction/migration effort, including
     migrating tests out of the core Kubernetes repository and into each
     cloud-provider repository. See
     https://github.com/kubernetes/cloud-provider/issues/25 to get started

## Membership

- Primary slack channel member count: 942
- Primary mailing list member count: 240
- Primary meeting attendee count (estimated, if needed): 11
- Primary meeting participant count (estimated, if needed): 11
- Unique reviewers for SIG-owned packages: 47
- Unique approvers for SIG-owned packages: 38

See [here](https://gist.github.com/nckturner/cddd64bc1a56eaec836c07a24f7fecf4) for reviewers/approvers count method.

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
- [cloud-provider-extraction-migration](https://docs.google.com/document/d/1KLsGGzNXQbsPeELCeF_q-f0h0CEGSe20xiwvcR2NlYM/edit)
- [AWS Subproject Meeting](https://docs.google.com/document/d/1-i0xQidlXnFEP9fXHWkBxqySkXwJnrGJP9OGyP2_P14/edit#)
- [Azure Subproject Meeting](https://docs.google.com/document/d/1SpxvmOgHDhnA72Z0lbhBffrfe9inQxZkU9xqlafOW9k/edit)
- [IBM Subproject Meeting](https://docs.google.com/document/d/1qd_LTu5GFaxUhSWTHigowHt3XwjJVf1L57kupj8lnwg/edit)

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cloud-provider/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cloud-provider/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

