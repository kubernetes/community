# Issue and PR Management

## Purpose

Speed up management of issues and PRs labeled with `sig/contributor-experience`.

The SIG Contributor Experience [issues](https://github.com/kubernetes/community/labels/sig%2Fcontributor-experience) and [PRs](https://github.com/kubernetes/community/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aopen+label%3Asig%2Fcontributor-experience+) are labeled with `sig/contributor-experience`.

Following are few example searches on issue and PR:
* [Open issues](https://github.com/kubernetes/community/labels/sig%2Fcontributor-experience)
* [Open issues for a specific milestone](https://github.com/kubernetes/community/issues?utf8=%E2%9C%93&q=is%3Aopen+label%3Asig%2Fcontributor-experience+milestone%3AMay)
* [Open PRs](https://github.com/kubernetes/community/pulls?utf8=%E2%9C%93&q=is%3A+pr+is%3Aopen+label%3Asig%2Fcontributor-experience+)
* [Open PRs for a specific milestone](https://github.com/kubernetes/community/pulls?utf8=%E2%9C%93&q=is%3Apr+label%3Asig%2Fcontributor-experience+milestone%3AMay+is%3Aopen)

## Scope

Manage incoming issues and PRs that are already identified and labeled with the `sig/contributor-experience` particularly in the [Kubernetes community repository](https://git.k8s.io/community) but in general across various Kubernetes related GitHub repositories like [Kubernetes repository](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fcontributor-experience).

## People involved

Everyone is welcome to help manage issues and PRs but the work and responsibilities discussed in this document is created with Triage team and Triage Captain roles in mind.

## Work involved

The main work is to keep track of issues and PRs, ping people, add appropriate labels as needed and report the overall status to the SIG Leads.

  - Validate if an issue is a bug - Validate if the issue is indeed a bug by following these [instructions](https://git.k8s.io/community/contributors/guide/issue-triage.md#validate-if-the-issue-is-bug).

  - Determine it belongs to contributor-experience SIG - An issue can be incorrectly labeled with the `sig/contributor-experience` label. If you think an issue does not belong to the `sig/contributor-experience`, add a comment stating so with a reason. Example response:

```code

`/remove-sig contributor-experience`

I don't think this is a SIG Contributor Experience issue because ....

If you think this is incorrect, please provide a reason and add it back with /sig contributor-experience label as a single line comment.

```

  - Define priorities - Define priority of an issue or PR by following these [instructions](https://git.k8s.io/community/contributors/guide/issue-triage.md#define-priority).

  - Verify important tasks and labels are in place - Make sure that one or more of the following tasks and labels are identified for issues and PRs. If any of these are missing, you can add one if you are authorized to do so. If you can not assign or can not decide the correct one, that’s fine, contact leaders of the SIG to do so.

    - Issues

        - Milestone
        - area/*
        - kind/*
        - good first issue
        - help wanted

    - PR

        - Milestone
        - Reviewers
        - Assignees
        - area/*
        - kind/*

  - Report - Maintain a list of issues and PRs which seem to be delaying, at risk of not making it in time, labeled with `lifecycle/stale` and `lifecycle/rotten` or in need of attention by the SIG Leads or other roles. Present this list in the weekly SIG meetings. If needed, also send the list of issues and PRs that aren't getting enough attention to the mailing list biweekly.

## How to Escalate
Whenever you find out that an issue or PR is not active and it needs to be taken care, try the following escalation path:
* Leave a comment on the GitHub issue or PR: "This issue hasn't been updated in 3 months. Are you still expecting to complete it in the current milestone?". It's helpful here to @ mention individuals you want an attention from.
* Send a message on the SIG slack channels or mailing list about the problem: It's helpful to directly @ mention the SIG Leads / Owners, and to condense multiple issues into a list, e.g. "Hey, these three issues haven't seen any attention, are they still valid for the current or future milestone?"
* Message individual owners and reviewers directly via Slack and/or email ID, if such an ID is available.
* Escalate to the SIG Leads with suggestions on what to do with non-responsive issues or PR.

