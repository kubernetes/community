---
kep-number: 31
title: Add ability to transfer PVC between Namespaces
authors:
  - "@j-griffith"
owning-sig: sig-storage
participating-sigs:
  - sig-architecture
reviewers:
  - TBD
approvers:
  - TBD
editor: @j-griffith
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: provisional
---

# Kubernetes PVC Namespace Transfer


## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.
[Tools for generating][] a table of contents from markdown are available.

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

Provide a mechanism to allow users to "give" a PVC to a user in another namespace based on mutual acceptance. A transfer should be initiated by updating an existing PVC with a transfer request annotation, and accepted by the intended destination namespace by the recipient requesting a PVC using the standard mechanism and specifying the correct PVC Name and transfer annotation as per the transfer request.

* Provide a Cluster wide mechanism to enable or disable the transfer of PVCs
* Expose standardized process for transferring a PVC from one Namespace to another
* Require explicit acceptance of a PVC into another users Namespace, ensuring a malitious user can't just dump PVCs into another users Namespace
* Utilize existing PVC creation semantics for creating a PVC from a Transfer
* Honor existing configuration options including quota checks, Storage Class access etc

The process of transferring a PVC/Volume is as follows:
1. Original user indicates they're willing to ``give`` the volume to another namespace
2. The receiving user indicates they'd like to ``accept`` the volume into their namespace
3. When the original user deletes the PVC from their namespace, the claim request from the receiving user is fulfilled

## Motivation

There are a number of uses cases where a user would like to have the ability to transfer an existing PVC from one namespace to another.  This is a valuable workflow for persistent storage and enables the ability to easily duplicate and transfer data sets from one environment to another.  These populated PVCs could be a clone of another volume, a volume from snapshot, or data that was written to the volume via an application (ie a database).

An example use case for this feature would be a cluster segmented into two namespaces; namespace-a for production staging, and namespace-b for production.  There are cases where an application could be developed and tested with the same production data without risking any modification or corruption of data in both environments.  Rather than reproducing the data in both namespaces, it would be much more efficient to be able to clone or restore the data from a snapshot in to a volume and then transfer that new volume to the desired namespace.

### Goals

* Provide a secure and standard mechanism to duplicate data sets on volumes across namespaces in the cluster
* Enhance existing API's and workflows to alow transferring of volumes across namespaces

### Non-Goals

  While this concept could be extended, volumes are a special case.  Most other resources can just be recreated easily, but this isn't as trivial with large data sets.
* Introducing a new API
* Introducing a new Transfer Object
* Enable transfer for resources other than volumes

## Proposal


Update the PVC with a transfer request annotation:
```yaml
apiversion: v1
kind: PersistentVolumeClaim
metadata:
    name: pvc-foo
	annotations:
	    transferRequest: development/pvc-fromfoo
...
```

The transferRequest update, will be detected by the PV controller, if the transfer feature is enabled, the controller will mark the PV backing the PVC for transfer.  This notation is used when the PVC is deleted from the original namespace.  The transfer request overrides the reclaim policy of the PV, and upon deletion of the claim from the original namespace, the claimref is updated to point to the new namespace/pvc-name, and the PV is not deleted.

The result is annotations added to the PV and PVC by the controller:
```yaml
apiversion: v1
kind: PersistentVolume
metadata:
    name: pvc-fromfoo
	annotations:
	    pv.kubernetes.io/transfer-status: pending
		pv.kubernetes.io/transfer-source: default/pvc-foo
		pv.kubernetes.io/transfer-destination: development/pvc-fromfoo
...
```

```yaml
apiversion: v1
kind: PersistentVolumeClaim
metadata:
    name: pvc-fromfoo
	annotations:
	    transferAccept: default/pvc-foo
	    pvc.kubernetes.io/transfer-status: pending
...
```

The transfer-status will remain in a ``pending`` state until:
1. The original user deletes their claim
2. The destination user requests a claim with the proper name and annotations

User in the development namespace requests a claim that matches the originators transfer request:
```yaml
apiversion: v1
kind: PersistentVolumeClaim
metadata:
    name: pvc-fromfoo
	annotations:
	    transferAccept: default/pvc-foo
...
```

The claim request in the destination namespace will be ``pending`` until the source user deletes claim ``pvc-foo`` from their namespace.  The source user can also delete the ``pvc-foo`` from their namespace before the claim request is made in the destination namespace as well.   Note that the sequence described preserves integrity of our users quotas on both sides, and even if a user in the reveiving namespace requests a claim with the names specified in a transfer request, the claim is ONLY fulfilled by that backing PV is they specified a ``transferAccept``.  This prevents accidental acceptance of transfers.

Upon success, the PV annotation: ``pv.kubernetes.io/transfer-status`` annotation will be updated by the controller to ``complete``.

### User Stories [optional]

#### Story 1

As a cluster user, I want to be able to clone a volume and use that volume in another namespace.

#### Story 2

As a cluster user, I want to be able to create a volume from a snapshot, and use that volume in another namespace.

#### Story 3

As a cluster user, I want to be able to populate a volume with data from my own application, and use that volume in another namespace.

### Implementation Details/Notes/Constraints [optional]

The main premise of this proposals implementation is to leverage existing user policies around storage, like quotas, enabled usage of storage classes etc.  It also seemed important to minimize changes to the API and the existing workflow.  This proposal attempts to keep the process of creating PVCs the same, regardless of whether it's a transfer or not.

The use of annotations may or may not be the most desirable method of exposing this to end users.  It's advantageous because it provides flexibility without breaking the API, but it's also less desirable in terms of the ability to version it and ensure breaking changes aren't introduced in subsequent Kubernetes releases.

A simple POC implementation to test the basics of the concept is available on this [GitHub branch](https://github.com/j-griffith/kubernetes/tree/pvc-transfer-poc), and a [screencast demo](https://asciinema.org/a/210854) is available on asciinema.

### Risks and Mitigations

This proposal attempts to leverage existing security mechanisms around storage, but there are still risks.  The biggest risk in the general idea is users 'dumping' volumes into other users namespaces.  The give/take handshake mechanism is the proposed mitigation for that.  Given that we require explicit acknowledgement on both the source and destination namespace of a transfer we're assuming a level of trust between users/namespaces in the cluster.  We're also providing a mechanism to disallow this on a cluster wide basis.

## Graduation Criteria

* Acceptance of enhancement proposal
* Acceptance/merge of changes to the PV controller to enable the new behavior

## Implementation History

## Drawbacks [optional]


## Alternatives [optional]

* Transfer data via external applications and use things like Object Stores as data transfer vehicles
* Implement a new "transfer" controller specifically for handling transfer of resources

## Infrastructure Needed [optional]

