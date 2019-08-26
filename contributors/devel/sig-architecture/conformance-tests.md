# Conformance Testing in Kubernetes

The Kubernetes Conformance test suite is a subset of e2e tests that SIG
Architecture has approved to define the core set of interoperable features that
all conformant Kubernetes clusters must support. The tests verify that the
expected behavior works as a user might encounter it in the wild.

The process to add new conformance tests is intended to decouple the development
of useful tests from their promotion to conformance:
- Contributors write and submit e2e tests, to be approved by owning SIGs
- Tests are proven to meet the [conformance test requirements] by review
  and by accumulation of data on flakiness and reliability
- A follow up PR is submitted to [promote the test to conformance](#promoting-tests-to-conformance)

NB: This should be viewed as a living document in a few key areas:
- The desired set of conformant behaviors is not adequately expressed by the
  current set of e2e tests, as such this document is currently intended to
  guide us in the addition of new e2e tests than can fill this gap
- This document currently focuses solely on the requirements for GA,
  non-optional features or APIs. The list of requirements will be refined over
  time to the point where it as concrete and complete as possible.
- There are currently conformance tests that violate some of the requirements
  (e.g., require privileged access), we will be categorizing these tests and
  deciding what to do once we have a better understanding of the situation
- Once we resolve the above issues, we plan on identifying the appropriate areas
  to relax requirements to allow for the concept of conformance Profiles that
  cover optional or additional behaviors

## Conformance Test Requirements

Conformance tests currently test only GA, non-optional features or APIs. More
specifically, a test is eligible for promotion to conformance if:

- it tests only GA, non-optional features or APIs (e.g., no alpha or beta
  endpoints, no feature flags required, no deprecated features)
- it does not require direct access to kubelet's API to pass (nor does it
  require indirect access via the API server node proxy endpoint); it MAY
  use the kubelet API for debugging purposes upon failure
- it works for all providers (e.g., no `SkipIfProviderIs`/`SkipUnlessProviderIs`
  calls)
- it is non-privileged (e.g., does not require root on nodes, access to raw
  network interfaces, or cluster admin permissions)
- it works without access to the public internet (short of whatever is required
  to pre-pull images for conformance tests)
