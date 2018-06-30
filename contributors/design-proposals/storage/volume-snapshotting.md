Kubernetes Volume Snapshot Proposal
================================

**Authors:** [Jing Xu](https://github.com/jingxu97), [Xing Yang](https://github.com/xing-yang), [Tomas Smetana](https://github.com/tsmetana), [Huamin Chen ](https://github.com/rootfs), [Cindy Wang](https://github.com/ciwang), 

## Background

Many storage systems (GCE PD, Amazon EBS, etc.) provide the ability to create "snapshots" of persistent volumes to protect against data loss. Snapshots can be used in place of a traditional backup system to back up and restore primary and critical data. Snapshots allow for quick data backup (for example, it takes a fraction of a second to create a GCE PD snapshot) and offer fast recovery time objectives (RTOs) and recovery point objectives (RPOs). Snapshots can also be used for data replication, distribution and migration. 

As the initial effort to support snapshot in Kubernetes,  volume snapshotting has been released as a prototype in Kubernetes 1.8. An external controller and provisioner (i.e. two separate binaries) have been added in the [external storage repo](https://github.com/kubernetes-incubator/external-storage/tree/master/snapshot). The prototype currently supports GCE PD, AWS EBS, OpenStack Cinder and Kubernetes hostPath volumes. Volume snapshtos APIs are using [CRD](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/)

To continue that effort, this design is proposed to move the Kubernetes volume snapshot support in-tree by providing snapshot API and snapshot controller in-tree. The volume snapshot feature will support both in-tree and out-of-tree CSI volume drivers. To be consistent with the existing CSI volume driver support documented [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md), a sidecar "Kubernetes to CSI" proxy container called "external-snapshotter" will be provided to watch the Kubernetes API on behalf of the external CSI volume driver and trigger the appropriate operations (i.e., create snapshot and delete snapshot) against the "CSI volume driver" container. The CSI snapshot spec is proposed [here](https://github.com/container-storage-interface/spec/pull/224).

## Objectives

The main goal of this feature is to offer a standardized snapshot API for creating, listing, deleting, and restoring snapshots on an arbitrary volume. For the first version of volume snapshot support in Kubernetes, only on-demand snapshots will be supported. Features listed in the roadmap for future versions are nongoals.

* Goal 1: Enable *on-demand* snapshots of Kubernetes persistent volumes by application developers. Expose standardized snapshotting operations to create, list and delete snapshots in Kubernetes REST API.

    * Nongoal: Enable *automatic* periodic snapshotting for volumes. 
    
* Goal 2: Implement volume snapshotting interface for in-tree plugins including Amazon EBS, GCE PDs, OpenStack Cinder etc.

* Goal 3: Add CSI volume snapshot support

* Goal 4: Offer application-consistent snapshot by providing pre/post snapshot hooks to freeze/unfreeze applications and/or unmount/mount file system.

* Goal 5: Provide a convenient way of creating new and restoring existing volumes from snapshots.

* Goal 6: Provide higher-level management of backing up and restoring a pod and statefulSet.

### Feature Roadmap

Major features, planned for the first version:

* On demand snapshots

    * API to create new snapshots

    * API to list snapshots available to the user

    * API to delete existing snapshots

    * API to create a new persistent volume with persistent volume claim from a snapshot

### Future Features

The features that are not planned for the first version of the API but should be considered in future versions:

* Creating snapshots

    * Support application-consistent snapshots (provide pre/post snapshot hooks)

    * Scheduled and periodic snapshots
    
    * coordinate distributed snapshots across multiple volumes

    * Support snapshot per pod or StatefulSet

    * Enable to create a pod/statefulsets with snapshots

* List snapshots

    * Enable to get the list of all snapshots for a specified persistent volume

    * Enable to get the list of all snapshots for a pod/StatefulSet

* Delete snapshots

    * Enable to automatic garbage collect older snapshots when storage is limited

* In-place restore
    
    * Enable to restore snapshot to a volume that is represented by an existing PVC (PersistentVolumeClaim)

* Quota management

    * Enable to set quota for limiting how many snapshots could be taken and saved

    * When quota is exceeded, delete the oldest snapshots automatically

## Requirements

### Performance

* Time SLA from issuing a snapshot to completion:

* The period we are interested in is the time between the scheduled snapshot time, the time the snapshot is cut (it is safe to write to the snapshotted volume again) and the time the snapshot is finished uploading to its storage location.

* This should be on the order of a few minutes.

### Reliability

* Data corruption

    * Though it is generally recommended to stop application writes before executing the snapshot command, we will not do this automatically for several reasons:

        * GCE and Amazon can create snapshots while the application is running.

        * Stopping application writes cannot be done from the master and varies by application, so doing so will introduce unnecessary complexity and permission issues in the code.

        * Some file systems and server applications are (and should be) able to restore inconsistent snapshots the same way as a disk that underwent an unclean shutdown.

    * The data consistency would be best-effort only: e.g., call fsfreeze prior to the snapshot on filesystems that support it.

    * There are several proposed solutions that would enable the users to specify the action to perform prior to/after the
    snapshots: e.g. use pod annotations. This will be addressed in a different design proposal.

* Snapshot failure

    * Case: Failure during external process, such as during API call or upload

        * Log error, do not attempt to retry

    * Case: Failure within Kubernetes, such as controller restarts

        *  If the master restarts in the middle of a snapshot operation, then the controller will find a snapshot request in pending state and should be able to successfully finish the operation.

## Solution Overview

The following lists the basic volume snapshot functions that will be supported. 

* **Create:**

    1. The user creates a `VolumeSnapshot` referencing a persistent volume claim bound to a persistent volume.

    2. Through pre-snapshot hook, application can be quiesced and file system can be freezed and unmounted. (This will be supported in the next version)

    3. The controller fulfils the `VolumeSnapshot` by creating a snapshot using the volume plugins.

    4. Through post-snapshot hook, application can be resumed and file system will be unfreezed and mounted. (This will be supported in the next version)
    
    5. A new object `VolumeSnapshotData` is created to represent the actual snapshot binding the `VolumeSnapshot` with
       the on-disk snapshot.

* **List:**

    1. The user is able to list all the `VolumeSnapshot` objects in the namespace.

* **Delete:**

    1. The user deletes the `VolumeSnapshot`

    2. The controller removes the on-disk snapshot. Note: snapshots have no notion of "reclaim policy" - there is
       no way to recover the deleted snapshot.

    3. The controller removes the `VolumeSnapshotData` object.

* **Promote snapshot to PV:**

    1. The user creates a persistent volume claim referencing the snapshot object in the annotation.
       Note: The special annotation might get replaced by a dedicated attribute of the
       `PersistentVolumeClaim` in the future.

    2. The controller will use the `VolumeSnapshotData` object to create a persistent volume using the corresponding
       volume snapshot plugin.

    3. The PVC is bound to the newly created PV containing the data from the snapshot.


There are a few uniqueness related to snapshots:

* Both users and admins might create snapshots. Users should only get access to the snapshots belonging to their namespaces. For this aspect, snapshot objects should be in user namespace. Admins might want to choose to expose the snapshots they created to some users who have access to those volumes.

* After snapshots are taken, users might use them to create new volumes or restore the existing volumes back to the time when the snapshot is taken.

* There are use cases that data from snapshots taken from one namespace need to be accessible by users in another namespace.

* For security purpose, if a snapshot object is created by a user, kubernetes should prevent other users duplicating this object in a different namespace if they happen to get the snapshot name.

* There might be some existing snapshots taken by admins/users and they want to use those snapshots through kubernetes API interface.



## API Design

* The `VolumeSnapshot` object

```
// The volume snapshot object accessible to the user. Upon successful creation of the actual
// snapshot by the volume provider it is bound to the corresponding VolumeSnapshotData through
// the VolumeSnapshotSpec
type VolumeSnapshot struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`

	// Spec represents the desired state of the snapshot
	// +optional
	Spec VolumeSnapshotSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the latest observer state of the snapshot
	// +optional
	Status VolumeSnapshotStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
}

type VolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ListMeta  `json:"metadata"`
	Items           []VolumeSnapshot `json:"items"`
}

// VolumeSnapshotSpec is the desired state of the volume snapshot
type VolumeSnapshotSpec struct {
        // PersistentVolumeClaimName is the name of the PVC being snapshotted
        // +optional
        PersistentVolumeClaimName string `json:"persistentVolumeClaimName" protobuf:"bytes,1,opt,name=persistentVolumeClaimName"`

        // SnapshotDataName binds the VolumeSnapshot object with the VolumeSnapshotData
        // +optional
        SnapshotDataName string `json:"snapshotDataName" protobuf:"bytes,2,opt,name=snapshotDataName"`

        // Name of the StorageClass required by the volume snapshot. This
        // StorageClass can be the same as or different from the one used in
        // the source persistent volume claim. If not specified, the StorageClass
        // in the persistent volume claim will be used for creating the snapshot.
	// If persistent volume claim does not have a StorageClass, this field is required.
        // +optional
        StorageClassName string `json:"storageClassName" protobuf:"bytes,3,opt,name=storageClassName"`
}

type VolumeSnapshotStatus struct {
	// The time the snapshot was successfully created
	// +optional
	CreationTimestamp metav1.Time `json:"creationTimestamp" protobuf:"bytes,1,opt,name=creationTimestamp"`

	// Represents the lates available observations about the volume snapshot
	Conditions []VolumeSnapshotCondition `json:"conditions" protobuf:"bytes,2,rep,name=conditions"`
}

type VolumeSnapshotConditionType string

// These are valid conditions of a volume snapshot.
const (
        // VolumeSnapshotConditionCreating means the snapshot is being created but
        // it is not cut yet.
        VolumeSnapshotConditionCreating VolumeSnapshotConditionType = "Creating"
        // VolumeSnapshotConditionUploading means the snapshot is cut and the application
        // can resume accessing data if core_v1.ConditionStatus is True. It corresponds
        // to "Uploading" in GCE PD or "Pending" in AWS and core_v1.ConditionStatus is True.
        // This condition type is not applicable in OpenStack Cinder.
        VolumeSnapshotConditionUploading VolumeSnapshotConditionType = "Uploading"
        // VolumeSnapshotConditionReady is added when the snapshot has been successfully created and is ready to be used.
        VolumeSnapshotConditionReady VolumeSnapshotConditionType = "Ready"
        // VolumeSnapshotConditionError means an error occurred during snapshot creation.
        VolumeSnapshotConditionError VolumeSnapshotConditionType = "Error"
)

// VolumeSnapshot Condition describes the state of a volume snapshot at a certain point.
type VolumeSnapshotCondition struct {
	// Type of replication controller condition.
	Type VolumeSnapshotConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=VolumeSnapshotConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status core_v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message" protobuf:"bytes,5,opt,name=message"`
}
```

* The `VolumeSnapshotData` object

```
// +genclient=true
// +nonNamespaced=true

// VolumeSnapshotData represents the actual "on-disk" snapshot object
type VolumeSnapshotData struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	Metadata metav1.ObjectMeta `json:"metadata"`

	// Spec represents the desired state of the snapshot
	// +optional
	Spec VolumeSnapshotDataSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the latest observed state of the snapshot
	// +optional
	Status VolumeSnapshotDataStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
}

