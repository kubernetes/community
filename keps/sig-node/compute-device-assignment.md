---
kep-number: 18
title: Kubelet endpoint for device assignment observation details 
authors:
  - "@dashpole" 
  - "@vikaschoudhary16"
owning-sig: sig-node
reviewers:
  - "@thockin"
  - "@derekwaynecarr"
  - "@dchen1107"
  - "@vishh"
approvers:
  - "@sig-node-leads"
editors:
  - "@dashpole"
  - "@vikaschoudhary16"
creation-date: "2018-07-19"
last-updated: "2018-07-19"
status: provisional
---
# Kubelet endpoint for device assignment observation details 

Table of Contents
=================
* [Abstract](#abstract)
* [Background](#background)
* [Objectives](#objectives)
* [User Journeys](#user-journeys)
  * [Device Monitoring Agents](#device-monitoring-agents)
* [Changes](#changes)
* [Potential Future Improvements](#potential-future-improvements)
* [Alternatives Considered](#alternatives-considered)

## Abstract
In this document we will discuss the motivation and code changes required for introducing a kubelet endpoint to expose device to container bindings.

## Background
[Device Monitoring](https://docs.google.com/document/d/1NYnqw-HDQ6Y3L_mk85Q3wkxDtGNWTxpsedsgw4NgWpg/edit?usp=sharing) requires external agents to be able to determine the set of devices in-use by containers and attach pod and container metadata for these devices.

## Objectives

* To remove current device-specific knowledge from the kubelet, such as [accellerator metrics](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/apis/stats/v1alpha1/types.go#L229)
* To enable future use-cases requiring device-specific knowledge to be out-of-tree

## User Journeys

### Device Monitoring Agents

* As a _Cluster Administrator_, I provide a set of devices from various vendors in my cluster.  Each vendor independently maintains their own agent, so I run monitoring agents only for devices I provide.  Each agent adheres to to the [node monitoring guidelines](https://docs.google.com/document/d/1_CdNWIjPBqVDMvu82aJICQsSCbh2BR-y9a8uXjQm4TI/edit?usp=sharing), so I can use a compatible monitoring pipeline to collect and analyze metrics from a variety of agents, even though they are maintained by different vendors.
* As a _Device Vendor_, I manufacture devices and I have deep domain expertise in how to run and monitor them. Because I maintain my own Device Plugin implementation, as well as Device Monitoring Agent, I can provide consumers of my devices an easy way to consume and monitor my devices without requiring open-source contributions. The Device Monitoring Agent doesn't have any dependencies on the Device Plugin, so I can decouple monitoring from device lifecycle management. My Device Monitoring Agent works by periodically querying the `/devices/<ResourceName>` endpoint to discover which devices are being used, and to get the container/pod metadata associated with the metrics:

![device monitoring architecture](https://user-images.githubusercontent.com/3262098/43926483-44331496-9bdf-11e8-82a0-14b47583b103.png)


## Changes

Add a v1alpha1 Kubelet GRPC service, at `/var/lib/kubelet/pod-resources/kubelet.sock`, which returns information about the kubelet's assignment of devices to containers. It obtains this information from the internal state of the kubelet's Device Manager. The GRPC Service returns a single PodResourcesResponse, which is shown in proto below:
```protobuf
// PodResources is a service provided by the kubelet that provides information about the
// node resources consumed by pods and containers on the node
service PodResources {
    rpc List(ListPodResourcesRequest) returns (ListPodResourcesResponse) {}
}

// ListPodResourcesRequest is the request made to the PodResources service
message ListPodResourcesRequest {}

// ListPodResourcesResponse is the response returned by List function
message ListPodResourcesResponse {
    repeated PodResources pod_resources = 1;
}

// PodResources contains information about the node resources assigned to a pod
message PodResources {
    string name = 1;
    string namespace = 2;
    repeated ContainerResources containers = 3;
}

// ContainerResources contains information about the resources assigned to a container
message ContainerResources {
    string name = 1;
    repeated ContainerDevices devices = 2;
}

// ContainerDevices contains information about the devices assigned to a container
message ContainerDevices {
    string resource_name = 1;
    repeated string device_ids = 2;
}
```

### Potential Future Improvements

* Add `ListAndWatch()` function to the GRPC endpoint so monitoring agents don't need to poll.
* Add identifiers for other resources used by pods to the `PodResources` message.
  * For example, persistent volume location on disk

## Alternatives Considered

### Add v1alpha1 Kubelet GRPC service, at `/var/lib/kubelet/pod-resources/kubelet.sock`, which returns a list of [CreateContainerRequest](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/apis/cri/runtime/v1alpha2/api.proto#L734)s used to create containers.
* Pros:
  * Reuse an existing API for describing containers rather than inventing a new one
* Cons:
  * It ties the endpoint to the CreateContainerRequest, and may prevent us from adding other information we want in the future
  * It does not contain any additional information that will be useful to monitoring agents other than device, and contains lots of irrelevant information for this use-case.
* Notes:
  * Does not include any reference to resource names.  Monitoring agentes must identify devices by the device or environment variables passed to the pod or container.

### Add a field to Pod Status. 
* Pros:
  * Allows for observation of container to device bindings local to the node through the `/pods` endpoint
* Cons:
  * Only consumed locally, which doesn't justify an API change
  * Device Bindings are immutable after allocation, and are _debatably_ observable (they can be "observed" from the local checkpoint file).  Device bindings are generally a poor fit for status.

### Use the Kubelet Device Manager Checkpoint file
* Allows for observability of device to container bindings through what exists in the checkpoint file
  * Requires adding additional metadata to the checkpoint file as required by the monitoring agent
* Requires implementing versioning for the checkpoint file, and handling version skew between readers and the kubelet
* Future modifications to the checkpoint file are more difficult.

### Add a field to the Pod Spec:
* A new object `ComputeDevice` will be defined and a new variable `ComputeDevices` will be added in the `Container` (Spec) object which will represent a list of `ComputeDevice` objects.
```golang
// ComputeDevice describes the devices assigned to this container for a given ResourceName
type ComputeDevice struct {
	// DeviceIDs is the list of devices assigned to this container
	DeviceIDs []string
	// ResourceName is the name of the compute resource
	ResourceName string
}

// Container represents a single container that is expected to be run on the host.
type Container struct {
    ...
	// ComputeDevices contains the devices assigned to this container
	// This field is alpha-level and is only honored by servers that enable the ComputeDevices feature.
	// +optional
	ComputeDevices []ComputeDevice
	...
}
```
* During Kubelet pod admission, if `ComputeDevices` is found non-empty, specified devices will be allocated otherwise behaviour will remain same as it is today.
* Before starting the pod, the kubelet writes the assigned `ComputeDevices` back to the pod spec.  
  * Note: Writing to the Api Server and waiting to observe the updated pod spec in the kubelet's pod watch may add significant latency to pod startup.
* Allows devices to potentially be assigned by a custom scheduler.
* Serves as a permanent record of device assignments for the kubelet, and eliminates the need for the kubelet to maintain this state locally.

