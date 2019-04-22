# Kubernetes Contributor Cheat Sheet

A list of common resources when contributing to Kubernetes, tips, tricks, and common best practices used within the
Kubernetes project. It is a "TL;DR" or quick reference of useful information to
make your GitHub contribution experience better.

**Table of Contents**
- [Getting Started](#getting-started)
- [Workflow](#workflow)
- [SIGs and Working Groups](#sigs-and-working-groups)
- [Community](#community)
- [Tests](#tests)
- [Email Aliases](#email-aliases)
- [Communicating effectively on GitHub](#Communicating-Effectively-on-GitHub)
  - [How to be Excellent to Each Other](#How-to-be-Excellent-to-Each-Other)
    - [Examples of Good/Bad Communication](#Examples-of-GoodBad-Communication)
- [Submitting a Contribution](#Submitting-a-Contribution)
  - [Opening and Responding to Issues](#Opening-and-Responding-to-Issues)
    - [Creating an Issue](#Creating-an-Issue)
    - [Responding to an Issue](#Responding-to-an-Issue)
  - [Opening a Pull Request](#Opening-a-pull-Request)
    - [Creating a Pull Request](#Creating-a-Pull-Request)
    - [Example PR Description](#Example-PR-Description)
    - [Troubleshooting a Pull Request](#Troubleshooting-a-Pull-Request)
  - [Labels](#Labels)
- [Working Locally](#Working-Locally)
  - [Branch Strategy](#Branch-Strategy)
    - [Adding Upstream](#Adding-Upstream)
    - [Keeping Your Fork in Sync](#Keeping-Your-Fork-in-Sync)
  - [Squashing Commits](#Squashing-Commits)
- [Other](#other)

| Repo | PRs | Issues | Notes |
| ---- | --- | ------ | ----- |
| [Kubernetes](https://github.com/kubernetes/kubernetes) | [PRs](https://github.com/kubernetes/kubernetes/pulls) | [Issues](https://github.com/kubernetes/kubernetes/issues) | [Meeting Notes](http://bit.ly/kubenotes)
| [Community](https://github.com/kubernetes/community) | [PRs](https://github.com/kubernetes/community/pulls) | [Issues](https://github.com/kubernetes/community/issues) |
| [Docs](https://github.com/kubernetes/website) | [PRs](https://github.com/kubernetes/website/pulls) | [Issues](https://github.com/kubernetes/website/issues)

## Getting Started

- [Contributor Guide](https://github.com/kubernetes/community/blob/master/contributors/guide/README.md) 

## Workflow

- [Gubernator Dashboard - k8s.reviews](https://k8s-gubernator.appspot.com/pr)
- [Tide](https://prow.k8s.io/tide)
- [Bot commands](https://go.k8s.io/bot-commands)
- [GitHub labels](https://go.k8s.io/github-labels)
- [Release Buckets](https://gcsweb.k8s.io/gcs/kubernetes-release/)
- Developer Guide
  - [Cherry Picking Guide](/contributors/devel/sig-release/cherry-picks.md)
- [Kubernetes Code Search](https://cs.k8s.io/), maintained by [@dims](https://github.com/dims)


## SIGs and Working Groups

- [Master SIG list](/sig-list.md#master-sig-list)

## Community

- [Calendar](https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com)
- [kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)
- [Kubernetes Forums](https://discuss.kubernetes.io)
- [Slack channels](http://slack.k8s.io/)
- [StackOverflow](https://stackoverflow.com/questions/tagged/kubernetes)
- [YouTube Channel](https://www.youtube.com/c/KubernetesCommunity/)

## Tests

- [Current Test Status](https://prow.k8s.io/)
- [Aggregated Failures](https://go.k8s.io/triage)
- [Test Grid](https://testgrid.k8s.io)
- [Test Health](https://go.k8s.io/test-health)
- [Test History](https://go.k8s.io/test-history)

## Email Aliases

- community@kubernetes.io - Mail someone on the community team (SIG Contributor Experience) about a community issue.
- social@cncf.io - Contact the CNCF social team; blog, twitter account, and other social properties.
- steering@kubernetes.io - Mail the steering committee. Public address with public archive.
- steering-private@kubernetes.io - Mail the steering committee privately, for sensitive items.
- conduct@kubernetes.io - Contact the Code of Conduct committee, private mailing list.

## Communicating Effectively on GitHub


### How to be Excellent to Each Other

As a first step, familiarize yourself with the [Code of Conduct].


#### Examples of Good/Bad Communication

When raising an issue, or seeking assistance, please be polite with your request:

  🙂 “X doesn’t compile when I do Y, do you have any suggestions?”

  😞 “X doesn’t work! Please fix it!”

When closing a PR, convey an explanatory and cordial message explaining
why it does not meet the requirements to be merged.

🙂 “I’m closing this PR because this feature can’t support the use case X. In
   it's proposed form, it would be a better to be implemented with Y tool. Thank
    you for working on this.”

😞 “Why isn’t this following the API conventions? This should be done elsewhere!”


## Submitting a Contribution

### Opening and Responding to Issues

GitHub Issues are the primary means of tracking things such as bug reports,
enhancement requests, or reporting other issues such as failing tests. They are
**not** intended for [user support requests]. For those, please check with the
[troubleshooting guide], report the problem to [Stack Overflow] or follow up on
the Kubernetes [User forum].

**References:**
- [Labels]
- [Prow commands][commands]


#### Creating an Issue

- Use an issue template if one is available. Using the correct one will aid other
  contributors in responding to your issue.
  - Follow any directions described in the issue template itself.
- Be descriptive with the issue you are raising.
- Assign appropriate [labels]. If you are unsure, the [k8s-ci-robot][prow] bot
  ([Kubernetes CI bot][prow]) will reply to your issue with the needed labels
  for it to be effectively triaged.
- Be selective when assigning issues using [`/assign @<username>`][assign] or
  [`/cc @<username>`][cc]. Your issue will be triaged more effectively applying
  correct labels over assigning more people to the issue.

#### Responding to an Issue

- When tackling an issue, comment on it letting others know you are working on
  it to avoid duplicate work.
- When you have resolved something on your own at any future time, comment on
  the issue letting people know before closing it.
- Include references to other PRs or issues (or any accessible materials),
  example: _"ref: #1234"_. It is useful to identify that related work has been
  addressed somewhere else.


### Opening a Pull Request

Pull requests (PR) are the main means of contributing code, documentation or
other forms of work that would be stored within a git repository.

**References:**
- [Labels]
- [Prow commands][commands]
- [Pull request process]
- [Github workflow]

#### Creating a Pull Request

- Follow the directions of the pull request template if one is available. It
  will help those that respond to your PR.
- If a [trivial fix] such as a broken link, typo or grammar mistake, review the
  entire document for other potential mistakes. Do not open multiple PRs for
  small fixes in the same document.
- Reference any issues related to your PR, or issues that PR may solve.
- Avoid creating overly large changes in a single commit. Instead, break your PR
  into multiple small, logical commits. This makes it easier for your PR to be
  reviewed.
- Comment on your own PR where you believe something may need further
  explanation.
- Be selective when assigning your PR with [`/assign @<username>`][assign].
  Assigning excessive reviewers will not yield a quicker PR review.
- If your PR is considered a _"Work in progress"_ prefix the name with `[WIP]`
  or use the [`/hold`][hold] command. This will prevent the PR from being merged
  till the `[WIP]` or hold is lifted.
- If you have not had your PR reviewed, do not close and open a new PR with the
  same changes. Ping your reviewers in a comment with `@<github username>`.


#### Example PR Description

```
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

What's in that PR:
- **Line 1** - Reference to other issues or PRs (#3064 #3097).
- **Line 2** - A brief description of what is being done in the PR.
- **Line 4** - [SIG][sigs] assignment with the [command][commands]
  `/sig contributor-experience`..
- **Line 5** - Reviewers that may have interest on this specific issue or PR are
  specified with the [`/cc`][cc] command.
- **Line 6** - The [`/kind cleanup`][kind] command add a [label][labels] that
  categorizes issue or PR as related to cleaning up code, process, or technical
  debt.
- **Line 7** - The [`/area developer-guide`][kind] command categorizes issue or
  PR as related to the developer guide.
- **Line 8** - The command [`/assign`][assign] assigns an approver to the PR.
  An approver will be suggested by the [k8s-ci-robot][prow] and is selected from
  the list of owners in the [OWNERS] file. They will add the
  [`/approve`][approve] label to the PR after it has been reviewed.


#### Troubleshooting a Pull Request

<!---
include link to CI troubleshooting when  this is merged:
https://github.com/kubernetes/community/pull/3143
--->

After your PR is proposed, a series of tests are executed by the Kubernetes CI
platform, [Prow]. If any of the tests failed, the [k8s-ci-robot][prow]
will reply to the PR with links to the failed tests and available logs.

Pushing new commits to your PR will automatically trigger the tests to re-run.

Occasionally there can be issues with Kubernetes CI platform. These can occur
for a wide variety of reasons even if your contribution passes all local
tests. You can trigger a re-run of the tests with the `/retest` command.


### Labels

Kubernetes uses [labels] to categorize and triage issues and Pull Requests.
Applying the right labels will help your issue or PR be triaged more
effectively.

**References:**
- [Labels]
- [Prow commands][commands]

Frequently used labels:
- [`/sig <sig name>`][kind] Assign a [SIG][SIGs] to the ownership of the issue
  or PR.
- [`/area <area name>`][kind] Associate the issue or PRs to a specific
  [area][labels].
- [`/kind <category>`][kind] [Categorizes][labels] the issue or PR.


## Working Locally

Before you propose a pull request, you will have to do some level of work
locally. If you are new to git, the [Atlassian git tutorial] is a good starting
point. As an alternative, Stanford's [Git magic] tutorial is a good
multi-language option.

**References:**
- [Atlassian git tutorial]
- [Git magic]
- [Github workflow]
- [Testing locally]
- [Developer guide]


### Branch Strategy

The Kubernetes project uses a _"Fork and Pull"_ workflow that is standard to
GitHub. In git terms, your personal fork is referred to as the _"`origin`"_ and
the actual project's git repository is called _"`upstream`"_. To keep your
personal branch (`origin`) up to date with the project (`upstream`), it must be
configured within your local working copy.


#### Adding Upstream

Add `upstream` as a remote, and configure it so you cannot push to it.

```
# replace <upstream git repo> with the upstream repo url
# example:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <upstream git repo>
git remote set-url --push upstream no_push
```

This can be verified by running `git remote -v` which will list your configured
remotes.


#### Keeping Your Fork in Sync

Fetch all the changes from `upstream` and _"rebase"_ them on your local `master`
branch. This will sync your local repo with the `upstream` project.

```
git fetch upstream
git checkout master
git rebase upstream/master
```

You should do this minimally before creating a new branch to work on your
feature or fix.

```
git checkout -b myfeature
```

#### Squashing Commits

The main purpose of [squashing commits] is to create a clean readable git
history or log of the changes that were made. Usually this is done in last
phase of a PR revision. If you are unsure if you should squash your commits, it
is better to err on the side of having more and leave it up to the judgement of
the other contributors assigned to review and approve your PR.

## Other

- [Developer Statistics](https://k8s.devstats.cncf.io)

[code of conduct]: /code-of-conduct.md
[user support request]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[troubleshooting guide]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[stack overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[kubernetes forum]: https://discuss.kubernetes.io/
[pull request process]: /contributors/guide/pull-requests.md
[github workflow]: /contributors/guide/github-workflow.md
[prow]: https://git.k8s.io/test-infra/prow#prow
[commands]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[cc]: https://prow.k8s.io/command-help#hold
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[trivial fix]: /contributors/guide/pull-requests.md#10-trivial-edits
[Github workflow]: /contributors/guide/github-workflow.md#3-branch
[squashing commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[owners]: /contributors/guide/owners.md
[testing locally]: /contributors/guide/README.md#testing
[developer guide]: /contributors/devel/README.md
[Atlassian git tutorial]: https://www.atlassian.com/git/tutorials
[git magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
