# SIG Autoscaling Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Autoscaling deals with any and all subjects related to the automatic adjustment of the sizing of cluster resources.  This includes horizontal autoscaling of replica counts, vertical autoscaling and autosizing of pod resources, and the autoscaling of clusters.

We collaborate with SIG Instrumentation to develop metrics APIs to suit those needs, but do not own those metrics APIs. We collaborate with SIG Scheduling and other SIGs as needed, to make sure that autoscaling decisions are consistent with the rest of the Kubernetes ecosystem.

### In scope

#### Code, Binaries, Docs, and Services

- [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
- [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)
- [Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)
- Autoscalers for system components
- Related APIs

#### Cross-cutting and Externally Facing Processes

- Metric pipelines for HPA and VPA.
- CA interactions with Scheduler.
- CA interactions with GPUs.

### Out of scope

- Support for Custom Metrics for HPA in various monitoring solutions.

## Roles and Organization Management

This sig follows and adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Manage and curate the project boards associated with all sub-projects ahead of every SIG meeting so they may be discussed
- Ensure the agenda is populated 24 hours in advance of the meeting, or the meeting is then cancelled
- Report the SIG status at events and community meetings wherever possible
- Actively promote diversity and inclusion in the SIG
- Uphold the Kubernetes Code of Conduct especially in terms of personal behavior and responsibility

### Deviations from [sig-governance]

### Subproject Creation

Federation of Subprojects as defined in [sig-governance]

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-architecture/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[here]: https://docs.google.com/document/d/1TTcfvf8T_tBhGDm-wjgg31WrWjYg8IZEmo3b1mpUXh0/edit?usp=sharing
[conflicts]: https://github.com/kubernetes/community/pull/2074#discussion_r184466503
[process]: https://github.com/kubernetes/sig-release/blob/master/release-team/role-handbooks/enhancements/README.md
[subproject]: https://github.com/kubernetes/sig-release/blob/master/release-team/README.md
