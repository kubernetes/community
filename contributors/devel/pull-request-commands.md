# Overview

There are several steps to authoring or reviewing a pull request in Kubernetes.
Due to limitations on how GitHub permissions can be expressed, the Kubernetes
authors developed a workflow using comments to run tests and set labels
used for merging PRs.  Merging a PR requires the following steps to be
completed before the PR will automatically be merged:

- Signing a CLA
- Setting release notes
- Having all e2e tests pass
- Getting an LGTM from a reviewer

# Master Branch Workflow

## Sign the CLA

If you have not already, the `k8s-ci-robot` will leave a comment with
instructions on how to sign the CLA.

**Important** the email you sign the CLA with must be the same email
attached to the commits in your PR.  If it is not, you may need to change
the email and repush the commits.

## Set Release Notes

Every PR must be labeled either `release-note` or `release-note-none`.
This can be done by adding a release-note section to the pr description:

For PRs with a release note:

    ```release-note
    Your release note here
    ```



For PRs without a release note:

    ```release-note
    NONE
    ```

Release notes should be present for any PRs with user visible changes such as
bug-fixes, feature additions, and output format changes.

Additionally, commenting either `/release-note` or `/release-note-none`
will also set the `release-note` or `release-note-none` labels respectively.

## Run e2e Tests

End-to-end tests are run and post the status results to the PR.  PRs authored by
regular contributors have the tests run automatically.  PRs authored by new
community members require the reviewer to mark the tests as safe to run by
commenting `@k8s-bot ok to test`.

If an e2e test fails, `k8s-ci-robot` will comment on the PR with the test history
and the comment-command to re-run that test.  e.g.

>
The magic incantation to run this job again is @k8s-bot unit test this. Please help us cut down flakes by linking to an open flake issue when you hit one in your PR.

## LGTM and Approval

A reviewer will be automatically assigned to your PR by the `k8s-merge-robot`.  The
reviewer will leave comments on your PR.  Once all comments have been addressed,
squash the commits and the reviewer will mark the PR as looking good.  This
can be done with the `/lgtm` command.

## PR merge

After all of the checks have passed, the PR will enter the merge queue:
[http://submit-queue.k8s.io](http://submit-queue.k8s.io).  
The merge queue re-runs the tests for PRs and then merges them if they pass.  The
merge queue is needed to make sure no incompatible changes have been introduced by other
PRs since the tests were last run on your PR.

# Comment Commands Reference

Documented [here](https://github.com/kubernetes/test-infra/blob/master/prow/commands.md)
