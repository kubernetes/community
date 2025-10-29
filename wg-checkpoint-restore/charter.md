
# WG Checkpoint Restore Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The Checkpoint/Restore Working Group aims to solve the problem of transparently
checkpointing and restoring workloads in Kubernetes, a [functionality discussed
for over five years][kep2008]. The group will deliver the design and
implementation of Checkpoint/Restore functionality in Kubernetes, serving as a
central hub for community information and discussion. This initiative addresses
a wide range of problems, including fault tolerance, improved resource
utilization, and accelerated application startup times.

### In scope

- Identify core Kubernetes checkpoint/restore use cases (e.g., live migration,
  fault tolerance, debugging, snapshotting) and gather stakeholder requirements.
- Investigate and propose Kubernetes APIs for checkpoint/restore operations.
- Work with SIGs for the best integration of checkpoint/restore functionality
  and APIs.
- Provide guidance for developers on checkpoint-friendly app design and
  recommendations for operators on feature management.
- Work closely with relevant upstream projects (CRI-O, containerd, CRIU, gVisor)
  for alignment and integration.
- Revisit the existing implementations to find and remedy possible inefficiencies.
  One example is the existing checkpoint archive format which has already been
  identified as being a major source of slowdown.

### Out of scope

- Not focused on general OS-level checkpointing outside Kubernetes
  pods/containers.
- Will not dictate internal application checkpointing logic; focuses on
  Kubernetes platform orchestration of *container/pod state.

## Stakeholders

Stakeholders in this working group span multiple SIGs that own parts of the
code in core kubernetes components and addons.

- SIG Node
- SIG Scheduling
- SIG Auth
- SIG Apps

## Deliverables

The list of deliverables include the following high level features:

- In the early stage, we mainly want to offer a well-defined location for the
  community to find information, ask questions, and discuss the next steps of
  enabling checkpoint and restore in Kubernetes.

Later:

- Ability to checkpoint and restore a container using kubectl
- Ability to checkpoint and restore a pod using kubectl
- Integration of container/pod checkpointing in scheduling decisions

## Roles and Organization Management

This WG adheres to the Roles and Organization Management outlined in [wg-governance]
and opts-in to updates and modifications to [wg-governance].

[wg-governance]: /committee-steering/governance/wg-governance.md

Additionally, the WG commits to:

- maintain a solid communication line between the Kubernetes groups and the
  wider CNCF community

## Timelines and Disbanding

As a first mandate, the WG will propose a draft roadmap and identify key tasks in the first quarter of operation.

After that, the WG will facilitate collaboration among community members to explore possible APIs and draft proposals for their integration into Kubernetes, which will then be presented to the relevant SIGs.

Achieving the aforementioned deliverables, also mentioned in the `In Scope`
section, will allow us to decide when to disband this WG.  There is no
expectations that the Working Group will be converted into a SIG long term.

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[kep2008]: https://github.com/kubernetes/enhancements/issues/2008
