# Overview

This document explains how cherry picks are managed on release
branches within the Kubernetes projects.

## Prerequisites
 * [Contributor License Agreement](http://git.k8s.io/community/CLA.md)
   is considered implicit for all code within cherry-pick pull requests,
   ***unless there is a large conflict***.
 * A pull request merged against the master branch.
 * [Release branch](https://git.k8s.io/release/docs/branching.md) exists.
 * The normal git and GitHub configured shell environment for pushing
   to your kubernetes `origin` fork on GitHub and making a pull request
   against a configured remote `upstream` that tracks
   "https://github.com/kubernetes/kubernetes.git", including
   `GITHUB_USER`.
 * Have `hub` installed, which is most easily installed via `go get
   github.com/github/hub` assuming you have a standard golang development
   environment.

## Initiate a Cherry Pick
 * Run the [cherry pick script](https://git.k8s.io/kubernetes/hack/cherry_pick_pull.sh).
   This example applies a master branch PR #98765 to the remote branch
   `upstream/release-3.14`: `hack/cherry_pick_pull.sh upstream/release-3.14
   98765`
 * Your cherrypick PR targeted to the release branch will immediately get the
   `do-not-merge/cherry-pick-not-approved` label. The release branch owner
   will triage PRs targeted to the branch.  Normal rules apply for code merge.
   * Reviewers `/lgtm` and owners `/approve` as they deem appropriate.
   * The approving release branch owner is responsible for applying the
     `cherrypick-approved` label.

## Cherry Pick Review

Cherry pick pull requests are reviewed differently than normal pull requests. In
particular, they may be self-merged by the release branch owner without fanfare,
in the case the release branch owner knows the cherry pick was already
requested - this should not be the norm, but it may happen.

## Searching for Cherry Picks

See the [cherrypick queue dashboard](http://cherrypick.k8s.io/#/queue) for
status of PRs labeled as `cherrypick-candidate`.
