---
kep-number: 4
title: Cloud Provider Template
authors:
  - "@janedoe"
owning-sig: sig-cloud-provider
participating-sigs:
  - sig-aaa
  - sig-bbb
reviewers:
  - TBD
  - "@alicedoe"
approvers:
  - "@andrewsykim"
  - "@hogepodge"
  - "@jagosan"
editor: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: provisional
see-also:
  - KEP-1
  - KEP-2
replaces:
  - KEP-3
superseded-by:
  - KEP-100
---

# Cloud Provider FooBar

This is a KEP template, outlining how to propose a new cloud provider into the Kubernetes ecosystem.

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Requirements](#requirements)
* [Proposal](#proposal)

## Summary

This is where you add a summary of your cloud provider and other additional information about your cloud provider that others may find useful.

## Motivation

### Goals

This is where you can specify any goals you may have for your cloud provider.

### Non-Goals

This is where you can specify any work that you think is outside the scope of your cloud provider.

## Prerequisites

This is where you outline all the prerequisites for new providers that have been met.

### Repository Requirements

For [repository requirements](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0002-cloud-controller-manager.md#repository-requirements) you are expected to have a repo (belonging to any organization, ideally owned by your cloud provider) that has a working implementation of the [Kubernetes Cloud Controller Manager](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/). Note that the list of requirements are subject to change.

### User Experience Reports

There must be a reasonable amount of user feedback about running Kubernetes for this cloud provider. You may want to link to sources that indicate this such as github issues, product data, customer tesitimonials, etc.

### Testgrid Integration

Your cloud provider is reporting conformance test results to TestGrid as per the [Reporting Conformance Test Results to Testgrid KEP](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0003-testgrid-conformance-e2e.md).

### CNCF Certified Kubernetes

Your cloud provider is accepted as part of the [Certified Kubernetes Conformance Program](https://github.com/cncf/k8s-conformance).

### Documentation

There is documentation on running Kubernetes on your cloud provider as per the [cloud provider documentation KEP](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0004-cloud-provider-documentation.md).

### Technical Leads are members of the Kubernetes Organization

All proposed technical leads for this provider must be members of the Kubernetes organization. Membership is used as a signal for technical ability, commitment to the project, and compliance to the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) which we believe are important traits for subproject technical leads. Learn more about Kubernetes community membership [here](https://github.com/kubernetes/community/blob/master/community-membership.md).

## Proposal

This is where you can talk about what resources from the Kubernetes community you would like such as a repository in the Kubernetes organization to host your provider code.

### Subproject Leads

This is where you indicate the leads for the subproject. Make sure you include their github handles. See the [SIG Charter](https://github.com/kubernetes/community/blob/master/sig-cloud-provider/CHARTER.md#subprojectprovider-owners) for more details on expectations from subproject leads.

### Repositories

This is where you propose a repository within the Kubernetes org, it's important you specify the name of the repository you would like. Cloud providers typically have at least 1 repository named `kubernetes/cloud-provider-foobar`. It's also important to indiciate who the initial owners of the repositories will be. These owners will be added to the initial OWNERS file. The owners of the subproject must be owners of the repositories but you can add more owners in the repo if you'd like. If you are requesting any repositories, be sure to add them to the SIG Cloud Provider [subproject list](https://github.com/kubernetes/community/tree/master/sig-cloud-provider#subprojects).

### Meetings

This where you specify when you will have meetings to discuss development of your cloud provider. SIG Cloud Provider will provide zoom/youtube channels as required. Note that these meetings are in addition to the biweekly SIG Cloud Provider meetings that subproject leads are strongly encouraged to attend.


### Others

Feel free to add anything else you may need.
