# SIG UI Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG-UI covers GUI-related aspects of the Kubernetes project. Efforts are centered around the Kubernetes Dashboard: a general purpose, web-based UI for Kubernetes clusters. It allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

#### Code, Binaries and Services

- [Kubernetes Dashboard](https://github.com/kubernetes/dashboard) (Archived)
- [Headlamp](https://github.com/kubernetes-sigs/headlamp)

#### Cross-cutting and Externally Facing Processes

- Cutting a new release of the Dashboard

### Out of Scope

Tools that contributors use in support of the project (eg. Prow, Test Grid)

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance] and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

SIG UI governance is structured according to the [ROLES.md] document within the [Kubernetes/Dashboard] repo

### Subproject Creation

SIG UI delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
[Kubernetes/Dashboard]: https://github.com/kubernetes/dashboard
[ROLES.md]: https://github.com/kubernetes/dashboard/blob/master/ROLES.md
