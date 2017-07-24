# Local Raw Block Consumption via Persistent Volume Source

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume soruce API as a consistent model
of consumption.

# Goals
* Enable durable access to block storage
* Support storage requirements for all workloads supported by Kubernetes
* Provide flexibility for users/vendors to utilize various types of storage devices
* Agree on API changes for block
* Provide a consistent security model for block devices 
* Provide block storage usage isolation
* Provide a means for running containerized block storage offerings as non-privileged container

# Non Goals
* Support all storage devices natively in upstream Kubernetes. Non-standard storage devices are expected to be managed using extension
  mechanisms.
   
# Value add to Kubernetes

  Before the advent of storage plugins, emptyDir and hostPath were widely used to quickly prototype stateless applications in Kube. 
  Both have limitations for their use in application that need to store persistent data or state. 
  EmptyDir, though quick and easy to use, provides no real guarantee of persistence for suitable amount of time. 
  Appropriately used as scratch space, it does not have the HostPath, became an initial offering for local storage, but had many
  drawbacks. Without having the ability to guarantee space & ensure ownership, one would lose data once a node was rescheduled. 
  Therefore, the risk outweighed the reward when trying to  leverage the power of local storage needed for stateful applications like 
  databases.
    
# Design Overview

  The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The consumption and 
  defintions for the block devices will be driven through the PVC and PV and Storage Class definitions. Along with Storage
  Resource definitions, this will provide the admin with a consistent way of managaging all storage. 
  
  The API changes proposed in the following section are minimal with the idea of defining a volumeType to indicate bothe the defintion
  and consumption of the devices. Since it's possible to create a volume as a block device and then later consume it by provisioning
  a filesystem on top, the design requires explict intent for how the volume will be used.
  
  This intentional use direction also helps us drive the means by how the device was utilized and provide a proper scrubbing of the 
  device as the mechinism for this will differ depending on whether it has a filesystem or not.
  
# Proposed API Changes
   
  ## Persistent Volume Claim API Changes:
  In the simplest case of static provisioning, a user asks for a volumeType of block. The binder will only bind to a PV defined 
  with the same label.
  
  ```
  kind: PersistentVolumeClaim
  apiVersion: v1
  metadata:
    name: myclaim
  spec:
    **volumeType: block**
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: 80Gi
  ```
 For dynamic provisioning and the use of the storageClass, the user also specifically defines the intent of the volume by 
 indicating the volumeType as block. The provisioner for this class will validate whether or not it supports block and return
 an error if it does not.
  
  ```
  kind: PersistentVolumeClaim
  apiVersion: v1
  metadata:
    name: myclaim
  spec:
    **storageClassName: local-fast**
    **volumeType: block**
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: 80Gi
    ```
   ## Persistent Volume API Changes:
  For static provisioning the admin creates the volume and also is intentional about how the volume should be consumed. For backwards
  compatibility, the absense of volumeType will default to volumes work today, which are formatted with a filesystem depending on 
  the plug-in chosen
  
   ```
   kind: PersistentVolume
   apiVersion: v1
   metadata:
     name: local-raw-pv
   spec:
     **volumeType: block**
     capacity:
       storage: 100Gi
     local:
       path: /dev/xvdc
     accessModes:
       - ReadWriteOnce
     persistentVolumeReclaimPolicy: Delete
   ```
   
 For dynamic provisioning and the use of the storageClass, the user also specifically defines the intent of the volume by 
 indicating the volumeType as block. The provisioner for this class will validate whether or not it supports block and return
 an error if it does not.
  
        ```
  kind: PersistentVolumeClaim
  apiVersion: v1
  metadata:
    name: myclaim
  spec:
    **storageClassName: local-fast**
    **volumeType: block**
    accessModes:
      - ReadWriteOnce
    resources:
      requests:
        storage: 80Gi
  
[TBD]

