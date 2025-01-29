# SIG Node Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Node is responsible for the components that support the controlled
interactions between pods and host resources.  We manage the lifecycle of pods
that are scheduled to a node.  We focus on enabling a broad set of workload
types, including workloads with hardware specific or performance sensitive requirements.  We maintain
isolation boundaries between pods on a node, as well as the pod and the host.  We
aim to continuously improve node reliability.

### In scope

SIG [readme]

#### Code, Binaries and Services

- Kubelet and its features
- Pod API and Pod behaviors (with [sig-architecture](../sig-architecture))
- Node API (with [sig-architecture](../sig-architecture))
- Node controller
- Node level performance and scalability (with [sig-scalability](../sig-scalability))
- Node reliability (problem detection and remediation)
- Node lifecycle management (with [sig-cluster-lifecycle](../sig-cluster-lifecycle))
- Container runtimes
- Device management
- Image management
- Node-level resource management (with [sig-scheduling](../sig-scheduling))
- Hardware discovery
- Issues related to node, pod, container monitoring (with [sig-instrumentation](../sig-instrumentation))
- Node level security and Pod isolation (with [sig-auth](../sig-auth))
- Host OS and/or kernel interactions (to a limited extent)

#### Cross-cutting and Externally Facing Processes

- CRI [validation] and [testing policy]
- Node [test grid]

### Out of scope

- network management ([sig-network](../sig-network))
- persistent storage management ([sig-storage](../sig-storage))

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Technical leads seeded by legacy SIG chairs from existing subproject owners

### Additional responsibilities of Tech Leads

None

### Deviations from [sig-governance]

None

### Subproject Creation

SIG Technical Leads


[testing policy]: /contributors/devel/sig-node/cri-testing-policy.md
[test grid]: https://testgrid.k8s.io/sig-node#Summary
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[readme]: https://github.com/kubernetes/community/tree/master/sig-node
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
