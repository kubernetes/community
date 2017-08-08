# OWNERS files

## Overview

OWNERS files are used to designate responsibility over different parts of the Kubernetes codebase.
Today, we use them to assign the **reviewer** and **approver** roles used in our two-phase code
review process. Our OWNERS files were inspired by [Chromium OWNERS
files](https://chromium.googlesource.com/chromium/src/+/master/docs/code_reviews.md), which in turn
inspired [GitHub's CODEOWNERS files](https://help.github.com/articles/about-codeowners/).

The velocity of a project that uses code review is limited by the number of people capable of
reviewing code. The quality of a person's code review is limited by their familiarity with the code
under review. Our goal is to address both of these concerns through the prudent use and maintenance
of OWNERS files

## OWNERS spec

The [mungegithub gitrepos
feature](https://github.com/kubernetes/test-infra/blob/master/mungegithub/features/repo-updates.go)
is the main consumer of OWNERS files.  If this page is out of date, look there.

Each directory that contains a unit of independent code or content may also contain an OWNERS file.
This file applies to everything within the directory, including the OWNERS file itself, sibling
files, and child directories.

OWNERS files are in YAML format and support the following keys:

- `approvers`: a list of GitHub usernames or aliases that can `/approve` a PR
- [DEPRECATED] `assignees`: do not use, equivalent to `approvers`
  ([kubernetes/test-infra#3851](https://github.com/kubernetes/test-infra/issues/3851))
- `labels`: a list of GitHub labels to automatically apply to a PR
- `reviewers`: a list of GitHub usernames or aliases that are good candidates to `/lgtm` a PR

All users are expected to be assignable. In GitHub terms, this means they are either collaborators
of the repo, or members of the organization to which the repo belongs.

A typical OWNERS file looks like:

```
approvers:
  - alice
  - bob     # this is a comment
reviewers:
  - alice
  - carol   # this is another comment
  - sig-foo # this is an alias
```

Each repo may contain at its root an OWNERS_ALIAS file.

OWNERS_ALIAS files are in YAML format and support the following keys:

- `aliases`: a mapping of alias name to a list of GitHub usernames

We use aliases for groups instead of GitHub Teams, because changes to GitHub Teams are not
publicly auditable.

A sample OWNERS_ALISES file looks like:

```
aliases:
  sig-foo:
    - david
    - erin
  sig-bar:
    - bob
    - frank
```

GitHub usernames and aliases listed in OWNERS files are case-insensitive.

## Code Review Process

This is a simplified description of our [full PR testing and merge
workflow](https://github.com/kubernetes/community/blob/master/contributors/devel/pull-requests.md#the-testing-and-merge-workflow)
that conveniently forgets about the existence of tests, to focus solely on the roles driven by
OWNERS files.

- The **author** submits a PR
- Phase 0: Automation determines **reviewers** and **approvers** for the PR
  - Determine the set of OWNERS files nearest to the code being changed
  - Choose two **reviewers**, and assign them to the PR (choose more than one to avoid assigning to
    an inactive / unavailable reviewer)
  - Choose suggested **approvers**, one from each OWNERS file, and list them in a comment on the PR
- Phase 1: Humans review the PR
  - **Reviewers** look for general code quality, correctness, sane software engineering, style, etc.
  - The PR **author** cannot be their own **reviewer**
  - If the code changes look good to them, a **reviewer** types `/lgtm` in a PR comment or review;
    if they change their mind, they `/lgtm cancel`
  - Once a **reviewer** has `/lgtm`'ed, `prow` ([@k8s-ci-robot](https://github.com/k8s-ci-robot/))
    applies an `lgtm` label to the PR
- Phase 2: Humans approve the PR
  - The PR **author** `/assign`'s all suggested **approvers** to the PR, and optionally notifies
    them (eg: "pinging @foo for approval")
  - The PR **author** can be their own **approver** (assuming they are listed in the appropriate
    OWNERS files)
  - **Approvers** look for holistic acceptance criteria, including dependencies with other features,
    forwards/backwards compatibility, API and flag definitions, etc
  - If the code changes look good to them, an **approver** types `/approve` in a PR comment or
    review; if they change their mind, they `/approve cancel`
  - `mungegithub` ([@k8s-merge-robot](https://github.com/k8s-merge-robot/)) updates its comment in
    the PR to indicate which **approvers** still need to approve
  - Once all **approvers** (one from each of the previously identified OWNERS files) have approved,
    `mungegithub` ([@k8s-merge-robot](https://github.com/k8s-merge-robot/)) applies an `approved`
label
- Phase 3: Automation merges the PR
  - All required labels are present (eg: `lgtm`, `approved`)
  - All required status checks for the PR are verified as green
  - The PR is merged

## Quirks of the Process

There are a number of behaviors we've observed that while _possible_ are discouraged, as they go
against the intent of this review process.  Some of these could be prevented in the future, but this
is the state of today.

- An **approver**'s `/lgtm` is simultaneously interpreted as an `/approve`
  - While a convenient shortcut for some, it can be surprising that the same command is interpreted
    in one of two ways depending on who the commenter is
  - Instead, explicitly write out `/lgtm` and `/approve` to help observers, or save the `/lgtm` for
    a **reviewer**
  - This goes against the idea of having at least two sets of eyes on a PR, and may be a sign that
    there are too few **reviewers** (who aren't also **approves)
- An **approver** can `/approve no-issue` to bypass the requirement that PR's must have linked
  issues
  - There is disagreement within the community over whether requiring every PR to have a linked
    issue provides value
  - Protest is being expressed in the form of overuse of `/approve no-issue`
  - Instead, suggest to the PR **author** that they edit the PR description to include a linked
    issue
  - This is a sign that we need to actually deliver value with linked issues, or be able to define
    what a "trivial" PR is in a machine-enforceable way to be able to automatically waive the linked
    issue requirement
- Technically, anyone who is a member of the kubernetes GitHub organization can drive-by `/lgtm` a
  PR
  - Drive-by reviews from non-members are encouraged as a way of demonstrating experience and
    intent to become a collaborator or reviewer
  - Drive-by `/lgtm`'s from members may be a sign that our OWNERS files are too small, or that the
    existing **reviewers** are too unresponsive
  - This goes against the idea of specifying **reviewers** in the first place, to ensure that
    **author** is getting actionable feedback from people knowledgeable with the code
- **Reviewers**, and **approvers** are unresponsive
  - This causes a lot of frustration for **authors** who often have little visibility into why their
    PR is being ignored
  - Many **reviewers** and **approvers** are so overloaded by GitHub notifications that @mention'ing
    is unlikely to get a quick response
  - If an **author** `/assign`'s a PR, **reviewers** and **approvers** will be made aware of it on
    their [PR dashboard](https://k8s-gubernator.appspot.com/pr)
  - An **author** can work around this by manually reading the relevant OWNERS files,
    `/unassign`'ing unresponsive individuals, and `/assign`'ing others
  - This is a sign that our OWNERS files are stale; pruning the **reviewers** and **approvers** lists
    would help with this
- **Authors** are unresponsive
  - This costs a tremendous amount of attention as context for an individual PR is lost over time
  - This hurts the project in general as its general noise level increases over time
  - Instead, close PR's that are untouched after too long (we currently have a bot do this after 90
    days)

## Implementation

### [`mungegithub`](https://github.com/kubernetes/test-infra/tree/master/mungegithub)

Mungegithub polls GitHub, and "munges" things it finds, including issues and pull requests. It is
stateful, in that restarting it means it loses track of which things it has munged at what time.

- [feature:
  gitrepos](https://github.com/kubernetes/test-infra/blob/master/mungegithub/features/repo-updates.go)
  - responsible for parsing OWNERS and OWNERS_ALIAS files
  - if its `use-reviewers` flag is set to false, **approvers** will also be **reviewers**
  - if its `enable-md-yaml` flag is set, `.md` files will also be parsed to see if they have
    embedded OWNERS content (this is only used by
[kubernetes.github.io](https://github.com/kubernetes/kubernetes.github.io/))
  - used by other mungers to get the set of **reviewers** or **approvers** for a given path
- [munger:
  blunderbuss](https://github.com/kubernetes/test-infra/blob/master/mungegithub/mungers/blunderbuss.go)
  - responsible for determining **reviewers** and assigning to them
  - chooses from people in the deepest/closest OWNERS files to the code being changed
  - weights its choice based on the magnitude of lines changed for each file
  - randomly chooses to ensure the same people aren't chosen every time
  - if its `blunderbuss-number-assignees` flag is unset, it will default to 2 assignees
- [munger:
  approval-handler](https://github.com/kubernetes/test-infra/blob/master/mungegithub/mungers/approval-handler.go)
  - responsible for adding the  `approved` label once an **approver** for each of the required
    OWNERS files has `/approve`'d
  - responsible for commenting as required OWNERS files are satisfied
  - responsible for removing outdated approval status comments
  - [full description of the
    algorithm](https://github.com/kubernetes/test-infra/blob/6f5df70c29528db89d07106a8156411068518cbc/mungegithub/mungers/approval-handler.go#L99-L111)
- [munger:
  submit-queue](https://github.com/kubernetes/test-infra/blob/master/mungegithub/mungers/submit-queue.go)
  - responsible for merging PR's
  - responsible for updating a GitHub status check explaining why a PR can't be merged (eg: a
    missing `lgtm` or `approved` label)

### [`prow`](https://github.com/kubernetes/test-infra/tree/master/prow)

Prow receives events from GitHub, and reacts to them. It is effectively stateless.

- [plugin: lgtm](https://github.com/kubernetes/test-infra/tree/master/prow/plugins/lgtm)
  - responsible for adding the `lgtm` label when a **reviewer** comments `/lgtm` on a PR
  - the **PR author** may not `/lgtm` their own PR
- [plugin: assign](https://github.com/kubernetes/test-infra/tree/master/prow/plugins/assign)
  - responsible for assigning GitHub users in response to `/assign` comments on a PR
  - responsible for unassigning GitHub users in response to `/unassign` comments on a PR

### GitHub

GitHub provides a few integration points for our automation:

- [Status Checks](https://help.github.com/articles/about-required-status-checks/): our tests and
  submit queue use these to communicate whether a commit is OK
- [Protected Branches](https://help.github.com/articles/about-protected-branches/): ensure that a
  branch cannot be merged unless all status checks are green
- [Merge Button](https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-button): the
  submit queue effectively uses an API driven click of this button


## Updating OWNERS

OWNERS files should be regularly maintained.

We should strive to:

- grow the number of OWNERS files
- add new people to OWNERS files
- ensure OWNERS files only contain org members and repo collaborators
- ensure OWNERS files only contain people are actively contributing to or reviewing the code they own
- remove inactive people from OWNERS files

Bad examples of OWNERS usage:

- directories that lack OWNERS files, resulting in too many hitting root OWNERS
- OWNERS files that have a single person as both approver and reviewer
- OWNERS files that haven't been touched in over 6 months
- OWNERS files that have non-collaborators present

Good examples of OWNERS usage:

- team aliases are used that correspond to sigs
- there are more `reviewers` than `approvers`
- the `approvers` are not in the `reviewers` section
- OWNERS files that are regularly updated (at least once per release)
