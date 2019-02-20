# SIG API Machinery Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG API Machinery is responsible for the development and enhancement of Kubernetes cluster control plane.  The scope covers API server, persistence layer (etcd), controller manager, cloud controller manager, CustomResourceDefinition and webhooks.

### In scope

#### Code, Binaries and Services

All aspects of 
* API server 
* API registration and discovery
* Generic API CRUD semantics
* Admission control
* Encoding/decoding
* Conversion
* Defaulting
* Persistence layer (etcd)
* OpenAPI
* The informer libraries
* CustomResourceDefinition
* Webhooks
* Garbage collection
* Namespace lifecycle
* Client libraries

#### Cross-cutting and Externally Facing Processes

Client library releases

### Out of scope

The contents of individual APIs are owned by SIG Architecture

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

Technical leads seeded by legacy SIG chairs from existing subproject owners

### Additional responsibilities of Tech Leads

N/A

### Deviations from [sig-governance]

#### Sub-project Repo Creation

The following individuals may approve kubernetes-sigs repo creation requests to be owned by any api-machinery
sub-projects:

- @cheftako
- @sttts

The following individuals may approve kubernetes-sigs repo creation requests to be owned by specific api-machinery
sub-projects:

- server-sdk
  - @pwittrock


### Subproject Creation

SIG delegates subproject approval to Technical Leads. See [Subproject creation - Option 1.]

[Subproject creation - Option 1.]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-api-machinery/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
