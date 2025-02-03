# 2024 Annual Report: WG Batch

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

See [2024 Highlights](#2024-highlights).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  Yes, JobSet has 1 active owner at the moment.

### 2024 Highlights

We will breakdown our highlights into Sub Projects, KEPs, talks, community adoption.

#### Sub Projects

##### Kueue

Kueue has had 5 minor releases in 2024.

- [Release 0.6](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.6.0)

- [Release 0.7](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.7.0)

- [Release 0.8](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.8.0)

- [Release 0.9](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.9.0)

- [Release 0.10](https://github.com/kubernetes-sigs/kueue/releases/tag/v0.10.0)

In 2024, the kueue community would like to highlight Topology aware scheduling, MultiKueue, Kueue Dashboard, KueueCtrl, Deployment/Statefulset integration for serving and Fair sharing.

[Topology aware scheduling](https://kueue.sigs.k8s.io/docs/concepts/topology_aware_scheduling/) facilitates scheduling of workloads that take into account data center topology.
Workloads benefit from using interconnects that are physically close together.

[MultiKueue](https://kueue.sigs.k8s.io/docs/concepts/multikueue/) provides a way of dispatching batch workloads to worker clusters.
Kueue provides multicluster dispatching for popular batch workloads such as Ray, Job, Kubeflow and JobSet.
This feature went beta in 0.9.

[Kueue Dashboards](https://github.com/kubernetes-sigs/kueue/tree/release-0.10/cmd/experimental/kueue-viz) has been a popular ask for Kueue.
Users would like to have a visualization representation of queueing and we are happy to announce that a dashboard has been created for Kueue.
This went into kueue in late 2024 and a big focus of 2025 will be to harden this for production.

[KueueCtl](https://kueue.sigs.k8s.io/docs/reference/kubectl-kueue/) provides a cli for creating kueue objects.
The plugin is hosted in krew and is easily installed as a kueue plugin.

[Deployment](https://kueue.sigs.k8s.io/docs/tasks/run/deployment/) and [StatefulSet](https://kueue.sigs.k8s.io/docs/tasks/run/statefulset/) integration provides an avenue for the usage of Kueue for serving workloads. Serving leads to a need for sharing/preemption of model servers that may leverage accelerators. Kueue provides an integration with popular methods of deploying services (Deployment/StatefulSet).

##### JobSet

Jobset has had 4 minor releases in 2024.

- [Release 0.4](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.4.0)

- [Release 0.5](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.5.0)

- [Release 0.6](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.6.0)

- [Release 0.7](https://github.com/kubernetes-sigs/jobset/releases/tag/v0.7.0)

A major achievement of JobSet has been the adoption of JobSet as a component for [Kubeflow Training Operator](https://github.com/kubeflow/training-operator) V2.
There has been a collaborative effort with the Kubeflow community and the batch community to implement the features needed for this integration.

[Metaflow](https://github.com/Netflix/metaflow/pull/1804) has adopted the use of JobSet for distributed ML training.

##### KJob

[KJob](https://github.com/kubernetes-sigs/kjob) has been started to provide a CLI friendly way for users to submit batch jobs.
The HPC/ML community tend to prefer CLI over YAML so the focus was to provide a templated solution for submitting batch jobs.
Another focus of this project is to provide a smooth transition for Slurm users.

#### KEPs

WG-Batch provided a series of kubernetes enhancements that improved the experience of batch workloads on Kubernetes. In 2024, this group proposed/implemented the following KEPs.

- [Job Managed By](https://github.com/kubernetes/enhancements/issues/4368)
  - Promoted to beta.

- [Job Success Policy](https://github.com/kubernetes/enhancements/issues/3998)
  - Promoted to beta.

- [Elastic Index Jobs](https://github.com/kubernetes/enhancements/issues/3715)
  - Promoted to stable.

- [Pod Failure Policy](https://github.com/kubernetes/enhancements/issues/3329)
  - Promoted to stable.

- [Pod Index Label](https://github.com/kubernetes/enhancements/issues/4017)
  - Promoted to stable.

### Talks

- WG-Batch Update at Kubecon
  - Speakers: Kevin Hannon and Marcin Wielgus
  - Kubecon NA, Salt Lake City
  - [Recording](https://www.youtube.com/watch?v=C2ABOEzZTWg&list=PLj6h78yzYM2Pw4mRw4S-1p_xLARMqPkA7&index=283&pp=iAQB)

- Keynote: MultiCluster Batch Jobs Dispatching with Kueue at CERN
  - Speakers: Ricardo Rocha and Marcin Wielgus
  - Kubecon NA, Salt Lake City
  - [Recording](https://www.youtube.com/watch?v=xMmskWIlktA&list=PLj6h78yzYM2Pw4mRw4S-1p_xLARMqPkA7&index=193&pp=iAQB)

- Multitenancy and Fairness at Scale with Kueue: A Case Study
  - Speakers: Aldo Culquicondor and Rajat Phull
  - Kubecon NA, Salt Lake City
  - [Recording](https://www.youtube.com/watch?v=GYiuTQCvTx8&list=PLj6h78yzYM2Mvqk_mNejD7kbe3tldxxsr&index=5&pp=iAQB)

- Advanced Resource Management for Running AI/ML Workloads with Kueue
  - Speakers: Michał Woźniak and Yuki Iwai
  - Kubecon EU, Paris
  - [Recording](https://www.youtube.com/watch?v=6k_8Go3u8Qk)

- Scale Your Batch / Big Data / AI Workloads Beyond the Kubernetes Scheduler
  - Speaker: Antonin Stefanutti and Anish Asthana
  - KubeCon EU, Paris
  - [Recording](https://www.youtube.com/watch?v=Ij5EAnuF-jk&list=PLj6h78yzYM2PWGv34W6w5ssq1b1meRmY7&index=15&pp=iAQB)

- WG-Batch Update
  - Speaker: Michał Woźniak and Yuki Iwai
  - KubeCon EU, Paris
  - [Recording](https://www.youtube.com/watch?v=2D2QSzUnS0M&list=PLj6h78yzYM2N8nw1YcqqKveySH6_0VnI0&index=84&pp=iAQB)

- How the Kubernetes Community is Improving Kubernetes for HPC/AI/ML Workloads
  - Author: Kevin Hannon
  - FOSDEM 2024
  - [Recording](https://live.fosdem.org/watch/ua2118)

### Community adoption

- [Kubeflow Training Operator v2](https://github.com/kubeflow/training-operator/blob/0c30f5cd306611f061b6dd529d3c7b7981a7d27c/docs/proposals/2170-kubeflow-training-v2/README.md#kep-2170-kubeflow-training-v2-api) will be using JobSet as a critical component for training and finetuning.

- [Metaflow supports JobSet](https://github.com/Netflix/metaflow/pull/1804) for distributed training.

- Airflow has built an [integration](https://airflow.apache.org/docs/apache-airflow-providers-cncf-kubernetes/stable/_api/airflow/providers/cncf/kubernetes/operators/kueue/index.html) with Kueue.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2024
      - [WG-Batch Updates at Kubecon EU 2024](https://www.youtube.com/watch?v=2D2QSzUnS0M&list=PLj6h78yzYM2N8nw1YcqqKveySH6_0VnI0&index=84&pp=iAQB)
      - [WG-Batch Updates at Kubecon NA 2024](https://www.youtube.com/watch?v=C2ABOEzZTWg&list=PLj6h78yzYM2Pw4mRw4S-1p_xLARMqPkA7&index=283&pp=iAQB)
[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-batch/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
