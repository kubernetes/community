---
title: "Issue Triage Guidelines"
weight: 10
slug: "issue-triage"
---

# Issue Triage: A Primer

## Table of Contents
- [Scope]
- [What is triaging?]
- [Why triage?]
- [Tools for Triage]
- [A Sample Step-by-Step]
 - [Review newly created open issues]
  - [Determine if an issue is a support request]
  - [Find the Right SIG]
    - [Set the `kind` issue to validate type (feature, documentation, bug, etc.)]
    - [Set the priority label to define urgency]
    - [Set the `sig/` label to define ownership]
- [After you label]
    * Poke issue owner if PR is not created for it in 30 days
    * Poke SIG if a SIG label is assigned but no comment was added by SIG in 30 days
* Plan milestones
* Closing issues
* Help Wanted

## Scope
These guidelines serve as a primary document for triaging incoming issues to Kubernetes. SIGs and projects are encouraged to use these guidelines as a starting point, with customization encouraged to addresss specific triaging needs.

**Note:** These guidelines only apply to the Kubernetes repository. Usage for other GitHub repositories related to Kubernetes is TBD.

## What is triaging?
Similar to what happens in a hospital emergency room, issue triage is a process by which a SIG intakes issues and requests, then applies a collection of important GitHub labels meant to categorize and organize that work based on these factors:
- urgency/priority
- the SIG or SIGs responsible for taking an action on the issue
- the kind of work 

## Why triage?
Triaging offers several benefits to SIGs:
- Speeds up issue management
- Quicker response times keeps contributors engaged
- Reduces "special requests" and context switches as issues and PRs are managed via a process
- Leads to greater transparency, broader input and more informed decision-making about priorities

Kubernetes issues are listed at https://github.com/kubernetes/kubernetes/issues. New issues start out without any labels.

## SIG Labels
The detailed list of labels can be found at
https://github.com/kubernetes/kubernetes/labels.


SIG label | example/`sig/network` | 

Following are few predetermined searches on issues for convenience:
* [Longest untriaged issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acreated-asc) (sorted by age)
* [Needs to be assigned to a SIG](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Aneeds-sig)
* [Newest incoming issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue)
* [Busy untriaged issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-desc) (sorted by number of comments)
* [Issues that need more attention](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-asc)

## Permissions and the Bot
Opening new issues, and leaving comments on other people's issues, are both possible for most people. However, permission to assign specific labels (e.g. `triaged`), change milestones or close other contributors' issues is not. For this reason, we use a bot to manage labelling and triaging. 

The bot has a set of [commands and permissions](https://go.k8s.io/bot-commands). Here, we'll cover the basic ones.

## Determine if it's a support request

Sometimes users ask for support requests in issues; these are usually requests
from people who need help configuring some aspect of Kubernetes. These issues
should be labeled with `triage/support`, directed to our support structures
(see below) and then closed. Also, if the issue is clearly abandoned or in the
wrong place, it should be closed. Keep in mind that only issue reporters,
assignees and component organization members can close issues. If you do not
have such privilege, just comment your findings. Otherwise, first `/assign`
issue to yourself and then `/close`.

### Support Structures

Support requests should be directed to the following:

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

## Find the right SIG(s)
Components are divided among [Special Interest Groups (SIGs)](/sig-list.md). Find a proper SIG for the ownership of the issue using the bot:

* Typing `/sig network` in a comment should add the sig/network label, for
example.
* Multiword SIGs use dashes, for example `/sig cluster-lifecycle`.

Keep in mind that these commands must be on their own lines, and at the front of the
comment.

## Validate if the issue is a bug

Validate if the problem is a bug by reproducing it. If reproducible, move to
the next step of defining priority. You may need to contact the issue reporter
in the following cases:
* Do a quick duplicate search to see if the issue has been reported already.
If a duplicate is found, let the issue reporter know it by marking it
duplicate. Label such issues as `triage/duplicate`.
* If you can not reproduce the issue, label it as a `triage/not-reproducible`.
Contact the issue reporter with your findings and close the issue if both the
parties agree that it could not be reproduced.
* If you need more information to further work on the issue, let the reporter
know it by adding an issue comment followed by label
`triage/needs-information`.

In all cases, if you do not get a response in 20 days then close the issue
with an appropriate comment.

## Define priority

We use GitHub issue labels for prioritization. The absence of a priority label
means the bug has not been reviewed and prioritized yet.

We try to apply these priority labels consistently across the entire project,
but if you notice an issue that you believe to be incorrectly prioritized,
please do let us know and we will evaluate your counter-proposal.

- **priority/critical-urgent**: Must be actively worked on as someone's top
priority right now. Stuff is burning. If it's not being actively worked on,
someone is expected to drop what they're doing immediately to work on it. Team
leaders are responsible for making sure that all the issues, labeled with this
priority, in their area are being actively worked on. Examples include
user-visible bugs in core features, broken builds or tests and critical
security issues.

- **priority/important-soon**: Must be staffed and worked on either currently,
or very soon, ideally in time for the next release.

- **priority/important-longterm**: Important over the long term, but may not be
currently staffed and/or may require multiple releases to complete.

- **priority/backlog**: There appears to be general agreement that this would be
good to have, but we may not have anyone available to work on it right now or in
the immediate future. Community contributions would be most welcome in the meantime 
(although it might take a while to get them reviewed if 
reviewers are fully occupied with higher priority issues, for example immediately before a release).

- **priority/awaiting-more-evidence**: Possibly useful, but not yet enough
support to actually get it done. These are mostly place-holders for potentially
good ideas, so that they don't get completely forgotten, and can be referenced
/deduped every time they come up.

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
Issues that are identified as a support request, duplicate, not-reproducible
or lacks enough information from reporter should be closed following guidelines
explained in this file. Also, any issues that can not be resolved because of
any particular reason should be closed. These issues should have one or more
of following self-readable labels:
* `triage/support`: Indicates an issues is not a bug but a support request.
* `triage/duplicate`: Indicates an issue is a duplicate of other open issue.
* `triage/not-reproducible`: Indicates an issue can not be reproduced as
described.
* `triage/needs-information`: Indicates an issue needs more information in
order to work on it.
* `triage/unresolved`: Indicates an issue that can not be resolved.

A triage engineer should add these labels appropriately. Kubernetes GitHub
Org members can search [open issues per these labels](https://github.com/kubernetes/kubernetes/labels?utf8=%E2%9C%93&q=triage%2F+kind%2Fsupport+is%3Aopen) to find ones that can be
quickly closed.

Also note that, `fejta-bot` will add `lifecycle/stale` label to issues with no
activity for 90 days. Such issues will be eventually auto closed if the label is
not removed with the `/remove-lifecycle stale` label or prevented with the
`/lifecycle frozen` label. Refer to the `fejta-bot` added comments in the issue
for more details. It is fine to add any of the `triage/*` labels described in
this issue triage guidelines to issues triaged by the `fejta-bot` for a better
understanding of the issue and closing of it.

## Help Wanted issues

We use two labels [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)
and [good first issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22)
to identify issues that have been specially groomed for new contributors.

We have specific [guidelines](/contributors/guide/help-wanted.md)
for how to use these labels. If you see an issue that satisfies these
guidelines, you can add the `help wanted` label with the `/help` command
and the `good first issue` label with the `/good-first-issue` command.
Please note that adding the `good first issue` label will also automatically
add the `help wanted` label.

If an issue has these labels but does not satisfy the guidelines, please
ask for more details to be added to the issue or remove the labels using
`/remove-help` or `/remove-good-first-issue` commands.
