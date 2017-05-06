# Multi-fields Merge Key in Strategic Merge Patch

## Abstract

Support multi-fields merge key in Strategic Merge Patch.

## Background

Strategic Merge Patch is covered in this [doc](https://github.com/kubernetes/community/blob/master/contributors/devel/strategic-merge-patch.md).
In Strategic Merge Patch, Merge Key is the key to identify the entries in the list of non-primitive types.
It must always be present and unique to perform the merge on the list of non-primitive types,
and will be preserved.

The merge key exists in the struct tag (e.g. in [types.go](https://github.com/kubernetes/kubernetes/blob/5a9759b0b41d5e9bbd90d5a8f3a4e0a6c0b23b47/pkg/api/v1/types.go#L2831))
and the [OpenAPI spec](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json).

## Motivation

The current implementation requires a single field that uniquely identifies each element in a list.
For some element Kinds, the identity is defined using multiple fields.
An [example](https://github.com/kubernetes/kubernetes/issues/39188) is the service.spec.ports,
which is identified by both `protocol` and `port`.

As a result we need to also support a set of keys as a Merge Key.
This will not only fix the APIs that cannot be effectively identified by one single field
but also benefit new APIs when we want to use multi-field merge key for them.

## Proposed Change

### API Change

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

We also need to introduce a new optional directive `$patchMergeKey` in the patch.
This directive contains a list of strings which are the merge keys used in this patch.

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

We will use the default merge key for merging if the $patchMergeKey directive is not present in the patch.
Otherwise, we will use the merge keys provided by the directive.

*Note*: Operations will take effect on all matching items in the list.

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
the server will apply the patch to all the matching items.
So an update or a deletion may apply to multiple items in the list.

*Note*: In last release(1.6) implementation, update operation will apply to the first matching items.
Delete operation will apply to all matching items.

E.g.
foo is the default merge key.
foo and bar are the recommended merge keys.

Live list:
```yaml
list:
- foo: a
  bar: x
  another: 1
- foo: a
  bar: y
  another: 2
- foo: b
  bar: x
```

Patch 1:
```yaml
list:
- $patchMergeKey:
  - foo
  foo: a
  bar: z
  other: val
```

Result after merging patch 1:
```yaml
list:
- foo: a
  bar: z
  other: val
  another: 1
- foo: a
  bar: z
  other: val
  another: 2
- foo: b
  bar: x
```

Patch 2:
```yaml
list:
- $patchMergeKey:
  - foo
  $patch: delete
  foo: a
```

Result after merging patch 2:
```yaml
list:
- foo: b
  bar: x
```

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

### the user change a field which is a merge key

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

### Docs

Document what the developer should consider when adding an API with `mergeKey`.

## Version Skew and Backward Compatibility

An old client sending a old patch without `$patchMergeKey` is backward compatible and
has been described in section [When no `$patchMergeKey` in the patch](#when-no-patchmergekey-in-the-patch).

An old server will drop the field it doesn't recognise.
And the new patch is the superset of the old patch and the difference is the directve.
Thus, the behavior of the old server will not be affected.

## Impacted APIs

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
