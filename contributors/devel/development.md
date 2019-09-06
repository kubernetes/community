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

### rsync

Kubernetes build system requires `rsync` command present in the development
platform.

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
| 1.11           | 1.10.2      |
| 1.12           | 1.10.4      |
| 1.13           | 1.11.13     |
| 1.14+          | 1.12.9      |

Note that Go 1.13 is not supported yet.

Ensure your GOPATH and PATH have been configured in accordance with the Go
environment instructions.

#### Upgrading Go

Upgrading Go requires specific modification of some scripts and container
images.

- The image for cross compiling in [build/build-image/cross].
  The `VERSION` file and `Dockerfile`.
- The cross tag `KUBE_BUILD_IMAGE_CROSS_TAG` in [build/common.sh].
- The `go_version` in the [`go_register_toolchains`](https://git.k8s.io/kubernetes/build/root/WORKSPACE) bazel rule.
- The desired Go verion in
  [test/images/Makefile](https://git.k8s.io/kubernetes/test/images/Makefile).

### Quick Start

The following section is a quick start on how to build Kubernetes locally, for more detailed information you can see [kubernetes/build](https://git.k8s.io/kubernetes/build/README.md).
The best way to validate your current setup is to build a small part of Kubernetes. This way you can address issues without waiting for the full build to complete. To build a specific part of Kubernetes use the `WHAT` environment variable to let the build scripts know you want to build only a certain package/executable.

```sh
make WHAT=cmd/{$package_you_want}
```

*Note:* This applies to all top level folders under kubernetes/cmd.

So for the cli, you can run:

```sh
make WHAT=cmd/kubectl
```

If everything checks out you will have an executable in the `_output/bin` directory to play around with.

*Note:* If you are using `CDPATH`, you must either start it with a leading colon, or unset the variable. The make rules and scripts to build require the current directory to come first on the CD search path in order to properly navigate between directories.

```sh
cd $working_dir/kubernetes
make
```

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

#### Install etcd

```sh
cd $working_dir/kubernetes

# Installs in ./third_party/etcd
hack/install-etcd.sh

# Add to PATH
echo export PATH="\$PATH:$working_dir/kubernetes/third_party/etcd" >> ~/.profile
```

#### Test

```sh
cd $working_dir/kubernetes

# Run all the presubmission verification. Then, run a specific update script (hack/update-*.sh)
# for each failed verification. For example:
#   hack/update-gofmt.sh (to make sure all files are correctly formatted, usually needed when you add new files)
#   hack/update-bazel.sh (to update bazel build related files, usually needed when you add or remove imports)
make verify

# Alternatively, run all update scripts to avoid fixing verification failures one by one.
make update

# Run every unit test
make test

# Run package tests verbosely
make test WHAT=./pkg/api/helper GOFLAGS=-v

# Run integration tests, requires etcd
# For more info, visit https://git.k8s.io/community/contributors/devel/sig-testing/testing.md#integration-tests
make test-integration

# Run e2e tests by building test binaries, turn up a test cluster, run all tests, and tear the cluster down
# Equivalent to: go run hack/e2e.go -- -v --build --up --test --down
# Note: running all e2e tests takes a LONG time! To run specific e2e tests, visit:
# ./e2e-tests.md#building-kubernetes-and-running-the-tests
make test-e2e
```

See the [testing guide](./sig-testing/testing.md) and [end-to-end tests](./sig-testing/e2e-tests.md)
for additional information and scenarios.

Run `make help` for additional information on these make targets.

#### Dependency management

Kubernetes uses [go modules](https://github.com/golang/go/wiki/Modules) to manage
dependencies.

Developers who need to manage dependencies in the `vendor/` tree should read
the docs on [using go modules to manage dependencies](/contributors/devel/sig-architecture/vendor.md).


## Build with Bazel/Gazel

Building with Bazel is currently experimental.  For more information,
see [Build with Bazel].


## GitHub workflow

To check out code to work on, please refer to [this guide](/contributors/guide/github-workflow.md).


[macOS GNU tools]: https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x
[build/build-image/cross]: https://git.k8s.io/kubernetes/build/build-image/cross
[build/common.sh]: https://git.k8s.io/kubernetes/build/common.sh
[etcd-latest]: https://coreos.com/etcd/docs/latest
[etcd-install]: sig-testing/integration-tests.md#install-etcd-dependency
<!-- https://github.com/coreos/etcd/releases -->
[go-workspace]: https://golang.org/doc/code.html#Workspaces
[issue]: https://github.com/kubernetes/kubernetes/issues
[kubectl user guide]: https://kubernetes.io/docs/user-guide/kubectl
[kubernetes.io]: https://kubernetes.io
[mercurial]: http://mercurial.selenic.com/wiki/Download
[Build with Bazel]: sig-testing/bazel.md
