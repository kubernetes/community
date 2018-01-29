# Proposal for Growing FlexVolume Size

**Authors:** [xingzhou](https://github.com/xingzhou)

## Goals

Since PVC resizing is introduced in Kubernetes v1.8, several volume plugins have already supported this feature, e.g. GlusterFS, AWS EBS. In this proposal, we are proposing to support FlexVolume expansion. So when user uses FlexVolume and corresponding volume driver to connect to his/her backend storage system, he/she can expand the PV size by updating PVC in Kubernetes.

## Non Goals

* We only consider expanding FlexVolume size in this proposal. Decreasing size of FlexVolume will be designed in the future.
* In this proposal, user can only expand the FlexVolume size manually by updating PVC. Auto-expansion of FlexVolume based on specific meterings is not considered.
* The proposal only contains the changes made in FlexVolume, volume driver changes which should be made by user are not included.

## Implementation Designs

### Prerequisites

* Kubernetes should be at least v1.8.
* Enable resizing by setting feature gate `ExpandPersistentVolumeGate` to `true`.
* Enable `PersistentVolumeClaimResize` admission plugin(optional).
* Follow the UI of PV resizing, including:
  * Only dynamic provisioning supports volume resizing
  * Set StorageClass attribute `allowVolumeExpansion` to `true`

### Admission Control Changes

Whether or not a specific volume plugin supports volume expansion is validated and checked in PV resize admission plugin. In general, we can list FlexVolume as the ones that support volume expansion and leave the actual expansion capability check to the underneath volume driver when PV resize controller calls the `ExpandVolumeDevice` method of FlexVolume.

In PV resize admission plugin, add the following check to `checkVolumePlugin` method:
```
// checkVolumePlugin checks whether the volume plugin supports resize
func (pvcr *persistentVolumeClaimResize) checkVolumePlugin(pv *api.PersistentVolume) bool {
  ...
  if pv.Spec.FlexVolume != nil {
    return true
  }
  ...
}
```

### FlexVolume Plugin Changes

FlexVolume relies on underneath volume driver to implement various volume functions, e.g. attach/detach. As a result, volume driver will decide whether volume can be expanded or not. 

By default, we assume all kinds of flex volume drivers support resizing. If they do not, flex volume plugin can detect this during resizing call to flex volume driver and always throw out error to stop the resizing process. So as a result, to implement resizing feature in flex volume plugin, the plugin itself must implement the following `ExpandableVolumePlugin` interfaces:

#### ExpandVolumeDevice

Volume resizing controller invokes this method while receiving a valid PVC resizing request. FlexVolume plugin calls the underneath volume driver’s corresponding `expandvolume` method with three parameters, including new size of volume(number in bytes), old size of volume(number in bytes) and volume spec, to expand PV. Once the expansion is done, volume driver should return the new size(number in bytes) of the volume to FlexVolume.

A sample implementation of `ExpandVolumeDevice` method is like:
```
func (plugin *flexVolumePlugin) ExpandVolumeDevice(spec *volume.Spec, newSize resource.Quantity, oldSize resource.Quantity) (resource.Quantity, error) {
  const timeout = 10*time.Minute
 
  call := plugin.NewDriverCallWithTimeout(expandVolumeCmd, timeout)
  call.Append(newSize.Value())
  call.Append(oldSize.Value())
  call.AppendSpec(spec, plugin.host, nil)

  // If the volume driver does not support resizing, Flex Volume Plugin can throw out error here
  // to stop expand controller's resizing process. 
  ds, err := call.Run()
  if err != nil {
      return resource.NewQuantity(0, resource.BinarySI), err
  }
 
  return  resource.NewQuantity(ds.ActualVolumeSize, resource.BinarySI), nil
}
```

Add a new field in type `DriverStatus` named `ActualVolumeSize` to identify the new expanded size of the volume returned by underneath volume driver:
```
// DriverStatus represents the return value of the driver callout.
type DriverStatus struct {
  ...
  ActualVolumeSize int64 `json:"volumeNewSize,omitempty"`
}
```

#### RequiresFSResize

`RequiresFSResize` is a method to implement `ExpandableVolumePlugin` interface. The return value of this method identifies whether or not a file system resize is required once physical volume get expanded. If the return value is `true`, PV resize controller will consider the volume resize operation is done and then update the PV object’s capacity in K8s directly; If the return value is `false`, PV resize controller will leave kubelet to do the file system resize, and kubelet on worker node will call `ExpandFS` method of FlexVolume to finish the file system resize step(at present, only offline FS resize is supportted, online resize support is under community discussion [here](https://github.com/kubernetes/community/pull/1535)). 

The return value of `RequiresFSResize` is collected from underneath volume driver when FlexVolume invokes `init` method of volume driver. The sample code of `RequiresFSResize` in FlexVolume looks like:
```
func (plugin *flexVolumePlugin) RequiresFSResize() bool {
  return plugin.capabilities.RequiresFSResize
}
```

And as a result, the FlexVolume type `DriverCapability` can be redefined as:
```
type DriverCapabilities struct {
  Attach           bool `json:"attach"`
  RequiresFSResize bool `json:"requiresFSResize"`
  SELinuxRelabel   bool `json:"selinuxRelabel"`
}

func defaultCapabilities() *DriverCapabilities {
  return &DriverCapabilities{
    Attach:           true,
    RequiresFSResize: true, //By default, we require file system resize which will be done by kubelet
    SELinuxRelabel:   true,
  }
}
```

#### ExpandFS

`ExpandFS` is another method to implement `ExpandableVolumePlugin` interface. This method allows volume plugin itself instead of kubelet to resize the file system. If volume plugin returns `true` for `RequiresFSResize`, PV resize controller will leave FS resize to kubelet on worker node. Kubelet then will call FlexVolume `ExpandFS` to resize file system once physical volume expansion is done.

As `ExpandFS` is called on worker node, volume driver can also take this chance to do physical volume resize together with file system resize as well. Also, current code only supports offline FS resize, online resize support is under dicsussion [here](https://github.com/kubernetes/community/pull/1535). Once online resize is implemented, we can also leverage online resize for FlexVolume by `ExpandFS` method.

Note that `ExpandFS` is a new API for `ExpandableVolumeDriver`, the community ticket can be found [here](https://github.com/kubernetes/kubernetes/issues/58786).

`ExpandFS` will call underneath volume driver `expandfs` method to finish FS resize. The sample code looks like:
```
func (plugin *flexVolumePlugin) ExpandFS(spec *volume.Spec, newSize resource.Quantity, oldSize resource.Quantity)  error {
  const timeout = 10*time.Minute
 
  call := plugin.NewDriverCallWithTimeout(expandFSCmd, timeout)
  call.Append(newSize.Value())
  call.Append(oldSize.Value())
  call.AppendSpec(spec, plugin.host, nil)
 
  _, err := call.Run()
 
  return  err
}
```

For more design and details on how kubelet resizes volume file system, please refer to volume resizing proposal at:
https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/grow-volume-size.md


Based on the above design, the resizing process for flex volume can be summarized as:
* If flex volume driver does not support resizing, driver shall not implement `expandvolume` method and flex volume plugin will throw out error to stop the expand volume controller's resizing process.
* If flex volume driver supports resizing, it shall implement `expandvolume` method and at least, the volume driver shall be installed on master node.
* If flex volume driver supports resizing and does not need file system resizing, it shall set "requiresFSResize" capability to `false`. Otherwise kubelet on worker node will call `ExpandFS` to resize the file system.
* If flex volume driver supports resizing and requires file system resizing(`RequiresFSResize` returns `true`), after the physical volume resizing is done, `ExpandFS` will be called from kubelet on worker node.
* If flex volume driver supports resizing and requires to resize the physical volume from worker node, the driver shall be installed on both master node and worker node. The driver on master node can do a non-op process for `ExpandVolumeDevice` and returns success message. For `RequiresFSResize`, driver on master node must return `true`. This process gives drivers on worker nodes a chance to make `physical volume resize` and `file system resize` together through `ExpandFS` call from kubelet. This scenario is useful for some local storage resizing cases.

### Volume Driver Changes

Volume driver needs to implement two new interfaces: `expandvolume` and `expandfs` to support volume resizing.

For `expandvolume`, it takes three parameters: new size of volume(number in bytes), old size of volume(number in bytes) and volume spec json string. `expandvolume` expands the physical backend volume size and return the new size(number in bytes) of volume.

For those volume plugins who need file system resize after physical volume is expanded, the `expandfs` method can take the FS resize work. If volume driver set the `requiresFSResize` capability to true, this method will be called from kubelet on worker node. Volume driver can do the file system resize (or physical volume resize together with file system resize) inside this method

In addition, those volume drivers who support resizing but do not require fils system resizing shall set `requiresFSResize` capability to `false`:
```
if [ "$op" = "init" ]; then
        log '{"status": "Success", "capabilities": {“requiresFSResize”: false}}'
        exit 0
fi
```

### UI

Expand FlexVolume size follows the same process as expanding other volume plugins, like GlusterFS. User creates and binds PVC and PV first. Then by using `kubectl edit pvc xxx` command, user can update the new size of PVC.

## References

* [Proposal for Growing Persistent Volume Size](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/grow-volume-size.md)
* [PR for Volume Resizing Controller](https://github.com/kubernetes/kubernetes/commit/cd2a68473a5a5966fa79f455415cb3269a3f7462)
* [Online FS resize support](https://github.com/kubernetes/community/pull/1535)
* [Add “ExpandFS” method to “ExpandableVolumePlugin” interface](https://github.com/kubernetes/kubernetes/issues/58786)
