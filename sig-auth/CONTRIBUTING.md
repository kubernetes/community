# Contributing to SIG Auth

Welcome to contributing to SIG Auth.

If you haven't seen them already, the Kubernetes project has:

- A [Contributor Guide][contrib-guide] - some
  [kubernetes/kubernetes]-specific content, but lots of info for the
  entire project
- A [Contributor Cheat Sheet][contrib-cheatsheet] - lots of resources
  and handy links

SIG Auth has multiple areas you can contribute to. Those contributions
can be in the form of code, documentation, support being involved in
mailing list discussions, attending meetings, and more. This guide
describes different major functional areas SIG Apps is involved in,
provides an overview of the areas, and gives pointers on getting more
involved in each area. Consider this a launching point or the start of
a [choose your own
adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure)
for SIG Auth.

## Workflow

Just like rest of the Kubernetes project, we also use the same PR and
review based workflow. Which means [use of the CNCF CLA][cla], [code
review by reviewers and approvers listed in OWNERS files][owners], and
tests that automatically exercise code or enforce conventions.

Some other points to keep in mind while contributing to this repo:

- For large code changes, please write a design doc or [KEP] and get
  signoff from SIG Auth before trying to land code. If you're not sure
  what a KEP should look like,
  [kuberenetes/enhancements/keps/sig-auth] has some examples. If
  you're not sure what "large" means, [come ask us](#contact)
- We find it polite to apply the `do-not-merge/hold` label via the `/hold`
  command when in doubt about whether/when a PR should merge. If we don't
  explain why we're adding a `/hold`, this is usually because the PR will
  cause changes to be deployed on merge, and we want the person responsible
  for deploying and monitoring the changes to decide when to `/hold cancel`.
  In some cases this is the PR author, in other cases this may be whomever
  is on-call for prow.k8.io.  If you are unsure which, please ask.
- Many of us use [gubernator.k8s.io/pr] to keep track of which PRs require
  our attention. Use of the [`/cc @person`][command-cc] or
  [`/assign @person`][command-assign] commands is the most effective way to
  put PRs on our radar. You can also [contact us](#contact) on slack.


## Issue Triaging

Issues are ideally labeled with:

- milestone: during which release cycle do we plan on working on this issue
- `sig/foo`: which SIG owns this work
- `area/foo`: which subproject or code is this issue related to
- `kind/foo`: which kind of work is this issue describing
- `priority/foo`: how important is this issue

For example, an issue related to cleaning up and consolidating (`kind/cleanup`)
release-related (`sig/release`) jobs and dashboards (`area/config`) for the
v1.16 cycle (`milestone: v1.16`) that may not get completed by the end of the
cycle if more important or more urgent arises (`priority/important-longterm`).

We try to have a non-stale pool of issues that are available for new
contributors who want to help out but aren't sure what to work on or where to
get started:

- [`label:"good first issue"`][good-first-issue] - triaged per [Good First Issue docs][good-first-issue-docs]
- [`label:"help wanted"`][help-wanted] - triaged per [Help Wanted
  docs][help-wanted-docs]
  
## Guides

If you're not sure where to contribute or what any of these mean, please see
[/README.md] for a brief description of the various codebases in this repo.
