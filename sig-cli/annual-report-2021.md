# 2021 Annual Report: SIG CLI

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

    - [kubectl events alpha command](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cli/1440-kubectl-events/README.md).
    - [KRM Functions subproject started](https://github.com/kubernetes-sigs/krm-functions-registry).
    - New changes to leadership.
        - [@KnVerey](https://github.com/knverey) brought on as new Co-Chair and Tech Lead.
        - [@soltysh](https://github.com/soltysh) stepped down from Co-Chair to focus on Tech Lead.
        - [@pwittrock](https://github.com/pwittrock) moved to emeritus.
        - [@monopole](https://github.com/monopole) moved to emeritus for Kustomize.
    - [Started a new monthly Kustomize bug scrub](https://github.com/kubernetes/community/tree/master/sig-cli#meetings).
    - [Upgraded the version of Kustomize that ships with kubectl](https://github.com/kubernetes/kubernetes/pull/98946).
    - [Implemented native Go shell completions](https://github.com/kubernetes/kubernetes/pull/96087).
    - [Replicated](https://www.replicated.com/) donated [kubectl.io](https://kubectl.io) and [kustomize.io](https://kustomize.io) to the project.
    - [IBM](https://ibm.com) donated the [Kui](https://github.com/kubernetes-sigs/kui) project.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - [The Kustomize Roadmap](https://github.com/kubernetes-sigs/kustomize/blob/master/ROADMAP.md).
   - [Refactoring old kubectl commands](https://github.com/kubernetes/kubectl/issues/1046).

3. KEP work in 2021 (1.x, 1.y, 1.z):

    - Stable
        - [KEP-555 - Server-side apply](https://github.com/kubernetes/enhancements/issues/555) - 1.22.stable
    - Beta
        - [KEP-1441 - kubectl debug](https://github.com/kubernetes/enhancements/issues/1441) - 1.20.beta, continued to evolve the beta through the year
        - [KEP-859 - kubectl command metadata in http request headers](https://github.com/kubernetes/enhancements/issues/859) - 1.22.beta
    - Alpha
        - [KEP-1440 - kubectl events](https://github.com/kubernetes/enhancements/issues/1440) - 1.23.alpha
        - [KEP-2227 - Default container annotation to be used by kubectl](https://github.com/kubernetes/enhancements/issues/2227) - 1.21.alpha
    - Pre-alpha
        - [KEP-2985 - Public KRM functions registry](https://github.com/kubernetes/enhancements/issues/2985)
        - [KEP-2953 - Kustomize Plugin Graduation](https://github.com/kubernetes/enhancements/issues/2953)
    - Rejected
        - [KEP-2229 - Use XDG Base Directory Specification](https://github.com/kubernetes/enhancements/issues/2229)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [Optimizing kubectl memory usage](https://github.com/kubernetes/kubectl/issues/978).
   - [Kustomize only has two maintainers](https://github.com/kubernetes-sigs/kustomize/blob/master/OWNERS_ALIASES#L4-L6).
   - Our docs, which are joint for Kustomize and Kubectl, need some love. They are built off [cli-experimental](https://github.com/kubernetes-sigs/cli-experimental), are outdated and need SEO improvements. The sites aren't in the first several pages of Google results for "kustomize docs" / "kubectl docs". The donated kustomize.io and kubectl.io sites/domains need to be integrated as well.

2. What metrics/community health stats does your group care about and/or measure?

   - [Open untriaged issues and PRs](https://cli.triage.k8s.io/s/kubectl).
   - New reviewers and approvers added.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No this document is out of date and difficult to maintain. We need to update it and model it after some other groups.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We’re working with SIG-Apps on starting a monthly review club, details to be announced soon.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes all of the leads are from different companies and we see a spread of contributions from other companies. That said we would love to see further investment.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Kustomize is in need of new maintainers. Ideally folks that depend on it daily and able to jump right in.
   - kubectl is a large tool and we need new folks to stick around and learn its different pieces.

## Membership

- Primary slack channel member count: 2014
- Primary mailing list member count: 329
- Primary meeting attendee count (estimated, if needed): 10+
- Primary meeting participant count (estimated, if needed): 6
- Unique reviewers for SIG-owned packages: 9
- Unique approvers for SIG-owned packages: 13

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [KRM Functions](https://github.com/kubernetes/community/blob/master/sig-cli/README.md#krm-functions)

Continuing:
- [kui](https://git.k8s.io/community/sig-cli#kui)
  - Open sourced "terminal as notebook" support. Example: https://playground.guidebooks.dev/
  - 1500 -> 2100 stars
  - User study of user preference for kubectl versus web consoles with [@algebrot](https://github.com/algebrot) (Cora Coleman). Summary of findings:
    > We found that among a survey of 60 cloud developers (beginner, intermediate, expert), all strongly prefer using the CLI modality for completing CRUD and debugging tasks and only the intermediates and experts prefer using a different tool modality (web console) for monitoring tasks. Similarly we observed strong preference for the CLI among four developers that completed a task using the CLI and the web console
- [cli-experimental](https://git.k8s.io/community/sig-cli#cli-experimental)
- [cli-sdk](https://git.k8s.io/community/sig-cli#cli-sdk)
- [cli-utils](https://git.k8s.io/community/sig-cli#cli-utils)
- [krew](https://git.k8s.io/community/sig-cli#krew)
- [krew-index](https://git.k8s.io/community/sig-cli#krew-index)
  - Steady growth with 42 new plugins added
- [kubectl](https://git.k8s.io/community/sig-cli#kubectl)
- [kustomize](https://git.k8s.io/community/sig-cli#kustomize)

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
  - [KubeCon EU 2021](https://youtu.be/f_P-wKjXrTs)
  - [KubeCon NA 2021](https://youtu.be/2o7WDLiXrW4)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cli/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cli/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

