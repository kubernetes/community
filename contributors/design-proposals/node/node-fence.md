# Node fencing
Status: Pending

Version: Alpha

Implementation Owner: Yaniv Bronhaim (ybronhei@redhat.com)

Current Repository: https://github.com/rootfs/node-fencing

## Overview
Rationale and architecture for the addition of a fencing mechanism in Kubernetes clusters. In [pod-safety](https://github.com/kubernetes/community/blob/16f88595883a7461010b6708fb0e0bf1b046cf33/contributors/design-proposals/pod-safety.md) proposal we describe the need and desire solutions for unrecoverable cluster partition. In the following we define motivations and flows to provide fence actions using fence controller to free the partitioned entities and recover the workload by isolation and power management operations.

### Motivation
Kubernetes cluster can get into a network partition state between nodes to the api server running on the master node (e.g. network device failure). When that happens their nodes’ status is changed to “not ready” by the node controller and the scheduler cannot know the status of pods running on that node without reaching the nodes’ kubelet service. From such scenario, k8s (since 1.8) defines an eviction timeout (forced by the node controller) such that after 5 minutes pods will enter a termination state in the api server and will try to shutdown gracefully when possible (this happens only if the connectivity returns).

If the pod belongs to a ReplicaSet, once the termination state is set, the ReplicaSet controller will immediately start another instance of that pod.  However when the pod is part of a StatefulSet, the StatefulSet controller won’t start new instance of that pod until the kubelet responds with the pod’s status. This is because there is a permanent association between a StatefulSet member and its storage.  Not waiting would potentially result in multiple copies of a member, all writing to the same volume (leading to corruption or worse).

### Cloud deployments
When k8s is deployed on a cloud such as AWS or GCE, the autoscaler uses the Cloud Provider APIs to recognize unhealth nodes and effectively bounding the amount of time the node will stay in the “not ready” state until the node-controller removes the node (if possible, kernel panic can't be handled that way), and how long the scheduler will need to wait before it can safely start the pod elsewhere.
This treatment is not immediate and not covered in bare metal deployments. Also, there is no mechanism to recover quicker, but only to downscale the partitioned nodes. This leads to admin intervention which we can avoid using automate fence controller.
This is particularly problematic because all StatefulSet scaling events (up and down) block until existing failures have been recovered.  From some perspectives, it could be considered that each node hosting a StatefulSet member has become a single point of failure for that set.

### Scope of this proposal
We cover isolation (storage-fence, cluster isolation) and power-management (rebooting the node’s machine) actions while node is partitioned from cluster until it becomes responsive again.
We don’t cover:
- HA applications
- Collaborative storage based fencing (e.g. sanlock)
- Network fencing

We assume:
- Containers that lose access to network and/or storage will self terminate (important for environments relying on network and/or storage fencing)
- With isolating the node we benefit by allowing quicker and intensive actions when node that runs stateful applications becomes unreachable. 

### Solution proposal
To address this we propose the addition of the fence controller which, in the event of a node-level connectivity failure, signals to a new fencing executor that isolates the affected node and notifies the apiserver that it is safe to continue recovery.

This functionality, if configured by the admin, applies exclusively to nodes running StatefulSets as they are currently the only construct that provides at-most-one semantics for its members.  In the absence of this feature, an end-user has no ability to safely or reliably allow StatefulSets to be recovered and as such end-users will not be provided with a mechanism to enable/disable this functionality on a set-by-set basis.

Depending on the deployment, the fencing executors will have capabilities such as:
- Power management fencing: powering off\on\rebooting a node
- Storage fencing: disconnection from specific storage to prevent multiple writers, and unfence when connectivity retuned
- Cluster fence: Cordon node, removing node wordload from api server

Once the node has been made safe by using one or more of the fencing mechanisms listed above, we can know that either the pods are not running anymore or are not accessing shared resources (storage or network) and we can safely delete their objects from the api server to allow the scheduler to initiate new instance.

The design and implementation acknowledge that other entities, such as the autoscaler, are likely to be present and performing similar monitoring and recovery actions. Therefore it is critical that the fencing controller not create or be susceptible to race conditions.

## User experience
The loss of a worker node should be transparent to user's of StatefulSets. Recovery time for affected Pods should be bounded and short, allowing scale up/down events to proceed as normal afterwards.

In the absence of this feature, an end-user has no ability to safely or reliably allow StatefulSets to be recovered and as such end-users will not be provided with a mechanism to enable/disable this functionality on a set-by-set basis.

### Use Cases
1. In a bare metal Kubernetes deployment, StatefulSets should not require admin intervention in order to restore capacity when a member Pod was active on a lost worker
1. In a bare metal Kubernetes deployment, StatefulSets should not require admin intervention in order to scale up or down when a member Pod was active on a lost worker
1. In a Cloud Kubernetes deployment, StatefulSets without an autoscaler should not require admin intervention in order to scale up or down when a member Pod was active on a lost worker

In other words, the failure of a worker node should not represent a single point of failure for StatefulSets.

## Admin experience
### Configurations required
- How to trigger fence devices/apis - general template parameters (e.g. cluster UPS address and credentials) and overrides values per node for specific fields (e.g. ports related to node)
- How to run a fence flow for each node (each node can be attached to several fence devices\apis and the fence flow can defer)

The following is stored in `ConfigMap` objects:

#### Node Fence Config
For each node that supports fencing we will configure NodeFenceConfig - This object is needed to centralize the information about “how” the node can be “fenced” from the cluster - we will have 3 steps for fence “isolation”, “power-management”, “recovery”. In each we will have a list of methods to perform.

Note: 
```yaml
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-config-host1
  data:
   config.properties: |-
    node_name=host1
    isolation=fc-off
    power_management=eaton-off eaton-on
    recovery=fc-on
```
Notes:
- To keep those fence configuration we had two options - Adding attributes in metadata section to the existing node object or create ConfigMap that describe this information for each node with fencing options.
- Each configmap relates to specific node name.
- Isolation\power_managment\recovery are parameters with list value which represent the method to run in each step of the fencing.
- Tha name of the configmap must be "fecne-config-[NODE_NAME]"

#### Fence Method Config ###
Represents the fence method parameters. This is done by configuring a “template” and specific config for each node with overloaded values.
- Template is an abstract structure for a fence method (see examples below for fc_switch device).
- For each template the admin can create specific method config with different properties’ values.
The template allows to reuse parameters for common devices and agents
```yaml
# This fence-method uses the fc-switch-brocade template and adds the plug number for host0,
# when executor reads fence-method-fc-on-host0 for host0, it will also search for the
# conigmap template to add its parameters to the action.
# The name of each fence method must be fence-method-[method_name]-[node_name]
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-fc-on-host0
   namespace: default
  data:
   method.properties: |-
          template=fence-method-template-fc-switch-brocade
          plug=2

# This tamplate configmap for fc-swtich-brocase. The name of templates must be fence-method-template-[method_name]
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-template-fc-switch-brocade
   namespace: default
  data:
   template.properties: |-
          name=fc_switch_brocade
          agent_name=fence_brocade
          ipaddr=192.168.1.2
          password=brocade_password
          username=brocade_admin
```

### Implementation
#### Fence Controller
The fence controller is a stateless controller that can be deployed as pod in cluster or process running outside the cluster. The controller identifies unresponsive node by getting events from the apiserver, once node becomes “not ready” the controller posts crd for fence to initiate fence flows.

The controller proidically polls fencenode crds and manage them as follow:
- status:new - Controller creates job objects for each method in step, move to status:running and update jobs list in nodefence object.
- status:running - Poll running jobs and check if all done successfully. Set status:done or status:error based on jobs condition.
- status:done - Check node readiness, if node is ready move to step:recovery. If node still unresponsive, move to step:power-management. If on step:power-management already and not ready move to status:error (move to error after configurable number of pollings before changing status - see cluster-fence-config)
- status:error - Delete all related jobs and move to status:new to retrigger jobs on different nodes.

Job creation gets the authenticatoin parameters to execute fence agent - This is parsed by the controller from the configmaps as described above. On init the controller reads all fence agents' meta-data to perform parameters extraction for creating the job command. New fence-scripts can be dynamically added by dropping scripts to fence-agents folder and rebuilt the agent image.

#### Executor Job
Executor is k8s Job that is posted to cluster by the controller. The job is based on centos image including all fence scripts that are available in cluster. The job only executes one command and returns the exit status. The job is monitored and mantained by the controller as described in Fence Controller section.
Follwing is a list of agents we will integrate using fence scripts:
- Cluster fence operation - E.g: 1) cordon node 2) cleaning resources - deleting pods from apiserver.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/aws - Cloud provider agent for rebooting ec2 machines.
- https://github.com/ClusterLabs/fence-agents - Scripts for executing pm devices.
- https://github.com/ClusterLabs/fence-agents/blob/master/fence/agents/compute/fence_compute.py - Fence agent for the automatic resurrection of OpenStack compute instances.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/vmware_soap - Cloud provider agent for rebooting vmware machines.
- https://github.com/ClusterLabs/fence-agents/tree/master/fence/agents/rhevm - Cloud provider agent for rebooting oVirt hosts.

