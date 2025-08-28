# WG Batch Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [wg-governance].

[Kubernetes Charter README]: /committee-steering/governance/README.md

## Scope

Discuss and enhance the support for Batch (eg. HPC, AI/ML, data analytics, CI)
workloads in core Kubernetes. We want to unify the way users deploy batch
workloads to improve portability and to simplify supportability for Kubernetes
providers.

### In scope

- To reduce fragmentation in the k8s batch ecosystem: congregate leads and users from
  different external and internal projects and user groups (CNCF TAGs, k8s sub-projects
  focused on batch-related features such as topology-aware scheduling) in the batch ecosystem to
  gather requirements, validate designs and encourage reutilization of core kubernetes APIs.
- The following recommendations for enhancements:
  - Additions to the batch API group, currently including Job and CronJob resources
    that benefit batch use cases such as HPC, AI/ML, data analytics and CI.
  - Primitives for job-level queueing, not limited to the k8s Job resource. Long-term,
    this could include multi-cluster support.
  - Primitives to control and maximize utilization of resources in fixed-size clusters
    (on-prem) and elastic clusters (cloud).
  - Runtime and scheduling support for specialized hardware (GPUs, NUMA, RDMA, etc.)

### Out of scope

- Addition of new API kinds that serve a specialized type of workload. The focus
  should be on general APIs that specialized controllers can build on top of.
- Uses of the batch APIs as support for serving workloads (eg. backups,
  upgrades, migrations). These can be served by existing SIGs.
- Proposals that duplicate the functionality of core kubernetes components
  (job-controller, kube-scheduler, cluster-autoscaler).
- Job workflows or pipelines. Mature third party frameworks serve these
  use cases with the current kubernetes primitives. But additional primitives
  to support these frameworks could be in scope.

## Stakeholders

Stakeholders in this working group span multiple SIGs that own parts of the
code in core kubernetes components and addons.

- Apps
- Autoscaling
- Node
- Scheduling

## Deliverables

The list of deliverables include the following high level features:

- To SIG Apps:
  - Updated Job API that fulfills the needs of a wider range of batch applications.
  - A performant job controller that can scale to thousands of pods per minute.
- To SIG Scheduling and Autoscaling
  - A set of APIs to support job queueing, a framework to support different
    queueing policies and a ready-to-use implementation as a subproject.
  - Scheduling plugin(s) to support different batch needs.
- To SIG Autoscaling:
  - Capabilities for job-level provisioning.
- To SIG Node:
  - Runtime support for specialized hardware.

## Roles and Organization Management

This wg adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

[wg-governance]: /committee-steering/governance/wg-governance.md

Additionally, the wg commits to:

- maintain a solid communication line between the Kubernetes groups and the wider CNCF community;

## Timelines and Disbanding

As a first mandate, the wg will define a roadmap in the first quarter
of operation. We envision three timelines for the exit criteria, the focus will
be on early exit, but a determination on whether or not to go beyond
that is left until we reach that milestone.

1. Early exit: define "recommendations" for the deliverables mentioned above, those
   recommendations would be left to the respective sigs to implement. The WG could
   start implementing those recommendations in the context of the owning sig to generate
   some momentum.
2. Milestone 2, Late exit: The WG continues the implementation of the recommendations until they reach GA,
   and then disband.
2. Convert to SIG: The WG observes a constant influx of requirements for the artifacts and there
   is the risk that the SIGs don't have enough capacity to maintain them.
   Then, the WG will propose the graduation into a SIG, taking ownership of the
   APIs, controllers and scheduling plugins.
