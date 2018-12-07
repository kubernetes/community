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
* As a user I am running an application on glusterfs. I should be able to resize the gluster volume without losing data or mount point. (*online and with data and without taking pod offline*)
* In the logging project we run on dedicated clusters, we start out with 187Gi PVs for each of the elastic search pods. However, the amount of logs being produced can vary greatly from one cluster to another and its not uncommon that these volumes fill and we need to grow them.

## Volume Plugin Matrix


| Volume Plugin   | Supports Resize   | Requires File system Resize | Supported in 1.8 Release |
| ----------------| :---------------: | :--------------------------:| :----------------------: |
| EBS             | Yes               | Yes                         | Yes                      |
| GCE PD          | Yes               | Yes                         | Yes                      |
| GlusterFS       | Yes               | No                          | Yes                      |
| Cinder          | Yes               | Yes                         | Yes                      |
| Vsphere         | Yes               | Yes                         | No                       |
| Ceph RBD        | Yes               | Yes                         | No                       |
| Host Path       | No                | No                          | No                       |
| Azure Disk      | Yes               | Yes                         | No                       |
| Azure File      | No                | No                          | No                       |
| Cephfs          | No                | No                          | No                       |
| NFS             | No                | No                          | No                       |
| Flex            | Yes               | Maybe                       | No                       |
| LocalStorage    | Yes               | Yes                         | No                       |


## Implementation Design

For volume type that requires both file system expansion and a volume plugin based modification, growing persistent volumes will be two
step process.


For volume types that only require volume plugin based api call, this will be one step process.

### Prerequisite

* `pvc.spec.resources.requests.storage` field of pvc object will become mutable after this change.
* #sig-api-machinery has agreed to allow pvc's status update from kubelet as long as pvc and node relationship
  can be validated by node authorizer.
* This feature will be protected by an alpha feature gate, so as API changes needed for it.


### Admission Control and Validations

* Resource quota code has to be updated to take into account PVC expand feature.
* In case volume plugin doesn’t support resize feature. The resize API request will be rejected and PVC object will not be saved. This check will be performed via an admission controller plugin.
* In case requested size is smaller than current size of PVC. A validation will be used to reject the API request. (This could be moved to admission controller plugin too.)
* Not all PVCs will be resizable even if underlying volume plugin allows that. Only dynamically provisioned volumes
which are explicitly enabled by an admin will be allowed to be resized. A plugin in admission controller will forbid
size update for PVCs for which resizing is not enabled by the admin.
* The design proposal for raw block devices should make sure that, users aren't able to resize raw block devices.


### Controller Manager resize

A new controller called `volume_expand_controller` will listen for pvc size expansion requests and take action as needed. The steps performed in this
new controller will be:

* Watch for pvc update requests and add pvc to controller's work queue if a increase in volume size was requested. Once PVC is added to
  controller's work queue - `pvc.Status.Conditions` will be updated with `ResizeStarted: True`.
* For unbound or pending PVCs - resize will trigger no action in `volume_expand_controller`.
* If `pv.Spec.Capacity` already is of size greater or equal than requested size, similarly no action will be performed by the controller.
* A separate goroutine will read work queue and perform corresponding volume resize operation. If there is a resize operation in progress
  for same volume then resize request will be pending and retried once previous resize request has completed.
* Controller resize in effect will be level based rather than edge based. If there are more than one pending resize request for same PVC then
  new resize requests for same PVC will replace older pending request.
* Resize will be performed via volume plugin interface, executed inside a goroutine spawned by `operation_executor`.
* A new plugin interface called `volume.Expander` will be added to volume plugin interface. The `Expander` interface
  will also define if volume requires a file system resize:

  ```go
    type Expander interface {
      // ExpandVolume expands the volume
      ExpandVolumeDevice(spec *Spec, newSize resource.Quantity, oldSize resource.Quantity) error
      RequiresFSResize() bool
    }
  ```

* The controller call to expand the PVC will look like:

