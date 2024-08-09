# 2023 Annual Report: SIG Contributor Experience

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?
    <!--
        Some example items that might be worth highlighting:
       - Major KEP advancement
       - Important initiatives that aren't tracked via KEPs
       - Paying down significant tech debt
       - Governance and leadership changes
    -->

    - [Onboarded New Chairs & Technical Leads](https://groups.google.com/g/kubernetes-sig-contribex/c/tNrxVKmNCQg/m/GaWWtvTQAgAJ)

    - GitHub Administration:
        - New members:
            - Onboarded 2 new GitHub Admins, and rotated 2 existing to Emeritus ([1](https://groups.google.com/g/kubernetes-sig-contribex/c/36Y3tT06FyQ/m/WYKUHcAaDAAJ), [2](https://groups.google.com/g/kubernetes-sig-contribex/c/PtLSBDS2yms/m/oFpc3PqUAQAJ))
            - [Onboarded 3 "New Membership Coordinators"](https://github.com/kubernetes/community/pull/7627), 1 active and [promoted as approver](https://github.com/kubernetes/org/pull/4719) and 2 dropped off.
        - [Kubernetes GitHub Org Membership Audit & Cleanup](https://groups.google.com/a/kubernetes.io/g/dev/c/Psiy9Iw-xCY/m/1-LVXnqRAAAJ)
        - [Updates to Kubernetes Org Membership Requirements](https://groups.google.com/a/kubernetes.io/g/dev/c/X6zDzaDRYaU/m/nVgDcG1rAwAJ)
        - [Onboarded SIG Etcd to Kubernetes GitHub Organisation & Prow](https://github.com/kubernetes/community/pull/7372) (major thanks to @cblecker, @mrbobbytables)
        - [Peribolos tool, extended to manage repo level permissions under GitHub teams](https://github.com/kubernetes/org/pull/2614)

    - Events:
        - 3 successful Kubernetes contributors summits:
            - [Europe, hosted in Amsterdam, Netherlands](https://www.kubernetes.dev/events/2023/kcseu/)
            - [China, hosted in Shanghai](https://www.kubernetes.dev/events/2023/kcscn/) (special thanks to @pacoxu for [leading the event](https://github.com/kubernetes/community/issues/7510))
            - [North America, hosted in Chicago, Illinois](https://www.kubernetes.dev/events/2023/kcsna/)

    - Mentoring:
        - 1 successful mentoring cohort:
            - [Kustomize maintainer training cohort](https://groups.google.com/a/kubernetes.io/g/dev/c/M5OphEVsv5o/m/zc6G4H15AAAJ)

    - Elections:
        - [Kubernetes Steering Committee Elections, for the 2023 cycle](https://groups.google.com/a/kubernetes.io/g/dev/c/bZpmcKA_900/m/EnlPmiDBAQAJ)
        - Facilitated [Kubernetes Code of Conduct Committee Election, for the 2023 cycle](https://groups.google.com/a/kubernetes.io/g/dev/c/YfyhGdTr3iA/m/AxT0OtIjBAAJ)

    - Contributor-comms:
        - [Subproject name change](https://github.com/kubernetes/community/commit/3613c1d2b33ab2283dc28f0215ff772668add39c): "Marketing Council" (earlier) -> "Contributor Communications" (Contributor Comms for short)
        - Got access to [Kubernetes Linkedin account](https://www.linkedin.com/company/kubernetes/)
        - Got access to Kubernetes X (earlier, Twitter) account – [@kubernetesio](https://x.com/kubernetesio)
        - LWKD (Last Week in Kubernetes Development) group became part of contributor-comms, onboarded 3 new dedicated maintainers [1](https://github.com/kubernetes/org/pull/4519) [2](https://github.com/kubernetes/org/pull/4818)
        - 5 SIG spotlight blogs (2/5 contributed by new contributors)
            - [Spotlight on SIG CLI](https://kubernetes.io/blog/2023/07/20/sig-cli-spotlight-2023/)
            - [Spotlight on SIG Testing](https://kubernetes.io/blog/2023/11/24/sig-testing-spotlight-2023/)
            - [Spotlight on SIG Architecture: Conformance](https://kubernetes.io/blog/2023/10/05/sig-architecture-conformance-spotlight-2023/)
            - [Spotlight on SIG Architecture: Production Readiness](https://kubernetes.io/blog/2023/11/02/sig-architecture-production-readiness-spotlight-2023/)
            - [Introducing SIG etcd](https://kubernetes.io/blog/2023/11/07/introducing-sig-etcd/)
        - New leadership roles in subproject:
            - Onboarded 1 [Comms Tech Lead](https://github.com/kubernetes/community/commit/ba3d34a001abd3d50004f87ae44385efa5b2695d)
            - Onboarded 1 [Social Media Coordinator shadow](https://github.com/kubernetes/community/pull/7438)


    - Slack moderation:
        - [Updates to Kubernetes core contributor slack channels](https://groups.google.com/a/kubernetes.io/g/dev/c/E1SeJvWQgzQ/m/EkN_w2E1AQAJ) - #kubernetes-new-contributors, #kubernetes-org-members



2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?


    - Mentoring subproject needs a dedicated leadership, responsible for running both (i) group mentoring cohorts for Kubernetes SIGs & WGs, (ii) external mentoring programs (LFX Mentorship, Google Summer of Code, Google Summer of Docs, Outreachy, etc)


    - Elections subproject is in need of a lead (someone preferably with Python programming experience), who will own & maintain the [election software](https://github.com/kubernetes/community/tree/master/elections#software), pick [election officers](https://github.com/kubernetes/community/tree/master/elections#recommending-election-officers), and  help groups in the Kubernetes project who wants to use the election software.

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- KubeCon + CloudNativeCon, Europe 2023
    - [From Automation to Community: A Deep Dive into SIG Contributor Experience](https://sched.co/1Hzcc)

- KubeCon + CloudNativeCon, North America 2023
    - [Inflections and Reflections from Kubernetes SIG ContribEx on Community Growth & Sustainability](https://sched.co/1R2ot)
    - [Keynote Kubernetes Project Updates (SIG Contributor Experience)](https://www.youtube.com/watch?v=Lg5WFV-BWnk&t=194s)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

    None


## [Subprojects](https://git.k8s.io/community/sig-contributor-experience#subprojects)


**New in 2023:**
  - [sigs-github-actions](https://git.k8s.io/community/sig-contributor-experience#sigs-github-actions)
**Continuing:**
  - community
  - community-management
  - contributor-comms
  - contributors-documentation
  - devstats
  - elections
  - events
  - github-management
  - mentoring
  - slack-infra

## [Working groups](https://git.k8s.io/community/sig-contributor-experience#working-groups)

   None

## Operational

Operational tasks in [sig-governance.md]:
- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [X] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-contributor-experience/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-contributor-experience/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