// VolumeSnapshotDataList is a list of VolumeSnapshotData objects
type VolumeSnapshotDataList struct {
        metav1.TypeMeta `json:",inline"`
        Metadata        metav1.ListMeta      `json:"metadata"`
        Items           []VolumeSnapshotData `json:"items"`
}

// The desired state of the volume snapshot
type VolumeSnapshotDataSpec struct {
	// Source represents the location and type of the volume snapshot
	VolumeSnapshotDataSource `json:",inline" protobuf:"bytes,1,opt,name=volumeSnapshotDataSource"`

	// VolumeSnapshotRef is part of bi-directional binding between VolumeSnapshot
	// and VolumeSnapshotData
	// +optional
	VolumeSnapshotRef *core_v1.ObjectReference `json:"volumeSnapshotRef" protobuf:"bytes,2,opt,name=volumeSnapshotRef"`

	// PersistentVolumeRef represents the PersistentVolume that the snapshot has been
	// taken from
	// +optional
	PersistentVolumeRef *core_v1.ObjectReference `json:"persistentVolumeRef" protobuf:"bytes,3,opt,name=persistentVolumeRef"`
}

// VolumeSnapshotDataStatus is the actual state of the volume snapshot
type VolumeSnapshotDataStatus struct {
        // The time the snapshot was successfully created
        // +optional
        CreationTimestamp metav1.Time `json:"creationTimestamp" protobuf:"bytes,1,opt,name=creationTimestamp"`

        // Represents the lates available observations about the volume snapshot
        Conditions []VolumeSnapshotDataCondition `json:"conditions" protobuf:"bytes,2,rep,name=conditions"`
}

