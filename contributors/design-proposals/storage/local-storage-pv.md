# Local Storage Persistent Volumes

Authors: @msau42, @vishh, @dhirajh, @ianchakeres

This document presents a detailed design for supporting persistent local storage,
as outlined in [Local Storage Overview](local-storage-overview.md).
Supporting all the use cases for persistent local storage will take many releases,
so this document will be extended for each new release as we add more features.

## Goals

* Allow pods to mount any local block or filesystem based volume.
* Allow pods to mount dedicated local disks, or channeled partitions as volumes for
IOPS isolation.
* Allow pods do access local volumes without root privileges.
* Allow pods to access local volumes without needing to understand the storage
layout on every node.
* Persist local volumes and provide data gravity for pods.  Any pod
using the local volume will be scheduled to the same node that the local volume
is on.
* Allow pods to release their local volume bindings and lose that volume's data
during failure conditions, such as node, storage or scheduling failures, where
the volume is not accessible for some user-configurable time.
* Allow pods to specify local storage as part of a Deployment or StatefulSet.
* Allow administrators to set up and configure local volumes with simple methods.
* Do not require administrators to manage the local volumes once provisioned
for a node.

## Non-Goals

* Provide data availability for a local volume beyond its local node.
* Support the use of HostPath volumes and Local PVs on the same volume.

## Background

In Kubernetes, there are two main types of storage: remote and local.

Remote storage is typically used with persistent volumes where the data can
persist beyond the lifetime of the pod.

Local storage is typically used with ephemeral volumes where the data only
persists during the lifetime of the pod.

There is increasing demand for using local storage as persistent volumes,
especially for distributed filesystems and databases such as GlusterFS and
Cassandra.  The main motivations for using persistent local storage, instead
of persistent remote storage include:

* Performance:  Local SSDs achieve higher IOPS and throughput than many
remote storage solutions.

* Cost: Operational costs may be reduced by leveraging existing local storage,
especially in bare metal environments.  Network storage can be expensive to
setup and maintain, and it may not be necessary for certain applications.

## Use Cases

### Distributed filesystems and databases

Many distributed filesystem and database implementations, such as Cassandra and
GlusterFS, utilize the local storage on each node to form a storage cluster.
These systems typically have a replication feature that sends copies of the data
to other nodes in the cluster in order to provide fault tolerance in case of
node failures.  Non-distributed, but replicated databases, like MySQL, can also
utilize local storage to store replicas.

The main motivations for using local persistent storage are performance and
cost.  Since the application handles data replication and fault tolerance, these
application pods do not need networked storage to provide shared access to data.
In addition, installing a high-performing NAS or SAN solution can be more
expensive, and more complex to configure and maintain than utilizing local
disks, especially if the node was already pre-installed with disks.  Datacenter
infrastructure and operational costs can be reduced by increasing storage
utilization.

These distributed systems are generally stateful, infrastructure applications
that provide data services to higher-level applications.  They are expected to
run in a cluster with many other applications potentially sharing the same
nodes.  Therefore, they expect to have high priority and node resource
guarantees.  They typically are deployed using StatefulSets, custom
controllers, or operators.

### Caching

Caching is one of the recommended use cases for ephemeral local storage.  The
cached data is backed by persistent storage, so local storage data durability is
not required.  However, there is a use case for persistent local storage to
achieve data gravity for large caches.  For large caches, if a pod restarts,
rebuilding the cache can take a long time.  As an example, rebuilding a 100GB
cache from a hard disk with 150MB/s read throughput can take around 10 minutes.
If the service gets restarted and all the pods have to restart, then performance
and availability can be impacted while the pods are rebuilding.  If the cache is
persisted, then cold startup latencies are reduced.

Content-serving applications and producer/consumer workflows commonly utilize
caches for better performance.  They are typically deployed using Deployments,
and could be isolated in its own cluster, or shared with other applications.

## Environments

### Baremetal

In a baremetal environment, nodes may be configured with multiple local disks of
varying capacity, speeds and mediums.  Mediums include spinning disks (HDDs) and
solid-state drives (SSDs), and capacities of each disk can range from hundreds
of GBs to tens of TB. Multiple disks may be arranged in JBOD or RAID configurations 
to consume as persistent storage.

