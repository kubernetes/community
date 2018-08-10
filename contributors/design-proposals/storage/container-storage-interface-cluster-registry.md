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
* A CSI volume plugin to register itself with Kubernetes when it is deployed.
* A CSI volume plugin to customize how Kubernetes interacts with it (e.g. skip attach process because the plugin doesn't support ControllerPublish, etc.).
* A user or cluster admin to easily discover which CSI volume plugins are deployed on their Kubernetes cluster.
* Be optional -- a plugin may choose not to use it (and will get the default set of behaviors).


## Background & Motivations

Kubernetes supports the Container Storage Interface (CSI) to enable third party storage developers to deploy volume plugins exposing new storage systems in Kubernetes without having to touch the core Kubernetes code.
Support for CSI was introduced as alpha in Kubernetes v1.9, and moved to beta in v1.10.
See "CSI Volume Plugins in Kubernetes Design Doc" in the "Links" section below for details on how an arbitrary CSI volume plugin is deployed and interacts with Kubernetes.
The beta implementation of CSI has a number of limitations, including:
* CSI drivers must be deployed with the provided CSI external-attacher sidecar container, even if they don’t implement ControllerPublishVolume.
  * Meaning a CSI volume plugin deployed on kubernetes has to deploy a CSI `external-attacher` container even if the volume plugin doesn't require a "volume attach" operation.
  * In this case, the CSI `external-attacher` container basically does a no-op in response to Kubernetes `VolumeAttachment` objects to allow Kubernetes to continue with the mounting process.
* Users and cluster admins have no easy way to discover what CSI volume plugins are deployed on their Kubernetes cluster.

The proposed "CSI cluster registration mechanism" should address these issues.

### Links

* [Container Storage Interface (CSI) Spec](https://github.com/container-storage-interface/spec/blob/master/spec.md)
* [CSI Volume Plugins in Kubernetes Design Doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md)

## Objective

The objective of this document is to document all the requirements for enabling a cluster level registry for CSI compliant volume plugins (a CSI volume driver) in Kubernetes.

## Goals

* Allow a CSI volume plugin (that opts-in to new mechanism) to configure how Kubernetes should interact with it.
* Improve discoverability by users/cluster admins of a CSI volume plugin (that opts-in to new mechanism) deployed on the Kubernetes cluster.

## Non-Goals

* Define how CSI volume plugin should be deployed on Kubernetes.
* Require use of new CSI cluster registry mechanism for all CSI volume plugins deployed on Kubernetes (the mechanism will be optional/opt-in).

## Design Overview

A new custom resource will be automatically be installed on Kubernetes clusters.

Upon deployment a plugin must create a new custom resource object.

## Design Details

### `CSIPlugin` Object

#### Proposed API

```go
// CSIPlugin captures information about a Container Storage Interface (CSI)
// volume plugin deployed on the cluster.
//
// CSIPlugin objects are non-namespaced.
type CSIPlugin struct {
    metav1.TypeMeta   `json:",inline"`

	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    metav1.ObjectMeta `json:"metadata"`

    // Specification describing the CSI volume plugin and any custom
    // configuration for it.
	Spec CSIPluginSpec `json:"spec"`

	// Status of the CSI volume plugin.
	Status CSIPluginStatus `json:"status,omitempty"`
}

// CSIPluginList is a collection of CSIPlugin objects.
type CSIPluginList struct {
    metav1.TypeMeta `json:",inline"`
    
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of CSIPlugins
	Items []CSIPlugin `json:"items"`
}

// CSIPluginSpec is the specification of a CSIPlugin.
type CSIPluginSpec struct {
    // Driver indicates the name of the CSI Plugin that this object refers to.
    // This MUST be the same name returned by the CSI GetPluginName() call for
    // that driver.
    Driver string `json:"driver"`
    
    // Indicates this CSI volume plugin does not require an attach operation
    // (because it does not implement the ControllerPublishVolume() method),
    // and that Kubernetes should not call or wait for any attach operation and
    // just skip to mounting.
	SkipAttach bool `json:"skipAttach"`
}

// CSIPluginStatus is the status of a CSIPlugin.
type CSIPluginStatus struct {
    // Indicates the volume plugin has been successfully deployed on the cluster
    // and is ready to use.
	Ready bool `json:"ready"`
}

```

#### CRD Installation

The `CSIPlugin` object schema will checked in to kubernetes/kubernetes repo under a storage directory.
The `CSIPlugin` `CustomResourceDefinition` (CRD) will be installed by the exiting Kubernetes attach/detach controller that is part of the kube-controller-manager binary.
The controller will periodically verify that the CRD is still installed, and recreate it, if it is not.

#### CR Creation

When a CSI volume plugin is deployed on Kubernetes, it may optionally, register it self with Kubernetes by creating a new `CSIPlugin` object.
The Kubernetes team will modify the CSI [`driver-registrar` container](https://github.com/kubernetes-csi/driver-registrar) to automatically do this.
The plugin must set the `Driver` field to the same name returned by the CSI `GetPluginName()` call for that driver.
The plugin may set any optional configuration fields (like `SkipAttach`) as appropriate.
When the plugin is ready to serve, it must set `Ready` in the status to `true`.

#### Upgrade/Downgrade
TODO

## Alternatives Considered
TODO