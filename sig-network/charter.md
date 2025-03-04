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
- Multi-cluster networking (shared responsibility with [sig-multicluster]).

#### Code, Binaries and Services

- Services
  - APIs for defining and grouping network endpoints (i.e. [EndpointSlices], or the older [Endpoints] API)
  - APIs for defining L3/4 loadbalancing (i.e. [Service], [Gateway API])
  - Reference implementations (i.e. [kube-proxy]).
- Ingress
  - APIs for defining ingress loadbalancing (i.e. [Ingress], [Gateway API], [Gateway API Inference Extension])
  - API Implementations (i.e. [ingress-nginx], [InGate], [Blixt])
- Network Policy
  - APIs for defining network policies (i.e. [NetworkPolicy], [AdminNetworkPolicy], [BaselineAdminNetworkPolicy])
  - Reference implementations (i.e. [kube-network-policies])
- Cluster DNS.
- Integration points with networking implementations (i.e. [Container Network Interface (CNI)][CNI]).
- [Container Runtime Interface (CRI)][CRI] (With [sig-node]).
- Cloud provider network integrations (With [sig-cloud-provider]).

#### Cross-cutting and Externally Facing Processes

### Out of scope

- The [CNI] specification itself, which is maintained outside the Kubernetes project
- Particular implementations of the [CNI] specification

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

[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

[sig-cloud-provider]: https://github.com/kubernetes/community/tree/master/sig-cloud-provider
[sig-node]: https://github.com/kubernetes/community/tree/master/sig-node
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-network/README.md#subprojects
[sig-multicluster]: https://github.com/kubernetes/community/blob/master/sig-multicluster/README.md

[EndpointSlices]: https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/
[Endpoints]: https://kubernetes.io/docs/concepts/services-networking/service/#endpoints
[Service]: https://kubernetes.io/docs/concepts/services-networking/service/
[kube-proxy]: https://kubernetes.io/docs/concepts/overview/components/#kube-proxy

[Ingress]: https://kubernetes.io/docs/concepts/services-networking/ingress/
[Gateway API]: https://gateway-api.sigs.k8s.io/
[Gateway API Inference Extension]: https://github.com/kubernetes-sigs/gateway-api-inference-extension
[ingress-nginx]: https://github.com/kubernetes/ingress-nginx/
[InGate]: https://github.com/kubernetes-sigs/ingate
[Blixt]: https://github.com/kubernetes-sigs/blixt

[NetworkPolicy]: https://kubernetes.io/docs/concepts/services-networking/network-policies/
[AdminNetworkPolicy]: https://network-policy-api.sigs.k8s.io/api-overview/#the-adminnetworkpolicy-resource
[BaselineAdminNetworkPolicy]: https://network-policy-api.sigs.k8s.io/api-overview/#the-baselineadminnetworkpolicy-resource
[kube-network-policies]: https://github.com/kubernetes-sigs/kube-network-policies

[CNI]: https://kubernetes.io/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-network-model
[CRI]: https://kubernetes.io/docs/concepts/architecture/cri/
