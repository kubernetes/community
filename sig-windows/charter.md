# SIG Windows Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The scope of SIG Windows is the operation of Kubernetes on the Windows operating system.
This includes maintaining the interface between Kubernetes and containers on Windows
as well as maintaining the pieces of Kubernetes (e.g. the kube-proxy) where there is a
Windows specific implementation.

### In scope

#### Code, Binaries and Services

- Windows specific code in all parts of the codebase
- Testing of Windows specific features and clusters

#### Cross-cutting and Externally Facing Processes

- Work with other SIGs on areas where Windows and Linux (and possibly other OSes in the future) deviate from one another in terms of functionality.
- Providing community guidance for external tools and paradigms necessary for production windows k8s administration
- Helping to transition from CSI/CNI tooling workarounds to the emerging (windows privileged containers)[https://github.com/ambguo/enhancements/blob/master/keps/sig-windows/1981-windows-privileged-container-support/20201981-windows-privileged-container-support.md] initiative

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

None

### Additional responsibilities of Tech Leads

None

### Deviations from [sig-governance]

None

### Subproject Creation

Federation of Subprojects

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
