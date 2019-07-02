# Job Controller

## Abstract

A proposal for implementing a new controller - Job controller - which will be responsible
for managing pod(s) that require running once to completion even if the machine
the pod is running on fails, in contrast to what ReplicationController currently offers.

Several existing issues and PRs were already created regarding that particular subject:
* Job Controller [#1624](https://github.com/kubernetes/kubernetes/issues/1624)
* New Job resource [#7380](https://github.com/kubernetes/kubernetes/pull/7380)


## Use Cases

1. Be able to start one or several pods tracked as a single entity.
1. Be able to run batch-oriented workloads on Kubernetes.
1. Be able to get the job status.
1. Be able to specify the number of instances performing a job at any one time.
1. Be able to specify the number of successfully finished instances required to finish a job.
1. Be able to specify a backoff policy, when job is continuously failing.


## Motivation

Jobs are needed for executing multi-pod computation to completion; a good example
here would be the ability to implement any type of batch oriented tasks.


## Backoff policy and failed pod limit

By design, Jobs do not have any notion of failure, other than a pod's `restartPolicy`
which is mistakenly taken as Job's restart policy ([#30243](https://github.com/kubernetes/kubernetes/issues/30243),
[#[43964](https://github.com/kubernetes/kubernetes/issues/43964)]).  There are
situation where one wants to fail a Job after some amount of retries over a certain
period of time, due to a logical error in configuration etc.  To do so we are going
to introduce the following fields, which will control the backoff policy: a number of
retries and an initial time of retry.  The two fields will allow fine-grained control
over the backoff policy.  Each of the two fields will use a default value if none
is provided,  `BackoffLimit` is set by default to 6 and `BackoffSeconds` to 10s.
This will result in the following retry sequence: 10s, 20s, 40s, 1m20s, 2m40s,
5m20s.  After which the job will be considered failed.

Additionally, to help debug the issue with a Job, and limit the impact of having
too many failed pods left around (as mentioned in [#30243](https://github.com/kubernetes/kubernetes/issues/30243)),
we are going to introduce a field which will allow specifying the maximum number
of failed pods to keep around.  This number will also take effect if none of the
limits described above are set. By default it will take value of 1, to allow debugging
job issues, but not to flood the cluster with too many failed jobs and their
accompanying pods.

All of the above fields will be optional and will apply when `restartPolicy` is
set to `Never` on a `PodTemplate`.  With restart policy `OnFailure` only `BackoffLimit`
applies. The reason for that is that failed pods are already restarted by the
kubelet with an [exponential backoff](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy).
Additionally, failures are counted differently depending on `restartPolicy`
setting.  For `Never` we count actual pod failures (reflected in `.status.failed`
field). With `OnFailure`, we take an approximate value of pod restarts (as reported
in `.status.containerStatuses[*].restartCount`).
When `.spec.parallelism` is set to a value higher than 1, the failures are an
overall number (as coming from `.status.failed`) because the controller does not
hold information about failures coming from separate pods.


## Implementation

Job controller is similar to replication controller in that they manage pods.
This implies they will follow the same controller framework that replication
controllers already defined.  The biggest difference between a `Job` and a
`ReplicationController` object is the purpose; `ReplicationController`
ensures that a specified number of Pods are running at any one time, whereas
`Job` is responsible for keeping the desired number of Pods to a completion of
a task.  This difference will be represented by the `RestartPolicy` which is
required to always take value of `RestartPolicyNever` or `RestartOnFailure`.


The new `Job` object will have the following content:

```go
// Job represents the configuration of a single job.
type Job struct {
    TypeMeta
    ObjectMeta

    // Spec is a structure defining the expected behavior of a job.
    Spec JobSpec

    // Status is a structure describing current status of a job.
    Status JobStatus
}

// JobList is a collection of jobs.
type JobList struct {
    TypeMeta
    ListMeta

    Items []Job
}
```

`JobSpec` structure is defined to contain all the information how the actual job execution
will look like.

```go
// JobSpec describes how the job execution will look like.
type JobSpec struct {

    // Parallelism specifies the maximum desired number of pods the job should
    // run at any given time. The actual number of pods running in steady state will
    // be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism),
    // i.e. when the work left to do is less than max parallelism.
    Parallelism *int32

    // Completions specifies the desired number of successfully finished pods the
    // job should be run with. Defaults to 1.
    Completions *int32

    // Optional duration in seconds relative to the startTime that the job may be active
    // before the system tries to terminate it; value must be a positive integer.
    // It applies to overall job run time, no matter of the value of completions
    // or parallelism parameters.
    ActiveDeadlineSeconds *int64

    // Optional number of retries before marking this job failed.
    // Defaults to 6.
    BackoffLimit *int32

    // Optional time (in seconds) specifying how long the initial backoff will last.
    // Defaults to 10s.
    BackoffSeconds *int64

    // Optional number of failed pods to retain.
    FailedPodsLimit *int32

    // Selector is a label query over pods running a job.
    Selector LabelSelector

    // Template is the object that describes the pod that will be created when
    // executing a job.
    Template *PodTemplateSpec
}
```

`JobStatus` structure is defined to contain information about pods executing
specified job.  The structure holds information about pods currently executing
the job.

```go
// JobStatus represents the current state of a Job.
type JobStatus struct {
    Conditions []JobCondition

    // CreationTime represents time when the job was created
    CreationTime unversioned.Time

    // StartTime represents time when the job was started
    StartTime unversioned.Time

    // CompletionTime represents time when the job was completed
    CompletionTime unversioned.Time

    // Active is the number of actively running pods.
    Active int32

    // Succeeded is the number of pods successfully completed their job.
    Succeeded int32

    // Failed is the number of pods failures, this applies only to jobs
    // created with RestartPolicyNever, otherwise this value will always be 0.
    Failed int32
}

type JobConditionType string

// These are valid conditions of a job.
const (
    // JobComplete means the job has completed its execution.
    JobComplete JobConditionType = "Complete"
)

// JobCondition describes current state of a job.
type JobCondition struct {
    Type               JobConditionType
    Status             ConditionStatus
    LastHeartbeatTime  unversioned.Time
    LastTransitionTime unversioned.Time
    Reason             string
    Message            string
}
```

## Events

Job controller will be emitting the following events:
* JobStart
* JobFinish

## Future evolution

Below are the possible future extensions to the Job controller:
* Be able to limit the execution time for a job, similarly to ActiveDeadlineSeconds for Pods. *now implemented*
* Be able to create a chain of jobs dependent one on another. *will be implemented in a separate type called Workflow*
* Be able to specify the work each of the workers should execute (see type 1 from
  [this comment](https://github.com/kubernetes/kubernetes/issues/1624#issuecomment-97622142))
* Be able to inspect Pods running a Job, especially after a Job has finished, e.g.
  by providing pointers to Pods in the JobStatus ([see comment](https://github.com/kubernetes/kubernetes/pull/11746/files#r37142628)).
* help users avoid non-unique label selectors ([see this proposal](selector-generation.md))

