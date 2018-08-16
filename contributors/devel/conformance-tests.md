# Conformance Testing in Kubernetes

The Kubernetes conformance test suite is a set of testcases, currently a
subset of the integration/e2e tests, that the Architecture SIG has approved
to define the core set of interoperable features that all Kubernetes
deployments must support.

Contributors must write and submit e2e tests first (approved by owning Sigs). 
Once the new tests prove to be stable in CI runs, later create a follow up PR 
to add the test to conformance. This approach also decouples the development 
of useful tests from their promotion to conformance.

A conformance test verifies the expected functionality works as a user might encounter it in the wild, 
and tests should begin by covering the most important and visible aspects of the function.

### Conformance Test Requirements

A test is eligible for promotion to conformance if it meets the following requirements:

- testing GA feature (not alpha or beta APIs, nor deprecated features)
- must be portable (not dependent on provider-specific capabilities or on the public internet)
- cannot test a feature which obviously cannot be supported on a broad range of platforms 
(i.e. testing of multiple disk mounts, GPUs, high density)
- cannot test an optional feature (e.g. not policy enforcement)
- should be non-privileged (neither root on nodes, network, nor cluster)
- cannot rely on any particular non-standard file system permissions granted to containers or users 
(i.e. sharing writable host /tmp with a container)
- should be stable and run consistently
- cannot skip providers (there should be no Skip like directives added to the test), 
especially in the Nucleus or Application layers as described 
[here](https://github.com/kubernetes/community/blob/master/contributors/devel/architectural-roadmap.md).
- cannot test cloud provider specific features (i.e. GCE monitoring, S3 Bucketing, ...)
- should work with default settings for all configuration parameters 
(example: the default list of admission plugins should not have to be tweaked for passing conformance).
- cannot rely on any binaries that are not required for the
linux kernel or for a kubelet to run (i.e. git)
- any container images used in the test must support all architectures for which kubernetes releases are built

### Conformance Test Version Skew Policy

As each new release of Kubernetes provides new functionality, the subset of
tests necessary to demonstrate conformance grows with each release. Conformance
is thus considered versioned, with the same backwards compatibility guarantees
as laid out in [our versioning policy](/contributors/design-proposals/release/versioning.md#supported-releases-and-component-skew).
Conformance tests for a given version should be run off of the release branch
that corresponds to that version. Thus `v1.2` conformance tests would be run
from the head of the `release-1.2` branch. eg:

 - A v1.3 development cluster should pass v1.1, v1.2 conformance tests

 - A v1.2 cluster should pass v1.1, v1.2 conformance tests

 - A v1.1 cluster should pass v1.0, v1.1 conformance tests, and fail v1.2
conformance tests


### Running Conformance Tests

Conformance tests are designed to be run with no cloud provider configured.
Conformance tests can be run against clusters that have not been created with
`hack/e2e.go`, just provide a kubeconfig with the appropriate endpoint and
credentials.

```sh
# build test binaries, ginkgo, and kubectl first:
make WHAT=test/e2e/e2e.test && make WHAT=ginkgo && make WHAT=cmd/kubectl

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
- as a prerequisite, the test case is already part of e2e and is not flaky 
- the testcase must use the `framework.ConformanceIt()` function rather
  than the `framework.It()` function
- the testcase must include a comment immediately before the
  `framework.ConformanceIt()` call that includes all of the required
  metadata about the test (see the [Test Metadata](#test-metadata) section)
- use "Promote xxx e2e test to Conformance" as template of your PR title
- tag your PR with "/area conformance" label
- send your PR to Sig-Architecture for review by adding "@kubernetes/sig-architecture-pr-reviews" 
also CC the relevant Sig and Sig-Architecture


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

