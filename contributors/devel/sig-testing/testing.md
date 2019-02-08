# Testing guide

**Table of Contents**

- [Testing guide](#testing-guide)
  - [Unit tests](#unit-tests)
    - [Run all unit tests](#run-all-unit-tests)
    - [Set go flags during unit tests](#set-go-flags-during-unit-tests)
    - [Run unit tests from certain packages](#run-unit-tests-from-certain-packages)
    - [Run specific unit test cases in a package](#run-specific-unit-test-cases-in-a-package)
    - [Stress running unit tests](#stress-running-unit-tests)
    - [Unit test coverage](#unit-test-coverage)
    - [Benchmark unit tests](#benchmark-unit-tests)
  - [Integration tests](#integration-tests)
  - [End-to-End tests](#end-to-end-tests)


This assumes you already read the [development guide](../development.md) to
install go, godeps, and configure your git client.  All command examples are
relative to the `kubernetes` root directory.

Before sending pull requests you should at least make sure your changes have
passed both unit and integration tests.

Kubernetes only merges pull requests when unit, integration, and e2e tests are
passing, so it is often a good idea to make sure the e2e tests work as well.

## Unit tests

* Unit tests should be fully hermetic
  - Only access resources in the test binary.
* All packages and any significant files require unit tests.
* The preferred method of testing multiple scenarios or input is
  [table driven testing](https://github.com/golang/go/wiki/TableDrivenTests)
  - Example: [TestNamespaceAuthorization](https://git.k8s.io/kubernetes/test/integration/auth/auth_test.go)
* Unit tests must pass on macOS and Windows platforms.
  - Tests using linux-specific features must be skipped or compiled out.
  - Skipped is better, compiled out is required when it won't compile.
* Concurrent unit test runs must pass.
* See [coding conventions](../../guide/coding-conventions.md).

### Run all unit tests

`make test` is the entrypoint for running the unit tests that ensures that
`GOPATH` is set up correctly.  If you have `GOPATH` set up correctly, you can
also just use `go test` directly.

```sh
cd kubernetes
make test  # Run all unit tests.
```

If any unit test fails with a timeout panic (see [#1594](https://github.com/kubernetes/community/issues/1594)) on the testing package, you can increase the `KUBE_TIMEOUT` value as shown below.

```sh
make test KUBE_TIMEOUT="-timeout 300s"
```

### Set go flags during unit tests

You can set [go flags](https://golang.org/cmd/go/) by setting the
`GOFLAGS` environment variable.

### Run unit tests from certain packages

`make test` accepts packages as arguments; the `k8s.io/kubernetes` prefix is
added automatically to these:

```sh
make test WHAT=./pkg/api                # run tests for pkg/api
```

To run multiple targets you need quotes:

```sh
make test WHAT="./pkg/api ./pkg/kubelet"  # run tests for pkg/api and pkg/kubelet
```

In a shell, it's often handy to use brace expansion:

```sh
make test WHAT=./pkg/{api,kubelet}  # run tests for pkg/api and pkg/kubelet
```

### Run specific unit test cases in a package

You can set the test args using the `KUBE_TEST_ARGS` environment variable.
You can use this to pass the `-run` argument to `go test`, which accepts a
regular expression for the name of the test that should be run.

```sh
# Runs TestValidatePod in pkg/api/validation with the verbose flag set
make test WHAT=./pkg/api/validation GOFLAGS="-v" KUBE_TEST_ARGS='-run ^TestValidatePod$'

# Runs tests that match the regex ValidatePod|ValidateConfigMap in pkg/api/validation
make test WHAT=./pkg/api/validation GOFLAGS="-v" KUBE_TEST_ARGS="-run ValidatePod\|ValidateConfigMap$"
```

For other supported test flags, see the [golang
documentation](https://golang.org/cmd/go/#hdr-Testing_flags).

### Stress running unit tests

Running the same tests repeatedly is one way to root out flakes.
You can do this efficiently.

```sh
# Have 2 workers run all tests 5 times each (10 total iterations).
make test PARALLEL=2 ITERATION=5
```

For more advanced ideas please see [flaky-tests.md](flaky-tests.md).

### Unit test coverage

Currently, collecting coverage is only supported for the Go unit tests.

To run all unit tests and generate an HTML coverage report, run the following:

```sh
make test KUBE_COVER=y
```

At the end of the run, an HTML report will be generated with the path
printed to stdout.

To run tests and collect coverage in only one package, pass its relative path
under the `kubernetes` directory as an argument, for example:

```sh
make test WHAT=./pkg/kubectl KUBE_COVER=y
```

Multiple arguments can be passed, in which case the coverage results will be
combined for all tests run.

### Benchmark unit tests

To run benchmark tests, you'll typically use something like:

```sh
go test ./pkg/apiserver -benchmem -run=XXX -bench=BenchmarkWatch
```

This will do the following:

1. `-run=XXX` is a regular expression filter on the name of test cases to run
2. `-bench=BenchmarkWatch` will run test methods with BenchmarkWatch in the name
  * See `grep -nr BenchmarkWatch .` for examples
3. `-benchmem` enables memory allocation stats

See `go help test` and `go help testflag` for additional info.

## Integration tests

Please refer to [Integration Testing in Kubernetes](integration-tests.md).

## End-to-End tests

Please refer to [End-to-End Testing in Kubernetes](e2e-tests.md).
