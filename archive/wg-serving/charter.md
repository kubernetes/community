# WG Serving Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [wg-governance].

[Kubernetes Charter README]: /committee-steering/governance/README.md

## Scope

Discuss and enhance serving workloads on Kubernetes, specifically focusing on
hardware-accelerated AI/ML inference. The working group will focus on the novel
challenges of compute-intensive online inference. Scenarios solving use cases
involving non-fungible accelerators will be prioritized over solutions against
generic CPU. However, all improvements should, where possible, benefit other
serving workloads like web services or stateful databases, be usable as
primitives by multiple ecosystem projects, and compose well into the workflows
of those deploying models to production. The Working Group Batch has a similar
scope. The difference in scope by a simplified definition is that the Serving WG
will generally concentrate on the workloads where Pods are running with
restartPolicy=Always, while WG Batch will generally be looking at Pods with the
restartPolicy=OnFailure. There are edge cases to this definition, but it creates
an easy enough framework to differentiate the scope of these two Working Groups.

### In scope

- Gather requirements for serving workloads (inference primarily, but benefiting
  other non-batch use cases where possible) that have broad community alignment
  from practitioners, distros, and vendors. Provide concrete input to other SIGs
  and WGs around needs for identified requirements. Do it in partnership
  with existing ecosystem projects like kServe, Seldon, Kaito, and
  others to identify, extract, or implement common shared problems (like Kueue
  abstracted deferred scheduling for multiple batch frameworks).
- Specific areas of improvement include:
  - Directly improve key kubernetes workload controllers when used with
    accelerators and the most common inference serving frameworks and model
    servers.
  - Explore new projects that improve orchestration, scaling, and load balancing
    of inference workloads and compose well with other workloads on Kubernetes
  - Being able to run serving workloads safely while giving up
    available slack capacity to batch frameworks

### Out of scope

- Training and batch inference, which are covered by WG Batch.
- Ability to describe the workflows for serving workloads is out of scope,
  Kubernetes will offer building blocks to MLOps platforms to build those.

## Stakeholders

Stakeholders in this working group span multiple SIGs that own parts of the
code in core kubernetes components and addons.

- SIG Apps as a primary SIG
- SIG Architecture
- SIG Node
- SIG Scheduling
- SIG Autoscaling
- SIG Network
- SIG Instrumentation
- SIG Storage

## Deliverables

The list of deliverables include the following high level features:

- To SIG Apps:
  - Ability to express the model serving workloads with easy to understand logical
    objects with the ability to scale to multi-host 
- To SIG Scheduling and Autoscaling
  - Faster scaling up and down
  - Ability to preempt workloads
- To SIG Node:
  - Runtime support for Pods preemption
  - Runtime support for devices partitioning

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

[wg-governance]: /committee-steering/governance/wg-governance.md

Additionally, the WG commits to maintain a solid communication line between the Kubernetes groups and the wider CNCF community.

## Timelines and Disbanding

As a first mandate, the WG will define a roadmap in the first quarter of operation.
We believe there will be a set of features the Working Group can identify and deliver
that will enable the majority of frameworks operate natively on Kubernetes.

Achieving the aforementioned deliverables, also mentioned in the `In Scope`
section, will allow us to decide when to disband this WG.  There is no
expectations that the Working Group will be converted into SIG long term,
however, there is a chance that a separate project or a sizeable sub-component
of SIG Apps can be created as a result of a Working Group.
