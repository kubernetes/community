# WG Workload Conformance - Creation Proposal


This document is the creation proposal for a new Kubernetes Working Group,
**WG Workload Conformance**. It answers the questions required by
[wg-governance.md] and follows the checklist in [sig-wg-lifecycle.md]. The
formal charter lives alongside this document in [charter.md].


## Summary


The Certified Kubernetes Conformance program standardizes and verifies the
cluster infrastructure layer, but there is no shared community standard that
verifies the operational behavior of the *workloads* running on top of it.
Workloads routinely break on cluster upgrades because they depend on deprecated
APIs, ship over-privileged security defaults, mis-size their resource requests,
or fail to survive routine node drains. These are not exotic failures; they
happen by default in the absence of a shared baseline, and every operator ends
up re-auditing every workload across every environment.


WG Workload Conformance is a time-boxed effort to define an open, objectively
verifiable standard for how a well-behaved workload runs on Kubernetes, to
deliver the tooling that verifies it. The specification covers a workload's live
runtime behavior — its security posture, operational resilience under
disruption, accurate resource footprint, storage portability, usage of stable
APIs (both the core Kubernetes API and node-local APIs such as the kubelet
API), and provide metrics endpoint — while explicitly excluding the
correctness of the workload's own application logic.


The Working Group defines the specification and the conformance tooling. It does
**not** run any certification program; operating a certification program on top
of this standard is CNCF's concern, not the Working Group's. The specification
and tooling live in a SIG Apps subproject from the outset — the
Working Group doesn't own any code itself, it's the time-boxed, cross-SIG
forum that shapes the spec while the subproject is still being defined. Once
the specification and tooling reach stability, SIG Apps transitions the
subproject to SIG Architecture, which owns it for the long term.

Detailed proposal: https://docs.google.com/document/d/1BGc4xVcrpQDEDVdGbj5DAvce_VRCSz9zRvQ-5FIdzrM/comment?tab=t.0

## Answers to the Working Group formation questions

These are the questions from [wg-governance.md] that a creation proposal must
answer.

### 1. What is the exact problem this group is trying to solve?

There is no shared, community-owned standard for how a workload should behave at
runtime on Kubernetes, and no common tooling to verify that behavior. As a
result:

- Workloads break during cluster upgrades because they call deprecated or
 removed APIs — not just the Kubernetes API surface, but node-local ones too,
 as with dockershim, the cgroup v1 removal, or the kubelet's read-only port.
- Workloads use insecure defaults — running as root, requesting excessive RBAC,
 requiring host namespaces — expanding the attack surface for no reason.
- Workloads are mis-sized: bloated requests force over-provisioning, missing
 limits cause node-level OOM cascades.
- Maintainers repeat the same validation and debugging across a wide matrix of
 providers and delivery formats (Helm, Operators, node agents, installers).
- Operators repeat per-workload audits that a shared baseline would make
 unnecessary.

The group's job is to turn "run like a good Kubernetes citizen" from tribal
knowledge into an objective, versioned, verifiable specification.


### 2. What is the artifact that this group will deliver, and to whom?

The Working Group delivers three artifacts:

1. **A Workload Conformance specification** — the set of verifiable behaviors a
  workload must exhibit, classified MUST / SHOULD, defined per Kubernetes minor
  version.
2. **A behavior verification methodology** — how each behavior in the
  specification is tested against a live workload.
3. **Conformance tooling** — a conformance test suite / image (in the spirit of
  the existing Kubernetes conformance image, e.g. runnable via Hydrophone) that
  executes the methodology against a running workload.

These live in a **SIG Apps** subproject from the outset; the Working Group
doesn't own any code, it's the cross-SIG forum that shapes the spec while the
subproject is still finding its shape. Once the specification and tooling
reach stability, SIG Apps transitions the subproject to **SIG Architecture**,
which owns it for the long term — the same way SIG Architecture owns the
existing Kubernetes conformance suite. Downstream, the CNCF certification
program and the broader ecosystem (workload maintainers, platform operators)
consume the standard.


Submission under the eventual CNCF certification program is intentionally
decoupled from Certified Kubernetes: the workload's owner or maintainer
submits, not the cluster vendor or distributor, since they're the ones who
control the manifests, RBAC, probes, and PDBs being assessed, not the cluster
underneath them. An ISV shipping a Helm chart for a message queue or an
Operator for a stateful database would run the suite against their own
product in its live, deployed state and self-attest, the same way GKE, EKS,
AKS, or OpenShift assert Certified Kubernetes for the cluster layer. These are
two independent claims about two different layers — infrastructure and
application — and the Working Group's specification only defines the latter.


### 3. How does the group know when the problem-solving process is complete, and it is time to dissolve?


The Working Group disbands when all of the following hold:


