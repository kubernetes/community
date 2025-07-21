---
title: "Issue Triage Guidelines"
weight: 10
description: |
  These guidelines serve as a primary document for triaging incoming issues to
  Kubernetes. SIGs and projects are encouraged to use this guidance as a
  starting point, and customize to address specific triaging needs.
---

## Table of Contents
- [Scope](#scope)
- [What Is Triaging?](#what-is-triaging)
- [Why Is Triaging Beneficial?](#why-is-triaging-beneficial)
- [How to Triage: A Step-by-Step Flow](#how-to-triage-a-step-by-step-flow)
   - [Triage-Related Tools](#triage-related-tools)
      - [Permissions and the Bot](#permissions-and-the-bot)
      - [Triage Party](#triage-party)
      - [GitHub Project Boards](#github-project-boards)
      - [DevStats](#devstats)
   - [Process Pointers and Advice from SIGs](#process-pointers-and-advice-from-sigs)
      - [Running a Triage Meeting: Tips from api-machinery](#running-a-triage-meeting-tips-from-api-machinery)
      - [Triage Guide by cluster-lifecycle](#triage-guide-by-cluster-lifecycle)
 - [Step One: Review Newly Created Open Issues](#step-one-review-newly-created-open-issues)
   - [Conducting Searches](#conducting-searches)
 - [Step Two: Triage Issues by Type](#step-two-triage-issues-by-type)
   - [Support Requests](#support-requests)
   - [Abandoned or Wrongly Placed Issues](#abandoned-or-wrongly-placed-issues)
   - [Needs More Information](#needs-more-information)
   - [Bugs](#bugs)
   - [Help Wanted/Good First Issues](#help-wantedgood-first-issues)
   - [Kind Labels](#kind-labels)
 - [Step Three: Define Priority](#step-three-define-priority)
 - [Step Four: Find and Set the Right SIG(s) to Own an Issue](#step-four-find-and-set-the-right-sigs-to-own-an-issue)
   - [Self-Assigning](#self-assigning)
 - [Step Five: Follow Up](#step-five-follow-up)
   - [If No PR Is Created for an Issue Within the Current Release Cycle](#if-no-pr-is-created-for-an-issue-within-the-current-release-cycle)
   - [If a SIG Label Is Assigned, but No Action Is Taken Within 30 Days](#if-a-sig-label-is-assigned-but-no-action-is-taken-within-30-days)
   - [If an Issue Has No Activity After 90 Days](#if-an-issue-has-no-activity-after-90-days)
 - [Further Notes](#further-notes)
   - [Support Requests: Channels](#support-requests-channels)
   - [User Support Response: Example](#user-support-response-example)

## Scope

These guidelines serve as a primary document for triaging incoming issues to Kubernetes. SIGs and projects are encouraged to use this guidance as a starting point, and customize to address specific triaging needs.

**Note:** These guidelines only apply to the Kubernetes repository. Usage for other Kubernetes-related GitHub repositories is TBD.

## What Is Triaging?

Issue triage is a process by which a SIG intakes and reviews new GitHub issues and requests, and organizes them to be actioned—either by its own members, or by other SIGs. Triaging involves categorizing issues and pull requests based on factors such as priority/urgency, SIG ownership of the issue, and the issue kind (bug, feature, etc.).

Triage can happen asynchronously and continuously, or in regularly scheduled meetings. Several Kubernetes SIGs and projects have adopted their own approaches to triaging.

## Why Is Triaging Beneficial?

SIGs who triage regularly say it offers a number of benefits, such as:

- Speeding up issue management
- Keeping contributors engaged by shortening response times
- Preventing work from lingering endlessly
- Replacing special requests and one-offs with a neutral process that acts like a boundary
- Greater transparency, interesting discussions, and more collaborative, informed decision-making
- Building prioritization, negotiation and decision-making skills, which are critical to most tech roles
- Reinforcement of SIG community and culture

People who enjoy product management and iterating on processes tend to enjoy triaging because it empowers their SIGs to maintain a steady, continuous flow of work that is assessed and prioritized based on feedback and value.

# How to Triage: A Step-by-Step Flow

This guide walks you through a standard triaging process, beginning with tools and tips.

## Triage-Related Tools

These are tools that your SIG can use to make the triage process simpler, more efficient and faster.

### Permissions and the Bot

Opening new issues and leaving comments on other people's issues are possible for all contributors. However, permission to assign specific labels (such as `triage`), change milestones, or close other contributors issues is only granted to the author of an issue, assignees, and organization members. For this reason, we use a bot to manage labelling and triaging. For a full list of the bot's commands and permissions, see the [Prow command reference page](https://go.k8s.io/bot-commands).

### Triage Party

[Triage Party](https://github.com/google/triage-party) is a tool for triaging incoming GitHub issues for large open-source projects, built with the GitHub API. Made public in April 2020, it facilitates "massively multi-player GitHub triage" and reduces contributor response latency.

Its features include:
- Queries across multiple repositories
- Queries that are not possible on GitHub:
   - conversation direction (`tag: recv`, `tag: send`)
   - duration (`updated: +30d`)
   - regexp (`label: priority/.*`)
   - reactions (`reactions: >=5`)
   - comment popularity (`comments-per-month: >0.9`)
- Multiplayer mode: for simultaneous group triage of a pool of issues
- Button to open issue groups as browser tabs (pop-ups must be disabled)
- "Shift-Reload" for live data pull

### GitHub Project Boards

GitHub offers project boards, set up like [kanban boards](https://en.wikipedia.org/wiki/Kanban), to help teams organize and track their workflow in order to get work done. The Release Team has come to depend on [their project board](https://github.com/orgs/kubernetes/projects/68) for planning new Kubernetes releases; they also use it as an archive to show the work done for past releases.

Other SIGs are also using project boards:
- [Apps](https://github.com/orgs/kubernetes/projects/167)
- [Auth](https://github.com/orgs/kubernetes/projects/116)
- [Scheduling](https://github.com/orgs/kubernetes/projects/165)

We encourage more SIGs to use project boards to enhance visibility and tracking. If you'd like some help getting started, visit [GitHub's documentation](https://help.github.com/en/github/managing-your-work-on-github/about-project-boards) or reach out to [SIG Contributor Experience](/sig-contributor-experience/README.md#contact).

### DevStats

The CNCF has created a [suite of Grafana dashboards and charts](https://devstats.cncf.io/) for collecting metrics related to all the CNCF projects. The [Kubernetes dashboard](https://k8s.devstats.cncf.io/d/12/dashboards?orgId=1&refresh=15m) can be used to help SIGs view real-time metrics on many aspects of their workflow, including:
- [Issue Velocity](https://k8s.devstats.cncf.io/d/12/dashboards?from=1587157094179&orgId=1&refresh=15m&to=1587758294179&viewPanel=8): How quickly issues are resolved
- [PR Velocity](https://k8s.devstats.cncf.io/d/12/dashboards?from=1587157166022&orgId=1&refresh=15m&to=1587758366022&viewPanel=9): Including PR workload per SIG, PR time to approve and merge, and other data

## Process Pointers and Advice from SIGs

Several SIGs consistently meet weekly or monthly to triage issues. Here are some details about their processes.

### Running a Triage Meeting: Tips from api-machinery

The [api-machinery SIG](/sig-api-machinery) has found that triage meetings offer valuable opportunities for newcomers to listen, learn, and start contributing. The SIG hold triage meetings every Tuesday and Thursday and archive recordings via their [YouTube playlist](https://www.youtube.com/playlist?list=PL69nYSiGNLP21oW3hbLyjjj4XhrwKxH2R).  [Watch an example of one of their meetings](https://www.youtube.com/watch?v=bRptR9vd4S8&list=PL69nYSiGNLP21oW3hbLyjjj4XhrwKxH2R&index=2&t=13s).

In a typical triage meeting, api-machinery members sort through every issue that they haven't triaged since the previous meeting, using a simple query and issue number to track open PRs and issues. They usually follow this process:
1. Read through the comments and the code briefly to understand what the issue is about.
1. Determine by consensus if it belongs to the api-machinery SIG or not. If not, remove the `sig/api-machinery` label.
1. Label other SIGs, if appropriate
1. Discuss briefly the technical implications
1. Assign people with expertise in the domain to review, comment, reject, etc.

The api-machinery SIG has found that consistently meeting on a regular, fixed schedule is key to the success of a triaging effort. More frequent, small meetings are better than infrequent, large meetings. They also offer a few other pointers for successful triage meetings:
- We try to balance the load, and ask people if they are okay taking on an issue before assigning it to them.
- We skip issues that are closed.
- We also skip cherrypicks, because we consider that the code change was reviewed in the original PR.
- We ensure participation from the entire SIG and support company diversity.
- We use this opportunity to add [`help wanted` and `good first issue`](#help-wantedgood-first-issues) labels.

### Triage Guide by cluster-lifecycle

The cluster-lifecycle SIG has developed a [triaging page](/sig-cluster-lifecycle/grooming.md) detailing their process, including the [Milestones](/sig-cluster-lifecycle/grooming.md#planning-a-milestone) stage. Here is a [March 2020 presentation](https://www.youtube.com/watch?v=Q07_PfkNjlw) delivered to the SIG chairs and leads group on their process.

## Step One: Review Newly Created Open Issues

The first step in a successful triage meeting is reviewing newly created open issues. Kubernetes issues are listed [here](https://github.com/kubernetes/kubernetes/issues). Labels are the primary tools for triaging. [Here's a comprehensive label list](https://github.com/kubernetes/kubernetes/labels).

New issues are automatically assigned a `needs-triage` label indicating that these issues are currently awaiting triage. After triaging an issue, the issue owning SIG will use the bot command `/triage accepted`. This command removes the `needs-triage` label and adds the `triage/accepted` label.

Note that adding labels requires Kubernetes GitHub org membership. If you are not an org member, you should add your triage findings as a comment.

### Conducting Searches

GitHub allows you to filter out types of issues and pull requests, which helps you discover items in need of triaging. This table includes some predetermined searches for convenience:


| Search                                                                                                       | What it sorts                                           |
|--------------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| [created-asc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acreated-asc)     | Untriaged issues by age                                 |
| [needs-sig](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Aneeds-sig)        | Issues that need to be assigned to a SIG                |
| [`is:open is:issue`](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue)                 | Newest incoming issues                                  |
| [comments-desc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-desc) | Busiest untriaged issues, sorted by # of comments       |
| [comments-asc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-asc)   | Issues that need more attention, based on # of comments |


We suggest preparing your triage by filtering out the oldest, unlabelled issues and pull requests first.

## Step Two: Triage Issues by Type

Use [these `triage/` and `kind/support` labels](https://github.com/kubernetes/kubernetes/labels?utf8=%E2%9C%93&q=triage%2F+kind%2Fsupport) to find open issues that can be quickly closed. A triage engineer can add the appropriate labels.

Depending on your permissions, either close or comment on any issues that are identified as support requests, duplicates, or not-reproducible bugs, or that lack enough information from the reporter.

### Support Requests

Some people mistakenly use GitHub issues to file support requests. Usually they are asking for help configuring some aspect of Kubernetes. To handle such an issue, direct the author to use our [support request channels](#support-requests-channels). Then apply the `kind/support` label, which is directed to our support structures, and apply the `close` label.

Please find more detailed information about Support Requests in the [Further Notes section](#further-notes).

### Abandoned or Wrongly Placed Issues

If an issue is abandoned or in the wrong place, either close or comment on it.

### Needs More Information

The `triage/needs-information` label indicates an issue needs more information in order for work to continue; comment on or close it.

### Bugs

First, validate if the problem is a bug by trying to reproduce it.

If you can reproduce it:
* [Define its priority](#step-three-define-priority).
* Search for duplicates to see if the issue has been reported already. If a duplicate is found, let the issue reporter know, reference the original issue, and close the duplicate.

If you can't reproduce it:
* Contact the issue reporter with your findings .
* Close the issue if both the parties agree that it could not be reproduced.

If you need more information to further work on the issue:
* Let the reporter know it by adding an issue comment. Include `/triage needs-information` in the comment to apply the `triage/needs-information` label.

In all cases, if you do not get a response within 20 days, close the issue with an appropriate comment. If you have permission to close someone else's issue, first `/assign` the issue to yourself, then `/close` it. If you do not, please leave a comment describing your findings.

### Help Wanted/Good First Issues

To identify issues that are specifically groomed for new contributors, we use the [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)
and [good first issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) labels. To use these labels:
* Review our specific [guidelines](/contributors/guide/help-wanted.md) for how to use them.
* If the issue satisfies these guidelines, you can add the `help wanted` label with the `/help` command.
and the `good first issue` label with the `/good-first-issue` command. Please note that adding the `good first issue` label will also automatically add the `help wanted` label.
* If an issue has these labels but does not satisfy the guidelines, please ask for more details to be added to the issue or remove the labels using the `/remove-help` or `/remove-good-first-issue` commands.

### Kind Labels

Usually the `kind` label is applied by the person submitting the issue. Issues that feature the wrong `kind` (for example, support requests labelled as bugs) can be corrected by someone triaging; double-checking is a good approach. Our [issue templates](https://github.com/kubernetes/kubernetes/issues/new/choose) aim to steer people to the right kind.

## Step Three: Define Priority

We use GitHub labels for prioritization. If an issue lacks a `priority` label, this means it has not been reviewed and prioritized yet.

We aim for consistency across the entire project. However, if you notice an issue that you believe to be incorrectly prioritized, please leave a comment offering your counter-proposal and we will evaluate it.


|Priority label|What it means|Examples|
|---|---|---|
| `priority/critical-urgent` | Team leaders are responsible for making sure that these issues (in their area) are being actively worked on—i.e., drop what you're doing. Stuff is burning. These should be fixed before the next release. | user-visible bugs in core features <br> broken builds <br> tests and critical security issues |
| `priority/important-soon` | Must be staffed and worked on either currently or very soon—ideally in time for the next release. Important, but wouldn't block a release. | [**XXXX**] |
| `priority/important-longterm` | Important over the long term, but may not be currently staffed and/or may require multiple releases to complete. Wouldn't block a release. | [**XXXX**]|
| `priority/backlog`  | General agreement that this is a nice-to-have, but no one's available to work on it anytime soon. Community contributions would be most welcome in the meantime, though it might take a while to get them reviewed if reviewers are fully occupied with higher-priority issues—for example, immediately before a release.| [**XXXX**]  |
| `priority/awaiting-more-evidence` | Possibly useful, but not yet enough support to actually get it done. | Mostly placeholders for potentially good ideas, so that they don't get completely forgotten, and can be referenced or deduped every time they come up |


## Step Four: Find and Set the Right SIG(s) to Own an Issue

Components are divided among [Special Interest Groups (SIGs)](/sig-list.md). [The bot](https://go.k8s.io/bot-commands) assists in finding a proper SIG to own an issue.

* For example, typing `/sig network` in a comment should add the `sig/network` label.
* Multiword SIGs use dashes: for example, `/sig cluster-lifecycle`.
* Keep in mind that these commands must be on their own lines, and at the front of the comment.
* If you are not sure about who should own an issue, defer to the SIG label only.
* If you feel an issue should warrant a notification, ping a team with an `@` mention, in this format: `@kubernetes/sig-<group-name>-<group-suffix>`. Here, the `<group-suffix>` can be one of:
    - `bugs`
    - `feature-requests`
    - `pr-reviews`
    - `test-failures`
    - `proposals`
  For example: `@kubernetes/sig-cluster-lifecycle-bugs, can you have a look at this?`

### Self-Assigning

If you think you can fix the issue, assign it to yourself with *just* the `/assign` command. If you cannot self-assign for permissions-related reasons, leave a comment that you'd like to claim it and [begin working on a PR](github-workflow.md).

When an issue already has an assignee, **do not** assign it to yourself or create a PR without talking to the existing assignee or going through the [Follow Up](#step-five-follow-up) steps as described in this document. Creating a PR when someone else is already working on an issue is not a good practice and is discouraged.

## Step Five: Follow Up

### If No PR is Created for an Issue Within the Current Release Cycle

If an issue is owned by a developer but a PR has not been created within 30 days, a triage engineer should contact the issue owner and ask them to either create a PR or release ownership.

### If a SIG Label Is Assigned, but No Action Is Taken Within 30 Days

If you find an issue with a SIG label assigned, but there's no evidence of movement or discussion within 30 days, then gently poke the SIG about this pending issue. Also, consider attending one of their meetings to bring up the issue.

### If an Issue Has No Activity After 90 Days

When an issue goes 90 days without activity, the [k8s-triage-robot](https://github.com/k8s-triage-robot) adds the `lifecycle/stale` label to that issue. You can block the bot by applying the `/lifecycle frozen` label preemptively, or remove the label with the `/remove-lifecycle stale` command. The k8s-triage-robot adds comments in the issue that include additional details. If you take neither step, the issue will eventually be auto-closed.

## Further Notes

### Support Requests: Channels

These should be directed to the following:
* [User documentation](https://kubernetes.io/docs/home/) and
[troubleshooting guide](https://kubernetes.io/docs/tasks/debug/)
* [Slack](https://kubernetes.slack.com) ([registration](https://slack.k8s.io))
* [Discussion forums](https://discuss.kubernetes.io)

### User Support Response: Example

If you see support questions on dev@kubernetes.io or issues asking for
support, try to redirect them to Discuss. Here is an example response:

> Please re-post your question to our [Discussion Forums](https://discuss.kubernetes.io).
>
> We are trying to consolidate the channels to which questions for help/support
> are posted so that we can improve our efficiency in responding to your requests,
> and to make it easier for you to find answers to frequently asked questions and
> how to address common use cases.
>
> We regularly see messages posted in multiple forums, with the full response
> thread only in one place or, worse, spread across multiple forums. Also, the
> large volume of support issues on GitHub is making it difficult for us to use
> issues to identify real bugs.
>
> Members of the Kubernetes community use Discussion Forums to field
> support requests. Before posting a new question, please search these for answers
> to similar questions, and also familiarize yourself with:
>
>  * [user documentation](https://kubernetes.io/docs/home/)
>  * [troubleshooting guide](https://kubernetes.io/docs/tasks/debug/)
>
> Again, thanks for using Kubernetes.
>
> The Kubernetes Team