- it works without non-standard filesystem permissions granted to pods
- it does not rely on any binaries that would not be required for the linux
  kernel or kubelet to run (e.g., can't rely on git)
- where possible, it does not depend on outputs that change based on OS (nslookup, ping, chmod, ls)
- any container images used within the test support all architectures for which
  kubernetes releases are built
- it passes against the appropriate versions of kubernetes as spelled out in
  the [conformance test version skew policy]
- it is stable and runs consistently (e.g., no flakes), and has been running
  for at least two weeks
- new conformance tests or updates to conformance tests for additional scenarios
  are only allowed before code freeze dates set by the release team to allow
  enough soak time of the changes and gives folks a chance to kick the tires
  either in the community CI or their own infrastructure to make sure the tests
  are robust
- it has a name that is a literal string

Examples of features which are not currently eligible for conformance tests:

- node/platform-reliant features, eg: multiple disk mounts, GPUs, high density,
  etc.
- optional features, eg: policy enforcement
- cloud-provider-specific features, eg: GCE monitoring, S3 Bucketing, etc.
- anything that requires a non-default admission plugin

Conformance tests are intended to be stable and backwards compatible according to 
the standard API deprecation policies. Therefore any test that relies on specific 
output that is not subject to the deprecation policy cannot be promoted to conformance. 
Examples of tests which are not eligible to conformance:
- anything that checks specific Events are generated, as we make no guarantees
  about the contents of events, nor their delivery
    - If a test depends on events it is recommended to change the test to 
      use an informer pattern and watch specific resource changes instead.
- anything that checks optional Condition fields, such as Reason or Message, as
  these may change over time (however it is reasonable to verify these fields
  exist or are non-empty)
    - If the test is checking for specific conditions or reasons, it is considered 
      overly specific and it is recommended to simply look for pass/failure criteria 
      where possible, and output the condition/reason for debugging purposes only.

Examples of areas we may want to relax these requirements once we have a
sufficient corpus of tests that define out of the box functionality in all
reasonable production worthy environments:
- tests may need to create or set objects or fields that are alpha or beta that
  bypass policies that are not yet GA, but which may reasonably be enabled on a
  conformant cluster (e.g., pod security policy, non-GA scheduler annotations)

### Windows & Linux Considerations

Windows node support is an optional but stable feature as of Kubernetes 1.14. This means that it is
not required by conformance testing. Nonetheless, it's important to verify that the behavior of Windows nodes match the behaviors tested in the conformance suite as much as possible. To that end, a
large number of conformance tests are already included in Windows testing. You can see what tests are already passing by looking at TestGrid for results of Windows tests running on
[Azure](https://testgrid.k8s.io/sig-windows#aks-engine-azure-windows-master) and
[GCE](https://testgrid.k8s.io/sig-windows#gce-windows-master)). Tests may be
scheduled for any PR with the bot command `/test pull-kubernetes-e2e-aks-engine-azure-windows`.

Generally speaking, the goals are to:

- Make sure tests that are already passing remain passing. If new OS-specific
functionality is added, it should be in a new test.
- Ensure that new tests covering Linux-specific functionality are tagged with `[LinuxOnly]` 
(see: [Kinds of Tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md#kinds-of-tests),
- Give future reviewers a reference to an active issue or documentation clarifying why a test
cannot run on Windows.

The tests that are running today:

- Rely only on container images that already have a multi-architecture manifest
including Windows versions, or have been ported by SIG-Windows
(see [kubernetes-sigs/windows-testing/images](https://github.com/kubernetes-sigs/windows-testing/tree/master/images)
- Do not depend on any functionality that is different or not available on Windows. The full list
is available in the Windows Kubernetes docs under [api](https://kubernetes.io/docs/setup/windows/intro-windows-in-kubernetes/#api). 
A brief summary is included here as a starting point. If the docs are insufficient
or there are more questions, please contact #SIG-Windows on Slack to get another
reviewer.

Some of the most common differences to watch for are:

- Container Images
  - Watch out for image names hardcoded into test cases or YAML files. These are often Linux-only. Instead, they should be adding to or using existing images from [tests/utils/image/manifest.go](https://github.com/kubernetes/kubernetes/blob/master/test/utils/image/manifest.go). This allows the container registry to be configured to one containing Windows images, and also supports testing on clusters with no internet access using a private registry. Multi-arch images supporting Windows are also acceptable.
- Container Options & Actions
  - Pod SecurityContext is set. Most of the fields are Linux specific, and any field set in the Pod's SecurityContext will result in the Pod not being able to spawn or not work as intended.
  - Privileged containers are not supported. Containers are always isolated.
  - Windows uses job objects or Hyper-V for pod isolation and resource controls, not CGroups. These are managed
implicitly by Docker or ContainerD, not by the kubelet. Do not check properties of CGroups as pass/fail criteria.
  - Running Linux-specific commands are not likely to work. Some commands may work using a Windows [busybox](https://github.com/kubernetes-sigs/windows-testing/tree/master/images/busybox) container. The paths of these binaries may differ from Linux, so it's best to rely on `PATH` rather than using Linux-specific paths such as `/usr/bin/nc`. As an alternative, you can use commands in the cross-platform [agnhost](https://github.com/kubernetes/kubernetes/tree/master/test/images/agnhost) image which is designed to return the same results regardless of OS.
- Storage
  - File permissions cannot be set on volumes. Tests using `DefaultMode` or `Mode` and checking the resulting permissions will fail.
  - Only NTFS volumes are supported. Volume mounts specifying other filesystems (ext4, xfs) or mediums (memory) are not supported
  - Mappings of individual files are not supported. Tests which are mounting or expecting such files to be mounted (including /etc/hosts, /etc/resolv.conf, /dev/termination-log) will fail.
  - Bidirectional mount propagation, specifically propagating mounts from a container to host, does not work.
- Networking
  - Pods set `HostNetwork=true`. Is not supported on Windows, and the Pod will not start.
  - Network and DNS settings must be passed through CNI. Windows does not use `/etc/resolv.conf`, so tests should not rely on reading that file to check DNS settings.
    - If you to check network settings such as dns search lists, please use [agnhost](https://github.com/kubernetes/kubernetes/tree/master/test/images/agnhost) to output needed data from the container. 
  - Windows treats all DNS lookups with a `.` to be FQDN, not PQDN. For example `kubernetes` will resolve as a PQDN,
  but `kubernetes.default` will be resolved as a FQDN and fail.
  - ICMP only works between pods on the same network, and are not routable to external networks. TCP/UDP are routable.
  - Windows containers do not support IPv6.

The existing tests which are affected by one of those criteria are tagged with `[LinuxOnly]` 
(see: [Kinds of Tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md#kinds-of-tests).

## Conformance Test Version Skew Policy

As each new release of Kubernetes provides new functionality, the subset of
tests necessary to demonstrate conformance grows with each release. Conformance
is thus considered versioned, with the same backwards compatibility guarantees
as laid out in the [kubernetes versioning policy]

To quote:

> For example, a v1.3 master should work with v1.1, v1.2, and v1.3 nodes, and
> should work with v1.2, v1.3, and v1.4 clients.

Conformance tests for a given version should be run off of the release branch
that corresponds to that version. Thus `v1.2` conformance tests would be run
from the head of the `release-1.2` branch.

For example, suppose we're in the midst of developing kubernetes v1.3. Clusters
with the following versions must pass conformance tests built from the
following branches:

| cluster version | master | release-1.3 | release-1.2 | release-1.1 |
| --------------- | -----  | ----------- | ----------- | ----------- |
| v1.3.0-alpha    | yes    | yes         | yes         | no          |
| v1.2.x          | no     | no          | yes         | yes         |
| v1.1.x          | no     | no          | no          | yes         |

## Running Conformance Tests

Conformance tests are designed to be run even when there is no cloud provider
configured. Conformance tests must be able to be run against clusters that have
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
export KUBERNETES_CONFORMANCE_TEST=y

# Option A: run all conformance tests serially
kubetest --provider=skeleton --test --test_args="--ginkgo.focus=\[Conformance\]"

# Option B: run parallel conformance tests first, then serial conformance tests serially
kubetest --ginkgo-parallel --provider=skeleton --test --test_args="--ginkgo.focus=\[Conformance\] --ginkgo.skip=\[Serial\]"
kubetest --provider=skeleton --test --test_args="--ginkgo.focus=\[Serial\].*\[Conformance\]"
```

## Kubernetes Conformance Document

For each Kubernetes release, a Conformance Document will be generated that lists
all of the tests that comprise the conformance test suite, along with the formal
specification of each test. For an example, see the [v1.9 conformance doc].
This document will help people understand what features are being tested without
having to look through the testcase's code directly.


## Conformance test review board

The conformance subproject uses the [Conformance Test Review board] to track
progress of PRs through to approval. The following types of PRs must go through
this approval process:

- promotion of tests to conformance
- demotion of tests from conformance
- changes to existing conformance tests
- changes to the conformance criteria or process
- changes to the conformance infrastructure code

There are six columns in this board. New PRs should enter in the To Triage
column, and [Conformance test reviewers] will pick it up from there and move it
through the process. New end-to-end tests that are intended to be promoted to
conformance tests in the future may be added to this board, but they will not
move all the way to the Needs Approval column, as that is intended only for the
types of PRs described above.

## Promoting Tests to Conformance

To promote a test to the conformance test suite, open a PR as follows:
- is titled "Promote xxx e2e test to Conformance"
- includes information and metadata in the description as follows:
  - "/area conformance" on a newline
  - "@kubernetes/sig-architecture-pr-reviews @kubernetes/sig-foo-pr-reviews
    @kubernetes/cncf-conformance-wg" on a new line, where sig-foo is whichever
    sig owns this test
  - any necessary information in the description to verify that the test meets
    [conformance test requirements], such as links to reports or dashboards that
    prove lack of flakiness
- contains no other modifications to test source code other than the following:
  - modifies the testcase to use the `framework.ConformanceIt()` function rather
    than the `framework.It()` function
  - adds a comment immediately before the `ConformanceIt()` call that includes
    all of the required [conformance test comment metadata]
  - `go run test/conformance/walk.go test/e2e > test/conformance/testdata/conformance.txt` 
    adds the test name to the [conformance.txt] file 
- add the PR to SIG Architecture's [Conformance Test Review board] in the To
  Triage column

Once you create the PR, please schedule the additional Windows tests with
`/test pull-kubernetes-e2e-aks-engine-azure-windows` to see if any existing tests
that pass on Windows are broken by the change.


### Conformance Test Comment Metadata

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

The corresponding portion of the Kubernetes Conformance Document for this test
would then look like this:

> ## [Kubelet: log output](https://github.com/kubernetes/kubernetes/tree/release-1.9/test/e2e_node/kubelet_test.go#L47)
>
> Release : v1.9
>
> By default the stdout and stderr from the process being executed in a pod MUST be sent to the pod's logs.

### Reporting Conformance Test Results

Conformance test results, by provider and releases, can be viewed in the
[testgrid conformance dashboard]. If you wish to contribute test results
for your provider, please see the [testgrid conformance README]

[kubernetes versioning policy]: /contributors/design-proposals/release/versioning.md#supported-releases-and-component-skew
[Conformance Test Review board]: https://github.com/orgs/kubernetes/projects/9
[Conformance test reviewers]: https://github.com/kubernetes/kubernetes/blob/master/test/conformance/testdata/OWNERS
[conformance test requirements]: #conformance-test-requirements
[conformance test comment metadata]: #conformance-test-comment-metadata
[conformance test version skew policy]: #conformance-test-version-skew-policy
[testgrid conformance dashboard]: https://testgrid.k8s.io/conformance-all
[testgrid conformance README]: https://github.com/kubernetes/test-infra/blob/master/testgrid/conformance/README.md
[v1.9 conformance doc]: https://github.com/cncf/k8s-conformance/blob/master/docs/KubeConformance-1.9.md
[conformance.txt]: https://github.com/kubernetes/kubernetes/blob/master/test/conformance/testdata/conformance.txt
