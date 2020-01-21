# Overview

This document explains how cherry-picks are managed on release branches within
the kubernetes/kubernetes repository.
A common use case for this task is backporting PRs from master to release 
branches.

-   [Prerequisites](#prerequisites)
-   [What Kind of PRs are Good for Cherry-picks](#what-kind-of-prs-are-good-for-cherry-picks)
-   [Initiate a Cherry-pick](#initiate-a-cherry-pick)
-   [Cherry-pick Review](#cherry-pick-review)
-   [Searching for Cherry-picks](#searching-for-cherry-picks)
-   [Troubleshooting Cherry-picks](#troubleshooting-cherry-picks)
-   [Cherry-picks for unsupported releases](#cherry-picks-for-unsupported-releases)

---

## Prerequisites
 * [Contributor License Agreement](http://git.k8s.io/community/CLA.md) is
   considered implicit for all code within cherry-pick pull requests,
   **unless there is a large conflict**.
 * A pull request merged against the master branch.
 * [Release branch](https://git.k8s.io/release/docs/branching.md) exists.
 * The normal git and GitHub configured shell environment for pushing to your
   kubernetes `origin` fork on GitHub and making a pull request against a 
   configured remote `upstream` that tracks
   "https://github.com/kubernetes/kubernetes.git", including `GITHUB_USER`.
 * Have `hub` installed, which is most easily installed via `go get
   github.com/github/hub` assuming you have a standard golang development
   environment.


## What Kind of PRs are Good for Cherry-Picks

Compared to the normal master branch's merge volume across time,
the release branches see one or two orders of magnitude less PRs.
This is because there is an order or two of magnitude higher scrutiny.
Again the emphasis is on critical bug fixes, eg:
 * Loss of data
 * Memory corruption
 * Panic, crash, hang
 * Security

If you are proposing a cherry-pick and it is not a clear and obvious
critical bug fix, please reconsider.  If upon reflection you wish to
continue, bolster your case by supplementing your PR with, eg:

 * A GitHub issue detailing the problem

 * Scope of the change

 * Risks of adding a change

 * Risks of associated regression

 * Testing performed, test cases added

 * Key stakeholder SIG reviewers/approvers attesting to their confidence in the
   change being a required backport

 * If the change is in cloud-provider-specific platform code (which is in the
   process of being moved out of core Kubernetes), describe the customer impact,
   how the issue escaped initial testing, remediation taken to prevent similar
   future escapes, and why the change cannot be carried in your downstream
   fork of the Kubernetes project branches.  It is critical that our full
   community is actively engaged on enhancements in the project.  If a
   released feature was not enabled on a particular provider's platform, this
   is a community miss that needs to be resolved in the master branch for
   subsequent releases.  Such enabling will not be backported to the patch
   release branches.


## Initiate a Cherry-pick
 * Run the [cherry-pick 
   script](https://git.k8s.io/kubernetes/hack/cherry_pick_pull.sh).
   This example applies a master branch PR #98765 to the remote branch
   `upstream/release-3.14`: `hack/cherry_pick_pull.sh upstream/release-3.14
   98765`
   * Be aware the cherry-pick script assumes you have a git remote called 
   `upstream` that points at the Kubernetes github org.
   Please see our [recommended Git workflow](https://git.k8s.io/community/contributors/guide/github-workflow.md#workflow).
   * You will need to run the cherry-pick script separately for each patch release you want to cherry-pick to.

 * Your cherry-pick PR will immediately get the `do-not-merge/cherry-pick-not-approved` label. 
   [Normal rules apply for code merge](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-release/release.md#tldr),
   with some additional caveats outlined in the next section of this document.

## Cherry-pick Review

As with any other PR, code OWNERS review (`/lgtm`) and approve (`/approve`) on
cherry-pick PRs as they deem appropriate.

The same release note requirements apply as normal pull requests,
except the release note stanza will auto-populate from the master
branch pull request from which the cherry-pick originated.  If this
is unsuccessful the `do-not-merge/release-note-label-needed` label
will be applied and the cherry-pick author must edit the pull request
description to [add a release note](https://git.k8s.io/community/contributors/guide/release-notes.md)
or include in a comment the `/release-note-none` command.

Cherry-pick pull requests are reviewed slightly differently than normal
pull requests on the master branch in that they:

 * Are by default expected to be `kind/bug` and `priority/critical-urgent`.

 * Milestones must be set on the PR reflecting the milestone for the target
   release branch (for example, milestone v1.11 for a cherry-pick onto branch
   release-1.11). This is normally done for you by automation.

 * Have one additional level of review in that they must be approved specifically
   for cherry-pick by branch approvers.

   The [Branch Manager](https://git.k8s.io/sig-release/release-team/role-handbooks/branch-manager)
   will triage PRs targeted to the next .0 minor release branch up until the 
   release, while the [Patch Release Team](https://git.k8s.io/sig-release/release-team/role-handbooks/patch-release-manager) 
   will handle all cherry-picks to patch releases.

   The [Branch Manager](https://git.k8s.io/sig-release/release-team/role-handbooks/branch-manager)
   or the [Patch Release Team](https://git.k8s.io/sig-release/release-team/role-handbooks/patch-release-manager)
   are the final authority on branch approval by removing the `do-not-merge/cherry-pick-not-approved`
   label and triggering a merge into the target branch.

   The team scrubs through incoming cherry-picks on at least a weekly basis, daily during
   burndown ahead of a .0 release.  Ahead of point releases, reminders of the
   cherry-pick deadline will be sent out to the community.  Cherry-pick PRs are
   often metered into the release branches to give more deliberate CI signal across
   changes.  For this reason your cherry-pick must be ready to merge ahead of
   the cherry-pick deadline, but those candidates may be merged during the days
   between the deadline and release.

   Open cherry-pick PRs which do not land in the current release will
   continue to be tracked by the team for consideration for inclusion in a next
   patch release.

   If you are concerned about the status of your cherry-pick, err on the
   side of overcommunicating and reach out to the branch reviewer(s):

   * During code freeze or after code thaw and ahead of a .0 release, to get attention on a cherry-pick by the current
     release team members see the [appropriate release folder](https://git.k8s.io/sig-release/releases)
     for the target release's team contact information. You may cc them with
     `@<githubusername>` on your cherry-pick PR.

   * For prior branches, check the [patch release schedule](https://git.k8s.io/sig-release/releases/patch-releases.md), which includes contact information for the patch release team.

## Searching for Cherry-picks

- [A sample search on kubernetes/kubernetes pull requests that are labeled as `cherry-pick-approved`](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Acherry-pick-approved)

- [A sample search on kubernetes/kubernetes pull requests that are labeled as `do-not-merge/cherry-pick-not-approved`](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Ado-not-merge%2Fcherry-pick-not-approved)


## Troubleshooting Cherry-picks

Contributors may encounter some of the following difficulties when initiating a cherry-pick.

- A cherry-pick PR does not apply cleanly against an old release branch.
In that case, you will need to manually fix conflicts.

- The cherry-pick PR includes code that does not pass CI tests.
In such a case you will have to fetch the auto-generated branch from your fork, amend the problematic commit and force push to the auto-generated branch.
Alternatively, you can create a new PR, which is noisier.

## Cherry-picks for unsupported releases

The release team only supports & patches `n-3` releases (`n` being the latest release of Kubernetes). In January of 2019 the community discovered a regression, that was introduced in a post-release patch, but was currently no longer supported.

As discussed in a sig-release meeting on 2019-01-15, a fix was backported to the non supported version.

Reference PR: [#72860](https://github.com/kubernetes/kubernetes/pull/72860#issuecomment-454072746)

The specific criteria driving the decision was:

- CI was still available for the version
- The regression was introduced as a patch (and not part of the official release)
- The issue being fixed is of sufficient **[severity & impact](#what-kind-of-prs-are-good-for-cherry-picks)**
- The fix is well understood and contained (doesn’t introduce risk of additional regressions)

A note about the specific case in [#72860](https://github.com/kubernetes/kubernetes/pull/72860#issuecomment-454072746):

- The patch was exceedingly tiny and very unlikely to introduce new problems
- Luckily, it was caught shortly after the release was supposed to be unsupported