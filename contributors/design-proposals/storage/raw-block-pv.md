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
  whereas previously any request for storage was always fulfilled with a formatted filesystem, even when the underlying storage was 
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
  Also use cases include dynamically provisioning and intelligent discovery of existing devices, which this proposal sets the
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
  the volumeMode: Block for both in-tree and external provisioners until a suitable implementation for provisioner versioning has been
  accepted and implemented in the community. In addition, in-tree provisioners should be able to gracefully ignore volumeMode API objects
  for plugins that haven't been updated to accept this value.
  
  It is important to note that when a PV is bound, it is either bound as a raw block device or formatted with a filesystem. Therefore, 
  the PVC drives the request and intended usage of the device by specifying the volumeMode as part of the API. This design lends itself
  to support of dynamic provisioning by also letting the request initiate from the PVC defining the role for the PV. It also
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
compatibility, the absence of volumeMode will default to filesystem which is how volumes work today, which are formatted with a filesystem depending on the plug-in chosen. Recycling will not be a supported reclaim policy as it has been deprecated. The path value in the local PV definition would be overloaded to define the path of the raw block device rather than the filesystem path.
```
kind: PersistentVolume
apiVersion: v1
metadata:
  name: local-raw-pv
  annotations:
        "volume.alpha.kubernetes.io/node-affinity": '{
            "requiredDuringSchedulingIgnoredDuringExecution": {
                "nodeSelectorTerms": [
                    { "matchExpressions": [
                        { "key": "kubernetes.io/hostname",
                          "operator": "In",
                          "values": ["ip-172-18-11-174.ec2.internal"]
                        }
                    ]}
                 ]}
              }'
spec:
  volumeMode: Block
  capacity:
    storage: 10Gi
  local:
    path: /dev/xvdf
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain  
```
## Pod API Changes:
This change intentionally calls out the use of a block device (volumeDevices) rather than the mount point on a filesystem.
```
apiVersion: v1
kind: Pod
metadata:
  name: my-db
spec:
    containers:
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
    - namee: mysql
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
## UC5:

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

**Since the PVC object is passed to the provisioner, it will be responsible for validating and handling whether or not it supports the volumeMode being passed**

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
Spec:
  volumeMode: Block
  capacity:
    storage: "10Gi"
  accessModes:
    - "ReadWriteOnce"
  gcePersistentDisk:
    pdName: "gce-disk-1"
```

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
* User creates a Pod yaml which will utilize both block and filesystem storage by its containers.

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
## UC9:

DESCRIPTION:
* A user wishes to read data from a read-only raw block device, an example might be a database for analytics processing. 

USER:
* User creates pod and specifies 'readOnly' as a parameter in the persistent volume claim to indicate they would
like to be bound to a PV with this setting enabled.

```
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod-block-001
spec:
  containers:
    - name: nginx-container
      image: nginx:latest
      ports:
      - containerPort: 80
      volumeDevices:
        - name: data
          devicePath: /dev/xvda
  volumes:
    - name: data
      persistentVolumeClaim:
        claimName: block-pvc001
        readOnly: true #flag indicating read-only for container runtime
```
**Note: the readOnly field already exists in the PersistentVolumeClaimVolumeSource above and will dictate the values set by the container runtime options**


# Container Runtime considerations
It is important the values that are passed to the container runtimes are valid and support the current implementation of these various runtimes. Listed below are a table of various runtime and the mapping of their values to what is passed from the kubelet.

| runtime engine    | runtime options  | accessMode       |  
| --------------    |:----------------:| ----------------:|
| docker/runc/rkt   |  mknod / RWM     | RWO              |
| docker/runc/rkt   |       R          | ROX              |

The accessModes would be passed as part of the options array and would need validate against the specific runtime engine. 
Since rkt doesn't use the CRI, the config values would need to be passed in the legacy method.
Note: the container runtime doesn't require a privileged pod to enable the device as RWX (RMW), but still requires privileges to mount as is consistent with the filesystem implementation today.

The runtime option would be placed in the DeviceInfo as such:
devices = append(devices, kubecontainer.DeviceInfo{PathOnHost: path, PathInContainer: path, Permissions: "XXX"}) 

The implementation plan would be to rename the current makeDevices to makeGPUDevices and create a separate function to add the raw block devices to the option array to be passed to the container runtime. This would iterate on the paths passed in for the pod/container.

Since the future of this in Kubernetes for GPUs and other plug-able devices is migrating to a device plugin architecture, there are 
still differentiating components of storage that are enough to not to enforce alignment to their convention. Two factors when
considering the usage of device plugins center around discoverability and topology of devices. Since neither of these are requirements
for using raw block devices, the legacy method of populating the devices and appending it to the device array is sufficient.


# Plugin interface changes
## New BlockVolume interface proposed design

