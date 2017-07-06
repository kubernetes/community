# New Resource API Proposal


- [Abstract](#abstract)
- [Background](#background)
- [Use Stories](#user-stories)
  - [As a cluster operator](#as-a-cluster-operator)
  - [As a developer](#as-a-developer)
  - [As a vendor](#as-a-vendor) 
- [Objectives](#objectives)
- [Non Objectives](#non-objectives)
- [Components](#components)
  - [ComputeResource API](#computeResource-api)
  - [ResourceClass API](#resourceclass-api)
- [Kubelet Extension](#kubelet-extension)
- [Scheduler Extension](#scheduler-extension)
- [Quota Extension](#quota-extension)
- [Roadmap](#roadmap)

_Authors:_

* @vikaschoudhary16 - Vikas Choudhary &lt;vichoudh@redhat.com&gt;
* @jiayingz - Jiaying Zhang &lt;jiayingz@google.com&gt;

## Abstract
In this document we will describe a new resource API model to better support non-native compute resources on Kubernetes.

## Background
We are seeing increasing needs to better support non-native compute resources on Kubernetes that cover a wide range of resources such as GPUs, High-performance NICs, Infiniband and FPGAs. Such resources often require vendor specific setup, and have rich sets of different properties even across devices of the same type. This brings new requirements for Kubernetes to better support non-native compute resources to allow vendor specific setup, dynamic resource exporting, flexible resource configuration, and portable resource specification.

The device plugin support added in Kubernetes 1.8 makes it easy for vendors to dynamically export their resources through a plugin API without changing Kubernetes core code. Taking a further step, this document proposes a new resource API model that can be used to describe, manage, and consume such vendor specific and metadata rich resources in simple and portable ways. This model introduces two API objects: ComputeResource API object that represents a type of physical compute resource with a unique set of common properties, and ResourceClass API object that maps from the physical ComputeResources matching certain specified property constraints to a portable name, with which cluster admins can further define the associated resource policies such as quota and limitRange.

## Use Stories
### As a cluster operator:
- Nodes in my cluster has GPU HW from different generations. I want to classify GPU nodes into one of the three categories, silver , gold and platinum depending upon the launch timeline of the GPU family eg: Kepler K20, K80, Pascal P40, P100, Volta V100. I want to charge each of the three categories differently. I want to offer my clients 3 GPU rates/classes to choose from.<br/>
**Motivation:** As time progresses in a cluster lifecycle, new advanced, high performance, expensive variants of GPUs gets added to the cluster nodes. At the same time older variants also co-exist. There are workloads which strictly wants latest GPUs and also there are workloads which are fine with older GPUs. But since there is a wide range of types, it will be hard to manage and confusing at the same time to have granularity at each GPU type. Grouping into few broad categories will be convenient to manage.<br/>
**Can this be solved without resource classes:** A unique taint can be used to represent a category like silver. Nodes can be tainted accordingly depending upon the type of GPUs availability. User pods can use tolerations to steer workloads to the appropriate nodes. But problem is how to restrict a user pod from not using the toleration that it should not be using?<br/>
**How Resource classes can solve this:** I, operator/admin, creates three resource classes: GPU-Platinum, GPU-Gold, GPU-Silver. Now since resource classes are quota controlled, end-user will be able to request resource classes only if quota is allocated.

- I want a mechanism where it is possible to offer a group of devices, which are co-located on a single node and shares a common property, as a single resource that can be requested in pod container spec. Example, N GPU units interconnected by NVLink or N cpu cores on same NUMA node.<br/>
**Motivation:** Increased performance because of local reference. Local reference also helps better use of cache<br/>
**How Resource classes can solve this:**  Property/attribute which forms the grouping can be advertised in the device attributes and then a resource can be created to form a grouped super-resource based on that property.<br/>
**Can this be solved without resource classes:** No

- I want to have quota control on the devices at the granularity of device properties. For example, I want to have a separate quota for ECC enabled GPUs. I want a specific user to not let use more than ‘N’ number of ECC enabled GPUs overall at namespace level.<br/>
**Motivation:**  This will make it possible to charge user per specialized hw consumption. Since special HW is costly, as an Operator i want to have this capability.<br/>
**How Resource classes can solve this:** Quota will be supported on resource class objects and by allowing resource request in user pods via resource class, charging policy can be linked with resource consumption.<br/>
**Can this be solved without resource classes:** No

- In my cluster, i have many different classes(different capabilities) of a device type (ex: NICs). End user’s expectations are met as long as device has a very small subset of these capabilities. I want a mechanism where end user can request devices which satisfies their minimum expectation.
Few nodes are connected to data network over 40 Gig NICs and others are connected over normal 1 Gig NICs. I want end user pods to be able to request:
Data network connectivity with high network performance 
In default case, data network connectivity via normal 1 Gbps NICs<br/>
**Motivation:** If some workloads demand higher network bandwidth, it should be possible to run these workloads on selected nodes.<br/>
**Can this be solved without resource classes:** Taints and tolerations can help in steering pods but the problem in that there is no way today to have access control over use of tolerations and therefore if multiple users are there, it is not possible to have control on allowed tolerations.<br/>
**How Resource classes can solve this:** I can define a ResourceClass for the high-performance NIC with minimum bandwidth requirements, and makes sure only users with proper quota can use such resources.

- I want to be able to utilize different 'types' of a HW resource while not losing workload portability when moving from one cluster to another. There can be Nvidia GPUs on one cluster and AMD GPUs on another cluster. This is example of different ‘types’ of a HW resource(GPU). I want to offer GPUs to be consumed under a same portable name, as long as their capabilities are almost same. If pods are consuming these GPUs with a generic resource class name, workload can be migrated from one cluster to another transparently.<br/>
**Motivation:**  less downtime, optimal use of resources<br/>
**How Resource classes can solve this:** Explained above<br/>
**Can this be solved without resource classes:** No

- I want to be able to query or list extended resources in my cluster through the standard Kubernetes API or CLI.<br/>
**Motivation:**  Improves introspectability of non-primary compute resources.<br/>
**How the new ComputeResource API can solve this:** By introducing ComputeResource API objects, we allow users with proper access to list all compute resources across the cluster to have a cluster level overview or get a list of ComputeResource objects with certain properties through standard Kubernetes API operations. This allows easy introspection of compute resources.
  
### As a developer:
- I want the ability to be able to request devices which have specific capabilities. Eg: GPUs that are Volta or newer.<br/>
  **Motivation:** I want minimum guaranteed compute performance<br/>
  **Can this be solved without resource classes:**<br/>
  - Yes, using node labels and NodeLabelSelectors.
    Problem: Same problem of lack of access control on using labelselectors at user level as with the use of tolerations. 
  - OR, Instead of using resource class, provide flexibility to query resource properties directly in pod container resource requests.
    Problem: In a large cluster, computing operators like “greater than”, “less than” at pod creation can be a very slow operation and is not scalable. 

  **How Resource classes can solve this:**
The Kubernetes scheduler is the central place to map container resource requests expressed through ResourceClass names to the underlying qualified physical resources, which automatically supports metadata aware resource scheduling.

- As a data scientist, I want my workloads to use advanced compute resources available in the running clusters without understanding the underlying hardware configuration details. I want the same workload to run on either on-prem Kubernetes clusters or on cloud, without changing its pod spec. When a new hardware driver comes out, I hope all the required resource configurations are handled properly by my cluster operators and things will just continue to work for any of my existing workloads.<br/>
**Motivation:** Separation of concerns between cluster operators and cluster users.<br/>
**Can this be solved without resource classes:**<br/>
Without the additional abstraction layer, consuming the non-standard, metadata-rich compute resources would be fragmented. More likely, we would see cluster providers implement their own solutions to address their user pains, and it would be hard to provide a consistent user experience for consuming extended resources in the future.

### As a vendor:
- I want an easy and extensible mechanism to export my resource to Kubernetes. I want to be able to roll out new hardware features to the users who require those features without breaking users who are using old versions of hardware.<br/>
**Motivation:** enables more compute resources and their advanced features on Kubernetes<br/>
**Can this be solved without resource classes:**<br/>
Yes, Using node labels and NodeLabelSelectors.<br/>
Problem: Lack of access control and lack of the ability to differentiate between hardware properties on the same node. E.g., if on the same node, some GPU devices are connected through nvlink while others are connected through PCI-e, vendors don’t have ways to export such resource properties that can have very different performance impacts.<br/>
**How Resource classes can solve this:**<br/>
Vendors can use DevicePlugin API to propagate new hardware features, and provide best-practice ResourceClass spec to consume their new hardware or new hardware features on Kubernetes. Vendors don’t need to worry supporting this new hardware would break existing use cases on old hardware because the Kubernetes scheduler takes the resource metadata into account during pod scheduling, and so only pods that explicitly request this new hardware through the corresponding ResourceClass name will be allocated with such resources.

## Objectives
By introducing this new API model, we are hoping to achieve the following objectives:
- **Allows workloads to request compute resources with wide range of properties in simple and standard ways.** By introducing ComputeResource API object, Kubelet can encapsulate the resource metadata information into the corresponding ComputeResource object, and propagate this information to the scheduler. With this information, the Kubernetes scheduler can determine the fitness of a node for a container resource request expressed through ResourceClass name by evaluating whether the node has enough capacity of a ComputeResource matching the property constraints specified in the ResourceClass. This allows the Kubernetes scheduler to take resource property into account to schedule pods on the right node whose hardware configuration meets the specific resource metadata constraints.
- **Allows cluster admins to configure and manage different kinds of non-native compute resources in flexible and simple ways.** A cluster admin creates a ResourceClass object that specifies a portable ResourceClass name (e.g., `fast-nic-gold`), and list of property matching constraints (e.g., `resourceName in (solarflare.com/fast-nic, intel.com/fast-nic)`, or `type=XtremeScale-8000`, or `bandwidth=100G`, or `zone in (us-west1-b, us-west1-c)`). The property matching constraints follow the generic LabelSelector format, which allows us to cover a wide range of resource specific properties. The cluster admin can then define and manage resource quota with the created ResourceClass object.
- **Allows vendors to export their resources on Kubernetes more easily.** The device plugin that a vendor needs to implement to make their resource available on Kubernetes lives outside Kubernetes core repository. The device plugin API will be extended to pass device properties from device plugin to Kubelet. Kubelet will propagate this information to the scheduler through ComputeResource and the scheduler will match a ComputeResource with certain properties to the best matching ResourceClass, and support resource requests expressed through ResourceClass name. The device plugin only needs to retrieve device properties through some device-specific API or tool, without needing to watch or understand either ComputeResource objects or ResourceClass objects.
- **Provides a unified interface to interpret compute resources across various system components such as Quota and CLI.** By introducing ComputeResource API objects, we allow users to list all compute resources across the cluster to have a cluster level overview or get a list of ComputeResource objects with certain properties through standard Kubernetes API operations. This allows easy introspection of compute resources.
- **Supports node-level resources as well as cluster-level resources.** Certain types of compute resources are tied to single nodes and are only accessible locally on those nodes. On the other hand, some types of compute resources such as network attached resources, can be dynamically bound to a chosen node that doesn’t have the resource available till the binding finishes. For such resources, the metadata constraints specified through ResourceClass can be consumed by a standalone controller or a scheduler extender during their resource provisioning or scheduler filtering so that the resource can be provisioned properly to meet the specified metadata constraints.
- **Supports cluster auto scaling for extended resources.** We have seen challenges on how to make cluster autoscaler work seamlessly with dynamically exported extended resources. In particular, for node level extended resources that are exported by a device plugin, cluster autoscaler needs to know what resources will be exported on a newly created node, how much of such resources will be exported, and how long it will take for the resource to be exported on the node. Otherwise, it would keep creating new nodes for the pending pod during this time gap. For cluster level extended resources, their resource provisionings are generally performed dynamically by a separate controller. Cluster autoscaler needs to be taught to filter out the resource requests for such resources for the pending pod so that it can create right type of node based on node level resource requests. Note that Kubelet and the scheduler have the similar need to ignore such resource requests during their PodFitsResources evaluation. By exporting this information through the ResourceClass API, we can ensure a consistent resource evaluation policy among these components.
- **Defines an easy and seamless migration path** for clusters to adopt ResourceClass even if they have existing pods requesting compute resources through raw resource name. In particular, suppose a cluster is running some workloads that have non native compute resources, such as `nvidia.com/gpu`, in their pod resource requirements. Such workloads should still be scheduled properly when the cluster admin creates a ResourceClass that matches the gpu resources in the cluster.


## Non Objectives
- Extends the current resource requirement API of the container spec. The current resource requirement API is basically a “name:value” list. A commonly arising question is whether we should extend this API to support resource metadata requirements. We can consider this as a possible extension orthogonal to the ComputeResource and ResourceClass proposal. A primary reason we propose to introduce ComputeResource and ResourceClass APIs first is because non-native compute resources usually lack standard resource properties. Although there are benefits to allow users to directly express their resource metadata requirements in their container spec, it may also compromise workload portability if not used carefully. By introducing ResourceClass as an additional resource abstraction layer, users can express their special resource requirements through a high-level portable name, and cluster admins can configure compute resources properly on different environments to meet such requirements. We feel this helps promote portability and separation of concerns, while still maintains API compatibility.
- Unifies with the StorageClass API. Although ResourceClass shares many similar motivations and requirements as the existing StorageClass API, they focus on different kinds of resources. StorageClass is used to represent storage resources that are stateful and contains special storage semantics. ResourceClass, on the other hand, focuses on stateless compute resources, whose usage is bound to container lifecycle and can’t be shared across multiple nodes at the same time. For these reasons, we don’t plan to unify the two APIs.
- Resource overcommitment, fractional resource requirements, native compute resource (i.e., cpu and memory) with special metadata requirements, and group compute resources. They are out of our current scope.

## Components
### ComputeResource API
We propose to introduce a new ComputeResource API object that represents a type of physical compute resource with unique sets of properties. We can start with the following simple spec:

```golang
struct ComputeResource {
  metav1.ObjectMeta
  Spec ComputeResourceSpec
  Status ComputeResourceStatus
}
struct ComputeResourceSpec {
     // For node-level resource, this is the name of the Node that owns this ComputeResource.
    // For cluster-level resource, this field is empty.
    NodeName string
     // Raw resource name. E.g.: nvidia.com/gpu
    ResourceName string
    // gpuType: k80
    // zone: us-west1-b
   // Note Kubelet adds a special property corresponding to the above ResourceName field.
   // This will allow a single ResourceClass (e.g., “gpu”) to match multiple types of 
   // resources (e.g., nvidia.com/gpu and amd.com/gpu) through general set selector.
   Properties map[string]string
}
struct ComputeResourceStatus {
   Capacity resource.Quantity
}
```

ComputeResource objects are similar to PV objects on storage. It is tied to the physical resource that can have a wide range of vendor specific properties. Kubelet will create or update ComputeResource objects upon any resource availability or property change for node-level resources. Once a node is configured to support ComputeResource API and the underlying resource is exported as a ComputeResource, its quantity should NOT be included in the conventional NodeStatus Capacity/Allocatable fields to avoid resource multiple counting. ComputeResource API can be included into NodeStatus to facilitate resource introspection.
For cluster level resources, a special controller or a scheduler extender can create a ComputeResource and dynamically bind that to a node during or after scheduling.

### ResourceClass API
ComputeResource mostly serves as an internal API to communicate resource states between different Kubernetes components, like Kubelet and the Scheduler. We will introduce ResourceClass as an additional abstraction layer for better portability and easy management, similar to why StorageClass was introduced to address similar requirements on storage. The scheduler will need to track both ResourceClass and ComputeResource objects to support resource requests expressed through ResourceClass names. 

During the initial phase, we propose to start with the following ResourceClass API spec that defines the basic ResourceClass name to the underlying compute resource plus metadata matching. We can extend this API to include auto provisioning configurations and support group resources in the following phases of development. However, this document will mostly focus on the design to support initial phase requirements.

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
  labelSelector:
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
  name: fast.nic
spec:
  labelSelector:
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

### Kubelet Extension
Kubelet needs to be extended to create or update ComputeResource objects upon any resource availability or property change for node-level resources. During the initial phase, we plan to start with exporting extended resources through the ComputeResource API but leaves primary resources in its current exporting model. We can extend the ComputeResource model to support primary resources later after getting more experience through the initial phase.

On node level, extended resources can be exported automatically by a third-party plugin through the Device Plugin API. We will extend the current Device Plugin API to allow device plugins to send to the Kubelet per-device properties during device listing. Exporting device properties at per-device level instead of per-resource level allows a device plugin to manage devices with heterogeneous properties.

After receiving device availability and property information from a device plugin, Kubelet will create a ComputeResource object to represent the device resource and the associated resource properties. Once a node is configured to support ComputeResource API and the underlying resource is exported as a ComputeResource, its quantity should NOT be included in the conventional NodeStatus Capacity/Allocatable fields to avoid resource multiple counting. ComputeResource API can be included into NodeStatus to facilitate resource introspection.

For cluster level resources, a special controller or a scheduler extender can create a ComputeResource and dynamically bind that to a node during or after scheduling.

The ComputeResource name needs to be unique and deterministically generated. We propose to use “nodeName-resourceName-propertyHash” naming convention for node level resources, where propertyHash is generated by calculating a hash over all resource properties. The properties of a physical resource may change, e.g., through a driver or node upgrade. When this happens, Kubelet will need to create a new ComputeResource object and delete the old one. However, some pods may already be scheduled on the node and allocated with the old ComputeResource, and it is hard for the scheduler to do its bookkeeping when such pods are still running on the node while the associated ComputeResource is removed. We have a few options to handle this situation.
- First, we may document that to change or remove a device resource, the node has to be drained. This requirement however will complicate device plugin upgrade process.
- Another option is that Kubelet can evict the pods that are allocated with an unexisting ComputeResource. Although simple, this approach may disturb long-running workloads during device plugin upgrade.
- To support a less disruptive model, upon resource property change, Kubelet can still export capacity at old ComputeResource name for the devices used by active pods, and exports capacity at new matching ComputeResource name for devices not in use. Only when those pods finish running, that particular node finishes its transition. This approach avoids resource multiple counting and simplifies the scheduler resource accounting. One potential downside is that the transition may take quite long process if there are long running pods using the resource on the nodes. In that case, cluster admins can still drain the node at convenient time to speed up the transition. Note that this approach does add certain code complexity on Kubelet DeviceManager component.

### Scheduler Extension
The scheduler needs to watch for ComputeResource objects and ResourceClass objects, and caches the binding information between the ResourceClass and the matching ComputeResource objects to serve container resource request expressed through ResourceClass names.

A natural question is how we should define the matching behavior. Suppose there are two ResourceClass objects. ResourceClass RC1 has metadata matching constraint “property1 = value1”, and ResourceClass RC2 has metadata matching constraint “property2 = value2”. Suppose a ComputeResource has both “property1: value1” and “property2: value2” properties, and so match both ResourceClasses. Now should the scheduler consider this ComputeResource as qualified for both ResourceClasses RC1 and RC2, or only one of them?
There are a few possible ways to solve multiple matching problem:
- One way to deal with this problem is to disallow ambiguous matching on ResourceClasses. i.e., two ResourceClasses can’t define non overlapping property constraints. This way, we can order ResourceClasses based on their number of constraints and match a ComputeResource to the ResourceClass with most constraints. Although we can enforce this policy during ResourceClass creation, we feel this requirement may complicate the use of ResourceClass, given that the underlying resources can have wide range of properties across different generations of products.
- Another option is to allow a ComputeResource to match any ResourceClasses as long as it qualifies their metadata constraints. Although this approach sounds natural, it may create deployment complexities. Suppose a cluster already has certain existing physical resources and the corresponding ResourceClasses matching these physical resources. Now suppose the cluster is provisioned with a new type of physical resources that can be distinguished through a new property but they also have other properties that match certain existing ResourceClasses. To prevent this new type of resources from being scheduled by existing workloads accidentally, the cluster admin would need to modify the existing ResourceClasses to not match the new property explicitly.
- Proposed solution: introduce a special Priority field in the ResourceClass spec. When a ComputeResource matches multiple ResourceClasses, the scheduler will choose the one with the highest Priority. With this approach, even though we allow cluster admins to create multiple ResourceClasses that may potentially match the same physical resource, the underlying resource can only be requested through a single ResourceClass name (we do allow the underlying resource to be requested through the raw Resource name for backward compatibility that we will discuss in more detail later). A valid question is whether users may desire the flexibility to request the same physical resource through multiple ResourceClass names. From our past experience, non-native computing resources are generally more expensive and users usually prefer to reserving such resources for workloads that really need them. I.e., ambiguous resource matching is generally depreferred. Therefore, we favor this approach for its simplicity, more deterministic scheduling behavior, and the model is relatively easy to understand by users. If multiple matching is indeed desired, users have the option to define multiple ResourceClasses with the same priority, in which case the scheduler will consider a ComputeResource to match all these ResourceClasses.

Because a ComputeResource object can match multiple ResourceClasses, Scheduler and Kubelet need to ensure a consistent view on ComputeResource to ResourceClass request binding. Let us consider an example to illustrate this problem. Suppose a node has two ComputeResources, CR1 and CR2, that have the same raw resource name but different sets of properties. Suppose they both satisfy the property constraints of ResourceClass RC1, but only CR2 satisfies the property constraints of another ResourceClass RC2. Suppose a Pod requesting RC1 is scheduled first. Because the RC1 resource request can be satisfied by either CR1 or CR2, it is important for the scheduler to record the binding information and propagate it to Kubelet, and Kubelet should honor this binding instead of making its own binding decision. This way, when another Pod comes in that requests RC2, the scheduler can determine whether Pod can fit on the node or not, depending on whether the previous RC1 request is bound to CR1 or CR2.

To maintain and propagate ComputeResource to ResourceClass binding information, the scheduler will need to record this information in a newly introduced PodSpec field, similar to the existing NodeName field, and Kubelet will need to consume this information. Note that if this field is pre-specified, scheduler should honor this binding instead of overwriting it, similar to how it handles pre-specified NodeName.

As mentioned earlier, we want to ensure that existing pods requesting resources through raw resource name can continue to be scheduled properly when administrators add ResourceClass in a cluster. To guarantee this, the scheduler may consider raw resource as a special ResourceClass that matches the ComputeResources of that raw resource.

Cluster admins may want to modify an existing ResourceClass, e.g., adding or removing some metadata constraints or changing its priority, or deleting some ResourceClasses that are still requested by some pods. By making the scheduler the single central place to make the ResourceClass to ComputeResource binding decision, and record the binding information in a PodSpec field, ResourceClass updates should have little impact on the pods already being scheduled, and the scheduler can make further scheduling decisions based on its current ResourceClass state and ComputeResource allocation state.

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
### Phase 1: Support ResourceClass based scheduling for node resources
Defines the ComputeResource API, the ResourceClass API, and extends Kubelet, scheduler, and device plugin API to support ResourceClass based resource requirements.

### Phase 2: Support auto provisioning and cluster resources
As we are offering more flexibility on introducing new types of compute resources into Kubernetes, we are also seeing more challenges on how to make auto provisioning work seamlessly with such wide range, dynamically exported, and probably dynamically bound resources. Here is the list of problems we have seen in the past related to auto scaling and auto provisioning in the presence of extended resources.

- For node level extended resources that are exported by a device plugin, cluster autoscaler needs to know what resources will be exported on a newly created node, how much of such resources will be exported, and how long it will take for the resource to be exported on the node. Otherwise, it would keep creating new nodes for the pending pod during this time gap.
- For cluster level extended resources, their resource provisionings are generally performed dynamically by a separate controller. Cluster autoscaler and auto provisioner need to be taught to filter out the resource requests for such resources for the pending pod so that it can create right type of node based on node level resource requests. Note that Kubelet and the scheduler have the similar need to ignore such resource requests during their PodFitsResources evaluation. By exporting this information through the ResourceClass API, we can ensure a consistent resource evaluation policy among these components.

During the second phase, we propose to extend the ResourceClass API with the following fields that can convey the resource provisioning related information.

```golang
ProvisionConfig {
	// CLUSTERLEVEL or NODELEVEL.
    // CLUSTERLEVEL resource can be ignored by scheduler, Kubelet, and
    // Cluster Autoscaler during their node resource fitness evaluation.
	ResourceType string
	// Resource specific node label keys attached on a node with the resource. The
	// value of such a node label indicates the expected quantity of the resource on the node.
	// It may contain multiple node label keys if the given ResourceClass matches multiple
	// kinds of node resources. E.g., we may have different node labels for different types
	// of GPU. A resource class matching all such GPUs should list all their node label keys.
	// Cluster autoscaler can use this field in combination with resource allocatable
    // to determine node readiness for the given resource, and the expected quantity
    // of the resource on the node.
	CapacityNodeLabelKeys []string
	// Represents latency upper bound for the resource to be exported on a new node.
	// Cluster Autoscaler can wait up to this bound before it considers the node not viable
	// for the resource.
	ProvisionSeconds int64
}
```
### Phase 3: Support group resources
i.e., a ResourceClass can represent a group of resources. E.g., a gpu-super may include two gpu devices with high affinity, an infiniband-super may include a high-performance nic plus 1G memory, etc.