```go
func (og *operationGenerator) GenerateExpandVolumeFunc(
    pvcWithResizeRequest *expandcache.PvcWithResizeRequest,
    resizeMap expandcache.VolumeResizeMap) (func() error, error) {

    volumePlugin, err := og.volumePluginMgr.FindExpandablePluginBySpec(pvcWithResizeRequest.VolumeSpec)
    expanderPlugin, err := volumePlugin.NewExpander(pvcWithResizeRequest.VolumeSpec)


    expandFunc := func() error {
        expandErr := expanderPlugin.ExpandVolumeDevice(pvcWithResizeRequest.ExpectedSize, pvcWithResizeRequest.CurrentSize)

        if expandErr != nil {
            og.recorder.Eventf(pvcWithResizeRequest.PVC, v1.EventTypeWarning, kevents.VolumeResizeFailed, expandErr.Error())
            resizeMap.MarkResizeFailed(pvcWithResizeRequest, expandErr.Error())
            return expandErr
        }

        // CloudProvider resize succeeded - lets mark api objects as resized
        if expanderPlugin.RequiresFSResize() {
            err := resizeMap.MarkForFileSystemResize(pvcWithResizeRequest)
            if err != nil {
                og.recorder.Eventf(pvcWithResizeRequest.PVC, v1.EventTypeWarning, kevents.VolumeResizeFailed, err.Error())
                return err
            }
        } else {
            err := resizeMap.MarkAsResized(pvcWithResizeRequest)

            if err != nil {
                og.recorder.Eventf(pvcWithResizeRequest.PVC, v1.EventTypeWarning, kevents.VolumeResizeFailed, err.Error())
                return err
            }
        }
        return nil

    }
    return expandFunc, nil
}
```

* Once volume expand is successful, the volume will be marked as expanded and new size will be updated in `pv.spec.capacity`. Any errors will be reported as *events* on PVC object.
* If resize failed in above step, in addition to events - `pvc.Status.Conditions` will be updated with `ResizeFailed: True`. Corresponding error will be added to condition field as well.
* Depending on volume type next steps would be:

    * If volume is of type that does not require file system resize, then `pvc.status.capacity` will be immediately updated to reflect new size. This would conclude the volume expand operation. Also `pvc.Status.Conditions` will be updated with `Ready: True`.
    * If volume is of type that requires file system resize then a file system resize will be performed on kubelet. Read below for steps that will be performed for file system resize.

* If volume plugin is of type that can not do resizing of attached volumes (such as `Cinder`) then `ExpandVolumeDevice` can return error by checking for
  volume status with its own API (such as by making Openstack Cinder API call in this case). Controller will keep trying to resize the volume until it is
  successful.

* To consider cases of missed PVC update events, an additional loop will reconcile bound PVCs with PVs. This additional loop will loop through all PVCs
  and match `pvc.spec.resources.requests` with `pv.spec.capacity` and add PVC in `volume_expand_controller`'s work queue if `pv.spec.capacity` is less
  than `pvc.spec.resources.requests`.

* There will be additional checks in controller that grows PV size - to ensure that we do not make volume plugin API calls that can reduce size of PV.

### File system resize on kubelet

A File system resize will be pending on PVC until a new pod that uses this volume is scheduled somewhere. While theoretically we *can* perform
online file system resize if volume type and file system supports it - we are leaving it for next iteration of this feature.

#### Prerequisite of File system resize

* `pv.spec.capacity` must be greater than `pvc.status.spec.capacity`.
* A fix in pv_controller has to made to fix `claim.Status.Capacity` only during binding. See comment by jan here - https://github.com/kubernetes/community/pull/657#discussion_r128008128
* A fix in attach_detach controller has to be made to prevent fore detaching of volumes that are undergoing resize.
This can be done by checking `pvc.Status.Conditions` during force detach. `AttachedVolume` struct doesn't hold a reference to PVC - so PVC info can either be directly cached in `AttachedVolume` along with PV spec or it can be fetched from PersistentVolume's ClaimRef binding info.

#### Steps for resizing file system available on Volume

* When calling `MountDevice` or `Setup` call of volume plugin, volume manager will in addition compare `pv.spec.capacity` and `pvc.status.capacity` and if `pv.spec.capacity` is greater
  than `pvc.status.spec.capacity` then volume manager will additionally resize the file system of volume.
