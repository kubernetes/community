# CSI Cluster Registry Design Doc

***Status:*** Pending

***Version:*** Alpha

***Author:*** Saad Ali ([@saad-ali](https://github.com/saad-ali), saadali@google.com)

## Terminology

Term | Definition
---|---
Container Storage Interface (CSI) | A specification attempting to establish an industry standard interface that Container Orchestration Systems (COs) can use to expose arbitrary storage systems to their containerized workloads.
CSI Volume Plugin | A new, in-tree volume plugin that acts as an adapter and enables out-of-tree, third-party CSI volume drivers to be used in Kubernetes.
CSI Volume Driver | An out-of-tree CSI compatible implementation of a volume plugin that can be used in Kubernetes through the Kubernetes CSI Volume Plugin.

# Summary

This document proposes a "CSI cluster registration mechanism".

This mechanism will enable:
* A CSI volume driver to register itself with Kubernetes when it is deployed.
* A CSI volume driver to customize how Kubernetes interacts with it (e.g. skip attach process because the driver doesn't support ControllerPublish, etc.).
* A user or cluster admin to easily discover which CSI volume drivers are deployed on their Kubernetes cluster.
* Be optional -- a driver may choose not to use it (and will get the default set of behaviors).


## Background & Motivations

Kubernetes supports the Container Storage Interface (CSI) to enable third party storage developers to deploy volume drivers exposing new storage systems in Kubernetes without having to touch the core Kubernetes code.
Support for CSI was introduced as alpha in Kubernetes v1.9, and moved to beta in v1.10.
See "CSI Volume Plugins in Kubernetes Design Doc" in the "Links" section below for details on how an arbitrary CSI volume driver is deployed and interacts with Kubernetes.
The beta implementation of CSI has a number of limitations, including:
* CSI drivers must be deployed with the provided CSI external-attacher sidecar container, even if they don’t implement ControllerPublishVolume.
  * Meaning a CSI volume driver deployed on kubernetes has to deploy a CSI `external-attacher` container even if the volume driver doesn't require a "volume attach" operation.
  * In this case, the CSI `external-attacher` container basically does a no-op in response to Kubernetes `VolumeAttachment` objects to allow Kubernetes to continue with the mounting process.
* Users and cluster admins have no easy way to discover what CSI volume drivers are deployed on their Kubernetes cluster.

The proposed "CSI cluster registration mechanism" should address these issues.

### Links

* [Container Storage Interface (CSI) Spec](https://github.com/container-storage-interface/spec/blob/master/spec.md)
* [CSI Volume Plugins in Kubernetes Design Doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md)

## Objective

The objective of this document is to document all the requirements for enabling a cluster level registry for a CSI compliant volume drivers in Kubernetes.

## Goals

* Allow a CSI volume driver (that opts-in to new mechanism) to configure how Kubernetes should interact with it.
* Improve discoverability by users/cluster admins of a CSI volume driver (that opts-in to new mechanism) deployed on the Kubernetes cluster.

## Non-Goals

* Define how CSI volume driver should be deployed on Kubernetes.
* Require use of new CSI cluster registry mechanism for all CSI volume drivers deployed on Kubernetes (the mechanism will be optional/opt-in).

## Design Overview

A new custom resource will be automatically be installed on Kubernetes clusters.

Upon deployment a driver must create a new custom resource object.

## Design Details

### `CSIDriver` Object

#### Proposed API

```go
// CSIDriver captures information about a Container Storage Interface (CSI)
// volume driver deployed on the cluster.
//
// CSIDriver objects are non-namespaced.
type CSIDriver struct {
    metav1.TypeMeta   `json:",inline"`

	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metav1.ObjectMeta `json:"metadata"`

    // Specification describing the CSI volume driver and any custom
    // configuration for it.
	Spec CSIDriverSpec `json:"spec"`

	// Status of the CSI volume driver.
	Status CSIDriverStatus `json:"status,omitempty"`
}

// CSIDriverList is a collection of CSIDriver objects.
type CSIDriverList struct {
    metav1.TypeMeta `json:",inline"`
    
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of CSIDrivers
	Items []CSIDriver `json:"items"`
}

// CSIDriverSpec is the specification of a CSIDriver.
type CSIDriverSpec struct {
    // Driver indicates the name of the CSI driver that this object refers to.
    // This MUST be the same name returned by the CSI GetPluginName() call for
    // that driver.
    Driver string `json:"driver"`
    
    // Indicates this CSI volume driver requires an attach operation (because it
    // implements the CSI ControllerPublishVolume() method), and that Kubernetes
    // should call attach and wait for any attach operation to complete before
    // proceeding to mounting.
    // If value is not specified, default is true -- meaning attach will be
    // called.
    // +optional
    AttachRequired *bool `json:"attachRequired"`
    
    // Indicates this CSI volume driver requires additional pod information
    // (like podName, podUID, etc.) during mount operations.
    // If this is set to true, Kubelet will pass pod information as
    // VolumeAttributes in the CSI NodePublishVolume() calls.
    // If value is not specified, default is false -- meaning pod information
    // will not be passed on mount. 
    // +optional
    PodInfoRequiredOnMount *bool `json:"podInfoRequiredOnMount"`
}

// CSIDriverStatus is the status of a CSIDriver.
type CSIDriverStatus struct {
    // Indicates the volume driver has been successfully deployed on the cluster
    // and is ready to use.
	Ready bool `json:"ready"`
}

```

#### CRD Installation

The `CSIDriver` object schema will checked in to kubernetes/kubernetes repo under a storage directory.
The `CSIDriver` `CustomResourceDefinition` (CRD) will be installed by a new Kubernetes controller that is responsible for ensuring required CRDs are installed.
The controller will periodically verify that the CRD is still installed, and recreate it, if it is not.

#### CR Creation

When a CSI volume driver is deployed on Kubernetes, it may optionally, register it self with Kubernetes by creating a new `CSIDriver` object.
The Kubernetes team will modify the CSI [`driver-registrar` container](https://github.com/kubernetes-csi/driver-registrar) to automatically do this.
If kubelet driver registration is enabled, the kubelet will automatically do this as part of plugin registration.
The driver must set the `Driver` field to the same name returned by the CSI `GetPluginName()` call for that driver.
The driver may set any optional configuration fields (like `SkipAttach`) as appropriate.
When the driver is ready to serve, it must set `Ready` in the status to `true`.

#### Upgrade/Downgrade
TODO

## Alternatives Considered
TODO