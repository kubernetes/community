# In-line CSI volumes in Pods

Author: @jsafrane

## Goal
* Define API and high level design for in-line CSI volumes in Pod.
* Make in-line CSI volumes secure for using ephemeral volumes (such as Secrets or ConfigMap).

## Motivation
Currently, CSI can be used only through PersistentVolume object. All other persistent volume sources support in-line volumes in Pods, CSI should be no exception. There are three main drivers:
* We want to move away from in-tree volume plugins to CSI, as designed in a separate proposal https://github.com/kubernetes/community/pull/2199/. In-line volumes should use CSI too.
* CSI drivers can be used to provide ephemeral volumes used to inject state, configuration, secrets, identity or similar information to pods, like Secrets and ConfigMap in-tree volumes do today. We don't want to force users to create PVs for each such volume, we should allow to use them in-line in pods as regular Secrets or ephemeral Flex volumes.
* Get the same features as Flex and deprecate Flex. (I.e. replace it with some CSI-Flex bridge. This bridge is out of scope of this proposal.)

## API
`VolumeSource` needs to be extended with CSI volume source:
```go
type VolumeSource struct {
    // <snip>

	// CSI (Container Storage Interface) represents storage that handled by an external CSI driver (Beta feature).
	// +optional
	CSI *CSIVolumeSource
}


// Represents storage that is managed by an external CSI volume driver (Alpha feature)
type CSIVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	// Required.
	Driver string

	// VolumeHandle is the unique ID of the volume. It is the volume ID used in
	// all CSI calls, optionally with a prefix based on VolumeHandlePrefix
	// value.
	// Required
	VolumeHandle string

	// VolumeHandlePrefix is type of prefix added to VolumeHandle before using
	// it as CSI volume ID. It ensures that volumes with the same VolumeHandle
	// in different pods or namespaces get unique CSI volume ID.
	// Required.
	VolumeHandlePrefix CSIVolumeHandlePrefix

	// Optional: The value to pass to ControllerPublishVolumeRequest.
	// Defaults to false (read/write).
	// +optional
	ReadOnly bool

	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string

	// Attributes of the volume. This corresponds to "volume_attributes" in some
	// CSI calls.
	// +optional
	VolumeAttributes map[string]string

	// ControllerPublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// ControllerPublishVolume and ControllerUnpublishVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	ControllerPublishSecretRef *LocalObjectReference

	// NodeStageSecretRef is a reference to the secret object containing sensitive
	// information to pass to the CSI driver to complete the CSI NodeStageVolume
	// and NodeStageVolume and NodeUnstageVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodeStageSecretRef *LocalObjectReference

	// NodePublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// NodePublishVolume and NodeUnpublishVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodePublishSecretRef *LocalObjectReference
}

type CSIVolumeHandlePrefix string
const (
	// VolumeHandle is prefixed by Pod UID.
	CSIVolumeHandlePrefixPod CSIVolumeHandlePrefix  = "Pod"
	// VolumeHandle is prefixed by UID of the namespace where the pod is located.
	CSIVolumeHandlePrefixNamespace CSIVolumeHandlePrefix  = "Namespace"
	// VolumeHandle is not modified.
	CSIVolumeHandlePrefixNone CSIVolumeHandlePrefix = "None"
)
```

The difference between `CSIVolumeSource` (in-lined in a pod) and `CSIPersistentVolumeSource` (in PV) are:

* All secret references in in-line volumes can refer only to secrets in the same namespace where the corresponding pod is running. This is common in all other volume sources that refer to secrets, incl. Flex.
* VolumeHandle in in-line volumes can have a prefix. This prefix (Pod UID, Namespace UID or nothing) is added to the VolumeHandle before each CSI call. It makes sure that each pod uses a different volume ID for its ephemeral volumes.  The prefix must be explicitly set by pod author, there is no default.
	* Users don't need to think about VolumeHandles used in other pods in their namespace, as each pod will get an unique prefix when `CSIVolumeHandlePrefixPod` is used. CSI volume ID with this prefix cannot accidentally conflict by another volume ID in another pod.
	* Each pod created by ReplicaSet, StatefulSet or DaemonSet will get the same copy of a pod template. `CSIVolumeHandlePrefixPod` makes sure that each pod gets its own unique volume ID and thus can get its own volume instance.
	* Without the prefix, user could guess volume ID of a secret-like CSI volume of another user and craft a pod with in-line volume referencing it. CSI driver, obeying idempotency, must then give the same volume to this pod. If users can use only`CSIVolumeHandlePrefixNamespace` or `CSIVolumeHandlePrefixPod`in their in-line volumes, we can make sure that they can't steal secrets of each other.
		* `PodSecurityPolicy` will be extended to allow / deny users using in-line volumes with no prefix.
	* Finally, `CSIVolumeHandlePrefixNone` allows selected users (based on PSP) to use persistent storage volumes in-line in pods.

## Implementation
#### Provisioning/Deletion
N/A, it works only with PVs and not with in-line volumes.

### Attach/Detach
Current `storage.VolumeAttachment` object contains only reference to PV that's being attached. It must be extended with VolumeSource for in-line volumes in pods.

