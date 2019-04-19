**Note**: Kubernetes now manages dependencies using go modules.
See [current documentation for working with dependencies](./vendor.md) for master branch development.
This document only applies to Kubernetes 1.14.x and earlier,
and should be removed once Kubernetes 1.14.x is no longer supported.

# Using godep to manage dependencies

This document is intended to show a way for managing `vendor/` tree dependencies
in Kubernetes. If you do not need to manage vendored dependencies, you probably
do not need to read this.

## Background

As a tool, `godep` leaves much to be desired.  It builds on `go get`, and adds
the ability to pin dependencies to exact git version. The `go get` tool itself
doesn't have any concept of versions, and tends to blow up if it finds a git
repo synced to anything but `master`, but that is exactly the state that
`godep` leaves repos.  This is a recipe for frustration when people try to use
the tools.

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

The `go` toolchain assumes a global workspace that hosts all of your Go code.

The `godep` tool operates by first "restoring" dependencies into your `$GOPATH`.
This reads the `Godeps.json` file, downloads all of the dependencies from the
internet, and syncs them to the specified revisions.  You can then make
changes - sync to different revisions or edit Kubernetes code to use new
dependencies (and satisfy them with `go get`).  When ready, you tell `godep` to
"save" everything, which it does by walking the Kubernetes code, finding all
required dependencies, copying them from `$GOPATH` into the `vendor/` directory,
and rewriting `Godeps.json`.

This does not work well, when combined with a global Go workspace.  Instead, we
will set up a private workspace for this process.

The Kubernetes build process uses this same technique, and offers a tool called
`run-in-gopath.sh` which sets up and switches to a local, private workspace,
including setting up `$GOPATH` and `$PATH`.  If you wrap commands with this
tool, they will use the private workspace, which will not conflict with other
projects and is easily cleaned up and recreated.

To see this in action, you can run an interactive shell in this environment:

```sh
# Run a shell, but don't run your own shell initializations.
hack/run-in-gopath.sh bash --norc --noprofile
```

## Restoring deps

To extract and download dependencies into `$GOPATH` we provide a script:
`hack/godep-restore.sh`.  If you run this tool, it will restore into your own
`$GOPATH`.  If you wrap it in `run-in-gopath.sh` it will restore into your
`_output/` directory.

```sh
hack/run-in-gopath.sh hack/godep-restore.sh
```

This script will try to optimize what it needs to download, and if it seems the
dependencies are all present already, it will return very quickly.

If there's ever any doubt about the correctness of your dependencies, you can
simply `make clean` or `rm -rf _output`, and run it again.

Now you should have a clean copy of all of the Kubernetes dependencies.

Downloading dependencies might take a while, so if you want to see progress
information use the `-v` flag:

```sh
hack/run-in-gopath.sh hack/godep-restore.sh -v
```

## Making changes

The most common things people need to do with deps are add and update them.
These are similar but different.

### Adding a dep

For the sake of examples, consider that we have discovered a wonderful Go
library at `example.com/go/frob`.  The first thing you need to do is get that
code into your workspace:

```sh
hack/run-in-gopath.sh go get -d example.com/go/frob
```

This will fetch, but not compile (omit the `-d` if you want to compile it now),
the library into your private `$GOPATH`.  It will pull whatever the default
revision of that library is, typically the `master` branch for git repositories.
If this is not the revision you need, you can change it, for example to
`v1.0.0`:

```sh
hack/run-in-gopath.sh bash -c 'git -C $GOPATH/src/example.com/go/frob checkout v1.0.0'
```

Now that the code is present, you can start to use it in Kubernetes code.
Because it is in your private workspace's `$GOPATH`, it might not be part of
your own `$GOPATH`, so tools like `goimports` might not find it.  This is an
unfortunate side-effect of this process.  You can either add the whole private
workspace to your own `$GOPATH` or you can `go get` the library into your own
`$GOPATH` until it is properly vendored into Kubernetes.

Another possible complication is a dep that uses `gopdep` itself.  In that case,
you need to restore its dependencies, too:

```sh
hack/run-in-gopath.sh bash -c 'cd $GOPATH/src/example.com/go/frob && godep restore'
```

If the transitive deps collide with Kubernetes deps, you may have to manually
resolve things.  This is where the ability to run a shell in this environment
comes in handy:

```sh
hack/run-in-gopath.sh bash --norc --noprofile
```

### Updating a dep

Sometimes we already have a dep, but the version of it is wrong.  Because of the
way that `godep` and `go get` interact (badly) it's generally easiest to hit it
with a big hammer:

```sh
hack/run-in-gopath.sh bash -c 'rm -rf $GOPATH/src/example.com/go/frob'
hack/run-in-gopath.sh go get -d example.com/go/frob
hack/run-in-gopath.sh bash -c 'git -C $GOPATH/src/example.com/go/frob checkout v2.0.0'
```

This will remove the code, re-fetch it, and sync to your desired version.

### Removing a dep

This happens almost for free.  If you edit Kubernetes code and remove the last
use of a given dependency, you only need to restore and save the deps, and the
`godep` tool will figure out that you don't need that dep any more:

## Saving deps

Now that you have made your changes - adding, updating, or removing the use of a
dep - you need to rebuild the dependency database and make changes to the
`vendor/` directory.

```sh
hack/run-in-gopath.sh hack/godep-save.sh
```

This will run through all of the primary targets for the Kubernetes project,
calculate which deps are needed, and rebuild the database.  It will also
regenerate other metadata files which the project needs, such as BUILD files and
the LICENSE database.

Commit the changes before updating deps in staging repos.

## Saving deps in staging repos

Kubernetes stores some code in a directory called `staging` which is handled
specially, and is not covered by the above.  If you modified any code under
staging, or if you changed a dependency of code under staging (even
transitively), you'll also need to update deps there:

```sh
./hack/update-staging-godeps.sh
```

Then commit the changes generated by the above script.

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
`git commit` which includes just the dependencies and Godeps files, and
another `git commit` that includes changes to Kubernetes code to use (or stop
using) the new/updated/removed dependency.  These commits can go into a single
pull request.

Before sending your PR, it's a good idea to sanity check that your
Godeps.json file and the contents of `vendor/ `are ok:

```sh
hack/run-in-gopath.sh hack/verify-godeps.sh
```

All this script will do is a restore, followed by a save, and then look for
changes.  If you followed the above instructions, it should be clean.  If it is
not, you get to figure out why.

## Manual updates

It is sometimes expedient to manually fix the `Godeps.json` file to
minimize the changes. However, without great care this can lead to failures
with the verifier scripts.  The kubernetes codebase does "interesting things"
with symlinks between `vendor/` and `staging/` to allow multiple Go import
paths to coexist in the same git repo.

The verifiers, including `hack/verify-godeps.sh` *must* pass for every pull
request.

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
- Look at the godeps file. Check that the only changes are what the PR claims
  them to be. 
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
