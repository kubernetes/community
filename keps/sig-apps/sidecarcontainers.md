---
title: Sidecar Containers
authors:
  - "@joseph-irving"
owning-sig: sig-apps
participating-sigs:
  - sig-apps
  - sig-node
reviewers:
  - "@fejta"
approvers:
  - "@enisoc"
  - "@kow3ns"
editor: TBD
creation-date: 2018-05-14
last-updated: 2018-11-20
status: provisional
---

# Sidecar Containers

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Alternatives](#alternatives)

## Summary

To solve the problem of container lifecycle dependency we can create a new class of container: a "sidecar container" that behaves primarily like a normal container but is handled differently during termination and startup.

## Motivation

SideCar containers have always been used in some ways but just not formally identified as such, they are becoming more common in a lot of applications and as more people have used them, more issues have cropped up.

Here are some examples of the main problems:

### Jobs
 If you have a Job with two containers one of which is actually doing the main processing of the job and the other is just facilitating it, you encounter a problem when the main process finishes; your sidecar container will carry on running so the job will never finish.

The only way around this problem is to manage the sidecar container's lifecycle manually and arrange for it to exit when the main container exits. This is typically achieved by building an ad-hoc signalling mechanism to communicate completion status between containers. Common implementations use a shared scratch volume mounted into all pods, where lifecycle status can be communicated by creating and watching for the presence of files. This pattern has several disadvantages:

* Repetitive lifecycle logic must be rewritten in each instance a sidecar is deployed.
* Third-party containers typically require a wrapper to add this behaviour, normally provided via an entrypoint wrapper script implemented in the k8s container spec. This adds undesirable overhead and introduces repetition between the k8s and upstream container image specs.
* The wrapping typically requires the presence of a shell in the container image, so this pattern does not work for minimal containers which ship without a toolchain.

### Startup
An application that has a proxy container acting as a sidecar may fail when it starts up as it's unable to communicate until its proxy has started up successfully. Readiness probes don't help if the application is trying to talk outbound.

### Shutdown
Applications that rely on sidecars may experience a high amount of errors when shutting down as the sidecar may terminate before the application has finished what it's doing.


## Goals

Solve issues so that they don't require application modification:
* [25908](https://github.com/kubernetes/kubernetes/issues/25908) - Job completion
* [65502](https://github.com/kubernetes/kubernetes/issues/65502) - Container startup dependencies

## Non-Goals

Allowing multiple containers to run at once during the init phase. //TODO See if we can solve this problem with this proposal

## Proposal

Create a way to define containers as sidecars, this will be an additional field to the Container Spec: `sidecar: true`. //TODO Decide on the API (see [Alternatives](#alternatives))

e.g:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp
    image: myapp
    command: ['do something']
  - name: sidecar
    image: sidecar-image
    sidecar: true
    command: ["do something to help my app"]

```
Sidecars will be started before normal containers but after init, so that they are ready before your main processes start.

This will change the Pod startup to look like this:
* Init containers start
* Init containers finish
* Sidecars start
* Containers start

During pod termination sidecars will be terminated last:
* Containers sent SIGTERM
* Once all Containers have exited: Sidecars sent SIGTERM

If Containers don't exit before the end of the TerminationGracePeriod then they will be sent a SIGKIll as normal, Sidecars will then be sent a SIGTERM with a short grace period of 5/10 Seconds (up for debate) to give them a chance to cleanly exit.

PreStop Hooks will be sent to sidecars and containers at the same time.
This will be useful in scenarios such as when your sidecar is a proxy so that it knows to no longer accept inbound requests but can continue to allow outbound ones until the the primary containers have shut down.  //TODO Discuss whether this is a valid use case (dropping inbound requests can cause problems with load balancers)

To solve the problem of Jobs that don't complete: When RestartPolicy!=Always if all normal containers have reached a terminal state (Succeeded for restartPolicy=OnFailure, or Succeeded/Failed for restartPolicy=Never), then all sidecar containers will be sent a SIGTERM.

### Implementation Details/Notes/Constraints

As this is a fairly large change I think it make sense to break this proposal down and phase in more functionality as we go, potential roadmap could look like:

* Add sidecar field, use it for the shutdown triggering when RestartPolicy!=Always
* Pre-stop hooks sent to sidecars before non sidecar containers
* Sidecars are terminated after normal containers
* Sidecars start before normal containers


As this is a change to the Container spec we will be using feature gating, you will be required to explicitly enable this feature on the api server as recommended [here](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#adding-unstable-features-to-stable-versions).

### Risks and Mitigations

You could set all containers to be `sidecar: true`, this seems wrong, so maybe the api should do a validation check that at least one container is not a sidecar.

Init containers would be able to have `sidecar: true` applied to them as it's an additional field to the container spec, this doesn't currently make sense as init containers are ran sequentially. We could get around this by having the api throw a validation error if you try to use this field on an init container or just ignore the field.

Older Kubelets that don't implement the sidecar logic could have a pod scheduled on them that has the sidecar field. As this field is just an addition to the Container Spec the Kubelet would still be able to schedule the pod, treating the sidecars as if they were just a normal container. This could potentially cause confusion to a user as their pod would not behave in the way they expect, but would avoid pods being unable to schedule.


## Graduation Criteria

//TODO

## Implementation History

- 14th May 2018: Proposal Submitted


## Alternatives

One alternative would be to have a new field in the Pod Spec of `sidecarContainers:` where you could define a list of sidecar containers, however this would require more work in terms of updating tooling to support this.

Another alternative would be to change the Job Spec to have a `primaryContainer` field to tell it which containers are important. However I feel this is perhaps too specific to job when this Sidecar concept could be useful in other scenarios.

Having it as a boolean could cause problems later down the line if more lifecycle related flags were added, perhaps it makes more sense to have something like `lifecycle: Sidecar` to make it more future proof.
