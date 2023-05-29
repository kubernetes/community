# SIG Node Triage Process

Currently, SIG Node is limiting its triage review to pull requests. In the
near-future, we hope to expand to issues as well.

This is intended to guide developers through the SIG Node triage process. You
will learn about the roles of authors, reviewers, and approvers, as well as
what labels are required for each stage of the PR process.

We track most node PRs on the [SIG Node Triage project board]. All test and
CI-related PRs are tracked on the [CI subproject board].

For help with the commands listed in the document below, review the [bot command documentation].

[SIG Node PR Triage project board]: https://github.com/orgs/kubernetes/projects/49
[SIG Node Bugs Triage project board]: https://github.com/orgs/kubernetes/projects/59
[CI subproject board]: https://github.com/orgs/kubernetes/projects/43
[bot command documentation]: https://go.k8s.io/bot-commands

## Triage

All new PRs added to the board should begin in the **Triage** column.

When a pull request is made against kubernetes/kubernetes, it will typically
have the following [labels], applied by the Prow bot:

- needs-ok-to-test (not needed for project members)
- needs-kind
- needs-priority
- needs-triage

In order to be moved out of this column, all of the above labels must be set.

[labels]: https://github.com/kubernetes/test-infra/blob/master/label_sync/labels.md

### needs-ok-to-test

Use your best judgment in determining whether a PR is OK to test. You don't
have to do a full review: just make sure that the code does not appear to be
actively malicious, and the PR appears to be doing something useful. Use the
command `/ok-to-test`.

Only [Kubernetes org members] can add this [label][labels].

If the PR is trivial and doesn't provide much value, feel free to close it
using the `/close` command and link the author to the [trivial edits policy].

