Support Union in API Server
=============

Generalize support for Unions in the API server. Instead of having the unionness of the field 
hard coded into the validation logic on a per-field basic, add first class support for validating unions.


## Proposed Changes

### APIs

**Scope**:

| Union Type | Supported | Tag |
|---|---|---|
| non-inlined non-discriminated union | Yes | `union:"oneof"` |
| non-inlined discriminated union | Yes | `union:"discriminator/<discriminatorName>"` |
| inlined union with [patchMergeKey](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#strategic-merge-patch) only | Yes | `union:"inlined/<patchMergeKey>"` |
| other inlined union | No | N/A |

Add a go struct tag / openAPI value indicating the field is a union and what type of union.

- Tag `union:"oneof"` means this is non-inlined non-discriminated union: only one of the fields in this struct can be set.

- Tag `union:"discriminator/<discriminatorName>"` means this field is a union with a discriminator in the struct.

- Tag `union:"inlined/<patchMergeKey>"` means this is an inlined union with only a `patchMergeKey`

Example of non-inlined non-discriminated union:
```go
type ContainerStatus struct {
	...
	// Add union:"oneof"
	State ContainerState `json:"state,omitempty" protobuf:"bytes,2,opt,name=state" union: "oneof"`
	...
}
```
Example of discriminated union:
```go
type DeploymentSpec struct {
	...
	// Add union:"discriminator/type"
	Strategy DeploymentStrategy `json:"strategy,omitempty" protobuf:"bytes,4,opt,name=strategy" union:"discriminator/type"`
	...
}
```

Example of inlined union with `patchMergeKey` only:
```go
type PodSpec struct {
	...
	// Add union:"inlined/name"
	Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge" patchMergeKey:"name" union: "inlined/name" protobuf:"bytes,1,rep,name=volumes"`
	...
}
```

We don't make any changes on other inlined unions.

### Server Changes

- Validate that neither *PATCH* nor *PUT* set multiple values within a Union
  - Add first class support for validating discriminators if there are any.
- *UPDATE* replaces the entire object, so there should be no need to clear anything in this case

For the inlined union that is not supported, we keep the validation code as it is.

### kubectl

When doing a *PATCH* to set one of the fields in a union, clear any other field in the union that
was previously set. The client explicitly expresses setting one field and clearing the other fields.
We provide the users a dry-run mode to let them check what is going to be sent to the server. E.g.
```yaml
containerstate:
  waiting:
    ...
  running: null
  terminated: null
```

### Open API change

We would need to add extensions to the openapi spec we publish. This is something we already need to do for the `patchStrategy` and `mergeKey` struct tags.

### Docs

Update `API-conventions-md` to include:
```
we should avoid adding new inlined unions in the future.
```

## Summary

Limitation: We don't support inlined union types. Because the validator doesn't have
enough metadata to distinguish the inlined union fields and other non-union fields.
