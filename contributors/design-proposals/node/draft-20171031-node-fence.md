# Node fence

## Abstract
Fencing in kubernetes is an additional set of API resources which are processed by managments pods in cluster. Those resources will map to fencing actions (e.g. power managment) on nodes in cases of un-known nodes' states.

## Motivation
Kubernetes cluster lacks the ability to perform power management fencing on nodes. In kubernetes 1.8 every pod can define different behaviour once the node it runs on becomes not responsive (see Pods eviction). Power management fencing allows HA configurations and management. With fencing control the admin can define more complex behaviour once node changes its connectivity status.

## Use Cases
Managing HA entities in cluster can relate to nodes and applications. In this scope we cover “node isolation” actions while node is not responsive until it becomes responsive again. Node isolation can be done by power management action (rebooting the node’s machine), storage fencing (preventing node from accessing), and cluster-wide isolation actions (pod evictions, kdump collection).

We do not cover:
- HA applications\pods management
- Collaborative storage based fencing (e.g. sanlock)

With isolating the node we benefit by allowing quicker and intensive actions when node that runs sensitive applications become unreachable.

## Configurations
Configurations for fencing includes:
1. How to communicate with the fence device\api
2. How to perform the fencing

For each node the admin can define fencing configurations which contain:
1. The Parameters related to this node to execute the fencing operation (cereditionals, addresses, ports..):

Those will be saved using `ConfigMap` objects in the following format:
```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  - name: fence_method
data:
  template: [template_name]
  parameter1: [value]
  parameter2: [value]
  ...
  parametern: [value]
```
```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  - name: fence_method_template_name
data:
  method_name: [agent the method will execute - e.g. fence_ipmilan]
  must_success: [yes\no]
  // default values for parameters that can overrided by specific fence_metho
  parameter1: [value]
  parameter2: [value]
  ...
  parametern: [value]
```

2. What steps to perform and when fencing a node:

Those will be saved using `ConfigMap` objects in the following format:
```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  - name: node_fence_config
data:
  node_name: [node_name]
  fence_stage1:[stage_name]
  fence_stage2:[stage_name]
  ...
  recover_stage:[stage_name]
```

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  - name: fence_stage_[name]
data:
  priority: 5
  execution_policy: [all\only one]
  fence_method1: [method_name]
  fence_method2: [method_name]
  ...
  fence_methodn: [method_name]
```


## Use cases
1. The admin of the cluster takes care for configuring FenceStageConfigs for each node that contains NodeFenceConfig. Node without NodeFenceConfig configmap won’t be managed by the FenceController.
2. The controller reacts when nodes with NodeFenceConfig becomes non-responsive, following the FenceStageConfig list until finishing Recover stage once host becomes responsive again.

### Fence Executor
1 or more pods instances that orchestrate on different physical nodes in cluster and can reach all PM devices in cluster. (note: If one of the PM devices is not reachable, the cluster status should be INACTIVE).
1. Executor reads from k8s api server “Fencing” CRD objects.
2. Executor reads the FenceMethodConfig list specified in the Fencing object.
3. Executor runs over FenceMethodConfig, creates agents related to each of them and executes agents following the executionPolicy specified in the Fencing object.
4. When done, Executor change the status in the Fencing object.

The internal operation performs the fence execution using Fence Agents.
List of agents implementations:
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/aws
- https://github.com/ClusterLabs/fence-agents
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/vmware_soap
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/rhevm

#### Executor Fail Handling
On fencing fail status will change to Error and another Executor in cluster will be picked it up and continue the process.
- Active passive method in executor to prevent re-run on failed executor (possibilities: jobs or active/active or with active/passive way to mark current node as passive)

### Fence Controller
One pod instance that can run anywhere in cluster. The controller reacts when node become non-responsive, following the FenceStageConfig list until finishing recover stage.
- Controller reads from k8s api server “node” objects that contains NodeFenceConfig attached to them.
- Controller keeps information about the managed node and their connectivity to the PVs in cluster.
- Once Node X changes status from active to not_active:
 - Controller validates conditions “Cluster fencing policy”.
 - Controller starts performing FenceMethodStages by creating Fencing CRD object.
 - Controller monitors status of the Fencing object until Done.
 - When stage is completed, controller creates next Fencing request with the following stage data.
 - Once is Active, performing Recover stage.

### Pods Eviction
Kubernetes follows taint-based-evictions. Taints and tolerations are ways to steer pods from nodes or delete (evict) pods that “should not” run.

Eviction for pods should be done as part of the fencing flow. Therefore, manual eviction should be done by an agent that runs as part of a stage that needs to evict pods by a certain rule (such as after storage fencing, we will run agent that evict all pods that are connected to the fenced storage).

Evictions rules:
- If storage fence was performed, all pods that used this storage can be evicted (The pods will be in fail state even if node becomes active again, so deleting them after storage fence is safe)
- If node stays NotActive after configured amount of time we might create a fence method that also cleans all node resources.
The orchestration need to take care for disabling default cluster eviction policy for nodes with fence config. This can be done by declaring tolerationSeconds var in the pods’ definitions.

### Cluster fence policies
Fencing policy allows each cluster to behave differently in case of connectivity issues. One of the main motivation for that is to prevent “fencing storms” from happening, but there are others. The controller is responsible to force the fence policy

Fencing storm is a situation in which fencing for a few nodes in the cluster is triggered at the same time, due to an environmental issue.

Examples:
- Switch failure - a switch used to connect a few nodes to the environment is failing, and there is no redundancy in the network environment. In such a case, the nodes will be reported as unresponsive, while they are still alive and kicking, and perhaps providing service through other networks they are attached to. If we fence those nodes, a few services will be offline until restarted, while it might not have been necessary.
- Management system network failure - if the management system has connectivity issues, it might cause it to identify that nodes are unresponsive, while the issue is with the management system itself.

Some ways to prevent fencing storms:
- Skip fencing if select % of hosts in cluster is non-responsive (will help for issue #2 above)
- Skip fencing if the host maintains a storage lease (will help for issue #1 above)

### Kdump integration
When kdump enabled is set in NodeFenceConfig the controller will check for kdump notifications once node becomes not active. Once dumping is recognized, we can evict all pods.
In parallel to wait for kernel dumping the controller will start to execute fence stage normally (admin should take care to configure PM method to run after enough timeout to let the dumping finish).
Kernel dumping is done by booting up node to kdump kernel that starts dumping to hard-coded fqdn  reachable in cluster that save the dumping data.

### API Resource
The NodeFence resource will be added to the main API v1:

```go
package v1
// NodeFencing is the node fencing object accessible to the fencing controller and executor
type NodeFencing struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`

	// Node represents the node to be fenced.
	// +optional
	Node core_v1.Node `json:"node" protobuf:"bytes,2,opt,name=node"`

	// Status represents the latest observer state of the node fencing
	// +optional
	Status NodeFencingStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
```

A `Registry` implementation for `NodeFencing` will be added to TODO.

This API allows the executing and monitoringfence operations by the fence controller on specific node.

### RBAC Rules
- Only fencing controller and agent can access the Fencing object
- Only executor reads NodeFenceConfigs
- Only controller reads FenceStageConfigs

### Examples

### Open Questions
- Recovery stage - how unfence can be performed and when. Before booting up fenced node we must unfence the storage from the node.
- Add links to related k8s documentation

### Related Proposals
- [pod-safety](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/0000-kep-template.md).
-
