---
title: "Issue Triage Guidelines"
weight: 10
slug: "issue-triage"
---

# Issue Triage: A Primer

## Table of Contents
- [Scope]
- [What is triaging?]
- [Why is triaging beneficial?]
- [Who should triage?]
- [A Sample Step-by-Step]
 - [Step One: Review newly created open issues]
  - [Conducting Searches]
  - [Permissions and the Bot]
 - [Step Two: Triage Issues by Type]
  - [Abandoned or wrongly placed issue]
  - [Support requests]
  - [Bugs]
  - [Help Wanted/Good First Issues]]
  
    - [Set the `kind` issue to validate type (feature, documentation, bug, etc.)]
    - [Set a `priority` label to define urgency]
    - [Set a `sig/` label to define ownership]
 - [Three: Follow up]
    - [Poke issue owner if PR is not created for it in 30 days]
   - [Poke SIG if a SIG label is assigned but no comment was added by SIG in 30 days]
 - [Four: Plan milestones]
 - [Five: Close out issues]

## Scope
These guidelines serve as a primary document for triaging incoming issues to Kubernetes. SIGs and projects are encouraged to use this guidance as a starting point, and customize to address specific triaging needs.

**Note:** These guidelines only apply to the Kubernetes repository. Usage for other Kubernetes-related GitHub repositories is TBD.

## What is triaging?
For our purposes, issue triage is a process by which a SIG intakes and reviews new GitHub issues and requests, and organizes them to be actioned—either by itself, or by other SIGs. Kubernetes SIGs do this mainly by applying GitHub labels that categorize issues and pull requests based on factors such as:
- priority/urgency
- the SIG or SIGs responsible for handling the issue or pull request
- the kind of work: bug, feature, etc.

Triage can happen asynchronously and continously, or in regularly scheduled meetings. Several Kubernetes SIGs and projects have adopted their own approaches to triaging. 

## Why is triaging beneficial?
Triaging offers several benefits to SIGs:
- Speeds up issue management
- Quicker response times keeps contributors engaged
- Prevents work from falling through the cracks
- Reduces "special requests" and sudden context switches—issues and PRs are managed via a process, and response-time SLAs can be communicated upfront
- Leads to greater transparency, broader input and more informed decision-making about priorities

## Who should triage?
Everyone belonging to a SIG is encouraged to triage. You might find it fulfilling for any of the following reasons:
- it leads to interesting discussions within your SIG
- it maintains a healthy, positive contributor experience
- it helps build prioritization, negotiation and decision-making skills, which are critical to most tech roles
- it reinforces SIG community and culture

That said, people who enjoy product management and iterating on processes tend to enjoy triaging because it empowers their SIGs to maintain a steady, continuous flow of work that is assessed and prioritized based on feedback and value. 

# A Sample Step-by-Step
This aims to walk you through a standard triaging process.

## Step One: Review newly created open issues
Kubernetes issues are listed at https://github.com/kubernetes/kubernetes/issues. New, untriaged issues start out without any labels attached. 

Labels are the primary tools for triaging. The detailed label list resides here:
https://github.com/kubernetes/kubernetes/labels.

### Conducting Searches
GitHub allows you to filter out types of issues and pull requests, which helps you discover items in need of triaging. This table includes some predetermined searches for convenience:

|  Search | What it sorts  |
|---|---|
| [created-asc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acreated-asc)  | untriaged issues by age |
|  [needs-sig](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Aneeds-sig) | issues that need to be assigned to a SIG  |
| [`is:open is:issue`](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue)   | Newest incoming issues  |
| [comments-desc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-desc)   | busiest untriaged issues, sorted by # of comments  |
| [comments-asc](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-asc)   | Issues that need more attention, based on # of comments  |

We suggest preparing your triage by filtering out the oldest, unlabelled issues and/or pull requests first.

