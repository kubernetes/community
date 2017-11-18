# Node fencing
Status: Pending

Version: Alpha

Implementation Owner: Yaniv Bronhaim (ybronhei@redhat.com)

## Overview
Rationale and architecture for the addition of a fencing mechanism in Kubernetes clusters.

### Motivation
In Kubernetes a cluster can get into a network partition state between nodes to the api server running on the master node (e.g. network device failure). When that happens their nodes’ status is changed to “not ready” by the node controller and the scheduler cannot know the status of pods running on that node without reaching the nodes’ kubelet service. From such scenario, k8s (since 1.8) defines an eviction timeout (forced by the node controller) such that after 5 minutes pods will enter a termination state in the api server and will try to shutdown gracefully when possible (this happens only if the connectivity returns).

If the pod belongs to a ReplicaSet, once the termination state is set, the ReplicaSet controller will immediately start another instance of that pod.  However when the pod is part of a StatefulSet, the StatefulSet controller won’t start new instance of that pod until the kubelet responds with the pod’s status. This is because there is a permanent association between a StatefulSet member and its storage.  Not waiting would potentially result in multiple copies of a member, all writing to the same volume (leading to corruption or worse).

### Cloud deployments
When k8s is deployed on a cloud such as AWS or GCE, the autoscaler uses the Cloud Provider APIs to remove failed nodes (Failure detection and allowing recovery of StatefulSets is more of a side effect of the autoscaler which is more focused on using the Cloud Provider APIs to monitor load and trigger scale up/down events), effectively bounding the amount of time the node will stay in the “not ready” state and how long the scheduler will need to wait before it can safely start the pod elsewhere.  However in the case of bare metal, no such mechanism exists and recovery will be blocked until an admin intervenes.

This is particularly problematic because all StatefulSet scaling events (up and down) block until existing failures have been recovered.  From some perspectives, it could be considered that each node hosting a StatefulSet member has become a single point of failure for that set.

### Scope of this proposal
Managing HA entities in cluster can relate to nodes and applications. In this scope we cover “node isolation” actions while node is not responsive until it becomes responsive again. Node isolation can be done by power management action (rebooting the node’s machine), storage fencing (preventing node from accessing), and cluster-wide isolation actions.

We don’t cover:
- HA applications
- Collaborative storage based fencing (e.g. sanlock)
- Network fencing

We assume:
Containers that lose access to network and/or storage will self terminate (important for environments relying on network and/or storage fencing)

With isolating the node we benefit by allowing quicker and intensive actions when node that runs stateful applications becomes unreachable. 

### Solution proposal
To address this we propose the addition of the fence controller which, in the event of a node-level connectivity failure, signals to a new fencing executor that isolates the affected node and notifies the apiserver that it is safe to continue recovery.

This functionality, if configured by the admin, applies exclusively to nodes running StatefulSets as they are currently the only construct that provides at-most-one semantics for its members.  In the absence of this feature, an end-user has no ability to safely or reliably allow StatefulSets to be recovered and as such end-users will not be provided with a mechanism to enable/disable this functionality on a set-by-set basis.

Depending on the deployment, the fencing controller will have capabilities such as:
- Power management fencing: powering off or rebooting a node, 
- Storage fencing: disconnection from specific storage to prevent multiple writers), and

Once the node has been made safe by using one or more of the fencing mechanisms listed above, we can know that either the pods are not running anymore or are not accessing shared resources (storage or network) and we can safely delete their objects from the api server to allow the scheduler to initiate new instance.

The design and implementation acknowledge that other entities, such as the autoscaler, are likely to be present and performing similar monitoring and recovery actions. Therefore it is critical that the fencing controller not create or be susceptible to race conditions.

## User experience
The loss of a worker node should be transparent to user's of
StatefulSets.  Recovery time for affected Pods should be bounded and
short, allowing scale up/down events to proceed as normal afterwards.

In the absence of this feature, an end-user has no ability to safely
or reliably allow StatefulSets to be recovered and as such end-users
will not be provided with a mechanism to enable/disable this
functionality on a set-by-set basis.

### Use Cases
1. In a bare metal Kubernetes deployment, StatefulSets should not
   require admin intervention in order to restore capacity when a
   member Pod was active on a lost worker

1. In a bare metal Kubernetes deployment, StatefulSets should not
   require admin intervention in order to scale up or down when a
   member Pod was active on a lost worker

1. In a Cloud Kubernetes deployment, StatefulSets without an
   autoscaler should not require admin intervention in order to scale
   up or down when a member Pod was active on a lost worker

