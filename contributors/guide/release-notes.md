---
title: "Adding Release Notes"
weight: 12
description: |
  Guidance on providing release notes for changes made to the main Kubernetes
  project repo.
---

On the [kubernetes/kubernetes repository][kubernetes-repository], release notes
are required for any pull request with user-visible changes, this could mean:

- User facing, critical bug-fixes
- Notable feature additions
- Output format changes
- Deprecations or removals
- Metrics changes
- Dependency changes
- API changes

Release notes are one of the most important reference points for users about to
install or upgrade to a particular release of Kubernetes.

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

No release notes are required for changes to the following:

- Tests
- Build infrastructure
- Fixes for unreleased bugs

## Contents of a Release Note

A release note needs a clear, concise description of the change. This includes:

1. An indicator if the pull request _Added_, _Changed_, _Fixed_, _Removed_,
   _Deprecated_ functionality or changed enhancement/feature maturity (alpha,
   beta, stable/GA)
2. An indicator if there is user _Action required_
3. The name of the affected API or configuration fields, CLI commands/flags or
   enhancement/feature
4. A link to relevant user documentation about the enhancement/feature

## Writing Good Release Notes

Good release notes can make a huge difference to the end user. To make sure that
your release notes are of the highest possible quality, follow these guidelines.

- Write your release notes in the past tense. For example, use "Added" instead
  of "Add", "Fixed" instead of "Fix", and "Updated" instead of "Update".
- Include a call to action if there is anything the end user needs to do, along
  with a link to a blog post or documentation item to provide better context.
- Keep your release notes relevant to the end user.
- Use Markdown to add links and emphasis, to make your release notes more
  readable and usable.
- Use good grammar and make sure you capitalize the first word of each
  sentence. 

Here are some pull requests with examples of exemplary release notes:
- https://github.com/kubernetes/kubernetes/pull/97252
- https://github.com/kubernetes/kubernetes/pull/105517

For more tips on writing good release notes, check out the [Release Notes Handbook].

## Applying a Release Note

On the `kubernetes/kubernetes` repository, release notes are required for any pull
request with user-visible changes.

To meet this requirement, do one of the following:
- Add notes in the release notes block, or
- Update the release note label

If you don't add release notes in the pull request template, the
`do-not-merge/release-note-label-needed` label is added to your pull request
automatically after you create it. There are a few ways to update it.

To add a release-note section to the pull request description, add your release
note beneath the question *Does this PR introduce a user-facing change?*

For pull requests that require additional action from users switching to the new
release, include the string "action required" (case insensitive) in the release
note.

For an example, see [this pull request](https://github.com/kubernetes/kubernetes/pull/107207).

For pull requests that don't need to be mentioned at release time, use the
`/release-note-none` Prow command to add the `release-note-none` label to the
PR. You can also write the string "NONE" as a release note in your PR
description.

For an example of a pull request without release notes, 
[take a look at this pull request](https://github.com/kubernetes/kubernetes/pull/107910).

Your release note should be written in clear and straightforward sentences. Most
often, users aren't familiar with the technical details of your PR, so consider
what they _need to know_ when you write your release note.

Some brief examples of release notes:

```
The deprecated flag --conntrack-max has been removed from kube-proxy. Users of this flag should switch to --conntrack-min and --conntrack-max-per-core instead. (#78399, @rikatz)

The --export flag for the kubectl get command, deprecated since v1.14, will be removed in v1.18.

Fixed a bug that prevents dry-run from being honored for the pod/eviction sub-resource. (#76969, @apelisse)
```

Pull Request titles and body comments can be modified at any time prior to the
release to make them friendly for release notes.

The release notes team maintains a
[template](https://github.com/kubernetes/sig-release/blob/master/release-team/role-handbooks/release-notes/relnotes-template.md)
for Kubernetes Release notes that may help clarify whether or not your PR
requires a release note. The most recent 
[Kubernetes Release notes](https://kubernetes.io/docs/setup/release/notes/) can
also provide insight into the writing style for release notes.

Release notes apply to pull requests on the master branch. For patch release
branches the automated cherry-pick pull requests process (see the 
[cherry-pick instructions](/contributors/devel/sig-release/cherry-picks.md))
should be followed.  That automation will pull release notes from the master
branch PR from which the cherry-pick originated. On a rare occasion a pull
request on a patch release branch is not a cherry-pick, but rather is targeted
directly to the non-master branch and in this case, a `release-note-*` label is
required for that non-master pull request.

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

[kubernetes-repository]: https://git.k8s.io/kubernetes/
[Release Notes Handbook]: https://github.com/kubernetes/sig-release/blob/master/release-engineering/release-notes.md
