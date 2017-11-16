## Kubernetes Volume Package Technical Overview

Latest active release branch as of today, 3/13/2017: [Kubernetes Release
1.6](https://github.com/kubernetes/kubernetes/tree/release-1.6)

The Kubernetes [volume
package](https://github.com/kubernetes/kubernetes/tree/master/pkg/volume) is a
modular implementation of storage integrations for storing or accessing
application data. Stores may be R/W or read-only. It consists of a core set of
interface definitions used throughout the rest of the Kubernetes application,
and a set of plugins that implement these interfaces to support providers such
as Google Persistent Volumes, Amazon Elastic Block Storage, GlusterFS, and NFS.
Volumes are necessary to run a stateful application such as a database on a
Kubernetes cluster, and a volume plugin must be implemented per supported block
storage provider. Components that depend on this package should be agnostic to
the underlying storage provider and rely solely on the generic interfaces
described in
[`pkg/volume/volume.go`](https://github.com/kubernetes/kubernetes/blob/master/pkg/volume/volume.go)
and
[`pkg/volume/plugins.go`](https://github.com/kubernetes/kubernetes/blob/master/pkg/volume/plugins.go).

### Volume Data Structures
```
├── VolumeOptions
├── VolumePluginMgr
├── Spec
├── VolumeConfig
├── Metrics
└── Attributes
```

##### `VolumeOptions`

`VolumeOptions` provides a generic way to specify mounting options for a
particular volume (e.g., one EBS instance) across platforms and clusters. It
holds metadata such as cluster-wide unique generated names, cloud tags,
parameters, cluster name, and the
[`PersistentVolumeClaim`](https://github.com/kubernetes/kubernetes/blob/f18a921a0308c4ecd2e7e8b7e30e39a8216a3447/pkg/api/v1/types.go#L513).
This is ultimately consumed by a method in [`VolumePlugin`](#volumeplugin)
implementations that creates a [`Mounter`](#mounter).

##### `VolumePluginMgr`

`VolumePluginMgr` is the managing data structure for all of a kubelet's volume
plugins. It maintains a map of unique plugin names to instances of their
implementing data structures. Each entry is tied to a particular volume (e.g.,
one EBS instance or one GPD instance). It implements an initialization method
and a variety of find-[`VolumePlugin`](#volumeplugin)-by-property methods. Each
kubelet instantiates an instance of `VolumePluginMgr` in its [volume
manager](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/volumemanager/volume_manager.go)
to keep track of all the requested volumes and their states. The volume manager
provides additional methods to find volumes by pods and to handle operations
such as mounting and unmounting.

##### `Spec`

The `Spec` is the internal representation of volumes that all APIs must be able
to convert to. Its fields include types of `*v1.Volume` and
`*v1.PersistentVolume` specified in
[pkg/api/v1/types.go](https://github.com/kubernetes/kubernetes/blob/release-1.6/pkg/api/v1/types.go#L223).

This file must be updated to support a new volume plugin, specifically by
adding a new struct named `YourPluginVolumeSource` containing the data
necessary to mount it. This new struct must then be added to the `VolumeSource`
definition and potentially the `PersistentVolumeSource` definition. See the
existing implementations for details.


##### `VolumeConfig`

`VolumeConfig` is used to provide configuration to volumes. It is different
from [`VolumeOptions`](#volumeoptions) in that `VolumeOptions` are used to
specify the provisioning, mounting, and unmounting of volumes, while
`VolumeConfig` can specify more "runtime" behavior such as [recycler
policy](https://kubernetes.io/docs/user-guide/persistent-volumes/#recycling).
It is currently only used by a few plugins.

##### `Metrics`

The `Metrics` data type is used to record general storage device metrics such
as total size, available space, inodes used, and more. It is tightly associated
with the [`MetricsProvider`](#metricsprovider) interface.

##### `Attributes`

The `Attributes` data type is used to describe further behaviors of the
[`Mounter`](#mounter) interface such as read-only.

### Volume Interfaces

This is a inheritance model of interfaces found in the volume package, with
descendents embedding their predecessesors.

```

├── VolumePlugin
│   ├── PersistentVolumePlugin
│   ├── RecyclableVolumePlugin
│   ├── DeletableVolumePlugin
│   ├── ProvisionableVolumePlugin
│   └── AttachableVolumePlugin
├── VolumeHost
├── MetricsProvider
│   └── Volume
│       ├── Mounter
│       ├── Unmounter
│       └── Deleter
├── Provisioner
├── Attacher
├── BulkVolumeVerifier
└── Detacher
```
#### `VolumePlugin`

`VolumePlugin` is the main interface that must be implemented by all volume
plugins. It is embedded in several other interfaces that provide special
functionality atop of it, as seen in the tree above. This design gives the
implementor a composable style with which to specify a storage provider's
working functionality by simply implementing the corresponding interfaces.

It requires the implementation of basic methods such as getters for the plugin
name and backing devices, initialization from a [`VolumeHost`](#volumehost),
tests for support of a `Spec`, and methods that return [`Mounter`](#mounter)
and [`Unmounter`](#unmounter).

##### `PersistentVolumePlugin`

`PersistentVolumePlugin` denotes the ability to persistently store data. It
requires the implementation of [`VolumePlugin`](#volumeplugin) methods along
with another method that describes the possible access modes via a listing of
[PersistentVolumeAccessMode](https://github.com/kubernetes/kubernetes/blob/release-1.6/pkg/api/v1/types.go#L611)
values. The current options are `ReadWriteOnce`, `ReadOnlyMany`, and
`ReadWriteMany`.

The correct list of options is specific to the vendor and product.

##### `RecyclableVolumePlugin`

`RecyclableVolumePlugin` denotes the ability of a volume plugin to recycle a
volume and become available for a new claim.
The implementation of recycling currently involves creating a new pod to scrub
clean the volume of an existing pod in a method called
`RecycleVolumeByWatchingPodUntilCompletion` found in
[pkg/volume/util.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/volume/util.go). By default, the recycling itself is handled by the pod template generated by `NewPersistentVolumeRecyclerPodTemplate` found in [pkg/volume/plugins.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/volume/plugins.go).

With the target volume mounted to `/scrub`, it currently runs in busybox:
```
test -e /scrub && rm -rf /scrub/..?* /scrub/.[!.]* /scrub/*  && test -z \"$(ls -A /scrub)\" || exit 1
```

##### `DeletableVolumePlugin`

`DeletableVolumePlugin` denotes the ability of a volume plugin to delete a
volume. It extends [`VolumePlugin`](#volumeplugin) and must implement a method
that returns an instance of [`Deleter`](#deleter).

##### `ProvisionableVolumePlugin`

`ProvisionableVolumePlugin` denotes the ability of a volume plugin to provision
a new volume. It extends [`VolumePlugin`](#volumeplugin) and must implement a
method that accepts [`VolumeOptions`](#volumeoptions), returning a new
instance of [`Provisioner`](#provisioner).

##### `AttachableVolumePlugin`

`AttachableVolumePlugin` denotes the requirement of a volume plugin to attach
and detach to a node before mounting and after unmounting, respectively. Block
stores in many cloud providers require this. It extends
[`VolumePlugin`](#volumeplugin) with methods that return instances of
[`Attacher`](#attacher) and [`Detacher`](#detacher) and a function that returns
all referencing paths to the mountable device. It's possible via symlinks and
hardlinks to have multiple references to the same device, and the
[`Mounter`](#mounter) needs a list of all of them to account for this.

#### `VolumeHost`

`VolumeHost` provides simplified access to a kubelet that restricts the method
set to methods relevant to volume plugins. If necessary, a kubelet client
interface can be returned for more advanced operations. This is passed into
each [`VolumePlugin`](#volumeplugin)'s initialization method.


#### `MetricsProvider`

The `MetricsProvider` interface requires a method to generate the `Metrics`
data structure mentioned above. The current implementations of
`MetricsProvider` are named `pkg/volume/metrics_*.go` and rely on system calls
such as `statfs` or command-line utilities such as `du` to determine storage
statistics. All implementations of [`Volume`](#volume) must extend this
interface, and they commonly do this by leveraging existing implementations
such as `volume.MetricsStatFS`. `Volume` plug-ins that don't support volume
metrics may use `volume.MetricsNil`.

#### `Volume`

The `Volume` interface currently must only implement `MetricsProvider` and a
method to return the path to which the volume should be mounted to for a pod.
This data type is not used very often in the codebase aside from the kubelet
[volume mount listing
methods](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet_volumes.go).
Instances of these data types are associated at the kubelet level.

Rather, this type is used as a building block for [`Mounter`](#mounter),
[`Unmounter`](#unmounter), and [`Deleter`](#deleter) to factor the common
necessity of a pod mountpoint.

Implementations typically use methods in the [`VolumeHost`](#volumehost)
interface to query the kubelet for this information.

##### `Mounter`

`Mounter` extends `Volume` by implementing several additional methods that
check if a volume can mount, do the actual mounting, and return the `Mounter`
implementation's [`Attributes`](#attributes).

Interfaces found in
[pkg/util/mount](https://github.com/kubernetes/kubernetes/blob/master/pkg/util/mount/)
are typically used to bind mount the global disk mount to a directory that the
individual pod can use.

##### `Unmounter`

`Unmounter` is the counterpart to `Mounter`, and it extends `Volume` with
methods that can tear down mounted volumes so that they can be safely detached
or deleted.

#### `Provisioner`

`Provisioner` extends the `Volume` interface with a method that provisions a
new volume on a storage provider. The configuration for the volume is stored in
an [`VolumeOptions`](#volumeoptions) instance which is eventually passed to
`ProvisionableVolumePlugin.NewProvisioner()`.

##### `Deleter`

`Deleter` extends the `Volume` interface with a method that deletes the
underlying volume from a storage provider. Implementations of `Deleter`
typically do not unmount or detach the volume before attempting deletion, and
instead, they return an error if something goes wrong.


#### `Attacher`

`Attacher` implements many methods used in the attachment of volumes and
verification of correct attachment. Its methods accept a [`Spec`](#spec)
instance and commit actions or return observations against it. `Attacher` also
knows how to mount block devices to global paths for individual pods to bind
mount.

Note that the `Attacher` doesn't extend the [`Volume`](#volume) interface.
Implementations typically contain a reference to [`VolumeHost`](#volumehost)
(simplified kubelet access) and the corresponding instance of
[`cloudprovider.Interface`](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go).

Interfaces found in
[pkg/util/mount](https://github.com/kubernetes/kubernetes/blob/master/pkg/util/mount/)
are typically used to mount the new block device to a global disk mount directory.

#### `BulkVolumeVerifier`

`BulkVolumeVerifier` implements a method that can query about attached volumes
statuses en-masse. It can be thought of as a higher-level
`Attacher.VolumesAreAttached` able to query volume states across different
nodes in a cluster using a single call. Currently, AWS is one of the few block
storage providers that support this.

#### `Detacher`

`Detacher` is the counterpart to [`Attacher`](#attacher), and implements
methods that unmount block devices from global paths and send detachment calls
to providers for specific volumes.

Note that the `Detacher` doesn't extend the [`Volume`](#volume) interface.
Implementations typically contain a reference to [`VolumeHost`](#volumehost)
(simplified kubelet access) and the corresponding instance of
[`cloudprovider.Interface`](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go).


## FAQ

##### How do cloud provider volume plugins figure out what mount device (e.g., `/dev/sdc`) a newly-attached volume will appear as?

This depends on the cloud provider. For example,


- The AWS EBS volume plugin uses
  [`volume/aws.getMountDevice`](https://github.com/kubernetes/kubernetes/blob/a2c7eb275443a5afea4052094a07daf627e4a724/pkg/cloudprovider/providers/aws/aws.go#L1207)
  and relies on having block device mappings readily available per instance via
  the AWS API. See the section titled "Viewing the EBS Volumes in an Instance
  Block Device Mapping" on this [support page](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html).
  The cloud provider interface can also predict what the next device mounting
  point would be for volumes that haven't yet been attached to a block device
  in `/dev`.
- The GCE PD volume plugin leverages that the GCP API allows specification of
  `deviceName` in the `attachDisk` call. See `deviceName` under "Request Body"
  on this [support page](https://cloud.google.com/compute/docs/reference/beta/instances/attachDisk)
  for details. This is done specifically in [`gcePersistentDiskAttacher.Attach`](https://github.com/kubernetes/kubernetes/blob/bf984aa328f3d3e5f8956f7d8e65c13c52102426/pkg/volume/gce_pd/attacher.go#L90).
- The Azure DD volume plugin uses the data disk's Logical Unit Number (LUN)
  returned by the
  [API](https://msdn.microsoft.com/en-us/library/azure/jj157199.aspx) to find a
  match in the devices. Due to the predictable naming scheme, the cloud
  provider interface can also return the next available LUN for a volume that
  hasn't yet been attached to a system block device in `/dev`.

