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
- Prow, testgrid, perf dashboard integrations to expand and maintain testing (e2e, jobs) and scale-testing (load, density)
- Support users on their issues and feature requests
- Documentation for all things Kubernetes on AWS

#### Cross-cutting and Externally Facing Processes

- Consult with other SIGs and the community on how to apply mechanisms owned by SIG
  AWS. Examples include:
    - Review escalation implications of feature and API designs as it relates to core Kubernetes components (etcd, kubelet, apiserver, controller manager, scheduler)
    - CSI, CRI implementation and design
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
- Reporting specific vulnerabilities in Kubernetes. Please report using these instructions: https://kubernetes.io/security/

## Roles and Organization Management

This SIG adheres to the Roles outlined in [sig-governance] and opts-in to updates and modifications to the same.

The following roles are required for the SIG to function properly. In the event that any role is unfilled, the SIG will make a best effort to fill it. Any decisions reliant on a missing role will be postponed until the role is filled.

### Chairs

- 2 or 3 chairs are required
- Run operations and processes governing the SIG
- An initial set of chairs was established at the time the SIG was founded.
- Chairs MAY decide to step down at anytime and propose a replacement, who must be approved by all of the other chairs. This SHOULD be supported by a majority of SIG Members.
- Chair election must be transparent and should follow the following steps:
  - initiate the proposal for a new Chair by sending an email to kubernetes-sig-aws@googlegroups.com
  - propose the new Chair as part of the agenda in the biweekly SIG meeting
  - gather votes (10+ preferred) from both forums to initiate a PR and change the entry for the lead in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml) 
- Chairs MAY select additional chairs using the same election process amongst SIG Members.
- Chairs MUST remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and MAY be removed by consensus of the other Chairs and members if not proactively working with other Chairs to fulfill responsibilities.
- Chairs WILL be asked to step down if there is inappropriate behavior or code of conduct issues
- SIG AWS cannot have more than 1 chair from any one company.

### Subproject/Provider Owners

- There should be at least 1 representative per subproject/provider (though 3 is recommended to avoid deadlock) as specified in the OWNERS file of each cloud provider repository.
- MUST be an escalation point for technical discussions and decisions in the subproject/provider
- MUST set milestone priorities or delegate this responsibility
- MUST remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and MAY be removed by consensus of other subproject owners and Chairs if not proactively working with other Subproject Owners to fulfill responsibilities.
- MAY decide to step down at anytime and propose a replacement. This can be done by updating the OWNERS file for any subprojects.
- MAY select additional subproject owners by updating the OWNERs file.
- WILL be asked to step down if there is inappropriate behavior or code of conduct issues

### SIG Members

Approvers and reviewers in the OWNERS file of all subprojects under SIG AWS.

## Organization Management

- Six months after this charter is first ratified, it MUST be reviewed and re-approved by the SIG in order to evaluate the assumptions made in its initial drafting
- SIG meets bi-weekly on zoom with agenda in meeting notes.
- SHOULD be facilitated by chairs unless delegated to specific Members
- The SIG MUST make a best effort to provide leadership opportunities to individuals who represent different races, national origins, ethnicities, genders, abilities, sexual preferences, ages, backgrounds, levels of educational achievement, and socioeconomic statuses

### Subproject Creation

SIG AWS delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].

### Subproject Retirement

SIG AWS subprojects may be retired if they do not satisfy requirements or the renewed scope for more than 6 months. Final decisions for retirement should be supported by a majority of SIG members using [lazy consensus](http://communitymgt.wikia.com/wiki/Lazy_consensus). Once retired any code related to that subproject will be archived into the kubernetes-retired organization.

Subprojects representing may be retired at any point given a lack of development or a lack of demand. Final decisions for retirement should be supported by a majority of SIG members, ideally from active maintainers and members. Once retired, any code related to that subproject will be archived into the kubernetes-retired organization.

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-aws/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
