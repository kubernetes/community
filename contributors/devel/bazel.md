# Build and test with Bazel

Building and testing Kubernetes with  Bazel is supported but not yet default.

Go rules are managed by the [`gazelle`](https://github.com/bazelbuild/rules_go/tree/master/go/tools/gazelle)
tool, with some additional rules managed by the [`kazel`](https://git.k8s.io/repo-infra/kazel) tool.
These tools are called via the `hack/update-bazel.sh` script.

Instructions for installing Bazel
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

## Planter
If you don't want to install Bazel, you can instead try using the unofficial
[Planter](https://git.k8s.io/test-infra/planter) tool,
which runs Bazel inside a Docker container.

For example, you can run
```console
$ ../test-infra/planter/planter.sh make bazel-test
$ ../test-infra/planter/planter.sh bazel build //cmd/kubectl
```

## Continuous Integration

There are several bazel CI jobs:
* [ci-kubernetes-bazel-build](http://k8s-testgrid.appspot.com/google-unit#bazel-build): builds everything
  with Bazel
* [ci-kubernetes-bazel-test](http://k8s-testgrid.appspot.com/google-unit#bazel-test): runs unit tests in
  with Bazel

Similar jobs are run on all PRs; additionally, several of the e2e jobs use
Bazel-built binaries when launching and testing Kubernetes clusters.

## Known issues

[Cross-compilation is not currently supported](https://github.com/bazelbuild/rules_go/issues/70),
so all binaries will be built for the host OS and architecture running Bazel.
(For example, you can't currently target linux/amd64 from macOS or linux/s390x
from an amd64 machine.)

Additionally, native macOS support is still a work in progress. Using Planter is
a possible workaround in the interim.

[Bazel does not validate build environment](https://github.com/kubernetes/kubernetes/issues/51623), thus make sure that needed
tools and development packages are installed in the system. Bazel builds require presence of `make`, `gcc`, `g++`, `glibc and libstdc++ development headers`, 
`glibc static development libraries` and `kernel development libraries`. Please check your distribution for exact names of the packages. Examples for some 
commonly used distributions are below:

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

To prevent Go rules from being updated, consult the [gazelle
documentation](https://github.com/bazelbuild/rules_go/tree/master/go/tools/gazelle).

Note that much like Go files and `gofmt`, BUILD files have standardized,
opinionated style rules, and running `hack/update-bazel.sh` will format them for you.

If you want to auto-format BUILD files in your editor, using something like
[Buildifier](https://github.com/bazelbuild/buildtools/blob/master/buildifier/README.md)
is recommended.

Updating the `BUILD` file for a package will be required when:
* Files are added to or removed from a package
* Import dependencies change for a package
* A `BUILD` file has been updated and needs to be reformatted
* A new `BUILD` file has been added (parent `BUILD` files will be updated)

## Contacts
For help or discussion, join the [#bazel](https://kubernetes.slack.com/messages/bazel)
channel on Kubernetes Slack.
