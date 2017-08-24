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
* Provide a means for full integration into the scheduler based on non-storage related requests (CPU, etc.)
* Provide a means of ensuring specific topology to ensure co-location of the data 
   
# Value add to Kubernetes
  
  By extending the API for volumes to specifically request a raw block device, we provide an explicit method for volume comsumption,
  whereas previously any request for storage was always fullfilled with a formatted fileystem, even when the underlying storage was 
  block. In addition, the ability to use a raw block device without a filesystem will allow
  Kubernetes better support of high performance applications that can utilize raw block devices directly for their storage. 
  Block volumes are critical to applications like databases (MongoDB, Cassandra) that require consistent I/O performance
  and low latency. For mission critical applications, like SAP, block storage is a requirement. 
  
  For applications that use block storage natively (like MongoDB) no additional configuration is required as the mount path passed
  to the application provides the device which MongoDB then uses for the storage path in the configuration file (dbpath). Specific
  tuning for each application to achieve the highest possibly performance is provided as part of its recommended configurations.
  
  Specific use cases around improved usage of storage consumption are included in the use cases listed below as follows:
  * An admin wishes to expose a block volume to be consumed as a block volume for the user  
  * A user wishes to utilitze block storage to fully realize the performance of an application tuned to using block devices
  * A user wishes to utilize raw block devices for consumption from a virtual machine
  * A user wishes to specify an inline volume as a block device in their pod
  Future use cases include dynamically provisioning and intelligent discovery of existing devices, which this proposal sets the 
  foundation for more fully developing these methods. 
  
  It is importnant to note that when a PV is bound, it is either bound as a raw block device or formatted with a filesystem. Therefore, 
  the PVC drives the request and intended usage of the device by specifying the volumeMode as part of the API. This design lends itself
  to future support of dynamic provisioning by also letting the request intiate from the PVC defining the role for the PV. It also allows
  flexibility in the implementation and storage plugins to determine their support of this feature.
 
# Design Overview

  The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The consumption and 
  definitions for the block devices will be driven through the PVC and PV definitions. Along with Storage
  Resource definitions, this will provide the admin with a consistent way of managing all storage. 
  The API changes proposed in the following section are minimal with the idea of defining a volumeMode to indicate both the definition
  and consumption of the devices. Since it's possible to create a volume as a block device and then later consume it by provisioning
  a filesystem on top, the design requires explicit intent for how the volume will be used.
  The additional benefit of explicitly defining how the volume is to be consumed will provide a means for indicating the method
  by which the device should be scrubbed when the claim is deleted, as this method will differ from a raw block device compared to a 
  filesystem. The ownership and responsibility of defining the rention policy shall be up to the plugin method being utilized and is not 
  covered in this proposal.
  
  The last design point is block devices should be able to be fully restricted by the admin in accordance with how inline volumes 
  are today. Ideally, the admin would want to be able to restrict either raw-local devices and or raw-network attached devices.
  
  To ensure backwards compatibility and a phased transition of this feature, the consensus from the community is to intentionally disable
  the volumeMode: Block for external provisioners until a suitable implementation for provisioner versioning has been accepted and 
  implemented in the community. This requirement is better described in the design PR discussion and will be implemented as a seperate
  initiative. Acceptable values for volumeMode are 'Block' and 'Filesystem'. Where 'Filesystem' is the default value today and not 
  required to be set in the PV/PVC.
  
# Proposed API Changes
   
## Persistent Volume Claim API Changes:
In the simplest case of static provisioning, a user asks for a volumeMode of block. The binder will only bind to a PV defined 
with the same volumeMode.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
name: myclaim
spec:
  volumeMode: Block #proposed API change
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi 
```

For dynamic provisioning and the use of the storageClass, the admin also specifically defines the intent of the volume by 
indicating the volumeMode as block. The provisioner for this class will validate whether or not it supports block and return
an error if it does not.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
name: myclaim
spec:
  storageClassName: local-fast 
  volumeMode: Block #proposed API change
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 80Gi 
```

