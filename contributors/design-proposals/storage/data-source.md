# Add DataSource for Volume Operations 

Note: this proposal is part of [Volume Snapshot](https://github.com/kubernetes/community/pull/2335) feature design, and also relevant to recently proposed Volume Clone feature. 

## Goal
In Volume Snapshot proposal, a snapshot is now represented as first-class CRD objects and an external snapshot controller is responsible for managing its lifecycle. With Snapshot API available, users could provision volumes from snapshot and data will be prepopulated to the volumes. Also considering clone and other possible storage operations, there could be many different types of sources used for populating the data to the volumes. In this proposal, we add a general "DataSource" which could be used to represent different types of data sources.

## Design
### API Change
A new DataSource field is proposed to add to PVC to represent the source of the data which is pre-populated to the provisioned volume. For DataSource field, we propose to define a new type “TypedLocalObjectReference”. It is similar to “LocalObjectReference” type with additional Kind field in order to support multiple data source types. In the alpha version, this data source is restricted in the same namespace of the PVC. The following are the APIs we propose to add.

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
}

```
### Error Handling
If an external-provisioner does not understand the new DataSource field and cannot populate the data to the volume, we propose to add a PV condition to detect this situation and fail the operation. The idea is simlar to the design of ["Pod Ready++"]https://github.com/kubernetes/community/blob/master/keps/sig-network/0007-pod-ready%2B%2B.md

First we introduce a condition called "VolumeDataPopulated" into PV. With the feature of snapshot and data source, volume might be created from snapshot source with data populated into the volume instead of an empty one. The external provision needs to understand the data source and also mark the condition "VolumeDataPopualted" to true. The PV controller will check this condition if the data source of PVC is specified. Only if the condition is true, the PVC can be marked as "Bound" phase.

```
// PersistentVolumeStatus is the current status of a persistent volume.
type PersistentVolumeStatus struct {
	// Phase indicates if a volume is available, bound to a claim, or released by a claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	// +optional
	Phase PersistentVolumePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumePhase"`
	...
	
	// +optional
	Conditions []PersistentVolumeCondition
}

type PersistentVolumeConditionType string

// These are valid conditions of Pvc
const (
	// An user trigger resize of pvc has been started
	PersistentVolumeDataPopulated PersistentVolumeConditionType = "dataPopulated"
)

type PersistentVolumeCondition struct {
	Type   PersistentVolumeConditionType
	Status ConditionStatus
	// +optional
	LastProbeTime metav1.Time
	// +optional
	LastTransitionTime metav1.Time
	// +optional
	Reason string
	// +optional
	Message string
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


 
 
