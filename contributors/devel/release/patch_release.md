# Building Kubernetes patch releases

This document describes the process for creating Kubernetes patch releases and communicating publishing
status regarding known issues impacting the current release.

## Process

After a Kubernetes minor release, patch releases are managed by the Kubenretes patch release manager.
The patch release manager works with SIGs to coordinate merging cherrypicks into the release branch,
then cutting and publishing a new patch release.

## Communicating fixes to the community

The canonical source of issues in a patch release will be communicated through GitHub issues.

To make it easy for the community to understand the status of each patch release, GitHub issues
must follow a consistent format and contain basic information about the issue.

GitHub issues targeted at patch releases must contain the following items:

### Priority label

The priority label describes the urgency and severity of the issue.

- `priority/critical-urgent`: the issue is severe enough that it requires an immediate patch release
  (storage or network issues that can cause data corruption or outages for instance)
- `priority/important-soon`: the issue should be fixed in a patch does not need to be fixed immediately

### Sig label

SIG labels should be applied for all SIGs involved in the issue / resolution.

- `sig/cli`
- `sig/node`

### Summary template

The issue description should be kept up-to-date by the issue owner through out the resolution process.  Items
discovered as part of triage should be reflected in the issue description.

While much of this information may exist as discussion comments on the issue,
providing the information in an easily understandable format and location
makes it much easier to quickly understand the state of upcoming patch release.

The template will start with a section describing which releases the issue was introduced and resolved in.  The
format will be machine parsable so that bots can apply labels and generate reports using the information in
the issue.

<p> &#96;&#96;&#96;release </p>
<p>introduced-in=vX.Y.Z </p>
<p>&#96;&#96;&#96;</p>

<p> &#96;&#96;&#96;release </p>
<p>resolved-in=vX.Y.Z </p>
<p>resolved-in=vX.Y+1.Z </p>
<p>&#96;&#96;&#96;</p>

The rest of the template is as follows:

```sh
## Symptoms

What users are experiencing

## Root cause

The technical cause of the symptoms including the list of components / binaries.

e.g.

Binaries:
- kubectl

Kubectl was incorrectly calculating the patch for apply.  When diffing foo...

## Impact

Why this is important enough to warrant a patch vs waiting until the next minor release

## Resolution

How the issue will be (was) fixed

## PRs

- #23456
```

## Communicating security fixes to the community

Due to the sensitive nature of security fixes, their details maybe omitted from the GitHub issue and
simply state that the owner is working with the [product security team](https://github.com/kubernetes/community/blob/master/contributors/devel/security-release-process.md) on a resolution.
