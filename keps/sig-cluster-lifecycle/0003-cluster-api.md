# Kubernetes Cluster Management API

## Metadata
```
---
kep-number: 3
title: Kubernetes Cluster Management API
status: accepted
authors:
    - "@roberthbailey"
    - "@pipejakob"
owning-sig: sig-cluster-lifecycle
reviewers:
  - "@thockin"
approvers:
  - "@roberthbailey"
editor: 
  - "@roberthbailey"
creation-date: 2018-01-19
last-updated: 2018-01-22

```

## Table of Contents

* [Kubernetes Cluster Management API](#kubernetes-cluster-management-api)
  * [Metadata](#metadata)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Motivation](#motivation)
    * [Goals](#goals)
    * [Non\-goals](#non-goals)
    * [Challenges and Open Questions](#challenges-and-open-questions)
  * [Proposal](#proposal)
    * [Driving Use Cases](#driving-use-cases)
    * [Cluster\-level API](#cluster-level-api)
    * [Machine API](#machine-api)
      * [Capabilities](#capabilities)
      * [Overview](#overview)
      * [In\-place vs\. Replace](#in-place-vs-replace)
      * [Omitted Capabilities](#omitted-capabilities)
      * [Conditions](#conditions)
      * [Types](#types)
  * [Graduation Criteria](#graduation-criteria)
  * [Implementation History](#implementation-history)
  * [Drawbacks](#drawbacks)
  * [Alternatives](#alternatives)

## Summary

We are building a set of Kubernetes cluster management APIs to enable common cluster lifecycle operations (install, upgrade, repair, delete) across disparate environments. 
We represent nodes and other infrastructure in Kubernetes-style APIs to enable higher level controllers to update the desired state of the cluster (e.g. the autoscaling controller requesting additional machines) and reconcile the world with that state (e.g. communicating with cloud providers to create or delete virtual machines). 
With the full state of the cluster represented as API objects, Kubernetes installers can use them as a common configuration language, and more sophisticated tooling can be built in an environment-agnostic way.

## Motivation

Kubernetes has a common set of APIs (see the [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md)) to orchestrate containers regardless of deployment mechanism or cloud provider. 
Kubernetes also has APIs for handling some infrastructure, like load-balancers, ingress rules, or persistent volumes, but not for creating new machines. 
As a result, the deployment mechanisms that manage Kubernetes clusters each have unique APIs and implementations for how to handle lifecycle events like cluster creation or deletion, master upgrades, and node upgrades.
Additionally, the cluster-autoscaler is responsible not only for determining when the cluster should be scaled, but also responsible for adding capacity to the cluster by interacting directly with the cloud provider to perform the scaling. 
When another component needs to create or destroy virtual machines, like the node auto provisioner, it would similarly need to reimplement the logic for interacting with the supported cloud providers (or reuse the same code to prevent duplication).

### Goals

* The cluster management APIs should be declarative, Kubernetes-style APIs that follow our existing [API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md).
* To the extent possible, we should separate state that is environment-specific from environment-agnostic.
   * However, we still want the design to be able to utilize environment-specific functionality, or else it likely won’t gain traction in favor of other tooling that is more powerful.

### Non-goals

* To add these cluster management APIs to Kubernetes core.
* To support infrastructure that is irrelevant to Kubernetes clusters.
   * We are not aiming to create terraform-like capabilities of creating any arbitrary cloud resources, nor are we interested in supporting infrastructure used solely by applications deployed on Kubernetes. The goal is to support the infrastructure necessary for the cluster itself.
* To convince every Kubernetes lifecycle product ([kops](https://github.com/kubernetes/kops), [kubespray](https://github.com/kubernetes-incubator/kubespray), [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/), [Azure Container Service](https://azure.microsoft.com/en-us/services/container-service/), [Elastic Container Service for Kubernetes](https://aws.amazon.com/eks/), etc.) to support these APIs.
   * There is value in having consistency between installers and broad support for the cluster management APIs and in having common infrastructure reconcilers used post-installation, but 100% adoption isn't an immediate goal.
* To model state that is purely internal to a deployer.
   * Many Kubernetes deployment tools have intermediate representations of resources and other internal state to keep track of. They should continue to use their existing methods to track internal state, rather than attempting to model it in these APIs.

### Challenges and Open Questions

* Should a single Kubernetes cluster only house definitions for itself?
   * If so, that removes the ability to have a single cluster control the reconciliation of infrastructure for other clusters.
   * However, with the concurrent [Cluster Registry](https://docs.google.com/a/google.com/document/d/1Oi9EO3Jwtp69obakl-9YpLkP764GZzsz95XJlX1a960/edit) project, a good separation of responsibilities would be that the Cluster Registry API is responsible for indexing multiple clusters, each of which would only have to know about itself. In order to achieve cross-cluster reconciliation, a controller would need to integrate with a Cluster Registry for discovery.
* Should a cluster’s control plane definition should be housed within that same cluster. 
   * If the control plane becomes unhealthy, then it won’t be able to rectify itself without external intervention. If the control plane configuration lives elsewhere, and the controllers reconciling its state are able to act in the face of control plane failure, then this API could be used to fix a misconfigured control plane that is unresponsive.
* Should our representation of Nodes allow declarative versioning of non-Kubernetes packages, like the container runtime, the Linux kernel, etc.?
   * It potentially enables the use case of smaller, in-place upgrades to nodes without changing the node image.
   * We may be able to leverage cloud-init to some extent, but since it isn’t supported across all cloud/distributions, and doesn’t support upgrades (or any actions beyond initialization), this may devolve into rolling our own solution.
* Should the Cluster API bother with control plane configuration, or expect each component to use component config?
   * One option is to allow arbitrary API objects to be defined during cluster initialization, which will be a combination of Cluster objects, NodeSet objects, and ConfigMaps for relevant component config. This makes the Cluster API less comprehensive, but avoids redundancy and more accurately reflects the desired state of the cluster.
   * Another option is to have key component config embedded in the Cluster API, which will then be created as the appropriate ConfigMaps during creation. This would be used as a convenience during cluster creation, and then the separate ConfigMaps become the authoritative configuration, potentially with a control loop to propagate changes from the embedded component config in the Cluster API to the appropriate (authoritative) ConfigMaps on an ongoing basis.
* Do we want to allow for arbitrary node boot scripts?
   * Some existing tools like kubicorn support this, but the user demand isn’t clear yet.
   * Also see https://github.com/kubernetes/kops/issues/387
      * Kops now has hooks
* Are there any environments in which it only makes sense to refer to a group of homogeneous nodes, instead of individual ones?
   * The current proposal is to start with individual objects to represent each declarative node (called a “Machine”), which allows us to build support for Sets and Deployments on top of them in the future. However, does this simplification break for any environment we want to support?


## Proposal

### Driving Use Cases

_TODO_: Separate out the use cases that are focused on the control plane vs. those focused on nodes.


These use cases are in scope for our v1alpha1 API design and initial prototype implementation:

* Initial cluster creation using these API objects in yaml files (implemented via client-side bootstrapping of resources)
   * Rather than each Kubernetes installer having its own custom APIs and cluster definitions, they could be fed the definition of the cluster via serialized API objects. This would lower the friction of moving between different lifecycle products.
* Declarative Kubernetes upgrades for the control plane and kubelets
* Declarative upgrades for node OS images
* Maintaining consistency of control plane and machine configuration across different clusters / clouds
   * By representing important cluster configuration via declarative objects, operations like “diffing” the configuration of two clusters becomes very straightforward. Also, reconcilers can be written to ensure that important cluster configuration is kept in sync between different clusters by simply copying objects.
* Cloud adoption / lift and shift / liberation

These use cases are in scope for the project, but post-v1alpha1:

* Server-side node draining
* Autoscaling
   * Currently, the OSS cluster autoscaler has the responsibility of determining the right size of the cluster and calling the cloud provider to perform the scaling (supporting every cloud provider directly). Modeling groups of nodes in a declarative way would allow autoscalers to only need to worry about the correct cluster size and error handling when that can’t be achieved (e.g. in the case of stockouts), and then separate cloud controllers can be responsible for creating and deleting nodes to reconcile that state and report any errors encountered.
* Integration with the Cluster Registry API
   * Automatically add a new cluster to a registry, support tooling that works across multiple clusters using a registry, delete a cluster from a registry.
* Supporting other common tooling, like monitoring

These use cases are out of scope entirely:

* Creating arbitrary cloud resources

### Cluster-level API

This level of the Cluster Management API describes the global configuration of a cluster. It should be capable of representing the versioning and configuration of the entire control plane, irrespective of the representation of nodes.

Given the recent efforts of SIG Cluster Lifecycle to make kubeadm the de facto standard toolkit for cloud- and vendor-agnostic cluster initialization, and because kubeadm has [an existing API](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubeadm/app/apis/kubeadm/v1alpha1/types.go) to define the global configuration for a cluster, it makes sense to coalesce the global portion of the Cluster API with the API used by “kubeadm init” to configure a cluster master.

A current goal is to make these APIs as cloud-agnostic as possible, so that the entire definition of a Cluster could remain reasonably in-sync across different deployments potentially in different cloud providers, which would help enable hybrid usecases where it’s desirable to have key configuration stay in sync across different clusters potentially in different clouds/environments. However, this goal is balanced against making the APIs coherent and usable, which strict separation may harm.

The full types for this API can be seen and were initially discussed in [kube-deploy#306](https://github.com/kubernetes/kube-deploy/pull/306).

### Machine API

#### Capabilities

The set of node capabilities that this proposal is targeting for v1alpha1 are:
1. A new Node can be created in a declarative way, including Kubernetes version and container runtime version. It should also be able to specify provider-specific information such as OS image, instance type, disk configuration, etc., though this will not be portable.
1. A specific Node can be deleted, freeing external resources associated with it.
1. A specific Node can have its kubelet version upgraded or downgraded in a declarative way\*.
1. A specific Node can have its container runtime changed, or its version upgraded or downgraded, in a declarative way\*.
1. A specific Node can have its OS image upgraded or downgraded in a declarative way\*.

\* It is an implementation detail of the provider if these operations are performed in-place or via Node replacement.

#### Overview

This proposal introduces a new API type: **Machine**.

A "Machine" is the declarative spec for a Node, as represented in Kubernetes core. If a new Machine object is created, a provider-specific controller will handle provisioning and installing a new host to register as a new Node matching the Machine spec. If the Machine's spec is updated, a provider-specific controller is responsible for updating the Node in-place or replacing the host with a new one matching the updated spec. If a Machine object is deleted, the corresponding Node should have its external resources released by the provider-specific controller, and should be deleted as well.

Fields like the kubelet version, the container runtime to use, and its version, are modeled as fields on the Machine's spec. Any other information that is provider-specific, though, is part of an opaque ProviderConfig string that is not portable between different providers.

The ProviderConfig is recommended to be a serialized API object in a format owned by that provider, akin to the [Component Config](https://docs.google.com/document/d/1arP4T9Qkp2SovlJZ_y790sBeiWXDO6SG10pZ_UUU-Lc/edit) pattern. This will allow the configuration to be strongly typed, versioned, and have as much nested depth as appropriate. These provider-specific API definitions are meant to live outside of the Machines API, which will allow them to evolve independently of it. Attributes like instance type, which network to use, and the OS image all belong in the ProviderConfig.

#### In-place vs. Replace

One simplification that might be controversial in this proposal is the lack of API control over "in-place" versus "replace" reconciliation strategies. For instance, if a Machine's spec is updated with a different version of kubelet than is actually running, it is up to the provider-specific controller whether the request would best be fulfilled by performing an in-place upgrade on the Node, or by deleting the Node and creating a new one in its place (or reporting an error if this particular update is not supported). One can force a Node replacement by deleting and recreating the Machine object rather than updating it, but no similar mechanism exists to force an in-place change.

Another approach considered was that modifying an existing Machine should only ever attempt an in-place modification to the Node, and Node replacement should only occur by deleting and creating a new Machine. In that case, a provider would set an error field in the status if it wasn't able to fulfill the requested in-place change (such as changing the OS image or instance type in a cloud provider).

The reason this approach wasn't used was because most cluster upgrade tools built on top of the Machines API would follow the same pattern:

```
for machine in machines:
    attempt to upgrade machine in-place
    if error:
        create new machine
        delete old machine
```

Since updating a Node in-place is likely going to be faster than completely replacing it, most tools would opt to use this pattern to attempt an in-place modification first, before falling back to a full replacement.

It seems like a much more powerful concept to allow every tool to instead say:

```
for machine in machines:
    update machine
```

and allow the provider to decide if it is capable of performing an in-place update, or if a full Node replacement is necessary.

#### Omitted Capabilities

**A scalable representation of a group of nodes**

Given the existing targeted capabilities, this functionality could easily be built client-side via label selectors to find groups of Nodes and using (1) and (2) to add or delete instances to simulate this scaling.

It is natural to extend this API in the future to introduce the concepts of MachineSets and MachineDeployments that mirror ReplicaSets and Deployments, but an initial goal is to first solidify the definition and behavior of a single Machine, similar to how Kubernetes first solidifed Pods.

A nice property of this proposal is that if provider controllers are written solely against Machines, the concept of MachineSets can be implemented in a provider-agnostic way with a generic controller that uses the MachineSet template to create and delete Machine instances. All Machine-based provider controllers will continue to work, and will get full MachineSet functionality for free without modification. Similarly, a MachineDeployment controller could then be introduced to generically operate on MachineSets without having to know about Machines or providers. Provider-specific controllers that are actually responsible for creating and deleting hosts would only ever have to worry about individual Machine objects, unless they explicitly opt into watching higher-level APIs like MachineSets in order to take advantage of provider-specific features like AutoScalingGroups or Managed Instance Groups.

However, this leaves the barrier to entry very low for adding new providers: simply implement creation and deletion of individual Nodes, and get Sets and Deployments for free.

**A provider-agnostic mechanism to request new nodes**

In this proposal, only certain attributes of Machines are provider-agnostic and can be operated on in a generic way. In other iterations of similar proposals, much care had been taken to allow the creation of truly provider-agnostic Machines that could be mapped to provider-specific attributes in order to better support usecases around automated Machine scaling. This introduced a lot of upfront complexity in the API proposals.

This proposal starts much more minimalistic, but doesn't preclude the option of extending the API to support these advanced concepts in the future.

**Dynamic API endpoint**

This proposal lacks the ability to declaratively update the kube-apiserver endpoint for the kubelet to register with. This feature could be added later, but doesn't seem to have demand now. Rather than modeling the kube-apiserver endpoint in the Machine object, it is expected that the cluster installation tool resolves the correct endpoint to use, starts a provider-specific Machines controller configured with this endpoint, and that the controller injects the endpoint into any hosts it provisions.

#### Conditions

[bgrant0607](https://github.com/bgrant0607) and [erictune](https://github.com/erictune) have indicated that the API pattern of having "Conditions" lists in object statuses is soon to be deprecated. These have generally been used as a timeline of state transitions for the object's reconciliation, and difficult to consume for clients that just want a meaningful representation of the object's current state. There are no existing examples of the new pattern to follow instead, just the guidance that we should use top-level fields in the status to represent meaningful information. We can revisit the specifics when new patterns start to emerge in core.

#### Types

The full Machine API types can be found and discussed in [kube-deploy#298](https://github.com/kubernetes/kube-deploy/pull/298).

## Graduation Criteria

__TODO__

## Implementation History

* **December 2017 (KubeCon Austin)**: Prototype implementation on Google Compute Engine using Custom Resource Definitions

## Drawbacks

__TODO__

## Alternatives

__TODO__
