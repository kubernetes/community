---
title: "OWNERS Files"
weight: 15
description: |
  OWNERS files are used to designate responsibility over different parts of the
  Kubernetes codebase and serve as the implementation mechanism for the two-phase
  code review process used by the project.
---

## Overview

OWNERS files are used to designate responsibility over different parts of the Kubernetes codebase.
Today, we use them to assign the **[reviewer][reviewer-role]** and **[approver][approver-role]** 
roles (defined in our [community membership doc]) that are used in our two-phase code review 
process. Our OWNERS files were inspired by [Chromium OWNERS files][chromium-owners] which in turn
inspired [GitHub's CODEOWNERS files][github-codeowners]

The velocity of a project that uses code review is limited by the number of people capable of
reviewing code. The quality of a person's code review is limited by their familiarity with the code
under review. Our goal is to address both of these concerns through the prudent use and maintenance
of OWNERS files.

## OWNERS spec

The [k8s.io/test-infra/prow/repoowners package](https://git.k8s.io/test-infra/prow/repoowners/repoowners.go)
is the main consumer of OWNERS files.  If this page is out of date, look there.

### OWNERS

Each directory that contains a unit of independent code or content may also contain an OWNERS file.
This file applies to everything within the directory, including the OWNERS file itself, sibling
files, and child directories.

OWNERS files are in YAML format and support the following keys:

- `approvers`: a list of GitHub usernames or aliases that can `/approve` a PR
- `labels`: a list of GitHub labels to automatically apply to a PR
- `options`: a map of options for how to interpret this OWNERS file, currently only one:
  - `no_parent_owners`: defaults to `false` if not present; if `true`, exclude parent OWNERS files.
    Allows the use case where `a/deep/nested/OWNERS` file prevents `a/OWNERS` file from having any
    effect on `a/deep/nested/bit/of/code`
- `reviewers`: a list of GitHub usernames or aliases that are good candidates to `/lgtm` a PR

The above keys constitute a *simple OWNERS configuration*.

All users are expected to be assignable. In GitHub terms, this means they must be
members of the organization to which the repo belongs.

A typical OWNERS file looks like:

```yaml
approvers:
  - alice
  - bob     # this is a comment
reviewers:
  - carol
  - david   # this is another comment
  - sig-foo # this is an alias
```

#### Filters

An OWNERS file may also include a `filters` key.
The `filters` key is a map whose keys are [Go regular expressions][go-regex] and whose values are [simple OWNERS configurations](#owners).
The regular expression keys are matched against paths relative to the OWNERS file in which the keys are declared.
For example:

```yaml
filters:
  ".*":
    labels:
    - re/all
  "\\.go$":
    labels:
    - re/go
```

If you set `filters` you must not set a [simple OWNERS configuration](#owners) outside of `filters`.
For example:

```yaml
# WARNING: This use of 'labels' and 'filters' as siblings is invalid.
labels:
- re/all
filters:
  "\\.go$":
    labels:
    - re/go
```

Instead, set a `.*` key inside `filters` (as shown in the previous example).

**WARNING**: The `approve` plugin [does not currently respect `filters`][test-infra-7690].
Until that is fixed, `filters` should only be used for the `labels` key (as shown in the above example).


#### Emeritus

It is inevitable, but there are times when someone may shift focuses, change jobs or step away from
a specific area in the project for a time. These people may be domain experts over certain areas
of the codebase, but can no longer dedicate the time needed to handle the responsibilities of
reviewing and approving changes. They are encouraged to add themselves as an _"emeritus"_ approver
under the `emeritus_approvers` key.

GitHub usernames listed under the `emeritus_approvers` key can no longer approve code (use the
`/approve` command) and will be ignored by prow for assignment. However, it can still be referenced
by a person looking at the OWNERS file for a possible second or more informed opinion.

When a contributor returns to being more active in that area, they may be promoted back to a
regular approver at the discretion of the current approvers.

```yaml
emeritus_approvers:
- david    # 2018-05-02
- emily    # 2019-01-05
```

#### Cleanup

In addition to the Emeritus process above, from time to time, it is necessary
to prune inactive members from OWNERS files. A core principle in maintaining a
healthy community is encouraging active participation. OWNERS have a higher
activity requirement, as they directly impact the ability of others to contribute.
If an owner should become inactive, they will be removed but may be added back
should they become active again.


An inactive OWNER is defined as someone with less than 10 Devstats recorded
contributions within the past year, as shown by this [dashboard]. This is a
conservative metric but should ensure only the removal of the most inactive
OWNERS.
- PR comments are less than 10 and Devstats count is less than 10 for a year



### OWNERS_ALIASES

Each repo may contain at its root an OWNERS_ALIASES file.

OWNERS_ALIASES files are in YAML format and support the following keys:

- `aliases`: a mapping of alias name to a list of GitHub usernames

We use aliases for groups instead of GitHub Teams, because changes to GitHub Teams are not
publicly auditable.

A sample OWNERS_ALIASES file looks like:

```yaml
aliases:
  sig-foo:
    - david
    - erin
  sig-bar:
    - bob
    - frank
```

GitHub usernames and aliases listed in OWNERS files are case-insensitive.

## Code Review using OWNERS files

[Code Review Process for the Kubernetes Project](review-process.md) is a simplified description
of our [full PR testing and merge workflow][pr-workflow] that conveniently forgets about the
existence of tests, to focus solely on the roles driven by OWNERS files.

Please see [below](#automation-using-owners-files) for details on how specific aspects of this
process may be configured on a per-repo basis.


## Automation using OWNERS files

Kubernetes uses the Prow Blunderbuss plugin and Tide.
Tide uses GitHub queries to select PRs into “tide pools”, runs as many in a
batch as it can (“tide comes in”), and merges them (“tide goes out”).

- [Blunderbuss plugin](https://git.k8s.io/test-infra/prow/plugins/blunderbuss):
  - responsible for determining **reviewers**
- [Tide](https://git.k8s.io/test-infra/prow/cmd/tide):
  - responsible for automatically running batch tests and merging multiple PRs together whenever possible.
  - responsible for retriggering stale PR tests.
  - responsible for updating a GitHub status check explaining why a PR can't be merged (eg: a
    missing `lgtm` or `approved` label)

### [`prow`](https://git.k8s.io/test-infra/prow)

Prow receives events from GitHub, and reacts to them. It is effectively stateless. The following
pieces of prow are used to implement the code review process above.

- [cmd: tide](https://git.k8s.io/test-infra/prow/cmd/tide)
  - per-repo configuration:
    - `labels`: list of labels required to be present for merge (eg: `lgtm`)
    - `missingLabels`: list of labels required to be missing for merge (eg: `do-not-merge/hold`)
    - `reviewApprovedRequired`: defaults to `false`; when true, require that there must be at least
      one [approved pull request review](https://help.github.com/articles/about-pull-request-reviews/)
      present for merge
    - `merge_method`: defaults to `merge`; when `squash` or `rebase`, use that merge method instead
      when clicking a PR's merge button
  - merges PR's once they meet the appropriate criteria as configured above
  - if there are any presubmit prow jobs for the repo the PR is against, they will be re-run one
    final time just prior to merge
- [plugin: assign](https://git.k8s.io/test-infra/prow/plugins/assign)
  - assigns GitHub users in response to `/assign` comments on a PR
  - unassigns GitHub users in response to `/unassign` comments on a PR
- [plugin: approve](https://git.k8s.io/test-infra/prow/plugins/approve)
  - per-repo configuration:
    - `issue_required`: defaults to `false`; when `true`, require that the PR description link to
      an issue, or that at least one **approver** issues a `/approve no-issue`
    - `implicit_self_approve`: defaults to `false`; when `true`, if the PR author is in relevant
      OWNERS files, act as if they have implicitly `/approve`'d
  - adds the  `approved` label once an **approver** for each of the required
    OWNERS files has `/approve`'d
  - comments as required OWNERS files are satisfied
  - removes outdated approval status comments
- [plugin: blunderbuss](https://git.k8s.io/test-infra/prow/plugins/blunderbuss)
  - determines **reviewers** and requests their reviews on PR's
- [plugin: lgtm](https://git.k8s.io/test-infra/prow/plugins/lgtm)
  - adds the `lgtm` label when a **reviewer** comments `/lgtm` on a PR
  - the **PR author** may not `/lgtm` their own PR
- [pkg: k8s.io/test-infra/prow/repoowners](https://git.k8s.io/test-infra/prow/repoowners/repoowners.go)
  - parses OWNERS and OWNERS_ALIAS files
  - if the `no_parent_owners` option is encountered, parent owners are excluded from having
    any influence over files adjacent to or underneath of the current OWNERS file

## Maintaining OWNERS files

OWNERS files should be regularly maintained.

We encourage people to self-nominate, self-remove or switch to [emeritus](#emeritus) from OWNERS
files via PR's. Ideally in the future we could use metrics-driven automation to assist in this
process.

We should strive to:

- grow the number of OWNERS files
- add new people to OWNERS files
- ensure OWNERS files only contain organization members
- ensure OWNERS files only contain people are actively contributing to or reviewing the code they own
- remove inactive people from OWNERS files

Bad examples of OWNERS usage:

- directories that lack OWNERS files, resulting in too many hitting root OWNERS
- OWNERS files that have a single person as both approver and reviewer
- OWNERS files that haven't been touched in over 6 months
- OWNERS files that have non organization members present

Good examples of OWNERS usage:

- team aliases are used that correspond to sigs
- there are more `reviewers` than `approvers`
- the `approvers` are not in the `reviewers` section
- OWNERS files that are regularly updated (at least once per release)

[go-regex]: https://golang.org/pkg/regexp/#pkg-overview
[test-infra-7690]: https://github.com/kubernetes/test-infra/issues/7690
[approver-role]: https://git.k8s.io/community/community-membership.md#approver
[reviewer-role]: https://git.k8s.io/community/community-membership.md#reviewer
[community membership doc]: https://git.k8s.io/community/community-membership.md
[chromium-owners]: https://chromium.googlesource.com/chromium/src/+/master/docs/code_reviews.md
[github-codeowners]: https://help.github.com/articles/about-codeowners/
[pr-workflow]: /contributors/guide/pull-requests.md#the-testing-and-merge-workflow
[dashboard]: https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-repo_name=kubernetes%2Fkubernetes&var-country_name=All
