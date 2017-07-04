# Service Policy Proposal

## Abstract

ServicePolicy allows cluster administrators to control the creation and validation of a service.
The intent of ServicePolicy is to manage the services.

## Motivation

As we know, field level control over allowed fields in services does not exist yet. we can not limit
the ip range for different users. and if we don't want people use the node port directly, we can not
implement it although the podsecuitypolicy can limit the "hostPort".

## Goals

1.  Associate [service accounts](../design-proposals/service_accounts.md), groups, and users with
a set of constraints that dictate how a security context is established for a service.
1.  Provide the ability for users and infrastructure components to use services with elevated privileges
on behalf of another user or within a namespace where privileges are more restrictive.
1.  Secure the ability to reference elevated permissions or to change the constraints under which
a user runs.

## Design

### Model

ServicePolicy objects exist in the root scope, outside of a namespace. The ServicePolicy will reference
users and groups that are allowed to operate under the constraints.

```go
// ServicePolicy governs the ability to make requests that affect the SecurityContext
// that will be applied to a pod and container.
type ServicePolicy struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata,omitempty"`

	// Spec defines the policy enforced.
	Spec ServicePolicySpec `json:"spec,omitempty"`
}

// ServicePolicySpec defines the policy enforced.
type ServicePolicySpec struct {
	// Type is the type that can be visit the service
	Type []AllowedType `json:"type,omitempty"`
	// NodePorts determines which node port ranges are allowed to be exposed.
	NodePorts []NodePortRange `json:"nodePorts,omitempty"`
	// ClusterIP determines which cluster ip ranges are allowed to be exposed.
	ClusterIP []ClusterIPRange `json:"clusterIP,omitempty"`
}

// NodePortRange defines a range of node ports that will be enabled by a policy
// for services to use.  It requires both the start and end to be defined.
type NodePortRange struct {
	// Start is the beginning of the port range which will be allowed.
	Start int `json:"start"`
	// End is the end of the port range which will be allowed.
	End int `json:"end"`
}

// ClusterIPRange defines a range of cluster ip that will be enabled by a policy
// for services to use.  It requires both the ip and netmask to be defined.
type ClusterIPRange struct {
	// Start is the beginning of the cluster ip range which will be allowed.
	Start int `json:"start"`
	// End is the end of the cluster ip range which will be allowed.
	End int `json:"end"`
}
```
