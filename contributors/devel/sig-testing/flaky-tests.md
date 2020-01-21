# Flaky tests

Any test that fails occasionally is "flaky". Since our merges only proceed when
all tests are green, and we have a number of different CI systems running the
tests in various combinations, even a small percentage of flakes results in a
lot of pain for people waiting for their PRs to merge.

Therefore, it's important we take flakes seriously. We should avoid flakes by
writing our tests defensively. When flakes are identified, we should prioritize
addressing them, either by fixing them or quarantining them off the critical
path.

# Avoiding Flakes

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

# Quarantining Flakes

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

# Hunting Flakes

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
- [velodrome.k8s.io] - dashboards driven by the results of queries run against test results using bigquery
  - http://velodrome.k8s.io/dashboard/db/job-health-merge-blocking?orgId=1 - includes flake rate and top flakes for merge-blocking jobs for kubernetes/kubernetes
  - http://velodrome.k8s.io/dashboard/db/job-health-release-blocking?orgId=1 - includes flake rate and top flakes for release-blocking jobs for kubernetes/kubernetes
- [`kind/flake` github query][flake] - open issues or PRs related to flaky jobs or tests for kubernetes/kubernetes

[go.k8s.io/triage]: https://go.k8s.io/triage
[testgrid.k8s.io]: https://testgrid.k8s.io
[velodrome.k8s.io]: http://velodrome.k8s.io

# GitHub Issues for Known Flakes

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

## Expectations when a flaky test is assigned to you

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

# Reproducing unit test flakes

Try the [stress command](https://godoc.org/golang.org/x/tools/cmd/stress).

Just

```
$ go install golang.org/x/tools/cmd/stress
```

Then build your test binary

```
$ go test -c -race
```

Then run it under stress

```
$ stress ./package.test -test.run=FlakyTest
```

It runs the command and writes output to `/tmp/gostress-*` files when it fails.
It periodically reports with run counts. Be careful with tests that use the
`net/http/httptest` package; they could exhaust the available ports on your
system!

# Hunting flaky unit tests in Kubernetes

Sometimes unit tests are flaky. This means that due to (usually) race
conditions, they will occasionally fail, even though most of the time they pass.

We have a goal of 99.9% flake free tests. This means that there is only one
flake in one thousand runs of a test.

Running a test 1000 times on your own machine can be tedious and time consuming.
Fortunately, there is a better way to achieve this using Kubernetes.

_Note: these instructions are mildly hacky for now, as we get run once semantics
and logging they will get better_

There is a testing image `brendanburns/flake` up on the docker hub. We will use
this image to test our fix.

Create a replication controller with the following config:

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: flakecontroller
spec:
  replicas: 24
  template:
    metadata:
      labels:
        name: flake
    spec:
      containers:
      - name: flake
        image: brendanburns/flake
        env:
        - name: TEST_PACKAGE
          value: pkg/tools
        - name: REPO_SPEC
          value: https://github.com/kubernetes/kubernetes
```

Note that we omit the labels and the selector fields of the replication
controller, because they will be populated from the labels field of the pod
template by default.

```sh
kubectl create -f ./controller.yaml
```

This will spin up 24 instances of the test. They will run to completion, then
exit, and the kubelet will restart them, accumulating more and more runs of the
test.

You can examine the recent runs of the test by calling `docker ps -a` and
looking for tasks that exited with non-zero exit codes. Unfortunately, docker
ps -a only keeps around the exit status of the last 15-20 containers with the
same image, so you have to check them frequently.

You can use this script to automate checking for failures, assuming your cluster
is running on GCE and has four nodes:

```sh
echo "" > output.txt
for i in {1..4}; do
  echo "Checking kubernetes-node-${i}"
  echo "kubernetes-node-${i}:" >> output.txt
  gcloud compute ssh "kubernetes-node-${i}" --command="sudo docker ps -a" >> output.txt
done
grep "Exited ([^0])" output.txt
```

Eventually you will have sufficient runs for your purposes. At that point you
can delete the replication controller by running:

```sh
kubectl delete replicationcontroller flakecontroller
```

If you do a final check for flakes with `docker ps -a`, ignore tasks that
exited -1, since that's what happens when you stop the replication controller.

Happy flake hunting!

