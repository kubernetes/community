# In-tree Storage Plugin to CSI Migration Design Doc

Authors: @davidz627, @jsafrane

This document presents a detailed design for migrating in-tree storage plugins
to CSI. This will be an opt-in feature turned on at cluster creation time that
will redirect in-tree plugin operations to a corresponding CSI Driver.

## Glossary

* ADC (Attach Detach Controller): Controller binary that handles Attach and Detach portion of a volume lifecycle
* Kubelet: Kubernetes component that runs on each node, it handles the Mounting and Unmounting portion of volume lifecycle
* CSI (Container Storage Interface): An RPC interface that Kubernetes uses to interface with arbitrary 3rd party storage drivers
* In-tree: Code that is compiled into native Kubernetes binaries
* Out-of-tree: Code that is not compiled into Kubernetes binaries, but can be run as Deployments on Kubernetes

## Background and Motivations

The Kubernetes volume plugins are currently in-tree meaning all logic and
handling for each plugin lives in the Kubernetes codebase itself. With the
Container Storage Interface (CSI) the goal is to move those plugins out-of-tree.
CSI defines a standard interface for communication between the Container
Orchestrator (CO), Kubernetes in our case, and the storage plugins.

As the CSI Spec moves towards GA and more storage plugins are being created and
becoming production ready, we will want to migrate our in-tree plugin logic to
use CSI plugins instead. This is motivated by the fact that we are currently
supporting two versions of each plugin (one in-tree and one CSI), and that we
want to eventually migrate all storage users to CSI.

In order to do this we need to migrate the internals of the in-tree plugins to
call out to CSI Plugins because we will be unable to deprecate the current
internal plugin API’s due to Kubernetes API deprecation policies. This will
lower cost of development as we only have to maintain one version of each
plugin, as well as ease the transition to CSI when we are able to deprecate the
internal APIs.


## Goals

* Compile all requirements for a successful transition of the in-tree plugins to
  CSI
    * As little code as possible remains in the Kubernetes Repo
    * In-tree plugin API is untouched, user Pods and PVs continue working after
      upgrades
    * Minimize user visible changes
* Design a robust mechanism for redirecting in-tree plugin usage to appropriate
  CSI drivers, while supporting seamless upgrade and downgrade between new
  Kubernetes version that uses CSI drivers for in-tree volume plugins to an old
  Kubernetes version that uses old-fashioned volume plugins without CSI.
* Design framework for migration that allows for easy interface extension by
  in-tree plugin authors to “migrate” their plugins.
    * Migration must be modular so that each plugin can have migration turned on
      and off separately

## Non-Goals

* Design a mechanism for deploying  CSI drivers on all systems so that users can
  use the current storage system the same way they do today without having to do
  extra set up.
* Implementing CSI Drivers for existing plugins
* Define set of volume plugins that should be migrated to CSI

## Implementation Schedule

Alpha [1.14]
* Off by default
* Proof of concept migration of at least 2 storage plugins [AWS, GCE]
* Framework for plugin migration built for Dynamic provisioning, pre-provisioned
  volumes, and in-tree volumes

Beta [Target 1.15]
* On by default
* Migrate all of the cloud provider plugins*

GA [TBD]
* Feature on by default, per-plugin toggle on for relevant cloud provider by
  default
* CSI Drivers for migrated plugins available on related cloud provider cluster
  by default

## Milestones

* Translation Library implemented in Kubernetes staging
* Migration Shim for Provision, Attach, Detach, Mount, Unmount (including Inline Volumes)
* Migration Shim for Resize, Block
* CSI Driver lifecycle manager
* GCE PD feature parity in CSI with in-tree implementation
* AWS EBS feature parity in CSI with in-tree implementation
* Cloud Driver feature parity in CSI with in-tree implementation

## Dependency Graph

![CSI Migration Dependency Diagram](csi-migration_dependencies.png?raw=true "CSI Migration Dependency Diagram")

## Feature Gating

We will have an alpha feature gate for the whole feature that can turn the CSI
migration on or off, when off all code paths should revert/stay with the in-tree
plugins. We will also have individual flags for each driver so that admins can
toggle them on or off.

The feature gate can exist at the interception points in the OperationGenerator
for Attach and Mount, as well as in the PV Controller for Provisioning.

We will also have one feature flag for each driver’s migration so that each
driver migration can be turned on and off individually. 

