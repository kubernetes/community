---
kep-number: 24
title: Coscheduling
authors:
  - "@k82cn"
owning-sig: sig-scheduling, machine-learning WG
reviewers:
  - "@bsalamat"
  - "@vishh"
approvers:
  - "@bsalamat"
  - "@vishh"
editor: TBD
creation-date: 2018-07-03
last-updated: 2018-10-12
status: provisional
---

# Coscheduling

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Motivation](#motivation)
* [Function Detail](#function%20detail)
  * [API Definitation](#api%20definition)
  * [Lifecycle Management](#lifecycle%20management)
  * [Scheduling](#scheduling)
  * [Customized Controller](#customized%20controller)
* [Feature Interaction](#feature%20interaction)
  * [Multi-scheduler](#multi-scheduler)
  * [Priority/Preemption](#priority/preemption)
  * [Pod RestartPolicy](#pod%20restartPolicy)
  * [Admission Controller](#admission%20controller)
  * [Kubectl](#kubectl)
* [References](#references)

## Motivation

Kubernetes has become a popular solution for orchestrating containerized workloads; it has been largely successful in orchestrating serving and storage workloads, and, with native K8s support for Spark. Meanwhile, the community also try to run Machine Learning (ML) workloads on Kubernetes, e.g. [kubeflow/tf-operator](https://github.com/kubeflow/tf-operator). When running a Tensorflow/MPI job, all tasks of a job must be start together; otherwise, did not start anyone of tasks. If the resource is enough to run all 'tasks', everything is fine; but it's not true for most of case, especially in the on-prem environment. In worst case, all jobs are pending here because of deadlock: every job only start part of tasks, and waits for the other tasks to start. It'll be worse in federation for cross-domain case which is out of scope in this doc.

After the discussion at [Coscheduling/Gang-scheduling](https://docs.google.com/document/d/1AUwcvTtULNvow5M9e428FnlvINO1uQ7ojRoTGuTp4DA/edit#heading=h.ckn8nv2jj0xv) proposal, we decide to implement Coscheduling in [kube-batch](https://github.com/kubernetes-sigs/kube-batch) by CRDs. kube-batch focuses on "batch" workload in kubernetes, and will share the same [scheduling frameworks](https://github.com/kubernetes/community/pull/2281) when it's ready. This document is used to provide definition of API object and the scheduler behaviour of Coscheduling.

## Function Detail

### API Definition

The following requirements are identified during the discussion of this feature:

1. Existing workload can use this feature without (or with a few) configuration changes
2. Pods of a group/gang may have different `PodSpec` (and/or belong to different collections)
3. Existing controllers which are responsible for managing life cycle of collections work well with this feature

To meet the requirements above, the following **Kind** is introduced by CRD under `incubator.scheduling.k8s.io/v1alpha1` **Group**/**Version**.

```go
// PodGroup defines the scheduling requirement of a pod group
type PodGroup struct {
    metav1.TypeMeta
    metav1.ObjectMeta

	// Spec defines the behavior of a pod group.
	// +optional
    Spec PodGroupSpec

	// Status represents the current information about a pod group.
	// This data may not be up to date.
	// +optional
    Status PodGroupStatus
}

// PodGroupSpec represents the template of a pod group.
type PodGroupSpec struct {
    // MinMembers defines the minimal number of members/tasks to run the pod group;
    // if there's not enough resources to start all tasks, the scheduler
    // will not start anyone.
    MinMembers int

    // TotalResources defines the total resource the PodGroup requests to run
    // Pods.
    TotalResources v1.ResourceList
}

// PodGroupStatus represents the current state of a pod group.
type PodGroupStatus struct {
    // The number of admitted pods.
    // +optional
    Admitted int32

    // The number of actively running pods.
    // +optional
    Running int32

    // The number of pods which reached phase Succeeded.
    // +optional
    Succeeded int32

    // The number of pods which reached phase Failed.
    // +optional
    Failed int32
}
```

The `PodGroup`, which is a namespaced object, specifies the attributes and status of a pod group, e.g. number of pods in a group. To define which pods are member of `PodGroup`, the following annotation key is introduced for `Pod`; the annotation key is used for this alpha feature, and it'll be changed to a more permanent form, such a field, when moving `PodGroup` to core.

```go
scheduling.k8s.io/group-name
```

The `scheduling.k8s.io/group-name` annotation specifies the `PodGroup` that it belongs to; and the pod can only belong to the `PodGroup` in the same namespace. The pod, controlled by different collections, can also belong to the same `PodGroup`.  Because of performance concern,  it does not use `LabelSelector` to build the relationship between `PodGroup` and `Pod`.

### Lifecycle Management

As the lifecycle of Pods in PodGroup may be different from controller to another, the lifecycle of the members is not managed by the coscheduling feature. Each collection controller may implement or already have the mean to manage lifecycle of its members. The scheduler'll record related events for controller to manage pods, e.g. `Unschedulable`. A controller of `PodGroup` will be introduced later for lifecycle management, e.g. restart the whole `PodGroup`, according to the configuration, when the number of running pods drop below `spec.MinMembers` at run-time.

The update to `PodGroup` is not supported for now; and deleting `PodGroup` does not impact Pod's status.

### Scheduling

The scheduler only watches `PodGroup` and `Pod`. It'll reconstruct 'Job' by annotation of Pod and `PodGroup`, the `Pod`s are considered as 'Task' of 'Job'; if annotation is empty, the scheduler records an unschedulable event of pod to ask user/controller to resubmit it. The schduler does not schedule pods until its `PodGroup` is created.

As batch scheduler and default scheduler may be running in parallel; the batch scheduler follows multi-scheduler feature to only handle the `Pod` that submitted to it. The batch scheduler does scheduling as follow:

1. Reconstructing 'Job' by the annotation of `Pod` and `PodGroup`
2. If there are less Pods than `minMembers` of `PodGroup`, the 'job' will not be scheduled; and an unschedulable event of `Pod` will be recorded
3. In `allocate` phase, scheduler will
   * record an `Unschedulable` event of `PodGroup` if some pods are running but `succeeded + pending + running < minMembers`, the controller takes action according to its configuration
   * allocate (but not bind) resource to Pods according to Pod's spec, e.g. `NodeAffinity`
   * bind all Pods to hosts until job is ready: if `minMembers` <= `allocated Pods` + `pending Pods`, it's ready when `minMembers` <= `allocated Pods`; otherwise, `numMember` <= `allocated Pods` + `succeeded Pods`
4. If can not allocate enough resources to the job, the pods stay pending; and the resource cannot be allocated to other job

That may make resources (less than job's resource request) idle for a while, e.g. a huge job. The solution, e.g. backfill other smaller jobs to improve the resource utilization, will be proposed in coming release. In `allocate` phase, only pod's `NodeAffinity` takes effect; the other predicates/priorities will be included on-demand in coming release.

### Customized Controller

A typical example of customized controller is [kubeflow/tf-operator](https://github.com/kubeflow/tf-operator), which managed the Pods for TensorFlow on Kubernetes, required `gang-scheduling` in upstream. Here's an example of customized controller that demonstrated the usage of  `gang-scheduling` in `kube-batch`.

Usually, CRD ([CustomResourceDefinitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)) feature is used to introduce a customized **Kind**, named `CRDJob` as example. The customized controller, named `CRDJobController`, watches it and manages lifecycle of it:

1. For each `CRDJob`, `CRDJobController` creates a `CRDJob` and `PodGroup` (one `CRDJob` with one `PodGroup` as example). The attributes of `PodGroup`should be set accordingly, e.g `numMember`; it's up to customized controller on how to manage relationship between `PodGroup` and `CRDJob`, e.g. `metadata.name`.
2. When `CRDJobController` create Pods, its annotation should be set accordingly. `kube-batch` follows gang-scheduling logic to schedule those pods in batch.
3. When pods failed/deleted/unschedulable, it is up to `CRDJobController` on how to manage `CRDJob`'s lifecycle. For example, if `CRDJobController` manages lifecycle itself, set `.spec.Policy` of `PodGroup` to nil; otherwise, `PodGroupController` will manage the lifecycle as described above.
4. If `CRDJob` was deleted, the `PodGroup` must be deleted accordingly.

## Feature Interaction

### Multi-scheduler

Since multiple schedulers work in parallel, there may be decision conflict between different schedulers; and the kubelet will reject one pod (failed) if conflict. The controller will handle rejected pods based on its lifecycle policy for failed pods. Users and cluster admins may reduce the probability of such conflicts by partitioning the clusters logically, for example, by placing node-affinity to distinct set of nodes on various groups of pods. 

### Priority/Preemption

A rejected or preempted batch/run-to-completion pod may trigger a restart of the whole `PodGroup`. This can have negative impact on performance.  The solution on how to handle conflicts better will be proposed in coming release.

The default scheduler should also consider `PodGroup` when preempting pods, similar to `PodDisruptionBudgets`.

### Pod RestartPolicy

Pod's `RestartPolicy` still works as before. But for batch/run-to-compelete workload, it's better to set `RestartPolicy` to `Never` to avoid endless restart loop.

### Admission Controller

If quota runs out in the middle of creating a group of pods, a few members of a `PodGroup` may be created, while the rest will be denied by the `ResourceQuota` admission controller. `.spec.TotalResource` is added in `PodGroup` to address this problem. When a `PodGroup` is created with `.spec.TotalResource`, so much quota is reserved for the group if there is available quota. Pods of group use the already reserved quota. By setting `.spec.TotalResource` properly, one can ensure that Pods of a `PodGroup` have enough quota at creation time. The design on `Quota` enhancement to support `.spec.TotalResource` will be proposed later for review.

### Kubectl

kubectl is enhanced to support `PodGroup` by [kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/), including its status.

## Roadmap

1. `PodGroup` CRD and Coscheduling (by kube-batch) (1.13)
1. `PodGroup` controller (by kube-batch) (1.13)
1. Admission controller for `.spec.TotalResource` (1.13, 1.14)

## References

* [Coscheduling in Kubernetes](https://docs.google.com/document/d/1AUwcvTtULNvow5M9e428FnlvINO1uQ7ojRoTGuTp4DA/edit#heading=h.ckn8nv2jj0xv)
* [Indexed Job](https://github.com/kubernetes/kubernetes/issues/14188)
* [Schedule a group of pods all at once](https://github.com/kubernetes/kubernetes/issues/16845)
* [kubeflow/tf-operator: Prevent scheduling deadlocks](https://github.com/kubeflow/tf-operator/issues/165)
