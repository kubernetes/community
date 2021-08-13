# Getting Kubernetes Builds

- [Version markers](#version-markers)
- [Using `hack/get-build.sh`](#using-hackget-buildsh)

## Version markers

Version markers are text files which act as sort of a public API for accessing
Kubernetes builds.

They are artifacts of a successful Kubernetes build which are stored in a
Google Cloud Storage bucket alongside the builds they make reference to.

You can read more about version markers [here](./kubernetes-versions.md).

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
