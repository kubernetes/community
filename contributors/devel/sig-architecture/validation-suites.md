# Validation Suite Testing in Kubernetes

A Kubernetes Validation suite is a subset of e2e tests that are used to
validate functionality of an optional capability of Kubernetes for use
with a specific release. Validation suites can grow or shrink over time, and can also
change as a functionality progresses from alpha, to beta, and to GA. The flexibility
of Validation suites to change based on experimentation and feedback is a key 
characteristic that differentiates Validation suites from 
[Conformance Tests](../sig-architecture/conformance-tests.md) and 
[Conformance profiles](https://github.com/cncf/k8s-conformance/blob/master/terms-conditions/Certified_Kubernetes_Terms.md). In addition, Validation suites differ from  
Conformance profiles in that they do not require conformance certification approval
in order to be created and utilized.  The creation of a Conformance Profile requires
consensus and discussion across the community because there is a desire to limit
the number of conformance profiles created; a large number of Conformance Profiles
would inhibit the portability of Kubernetes workloads. In contrast, Validation 
suites can be created to validate the functionality of an optional 
capability and the more lengthy community discussion of whether the validation suite
should be promoted to become a Conformance profile or merged into the base
Conformance Tests can be deferred.  Validation suites allow the Kuberentes e2e
testing team to rapidly create test suites for a variety of use cases.  They may
also be used to identify a group of tests that validate functionality on an
optional platform such as Kuberenetes on Windows.  Kubernetes SIG groups are
welcome to define Validation suites as needed.  Validation suites should not 
be confused with [Conformance Tests](../sig-architecture/conformance-tests.md)
which refers to the SIG Architecture core set of interoperable features that
all conformant Kubernetes clusters must support. However, it is important to note that
a Validation suite or portion of it may be promoted to become a 
[Conformance Test](../sig-architecture/conformance-tests.md) if deemed necessary. 

## Validation Suite Test Requirements

Validation Test Suites may be used as deemed necessary by the associated Kubernetes
SIG that is defining the Validation Suite.  Unlike Conformance Tests, Validation
Suites may change over time with additions, deletions, and modifications whenever
the SIG learns more about the problem space and adapts the implementations.
Hence, Validation Suites are not subject to the deprecation cycles such as those 
that apply to features and modification decisions are left to the decision making
process(es) of the SIG.  The end users of the Validation Suites are strongly 
encouraged to participate in the upkeep and maintenance of the suite so
they know what changes are coming up and can adapt their 
product/implementation to the updates of the test suite.

## Running Validation Suite Tests

Validation Suite tests are designed to be run even when there is no cloud provider
configured. The [Testing Commons](https://github.com/kubernetes/community/tree/master/sig-testing) 
subproject within the Kubernetes sig-testing
community is currently tasked with adding a framework for writing test suites and
this new framework may be used for running Validation Suites. This new framework is 
optional because there is not a prescribed approach for how a Validation Suite needs to 
be run.  Instead, the authors of the Validation Suite need to ensure that the 
mechanisms for running the Validation Suite are easy to use, provide sufficient
documentation, and follow patterns similar to existing e2e tests. 
In addition, the mechanism should provide the  ability to select some tests and 
skip others.  

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
metadata.  Note how the test contains a bracketed label with name of the Validation
Suite that this test belongs to. 

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


