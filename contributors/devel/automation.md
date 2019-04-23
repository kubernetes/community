# Kubernetes Development Automation

## Overview

Kubernetes uses a variety of automated tools in an attempt to relieve developers
of repetitive, low brain power work. This document attempts to describe these
processes.

## Tide

This project formerly used a Submit Queue, it has since been replaced by
[Tide](https://git.k8s.io/test-infra/prow/cmd/tide).

#### Ready to merge status

A PR is considered "ready for merging" by Tide if it matches the set
of conditions listed in the [Tide dashboard](https://prow.k8s.io/tide).
Please visit that page for more details.

### Closing stale pull-requests

Prow will close pull-requests that don't have human activity in the
last 90 days. It will warn about this process 60 days before closing the
pull-request, and warn again 30 days later. One way to prevent this from
happening is to add the `lifecycle/frozen` label on the pull-request.

Feel free to re-open and maybe add the `lifecycle/frozen` label if this happens to a
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

#### How can I ask my PR to be tested again for test failures?

PRs should only need to be manually re-tested if you believe there was a flake
during the original test. It would be good to file flakes as an
[issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fflake). 
`@k8s-ci-robot` will comment to tell you which test(s) failed and how to re-test. 
The simplest way is to comment `/retest`.

Any pushes of new code to the PR will automatically trigger a new test. No human
interaction is required. Note that if the PR has a `lgtm` label, it will be removed after the pushes.
