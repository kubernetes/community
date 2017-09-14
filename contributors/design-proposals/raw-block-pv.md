# Raw Block Consumption in Kubernetes

Authors: erinboyd@, screeley44@, mtanino@

This document presents a proposal for managing raw block storage in Kubernetes using the persistent volume source API as a consistent model of consumption.

# Terminology
* Raw Block Device - a physically attached device devoid of a filesystem
* Raw Block Volume - a logical abstraction of the raw block device as defined by a path
* Filesystem on Block - a formatted (ie xfs) filesystem on top of a raw block device

# Goals
* Enable durable access to block storage
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
  
  By extending the API for volumes to specifically request a raw block device, we provide an explicit method for volume consumption,
  whereas previously any request for storage was always fulfilled with a formatted fileystem, even when the underlying storage was 
  block. In addition, the ability to use a raw block device without a filesystem will allow
  Kubernetes better support of high performance applications that can utilize raw block devices directly for their storage. 
  Block volumes are critical to applications like databases (MongoDB, Cassandra) that require consistent I/O performance
  and low latency. For mission critical applications, like SAP, block storage is a requirement. 
  
  For applications that use block storage natively (like MongoDB) no additional configuration is required as the mount path passed
  to the application provides the device which MongoDB then uses for the storage path in the configuration file (dbpath). Specific
  tuning for each application to achieve the highest possibly performance is provided as part of its recommended configurations.
  
  Specific use cases around improved usage of storage consumption are included in the use cases listed below as follows:
  * An admin wishes to expose a block volume to be consumed as a block volume for the user  
  * An admin wishes to expose a block volume to be consumed as a block volume for an administrative function such 
    as bootstrapping 
  * A user wishes to utilize block storage to fully realize the performance of an application tuned to using block devices
  * A user wishes to read from a block storage device and write to a filesystem (big data analytics processing)
  Future use cases include dynamically provisioning and intelligent discovery of existing devices, which this proposal sets the 
  foundation for more fully developing these methods. 

 
# Design Overview

  The proposed design is based on the idea of leveraging well defined concepts for storage in Kubernetes. The consumption and 
  definitions for the block devices will be driven through the PVC and PV definitions. Along with Storage
  Resource definitions, this will provide the admin with a consistent way of managing all storage. 
  The API changes proposed in the following section are minimal with the idea of defining a volumeMode to indicate both the definition
  and consumption of the devices. Since it's possible to create a volume as a block device and then later consume it by provisioning
  a filesystem on top, the design requires explicit intent for how the volume will be used.
  The additional benefit of explicitly defining how the volume is to be consumed will provide a means for indicating the method
  by which the device should be scrubbed when the claim is deleted, as this method will differ from a raw block device compared to a 
  filesystem. The ownership and responsibility of defining the retention policy shall be up to the plugin method being utilized and is
  not covered in this proposal.
  
  Limiting use of the volumeMode to block can be executed through the use of storage resource quotas and storageClasses defined by the 
  administrator.
  
  To ensure backwards compatibility and a phased transition of this feature, the consensus from the community is to intentionally disable
  the volumeMode: Block for external provisioners until a suitable implementation for provisioner versioning has been accepted and 
  implemented in the community. In addition, in-tree provisioners should be able to gracefully ignore volumeMode API objects for plugins
  that haven't been updated to accept this value.
  
  It is important to note that when a PV is bound, it is either bound as a raw block device or formatted with a filesystem. Therefore, 
  the PVC drives the request and intended usage of the device by specifying the volumeMode as part of the API. This design lends itself
  to future support of dynamic provisioning by also letting the request initiate from the PVC defining the role for the PV. It also 
  allows flexibility in the implementation and storage plugins to determine their support of this feature. Acceptable values for 
  volumeMode are 'Block' and 'Filesystem'. Where 'Filesystem' is the default value today and not required to be set in the PV/PVC.
  
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
compatibility, the absence of volumeMode will default to filesystem which is how volumes work today, which are formatted with a filesystem depending on the plug-in chosen. Recycling will not be a supported reclaim policy as it has been deprecated. The path value in the local PV definition would be overloaded to define the path of the raw block device rather than the fileystem path.
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
This change intentionally calls out the use of a block device (volumeDevices) rather than the mount point on a filesystem.
```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers
    - name: mysql
      image: mysql
      volumeDevices: #proposed API change
      - name: my-db-data
        devicePath: /dev/xvda #proposed API change
    volumes:
    - name: my-db-data
      persistentVolumeClaim:
	claimName: raw-pvc
```
## Storage Class non-API Changes:
For dynamic provisioning, it is assumed that values passed in the parameter section are opaque, thus the introduction of utilizing
fstype in the StorageClass can be used by the provisioner to indicate how to create the volume. The proposal for this value is
defined here:
https://github.com/kubernetes/kubernetes/pull/45345 
This section is provided as a general guideline, but each provisioner may implement their parameters independent of what is defined
here. It is our recommendation that the volumeMode in the PVC be the guidance for the provisioner and overrides the value given in the fstype. Therefore a provisioner should be able to ignore the fstype and provision a block device if that is what the user requested via the PVC and the provisioner can support this.

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: block-volume
 provisioner: kubernetes.io/scaleio
 parameters:
  gateway: https://192.168.99.200:443/api
  system: scaleio
  protectionDomain: default
  storagePool: default
  storageMode: ThinProvisionned
  secretRef: sio-secret
  readOnly: false
  fsType: Block #suggested value
