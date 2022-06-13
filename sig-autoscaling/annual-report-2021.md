# 2021 Annual Report: SIG Autoscaling

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - [Promotion of autoscaling v2 API to GA.](https://github.com/kubernetes/enhancements/pull/2703)
   - [Continued improvements of VPA via community feedback](https://github.com/kubernetes/autoscaler/issues/3913)
   - [Improved extensibility of Cluster Autoscaler](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/plugable-provider-grpc.md)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Continuing to improve and extend functionality of owned Subprojects
   - Increasing out of band extensibility of projects via mechanisms such as gRPC providers for CA

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [2702 - Graduate v2beta2 Autoscaling API to GA](https://git.k8s.io/enhancements/keps/sig-autoscaling/2702-graduate-hpa-api-to-GA/README.md) - $milestone.stable

   - Subproject enhancements:
     - [Support Customized Recommenders for Vertical Pod Autoscalers](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/3919-customized-recommender-vpa)
     - [MinReplicas per VPA object](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/4566-min-replicas#kep-4566-minreplicas-per-vpa-object)
     - [https://github.com/kubernetes/autoscaler/blob/68c984472acce69cba89d96d724d25b3c78fc4a0/cluster-autoscaler/proposals/plugable-provider-grpc.md](https://github.com/kubernetes/autoscaler/blob/68c984472acce69cba89d96d724d25b3c78fc4a0/cluster-autoscaler/proposals/plugable-provider-grpc.md)
     - [Expander Plugin over gRPC](https://github.com/kubernetes/autoscaler/blob/66af6d1339f86e87a37d5f505109b59c729de198/cluster-autoscaler/proposals/expander-plugin-grpc.md)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Code contributors and reviewers
   - Issue triage/response
   - Whilst there are no areas with explicitly 2 or few owners, there are a number of areas where we have this few subject matter experts, acting as a significant bottleneck for significant changes. These currently include the core Cluster Autoscaler code, as well as the HPA controller.

2. What metrics/community health stats does your group care about and/or measure?

   - Currently lacking in measurement of community health stats
   - Care about growing number of contributors
   - Would like to begin measuring time to merge of community PRs

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Not currently

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A

5. Does the group have contributors from multiple companies/affiliations?

   - Yes - though the core approvers are almost all from a single company (Google). Other contributors/code owners are from a variety of companies, including end users (Datadog, Airbnb, Skyscanner...) and vendors (Microsoft, Red Hat, Amazon...), particularly for Cluster Autoscaler cloud provider implementations.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Increased support for some cloud provider implementations for the Cluster Autoscaler. See e.g. [issues](https://github.com/kubernetes/autoscaler/issues?q=is%3Aopen+is%3Aissue+label%3Aarea%2Fprovider%2Faws+) and [pull requests](https://github.com/kubernetes/autoscaler/pulls?q=is%3Aopen+is%3Apr+label%3Aarea%2Fprovider%2Fhetzner) which can be queried by cloud provider for starting points. (Modify the label filter as appropriate from the above links.)
   - Responding to issues - the SIG currently lacks the capacity to respond to all issues raised

## Membership

- Primary slack channel member count: 3824
- Primary mailing list member count: 346
- Primary meeting attendee count (estimated, if needed): ~10
- Primary meeting participant count (estimated, if needed): ~6
- Unique reviewers for SIG-owned packages:
  - Under k/k: 4
  - SIG Subprojects Core Code: 4 to 8 depending on the project
  - Cluster Autoscaler including cloud provider implementations: ~40<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages:
  - Under k/k: 4
  - SIG Subprojects Core Code: 4 to 8 depending on the project
  - Cluster Autoscaler including cloud provider implementations: ~40<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

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

- [horizontal-pod-autoscaler](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md#horizontal-pod-autoscaler)
- [vertical-pod-autoscaler](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md#vertical-pod-autoscaler)
- [cluster-autoscaler](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md#cluster-autoscaler)
- [addon-resizer](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md#addon-resizer)
- [scale-client](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md#scale-client)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:

- [WG-Batch](https://github.com/kubernetes/community/tree/master/wg-batch)

Retired in 2021:

- None

Continuing:

- None

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
      - [Kubecon EU 2021](https://www.youtube.com/watch?v=odxPyW_rZNQ)
      - [Kubecon NA 2021](https://www.youtube.com/watch?v=L4d7K83vq_0)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-autoscaling/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-autoscaling/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
