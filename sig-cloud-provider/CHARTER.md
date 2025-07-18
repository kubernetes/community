# SIG Cloud Provider Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Cloud Provider’s mission is to simplify, develop, and maintain cloud provider integrations as extensions, or add-ons, to Kubernetes clusters.

### In scope

#### Areas of Focus

- Cloud provider specific integrations and extension points that are not already covered by a more specific sig such as storage or networking.
- APIs/interfaces for efficiently provisioning/de-provisioning cloud resources (nodes, routes, load balancers, etc)
- Configuration of cluster components to enable cloud provider integrations
- Testing and testing frameworks to ensure vendor neutrality across all cloud providers

#### Code, Binaries and Services

The SIG offers standardization across cloud-provider-* repos that are owned by the sig. We establish basic structure and tooling expectations to help new contributors to understand the code and how to contribute.

- the [common interfaces](https://github.com/kubernetes/cloud-provider/blob/master/cloud.go) consumed by all cloud providers
- the [cloud-controller-manager](https://github.com/kubernetes/kubernetes/tree/master/cmd/cloud-controller-manager), which acts as the “out-of-tree” cloud provider component for clusters.
- core controllers (started by the cloud-controller-manager) that interact with cloud provider resources
- all [cloud provider repositories](https://github.com/kubernetes?utf8=%E2%9C%93&q=cloud-provider-&type=&language=) under the Kubernetes organization
- [e2e tests for cloud provider specific](https://github.com/kubernetes/kubernetes/tree/master/test/e2e/cloud) functionality
- the subproject [apiserver-network-proxy](https://github.com/kubernetes-sigs/apiserver-network-proxy), which is an extensible system which controls network traffic from the Kube API Server.
- all the subprojects formerly owned by [SIG-AWS](https://github.com/kubernetes/community/tree/master/sig-aws#subprojects), [SIG-AZURE](https://github.com/kubernetes/community/tree/master/sig-azure#subprojects), [SIG-GCP](https://github.com/kubernetes/community/tree/master/sig-gcp#subprojects), [SIG-IBMCloud](https://github.com/kubernetes/community/tree/master/sig-ibmcloud#subprojects), [SIG-Openstack](https://github.com/kubernetes/community/tree/master/sig-openstack#subprojects), [SIG-VMware](https://github.com/kubernetes/community/tree/master/sig-vmware#subprojects).
- any new subproject that is cloud provider specific, unless there is another SIG already sponsoring it.

#### Cross-cutting and Externally Facing Processes

- This SIG works with SIG Testing & SIG Release to ensure that cloud providers are actively testing & reporting test results to testgrid.
- This SIG works with SIG Docs to provide user-facing documentation on configuring Kubernetes clusters with cloud provider integration enabled.
- This SIG works with new cloud providers in the ecosystem that want to host their code in the kubernetes-sigs organization and have an interest in contributing back.
- A portion of the apiserver-network-proxy code needs to be compiled into the apiserver, which overlaps with SIG API Machinery.
- This SIG actively engages with SIGs owning other external components of Kubernetes (CNI, CSI, other networking and storage, apiserver, and similar) to ensure a consistent integration story for users.
- This SIG collaborates to create infrastructure-specific endpoints and extensions. This can entail participation in working groups or sponsorship of subprojects.

### Out of scope

- This SIG does not act as a line of support for Kubernetes users running their clusters on any cloud provider, though many members of the SIG represent cloud providers and are actively engaged with users.
- This SIG does not address features/bugs pertaining to cloud providers outside the scope of Kubernetes integrations (e.g. when will instance type X be available on cloud provider Y?)

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Selecting/prioritizing work to be done for a milestone
- Hosting the weekly SIG meeting, ensure that recordings are uploaded in a timely fashion.
- Ensuring that the breakout sessions the SIG hosts during the week have chairs.
- Organizing SIG sessions at KubeCon events (intro / deep dive sessions).
- Creating roadmaps for a given year or release, or reviewing and approving technical implementation plans (e.g. KEPs) in coordination with both SIG Cluster Lifecycle contributors and other SIGs.

### Deviations from [sig-governance]

- There should be no more than 1 chair from a single company. This ensures decisions are not being made in favor of a single provider/company.
- As SIG cloud provider contains a number of subprojects, the SIG has empowered subproject leads with a number of additional responsibilities, including but not limited to:
    * Releases: The subproject owners are responsible for determining the subproject release cadence, producing releases, and communicating releases with SIG Release and any other relevant SIG.
    * Backlog grooming: The subproject owners are responsible for ensuring that the issues for the subproject are correctly associated with milestones and that bugs are triaged in a timely manner.
PR timeliness: The subproject owners are responsible for ensuring that active pull requests for the subproject are addressed in a timely manner.
    * Repository ownership: The subproject owners are given admin permissions to repositories under the subproject. For example, the owners of the Azure subproject are given admin access to the `sigs.k8s.io/cloud-provider-azure` repository.


### Subproject Creation

Federation of Subprojects

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
