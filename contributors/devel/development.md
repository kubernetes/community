# Development Guide

This document is the canonical source of truth for things like supported
toolchain versions for building Kubernetes.

Please submit an [issue] on Github if you
* Notice a requirement that this doc does not capture.
* Find a different doc that specifies requirements (the doc should instead link
  here).

Development branch requirements will change over time, but release branch
requirements are frozen.

## Pre submit flight checks

Determine whether your issue or pull request is improving Kubernetes'
architecture or whether it's simply fixing a bug.

If you need a diagram, add it.  SEPARATE the description of the problem (e.g. Y
is a critical component that is too slow for an SLA that we care about) from the
solution (e.g. make X faster).

Some of these checks were less common in Kubernetes' earlier days. Now that we
have over 1000 contributors, each issue should be filed with care. No issue
should take more than 5 minutes to check for sanity (even the busiest of
reviewers can spare 5 minutes to review a patch that is thoughtfully justified).

### Is this just a simple bug fix?

Simple bug patches are easy to review since test coverage is submitted with the
patch.  Bug fixes don't usually require a lot of extra testing, but please
update the unit tests so they catch the bug!

### Is this an architecture improvement?

Some examples of "Architecture" improvements include:

- Adding a new feature or making a feature more configurable or modular.
- Improving test coverage.
- Decoupling logic or creation of new utilities.
- Making code more resilient (sleeps, backoffs, reducing flakiness, etc.).

These sorts of improvements are easily evaluated, especially when they decrease
lines of code without breaking functionality.  That said, please explain exactly
what you are 'cleaning up' in your Pull Request so as not to waste a reviewer's
time.

If you're making code more resilient, include tests that demonstrate the new
resilient behavior.  For example: if your patch causes a controller to better
handle inconsistent data, make a mock object which returns incorrect data a few
times and verify the controller's new behaviour.

### Is this a performance improvement ?

Performance bug reports MUST include data that demonstrates the bug.  Without
data, the issue will be closed.  You can measure performance using kubemark,
scheduler_perf, go benchmark tests, or e2e tests on a real cluster with metric
plots.

Examples of how NOT to suggest a performance bug (these lead to a long review
process and waste cycles):

- We *should* be doing X instead of Y because it *might* lead to better
  performance.
- Doing X instead of Y would reduce calls to Z.

The above statements have no value to a reviewer because neither is backed by
data. Writing issues like this lands your PR in a no-man's-land and waste your
reviewers' time.

Examples of possible performance improvements include (remember, you MUST
document the improvement with data):

- Improving a caching implementation.
- Reducing calls to functions which are O(n^2)
- Reducing dependence on API server requests.
- Changing the value of default parameters for processes, or making those values
  'smarter'.
- Parallelizing a calculation that needs to run on a large set of node/pod
  objects.

These issues should always be submitted with (in decreasing order or value):

- A golang Benchmark test.
- A visual depiction of reduced metric load on a cluster (measurable using
  metrics/ endpoints and grafana).
- A hand-instrumented timing test (i.e. adding some logs into the controller
  manager).

Here are some examples of properly submitted performance issues.  If you are new
to kubernetes and thinking about filing a performance optimization, re-read one
or all of these before you get started.

- https://github.com/kubernetes/kubernetes/issues/18266 (apiserver)
- https://github.com/kubernetes/kubernetes/issues/32833 (node)
- https://github.com/kubernetes/kubernetes/issues/31795 (scheduler)

Since performance improvements can be empirically measured, you should follow
the "scientific method" of creating a hypothesis, collecting data, and then
revising your hypothesis.  The above issues do this transparently, using figures
and data rather then conjecture. Notice that the problem is analyzed and a
correct solution is created before a single line of code is reviewed.

## Building Kubernetes with Docker

