# Getting Kubernetes Builds

- [Using `hack/get-build.sh`](#using-hackget-buildsh)
- [Using `gsutil`](#using-gsutil)
  - [Install `gsutil`](#install-gsutil)
  - [Examples](#examples)
    - [Output the latest CI version number](#output-the-latest-ci-version-number)
    - [List the contents of a CI release](#list-the-contents-of-a-ci-release)
    - [List all official releases and RCs](#list-all-official-releases-and-rcs)

## Using `hack/get-build.sh`

You can use [hack/get-build.sh](https://git.k8s.io/kubernetes/hack/get-build.sh)
to get a build or to use as a reference on how to get the most recent builds
with curl.

With `get-build.sh` you can grab the most recent stable build, the
most recent release candidate, or the most recent build to pass our CI and GCE
e2e tests (essentially a nightly build).

Run `./hack/get-build.sh -h` for its usage.

To get a build at a specific version (v1.18.3) use:

```shell
./hack/get-build.sh v1.18.3
```

To get the latest stable release:

```shell
./hack/get-build.sh release/stable
```

Use the "-v" option to print the version number of a build without retrieving
it.

For example, the following prints the version number for the latest ci
build:

```shell
./hack/get-build.sh -v ci/latest
```

## Using `gsutil`

You can also use the gsutil tool to explore the Google Cloud Storage release
buckets.

### Install `gsutil`

`gsutil` is part of the Google Cloud SDK.
Install instructions for the Google Cloud SDK can be found [here](https://cloud.google.com/sdk/install).

### Examples

#### Output the latest CI version number

```console
$ gsutil cat gs://kubernetes-release-dev/ci/latest.txt
v1.19.0-beta.2.32+35fc65dc2c614e
```

#### List the contents of a CI release

```console
$ gsutil ls gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA256SUMS
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA256SUMS.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA256SUMS.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA512SUMS
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA512SUMS.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/SHA512SUMS.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-client-linux-amd64.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-client-linux-amd64.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-client-linux-amd64.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-manifests.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-manifests.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-manifests.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-node-linux-amd64.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-node-linux-amd64.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-node-linux-amd64.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-server-linux-amd64.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-server-linux-amd64.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-server-linux-amd64.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-src.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-src.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-src.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-linux-amd64.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-linux-amd64.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-linux-amd64.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-portable.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-portable.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes-test-portable.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes.tar.gz
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes.tar.gz.sha256
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/kubernetes.tar.gz.sha512
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/bin/
gs://kubernetes-release-dev/ci/v1.19.0-beta.2.32+35fc65dc2c614e/extra/
```

#### List all official releases and RCs

```console
$ gsutil ls gs://kubernetes-release/release
gs://kubernetes-release/release/kube-register
gs://kubernetes-release/release/latest-1.0.txt
gs://kubernetes-release/release/latest-1.1.txt
gs://kubernetes-release/release/latest-1.10.txt
gs://kubernetes-release/release/latest-1.11.txt
gs://kubernetes-release/release/latest-1.12.txt
gs://kubernetes-release/release/latest-1.13.txt
gs://kubernetes-release/release/latest-1.14.txt
gs://kubernetes-release/release/latest-1.15.txt
gs://kubernetes-release/release/latest-1.16.txt
gs://kubernetes-release/release/latest-1.17.txt
gs://kubernetes-release/release/latest-1.18.txt
gs://kubernetes-release/release/latest-1.19.txt
gs://kubernetes-release/release/latest-1.2.txt
gs://kubernetes-release/release/latest-1.3.txt
gs://kubernetes-release/release/latest-1.4.txt
gs://kubernetes-release/release/latest-1.5.txt
gs://kubernetes-release/release/latest-1.6.txt
gs://kubernetes-release/release/latest-1.7.txt
gs://kubernetes-release/release/latest-1.8.txt
gs://kubernetes-release/release/latest-1.9.txt
gs://kubernetes-release/release/latest-1.txt
gs://kubernetes-release/release/latest.txt
gs://kubernetes-release/release/stable-1.0.txt
gs://kubernetes-release/release/stable-1.1.txt
gs://kubernetes-release/release/stable-1.10.txt
gs://kubernetes-release/release/stable-1.11.txt
gs://kubernetes-release/release/stable-1.12.txt
gs://kubernetes-release/release/stable-1.13.txt
gs://kubernetes-release/release/stable-1.14.txt
gs://kubernetes-release/release/stable-1.15.txt
gs://kubernetes-release/release/stable-1.16.txt
gs://kubernetes-release/release/stable-1.17.txt
gs://kubernetes-release/release/stable-1.18.txt
gs://kubernetes-release/release/stable-1.2.txt
gs://kubernetes-release/release/stable-1.3.txt
gs://kubernetes-release/release/stable-1.4.txt
gs://kubernetes-release/release/stable-1.5.txt
gs://kubernetes-release/release/stable-1.6.txt
gs://kubernetes-release/release/stable-1.7.txt
gs://kubernetes-release/release/stable-1.8.txt
gs://kubernetes-release/release/stable-1.9.txt
gs://kubernetes-release/release/stable-1.txt
gs://kubernetes-release/release/stable.txt

<snip>

gs://kubernetes-release/release/v1.18.0-alpha.0/
gs://kubernetes-release/release/v1.18.0-alpha.1/
gs://kubernetes-release/release/v1.18.0-alpha.2/
gs://kubernetes-release/release/v1.18.0-alpha.3/
gs://kubernetes-release/release/v1.18.0-alpha.5/
gs://kubernetes-release/release/v1.18.0-beta.0/
gs://kubernetes-release/release/v1.18.0-beta.1/
gs://kubernetes-release/release/v1.18.0-beta.2/
gs://kubernetes-release/release/v1.18.0-rc.1/
gs://kubernetes-release/release/v1.18.0/
gs://kubernetes-release/release/v1.18.1-beta.0/
gs://kubernetes-release/release/v1.18.1/
gs://kubernetes-release/release/v1.18.2-beta.0/
gs://kubernetes-release/release/v1.18.2/
gs://kubernetes-release/release/v1.18.3-beta.0/
gs://kubernetes-release/release/v1.18.3/
gs://kubernetes-release/release/v1.18.4-rc.0/
gs://kubernetes-release/release/v1.19.0-alpha.0/
gs://kubernetes-release/release/v1.19.0-alpha.1/
gs://kubernetes-release/release/v1.19.0-alpha.2/
gs://kubernetes-release/release/v1.19.0-alpha.3/
gs://kubernetes-release/release/v1.19.0-beta.0/
gs://kubernetes-release/release/v1.19.0-beta.1/
gs://kubernetes-release/release/v1.19.0-beta.2/

<snip>
```
