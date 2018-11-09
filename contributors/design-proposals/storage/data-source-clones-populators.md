 DataSource Proposal
================================

**Authors:** [John Griffith](https://github.com/j-griffith)

## Background
This document proposes the extension of the DataSources field for PVC creation in Kubernetes.  Snapshots as a DataSource have already been added to Kubernetes, enabling a user to create a new PVC based off of a specified Snapshot.  There are other common DataSources that can be used for PVC creation, including Cloning and Populators.

DataSources in general provides a clean abstraction to expose these other storage device related features (Cloning) and Services that retrieve data from remote sources implemented as CRDs (Populators).

## Terminology
* Clone - a duplicated volume created by the same storage technology as the original
* Snapshot - a point in time copy of the data the resides on a PVC, often used as a backup; the Snapshot is NOT a PVC but is it's own unique object.
* Populator - a CRD used to write data to a PVC, this includes sources like tar files which can be transferred to and expanded on a PVC during the creation process.

## Goals
* Enable common features provided by storage devices to provide better performance, automation and disaster recovery
* Utlize the DataSources field for standard supported features like cloning and snapshot, but also allow extensibility for external Controllers (Populators)
* In all cases, provide a single workflow for creating a PVC pre-populated with data, resulting in an independent PVC object

## Non Goals
* This proposal does NOT suggest implementing DataSources (Cloning or Populators) in Kubernetes.  The goal is instead to just expose these features in the API so that they can be used by Kubernetes users if they're available.

## Value add to Kubernetes

Clones, Snapshots and Populators are commonly used in storage to automate the process of providing specific/expected data on a Volume.  Most backend devices provide optimizations for things like Clones and Snapshots, and recommend their usage as best practices.  Adding the ability to use these features in Kubernetes provides Kubernetes users with the features they're already using and expecting.

The added ability to specify Populators implemented via CRD adds another level of convenience and standardization for users to create PVCs pre populated with data.

## Use Cases
Specific use cases around cloning are included in the use cases listed below as follows:
* A user wishes to create a duplicate environment for debug purposes
* A user wishes to use Clones or Snapshots as part of their DR strategy
* A user wishes to install data from an http, S3 or other source to their volume via a CRD

## Design Overview

The propsed design is to expand the DataSource Field that already exists in the PVC object.  Currently, the only accepted use of DataSource is to specify a Snapshot to create a PVC from.  This design would extend the DataSource field to accept:
* Other PVCs in the user namespace (Clones) for SC's that support it
* Populators which would take the form of an external controller or CRD on those systems that have Populator CRDs registered

It's important to note that the process of Cloning, Creating from Snapshot or Populating is NOT implemented in Kubernetes but is a capability of the Storage Device, or provided via an external controller.

#### Clone

The following example shows what a request to create a Clone of an existing PVC would look like.  This assumes that there is only one Storage Class on the system and that the Storage Class supports cloning (CSI Capabilities):

** Request:

```
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


** Result:
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
    name: pvc-2
    namespace: myns
spec:
    accessModes:
    - ReadWriteOnce
    resources:
    requests:
        storage: 10Gi
```
Where pvc-2 is a clone ov pvc-1 and is a new independent object, with it's own provisioned backend volume.  This relies upon the use of a CSI Provisioner, with a Storage Class that supports cloning capabilities.  If the selected Storage Class does NOT report support for Cloning in it's capabilities, then the CSI Provisioner should return an error for this request.  If the CSI Plugin DOES report support of cloning in it's capabilities, then it's up to the CSI External Provisioner to form the proper DataSource options in the request that is issued to storage plugin as per the CSI Specification.

#### Populator

The following example shows what a request to create a PVC from a populator would look like.  This assumes that the specified populator CRD has been deployed on the Cluster.  Our example uses a CRD `populator.acme.io` that knows how to handle requests of kind `HTTPSource`.

** Request:

```
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
** Result:
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-2
  Namespace: myns
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```

## Additional Requirements

This proposed expansion of the DataSource field in the PVC Object, is dependent upon work being proposed for a ReadinessGate.  The requirement is for a way to describe that a PVC is created on a backend, but that it may require some operations (ie populating from a Snapshot, or a Populator CRD) before it is actually available for use.  There is currently ongoing discussions to form a proposal for this feature as it pertains to Snapshots.  Ideally, that feature would be designed with the flexibility to be used for any DataSource and not just Snapshots.

