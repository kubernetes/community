# WG Device Management Charter

This charter adheres to the conventions described in the [Kubernetes Charter
README] and uses the Roles and Organization Management outlined in
[wg-governance].

## Scope

Enable simple and efficient configuration, sharing, and allocation of
accelerators and other specialized devices. This working group focuses on the
APIs, abstractions, and feature designs needed to configure, target, and share
the necessary hardware for both batch and serving (inference) workloads.

### In scope

- Enable efficient utilization of specialized hardware devices. This includes
  sharing one or more resources effectively (many workloads sharing a pool of
  devices), as well as sharing individual devices effectively (several workloads
  dividing up a single device for sharing).
- Enable workload authors to specify “just enough” details about their workload
  requirements to ensure it runs optimally, without having to understand exactly
  how the infrastructure team has provisioned the cluster.
- Enable the scheduler to choose the correct place to run a workload the vast
  majority of the time (rejections should be extremely rare).
- Enable cluster autoscalers and other node auto-provisioning components to
  predict whether creating additional resources will satisfy workload needs,
  before provisioning those resources.
- Enable the shift from “pods run on nodes” to “workloads consume capacity”.
  This allows Kubernetes to provision sets of pods on top of sets of nodes and
  specialized hardware, while taking into account the relationships between
  those infrastructure components.
- Enable in-node devices as well as network-accessible devices.
- Minimize workload disruption due to hardware failures.
- Address fragmentation of accelerator due to fractional use.
- Additional problems that may be identified and deemed in scope as we gather
  use cases and requirements from WG Serving, WG Batch, and other stakeholders.
- Address all of the above while with a simple API that is a natural extension
  of the existing Kubernetes APIs, and avoids or minimizes any transition
  effort.

### Out of Scope

- Higher-level workload controller APIs (for example, the equivalent of
  Deployment, StatefulSet, or DaemonSet) for specific types of workloads.
- General resource management requirements not related to devices.

## Deliverables

The WG will coordinate the delivery of KEPs and their implementations by the
participating SIGs. Interim artifacts will include documents capturing use
cases, requirements, and designs; however, all of those will eventually result
in KEPs and code owned by SIGs.

Specifically, we expect to need:

- APIs for publishing resource capacity of in-node and network-accessible
  devices, as well as sample code to ease creation of drivers to populate this
  information.
- APIs for specifying workload resource requirements with respect to devices.
- APIs, algorithms, and implementations for allocating access to and resources on devices, as well as
  persisting the results of those allocations.
- APIs, algorithms, and implementations for allowing adminstrators to control
  and govern access to devices.

## Stakeholders

- SIG Architecture
- SIG Autoscaling
- SIG Network
- SIG Node
- SIG Scheduling

Additionally a broad set of end users, device vendors, cloud providers,
Kubernetes distribution providers, and ecosystem projects (particularly
autoscaling-related projects) have expressed interest in this effort. There are
five primary groups of stakeholders from each of which we expect multiple participants:

- Device vendors that manufacture accelerators and other specialized hardware
  which they would like to make available to Kubernetes users.
- Kubernetes distribution and managed offering providers that would like to make
  specialized hardware available to their users.
- Kubernetes ecosystem projects that help manage workloads utilizing these
  accelerators (e.g., Karpenter, Kueue, Volcano)
- End user workload authors that will create workloads that take advantage of
  the specialized hardware.
- Cluster administrators that operate and govern clusters containing the
  specialized hardware.

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

## Exit Criteria

The working group will disband when the KEPs resulting from these discussions
have reached a terminal state.

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
