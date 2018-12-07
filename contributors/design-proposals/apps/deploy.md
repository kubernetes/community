- [Deploy through CLI](#deploy-through-cli)
  - [Motivation](#motivation)
  - [Requirements](#requirements)
  - [Related `kubectl` Commands](#related-kubectl-commands)
    - [`kubectl run`](#kubectl-run)
    - [`kubectl scale` and `kubectl autoscale`](#kubectl-scale-and-kubectl-autoscale)
    - [`kubectl rollout`](#kubectl-rollout)
    - [`kubectl set`](#kubectl-set)
    - [Mutating Operations](#mutating-operations)
    - [Example](#example)
  - [Support in Deployment](#support-in-deployment)
    - [Deployment Status](#deployment-status)
    - [Deployment Revision](#deployment-revision)
    - [Pause Deployments](#pause-deployments)
    - [Failed Deployments](#failed-deployments)


# Deployment rolling update design proposal 

**Author**: @janetkuo

**Status**: implemented

# Deploy through CLI

## Motivation

Users can use [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) or [`kubectl rolling-update`](https://kubernetes.io/docs/tasks/run-application/rolling-update-replication-controller/) to deploy in their Kubernetes clusters. A Deployment provides declarative update for Pods and ReplicationControllers, whereas `rolling-update` allows the users to update their earlier deployment without worrying about schemas and configurations. Users need a way that's similar to `rolling-update` to manage their Deployments more easily.

`rolling-update` expects ReplicationController as the only resource type it deals with. It's not trivial to support exactly the same behavior with Deployment, which requires:
- Print out scaling up/down events.
- Stop the deployment if users press Ctrl-c.
- The controller should not make any more changes once the process ends. (Delete the deployment when status.replicas=status.updatedReplicas=spec.replicas)

So, instead, this document proposes another way to support easier deployment management via Kubernetes CLI (`kubectl`).

## Requirements

The followings are operations we need to support for the users to easily managing deployments:

- **Create**: To create deployments.
- **Rollback**: To restore to an earlier revision of deployment.
- **Watch the status**: To watch for the status update of deployments.
- **Pause/resume**: To pause a deployment mid-way, and to resume it. (A use case is to support canary deployment.)
- **Revision information**: To record and show revision information that's meaningful to users. This can be useful for rollback.

## Related `kubectl` Commands

### `kubectl run`

`kubectl run` should support the creation of Deployment (already implemented) and DaemonSet resources.

### `kubectl scale` and `kubectl autoscale`

Users may use `kubectl scale` or `kubectl autoscale` to scale up and down Deployments (both already implemented).

### `kubectl rollout`

`kubectl rollout` supports both Deployment and DaemonSet. It has the following subcommands:
- `kubectl rollout undo` works like rollback; it allows the users to rollback to a previous revision of deployment.
- `kubectl rollout pause` allows the users to pause a deployment. See [pause deployments](#pause-deployments).
- `kubectl rollout resume` allows the users to resume a paused deployment.
- `kubectl rollout status` shows the status of a deployment.
- `kubectl rollout history` shows meaningful revision information of all previous deployments. See [development revision](#deployment-revision).

### `kubectl set`

`kubectl set` has the following subcommands:
- `kubectl set env` allows the users to set environment variables of Kubernetes resources. It should support any object that contains a single, primary PodTemplate (such as Pod, ReplicationController, ReplicaSet, Deployment, and DaemonSet).
- `kubectl set image` allows the users to update multiple images of Kubernetes resources. Users will use `--container` and `--image` flags to update the image of a container. It should support anything that has a PodTemplate.

`kubectl set` should be used for things that are common and commonly modified. Other possible future commands include:
- `kubectl set volume`
- `kubectl set limits`
- `kubectl set security`
- `kubectl set port`

### Mutating Operations

Other means of mutating Deployments and DaemonSets, including `kubectl apply`, `kubectl edit`, `kubectl replace`, `kubectl patch`, `kubectl label`, and `kubectl annotate`, may trigger rollouts if they modify the pod template.

`kubectl create` and `kubectl delete`, for creating and deleting Deployments and DaemonSets, are also relevant.

### Example

With the commands introduced above, here's an example of deployment management:

```console
# Create a Deployment
$ kubectl run nginx --image=nginx --replicas=2 --generator=deployment/v1beta1

# Watch the Deployment status
$ kubectl rollout status deployment/nginx

# Update the Deployment 
$ kubectl set image deployment/nginx --container=nginx --image=nginx:<some-revision>

# Pause the Deployment
$ kubectl rollout pause deployment/nginx

# Resume the Deployment
$ kubectl rollout resume deployment/nginx

# Check the change history (deployment revisions)
$ kubectl rollout history deployment/nginx

# Rollback to a previous revision.
$ kubectl rollout undo deployment/nginx --to-revision=<revision>
```

## Support in Deployment

### Deployment Status

Deployment status should summarize information about Pods, which includes:
- The number of pods of each revision.
- The number of ready/not ready pods.

See issue [#17164](https://github.com/kubernetes/kubernetes/issues/17164).

### Deployment Revision

We store previous deployment revision information in annotations `kubernetes.io/change-cause` and `deployment.kubernetes.io/revision` of ReplicaSets of the Deployment, to support rolling back changes as well as for the users to view previous changes with `kubectl rollout history`.
- `kubernetes.io/change-cause`, which is optional, records the kubectl command of the last mutation made to this rollout. Users may use `--record` in `kubectl` to record current command in this annotation.
- `deployment.kubernetes.io/revision` records a revision number to distinguish the change sequence of a Deployment's
ReplicaSets. A Deployment obtains the largest revision number from its ReplicaSets and increments the number by 1 upon update or creation of the Deployment, and updates the revision annotation of its new ReplicaSet.

When the users perform a rollback, i.e. `kubectl rollout undo`, the Deployment first looks at its existing ReplicaSets, regardless of their number of replicas. Then it finds the one with annotation `deployment.kubernetes.io/revision` that either contains the specified rollback revision number or contains the second largest revision number among all the ReplicaSets (current new ReplicaSet should obtain the largest revision number) if the user didn't specify any revision number (the user wants to rollback to the last change). Lastly, it
starts scaling up that ReplicaSet it's rolling back to, and scaling down the current ones, and then update the revision counter and the rollout annotations accordingly.

Note that ReplicaSets are distinguished by PodTemplate (i.e. `.spec.template`). When doing a rollout or rollback, a Deployment reuses existing ReplicaSet if it has the same PodTemplate, and its `kubernetes.io/change-cause` and `deployment.kubernetes.io/revision` annotations will be updated by the new rollout. All previous of revisions of this ReplicaSet will be kept in the annotation `deployment.kubernetes.io/revision-history`. For example, if we had 3 ReplicaSets in
Deployment history, and then we do a rollout with the same PodTemplate as revision 1, then revision 1 is lost and becomes revision 4 after the rollout, and the ReplicaSet that once represented revision 1 will then have an annotation `deployment.kubernetes.io/revision-history=1`.

To make Deployment revisions more meaningful and readable for users, we can add more annotations in the future. For example, we can add the following flags to `kubectl` for the users to describe and record their current rollout:
- `--description`: adds `description` annotation to an object when it's created to describe the object.
- `--note`: adds `note` annotation to an object when it's updated to record the change.
- `--commit`: adds `commit` annotation to an object with the commit id.

### Pause Deployments

Users sometimes need to temporarily disable a Deployment. See issue [#14516](https://github.com/kubernetes/kubernetes/issues/14516).

For more details, see [pausing and resuming a
Deployment](https://kubernetes.io/docs/user-guide/deployments/#pausing-and-resuming-a-deployment).

### Failed Deployments

The Deployment could be marked as "failed" when it gets stuck trying to deploy
its newest ReplicaSet without completing within the given deadline (specified
with `.spec.progressDeadlineSeconds`), see document about
[failed Deployment](https://kubernetes.io/docs/user-guide/deployments/#failed-deployment).
