# SIG VMware Charter

## Mission

The VMware SIG maintains and evolves the ability to run Kubernetes on VMware infrastructure.

## Scope

SIG VMware serves to bring together members of the VMware and Kubernetes community to maintain, support and provide guidance for runing Kubernetes on VMware platforms.

One focus area is implementing and maintaining the Kubernetes cloud provider integration code for vSphere.

The SIG also intends to work with sub-projects outside the scope of the cloud provider (examples: network, node, scalability, scheduling, storage, testing). The goal is to represent the perspective of those deploying Kubernetes with VMware software across the breadth of the Kubernetes project.

In scope:

* Deploying Kubernetes of any type (standard or a packaged distribution from any vendor) on vSphere infrastructure
* Utilizing networking and storage features of VMware infrastructure through Kubernetes
* Determining and documenting best practices for configuring Kubernetes when running on vSphere infrastructure.
* Determining and documenting best practices for configuring VMware infrastructure when supporting Kubernetes.
* Determining and documenting best practices for using common configuration management and orchestration tools for deploying and managing Kubernetes on VMware infrastructure
* Discussing bugs and feature requests recorded as Kubernetes or VMware cloud provider issues on GitHub. These issues should be tagged with sig/vmware

Out of scope:

* The VMware SIG is not for discussing bugs or feature requests outside the scope of Kubernetes. For example, the VMware SIG should not be used to discuss or resolve support requests related to VMware products. It should also not be used to discuss topics that other, more specialized SIGs own (to avoid overlap).

SIG is responsible for:

* directing and maintaining the vSphere cloud provider.
* maintaining documentation related to Kubernetes deployment on VMware infrastructure
* implementing and maintaining tests running on VMware infrastructure for coverage of both the vSphere cloud provider as well as general Kubernetes operation.

Responsibilities withing the SIG:

* SIG chair role
  * Plan, publicize and operate meetings.
  * Organize and track activity within the charter of the SIG.
  * Facilitate disussion and technical decision making process.
  * Maintain SIG community artifacts (github, mailing list, meeting recordings, Slack channel).
  * Deliver SIG Update presentations on a regular basis in the Kubernetes Community Meeting and at Contributor Summit and KubeCon events.
  * Chairs MAY decide to step down at anytime and propose a replacement. Use lazy consensus amongst chairs with fallback on majority vote to accept proposal. This SHOULD be supported by a majority of SIG Members.
  * Chairs MAY select additional chairs through a super-majority vote amongst chairs. This SHOULD be supported by a majority of SIG Members.
  * Chairs MUST remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and MAY be removed if not proactively working with other chairs to fulfill responsibilities. Coordinated leaves of absence serve as exception to this requirement.
  * number of chairs: 2 - 3
  * current chars are identified in sigs.yaml
* SIG technical lead role
  * SIG chairs hold operational and technical roles at this time. 
  * Establish new subprojects, and retire existing subprojects
  * Resolve cross-subproject technical issues and decisions, and escalations from subprojects
  * Decision-making MUST be by consensus. It’s particularly important for the technical leads to provide cohesive technical guidance for the project as a whole.

TODO

* Document initial reviewers and approvers for repositories.