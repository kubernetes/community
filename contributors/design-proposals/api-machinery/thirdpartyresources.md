# Moving ThirdPartyResources to beta

## Background
There are a number of important issues with the alpha version of
ThirdPartyResources that we wish to address to move TPR to beta. The list is
tracked [here](https://github.com/kubernetes/features/issues/95), and also
includes feedback from existing Kubernetes ThirdPartyResource users. This
proposal covers the steps we believe are necessary to move TPR to beta and to
prevent future challenges in upgrading.


## Goals
1. Ensure ThirdPartyResource APIs operate consistently with first party
Kubernetes APIs.
2. Enable ThirdPartyResources to specify how they will appear in API
discovery to be consistent with other resources and avoid naming conflicts
3. Move TPR into their own API group to allow the extensions group to be
[removed](https://github.com/kubernetes/kubernetes/issues/43214)
4. Support cluster scoped TPR resources
5. Identify other features required for TPR to become beta
6. Minimize the impact to alpha ThirdPartyResources consumers and define a
process for how TPR migrations / breaking changes can be accomplished (for
both the cluster and for end users)

Non-goals
1. Solve automatic conversion of TPR between versions or automatic migration of
existing TPR

### Desired API Semantics
TPRs are intended to look like normal kube-like resources to external clients.
In order to do that effectively, they should respect the normal get, list,
watch, create, patch, update, and delete semantics.

In "normal" Kubernetes APIs, if I have a persisted resource in the same group
with the same name in v1 and v2, they are backed by the same underlying object.
A change made to one is reflected in the other. API clients, garbage collection,
namespace cleanup, version negotiation, and controllers all build on this.

The convertibility of Kubernetes APIs provides a seamless interaction between
versions.  A TPR does not have the ability to convert between versions, which
focuses on the primary role of TPR as an easily extensible and simple mechanism
for adding new APIs. Conversion primarily allows structural, but not backwards
incompatible, changes. By not supporting conversion, all TPR use cases are
preserved, but a large amount of complexity is avoided for consumers of TPR.

Allowing a single, user specified version for a given TPR will provide this
semantic by preventing server-side versioning altogether.  All instances of a
single TPR must have the same version or the Kubernetes API semantic of always
returning a resource encoded to the matching version will not be maintained.
Since conversions (even native Kubernetes conversions) cannot be used to handle
behavioral changes, the same effect can be achieved for TPRs client-side with
overlapping serialization changes.


### Avoiding Naming Problems
There are several identifiers that a Kubernetes API resource has which share
value-spaces within an API group and must not conflict.  They are:
1. Resource-type value space
  1. plural resource-type name - like "configmaps"
  2. singular resource-type name - like "configmap"
  3. short names - like "cm"
2. Kind-type value space - for group "example.com"
  1. Kind name - like "ConfigMap"
  2. ListKind name - like "ConfigMapList"
If these values conflict within their value-spaces then no client will be able
to properly distinguish intent.

The actual name of the TPR-registration (resource that describes the TPR to
create) resource can only protect one of these values from conflict.  Since
Kubernetes API types are accessed via a URL that looks like `/apis/<group>/<version>/namespaces/<namespace-name>/<plural-resource-type>`,
the name of the TPR-registration object will be `<plural-resource-type>.<group>`.

Conflicts with other parts of the value-space can not be detected with static
validation, so there will be a spec/status split with `status.conditions` that
reflect the acceptance status of a TPR-registration.  For instance, you cannot
determine whether two TPRs in the same group have the same short name without
inspecting the current state of existing TPRs.

Parts of the value-space will be "claimed" by making an entry in TPR.status to
include the accepted names which will be served.  This prevents a new TPR from
disabling an existing TPR's name.


## New API
In order to: 
1. eliminate opaquely derived information - deriving camel-cased kind names
from lower-case dash-delimited values as for instance.
1. allow the expression of complex transformations - not all plurals are easily
determined (ox and oxen) and not all are English.  Fields for complete
specification eliminates ambiguity.
1. handle TPR-registration value-space conflicts
1. [stop using the extensions API group](https://github.com/kubernetes/kubernetes/issues/43214)

We can create a type `ThirdPartyResource.apiextension.k8s.io`.
```go
// ThirdPartyResourceSpec describe how a user wants their resource to appear
type ThirdPartyResourceSpec struct {
	// Group is the group this resource belongs in
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
	// Version is the version this resource belongs in
	Version string `json:"version" protobuf:"bytes,2,opt,name=version"`
	// Names holds the information about the resource and kind you have chosen which is
	// surfaced through discovery.
	Names ThirdPartyResourceNames

	// Scope indicates whether this resource is cluster or namespace scoped.  Default is namespaced
	Scope ResourceScope `json:"scope" protobuf:"bytes,8,opt,name=scope,casttype=ResourceScope"`
}

type ThirdPartyResourceNames struct {
	// Plural is the plural name of the resource to serve.  It must match the name of the TPR-registration
	// too: plural.group
	Plural string `json:"plural" protobuf:"bytes,3,opt,name=plural"`
	// Singular is the singular name of the resource.  Defaults to lowercased <kind>
	Singular string `json:"singular,omitempty" protobuf:"bytes,4,opt,name=singular"`
	// ShortNames are short names for the resource.
	ShortNames []string `json:"shortNames,omitempty" protobuf:"bytes,5,opt,name=shortNames"`
	// Kind is the serialized kind of the resource
	Kind string `json:"kind" protobuf:"bytes,6,opt,name=kind"`
	// ListKind is the serialized kind of the list for this resource.  Defaults to <kind>List
	ListKind string `json:"listKind,omitempty" protobuf:"bytes,7,opt,name=listKind"`
}

type ResourceScope string

const (
	ClusterScoped    ResourceScope = "Cluster"
	NamespaceScoped  ResourceScope = "Namespaced"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// ThirdPartyResourceConditionType is a valid value for ThirdPartyResourceCondition.Type
type ThirdPartyResourceConditionType string

const (
	// NameConflict means the resource or kind names chosen for this ThirdPartyResource conflict with others in the group.
	// The first TPR in the group to have the name reflected in status "wins" the name.
	NameConflict ThirdPartyResourceConditionType = "NameConflict"
	// Terminating means that the ThirdPartyResource has been deleted and is cleaning up.
	Terminating ThirdPartyResourceConditionType = "Terminating"
)

// ThirdPartyResourceCondition contains details for the current condition of this ThirdPartyResource.
type ThirdPartyResourceCondition struct {
	// Type is the type of the condition.
	Type ThirdPartyResourceConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ThirdPartyResourceConditionType"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
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

// ThirdPartyResourceStatus indicates the state of the ThirdPartyResource
type ThirdPartyResourceStatus struct {
	// Conditions indicate state for particular aspects of a ThirdPartyResource
	Conditions []ThirdPartyResourceCondition `json:"conditions" protobuf:"bytes,1,opt,name=conditions"`

	// AcceptedNames are the names that are actually being used to serve discovery
	// They may not be the same as names in spec.
	AcceptedNames ThirdPartyResourceNames
}

// +genclient=true

// ThirdPartyResource represents a resource that should be exposed on the API server.  Its name MUST be in the format
// <.spec.plural>.<.spec.group>.
type ThirdPartyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec describes how the user wants the resources to appear
	Spec ThirdPartyResourceSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// Status indicates the actual state of the ThirdPartyResource
	Status ThirdPartyResourceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ThirdPartyResourceList is a list of ThirdPartyResource objects.
type ThirdPartyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items individual ThirdParties
	Items []ThirdPartyResource `json:"items" protobuf:"bytes,2,rep,name=items"`
}
```


## Behavior
### Create
When a new TPR is created, no synchronous action is taken.
A controller will run to confirm that value-space of the reserved names doesn't
collide and sets the "KindNameConflict" condition to `false`.

A custom `http.Handler` will look at request and use the parsed out
GroupVersionResource information to match it to a ThirdPartyResource.  The ThirdPartyResource
will be checked to make sure its valid enough in .Status to serve and will
response appropriated.  If there is no ThirdPartyResource defined, it will delegate
to the next handler in the chain.

### Delete
When a TPR-registration is deleted, it will be handled as a finalizer like a
namespace is done today.  The `Terminating` condition will be updated (like
namespaces) and that will cause mutating requests to be rejected by the REST
handler (see above).  The finalizer will remove all the associated storage.
Once the finalizer is done, it will delete the TPR-registration itself.


## Migration from existing TPR
Because of the changes required to meet the goals, there is not a silent
auto-migration from the existing TPR to the new TPR.  It will be possible, but
it will be manual.  At a high level, you simply:
 1. Stop all clients from writing to TPR (revoke edit rights for all users) and
 stop controllers.
 2. Get all your TPR-data.  
 `$ kubectl get TPR --all-namespaces -o yaml > data.yaml`
 3. Delete the old TPR-data.  Be sure you orphan!  
 `$ kubectl delete TPR --all --all-namespaces --cascade=false`
 4. Delete the old TPR-registration.  
 `$ kubectl delete TPR/name`
 5. Create a new TPR-registration with the same GroupVersionKind as before.  
 `$ kubectl create -f new_tpr.name`
 6. Recreate your new TPR-data.  
 `$ kubectl create -f data.yaml`
 7. Restart controllers.

There are a couple things that you'll need to consider:
 1. Garbage collection.  You may have created links that weren't respected by
 the GC collector in 1.6.  Since you orphaned your dependents, you'll probably
 want to re-adopt them like the Kubernetes controllers do with their resources.
 2. Controllers will observe deletes.  Part of this migration actually deletes
 the resource.  Your controller will see the delete.  You ought to shut down
 your TPR controller while you migrate your data.  If you do this, your
 controller will never see a delete.

