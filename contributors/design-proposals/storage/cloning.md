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
* Host-Assisted Clone - using a pod to copy data between two persistent volumes

## Goals
* Provide consistent process to execute cloning for storage technologies both for intree and out of tree plugins

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
* An admin wishes to define a storage class that will allow the user to easily request a clone a volume without knowledge
* An admin wishes to restrict what namespaces can clone which images

## Design Overview
The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The cloning will be initiated 
by creating an annotated PVC to create a new volume using a defined storage class that supports cloning.
The actual cloning process must be implemented by the provisioning storage and properly understand and react to the annotation to clone. In the absence of a storage class that supports cloning a host-assisted clone can also be exected to create the copy. The host-assisted cloning deploys a controller to utlize pod streaming to execute the copy.

### API
Leveraging the proposed API change for 'dataSource' the PVC will make a clone request by specifying the storage class and the location of the object it wishes to clone. This this example, we are cloning a pvc. For version 1 of cloning, we will limit cloning of pvcs to the current namespace only.

To create a clonable pvc, the volume that is bound to the pvc must be marked as 'canClone' to indicate cloning of the volume contents is permitted.

** Pre-Request:
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-1
spec:
  capacity:
    storage: 100Gi
  volumeMode: Filesystem
annotations:
    k8s.io/CloneRequest: canClone  
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
    kind: PersistentVolumeClaim
    name: pvc-1
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
parameters:
  smartclone: "true"
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
  annotations:
    k8s.io/CloneOf: pvc-1
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
    kind: PersistentVolumeClaim
    name: pvc-1
  resources:
    requests:
      storage: 10Gi
  annotations:
    k8s.io/CloneOf: pvc-1
```   
   
      
