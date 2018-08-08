# Skip attach for non-attachable CSI volumes

Author: @jsafrane

## Goal
* Non-attachable CSI volumes should not require external attacher and `VolumeAttachment` instance creation. This will speed up pod startup.

## Motivation
Currently, CSI requires admin to start external CSI attacher for **all** CSI drivers, including those that don't implement attach/detach operation (such as NFS or all ephemeral Secrets-like volumes). Kubernetes Attach/Detach controller always creates `VolumeAttachment` objects for them and always waits until they're reported as "attached" by external CSI attacher.

We want to skip creation of `VolumeAttachment` objects in A/D controller for CSI volumes that don't require 3rd party attach/detach.

## Dependencies
In order to skip both A/D controller attaching a volume and kubelet waiting for the attachment, both of them need to know if a particular CSI driver is attachable or not. In this document we expect that proposal #2514 is implemented and both A/D controller and kubelet has informer on `CSIPlugin` so they can check if a volume is attachable easily.

## Design
### Volume plugins

#### CSI volume plugin
* Rework [`ProbeVolumePlugins`](https://github.com/kubernetes/kubernetes/blob/43f805b7bdda7a5b491d34611f85c249a63d7f97/pkg/volume/csi/csi_plugin.go#L58) to accept informer that watches CSIPlugin. The plugin will store the informer for later.
* Rework `Attach`, `Detach`, `VolumesAreAttached` and `WaitForAttach` to check for `CSIPlugin` instance using the informer.
	* If CSIPlugin for the driver exists and it's attachable, perform usual logic.
	* If CSIPlugin for the driver exists and it's not attachable, return success immediately (basically NOOP).
	* If CSIPlugin for the driver does not exist, perform usual logic (i.e. treat the volume as attachable).
	  * This keeps the behavior the same as in old Kubernetes version without CSIPlugin object.
	  * This also happens when CSIPlugin informer has not been quick enough. It is suggested that CSIPlugin instance is created **before** any pod that uses corresponding CSI driver can run.
	    * In case that CSIPlugin informer (or user) is too slow, CSI volume plugin `Attach()` will create `VolumeAttachment` instance and wait for (non-existing) external attacher to fulfill it. The CSI plugin shall recover when `CSIPlugin` instance is created and skip attach. Any `VolumeAttachment` instance created here will be deleted on `Detach()`, see the next bullet.
* In addition to the above, `Detach()` removes `VolumeAttachment` instance even if the volume is not attachable. This deletes `VolumeAttachment` instances created by old A/D controller or before `CSIPlugin` instance was created.

### A/D controller
* A/D controller must pass a (shared) informer to watch CSIPlugin and pass it to CSI volume plugin in [`ProbeVolumePlugins`](https://github.com/kubernetes/kubernetes/blob/8db5328c4c1f9467ab0d70ccb991a12d4675b6a7/cmd/kube-controller-manager/app/plugins.go#L82).

### Kubelet / VolumeManager
* Kubelet must create a (shared) informer to watch CSIPlugin and pass it to CSI volume plugin in [`ProbeVolumePlugins`](https://github.com/kubernetes/kubernetes/blob/8db5328c4c1f9467ab0d70ccb991a12d4675b6a7/cmd/kubelet/app/plugins.go#L101).

## API
No API changes.

## Upgrade
This chapter covers:
* Upgrade from old Kubernetes that has `CSISkipAttach` disabled to new Kubernetes with `CSISkipAttach` enabled.
* Update from Kubernetes that has `CSISkipAttach` disabled to the same Kubernetes with `CSISkipAttach` enabled.
* Creation of CSIPlugin instance with non-attachable CSI driver.

In all cases listed above, an "attachable" CSI driver becomes non-attachable. Upgrade does not affect attachable CSI drivers, both "old" and "new" Kubernetes processes them in the same way.

For non-attachable volumes, if the volume was attached by "old" Kubernetes (or "new" Kubernetes before CSIPlugin instance was created), it has `VolumeAttachment` instance. It will be deleted by `Detach()`, as it deletes `VolumeAttachment` instance also for non-attachable volumes.

## Downgrade
This chapter covers:
* Downgrade from new Kubernetes that has `CSISkipAttach` enabled to old Kubernetes with `CSISkipAttach disabled.
* Update from Kubernetes that has `CSISkipAttach` enabled to the same Kubernetes with `CSISkipAttach` disabled.
* Deletion of CSIPlugin instance with non-attachable CSI driver.

In all cases listed above, a non-attachable CSI driver becomes "attachable" (i.e. requires external attacher).  Downgrade does not affect attachable CSI drivers, both "old" and "new" Kubernetes processes them in the same way.

For non-attachable volumes, if the volume was mounted by "new" Kubernetes, it has no VolumeAttachment instance. "Old" A/D controller does not know about it. However, it will periodically call plugin's `VolumesAreAttached()` that checks for `VolumeAttachment` presence. Volumes without `VolumeAttachment` will be reported as not attached and A/D controller will call `Attach()` on these.


## Performance considerations

* Flow suggested in this proposal adds new `CSIPlugin` informer both to A/D controller and kubelet. We don't expect any high amount of instances of `CSIPlugin` nor any high frequency of updates. `CSIPlugin` should have negligible impact on performance.

* A/D controller will not create `VolumeAttachment` instances for non-attachable volumes. Etcd load will be reduced.

* On the other hand, all CSI volumes still must go though A/D controller. A/D controller **must** process every CSI volume and kubelet **must** wait until A/D controller marks a volume as attached, even if A/D controller basically does nothing. All CSI volumes must be added to `Node.Status.VolumesInUse` and `Node.Status.VolumesAttached`. This does not introduce any new API calls, all this is already implemented, however this proposal won't reduce `Node.Status` update frequency in any way. If *all* volumes will move to CSI eventually, pod startup will be slower than when using in-tree volume plugins that don't go through A/D controller.

## Implementation

Expected timeline:
* Alpha: 1.12 (behind feature gate `CSISkipAttach`)
* Beta: 1.13 (enabled by default)
* GA: 1.14

## Alternatives considered
A/D controller and kubelet can be easily extended to check if a given volume is attachable. This would make mounting of non-attachable volumes easier, as kubelet would not need to wait for A/D controller to mark the volume as attached. However, there would be issues when upgrading or downgrading Kubernetes (or marking CSIPlugin as attachable or non-attachable, which has basically the same handling).
* On upgrade (i.e. a previously attachable CSI volume becomes non-attachable), A/D controller could discover that an attached volume is not attachable any longer. A/D controller could clean up `Node.Status.VolumesAttached`, but since A/D controller does not know anything about `VolumeAttachment`, we would either need to introduce a new volume plugin call to clean it up in CSI volume plugin, or something else would need to clean it.
* On downgrade (i.e. a previously non-attachable CSI volume becomes attachable), kubelet must discover that already mounted volume has changed from non-attachable to attachable and put it into `Node.Status.VolumesInUse`. This would race with A/D controller detaching the volume when a pod was deleted at the same time a CSIPlugin instance was made attachable.

Passing all volumes through A/D controller saves us from these difficulties and even races.
