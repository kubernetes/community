# Add DataSource for Volume Operations 

Note: this proposal is part of [Volume Snapshot](https://github.com/kubernetes/community/pull/2335) feature design, and also relevant to recently proposed Volume Clone feature. 

## Goal
In Volume Snapshot proposal, a snapshot is now represented as first-class CRD objects and an external snapshot controller is responsible for managing its lifecycle. With Snapshot API available, users could provision volumes from snapshot and data will be prepopulated to the volumes. Also considering clone and other possible storage operations, there could be many different types of sources used for populating the data to the volumes. In this proposal, we add a general "DataSource" which could be used to represent different types of data sources.

## Design
The following are the APIs we propose to add

```

type PersistentVolumeClaimSpec struct {
        // If specified, volume will be prepopulated with data from the DataSource.
        // +optional
        DataSource *TypedLocalObjectReference `json:"dataSource" protobuf:"bytes,2,opt,name=dataSource"`
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
* Use snapshot to backup data: Alice wants to take a snapshot of his Mongo database, and accidentally delete her tables, she wants to restore her volumes from the snapshot.
To create a snapshot for a volume (represented by PVC), use the snapshot.yaml

```
apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: snapshot-pd-1
  namespace: myns
spec:
  source:
    Kind: PersistentVolumeClaim
    Name: podpvc
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
    type: VolumeSnapshot
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
    type: PersistentVolumeClaim
    name: pvc-1
  resources:
    requests:
      storage: 10Gi
      
```

 
