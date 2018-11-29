# Add DataSource for Volume Operations 

Note: this proposal is part of [Volume Snapshot](https://github.com/kubernetes/community/pull/2335) feature design, and also relevant to recently proposed [Volume Clone](https://github.com/kubernetes/community/pull/2533) feature. 

## Goal
Currently in Kubernetes, volume plugin only supports to provision an empty volume. With the new storage features (including [Volume Snapshot](https://github.com/kubernetes/community/pull/2335) and [volume clone](https://github.com/kubernetes/community/pull/2533)) being proposed, there is a need to support data population for volume provisioning. For example, volume can be created from a snapshot source, or volume could be cloned from another volume source. Depending on the sources for creating the volume, there are two scenarios
1. Volume provisioner can recognize the source and be able to create the volume from the source directly (e.g., restore snapshot to a volume or clone volume).
2. Volume provisioner does not recognize the volume source, and create an empty volume. Another external component (data populator) could watch the volume creation and implement the logic to populate/import the data to the volume provisioned. Only after data is populated to the volume, the PVC is ready for use.

There could be many different types of sources used for populating the data to the volumes. In this proposal, we propose to add a generic "DataSource" field to PersistentVolumeClaimSpec to represent different types of data sources.

## Design
### API Change
A new DataSource field is proposed to be added to PVC to represent the source of the data which is pre-populated to the provisioned volume. For DataSource field, we propose to define a new type “TypedLocalObjectReference”. It is similar to “LocalObjectReference” type with additional Kind field in order to support multiple data source types. In the alpha version, this data source is restricted in the same namespace of the PVC. The following are the APIs we propose to add.

```

type PersistentVolumeClaimSpec struct {
        // If specified, volume will be pre-populated with data from the specified data source.
        // +optional
        DataSource *TypedLocalObjectReference `json:"dataSource" protobuf:"bytes,2,opt,name=dataSource"`
}

// TypedLocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type TypedLocalObjectReference struct {
        // Name of the object reference.
        Name string
	// Kind indicates the type of the object reference.
	Kind string
	// APIGroup is the group for the resource being referenced
	APIGroup string
}

```
### Design Details
In the first alpha version, we only support data source from Snapshot. So the expected Kind in DataSource has to be "VolumeSnapshot". In this case, provisioner should provision volume and populate data in one step. There is no need for external data populator yet. 

For other types of data sources that require external data populator, volume creation and data population are two separate steps. Only when data is ready, PVC/PV can be marked as ready (Bound) so that users can start to use them. We are working on a separate proposal to address this using similar idea from ["Pod Ready++"](https://github.com/kubernetes/community/blob/master/keps/sig-network/0007-pod-ready%2B%2B.md).

Note: In order to use this data source feature, user/admin needs to update to the new external provisioner which can recognize snapshot data source. Otherwise, data source will be ignored and an empty volume will be created

## Use cases
* Use snapshot to backup data: Alice wants to take a snapshot of her Mongo database, and accidentally delete her tables, she wants to restore her volumes from the snapshot.
To create a snapshot for a volume (represented by PVC), use the snapshot.yaml

```
apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: snapshot-pd-1
  namespace: mynamespace
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
  Namespace: mynamespace
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
  Namespace: mynamespace
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
  Namespace: mynamespace
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


 
 
