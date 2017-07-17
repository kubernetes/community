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

For volume type that requires both file system expansion and a volume plugin based modification, growing persistent volumes will be two
step process.


For volume types that only require volume plugin based api call, this will be one step process.

### Prerequisite

* `pvc.spec.resources.requests.storage` field of pvc object will become mutable after this change.
* #sig-api-machinery has agreed to allow pvc's status update from kubelet as long as pvc and node relationship
  can be validated by node authorizer.
* This feature will be protected by an alpha feature gate.

### Admission Control and Validations

* Resource quota code has to be updated to take into account PVC expand feature.
* In case volume plugin doesn’t support resize feature. The resize API request will be rejected and PVC object will not be saved. This check will be performed via an admission controller plugin.
* In case requested size is smaller than current size of PVC. A validation will be used to reject the API request. (This could be moved to admission controller plugin too.)


### Controller Manager resize

A new controller called `volume_expand_controller` will listen for pvc size expansion requests and take action as needed. The steps performed in this
new controller will be:

* Watch for pvc update requests and add pvc to controller's desired state of world if a increase in volume size was requested.
* A reconciler will read desired state of world and perform corresponding volume resize operation. If there is a resize operation in progress
  for same volume then resize request will be pending and retried once previous resize request has completed.
* Controller resize in effect will be level based rather than edge based. If there are more than one pending resize request for same PVC then
  new resize requests for same PVC will replace older pending request.
* Resize will be performed via volume plugin interface, executed inside a goroutine spawned by `operation_exectutor`.
* A new plugin interface called `volume.Exander` will be added to volume plugin interface. The controller call to expand the PVC will look like:

```go
func (og *operationGenerator) GenerateExpandVolumeFunc(
    pvcWithResizeRequest *expandcache.PvcWithResizeRequest,
    dsow expandcache.DesiredStateOfWorld) (func() error, error) {

    volumePlugin, err := og.volumePluginMgr.FindExpandablePluginBySpec(pvcWithResizeRequest.VolumeSpec)

    if err != nil {
        return nil, fmt.Errorf("Error finding plugin for expanding volume: %q with error %v", pvcWithResizeRequest.UniquePvcKey(), err)
    }

    expanderPlugin, err := volumePlugin.NewExpander()

    if err != nil {
        return nil, fmt.Errorf("Error creating expander plugin for volume %q with error %v", pvcWithResizeRequest.UniquePvcKey(), err)
    }

    expandFunc := func() error {
        expandErr := expanderPlugin.ExpandVolumeDevice(pvcWithResizeRequest.VolumeSpec, pvcWithResizeRequest.ExpectedSize, pvcWithResizeRequest.CurrentSize)

        if expandErr != nil {
            glog.Errorf("Error expanding volume through cloudprovider : %v", expandErr)
            return expandErr
        }
        dsow.MarkAsResized(pvcWithResizeRequest)

        return nil
    }
    return expandFunc, nil
}
```

* Once volume expand is successful, the volume will be marked as expanded and new size will be updated in `pv.spec.capacity`. Any errors will be
reported as *events* on PVC object.
* Depending on volume type next steps would be:

    * If volume is of type that does not require file system resize, then `pvc.status.capacity` will be immediately updated to reflect new size. This would conclude the volume expand operation.
    * If volume if of type that requires file system resize then a file system resize will be performed on kubelet. Read below for steps that will be performed for file system resize.

* If volume plugin is of type that can not do resizing of attached volumes (such as `Cinder`) then `ExpandVolumeDevice` can return error by checking for
  volume status with its own API (such as by making Openstack Cinder API call in this case). Controller will keep trying to resize the volume until it is
  successful.

* To consider cases of missed PVC update events, an additional loop will reconcile bound PVCs with PVs. This additional loop will loop through all PVCs
  and match `pvc.spec.capactiy` with `pv.spec.capacity` and add PVC in `volume_expand_controller`'s desired state of world if `pv.spec.capacity` is less
  than `pvc.spec.capacity`.

* There will be additional checks in controller that grows PV size - to ensure that we do not make volume plugin API calls that can reduce size of PV.

### File system resize on kublet

* When calling `MountDevice` or `Setup` call of volume plugin, volume manager will in addition compare `pv.spec.capacity` and `pvc.status.capacity` and if `pv.spec.capacity` is greater
  than `pvc.status.spec.capacity` then volume manager will additionally resize the file system of volume.
* The call to resize file system will be performed inside `operation_generator.GenerateMountVolumeFunc`.  `VolumeToMount` struct will be enhanced to store PVC as well.
* Any errors during file system resize will be added as *events* to Pod object and mount operation will be failed.
* File System resize will not be performed on kubelet where volume being attached is ReadOnly. This is similar to pattern being used for performing formatting.
* After file system resize is successful, `pvc.status.capacity` will be updated to match `pv.spec.capacity` and volume expand operation will be considered complete.


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

### Other API changes

This proposal relies on ability to update PVC status from kubelet. While updating PVC's status
a PATCH request must be made from kubelet to update the status.

Also - an Admin can directly edit the PV and specify new size but controller will not perform
any automatic resize of underlying volume or file system in such cases.
