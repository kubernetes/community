# 2022 Annual Report: SIG CLI

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Aggregated discovery along with API Machinery
   - kubectl shell completion reworked
   - Making server-side apply the default in kubectl KEP 3805
   - kubectl translation improvements KEP 3655
   - Improve kubectl plugins for subcommands like "create my-thing" KEP 3638
   - Kustomize had a 5.0 release that made it into v1.27! Headline feature: kustomize localize
   - Add subresource support to kubectl KEP 2590
   - kui had 2 major releases and is plugging away at refinements
   - krew has 213 plugins
   - Mentoring cohort https://github.com/kubernetes/community/issues/6665 which resulted in more approvers (Arda) and reviewers (Marly)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Ever running options and flags refactor
   - Cutting down noisy CI
   - Increased test coverage
   - Kustomize docs revamp

3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2551 - kubectl return code normalization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2551-return-code-normalization) - v1.24
    - [2590 -  Kubectl Subresource Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2590-kubectl-subresource) - v1.24
    - [3104 - Introduce kuberc](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3104-introduce-kuberc) - v1.25
    - [3515 - Kubectl Explain OpenAPIv3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3515-kubectl-explain-openapiv3) - v1.26
  - beta:
    - [1441 - kubectl debug](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/1441-kubectl-debug) - v1.24
    - [2227 - kubectl default container](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2227-kubectl-default-container) - v1.24
    - [1440 - Kubectl events](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/1440-kubectl-events) - v1.26


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [Optimizing kubectl memory usage](https://github.com/kubernetes/kubectl/issues/978).
   - Kustomize only has two owners: https://github.com/kubernetes-sigs/kustomize/blob/master/OWNERS_ALIASES#L4-L6
   - Our docs, which are joint for Kustomize and Kubectl, need some love. They are built off [cli-experimental](https://github.com/kubernetes-sigs/cli-experimental), are outdated and need SEO improvements. The sites aren't in the first several pages of Google results for "kustomize docs" / "kubectl docs". The donated kustomize.io and kubectl.io sites/domains need to be integrated as well.


2. What metrics/community health stats does your group care about and/or measure?

   - [Open untriaged issues and PRs](https://cli.triage.k8s.io/s/kubectl).
   - New reviewers and approvers added.
     - Last year we added 1 reviewer and 1 approver.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No this document is out of date and difficult to maintain. We need to update it and model it after some other groups.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We recently did a reviewer mentorship cohort. Unclear if we would do another. The group consisted mostly of students not very experienced with Golang.

5. Does the group have contributors from multiple companies/affiliations?

    - The leads all represent different companies.
    - Contributors from 28 companies, 9 of which had 10+ contributions.
      - [Data](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20CLI&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=now-1y%2Fy&to=now-1y%2Fy).

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Kustomize is in need of new maintainers. Ideally folks that depend on it daily and able to jump right in.
   - kubectl is a large tool and we need new folks to stick around and learn its different pieces.

## Membership

- Primary slack channel member count: 2361
- Primary mailing list member count: 425
- Primary meeting attendee count (estimated, if needed): 10+
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: 32
- Unique approvers for SIG-owned packages: 29

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-cli#subprojects)

**Continuing:**

  - cli-experimental
  - cli-sdk
  - cli-utils
  - krew
  - krew-index
  - krm-functions
  - kubectl
  - kui
  - kustomize

## [Working groups](https://git.k8s.io/community/sig-cli#working-groups)

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon EU 2022](https://youtu.be/2o7WDLiXrW4)
      - [KubeCon US 2022](https://youtu.be/BDZFtYUnmCw)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cli/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cli/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
