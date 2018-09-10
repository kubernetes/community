# Cluster Scoped Resources

## Abstract
Cluster scoped resources are consumable resources that do not belong to any specific node but instead are available across mulitple nodes in a cluster. These resources are accounted as other consumable resources and should be usable by the scheduler while deciding if a pod can actually be scheduled.


## Motivation
Resources in Kubernetes such as cpu and memory are available at a node level and can be consumed by pods by requesting them. However there are some resources that do not belong to a specific node, but they are consumable across all or a group of nodes in the cluster. Few such use cases are mentioned below.

#### Use Cases
1. Software Licenses that can be shared by pods across the entire cluster

2. IP Addresses:
   Nodes in a cluster can be partitioned into multiple network scopes and each network scope can have a certain number of available IPs that can be assigned to pods running on those nodes. Hence, the nodes belonging to a network scope can collectively run as many pods as the IP capacity in that scope.

3. Rack storage:
   Locally attached shared storage in a rack, which is consumable by pods on nodes within a rack

4. Network Bandwidth:
   Network bandwidth shared by pods. Depending on the network topology of a cluster, pods on multiple nodes will be sharing network bandwidth with each other. In use cases where pods need guranteed network throughput, representing bandwidth as a cluster resource is essential for scheduling such pods.

## Goals
The goal is to define mechanisms to expose and consume cluster scoped resources

## Design

### ClusterResource type
```
// pkg/api/types.go:

// ClusterResourceQuantity represents quantity of a ClusterResource
type ClusterResourceQuantity struct {
	Quantity resource.Quantity `json:"quantity"`
	// NodeSelector is a label query over nodes which collectively provide resource Quantity
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

// ClusterResource represents a resource which is available at a cluster level
type ClusterResource struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status ClusterResourceStatus `json:"status,omitempty"`
}

type ClusterResourceStatus struct {
	// Capacity represents the total quantity of ClusterResource
	// +optional
	Capacity []ClusterResourceQuantity `json:"capacity,omitempty"`
	// Allocatable represents the quantity of ClusterResource that is available for scheduling
	// +optional
	Allocatable []ClusterResourceQuantity `json:"allocatable,omitempty"`
}
```
`ClusterResourceStatus` captures the capacity and allocatable quantity for a `ClusteResource` in the form of `ClusterResourceQuantity`. `ClusterResourceQuantity` represents the quantity of a `ClusterResource` which is collectively consumable on nodes selected by `NodeSelector`.
`NodeSelector` is a label query over the nodes, which collectively provide this `ClusterResource`. This field is optional, and if not specfied, it means that the `ClusterResource` is consumable across all nodes in the cluster.


### Consuming ClusterResources

ClusterResources are consumable by pods just like CPU and memory, by specifying it in the pod request. The scheduler should take care of the resource accounting for ClusterResources so that no more than the available amount is simultaneously allocated to Pods. The prefix used to identify a ClusterResource coule be 
```
pod.alpha.kubernetes.io/cluster-resource-
```

### Accounting in scheduler

ClusterResources should be tracked as normal consumable resources and should be considered by the scheduler when determining if a pod can actually be scheduled

```
// kubernetes/plugin/pkg/scheduler/schedulercache/cluster_info.go

// ClusterInfo is cluster level aggregated information.
type ClusterInfo struct {

	clusterResources map[string]*ClusterResource
}

// kubernetes/plugin/pkg/scheduler/schedulercache/cache.go

type schedulerCache struct {
   	... 
   	cluster   *ClusterInfo
}
```

`clusterinfo` is added to scheduler cache to do accounting for ClusterResources consumed by pods. `clusterInfo` will be exposed to the predicate and priority functions in order to take ClusterResources into consideration while making scheduling decisions.
