# Add DataSource for Volume Operations 

Note: this proposal is part of [Volume Snapshot](https://github.com/kubernetes/community/pull/2335) feature design, and also relevant to recently proposed Volume Clone feature. 

## Goal
In Volume Snapshot proposal, a snapshot is now represented as first-class CRD objects and an external snapshot controller is responsible for managing its lifecycle. With Snapshot API available, users could provision volumes from snapshot and data will be prepopulated to the volumes. Also considering clone and other possible storage operations, there could be many different types of sources used for populating the data to the volumes. In this proposal, we add a general "DataSource" which could be used to represent different types of data sources.

## Design
A new DataSource field is proposed to add to both PVC and PV to represent the source of the data which is prepopulated to the provisioned volume. If an external-provisioner does not understand the new DataSource field and cannot populate the data to the volume, PV/PVC controller should be able to detect that by comparing DataSource field in PV and PVC (i.e., PVC has DataSource but PV does not) and fail the operation. 

For DataSource, we propose to define a new type “TypedLocalObjectReference”. It is similar to “LocalObjectReference” type with additional Kind field in order to support multiple data source types. In the alpha version, this data source is restricted in the same namespace of the PVC. The following are the APIs we propose to add.

```

type PersistentVolumeClaimSpec struct {
        // If specified, volume will be pre-populated with data from the specified data source.
        // +optional
        DataSource *TypedLocalObjectReference `json:"dataSource" protobuf:"bytes,2,opt,name=dataSource"`
}

type PersistentVolumeSpec struct {
        // If specified, volume was pre-populated with data from the specified data source.
        // +optional
        DataSource *ypedLocalObjectReference `json:"dataSourceRef" protobuf:"bytes,2,opt,name=dataSourceRef"`
}

// TypedLocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type TypedLocalObjectReference struct {
        // Name of the object reference.
        Name string
	// Kind indicates the type of the object reference.
	Kind string
}

```

## Use cases
* Use snapshot to backup data: Alice wants to take a snapshot of her Mongo database, and accidentally delete her tables, she wants to restore her volumes from the snapshot.
To create a snapshot for a volume (represented by PVC), use the snapshot.yaml

```
apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: snapshot-pd-1
  namespace: myns
spec:
  source:
    kind: PersistentVolumeClaim
    name: podpvc
  snapshotClassName: snapshot-class
 
 ```
 After snapshot is ready, create a new volume from the snapshot

```
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: snapshot-pvc
  Namespace: myns
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: csi-gce-pd
  dataSource:
    kind: VolumeSnapshot
    name: snapshot-pd-1
  resources:
    requests:
      storage: 6Gi
```

* Clone volume: Bob want to copy the data from one volume to another by cloning the volume.

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

* Import data from Github repo: Alice want to import data from a github repo to her volume. The github repo is represented by a PVC (gitrepo-1). Compare with the user case 2 is that the data source should be the same kind of volume as the provisioned volume for cloning.

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
    name: gitrepo-1
  resources:
    requests:
      storage: 100Gi
```


 
 
