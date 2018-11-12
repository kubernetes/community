---
kep-number: 0
title: Extend usage of Volume DataSource field
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
creation-date: 2018-11-11
last-updated: 2018-11-11
status: provisional
see-also:
replaces:
superseded-by:
---

# Allow the use of the dataSource field for clones and external populators

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

This KEP proposes expanding the allowed uses of the DataSource field for PVC creation in Kubernetes to include PVCs and external Populator CRDs.  Currently we allow Snapshots, which is used by the CSI provisioner to enable creating volumes from snapshots.  In addition to Snaphsots however, the CSI Spec includes the ability to create "clones" (creating a new duplicate volume from an existing volume), in which case the DataSource is an existing Volume on the backend device.  Along with the existing Clone case, the other proposed inclusing would be for a new type of CRD labeleled as a ``Populator``.  Populators are implemented as external CRDs (similar to the Snaphsot CRD) and are comprised of applications that attach a newly created PVC to a POD and populate it with data.

## Motivation

Features like Cloning are common in almost all modern day storage devices, not only is the capability available on most devices, it's also frequently used in various use cases whether it be for duplicating data or to use as a disaster recovery method.  In addition to Clones, one other common process used by storage consumers is to "pre populate" volumes with data, this could be things like restoring a database backup, or seeding a specific data set prior to processing.  In this case it's currently up to the user to implement the population process using their own applications, the concept of an external populator however would greatly simplify this process for the user.

### Goals

* Extend the existing DataSource field for PVCs to allow specifying existing PVCs on the system
* Extend the existing DataSource field for PVCs to allow specifying external Populator CRDs that are registered on the cluster


### Non-Goals

* The actual clone logic is NOT proposed to be implemented in Kubernetes.  This proposal simply proposes the ability to communicate to the CSI Provisioner that we want to perform a clone, it's up to the CSI Plugin/Backend device to handle the implementation.
* The logic for a Populator would NOT be implemented in Kubernetes.  This would just be a mechanism to indicate that a user would like to use a Populator that's registered on the Cluster as an external CRD.

## Proposal

Modify the checks on the dataSource field to allow:
1. Existing PVCs in the users namespace (clones).
2. External Populators in the form of registered CRDs

In both cases the proposal is just to allow communication of intent, there's no implementation proposed in Kubernetes.  The Clone case is fairly straight forward given that the CSI spec already includes volumes as a data source and has a cloning capability that plugins are required to report:

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
    name: pvc-2
    namespace: myns
spec:
  capacity:
    storage: 10Gi
  dataSource:
    kind: PersistentVolumeClaim
    name: pvc-1
```

The result would be a new and independent PV/PVC (pvc-2) on the backend device that is a duplicate of the data that existed on pvc-1.

In the case of using a Populator, we'd rely on the use of the APIGroup field in the dataSource and the kind would depend upon object definitions in the selected APIGroup.

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-2
  Namespace: myns
spec:
  capacity:
    storage: 10Gi
  dataSource:
    APIGroup: populator.acme.io
    kind: HTTPSource # The kind HTTPSource is defined by the specified Populator CRD
    name: go1.11.2.tar.gz
-------------------------
kind: HTTPSource
metadata:
  name: my-data
  url: https://github.com/golang/go/archive/go1.11.2.tar.gz
  secretRef: "" #optional
```

The result would be a new and independent PV/VPVC (pvc-2) on the backend device that includes the data as provided by the external Populator.

### User Stories [optional]

Detail the things that people will be able to do if this KEP is implemented.
Include as much detail as possible so that people can understand the "how" of the system.
The goal here is to make this feel real for users without getting bogged down.

#### Story 1
As a cluster user, I want to be able to Clone a volume and run a different set of PODs/Applications against it

#### Story 2
As a cluster user, I want to be able to populate a newly created PVC with data from a source (ie tar file, release binary over HTTPS or other protocol) before using it in my application

### Implementation Details/Notes/Constraints [optional]

This proposal is intended to closely follow the work being proposed for Snapshots.  One of the key components needed still for Snapshots (and for this proposal) is the implementation of a readiness gate, whereby we can indicate that a PVC has been created, but is not ready for use.  This proposal is still under discussion and design, final completion of this KEP would depend upon that feature.

Along with the minor changes to Kubernetes, there will also be an implementation added to the CSI Provisioner to read the DataSource entry and in the case of clone, conduct the proper checks for things like:
* Storage classes match (src and destination)
* The storage class capabilities report support for cloning

On the topic of Populators, a reference CRD to handle common cases like HTTP transfers will be submitted as an external controller.  This should almost mirror the behavior and usage of the external snapshot controller.  We'd like to provide a community supported base Populator, but also allow users to write their own CRDs, register and use them.

### Risks and Mitigations

The most signfiicant risk for this proposal is the readiness gate component.  If there's not an accepted solution to that proposal it might be possible to introduce and use a concept of taints/tolerations for PVCs to achieve the same result.

## Graduation Criteria

There are a number of things that have to happen even after the use of DataSource is extended in order for this to be considered complete:
1. Cloning implementation in the CSI external provisioner
2. An external Populator CRD (similar to the Snapshot CRD)
3. A merged implementation of the readiness gate or some other equivalent substitute

## Implementation History

## Drawbacks [optional]

## Alternatives [optional]

Currently Cloning and Populators are implemented in an ad-hoc manner by various vendors/plugins through the use of annotations.  That does work, but it's inconsistent and is a less than ideal user experience.

## Infrastructure Needed [optional]