- An initial, agreed-upon set of workload conformance specs and their
 verification methodology are finalized within the SIG Apps subproject.
- The conformance tooling reaches a stable, runnable state (e.g. via
 Hydrophone) against real workloads.
- SIG Apps transitions the subproject — specs, methodology, and tooling — to
 SIG Architecture, which keeps evolving them on its own, in collaboration
 with SIG Apps and other participating SIGs, without the cross-SIG working
 group coordinating it.


The group expects this to take roughly a few Kubernetes release cycles. If
stability is reached earlier, it may wind down earlier.


### 4. Who are all the stakeholder SIGs involved in this problem?


- **SIG Apps** — hosts the subproject and owns the deliverables from the
 outset, representing the workload/application perspective and the
 archetypes being certified.
- **SIG Architecture** — owns conformance definitions and tooling long term;
 the subproject transitions to SIG Architecture once the specs and tooling
 reach stability.


Additional SIGs expected to be consulted as the spec touches their domains
(not necessarily formal stakeholders at founding): SIG Node (node-local APIs
and deprecations, e.g. the kubelet API, dockershim, cgroup v1), SIG Auth (RBAC, Pod
Security, admission), SIG Network (NetworkPolicy, Ingress, Gateway), SIG
Storage (access modes, portable volumes), SIG Instrumentation (metrics,
structured logs), SIG Testing (conformance test infrastructure), WG Node Lifecycle (pod lifecycle changes around node lifecycle) and WG Batch (batch workloads).


### 5. What are the meeting mechanics (frequency, duration, roles)?


- **Frequency:** every two weeks.
- **Duration:** 60 minutes.
- **Roles:** WG organizers run meetings, curate the agenda, and report to
 sponsoring SIG chairs. Meetings follow standard Kubernetes project practices
 (public agenda doc, recorded, notes archived).


### 6. Does the goal represent the needs of the project as a whole, or a narrow set of contributors/companies?


The whole project. The absence of a workload behavior baseline is a horizontal
problem that hits every operator and every maintainer regardless of vendor,
distribution, or delivery format. The specification is defined against upstream
Kubernetes behavior (Pod Security Standards, deprecated-API signals, standard
API fields) rather than any single vendor's platform, and the deliverables are
owned upstream — initially by SIG Apps, transitioning to SIG Architecture once
stable. Participation is open, and the archetype set
spans the ecosystem rather than any one company's portfolio.


### 7. Who will chair the group and ensure it continues to meet these requirements?

Proposed organizers:
- Nabarun Pal
- Maciej Szulik
- Liz Rice
- Ricardo Aravena

### 8. Is diversity well-represented in the Working Group?

Organizers span multiple companies and multiple areas, and the group
will actively recruit participants across the stakeholder SIGs and across the
workload archetypes to avoid a single-vendor or single-domain skew.

## In Scope


- Define what constitutes a "workload" for the purposes of conformance.
- Define the ideal characteristic runtime behavior of a workload on Kubernetes.
- Define the behaviors to be tested for each workload, per Kubernetes minor
 version, covering node-local APIs (e.g. the kubelet API) alongside the core
 Kubernetes API surface.
- Define the verification methodology for each behavior.
- Deliver conformance tooling / a conformance image, maintainable by a
 subproject — initially under SIG Apps, transitioning to SIG Architecture
 once stable — in the spirit of the Kubernetes conformance image.


## Out of Scope


- Running the certification program itself — that is CNCF's concern.
- Defining or validating the internal business logic of any workload.
- Re-testing the infrastructure layer already covered by Certified Kubernetes.
- Measuring a cluster's capability to run a class of workloads (e.g. AI/ML
 readiness), which is handled by Kubernetes AI Conformance.
- Preferring or mandating any workload packaging or delivery format (Helm
 charts, Operators, installers, etc.) — the specification covers a workload's
 runtime behavior, not how it's packaged to reach the cluster, and can't live
 inside any single packaging project's governance.


## Deliverables


- Workload Conformance specification (behaviors, MUST/SHOULD, per K8s version).
- Behavior verification methodology.
- Conformance test suite / image handed to a SIG Apps subproject, later
 transitioning to SIG Architecture.


## Disband Criteria


- Initial set of specs and methodology finalized within the SIG Apps subproject.
- Conformance tooling (via Hydrophone) reaches a stable state of operation.
- SIG Apps transitions the subproject to SIG Architecture, which owns and
 continues to evolve the specs, in collaboration with SIG Apps and other
 participating SIGs.


## Stakeholder SIGs


- SIG Apps
- SIG Architecture


[wg-governance.md]: /committee-steering/governance/wg-governance.md
[sig-wg-lifecycle.md]: /sig-wg-lifecycle.md
[charter.md]: ./charter.md
[community members]: /community-membership.md
