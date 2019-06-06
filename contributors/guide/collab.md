---
title: "Collaborative Development"
weight: 1
slug: "collab"
---

## Code reviews

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

## Assigned reviews

Maintainers can assign reviews to other maintainers, when appropriate. The
assignee becomes the shepherd for that PR and is responsible for merging the PR
once they are satisfied with it or else closing it. The assignee might request
reviews from non-maintainers.
