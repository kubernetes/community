# SIG Multicluster Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The scope of SIG Multicluster is limited to the following subprojects:

- The [cluster-registry](https://github.com/kubernetes/cluster-registry)
- Kubernetes federation:
  - [Federation v2](https://github.com/kubernetes-sigs/federation-v2)
  - [Federation v1](https://github.com/kubernetes/federation)
- [Kubemci](https://github.com/GoogleCloudPlatform/k8s-multicluster-ingress)

### In scope

See [SIG README].

#### Code, Binaries and Services

SIG Multicluster code and binaries are limited to those from one of the SIG subprojects.

#### Cross-cutting and Externally Facing Processes

- Consult with other SIGs and the community on how the in-scope mechanisms
  should work and integrate with other areas of the wider Kubernetes ecosystem

### Out of scope

- Software that creates or manages the lifecycle of Kubernetes clusters

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

SIG Multicluster delegates subproject approval to Technical Leads. See [Subproject creation - Option 1]. 

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-multicluster/README.md#subprojects
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L1042
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
[SIG README]: https://github.com/kubernetes/community/blob/master/sig-multicluster/README.md