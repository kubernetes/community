# VMware User Group Charter

This User Group adheres to the governance described in [ug-governance] under the community standards of the Kubernetes project.

## Scope
The VMware User Group facilitates communication among Kubernetes users and contributors on topics pertaining to running all [conformant] forms of Kubernetes on VMware infrastructure. Scope includes extensions and integration of Kubernetes with Kubernetes subprojects.

In addition to the usage of Kubernetes running on the cloud provider for the vSphere hypervisor, the User Group’s mission includes:

- Running Kubernetes on desktop hypervisors (Fusion for OS X and Workstation for Linux and Windows), as supported by Minikube.
- vCloud Director, which enables service providers to provision and manage multi tenant clouds, which can host Kubernetes.
- Utilizing networking and storage features of VMware infrastructure through Kubernetes.
- In the context of Kubernetes user use cases:
  - Facilitating determination and documentation of best practices for configuring Kubernetes when running on vSphere infrastructure.
  - Facilitating determination and documentation of best practices for configuring VMware infrastructure when supporting Kubernetes.
  - Facilitating determination and documentation of best practices for using common configuration management, installation and orchestration tools for deploying and managing Kubernetes on VMware infrastructure.
- Discussing bugs and feature requests recorded as Kubernetes, VMware cloud provider, or VMware storage plugin issues on GitHub. The UG itself is not responsible for actual management and resolution of issues, but merely facilitates members in creating well crafted issues in appropriate places - along with responses to comments and questions from prospective fixers within the issue.   

## Out of Scope
The VMware User Group does NOT:
- Produce or own any Kubernetes code or project deliverables
- Make Kubernetes project decisions
- Claim ownership of any specific topic
- Operate as a forum for discussing bugs outside the scope of Kubernetes and related open source cloud providers and plugins. For example, the VMware User Group should not be used to discuss or resolve support requests related to VMware products.
 
## UG Example Activities
- Facilitating collaboration between user group members - which may result in the production of - blogs, docs, guides, demos, presentations, prototypes, external contributions, etc. (UG is itself not responsible to create any such content, it merely facilitates members creating them).
- Anything produced within the context of the user group can be ultimately owned by a SIG as a subproject, or is owned by individuals in the user group.
- Publishing content such as FAQs, READMEs, or other forms of documentation to a GitHub directory owned by the user group. Access to this location is controlled by OWNERS files.
- Providing updates from the user group’s activities to the community at large.
  - Cadence - every release.
  - Suggested Improvements, Valued Contributions, Pain-points to be addressed, channeled through the user group as feedback to the ecosystem.

## Organization
As provided in the User Group [Prerequisites], the user group will start with two chairs, and two inaugural Kubernetes user representative members, all of whom are community [members]. The identity of people holding these positions will be documented in [sigs.yaml], with user representative members designated as tech leads. Procedures for continued staffing of the chair and tech lead member roles is governed by the [Roles] procedure described in [sig-governance.md] - noting that the WG will utilize these SIG Roles procedures by reference, but procedures related to subprojects in sig-governance.md do not apply since subprojects are not present in a UG.

[ug-governance]: ../committee-steering/governance/ug-governance.md
[conformant]: https://www.cncf.io/certification/software-conformance/
[Prerequisites]: ../sig-wg-lifecycle.md#prerequisites-for-a-ug
[members]: ../community-membership.md
[sigs.yaml]: ../sigs.yaml
[Roles]: ../committee-steering/governance/sig-governance.md#roles
[sig-governance.md]: ../committee-steering/governance/sig-governance.md