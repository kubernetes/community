Kubernetes Cloning Proposal
================================

**Authors:** [Erin Boyd](https://github.com/erinboyd) [Adam Litke](https://github.com/aglitke)

## Background

This document presents a proposal for defining best practices around security and implementation of cloning of persistent volumes in Kubernetes. 
Cloning is distinctly different than storage snapshots in that it leverages the specific storage technology to ‘clone’ the physical volume 
rather than copy the contents from one created PV to another. It also it not intended to capture a portion of data in time, 
but always provide a full and complete copy.

## Terminology
* Clone - a duplicated volume created by the same storage technology as the original
* Cloning - the process by which a storage technology can create a full copy of a volume

## Goals
* Provide consistent process to execute cloning for storage technologies both for out of tree plugins.

## Non Goals
* The process should not force storage vendors to implement cloning
* Provide governance for security to clone assets between namespaces

## Value add to Kubernetes  
 By providing a consistent method for cloning, users can leverage the native capability of their storage to quickly create a 
 copy of their persistent volumes without having to wait for a snapshot object to be created. This is especially desirable for large 
 data volumes. Cloning also is a common use case in disaster recovery and for pre-seeding storage in a virtualized environment that many 
 users accustomed to operating in those environments might find useful. As plug-ins move out of tree, it will be more difficult to 
 enforce conformance to best practices. By creating an accepted framework for this process, we can ensure it’s done securely regardless 
 of the storage provider.

## Use Cases 
Specific use cases around cloning are included in the use cases listed below as follows:
* An admin wishes to create a copy of a volume to spawn another container to use the same data. 
* An admin wishes to create a copy of a volume to pre-seed a process with information the container might need for bootstrapping.
* A user wishes to clone a volume for purposes of disaster recovery

## Design Overview
The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The cloning will be initiated 
by creating an PVC to create a new volume using the dataSource API created with snapshots.
The actual cloning process must be implemented by the provisioning storage and properly understand and react to the request to clone. 

### API
Leveraging the proposed API change for 'dataSource' the PVC will make a clone request by specifying the object it wishes to clone. In this example, we are cloning a pvc. For version 1 of cloning, we will limit cloning of pvcs to the current namespace only.

** Pre-Request:
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-1
  Namespace: myns
spec:
  capacity:
    storage: 100Gi
  volumeMode: Filesystem
accessModes:
    - ReadOnlyMany
  persistentVolumeReclaimPolicy: Delete
  storageClassName: csi-gce-pd
```
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-1
  Namespace: myns
spec:
  capacity:
    storage: 100Gi
```

** Request:
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: clone-pvc
  Namespace: myns
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: csi-gce-pd
  dataSource:
    kind: PersistentVolume
    name: pv-1
  resources:
    requests:
      storage: 10Gi
```
```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: csi-gce-pd
provisioner: kubernetes.io/gce-pd
```

** Result:
```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: pv-2
  Namespace: myns
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: csi-gce-pd
  resources:
    requests:
      storage: 10Gi
```
```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: clone-pvc
  Namespace: myns
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: csi-gce-pd
  dataSource:
    kind: PersistentVolume
    name: pv-1
  resources:
    requests:
      storage: 10Gi
```   

### CSI Spec
This PR does not propose an API change, however the CSI specification will need to be updated as follows to add an additional
option for the VolumeContentSource:

``` 
// Specifies what source the volume will be created from. One of the
// type fields MUST be specified.
message VolumeContentSource {
  message SnapshotSource {
    // Contains identity information for the existing source snapshot.
    // This field is REQUIRED. Plugin is REQUIRED to support creating
    // volume from snapshot if it supports the capability
    // CREATE_DELETE_SNAPSHOT.
    string id = 1;
  }
  message PersistentVolumeSource {
    // Contains identity information for the existing volume ID.
    // This field is REQUIRED. Plugin is REQUIRED to support creating
    // volume from clone if it supports the capability
    string id = 1;
  }
``` 

Care has been taking to ensure this design follows the groundwork laid by the following snapshot proposals:
https://github.com/container-storage-interface/spec/pull/244 

https://github.com/kubernetes-csi/external-provisioner/pull/123 
