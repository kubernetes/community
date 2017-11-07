# Subresources for CustomResources

Authors: @nikhita, @sttts

## Table of Contents

1. [Abstract](#abstract)
2. [Goals](#goals)
3. [Non-Goals](#non-goals)
4. [Proposed Extension of CustomResourceDefinition](#proposed-extension-of-customresourcedefinition)
    1. [API Types](#api-types)
    2. [Feature Gate](#feature-gate)
5. [Semantics](#semantics)
    1. [Validation Behavior](#validation-behavior)
        1. [Status](#status)
        2. [Scale](#scale)
    2. [Status Behavior](#status-behavior)
    3. [Scale Behavior](#scale-behavior)
        1. [Status Replicas Behavior](#status-replicas-behavior)
        2. [Selector Behavior](#selector-behavior)
4. [Implementation Plan](#implementation-plan)
5. [Alternatives](#alternatives)
    1. [Scope](#scope)

## Abstract

[CustomResourceDefinitions](https://github.com/kubernetes/community/pull/524) (CRDs) were introduced in 1.7. The objects defined by CRDs are called CustomResources (CRs). Currently, we do not provide subresources for CRs.

However, it is one of the [most requested features](https://github.com/kubernetes/kubernetes/issues/38113) and this proposal seeks to add  `/status` and `/scale` subresources for CustomResources.

## Goals

1. Support status/spec split for CustomResources:
    1. Status changes are ignored on the main resource endpoint.
    2. Support a `/status` subresource HTTP path for status changes.
    3. `metadata.Generation` is increased only on spec changes.
2. Support a `/scale` subresource for CustomResources.
3. Maintain backward compatibility by allowing CRDs to opt-in to enable subresources.
4. If a CustomResource is already structured using spec/status, allow it to easily transition to use the `/status` and `/scale` endpoint.
5. Work seamlessly with [JSON Schema validation](https://github.com/kubernetes/community/pull/708).

## Non-Goals

1. Allow defining arbitrary subresources i.e. subresources except `/status` and `/scale`.

## Proposed Extension of CustomResourceDefinition

### API Types

The addition of the following external types in `apiextensions.k8s.io/v1beta1` is proposed:

```go
type CustomResourceDefinitionSpec struct {
    ...
    // SubResources describes the subresources for CustomResources
    // This field is alpha-level and should only be sent to servers that enable
    // subresources via the CurstomResourceSubResources feature gate.
    // +optional
    SubResources *CustomResourceSubResources `json:"subResources,omitempty"`
}

// CustomResourceSubResources defines the status and scale subresources for CustomResources.
type CustomResourceSubResources struct {
    // Status denotes the status subresource for CustomResources
    Status *CustomResourceSubResourceStatus `json:"status,omitempty"`
    // Scale denotes the scale subresource for CustomResources
    Scale  *CustomResourceSubResourceScale  `json:"scale,omitempty"`
}

// CustomResourceSubResourceStatus defines how to serve the HTTP path <CR Name>/status.
type CustomResourceSubResourceStatus struct {
}

// CustomResourceSubResourceScale defines how to serve the HTTP path <CR name>/scale.
type CustomResourceSubResourceScale struct {
    // required, e.g. “.spec.replicas”. Must be under `.spec`.
    // Only JSON paths without the array notation are allowed.
    SpecReplicasPath string `json:"specReplicasPath"`
    // optional, e.g. “.status.replicas”. Must be under `.status`.
    // Only JSON paths without the array notation are allowed.
    StatusReplicasPath string `json:"statusReplicasPath,omitempty"`
    // optional, e.g. “.spec.labelSelector”. Must be under `.spec`.
    // Only JSON paths without the array notation are allowed.
    LabelSelectorPath string `json:"labelSelectorPath,omitempty"`
    // ScaleGroupVersion denotes the GroupVersion of the Scale
    // object sent as the payload for /scale. It allows transition
    // to future versions easily.
    // Today only autoscaling/v1 is allowed.
    ScaleGroupVersion schema.GroupVersion `json:"groupVersion"`
}
```

### Feature Gate

The `SubResources` field in `CustomResourceDefinitionSpec` will be gated under the `CustomResourceSubResources` alpha feature gate.
If the gate is not open, the value of the new field within `CustomResourceDefinitionSpec` is dropped on creation and updates of CRDs.

### Scale type

The `Scale` object is the payload sent over the wire for `/scale`. The [polymorphic `Scale` type](https://github.com/kubernetes/kubernetes/pull/53743) i.e. `autoscaling/v1.Scale` is used for the `Scale` object.

Since the GroupVersion of the `Scale` object is specified in `CustomResourceSubResourceScale`, transition to future versions (eg `autoscaling/v2.Scale`) can be done easily.

Note: If `autoscaling/v1.Scale` is deprecated, then it would be deprecated here as well.

## Semantics

### Validation Behavior

#### Status

The status endpoint of a CustomResource receives a full CR object. Changes outside of the `.status` subpath are ignored.
For validation, the JSON Schema present in the CRD is validated only against the `.status` subpath.

To validate only against the schema for the `.status` subpath, `oneOf` and `anyOf` constructs are not allowed within the root of the schema, but only under a properties sub-schema (with this restriction, we can project a schema to a sub-path). The following is forbidden in the CRD spec:

```yaml
validation:
    openAPIV3Schema:
        oneOf:
            ...
```

**Note**: The restriction for `oneOf` and `anyOf` allows us to write a projection function `ProjectJSONSchema(schema *JSONSchemaProps, path []string) (*JSONSchemaProps, error)` that can be used to apply a given schema for the whole object to only the sub-path `.status` or `.spec`.

#### Scale

Moreover, if the scale subresource is enabled:

On update, we copy the values from the `Scale` object into the specified paths in the CustomResource, if the path is set (`StatusReplicasPath` and `LabelSelectorPath` are optional).
If `StatusReplicasPath` or `LabelSelectorPath` is not set, we validate that the value in `Scale` is also not specified and return an error otherwise.

On `get` and on `update` (after copying the values into the CustomResource as described above), we verify that:

- The value at the specified JSON Path `SpecReplicasPath` (e.g. `.spec.replicas`) is a non-negative integer value and is not empty.

- The value at the optional JSON Path `StatusReplicasPath` (e.g. `.status.replicas`) is an integer value if it exists (i.e. this can be empty).

- The value at the optional JSON Path `LabelSelectorPath` (e.g. `.spec.labelSelector`) is a valid label selector if it exists (i.e. this can be empty).

**Note**: The values at the JSON Paths specified by `SpecReplicasPath`, `LabelSelectorPath` and `StatusReplicasPath` are also validated with the same rules when the whole object or, in case the `/status` subresource is enabled, the `.status` sub-object is updated.

### Status Behavior

If the `/status` subresource is enabled, the following behaviors change:

- The main resource endpoint will ignore all changes in the status subpath.
(note: it will **not** reject requests which try to change the status, following the existing semantics of other resources).

- The `.metadata.generation` field is updated if and only if the value at the `.spec` subpath changes.
Additionally, if the spec does not change, `.metadata.generation` is not updated.

- The `/status` subresource receives a full resource object, but only considers the value at the `.status` subpath for the update.
The value at the `.metadata` subpath is **not** considered for update as decided in https://github.com/kubernetes/kubernetes/issues/45539.

Both the status and the spec (and everything else if there is anything) of the object share the same key in the storage layer, i.e. the value at  `.metadata.resourceVersion` is increased for any kind of change. There is no split of status and spec in the storage layer.

The `/status` endpoint supports both `get` and `update` verbs.

### Scale Behavior

The number of CustomResources can be easily scaled up or down depending on the replicas field present in the `.spec` subpath.

Only `ScaleSpec.Replicas` can be written. All other values are read-only and changes will be ignored. i.e. upon updating the scale subresource, two fields are modified:

1. The replicas field is copied back from the `Scale` object to the main resource as specified by `SpecReplicasPath` in the CRD, e.g.  `.spec.replicas = scale.Spec.Replicas`.

2. The resource version is copied back from the `Scale` object to the main resource before writing to the storage: `.metadata.resourceVersion = scale.ResourceVersion`.
In other words, the scale and the CustomResource share the resource version used for optimistic concurrency.
Updates with outdated resource versions are rejected with a conflict error, read requests will return the resource version of the CustomResource.

The `/scale` endpoint supports both `get` and `update` verbs.

#### Status Replicas Behavior

As only the `scale.Spec.Replicas` field is to be written to by the CR user, the user-provided controller (not any generic CRD controller) counts its children and then updates the controlled object by writing to the `/status` subresource, i.e. the `scale.Status.Replicas` field is read-only.

#### Selector Behavior

`CustomResourceSubResourceScale.LabelSelectorPath` is the label selector over CustomResources that should match the replicas count.
The value in the `Scale` object is one-to-one the value from the CustomResource if the label selector is non-empty.
Intentionally we do not default it to another value from the CustomResource (e.g. `.spec.template.metadata.labels`) as this turned out to cause trouble (e.g. in `kubectl apply`) and it is generally seen as a wrong approach with existing resources.

## Implementation Plan

The `/scale` and `/status` subresources are mostly distinct. It is proposed to do the implementation in two phases (the order does not matter much):

1. `/status` subresource
2. `/scale` subresource

## Alternatives

### Scope

In this proposal we opted for an opinionated concept of subresources i.e. we restrict the subresource spec to the two very specific subresources: `/status` and `/scale`.
We do not aim for a more generic subresource concept. In Kubernetes there are a number of other subresources like `/log`, `/exec`, `/bind`. But their semantics is much more special than `/status` and `/scale`.
Hence, we decided to leave those other subresources to the domain of User provided API Server (UAS) instead of inventing a more complex subresource concept for CustomResourceDefinitions.

**Note**: The types do not make the addition of other subresources impossible in the future.

We also restrict the JSON path for the status and the spec within the CustomResource.
We could make them definable by the user and the proposed types actually allow us to open this up in the future.
For the time being we decided to be opinionated as all status and spec subobjects in existing types live under `.status` and `.spec`. Keeping this pattern imposes consistency on user provided CustomResources as well.
