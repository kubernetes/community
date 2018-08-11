---

title: Attacher/Detacher refactor for local storage

authors:
- "@NickrenREN"

owning-sig: sig-storage

participating-sigs:
  - nil

reviewers:
  - "@msau42"
  - "@jsafrane"

approvers:
  - "@jsafrane"
  - "@msau42"
  - "@saad-ali"

editor: TBD

creation-date: 2018-07-30

last-updated: 2018-07-30

status: provisional

---

## Table of Contents
 * [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
* [Implementation](#implementation)
    * [Volume plugin interface change](#volume-plugin-interface-change)
    * [MountVolume/UnmountDevice generation function change](#MountVolume/UnmountDevice-generation-function-change)
    * [Volume plugin change](#volume-plugin-change)
* [Future](#future)

## Summary

Today, the workflow for a volume to be used by pod is:

- attach a remote volume to the node instance (if it is attachable)
- wait for the volume to be attached (if it is attachable)
- mount the device to a global path (if it is attachable)
- mount the global path to a pod directory

It is ok for remote block storage plugins which have a remote attach api,such as `GCE PD`, `AWS EBS`
and remote fs storage plugins such as `NFS`, and `Cephfs`.

But it is not so good for plugins which need local attach such as `fc`, `iscsi` and `RBD`.

It is not so good for local storage neither which is not attachable but needs `MountDevice`


## Motivation

### Goals

 Update Attacher/Detacher interfaces for local storage

### Non-Goals

 Update `fc`, `iscsi` and `RBD` implementation according to the new interfaces

## Proposal

Here we propose to only update the Attacher/Detacher interfaces for local storage.
We may expand it in future to `iscsi`, `RBD` and `fc`, if we figure out how to prevent multiple local attach without implementing attacher interface.

## Implementation

### Volume plugin interface change

We can create a new interface `DeviceMounter`,  move `GetDeviceMountPath` and `MountDevice` from `Attacher`to it.

We can put `DeviceMounter` in `Attacher` which means any one who implements the `Attacher` interface must implement `DeviceMounter`.

```
// Attacher can attach a volume to a node.
type Attacher interface {
        DeviceMounter
        
        // Attaches the volume specified by the given spec to the node with the given Name.
        // On success, returns the device path where the device was attached on the
        // node.
        Attach(spec *Spec, nodeName types.NodeName) (string, error)

        // VolumesAreAttached checks whether the list of volumes still attached to the specified
        // node. It returns a map which maps from the volume spec to the checking result.
        // If an error is occurred during checking, the error will be returned
        VolumesAreAttached(specs []*Spec, nodeName types.NodeName) (map[*Spec]bool, error)

        // WaitForAttach blocks until the device is attached to this
        // node. If it successfully attaches, the path to the device
        // is returned. Otherwise, if the device does not attach after
        // the given timeout period, an error will be returned.
        WaitForAttach(spec *Spec, devicePath string, pod *v1.Pod, timeout time.Duration) (string, error)
}

// DeviceMounter can mount a block volume to a global path.
type DeviceMounter interface {
        // GetDeviceMountPath returns a path where the device should
        // be mounted after it is attached. This is a global mount
        // point which should be bind mounted for individual volumes.
        GetDeviceMountPath(spec *Spec) (string, error)

        // MountDevice mounts the disk to a global path which
        // individual pods can then bind mount
        // Note that devicePath can be empty if the volume plugin does not implement any of Attach and WaitForAttach methods.
        MountDevice(spec *Spec, devicePath string, deviceMountPath string) error
}

```

Note: we also need to make sure that if our plugin implements the `DeviceMounter` interface, 
then executing mount operation from multiple pods referencing the same volume in parallel should be avoided,
even if it does not implement the  `Attacher` interface. 

Since `NestedPendingOperations` can achieve this by setting the same volumeName and same or empty podName in one operation,
we just need to add another check in `MountVolume`: check if the volume is DeviceMountable.

We also need to create another new interface `DeviceUmounter`, and move `UnmountDevice` to it.
```
// Detacher can detach a volume from a node.
type Detacher interface {
        DeviceUnmounter
        
        // Detach the given volume from the node with the given Name.
        // volumeName is name of the volume as returned from plugin's
        // GetVolumeName().
        Detach(volumeName string, nodeName types.NodeName) error
}

// DeviceUnmounter can unmount a block volume from the global path.
type DeviceUnmounter interface {
        // UnmountDevice unmounts the global mount of the disk. This
        // should only be called once all bind mounts have been
        // unmounted.
        UnmountDevice(deviceMountPath string) error
}
```
Accordingly, we need to create a new interface `DeviceMountableVolumePlugin` and move `GetDeviceMountRefs` to it.
```
// AttachableVolumePlugin is an extended interface of VolumePlugin and is used for volumes that require attachment
// to a node before mounting.
type AttachableVolumePlugin interface {
        DeviceMountableVolumePlugin 
        NewAttacher() (Attacher, error)
        NewDetacher() (Detacher, error)
}

// DeviceMountableVolumePlugin is an extended interface of VolumePlugin and is used
// for volumes that requires mount device to a node before binding to volume to pod.
type DeviceMountableVolumePlugin interface {
        VolumePlugin
        NewDeviceMounter() (DeviceMounter, error)
        NewDeviceUmounter() (DeviceUmounter, error)
        GetDeviceMountRefs(deviceMountPath string) ([]string, error)
}
```

### MountVolume/UnmountDevice generation function change

Currently we will check if the volume plugin is attachable in `GenerateMountVolumeFunc`, if it is, we need to call `WaitForAttach` ,`GetDeviceMountPath` and `MountDevice` first, and then set up the volume.

After the refactor, we can split that into three sections: check if volume is attachable, check if it is deviceMountable and set up the volume.
```
devicePath := volumeToMount.DevicePath
if volumeAttacher != nil {
        devicePath, err = volumeAttacher.WaitForAttach(
                volumeToMount.VolumeSpec, devicePath, volumeToMount.Pod, waitForAttachTimeout)
        if err != nil {
                // On failure, return error. Caller will log and retry.
                return volumeToMount.GenerateError("MountVolume.WaitForAttach failed", err)
        }
        // Write the attached device path back to volumeToMount, which can be used for MountDevice.
        volumeToMount.DevicePath = devicePath
}

if volumeDeviceMounter != nil {
        deviceMountPath, err :=
                volumeDeviceMounter.GetDeviceMountPath(volumeToMount.VolumeSpec)
        if err != nil {
                // On failure, return error. Caller will log and retry.
                return volumeToMount.GenerateError("MountVolume.GetDeviceMountPath failed", err)
        }
        deviceMountPath, err := volumeDeviceMounter.MountDevice(volumeToMount.VolumeSpec, devicePath, deviceMountPath)
        if err != nil {
                // On failure, return error. Caller will log and retry.
                return volumeToMount.GenerateError("MountVolume.MountDevice failed", err)
        }

        glog.Infof(volumeToMount.GenerateMsgDetailed("MountVolume.MountDevice succeeded", fmt.Sprintf("device mount path %q", deviceMountPath)))

        // Update actual state of world to reflect volume is globally mounted
        markDeviceMountedErr := actualStateOfWorld.MarkDeviceAsMounted(
                volumeToMount.VolumeName)
        if markDeviceMountedErr != nil {
                // On failure, return error. Caller will log and retry.
                return volumeToMount.GenerateError("MountVolume.MarkDeviceAsMounted failed", markDeviceMountedErr)
        }
}
```
Note that since local storage plugin will not implement the Attacher interface, we can get the device path directly from `spec.PersistentVolume.Spec.Local.Path` when we run `MountDevice`

The device unmounting operation will be executed in `GenerateUnmountDeviceFunc`, we can update the device unmounting generation function as below:
```
// Get DeviceMounter plugin
deviceMountableVolumePlugin, err :=
        og.volumePluginMgr.FindDeviceMountablePluginByName(deviceToDetach.PluginName)
if err != nil || deviceMountableVolumePlugin == nil {
        return volumetypes.GeneratedOperations{}, deviceToDetach.GenerateErrorDetailed("UnmountDevice.FindDeviceMountablePluginByName failed", err)
}

volumeDeviceUmounter, err := deviceMountablePlugin.NewDeviceUmounter()
if err != nil {
        return volumetypes.GeneratedOperations{}, deviceToDetach.GenerateErrorDetailed("UnmountDevice.NewDeviceUmounter failed", err)
}

volumeDeviceMounter, err := deviceMountableVolumePlugin.NewDeviceMounter()
if err != nil {
        return volumetypes.GeneratedOperations{}, deviceToDetach.GenerateErrorDetailed("UnmountDevice.NewDeviceMounter failed", err)
}

unmountDeviceFunc := func() (error, error) {
                deviceMountPath, err :=
                                volumeDeviceMounter.GetDeviceMountPath(deviceToDetach.VolumeSpec)
                if err != nil {
                        // On failure, return error. Caller will log and retry.
                        return deviceToDetach.GenerateError("GetDeviceMountPath failed", err)
                }
                refs, err := deviceMountablePlugin.GetDeviceMountRefs(deviceMountPath)

                if err != nil || mount.HasMountRefs(deviceMountPath, refs) {
                        if err == nil {
                                err = fmt.Errorf("The device mount path %q is still mounted by other references %v", deviceMountPath, refs)
                        }
                        return deviceToDetach.GenerateError("GetDeviceMountRefs check failed", err)
                }
                // Execute unmount
                unmountDeviceErr := volumeDeviceUmounter.UnmountDevice(deviceMountPath)
                if unmountDeviceErr != nil {
                        // On failure, return error. Caller will log and retry.
                        return deviceToDetach.GenerateError("UnmountDevice failed", unmountDeviceErr)
                }
                // Before logging that UnmountDevice succeeded and moving on,
                // use mounter.PathIsDevice to check if the path is a device,
                // if so use mounter.DeviceOpened to check if the device is in use anywhere
                // else on the system. Retry if it returns true.
                deviceOpened, deviceOpenedErr := isDeviceOpened(deviceToDetach, mounter)
                if deviceOpenedErr != nil {
                        return nil, deviceOpenedErr
                }
                // The device is still in use elsewhere. Caller will log and retry.
                if deviceOpened {
                        return deviceToDetach.GenerateError(
                                "UnmountDevice failed",
                                fmt.Errorf("the device is in use when it was no longer expected to be in use"))
                }

                ...

                return nil, nil
        }

```

### Volume plugin change

We need to olny implement the DeviceMounter/DeviceUnmounter interface for local storage since it is not attachable.
And we can keep `fc`,`iscsi` and `RBD` unchanged at the first stage.

## Future
Update `iscsi`, `RBD` and `fc` volume plugins accordingly, if we figure out how to prevent multiple local attach without implementing attacher interface.
