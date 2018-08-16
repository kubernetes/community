# Hardware Accelerators

Author: Vishnu Kannan (vishh@)

## Introduction

Hardware Accelerators are becoming a widely used commodity across various industries.
Accelerators can bring down computing latency and/or costs significantly.
Many of these accelerators present unique technical constraints.
Kubernetes as an established Application Management Platform can play a crucial role in getting workloads that require hardware acceleration to be portable and function at scale.
Most of the large public clouds already provide various types of Hardware Accelerators.
There are quite a few Kubernetes users who are already managing Accelerators at scale on their own datacenters.

The remainder of this document provides a technical overview of how hardware accelerators will be integrated into Kubernetes.

## Goals

* Make Kubernetes Machine Learning and Deep Learning friendly.
* Extensible support for various types of hardware accelerators Nvidia GPUs, AMD GPUs and Google Tensor Processing Units
* Portability across Kubernetes clusters
* Limit feature creep to Kubernetes nucleus
* Provide baseline application performance guarantees

## Non Goals
* Support for Cloud Gaming, Simulations, Remote Desktops and other workloads
  * Support for these workloads will be tackled once support for Machine Learning matures

## System Design

The following sections highlight some of the critical design points for supporting Hardware Accelerators

### API

A plethora of Hardware accelerators exist in the world.
Some of them are general purpose, but many of them are purpose built for specific use cases.
Exposing all hardware accelerators as well known (first class) Compute Resource Types will bloat the API and compromise portability.
For this reason, Hardware Accelerators are expected to be handled as “Extended Compute Resources”.

Kubernetes nucleus will recommend and document a general purpose resource name for each family of accelerators - examples include `nvidia.com/gpu`, `amd.com/gpu`, `google.com/tpu`, etc., with a standard domain name that unambiguously identifies the vendor of a hardware followed by the hardware type `<hardware-vendor-domain>/<hardware-type>`.
It is expected that the hardware vendors will work with the kubernetes community to keep their resource names consistent across kubernetes clusters.

Nodes are expected to be homogenous and any attributes specific to hardware accelerators are expected to be exposed as node labels in Kubernetes to begin with.
Users can expose “extended resources” with other names and consume them in their own clusters.
The admission logic will be extended to allow any resource with a non empty non-default (not `kubernetes.io`) domain name.
The scheduler will be extended to treat such extended resources as an integer resource to begin with.

GPU workloads may wish to express soft and hard preferences for specific accelerator sub-types within a family.
For example, a CUDA app may benefit from running on the latest and greatest Nvidia GPU, but has a minimum GPU version requirement.
In addition to this, [feedback from Nvidia](https://docs.google.com/document/d/1lSwVh2ZfJ2FeLXIeyyiNqN_hKPYpahJiwN5X5cszjOk/edit) has indicated that newer version of GPUs are placed alongside older generations on the same machine (heterogenous nodes). 
To support a combination of these two use cases, new scheduling features have been proposed and discussed in the community [here](https://docs.google.com/document/d/1666PPUs4Lz56TqKygcy6mXkNazde-vwA7q4e5H92sUc/edit). 
Kubernetes will support heterogenous nodes in the future, but the initial plan is to support homogenous nodes with standard node labels prior to tackling heterogeneous nodes.
Node labels satisfy the former "soft" and "hard" preferences use case on homogeneous nodes.
Homogeneous nodes are recommended over heteregenous nodes since the latter may cause resource fragmentation at scale.

### SW Infrastructure for Accelerators

Hardware Accelerators often need vendor provided kernel and user space software.
These software at times introduce tight coupling between the host and applications.
Nvidia GPUs for example are consumed via higher level APIs like CUDA, CUVID, etc.
These APIs are available via user space libraries.
The libraries themselves are tied to the host image (kernel and Nvidia kernel driver versions primarily).
These APIs break the abstraction of containers where the general assumption is that applications inside a container bring all their libraries as part of the container image.

#### Extensibility

Instead of building a special solution for Nvidia GPUs in Kubernetes, a standard extension pipeline called "Hardware Device Plugin" [has been proposed](https://docs.google.com/a/google.com/document/d/1LHeTPx_fWA1PdZkHuALPzYxR0AYXUiiXdo3S0g2VSlo/edit?usp=drive_web) to support arbitrary hardware (and virtual) devices without requiring device specific changes to Kubernetes nucleus.
SW for hardware accelerators are expected to be shipped via standard containers. These containers are expected to be deployed on every node with accelerators. These containers are expected to install necessary SW for initializing hardware accelerators, register themselves with the Kubelet via standard device plugin APIs and exposing accelerators as consumable compute resources via Kubernetes APIs.
Kubelet will handle allocation of hardware accelerators to pods and containers.
Kubelet will communicate with the plugins to ensure that the necessary environment (SW, devices, env variables, etc.) to access hardware accelerators assigned to a pod/container are made accessible within the pod/container sandbox.

Kubernetes will not provide any primitives to manage lifecycle of SW for hardware accelerators that are handled outside of Hardware Device Plugins. 

### Monitoring

Hardware Accelerators are expensive and typically have unique hardware architectures.
Programming against these accelerators, improving performance and utilization is non-trivial.
Certain generic metrics like `utilization` and `usage_time`, and vendor specific metrics are expected to be exposed via cAdvisor and made available to monitoring solutions.
These metrics will not be available as part of the core Metrics APIs since Kubernetes nucleus isn't providing any functionality based on these metrics.

### Predictable performance

Accelerators are preferred over CPUs mainly for performance reasons.
Accelerators typically have extreme requirements at the hardware level in terms of power, hardware interconnect bandwidth, latency, etc.
These high performance devices require careful placement of user workloads on specific CPUs, Memory banks and Accelerator devices to reduce latency and guarantee application performance.
Kubernetes will support support performance isolation for these hardware accelerators, by allowing hardware device plugins to expose a hardware topology graph where each edge represents latency to access one or more CPUs.
Kubelet will combine graphs from multiple plugins along with the node’s NUMA topology to handle hardware device assignment.
Performance guarantees are expected to be tackled once basic support for Hardware Accelerators mature.

## Implementation Plan

The following implementation plan is a proposal. Actual delivery dates may change.

## Alpha

### Requirements
* Opt-in with no support
* No patch fixes or backwards compatibility guarantees.

### Timeline

Support for Nvidia GPUs has been in alpha since `v1.6`.

## Beta

### Requirements 

* Backwards compatibile Resource APIs
* Portability across Kubernetes clusters
* Adequate documentation on API including vendor specific attributes
* End to end tests that exercise the APIs

### Dependencies

* Hardware Device Plugin feature set and functionalities mature with backwards compatible APIs

### Timelines

* Current target is `v1.10`

## General Availability

### Requirements

* Minimal performance guarantees

### Dependencies

* Hardware topology aware scheduling in Kubelet along with support for topology in Hardware Device Plugin APIs

### Timelines

* Current target is `v1.12`
