# SIG etcd Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Scope

Owns the etcd project and how it is used by Kubernetes.

### In scope

#### Code, Binaries and Services

- Development of [etcd] and other repositories under [etcd-io organization]
- Maintenance of [etcd image] packaged with Kubernetes

[etcd]: https://github.com/etcd-io/etcd
[etcd-io organization]: https://github.com/etcd-io
[etcd image]: https://github.com/kubernetes/kubernetes/tree/master/cluster/images/etcd

#### Cross-cutting and Externally Facing Processes

- Specifying, testing and improving [The Implicit Kubernetes-ETCD Contract]
- Release process of etcd and other binaries belonging to [etcd-io organization]

[The Implicit Kubernetes-ETCD Contract]: https://docs.google.com/document/d/1NUZDiJeiIH5vo_FMaTWf0JtrQKCx0kpEaIIuPoj9P6A/edit?usp=sharing

### Out of scope

- Structure of data stored in etcd by Kubernetes components is owned by SIG API Machinery

## Roles and Organization Management

This SIG follows the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Tech Leads

- Release of etcd and other binaries belonging to [etcd-io organization]

### Deviations from [sig-governance]

- SIG etcd's participation in the Kubernetes release cycle is limited by etcd having a different schedule for its releases.
- SIG etcd communication utilizes pre-existing forums for communication:
  - Email: [etcd-dev](https://groups.google.com/forum/?hl=en#!forum/etcd-dev).
  - Slack: [#etcd](https://kubernetes.slack.com/messages/C3HD8ARJ5/details/) channel on Kubernetes.
- SIG etcd contributing instructions ([CONTRIBUTING.md]) be defined in etcd project.

[CONTRIBUTING.md]: https://github.com/etcd-io/etcd/blob/main/CONTRIBUTING.md

### Deviations from [kubernetes-repositories]

- SIG etcd repositories live in github.com/etcd-io
- SIG etcd repositories should (but not must) adopt merge bot, Kubernetes PR commands/bot.
- SIG etcd repositories will follow [rules for donated repositories].

[kubernetes-repositories]: https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md#sig-repositories
[rules for donated repositories]:  https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md#rules-for-donated-repositories

### Subproject Creation

By SIG Technical Leads
