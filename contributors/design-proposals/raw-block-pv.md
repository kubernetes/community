# Local Raw Block Consumption via Persistent Volume Source

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume source API as a consistent model of consumption.

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
  definitions for the block devices will be driven through the PVC and PV and Storage Class definitions. Along with Storage
  Resource definitions, this will provide the admin with a consistent way of managing all storage. 
  The API changes proposed in the following section are minimal with the idea of defining a volumeType to indicate both the definition
  and consumption of the devices. Since it's possible to create a volume as a block device and then later consume it by provisioning
  a filesystem on top, the design requires explicit intent for how the volume will be used.
  The additional benefit of explicitly defining how the volume is to be consumed will provide a means for indicating the method
  by which the device should be scrubbed when the claim is deleted, as this method will differ from a raw block device compared to a 
  filesystem. The ownership of scrubbing the device properly shall be up to the plugin method being utilized.
  The last design point is block devices should be able to be fully restricted by the admin in accordance with how inline volumes 
  are today. Ideally, the admin would want to be able to restrict either raw-local devices and or raw-network attached devices.
  
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
compatibility, the absence of volumeType will default to volumes work today, which are formatted with a filesystem depending on 
the plug-in chosen. Recycling will not be a supported reclaim policy. Once the user deletes the claim against a PV, the volume will 
be scrubbed according to how it was bound. The path value in the PV definition would be overloaded to define the path of the raw
block device rather than the fileystem path.

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
**path: /dev/xvdc**
accessModes:
  - ReadWriteOnce
persistentVolumeReclaimPolicy: Delete 
```

## Storage Class API Changes:
For dynamic provisioning, it is assumed that values pass in the parameter section are opaque, thus the introduction of utilizing
fsType in the StorageClass can be used by the provisioner to indicate how to create the volume. The proposal for this value is
defined here:
https://github.com/kubernetes/kubernetes/pull/45345 
Therefore, a provisioner could potentially provision a block device and install the filesystem onto it by indicating the volumeType
as 'block' but the fsType as 'xfs'.

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
name: block-volume
provisioner: kubernetes.io/local-block-glusterfs
parameters:
**volumeType: block**
fsType: xfs
```
The provisioner should validate the parameters and return and error if the combination specified is not supported.
This also allows the use case for leveraging a Storage Class for utilizing pre-defined static volumes.

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
name: block-volume
provisioner: no-provisioning 
parameters:
**volumeType: block**
```
  
# Use Cases

## UC1: 

DESCRIPTION: 

A developer wishes to enable their application to use a local raw block device as the volume for the container. The admin has already created PVs that the user will bind to by specifying 'block' as the volume type of their PVC.

BACKGROUND:

For example, existing on-premise legacy applications such as database utilizes raw block device directly on their iSCSI or FC enterprise storage. This feature enables these users to move their application environment onto container environment by lift and shift approach with using Block Volumes Support and StatefulSets.

Examples:
*   MariaDB

WORKFLOW:

ADMIN:

Admin creates devices on disk. Admin creates PV using the path to specify the device location.

```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: local-raw-pv
spec:
  volumeType: block
  capacity:
    storage: 100Gi
  local:
    path: /dev/xvdc
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
```

USER:

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: local-raw-pvc
spec:
  volumeType: block
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi
```

```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers:
    - name: mysql
      image: mysql
      volumeMounts:
      - name: my-db-data
	mountPath: /var/lib/mysql/data
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: local-raw-pvc
```
  NOTE: *accessModes correspond to the container runtime values. Where RWO == RWM (mknod) to enable the device to be written to and
  create new files. (Default is RWM) ROX == R
  **(RWX is NOT valid for block and should return an error.)** * This has been validated among runc, Docker and rocket. 

## UC2: 

DESCRIPTION: 

A user uses a container on virtual machine on hypervisor such as KVM, VMware and wishes to use raw block device for databases such as MariaDB. 

BACKGROUND:

Using hypervisor's device passthrough feature(KVM: device passthrough, VMware: RawDeviceMapping), storage admin can expose raw block device into virtual machine then user can consume it via PV and PVC. Also dynamic provisioning could work if external-provisioner support this.

WORKFLOW:

ADMIN:

* Admin creates a disk and attach it to server and pass-through the disk to KVM guest which Kubelet node1 is working inside.
* Admin creates a PV using the path to specify the device location.

```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: raw-pv
spec:
  volumeType: block
  capacity:
    storage: 100Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  fc:
    targetWWNs: ['500a0982991b8dc5']
    lun: 2
```

* Admin adds "raw-disk" label to the kubelet node1

```
% kubectl label nodes <node1-ip> type=raw-disk
```

USER:

