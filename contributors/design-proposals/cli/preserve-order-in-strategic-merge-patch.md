# Preserve Order in Strategic Merge Patch

Author: @mengqiy

## Motivation

Background of the Strategic Merge Patch is covered [here](/contributors/devel/sig-api-machinery/strategic-merge-patch.md).

The Kubernetes API may apply semantic meaning to the ordering of items within a list,
however the strategic merge patch does not keep the ordering of elements.
Ordering has semantic meaning for Environment variables,
as later environment variables may reference earlier environment variables,
but not the other way around.

One use case is the environment variables. We don't preserve the order which causes
issue [40373](https://github.com/kubernetes/kubernetes/issues/40373).

## Proposed Change

We will use the following notions through the doc.
Notion:
list to be merged: same as live list, which is the list current in the server.
parallel list: the list with `$setElementOrder` directive in the patch.
patch list: the list in the patch that contains the value changes.

Changes are all in strategic merge patch package.
The proposed solution is similar to the solution used for deleting elements from lists of primitives.

Add to the current patch, a directive ($setElementOrder) containing a list of element keys -
either the patch merge key, or for primitives the value. When applying the patch,
the server ensures that the relative ordering of elements matches the directive.

The server will reject the patch if it doesn't satisfy the following 2 requirements.
- the relative order of any two items in the `$setElementOrder` list
matches that in the patch list if they present.
- the items in the patch list must be a subset or the same as the `$setElementOrder` list if the directive presents.

The relative order of two items are determined by the following order:

1. relative order in the $setElementOrder if both items are present
2. else relative order in the patch if both items are present
3. else relative order in the server-side list if both items are present
4. else append to the end

If the relative order of the live config in the server is different from the order of the parallel list,
the user's patch will always override the order in the server.

Here is a simple example of the patch format:

Suppose we have a type called list. The patch will look like below.
The order from the parallel list ($setElementOrder/list) will be respected.

```yaml
$setElementOrder/list:
- A
- B
- C
list:
- A
- C
```

All the items in the server's live list but not in the parallel list will come before the parallel list.
The relative order between these appended items are kept.

The patched list will look like:

```
mergingList:
- serverOnlyItem1 \
  ...              |===> items in the server's list but not in the parallel list
- serverOnlyItemM /
- parallelListItem1 \
  ...                |===> items from the parallel list
- parallelListItemN /
```

### When $setElementOrder is not present and patching a list

The new directive $setElementOrder is optional.
When the $setElementOrder is missing,
relative order in the patch list will be respected.

Examples where A and C have been changed, B has been deleted and D has been added.

Patch:

```yaml
list:
- A'
- B'
- D
```

Live:

```yaml
list:
- B
- C
- A
```

Result:

```yaml
list:
- C # server-only item comes first
- A'
- B'
- D
```

### `$setElementOrder` may contain elements not present in the patch list

The $setElementOrder value may contain elements that are not present in the patch
but present in the list to be merged to reorder the elements as part of the merge.

Example where A & B have not changed:

Patch:

```yaml
$setElementOrder/list:
- A
- B
```

Live:

```yaml
list:
- B
- A
```

Result:

```yaml
list:
- A
- B
```

### When the list to be merged contains elements not found in `$setElementOrder`

If the list to be merged contains elements not found in $setElementOrder,
they will come before all elements defined in $setElementOrder, but keep their relative ordering.

Example where A & B have been changed:

Patch:

```yaml
$setElementOrder/list:
- A
- B
list:
- A
- B
```

Live:

```yaml
list:
- C
- B
- D
- A
- E
```

Result:

```yaml
list:
- C
- D
- E
- A
- B
```

### When `$setElementOrder` contains elements not found in the list to be merged

If `$setElementOrder` contains elements not found in the list to be merged,
the elements that are not found will be ignored instead of failing the request.

Patch:
```yaml
$setElementOrder/list:
- C
- A
- B
list:
- A
- B
```

Live:
```yaml
list:
- A
- B
```

Result:

```yaml
list:
- A
- B
```

## Version Skew and Backwards Compatibility

The new version patch is always a superset of the old version patch.
The new patch has one additional parallel list which will be dropped by the old server.

As mentioned [above](#when-setelementorder-is-not-present-and-patching-a-list),
the new directive is optional.
Patch requests without the directive will change a little,
but still be fully backward compatible.

### kubectl
If an old kubectl sends a old patch to a new server,
the server will honor the order in the list as mentioned above.
The behavior is a little different from before but is not a breaking change.

If a new kubectl sends a new patch to an old server, the server doesn't recognise the parallel list and will drop it.
So it will behave the same as before.

## Example

### List of Maps

We take environment variables as an example.
Environment variables is a list of maps with merge patch strategy.

Suppose we define a list of environment variables and we call them
the original environment variables:

```yaml
env:
- name: ENV1
  value: foo
- name: ENV2
  value: bar
- name: ENV3
  value: baz
```

Then the server appends two environment variables and reorder the list:

```yaml
env:
- name: ENV2
  value: bar
- name: ENV5
  value: server-added-2
- name: ENV1
  value: foo
- name: ENV3
  value: baz
- name: ENV4
  value: server-added-1
```

Then the user wants to change it from the original to the following using `kubectl apply`:

```yaml
env:
- name: ENV1
  value: foo
- name: ENV2
  value: bar
- name: ENV6
  value: new-env
```

The old patch without parallel list will looks like:

```yaml
env:
- name: ENV3
  $patch: delete
- name: ENV6
  value: new-env
```

The new patch will looks like below. It is the

```yaml
$setElementOrder/env:
- name: ENV1
- name: ENV2
- name: ENV6
env:
- name: ENV3
  $patch: delete
- name: ENV6
  value: new-env
```

After server applying the new patch:

```yaml
env:
- name: ENV5
  value: server-added-2
- name: ENV4
  value: server-added-1
- name: ENV1
  value: foo
- name: ENV2
  value: bar
- name: ENV6
  value: new-env
```

### List of Primitives

We take finalizers as an example.
finalizers is a list of strings.

Suppose we define a list of finalizers and we call them
the original finalizers:

```yaml
finalizers:
- a
- b
- c
```

Then the server appends two finalizers and reorder the list:

```yaml
finalizers:
- b
- e
- a
- c
- d
```

Then the user wants to change it from the original to the following using `kubectl apply`:

```yaml
finalizers:
- a
- b
- f
```

The old patch without parallel list will looks like:

```yaml
$deleteFromPrimitiveList/finalizers:
- c
finalizers:
- f
```

The new patch will looks like below. It is the

```yaml
$setElementOrder/finalizers:
- a
- b
- f
$deleteFromPrimitiveList/finalizers:
- c
finalizers:
- f
```

After server applying the patch:

```yaml
finalizers:
- e
- d
- a
- b
- f
```


# Alternative Considered

# 1. Use the patch list to set order

## Proposed Change

This approach can considered as merging the parallel list and patch list into one single list.

For list of maps, the patch list will have all entries that are
either a map that contains the mergeKey and other changes
or a map that contains the mergeKey only.

For list of primitives, the patch list will be the same as the list in users' local config.

## Reason of Rejection

It cannot work correctly in the following concurrent writers case,
because PATCH in k8s doesn't use optimistic locking, so the following may happen.

Live config is:

```yaml
list:
- mergeKey: a
  other: A
- mergeKey: b
  other: B
- mergeKey: c
  other: C
```

Writer foo first GET the object from the server.
It wants to delete B, so it calculate the patch and is about to send it to the server:

```yaml
list:
- mergeKey: a
- mergeKey: b
  $patch: delete
- mergeKey: c
```

Before foo sending the patch to the server,
writer bar GET the object and it want to update A.

Patch from bar is:

```yaml
list:
- mergeKey: a
  other: A'
- mergeKey: b
- mergeKey: c
```

After the server first applying foo's patch and then bar's patch,
the final result will be wrong.
Because entry b has been recreated which is not desired.

```yaml
list:
- mergeKey: a
  other: A
- mergeKey: b
- mergeKey: c
  other: C
```

# 2. Use $position Directive

## Proposed Change

Use an approach similar to [MongoDB](https://docs.mongodb.com/manual/reference/operator/update/position/).
When patching a list of maps with merge patch strategy,
use a new directive `$position` in each map in the list.

If the order in the user's config is different from the order of the live config,
we will insert the `$position` directive in each map in the list.
We guarantee that the order of the user's list will always override the order of live list.

All the items in the server's live list but not in the patch list will be append to the end of the patch list.
The relative order between these appended items are kept.
If the relative order of live config in the server is different from the order in the patch,
user's patch will always override the order in the server.

When patching a list of primitives with merge patch strategy,
we send a whole list from user's config.

## Version Skew

It is NOT backward compatible in terms of list of primitives.

When patching a list of maps:
- An old client sends an old patch to a new server, the server just merges the change and no reordering.
The server behaves the same as before.
- A new client sends a new patch to an old server, the server doesn't understand the new directive.
So it just simply does the merge.

When patching a list of primitives:
- An old client sends an old patch to a new server, the server will reorder the patch list which is sublist of user's.
The server has the WRONG behavior.
- A new client sends a new patch to an old server, the server will deduplicate after merging.
The server behaves the same as before.

## Example

For patching list of maps:

Suppose we define a list of environment variables and we call them
the original environment variables:
```yaml
env:
  - name: ENV1
    value: foo
  - name: ENV2
    value: bar
  - name: ENV3
    value: baz
```

Then the server appends two environment variables and reorder the list:
```yaml
env:
  - name: ENV2
    value: bar
  - name: ENV5
    value: server-added-2
  - name: ENV1
    value: foo
  - name: ENV3
    value: baz
  - name: ENV4
    value: server-added-1
```

Then the user wants to change it from the original to the following using `kubectl apply`:
```yaml
env:
  - name: ENV1
    value: foo
  - name: ENV2
    value: bar
  - name: ENV6
    value: new-env
```

The patch will looks like:
```yaml
env:
  - name: ENV1
    $position: 0
  - name: ENV2
    $position: 1
  - name: ENV6
    value: new-env
    $position: 2
  - name: ENV3
    $patch: delete
```

After server applying the patch:
```yaml
env:
  - name: ENV1
    value: foo
  - name: ENV2
    value: bar
  - name: ENV6
    value: new-env
  - name: ENV5
    value: server-added-2
  - name: ENV4
    value: server-added-1
```


