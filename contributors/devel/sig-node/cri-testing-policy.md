# Container Runtime Interface: Testing Policy

**Owner: SIG-Node**

This document describes testing policy and process for runtimes implementing the
[Container Runtime Interface (CRI)](/contributors/devel/sig-node/container-runtime-interface.md)
to publish test results in a federated dashboard. The objective is to provide
the Kubernetes community an easy way to track the conformance, stability, and
supported features of a CRI runtime.

This document focuses on Kubernetes node/cluster end-to-end (E2E) testing
because many features require integration of runtime, OS, or even the cloud
provider. A higher-level integration tests provider better signals on vertical
stack compatibility to the Kubernetes community. On the other hand, runtime
developers are strongly encouraged to run low-level
[CRI validation test suite](https://github.com/kubernetes-sigs/cri-tools/blob/master/docs/validation.md)
for validation as part of their development process.

## Required and optional tests

Runtime maintainers are **required** to submit the tests listed below.
 1. Node conformance test suite
 2. Node feature test suite

Node E2E tests qualify an OS image with a pre-installed CRI runtime. The
runtime maintainers are free to choose any OS distribution, packaging, and
deployment mechanism. Please see the
[tutorial](e2e-node-tests.md)
to know more about the Node E2E test framework and tests for validating a
compatible OS image.

The conformance suite is a set of platform-agnostic (e.g., OS, runtime, and
cloud provider) tests that validate the conformance of the OS image. The feature
suite allows the runtime to demonstrate what features are supported with the OS
distribution.

In addition to the required tests, the runtime maintainers are *strongly
recommended to run and submit results from the Kubernetes conformance test
suite*. This cluster-level E2E test suite provides extra test signal for areas
such as Networking, which cannot be covered by CRI, or Node-level
tests. Because networking requires deep integration between the runtime, the
cloud provider, and/or other cluster components, runtime maintainers are
recommended to reach out to other relevant SIGs (e.g., SIG-GCP or SIG-AWS) for
guidance and/or sponsorship.

## Process for publishing test results

To publish tests results, please submit a proposal in the
[Kubernetes community repository](https://github.com/kubernetes/community)
briefly explaining your runtime, providing at least two maintainers, and
assigning the proposal to the leads of SIG-Node.

These test results should be published under the `sig-node` tab, organized
as follows.

```
sig-node -> sig-node-cri-{Kubernetes-version} -> [page containing the required jobs]
```

Only the last three most recent Kubernetes versions and the master branch are
kept at any time. This is consistent with the Kubernetes release schedule and
policy.

## Test job maintenance

Tests are required to run at least nightly.

The runtime maintainers are responsible for keeping the tests healthy. If the
tests are deemed not actively maintained, SIG-Node may remove the tests from
the test grid at their discretion.

## Process for adding pre-submit testing

If the tests are in good standing (i.e., consistently passing for more than 2
weeks), the runtime maintainers may request that the tests to be included in the
pre-submit Pull Request (PR) tests. Please note that the pre-submit tests
require significantly higher testing capacity, and are held at a higher standard
since they directly affect the development velocity.

If the tests are flaky or failing, and the maintainers are unable to respond and
fix the issues in a timely manner, the SIG leads may remove the runtime from
the presubmit tests until the issues are resolved.

As of now, SIG-Node only accepts promotion of Node conformance tests to
pre-submit because Kubernetes conformance tests involve a wider scope and may
need co-sponsorships from other SIGs.

## FAQ

 *1. Can runtime maintainers publish results from other E2E tests?*

Yes, runtime maintainers can publish additional Node E2E tests results. These
test jobs will be displayed in the `sig-node-{runtime-name}` page. The same
policy for test maintenance applies.

As for additional Cluster E2E tests, SIG-Node may agree to host the
results. However, runtime maintainers are strongly encouraged to seek for a more
appropriate SIG to sponsor or host the results.

 *2. Can these runtime-specific test jobs be considered release blocking?*

This is beyond the authority of SIG-Node, and requires agreement and consensus
across multiple SIGs (e.g., Release, the relevant cloud provider SIG, etc).

 *3. How to run the aforementioned tests?*

It is hard to keep instructions are even links to them up-to-date in one
document. Please contact the relevant SIGs for assistance.

 *4. How can I change the test-grid to publish the test results?*

Please contact SIG-Node for the detailed instructions.

 *5. How does this policy apply to Windows containers?*

Windows containers are still in the early development phase and the features
they support change rapidly. Therefore, it is suggested to treat it as a
feature with select, whitelisted tests to run.