The new feature gates for alpha are:
```
// Enables the in-tree storage to CSI Plugin migration feature.
CSIMigration utilfeature.Feature = "CSIMigration"

// Enables the GCE PD in-tree driver to GCE CSI Driver migration feature.
CSIMigrationGCE utilfeature.Feature = "CSIMigrationGCE"

// Enables the AWS in-tree driver to AWS CSI Driver migration feature.
CSIMigrationAWS utilfeature.Feature = "CSIMigrationAWS"
```

## Translation Layer

The main mechanism we will use to migrate plugins is redirecting in-tree
operation calls to the CSI Driver instead of the in-tree driver, the external
components will pick up these in-tree PV's and use a translation library to
translate to CSI Source.

Pros:
* Keeps old API objects as they are
* Facilitates gradual roll-over to CSI

Cons:
* Somewhat complicated and error prone.
* Bespoke translation logic for each in-tree plugin

### Dynamically Provisioned Volumes

#### Kubernetes Changes

Dynamically Provisioned volumes will continue to be provisioned with the in-tree
`PersistentVolumeSource`. The CSI external-provisioner to pick up the
in-tree PVC's when migration is turned on and provision using the CSI Drivers;
it will then use the imported translation library to return with a PV that contains an equivalent of the original
in-tree PV. The PV will then go through all the same steps outlined below in the
"Non-Dynamic Provisioned Volumes" for the rest of the volume lifecycle.

#### Leader Election

There will have to be some mechanism to switch between in-tree and external
provisioner when the migration feature is turned on/off. The two should be
compatible as they both will create the same volume and PV based on the same
PVC, as well as both be able to delete the same PV/PVCs. The in-tree provisioner
will have logic added so that it will stand down and mark the PV as "migrated"
with an annotation  when the migration is turned on and the external provisioner
will take care of the PV when it sees the annotation.

### Translation Library

In order to make this on-the-fly translation work we will develop a separate
translation library. This library will have to be able to translate from in-tree
PV Source to the equivalent CSI Source. This library can then be imported by
both Kubernetes and the external CSI Components to translate Volume Sources when
necessary. The cost of doing this translation will be very low as it will be an
imported library and part of whatever binary needs the translation (no extra
API or RPC calls).

#### Library Interface

```
type CSITranslator interface {
  // TranslateInTreePVToCSI takes a persistent volume and will translate
  // the in-tree source to a CSI Source if the translation logic
  // has been implemented. The input persistent volume will not
  // be modified
  TranslateInTreePVToCSI(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {

  // TranslateCSIPVToInTree takes a PV with a CSI PersistentVolume Source and will translate
  // it to a in-tree Persistent Volume Source for the specific in-tree volume specified
  // by the `Driver` field in the CSI Source. The input PV object will not be modified.
  TranslateCSIPVToInTree(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {

  // TranslateInTreeInlineVolumeToPVSpec takes an inline intree volume and will translate
  // the in-tree volume source to a PersistentVolumeSpec containing a CSIPersistentVolumeSource
  TranslateInTreeInlineVolumeToPVSpec(volume *v1.Volume) (*v1.PersistentVolumeSpec, error) {

  // IsMigratableByName tests whether there is Migration logic for the in-tree plugin
  // for the given `pluginName`
  IsMigratableByName(pluginName string) bool {

  // GetCSINameFromIntreeName maps the name of a CSI driver to its in-tree version
  GetCSINameFromIntreeName(pluginName string) (string, error) {

  // IsPVMigratable tests whether there is Migration logic for the given Persistent Volume
  IsPVMigratable(pv *v1.PersistentVolume) bool {

  // IsInlineMigratable tests whether there is Migration logic for the given Inline Volume
  IsInlineMigratable(vol *v1.Volume) bool {
}
```

#### Library Versioning

Since the library will be imported by various components it is imperative that
all components import a version of the library that supports in-tree driver x
before the migration feature flag for x is turned on. If not, the TranslateToCSI
function will return an error when the translation is attempted.


### Pre-Provisioned Volumes (and volumes provisioned before migration)

In the OperationGenerator at the start of each volume operation call we will
check to see whether the plugin has been migrated.

