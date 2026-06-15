# WG Node Identity Charter

This charter adheres to the conventions described in the [Kubernetes Charter README][]
and uses the Roles and Organization Management outlined in [wg-governance][].

[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md

## Scope

The scope of this WG is to define “node attestation” and propose
deliverables that need to be adopted to deploy and secure
cloud–provider/hardware-backed node identity and attestation with a common
model across environments. The goal is for attestation to be enabled by
default and require minimal configuration on platforms that support it, making
secure node identity the easy path rather than an advanced option. We succeed
if following a Getting Started guide in the Kubernetes docs gets you an
attested cluster.

### In Scope

- Defining terms such as a Kubernetes-native node attestation that allows
  kubelets to prove their identity to the control plane using platform-provided
  attestation (TPM, vTPM, cloud instance identity documents, or similar mechanisms)
- Define use cases for Kubernetes users that address the heterogeneity of environments.
- Determining which common features (e.g. TPM attestation) can be covered by APIs.
- Propose new sub-projects, to be hosted by a stakeholder SIG, if existing
  sub-projects are not sufficient.
- Producing a threat model for the node identity lifecycle covering bootstrap,
  certificate rotation, and node impersonation attacks

### Out of Scope

- Developing any node attestation mechanism outside of the ownership of any
  sponsoring SIGs.
- Long term ownership of cloud provider-specific implementations of attestation verifiers
- Workload identity (existing solutions exist)
- Removal of bootstrap token support — bootstrap tokens will continue to
  remain available

## Stakeholders

* SIG Autoscaling
* SIG Cloud Provider
* SIG Cluster Lifecycle

## Deliverables

* An accepted KEP
* A common threat model applicable to multiple cluster lifecycle management
  tools, clouds, virtualized environments and bare metal.
* Provide a space for collaboration across cloud providers and practitioners to
  determine common, viable features that should be supported by Kubernetes SIGs.
  Given consensus in one or more ideas, the WG will facilitate and coordinate
  the delivery of proposals in the appropriate areas with owning SIGs.

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in
[wg-governance][] and opts-in to updates and calculation of the
[Conditions for Disbanding][].

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Conditions for Disbanding]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md#disbandment-process-description

### Chairs

- Taahir Ahmed (@ahmedtd)
- Rodrigo Campos Catelin (@rata)
- Ciprian Hacman (@hakman)
- Naadir Jeewa (@randomvariable)
- Michael McCune (@elmiko)
- Josephine Pfeiffer (@pfeifferj)

## Timelines and Disbanding

This WG will disband when its deliverables are complete

1. Completed definitions and key use cases for cluster owners and lifecycle management implementations, documented and committed.
2. Key common features that Kubernetes or a sub-project needs to best support the defined use cases.
3. For each feature in the list, proposals/KEPs which support them are made to the relevant sub-project OR propose new sub-projects if deemed necessary. The KEPs must move to Accepted as a condition of exit.

### Alternative Exit Paths

- **Early exit**: If the WG determines during Phase 1 that node attestation cannot be securely implemented within the existing Kubernetes certificate infrastructure without unacceptable complexity, it will publish findings and recommendations and disband
- **Partial exit**: If the security review identifies fundamental issues with moving attestation into core or a subproject, the WG will document the findings and recommend that attestation remain an out-of-tree extension point
