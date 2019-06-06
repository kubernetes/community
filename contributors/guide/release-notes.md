---
title: "Adding Release Notes"
weight: 1
slug: "adding-release-notes" 
---

On the kubernetes/kubernetes repository, release notes are required for any pull request with user-visible changes, such as bug-fixes, feature additions, and output format changes.

To meet this requirement, do one of the following:
- Add notes in the release notes block, or
- Update the release note label

If you don't add release notes in the pull request template, the `do-not-merge/release-note-label-needed` label is added to your pull request automatically after you create it. There are a few ways to update it.

To add a release-note section to the pull request description:

For pull requests with a release note:

    ```release-note
    Your release note here
    ```

For pull requests that require additional action from users switching to the new release, include the string "action required" (case insensitive) in the release note:

    ```release-note
    action required: your release note here
    ```

For pull requests that don't need to be mentioned at release time, use the `/release-note-none` Prow command to add the `release-note-none` label to the PR. You can also write the string "NONE" as a release note in your PR description:

    ```release-note
    NONE
    ```

To see how to format your release notes, view the kubernetes/kubernetes [pull request template](https://git.k8s.io/kubernetes/.github/PULL_REQUEST_TEMPLATE.md) for a brief example. Pull Request titles and body comments can be modified at any time prior to the release to make them friendly for release notes.

Release notes apply to pull requests on the master branch. For patch release branches the automated cherry-pick pull requests process (see the [cherry-pick instructions](/contributors/devel/sig-release/cherry-picks.md)) should be followed.  That automation will pull release notes from the master branch PR from which the cherry-pick originated. On a rare occasion a pull request on a patch release branch is not a cherry-pick, but rather is targeted directly to the non-master branch and in this case, a `release-note-*` label is required for that non-master pull request.