```
// BlockVolume interface provides methods to generate global map path
// and pod device map path.
type BlockVolume interface {
	// GetGlobalMapPath returns a global map path which contains
	// symbolic links associated to a block device.
	// ex. plugins/kubernetes.io/{PluginName}/{DefaultKubeletVolumeDevicesDirName}/{volumePluginDependentPath}/{pod uuid}
	GetGlobalMapPath(spec *Spec) (string, error)
	// GetPodDeviceMapPath returns a pod device map path
	// and name of a symbolic link associated to a block device.
	// ex. pods/{podUid}}/{DefaultKubeletVolumeDevicesDirName}/{escapeQualifiedPluginName}/{volumeName}
	GetPodDeviceMapPath() (string, string)
}
```

## New BlockVolumePlugin interface proposed design

```
// BlockVolumePlugin is an extend interface of VolumePlugin and is used for block volumes support.
type BlockVolumePlugin interface {
	VolumePlugin
	// NewBlockVolumeMapper creates a new volume.BlockVolumeMapper from an API specification.
	// - spec: The v1.Volume spec
	// - pod: The enclosing pod
	NewBlockVolumeMapper(spec *Spec, podRef *v1.Pod, opts VolumeOptions) (BlockVolumeMapper, error)
	// NewBlockVolumeUnmapper creates a new volume.BlockVolumeUnmapper from recoverable state.
	// - name: The volume name, as per the v1.Volume spec.
	// - podUID: The UID of the enclosing pod
	NewBlockVolumeUnmapper(name string, podUID types.UID) (BlockVolumeUnmapper, error)
	// ConstructBlockVolumeSpec constructs a volume spec based on the given
	// pod name, volume name and a pod device map path.
	// The spec may have incomplete information due to limited information
	// from input. This function is used by volume manager to reconstruct
	// volume spec by reading the volume directories from disk.
	ConstructBlockVolumeSpec(podUID types.UID, volumeName, mountPath string) (*Spec, error)
}
```

## New BlockVolumeMapper/BlockVolumeUnmapper interface proposed design

```
// BlockVolumeMapper interface provides methods to set up/map the volume.
type BlockVolumeMapper interface {
	BlockVolume
	// SetUpDevice prepares the volume to a self-determined directory path,
	// which may or may not exist yet and returns combination of physical
	// device path of a block volume and error.
	// If the plugin is non-attachable, it should prepare the device
	// in /dev/ (or where appropriate) and return unique device path.
	// Unique device path across kubelet node reboot is required to avoid
	// unexpected block volume destruction.
	// If the plugin is attachable, it should not do anything here,
	// just return empty string for device path.
	// Instead, attachable plugin have to return unique device path
	// at attacher.Attach() and attacher.WaitForAttach().
	// This may be called more than once, so implementations must be idempotent.
	SetUpDevice() (string, error)
}

// BlockVolumeUnmapper interface provides methods to cleanup/unmap the volumes.
type BlockVolumeUnmapper interface {
	BlockVolume
	// TearDownDevice removes traces of the SetUpDevice procedure under
	// a self-determined directory.
	// If the plugin is non-attachable, this method detaches the volume
	// from devicePath on kubelet node.
	TearDownDevice(mapPath string, devicePath string) error
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
Therefore, instead of volume mount, we use symbolic link which is associated to raw block device.
Kubelet creates a new symbolic link under the new `global map path` and `pod device map path`.

#### Global map path for "Block" volumeMode volume
Kubelet creates a new symbolic link under the new global map path when volume is attached to a Pod.
Number of symbolic links are equal to the number of Pods which use the same volume. Kubelet needs
to manage both creation and deletion of symbolic links under the global map path. The name of the
symbolic link is same as pod uuid.
There are two usages of Global map path.

1. Manage number of references from multiple pods
1. Retrieve `{volumePluginDependentPath}` during `Block volume reconstruction`

```
/var/lib/kubelet/plugins/kubernetes.io/{pluginName}/volumeDevices/{volumePluginDependentPath}/{pod uuid1}
/var/lib/kubelet/plugins/kubernetes.io/{pluginName}/volumeDevices/{volumePluginDependentPath}/{pod uuid2}
...
```

- {volumePluginDependentPath} example:
```
FC plugin: {wwn}-lun-{lun} or {wwid}
ex. /var/lib/kubelet/plugins/kubernetes.io/fc/volumeDevices/500a0982991b8dc5-lun-0/f527ca5b-6d87-11e5-aa7e-080027ff6387
iSCSI plugin: {portal ip}-{iqn}-lun-{lun}
ex. /var/lib/kubelet/plugins/kubernetes.io/iscsi/volumeDevices/1.2.3.4:3260-iqn.2001-04.com.example:storage.kube.sys1.xyz-lun-1/f527ca5b-6d87-11e5-aa7e-080027ff6387
 ```
 
#### Pod device map path for "Block" volumeMode volume
Kubelet creates a symbolic link under the new pod device map path. The file of {volumeName} is
symbolic link and the link is associated to raw block device. If a Pod has multiple block volumes,
multiple symbolic links under the pod device map path will be created with each volume name.
The usage of pod device map path is;

1. Retrieve raw block device path(ex. /dev/sdX) during `Container initialization` and `Block volume reconstruction`

```
/var/lib/kubelet/pods/{podUID}/volumeDevices/{escapeQualifiedPluginName}/{volumeName1}
/var/lib/kubelet/pods/{podUID}/volumeDevices/{escapeQualifiedPluginName}/{volumeName2}
...
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

