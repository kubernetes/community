# WG Node Identity Charter

This charter adheres to the conventions described in the [Kubernetes Charter README](https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md) and uses the Roles and Organization Management outlined in [wg-governance](https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md).

## Background

Node attestation isn't new in the Kubernetes ecosystem, but every distribution that takes it seriously has built its own approach \- [GKE attests nodes with a virtual TPM](https://cloud.google.com/kubernetes-engine/docs/how-to/shielded-gke-nodes), [kOps verifies cloud-native instance identity](https://kops.sigs.k8s.io/architecture/kops-controller/), and [EKS authenticates nodes through their IAM roles](https://github.com/kubernetes-sigs/aws-iam-authenticator). Each proves the same thing: that a node is what it claims to be, but none of these interoperate, none share a common security review, and an operator moving between distributions or wanting to combine different clouds meets a different node identity model each time.

For self-managed clusters the common case is kubeadm bootstrap tokens, which are bearer credentials that can't be scoped to a specific node. [Cluster API](https://github.com/kubernetes-sigs/cluster-api), and [Karpenter on self-managed kubeadm](https://github.com/aws/karpenter-provider-aws/issues/4492) both inherit this gap. The 2022 [Cluster API Security Self-Assessment](https://github.com/kubernetes/sig-security/blob/main/sig-security-assessments/cluster-api/self-assessment.md) flagged it as critical — an attacker who reads a token can register as any node, including a control plane node — and its fix, [replace the kubeadm join with host attestation](https://github.com/kubernetes-sigs/cluster-api/issues/3762), was marked "Planned" and never landed.

The appetite to fix this was clear at the [KubeCon EU 2026 maintainer panel "From Static Tokens to Attestation"](https://www.youtube.com/watch?v=MIO3tDk0GnI), which drew roughly 150 attendees, ran past time, and spilled into the hallway track. The recurring theme: this has stayed unsolved not for lack of building blocks but because no SIG owns the end-to-end problem.

The proposers bar is that a solution be low-effort to *consume*; solutions like SPIFFE exist, is a possible backend, but require additional infrastructure that strays from the ease of use of cluster bring up with common tooling.

We acknowledge that clouds differ too much for one uniform implementation \- some expose a vTPM and signed metadata, others hand you a bare VM and an IP. So this WG doesn't assume one mechanism fits all. It aims to give the problem a cross-SIG home and a legible model: a consistent interface in core Kubernetes with thin per-provider implementations behind it \- so what differs is the mechanism, not the user's mental model.

Above all, it's a venue to gather the people already interested, scattered today across SIGs, distributions, and practitioners in one place to solve it.

## Scope

The scope of this WG is to define “node attestation” and propose deliverables that need to be adopted to deploy, secure, cloud–provider/hardware-backed node identity and attestation with a common model across environments. The goal is for attestation to be enabled by default and require minimal configuration on platforms that support it, making secure node identity the easy path rather than an advanced option. We succeed if following a Getting Started guide in the Kubernetes docs gets you an attested cluster.

### In Scope

- Defining terms such as a Kubernetes-native node attestation that allows kubelets to prove their identity to the control plane using platform-provided attestation (TPM, vTPM, cloud instance identity documents, or similar mechanisms)
- Define use cases for Kubernetes users that address the heterogeneity of environments.
- Determining which common features (e.g. TPM attestation) can be covered by APIs.
- Propose new sub-projects if existing sub-projects are not sufficient.
- Producing a threat model for the node identity lifecycle covering bootstrap, certificate rotation, and node impersonation attacks

### Out of Scope

- Developing any node attestation mechanism outside of the ownership of a sponsoring SIG.
- Long term ownership of cloud provider-specific implementations of attestation verifiers
- Workload identity (existing solutions exist)
- Removal of bootstrap token support — bootstrap tokens will remain available as a fallback, but attestation should be the default for distributions that support it

## Deliverables

* An accepted KEP and sponsoring SIG
* A common threat model applicable to multiple cluster lifecycle management tools, clouds, virtualized environments and bare metal.
* Provide  a space for collaboration across cloud providers and practitioners to determine common, viable features that should be supported by Kubernetes SIGs. Given consensus in one or more ideas, the WG will facilitate and coordinate the delivery of proposals in the appropriate areas with owning SIGs.

## Stakeholders

* SIG Node
* SIG Cluster Lifecycle
* SIG Autoscaling
* SIG Cloud Provider

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in [wg-governance](https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md) and opts-in to updates and calculation of the [Conditions for Disbanding](https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md#disbanding).

### Chairs

- Rodrigo Campos Catelin (@rata)
- Ciprian Hacman (@hakman)
- Naadir Jeewa (@randomvariable)
- Michael McCune (@elmiko)
- Josephine Pfeiffer (@pfeifferj)

### Steering Committee Liaison

- TBD (assigned after approval)

## Exit Criteria

This WG will disband when its deliverables are complete

1. Completed definitions and key use cases for cluster owners and lifecycle management implementations, documented and committed.
2. Key common features that Kubernetes or a sub-project needs to best support the defined use cases.
3. For each feature in the list, proposals/KEPs which support them are made to the relevant sub-project OR propose new sub-projects if deemed necessary. The KEPs must move to Accepted as a condition of exit.

### Alternative Exit Paths

- **Early exit**: If the WG determines during Phase 1 that node attestation cannot be securely implemented within the existing Kubernetes certificate infrastructure without unacceptable complexity, it will publish findings and recommendations and disband
- **Partial exit**: If the security review identifies fundamental issues with moving attestation into core or a subproject, the WG will document the findings and recommend that attestation remain an out-of-tree extension point

### Progress Reporting

- Annual reports to the steering committee (due by 1 March)
- Quarterly updates to sponsoring SIG mailing lists (sig-node@, sig-cluster-lifecycle@, @sig-cloud-provider)
- Monthly WG meetings open to the community