### Permissions and the Bot
Opening new issues and leaving comments on other people's issues are possible for all contributors. However, permission to assign specific labels (e.g. `triaged`), change milestones, or close other contributors' issues is only granted to the author of an issue, assignees, and component organization members. For this reason, we use a bot to manage labelling and triaging. The bot has a set of [commands and permissions](https://go.k8s.io/bot-commands).  

## Step Two: Triage Issues by Type
Use [these labels](https://github.com/kubernetes/kubernetes/labels?utf8=%E2%9C%93&q=triage%2F+kind%2Fsupport+is%3Aopen) to find open issues that can be quickly closed. A triage engineer can add the appropriate labels.

Depending on your permissions, either close or comment on any issues that are identified as support requests, duplicates, or not-reproducible bugs, or that lack enough information from the reporter.
 
### Support requests
Some people mistakenly use GitHub issues to file support requests— usually asking for help configuring some aspect of Kubernetes.
* First, apply the `triage/support` label, which is directed to our support structures (see below) 
* Then, close or comment

Please find more detailed information about Support Requests in the [Footnotes section](#footnotes).

### Duplicates
* Duplicates of other open issues should have the self-readable label `triage/duplicate`, then be commented on or closed. 

### Abandoned or wrongly placed issues
Depending on your permissions, either close or comment on it.

### Needs more information
* The `triage/needs-information` label indicates an issue needs more information in order to work on it; comment on or close it.

### Unresolved
* The `triage/unresolved` label indicates an issue that can not be resolved.

### Bugs
First, validate if the problem is a bug by trying to reproduce it.

If you can reproduce it:
* Define its priority
* Do a quick duplicate search to see if the issue has been reported already. If a duplicate is found, let the issue reporter know it by marking it duplicate. Label such issues as `triage/duplicate`.

If you can't reproduce it:
* label it as a `triage/not-reproducible`
* Contact the issue reporter with your findings 
* Close the issue if both the parties agree that it could not be reproduced.

If you need more information to further work on the issue:
* let the reporter know it by adding an issue comment followed by label `triage/needs-information`.

In all cases, if you do not get a response in 20 days then close the issue with an appropriate comment. If you have permission to close someone else's issue, first `/assign` the issue to yourself, then `/close` it. If you do not, just comment your findings. 

### Help Wanted/Good First Issues
To identify issues that are specifically groomed for new contributors, we use the [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)
and [good first issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) labels. To use these labels:
* review our specific [guidelines](/contributors/guide/help-wanted.md) for how to use them.
* If the issue satisfies these guidelines, you can add the `help wanted` label with the `/help` command
and the `good first issue` label with the `/good-first-issue` command. Please note that adding the `good first issue` label will also automatically add the `help wanted` label.
* If an issue has these labels but does not satisfy the guidelines, please ask for more details to be added to the issue or remove the labels using the `/remove-help` or `/remove-good-first-issue` commands.

## Define priority
We use GitHub labels for prioritization. The absence of a `priority` label
means the issue has not been reviewed and prioritized yet. We aim for consistency across the entire project,
but if you notice an issue that you believe to be incorrectly prioritized,
please let us know by leaving a comment. We will evaluate your counter-proposal.

|  Priority label | What it means  | Examples |
|---|---|---|
| **priority/critical-urgent**  | Team leaders are responsible for making sure that these issues (in their area) are being actively worked on. Someone is expected to drop what they're doing immediately to work on it. Stuff is burning. | * user-visible bugs in core features * broken builds * tests and critical
security issues | 
| **priority/important-soon**  | Must be staffed and worked on either currently,
or very soon, ideally in time for the next release. | [**XXXX**] |  
| **priority/important-longterm**  | Important over the long term, but may not be
currently staffed and/or may require multiple releases to complete. | [**XXXX**] |  
| **priority/backlog**  | General agreement that this is a nice-to-have, but no one's available to work on it anytime soon. Community contributions would be most welcome in the meantime, though it might take a while to get them reviewed if reviewers are fully occupied with higher-priority issues—for example, immediately before a release.| [**XXXX**]  | 
| **priority/awaiting-more-evidence**  | Possibly useful, but not yet enough
support to actually get it done. | * mostly placeholders for potentially
good ideas, so that they don't get completely forgotten, and can be referenced
/deduped every time they come up | 

## Find the right SIG(s)
Components are divided among [Special Interest Groups (SIGs)](/sig-list.md). Find a proper SIG for the ownership of the issue using the bot:

* Typing `/sig network` in a comment should add the sig/network label, for
example.
* Multiword SIGs use dashes, for example `/sig cluster-lifecycle`.

Keep in mind that these commands must be on their own lines, and at the front of the
comment.

## Set ownership
If you are not sure of who should own an issue, defer to the
SIG label only. If you feel the issue should warrant a notification, you can ping
a team with an @ mention, in this format, `@kubernetes/sig-<group-name>-<group-suffix>`.
Here the `<group-suffix>` can be one of `bugs, feature-requests, pr-reviews, test-failures, proposals`.
For example, `@kubernetes/sig-cluster-lifecycle-bugs, can you have a look at this?`

If you think you can fix the issue and you are an issue reporter or a component
organization member, assign it to yourself with just `/assign`. If you cannot
self-assign, leave a comment that you are willing to work on it and work on
creating a PR.

## Poke issue owner if PR is not created for it in 30 days

If you see any issue which is owned by a developer but a PR is not created in 30
days, a Triage engineer should contact the issue owner and ask for PR or release
ownership as needed.

## Poke SIG if a SIG label is assigned but no comment was added by SIG in 30 days

Ideally the SIG lead should have a SIG member that is a first point
of contact for SIG new issues. If an issue has a SIG label assigned and no
action is taken by SIG in 30 days (e.g. no comment was added by SIG or no
discussion was initiated) then gently poke SIG about this pending issue.
Also, consider attending one of the SIG meetings and bringing up the issue, if
you feel this is appropriate.

## Milestones

We additionally use milestones, based on minor version, for determining if a bug
should be fixed for the next release. These milestones will be especially
scrutinized as we get to the weeks just before a release. We can release a new
version of Kubernetes once they are empty. We will have two milestones per minor
release.

- **vX.Y**: The list of bugs that will be merged for that milestone once ready.

- **vX.Y-candidate**: The list of bugs that we might merge for that milestone. A
bug shouldn't be in this milestone for more than a day or two towards the end of
a milestone. It should be triaged either into vX.Y, or moved out of the release
milestones.

The above [priority](#define-priority) scheme still applies. The
`priority/critical-urgent` issues are work we feel must get done before
release.  The `priority/important-soon` and `priority/important-longterm`
issues are work we would merge into the release if it gets done, but we wouldn't
block the release on it. A few days before release, we will probably move all
`priority/important-soon` and `priority/important-longterm` bugs out of
that milestone in bulk.

More information can be found in the developer guide section for
[targeting issues and PRs to a milestone release](/contributors/devel/sig-release/release.md).

## Closing issues
Also note that, `fejta-bot` will add `lifecycle/stale` label to issues with no
activity for 90 days. Such issues will be eventually auto closed if the label is
not removed with the `/remove-lifecycle stale` label or prevented with the
`/lifecycle frozen` label. Refer to the `fejta-bot` added comments in the issue
for more details. It is fine to add any of the `triage/*` labels described in
this issue triage guidelines to issues triaged by the `fejta-bot` for a better
understanding of the issue and closing of it.

## Footnotes 
### Support requests
These should be directed to the following:
* [User documentation](https://kubernetes.io/docs/home/) and
[troubleshooting guide](https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/)

* [Stack Overflow](http://stackoverflow.com/questions/tagged/kubernetes) and
[ServerFault](http://serverfault.com/questions/tagged/kubernetes)

* [Slack](https://kubernetes.slack.com) ([registration](http://slack.k8s.io))

* [Discussion forums](https://discuss.kubernetes.io)

### User support response example

If you see support questions on kubernetes-dev@googlegroups.com or issues asking for
support try to redirect them to Stack Overflow. Example response:

```code
Please re-post your question to [Stack Overflow](http://stackoverflow.com/questions/tagged/kubernetes)
or our [Discussion Forums](https://discuss.kubernetes.io).

We are trying to consolidate the channels to which questions for help/support
are posted so that we can improve our efficiency in responding to your requests,
and to make it easier for you to find answers to frequently asked questions and
how to address common use cases.

We regularly see messages posted in multiple forums, with the full response
thread only in one place or, worse, spread across multiple forums. Also, the
large volume of support issues on GitHub is making it difficult for us to use
issues to identify real bugs.

Members of the Kubernetes community use Stack Overflow and Discussion Forums to field
support requests. Before posting a new question, please search these for answers
to similar questions, and also familiarize yourself with:

  * [user documentation](https://kubernetes.io/docs/home/)
  * [troubleshooting guide](https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/)

Again, thanks for using Kubernetes.

The Kubernetes Team
```
