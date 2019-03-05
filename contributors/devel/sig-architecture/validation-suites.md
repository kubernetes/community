# Validation Suite Testing in Kubernetes

A Kubernetes Validation suite is a subset of e2e tests that are used to
validate correct functionality of an optional capability of Kubernetes. For example,
Validation suites may be used to verify that optional storage or networking 
capabilities are performing correctly.  They may also 
be used to identify a group of tests that validate correct functionality on an
optional platform such as Kuberenetes on Windows.  Kubernetes SIG groups are
welcome to define Validation suites as needed.  Validation suites should not 
be confused with [Conformance Tests](../sig-architecture/conformance-tests.md)
which instead refers to the SIG Architecture core set of interoperable features that
all conformant Kubernetes clusters must support. Instead, Validation suites 
validate correct behavior on an optional or extension capability.  Over time a
a Validation suite or portion of it may be promoted to become a 
[Conformance Test](../sig-architecture/conformance-tests.md) if deemed necessary. 

## Validation Suite Test Requirements

Validation Test Suites may be used as deemed necessary by the associated Kubernetes
SIG that is defining the Validation Suite. 

## Running Validation Suite Tests

Validation Suite tests are designed to be run even when there is no cloud provider
configured. Validation Suite tests must be able to be run against clusters that have
not been created with `hack/e2e.go`, just provide a kubeconfig with the
appropriate endpoint and credentials.

These commands are intended to be run within a kubernetes directory, either
cloned from source, or extracted from release artifacts such as
`kubernetes.tar.gz`. They assume you have a valid golang installation.

```sh
# ensure kubetest is installed
go get -u k8s.io/test-infra/kubetest

# build test binaries, ginkgo, and kubectl first:
make WHAT="test/e2e/e2e.test vendor/github.com/onsi/ginkgo/ginkgo cmd/kubectl"

# setup for conformance tests
export KUBECONFIG=/path/to/kubeconfig

# Option A: run all validation suite tests serially
kubetest --provider=skeleton --test --test_args="--ginkgo.focus=\[*Validation\]"

# Option B: run parallel validations suite tests first, then serial validations tests serially
kubetest --ginkgo-parallel --provider=skeleton --test --test_args="--ginkgo.focus=\[*Validation\] --ginkgo.skip=\[Serial\]"
kubetest --provider=skeleton --test --test_args="--ginkgo.focus=\[Serial\].*\[*Validation\]"
```

Note that if you would like to run a specific Validation Suite instead of all the
Validation Suites, replace `*Validation` in `--ginkgo.focus=\[*Validation\]` with
a more specific Validation Suite name such as `StorageValidation`



### Validation Suite Test Comment Metadata

Each Validation suite test must include the following piece of metadata
within its associated comment:

- `Release`: indicates the Kubernetes release that the test was added to the
  Validation test suite. If the test was modified in subsequent releases
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

### Sample Validation Suite Test

The following snippet of code shows a sample Storage Validation Suite test's 
metadata:

```
/*
  Release : v1.14
  Testname: CSI Attach, Non-attachable volume used in a Pod
  Description: When CSI driver attach is called on non attachable volumes it should not create a volume attachment.
*/
It("should not require VolumeAttach for drivers without attachment [StorageValidation]", func() {
  ...
})

```


