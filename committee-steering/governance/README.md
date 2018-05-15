# SIG Charter Guide

All Kubernetes SIGs must define a charter defining the scope and governance of the SIG.

- The scope must define what areas the SIG is responsible for directing and maintaining.
- The governance must outline the responsibilities within the SIG as well as the roles
  owning those responsibilities.

## Steps to create a SIG charter

1. Copy the template into a new file under community/sig-*YOURSIG*/charter.md ([sig-architecture example])
2. Read the [Recommendations and requirements] so you have context for the template
3. Customize your copy of the template for your SIG.  Feel free to make adjustments as needed.
4. Update [sigs.yaml] with the individuals holding the roles as defined in the template.
5. Add subprojects owned by your SIG to the [sigs.yaml]
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

When developing or modifying a SIG governance doc, the intention is for SIGs to use the templates (*under development*)
as a common set of options SIGs may choose to incorporate into their own governance structure.  It is recommended that
SIGs start by looking at the [Recommendations and requirements] for SIG governance docs and consider what structure
they think will work best for them before pulling items from the templates.

The expectation is that SIGs will pull and adapt the options in the templates to best meet the needs of the both the SIG
and project.

- [Recommendations and requirements]

## Templates

- [Short Template]

## Goals

The following documents outline recommendations and requirements for SIG charters and provide
template documents for SIGs to adapt.  The goals are to define the baseline needs for SIGs to
self govern and exercise ownership over an area of the Kubernetes project.

The documents are focused on:

- Defining SIG scope
- Outlining organizational responsibilities
- Outlining organizational roles
- Outlining processes and tools

Specific attention has been given to:

- The role of technical leadership
- The role of operational leadership
- Process for agreeing upon technical decisions
- Process for ensuring technical assets remain healthy

## FAQ

See [frequently asked questions]

[Recommendations and requirements]: sig-governance-requirements.md
[Short Template]: sig-governance-template-short.md
[frequently asked questions]: FAQ.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[sig-architecture example]: ../../sig-architecture/charter.md
