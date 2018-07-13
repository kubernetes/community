---
kep-number: 14
title: New Resource API Proposal
authors:
  - "@vikaschoudhary16"
  - "@jiayingz"
owning-sig: sig-node
participating-sigs:
  - sig-scheduling
reviewers:
  - "@thockin"
  - "@derekwaynecarr"
  - "@dchen1107"
  - "@bsalamat"
  - "@vishh"
approvers:
  - "@sig-node-leads"
  - "@sig-scheduling-leads"
editor: "@vikaschoudhary16"
creation-date: "2018-06-14"
last-updated: "2018-06-14"
status: provisional
---
# New Resource API Proposal

Table of Contents
=================
* [Abstract](#abstract)
* [Background](#background)
* [Use Stories](#user-stories)
* [Objectives](#objectives)
* [Non Objectives](#non-objectives)
* [Components](#components)
  * [ResourceClass API](#resourceclass-api)
  * [Kubelet Extension](#kubelet-extension)
  * [Scheduler Extension](#scheduler-extension)
  * [Quota Extension](#quota-extension)
* [Roadmap](#roadmap)

## Abstract
In this document we will describe a new resource API model to better support non-native compute resources on Kubernetes.

## Background
We are seeing increasing needs to better support non-native compute resources on Kubernetes that cover a wide range of resources such as GPUs, High-performance NICs, Infiniband and FPGAs. Such resources often require vendor specific setup, and have rich sets of different properties even across devices of the same type. This brings new requirements for Kubernetes to better support non-native compute resources to allow vendor specific setup, dynamic resource exporting, flexible resource configuration, and portable resource specification.

The device plugin support added in Kubernetes 1.8 makes it easy for vendors to dynamically export their resources through a plugin API without changing Kubernetes core code. Taking a further step, this document proposes a new resource abstraction API, ResourceClass, that can be used to describe, manage and consume such vendor specific and metadata rich resources in simple and portable ways.

## Use Stories
- As a cluster operator, I manage different types of GPUs in my cluster. I want the workloads running in the cluster to be able to request and consume different types of GPUs easily, as long as they have enough quota. I want to assign different resource quota on different type of GPUs, eg: Kepler K20, K80, Pascal P40, P100, Volta V100, as they can have very different performance, price, and available units.<br/>
**Motivation:** Empower enterprise customers to consume and manage non-primary resources easily, similar to how they consume and manage primary resources today.<br/>
**Can this be solved without resource classes:** Without ResourceClass, people would rely on `NodeLabels`, `NodeAffinity`, `Taints`, and `Tolerations` to steer workloads to the appropriate nodes, or build their own [non-upstream solutions](https://github.com/NVIDIA/kubernetes/blob/875873bec8f104dd87eea1ce123e4b81ff9691d7/pkg/apis/core/types.go#L2576) to allow users to specify their resource specific metadata requirements. Workloads would have different experience on consuming non-primary compute resources on k8s. As time goes and more non-upstream solutions were deployed, user experience becomes fragmented across different environments. Furthermore, `NodeLabels` and `Taints` were designed as node level properties. They can't support multiple types of compute resources on a single node, and don't integrate well with resource quota. Even with the recent [Pod Scheduling Policy proposal](https://github.com/kubernetes/community/pull/1937), cluster admins can either allow or deny pods in a namespace to specify a `NodeAffinity` or `Toleration`, but cannot assign different quota to different namespaces.<br/>
**How Resource classes can solve this:** I, operator/admin, create different ResourceClasses for different types of GPUs. User workloads can request different types of GPUs in their `ContainerSpec` resource requests/limits through the corresponding ResourceClass name, in the same way as they request primary resources. Now since resource classes are quota controlled, end-user will be able to consume the requested GPUs only if they have enough quota.<br/>
**Similar use case for network devices:** A cluster can have different types of high-performance NICs and/or infiniband cards, with different performance and cost. E.g., some nodes may have 40 Gig high-performance NICs and some may have 10 Gig high-performance NICs. Some devices may support RDMA and some may not. Different workloads may desire to use different type of high-network access devices depending on their performance and cost tradeoff.</br>
**Similar use case for FPGA:** A cluster can have different FPGA cards programmed with different functions or supports different functionalities. For example, some might have embedded SDN control plane functionality and some might have embedded crypto logic. One workload may want a subset of these FPGA functionalities advertised as resource attributes.</br>

- As a cluster operator, nodes in my cluster have GPU HW from different generations. I want to classify GPU nodes into one of the three categories, silver, gold and platinum depending upon the launch timeline of the GPU family eg: Kepler K20, K80, Pascal P40, P100, Volta V100. I want to charge each of the three categories differently. I want to offer my clients 3 GPU rates/classes to choose from.<br/>
**Motivation:** As time progresses in a cluster lifecycle, new advanced, high performance, expensive variants of GPUs get added to the cluster nodes. At the same time older variants also co-exist. There are workloads which strictly wants latest GPUs and also there are workloads which are fine with older GPUs. But since there is a wide range of types, it will be hard to manage and confusing at the same time to have granularity at each GPU type. Grouping into few broad categories will be convenient to manage.<br/>
**Can this be solved without resource classes:** A unique taint can be used to represent a category like silver. Nodes can be tainted accordingly depending upon the type of GPUs availability. User pods can use tolerations to steer workloads to the appropriate nodes. Now there are two problems. First, access control on tolerations and second, mechanism to quota resources. Though a new feature, [Pod Scheduling Policy](https://github.com/kubernetes/community/pull/1937), is under design that can address first problem of access control on tolerations, there is no solution for second problem i.e quota control.<br/>
**How Resource classes can solve this:** I, operator/admin, creates three resource classes: GPU-Platinum, GPU-Gold, GPU-Silver. Now since resource classes are quota controlled, end-user will be able to request resource classes only if quota is allocated.<br/>
**Similar use case for network devices:** A cluster may have different classes (different capabilities shown as different resource attributes) of a network device. End user’s expectations are met as long as device has a very small subset of these capabilities. Cluster operators want a mechanism where end user can request devices which satisfies their minimum expectation.<br/>
**Similar use case for FPGA:** Some FPGA devices might have embedded SDN control plane functionality, some might have embedded crypto logic. One workload may want a subset of these FPGA functionalities advertised as resource attributes.<br/>

- As an user, I want to be able to utilize different 'types' of a HW resource (may be from the same vendor) while not losing workload portability when moving from one cluster/cloud to another. There can be one type of High-performance NIC on one cluster and another type of high-performance NIC on another cluster. I want to offer high-performance NICs to be consumed under a same portable name, as long as their capabilities are almost same. If pods are consuming these high-performance NICs with a generic resource class name, workload can be migrated from one cluster to another transparently.<br/>
**Motivation:**  Promotes workload portability and less down time.<br/>
**How Resource classes can solve this:** The user can create different resource classes in different environments to match the underlying hardware configurations, but with the same ResourceClass name. This allows workloads to migrate around different environments without changing their workload specs.<br/>
**Can this be solved without resource classes:** No.<br/>

- As a vendor, I want an easy and extensible mechanism to export my resource to Kubernetes. I want to be able to roll out new hardware features to the users who require those features without breaking users who are using old versions of hardware. I want to provide my users some example best-practice configuration specs, with which their applications can use my resource more efficiently.<br/>
**Motivation:** Enables more compute resources and their advanced features on Kubernetes<br/>
**Can this be solved without resource classes:**<br/>
Yes, Using node labels and NodeLabelSelectors.<br/>
Problem: Lack of access control and lack of the ability to differentiate between hardware properties on the same node. E.g., if on the same node, some GPU devices are connected through nvlink while others are connected through PCI-e, vendors don’t have ways to export such resource properties that can have very different performance impacts.<br/>
**How Resource classes can solve this:**<br/>
Vendors can use DevicePlugin API to propagate new hardware features, and provide best-practice ResourceClass spec to consume their new hardware or new hardware features on Kubernetes. Vendors don’t need to worry supporting this new hardware would break existing use cases on old hardware because the Kubernetes scheduler takes the resource metadata into account during pod scheduling, and so only pods that explicitly request this new hardware through the corresponding ResourceClass name will be allocated with such resources.</br>

- I want a mechanism where it is possible to offer a group of devices, which are co-located on a single node and share a common property, as a single resource that can be requested in pod container spec. Example, N GPU units interconnected by NVLink or N cpu cores on same NUMA node.<br/>
**Motivation:** Provides an infrastructure building block to allow more flexible resource scheduling, through which people can get more optimal use of resources.</br>
**How Resource classes can solve this:**  Property/attribute which forms the grouping can be advertised in the device attributes and then a resource can be created to form a grouped super-resource based on that property.<br/>
**Can this be solved without resource classes:** No

## Objectives
Essentially, a ResourceClass object maps a non-native compute resource with a specific set of properties to a portable name. Cluster admins can create different ResourceClass objects with the same generic name on different clusters to match the underlying hardware configuration. Users can then use the portable names to consume the matching compute resources. Through this extra abstraction layer, we are hoping to achieve the following goals:
- **Allows workloads to request compute resources with wide range of properties in simple and standard way.** We propose to introduce a new `ComputeResource` API field in `Node.Status` to store a list of `ComputeResource` objects. Kubelet can encapsulate the resource metadata information associated with its underlying physical compute resources and propagate this information to the scheduler by appending it to the `ComputeResource` list in the `Node.Status`. With the resource metadata information, the Kubernetes scheduler can determine the fitness of a node for a container resource request expressed through ResourceClass name by evaluating whether the node has enough unallocated ComputeResource matching the property constraints specified in the ResourceClass. This allows the Kubernetes scheduler to take resource property into account to schedule pods on the right node whose hardware configuration meets the specific resource metadata constraints.
- **Allows cluster admins to configure and manage different kinds of non-native compute resources in flexible and simple ways.** A cluster admin creates a ResourceClass object that specifies a portable ResourceClass name (e.g., `fast-nic-gold`), and list of property matching constraints (e.g., `resourceName in (solarflare.com/fast-nic intel.com/fast-nic)`, or `type=XtremeScale-8000`, or `bandwidth=100`). The property matching constraints follow the generic LabelSelector format, which allows us to cover a wide range of resource specific properties. The cluster admin can then define and manage resource quota with the created ResourceClass object.
- **Allows vendors to export their resources on Kubernetes more easily.** The device plugin that a vendor needs to implement to make their resource available on Kubernetes lives outside Kubernetes core repository. The device plugin API will be extended to pass device properties from device plugin to Kubelet. Kubelet will propagate this information to the scheduler through ComputeResource and the scheduler will match a ComputeResource with certain properties to the best matching ResourceClass, and support resource requests expressed through ResourceClass name. The device plugin only needs to retrieve device properties through some device-specific API or tool, without needing to watch or understand either ComputeResource objects or ResourceClass objects.
- **Provides a unified interface to interpret compute resources across various system components such as Quota and Container Resource Spec.** By introducing ResourceClass as a first-class API object, we provide a built-in solution for users to define their special resource constraints through this API, to request such resources through the existing Container Resource Spec, to limit access for such resources through the existing resource Quota component and to ensure their Pods land on the nodes with the matching physical resources through the default Kubernetes scheduling.
- **Supports node-level resources as well as cluster-level resources.** Certain types of compute resources are tied to single nodes and are only accessible locally on those nodes. On the other hand, some types of compute resources such as network attached resources, can be dynamically bound to a chosen node that doesn’t have the resource available till the binding finishes. For such resources, the metadata constraints specified through ResourceClass can be consumed by a standalone controller or a scheduler extender during their resource provisioning or scheduler filtering so that the resource can be provisioned properly to meet the specified metadata constraints.
- **Supports cluster auto scaling for extended resources.** We have seen challenges on how to make cluster autoscaler work seamlessly with dynamically exported extended resources. In particular, for node level extended resources that are exported by a device plugin, cluster autoscaler needs to know what resources will be exported on a newly created node, how much of such resources will be exported and how long it will take for the resource to be exported on the node. Otherwise, it would keep creating new nodes for the pending pod during this time gap. For cluster level extended resources, their resource provisionings are generally performed dynamically by a separate controller. Cluster autoscaler needs to be taught to filter out the resource requests for such resources for the pending pod so that it can create right type of node based on node level resource requests. Note that Kubelet and the scheduler have the similar need to ignore such resource requests during their `PodFitsResources` evaluation. As we are introducing the new resource API that can be used to export arbitrary resource metadata along with extended resources, we need to define a general mechanism for cluster autoscaler to learn the upcoming resource property and capacity on a new node and ensure a consistent resource evaluation policy among cluster autoscaler, scheduler and Kubelet.
- **Defines an easy and seamless migration path** for clusters to adopt ResourceClass even if they have existing pods requesting compute resources through raw resource name. In particular, suppose a cluster is running some workloads that have non native compute resources, such as `nvidia.com/gpu`, in their pod resource requirements. Such workloads should still be scheduled properly when the cluster admin creates a ResourceClass that matches the gpu resources in the cluster. Furthermore, we want to support the upgrade scenario that new resource properties can be added for a resource, e.g., through device plugin upgrade and cluster admins may define a new ResourceClass based on the newly exported resource properties without breaking the use of old ResourceClasses.

## Non Objectives
- Extends the current resource requirement API of the container spec. The current resource requirement API is basically a “name:value” list. A commonly arising question is whether we should extend this API to support resource metadata requirements. We decide to not include this API change in this proposal for the following reasons. First, in a large cluster, computing operators like “greater than”, “less than” during pod scheduling can be a very slow operation and is not scalable. It can cause scaling issues on the scheduler side. Second, non-primary compute resources usually lack standard resource properties. Although there are benefits to allow users to directly express their resource metadata requirements in their container spec, it may also compromise workload portability in longer term. Third, resource quota control will become harder. That is because the current quota admission handler is very simple and just watches for Pod updates and does simple resource request counting to see whether its resource requests in a given namespace is beyond any specified limit or not. If we add resource property selector directly in ContainerSpec (Nvidia non-upstream approach), we will need to extend the current resource quota spec and the quota admission handler quite a lot. It wil also be quite tricky to make sure all pod resource requests are properly quota controlled with the multiple matching behavior allowed by the use of resource property selectors. Fourth, we may consider the resource requirement API change as a possible extension orthogonal to the ResourceClass proposal. By introducing ResourceClass as an additional resource abstraction layer, users can express their special resource requirements through a high-level portable name, and cluster admins can configure compute resources properly on different environments to meet such requirements. We feel this helps promote portability and separation of concerns, while still maintains API compatibility.
- Unifies with the StorageClass API. Although ResourceClass shares many similar motivations and requirements as the existing StorageClass API, they focus on different kinds of resources. StorageClass is used to represent storage resources that are stateful and contains special storage semantics. ResourceClass, on the other hand, focuses on stateless compute resources, whose usage is bound to container lifecycle and can’t be shared across multiple nodes at the same time. For these reasons, we don’t plan to unify the two APIs.
- Resource overcommitment, fractional resource requirements, native compute resource (i.e., cpu and memory) with special metadata requirements, and group compute resources. They are out of our current scope.

## Components
### ResourceClass API
During the initial phase, we propose to start with the following ResourceClass API spec that defines the basic ResourceClass name to the underlying node level compute resource plus metadata matching. We can extend this API to better support cluster resources and group resources in the following phases of development. However, this document will mostly focus on the design to support initial phase requirements.

```golang
// +nonNamespaced=true
// +genclient=true

type ResourceClass struct {
        metav1.TypeMeta
        metav1.ObjectMeta
        Spec ResourceClassSpec
        // +optional
}

type ResourceClassSpec struct {
        // raw resource name. E.g.: nvidia.com/gpu
        ResourceName string
        // defines general resource property matching constraints.
        // e.g.: zone in { us-west1-b, us-west1-c }; type: k80
        MetadataRequirements metav1.LabelSelector
        // used to compare preference of two matching ResourceClasses
        // The purpose to introduce this field is explained more later
        Priority int
}
```

YAML example 1:
```yaml
kind: ResourceClass
metadata:
  name: nvidia.high.mem
spec:
  resourceName: "nvidia.com/nvidia-gpu"
  labelSelector:
    - matchExpressions:
        - key: "memory"
          operator: "GtEq"
          values:
            - "15G"

kind: Pod
metadata:
  name: example-pod
spec:
  containers:
    - name: example-container
      resources:
        limits:
          nvidia.high.mem: 2

```
Above resource class will select all the nvidia-gpus which have memory greater
than and equal to 30 GB.

YAML example 2:
```yaml
kind: ResourceClass
metadata:
  name: fast.nic
spec:
  resourceName: "sfc.com/smartNIC"
  labelSelector:
    - matchExpressions:
        - key: "speed"
          operator: "GtEq"
          values:
            - "40GBPS"

kind: Pod
metadata:
  name: example-pod
spec:
  containers:
    - name: example-container
      resources:
        limits:
          fast.nic: 1
```
Above resource class will select all the NICs with speed greater than equal to
40 GBPS.

Possible fields we may consider to add later include:
- `AutoProvisionConfig`. This field can be used to specify resource auto provisioning config in different cloud environments.
- `Scope`. Indicate whether it maps to node level resource or cluster level resource. For cluster level resource, scheduler, Kubelet, and cluster autoscaler can skip the PodFitsResources predicate evaluation. This allows consistent resource predicate evaluation among these components.
- `ResourceRequestParameters`. This field can be used to indicate special resource request prameters that device plugins may need to perform special configurations on their devices to be consumed by workload pods requesting this resource.

Note we intentially leave these fields out of the initial design to limit the scope
of this proposal.

### Kubelet Extension
On node level, extended resources can be exported automatically by a third-party plugin through the Device Plugin API. We propose to extend the current Device Plugin API that allows device plugins to send to the Kubelet per-device properties during device listing. Exporting device properties at per-device level instead of per-resource level allows a device plugin to manage devices with heterogeneous properties.

After receiving device availability and property information from a device plugin, Kubelet needs to propagate this information to the scheduler so that scheduler can take resource metadata into account when it is making the scheduling decision. We propose to add a new `ComputeResource` list field in NodeStatus API, where each instance in the list represents a device resource and the associated resource properties. Once a node is configured to support ComputeResource API and the underlying resource is exported as a ComputeResource, its quantity should NOT be included in the conventional NodeStatus Capacity/Allocatable fields to avoid resource multiple counting. During the initial phase, we plan to start with exporting extended resources through the ComputeResource API but leaves primary resources in its current exporting model. We can extend the ComputeResource model to support primary resources later after getting more experience through the initial phase. Kubelet will update ComputeResources field upon any resource availability or property change for node-level resources.

We propose to start with the following struct definition:

```golang
type NodeStatus struct {
        …
        ComputeResources []ComputeResource
        …
}
type ComputeResource struct {
        // unique and deterministically generated. “nodeName-resourceName-propertyHash” naming convention,
        // where propertyHash is generated by calculating a hash over all resource properties
        Name string
        // raw resource name. E.g.: nvidia.com/nvidia-gpu
        ResourceName string
        // resource metadata received from device plugin.
        // e.g., gpuType: k80, zone: us-west1-b
        Properties map[string]string
        // list of deviceIds received from device plugin.
        // e.g., ["nvida0", "nvidia1"]
        Devices []string
        // similar to the above but only contains allocatable devices.
        AllocatableDevices []string
}
```
The ComputeResource name needs to be unique and deterministically generated. We propose to use “nodeName-resourceName-propertyHash” naming convention for node level resources, where propertyHash is generated by calculating a hash over all resource properties. The properties of a physical resource may change, e.g., through a driver or node upgrade. When this happens, Kubelet will need to create a new ComputeResource object and delete the old one. However, some pods may already be scheduled on the node and allocated with the old ComputeResource, and it is hard for the scheduler to do its bookkeeping when such pods are still running on the node while the associated ComputeResource is removed. We have a few options to handle this situation.
- First, we may document that to change or remove a device resource, the node has to be drained. This requirement however may complicate device plugin upgrade process.
- Another option is that Kubelet can evict the pods that are allocated with an unexisting ComputeResource. Although simple, this approach may disturb long-running workloads during device plugin upgrade.
- To support a less disruptive model, upon resource property change, Kubelet can still export capacity at old ComputeResource name for the devices used by active pods, and exports capacity at new matching ComputeResource name for devices not in use. Only when those pods finish running, that particular node finishes its transition. This approach avoids resource multiple counting and simplifies the scheduler resource accounting. One potential downside is that the transition may take quite long process if there are long running pods using the resource on the nodes. In that case, cluster admins can still drain the node at convenient time to speed up the transition. Note that this approach does add certain code complexity on Kubelet DeviceManager component.

Possible fields we may consider to add later include:
- `DeviceUnits resource.Quantity`. This field can be used to support fractional
  resource or infinite resource. In a more advanced use case, a device plugin may
  even advertise a single Device with X DeviceUnits so that it can make its own
  device allocation decisions, although this usually require the device plugin
  to implement its own complex logic to track resource life cycle.
- `Owner string`. Can be Kubelet or some cluster-level controller to indicate
  the ownership and scope of the resource.
- `IsolationGuarantee string`. Can map to "ContainerLevel", or "PodLevel", or
  "NodeLevel" to support resource sharing with different levels of isolation
  guarantees.

Note we intentially leave these fields out of the initial design to limit the scope
of this proposal.

### Scheduler Extension
The scheduler needs to watch for NodeStatus ComputeResources field changes and ResourceClass object updates and caches the binding information between the ResourceClass and the matchingComputeResources so that it can serve container resource request expressed through ResourceClass names.

A natural question is how we should define the matching behavior. Suppose there are two ResourceClass objects. ResourceClass RC1 has metadata matching constraint “property1 = value1”, and ResourceClass RC2 has metadata matching constraint “property2 = value2”. Suppose a ComputeResource has both “property1: value1” and “property2: value2” properties, and so match both ResourceClasses. Now should the scheduler consider this ComputeResource as qualified for both ResourceClasses RC1 and RC2, or only one of them?

We feel the desired answer to this question may vary across different types of resources, properties and use cases. To illustrate this, lets consider the following example: A GPU device plugin in a cluster with different types of GPUs may be configured to advertise under a common ResourceName as "nvidia.com/nvidia-tesla-k80", at the beginning. To support per GPU type resource quota, cluster admins may define the following ResourceClasses:

```yaml
kind: ResourceClass
metadata:
  name: nvidia-k80
spec:
  resourceName: "nvidia.com/nvidia-tesla-k80"

kind: ResourceClass
metadata:
  name: nvidia-p100
spec:
  resourceName: "nvidia.com/nvidia-tesla-p100"
```

Later on, suppose the cluster admins add a new GPU node group with a new version of GPU device plugin that exports another resource property "Nvlink" which will be set true for nvidia-tesla-p100 GPUs connected through nvlinks. To utilize this new feature, the cluster admins define the following new ResourceClass with nvlink constraints:

```yaml
kind: ResourceClass
metadata:
  name: nvidia-p100-nvlink
spec:
  resourceName: "nvidia.com/nvidia-tesla-p100"
  labelSelector:
        - key: "Nvlink"
          operator: "Eq"
          values:
            - "true"
```

Now we face the question that whether the scheduler should allow Pods requesting
"nvidia-p100" to land on a node in this new GPU node groups. So far, we have
received different feedbacks on this question. In some use cases, users would
like to have minimum matching behavior that as long as the underlying hardware
matches the minimal requirements specified through ResourceClass contraints,
they want to allow Pods to be scheduled on the hardware. On the other hand, some users desire to reserve expensive hardware resources for users who explicitly request them.
We feel both use cases are valid requirements. Allowing a ComputeResource to match
multiple ResourceClasses as long as it matches their matching constraints
perhaps yields least surprising behavior to users and also simplies upgrade
scenario as new resource properties are introduced into the system. Therefore we
support this behavior by default. To also provide an easy way for cluster admins
to reserve expensive compute resources and control their access with resource
quota, we propose to include a Priority field in ResourceClass API.
By default, the value of this field is set to zero, but cluster admins can set
it to a higher value, which would prevent its matching compute resources from
being matched by lower priority ResourceClasses. i.e.,
when a ComputeResource matches multiple ResourceClasses with different Priority values, the scheduler will choose those with the highest Priority.
Supporting multiple ResourceClass matching also makes it easy to ensure that existing pods requesting resources through raw resource name can continue to be scheduled properly when administrators add ResourceClass in a cluster. To guarantee this, the scheduler may just consider raw resource as a special ResourceClass with empty resource metadata constraints and priority higher than any resource class.

Because a ComputeResource can match multiple ResourceClasses, Scheduler and Kubelet need to ensure a consistent view on ComputeResource to ResourceClass request binding. Let us consider an example to illustrate this problem. Suppose a node has two ComputeResources, CR1 and CR2, that have the same raw resource name but different sets of properties. Suppose they both satisfy the property constraints of ResourceClass RC1, but only CR2 satisfies the property constraints of another ResourceClass RC2. Suppose a Pod requesting RC1 is scheduled first. Because the RC1 resource request can be satisfied by either CR1 or CR2, it is important for the scheduler to record the binding information and propagate it to Kubelet, and Kubelet should honor this binding instead of making its own binding decision. This way, when another Pod comes in that requests RC2, the scheduler can determine whether Pod can fit on the node or not, depending on whether the previous RC1 request is bound to CR1 or CR2.

To maintain and propagate ResourceClass to ComputeResource binding information, the scheduler will need to record this information in a newly introduced ContainerSpec field, similar to the existing NodeName field, and Kubelet will need to consume this information. During the initial implementation, we propose to encode the ResourceClass to the underlying compute resource binding information in a new `AllocatedComputeResources` field in ContainerSpec.
```golang
AllocatedComputeResources map[string]AllocatedResourceList
type AllocatedResourceList struct {
       ComputeResourceName string
       Count int32
}
```

For the purpose to support ResourceClass, we will extend the scheduler NodeInfo cache to store ResourceClass to the matching ComputeResource information on the node. For a given ComputeResource, its capacity will be reflected in NodeInfo.allocatableResource with all matching ResourceClass names. This way, the current node resource fitness evaluation will stay most the same. After a pod is bound to a node, the scheduler will choose the matching ComputeResource on the node, and record this information in the mentioned new field. After that, it increases the NodeInfo.requestedResource for all of the matching ResourceClass names of that ComputeResource. Note that if the `AllocatedComputeResource` field is pre-specified, scheduler should honor this binding instead of overwriting it, similar to how it handles pre-specified NodeName.

The matching from a ResourceClass to the underlying compute resources may change
when cluster admins add, delete, or modify a ResourceClass by adding or removing some metadata constraints or changing its priority.
In such cases, as long as scheduler has already assigned the pod to a node with a ComputeResource, it doesn't matter whether the old ResourceClass would be valid or not. As mentioned above, the scheduler would update NodeInfo.requestedResource for all of the matching ResourceClass names of that ComputeResource. E.g., if a ComputeResource matches ResourceClass RC-A first, and the scheduler assign that ComputeResource to two pods, it will have "RC-A: 2" in NodeInfo.requestedResource. Then suppose there is a ResourceClass update comes in and that same ComputeResource now matches another ResourceClass RC-B. The scheduler will have both "RC-A: 2" and "RC-B: 2" in its NodeInfo.requestedResource cache, even though the ComputeResource no longer matches RC-A. This makes sure we will not over-allocate compute resources. And by recording this info in ContainerSpec, scheduler can re-build this cache info after restarts, even though the matching relationship may have changed.

We do notice that keeping track of ResourceClass to the underlying compute
resource binding may bring scaling concern on the scheduler. In particular,
during ResourceClass addition, deletion, and update, the scheduler needs to scan
through all the cached NodeInfo in the cluster to update the cached
ResourceClass to ComputeResource matching information. This can be an expensive
operation when the cluster has a lot of nodes, and during the time, pods can not
be scheduled as the scheduler needs to hold cache lock during its NodeInfo
traversal.

Our proposed plan is to manage the ResourceClass to the underlying compute
resource matching at the scheduler during the initial implementation, so that we have a central place to track this
information.
With the initial implementation, we will add scaling tests to evaluate the
system scaling limits in different dimensions, such as the number of properties
a ComputeResource may expose, the number of devices a ComputeResource may have, the number of nodes in a cluster can have ComputeResource, and the number of ResourceClasses a cluster can have. Based on the performance results we get, we can explore further optimizations such as having a separate controller to watch NodeStatus ComputeResource updates and propagates the cluster-level aggregated compute resource information to the scheduler through a new ComputeResource API object, or limit the dynamic level of updates we allow for ComputeResource property changes (e.g., we may require that ComputeResource property updates on a node requires node drain) and ResourceClass changes (e.g., don't allow existing ResourceClass to be modified).

### Quota Extension
As mentioned earlier, an important goal of introducing ResourceClass is to allow cluster admins to manage non-native compute resource quota more easily. With the current resource quota system, admins can define the hard limit on the overall quantity of pod resource requests through raw resource name. They don’t have a mechanism to express finer-granularity resource request constraints that map to resources with certain metadata constraints. Following the previous GPU example, admins may want to constraint P100 and K80 GPU requests in a given namespace, other than the total GPU resource requests. By allowing admins to express resource quota constraints at ResourceClass level, we can now support more flexible resource quota specification.

This flexibility also brings an interesting question on how we can enable this benefit while also maintain backward compatibility. We have previously discussed the importance to maintain backward compatibility and how we may extend the scheduler so that pods requesting extended resources through raw resource name can still be scheduled properly. Now a natural question is whether ResourceClass quota constraints should also limit how pods’ raw resource requests are bound to available ResourceClass resources. 

Note that resource quota enforcement and resource request fitting validation are performed by different components. The former enforcement is implemented in the Quota admission controller, while the later task is performed by the scheduler. Enforcing quota constraints at scheduling time would require significant amount of change. We hope to avoid this complexity at the initial stage. Instead, we propose to keep resource quota at its current meaning, i.e., a resource quota limits how much resource can be requested by all of the pods in their container specs in a given namespace. We feel this still allows cluster admins to gradually take advantage of ResourceClass through a manageable migration path. Again, let us walk through an example migration scenario to illustrate how cluster admins can gradually use ResourceClass to express finer-granularity quota constraints:

Step 1: Suppose a cluster already has the following quota constraint to limit total number of gpu devices that pods in a given namespace may request:

```yaml
kind: ResourceQuota
metadata:
  name: gpu-quota-example
spec:
  Hard:
    nvidia.com/gpu: “100”
```
Step 2: Now suppose the cluster admin want to enforce finer-granularity resource constraints through ResourceClass. Following the previous example, the admin defines two resource classes, gpu-silver that maps to K80 gpu devices, and gpu-gold that maps to P100 gpu devices. The cluster admin then migrates a small percent, say 10%, of workloads in the namespace to request resources through the created ResourceClass names in e.g., 3:2 ratio. The admin can then modify the existing quota spec as follows to express the current resource constraints:

```yaml
kind: ResourceQuota
metadata:
  name: gpu-quota-example
spec:
  Hard:
    nvidia.com/gpu: “90”
    gpu-silver: "6"
    gpu-gold: “4”
```
Step 3: Now suppose the experiment of using ResourceClass is successful, and the cluster admin have converted all the workloads running in the namespace to request resources through ResourClass name, the admin can then enforce the finer granularity resource quota across all workloads by modifying the quota spec as follows:

```yaml
kind: ResourceQuota
metadata:
  name: gpu-quota-example
spec:
  Hard:
    nvidia.com/gpu: “0”
    gpu-silver: "60"
    gpu-gold: “40”
```

It is easy to notice that we propose to avoid introducing any changes to the existing quota tracking system. It is possible that in the future, we may need to extend the scheduler to enforce certain quota constraints to support more advanced use cases like better batch job scheduling. When that time comes, the ResourceClass API can be extended to directly express quota constraints, but we will leave that discussion outside the current design.

## Roadmap
### Phase 1: Support ResourceClass based scheduling for node level resources
Defines the ComputeResource API, the ResourceClass API, and extends Kubelet, scheduler, and device plugin API to support ResourceClass based resource requirements.

### Phase 2: Support auto scaling and cluster resources
As we are offering more flexibility on introducing new types of compute resources into Kubernetes, we are also seeing more challenges on how to make cluster autoscaling work seamlessly with such wide range, dynamically exported, and probably dynamically bound resources. Here is the list of problems we have seen in the past related to auto scaling and auto provisioning in the presence of extended resources.

- For node level extended resources that are exported by a device plugin, cluster autoscaler needs to know what resources will be exported on a newly created node, how much of such resources will be exported, and how long it will take for the resource to be exported on the node. Otherwise, it would keep creating new nodes for the pending pod during this time gap.
- For cluster level extended resources, their resource provisionings are generally performed dynamically by a separate controller. Cluster autoscaler and auto provisioner need to be taught to filter out the resource requests for such resources for the pending pod so that it can create right type of node based on node level resource requests. Note that Kubelet and the scheduler have the similar need to ignore such resource requests during their PodFitsResources evaluation. By exporting this information through the ResourceClass API, we can ensure a consistent resource evaluation policy among these components.

During the second phase, we will design and implement a general solution to
support both node level and cluster level extended compute resources with cluster autoscaler.
We may introduce a Scope field in the ResourceClass API to indicate whether it
maps to node level resource or cluster level resource. For node level extended resources,
template ComputeResources can be pre-defined at node group level for cluster
autoscaler to evaluate whether scaling up the node group can satisfy pending pod
resource requirements.

### Phase 3: Support group resources
I.e., a ResourceClass can represent a group of resources. E.g., a gpu-super may include two gpu devices with high affinity, an infiniband-super may include a high-performance nic plus 1G memory, etc.
