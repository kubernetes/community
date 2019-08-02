---
title: "Pull Request Process"
weight: 1
slug: "pull-requests"
---

This doc explains the process and best practices for submitting a pull request to the [Kubernetes project](https://github.com/kubernetes/kubernetes) and its associated sub-repositories. It should serve as a reference for all contributors, and be useful especially to new and infrequent submitters.

- [Before You Submit a Pull Request](#before-you-submit-a-pull-request)
  * [Run Local Verifications](#run-local-verifications)
- [The Pull Request Submit Process](#the-pull-request-submit-process)
  * [The Testing and Merge Workflow](#the-testing-and-merge-workflow)
  * [Marking Unfinished Pull Requests](#marking-unfinished-pull-requests)
  * [Comment Commands Reference](#comment-commands-reference)
  * [Automation](#automation)
  * [How the e2e Tests Work](#how-the-e2e-tests-work)
- [Why was my Pull Request closed?](#why-was-my-pull-request-closed)
- [Why is my Pull Request not getting reviewed?](#why-is-my-pull-request-not-getting-reviewed)
- [Best Practices for Faster Reviews](#best-practices-for-faster-reviews)
  * [0. Familiarize yourself with project conventions](#0-familiarize-yourself-with-project-conventions)
  * [1. Is the feature wanted? File a Kubernetes Enhancement Proposal](#1-is-the-feature-wanted-file-a-kubernetes-enhancement-proposal)
  * [2. Smaller Is Better: Small Commits, Small Pull Requests](#2-smaller-is-better-small-commits-small-pull-requests)
  * [3. Open a Different Pull Request for Fixes and Generic Features](#3-open-a-different-pull-request-for-fixes-and-generic-features)
  * [4. Comments Matter](#4-comments-matter)
  * [5. Test](#5-test)
  * [6. Squashing and Commit Titles](#6-squashing-and-commit-titles)
  * [7. KISS, YAGNI, MVP, etc.](#7-kiss-yagni-mvp-etc)
  * [8. It's OK to Push Back](#8-its-ok-to-push-back)
  * [9. Common Sense and Courtesy](#9-common-sense-and-courtesy)
  * [10. Trivial Edits](#10-trivial-edits)

# Before You Submit a Pull Request

This guide is for contributors who already have a pull request to submit. If you're looking for information on setting up your developer environment and creating code to contribute to Kubernetes, see the [development guide](/contributors/devel/development.md).

First-time contributors should head to the [Contributor Guide](/contributors/guide/README.md) to get started.

**Make sure your pull request adheres to our best practices. These include following project conventions, making small pull requests, and commenting thoroughly. Please read the more detailed section on [Best Practices for Faster Reviews](#best-practices-for-faster-reviews) at the end of this doc.**

## Run Local Verifications

You can run these local verifications before you submit your pull request to predict the
pass or fail of continuous integration.

* Run and pass `make verify` (can take 30-40 minutes)
* Run and pass `make test`
* Run and pass `make test-integration`

# The Pull Request Submit Process

Merging a pull request requires the following steps to be completed before the pull request will be merged automatically.

- [Sign the CLA](https://git.k8s.io/community/CLA.md) (prerequisite)
- [Open a pull request](https://help.github.com/articles/about-pull-requests/)
  - *For kubernetes/kubernetes repository only:* Add [release notes](/contributors/guide/release-notes.md) if needed.
- Pass all e2e tests
- Get all necessary approvals from reviewers and code owners

## The Testing and Merge Workflow

The Kubernetes merge workflow uses labels, applied by [commands](https://prow.k8s.io/command-help) via comments. These will trigger actions on your pull request. Different Kubernetes repositories may require different labels on the path to approval. A generic explanation of how labels are used in pull requests can be found [here](/contributors/guide/owners.md#code-review-using-owners-files). The pull request bot will also automatically apply and/or suggest labels.

_Example:_ To apply a SIG label, you would type in a comment:
```
/sig aws
```

*NOTE: For pull requests that are in progress but not ready for review,
prefix the pull request title with `WIP` or `[WIP]` and track any remaining TODOs
in a checklist in the pull request description.*

Here's the process the pull request goes through on its way from submission to merging:

1. Make the pull request
1. `@k8s-ci-robot` assigns reviewers

1. If you're **not** a member of the Kubernetes organization, a Reviewer/Kubernetes Member checks that the pull request is safe to test. If so, they comment `/ok-to-test`. Pull requests by Kubernetes organization [members](/community-membership.md) do not need this step. Now the pull request is considered to be trusted, and the pre-submit tests will run:

    1. Automatic tests run. See the current list of tests at this [link](https://prow.k8s.io/?repo=kubernetes%2Fkubernetes&type=presubmit)
    1. If tests fail, resolve issues by pushing edits to your pull request branch
    1. If the failure is a flake, anyone on trusted pull requests can comment `/retest` to rerun failed tests

1. Reviewer suggests edits
1. Push edits to your pull request branch
1. Repeat the prior two steps as needed until the reviewer(s) add `/lgtm` label. The `/lgtm` label, when applied by someone listed as a `reviewer` in the corresponding project `OWNERS` file, is a signal that the code has passed review from one or more trusted reviewers for that project
1. (Optional) Some reviewers prefer that you squash commits at this step
1. Follow the bot suggestions to assign an OWNER who will add the `/approve` label to the pull request. The `/approve` label, when applied by someone listed as an `approver` in the corresponding project `OWNERS`, is a signal that the code has passed final review and is ready to be automatically merged

The behavior of Prow is configurable across projects. You should be aware of the following configurable behaviors.

* If you are listed as an `/approver` in the `OWNERS` file, an implicit `/approve` can be applied to your pull request. This can result in a merge being triggered by a `/lgtm` label. This is the configured behavior in many projects, including `kubernetes/kubernetes`. You can remove the implicit `/approve` with `/approve cancel`
* `/lgtm` can be configured so that from someone listed as both a `reviewer` and an `approver` will cause both labels to be applied. For `kubernetes/kuebernetes` and many other projects this is _not_ the default behavior, and `/lgtm` is decoupled from `/approve`

Once the tests pass and the reviewer adds the `lgtm` and `approved` labels, the pull request enters the final merge pool. The merge pool is needed to make sure no incompatible changes have been introduced by other pull requests since the tests were last run on your pull request.
<!-- TODO: create parallel instructions for reviewers -->

[Tide](https://git.k8s.io/test-infra/prow/cmd/tide) will manage the merge pool
automatically. It uses GitHub queries to select PRs into “tide pools”,
runs as many in a batch as it can (“tide comes in”), and merges them (“tide goes out”).

1. The pull request enters the [merge pool](https://prow.k8s.io/tide)
if the merge criteria are met. The [PR dashboard](https://prow.k8s.io/pr) shows
the difference between your PR's state and the merge criteria so that you can
easily see all criteria that are not being met and address them.
1. If tests fail, resolve issues by pushing edits to your pull request branch
1. If the failure is a flake, anyone can comment `/retest` if the pull request is trusted
1. If tests pass, Tide automatically merges the pull request

That's the last step. Your pull request is now merged.

## Marking Unfinished Pull Requests

If you want to solicit reviews before the implementation of your pull request is complete, you should hold your pull request to ensure that Tide does not pick it up and attempt to merge it. There are two methods to achieve this:

1. You may add the `/hold` or `/hold cancel` comment commands
2. You may add or remove a `WIP` or `[WIP]` prefix to your pull request title

The GitHub robots will add and remove the `do-not-merge/hold` label as you use the comment commands and the `do-not-merge/work-in-progress` label as you edit your title. While either label is present, your pull request will not be considered for merging.

## Pull Requests and the Release Cycle

If a pull request has been reviewed but held or not approved, it might be due to the current phase in the [Release Cycle](/contributors/devel/sig-release/release.md). Occasionally, a SIG may freeze their own code base when working towards a specific feature or goal that could impact other development. During this time, your pull request could remain unmerged while their release work is completed.

If you feel your pull request is in this state, contact the appropriate [SIG](https://git.k8s.io/community/sig-list.md) or [SIG-Release](https://git.k8s.io/sig-release) for clarification.

## Comment Commands Reference

[The commands doc](https://go.k8s.io/bot-commands) contains a reference for all comment commands.

## Automation

The Kubernetes developer community uses a variety of automation to manage pull requests.  This automation is described in detail [in the automation doc](/contributors/devel/automation.md).

## How the e2e Tests Work

The end-to-end tests will post the status results to the pull request. If an e2e test fails,
`@k8s-ci-robot` will comment on the pull request with the test history and the
comment-command to re-run that test. e.g.

> The following tests failed, say /retest to rerun them all.

# Why was my pull request closed?

Pull requests older than 90 days will be closed. Exceptions can be made for pull requests that have active review comments, or that are awaiting other dependent pull requests. Closed pull requests are easy to recreate, and little work is lost by closing a pull request that subsequently needs to be reopened. We want to limit the total number of pull requests in flight to:
* Maintain a clean project
* Remove old pull requests that would be difficult to rebase as the underlying code has changed over time
* Encourage code velocity

# Why is my pull request not getting reviewed?

A few factors affect how long your pull request might wait for review.

If it's the last few weeks of a milestone, we need to reduce churn and stabilize.

Or, it could be related to best practices. One common issue is that the pull request is too big to review. Let's say you've touched 39 files and have 8657 insertions. When your would-be reviewers pull up the diffs, they run away - this pull request is going to take 4 hours to review and they don't have 4 hours right now. They'll get to it later, just as soon as they have more free time (ha!).

There is a detailed rundown of best practices, including how to avoid too-lengthy pull requests, in the next section.

But, if you've already followed the best practices and you still aren't getting any pull request love, here are some
things you can do to move the process along:

   * Make sure that your pull request has an assigned reviewer (assignee in GitHub). If not, reply to the pull request comment stream asking for a reviewer to be assigned. This is done via a [bot command](https://prow.k8s.io/command-help) (the bot may have suggestions for this) and looks like this: `/assign @username`.

   * Ping the assignee (@username) on the pull request comment stream, and ask for an estimate of when they can get to the review.

   * Ping the assignee on [Slack](http://slack.kubernetes.io). Remember that a person's GitHub username might not be the same as their Slack username.

   * Ping the assignee by email (many of us have publicly available email addresses).

   * If you're a member of the organization ping the [team](https://github.com/orgs/kubernetes/teams) (via @team-name) that works in the area you're submitting code to.

   * If you have fixed all the issues from a review, and you haven't heard back, you should ping the assignee on the comment stream with a "please take another look" (`PTAL`) or similar comment indicating that you are ready for another review.

Read on to learn more about how to get faster reviews by following best practices.

# Best Practices for Faster Reviews

Most of this section is not specific to Kubernetes, but it's good to keep these best practices in mind when you're making a pull request.

You've just had a brilliant idea on how to make Kubernetes better. Let's call that idea Feature-X. Feature-X is not even that complicated. You have a pretty good idea of how to implement it. You jump in and implement it, fixing a bunch of stuff along the way. You send your pull request - this is awesome! And it sits. And sits. A week goes by and nobody reviews it. Finally, someone offers a few comments, which you fix up and wait for more review. And you wait. Another week or two go by. This is horrible.

Let's talk about best practices so your pull request gets reviewed quickly.

## 0. Familiarize yourself with project conventions

* [Development guide](/contributors/devel/development.md)
* [Coding conventions](../guide/coding-conventions.md)
* [API conventions](/contributors/devel/sig-architecture/api-conventions.md)
* [Kubectl conventions](/contributors/devel/sig-cli/kubectl-conventions.md)

## 1. Is the feature wanted? File a Kubernetes Enhancement Proposal
Are you sure Feature-X is something the Kubernetes team wants or will accept? Is it implemented to fit with other changes in flight? Are you willing to bet a few days or weeks of work on it?

It's better to get confirmation beforehand.

When you want to make a large or otherwise significant change, you should follow the [Kubernetes Enhancement Proposal process](https://git.k8s.io/enhancements/keps/0001-kubernetes-enhancement-proposal-process.md).

Even for small changes, it is often a good idea to gather feedback on an issue you filed, or even simply ask in the appropriate SIG's Slack channel to invite discussion and feedback from code owners. Here's a [list of SIGs](/sig-list.md).


## 2. Smaller Is Better: Small Commits, Small Pull Requests

Small commits and small pull requests get reviewed faster and are more likely to be correct than big ones.

Attention is a scarce resource. If your pull request takes 60 minutes to review, the reviewer's eye for detail is not as keen in the last 30 minutes as it was in the first. It might not get reviewed at all if it requires a large continuous block of time from the reviewer.

**Breaking up commits**

Break up your pull request into multiple commits, at logical break points.

Making a series of discrete commits is a powerful way to express the evolution of an idea or the
different ideas that make up a single feature. Strive to group logically distinct ideas into separate commits.

For example, if you found that Feature-X needed some prefactoring to fit in, make a commit that JUST does that prefactoring. Then make a new commit for Feature-X.

Strike a balance with the number of commits. A pull request with 25 commits is still very cumbersome to review, so use
judgment.

**Breaking up Pull Requests**

Or, going back to our prefactoring example, you could also fork a new branch, do the prefactoring there and send a pull request for that. If you can extract whole ideas from your pull request and send those as pull requests of their own, you can avoid the painful problem of continually rebasing.

Kubernetes is a fast-moving codebase - lock in your changes ASAP with your small pull request, and make merges be someone else's problem.

Multiple small pull requests are often better than multiple commits. Don't worry about flooding us with pull requests. We'd rather have 100 small, obvious pull requests than 10 unreviewable monoliths.

We want every pull request to be useful on its own, so use your best judgment on what should be a pull request vs. a commit.

As a rule of thumb, if your pull request is directly related to Feature-X and nothing else, it should probably be part of the Feature-X pull request. If you can explain why you are doing seemingly no-op work ("it makes the Feature-X change easier, I promise") we'll probably be OK with it. If you can imagine someone finding value independently of Feature-X, try it as a pull request. (Do not link pull requests by `#` in a commit description, because GitHub creates lots of spam. Instead, reference other pull requests via the pull request your commit is in.)

## 3. Open a Different pull request for Fixes and Generic Features

**Put changes that are unrelated to your feature into a different pull request.**

Often, as you are implementing Feature-X, you will find bad comments, poorly named functions, bad structure, weak type-safety, etc.

You absolutely should fix those things (or at least file issues, please) - but not in the same pull request as your feature. Otherwise, your diff will have way too many changes, and your reviewer won't see the forest for the trees.

**Look for opportunities to pull out generic features.**

For example, if you find yourself touching a lot of modules, think about the dependencies you are introducing between packages. Can some of what you're doing be made more generic and moved up and out of the Feature-X package? Do you need to use a function or type from an otherwise unrelated package? If so, promote! We have places for hosting more generic code.

Likewise, if Feature-X is similar in form to Feature-W which was checked in last month, and you're duplicating some tricky stuff from Feature-W, consider prefactoring the core logic out and using it in both Feature-W and
Feature-X. (Do that in its own commit or pull request, please.)

## 4. Comments Matter

In your code, if someone might not understand why you did something (or you won't remember why later), comment it. Many code-review comments are about this exact issue.

If you think there's something pretty obvious that we could follow up on, add a TODO.

Read up on [GoDoc](https://blog.golang.org/godoc-documenting-go-code) - follow those general rules for comments.

## 5. Test

Nothing is more frustrating than starting a review, only to find that the tests are inadequate or absent. Very few pull requests can touch the code and NOT touch tests.

If you don't know how to test Feature-X, please ask!  We'll be happy to help you design things for easy testing or to suggest appropriate test cases.

## 6. Squashing and Commit Titles

Your reviewer has finally sent you feedback on Feature-X.

Make the fixups, and don't squash yet. Put them in a new commit, and re-push. That way your reviewer can look at the new commit on its own, which is much faster than starting over.

We might still ask you to clean up your commits at the very end for the sake of a more readable history, but don't do this until asked: typically at the point where the pull request would otherwise be tagged `LGTM`.

Each commit should have a good title line (<70 characters) and include an additional description paragraph describing in more detail the change intended.

**General squashing guidelines:**

* Sausage => squash

 Do squash when there are several commits to fix bugs in the original commit(s), address reviewer feedback, etc. Really we only want to see the end state and commit message for the whole pull request.

* Layers => don't squash

 Don't squash when there are independent changes layered to achieve a single goal. For instance, writing a code munger could be one commit, applying it could be another, and adding a precommit check could be a third. One could argue they should be separate pull requests, but there's really no way to test/review the munger without seeing it applied, and there needs to be a precommit check to ensure the munged output doesn't immediately get out of date.

A commit, as much as possible, should be a single logical change.

## 7. KISS, YAGNI, MVP, etc.

Sometimes we need to remind each other of core tenets of software design - Keep It Simple, You Aren't Gonna Need It, Minimum Viable Product, and so on. Adding a feature "because we might need it later" is antithetical to software that ships. Add the things you need NOW and (ideally) leave room for things you might need later - but don't implement them now.

## 8. It's OK to Push Back

Sometimes reviewers make mistakes. It's OK to push back on changes your reviewer requested. If you have a good reason for doing something a certain way, you are absolutely allowed to debate the merits of a requested change. Both the reviewer and reviewee should strive to discuss these issues in a polite and respectful manner.

You might be overruled, but you might also prevail. We're pretty reasonable people. Mostly.

Another phenomenon of open-source projects (where anyone can comment on any issue) is the dog-pile - your pull request gets so many comments from so many people it becomes hard to follow. In this situation, you can ask the primary reviewer (assignee) whether they want you to fork a new pull request to clear out all the comments. You don't HAVE to fix every issue raised by every person who feels like commenting, but you should answer reasonable comments with an explanation.

## 9. Common Sense and Courtesy

No document can take the place of common sense and good taste. Use your best judgment, while you put
a bit of thought into how your work can be made easier to review. If you do these things your pull requests will get merged with less friction.

## 10. Trivial Edits

Each incoming Pull Request needs to be reviewed, checked, and then merged.

While automation helps with this, each contribution also has an engineering cost. Therefore it is appreciated if you do NOT make trivial edits and fixes, but instead focus on giving the entire file a review.

If you find one grammatical or spelling error, it is likely there are more in that file, you can really make your Pull Request count by checking the formatting, checking for broken links, and fixing errors and then submitting all the fixes at once to that file.

**Some questions to consider:**

* Can the file be improved further?
* Does the trivial edit greatly improve the quality of the content?
