# Development Guide

This document is the canonical source of truth for things like supported
toolchain versions for building Kubernetes.

Please submit an [issue] on GitHub if you

- Notice a requirement that this doc does not capture.
- Find a different doc that specifies requirements (the doc should instead link
  here).

Development branch requirements will change over time, but release branch
requirements are frozen.

## Pre submit flight checks

Determine whether your issue or pull request is improving Kubernetes'
architecture or whether it's simply fixing a bug.

If you need a diagram, add it. SEPARATE the description of the problem (e.g. Y
is a critical component that is too slow for an SLA that we care about) from the
solution (e.g. make X faster).

Some of these checks were less common in Kubernetes' earlier days. Now that we
have over 50000 contributors, each issue should be filed with care. No issue
should take more than 5 minutes to check for sanity (even the busiest of
reviewers can spare 5 minutes to review a patch that is thoughtfully justified).

### Is this just a simple bug fix?

Simple bug patches are easy to review since test coverage is submitted with the
patch. Bug fixes don't usually require a lot of extra testing, but please
update the unit tests so they catch the bug!

### Is this an architecture improvement?

Some examples of "Architecture" improvements include:

- Adding a new feature or making a feature more configurable or modular.
- Improving test coverage.
- Decoupling logic or creation of new utilities.
- Making code more resilient (sleeps, backoffs, reducing flakiness, etc.).

These sorts of improvements are easily evaluated, especially when they decrease
lines of code without breaking functionality. That said, please explain exactly
what you are 'cleaning up' in your Pull Request so as not to waste a reviewer's
time.

If you're making code more resilient, include tests that demonstrate the new
resilient behavior. For example: if your patch causes a controller to better
handle inconsistent data, make a mock object which returns incorrect data a few
times and verify the controller's new behaviour.

### Is this a performance improvement?

Performance bug reports MUST include data that demonstrates the bug. Without
data, the issue will be closed. You can measure performance using kubemark,
scheduler_perf, go benchmark tests, or e2e tests on a real cluster with metric
plots.

Examples of how NOT to suggest a performance bug (these lead to a long review
process and waste cycles):

- We _should_ be doing X instead of Y because it _might_ lead to better
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

Here are some examples of properly submitted performance issues. If you are new
to kubernetes and thinking about filing a performance optimization, re-read one
or all of these before you get started.

- https://github.com/kubernetes/kubernetes/issues/18266 (apiserver)
- https://github.com/kubernetes/kubernetes/issues/32833 (node)
- https://github.com/kubernetes/kubernetes/issues/31795 (scheduler)

Since performance improvements can be empirically measured, you should follow
the "scientific method" of creating a hypothesis, collecting data, and then
revising your hypothesis. The above issues do this transparently, using figures
and data rather than conjecture. Notice that the problem is analyzed and a
correct solution is created before a single line of code is reviewed.

## Building Kubernetes with Docker

