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

Manage incoming issues and PRs that are already identified and labeled with the `sig/contributor-experience` particularly in the [Kubernetes community repository](https://github.com/kubernetes/community) but in general across various Kubernetes related GitHub repositories like [Kubernetes repository](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fcontributor-experience).

## People involved

Everyone is welcome to help manage issues and PRs but the work and responsibilities discussed in this document is created with [Triage Team](TODO-describe the team in a separate doc and add link) and [Triage Captain](TODO-describe the team in a separate doc and add link) role in mind.

## Work involved

The main work is to keep track of issues and PRs, ping people, add appropriate labels as needed and report the overall status to the SIG Leads.

- **Validate if an issue is a bug**: Validate if the issue is indeed a bug by following these [instructions](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md#validate-if-the-issue-is-bug).

- **Determine it belongs to contributor-experience SIG**: An issue can be incorrectly labeled with the `sig/contributor-experience` label. If you think an issue does not belong to the `sig/contributor-experience`, add a comment stating so with a reason.

- **Define priorities**: Define priority of an issue or PR by following these [instructions](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md#define-priority).

- **Verify important labels are in place**: Make sure that proper assignees are added in issue, reviewers are added to the PR and milestone is identified. If any of these labels are missing, you can add one if you are authorized to do so. If you can not assign labels or you can not decide the correct label, that’s fine, contact technical leaders of the SIG to do so.

- **Report**: Maintain a list of issues and PRs which seem to be delaying, at risk of not making it in time, or in need of attention by the SIG Leads or other roles, and present it in the weekly SIG meetings.

## How to Escalate
Whenever you find out that an issue or PR is not active and it needs to be taken care, try the following escalation path:
* Leave a comment on the GitHub issue or PR: "This issue hasn't been updated in 3 months. Are you still expecting to complete it in the current milestone?". It's helpful here to @ mention individuals you want an attention from.
* Send a message on the SIG slack channels or mailing list about the problem: It's helpful to directly @ mention the SIG Leads / Owners, and to condense multiple issues into a list, e.g. "Hey, these three issues haven't seen any attention, are they still valid for the current or future milestone?"
* Message individual owners and reviewers directly via Slack and/or email. Individual's email addresses may be harder to find than GitHub ID, but are usually in the Git commit history. Sometimes Slack handles are hard to find. There is no master list mapping human names to GitHub ID, email or Slack ID. If you can't find contact info, asking in the SIG's Slack channel will usually get you pointed in the right direction.
* Escalate to the SIG Leads with suggestions on what to do with non-responsive issues or PR.

