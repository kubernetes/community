# 2022 Annual Report: SIG Windows

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - [HostProcess containers](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1981-windows-privileged-container-support/README.md) graduated to `stable` in v1.26
   - Windows performance test results are now reported on the [k8s perf dashboard](https://perf-dash.k8s.io/#/?jobname=soak-tests-capz-windows-2019&metriccategoryname=E2E&metricname=CPUUsage)
   - [Guide for adding Windows nodes using kubeadm](https://github.com/kubernetes-sigs/sig-windows-tools/blob/master/guides/guide-for-adding-windows-node.md)
   - [Windows-node-exporter](https://github.com/prometheus-community/windows_exporter) can now be run as a daemonSet with HostProcess containers!

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Out-of-tree kube-proxy for Windows based on KNPG [windows-service-proxy](https://github.com/kubernetes-sigs/windows-service-proxy)
   - Enabling Hyper-V isolated containers with containerd
     - https://github.com/containerd/containerd/issues/6862
     - https://github.com/kubernetes-sigs/windows-testing/issues/364
   - Unit tests for Windows - https://testgrid.k8s.io/sig-windows-signal#windows-unit-master
   - ContainerD work and KEP updates to support [KEP 2371 - cAdvisor-less, CRI-full Container and Pod Stats](https://github.com/kubernetes/enhancements/issues/2371) for Windows
   - [sig-windows-dev-tools](https://github.com/kubernetes-sigs/sig-windows-dev-tools) continual improvements
     - Add support for M1/M2 macbooks

3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2578 - Windows Conformance](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2578-windows-conformance) - v1.24
    - [3503 - Host Network Support for Windows Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/3503-host-network-support-for-windows-pods) - v1.26 
  - stable:
    - [1981 - Windows Privileged Container Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/1981-windows-privileged-container-support) - v1.26
    - [2802 - Identify Pod's OS during API Server admission](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2802-identify-windows-pods-apiserver-admission) - v1.25

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Maintaining SIG-Windows E2E test passes.
   - Attracting long-term contributors for all SIG-windows sub-projects.

2. What metrics/community health stats does your group care about and/or measure?

   - open / stale issues in Windows specific k-sigs/ repos
   - stars for k-sigs/windows specific repos

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - yes however setting up K8s clusters with Windows nodes for local development scenarios continues to be difficult and time consuming for new contributors. We are continually trying to improve this experience with the sig-windows-dev-tool project.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - yes

5. Does the group have contributors from multiple companies/affiliations?

   - yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Help maintaining [sig-windows-dev-tools](https://github.com/kubernetes-sigs/sig-windows-dev-tools).
     This project is aimed to make testing Windows-related changes in core Kubernetes components easier which is consistently idefintied as one of the hurdles to contrubting in the SIG. 
     From the project's [readme](https://github.com/kubernetes-sigs/sig-windows-dev-tools#goal):
     > Our goal is to make Windows ridiculously easy to contribute to, play with, and learn about for anyone interested in using or contributing to the ongoing Kubernetes-on-Windows story. Windows is rapidly becoming an increasingly viable alternative to Linux thanks to the recent introduction of Windows HostProcess containers and Windows support for NetworkPolicies + Containerd integration.

## Membership

- Primary slack channel member count: 1687
- Primary mailing list member count: 212
- Primary meeting attendee count (estimated, if needed): 12
- Primary meeting participant count (estimated, if needed): 6-9
- Unique reviewers for SIG-owned packages: 8
- Unique approvers for SIG-owned packages: 8

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-windows#subprojects)

**New in 2022:**

- windows-operational-readiness
- windows-service-proxy

**Continuing:**

- windows-gmsa
- windows-samples
- windows-testing
- windows-tools

## [Working groups](https://git.k8s.io/community/sig-windows#working-groups)

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon EU 2022 maintainer talk](https://www.youtube.com/watch?v=THaDy6u-Cgk)
      - [KubeCon NA 2022 maintainer talk](https://www.youtube.com/watch?v=rELpBRmXaTw)
      - [SIG leadership updates](https://groups.google.com/g/kubernetes-sig-windows/c/jFLFUAQpM2c)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
