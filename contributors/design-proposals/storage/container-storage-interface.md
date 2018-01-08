# CSI Volume Plugins in Kubernetes Design Doc

***Status:*** Pending

***Version:*** Alpha

***Author:*** Saad Ali ([@saad-ali](https://github.com/saad-ali), saadali@google.com)

*This document was drafted [here](https://docs.google.com/document/d/10GDyPWbFE5tQunKMlTXbcWysUttMFhBFJRX8ntaS_4Y/edit?usp=sharing).*

## Terminology

Term | Definition
---|---
Container Storage Interface (CSI) | A specification attempting to establish an industry standard interface that Container Orchestration Systems (COs) can use to expose arbitrary storage systems to their containerized workloads.
in-tree | Code that exists in the core Kubernetes repository.
out-of-tree | Code that exists somewhere outside the core Kubernetes repository.
CSI Volume Plugin | A new, in-tree volume plugin that acts as an adapter and enables out-of-tree, third-party CSI volume drivers to be used in Kubernetes.
CSI Volume Driver | An out-of-tree CSI compatible implementation of a volume plugin that can be used in Kubernetes through the Kubernetes CSI Volume Plugin.


## Background & Motivations

Kubernetes volume plugins are currently “in-tree” meaning they are linked, compiled, built, and shipped with the core kubernetes binaries. Adding a new storage system to Kubernetes (a volume plugin) requires checking code into the core Kubernetes code repository. This is undesirable for many reasons including:

1. Volume plugin development is tightly coupled and dependent on Kubernetes releases.
2. Kubernetes developers/community are responsible for testing and maintaining all volume plugins, instead of just testing and maintaining a stable plugin API.
3. Bugs in volume plugins can crash critical Kubernetes components, instead of just the plugin.
4. Volume plugins get full privileges of kubernetes components (kubelet and kube-controller-manager).
5. Plugin developers are forced to make plugin source code available, and can not choose to release just a binary.

The existing [Flex Volume](/contributors/devel/flexvolume.md) plugin attempted to address this by exposing an exec based API for mount/unmount/attach/detach. Although it enables third party storage vendors to write drivers out-of-tree, it requires access to the root filesystem of node and master machines in order to deploy the third party driver files.

Additionally, it doesn’t address another pain of in-tree volumes plugins: dependencies. Volume plugins tend to have many external requirements: dependencies on mount and filesystem tools, for example. These dependencies are assumed to be available on the underlying host OS, which often is not the case, and installing them requires direct machine access. There are efforts underway, for example https://github.com/kubernetes/community/pull/589, that are hoping to address this for in-tree volume plugins. But, enabling volume plugins to be completely containerized will make dependency management much easier.

While Kubernetes has been dealing with these issues, the broader storage community has also been dealing with a fragmented story for how to make their storage system available in different Container Orchestration Systems (COs). Storage vendors have to either write and support multiple volume drivers for different COs or choose to not support some COs.

The Container Storage Interface (CSI) is a specification that resulted from cooperation between community members from various COs--including Kubernetes, Mesos, Cloud Foundry, and Docker. The goal of this interface is to establish a standardized mechanism for COs to expose arbitrary storage systems to their containerized workloads.

The primary motivation for Storage vendors to adopt the interface is a desire to make their system available to as many users as possible with as little work as possible. The primary motivation for COs to adopt the interface is to invest in a mechanism that will enable their users to use as many different storage systems as possible. In addition, for Kubernetes, adopting CSI will have the added benefit of moving volume plugins out of tree, and enabling volume plugins to be containerized.

### Links

* [Container Storage Interface (CSI) Spec](https://github.com/container-storage-interface/spec/blob/master/spec.md)

## Objective

The objective of this document is to document all the requirements for enabling a CSI compliant volume plugin (a CSI volume driver) in Kubernetes.

## Goals

* Define Kubernetes API for interacting with an arbitrary, third-party CSI volume drivers.
* Define mechanism by which Kubernetes master and node components will securely communicate with an arbitrary, third-party CSI volume drivers.
* Define mechanism by which Kubernetes master and node components will discover and register an arbitrary, third-party CSI volume driver deployed on Kubernetes.
* Recommend packaging requirements for Kubernetes compatible, third-party CSI Volume drivers.
* Recommend deployment process for Kubernetes compatible, third-party CSI Volume drivers on a Kubernetes cluster.

## Non-Goals
* Replace [Flex Volume plugin](/contributors/devel/flexvolume.md)
  * The Flex volume plugin exists as an exec based mechanism to create “out-of-tree” volume plugins.
  * Because Flex drivers exist and depend on the Flex interface, it will continue to be supported with a stable API.
  * The CSI Volume plugin will co-exist with Flex volume plugin.

## Design Overview

To support CSI Compliant Volume plugins, a new in-tree CSI Volume plugin will be introduced in Kubernetes. This new volume plugin will be the mechanism by which Kubernetes users (application developers and cluster admins) interact with external CSI volume drivers.

The `SetUp`/`TearDown` calls for the new in-tree CSI volume plugin will directly invoke `NodePublishVolume` and `NodeUnpublishVolume` CSI RPCs through a unix domain socket on the node machine.

Provision/delete and attach/detach must be handled by some external component that monitors the Kubernetes API on behalf of a CSI volume driver and invokes the appropriate CSI RPCs against it.

To simplify integration, the Kubernetes team will offer a containers that captures all the Kubernetes specific logic and act as adapters between third-party containerized CSI volume drivers and Kubernetes (each deployment of a CSI driver would have it’s own instance of the adapter).

## Design Details

### Third-Party CSI Volume Drivers

Kubernetes is as minimally prescriptive on the packaging and deployment of a CSI Volume Driver as possible. Use of the *Communication Channels* (documented below) is the only requirement for enabling an arbitrary external CSI compatible storage driver in Kubernetes.

This document recommends a standard mechanism for deploying an arbitrary containerized CSI driver on Kubernetes. This can be used by a Storage Provider to simplify deployment of containerized CSI compatible volume drivers on Kubernetes (see the “Recommended Mechanism for Deploying CSI Drivers on Kubernetes” section below). This mechanism, however, is strictly optional.

### Communication Channels

#### Kubelet to CSI Driver Communication

Kubelet (responsible for mount and unmount) will communicate with an external “CSI volume driver” running on the same host machine (whether containerized or not) via a Unix Domain Socket.

CSI volume drivers should create a socket at the following path on the node machine: `/var/lib/kubelet/plugins/[SanitizedCSIDriverName]/csi.sock`. For alpha, kubelet will assume this is the location for the Unix Domain Socket to talk to the CSI volume driver. For the beta implementation, we can consider using the [Device Plugin Unix Domain Socket Registration](/contributors/design-proposals/resource-management/device-plugin.md#unix-socket) mechanism to register the Unix Domain Socket with kubelet. This mechanism would need to be extended to support registration of both CSI volume drivers and device plugins independently.

`Sanitized CSIDriverName` is CSI driver name that does not contain dangerous character and can be used as annotation name. It can follow the same pattern that we use for [volume plugins](https://git.k8s.io/kubernetes/pkg/util/strings/escape.go#L27). Too long or too ugly driver names can be rejected, i.e. all components described in this document will report an error and won't talk to this CSI driver. Exact sanitization method is implementation detail (SHA in the worst case).

Upon initialization of the external “CSI volume driver”, some external component must call the CSI method `GetNodeId` to get the mapping from Kubernetes Node names to CSI driver NodeID. It must then add the CSI driver NodeID to the `csi.volume.kubernetes.io/nodeid` annotation on the Kubernetes Node API object. The key of the annotation must be `csi.volume.kubernetes.io/nodeid`. The value of the annotation is a JSON blob, containing key/value pairs for each CSI driver.

For example:
```
csi.volume.kubernetes.io/nodeid: "{ \"driver1\": \"name1\", \"driver2\": \"name2\" }
```

This will enable the component that will issue `ControllerPublishVolume` calls to use the annotation as a mapping from cluster node ID to storage node ID.

To enable easy deployment of an external containerized CSI volume driver, the Kubernetes team will provide a sidecar "Kubernetes CSI Helper" container that can manage the unix domain socket registration and NodeId initialization. This is detailed in the “Suggested Mechanism for Deploying CSI Drivers on Kubernetes” section below.

#### Master to CSI Driver Communication

Because CSI volume driver code is considered untrusted, it might not be allowed to run on the master. Therefore, the Kube controller manager (responsible for create, delete, attach, and detach) can not communicate via a Unix Domain Socket with the “CSI volume driver” container. Instead, the Kube controller manager will communicate with the external “CSI volume driver” through the Kubernetes API.

More specifically, some external component must watch the Kubernetes API on behalf of the external CSI volume driver and trigger the appropriate operations against it. This eliminates the problems of discovery and securing a channel between the kube-controller-manager and the CSI volume driver.

To enable easy deployment of an external containerized CSI volume driver on Kubernetes, without making the driver Kubernetes aware, Kubernetes will provide a sidecar “Kubernetes to CSI” proxy container that will watch the Kubernetes API and trigger the appropriate operations against the “CSI volume driver” container. This is detailed in the “Suggested Mechanism for Deploying CSI Drivers on Kubernetes” section below.

The external component watching the Kubernetes API on behalf of the external CSI volume driver must handle provisioning, deleting, attaching, and detaching.

##### Provisioning and Deleting

Provisioning and deletion operations are handled using the existing [external provisioner mechanism](https://github.com/kubernetes-incubator/external-storage/tree/master/docs), where the external component watching the Kubernetes API on behalf of the external CSI volume driver will act as an external provisioner.

In short, to dynamically provision a new CSI volume, a cluster admin would create a `StorageClass` with the provisioner corresponding to the name of the external provisioner handling provisioning requests on behalf of the CSI volume driver.

To provision a new CSI volume, an end user would create a `PersistentVolumeClaim` object referencing this `StorageClass`. The external provisioner will react to the creation of the PVC and issue the `CreateVolume` call against the CSI volume driver to provision the volume. The `CreateVolume` name will be auto-generated as it is for other dynamically provisioned volumes. The `CreateVolume` capacity will be taken from the `PersistentVolumeClaim` object. The `CreateVolume` parameters will be passed through from the `StorageClass` parameters (opaque to Kubernetes). Once the operation completes successfully, the external provisioner creates a `PersistentVolume` object to represent the volume using the information returned in the `CreateVolume` response. The `PersistentVolume` object is bound to the `PersistentVolumeClaim` and available for use.

To delete a CSI volume, an end user would delete the corresponding `PersistentVolumeClaim` object. The external provisioner will react to the deletion of the PVC and based on its reclamation policy it will issue the `DeleteVolume` call against the CSI volume driver commands to delete the volume. It will then delete the `PersistentVolume` object.

##### Attaching and Detaching

Attach/detach operations must also be handled by an external component (an “attacher”). The attacher watches the Kubernetes API on behalf of the external CSI volume driver for new `VolumeAttachment` objects (defined below), and triggers the appropriate calls against the CSI volume driver to attach the volume. The attacher must watch for `VolumeAttachment` object and mark it as attached even if the underlying CSI driver does not support `ControllerPublishVolume` call, as Kubernetes has no knowledge about it.

More specifically, an external “attacher” must watch the Kubernetes API on behalf of the external CSI volume driver to handle attach/detach requests.

Once the following conditions are true, the external-attacher should call `ControllerPublishVolume` against the CSI volume driver to attach the volume to the specified node:

1. A new `VolumeAttachment` Kubernetes API objects is created by Kubernetes attach/detach controller.
2. The `VolumeAttachment.Spec.Attacher` value in that object corresponds to the name of the external attacher.
3. The `VolumeAttachment.Status.Attached` value is not yet set to true.
4. A Kubernetes Node API object exists with the name matching `VolumeAttachment.Spec.NodeName` and that object contains a `csi.volume.kubernetes.io/nodeid` annotation. This annotation contains a JSON blob, a list of key/value pairs, where one of they keys corresponds with the CSI volume driver name, and the value is the NodeID for that driver. This NodeId mapping can be retrieved and used in the `ControllerPublishVolume` calls.
5. The `VolumeAttachment.Metadata.DeletionTimestamp` is not set.

Before starting the `ControllerPublishVolume` operation, the external-attacher should add these finalizers to these Kubernetes API objects:

* To the `VolumeAttachment` so that when the object is deleted, the external-attacher has an opportunity to detach the volume first. External attacher removes this finalizer once the volume is fully detached from the node.
* To the `PersistentVolume` referenced by `VolumeAttachment` so the the PV cannot be deleted while the volume is attached. External attacher needs information from the PV to perform detach operation. The attacher will remove the finalizer once all `VolumeAttachment` objects that refer to the PV are deleted, i.e. the volume is detached from all nodes.

If the operation completes successfully, the external-attacher will:

1. Set `VolumeAttachment.Status.Attached` field to true to indicate the volume is attached.
2. Update the `VolumeAttachment.Status.AttachmentMetadata` field with the contents of the returned `PublishVolumeInfo`.
3. Clear the `VolumeAttachment.Status.AttachError` field.

If the operation fails, the external-attacher will:

1. Ensure the `VolumeAttachment.Status.Attached` field to still false to indicate the volume is not attached.
2. Set the `VolumeAttachment.Status.AttachError` field detailing the error.
3. Create an event against the Kubernetes API associated with the `VolumeAttachment` object to inform users what went wrong.

The external-attacher may implement it’s own error recovery strategy, and retry as long as conditions specified for attachment above are valid. It is strongly recommended that the external-attacher implement an exponential backoff strategy for retries.

The detach operation will be triggered by the deletion of the `VolumeAttachment` Kubernetes API objects. Since the `VolumeAttachment` Kubernetes API object will have a finalizer added by the external-attacher, it will wait for confirmation from the external-attacher before deleting the object.

Once all the following conditions are true, the external-attacher should call `ControllerUnpublishVolume` against the CSI volume driver to detach the volume from the specified node:
1. A `VolumeAttachment` Kubernetes API object is marked for deletion: the value for the `VolumeAttachment.metadata.deletionTimestamp` field is set.

If the operation completes successfully, the external-attacher will:
1. Remove its finalizer from the list of finalizers on the `VolumeAttachment` object permitting the delete operation to continue.

If the operation fails, the external-attacher will:

1. Ensure the `VolumeAttachment.Status.Attached` field remains true to indicate the volume is not yet detached.
2. Set the `VolumeAttachment.Status.DetachError` field detailing the error.
3. Create an event against the Kubernetes API associated with the `VolumeAttachment` object to inform users what went wrong.

The new API object called `VolumeAttachment` will be defined as follows:

```GO

// VolumeAttachment captures the intent to attach or detach the specified volume
// to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachment struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired attach/detach volume behavior.
	// Populated by the Kubernetes system.
	Spec VolumeAttachmentSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status of the VolumeAttachment request.
	// Populated by the entity completing the attach or detach
	// operation, i.e. the external-attacher.
	// +optional
	Status VolumeAttachmentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// The specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this
	// request. This is the name returned by GetPluginName() and must be the
	// same as StorageClass.Provisioner.
	Attacher string `json:"attacher" protobuf:"bytes,1,opt,name=attacher"`

	// AttachedVolumeSource represents the volume that should be attached.
	VolumeSource AttachedVolumeSource `json:"volumeSource" protobuf:"bytes,2,opt,name=volumeSource"`

	// Kubernetes node name that the volume should be attached to.
	NodeName string `json:"nodeName" protobuf:"bytes,3,opt,name=nodeName"`
}

// VolumeAttachmentSource represents a volume that should be attached.
// Right now only PersistentVolumes can be attached via external attacher,
// in future we may allow also inline volumes in pods.
// Exactly one member can be set.
type AttachedVolumeSource struct {
	// Name of the persistent volume to attach.
	// +optional
	PersistentVolumeName *string `json:"persistentVolumeName,omitempty" protobuf:"bytes,1,opt,name=persistentVolumeName"`

	// Placeholder for *VolumeSource to accommodate inline volumes in pods.
}

// The status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// Indicates the volume is successfully attached.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	Attached bool `json:"attached" protobuf:"varint,1,opt,name=attached"`

	// Upon successful attach, this field is populated with any
	// information returned by the attach operation that must be passed
	// into subsequent WaitForAttach or Mount calls.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
	AttachmentMetadata map[string]string `json:"attachmentMetadata,omitempty" protobuf:"bytes,2,rep,name=attachmentMetadata"`

	// The most recent error encountered during attach operation, if any.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
    AttachError *VolumeError `json:"attachError,omitempty" protobuf:"bytes,3,opt,name=attachError,casttype=VolumeError"`

	// The most recent error encountered during detach operation, if any.
	// This field must only be set by the entity completing the detach
	// operation, i.e. the external-attacher.
	// +optional
	DetachError *VolumeError `json:"detachError,omitempty" protobuf:"bytes,4,opt,name=detachError,casttype=VolumeError"`
}

// Captures an error encountered during a volume operation.
type VolumeError struct {
	// Time the error was encountered.
	// +optional
	Time metav1.Time `json:"time,omitempty" protobuf:"bytes,1,opt,name=time"`

	// String detailing the error encountered during Attach or Detach operation.
	// This string may be logged, so it should not contain sensitive
	// information.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

```

### Kubernetes In-Tree CSI Volume Plugin

A new in-tree Kubernetes CSI Volume plugin will contain all the logic required for Kubernetes to communicate with an arbitrary, out-of-tree, third-party CSI compatible volume driver.

The existing Kubernetes volume components (attach/detach controller, PVC/PV controller, Kubelet volume manager) will handle the lifecycle of the CSI volume plugin operations (everything from triggering volume provisioning/deleting, attaching/detaching, and mounting/unmounting) just as they do for existing in-tree volume plugins.

#### Proposed API

A new `CSIPersistentVolumeSource` object will be added to the Kubernetes API. It will be part of the existing `PersistentVolumeSource` object and thus can be used only via PersistentVolumes. CSI volumes will not be allow referencing directly from Pods without a `PersistentVolumeClaim`.

```GO
type CSIPersistentVolumeSource struct {
  // Driver is the name of the driver to use for this volume.
  // Required.
  Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver"`

  // VolumeHandle is the unique volume name returned by the CSI volume
  // plugin’s CreateVolume to refer to the volume on all subsequent calls.
  VolumeHandle string `json:"volumeHandle" protobuf:"bytes,2,opt,name=volumeHandle"`

  // Optional: The value to pass to ControllerPublishVolumeRequest.
  // Defaults to false (read/write).
  // +optional
  ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,5,opt,name=readOnly"`
}
```

#### Internal Interfaces

The in-tree CSI volume plugin will implement the following internal Kubernetes volume interfaces:

1. `VolumePlugin`
    * Mounting/Unmounting of a volume to a specific path.
2. `AttachableVolumePlugin`
    * Attach/detach of a volume to a given node.

Notably, `ProvisionableVolumePlugin` and `DeletableVolumePlugin` are not implemented because provisioning and deleting for CSI volumes is handled by an external provisioner.

#### Mount and Unmount

The in-tree volume plugin’s SetUp and TearDown methods will trigger the `NodePublishVolume` and `NodeUnpublishVolume` CSI calls via Unix Domain Socket. Kubernetes will generate a unique `target_path` (unique per pod per volume) to pass via `NodePublishVolume` for the CSI plugin to mount the volume. Upon successful completion of the `NodeUnpublishVolume` call (once volume unmount has been verified), Kubernetes will delete the directory.

The Kubernetes volume sub-system does not currently support block volumes (only file), so for alpha, the Kubernetes CSI volume plugin will only support file.

#### Attaching and Detaching

The attach/detach controller,running as part of the kube-controller-manager binary on the master, decides when a CSI volume must be attached or detached from a particular node.

When the controller decides to attach a CSI volume, it will call the in-tree CSI volume plugin’s attach method. The in-tree CSI volume plugin’s attach method will do the following:

1. Create a new `VolumeAttachment` object (defined in the “Communication Channels” section) to attach the volume.
    * The name of the of the `VolumeAttachment` object will be `pv-<SHA256(PVName+NodeName)>`.
      * `pv-` prefix is used to allow using other scheme(s) for inline volumes in the future, with their own prefix.
      * SHA256 hash is to reduce length of `PVName` plus `NodeName` string, each of which could be max allowed name length (hexadecimal representation of SHA256 is 64 characters).
      * `PVName` is `PV.name` of the attached PersistentVolume.
      * `NodeName` is `Node.name` of the node where the volume should be attached to.
    * If a `VolumeAttachment` object with the corresponding name already exists, the in-tree volume plugin will simply begin to poll it as defined below. The object is not modified; only the external-attacher should change the status fields; and the external-attacher is responsible for it’s own retry and error handling logic.
2. Poll the `VolumeAttachment` object waiting for one of the following conditions:
    * The `VolumeAttachment.Status.Attached` field to become `true`.
      * The operation completes successfully.
    * An error to be set in the `VolumeAttachment.Status.AttachError` field.
      * The operation terminates with the specified error.
    * The operation to timeout.
      * The operation terminates with timeout error.
    * The `VolumeAttachment.DeletionTimestamp` is set.
      * The operation terminates with an error indicating a detach operation is in progress.
      * The `VolumeAttachment.Status.Attached` value must not be trusted. The attach/detach controller has to wait until the object is deleted by the external-attacher before creating a new instance of the object.

When the controller decides to detach a CSI volume, it will call the in-tree CSI volume plugin’s detach method. The in-tree CSI volume plugin’s detach method will do the following:

1. Delete the corresponding `VolumeAttachment` object (defined in the “Communication Channels” section) to indicate the volume should be detached.
2. Poll the `VolumeAttachment` object waiting for one of the following conditions:
    * The `VolumeAttachment.Status.Attached` field to become false.
      * The operation completes successfully.
    * An error to be set in the `VolumeAttachment.Status.DetachError` field.
      * The operation terminates with the specified error.
    * The object to no longer exists.
      * The operation completes successfully.
    * The operation to timeout.
      * The operation terminates with timeout error.

### Recommended Mechanism for Deploying CSI Drivers on Kubernetes

Although, Kubernetes does not dictate the packaging for a CSI volume driver, it offers the following recommendations to simplify deployment of a containerized CSI volume driver on Kubernetes.

![Recommended CSI Deployment Diagram](container-storage-interface_diagram1.png?raw=true "Recommended CSI Deployment Diagram")

To deploy a containerized third-party CSI volume driver, it is recommended that storage vendors:

  * Create a “CSI volume driver” container that implements the volume plugin behavior and exposes a gRPC interface via a unix domain socket, as defined in the CSI spec (including Controller, Node, and Identity services).
  * Bundle the “CSI volume driver” container with helper containers (external-attacher, external-provisioner, Kubernetes CSI Helper) that the Kubernetes team will provide (these helper containers will assist the “CSI volume driver” container in interacting with the Kubernetes system). More specifically, create the following Kubernetes objects:
    * A `StatefulSet` (to facilitate communication with the Kubernetes controllers) that has:
      * Replica size 1
        * Guarantees that no more than 1 instance of the pod will be running at once (so we don’t have to worry about multiple instances of the `external-provisioner` or `external-attacher` in the cluster).
      * The following containers
        * The “CSI volume driver” container created by the storage vendor.
        * The `external-attacher` container provided by the Kubernetes team.
        * The `external-provisioner` container provided by the Kubernetes team.
      * The following volumes:
        * `emptyDir` volume
          * Mounted inside all containers at `/var/lib/csi/sockets/pluginproxy/`
          * The “CSI volume driver” container should create its Unix Domain Socket in this directory to enable communication with the Kubernetes helper container(s) (`external-provisioner`, `external-attacher`).
    * A `DaemonSet` (to facilitate communication with every instance of kubelet) that has:
      * The following containers
        * The “CSI volume driver” container created by the storage vendor.
        * The “Kubernetes CSI Helper” container provided by the Kubernetes team
          * Responsible for registering the unix domain socket with kubelet and initializing NodeId.
      * The following volumes:
        * `hostpath` volume
          * Expose `/var/lib/kubelet/device-plugins/kubelet.sock` from the host.
          * Mount only in “Kubernetes CSI Helper” container at `/var/lib/csi/sockets/kubelet.sock`
          * The Kubernetes to CSI proxy container will use this unix domain socket to register the CSI driver’s unix domain socket with kubelet.
        * `hostpath` volume
          * Expose `/var/lib/kubelet/` from the host.
          * Mount only in “CSI volume driver” container at `/var/lib/kubelet/`
          * Ensure [bi-directional mount propagation](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation) is enabled, so that any mounts setup inside this container are propagated back to the host machine.
        * `hostpath` volume
          * Expose `/var/lib/kubelet/plugins/[SanitizedCSIDriverName]/` from the host as `hostPath.type = "DirectoryOrCreate"`.
          * Mount inside “CSI volume driver” container at the path the CSI gRPC socket will be created.
          * This is the primary means of communication between Kubelet and the “CSI volume driver” container (gRPC over UDS).
  * Have cluster admins deploy the above `StatefulSet` and `DaemonSet` to add support for the storage system in their Kubernetes cluster.

Alternatively, deployment could be simplified by having all components (including external-provisioner and external-attacher) in the same pod (DaemonSet). Doing so, however, would consume more resources, and require a leader election protocol (likely https://git.k8s.io/contrib/election) in the `external-provisioner` and `external-attacher` components.

### Example Walkthrough

#### Provisioning Volumes

1. A cluster admin creates a `StorageClass` pointing to the CSI driver’s external-provisioner and specifying any parameters required by the driver.
2. A user creates a `PersistentVolumeClaim` referring to the new `StorageClass`.
3. The persistent volume controller realizes that dynamic provisioning is needed, and marks the PVC with a `volume.beta.kubernetes.io/storage-provisioner` annotation.
4. The external-provisioner for the CSI driver sees the `PersistentVolumeClaim` with the `volume.beta.kubernetes.io/storage-provisioner` annotation so it starts dynamic volume provisioning:
    1. It dereferences the `StorageClass` to collect the opaque parameters to use for provisioning.
    2. It calls `CreateVolume` against the CSI driver container with parameters from the `StorageClass` and `PersistentVolumeClaim` objects.
5. Once the volume is successfully created, the external-provisioner creates a `PersistentVolume` object to represent the newly created volume and binds it to the `PersistentVolumeClaim`.

#### Deleting Volumes

1. A user deletes a `PersistentVolumeClaim` object bound to a CSI volume.
2. The external-provisioner for the CSI driver sees the the `PersistentVolumeClaim` was deleted and triggers the retention policy:
  1. If the retention policy is `delete`
    1. The external-provisioner triggers volume deletion by issuing a `DeleteVolume` call against the CSI volume plugin container.
    2. Once the volume is successfully deleted, the external-provisioner deletes the corresponding `PersistentVolume` object.
  2. If the retention policy is `retain`
    1. The external-provisioner does not delete the `PersistentVolume` object.

#### Attaching Volumes

1. The Kubernetes attach/detach controller, running as part of the `kube-controller-manager` binary on the master, sees that a pod referencing a CSI volume plugin is scheduled to a node, so it calls the in-tree CSI volume plugin’s attach method.
2. The in-tree volume plugin creates a new `VolumeAttachment` object in the kubernetes API and waits for its status to change to completed or error.
3. The external-attacher sees the `VolumeAttachment` object and triggers a `ControllerPublish` against the CSI volume driver container to fulfil it (meaning the external-attacher container issues a gRPC call via underlying UNIX domain socket to the CSI driver container).
4. Upon successful completion of the `ControllerPublish` call the external-attacher updates the status of the `VolumeAttachment` object to indicate the volume is successfully attached.
5. The in-tree volume plugin watching the status of the `VolumeAttachment` object in the Kubernetes API, sees the `Attached` field set to true indicating the volume is attached, so it updates the attach/detach controller’s internal state to indicate the volume is attached.

#### Detaching Volumes

1. The Kubernetes attach/detach controller, running as part of the `kube-controller-manager` binary on the master, sees that a pod referencing an attached CSI volume plugin is terminated or deleted, so it calls the in-tree CSI volume plugin’s detach method.
2. The in-tree volume plugin deletes the corresponding `VolumeAttachment` object.
3. The external-attacher sees a `deletionTimestamp` set on the `VolumeAttachment` object and triggers a `ControllerUnpublish` against the CSI volume driver container to detach it.
4. Upon successful completion of the `ControllerUnpublish` call, the external-attacher removes the finalizer from the `VolumeAttachment` object to indicate successful completion of the detach operation allowing the `VolumeAttachment` object to be deleted.
5. The in-tree volume plugin waiting for the `VolumeAttachment` object sees it deleted and assumes the volume was successfully detached, so It updates the attach/detach controller’s internal state to indicate the volume is detached.

#### Mounting Volumes

1. The volume manager component of kubelet notices a new volume, referencing a CSI volume, has been scheduled to the node, so it calls the in-tree CSI volume plugin’s `WaitForAttach` method.
2. The in-tree volume plugin’s `WaitForAttach` method watches the `Attached` field of the `VolumeAttachment` object in the kubernetes API to become `true`, it then returns without error.
3. Kubelet then calls the in-tree CSI volume plugin’s `MountDevice` method which is a no-op and returns immediately.
4. Finally kubelet calls the in-tree CSI volume plugin’s mount (setup) method, which causes the in-tree volume plugin to issue a `NodePublishVolume` call via the registered unix domain socket to the local CSI driver.
5. Upon successful completion of the `NodePublishVolume` call the specified path is mounted into the pod container.

#### Unmounting Volumes
1. The volume manager component of kubelet, notices a mounted CSI volume, referenced by a pod that has been deleted or terminated, so it calls the in-tree CSI volume plugin’s `UnmountDevice` method which is a no-op and returns immediately.
2. Next kubelet calls the in-tree CSI volume plugin’s unmount (teardown) method, which causes the in-tree volume plugin to issue a `NodeUnpublishVolume` call via the registered unix domain socket to the local CSI driver. If this call fails from any reason, kubelet re-tries the call periodically.
3. Upon successful completion of the `NodeUnpublishVolume` call the specified path is unmounted from the pod container.


### CSI Credentials

This part of proposal is not going to be implemented in alpha release.

#### End user credentials
CSI allows specifying *end user credentials* in all operations. Kubernetes does not have facility to configure a Secret per *user*, we usually track objects per *namespace*. Therefore we decided to postpone implementation of these credentials and wait until CSI is clarified.

#### Volume specific credentials
Some storage technologies (e.g. iSCSI with CHAP) require credentials tied to the volume (iSCSI LUN) that must be used during `NodePublish` request. It is expected that these credentials will be provided during dynamic provisioning of the volume, however CSI `CreateVolume` response does not provide any. In case it gets fixed soon external provisioner can save the secrets in a dedicated namespace and make them available to external attacher and internal CSI volume plugin using these `CSIPersistentVolumeSource` fields:

// ...
```go
type CSIPersistentVolumeSource struct {

    // Optional: MountSecretRef is a reference to the secret object containing
    // sensitive information to pass to the CSI driver during NodePublish.
    // This may be empty if no secret is required. If the secret object contains
    // more than one secret, all secrets are passed.
    // +optional
    MountSecretRef *SecretReference `json:"mountSecretRef,omitempty" protobuf:"bytes,3,opt,name=mountSecretRef"`

    // Optional: AttachSecretRef is a reference to the secret object containing
    // sensitive information to pass to the CSI driver during ControllerPublish.
    // This may be empty if no secret is required. If the secret object contains
    // more than one secret, all secrets are passed.
    // +optional
    AttachSecretRef *SecretReference `json:"attachSecretRef,omitempty" protobuf:"bytes,4,opt,name=attachSecretRef"`
}
```

Note that a malicious provisioner could obtain an arbitrary secret by setting the mount secret in PV object to whatever secret it wants. It is assumed that cluster admins will only run trusted provisioners.

Because the kubelet would be responsible for fetching and passing the mount secret to the CSI driver,the Kubernetes NodeAuthorizer must be updated to allow kubelet read access to mount secrets.

## Alternatives Considered

### Extending PersistentVolume Object

Instead of creating a new `VolumeAttachment` object, another option we considered was extending the existing `PersistentVolume` object.

`PersistentVolumeSpec` would be extended to include:
* List of nodes to attach the volume to (initially empty).

`PersistentVolumeStatus` would be extended to include:
* List of nodes the volume was successfully attached to.

We dismissed this approach because having attach/detach triggered by the creation/deletion of an object is much easier to manage (for both external-attacher and Kubernetes) and more robust (fewer corner cases to worry about).
