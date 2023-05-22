# Overview

This document explains how cherry picks are managed on release branches within
the `kubernetes/kubernetes` repository.
A common use case for this task is backporting PRs from master to release
branches.

- [Prerequisites](#prerequisites)
- [What Kind of PRs are Good for Cherry Picks](#what-kind-of-prs-are-good-for-cherry-picks)
- [Initiate a Cherry Pick](#initiate-a-cherry-pick)
- [Cherry Pick Review](#cherry-pick-review)
- [Searching for Cherry Picks](#searching-for-cherry-picks)
- [Troubleshooting Cherry Picks](#troubleshooting-cherry-picks)
- [Cherry Picks for Unsupported Releases](#cherry-picks-for-unsupported-releases)

## Prerequisites

- [Contributor License Agreement](http://git.k8s.io/community/CLA.md) is
  considered implicit for all code within cherry pick pull requests,
  **unless there is a large conflict**.
- A pull request merged against the `master` branch.
- The release branch exists (example: [`release-1.18`](https://github.com/kubernetes/kubernetes/tree/release-1.18))
- The normal git and GitHub configured shell environment for pushing to your
  kubernetes `origin` fork on GitHub and making a pull request against a
  configured remote `upstream` that tracks
  `https://github.com/kubernetes/kubernetes.git`, including `GITHUB_USER`.
- Have GitHub CLI (`gh`) installed following [installation instructions](https://github.com/cli/cli#installation).
- A github personal access token which has permissions "repo" and "read:org".
  Permissions are required for [gh auth login](https://cli.github.com/manual/gh_auth_login)
  and not used for anything unrelated to cherry-pick creation process
  (creating a branch and initiating PR).

## What Kind of PRs are Good for Cherry Picks

Compared to the normal master branch's merge volume across time,
the release branches see one or two orders of magnitude less PRs.
This is because there is an order or two of magnitude higher scrutiny.
Again, the emphasis is on critical bug fixes, e.g.,

- Loss of data
- Memory corruption
- Panic, crash, hang
- Security

A bugfix for a functional issue (not a data loss or security issue) that only
affects an alpha feature does not qualify as a critical bug fix.

If you are proposing a cherry pick and it is not a clear and obvious critical
bug fix, please reconsider. If upon reflection you wish to continue, bolster
your case by supplementing your PR with e.g.,

- A GitHub issue detailing the problem

- Scope of the change

- Risks of adding a change

- Risks of associated regression

- Testing performed, test cases added

- Key stakeholder SIG reviewers/approvers attesting to their confidence in the
  change being a required backport

If the change is in cloud provider-specific platform code (which is in the
process of being moved out of core Kubernetes), describe the customer impact,
how the issue escaped initial testing, remediation taken to prevent similar
future escapes, and why the change cannot be carried in your downstream fork of
the Kubernetes project branches.

It is critical that our full community is actively engaged on enhancements in
the project. If a released feature was not enabled on a particular provider's
platform, this is a community miss that needs to be resolved in the `master`
branch for subsequent releases. Such enabling will not be backported to the
patch release branches.

## Initiate a Cherry Pick

### Before you begin

- Plan to initiate a cherry-pick against _every_ supported release branch. If you decide to skip some release branch, explain your decision in a comment to the PR being cherry-picked.

- Initiate cherry-picks in order, from newest to oldest supported release branches. For example, if 1.27 is the newest supported release branch, then, before cherry-picking to 1.25, make sure the cherry-pick PR already exists for in 1.26 and 1.27. This helps to prevent regressions as a result of an upgrade to the next release.

### Steps

- Run the [cherry pick script][cherry-pick-script]

  This example applies a master branch PR #98765 to the remote branch
  `upstream/release-3.14`:

  ```shell
  hack/cherry_pick_pull.sh upstream/release-3.14 98765
  ```

  - Be aware the cherry pick script assumes you have a git remote called
    `upstream` that points at the Kubernetes github org.

    Please see our [recommended Git workflow](/contributors/guide/github-workflow.md#workflow).

  - You will need to run the cherry pick script separately for each patch
    release you want to cherry pick to. Cherry picks should be applied to all
    [active](https://github.com/kubernetes/website/blob/main/content/en/releases/patch-releases.md#detailed-release-history-for-active-branches)
    release branches where the fix is applicable.

  - If `GITHUB_TOKEN` is not set you will be asked for your github password:
    provide the github [personal access token](https://github.com/settings/tokens) rather than your actual github
    password. If you can securely set the environment variable `GITHUB_TOKEN`
    to your personal access token then you can avoid an interactive prompt.
    Refer [https://github.com/github/hub/issues/2655#issuecomment-735836048](https://github.com/github/hub/issues/2655#issuecomment-735836048)


- Your cherry pick PR will immediately get the
  `do-not-merge/cherry-pick-not-approved` label.

  [Normal rules apply for code merge](/contributors/devel/sig-release/release.md#tldr),
  with some additional caveats outlined in the next section of this document.

## Cherry Pick Review

As with any other PR, code OWNERS review (`/lgtm`) and approve (`/approve`) on
cherry pick PRs as they deem appropriate.

The same release note requirements apply as normal pull requests, except the
release note stanza will auto-populate from the master branch pull request from
which the cherry pick originated.

If this is unsuccessful, the `do-not-merge/release-note-label-needed` label
will be applied and the cherry pick author must edit the pull request
description to [add a release note](/contributors/guide/release-notes.md) or
include in a comment the `/release-note-none` command.

Cherry pick pull requests are reviewed slightly differently than normal
pull requests on the `master` branch in that they:

- Are by default expected to be `kind/bug` and `priority/critical-urgent`.

- The original change to the `master` branch is expected to be merged for
  some time and no related CI failures or test flakiness must be discovered.

- The easy way to compare changes from the original change and cherry-pick
  is to compare PRs `.patch` files. To generate the patch from
  PR, just add the `.patch` to PR url. For example, for PR #100972 in
  kubernetes repositry, ptach can be downloaded following this URL:

  `https://github.com/kubernetes/kubernetes/pull/100972.patch`

- Milestones must be set on the PR reflecting the milestone for the target
  release branch (for example, milestone v1.11 for a cherry pick onto branch
  `release-1.11`). This is normally done for you by automation.

- A separate cherry pick pull request should be open for every applicable target
  branch. This ensures that the fix will be present on every active branch for a
  given set of patch releases. If a fix is only applicable to a subset of active
  branches, it is helpful to note why that is the case on the parent pull
  request or on the cherry pick pull requests to the applicable branches.

- Have one additional level of review in that they must be approved
  specifically for cherry pick by branch approvers.

  The [Release Managers][release-managers] are the final approvers on release
  branches.

  Approval is signified by a Release Manager manually applying the
  `cherry-pick-approved` label. This action removes the
  `do-not-merge/cherry-pick-not-approved` label and triggers a merge into the
  target branch.

  The team scrubs through incoming cherry picks on at least a weekly basis,
  daily during burndown ahead of a .0 release. Ahead of point releases,
  reminders of the cherry pick deadline will be sent out to the community.
  Cherry pick PRs are often metered into the release branches to give more
  deliberate CI signal across changes. For this reason your cherry pick must be
  ready to merge ahead of the cherry pick deadline, but those candidates may be
  merged during the days between the deadline and release.

  Open cherry pick PRs which do not land in the current release will continue
  to be tracked by the team for consideration for inclusion in a next patch
  release.

  If you are concerned about the status of your cherry pick, err on the
  side of overcommunicating and reach out to the
  [Release Managers][release-managers].

## Searching for Cherry Picks

Examples (based on cherry picks targeting the `release-1.18` branch):

- [`cherry-pick-approved`](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Acherry-pick-approved+base%3Arelease-1.18)
- [`do-not-merge/cherry-pick-not-approved`](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Ado-not-merge%2Fcherry-pick-not-approved+base%3Arelease-1.18)

## Troubleshooting Cherry Picks

Contributors may encounter some of the following difficulties when initiating a
cherry pick.

- A cherry pick PR does not apply cleanly against an old release branch. In
  that case, you will need to manually fix conflicts.

- The cherry pick PR includes code that does not pass CI tests. In such a case
  you will have to fetch the auto-generated branch from your fork, amend the
  problematic commit and force push to the auto-generated branch.
  Alternatively, you can create a new PR, which is noisier.

## Cherry Picks for Unsupported Releases

The community supports & patches releases for approximately 1 year
for releases 1.19 and newer.  For releases 1.18 and older the patch
support extended for approximately 9 months, which was derived from
keeping `n-3` releases (`n` being the latest -release of Kubernetes)
in support and a quarterly release cycle.

The community makes no guarantees, but in the event of a high
severity issue with a patch that is backportable and can be proved
with CI signal, this extra support may occasionally be given.

For example, in January of 2019 the community discovered a regression, that was
introduced in a post-release patch, but was currently no longer
supported.  As discussed in a SIG Release meeting on 2019-01-15, a
fix was backported to the non supported version.

Reference PR: [#72860](https://github.com/kubernetes/kubernetes/pull/72860)

The specific criteria driving the decision was:

- CI was still available for the version
- The regression was introduced as a patch (and not part of the official
  release)
- The issue being fixed is of sufficient **[severity & impact](#what-kind-of-prs-are-good-for-cherry-picks)**
- The fix is well understood and contained (doesn’t introduce risk of
  additional regressions)

A note about the specific case in [#72860](https://github.com/kubernetes/kubernetes/pull/72860#issuecomment-454072746):

- The patch was exceedingly tiny and very unlikely to introduce new problems
- Luckily, it was caught shortly after the release was supposed to be
  unsupported

[cherry-pick-script]: https://git.k8s.io/kubernetes/hack/cherry_pick_pull.sh
[release-managers]: https://kubernetes.io/releases/release-managers/
