# SIG Charter Guide

All Kubernetes SIGs must define a charter defining the scope and governance of the SIG.

- The scope must define what areas the SIG is responsible for directing and maintaining.
- The governance must outline the responsibilities within the SIG as well as the roles
  owning those responsibilities.

## Steps to create a SIG charter

1. Copy [the template][Short Template] into a new file under community/sig-*YOURSIG*/charter.md ([sig-architecture example])
2. Read the [Recommendations and requirements] so you have context for the template
3. Fill out the template for your SIG
4. Update [sigs.yaml] with the individuals holding the roles as defined in the template.
5. Add subprojects owned by your SIG in the [sigs.yaml]
5. Create a pull request with a draft of your charter.md and sigs.yaml changes.  Communicate it within your SIG
   and get feedback as needed.
6. Send the SIG Charter out for review to steering@kubernetes.io.  Include the subject "SIG Charter Proposal: YOURSIG"
   and a link to the PR in the body.
7. Typically expect feedback within a week of sending your draft.  Expect longer time if it falls over an
   event such as KubeCon/CloudNativeCon or holidays.  Make any necessary changes.
8. Once accepted, the steering committee will ratify the PR by merging it.

## Steps to update an existing SIG charter

- For significant changes, or any changes that could impact other SIGs, such as the scope, create a
  PR and send it to the steering committee for review with the subject: "SIG Charter Update: YOURSIG"
- For minor updates to that only impact issues or areas within the scope of the SIG the SIG Chairs should
  facilitate the change.

## SIG Charter approval process

When introducing a SIG charter or modification of a charter the following process should be used.
As part of this we will define roles for the [OARP] process (Owners, Approvers, Reviewers, Participants)

- Identify a small set of Owners from the SIG to drive the changes.
  Most typically this will be the SIG chairs.
- Work with the rest of the SIG in question (Reviewers) to craft the changes.
  Make sure to keep the SIG in the loop as discussions progress with the Steering Committee (next step).
  Including the SIG mailing list in communications with the steering committee would work for this.
- Work with the steering committee (Approvers) to gain approval.
  This can simply be submitting a PR and sending mail to [steering@kubernetes.io].
  If more substantial changes are desired it is advisable to socialize those before drafting a PR.
    - The steering committee will be looking to ensure the scope of the SIG as represented in the charter is reasonable (and within the scope of Kubernetes) and that processes are fair.
- For large changes alert the rest of the Kubernetes community (Participants) as the scope of the changes becomes clear.
  Sending mail to [kubernetes-dev@googlegroups.com] and/or announcing at the community meeting are a good ways to do this.

If there are questions about this process please reach out to the steering committee at [steering@kubernetes.io].

## How to use the templates

SIGs should use [the template][Short Template] as a starting point. This document links to the recommended [SIG Governance][sig-governance] but SIGs may optionally record deviations from these defaults in their charter.


## Goals

The primary goal of the charters is to define the scope of the SIG within Kubernetes and how the SIG leaders exercise ownership of these areas by taking care of their responsibilities. A majority of the effort should be spent on these concerns.

## FAQ

See [frequently asked questions]

[OARP]: https://stumblingabout.com/tag/oarp/
[Recommendations and requirements]: sig-governance-requirements.md
[sig-governance]: sig-governance.md
[Short Template]: sig-charter-template.md
[frequently asked questions]: FAQ.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[sig-architecture example]: ../../sig-architecture/charter.md
[steering@kubernetes.io]: mailto:steering@kubernetes.io
[kubernetes-dev@googlegroups.com]: mailto:kubernetes-dev@googlegroups.com
