# In-tree Storage Plugin to CSI Migration Design Doc

Authors: @davidz627, @jsafrane

This document presents a detailed design for migrating in-tree storage plugins
to CSI. This will be an opt-in feature turned on at cluster creation time that
will redirect in-tree plugin operations to a corresponding CSI Driver.

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
want to eventually transition all storage users to CSI.

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
* Manual installation of drivers

Beta [Target 1.15]
* On by default
* Migrate all of the cloud provider plugins*
* Plugin installation solved

GA [TBD]
* Feature on by default, per-plugin toggle on for relevant cloud provider by
  default
* CSI Drivers for migrated plugins available on related cloud provider cluster
  by default

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
  // TranslateToCSI takes a volume.Spec and will translate it to a
  // CSIPersistentVolumeSource if the translation logic for that
  // specific in-tree volume spec has been implemented
  TranslateToCSI(spec volume.Spec) (CSIPersistentVolumeSource, error)

  // TranslateToIntree takes a CSIPersistentVolumeSource and will translate
  // it to a volume.Spec for the specific in-tree volume specified by
  //`inTreePlugin`, if that translation logic has been implemented
  TranslateToInTree(source CSIPersistentVolumeSource, inTreePlugin string) (volume.Spec, error)

  // IsMigrated returns true if the plugin has migration logic
  // false if it does not
  IsMigrated(inTreePlugin string) bool
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


### In-Line Volumes
In-line controller calls are a special case because there is no PV. In this case
we will add the CSI Source JSON to the VolumeToAttach object and in Attach we
will put the Source in a new field in the VolumeAttachment object
VolumeAttachment.Spec.Source.VolumeAttachmentSource.InlineVolumeSource. The CSI Attacher will have to
be modified to also check this location for a source before checking the PV
itself.

We need to be careful with naming VolumeAttachments for in-line volumes. The
name needs to be unique and A/D controller must be able to find the right
VolumeAttachment when a pod is deleted (i.e. using only info in Node.Status).
CSI driver in kubelet must be able to find the VolumeAttachment too to get
AttachmentMetadata for NodeStage/NodePublish.

In downgrade scenario where the migration is then turned off we will have to
remove these floating VolumeAttachment objects, the same issue is outlined above
in the Non-Dynamic Provisioned Volumes section.

For more details on this see the PR that specs out CSI Inline Volumes in more detail:
https://github.com/kubernetes/community/pull/2273. Basically we will just translate
the in-tree inline volumes into the format specified/implemented in the 
container-storage-interface-inline-volumes proposal.

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


## Upgrade/Downgrade, Migrate/Un-migrate
### Kubelet Node Annotation
When the Kubelet starts, it will check whether the feature gate is
enabled and if so will annotate its node with `csi.attach.kubernetes.io/gce-pd`
for example to communicate to the A/D Controller that it supports migration of
the gce-pd to CSI. The A/D Controller will have to choose on a per-node basis
whether to use the CSI or the in-tree plugin for attach based on 3 criterea:
1. Feature gate
2. Plugin Migratable (Implements MigratablePlugin interface)
3. Node to Attach to has requisite Annotation

Note: All 3 criterea must be satisfied for A/D controller to Attach/Detach with
CSI instead of in-tree plugin. For example if a Kubelet has feature on and marks
the annotation, but the A/D Controller does not have the feature gate flipped,
we consider this user error and will throw some errors.

This can cause a race between the A/D Controller and the Kubelet annotating, if
a volume is attached before the Kubelet completes annotation the A/D controller
could attach using in-tree plugin instead of CSI while the Kubelet is expecting
a CSI Attach. The same issue exists on downgrade if the Annotation is not
removed before a volume is attached. An additional consideration is that we
cannot have the Kubelet downgraded to a version that does not have the
Annotation removal code.

### Node Drain Requirement
We require node's to be drained whenever the Kubelet is Upgrade/Downgraded or
Migrated/Unmigrated to ensure that the entire volume lifecycle is maintained
inside one code branch (CSI or In-tree). This simplifies upgrade/downgrade
significantly and reduces chance of error and races.

### Upgrade/Downgrade Migrate/Unmigrate Scenarios
For upgrade, starting from a non-migrated cluster you must turn on migration for
A/D Controller first, then drain your node before turning on migration for the
Kubelet. The workflow is as follows:
1. A/D Controller and Kubelet are both not migrated
2. A/D Controller restarted and migrated (flags flipped)
3. A/D Controller continues to use in-tree code for this node b/c node
   annotation doesn't exist
4. Node drained and made unschedulable. All volumes unmounted/detached with in-tree code
5. Kubelet restarted and migrated (flags flipped)
6. Kubelet annotates node to tell A/D controller this node has been migrated
7. Kubelet is made schedulable
8. Both A/D Controller & Kubelet Migrated, node is in "fresh" state so all new
   volumes lifecycle is CSI

