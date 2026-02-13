# 2025 Annual Report: WG Batch

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

See [2025 Highlights](#2025-highlights).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  No, all subprojects have sufficient active owners.

### 2025 Highlights

We will break down our highlights into Sub Projects, KEPs, talks, community adoption.

#### Sub Projects

##### Kueue

Kueue has had 5 minor releases in 2025.

- [Release 0.11](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.11.0)

- [Release 0.12](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.12.0)

- [Release 0.13](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.13.0)

- [Release 0.14](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.14.0)

- [Release 0.15](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.15.0)

In 2025, the Kueue community would like to highlight Topology Aware Scheduling, MultiKueue, Admission Fair Sharing, Elastic Jobs, DRA Integration, v1beta2 API and KueueViz Dashboard.

[Topology Aware Scheduling](https://kueue.sigs.k8s.io/docs/concepts/topology_aware_scheduling/) matured from alpha to beta (enabled by default in 0.14), facilitating scheduling of workloads that take into account data center topology.
Workloads benefit from using interconnects that are physically close together.

[MultiKueue](https://kueue.sigs.k8s.io/docs/concepts/multikueue/) expanded to support RayCluster, RayJob, and Pods.
An external dispatcher API was introduced in 0.13 for nominating worker clusters.
Security hardening with kubeconfig validation was added in 0.15.

[Admission Fair Sharing](https://kueue.sigs.k8s.io/docs/concepts/admission_fair_sharing/) progressed from alpha (0.12) to beta (0.15).
This feature orders workloads based on recent LocalQueue usage rather than just priority, preventing queue manipulation and ensuring fair resource distribution.

[Elastic Jobs](https://kueue.sigs.k8s.io/docs/concepts/elastic_workload/) via WorkloadSlices was introduced in 0.13 as an alpha feature.
This enables dynamic job resizing without suspension or requeueing.

[DRA Integration](https://github.com/kubernetes-sigs/kueue/tree/main/keps/2941-DRA) was introduced in 0.14 as an alpha feature, providing Dynamic Resource Allocation support for specialized hardware.

[v1beta2 API](https://kueue.sigs.k8s.io/docs/reference/kueue.v1beta2/) was introduced in 0.15, representing API maturation toward stability.

[KueueViz Dashboard](https://kueue.sigs.k8s.io/docs/tasks/manage/enable_kueueviz/) was hardened for production with Helm charts for installation and rebranded with the CNCF logo.

##### JobSet

JobSet has had 3 minor releases in 2025.

- [Release 0.8](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.8.0)

- [Release 0.9](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.9.0)

- [Release 0.10](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.10.0)

In 2025, the JobSet community would like to highlight the Kubernetes blog post, Failure Policy, DependsOn API, and Coordinator.

The official Kubernetes blog post [Introducing JobSet](https://kubernetes.io/blog/2025/03/23/introducing-jobset/) was published in March 2025.

[Failure Policy](https://jobset.sigs.k8s.io/docs/tasks/failure_policy/) added the `onJobFailureMessagePatterns` field to [distinguish retriable from non-retriable Pod failure policies](https://github.com/kubernetes-sigs/jobset/pull/1027).
Previously, all failures matching a Pod Failure Policy resulted in the same `PodFailurePolicy` job failure reason, making it impossible to apply different JobSet actions.
With `onJobFailureMessagePatterns`, rules can match against failure message strings to differentiate between cases such as a non-retriable exit code (fail immediately) and a node eviction (restart without counting toward the limit).

[DependsOn API](https://github.com/kubernetes-sigs/jobset/tree/main/keps/672-serial-job-execution) defines execution dependencies between ReplicatedJobs within a JobSet, enabling serial/sequential execution.
A ReplicatedJob can declare that it depends on another ReplicatedJob reaching either `Ready` or `Complete` status before it starts.
This supports multi-stage workflows like data initialization followed by distributed training.

[Coordinator](https://jobset.sigs.k8s.io/docs/concepts/) designates a specific Pod within a JobSet as the coordinator for distributed ML/HPC workloads.
A stable network endpoint for the coordinator is added as a label and annotation to every Job and Pod in the JobSet, enabling other pods to discover the coordinator without hardcoding addresses.

##### KJob

[KJob](https://github.com/kubernetes-sigs/kjob) had its first release (v0.1.0) in 2025, providing the base functionality for CLI-friendly batch job submission.
KJob provides a template-based job execution with built-in SLURM support and kubectl plugin integration.
The HPC/ML community tend to prefer CLI over YAML so the focus was to provide a templated solution for submitting batch jobs and a smooth transition for Slurm users.

#### KEPs

WG-Batch provided a series of Kubernetes enhancements that improved the experience of batch workloads on Kubernetes. In 2025, this group proposed/implemented/consulted the following KEPs.

- [Job Success Policy](https://github.com/kubernetes/enhancements/issues/3998)
  - Promoted to stable.

- [Backoff Limit Per Index](https://github.com/kubernetes/enhancements/issues/3850)
  - Promoted to stable.

- [Pod Replacement Policy](https://github.com/kubernetes/enhancements/issues/3939)
  - Promoted to stable.

- [Job Managed By](https://github.com/kubernetes/enhancements/issues/4368)
  - Promoted to stable.

- [Gang Scheduling / Workload API](https://github.com/kubernetes/enhancements/issues/4671)
  - Introduced as alpha.

- [Mutable Container Resources for Suspended Jobs](https://github.com/kubernetes/enhancements/issues/5440)
  - Introduced as alpha.

### Talks

- Accelerate Your AI/ML Workloads With Topology-Aware Scheduling in Kueue
  - Speakers: Michal Wozniak and Yuki Iwai
  - KubeCon EU, London
  - [Recording](https://www.youtube.com/watch?v=F55pFM1M1bU)

- Tutorial: Build, Operate, and Use a Multi-Tenant AI Cluster Based Entirely on Open Source
  - Speakers: Claudia Misale, Olivier Tardieu, and David Grove
  - KubeCon EU, London
  - [Recording](https://www.youtube.com/watch?v=Ab7mRoJYsMo)

- Kueue: Save Some QPS for the Rest of Us! How To Manage 100k Updates Per Second
  - Speaker: Patryk Bundyra
  - KubeCon EU, London
  - [Recording](https://www.youtube.com/watch?v=njNXlZNT3dw)

- WG-Batch Updates: What's New and What Is Next?
  - Speaker: Marcin Wielgus
  - KubeCon EU, London
  - [Recording](https://www.youtube.com/watch?v=aWxuaEFSarU)

- From High Performance Computing to AI Workloads on Kubernetes: MPI Runtime in Kubeflow TrainJob
  - Speakers: Andrey Velichkevich and Yuki Iwai
  - KubeCon EU, London
  - [Recording](https://www.youtube.com/watch?v=Fnb1a5Kaxgo)

- Resource Fairness and Utilization for Heterogeneous Batch/ML Platforms With Kueue
  - Speakers: Yuki Iwai and Gabe Saba
  - KubeCon NA, Atlanta
  - [Recording](https://www.youtube.com/watch?v=dKhF-hZi7CI)

- WG-Batch Updates: What's New and What Is Next?
  - Speakers: Michal Wozniak and Yuki Iwai
  - KubeCon Japan, Tokyo
  - [Recording](https://www.youtube.com/watch?v=jeRhDmp_i2M)

### Community adoption

- [CNCF Kubernetes AI Conformance Program](https://www.cncf.io/announcements/2025/11/11/cncf-launches-certified-kubernetes-ai-conformance-program-to-standardize-ai-workloads-on-kubernetes/) was launched in November 2025 to standardize AI workloads on Kubernetes, with JobSet and Kueue as a key component in the ecosystem.

- [Kueue integrated with Kubeflow TrainJob](https://kueue.sigs.k8s.io/docs/tasks/run/trainjobs/), enabling quota management and admission control for distributed training workloads submitted via Kubeflow Trainer v2.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025
      - [WG-Batch Updates at KubeCon EU 2025](https://sched.co/1tczF)
      - [WG-Batch Updates at KubeCon Japan 2025](https://sched.co/1x6ze)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-batch/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
