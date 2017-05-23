# Growing Persistent Volume size

## Goals

Enable users to increase size of PVs that their pods are using. The user will update PVC for requesting a new size. Underneath we expect that - a controller will apply the change to PV which is bound to the PVC.

## Non Goals

* Reducing size of Persistent Volumes: We realize that, reducing size of PV is way riskier than increasing it. Reducing size of a PV could be a destructive operation and it requires support from underlying file system and volume type. In most cases it also requires that file system being resized is unmounted.

* Rebinding PV and PVC: Kubernetes will only attempt to resize the currently bound PV and PVC and will not attempt to relocate data  from a PV to a new PV and rebind the PVC to newly created PV.

## Use Cases

* As a user I am running Mysql on a 100GB volume - but I am running out of space, I should be able to increase size of volume mysql is using without losing all my data. (*online and with data*)
* As a user I created a PVC requesting 2GB space. I am yet to start a pod with this PVC but I realize that I probably need more space. Without having to create a new PVC, I should be able to request more size with same PVC. (*offline and no data on disk*)
* As a user I was running a rails application with 5GB of assets PVC. I have taken my application offline for maintenance but I would like to grow asset PVC to 10GB in size. (*offline but with data*)

## Volume Plugin Matrix


| Volume Plugin   | Supports Resize   | Requires File system Resize | Supported in 1.7 Release |
| ----------------| :---------------: | :--------------------------:| :----------------------: |
| EBS             | Yes               | Yes                         | Yes                      |
| GCE PD          | Yes               | Yes                         | Yes                      |
| Azure Disk      | Yes               | Yes                         | No                       |
| Cinder          | Yes               | Yes                         | Yes                      |
| Vsphere         | Yes               | Yes                         | No                       |
| Ceph RBD        | Yes               | Yes                         | No                       |
| Host Path       | No                | No                          | No                       |
| GlusterFS       | Yes               | No                          | Yes                      |
| Azure File      | No                | No                          | No                       |
| Cephfs          | No                | No                          | No                       |
| NFS             | No                | No                          | No                       |


## Implementation Design

For volume type that supports growing the PV size, this will be a two step operation:

* A controller in master-controller will listen for PVC events and perform corresponding cloudprovider operation. If successful - controller will store new device size in PV. Some cloudproviders (such as cinder) - do not allow resizing of attached volumes. In such cases - it is upto volume plugin maintainer to decide appropriate behaviour. Volume Plugin maintainer can choose to ignore resize request if disk is attached to a pod (and add appropriate error events to PVC object).  Resize request will keep failing until user corrects the error. User can take necessary action in such cases (such as scale down the pod) which will allow resize to proceed normally.

  In case where volume type requires no file system resize, both PV & PVC objects will be updated accordingly and `status.capacity` of both objects will reflect new size.
  For volume plugins that require file system resize - an additional annotation called `volume.alpha.kubernetes.io/fs-resize-pending` will be added to PV to communicate
  to the Kubelet that File system must be resized when a new pod is started using the PV.

* In case volume plugin doesn’t support resize feature. The resize API request will be rejected and PVC object will not be saved. This check will be performed via an admission controller plugin.

* In case requested size is smaller than current size of PVC. A validation will be used to reject the API request. (This could be moved to admission controller plugin too.)

* There will be additional checks in controller that grows PV size - to ensure that we do not make cloudprovider API calls that can reduce size of PV.

* To consider cases of missed PVC update events, an additional loop will reconcile bound PVCs with PVs.

* Resource Quota code in admission controller has to be updated to consider PVC updates.

* The resize of file system will be performed on kubelet. If there is a running pod - no operation will be performed. Only when a new pod is started using same PVC - then kubelet will match device size and size of pv and attempt a resize of file system.  resizing filesystem will be a volume plugin function. It is upto volume plugin maintainer to correctly implement this. In following cases no resize will be necessary and hence volume plugin can return success without actually doing anything.

  * If disk being attached to the pod is unformatted. In which case since kubelet formats the disk, no resize is necessary.
  * If PVC being attached to pod is of volume type that requires no file system level resize. Such as glusterfs.

  Once file system resize is successful - kubelet will update `pv.spec.status.capacity` and `pvc.spec.status.capacity`field to reflect updated size. Kubelet will also
  update `storageCapacityCondition` and remove the `volume.alpha.kubernetes.io/fs-resize-pending` annotation.

* File System resize will not be performed on kubelet where volume being attached is ReadOnly.
* Once disk has been provisioned with new size, it will be mounted and used in a pod as usual.

## API and UI Design

Given a PVC definition:

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
 name: volume-claim
 annotations:
   volume.beta.kubernetes.io/storage-class: "generalssd"
spec:
 accessModes:
   - ReadWriteOnce
 resources:
   requests:
     storage: 1Gi
```

Users can request new size of underlying PV by simply editing the PVC and requesting new size:

```
~> kubectl edit pvc volume-claim
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
 name: volume-claim
 annotations:
   volume.beta.kubernetes.io/storage-class: "generalssd"
spec:
 accessModes:
   - ReadWriteOnce
 resources:
   requests:
     storage: 10Gi
```

## API Changes

### PV API Change

Two new fields will be added to `PersistentVolumeStatus` object. One is `capacity` and another is `storageCapacityCondition`.

`storageCapacityCondition` field could be just annotation in Alpha. This field will become true if `spec.capacity.storage` and `status.capacity.storage` match their values.
An additional `volume.alpha.kubernetes.io/fs-resize-pending` annotation will be added by controller to indicate that - `PersistentVolume` needs file system resize.


```go
type ResourceList map[ResourceName]resource.Quantity

type PersistentVolumeStatus struct {
        Capacity                 ResourceList
        StorageCapacityCondition bool
}
```

For example - YAML representation of a PV undergoing resize will become:

```yaml
apiVersion: v1
  kind: PersistentVolume
  metadata:
  name: pv0003
  spec:
    capacity:
      # size requested
      storage: 10Gi
    accessModes:
       - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  status:
    capacity:
      # actual size
      storage: 5Gi
    storageCapacityCondition: false
```


### PVC API Change

`pvc.spec.resources.requests.storage` field of pvc object will become mutable after this change.

Similar to PV, PVC API object will have `storageCapacityCondition` field:
`storageCapacityCondition` field could be just annotation in Alpha.

### Other API changes

This proposal relies on ability to update PV & PVC objects from kubelet. Kubelet policy has to be relaxed
to enabled that - https://github.com/kubernetes/kubernetes/blob/master/plugin/pkg/auth/authorizer/rbac/bootstrappolicy/policy.go#L204-L247

Also - an Admin can directly edit the PV and specify new size but controller will not perform
any automatic resize of underlying volume or file system in such cases.
