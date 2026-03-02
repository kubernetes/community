# Flaky Tests

Any test that fails occasionally is "flaky". Since our merges only proceed when
all tests are green, and we have a number of different CI systems running the
tests in various combinations, even a small percentage of flakes results in a
lot of pain for people waiting for their PRs to merge.

Therefore, it's important we take flakes seriously. We should avoid flakes by
writing our tests defensively. When flakes are identified, we should prioritize
addressing them, either by fixing them or quarantining them off the critical
path.

The project has a "zero-flake" policy. Test jobs must not automatically retry on test failures.
This was announced and implemented in effect from 2019:
[No more ginkgo.flakeAttempts=2 for e2e tests as of 2019-12-13](https://groups.google.com/g/kubernetes-dev/c/NNmEGUsJObg/m/dmI2mVc_AAAJ)
(and then confirmed as policy in 2023).

For more information about deflaking Kubernetes tests, you can watch:
- @liggitt's [presentation from Kubernetes SIG Testing - 2020-08-25](https://www.youtube.com/watch?v=Ewp8LNY_qTg).
- @aojea's [presentation from Kubernetes SIG Testing - 2022-11-15](https://www.youtube.com/watch?v=x2Lj-ldR0AA&t=2660s).
- @aojea's [Contributor Summit: "The art of deflaking Kubernetes tests"](https://www.youtube.com/watch?v=wyMyQdvg1Qw).

**Table of Contents**

- [Flaky Tests](#flaky-tests)
  - [Avoiding Flakes](#avoiding-flakes)
  - [Quarantining Flakes](#quarantining-flakes)
  - [Hunting Flakes](#hunting-flakes)
  - [GitHub Issues for Known Flakes](#github-issues-for-known-flakes)
    - [Expectations when a flaky test is assigned to you](#expectations-when-a-flaky-test-is-assigned-to-you)
    - [Writing a good flake report](#writing-a-good-flake-report)
  - [Deflaking unit tests](#deflaking-unit-tests)
  - [Deflaking integration tests](#deflaking-integration-tests)
  - [Deflaking e2e tests](#deflaking-e2e-tests)
    - [Gathering information](#gathering-information)
    - [Filtering and correlating information](#filtering-and-correlating-information)
    - [What to look for](#what-to-look-for)
  - [Hunting flaky unit tests in Kubernetes](#hunting-flaky-unit-tests-in-kubernetes)

## Avoiding Flakes

Write tests defensively. Remember that "almost never" happens all the time when
tests are run thousands of times in a CI environment. Tests need to be tolerant
of other tests running concurrently, resource contention, and things taking
longer than expected.

There is a balance to be had here. Don't log too much, but don't log too little.
Don't assume things will succeed after a fixed delay, but don't wait forever.

- Ensure the test functions in parallel with other tests
  - Be specific enough to ensure a test isn't thrown off by other tests' assets
    - https://github.com/kubernetes/kubernetes/pull/85849 - eg: ensure resource name and namespace match
    - https://github.com/kubernetes/kubernetes/pull/85967 - eg: tolerate errors for non k8s.io APIs
    - https://github.com/kubernetes/kubernetes/pull/85619 - eg: tolerate multiple storage plugins
- Ensure the test functions in a resource constrained environment
  - Only ask for the resources you need
    - https://github.com/kubernetes/kubernetes/pull/84975 - eg: drop memory constraints for test cases that only need cpu
  - Don't use overly tight deadlines (but not overly broad either, non-[Slow] tests timeout after 5min)
    - https://github.com/kubernetes/kubernetes/pull/85847 - eg: poll for `wait.ForeverTestTimeout` instead of 10s
    - https://github.com/kubernetes/kubernetes/pull/84238 - eg: poll for 2min instead of 1min
    - mark tests as [Slow] if they are unable to pass within 5min
  - Do not expect actions to happen instantaneously or after a fixed delay
  - Prefer informers and wait loops
- Ensure the test provides sufficient context in logs for forensic debugging
  - Explain what the test is doing, eg:
    - "creating a foo with invalid configuration"
    - "patching the foo to have a bar"
  - Explain what specific check failed, and how, eg:
    - "failed to create resource foo in namespace bar because of err"
    - "expected all items to be deleted, but items foo, bar, and baz remain"
  - Explain why a polling loop is failing, eg: 
    - "expected 3 widgets, found 2, will retry"
    - "expected pod to be in state foo, currently in state bar, will retry"

## Quarantining Flakes

- When quarantining a presubmit test, ensure an issue exists in the current
  release milestone assigned to the owning SIG. The issue should be labeled
  `priority/critical-urgent`, `lifecycle/frozen`, and `kind/flake`. The
  expectation is for the owning SIG to resolve the flakes and reintroduce the
  test, or determine the tested functionality is covered via another method
  and delete the test in question.
- Quarantine a single test case by adding `[Flaky]` to the test name in question,
  most CI jobs exclude these tests. This makes the most sense for flakes that
  are merge-blocking and taking too long to troubleshoot, or occurring across
  multiple jobs.    - eg: https://github.com/kubernetes/kubernetes/pull/83792
  - eg: https://github.com/kubernetes/kubernetes/pull/86327
- Quarantine an entire set of tests by adding `[Feature:Foo]` to the test(s) in
  question. This will require creating jobs that focus specifically on this
  feature. The majority of release-blocking and merge-blocking suites avoid
  these jobs unless they're proven to be non-flaky.

## Hunting Flakes

We offer the following tools to aid in finding or troubleshooting flakes

- [flakes-latest.json](http://storage.googleapis.com/k8s-metrics/flakes-latest.json)
  - shows the top 10 flakes over the past week for all PR jobs
- [go.k8s.io/triage] - an interactive test failure report providing filtering and drill-down by job name, test name, failure text for failures in the last two weeks
  - https://storage.googleapis.com/k8s-gubernator/triage/index.html?pr=1&job=pull-kubernetes-e2e-gce%24 - all failures that happened in the `pull-kubernetes-e2e-gce` job
  - https://storage.googleapis.com/k8s-gubernator/triage/index.html?text=timed%20out - all failures containing the text `timed out`
  - https://storage.googleapis.com/k8s-gubernator/triage/index.html?test=%5C%5Bsig-apps%5C%5D - all failures that happened in tests with `[sig-apps]` in their name
- [testgrid.k8s.io] - display test results in a grid for visual identififcation of flakes
  - https://testgrid.k8s.io/presubmits-kubernetes-blocking - all merge-blocking jobs
  - https://testgrid.k8s.io/presubmits-kubernetes-blocking#pull-kubernetes-e2e-gce&exclude-filter-by-regex=BeforeSuite&sort-by-flakiness= - results for the pull-kubernetes-e2e-gce job sorted by flakiness
  - https://testgrid.k8s.io/sig-release-master-informing#gce-cos-master-default&sort-by-flakiness=&width=10 - results for the equivalent CI job
- [`kind/flake` github query][flake] - open issues or PRs related to flaky jobs or tests for kubernetes/kubernetes

[go.k8s.io/triage]: https://go.k8s.io/triage
[testgrid.k8s.io]: https://testgrid.k8s.io

## GitHub Issues for Known Flakes

Because flakes may be rare, it's very important that all relevant logs be
discoverable from the issue.

1. Search for the test name. If you find an open issue and you're 90% sure the
   flake is exactly the same, add a comment instead of making a new issue.
2. If you make a new issue, you should title it with the test name, prefixed by
   "[Flaky test]"
3. Reference any old issues you found in step one. Also, make a comment in the
   old issue referencing your new issue, because people monitoring only their
   email do not see the backlinks github adds. Alternatively, tag the person or
   people who most recently worked on it.
4. Paste, in block quotes, the entire log of the individual failing test, not
   just the failure line.
5. Link to spyglass to provide access to all durable artifacts and logs (eg: https://prow.k8s.io/view/gcs/kubernetes-jenkins/logs/ci-kubernetes-e2e-gci-gce-flaky/1204178407886163970)

Find flaky tests issues on GitHub under the [kind/flake issue label][flake].
There are significant numbers of flaky tests reported on a regular basis. Fixing
flakes is a quick way to gain expertise and community goodwill.

[flake]: https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fflake

### Expectations when a flaky test is assigned to you

Note that we won't randomly assign these issues to you unless you've opted in or
you're part of a group that has opted in. We are more than happy to accept help
from anyone in fixing these, but due to the severity of the problem when merges
are blocked, we need reasonably quick turn-around time on merge-blocking or
release-blocking flakes. Therefore we have the following guidelines:

1. If a flaky test is assigned to you, it's more important than anything else
   you're doing unless you can get a special dispensation (in which case it will
   be reassigned).  If you have too many flaky tests assigned to you, or you
   have such a dispensation, then it's *still* your responsibility to find new
   owners (this may just mean giving stuff back to the relevant Team or SIG Lead).
2. You should make a reasonable effort to reproduce it. Somewhere between an
   hour and half a day of concentrated effort is "reasonable". It is perfectly
   reasonable to ask for help!
3. If you can reproduce it (or it's obvious from the logs what happened), you
   should then be able to fix it, or in the case where someone is clearly more
   qualified to fix it, reassign it with very clear instructions.
4. Once you have made a change that you believe fixes a flake, it is conservative
   to keep the issue for the flake open and see if it manifests again after the
   change is merged.
5. If you can't reproduce a flake: __don't just close it!__ Every time a flake comes
   back, at least 2 hours of merge time is wasted. So we need to make monotonic
   progress towards narrowing it down every time a flake occurs. If you can't
   figure it out from the logs, add log messages that would have help you figure
   it out.  If you make changes to make a flake more reproducible, please link
   your pull request to the flake you're working on.
6. If a flake has been open, could not be reproduced, and has not manifested in
   3 months, it is reasonable to close the flake issue with a note saying
   why.
7. If you are unable to deflake the test, consider adding `[Flaky]` to the test
   name, which will result in the test being quarantined to only those jobs that
   explicitly run flakes (eg: https://testgrid.k8s.io/google-gce#gci-gce-flaky)

### Writing a good flake report

If you are reporting a flake, it is important to include enough information for
others to reproduce the issue. When filing the issue, use the
[flaking test template](https://github.com/kubernetes/kubernetes/issues/new?labels=kind%2Fflake&template=flaking-test.yaml). In
your issue, answer these following questions:

- Is this flaking in multiple jobs? You can search for the flaking test or error
  messages using the
  [Kubernetes Aggregated Test Results](http://go.k8s.io/triage) tool.
- Are there multiple tests in the same package or suite failing with the same apparent error?

In addition, be sure to include the following information:

- A link to [testgrid](https://testgrid.k8s.io/) history for the flaking test's
  jobs, filtered to the relevant tests 
- The failed test output &mdash; this is essential because it makes the issue searchable
- A link to the triage query
- A link to specific failures
- Be sure to tag the relevant SIG, if you know what it is.

For a good example of a flaking test issue,
[check here](https://github.com/kubernetes/kubernetes/issues/93358).

## Deflaking unit tests

To get started with deflaking unit tests, you will need to first
reproduce the flaky behavior. Start with a simple attempt to just run
the flaky unit test. For example: 

```sh
go test ./pkg/kubelet/config -run TestInvalidPodFiltered
```

Also make sure that you bypass the `go test` cache by using an uncachable
command line option:

```sh
go test ./pkg/kubelet/config -count=1 -run TestInvalidPodFiltered
```

If even this is not revealing issues with the flaky test, try running with
[race detection](https://golang.org/doc/articles/race_detector.html) enabled:

```sh
go test ./pkg/kubelet/config -race -count=1 -run TestInvalidPodFiltered
```

Finally, you can stress test the unit test using the
[stress command](https://godoc.org/golang.org/x/tools/cmd/stress). Install it
with this command: 

```sh
# go version 1.17 and later
go install golang.org/x/tools/cmd/stress@latest

# go version prior to 1.17
go get golang.org/x/tools/cmd/stress
```

Then build your test binary:

```sh
go test ./pkg/kubelet/config -race -c
```

Then run it under stress:

```sh
stress ./config.test -test.run TestInvalidPodFiltered -test.timeout=1m
```

The stress command runs the test binary repeatedly, reporting when it fails. It
will periodically report how many times it has run and how many failures have
occurred. 

You should see output like this:

```
411 runs so far, 0 failures
/var/folders/7f/9xt_73f12xlby0w362rgk0s400kjgb/T/go-stress-20200825T115041-341977266
--- FAIL: TestInvalidPodFiltered (0.00s)
    config_test.go:126: Expected no update in channel, Got types.PodUpdate{Pods:[]*v1.Pod{(*v1.Pod)(0xc00059e400)}, Op:1, Source:"test"}
FAIL
ERROR: exit status 1
815 runs so far, 1 failures
```

Be careful with tests that use the `net/http/httptest` package; they could
exhaust the available ports on your system!

The `-test.timeout=1m` is necessary to detect when a test run gets stuck.
In contrast to `go test` invocations there is no default timeout when
`stress` invokes the test binary directly. Use a timeout that is long
for the test.

## Deflaking integration tests

Integration tests run similarly to unit tests, but they almost always expect a
running `etcd` instance. You should already have `etcd` installed if you have
followed the instructions in the [Development Guide](../development.md). Run
`etcd` in another shell window or tab.

Compile your integration test using a command like this:

```sh
go test -c -race ./test/integration/endpointslice
```

And then stress test the flaky test using the `stress` command:

```sh
stress ./endpointslice.test -test.run TestEndpointSliceMirroring
```

For an example of a failing or flaky integration test,
[read this issue](https://github.com/kubernetes/kubernetes/issues/93496#issuecomment-678375312). 

Sometimes, but not often, a test will fail due to timeouts caused by
deadlocks. This can be tracked down by stress testing an entire package. The way
to track this down is to stress test individual tests in a package. This process
can take extra effort. Try following these steps:

1. Run each test in the package individually to figure out the average runtime.
2. Stress each test individually, bounding the timeout to 100 times the average run time.
3. Isolate the particular test that is deadlocking.
4. Add debug output to figure out what is causing the deadlock.

Hopefully this can help narrow down exactly where the deadlock is occurring,
revealing a simple fix!

## Deflaking e2e tests

A flaky [end-to-end (e2e) test](e2e-tests.md) offers its own set of
challenges. In particular, these tests are difficult because they test the
entire Kubernetes system. This can be both good and bad. It can be good because
we want the entire system to work when testing, but an e2e test can also fail
because of something completely unrelated, such as failing infrastructure or
misconfigured volumes. Be aware that you can't simply look at the title of an
e2e test to understand exactly what is being tested. If possible, look for unit
and integration tests related to the problem you are trying to solve.

### Gathering information

The first step in deflaking an e2e test is to gather information. We capture a
lot of information from e2e test runs, and you can use these artifacts to gather
information as to why a test is failing.

Use the [Prow Status](https://prow.k8s.io/) tool to collect information on
specific test jobs. Drill down into a job and use the **Artifacts** tab to
collect information. For example, with
[this particular test job](https://prow.k8s.io/view/gcs/kubernetes-jenkins/pr-logs/directory/pull-kubernetes-e2e-gce/1296558932902285312),
we can collect the following:

* `build-log.txt`
* In the control plane directory: `artifacts/e2e-171671cb3f-674b9-master/`
  * `kube-apiserver-audit.log` (and rotated files)
  * `kube-apiserver.log`
  * `kube-controller-manager.log`
  * `kube-scheduler.log`
  * And more!

The `artifacts/` directory will contain much more information. From inside the
directories for each node:
- `e2e-171671cb3f-674b9-minion-group-drkr`
- `e2e-171671cb3f-674b9-minion-group-lr2z`
- `e2e-171671cb3f-674b9-minion-group-qkkz`

Look for these files:
* `kubelet.log`
* `docker.log`
* `kube-proxy.log`
* And so forth.

### Filtering and correlating information

Once you have gathered your information, the next step is to filter and
correlate the information. This can require some familiarity with the issue you are tracking
down, but look first at the relevant components, such as the test log, logs for the API
server, controller manager, and `kubelet`.

Filter the logs to find events that happened around the time of the failure and
events that occurred in related namespaces and objects.

The goal is to collate log entries from all of these different files so you can
get a picture of what was happening in the distributed system. This will help
you figure out exactly where the e2e test is failing. One tool that may help you
with this is [k8s-e2e-log-combiner](https://github.com/brianpursley/k8s-e2e-log-combiner)

Kubernetes has a lot of nested systems, so sometimes log entries can refer to
events happening three levels deep. This means that line numbers in logs might
not refer to where problems and messages originate. Do not make any assumptions
about where messages are initiated!

If you have trouble finding relevant logging information or events, don't be
afraid to add debugging output to the test. For an example of this approach,
[see this issue](https://github.com/kubernetes/kubernetes/pull/88297#issuecomment-588607417).

### What to look for

One of the first things to look for is if the test is assuming that something is
running synchronously when it actually runs asynchronously. For example, if the
test is kicking off a goroutine, you might need to add delays to simulate slow
operations and reproduce issues.

Examples of the types of changes you could make to try to force a failure:
  - `time.Sleep(time.Second)` at the top of a goroutine
  - `time.Sleep(time.Second)` at the beginning of a watch event handler
  - `time.Sleep(time.Second)` at the end of a watch event handler
  - `time.Sleep(time.Second)` at the beginning of a sync loop worker
  - `time.Sleep(time.Second)` at the end of a sync loop worker

Sometimes,
[such as in this example](https://github.com/kubernetes/kubernetes/issues/93496#issuecomment-675631856),
a test might be causing a race condition with the system it is trying to
test. Investigate if the test is conflicting with an asynchronous background
process. To verify the issue, simulate the test losing the race by putting a
`time.Sleep(time.Second)` between test steps. 

If a test is assuming that an operation will happen quickly, it might not be
taking into account the configuration of a CI environment. A CI environment will
generally be more resource-constrained and will run multiple tests in
parallel. If it runs in less than a second locally, it could take a few seconds
in a CI environment.

Unless your test is specifically testing performance/timing, don't set tight
timing tolerances. Use `wait.ForeverTestTimeout`, which is a reasonable stand-in
for operations that should not take very long. This is a better approach than
polling for 1 to 10 seconds.

Is the test incorrectly assuming deterministic output? Remember that map iteration in go is
non-deterministic. If there is a list being compiled or a set of steps are being
performed by iterating over a map, they will not be completed in a predictable
order. Make sure the test is able to tolerate any order in a map.

Be aware that if a test is mixing random allocation with static allocation, that
there will be intermittent conflicts.

Finally, if you are using a fake client with a watcher, it can relist/rewatch at any point.
It is better to look for specific actions in the fake client rather than
asserting exact content of the full set.
