# Build with Bazel

Building with Bazel is currently experimental. Automanaged `BUILD` rules have the
tag "automanaged" and are maintained by
[gazel](https://github.com/mikedanese/gazel). Instructions for installing Bazel
can be found [here](https://www.bazel.io/versions/master/docs/install.html).

Several `make` rules have been created for common operations:

* `make bazel-build`: builds all binaries in tree
* `make bazel-test`: runs all unit tests
* `make bazel-test-integration`: runs all integration tests
* `make bazel-release`: builds release tarballs, Docker images (for server
  components), and Debian images

You can also interact with Bazel directly; for example, to run all `kubectl` unit
tests, run

```console
$ bazel test //pkg/kubectl/...
```

## Continuous Integration

The [Bazel CI job](http://k8s-testgrid.appspot.com/google-unit#bazel) runs
`make bazel-build`, `make bazel-test`, and (transitively) `make bazel-release`.
A similar job is run on all PRs.

Many steps are cached, so the Bazel job usually executes fairly quickly.

## Known issues

[Cross-compilation is not currently supported](https://github.com/bazelbuild/rules_go/issues/70),
so all binaries will be built for the host architecture running Bazel.
Additionally, Go build tags are not supported. This means that builds on macOS may not work.

[Binaries produced by Bazel are not statically linked](https://github.com/bazelbuild/rules_go/issues/161),
and they are not currently tagged with version information.

[Bazel does not validate build environment](https://github.com/kubernetes/kubernetes/issues/51623), thus make sure that needed
tools and development packages are installed in the system. Bazel builds require presense of `make`, `gcc`, `g++`, `glibc and libstdc++ development headers` and `glibc static development libraries`. Please check your distribution for exact names of the packages. Examples for some commonly used distributions are below:

|     Dependency        | Debian/Ubuntu                 | CentOS                         | OpenSuSE                                |
|:---------------------:|-------------------------------|--------------------------------|-----------------------------------------|
| Build essentials      | `apt install build-essential` | `yum groupinstall development` | `zypper install -t pattern devel_C_C++` |
| GCC C++               | `apt install g++`             | `yum install gcc-c++`          | `zypper install gcc-c++`                |
| GNU Libc static files | `apt install libc6-dev`       | `yum install glibc-static`     | `zypper install glibc-devel-static`     |


## Updating `BUILD` files

To update `BUILD` files, run:

```console
$ ./hack/update-bazel.sh
```

**NOTE**: `update-bazel.sh` only works if check out directory of Kubernetes is `$GOPATH/src/k8s.io/kubernetes`.

Only rules which are automanaged will be updated, but all rules will be
auto-formatted.

Updating the `BUILD` file for a package will be required when:
* Files are added to or removed from a package
* Import dependencies change for a package
* A `BUILD` file has been updated and needs to be reformatted
* A new `BUILD` file has been added (parent `BUILD` files will be updated)

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/bazel.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