* User creates a persistent volume claim with volumeType: block option.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: raw-pvc
spec:
  volumeType: block
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi
```

*   User creates a Pod yaml which uses raw-pvc PVC and selects a node that Admin attached raw disk using nodeSelector option.

```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers
    - name: mysql
      image: mysql
      volumeMounts:
      - name: my-db-data
	mountPath: /dev/xvda
    nodeSelector:
      type: raw-disk
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: raw-pvc
```

## UC3: 

DESCRIPTION: An admin wishes to pre-create a series of raw block devices to expose as PVs for consumption. The admin wishes to specify the purpose of these devices by specifying 'block' as the volumeType for the PVs.

WORKFLOW:

ADMIN:

```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: local-raw-pv
spec:
  volumeType: block
  capacity:
    storage: 100Gi
  local:
    path: /dev/xvdc
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
```

## UC4: 

DESCRIPTION: StorageClass with non-dynamically created volumes

BACKGROUND: The admin wishes to create a storage class that will identify pre-provisioned block PVs based on a user's PVC request for volumeType: Block. 

WORKFLOW: 

ADMIN:

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: block-volume
provisioner: no-provisioning 
parameters:
  volumeType: block

```
* Sample of pre-created volume definition:

```
apiVersion: v1
kind: PersistentVolume
metadata:
 name: pv-block-volume
spec:
 volumeType: block
 storageClassName: block-volume
 capacity:
   storage: 35Gi
 accessModes:
   - ReadWriteOnce
 local:
    path: /dev/xvdc

```
## [FUTURE] UC5: 

DESCRIPTION: StorageClass with dynamically created volumes 

BACKGROUND: The admin wishes to create a storage class that will dynamically create block PVs based on a user's PVC request for volumeType: Block. The admin desires the volumes be created dynamically and deleted when the PV definition is deleted. 

WORKFLOW:

ADMIN:

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: local-fast
provisioner: kubernetes.io/local-block-ssd
parameters:
  volumeType: block

```

***This has implementation details that have yet to be determined. It is included in this proposal for completeness of design ****

## UC6: 

DESCRIPTION: The developer wishes to request a block device via a Storage Class.

WORKFLOW:

USER: 

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
 name: pvc-local-block
spec:
 volumeType: block
 storageClassName: local-fast
 accessModes:
  - ReadWriteOnce
 resources:
   requests:
     storage: 10Gi

```  

## UC7:

DESCRIPTION: User wishes to install a software defined storage solution  (i.e.GlusterFS/Ceph(bluestore)) on top of raw block device. Using something like a StateFulSet for GlusterFS, for example, the system could provide it with provisioned block storage, and allow the software defined storage solution to consume it as needed. 

BACKGROUND: This provides a way to dynamically create raw block devices to later be consumed by software defined storage.

WORKFLOW:

ADMIN:

```
kind: StorageClass
apiVersion: storage.k8s.io/v1beta1
metadata:
  name: aws-ebs-raw
provisioner: kubernetes.io/aws-ebs-raw
parameters:
  type: gp2
  fstype: raw 
```

* Storage Class Stateful Set

```
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 2
  template:
    metadata:
      labels:
	app: nginx
    spec:
      containers:
      - name: nginx
	image: gcr.io/google_containers/nginx-slim:0.8
	ports:
	- containerPort: 80
	  name: web
	volumeMounts:
	- name: www
	  mountPath: /usr/share/nginx/html
 volumeClaimTemplates:
  - metadata:
      name: datadir
      annotations:
	volume.beta.kubernetes.io/storage-class: aws-ebs-raw
    spec:
      accessModes:
	- "ReadWriteOnce"
      resources:
	requests:
	  storage: 10Gi
```
***SUITABLE FOR: NETWORK ATTACHED BLOCK***

## UC8: 

DESCRIPTION: Admin creates network raw block devices

BACKGROUND: Admin wishes to pre-create Persistent Volumes in GCE as raw block devices

WORKFLOW:

ADMIN:

```
apiVersion: "v1"
kind: "PersistentVolume"
metadata:
  name: gce-disk-1
  annotations:
    volume.beta.kubernetes.io/mount-options: "discard"
Spec:
  volumeType: block
  capacity:
    storage: "10Gi"
  accessModes:
    - "ReadWriteOnce"
  gcePersistentDisk:
    fsType: "raw"
    pdName: "gce-disk-1"

```

***If admin specifies volumeType: block + fstype: ext4 then they would get what they already get today ***

## UC9: 

DESCRIPTION: Developer wishes to consumes raw device

WORKFLOW:

USER:

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-raw-block
spec:
  volumeType: block
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi
```

```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers:
    - name: mysql
      image: mysql
      volumeMounts:
      - name: my-db-data
	mountPath: /var/lib/mysql/data
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: pvc-raw-block
```

# Implementation Plan

Phase 1: Pre-provisioned PVs to precreated devices

               API changes

               Special considerations: Access, Scrubbing device after PV is deleted
               
               Changes to the mounter interface as today it is assumed 'file' as the default.

Phase 2:  Discovery of block devices 

                Dynamically provisioned PVs to dynamically allocated devices

                Priv container concerns    

Other considerations:

               Expose volumeType to users via kubectl

               Adds enable/disable configuration to securityContext

               Reference volume driver change(Attach/Detach logic) for pre-provisioned PVs

               Reference volume driver change(Provision/Delete logic) for dynamic provisioning

## Mounter interface proposed design

```
 type BlockMounter interface {
	Mounter
 	CanSupportFileOnBlock() error
            // kubelet needs volume's device path when attaches a volume to container
 	GetVolumePath() string
 	GetVolumeType() string   //TBD
 }

type BlockUnmounter interface {
 	Unmounter
 	GetVolumePath() string
}
```
