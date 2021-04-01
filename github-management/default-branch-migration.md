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

These changes are **non-disruptive**  and can be made anytime before renaming
the branch.

- [ ] If a presubmit or postsubmit prowjob triggers on the `master` branch
  (`branches` field of the prowjob), add the `main` branch to the list
  (see [kubernetes/test-infra#20665] for an example).

- [ ] If the [`milestone_applier`] prow config references the `master` branch,
add the `main` branch to the config (see [kubernetes/test-infra#20675] for an example).

- [ ] If the [`branch_protection`] prow config references the `master` branch,
add the `main` branch to the config.

### Just before rename

These changes are **disruptive** and should be made just before renaming the
branch.

- [ ] For periodic prowjobs, or any prowjob that mentions the `master` branch
in `base_ref`, update them to the `main` branch. Ensure that these changes
happen in lock-step with the branch rename (jobs triggered in between landing
these changes and renaming the branch will fail).
  - For bootstrap-based jobs, ensure the branch is explicitly specified,
    e.g. `kubernetes/foo=main`. [kubernetes/test-infra#20667] may eventually
    allow for non-disruptive changes.
  - For pod-utils based jobs, ensure the branch is explicitly specified,
    e.g. `base_ref: main`. [kubernetes/test-infra#20672] may eventually allow
    for non-disruptive changes.

- [ ] If a prowjob mentions `master` in its name, rename the job to not include
the branch name, e.g. `pull-repo-verify-master` -> `pull-repo-verify`.
[`status-reconciler`] should automatically migrate PR status contexts to the
new job name, and retrigger accordingly, but we have anecdotally found it
sometimes misses changes.
  - NOTE: our infrastructure doesn't understand the concept of job renames, so
  from the perspective of e.g. https://testgrid.k8s.io the job will appear to
  have lost history and start from scratch.

- [ ] If a prowjob calls scripts or code in your repo that explicitly
reference `master`, update all references to use `main`, or auto-detect the
remote branch
  - e.g. using git to auto-detect
  ```sh
  # for existing clones, update their view of the remote
  git fetch origin
  git remote set-head origin -a
  # for new clones, or those updated as above, this prints "main" post-rename
  echo $(git symbolic-ref refs/remotes/origin/HEAD)
  ```
  - e.g. using github's api to auto-detect
  ```sh
  # gh is https://github.com/cli/cli, this will print "main" post-rename
  gh api /repos/kubernetes-sigs/slack-infra | jq -r .default_branch
  ```

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
  - If there are any outstanding PRs you can /approve to merge, do so to verify
  that presubmits and postsubmits work as expected

- [ ] Trial the local development experience with a pre-rename clone.
  - ensure [Github instructions to rename your local branch] work
  - consider updating your fork's default remote branch name such that if you
  have git autocompletion enabled, typing `ma<tab>` will autocomplete to `main`

- [ ] Send a notice about the branch rename to your SIG's mailing list.
Include the link to the [GitHub instructions to rename your local branch].

[kubernetes/org]: https://github.com/kubernetes/org/issues
[@kubernetes/owners]: https://github.com/orgs/kubernetes/teams/owners
[#github-management]: https://kubernetes.slack.com/messages/github-management
[kubernetes/test-infra#20665]: https://github.com/kubernetes/test-infra/pull/20665
[kubernetes/test-infra#20667]: https://github.com/kubernetes/test-infra/issues/20667
[kubernetes/test-infra#20669]: https://github.com/kubernetes/test-infra/pull/20669
[kubernetes/test-infra#20672]: https://github.com/kubernetes/test-infra/issues/20672
[kubernetes/test-infra#20675]: https://github.com/kubernetes/test-infra/pull/20675
[`status-reconciler`]: https://github.com/kubernetes/test-infra/tree/master/prow/cmd/status-reconciler
[`branch_protection`]: https://github.com/kubernetes/test-infra/blob/ca6273046b355d38eade4c4bd435bd13fbb55043/config/prow/config.yaml#L131
[`milestone_applier`]: https://github.com/kubernetes/test-infra/blob/ca6273046b355d38eade4c4bd435bd13fbb55043/config/prow/plugins.yaml#L324
[official instructions]: https://github.com/github/renaming#renaming-existing-branches
[GitHub instructions to rename your local branch]: https://docs.github.com/en/github/administering-a-repository/renaming-a-branch#updating-a-local-clone-after-a-branch-name-changes