In other words, the failure of a worker node should not represent a
single point of failure for StatefulSets.

## Admin experience
### Configurations required:
- How to trigger fence devices/apis - general template parameters (e.g. cluster UPS address and credentials) and overrides values per node for specific fields (e.g. ports related to node)
- How to run a fence flow for each node (each node can be attached to several fence devices\apis and the fence flow can defer)

The following is done by `ConfigMap` objects:

#### NodeFenceConfig
For each node that supports fencing we will configure NodeFenceConfig - This object is needed to centralize the information about “how” the node can be “fenced” from the cluster - we will have 3 steps for fence “isolation”, “power-management”, “recovery”. In each we will have a list of methods to perform.

Note: to keep those fence configuration we had two options - Adding attributes in metadata section to the existing node object or create ConfigMap that describe this information for each node with fencing options.
```yaml
kind: ConfigMap
apiVersion: v1
metadata:
name: node_fence_config
data: 
  node_fence.properties: | 
node_name=xxx
isolation=[method-storage-iso1]]
power_managment=[method-pm1, method-kdump]
recovery=[method1,method2,method3]
execution_policy=[seq\cun]
```
Notes:
- Each configmap we relate a node name - this should be unique, and in the properties we set order for stages’ names.
- execution_policy means: In list of method we might want to run all methods until all finish, or until only one finishes successfully (useful for sequential pm actions). 
- Isolation\power_managment\recovery are parameters with list value which represent the method to run in each step of the fencing.
#### Fence Method Config ###
Represents the fence method parameters. This is done by configuring a “template” and specific config if needed with overloaded parameters.
- Template is an abstract structure for a fence method (see example below for fence_ipmi device).
- For each template the admin can create specific method config with different properties’ values.
The template allows to reuse parameters for common devices and agents
```yaml
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
version: 1.0
cluster_name: cluster1
spec: 
Name: node1_iscsi_pm_uenf324
template: fence_method_template1
parameters: 
template_param1_name: value
template_param2_name: value
template_param3_name: value
…

kind: FenceMethodTemplate
apiVersion: 1.0
metadata:
version: 1.0
spec: 
Method_template_name: ipmilan_pm_device
agent_name: fence_ipmilan
must_success: Yes\No
parameters*: (default values) 
port:
address:
...
```

### Implementation
#### Fence Controller
The fence controller is a stateless pod instance running in cluster that supports fence. The controller contains two main loop: 1) identifies unresponsive nodes. 2) following fence crds.

First controller will trigger one executor job to live in cluster.

Identification: The controller will identify unresponsive node by check their apiserver object, once node becomes “not ready” a fence treatment will be triggered by posting crd for fence that includes all the information to execute fence stages by Fence Executor. 

Following fence process using crd: The controller will check iteratively all fence objects and validate their execution (base on their status and state in stage). If execution is recognized as stuck or failed, the controller will handle the running job and run a new one on different resource to handle the request. Also, the controller is responsible to move to next “step” when node is still not ready, crd already exists, and the step is running for a configuration amount of time.

#### Fence Executor
Executor is k8s Job that is created in cluster by Fence Controller once the controller starts up. The executor will fetch “fence” crd objects to handle by the following:
- If status:wait- executor reads node’s fence config and specifically the step’s method list, and then start running fence agents using the methods’ config values.
- Any other status means - another executor already process the request, we already fail or success.
- On fail will modify to status:fail and controller will handle the cleanup of the job and retriggering new one to handle the fence request again.

The internal operation performs the fence execution using Fence Agents.

Follwing is a list of agents we will integrate with:
- Cluster fence operation - E.g: 1) cordon node 2) cleaning resources - deleting pods from apiserver.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/aws - Cloud provider agent for rebooting ec2 machines.
- https://github.com/ClusterLabs/fence-agents - Scripts for executing pm devices.
- https://github.com/ClusterLabs/fence-agents/blob/master/fence/agents/compute/fence_compute.py - Fence agent for the automatic resurrection of OpenStack compute instances.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/vmware_soap - Cloud provider agent for rebooting vmware machines.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/rhevm - Cloud provider agent for rebooting oVirt hosts.

Cloud provider allows us to implement fencing agents that perform power management reboot. In k8s autoscaler implementation the concept of cloud provider is already implemented - we will integrate with that code to support PM operations over AWS and GCE.

#### Fence CRD
This is new proposed crd object in k8s cluster API server. The idea behind it is: 1) to allow the fence controller to be “stateless” - means that the crd will hold the fence operation state. 2) to allow triggering executors to perform actions.

