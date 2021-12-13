---
title: Code Review Process for the Kubernetes Project
slug: review-processes
description: |
  Details of the code review process, from the perspective of someone performing
  the review.
---

Most projects in Kubernetes follow this code review process.

This pages gives you a simplified description of our [full PR testing and merge workflow][pr-workflow]
that conveniently forgets about the existence of tests, to focus solely on the roles taken
by reviewers and by code owners.

[OWNERS Files](https://www.kubernetes.dev/docs/guide/owners/) explains how `OWNERS` files work
in Kubernetes.

## Code review process

- The **author** submits a pull request (PR).
- Phase 0: Automation suggests **[reviewers][reviewer-role]** and **[approvers][approver-role]** for the PR
  - Determine the set of OWNERS files nearest to the code being changed;
  - Choose at least two suggested **reviewers**, trying to find a unique reviewer for every leaf
    OWNERS file, and request their reviews on the PR;
  - Choose suggested **approvers**, one from each OWNERS file, and list them in a comment on the PR.
- Phase 1: Humans review the PR
  - **Reviewers** look for general code quality, correctness, sane software engineering, style, etc.
  - Anyone in the organization can act as a **reviewer** - with the exception of the individual who
    opened the PR.
  - If the code changes look good to them, a **reviewer** types `/lgtm` in a PR comment or review;
    if they change their mind, they `/lgtm cancel`.
  - Once a **reviewer** has `/lgtm`'ed, [prow](https://prow.k8s.io)
    ([@k8s-ci-robot](https://github.com/k8s-ci-robot/)) applies an `lgtm` label to the PR.
- Phase 2: Humans approve the PR
  - The PR **author** `/assign`'s all suggested **approvers** to the PR, and optionally notifies
    them (eg: "pinging @foo for approval");
  - Only people listed in the relevant OWNERS files, either directly or through an alias,
    can act as **approvers**, including the individual who opened the PR.
    See [`OWNERS_ALIASES`](https://www.kubernetes.dev/docs/guide/owners/#owners_aliases) for details
    about aliases.
  - **Approvers** look for holistic acceptance criteria, such as: dependencies with other features,
    forwards/backwards compatibility, API and flag definitions.
  - If the code changes look good to them, an **approver** types `/approve` in a PR comment or
    review; if they change their mind, they `/approve cancel`;
  - [prow](https://prow.k8s.io) ([@k8s-ci-robot](https://github.com/k8s-ci-robot/)) updates its
    comment in the PR to indicate which **approvers** still need to approve'
  - Once all **approvers** (one from each of the previously identified OWNERS files) have approved,
    [prow](https://prow.k8s.io) ([@k8s-ci-robot](https://github.com/k8s-ci-robot/)) applies an
    `approved` label.
- Phase 3: Automation merges the PR:
  - If all of the following are true:
    - All required labels are present (eg: `lgtm`, `approved`)
    - Any blocking labels are missing (eg: there is no `do-not-merge/hold`, `needs-rebase`)
  - And if any of the following are true:
    - there are no presubmit prow jobs configured for this repo
    - there are presubmit prow jobs configured for this repo, and they all pass after automatically
      being re-run one last time
  - Then the PR will automatically be merged.

### Quirks of the process

There are a number of behaviors we've observed that while _possible_ are discouraged, as they go
against the intent of this review process.  Some of these could be prevented in the future, but this
is the state of today.

- An **approver**'s `/lgtm` is simultaneously interpreted as an `/approve`
  - While a convenient shortcut for some, it can be surprising that the same command is interpreted
    in one of two ways depending on who the commenter is
  - Instead, explicitly write out `/lgtm` and `/approve` to help observers, or save the `/lgtm` for
    a **reviewer**
  - This goes against the idea of having at least two sets of eyes on a PR, and may be a sign that
    there are too few **reviewers** (who aren't also **approver**)
- Technically, anyone who is a member of the Kubernetes GitHub organization can drive-by `/lgtm` a
  PR
  - Drive-by reviews from non-members are encouraged as a way of demonstrating experience and
    intent to become a member or reviewer.
  - Drive-by `/lgtm`'s from members may be a sign that our OWNERS files are too small, or that the
    existing **reviewers** are too unresponsive
  - This goes against the idea of specifying **reviewers** in the first place, to ensure that
    **author** is getting actionable feedback from people knowledgeable with the code
- **Reviewers**, and **approvers** are unresponsive
  - This causes a lot of frustration for **authors** who often have little visibility into why their
    PR is being ignored
  - Many **reviewers** and **approvers** are so overloaded by GitHub notifications that @mention'ing
    is unlikely to get a quick response
  - If an **author** `/assign`'s a PR, **reviewers** and **approvers** will be made aware of it on
    their [PR dashboard](https://gubernator.k8s.io/pr)
  - An **author** can work around this by manually reading the relevant OWNERS files,
    `/unassign`'ing unresponsive individuals, and `/assign`'ing others
  - This is a sign that our OWNERS files are stale; pruning the **reviewers** and **approvers** lists
    would help with this
- **Authors** are unresponsive
  - This costs a tremendous amount of attention as context for an individual PR is lost over time
  - This hurts the project in general as its general noise level increases over time
  - Instead, close PR's that are untouched after too long (we currently have a bot do this after 90
    days)

## Prow commands for reviewing

[Prow](https://github.com/kubernetes/test-infra/blob/master/prow/README.md) is
the Kubernetes-based CI/CD system that runs jobs against pull requests (PRs). Prow
enables chatbot-style commands to handle GitHub actions across the Kubernetes
organization, like [adding and removing labels](#adding-and-removing-issue-labels),
closing issues, and assigning an approver. Enter Prow commands as GitHub comments using the
`/<command-name>` format.

The most common prow commands reviewers and approvers use are:

Prow Command | Role Restrictions | Description
:------------|:------------------|:-----------
`/lgtm` | Organization members | Signals that you've finished reviewing a PR and are satisfied with the changes.
`/approve` | Approvers | Approves a PR for merging.
`/assign` | Reviewers or Approvers | Assigns a person to review or approve a PR
`/close` | Reviewers or Approvers | Closes an issue or PR.
`/hold` | Organization members | Adds the `do-not-merge/hold` label, indicating the PR cannot be automatically merged. You should also explain why you are adding the hold.
`/hold cancel` | Organization members | Removes the `do-not-merge/hold` label.

See [the Prow command reference](https://prow.k8s.io/command-help) to see the full list
of commands you can use in a PR.

## Committing into another person's PR

Leaving PR comments is helpful, but there might be times when you need to commit
into another person's PR instead.

Do not "take over" for another person unless they explicitly ask
you to, or you have asked them first and they agree. While it may be faster
in the short term, it deprives the person of the chance to contribute.

The process you use depends on whether you need to edit a file that is already
in the scope of the PR, or a file that the PR has not yet touched.

You can't commit into someone else's PR if either of the following things is
true:

- The PR author explicitly disallows edits from approvers.

- If the PR author pushed their branch directly to the upstream repository.
  Only a reviewer with push access can commit to another user's PR.


[pr-workflow]: https://www.kubernetes.dev/docs/guide/pull-requests/#the-testing-and-merge-workflow
