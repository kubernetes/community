# Changing the API

This document is oriented at developers who want to change existing APIs.
A set of API conventions, which applies to new APIs and to changes, can be
found at [API Conventions](api-conventions.md).

**Table of Contents**

- [So you want to change the API?](#so-you-want-to-change-the-api)
- [Operational overview](#operational-overview)
- [On compatibility](#on-compatibility)
  - [Adding a field](#adding-a-field)
  - [Making a singular field plural](#making-a-singular-field-plural)
    - [Single-Dual ambiguity](#single-dual-ambiguity)
  - [Multiple API versions](#multiple-api-versions)
- [Backward compatibility gotchas](#backward-compatibility-gotchas)
- [Incompatible API changes](#incompatible-api-changes)
- [Changing versioned APIs](#changing-versioned-apis)
  - [Edit types.go](#edit-typesgo)
  - [Edit defaults.go](#edit-defaultsgo)
  - [Edit conversion.go](#edit-conversiongo)
- [Changing the internal structures](#changing-the-internal-structures)
  - [Edit types.go](#edit-typesgo-1)
- [Edit validation.go](#edit-validationgo)
- [Edit version conversions](#edit-version-conversions)
- [Generate Code](#generate-code)
  - [Generate protobuf objects](#generate-protobuf-objects)
  - [Generate Clientset](#generate-clientset)
  - [Generate Listers](#generate-listers)
  - [Generate Informers](#generate-informers)
  - [Edit json (un)marshaling code](#edit-json-unmarshaling-code)
- [Making a new API Version](#making-a-new-api-version)
- [Making a new API Group](#making-a-new-api-group)
- [Update the fuzzer](#update-the-fuzzer)
- [Update the semantic comparisons](#update-the-semantic-comparisons)
- [Implement your change](#implement-your-change)
- [Write end-to-end tests](#write-end-to-end-tests)
- [Examples and docs](#examples-and-docs)
- [Alpha, Beta, and Stable Versions](#alpha-beta-and-stable-versions)
  - [Adding Unstable Features to Stable Versions](#adding-unstable-features-to-stable-versions)
    - [New field in existing API version](#new-field-in-existing-api-version)
    - [Ratcheting validation](#ratcheting-validation)
    - [New enum value in existing field](#new-enum-value-in-existing-field)
    - [New alpha API version](#new-alpha-api-version)

## So you want to change the API?

Before attempting a change to the API, you should familiarize yourself with a
number of existing API types and with the [API conventions](api-conventions.md).
If creating a new API type/resource, we also recommend that you first send a PR
containing just a proposal for the new API types.

The Kubernetes API has two major components - the internal structures and
the versioned APIs. The versioned APIs are intended to be stable, while the
internal structures are implemented to best reflect the needs of the Kubernetes
code itself.

What this means for API changes is that you have to be somewhat thoughtful in
how you approach changes, and that you have to touch a number of pieces to make
a complete change.  This document aims to guide you through the process, though
not all API changes will need all of these steps.

## Operational overview

It is important to have a high level understanding of the API system used in
Kubernetes in order to navigate the rest of this document.

As mentioned above, the internal representation of an API object is decoupled
from any one API version. This provides a lot of freedom to evolve the code,
but it requires robust infrastructure to convert between representations. There
are multiple steps in processing an API operation - even something as simple as
a GET involves a great deal of machinery.

The conversion process is logically a "star" with the internal form at the
center. Every versioned API can be converted to the internal form (and
vice-versa), but versioned APIs do not convert to other versioned APIs directly.
This sounds like a heavy process, but in reality we do not intend to keep more
than a small number of versions alive at once. While all of the Kubernetes code
operates on the internal structures, they are always converted to a versioned
form before being written to storage (disk or etcd) or being sent over a wire.
Clients should consume and operate on the versioned APIs exclusively.

To demonstrate the general process, here is a (hypothetical) example:

   1. A user POSTs a `Pod` object to `/api/v7beta1/...`
   2. The JSON is unmarshalled into a `v7beta1.Pod` structure
   3. Default values are applied to the `v7beta1.Pod`
   4. The `v7beta1.Pod` is converted to an `api.Pod` structure
   5. The `api.Pod` is validated, and any errors are returned to the user
   6. The `api.Pod` is converted to a `v6.Pod` (because v6 is the latest stable
version)
   7. The `v6.Pod` is marshalled into JSON and written to etcd

Now that we have the `Pod` object stored, a user can GET that object in any
supported api version. For example:

   1. A user GETs the `Pod` from `/api/v5/...`
   2. The JSON is read from etcd and unmarshalled into a `v6.Pod` structure
   3. Default values are applied to the `v6.Pod`
   4. The `v6.Pod` is converted to an `api.Pod` structure
   5. The `api.Pod` is converted to a `v5.Pod` structure
   6. The `v5.Pod` is marshalled into JSON and sent to the user

The implication of this process is that API changes must be done carefully and
backward-compatibly.

## On compatibility

Before talking about how to make API changes, it is worthwhile to clarify what
we mean by API compatibility.  Kubernetes considers forwards and backwards
compatibility of its APIs a top priority.  Compatibility is *hard*, especially
handling issues around rollback-safety.  This is something every API change
must consider.

An API change is considered compatible if it:

   * adds new functionality that is not required for correct behavior (e.g.,
does not add a new required field)
   * does not change existing semantics, including:
     * the semantic meaning of default values *and behavior*
     * interpretation of existing API types, fields, and values
     * which fields are required and which are not
     * mutable fields do not become immutable
     * valid values do not become invalid
     * explicitly invalid values do not become valid

Put another way:

1. Any API call (e.g. a structure POSTed to a REST endpoint) that succeeded
before your change must succeed after your change.
2. Any API call that does not use your change must behave the same as it did
before your change.
3. Any API call that uses your change must not cause problems (e.g. crash or
degrade behavior) when issued against an API servers that do not include your
change.
4. It must be possible to round-trip your change (convert to different API
versions and back) with no loss of information.
5. Existing clients need not be aware of your change in order for them to
continue to function as they did previously, even when your change is in use.
6. It must be possible to rollback to a previous version of API server that
does not include your change and have no impact on API objects which do not use
your change.  API objects that use your change will be impacted in case of a
rollback.

If your change does not meet these criteria, it is not considered compatible,
and may break older clients, or result in newer clients causing undefined
behavior.  Such changes are generally disallowed, though exceptions have been
made in extreme cases (e.g. security or obvious bugs).

Let's consider some examples.

### Adding a field

In a hypothetical API (assume we're at version v6), the `Frobber` struct looks
something like this:

```go
// API v6.
type Frobber struct {
  Height int    `json:"height"`
  Param  string `json:"param"`
}
```

You want to add a new `Width` field. It is generally allowed to add new fields
without changing the API version, so you can simply change it to:

```go
// Still API v6.
type Frobber struct {
  Height int    `json:"height"`
  Width  int    `json:"width"`
  Param  string `json:"param"`
}
```

The onus is on you to define a sane default value for `Width` such that rules
#1 and #2 above are true - API calls and stored objects that used to work must
continue to work.

### Making a singular field plural

For your next change you want to allow multiple `Param` values. You can not
simply remove `Param string` and add `Params []string` (without creating a
whole new API version) - that fails rules #1, #2, #3, and #6.  Nor can you
simply add `Params []string` and use it instead - that fails #2 and #6.

You must instead define a new field and the relationship between that field and
the existing field(s).  Start by adding the new plural field:

```go
// Still API v6.
type Frobber struct {
  Height int           `json:"height"`
  Width  int           `json:"width"`
  Param  string        `json:"param"`  // the first param
  Params []string      `json:"params"` // all of the params
}
```

This new field must be inclusive of the singular field.  In order to satisfy
the compatibility rules you must handle all the cases of version skew, multiple
clients, and rollbacks.  This can be handled by admission control or API
registry logic (e.g. strategy) linking the fields together with context from
the API operation to get as close as possible to the user's intentions.

Upon any read operation:
  * If plural is not populated, API logic must populate plural as a one-element
    list, with plural[0] set to the singular value.

Upon any create operation:
  * If only the singular field is specified (e.g. an older client), API logic
    must populate plural as a one-element list, with plural[0] set to the
    singular value.  Rationale: It's an old client and they get compatible
    behavior.
  * If both the singular and plural fields are specified, API logic must
    validate that plural[0] matches the singular value.
  * Any other case is an error and must be rejected.  This includes the case of
    the plural field being specified and the singular not.  Rationale: In an
    update, it's impossible to tell the difference between an old client
    clearing the singular field via patch and a new client setting the plural
    field.  For compatibility, we must assume the former, and we don't want
    update semantics to differ from create (see [Single-Dual
    ambiguity](#single_dual_ambiguity) below.

For the above: "is specified" means the field is present in the user-provided
input (including defaulted fields).

Upon any update operation (including patch):
  * If singular is cleared and plural is not changed, API logic must clear
    plural.  Rationale: It's an old client clearing the field it knows about.
  * If plural is cleared and singular is not changed, API logic must populate
    the new plural with the same values as the old.  Rationale: It's an old
    client which can't send fields it doesn't know about.
  * If the singular field is changed (but not cleared) and the plural field is
    not changed, API logic must populate plural as a one-element list, with
    plural[0] set to the singular value.  Rationale: It's an old client
    changing the field they know about.

Expressed as code, this looks like the following:

```
// normalizeParams adjusts Params based on Param.  This must not consider
// any other fields.
func normalizeParams(after, before *api.Frobber) {
     // Validation  will be called on the new object soon enough.  All this
     // needs to do is try to divine what user meant with these linked fields.
     // The below is verbosely written for clarity.

     // **** IMPORTANT *****
     // As a governing rule. User must either:
     //   a) Use singular field only (old client)
     //   b) Use singular *and* plural fields (new client)

     if before == nil {
         // This was a create operation.

         // User specified singular and not plural (an old client), so we can
         // init plural for them.
         if len(after.Param) > 0 && len(after.Params) == 0 {
             after.Params = []string{after.Param}
             return
         }

         // Either both were specified or both were not.  Catch this in
         // validation.
         return
     }

     // This was an update operation.

     // Plural was cleared by an old client which was trying to patch
     // some field and didn't provide it.
     if len(before.Params) > 0 && len(after.Params) == 0 {
         // If singular is unchanged, then it is an old client trying to
         // patch, and didn't provide plural.  Bring the old value forward.
         if before.Param == after.Param {
             after.Params = before.Params
         }
     }

     if before.Param != after.Param {
         // Singular is changed.

         if len(before.Param) > 0 && len(after.Param) == 0 {
             // If singular was cleared and plural is unchanged, then we can
             // clear plural to match.
             if sameStringSlice(before.Params, after.Params) {
                 after.Params = nil
             }
             // Else they also changed plural - check it in validation.
         } else {
             // If singular was changed (but not cleared) and plural was not,
             // then we can set plural based on singular (same as create).
             if sameStringSlice(before.Params, after.Params) {
                 after.Params = []string{after.Param}
             }
         }
     }
 }
```

Older clients that only know the singular field will continue to succeed and
produce the same results as before the change.  Newer clients can use your
change without impacting older clients.  The API server can be rolled back and
only objects that use your change will be impacted.

Part of the reason for versioning APIs and for using internal types that are
distinct from any one version is to handle growth like this. The internal
representation can be implemented as:

```go
// Internal, soon to be v7beta1.
type Frobber struct {
  Height int
  Width  int
  Params []string
}
```

The code that converts to/from versioned APIs can decode this into the
compatible structure. Eventually, a new API version, e.g. v7beta1,
will be forked and it can drop the singular field entirely.

#### Single-Dual ambiguity

Assume the user starts with:

```
kind: Frobber
height: 42
width: 3
param: "super"
```

On create we can set `params: ["super"]`.

On an unrelated POST (aka replace), an old client would send:

```
kind: Frobber
height: 3
width: 42
param: "super"
```

If we don't require new clients to use both singular and plural fields, a new
client would send:

```
kind: Frobber
height: 3
width: 42
params: ["super"]
```

That seems clear enough - we can assume `param: "super"`.

But the old client could send this, via patch:

```
PATCH  /frobbers/1
{ param: "" }
```

That gets applied to the old object before registry code can see it, and we end up with:

```
kind: Frobber
height: 42
width: 3
params: ["super"]
```

By the previous logic, we would copy `params[0]` to `param` and end up with
`param: "super"`.  But that's not what the user wanted and more importantly is
different than what happened before we pluralized.

To disambiguate that, we require users of plural to always specify singular,
too.

### Multiple API versions

We've seen how to satisfy rules #1, #2, and #3. Rule #4 means that you can not
extend one versioned API without also extending the others. For example, an
API call might POST an object in API v7beta1 format, which uses the new
`Params` field, but the API server might store that object in trusty old v6
form (since v7beta1 is "beta"). When the user reads the object back in the
v7beta1 API it would be unacceptable to have lost all but `Params[0]`. This
means that, even though it is ugly, a compatible change must be made to the v6
API, as above.

For some changes, this can be challenging to do correctly. It may require multiple
representations of the same information in the same API resource, which need to
be kept in sync should either be changed.

For example, let's say you decide to rename a field within the same API
version. In this case, you add units to `height` and `width`. You implement
this by adding new fields:

```go
type Frobber struct {
  Height         *int          `json:"height"`
  Width          *int          `json:"width"`
  HeightInInches *int          `json:"heightInInches"`
  WidthInInches  *int          `json:"widthInInches"`
}
```

You convert all of the fields to pointers in order to distinguish between unset
and set to 0, and then set each corresponding field from the other in the
defaulting logic (e.g. `heightInInches` from `height`, and vice versa).  That
works fine when the user creates a sends a hand-written configuration --
clients can write either field and read either field.

But what about creation or update from the output of a GET, or update via PATCH
(see [In-place updates](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#in-place-updates-of-resources))?
In these cases, the two fields will conflict, because only one field would be
updated in the case of an old client that was only aware of the old field
(e.g. `height`).

Suppose the client creates:

```json
{
  "height": 10,
  "width": 5
}
```

and GETs:

```json
{
  "height": 10,
  "heightInInches": 10,
  "width": 5,
  "widthInInches": 5
}
```

then PUTs back:

```json
{
  "height": 13,
  "heightInInches": 10,
  "width": 5,
  "widthInInches": 5
}
```

As per the compatibility rules, the update must not fail, because it would have
worked before the change.

## Backward compatibility gotchas

* A single feature/property cannot be represented using multiple spec fields
  simultaneously within an API version.  Only one representation can be
  populated at a time, and the client needs to be able to specify which field
  they expect to use (typically via API version), on both mutation and read. As
  above, older clients must continue to function properly.

* A new representation, even in a new API version, that is more expressive than an
  old one breaks backward compatibility, since clients that only understood the
  old representation would not be aware of the new representation nor its
  semantics. Examples of proposals that have run into this challenge include
  [generalized label selectors](http://issues.k8s.io/341) and [pod-level security context](http://prs.k8s.io/12823).

* Enumerated values cause similar challenges. Adding a new value to an enumerated set
  is *not* a compatible change. Clients which assume they know how to handle all possible
  values of a given field will not be able to handle the new values. However, removing a
  value from an enumerated set *can* be a compatible change, if handled properly (treat the
  removed value as deprecated but allowed). For enumeration-like fields that expect to add
  new values in the future, such as `reason` fields, document that expectation clearly
  in the API field description in the first release the field is made available,
  and describe how clients should treat an unknown value. Clients should treat such
  sets of values as potentially open-ended.

* For [Unions](api-conventions.md#unions), sets of fields where at most one should
  be set, it is acceptable to add a new option to the union if the [appropriate
  conventions](api-conventions.md#objects) were followed in the original object.
  Removing an option requires following the [deprecation process](https://kubernetes.io/docs/reference/deprecation-policy/).

* Changing any validation rules always has the potential of breaking some client, since it changes the
  assumptions about part of the API, similar to adding new enum values. Validation rules on spec fields can
  neither be relaxed nor strengthened. Strengthening cannot be permitted because any requests that previously
  worked must continue to work. Weakening validation has the potential to break other consumers and generators
  of the API resource. Status fields whose writers are under our control (e.g., written by non-pluggable
  controllers), may potentially tighten validation, since that would cause a subset of previously valid
  values to be observable by clients.

* Do not add a new API version of an existing resource and make it the preferred version in the same
  release, and do not make it the storage version. The latter is necessary so that a rollback of the
  apiserver doesn't render resources in etcd undecodable after rollback.

* Any field with a default value in one API version must have a *non-nil* default
  value in all API versions.  This can be split into 2 cases:
  * Adding a new API version with a default value for an existing non-defaulted
    field: it is required to add a default value semantically equivalent to
    being unset in all previous API versions, to preserve the semantic meaning
    of the value being unset.
  * Adding a new field with a default value: the default values must be
    semantically equivalent in all currently supported API versions.

## Incompatible API changes

There are times when incompatible changes might be OK, but mostly we want
changes that meet the above definitions. If you think you need to break
compatibility, you should talk to the Kubernetes API reviewers first.

Breaking compatibility of a beta or stable API version, such as v1, is
unacceptable. Compatibility for experimental or alpha APIs is not strictly
required, but breaking compatibility should not be done lightly, as it disrupts
all users of the feature. Alpha and beta API versions may be deprecated and
eventually removed wholesale, as described in the [deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/).

If your change is going to be backward incompatible or might be a breaking
change for API consumers, please send an announcement to
`dev@kubernetes.io` before the change gets in. If you are unsure,
ask. Also make sure that the change gets documented in the release notes for the
next release by labeling the PR with the "release-note-action-required" github label.

If you found that your change accidentally broke clients, it should be reverted.

In short, the expected API evolution is as follows:

* `newapigroup/v1alpha1` -> ... -> `newapigroup/v1alphaN` ->
* `newapigroup/v1beta1` -> ... -> `newapigroup/v1betaN` ->
* `newapigroup/v1` ->
* `newapigroup/v2alpha1` -> ...

While in alpha we expect to move forward with it, but may break it.

Once in beta we will preserve forward compatibility, but may introduce new
versions and delete old ones.

v1 must be backward-compatible for an extended length of time.

## Changing versioned APIs

For most changes, you will probably find it easiest to change the versioned
APIs first. This forces you to think about how to make your change in a
compatible way. Rather than doing each step in every version, it's usually
easier to do each versioned API one at a time, or to do all of one version
before starting "all the rest".

### Edit types.go

The struct definitions for each API are in
`staging/src/k8s.io/api/<group>/<version>/types.go`. Edit those files to reflect
the change you want to make. Note that all types and non-inline fields in
versioned APIs must be preceded by descriptive comments - these are used to
generate documentation. Comments for types should not contain the type name; API
documentation is generated from these comments and end-users should not be
exposed to golang type names.

For types that need the generated
[DeepCopyObject](https://github.com/kubernetes/kubernetes/commit/8dd0989b395b29b872e1f5e06934721863e4a210#diff-6318847735efb6fae447e7dbf198c8b2R3767)
methods, usually only required by the top-level types like `Pod`, add this line
to the comment
([example](https://github.com/kubernetes/kubernetes/commit/39d95b9b065fffebe5b6f233d978fe1723722085#diff-ab819c2e7a94a3521aecf6b477f9b2a7R30)):

```golang
  // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
```

Optional fields should have the `,omitempty` json tag; fields are interpreted as
being required otherwise.

### Edit defaults.go

If your change includes new fields for which you will need default values, you
need to add cases to `pkg/apis/<group>/<version>/defaults.go`.

**Note:** When adding default values to new fields, you *must* also add default
values in all API versions, instead of leaving new fields unset (e.g. `nil`) in
old API versions. This is required because defaulting happens whenever a
serialized version is read (see [#66135]). When possible, pick meaningful values
as sentinels for unset values.

In the past the core v1 API
was special. Its `defaults.go` used to live at `pkg/api/v1/defaults.go`.
If you see code referencing that path, you can be sure its outdated. Now the core v1 api lives at
`pkg/apis/core/v1/defaults.go` which follows the above convention.

Of course, since you have added code, you have to add a test:
`pkg/apis/<group>/<version>/defaults_test.go`.

Do use pointers to scalars when you need to distinguish between an unset value
and an automatic zero value.  For example,
`PodSpec.TerminationGracePeriodSeconds` is defined as `*int64` the go type
definition.  A zero value means 0 seconds, and a nil value asks the system to
pick a default.

Don't forget to run the tests!

[#66135]: https://github.com/kubernetes/kubernetes/issues/66135

### Edit conversion.go

Given that you have not yet changed the internal structs, this might feel
premature, and that's because it is. You don't yet have anything to convert to
or from. We will revisit this in the "internal" section. If you're doing this
all in a different order (i.e. you started with the internal structs), then you
should jump to [that topic below](#edit-version-conversions). In the very rare
case that you are making an incompatible change you might or might not want to
do this now, but you will have to do more later. The files you want are
`pkg/apis/<group>/<version>/conversion.go` and
`pkg/apis/<group>/<version>/conversion_test.go`.

Note that the conversion machinery doesn't generically handle conversion of
values, such as various kinds of field references and API constants. [The client
library](https://github.com/kubernetes/client-go/blob/v4.0.0-beta.0/rest/request.go#L352)
has custom conversion code for field references. You also need to add a call to
`AddFieldLabelConversionFunc` of your scheme with a mapping function that
understands supported translations, like this
[line](https://github.com/kubernetes/kubernetes/blob/v1.8.0-alpha.2/pkg/api/v1/conversion.go#L165).

## Changing the internal structures

Now it is time to change the internal structs so your versioned changes can be
used.

### Edit types.go

Similar to the versioned APIs, the definitions for the internal structs are in
`pkg/apis/<group>/types.go`. Edit those files to reflect the change you want to
make. Keep in mind that the internal structs must be able to express *all* of
the versioned APIs.

Similar to the versioned APIs, you need to add the `+k8s:deepcopy-gen` tag to
types that need generated DeepCopyObject methods.

## Edit validation.go

Most changes made to the internal structs need some form of input validation.
Validation is currently done on internal objects in
`pkg/apis/<group>/validation/validation.go`. This validation is the one of the
first opportunities we have to make a great user experience - good error
messages and thorough validation help ensure that users are giving you what you
expect and, when they don't, that they know why and how to fix it. Think hard
about the contents of `string` fields, the bounds of `int` fields and the
optionality of fields.

Of course, code needs tests - `pkg/apis/<group>/validation/validation_test.go`.

## Edit version conversions

At this point you have both the versioned API changes and the internal
structure changes done.  If there are any notable differences - field names,
types, structural change in particular - you must add some logic to convert
versioned APIs to and from the internal representation.  If you see errors from
the `serialization_test`, it may indicate the need for explicit conversions.

Performance of conversions very heavily influence performance of apiserver.
Thus, we are auto-generating conversion functions that are much more efficient
than the generic ones (which are based on reflections and thus are highly
inefficient).

The conversion code resides with each versioned API. There are two files:

   - `pkg/apis/<group>/<version>/conversion.go` containing manually written
     conversion functions
   - `pkg/apis/<group>/<version>/zz_generated.conversion.go` containing
     auto-generated conversion functions

Since auto-generated conversion functions are using manually written ones,
those manually written should be named with a defined convention, i.e. a
function converting type `X` in pkg `a` to type `Y` in pkg `b`, should be named:
`convert_a_X_To_b_Y`.

**Note:** You should, for efficiency reasons and future updates, use auto-generated
conversion functions when writing your conversion functions.

Adding manually written conversion also requires you to add tests to
`pkg/apis/<group>/<version>/conversion_test.go`.

Once all the necessary manually written conversions are added, you need to
regenerate auto-generated ones. To regenerate them run:

```sh
make clean && make generated_files
```

`make clean` is important, otherwise the generated files might be stale, because
the build system uses custom cache.

`make all` will invoke `make generated_files` as well.

The `make generated_files` will also regenerate the `zz_generated.deepcopy.go`,
`zz_generated.defaults.go`, and `api/openapi-spec/swagger.json`.

If regeneration is somehow not possible due to compile errors, the easiest
workaround is to remove the files causing errors and rerun the command.

## Generate Code

Apart from the `defaulter-gen`, `deepcopy-gen`, `conversion-gen` and
`openapi-gen`, there are a few other generators:
 - `go-to-protobuf`
 - `client-gen`
 - `lister-gen`
 - `informer-gen`
 - `codecgen` (for fast json serialization with ugorji codec)

Many of the generators are based on
[`gengo`](https://github.com/kubernetes/gengo) and share common
flags. The `--verify-only` flag will check the existing files on disk
and fail if they are not what would have been generated.

The generators that create go code have a `--go-header-file` flag
which should be a file that contains the header that should be
included. This header is the copyright that should be present at the
top of the generated file and should be checked with the
[`repo-infra/verify/verify-boilerplane.sh`](https://git.k8s.io/repo-infra/verify/verify-boilerplate.sh)
script at a later stage of the build.

To invoke these generators, you can run `make update`, which runs a bunch of
[scripts](https://github.com/kubernetes/kubernetes/blob/release-1.23/hack/make-rules/update.sh#L47-L55).
Please continue to read the next a few sections, because some generators have
prerequisites, also because they introduce how to invoke the generators
individually if you find `make update` takes too long to run.

### Generate protobuf objects

For any core API object, we also need to generate the Protobuf IDL and marshallers.
That generation is invoked with

```sh
hack/update-generated-protobuf.sh
```

The vast majority of objects will not need any consideration when converting
to protobuf, but be aware that if you depend on a Golang type in the standard
library there may be additional work required, although in practice we typically
use our own equivalents for JSON serialization. The `pkg/api/serialization_test.go`
will verify that your protobuf serialization preserves all fields - be sure to
run it several times to ensure there are no incompletely calculated fields.

### Generate Clientset

`client-gen` is a tool to generate clientsets for top-level API objects.

`client-gen` requires the `// +genclient` annotation on each
exported type in both the internal `pkg/apis/<group>/types.go` as well as each
specifically versioned `staging/src/k8s.io/api/<group>/<version>/types.go`.

If the apiserver hosts your API under a different group name than the `<group>`
in the filesystem, (usually this is because the `<group>` in the filesystem
omits the "k8s.io" suffix, e.g., admission vs. admission.k8s.io), you can
instruct the `client-gen` to use the correct group name by adding the `//
+groupName=` annotation in the `doc.go` in both the internal
`pkg/apis/<group>/doc.go` as well as in each specifically versioned
`staging/src/k8s.io/api/<group>/<version>/types.go`.

Once you added the annotations, generate the client with

```sh
hack/update-codegen.sh
```

Note that you can use the optional `// +groupGoName=` to specify a CamelCase
custom Golang identifier to de-conflict e.g. `policy.authorization.k8s.io` and
`policy.k8s.io`. These two would both map to `Policy()` in clientsets.

client-gen is flexible. See [this document](../sig-api-machinery/generating-clientset.md) if you need
client-gen for non-kubernetes API.

### Generate Listers

`lister-gen` is a tool to generate listers for a client. It reuses the
`//+genclient` and the `// +groupName=` annotations, so you do not need to
specify extra annotations.

Your previous run of `hack/update-codegen.sh` has invoked `lister-gen`.

### Generate Informers

`informer-gen` generates the very useful Informers which watch API
resources for changes. It reuses the `//+genclient` and the
`//+groupName=` annotations, so you do not need to specify extra annotations.

Your previous run of `hack/update-codegen.sh` has invoked `informer-gen`.

### Edit json (un)marshaling code

We are auto-generating code for marshaling and unmarshaling json representation
of api objects - this is to improve the overall system performance.

The auto-generated code resides with each versioned API:

   - `staging/src/k8s.io/api/<group>/<version>/generated.proto`
   - `staging/src/k8s.io/api/<group>/<version>/generated.pb.go`

To regenerate them run:

```sh
hack/update-generated-protobuf.sh
```

## Making a new API Version

This section is under construction, as we make the tooling completely generic.

If you are adding a new API version to an existing group, you can copy the
structure of the existing `pkg/apis/<group>/<existing-version>` and
`staging/src/k8s.io/api/<group>/<existing-version>` directories.

It is helpful to structure the PR in layered commits to make it easier for
reviewers to see what has changed between the two versions:
1. A commit that just copies the `pkg/apis/<group>/<existing-version>` and
   `staging/src/k8s.io/api/<group>/<existing-version>` packages to the
   `<new-version>`.
1. A commit that renames `<existing-version>`to `<new-version>` in the new files.
1. A commit that makes any new changes for `<new-version>`.
1. A commit that contains the generated files from running `make generated_files`, `make update`, etc.

Due to the fast changing nature of the project, the following content is probably out-dated:
* You must add the version to
  [pkg/controlplane/instance.go](https://github.com/kubernetes/kubernetes/blob/v1.21.2/pkg/controlplane/instance.go#L662)
  is be enabled by default for stable versions, or disabled by default
  for alpha and beta versions.
* You must add the new version to
  `pkg/apis/group_name/install/install.go` (for example, [pkg/apis/apps/install/install.go](https://github.com/kubernetes/kubernetes/blob/v1.21.2/pkg/apis/apps/install/install.go)).
* You must add the new version to
  [hack/lib/init.sh#KUBE_AVAILABLE_GROUP_VERSIONS](https://github.com/kubernetes/kubernetes/blob/v1.21.2/hack/lib/init.sh#L65).
* You must add the new version  to
  [cmd/kube-apiserver/app#apiVersionPriorities](https://github.com/kubernetes/kubernetes/blob/v1.21.2/cmd/kube-apiserver/app/aggregator.go#L247).
* You must setup storage for the new version in
  `pkg/registry/group_name/rest` (for example, [pkg/registry/authentication/rest](https://github.com/kubernetes/kubernetes/blob/v1.21.2/pkg/registry/authentication/rest/storage_authentication.go)).
* For `kubectl get` you must add a table definition to [pkg/printers/internalversion/printers.go](https://github.com/kubernetes/kubernetes/blob/v1.23.0/pkg/printers/internalversion/printers.go). Integration tests for this are in [test/integration/apiserver/print_test.go](https://github.com/kubernetes/kubernetes/blob/v1.23.0/test/integration/apiserver/print_test.go).

You need to regenerate the generated code as instructed in the sections above.

### Testing

Some updates to tests are required.

* You must add the new storage version hash published in API discovery data to
  [pkg/controlplane/storageversionhashdata/datago#GVRToStorageVersionHash](https://github.com/kubernetes/kubernetes/blob/v1.21.2/pkg/controlplane/storageversionhashdata/data.go#L44).
    * Run `go test ./pkg/controlplane -run StorageVersion` to verify.
* You must add the new version stub to the persisted versions stored in etcd in [test/integration/etcd/data.go](https://github.com/kubernetes/kubernetes/blob/v1.21.2/test/integration/etcd/data.go#L40).
    * Run `go test ./test/integration/etcd` to verify
* Sanity test the changes by bringing up a cluster (i.e.,
local-up-cluster.sh, kind, etc) and running `kubectl get
<resource>.<version>.<group>`.
* [Integration tests](../sig-testing/integration-tests.md)
are also good for testing the full CRUD lifecycle along with the controller.
  * To write integration tests for beta APIs you will need to selectively enable the resources you need.
    You can do this using [cmd/kube-apiserver/app/testing/testserver.go#StartTestServerOrDie](https://github.com/kubernetes/kubernetes/blob/2b1b849d6a8bdeb7dc0807438cfd0ff2a9d752c1/cmd/kube-apiserver/app/testing/testserver.go#L325).
    You will then pass the `--runtime-config=groupname/v1beta1/resourcename` as a flag to enable the beta API.
* For beta APIs, e2e tests need to perform discovery checks against the kube-apiserver to determine if
  a beta API is enabled or not.  See [test/e2e/apimachinery/discovery.go](https://github.com/kubernetes/kubernetes/blob/2b1b849d6a8bdeb7dc0807438cfd0ff2a9d752c1/test/e2e/apimachinery/discovery.go#L50)
  for an example.
  There is a [prow dashboard for beta API jobs](https://prow.k8s.io/?job=*betaapis*) to watch your results.

## Making a new API Group

You'll have to make a new directory under `pkg/apis/` and
`staging/src/k8s.io/api`; copy the directory structure of an existing API group,
e.g. `pkg/apis/authentication` and `staging/src/k8s.io/api/authentication`;
replace "authentication" with your group name and replace versions with your
versions; replace the API kinds in
[versioned](https://github.com/kubernetes/kubernetes/blob/v1.8.0-alpha.2/staging/src/k8s.io/api/authentication/v1/register.go#L47)
and
[internal](https://github.com/kubernetes/kubernetes/blob/v1.8.0-alpha.2/pkg/apis/authentication/register.go#L47)
register.go, and
[install.go](https://github.com/kubernetes/kubernetes/blob/v1.8.0-alpha.2/pkg/apis/authentication/install/install.go#L43)
with your kinds.

You'll have to add your API group/version to a few places in the code base, as
noted in [Making a new API Version](#making-a-new-api-version) section.

You need to regenerate the generated code as instructed in the sections above.

## Update the fuzzer

Part of our testing regimen for APIs is to "fuzz" (fill with random values) API
objects and then convert them to and from the different API versions. This is
a great way of exposing places where you lost information or made bad
assumptions.

The fuzzer works by creating a random API object and calling the custom fuzzer
function in `pkg/apis/$GROUP/fuzzer/fuzzer.go`. The resulting object is then
round-tripped from one api version to another, and verified to be the same as
what was started with. Validation is not run during this process, but defaulting
is.

If you have added any fields which need very careful formatting (the test does
not run validation) or if you have made assumptions during defaulting such as
"this slice will always have at least 1 element", you may get an error or even a
panic from the `k8s.io/kubernetes/pkg/api/testing.TestRoundTripTypes` in
`./pkg/api/testing/serialization_test.go`.

If you default any fields, you must check that in the custom fuzzer function,
because the fuzzer may leave some fields empty. If your object has a structure
reference, the fuzzer may leave that nil, or it may create a random object. Your
custom fuzzer function must ensure that defaulting does not further change the
object, as that will show up as a diff in the round trip test.

Finally, the fuzz test runs without any feature gate configuration. If
defaulting or other behavior is behind a feature gate, beware that the fuzz
behavior will change when the feature gate becomes default on.

## Update the semantic comparisons

VERY VERY rarely is this needed, but when it hits, it hurts. In some rare cases
we end up with objects (e.g. resource quantities) that have morally equivalent
values with different bitwise representations (e.g. value 10 with a base-2
formatter is the same as value 0 with a base-10 formatter). The only way Go
knows how to do deep-equality is through field-by-field bitwise comparisons.
This is a problem for us.

The first thing you should do is try not to do that. If you really can't avoid
this, I'd like to introduce you to our `apiequality.Semantic.DeepEqual` routine.
It supports custom overrides for specific types - you can find that in
`pkg/api/helper/helpers.go`.

There's one other time when you might have to touch this: `unexported fields`.
You see, while Go's `reflect` package is allowed to touch `unexported fields`,
us mere mortals are not - this includes `apiequality.Semantic.DeepEqual`.
Fortunately, most of our API objects are "dumb structs" all the way down - all
fields are exported (start with a capital letter) and there are no unexported
fields. But sometimes you want to include an object in our API that does have
unexported fields somewhere in it (for example, `time.Time` has unexported fields).
If this hits you, you may have to touch the `apiequality.Semantic.DeepEqual`
customization functions.

## Implement your change

Now you have the API all changed - go implement whatever it is that you're
doing!

## Write end-to-end tests

Check out the [E2E docs](../sig-testing/e2e-tests.md) for detailed information about how to
write end-to-end tests for your feature.
Make sure the E2E tests are running in the default presubmits for a feature/API that
is enabled by default.

## Examples and docs

At last, your change is done, all unit tests pass, e2e passes, you're done,
right? Actually, no. You just changed the API. If you are touching an existing
facet of the API, you have to try *really* hard to make sure that *all* the
examples and docs are updated. There's no easy way to do this, due in part to
JSON and YAML silently dropping unknown fields. You're clever - you'll figure it
out. Put `grep` or `ack` to good use.

If you added functionality, you should consider documenting it and/or writing
an example to illustrate your change.

Make sure you update the swagger and OpenAPI spec by running:

```sh
make update
```

The API spec changes should be in a commit separate from your other changes.

## Alpha, Beta, and Stable Versions

New feature development proceeds through a series of stages of increasing
maturity:

- Development level
  - Object Versioning: no convention
  - Availability: not committed to main kubernetes repo, and thus not available
in official releases
  - Audience: other developers closely collaborating on a feature or
proof-of-concept
  - Upgradeability, Reliability, Completeness, and Support: no requirements or
guarantees
- Alpha level
  - Object Versioning: API version name contains `alpha` (e.g. `v1alpha1`)
  - Availability: committed to main kubernetes repo;  appears in an official
release; feature is disabled by default, but may be enabled by flag
  - Audience: developers and expert users interested in giving early feedback on
features
  - Completeness: some API operations, CLI commands, or UI support may not be
implemented; the API need not have had an *API review* (an intensive and
targeted review of the API, on top of a normal code review)
  - Upgradeability: the object schema and semantics may change in a later
software release, without any provision for preserving objects in an existing
cluster; removing the upgradability concern allows developers to make rapid
progress; in particular, API versions can increment faster than the minor
release cadence and the developer need not maintain multiple versions;
developers should still increment the API version when object schema or
semantics change in an [incompatible way](#on-compatibility)
  - Cluster Reliability: because the feature is relatively new, and may lack
complete end-to-end tests, enabling the feature via a flag might expose bugs
with destabilize the cluster (e.g. a bug in a control loop might rapidly create
excessive numbers of object, exhausting API storage).
  - Support: there is *no commitment* from the project to complete the feature;
the feature may be dropped entirely in a later software release
  - Recommended Use Cases: only in short-lived testing clusters, due to
complexity of upgradeability and lack of long-term support and lack of
upgradability.
- Beta level:
  - Object Versioning: API version name contains `beta` (e.g. `v2beta3`)
  - Availability: in official Kubernetes releases; API is disabled by default
but may be enabled by a flag.
(Note: beta APIs introduced before v1.24 were enabled by default, but this
[changed for new beta APIs](https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/3136-beta-apis-off-by-default/README.md))
  - Audience: users interested in providing feedback on features
  - Completeness: all API operations, CLI commands, and UI support should be
implemented; end-to-end tests complete; the API has had a thorough API review
and is thought to be complete, though use during beta may frequently turn up API
issues not thought of during review
  - Upgradeability: the object schema and semantics may change in a later
software release; when this happens, an upgrade path will be documented; in some
cases, objects will be automatically converted to the new version; in other
cases, a manual upgrade may be necessary; a manual upgrade may require downtime
for anything relying on the new feature, and may require manual conversion of
objects to the new version; when manual conversion is necessary, the project
will provide documentation on the process
  - Cluster Reliability: since the feature has e2e tests, enabling the feature
via a flag should not create new bugs in unrelated features; because the feature
is new, it may have minor bugs
  - Support: the project commits to complete the feature, in some form, in a
subsequent Stable version; typically this will happen within 3 months, but
sometimes longer; releases should simultaneously support two consecutive
versions (e.g. `v1beta1` and `v1beta2`; or `v1beta2` and `v1`) for at least one
minor release cycle (typically 3 months) so that users have enough time to
upgrade and migrate objects
  - Recommended Use Cases: in short-lived testing clusters; in production
clusters as part of a short-lived evaluation of the feature in order to provide
feedback
- Stable level:
  - Object Versioning: API version `vX` where `X` is an integer (e.g. `v1`)
  - Availability: in official Kubernetes releases, and enabled by default
  - Audience: all users
  - Completeness: must have conformance tests, approved by SIG Architecture,
in the appropriate conformance profile (e.g., non-portable and/or optional
features may not be in the default profile)
  - Upgradeability: only [strictly compatible](#on-compatibility) changes
allowed in subsequent software releases
  - Cluster Reliability: high
  - Support: API version will continue to be present for many subsequent
software releases;
  - Recommended Use Cases: any

### Adding Unstable Features to Stable Versions

When adding a feature to an object which is already Stable, the new fields and
new behaviors need to meet the Stable level requirements. If these cannot be
met, then the new field cannot be added to the object.

For example, consider the following object:

```go
// API v6.
type Frobber struct {
  // height ...
  Height *int32 `json:"height"
  // param ...
  Param  string `json:"param"
}
```

A developer is considering adding a new `Width` parameter, like this:

```go
// API v6.
type Frobber struct {
  // height ...
  Height *int32 `json:"height"
  // param ...
  Param  string `json:"param"
  // width ...
  Width  *int32 `json:"width,omitempty"
}
```

However, the new feature is not stable enough to be used in a stable version
(`v6`). Some reasons for this might include:

- the final representation is undecided (e.g. should it be called `Width` or `Breadth`?)
- the implementation is not stable enough for general use (e.g. the `Area()` routine sometimes overflows.)

The developer cannot add the new field unconditionally until stability is met. However,
sometimes stability cannot be met until some users try the new feature, and some
users are only able or willing to accept a released version of Kubernetes. In
that case, the developer has a few options, both of which require staging work
over several releases.

The mechanism used depends on whether a new field is being added,
or a new value is being permitted in an existing field.

#### New field in existing API version

Previously, annotations were used for experimental alpha features, but are no longer recommended for several reasons:

* They expose the cluster to "time-bomb" data added as unstructured annotations against an earlier API server (https://issue.k8s.io/30819)
* They cannot be migrated to first-class fields in the same API version (see the issues with representing a single value in multiple places in [backward compatibility gotchas](#backward-compatibility-gotchas))

The preferred approach adds an alpha field to the existing object, and ensures it is disabled by default:

1. Add a [feature gate](feature-gates.md) to the API server to control enablement of the new field:

    In [staging/src/k8s.io/apiserver/pkg/features/kube_features.go](https://git.k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/features/kube_features.go):

    ```go
    // owner: @you
    // alpha: v1.11
    //
    // Add multiple dimensions to frobbers.
    Frobber2D utilfeature.Feature = "Frobber2D"

    var defaultKubernetesFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{
      ...
      Frobber2D: {Default: false, PreRelease: utilfeature.Alpha},
    }
    ```

2. Add the field to the API type:

    * ensure the field is [optional](api-conventions.md#optional-vs-required)
        * add the `omitempty` struct tag
        * add the `// +optional` comment tag
        * add the `// +featureGate=<gate-name>` comment tag
        * ensure the field is entirely absent from API responses when empty (optional fields must be pointers)
    * include details about the alpha-level in the field description

    ```go
    // API v6.
    type Frobber struct {
      // height ...
      Height int32  `json:"height"`
      // param ...
      Param  string `json:"param"`
      // width indicates how wide the object is.
      // This field is alpha-level and is only honored by servers that enable the Frobber2D feature.
      // +optional
      // +featureGate=Frobber2D
      Width  *int32 `json:"width,omitempty"`
    }
    ```

3. Before persisting the object to storage, clear disabled alpha fields on create,
and on update if the existing object does not already have a value in the field.
This prevents new usage of the feature while it is disabled, while ensuring existing data is preserved.
Ensuring existing data is preserved is needed so that when the feature is enabled by default in a future version *n*
and data is unconditionally allowed to be persisted in the field, an *n-1* API server
(with the feature still disabled by default) will not drop the data on update.
The recommended place to do this is in the REST storage strategy's PrepareForCreate/PrepareForUpdate methods:

    ```go
    func (frobberStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
      frobber := obj.(*api.Frobber)

      if !utilfeature.DefaultFeatureGate.Enabled(features.Frobber2D) {
        frobber.Width = nil
      }
    }

    func (frobberStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
      newFrobber := obj.(*api.Frobber)
      oldFrobber := old.(*api.Frobber)

      if !utilfeature.DefaultFeatureGate.Enabled(features.Frobber2D) && oldFrobber.Width == nil {
        newFrobber.Width = nil
      }
    }
    ```

4. To future-proof your API testing, when testing with feature gate on and off, ensure that the gate is deliberately set as desired. Don't assume that gate is off or on. As your feature
progresses from `alpha` to `beta` and then `stable` the feature might be turned on or off by default across the entire code base. The below example
provides some details

   ```go
   func TestAPI(t *testing.T){
    testCases:= []struct{
      // ... test definition ...
    }{
       {
        // .. test case ..
       },
       {
       // ... test case ..
       },
   }

   for _, testCase := range testCases{
     t.Run("..name...", func(t *testing.T){
      // run with gate on
      defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features. Frobber2D, true)()
       // ... test logic ...
     })
     t.Run("..name...", func(t *testing.T){
      // run with gate off, *do not assume it is off by default*
      defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features. Frobber2D, false)()
      // ... test gate-off testing logic logic ...
     })
   }
   ```

5. In validation, validate the field if present:

    ```go
    func ValidateFrobber(f *api.Frobber, fldPath *field.Path) field.ErrorList {
      ...
      if f.Width != nil {
        ... validation of width field ...
      }
      ...
    }
    ```

In future Kubernetes versions:

* if the feature progresses to beta or stable status, the feature gate can be removed or be enabled by default.
* if the schema of the alpha field must change in an incompatible way, a new field name must be used.
* if the feature is abandoned, or the field name is changed, the field should be removed from the go struct, with a tombstone comment ensuring the field name and protobuf tag are not reused:

    ```go
    // API v6.
    type Frobber struct {
      // height ...
      Height int32  `json:"height" protobuf:"varint,1,opt,name=height"`
      // param ...
      Param  string `json:"param" protobuf:"bytes,2,opt,name=param"`

      // +k8s:deprecated=width,protobuf=3
    }
    ```

#### Ratcheting validation

The word "ratcheting" refers to a process of incremental and often irreversible
progression or change, typically in a single direction. The term originates from
a mechanical device called a [ratchet](https://en.wikipedia.org/wiki/Ratchet_(device)),
which consists of a toothed wheel or bar and a pawl (a catch) that allows movement
in only one direction while preventing backward motion.

In the Kubernetes world, a ratcheting validation refers to an incremental tightening
of validation. This means we allow current resources to either remain invalid or
be fixed, but all new resources must pass the validation. The following table
best illustrates these cases:

| Resource    | Validation |
|-------------|------------|
| new valid   | succeeds   |
| new invalid | fails      |
| old valid   | succeeds   |
| old invalid | succeeds   |

A good example of ratcheting validation was introduced in [this pull request](https://github.com/kubernetes/kubernetes/pull/130233).
It introduced validation for the optional `.spec.serviceName` field for StatefulSet,
such that old resources (nregarldess of whether they are valid or not) will succeed
the validation check, but new resources must adhere to stricter validation rules
for that field. The relevant changes include:
- A struct with options passed to validation methods (here it's the `StatefulSetValidationOptions`
  struct, with `AllowInvalidServiceName` to handle this specific case).
- Appropriate changes inside `Validate*` methods which ensure the rules from the
  table above are implemented.
- Tests ensuring all the cases from the above table are covered.

#### New enum value in existing field

A developer is considering adding a new allowed enum value of `"OnlyOnTuesday"`
to the following existing enum field:

```go
type Frobber struct {
  // restartPolicy may be set to "Always" or "Never".
  // Additional policies may be defined in the future.
  // Clients should expect to handle additional values,
  // and treat unrecognized values in this field as "Never".
  RestartPolicy string `json:"policy"
}
```

Older versions of expected API clients must be able handle the new value in a safe way:

* If the enum field drives behavior of a single component, ensure all versions of that component
  that will encounter API objects containing the new value handle it properly or fail safe.
  For example, a new allowed value in a `Pod` enum field consumed by the kubelet must be handled
  safely by kubelets up to three versions older than the first API server release that allowed the new value.
* If an API drives behavior that is implemented by external clients (like `Ingress` or `NetworkPolicy`),
  the enum field must explicitly indicate that additional values may be allowed in the future,
  and define how unrecognized values must be handled by clients. If this was not done in the first release
  containing the enum field, it is not safe to add new values that can break existing clients.

If expected API clients safely handle the new enum value, the next requirement is to begin allowing it
in a way that does not break validation of that object by a previous API server.
This requires at least two releases to accomplish safely:

Release 1:

* Only allow the new enum value when updating existing objects that already contain the new enum value
* Disallow it in other cases (creation, and update of objects that do not already contain the new enum value)
* Verify that known clients handle the new value as expected, honoring the new value or using previously defined "unknown value" behavior,
  (depending on whether the associated feature gate is enabled or not)


Release 2:

* Allow the new enum value in create and update scenarios

This ensures a cluster with multiple servers at skewed releases (which happens during a rolling upgrade),
will not allow data to be persisted which the previous release of the API server would choke on.

Typically, a [feature gate](feature-gates.md) is used to do this rollout, starting in alpha and disabled by default in release 1,
and graduating to beta and enabled by default in release 2.

1. Add a feature gate to the API server to control enablement of the new enum value (and associated function):

    In [staging/src/k8s.io/apiserver/pkg/features/kube_features.go](https://git.k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/features/kube_features.go):

    ```go
    // owner: @you
    // alpha: v1.11
    //
    // Allow OnTuesday restart policy in frobbers.
    FrobberRestartPolicyOnTuesday utilfeature.Feature = "FrobberRestartPolicyOnTuesday"

    var defaultKubernetesFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{
      ...
      FrobberRestartPolicyOnTuesday: {Default: false, PreRelease: utilfeature.Alpha},
    }
    ```

2. Update the documentation on the API type:

    * include details about the alpha-level in the field description

    ```go
    type Frobber struct {
      // restartPolicy may be set to "Always" or "Never" (or "OnTuesday" if the alpha "FrobberRestartPolicyOnTuesday" feature is enabled).
      // Additional policies may be defined in the future.
      // Unrecognized policies should be treated as "Never".
      RestartPolicy string `json:"policy"
    }
    ```

3. When validating the object, determine whether the new enum value should be allowed.
This prevents new usage of the new value when the feature is disabled, while ensuring existing data is preserved.
Ensuring existing data is preserved is needed so that when the feature is enabled by default in a future version *n*
and data is unconditionally allowed to be persisted in the field, an *n-1* API server
(with the feature still disabled by default) will not choke on validation.
The recommended place to do this is in the REST storage strategy's Validate/ValidateUpdate methods:

    ```go
    func (frobberStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
      frobber := obj.(*api.Frobber)
      return validation.ValidateFrobber(frobber, validationOptionsForFrobber(frobber, nil))
    }

    func (frobberStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
      newFrobber := obj.(*api.Frobber)
      oldFrobber := old.(*api.Frobber)
      return validation.ValidateFrobberUpdate(newFrobber, oldFrobber, validationOptionsForFrobber(newFrobber, oldFrobber))
    }

    func validationOptionsForFrobber(newFrobber, oldFrobber *api.Frobber) validation.FrobberValidationOptions {
      opts := validation.FrobberValidationOptions{
        // allow if the feature is enabled
        AllowRestartPolicyOnTuesday: utilfeature.DefaultFeatureGate.Enabled(features.FrobberRestartPolicyOnTuesday)
      }

      if oldFrobber == nil {
        // if there's no old object, use the options based solely on feature enablement
        return opts
      }

      if oldFrobber.RestartPolicy == api.RestartPolicyOnTuesday {
        // if the old object already used the enum value, continue to allow it in the new object
        opts.AllowRestartPolicyOnTuesday = true
      }
      return opts
    }
    ```

4. In validation, validate the enum value based on the passed-in options:

    ```go
    func ValidateFrobber(f *api.Frobber, opts FrobberValidationOptions) field.ErrorList {
      ...
      validRestartPolicies := sets.NewString(RestartPolicyAlways, RestartPolicyNever)
      if opts.AllowRestartPolicyOnTuesday {
        validRestartPolicies.Insert(RestartPolicyOnTuesday)
      }

      if f.RestartPolicy == RestartPolicyOnTuesday && !opts.AllowRestartPolicyOnTuesday {
        allErrs = append(allErrs, field.Invalid(field.NewPath("restartPolicy"), f.RestartPolicy, "only allowed if the FrobberRestartPolicyOnTuesday feature is enabled"))
      } else if !validRestartPolicies.Has(f.RestartPolicy) {
        allErrs = append(allErrs, field.NotSupported(field.NewPath("restartPolicy"), f.RestartPolicy, validRestartPolicies.List()))
      }
      ...
    }
    ```

5. After at least one release, the feature can be promoted to beta or GA and enabled by default.

    In [staging/src/k8s.io/apiserver/pkg/features/kube_features.go](https://git.k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/features/kube_features.go):

    ```go
    // owner: @you
    // alpha: v1.11
    // beta: v1.12
    //
    // Allow OnTuesday restart policy in frobbers.
    FrobberRestartPolicyOnTuesday utilfeature.Feature = "FrobberRestartPolicyOnTuesday"

    var defaultKubernetesFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{
      ...
      FrobberRestartPolicyOnTuesday: {Default: true, PreRelease: utilfeature.Beta},
    }
    ```

#### New alpha API version

Another option is to introduce a new type with an new `alpha` or `beta` version
designator, like this:

```go
// API v7alpha1
type Frobber struct {
  // height ...
  Height *int32 `json:"height"`
  // param ...
  Param  string `json:"param"`
  // width ...
  Width  *int32 `json:"width,omitempty"`
}
```

The latter requires that all objects in the same API group as `Frobber` to be
replicated in the new version, `v7alpha1`. This also requires user to use a new
client which uses the other version. Therefore, this is not a preferred option.

A related issue is how a cluster manager can roll back from a new version
with a new feature, that is already being used by users. See
https://github.com/kubernetes/kubernetes/issues/4855.
