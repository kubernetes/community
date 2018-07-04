---
kep-number: 0
title: New pod restartPolicy to restart the whole pod instead of just a container
authors:
  - "@amshuman-kr"
owning-sig: sig-node
participating-sigs:
  - sig-apps
reviewers:
  - "@mtaufen"
  - "@smarterclayton"
approvers:
  - "@liggitt"
  - "@derekwaynecarr"
editor:
creation-date: 2018-07-04
last-updated: 2018-07-04
status: provisional
see-also:
replaces:
superseded-by:
---

# New pod restartPolicy to restart the whole pod instead of just a container

## Table of Contents
* [New pod restartPolicy to restart the whole pod instead of just a container](#new-pod-restartpolicy-to-restart-the-whole-pod-instead-of-just-a-container)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Motivation](#motivation)
    * [Goals](#goals)
    * [Non\-Goals](#non-goals)
  * [Proposal](#proposal)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
      * [Recreate Pod Sandbox](#recreate-pod-sandbox)
        * [Pros](#pros)
        * [Cons](#cons)
      * [Re\-use Pod Sandbox](#re-use-pod-sandbox)
        * [Pros](#pros-1)
        * [Cons](#cons-1)
    * [Risks and Mitigations](#risks-and-mitigations)
  * [Graduation Criteria](#graduation-criteria)
  * [Implementation History](#implementation-history)
  * [Alternatives](#alternatives)

## Summary

In a pod with multiple containers, if one of the containers terminates and if the restart policy mandates that it be restarted (as in [`restartPolicy`][rp] of `OnFailure` or `Always`) then the kubelet restarts only the terminated container. The other containers are left as they are.

This behaviour, while generally desirable, makes some scenarios such as the ones involving [initContainers][issue] or some complex interaction between the containers of the pod cumbersome to implement.

To address such scenarios, this proposal introduces a new pod `restartPolicy` called `AlwaysPod` to make it possible to restart the whole pod (including the initContainers) whenever the `restartPolicy` `Always` would have restarted just one of the containers of the pod.

## Motivation

The `OnFailure` and `Always` [restart polices][rp] efficiently manage the life-cycle of the containers of a pod. The support for multiple containers in a pod also enable better modularity and separation of concerns between different containers. It also promotes looser coupling between components.

[Init containers] provide some additional support for modularity and looser coupling for the functionality of initialization of the pod. They make it possible to separate the initialization from the rest of the pod to enhance both modularity as well as security.

But both the `OnFailure` as well as the `Always` restart policies restart the individual containers in question and not the whole pod. This is, for the most part, desirable, even optimal.

However, there are scenarios (some documented in [this issue][issue]) where the many containers in the pod (including init containers) might be interlinked or inter-dependent in such a way as to require closer co-ordination when any one of its containers are restarted.

### Goals

Make it possible to declaratively specify that the whole pod (including the init containers) should be restarted if any container is going to be restarted.

This can simplify many scenarios requiring close co-ordination of containers of a pod during individual container restart.

For example, if init containers are used to verify, initialize and if necessary restore from backup the data for some persistent services, then restarting the pod when any of its regular containers crash or restart would make sure that the data is always consistent and ready before the regular containers are started or restarted. This can make the services more self-managed. It can also enhance the reach of the init containers into many other use-cases where they cannot be used right now.

Many other use-cases are documented in the [upstream issue][issue].

### Non-Goals

It is not a goal of this proposal to address all possible ways to co-ordinate between the containers of the same pod.
The only goal is to address the only co-ordination of restarting of containers of the same pod.

## Proposal

Introduce a new value (`AlwaysPod`) for `restartPolicy` which works almost exactly like `Always` except that whenever the `Always` restart policy would have restarted any one of the containers of a pod, the `AlwaysPod` restart policy would restart the whole pod (including re-executing it's init containers).

### Implementation Details/Notes/Constraints

#### Recreate Pod Sandbox

We can trigger the pod restart by triggering the recreation of the [pod sandbox][sandbox]. A work-in-progress implementation of this approach can be found [in this pull request][pr].

##### Pros
  * This approach as the benefit of having only a small amount of change to the existing code-base.
  * It also leverages the existing mechanisms for pod restart such as during changes to the pod specification.
##### Cons
  * This approach is sub-optimal in that it discards the existing sandbox and the associated container instances.

#### Re-use Pod Sandbox

We can think of re-using the existing pod sandbox and merely restart the existing container instances in the right order. This is made a bit more challenging by the fact that the kublet currently optimizes by [pruning][prune] the container instances of successfully executed init containers.

##### Pros
  * This approach is more optimal.
##### Cons
  * This approach is more complex.
  * The impact on the rest of the kubelet behaviour also might be larger.

### Risks and Mitigations

The `restartPolicy` or `AlwaysPod` would be a new value for an existing field in the pod specefication. So, the question of backward compatibility may not apply.

## Graduation Criteria

## Implementation History

## Alternatives

Co-ordination between containers can be achieved by baking in some mechanism into the individual containers of the pod. 
But this has the limitations that the some customisation would be required in the source code or at least the image of the individual containers. Such customisations may not always be possible or desirable.

The init containers already address this need for some custom behaviour without modifying the individual container images. But the scope of init containers is somewhat limited by the current behaviour during container restart.

[rp]: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy 
[issue]: https://github.com/kubernetes/kubernetes/issues/52345
[ic]: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
[sandbox]: https://github.com/kubernetes/kubernetes/blob/e3fa9133af250ee8abfc84b200c6495612114c08/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L383
[pr]: https://github.com/kubernetes/kubernetes/pull/65619
[prune]: https://github.com/kubernetes/kubernetes/blob/e3fa9133af250ee8abfc84b200c6495612114c08/pkg/kubelet/kuberuntime/kuberuntime_container.go#L621
