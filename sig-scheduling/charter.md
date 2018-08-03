# SIG Scheduling Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Scheduling is responsible for the components that make Pod placement decisions.
We build Kubernetes schedulers and scheduling features for Pods. We design and
implement features that allows users to customize placement of Pods on the nodes
of a cluster. These features include those that improve reliability of workloads,
more efficient use of cluster resources, and/or enforces placement policies. 

### In scope

Link to SIG section in [sigs.yaml]

#### Code, Binaries and Services

- Scheduling related features (e.g. Node Affinity)
- Kube-scheduler performance and scalability (with [sig-scalability](../sig-scalability))
- Kube-scheduler reliability (problem detection and remediation)
- Pod scheduling APIs (with [sig-api-machinery](../sig-api-machinery))
- Node resource management (with [sig-node](../sig-node))
- Cluster resource management (with [wg-resource-management](../wg-resource-management))
- Pod scheduling policies (with [wg-policy](../wg-poicy))

#### Cross-cutting and Externally Facing Processes

- Kube-scheduler [test grid] and [perf dashboard]

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

Pick one:

1. SIG Technical Leads (x)
2. Federation of Subprojects

[test grid]: https://k8s-testgrid.appspot.com/sig-scheduling#Summary
[perf dashboard]: http://perf-dash.k8s.io/
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1434
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md