# 2025 Annual Report: SIG Cluster Lifecycle

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- Cluster API added
[v1beta2](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.11.0) of its API.
- Cluster API added support for in
[place updates](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/proposals/20240807-in-place-updates.md).
- Cluster API added support for
[chained upgrades](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/proposals/20250513-chained-and-efficient-upgrades-for-clusters-with-managed-topologies.md).
- Image Builder made
[various improvements](https://github.com/kubernetes-sigs/image-builder/releases)
for Windows support.
- kOps had their [first contribution](https://github.com/kubernetes/kops/pull/17767)
from GitHub's Copilot.
- kubeadm graduated the feature gates
[NodeLocalCRISocket](https://github.com/kubernetes/kubernetes/pull/135742)
and [ControlPlaneKubeletLocalMode](https://github.com/kubernetes/kubernetes/pull/134106) to GA.
- Kubespray added support for control plane
[reconfiguration on upgrades](https://github.com/kubernetes-sigs/kubespray/releases/tag/v2.27.1).
- Minikube made it possible to run
[GPU workloads on MacOS](https://github.com/kubernetes/minikube/releases/tag/v1.37.0).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- cluster-api-provider-aws
- cluster-api-provider-digitalocean
- cluster-api-provider-cloudstack

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

- [Kubespray: Driving Cost-Efficiency for AI on Kubernetes](https://kccnceu2025.sched.com/event/1td0D/kubespray-driving-cost-efficiency-for-ai-on-kubernetes-antoine-legrand-conny-gmbh-mohamed-zaian-new-work-se)
- [Resilient Multi-Cloud Strategies: Harnessing Kubernetes, Cluster API, and Cell-Based Architecture](https://kccnceu2025.sched.com/event/1txDE/resilient-multi-cloud-strategies-harnessing-kubernetes-cluster-api-and-cell-based-architecture-tasdik-rahman-javi-mosquera-new-relic)
- [Kubernetes CRD Design for the Long Haul: Tips, Tricks, and Lessons Learned](https://kccnceu2025.sched.com/event/1tx6k/kubernetes-crd-design-for-the-long-haul-tips-tricks-and-lessons-learned-christian-schlotter-broadcom-fabrizio-pandini-vmware-by-broadcom)
- [Cluster API status update and discussion](https://maintainersummiteu2025.sched.com/event/1uSNe/project-meeting-cluster-api-status-update-open-discussion)

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Alpha
    - [5607 - Allow HostNetwork Pods to User Namespaces](http://git.k8s.io/enhancements/keps/sig-node/5607-hostnetwork-userns) - v1.35

  - Stable
    - [4656 - Add kubelet instance configuration to configure CRI socket for each node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/4656-add-kubelet-instance-configuration) - v1.35

## [Subprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects)


**Retired in 2025:**
  - cluster-api-provider-packet
**Continuing:**
  - cluster-addons
  - cluster-api
  - cluster-api-addon-provider-helm
  - cluster-api-ipam-provider-in-cluster
  - cluster-api-operator
  - cluster-api-provider-aws
  - cluster-api-provider-azure
  - cluster-api-provider-cloudstack
  - cluster-api-provider-digitalocean
  - cluster-api-provider-gcp
  - cluster-api-provider-ibmcloud
  - cluster-api-provider-kubemark
  - cluster-api-provider-kubevirt
  - cluster-api-provider-openstack
  - cluster-api-provider-vsphere
  - image-builder
  - kOps
  - karpenter-provider-cluster-api
  - kubeadm
  - kubespray
  - minikube

## [Working groups](https://git.k8s.io/community/sig-cluster-lifecycle#working-groups)

**New in 2025:**
 - Node Lifecycle
**Continuing:**
 - LTS
 - etcd Operator

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cluster-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