```
The provisioner (if applicable) should validate the parameters and return an error if the combination specified is not supported.
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
* A user uses a raw block device for database applications such as MariaDB.
* User creates a persistent volume claim with "volumeMode: Block" option to bind pre-created iSCSI PV. 

WORKFLOW:

ADMIN:
* Admin creates a disk and exposes it to all kubelet worker nodes. (This is done by storage operation).
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
* During Pod creation, iSCSI Plugin attaches iSCSI volume to the kubelet worker node using storage information.


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
    pdName: "gce-disk-1"
```
***Since the PVC object is passed to the provisioner, it will be responsible for validating and handling whether or not it supports the volumeMode being passed ***

## UC8:

DESCRIPTION: 
* A user uses a raw block device for database applications such as mysql to read data from and write the results to a disk that 
  has a formatted filesystem to be displayed via nginx web server.

ADMIN:
* Admin creates a 2 block devices and formats one with a filesystem

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
  gcePersistentDisk:
    pdName: "gce-disk-1"
  
```
```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: gluster-pv
spec:
  volumeMode: Filesystem
  capacity:
    storage: 100Gi
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Delete
  glusterfs: 
    endpoints: glusterfs-cluster 
    path: glusterVol
```
USER:

* User creates a persistent volume claim with volumeMode: Block option to bind pre-created block volume.

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
* User creates a persistent volume claim with volumeMode: Filesystem to the pre-created gluster volume.

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: gluster-pvc
spec:
  volumeMode: Filesystem
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 50Gi
```
* User creates a Pod yaml which will utilitze both block and filesystem storage by its containers.

```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
  volumes:
  - name: my-db-data
    persistentVolumeClaim:
      claimName: raw-pvc
  - name: my-nginx-data
    persistentVolumeClaim:
      claimName: gluster-pvc
  containers
    - name: mysql
      image: mysql
      volumeDevices: 
      - name: my-db-data
        devicePath: /var/lib/mysql/data
    - name: nginx
      image: nginx
      ports:
      - containerPort: 80
      volumeMounts:
      - mountPath: /usr/share/nginx/html
        name: my-nginx-data 
	readOnly: false
```

# Container Runtime considerations
It is important the values that are passed to the container runtimes are valid and support the current implementation of these various runtimes. Listed below are a table of various runtime and the mapping of their values to what is passed from the kubelet.

| runtime engine    | runtime options  | accessMode       |  
| --------------    |:----------------:| ----------------:|
| docker/runc/rkt   |  mknod / RWM     | RWO              |
| docker/runc/rkt   |       R          | ROX              |

The accessModes would be passed as part of the options array and would need validate against the specific runtime engine. 
Since rkt doesn't use the CRI, the config values would need to be passed in the legacy method.
Note: the container runtime doesn't require a privileged pod to enable the device as RWX (RMW), but still requires privileges to mount as is consistent with the filesystem implemenatation today.

The runtime option would be placed in the DeviceInfo as such:
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "XXX"}) 

The implemenation plan would be to rename the current makeDevices to makeGPUDevices and create a separate function to add the raw block devices to the option array to be passed to the container runtime. This would iterate on the paths passed in for the pod/container.

Since the future of this in Kubernetes for GPUs and other plug-able devices is migrating to a device plugin architecture, there are 
still differentiating components of storage that are enough to not to enforce alignment to their convention. Two factors when
considering the usage of device plugins center around discoverability and topology of devices. Since neither of these are requirements
for using raw block devices, the legacy method of populating the devices and appending it to the device array is sufficient.