Currently, the methods to use the additional disks are to:

* Configure a distributed filesystem
* Configure a HostPath volume

It is also possible to configure a NAS or SAN on a node as well.  Speeds and
capacities will widely vary depending on the solution.

### GCE/GKE

GCE and GKE both have a local SSD feature that can create a VM instance with up
to 8 fixed-size 375GB local SSDs physically attached to the instance host and
appears as additional disks in the instance.  The local SSDs have to be
configured at the VM creation time and cannot be dynamically attached to an
instance later.  If the VM gets shutdown, terminated, pre-empted, or the host
encounters a non-recoverable error, then the SSD data will be lost.  If the
guest OS reboots, or a live migration occurs, then the SSD data will be
preserved.

### EC2

In EC2, the instance store feature attaches local HDDs or SSDs to a new instance
as additional disks.  HDD capacities can go up to 24 2TB disks for the largest
configuration.  SSD capacities can go up to 8 800GB disks or 2 2TB disks for the
largest configurations.  Data on the instance store only persists across
instance reboot.

## Limitations of current volumes

The following is an overview of existing volume types in Kubernetes, and how
they cannot completely address the use cases for local persistent storage.

* EmptyDir: A temporary directory for a pod that is created under the kubelet
root directory.  The contents are deleted when a pod dies.  Limitations:

  * Volume lifetime is bound to the pod lifetime.  Pod failure is more likely
than node failure, so there can be increased network and storage activity to
recover data via replication and data backups when a replacement pod is started.
  * Multiple disks are not supported unless the administrator aggregates them
into a spanned or RAID volume.  In this case, all the storage is shared, and
IOPS guarantees cannot be provided.
  * There is currently no method of distinguishing between HDDs and SDDs.  The
“medium” field could be expanded, but it is not easily generalizable to
arbitrary types of mediums.

* HostPath: A direct mapping to a specified directory on the node.  The
directory is not managed by the cluster.  Limitations:

  * Admin needs to manually setup directory permissions for the volume’s users.
  * Admin has to manage the volume lifecycle manually and do cleanup of the data and
directories.
  * All nodes have to have their local storage provisioned the same way in order to
use the same pod template.
  * There can be path collision issues if multiple pods get scheduled to the same
node that want the same path
  * If node affinity is specified, then the user has to do the pod scheduling
manually.

* Provider’s block storage (GCE PD, AWS EBS, etc): A remote disk that can be
attached to a VM instance.  The disk’s lifetime is independent of the pod’s
lifetime.  Limitations:

  * Doesn’t meet performance requirements.
