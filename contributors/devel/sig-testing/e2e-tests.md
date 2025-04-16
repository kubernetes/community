# End-to-End Testing in Kubernetes

**Table of Contents**

- [End-to-End Testing in Kubernetes](#end-to-end-testing-in-kubernetes)
  - [Overview](#overview)
  - [Building Kubernetes and Running the Tests](#building-kubernetes-and-running-the-tests)
    - [Cleaning up](#cleaning-up)
  - [Advanced testing](#advanced-testing)
    - [Extracting a specific version of Kubernetes](#extracting-a-specific-version-of-kubernetes)
    - [Bringing up a cluster for testing](#bringing-up-a-cluster-for-testing)
    - [Debugging clusters](#debugging-clusters)
    - [Debugging an E2E test with a debugger (delve)](#debugging-an-e2e-test-with-a-debugger-delve)
    - [Local clusters](#local-clusters)
      - [Testing against local clusters](#testing-against-local-clusters)
    - [Version-skewed and upgrade testing](#version-skewed-and-upgrade-testing)
      - [Test jobs naming convention](#test-jobs-naming-convention)
  - [Kinds of tests](#kinds-of-tests)
    - [Viper configuration and hierarchichal test parameters.](#viper-configuration-and-hierarchichal-test-parameters)
    - [Conformance tests](#conformance-tests)
  - [Continuous Integration](#continuous-integration)
    - [What is CI?](#what-is-ci)
    - [What runs in CI?](#what-runs-in-ci)
      - [Non-default tests](#non-default-tests)
    - [The PR-builder](#the-pr-builder)
    - [Adding a test to CI](#adding-a-test-to-ci)
    - [Moving a test out of CI](#moving-a-test-out-of-ci)
  - [Performance Evaluation](#performance-evaluation)
  - [One More Thing](#one-more-thing)


## Overview

End-to-end (e2e) tests for Kubernetes provide a mechanism to test end-to-end
behavior of the system, and is the last signal to ensure end user operations
match developer specifications. Although unit and integration tests provide a
good signal, in a distributed system like Kubernetes it is not uncommon that a
minor change may pass all unit and integration tests, but cause unforeseen
changes at the system level.

The primary objectives of the e2e tests are to ensure a consistent and reliable
behavior of the Kubernetes code base, and to catch hard-to-test bugs before
users do, when unit and integration tests are insufficient.

**NOTE:** If you want test against a cluster, you can use `test/e2e` framework. This page is written about `test/e2e`. If you want to test the `kubelet` code, you can use `test/e2e_node` framework. If you want to know `test/e2e_node` , please see the [e2e-node-tests](../sig-node/e2e-node-tests.md).

The e2e tests in Kubernetes are built atop of
[Ginkgo](http://onsi.github.io/ginkgo/) and
[Gomega](http://onsi.github.io/gomega/). There are a host of features that this
Behavior-Driven Development (BDD) testing framework provides, and it is
recommended that the developer read the documentation prior to diving into the
 tests.

The purpose of *this* document is to serve as a primer for developers who are
looking to execute or add tests using a local development environment.

Before writing new tests or making substantive changes to existing tests, you
should also read [Writing Good e2e Tests](writing-good-e2e-tests.md)

## Building Kubernetes and Running the Tests

There are a variety of ways to run e2e tests, but we aim to decrease the number
of ways to run e2e tests to a canonical way: `kubetest`.

For information on installing `kubetest`, please see the
[installation section](https://github.com/kubernetes/test-infra/tree/master/kubetest#installation) of the
[Kubetest project documentation](https://github.com/kubernetes/test-infra/tree/master/kubetest).

You can run an end-to-end test which will bring up a master and nodes, perform
some tests, and then tear everything down. Make sure you have followed the
getting started steps for your chosen cloud platform (which might involve
changing the --provider flag value to something other than "gce").

You can quickly recompile the e2e testing framework via `go install ./test/e2e`.
This will not do anything besides allow you to verify that the go code compiles.
If you want to run your e2e testing framework without re-provisioning the e2e setup,
you can do so via `make WHAT=test/e2e/e2e.test`, and then re-running the ginkgo tests.

To build Kubernetes, up a cluster, run tests, and tear everything down, use:

```sh
kubetest --build --up --test --down
```

If you'd like to just perform one of these steps, here are some examples:

```sh
# Build binaries for testing
kubetest --build

# Create a fresh cluster.  Deletes a cluster first, if it exists
kubetest --up

# Run all tests
kubetest --test

# Run tests which have been labeled with "Feature:Performance" against a local cluster
# Specify "--provider=local" flag when running the tests locally
kubetest --test --test_args='--ginkgo.label-filter=Feature:Performance' --provider=local

# Conversely, exclude tests that match the regex "Pods.*env"
kubetest --test --test_args='--ginkgo.skip=Pods.*env'

# Exclude tests that require a certain minimum version of the kubelet
kubetest --test --test_args='--ginkgo.label-filter=!MinimumKubeletVersion:1.20'

# Run tests in parallel, skip any that must be run serially
GINKGO_PARALLEL=y kubetest --test --test_args='--ginkgo.label-filter=!Serial'

# Run tests in parallel, skip any that must be run serially and keep the test namespace if test failed
GINKGO_PARALLEL=y kubetest --test --test_args='--ginkgo.label-filter=!Serial --delete-namespace-on-failure=false'

# Flags can be combined, and their actions will take place in this order:
# --build, --up, --test, --down
#
# You can also specify an alternative provider, such as 'aws'
#
# e.g.:
kubetest --provider=aws --build --up --test --down
```

The tests are built into a single binary which can be used to deploy a
Kubernetes system or run tests against an already-deployed Kubernetes system.
See `kubetest --help` for more options, such as reusing an existing cluster.

### Cleaning up

During a run, pressing `control-C` should result in an orderly shutdown, but if
something goes wrong and you still have some VMs running you can force a cleanup
with this command:

```sh
kubetest --down
```

## Advanced testing

### Extracting a specific version of Kubernetes

The `kubetest` binary can download and extract a specific version of Kubernetes,
both the server, client and test binaries. The `--extract=E` flag enables this
functionality.

There are a variety of values to pass this flag:

```sh
# Official builds: <ci|release>/<latest|stable>[-N.N]
kubetest --extract=ci/latest --up  # Deploy the latest ci build.
kubetest --extract=ci/latest-1.5 --up  # Deploy the latest 1.5 CI build.
kubetest --extract=release/latest --up  # Deploy the latest RC.
kubetest --extract=release/stable-1.5 --up  # Deploy the 1.5 release.

# A specific version:
kubetest --extract=v1.5.1 --up  # Deploy 1.5.1
kubetest --extract=v1.5.2-beta.0  --up  # Deploy 1.5.2-beta.0
kubetest --extract=gs://foo/bar  --up  # --stage=gs://foo/bar

# Whatever GKE is using (gke, gke-staging, gke-test):
kubetest --extract=gke  --up  # Deploy whatever GKE prod uses

# Using a GCI version:
kubetest --extract=gci/gci-canary --up  # Deploy the version for next gci release
kubetest --extract=gci/gci-57  # Deploy the version bound to gci m57
kubetest --extract=gci/gci-57/ci/latest  # Deploy the latest CI build using gci m57 for the VM image

# Reuse whatever is already built
kubetest --up  # Most common. Note, no extract flag
kubetest --build --up  # Most common. Note, no extract flag
kubetest --build --stage=gs://foo/bar --extract=local --up  # Extract the staged version
```

### Bringing up a cluster for testing

If you want, you may bring up a cluster in some other manner and run tests
against it. To do so, or to do other non-standard test things, you can pass
arguments into Ginkgo using `--test_args` (e.g. see above). For the purposes of
brevity, we will look at a subset of the options, which are listed below:

```
--ginkgo.dryRun=false: If set, ginkgo will walk the test hierarchy without
actually running anything.

--ginkgo.failFast=false: If set, ginkgo will stop running a test suite after a
failure occurs.

--ginkgo.failOnPending=false: If set, ginkgo will mark the test suite as failed
if any specs are pending.

--ginkgo.focus="": If set, ginkgo will only run specs that match this regular
expression.

--ginkgo.skip="": If set, ginkgo will only run specs that do not match this
regular expression.

--ginkgo.label-filter="": If set, select tests based on their labels as described under
"Spec Labels" in https://onsi.github.io/ginkgo/#filtering-specs. This can focus
on tests and exclude others in a single parameter without using regular expressions.

--ginkgo.noColor="n": If set to "y", ginkgo will not use color in the output

--ginkgo.trace=false: If set, default reporter prints out the full stack trace
when a failure occurs

--ginkgo.v=false: If set, default reporter print out all specs as they begin.

--host="": The host, or api-server, to connect to

--kubeconfig="": Path to kubeconfig containing embedded authinfo.

--provider="": The name of the Kubernetes provider (gce, gke, local, vagrant,
etc.)

--repo-root="../../": Root directory of Kubernetes repository, for finding test
files.
```

Prior to running the tests, you may want to first create a simple auth file in
your home directory, e.g. `$HOME/.kube/config`, with the following:

```
{
  "User": "root",
  "Password": ""
}
```

As mentioned earlier there are a host of other options that are available, but
they are left to the developer.

**NOTE:** If you are running tests on a local cluster repeatedly, you may need
to periodically perform some manual cleanup:

  - `rm -rf /var/run/kubernetes`, clear kube generated credentials, sometimes
stale permissions can cause problems.

  - `sudo iptables -F`, clear ip tables rules left by the kube-proxy.

### Reproducing failures in flaky tests

You can run a test repeatedly until it fails. This is useful when debugging
flaky tests. In order to do so, you need to set the following environment
variable:
```sh
$ export GINKGO_UNTIL_IT_FAILS=true
```

After setting the environment variable, you can run the tests as before. The e2e
script adds `--untilItFails=true` to ginkgo args if the environment variable is
set. The flags asks ginkgo to run the test repeatedly until it fails.

### Debugging clusters

If a cluster fails to initialize, or you'd like to better understand cluster
state to debug a failed e2e test, you can use the `cluster/log-dump.sh` script
to gather logs.

This script requires that the cluster provider supports ssh. Assuming it does,
running:

```sh
$ cluster/log-dump.sh <directory>
```

will ssh to the master and all nodes and download a variety of useful logs to
the provided directory (which should already exist).

The Google-run Jenkins builds automatically collected these logs for every
build, saving them in the `artifacts` directory uploaded to GCS.

### Debugging an E2E test with a debugger (delve)

When debugging E2E tests it's sometimes useful to pause in the middle of an E2E test
to check the value of a variable or to check something in the cluster, instead of adding
`time.Sleep(...)` we can run the E2E test with `delve`

Requirements:

- delve (https://github.com/go-delve/delve/tree/master/Documentation/installation)

For this example we'll debug a [sig-storage test that will provision storage from a snapshot](https://github.com/kubernetes/kubernetes/blob/3ed71cf190a3d6a6dcb965cf73224538059e8e5e/test/e2e/storage/testsuites/provisioning.go#L200-L236)

First, compile the E2E test suite with additional compiler flags

```sh
# DBG=1 enables necessary debug options and disables stripping binaries
# see the makefile upstream, or use KUBE_VERBOSE=3 to get the actual build commands
make WHAT=test/e2e/e2e.test DBG=1
```

Then set the env var `E2E_TEST_DEBUG_TOOL=delve` and then run the test with `./hack/ginkgo.sh` instead of `kubetest`, you should see the delve command line prompt

```sh
E2E_TEST_DEBUG_TOOL=delve ./hack/ginkgo-e2e.sh --ginkgo.focus="sig-storage.*csi-hostpath.*Dynamic.PV.*default.fs.*provisioning.should.provision.storage.with.snapshot.data.source" --allowed-not-ready-nodes=10
---
Setting up for KUBERNETES_PROVIDER="gce".
Project: ...
Network Project: ...
Zone: ...
Trying to find master named '...'
Looking for address '...'
Using master: ... (external IP: XX.XXX.XXX.XX; internal IP: (not set))
Type 'help' for list of commands.
(dlv)
```

Use the commands described in the [delve command lists](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md), for our example we'll set a breakpoint at the start of the method

```sh
(dlv) break test/e2e/storage/testsuites/provisioning.go:201
Breakpoint 1 set at 0x72856f2 for k8s.io/kubernetes/test/e2e/storage/testsuites.(*provisioningTestSuite).DefineTests.func4() _output/local/go/src/k8s.io/kubernetes/test/e2e/storage/testsuites/provisioning.go:201
```

When you're done setting breakpoints execute `continue` to continue the test, once the breakpoint hits you have the chance to explore variables in the test

```sh
(dlv) continue
Apr 16 20:29:18.724: INFO: Fetching cloud provider for "gce"
I0416 20:29:18.725327 3669683 gce.go:909] Using DefaultTokenSource &oauth2.reuseTokenSource{new:(*oauth2.tokenRefresher)(0xc002b65d10), mu:sync.Mutex{state:0, sema:0x0}, t:(*oauth2.Token)(0xc0028e43c0)}
W0416 20:29:18.891866 3669683 gce.go:477] No network name or URL specified.
I0416 20:29:18.892058 3669683 e2e.go:129] Starting e2e run "ae1b58af-9e9e-4745-b1f4-27d763451f8e" on Ginkgo node 1
{"msg":"Test Suite starting","total":1,"completed":0,"skipped":0,"failed":0}
Running Suite: Kubernetes e2e suite
===================================
Random Seed: 1618604956 - Will randomize all specs
Will run 1 of 5745 specs
...
------------------------------
[sig-storage] CSI Volumes [Driver: csi-hostpath] [Testpattern: Dynamic PV (default fs)] provisioning
  should provision storage with snapshot data source [Feature:VolumeSnapshotDataSource]
  _output/local/go/src/k8s.io/kubernetes/test/e2e/storage/testsuites/provisioning.go:200
[BeforeEach] [Testpattern: Dynamic PV (default fs)] provisioning
  _output/local/go/src/k8s.io/kubernetes/test/e2e/storage/framework/testsuite.go:51
[BeforeEach] [Testpattern: Dynamic PV (default fs)] provisioning
  _output/local/go/src/k8s.io/kubernetes/test/e2e/framework/framework.go:185
STEP: Creating a kubernetes client
Apr 16 20:29:24.747: INFO: >>> kubeConfig: ...
STEP: Building a namespace api object, basename provisioning
W0416 20:29:24.901750 3669683 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
Apr 16 20:29:24.901: INFO: No PodSecurityPolicies found; assuming PodSecurityPolicy is disabled.
STEP: Waiting for a default service account to be provisioned in namespace
[It] should provision storage with snapshot data source [Feature:VolumeSnapshotDataSource]
  _output/local/go/src/k8s.io/kubernetes/test/e2e/storage/testsuites/provisioning.go:200
> k8s.io/kubernetes/test/e2e/storage/testsuites.(*provisioningTestSuite).DefineTests.func4() _output/local/go/src/k8s.io/kubernetes/test/e2e/storage/testsuites/provisioning.go:201 (hits goroutine(165):1 total:1) (PC: 0x72856f2)
Warning: listing may not match stale executable
   196:
   197:                 l.testCase.TestDynamicProvisioning()
   198:         })
   199:
   200:         ginkgo.It("should provision storage with snapshot data source [Feature:VolumeSnapshotDataSource]", func() {
=> 201:                 if !dInfo.Capabilities[storageframework.CapSnapshotDataSource] {
   202:                         e2eskipper.Skipf("Driver %q does not support populate data from snapshot - skipping", dInfo.Name)
   203:                 }
   204:                 if !dInfo.SupportedFsType.Has(pattern.FsType) {
   205:                         e2eskipper.Skipf("Driver %q does not support %q fs type - skipping", dInfo.Name, pattern.FsType)
   206:                 }
(dlv) print dInfo
*k8s.io/kubernetes/test/e2e/storage/framework.DriverInfo {
        Name: "csi-hostpath",
        InTreePluginName: "",
        FeatureTag: "",
        MaxFileSize: 104857600,
        SupportedSizeRange: k8s.io/kubernetes/test/e2e/framework/volume.SizeRange {Max: "", Min: "1Mi"},
        SupportedFsType: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/sets.String [
                "": {},
        ],
        SupportedMountOption: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/sets.String nil,
        RequiredMountOption: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/sets.String nil,
        Capabilities: map[k8s.io/kubernetes/test/e2e/storage/framework.Capability]bool [
                "persistence": true,
                "snapshotDataSource": true,
                "multipods": true,
                "block": true,
                "pvcDataSource": true,
                "controllerExpansion": true,
                "singleNodeVolume": true,
                "volumeLimits": true,
        ],
        RequiredAccessModes: []k8s.io/kubernetes/vendor/k8s.io/api/core/v1.PersistentVolumeAccessMode len: 0, cap: 0, nil,
        TopologyKeys: []string len: 0, cap: 0, nil,
        NumAllowedTopologies: 0,
        StressTestOptions: *k8s.io/kubernetes/test/e2e/storage/framework.StressTestOptions {NumPods: 10, NumRestarts: 10},
        VolumeSnapshotStressTestOptions: *k8s.io/kubernetes/test/e2e/storage/framework.VolumeSnapshotStressTestOptions {NumPods: 10, NumSnapshots: 10},}
```

### Local clusters

It can be much faster to iterate on a local cluster instead of a cloud-based
one. To start a local cluster, you can run:

```sh
# The PATH construction is needed because PATH is one of the special-cased
# environment variables not passed by sudo -E
sudo PATH=$PATH hack/local-up-cluster.sh
```

This will start a single-node Kubernetes cluster than runs pods using the local
docker daemon. Press Control-C to stop the cluster.

You can generate a valid kubeconfig file by following instructions printed at the
end of aforementioned script.

#### Testing against local clusters

In order to run an E2E test against a locally running cluster, first make sure
to have a local build of the tests:

```sh
kubetest --build
```

Then point the tests at a custom host directly:

```sh
export KUBECONFIG=/path/to/kubeconfig
kubetest --provider=local --test
```

To control the tests that are run:

```sh
kubetest --provider=local --test --test_args="--ginkgo.focus=Secrets"
```

You will also likely need to specify `minStartupPods` to match the number of
nodes in your cluster. If you're testing against a cluster set up by
`local-up-cluster.sh`, you will need to do the following:

```sh
kubetest --provider=local --test --test_args="--minStartupPods=1 --ginkgo.focus=Secrets"
```

### Version-skewed and upgrade testing

We run version-skewed tests to check that newer versions of Kubernetes work
similarly enough to older versions.  The general strategy is to cover the following cases:

1. One version of `kubectl` with another version of the cluster and tests (e.g.
   that v1.2 and v1.4 `kubectl` doesn't break v1.3 tests running against a v1.3
   cluster).
1. A newer version of the Kubernetes master with older nodes and tests (e.g.
   that upgrading a master to v1.3 with nodes at v1.2 still passes v1.2 tests).
1. A newer version of the whole cluster with older tests (e.g. that a cluster
   upgraded---master and nodes---to v1.3 still passes v1.2 tests).
1. That an upgraded cluster functions the same as a brand-new cluster of the
   same version (e.g. a cluster upgraded to v1.3 passes the same v1.3 tests as
   a newly-created v1.3 cluster).

[kubetest](https://git.k8s.io/test-infra/kubetest) is
the authoritative source on how to run version-skewed tests, but below is a
quick-and-dirty tutorial.

```sh
# Assume you have two copies of the Kubernetes repository checked out, at
# ./kubernetes and ./kubernetes_old

# If using GKE:
export CLUSTER_API_VERSION=${OLD_VERSION}

# Deploy a cluster at the old version; see above for more details
cd ./kubernetes_old
kubetest --up

# Upgrade the cluster to the new version
#
# If using GKE, add --upgrade-target=${NEW_VERSION}
#
# You can target Feature:MasterUpgrade or Feature:ClusterUpgrade
cd ../kubernetes
kubetest --provider=gke --test --check-version-skew=false --test_args="--ginkgo.label-filter=Feature:MasterUpgrade"

# Run old tests with new kubectl
cd ../kubernetes_old
kubetest --provider=gke --test --test_args="--kubectl-path=$(pwd)/../kubernetes/cluster/kubectl.sh"
```

If you are just testing version-skew, you may want to just deploy at one
version and then test at another version, instead of going through the whole
upgrade process:

```sh
# With the same setup as above

# Deploy a cluster at the new version
cd ./kubernetes
kubetest --up

# Run new tests with old kubectl
kubetest --test --test_args="--kubectl-path=$(pwd)/../kubernetes_old/cluster/kubectl.sh"

# Run old tests with new kubectl
cd ../kubernetes_old
kubetest --test --test_args="--kubectl-path=$(pwd)/../kubernetes/cluster/kubectl.sh"
```

#### Test jobs naming convention

**Version skew tests** are named as
`<cloud-provider>-<master&node-version>-<kubectl-version>-<image-name>-kubectl-skew`
e.g: `gke-1.5-1.6-cvm-kubectl-skew` means cloud provider is GKE;
master and nodes are built from `release-1.5` branch;
`kubectl` is built from `release-1.6` branch;
image name is cvm (container_vm).
The test suite is always the older one in version skew tests. e.g. from release-1.5 in this case.

**Upgrade tests**:

If a test job name ends with `upgrade-cluster`, it means we first upgrade
the cluster (i.e. master and nodes) and then run the old test suite with new kubectl.

If a test job name ends with `upgrade-cluster-new`, it means we first upgrade
the cluster (i.e. master and nodes) and then run the new test suite with new kubectl.

If a test job name ends with `upgrade-master`, it means we first upgrade
the master and keep the nodes in old version and then run the old test suite with new kubectl.

There are some examples in the table,
where `->` means upgrading; container_vm (cvm) and gci are image names.

| test name | test suite | master version (image) | node version (image) | kubectl
| --------- | :--------: | :----: | :---:| :---:
| gce-1.5-1.6-upgrade-cluster | 1.5 | 1.5->1.6 | 1.5->1.6 | 1.6
| gce-1.5-1.6-upgrade-cluster-new | 1.6 | 1.5->1.6 | 1.5->1.6 | 1.6
| gce-1.5-1.6-upgrade-master | 1.5 | 1.5->1.6 | 1.5 | 1.6
| gke-container_vm-1.5-container_vm-1.6-upgrade-cluster | 1.5 | 1.5->1.6 (cvm) | 1.5->1.6 (cvm) | 1.6
| gke-gci-1.5-container_vm-1.6-upgrade-cluster-new | 1.6 | 1.5->1.6 (gci) | 1.5->1.6 (cvm) | 1.6
| gke-gci-1.5-container_vm-1.6-upgrade-master | 1.5 | 1.5->1.6 (gci) | 1.5 (cvm) | 1.6

## Kinds of tests

Tests can be labeled. Labels appear with square brackets inside the test names
(the traditional approach) *and* are Ginkgo v2 labels (since Kubernetes v1.29).
Available labels in order of increasing precedence (that is, each label listed
below supersedes the previous ones):

  - If a test has no labels, it is expected to run fast (under five minutes), be
able to be run in parallel, and be consistent.

  - `[Slow]`: If a test takes more than five minutes to run (by itself or in
parallel with many other tests), it is labeled `[Slow]`. This partition allows
us to run almost all of our tests quickly in parallel, without waiting for the
stragglers to finish.

  - `[Serial]`: If a test cannot be run in parallel with other tests (e.g. it
takes too many resources or restarts nodes), it is labeled `[Serial]`, and
should be run in serial as part of a separate suite.

  - `[Disruptive]`: If a test may impact workloads that it didn't create,
 it should be marked as `[Disruptive]`. Examples of disruptive behavior
include, but are not limited to, restarting components or tainting nodes. Any
`[Disruptive]` test is also assumed to qualify for the `[Serial]` label, but
need not be labeled as both. These tests are not run against soak clusters to
avoid restarting components.

  - `[Flaky]`: If a test is found to be flaky and we have decided that it's too
hard to fix in the short term (e.g. it's going to take a full engineer-week), it
receives the `[Flaky]` label until it is fixed. The `[Flaky]` label should be
used very sparingly, and should be accompanied with a reference to the issue for
de-flaking the test, because while a test remains labeled `[Flaky]`, it is not
monitored closely in CI. `[Flaky]` tests are by default not run, unless a
`focus` or `skip` argument is explicitly given.

  - `[Feature:.+]`: If a test has non-default requirements to run or targets
some non-core functionality, and thus should not be run as part of the standard
suite, it receives a `[Feature:.+]` label. This non-default requirement could
be some special cluster setup (e.g. `Feature:IPv6DualStack` indicates that the
cluster must support dual-stack pod and service networks) or that the test has
special behavior that makes it unsuitable for a normal test run (e.g.
`Feature:PerformanceDNS` marks a test that stresses cluster DNS performance
with many services). `[Feature:.+]` tests are not run in our core suites,
instead running in custom suites. If a feature is experimental or alpha and is
not enabled by default due to being incomplete or potentially subject to
breaking changes, it does *not* block PR merges, and thus should run in
some separate test suites owned by the feature owner(s)
(see [Continuous Integration](#continuous-integration) below).

  - `[MinimumKubeletVersion:.+]`: This label must be set on tests that require
a minimum version of the kubelet. Invocations of the test suite can then decide
to `skip` the same tests if kubelets in the cluster do not satisfy the requirement.
For example, `[MinimumKubeletVersion:(1.20|1.21)]` would `skip` tests with minimum
kubelet versions `1.20` and `1.21`.

  - `[Conformance]`: Designate that this test is included in the Conformance
test suite for [Conformance Testing](../sig-architecture/conformance-tests.md). This test must
meet a number of [requirements](../sig-architecture/conformance-tests.md#conformance-test-requirements)
to be eligible for this tag. This tag does not supersed any other labels.

  - `[LinuxOnly]`: If a test is known to be using Linux-specific features
(e.g.: seLinuxOptions) or is unable to run on Windows nodes, it is labeled
`[LinuxOnly]`. When using Windows nodes, this tag should be added to the
`skip` argument. This is not using `[Feature:LinuxOnly]` because that
would have implied changing all CI jobs which skip tests with unknown
requirements.

  - The following tags are not considered to be exhaustively applied, but are
intended to further categorize existing `[Conformance]` tests, or tests that are
being considered as candidate for promotion to `[Conformance]` as we work to
refine requirements:
    - `[Privileged]`: This is a test that requires privileged access
    - `[Deprecated]`: This is a test that exercises a deprecated feature

  - For tests that depend on feature gates, the following are set automatically:
    - `[Alpha]`: This is a test that exercises an alpha feature
    - `[Beta]`: This is a test that exercises a beta feature

    Conceptually, these are non-default requirements as defined above under
    `[Feature:.+]`, but for historic reasons and the sake of brevity they don't
    have that prefix when embedded in test names. They *do* have that prefix in the
    Ginkgo v2 label, so use e.g. `--filter-label=Feature: containsAny Alpha` to
    run them. The normal `--filter-label=Feature: isEmpty` excludes them.

    Note that at the moment, not all jobs filter out tests with `Alpha` or `Beta`
    requirements like that. Therefore all tests with such a requirement also
    have to be annotated with a `[Feature]` tag. This restriction will be lifted
    once migration of jobs to `--filter-label` is completed.

Every test should be owned by a [SIG](/sig-list.md),
and have a corresponding `[sig-<name>]` label.

## Selecting tests to run

See https://onsi.github.io/ginkgo/#filtering-specs for a general introduction.

Focusing on a specific test by its name is useful when interactively running
just one or a few related tests. The test name is a concatenation of multiple
strings. To get a list of all full test names, run:

```console
$ e2e.test -list-tests
The following spec names can be used with 'ginkgo run --focus/skip':
    test/e2e/apimachinery/watchlist.go:41: [sig-api-machinery] API Streaming (aka. WatchList) [Serial] [Feature:WatchList] should be requested when ENABLE_CLIENT_GO_WATCH_LIST_ALPHA is set
    test/e2e/apimachinery/flowcontrol.go:65: [sig-api-machinery] API priority and fairness should ensure that requests can be classified by adding FlowSchema and PriorityLevelConfiguration
    test/e2e/apimachinery/flowcontrol.go:190: [sig-api-machinery] API priority and fairness should ensure that requests can't be drowned out (fairness)
...
```

Or within the Kubernetes repo:

```console
$ go test -v ./test/e2e -args -list-tests
The following spec names can be used with 'ginkgo run --focus/skip':
    test/e2e/apimachinery/watchlist.go:41: [sig-api-machinery] API Streaming (aka. WatchList) [Serial] [Feature:WatchList] should be requested when ENABLE_CLIENT_GO_WATCH_LIST_ALPHA is set
...
```

The same works for other Kubernetes E2E suites, like `e2e_node`.

In Prow jobs, selection by labels is often simpler. See
[below]((#kinds-of-tests) for documentation of the different labels that are in
use. A full list of labels used by a specific E2E suite can be obtained with
`--list-labels`.

A common pattern is to run only tests which have no special cluster setup
requirements and are not flaky:

    --filter-label='Feature: isEmpty && !Flaky'

Feature owners have to ensure that tests excluded that way from shared CI
jobs are executed in dedicated jobs (more on CI below):

    --filter-label='Feature: containsAny MyAwesomeFeature'

In jobs that support certain well-known features it is possible to run tests
which have no special requirements or at least only depend on the supported
features:

    # Alpha APIs and features enabled, allow tests depending on that as
    # long as they have no other special requirements.
    --filter-label='Feature: isSubsetOf Alpha'

### Viper configuration and hierarchichal test parameters.

The future of e2e test configuration idioms will be increasingly defined using viper, and decreasingly via flags.

Flags in general fall apart once tests become sufficiently complicated.  So, even if we could use another flag library, it wouldn't be ideal.

To use viper, rather than flags, to configure your tests:

- Just add "e2e.json" to the current directory you are in, and define parameters in it... i.e. `"kubeconfig":"/tmp/x"`.

Note that advanced testing parameters, and hierarchichally defined parameters, are only defined in viper, to see what they are, you can dive into [TestContextType](https://git.k8s.io/kubernetes/test/e2e/framework/test_context.go).

In time, it is our intent to add or autogenerate a sample viper configuration that includes all e2e parameters, to ship with Kubernetes.

### Pod Security Admission

With introducing Pod Security admission in Kubernetes by default, it is desired to execute e2e tests within bounded pod security policy levels. The default pod security policy in e2e tests is [restricted](https://kubernetes.io/docs/concepts/security/pod-security-admission/#pod-security-levels). This is set in https://github.com/kubernetes/kubernetes/blob/master/test/e2e/framework/framework.go. This ensures that e2e tests follow best practices for hardening pods by default.

Two helper functions are available for returning a minimal [restricted pod security context](https://github.com/kubernetes/kubernetes/blob/c876b30c2b30c0355045d7548c22b6cd42ab58da/test/e2e/framework/pod/utils.go#L156) and a [restricted container security context](https://github.com/kubernetes/kubernetes/blob/c876b30c2b30c0355045d7548c22b6cd42ab58da/test/e2e/framework/pod/utils.go#L172). These can be used to initialize pod or container specs to ensure adherence for the most restricted pod security policy.

If pods need to elevate privileges to either `baseline` or `privileged` a new field - `NamespacePodSecurityEnforceLevel` - was introduced to the e2e framework to specify the necessary namespace enforcement level. Note that namespaces get created in the `BeforeEach()` phase of ginkgo tests.

```
import (
...
  admissionapi "k8s.io/pod-security-admission/api"
...
)


var _ = SIGDescribe("Test", func() {
  ...
  f := framework.NewDefaultFramework("test")
  f.NamespacePodSecurityEnforceLevel = admissionapi.LevelPrivileged
  ...
}
```

This ensures that the namespace returned by `f.Namespace.Name` includes the configured pod security policy level. Note that creating custom namespace names is not encouraged and will not include the configured settings.

### Conformance tests

For more information on Conformance tests please see the [Conformance Testing](../sig-architecture/conformance-tests.md)

## Continuous Integration

A quick overview of how we run e2e CI on Kubernetes.

### What is CI?

We run a battery of [release-blocking jobs](https://testgrid.k8s.io/sig-release-master-blocking)
against `HEAD` of the master branch on a continuous basis, and block merges
via [Tide](https://sigs.k8s.io/prow/cmd/tide) on a subset of those
tests if they fail.

CI results can be found at [ci-test.k8s.io](http://ci-test.k8s.io), e.g.
[ci-test.k8s.io/kubernetes-e2e-gce/10594](http://ci-test.k8s.io/kubernetes-e2e-gce/10594).

### What runs in CI?

We run all default tests (those that aren't marked `[Flaky]` or `[Feature:.+]`)
against GCE and GKE. To minimize the time from regression-to-green-run, we
partition tests across different jobs:

  - `kubernetes-e2e-<provider>` runs all non-`[Slow]`, non-`[Serial]`,
non-`[Disruptive]`, non-`[Flaky]`, non-`[Feature:.+]` tests in parallel.

  - `kubernetes-e2e-<provider>-slow` runs all `[Slow]`, non-`[Serial]`,
non-`[Disruptive]`, non-`[Flaky]`, non-`[Feature:.+]` tests in parallel.

  - `kubernetes-e2e-<provider>-serial` runs all `[Serial]` and `[Disruptive]`,
non-`[Flaky]`, non-`[Feature:.+]` tests in serial.

  - `ci-kubernetes-e2e-kind-alpha-features` runs all tests without any special
    requirements and tests that only have alpha feature gates and API groups
    as requirement.

  - `ci-kubernetes-e2e-kind-beta-features` runs all tests without any special
    requirements and tests that only have beta feature gates and API groups
    as requirement.

We also run non-default tests if the tests exercise general-availability ("GA")
features that require a special environment to run in, e.g.
`kubernetes-e2e-gce-scalability` and `kubernetes-kubemark-gce`, which test for
Kubernetes performance.

#### Non-default tests

Many `[Feature:.+]` tests we don't run in CI. These tests are for features that
are experimental (often in the `experimental` API), and aren't enabled by
default.

### The PR-builder

We also run a battery of tests against every PR before we merge it. These tests
are equivalent to `kubernetes-gce`: it runs all non-`[Slow]`, non-`[Serial]`,
non-`[Disruptive]`, non-`[Flaky]`, non-`[Feature:.+]` tests in parallel. These
tests are considered "smoke tests" to give a decent signal that the PR doesn't
break most functionality. Results for your PR can be found at
[pr-test.k8s.io](http://pr-test.k8s.io), e.g.
[pr-test.k8s.io/20354](http://pr-test.k8s.io/20354) for #20354.

### Adding a test to CI

As mentioned above, prior to adding a new test, it is a good idea to perform a
`-ginkgo.dryRun=true` on the system, in order to see if a behavior is already
being tested, or to determine if it may be possible to augment an existing set
of tests for a specific use case.

If a behavior does not currently have coverage and a developer wishes to add a
new e2e test, navigate to the ./test/e2e directory and create a new test using
the existing suite as a guide.

**NOTE:** To build/run with tests in a new directory within ./test/e2e, add the
directory to import list in ./test/e2e/e2e_test.go

When writing a test, consult #kinds-of-tests above to determine how your test
should be marked, (e.g. `[Slow]`, `[Serial]`; remember, by default we assume a
test can run in parallel with other tests!).

When first adding a test it should *not* go straight into CI, because failures
block ordinary development. A test should only be added to CI after is has been
running in some non-CI suite long enough to establish a track record showing
that the test does not fail when run against *working* software. Note also that
tests running in CI are generally running on a well-loaded cluster, so must
contend for resources; see above about [kinds of tests](#kinds_of_tests).

Generally, a feature starts as `experimental`, and will be run in some suite
owned by the team developing the feature. If a feature is in beta or GA, it
*should* block PR merges and releases. In moving from experimental to beta or GA, tests
that are expected to pass by default should simply remove the `[Feature:.+]`
label, and will be incorporated into our core suites. If tests are not expected
to pass by default, (e.g. they require a special environment such as added
quota,) they should remain with the `[Feature:.+]` label.

Occasionally, we'll want to add tests to better exercise features that are
already GA. These tests also shouldn't go straight to CI. They should begin by
being marked as `[Flaky]` to be run outside of CI, and once a track-record for
them is established, they may be promoted out of `[Flaky]`.

### Moving a test out of CI

If we have determined that a test is known-flaky and cannot be fixed in the
short-term, we may move it out of CI indefinitely. This move should be used
sparingly, as it effectively means that we have no coverage of that test. When a
test is demoted, it should be marked `[Flaky]` with a comment accompanying the
label with a reference to an issue opened to fix the test.

## Performance Evaluation

Another benefit of the e2e tests is the ability to create reproducible loads on
the system, which can then be used to determine the responsiveness, or analyze
other characteristics of the system. For example, the density tests load the
system to 30,50,100 pods per/node and measures the different characteristics of
the system, such as throughput, api-latency, etc.

For a good overview of how we analyze performance data, please read the
following [post](https://kubernetes.io/blog/2015/09/kubernetes-performance-measurements-and/)

For developers who are interested in doing their own performance analysis, we
recommend setting up [prometheus](http://prometheus.io/) for data collection,
and using [grafana](https://prometheus.io/docs/visualization/grafana/) to
visualize the data.  There also exists the option of pushing your own metrics in
from the tests using a
[prom-push-gateway](http://prometheus.io/docs/instrumenting/pushing/).
Containers for all of these components can be found
[here](https://hub.docker.com/u/prom/).

For more accurate measurements, you may wish to set up prometheus external to
Kubernetes in an environment where it can access the major system components
(api-server, controller-manager, scheduler). This is especially useful when
attempting to gather metrics in a load-balanced api-server environment, because
all api-servers can be analyzed independently as well as collectively. On
startup, configuration file is passed to prometheus that specifies the endpoints
that prometheus will scrape, as well as the sampling interval.

```
#prometheus.conf
job: {
  name: "kubernetes"
  scrape_interval: "1s"
  target_group: {
    # apiserver(s)
    target: "http://localhost:8080/metrics"
    # scheduler
    target: "http://localhost:10251/metrics"
    # controller-manager
    target: "http://localhost:10252/metrics"
  }
}
```

Once prometheus is scraping the Kubernetes endpoints, that data can then be
plotted using [grafana](https://prometheus.io/docs/visualization/grafana/),
and alerts can be created against the assortment of metrics that Kubernetes
provides.

## One More Thing

You should also know the [testing conventions](../../guide/coding-conventions.md#testing-conventions).

**HAPPY TESTING!**
