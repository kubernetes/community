# Contributing to sig-Cli

This document explains the process for contributing code to the Kubernetes
repo through the sig-cli community.

## TL;DR

- New contributors should reach out on slack sig-cli - [link](https://github.com/kubernetes/community/tree/master/sig-cli) for an issue to work on that has an approved design
- New contributors should attend the next sig-cli meeting and introduce themselves
- Write tests

## Important

The [sig-cli community page](https://github.com/kubernetes/community/tree/master/sig-cli)
contains the GitHub handles for the sig-cli leads, slack channels,
email discussion group, and bi-weekly meeting time.

## Message to new contributors

Welcome to the Kubernetes sig-cli contributing guide.  We are excited
about the prospect of you joining our community.  Please understand
that all contributions to Kubernetes require time and commitment from
the project maintainers to review the ux, software design, and code.  We
recommend that you reach out to one of the sig leads
and ask the best way to get started.  It is likely that there a couple
issues perfect for a new contributor to get started on.

sig leads can be reached in the sig-cli slack channel or at the sig-cli
bi-weekly meeting.  See the [sig-cli community page](https://github.com/kubernetes/community/tree/master/sig-cli)
for a link to the slack channel, list of sig leads, and the community meeting time and location.

Example message for asking how to become a contributor:

    Hello, my name is <your name>.  I would like to get involved in
    contributing to the Kubernetes project.  What is the best way to
    get started?

Please understand that mentoring and on boarding new contributors takes
a lot time and energy from maintainers who have many other
responsibilities.

## Before you begin

We are so glad you are ready to get started.  Please complete the following steps
to join the community.

- Read the Kubectl user facing documentation to make sure you understand the tool
  - Complete the [Kubernetes Basics Tutorial](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
  - Read the concept guides starting with the [Overview](https://kubernetes.io/docs/concepts/tools/kubectl/object-management-overview/)
    - This is really important so that you have a good understanding of the tool before you start trying to work on it
- Visit the [sig-cli community page](https://github.com/kubernetes/community/tree/master/sig-cli)
  - Join the `kubernetes-sig-cli@googlegroups.com` so you get email updates
  - Join the Kubernetes slack channel `sig-cli`
- Introduce yourself to the community by sending an email to kubernetes-sig-cli@googlegroups.com and let us know you want to be a contributor

Completing the following steps will make sure you are ready to immediately
get started once you have been assigned a piece of work.  Do these right
away.

- Setup your development environment so you can build and run Kubernetes. See [this guide for details](https://github.com/kubernetes/community/blob/master/contributors/devel/development.md)
- Starting taking a look at the code:
  - `kubernetes/cmd/kubectl`: This is the entry point
  - `kubernetes/pkg/kubectl`: This contains the implementation
  - Look at how some of the other commands are implemented
- Try adding a new command to do something simple:
  - Add `kubectl hello-world`: print "Hello World"
  - Add `kubectl hello-kubernetes -f file`: Print "Hello <kind of resource> <name of resource>"
  - Add `kubectl hello-kubernetes type/name`: Print "Hello <kind of resource> <name of resource> <creation time>"

## Finding an existing issue to work on

The preferred way to own a piece of work is to directly reach out
to one of the sig-cli leads.  This will ensure that the issue is
high-priority enough that it will receive PR review bandwidth from
the sig-cli community.

1. In the Kubernetes sig-cli slack channel, mention @pwittrock, @adohe, or  @fabianofranz and ask if there are any issues you could pick up
2. Send an email to the `kubernetes-sig-cli@googlegroups.com`

        Subject: `New Sig-Cli Contributor <Your Name>`
        Body: Hello, my name is <your name>.  I would like to get involved in
        contributing to the Kubernetes project.  I have read all of the
        user documentation listed on the community contributring page.
        What should I do next to get started?

3. Attend the sig-cli [bi-weekly meeting](https://github.com/kubernetes/community/tree/master/sig-cli).
  - Introduce yourself at the beginning of the meeting

## Expectations

If a sig-cli identifies a bug or feature to you to work on,
they will need your help to ensure that continual progress is
made and the fix / feature makes it into a Kubernetes release.
While you are working on the issue, you must leave a weekly update
including:

1. What has been finished
2. What is being worked on
3. What if anything is blocking progress

## Life of a sig-cli bug overview

1. File a GitHub issue
  - Mention `@kubernetes/sig-cli-bugs`
  - Describe steps to reproduce the issue including client / server version
2. Implement the fix
  - Send a PR
  - Add `Closes #<Issue Number>` to the description
  - Mention `@kubernetes/sig-cli-pr-reviews` in a comment
  - Incorporate review feedback
  - **Note:** Include unit and e2e tests
7. Release
  - Wait for your feature to appear in the next Kubernetes release!

## Life of a sig-cli feature overview

Picking up an issue and implementing it without getting approval for
the user experience and software design often results in wasted time
and effort due to decisions such as flag-names, command names, specific
command behavior.

In order to minimize wasted work and improve communication across efforts,
the user experience and software design must be agreed upon before
any PRs are sent for code review.

1. Identify a problem - GitHub issue with basic description
2. Propose a solution - GitHub design proposal PR
  - **Action:** Write a design proposal using the [template](contributors/design-proposals/sig-cli/template.md).  Once
    complete, create a pull request and mention `@kubernetes/sig-cli-misc`
  - Avoid "Work In Progress" PRs (WIP), only send the PR after you
    consider it complete.  To get early feedback, consider messaging the
    email discussion group.
  - Expect feedback from 2-3 different sig-cli community members in the form of PR comments
  - Incorporate feedback from sig-cli community members and comment "PTAL" (please take another look)
  - Proposal accepted by at least one sig-cli lead
  - Proposal merged
3. Communicate what will be done - Discussion at sig-cli bi-weekly meeting - See the [sig-cli community page](https://github.com/kubernetes/community/tree/master/sig-cli) for meeting time and location.
  - **Note:** This should be done *before* the proposal has been merged, and after an initial round of feedback - e.g. several folks have looked at it and left comments.
  - **Action:** Add the design proposal as an item to the [bi-weekly meeting agenda](https://docs.google.com/document/d/1r0YElcXt6G5mOWxwZiXgGu_X6he3F--wKwg-9UBc29I/edit)
  - Make sure the sig-cli community is aware of the work that is being done and can comment on it
  - Should be included in meeting notes sent to the sig-cli googlegroup
4. Add feature to Kubernetes release feature repo in [kubernetes/features](https://github.com/kubernetes/features/)
  - **Action:** *Done by sig-cli lead* - add the feature to the feature tracking repo tag
    to a release.  This is done to ensure release related decisions are
    properly made and communicated.  Such as defining alpha / beta / ga
    release, proper testing, and documentation are completed.
5. Implement the code - GitHub implementation PR
  - **Action:** Implement the solution described in the proposal, then
    send a PR - **include a link to the design proposal in the description**
  - Incorporate review feedback
  - **Note:** Include unit and e2e tests
6. Update related documentation
  - [Concept Docs](https://github.com/kubernetes/kubernetes.github.io/tree/master/docs/concepts/tools/kubectl)
7. Release
  - Wait for your feature to appear in the next Kubernetes release!

# New feature development

## Creating a GitHub issue

**Note:** For new contributors, it is recommended that you pick up an
existing issue with an approved design proposal by reaching out a sig-lead.
See the [Message to new contributors](#message-to-new-contributors).

In order to start a discussion around a feature, create a new GitHub issue
in the https://github.com/kubernetes/kubernetes repo and mention `@kubernetes/sig-cli-misc`.

Keep the issue to 2-4 sentence description of the basic description of
the problem and proposed solution.

**Note:** You must mention `@kubernetes/sig-cli-feature-requests` in the description of the
issue or in a comment for someone from the sig-cli community to look
at the issue.

## Creating a design proposal

**Note:** For new contributors, it is recommended that you pick up an
existing issue with an approved design proposal by reaching out to a sig-lead.
See the [Message to new contributors](#message-to-new-contributors).

Once at least one sig-lead has agreed that design and code review resources
can be allocated to tackle a particular issue, the fine details
of the user experience and design should be agreed upon.

A PR is filed against the https://github.com/kubernetes/community repo.
This will provide a chance for community members to comment on the
user experience and software design.

**Note:** This step is important to prevent a lot of code churn and
thrashing around issues like flag names, command names, etc.

**Note:** It is normal for sig-cli community members to push back on
feature requests and proposals. sig-cli development and review resources are extremely constrained.
If your proposal is not high-impact or urgent it may not be accepted
in favor of more pressing needs.  sig-cli community members should
be free to say: "No, not this release." or "No, not this year." or "No, not right now because although this
is desirable we need help on these other concrete issues before we can
tackle this." or "No, this problem should be solved in another way."

Create a design proposal PR under
https://github.com/kubernetes/community/tree/master/contributors/design-proposals/sig-cli/
by copying template.md.

Mention `@kubernetes/sig-cli-proposals` on the proposal and link to
the proposal from the original issue.  Expect to get feedback
from 2-3 community members.  At least
one sig lead must approve the design proposal for it to be accepted.

## Discussing at sig-cli

Finally, when a contributor picks up a design proposals and is ready
to begin implementing it, we should put it as an item on the next
sig-cli agenda.  This ensures that everyone in sig-cli community
is aware that work is being started.

## Updating the feature repo

The kubernetes/feature repo exists to help PMs, folks managing the release,
and sig leads coordinate the work going into a given release.  This
includes items like:

- ensuring all items slated for the release have been completed, or have
  been disabled and pushed into the next release.
- support level for items has been properly defined - alpha / beta / ga
- api changes have been properly vetted by the appropriate parties
- user facing documentation has been updated

**Note:** in most cases this will be done by sig leads on behalf of the
other community members.

## Implementing the feature

Contributors can begin implementing a feature before any of the above
steps have been completed, but should not send a PR until
the design proposal has been approved / merged.

Go [here](https://github.com/kubernetes/community/blob/master/contributors/devel/development.md)
for instructions on setting up the Kubernetes development environment.

**Note:** All new features must have automated tests in the same PR.
Small features and flag changes require only unit/integration tests,
while larger changes require both unit/integration tests and e2e tests.

If you get stuck or need help, reach out in the Kubernetes slack
channel and mention @jess and @adohe.

## Updating documentation

Most new features should include user facing documentation.  This is
important to make sure that users are aware of the cool new thing that
was added.  Depending on the contributor and size of the feature, this
may be done either by the same contributor that implemented the feature,
or another contributor who is more familiar with the existing docs
templates.

## Notes on release

Several weeks before a Kubernetes release, development enters a stabilization
period where no new features are merged.  For a feature to be accepted
into a release, it must be fully merged and tested by this time.  If
your feature is not fully complete - including tests - it will have
to wait until the next release.

# Escalation

## What to do if your bug issue is stuck

If an issue isn't getting any attention and is unresolved, mention
@kubernetes/sig-cli-bugs.
Highlight the severity and urgency of the issue.  For severe issues
escalate by contacting sig leads and attending the bi-weekly sig-cli
meeting.

## What to do if your feature request issue is stuck

If an issue isn't getting any attention and is unresolved, mention
@kubernetes/sig-cli-feature-requests.
If a particular issue has a high impact for you or your business,
make sure this is clear on the bug, and reach out to the sig leads
directly.  Consider attending the sig meeting to discuss over video
conference.

## What to do if your PR is stuck

It may happen that your PR seems to be stuck without clear actionable
feedback for a week or longer.  If your PR is based off a design proposal,
the chances of the PR getting stuck for a long period of time
are much smaller.  However, if it happens do the following:

- If your PR is stuck for a week or more because it has never gotten any
  comments, mention @kubernetes/sig-cli-pr-reviews and ask for attention.
  If this is not successful in a day or two, escalate to sig leads or
  attend the bi-weekly meeting.
- If your PR is stuck for a week or more after it got comments, but
  the attention has died down.  Mention the reviewer and comment with
  "PTAL" (please take another look).  If you are still not able to
  get any attention after a couple days, escalate to sig leads by
  mentioning them.

## What to do if your design docs is stuck

It may happen that your design doc gets stuck without getting merged
or additional feedback. This may happen for a number of reasons.  If you believe that your
design is important and has been dropped, or it is not moving forward,
please add it to the sig cli bi-weekly meeting agenda and email
kubernetes-sig-cli@googlegroups.com notifying the community you would
like to discuss it at the next meeting.

Merge state meanings:

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

## General escalation instructions if you get stuck

See the [sig-cli community page](https://github.com/kubernetes/community/tree/master/sig-cli) for points of contact and meeting times:

- attend the sig-cli bi-weekly meeting
- message one of the sig leads [on slack](https://kubernetes.slack.com/messages/sig-cli/)
- send an email to the [kubernetes-sig-cli@googlegroups.com](https://groups.google.com/forum/#!forum/kubernetes-sig-cli)
