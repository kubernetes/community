---
title: Default Branch Migration
linkTitle: Default Branch Migration
description:  |
  Instructions on how to migrate the default branch from master to main.
weight: 99
type: docs
aliases: [ "/rename" ]
---

# Default Branch Migration

This document outlines steps needed to migrate the default branch
of your repo from `master` to `main`.

Note: This document is currently a work in progress.

If you have questions about the process, reach out to the GitHub Management Team
on the [#github-management] channel on slack or open an issue in the [kubernetes/org] repo.

## Prerequisites

- [ ] Ensure that your repo has low PR volume (<20 open PRs) and
less number of periodic jobs. The branch rename will re-trigger
prow on _all_ open PRs, which will cause a huge spike in the CI load.

- [ ] Create an issue in your repo to track the branch rename.
You can paste this checklist in the issue body.

- [ ] If you are not a root approver for the repo, assign a root
approver for approval.

- [ ] Once the issue has been approved, send a notice to your SIG's
mailing list about the potential branch rename.

## Changes pre-rename

Make the following changes _before_ renaming the branch the `master` branch.

Note: There might be additional changes required that have not been
covered in this checklist.

### Anytime

These changes are **non-disruptive**  and can be made anytime before
renaming the branch.

- [ ] If a prowjob triggers on the `master` branch (`branches` field
of the prowjob), add the `main` branch to the list
(see [kubernetes/test-infra#20665] for an example).

- [ ] If the [`milestone_applier`] prow config references the `master` branch,
add the `main` branch to the config (see [kubernetes/test-infra#20675] for an example).

- [ ] If the [`branch_protection`] prow config references the `master` branch,
add the `main` branch to the config.

### Just before rename

These changes are **disruptive** and should be made just before
renaming the branch.

- [ ] If a prowjob mentions the `master` branch in `base_ref`,
update it to the `main` branch. For a periodic job, ensure that
the branch is renamed between periodic job runs.

- [ ] If a prowjob mentions `master` in its name, rename the job to
to not include the branch name. [`status-reconciler`] should automatically
migrate the PR status contexts to the new job name but this has not been tested yet.
The job with the new name will also appear as a differt job in Testgrid.

- [ ] If a prowjob calls scripts or code in your repo that explicitly
reference `master`, update all references to use `main`.

- [ ] If the repo has netlify configured for it, ask a member of the GitHub
Management Team to rename the `master` branch to `main` in the netlify site config.
It can't be controlled through the netlify config in the repo.

### Approval

- [ ] Once all non-disruptive tasks have been completed and disruptive tasks
have been identified, assign the GitHub Management team ([@kubernetes/owners])
for approval.

## Rename the default branch

- [ ] Rename the default branch from `master` to `main` using the GitHub UI
by following the [official instructions].

## Changes post-rename

After the default branch has been renamed to `main`, make the following
changes.

Note: There might be additional changes required that have not been
covered in this checklist.

### Prowjobs

- [ ] If a prowjob still references the `master` branch in the `branches` field,
remove the `master` branch (see [kubernetes/test-infra#20669] for an example).

### Prow config

- [ ] If the [`milestone_applier`] prow config references the `master` branch,
remove it from the config.

- [ ] If the [`branch_protection`] prow config references the `master` branch,
remove it from the config.

### Other

- [ ] If any docs reference the `master` branch, update to `main`
(URLs will be automatically redirected).

- [ ] Ensure that CI and PR tests work fine.

- [ ] Trial the local development experience with a pre-rename clone.

- [ ] Send a notice about the branch rename to your SIG's mailing list.
Include the link to the [GitHub instructions to rename your local branch].

[kubernetes/org]: https://github.com/kubernetes/org/issues
[@kubernetes/owners]: https://github.com/orgs/kubernetes/teams/owners
[#github-management]: https://kubernetes.slack.com/messages/github-management
[kubernetes/test-infra#20665]: https://github.com/kubernetes/test-infra/pull/20665
[kubernetes/test-infra#20667]: https://github.com/kubernetes/test-infra/issues/20667
[kubernetes/test-infra#20669]: https://github.com/kubernetes/test-infra/pull/20669
[kubernetes/test-infra#20675]: https://github.com/kubernetes/test-infra/pull/20675
[`status-reconciler`]: https://github.com/kubernetes/test-infra/tree/master/prow/cmd/status-reconciler
[`branch_protection`]: https://github.com/kubernetes/test-infra/blob/ca6273046b355d38eade4c4bd435bd13fbb55043/config/prow/config.yaml#L131
[`milestone_applier`]: https://github.com/kubernetes/test-infra/blob/ca6273046b355d38eade4c4bd435bd13fbb55043/config/prow/plugins.yaml#L324
[official instructions]: https://github.com/github/renaming#renaming-existing-branches
[GitHub instructions to rename your local branch]: https://docs.github.com/en/github/administering-a-repository/renaming-a-branch#updating-a-local-clone-after-a-branch-name-changes
