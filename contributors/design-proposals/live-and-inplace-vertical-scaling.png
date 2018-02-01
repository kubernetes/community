Live and In-place Vertical Scaling Proposal
================

_Authors:_
* @YoungjaeLee - Youngjae Lee &lt;leeyo@us.ibm.com&gt;
* @karthickrajamani - Karthick Rajamani &lt;karthick@us.ibm.com&gt;

# Abstract

This proposal is to enable the resizing of resources allocated to running containers of a pod without having to restart the pod.
The initial focus is on enabling this for Statefulsets, in particular, though this may be use-able for stand-alone/controller-free pods and Deployments as well.

# Motivation

Resources needed by containers can change over time for a variety of reasons - moving from live-test mode to production usage, change in user load or dataset sizes each of which again might come about for a variety of reasons.
`Statefulset` supports the capability to change Request and Limit values specified for a container through supported pod spec update methods.
However, they currently require the pods be restarted to run with the new resource sizes.

There are a few reasons why restarting may not be desirable particularly for stateful services.

1. Moving or copying existing state information or re-sharding employing horizontal scaling may be expensive or even infeasible depending on the service or underlying application implementation.

2. Services with large amount of state associated with specific instances may encounter temporary loss in access to the specific state information while its corresponding container is being restarted. Or they might encounter non-trivial delays in getting back to expected performance.

3. State may be maintained on node-local storage/devices and restart might place the container/pod on a different node (unless suitably constrained). Might create more fragile scaling option for service and/or additional burden (to constrain placement) on user.

4. Where the restart is not needed, unnecessary load may be created on Kubernetes components that are responsible for pod deletion and creation potentially impacting Kubernetes scalability.
Resizing running containers of a pod in place could avoid that additional load.

Conversely, resizing without restart may not be the right option for all applications and the right option may also be resource dependent.
For example, an application might (or have the means to) assess resources available to it only at its initialization time and be unable to adjust its usage if those are dynamically changed.
Or it might be able to adjust to changes in certain resources (most can for CPU) but not for others (some application have particular difficulty freeing memory they have taken up).
So we need the means to express and implement the best approach for resizing for each workload and resource.

# Use Cases

* After deploying a service/application through `StatefulSet`, I want to resize, by giving a updated pod spec with new resource sizes, the resources (for now, CPU and memory) allocated to each pod of the statefulset without restarting the pods.

