# Local volume management

**Authors:** [mlmhl](https://github.com/mlmhl)

- [Background](#background)
- [Object](#object)
- [Roadmap](#roadmap)
- [Implementation Design](#design)
	- [Kubernetes Side Refactor](#k8s_refactor)
		- [Volum Plugin Interfaces Refactor](#k8s_interfaces_refactor)
		- [Kubelet Volume Manager Refactor](#k8s_vm_refactor)
			- [Mount Device Operation](#k8s_vm_mount)
			- [Umount Device Operation](#k8s_vm_umount)
		- [A/D Controller Refactor](#k8s_ad_controller_refactor)
	- [Volume Plugin Side Refactor](#volume_plugin_refactor)
		- [Mount Device Operation](#volume_plugin_mount_refactor)
		- [Umount Device Operation](#volume_plugin_umount_refactor)

## <a name="background"></a>Background

For all attachable volume plugins, there are three major steps when provisioning a volume for a `Pod`:

1. Attach the volume to according node the pod is scheduled to
2. Mount the attached device to a global mount path
3. Bind the global mount path to a pod specified path

The recycle of a volume is similar, umount/detach the path/device in a reverse order.

Today, A/D controller is responsible for the attach/detach operation of volumes, kubelet is responsible for the other two steps. Before we start, let's make a quick review of the whole workflow:

1. A/D controller watches all scheduled `Pods` from apiserver, and add all volumes in `Pod.Spec.Volumes` to its `DesiredStateOfWorld`. This cache maintains volumes need to be attached
2. A/D controller periodically checks volumes in `DesiredStateOfWorld` to make sure if they are attached. If not, controller will invoke the attach operation of this volume. After attach operation finished, controller will add this volume to `Node.Status.VolumesAttached`
3. Kubelet watches all pod scheduled to its own, and add all volumes in `Pod.Spec.Volumes` to its `DesiredStateOfWorld`
4. Kubelet periodically checks volumes in `DesiredStateOfWorld` to make sure if they are mounted. If not, kubelet will:
	- Wait for the attach operation finished
	- Mount the attached device to the global mount path
	- Bind the global mount path to a pod specified path

This workflow if nicely for all cloud volumes(`AWS EBS`, `GCE PD`, etc.) which have a remote attach API, but not very suitable for local volumes(`fc`, `iscsi`, `rbd`, etc.) which need a local attach operation. So current implementation make the `Attach/Detach` method of these local volumes do nothing, and execute the real attach operations in `WaitForAttach` method. This led to the following problems:

1. Device path missed in `Node.Status.VolumesAttached`.  A/D controller gets volumes' device path from the returned parameter of `Attach` method, and  updates according Node API object. Local volumes' `Attach` method does nothing, so the returned device path is empty. These absent fields may make user confused.

2. Unexpected umount behavior in kubelet. Kubelet will umount the device from node if no pods use a volume, the `UnmountDeviceFunc` will use the device path stored in `Node.Status.VolumesAttached` to check if this device is still opened or not. As this device path is empty for local volumes, there may be unexpected behaviors.

3. Complex and confusing code both in kubernetes side(`kubelet Volumemanager` and `A/D controller`) and volume plugin side, as local volumes have to implement `AttachableVolumePlugin` interface.

## <a name="object"></a>Object

To solve the problems as above, we need to split the `AttachableVolumePlugin` interface into two independent interfaces: One is still called `AttachableVolumePlugin` and another is `DeviceMountableVolumePlugin` . The new `AttachableVolumePlugin` has the same interface with the old one and implemented by the cloud volumes only, `DeviceMountableVolumePlugin` has the device mount/umount interface only and implemented by the local volumes. Accordingly, we also need to split `Attacher` intreface to `Attacher` and `DeviceMounter`, `Detacher` to `Detacher` and `DeviceUmounter`.

## <a name="roadmap"></a>Roadmap

* Volume relative interface refactor
	* Refactor volume plugin interfaces(`AttachableVolumePlugin`) and volume interfaces(`Attacher/Detacher`)
	* Refactor `A/D controller` and `Kubelet VolumeManager` to use new interfaces
* Refactor local volumes to implement `DevicemountableVolumePlugin` only
	* Refactor `rbd`
	* Refactor `fc`, `iscsi`, etc.

## <a name="design"></a>Implementation Design

### <a name="k8s_refactor"></a>Kubernetes Side Refactor

#### <a name="k8s_interfaces_refactor"></a>Volume Plugin Interfaces refactor

We extract the device mount relative methods from `Attacher` to a new interface, and embed this new interface back to `Attacher` so that volumes implement `Attacher` need also implement `DeviceMounter`:

```go
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

// DeviceMounter can mount a volume to node.
type DeviceMounter interface {
	// GetDeviceMountPath returns a path where the device should
	// be mounted after it is attached. This is a global mount
	// point which should be bind mounted for individual volumes.
	GetDeviceMountPath(spec *Spec) (string, error)

	// MountDevice mounts the disk to a global path which individual pods
	// can then bind mount. Returns the global mount path if succeeded,
	// otherwise returns an error.
	MountDevice(spec *Spec, devicePath string, pod *v1.Pod) (string, error)
}
```

A `*v1.Pod` pararmeter is added for `MountDevice` method because we need to use pod's namespace to fetch according secret if the volume is a pod inline volume, or the secret namespace is not set.

Similarly we apply the refactor to `Detacher`:

```go
// Detacher can detach a volume from a node.
type Detacher interface {
	DeviceUmounter
	// Detach the given volume from the node with the given Name.
	// volumeName is name of the volume as returned from plugin's
	// GetVolumeName().
	Detach(volumeName string, nodeName types.NodeName) error
}

// Deviceumounter can umount a volume from a node.
type DeviceUmounter interface {
	// UnmountDevice unmounts the global mount of the disk. This
	// should only be called once all bind mounts have been
	// unmounted.
	UnmountDevice(deviceMountPath string) error
}
```

Accordingly, extract a `DeviceMountableVolumePlugin` interface from `AttachableVolumePlugin`:

```go
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

#### <a name="k8s_vm_refactor"></a>Kubelet Volume Manager refactor

##### <a name="k8s_vm_mount"></a>Mount Device Operation

Currently, kubelet volume manager invokes `operationGenerator.GenerateMountVolumeFunc` to mount volumes.  If the volume plugin is attachable, `GenerateMountVolumeFunc` will check if the volume is attached or not by `WaitForAttach` method, and then mount the device to global path.

After the refactoring of volume/plugin interfaces, Volume manager should mount the device to global mount path only if according volume plugin implemented `DeviceMountableVolumePlugin` interface, the code segment looks like this:

```go
if volumeAttacher != nil {
	devicePath, err := volumeAttacher.WaitForAttach(
		volumeToMount.VolumeSpec, volumeToMount.DevicePath, volumeToMount.Pod, waitForAttachTimeout)
	if err != nil {
		// On failure, return error. Caller will log and retry.
		return volumeToMount.GenerateError("MountVolume.WaitForAttach failed", err)
	}
	// Write the attached device path back to volumeToMount, which can be used for MountDevice.
	volumeToMount.DevicePath = devicePath
}

if volumeDeviceMounter != nil {
	deviceMountPath, err := volumeDeviceMounter.MountDevice(volumeToMount.VolumeSpec, volumeToMount.DevicePath, volumeToMount.Pod)
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

##### <a name="k8s_vm_umount"></a>Umount Device Operation

similarly, volume manager invokes `operationGenerator.GenerateUnmountDeviceFunc` to umount the volume device from global path. Instead of `AttachableVolumePlugin`, we should use `DeviceMountableVolumePlugin` to do this umount operation, this is a simple 'Name Replace' refactor.

Moreover, volume manager uses `AttachedVolume.DevicePath` to check if the device is opened, however, this field is retrieved from `Node.Status.VolumesAttached`, which is empty for non attachable volume plugins, so we retrieve device path from the global mount path for volumes only implement `DeviceMountableVolumePlugin` before the umount operation, this can be done with the `mount.GetDeviceNameFromMount` method.

It must be noted that, although we keep this validation for volumes only implement `DeviceMountableVolumePlugin`, this is usually a useless check. Because we do this check after the `UnmountDevice` method succeeded, for `DeviceMountableVolumePlugin` volumes, this means the device is already detached from the node, so this validation can always be passed. Consider this, we require the volume plugins should also execute this validation.

#### <a name="k8s_ad_controller_refactor"></a>A/D Controller refactor

A/D controller needn't to change any implementation details as it only handle `AttachableVolumePlugin` and we actually don't modify this interface and according volume plugins. The only thing we need to do is remove local volume plugins from the returned list of `ProbeAttachableVolumePlugins` method as they aren't `AttachableVolumePlugin` any more.

### <a name="volume_plugin_refactor"></a>Volume Plugin Side Refactor

Currently, some local volumes('rbd', `fc`, `iscsi`) also implement `AttachableVolumePlugin`. As these volumes don't have a remote `Attach` API, the local attach operation is executed inside `WaitForAttach` method, the umount and detach operation both performed inside `UnmountDevice` method, this implement leads to the problems mentioned in the beginning of this article. So we should make these volume plugins implementing `DeviceMountableVolumePlugin` only.

#### <a name="volume_plugin_mount_refactor"></a>Mount Device Operation

- Move the local attach operation from `Attacher.WaitForAttach` method to `DeviceMounter.UnmountDevice` method. This is a simple and relative small modification.

- Add security check before attaching device. Concretely, as these local volumes are not handled by A/D controller any more, we should check if the volume is requested as `ReadWrite`(in other words, the `ReadOnly` field is `false`) mode or not before attaching device, if so, it means this volume is requested as exclusive, we should try to lock the volume before we execute the local attach operation, and report any error if we couldn't safely use this volume.

#### <a name="volume_plugin_umount_refactor"></a>Umount Device Operation

As mentioned in [volume manager umount operation refactor](#k8s_vm_umount), we should implement the `isDeviceOpened` check inside `UnmountDevice` method for local volumes, so the process steps weill be:

1. Check if the device path is only mounted to global mount path or not, this can be done by pkugins' `GetDeviceMountRefs` method. If the device is mounted to more than one path, report an error to reflect this
2. Umount the device from global mount path
3. Detach the device from this node
4. Remove the device lock if according volume is requested as `ReadWrite` mode
5. Remove global mount path

Only the first step is a new operation, other subsequent operations are exactly the same as current implementation.