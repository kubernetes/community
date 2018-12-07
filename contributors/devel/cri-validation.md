# Container Runtime Interface (CRI) Validation Testing

CRI validation testing provides a test framework and a suite of tests to validate that the Container Runtime Interface (CRI) server implementation meets all the requirements. This allows the CRI runtime developers to verify that their runtime conforms to CRI, without needing to set up Kubernetes components or run Kubernetes end-to-end tests.

CRI validation testing is GA since v1.11.0 and is hosted at the [cri-tools](https://github.com/kubernetes-sigs/cri-tools) repository. We encourage the CRI developers to report bugs or help extend the test coverage by adding more tests.

## Install

The test suites can be downloaded from cri-tools [release page](https://github.com/kubernetes-sigs/cri-tools/releases):

```sh
VERSION="v1.11.0"
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/$VERSION/critest-$VERSION-linux-amd64.tar.gz
sudo tar zxvf critest-$VERSION-linux-amd64.tar.gz -C /usr/local/bin
rm -f critest-$VERSION-linux-amd64.tar.gz
```

critest requires [ginkgo](https://github.com/onsi/ginkgo) to run parallel tests. It could be installed by

```sh
go get -u github.com/onsi/ginkgo/ginkgo
```

*Note: ensure GO is installed and GOPATH is set before installing ginkgo.*

## Running tests

### Prerequisite

Before running the test, you need to _ensure that the CRI server under test is running and listening on a Unix socket_. Because the validation tests are designed to request changes (e.g., create/delete) to the containers and verify that correct status is reported, it expects to be the only user of the CRI server. Please make sure that 1) there are no existing CRI-managed containers running on the node, and 2) no other processes (e.g., Kubelet) will interfere with the tests.

### Run

```sh
critest
```

This will

- Connect to the shim of CRI container runtime
- Run the tests using `ginkgo`
- Output the test results to STDOUT

critest connects to `unix:///var/run/dockershim.sock` by default. For other runtimes, the endpoint can be set by flags `-runtime-endpoint` and `-image-endpoint`.

## Additional options

- `-ginkgo.focus`: Only run the tests that match the regular expression.
- `-image-endpoint`: Set the endpoint of image service. Same with runtime-endpoint if not specified.
- `-runtime-endpoint`: Set the endpoint of runtime service. Default to `unix:///var/run/dockershim.sock`.
- `-ginkgo.skip`: Skip the tests that match the regular expression.
- `-parallel`: The number of parallel test nodes to run (default 1). ginkgo must be installed to run parallel tests.
- `-h`: Show help and all supported options.
