# Integration Testing in Kubernetes

**Table of Contents**

- [Integration testing in Kubernetes](#integration-tests)
  - [Install etcd dependency](#install-etcd-dependency)
  - [Etcd test data](#etcd-test-data)
  - [Run integration tests](#run-integration-tests)
  - [Run a specific integration test](#run-a-specific-integration-test)

This assumes you already read the [testing guide](testing.md).

## Integration tests

* Integration tests should only access other resources on the local machine
  - Most commonly etcd or a service listening on localhost.
* All significant features require integration tests.
  - This includes kubectl commands
* The preferred method of testing multiple scenarios or inputs
is [table driven testing](https://github.com/golang/go/wiki/TableDrivenTests)
  - Example: [TestNamespaceAuthorization](https://git.k8s.io/kubernetes/test/integration/auth/auth_test.go)
* Each test should create its own master, httpserver and config.
  - Example: [TestPodUpdateActiveDeadlineSeconds](https://git.k8s.io/kubernetes/test/integration/pods/pods_test.go)
* See [coding conventions](../../guide/coding-conventions.md).

### Install etcd dependency

Kubernetes integration tests require your `PATH` to include an
[etcd](https://github.com/coreos/etcd/releases) installation. Kubernetes
includes a script to help install etcd on your machine.

```sh
# Install etcd and add to PATH

# Option a) install inside kubernetes root
hack/install-etcd.sh  # Installs in ./third_party/etcd
echo export PATH="\$PATH:$(pwd)/third_party/etcd" >> ~/.profile  # Add to PATH

# Option b) install manually
grep -E "image.*etcd" cluster/gce/manifests/etcd.manifest  # Find version
# Install that version using yum/apt-get/etc
echo export PATH="\$PATH:<LOCATION>" >> ~/.profile  # Add to PATH
```

### Etcd test data

Many tests start an etcd server internally, storing test data in the operating system's temporary directory.

If you see test failures because the temporary directory does not have sufficient space,
or is on a volume with unpredictable write latency, you can override the test data directory
for those internal etcd instances with the `TEST_ETCD_DIR` environment variable.

### Run integration tests

The integration tests are run using `make test-integration`.
The Kubernetes integration tests are written using the normal golang testing
package but expect to have a running etcd instance to connect to.  The `test-integration.sh`
script wraps `make test` and sets up an etcd instance for the integration tests to use.

```sh
make test-integration  # Run all integration tests.
```

This script runs the golang tests in package
[`test/integration`](https://git.k8s.io/kubernetes/test/integration).

### Run a specific integration test

You can also use the `KUBE_TEST_ARGS` environment variable with the `make test-integration`
to run a specific integration test case:

```sh
# Run integration test TestPodUpdateActiveDeadlineSeconds with the verbose flag set.
make test-integration WHAT=./test/integration/pods GOFLAGS="-v" KUBE_TEST_ARGS="-run ^TestPodUpdateActiveDeadlineSeconds$"
```

If you set `KUBE_TEST_ARGS`, the test case will be run with only the `v1` API
version and the watch cache test is skipped.
