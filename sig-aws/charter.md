# SIG AWS Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG AWS is responsible for the creation and maintenance of subprojects (features/innovations) necessary to integrate AWS services for the operation and management of Kubernetes on AWS. SIG AWS also acts as a forum for Kubernetes on AWS users/developers to raise their feature requests and support issues with. SIG leads in collaboration with SIG members will make a best effort to triage known problems within one or two release cycle of issues being reported. SIG AWS in collaboration with SIG-Testing, SIG-Scalability and SIG-Docs is responsible for integration and maintenance of tests (e2e, periodic jobs, postsubmit jobs etc.); scale-tests (load, density tests) and documentation for the scope within the purview of this charter.

### In scope

Link to SIG [subprojects](https://github.com/kubernetes/community/tree/master/sig-aws#subprojects)

#### Code, Binaries and Services

Kubernetes integrations specific to AWS including:
- Integrations, interfaces, libraries and extension points for all AWS services such as IAM, storage, networking, loadbalancers, registry, security, monitoring/logging at the instance or container level
- Tools for Kubernetes APIs to work with AWS services including Amazon EKS
- Prow, testgrid, perf dashboard integrations to expand and maintain testing (e2e, jobs) and scale-testing (load, density) on AWS and Amazon EKS
- Support users on their issues and feature requests
- Documentation for all things Kubernetes on AWS

#### Cross-cutting and Externally Facing Processes

- Consult with other SIGs and the community on how to apply mechanisms owned by SIG
  AWS. Examples include:
    - Review escalation implications of feature and API designs as it relates to core Kubernetes components (etcd, kubelet, apiserver, controller manager, scheduler)
    - CSI, CNI, CRI implementation and design
    - Cloud provider implementation and design
    - Best practices for hardening add-ons or other external integrations such as KMS, LB, others.
    - Implementing and hardening tests, scale tests and documentation

### Out of scope

SIG AWS is not for discussing bugs or feature requests outside the scope of Kubernetes. For example, SIG AWS should not be used to discuss or resolve support requests related to AWS Services. It should also not be used to discuss topics that other, more specialized SIGs own (to avoid overlap). Examples of such scenarios include:
- Specification of CSI, CRI interfaces, cloudprovider binary (prefer: sig-storage, sig-node and sig-cloudprovider)
- Container runtime (prefer: sig-node and sig-networking)
- Resource quota (prefer: sig-scheduling)
- Resource availability (prefer: sig-apimachinery, sig-network, sig-node)
- Detailed design and scope of tests or tooling to run tests (prefer: sig-testing)
- Detailed design and scope of scale tests or tooling to run scale tests (prefer: sig-scalability)
- Troubleshooting and maintenance of test jobs related to kops (prefer: sig-cluster-lifecyle)
- Reporting specific vulnerabilities in Kubernetes. Please report using these instructions: https://kubernetes.io/security/

## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

SIG AWS delegates subproject approval to Chairs. Chairs also act as Technical Leads in SIG AWS. See [Subproject creation - Option 1].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-aws/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
