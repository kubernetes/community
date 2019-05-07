# Targeting Features, Issues and PRs to Release Milestones

This document is focused on Kubernetes developers and contributors
who need to create a feature, issue, or pull request which targets a specific
release milestone.

-   [TL;DR](#tldr)
-   [Definitions](#definitions)
-   [The Release Cycle](#the-release-cycle)
-   [Removal Of Items From The Milestone](#removal-of-items-from-the-milestone)
-   [Adding An Item To The Milestone](#adding-an-item-to-the-milestone)
    -   [Milestone Maintainers](#milestone-maintainers)
    -   [Feature additions](#feature-additions)
    -   [Issue additions](#issue-additions)
    -   [PR Additions](#pr-additions)
-   [Other Required Labels](#other-required-labels)
    -   [SIG Owner Label](#sig-owner-label)
    -   [Priority Label](#priority-label)
    -   [Issue Kind Label](#issue-kind-label)

The process for shepherding features, issues, and pull requests
into a Kubernetes release spans multiple stakeholders:
* the feature, issue, or pull request owner
* SIG leadership
* the release team

Information on workflows and interactions are described below.

As the owner of a feature, issue, or pull request (PR), it is your
responsibility to ensure release milestone requirements are met.
Automation and the release team will be in contact with you if
updates are required, but inaction can result in your work being
removed from the milestone.  Additional requirements exist when the
target milestone is a prior release (see [cherry pick
process](cherry-picks.md) for more information).

## TL;DR

If you want your PR to get merged, it needs the following required labels and milestones, represented here by the Prow /commands it would take to add them:
<table>
<tr>
<td></td>
<td>Normal Dev</td>
<td>Code Freeze</td>
<td>Post-Release</td>
</tr>
<tr>
<td></td>
<td>Weeks 1-8</td>
<td>Weeks 9-11</td>
<td>Weeks 11+</td>
</tr>
<tr>
<td>Required Labels</td>
<td>
<ul>
<!--Weeks 1-8-->
<li>/sig {name}</li>
<li>/kind {type}</li>
<li>/lgtm</li>
<li>/approved</li>
</ul>
</td>
<td>
<ul>
<!--Weeks 9-11-->
<li>/milestone {v1.y}</li>
<li>/sig {name}</li>
<li>/kind {bug, failing-test}</li>
<li>/lgtm</li>
<li>/approved</li>
</ul>
</td>
<td>
<!--Weeks 11+-->
Return to 'Normal Dev' phase requirements:
<ul>
<li>/sig {name}</li>
<li>/kind {type}</li>
<li>/lgtm</li>
<li>/approved</li>
</ul>

Merges into the 1.y branch are now [via cherrypicks](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-release/cherry-picks.md), approved by release branch manager.
</td>
<td>
<ul>
</td>
</tr>
</table>

In the past there was a requirement for a milestone targeted pull
request to have an associated GitHub issue opened, but this is no
longer the case.  Features are effectively GitHub issues or
[KEPs](https://git.k8s.io/community/keps)
which lead to subsequent PRs. The general labeling process should
be consistent across artifact types.

---

## Definitions

- *issue owners*: Creator, assignees, and user who moved the issue into a release milestone.
- *release team*: Each Kubernetes release has a team doing project
  management tasks described
  [here](https://git.k8s.io/sig-release/release-team/README.md).  The
  contact info for the team associated with any given release can be
  found [here](https://git.k8s.io/sig-release/releases/).
- *Y days*: Refers to business days (using the location local to the release-manager M-F).
- *feature*: see "[Is My Thing a Feature?](http://git.k8s.io/features/README.md#is-my-thing-a-feature)
- *release milestone*: semantic version string or [GitHub milestone](https://help.github.com/articles/associating-milestones-with-issues-and-pull-requests/) referring to a release MAJOR.MINOR vX.Y version.  See also [release versioning](http://git.k8s.io/community/contributors/design-proposals/release/versioning.md)
- *release branch*: Git branch "release-X.Y" created for the vX.Y milestone.  Created at the time of the vX.Y-beta.0 release and maintained after the release for approximately 9 months with vX.Y.Z patch releases.

## The Release Cycle

![Image of one Kubernetes release cycle](release-cycle.png)

Kubernetes releases currently happen four times per year.  The release
process can be thought of as having three main phases:
* Feature Definition
* Implementation
* Stabilization

But in reality this is an open source and agile project, with feature
planning and implementation happening at all times.  Given the
project scale and globally distributed developer base, it is critical
to project velocity to not rely on a trailing stabilization phase and
rather have continuous integration testing which ensures the
project is always stable so that individual commits can be
flagged as having broken something.

With ongoing feature definition through the year, some set of items
will bubble up as targeting a given release.  The **enhancement freeze**
starts ~4 weeks into release cycle.  By this point all intended
feature work for the given release has been defined in suitable
planning artifacts in conjunction with the Release Team's [enhancements
lead](https://git.k8s.io/sig-release/release-team/role-handbooks/enhancements/README.md).

Implementation and bugfixing is ongoing across the cycle, but
culminates in a code freeze period:
* The **code freeze** starts in week ~10 and continues for ~2 weeks.
  Only critical bug fixes are accepted into the release codebase.

There are approximately two weeks following code freeze, and preceding
release, during which all remaining critical issues must be resolved
before release.  This also gives time for documentation finalization.

When the code base is sufficiently stable, the master branch re-opens
for general development and work begins there for the next release
milestone. Any remaining modifications for the current release are cherry
picked from master back to the release branch.  The release is built from
the release branch.

Following release, the [Release Branch
Manager](https://git.k8s.io/sig-release/release-team/role-handbooks/branch-manager/README.md)
cherry picks additional critical fixes from the master branch for
a period of around 9 months, leaving an overlap of three release
versions forward support.  Thus, each release is part of a broader
Kubernetes lifecycle:

![Image of Kubernetes release lifecycle spanning three releases](release-lifecycle.png)

## Removal Of Items From The Milestone

Before getting too far into the process for adding an item to the
milestone, please note:

Members of the Release Team may remove Issues from the milestone
if they or the responsible SIG determine that the issue is not
actually blocking the release and is unlikely to be resolved in a
timely fashion.

Members of the Release Team may remove PRs from the milestone for
any of the following, or similar, reasons:

* PR is potentially de-stabilizing and is not needed to resolve a blocking issue;
* PR is a new, late feature PR and has not gone through the features process or the exception process;
* There is no responsible SIG willing to take ownership of the PR and resolve any follow-up issues with it;
* PR is not correctly labelled;
* Work has visibly halted on the PR and delivery dates are uncertain or late.

While members of the Release Team will help with labelling and
contacting SIG(s), it is the responsibility of the submitter to
categorize PRs, and to secure support from the relevant SIG to
guarantee that any breakage caused by the PR will be rapidly resolved.

Where additional action is required, an attempt at human to human
escalation will be made by the release team through the following
channels:

- Comment in GitHub mentioning the SIG team and SIG members as appropriate for the issue type
- Emailing the SIG mailing list
  - bootstrapped with group email addresses from the [community sig list](/sig-list.md)
  - optionally also directly addressing SIG leadership or other SIG members
- Messaging the SIG's Slack channel
  - bootstrapped with the slackchannel and SIG leadership from the [community sig list](/sig-list.md)
  - optionally directly "@" mentioning SIG leadership or others by handle

## Adding An Item To The Milestone

### Milestone Maintainers

The members of the GitHub [“kubernetes-milestone-maintainers”
team](https://github.com/orgs/kubernetes/teams/kubernetes-milestone-maintainers/members)
are entrusted with the responsibility of specifying the release milestone on
GitHub artifacts.  This group is [maintained by
SIG-Release](https://git.k8s.io/sig-release/release-team/README.md#milestone-maintainers)
and has representation from the various SIGs' leadership.

### Feature additions

Feature planning and definition takes many forms today, but a typical
example might be a large piece of work described in a
[KEP](https://git.k8s.io/community/keps), with associated
task issues in GitHub.  When the plan has reached an implementable state and
work is underway, the feature or parts thereof are targeted for an upcoming
milestone by creating GitHub issues and marking them with the Prow "/milestone"
command.

For the first ~4 weeks into the release cycle, the release team's
Enhancements Lead will interact with SIGs and feature owners via GitHub,
Slack, and SIG meetings to capture all required planning artifacts.

If you have a feature to target for an upcoming release milestone, begin a
conversation with your SIG leadership and with that release's Enhancements
Lead.

### Issue additions

Issues are marked as targeting a milestone via the Prow
"/milestone" command.

The release team's [Bug Triage
Lead](https://git.k8s.io/sig-release/release-team/role-handbooks/bug-triage/README.md) and overall community watch
incoming issues and triage them, as described in the contributor
guide section on [issue triage](/contributors/guide/issue-triage.md).

Marking issues with the milestone provides the community better
visibility regarding when an issue was observed and by when the community
feels it must be resolved.  During code freeze, to merge a PR it is required
that a release milestone is set.

An open issue is no longer required for a PR, but open issues and
associated PRs should have synchronized labels.  For example a high
priority bug issue might not have its associated PR merged if the PR is
only marked as lower priority.

### PR Additions

PRs are marked as targeting a milestone via the Prow
"/milestone" command.

This is a blocking requirement during code freeze as described above.

## Other Required Labels

*Note* [Here is the list of labels and their use and purpose.](https://git.k8s.io/test-infra/label_sync/labels.md#labels-that-apply-to-all-repos-for-both-issues-and-prs)

### SIG Owner Label

The SIG owner label defines the SIG to which we escalate if a
milestone issue is languishing or needs additional attention.  If
there are no updates after escalation, the issue may be automatically
removed from the milestone.

These are added with the Prow "/sig" command. For example to add
the label indicating SIG Storage is responsible, comment with `/sig
storage`.

### Priority Label

Priority labels are used to determine an escalation path before
moving issues out of the release milestone.  They are also used to
determine whether or not a release should be blocked on the resolution
of the issue.

- `priority/critical-urgent`: Never automatically move out of a release milestone; continually escalate to contributor and SIG through all available channels.
  - considered a release blocking issue
  - code freeze: issue owner update frequency: daily
  - would require a patch release if left undiscovered until after the minor release.
- `priority/important-soon`: Escalate to the issue owners and SIG owner; move out of milestone after several unsuccessful escalation attempts.
  - not considered a release blocking issue
  - would not require a patch release
  - will automatically be moved out of the release milestone at code freeze after a 4 day grace period
- `priority/important-longterm`: Escalate to the issue owners; move out of the milestone after 1 attempt.
  - even less urgent / critical than `priority/important-soon`
  - moved out of milestone more aggressively than `priority/important-soon`

### Issue/PR Kind Label

The issue kind is used to help identify the types of changes going
into the release over time.  This may allow the release team to
develop a better understanding of what sorts of issues we would
miss with a faster release cadence.

For release targeted issues, including pull requests, one of the following 
issue kind labels must be set:

- `kind/api-change`: Adds, removes, or changes an API
- `kind/bug`: Fixes a newly discovered bug.
- `kind/cleanup`: Adding tests, refactoring, fixing old bugs.
- `kind/design`: Related to design
- `kind/documentation`: Adds documentation
- `kind/failing-test`: CI test case is failing consistently.
- `kind/feature`: New functionality.
- `kind/flake`: CI test case is showing intermittent failures.
