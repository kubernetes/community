# WG Kubernetes Workload Conformance Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [wg-governance].

## Summary

The working group defines an open, objectively verifiable standard for how a
workload should behave at runtime on Kubernetes, and delivers the tooling that
verifies it. The specification covers a workload's live runtime behavior - its
security posture, operational resilience under disruption, resource footprint,
networking, storage, API stability (both the core Kubernetes API and
node-local APIs such as the kubelet API), and observability - defined per
Kubernetes minor version. It covers how a workload runs, not what it does: the
workload's own application logic is out of scope.

The specification, the verification methodology, and the conformance tooling
live in a subproject under SIG Apps from the outset. The working group
doesn't own any code itself; it's the cross-SIG forum that steers the
subproject's direction while it's still being defined. Once the
specification and tooling reach stability, SIG Apps transitions the
subproject to SIG Architecture, which owns it for the long term. Running any
certification program on top of the standard is out of scope and is CNCF's
concern.

### Meetings

The working group will meet every two weeks for 60 minutes to discuss the
specifications. Ideally, it would take 2-3 Kubernetes release cycles to reach
stability. In case stability of specifications is reached earlier than that, the
working group might decide to wrap even before.

The regular meetings will follow all the same standard practices followed in the Kubernetes project.

## In Scope

- Define a workload.
- Define the ideal characteristic behavior of an workload when running on Kubernetes.
- Define the specs to be tested for each workload against each Kubernetes
 version.
- Define the verification methodology for each behavior in the specification.
- Deliver a conformance image that can be maintained by a subproject similar to the Kubernetes conformance image.

## Out of Scope

- Run any certification program. It is in scope of CNCF instead.
- Define/test the internal business logic of any workload.
- Prefer or mandate any workload packaging or delivery format (Helm,
 Operators, installers, etc.).

## Organizers

- Nabarun Pal
- Maciej Szulik
- Liz Rice
- Ricardo Aravena

All organizers and leadership-role holders must be [community members].

## Responsibilities of chairs/organizers

- Run the meetings and guide the direction of the group
- Ensure healthy agenda for each meeting.
- Report progress to the sponsoring SIG chairs.

## Disband Criteria

- Initial set of specs are finalized and delivered to a subproject.
- Workload tests reach a stable state of operation.
- Workload tests are packaged to a conformance tooling via Hydrophone.
- SIG Apps transitions the conformance tooling and artifacts to a SIG Architecture subproject, and the specs keep on evolving with collaboration from SIG Apps and other participating SIGs.

### Stakeholder SIGs

- SIG Apps
- SIG Architecture

[community members]: https://github.com/kubernetes/community/blob/master/community-membership.md
[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
