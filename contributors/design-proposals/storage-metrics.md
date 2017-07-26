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

The first part of this proposal is to add volume usage information indexed by PVC name in Kubelet Summary API

### volume usage information indexed by PVC

The basic idea is to cache PVC and the volume information in kubelet volume manager which is similar to caching the pod and volume information. In Summary API, we could add new API type PVCStats which is similar to NodeStats and PodStats

```
// PVCStats holds pod-level unprocessed sample volume stats refered by PVC.
type PVCStats struct {
	// Reference to the measured PVC.
	PVCRef PVCReference `json:"podRef"`
	// The time at which data collection for the PVC-scoped volume stats was (re)started.
	StartTime metav1.Time `json:"startTime"`
	// Stats pertaining to volume usage of filesystem resources.
	// VolumeStats.UsedBytes is the number of bytes used by the Volume
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	VolumeStats VolumeStats `json:"volume,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
}
```

## Implementation Timeline:
The feature is targeted for kubernetes v1.8