```go
// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
    // <snip>

	// Source represents the volume that should be attached.
	Source VolumeAttachmentSource
}

// VolumeAttachmentSource represents a volume that should be attached, either
// PersistentVolume or a volume in-lined in a Pod.
// Exactly one member can be set.
type VolumeAttachmentSource struct {
	// Name of the persistent volume to attach.
	// +optional
	PersistentVolumeName *string

	// InlineVolumeSource represents the source location of a in-line volume in a pod to attach.
	// +optional
    InlineVolumeSource *InlineVolumeSource
}

// InlineVolumeSource represents the source location of a in-line volume in a pod.
type InlineVolumeSource struct {
	// VolumeSource is copied from the pod. It ensures that attacher has enough
	// information to detach a volume when the pod is deleted before detaching.
	// Only CSIVolumeSource can be set.
	// Required.
	VolumeSource v1.VolumeSource

	// Namespace of the pod with in-line volume. It is used to resolve
	// references to Secrets in VolumeSource.
	// Required.
	Namespace string
}
```

* A/D controller **copies whole `VolumeSource`**  from `Pod` into `VolumeAttachment`. This allows external CSI attacher to detach volumes for deleted pods without keeping any internal database of attached VolumeSources.
* Using whole `VolumeSource` allows us to re-use `VolumeAttachment` for any other in-line volume in the future. We provide validation that this `VolumeSource` contains only `CSIVolumeSource` to clearly state that only CSI is supported now.
* External CSI attacher must be extended to  process either `PersistentVolumeName` or `VolumeSource`.
* Since in-line volume in a pod can refer to a secret in the same namespace as the pod, **external attacher may need permissions to read any Secrets in any namespace**.
* CSI `ControllerUnpublishVolume` call (~ volume detach) requires the Secrets to be available at detach time. Current CSI attacher implementation simply expects that the Secrets are available at detach time.
* Secrets for PVs are "global", out of user's namespace, so this assumption is probably OK.
* Secrets for in-line volumes must be in the same namespace as the pod that contains the volume. Users can delete them before the volume is detached. We deliberately choose to let the external attacher to fail when such Secret cannot be found on detach time and keep the volume attached, reporting errors about missing Secrets to user.
	* Since access to in-line volumes can be configured by `PodSecurityPolicy` (see below), we expect that cluster admin gives access to CSI drivers that require secrets at detach time only to educated users that know they should not delete Secrets used in volumes.
	* Number of CSI drivers that require Secrets on detach is probably very limited. No in-tree Kubernetes volume plugin requires them on detach.
	* We will provide clear documentation that using in-line volumes with drivers that require credentials on detach may leave orphaned attached volumes that Kubernetes is not able to detach. It's up to the cluster admin to decide if using such CSI driver is worth it.

### Kubelet (MountDevice/SetUp/TearDown/UnmountDevice)
In-tree CSI volume plugin calls in kubelet get universal `volume.Spec`, which contains either `v1.VolumeSource` from Pod (for in-line volumes) or `v1.PersistentVolume`. We need to modify CSI volume plugin to check for presence of `VolumeSource` or `PersistentVolume` and read NodeStage/NodePublish secrets from appropriate source. Kubelet does not need any new permissions, it already can read secrets for pods that it handles. These secrets are needed only for `MountDevice/SetUp` calls and don't need to be cached until `TearDown`/`UnmountDevice`.

### `PodSecurityPolicy`
* `PodSecurityPolicy` must be enhanced to limit pods in using in-line CSI volumes. It will be modeled following existing Flex volume policy. There is no default, users can't use in-line CSI volumes unless some CSI drivers are explicitly allowed.
  ```go
  type PodSecurityPolicySpec struct {
	// <snip>

	// AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all
	// Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes
	// is allowed in the "Volumes" field.
	// +optional
	AllowedFlexVolumes []AllowedFlexVolume

	// AllowedCSIVolumes is a whitelist of allowed CSI volumes.  Empty or nil indicates that all
	// CSI volumes may be used.  This parameter is effective only when the usage of the CSI volumes
	// is allowed in the "Volumes" field.
	// +optional
	AllowedCSIVolumes []AllowedCSIVolume
  }

  // AllowedCSIVolume represents a single CSI volume that is allowed to be used.
  type AllowedCSIVolume struct {
	// Driver is the name of the CSI volume driver.
	Driver string
  }
  ```
* `PodSecurityPolicy` must be extended to allow users to use in-line volumes with no prefixes. This prevents users from stealing data from Secrets-like ephemeral volumes inlined in pods by guessing volume ID of someone else. There is no default, users can't use in-line CSI volumes unless some prefixes are explicitly allowed.
    ```go
    type PodSecurityPolicySpec struct {
	  // <snip>
	  // AllowedCSIVolumeHandlePrefixes is a whitelist of volume prefixes
	  // allowed to be used in CSI volumes in-lined in pods.
	  AllowedCSIVolumeHandlePrefixes []core.CSIVolumeHandlePrefix
    }
    ```

* `PodSecurityPolicy` must be extended to allow users to use attachable in-line CSI volumes. This prevents users from leaving orphaned attached volumes when they delete Secrets required to detach volumes. **Kubernetes currently does not know which CSI volumes are attachable or not. There are several options considered and it will be handled in a separate proposal.**
   ```go
   type PodSecurityPolicySpec struct {
	  // <snip>
	  // AllowAttachableCSIVolumes allows users to use attachable CSI volumes
	  // in-line in pod definitions.
	  AllowAttachableCSIVolumes bool
    }
    ```

### Security considerations
As written above, external attacher may requrie permissions to read Secrets in any namespace. It is up to CSI driver author to document if the driver needs such permission (i.e. access to Secrets at attach/detach time) and up to cluster admin to deploy the driver with these permissions or restrict external attacher to access secrets only in some namespaces.
