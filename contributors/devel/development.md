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

These issues should always be submitted with (in decreasing order of value):

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

macOS ships with outdated BSD-based tools. We recommend installing [macOS GNU
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
| 1.7            | 1.8.1       |
| 1.8            | 1.8.3       |
| 1.9            | 1.9.1       |
| 1.10           | 1.9.1       |
| 1.11+          | 1.10.1      |

Ensure your GOPATH and PATH have been configured in accordance with the Go
environment instructions.

#### Upgrading Go

Upgrading Go requires specific modification of some scripts and container
images.

- The image for cross compiling in [build/build-image/cross].
  The `VERSION` file and `Dockerfile`.
- Update the desired Go version in Dockerfile for the [e2e][e2e-image] and [test][test-image].
  This requires pushing the [e2e][e2e-image] and [test][test-image] images that are `FROM` the desired Go version.
- The cross tag `KUBE_BUILD_IMAGE_CROSS_TAG` in [build/common.sh].


#### Dependency management

Kubernetes uses [`godep`](https://github.com/tools/godep) to manage
dependencies.

Developers who need to manage dependencies in the `vendor/` tree should read
the docs on [using godep to manage dependencies](godep.md).


## Build with Bazel/Gazel

Building with Bazel is currently experimental.  For more information,
see [Build with Bazel].


## GitHub workflow

To check out code to work on, please refer to [this guide](/contributors/guide/github-workflow.md).


[macOS GNU tools]: https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x
[build/build-image/cross]: https://git.k8s.io/kubernetes/build/build-image/cross
[build/common.sh]: https://git.k8s.io/kubernetes/build/common.sh
[e2e-image]: https://git.k8s.io/test-infra/jenkins/e2e-image
[etcd-latest]: https://coreos.com/etcd/docs/latest
[etcd-install]: testing.md#install-etcd-dependency
<!-- https://github.com/coreos/etcd/releases -->
[go-workspace]: https://golang.org/doc/code.html#Workspaces
[issue]: https://github.com/kubernetes/kubernetes/issues
[kubectl user guide]: https://kubernetes.io/docs/user-guide/kubectl
[kubernetes.io]: https://kubernetes.io
[mercurial]: http://mercurial.selenic.com/wiki/Download
[test-image]: https://git.k8s.io/test-infra/jenkins/test-image
[Build with Bazel]: bazel.md
