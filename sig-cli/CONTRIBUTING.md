# Contributing

Welcome to the Kubernetes sig-cli contributing guide.  We are excited
about the prospect of you joining our [community][community page]!

## Before You Begin

We strongly recommend you to understand the main [Kubernetes Contributor Guide](https://git.k8s.io/community/contributors/guide) and adhere to the contribution rules (specially signing the CLA).

You can also check the [Contributor Cheat Sheet](/contributors/guide/contributor-cheatsheet/), with common resources for existing developers.

The process for contributing code to Kubernetes via SIG-CLI [community][community page].

Please be aware that all contributions to Kubernetes require time and commitment from project maintainers to direct and review work. This is done in additional to many other maintainer responsibilities, and direct engagement from maintainers is a finite resource.

### Learn a bit about the kubectl cli

This is important.

Learn about using kubectl with Kubernetes in the [Kubernetes Basics Tutorial].

Learn about managing configuration in the [kubectl docs].

## Pick your track

Determine in what capacity you are looking to contribute:

### Guided

**Who is this for?**

Contributors looking to engage with the SIG cli community for
a sustained period of time and looking to build working relationships
with existing members.  Route to becoming a SIG cli member as
a reviewer or approver.

**How does it work?**

Work items come from a backlog of groomed items provided by SIG cli community members.
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
  - Look at how some of the other commands are implemented
  - [Codebase Tour]
- Try adding a new command to do something simple:
  - Add `kubectl hello-world`: print "Hello World"
  - Add `kubectl hello-kubernetes -f file`: Print "Hello \<kind of resource\> \<name of resource\>"
  - Add `kubectl hello-kubernetes type/name`: Print "Hello \<kind of resource\> \<name of resource\> \<creation time\>"

**Note:** Consider publishing your command to a fork so a maintainer can look at it.

## Your first contribution

### Adopt an issue

Pick up an [issue] from the backlog by commenting on the issue that you would like to work on it.
Be sure to mention the author of the issue as well as the SIG cli members `@mpuckett159` and `@ardaguclu`.

Using the following comment will make it easier for us to search for issues folks want to have
assigned to them:

`cc @mpuckett159 @ardaguclu I would like to take this`

**Note:** Don't do this unless you will start work on the issue within a few days of being assigned.

**Note:** GitHub only allows issues to be assigned to GitHub accounts that are part
of the organization.

**Picking your first issue**

For your first issue, we recommend picking an issue labeled with "good first issue" from the [issue] backlog.

**Picking the right size of issue**

Be sure to pick up an issue that is appropriate to the time you are able to commit.
We recommend first time contributors start with small or medium issues.

Following are very rough estimates, but are best effort only.  They assume you have a
development environment already set up and are able to build a kubectl binary and
use it against a cluster.  These estimates assume some knowledge of Go.

- `size/S`
  - 4-10 hours
- `size/M`
  - 10-20 hours
- `size/L`
  - 20+ hours
- `size/XL`
  - 40-80 hours

Meta/Umbrella issues may have multiple components.  By signing up for a Meta/Umbrella issue,
you are only committing to one piece of it.  Let the issue author know when you have completed
some piece of it, and if you would like to continue working on it, or have it unassigned.

**Picking the right kind of issue**

Guided issues have a *type* defining the type of work to be done.  Pick up an
issue that fits your experience level and interest.  Documentation and
test-coverage issues typically are smaller in scope and easier to complete than
features and cleanup issues.

- `type/code-cleanup`
  - Usually some refactoring or small rewrites of code.
- `type/code-documentation`
  - Write `doc.go` with package overview and examples or add code comments to document
    existing types and functions.
- `type/code-feature`
  - Usually a new go package / library for some functionality that is requested.
    Should be encapsulated in its own interfaces with thorough unit tests for the new library.
- `type/code-test-coverage`
  - Audit tests for a package.  Run coverage tools and also manually look at what functions
    are missing unit or integration tests.  Write tests for these functions.

**Provide periodic status updates**

Once you have requested an issue and it has been accepted, you will be expected
to provide periodic updates to it.  Do update the issue with your status at least every
week, and publish your work to a fork so the community can see your progress and
provide early feedback.

If you find the issue is too challenging, time consuming, or you are no longer able to work on it,
this is perfectly acceptable and please let the issue author know.
If you like, you may pick up a different issue immediately or sometime in the future.

**Summary**:

- Don't pick up an issue until you are ready to start working on it
- When you want to pick up an issue, be sure to comment `@mpuckett159` and `@ardaguclu`.
  Expect a response within 2 days.
- Update the issue every week with your progress so we know it is being actively worked on.
- There is an expectation that some time will be committed to working on the issue each
  week until it is completed, or you are blocked on a maintainer.

### Meet the community

Engage with the SIG cli community!  Let us know who you are and how things are going!

- Fill out the [about me form] so we know a bit about you and can direct your work accordingly.
  - **Note:** After filling out the form, reach out via slack or the googlegroup and let us know.

- In [slack][slack-messages] (signup [here][slack-signup]),
  @mention a [lead][leads] and ask if there are any issues you could pick up, or
  let them know what you are working on.

- Attend a sig-cli [meeting] and introduce yourself and what you are working on.

- The sig-cli [community page] lists sig-cli [leads], channels of [communication],
and group [meeting] times.

## Information about how Features are developed

Once you have made several contributions, you may want to start developing
features that you come up with.  This section is about how to propose new
features and get them accepted.

## Feature requests

__New contributors:__ Please start by adopting an [existing issue].

A feature request is an [issue] mentioning `@kubernetes/sig-cli-feature-requests`.

To encourage readership, the issue description should _concisely_ (2-4 sentence) describe
the problem that the feature addresses.

### Feature lifecycle

Working on a feature without getting approval for the user experience
and software design often results in wasted time and effort due to
decisions around flag names, command names, and specific command
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
7. Update [kubectl docs].
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

The proposal can be merged into the [design repo] after [leads][leads]
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
- send an email to the _sig-cli@kubernetes.io_ [group].

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
[communication]:  /sig-cli/README.md#contact
[community page]: /sig-cli
[design proposal]: #design-proposals
[design repo]: https://git.k8s.io/design-proposals-archive/cli
[design template]: https://git.k8s.io/design-proposals-archive/Design_Proposal_TEMPLATE.md
[development guide]: /contributors/devel/development.md
[existing issue]: #adopt-an-issue
[feature repo]: https://github.com/kubernetes/features
[feature request]: #feature-requests
[feature]: https://github.com/kubernetes/features
[group]: https://groups.google.com/a/kubernetes.io/g/sig-cli
[issue]: https://github.com/kubernetes/kubectl/issues?q=is%3Aissue%20state%3Aopen%20label%3Apriority%2Fbacklog
[kubectl docs]: https://kubernetes.io/docs/tutorials/object-management-kubectl/object-management/
[kubernetes/cmd/kubectl]: https://git.k8s.io/kubernetes/cmd/kubectl
[kubernetes/staging/src/k8s.io/kubectl/pkg]: https://git.k8s.io/kubernetes/staging/src/k8s.io/kubectl/pkg
[Codebase Tour]:  https://youtu.be/eZeCFRh2uGg?t=538
[leads]: /sig-cli/README.md#leadership
[management overview]: https://kubernetes.io/docs/concepts/tools/kubectl/object-management-overview
[meeting]: /sig-cli/README.md#meetings
[release]: #release
[slack-messages]: https://kubernetes.slack.com/messages/sig-cli
[slack-signup]: http://slack.k8s.io/
[tests]: /contributors/devel/sig-testing/testing.md
[about me form]: https://docs.google.com/forms/d/1ID6DX1abiDr9Z9_sXXC0DsMwuyHb_NeFdB3xeRa4Vf0
