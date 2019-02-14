Add new patchStrategy to clear fields not present in the patch
=============

We introduce a new struct tag `patchStrategy:"retainKeys"` and
a new optional directive `$retainKeys: <list of fields>` in the patch.

The proposal of Full Union is in [kubernetes/community#388](https://github.com/kubernetes/community/pull/388).

| Capability | Supported By This Proposal | Supported By Full Union |
|---|---|---|
| Auto clear missing fields on patch | X | X |
| Merge union fields on patch | X | X |
| Validate only 1 field set on type |  | X |
| Validate discriminator field matches one-of field |  | X |
| Support non-union patchKey | X | TBD |
| Support arbitrary combinations of set fields | X |  |

## Use cases

- As a user patching a map, I want keys mutually exclusive with those that I am providing to automatically be cleared.

- As a user running kubectl apply, when I update a field in my configuration file,
I want mutually exclusive fields never specified in my configuration to be cleared.

## Examples:

- General Example: Keys in a Union are mutually exclusive. Clear unspecified union values in a Union that contains a discriminator.

- Specific Example: When patching a Deployment .spec.strategy, clear .spec.strategy.rollingUpdate
if it is not provided in the patch so that changing .spec.strategy.type will not fail.

- General Example: Keys in a Union are mutually exclusive. Clear unspecified union values in a Union
that does not contain a discriminator.

- Specific Example: When patching a Pod .spec.volume, clear all volume fields except the one specified in the patch.

## Proposed Changes

### APIs

**Scope**:

| Union Type | Supported |
|---|---|
| non-inlined non-discriminated union | Yes |
| non-inlined discriminated union | Yes |
| inlined union with [patchMergeKey](/contributors/devel/sig-architecture/api-conventions.md#strategic-merge-patch) only | Yes |
| other inlined union | No |

For the inlined union with patchMergeKey, we move the tag to the parent struct's instead of
adding some logic to lookup the metadata in go struct of the inline union.
Because the limitation of the latter is that the metadata associated with
the inlined APIs will not be reflected in the OpenAPI schema.

#### Tags

old tags:

1) `patchMergeKey`:
It is the key to distinguish the entries in the list of non-primitive types. It must always be
present to perform the merge on the list of non-primitive types, and will be preserved.

2) `patchStrategy`:
It indicates how to generate and merge a patch for lists. It could be `merge` or `replace`. It is optional for lists.

new tags:

`patchStrategy: "retainKeys"`:

We introduce a new optional directive `$retainKeys` to support the new patch strategy.

`$retainKeys` directive has the following properties:
- It contains a list of strings.
- All fields needing to be preserved must be present in the `$retainKeys` list.
- The fields that are present will be merged with live object.
- All of the missing fields will be cleared when patching.
- All fields in the `$retainKeys` list must be a superset or the same as the fields present in the patch.

A new patch will have the same content as the old patch and an additional new directive.
It will be backward compatible.

#### When the patch doesn't have `$retainKeys` directive

When the patch doesn't have `$retainKeys` directive, even for a type with `patchStrategy: "retainKeys"`,
the server won't treat the patch with the retainKeys logic.

This will guarantee the backward compatibility: old patch behaves the same as before on the new server.

#### When the patch has fields that not present in the `$retainKeys` list

The server will reject the patch in this case.

This is an invalid patch:

```yaml
union:
  $retainKeys:
  - foo
  foo: a
  bar: x
```

#### When the `$retainKeys` list has fields that are not present in the patch

The server will merge the change and clear the fields not present in the `$retainKeys` list

This is a valid patch:
```yaml
union:
  $retainKeys:
  - foo
  - bar
  foo: a
```

#### Examples

1) Non-inlined non-discriminated union:

Type definition:
```go
type ContainerStatus struct {
	...
	// Add patchStrategy:"retainKeys"
	State ContainerState `json:"state,omitempty" protobuf:"bytes,2,opt,name=state" patchStrategy:"retainKeys"``
	...
}
```
Live object:
```yaml
state:
  running:
    startedAt: ...
