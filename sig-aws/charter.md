# SIG AWS Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG AWS is responsible for the creation and maintenance of subprojects (features/innovations) necessary to operate Kubernetes on AWS (see below). SIG AWS acts as a forum for Kubernetes on AWS users/developers to raise their feature requests and support issues with. SIG AWS in collaboration with SIG-Testing, SIG-Scalability and SIG-Docs is responsible for integration and maintenance of tests, scale-tests and documentation for all things Kubernetes on AWS.

### In scope

Link to SIG section in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml) and SIG [subprojects](https://github.com/kubernetes/community/tree/master/sig-aws#subprojects)

#### Code, Binaries and Services

Kubernetes integrations specific to AWS including:
- Integrations, interfaces, libraries and extension points for all AWS services such as IAM, storage, networking, loadbalancers, registry, security, monitoring/logging at the instance or container level
- Tools for Kubernetes APIs to work with AWS services including Amazon EKS
- Prow, testgrid, perf dashboard integrations to expand and maintain testing and scale-testing
- Support users on their issues and feature requests
- Socumentation for all things Kubernetes on AWS

#### Cross-cutting and Externally Facing Processes

- Consult with other SIGs and the community on how to apply mechanisms owned by SIG
  AWS. Examples include:
    - Review escalation implications of feature and API designs as it relates to core Kubernetes components (etcd, kubelet, apiserver, controller manager, scheduler)
    - CSI, CRI implementation and design
    - Cloud provider implementation and design
    - Best practices for hardening add-ons or other external integrations such as KMS, LB, others.
    - Implementing and hardening tests, scale tests and documentation

### Out of scope

SIG AWS is not for discussing bugs or feature requests outside the scope of Kubernetes. For example, the SIG AWS should not be used to discuss or resolve support requests related to AWS Services. It should also not be used to discuss topics that other, more specialized SIGs own (to avoid overlap). Examples of such scenarios include:
- Specification of CSI, CRI interfaces, cloudprovider binary (prefer: sig-storage, sig-node and sig-cloudprovider)
- Container runtime (prefer: sig-node and sig-networking)
- Resource quota (prefer: sig-scheduling)
- Resource availability (prefer: sig-apimachinery, sig-network, sig-node)
- Testing and scale testing scope (prefer: sig-testing, sig-scalability)
- Reporting specific vulnerabilities in Kubernetes. Please report using these instructions: https://kubernetes.io/security/

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

SIG AWS delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
