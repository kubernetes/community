# 2021 Annual Report: SIG Testing

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - kubetest2 is maturing thanks to some great contributions.
   - Automated secret syncing for ProwJob secrets.
   - Prow work for other open source projects and communities.
     - GitHub Apps support.
     - Improved job config validation (strict field checks, build cluster existence).
     - Improved In-repo config support and performance.
     - Support for config file sharding to better manage approval permissions.
     - New monitoring stack solution that doesn’t rely on Grafana (GKE Workload Metrics + Cloud Monitoring).
     - OSS-Fuzz integration.
     - Private repo multitenancy (multiple private front ends).
   - kubernetes/kubernetes Bazel removal GA
   - Bazel removal from kubernetes/test-infra almost complete

2. What initiatives are you working on that aren't being tracked in KEPs?

    Most of our work is not tracked via KEPs. We need to do better with bringing issues to groups and individuals to gain consensus. The current KEP process is better suited for end user features and speaks to end users more than contributors.

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [KEP 2420 - Reducing Kubernetes Build Maintenance](https://github.com/kubernetes/enhancements/issues/2420) - 1.23.stable
   - Beta
     - [KEP 2539 - Continuously Deploy K8s Prow](https://github.com/kubernetes/enhancements/issues/2539) - 1.21.beta
     - [KEP 2464 - kubetest2 CI migration](https://github.com/kubernetes/enhancements/issues/2464) - 1.23.beta

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [Boskos](https://github.com/kubernetes-sigs/boskos/blob/master/OWNERS)
     - Stable but if it breaks we are in a bad place for the project's CI.
   - [K8s GSM Tool](https://github.com/kubernetes-sigs/k8s-gsm-tools/blob/master/OWNERS)
     - This was an intern project that did not gain traction. Consider archiving and discussing with Aaron.
   - [kubetest2](https://github.com/kubernetes-sigs/kubetest2/blob/master/OWNERS)
     - Ben is the only active reviewer and approver.
   - [Prow](https://github.com/kubernetes/test-infra/tree/master/prow/OWNERS)
     - Google cannot continue to maintain https://monitoring.prow.k8s.io due to Grafana license change. We have switched to using Google Cloud Monitoring, but cannot make the dashboards publicly visible. We need SIG-k8s-infra to take over here or we risk losing public monitoring dashboards soon. The boskos metrics dashboard is probably the most notable loss. Currently the grafana instance is frozen.
   - [Triage](https://github.com/kubernetes/test-infra/blob/master/triage/OWNERS) + [Kettle](https://github.com/kubernetes/test-infra/blob/master/kettle/OWNERS)
     - Not a top priority but is used by many when debugging flake.

2. What metrics/community health stats does your group care about and/or measure?

   - Reviewers and approvers

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No this file does not exist.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   -

5. Does the group have contributors from multiple companies/affiliations?

   - Yes Google and Red Hat are the majority OWNERS but we see folks from ii.co, VMware, and others.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - We need folks to show up and stick around to climb the contributor ladder.

## Membership

- Primary slack channel member count: 2117
- Primary mailing list member count: 341
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
- [boskos](https://git.k8s.io/community/sig-testing#boskos)
- [e2e-framework](https://git.k8s.io/community/sig-testing#e2e-framework)
- [kind](https://git.k8s.io/community/sig-testing#kind)
- [kubetest2](https://git.k8s.io/community/sig-testing#kubetest2)
- [prow](https://git.k8s.io/community/sig-testing#prow)

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

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-testing/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-testing/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