For Controller calls, we will call the CSI calls instead of the in-tree calls.
The OperationGenerator can do the translation of the PV Source before handing it
to the CSI calls, therefore the CSI in-tree plugin will only have to deal with
what it sees as a CSI Volume. Special care must be taken that `volumeHandle` is
unique and also deterministic so that we can always find the correct volume. 
We also foresee that future controller calls such as resize and snapshot will use a similar mechanism. All these external components
will also need to be updated to accept PV's of any source type when it is given
and use the translation library to translate the in-tree PV Source into a CSI
Source when necessary.

For Node calls, the VolumeToMount object will contain the in-tree PV Source,
this can then be translated by the translation library when needed and
information can be fed to the CSI components when necessary.

Then the rest of the code in the Operation Generator can execute as normal with
the CSI Plugin and the annotation in the requisite locations.

Caveat: For ALL detach calls of plugins that MAY have already been migrated we
have to attempt to DELETE the VolumeAttachment object that would have been
created if that plugin was migrated. This is because Attach after migration
creates a VolumeAttachment object, and if for some reason we are doing a detach
with the in-tree plugin, the VolumeAttachment object becomes orphaned.


### In-line Volumes

In-line controller calls are a special case because there is no PV. In this case,
we will translate the in-line Volume into a PersistentVolumeSpec using
plugin-specific translation logic in the CSI translation library method,
`TranslateInTreeInlineVolumeToPVSpec`. The resulting PersistentVolumeSpec will
be stored in a new field `VolumeAttachment.Spec.Source.VolumeAttachmentSource.InlineVolumeSpec`.

The plugin-specific CSI translation logic invoked by `TranslateInTreeInlineVolumeToPVSpec`
will need to populate the `CSIPersistentVolumeSource` field along with appropriate
values for `AccessModes` and `MountOptions` fields in
`VolumeAttachment.Spec.Source.VolumeAttachmentSource.InlineVolumeSpec`. Since
`AccessModes` and `MountOptions` are not specified for inline volumes, default values
for these fields suitable for the CSI plugin will need to be populated in addition
to translation logic to populate `CSIPersistentVolumeSource`.

The VolumeAttachment name must be made with the CSI translated version of the
VolumeSource in order for it to be discoverable by Detach and WaitForAttach
(described in more detail below).

The CSI Attacher will have to be modified to also check for `InlineVolumeSpec`
besides the `PersistentVolumeName`. Only one of the two may be specified. If `PersistentVolumeName`
is empty and `InlineVolumeSpec` is set, the CSI Attacher will not look for
an associated PV in it's PV informer cache as it implies the inline volume scenario
(where no PVs are created).

The CSI Attacher will have access to all the data it requires for handling in-line
volumes attachment (through the CSI plugins) from fields in the `InlineVolumeSpec`.