## Persistent Volume API Changes:
For static provisioning the admin creates the volume and also is intentional about how the volume should be consumed. For backwards
compatibility, the absence of volumeMode will default to file which is how volumes work today, which are formatted with a filesystem depending on the plug-in chosen. Recycling will not be a supported reclaim policy as it has been deprecated. The path value in the local PV definition would be overloaded to define the path of the raw block device rather than the fileystem path.
```
kind: PersistentVolume
apiVersion: v1
metadata:
name: local-raw-pv
spec:
  volumeMode: Block #proposed API change
  capacity:
    storage: 100Gi
  local:
    path: /dev/xvdc #device path
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete 
```
## Pod API Changes:
To provide better specificy and ensure support of inline volumes, the following changes are proposed in the pod specification.
```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers
    - name: mysql
      image: mysql
      volumeDevices:
      - name: my-db-data
        devicePath: /dev/xvda
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: raw-pvc
```
## Storage Class non-API Changes:
For dynamic provisioning, it is assumed that values pass in the parameter section are opaque, thus the introduction of utilizing
fsType in the StorageClass can be used by the provisioner to indicate how to create the volume. The proposal for this value is
defined here:
https://github.com/kubernetes/kubernetes/pull/45345 
Therefore, a provisioner could potentially provision a block device and install the filesystem onto it by indicating the volumeMode
as 'block' but the fsType as 'xfs'.
This section is provided as a general guideline, but each provisioner may implement their parameters independent of what is defined
here. It is our recommendation that the volumeMode be the guidance for the provisioner and overrides the value given in the fstype. Therefore a provisioner should be able to ignore the fstype and provision a block device if that is what the user requested via the PVC and the provisioner can support this.

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: block-volume
provisioner: kubernetes.io/local-block-glusterfs
parameters:
  volumeMode: Block #opaque value  / plug-in dependent -AND/OR-
  fsType: block
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

# Pod Security Policy (PSP) Changes:
Since the utilization of block devices can pose a risk to possible kernel manipulation by malicious users, it may be desirable for the administrator to restrict the usage entirely within a cluster. A similar convention is used with hostPath in that for some kubernetes 
implmentations it is disabled by default.
Thus, the PSP can define whether a validating pod can request the usage of such devices through the volumeMode label in the PV/PVC.
The changes to the pod security policy for volumes would include a 'localRawBlockVolumes' and 'newtworkRawBlockVolumes' parameters as such:

```
NAME               PRIV      CAPS      SELINUX     RUNASUSER          FSGROUP     SUPGROUP    PRIORITY   READONLYROOTFS   VOLUMES
anyuid             false     []        MustRunAs   RunAsAny           RunAsAny    RunAsAny    10         false            [configMap downwardAPI emptyDir persistentVolumeClaim secret localRawBlockVolumes networkRawBlockVolumes]
```

# Use Cases

## UC1: 

DESCRIPTION: An admin wishes to pre-create a series of local raw block devices to expose as PVs for consumption. The admin wishes to specify the purpose of these devices by specifying 'block' as the volumeMode for the PVs.

WORKFLOW:

ADMIN:

```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: local-raw-pv
spec:
  volumeMode: Block
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
  volumeMode: Block
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

* User creates a persistent volume claim with volumeMode: Block option to bind pre-created iSCSI PV.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: raw-pvc
spec:
  volumeMode: Block
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
      volumeDevices:
      - name: my-db-data
        devicePath: /dev/xvda
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
  volumeMode: Block
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
      volumeDevices:
      - name: my-db-data
        devicePath: /var/lib/mysql/data
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: local-raw-pvc
```

## UC4: 

DESCRIPTION: StorageClass with non-dynamically created volumes

BACKGROUND: The admin wishes to create a storage class that will identify pre-provisioned block PVs based on a user's PVC request for volumeMode: Block. 

WORKFLOW: 

ADMIN:

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: block-volume
provisioner: no-provisioning 
parameters:
```
* Sample of pre-created volume definition:

```
apiVersion: v1
kind: PersistentVolume
metadata:
 name: pv-block-volume
spec:
 volumeMode: Block
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

BACKGROUND: The admin wishes to create a storage class that will dynamically create block PVs based on a user's PVC request for volumeMode: Block. The admin desires the volumes be created dynamically and deleted when the PV definition is deleted. 

WORKFLOW:

