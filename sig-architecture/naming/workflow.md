# Recommendations Workflow

This document is a workflow for proposing language recommendations to
WG Naming.

## Make a proposal

Begin by sending a message outlining the proposed change to
the [WG Naming mailing list][mailing-list].

While WG Naming also discusses proposals in [#wg-naming][Slack], Slack is a lossy
communication forum which is much harder to review than the mailing list.

Proposals should include information required to file a WG Naming
Architecture Decision Record (ADR):

- The recommendation to be made
- Brief reasoning behind the change
- The context in which the existing term might be used
- Alternative recommendations
- A reasonable guess at the work required to make the change

For more information about the WG Naming ADR, review the [template].

Following a mailing list proposal, contributors can continue discussion both on Slack
and at WG Naming meetings.

### Filing a recommendation ADR

If the working group's discussion determines that the recommendation is
reasonable and in line with our [framework] for language evaluation,
WG Naming leads formalize the recommendation with an ADR.

For more information about the WG Naming ADR, review the [template].

An ADR must include:

- Groups responsible for implementing the change
- The scope of the change in the Kubernetes project, as well as downstream
  implications

ADRs must be opened with a `/hold` to give an opportunity to seek approval
with the governance groups that with be responsible for implementation.

After opening an ADR pull request, WG Naming leads
should:

- Reply to any mailing list threads about the recommendation with a link to the
  newly-opened PR, and CC any stakeholder groups
- Place the recommendation on the next meeting's agenda for review

## Approval

_This approval process is still under discussion, so here we will list out some
frequently-asked questions from our discussions thus far._

### What if a recommendation requires a KEP?

ADRs should remain on hold until scoped area agrees with the direction.

### What do we do when stakeholders disagree with a recommendation?

Again, do not merge a recommendation until code owners from the scoped area
agree to it.

### General guidance

- SIG Architecture records decisions to "...not make the mistakes we made in
  the past"
- Don’t block recording a recommendation on a plan to remediate all existing
  uses; once the direction is agreed on by the code/content owners from the
  scoped area, a recorded recommendation has value in guiding new/future work

### Requirements

- ADR is on hold until approved by scoped areas (e.g., SIG Architecture, SIG
  Docs)
- Steering is tagged on the ADR for approval
- WG Naming lead establishes a Steering review/approval period with a lazy
  consensus deadline of 3-5 business days
- WG Naming lead releases the hold and merges ADR

## Implementation

- WG Naming leads record accepted recommendation in a canonical location (TBD)
  (for example, a style guide)
- Areas in scope are now responsible for implementation

[framework]: language-evaluation-framework.md
[mailing-list]: https://groups.google.com/forum/#!forum/kubernetes-wg-naming
[slack]: https://kubernetes.slack.com/messages/wg-naming
[template]: ./recommendations/template.md
