# Raw Block Consumption in Kubernetes

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume source API as a consistent model of consumption.

# Terminology
* Raw Block Device - a physically attached device devoid of a filesystem
* Raw Block Volume - a logical abstraction of the raw block device as defined by a path
* File on Block/Filesystem on Block - a formatted (ie xfs) filesystem on top of a raw block device

# Goals
* Enable durable access to block storage
* Support storage requirements for all workloads supported by Kubernetes
* Provide flexibility for users/vendors to utilize various types of storage devices
* Agree on API changes for block
* Provide a consistent security model for block devices 
* Provide a means for running containerized block storage offerings as non-privileged container

# Non Goals
* Support all storage devices natively in upstream Kubernetes. Non-standard storage devices are expected to be managed using extension
  mechanisms.
   
# Value add to Kubernetes

  Before the advent of storage plugins, emptyDir and hostPath were widely used to quickly prototype stateless applications in Kubernetes. 
  Both have limitations for their use in application that need to store persistent data or state. 
  EmptyDir, though quick and easy to use, provides no real guarantee of persistence for suitable amount of time. 
  Appropriately used as scratch space, it does not have the HostPath, became an initial offering for local storage, but had many
  drawbacks. Without having the ability to guarantee space & ensure ownership, one would lose data once a node was rescheduled. 
  Therefore, the risk outweighed the reward when trying to  leverage the power of local storage needed for stateful applications like 
  databases.
  
  By extending the API for volumes to specifically request a raw block device, we provide an explicit method for volume comsumption,
  whereas previously it was always a fileystem. In addition, the ability to use a raw block device without a filesystem will allow
  Kubernetes better support of high performance applications that can utilitze raw block devices directly for their storage. 
  For example, MariaDB or MongoDB.
    
# Design Overview

  The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The consumption and 
  definitions for the block devices will be driven through the PVC and PV and Storage Class definitions. Along with Storage
  Resource definitions, this will provide the admin with a consistent way of managing all storage. 
  The API changes proposed in the following section are minimal with the idea of defining a volumeType to indicate both the definition
  and consumption of the devices. Since it's possible to create a volume as a block device and then later consume it by provisioning
  a filesystem on top, the design requires explicit intent for how the volume will be used.
  The additional benefit of explicitly defining how the volume is to be consumed will provide a means for indicating the method
  by which the device should be scrubbed when the claim is deleted, as this method will differ from a raw block device compared to a 
  filesystem. The ownership and responsibility of scrubbing the device properly shall be up to the plugin method being utilized. 
  By explicitly having volumeType in the PV and PVC can help the storage provider determine what scrubbing method to use. As an example,
  for local storage:
  
  PV = block, PVC = block: zero the bytes
  PV = block, PVC = file: destroy the filesystem
  PV = file, PVC = file: delete the files

  The last design point is block devices should be able to be fully restricted by the admin in accordance with how inline volumes 
  are today. Ideally, the admin would want to be able to restrict either raw-local devices and or raw-network attached devices.
  
# Proposed API Changes
   
## Persistent Volume Claim API Changes:
In the simplest case of static provisioning, a user asks for a volumeType of block. The binder will only bind to a PV defined 
with the same volumeType.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
name: myclaim
spec:
  volumeType: block #proposed API change
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi 
```

For dynamic provisioning and the use of the storageClass, the admin also specifically defines the intent of the volume by 
indicating the volumeType as block. The provisioner for this class will validate whether or not it supports block and return
an error if it does not.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
name: myclaim
spec:
  storageClassName: local-fast 
  volumeType: block #proposed API change
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi 
```

## Persistent Volume API Changes:
For static provisioning the admin creates the volume and also is intentional about how the volume should be consumed. For backwards
compatibility, the absence of volumeType will default to file which is how volumes work today, which are formatted with a filesystem depending on the plug-in chosen. Recycling will not be a supported reclaim policy. Once the user deletes the claim against a PV, the volume will be scrubbed according to how it was bound. The path value in the local PV definition would be overloaded to define the path of the raw block device rather than the fileystem path.
```
kind: PersistentVolume
apiVersion: v1
metadata:
name: local-raw-pv
spec:
  volumeType: block #proposed API change
  capacity:
    storage: 100Gi
  local:
    path: /dev/xvdc 
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
  volumeType: block #opaque value  / plug-in dependent
  fsType: xfs
```
The provisioner (if applicable) should validate the parameters and return and error if the combination specified is not supported.
This also allows the use case for leveraging a Storage Class for utilizing pre-defined static volumes. By labeling the Persistent Volumes
with the Storage Class, volumes can be grouped and used according to how they are defined in the class.
```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: block-volume
provisioner: no-provisioning 
parameters:
```
  
# Use Cases

## UC1: 

