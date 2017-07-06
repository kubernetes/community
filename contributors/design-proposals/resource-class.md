# Resource Classes Proposal

  1. [Abstract](#abstract)
  2. [Motivation](#motivation)
  3. [Use Cases](#use-cases)
  4. [Objectives](#objectives)
  5. [Non Objectives](#non-objectives)
  6. [Resource Class](#resource-class)
  7. [API Changes](#api-changes)
  8. [Scheduler Changes](#sch-changes)
  9. [Kubelet Changes](#kubelet-changes)
 10. [Opaque Integer Resources](#oir)
 11. [Future Scope](#future-scope)

_Authors:_

* @vikaschoudhary16 - Vikas Choudhary &lt;vichoudh@redhat.com&gt;
* @aveshagarwal - Avesh Agarwal &lt;avagarwa@redhat.com&gt;

## Abstract
In this document we will describe *resource classes* which is a new model to
represent compute resources in Kubernetes. This document should be seen as a
successor to [device plugin proposal](https://github.com/kubernetes/community/pull/695/files)
and has a dependency on the same.

## Motivation
Kubernetes system knows only two resource types, 'CPU' and 'Memory'. Any other
resource can be requested by pod using opaque-integer-resource(OIR) mechanism.
OIR is a key-value pair with the key being a string and the value being a 'Quantity'
which can (optionally) be fractional. The current model is great for supporting
simple compute resources like CPU or Memory, which are available across all
kubernetes deployments.
But there is a problem in representing resources like GPUs, ASICs, local storage
etc in the form of OIRs. Each of such resource type generally has a rich
metadata like version, capabilities etc. to describe the resource. A particular
application pod may perform well only with resources which have a certain
capability, for example GPUs greater than version 'V'. OIR does not allow such,
metadata based, selection/filtering of resources.
The current model requires identity mapping between available resources and requested
resources. Identity mapping means there must be a one-to-one mapping between the
resource reference used in the spec to request the resource and resource reference
used to advertise the resource availablity.

Since 'CPU' and 'Memory' are resources that are available across all kubernetes
deployments and need no metadata to describe the resource type any further,
the current user facing API (Pod Specification) remains portable as long as pod
requests only CPU and Memory. However the current model cannot support
complex resources like GPUs, ASICs, NICs, local storage, etc.
To support heterogeneity, portability and management at scale, such resources
must be represented(advertised) in a form which allows metadata inclusion and a
metadata based resource selection mechanism.

_GPU Integration Example:_
  * [Enable "kick the tires" support for Nvidia GPUs in COS](https://github.com/kubernetes/kubernetes/pull/45136)
  * [Extend experimental support to multiple Nvidia GPUs](https://github.com/kubernetes/kubernetes/pull/42116)

_Kubernetes Meeting Notes On This:_
  * [Meeting notes](https://docs.google.com/document/d/1Qg42Nmv-QwL4RxicsU2qtZgFKOzANf8fGayw8p3lX6U/edit#)
  * [Better Abstraction for Compute Resources in Kubernetes](https://docs.google.com/document/d/1666PPUs4Lz56TqKygcy6mXkNazde-vwA7q4e5H92sUc)
  * [Extensible support for hardware devices in Kubernetes (join kubernetes-dev@googlegroups.com for access)](https://docs.google.com/document/d/1LHeTPx_fWA1PdZkHuALPzYxR0AYXUiiXdo3S0g2VSlo/edit)

## Use Cases

  * I want to have a compute resource type, 'Resource Class',
    which can be created with meaningful and portable names. This compute
    resource can hold additional metadata as well, for example:
    * `nvidia.gpu.high.mem` is the name and metadata is memory greater than 'X' GB.
    * `fast.nic` is the name and associated metadata is bandwidth greater than
      'B' gbps.
  * If I request a resource `nvidia.gpu.high.mem` for my pod, any 'nvidia-gpu'
    type device which has memory greater than or equal to 'X' GB, should be able
    to satisfy this request, independent of other device capabilities such as
    'version' or 'nvlink locality' etc.
  * Similarly, if I request a resource `fast.nic`, any nic device with speed
    greater than 'B' gbps should be able to meet the request.
  * I want a rich metadata selection interface where operators like 'Eq' for
    'equals to', 'Lt' for 'less than', 'LtEq' for 'less than equal to', 'Gt' for
    'greater than', 'GtEq' for 'greater than and equal to' and 'In' for
    'a set of accepted values' are supported on the compute resource metadata.

## Objectives

1. Define and add support in the API for a new type, *Resource Class*.
2. Add support for *Resource Class* in the scheduler.

## Non Objectives
1. Discovery, advertisement, allocation/deallocation of devices is expected to
   be addressed by [device plugin proposal](https://github.com/kubernetes/community/pull/695/files)

## Resource Class
*Resource Class* is a new type, objects of which provide abstraction over
[Devices](https://github.com/RenaudWasTaken/community/blob/a7762d8fa80b9a805dbaa7deb510e95128905148/contributors/design-proposals/device-plugin.md#resourcetype).
A *Resource Class* object selects devices using `matchExpressions`, a list of
(operator, key, value). A *Resource Class* object selects a device if at least
one of the `matchExpressions` matches with device details. Within a matchExpression,
all the (operator,key,value) are ANDed together to evaluate the result.

*Resource Class* object is non-Namespaced kind of object and post created object
is immutable.

YAML example 1:
```yaml
kind: ResourceClass
metadata:
  name: nvidia.high.mem
spec:
  resourceSelector:
    - matchExpressions:
        - key: "Kind"
          operator: "In"
          values:
            - "nvidia-gpu"
        - key: "memory"
          operator: "GtEq"
          values:
            - "30G"
```
Above resource class will select all the nvidia-gpus which have memory greater
than and equal to 30 GB.

YAML example 2:
```yaml
kind: ResourceClass
metadata:
  name: hugepages-1gig
spec:
  resourceSelector:
    - matchExpressions:
        - key: "Kind"
          operator: "In"
          values:
            - "huge-pages"
        - key: "size"
          operator: "GtEq"
          values:
            - "1G"
```
Above resource class will select all the hugepages with size greater than and
equal to 1 GB.

YAML example 3:
```yaml
kind: ResourceClass
metadata:
  name: fast.nic
spec:
  resourceSelector:
    - matchExpressions:
        - key: "Kind"
          operator: "In"
          values:
            - "nic"
        - key: "speed"
          operator: "GtEq"
          values:
            - "40GBPS"
```
Above resource class will select all the NICs with speed greater than equal to
40 GBPS.


## API Changes
### ResourceClass

Internal representation of *Resource Class*:

```golang
// +nonNamespaced=true
// +genclient=true

type ResourceClass struct {
        metav1.TypeMeta
        metav1.ObjectMeta
        // Spec defines resources required
        Spec ResourceClassSpec
        // +optional
        Status ResourceClassStatus
}
// Spec defines resources required
type ResourceClassSpec struct {
        // Resource Selector selects resources
        ResourceSelector []ResourcePropertySelector
}

// A null or empty selector matches no resources
type ResourcePropertySelector struct {
        // A list of resource/device selector requirements. ANDed from each ResourceSelectorRequirement
        MatchExpressions []ResourceSelectorRequirement
}

// A resource selector requirement is a selector that contains values, a key, and an operator
// that relates the key and values
type ResourceSelectorRequirement struct {
        // The label key that the selector applies to
        // +patchMergeKey=key
        // +patchStrategy=merge
        Key string
        // +optional
        Values []string
        // operator
        Operator ResourceSelectorOperator
}
type ResourceSelectorOperator string

const (
        ResourceSelectorOpIn           ResourceSelectorOperator = "In"
        ResourceSelectorOpEq           ResourceSelectorOperator = "Eq"
        ResourceSelectorOpNotIn        ResourceSelectorOperator = "NotIn"
        ResourceSelectorOpExists       ResourceSelectorOperator = "Exists"
        ResourceSelectorOpDoesNotExist ResourceSelectorOperator = "DoesNotExist"
        ResourceSelectorOpGt           ResourceSelectorOperator = "Gt"
        ResourceSelectorOpGtEq         ResourceSelectorOperator = "GtEq"
        ResourceSelectorOpLt           ResourceSelectorOperator = "Lt"
        ResourceSelectorOpLtEq         ResourceSelectorOperator = "LtEq"
)
```
### ResourceClassStatus
```golang
type ResourceClassStatus struct {
        Allocatable resources.Quantity
        Requested   resources.Quantity
}
```
ResourceClass status is updated by the scheduler at:
1. New *Resource Class* object creation.
2. Node addition to the cluster.
3. Node removal from the cluster.
4. Pod creation if pod requests a resource class.
5. Pod deletion if pod was consuming resource class.

`ResourceClassStatus` serves the following two purposes:
* Scheduler predicates evaluation while pod creation. For details, please refer
  further sections
* User can view the current usage/availability details about the resource class
  using kubectl.

### User story
The administrator has deployed device plugins to support hardware present in the
cluster. Device plugins, running on nodes, will update node status indicating
the presence of this hardware. To offer this hardware to applications deployed
on kubernetes in a portable way, the administrator creates a number of resource
classes to represent that hardware. These resource classes will include metadata
about the devices as selection criteria.

1. A user submits a pod spec requesting 'X' resource classes.
2. The scheduler filters the nodes which do not match the resource requests.
3. scheduler selects a device for each resource class requested and annotates
   the pod object with device selection info. eg:
   `scheduler.alpha.kubernetes.io/resClass_test-res-class_nvidia-tesla-gpu=4`
   where `scheduler.alpha.kubernetes.io/resClass` is the common prefix for all the
   device annotations, `tes-res-class` is resource class name,
   `nvidia-tesla-gpu` is the selected device name and `4` is the quantity requested.

4. Kubelet reads the device request from pod annotation and calls `Allocate` on
   the matching Device Plugins.
5. The user deletes the pod or the pod terminates
6. Kubelet reads pod object annotation for devices consumed and calls `Deallocate`
   on the matching Device Plugins

In addition to node selection, the scheduler is also responsible for selecting a
device that matches the resource class requested by the user.

### Reason for preferring device selection at the scheduler and not at the kubelet
Kubelet does not maintain any cache. Therefore to know the availability of a
device, which is requested by the new incoming pod, kubelet calculates how many
devices are consumed by all already admitted pods, by iterating over all the admitted
pods running on the node. This is done while running predicates for each new
incoming pod at kubelet. Even if we assume that scheduler cache and consumption
state that is created at runtime for each pod, are exactly same, current api
interfaces does not allow to pass selected device to container manager (where
actually device plugin will be invoked from). This problem occurs because
requested resource classes are translated into devices internally through code
and user does not mention device in pod object. While other resource requests
can be determined from the pod object directly.
To summarize, device selection at the kubelet can be done in one of the following
two ways:
* Select device at pod admission while applying predicates and change all api
  interfaces that are required to pass selected device to container manager.
* Create resource consumption state again at container manager and select device.

None of the above approach seems cleaner than doing device selection at scheduler,
which helps to retain cleaner api interfaces between packages.

## Scheduler Changes
Scheduler already listens and maintains state in the cache for any changes in
node or pod objects. We will enhance the logic:
1. To listen and maintain the state in cache for user created *Resource Class* objects.
2. To look for device related details in node objects and maintain accounting for
   devices as well.

From the events perspective, handling for the following events will be added/updated:

### Resource Class Creation
1. Initialize and add resource class info into local cache
2. Iterate over all existing nodes in cache to figure out if there are devices
   on these nodes which are selectable by resource class. If found, update the
   resource class availability status in local cache.
3. Patch the status of resource class api object with availability state in local
   cache

### Resource Class Deletion
Delete the resource class info from the cache.

### Node Addition
Scheduler already caches `NodeInfo`. Now additionally update device state:
1. Check in the node status if any devices are present.
2. For each device found, iterate over all existing resource classes in the cache
   to find resource classes which can select this particular device. For all
   such resource classes, update the availability state in the local cache.
3. ResourceClass api object's status, `ResourceClassStatus` will be patched
   as per the updated availability state in the cache.

### Node Deletion
If node has devices which are selectable by existing resource classes:
1. Adjust resource class state in local cache.
2. Update resource class status by patching api object.

### Pod Creation
1. Get the requested resource class name and quantity from pod spec.
2. Select nodes by applying predicates according to requested quantity and Resource
   class's state present in the cache.
3. On the selected node, select a Device from the stored devices info in cache
   after matching key,value from requested resource class.
4. After device selection, update(decrease) 'Requested' for all the resource
   classes which could select this device in the cache.
5. Patch the resource class objects with new 'Requested' in the `ResourceClassStatus`.
6. Add the pod reference in local DeviceToPod mapping structure in the cache.
7. Patch the pod object with selected device annotation with prefix 'scheduler.alpha.kubernetes.io/resClass'

NOTE: This proposal propose only 'first fit' as device selection strategy.
      In the future, this can be extended to multiple algorithms available for
      the user to choose from, in a configurable manner.

### Pod Delete
1. Iterate over the all the devices on the at which pod was scheduled to and
   find out the devices being used by pod.
2. For each device consumed by pod, update availability state of Resource classes
   which can select this device in the cache.
3. Patch `ResourceClassStatus` with new availability state.

## Kubelet Changes
Update logic at container runtime manager to look for device annotations,
prefixed by 'scheduler.alpha.kubernetes.io/resClass' and call matching device
plugins.

## Opaque Integer Resources
This API will supersede the [Opaque Integer Resources](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#opaque-integer-resources-alpha-feature)
(OIR). External agents can continue to attach additional 'opaque' resources to
nodes, but the special naming scheme that is part of the current OIR approach
will no longer be necessary. Any existing resource discovery tool which updates
node objects with OIR, will adapt to update node status with devices instead.


## Future Scope
* RBAC: It can further be explored that how to tie resource classes with RBAC
  like any other existing API resource objects.
* Nested Resource Classes: In future device plugins and resource classes can be 
  extended to support the nested resource class functionality where one resource
  class could be comprised of a group of sub-resource classes. For example 'numa-node'
  resource class comprised of sub-resource classes, 'single-core'.
* Multiple device selection algorithms, each with a different selection strategy,
  will be added to the scheduler and cluster admin will be able to configure one
  as per his/her choice.
