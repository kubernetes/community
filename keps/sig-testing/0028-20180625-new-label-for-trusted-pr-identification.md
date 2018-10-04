---
kep-number: 28
title: New label for trusted PR identification
authors:
  - "@matthyx"
owning-sig: sig-testing
participating-sigs:
  - sig-contributor-experience
reviewers:
  - "@fejta"
  - "@cjwagner"
  - "@BenTheElder"
  - "@cblecker"
  - "@stevekuznetsov"
approvers:
  - TBD
editor: TBD
creation-date: 2018-06-25
last-updated: 2018-09-03
status: provisional
---

# New label for trusted PR identification

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non-Goals](#non-goals)
* [Proposal](#proposal)
  * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
  * [Benefits](#benefits)
  * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Future evolutions](#future-evolutions)
* [References](#references)
* [Implementation History](#implementation-history)

## Summary

This document describes a major change to the way the `trigger` plugin determines if test jobs should be started on a pull request (PR).

We propose introducing a new label named `ok-to-test` that will be applied on non-member PRs once they have been `/ok-to-test` by a legitimate reviewer.

## Motivation

PR test jobs are started by the trigger plugin on *trusted PR* events, or when a *untrusted* PR becomes *trusted*.
> A PR is considered trusted if the author is a member of the *trusted organization* for the repository or if such a member has left an `/ok-to-test` command on the PR.

It is easy spot an untrusted PR opened by a non-member of the organization by its `needs-ok-to-test` label. However the contrary is difficult and involves scanning every comment for a `/ok-to-test`, which increases code complexity and API token consumption.

### Goals

This KEP will only target PRs authored from non-members of the organization:

* introduce a new `ok-to-test` label
* modify `/ok-to-test` command to apply `ok-to-test` on success
* modify `trigger` plugin and other tools to use `ok-to-test` for PR trust
* support `/ok-to-test` calls inside review comments

### Non-Goals

This KEP will not change the current process for members of the organization.

## Proposal

We suggest introducing a new label named `ok-to-test` that would be required on any non-member PR before automatic test jobs can be started by the `trigger` plugin.

This label will be added by members of the *trusted organization* for the repository using the `/ok-to-test` command, detected with a single GenericCommentEvent handler on corresponding events (issue_comment, pull_request_review, and pull_request_review_comment).

### Implementation Details/Notes/Constraints

1. PR: declare `ok-to-test`
   * add `ok-to-test` to `label_sync`
1. (custom tool needed) batch add `ok-to-test` label to non-members trusted PRs
   * for all PR without `ok-to-test` or `needs-ok-to-test`
   * if author is not a member of trusted org
   * add `ok-to-test`
1. PR: switch to `ok-to-test`
   * remove `needs-ok-to-test` from `missingLabels` in `prow/config.yaml`
   * edit `prow/config/jobs_test.go`
   * edit `prow/cmd/deck/static/style.css`
   * edit `prow/cmd/tide/README.md`
   * code changes in `trigger`:
      * `/ok-to-test` adds `ok-to-test`
      * PR trust relies on `ok-to-test`
      * if PR has both labels, drop `needs-ok-to-test`
      * edit all references to `needs-ok-to-test`
1. run batch job again, to catch new PRs that arrived between first run and merge/deploy
1. (to be discussed) periodically check for and report PRs with both `ok-to-test` and `needs-ok-to-test` labels

### Benefits

* Trusted PRs are easily identified by either being authored by org members, or by having the `ok-to-test` label.
* Race conditions can no longer happen when checking if a PR is trusted.
* API tokens are saved by avoiding listing the comments, reviews, and review comments every time we need to check if a PR is trusted.

### Risks and Mitigations

TODO

## Graduation Criteria

TODO

## Future evolutions

In the future, we might decide to require the new label for all PRs, which means that organization members will also need the `ok-to-test` label applied to their PRs before automatic testing can be triggered.

Trusted and untrusted PRs will be even easier to tell apart.

This would require adding automatically the `ok-to-test` label to member authored PRs to keep the current functionality.

## References

* https://github.com/kubernetes/test-infra/issues/3827
* https://github.com/kubernetes/test-infra/issues/7801
* https://github.com/kubernetes/test-infra/pull/5246

## Implementation History

* 2018-06-25: creation of the KEP
* 2018-07-09: KEP content LGTM during sig-testing presentation
* 2018-07-24: KEP updated to keep `needs-ok-to-test` for better UX
* 2018-09-03: KEP rewritten with template