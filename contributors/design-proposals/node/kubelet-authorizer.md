# Scoped Kubelet API Access

Author: Jordan Liggitt (jliggitt@redhat.com)

## Overview

Kubelets are primarily responsible for:
* creating and updating status of their Node API object
* running and updating status of Pod API objects bound to their node
* creating/deleting "mirror pod" API objects for statically-defined pods running on their node

To run a pod, a kubelet must have read access to the following objects referenced by the pod spec:
* Secrets
* ConfigMaps
* PersistentVolumeClaims (and any bound PersistentVolume or referenced StorageClass object)

As of 1.6, kubelets have read/write access to all Node and Pod objects, and 
read access to all Secret, ConfigMap, PersistentVolumeClaim, and PersistentVolume objects.
This means that compromising a node gives access to credentials that allow modifying other nodes,
pods belonging to other nodes, and accessing confidential data unrelated to the node's pods.

This document proposes limiting a kubelet's API access using a new node authorizer, admission plugin, and additional API validation:
* Node authorizer
  * Authorizes requests from identifiable nodes using a fixed policy identical to the default RBAC `system:node` cluster role
  * Further restricts secret, configmap, persistentvolumeclaim and persistentvolume access to only allow reading objects referenced by pods bound to the node making the request
* Node admission
  * Limit identifiable nodes to only be able to mutate their own Node API object
  * Limit identifiable nodes to only be able to create mirror pods bound to themselves
  * Limit identifiable nodes to only be able to mutate mirror pods bound to themselves
  * Limit identifiable nodes to not be able to create mirror pods that reference API objects (secrets, configmaps, service accounts, persistent volume claims)
* Additional API validation
  * Reject mirror pods that are not bound to a node
  * Reject pod updates that remove mirror pod annotations

## Alternatives considered

**Can this just be enforced by authorization?**

Authorization does not have access to request bodies (or the existing object, for update requests),
so it could not restrict access based on fields in the incoming or existing object.

**Can this just be enforced by admission?**

Admission is only called for mutating requests, so it could not restrict read access.

**Can an existing authorizer be used?**

Only one authorizer (RBAC) has in-tree support for dynamically programmable policy.

Manifesting RBAC policy rules to give each node access to individual objects within namespaces
would require large numbers of frequently-modified roles and rolebindings, resulting in
significant write-multiplication.

Additionally, not all clusters will use RBAC, but all useful clusters will have nodes.
A node-specific authorizer allows cluster admins to continue to use their authorization mode of choice.

## Node identification

The first step is to identify whether a particular API request is being made by 
a node, and if so, from which node.

The proposed node authorizer and admission plugin will take a `NodeIdentifier` interface:

```go
type NodeIdentifier interface {
  // IdentifyNode determines node information from the given user.Info.
  // nodeName is the name of the Node API object associated with the user.Info,
  // and may be empty if a specific node cannot be determined.
  // isNode is true if the user.Info represents an identity issued to a node.
  IdentifyNode(user.Info) (nodeName string, isNode bool)
}
```

The default `NodeIdentifier` implementation:
* `isNode` - true if the user groups contain the `system:nodes` group and the user name is in the format `system:node:<nodeName>`
* `nodeName` - set if `isNode` is true, by extracting the `<nodeName>` portion of the `system:node:<nodeName>` username

