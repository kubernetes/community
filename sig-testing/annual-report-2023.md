# 2023 Annual Report: SIG Testing

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

Crossover with SIG K8s Infra. Migration of CI (mainly jobs, AWS build cluster,
etc.) to community-owned infra, which is also a prerequisite for migrating the
Prow control plane in 2024. See the umbrella issue at
[k8s/test-infra#29722](https://github.com/kubernetes/test-infra/issues/29722).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- [prow](https://sigs.k8s.io/prow/)
- [testgrid](https://sigs.k8s.io/testgrid/)
- [boskos](https://sigs.k8s.io/boskos/)
- [kettle](https://git.k8s.io/test-infra/kettle/)
- [kubetest2](https://sigs.k8s.io/kubetest2/)

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

- KubeCon NA: [SIG Testing: Intro and Updates - Mahamed Ali, Cisco](https://youtu.be/aXftW1MRxJ0?si=uOv-jr8NKPYNiQQa)
- [Spotlight on SIG Testing](https://kubernetes.io/blog/2023/11/24/sig-testing-spotlight-2023/)

4. KEP work in 2023 (v1.27, v1.28, v1.29): None

## [Subprojects](https://git.k8s.io/community/sig-testing#subprojects)

**New in 2023:**
  - [Cloud Provider for KIND](https://sigs.k8s.io/cloud-provider-kind/)
  - [Hydrophone](https://sigs.k8s.io/hydrophone/)
**Continuing:**
  - boskos
  - e2e-framework
  - kind
  - kubetest2
  - prow
  - sig-testing
  - test-infra
  - testing-commons

## [Working groups](https://git.k8s.io/community/sig-testing#working-groups)

**New in 2023:**
 - LTS
**Retired in 2023:**
 - Reliability

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-testing/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-testing/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
