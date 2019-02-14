# Multi-fields Merge Key in Strategic Merge Patch

## Abstract

Support multi-fields merge key in Strategic Merge Patch.

## Background

Strategic Merge Patch is covered in this [doc](/contributors/devel/sig-api-machinery/strategic-merge-patch.md).
In Strategic Merge Patch, we use Merge Key to identify the entries in the list of non-primitive types.
It must always be present and unique to perform the merge on the list of non-primitive types,
and will be preserved.

The merge key exists in the struct tag (e.g. in [types.go](https://github.com/kubernetes/kubernetes/blob/5a9759b0b41d5e9bbd90d5a8f3a4e0a6c0b23b47/pkg/api/v1/types.go#L2831))
and the [OpenAPI spec](https://git.k8s.io/kubernetes/api/openapi-spec/swagger.json).

## Motivation

The current implementation only support a single field as merge key.
For some element Kinds, the identity is actually defined using multiple fields.
[Service port](https://github.com/kubernetes/kubernetes/issues/39188) is an evidence indicating that
we need to support multi-fields Merge Key.

## Scope

This proposal only covers how we introduce ability to support multi-fields merge key for strategic merge patch.
It will cover how we support new APIs with multi-fields merge key.

This proposal does NOT cover how we change the merge keys from one single field to multi-fields
for existing APIs without breaking backward compatibility,
e.g. we are not addressing the service port issue mentioned above.
That part will be addressed by [#476](https://github.com/kubernetes/community/pull/476).

## Proposed Change

### API Change

If a merge key has multiple fields, it will be a string of merge key fields separated by ",", i.e. `patchMergeKey:"<key1>,<key2>,<key3>"`.

If a merge key only has one field, it will be the same as before, i.e. `patchMergeKey:"<key1>"`.

There are no patch format changes.
Patches for fields that have multiple fields in the merge key must include all of the fields of the merge key in the patch.

If a new API uses multi-fields merge key, all the fields of the merge key are required to present.
Otherwise, the server will reject the patch.

E.g.
foo and bar are the merge keys.

Live list:
```yaml
list:
- foo: a
  bar: x
  other: 1
- foo: a
  bar: y
  other: 2
- foo: b
  bar: x
  other: 3
```

Patch 1:
```yaml
list:
- foo: a # field 1 of merge key
  bar: x # field 2 of merge key
  other: 4
  another: val
```

Result after merging patch 1:
```yaml
list:
- foo: a
  bar: x
  other: 4
  another: val
- foo: a
  bar: y
  other: 2
- foo: b
  bar: x
  other: 3
```

Patch 2:
```yaml
list:
- $patch: delete
  foo: a # field 1 of merge key
  bar: x # field 2 of merge key
```

Result after merging patch 2:
```yaml
list:
- foo: a
  bar: y
  other: 2
- foo: b
  bar: x
  other: 3
```

### Strategic Merge Patch pkg

We will add logic to support
- returning a list of fields instead of one single field when looking up merge key.
- merging list respecting a list of fields as merge key.

### Open API

Open API will not be affected,
since multi-fields merge key is still in one single string as an extension in Open API spec.

### Docs

Document that the developer should make sure the merge key can uniquely identify an entry in all cases.

## Version Skew and Backward Compatibility

It is fully backward compatibility,
because there are no patch format changes and no changes to the existing APIs.