Cloud provider allows us to implement fencing agents that perform power management reboot. In k8s autoscaler implementation the concept of cloud provider is already [implemented](https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider) - we might integrate with that code to support PM operations over AWS and GCE.

#### Fence CRD
This is new proposed crd object in k8s cluster API server. The idea behind it is: 
1. to allow the fence controller to be “stateless” - means that the crd will hold the fence operation state and if controller was restarted all the info to continue fence operation will be specified in those objects.
1. to allow triggering jobs to perform actions and sign if they finished successfully or failed.

```yaml
apiVersion: ha.k8s.io/v1
jobs:
- gcloud-reset-inst-9cf8f46e-f44b-11e7-8977-68f728ac95ea
- clean-pods-9d0b1321-f44b-11e7-8977-68f728ac95ea
kind: NodeFence
metadata:
  clusterName: ""
  creationTimestamp: 2018-01-08T08:11:03Z
  generation: 0
  name: node-fence-k8s-cluster-host1.usersys.redhat.com
  namespace: ""
  resourceVersion: "163898"
  selfLink: /apis/ha.k8s.io/v1/node-fence-k8s-cluster-host1.usersys.redhat.com
  uid: 77fdb386-f44b-11e7-bec9-001a4a16015a
node: k8s-cluster-host1.usersys.redhat.com
retries: 2
status: Done
step: Power-Management
```
- step can be isolation\power_managment\recover - this refer to the set of method configured in the NodeFenceConfig.
- status is new, done, running or error - see Fence Controller management for each status.