```
Local file config:
```yaml
state:
  terminated:
    exitCode: 0
    finishedAt: ...
```
Patch:
```yaml
state:
  $retainKeys:
  - terminated
  terminated:
    exitCode: 0
    finishedAt: ...
```
Result after merging
```yaml
state:
  terminated:
    exitCode: 0
    finishedAt: ...
```

2) Non-inlined discriminated union:

Type definition:
```go
type DeploymentSpec struct {
	...
	// Add patchStrategy:"retainKeys"
	Strategy DeploymentStrategy `json:"strategy,omitempty" protobuf:"bytes,4,opt,name=strategy" patchStrategy:"retainKeys"`
	...
}
```
Since there are no fields associated with `recreate` in `DeploymentSpec`, I will use a generic example.

Live object:
```yaml
unionName:
  discriminatorName: foo
  fooField:
    fooSubfield: val1
```
Local file config:
```yaml
unionName:
  discriminatorName: bar
  barField:
    barSubfield: val2
```
Patch:
```yaml
unionName:
  $retainKeys:
  - discriminatorName
  - barField
  discriminatorName: bar
  barField:
    barSubfield: val2
```
Result after merging
```yaml
unionName:
  discriminatorName: bar
  barField:
    barSubfield: val2
```

3) Inlined union with `patchMergeKey` only.
This case is special, because `Volumes` already has a tag `patchStrategy:"merge"`.
We change the tag to `patchStrategy:"merge|retainKeys"`

Type definition:
```go
type PodSpec struct {
	...
	// Add another value "retainKeys" to patchStrategy
	Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge|retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
	...
}
```
Live object:
```yaml
spec:
  volumes:
  - name: foo
    emptyDir:
      medium:
        ...
```
Local file config:
```yaml
spec:
  volumes:
  - name: foo
    hostPath:
      path: ...
```
Patch:
```yaml
spec:
  volumes:
  - $retainKeys:
    - name
    - hostPath
    name: foo
    hostPath:
      path: ...
```
Result after merging
```yaml
spec:
  volumes:
  - name: foo
    hostPath:
      path: ...