This group and user name format match the identity created for each kubelet as part of [kubelet TLS bootstrapping](https://kubernetes.io/docs/admin/kubelet-tls-bootstrapping/).

## Node authorizer

A new node authorization mode (`Node`) will be made available for use in combination 
with other authorization modes (for example `--authorization-mode=Node,RBAC`).

The node authorizer does the following:
1. If a request is not from a node (`IdentifyNode()` returns isNode=false), reject
2. If a specific node cannot be identified (`IdentifyNode()` returns nodeName=""), reject
3. If a request is for a secret, configmap, persistent volume or persistent volume claim, reject unless the verb is `get`, and the requested object is related to the requesting node:

    * node <-pod
    * node <-pod-> secret
    * node <-pod-> configmap
    * node <-pod-> pvc
    * node <-pod-> pvc <-pv
    * node <-pod-> pvc <-pv-> secret
4. For other resources, allow if allowed by the rules in the default `system:node` cluster role

Subsequent authorizers in the chain can run and choose to allow requests rejected by the node authorizer.

## Node admission

A new node admission plugin (`--admission-control=...,NodeRestriction,...`) is made available that does the following:

1. If a request is not from a node (`IdentifyNode()` returns isNode=false), allow the request
2. If a specific node cannot be identified (`IdentifyNode()` returns nodeName=""), reject the request
3. For requests made by identifiable nodes:
  * Limits `create` of node resources:
    * only allow the node object corresponding to the node making the API request
  * Limits `create` of pod resources:
    * only allow pods with mirror pod annotations
    * only allow pods with nodeName set to the node making the API request
    * do not allow pods that reference any API objects (secrets, serviceaccounts, configmaps, or persistentvolumeclaims)
  * Limits `update` of node and nodes/status resources:
    * only allow updating the node object corresponding to the node making the API request
  * Limits `update` of pods/status resources:
    * only allow reporting status for pods with nodeName set to the node making the API request
  * Limits `delete` of node resources:
    * only allow deleting the node object corresponding to the node making the API request
  * Limits `delete` of pod resources:
    * only allow deleting pods with nodeName set to the node making the API request

## API Changes

Change Pod validation for mirror pods:
  * Reject `create` of pod resources with mirror pod annotations that do not specify a nodeName
  * Reject `update` of pod resources with mirror pod annotations that modify or remove the mirror pod annotation

## RBAC Changes

In 1.6, the `system:node` cluster role is automatically bound to the `system:nodes` group when using RBAC.
Because the node authorizer accomplishes the same purpose, with the benefit of additional restrictions
on secret and configmap access, the automatic binding of the `system:nodes` group to the `system:node` role will be deprecated in 1.7.

In 1.7, the binding will not be created if the `Node` authorization mode is used.

In 1.8, the binding will not be created at all.

The `system:node` cluster role will continue to be created when using RBAC,
for compatibility with deployment methods that bind other users or groups to that role.

## Migration considerations

### Kubelets outside the `system:nodes` group

Kubelets outside the `system:nodes` group would not be authorized by the `Node` authorization mode,
and would need to continue to be authorized via whatever mechanism currently authorizes them.
The node admission plugin would not restrict requests from these kubelets.

### Kubelets with undifferentiated usernames

In some deployments, kubelets have credentials that place them in the `system:nodes` group,
but do not identify the particular node they are associated with.

These kubelets would not be authorized by the `Node` authorization mode,
and would need to continue to be authorized via whatever mechanism currently authorizes them.

The `NodeRestriction` admission plugin would ignore requests from these kubelets,
since the default node identifier implementation would not consider that a node identity.

### Upgrades from previous versions

Upgraded 1.6 clusters using RBAC will continue functioning as-is because the `system:nodes` group binding will already exist.

If a cluster admin wishes to start using the `Node` authorizer and `NodeRestriction` admission plugin
to limit node access to the API, they can do that non-disruptively:
1. Enable the `Node` authorization mode (`--authorization-mode=Node,RBAC`) and the `NodeRestriction` admission plugin
2. Ensure all their kubelets' credentials conform to the group/username requirements
3. Audit their apiserver logs to ensure the `Node` authorizer is not rejecting requests from kubelets (no `NODE DENY` messages logged)
4. Delete the `system:node` cluster role binding

## Future work

Node and pod mutation, and secret and configmap read access are the most critical permissions to restrict.
Future work could further limit a kubelet's API access:
* only write events with the kubelet set as the event source
* only get endpoints objects referenced by pods bound to the kubelet's node (currently only needed for glusterfs volumes)
* only get/list/watch pods bound to the kubelet's node (requires additional list/watch authorization capabilities)
* only get/list/watch it's own node object (requires additional list/watch authorization capabilities)

Features that expand or modify the APIs or objects accessed by the kubelet will need to involve the node authorizer.
Known features in the design or development stages that might modify kubelet API access are:
* [Dynamic kubelet configuration](https://github.com/kubernetes/features/issues/281)
* [Local storage management](/contributors/design-proposals/storage/local-storage-overview.md)
* [Bulk watch of secrets/configmaps](https://github.com/kubernetes/community/pull/443)