* The call to resize file system will be performed inside `operation_generator.GenerateMountVolumeFunc`.  `VolumeToMount` struct will be enhanced to store PVC as well.
* The flow of file system resize will be as follow:
    * Perform a resize based on file system used inside block device.
    * If resize succeeds, proceed with mounting the device as usual.
    * If resize failed with an error that shows no file system exists on the device, then log a warning and proceed with format and mount.
    * If resize failed with any other error then fail the mount operation.
* Any errors during file system resize will be added as *events* to Pod object and mount operation will be failed.
* If there are any errors during file system resize `pvc.Status.Conditions` will be updated with `ResizeFailed: True`. Any errors will be added to
  `Conditions` field.
* File System resize will not be performed on kubelet where volume being attached is ReadOnly. This is similar to pattern being used for performing formatting.
* After file system resize is successful, `pvc.status.capacity` will be updated to match `pv.spec.capacity` and volume expand operation will be considered complete. Also `pvc.Status.Conditions` will be updated with `Ready: True`.

#### Reduce coupling between resize operation and file system type

A file system resize in general requires presence of tools such as `resize2fs` or `xfs_growfs` on the host where kubelet is running. There is a concern
that open coding call to different resize tools directly in Kubernetes will result in coupling between file system and resize operation. To solve this problem
we have considered following options:

1. Write a library that abstracts away various file system operations, such as - resizing, formatting etc.

   Pros:
   * Relatively well known pattern

   Cons:
   * Depending on version with which Kubernetes is compiled with, we are still tied to which file systems are supported in which version
     of kubernetes.
2. Ship a wrapper shell script that encapsulates various file system operations and as long as the shell script supports particular file system
   the resize operation is supported.
   Pros:
   * Kubernetes Admin can easily replace default shell script with her own version and thereby adding support for more file system types.

   Cons:
   * I don't know if there is a pattern that exists in kube today for shipping shell scripts that are called out from code in Kubernetes. Flex is
     different because, none of the flex scripts are shipped with Kubernetes.
3. Ship resizing tools in a container.


Of all options - #3 is our best bet but we are not quite there yet. Hence, I would like to propose that we ship with support for
most common file systems in current release and we revisit this coupling and solve it in next release.

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

### PVC API Change

`pvc.spec.resources.requests.storage` field of pvc object will become mutable after this change.

In addition to that PVC's status will have a `Conditions []PvcCondition` - which will be used
to communicate the status of PVC to the user.

The API change will be protected by Alpha feature gate and api-server will not allow PVCs with
`Status.Conditions` field if feature is not enabled. `omitempty` in serialization format will
prevent presence of field if not set.

So the `PersistentVolumeClaimStatus` will become:

```go
type PersistentVolumeClaimStatus struct {
    Phase PersistentVolumeClaimPhase
    AccessModes []PersistentVolumeAccessMode
    Capacity ResourceList
    // New Field added as part of this Change
    Conditions []PVCCondition
}

// new API type added
type PVCCondition struct {
   Type PVCConditionType
   Status ConditionStatus
   LastProbeTime metav1.Time
   LastTransitionTime metav1.Time
   Reason string
   Message string
}

// new API type
type PVCConditionType string

// new Constants
const (
   PVCReady PVCConditionType = "Ready"
   PVCResizeStarted PVCConditionType = "ResizeStarted"
   PVCResizeFailed  PVCResizeFailed = "ResizeFailed"
)
```

### StorageClass API change

A new field called `AllowVolumeExpand` will be added to StorageClass. The default of this value
will be `false` and only if it is true - PVC expansion will be allowed.

```go
type StorageClass struct {
    metav1.TypeMeta
    metav1.ObjectMeta
    Provisioner string
    Parameters map[string]string
    // New Field added
    // +optional
    AllowVolumeExpand bool
}
```

### Other API changes

This proposal relies on ability to update PVC status from kubelet. While updating PVC's status
a PATCH request must be made from kubelet to update the status.
