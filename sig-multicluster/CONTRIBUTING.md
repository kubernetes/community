# Contributing

The process for contributing code to Kubernetes via the sig-multicluster [community][community page].

## TL;DR

- The sig-multicluster [community page] lists sig-multicluster [leads]
  and group [meeting] times.
- Request a feature by making an [issue] and mentioning
  `@kubernetes/sig-multicluster-feature-requests`.
- Write a design proposal before starting work on a new feature.
- Write [tests]!

## Before You Begin

Welcome to the Kubernetes sig-multicluster contributing guide.  We are excited
about the prospect of you joining our [community][community page]!

Please understand that all contributions to Kubernetes require time
and commitment from the project maintainers to review the ux, software
design, and code.  Mentoring and on-boarding new contributors is done
in addition to many other responsibilities.

If you are interested in contributing to Kubernetes as a whole there is
a top level [contributor's guide][contributors guide] 

### Understand the big picture

- Complete the [Kubernetes Basics Tutorial].

### Agree to contribution rules

Follow the [CLA signup instructions](../CLA.md).

### Adopt an issue

New contributors can try the following to work on an existing bug or approved design:

- In [slack][slack-messages] (signup [here][slack-signup]),
  @mention a [lead][leads] and ask if there are any issues you could pick up.
  We also maintain a list of [multi cluster issues where help is wanted][multicluster_help_wanted_issues].
  Most of them are not very complex, so that's probably a good starting point.
  Leads can recommend issues that have enough priority to receive PR review bandwidth.
- Send an email to the _kubernetes-sig-multicluster@googlegroups.com_ [group]

  > Subject: New sig-multicluster contributor _${yourName}_
  >
  > Body: Hello, my name is _${yourName}_.  I would like to get involved in
  > contributing to the Kubernetes project.  I have read all of the
  > user documentation listed on the community contributing page.
  > What should I do next to get started?

- Attend a sig-multicluster [meeting] and introduce yourself as looking to get started.

### Bug lifecycle

#### Filing a bug

1. An [issue] is filed that
  - includes steps to reproduce the issue including client / server version,
  - mentions `@kubernetes/sig-multicluster-bugs`.

#### Sending a fix

2. A [PR] fixing the issue is implemented that
  - __includes unit and test-cmd tests__,
  - incorporates review feedback,
  - description includes `Closes #<Issue Number>` or `Fixes #<Issue Number>`,
  - description or comment @mentions `@kubernetes/sig-multicluster-pr-reviews`.
3. Fix appears in the next Kubernetes release!

## Feature requests

__New contributors:__ Please start by adopting an [existing issue].

A feature request is an [issue] mentioning `@kubernetes/sig-multicluster-feature-requests`.

To encourage readership, the issue description should _concisely_ (2-4 sentence) describe
the problem that the feature addresses.

### Feature lifecycle

Working on a feature without getting approval for the user experience
and software design often results in wasted time and effort due to
decisions around names and user experience.

To minimize wasted work and improve communication across efforts,
the user experience and software design must be agreed upon before
any PRs are sent for code review.

1. Identify a problem by filing an [issue] (mention `@kubernetes/sig-multicluster-feature-requests`).
2. Submit a [design proposal] and get it approved by a lead.
3. Announce the proposal as an [agenda] item for the sig-multicluster [meeting].
  - Ensures awareness and feedback.
  - Should be included in meeting notes sent to the sig-multicluster [group].
4. _Merge_ the proposal PR after approval and announcement.
5. A [lead][leads] adds the associated feature to the [feature repo], ensuring that
   - release-related decisions are properly made and communicated,
   - API changes are vetted,
   - testing is completed,
   - docs are completed,
   - feature is designated _alpha_, _beta_ or _GA_.
6. Implement the code per discussion in [bug lifecycle][bug].
7. Update docs.
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
- Mention `@kubernetes/sig-multicluster-proposals` in the description.
- Mention the related [feature request].

Expect feedback from 2-3 different sig-multicluster community members.

Incorporate feedback and comment [`PTAL`].

Once a [lead][leads] has agreed (via review commentary) that design
and code review resources can be allocated to tackle the proposal, the
details of the user experience and design should be discussed in the
community.

This step is _important_; it prevents code churn and thrashing around
issues like flag names, command names, etc.

It is normal for sig-multicluster community members to push back on feature
proposals. sig-multicluster development and review resources are extremely
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
- mention `@kubernetes/sig-multicluster-pr-reviews`,
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
`@kubernetes/sig-multicluster-bugs`.

Highlight the severity and urgency of the issue.  For severe issues
escalate by contacting sig [leads] and attending the [meeting].

### If your feature request issue is stuck

If an issue isn't getting any attention and is unresolved, mention
`@kubernetes/sig-multicluster-feature-requests`.

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
  comments, mention `@kubernetes/sig-multicluster-pr-reviews` and ask for attention.
- If your PR is stuck for a week or more _after_ it got comments, but
  the attention has died down.  Mention the reviewer and comment with
  [`PTAL`].

If you are still not able to get any attention after a couple days,
escalate to sig [leads] by mentioning them.

### If your design proposal issue is stuck

It may happen that your design doc gets stuck without getting merged
or additional feedback. If you believe that your design is important
and has been dropped, or it is not moving forward, please add it to
the sig multicluster bi-weekly meeting [agenda] and mail the [group] saying
you'd like to discuss it.

### General escalation instructions

See the sig-multicluster [community page] for points of contact and meeting times:

- attend the sig-multicluster [meeting]
- message one of the sig leads on [slack][slack-messages] (signup [here][slack-signup])
- send an email to the _kubernetes-sig-multicluster@googlegroups.com_ [group].

## Use of [@mentions]

- `@{any lead}` solicit opinion or advice from [leads].
- `@kubernetes/sig-multicluster-bugs` sig-multicluster centric bugs.
- `@kubernetes/sig-multicluster-pr-reviews` triggers review of code fix PR.
- `@kubernetes/sig-multicluster-feature-requests` flags a feature request.
- `@kubernetes/sig-multicluster-proposals` flags a design proposal.

[@mentions]: https://help.github.com/articles/basic-writing-and-formatting-syntax/#mentioning-users-and-teams
[Kubernetes Basics Tutorial]: https://kubernetes.io/docs/tutorials/kubernetes-basics
[PR]: https://help.github.com/articles/creating-a-pull-request
[`PTAL`]: https://en.wiktionary.org/wiki/PTAL
[agenda]: https://docs.google.com/document/d/18mk62nOXE_MCSSnb4yJD_8UadtzJrYyJxFwbrgabHe8/edit
[bug]: #bug-lifecycle
[community page]: /sig-multicluster
[contributors guide]: /contributors/guide
[design proposal]: #design-proposals
[design repo]: /contributors/design-proposals/multicluster
[design template]: /contributors/design-proposals/Design_Proposal_TEMPLATE.md
[development guide]: /contributors/devel/development.md
[existing issue]: #adopt-an-issue
[feature repo]: https://github.com/kubernetes/features
[feature request]: #feature-requests
[feature]: https://github.com/kubernetes/features
[group]: https://groups.google.com/forum/#!forum/kubernetes-sig-multicluster
[issue]: https://github.com/kubernetes/kubernetes/issues
[multicluster_help_wanted_issues]: https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3A"help+wanted"+label%3Asig%2Fmulticluster
[kubectl concept docs]: https://git.k8s.io/kubernetes.github.io/docs/concepts/tools/kubectl
[kubectl docs]: https://kubernetes.io/docs/user-guide/kubectl-overview
[kubernetes/cmd/kubectl]: https://git.k8s.io/kubernetes/cmd/kubectl
[kubernetes/pkg/kubectl]: https://git.k8s.io/kubernetes/pkg/kubectl
[leads]: /sig-multicluster#leads
[management overview]: https://kubernetes.io/docs/concepts/tools/kubectl/object-management-overview
[meeting]: /sig-multicluster#meetings
[release]: #release
[slack-messages]: https://kubernetes.slack.com/messages/sig-multicluster
[slack-signup]: http://slack.k8s.io/
[tests]: /contributors/devel/sig-testing/testing.md