In below design we describe fence controller and fence executor, the former triggers fence actions using this object, and should be the only creator for fence crd objects in cluster.

In the following the controller created new fence request for node1

```yaml
kind: Fence
apiVersion: 1.0
metadata:
 - namespace: ..
 - name: ..
 - uid: ..
 ...
spec: 
 - node: node1
 - step: iosolation
 - start_timestep: 19:24:23
 - status: running
 - reources_cleaned: no
 - executor: pod1.dev-8.executor1
status: ..
```
- step is isolation or power_managment or recover - this refer to the set of method configured in the NodeFenceConfig.
- start_timestep allows the executor to know how long operation takes and how to proceed with the fence. 
- status is running while executor starts working, and when it finishes its either fail or success. Before executor picks it the controller creates the object with status:wait.
- resources_cleaned points on the status of the pods cleanup by the controller - if fence is done and resource should be cleaned. Once resource is handled by the fence controller the parameter changes to yes.

For example: controller saw non-response node and created new crd to fence the new, this initialized to “step: isolation” and current timestep. Executor picks this and change status to running. Meanwhile controller check if status was changed to success or failed. 
- If failed will run new executor job to try it
- If finished begin pod treatment - once finish set resources_handled:yes
- If node becomes “ready” in cluster, controller will change the crd to step:recovery and let executor to perform recovery action until done.

After 5min (configurable) if node is still “not ready” controller will change to step:power_managment and back to resources_handled:no and status:wait to let executor pick it up again.

#### Pods treatment
Pod treatment is done by “cluster fence agents” which will be run as part of a node fence treatment.

Kubernetes follows taint-based-evictions. Taints and tolerations are ways to steer pods from nodes or evict pods that should stop.
Pods states are changed once fence flow starts. Therefore, manual status change needs to be done:
Fence pods treatment rules:
- If storage fence was performed, all pods that used this storage can be deleted from the api server.
- If power management for reboot the node ran - all node resources can be removed from cluster - In contrast to the autoscaler, the fence controller cannot control the re-scheduling load once a node is fenced (on scaledown their controller gracefully treat the pods liveness). All resources are released at once, which can lead to overload and scale issues, which happens in parallel to the fence operation.
Default node controller eviction policy doesn’t interfere with this logic. Evictions set the pod for terminating state until node is responsive to perform graceful shutdown. In fencing controller we will delete the pods’ objects only if the node does not get to ready and fenced. This action can trigger autoscaler to scaleup the cluster when big overload is removed from specific node and immediately reschedualed. 
In power management fencing we usually expect the machine to come up again and re-join to the cluster. Therefore, we do not clean the node object, but leave it in “not ready” state bounded until connectivity is returned. 
Note: A fencing method can be also to remove the node from the cluster using the cloud provider api.

## Example configuration
2 node cluster (nodeA - 192.168.1.11, nodeB - 192.168.1.12) with two distinct PDUs (192.168.1.3 - apc, 192.168.1.4 - eaton) using brocade FC switch (192.168.1.2).

Brocade template:
```
kind: FenceMethodTemplate
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Method_template_name: fc_switch_brocade
  - agent_name: fence_brocade
  - parameters:
     - ipaddr: 192.168.1.2
     - password: brocade_password
     - username: brocade_admin
```

APC template:

```
kind: FenceMethodTemplate
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Method_template_name: apc_pdu
  - agent_name: fence_apc_snmp
  - must_sucess: yes
  - parameters:
     - ipaddr: 192.168.1.3
     - password: apc_password
     - username: apc_admin
```

Eaton template:
```
kind: FenceMethodTemplate
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Method_template_name: eaton_pdu
  - agent_name: fence_eaton_snmp
  - must_sucess: yes
  - parameters:
     - ipaddr: 192.168.1.4
     - password: eaton_password
     - username: eaton_admin
     - snmp-priv-prot: AES
     - snmp-priv-passwd: eaton_snmp_passwd
     - snmp-sec-level: authPriv
     - inet4-only
```

Method for enable and disable of brocade for node A:

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_fc_off
  - template: fc_switch_brocade
  - parameters:
     - plug: 1
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_fc_on
  - template: fc_switch_brocade
  - parameters:
     - plug: 1
     - action: on
```

Method for enable and disable of brocade for node B:

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeB_fc_on
  - template: fc_switch_brocade
  - parameters:
     - plug: 2
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeB_fc_off
  - template: fc_switch_brocade
  - parameters:
     - plug: 2
     - action: on
```

