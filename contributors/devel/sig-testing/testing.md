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
install go and configure your git client.  All command examples are
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
`GOPATH` is set up correctly.

```sh
cd kubernetes
make test  # Run all unit tests.
```
If you have `GOPATH` set up correctly, you can
also just use `go test` directly.

```sh
cd kubernetes
go test ./... # Run all unit tests
```

The remainder of this documentation presumes that you use `Make` as an
entry point, but remember that the ability to use `go test` exists should you
desire.

If any unit test fails with a timeout panic (see [#1594](https://github.com/kubernetes/community/issues/1594)) on the testing package, you can increase the `KUBE_TIMEOUT` value as shown below.

```sh
make test KUBE_TIMEOUT="-timeout=300s"
```

### Set go flags during unit tests

You can set [go flags](https://golang.org/cmd/go/) by setting the
`GOFLAGS` environment variable.

### Run unit tests from certain packages

`make test` accepts packages as arguments; the `k8s.io/kubernetes` prefix is
added automatically to these:

```sh
make test WHAT=./pkg/kubelet                # run tests for pkg/kubelet
```

Expressed strictly with `go test`, the above command is equivalent to the following:

```sh
go test ./pkg/kubelet
```

To run tests for a package and all of its subpackages, you need to append `...`
to the package path:

```sh
make test WHAT=./pkg/api/...  # run tests for pkg/api and all its subpackages
```

To run multiple targets you need quotes:

```sh
make test WHAT="./pkg/kubelet ./pkg/scheduler"  # run tests for pkg/kubelet and pkg/scheduler
```

In a shell, it's often handy to use brace expansion:

```sh
make test WHAT=./pkg/{kubelet,scheduler}  # run tests for pkg/kubelet and pkg/scheduler
```

### Run specific unit test cases in a package

You can set the test args using the `KUBE_TEST_ARGS` environment variable.
You can use this to pass the `-run` argument to `go test`, which accepts a
regular expression for the name of the test that should be run.

```sh
# Runs TestValidatePod in pkg/api/validation with the verbose flag set
make test WHAT=./pkg/apis/core/validation GOFLAGS="-v" KUBE_TEST_ARGS='-run ^TestValidatePod$'

# Runs tests that match the regex ValidatePod|ValidateConfigMap in pkg/api/validation
make test WHAT=./pkg/apis/core/validation GOFLAGS="-v" KUBE_TEST_ARGS="-run ValidatePod\|ValidateConfigMap$"
```

Or if we are using `go test` as our entry point, we could run:

```sh
go test ./pkg/apis/core/validation -v -run ^TestValidatePods$
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
make test WHAT=./pkg/scheduler/internal/cache KUBE_TEST_ARGS='-benchmem -run=XXX -bench=BenchmarkExpirePods'
```

Alternatively, to express in pure Go, you could write the following:

```sh
go test ./pkg/scheduler/internal/cache -benchmem -run=XXX -bench=Benchmark
```

This will do the following:

1. `-run=XXX` is a regular expression filter on the name of test cases to run.
   Go will execute both the tests matching the `-bench` regex and the `-run`
   regex. Since we only want to execute benchmark tests, we set the `-run` regex
   to XXX, which will not match any tests.
2. `-bench=Benchmark` will run test methods with Benchmark in the name
  * See `grep -nr Benchmark .` for examples
3. `-benchmem` enables memory allocation stats

See `go help test` and `go help testflag` for additional info.

## Integration tests

Please refer to [Integration Testing in Kubernetes](integration-tests.md).

## End-to-End tests

Please refer to [End-to-End Testing in Kubernetes](e2e-tests.md).

## Running your contribution through Kubernetes CI
Once you open a PR, [`prow`][prow-url] runs pre-submit tests in CI. You can find more about `prow` in [kubernetes/test-infra][prow-git] and in [this blog post][prow-doc] on automation involved in testing PRs to Kubernetes.

If you are not a [Kubernetes org member][membership], another org member will need to run [`/ok-to-test`][ok-to-test] on your PR.

Find out more about [other commands][prow-cmds] you can use to interact with prow through GitHub comments.

### Troubleshooting a failure
Click on `Details` to look at artifacts produced by the test and the cluster under test, to help you debug the failure. These artifacts include:
- test results
- metadata on the test run (including versions of binaries used, test duration)
- output from tests that have failed
- build log showing the full test run
- logs from the cluster under test (k8s components such as kubelet and apiserver, possibly other logs such as etcd and kernel)
- junit xml files
- test coverage files

If the failure seems unrelated to the change you're submitting:
- Is it a flake?
  - Check if a GitHub issue is already open for that flake
    - If not, open a new one (like [this example][new-issue-example]) and [label it `kind/flake`][kind/flake]
    - If yes, any help troubleshooting and resolving it is very appreciated. Look at [Helping with known flakes](#helping-with-known-flakes) for how to do it.
  - Run [`/retest`][retest] on your PR to re-trigger the tests

- Is it a failure that shouldn't be happening (in other words; is the test expectation now wrong)?
  - Get in touch with the SIG that your PR is labeled after
    - preferably as a comment on your PR, by tagging the [GitHub team][k-teams] (for example a [reviewers team for the SIG][k-teams-review])
    - write your reasoning as to why you think the test is now outdated and should be changed
    - if you don't get a response in 24 hours, engage with the SIG on their channel on the [Kubernetes slack](http://slack.k8s.io/) and/or attend one of the [SIG meetings][sig-meetings] to ask for input.

[prow-url]: https://prow.k8s.io
[prow-git]: https://git.k8s.io/test-infra/prow
[prow-doc]: https://kubernetes.io/blog/2018/08/29/the-machines-can-do-the-work-a-story-of-kubernetes-testing-ci-and-automating-the-contributor-experience/#enter-prow
[membership]: https://github.com/kubernetes/community/blob/master/community-membership.md#member
[k-teams]: https://github.com/orgs/kubernetes/teams
[k-teams-review]: https://github.com/orgs/kubernetes/teams?utf8=%E2%9C%93&query=review
[ok-to-test]: https://prow.k8s.io/command-help#ok_to_test
[prow-cmds]: https://prow.k8s.io/command-help
[retest]: https://prow.k8s.io/command-help#retest
[new-issue-example]: https://github.com/kubernetes/kubernetes/issues/71430
[kind/flake]: https://prow.k8s.io/command-help#kind
[sig-meetings]: https://github.com/kubernetes/community/blob/master/sig-list.md

#### Helping with known flakes
For known flakes (i.e. with open GitHub issues against them), the community deeply values help in troubleshooting and resolving them. Starting points could be:
- add logs from the failed run you experienced, and any other context to the existing discussion
- if you spot a pattern or identify a root cause, notify or collaborate with the SIG that owns that area to resolve them

#### Escalating failures to a SIG
- Figure out corresponding SIG from test name/description
- Mention the SIG's GitHub handle on the issue, optionally `cc` the SIG's chair(s) (locate them under kubernetes/community/sig-<name\>)
- Optionally (or if you haven't heard back on the issue after 24h) reach out to the SIG on slack

### Testgrid
[`testgrid`](https://testgrid.k8s.io/) is a visualization of the Kubernetes CI status.

It is useful as a way to:
- see the run history of a test you are debugging (access it starting from a gubernator report for that test)
- get an overview of the project's general health

`testgrid` is organised in:
- tests
  - collection of assertions in a test file
  - each test is typically owned by a single SIG
  - each test is represented as a row on the grid
- jobs
  - collection of tests
  - each job is typically owned by a single SIG
  - each job is represented as a tab
- dashboards
  - collection of jobs
  - each dashboard is represented as a button
  - some dashboards collect jobs/tests in the domain of a specific SIG (named after and owned by those SIGs), and dashboards to monitor project wide health (owned by SIG-release)
