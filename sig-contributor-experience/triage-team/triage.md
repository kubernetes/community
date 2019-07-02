# Issue and PR Management

## Purpose

Speed up management of issues and PRs under [k/community repo](https://git.k8s.io/community).

## Scope

Incoming issues and PRs in the [Kubernetes community repository](https://git.k8s.io/community). The goal is to maintain a healthy community repository by timely handling issues and PRs. The scope may be enhanced in the future.

## People involved

Everyone is welcome to help manage issues and PRs, however, the work and responsibilities discussed in this document is created with `sig/contributor-experience` triage team and [OWNERS](https://git.k8s.io/community/sig-contributor-experience/OWNERS) in mind.

## Work involved

The main work is to keep track of issues and PRs, ping people, add labels and milestone, and discuss the overall status with the SIG leads. If you can not add or can not decide the correct label or milestone, that’s fine, contact leaders of the SIG to do so.

### Issues

#### Validate if an issue is a bug or a support request

Validate if the issue is indeed a bug by following these [instructions](https://git.k8s.io/community/contributors/guide/issue-triage.md#determine-if-its-a-support-request). If it is a support request, provide a response similar to this [format](https://git.k8s.io/community/contributors/guide/issue-triage.md#user-support-response-example) and close it.

#### Determine SIG

In order to determine a correct SIG, charter of the SIG can be helpful. For example, you can determine if an issue belongs to one of the `sig/contributor-experience` [Subprojects](https://git.k8s.io/community/sig-contributor-experience#subprojects) by looking at the [charter](https://git.k8s.io/community/sig-contributor-experience/charter.md).

An issue can have an incorrect SIG label. If you think an issue does not belong to the labeled SIG, remove the label and provide a reason by adding a comment. For example, provide this response when an issue is incorrectly labeled with the `sig/contributor-experience`:

```code

`/remove-sig contributor-experience`

I don't think this is a SIG Contributor Experience issue because ....

If you think this is incorrect, please provide a reason and add it back with /sig contributor-experience label as a single line comment.

```

#### Determine milestone

[Milestone](https://git.k8s.io/community/milestones) should be defined considering the [priority](#determine-priorities). If the priority is `priority/critical-urgent`, the milestone should be the current milestone. If priority is not critical, the milestone can be set to current or future one. Use your experience and reference of related past issues while determining a milestone.

#### Determine projects

The project is defined based off of [Subprojects](https://git.k8s.io/community/sig-contributor-experience#subprojects).

#### Determine area

An issue `area/*` label is defined around the [Subprojects](https://git.k8s.io/community/sig-contributor-experience#subprojects). e.g. `area/devstats`. Determine the area label considering Subprojects, and using your experience and reference of related past issues.

#### Determine priorities

WIP. For now, define priority of an issue by following these generic [instructions](https://git.k8s.io/community/contributors/guide/issue-triage.md#define-priority).

#### Determine kind

For the `kind/*` label, we follow the instruction as defined in the [release guidelines](https://git.k8s.io/community/contributors/devel/sig-release/release.md#issuepr-kind-label).

#### Determine if an issue is a candidate for new contributors

An issue can be a good candidate for a new contributor to work. Such issues should be labeled with `good first issue` or `help wanted` as described in more detail [here](https://git.k8s.io/community/contributors/guide/help-wanted.md).

#### Discuss the work

Maintain a list of issues and PRs which seem to be delaying, at risk of not making it in time, labeled with `lifecycle/stale` and `lifecycle/rotten` or in need of attention by the SIG leads or other roles. Discuss this list in the bi-weekly SIG meetings. If needed, also send the list of issues and PRs that aren't getting enough attention to the mailing list biweekly.

### PRs

TODO - PR specific work.

## How to Escalate

Whenever you find out that an issue or PR is not active and it needs to be taken care, try the following escalation path.

### Leave a comment in the issue or PR

"This issue hasn't been updated in 3 months. Are you still expecting to complete it in the current milestone?". It's helpful here to @ mention individuals you want an attention from.

### Send a message on the SIG slack channel or mailing list

It's helpful to directly @ mention the SIG Leads / Owners, and to condense multiple issues into a list, e.g. "Hey, these three issues haven't seen any attention, are they still valid for the current or future milestone?"

### Message individual owners and reviewers directly

Send a direct message via Slack and/or email ID, if such an ID is available.

### Escalate to the SIG Leads

Talk to SIG leads with suggestions on what to do with non-responsive issues or PR.

