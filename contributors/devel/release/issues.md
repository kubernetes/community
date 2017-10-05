# Targeting issues and PRs to release milestones

This document describes how to target issues and PRs to a release.
The process for shepherding issues into the release by the owner, release team, and GitHub bot
is outlined below.

## Definitions

- *issue owners*: creator, assignees, and user who moved the issue into a release milestone
- *Y days*: refers to business days (using the location local to the release-manager M-F)
- *code slush*: starts when master branch only accepts PRs for release milestone.  no additional feature development is merged after this point.
- *code freeze*: starts 2 weeks after code slush.  only critical bug fixes are accepted into the release codebase.

## Requirements for adding an issue to the milestone

**Note**: Issues with unmet label requirements will automatically be removed from the release milestone.

When adding an issue to a milestone, the Kubernetes bot will check that the following
labels are set, and comment on the issue with the appropriate instructions.  The
bot will attempt to contact the issue creator 3 times (over 3 days)
before automatically removing the issue from the milestone.

Label categories:

- SIG label owner
- Priority
- Issue type

### SIG owner label

The SIG owner label defines the SIG to which the bot will escalate if the issue is not resolved
or updated by the deadline.  If there are no updates after escalation, the
issue may automatically removed from the milestone.

e.g. `sig/node`, `sig/federation`, `sig/apps`, `sig/network`

**Note:**
  - For test-infrastructure issues use `sig/testing`.
  - For GKE and GCE issues use `sig/gcp` once it is created, and `sig/cluster-lifecycle` until then.

### Priority

Priority label used by the bot to determine escalation path before moving an issues
out of the release milestone.  Also used to determine whether or not a release should be
blocked on the resolution of the issue.

- `priority/critical-urgent`: Never automatically move out of a release milestone; continually escalate to contributor and SIG through all available channels.
  - considered a release blocking issue
  - code slush: issue owner update frequency: every 3 days
  - code freeze: issue owner update frequency: daily
  - would require a patch release if left undiscovered until after the minor release.
- `priority/important-soon`: Escalate to the issue owners and SIG owner; move out of milestone after several unsuccessful escalation attempts.
  - not considered a release blocking issue
  - would not require a patch release
  - will automatically be moved out of the release milestone at code freeze
- `priority/important-longterm`: Escalate to the issue owners; move out of the milestone after 1 attempt.
  - even less urgent / critical than `priority/important-soon`
  - moved out of milestone more aggressively than `priority/important-soon`

### Issue type

The issue type is used to help identify the types of changes going into the release over time.
This will allow us to develop a better understanding of what sorts of issues we would miss
with a faster release cadence.

This will also be used to escalate to the correct SIG GitHub team.

- `kind/bug`: Fixes a newly discovered bug.
  - were not known issues at the start of the development period.
- `kind/feature`: New functionality.
- `kind/cleanup`: Adding tests, refactoring, fixing old bugs.

## Bot communication

The bot will communicate the state of an issue in the active milestone via comments and labels.

### Comments

All bot comments will mention issue owners and link to this doc.  Bot
comments in the workflow section of this document will only include
the message but be presumed to include the header and footer in the
following example:

```
@pwittrock @droot

<message>

Additional instructions available [here](<link to this doc>)
```

### Labels

The following labels are used by the bot to track the state of an
issue in the milestone:

 - milestone/incomplete-labels - one or more of the required `kind/`, `priority`/ or `sig/` labels are missing
 - milestone/needs-approval - the `status/approved-for-milestone` label is missing
 - milestone/needs-attention - a status label is missing or an update is required
 - milestone/removed - the issue was removed from the milestone

These labels are mutually exclusive - only one will appear on an issue at once.

## Workflow

