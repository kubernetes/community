---
kep-number: 8
title: VMware Cloud Provider transition
authors:
  - "@frapposelli"
owning-sig: sig-vmware
reviewers:
  - "@frapposelli"
  - "@dougm"
approvers:
  - "@frapposelli"
  - "@dougm"  
creation-date: 2018-04-16
last-updated: 2018-04-26
status: implementable
---

# VMware Cloud Provider transition

## Table of Contents

<!-- TOC -->

- [VMware Cloud Provider transition](#vmware-cloud-provider-transition)
  - [Table of Contents](#table-of-contents)
  - [Summary](#summary)
  - [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
  - [Proposal](#proposal)
    - [Graduation to Beta](#graduation-to-beta)

<!-- /TOC -->

## Summary

We want to align the development of the VMware vSphere cloud provider to [KEP0002](https://github.com/kubernetes/community/blob/master/keps/0002-controller-manager.md), this requires a refactor of the existing in-tree vSphere cloud provider and the transition to the Cloud Controller Manager model, this is the premise for the creation of a cloud-provider-vsphere subproject under the SIG-VMware ownership.

## Motivation

As Kubernetes core is trying to remove any dependencies tied to specific cloud providers, we want to facilitate this process by proactively working on the transition of the vSphere cloud provider.

This will also enable faster iteration on the cloud provider for further enhancements and bug fixes.

### Goals

- Refactor in-tree cloud provider code to support migration to Cloud Controller Manager
- Reach functional parity between in-tree provider and CCM
- Have E2E test framework in place for CCM
- Move code from in-tree provider to kubernetes/cloud-provider-vsphere

### Non-Goals

- Expand the cloud provider outside the vSphere platform

## Proposal

Our proposal is to establish a cloud-provider-vsphere subproject under VMware SIG and create a repo under kubernetes/cloud-provider-vsphere to host the CCM code for vSphere.

The plan of action includes 

- Scheduling a weekly meeting to discuss and manage the backlog
- Establish relationship with the SIGs that own and participate in [KEP0002](https://github.com/kubernetes/community/blob/master/keps/0002-controller-manager.md): apimachinery, storage, apps, network.
- Participate in the Kubernetes Cloud Provider Refactoring working group to update on the status of the vSphere transition. 
- Scope the work required to create a Cloud Controller Manager for vSphere
- Scope the impact to storage and network interfaces (like CSI)
- Define targets and milestones
- Refactor in-tree cloud provider code to support migration to Cloud Controller Manager
- Reach functional parity between in-tree provider and CCM
- Have E2E test framework in place for CCM
- Move code from in-tree provider to kubernetes/cloud-provider-vsphere

### Graduation to Beta

As part of the graduation to `stable` or General Availability (GA), we will require to reach functional parity between in-tree provider and CCM.