# SIG Storage Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Storage is responsible for ensuring that different types of file and block storage
(whether ephemeral or persistent, local or remote) are available wherever a container is
scheduled (including provisioning/creating, attaching, mounting, unmounting, detaching,
and deleting of volumes), storage capacity management (container ephemeral storage
usage, volume resizing, etc.), influencing scheduling of containers based on storage
(data gravity, availability, etc.), and generic operations on storage (snapshoting, etc.).

### In scope

Some notable examples of features owned by SIG Storage:

* Persistent Volume Claims and Persistent Volumes
* Storage Classes and Dynamic Provisioning
* Kubernetes volume plugins
* Container Storage Interface (CSI)
* Secret Volumes, ConfigMap Volumes, DownwardAPI Volumes, EmptyDir Volumes (co-owned with SIG-Node)

#### Code, Binaries and Services

* Kubernetes internal controllers and APIs responsible for exposing file and block storage to Kubernetes workloads.
* Kubernetes external sidecar containers and binaries required for exposing file and block storage to Kubernetes workloads.
* Interfaces required for exposing file and block storage to Kubernetes workloads.
* Unit, Integration, and End-to-End (E2E) Tests validating and preventing regressions in the above.

#### Cross-cutting and Externally Facing Processes

* Defining interface and requirements for connecting third party storage systems to Kubernetes.

### Out of scope

SIG Storage is not responsible for

* Data path of remote storage (GCE PD, AWS EBS, NFS, etc.)
  * How bits are transferred.
  * Where bits are stored.
* Container writable layer (SIG Node handles that).
* The majority of storage plugins/drivers (generally owned by storage vendors).

## Roles and Organization Management

SIG Storage adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

SIG Storage does not have separate tech leads: SIG Storage chairs serve as tech leads.

### Subproject Creation

SIG Storage delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
