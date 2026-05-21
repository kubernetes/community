# WG Workload-aware Scheduling Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [wg-governance].

## Scope

Native support for workload-aware and topology-aware scheduling in core Kubernetes. Today, 
these concerns are spread across multiple SIGs and ecosystem projects, and the lack of a 
shared model for expressing workload-level requirements and infrastructure topology makes it 
difficult to coordinate placement decisions consistently and efficiently.

The goal of this working group is to establish a common scheduling substrate and native APIs 
in core Kubernetes to address these gaps.

### In scope

- Enable core APIs for expressing workload-level scheduling requirements, including coupling
  constraints, topology preferences, and collective resource needs.
- Enable disruption handling that respects workload boundaries and semantics (e.g. understanding
  that evicting one member of a tightly-coupled workload group may disrupt the entire workload).
- Enable higher-level controllers to consume workload-aware scheduling primitives.
- Enable capacity provisioning that understands collective workload requirements.
- Preserve scheduling latency and throughput for existing workloads while achieving acceptable
  performance and scalability for workload-aware scheduling paths.
- Cross-SIG design alignment, use case gathering, and terminology alignment across stakeholder
  SIGs.

### Out of scope

- Implementing vendor-specific or cloud-provider-specific scheduling logic.
- Defining new workload controller APIs (e.g. equivalents of Deployment, StatefulSet, or
  DaemonSet) for specialized workload types.
- Building a job queueing or quota management system. The WG focuses on providing core 
scheduling primitives that ecosystem projects (e.g., Kueue, Volcano) can build upon 
and consume, rather than replacing them.

## Stakeholders

Stakeholders in this working group span multiple SIGs that own parts of the code in core
Kubernetes components and addons.

- SIG Scheduling
- SIG Autoscaling
- SIG Apps
- WG Device Management
- WG Batch

Additionally, a broad set of end users, infrastructure vendors, cloud providers, and ecosystem
projects have expressed interest in this effort. There are five primary groups of stakeholders
from each of which we expect multiple participants:

- Cloud providers and Kubernetes distribution providers.
- Workload authors and platform teams building AI, HPC, batch, serving, and stateful systems
  on Kubernetes.
- Kubernetes ecosystem projects that help manage workload scheduling (e.g. Cluster Autoscaler,
  Karpenter, Kueue, Volcano, LWS, JobSet, TrainJob).
- End user workload authors that will create workloads that take advantage of workload awareness.
- Cluster operators managing heterogeneous and accelerator-backed infrastructure.

## Deliverables

The WG coordinates the delivery of KEPs and their implementations by the participating SIGs.
Interim artifacts will include documents capturing use cases, requirements, and designs;
however, all of those will eventually result in KEPs and code owned by SIGs.

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

### Additional responsibilities of Chairs

- The Working Group will have designated Chair(s) responsible for guiding discussions and ensuring progress.
- Agendas and meeting notes will be publicly accessible.


## Timelines and Disbanding

The working group will disband when the KEPs resulting from these discussions have reached a
terminal state and all use cases and requirements have been met or implemented. The cross-SIG
coordination gaps that motivated this WG should be resolved, and ongoing work can be fully
owned by individual SIGs without requiring cross-cutting design alignment through this forum.

[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
