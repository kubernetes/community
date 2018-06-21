# SIG VMware Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The VMware SIG maintains and evolves the ability to run Kubernetes on VMware infrastructure.

In addition to the cloud provider for the vSphere hypervisor, the SIG’s mission includes:

- NSX and NSX-T, infrastructure components that provide software based networking service, with automated micro-segmentation at application and control plane levels - along with security, monitoring, logging, and features for multi-cluster, multi-site, and cross cloud integration. NSX-T is not vSphere specific, and is supported on other cloud providers (e.g. OpenStack). The scope of NSX and NSX-T interaction with Kubernetes goes beyond the CNI plug-in abstraction.
- Desktop hypervisors (Fusion for OS X and Workstation for Linux and Windows), as supported by Minikube.
- vCloud Director, which enables service providers to provision and manage multi tenant clouds, which can host Kubernetes.


### In scope
- Deploying Kubernetes of any type (standard or a packaged distribution from any vendor) on vSphere infrastructure.
- Utilizing networking and storage features of VMware infrastructure through Kubernetes.
- Determining and documenting best practices for configuring Kubernetes when running on vSphere infrastructure.
- Determining and documenting best practices for configuring VMware infrastructure when supporting Kubernetes.
- Determining and documenting best practices for using common configuration management and orchestration tools for deploying and managing Kubernetes on VMware infrastructure.
- Discussing bugs and feature requests recorded as Kubernetes or VMware cloud provider issues on GitHub. These issues should be tagged with ``sig/vmware``.

The authoritative source for SIG information is the [sigs.yaml] file.

#### Code, Binaries and Services

- vSphere Cloud Provider
- vRealize Automation Cluster API Provider

#### Cross-cutting and Externally Facing Processes

SIG VMware serves to bring together members of the VMware and Kubernetes community to maintain, support and provide guidance for running Kubernetes on VMware platforms.

The VMware SIG will provide a forum for hosting related architectural planning and discussion. Associated activities related to development, testing, and documentation will be tracked and reported by the SIG.

- The vSphere platform’s availability and resource management can create pools that extend across Kubernetes worker node, management node, and namespace boundaries. Unique features of the vSphere platform related to high availability, load balancing, and resource management (storage, networking, compute) can supplement existing capabilities of Kubernetes, container runtime, and OS functionality.

  - Users can do some of this via configuration today, but opportunities exist to achieve more efficient, predictable, and dependable service through additional new CRDs, plug-ins, and enhancement proposals (KEPs) related to other Kubernetes components.
  - Interaction is expected to cut across multiple aspects of Kubernetes in ways not firmly bounded by existing SIG dividing lines (e.g. node, scheduling, scalability).

### Out of scope

The VMware SIG is not for discussing bugs or feature requests outside the scope of Kubernetes. For example, the VMware SIG should not be used to discuss or resolve support requests related to VMware products. It should also not be used to discuss topics that other, more specialized SIGs own (to avoid overlap).

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

Subprojects associated with this SIG may be created following the procedure described in [sig-governance] as:

> by Federation of Subprojects

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1706
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md