The new VolumeAttachmentSource API will look as such:
```
// VolumeAttachmentSource represents a volume that should be attached.
// Inline volumes and Persistent volumes can be attached via external attacher.
// Exactly one member can be set.
type VolumeAttachmentSource struct {
	// Name of the persistent volume to attach.
	// +optional
	PersistentVolumeName *string `json:"persistentVolumeName,omitempty" protobuf:"bytes,1,opt,name=persistentVolumeName"`

	// A PersistentVolumeSpec whose fields contain translated data from a pod's inline
	// VolumeSource to support shimming of in-tree inline volumes to a CSI backend.
	// This field is alpha-level and is only honored by servers that
	// enable the CSIMigration feature.
	// +optional
	InlineVolumeSpec *v1.PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty" protobuf:"bytes,2,opt,name=inlineVolumeSpec"`
}
```

We need to be careful with naming VolumeAttachments for in-line volumes. The
name needs to be unique and ADC must be able to find the right VolumeAttachment
when a pod is deleted (i.e. using only info in Node.Status). CSI driver in
kubelet must be able to find the VolumeAttachment too to call WaitForAttach and
VolumesAreAttached.

The attachment name is usually a hash of the volume name, CSI Driver name, and
Node name. We are able to get all this information for Detach and WaitForAttach
by translating the in-tree inline volume source to a CSI volume source before
passing it to to the volume operations.

There is currently a race condition in in-tree inline volumes where if a pod
object is deleted and the ADC restarts we lose the information for the inline
volume and will not be able to detach the volume. This is a known issue and we
will retain the same behavior with migrated inline volumes. However, we may be
able to solve this in the future by reconciling the VolumeAttachment object with
existing Pods in the ADC.


### Volume Resize
#### Offline Resizing
For controller expansion, in the in-tree resize controller, we will create a new PVC annotation `volume.kubernetes.io/storage-resizer`
and set the value to the name of resizer. If the PV is CSI PV or migrated in-tree PV, the annotation will be set to 
the name of CSI driver; otherwise, it will be set to the name of in-tree plugin.

For migrated volume, The CSI resizer name will be derived from translating in-tree plugin name
to CSI driver name by translation library. We will also add an event to PVC about resizing being handled
by external controller.

For external resizer, we will update it to expand volume for both CSI volume and in-tree 
volume (only if migration is enabled). For migrated in-tree volume, it will update in-tree PV object
with new volume size and mark in-tree PVC as resizing finished.

To synchronize between in-tree resizer and external resizer, external resizer will find resizer name
using PVC annotation `volume.kubernetes.io/storage-resizer`. Since `volume.kubernetes.io/storage-resizer`
annotation defines the CSI plugin name which will handle external resizing, it should
match driver running with external-resizer, hence external resizer will proceed with volume resizing. Otherwise,
it will yield to in-tree resizer.

For filesystem expansion, in the OperationGenerator, `GenerateMountVolumeFunc` is used to expand file system after volume
is expanded and staged/mounted. The migration logic is covered by previous migration of volume mount.

#### Online Resizing
Handling online resizing does not require anything special in control plane. The behaviour will be
same as offline resizing. 

To handle expansion on kubelet - we will convert volume spec to CSI spec before handling the call
to volume plugin inside `GenerateExpandVolumeFSWithoutUnmountingFunc`.

### Raw Block
In the OperationGenerator, `GenerateMapVolumeFunc`, `GenerateUnmapVolumeFunc` and 
`GenerateUnmapDeviceFunc` are used to prepare and mount/umount block devices. At the 
beginning of each API, we will check whether migration is enabled for the plugin. If
enabled, volume spec will be translated from the in-tree spec to out-of-tree spec using
CSI as the persistence volume source.

Caveat: the original spec needs to be used when setting the state of `actualStateOfWorld`
for where is it used before the translation.

### Volume Reconstruction

Volume Reconstruction is currently a routine in the reconciler that runs on the
nodes when a Kubelet restarts and loses its cached state (`desiredState` and
`actualState`). It is kicked off in `syncStates()` in
`pkg/kubeletvolumemanager/reconciler/reconciler.go` and attempts to reconstruct
a volume based on the mount path on the host machine.

When CSI Migration is turned on, when the reconstruction code is run and it
finds a CSI mounted volume we currently do not know whether it was mounted as a
native CSI volume or migrated from in-tree. To solve this issue we will save a
`migratedVolume` boolean in the `saveVolumeData` function when the `NewMounter`
is created during the `MountVolume` call for that particular volume in the
Operation generator.

When the Kubelet is restarted and we lose state the Kubelet will call
`reconstructVolume` we can `loadVolumeData` and determine whether that CSI
volume was migrated or not, as well as get the information about the original
plugin requested. With that information we should be able to call the
`ReconstructVolumeOperation` with the correct in-tree plugin to get the original
in-tree spec that we can then pass to the rest of volume reconstruction. The
rest of the volume reconstruction code will then use this in-tree spec passed to
the `desiredState`, `actualState`, and `operationGenerator` and the volume will
go through the standard volume pathways and go through the standard migrated
volume lifecycles described above in the "Pre-Provisioned Volumes" section.

### Volume Limit

TODO: Design

## Interactions with PV-PVC Protection Finalizers

PV-PVC Protection finalizers prevent deletion of a PV when it is bound to a PVC,
and prevent deletion of a PVC when it is in use by a pod.

There is no known issue with interaction here. The finalizers will still work in
the same ways as we are not removing/adding PV’s or PVC’s in out of the ordinary
ways.

## Dealing with CSI Driver Failures

Plugin should fail if the CSI Driver is down and migration is turned on. When
the driver recovers we should be able to resume gracefully.

We will also create a playbook entry for how to turn off the CSI Driver
migration gracefully, how to tell when the CSI Driver is broken or non-existent,
and how to redeploy a CSI Driver in a cluster.

## API Changes

### CSINodeInfo API

Changes in: https://github.com/kubernetes/kubernetes/pull/70515

#### Old CSINodeInfo API

```
// CSINodeInfo holds information about all CSI drivers installed on a node.
type CSINodeInfo struct {
	metav1.TypeMeta `json:",inline"`

	// metadata.name must be the Kubernetes node name.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// List of CSI drivers running on the node and their properties.
	// +patchMergeKey=driver
	// +patchStrategy=merge
	CSIDrivers []CSIDriverInfo `json:"csiDrivers" patchStrategy:"merge" patchMergeKey:"driver"`
}

