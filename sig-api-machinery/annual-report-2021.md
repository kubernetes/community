# 2021 Annual Report: SIG API Machinery

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Server Side Apply promoted to GA
   - Deprecated v1beta1 of Custom Resources and Webhooks (in favor of GA version)
   - API Priority and Fairness introduced v1beta2
   - A massive ammount of improvements in the API Expression WG

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We are evaluating long term the potential benefits that generics in go1.19 could provide.


3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

Need to pull from https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+label%3Asig%2Fapi-machinery+updated%3A%3E%3D2021-01-01+is%3Aclosed

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
   
   The SIG sponsors some working groups that are largely independent. 

   There are several areas where regularly the SIG becomes under pressure, especially closer to code freezes and the
   vast amount of code owned by API Machinery.

   The ecosystem of the different Kubernetes Clients that we own grows more or less organically. Client-go and
   Python-client are probably the bigger ones.

   There are some packages that API Machinery owns and come out usually in our triage meetings, and that we most likely
   don't know much about: this happens often when Kubernetes is upgrading libraries for example. 
   
   Technical support for triages could have improvements, for example sometimes we remove API Machinery in PRs, and every update to the issue / pr get's us re-        tagged again. 

2. What metrics/community health stats does your group care about and/or measure?

On the technical health of the SIG, we look at
   - the ratio of open/close PRs
   - the ratio of open/close Issues
   - overall age of open Issues
   - Number of active contributors to the sig

On the inclusion health of the SIG, we look at:
   - representation of diversity and of multiple companies in the sig participants

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No, it can be improved. Usually we find many contributors via slack, PRs, and issues.


4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - No, it can be improved. The main concern has always been the large ammount of investment that is required to ramp up new contributors

