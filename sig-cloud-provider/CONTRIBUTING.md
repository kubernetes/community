# Contributing

Welcome to the Kubernetes SIG Cloud Provider contributing guide.  We are excited
about the prospect of you joining our community!

## Before You Begin

We strongly recommend you to understand the main [Kubernetes Contributor
Guide](http://git.k8s.io/community/contributors/guide) and adhere to the
contribution rules (specially signing the CLA).

You can also check the [Contributor Cheat
Sheet](/contributors/guide/contributor-cheatsheet/), with common resources for
existing developers.

Read the [developer guide].

Please be aware that all contributions to Kubernetes require time and
commitment from project maintainers to direct and review work. This is done in
additional to many other maintainer responsibilities, and direct engagement
from maintainers is a finite resource.

### Learn about our work

* [The Future of Cloud Providers in Kubernetes](https://kubernetes.io/blog/2019/04/17/the-future-of-cloud-providers-in-kubernetes/)
* [Cloud Controller Manager](https://kubernetes.io/docs/concepts/architecture/cloud-controller/)

## Your first contribution

### Adopt an issue

Pick up an [issue] from the backlog by commenting on the issue that you would
like to work on it.  Be sure to mention the author of the issue as well as the
SIG Cloud Provider members.

**Note:** Don't do this unless you will start work on the issue within a few
days of being assigned.

**Note:** GitHub only allows issues to be assigned to GitHub accounts that are
part of the organization.

**Picking your first issue**

For your first issue, we recommend picking an issue labeled with "good first
issue" from the [issue] backlog.  Work with active members of the SIG to find a
suitable issue if you need help.

**Picking the right size of issue**

Be sure to pick up an issue that is appropriate to the time you are able to
commit.  We recommend first time contributors start with small or medium
issues.

Following are very rough estimates, but are best effort only.  They assume you
have a development environment already set up and are able to build a kubectl
binary and use it against a cluster.  These estimates assume some knowledge of
Go.

- `size/S`
  - simple complexity, good for novices to project (4-10 hours)
- `size/M`
  - moderate complexity (10-20 hours)
- `size/L`
  - high complexity (20+ hours)
- `size/XL`
  - very high complexity, might require help from community members (40-80 hours)

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

**Testing**

Look at [tests] for more information about testing.

**Summary**:

- Don't pick up an issue until you are ready to start working on it
- When you want to pick up an issue, be sure to comment to the [leads] that you
  are taking the issue.
- Update the issue every week with your progress so we know it is being
  actively worked on.
- There is an expectation that some time will be committed to working on the
  issue each week until it is completed, or you are blocked on a maintainer.

### Meet the community

Engage with the SIG cloud provider community!  Let us know who you are and how
things are going!

- In [slack][slack-messages] (signup [here][slack-signup]), @mention a
  [lead][leads] and ask if there are any issues you could pick up, or let them
  know what you are working on.

- Attend a sig-cloud-provider [meeting] and introduce yourself and what you are
  working on.

- The sig-cloud-provider [community page] lists sig-cloud-provider [leads],
  channels of [communication], and group [meeting] times.

## Information about how Features are developed

Kubernetes uses a process called a KEP (Kubernetes enhancement proposal) to
drive feature development.  See [enhancements] for the most up to date
information about how enhancements are planned and developed in the Kubernetes
community.

## Escalation

### If your bug issue is stuck

If an issue isn't getting any attention and is unresolved, mention
`@kubernetes/sig-cloud-provider-bugs`.

Highlight the severity and urgency of the issue.  For severe issues
escalate by contacting sig [leads] and attending the [meeting].

### If your feature request issue is stuck

If an issue isn't getting any attention and is unresolved, mention
`@kubernetes/sig-cloud-provider-feature-requests`.

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
  comments, mention `@kubernetes/sig-cloud-provider-pr-reviews` and ask for attention.
- If your PR is stuck for a week or more _after_ it got comments, but
  the attention has died down.  Mention the reviewer and comment with
  [`PTAL`].

If you are still not able to get any attention after a couple days,
escalate to sig [leads] by mentioning them.

### If your KEP is stuck

It may happen that your KEP gets stuck without getting merged or additional
feedback. If you believe that your design is important and has been dropped, or
it is not moving forward, please add it to the sig-cloud-provider bi-weekly
meeting [agenda] and mail the [group] saying you'd like to discuss it.

### General escalation instructions

See the sig-cloud-provider [community page] for points of contact and meeting times:

- Attend the sig-cloud-provider [meeting]
- Message one of the sig leads on [slack][slack-messages] (signup [here][slack-signup])
- Send an email to the _kubernetes-cloud-provider@googlegroups.com_ [group].

## Use of [@mentions]

- `@{any lead}` solicit opinion or advice from [leads].
- `@kubernetes/sig-cloud-provider-bugs` sig-cloud-provider centric bugs.
- `@kubernetes/sig-cloud-provider-pr-reviews` triggers review of code fix PR.
- `@kubernetes/sig-cloud-provider-feature-requests` flags a feature request.
- `@kubernetes/sig-cloud-provider-proposals` flags a design proposal.

[@mentions]: https://help.github.com/articles/basic-writing-and-formatting-syntax/#mentioning-users-and-teams
[Kubernetes Basics Tutorial]: https://kubernetes.io/docs/tutorials/kubernetes-basics
[PR]: https://help.github.com/articles/creating-a-pull-request
[`PTAL`]: https://en.wiktionary.org/wiki/PTAL
[agenda]: https://docs.google.com/document/d/1OZE-ub-v6B8y-GuaWejL-vU_f9jsjBbrim4LtTfxssw/edit#
[communication]:  /sig-cloud-provider/README.md#contact
[community page]: /sig-cloud-provider
[developer guide]: /contributors/devel/development.md
[enhancements]: https://github.com/kubernetes/enhancements
[group]: https://groups.google.com/forum/#!forum/kubernetes-sig-cloud-provider
[issue]: https://github.com/kubernetes/cloud-provider/issues
[leads]: /sig-cloud-provider/README.md#leadership
[meeting]: /sig-cloud-provider/README.md#meetings
[slack-messages]: https://kubernetes.slack.com/messages/sig-cloud-provider
[slack-signup]: http://slack.k8s.io/
[tests]: /contributors/devel/sig-testing/testing.md
