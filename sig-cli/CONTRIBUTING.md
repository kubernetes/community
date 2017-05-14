# Contributing

The process for contributing code to Kubernetes via the sig-cli [community][community page].

## TL;DR

- The sig-cli [community page] lists sig-cli [leads],
  channels of [communication], and group [meeting] times.
- New contributors: please start by adopting an [existing issue].
- Request a feature by making an [issue] and mentioning
  `@kubernetes/sig-cli-feature-requests`.
- Write a [design proposal] before starting work on a new feature.
- Write [tests]!

## Before You Begin

Welcome to the Kubernetes sig-cli contributing guide.  We are excited
about the prospect of you joining our [community][community page]!

Please understand that all contributions to Kubernetes require time
and commitment from the project maintainers to review the ux, software
design, and code.  Mentoring and on-boarding new contributors is done
in addition to many other responsibilities.

### Understand the big picture

- Complete the [Kubernetes Basics Tutorial].
- Be familiar with [kubectl user facing documentation ][kubectl docs].
- Read the concept guides starting with the [management overview].

### Modify your own `kubectl` fork

Make sure you are ready to immediately get started once you have been
assigned a piece of work.  Do this right away.

- Setup your [development environment][development guide].
- Look at code:
  - [kubernetes/cmd/kubectl] is the entry point
  - [kubernetes/pkg/kubectl] is the implementation
  - Look at how some of the other commands are implemented
- Add a new command to do something simple:
  - Add `kubectl hello-world`: print "Hello World"
  - Add `kubectl hello-kubernetes -f file`: Print "Hello \<kind of resource\> \<name of resource\>"
  - Add `kubectl hello-kubernetes type/name`: Print "Hello \<kind of resource\> \<name of resource\> \<creation time\>"

### Agree to contribution rules

Follow the [CLA signup instructions](../CLA.md).

### Adopt an issue

New contributors can try the following to work on an existing [bug] or [approved design][design repo]:

- In [slack][slack-messages] (signup [here][slack-signup]),
  @mention a [lead][leads] and ask if there are any issues you could pick up.
  Leads can recommend issues that have enough priority to receive PR review bandwidth.
  We also maintain a list of [CLI issues where help is wanted][cli_help_wanted_issues].
  Most of them are not very complex, so that's probably a good starting point.
- Send an email to the _kubernetes-sig-cli@googlegroups.com_ [group]

  > Subject: New sig-cli contributor _${yourName}_
  >
  > Body: Hello, my name is _${yourName}_.  I would like to get involved in
  > contributing to the Kubernetes project.  I have read all of the
  > user documentation listed on the community contributing page.
  > What should I do next to get started?

- Attend a sig-cli [meeting] and introduce yourself as looking to get started.

### Bug lifecycle

1. An [issue] is filed that
  - includes steps to reproduce the issue including client / server version,
  - mentions `@kubernetes/sig-cli-bugs`.
2. A [PR] fixing the issue is implemented that
  - __includes unit and test-cmd tests__,
  - incorporates review feedback,
  - description includes `Closes #<Issue Number>` or `Fixes #<Issue Number>`,
  - description or comment @mentions `@kubernetes/sig-cli-pr-reviews`.
3. Fix appears in the next Kubernetes release!

## Feature requests

__New contributors:__ Please start by adopting an [existing issue].

A feature request is an [issue] mentioning `@kubernetes/sig-cli-feature-requests`.

To encourage readership, the issue description should _concisely_ (2-4 sentence) describe
the problem that the feature addresses.

### Feature lifecycle

Working on a feature without getting approval for the user experience
and software design often results in wasted time and effort due to
decisions around flag-names, command names, and specific command
behavior.

To minimize wasted work and improve communication across efforts,
the user experience and software design must be agreed upon before
any PRs are sent for code review.

