# Guidelines for contributing API changes

## Overview

Modifications to the core Kubernetes API have a material impact on the Kubernetes project, and
undergo thorough review and discussion before they are accepted.  This document outlines the
review process and what is expected of API contributors.

## Process overview

*Proposal process*

All API changes require a formal proposal to be made, reviewed and accepted.  The proposal will be
considered accepted once the proposal PR has been merged.

*Feature process*

All API changes require an issue to be filed in the [kubernetes/features] repo.  The
issue will be used by release coordinators to track which features are slated for a
release milestone.

## Changes that qualify as API changes

All changes to existing or new APIs are subject to the process outlined in this document.

The following qualify as API changes:

- Adding, removing, or modifying API types (e.g. through types.go, or by defining a new annotation)
- Changing the version of an existing type (e.g. v1alpha1 to v1beta1)
- Adding, removing, or modifying subresources for existing types. (e.g. /scale)
- Material changes to how fields on existing types are interpreted / used - including, but not limited to
  validation and defaulting of the fields.

The following on their own do not qualify as API changes, but should follow the process defined by the related
[special interest groups].

- Adding or changing a command or flag to kubectl: See [contributing to SIG cli]
- Building new command line tools: See [contributing to SIG cli]
- Adding or changing a flag to kubelet: sig-node
- Adding or changing a flag to apiserver: sig-api-machinery
- Refactoring code
- Building extensions

## Life of an API change

1. Start a discussion
2. Write a short description
3. Write a detailed design proposal in the [kubernetes/community] repo
4. Schedule a design review
5. File an issue in the [kubernetes/features] repo that links to the proposal
6. Implement the proposal in the [kubernetes/kubernetes] repo
7. Write user facing documentation on [kubernetes/kubernetes.github.io]

### Start a discussion.

Expected time for decision: 1-4 weeks (depending on the SIG and feature)

Start a discussion with the relevant [special interest groups] that will [own] the changes.
Depending on the SIG, this may be done via SIG meetings, email groups, slack channel, etc.

### Write a short description

Expected time for decision: 2 weeks

Write a short description of the API changes by copying the [design summary template]
into a new issue, and [@mention] @kubernetes/api-reviewers and the owning SIG reviewers
group.

Determine whether @kubernetes/api-reviewers is willing to accept a design proposal
for your feature.  The answer could be "no, we don't want to solve this problem this way."
or "no, this is something we want, but cannot commit the resources to reviewing it this release."

### Write a detailed design proposal

Expected time for initial feedback: 3-6 weeks

Write a detailed design proposal PR using the [design proposal PR template].  [@mention]
@kubernetes/api-reviewers and the appropriate SIGs on the PR and reference the original
issue in the PR description.

### Schedule a design review

Expected time for next slot: 4-8 weeks

Schedule a time for your design to be reviewed.

See the current list of [design approvers](https://github.com/orgs/kubernetes/teams/api-approvers)

**Note**: these are the current reviewers, we are still figuring out how to
add/remove people to this list.  For more information, see the
[governance discussion](https://groups.google.com/forum/#!topic/kubernetes-dev/4e8WOnMvZC0)

### Implement the proposal

Once the proposal has been accepted an merged, you can begin implementing the solution.

## How to escalate

Escalation should be done through the owning SIG.  If you need help getting attention, reach out to your
SIG.

## API change considerations

First read the [API changes] document for background on how to make changes in the API.

When writing a design proposal, you must consider the following:

### API versions

See [API versions] for details on the semantic meaning of the API versions.  Changes to how
existing fields are validated, defaulted, or interpreted requires an API version change.

### API groups

It is important to pick the correct API group for your API.  This will ensure that it is discoverable
by users and is maintained in concert with related APIs.  Current API groups:

| Group  | Description |
| ------ | ----------- |
| abac   |             |
| apps |   |
| authentication |   |
| authorization |   |
| autoscaling |   |
| batch |   |
| certificates |   |
| imagepolicy |   |
| policy |   |
| rbac |   |
| settings  |   |
| storage |   |


## Related documents

- [API changes]
- [API conventions]
- [OARP roles model](https://stumblingabout.com/tag/oarp/)
- [RACI roles model](http://www.valuebasedmanagement.net/methods_raci.html)

[API versions]: https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#alpha-beta-and-stable-versions
[@mention]: https://help.github.com/articles/basic-writing-and-formatting-syntax/#mentioning-users-and-teams
[own]: https://github.com/kubernetes/community/blob/master/contributors/sig-ownership.md
[special interest groups]: https://github.com/kubernetes/community/blob/master/README.md#special-interest-groups-sig
[design proposal PR template]: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-proposal-design-template.md
[design summary template]: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-proposal-issue-template.md
[API changes]: https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md
[contributing to SIG cli]: https://github.com/kubernetes/community/blob/master/sig-cli/contributing.md
[kubernetes/kubernetes]: https://github.com/kubernetes/kubernetes/
[kubernetes/features]: https://github.com/kubernetes/features/
[kubernetes/kubernetes.github.io]: https://github.com/kubernetes/kubernetes.github.io/
[kubernetes/community]: https://github.com/kubernetes/community/
[API changes]: https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md
[API conventions]: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md