# Development Guide

This document is the canonical source of truth for things like supported
toolchain versions for building Kubernetes.

**Table of Contents**

- [Getting the code](#getting-the-code)
    - [GitHub workflow](#github-workflow)
- [Building Kubernetes](#building-kubernetes)
  - [Building Kubernetes with Docker](#building-kubernetes-with-docker)
  - [Building Kubernetes on a local OS/shell environment](#building-kubernetes-on-a-local-OS/shell-environment)
    - [rsync](#rsync)
    - [etcd](#etcd)
    - [Go](#go)
    - [Upgrading Go](#upgrading-go)
    - [Dependency management](#dependency-management)
    - [Import paths](#import-paths)

Please submit an [issue] on Github if you
* Notice a requirement that this doc does not capture.
* Find a different doc that specifies requirements (the doc should instead link
  here).

Development branch requirements will change over time, but release branch
requirements are frozen.

## Getting the code

Kubernetes is developed using git for version control.

### GitHub workflow

To check out code to work on, please refer to [this guide](/contributors/guide/github-workflow.md).

For tips, tricks, and common practices used within the Kubernetes project read the [Kubernetes GitHub Cheat Sheet][github-cheat-sheet].

## Building Kubernetes

There are two recommended ways of building Kubernetes, and the choice of one or the other depends on if you do plan on developing upstream. One way is building a release from source, this is recommended for development purposes. If there is no development involved, it is suggested to use a pre-built version of the current release.

### Building Kubernetes with Docker

Official releases are built using Docker containers. To build Kubernetes using
Docker please follow [these
instructions][these-instructions].

### Building Kubernetes on a local OS/shell environment

Kubernetes development helper scripts assume an up-to-date GNU tools
environment. Recent Linux distributions should work out-of-the-box.

macOS ships with outdated BSD-based tools. We recommend installing [macOS GNU
tools].

#### rsync

Kubernetes build system requires `rsync` command present in the development
platform.

#### etcd

etcd is a distributed key value store that provides a reliable way to store data across a cluster of machines. Kubernetes maintains state in [`etcd`][etcd-latest]. Version 2 as was deprecated and support was removed in Kubernetes 1.13; thereby, make sure to install etcd version 3.x.

Please [install it locally][etcd-install] to run local integration tests. Also, add it to PATH.

Options for local installation:
  1. Install inside kubernetes root. Use `hack/install-etcd.sh`
  2. Install manually.  
    Find version with `grep -E "image.*etcd" cluster/gce/manifests/etcd.manifest` and install with your OS package manager.

#### Go

Kubernetes is written in [Go](http://golang.org). If you don't have a Go
development environment, please [set one up](https://golang.org/doc/install).

Great introduction to the [development process in Go](http://golang.org/doc/code.html).


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
| 1.13           | 1.11.2      |
| 1.13+          | 1.11.4      |

Ensure your GOPATH and PATH have been configured in accordance with the Go
environment instructions.

##### Upgrading Go

Upgrading Go requires specific modification of some scripts and container
images.

- The image for cross compiling in [build/build-image/cross].
  The `VERSION` file and `Dockerfile`.
- Update the desired Go version in Dockerfile for the [e2e][e2e-image] and [test][test-image].
  This requires pushing the [e2e][e2e-image] and [test][test-image] images that are `FROM` the desired Go version.
- The cross tag `KUBE_BUILD_IMAGE_CROSS_TAG` in [build/common.sh].


##### Dependency management

Kubernetes uses [`godep`](https://github.com/tools/godep) to manage
dependencies.

Developers who need to manage dependencies in the `vendor/` tree should read
the docs on [using godep to manage dependencies](sig-architecture/godep.md).

##### Import paths

Kubernetes import paths are of the form `k8s.io/X`. Those paths map to `github.com/kubernetes/kubernetes/staging/X` whose code is placed in the staging area directory `staging/src/k8s.io`.

Aditionally, that code is found by the compiler because of the symlinks located at `vendor/k8s.io`. For more information,
see [staging/README.md].


### Build with Bazel/Gazel

Building with Bazel is currently experimental.  For more information,
see [Build with Bazel].


[macOS GNU tools]: https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x
[build/build-image/cross]: https://git.k8s.io/kubernetes/build/build-image/cross
[build/common.sh]: https://git.k8s.io/kubernetes/build/common.sh
[e2e-image]: https://git.k8s.io/test-infra/jenkins/e2e-image
[etcd-latest]: https://coreos.com/etcd/docs/latest
[etcd-install]: sig-testing/integration-tests.md#install-etcd-dependency
<!-- https://github.com/coreos/etcd/releases -->
[go-workspace]: https://golang.org/doc/code.html#Workspaces
[issue]: https://github.com/kubernetes/kubernetes/issues
[kubectl user guide]: https://kubernetes.io/docs/user-guide/kubectl
[kubernetes.io]: https://kubernetes.io
[mercurial]: http://mercurial.selenic.com/wiki/Download
[test-image]: https://git.k8s.io/test-infra/jenkins/test-image
[Build with Bazel]: sig-testing/bazel.md
[official git documentation]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
[github-workflow.md]: https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md
[staging/README.md]: https://git.k8s.io/kubernetes/staging/README.md
[these-instructions]: https://git.k8s.io/kubernetes/build/README.md

[github-cheat-sheet]: https://github.com/kubernetes/community/blob/master/devel/github-cheat-sheet.md