1. An issue is added to the current release milestone (either through creation or update)
  - Bot checks to make sure all required labels are set on the issue
  - If any labels are missing, the bot comments listing the missing labels and applies the `milestone/incomplete-labels` label.
    ```
    **Action required**: Issue is missing the following required labels.  Set the labels or the issue
    will be moved out of the milestone within 3 days.

    - priority
    - severity
    ```
  - **If required labels are not applied within 3 days of being moved to the milestone, the bot will move the issue out of the milestone and apply the `milestone/removed` label (unless the issue is critical-urgent).**
  - If the required labels are present, the bot checks whether the issue has the `status/approved-for-milestone` label.
  - If the approved label is not present, the bot comments indicating that the label must be applied by a SIG maintainer and applies the `milestone/needs-approval` label.
    ```
    **Action required**: This issue must have the `status/approved-for-milestone` label applied
    by a SIG maintainer.
    ```
  - If the approved label is present, the bot comments summarizing the label state and removes the other `milestone/*` labels.
    ```
    Issue label settings:

    sig/node: Issue will be escalated to SIG node if needed
    priority/critical-urgent: Never automatically move out of a release milestone.
                              Escalate to SIG and contributor through all available channels.
    kind/bug: Fixes a bug.
    ```
  - **If the approved label is not applied within 7 days of the `milestone/needs-approval` label being applied, the bot will move the issue out of the milestone and apply the `milestone/removed` label (unless the issue is critical-urgent).**
2. If labels change, the bot checks that the needed labels are present and updates its comment and labeling to reflect the issue's current state.
3. Code slush
  - All issues are required to have a status label - one of `status/in-review` or `status/in-progress`.
  - If an issue does not have a status label, the bot comments indicating the required action and applies the `milestone/needs-attention` label.
    ```
    **Action required**: Must specify at most one of `status/in-review` or `status/in-progress`.
    ```
  - **priority/important- issues**
    - The bot includes a warning in the issue comment that the issue will be moved out of the milestone at code freeze.
    ```
    **Note**: This issue must be resolved or labeled as priority/critical-urgent by
    <date of code freeze> or it will automatically be moved out of the milestone
    ```
  - **priority/critical-urgent issues**
    - The bot includes a warning in the issue comment that the issue must be updated regularly.
      ```
      **Note**: This issue is marked as priority/critical-urgent, and is expected to be updated at
      least every 3 days.
      ```
    - If an issue hasn't been updated for more than 3 days, the bot comments and adds the `milestone/needs-attention` label.
      ```
      **Action Required**: This issue is marked as priority/critical-urgent, but has not been updated
      in 3 days.  Please provide an update.
      ```
    - Owner updates can be a short ACK, but should include an ETA for completion and any risk factors.
      ```
      ACK.  In progress
      ETA: DD/MM/YYYY
      Risks: Complicated fix required
      ```

      ```
      ACK.  In progress
      ETA: ???
      Risks: Root cause unknown.
      ```
4. Code freeze
  - **priority/important- issues**
    - The bot removes non-blocker issues from the milestone, comments as to why this was done, and adds the `milestone/removed` label.
      ```
      **Important**: Code freeze is in effect and only issues with priority/critical-urgent may remain
      in the active milestone.  Removing it from the milestone.
      ```
  - **priority/critical-urgent issues**
    - If an issue has not been updated within 2 days, the bot comments and adds the `milestone/needs-attention` label.

## Escalation

SIGs will have issues needing attention escalated through the following channels

- Comment mentioning the sig team appropriate for the issue type
- Email the SIG googlegroup list
  - bootstrapped with the emails from the [community sig list](https://github.com/kubernetes/community/blob/master/sig-list.md)
  - maybe configured to email alternate googlegroup
  - maybe configured to directly email SIG leads or other SIG members
- Message the SIG slack channel, mentioning the SIG leads
  - bootstrapped with the slackchannel and SIG leads from the [community sig list](https://github.com/kubernetes/community/blob/master/sig-list.md)
  - maybe configured to message alternate slack channel and users

## Issues tracked in other repos

Some issues are filed in repos outside the [kubernetes repo].  The bot must also be run against these
repos and follow the same pattern.  The release team can query issues across repos in the kubernetes org using
a query like [this](https://github.com/issues?utf8=%E2%9C%93&q=is%3Aopen+is%3Aissue+milestone%3Av1.7+org%3Akubernetes+)

If the bot is not setup against the split repo, the repo owners should setup an umbrella tracking issue
in the kubernetes/kubernetes repo and aggregate the status.

`Release 1.<minor version> blocking umbrella: <repo name> (size: <number of open issues>)`

it must also include:

- a link to the repo with a query for issues in the milestone. See [this](https://github.com/kubernetes/kubectl/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20milestone%3Av1.7) example.
- a list of unresolved issues blocking the release. See
[this](https://github.com/kubernetes/kubernetes/issues/47747) example.


[kubernetes repo]: (https://github.com/kubernetes/kubernetes)
