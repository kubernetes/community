# Cloud Provider BaiduCloud

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Requirements](#requirements)
* [Proposal](#proposal)

## Summary

Baidu is a gold member of CNCF and we have a large team working on Kubernetes and related projects like complex scheduling, heterogeneous computing, auto-scaling etc. We build cloud platform to support Baidu emerging business including autonomous driving, deep learning, blockchain by leveraging Kubernetes. We also provide public container services named cloud container engine(CCE).

## Motivation

### Goals

- Building, deploying, maintaining, supporting, and using Kubernetes on Baidu Cloud Container Engine(CCE) and Baidu Private Cloud(BPC). Both of the project are built on Kubernetes and related CNCF project.

- Designing, discussing, and maintaining the cloud-provider-baidu repository under Github Kubernetes project. 

### Non-Goals

- Identify domain knowledge and work that can be contributed back to Kubernetes and related CNCF projects.

- Mentor CCE and BPC developers to contribute to CNCF projects.

- Focus on Kubernetes and CNCF related projects, the discussion of development issue for CCE and BCP will not be included in the SIG.

## Prerequisites

### Repository Requirements

The repository url which meets all the requirements is: https://github.com/baidu/cloud-provider-baiducloud

### User Experience Reports

![CCE-ticket-1](http://agroup-bos.su.bcebos.com/c34021571744895b5d9fffd8c22d8409469f47b3)
CCE-ticket-1: User want to get the Kubernetes cluster config file by using account's aksk.

![CCE-ticket-2](http://agroup-bos.su.bcebos.com/756c9463c8487dee9c26d7725e127c5b64975fc4)
CCE-ticket-2: User want to modify the image repository's username.

![CCE-ticket-3](http://agroup-bos.su.bcebos.com/7a4506fcb1fbeeb15c86060cfbb6e69d090c8984)
CCE-ticket-3: User want to have multi-tenant ability in a shared large CCE cluster.


## Proposal

### Subproject Leads

The subproject will have 3 leaders at any given time. I will be an initial point of contact as we work on creating the subporject. My github account is: tizhou86

I will be the subproject leader at this moment. My github account is: tizhou86.

### Repositories

The repository we propose at this moment is: kubernetes/cloud-provider-baiducloud, I'll be the initial point of contact.

### Meetings

We plan to have bi-week online meeting at https://zoom.us/j/5134183949 on every next Wednesday 6pm PST.


### Others

NA at this moment.

