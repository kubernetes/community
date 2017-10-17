# Creating PRs for release notes

This document describes how to create PRs in Kubernetes repos so that the changes can be tracked with the
Kubernetes release and published in release notes.

## Curated release notes

The release notes for each Kubernetes release are divided into several sections.

All sections are grouped into subsections by component.

#### Action required

The action required section outlines any steps the user must perform before upgrading their cluster.
The may include actions such as migrating APIs or commands to new versions, upgrading /
installing tools, or validating the cluster environment.

#### Deprecations

The deprecation section outlines any APIs, tools, commands or flags that are being deprecated.
They will continue to be available for this release, but will be removed in a subsequent release.

#### Feature highlights

The feature highlights section is manually populated from the [features repo] and contains
a list of the feature highlights for the release. 

#### Notable changes

The notable changes section is generated from the PR release notes, and contains a list
of features and bug fixes going into the release.

## Release note labels

All  PRs with user visible changes must have a release note attached.  This will automatically

populate them in the release notes.

Release notes are added applying on of the `release-note*` labels by adding a section
to the PR description containing the test that will appear in the release notes:

>\```release-note*
>
> User visible change description
>
> \```

The change description will be directly added to the release notes, and should include the visible
impact to the end user.  Individual SIGs should consider developing guidelines around how to write
release notes for the components they own.  The release-note labels follow:

- `release-note-action-required`

The `release-note-action-required` label will cause the PR to be added to the Action Required section
of the release notes.  This label should be added for any PRs that result in the user
needing to take action before, during or after upgrading their cluster to the new release.  The release
note should outline what action the user must perform.  e.g.

> kubectl --foo has been removed and replaced with --bar.  All references to --foo must be updated.

- `release-note-deprecation`

The `release-note-deprecation` label will cause the PR to be added to the Deprecation section
of the release notes.  This label should be added for any PRs that mark APIs, tools, commands or
flags as deprecated.

- `release-note`

The `release-note` label will cause the PR to be added to the Notable changes section.

- `release-note-none`

The `release-note-none` should be applied to any changes that are not visible to the end user. e.g.
refactorings, test changes, and disabled code.

## Additional required labels

- `sig/*` e.g. `sig/cli`

The `sig` label will be used by SIGs to process and edit release notes before the release is published.
Release notes often must be ordered and reworded so they are consistent.  The SIG labels will not
appear in the final release notes.

- `component/*`

The component label will be used to further group the release notes by component.  This is useful for
SIGs with multiple components - e.g. SIG node with kubelet, kubeproxy, cri

- `kind/bug`, `kind/feature`

The `kind` label will be used to sort the release notes as either bug fixes or new features.

[kubernetes repo]: (https://github.com/kubernetes/kubernetes)
[features repo]: (https://github.com/kubernetes/features)
