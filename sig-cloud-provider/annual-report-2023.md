# 2023 Annual Report: SIG Cloud Provider

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

[**Bridget Kromhout and Michael McCune elected as new SIG chairs**](https://groups.google.com/g/kubernetes-sig-cloud-provider/c/BAh7gGZhfvo/m/5eJtVdoIDwAJ)

Previous chairs Andrew Sy Kim and Nick Turner stepped down during 2023, with Andrew staying on as a technical lead and Nick becoming an emeritus lead. Many thanks for their leadership and guidance over the years.

[**[KEP-2395] Update 2395-removing-in-tree-cloud-providers to beta status**](https://github.com/kubernetes/enhancements/pull/4190)

Updating KEP-2395 was the largest task that the SIG accomplished during 2023. During the update
of this KEP we changed the default behavior for Kubernetes to use external cloud controllers.
This update could not have been achieved without the support of the wider Kubernetes community as
there were several crucial changes to the CI and testing infrastructure which needed to be
coordinated. Thank you to Antonio Ojea, Andrew Sy Kim, Davanum Srinivas, Joel Speed, and Michael McCune
to name a few.

[**KEP-2699: Webhook KEP alpha in 1.27**](https://github.com/kubernetes/enhancements/pull/3813)

Webhooks for cloud controllers is a feature that will enable an alternate path for automation.
The initial example use case for the webhooks is to allow providers to stop using the persistent volume
label admission controller in favor of labelling in the webhook. This KEP is now properly marked in
alpha status and implementation work continues.

**Ending separate meeting for extraction/migration subproject**

With the advancement of KEP 2395 to beta status, the SIG indefinitely postponed the
separate subproject meetings in favor of handling this business during the main
SIG office hours. This decision was based on the notion that the work to finalize
the KEP into stable status will be much less than the previous work and the extra
time slot is no longer needed.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

[**Cloud Provider Baidu Cloud**](https://github.com/kubernetes-sigs/cloud-provider-baiducloud)

This repository has not had an update since April 2021. The SIG will review its status
and ownership for removal in the future.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

**KubeCon EU 2023 - The Ins and Outs of the Cloud Provider in Kubernetes**
[YouTube](https://www.youtube.com/watch?v=i2dNGzB0PMI) - [Sched](https://kccnceu2023.sched.com/event/1Hzci/the-ins-and-outs-of-the-cloud-provider-in-kubernetes-michael-mccune-joel-speed-red-hat-bridget-kromhout-microsoft)

Speakers: Bridget Kromhout (Microsoft), Michael McCune & Joel Speed (Red Hat)

> How do Kubernetes clusters interact with cloud services? In this session, the maintainers of SIG Cloud Provider will take a deep dive into the cloud provider framework, including how to implement an external cloud provider using the cloud provider interface, the cloud controller manager responsibilities, and an overview of the Kubelet image credential provider. We will also discuss the migration to external cloud providers in an HA configuration. We will identify trouble spots and processes that you should be aware of as you plan your migrations, and we will walk through the steps you can take to ensure zero downtime Kubernetes clusters as you perform this migration. Expect to walk away from this session with newfound knowledge about how Kubernetes interacts with cloud providers, an understanding of how to build an external cloud controller manager, and a solid plan of action for how you can migrate to external cloud controller managers without downtime.

**KubeCon NA 2023 - Improving Kubernetes Security with the Konnectivity Proxy**
[YouTube](https://www.youtube.com/watch?v=wTRezbXnlj8) - [Sched](https://kccncna2023.sched.com/event/1R2oN/improving-kubernetes-security-with-the-konnectivity-proxy-michael-mccune-red-hat-joseph-anttila-hall-google)

Speakers: Joseph Anttila Hall (Google)

> When architecting secure Kubernetes deployments it is often desirable to isolate the various streams of network traffic that exist within a cluster. For user-initiated traffic this process involves proxies and well-crafted firewall rules, but how do you properly separate API server-initiated traffic that flows to pods, nodes, and service networks?

> Konnectivity Proxy simplifies this question by providing a common methodology for shaping the network egress traffic from Kubernetes API servers. Securing network traffic within Kubernetes clusters is a vital step to ensure that user data is protected and that cloud resources are not exploited. Project maintainers will cover an overview of the Konnectivity proxy, the goals of this project as a collaboration between SIG Cloud Provider and SIG API Machinery, its current status, and will share experiences running Konnectivity at GKE. Attendees will leave with new knowledge and tools to secure their Kubernetes clusters.

**KubeCon CN 2023 - Navigating Kubernetes Cloud Provider**
[YouTube](https://www.youtube.com/watch?v=kmvQQW56nmQ) - [Sched](https://kccncosschn2023.sched.com/event/1PTKN/zhen-kubernetesss-navigating-kubernetes-cloud-provider-pengfei-ni-microsoft)

Speakers: Pengfei Ni (Microsoft)

_Please note, this talk is presented in Mandarin Chinese with slides in English._

> 让我们讨论一下Kubernetes云提供商是如何随着时间的推移发展的，以及您现在为了实现目标需要了解的内容。对云提供商内部的详细了解始于对四个控制器（节点、节点生命周期、路由和服务）的深入探索。 从内部云提供商迁移的关键概念将涵盖何时以及如何运行云控制器管理器，包括Kubelet镜像凭据提供程序和CSI驱动程序的外部迁移路径。了解如何构建自己的CCM将为您实现最佳结果提供一个框架。 外部迁移的最佳实践将包括最新更新的所有建议。您将了解Kubernetes如何与云提供商进行交互，以及如何在没有停机时间的情况下迁移到外部云控制器管理器。最后，我们将介绍如何与Kubernetes云提供商社区进行连接，您将获得满足云提供商需求的见解。

> Let's discuss how Kubernetes Cloud Provider has evolved over time and what you need to know right now for your goals. A detailed look at cloud provider internals starts with an in-depth exploration of the four controllers (Node, Node Lifecycle, Route, and Service). Key concepts for migration from the in-tree cloud provider will cover when and how to run the cloud controller manager, including out-of-tree migration paths for the Kubelet image credential provider and the CSI drivers. Looking at how to build your own CCM will give you a framework for achieving optimal results. Best practices for out-of-tree migration will include all the latest recommendations in light of recent updates. You'll see how Kubernetes interacts with cloud providers and how to migrate to an external cloud controller manager without downtime. Wrapping up with a look at how you can connect with the Kubernetes Cloud Provider community, you will leave with insights into how to meet your cloud provider needs. 

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [2699 - Add webhook hosting to CCM.](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2699-add-webhook-hosting-to-ccm) - v1.27

  - Beta
    - [2395 - Removing In-Tree Cloud Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2395-removing-in-tree-cloud-providers) - v1.29


## [Subprojects](https://git.k8s.io/community/sig-cloud-provider#subprojects)


**New in 2023:**

  - [provider-equinix-metal](https://git.k8s.io/community/<no value>#provider-equinix-metal)

**Continuing:**

  - cloud-provider-extraction-migration
  - kubernetes-cloud-provider
  - provider-alibaba-cloud
  - provider-aws
  - provider-azure
  - provider-baiducloud
  - provider-gcp
  - provider-huaweicloud
  - provider-ibmcloud
  - provider-oci
  - provider-openstack
  - provider-vsphere

## [Working groups](https://git.k8s.io/community/sig-cloud-provider#working-groups)

**Continuing:**
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cloud-provider/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cloud-provider/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
