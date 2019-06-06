**Note**: This document only applies to Kubernetes development after 1.14.x.
See [previous godep documentation for working with dependencies](./godep.md) for Kubernetes 1.14.x and earlier.

# Using go modules to manage dependencies

This document is intended to show a way for managing `vendor/` tree dependencies
in Kubernetes. If you do not need to manage vendored dependencies, you probably
do not need to read this.

## Background

Go modules allow recording desired versions of dependencies, and allow the main
module in a build to pin dependencies to specific versions.

This doc will focus on predictability and reproducibility.

## Justifications for an update

Before you update a dependency, take a moment to consider why it should be 
updated. Valid reasons include:
 1. We need new functionality that is in a later version.
 2. New or improved APIs in the dependency significantly improve Kubernetes code.
 3. Bugs were fixed that impact Kubernetes.
 4. Security issues were fixed even if they don't impact Kubernetes yet.
 5. Performance, scale, or efficiency was meaningfully improved.
 6. We need dependency A and there is a transitive dependency B.
 7. Kubernetes has an older level of a dependency that is precluding being able
to work with other projects in the ecosystem.

## Theory of operation

The `go.mod` file in the root of `k8s.io/kubernetes` describes dependencies using two directives:

* `require` directives list the preferred version of dependencies (this is auto-updated by go tooling to the maximum preferred version of the module)
* `replace` directives pin to specific tags or commits

## Adding or updating a dependency

The most common things people need to do with deps are add and update them.
These operations are handled the same way:

For the sake of examples, consider that we have discovered a wonderful Go
library at `example.com/go/frob`.

Step 1: Ensure there is go code in place that references the packages you want to use.
```go
import "example.com/go/frob"
...
frob.DoStuff()
```

Step 2: Determine what version of the dependency you want to use, and add that version to the go.mod file:

```sh
hack/pin-dependency.sh example.com/go/frob v1.0.4
```

This fetches the dependency, resolves the specified sha or tag, and adds two entries to the `k8s.io/kubernetes` `go.mod` file:

```
require (
    example.com/go/frob v1.0.4
    ...
)

replace (
    example.com/go/frob => example.com/go/frob v1.0.4
    ...
)
```

The `require` directive indicates our module requires `example.com/go/frob` >= `v1.0.4`.
If our module was included as a dependency in a build with other modules that also required `example.com/go/frob`,
the maximum required version would be selected (unless the main module in that build pinned to a lower version).

The `replace` directive pins us to the desired version when running go commands within kubernetes/kubernetes.

Step 3: Rebuild the `vendor` directory and update the `go.mod` files for all staging repositories:
```sh
hack/update-vendor.sh
```

Step 4: Check if the new dependency requires newer versions of existing dependencies we have pinned.
You can check this by:
1. running `hack/lint-dependencies.sh` against your branch and against `master` and comparing the results
2. checking if any new `replace` directives were added to `go.mod` files of components inside the staging directory.

Staging components with `replace` directives are the most problematic, because consumers of those components
will use different versions of libraries than the ones we build and test Kubernetes with by default.

If transitive dependencies need to be updated as a result of the new dependency,
run `hack/pin-dependency.sh` to update their version, and `hack/update-vendor.sh` again.
Repeat until step 4 shows no new transitive version requirements, compared to `master`.


### Removing a dependency

This happens almost for free.  If you edit Kubernetes code and remove the last
use of a given dependency, you only need to run `hack/update-vendor.sh`, and the
tooling will figure out that you don't need that dependency any more and remove it,
along with any unused transitive dependencies.

## Commit messages

Terse messages like "Update foo.org/bar to 0.42" are problematic
for maintainability.  Please include in your commit message the
detailed reason why the dependencies were modified.

Too commonly dependency changes have a ripple effect where something
else breaks unexpectedly.  The first instinct during issue triage
is to revert a change.  If the change was made to fix some other
issue and that issue was not documented, then a revert simply
continues the ripple by fixing one issue and reintroducing another
which then needs refixed.  This can needlessly span multiple days
as CI results bubble in and subsequent patches fix and refix and
rerefix issues.  This may be avoided if the original modifications
recorded artifacts of the change rationale.

## Sanity checking

After all of this is done, `git status` should show you what files have been
modified and added/removed.  Make sure to sanity-check them with `git diff`, and
to `git add` and `git rm` them, as needed.  It is commonly advised to make one
`git commit` which includes just the dependencies and `go.mod` and `go.sum` files, and
another `git commit` that includes changes to Kubernetes code to use (or stop
using) the new/updated/removed dependency.  These commits can go into a single
pull request.

Before sending your PR, it's a good idea to sanity check that your
go.mod, go.sum files and the contents of `vendor/` are ok:

```sh
hack/verify-vendor.sh
```

## Reviewing and approving dependency changes

Particular attention to detail should be exercised when reviewing and approving
PRs that add/remove/update dependencies. Importing a new dependency should bring
a certain degree of value as there is a maintenance overhead for maintaining
dependencies into the future.

When importing a new dependency, be sure to keep an eye out for the following:
- Is the dependency maintained?
- Does the dependency bring value to the project? Could this be done without
  adding a new dependency?
- Is the target dependency the original source, or a fork?
- Is there already a dependency in the project that does something similar?
- Does the dependency have a license that is compatible with the Kubernetes
  project?

Additionally:
- Look at the `go.mod` changes in `k8s.io/kubernetes`.
  Check that the only changes are what the PR claims them to be. 
- Look at the `go.mod` changes in the staging components.
  Avoid adding new `replace` directives in staging component `go.mod` files.
  New `replace` directives are problematic for consumers of those libraries,
  since it means we are pinned to older versions than would be selected by go
  when our module is used as a library.
- Check if there is a tagged release we can vendor instead of a random hash
- Scan the imported code for things like init() functions
- Look at the Kubernetes code changes and make sure they are appropriate
  (e.g. renaming imports or similar). You do not need to do feature code review.
- If this is all good, approve, but don't LGTM, unless you also do code review
  or unless it is trivial (e.g. moving from k/k/pkg/utils -> k/utils).

All new dependency licenses should be reviewed by @kubernetes/dep-approvers to ensure that they
are compatible with the Kubernetes project license. It is also important to note
and flag if a license has changed when updating a dependency, so that these can
also be reviewed.

For reference, whitelisted licenses as per the CNCF Whitelist Policy are
mentioned [here](https://git.k8s.io/sig-release/licensing/README.md#licenses-for-dependencies).
