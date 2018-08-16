---
kep-number: 26
title: TTL After Finished
authors:
  - "@janetkuo"
owning-sig: sig-apps
participating-sigs:
  - sig-api-machinery
reviewers:
  - @enisoc
  - @tnozicka
approvers:
  - @kow3ns
editor: TBD
creation-date: 2018-08-16
last-updated: 2018-08-16
status: provisional
see-also:
  - n/a
replaces:
  - n/a
superseded-by:
  - n/a
---

# TTL After Finished Controller

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.
[Tools for generating][] a table of contents from markdown are available.

   * [TTL After Finished Controller](#ttl-after-finished-controller)
      * [Table of Contents](#table-of-contents)
      * [Summary](#summary)
      * [Motivation](#motivation)
         * [Goals](#goals)
      * [Proposal](#proposal)
         * [Concrete Use Cases](#concrete-use-cases)
         * [Detailed Design](#detailed-design)
            * [Feature Gate](#feature-gate)
            * [API Object](#api-object)
               * [Validation](#validation)
         * [User Stories](#user-stories)
         * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
            * [TTL Controller](#ttl-controller)
            * [Finished Jobs](#finished-jobs)
            * [Finished Pods](#finished-pods)
            * [Owner References](#owner-references)
         * [Risks and Mitigations](#risks-and-mitigations)
      * [Graduation Criteria](#graduation-criteria)
      * [Implementation History](#implementation-history)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

We propose a TTL mechanism to limit the lifetime of finished resource objects,
including Jobs and Pods, to make it easy for users to clean up old Jobs/Pods
after they finish. The TTL timer starts when the Job/Pod finishes, and the
finished Job/Pod will be cleaned up after the TTL expires.

## Motivation

In Kubernetes, finishable resources, such as Jobs and Pods, are often
frequently-created and short-lived. If a Job or Pod isn't controlled by a
higher-level resource (e.g. CronJob for Jobs or Job for Pods), or owned by some
other resources, it's difficult for the users to clean them up automatically,
and those Jobs and Pods can accumulate and overload a Kubernetes cluster very
easily. Even if we can avoid the overload issue by implementing a cluster-wide
(global) resource quota, users won't be able to create new resources without
cleaning up old ones first. See [#64470][].

The design of this proposal can be later generalized to other finishable
frequently-created, short-lived resources, such as completed Pods or finished
custom resources.

[#64470]: https://github.com/kubernetes/kubernetes/issues/64470

### Goals

Make it easy to for the users to specify a time-based clean up mechanism for
finished resource objects. 
* It's configurable at resource creation time and after the resource is created.

## Proposal

[K8s Proposal: TTL controller for finished Jobs and Pods][]

[K8s Proposal: TTL controller for finished Jobs and Pods]: https://docs.google.com/document/d/1U6h1DrRJNuQlL2_FYY_FdkQhgtTRn1kEylEOHRoESTc/edit

### Concrete Use Cases

* [Kubeflow][] needs to clean up old finished Jobs (K8s Jobs, TF Jobs, Argo
  workflows, etc.), see [#718][].

* [Prow][] needs to clean up old completed Pods & finished Jobs. Currently implemented with Prow sinker.

* [Apache Spark on Kubernetes][] needs proper cleanup of terminated Spark executor Pods.

* Jenkins Kubernetes plugin creates slave pods that execute builds. It needs a better way to clean up old completed Pods.

[Kubeflow]: https://github.com/kubeflow
[#718]: https://github.com/kubeflow/tf-operator/issues/718
[Prow]: https://github.com/kubernetes/test-infra/tree/master/prow
[Apache Spark on Kubernetes]: http://spark.apache.org/docs/latest/running-on-kubernetes.html

### Detailed Design 

#### Feature Gate

This will be launched as an alpha feature first, with feature gate
`TTLAfterFinished`.

#### API Object

We will add the following API fields to `JobSpec` (`Job`'s `.spec`).

```go
type JobSpec struct {
 	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished
	// execution (either Complete or Failed). If this field is set, once the Job
	// finishes, it will be deleted after ttlSecondsAfterFinished expires. When
	// the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will
	// be honored. If this field is unset, ttlSecondsAfterFinished will not
	// expire. If this field is set to zero, ttlSecondsAfterFinished expires
	// immediately after the Job finishes.
	// This field is alpha-level and is only honored by servers that enable the
	// TTLAfterFinished feature.
	// +optional
	TTLSecondsAfterFinished *int32
}
```

This allows Jobs to be cleaned up after they finish and provides time for
asynchronous clients to observe Jobs' final states before they are deleted.


Similarly, we will add the following API fields to `PodSpec` (`Pod`'s `.spec`).

```go
type PodSpec struct {
 	// ttlSecondsAfterFinished limits the lifetime of a Pod that has finished
	// execution (either Succeeded or Failed). If this field is set, once the Pod
	// finishes, it will be deleted after ttlSecondsAfterFinished expires. When
	// the Pod is being deleted, its lifecycle guarantees (e.g. finalizers) will
	// be honored. If this field is unset, ttlSecondsAfterFinished will not
	// expire. If this field is set to zero, ttlSecondsAfterFinished expires
	// immediately after the Pod finishes.
	// This field is alpha-level and is only honored by servers that enable the
	// TTLAfterFinished feature.
	// +optional
	TTLSecondsAfterFinished *int32
}
```

##### Validation

Because Job controller depends on Pods to exist to work correctly. In Job
validation, `ttlSecondsAfterFinished` of its pod template shouldn't be set, to
prevent users from breaking their Jobs. Users should set TTL seconds on a Job,
instead of Pods owned by a Job.

It is common for higher level resources to call generic PodSpec validation;
therefore, in PodSpec validation, `ttlSecondsAfterFinished` is only allowed to
be set on a PodSpec with a `restartPolicy` that is either `OnFailure` or `Never`
(i.e. not `Always`).

### User Stories

The users keep creating Jobs in a small Kubernetes cluster with 4 nodes.
The Jobs accumulates over time, and 1 year later, the cluster ended up with more
than 100k old Jobs. This caused etcd hiccups, long high latency etcd requests,
and eventually made the cluster unavailable.

The problem could have been avoided easily with TTL controller for Jobs.

The steps are as easy as:

1. When creating Jobs, the user sets Jobs' `.spec.ttlSecondsAfterFinished` to
   3600 (i.e. 1 hour).
1. The user deploys Jobs as usual.
1. After a Job finishes, the result is observed asynchronously within an hour
   and stored elsewhere.
1. The TTL collector cleans up Jobs 1 hour after they complete.

### Implementation Details/Notes/Constraints

#### TTL Controller
We will add a TTL controller for finished Jobs and finished Pods. We considered
adding it in Job controller, but decided not to, for the following reasons:

1. Job controller should focus on managing Pods based on the Job's spec and pod
   template, but not cleaning up Jobs.
1. We also need the TTL controller to clean up finished Pods, and we consider
   generalizing TTL controller later for custom resources. 

The TTL controller utilizes informer framework, watches all Jobs and Pods, and
read Jobs and Pods from a local cache.

#### Finished Jobs

When a Job is created or updated:

1. Check its `.status.conditions` to see if it has finished (`Complete` or
   `Failed`). If it hasn't finished, do nothing. 
1. Otherwise, if the Job has finished, check if Job's 
   `.spec.ttlSecondsAfterFinished` field is set. Do nothing if the TTL field is
   not set. 
1. Otherwise, if the TTL field is set, check if the TTL has expired, i.e. 
   `.spec.ttlSecondsAfterFinished` + the time when the Job finishes
   (`.status.conditions.lastTransitionTime`) > now. 
1. If the TTL hasn't expired, delay re-enqueuing the Job after a computed amount
   of time when it will expire. The computed time period is:
   (`.spec.ttlSecondsAfterFinished` + `.status.conditions.lastTransitionTime` -
   now).
1. If the TTL has expired, `GET` the Job from API server to do final sanity
   checks before deleting it.
1. Check if the freshly got Job's TTL has expired. This field may be updated
   before TTL controller observes the new value in its local cache.
   * If it hasn't expired, it is not safe to delete the Job. Delay re-enqueue
     the Job after a computed amount of time when it will expire.
1. Delete the Job if passing the sanity checks. 

#### Finished Pods

When a Pod is created or updated:
1. Check its `.status.phase` to see if it has finished (`Succeeded` or `Failed`).
   If it hasn't finished, do nothing. 
1. Otherwise, if the Pod has finished, check if Pod's
   `.spec.ttlSecondsAfterFinished` field is set. Do nothing if the TTL field is
   not set. 
1. Otherwise, if the TTL field is set, check if the TTL has expired, i.e.
   `.spec.ttlSecondsAfterFinished` + the time when the Pod finishes (max of all
   of its containers termination time
   `.containerStatuses.state.terminated.finishedAt`) > now. 
1. If the TTL hasn't expired, delay re-enqueuing the Pod after a computed amount
   of time when it will expire. The computed time period is:
   (`.spec.ttlSecondsAfterFinished` + the time when the Pod finishes - now).
1. If the TTL has expired, `GET` the Pod from API server to do final sanity
   checks before deleting it.
1. Check if the freshly got Pod's TTL has expired. This field may be updated
   before TTL controller observes the new value in its local cache.
   * If it hasn't expired, it is not safe to delete the Pod. Delay re-enqueue
     the Pod after a computed amount of time when it will expire.
1. Delete the Pod if passing the sanity checks. 

#### Owner References

We have considered making TTL controller leave a Job/Pod around even after its
TTL expires, if the Job/Pod has any owner specified in its
`.metadata.ownerReferences`.

We decided not to block deletion on owners, because the purpose of
`.metadata.ownerReferences` is for cascading deletion, but not for keeping an
owner's dependents alive. If the Job is owned by a CronJob, the Job can be
cleaned up based on CronJob's history limit (i.e. the number of dependent Jobs
to keep), or CronJob can choose not to set history limit but set the TTL of its
Job template to clean up Jobs after TTL expires instead of based on the history
limit capacity. 

Therefore, a Job/Pod can be deleted after its TTL expires, even if it still has
owners. 

Similarly, the TTL won't block deletion from generic garbage collector. This
means that when a Job's or Pod's owners are gone, generic garbage collector will
delete it, even if it hasn't finished or its TTL hasn't expired. 

### Risks and Mitigations

Risks:
* Time skew may cause TTL controller to clean up resource objects at the wrong
  time.

Mitigations:
* In Kubernetes, it's required to run NTP on all nodes ([#6159][]) to avoid time
  skew. We will also document this risk.

[#6159]: https://github.com/kubernetes/kubernetes/issues/6159#issuecomment-93844058

## Graduation Criteria

We want to implement this feature for Pods/Jobs first to gather feedback, and
decide whether to generalize it to custom resources. This feature can be
promoted to beta after we finalize the decision for whether to generalize it or
not, and when it satisfies users' need for cleaning up finished resource
objects, without regressions.

This will be promoted to GA once it's gone a sufficient amount of time as beta
with no changes. 

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

TBD
