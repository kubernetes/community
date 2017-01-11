# Core Metrics Pipeline in kubelet

**Author**: David Ashpole (@dashpole)

**Last Updated**: 1/10/2017

**Status**: Draft Proposal (WIP)

This document proposes a design for an internal Core Metrics Pipeline.

<!-- BEGIN MUNGE: GENERATED_TOC -->

- [Core Metrics Pipeline in kubelet](#core-metrics-pipeline-in-kubelet)
  - [Introduction](#introduction)
    - [Definitions](#definitions)
    - [Background](#background)
    - [Motivations](#motivations)
    - [Proposal](#proposal)
    - [Non Goals](#non-goals)
  - [Design](#design)
    - [Metric Requirements:](#metric-requirements)
    - [Proposed Core Metrics API:](#proposed-core-metrics-api)
  - [Implementation Plan](#implementation-plan)
  - [Rollout Plan](#rollout-plan)
  - [Implementation Status](#implementation-status)

<!-- END MUNGE: GENERATED_TOC -->

## Introduction

### Definitions
"Kubelet": The daemon that runs on every kubernetes node and controls pod and container lifecycle, among many other things.  
["cAdvisor":](https://github.com/google/cadvisor) An open source container monitoring solution which only monitors containers, and has no concept of k8s constructs like pods or volumes.  
["Summary API":](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/stats/types.go) A kubelet API which currently exposes node metrics for use by both system components and monitoring systems.  
["CRI":](https://github.com/kubernetes/community/blob/master/contributors/devel/container-runtime-interface.md) The Container Runtime Interface designed to provide an abstraction over runtimes (docker, rkt, etc).  
"Core Metrics": A set of metrics described in the [Monitoring Architecture](https://github.com/kubernetes/kubernetes/blob/master/docs/design/monitoring_architecture.md) whose purpose is to provide system components with metrics for the purpose of [resource feasibility checking](https://github.com/eBay/Kubernetes/blob/master/docs/design/resources.md#the-resource-model) or node resource management.  

### Background
The [Monitoring Architecture](https://github.com/kubernetes/kubernetes/blob/master/docs/design/monitoring_architecture.md) proposal contains a blueprint for a set of metrics referred to as "Core Metrics".  The purpose of this proposal is to specify what those metrics are, and how they will be collected on the node.

Kubernetes vendors cAdvisor into its codebase, and the kubelet uses cAdvisor as a library that enables it to collect metrics on containers.  The kubelet can then combine container-level metrics from cAdvisor with the kubelet's knowledge of k8s constructs (e.g. pods) to produce the kubelet Summary statistics, which provides metrics for use by the kubelet, or by users through the [Summary API](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/stats/types.go).  cAdvisor works by collecting metrics at an interval (10 seconds, by default), and the kubelet then simply queries these cached metrics whenever it has a need for them.

Currently, cAdvisor collects a large number of metrics related to system and container performance. However, only some of these metrics are consumed by the kubelet summary API, and many are not used.  The kubelet [Summary API](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/stats/types.go) is published to the kubelet summary API endpoint (stats/summary).  Some of the metrics provided by the summary API are consumed by kubernetes system components, but most are not for this purpose.

### Motivations
The [Monitoring Architecture](https://github.com/kubernetes/kubernetes/blob/master/docs/design/monitoring_architecture.md) proposal explains why a separate monitoring pipeline is required.

By publishing core metrics, the summary API is relieved of its responsibility to provide metrics to system components.  This will allow the summary API to evolve into a much richer set of metrics for monitoring, since the overhead burden falls on the third party monitoring pipeline.

cAdvisor is structured to collect metrics on an interval, which is appropriate for a stand-alone metrics collector.  However, many functions in the kubelet are latency-sensitive (eviction, for example), and would benifit from a more "On-Demand" metrics collection design.

### Proposal
I propose to use this set of core metrics, collected by the kubelet, and used solely by kubernetes system compenents to support resource feasibility checking and resource management on the node.

The target "Users" of this set of metrics are kubernetes components.  This set of metrics itself is not designed to be user-facing, but is designed to be general enough to support user-facing components.

### Non Goals
Everything covered in the [Monitoring Architecture](https://github.com/kubernetes/kubernetes/blob/master/docs/design/monitoring_architecture.md) design doc will not be covered in this proposal.  This includes the third party metrics pipeline, and the methods by which the metrics found in this proposal are provided to other kubernetes components.

Integration with CRI will not be covered in this proposal.  In future proposals, integrating with CRI may provide a better abstraction of information required by the core metrics pipeline to collect metrics.

## Design
This design covers only the Core Metrics Pipeline.

High level requirements for the design are as follows:
 - Do not break existing users.  We should continue to provide the full summary API as an optional add-on.  Once the monitoring pipeline is completed, the summary API will be provided by the monitoring pipeline, possibly through a stand-alone version of cAdvisor.
 - The kubelet collects the minimum possible number of metrics for complete portable kubernetes functionalities.
 - Metrics can be fetched "On Demand", giving the kubelet more up-to-date stats.

More details on how I intend to achieve these high level goals can be found in the Implementation Plan.

This Core Metrics API will be versioned to account for version-skew between kubernetes components.

This proposal purposefully omits many metrics that may eventually become core metrics.  This is by design.  Once metrics are needed to support resource feasibility checking or node resource management, they can be added to the core metrics API.

### Metric Requirements
The core metrics api is designed to provide metrics for two use-cases within kubernetes:
 - Resource Feasibility Checking
 - Node Resource Management

Many kubernetes system components currently support these features.  Many more components that support these features are in development.
The following is meant not meant to be an exhaustive list, but gives the current set of use cases for these metrics.

Metrics requirements for resource feasibility checking and node resource management, based on kubernetes component needs, are as follows:
 - Kubelet
  - Node-level capacity and availability metrics for Disk, Memory, and CPU
  - Pod-level usage metrics for Disk and Memory
 - Scheduler (Possibly through [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md))
  - Node-level capacity and availability metrics for Disk, CPU, and Memory
  - Pod-level usage metrics for Disk, CPU, and Memory
  - Container-level usage metrics for Disk, CPU, and Memory
 - Horizontal-Pod-Autoscaler (Exposed through [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md)))
  - Node-level capacity and availability metrics for CPU and Memory
  - Pod-level usage metrics for CPU and Memory
 - Cluster Federation (Exposed through [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md))
  - Node-level capacity and availability metrics for Disk, Memory, and CPU
 - kubectl top (Exposed through [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md))
  - Node-level capacity and availability metrics for Disk, Memory, and CPU
  - Pod-level usage metrics for Disk, Memory, and CPU
  - Container-level usage metrics for Disk, CPU, and Memory
 - Kubernetes Dashboard (Exposed through [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md))
  - Node-level capacity and availability metrics for Disk, Memory, and CPU
  - Pod-level usage metrics for Disk, Memory, and CPU
  - Container-level usage metrics for Disk, CPU, and Memory


### Proposed Core Metrics API:

An important difference between the current summary api and the proposed core metrics api is that per-pod stats in the core metrics api contain only usage data, and not capacity-related statistics.  This is more accurate since a pod's resource capacity is really defined by its "requests" and "limits", and it is a better reflection of how components use these metrics.  The kubelet, for example, finds which resources are constrained using node-level capacity and availability data, and then chooses which pods to take action on based on the pod's usage of the constrained resource.  If neccessary, capacity for resources a pod consumes can still be correlated with node-level resources using this format of stats.

```go
// CoreStats is a top-level container for holding NodeStats and PodStats.  
type CoreStats struct {  
  // Overall node resource stats.  
  Node NodeResources `json:"node"`  
  // Per-pod usage stats.  
  Pods []PodUsage `json:"pods"`  
}  

// NodeStats holds node-level stats.  NodeStats contains capacity and availibility for Node Resources.  
type NodeResources struct {  
  // The filesystem device used by node k8s components.  
  // +optional  
  KubeletFsDevice string `json:"kubeletfs,omitempty"`  
  // The filesystem device used by node runtime components.  
  // +optional  
  RuntimeFsDevice string `json:"runtimefs,omitempty"`  
  // Stats pertaining to cpu resources.  
  // +optional  
  CPU *CpuResources `json:"cpu,omitempty"`  
  // Stats pertaining to memory (RAM) resources.  
  // +optional  
  Memory *MemoryResources `json:"memory,omitempty"`  
  // Stats pertaining to node filesystem resources.  
  // +optional  
  Filesystems []FilesystemResources `json:"filesystems, omitempty" patchStrategy:"merge" patchMergeKey:"device"`  
}  

// CpuResources containes data about cpu resource usage  
type CpuResources struct {  
  // The number of cores in this machine.  
  NumCores int `json:"numcores"`  
  // The current Usage of CPU resources  
  TotalUsage *CpuUsage `json:"cpuusage,omitempty"`  
}  

// MemoryResources contains data about memory resource usage.  
type MemoryResources struct {  
  // The time at which these stats were updated.  
  Timestamp metav1.Time `json:"time"`  
  // The memory capacity, in bytes  
  CapacityBytes *uint64 `json:"capacitybytes,omitempty"`  
  // The available memory, in bytes  
  // This is the number of bytes which are not included in the working set memory
  // Working set memory includes recently accessed memory, dirty memory, and kernel memory.
  AvailableBytes *uint64 `json:"availablebytes,omitempty"`  
}  

// FilesystemResources contains data about filesystem disk resources.  
type FilesystemResources struct {  
  // The time at which these stats were updated.  
  Timestamp metav1.Time `json:"time"`  
  // The device that this filesystem is on  
  Device string `json:"device"`  
  // AvailableBytes represents the storage space available (bytes) for the filesystem.  
  // +optional  
  AvailableBytes *uint64 `json:"availableBytes,omitempty"`  
  // CapacityBytes represents the total capacity (bytes) of the filesystems underlying storage.  
  // +optional  
  CapacityBytes *uint64 `json:"capacityBytes,omitempty"`  
  // InodesFree represents the free inodes in the filesystem.  
  // +optional  
  InodesFree *uint64 `json:"inodesFree,omitempty"`  
  // Inodes represents the total inodes in the filesystem.  
  // +optional  
  Inodes *uint64 `json:"inodes,omitempty"`  
}  

// PodUsage holds pod-level unprocessed sample stats.  
type PodUsage struct {  
  // UID of the pod  
  PodUID string `json:"uid"`  
  // Stats pertaining to pod total usage of cpu  
  // This may include additional overhead not included in container usage statistics.  
  // +optional  
  CPU *CpuUsage `json:"cpu,omitempty"`  
  // Stats pertaining to pod total usage of system memory  
  // This may include additional overhead not included in container usage statistics.  
  // +optional  
  Memory *MemoryUsage `json:"memory,omitempty"`  
  // Stats of containers in the pod.  
  Containers []ContainerUsage `json:"containers" patchStrategy:"merge" patchMergeKey:"uid"`  
  // Stats pertaining to volume usage of filesystem resources.  
  // +optional  
  Volumes []VolumeUsage `json:"volume,omitempty" patchStrategy:"merge" patchMergeKey:"name"`  
}  

// ContainerUsage holds container-level usage stats.  
type ContainerUsage struct {  
  // UID of the container  
  ContainerUID string `json:"uid"`  
  // Stats pertaining to container usage of cpu  
  // +optional  
  CPU *CpuUsage `json:"memory,omitempty"`  
  // Stats pertaining to container usage of system memory  
  // +optional  
  Memory *MemoryUsage `json:"memory,omitempty"`  
  // Stats pertaining to container rootfs usage of disk.  
  // Rootfs.UsedBytes is the number of bytes used for the container write layer.  
  // +optional  
  Rootfs *FilesystemUsage `json:"rootfs,omitempty"`  
  // Stats pertaining to container logs usage of Disk.  
  // +optional  
  Logs *FilesystemUsage `json:"logs,omitempty"`  
}  

// CpuUsage holds statistics about the amount of cpu time consumed  
type CpuUsage struct {  
  // The time at which these stats were updated.  
  Timestamp metav1.Time `json:"time"`  
  // Average CPU usage rate over sample window (across all cores), in "cores".  
  // The "core" unit represents nanoseconds of CPU time consumed per second.  
  // For example, 5 nanocores means the process averaged 5 nanoseconds 
  // of cpu time per second during the sample window.
  // +optional  
  UsageRateNanoCores *uint64 `json:"usageNanoCores,omitempty"`  
  // Cumulative CPU usage (sum of all cores) since object creation.  
  // +optional  
  AggregateUsageCoreNanoSeconds *uint64 `json:"usageCoreNanoSeconds,omitempty"`  
}  

// MemoryUsage holds statistics about the quantity of memory consumed  
type MemoryUsage struct {  
  // The time at which these stats were updated.  
  Timestamp metav1.Time `json:"time"`  
  // The amount of working set memory. This includes recently accessed memory,  
  // dirty memory, and kernel memory.  
  // +optional  
  UsageBytes *uint64 `json:"usageBytes,omitempty"`  
}  

// VolumeUsage holds statistics about the quantity of disk resources consumed for a volume  
type VolumeUsage struct {  
  // Embedded FilesystemUsage  
  FilesystemUsage  
  // Name is the name given to the Volume  
  // +optional  
  Name string `json:"name,omitempty"`  
}  

// FilesystemUsage holds statistics about the quantity of disk resources consumed  
type FilesystemUsage struct {  
  // The time at which these stats were updated.  
  Timestamp metav1.Time `json:"time"`  
  // The device on which resources are consumed  
  Device string `json:"device"`  
  // UsedBytes represents the disk space consumed on the device, in bytes.  
  // +optional  
  UsedBytes *uint64 `json:"usedBytes,omitempty"`  
  // InodesUsed represents the inodes consumed on the device  
  // +optional  
  InodesUsed *uint64 `json:"inodesUsed,omitempty"`  
}  
```

## Implementation Plan

@dashpole will internally separate core metrics from summary metrics and make the kubelet use the core metrics.  
@dashpole will create a separate endpoint TBD to publish this set of core metrics.  
@dashpole will modify volume stats collection so that it relies on this code.   
@dashpole will modify the structure of stats collection code to be "On-Demand".   

Suggested, tentative future work, which may be covered by future proposals:  
 - Obtain all runtime-specific information needed to collect metrics from the CRI.   
 - Modify cAdvisor to be "stand alone", and run in a seperate binary from the kubelet.  It will consume the above metadata API, and provide the summary API.  
 - The kubelet no longer provides the summary API, and starts, by default, cAdvisor stand-alone (which provides the summary API).  Include flag to disable running stand-alone cAdvisor.  

## Rollout Plan
The core metrics endpoint (TBD) will be added alongside the current Summary API for the upcoming release.  This should allow concurrent developments of other portions of the system metrics pipeline (metrics-server, for example).  Once this addition is made, all other changes will be internal, and will not require any API changes.  
Once the [implementation work](#implementation-plan) is completed, @dashpole will start discussions on how to provide the summary API through a means separate from the kubelet.  One current idea is a standalone verison of cAdvisor, but any third party metrics solution could serve this function as well.

## Implementation Status

The implementation goals of the first milestone are outlined below.
- [ ] Create the proposal
- [ ] Implement collection and consumption of core metrics.
- [ ] Create Kubelet API endpoint for core metrics.
- [ ] Modify volume stats collection so that it relies on this code.
- [ ] Modify the structure of stats collection code to be "On-Demand"



<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/core-metrics-pipeline.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