5. Does the group have contributors from multiple companies/affiliations?

   Yes, there are contributors from [multiple companies](https://k8s.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20API%20Machinery&var-country_name=All&var-company_name=All&var-company=all).
   We see all sorts of contributions, varying from issues, to comments, to PRs, to designs, to sig meeting participation,
   and user-survey data.
   
6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - This is a topic that appears frequently. API Machinery owns ~40% of the Kubernetes code base. In the past, we had some newcomers show up in SIG meetings with very concrete support questions. This is not the right forum, and that is not what we want to turn those meetings into. 
   - There is plenty of areas were the SIG could use help in reviewing and maintaining the code, but it's a complicated code base to ramp up. If people don't stick around long enough, the effort does not make sense.

## Membership

- Primary slack channel member count: 3,384
- Primary mailing list member count: 689
- Primary meeting attendee count (estimated, if needed): 30
- Primary meeting participant count (estimated, if needed): 10
- Unique reviewers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
The following [subprojects][subproject-definition] are owned by sig-api-machinery:
### component-base
- **Owners:**
  - [kubernetes-sigs/legacyflag](https://github.com/kubernetes-sigs/legacyflag/blob/master/OWNERS)
  - [kubernetes/component-base](https://github.com/kubernetes/component-base/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/component-base](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/component-base/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/component-base/version](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/component-base/version/OWNERS)
### control-plane-features
- **Owners:**
  - [kubernetes-sigs/kube-storage-version-migrator](https://github.com/kubernetes-sigs/kube-storage-version-migrator/blob/master/OWNERS)
  - [kubernetes-sigs/kubectl-check-ownerreferences](https://github.com/kubernetes-sigs/kubectl-check-ownerreferences/blob/master/OWNERS)
  - [kubernetes/kubernetes/pkg/controller/garbagecollector](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/garbagecollector/OWNERS)
  - [kubernetes/kubernetes/pkg/controller/namespace](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/namespace/OWNERS)
  - [kubernetes/kubernetes/pkg/controller/resourcequota](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/resourcequota/OWNERS)
  - [kubernetes/kubernetes/pkg/quota/v1](https://github.com/kubernetes/kubernetes/blob/master/pkg/quota/v1/OWNERS)
### idl-schema-client-pipeline
- **Owners:**
  - [kubernetes-client/gen](https://github.com/kubernetes-client/gen/blob/master/OWNERS)
  - [kubernetes-sigs/structured-merge-diff](https://github.com/kubernetes-sigs/structured-merge-diff/blob/master/OWNERS)
  - [kubernetes/code-generator](https://github.com/kubernetes/code-generator/blob/master/OWNERS)
  - [kubernetes/gengo](https://github.com/kubernetes/gengo/blob/master/OWNERS)
  - [kubernetes/kube-openapi](https://github.com/kubernetes/kube-openapi/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/code-generator](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/code-generator/OWNERS)
### json
- **Owners:**
  - [kubernetes-sigs/json](https://github.com/kubernetes-sigs/json/blob/main/OWNERS)
### kubernetes-clients
- **Owners:**
  - [kubernetes-client/c](https://github.com/kubernetes-client/c/blob/master/OWNERS)
  - [kubernetes-client/csharp](https://github.com/kubernetes-client/csharp/blob/master/OWNERS)
  - [kubernetes-client/go-base](https://github.com/kubernetes-client/go-base/blob/master/OWNERS)
  - [kubernetes-client/go](https://github.com/kubernetes-client/go/blob/master/OWNERS)
  - [kubernetes-client/haskell](https://github.com/kubernetes-client/haskell/blob/master/OWNERS)
  - [kubernetes-client/java](https://github.com/kubernetes-client/java/blob/master/OWNERS)
  - [kubernetes-client/javascript](https://github.com/kubernetes-client/javascript/blob/master/OWNERS)
  - [kubernetes-client/perl](https://github.com/kubernetes-client/perl/blob/master/OWNERS)
  - [kubernetes-client/python-base](https://github.com/kubernetes-client/python-base/blob/master/OWNERS)
  - [kubernetes-client/python](https://github.com/kubernetes-client/python/blob/master/OWNERS)
  - [kubernetes-client/ruby](https://github.com/kubernetes-client/ruby/blob/master/OWNERS)
  - [kubernetes-sigs/clientgofix](https://github.com/kubernetes-sigs/clientgofix/blob/master/OWNERS)
  - [kubernetes/client-go](https://github.com/kubernetes/client-go/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/client-go](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/OWNERS)
### server-api-aggregation
- **Owners:**
  - [kubernetes/kube-aggregator](https://github.com/kubernetes/kube-aggregator/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/kube-aggregator](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kube-aggregator/OWNERS)
### server-binaries
- **Owners:**
  - [kubernetes/kubernetes/cmd/cloud-controller-manager](https://github.com/kubernetes/kubernetes/blob/master/cmd/cloud-controller-manager/OWNERS)
  - [kubernetes/kubernetes/cmd/kube-apiserver](https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-apiserver/OWNERS)
  - [kubernetes/kubernetes/cmd/kube-controller-manager](https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-controller-manager/OWNERS)
  - [kubernetes/kubernetes/pkg/kubeapiserver](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubeapiserver/OWNERS)
  - [kubernetes/kubernetes/pkg/master](https://github.com/kubernetes/kubernetes/blob/master/pkg/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/controller-manager](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/controller-manager/OWNERS)
### server-crd
- **Owners:**
  - [kubernetes/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/apiextensions-apiserver](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiextensions-apiserver/OWNERS)
### server-frameworks
- **Owners:**
  - [kubernetes/apiserver](https://github.com/kubernetes/apiserver/blob/master/OWNERS)
  - [kubernetes/controller-manager](https://github.com/kubernetes/controller-manager/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/apiserver](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/controller-manager](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/controller-manager/OWNERS)
### server-sdk
- **Owners:**
  - [kubernetes-sigs/apiserver-builder-alpha](https://github.com/kubernetes-sigs/apiserver-builder-alpha/blob/master/OWNERS)
  - [kubernetes-sigs/apiserver-runtime](https://github.com/kubernetes-sigs/apiserver-runtime/blob/master/OWNERS)
  - [kubernetes-sigs/controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/blob/master/OWNERS)
  - [kubernetes-sigs/controller-tools](https://github.com/kubernetes-sigs/controller-tools/blob/master/OWNERS)
  - [kubernetes-sigs/kubebuilder-declarative-pattern](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/blob/master/OWNERS)
  - [kubernetes-sigs/kubebuilder-release-tools](https://github.com/kubernetes-sigs/kubebuilder-release-tools/blob/master/OWNERS)
  - [kubernetes-sigs/kubebuilder](https://github.com/kubernetes-sigs/kubebuilder/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/sample-apiserver](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/sample-apiserver/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/sample-controller](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/sample-controller/OWNERS)
  - [kubernetes/sample-apiserver](https://github.com/kubernetes/sample-apiserver/blob/master/OWNERS)
  - [kubernetes/sample-controller](https://github.com/kubernetes/sample-controller/blob/master/OWNERS)
- **Contact:**
  - [Mailing List](https://groups.google.com/forum/#!forum/kubebuilder)
### universal-machinery
- **Owners:**
  - [kubernetes/apimachinery](https://github.com/kubernetes/apimachinery/blob/master/OWNERS)
  - [kubernetes/kubernetes/staging/src/k8s.io/apimachinery](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/OWNERS)
### yaml
- **Owners:**
  - [kubernetes-sigs/yaml](https://github.com/kubernetes-sigs/yaml/blob/master/OWNERS)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
The following working groups are sponsored by sig-api-machinery:

- [WG API Expression](https://github.com/kubernetes/community/tree/master/wg-api-expression)
- [WG Multitenancy](https://github.com/kubernetes/community/tree/master/wg-multitenancy)
- [WG Structured Logging](https://github.com/kubernetes/community/tree/master/wg-structured-logging)


## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x]  Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Kubecon NA 2021](https://www.youtube.com/watch?v=oiC2w1PVjrQ)
      - [Community update](https://docs.google.com/presentation/d/1UWRaMVtTD3yVhJ3MGBpt7LRIaRHTaQZoGlDT7Bl7jLE/edit#slide=id.g401c104a3c_0_0)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

