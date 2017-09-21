# Kubernetes Development Automation

## Overview

Kubernetes uses a variety of automated tools in an attempt to relieve developers
of repetitive, low brain power work. This document attempts to describe these
processes.


## Submit Queue

In an effort to
   * reduce load on core developers
   * maintain e2e stability
   * load test github's label feature

We have added an automated [submit-queue](https://github.com/kubernetes/test-infra/tree/master/mungegithub/submit-queue)
to the
[github "munger"](https://github.com/kubernetes/test-infra/tree/master/mungegithub)
for kubernetes.

The submit-queue does the following:

```go
for _, pr := range readyToMergePRs() {
    if testsAreStable() {
        if retestPR(pr) == success {
            mergePR(pr)
        }
    }
}
```

The status of the submit-queue is [online.](http://submit-queue.k8s.io/)

### Ready to merge status

The submit-queue lists what it believes are required on the [merge requirements tab](http://submit-queue.k8s.io/#/info) of the info page. That may be more up to date.

A PR is considered "ready for merging" by the submit queue if it matches the set
of conditions listed in the merge requirements tab mentioned above.
Please visit that page for more details.

### Merge process

If the PR has the `retest-not-required` label, it is simply merged. If the PR does
not have this label, the aforementioned required tests are re-run.
If these tests pass a second time, the PR will be merged when this PR finishes retesting.

## Github Munger

We run [github "mungers"](https://github.com/kubernetes/test-infra/tree/master/mungegithub).

This runs repeatedly over github pulls and issues and runs modular "mungers"
similar to "mungedocs". The mungers include the "submit-queue" referenced above along
with numerous other functions. See the README in the link above.

Please feel free to unleash your creativity on this tool, send us new mungers
that you think will help support the Kubernetes development process.

### Closing stale pull-requests

Github Munger will close pull-requests that don't have human activity in the
last 90 days. It will warn about this process 60 days before closing the
pull-request, and warn again 30 days later. One way to prevent this from
happening is to add the `keep-open` label on the pull-request.

Feel free to re-open and maybe add the `keep-open` label if this happens to a
valid pull-request. It may also be a good opportunity to get more attention by
verifying that it is properly assigned and/or mention people that might be
interested. Commenting on the pull-request will also keep it open for another 90
days.

## PR builder

We also run a robotic PR builder that attempts to run tests for each PR.

Before a PR from an unknown user is run, the PR builder bot (`@k8s-ci-robot`) asks to
a message from a kubernetes member that a PR is safe to test, the member can
reply with the `/ok-to-test` command on a single line to begin CI testing.

## FAQ:

#### How can I ask my PR to be tested again for Jenkins failures?

PRs should only need to be manually re-tested if you believe there was a flake
during the original test. It would be good to file flakes as an
[issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fflake). 
`@k8s-ci-robot` will comment to tell you which test(s) failed and how to re-test. 
The simplest way is to comment `/retest`.

Any pushes of new code to the PR will automatically trigger a new test. No human
interraction is required. Note that if the PR has a `lgtm` label, it will be removed after the pushes.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/automation.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
