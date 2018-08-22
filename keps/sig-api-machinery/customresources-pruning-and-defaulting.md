---
kep-number: 24
title: Pruning for Custom Resources
status: provisional
authors:
  - "@sttts"
owning-sig: sig-api-machinery
participating-sigs:
  - sig-api-machinery
  - sig-architecture
reviewers:
  - "@deads2k"
  - "@lavalamp"
  - "@liggitt"
  - "@erictune"
  - "@mbohlool"
  - "@apelisse"
approvers:
  - "@deads2k"
  - "@lavalamp"
editor:
  name: "@sttts"
creation-date: 2018-07-31
last-updated: 2018-07-31
---

# Pruning for Custom Resources
    
## Table of Contents
* [Pruning for Custom Resources](#pruning-for-custom-resources)
    * [Table of Contents](#table-of-contents)
    * [Overview](#overview)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
    * [Motivation](#motivation)
        * [Pruning](#pruning)
        * [Mixing of Schema and Value Validation](#mixing-schema-and-value-validation)
    * [Formal Proposal – following pruning option 1]()
        * [Types and Formats](#types-and-formats)
        * [Polymorphic Fields](#polymorphic-fields)
        * [Excluding values from Pruning](#excluding-values-from-pruning)
    * [References](#references)
    * [Alternatives Considered](#alternatives-considered)

## Overview

Native Golang based resources do not persist JSON fields which are not part of the Golang structs which are backing them in the API server memory. This artifact of the use of typed Golang structs inside of the REST implementation has turned into an API convention. Major parts of the Kubernetes depend on this to ensure consistency of data persisted to etcd and returned from the REST API. 

Without consistency of data in etcd, objects can suddenly render unaccessible on version upgrade because unexpected data may break decoding (e.g. in generated, typed clients which are not based on forgiving `unstructured.Unstructured`). Even if persisted data has correct format and decodes correctly, having it not gone through validation and admission when it was stored can break cluster-wide invariants. For example assume the `privileged: true/false` field is added to a type in Kubernetes version X+1. In version X, there is no security check around this. So every user could set that flag if we didn’t drop unknown fields. When the field is added in X+1, that user suddenly has escalated access (note: on read we do not run admission). This is a serious security risk.

CustomResources are persisted as JSON blobs today (with the exception of `ObjectMeta` which is pruned since Kubernetes 1.11), i.e. we do not drop unknown fields, with the described consequences. This proposal is about adding a decoding step named "pruning" into the decoder from JSON to `unstructured.Unstructured` inside the apiextensions-apiserver. This pruning step will drop fields not specified in the OpenAPI validation spec, leading to the same persistence semantics as for native types.

## Goals

* Prune unknown fields from CustomResources silently. Unknown means not specified in the OpenAPI validation spec.
* Allow to opt-out of pruning via the OpenAPI validation spec for a whole subtree of the JSON objects.
* Have simple semantics for pruning.
* Be extensible to defaulting at a later point.

## Non-Goals
* Add a strict mode to the REST API which rejects objects with unknown fields.
* Propose or decide anything about defaulting.

## Motivation
### Pruning
Pruning of a JSON value means to remove "unknown fields". A field (given as a JSON path) is unknown if it is not specified in the OpenAPI validation schema.

A JSON path `<JSON path>.x` is specified in an OpenAPI validation schema if `"x"` is a key in a corresponding `properties` field in the schema.

#### Example 1

>Assume the OpenAPI schema
>
>```json
>{"properties": {"a": {}, "b": {"type": "string"}, "c": {"not": {}}}}
>```
>
>Then a JSON object `{"a":1, "c": 3, "d": 4}` is pruned to `{"a":1, "c":3}`.
>
> Note that the pruned object does not validate.

#### Example 2

> Assume the OpenAPI schema
>
>```json
>{"anyOf": [{"properties": {"a": {}}}, {"properties": {"b": {}}}]}
>```
>
>Then a JSON object `{"a":1, "b": 2, "c": 3}` is pruned to `{"a":1, "b":2}` with the given semantics. Note that also `{"a":1}`, `{"b":2}` would be natural pruning results if we defined "specified" in a different way.

#### Example 3

> Assume the OpenAPI schema
>
>```json
>{
>   "properties": {
>       "a": {}, 
>       "b": {"type": "string", "properties": {"x": {}}}
>   }, 
>   "anyOf": [
>	    { "properties": {"c":{}} }, 
>       { "properties": {"d":{"not":{}}}}
>   ]
>}
>```
>
> Then `a`, `b`, `b.x`, `c`, `d` are specified, `b.y` and `e` are not specified.

**Question:** is it natural semantics to assume `d` to be specified and not to prune it? Note that there is no object with `d` that validates against `"d":{not:{}}`. But due to the `anyOf` there are objects which will validate against the complete schema.

#### Pruning Options

Motivated by the examples, we have different options to define the pruning semantics:
1. Use the described semantics of `specified`.
2. Only consider `properties` fields in the schema which actually successfully validate a given object (then `d` would be pruned in example 3).
3. Only consider `properties` fields outside of `anyOf`, `allOf`, `oneOf`, `not` during pruning.
4. Only consider `properties` fields outside of `anyOf`, `allOf`, `oneOf`, `not` during pruning, but enforce that every `properties` key inside `anyOf`, `allOf`, `oneOf`, `not` also appears outside all of those.

From these options:
1. leads to fields being kept from pruning although they only appear in branches of the OpenAPI validation schema which do not contribute to validation of an object.
2. is ambiguous if you have multiple branches of `anyOf` validating. Should we drop the fields of the first or the second branch?
3. leads to surprising pruning of fields that the user forgot to specify outside of propositional logic of `anyOf`, `allOf`, `oneOf`, `not`.
4. forces the user to re-specify all those properties outside of `anyOf`, `allOf`, `oneOf`, `not` which appear inside of them. The outcome will match that of 1, with the difference that it makes the skeleton explicit.

**Here we propose not to follow 2 and 3 due to ambiguity of 2 and the danger of user mistakes of 3. Both 1 and 4 lead to the same pruning behaviour and only differ in whether the user has to re-specify property keys or whether they are automatically derived.**

### Mixing of Schema and Value Validation

The OpenAPI validation schema mixes the actual structural schema validation (which is usually done by the Golang struct JSON decoding for native types) 
and the value validation (usually done in the validation step of the API server handler pipeline for native types). 

For CRDs we cannot distinguish both. This was first noticed by @lavalamp in https://github.com/kubernetes/kubernetes/pull/64907#issuecomment-397015030.

#### Example 4

> Assume 
>
>```json
>{
>   "properties": {
>       "a": {"type": "string", "pattern": "<some-long-regex-for-ips>", "format": "ip"},
>       "b": {"properties": {"x": {}, "y": {}}}
>   }
>}
>```
>
> The type and properties are usually called schema validation, while regex patterns and formats are about values. The latter would be tested in the validation phase, the former would be checked during decoding for native types in the JSON decoder.

**Remark:** the line between type and format is blurry. Format is optional to be processed by tools, and the value is open ended. In Kube we have some formats (like `date`) which correspond to custom JSON unmarshallers in the native types. That format would be verified during decoding, although technically a date is just a string. So we might want to replicate that behaviour at some point.

For pruning (and possibly later defaulting) we have to apply the OpenAPI validation schema inside of the decoder step (the left box inside apiextensions-apiserver in the figure above). This is considerably earlier than for native types. This has a number of implications:
* We have to apply full OpenAPI value validation (e.g. regular expressions, propositional evaluation) during decoding.
* Our generic registry applies validation after defaulting. We need the other way around for CRDs.
* Our generic registry expects defaulting to always succeed (no error result type). CRD validation needed by OpenAPI defaulting would be able to fail.

To avoid these and to get sane semantics, **we propose to split the CRD validation into two top-level steps:**

1. During decoding we validate using a skeleton schema which lacks value validations and propositional logic (`anyOf`, `allOf`, `oneOf`, `not`). 
2. During standard generic registry validation phase we validate using the full validation schema.

Step 1 is extensible to defaulting based on the skeleton validation result. Step 2 would then catch wrong types of the used defaults.

## Formal Proposal – following pruning option 1

The examples in the motivational section show that the semantics of full OpenAPI validation schemata are not trivial in respect to pruning. To simplify the algorithms and to enforce "sane" schemata which allow to split schema and value validation, we propose to derive a skeleton schema from the full user-given OpenAPI validation schema
* which does not contain value validations
* which is complete enough for pruning (and possibly later defaulting).

**Remark:** we have two options to define and implement pruning:
1. directly on the full OpenAPI validation schema with two custom algorithms doing parallel recursion over the schema and the input object.
2. via an intermediate representation (the skeleton schema) and using go-openapi pruning. Both algorithms are ten-liners based on the go-openapi/validate output. With the skeleton schema the go-openapi pruning algorithm coincides with our pruning option 1.

Both routes lead to the same algorithm: the intermediate representation of the skeleton schema makes merging of OpenAPI validation schema constraints explicit, while the custom algorithms would hide that in its recursion code.

Moreover, the main reason for the intermediate schema: the custom algorithm would have to replicate a lot of the validation logic of go-openapi, e.g. the semantics of `properties`, `additionalPropoerties`, `patternProperties` and the same for items of an array. With route 2 we get all of this for free. 

### Definition: skeleton schema

> For a given OpenAPI validation schema `s` the skeleton schema `skel(s)` is derived by 
> 1. applying `skel` to all elements of `.allOf`, `.anyOf`, `.oneOf` and `.not` giving `s_1, …, s_n`,
> then dropping all fields from `s` other than 
> * `type`, 
> * `items`,
> * `additionalItems`,
> * `properties`,
> * `patternProperties`,
> * `additionalProperties`
> giving `drop(s)`,
> 2. then merging `s_i` into it’s containing object.
>
>    I.e.: `skel(s) := merge(drop(s), s_1, …, s_n)` with
>
>```
>merge(x_1, …, x_n) := { 
>   "type": t if all x_i agree on t as type, undefined otherwise
>   "items": [ merge(x_i1, …, x_ik) ] where s_ij = s_j.items[i] if defined,
>   "additionaItems": merge(p_1, …, p_i),
>       for all x_i with defined additionalItems p_i
>   "properties": { k_i: merge(v_i1, …, v_ik) for keys appearing in x_ij with values x_ij },
>   "patternProperties": { k_i: merge(v_i1, …, v_ik) for keys appearing in x_ij with values x_ij },
>   "additionalProperties": merge(p_1, …, p_i),
>       for all x_i with defined additionalProperties p_i
>}
>```

The skeleton schema especially lacks:
* all value validations like `pattern`, `format`, `minValue`, `maxValue`, ...
* all propositional operators like `anyOf`, `allOf`, `oneOf`, `not`.

The skeleton schema `skel(s)` puts
* less or equal constraints on `type` than `s`.
* every field specified by a `properties` key in `s` is specified in `skel(s)`.

The computation of `skel(s)` is `O(size(s))` and `size(skel(s)) <= size(s)`.

Note that a field might be constrained by the `type` construct in the full OpenAPI validation schema, but not in its skeleton. This is fine because we have to support polymorphic fields like `IntOrString`, but avoid  `allOf`, `anyOf`, `oneOf`, `not` in the skeleton.

**Property:** if the OpenAPI validation schema applies to an object, so does its skeleton schema.

Pruning will be implemented based on the skeleton schema of the specified OpenAPI validation schema.

Optionally for debugging, the skeleton schema derived from the validation schema by the apiextensions-apiserver can be stored in the CRD status.

#### Example 5

>Assume 
>
>```
>s := {
>    "anyOf": [
>        {"properties": {"a": {"type": "string"}}},
>        {"properties": {"a": {"type": "integer"}, "b": {"type": "string"}}}
>    ],
>    "properties": {"c": {}}
>}
>```
>
>Then 
>
>```
>skel(s) = {
>    "properties": {
>	    "a": {},
>	    "b": {"type": "string"},
>	    "c": {}
>    }
>}
>```
>
> The `type` of `a` does not match on all paths. Hence, it is omitted in the skeleton. In contrast, `b` has a unique `type` of `"string"`, so it stays in the skeleton.

### Types and Formats

Note, that having less `type` constraints in the skeleton than in the whole OpenAPI validation schema means that admission and conversion will see possibly wrong types. The final validation in the registry validation phase will check for the complete OpenAPI validation schema and catch those type errors.

**Question:** we could extent the `skel` function to keep a list of `type` values. As long as all branches define the type, we can add `anyOf: [{type: "type1", type: "type2, ….}]` to the skeleton. If one branch does not define the type though, this is still not possible.

In native types, we have custom unmarshallers for date / timestamps. In OpenAPI these would be rendered as `{type: "string", format: "date"}`. With the proposed skeleton algorithm, we would not verify these fields before full OpenAPI validation in the registry. We could move non-contradicting `format` constraints into the skeleton as well, like we do for `type` already. Then admission would be protected from invalid format (if admission uses Golang decoding, this might be relevant to get proper error messages).

### Polymorphic Fields

#### Example 6

> Assume
>
>```
>s := {
>    "anyOf": [
>        {"properties": {"a": {"type": "string"}}},
>        {"additionalProperties": {"type": "integer"}}
>    ],
>}
>```
>
>Then
>
>```
>   skel(s) = {"properties": {}, "additionalProperties": {}}
>```
>
> In the CRD validation, we reject OpenAPI validation schemata like `skel(s)` with `properties` and one of `additionalProperties` or  `patternProperties` being defined at the same time. Having `additionalProperties` or `patternProperties` defined there means to have a `map[string]T` like field where we don’t want to prune the unknown "keys". 
> 
> The schema `s` above means to either have a `map[string]int64` or a `struct` with `a` of type string. This is a special case of polymorphism because for the same JSON path is typed with two different types in the Golang sense (struct and `map[string]T`), but same types in the JSON sense (JSON object).

Hence, we have to add the following restriction to OpenAPI validation schemata:

**Restriction 1 on Object Polymorphism:** reject CRD OpenAPI validation schemata `s` with `skel(s)` having `properties` and one of  `additionalProperties` or `patternProperties` being defined for the same JSON path.

### Excluding values from Pruning

There are cases where parts of an object are verbatim JSON, i.e. without any applied schema and especially without a complete specification which allows to apply pruning.

Hence, we need a mechanism to express that in the OpenAPI validation schema (compare https://github.com/kubernetes/kubernetes/pull/64558#issuecomment-403564033).

**Raw JSON Option 1:** add a format `json`, e.g.:

```json
{"properties": {"x": {"format": "json"}}}
```

In this example "x" is excluded from pruning. Note that you can still use any kind of OpenAPI validation schema constructs to restrict `"x"` further.

Note that we lose expressiveness of the existing `format` strings: we either apply the format to `"json"` or any other pre-defined format. I.e. we cannot express that pruning should be disabled, but if it is an integer, it should be an int32.

**Question:** this feels like a reasonable loss of expressivity. Do we accept that?

**Raw JSON Option 2:** add an extension property, e.g.

```json
{"properties": {"x": {"x-kubernetes-no-pruning": true}}}
```

This would not lead to a loss in expressivity. It is formulated negatively intentionally to have `false` as the default with `omitempty`.

**We propose to follow the second option with `x-kubernetes-no-prune`, because it is more explicit, does not reduce expressivity and does not mangle with the already very vague `format` field definition.**

### Nested x-no-pruning

**Question:** should we support nested `x-kubernetes-no-prune`, i.e. disabling pruning for a sub-object, but re-enable it something deep inside of it? E.g.

```json
{
    "properties": {
        "x": {
            "x-no-pruning": true,
            "properties": {
		        "y": { 
                    "x-kubernetes-no-pruning": false,
                    "properties": { "z": {} }
                }
            }
        }
    }
}
```

If we do, the object 

```json
{
    "a": 1,
    "x": {
        "b": 2,
        "y": {
            "c": 3,
            "z": 42
        }
    }
}
```

would be pruned to  `{"x":{"b": 2, "y":{"z": 42}}}`.

**We propose to disallow nesting of `x-kubernetes-no-prune` and to disallow setting it to false, i.e. `x-kubernetes-no-prune: false`.** We can add nesting later if necessary.

## Opt-in and Opt-out of Pruning on CRD Level

We will add a pruning flag to `CustomResourceDefinitionSpec` of `apiextensions.k8s.io/v1beta1`:

```golang
type CustomResourceDefinitionSpec struct {
  ...
	
  // Prune enables pruning of unspecified fields. Defaults to false.
  // Note: this will default to true in version v1.
  Prune *bool
}
```

I.e. for `apiextensions.k8s.io/v1beta1` this will default to `false`.

For `apiextensions.k8s.io/v1` we will change the default to `true` and forbid `false` during creation and updates. In `v1` the only way to opt-out from pruning is via setting `x-kubernetes-no-prune: true` in the schema.

When [CRD conversion](https://github.com/mbohlool/community/blob/master/contributors/design-proposals/api-machinery/customresource-conversion-webhook.md) is implemented before this KEP, we will add the pruning field to `type CustomResourceDefinitionVersion`, in analogy to subresources and `additionalPrinterColumns`.

## References

* Pruning implementation PR https://github.com/kubernetes/kubernetes/pull/64558 
* [OpenAPI v3 specification](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)
* [JSON Schema](http://json-schema.org/)
* [pruning algorithm in go-openapi](https://github.com/go-openapi/validate/blob/master/post/prune.go)

## Alternatives Considered

* we have explored pruning option 4 in the [GDoc which preceded this KEP](https://docs.google.com/document/d/1rBn6SZM7NsWxzBN41J2kO2Odf07PeGPygatM_1RwofY/edit#heading=h.4qdisqud6z3t), but decided against it as it put a lot of burden on the CRD author. The approach shown in this KEP leads to the same final outcome, but derives the skeleton automatically.
