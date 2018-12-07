# Skip attach for non-attachable CSI volumes

Author: @jsafrane

## Goal
* Non-attachable CSI volumes should not require external attacher and `VolumeAttachment` instance creation. This will speed up pod startup.

## Motivation
Currently, CSI requires admin to start external CSI attacher for **all** CSI drivers, including those that don't implement attach/detach operation (such as NFS or all ephemeral Secrets-like volumes). Kubernetes Attach/Detach controller always creates `VolumeAttachment` objects for them and always waits until they're reported as "attached" by external CSI attacher.

We want to skip creation of `VolumeAttachment` objects in A/D controller for CSI volumes that don't require 3rd party attach/detach.

## Dependencies
In order to skip both A/D controller attaching a volume and kubelet waiting for the attachment, both of them need to know if a particular CSI driver is attachable or not. In this document we expect that proposal #2514 is implemented and both A/D controller and kubelet has informer on `CSIDriver` so they can check if a volume is attachable easily.

## Design
### CSI volume plugin
* Rework [`Init`](https://github.com/kubernetes/kubernetes/blob/43f805b7bdda7a5b491d34611f85c249a63d7f97/pkg/volume/csi/csi_plugin.go#L58) to get or create informer to cache CSIDriver instances.
  * Depending on where the API for CSIDriver ends up, we may:
    * Rework VolumeHost to either provide the informer. This leaks CSI implementation details to A/D controller and kubelet
    * Or the CSI volume plugin can create and run CSIDriver informer by itself. No other component in controller-manager or kubelet needs the informer right now, so a non-shared informer is viable option. Depending on when the API for CSIDriver ends up, `VolumeHost` may need to be extended to provide client interface to the API and kubelet and A/D controller may need to be updated to create the interface (somewhere in `cmd/`, where RESTConfig is still available to create new clients ) and pass it to their `VolumeHost` implementations.
* Rework `Attach`, `Detach`, `VolumesAreAttached` and `WaitForAttach` to check for `CSIDriver` instance using the informer.
	* If CSIDriver for the driver exists and it's attachable, perform usual logic.
	* If CSIDriver for the driver exists and it's not attachable, return success immediately (basically NOOP). A/D controller will still mark the volume as attached in `Node.Status.VolumesAttached`.
	* If CSIDriver for the driver does not exist, perform usual logic (i.e. treat the volume as attachable).
	  * This keeps the behavior the same as in old Kubernetes version without CSIDriver object.
	  * This also happens when CSIDriver informer has not been quick enough. It is suggested that CSIDriver instance is created **before** any pod that uses corresponding CSI driver can run.
	    * In case that CSIDriver informer (or user) is too slow, CSI volume plugin `Attach()` will create `VolumeAttachment` instance and wait for (non-existing) external attacher to fulfill it. The CSI plugin shall recover when `CSIDriver` instance is created and skip attach. Any `VolumeAttachment` instance created here will be deleted on `Detach()`, see the next bullet.
* In addition to the above, `Detach()` removes `VolumeAttachment` instance even if the volume is not attachable. This deletes `VolumeAttachment` instances created by old A/D controller or before `CSIDriver` instance was created.


### Authorization
* A/D controller and kubelet must be allowed to list+watch CSIDriver instances. Updating RBAC rules should be enough.

## API
No API changes.

## Upgrade
This chapter covers:
* Upgrade from old Kubernetes that has `CSISkipAttach` disabled to new Kubernetes with `CSISkipAttach` enabled.
* Update from Kubernetes that has `CSISkipAttach` disabled to the same Kubernetes with `CSISkipAttach` enabled.
* Creation of CSIDriver instance with non-attachable CSI driver.

In all cases listed above, an "attachable" CSI driver becomes non-attachable. Upgrade does not affect attachable CSI drivers, both "old" and "new" Kubernetes processes them in the same way.

For non-attachable volumes, if the volume was attached by "old" Kubernetes (or "new" Kubernetes before CSIDriver instance was created), it has `VolumeAttachment` instance. It will be deleted by `Detach()`, as it deletes `VolumeAttachment` instance also for non-attachable volumes.

## Downgrade
This chapter covers:
* Downgrade from new Kubernetes that has `CSISkipAttach` enabled to old Kubernetes with `CSISkipAttach disabled.
* Update from Kubernetes that has `CSISkipAttach` feature enabled to the same Kubernetes with `CSISkipAttach` disabled.
* Deletion of CSIDriver instance with non-attachable CSI driver.

In all cases listed above, a non-attachable CSI driver becomes "attachable" (i.e. requires external attacher).  Downgrade does not affect attachable CSI drivers, both "old" and "new" Kubernetes processes them in the same way.

For non-attachable volumes, if the volume was mounted by "new" Kubernetes, it has no VolumeAttachment instance. "Old" A/D controller does not know about it. However, it will periodically call plugin's `VolumesAreAttached()` that checks for `VolumeAttachment` presence. Volumes without `VolumeAttachment` will be reported as not attached and A/D controller will call `Attach()` on these. Since "old" Kubernetes required an external attacher even for non-attachable CSI drivers, the external attacher will pick the `VolumeAttachment` instances and fulfil them in the usual way.


## Performance considerations

* Flow suggested in this proposal adds new `CSIDriver` informer both to A/D controller and kubelet. We don't expect any high amount of instances of `CSIDriver` nor any high frequency of updates. `CSIDriver` should have negligible impact on performance.

* A/D controller will not create `VolumeAttachment` instances for non-attachable volumes. Etcd load will be reduced.

* On the other hand, all CSI volumes still must go though A/D controller. A/D controller **must** process every CSI volume and kubelet **must** wait until A/D controller marks a volume as attached, even if A/D controller basically does nothing. All CSI volumes must be added to `Node.Status.VolumesInUse` and `Node.Status.VolumesAttached`. This does not introduce any new API calls, all this is already implemented, however this proposal won't reduce `Node.Status` update frequency in any way.
  * If *all* volumes move to CSI eventually, pod startup will be slower than when using in-tree volume plugins that don't go through A/D controller and `Node.Status` will grow in size.

## Implementation

Expected timeline:
* Alpha: 1.12 (behind feature gate `CSISkipAttach`)
* Beta: 1.13 (enabled by default)
* GA: 1.14

## Alternatives considered
A/D controller and kubelet can be easily extended to check if a given volume is attachable. This would make mounting of non-attachable volumes easier, as kubelet would not need to wait for A/D controller to mark the volume as attached. However, there would be issues when upgrading or downgrading Kubernetes (or marking CSIDriver as attachable or non-attachable, which has basically the same handling).
* On upgrade (i.e. a previously attachable CSI volume becomes non-attachable, e.g. when user creates CSIDriver instance while corresponding CSI driver is already running), A/D controller could discover that an attached volume is not attachable any longer. A/D controller could clean up `Node.Status.VolumesAttached`, but since A/D controller does not know anything about `VolumeAttachment`, we would either need to introduce a new volume plugin call to clean it up in CSI volume plugin, or something else would need to clean it.
* On downgrade (i.e. a previously non-attachable CSI volume becomes attachable, e.g. when user deletes CSIDriver instance or downgrades to old Kubernetes without this feature), kubelet must discover that already mounted volume has changed from non-attachable to attachable and put it into `Node.Status.VolumesInUse`. This would race with A/D controller detaching the volume when a pod was deleted at the same time a CSIDriver instance was made attachable.

Passing all volumes through A/D controller saves us from these difficulties and even races.
