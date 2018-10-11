# SIG Governance Requirements

## Goals

This document outlines the recommendations and requirements for defining SIG and subproject governance.

This doc uses [rfc2119](https://www.ietf.org/rfc/rfc2119.txt) to indicate keyword requirement levels.
Sub elements of a list inherit the requirements of the parent by default unless overridden.

## Checklist

Following is the checklist of items that should be considered as part of defining governance for
any subarea of the Kubernetes project.

### Roles

- *MUST* enumerate any roles within the SIG and the responsibilities of each
- *MUST* define process for changing the membership of roles
  - When and how new members are chosen / added to each role
  - When and how existing members are retired from each role
- *SHOULD* define restrictions / requirements for membership of roles
- *MAY* define target staffing numbers of roles

### Organizational management

- *MUST* define when and how collaboration between members of the SIG is organized
  - *SHOULD* define how periodic video conference meetings are arranged and run
  - *SHOULD* define how conference / summit sessions are arranged
  - *MAY* define periodic office hours on slack or video conference

- *MAY* define process for new community members to contribute to the area
  - e.g. read a contributing guide, show up at SIG meeting, message the google group

- *MUST* define how subprojects are managed
  - When and how new subprojects are created
  - Subprojects *MUST* define roles (and membership) within subprojects

### Project management

The following checklist applies to both SIGs and subprojects of SIGs as appropriate:

- *MUST* define how milestones / releases are set
  - How target dates for milestones / releases are proposed and accepted
  - What priorities are targeted for milestones
  - The process for publishing a release

- *SHOULD* define how priorities / commitments are managed
  - How priorities are determined
  - How priorities are staffed

### Technical processes

All technical assets *MUST* be owned by exactly 1 SIG subproject.  The following checklist applies to subprojects:

- *MUST* define how technical decisions are communicated and made within the SIG or project
  - Process for proposal, where and how it is published and discussed, when and how a decision is made
    (e.g. [KEP] process)
  - Who are the decision makers on proposals (e.g. anyone in the world can block, just reviewers on the PR,
    just approvers in OWNERs, etc)
  - How disagreements are resolved within the area (e.g. discussion, fallback on voting, escalation, etc)
  - How and when disagreements may be escalated
  - *SHOULD* define expectations and recommendations for proposal process (e.g. escalate if not progress towards
    resolution in 2 weeks)
  - *SHOULD* define a level of commitment for decisions that have gone through the formal process
    (e.g. when is a decision revisited or reversed)

- *MUST* define how technical assets of project remain healthy and can be released
  - Publicly published signals used to determine if code is in a healthy and releasable state
  - Commitment and process to *only* release when signals say code is releasable
  - Commitment and process to ensure assets are in a releasable state for milestones / releases
    coordinated across multiple areas / subprojects (e.g. the Kubernetes OSS release)
  - *SHOULD* define target metrics for health signal (e.g. broken tests fixed within N days)
  - *SHOULD* define process for meeting target metrics (e.g. all tests run as presubmit, build cop, etc)

[lazy-consensus]: http://en.osswiki.info/concepts/lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[warnocks-dilemma]: http://communitymgt.wikia.com/wiki/Warnock%27s_Dilemma
[slo]: https://en.wikipedia.org/wiki/Service_level_objective
[steering-committee]: https://github.com/kubernetes/steering#contact
[business-operations]: http://www.businessdictionary.com/definition/business-operation.html
[KEP]: https://kubernetes.io/docs/imported/community/keps/
