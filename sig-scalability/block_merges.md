# Blocking PR merges in the event of regression.

As mentioned in the charter, SIG scalability has a right to block all PRs
from merging into the relevant repos. This document describes the underlying
"Rules of engagement" of this process and the rationale why this is needed.

### Rules of engagement.
The rules of engagement for blocking merges are as following:

- Observe as scalability regression on one of release-blocking test suites
  (defined as green to red transition - if tests were already failing, we
  don't have a right to declare a regression).
- Block merges of all PRs to the relevant repos in the affected branch,
  declaring which repos those are and why.
- Identify the PR which caused the regression:
  - this can be done by reading code changes, bisecting, debugging based on
    metrics and/or logs, etc.
  - we say a PR is identified as the cause when we are reasonably confident
    that it indeed caused a regression, even if the mechanism is not 100%
    understood to minimize the time when merges are blocked
- Mitigate the regression. This may mean e.g.:
  - reverting the PR
  - switching a feature off (preferably by default, as last resort only in tests)
  - fixing the problem (if it's easy and quick to fix)
- Unblock merges of all PRs to the relevant repos in the affected branch.

The exact technical mechanisms for it are out of scope for this document.

### Rationale
The process described above is quite drastic, but we believe it is justified
if we want kubernetes to maintain scalability SLOs. The reasoning is:
- reliably testing for regressions takes a lot of time:
  - key scalability e2e tests take too long to execute to be a prerequisite
    for merging all PRs, this is an inherent characteristic of testing at scale,
  - end-to-end tests are flaky (even when not at scale) requiring retries,
- we need to prevent regression pile-ups:
  - once a regression is merged, and no other action is taken, it is only
    a matter of time until another regression is merged on top of it,
  - debugging the cause of two simultaneous (piled-up) regressions is 
    exponentially harder, see issue [53255](http://pr.k8s.io/53255) which
    links to past experience
- we need to keep flakiness of merge-blocking jobs very low:
- regarding benchmarks, there were several scalability issues in the past
  caught by (costly) large-scale e2e tests, which could have been caught and
  fixed earlier and with far less human effort if we had benchmark-like
  tests. Examples include:
  - scheduler anti-affinity affecting kube-dns,
  - kubelet network plugin increasing pod-startup latency,
  - large responses from apiserver violating gRPC MTU.

As explained in detail in an issue, not being able to maintain passing scalability
tests adversely affect:
- release quality
- release schedule
- engineer productivity
