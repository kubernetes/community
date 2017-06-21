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

## Workflow

1. An issue is added to the current release milestone (either through creation or update)
  - Bot checks to make sure all required labels are set on the issue
  - If and only if any labels are missing, comment listing the missing labels and include a link to this document.  Mention all issue owners.
    ```
    @pwittrock @droot
    
    **Action required**: Issue is missing the following required labels.  Set the labels or the issue will be moved
    out of the milestone within 2 days.
    
    - priority
    - severity
    
    Additional instructions available [here](<link to this doc>)
    ```
  - If and only if all labels are present, comment summarizing the state.  Mention all issue owners.
    ```
    @pwittrock @droot
    
    Issue label settings:
    sig/node: Issue will be escalated to SIG node if needed
    priority/critical-urgent: Never automatically move out of a release milestone.  Escalate to SIG and contributor through all available channels.
    kind/bug: Fixes a bug.
    
    Additional instructions available [here](<link to this doc>)
    ```
  - **If required labels are not applied within 3 days of being moved to the milestone,  the issue is moved back out of the milestone.  (unless it is critical-urgent)**
2. If labels change, bot checks that the needed labels are present and comments with updated meaning.
3. Code slush
  - **priority/important-* issues**
    - bot comments mentioning issue owners, expects a 1 time response.
    ```
    **Action Required**:
    This issue is marked as priority/important-x, and is in the 1.y milestone.  Please confirm the following
    or the issue will be removed from the milestone.
    
    1. it is still targeted at the milestone
    2. it has not been completed and should remain open
    3. it is being actively worked on by the assignee
    
    **Note**: This issue must be resolved or moved out of the milestone by <date of code freeze> or it will
    automatically be moved out of the milestone
    ```
    - after 2 days without a reply, bot follows escalation procedure derived from priority label
      - important-longterm, remove from milestone and notify the owners
      - important-soon, escalate to SIG daily for 2 days.
    - removal should add the `release/removed-from-milestone` label and contain the following message
      ```
      **Important**:
      There has been no confirmation that this issue belongs in the v1.X milestone.
      Removing it from the milestone.
      ```
  - **priority/critical-urgent issues**
    - bot comments on issues not updated within 3 days mentioning issue owners, expects updates every 3 days going forward.
      ```
      **Action Required**:
      This issue is marked as priority/critical-urgent, but has not been updated in 3 days.  Please provide
      an update or the issue will be escalated to the SIG.
      ```
    - owner updates can be a short ACK, but should include an ETA for completion and any risk factors.
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
  - **priority/important-* issues**
    - bot comments mentioning issue owners, expects the issue to be escalated in priority or removed from milestone within 2 days.
    ```
    **Action Required**:
    This issue is marked as priority/important-x, and is in the 1.y milestone.  Only critical-urgent issue
    are being tracked for the 1.y release.  Escalate the priority to critical-urgent within 1 day, **or the issue
    will be removed from the milestone**.
    ```
    - bot immediately escalates to the SIG.  escalation message includes link to all priority/important issues
      for the SIG remaining in the milestone.
  - **priority/critical-urgent issues**
    - bot comments on issues not updated within 1 day mentioning issue owners, expects updates every day going forward.
    - bot escalates to SIG after 2 days without updates and applies the label `release/needs-attention`

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