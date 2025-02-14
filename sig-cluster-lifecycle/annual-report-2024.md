# 2024 Annual Report: SIG Cluster Lifecycle

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- No major SIG level KEPs.
- No governance or leadership changes.
- The SIG started participating in the cross-SIG WG etcd operator with SIG etcd.
- We retired the inactive etcdadm and cluster-api-provider-nested subprojects.
- kubeadm released v1beta4 and added phase support for `upgrade apply` phases.
- minikube added support for multi-control plane clusters (HA) and improved the GPU support.
- CAPI started work on v1beta2 and a new process for cluster Conditions.
- kOps improved its Azure support and added the `reconcile` command to handle k8s skew issues.
- kubespray added support for more Linux distributions and exposed more configuration options.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- In 2023 we collected a
[subproject self-assessment](https://docs.google.com/forms/d/e/1FAIpQLSc0CqmfaOIK4bCbEDhh0qiF5wCHi6Uvy0uR_k8egOtafalpow/viewanalytics),
which still has a valid list of areas that need help.

The most common areas where we need help are:
- Contributing to code
- Reviewing code

Subprojects that might need help with additional OWNERS, in no particular order,
and maybe an incomplete list:
- cluster-api
- cluster-api-addon-provider-helm
- cluster-api-ipam-provider-in-cluster
- cluster-api-provider-aws
- cluster-api-provider-gcp
- cluster-api-provider-kubemark
- cluster-api-provider-openstack
- cluster-api-provider-packet
- kubeadm

We encourage contributors who are interested in our subprojects to reach out to the
SIG mailing list or to individual subproject Slack channels.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

KubeCon EU 2024:
- [how-to-deploy-an-ai-optimized-k8s-cluster-with-kubespray](https://kccnceu2024.sched.com/event/1YhhP/how-to-deploy-an-ai-optimized-k8s-cluster-with-kubespray-kay-yan-daocloud-mohamed-zaian-new-work-se)

KubeCon NA 2024:
- [luster-api-deep-dive-roadmap-to-api-graduation](https://kccncna2024.sched.com/event/1howc/cluster-api-deep-dive-roadmap-to-api-graduation-christian-schlotter-broadcom-vince-prignano-apple)
- [bare-metal-kubernetes-with-kops-gathering-community-wisdom](https://kccncna2024.sched.com/event/1howo/bare-metal-kubernetes-with-kops-gathering-community-wisdom-justin-santa-barbara-google-ciprian-hacman-microsoft)

KubeCon China 2024:
- [kubespray-unleashed-navigating-bare-metal-services](https://kccncossaidevchn2024.sched.com/event/1eYXD/kubespray-unleashed-navigating-bare-metal-services-in-kubernetes-for-llm-and-rag-kubespraydaepnanokubernetesllmreragzhu-ya-lie-shu-kay-yan-daocloud-alan-leung-equinix)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

<!--
  - Alpha
    - [4656 - Add kubelet instance configuration to configure CRI socket for each node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/4656-add-kubelet-instance-configuration) - v1.32

 -->

While we have a few SIG level KEPs in `kubernetes/enhancements`, kubeadm is the only
subproject that we have that is part of the Kubernetes release and more actively receives KEPs.

In 2024 there were 2 new kubeadm KEPs:
- [4471-cp-join-kubelet-local-apiserver](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/4471-cp-join-kubelet-local-apiserver)
- [4656-add-kubelet-instance-configuration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/4656-add-kubelet-instance-configuration)

## [Subprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects)


**New in 2024:**
  - karpenter-provider-cluster-api
**Retired in 2024:**
  - etcdadm
  - cluster-api-provider-nested
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
  - cluster-api-provider-packet
  - cluster-api-provider-vsphere
  - image-builder
  - kOps
  - kubeadm
  - kubespray
  - minikube

## [Working groups](https://git.k8s.io/community/sig-cluster-lifecycle#working-groups)

**New in 2024:**
 - etcd Operator
**Continuing:**
 - LTS

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cluster-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