1. Identify a problem by filing an [issue] (mention `@kubernetes/sig-cli-feature-requests`).
2. Submit a [design proposal] and get it approved by a lead.
3. Announce the proposal as an [agenda] item for the sig-cli [meeting].
  - Ensures awareness and feedback.
  - Should be included in meeting notes sent to the sig-cli [group].
4. _Merge_ the proposal PR after approval and announcement.
5. A [lead][leads] adds the associated feature to the [feature repo], ensuring that
   - release-related decisions are properly made and communicated,
   - API changes are vetted,
   - testing is completed,
   - docs are completed,
   - feature is designated _alpha_, _beta_ or _GA_.
6. Implement the code per discussion in [bug lifecycle][bug].
7. Update [kubectl concept docs].
8. Wait for your feature to appear in the next Kubernetes release!


## Design Proposals

__New contributors:__ Please start by adopting an [existing issue].

A design proposal is a single markdown document in the [design repo]
that follows the [design template].

To make one,
- Prepare the markdown document as a PR to that repo.
  - Avoid _Work In Progress_ (WIP) PRs (send it only after
    you consider it complete).
  - For early feedback, use the email discussion [group].
- Mention `@kubernetes/sig-cli-proposals` in the description.
- Mention the related [feature request].

Expect feedback from 2-3 different sig-cli community members.

Incorporate feedback and comment [`PTAL`].

Once a [lead][leads] has agreed (via review commentary) that design
and code review resources can be allocated to tackle the proposal, the
details of the user experience and design should be discussed in the
community.

This step is _important_; it prevents code churn and thrashing around
issues like flag names, command names, etc.

It is normal for sig-cli community members to push back on feature
proposals. sig-cli development and review resources are extremely
constrained. Community members are free to say

- No, not this release (or year).
- This is desirable but we need help on these other existing issues before tackling this.
- No, this problem should be solved in another way.

The proposal can be merged into the [design repo] after [lead][leads]
approval and discussion as a meeting [agenda] item.

Then coding can begin.

## Implementation

Contributors can begin implementing a feature before any of the above
steps have been completed, but _should not send a PR until
the [design proposal] has been merged_.

See the [development guide] for instructions on setting up the
Kubernetes development environment.

Implementation PRs should
- mention the issue of the associated design proposal,
- mention `@kubernetes/sig-cli-pr-reviews`,
- __include tests__.

Small features and flag changes require only unit/integration tests,
while larger changes require both unit/integration tests and e2e tests.

### Report progress

_Leads need your help to ensure that progress is made to
get the feature into a [release]._

While working on the issue, leave a weekly update on the issue
including:

1. What's finished?
2. What's part is being worked on now?
3. Anything blocking?


## Documentation

_Let users know about cool new features by updating user facing documentation._

Depending on the contributor and size of the feature, this
may be done either by the same contributor that implemented the feature,
or another contributor who is more familiar with the existing docs
templates.

## Release

Several weeks before a Kubernetes release, development enters a stabilization
period where no new features are merged.  For a feature to be accepted
into a release, it must be fully merged and tested by this time.  If
your feature is not fully complete, _including tests_, it will have
to wait until the next release.

## Merge state meanings

- Merged:
  - Ready to be implemented.
- Unmerged:
  - Experience and design still being worked out.
  - Not a high priority issue but may implement in the future: revisit
    in 6 months.
  - Unintentionally dropped.
- Closed:
  - Not something we plan to implement in the proposed manner.
  - Not something we plan to revisit in the next 12 months.

## Escalation

### If your bug issue is stuck

If an issue isn't getting any attention and is unresolved, mention
`@kubernetes/sig-cli-bugs`.

Highlight the severity and urgency of the issue.  For severe issues
escalate by contacting sig [leads] and attending the [meeting].

### If your feature request issue is stuck

If an issue isn't getting any attention and is unresolved, mention
`@kubernetes/sig-cli-feature-requests`.

If a particular issue has a high impact for you or your business,
make sure this is clear on the bug, and reach out to the sig leads
directly.  Consider attending the sig meeting to discuss over video
conference.

