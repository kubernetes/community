# 2025 Annual Report: SIG Windows

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

   - [Windows Graceful Node Shutdown (KEP-4802)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/4802-windows-node-shutdown) promoted to beta in v1.34, enabled by default - Windows nodes now have the same graceful shutdown handling that Linux nodes have had
   - [DSR and Overlay support in Windows kube-proxy (KEP-5100)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/5100-windows-dsr-and-overlay-support) promoted to stable in v1.34
   - Continued work on [Windows CPU and Memory Affinity (KEP-4885)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/4885-windows-cpu-and-memory-affinity) - enabling CPU, Memory and Topology Managers in kubelet for Windows
   - Continued collaboration on [Image pull per runtime class (KEP-4216)](https://github.com/kubernetes/enhancements/issues/4216) with sig-node - important for Windows HyperV isolation scenarios
   - Windows Server 2025 support added to Kubernetes testing infrastructure
   - Ongoing CI and e2e testing improvements in the community infrastructure

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   No

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

   - [SIG Windows Updates, KubeCon NA](https://kccncna2025.sched.com/) - Mark Rossetti (Microsoft) and Jose Valdes (Red Hat)

4. KEP work in 2025 (v1.33, v1.34, v1.35):

  - Beta
    - [4802 - Windows Graceful Node Shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/4802-windows-node-shutdown) - v1.34
  - Stable
    - [5100 - DSR and Overlay support in Windows kube-proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/5100-windows-dsr-and-overlay-support) - v1.34

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
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
