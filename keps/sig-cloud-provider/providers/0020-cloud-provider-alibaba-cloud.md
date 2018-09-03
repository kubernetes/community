---
kep-number: 20
title: Cloud Provider for Alibaba Cloud
authors:
  - "@aoxn"
owning-sig: sig-cloud-provider
reviewers:
  - "@andrewsykim"
approvers:
  - "@andrewsykim"
  - "@hogepodge"
  - "@jagosan"
editor: TBD
creation-date: 2018-06-20
last-updated: 2018-06-20
status: provisional

---

# Cloud Provider for Alibaba Cloud

This is a KEP for adding ```Cloud Provider for Alibaba Cloud``` into the Kubernetes ecosystem.

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Requirements](#requirements)
* [Proposal](#proposal)

## Summary

Alibaba Cloud provides the Cloud Provider interface implementation as an out-of-tree cloud-controller-manager. It allows Kubernetes clusters to leverage the infrastructure services of Alibaba Cloud .
It is original open sourced project is [https://github.com/AliyunContainerService/alicloud-controller-manager](https://github.com/AliyunContainerService/alicloud-controller-manager)

## Motivation

### Goals

Cloud Provider of Alibaba Cloud  implements interoperability between Kubernetes cluster and Alibaba Cloud. In this project, we will dedicated in:
- Provide reliable, secure and optimized integration with Alibaba Cloud for Kubernetes

- Help on the improvement for decoupling cloud provider specifics from Kubernetes implementation.



### Non-Goals

The networking and storage support of Alibaba Cloud for Kubernetes will be provided by other projects.

E.g.

* [Flannel network for Alibaba Cloud VPC](https://github.com/coreos/flannel)
* [FlexVolume for Alibaba Cloud](https://github.com/AliyunContainerService/flexvolume)


## Prerequisites

1. The VPC network is supported in this project. The support for classic network or none ECS environment will be out-of-scope.
2. When using the instance profile for authentication, an instance role is required to attach to the ECS instance firstly.
3. Kubernetes version v1.7 or higher

### Repository Requirements

[Alibaba Cloud Controller Manager](https://github.com/AliyunContainerService/alicloud-controller-manager) is a working implementation of the [Kubernetes Cloud Controller Manager](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/).

The repo requirements is mainly a copy from [cloudprovider KEP](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0002-cloud-controller-manager.md#repository-requirements). Open the link for more detail.

### User Experience Reports
As a CNCF Platinum member, Alibaba Cloud is dedicated in providing users with highly secure , stable and efficient cloud service.
Usage of aliyun container services can be seen from github issues in the existing alicloud controller manager repo: https://github.com/AliyunContainerService/alicloud-controller-manager/issues

### Testgrid Integration
 Alibaba cloud provider is reporting conformance test results to TestGrid as per the [Reporting Conformance Test Results to Testgrid KEP](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0003-testgrid-conformance-e2e.md).
 See [report](https://k8s-testgrid.appspot.com/conformance-alibaba-cloud-provider#Alibaba%20Cloud%20Provider,%20v1.10) for more details.

### CNCF Certified Kubernetes
 Alibaba cloud provider is accepted as part of the [Certified Kubernetes Conformance Program](https://github.com/cncf/k8s-conformance).
 For v1.11 See [https://github.com/cncf/k8s-conformance/tree/master/v1.11/alicloud](https://github.com/cncf/k8s-conformance/tree/master/v1.11/alicloud)
 For v1.10 See [https://github.com/cncf/k8s-conformance/tree/master/v1.10/alicloud](https://github.com/cncf/k8s-conformance/tree/master/v1.10/alicloud)
 For v1.9 See [https://github.com/cncf/k8s-conformance/tree/master/v1.9/alicloud](https://github.com/cncf/k8s-conformance/tree/master/v1.9/alicloud)
 For v1.8 See [https://github.com/cncf/k8s-conformance/tree/master/v1.8/alicloud](https://github.com/cncf/k8s-conformance/tree/master/v1.8/alicloud)

### Documentation
 
 Alibaba CloudProvider provide users with multiple documentation on build & deploy & utilize CCM. Please refer to [https://github.com/AliyunContainerService/alicloud-controller-manager/tree/master/docs](https://github.com/AliyunContainerService/alicloud-controller-manager/tree/master/docs) for more details.
 
### Technical Leads are members of the Kubernetes Organization

The Leads run operations and processes governing this subproject.

-  @cheyang Special Tech Leader, Alibaba Cloud. Kubernetes Member

## Proposal

Here we propose a repository from Kubernetes organization to host our cloud provider implementation.  Cloud Provider of Alibaba Cloud would be a subproject under Kubernetes community.

### Repositories

Cloud Provider of Alibaba Cloud will need a repository under Kubernetes org named ```kubernetes/cloud-provider-alibaba-cloud``` to host any cloud specific code.
The initial owners will be indicated in the initial OWNER files.

Additionally, SIG-cloud-provider take the ownership of the repo but Alibaba Cloud should have the fully autonomy permission to operator on this subproject.

### Meetings

Cloud Provider meetings is expected to have biweekly. SIG Cloud Provider will provide zoom/youtube channels as required. We will have our first meeting after repo has been settled.

Recommended Meeting Time: Wednesdays at 20:00 PT (Pacific Time) (biweekly). [Convert to your timezone](http://www.thetimezoneconverter.com/?t=20:00&tz=PT%20%28Pacific%20Time%29).
- Meeting notes and Agenda.
- Meeting recordings.


### Others
