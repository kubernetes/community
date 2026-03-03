# 2025 Annual Report: SIG etcd

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

- Leadership Change – [Onboarded @ivanvc as new chair](https://github.com/kubernetes/org/pull/5491).
- Leadership Change – [@siyuanfoundation replaced @wenjiaswe as chair](https://groups.google.com/g/etcd-dev/c/1I9Q8i97Lts).
- Leadership Change – [Onboarded @fuweid as Technical Leader](https://groups.google.com/g/etcd-dev/c/h-frvWLYtKw).

- Released etcd [v3.6.0](https://etcd.io/blog/2025/announcing-etcd-3.6/) and seven patches.
- Delivered 8 patch releases for the 3.5 stable minor.
- Delivered 4 patch releases for the 3.4 stable minor.

- [Collaborated with Antithesis](https://etcd.io/blog/2025/autonomus_testing_with_antithesis/) to start a strategic alliance to improve the execution of etcd robustness tests.
- Identified and fixed multiple fail scenarios when updating from 3.5 to 3.6, including the [replacement of the v2store with v3store](https://etcd.io/blog/2025/upgrade_from_3.5_to_3.6_issue/), [during rolling replacement updates](https://etcd.io/blog/2025/upgrade_from_3.5_to_3.6_issue_followup/), and with [possible zombie members](https://etcd.io/blog/2025/zombie_members_upgrade/).

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- [etcdlabs](https://github.com/etcd-io/etcdlabs)
- [gofail](https://github.com/etcd-io/gofail)
- [govanityurls](https://github.com/etcd-io/govanityurls)
- [jetcd](https://github.com/etcd-io/jetcd)
- [protodoc](https://github.com/etcd-io/protodoc)

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- KubeCon EU 2025: [etcd v3.6.0 and etcd-operator v0.1.0](https://www.youtube.com/watch?v=_xoDbpm-Qks).
- KubeCon NA 2025: [etcd v3.6 and Beyond + etcd-operator Updates](https://www.youtube.com/watch?v=tF7UOwhgetU).
- [etcd mentorship program cohort 2](https://groups.google.com/a/kubernetes.io/g/dev/c/JbdZJVjmkwQ/m/4VW6bmhNAAAJ).

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

- None.

## [Subprojects](https://git.k8s.io/community/sig-etcd#subprojects)


**New in 2025:**
  - etcd-operator
**Continuing:**
  - auger
  - bbolt
  - cetcd
  - dbtester
  - discovery.etcd.io
  - discoveryserver
  - etcd
  - etcd-manager
  - etcd-play
  - etcdlabs
  - gofail
  - govanityurls
  - jetcd
  - maintainers
  - protodoc
  - raft
  - website
  - zetcd

## [Working groups](https://git.k8s.io/community/sig-etcd#working-groups)

**Continuing:**
 - etcd Operator

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-etcd/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-etcd/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