Official releases are built using Docker containers. To build Kubernetes using
Docker please follow [these
instructions](http://releases.k8s.io/HEAD/build/README.md).

## Building Kubernetes on a local OS/shell environment

Kubernetes development helper scripts assume an up-to-date GNU tools
environment. Recent Linux distros should work out-of-the-box.

Mac OS X ships with outdated BSD-based tools. We recommend installing [OS X GNU
tools].

### etcd

Kubernetes maintains state in [`etcd`][etcd-latest], a distributed key store.

Please [install it locally][etcd-install] to run local integration tests.

### Go

Kubernetes is written in [Go](http://golang.org). If you don't have a Go
development environment, please [set one up](http://golang.org/doc/code.html).


| Kubernetes     | requires Go |
|----------------|-------------|
| 1.0 - 1.2      | 1.4.2       |
| 1.3, 1.4       | 1.6         |
| 1.5, 1.6       | 1.7 - 1.7.5 |
| 1.7+           | 1.8.1       |

Ensure your GOPATH and PATH have been configured in accordance with the Go
environment instructions.

#### Upgrading Go

Upgrading Go requires specific modification of some scripts and container
images.

- The image for cross compiling in [build/build-image/cross].
  The `VERSION` file and `Dockerfile`.
- Update [dockerized-e2e-runner.sh] to run a kubekins-e2e with the desired Go
  version. This requires pushing the [e2e][e2e-image] and [test][test-image]
  images that are `FROM` the desired Go version.
- The cross tag `KUBE_BUILD_IMAGE_CROSS_TAG` in [build/common.sh].


#### Dependency management

Kubernetes build/test scripts use [`godep`](https://github.com/tools/godep) to
manage dependencies.

```sh
go get -u github.com/tools/godep
```

The Godep version that Kubernetes is using is listed in `Godep/Godep.json` (in
the kubernetes repo root). See what version you are running with this command:

```sh
godep version
```

Developers planning to manage dependencies in the `vendor/` tree may want to
explore alternative environment setups. See [using godep to manage
dependencies](godep.md).



## Build with Bazel/Gazel

Building with Bazel is currently experimental.  For more information,
see [Build with Bazel].



## Workflow

![Git workflow](git_workflow.png)

### 1 Fork in the cloud

1. Visit https://github.com/kubernetes/kubernetes
2. Click `Fork` button (top right) to establish a cloud-based fork.

### 2 Clone fork to local storage

Per Go's [workspace instructions][go-workspace], place Kubernetes' code on your
`GOPATH` using the following cloning procedure.

Define a local working directory:

```sh
# If your GOPATH has multiple paths, pick
# just one and use it instead of $GOPATH here.
# You must follow exactly this pattern,
# neither `$GOPATH/src/github.com/${your github profile name/`
# nor any other pattern will work.
working_dir=$GOPATH/src/k8s.io
```

> If you already do Go development on github, the `k8s.io` directory
> will be a sibling to your existing `github.com` directory.

Set `user` to match your github profile name:

```sh
user={your github profile name}
```

Both `$working_dir` and `$user` are mentioned in the figure above.

Create your clone:

```sh
mkdir -p $working_dir
cd $working_dir
git clone https://github.com/$user/kubernetes.git
# or: git clone git@github.com:$user/kubernetes.git

cd $working_dir/kubernetes
git remote add upstream https://github.com/kubernetes/kubernetes.git
# or: git remote add upstream git@github.com:kubernetes/kubernetes.git

# Never push to upstream master
git remote set-url --push upstream no_push

# Confirm that your remotes make sense:
git remote -v
```

### 3 Branch

Get your local master up to date:

```sh
cd $working_dir/kubernetes
git fetch upstream
git checkout master
git rebase upstream/master
```

Branch from it:
```sh
git checkout -b myfeature
```

Then edit code on the `myfeature` branch.

#### Build

```sh
cd $working_dir/kubernetes
make
```

(Note that this requires an unset `CDPATH` environment variable or one starting
with the null directory (i.e., contains a leading colon).

To remove the limit on the number of errors the Go compiler reports (default
limit is 10 errors):
```sh
make GOGCFLAGS="-e"
```

To build with optimizations disabled (enables use of source debug tools):

```sh
make GOGCFLAGS="-N -l"
```

To build binaries for all platforms:

```sh
make cross
```

#### Test

```sh
cd $working_dir/kubernetes

# Run every unit test
make test

# Run package tests verbosely
make test WHAT=pkg/util/cache GOFLAGS=-v

# Run integration tests, requires etcd
make test-integration

# Run e2e tests
make test-e2e
```

See the [testing guide](testing.md) and [end-to-end tests](e2e-tests.md)
for additional information and scenarios.

Run `make help` for additional information on these make targets.

### 4 Keep your branch in sync

```sh
# While on your myfeature branch
git fetch upstream
git rebase upstream/master
```

### 5 Commit

Commit your changes.

```sh
git commit
```
Likely you go back and edit/build/test some more then `commit --amend`
in a few cycles.

### 6 Push

When ready to review (or just to establish an offsite backup or your work),
push your branch to your fork on `github.com`:

```sh
git push -f ${your_remote_name} myfeature
```

### 7 Create a pull request

1. Visit your fork at https://github.com/$user/kubernetes
2. Click the `Compare & Pull Request` button next to your `myfeature` branch.
3. Check out the pull request [process](pull-requests.md) for more details.

_If you have upstream write access_, please refrain from using the GitHub UI for
creating PRs, because GitHub will create the PR branch inside the main
repository rather than inside your fork.

#### Get a code review

Once your pull request has been opened it will be assigned to one or more
reviewers.  Those reviewers will do a thorough code review, looking for
correctness, bugs, opportunities for improvement, documentation and comments,
and style.

Commit changes made in response to review comments to the same branch on your
fork.

Very small PRs are easy to review.  Very large PRs are very difficult to review.
At the assigned reviewer's discretion, a PR may be switched to use
[Reviewable](https://reviewable.k8s.io) instead.  Once a PR is switched to
Reviewable, please ONLY send or reply to comments through Reviewable.  Mixing
code review tools can be very confusing.

See [Faster Reviews](faster_reviews.md) for some thoughts on how to streamline
the review process.


#### Squash and Merge

Upon merge (by either you or your reviewer), all commits left on the review
branch should represent meaningful milestones or units of work.  Use commits to
add clarity to the development and review process.

Before merging a PR, squash any _fix review feedback_, _typo_, and _rebased_
sorts of commits.

It is not imperative that every commit in a PR compile and pass tests
independently, but it is worth striving for.

For mass automated fixups (e.g. automated doc formatting), use one or more
commits for the changes to tooling and a final commit to apply the fixup en
masse. This makes reviews easier.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/development.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->

[OS X GNU tools]: https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x
[build/build-image/cross]: https://github.com/kubernetes/kubernetes/blob/master/build/build-image/cross
[build/common.sh]: https://github.com/kubernetes/kubernetes/blob/master/build/common.sh
[dockerized-e2e-runner.sh]: https://github.com/kubernetes/test-infra/blob/master/jenkins/dockerized-e2e-runner.sh
[e2e-image]: https://github.com/kubernetes/test-infra/tree/master/jenkins/e2e-image
[etcd-latest]: https://coreos.com/etcd/docs/latest
[etcd-install]: testing.md#install-etcd-dependency
<!-- https://github.com/coreos/etcd/releases -->
[go-workspace]: https://golang.org/doc/code.html#Workspaces
[issue]: https://github.com/kubernetes/kubernetes/issues
[kubectl user guide]: https://kubernetes.io/docs/user-guide/kubectl
[kubernetes.io]: https://kubernetes.io
[mercurial]: http://mercurial.selenic.com/wiki/Download
[test-image]: https://github.com/kubernetes/test-infra/tree/master/jenkins/test-image
[Build with Bazel]: bazel.md
