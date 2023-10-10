# 2022 Annual Report: SIG Cloud Provider

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   -  In 2022, SIG Cloud Provider made progress in the extraction migration process with the [Openstack in-tree provider removal](https://github.com/kubernetes/kubernetes/pull/67782).
   - We merged the [KEP for CCM webhooks](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2699-add-webhook-hosting-to-ccm), finalizing the design. This was to solve a problem around persistent volume labeling and to allow for further customization.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Michael McCune (@elmiko) began [experimental work on e2e test refactors](https://hackmd.io/@elmiko/BJGn1SQU3), building a set of generic ccm tests, based on the current tests, that will be available for all providers to utilize to exercise their own provider. We currently have limited tests for selected providers, and want to expand beyond that.


3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2699 - Add webhook hosting to CCM.](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2699-add-webhook-hosting-to-ccm) - v1.26
  - stable:
    - [1959 - Service Type=LoadBalancer Class Field](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/1959-service-lb-class-field) - v1.24
    - [2436 - Controller Manager Leader Migration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2436-controller-manager-leader-migration) - v1.24
    - [2133: Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers) - v1.26


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - While we always need stronger connections to the individual cloud providers and better representation across the community, we specifically could use help with the individual cloud provider implementations. Most providers have a relatively small subset of active maintainers compared to other SIGs in Kubernetes overall. (The OWNERS files appear to need an audit.)

2. What metrics/community health stats does your group care about and/or measure?

   - SIG Cloud Provider monitors the number and variety of cloud providers who are attending SIG meetings and providing feedback about their concerns with issues on their infrastructures or with difficulties implementing SIG-owned keps and related technical assistance. We measure this through attendance and agenda items at our bi-weekly SIG meetings.


3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Given the vendor-specific nature of the out-of-tree cloud provider code bases, when new potential contributors ask about opportunities we recommend they look at the provider for the cloud they are most interested in using. Pairing with an established contributor or partnering on a bugfix is a good approach. We welcome documentation contributions for topics like how CCMs work (and also docs about release engineering, QA, and testing).

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Our SIG has no special requirements beyond the general contributor guide.


5. Does the group have contributors from multiple companies/affiliations?

   - Yes, SIG Cloud Provider has contributors from multiple clouds and large-scale vendors. The nature of cross-cloud and cross-vendor collaboration means that it’s necessary to have contributors from the major areas of cloud effort in the ecosystem.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Interested users who work at end user organizations can contribute to the cloud provider work via testing, bug reports, and feature requests. They can also contribute to the specific individual provider codebases as relevant to their needs.

## Membership

Note: Estimated numbers as of July 2023 when data was collected; 2022 numbers not available.


- Primary slack channel member count: 1,157
- Primary mailing list member count: 265
- Primary meeting attendee count (estimated, if needed): ~6-8 (estimate/average)
- Primary meeting participant count (estimated, if needed): ~5-6 (estimate/average)
- Unique reviewers for SIG-owned packages: 109
- Unique approvers for SIG-owned packages: 116

Note: for 2022, we retrieved reviewer and approver numbers from all the subprojects listed in https://github.com/kubernetes/community/tree/master/sig-cloud-provider#subprojects which means this figure is not comparable to the limited subset counted in 2021. However, a cursory overview of the OWNERS files reveals that they change infrequently and likely need an audit.

## [Subprojects](https://git.k8s.io/community/sig-cloud-provider#subprojects)



**New in 2022:**

  - provider-oci

**Continuing:**

  - cloud-provider-extraction-migration
  - kubernetes-cloud-provider
  - provider-alibaba-cloud
  - provider-aws
  - provider-azure
  - provider-baiducloud
  - provider-gcp
  - provider-huaweicloud
  - provider-ibmcloud
  - provider-openstack
  - provider-vsphere


## [Working groups](https://git.k8s.io/community/sig-cloud-provider#working-groups)


**Continuing:**

 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
    - KubeCon EU 2022 - [SIG Cloud Provider: Portable K8s Across all Clouds, Roadmap and Updates - Nick Turner, Amazon & Steve Wong, VMware](https://sched.co/ytow)
    - KubeCon NA 2022 - [SIG Cloud Provider Update - Michael McCune, Red Hat & Bridget Kromhout, Microsoft](https://www.youtube.com/watch?v=jnc4Eysh1g0) & [slides](https://sched.co/1C89O)
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cloud-provider/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cloud-provider/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
