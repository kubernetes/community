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

The existing [Flex Volume] plugin attempted to address this by exposing an exec based API for mount/unmount/attach/detach. Although it enables third party storage vendors to write drivers out-of-tree, it requires access to the root filesystem of node and master machines in order to deploy the third party driver files.

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
* Replace [Flex Volume plugin]
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

Upon initialization of the external “CSI volume driver”, kubelet must call the CSI method `NodeGetInfo` to get the mapping from Kubernetes Node names to CSI driver NodeID and the associated `accessible_topology`. It must:

  * Create/update a `CSINodeInfo` object instance for the node with the NodeID and topology keys from `accessible_topology`.
    * This will enable the component that will issue `ControllerPublishVolume` calls to use the `CSINodeInfo` as a mapping from cluster node ID to storage node ID.
    * This will enable the component that will issue `CreateVolume` to reconstruct `accessible_topology` and provision a volume that is accesible from specific node.
    * Each driver must completely overwrite its previous version of NodeID and topology keys, if they exist.
    * If the `NodeGetInfo` call fails, kubelet must delete any previous NodeID and topology keys for this driver.
    * When kubelet plugin unregistration mechanism is implemented, delete NodeID and topology keys when a driver is unregistered.

  * Update Node API object with the CSI driver NodeID as the `csi.volume.kubernetes.io/nodeid` annotation. The value of the annotation is a JSON blob, containing key/value pairs for each CSI driver. For example:
    ```
    csi.volume.kubernetes.io/nodeid: "{ \"driver1\": \"name1\", \"driver2\": \"name2\" }
    ```

    *This annotation is deprecated and will be removed according to deprecation policy (1 year after deprecation). TODO mark deprecation date.*
    * If the `NodeGetInfo` call fails, kubelet must delete any previous NodeID for this driver.
    * When kubelet plugin unregistration mechanism is implemented, delete NodeID and topology keys when a driver is unregistered.

  * Create/update Node API object with `accessible_topology` as labels.
    There are no hard restrictions on the label format, but for the format to be used by the recommended setup, please refer to [Topology Representation in Node Objects](#topology-representation-in-node-objects).

To enable easy deployment of an external containerized CSI volume driver, the Kubernetes team will provide a sidecar "Kubernetes CSI Helper" container that can manage the unix domain socket registration and NodeId initialization. This is detailed in the “Suggested Mechanism for Deploying CSI Drivers on Kubernetes” section below.

The new API object called `CSINodeInfo` will be defined as follows:

```go
// CSINodeInfo holds information about status of all CSI drivers installed on a node.
type CSINodeInfo struct {
    metav1.TypeMeta
    // ObjectMeta.Name must be node name.
    metav1.ObjectMeta

    // List of CSI drivers running on the node and their properties.
    CSIDrivers []CSIDriverInfo
}

// Information about one CSI driver installed on a node.
type CSIDriverInfo struct {
    // CSI driver name.
    Name string

    // ID of the node from the driver point of view.
    NodeID string

    // Topology keys reported by the driver on the node.
    TopologyKeys []string
}
```

A new object type `CSINodeInfo` is chosen instead of `Node.Status` field because Node is already big enough and there are issues with its size. `CSINodeInfo` is CRD installed by TODO (jsafrane) on cluster startup and defined in `kubernetes/kubernetes/pkg/apis/storage-csi/v1alpha1/types.go`, so k8s.io/client-go and k8s.io/api are generated automatically. All users of `CSINodeInfo` will tolerate if the CRD is not installed and retry anything they need to do with it with exponential backoff and proper error reporting. Especially kubelet is able to serve its usual duties when the CRD is missing.

Each node must have zero or one `CSINodeInfo` instance. This is ensured by `CSINodeInfo.Name == Node.Name`. TODO: how to validate this? Each `CSINodeInfo` is "owned" by corresponding Node for garbage collection.


#### Master to CSI Driver Communication

Because CSI volume driver code is considered untrusted, it might not be allowed to run on the master. Therefore, the Kube controller manager (responsible for create, delete, attach, and detach) can not communicate via a Unix Domain Socket with the “CSI volume driver” container. Instead, the Kube controller manager will communicate with the external “CSI volume driver” through the Kubernetes API.

More specifically, some external component must watch the Kubernetes API on behalf of the external CSI volume driver and trigger the appropriate operations against it. This eliminates the problems of discovery and securing a channel between the kube-controller-manager and the CSI volume driver.

To enable easy deployment of an external containerized CSI volume driver on Kubernetes, without making the driver Kubernetes aware, Kubernetes will provide a sidecar “Kubernetes to CSI” proxy container that will watch the Kubernetes API and trigger the appropriate operations against the “CSI volume driver” container. This is detailed in the “Suggested Mechanism for Deploying CSI Drivers on Kubernetes” section below.

The external component watching the Kubernetes API on behalf of the external CSI volume driver must handle provisioning, deleting, attaching, and detaching.

##### Provisioning and Deleting

Provisioning and deletion operations are handled using the existing [external provisioner mechanism](https://github.com/kubernetes-incubator/external-storage/tree/master/docs), where the external component watching the Kubernetes API on behalf of the external CSI volume driver will act as an external provisioner.

In short, to dynamically provision a new CSI volume, a cluster admin would create a `StorageClass` with the provisioner corresponding to the name of the external provisioner handling provisioning requests on behalf of the CSI volume driver.

To provision a new CSI volume, an end user would create a `PersistentVolumeClaim` object referencing this `StorageClass`. The external provisioner will react to the creation of the PVC and issue the `CreateVolume` call against the CSI volume driver to provision the volume. The `CreateVolume` name will be auto-generated as it is for other dynamically provisioned volumes. The `CreateVolume` capacity will be taken from the `PersistentVolumeClaim` object. The `CreateVolume` parameters will be passed through from the `StorageClass` parameters (opaque to Kubernetes).

If the `PersistentVolumeClaim` has the `volume.alpha.kubernetes.io/selected-node` annotation set (only added if delayed volume binding is enabled in the `StorageClass`), the provisioner will get relevant topology keys from the corresponding `CSINodeInfo` instance and the topology values from `Node` labels and use them to generate preferred topology in the `CreateVolume()` request. If the annotation is unset, preferred topology will not be specified (unless the PVC follows StatefulSet naming format, discussed later in this section). `AllowedTopologies` from the `StorageClass` is passed through as requisite topology. If `AllowedTopologies` is unspecified, the provisioner will pass in a set of aggregated topology values across the whole cluster as requisite topology.

To perform this topology aggregation, the external provisioner will cache all existing Node objects. In order to prevent a compromised node from affecting the provisioning process, it will pick a single node as the source of truth for keys, instead of relying on keys stored in `CSINodeInfo` for each node object. For PVCs to be provisioned with late binding, the selected node is the source of truth; otherwise a random node is picked. The provisioner will then iterate through all cached nodes that contain a node ID from the driver, aggregating labels using those keys. Note that if topology keys are different across the cluster, only a subset of nodes matching the topology keys of the chosen node will be considered for provisioning.

To generate preferred topology, the external provisioner will generate N segments for preferred topology in the `CreateVolume()` call, where N is the size of requisite topology. Multiple segments are included to support volumes that are available across multiple topological segments. The topology segment from the selected node will always be the first in preferred topology. All other segments are some reordering of remaining requisite topologies such that given a requisite topology (or any arbitrary reordering of it) and a selected node, the set of preferred topology is guaranteed to always be the same.

If immediate volume binding mode is set and the PVC follows StatefulSet naming format, then the provisioner will choose, as the first segment in preferred topology, a segment from requisite topology based on the PVC name that ensures an even spread of topology across the StatefulSet's volumes. The logic will be similar to the name hashing logic inside the GCE Persistent Disk provisioner. Other segments in preferred topology are ordered the same way as described above. This feature will be flag-gated in the external provisioner provided as part of the recommended deployment method.

Once the operation completes successfully, the external provisioner creates a `PersistentVolume` object to represent the volume using the information returned in the `CreateVolume` response. The topology of the returned volume is translated to the `PersistentVolume` `NodeAffinity` field. The `PersistentVolume` object is then bound to the `PersistentVolumeClaim` and available for use.

The format of topology key/value pairs is defined by the user and must match among the following locations:
* `Node` topology labels
* `PersistentVolume` `NodeAffinity` field
* `StorageClass` `AllowedTopologies` field
When a `StorageClass` has delayed volume binding enabled, the scheduler uses the topology information of a `Node` in the following ways:
  1. During dynamic provisioning, the scheduler selects a candidate node for the provisioner by comparing each `Node`'s topology with the `AllowedTopologies` in the `StorageClass`.
  1. During volume binding and pod scheduling, the scheduler selects a candidate node for the pod by comparing `Node` topology with `VolumeNodeAffinity` in `PersistentVolume`s.

A more detailed description can be found in the [topology-aware volume scheduling design doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/volume-topology-scheduling.md). See [Topology Representation in Node Objects](#topology-representation-in-node-objects) for the format used by the recommended deployment approach.

To delete a CSI volume, an end user would delete the corresponding `PersistentVolumeClaim` object. The external provisioner will react to the deletion of the PVC and based on its reclamation policy it will issue the `DeleteVolume` call against the CSI volume driver commands to delete the volume. It will then delete the `PersistentVolume` object.

##### Attaching and Detaching

Attach/detach operations must also be handled by an external component (an “attacher”). The attacher watches the Kubernetes API on behalf of the external CSI volume driver for new `VolumeAttachment` objects (defined below), and triggers the appropriate calls against the CSI volume driver to attach the volume. The attacher must watch for `VolumeAttachment` object and mark it as attached even if the underlying CSI driver does not support `ControllerPublishVolume` call, as Kubernetes has no knowledge about it.

More specifically, an external “attacher” must watch the Kubernetes API on behalf of the external CSI volume driver to handle attach/detach requests.

Once the following conditions are true, the external-attacher should call `ControllerPublishVolume` against the CSI volume driver to attach the volume to the specified node:

1. A new `VolumeAttachment` Kubernetes API objects is created by Kubernetes attach/detach controller.
2. The `VolumeAttachment.Spec.Attacher` value in that object corresponds to the name of the external attacher.
3. The `VolumeAttachment.Status.Attached` value is not yet set to true.
4. * Either a Kubernetes Node API object exists with the name matching `VolumeAttachment.Spec.NodeName` and that object contains a `csi.volume.kubernetes.io/nodeid` annotation. This annotation contains a JSON blob, a list of key/value pairs, where one of they keys corresponds with the CSI volume driver name, and the value is the NodeID for that driver. This NodeId mapping can be retrieved and used in the `ControllerPublishVolume` calls.
   * Or a `CSINodeInfo` API object exists with the name matching `VolumeAttachment.Spec.NodeName` and the object contains `CSIDriverInfo` for the CSI volume driver. The `CSIDriverInfo` contains NodeID for `ControllerPublishVolume` call.
5. The `VolumeAttachment.Metadata.DeletionTimestamp` is not set.

Before starting the `ControllerPublishVolume` operation, the external-attacher should add these finalizers to these Kubernetes API objects:

* To the `VolumeAttachment` so that when the object is deleted, the external-attacher has an opportunity to detach the volume first. External attacher removes this finalizer once the volume is fully detached from the node.
* To the `PersistentVolume` referenced by `VolumeAttachment` so the PV cannot be deleted while the volume is attached. External attacher needs information from the PV to perform detach operation. The attacher will remove the finalizer once all `VolumeAttachment` objects that refer to the PV are deleted, i.e. the volume is detached from all nodes.

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
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
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
    * The name of the `VolumeAttachment` object will be `pv-<SHA256(PVName+NodeName)>`.
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

Although Kubernetes does not dictate the packaging for a CSI volume driver, it offers the following recommendations to simplify deployment of a containerized CSI volume driver on Kubernetes.

![Recommended CSI Deployment Diagram](container-storage-interface_diagram1.png?raw=true "Recommended CSI Deployment Diagram")

To deploy a containerized third-party CSI volume driver, it is recommended that storage vendors:

  * Create a “CSI volume driver” container that implements the volume plugin behavior and exposes a gRPC interface via a unix domain socket, as defined in the CSI spec (including Controller, Node, and Identity services).
  * Bundle the “CSI volume driver” container with helper containers (external-attacher, external-provisioner, node-driver-registrar, cluster-driver-registrar, external-resizer, external-snapshotter, livenessprobe) that the Kubernetes team will provide (these helper containers will assist the “CSI volume driver” container in interacting with the Kubernetes system). More specifically, create the following Kubernetes objects:
    * To facilitate communication with the Kubernetes controllers, a `StatefulSet` or a `Deployment` (depending on the user's need; see [Cluster-Level Deployment](#cluster-level-deployment)) that has:
      * The following containers
        * The “CSI volume driver” container created by the storage vendor.
        * Containers provided by the Kubernetes team (all of which are optional):
          * `cluster-driver-registrar` (refer to the README in `cluster-driver-registrar` repository for when the container is required)
          * `external-provisioner` (required for provision/delete operations)
          * `external-attacher` (required for attach/detach operations. If you wish to skip the attach step, CSISkipAttach feature must be enabled in Kubernetes in addition to omitting this container)
          * `external-resizer` (required for resize operations)
          * `external-snapshotter` (required for volume-level snapshot operations)
          * `livenessprobe`
      * The following volumes:
        * `emptyDir` volume
          * Mounted by all containers, including the “CSI volume driver”.
          * The “CSI volume driver” container should create its Unix Domain Socket in this directory to enable communication with the Kubernetes helper container(s).
    * A `DaemonSet` (to facilitate communication with every instance of kubelet) that has:
      * The following containers
        * The “CSI volume driver” container created by the storage vendor.
        * Containers provided by the Kubernetes team:
          * `node-driver-registrar` - Responsible for registering the unix domain socket with kubelet.
          * `livenessprobe` (optional)
      * The following volumes:
        * `hostpath` volume
          * Expose `/var/lib/kubelet/plugins_registry` from the host.
          * Mount only in `node-driver-registrar` container at `/registration`
          * `node-driver-registrar` will use this unix domain socket to register the CSI driver’s unix domain socket with kubelet.
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

Containers provided by Kubernetes are maintained in [GitHub kubernetes-csi organization](https://github.com/kubernetes-csi).

#### Cluster-Level Deployment
Containers in the cluster-level deployment may be deployed in one of the following configurations:

1. StatefulSet with single replica. Good for clusters with a single dedicated node to run the cluster-level pod. A StatefulSet guarantees that no more than 1 instance of the pod will be running at once. One downside is that if the node becomes unresponsive, the replica will never be deleted and recreated.
1. Deployment with multiple replicas and leader election enabled (if supported by the container). Good for admins who prefer faster recovery time in case the main replica fails, at a cost of higher resource usage (especially memory).
1. Deployment with a single replica and leader election enabled (if supported by the container). A compromise between the above two options. If the replica is detected to be failed, a new replica can be scheduled almost immediately.

Note that certain cluster-level containers, such as `external-provisioner`, `external-attacher`, `external-resizer`, and `external-snapshotter`, may require credentials to the storage backend, and as such, admins may choose to run them on dedicated "infrastructure" nodes (such as master nodes) that don't run user pods.

#### Topology Representation in Node Objects
Topology information will be represented as labels.

Requirements:
* Must adhere to the [label format](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
* Must support different drivers on the same node.
* The format of each key/value pair must match those in `PersistentVolume` and `StorageClass` objects, as described in the [Provisioning and Deleting](#provisioning-and-deleting) section.

Proposal: `"com.example.topology/rack": "rack1"`
The list of topology keys known to the driver is stored separately in the `CSINodeInfo` object.

Justifications:
* No strange separators needed, comparing to the alternative. Cleaner format.
* The same topology key could be used across different components (different storage plugin, network plugin, etc.)
* Once NodeRestriction is moved to the newer model (see [here](https://github.com/kubernetes/community/pull/911) for context), for each new label prefix introduced in a new driver, the cluster admin has to configure NodeRestrictions to allow the driver to update labels with the prefix. Cluster installations could include certain prefixes for pre-installed drivers by default. This is less convenient compared to the alternative, which can allow editing of all CSI drivers by default using the “csi.kubernetes.io” prefix, but often times cluster admins have to whitelist those prefixes anyway (for example ‘cloud.google.com’)

Considerations:
* Upon driver deletion/upgrade/downgrade, stale labels will be left untouched. It’s difficult for the driver to decide whether other components outside CSI rely on this label.
* During driver installation/upgrade/downgrade, controller deployment must be brought down before node deployment, and node deployment must be deployed before the controller deployment, because provisioning relies on up-to-date node information. One possible issue is if only topology values change while keys remain the same, and if AllowedTopologies is not specified, requisite topology will contain both old and new topology values, and CSI driver may fail the CreateVolume() call. Given that CSI driver should be backward compatible, this is more of an issue when a node rolling upgrade happens before the controller update. It's not an issue if keys are changed as well since requisite and preferred topology generation handles it appropriately.
* During driver installation/upgrade/downgrade, if a version of the controller (either old or new) is running while there is an ongoing rolling upgrade with the node deployment, and the new version of the CSI driver reports different topology information, nodes in the cluster may have different versions of topology information. However, this doesn't pose an issue. If AllowedTopologies is specified, a subset of nodes matching the version of topology information in AllowedTopologies will be used as provisioning candidate. If AllowedTopologies is not specified, a single node is used as the source of truth for keys
* Topology keys inside `CSINodeInfo` must reflect the topology keys from drivers currently installed on the node. If no driver is installed, the collection must be empty. However, due to the possible race condition between kubelet (the writer) and the external provisioner (the reader), the provisioner must gracefully handle the case where `CSINodeInfo` is not up-to-date. In the current design, the provisioner will erroneously provision a volume on a node where it's inaccessible.

Alternative:
1. `"csi.kubernetes.io/topology.example.com_rack": "rack1"`

#### Topology Representation in PersistentVolume Objects
There exists multiple ways to represent a single topology as NodeAffinity. For example, suppose a `CreateVolumeResponse` contains the following accessible topology:

```yaml
- zone: "a"
  rack: "1"
- zone: "b"
  rack: "1"
- zone: "b"
  rack: "2"
```

There are at least 3 ways to represent this in NodeAffinity (excluding `nodeAffinity`, `required`, and `nodeSelectorTerms` for simplicity):

Form 1 - `values` contain exactly 1 element.
```yaml
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "a"
  - key: rack
    operator: In
    values:
    - "1"
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "b"
  - key: rack
    operator: In
    values:
    - "1"
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "b"
  - key: rack
    operator: In
    values:
    - "2"
```

Form 2 - Reduced by `rack`.
```yaml
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "a"
    - "b"
  - key: rack
    operator: In
    values:
    - "1"
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "b"
  - key: rack
    operator: In
    values:
    - "2"
```
Form 3 - Reduced by `zone`.
```yaml
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "a"
  - key: rack
    operator: In
    values:
    - "1"
- matchExpressions:
  - key: zone
    operator: In
    values:
    - "b"
  - key: rack
    operator: In
    values:
    - "1"
    - "2"
```
The provisioner will always choose Form 1, i.e. all `values` will have at most 1 element. Reduction logic could be added in future versions to arbitrarily choose a valid and simpler form like Forms 2 & 3.

#### Upgrade & Downgrade Considerations
When drivers are uninstalled, topology information stored in Node labels remain untouched. The recommended label format allows multiple sources (such as CSI, networking resources, etc.) to share the same label key, so it's nontrivial to accurately determine whether a label is still used.

In order to upgrade drivers using the recommended driver deployment mechanism, the user is recommended to tear down the StatefulSet (controller components) before the DaemonSet (node components), and deploy the DaemonSet before the StatefulSet. There may be design improvements to eliminate this constraint, but it will be evaluated at a later iteration.

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
2. The external-provisioner for the CSI driver sees the `PersistentVolumeClaim` was deleted and triggers the retention policy:
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

CSI allows specifying credentials in CreateVolume/DeleteVolume, ControllerPublishVolume/ControllerUnpublishVolume, NodeStageVolume/NodeUnstageVolume, and NodePublishVolume/NodeUnpublishVolume operations.

Kubernetes will enable cluster admins and users deploying workloads on the cluster to specify these credentials by referencing Kubernetes secret object(s). Kubernetes (either the core components or helper containers) will fetch the secret(s) and pass them to the CSI volume plugin.

If a secret object contains more than one secret, all secrets are passed.

#### Secret to CSI Credential Encoding

CSI accepts credentials for all the operations specified above as a map of string to string (e.g. `map<string, string> controller_create_credentials`).

Kubernetes, however, defines secrets as a map of string to byte-array (e.g. `Data map[string][]byte`). It also allows specifying text secret data in string form via a write-only convenience field `StringData` which is a map of string to string.

Therefore, before passing secret data to CSI, Kubernetes (either the core components or helper containers) will convert the secret data from bytes to string (Kubernetes does not specify the character encoding, but Kubernetes internally uses golang to cast from string to byte and vice versa which assumes UTF-8 character set).

Although CSI only accepts string data, a plugin MAY dictate in its documentation that a specific secret contain binary data and specify a binary-to-text encoding to use (base64, quoted-printable, etc.) to encode the binary data and allow it to be passed in as a string. It is the responsibility of the entity (cluster admin, user, etc.) that creates the secret to ensure its content is what the plugin expects and is encoded in the format the plugin expects.

#### CreateVolume/DeleteVolume Credentials

The CSI CreateVolume/DeleteVolume calls are responsible for creating and deleting volumes.
These calls are executed by the CSI external-provisioner.
Credentials for these calls will be specified in the Kubernetes `StorageClass` object.

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: fast-storage
provisioner: com.example.team.csi-driver
parameters:
  type: pd-ssd
  csiProvisionerSecretName: mysecret
  csiProvisionerSecretNamespace: mynamespaace
```

The CSI external-provisioner will reserve the parameter keys `csiProvisionerSecretName` and `csiProvisionerSecretNamespace`. If specified, the CSI Provisioner will fetch the secret `csiProvisionerSecretName` in the Kubernetes namespace `csiProvisionerSecretNamespace` and pass it to:
1. The CSI `CreateVolumeRequest` in the `controller_create_credentials` field.
2. The CSI `DeleteVolumeRequest` in the `controller_delete_credentials` field.

See "Secret to CSI Credential Encoding" section above for details on how secrets will be mapped to CSI credentials.

It is assumed that since `StorageClass` is a non-namespaced field, only trusted users (e.g. cluster administrators) should be able to create a `StorageClass` and, thus, specify which secret to fetch.

The only Kubernetes component that needs access to this secret is the CSI external-provisioner, which would fetch this secret. The permissions for the external-provisioner may be limited to the specified (external-provisioner specific) namespace to prevent a compromised provisioner from gaining access to other secrets.

#### ControllerPublishVolume/ControllerUnpublishVolume Credentials

The CSI ControllerPublishVolume/ControllerUnpublishVolume calls are responsible for attaching and detaching volumes.
These calls are executed by the CSI external-attacher.
Credentials for these calls will be specified in the Kubernetes `CSIPersistentVolumeSource` object.

```go
type CSIPersistentVolumeSource struct {

  // ControllerPublishSecretRef is a reference to the secret object containing
  // sensitive information to pass to the CSI driver to complete the CSI
  // ControllerPublishVolume and ControllerUnpublishVolume calls.
  // This secret will be fetched by the external-attacher.
  // This field is optional, and  may be empty if no secret is required. If the
  // secret object contains more than one secret, all secrets are passed.
  // +optional
  ControllerPublishSecretRef *SecretReference
}
```

If specified, the CSI external-attacher will fetch the Kubernetes secret referenced by `ControllerPublishSecretRef` and pass it to:
1. The CSI `ControllerPublishVolume` in the `controller_publish_credentials` field.
2. The CSI `ControllerUnpublishVolume` in the `controller_unpublish_credentials` field.

See "Secret to CSI Credential Encoding" section above for details on how secrets will be mapped to CSI credentials.

It is assumed that since `PersistentVolume` objects are non-namespaced and `CSIPersistentVolumeSource` can only be referenced via a `PersistentVolume`, only trusted users (e.g. cluster administrators) should be able to create a `PersistentVolume` objects and, thus, specify which secret to fetch.

The only Kubernetes component that needs access to this secret is the CSI external-attacher, which would fetch this secret. The permissions for the external-attacher may be limited to the specified (external-attacher specific) namespace to prevent a compromised attacher from gaining access to other secrets.

#### NodeStageVolume/NodeUnstageVolume Credentials

The CSI NodeStageVolume/NodeUnstageVolume calls are responsible for mounting (setup) and unmounting (teardown) volumes.
These calls are executed by the Kubernetes node agent (kubelet).
Credentials for these calls will be specified in the Kubernetes `CSIPersistentVolumeSource` object.

```go
type CSIPersistentVolumeSource struct {

  // NodeStageSecretRef is a reference to the secret object containing sensitive
  // information to pass to the CSI driver to complete the CSI NodeStageVolume
  // and NodeStageVolume and NodeUnstageVolume calls.
  // This secret will be fetched by the kubelet.
  // This field is optional, and  may be empty if no secret is required. If the
  // secret object contains more than one secret, all secrets are passed.
  // +optional
  NodeStageSecretRef *SecretReference
}
```

If specified, the kubelet will fetch the Kubernetes secret referenced by `NodeStageSecretRef` and pass it to:
1. The CSI `NodeStageVolume` in the `node_stage_credentials` field.
2. The CSI `NodeUnstageVolume` in the `node_unstage_credentials` field.

See "Secret to CSI Credential Encoding" section above for details on how secrets will be mapped to CSI credentials.

It is assumed that since `PersistentVolume` objects are non-namespaced and `CSIPersistentVolumeSource` can only be referenced via a `PersistentVolume`, only trusted users (e.g. cluster administrators) should be able to create a `PersistentVolume` objects and, thus, specify which secret to fetch.

The only Kubernetes component that needs access to this secret is the kubelet, which would fetch this secret. The permissions for the kubelet may be limited to the specified (kubelet specific) namespace to prevent a compromised attacher from gaining access to other secrets.

The Kubernetes API server's node authorizer must be updated to allow kubelet to access the secrets referenced by `CSIPersistentVolumeSource.NodeStageSecretRef`.

#### NodePublishVolume/NodeUnpublishVolume Credentials

The CSI NodePublishVolume/NodeUnpublishVolume calls are responsible for mounting (setup) and unmounting (teardown) volumes.
These calls are executed by the Kubernetes node agent (kubelet).
Credentials for these calls will be specified in the Kubernetes `CSIPersistentVolumeSource` object.

```go
type CSIPersistentVolumeSource struct {

  // NodePublishSecretRef is a reference to the secret object containing
  // sensitive information to pass to the CSI driver to complete the CSI
  // NodePublishVolume and NodeUnpublishVolume calls.
  // This secret will be fetched by the kubelet.
  // This field is optional, and  may be empty if no secret is required. If the
  // secret object contains more than one secret, all secrets are passed.
  // +optional
  NodePublishSecretRef *SecretReference
}
```

If specified, the kubelet will fetch the Kubernetes secret referenced by `NodePublishSecretRef` and pass it to:
1. The CSI `NodePublishVolume` in the `node_publish_credentials` field.
2. The CSI `NodeUnpublishVolume` in the `node_unpublish_credentials` field.

See "Secret to CSI Credential Encoding" section above for details on how secrets will be mapped to CSI credentials.

It is assumed that since `PersistentVolume` objects are non-namespaced and `CSIPersistentVolumeSource` can only be referenced via a `PersistentVolume`, only trusted users (e.g. cluster administrators) should be able to create a `PersistentVolume` objects and, thus, specify which secret to fetch.

The only Kubernetes component that needs access to this secret is the kubelet, which would fetch this secret. The permissions for the kubelet may be limited to the specified (kubelet specific) namespace to prevent a compromised attacher from gaining access to other secrets.

The Kubernetes API server's node authorizer must be updated to allow kubelet to access the secrets referenced by `CSIPersistentVolumeSource.NodePublishSecretRef`.

## Alternatives Considered

### Extending PersistentVolume Object

Instead of creating a new `VolumeAttachment` object, another option we considered was extending the existing `PersistentVolume` object.

`PersistentVolumeSpec` would be extended to include:
* List of nodes to attach the volume to (initially empty).

`PersistentVolumeStatus` would be extended to include:
* List of nodes the volume was successfully attached to.

We dismissed this approach because having attach/detach triggered by the creation/deletion of an object is much easier to manage (for both external-attacher and Kubernetes) and more robust (fewer corner cases to worry about).


[Flex Volume]: /contributors/devel/sig-storage/flexvolume.md
[Flex Volume plugin]: /contributors/devel/sig-storage/flexvolume.md