ADMIN:

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: local-fast
provisioner: kubernetes.io/local-block-ssd
parameters:
  volumeMode: Block #suggested value - this is plugin/provisioner dependent
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
 volumeMode: Block
 storageClassName: local-fast
 accessModes:
  - ReadWriteOnce
 resources:
   requests:
     storage: 10Gi
```  

## UC7: 

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
  volumeMode: Block
  capacity:
    storage: "10Gi"
  accessModes:
    - "ReadWriteOnce"
  gcePersistentDisk:
    fsType: "block"
    pdName: "gce-disk-1"
```

***If admin specifies volumeMode: Block + fstype: ext4 then they would have the default behavior of files on block ***
***fsType values will be provisioner dependent. Block is suggested for development simplicity. Since the PVC object is passed
   to the provisioner, it will be responsible for validating and handling whether or not it supports the volumeMode being passed ***

## UC8: 

DESCRIPTION: 

A developer wishes to enable their application to use a raw block device as an inline volume in the pod. 

WORKFLOW:

USER:

```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers:
    - name: mysql
      image: mysql
      volumeDevices:
      - name: my-db-data
        devicePath: /var/lib/mysql/data
    volumes:
    - name: my-db-data
      local:
        path: /dev/sdb
```
# Container Runtime considerations
It is important the values that are passed to the container runtimes are valid and support the current implementation of these various runtimes. Listed below are a table of various runtime and the mapping of their values to what is passed from the kubelet.

| runtime engine    | runtime options  | accessMode       |  
| --------------    |:----------------:| ----------------:|
| docker/runc/rkt   |  mknod / RWM     | RWO              |
| docker/runc/rkt   |       R          | ROX              |

The accessModes would be passed as part of the options array and would need validate against the specific runtime engine. 
Since rkt doesn't use the CRI, the config values would need to be passed in the legacy method.
Note: the container runtime doesn't require a priviledged pod to enable the device as RWX (RMW).

The runtime option would be placed in the DeviceInfo as such:
accessMode == RWX or RWO would map to:
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "rmw"}) 
accessMode == ROX would map to:
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "r"}) 

The implemenation plan would be to rename the current makeDevices to makeGPUDevices and create a seperate function to add the raw block devices to the option array to be passed to the container runtime. This would interate on the paths passed in for the pod/container.

Since the future of this in Kubernetes for GPUs and other plugable devices is migrating to a device plugin architecture, there are 
still differentiating components of storage that are enough to not to enforce alignment to their convention. Two factors when
considering the usage of device plugins center around discoverability and topology of devices. Since neither of these are requirements
for using raw block devices, the legacy method of populating the devices and appending it to the device array is sufficient.

# Implementation Plan, Features & Milesones

Phase 1: v1.8
Feature: Pre-provisioned PVs to precreated devices 

               Milestone 1: API changes

               Milestone 2: Restricted Access 
              
               Milestone 3: Changes to the mounter interface as today it is assumed 'file' as the default.
               
               Milestone 4: Expose volumeMode to users via kubectl
               
               Milestone 5: Adds enable/disable configuration to securityContext in PSP (Pod Security Policy) similar to hostPath
               
               Milestone 6: Validate container runtime options with user specifcations as indicated in UC3
	       
	       Milestone 7: Container Runtime changes

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
# Volume binding matrix for statically provisioned volumes:

| PV volumeMode | PVC volumeMode  | Result           |
| --------------|:---------------:| ----------------:|
|   unspecified | unspecified     | BIND             |
|   unspecified | block           | NO BIND          |
|   unspecified | file            | BIND             |
|   block       | unspecified     | NO BIND          |
|   block       |  block          | BIND             |
|   block       |  file           | NO BIND          |
|   file        |  file           | BIND             |
|   file        |  block          | NO BIND          |
|   file        | unspecified     | BIND             |



* unspecified defaults to 'file/ext4' today for backwards compatibility and in mount_linux.go  


# Volume binding matrix for dynamically provisioned volumes:

Note: The value used for the plugin to indicate is it provisioning 
block will be plugin dependent and is an opaque parameter. Binding will also be plugin dependent and must handle the parameter being passed and indicate whether or not it supports block. 