Official releases are built using Docker containers. To build Kubernetes using
Docker please follow [these
instructions](http://releases.k8s.io/HEAD/build/README.md).

## Building Kubernetes on a local OS/shell environment

While building via Docker can be simpler, sometimes it makes sense to
do development on your local workstation or some other shell
environment. The details below outline the hardware and software
requirements for building on Linux, Windows, and macOS.

### Hardware Requirements

Kubernetes is a large project, and compiling it can use a lot of
resources. We recommend the following for any physical or virtual
machine being used for building Kubernetes.

- 8GB of RAM
- 50GB of free disk space

### Preparing Your Local Operating System

Where needed, each piece of required software will have separate
instructions for Linux, Windows, or macOS.

#### Setting Up Windows

If you are running Windows, you will need to use one of two methods
to set up your machine for Kubernetes development. To figure out which
method is the best choice, you will first need to determine which version of
Windows you are running. To do this, press **Windows logo key + R**,
type **winver**, and click **OK**. You may also enter the `ver` command at
the Windows Command Prompt.

1. If you're using Windows 10, Version 2004, Build 19041 or higher,
   you can use Windows Subsystem for Linux (WSL) to build
   Kubernetes. [Follow these instructions to install WSL2.](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
2. If you're using an earlier version of Windows, then create a Linux
   virtual machine with at least 8GB of memory and 60GB of disk space.

Once you have finished setting up your WSL2 installation or Linux VM,
follow the instructions below to configure your system for building
and developing Kubernetes.

#### Setting Up macOS

Since Kubernetes assumes you are using GNU command line tools, you will need to
install those tools on your
system. [Follow these directions to install the tools](https://ryanparman.com/posts/2019/using-gnu-command-line-tools-in-macos-instead-of-freebsd-tools/).
In particular, this command installs the necessary packages:

```sh
brew install coreutils ed findutils gawk gnu-sed gnu-tar grep make jq
```

### Installing Required Software

#### GNU Development Tools

Kubernetes development helper scripts require an up-to-date GNU
development tools environment. The method for installing these tools
varies from system to system.

##### Installing on Linux

All Linux distributions have the GNU tools available. The most popular
distributions and commands used to install these tools are below.

- Debian/Ubuntu
  ```sh
  sudo apt update
  sudo apt install build-essential
  ```
- Fedora/RHEL/CentOS
  ```sh
  sudo yum update
  sudo yum groupinstall "Development Tools"
  ```
- OpenSUSE
  ```sh
  sudo zypper update
  sudo zypper install -t pattern devel_C_C++
  ```
- Arch
  ```sh
  sudo pacman -Sy base-devel
  ```
  Alpin linux
  '''sh
  sudo apk update
  sudo apk add build-base
  '''
  Clear linux
  '''sh
  sudo swupd update
  sodo swupd bundle-add c-basic dev-utils
  '''
  Solus
  '''sh
  sudo eopkg update-repo
  sudo eopkg install-c system.devel

Once you have finished, confirm that `gcc` and `make` are installed.

##### Installing on macOS

Some of the build tools were installed when you prepared your system
with the GNU command line tools earlier. However, you will also need
to install the
[Command Line Tools for Xcode](https://developer.apple.com/library/archive/technotes/tn2339/_index.html).

#### Docker

Kubernetes development requires Docker to run certain verifications. To
install Docker in your development environment,
[follow the instructions from the Docker website](https://docs.docker.com/get-docker/).

**Note:** If you are running macOS, make sure that `/usr/local/bin` is
in your `PATH`.

#### rsync

The Kubernetes build system requires that `rsync`, a common file
synchronization and transfer tool, be present in the
development environment. Most modern operating systems come with
`rsync` already installed. If this is not the case, your operating
system's package manager can most likely install the `rsync`
package.

If this fails, check the [rsync download instructions page](https://rsync.samba.org/download.html).

#### jq

Some of the Kubernetes helper scripts require `jq`, a command-line JSON processor, to be
installed in your development environment. The
[jq installation guide](https://stedolan.github.io/jq/download/)
provides detailed instructions for supported platforms.

#### gcloud

If you plan to build remotely or run end-to-end (e2e) tests, you will
need to install the command line interface to the Google Cloud
Platform. [Follow the `gcloud` installation instructions for your operating system.](https://cloud.google.com/sdk/downloads)

#### Go

Kubernetes is written in [Go](http://golang.org). If you don't have a Go
development environment, please follow the instructions in the
[Go Getting Started guide](https://golang.org/doc/install).

Confirm that your `GOPATH` and `GOBIN` environment variables are
correctly set as detailed in
[How to Write Go Code](https://golang.org/doc/code.html) before
proceeding.

**Note:** Building and developing Kubernetes requires a very recent
version of Go. Please install the newest stable version available for
your system. The table below lists the required Go versions for
different versions of Kubernetes.

| Kubernetes  | requires Go |
| ----------- | ----------- |
| 1.0 - 1.2   | 1.4.2       |
| 1.3, 1.4    | 1.6         |
| 1.5, 1.6    | 1.7 - 1.7.5 |
| 1.7         | 1.8.1       |
| 1.8         | 1.8.3       |
| 1.9         | 1.9.1       |
| 1.10        | 1.9.1       |
| 1.11        | 1.10.2      |
| 1.12        | 1.10.4      |
| 1.13        | 1.11.13     |
| 1.14 - 1.16 | 1.12.9      |
| 1.17 - 1.18 | 1.13.15     |
| 1.19 - 1.20 | 1.15.5      |
| 1.21 - 1.22 | 1.16.7      |
| 1.23        | 1.17        |
| 1.24        | 1.18        |
| 1.25        | 1.20.10     |
| 1.26 - 1.29 | 1.21.7      |
| 1.30        | 1.22.1      |

[Go version for latest Kubernetes](https://cs.k8s.io/?q=golang%3A%20upstream%20version&i=nope&files=&excludeFiles=&repos=)

To find which Go is required for a specific Kubernetes version,
run the following commands in your Kubernetes working directory.
Below example looks for all 1.29.z Kubernetes releases.

```sh
K8S_VERSION=1.29
for tag in $(git tag | grep $K8S_VERSION);do git checkout -q tags/$tag;goVersion=$(cat ./build/dependencies.yaml | grep "golang: upstream version" -A 1 | grep version: | awk '{$1=$1;print}' );echo "Kubernetes $tag requires Go $goVersion";done
```

An example output will be

```sh
Kubernetes v1.29.0 requires Go version: 1.21.5
Kubernetes v1.29.0-alpha.0 requires Go version: 1.20.6
Kubernetes v1.29.0-alpha.1 requires Go version: 1.21.1
Kubernetes v1.29.0-alpha.2 requires Go version: 1.21.2
Kubernetes v1.29.0-alpha.3 requires Go version: 1.21.3
Kubernetes v1.29.0-rc.0 requires Go version: 1.21.4
Kubernetes v1.29.0-rc.1 requires Go version: 1.21.4
Kubernetes v1.29.0-rc.2 requires Go version: 1.21.5
Kubernetes v1.29.1 requires Go version: 1.21.6
Kubernetes v1.29.2 requires Go version: 1.21.7
```

##### A Note on Changing Go Versions

If you have already compiled Kubernetes but are now trying with a
different version of Go, please refer to the
[SIG Release documentation](https://github.com/kubernetes/sig-release).

#### PyYAML

Some Kubernetes verification tests use [PyYAML](https://pyyaml.org/) and it therefore needs to be installed to successfully run all verification tests in your local environment.
You can use the
[PyYAML documentation](https://pyyaml.org/wiki/PyYAMLDocumentation) to
find the installation instructions for your platform.

**Note:** If you are running macOS, you may need to use the `pip3`
command instead of the `pip` command to install PyYAML.

#### Cloning the Kubernetes Git Repository

You are now ready to clone the Kubernetes git repository. See the [GitHub Workflow](/contributors/guide/github-workflow.md) document from the Contributor Guide for instructions.

#### etcd

To test Kubernetes, you will need to install a recent version of [etcd](https://etcd.io/), a consistent and highly-available key-value store. To install a local version of etcd, run the following command in your Kubernetes working directory.

```sh
./hack/install-etcd.sh
```

This script will instruct you to make a change to your `PATH`. To make
this permanent, add this to your `.bashrc` or login script:

```sh
export PATH="$GOPATH/src/k8s.io/kubernetes/third_party/etcd:${PATH}"
```

##### BASH version requirement

To successfully run unit tests in Kubernetes, you will need bash version installed to be >4.3.

Once you have installed all required software, you can proceed to the
[Building Kubernetes](#building-kubernetes) section to test if it all works properly.

## Building Kubernetes

The best way to validate your development environment is to build part of Kubernetes. This allows you to address issues and correct your configuration without waiting for a full build to complete. This section briefly describes various methods for compiling Kubernetes subsystems. For more detailed instructions, see [Building Kubernetes](https://github.com/kubernetes/kubernetes/blob/master/build/README.md) in the official Kubernetes documentation.

To build a specific part of Kubernetes use the `WHAT` environment variable. In `$GOPATH/src/k8s.io/kubernetes/`, the Kubernetes project directory, run the following command:

```sh
make WHAT=cmd/<subsystem>
```

Replace `<subsystem>` with one of the command folders under the `cmd/` directory. For example, to build the `kubectl` CLI, run the following:

```sh
make WHAT=cmd/kubectl
```

If this command succeeds, you will now have an executable at `_output/bin/kubectl` off of your Kubernetes project directory.

To build the entire Kubernetes project, run the following command:

```sh
make all
```

**Note:** You can omit `all` and just run `make`.

The Kubernetes build system defaults to limiting the number of reported Go compiler errors to 10. If you would like to remove this limit, add `GOGCFLAGS="-e"` to your command line. For example:

```sh
make WHAT="cmd/kubectl" GOGCFLAGS="-e"
```

If you need to use debugging inspection tools on your compiled Kubernetes executables, set DBG=1. For example:

```sh
make WHAT="cmd/kubectl" DBG=1
```

To cross-compile Kubernetes for all platforms, run the following command:

```sh
make cross
```

To build binaries for a specific platform, add `KUBE_BUILD_PLATFORMS=<os>/<arch>`. For example:

```sh
make cross KUBE_BUILD_PLATFORMS=windows/amd64
```

## A Quick Start for Testing Kubernetes

Because kubernetes only merges pull requests when unit, integration, and e2e tests are
passing, your development environment needs to run all tests successfully. While this quick start will get you going,
to really understand the testing infrastructure, read the
[Testing Guide](sig-testing/testing.md) and check out the
[SIG Architecture developer guide material](README.md#sig-architecture).

Note that all of the commands in this section are run in your
Kubernetes project directory at `$GOPATH/src/k8s.io/kubernetes/`
unless otherwise specified.

**Note:** You can get additional information for many of the commands
mentioned here by running `make help`.

### Presubmission Verification

Presubmission verification provides a battery of checks and tests to
give your pull request the best chance of being accepted. Developers need to run as many verification tests as possible
locally.

You can view a list of all verification tests in `hack/verify-*.sh`
off of your Kubernetes project directory.

To run all presubmission verification tests, use this command:

```sh
make verify
```

If a specific verification test is failing, there could be an update
script to help fix the problem. These are located in
`hack/update-*.sh`. For example, `hack/update-gofmt.sh` makes sure
that all source code files are correctly formatted. This is usually
needed when you add new files to the project.

You can also run all update scripts with this command:

```sh
make update
```

### Unit Tests

Pull requests need to pass all unit tests. To run every unit test, use
this command:

```sh
make test
```

You can also use the `WHAT` option to control which packages and
subsystems are testing and use `GOFLAGS` to change how tests are
run. For example, to run unit tests verbosely against just one
package, use a command like this:

```
make test WHAT=./pkg/apis/core/helper GOFLAGS=-v
```

### Integration Tests

All integration tests need to pass for a pull request to be
accepted. Note that for this stage, in particular, it is important that
[etcd](#etcd) be properly installed. Without it, integration testing
will fail.

To run integration tests, use this command:

```sh
make test-integration
```

To learn more about integration testing, read the
[SIG Testing Integration Tests guide](./sig-testing/integration-tests.md).

### E2E Tests

End-to-end (E2E) tests provide a mechanism to test the end-to-end behavior
of the system. The primary objective of the E2E tests is to ensure
consistent and reliable behavior of the Kubernetes code base,
especially in areas where unit and integration tests are insufficient.

E2E tests build test binaries, spin up a test cluster,
run the tests, and then tear the cluster down.

**Note:** Running all E2E tests takes a _very long time_!

For more information on E2E tests, including methods for saving time
by just running specific tests, read
[End-to-End Testing in Kubernetes](./sig-testing/e2e-tests.md) and the
[getting started guide for `kubetest2`](./sig-testing/e2e-tests-kubetest2.md).

## Dependency management

Kubernetes uses [go modules](https://github.com/golang/go/wiki/Modules) to manage
dependencies.

Developers who need to manage dependencies in the `vendor/` tree should read
the docs on [using go modules to manage dependencies](/contributors/devel/sig-architecture/vendor.md).

### Building Kubernetes Using A Specific Version of Go

There exists a [`.go-version`](https://github.com/kubernetes/kubernetes/blob/f563910656ad325a7e1f8ab5848746bc2eba4d7f/.go-version)
file in the root of the Kubernetes repo. This file defines what version of `go` should be used
to build Kubernetes. So, for example, if you'd like to build with `go1.20.4` specifically, you
would change the contents of this file to just `1.20.4`.

The way that the build targets choose what `go` version to use is as follows:

- If the `go` version that exists on your system (determined by output of `go version`) does
  not match the version defined in `.go-version`, then default to the version specified in `.go-version`.
- If you do not want this behaviour, you can do one of the following:
  - Set the `GO_VERSION` environment variable. `GO_VERSION` defines the _desired_ version of
    `go` to be used. Even if the `go` version on your system does not match the one in `.go-version`,
    the version specified by `GO_VERSION` will be used (even if it needs to be downloaded).
    - The format of the version specified as part of `GO_VERSION` is the same as how a version
      would be defined in the `.go-version` file. So if you wanted to build with `go1.20.4`, you'd
      set `GO_VERSION=1.20.4`.
  - Set the `FORCE_HOST_GO` environment variable to a non-empty value. This will skip all the above
    logic and just use the `go` version that exists on your system's `$PATH`.

Some examples:

If you want to build using a `go` version (let's assume this is go1.20.4) that neither exists on your
system nor is the one that is specified in the `.go-version` file:

```
GO_VERSION=1.20.4 make WHAT=cmd/<subsystem>
```

If you want to build using the `go` version that exists on your system already and not really with what
exists in the `.go-version` file:

```
FORCE_HOST_GO=y make WHAT=cmd/<subsystem>
```

Or you can just change the contents of the `.go-version` file to your desired `go` version!  
`.go-version`:

```
1.20.4
```

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
