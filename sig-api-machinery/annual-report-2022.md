# 2022 Annual Report: SIG API Machinery

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Storage layer test unification (https://github.com/kubernetes/kubernetes/issues/109831)
   -
   -

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Storage layer test unification (https://github.com/kubernetes/kubernetes/issues/109831)
   -
   -



3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [3156 - HTTP3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3156-http3) - v1.24
  - beta:
    - [1965 - kube-apiserver identity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1965-kube-apiserver-identity) - v1.26
    - [2876 - CRD Validation Expression Language](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2876-crd-validation-expression-language) - v1.25
  - stable:
    - [1164 - Deprecate and remove SelfLink](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1164-remove-selflink) - v1.24
    - [1904 - Efficient watch resumption after kube-apiserver reboot](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption) - v1.24


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

   - No, it can be improved. The README points contributors to SIG meetings and triage meetings that happens twice a week.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - No, it can be improved.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, there are contributors from [multiple companies](https://k8s.devstats.cncf.io/d/74/contributions-chart?var-period=m&var-metric=contributions&var-repogroup_name=SIG%20API%20Machinery&var-repo_name=kubernetes%2Fkubernetes&var-country_name=All&var-company_name=All&var-company=all).
   We see all sorts of contributions, varying from issues, to comments, to PRs, to designs, to sig meeting participation,
   and user-survey data.
   

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count: 3912
- Primary mailing list member count: 771
- Primary meeting attendee count (estimated, if needed): 30 <!-- carried over from last year -->
- Primary meeting participant count (estimated, if needed): 10 <!-- carried over from last year -->
- Unique reviewers for SIG-owned packages: 111 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 94 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-api-machinery#subprojects)



**Continuing:**

  - component-base
  - control-plane-features
  - idl-schema-client-pipeline
  - json
  - kubernetes-clients
  - server-api-aggregation
  - server-binaries
  - server-crd
  - server-frameworks
  - server-sdk
  - universal-machinery
  - yaml


## [Working groups](https://git.k8s.io/community/sig-api-machinery#working-groups)


**Continuing:**

 - API Expression
 - Multitenancy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - KubeCon NA and EU 22 maintainer track
      - 

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
