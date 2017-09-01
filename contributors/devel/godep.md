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

If there's every any doubt about the correctness of your dependencies, you can
simply `make clean` or `rm -rf _output`, and run it again.

Now you should have a clean copy of all of the Kubernetes dependencies.

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

After all of this is done, `git status` should show you what files have been
modified and added/removed.  Make sure to sanity-check them with `git diff`, and
to `git add` and `git rm` them, as needed.  It is commonly advised to make one
`git commit` which includes just the dependencies and Godeps files, and
another `git commit` that includes changes to Kubernetes code to use (or stop
using) the new/updated/removed dependency.  These commits can go into a single
pull request.

## Sanity checking

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


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/godep.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
