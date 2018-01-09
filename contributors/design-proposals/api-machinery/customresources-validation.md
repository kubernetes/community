# Validation for CustomResources

Authors: @nikhita, @sttts, some ideas integrated from @xiao-zhou’s proposal<sup id="f1">[1](#footnote1)</sup>


## Table of Contents

1. [Overview](#overview)
2. [Background](#background)
    1. [Goals](#goals)
    2. [Non-Goals](#non-goals)
3. [Proposed Extension of CustomResourceDefinition](#proposed-extension-of-customresourcedefinition)
    1. [API Types](#api-types)
    2. [Examples](#examples)
        1. [JSON-Schema](#json-schema)
        2. [Error messages](#error-messages)
4. [Validation Behavior](#validation-behavior)
    1. [Metadata](#metadata)
    2. [Server-Side Validation](#server-side-validation)
    3. [Client-Side Validation](#client-side-validation)
    4. [Comparison between server-side and client-side Validation](#comparison-between-server-side-and-client-side-Validation)
    5. [Existing Instances and changing the Schema](#existing-instances-and-changing-the-schema)
    6. [Outlook to Status Sub-Resources](#outlook-to-status-sub-resources)
    7. [Outlook Admission Webhook](#outlook-admission-webhook)
5. [Implementation Plan](#implementation-plan)
6. [Appendix](#appendix)
    1. [Expressiveness of JSON-Schema](#expressiveness-of-json-schema)
    2. [JSON-Schema Validation Runtime Complexity](#json-schema-validation-runtime-complexity)
    3. [Alternatives](#alternatives)
        1. [Direct Embedding of the Schema into the Spec](#direct-embedding-of-the-schema-into-the-spec)
        2. [External CustomResourceSchema Type](#external-customresourceschema-type)

## Overview

This document proposes the design and describes a way to add JSON-Schema based validation for Custom Resources.

## Background

ThirdPartyResource (TPR) is deprecated and CustomResourceDefinition (CRD) is the successor which solves the fundamental [issues](https://github.com/kubernetes/features/issues/95) of TPRs to form a stable base for further features.

Currently we do not provide validation for CustomResources (CR), i.e. the CR payload is free-form JSON. However, one of the most requested [[1](https://github.com/kubernetes/features/issues/95#issuecomment-296416969)][[2](https://github.com/kubernetes/features/issues/95#issuecomment-298791881)] features is validation and this proposal seeks to add it.

## Goals

1. To provide validation for CustomResources using a declarative specification language for JSON data.
2. To keep open the door to add other validation mechanisms later.<sup id="f2">[2](#footnote2)</sup>
3. To allow server-side validation.
4. To be able to integrate into the existing client-side validation of kubectl.
5. To be able to define defaults in the specification (at least in a follow-up after basic validation support).

## Non-Goals

1. The JSON-Schema specs can be used for creating OpenAPI documentation for CRs. The format is compatible but we won’t propose an implementation for that.
2. A turing-complete specification language is not proposed. Instead a declarative way is proposed to express the vast majority of validations.
3. For now, CRD only allows 1 version at a time. Supporting multiple versions of CRD and/or conversion of CRD is not within the scope of this proposal.

## Proposed Extension of CustomResourceDefinition

We propose to add a field `validation` to the spec of a CustomResourceDefinition. As a first validation format we propose to use [JSON-Schema](http://json-schema.org/) under `CRD.Spec.Validation.JSONSchema`.

JSON-Schema is a [standardized](https://tools.ietf.org/html/draft-zyp-json-schema-04) declarative specification language. Different keywords may be utilized to put constraints on the data. Thus it provides ways to make assertions about what a valid document must look like. 

It is already used in Swagger/OpenAPI specs in Kubernetes and hence such a CRD specification integrates cleanly into the existing infrastructure of the API server which serves these specifications,
* into kubectl which is able to verify YAML and JSON objects against the returned specification.
* With the https://github.com/go-openapi/validate library, we have a powerful JSON-Schema validator which can be used client and server-side.

## API Types

The schema is referenced in [`CustomResourceDefinitionSpec`](https://github.com/kubernetes/kubernetes/commit/0304ef60a210758ab4ac43a468f8a5e19f39ff5a#diff-0e64a9ef2cf809a2a611b16fd44d22f8). `Validation` is of the type `CustomResourceValidation`. The JSON-Schema is stored in a field of `Validation`. This way we can make the validation generic and add other validations in the future as well.

The schema types follow those of the OpenAPI library, but we decided to define them independently for the API to have full control over the serialization and versioning. Hence, it is easy to convert our types into those used for validation or to integrate them into an OpenAPI spec later.

Reference http://json-schema.org is also used by OpenAPI. We propose this as there are implementations available in Go and with OpenAPI, we will also be able to serve OpenAPI specs for CustomResourceDefinitions.

```go
// CustomResourceSpec describes how a user wants their resource to appear
type CustomResourceDefinitionSpec struct {
    Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
    Version string `json:"version" protobuf:"bytes,2,opt,name=version"`
    Names CustomResourceDefinitionNames `json:"names" protobuf:"bytes,3,opt,name=names"`
    Scope ResourceScope `json:"scope" protobuf:"bytes,8,opt,name=scope,casttype=ResourceScope"`
    // Validation describes the validation methods for CustomResources
    Validation CustomResourceValidation `json:"validation,omitempty"`
}

// CustomResourceValidation is a list of validation methods for CustomResources
type CustomResourceValidation struct {
    // JSONSchema is the JSON Schema to be validated against.
    // Can add other validation methods later if needed.
    JSONSchema *JSONSchemaProps `json:"jsonSchema,omitempty"`
}

// JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
type JSONSchemaProps struct {
	ID                   string                     `json:"id,omitempty"`
	Schema               JSONSchemaURL              `json:"-,omitempty"`
	Ref                  JSONSchemaRef              `json:"-,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Type                 StringOrArray              `json:"type,omitempty"`
	Format               string                     `json:"format,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Default              interface{}                `json:"default,omitempty"`
	Maximum              *float64                   `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                       `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64                   `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                       `json:"exclusiveMinimum,omitempty"`
	MaxLength            *int64                     `json:"maxLength,omitempty"`
	MinLength            *int64                     `json:"minLength,omitempty"`
	Pattern              string                     `json:"pattern,omitempty"`
	MaxItems             *int64                     `json:"maxItems,omitempty"`
	MinItems             *int64                     `json:"minItems,omitempty"`
	// disable uniqueItems for now because it can cause the validation runtime
	// complexity to become quadratic.
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	MultipleOf           *float64                   `json:"multipleOf,omitempty"`
	Enum                 []interface{}              `json:"enum,omitempty"`
	MaxProperties        *int64                     `json:"maxProperties,omitempty"`
	MinProperties        *int64                     `json:"minProperties,omitempty"`
	Required             []string                   `json:"required,omitempty"`
	Items                *JSONSchemaPropsOrArray    `json:"items,omitempty"`
	AllOf                []JSONSchemaProps          `json:"allOf,omitempty"`
	OneOf                []JSONSchemaProps          `json:"oneOf,omitempty"`
	AnyOf                []JSONSchemaProps          `json:"anyOf,omitempty"`
	Not                  *JSONSchemaProps           `json:"not,omitempty"`
	Properties           map[string]JSONSchemaProps `json:"properties,omitempty"`
	AdditionalProperties *JSONSchemaPropsOrBool     `json:"additionalProperties,omitempty"`
	PatternProperties    map[string]JSONSchemaProps `json:"patternProperties,omitempty"`
	Dependencies         JSONSchemaDependencies     `json:"dependencies,omitempty"`
	AdditionalItems      *JSONSchemaPropsOrBool     `json:"additionalItems,omitempty"`
	Definitions          JSONSchemaDefinitions      `json:"definitions,omitempty"`
}

// JSONSchemaRef represents a JSON reference that is potentially resolved.
// It is marshaled into a string using a custom JSON marshaller.
type JSONSchemaRef struct {
	ReferencePointer JSONSchemaPointer
	HasFullURL       bool
	HasURLPathOnly   bool
	HasFragmentOnly  bool
	HasFileScheme    bool
	HasFullFilePath  bool
}

// JSONSchemaPointer is the JSON pointer representation.
type JSONSchemaPointer struct {
	ReferenceTokens []string
}

// JSONSchemaURL represents a schema url. Defaults to JSON Schema Specification Draft 4.
type JSONSchemaURL string

const (
	// JSONSchemaDraft4URL is the url for JSON Schema Specification Draft 4.
	JSONSchemaDraft4URL SchemaURL = "http://json-schema.org/draft-04/schema#"
)

// StringOrArray represents a value that can either be a string or an array of strings.
// Mainly here for serialization purposes.
type StringOrArray []string

// JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps
// or an array of JSONSchemaProps. Mainly here for serialization purposes.
type JSONSchemaPropsOrArray struct {
	Schema      *JSONSchemaProps
	JSONSchemas []JSONSchemaProps
}

// JSONSchemaPropsOrBool represents JSONSchemaProps or a boolean value.
// Defaults to true for the boolean property.
type JSONSchemaPropsOrBool struct {
	Allows bool
	Schema *JSONSchemaProps
}

// JSONSchemaDependencies represent a dependencies property.
type JSONSchemaDependencies map[string]JSONSchemaPropsOrStringArray

// JSONSchemaPropsOrStringArray represents a JSONSchemaProps or a string array.
type JSONSchemaPropsOrStringArray struct {
	Schema   *JSONSchemaProps
	Property []string
}

// JSONSchemaDefinitions contains the models explicitly defined in this spec.
type JSONSchemaDefinitions map[string]JSONSchemaProps
```

Note: A reflective test to check for drift between the types here and the OpenAPI types for runtime usage will be added.

## Examples

### JSON-Schema

The following example illustrates how a schema can be used in `CustomResourceDefinition`. It shows various restrictions that can be achieved for validation using JSON-Schema.

```json
{
    "apiVersion": "apiextensions.k8s.io/v1beta1",
    "kind": "CustomResourceDefinition",
    "metadata": {
        "name": "noxus.mygroup.example.com"
    },
    "spec": {
        "group": "mygroup.example.com",
        "version": "v1alpha1",
        "scope": "Namespaced",
        "names": {
            "plural": "noxus",
            "singular": "noxu",
            "kind": "Noxu",
            "listKind": "NoxuList"
        },
        "validation": {
            "jsonSchema": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "type": "object",
                "description": "Noxu is a kind of Custom Resource which has only fields that are specified",
                "required": [
                    "alpha",
                    "beta",
                    "gamma",
                    "delta",
                    "epsilon",
                    "zeta"
                ],
                "properties": {
                    "alpha": {
                        "description": "Alpha is an alphanumeric string with underscores which defaults to foo_123",
                        "type": "string",
                        "pattern": "^[a-zA-Z0-9_]*$",
                        "default": "foo_123"
                    },
                    "beta": {
                        "description": "We need at least 10 betas. If not specified, it defaults to 10.",
                        "type": "number",
                        "minimum": 10,
                        "default": 10
                    },
                    "gamma": {
                        "description": "Gamma is restricted to foo, bar and baz",
                        "type": "string",
                        "enum": [
                            "foo",
                            "bar",
                            "baz"
                        ]
                    },
                    "delta": {
                        "description": "Delta is a string with a maximum length of 5 or a number with a minimum value of 0",
                        "anyOf": [
                            {
                                "type": "string",
                                "maxLength": 5
                            },
                            {
                                "type": "number",
                                "minimum": 0
                            }
                        ]
                    },
                    "epsilon": {
                        "description": "Epsilon is either of type one zeta or two zeta",
                        "allOf": [
                            {
                                "$ref": "#/definitions/zeta"
                            },
                            {
                                "properties": {
                                    "type": {
                                        "enum": [
                                            "one",
                                            "two"
                                        ]
                                    }
                                },
                                "required": [
                                    "type"
                                ],
                                "additionalProperties": false
                            }
                        ]
                    },
                    "additionalProperties": false,
                    "definitions": {
                        "zeta": {
                            "description": "Every zeta needs to have foo, bar and baz",
                            "type": "object",
                            "properties": {
                                "foo": {
                                    "type": "string"
                                },
                                "bar": {
                                    "type": "number"
                                },
                                "baz": {
                                    "type": "boolean"
                                }
                            },
                            "required": [
                                "foo",
                                "bar",
                                "baz"
                            ],
                            "additionalProperties": false
                        }
                    }
                }
            }
        }
    }
}
```

### Error messages

The following examples illustrate the type of validation errors generated by using the go-openapi validate library.

The description is not taken into account, but a better error output can be easily [added](https://github.com/go-openapi/errors/blob/master/headers.go#L23) to go-openapi.

1. `data.foo in body should be at least 4 chars long`
2. `data.foo in body should be greater than or equal to 10`
3. `data.foo in body should be one of [bar baz]`
4. `data.foo in body must be of type integer: "string"`
5. `data.foo in body should match '^[a-zA-Z0-9_]*$'`
6. `data.foo in body is required`
7. When foo validates if it is a multiple of 3 and 5:
```
data.foo in body should be a multiple of 5
data.foo in body should be a multiple of 3
must validate all the schemas (allOf)
```

## Validation Behavior

The schema will be described in the `CustomResourceDefinitionSpec`. The validation will be carried out using the [go-openapi validation library](https://github.com/go-openapi/validate).
 
While creating/updating the CR, the metadata is first validated. To validate the CR against the spec in the CRD, we _must_ have server-side validation and we _can_ have client-side validation.

### Metadata

ObjectMeta and TypeMeta are implicitly specified. They do not have to be added to the JSON-Schema of a CRD. The validation already happens today as part of the apiextensions-apiserver REST handlers.

### Server-Side Validation

The server-side validation is carried out after sending the request to the apiextensions-apiserver, i.e. inside the CREATE and UPDATE handlers for CRs.

We do a schema pass there using the https://github.com/go-openapi/validate validator with the provided schema in the corresponding CRD. Validation errors are returned to the caller as for native resources.

JSON-Schema also allows us to reject additional fields that are not defined in the schema and only allow the fields that are specified. This can be achieved by using `"additionalProperties": false` in the schema. However, there is danger in allowing CRD authors to set `"additionalProperties": false` because it breaks version skew (new client can send new optional fields to the old server). So we should not allow CRD authors to set `"additionalProperties": false`.

### Client-Side Validation

The client-side validation is carried out before sending the request to the api-server, or even completely offline. This can be achieved while creating resources through the client i.e. kubectl using the --validate option.

If the API type serves the JSON-Schema in the swagger spec, the existing kubectl code will already be able to also validate CRs. This will be achieved as a follow-up.

### Comparison between server-side and client-side Validation

The table below shows the cases when server-side and client-side validation methods are applicable.

| Case                                               | Server-Side  | Client-Side   |
|:--------------------------------------------------:|:------------:|:-------------:|
| Kubectl create/edit/replace with validity feedback | &#x2713;     | &#x2713;      | 
| Custom controller creates/updates CRs              | &#x2713;     | &#x2717;      |
| CRs are created by an untrusted party              | &#x2713;     | &#x2717;      |
| Not making validation for CRs a special case       | &#x2713;     | &#x2717;      |

The above table is an evidence that we need server-side validation as well, next to the client-side validation we easily get, nearly for free, by serving Swagger/OpenAPI specs in apiextension-apiserver. 

This is especially true in situations when CRs are used by components that are out of the control of the admin. Example: A user can create a database CR for a Database-As-A-Service. In this case, only server-side validation can give confidence that the CRs are well formed.

### Existing Instances and changing the Schema

If the schema is made stricter later, the existing CustomResources might no longer comply with the spec. This will make them unchangeable and essentially read-only.

To avoid this, it is the responsibility of the user to make sure that any changes made to the schema are such that the existing CustomResources remain validated. 
 
Note:

1. This is the same behavior that we require for native resources. Validation cannot be made stricter in later Kubernetes versions without breaking compatibility.

2. For migration of CRDs with no validation to CRDs with validation, we can create a controller that will validate and annotate invalid CRs once the spec changes, so that the custom controller can choose to delete them (this is also essentially the status condition of the CRD). This can be achieved, but it is not part of the proposal.

### Outlook to Status Sub-Resources

As another most-wanted feature, a Status sub-resource might be proposed and implemented for CRDs. The JSON-Schema proposed here might as well cover the Status field of a CR. For now this is not handled or validated in a particular way.
 
When the Status sub-resource exists some day, the /status endpoint will receive a full CR object, but only the status field is to be validated. We propose to enforce the JSON-Schema structure to be of the shape:

```json
{"type":"object", "properties":{"status": ..., "a": ..., "b": ...}}
```
Then we can validate the status against the sub-schema easily. Hence, this proposal will be compatible with a later sub-resource extension.

### Outlook Admission Webhook

Apiextensions-apiserver uses the normal REST endpoint implementation and only customizes the registry and the codecs. The admission plugins are inherited from the kube-apiserver (when running inside of it via apiserver delegation) and therefore they are supposed to apply to CRs as well.

It is [verified](https://github.com/kubernetes/kubernetes/pull/47252) that CRDs work well with initializers. It is also expected that webhook admission prototyped at https://github.com/kubernetes/kubernetes/pull/46316 will work with CRs out of the box. Hence, for more advanced validation webhook admission is an option as well (when it is merged).

JSON-Schema based validation does not preclude implementation of other validation methods. Hence, advanced webhook-based validation can also be implemented in the future.

## Implementation Plan

The implementation is planned in the following steps:

1. Add the proposed types to the v1beta1<sup id="f3">[3](#footnote3)</sup> version of the CRD type.
2. Add a validation step to the CREATE and UPDATE REST handlers of the apiextensions-apiserver.

Independently, from 1. and 2. add defaulting support:

3. [Add defaulting support to go-openapi](https://github.com/go-openapi/validate/pull/27). Before this PR, we will reject JSON-Schemas which define defaults.
 
As an optional follow-up, we can implement the OpenAPI part and with that enable client-side validation:

4. Export the JSON-Schema via a dynamically served OpenAPI spec.

## Appendix

### Expressiveness of JSON-Schema

The following example properties cannot be expressed using JSON-Schema:
1. “In a PodSpec, for each `spec.Containers[*].volumeMounts[*].Name` there must be a `spec.Volumes[*].Name`”
2. “The volume names in `PodSpec.Volumes` are unique” (`uniqueItems` only compares the complete objects, it cannot compare by key)
 
Different versions within one CRD with a custom version field (i.e. not the one in apiVersion) **can** be expressed:

```json
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "child_schema",
    "type": "object",
    "anyOf": [
        {
            "properties": {
                "version": {
                    "type": "string",
                    "pattern": "^a$"
                },
                "spec": {
                    "type": "object",
                    "properties": {
                        "foo": {}
                    },
                    "additionalProperties": false,
                },
            }
        },
        {
            "properties": {
                "version": {
                    "type": "string",
                    "pattern": "^b$"
                },
                "spec": {
                    "type": "object",
                    "properties": {
                        "bar": {}
                    },
                    "additionalProperties": false,
                },
            }
        }
    ],
}
```

This validates: 
* `{"version": "a",  "spec": {"foo": 42}}`
* `{"version": "b",  "spec": {"bar": 42}}`

but not:
* `{"version": "a",  "spec": {"bar": 42}}`.
 
Note: this is a workaround while we do not support multiple versions and conversion for custom resources.

### JSON-Schema Validation Runtime Complexity

Following “JSON: data model, query languages and schema specification<sup id="f4">[4](#footnote4)</sup>” and “Formal Specification, Expressiveness and Complexity analysis for JSON Schema<sup id="f5">[5](#footnote5)</sup>”, JSON-Schema validation
* without the uniqueItems operator and 
* without recursion for the $ref operator
has linear runtime in the size of the JSON input and the size of the schema (Th. 1 plus Prop. 7).
 
If we allow uniqueItems, the runtime complexity becomes quadratic in the size of the JSON input. Hence, we might want to consider forbidding the uniqueItems operator in order to avoid DDoS attacks, at least if the schema definitions of CRDs cannot be trusted.

The CRD JSON-Schema will be validated to have neither recursion, nor `uniqueItems=true` being set.

### Alternatives

#### Direct Embedding of the Schema into the Spec

An alternative approach to describe the schema in the spec can be as shown below. We directly specify the schema in the spec without the using a Validation field. While simpler, this will limit later extensions, e.g. with non-declarative validation.

```go
// CustomResourceSpec describes how a user wants their resource to appear
type CustomResourceDefinitionSpec struct {
    Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
    Version string `json:"version" protobuf:"bytes,2,opt,name=version"`
    Names CustomResourceDefinitionNames `json:"names" protobuf:"bytes,3,opt,name=names"`
    Scope ResourceScope `json:"scope" protobuf:"bytes,8,opt,name=scope,casttype=ResourceScope"`
    // Schema is the JSON-Schema to be validated against.
    Schema JSONSchema
}
```

#### External CustomResourceSchema Type

In this proposal the JSON-Schema is directly stored in the CRD. Alternatively, one could create a separate top-level API type CustomResourceValidator and reference this from a CRD. Compare @xiao-zhou’s [proposal](https://docs.google.com/document/d/1lKJf9pYBNRcbM7il1VjSJNMDLaf3cFPnquIPPGbEjr4/) for a more detailed sketch of this idea.
 
We do not follow the idea of separate API types in this proposal because CustomResourceDefinitions are highly coupled in practice with the validation of the instances. It doesn’t look like a common use-case to reference a schema from different CRDs and to modify the schema for all of them concurrently.

Hence, the additional complexity for an extra type doesn’t look to be justified.


#### Footnotes

<a name="footnote1">1</a>: https://docs.google.com/document/d/1lKJf9pYBNRcbM7il1VjSJNMDLaf3cFPnquIPPGbEjr4 [↩](#f1)

<a name="footnote2">2</a>: Admission webhooks and embedded programming languages like JavaScript or LUA have been discussed. [↩](#f2)

<a name="footnote3">3</a>: It is common to have alpha fields in beta objects in Kubernetes, compare: FlexVolume, component configs. [↩](#f3)

<a name="footnote4">4</a>: https://arxiv.org/pdf/1701.02221.pdf [↩](#f4)

<a name="footnote5">5</a>: https://repositorio.uc.cl/bitstream/handle/11534/16908/000676530.pdf [↩](#f5)