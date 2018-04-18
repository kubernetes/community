# Dynamic attached volume limits

## Goals

Currently the number of volumes attachable to a node is either hard-coded or only configurable via an environment variable. Also
existing limits only apply to well known volume types like EBS, GCE and is not available to all volume plugins.

This proposal enables any volume plugin to specify those limits and also allows same volume type to have different volume
limits depending on type of node.

## Implementation Design

### Prerequisite

* This feature will be protected by an alpha feature gate, so as API and CLI changes needed for it. We are planning to call
  the feature `AttachVolumeLimit`.

### API Changes

There is no API change needed for this feature. However existing `node.Status.Capacity` and `node.Status.Allocatable` will
be extended to cover volume limits available on the node too.

The key name that will store volume will be start with prefix `storage-attach-limits-`. The volume limit key will respect
format restrictions applied to Kubernetes Resource names. Volume limit key for existing plugins might look like:


* `storage-attach-limits-aws-ebs`
* `storage-attach-limits-gce-pd`

`IsScalarResourceName` check will be extended to cover storage limits:

```go
func IsStorageAttachLimit(name v1.ResourceName) bool {
    return strings.HasPrefix(string(name), v1.ResourceStoragePrefix)
}

// Extended and Hugepages resources
func IsScalarResourceName(name v1.ResourceName) bool {
    return IsExtendedResourceName(name) || IsHugePageResourceName(name) ||
        IsPrefixedNativeResource(name) || IsStorageAttachLimit(name)
}
```

The prefix `storage-attach-limits-*` can not be used as a resource in pods, because it does not adhere to specs defined in following function:


```go
func IsStandardContainerResourceName(str string) bool {
    return standardContainerResources.Has(str) || IsHugePageResourceName(core.ResourceName(str))
}
```

Additional validation tests will be added to make sure we don't accidentally break this.

#### Alternative to using "storage-" prefix
We also considered using currently defined `GetPluginName` interface(of Volume Plugins) for using as key in the `node.Status.Capacity`. Ultimately
we decided against using it, because most in-tree plugins start with `kubernetes.io/` and we needed a uniform way to identify storage
related capacity limits in `node.Status`.

### Changes to scheduler

Scheduler will retrieve available attachable limit on a node from `node.Status.Allocatable` and store it in `nodeInfo` cache. Volume
limits will be treated like any other scalar resource.

For `AWS-EBS`, `AzureDisk` and `GCE-PD` volume types, existing `MaxPD*` predicates will be updated to use volume attach limits available
from node's allocatable property. To be backward compatible - the scheduler will fallback to older logic, if no limit is set in `node.Status.Allocatable` for AWS, GCE and Azure volume types.

### Setting of limit for existing in-tree volume plugins

The volume limit for existing volume plugins will be set by querying the volume plugin. Following function
will be added to volume plugin interface:

```go
type VolumePluginWithAttachLimits interface {
    // Return key name that is used for storing volume limits inside node Capacity
    // must start with storage- prefix
    VolumeLimitKey(spec *Spec) string
    // Return volume limits for plugin
    GetVolumeLimits() (map[string]int64, error)
}
```

When querying the plugin - plugin will use `ProviderName` function of CloudProvider to check
if plugin is usable on the node. For example - querying for `GetVolumeLimits` from `aws-ebs` plugin with `gce` cloudprovider
will result in error.

Kubelet will query the volume plugins inside `kubelet.initialNode` function and populate `node.Status` with returned values.

For GCE and AWS - `GetVolumeLimits` will return limits depending on node type. Plugin already has node name accessible
via `VolumeHost` interface and hence it will check the node type and return the volume limits.

We do not aim to cover all in-tree volume types. We will support dynamic volume limits proposed here for following volume types:

* GCE-PD
* AWS-EBS
* AzureDisk

We expect to add incremental support for other volume types.

### Phase 2

CSI implementation will be done in phase 2. The CSI spec change that reports
volume limits from CSI plugin has been already merged, however we still need to decide
what is the best way Kubernetes scheduler can identify individual CSI plugin and compare it against
limit available from node's allocatable.