// VolumeSnapshotDataConditionType is the type of the VolumeSnapshotData condition
type VolumeSnapshotDataConditionType string

// These are valid conditions of a volume snapshot.
const (
        // VolumeSnapshotDataReady is added when the on-disk snapshot has been successfully created.
        VolumeSnapshotDataConditionReady VolumeSnapshotDataConditionType = "Ready"
        // VolumeSnapshotDataUploading is added when the on-disk snapshot has been cut but is being uploaded.
        VolumeSnapshotDataConditionUploading VolumeSnapshotDataConditionType = "Uploading"
        // VolumeSnapshotDataError is added when the on-disk snapshot is failed to be created
        VolumeSnapshotDataConditionError VolumeSnapshotDataConditionType = "Error"
)

// VolumeSnapshotDataCondition describes the state of a volume snapshot  at a certain point.
type VolumeSnapshotDataCondition struct {
        // Type of volume snapshot condition.
        Type VolumeSnapshotDataConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=VolumeSnapshotDataConditionType"`
        // Status of the condition, one of True, False, Unknown.
        Status core_v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
        // The last time the condition transitioned from one status to another.
        // +optional
        LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`
        // The reason for the condition's last transition.
        // +optional
        Reason string `json:"reason" protobuf:"bytes,4,opt,name=reason"`
        // A human readable message indicating details about the transition.
        // +optional
        Message string `json:"message" protobuf:"bytes,5,opt,name=message"`
}

// Represents the actual location and type of the snapshot. Only one of its members may be specified.
type VolumeSnapshotDataSource struct {
	// HostPath represents a directory on the host.
	// Provisioned by a developer or tester.
	// This is useful for single-node development and testing only!
	// On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	HostPath *HostPathVolumeSnapshotSource `json:"hostPath,omitempty"`
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSnapshotSource `json:"awsElasticBlockStore,omitempty"`
        // CSI (Container Storage Interface) represents storage that handled by an external CSI driver (Alpha feature).
        // +optional
        CSI *CSIVolumeSnapshotSource `json:"csiVolumeSnapshotSource,omitempty"`
        // ... etc. for other snapshot types
}
```

An example of the `VolumeSnapshotDataSource` for Amazon EBS snapshots:

```
// AWS EBS volume snapshot source
type AWSElasticBlockStoreVolumeSnapshotSource struct {
	// Unique id of the persistent disk snapshot resource. Used to identify the disk snapshot in AWS
	SnapshotID string `json:"snapshotId"`
}
```

An example of the `VolumeSnapshotDataSource` for CSI snapshots:

```
// Represents volume snapshot that is managed by an external CSI volume driver (Alpha feature)
type CSIVolumeSnapshotSource struct {
        // Driver is the name of the driver to use for this snapshot.
        // Required.
        Driver string `json:"driver"`

        // SnapshotHandle is the unique snapshot id returned by the CSI volume
        // plugin’s CreateSnapshot to refer to the snapshot on all subsequent calls.
        // Required.
        SnapshotHandle string `json:"snapshotHandle"`

       // Timestamp when the point-in-time snapshot is taken on the storage
       // system. The format of this field should be a Unix nanoseconds time
       // encoded as an int64. On Unix, the command `date +%s%N` returns the
       // current time in nanoseconds since 1970-01-01 00:00:00 UTC. This
       // field is REQUIRED.
       int64 created_at = 4;
}
```


* Add `SnapshotParameters` to the `StorageClass` object

```
type StorageClass struct {
        // SnapshotParameters holds parameters for creating a snapshot.
        // These values are opaque to the system and are passed directly
        // to the provisioner.  The only validation done on keys is that they are
        // not empty.  The maximum number of parameters is
        // 512, with a cumulative max size of 256K
        // +optional
        SnapshotParameters map[string]string
}
```

## Snapshot Controller Design

The in-tree snapshot controller will be watching the add/delete `VolumeSnapshot` events. If an add `VolumeSnapshot` event is received, it compares the `Provisioner` specified in the `StorageClass` in the `VolumeSnapshot` object with the one specified in the `PersistentVolumeClaim` object. If they are different, abort the operation.

It also checks if the PersistentVolumeSource in the PersistentVolumeSpec contains CSI.  If so, the creating/deleting snapshot operation will be handled by the out-of-tree snapshot controller `external-snapshotter`; otherwise it will be handled by the in-tree snapshot controller.

The in-tree snapshot controller binds the `VolumeSnapshot` and `VolumeSnapshotData` API objects after the snapshot is created successfully.

`External-snapshotter` follows [controller](https://github.com/kubernetes/community/blob/master/contributors/devel/controllers.md) pattern and uses informers to watch for `VolumeSnapshot` create/update/delete events. It filters out `VolumeSnapshot` instances with `Snapshotter==<CSI driver name>` and processes these events in workqueues with exponential backoff. The `external-snapshotter` creates `CreateSnapshotRequest` and calls `CreateSnapshot` through the CSI `ControllerClient`. It gets `CreateSnapshotResponse` from the CSI plugin and creates a `VolumeSnapshotData` API object with `VolumeSnapshotDataSource`.


### Create Snapshot Logic

To create a snapshot:

* Acquire operation lock for volume so that only one snapshot creation operation is running for the specified volume

    * Abort if there is already a pending operation.

* Spawn a new thread:

    * Execute the volume-specific logic to create a snapshot of the persistent volume referenced by the PVC.

    * For any errors, log the error, send it as an event on the corresponding `VolumeSnapshot`, and terminate the thread (the main controller will retry as needed).

    * Once a snapshot is created successfully:

        * Make a call to the API server to add the new snapshot ID/timestamp to the `VolumeSnapshotData` API object, update its status.

### Snapshot to PV promotion logic

For the `PersistentVolumeClaim` used to restore the snapshot:

* Check the `PeristentVolumeClaim` annotation and get the `VolumeSnapshot` name.

* Retrieve the `VolumeSnapshot` object from the API server:

    * Verify the `Provisioner` used for `VolumeSnapshot` and `PersistentVolumeClaim` are the same.

    * Verify both the `VolumeSnapshot` and `PersistentVolumeClaim` belong to the same namespace.

    * Verify the `VolumeSnapshotData` referenced by the `VolumeSnapshot` exists and that its `VolumeSnapshot` reference points "back" to the given `VolumeSnapshot` object.

* Find the correct plugin to use to create a `PersistentVolume` bound to the given PVC.

## Example Use Case

### Alice wants to backup her MySQL database data

Alice is a DB admin who runs a MySQL database and needs to backup the data on a remote server prior to the database
upgrade. She has a short maintenance window dedicated to the operation that allows her to pause the dabase only for
a short while. Alice will therefore stop the database, create a snapshot of the data, re-start the database and after
that start time-consuming network transfer to the backup server.

The database is running in a pod with the data stored on a persistent volume:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mysql
  labels:
    name: mysql
spec:
  containers:
    - resources:
        limits :
          cpu: 0.5
      image: openshift/mysql-55-centos7
      name: mysql
      env:
        - name: MYSQL_ROOT_PASSWORD
          value: rootpassword
        - name: MYSQL_USER
          value: wp_user
        - name: MYSQL_PASSWORD
          value: wp_pass
        - name: MYSQL_DATABASE
          value: wp_db
      ports:
        - containerPort: 3306
          name: mysql
      volumeMounts:
        - name: mysql-persistent-storage
          mountPath: /var/lib/mysql/data
  volumes:
    - name: mysql-persistent-storage
      persistentVolumeClaim:
      claimName: claim-mysql
```

The persistent volume is bound to the `claim-mysql` PVC which needs to be snapshotted. Since Alice has some downtime
allowed she may lock the database tables for a moment to ensure the backup would be consistent:
```
mysql> FLUSH TABLES WITH READ LOCK;
```
Now she is ready to create a snapshot of the `claim-mysql` PVC. She creates a vs.yaml:
```yaml
apiVersion: storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: mysql-snapshot
  namespace: default
spec:
  persistentVolumeClaim: claim-mysql
```

```
$ kubectl create -f vs.yaml
```

This will result in a new snapshot being created by the controller. Alice would wait until the snapshot is complete:
```
$ kubectl get volumesnapshots

NAME             STATUS
mysql-snapshot   ready
```
Now it's OK to unlock the database tables and the database may return to normal operation:
```
mysql> UNLOCK TABLES;
```
Alice can now get to the snapshotted data and start syncing them to the remote server. First she needs to promote the
snapshot to a PV by creating a new PVC. Alice can create the PVC referencing the snapshot in the annotations.
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: snapshot-data-claim
annotations:
    snapshot.alpha.kubernetes.io/snapshot: mysql-snapshot
spec:
  accessModes:
    - ReadWriteOnce
```
Once the claim is bound to a persistent volume Alice creates a job to sync the data with a remote backup server:
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: mysql-sync
spec:
  template:
    metadata:
      name: mysql-sync
    spec:
      containers:
      - name: mysql-sync
        image: rsync
        command: "rsync -av /mnt/data alice@backup.example.com:mysql_backups"
      restartPolicy: Never
      volumeMounts:
        - name: snapshot-data
          mountPath: /mnt/data
  volumes:
    - name: snapshot-data
      persistentVolumeClaim:
      claimName: snapshot-data-claim
```

Alice will wait for the job to finish and then may delete both the `snapshot-data-claim` PVC as well as `mysql-snapshot`
request (which will delete also the snapshot object):
```
$ kubectl delete pvc snapshot-data-claim
$ kubectl delete volumesnapshot mysql-snapshot
```
