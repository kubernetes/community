# Overview

This document explains how cherry-picks are managed on release branches within
the kubernetes/kubernetes repository.
A common use case for this task is backporting PRs from master to release 
branches.

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
   The [Branch Manager](https://git.k8s.io/sig-release/release-team/role-handbooks/branch-manager)
   will triage PRs targeted to the next .0 minor release branch up until the 
   release, while the [Patch Release Team](https://git.k8s.io/sig-release/release-team/role-handbooks/patch-release-manager) 
   will handle all cherry-picks to patch releases.
   Normal rules apply for code merge.
   * Reviewers `/lgtm` and owners `/approve` as they deem appropriate.
   * Milestones on cherry-pick PRs should be the milestone for the target 
   release branch (for example, milestone 1.11 for a cherry-pick onto 
   release-1.11).
   *  You can find the current release team members in the 
   [appropriate release folder](https://git.k8s.io/sig-release/releases) for the target release.
   You may cc them with `<@githubusername>` on your cherry-pick PR.

## Cherry-pick Review

Cherry-pick pull requests have an additional requirement compared to normal pull
requests.
They must be approved specifically for cherry-pick by Approvers.
The [Branch Manager](https://git.k8s.io/sig-release/release-team/role-handbooks/branch-manager) 
or the [Patch Release Team](https://git.k8s.io/sig-release/release-team/role-handbooks/patch-release-manager)
are the final authority on removing the `do-not-merge/cherry-pick-not-approved`
label  and triggering a merge into the target branch.

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
