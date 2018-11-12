# SIG API Machinery Charter

This charter is a WIP.

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG API Machinery is responsible for the development and enhancement of Kubernetes master node.  The scope covers API server, persistence layer (etcd), controller manager, cloud controller manager, CustomResourceDefinition, scheduler and webhooks.

### In scope

#### Code, Binaries and Services

All aspects of API server, API registration and discovery, generic API CRUD semantics, admission control, encoding/decoding, conversion, defaulting, persistence layer (etcd), OpenAPI, CustomResourceDefinition, webhooks, garbage collection, namespace lifecycle, and client libraries.

#### Cross-cutting and Externally Facing Processes

N/A

### Out of scope

The contents of individual APIs are owned by SIG Architecture

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

N/A

### Additional responsibilities of Tech Leads

N/A

### Deviations from [sig-governance]

N/A

### Subproject Creation

SIG Auth delegates subproject approval to Technical Leads. See [Subproject creation - Option 1.]

[Subproject creation - Option 1.]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
