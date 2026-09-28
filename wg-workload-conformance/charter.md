# WG Workload Conformance Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [wg-governance].

## Summary

The Certified Kubernetes Conformance program verifies the cluster
infrastructure layer. No shared community standard verifies the runtime
behavior of the workloads that run on top of it. Workloads break on cluster
upgrades because they depend on deprecated APIs. They ship over-privileged
security defaults, set incorrect resource requests, or fail during routine node
drains. Without a shared baseline, every operator audits every workload again
in every environment.

The working group defines an open, objectively verifiable standard for the
expected runtime behavior of a workload on Kubernetes. It also drives the
creation of a tooling, owned by SIG Architecture, that verifies this behavior.
The specification covers the live runtime behavior of a workload, defined per
Kubernetes minor version. This includes security posture, operational resilience
under disruption, resource footprint, networking, storage, API usage, and
observability. The specification covers how a workload runs, not what it does:
the workload's own application logic is out of scope. The conformance
specification will only have requirements that the tooling can verify
objectively and programmatically. The standard will not have requirements that
are subjective in nature. The working group can work with other SIGs to make a
requirement objective, and only then introduce it into the specification.

The specification, the verification methodology, and the conformance tooling
live in a subproject under SIG Architecture. When the first artifacts exist, SIG
Architecture creates the subproject. This is similar to the existing
conformance-definition and ai-conformance subprojects. SIG Apps
participates in the working group to standardize workload behavior and define
the specs. The working group does not own any code. It is the cross-SIG forum
that provides input on the direction of the subproject while the specification
is still being defined. When the specification and tooling are stable, SIG
Architecture completes a thorough review and gives its final approval. The
subproject then owns the specification and tooling for the long term. If SIG
Architecture does not give that approval, Steering decides whether the working
group continues or disbands. SIG Architecture decides whether to keep or archive
the subproject and its artifacts.

The Kubernetes project will not endorse or support a certification program until
the standard gets created with approval from SIG Architecture. Running that
program is out of scope for this working group and is CNCF's concern. The
subproject must be staffed to continue the work that this working group starts.
The subproject works with CNCF and helps end users who run the conformance
tooling against their workloads. Ongoing maintenance includes updating the
requirements and the tests for newer Kubernetes releases, and dropping
requirements and tests for deprecated features.

The charter takes inspiration from the
[WG Workload Conformance proposal document].

### Meetings

The working group will meet every two weeks to discuss the
specifications. The working group expects to reach stability in 2-3
Kubernetes release cycles. If the specifications become stable earlier, the
working group can wind down earlier.

The regular meetings will follow all the same standard practices followed in the
Kubernetes project.

## In Scope

- Define a workload.
- Define the generic characteristic behavior of a workload on Kubernetes.
- Define the specs to be tested for each Kubernetes minor version.
- Define the verification methodology for each behavior in the specification.
- Drives the creation of upstream conformance tooling that the SIG Architecture
  subproject maintains, similar to the Kubernetes conformance tests.

## Out of Scope

- Run any certification program. It is in scope of CNCF instead.
- Define/test the internal business logic of any workload.
- Prefer or mandate any workload packaging or delivery format, for example
  Helm, Operators, or installers.
- Prefer or mandate any policy engine, scanner, or vendor tool for verification.
- Test the cluster infrastructure layer again. Certified Kubernetes covers it.
- Measure the capability of a cluster to run a class of workloads, for example
  AI/ML readiness. Kubernetes AI Conformance covers it.

## Responsibilities of chairs/organizers

- Run the meetings and guide the direction of the group
- Ensure healthy agenda for each meeting.
- Report progress to the sponsoring SIG chairs.
- Present the annual report to the Steering Committee.

## Disband Criteria

- Initial set of specs are finalized, workload tests reach a stable state of
  operation, and tests are packaged into upstream conformance tooling.
- SIG Architecture completes a thorough review and gives its final approval.
  The SIG Architecture subproject then owns the specs and tooling for the
  long term.
- If SIG Architecture provides their approval, the path opens for an official
  CNCF Workload Conformance Program on top of the standard.
- If SIG Architecture does not give that approval, Steering decides whether the
  working group continues or disbands. SIG Architecture decides whether to keep
  or archive the subproject and its artifacts.

### Stakeholder SIGs

- SIG Architecture: owns the subproject that hosts the specification and
  tooling.
- SIG Apps: participates in the working group and brings the workload and
  application perspective.

When the specification touches their areas, the working group also consults
these groups: SIG Node, SIG Auth, SIG Security, SIG Network, SIG Storage,
SIG Autoscaling, SIG Instrumentation, SIG Testing, WG Node Lifecycle, and
WG Batch.

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[WG Workload Conformance proposal document]: https://docs.google.com/document/d/1BGc4xVcrpQDEDVdGbj5DAvce_VRCSz9zRvQ-5FIdzrM/comment?tab=t.0
