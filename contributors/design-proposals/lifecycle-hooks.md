# Lifecycle Hooks

**Author**: @tnozicka

**Status**: Proposal

**RFC Issue**: https://github.com/kubernetes/kubernetes/issues/14512

## Abstract
The intent of this proposal it to support lifecycle hooks in all objects having rolling/recreate update; currently those are Deployments, StatefulSets and DaemonSets. Lifecycle hooks should allow users to better customize updates (rolling/recreate) without having to write their own controllers. 

## Lifecycle Hooks
Lifecycle hook is a Job that will get triggered by the update reaching certain progress point (like 50%). The progress point is determined by the ratio of new.availableReplicas to the declared replicas or by explicitly stating the number of new.availableReplicas.

### Previous Implementations
Lifecycle hooks are already implemented in OpenShift but we want to enhance them and implement them natively in Kubernetes.
 - [OpenShift proposal](https://github.com/openshift/origin/blob/master/docs/proposals/post-deployment-hooks.md)
 - [OpenShift docs](https://docs.openshift.org/latest/dev_guide/deployments/deployment_strategies.html#lifecycle-hooks)

### Previous Kubernetes Proposals
 - https://github.com/kubernetes/kubernetes/pull/33545

### Use Cases
The most common use case for lifecycle hooks is to have pre, mid and/or post hook. It is mostly used to run some kind of acceptance check in the middle and/or at the end to fully verify the update is working as expected and rollback if it isn't. The acceptance check may be time consuming and more thorough than what readiness and liveness probes are intended for. You can also notify external services from them, migrate database in the middle of an update, send messages to IRC channel or do anything else.

[A short demo](https://youtu.be/GVNTm_K43iI) simulating lifecycle hooks using auto-pausing that has been presented at SIG-Apps meeting on August 21, 2017.

#### Reusability
If you are a big shop and you are running several instances of e.g. your database you want to reuse definition of lifecycle hooks e.g. for several instances of your database. This is reflected in the design bellow by having separate object to define lifecycle hooks and reference it from the objects.

If you would be worried about having shared definitions so e.g. your mistakes won't spread too much you can always choose not to share those definitions and reference unique lifecycle hook objects from every instance. 

## API Objects
### New Objects
```go
type LifecycleTemplate struct {
    TypeMeta
    ObjectMeta
    Spec LifecycleTemplateSpec
}

type LifecycleTemplateSpec struct {
    // "Always" - keep all Jobs 
    // "OnFailure" - keep only failed Jobs
    // "Never" - delete Jobs immediately 
    RetainPolicy string
    
    // After reaching this limit the Job history will be pruned.
    RevisionHistoryLimit *int32
    
    // List of lifecycle hooks
    // Can have multiple hooks at the same ProgressPoint; order independent
    Hooks []LifecycleHook
}

type LifecycleHook struct {
    // Unique name
    Name string
    
    // ProgressPoint specifies the point during update when the hook should be triggered.
    // Accepts both the number of new pods available or a percentage value representing the ratio
    // of new.availableReplicas to the declared replicas. In case of getting a percentage
    // it will try to reach the exact or the closest possible point right after it,
    // trigger the job, wait for it to complete and then continue. 
    // If such situation shall occur that two different ProgressPoints should be reached at
    // the same time, all the hooks for an earlier ProgressPoint will be ran (and finished)
    // before any later one.
    ProgressPoint intstr.IntOrString
    
    // "Abort" - Failure for this hook is fatal, deployment is considered failed and should be rollbacked.
    // "Ignore" - Even if the hook fails deployment continues rolling out new version.
    FailurePolicy string
    
    JobSpec v1.JobSpec
}

type LifecycleHookStatus struct {
    // Name of the hook
    Name string
    
    ProgressPoint intstr.IntOrString
    
    // States: "Running", "Succeeded", "Failed" 
    State string
    
    // Reference to locate the Job created when executing the hook
    JobRef LocalObjectReference
}

type LifecycleHookRevisionStatus struct {
    // Revision of the object that this hook was run for
    Revision string
    
    LifecycleHookStatuses []LifecycleHookStatus
}
```

### Affected Objects
As of now the lifecycle hooks are relevant for Deployments, StatefulSets and DaemonSets. They will require new optional field to be able to reference LifecycleTemplate and extending its status by []LifecycleHookRevisionStatus. Also the controller will need to be slightly adjusted to scale in appropriate chunks and trigger the hooks.

```go
type DeploymentSpec struct {
    // ...
    // Addition
    LifecycleTemplate *ObjectReference
}

type DeploymentStatus struct {
    // ...
    // Addition
    LifecycleHookRevisionStatuses []LifecycleHookRevisionStatus
}
```

```go
type StatefulSetSpec struct {
    // ...
    // Addition
    LifecycleTemplate *ObjectReference
}

type StatefulSetStatus struct {
    // ...
    // Addition
    LifecycleHookRevisionStatuses []LifecycleHookRevisionStatus
}

```

```go
type DaemonSetSpec struct {
    // ...
    // Addition
    LifecycleTemplate *ObjectReference
}

type DaemonSetStatus struct {
    // ...
    // Addition
    LifecycleHookRevisionStatuses []LifecycleHookRevisionStatus
}
```

## Algorithm
If there is a LifecycleTemplate referenced from an object:

1. Calculate next partition point to reach the closest lifecycle hook progress point and scale replicas in update appropriately. If there is no hook remaining, GOTO 5.
2. When partition point is reached, run the hook by creating a Job using LifecycleHook.JobSpec
3.
    1. If the hook failed and FailurePolicy is Abort - emit event, fail the update and initiate rollback (GOTO 5.) 
    2. If the hook failed and FailurePolicy is Ignore - only emit event
4. GOTO 1.
5. Finish

(In case of rollover cancel any hooks running and don't execute new ones.)

This can be applied to all the existing update strategies as that's essentially dependent only on the ability to have a certain ratio of new.availableReplicas to the declared replicas and running the hook between changing that ratio.
