---
title: "Pull Request Review Guidelines"
weight: 13
description: |
  A collection of of tips to help both those looking to get their PR reviewed
  and the code reviewers themselves.
---
- [Tips for getting your PR reviewed](#tips-for-getting-your-pr-reviewed)
  - [Use Tokens to signal the state of your PR](#use-tokens-to-signal-the-state-of-your-pr)
  - [Reduce the scope of your PR](#reduce-the-scope-of-your-pr)
  - [Split your PR into logical smaller commits](#split-your-pr-into-logical-smaller-commits)
  - [Squash or remove unnecessary commits](#squash-or-remove-unnecessary-commits)
  - [Explain your changes (comments)](#explain-your-changes-comments)
  - [Ensure you properly test your changes](#ensure-you-properly-test-your-changes)
  - [Commit messages should be thorough and detailed](#commit-messages-should-be-thorough-and-detailed)
  - [Pull Request comments should summarize the PR](#pull-request-comments-should-summarize-the-pr)
  - [Minor changes (typos / link fixes)](#minor-changes-typos--link-fixes)
- [Tips for code reviewers](#tips-for-code-reviewers)
  - [Managing time](#managing-time)
    - [Block time off to review](#block-time-off-to-review)
    - [Taking a break or stepping away](#taking-a-break-or-stepping-away)
  - [Loop in others when domain specific knowledge is needed](#loop-in-others-when-domain-specific-knowledge-is-needed)
  - [Asking questions](#asking-questions)
  - [Asking for changes](#asking-for-changes)
    - [Commit Hygiene](#commit-hygiene)
  - [Be clear on progress and time](#be-clear-on-progress-and-time)
  - [Checking out a PR](#checking-out-a-pr)
- [Additional Resources](#additional-resources)


## Tips for getting your PR reviewed


### Use Tokens to signal the state of your PR

Token phrases like `WIP` or `do not merge` signal to reviewers that the PR is
potentially in a state that is not complete. This can save significant time for
reviewers on prioritizing which PR should be reviewed. They can either hold off
on reviewing altogether or will review the PR with the knowledge that it is not
complete.


### Reduce the scope of your PR

Extremely large PRs take significant amounts of time to both review and evaluate
for potential upstream/downstream impacts. If it is possible to achieve the
desired goal over multiple smaller, more concise PR ("Pre-factor"), it decreases
risk and will have a greater chance of being accepted.


### Split your PR into logical smaller commits

Multiple smaller commits that contain the logical pieces of work simplify the
review process for others to review your code. These commits should be able to
stand on their own and not necessarily depend on other things.

**Example: Markdown**

A PR that applies both formatting changes and content changes. If taken together
in a single commit, it will appear to your reviewer that the entire document has
changed and will require more time to compare and separate the actual content
changes from the style changes.

If content and style changes were made in separate commits, they could be reviewed
independently for their changes.

**Example: Generated code**

A PR that has both authored code and generated code can be difficult to parse.
If the authored code and generated code are in separate commits, the generated
code can largely be ignored by the reviewer, and they can focus their time on
the authored code.

### Squash or remove unnecessary commits

Often you may rapidly commit small changes or fixes. These are even more common
when accepting suggestions through the GitHub UI where each change is done as an
independent small commit. [Squashing] these down into a logical commit will both
help your reviewer when performing a final pass, and provide for a clean git
commit history.


### Explain your changes (comments)

Comments should be used for **ANY** non-obvious changes. They should explain why
something was done and the intent behind it; leave no surprises for the next
person to view or touch your code. If you spent time debugging, testing, or 
searching for prior art on a means to do something - document it in the comments.

If you leave a `TODO`, it should explain why it needs to be revisited along with
any relevant information that could help the next person looking to update that
part of the code. Consider following insertion of a `TODO` with creating a new
_help wanted_ issue, if appropriate.

This is also true for things that should **NOT** be updated or updated with great
care. It should document why it cannot be further reduced or changed. For an
example of this, see the _"[KEEP THE SPACE SHUTTLE FLYING]"_ portion of the
persistent volume controller as an example.


### Ensure you properly test your changes

Code submitted without tests will almost always be rejected at first pass. If an
area you're updating did not previously have them, adding them in your PR will
be appreciated and helps improve the overall health of the project.

If something you change has the potential to impact the performance or scalability
of the project, include a benchmark. These are essential to demonstrate the 
efficacy of your change and will help in preventing potential regressions.


### Commit messages should be thorough and detailed

PR comments are not represented in the commit history. Important details
regarding your changes should be kept in your commit message.

The commit message subject line should be no more than 72 characters and
summarize the changes being made.

Use the body of the commit message to describe **WHY** the change is being made.
Describing why something has changed is more important than how. You are
providing context to both your reviewer and the next person that has to touch
your code.

For more guidance on commit messages, see our [commit message guidelines].


### Pull Request comments should summarize the PR

The GitHub pull request should contain a summary of the changes being made as
well as links to all the relevant information related to the PR itself. These
are things such as issues it closes, KEPs, and any other relevant documentation.

Accurately filling out the issue template and applying the necessary labels
will greatly speed up the initial triage of your PR.

If there is a portion of your PR that should be explained or linked that is not
committed as a comment, you can review your own PR, applying additional context
or links.


### Minor changes (typos / link fixes)

While these sort of fixes are appreciated, try and fix multiple instances at the
same time if possible, instead of issuing subsequent individual PRs for each fix.

---

## Tips for code reviewers

### Managing time

#### Block time off to review

Often it can be hard to find dedicated, uninterrupted time to review PRs. If you
can, allocate some time and block it off on your calendar to help stay on top of
the incoming queue.

#### Taking a break or stepping away

If you are taking a break, going on vacation, or stepping away for a bit -- you
can set your [GitHub status] to busy; this will signal [Blunderbuss] to not
automatically assign you to reviews (It does not block manual assignment).

If for an extended period of time, or if you need to focus on other areas,
consider setting yourself as an `[emeritus_approver]` in some of the areas your
an approver.


### Loop in others when domain specific knowledge is needed

Kubernetes has an incredibly large and complex code base with sections that may
require domain-specific knowledge. If you are unsure or uncomfortable reviewing
a portion of a PR, it is better to decline a review and reassign to an owner or 
contributor with more expertise in that area.

If you are brought in for your knowledge in a specific area, try and provide
meaningful comments to serve as breadcrumbs in the future to help others gain a
better understanding of that area.


### Asking questions

You are encouraged to ask questions and seek an understanding of what the PR is
doing; however, your question might be answered further into the review. You can
stage your questions, and before you submit your review, revisit your own comments
to see if they're still relevant or update them after gaining further context. 

Often a question may turn into a request for further comments or changes to
explain what is happening at that specific point.

In your questions, try and be empathetic when phrasing. Instead of:

_"Why did you do this?"_

try

_"Am I understanding this correctly? Can you explain why...?"_

Remember a review is a discussion, often with multiple parties -- be reasonable.
Try to focus and summarize in ways which constructively move the conversation
forward instead of retreading ground.

### Asking for changes

It's okay to ask for changes to be made on a PR. In your comments, you should be
clear on what is a 'nit' or small thing to be improved and a **required** change
needed to accept the PR.

Be clear and state upfront architectural or larger changes. These should be
resolved first before addressing any further nits.

It's also okay to say _"No"_ to a PR. As a community, we want to respect people's
time and efforts, but sometimes things just don't make sense to accept. As
reviewers, you are the stewards of the code-base and sometimes that means pushing
back on potential changes.


#### Commit Hygiene 

It can be seen as trivial, but you can ask the PR author to break apart their
PR into smaller chunks, or change a commit message to be more informative. They
are the _"permanent record"_ of the change and should accurately describe both
what and why it is being done.


### Be clear on progress and time

Be upfront with the PR author about where the state of their PR is and what
needs to be completed for it to be accepted.

No one likes it if their PR misses a release, but it is a fact of life. Try and
be upfront about it. Don't push a PR through out of guilt or deadlines. Remember,
you are a steward of the codebase.


### Checking out a PR

If a PR is too complex to review through the GitHub UI, you can pull it down
locally to evaluate. You can do so using the following command:

```
git fetch origin pull/<PR ID>/head:<BRANCHNAME>
git checkout <BRANCHNAME>
```

**Example:**
```
git fetch upstream pull/1245/head:foo
git checkout foo
```


## Additional Resources

- [Keeping the Bar High - How to be a bad ass Code Reviewer, Tim Hockin] - A
  presentation by Tim from the Kubernetes Contributor Summit in San Diego. It is
  largely what this document is based off of.
- [Kubernetes Code Reviewing with Tim Hockin] - Some notes and tips from Tim on
  how to effective at reviewing Kubernetes Code.
- [Live API Review, Jordan Liggitt] - A presentation by Jordan from the Kubernetes
  Contributor Summit in San Diego covering the sort of things they look for when
  performing an API review for a new feature.
- [The Gentle Art of Patch Review, Sage Sharp] - A blog post from Sage Sharp on
  some methods of approaching PR review



[squashing]: ./github-workflow.md#squash-commits
[KEEP THE SPACE SHUTTLE FLYING]: https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/volume/persistentvolume/pv_controller.go#L57-L117
[commit message guidelines]: ./pull-requests.md#7-commit-message-guidelines
[GitHub Status]: https://help.github.com/en/github/setting-up-and-managing-your-github-profile/personalizing-your-profile#setting-a-status
[Blunderbuss]: https://git.k8s.io/test-infra/prow/plugins/approve/approvers/README.md#blunderbuss-and-reviewers
[emeritus_approver]: ./owners.md#emeritus
[Keeping the Bar High - How to be a bad ass Code Reviewer, Tim Hockin]: https://www.youtube.com/watch?v=OZVv7-o8i40
[Kubernetes Code Reviewing with Tim Hockin]: https://docs.google.com/document/d/15y8nIgWMzptHcYIeqf4vLJPttE3Fj_ht4I6Nj4ghDLA/edit#heading=h.3dchnigrxf5y
[Live API Review, Jordan Liggitt]: https://www.youtube.com/watch?v=faRARV3C7Fk
[The Gentle Art of Patch Review, Sage Sharp]: https://sage.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/