```

**Impacted APIs** are listed in the [Appendix](#appendix).

### API server

No required change.
Auto clearing missing fields of a patch relies on package Strategic Merge Patch.
We don't validate only 1 field is set in union in a generic way. We don't validate discriminator
field matches one-of field. But we still rely on hardcoded per field based validation.

### kubectl

No required change.
Changes about how to generate the patch rely on package Strategic Merge Patch.

### Strategic Merge Patch
**Background**
Strategic Merge Patch is a package used by both client and server. A typical usage is that a client
calls the function to calculate the patch and the API server calls another function to merge the patch.

We need to make sure the new client always sends its patches with the `$retainKeys` directive.
When merging, auto clear missing fields of a patch if the patch has a directive `$retainKeys`

### Open API

Update OpenAPI schema.

## Version Skew

The changes are all backward compatible.

Old kubectl vs New server: All behave the same as before, since no new directive in the patch.

New kubectl vs Old server: All behave the same as before, since new directive will not be recognized
by the old server and it will be dropped in conversion.

# Alternatives Considered

# 1. Use directive `$patch: retainKeys` in the patch

Add tags `patchStrategy:"retainKeys"`.
For a given type that has the tag, all keys/fields missing
from the request will be cleared when patching the object.
Each field present in the request will be merged with the live config.

## Analysis

There are 2 reasons of avoiding this logic:
- Using `$patch` as directive key will break backward compatibility.
But can easily be fixed by using a different key, e.g. `retainKeys: true`.
Reason is that `$patch` has been used in earlier releases.
If we add new value to this directive,
the old server will reject the new patch due to not knowing the new value.
- The patch has to include the entire struct to hold the place in a list with `replace` patch strategy,
even though there may be no changes at all.
This is less efficient compared to the approach above.

The proposals below are not mutually exclusive with the proposal above, and maybe can be added at some point in the future.

# 2. Add Discriminators in All Unions/OneOf APIs

Original issue is described in kubernetes/kubernetes#35345

## Analysis

### Behavior

If the discriminator were set, we'd require that the field corresponding to its value were set and the APIServer (registry) could automatically clear the other fields.

If the discriminator were unset, behavior would be as before -- exactly one of the fields in the union/oneof would be required to be set and the operation would otherwise fail validation.

We should set discriminators by default. This means we need to change it accordingly when the corresponding union/oneof fields were set and unset.

## Proposed Changes

### API
Add a discriminator field in all unions/oneof APIs. The discriminator should be optional for backward compatibility. There is an example below, the field `Type` works as a discriminator.
```go
type PersistentVolumeSource struct {
...
	// Discriminator for PersistentVolumeSource, it can be "gcePersistentDisk", "awsElasticBlockStore" and etc.
	// +optional
	Type *string `json:"type,omitempty" protobuf:"bytes,24,opt,name=type"`
}
```

### API Server

We need to add defaulting logic described in the [Behavior](#behavior) section.

### kubectl

No change required on kubectl.

## Summary

Limitation: Server-side automatically clearing fields based on discriminator may be unsafe.

# Appendix

## List of Impacted APIs
In `pkg/api/v1/types.go`:
- [`VolumeSource`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L235):
It is inlined. Besides `VolumeSource`. its parent [Volume](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L222) has `Name`.
- [`PersistentVolumeSource`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L345):
It is inlined. Besides `PersistentVolumeSource`, its parent [PersistentVolumeSpec](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L442) has the following fields:
```go
Capacity ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
// +optional
AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
// +optional
ClaimRef *ObjectReference `json:"claimRef,omitempty" protobuf:"bytes,4,opt,name=claimRef"`
// +optional
PersistentVolumeReclaimPolicy PersistentVolumeReclaimPolicy `json:"persistentVolumeReclaimPolicy,omitempty" protobuf:"bytes,5,opt,name=persistentVolumeReclaimPolicy,casttype=PersistentVolumeReclaimPolicy"`
```
- [`Handler`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L1485):
It is inlined. Besides `Handler`, its parent struct [`Probe`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L1297) also has the following fields:
```go
// +optional
InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds"`
// +optional
TimeoutSeconds int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds"`
// +optional
PeriodSeconds int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds"`
// +optional
SuccessThreshold int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold"`
// +optional
FailureThreshold int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold"`
````
- [`ContainerState`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L1576):
It is NOT inlined.
- [`PodSignature`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/api/v1/types.go#L2953):
It has only one field, but the comment says "Exactly one field should be set". Maybe we will add more in the future? It is NOT inlined.
In `pkg/authorization/types.go`:
- [`SubjectAccessReviewSpec`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/authorization/types.go#L108):
Comments says: `Exactly one of ResourceAttributes and NonResourceAttributes must be set.`
But there are some other non-union fields in the struct.
So this is similar to INLINED struct.
- [`SelfSubjectAccessReviewSpec`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/authorization/types.go#L130):
It is NOT inlined.

In  `pkg/apis/extensions/v1beta1/types.go`:
- [`DeploymentStrategy`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/extensions/types.go#L249):
It is NOT inlined.
- [`NetworkPolicyPeer`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/extensions/v1beta1/types.go#L1340):
It is NOT inlined.
- [`IngressRuleValue`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/extensions/v1beta1/types.go#L876):
It says "exactly one of the following must be set". But it has only one field.
It is inlined. Its parent [`IngressRule`](https://github.com/kubernetes/kubernetes/blob/v1.5.2/pkg/apis/extensions/v1beta1/types.go#L848) also has the following fields:
```go
// +optional
Host string `json:"host,omitempty" protobuf:"bytes,1,opt,name=host"`
```
