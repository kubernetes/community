# SIG Network Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Network is responsible for the components, interfaces, and APIs which expose networking capabilities to Kubernetes
users and workloads. SIG Network also provides some reference implementations of these APIs, for
example kube-proxy as a reference implementation of the Service API.

### In scope

The following topics fall under ownership of this SIG:

- Networking control plane and data paths.
- Network service abstractions.
- Service discovery (DNS).
- Service load balancing (L4, L7).
- Network security and identity.
- Cluster connectivity.
- Cross-cutting concerns such as scalability.
- Metrics and monitoring associated with networking components.

#### Code, Binaries and Services

- Services
  - APIs for defining and grouping networking endpoints.
  - APIs for defining L3/4 loadbalancing.
  - Reference implementation (kube-proxy).
- Ingress
  - APIs for defining L7 loadbalancing.
- APIs for defining network policy.
- Cluster DNS.
- Integration points with networking implementations (e.g. CNI integration).
- Container runtime interface (CRI) (With [sig-node]).
- Cloud provider network integrations (With [sig-cloud-provider]).

#### Cross-cutting and Externally Facing Processes

### Out of scope

- The CNI specification itself, which is maintained outside the Kubernetes project
- Particular implementations of the CNI specification
- Particular implementations of the NetworkPolicy API
- Particular implementations of the Ingress API

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- None

### Additional responsibilities of Tech Leads

- None

### Deviations from [sig-governance]

- None

### Subproject Creation

SIG Technical Leads

[sig-cloud-provider]: https://github.com/kubernetes/community/tree/master/sig-cloud-provider
[sig-node]: https://github.com/kubernetes/community/tree/master/sig-node

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-network/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
