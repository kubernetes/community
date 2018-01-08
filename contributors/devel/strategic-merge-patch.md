Strategic Merge Patch
=====================

# Background

Kubernetes supports a customized version of JSON merge patch called strategic merge patch.  This
patch format is used by `kubectl apply`, `kubectl edit` and `kubectl patch`, and contains
specialized directives to control how specific fields are merged.

In the standard JSON merge patch, JSON objects are always merged but lists are
always replaced. Often that isn't what we want. Let's say we start with the
following Pod:

```yaml
spec:
  containers:
    - name: nginx
      image: nginx-1.0
```

and we POST that to the server (as JSON). Then let's say we want to *add* a
container to this Pod.

```yaml
PATCH /api/v1/namespaces/default/pods/pod-name
spec:
  containers:
    - name: log-tailer
      image: log-tailer-1.0
```

If we were to use standard Merge Patch, the entire container list would be
replaced with the single log-tailer container. However, our intent is for the
container lists to merge together based on the `name` field.

To solve this problem, Strategic Merge Patch uses the go struct tag of the API
objects to determine what lists should be merged and which ones should not.
The metadata is available as struct tags on the API objects
themselves and also available to clients as [OpenAPI annotations](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/README.md#x-kubernetes-patch-strategy-and-x-kubernetes-patch-merge-key).
In the above example, the `patchStrategy` metadata for the `containers`
field would be `merge` and the `patchMergeKey` would be `name`.


# Basic Patch Format

Strategic Merge Patch supports special operations through directives.

There are multiple directives:

- replace
- merge
- delete
- delete from primitive list

`replace`, `merge` and `delete` are mutual exclusive.

## `replace` Directive

### Purpose

`replace` directive indicates that the element that contains it should be replaced instead of being merged.

### Syntax

`replace` directive is used in both patch with directive marker and go struct tags.

Example usage in the patch:

```
$patch: replace
```

### Example

`replace` directive can be used on both map and list.

#### Map

To indicate that a map should not be merged and instead should be taken literally:

```yaml
$patch: replace  # recursive and applies to all fields of the map it's in
containers:
- name: nginx
  image: nginx-1.0
```

#### List of Maps

To override the container list to be strictly replaced, regardless of the default:

```yaml
containers:
  - name: nginx
    image: nginx-1.0
  - $patch: replace   # any further $patch operations nested in this list will be ignored
```


## `delete` Directive

### Purpose

`delete` directive indicates that the element that contains it should be deleted.

### Syntax

`delete` directive is used only in the patch with directive marker.
It can be used on both map and list of maps.
```
$patch: delete
```

### Example

#### List of Maps

To delete an element of a list that should be merged:

```yaml
containers:
  - name: nginx
    image: nginx-1.0
  - $patch: delete
    name: log-tailer  # merge key and value goes here
```

Note: Delete operation will delete all entries in the list that match the merge key.

#### Maps

One way to delete a map is using `delete` directive.
Applying this patch will delete the rollingUpdate map.
```yaml
rollingUpdate:
  $patch: delete
```

An equivalent way to delete this map is
```yaml
rollingUpdate: null
```

## `merge` Directive

### Purpose

`merge` directive indicates that the element that contains it should be merged instead of being replaced.

### Syntax

`merge` directive is used only in the go struct tags.


## `deleteFromPrimitiveList` Directive

### Purpose

We have two patch strategies for lists of primitives: replace and merge.
Replace is the default patch strategy for list, which will replace the whole list on update and it will preserve the order;
while merge strategy works as an unordered set. We call a primitive list with merge strategy an unordered set.
The patch strategy is defined in the go struct tag of the API objects.

`deleteFromPrimitiveList` directive indicates that the elements in this list should be deleted from the original primitive list.

### Syntax

It is used only as the prefix of the key in the patch.
```
$deleteFromPrimitiveList/<keyOfPrimitiveList>: [a primitive list]
```

### Example

##### List of Primitives (Unordered Set)

`finalizers` uses `merge` as patch strategy.
```go
Finalizers []string `json:"finalizers,omitempty" patchStrategy:"merge" protobuf:"bytes,14,rep,name=finalizers"`
```

Suppose we have defined a `finalizers` and we call it the original finalizers:

```yaml
finalizers:
  - a
  - b
  - c
```

To delete items "b" and "c" from the original finalizers, the patch will be:

```yaml
# The directive includes the prefix $deleteFromPrimitiveList and
# followed by a '/' and the name of the list.
# The values in this list will be deleted after applying the patch.
$deleteFromPrimitiveList/finalizers:
  - b
  - c
```

After applying the patch on the original finalizers, it will become:

```yaml
finalizers:
  - a
```

Note: When merging two set, the primitives are first deduplicated and then merged.
In an erroneous case, the set may be created with duplicates. Deleting an
item that has duplicates will delete all matching items.

## `setElementOrder` Directive

### Purpose

`setElementOrder` directive provides a way to specify the order of a list.
The relative order specified in this directive will be retained.
Please refer to [proposal](/contributors/design-proposals/cli/preserve-order-in-strategic-merge-patch.md) for more information.

### Syntax

It is used only as the prefix of the key in the patch.
```
$setElementOrder/<keyOfList>: [a list]
```

### Example

#### List of Primitives

Suppose we have a list of `finalizers`:
```yaml
finalizers:
  - a
  - b
  - c
```

To reorder the elements order in the list, we can send a patch:
```yaml
# The directive includes the prefix $setElementOrder and
# followed by a '/' and the name of the list.
$setElementOrder/finalizers:
  - b
  - c
  - a
```

After applying the patch, it will be:
```yaml
finalizers:
  - b
  - c
  - a
```

#### List of Maps

Suppose we have a list of `containers` whose `mergeKey` is `name`:
```yaml
containers:
  - name: a
    ...
  - name: b
    ...
  - name: c
    ...
```

To reorder the elements order in the list, we can send a patch:
```yaml
# each map in the list should only include the mergeKey
$setElementOrder/containers:
  - name: b
  - name: c
  - name: a
```

After applying the patch, it will be:
```yaml
containers:
  - name: b
    ...
  - name: c
    ...
  - name: a
    ...
```


## `retainKeys` Directive

### Purpose

`retainKeys` directive provides a mechanism for union types to clear mutual exclusive fields.
When this directive is present in the patch, all the fields not in this directive will be cleared.
Please refer to [proposal](/contributors/design-proposals/api-machinery/add-new-patchStrategy-to-clear-fields-not-present-in-patch.md) for more information.

### Syntax

```
$retainKeys: [a list of field keys]
```

### Example

#### Map

Suppose we have a union type:
```
union:
  foo: a
  other: b
```

And we have a patch:
```
union:
  retainKeys:
    - another
    - bar
  another: d
  bar: c
```

After applying this patch, we get:
```
union:
  # Field foo and other have been cleared w/o explicitly set them to null.
  another: d
  bar: c
```

# Changing patch format

As issues and limitations have been discovered with the strategic merge
patch implementation, it has been necessary to change the patch format
to support additional semantics - such as merging lists of
primitives and defining order when merging lists.

## Requirements for any changes to the patch format

**Note:** Changes to the strategic merge patch must be backwards compatible such
that patch requests valid in previous versions continue to be valid.
That is, old patch formats sent by old clients to new servers with
must continue to function correctly.

Previously valid patch requests do not need to keep the exact same
behavior, but do need to behave correctly.

**Example:** if a patch request previously randomized the order of elements
in a list and we want to provide a deterministic order, we must continue
to support old patch format but we can make the ordering deterministic
for the old format.

### Client version skew

Because the server does not publish which patch versions it supports,
and it silently ignores patch directives that it does not recognize,
new patches should behave correctly when sent to old servers that
may not support all of the patch directives.

While the patch API must be backwards compatible, it must also
be forward compatible for 1 version.  This is needed because `kubectl` must
support talking to older and newer server versions without knowing what
parts of patch are supported on each, and generate patches that work correctly on both.

## Strategies for introducing new patch behavior

#### 1. Add optional semantic meaning to the existing patch format.

**Note:** Must not require new data or elements to be present that was not required before.  Meaning must not break old interpretation of old patches.

**Good Example:**

Old format
  - ordering of elements in patch had no meaning and the final ordering was arbitrary

New format
  - ordering of elements in patch has meaning and the final ordering is deterministic based on the ordering in the patch

**Bad Example:**

Old format
  - fields not present in a patch for Kind foo are ignored
  - unmodified fields for Kind foo are optional in patch request

New format
  - fields not present in a patch for Kind foo are cleared
  - unmodified fields for Kind foo are required in patch request

This example won't work, because old patch formats will contain data that is now
considered required.  To support this, introduce a new directive to guard the
new patch format.

#### 2. Add support for new directives in the patch format

- Optional directives may be introduced to change how the patch is applied by the server - **backwards compatible** (old patch against newer server).
  - May control how the patch is applied
  - May contain patch information - such as elements to delete from a list
  - Must NOT impose new requirements on the old patch format

- New patch requests should be a superset of old patch requests - **forwards compatible** (newer patch against older server)
  - *Old servers will ignore directives they do not recognize*
  - Must include the full patch that would have been sent before the new directives were added.
  - Must NOT rely on the directive being supported by the server

**Good Example:**

Old format
  - fields not present in a patch for Kind foo are ignored
  - unmodified fields for Kind foo are optional in patch request

New format *without* directive
  - Same as old

New format *with* directive
  - fields not present in a patch for Kind foo are cleared
  - unmodified fields for Kind foo are required in patch request

In this example, the behavior was unchanged when the directive was missing,
retaining the old behavior for old patch requests.

**Bad Example:**

Old format
  - fields not present in a patch for Kind foo are ignored
  - unmodified fields for Kind foo are optional in patch request

New format *with* directive
  - Same as old

New format *without* directive
  - fields not present in a patch for Kind foo are cleared
  - unmodified fields for Kind foo are required in patch request

In this example, the behavior was changed when the directive was missing,
breaking compatibility.

## Alternatives

The previous strategy is necessary because there is no notion of
patch versions.  Having the client negotiate the patch version
with the server would allow changing the patch format, but at
the cost of supporting multiple patch formats in the server and client.
Using client provided directives to evolve how a patch is merged
provides some limited support for multiple versions.

