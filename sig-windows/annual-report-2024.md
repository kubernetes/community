# 2024 Annual Report: SIG Windows

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  - Continued working on [Image pull per runtime class](https://github.com/kubernetes/enhancements/issues/4216)
  - [Windows support for memory pressure eviction](https://github.com/kubernetes/kubernetes/pull/122922)
  - Migrated CI jobs to community infra
  - Unit and e2e testing improvements
  - [Documentation](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/adding-windows-nodes/) improvements for adding Windows nodes

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   No

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

  - [SIG Windows Retrospective and Windows Image Building Deep Dive, Kubecon EU](https://sched.co/1YhiZ)
  - [What's new with SIG Windows, Kubecon NA](https://sched.co/1hoxQ)

4. KEP work in 2024 (v1.30, v1.31, v1.32):

  - Alpha
    - [4802 - Windows Graceful Node Shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/4802-windows-node-shutdown) - v1.32
    - [4885 - Windows CPU and Memory Affinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/4885-windows-cpu-and-memory-affinity) - v1.32

  - Beta
    - [2258 - Node log query](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2258-node-log-query) - v1.30

## [Subprojects](https://git.k8s.io/community/sig-windows#subprojects)


**Continuing:**
  - windows-gmsa
  - windows-operational-readiness
  - windows-samples
  - windows-service-proxy
  - windows-testing
  - windows-tools

## [Working groups](https://git.k8s.io/community/sig-windows#working-groups)


## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