For downgrade, starting from a fully migrated cluster you must drain your node
first, then turn off migration for your Kubelet, then turn off migration for the
A/D Controller. The workflow is as follows:
1. A/D Controller and Kubelet are both migrated
2. Kubelet drained and made unschedulable, all volumes unmounted/detached with CSI code
3. Kubelet restarted and un-migrated (flags flipped)
4. Kubelet removes node annotation to tell A/D Controller this node is not
   migrated. In case kubelet does not have annotation removal code, admin must
   remove the annotation manually.
5. Kubelet is made schedulable.
5. At this point all volumes going onto the node would be using in-tree code for
   both A/D Controller(b/c of annotation) and Kublet
6. Restart and un-migrate A/D Controller

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

## Plugin Availability (Installation)
When Migration goes Beta and is turned on by default we have the issue of how to
get the drivers on to the clusters. We have to consider what to do with both new
clusters, as well as existing clusters that may already be using storage. We
also have to consider what to do with both cloud storage solutions, and all the
other non-cloud specific drivers.

### New Clusters (Cloud Drivers)
For new clusters we can modify the `cluster-up` scripts to deploy the cloud
storage drivers automatically with cluster bring-up depending on which cloud the
cluster is being brought up on. For example if we `cluster-up` on GCE we would
just install the GCE PD CSI Driver by default. Each Cloud Provider can also
provide their own mechanism adding drivers in their own managed solutions.

### Upgraded Clusters (Cloud Drivers)
For upgraded clusters the respective cloud drivers should be installed on
upgrade. If there is no good universal hook to install things onto a cluster on
upgrade then we fallback to the behavior described in the next section "Missing
Drivers."

### Non-Cloud and Missing Drivers
For non-cloud and other miscellaneous missing drivers the question of
installation becomes more nuanced. We have to strike a balance between
unexpected resource consumption (installing everything would be unnecessary and
wasteful) and user grief (in-tree drivers no longer work because CSI Drivers are
not installed). There are a couple options we could take here:

#### Automatic Installation
One option is to create a tool that installs CSI Drivers automatically based on
storage classes present. Since the storage classes for in-tree plugins have
"well known" provisioner names we could concieve of a tool that scans all
storage classes in the system and installs the drivers present in the storage
classes on the cluster. This tool could potentially also be run as a controller
that watches storage classes that acts whenever a storage class is updated,
created, or deleted.

Pros:
* All "used" drivers are installed
* All "unused" drivers can be uninstalled
* Set it and forget it
* Does not affect user workloads (all the work and changes happen on the admin
  side when storage classes change)

Cons:
* Extra controller watching the Storage Class API
* Tool has a dependency on CSI Driver Registry to know what drivers are already
  installed

Another option is to install the driver on first use of the plugin, when we try
to use an in-tree plugin we check for existence of the driver and install it if
it doesn't exist. Pros:
* No manual intervention necessary
* Drivers installed if and only if they are necessary

Cons:
* Hard to tell whether driver exists or just in transient failure
* Resource usage inscreases because of driver deployment at unexpected time
  (attach of a volume)
* Latency for first usage of in-tree driver


#### Manual Installation
A manual installation option (that can also be used as fallback if automatic
installation fails or is not done) also requires some choosing. The main concern
here is that the controllers should not panic when the drivers do not exist or
are not responding. The ideal situation is that they print out a nice error
message such as "Your administrator has turned on the CSI Migration feature, in
order to continue using this plugin please contact your administrator to install
the CSI DRIVER (LINK)." This could be enforced as a validating admission webhook
for new objects, as well as in the operation generator layer for existing
objects or ones that were not caught by the admission validation (we have CSI
Driver Registration objects now to see what drivers are installed).

#### Installation Conclusions
The problem with using a Manual installation option only is that this "breaks
the in-tree API" that we are trying our hardest not to break. The Manual
Installation option should be implemented as a backup but should only be
encountered in a situation where the Automatic Installation method has either
failed or did not run. Based on the fact that we want to minimize user
disruption, it seems that the best Automatic Installation method to use is the
installation tool that runs as a controller.

## Testing
### Standard
Good news is that all “normal functionality” can be tested by simply bringing up
a cluster with “migrated” drivers and running the existing e2e tests for that
driver. We will create CI jobs that run in this configuration for each new
volume plugin

### Migration/Non-migration (Upgrade/Downgrade)
Write tests were in a normal workflow of attach/mount/unmount/detach, we have
any one of these operations actually happen with the old volume plugin, not the
CSI one This makes sure that the workflow is resiliant to rollback at any point
in time.

### Version Skew
Master/Node can have up to 2 version skw. Master must always be equal or higher
version than the node. It should be covered by the tests in
Migration/Non-migration section.
