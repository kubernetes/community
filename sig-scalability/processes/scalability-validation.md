# Scalability validation of the release

_by Shyam JVS, Google Inc_

**July 2017**

## Introduction

Scalability is an integral part of k8s. Numerous issues were identified during release 1.7 while running tests on large clusters (2k-5k nodes). A majority of these manifested only in large clusters - more info under umbrella issue [#47344]. The issues ranged from large cluster setup (both cloud-provider specific and independent ones) to test-infra problems to previously uncaught regressions with performance and resource usage and so on.

We started [supporting] 5000-node clusters from k8s 1.6 and need to make sure this scale is supported in future releases too. We did it this time by running performance and correctness tests under different configurations manually. But this testing procedure needs to be automated, concrete & well-maintained.

## Goals

In this document, we address the following process-related problems wrt scale testing:

- Automate large cluster tests with a reasonable frequency
- Concretely define the testing configuration used for the release
- Make scalability validation a release prerequisite
- Clearly lay down responsibilities for scale tests

## Non-Goals

This document does not intend to:

- Define the set of tests that comprise scalability and correctness suite
- Define SLIs/SLOs (that’s discussed [here]) and thresholds for the tests
- Discuss particular performance issues we’re facing wrt different releases

While these are interesting from scalability perspective, they are not testing process issues but details of the tests themselves. And can change without requiring much changes in the process.

## Proposed design

For each of the problems above we propose solutions and discuss some caveats. This is an open discussion and better solutions may evolve in future.

### Automate large cluster tests

Currently, there are 2 kinds of tests needed to validate the claim “kubernetes supports X-node clusters”. These are:

- Correctness tests - e2e tests verifying expected system behaviours
- Performance tests - e2e tests for stress-testing the system perf ([Feature:Performance])

We need to run them on 5k-node clusters, but they’re:

- Time-consuming (12-24 hrs)
- Expensive (tens of thousands of core hours per run)
- Blocking other large tests (quota limitations + only one large test project available viz. 'kubernetes-scale')

So we don’t want to run them too frequently. On the other hand, running them too infrequently means 
late identification and piling up of regressions. So we choose the following middleground. \
(**B** = release-blocking job, all times in UTC)


| Day | |
| ------------- |:-------------:| 
| Mon | GCE 5k-node correctness (**B**)  @ 03:01 AM UTC <br /> GCE 5k-node performance (**B**) @ 08:01 AM UTC |
| Tue | GCE 5k-node correctness (**B**)  @ 03:01 AM UTC <br /> GCE 5k-node performance (**B**) @ 08:01 AM UTC |
| Wed | GCE 5k-node correctness (**B**)  @ 03:01 AM UTC <br /> GCE 5k-node performance (**B**) @ 08:01 AM UTC |
| Thu | GCE 5k-node correctness (**B**)  @ 03:01 AM UTC <br /> GCE 5k-node performance (**B**) @ 08:01 AM UTC |
| Fri | GCE 5k-node correctness (**B**)  @ 03:01 AM UTC <br /> GCE 5k-node performance (**B**) @ 08:01 AM UTC |
| Sat | GKE 5k-node correctness @ 03:01 AM UTC <br /> GKE 5k-node performance @ 08:01 AM UTC |
| Sun | GKE 2k-node performance @ 08:01 AM UTC <br /> GKE 2k-node performance (regional) @ 08:01 AM UTC |

Note: The above schedule is subject to change based on job health, release requirements, etc.

Why this schedule?

- 5k tests might need special attention in case of failures so they should mostly run on weekdays.
- Running a large-scale performance job and a large-scale correctness job each day would:
  - help catch regressions on a daily basis
  - help verify fixes with low latency
  - ensure a good release signal
- Running large scale tests on GKE once a week would help verify GKE setup also, at no real loss of signal ideally

Why run GKE tests at all?

Google is currently using a single project for scalability testing, on both GCE and GKE. As a result we need to schedule them together. There's a plan for CNCF becoming responsible for funding k8s testing, and GCE/GKE tests would be separated to different projects when that happens, with only GCE being funded by them. This ensures fairness across all cloud providers.

### Concretely define test configuration

This is a relatively minor issue but it is important that we clearly define the test configuration we use for the release. E.g. there was a confusion this time around testing k8s services, machine-type and no. of the nodes we used (we tested 4k instead of 5k due to a CIDR-setup problem). For ref - [#47344] [#47865]. To solve this, we need to document it using the below template in a file named scalability-validation-report.md placed under kubernetes/features/release-&gt;N&lt;. And this file should be linked from under the scalability section in the release's CHANGELOG.md.

```
Validated large cluster performance under the following configuration:
- Cloud-provider - [GCE / GKE / ..]
- No. of nodes - [5k (desired) / 4k / ..]
- Node size, OS, disk size/type
- Master size, OS, disk size/type
- Any non-default config used - (monitoring with stackdriver, logging with elasticsearch, etc)
- Any important test details - (services disabled in load test, pods increased in density test, etc)
- <job-name, run#> of the validating run (to know other specific details from the logs)

Validated large cluster correctness under the following configuration:
- <similar to above>

Misc:
<Any important scalability insights/issues/improvements in the release>
```

### Make scalability validation a release prerequisite

The model we followed this time was to create an umbrella issue ([#47344]) for scalability testing and labeling it as a release-blocker. While it helped block the release, it didn’t receive enough traction from individual SIGs as scale tests were not part of the release-blocking suite. As a result, the onus for fixing issues fell almost entirely on sig-scalability due to time constraints. Thus the obvious requirement here is to make the 5k-node tests release blockers. This along with test automation ensures failures are identified quickly and get traction from relevant SIGs.

### Clearly lay down responsibilities for scale tests

Responsibilities that lie with sig-scalability:

- Setting and tuning the schedule of the tests on CI
- Ensuring the project is healthy and quotas are sufficient
- Documenting the testing configuration for the release
- Identifying and fixing performance test failures (triaging/delegating as needed)

Responsibilities lying with other SIGs/teams as applicable (could be sig-scalability too):

- Fixing failing correctness test - Owner/SIG for the e2e test (as specified by [test_owners.csv])
- Fixing performance regression - Owner/SIG of relevant component (as delegated by sig-scalability)
- Testing infrastructure issues - @sig-testing
- K8s-specific large-cluster setup issues - @sig-cluster-lifecycle
- GKE-specific large-cluster setup issues - @goog-gke


[#47344]: https://github.com/kubernetes/kubernetes/issues/47344
[supporting]: http://blog.kubernetes.io/2017/03/scalability-updates-in-kubernetes-1.6.html
[here]: https://docs.google.com/document/d/15rD6XBtKyvXXifkRAsAVFBqEGApQxDRWM3H1bZSBsKQ
[#47865]: https://github.com/kubernetes/kubernetes/issues/47865
[test_owners.csv]: https://github.com/kubernetes/kubernetes/blob/master/test/test_owners.csv
[calendar]: https://calendar.google.com/calendar?cid=Z29vZ2xlLmNvbV9tNHA3bG1jODVubGlmazFxYzRnNTRqZjg4a0Bncm91cC5jYWxlbmRhci5nb29nbGUuY29t

