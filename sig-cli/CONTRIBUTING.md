# Contributing

Welcome to the Kubernetes SIG CLI contributing guide.  We are excited
about the prospect of you joining our [community][community page]!

## Before You Begin

We strongly recommend you to understand the main [Kubernetes Contributor Guide](https://git.k8s.io/community/contributors/guide) and adhere to the contribution rules (specially signing the CLA).

You can also check the [Contributor Cheat Sheet](/contributors/guide/contributor-cheatsheet/), with common resources for existing developers.

Please be aware that all contributions to Kubernetes require time and commitment from project maintainers to direct and review work. This is done in additional to many other maintainer responsibilities, and direct engagement from maintainers is a finite resource.

### Learn a bit about the kubectl cli

This is important.

Learn about using kubectl with Kubernetes in the [Kubernetes Basics Tutorial].

Learn about managing configuration in the [kubectl docs].

## Pick your track

Determine in what capacity you are looking to contribute:

### Guided

**Who is this for?**

Contributors looking to engage with the SIG CLI community for
a sustained period of time and looking to build working relationships
with existing members.  Route to becoming a SIG CLI member as
a reviewer or approver.

**How does it work?**

Work items come from a backlog of groomed items provided by SIG CLI community members.
Each items has a stake holder willing to provide limited direction to contributors
working on it.  Contributors typically need to put in 10x the time per-issue as the
maintainers providing direction.  Contributors are expected to learn and do research
to complete the task independently with only periodic direction (~weekly).

**What is expected of contributors?**

Contributors are expected to make progress on items weekly and
provide periodic updates to any issue they are working on.
Contributors are expected exercise ownership of their code by fixing bugs
that are discovered.

### Self service

**Who is this for?**

Contributors that are looking to contribute only 1 or 2 items, or
have a specific issue they would like to like resolve and are willing
to contribute the solution.

**How does it work?**

Contributors are free to pick up any work items that they like.  Maintainers
will be focused on directing contributors working on Guided items, so contributors
picking up non-Guided items will have almost no direction or support from maintainers.

**What is expected of contributors?**

Contributions must be relatively small, simple, well documented and well tested.
Since maintainers will need to own any code for these contributions, these should
be very limited in scope and contain minimal risk
(e.g. simple regression fixes, improved documentation, improved testing).


### Modify your own `kubectl` fork

Make sure you are ready to immediately get started before you claim any piece of
work.

- Setup your [development environment][development guide].
  - This is hard.  Sorry.  We want to make this easier.
- Familiarize yourself with the code:
  - [kubernetes/cmd/kubectl] is the entry point
  - [kubernetes/staging/src/k8s.io/kubectl/pkg] is the implementation
  - Look at how some of the other commands are implemented:
    - [Almost new style (`apiresources`)](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kubectl/pkg/cmd/apiresources/apiresources.go)
    - [Latest style (`wait`)](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kubectl/pkg/cmd/wait/wait.go)
  - [Codebase Tour(old)]
  - [Codebase Tour(new)]
- Try adding a new command to do something simple:
  - Add `kubectl hello-world`: print "Hello World"
  - Add `kubectl hello-kubernetes -f file`: Print "Hello \<kind of resource\> \<name of resource\>"
  - Add `kubectl hello-kubernetes type/name`: Print "Hello \<kind of resource\> \<name of resource\> \<creation time\>"


## Your first contribution

### Adopt an issue

Pick up an [issue] from the backlog by commenting on the issue that you would like to work on it.
Make sure to mention the author of the issue and the SIG CLI [leads].

Using the following comment will make it easier for us to search for issues folks want to have
assigned to them:

`@{author} @{lead} I would like to take this`

**Note:** Don't do this unless you will start work on the issue within a few days of being assigned.

**Note:** GitHub only allows issues to be assigned to GitHub accounts that are part
of the organization.

**Picking your first issue**

In order to find your first issue, the best way is to attend the SIG CLI calls, especially the bug scrubs.
Check out the [Meet the Community](#meet-the-community) section for more information.

**Picking the right kind of issue**

Guided issues have a *kind* label defining the kind of work to be done.  Pick up an
issue that fits your experience level and interest.  Documentation and
test-coverage issues typically are smaller in scope and easier to complete than
features and cleanup issues.

- `kind/cleanup`
  - Usually some refactoring or small rewrites of code.
- `kind/documentation`
  - Write `doc.go` with package overview and examples or add code comments to document
    existing types and functions.
- `kind/feature`
  - Usually a new go package / library for some functionality that is requested.
    Should be encapsulated in its own interfaces with thorough unit tests for the new library.

**Provide periodic status updates**

Once you have requested an issue and it has been accepted, you will be expected
to provide periodic updates to it.  Do update the issue with your status at least every
week, and publish your work to a fork so the community can see your progress and
provide early feedback.

Check out the [Report Progress](#report-progress) section to know when to provide an update

If you find the issue is too challenging, time consuming, or you are no longer able to work on it,
this is perfectly acceptable and please let the issue author know.
If you like, you may pick up a different issue immediately or sometime in the future.

**Summary**:

- Don't pick up an issue until you are ready to start working on it
- When you want to pick up an issue add a comment, and mention a [lead][leads].
  Expect a response within 2 days.
- Update the issue every week with your progress so we know it is being actively worked on.
- There is an expectation that some time will be committed to working on the issue each
  week until it is completed, or you are blocked on a maintainer.

### Meet the community

Engage with the SIG CLI community!  Let us know who you are and how things are going!

- In [slack][slack-messages] (signup [here][slack-signup]) ask if there are any issues you could pick up, or let everyone know what you are working on.

- Attend a SIG CLI [meeting] and introduce yourself and what you are working on.

- The SIG CLI [community page] lists SIG CLI [leads], channels of [communication],
and group [meeting] times.

## Information about how Features are developed

Once you have made several contributions, you may want to start developing
features that you come up with.  This section is about how to propose new
features and get them accepted.

## Feature requests

__New contributors:__ Please start by adopting an [existing issue].

For feature requests please check out the [enhancement's README.md][KEP README]
Starting a discussion with SIG CLI is the best way to forward your Kubernetes Enhancement Proposal ([KEP][KEP README]).

Check out the [If your KEP issue is stuck](#if-your-kep-issue-is-stuck) section
to learn how to start a discussion about your KEP.

### Feature lifecycle

Working on a feature without getting approval for the user experience
and software design often results in wasted time and effort due to
decisions around flag names, command names, and specific command
behavior.

To minimize wasted work and improve communication across efforts,
the user experience and software design must be agreed upon before
any PRs are sent for code review.

1. Identify a problem by filing an [issue] (mention `@kubernetes/sig-cli-feature-requests`).
2. Announce the proposal as an [agenda] item for the SIG CLI [meeting].
  - Ensures awareness and feedback.
  - Should be included in meeting notes sent to the SIG CLI [group].
3. Submit a [KEP] and get it approved by a [lead][leads].
4. _Merge_ the proposal PR after approval and announcement.
5. A [lead][leads] adds the associated feature to the [feature repo], ensuring that
   - release-related decisions are properly made and communicated,
   - API changes are vetted,
   - testing is completed,
   - docs are completed,
   - feature is designated _alpha_, _beta_ or _GA_.
6. Implement the code per the [KEP].
7. Update [kubectl docs].
8. Wait for your feature to appear in the next Kubernetes release!


## Kubernetes Enhancement Proposals

__New contributors:__ Please start by adopting an [existing issue].

A Kubernetes Enhancement Proposal ([KEP]) is a design document proposing and
tracking a significant change to Kubernetes. It describes the changes
motivation, implementation, risks, testing, and path from alpha to stable.

Please follow the [KEP README] for more details on the lifecycle of KEPs.

## Implementation

Contributors can begin implementing a feature before any of the above
steps have been completed, but _should not send a PR until
the [KEP][KEP README] has been merged_.

See the [development guide] for instructions on setting up the
Kubernetes development environment.

Implementation PRs should
- mention the issue of the associated [KEP][KEP README],
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

### If your [KEP] issue is stuck

It may happen that your [KEP] gets stuck without getting merged
or additional feedback. If you believe that your design is important
and has been dropped, or it is not moving forward, please add it to
the SIG CLI bi-weekly meeting [agenda] and mail the [group] saying
you'd like to discuss it.

### General escalation instructions

See the SIG CLI [community page] for points of contact and meeting times:

- attend the SIG CLI [meeting]
- message one of the sig leads on [slack][slack-messages] (signup [here][slack-signup])
- send an email to the _sig-cli@kubernetes.io_ [group].

## Use of [@mentions]

- `@{any lead}` solicit opinion or advice from [leads].
- `@kubernetes/sig-cli-bugs` SIG CLI centric bugs.
- `@kubernetes/sig-cli-pr-reviews` triggers review of code fix PR.
- `@kubernetes/sig-cli-feature-requests` flags a feature request.
- `@kubernetes/sig-cli-proposals` flags a design proposal.

[@mentions]: https://help.github.com/articles/basic-writing-and-formatting-syntax/#mentioning-users-and-teams
[Kubernetes Basics Tutorial]: https://kubernetes.io/docs/tutorials/kubernetes-basics
[PR]: https://help.github.com/articles/creating-a-pull-request
[`PTAL`]: https://en.wiktionary.org/wiki/PTAL
[agenda]: https://docs.google.com/document/d/1I1UFGHMDO7mMbDbioQp52DEJXEhk1qymch3qL5-EN10/edit
[communication]:  /sig-cli/README.md#contact
[community page]: /sig-cli
[KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli
[KEP README]: https://github.com/kubernetes/enhancements/blob/master/README.md
[development guide]: /contributors/devel/development.md
[existing issue]: #adopt-an-issue
[feature repo]: https://github.com/kubernetes/features
[feature request]: #feature-requests
[group]: https://groups.google.com/a/kubernetes.io/g/sig-cli
[issue]: https://github.com/kubernetes/kubectl/issues?q=is%3Aissue%20state%3Aopen%20label%3Apriority%2Fbacklog
[kubectl docs]: https://kubernetes.io/docs/tutorials/object-management-kubectl/object-management/
[kubernetes/cmd/kubectl]: https://git.k8s.io/kubernetes/cmd/kubectl
[kubernetes/staging/src/k8s.io/kubectl/pkg]: https://git.k8s.io/kubernetes/staging/src/k8s.io/kubectl/pkg
[Codebase Tour(old)]:  https://youtu.be/eZeCFRh2uGg?t=538
[Codebase Tour(new)]:  https://youtu.be/Un9N9UNiO5w?t=21
[leads]: /sig-cli/README.md#leadership
[management overview]: https://kubernetes.io/docs/concepts/tools/kubectl/object-management-overview
[meeting]: /sig-cli/README.md#meetings
[release]: #release
[slack-messages]: https://kubernetes.slack.com/messages/sig-cli
[slack-signup]: http://slack.k8s.io/