### If your PR is stuck

It may happen that your PR seems to be stuck without clear actionable
feedback for a week or longer.  A PR _associated with a bug or design
proposal_ is much less likely to be stuck than a dangling PR.

However, if it happens do the following:

- If your PR is stuck for a week or more because it has never gotten any
  comments, mention `@kubernetes/sig-cli-pr-reviews` and ask for attention.
- If your PR is stuck for a week or more _after_ it got comments, but
  the attention has died down.  Mention the reviewer and comment with
  [`PTAL`].

If you are still not able to get any attention after a couple days,
escalate to sig [leads] by mentioning them.

### If your design proposal issue is stuck

It may happen that your design doc gets stuck without getting merged
or additional feedback. If you believe that your design is important
and has been dropped, or it is not moving forward, please add it to
the sig cli bi-weekly meeting [agenda] and mail the [group] saying
you'd like to discuss it.

### General escalation instructions

See the sig-cli [community page] for points of contact and meeting times:

- attend the sig-cli [meeting]
- message one of the sig leads on [slack][slack-messages] (signup [here][slack-signup])
- send an email to the _kubernetes-sig-cli@googlegroups.com_ [group].

## Use of [@mentions]

- `@{any lead}` solicit opinion or advice from [leads].
- `@kubernetes/sig-cli-bugs` sig-cli centric bugs.
- `@kubernetes/sig-cli-pr-reviews` triggers review of code fix PR.
- `@kubernetes/sig-cli-feature-requests` flags a feature request.
- `@kubernetes/sig-cli-proposals` flags a design proposal.

[@mentions]: https://help.github.com/articles/basic-writing-and-formatting-syntax/#mentioning-users-and-teams
[Kubernetes Basics Tutorial]: https://kubernetes.io/docs/tutorials/kubernetes-basics
[PR]: https://help.github.com/articles/creating-a-pull-request
[`PTAL`]: https://en.wiktionary.org/wiki/PTAL
[agenda]: https://docs.google.com/document/d/1r0YElcXt6G5mOWxwZiXgGu_X6he3F--wKwg-9UBc29I/edit
[bug]: #bug-lifecycle
[communication]:  https://github.com/kubernetes/community/tree/master/sig-cli#communication
[community page]: https://github.com/kubernetes/community/tree/master/sig-cli
[design proposal]: #design-proposals
[design repo]: https://github.com/kubernetes/community/tree/master/contributors/design-proposals/sig-cli
[design template]: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/sig-cli/template.md
[development guide]: https://github.com/kubernetes/community/blob/master/contributors/devel/development.md
[existing issue]: #adopt-an-issue
[feature repo]: https://github.com/kubernetes/features
[feature request]: #feature-requests
[feature]: https://github.com/kubernetes/features
[group]: https://groups.google.com/forum/#!forum/kubernetes-sig-cli
[issue]: https://github.com/kubernetes/kubernetes/issues
[cli_help_wanted_issues]: https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fcli+label%3Ahelp-wanted
[kubectl concept docs]: https://github.com/kubernetes/kubernetes.github.io/tree/master/docs/concepts/tools/kubectl
[kubectl docs]: https://kubernetes.io/docs/user-guide/kubectl-overview
[kubernetes/cmd/kubectl]: https://github.com/kubernetes/kubernetes/tree/master/cmd/kubectl
[kubernetes/pkg/kubectl]: https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl
[leads]: https://github.com/kubernetes/community/tree/master/sig-cli#leads
[management overview]: https://kubernetes.io/docs/concepts/tools/kubectl/object-management-overview
[meeting]: https://github.com/kubernetes/community/tree/master/sig-cli#meetings
[release]: #release
[slack-messages]: https://kubernetes.slack.com/messages/sig-cli
[slack-signup]: http://slack.k8s.io/
[tests]: https://github.com/kubernetes/community/blob/master/contributors/devel/testing.md