// CSIDriverInfo contains information about one CSI driver installed on a node.
type CSIDriverInfo struct {
	// driver is the name of the CSI driver that this object refers to.
	// This MUST be the same name returned by the CSI GetPluginName() call for
	// that driver.
	Driver string `json:"driver"`

	// nodeID of the node from the driver point of view.
	// This field enables Kubernetes to communicate with storage systems that do
	// not share the same nomenclature for nodes. For example, Kubernetes may
	// refer to a given node as "node1", but the storage system may refer to
	// the same node as "nodeA". When Kubernetes issues a command to the storage
	// system to attach a volume to a specific node, it can use this field to
	// refer to the node name using the ID that the storage system will
	// understand, e.g. "nodeA" instead of "node1".
	NodeID string `json:"nodeID"`

	// topologyKeys is the list of keys supported by the driver.
	// When a driver is initialized on a cluster, it provides a set of topology
	// keys that it understands (e.g. "company.com/zone", "company.com/region").
	// When a driver is initialized on a node it provides the same topology keys
	// along with values that kubelet applies to the coresponding node API
	// object as labels.
	// When Kubernetes does topology aware provisioning, it can use this list to
	// determine which labels it should retrieve from the node object and pass
	// back to the driver.
	TopologyKeys []string `json:"topologyKeys"`
}
```

#### New CSINodeInfo API 

```
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSINodeInfo holds information about all CSI drivers installed on a node.
// CSI drivers do not need to create the CSINodeInfo object directly. As long as
// they use the node-driver-registrar sidecar container, the kubelet will
// automatically populate the CSINodeInfo object for the CSI driver as part of
// kubelet plugin registration.
// CSINodeInfo has the same name as a node. If it is missing, it means either
// there are no CSI Drivers available on the node, or the Kubelet version is low
// enough that it doesn't create this object.
// CSINodeInfo has an OwnerReference that points to the corresponding node object.
type CSINodeInfo struct {
	metav1.TypeMeta

	// metadata.name must be the Kubernetes node name.
	metav1.ObjectMeta

	// spec is the specification of CSINodeInfo
	Spec CSINodeInfoSpec
}

// CSINodeInfoSpec holds information about the specification of all CSI drivers installed on a node
type CSINodeInfoSpec struct {
	// drivers is a list of information of all CSI Drivers existing on a node.
	// It can be empty on initialization.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Drivers []CSIDriverInfoSpec
}

