# SIG Cluster Lifecycle Charter

## Scope

SIG Cluster Lifecycle’s objective is to simplify creation, configuration, upgrade, downgrade, and teardown of Kubernetes clusters and their components.

### In scope

The following topics fall under ownership of this SIG:

- Improving the Kubernetes user experience for cluster administration.
- Tools that assist in the creation, configuration, upgrade, downgrade, and teardown of Kubernetes control plane components. 
- Portable APIs for provisioning, configuration, upgrade/downgrade, and de-provisioning of nodes.
- Tools that assist in management of configuration of Kubernetes components.
- The configuration of core add-ons that are required for cluster bootstrapping.

#### Code, Binaries and Services

- Everything that falls in the scope of the SIG.
- Tools that are provider specific implementation for infrastructure management.
- Core add-ons (e.g. DNS) that are required for cluster bootstrapping.

#### Cross-cutting and Externally Facing Processes

- This SIG works closely with SIG Release during the end of the cycle, because of how the release of kubeadm is currently managed. The process requires a lock-step coordination which is outlined [here](https://github.com/kubernetes/kubeadm/blob/master/docs/release-cycle.md).
- The SIG recommends and verifies compatibility of critical cluster add-ons for networking, network policy, service discovery, etc. The SIG maintains the health check of container images for some add-ons that are required for cluster bootstrapping. While the SIG could provide support to users that have add-on related issues, the SIG can decide to delegate issues to the add-on maintainers or other SIGs.
- The SIG owns the framework for upgrade / downgrade testing of Kubernetes. However, the SIG does not own the individual tests for each feature. While the SIG is often involved in triaging and debugging upgrade or downgrade test failures, the SIG will normally delegate issues to the feature owner for the failing test.
- The SIG collaborates regularly with SIG Auth in an effort to follow best practices in order to promote secure default clusters.
- The SIG co-owns cloud provider specific code related to cluster and machine provisioning with the respective SIGs for each cloud provider but does not own the cloud controller manager or any other provider specific code.

### Out of scope

- Networking related issues (see [sig-network](../sig-network)).
- User interface, or user experience, issues other than cluster bootstrapping or management (see [sig-ui](../sig-ui) and  [sig-cli](../sig-cli)).
- Node related issues (see [sig-node](../sig-node)).
- Kubernetes control plane issues:
   - Control plane component related issues (see [sig-api-machinery](../sig-api-machinery)).

## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Selecting features for a given milestone.
- Hosting the weekly SIG meeting, ensure that recordings are uploaded in a timely fashion.
- Ensuring that the breakout sessions the SIG hosts during the week have chairs.
- Organizing SIG sessions at KubeCon events (intro / deep dive sessions).
- Creating roadmaps for a given year or release, or reviewing and approving technical implementation plans (e.g. KEPs) in coordination with both SIG Cluster Lifecycle contributors and other SIGs.

### Deviations from [sig-governance]

- As SIG cluster lifecycle contains a number of subprojects, the SIG has empowered subproject leads with a number of additional responsibilities, including but not limited to:
   * Releases: The subproject owners are responsible for determining the subproject release cadence, producing releases, and communicating releases with SIG Release and SIG Cluster Lifecycle.
   * Backlog grooming: The subproject owners are responsible for ensuring that the issues for the subproject are correctly associated with milestones and that bugs are triaged in a timely manner.
   * PR timeliness: The subproject owners are responsible for ensuring that active pull requests for the subproject are addressed in a timely manner.

### Subproject Creation

- Federation of Subprojects

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