# Implementation Plan, Features & Milestones

Phase 1: v1.8
Feature: Pre-provisioned PVs to precreated devices 

               Milestone 1: API changes

               Milestone 2: Restricted Access 
              
               Milestone 3: Changes to the mounter interface as today it is assumed 'file' as the default.
               
               Milestone 4: Expose volumeMode to users via kubectl
               
               Milestone 5: PV controller binding changes for block devices
               	       
	       Milestone 6: Container Runtime changes
	       
	       Milestone 7: Initial Plugin changes (GCE, AWS & GlusterFS)
	       
	       Milestone 8: Disabling of provisioning where volumeMode == Block is not supported

Phase 2:  v1.9
Feature: Discovery of block devices 

                Milestone 1: Dynamically provisioned PVs to dynamically allocated devices

                Milestone 2: Privileged container concerns    
		
		Milestone 3: Flex volume update
		
Other considerations:

               Reference volume driver change(Attach/Detach logic) for pre-provisioned PVs

               Reference volume driver change(Provision/Delete logic) for dynamic provisioning

## Mounter interface proposed design
# Plugin changes
## New BlockVolumeMapper interface proposed design

```
 type BlockVolumeMapper interface {
       Volume
       CanBlockMap() error
       SetUpDevice(podUID types.UID) error
       SetUpDeviceAt(dir string, podUID types.UID) error
       GetDevicePath() (string, error)
       GetVolumeDeviceMapPath(spec *Spec) (string, error)
 }
 type BlockVolumeUnmapper interface {
       Volume
       TearDownDevice() error
       TearDownDeviceAt(dir string) error
       GetVolumeDeviceUnmapPath(spec *Spec) (string, error)
 }
```
## Changes for volume mount points

Currently, a volume which has filesystem is mounted to the following two paths on a kubelet node when the volumes is in-use.
The purpose of those mount points are that Kubernetes manages volume attach/detach status using these mount points and number
of references to these mount points.

```
- Global mount path
/var/lib/kubelet/plugins/kubernetes.io/{pluginName}/{volumePluginDependentPath}/

- Volume mount path
/var/lib/kubelet/pods/{podUID}/volumes/{escapeQualifiedPluginName}/{volumeName}/
```

Even if the volumeMode is "Block", similar scheme is needed. However, the volume which 
doesn't have filesystem can't be mounted.
Therefore, instead of volume mount, we use symbolic link to map raw block device.
Kubelet creates a new symbolic link under the new global map path when volume is attached to a Pod. Number of symbolic links
are equal to the number of Pods which uses the same volume. Kubelet needs to manage both creation and deletion of symbolic links
under the global map path. The name of the symbolic link is same as pod uuid.
 
```
Global map path for "Block" volumeMode volume
/var/lib/kubelet/plugins/kubernetes.io/{pluginName}/volumeDevices/{volumePluginDependentPath}/pod-uuid1
/var/lib/kubelet/plugins/kubernetes.io/{pluginName}/volumeDevices/{volumePluginDependentPath}/pod-uuid2
```
 
Plugin creates a symbolic link under the new volume map path. This symbolic link is not used to manage volume attach/detach
status but is needed to keep compatibility of current scheme.

```
Volume map path for "Block" volumeMode volume
/var/lib/kubelet/pods/{podUID}/volumeDevices/{escapeQualifiedPluginName}/{volumeName}/symlink
```
 
# Volume binding matrix for statically provisioned volumes:

| PV volumeMode | PVC volumeMode  | Result           |
| --------------|:---------------:| ----------------:|
|   unspecified | unspecified     | BIND             |
|   unspecified | Block           | NO BIND          |
|   unspecified | Filesystem      | BIND             |
|   Block       | unspecified     | NO BIND          |
|   Block       | Block           | BIND             |
|   Block       | Filesystem      | NO BIND          |
|   Filesystem  | Filesystem      | BIND             |
|   Filesystem  | Block           | NO BIND          |
|   Filesystem  | unspecified     | BIND             |



* unspecified defaults to 'file/ext4' today for backwards compatibility and in mount_linux.go  


# Volume binding considerations for dynamically provisioned volumes:
The value used for the plugin to indicate is it provisioning block will be plugin dependent and is an opaque parameter. Binding will also be plugin dependent and must handle the parameter being passed and indicate whether or not it supports block. 
