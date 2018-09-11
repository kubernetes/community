# SIG VMware Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The VMware SIG maintains and evolves the ability to run Kubernetes on VMware infrastructure.

In addition to the cloud provider for the vSphere hypervisor, the SIG’s mission includes:

- NSX and NSX-T, infrastructure components that provide software based networking service, with automated micro-segmentation at application and control plane levels - along with security, monitoring, logging, and features for multi-cluster, multi-site, and cross cloud integration. NSX-T is not vSphere specific, and is supported on other cloud providers (e.g. OpenStack). The scope of NSX and NSX-T interaction with Kubernetes goes beyond the CNI plug-in abstraction.
- Desktop hypervisors (Fusion for OS X and Workstation for Linux and Windows), as supported by Minikube.
- vCloud Director, which enables service providers to provision and manage multi tenant clouds, which can host Kubernetes.
- Hosting architectural planning and discussion related to new CRDs, plug-ins and KEPs that allow the vSphere platform to supplement the existing capabilities of Kubernetes, container runtime, and OS functionality


### In scope
- Deploying Kubernetes of any type (standard or a packaged distribution from any vendor) on vSphere infrastructure.
- Utilizing networking and storage features of VMware infrastructure through Kubernetes.
- Determining and documenting best practices for configuring Kubernetes when running on vSphere infrastructure.
- Determining and documenting best practices for configuring VMware infrastructure when supporting Kubernetes.
- Determining and documenting best practices for using common configuration management and orchestration tools for deploying and managing Kubernetes on VMware infrastructure.
- Discussing bugs and feature requests recorded as Kubernetes or VMware cloud provider issues on GitHub. These issues should be tagged with ``sig/vmware``.

A directory of specialized subject area GitHub teams for issues, PRs, and design proposals is defined in [SIG GitHub Teams](https://github.com/kubernetes/community/tree/master/sig-vmware#github-teams).

#### Code, Binaries and Services

- vSphere Cloud Provider
- vRealize Automation Cluster API Provider

Code locations and projects are defined in [SIG Subprojects](https://github.com/kubernetes/community/tree/master/sig-vmware#subprojects)

#### Cross-cutting and Externally Facing Processes

SIG VMware serves to bring together members of the VMware and Kubernetes community to maintain, support and provide guidance for running Kubernetes on VMware platforms.

The VMware SIG will provide a forum for hosting related architectural planning and discussion. Associated activities related to development, testing, and documentation will be tracked and reported by the SIG.
- The vSphere platform’s availability and resource management can create pools that extend across Kubernetes worker node, management node, and namespace boundaries. Unique features of the vSphere platform related to high availability, load balancing, and resource management (storage, networking, compute) can supplement existing capabilities of Kubernetes, container runtime, and OS functionality.
- We drive changes that cross existing SIG boundaries by collaborating with the appropriate SIGs (eg: node, schedule, scalability, etc.)
- We co-own vSphere-specific code related to cluster and machine provisioning with sig-cluster-lifecycle and the cluster-api subproject.
- We co-own the cloud-provider-vsphere code with sig-cloud-provider
- We define requirements, interfaces, and end-to-end tests related to vSphere-specific storage with sig-storage


### Out of scope

The VMware SIG is not for discussing bugs or feature requests outside the scope of Kubernetes. For example, the VMware SIG should not be used to discuss or resolve support requests related to VMware products. It should also not be used to discuss topics that other, more specialized SIGs own (to avoid overlap).

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

Associated subprojects are created following the Federation of Subprojects procedure described in [sig-governance]

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md