[Performance benchmarks on GCE](https://cloud.google.com/compute/docs/disks/performance)
show that local SSD can perform better than SSD persistent disks:

    * 16x read IOPS
    * 11x write IOPS
    * 6.5x read throughput
    * 4.5x write throughput

* Networked filesystems (NFS, GlusterFS, etc): A filesystem reachable over the
network that can provide shared access to data.  Limitations:

  * Requires more configuration and setup, which adds operational burden and
cost.
  * Requires a high performance network to achieve equivalent performance as
local disks, especially when compared to high-performance SSDs.

Due to the current limitations in the existing volume types, a new method for
providing persistent local storage should be considered.

## Feature Plan

A detailed implementation plan can be found in the
[Storage SIG planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/view#gid=1566770776).
The following is a high level summary of the goals in each phase.

### Phase 1

* Support Pod, Deployment, and StatefulSet requesting a single local volume
* Support pre-configured, statically partitioned, filesystem-based local volumes

### Phase 2

* Block devices and raw partitions
* Smarter PV binding to consider local storage and pod scheduling constraints,
such as pod affinity/anti-affinity, and requesting multiple local volumes

### Phase 3

* Support common partitioning patterns
* Volume taints and tolerations for unbinding volumes in error conditions

### Phase 4

* Dynamic provisioning

## Design

A high level proposal with user workflows is available in the
[Local Storage Overview](local-storage-overview.md).

This design section will focus on one phase at a time.  Each new release will
extend this section.

### Phase 1: 1.7 alpha

#### Local Volume Plugin

A new volume plugin will be introduced to represent logical block partitions and
filesystem mounts that are local to a node.  Some examples include whole disks,
disk partitions, RAID volumes, LVM volumes, or even directories in a shared
partition.  Multiple Local volumes can be created on a node, and is
accessed through a local mount point or path that is bind-mounted into the
container.  It is only consumable as a PersistentVolumeSource because the PV
interface solves the pod spec portability problem and provides the following:

* Abstracts volume implementation details for the pod and expresses volume
requirements in terms of general concepts, like capacity and class.  This allows
for portable configuration, as the pod is not tied to specific volume instances.
* Allows volume management to be independent of the pod lifecycle.  The volume can
survive container, pod and node restarts.
* Allows volume classification by StorageClass.
* Is uniquely identifiable within a cluster and is managed from a cluster-wide
view.

There are major changes in PV and pod semantics when using Local volumes
compared to the typical remote storage volumes.

* Since Local volumes are fixed to a node, a pod using that volume has to
always be scheduled on that node.
* Volume availability is tied to the node’s availability.  If the node is
unavailable, then the volume is also unavailable, which impacts pod
availability.
* The volume’s data durability characteristics are determined by the underlying
storage system, and cannot be guaranteed by the plugin.  A Local volume
in one environment can provide data durability, but in another environment may
only be ephemeral.  As an example, in the GCE/GKE/AWS cloud environments, the
data in directly attached, physical SSDs is immediately deleted when the VM
instance terminates or becomes unavailable.

Due to these differences in behaviors, Local volumes are not suitable for
general purpose use cases, and are only suitable for specific applications that
need storage performance and data gravity, and can tolerate data loss or
unavailability.  Applications need to be aware of, and be able to handle these
differences in data durability and availability.

Local volumes are similar to HostPath volumes in the following ways:

* Partitions need to be configured by the storage administrator beforehand.
* Volume is referenced by the path to the partition.
* Provides the same underlying partition’s support for IOPS isolation.
* Volume is permanently attached to one node.
* Volume can be mounted by multiple pods on the same node.

However, Local volumes will address these current issues with HostPath
volumes:

* Security concerns allowing a pod to access any path in a node.  Local
volumes cannot be consumed directly by a pod.  They must be specified as a PV
source, so only users with storage provisioning privileges can determine which
paths on a node are available for consumption.
* Difficulty in permissions setup.  Local volumes will support fsGroup so
that the admins do not need to setup the permissions beforehand, tying that
particular volume to a specific user/group.  During the mount, the fsGroup
settings will be applied on the path.  However, multiple pods
using the same volume should use the same fsGroup.
* Volume lifecycle is not clearly defined, and the volume has to be manually
cleaned up by users.  For Local volumes, the PV has a clearly defined
lifecycle.  Upon PVC deletion, the PV will be released (if it has the Delete
policy), and all the contents under the path will be deleted.  In the future,
advanced cleanup options, like zeroing can also be specified for a more
comprehensive cleanup.

##### API Changes

All new changes are protected by a new feature gate, `PersistentLocalVolumes`.

A new `LocalVolumeSource` type is added as a `PersistentVolumeSource`.  For this
initial phase, the path can only be a mount point or a directory in a shared
filesystem.

```
type LocalVolumeSource struct {
        // The full path to the volume on the node
        // For alpha, this path must be a directory
        // Once block as a source is supported, then this path can point to a block device
        Path string
}

type PersistentVolumeSource struct {
    <snip>
    // Local represents directly-attached storage with node affinity.
    // +optional
    Local *LocalVolumeSource
}
```

The relationship between a Local volume and its node will be expressed using
PersistentVolume node affinity, described in the following section.

Users request Local volumes using PersistentVolumeClaims in the same manner as any
other volume type. The PVC will bind to a matching PV with the appropriate capacity,
AccessMode, and StorageClassName.  Then the user specifies that PVC in their
Pod spec.  There are no special annotations or fields that need to be set in the Pod
or PVC to distinguish between local and remote storage.  It is abstracted by the
StorageClass.

#### PersistentVolume Node Affinity

PersistentVolume node affinity is a new concept and is similar to Pod node affinity,
except instead of specifying which nodes a Pod has to be scheduled to, it specifies which nodes
a PersistentVolume can be attached and mounted to, influencing scheduling of Pods that
use local volumes.

For a Pod that uses a PV with node affinity, a new scheduler predicate
will evaluate that node affinity against the node's labels.  For this initial phase, the
PV node affinity is only considered by the scheduler for already-bound PVs.  It is not
considered during the initial PVC/PV binding, which will be addressed in a future release.

Only the `requiredDuringSchedulingIgnoredDuringExecution` field will be supported.

##### API Changes

For the initial alpha phase, node affinity is expressed as an optional
annotation in the PersistentVolume object.

```
// AlphaStorageNodeAffinityAnnotation defines node affinity policies for a PersistentVolume.
// Value is a string of the json representation of type NodeAffinity
AlphaStorageNodeAffinityAnnotation = "volume.alpha.kubernetes.io/node-affinity"
```

#### Local volume initial configuration

There are countless ways to configure local storage on a node, with different patterns to
follow depending on application requirements and use cases.  Some use cases may require
dedicated disks; others may only need small partitions and are ok with sharing disks.
Instead of forcing a partitioning scheme on storage administrators, the Local volume
is represented by a path, and lets the administrators partition their storage however they
like, with a few minimum requirements:

* The paths to the mount points are always consistent, even across reboots or when storage
is added or removed.
* The paths are backed by a filesystem (block devices or raw partitions are not supported for
the first phase)
* The directories have appropriate permissions for the provisioner to be able to set owners and
cleanup the volume.

#### Local volume management

Local PVs are statically created and not dynamically provisioned for the first phase.
To mitigate the amount of time an administrator has to spend managing Local volumes,
a Local static provisioner application will be provided to handle common scenarios.  For
uncommon scenarios, a specialized provisioner can be written.

The Local static provisioner will be developed in the
[kubernetes-incubator/external-storage](https://github.com/kubernetes-incubator)
repository, and will loosely follow the external provisioner design, with a few differences:

* A provisioner instance needs to run on each node and only manage the local storage on its node.
* For phase 1, it does not handle dynamic provisioning.  Instead, it performs static provisioning
by discovering available partitions mounted under configurable discovery directories.

The basic design of the provisioner will have two separate handlers: one for PV deletion and
cleanup, and the other for static PV creation.  A PersistentVolume informer will be created
and its cache will be used by both handlers.

PV deletion will operate on the Update event.  If the PV it provisioned changes to the “Released”
state, and if the reclaim policy is Delete, then it will cleanup the volume and then delete the PV,
removing it from the cache.

PV creation does not operate on any informer events.  Instead, it periodically monitors the discovery
directories, and will create a new PV for each path in the directory that is not in the PV cache.  It
sets the "pv.kubernetes.io/provisioned-by" annotation so that it can distinguish which PVs it created.

For phase 1, the allowed discovery file types are directories and mount points.  The PV capacity
will be the capacity of the underlying filesystem.  Therefore, PVs that are backed by shared
directories will report its capacity as the entire filesystem, potentially causing overcommittment.
Separate partitions are recommended for capacity isolation.

The name of the PV needs to be unique across the cluster.  The provisioner will hash the node name,
StorageClass name, and base file name in the volume path to generate a unique name.

##### Packaging

The provisioner is packaged as a container image and will run on each node in the cluster as part of
a DaemonSet.  It needs to be run with a user or service account with the following permissions:

* Create/delete/list/get PersistentVolumes - Can use the `system:persistentvolumeprovisioner` ClusterRoleBinding
* Get ConfigMaps - To access user configuration for the provisioner
* Get Nodes - To get the node's UID and labels

These are broader permissions than necessary (a node's access to PVs should be restricted to only
those local to the node).  A redesign will be considered in a future release to address this issue.

In addition, it should run with high priority so that it can reliably handle all the local storage
partitions on each node, and with enough permissions to be able to cleanup volume contents upon
deletion.

The provisioner DaemonSet requires the following configuration:

* The node's name set as the MY_NODE_NAME environment variable
* ConfigMap with StorageClass -> discovery directory mappings
* Each mapping in the ConfigMap needs a hostPath volume
* User/service account with all the required permissions

Here is an example ConfigMap:

```
kind: ConfigMap
metadata:
  name: local-volume-config
  namespace: kube-system
data:
  storageClassMap: |
    local-fast:
      hostDir: "/mnt/ssds"
      mountDir: "/local-ssds"
    local-slow:
      hostDir: "/mnt/hdds"
      mountDir: "/local-hdds"
```

The `hostDir` is the discovery path on the host, and the `mountDir` is the path it is mounted to in
the provisioner container.  The `hostDir` is required because the provisioner needs to create Local PVs
with the `Path` based off of `hostDir`, not `mountDir`.

The DaemonSet for this example looks like:
```

apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: local-storage-provisioner
  namespace: kube-system
spec:
  template:
    metadata:
      labels:
        system: local-storage-provisioner
    spec:
      containers:
      - name: provisioner
        image: "k8s.gcr.io/local-storage-provisioner:v1.0"
        imagePullPolicy: Always
        volumeMounts:
        - name: vol1
          mountPath: "/local-ssds"
        - name: vol2
          mountPath: "/local-hdds"
        env:
        - name: MY_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
      volumes:
      - name: vol1
        hostPath:
          path: "/mnt/ssds"
      - name: vol2
        hostPath:
          path: "/mnt/hdds"
      serviceAccount: local-storage-admin
```

##### Provisioner Boostrapper

Manually setting up this DaemonSet spec can be tedious and it requires duplicate specification
of the StorageClass -> directory mappings both in the ConfigMap and as hostPath volumes. To
make it simpler and less error prone, a boostrapper application will be provided to generate
and launch the provisioner DaemonSet based off of the ConfigMap.  It can also create a service
account with all the required permissions.

The boostrapper accepts the following optional arguments:

* -image: Name of local volume provisioner image (default
"quay.io/external_storage/local-volume-provisioner:latest")
* -volume-config: Name of the local volume configuration configmap. The configmap must reside in the same
namespace as the bootstrapper. (default "local-volume-default-config")
* -serviceaccount: Name of the service account for local volume provisioner (default "local-storage-admin")

The boostrapper requires the following permissions:

* Get/Create/Update ConfigMap
* Create ServiceAccount
* Create ClusterRoleBindings
* Create DaemonSet

Since the boostrapper generates the DaemonSet spec, the ConfigMap can be simplified to just specify the
host directories:

```
kind: ConfigMap
metadata:
  name: local-volume-config
  namespace: kube-system
data:
 storageClassMap: |
    local-fast:
      hostDir: "/mnt/ssds"
    local-slow:
      hostDir: "/mnt/hdds"
```

The boostrapper will update the ConfigMap with the generated `mountDir`.  It generates the `mountDir`
by stripping off the initial "/" in `hostDir`, replacing the remaining "/" with "~", and adding the
prefix path "/mnt/local-storage".

In the above example, the generated `mountDir` is `/mnt/local-storage/mnt ~ssds` and
`/mnt/local-storage/mnt~hdds`, respectively.

#### Use Case Deliverables

This alpha phase for Local PV support will provide the following capabilities:

* Local directories to be specified as Local PVs with node affinity
* Pod using a PVC that is bound to a Local PV will always be scheduled to that node
* External static provisioner DaemonSet that discovers local directories, creates, cleans up,
and deletes Local PVs

#### Limitations

However, some use cases will not work:

* Specifying multiple Local PVCs in a pod.  Most likely, the PVCs will be bound to Local PVs on
different nodes,
making the pod unschedulable.
* Specifying Pod affinity/anti-affinity with Local PVs.  PVC binding does not look at Pod scheduling
constraints at all.
* Using Local PVs in a highly utilized cluster.  PVC binding does not look at Pod resource requirements
and Node resource availability.

These issues will be solved in a future release with advanced storage topology scheduling.

As a workaround, PVCs can be manually prebound to Local PVs to essentially manually schedule Pods to
specific nodes.

#### Test Cases

##### API unit tests

* LocalVolumeSource cannot be specified without the feature gate
* Non-empty PV node affinity is required for LocalVolumeSource
* Preferred node affinity is not allowed
* Path is required to be non-empty
* Invalid json representation of type NodeAffinity returns error

##### PV node affinity unit tests

* Nil or empty node affinity evaluates to true for any node
* Node affinity specifying existing node labels evaluates to true
* Node affinity specifying non-existing node label keys evaluates to false
* Node affinity specifying non-existing node label values evaluates to false

##### Local volume plugin unit tests

* Plugin can support PersistentVolumeSource
* Plugin cannot support VolumeSource
* Plugin supports ReadWriteOnce access mode
* Plugin does not support remaining access modes
* Plugin supports Mounter and Unmounter
* Plugin does not support Provisioner, Recycler, Deleter
* Plugin supports readonly
* Plugin GetVolumeName() returns PV name
* Plugin ConstructVolumeSpec() returns PV info
* Plugin disallows backsteps in the Path

##### Local volume provisioner unit tests

* Directory not in the cache and PV should be created
* Directory is in the cache and PV should not be created
* Directories created later are discovered and PV is created
* Unconfigured directories are ignored
* PVs are created with the configured StorageClass
* PV name generation hashed correctly using node name, storageclass and filename
* PV creation failure should not add directory to cache
* Non-directory type should not create a PV
* PV is released, PV should be deleted
* PV should not be deleted for any other PV phase
* PV deletion failure should not remove PV from cache
* PV cleanup failure should not delete PV or remove from cache

##### E2E tests

* Pod that is bound to a Local PV is scheduled to the correct node
and can mount, read, and write
* Two pods serially accessing the same Local PV can mount, read, and write
* Two pods simultaneously accessing the same Local PV can mount, read, and write
* Test both directory-based Local PV, and mount point-based Local PV
* Launch local volume provisioner, create some directories under the discovery path,
and verify that PVs are created and a Pod can mount, read, and write.
* After destroying a PVC managed by the local volume provisioner, it should cleanup
the volume and recreate a new PV.
* Pod using a Local PV with non-existent path fails to mount
* Pod that sets nodeName to a different node than the PV node affinity cannot schedule.


### Phase 2: 1.9 alpha

#### Smarter PV binding

The issue of PV binding not taking into account pod scheduling requirements affects any
type of volume that imposes topology constraints, such as local storage and zonal disks.

Because this problem affects more than just local volumes, it will be treated as a
separate feature with a separate proposal.  Once that feature is implemented, then the
limitations outlined above will be fixed.

#### Block devices and raw partitions

Pods accessing raw block storage is a new alpha feature in 1.9.  Changes are required in
the Local volume plugin and provisioner to be able to support raw block devices. The local
volume provisioner will be enhanced to support discovery of block devices and creation of
PVs corresponding to those block devices. In addition, when a block device based PV is
released, the local volume provisioner will cleanup the block devices. The cleanup
mechanism will be configurable and also customizable as no single mechanism covers all use
cases.


##### Discovery

Much like the current file based PVs, the local volume provisioner will look for block devices
under designated directories that have been mounted on the provisioner container. Currently, for
each storage class, the provisioner has a configmap entry that looks like this:

```
data:
  storageClassMap: |
    local-fast:
      hostDir: "/mnt/disks"
      mountDir: "/local-ssds"
```

With this current approach, filesystems that were meant to be exposed as PVs are supposed to be
mounted on sub-directories under hostDir and the provisioner running in a container would walk
through the corresponding "mountDir" to find all the PVs.  

For block discovery, we will extend the same approach to enable discovering block devices. The
admin can create symbolic links under hostDir for each block device that should be discovered
under that storage class. The provisioner would use the same configMap and its logic will be
enhanced to auto detect if the entry under the directory is a block device or a file system. If
it is a block device, then a block based PV is created, otherwise a file based PV is created.

##### Cleanup after Release

Cleanup of a block device can be a bit more involved for the following reasons:

* With file based PVs, a quick deletion of all files (inode information) was sufficient, with
block devices one might want to wipe all current content.
* Overwriting SSDs is not guaranteed to securely cleanup all previous content as there is a
layer of indirection in SSDs called the FTL (flash translation layer) and also wear leveling
techniques in SSDs that prevent reliable overwrite of all previous content. 
* SSDs can also suffer from wear if they are repeatedly subjected to zeroing out, so one would
need different tools and strategies for HDDs vs SSDs
* A cleanup process which favors overwriting every block in the disk can take several hours.

For this reason, the cleanup process has been made configurable and extensible, so that admin
can use the most appropriate method for their environment.
 
Block device cleanup logic will be encapsulated in separate scripts or binaries. There will be
several scripts that will be made available out of the box, for example:


| Cleanup Method | Description | Suitable for Device |
|:--------------:|-------------|:-------------------:|
|dd-zero| Used for zeroing the device repeatedly | HDD |
|blkdiscard| Discards sectors on the device. This cleanup method may not be supported by all devices.| SSD |
|fs-reset| A non-secure overwrite of any existing filesystem with mkfs, followed by wipefs to remove the signature of the file system | SSD/HDD |
|shred|Repeatedly writes random values to the block device. Less effective with wear levelling in SSDs.| HDD |
| hdparm| Issues [ATA secure erase](https://ata.wiki.kernel.org/index.php/ATA_Secure_Erase) command to erase data on device. See ATA Secure Erase. Please note that the utility has to be supported by the device in question. | SSD/HDD |

The fs-reset method is a quick and minimal approach as it does a reset of any file system, which
works for both SSD and HDD and will be the default choice for cleaning. For SSDs, admins could
opt for either blkdiscard which is also quite fast or hdparm. For HDDs they could opt for
dd-zeroing or shred, which can take some time to run. Finally, the user is free to create new
cleanup scripts of their own and have them specified in the configmap of the provisioner.

The configmap from earlier section will be enhanced as follows
```
data:
  storageClassMap: |
    local-fast:
      hostDir: "/mnt/disks"
      mountDir: "/local-ssds"
      blockCleanerCommand:
         - "/scripts/dd_zero.sh"
         - "2"
 ```

The block cleaner command will specify the script and any arguments that need to be passed to it.
The actual block device being cleaned will be supplied to the script as an environment variable
(LOCAL_PV_BLKDEVICE) as opposed to command line, so that the script command line has complete
freedom on its structure. The provisioner will validate that the block device path is actually
within the directory managed by the provisioner, to prevent destructive operations on arbitrary
paths.

The provisioner logic currently does each volume’s cleanup as a synchronous serial activity.
However, with cleanup now potentially being a multi hour activity, the processes will have to
be asynchronous and capable of being executed in parallel. The provisioner will ensure that all
current asynchronous cleanup processes are tracked. Special care needs to be taken to ensure that
when a disk has only been partially cleaned. This scenario can happen if some impatient user
manually deletes a PV and the provisioner ends up re-creating pv ready for use (but only partially
cleaned). This issue will be addressed in the re-design of the provisioner (details will be provided
in the re-design section). The re-design will ensure that all disks being cleaned will be tracked
through custom resources, so no disk being cleaned will be re-created as a PV.

The provisioner will also log events to let the user know that cleaning is in progress and it can
take some time to complete.

##### Testing

The unit tests in the provisioner will be enhanced to test all the new block discover, block cleaning
and asynchronous cleaning logic. The tests include
* Validating that a discovery directory containing both block and file system volumes are appropriately discovered and have PVs created.
* Validate that both success and failure of asynchronous cleanup processes are properly tracked by the provisioner
* Ensure a new PV is not created while cleaning of volume behind the PV is still in progress
* Ensure two simultaneous cleaning operations on the same PV do not occur

 In addition, end to end tests will be added to support block cleaning. The tests include:
* Validate block PV are discovered and created
* Validate cleaning of released block PV using each of the block cleaning scripts included.
* Validate that file and block volumes in the same discovery path have correct PVs created, and that they are appropriately cleaned up.
* Leverage block PV via PVC and validate that serially writes data in one pod, then reads and validates the data from a second pod.
* Restart of the provisioner during cleaning operations, and validate that the PV is not recreated by the provisioner until cleaning has occurred.

#### Provisioner redesign for stricter K8s API access control

In 1.7, each instance of the provisioner on each node has full permissions to create and
delete all PVs in the system.  This is unnecessary and potentially a vulnerability if the
node gets compromised.

To address this issue, the provisioner will be redesigned into two major components:

1. A central manager pod that handles the creation and deletion of PV objects.
This central pod can run on a trusted node and be given PV create/delete permissions.
2. Worker pods on each node, run as a DaemonSet, that discovers and cleans up the local
volumes on that node.  These workers do not interact with PV objects, however
they still require permissions to be able to read the `Node.Labels` on their node.

The central manager will poll each worker for their discovered volumes and create PVs for
them.  When a PV is released, then it will send the cleanup request to the worker.

Detailed design TBD
