Strategic Merge Patch
=====================

# Background

TODO: @pwittrock complete this section

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
Currently the metadata is available as struct tags on the API objects
themselves, but will become available to clients as Swagger annotations in the
future. In the above example, the `patchStrategy` metadata for the `containers`
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


TODO: @pwittrock
# Changing Patch Format

## Purpose

## Requirement

### Version Skew

## Strategy

## Example
