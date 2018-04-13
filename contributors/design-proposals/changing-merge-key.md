# Support Sending Merge Key in the Patch by Clients

## Abstract

Include the merge key in a patch directive sending by the clients.

## Background

Strategic Merge Patch is covered in this [doc](https://github.com/kubernetes/community/blob/master/contributors/devel/strategic-merge-patch.md).
In Strategic Merge Patch, Merge Key is the key to identify the entries in the list of non-primitive types.
It must always be present and unique to perform the merge on the list of non-primitive types,
and will be preserved.

The merge key exists in the struct tag (e.g. in [types.go](https://github.com/kubernetes/kubernetes/blob/5a9759b0b41d5e9bbd90d5a8f3a4e0a6c0b23b47/pkg/api/v1/types.go#L2831))
and the [OpenAPI spec](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json).

`patchMergeKey` tags always exist with `patchStrategy: "merge"`.
`patchStrategy: "merge"` may exist without `patchMergeKey` tags for list of primitives.

We are supporting multi-fields merge key in [#610](https://github.com/kubernetes/community/pull/610).
This proposal depends on #610.

## Motivation

There are some existing APIs that are using merge keys incorrectly.

The current existing APIs requires a single field that uniquely identifies each element in a list.
For some element Kinds, the identity is defined using multiple fields.
An [example](https://github.com/kubernetes/kubernetes/issues/39188) is the service.spec.ports,
which is identified by both `protocol` and `port`.
Therefore, we should change the merge key to multi-fields.

There are some existing APIs that put the merge key tag at the wrong place.
We should also correct them.

## Goal

Propose a solution of changing the merge keys for existing APIs without breaking backward compatibility.

## Proposed Change - Phase 1

There are no API changes in phase 1.
All the changes are in Strategic Merge Patch package.
The changes are to support a new optional directive.

### Strategic Merge Patch pkg

We first need to introduce a new optional directive `$patchMergeKey` in the patch.

`$patchMergeKey` directive has the following properties:
- It should present in each entry in the list when using.
- It contains a list of strings which are the merge keys used in this patch.
- It overrides the patch strategy to `merge` strategy when present.

The reason for the third one is that we may change the patch strtegy from `replace` to `merge`.
It will be helpful for the case in section [API Change - Move Merge Key to the Correct Level](#api-change---move-merge-key-to-the-correct-level).

#### Clients Sends Directive `$patchMergeKey`

The new client will always send a patch with `$patchMergeKey` to
explicitly tell the server what merge keys it knows and want to use.

An example patch will look like:
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  foo: a
  bar: x
  other: val
```

#### Server Respects Directive `$patchMergeKey`

The new server will use the compiled-in merge key for merging
if the $patchMergeKey directive is not present in the patch.
Otherwise, we will use the merge keys provided by the `$patchMergeKey` directive.

#### Relation with `setElementOrder` Directive

The fields used by `$patchMergeKey` must uniquely identify each item in the list.
This is the guarantee for [`setElementOrder`](https://github.com/kubernetes/community/pull/537) directive to work correctly.

We may need to add a validation check to enforce no two items using identical merge keys.


## Proposed Change - Phase 2

In phase 2, we are actually changing merge keys of the existing APIs in a backward compatible way.

This includes 2 kinds of changes:
- supporting multi-fields merge key for existing APIs.
See example in section [Support Multi-Fields Merge Key](#support-multi-fields-merge-key)
- moving merge keys to the correct level of the struct.
See example in section [Move Merge Key to the Correct Level](#move-merge-key-to-the-correct-level)

### API Change - Multi-Fields Merge Key

We introduce a separate set of merge keys, called the recommended merge keys,
for all of the fields using merge key.

It will be in a new struct tag with key `recommendedPatchMergeKey`.
The value will be merge keys seperated by ",", i.e. `recommendedPatchMergeKey:"<key1>,<key2>,<key3>"`.

We will keep the old merge key to keep backward compatibility, and we call them the default merge key.

- For API resources that cannot be effectively merged with a single merge key,
we will introduce additional merge keys in the recommended merge keys.
- For the others, the recommended merge keys will be the same as the default merge key.

Requirements for the recommended merge keys to keep backward compatibility:
- the recommended merge keys must have the default merge key
- the default merge key must be the first one in the recommended merge keys, i.e. `recommendedPatchMergeKey:"<default-merge-key>,<additional-key1>,<additional-key2>"`
- the default merge key must be present in the patch, i.e. the first key in the recommended merge keys must be used.

E.g. [`Ports` in `ServiceSpec`](https://github.com/kubernetes/kubernetes/blob/c51efa9ba0929a643544078d5c182ba75e4b4087/pkg/api/v1/types.go#L2825-L2831).
```go
type ServiceSpec struct {
  // add recommendedPatchMergeKey "port,protocol"
  Ports []ServicePort `patchMergeKey:"port" recommendedPatchMergeKey:"port,protocol" ...`
  ...
}
```

#### When `$patchMergeKey` contains items that are not in the patch

The item in the `$patchMergeKey` list but not in the patch will be considered as:
when merging, this key is required to be not present to match an item in the list.

This will be helpful when we want to distinguish the 2 items in the following list when merging:
```yaml
list:
# foo is the default merge key; foo and bar are the recommended merge key.
- foo: a
- foo: a
  bar: x
```

To match the first item, the patch should look like:
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  foo: a
  # bar not present
  other: val
```

To match the second item, the patch should look like:
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  foo: a
  bar: x # bar presents
  other: val
```

#### When no `$patchMergeKey` in the patch

When there is no `$patchMergeKey` list in the patch,
the server will use the default merge key when merging.

E.g.
foo is the default merge key.
foo and bar are the recommended merge keys.

Patch:
```yaml
list:
- foo: a
  bar: x
  other: val
```

Live list:
```yaml
list:
- foo: a
  bar: y
```

Result after merging:
```yaml
list:
- foo: a
  bar: x
  other: val
```

#### When `$patchMergeKey` only has part of all merge keys

When `$patchMergeKey` only has part of all merge keys,
the server will try to use the existing part of merge keys to uniquely identify an entry in the list.
The server will reject the patch if the existing part of merge keys actually matches more than one entry in the live list.

E.g.
foo is the default merge key.
foo and bar are the recommended merge keys.

Live list:
```yaml
list:
- foo: a
  other: 0
- foo: b
  bar: x
  other: 1
- foo: b
  bar: x
  other: 2
```

Patch 1 (valid):
```yaml
list:
- $patchMergeKey:
  - foo
  foo: a
  other: 3
```

Result after merging patch 1:
```yaml
list:
- foo: a
  other: 3
- foo: b
  bar: x
  other: 1
- foo: b
  bar: x
  other: 2
```

Patch 2 (invalid):
```yaml
list:
- $patchMergeKey:
  - foo
  foo: b
  other: 3
```

Server will reject patch 2, since `foo: b` cannot uniquely identify an entry in the list.

#### When `$patchMergeKey` has all of the merge keys

When `$patchMergeKey` has all of the merge keys,
it should uniquely identify an item in the list.

This case is straightforward, so we don't provide an example here.

#### When the user add|delete a field which is a merge key

When the user adds or deletes a field which is a merge key,
the patch will use the the original values of the merge keys to identify the items.

E.g.
foo is the default merge key.
foo, bar and **baz** are the recommended merge keys.

Live list:
```yaml
list:
- foo: a
  bar: x
- foo: a
  bar: y
```

Patch 1 (add a field)
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  foo: a
  bar: x
  baz: m # add an additional merge key
```

Result after merging patch 1:
```yaml
list:
- foo: a
  bar: x
  baz: m
- foo: a
  bar: y
```

Patch 2 (delete a field)
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  foo: a
  bar: null # delete a merge key
```

Result after merging patch 1:
```yaml
list:
- foo: a
- foo: a
  bar: y
```

### When the user changes a field which is a merge key

Live list:
```yaml
list:
- foo: a
  bar: x
  other: val
```

When the user changes a field which is a merge key and it is in the $patchMergeKey list,
the update operation will be break into a delete operation and a create operation.
The server will apply the deletion and then the create operation.

Patch:
```yaml
list:
- $patchMergeKey:
  - foo
  - bar
  $patch: delete
  foo: a
  bar: x
- $patchMergeKey:
  - foo
  - bar
  foo: a
  bar: y
```

Result after applying patch on live list:
```yaml
list:
- foo: a
  bar: y
```

When the user changes a field which is a merge key and it is NOT in the $patchMergeKey list,
the update operation will be just considered as a regular update as non-mergeKey fields.

Patch:
```yaml
list:
- $patchMergeKey:
  - foo
  foo: a
  bar: y
```

Result after applying patch on live list:
```yaml
list:
- foo: a
  bar: y
  other: val
```

All the impacted APIs are listed in section [Impacted APIs](#impacted-apis)

### API Change - Move Merge Key to the Correct Level

When the merge key is at the wrong level (child level),
it means the field should have the merge key is actually using the default `replace` strategy.

The changes to move merge key to the parent level include:
- delete the old `patchStrategy` and `patchMergeKey` tags.
- add `recommendedPatchMergeKey` for these fields that should have merge keys.

E.g. [`Taints` in `NodeSpec`](https://github.com/kubernetes/kubernetes/blob/c51efa9ba0929a643544078d5c182ba75e4b4087/pkg/api/v1/types.go#L3159).
```go
type NodeSpec struct {
  Taints []Taint `recommendedPatchMergeKey:"key" ...`
  ...
}
```
```go
type Taint struct {
    // Deleted `patchStrategy:"merge" patchMergeKey:"key"`
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`
	...
}
```

**backward compatibility**:
- No one is rely on the `patchStrategy` and `patchMergeKey` tags at the child level.
- `recommendedPatchMergeKey` will use `$patchMergeKey` directive.

#### When no `$patchMergeKey` in the patch

When there is no `$patchMergeKey` in the patch,
the server will use the old patch metadata,
which is using `replace` as patch strategy.

#### When there is `$patchMergeKey` in the patch

The server will use `merge` as patch strategy and
use the fields in the `$patchMergeKey` list as merge key,
even though the old patch strategy may be `replace`.

### Open API

The recommended merge keys should have its corresponding OpenAPI extension as `patchMergeKey` does.

Update [Open API schema](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json)
to reflect the additions of `recommendedPatchMergeKey` struct tags.
E.g. add an additional extension similar to
```
"x-kubernetes-patch-merge-key": "foo"
"x-kubernetes-recommended-patch-merge-key": "foo,bar"
```

### Strategic Merge Patch pkg

We will use the behavior we discussed above when merging the patch.

When calculating the patch, we should always use all of the recommended merge keys.
It means we will include all the merge keys in the $patchMergeKey list,
even though there may be only part of them present in the patch as the case in
section [When `$patchMergeKey` contains items that are not in the patch](#when-patchmergekey-contains-items-that-are-not-in-the-patch).

## Version Skew and Backward Compatibility

An old client sending a old patch without `$patchMergeKey` is backward compatible and
has been described in section When no `$patchMergeKey` in the patch. See [here](#when-no-patchmergekey-in-the-patch) and [here](#when-no-patchmergekey-in-the-patch-1).

An old server will drop the field it doesn't recognise.
And the new patch is the superset of the old patch and the difference is the directve.
Thus, the behavior of the old server will not be affected.

## Impacted APIs

### Support Multi-Fields Merge Key

(1) `ContainerPort`: Change merge key from `containerPort` to `name,containerPort`.

Usage of [ContainerPort](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1637)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1286).
```go
type Container struct {
Ports []ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
...
}
```
```go
type ContainerPort struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// +optional
	HostPort int32 `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"`
	ContainerPort int32 `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"`
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"`
}
```

(2) `ServicePort`: Similar to `ContainerPort`. Change merge key from `port` to `name,port`.

Usage of [ServicePort](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2777)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2867).
```go
type ServiceSpec struct {
	Ports []ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
  ...
}
```
```go
type ServicePort struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol,casttype=Protocol"`
	Port int32 `json:"port" protobuf:"varint,3,opt,name=port"`
	// +optional
	TargetPort intstr.IntOrString `json:"targetPort,omitempty" protobuf:"bytes,4,opt,name=targetPort"`
	// +optional
	NodePort int32 `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort"`
}
```

### Move Merge Key to the Correct Level

(1) `NodeSelectorRequirement`: The merge key and patch strategy is not used correctly, they should be in the list
level (`[]NodeSelectorRequirement`) instead of field level. We should move them out.

Usage of [MatchExpressions](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1971)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1976).

```go
type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions" protobuf:"bytes,1,rep,name=matchExpressions"`
  ...
}
```
```go
type NodeSelectorRequirement struct {
	Key string `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator NodeSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=NodeSelectorOperator"`
	// +optional
	Values []string `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}
```

(2) `Taint`: Similar to `NodeSelectorRequirement`. Tags should be in the list level.

Usage of [Taints](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L3114)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2160).
```go
type NodeSpec struct {
  Taints []Taint `json:"taints,omitempty" protobuf:"bytes,5,opt,name=taints"`
  ...
}
```
```go
type Taint struct {
	Key string `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
	Effect TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect"`
	// +optional
	TimeAdded metav1.Time `json:"timeAdded,omitempty" protobuf:"bytes,4,opt,name=timeAdded"`
}
```

(3) `Toleration`: Similar to `NodeSelectorRequirement`. Tags should be in the list level.

Usage of [Tolerations](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2375)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2200)
```go
type PodSpec struct {
  Tolerations []Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
  ...
}
```
```go
type Toleration struct {
	// +optional
	Key string `json:"key,omitempty" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	// +optional
	Operator TolerationOperator `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"`
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`
	// +optional
	Effect TaintEffect `json:"effect,omitempty" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"`
	// +optional
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty" protobuf:"varint,5,opt,name=tolerationSeconds"`
}
```

### Need auditing APIs

The API owner should audit these APIs to make sure the Merge Key are guaranteed to be the unique identifier.

1) [OwnerReferences](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L200)
```go
type OwnerReference struct {
  OwnerReferences []metav1.OwnerReference `json:"ownerReferences,omitempty" patchStrategy:"merge" patchMergeKey:"uid" protobuf:"bytes,13,rep,name=ownerReferences"`
...
}
```

2) [Env](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1649)
```go
type Container struct {
  Env []EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
  ...
}
```

3) [Volumes](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2261),
[InitContainers](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2275),
[Containers](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2281)
and [ImagePullSecrets](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2357).
```go
type PodSpec struct {
  Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
  InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
  Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
  ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
  ...
}
```

4) [Secrets](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2962)
```go
type ServiceAccount struct {
  Secrets []ObjectReference `json:"secrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=secrets"`
  ...
}
```

5) [Addresses](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L3187)
```go
type NodeStatus struct {
  Addresses []NodeAddress `json:"addresses,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,5,rep,name=addresses"`???
  ...
}
```

6) All `Conditions` using the same pattern.
In `v1/`:
```go
type PodStatus struct {
  Conditions []PodCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
  ...
}

type ReplicationControllerStatus struct {
  Conditions []ReplicationControllerCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
  ...
}

type NodeStatus struct {
  Conditions []NodeCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions"`???
  ...
}

type ComponentStatus struct {
  Conditions []ComponentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
  ...
}
```

In `extensions/`:
```go
type DeploymentStatus struct {
  Conditions []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
  ...
}

type ReplicaSetStatus struct {
  Conditions []ReplicaSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
  ...
}

type DeploymentStatus struct {
  Conditions []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
  ...
}
```

In `apps/`:
```go
type DeploymentStatus struct {
  Conditions []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
  ...
}
```

In `batch/`:
```go
type JobStatus struct {
  Conditions []JobCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
  ...
}
```
