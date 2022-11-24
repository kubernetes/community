# 2021 Annual Report: SIG Windows

## Current initiatives

1. What work did the SIG do this year that should be highlighted?
   - Implemented hostProcess container support in Kubernetes (now in beta) and pomoted adoption in multiple open source communities
     - https://github.com/kubernetes-sigs/sig-windows-tools/tree/master/hostprocess - for examples of running flannel, calico, csi-proxy, kube-proxy as hostProcess containers
     - https://github.com/weaveworks/kured/pull/460  - KuReD Windows support
     - https://github.com/prometheus-community/windows_exporter/pull/864 - node exporter support
   - Defined the `kubectl node logs` command interface.
   - Made the developer UX for windows transparent with sig-windows-dev-tools.
   - Defined windows operational readiness standards.
   - Defined the pod OS field.

2. What initiatives are you working on that aren't being tracked in KEPs?
   - Migration of the windows kube-proxy to KPNG.
   - Migration of testgrid reporting jobs from aks-engine to cluster-api/cluster-api-provder-azure.
   - Dockershim removal / validation for Windows nodes.

3. KEP work in 2021 (1.x, 1.y, 1.z):
   - Stable
     - (1.22) [1122 - windows-csi-support](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1122-windows-csi-support/README.md)
   - Beta
     - (1.23) [1981 - Windows Privileged Container Support](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1981-windows-privileged-container-support/README.md)
     - (1.23) [2802 -Identify Windows pods at API admission level authoritatively](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/2802-identify-windows-pods-apiserver-admission/kep.yaml)
   - Alpha
     - (1.22) [1981 - Windows Privileged Container Support](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1981-windows-privileged-container-support/README.md)
     - (1.23) [2802 -Identify Windows pods at API admission level authoritatively](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2802-identify-windows-pods-apiserver-admission/README.md)
   - Pre-alpha (Targeting 1.24)
     - [2578 - Windows Operational Readiness](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/2578-windows-conformance/kep.yaml)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)
   - csi-proxy and storage: this seems like an underserved area for windows https://github.com/kubernetes-csi/csi-proxy (meeting info is there).
2. What metrics/community health stats does your group care about and/or measure?
   - stars for ksigs/windows specific repos
     - sig-windows-dev-tools
       - https://github.com/kubernetes-sigs/sig-windows-dev-tools -> up to 46, represents interest
     - sig-windows-tools
       - represents people trying to install windows on k8s nodes
     - windows-gmsa
       - represents enterprises integrating windows pods into GMSA
3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?
   - yes
4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?
   - yes

5. Does the group have contributors from multiple companies/affiliations?
   - yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?
   - testing hostProcess implementations on several windows apps
   - improving our dev tools environment to grow the community
   - hardening the CSI proxy and CSI support ecosystem
   - performance testing Kubernetes on Windows extensively and publishing results in cncf blog posts

## Membership

- Primary slack channel member count: 1507
- Primary mailing list member count: 188
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 10
- Unique reviewers for SIG-owned packages: 6
- Unique approvers for SIG-owned packages: 4

Include any other ways you measure group membership

## Subprojects

- windows csi-proxy subproject is active and healthy https://github.com/kubernetes-csi/csi-proxy
  - meetings going well
  - new optimization issue came up recently community engaged on it across companies vmware,rancher

## Working groups

n/a

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
  - [KubeCon EU 2021 virtual talk](https://www.youtube.com/watch?v=zJw4lrB7kKs)
  - [KubeCon NA 2021 virtual talk](https://www.youtube.com/watch?v=fSmDmwKwFfQ)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
