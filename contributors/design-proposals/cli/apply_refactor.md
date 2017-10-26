# Apply v2

## Background

`kubectl apply` reads a file or set of files, and updates the cluster state based off the file contents.
It does a couple things:

1. Create / Update / (Delete) the live resources based on the file contents
2. Update currently and previously configured fields, without clobbering fields set by other means,
   such as imperative kubectl commands, other deployment and management tools, admission controllers,
   initializers, horizontal and vertical autoscalers, operators, and other controllers.

Essential complexity in the apply code comes from supporting custom strategies for
merging fields built into the object schema -
such as merging lists together based on a field `key` and deleting individual
items from a list by `key`.

Accidental complexity in the apply code comes from the structure growing organically in ways that have
broken encapsulation and separation of concerns.  This has lead to maintenance challenges as
keeping ordering for items in a list, and correctly merging lists of primitives (`key`less).

Round tripping changes through PATCHes introduces additional accidental complexity,
as they require imperative directives that are not part of the object schema.

## Objective


Reduce maintenance burden by minimizing accidental complexity in the apply codebase.

This should help:

- Simplify introducing new merge semantics
- Simplify enabling / disabling new logic with flags

## Changes

Implementation of proposed changes under review in PR [52349](https://github.com/kubernetes/kubernetes/pull/52349)

### Use read-update instead of patch

#### Why

Building a PATCH from diff creates additional code complexity vs directly updating the object.

- Need to generate imperative delete directives instead of simply deleting an item from a list.
- Using PATCH semantics and directives is less well known and understood by most users
  than using the object schema itself.  This makes it harder for non-experts to maintain the codebase.
- Using PATCH semantics is more work to implement a diff of the changes as
  PATCH must be separately merged on the remote object for to display the diff.

#### New approach

1. Read the live object
2. Compare the live object to last-applied and local files
3. Update the fields on the live object that was read
4. Send a PUT to update the modified object
5. If encountering optimistic lock failure, retry back to 1.

### Restructure code into modular components

In the current implementation of apply - parsing and traversing the object trees, diffing the
contents and generating the patch are entangled.  This creates maintenance and
testing challenges.  We should instead encapsulate discrete responsibilities in separate packages -
such as collating the object values and updating the target object.

#### Phase 1: Parse last-applied, local, live objects and collate

Provide a structure that contains the last, local and live value for each field.  This
will make it easy to walk the a single tree when making decisions about how to update the object.
Decisions about ordering of lists or parsing metadata for fields are made here.

#### Phase 2: Diff and update objects

Use the visitor pattern to encapsulate how to update each field type for each merge strategy.
Unit test each visit function.  Decisions about how to replace, merge, or delete a field or
list item are made here.

## Notable items

- Merge will use openapi to get the schema from the server
- Merge can be run either on the server side or the client side
- Merge can handle 2-way or 3-way merges of objects (initially will not support PATCH directives)

## Out of scope of this doc

In order to make apply sufficiently maintainable and extensible to new API types, as well as to make its
behavior more intuitive for users, the merge behavior, including how it is specified in the API schema,
must be systematically redesigned and more thoroughly tested.

Examples of issues that need to be resolved

- schema metadata `patchStrategy` and `mergeKey` are implicit, unversioned and incorrect in some cases.
  to fix the incorrect metadata, the metadata must be versioned so PATCHes generated will old metadata continue
  to be merged by the server in the manner they were intended
  - need to version all schema metadata for each objects and provide this are part of the request
  - e.g. container port [39188](https://github.com/kubernetes/kubernetes/issues/39188)
- no semantic way to represent union fields [35345](https://github.com/kubernetes/kubernetes/issues/35345)


## Detailed analysis of structure and impact today

The following PRs constitute the focus of ~6 months of engineering work.  Each of the PRs is very complex
for the work what it is solving.

### Patterns observed

- PRs frequently closed or deferred because maintainers / reviewers cannot reason about the impact or
  correctness of the changes
- Relatively simple changes
  - are 200+ lines of code
  - modify dozens of existing locations in the code
  - are spread across 1000+ lines of existing code
- Changes that add new directives require updates in multiple locations - create patch + apply patch

### PRs

[38665](https://github.com/kubernetes/kubernetes/pull/38665/files)
- Support deletion of primitives from lists
- Lines (non-test): ~200
- ~6 weeks
[44597](https://github.com/kubernetes/kubernetes/pull/44597/files)
- Support deleting fields not listed in the patch
- Lines (non-test): ~250
- ~6 weeks
[45980](https://github.com/kubernetes/kubernetes/pull/45980/files#diff-101008d96c4444a5813f7cb6b54aaff6)
- Keep ordering of items when merging lists
- Lines (non-test): ~650
[46161](https://github.com/kubernetes/kubernetes/pull/46161/files#diff-101008d96c4444a5813f7cb6b54aaff6)
- Support using multiple fields for a merge key
- Status: Deferred indefinitely - too hard for maintainers to understand impact and correctness of changes
[46560](https://github.com/kubernetes/kubernetes/pull/46560/files)
- Support diff apply (1st attempt)
- Status: Closed - too hard for maintainers to understand impact and correctness of changes
[49174](https://github.com/kubernetes/kubernetes/pull/49174/files)
- Support diff apply (2nd attempt)
- Status: Deferred indefinitely - too hard for maintainers to understand impact and correctness of changes
- Maintainer reviews: 3


### Analysis - causes of complexity

Apply is implemented by diffing the 3 sources (last-applied, local, remote) as 2 2-way diffs and then
merging the results of those 2 diffs into a 3rd result.  The diffs can each produce patch request where
a single logic update (e.g. remove 'foo' and add 'bar' to a field that is a list) may require spreading the
patch result across multiple pieces of the patch (a 'delete' directive, an 'order' directive
and the list itself).

Because of the way diff is implemented with 2-way diffs, a simple bit of logic
"compare local to remote" and do X - is non-trivial to define.  The code that compares local to remote
is also executed to compare last-applied to local, but with the local argument differing in location.
To compare local to remote means understanding what will happen when the same code is executed
comparing last-applied to local, and then putting in the appropriate guards to short-circuit the
logic in one context or the other as needed.  last-applied and remote are not compared directly, and instead
are only compared indirectly when the 2 diff results are merged.  Information that is redundant or
should be checked for consistency across all 3 sources (e.g. checking for conflicts) is spread across
3 logic locations - the first 2-way diff, the second 2-way diff and the merge of the 2 diffs.

That the diffs each may produce multiple patch directives + results that constitute an update to a single
field compounds the complexity of that comparing a single field occurs across 3 locations.
 
The diff / patch logic itself does not follow any sort of structure to encapsulate complexity
into components so that logic doesn't bleed cross concerns.  The logic to collate the last-applied, local and
remote field values, the logic to diff the field values and the logic to create the patch is
all combined in the same group of package-scoped functions, instead of encapsulating
each of these responsibilities in its own interface.

Sprinkling the implementation across dozens of locations makes it very challenging to
flag guard the new behavior.  If issues are discovered during the stabilization period we cannot
easily revert to the previous behavior by changing a default flag value.  The inability to build
in these sorts of break-glass options further degrades confidence in safely accepting PRs.

This is a text-book example of what the [Visitor pattern](https://en.wikipedia.org/wiki/Visitor_pattern) 
was designed to address.

- Encapsulate logic in *Element*s and *Visitor*s
- Introduce logic for new a field type by adding a new *Element* type
- Introduce logic for new a merge strategy by defining a new *Visitor* implementation
- Introduce logic on structuring of a field by updating the parsing function for that field type

If the apply diff logic was redesigned, most of the preceding PRs could be implemented by
only touching a few existing code locations to introduce the new type / method, and
then encapsulating the logic in a single type.  This would make it simple to flag guard
new behaviors before defaulting them to on.

