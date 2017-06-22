# Container Runtime Interface (CRI) Validation Testing

CRI validation testing provides a test framework and a suite of tests to validate that the Container Runtime Interface (CRI) server implementation meets all the requirements. This allows the CRI runtime developers to verify that their runtime conforms to CRI, without needing to set up Kubernetes components or run Kubernetes end-to-end tests.

CRI validation testing is currently Alpha and is hosted at the [cri-tools](https://github.com/kubernetes-incubator/cri-tools) repository. Performance benchmarking will be added in the future. We encourage the CRI developers to report bugs or help extend the test coverage by adding more tests.

## Install

The test suites can be installed easily via `go get` command:

```sh
go get github.com/kubernetes-incubator/cri-tools/cmd/critest
```

Then `critest` binary can be found in `$GOPATH/bin`.

*Note: ensure GO is installed and GOPATH is set before installing critest.*

## Running tests

### Prerequisite

Before running the test, you need to _ensure that the CRI server under test is running and listening on a Unix socket_. Because the validation tests are designed to request changes (e.g., create/delete) to the containers and verify that correct status is reported, it expects to be the only user of the CRI server. Please make sure that 1) there are no existing CRI-managed containers running on the node, and 2) no other processes (e.g., Kubelet) will interfere with the tests.

### Run

```sh
critest validation
```

This will

- Connect to the shim of CRI container runtime
- Run the tests using `ginkgo`
- Output the test results to STDOUT

critest connects to `/var/run/dockershim.sock` by default. For other runtimes, the endpoint can be set in two ways:

- By setting flags `--runtime-endpoint` and `--image-endpoint`
- By setting environment variables `CRI_RUNTIME_ENDPOINT` and `CRI_IMAGE_ENDPOINT`

## Additional options

- `--focus`, `-f`: Only run the tests that match the regular expression.
- -`-ginkgo-flags`, `-g`: Space-separated list of arguments to pass to Ginkgo test runner.
- `--image-endpoint`, `-i`: Set the endpoint of image service. Same with runtime-endpoint if not specified.
- `--runtime-endpoint`, `-r`: Set the endpoint of runtime service. Default to `/var/run/dockershim.sock`.
- `--skip`, `-s`: Skip the tests that match the regular expression.
