# Device Manager Proposal

  1. [Abstract](#abstract)
  2. [Motivation](#motivation)
  3. [Use Cases](#use-cases)
  4. [Objectives](#objectives)
  5. [Non Objectives](#non-objectives)
  6. [Stories](#stories)
      * [Vendor story](#vendor-story)
      * [User story](#user-story)
  8. [Device Plugin](#device-plugin)
      * [Protocol Overview](#protocol-overview)
      * [Protobuf specification](#protobuf-specification)
      * [Installation](#installation)
      * [API Changes](#api-changes)
      * [Versioning](#versioning)

_Authors:_

* @RenaudWasTaken - Renaud Gaubert &lt;rgaubert@NVIDIA.com&gt;

## Abstract

This document describes a vendor independant solution to:
  * Discovering and representing external devices
  * Making these devices available to the container and cleaning them up
    afterwards
  * Health Check of these devices

Because devices are vendor dependant and have their own sets of problems
and mechanisms, the solution we describe is a plugin mechanism managed by
Kubelet.

At their core, device plugins are simple gRPC servers that may run in a
container deployed through the pod mechanism.

These servers implement the gRPC interface defined later in this design
document and once the device plugin makes itself know to kubelet, kubelet
will interact with the device through three simple functions:
  1. A `Discover` function for the kubelet to Discover the devices and
     their properties.
  2. An `Allocate` and `Deallocate` function which are called respectively
     before container creation and after container deletion with the
     devices to allocate and deallocate.
  3. A `Monitor` function to notify Kubelet whenever a device becomes
     unhealthy.

The goal is for a user to be able to enable vendor devices (e.g: GPUs) through
the simple following steps:
  * `kubectl create -f http://vendor.com/device-plugin-daemonset.yaml`
  * When launching `kubectl describe nodes`, the devices appear in the node spec
  * In the long term users will be able to select them through Resource Class

We expect the plugins to be deployed across the clusters through DaemonSets.
The targeted devices are GPUs, NICs, FPGAs, InfiniBand, Storage devices, ....


## Motivation

Kubernetes currently supports discovery of CPU and Memory primarily to a
minimal extent. Very few devices are handled natively by Kubelet.

It is not a sustainable solution to expect every vendor to add their vendor
specific code inside Kubernetes. This approach does not scale and is not
portable.

We want a solution for those vendors to be able to advertise their resources
to kubelet and monitor them.
We also want a way for the user to specify which resource their jobs will use
and what constraints are associated to these resources.

In order to solve this problem it is obvious that we need a plugin system in
order to have vendors advertise and monitor their resources on behalf
of Kubelet.

Additionally, we introduce the concept of Device to be able to select
resources with constraints in a pod spec.

_GPU Integration Example:_
  * [Enable "kick the tires" support for NVIDIA GPUs in COS](https://github.com/Kubernetes/Kubernetes/pull/45136)
  * [Extend experimental support to multiple NVIDIA GPUs](https://github.com/Kubernetes/Kubernetes/pull/42116)

_Kubernetes Meeting Notes On This:_
  * [Meeting notes](https://docs.google.com/document/d/1Qg42Nmv-QwL4RxicsU2qtZgFKOzANf8fGayw8p3lX6U/edit#)
  * [Better Abstraction for Compute Resources in Kubernetes](https://docs.google.com/document/d/1666PPUs4Lz56TqKygcy6mXkNazde-vwA7q4e5H92sUc)
  * [Extensible support for hardware devices in Kubernetes (join Kubernetes-dev@googlegroups.com for access)](https://docs.google.com/document/d/1LHeTPx_fWA1PdZkHuALPzYxR0AYXUiiXdo3S0g2VSlo/edit)

## Use Cases

  * I want to use a particular device type (GPU, InfiniBand, FPGA, etc.)
    in my pod.
  * I should be able to use that device without writing custom Kubernetes code.
  * I want a consistent and portable solution to consume hardware devices
    across k8s clusters.

## Objectives

1. Add support for vendor specific Devices in kubelet:
    * Through a pluggable mechanism.
    * Which allows discovery and monitoring of devices.
    * Which allows hooking the runtime to make devices available in containers
      and cleaning them up.
2. Define a deployment mechanism for this new API.
3. Define a versioning mechanism for this new API.

## Non Objectives
1. Advanced scheduling and resource selection (solved through [#782](https://github.com/Kubernetes/community/pull/782)).
   We will only try to give basic selection primitives to the devices
2. Metrics: this should be the job of cadvisor and should probably either be
   addressed there (cadvisor) or if people feel there is a case to be made
   for it being addressed in the Device Plugin, in a follow up proposal.

## Stories

### Vendor story

Kubernetes provides to vendors a mechanism called device plugins to:
  * advertise devices.
  * monitor devices (currently perform health checks).
  * hook into the runtime to instruct Kubelet what are the steps to
    take in order to make the device available (or cleanup the device).

A device plugin at it's core is a simple gRPC server usually running in
a container and deployed across clusters through a daemonSet.

```gRPC
service DevicePlugin {
	rpc Discover(Empty) returns (stream Device) {}
	rpc Monitor(Empty) returns (stream DeviceHealth) {}

	rpc Allocate(AllocateRequest) returns (AllocateResponse) {}
	rpc Deallocate(DeallocateRequest) returns (Empty) {}
}

```

The gRPC server that the device plugin must implement is expected to
be advertised on a unix socket in a mounted hostPath (e.g:
`/var/run/Kubernetes/vendor.sock`).

Finally, to notify Kubelet of the existence of the device plugin,
the vendor's device plugin will have to make a request to Kubelet's
onwn gRPC server.
Only then will kubelet start interacting with the vendor's device plugin
through the gRPC apis.

### End User story

When setting up the cluster the admin knows what kind of devices are present
on the different machines and therefore can select what devices they want to
enable.

The cluster admins knows his cluster has NVIDIA GPUs therefore he deploys
the NVIDIA device plugin through:
`kubectl create -f NVIDIA.io/device-plugin.yml`

The device plugin lands on all the nodes of the cluster and if it detects that
there are no GPUs it terminates. However, when there are GPUs it reports them
to Kubelet.
For device plugins reporting non-GPU Devices these are advertised as
OIRs and selected through the same method.

1. A user submits a pod spec requesting X GPUs (or devices)
2. The scheduler filters the nodes which do not match the resource requests
3. The pod lands on the node and Kubelet decides which device
   should be assigned to the pod
4. Kubelet calls `Allocate` on the matching Device Plugins
5. The user deletes the pod or the pod terminates
6. Kubelet calls `Deallocate` on the matching Device Plugins

When receiving a pod which requests Devices kubelet is in charge of:
  * deciding which device to assign to the pod's containers (this will
    change in the future)
  * advertising the changes to the node's `Available` list
  * advertising the changes to the pods's `Allocated` list
  * Calling the `Allocate` function with the list of devices

The scheduler is still be in charge of filtering the nodes which cannot
satisfy the resource requests.
He might in the future be in charge of selecting the device.

## Device Plugin

### Introduction
The device plugin is structured in 5 parts:
1. Registration: The device plugin advertises it's presence to Kubelet
2. Discovery: Kubelet calls the device plugin to list it's devices
3. Allocate / Deallocate: When creating/deleting containers requesting the
   devices advertised by the device plugin, Kubelet calls the device plugin's
   `Allocate` and `Deallocate` functions.
4. Cleanup: Kubelet terminates the communication through a "Stop"
4. Heartbeat: The device plugin polls Kubelet to know if it's still alive
   and if it has to re-issue a Register request

### Registration

When starting the device plugin is expected to make a (client) gRPC call
to the `Register` function that Kubelet exposes.

The communication between Kubelet is expected to happen only through Unix
sockets and follow this simple pattern:
1. The device plugins starts it's gRPC server
2. The device plugins sends a `RegisterRequest` to Kubelet (through a
   gRPC request)
4. Kubelet starts it's Discovery phase and calls `Discover` and `Monitor`
5. Kubelet answers to the `RegisterRequest` with a `RegisterResponse`
   containing any error Kubelet might have encountered

### Unix Socket

Device Plugins are expected to communicate with Kubelet through gRPC
on an Unix socket.
When starting the gRPC server, they are expected to create a unix socket
at the following host path: `/var/run/Kubernetes`.

For non bare metal device plugin this means they will have to mount the folder
as a volume in their pod spec ([see Installation](##installation)).

Device plugins can expect to find the socket to register themselves on
the host at the following path:
`/var/run/Kubernetes/kubelet.sock`.

### Protocol Overview

When first registering themselves against Kubelet, the device plugin
will send:
  * The name of their unix socket
  * [The API version against which they were built](#versioning).
  * Their `Vendor` ID or name of the device plugin

Kubelet answers with the minimum version it supports and whether or
not there was an error. The errors may include (but not limited to):
  * API version not supported
  * A device plugin was already registered for this vendor
  * A device plugin already registered this device
  * Vendor is not consistent across discovered devices

Kubelet will then interact with the plugin through the following functions:
  * `Discover`: List Devices
  * `Monitor`: Returns a stream that is written to when a
     Device becomes unhealty
  * `Allocate`: Called when creating a container with a list of devices
     can request changes to the Container config
  * `Deallocate`: Called when deleting a container can be used for cleanup

The device plugin is also expected to periodically call the `Heartbeat` function
exposed by Kubelet and issue a `Registration` request when it either can't reach
Kubelet or Kubelet answers with a `KO` response.

![Process](./device-plugin.png)


### Protobuf specification

```go
service PluginRegistration {
	rpc Register(RegisterRequest) returns (RegisterResponse) {}
	rpc Heartbeat(HeartbeatRequest) returns (HeartbeatResponse) {}
}

service DevicePlugin {
	rpc Discover(Empty) returns (stream Device) {}
	rpc Monitor(Empty) returns (stream DeviceHealth) {}

	rpc Allocate(AllocateRequest) returns (AllocateResponse) {}
	rpc Deallocate(DeallocateRequest) returns (Empty) {}
}

message RegisterRequest {
	// Version of the API the Device Plugin was built against
	string version = 1;
	// Name of the unix socket the device plugin is listening on
	string unixsocket = 2;
	// Name of the devices the device plugin wants to register
	// A device plugin can only register one kind of devices
	string vendor = 3;
}

message RegisterResponse {
	// Minimum version the Kubelet API supports.
	string version = 1;
	// Kubelet fills this field if it encounters any errors
	// during the registration process or discover process
	Error error = 2;
}

message HeartbeatRequest {
	string vendor = 1;
}

message HeartbeatResponse {
	// Kubelet answers with a string telling the device
	// plugin to either re-register itself or not
	string response = 1;
	// Kubelet fills this field if it encountered any errors
	Error error = 2;
}

message AllocateRequest {
	repeated Device devices = 1;
}

message AllocateResponse {
	// List of environment variable to set in the container.
	repeated KeyValue envs = 1;
	// Mounts for the container.
	repeated Mount mounts = 2;
}

message DeallocateRequest {
	repeated Device devices = 1;
}

message Error {
	bool error = 1;
	string reason = 2;
}

// E.g:
// struct Device {
//    Kind: "NVIDIA-gpu"
//    Name: "GPU-fef8089b-4820-abfc-e83e-94318197576e"
//    Properties: {
//        "Family": "Pascal",
//        "Memory": "4G",
//        "ECC"   : "True",
//    }
//}
//
message Device {
	string Kind = 1;
	string Name = 2;
	string Health = 3;
	string Vendor = 4;
	map<string, string> properties = 5; // Could be [1, 1.2, 1G]
}

message DeviceHealth {
	string Name = 1;
	string Kind = 2;
	string Vendor = 4;
	string Health = 3;
}
```

## Installation

The installation process should be straightforward to the user, transparent
and similar to other regular Kubernetes actions.
The device plugin should also run in containers so that Kubernetes can
deploy them and restart the plugins when they fail.
However, we should not prevent the user from deploying a bare metal device
plugin.

Deploying the device plugins through DemonSets makes sense as the cluster
admin would be able to specify which machines it wants the device plugins to
run on, the process is similar to any Kubernetes action and does not require
to change any parts of Kubernetes.

Additionally, for integrated solutions such as `kubeadm` we can add support
to auto-deploy community vetted Device Plugins.
Thus not fragmenting once more the Kubernetes ecosystem.

For users installing Kubernetes without using an integrated solution such
as `kubeadm` they would use the examples that we would provide at:
`https://github.com/Kubernetes/Kubernetes/tree/master/examples/device-plugin.yaml`

YAML example:
```yaml
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
spec:
    template:
        metadata:
            labels:
                - name: device-plugin
        spec:
            containers:
                name: device-plugin-ctr
                image: NVIDIA/device-plugin:1.0
                volumeMounts:
                  - mountPath: /device-plugin
                  - name: device-plugin
           volumes:
             - name: device-plugin
               hostPath:
                   path: /var/run/Kubernetes
```

## API Changes
### Device

When discovering the devices, Kubelet will be in charge of advertising those
resources to the API server.

We will advertise each device returned by the Device Plugin in a new structure
called `Device`.
It is defined as follows:

```golang
type Device struct {
	Kind       string
	Vendor     string
	Name       string
	Health     DeviceHealthStatus
	Properties map[string]string
}
```

Because the current API (Capacity) can not be extended to support Device,
we will need to create two new attributes in the NodeStatus structure:
  * `DevCapacity`: Describing the device capacity of the node
  * `DevAvailable`: Describing the available devices

```golang
type NodeStatus struct {
	DevCapacity []Device
	DevAvailable []Device
}
```

We also introduce the `Allocated` field in the pod's status so that user
can know what devices were assigned to the pod. It could also be useful in
the case of monitoring

```golang
type ContainerStatus struct {
	Devices []Device
}
```

# Versioning

Currently there is only one part (CRI) of Kubernetes which is based on
a protobuf model.

The model used by CRI as of now involves the client (kubelet) checking
if the server (runtime) version is compatible and then continuing to
communicate with the server.
Currently for CRI, compatible means matching the exact version.
This means that every time the CRI spec changes the CRI clients needs to
be updated.

CRI also uses gRPC-go, which requires the same package name between client
and server.
If they are not same, then no API calls can succeed because the generated grpc
code registers a service using the `package_name.service_name` convention,
e.g., The StopPodSandbox method is known as `/v1alpha1.RuntimeService/StopPodSandbox.`

To work around this restriction, CRI adopted the strategy to freeze the
package name at `pkg/kubelet/apis/cri/v1alpha1/runtime`.

Considering the restrictions it seems reasonable to follow the same pattern for
the device plugin proposal to prevent API breaking:
  * Follow protobuf guidelines on versionning:
    * Do not change ordering
    * Do not remove fields or change types
    * Add optional fields
  * Freeze the package name to `apis/device-plugin/v1alpha1`
  * Have kubelet and the Device Plugin negotiate versions if we do break the API

Negotiation would take place in the registration:
1. When registering itself with Kubelet, the Device plugin sends the version
   against which it was built.
2. Kubelet returns the minimum version it supports and if the version sent
   is supported.
3. If Kubelet supports the version sent by the Device Plugin, it
   contacts the Device Plugin
4. If the Device Plugin supports the version sent by Kubelet it can and should
   answer the different calls made by Kubelet
