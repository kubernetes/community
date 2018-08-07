---
kep-number: draft-20180806
title: Namespace Initializer
authors:
  - "@easeway"
owning-sig: sig-api-machinery
participating-sigs:
  - sig-architecture
reviewers:
  - "@bgrant0607"
  - "@smartclayton"
  - "@jbeda"
  - "@lavalamp"
  - "@liggitt"
  - "@tallclair"
  - "@davidopp"
approvers:
  - "@bgrant0607"
  - "@lavalamp"
  - "@liggitt"
  - "@tallclair"
  - "@smartclayton"
editor: "@easeway"
creation-date: 2018-08-06
last-updated: 2018-10-23
status: provisional
---

# Namespace Initializer

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
      * [Automatic Policy Population](#automatic-policy-population)
      * [Automated External Service Provisioning](#automated-external-service-provisioning)
    * [Implementation Details](#implementation-details)
      * [Core Kubernetes Changes](#core-kubernetes-changes)
	  * [Initializers Configuration](#initializers-configuration)
    * [Security Considerations](#security-considerations)
	  * [Scalability Considerations](#scalability-considerations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

Namespace Initializer is a proposal for an initialization mechanism specialized for namespaces.
The mechanism is derived from more generic [Initializers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#initializers)
which is going to be deprecated.

## Motivation

Initialization is very useful for _container_ resources (namespace is the only one in current Kubernetes).
Compared to other resources in general,
initialization on namespaces is special because it creates resources inside the namespace,
without updating the namespace itself.
So this can't be achieved by mutating admission webhooks which are normally used on a single resource.
During initialization, a group of resources are created.
And usually, these resources are expected to be created consistently before the _container_ resource is ready for use.

As the more generic Initializers mechanism is being deprecated,
this proposal attempts to keep the mechanism specialized for namespaces and push towards beta and GA.

This mechanism is essential for some applications as currently there's no other way:

- Namespace Population (aka Namespace Template) [KEP PR](https://github.com/kubernetes/community/pull/2177)
- Enterprise cluster administrators want to automatically populate namespaces with policies like ResourceQuota
- Kubernetes cluster shared by multiple tenants automatically populate policies when tenant users create namespaces by themselves

### Goals

- Core Kubernetes API changes for namespace initialization
  - Adding fields to v1 Namespace
  - Ensure backward compatibility
- Security considerations
- Extendability on namespace initializers
- Consideration on scalability
- Recommendations for client/CLI changes (no change required to keep existing clients/CLI working, but change for better)

### Non-Goals

- The details of existing Initializers mechanism
- Changes of initialization mechanism compared to Initializers
- Namespace initializers configured to be applied to a subset of namespaces
- Application specific use of namespace initialization
- V2 Namespace

## Proposal

### User Stories

#### Automatic Policy Population

In a large enterprise, a single Kubernetes cluster is shared by multiple engineering teams.
The cluster administrators allow the engineering teams to do self-service namespace creation (using `kubectl create namespace`) to make the process simpler,
but they still want to enforce the policies in each namespace automatically, including RBAC, ResourceQuota, NetworkPolicy etc.
They define these policies in a centralized places (e.g. git, corp AD, etc.),
and use a controller watching namespace creation and populate the predefined policies during namespace creation before the creator gains full permission inside the namespace.

This is the user story demonstrated by the [Namespace Population](https://github.com/kubernetes/community/pull/2177) proposal.

This scenario for self-service namesapce creation is normally done by only granting namespace creators the `create` and `get` permissions on namespace resource, e.g.

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-creator
rules:
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["get", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: namespace-creators
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: namespace-creator
subjects:
- kind: Group
  apiGroup: rbac.authorization.k8s.io
  name: 'teams@corp.com'
```

And the users with permissions to create namespace doesn't have the control of the cluster.

After that, when a namespace is created, the [Namespace Population](https://github.com/kubernetes/community/pull/2177) mechanism will further create a `RoleBinding`
inside the namespace to grant the creater `admin` permission inside the namespace, e.g.

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: namespace-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: admin
subjects:
- kind: User
  apiGroup: rbac.authorization.k8s.io
  name: 'user1@corp.com'
```

#### Automated External Service Provisioning

The Kubernetes cluster administrator wants to associate namespaces with log analysis system backed by ElasticSearch.
During the namespace creation, they want to automatically provision ElasticSearch with a new index dedicated for logs sent from this namespace.
They deploy a controller to perform the provisioning when watch a new namespace creation,
and they don't want the users to access the namespace before the provisioning is completed.

### Implementation Details

The mechanism is derived from exising [Initializers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#initializers).
For more background and initial discussion, please read comments in [Namespace Population KEP](https://github.com/kubernetes/community/pull/2177).

Briefly, namespace initialization involves a few components:

- A built-in admission controller `NamespaceInitializers`
- A resource of type `NamespaceInitializerConfiguration`
- A few namespace initialization controllers: custom controllers which initialize a namespace by creating resources inside the namespace

They work in the process of:

- The namespace initialization controllers register themselves by creating `NamespaceInitializerConfiguration` resources
- The namespace initialization controllers must run under a service account granted `initialize` RBAC permission (see [Security considerations](#security-considerations) section for details).
- When a namespace is newly created
  - The `NamespaceInitializers` admission controller aggregates all `NamespaceInitializersConfiguration` resources and generate `spec.Initializers` in this namespace resource
  - Once the namespace resource is persisted, the following happens in parallel:
    - Namespace initialization controllers get notified about the newly created namespace,
      and they start creating resources inside namespace and also update `spec.Initializers` to remove themselves once the initialization is done;
    - The `NamespaceInitializers` admission controller will check all resource operations against `initialize` RBAC permission if they are inside a namespace which is being initialized.

#### Core Kubernetes Changes

The generic Initializers uses `metadata.Initializers` to record all pending initializations.
As this is going to be deprecated, for namespaces, `Initializers` is added to `NamespaceSpec` as an _alpha_ field:

```go
// NamespaceSpec describes the attributes on a Namespace.
type NamespaceSpec struct {
  // Finalizers is an opaque list of values that must be empty to permanently remove object from storage.
  // More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
  // +optional
  Finalizers []FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,1,rep,name=finalizers,casttype=FinalizerName"`

  // An initializer is a controller which enforces some system invariant at object creation time.
  // This field is a list of initializers that have not yet acted on this object. If nil or empty,
  // this object has been completely initialized. Otherwise, the object is considered uninitialized
  // and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to
  // observe uninitialized objects.
  //
  // When an object is created, the system will populate this list with the current set of initializers.
  // Only privileged users may set or modify this list. Once it is empty, it may not be modified further
  // by any user.
  // +optional
  Initializers *NamespaceInitializers `json:"initializers,omitempty" protobuf:"bytes,16,opt,name=initializers"`
}

// NamespaceInitializers tracks the progress of initialization.
type NamespaceInitializers struct {
  // Pending is a list of initializers that must execute in order before this object is visible.
  // When the last pending initializer is removed, and no failing result is set, the initializers
  // struct will be set to nil and the object is considered as initialized and visible to all
  // clients.
  // +patchMergeKey=name
  // +patchStrategy=merge
  Pending []NamespaceInitializer `json:"pending" protobuf:"bytes,1,rep,name=pending" patchStrategy:"merge" patchMergeKey:"name"`
  // If result is set with the Failure field, the object will be persisted to storage and then deleted,
  // ensuring that other clients can observe the deletion.
  Result *metav1.Status `json:"result,omitempty" protobuf:"bytes,2,opt,name=result"`
}

// NamespaceInitializer is information about an initializer that has not yet completed.
type NamespaceInitializer struct {
  // name of the process that is responsible for initializing this object.
  Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}
```

Similarly to `InitializerConfiguration`,
a new type `NamespaceInitializerConfiguration` (v1alpha1) is introduced but simplified and specific to namespace:

```go
// NamespaceInitializerConfiguration describes the configuration of initializers.
type NamespaceInitializerConfiguration struct {
  metav1.TypeMeta `json:",inline"`
  // Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
  // +optional
  metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

  // Initializers is a list of resources and their default initializers
  // Order-sensitive.
  // When merging multiple NamespaceInitializerConfigurations, we sort the initializers
  // from different NamespaceInitializerConfigurations by the name of the
  // NamespaceInitializerConfigurations; the order of the initializers from the same
  // NamespaceInitializerConfiguration is preserved.
  // +patchMergeKey=name
  // +patchStrategy=merge
  // +optional
  Initializers []NamespaceInitializer `json:"initializers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=initializers"`
}

// NamespaceInitializer describes the name of an initializer
type NamespaceInitializer struct {
  // Name is the identifier of the initializer. It will be added to the
  // object that needs to be initialized.
  // Name should be fully qualified, e.g., alwayspullimages.kubernetes.io, where
  // "alwayspullimages" is the name of the webhook, and kubernetes.io is the name
  // of the organization.
  // Required
  Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}
```

A built-in mutating admission control `NamespaceInitializers` is also introduced to consume `NamespaceInitializersConfiguration` objects
and injecting `Initializers` list into `Namespace` object being created.

To better demonstrate the current status of the namespace, `NamespaceCondition` is introduced to `NamespaceStatus`,
as the existing `phase` is deprecated, see discussion [here](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#typical-status-properties).

```go
// NamespaceStatus is information about the current status of a Namespace.
type NamespaceStatus struct {
  // Phase is the current lifecycle phase of the namespace.
  // More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
  // +optional
  Phase NamespacePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=NamespacePhase"`
  // Current service state of namespace.
  // +optional
  // +patchMergeKey=type
  // +patchStrategy=merge
  Conditions []NamespaceCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

// NamespaceCondition contains details for the current condition of this namespace.
type NamespaceCondition struct {
  // Type is the type of the condition.
  // Currently only Ready.
  Type NamespaceConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NamespaceConditionType"`
  // Status is the status of the condition.
  // Can be True, False, Unknown.
  Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
  // Last time we probed the condition.
  // +optional
  LastProbeTime metav1.Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
  // Last time the condition transitioned from one status to another.
  // +optional
  LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
  // Unique, one-word, CamelCase reason for the condition's last transition.
  // +optional
  Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
  // Human-readable message indicating details about last transition.
  // +optional
  Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

// NamespaceConditionType is a valid value for NamespaceCondition.Type
type NamespaceConditionType string

// These are valid conditions of namespace.
const (
  // NamespaceInitializing means the namespace is being initialized, and not ready for use.
  NamespaceInitializing NamespaceConditionType = "Initializing"
)
```

With `NamespaceCondition`, the clients and CLI may work better in a situation where initialization takes long as there's a way to tell whether the namespace is ready.
See [Scalability Considerations](#scalability-considerations) below for more details.

Feature gate `NamespaceInitializers` (v1alpha1) is introduced to turn on `Initializers` field in v1 `NamespaceSpec`,
`Conditions` in `NamespaceStatus`,
and also enabling the `NamespaceInitializers` admission control.

##### Alternative Design

If Namespace is not the only container resource (so far it is), an alternative design is preferred to make
initialization common for container resources.
Instead of putting `Initializers` and `Finalizers` into `spec`, define an extra metadata type (extending `ObjectMeta`):

```go

type ContainerMeta struct {
  ObjectMeta

  // Initializers is a list of resources and their default initializers
  // Order-sensitive.
  // +patchMergeKey=name
  // +patchStrategy=merge
  // +optional
  Initializers []Initializer `json:"initializers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=initializers"`

  // Finalizers is an opaque list of values that must be empty to permanently remove object from storage.
  // More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
  // +optional
  Finalizers []FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,1,rep,name=finalizers,casttype=FinalizerName"`
}

```

Note: no change to current `Initializer` schema, and not renaming it to `NamespaceInitializer`.

Then, container resources, e.g. `Namespace` will be defined as:

```go

type Namespace struct {
  metav1.TypeMeta `json:",inline"`
  // Container resource's metadata.
  // More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  // +optional
  metav1.ContainerMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

```

To maintain the compatibility, the existing `Finalizers` in `spec` and existing `Initializers` in `metadata` are not changed.
New logic will be added to specially handle resources with `ContainerMeta`.

#### Initializers Configuration

Similar to the existing `InitializersConfiguration`,
namespace initializers can create `NamespaceInitializerConfiguration` as a convinient way to inject themselves into `Initializers` list
of newly created namespaces.

`NamespaceInitializersConfiguration` is specific for namespaces,
and doesn't support selection of namespaces,
so there's _no_ `Rules` field in the object.

#### Security Considerations

To effectively protect namespaces being initialized,
only the controllers performing namespace initialization are allowed to manipulate resources inside the namespaces.
This can be achieved by introducing the `initialize` verb into RBAC rules and create a cluster role binding with the service account of the controller:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-initializer
rules:
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["initialize"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: namespace-initializer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: namespace-initializer
subjects:
- kind: ServiceAccount
  apiGroup: rbac.authorization.k8s.io
  name: ns-initializer-sa
  namespace: kube-system
```

The `NamespaceInitializer` admission controller will perform the validation for any operations of resources inside a namespace being initialized.

#### Scalability Considerations

Regarding backward compatibility, there's an issue with this mechanism.
For example, the namespace creator may use the following manifest to create a new namespace and populate some resources inside, e.g.

 ```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: ns1
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: ns1
  name: namespace-editors
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: edit
subjects:
- kind: Group
  apiGroup: rbac.authorization.k8s.io
  name: team-a@corp.com
```

If `kubectl apply -f` is used with the above manifest as input, `kubectl` will move forward to create `RoleBinding` in namespace `ns1` which may not finish the initialization process.
And the command will fail as the `RoleBinding` of `namespace-admin` for current creater has not been created yet.
The existing `Initializers` mechanism will make sure the namespace creation response reaches to `kubectl` after the namespace is fully populated, and then the above process will not fail.

However, the time-limited wait doesn't scale in the case of creating large number of namespaces simultanously, like in a certain automation system.
The namespace initializers can be overloaded and react slowly and result in namespace creation timeout.
And also, the hard-coded time-limited is not flexible enough to make the initialization work perfectly with external systems
(see the user story about [Automated External Service Provisioning](#automated-external-service-provisioning)) if the configuration of external systems takes longer.

The best way to solve the problem is changing the behavior of the clients or CLI.
The namespace creation should be treated as a long running asynchronous operation.
Once namespace creation request is completed, it doesn't mean the namespace is ready.
A client or CLI should poll the status of the namespace until it becomes ready.
The recommended way is to poll `Conditions` from `NamespaceStatus` (see the section above).

As Kubernetes works in eventual-consistent way, the client should expect certain failures manipulating a resource and perform retries to reach the consistent state eventually.
It doesn't worth to introduce the complex mechanism (not scalable) of holding the namespace creation request in storage layer.

`kubectl` needs to be changed to follow the recommended way for namespace creation.

## Graduation Criteria

This change adds new alpha features into stable v1 API objects.
New feature gate `NamespaceInitializers` is defined to turn on this feature in alpha stage.
Other features (like [Namespace Population](https://github.com/kubernetes/community/pull/2177))
can be built on-top of this mechanism.

Once other features are proved to be useful with more customer demand,
the `NamespaceInitializers` may be promoted to beta, and finally GA.

## Implementation History

- initial draft
