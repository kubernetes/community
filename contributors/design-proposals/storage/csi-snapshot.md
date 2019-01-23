Kubernetes CSI Snapshot Proposal
================================

**Authors:** [Jing Xu](https://github.com/jingxu97), [Xing Yang](https://github.com/xing-yang), [Tomas Smetana](https://github.com/tsmetana), [Huamin Chen ](https://github.com/rootfs)

## Background

Many storage systems (GCE PD, Amazon EBS, etc.) provide the ability to create "snapshots" of persistent volumes to protect against data loss. Snapshots can be used in place of a traditional backup system to back up and restore primary and critical data. Snapshots allow for quick data backup (for example, it takes a fraction of a second to create a GCE PD snapshot) and offer fast recovery time objectives (RTOs) and recovery point objectives (RPOs). Snapshots can also be used for data replication, distribution and migration.

As the initial effort to support snapshot in Kubernetes, volume snapshotting has been released as a prototype in Kubernetes 1.8. An external controller and provisioner (i.e. two separate binaries) have been added in the [external storage repo](https://github.com/kubernetes-incubator/external-storage/tree/master/snapshot). The prototype currently supports GCE PD, AWS EBS, OpenStack Cinder, GlusterFS, and Kubernetes hostPath volumes. Volume snapshots APIs are using [CRD](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/).

To continue that effort, this design is proposed to add the snapshot support for CSI Volume Drivers. Because the overall trend in Kubernetes is to keep the core APIs as small as possible and use CRD for everything else, this proposal adds CRD definitions to represent snapshots, and an external snapshot controller to handle volume snapshotting. Out-of-tree external provisioner can be upgraded to support creating volume from snapshot. In this design, only CSI volume drivers will be supported. The CSI snapshot spec is proposed [here](https://github.com/container-storage-interface/spec/pull/224).


## Objectives

For the first version of snapshotting support in Kubernetes, only on-demand snapshots for CSI Volume Drivers will be supported.


### Goals

* Goal 1: Expose standardized snapshotting operations to create, list, and delete snapshots in Kubernetes REST API.
Currently the APIs will be implemented with CRD (CustomResourceDefinitions).

* Goal 2: Implement CSI volume snapshot support.
An external snapshot controller will be deployed with other external components (e.g., external-attacher, external-provisioner) for each CSI Volume Driver.

* Goal 3: Provide a convenient way of creating new and restoring existing volumes from snapshots.


### Non-Goals

The following are non-goals for the current phase, but will be considered at a later phase.

* Goal 4: Offer application-consistent snapshots by providing pre/post snapshot hooks to freeze/unfreeze applications and/or unmount/mount file system.

* Goal 5: Provide higher-level management, such as backing up and restoring a pod and statefulSet, and creating a consistent group of snapshots.


## Design Details

In this proposal, volume snapshots are considered as another type of storage resources managed by Kubernetes. Therefore the snapshot API and controller follow the design pattern of existing volume management. There are three APIs, VolumeSnapshot and VolumeSnapshotContent, and VolumeSnapshotClass which are similar to the structure of PersistentVolumeClaim and PersistentVolume, and storageClass. The external snapshot controller functions similar to the in-tree PV controller. With the snapshots APIs, we also propose to add a new data source struct in PersistentVolumeClaim (PVC) API in order to support restore snapshots to volumes. The following section explains in more details about the APIs and the controller design.


### Snapshot API Design

The API design of VolumeSnapshot and VolumeSnapshotContent is modeled after PersistentVolumeClaim and PersistentVolume. In the first version, the VolumeSnapshot lifecycle is completely independent of its volumes source (PVC). When PVC/PV is deleted, the corresponding VolumeSnapshot and VolumeSnapshotContents objects will continue to exist. However, for some volume plugins, snapshots have a dependency on their volumes. In a future version, we plan to have a complete lifecycle management which can better handle the relationship between snapshots and their volumes. (e.g., a finalizer to prevent deleting volumes while there are snapshots depending on them).

#### The `VolumeSnapshot` Object

```GO

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshot is a user's request for taking a snapshot. Upon successful creation of the actual
// snapshot by the volume provider it is bound to the corresponding VolumeSnapshotContent. 
// Only the VolumeSnapshot object is accessible to the user in the namespace. 
type VolumeSnapshot struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the desired characteristics of a snapshot requested by a user.
	Spec VolumeSnapshotSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the latest observed state of the snapshot
	// +optional
	Status VolumeSnapshotStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotList is a list of VolumeSnapshot objects
type VolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of VolumeSnapshots
	Items []VolumeSnapshot `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// VolumeSnapshotSpec describes the common attributes of a volume snapshot
type VolumeSnapshotSpec struct {
	// Source has the information about where the snapshot is created from.
	// In Alpha version, only PersistentVolumeClaim is supported as the source.
	// If not specified, user can create VolumeSnapshotContent and bind it with VolumeSnapshot manually.
	// +optional
	Source *TypedLocalObjectReference `json:"source" protobuf:"bytes,1,opt,name=source"`

	// SnapshotContentName binds the VolumeSnapshot object with the VolumeSnapshotContent
	// +optional
	SnapshotContentName string `json:"snapshotContentName" protobuf:"bytes,2,opt,name=snapshotContentName"`

	// Name of the VolumeSnapshotClass used by the VolumeSnapshot. If not specified, a default snapshot class will
	// be used if it is available.
	// +optional
	VolumeSnapshotClassName *string `json:"snapshotClassName" protobuf:"bytes,3,opt,name=snapshotClassName"`
}

// VolumeSnapshotStatus is the status of the VolumeSnapshot
type VolumeSnapshotStatus struct {
	// CreationTime is the time the snapshot was successfully created. If it is set,
	// it means the snapshot was created; Otherwise the snapshot was not created.
	// +optional
	CreationTime *metav1.Time `json:"creationTime" protobuf:"bytes,1,opt,name=creationTime"`

	// When restoring volume from the snapshot, the volume size should be equal or 
	// larger than the Restoresize if it is specified. If RestoreSize is set to nil, it means
	// that the storage plugin does not have this information available.
	// +optional
	RestoreSize *resource.Quantity `json:"restoreSize" protobuf:"bytes,2,opt,name=restoreSize"`
	
	// Ready is set to true only if the snapshot is ready to use (e.g., finish uploading if
	// there is an uploading phase) and also VolumeSnapshot and its VolumeSnapshotContent
	// bind correctly with each other. If any of the above condition is not true, Ready is
	// set to false
	// +optional
	Ready bool `json:"ready" protobuf:"varint,3,opt,name=ready"`

	// The last error encountered during create snapshot operation, if any.
	// This field must only be set by the entity completing the create snapshot
	// operation, i.e. the external-snapshotter.
	// +optional
	Error *storage.VolumeError
}

```

Note that if an error occurs before the snapshot is cut, `Error` will be set and none of `CreatedAt`/`AvailableAt` will be set. If an error occurs after the snapshot is cut but before it is available, `Error` will be set and `CreatedAt` should still be set, but `AvailableAt` will not be set. If an error occurs after the snapshot is available, `Error` will be set and `CreatedAt` should still be set, but `AvailableAt` will no longer be set.

#### The `VolumeSnapshotContent` Object

```GO

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotContent represents the actual snapshot object
type VolumeSnapshotContent struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines a specification of a volume snapshot
	Spec VolumeSnapshotContentSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotContentList is a list of VolumeSnapshotContent objects
type VolumeSnapshotContentList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of VolumeSnapshotContents
	Items []VolumeSnapshotContent `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// VolumeSnapshotContentSpec is the spec of the volume snapshot content
type VolumeSnapshotContentSpec struct {
	// Source represents the location and type of the volume snapshot
	VolumeSnapshotSource `json:",inline" protobuf:"bytes,1,opt,name=volumeSnapshotSource"`

	// VolumeSnapshotRef is part of bi-directional binding between VolumeSnapshot
	// and VolumeSnapshotContent. It becomes non-nil when bound.
	// +optional
	VolumeSnapshotRef *core_v1.ObjectReference `json:"volumeSnapshotRef" protobuf:"bytes,2,opt,name=volumeSnapshotRef"`

	// PersistentVolumeRef represents the PersistentVolume that the snapshot has been
	// taken from. It becomes non-nil when VolumeSnapshot and VolumeSnapshotContent are bound.
	// +optional
	PersistentVolumeRef *core_v1.ObjectReference `json:"persistentVolumeRef" protobuf:"bytes,3,opt,name=persistentVolumeRef"`
	// Name of the VolumeSnapshotClass used by the VolumeSnapshotContent. If not specified, a default snapshot class will
	// be used if it is available.
	// +optional
	VolumeSnapshotClassName *string `json:"snapshotClassName" protobuf:"bytes,4,opt,name=snapshotClassName"`
}

// VolumeSnapshotSource represents the actual location and type of the snapshot. Only one of its members may be specified.
type VolumeSnapshotSource struct {
	// CSI (Container Storage Interface) represents storage that handled by an external CSI Volume Driver (Alpha feature).
	// +optional
	CSI *CSIVolumeSnapshotSource `json:"csiVolumeSnapshotSource,omitempty"`
}

// Represents the source from CSI volume snapshot
type CSIVolumeSnapshotSource struct {
	// Driver is the name of the driver to use for this snapshot.
	// Required.
	Driver string `json:"driver"`

	// SnapshotHandle is the unique snapshot id returned by the CSI volume
	// plugin’s CreateSnapshot to refer to the snapshot on all subsequent calls.
	// Required.
	SnapshotHandle string `json:"snapshotHandle"`

	// Timestamp when the point-in-time snapshot is taken on the storage
	// system. This timestamp will be generated by the CSI volume driver after
	// the snapshot is cut. The format of this field should be a Unix nanoseconds
	// time encoded as an int64. On Unix, the command `date +%s%N` returns
	// the  current time in nanoseconds since 1970-01-01 00:00:00 UTC.
	CreationTime *int64 `json:"creationTime,omitempty" protobuf:"varint,3,opt,name=creationTime"`
	
	// When restoring volume from the snapshot, the volume size should be equal or 
	// larger than the Restoresize if it is specified. If RestoreSize is set to nil, it means
	// that the storage plugin does not have this information available.
	// +optional
	RestoreSize *resource.Quantity `json:"restoreSize" protobuf:"bytes,2,opt,name=restoreSize"`
}

```

#### The `VolumeSnapshotClass` Object

A new VolumeSnapshotClass API object will be added instead of reusing the existing StorageClass, in order to avoid mixing parameters between snapshots and volumes. Each CSI Volume Driver can have its own default VolumeSnapshotClass. If VolumeSnapshotClass is not provided, a default will be used. It allows to add new parameters for snapshots.

```

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotClass describes the parameters used by storage system when
// provisioning VolumeSnapshots from PVCs.
// The name of a VolumeSnapshotClass object is significant, and is how users can request a particular class.
type VolumeSnapshotClass struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Snapshotter is the driver expected to handle this VolumeSnapshotClass.
	Snapshotter string `json:"snapshotter" protobuf:"bytes,2,opt,name=snapshotter"`

	// Parameters holds parameters for the snapshotter.
	// These values are opaque to the system and are passed directly
	// to the snapshotter.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`
}


```
### Volume API Changes

With Snapshot API available, users could provision volumes from snapshot and data will be pre-populated to the volumes. Also considering clone and other possible storage operations, there could be many different types of sources used for populating the data to the volumes. In this proposal, we add a general "DataSource" which could be used to represent different types of data sources.

#### The `DataSource` Object in PVC

Add a new `DataSource` field into PVC to represent the source of the data which is populated to the provisioned volume. External-provisioner will check `DataSource` field and try to provision volume from the sources. In the first version, only VolumeSnapshot is the supported `Type` for data source object reference. Other types will be added in a future version. If unsupported `Type` is used, the PV Controller SHALL fail the operation. Please see more details in [here](https://github.com/kubernetes/community/pull/2495)

Possible `DataSource` types may include the following:

    * VolumeSnapshot: restore snapshot to a new volume
    * PersistentVolumeClaim: clone volume which is represented by PVC

```
type PersistentVolumeClaimSpec struct {
        // If specified when creating, volume will be prepopulated with data from the DataSource.
        // +optional
        DataSource *TypedLocalObjectReference `json:"dataSource" protobuf:"bytes,2,opt,name=dataSource"`
}

```

Add a TypedLocalObjectReference in core API.

```

// TypedLocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type TypedLocalObjectReference struct {
        // Name of the object reference.
        Name string
	// Kind indicates the type of the object reference.
	Kind string
}

```

### Snapshot Controller Design
As the figure below shows, the CSI snapshot controller architecture consists of an external snapshotter which talks to out-of-tree CSI Volume Driver over socket (/run/csi/socket by default, configurable by -csi-address). External snapshotter is part of Kubernetes implementation of [Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec). It is an external controller that monitors `VolumeSnapshot` and `VolumeSnapshotContent` objects and creates/deletes snapshot.
![CSI Snapshot Diagram](csi-snapshot_diagram.png?raw=true "CSI Snapshot Diagram")

* External snapshotter uses ControllerGetCapabilities to find out if CSI driver supports CREATE_DELETE_SNAPSHOT calls. It degrades to trivial mode if not.

* External snapshotter is responsible for creating/deleting snapshots and binding snapshot and SnapshotContent objects. It follows [controller](/contributors/devel/sig-api-machinery/controllers.md) pattern and uses informers to watch for `VolumeSnapshot` and `VolumeSnapshotContent` create/update/delete events. It filters out `VolumeSnapshot` instances with `Snapshotter==<CSI driver name>` and processes these events in workqueues with exponential backoff. 

* For dynamically created snapshot, it should have a VolumeSnapshotClass associated with it. User can explicitly specify a VolumeSnapshotClass in the VolumeSnapshot API object. If user does not specify a VolumeSnapshotClass, a default VolumeSnapshotClass created by the admin will be used. This is similar to how a default StorageClass created by the admin will be used for the provisioning of a PersistentVolumeClaim.

* For statically binding snapshot, user/admin must specify bi-pointers correctly for both VolumeSnapshot and VolumeSnapshotContent, so that the controller knows how to bind them. Otherwise, if VolumeSnapshot points to a non-exist VolumeSnapshotContent, or VolumeSnapshotContent does not point back to the VolumeSnapshot, the Error status will be set for VolumeSnapshot

* External snapshotter is running in the sidecar along with external-attacher and external-provisioner for each CSI Volume Driver.

* In current design, when the storage system fails to create snapshot, retry will not be performed in the controller. This is because users may not want to retry when taking consistent snapshots or scheduled snapshots when the timing of the snapshot creation is important. In a future version, a maxRetries flag or retry termination timestamp will be added to allow users to control whether retries are needed.


#### Changes in CSI External Provisioner

`DataSource` is available in `PersistentVolumeClaim` to represent the source of the data which is prepopulated to the provisioned volume. The operation of the provisioning of a volume from a snapshot data source will be handled by the out-of-tree CSI External Provisioner. The in-tree PV Controller will handle the binding of the PV and PVC once they are ready.
 

#### CSI Volume Driver Snapshot Support

The out-of-tree CSI Volume Driver creates a snapshot on the backend storage system or cloud provider, and calls CreateSnapshot through CSI ControllerServer and returns CreateSnapshotResponse. The out-of-tree CSI Volume Driver needs to implement the following functions:

* CreateSnapshot, DeleteSnapshot, and create volume from snapshot if it supports CREATE_DELETE_SNAPSHOT.
* ListSnapshots if it supports LIST_SNAPSHOTS.

ListSnapshots can be an expensive operation because it will try to list all snapshots on the storage system. For a storage system that takes nightly periodic snapshots, the total number of snapshots on the system can be huge. Kubernetes should try to avoid this call if possible. Instead, calling ListSnapshots with a specific snapshot_id as filtering to query the status of the snapshot will be more desirable and efficient.

CreateSnapshot is a synchronous function and it must be blocking until the snapshot is cut. For cloud providers that support the uploading of a snapshot as part of creating snapshot operation, CreateSnapshot function must also be blocking until the snapshot is cut and after that it shall return an operation pending gRPC error code until the uploading process is complete.

Refer to [Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec) for detailed instructions on how CSI Volume Driver shall implement snapshot functions.


## Transition to the New Snapshot Support

### Existing Implementation in External Storage Repo

For the snapshot implementation in [external storage repo](https://github.com/kubernetes-incubator/external-storage/tree/master/snapshot), an external snapshot controller and an external provisioner need to be deployed.

* The old implementation does not support CSI volume drivers.
* VolumeSnapshotClass concept does not exist in the old design.
* To restore a volume from the snapshot, however, user needs to create a new StorageClass that is different from the original one for the PVC.

Here is an example yaml file to create a snapshot in the old design:

```GO

apiVersion: volumesnapshot.external-storage.k8s.io/v1
kind: VolumeSnapshot
metadata:
  name: hostpath-test-snapshot
spec:
  persistentVolumeClaimName: pvc-test-hostpath

```

### New Snapshot Design for CSI

For the new snapshot model, a sidecar "Kubernetes to CSI" proxy container called "external-snapshotter" needs to be deployed in addition to the sidecar container for the external provisioner. This deployment model is shown in the CSI Snapshot Diagram in the CSI External Snapshot Controller section.

* The new design supports CSI volume drivers.
* To create a snapshot for CSI, a VolumeSnapshotClass can be created and specified in the spec of VolumeSnapshot.
* To restore a volume from the snapshot, users could use the same StorageClass that is used for the original PVC.

Here is an example to create a VolumeSnapshotClass and to create a snapshot in the new design:

```GO

apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshotClass
metadata:
  name: csi-hostpath-snapclass
snapshotter: csi-hostpath
---
apiVersion:snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: snapshot-demo
spec:
  snapshotClassName: csi-hostpath-snapclass
  source:
    name: hpvc
    kind: PersistentVolumeClaim

```
