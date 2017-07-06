# Kubernetes test sustaining engineering

This document describes how Kubernetes automated tests are maintained as part
of the development process.

## Definitions

The following definitions are for tests continuously run as part of CI.

- *test*
  - *artifact*: a row in [test grid]
  - a single test run as part of a test job
  - maybe either an e2e test or an integration / unit test
- *test job*
  - *artifact*: a tab in [test grid]
  - a collection of tests that are run together in shared environment. may:
    - run in a specific environment - e.g. [gce, gke, aws], [cvm, gci]
    - run under specific conditions - e.g. []upgrade, version skew, soak, serial]
    - test a specific component - e.g. [federation, node]
- *test infrastructure*
  - not directly shown in the test grid
  - libraries and infrastructure common across tests
- *test failure*
  - persistently failing test runs for a given test
- *test flake*
  - periodically failling test runs for a given test
  
## Ownership

Each test must have an escalation point (email + slack).  The escalation point is responsible for
keeping the test healthy.  Fixes for test failures caused by areas of ownership outside the
responsibility of the escalation point should be coordinated with other teams by the
test escalation point.

Escalation points are expected to be responsive within 24 hours, and prioritize test failure
issues over other issues.

### test

Each test must have an owning SIG or group that serves as the escalation point for flakes and failures.
The name of the owner should be present in the test name so that it is displayed in the test grid.

Owners are expected to maintain a dashboard of the tests that they own and
maintain the test health.

**Note:** e2e test owners are present in the test name

### test job

Each test job must have an owning SIG or group that is responsible for the health of the test job.  The
owner may also serve as an escalation point for issues impacting a test only in that specific test job
(passing in other test jobs).  e.g. If a test only fails on aws or only on gke test jobs, the test job
owner and test owner must identify the owner for resolving the failure.

Owners of test jobs are expected to maintain a dashboard of the test jobs they own and
maintain the test job health.

SIGs should update the [job config] and mark the tests that they own.

### test infrastructure

Issues with underlying test infrastructure (e.g. prow) should be escalated to sig/testing.

## Monitoring project wide test health

Dashboards for Kubernetes release blocking test are present on the [test grid].

The following dashboards are expected to remain healthy throughout the development cycle.

- [release-master-blocking](https://k8s-testgrid.appspot.com/release-master-blocking)
  - Tests run against the master branch
- [1.7-master-upgrade & 1.6-master-upgrade](https://k8s-testgrid.appspot.com/master-upgrade)
  - Upgrade a cluster from 1.7 to the master branch and run tests
  - Upgrade a cluster from 1.6 to the master branch and run tests
- [1.7-master-kubectl-skew](https://k8s-testgrid.appspot.com/master-kubectl-skew)
  - Run tests skewing the master and kubectl by +1/-1 version

## Triaging ownership for test failures

When a test is failing, it must be quickly escalated to the correct owner.  Tests that
are left to fail for days or weeks become toxic and create noise in the system health
metrics.

The [build cop] is expected to ensure that the release blocking tests remain
perpetually healthy by monitoring the test grid and escalating failures.

On test failures, the build cop will follow the [sig escalation](#sig-test-escalation) path.

*Tests without a responsive owner should be assigned a new owner or disabled.*

### test failure

A test is failing.

*Symptom*: A row in the test grid is consistently failing across multiple jobs

*How to check for symptom*: Go to the [triage tool], and
search for the failing test by name.  Check to see if it is failing across
multiple jobs, or just one.

*Action*: Escalate to the owning SIG present in the test name (e.g. SIG-cli)

### test job failure

A test *job* is unhealthy causing multiple unrelated tests to fail.

*Symptom*: Multiple unrelated rows in the test grid are consistently failing in a single job,
but passing in others jobs.

*How to check for symptom*: Go to the [test grid].  Are a bunch of tests failing or just a couple?  Are
those tests passing on other jobs?

*Action*: Escalate to the owning SIG for the test job.

### test failure (only on specifics job)

A test is failing, but only on specific jobs.

*Symptom*: A row in the test grid is consistently failing on a single job, but passing on other jobs.

*How to check for symptom*: Go to the [triage tool], and
search for the failing test by name.  Check to see if it is failing across
multiple jobs, or just one.

*Action*: Escalate to the owning SIG present in the test name (e.g. SIG-cli).  They
will coordinate a fix with the test job owner.

## Triaging ownership for test flakes

To triage ownership flakes, follow the same escalation process for failures.  Flakes are considered less
urgent than persistent failures, but still expected to have a root cause investigation within 1 week.

## Broken test workflow

SIGs are expected to proactively monitor and maintain their tests.  The build cop will also
monitor the health of the entire project, but is intended as backup who will escalate
failures to the owning SIGs.

- File an issue for the broken test so it can be referenced and discovered
  - Set the following labels: `priority/failing-test`, `sig/*`
  - Assign the issue to whoever is working on it
  - Mention the current build cop (TODO: publish this somewhere)
- Root cause analysis of the test failure is performed by the owner
- **Note**: The owning SIG for a test can reassign ownership of a resolution to another SIG only after getting
  approval from that SIG
  - This is done by the target SIG reassigning to themselves, not the test owning SIG assigning to someone else.
- Tests failure is resolved either by fixing the underlying issue or disabling the test
  - Disabling a test maybe the correct thing to do in some cases - such as upgrade tests running e2e tests for alpha
    features disable in newer releases.
- SIG owner monitors the test grid to make sure the tests begin to pass
- SIG owner closes the issue

## SIG test escalation

The build cop should monitor the overall test health of the project, and ensure ownership for any given
test does not fall through the cracks.  When the build cop observer a test failure, they should first
search to see if an issue has been filed already, and if not (optionally file an issue and) escalate to the SIG
escalation point.  If the escalation point is unresponsive within a day, the build cop should escalate to the SIG
googlegroup and/or slack channel, mentioning the SIG leads.  If escalation through the SIG googlegroup,
slack channel and SIG leads is unsuccessful, the build cop should escalate to SIG release through the
googlegroup and slack - mentioning the SIG leads.

The SIG escalation points should be bootstrapped from the [community sig list].

## SIG Recommendations

- Figure out which e2e test jobs are release blocking for your SIG.
- Develop a process for making sure the SIGs test grid remains healthy and resolving test failures.
- Consider moving the e2e tests for the SIG into their own test jobs if this would make maintaining them easier.
- Consider developing a playbook for how to resolve test failures and how do identify whether or not another SIG owns the resolution of the issue.

[community sig list]: https://github.com/kubernetes/community/blob/master/sig-list.md
[triage tool]: https://storage.googleapis.com/k8s-gubernator/triage/index.html
[test grid]: https://k8s-testgrid.appspot.com/
[build cop]: https://github.com/kubernetes/community/blob/master/contributors/devel/on-call-build-cop.md
[release-master-blocking]: https://k8s-testgrid.appspot.com/release-master-blocking#Summary
[1.7-master-upgrade]: https://k8s-testgrid.appspot.com/1.7-master-upgrade#Summary
[1.6-master-upgrade]: https://k8s-testgrid.appspot.com/1.6-master-upgrade#Summary
[1.7-master-kubectl-skew]: https://k8s-testgrid.appspot.com/1.6-1.7-kubectl-skew
[job config]: https://github.com/kubernetes/test-infra/blob/master/jobs/config.json
