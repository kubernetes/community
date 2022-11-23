# Kubernetes Version Markers

Version markers are text files which act as sort of a public API for accessing
Kubernetes builds.

They are artifacts of a successful Kubernetes build which are stored in a
Google Cloud Storage bucket alongside the builds they make reference to.

You may see version markers leveraged in a variety of places, including:

- extraction strategies for e2e tests
- Release Engineering tooling
- external user-maintained scripts

## Table of Contents <!-- omit in toc -->

- [tl;dr](#tldr)
  - [CI - cross build](#ci---cross-build)
  - [CI - linux/amd64 (fast) build](#ci---linuxamd64-fast-build)
  - [Release - Official release build](#release---official-release-build)
  - [Release - Pre-release build](#release---pre-release-build)
- [Usage](#usage)
  - [Format](#format)
  - [Access](#access)
    - [gsutil](#gsutil)
    - [curl/wget](#curlwget)
    - [gcsweb](#gcsweb)
  - [Querying a build](#querying-a-build)
- [Marker types](#marker-types)
  - [CI](#ci)
    - [latest](#latest)
    - [latest-fast](#latest-fast)
    - [**DEPRECATED** - generic](#deprecated---generic)
  - [Release](#release)
    - [Official](#official)
    - [Pre-release](#pre-release)
- [Future Plans](#future-plans)
- [Known Issues](#known-issues)
  - [Generic version markers are not explicit](#generic-version-markers-are-not-explicit)
  - [Manually created jobs using generic version markers can be inaccurate](#manually-created-jobs-using-generic-version-markers-can-be-inaccurate)
- [Previous Issues](#previous-issues)
  - [linux/amd64 version markers are colliding with cross builds](#linuxamd64-version-markers-are-colliding-with-cross-builds)
  - [Cross builds are stored in a separate GCS bucket](#cross-builds-are-stored-in-a-separate-gcs-bucket)
  - [Generated jobs may not represent intention](#generated-jobs-may-not-represent-intention)
  - [bazel version markers appear to be unused](#bazel-version-markers-appear-to-be-unused)

## tl;dr

You need a...

### CI - cross build version

Use `curl -Ls https://dl.k8s.io/ci/latest.txt` (`master` branch)

OR

`curl -Ls https://dl.k8s.io/ci/latest-x.y.txt`, where `x` is the
Kubernetes major version and `y` is the Kubernetes minor version (release branches).

### CI - linux/amd64 (fast) build version

Use `curl -Ls https://dl.k8s.io/ci/latest-fast.txt` (**_only available
on `master` branch_**).

### Release - Official release build version

Use `curl -Ls https://dl.k8s.io/release/stable-x.y.txt`, where `x` is the
Kubernetes major version and `y` is the Kubernetes minor version.

### Release - Pre-release build version

Use `curl -Ls https://dl.k8s.io/release/latest-x.y.txt`, where `x` is the
Kubernetes major version and `y` is the Kubernetes minor version.

## Usage

### Format

All version markers have similar endpoints:

```console
<base-uri>/<directory>/<marker>
```

Expected output is [semver](https://semver.org/spec/v2.0.0.html)-compliant
version, prepended with a `v`.

Example:

```console
v1.20.0-alpha.0.391+575c4925be8c39
```

### Access

Version markers are accessible via HTTP, so there are several ways to get them,
depending on your use case.

#### curl/wget

The pattern to use is:

```
https://dl.k8s.io/<directory>/<marker>
```

For example, this generates a list of copy-pastable curl commands for each marker:
```shell
for p in ci release; do \
  for v in "" -1 -1.21 -1.20 -1.19; do \
    for t in latest stable; do \
      m=$p/$t$v; echo curl -Ls https://dl.k8s.io/$m.txt: $(curl -Ls https://dl.k8s.io/$m.txt); \
    done; 
  done; 
done | grep -v NoSuchKey
```

```
curl -Ls https://dl.k8s.io/ci/latest.txt: v1.22.0-beta.2.39+33aba7ee025dfd
curl -Ls https://dl.k8s.io/ci/latest-1.txt: v1.22.0-beta.2.39+33aba7ee025dfd
curl -Ls https://dl.k8s.io/ci/latest-1.21.txt: v1.21.4-rc.0.8+e33079364f8f66
curl -Ls https://dl.k8s.io/ci/latest-1.20.txt: v1.20.10-rc.0.3+c0f757a067f32a
curl -Ls https://dl.k8s.io/ci/latest-1.19.txt: v1.19.14-rc.0.1+4900b04db38214
curl -Ls https://dl.k8s.io/release/latest.txt: v1.22.0-beta.2
curl -Ls https://dl.k8s.io/release/stable.txt: v1.21.3
curl -Ls https://dl.k8s.io/release/latest-1.txt: v1.22.0-beta.2
curl -Ls https://dl.k8s.io/release/stable-1.txt: v1.21.3
curl -Ls https://dl.k8s.io/release/latest-1.21.txt: v1.21.4-rc.0
curl -Ls https://dl.k8s.io/release/stable-1.21.txt: v1.21.3
curl -Ls https://dl.k8s.io/release/latest-1.20.txt: v1.20.10-rc.0
curl -Ls https://dl.k8s.io/release/stable-1.20.txt: v1.20.9
curl -Ls https://dl.k8s.io/release/latest-1.19.txt: v1.19.14-rc.0
curl -Ls https://dl.k8s.io/release/stable-1.19.txt: v1.19.13
```

#### gcsweb

Navigate via web browser to
`https://gcsweb.k8s.io/gcs/<gcs-bucket>/<directory>/<marker>`.

For CI builds, `<gcs-bucket>` is `k8s-release-dev`

Example:

```console
$ open https://gcsweb.k8s.io/gcs/k8s-release-dev/ci
```

For release builds, `<gcs-bucket>` is `kubernetes-release`

Example:

```console
$ open https://gcsweb.k8s.io/gcs/kubernetes-release/release
```

### Querying a build

Once you've successfully retrieved a version marker, you can use it to list the
contents of a Kubernetes build.

Here's an example:

```console
$ gsutil ls gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/
```

Output:

<details>

```
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA256SUMS
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA256SUMS.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA256SUMS.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA512SUMS
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA512SUMS.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/SHA512SUMS.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-darwin-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-386.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-386.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-386.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-ppc64le.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-ppc64le.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-ppc64le.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-s390x.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-s390x.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-linux-s390x.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-386.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-386.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-386.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-client-windows-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-manifests.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-manifests.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-manifests.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-ppc64le.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-ppc64le.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-ppc64le.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-s390x.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-s390x.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-linux-s390x.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-windows-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-windows-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-node-windows-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-ppc64le.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-ppc64le.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-ppc64le.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-s390x.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-s390x.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-server-linux-s390x.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-src.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-src.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-src.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-darwin-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-arm64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-ppc64le.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-ppc64le.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-ppc64le.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-s390x.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-s390x.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-linux-s390x.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-portable.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-portable.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-portable.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-windows-amd64.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-windows-amd64.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes-test-windows-amd64.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes.tar.gz
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes.tar.gz.sha256
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/kubernetes.tar.gz.sha512
gs://k8s-release-dev/ci/v1.22.0-beta.2.39+33aba7ee025dfd/bin/
```

</details>

A general, non-exhaustive list of the expected artifacts of a Kubernetes build
can be found [here](https://git.k8s.io/sig-release/release-engineering/artifacts.md).

## Marker types

Version markers are broken into two primary types: `CI` and `release` markers

### CI

**GCS Bucket:** `k8s-release-dev`

**Directory:** `ci`

#### latest

**Path:** `https://dl.k8s.io/ci/latest[-x.y].txt`

**kubekins-e2e `extract`:** `ci/latest[-x.y].txt`

`latest` markers reference cross builds generated via the `ci-kubernetes-build`
jobs, which run approximately every hour.

This version marker exists for all active Kubernetes release branches.

#### latest-fast

**Path:** `https://dl.k8s.io/ci/latest-fast.txt`

**kubekins-e2e `extract`:** `ci/latest-fast.txt`

`latest-fast` markers reference linux/amd64-only builds generated via the
`ci-kubernetes-build-fast` job, which run approximately every 5 minutes.

This version marker only exists for the `master` branch.

#### **DEPRECATED** - generic

**Path:** `https://dl.k8s.io/ci/k8s-<generic-version>.txt`

**kubekins-e2e `extract`:** `ci/k8s-<generic-version>.txt`

The following generic markers are available:

- `k8s-master`
- `k8s-beta`*
- `k8s-stable1`
- `k8s-stable2`
- `k8s-stable3`
- `k8s-stable4`

Generic markers reference cross builds generated via the `ci-kubernetes-build`
jobs, which run approximately every hour.

This version marker exists for all active Kubernetes release branches.

**Whenever possible, prefer using the latest version markers instead of
generic ones. The meaning of these markers changes throughout the release
cycles, which frequently leads to job misconfiguration and increased
difficulty in debugging failures. These generic markers only continue to exist
to prevent breakage in jobs that currently use them. They will be disabled at a
future date.**

### Release

**GCS Bucket:** `kubernetes-release`

**Directory:** `release`

#### Official

**Path:** `https://dl.k8s.io/release/stable-x.y.txt`

**kubekins-e2e `extract`:** `release/stable-x.y.txt`

`stable` markers reference cross builds generated via the official Kubernetes
releases and can only be produced by [Release Managers][release-managers].

This version marker exists for all active Kubernetes release branches that have
had a minor (`x.y.0`) release.

#### Pre-release

**Path:** `https://dl.k8s.io/release/latest-x.y.txt`

**kubekins-e2e `extract`:** `release/latest-x.y.txt`

`latest` markers reference cross builds generated via the Kubernetes
pre-releases and can only be produced by [Release Managers][release-managers].

This version marker exists for all active Kubernetes release branches.

## Future Plans

(An up-to-date tracking issue for this work can be found [here](https://github.com/kubernetes/sig-release/issues/850).)

- [x] (https://github.com/kubernetes/test-infra/pull/15564) Use explicit (`latest-x.y`) version markers in generated jobs
- [ ] (https://github.com/kubernetes/test-infra/pull/18290, https://github.com/kubernetes/release/pull/1389) Publish fast builds to separate subdirectory to prevent collisions
- [ ] (https://github.com/kubernetes/test-infra/pull/18169) Add a new `stable4` field to the kubernetes version lists in [`releng/test_config.yaml`][test_config.yaml] to remove confusion around jobs with `beta` in their name
- [ ] Refactor any non-generated jobs using generic version markers
- [ ] Refactor any config-forked release jobs using generic version markers
- [ ] Refactor any jobs using `fork-per-release-generic-suffix: "true"` annotations
- [ ] Disable usage of `fork-per-release-generic-suffix: "true"` annotations
- [ ] Rewrite the [Kubernetes versions doc](https://github.com/kubernetes/test-infra/blob/master/docs/kubernetes-versions.md) and put it in a more visible location
- [ ] Refactor [`releng/test_config.yaml`][test_config.yaml] to remove references to generic versions (e.g., prefer `ci-kubernetes-e2enode-ubuntu1-latest-1-19-gkespec` over `ci-kubernetes-e2enode-ubuntu1-k8sbeta-gkespec`)

## Known Issues

Unfortunately, the way certain version markers are generated and utilized can
at best be confusing, and at worst, disruptive.

There are a variety of problems, some of which are symptoms of the other ones...

### Generic version markers are not explicit

We publish a set of additional generic version markers:

- `k8s-master`
- `k8s-beta`
- `k8s-stable1`
- `k8s-stable2`
- `k8s-stable3`
- `k8s-stable4`

Depending on the point in the release cycle, the meaning of these markers can
change.

- `k8s-master` always points to the version on `master`.
- `k8s-beta` may represent:
  - `master`s build version (pre-branch cut)
  - a to-be-released build version (post-branch cut)
  - a recently released build version (post-release)

Knowing what these markers mean at any one time presumes knowledge of the
build/release process or a correct interpretation of this document, which has
frequently out of date and lives in a low-visibility location.

### Manually created jobs using generic version markers can be inaccurate

Non-generated jobs using generic version markers do not get the same level of
scrutiny as ones that are generated via
[`releng/test_config.yaml`][test_config.yaml].

This leads to inaccuracies between the versions presumed to be used in test
and the versions that may be displayed in testgrid.

`ci-kubernetes-e2e-gce-beta-stable1-gci-kubectl-skew` is a great example:

https://github.com/kubernetes/test-infra/blob/96e08f4be2a86189f59c72055785f817ac346d30/config/jobs/kubernetes/sig-cli/sig-cli-config.yaml#L85-L112

All variants of that prowjob have landed on the `sig-release-job-config-errors`
dashboard for various misconfiguration issues that are the result of generic
version markers.

## Previous Issues

### linux/amd64 version markers are colliding with cross builds

(Fixed in https://github.com/kubernetes/test-infra/pull/18290.)

"Fast" (linux/amd64-only) builds run every 5 minutes, while cross builds run
every hour.
They also write to the same version markers (`latest.txt`,
`latest-<major>.txt`, `latest-<major>.<minor>.txt`).

The Kubernetes build jobs have a mechanism for checking if a build already
exists and will exit early to save on test cycles.

What this means is if a "fast" build has already happened for a commit, then
the corresponding cross build will exit without building.

This has been happening pretty consistently lately, so cross build consumers
are using much older versions of Kubernetes than intended.

(Note that this condition only happens on `master`.)

### Cross builds are stored in a separate GCS bucket

(Fixed in https://github.com/kubernetes/test-infra/pull/14030.)

This makes long-term usage of cross builds a little more difficult, since
scripts utilizing version markers tend to consider only the version marker
filename, while the GCS bucket name remains unparameterized.

### Generated jobs may not represent intention

(Fixed in https://github.com/kubernetes/test-infra/pull/15564.)

As the generic version markers can shift throughout the release cycle, every
time we regenerate jobs, they may not represent what we intend to test.

The best examples of this are pretty much every job using the `k8s-beta`
version marker, and more specifically, skew and upgrade jobs.

### bazel version markers appear to be unused

(Fixed in https://github.com/kubernetes/test-infra/pull/15612.)


[release-managers]: https://kubernetes.io/releases/release-managers/
[test_config.yaml]: https://github.com/kubernetes/test-infra/blob/master/releng/test_config.yaml
