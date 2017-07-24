# Kubernetes Issue Triage Guidelines

## Purpose

Speed up issue management.

The Kubernetes issues are listed at https://github.com/kubernetes/kubernetes/issues
and are identified with labels. For example, an issue that belongs to SIG
Network group will eventually be set to label `sig/network`. New issues will
start out without any labels. The detailed list of labels can be found at
https://github.com/kubernetes/kubernetes/labels. While working on triaging
issues you may not have privilege to assign specific label (e.g. `triaged`)
and in that case simply add a comment in the issue with your findings.

Following are few predetermined searches on issues for convenience:
* [Longest untriaged issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acreated-asc) (sorted by age)
* [Needs to be assigned to a SIG](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acreated-asc)
* [Newest incoming issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue)
* [Busy untriaged issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-desc) (sorted by number of comments)
* [Issues that need more attention](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+sort%3Acomments-asc)

## Scope

These guidelines serves as a primary document for triaging incoming issues to
Kubernetes. SIGs and projects are encouraged to either use these guidelines, or
use this as a starting point if necessary. For example if your SIG has specific
triaging needs, extend these guidelines.
**Note:** These guidelines only applies to the kubernetes repository. Its usage
for other github repositories related to Kubernetes is TBD.

## Using the bot

Most people can leave comments and open issues. They don't have the ability to
set labels, change milestones and close other peoples issues. For that we use
a bot to manage labelling and triaging. The bot has a set of
[commands and permissions](https://github.com/kubernetes/test-infra/blob/master/commands.md)
and this document will cover the basic ones.

## Determine if it’s a support request

Sometimes users ask for support requests in issues; these are usually requests
from people who need help configuring some aspect of Kubernetes. These should be
directed to our [support structures](https://github.com/kubernetes/community/blob/master/contributors/devel/on-call-user-support.md) and then closed. Also, if the issue is clearly abandoned or in
the wrong place, it should be closed. Keep in mind that only issue reporter,
assignees and component organization members can close issue. If you do not
have such privilege, just comment your findings. Otherwise, first `/assign`
issue to yourself and then `/close`.

## Find the right SIG(s)
Components are divided among [Special Interest Groups (SIGs)](https://github.com/kubernetes/community/blob/master/sig-list.md). Find a proper SIG for the ownership of the issue using the bot:

* Typing `/sig network` in a comment should add the sig/network label, for
example.
* Multiword SIGs use dashes, for example `/sig cluster-lifecycle`.

Keep in mind that these commands must be on its own and at the front of the
comment.

## Validate if the issue is bug

Validate if the problem is a bug by reproducing it. If reproducible, move to the
next step of defining priority. You may need to contact the issue reporter in
the following cases:
* Do a quick duplicate search to see if the issue has been reported already. If
a duplicate is found, let the issue reporter know it by marking it duplicate.
* If you can not reproduce the issue, contact the issue reporter with your
findings and close the issue if both the parties agree that it could not be
reproduced.
* If the issue is non-trivial to reproduce, work with issue reporter and let SIG
know of your findings.
* If you do not get a response in 20 days, then close the issue with appropriate
comment.

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

- **priority/failing-test**: Automatically filed frequently failing test. Needs
to be investigated.

- **priority/important-soon**: Must be staffed and worked on either currently,
or very soon, ideally in time for the next release.

- **priority/important-longterm**: Important over the long term, but may not be
currently staffed and/or may require multiple releases to complete.

- **priority/backlog**: There appears to be general agreement that this would be
good to have, but we may not have anyone available to work on it right now or in
the immediate future. Community contributions would be most welcome in the mean
time (although it might take a while to get them reviewed if reviewers are fully
occupied with higher priority issues, for example immediately before a release).

- **priority/awaiting-more-evidence**: Possibly useful, but not yet enough
support to actually get it done. These are mostly place-holders for potentially
good ideas, so that they don't get completely forgotten, and can be referenced
/deduped every time they come up.

## Set ownership

If you are not sure of who should own issue, defer to the
SIG label only. If you feel the issue should warrant a notification,you can ping
a team with an @ mention, in this format, `@kubernetes/sig-<group-name>-<group-suffix>`.
Here the `<group-suffix>` can be one of `bugs, feature-requests, pr-reviews, test-failures, proposals`.
For example, `@kubernetes/sig-cluster-lifecycle-bugs, can you have a look at this?`

If you think you can fix the issue and you are an issue reporter or a component
organization member, assign it to yourself with just `/assign`. If you can not
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
Also consider attending one of the SIG meetings and brig up issue, if you feel
this is appropriate.

## Milestones

We additionally use milestones, based on minor version, for determining if a bug
should be fixed for the next release. These milestones will be especially
scrutinized as we get to the weeks just before a release. We can release a new
version of Kubernetes once they are empty. We will have two milestones per minor
release.

- **vX.Y**: The list of bugs that will be merged for that milestone once ready.

- **vX.Y-candidate**: The list of bug that we might merge for that milestone. A
bug shouldn't be in this milestone for more than a day or two towards the end of
a milestone. It should be triaged either into vX.Y, or moved out of the release
milestones.

The above [priority](#define-priority) scheme still applies. The `priority/critical-urgent`
and `priority/failing-test` issues are work we feel must get done before
release.  The `priority/important-soon` and `priority/important-longterm`
issues are work we would merge into the release if it gets done, but we wouldn't
block the release on it. A few days before release, we will probably move all
`priority/important-soon` and `priority/important-longterm` bugs out of
that milestone in bulk.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/issues.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