* A (vertical) pod auto-scaler (c.f. [a related issue](https://github.com/kubernetes/features/issues/21)) can utilize this feature, while currently it does resize a pod with restart only.

# Objectives

1. Enable live and in-place resource resizing on a pod.
2. Add support for live and in-place resource resizing in `StatefulSet` controller.

# API and Usage

To express the policy for resizing in a pod spec, we introduce resource attribute `resizePolicy` with the following choices for value:
* RestartOnly. (the current behavior, default)
* LiveResizeable.

This attribute will be available per resource (such as cpu, memory) and so is adequate to indicate whether the workload can handle and prefer a change in each resource’s allocation for it without restarting.
With potentially multiple containers and multiple resizeable resources for each in a Pod, the response to an update of the pod spec will be determined by the a precedence order among the attribute values with RestartOnly dominating LiveResizeable, i.e., if two resources have been resized in the update to the spec and one of them has a policy of RestartOnly then the pod would be restarted to realize both updates.

We can optionally introduce an action annotation `resizeAction` with following choices for value:
* Restart. (default)
* LiveResize.
* LiveResizePreferred.

If used, this would be included as part of the patch or appropriate update command providing the spec update for the resize.
It would indicate the preference of user at the time of resize.
Specifically, Restart for `resizeAction` would indicate the pod be restarted for the corresponding resizing of resource(s), LiveResize would indicate the pod not be restarted the resize be realized live, and LiveResizePreferred would indicate that the resize be realized preferrably live but if that fails for any reason to accomplish it with a restart.

An example of the usage of resizePolicy attribute in a pod spec:

```yaml
resources:
    requests:
        cpu: 100m
        memory: 1Gi
    limits:
        cpu: 1000m
        memory: 1Gi
    resizePolicy:
        cpu: LiveResizeable
        memory: RestartOnly
```

For the above example, if there is a change to cpu request or limit it can be vertically scaled only if the memory request and limit remained the same, otherwise the RestartOnly policy for memory would override the policy for CPU, and the Pod (container, if container-alone restart is allowed) would need to be restarted.

An example of change in CPU resource `request` size and usage of `resizeAction` annotation with the updated spec.

// before the change
```yaml
resources:
    requests:
        cpu: **100m**
        memory: 1Gi
    limits:
        cpu: 1000m
        memory: 1Gi
    resizePolicy:
        cpu: LiveResizeable
        memory: RestartOnly
```

// after the change
```yaml
annotation:
resizeAction: LiveResizePreferred
...
resources:
    requests:
        cpu: **400m**
        memory: 1Gi
    limits:
        cpu: 1000m
        memory: 1Gi
    resizePolicy:
        cpu: LiveResizeable
        memory: RestartOnly
```

## Combining Stateful Set Update options with Vertical Scaling

`StatefulSet` supports two update options with `spec.updateStrategy.type`, where OnDelete applies the changed spec when restarting a pod after an explicit delete command and the RollingUpdate is applied by the controller by restarting each of the pods that are members of the Stateful set with the changed values, in reverse ordinal sequence (see [https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/#updating-statefulsets](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/#updating-statefulsets) for more details).

In the table below we propose how the resource resizing directives for vertical scaling can be applied in conjunction with the 'StatefulSet' update strategy.

|resizePolicy (resizeAction)|OnDelete|RollingUpdate|
|---|---|---|
|RestartOnly (or Action=Restart)|Resize resource with restart (with delete command)|Resize resource with restart (with allowed spec update commands)|
|LiveResizeable (and Action=LiveResize)|Resize resource live only (with allowed spec update commands)|Resize resource live only (with allowed spec update commands)|
LiveResizeable (and Action=LiveResizePreferred or no Action specified)|Resize resource with live resize preferred (with allowed spec update commands), i.e., resize resource with restart (with delete command) if not able to resize live.|Resize resource with live resize preferred (with allowed spec update commands), i.e., resize resource with restart (managed by controller) if not able to resize live.|

## Desired Approach

Any valid method to update the pod spec should be applicable for vertical scaling, e.g., using kubectl commands set, patch, apply, edit.
Logs associated with the pod will capture the failure/success of the resize command.
The controller will continue to attempt the update to the spec while there is a difference between the current size and size in updated spec.
If an update is partially successful, user can know this from the logs and attempt to rectify the situation by submitting new updates (that can restore original size(s) or go for a size feasible on all the nodes). 

The policy for vertical scaling can itself be mutable.
It is desirable to submit a change in policy, confirm its acceptance (by re-reading the changed pod spec) and then submit a change to a resource size impacted by that policy. 

# Design and Implementation

This section highlights some of implementation details by describing API changes and the workflow of vertical scaling on a pod.

## Changes on API and key components

This section describes briefly API changes for pod-level vertical scaling and the related changes on the 'API server', the 'Scheduler', and the 'Kubelet'.

* **ResizeRequest**

A new data structure, ResizeRequest, is added to v1.PodSpec:

```go
const (
    ResizeRequested ResizeStatus = "Requested"
    ResizeAccepted ResizeStatus = “Accepted”
    ResizeRejected ResizeStatus = “Rejected”
    ResizeNone      ResizeStatus = "None"
)

type ResizeRequest struct {
    RequestStatus ResizeStatus
    NewResources  []ResourceRequirements // indexed by containers’ index
}

Type PodSpec {
    …
    ResizeRequest ResizeRequest
    ...
}
```

ResizeRequest has two variables, RequestStatus and NewResources.
RequestStatus represents the status of a resource resizing request.
ResizeRequested indicates resource resizing for a pod is requested.
ResizedAccepted and ResizedRejected means that the requested resource resizing is accepted and rejected, respectively, by the Scheduler.
The NewResources is an array indexed by a container’s index and its each entry holds new resource requirements of a container that needs to resize.

Given a new PodSpec with new resource requirements from a client, first the 'API server' validates it.
If it is valid, the 'API server' sets the RequestStatus to ResizedRequested and copies the new resource requirements of each container into the NewResources.
Also, the 'API server' restores the resource requirements of each container of the PodSpec to the original and writes the revised PodSpec to ETCD to communicate with the Scheduler.
This is because at this moment the PodSpec on ETCD shouldn’t be updated with new resource requirements.

For a pod with ResizeRequested, the 'Scheduler' checks if the node on which the pod currently runs has enough resources to resize the pod.
The Scheduler notify the 'API server' of the result via 'Resizing' API operation, which will be describe below.

* **Resizing**

A new API, Resizing, for the 'scheduler' is introduced:

```go
// Resizing resizes the resources allocated to a pod
type Resizing struct {
    metav1.TypeMeta
    metav1.ObjectMeta
    Request ResizeRequest
}
```

Resizing has the metadata of a pod to resize and a value of ResizeRequest that holds the status of a resizing request, which indicates whether the resizing is feasible or not, and new resource requirements of the pod.

Once the 'Scheduler' determine whether a resource resizing on a pod is feasible, or not, it notifies to the API server via this Resizing API.
ObjectMeta.Namespace/Name are set to the Namespace and Name of the pod to resize and ResizeRequest is set with new resource requirements.

Given a Resizing API operation, the ResizeStatus of the PodSpec of a Pod is updated according to that of the Resizing operation.
If the ResizeStatus is ResizeAccepted, the API server updates the ResourceRequirement of each container of a pod with new resource requirements on ETCD.

* **PodResized**

A new pod condition, PodResized, and condition status for that is added to v1.PodCondition:

```go
// These are valid conditions of pod.
const (
       // PodResized represents the status of the resizing process for this pod
    PodResized PodConditionType = "PodResized"
    PodReasonUnresizable   = "Unresizable"
    PodReasonResizerFailed = "ResizerFailed"
)
const (
    ConditionRequested ConditionStatus = "Requested"
    ConditionAccepted  ConditionStatus = "Accepted"
    ConditionRejected  ConditionStatus = "Rejected"
    ConditionDone      ConditionStatus = "Done"
)
```

The pod condition of PodResized represents the status of the resizing process.
The PodResized Condition is updated by the `Kubelet` according to the ResizeStatus, which is updated by the 'API server'.

Basically, when the ResizeStatus is changed, the `Kubelet` updates the PodResized condition accordingly.
In case of ConditionDone, the `Kubelet` sets the PodResized of a Pod to it when all the containers that need to be resized complete to be resized.


* **A new additional hash, called `expectedHashNoResources`, added for `Kubelet` to detect a change on resource requirements**

In order to watch resource requirement changes efficiently, a new additional hash is added to kubecontainer.ContainerStatus (and that is also stored as a one of the container’s labels).
This hash is calculated with a container’s spec munged with an empty v1.ResourceRequirements.
In addition to the existing container spec’s hash, Kubelet uses this new hash to detect a change on resource requirements of a container.
Without this new hash, still Kubelet could detect a change on resource requirement with the existing hash, but needs to compare every entry in the container’s spec to identify which entry changes.
The following is the details. 

```go
expectedHash := kubecontainer.HashContainer(&container)
containerChanged := containerStatus.Hash != expectedHash
if containerChanged {
    // Something in the container’s spec changed. 
    // So, see if it is just a change only on resource requirements.
    mungedContainer := container
    mungedContainer.Resources = v1.ResourceRequirements{}
    expectedHashNoResources := kubecontainer.HashContainer(&mungedContainer)
    containerChangedToStart = containerStatus.HashNoResources != expectedHashNoResources

    if containerChangedToStart {
        // This is a change on something other than resource requirements, so it needs to restart a container
    } else {
        // This is a change only on resource requirements.
    }
}
```

## Workflow

This describes the sequence of a pod-level vertical scaling example to resize the CPU resource of a pod from 1 to 2.

![Process](live-and-inplace-vertical-scaling.png)

0. A pod has 1 CPU.

1. A client requests resource resizing on the pod with a new PodSpec with 2 CPU.

2. The API server updates the PodSpec with ResizeRequest on etcd

3. The PodSpec is updated on etcd.

4. The Scheduler checks if the resizing is feasible and, if so, issues a Resizing API operation with the ResizeRequest of the “Accepted” status.

5. The API server updates the PodSpec with the new resource requirement on etcd and also modifies the ResizeRequest of the PodSpec to “Accepted”.

6., 7. The Kubelet updates the PodResized condition to “Accepted”.

8. The Kubelet detects the change of resource requirements on the container and updates the cgroup configuration of the container via UpdateContainerResources CRI interface.

9. The Kubelet modifies the status of the PodResized condition to Done after every update on the cgroup configuration of all containers to resize is done. 

10. It completes to resize the pod to have 2 CPUs.


## Implementation Phases

* Phase 1 - Introduce live and in-place vertical scaling on a pod

Status: In progress, a working prototype implemented originally in the master branch (v1.10-alpha)

The working code is available at [https://github.com/YoungjaeLee/kubernetes](https://github.com/YoungjaeLee/kubernetes) (the qos-master branch)

* Phase 2 - Adding support for live and in-place vertical scaling in StatefulSet

Status: In progress in the master branch (v1.10-alpha)


# Related issues

1. QoS class change by resize is not supported.

For each of the Burstable and the Best-effort class, the Kubelet maintains a class-level cgroup under the ‘kubepods’ cgroup, which is the parent cgroup of pods of each QoS class. (e.g. kubepods/burstable is the parent cgroup of Burstable pods).
So, in order to change the QoS class of a pod, it needs not only to resize resources, but also to change its parent cgroup properly.
But, once a pod is created, its parent cgroup cannot be changed with the current Docker API. So, the QoS-class of a pod cannot be changed by resource resizing.

2. Memory-resizing to change a request value might not take effect for Burstable pods.

For Burstable pods, a request value for memory resource determines the value of a score for the OOM killer, but Docker doesn’t support to change dynamically the score of an existing container.
So, the change to a memory request value doesn’t take effect for Burstable pods on the OOM killer’s behavior.
But, for Guaranteed and Best-effort pods, this is not an issue because the score is fixed regardless of its memory request value. (in this case, the memory request value is used only for admission control by Scheduler) 

3. Memory-resizing to decrease its limit may fail on the Kubelet in some circumstances.

By default, the Kubelet requires disabling swap. With swap disabled, memory-resizing to decrease fails when there is not enough free memory that can be reclaimed.
Specifically, the cgroup change to write a new limit value to memory.limit_in_bytes that is smaller than the current value fails as the Linux kernel fails to reclaim memory in use exceeding the new value.

4. Memory-resizing on memory-backed storage (emptydir backed by memory)

In the perspective of resizing, there is no difference between normal memory and the emptydir memory.
Basically, if the memory allocated to a container is resized, the amount of memory available to an emptydir changes accordingly.

With respect to usage accounting, the emptydir memory has a little difference (not specific to resizing.).
The emptydir memory is accounted to a container that allocates the memory to store a file in the emptydir.
For example, in a case where two containers share a memory-backed emptydir, the memory used to store a file is accounted to one of the containers that created the file and wrote its data, even though the other container is able to read/modify the file as if the file resides on its own allocated memory.
If the second container appends some data at the end of the file, the corresponding memory is accounted to the second container since it is allocated by the second container.

