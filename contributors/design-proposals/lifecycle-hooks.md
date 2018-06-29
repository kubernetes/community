# Lifecycle Hooks

**Author**: @tnozicka

**Status**: Proposal

**RFC Issue**: https://github.com/kubernetes/kubernetes/issues/14512

## Abstract
The intent of this proposal is to support lifecycle hooks in all objects having rolling/recreate update; currently those are Deployments, StatefulSets and DaemonSets. 
Lifecycle hooks should allow users to better customize updates (rolling/recreate) without having to write their own controllers. 

## Lifecycle Hooks
Lifecycle hook is a Job that will get triggered by the update reaching certain progress point (like 50%). 
The progress point is determined by the ratio of new.availableReplicas to the declared replicas or by explicitly stating the number of new.availableReplicas. 
The API is open for other types of triggering conditions in the future. 

### Previous Implementations
Lifecycle hooks are already implemented in OpenShift but we want to enhance them and implement them natively in Kubernetes.
 - [OpenShift proposal](https://github.com/openshift/origin/blob/master/docs/proposals/post-deployment-hooks.md)
 - [OpenShift docs](https://docs.openshift.org/latest/dev_guide/deployments/deployment_strategies.html#lifecycle-hooks)

### Previous Kubernetes Proposals
 - https://github.com/kubernetes/kubernetes/pull/33545

### Use Cases
The most common use case for lifecycle hooks is to have pre, mid and/or post hook. 
It is mostly used to run some kind of acceptance check in the middle and/or at the end to fully verify the update is working as expected and rollback if it isn't. 
The acceptance check may be time consuming and more thorough than what readiness and liveness probes are intended for. 
You can also notify external services from them, migrate database in the middle of an update, send messages to IRC channel or do anything else.

[A short demo](https://youtu.be/GVNTm_K43iI) simulating lifecycle hooks using auto-pausing that has been presented at SIG-Apps meeting on August 21, 2017.


## API Objects
### New Objects
```go
type HookTriggerType string

const (
	ProgressAvailableHookTriggerType HookTriggerType = "ProgressAvailable"
)

type ProgressAvailableParams struct {
    // ProgressPoint specifies the point during update when the hook should be triggered.
    // Accepts both the number of new pods available or a percentage value representing the ratio
    // of new.availableReplicas to the declared replicas. In case of getting a percentage
    // it will try to reach the exact or the closest possible point right after it,
    // trigger the job, wait for it to complete and then continue. 
    // If such situation shall occur that two different ProgressPoints should be reached at
    // the same time, all the hooks for an earlier ProgressPoint will be ran (and finished)
    // before any later one.
    ProgressPoint intstr.IntOrString
}

type HookTrigger struct {
    // Type of hook trigger. Can be ProgressAvailable. Default is ProgressAvailable.
    Type HookTriggerType
    
    // ProgressAvailable config params. Present only if type = "ProgressAvailable".
    ProgressAvailableParams *ProgressAvailableParams
}

type LifecycleHook struct {
    // Name of the hook
    Name string
    
    // Triggering config
    Trigger HookTrigger
    
    // "Abort" - Failure for this hook is fatal, deployment becomes perma-failed
    // "Ignore" - Even if the hook fails deployment continues rolling out new version.
    FailurePolicy string
    
    // "Always" - always keep this jobs 
    // "OnFailure" - keep this Job only if it fails
    // "Never" - delete this job immediately after it is done 
    RetainPolicy string
    
    JobSpec v1.JobSpec
}
```

### Affected Objects
As of now the lifecycle hooks are relevant for Deployments, StatefulSets and DaemonSets.
New optional field named `LifecycleHooks` will be added to provide configuration for lifecycle hooks. 
The related state will be reported using Conditions and details will be available on the created Job objects.
Also the controllers will need to be adjusted to scale in appropriate chunks and trigger the hooks.

```go
type DeploymentSpec struct {
    // ...
    // Addition
    // List of lifecycle hooks
    // Can have multiple hooks at the same ProgressPoint; order independent
    LifecycleHooks []LifecycleHook
}
```

```go
type StatefulSetSpec struct {
    // ...
    // Addition
    // List of lifecycle hooks
    // Can have multiple hooks at the same ProgressPoint; order independent
    LifecycleHooks []LifecycleHook
}
```

```go
type DaemonSetSpec struct {
    // ...
    // Addition
    // List of lifecycle hooks
    // Can have multiple hooks at the same ProgressPoint; order independent
    LifecycleHooks []LifecycleHook
}
```

## Algorithm
If there are LifecycleHooks present in an object:

1. Calculate next partition point to reach the closest lifecycle hook progress point and scale replicas in update appropriately. If there is no hook remaining, GOTO 5.
2. When partition point is reached, run the hook by creating a Job using LifecycleHook.JobSpec. 
(Controller is allowed to run more hooks for the same progress point in parallel but has to wait for all of them to finish before it can move to the next progress point.)
3.
    1. If the hook failed and FailurePolicy is Abort - emit event, permanently fail the update (GOTO 5.) 
    2. If the hook failed and FailurePolicy is Ignore - only emit event
4. GOTO 1.
5. Finish

(In case of rollover cancel any hooks running and don't execute new ones.)

This can be applied to all the existing update strategies as that's essentially dependent only on the ability to have a certain ratio of new.availableReplicas to the declared replicas and running the hook between changing that ratio.

Hooks (created Jobs) are tracked using ownerRefs and by being labeled with a revision they were created for. By having an ownerRef to particular controllerRevision (or RS) pruning will propagate there.

## Injecting deployment information into hooks
To help scripts make more informed decisions following environment variables will be injected into hook jobs.

`K8S_ROLLOUT_FROM_OBJECT` and `K8S_ROLLOUT_TO_OBJECT` will allow scripts to use kubectl or API to get more details.
```bash
K8S_ROLLOUT_FROM_OBJECT="namespace/name"
K8S_ROLLOUT_TO_OBJECT="namesapce/name"
```

`K8S_ROLLOUT_TYPE` will help scripts to easily detect a rollback.
```
K8S_ROLLOUT_TYPE=[ROLLOUT,ROLLBACK]
```

In case any of these environment variables is already present on the object, it won't be set by the controller.