DESCRIPTION: An admin wishes to pre-create a series of local raw block devices to expose as PVs for consumption. The admin wishes to specify the purpose of these devices by specifying 'block' as the volumeType for the PVs.

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

## UC2: 

DESCRIPTION: 

A user uses a container on virtual machine on hypervisor such as KVM, VMware and wishes to use raw block device for databases such as MariaDB. 

WORKFLOW:

ADMIN:
* Admin creates a disk and exposes it to all kubelet worker node VMs which are running on KVM hypervisor.(This is done by storage operation).
* Admin creates an iSCSI persistent volume using storage information such as portal IP, iqn and lun.

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
  iscsi:
    targetPortal: 1.2.3.4:3260
    iqn: iqn.2017-05.com.example:test
    lun: 0
```

USER:

* User creates a persistent volume claim with volumeType: block option to bind pre-created iSCSI PV.

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

* User creates a Pod yaml which uses raw-pvc PVC.

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
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: raw-pvc
```
* During Pod creation, iSCSI Plugin attaches iSCSI volume to the kubelet worker node VM using storage information.

## UC3: 

DESCRIPTION: 

A developer wishes to enable their application to use a local raw block device as the volume for the container. The admin has already created PVs that the user will bind to by specifying 'block' as the volume type of their PVC.

BACKGROUND:

For example, an admin has already created the devices locally and wishes to expose them to the user in a consistent manner through the 
Persistent Volume API.

WORKFLOW:

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
    fsType: "block"
    pdName: "gce-disk-1"
```

***If admin specifies volumeType: block + fstype: ext4 then they would have the default behavior of files on block ***
***fsType values will be provisioner dependent. Block is suggested for development simplicity. Since the PVC object is passed
   to the provisioner, it will be responsible for validating and handling whether or not it supports the volumeType being passed ***

# Container Runtime considerations
It is important the values that are passed to the container runtimes are valid and support the current implementation of these various runtimes. Listed below are a table of various runtime and the mapping of their values to what is passed from the mount. 

| runtime engine    | runtime options  | accessMode       |  
| --------------    |:----------------:| ----------------:|
| docker/runc/rkt   |  mknod / RWM     | RWO              |
| docker/runc/rkt   |       R          | ROX              |

The accessModes would be passed as part of the options array and would need validate against the specific runtime engine. 
Since rkt doesn't use the CRI, the config values would need to be passed in the legacy method.

The runtime option would be placed in the DeviceInfo as such:
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "mrw"}) for RWO
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "mr"}) for ROX

Today, this is defaulted always to mrw, thus it would need to be updated with what is passed in.

# Implementation Plan, Features & Milesones

Phase 1: v1.8
Feature: Pre-provisioned PVs to precreated devices 

               Milestone 1: API changes

               Milestone 2: Restricted Access 
               
               Milestone 3: Scrubbing device after PV is deleted
               
               Milestone 4: Changes to the mounter interface as today it is assumed 'file' as the default.
               
               Milestone 5: Expose volumeType to users via kubectl
               
               Milestone 6: Adds enable/disable configuration to securityContext in PSP (Pod Security Policy) similar to hostPath
               
               Milestone 7: Validate container runtime options with user specifcations as indicated in UC3
	       
	       Milestone 8: Container Runtime changes

Phase 2:  v1.9
Feature: Discovery of block devices 

                Milestone 1: Dynamically provisioned PVs to dynamically allocated devices

                Milestone 2: Privileged container concerns    

Other considerations:

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
# Mounter binding matrix for statically provisioned volumes:

| PV volumeType | PVC volumeType  | Result           |
| --------------|:---------------:| ----------------:|
|   unspecified | unspecified     | BIND             |
|   file        | file            | BIND             |
|   block       | unspecified     | NO BIND          |
|   block       |  block          | BIND             |
|   unspecified | block           | NO BIND          |
|   block       |  file           | NO BIND          |
|   file        |  block          | BIND**           |
|  unspecified  | file            | BIND             |


* unspecified defaults to 'file/ext4' today for backwards compatibility and in mount_linux.go  
**this is plugin dependent 

# Mounter binding matrix for dynamically provisioned volumes:

Note: The value used for the plugin to indicate is it provisioning 
block will be plugin dependent and is an opaque parameter. Thus, not an API
change and possibly inconsistent between plugins. We are suggesting using 'block'
to simplify validation in the code (rather than raw that what proposed before).

| PV volumeType | Plugin fstype | PVC volumeType  | Result           |
| --------------|:-------------:| ---------------:|-----------------:|
|  --           | ext4/xfs      | block           | NO BIND          |
|  --           | ext4/xfs      | unspecified     | BIND             |
|  --           | block         | block           | BIND             |

* unspecified defaults to file today for backwards compatibility.
