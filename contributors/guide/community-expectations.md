# Community Expectations

Kubernetes is a community project. Consequently, it is wholly dependent on
its community to provide a productive, friendly and collaborative environment.

The first and foremost goal of the Kubernetes community to develop orchestration
technology that radically simplifies the process of creating reliable
distributed systems. However a second, equally important goal is the creation
of a community that fosters easy, agile development of such orchestration
systems.

We therefore describe the expectations for
members of the Kubernetes community.  This document is intended to be a living one
that evolves as the community evolves via the same PR and code review process
that shapes the rest of the project.  It currently covers the expectations
of conduct that govern all members of the community as well as the expectations
around code review that govern all active contributors to Kubernetes.

## Code review

As a community we believe in the value of code review for all contributions.
Code review increases both the quality and readability of our codebase, which
in turn produces high quality software.

However, the code review process can also introduce latency for contributors
and additional work for reviewers that can frustrate both parties.

Consequently, as a community we expect that all active participants in the
community will also be active reviewers. The 
[community membership](../community-membership.md) outlines the responsibilities
of the different contributor roles. 

All changes must be code reviewed. For non-maintainers this is obvious, since
you can't commit anyway. But even for maintainers, we want all changes to get at
least one review, preferably (for non-trivial changes obligatorily) from someone
who knows the areas the change touches. For non-trivial changes we may want two
reviewers. The primary reviewer will make this decision and nominate a second
reviewer, if needed. Except for trivial changes, PRs should not be committed
until relevant parties (e.g. owners of the subsystem affected by the PR) have
had a reasonable chance to look at PR in their local business hours.

Most PRs will find reviewers organically. If a maintainer intends to be the
primary reviewer of a PR they should set themselves as the assignee on GitHub
and say so in a reply to the PR. Only the primary reviewer of a change should
actually do the merge, except in rare cases (e.g. they are unavailable in a
reasonable timeframe).

If a PR has gone 2 work days without an owner emerging, please poke the PR
thread and ask for a reviewer to be assigned.

Except for rare cases, such as trivial changes (e.g. typos, comments) or
emergencies (e.g. broken builds), maintainers should not merge their own
changes.

Expect reviewers to request that you avoid [common go style
mistakes](https://github.com/golang/go/wiki/CodeReviewComments) in your PRs.

## Expectations of reviewers: Review comments

Because reviewers are often the first points of contact between new members of
the community and can significantly impact the first impression of the
Kubernetes community, reviewers are especially important in shaping the
Kubernetes community.  Reviewers are highly encouraged to review the
[code of conduct](../../governance.md#code-of-conduct) and are strongly 
encouraged to go above and beyond the code of conduct to promote a collaborative, 
respectful Kubernetes community.

## Expectations of reviewers: Review latency

Reviewers are expected to respond in a timely fashion to PRs that are assigned
to them.  Reviewers are expected to respond to an *active* PRs with reasonable
latency, and if reviewers fail to respond, those PRs may be assigned to other
reviewers.

*Active* PRs are considered those which have a proper CLA (`cla:yes`) label
and do not need rebase to be merged.  PRs that do not have a proper CLA, or
require a rebase are not considered active PRs.

## Thanks

Many thanks in advance to everyone who contributes their time and effort to
making Kubernetes both a successful system as well as a successful community.
The strength of our software shines in the strengths of each individual
community member.  Thanks!
