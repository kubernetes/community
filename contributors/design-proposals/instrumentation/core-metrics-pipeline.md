# Core Metrics in kubelet

**Author**: David Ashpole (@dashpole)

**Last Updated**: 1/31/2017

**Status**: Proposal

This document proposes a design for the set of metrics included in an eventual Core Metrics Pipeline.


- [Core Metrics in kubelet](#core-metrics-in-kubelet)
  - [Introduction](#introduction)
    - [Definitions](#definitions)
    - [Background](#background)
    - [Motivations](#motivations)
    - [Proposal](#proposal)
    - [Non Goals](#non-goals)
  - [Design](#design)
    - [Metric Requirements:](#metric-requirements)
    - [Proposed Core Metrics:](#proposed-core-metrics)
    - [On-Demand Design:](#on-demand-design)
  - [Future Work](#future-work)


## Introduction

### Definitions
"Kubelet": The daemon that runs on every kubernetes node and controls pod and container lifecycle, among many other things.  
["cAdvisor":](https://github.com/google/cadvisor) An open source container monitoring solution which only monitors containers, and has no concept of kubernetes constructs like pods or volumes.  
["Summary API":](https://git.k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1/types.go) A kubelet API which currently exposes node metrics for use by both system components and monitoring systems.  
["CRI":](/contributors/devel/sig-node/container-runtime-interface.md) The Container Runtime Interface designed to provide an abstraction over runtimes (docker, rkt, etc).  
"Core Metrics": A set of metrics described in the [Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md) whose purpose is to provide metrics for first-class resource isolation and utilization features, including [resource feasibility checking](https://github.com/eBay/Kubernetes/blob/master/docs/design/resources.md#the-resource-model) and node resource management.
"Resource": A consumable element of a node (e.g. memory, disk space, CPU time, etc).  
"First-class Resource": A resource critical for scheduling, whose requests and limits can be (or soon will be) set via the Pod/Container Spec.  
"Metric": A measure of consumption of a Resource.  

### Background
The [Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md) proposal contains a blueprint for a set of metrics referred to as "Core Metrics".  The purpose of this proposal is to specify what those metrics are, to enable work relating to the collection, by the kubelet, of the metrics.

Kubernetes vendors cAdvisor into its codebase, and the kubelet uses cAdvisor as a library that enables it to collect metrics on containers.  The kubelet can then combine container-level metrics from cAdvisor with the kubelet's knowledge of kubernetes constructs (e.g. pods) to produce the kubelet Summary statistics, which provides metrics for use by the kubelet, or by users through the [Summary API](https://git.k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1/types.go).  cAdvisor works by collecting metrics at an interval (10 seconds, by default), and the kubelet then simply queries these cached metrics whenever it has a need for them.

Currently, cAdvisor collects a large number of metrics related to system and container performance. However, only some of these metrics are consumed by the kubelet summary API, and many are not used.  The kubelet [Summary API](https://git.k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1/types.go) is published to the kubelet summary API endpoint (stats/summary).  Some of the metrics provided by the summary API are consumed by kubernetes system components, but many are included for the sole purpose of providing metrics for monitoring.

### Motivations
The [Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md) proposal explains why a separate monitoring pipeline is required.

By publishing core metrics, the kubelet is relieved of its responsibility to provide metrics for monitoring.
The third party monitoring pipeline also is relieved of any responsibility to provide these metrics to system components.

cAdvisor is structured to collect metrics on an interval, which is appropriate for a stand-alone metrics collector.  However, many functions in the kubelet are latency-sensitive (eviction, for example), and would benefit from a more "On-Demand" metrics collection design.

### Proposal
This proposal is to use this set of core metrics, collected by the kubelet, and used solely by kubernetes system components to support "First-Class Resource Isolation and Utilization Features".  This proposal is not designed to be an API published by the kubelet, but rather a set of metrics collected by the kubelet that will be transformed, and published in the future.

The target "Users" of this set of metrics are kubernetes components (though not necessarily directly).  This set of metrics itself is not designed to be user-facing, but is designed to be general enough to support user-facing components.

### Non Goals
Everything covered in the [Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md) design doc will not be covered in this proposal.  This includes the third party metrics pipeline, and the methods by which the metrics found in this proposal are provided to other kubernetes components.

Integration with CRI will not be covered in this proposal.  In future proposals, integrating with CRI may provide a better abstraction of information required by the core metrics pipeline to collect metrics.

The kubelet API endpoint, including the format, url pattern, versioning strategy, and name of the API will be the topic of a follow-up proposal to this proposal.

## Design
This design covers only metrics to be included in the Core Metrics Pipeline.

High level requirements for the design are as follows:
 - The kubelet collects the minimum possible number of metrics to provide "First-Class Resource Isolation and Utilization Features".
 - Metrics can be fetched "On Demand", giving the kubelet more up-to-date stats.

This proposal purposefully omits many metrics that may eventually become core metrics.  This is by design.  Once metrics are needed to support First-Class Resource Isolation and Utilization Features, they can be added to the core metrics API.

### Metric Requirements
The core metrics api is designed to provide metrics for "First Class Resource Isolation and Utilization Features" within kubernetes.

Many kubernetes system components currently support these features.  Many more components that support these features are in development.
The following is not meant to be an exhaustive list, but gives the current set of use cases for these metrics.

Metrics requirements for "First Class Resource Isolation and Utilization Features", based on kubernetes component needs, are as follows:  

 - Kubelet
   - Node-level usage metrics for Filesystems, CPU, and Memory  
   - Pod-level usage metrics for Filesystems and Memory  
 - Metrics Server (outlined in [Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md)), which exposes the [Resource Metrics API](/contributors/design-proposals/instrumentation/resource-metrics-api.md) to the following system components:
   - Scheduler  
     - Node-level usage metrics for Filesystems, CPU, and Memory  
     - Pod-level usage metrics for Filesystems, CPU, and Memory  
   - Vertical-Pod-Autoscaler  
     - Node-level usage metrics for Filesystems, CPU, and Memory  
     - Pod-level usage metrics for Filesystems, CPU, and Memory  
     - Container-level usage metrics for Filesystems, CPU, and Memory  
   - Horizontal-Pod-Autoscaler  
     - Node-level usage metrics for CPU and Memory  
     - Pod-level usage metrics for CPU and Memory  
   - Cluster Federation  
     - Node-level usage metrics for Filesystems, CPU, and Memory  
   - kubectl top and Kubernetes Dashboard  
     - Node-level usage metrics for Filesystems, CPU, and Memory  
     - Pod-level usage metrics for Filesystems, CPU, and Memory  
     - Container-level usage metrics for Filesystems, CPU, and Memory  

### Proposed Core Metrics:
This section defines "usage metrics" for filesystems, CPU, and Memory.  
As stated in Non-Goals, this proposal does not attempt to define the specific format by which these are exposed.  For convenience, it may be necessary to include static information such as start time, node capacities for CPU, Memory, or filesystems, and more.

```go
// CpuUsage holds statistics about the amount of cpu time consumed  
type CpuUsage struct {  
  // The time at which these Metrics were updated.  
  Timestamp metav1.Time  
  // Cumulative CPU usage (sum of all cores) since object creation.  
  CumulativeUsageNanoSeconds *uint64   
}  

// MemoryUsage holds statistics about the quantity of memory consumed  
type MemoryUsage struct {  
  // The time at which these metrics were updated.  
  Timestamp metav1.Time  
  // The amount of "working set" memory. This includes recently accessed memory,  
  // dirty memory, and kernel memory.  
  UsageBytes *uint64  
}   

// FilesystemUsage holds statistics about the quantity of local storage (e.g. disk) resources consumed  
type FilesystemUsage struct {  
  // The time at which these metrics were updated.  
  Timestamp metav1.Time  
  // StorageIdentifier must uniquely identify the node-level storage resource that is consumed.
  // It may utilize device, partition, filesystem id, or other identifiers.  
  StorageIdentifier string  
  // UsedBytes represents the disk space consumed, in bytes.  
  UsedBytes *uint64  
  // UsedInodes represents the inodes consumed  
  UsedInodes *uint64  
}  
```

### On-Demand Design
Interface:  
The interface for exposing these metrics within the kubelet contains methods for fetching each relevant metric.  These methods contains a "recency" parameter which specifies how recently the metrics must have been computed.  Kubelet components which require very up-to-date metrics (eviction, for example), use very low values.  Other components use higher values.  

Implementation:  
To keep performance bounded while still offering metrics "On-Demand", all calls to get metrics are cached, and a minimum recency is established to prevent repeated metrics computation.  Before computing new metrics, the previous metrics are checked to see if they meet the recency requirements of the caller.  If the age of the metrics meet the recency requirements, then the cached metrics are returned.  If not, then new metrics are computed and cached.  

## Future work
Suggested, tentative future work, which may be covered by future proposals:  
 - Decide on the format, name, and kubelet endpoint for publishing these metrics.
 - Integrate with the CRI to allow compatibility with a greater number of runtimes, and to create a better runtime abstraction.   

