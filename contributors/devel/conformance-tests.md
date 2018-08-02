# Conformance Testing in Kubernetes

The Kubernetes Conformance test suite is a subset of e2e tests that 
SIG Architecture has approved to define the core set of interoperable 
features that all conformant Kubernetes clusters must support. The 
tests verify that the expected behavior works as a user might encounter 
it in the wild.

The process to add new conformance tests is intended to decouple the
development of useful tests from their promotion to conformance:

- Contributors write and submit e2e tests, to be approved by owning SIGs
- Tests are proven to meet the conformance test requirements
- A follow up PR is submitted to promote the test to conformance

and tests should begin by covering the most important and visible aspects of the function.

### Conformance Test Requirements

Conformance tests test only GA, non-optional features or APIs.  More 
specifically, a test is eligible for promotion to conformance if:

- it tests only GA, non-optional features or APIs (ie: no alpha or beta endpoints, 
  no feature flags required, no deprecated features)
- it works for all providers (ie: no `SkipIfProviderIs`/`SkipUnlessProviderIs` calls)
- it is non-privileged (ie: no root on nodes, network, nor cluster) ([citation needed] @timothysc what existing tests violate this requirement)
- it works without access to the public internet
- it works without non-standard filesystem permissions granted to containers or users
  (ie: can't assume a writable host /tmp shared with container)
- it does not rely on any binaries that would not be required for the linux kernel or
  kubelet to run (ie: can't rely on git)
- any container images used within the test support all architectures for which
  kubernetes releases are built
- it is stable and runs consistently (ie: no flakes)

Examples of features which are not eligible for conformance tests:

- node/platform-reliant features, eg: multiple disk mounts, GPUs, high density, etc.
- optional features, eg: policy enforcement
- cloud-provider-specific features, eg: GCE monitoring, S3 Bucketing, etc.
- anything that requires a non-default admission plugin

Examples of tests which are not eligible for promotion to conformance:

- anything that checks specific Events are generated ([citation needed]
  @bgrant0607 we don't make guarantee anything about Events)
- anything that checks optional Condition fields, such as Reason or Message,
  as these may change over time ([citation needed] @bgrant0607 in some cases 
  they could be tested to be non-empty however)

Our intent is to refine the above list of requirements over time to the point
where it is as concrete and complete as possible. Once we reach this point, we
plan on identifying the appropriate areas to relax these requirements to allow
for the concept of conformance Profiles that cover optional or additional 
behaviors.

### Conformance Test Version Skew Policy

As each new release of Kubernetes provides new functionality, the subset of
tests necessary to demonstrate conformance grows with each release. Conformance
is thus considered versioned, with the same backwards compatibility guarantees
as laid out in [the kubernetes versioning policy]

To quote:

> For example, a v1.3 master should work with v1.1, v1.2, and v1.3 nodes, and 
> should work with v1.2, v1.3, and v1.4 clients.

Conformance tests for a given version should be run off of the release branch
that corresponds to that version. Thus `v1.2` conformance tests would be run
from the head of the `release-1.2` branch.

For example, suppose we're in the midst of developing kubernetes v1.3. The
following clusters must pass conformance tests built from the following branches:

| cluster version | master | release-1.3 | release-1.2 | release-1.1 |
| --------------- | -----  | ----------- | ----------- | ----------- |
| v1.3.0-alpha    | yes    | yes         | yes         | no          |
| v1.2.x          | no     | no          | yes         | yes         |
| v1.1.x          | no     | no          | no          | yes         |

### Running Conformance Tests

Conformance tests are designed to be run with no cloud provider configured.
Conformance tests can be run against clusters that have not been created with
`hack/e2e.go`, just provide a kubeconfig with the appropriate endpoint and
credentials.

```sh
# build test binaries, ginkgo, and kubectl first:
make WHAT="test/e2e/e2e.test vendor/github.com/onsi/ginkgo/ginkgo cmd/kubectl"

# setup for conformance tests
export KUBECONFIG=/path/to/kubeconfig
export KUBERNETES_CONFORMANCE_TEST=y

# run all conformance tests
go run hack/e2e.go -- --provider=skeleton --test --test_args="--ginkgo.focus=\[Conformance\]"

# run all parallel-safe conformance tests in parallel
GINKGO_PARALLEL=y go run hack/e2e.go -- --provider=skeleton --test --test_args="--ginkgo.focus=\[Conformance\] --ginkgo.skip=\[Serial\]"

# ... and finish up with remaining tests in serial
go run hack/e2e.go -- --provider=skeleton --test --test_args="--ginkgo.focus=\[Serial\].*\[Conformance\]"
```

### Kubernetes Conformance Document
For each Kubernetes release, a Conformance Document will be generated
that lists all of the tests that comprise the conformance test suite, along
with the formal specification of each test. For example conformance document for 
1.9 can be found [here](https://github.com/cncf/k8s-conformance/blob/master/docs/KubeConformance-1.9.md).
This document will help people understand what features are being tested without having to look through
the testcase's code directly.


## Adding New Tests

To promote a testcase to the conformance test suite, the following
steps must be taken:
- the test case must already be a part of e2e, and demonstrated to be not flaky
  ([citation needed] @spiffxp how do we demonstrate the test isn't flaky)
- the testcase must use the `framework.ConformanceIt()` function rather
  than the `framework.It()` function
- the testcase must include a comment immediately before the
  `framework.ConformanceIt()` call that includes all of the required
  metadata about the test (see the [Test Metadata](#test-metadata) section)
- use "Promote xxx e2e test to Conformance" as template of your PR title
- tag your PR with "/area conformance" label
- send your PR to Sig-Architecture for review by adding "@kubernetes/sig-architecture-pr-reviews" 
also CC the relevant Sig and Sig-Architecture
- add your PR to SIG Architecture's [Conformance Test Review board] 


### Test Metadata

Each conformance test must include the following piece of metadata
within its associated comment:

- `Release`: indicates the Kubernetes release that the test was added to the
  conformance test suite. If the test was modified in subsequent releases
  then those releases should be included as well (comma separated)
- `Testname`: a human readable short name of the test
- `Description`: a detailed description of the test. This field must describe
  the required behaviour of the Kubernetes components being tested using 
  [RFC2119](https://tools.ietf.org/html/rfc2119) keywords. This field
  is meant to be a "specification" of the tested Kubernetes features, as
  such, it must be detailed enough so that readers can fully understand
  the aspects of Kubernetes that are being tested without having to read
  the test's code directly. Additionally, this test should provide a clear
  distinction between the parts of the test that are there for the purpose
  of validating Kubernetes rather than simply infrastructure logic that
  is necessary to setup, or clean up, the test.

### Sample Conformance Test

The following snippet of code shows a sample conformance test's metadata:

```
/*
  Release : v1.9
  Testname: Kubelet: log output
  Description: By default the stdout and stderr from the process being
  executed in a pod MUST be sent to the pod's logs.
*/
framework.ConformanceIt("it should print the output to logs", func() {
  ...
})
```

The corresponding portion of the Kubernetes Conformance Documentfor this test would then look
like this:

>
> ## [Kubelet: log output](https://github.com/kubernetes/kubernetes/tree/release-1.9/test/e2e_node/kubelet_test.go#L47)
> 
> Release : v1.9
> 
> By default the stdout and stderr from the process being executed in a pod MUST be sent to the pod's logs.

### Reporting Conformance Test Results

Conformance test results, by provider and releases, can be viewed in the 
federated [Conformance TestGrid dashboard](https://k8s-testgrid.appspot.com/conformance-all). 
If you wish to contribute conformance test results for your provider, 
please follow this [on-boarding document](https://docs.google.com/document/d/1lGvP89_DdeNO84I86BVAU4qY3h2VCRll45tGrpyx90A/edit#).


[the kubernetes versioning policy]: /contributors/design-proposals/release/versioning.md#supported-releases-and-component-skew
[Conformance Test Review board]: https://github.com/kubernetes-sigs/architecture-tracking/projects/1
