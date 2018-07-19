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
   event such as Kubecon or holidays.  Make any necessary changes.
8. Once accepted, the steering committee will ratify the PR by merging it.

## Steps to update an existing SIG charter

- For significant changes, or any changes that could impact other SIGs, such as the scope, create a
  PR and send it to the steering committee for review with the subject: "SIG Charter Update: YOURSIG"
- For minor updates to that only impact issues or areas within the scope of the SIG the SIG Chairs should
  facilitate the change.

## How to use the templates

SIGs should use [the template][Short Template] as a starting point. This document links to the recommended [SIG Governance][sig-governance] but SIGs may optionally record deviations from these defaults in their charter.


## Goals

The primary goal of the charters is to define the scope of the SIG within Kubernetes and how the SIG leaders exercise ownership of these areas by taking care of their responsibilities. A majority of the effort should be spent on these concerns.

## FAQ

See [frequently asked questions]

[Recommendations and requirements]: sig-governance-requirements.md
[sig-governance]: sig-governance.md
[Short Template]: sig-charter-template.md
[frequently asked questions]: FAQ.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[sig-architecture example]: ../../sig-architecture/charter.md
