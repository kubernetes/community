---
title: "Adding Release Notes"
weight: 12
slug: "release-notes"
---

On the kubernetes/kubernetes repository, release notes are required for any pull
request with user-visible changes, such as bug fixes, feature additions, and
output format changes. Release notes are one of the most important reference
points for users about to install or upgrade to a particular release of
Kubernetes.

## Does my pull request need a release note?

Any user-visible or operator-visible change qualifies for a release note. This
could be a:

- CLI change
- API change
- UI change
- configuration schema change
- behavioral change
- change in non-functional attributes such as efficiency or availability,
  availability of a new platform
- a warning about a deprecation
- fix of a previous _Known Issue_
- fix of a vulnerability (CVE)

No release notes are required for changes to:

- tests
- build infrastructure
- fixes of bugs which have not been released

## Contents of a Release Note

A release note needs a clear, concise description of the change. This includes:

1. an indicator if the pull request _Added_, _Changed_, _Fixed_, _Removed_,
   _Deprecated_ functionality or changed enhancement/feature maturity (alpha,
   beta, stable/GA)
2. an indicator if there is user _Action required_
3. the name of the affected API or configuration fields, CLI commands/flags or
   enhancement/feature
4. a link to relevant user documentation about the enhancement/feature

## Applying a Release Note

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

## Reviewing Release Notes

Reviewing the release notes of a pull request should be a dedicated step in the
overall review process. It is necessary to rely on the same metrics as other
reviewers to be able to distinguish release notes which might need to be
rephrased.

As a guideline, a release notes entry needs to be rephrased if one of the
following cases apply:

- The release note does not communicate the full purpose of the change.
- The release note has no impact on any user.
- The release note is grammatically incorrect.

In any other case the release note should be fine.

## Related

* [Behind The Scenes: Kubernetes Release Notes Tips & Tricks - Mike Arpaia, Kolide (KubeCon 2018 Lightning Talk)](https://www.youtube.com/watch?v=n62oPohOyYs)
