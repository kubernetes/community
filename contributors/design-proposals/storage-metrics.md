# [WIP] Expose Storage Metrics to Users

**Author**: Jing Xu (@jingxu97)

**Last Updated**: 7/26/2017

**Status**: Proposal

This document explores different ways of exposing storage metrics to users

## Background and Motivation

Monitoring resource usage is critical for users/admins to manage their systems. For storage, there are three levels of metrics:
 - Node-level: the capacity, used, and allocatable storage on local host
 - Pod-level: Pod volume could be network attached (GCE PD, AWS EBS etc.) or local (emptyDir, hostPath, or local storage volume). Each volume has a total capacity and used bytes
 - Container-level: How much storage is used by each container’s writable layer and logs (standard input/output).

All the above metrics are used for operation of Kubernetes internal components and core utilities as explains below.

 - Eviction manager checks the allocatable local storage (scratch space) and take actions. Scheduler also needs the node-level local storage capacity and availability to determine how to place pods on nodes 
 - In 1.7, an alpha feature is added to isolation emptyDir volume usage and container’s overlay in local storage. User can set a sizeLimit for Pod EmptyDir Volume and eviction manager will take action if the emptyDir volume usage exceeds the configured size. Similarly, user can set request/limit on container’s overlay and eviction manager can take actions based on the monitoring data and the limits.
 - There is also a proposal targeting for 1.8 [volume resize proposal](https://github.com/kubernetes/community/pull/657) to dynamically resize the volume (network attached) if the disk resource is insufficient. It requires a monitoring pipeline to dynamically check the the disk usage for each volume (represented by PVC)

### Current Monitoring Options

There are a number of monitoring options supported by Kubernetes
 - Resource Metrics API (Alpha feature in progress [Resource Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md)
 - Collect metrics in Heapster. (metrics has to exposed by Kubelet in Summary API)
 - Export metrics through StackDriver 
 - Kubectl Command (kubectl top, kubectl describe)

All the storage metrics mentioned in Background are collected by Kubelet Summary API so that it could be integrated with the above monitoring system. However, pod volume metrics is provided in PodStat object’s VolumeStat and VolumeStat only has volume name specified in pod specification. There is no direct way of checking volume usage information by PVC name, which is more preferable by users.


## Goals
The gola of this proposal is to expose storage usage metrics to users for monitoring their storage systems.

## Proposal

This proposal has three parts, the first part is trying to address the issue of adding volume usage information indexed by PVC and PV name in Kubelet Summary API. The second part is to register the volume metrics to . The third part is to add more storage metrics to Heapster.


### volume usage information indexed by PVC/PV in Kubelet Summary API

The basic idea is to cache PVC and the volume information in kubelet volume manager which is similar to caching the pod and volume information. In Summary API, In Summary API, we could add PVC and PV references into VolumeStats which is included in PodStats.

```
// VolumeStats contains data about Volume filesystem usage.
type VolumeStats struct {
	// Reference to the measured PVC.
	PVCRef PVCReference `json:"podRef"`
	// Reference to the measured PV.
	PVRef PVReference `json:"podRef"`
	// Embedded FsStats
	FsStats
	// Name is the name given to the Volume
	// +optional
	Name string `json:"name,omitempty"`
}

// PVCReference contains enough information to locate the referenced PVC.
type PVCReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UID       string `json:"uid"`
}

// PVReference contains enough information to locate the referenced PV.
type PVReference struct {
	Name      string `json:"name"`
	UID       string `json:"uid"`
}

```
### Register the VolumeStats metrics to Prometheus 

The following metrics could be registered to Prometheus


| Metric name | Metric type | Labels/tags |
|-------------|-------------|-------------|
| volume_stats_capacityBytes | Gauge | namespace=\<persistentvolumeclaim-namespace\> <br/> persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=\<persistentvolume-name\> |
| volume_stats_usedBytes | Gauge | namespace=\<persistentvolumeclaim-namespace\> <br/>  persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=\<persistentvolume-name\> |
| volume_stats_availableBytes | Gauge | namespace=\<persistentvolumeclaim-namespace\> <br/> persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=<persistentvolume-name> |
| volume_stats_InodesFree | Gauge | nnamespace=\<persistentvolumeclaim-namespace\> <br/> persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=\<persistentvolume-name\> |
| volume_stats_Inodes | Gauge | namespace=\<persistentvolumeclaim-namespace\> <br/> persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=\<persistentvolume-name\> |
| volume_stats_InodesUsed | Gauge | namespace=\<persistentvolumeclaim-namespace\> <br/> persistentvolumeclaim=\<persistentvolumeclaim-name\>  <br/> persistentvolume=\<persistentvolume-name\> |

## Implementation Timeline:
The feature is targeted for kubernetes v1.8

