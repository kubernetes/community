# 2022 Annual Report: SIG API Machinery

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Storage layer test unification (https://github.com/kubernetes/kubernetes/issues/109831)
   - CEL for admission control to alpha1 (https://github.com/kubernetes/enhancements/issues/3488)
   - Aggregated Discovery to alpha (https://github.com/kubernetes/enhancements/issues/3352)
   - StorageVersions revived interest from sig-auth (https://github.com/kubernetes/enhancements/issues/2339)
   - APIServer identity, revived interest from sig-auth (https://github.com/kubernetes/enhancements/issues/1965)
   - CRD CEL validation to beta (https://github.com/kubernetes/enhancements/issues/2876)
   - Efficient watch resumption to stable (https://github.com/kubernetes/enhancements/issues/1904)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Contacting individuals to widen the reviewer pool (including, not limited to https://github.com/kubernetes/kubernetes/pull/113904, https://github.com/kubernetes/kubernetes/pull/113959 as examples)
   - Twice a week issue/PR triage meetings.
   - Storage layer test unification (https://github.com/kubernetes/kubernetes/issues/109831)


3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [3488 - CEL for Admission Control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3488-cel-admission-control) - v1.26
    - [3352 - Aggregated Discovery](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3352-aggregated-discovery) - v1.26
    - [3156 - HTTP3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3156-http3) - v1.24
  - beta:
    - [1965 - kube-apiserver identity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1965-kube-apiserver-identity) - v1.26
    - [2876 - CRD Validation Expression Language](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2876-crd-validation-expression-language) - v1.25
    - [2885 - Server Side Unknown Field Validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2885-server-side-unknown-field-validation) - v1.25
    - [2896 - OpenAPI V3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2896-openapi-v3) - 1.24
  - stable:
    - [1164 - Deprecate and remove SelfLink](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1164-remove-selflink) - v1.24
    - [1904 - Efficient watch resumption after kube-apiserver reboot](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1904-efficient-watch-resumption) - v1.24


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - There are some packages that API Machinery owns and come out usually in our triage meetings, and that we most likely don't know much about: this happens often when Kubernetes is upgrading libraries for example - dependency updates span multiple SIGs due to which the entire PR also comes under the purview of API Machinery. 
   - Technical support for triages could have improvements, for example sometimes we remove API Machinery in PRs, and every update to the issue / PR get's us re-tagged again.
   - The ecosystem of the different Kubernetes Clients that we own grows more or less organically. Client-go and
   Python-client are probably the bigger ones.
   - The subprojects that have a total reviewer + approver count <= 2 across all its OWNERS files are as follows:
      - https://github.com/kubernetes-client/c (count: 2)
      - https://github.com/kubernetes-client/go-base (count: 2)
      - https://github.com/kubernetes-client/perl (count: 1)
   - There also exist parts of a subproject that come under review/approver crunch. One example is the [`cacher`](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/apiserver/pkg/storage/cacher) component. Components like this come under a subproject with enough OWNERS but the particular component itself might be dependent on the expertise of a small subset of these OWNERS.

2. What metrics/community health stats does your group care about and/or measure?

On the technical health of the SIG, we look at the following metrics tracked in [Devstats](https://k8s.devstats.cncf.io/d/25/open-pr-age-by-repository-group?orgId=1&var-period=q&var-repogroup_name=SIG%20API%20Machinery&var-kind_name=All&from=1640991600000&to=1672527599000&var-repo_name=kubernetes%2Fkubernetes):
   - Ratio of open/close PRs
   - Ratio of open/close issues
   - Overall age of open issues
   - Number of active contributors to the SIG

On the inclusion health of the SIG, we look at:
   - Representation of diversity and of multiple companies in the SIG participants

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No, one can be created. The README points contributors to SIG meetings and triage meetings that happens twice a week.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - No, it can be improved.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes, there are contributors from [multiple companies](https://k8s.devstats.cncf.io/d/74/contributions-chart?var-period=m&var-metric=contributions&var-repogroup_name=SIG%20API%20Machinery&var-repo_name=kubernetes%2Fkubernetes&var-country_name=All&var-company_name=All&var-company=all).
   We see all sorts of contributions, varying from issues, to comments, to PRs, to designs, to sig meeting participation, and user-survey data.
   

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - There currently isn't a charted course for end users/companies to get more involved.
   - The most common, albeit ad-hoc, method has been showing up in the SIG meeting and starting/continuing discussions on the slack channel.
   - There is plenty of areas were the SIG could use help in reviewing and maintaining the code, but it's a complicated code base to ramp up. If people don't stick around long enough, the effort does not make sense.

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
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon + CloudNativeCon NA 2022 Maintainer Track](https://www.youtube.com/watch?v=oLbzn_hYd5E)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
