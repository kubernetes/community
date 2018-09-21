# SIG IBMCloud Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The IBMCloud SIG enables activities and discussion around building, deploying, maintaining, supporting,
and using Kubernetes on IBM Public and Private Clouds.

### In scope
- Determining and documenting best practices for configuring Kubernetes on IBM Cloud Kubernetes Service (IKS).
- Determining and documenting best practices for configuring Kubernetes on IBM Cloud Private (ICP).
- Discussing IKS and ICP tracking of Kubernetes features and releases.
- Utilizing Kubernetes and related CNCF projects (e.g. Helm, Istio) by IKS and ICP.
- Discussing bugs and feature requests recorded as Kubernetes upstream issues on GitHub. These issues should be tagged with `sig/ibmcloud`.

#### Code, Binaries and Services

The work to have a cloud provider specific public code repository is in progress. This section will be updated once the work is complete. Kubernetes upstream code that
is directly related to IKS or ICP issues or features can be discussed.

### Out of scope

* Internal or commercial aspects of IKS and ICP.

## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance], with exception that the SIG only
has the Chair role at present. Chairs must also fulfill all of the responsibilities of the Tech Lead role as outlined in [sig-governance].

### Subproject Creation
Associated subprojects are created following the `by SIG Technical Leads` option procedure described in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[SIG README]: https://github.com/kubernetes/community/blob/master/sig-ibmcloud/README.md
