---
kep-number: 4
title: Cloud Provider Documentation
authors:
  - "@d-nishi"
  - "@hogepodge"
owning-sig: sig-cloud-provider
participating-sigs:
  - sig-docs
  - sig-cluster-lifecycle
  - sig-aws
  - sig-azure
  - sig-gcp
  - sig-openstack
  - sig-vmware
reviewers:
  - "@andrewsykim"
  - "@calebamiles"
  - "@hogepodge"
  - "@jagosan"
approvers:
  - "@thockin"
editor: TBD
status: provisional
---
## Transfer the responsibility of maintaining valid documentation for Cloud Provider Code to the Cloud Provider

### Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
* [User Stories [optional]](#user-stories)
    * [Story 1](#story-1)
    * [Story 2](#story-2)
* [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints)
* [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Alternatives [optional]](#alternatives)

### Summary
This KEP describes the documentation requirements for both in-tree and out-of-tree cloud controller managers. 
These requirements are meant to capture critical usage documentation that is common between providers, set requirements for individual documentation, and create consistent standards across provider documentation. The scope of this document is limited to in-tree code that interfaces with kube-controller-manager, and out-of-tree code that interfaces with cloud-controller-manager

### Motivation
Currently documentation for cloud providers for both in-tree and out-of-tree managers is limited in both scope, consistency, and quality. This KEP describes requirements, to be reached in the 1.12 release cycle, to create and maintain consistent documentation across all cloud provider manager code. By establishing these standards, SIG-Cloud-Provider will benefit the user-community by offering a single discoverable source of reliable documentation while relieving the SIG-Docs team from the burden of maintaining out-dated duplicated documentation.

#### Goals
* Produce a common document that describes how to configure any in-tree cloud provider that can be reused by tools such as kubeadm, to create minimum viable Kubernetes clusters.
  * Create documentation requirements on how to configure in-tree cloud providers.
  * Produce documentation for every in-tree cloud provider.
* Provide a common document that describes how to configure any out-of-tree cloud-controller-manager by provider.
  * Create documentation requirements on how to configure out-of-tree cloud providers.
  * Produce documentation for every out-of-tree cloud provider.
* Maintain developer documentation for anyone wanting to build a new cloud-controller-manager.
* Generate confidence in SIG-docs to confidently link to SIG-Cloud-Provider documentation for all future releases.

#### Non-Goals
This KEP is limited to documenting requirements for control plane components for in-tree implementation and cloud-controller-manager for out-of-tree implementation. It is not currently meant to document provider-specific drivers or code (example: Identity & access management: Keystone for Openstack, IAM for AWS etc).
SIG-Docs is not expected to produce or maintain any of this documentation.

### Proposal

#### Goal 1: 
Produce a common document that describes how to configure any in-tree cloud provider that can be reused by tools such as kubeadm, to create minimum viable Kubernetes clusters.

Kubernetes documentation lists details of current cloud-provider [here](https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/). Additional documentation [(1),](https://kubernetes.io/docs/concepts/services-networking/service/) [(2)](https://kubernetes.io/docs/tasks/administer-cluster/developing-cloud-controller-manager/) that link to cloud-provider code currently remains detached and poorly maintained. 

#### Requirement 1: 
Provide validated manifests for kube-controller-manager, kubelet and kube-apiserver by cloud-provider to enable a Kubernetes administrator to run cloud-provider=<providername> in-tree with kube-controller-manager as is feasible today. This is only relevant to environments where a Kubernetes administrator/user has (or) wants access to the control plane. Environments such as Amazon EKS, Google’s GKE, Azure’s AKS and other such platforms are out of context here.

* Add --cloud-provider=<providername> to the kube-apiserver, kube-controller-manager and every kubelet. These manifests should be regularly updated and listed at this location: https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/aws/docs/
  * Example of an [apiserver manifest](https://gist.github.com/d-nishi/1109fec153930e8de04a1bf160cacffb)
  * Example of [kube-controller-manager](https://gist.github.com/d-nishi/a41691cdf50239986d1e725af4d20033)
  * Example of a [systemd service for kubelet](https://gist.github.com/d-nishi/289cb82367580eb0cb129c9f967d903d) and [kubelet config](https://gist.github.com/d-nishi/d7f9a1b59c0441d476646dc7cce7e811)
* Run in-tree cloud-controller-manager as a [daemon-set](https://gist.github.com/d-nishi/38e3b7051029b5d1a1772f3862f62ce9)/deployment/replicaset/static pod on the cluster.

#### Requirement 2: 
Provide validated/tested descriptions with examples of controller features (annotations or tags) that are cloud-provider dependent that can be reused by any Kubernetes administrator to run `cloud-provider-<providername>` in-tree with `kube-controller-manager` as is described in the code <cloudprovider.go> Example: aws.go
These manifests should be regularly tested and updated post testing in the relevant provider location E.g.: https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/aws/docs/
* Node Controller (or) NodeName
* Service Controller (or) LoadBalancer
* Volume Controller (or) persistent volume labels controller
* Other provider-specific-Controller e.g. Route controller for GCP.

#### Goal 2: 
Provide a common document that describes how to configure any out-of-tree cloud-controller-manager by provider.

#### Requirement 1:  
Provide validated manifests for kube-controller-manager, kubelet and kube-apiserver by cloud-provider to enable a Kubernetes administrator to run out-of-tree cloud-controller-manager.
* Remove --cloud-provider flag from kube-apiserver and kube-controller-manager. Remove this flag from the manifest when the flag is deprecated in a future release. Run kubelet with --cloud-provider=external.
* Run out-of-tree cloud-controller-manager as a daemon-set/deployment/replicaset/static pod on the cluster.

#### Requirement 2: 
List out the latest annotations or tags that are cloud-provider dependent and will be used by the Kubernetes administrator to run `cloud-provider-<providername>` in-tree with `kube-controller-manager` as is described in the code <cloudprovider.go> Eg. aws.go
These manifests should be regularly tested and updated in the relevant provider location E.g.: https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/aws/docs/
* Node Controller (or) NodeName
* Service Controller (or) LoadBalancer
* Volume Controller (or) persistent volume labels controller
* Other provider-specific-Controller e.g. Route controller for GCP

### User Stories [optional]

#### Story 1
Sally is a devops engineer wants to run Kubernetes clouds across her on-premise environment and public cloud sites. She wants to use ansible or terraform to bring up Kubernetes v1.11. She references the cloud-provider documentation to understand how to enable in-tree provider code, and has a consistent set of documentation to help her write automation to target each individual cloud.

#### Story 2
Sam wants to add advanced features to external cloud provider. By consulting the external cloud provider documents, they are able to set up a development and test environment. Where previously documentation was inconsistent and spread across multiple sources, there is a single document that allows them to immediately launch provider code within their target cloud.

### Implementation Details/Notes/Constraints [optional]
The requirements set forward need to accomplish several things:
* Identify and abstract common documentation across all providers.
* Create a consistent format that makes it easy to switch between providers.
* Allow for provider-specific documentation, quirks, and features.

### Risks and Mitigations
This proposal relies heavily on individual cloud-provider developers to provide expertise in document generation and maintenance. Documentation can easily drift from implementation, making for a negative user experience.
To mitigate this, SIG-Cloud-Provider membership will work with developers to keep their documentation up to date. This will include a review of documents along release-cycle boundaries, and adherence to release-cycle deadlines.
SIG-Cloud-Provider will work with SIG-Docs to establish quality standards and with SIG-Node and SIG Cluster Lifecycle to keep common technical documentation up-to-date.

### Graduation Criteria
This KEP represents an ongoing effort for the SIG-Cloud-Provider team.
* Immediate success is measured by the delivery of all goals outlined in the Goals (1) section.
* Long Term success is measured by the delivery of goals outlined in the Goals (2) section.
* Long Term success is also measured by the regular upkeep of all goals in Goals (1) and (2) sections.

### Implementation History
Major milestones in the life cycle of a KEP should be tracked in Implementation History. Major milestones might include:
* the Summary and Motivation sections being merged signaling SIG acceptance
* the Proposal section being merged signaling agreement on a proposed design
* the date implementation started - July 25 2018
* the first Kubernetes release where an initial version of the KEP was available - v1.12
* the version of Kubernetes where the KEP graduated to general availability - v1.14
* the date when the KEP was retired or superseded - NA

### Alternatives [optional]
The Alternatives section is used to highlight and record other possible approaches to delivering the value proposed by a KEP.
* SIG docs could tag cloudprovider documentation as a blocking item for Kubernetes releases
* SIG docs could also assign SIG-<provider> leads to unblock cloudprovider documentation in the planning phase for the release.

