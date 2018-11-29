---
kep-number: 34
title: aws-ebs-csi-driver
authors:
  - "@leakingtapan"
owning-sig: sig-aws
reviewers:
  - "@d-nishi"
  - "@jsafrane"
approvers:
  - "@d-nishi"
  - "@jsafrane"
editor: TBD
creation-date: 2018-11-27
last-updated: 2018-11-27
status: provisional
---

# AWS Elastic Block Store (EBS) CSI Driver

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
        * [Static Provisioning](#static-provisioning)
        * [Volume Schduling](#volume-scheduling)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary
AWS EBS CSI Driver implements [Container Storage Interface](https://github.com/container-storage-interface/spec/tree/master) which is the standard of storage interface for container. It provides the same in-tree AWS EBS plugin features including volume creation, volume attachment, volume mounting and volume scheduling. It is also configurable on what is the EBS volume type to create, what is the file system file should be formatted, which KMS key to use to create encrypted volume, etc.

## Motivation
Similar to CNI plugins, AWS EBS CSI driver will be a stand alone plugin that lives out-of-tree of kuberenetes. Being out-of-tree, it will be benefit from being modularized, maintained and optimized without affecting kubernetes core code base. Aside from those benefits, it could also be consumed by other container orchestrators such as ECS.

### Goals
AWS EBS CSI driver will provide similar user experience as in-tree EBS plugin:
* As an application developer, he will not even notice any difference between EBS CSI driver and in-tree plugin. His workflow will stay the same as current.
* As an infrastructure operator, he just need to create/update storage class to use CSI driver to manage underlying storage backend.

List of driver features include volume creation/deletion, volume attach/detach, volume mount/unmount, volume scheduling, create volume configurations, volume snapshotting, mount options, raw block volume, etc.

### Non-Goals
* Supporting non AWS block storage
* Supporting other AWS storage serivces such as Dynamodb, S3, etc.

## Proposal

### User Stories

#### Static Provisioning
Operator creates a pre-created EBS volume on AWS and a PV that refer the EBS volume on cluster. Developer creates PVC and a Pod that uses the PVC. Then developer deploys the Pod during which time the PV will be attached to container inside Pod after PVC bonds to PV successfully.

#### Volume Scheduling
Operation creates StorageClass with  volumeBindingMode = WaitForFirstConsumer. When developer deploys a Pod that has PVC that is trying to claim for a PV, a new PV will be created, attached, formatted and mounted inside Pod&#39;s container by the EBS CSI driver. Topology information provided by EBS CSI driver will be used during Pod scheduling to guarantee that both Pod and volume are collocated in the same availability zone.

### Risks and Mitigations
* *Information disclosure* - AWS EBS CSI driver requires permission to perform AWS operation on users&#39; behave. EBS CSI driver will make sure non of credentials are logged. And we will instruct user to grant only required permission to driver as best securtiy practise.
* *Escalation of Privileges* - Since EBS CSI driver is formatting and mounting volumes, it requires root privilege to permform the operations. So that driver will have higher privilege than other containers in the cluster. The driver will not execute random command provided by untrusted user. All of its interfaces are only provided for kuberenetes system components to interact with. The driver will also validate requests to make sure it aligns with its assumption.

## Graduation Criteria
AWS EBS CSI driver provides the same features as in-tree plugin.

## Implementation History
* 2018-11-26 Initial proposal to SIG
* 2018-11-26 Initial KEP draft
* 2018-12-03 Alpha release with kuberentes 1.13