# Dynamically provisioning

Using dynamic provisioning, user is able to create block volume via provisioners. Currently,
we have two types of provisioners, internal provisioner and external provisioner.
During volume creation via dynamic provisioner, user passes persistent volume claim which
contains `volumeMode` parameter, then the persistent volume claim object is passed to
provisioners. Therefore, in order to create block volume, provisioners need to support
`volumeMode` and then create persistent volume with `volumeMode`.

If a storage and plugin don't have an ability to create raw block type of volume,
then `both internal and external provisioner don't need any update` to support `volumeMode`
because `volumeMode` in PV and PVC are automatically set to `Filesystem` as a default when
these volume object are created.
However, there is a case that use specifies `volumeMode` as `Block` even if both plugin and
provisioner don't support. As a result, PVC will be created, PV will be provisioned
but both of them will stuck Pending status since `volumeMode` between them don't match.
For this situation, we will add error propagation into persistent volume controller to make
it more clear to the user what's wrong.

If admin provides external provisioner to provision both filesystem and block volume,
admin have to carefully prepare Kubernetes environment for their users because both
Kubernetes itself and external provisioner have to support block volume functionality.
This means Kubernetes v1.9 or later must be used to provide block volume with external
provisioner which supports block volume.

Regardless of the volumeMode, provisioner can set `FSType` into the plugin's volumeSource
but the value will be ignored at the volume plugin side if `volumeMode` is `Block`.

## Internal provisioner

If internal plugin has own provisioner, the plugin needs to support `volumeMode` to provision
block volume. This is the example implementation of `volumeMode` support for GCE PD plugin.

```
// Obtain volumeMode from PVC Spec VolumeMode
var volumeMode v1.PersistentVolumeMode
if options.PVC.Spec.VolumeMode != nil {
  volumeMode = *options.PVC.Spec.VolumeMode
}

// Set volumeMode into PersistentVolumeSpec
pv := &v1.PersistentVolume{
  Spec: v1.PersistentVolumeSpec{
    VolumeMode: &volumeMode,
    PersistentVolumeSource: v1.PersistentVolumeSource{
      GCEPersistentDisk: &v1.GCEPersistentDiskVolumeSource{
        PDName:      options.Parameters["pdName"],
        FSType:      options.Parameters["fsType"],
        ...
      },
    },
  },
}
```


## External provisioner

We have a "protocol" to allow dynamic provisioning by external software called external provisioner.
In order to support block volume via external provisioner, external provisioner needs to support
`volumeMode` and then create persistent volume with `volumeMode`. This is the example implementation
of `volumeMode` support for external provisioner of Local volume plugin.

```
// Obtain volumeMode from PVC Spec VolumeMode
var volumeMode v1.PersistentVolumeMode
if options.PVC.Spec.VolumeMode != nil {
  volumeMode = *options.PVC.Spec.VolumeMode
}

// Set volumeMode into PersistentVolumeSpec
pv := &v1.PersistentVolume{
  Spec: v1.PersistentVolumeSpec{
    VolumeMode: &volumeMode,
    PersistentVolumeSource: v1.PersistentVolumeSource{
      Local: &v1.LocalVolumeSource{
        Path:      options.Parameters["Path"],
      },
    },
  },
}
```


# Volume binding considerations for dynamically provisioned volumes:
The value used for the plugin to indicate is it provisioning block will be plugin dependent and is an opaque parameter. Binding will also be plugin dependent and must handle the parameter being passed and indicate whether or not it supports block. 

# Implementation Plan, Features & Milestones

Phase 1: v1.9
Feature: Pre-provisioned PVs to precreated devices

               Milestone 1: API changes

               Milestone 2: Restricted Access

               Milestone 3: Changes to the mounter interface as today it is assumed 'file' as the default.

               Milestone 4: Expose volumeMode to users via kubectl

               Milestone 5: PV controller binding changes for block devices

               Milestone 6: Container Runtime changes

               Milestone 7: Initial Plugin changes (FC & Local storage)

Phase 2:  v1.10
Feature: Discovery of block devices

                Milestone 1: Dynamically provisioned PVs to dynamically allocated devices

                Milestone 2: Plugin changes with dynamic provisioning support (RBD, iSCSI, GCE, AWS & GlusterFS)