Methods for on/off APC (both nodes):

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_apc_off
  - template: apc_pdu
  - parameters:
     - plug: 1
     - action: off
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeB_apc_off
  - template: apc_pdu
  - parameters:
     - plug: 2
     - action: off
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_apc_on
  - template: apc_pdu
  - parameters:
     - plug: 1
     - action: on
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeB_apc_on
  - template: apc_pdu
  - parameters:
     - plug: 2
     - action: on
```

Method for on/off eaton (both nodes)

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_eaton_off
  - template: eaton_pdu
  - parameters:
     - plug: 1
     - action: off
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_eaton_off
  - template: eaton_pdu
  - parameters:
     - plug: 2
     - action: off
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_eaton_on
  - template: eaton_pdu
  - parameters:
     - plug: 1
     - action: on
```

```
kind: FenceMethodConfig
apiVersion: 1.0
metadata:
  - version: 1.0
spec:
  - Name: NodeA_eaton_on
  - template: eaton_pdu
  - parameters:
     - plug: 2
     - action: on
```

Finally fence configs:

```
kind: ConfigMap
apiVersion: v1
metadata:
  - name: nodeA_fence_config
data:
  node_fence.properties:
    - node_name: NodeA
    - isolation: NodeA_fc_off
    - power_management: [NodeA_apc_off, NodeA_eaton_off, nodeA_eaton_on, nodeA_apc_on]
    - recovery: NodeA_fc_on
```

```
kind: ConfigMap
apiVersion: v1
metadata:
  - name: nodeA_fence_config
data:
  node_fence.properties:
    - node_name: NodeB
    - isolation: NodeB_fc_off
    - power_management: [NodeB_apc_off, NodeB_eaton_off, nodeB_eaton_on, nodeB_apc_on]
    - recovery: NodeB_fc_on
```


## Aditional implementations details
### Cluster fence policies
Fencing policy allows each cluster to behave differently in case of connectivity issues. One of the main motivation for that is to prevent “fencing storms” from happening, but there are others. The controller is responsible to force the fence policy

Fencing storm is a situation in which fencing for a few nodes in the cluster is triggered at the same time, due to an environmental issue.

Examples:
1. Switch failure - a switch used to connect a few nodes to the environment is failing, and there is no redundancy in the network environment. In such a case, the nodes will be reported as unresponsive, while they are still alive and kicking, and perhaps providing service through other networks they are attached to. If we fence those nodes, a few services will be offline until restarted, while it might not have been necessary.
2. Management system network failure - if the management system has connectivity issues, it might cause it to identify that nodes are unresponsive, while the issue is with the management system itself.

Some ways to prevent fencing storms:
- Skip fencing if select % of hosts in cluster is non-responsive (will help for issue #2 above)
- Skip fencing if detected the host cannot connect to storage.

### Kdump integration
When kdump enabled is set in NodeFenceConfig the controller will check for kdump notifications once node becomes not ready. Once dumping is recognized, we can delete all pods.
In parallel to wait for kernel dumping the controller will start to execute fence stage normally (admin should take care to configure PM method to run after enough timeout to let the dumping finish).
Kernel dumping is done by booting up node to kdump kernel that starts dumping to hard-coded fqdn  reachable in cluster that save the dumping data.

### Alternatives considered
1. Create a new Cloud Provider allowing the autoscaler to function for
   bare metal deployments.
   
   This was considered however the existing APIs are load balancer
   centric and hard to map to the concept of powering on and off nodes.
   
   If the Cloud Provider API evolves in a compatible direction, it
   might be advisable to persue a Bare Metal provider and have it be
   responsible for much of the fencing configuration.

1. A solution that focused exclusively on power fencing.
   
   While this would dramatically simplify the configuration required,
   many admins see power fencing as a last resort and would prefer
   less destructive way to isolate a misbehaving node, such as network
   and/or disk fencing.
   
   We also see a desire from admins to use tools such as `kdump` to
   obtain additional diagnostics, when possible, prior to powering off
   the node.

1. Attaching fencing configuration to nodes.
   
   While it is tempting to add details on how to fence a node to the
   kubernetes Node objects, this scales poorly from a maintenance
   perspective, preventing nodes from sharing common methods (such as
   `kdump`).
   
   This is especially true for cloud deployments where all nodes are
   controlled with the same credentials. However, even on bare metal
   the only point of differention is often the the IP addresses of the
   IPMI device, or the port number for a network switch, and it would
   be advantageous to manage the rest in one place.

### RBAC rules
### Open questions
- Using https://github.com/kubernetes/node-problem-detector in extend to only check readness by the controller.
### Related proposals