Flow example: controller saw non-response node and created new crd to fence the new, this initialized to “step: isolation” and current timestep. The controller create jobs related to the step and move status to running.
- If node becomes “ready” in cluster, controller will change the crd to step:recovery and start triggering recovery jobs.

After 5min (configurable) if node is still “not ready” controller will change to step:power_managment and status:new to start triggering pm jobs.

#### Pods treatment
Pod treatment is done by “cluster fence agents” which will be run as part of a node fence treatment.

Kubernetes follows taint-based-evictions. Taints and tolerations are ways to steer pods from nodes or evict pods that should stop.
Pods states are changed once fence flow starts. Therefore, manual status change needs to be done:
Fence pods treatment rules:
- If storage fence was performed, all pods that used this storage can be deleted from the api server.
- If power management for reboot the node ran - all node resources can be removed from cluster.

In contrast to the autoscaler, the fence controller cannot control the re-scheduling load once a node is fenced (on scaledown their controller gracefully treat the pods liveness). All resources are released at once, which can lead to overload and scale issues, which happens in parallel to the fence operation.

Default node controller eviction policy doesn’t interfere with this logic. Evictions set the pod for terminating state until node is responsive to perform graceful shutdown. In pod treatment agents we will delete the pods’ objects only if the node does not get to ready and fenced. This action can trigger autoscaler to scaleup the cluster when big overload is removed from specific node and immediately reschedualed. 

In power management fencing we usually expect the machine to come up again and re-join to the cluster. Therefore, we do not clean the node object, but leave it in “not ready” state bounded until connectivity is returned. 

Note: A fencing method can be also to remove the node from the cluster using the cloud provider api.

## Example configuration
2 node cluster (host0 - 192.168.1.11, host1 - 192.168.1.12) with two distinct PDUs (192.168.1.3 - apc, 192.168.1.4 - eaton) using brocade FC switch (192.168.1.2) for storage isolation.

Brocade template:
```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-template-fc-switch-brocade
   namespace: default
  data:
   template.properties: |-
          name=fc_switch_brocade
          agent_name=fence_brocade
          ipaddr=192.168.1.2
          password=brocade_password
          username=brocade_admin
```

APC template:

```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-template-apc-pdu
   namespace: default
  data:
   template.properties: |-
          name=apc_pdu
          agent_name=fence_apc_snmp
          must_sucess=yes
          ipaddr=192.168.1.3
          password=apc_password
          username=apc_admin

```

Eaton template:
```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-template-eaton-pdu
   namespace: default
  data:
   template.properties: |-
          name=eaton_pdu
          agent_name=fence_eaton_snmp
          must_sucess=yes
          ipaddr=192.168.1.4
          password=eaton_password
          username=eaton_admin
          snmp-priv-prot=AES
          snmp-priv-passwd=eaton_snmp_passwd
          snmp-sec-level=authPriv
          inet4-only=true
```

Method for enable and disable of brocade for host0:

```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-eaton-off-host0
   namespace: default
  data:
   method.properties: |-
        template=fence-method-template-eaton-pdu
        plug=1
        action=off

- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-eaton-on-host0
   namespace: default
  data:
   method.properties: |-
        template=fence-method-template-eaton-pdu
        plug=1
        action=on
```

Method for enable and disable of eaton for host1:

```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-eaton-off-host1
   namespace: default
  data:
   method.properties: |-
        template=fence-method-template-eaton-pdu
        plug=2
        action=off

- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-method-eaton-on-host1
   namespace: default
  data:
   method.properties: |-
        template=fence-method-template-eaton-pdu
        plug=2
        action=on
```

Same can be defined APC and eaton for both nodes.

Finally fence configs:

```
- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-config-lago-kube-host0
   namespace: default
  data:
   config.properties: |-
    node_name=lago-kube-host0
    isolation=fc-off
    power_management=eaton-off eaton-on
    recovery=fc-on

- kind: ConfigMap
  apiVersion: v1
  metadata:
   name: fence-config-host1
  data:
   config.properties: |-
    node_name=host
    isolation=fc-off
    power_management=eaton-off eaton-on
    recovery=fc-on
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
Not defined yet explicitly.

### Open questions
- Using https://github.com/kubernetes/node-problem-detector in extend to only check readness by the controller.
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/cloud-provider/cloud-provider-refactoring.md