// CSIDriverInfoSpec holds information about the specification of one CSI driver installed on a node
type CSIDriverInfoSpec struct {
	// This is the name of the CSI driver that this object refers to.
	// This MUST be the same name returned by the CSI GetPluginName() call for
	// that driver.
	Name string

	// nodeID of the node from the driver point of view.
	// This field enables Kubernetes to communicate with storage systems that do
	// not share the same nomenclature for nodes. For example, Kubernetes may
	// refer to a given node as "node1", but the storage system may refer to
	// the same node as "nodeA". When Kubernetes issues a command to the storage
	// system to attach a volume to a specific node, it can use this field to
	// refer to the node name using the ID that the storage system will
	// understand, e.g. "nodeA" instead of "node1".
	// This field must be populated. An empty string means NodeID is not initialized
	// by the driver and it is invalid.
	NodeID string

	// topologyKeys is the list of keys supported by the driver.
	// When a driver is initialized on a cluster, it provides a set of topology
	// keys that it understands (e.g. "company.com/zone", "company.com/region").
	// When a driver is initialized on a node, it provides the same topology keys
	// along with values. Kubelet will expose these topology keys as labels
	// on its own node object.
	// When Kubernetes does topology aware provisioning, it can use this list to
	// determine which labels it should retrieve from the node object and pass
	// back to the driver.
	// It is possible for different nodes to use different topology keys.
	// This can be empty if driver does not support topology.
	// +optional
	TopologyKeys []string
}
```

#### API Lifecycle

A new `CSINodeInfo` API object is created for each node by the Kubelet on
Kubelet initialization before pods are able to be scheduled. A driver will be
added with all of its information populated when a driver is registered through
the plugin registration mechanism. When the driver is unregistered through the
plugin registration mechanism it's entry will be removed from the `Drivers` list
in the `CSINodeInfoSpec`.

#### Kubelet Initialization & Migration Annotation

On Kubelet initialization we will also pre-populate an annotation for that
node's `CSINodeInfo`. The key will be
`storage.alpha.kubernetes.io/migrated-plugins` and the value will be a list of
in-tree plugin names that the Kubelet has the migration shim turned on for
(through feature flags). This must be populated before the Kubelet becomes
schedulable in order to achieve synchronization described in the "ADC and
Kubelete CSI/In-tree Sync" section below".

## Upgrade/Downgrade, Migrate/Un-migrate

### Feature Flags

ADC and Kubelet use the "same" feature flags, but in reality they are passed in
to each binary separately. There will be a feature flag per driver as well as
one for CSIMigration in general.

Kubelet will use its own feature flags to determine whether to use the in-tree
or csi backend for Kubelet storage lifecycle operations, as well as to add the
plugins that have the feature flag on to the
`storage.alpha.kubernetes.io/migrated-plugins` annotation of `CSINodeInfo` for
the node that Kubelet is running on.

The ADC will also use its own feature flags to help make the determination
whether to use in-tree or CSI backend for ADC storage lifecycle operations. The
other component to help determine which backend to use will be outlined below in
the "ADC and Kubelet CSI/In-tree Sync" section.

### ADC and Kubelet CSI/In-tree Sync

Some plugins have subtly different behavior on both ADC and Kubelet side between
in-tree and CSI implementations. Therefore it is important that if the ADC is to
use the in-tree implementation, the Kubelet must as well - and if the ADC is to
use the CSI Migrated implementation, the Kubelet must as well. Therefore we will
implement a mechanism to keep the ADC and the Kubelet in sync about the Kubelets
abilities as well as the feature gates active in each.

In order for the ADC controller to have the requisite information from the
Kubelet to make informed decisions the Kubelet must propagate the
`storage.alpha.kubernetes.io/migrated-plugins` annotation information for each
potentially migrated driver on Kubelet startup and be considered `NotReady`
until that information is synced to the API server. This gives is the following
guarantees:
* If `CSINodeInfo` for the node does not exist, then ADC can infer the Kubelet
  is not at a version with migration logic and should therefore fall-back to
  in-tree implementation
* If `CSINodeInfo` exists, and `storage.alpha.kubernetes.io/migrated-plugins`
  doesn't include the plugin name, then ADC can infer Kubelet has migration
  logic however the Feature Flag for that particular plugin is `off` and the ADC
  should therefore fall-back to in-tree storage implementation
* If `CSINodeInfo` exists, and `storage.alpha.kubernetes.io/migrated-plugins`
  does include the plugin name, then ADC can infer Kubelet has migration logic
  and the Feature Flag for that particular plugin is `on` and the ADC should
  therefore use the csi-plugin migration implementation
* If `CSINodeInfo` exists, and `storage.alpha.kubernetes.io/migrated-plugins`
  does include the plugin name but the ADC feature flags for that driver are off
  (`in-tree`), then an error should be thrown notifying users that Kubelet
  requested `csi-plugin` volume plugin mechanism but it was not specified on the
  ADC

In each of these above cases, the decision the ADC makes to use in-tree or csi
migration implemtnation will be mirror the Kubelets logic therefore guaranteeing
the entire lifecycle of a volume from controller to Kubelet will be done with
the same implementation.

### Node Drain Requirement

We require node's to be drained whenever the Kubelet is Upgrade/Downgraded or
Migrated/Unmigrated to ensure that the entire volume lifecycle is maintained
inside one code branch (CSI or In-tree). This simplifies upgrade/downgrade
significantly and reduces chance of error and races.

### Upgrade/Downgrade Migrate/Unmigrate Scenarios

For upgrade, starting from a non-migrated cluster you must turn on migration for
ADC first, then drain your node before turning on migration for the
Kubelet. The workflow is as follows:
1. ADC and Kubelet are both not migrated
2. ADC restarted and migrated (flags flipped)
3. ADC continues to use in-tree code for this node b/c
   `storage.alpha.kubernetes.io/migrated-plugins` does NOT include the plugin
   name
4. Node drained and made unschedulable. All volumes unmounted/detached with
   in-tree code
6. Kubelet restarted and migrated (flags flipped)
7. Kubelet updates CSINodeInfo node to tell ADC (without informer) whether each
   node/driver has been migrated by adding the plugin to the
   `storage.alpha.kubernetes.io/migrated-plugins` annotation
8. Kubelet is made schedulable
9. Both ADC & Kubelet Migrated, node is in "fresh" state so all new
   volumes lifecycle is CSI

For downgrade, starting from a fully migrated cluster you must drain your node
first, then turn off migration for your Kubelet, then turn off migration for the
ADC. The workflow is as follows:
1. ADC and Kubelet are both migrated
2. Kubelet drained and made unschedulable, all volumes unmounted/detached with
   CSI code
3. Kubelet restarted and un-migrated (flags flipped)
4. Kubelet removes the plugin in question to
   `storage.alpha.kubernetes.io/migrated-plugins`. In case kubelet does not have
   `storage.alpha.kubernetes.io/migrated-plugins` update code, admin must update
   the field manually.
5. Kubelet is made schedulable.
5. At this point all volumes going onto the node would be using in-tree code for
   both ADC(b/c of annotation) and Kublet
6. Restart and un-migrate ADC

With these workflows a volume attached with CSI will be handled by CSI code for
its entire lifecycle, and a volume attached with in-tree code will be handled by
in-tree code for its entire lifecycle.

## Cloud Provider Requirements

There is a push to remove CloudProvider code from kubernetes.

There will not be any general auto-deployment mechanism for ALL CSI drivers
covered in this document so the timeline to remove CloudProvider code using this
design is undetermined: For example: At some point GKE could auto-deploy the GCE
PD CSI driver and have migration for that turned on by default, however it may
not deploy any other drivers by default. And at this point we can only remove
the code for the GCE In-tree plugin (this would still break anyone doing their
own deployments while using GCE unless they install the GCE PD CSI Driver).

We could have auto-deploy depending on what cloud provider kubernetes is running
on. But AFAIK there is no standard mechanism to guarantee this on all Cloud
Providers.

For example the requirements for just the GCE Cloud Provider code for storage
with minimal disruption to users would be:
* In-tree to CSI Plugin migration goes GA
* GCE PD CSI Driver deployed on GCE/GKE by default (resource requirements of
  driver need to be determined)
* GCE PD CSI Migration turned on by default
* Remove in-tree plugin code and cloud provider code

And at this point users doing their own deployment and not installing the GCE PD
CSI driver encounter an error.

## Testing

### Migration Shim Testing
Run all existing in-tree plugin driver tests
* If migration is on for that plugin, add infrastructure piece that inspects CSI
  Drivers logs to make sure that the driver is servicing the operations
* Also observer that none of the in-tree code is being called

Additionally, we must test that a PV created from migrated dynamic provisioning
is identical to the PV created from the in-tree plugin

This should cover all use cases of volume operations, including volume
reconstruction. 

### Upgrade/Downgrade/Skew Testing
We need to have test clusters brought up that have different feature flags
enabled on different components (ADC and Kubelet). Once these feature flag skew
configurations are brought up the test itself would have to know what
configuration it’s running in and validate the expected result.

Configurations to test:

| ADC               | Kubelet                                            | Expected Result                                                          |
|-------------------|----------------------------------------------------|--------------------------------------------------------------------------|
| ADC Migration On  | Kubelet Migration On                               | Fully migrated - result should be same as “Migration Shim Testing” above |
| ADC Migration On  | Kubelet Migration Off (or Kubelet version too low) | No calls made to driver. All operations serviced by in-tree plugin       |
| ADC Migration Off | Kubelet Migration On                               | Not supported config - Undefined behavior                                |
| ADC Migration Off | Kubelet Migration Off                              | No calls made to driver. All operations service by in-tree plugin        |

### CSI Driver Feature Parity Testing

We will need some way to automatically qualify drivers have feature parity
before promoting their migration features to Beta (on by default). 

This is as simple as on the feature flags and run through our “Migration Shim
Testing” tests. If the driver passes all of them then they have parity. If not,
we need to revisit in-tree plugin tests and make sure they test the entire suite
of possible tests.
