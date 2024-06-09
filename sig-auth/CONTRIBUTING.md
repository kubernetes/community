# Contributing to SIG Auth

Welcome to contributing to SIG Auth.

If you haven't seen them already, the Kubernetes project has:

- A [Contributor Guide](https://git.k8s.io/community/contributors/guide) - some
  [kubernetes/kubernetes]-specific content, but lots of info for the
  entire project
- A [Contributor Cheat
  Sheet](https://github.com/kubernetes/community/blob/master/contributors/guide/contributor-cheatsheet) -
  lots of resources and handy links

SIG Auth has multiple areas you can contribute to. Those contributions
can be in the form of code, documentation, support being involved in
mailing list discussions, attending meetings, and more. This guide
describes different major functional areas SIG Auth is involved in,
provides an overview of the areas, and gives pointers on getting more
involved in each area. Consider this a launching point or the start of
a [choose your own
adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure)
for SIG Auth.

## Workflow

Just like rest of the Kubernetes project, we also use the same PR and
review based workflow. Which means [use of the CNCF
CLA](https://github.com/kubernetes/community/blob/master/CLA.md),
[code review by reviewers and approvers listed in OWNERS
files](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md),
and tests that automatically exercise code or enforce conventions.

## Issue Triaging

Issues are ideally labeled with:

- milestone: during which release cycle do we plan on working on this issue
- `sig/foo`: which SIG owns this work
- `area/foo`: which subproject or code is this issue related to
- `kind/foo`: which kind of work is this issue describing
- `priority/foo`: how important is this issue

For example, an issue related to cleaning up and consolidating
(`kind/cleanup`) release-related (`sig/release`) jobs and dashboards
(`area/config`) for the v1.16 cycle (`milestone: v1.16`) that may not
get completed by the end of the cycle if more important or more urgent
work arises (`priority/important-longterm`).

We try to have a non-stale pool of issues that are available for new
contributors who want to help out but aren't sure what to work on or where to
get started:

- [`label:"good first issue"`][good-first-issue] - triaged per
  [good-first-issue](https://github.com/issues?q=repo%3Akubernetes%2Fkubernetes+is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22+label%3A%22sig%2Fauth%22)
- [good-first-issue-docs](https://git.k8s.io/community/contributors/guide/help-wanted.md#good-first-issue)
- [`label:"help wanted"`][help-wanted] - triaged per
  [help-wanted](https://github.com/issues?q=repo%3Akubernetes%2Fkubernetes+is%3Aissue+is%3Aopen+label%3A%22help+wanted%22++label%3A%22sig%2Fauth%22)
- [help-wanted-docs](https://git.k8s.io/community/contributors/guide/help-wanted.md#help-wanted)
  
## Guides

If you're not sure where to contribute or what any of these mean,
please see
[README.md](https://github.com/kubernetes/community/blob/master/README.md)
for a brief description of the various codebases in this repo.