You can use the following message:

    /close

    Thank you for your PR. It seems to only contain trivial edits. Please read our [trivial edits policy](https://github.com/kubernetes/community/blob/master/contributors/guide/pull-requests.md#trivial-edits). We encourage you to take a look at confirmed issues and bugs or issues marked `help-wanted`.

[Kubernetes org members]: https://github.com/kubernetes/community/blob/master/community-membership.md#member
[trivial edits policy]: https://github.com/kubernetes/community/blob/master/contributors/guide/pull-requests.md#trivial-edits

### needs-kind

Most authors will already set this, as it's part of the PR template, but they
may not set it correctly. A PR may have multiple "kind" labels, so ensure only
the correct ones are applied.

Anyone can add these [labels].

- **kind/api-change:** API change, that will require special API review
- **kind/bug:** related to a bug
- **kind/cleanup:** cleaning up code, process, or technical debt
- **kind/deprecation:** deprecation, that will require special API review
- **kind/documentation:** related to documentation (including code comments)
- **kind/failing-test:** related to a consistently or frequently failing test
- **kind/feature:** related to a new feature or enhancement; should have an
  associated KEP linked
- **kind/flake:** related to a flaky test
- **kind/regression:** related to a regression in performance or functionality
  from a prior release
- **kind/support:** not applicable to PRs

### needs-priority

You can take a quick look at what the PR is addressing and then apply a
priority label.

Anyone can add this [label][labels].

- **priority/critical-urgent:** Urgent bug fix, required ASAP. If not
  addressed, will block a release. These issues should always be discussed in
  the `#sig-node` channel on Slack.
- **priority/important-soon:** Needs to be completed this release. Important
  bug fixes + KEPs targeted for the current milestone.
- **priority/important-longterm:** Has an attached issue/KEP, but unclear what
  the specific priority is.
- **priority/backlog:** Non-urgent changes such as minor performance
  optimizations, improving error logs, increasing test code coverage, code
  refactoring, and addressing static code analysis issues.

### needs-triage

There are two more things to check before accepting a PR for triage.

The first is whether the SIG is correct. If the PR does not appear to touch SIG
Node code or require a SIG Node approver, you should remove the SIG Node label,
and add other SIG labels as appropriate.

The second is verifying the kind of PR. Most will be bug fixes, cleanups, or
documentation. Feature PRs should generally have an attached KEP. API changes
and deprecations require special review; those labels may be mistakenly
applied, so check over the PR to make sure they're accurate.

Once you've quickly looked over a PR, applied the appropriate labels, and it
looks ready to proceed to review (i.e. it doesn't have any labels that would
mean it's waiting on more work from the author), you can mark the PR as triaged
with `/triage accepted`.

Only [Kubernetes org members] can add this [label][labels].

## Waiting on Author

This column means that the PR is waiting on some action from the author. A
reviewer may have requested changes, or a PR may have one of the following
do-not-merge [labels]:

- **do-not-merge/hold:** usually set by a reviewer
- **do-not-merge/work-in-progress:** usually set by an author
- **do-not-merge/release-note-needed:** needs a release note to be added by the
  author, [Kubernetes org members] can override with `/release-note-none` if
  not required
- **do-not-merge/contains-merge-commits:** PR needs to be rebased
- **needs-rebase:** PR needs to be rebased

If [tests are failing] due to an issue with the change (rather than a [flake]), PRs
should also be assigned to this column.

Authors are encouraged to fix any of the label issues above, resolve or reply
to all PR feedback, and leave a comment indicating when their PR is ready for
review.

PRs that do not have any of the above labeled in this column should be
evaluated occasionally to see if they are ready for review.

If PRs are not updated for a long period (90d), they will be marked as stale.
After 30d more, they will be marked as rotten, and then closed automatically.
Reviewers should feel free to close stale PRs (4+ months of no changes) with a
note that the author can reopen when they are ready to work on it.

[tests are failing]: /contributors/devel/sig-testing/testing.md#troubleshooting-a-failure
[flake]: /contributors/devel/sig-testing/flaky-tests.md

## Waiting on Reviewer

This PR needs review! If you're not sure how to review a PR, start by
familiarizing yourself with the Kubernetes [pull request guidelines] and
[review guidelines].

PRs in this column must have the following [labels] set:

- triage-accepted
- priority
- kind

Only [Kubernetes org members] can add an `/lgtm`.

If you want to become an official Node reviewer, you should read through the
[reviewer responsibilities and requirements].

As part of code review, you should ensure that:

- the change is needed
- the metadata on the PR and the release note are accurate
- the change works the way it is intended to
- alternative implementations have been explored and this is an appropriate
  solution
- the code has been reviewed by everyone it needs to be: it may have an LGTM
  label already, but still needs feedback from Node reviewers

TODO: Add some node-specific stuff here.

[pull request guidelines]: https://github.com/kubernetes/community/blob/master/contributors/guide/pull-requests.md
[review guidelines]: https://github.com/kubernetes/community/blob/master/contributors/guide/review-guidelines.md
[reviewer responsibilities and requirements]: https://github.com/kubernetes/community/blob/master/community-membership.md#reviewer

## Waiting on Approver

These PRs are waiting on an `approved` label that can only be provided by
approvers. The bot will always tell you whose approval is needed on which
directories ([example comment]), and will update its comment as approvals are
provided.

PRs in this column must have the following [labels] set:

- lgtm
- triage-accepted
- priority
- kind

Check for the bot's comment to see which files still need approvers from the
appropriate OWNERS. If the PR already has an approval for the node components
(commonly, anything in `./pkg/kubelet/*`), you can mark the PR as Done manually
while waiting on other approvers.

Only [Kubernetes approvers] can use `/approve` on a PR to address this
requirement.

[example comment]: https://github.com/kubernetes/kubernetes/pull/97992#issuecomment-759450299
[Kubernetes approvers]: https://github.com/kubernetes/community/blob/master/community-membership.md#approver

## Done

This column has automation to include all closed and merged PRs on the board.
Huzzah!

It may also include PRs that have LGTMs and approvals, but are not yet merged
(i.e. requires a non-node approver signoff, API review, or a release team
cherry-pick approval).

TODO: We should archive this column per release.
