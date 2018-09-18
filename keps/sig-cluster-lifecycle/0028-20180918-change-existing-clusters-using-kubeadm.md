---
kep-number: 28
title: Change existing clusters using kubeadm
authors:
    - "@fabriziopandini"
owning-sig: sig-cluster-lifecycle
reviewers:
    - "@chuckha"
    - "@detiber"
    - "@liztio"
    - "@neolit123"
approvers:
    - "@luxas"
    - "@timothysc"
editor:
    - "@fabriziopandini"
creation-date: 2018-09-18
status: provisional
last-updated:
---

# Change existing clusters using kubeadm

## Table of Contents

<!-- TOC -->

- [Change existing clusters using kubeadm](#change-existing-clusters-using-kubeadm)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
    - [Motivation](#motivation)
        - [Goals](#goals)
        - [Non-Goals](#non-goals)
    - [Proposal](#proposal)
        - [User Stories](#user-stories)
        - [Implementation Details](#implementation-details)
            - [How to define target state of the cluster](#how-to-define-target-state-of-the-cluster)
            - [How to define components to be modified](#how-to-define-components-to-be-modified)
        - [Notes](#notes)
        - [Constraints](#constraints)
        - [Risks and Mitigations](#risks-and-mitigations)
    - [Graduation Criteria](#graduation-criteria)
    - [Implementation History](#implementation-history)
    - [Drawbacks](#drawbacks)
    - [Alternatives](#alternatives)
    - [Infrastructure Needed](#infrastructure-needed)

<!-- /TOC -->

## Summary

Sometimes it is necessary to change attributes of an existing Kubernetes cluster: enable a new
api group in the api server, change kubelet settings, etc. This KEP is about creating a
new kubeadm sub-command that will allow to change Kubernetes cluster attributes in a
controlled way.

## Motivation

Kubeadm currently allows you to init a Kubernetes cluster, join nodes and upgrade the cluster
to a newer version, but there is no a dedicated workflow for changing cluster attributes.

According to users feedbacks in the kubeadm repo/slack channel, today two workarounds are
used for changing cluster attributes:

1. Editing manually manifest files under `\etc\kubernetes\manifests` or component config
   ConfigMaps in the `kube-system` namespace

2. Running `kubeadm upgrade apply` with a modified kubeadm configuration file but with
   `new KubernetesVersion == current KubernetesVersion` (afterwards referred as `kubeadm upgrade
   apply` workaround)

While the possible risks of the first approach are easy to understand, also the
`kubeadm upgrade apply` workaround has many potential pitfalls.

First of all, there is no control on which attributes the user can *safely* change, and
this can potentially lead to unstable/nonfunctional clusters. e.g. if the user changes
the `controlPlaneAddress` or the `advertiseAddress`, all the nodes probably will stop working.

The second source of potential problems is that the `kubeadm upgrade apply` workflow basically
re-recreates all the kubernetes core components, no matter if a component is impacted by
the desired changes of attributes or not, e.g. even if only kubelet attributes are changing,
kubeadm upgrade apply re-creates the api-server, controller-manager, scheduler, addons etc.

Recreating all the kubernetes core components adds un-necessary risks and complexity to this
operations and makes any potential problem more difficult to be investigated.

Another subtle pitfall of the `kubeadm upgrade apply` workaround is that the operation, once
completed, doesn't provide to the user any instruction about the additional, mandatory steps
that are required to make some type of changes really effective. e.g. when you change kubelet
attributes, those attributes are applied to the kubelet component config ConfigMap and to
the current node only; to make the change effective on all the other nodes, the user should
execute `kubeadm upgrade node` on each of them.

### Goals

This KEP aims to avoid/limiting all the above risks and pitfalls by:

- Implement a new kubeadm sub command that allows user to change cluster attributes in a more
  controlled way, and more specifically:
  - To validate requested changes, and block or provide advice for potentially disruptive
    changes.
  - To reduce the complexity and the risk of the set of actions to be performed by kubeadm,
    by determining the minimum set of components to be modified for executing the requested
    changes.
  - When necessary, to provide the users with the instructions about additional, mandatory
    steps required for making changes really effective.

- Validate the above approach by implementing support for changing a limited set of
  cluster attributes:
  - Controlplane components `ExtraArgs` and `HostVolumes`.
  - Component configs ConfigMaps.

- Once the approach is validated:
  - Lock down the `kubeadm upgrade apply` workaround.
  - Progressively extend the same approach to other attributes.

### Non-Goals

- To distill all the knowledge/lesson learned for changing all the existing cluster attributes - and
  combination of attributes -. This will take take time and potentially further iterations on this KEP.
- To manage CoreDNS settings - currently based on a custom format / with a different
  approach then other component configs

## Proposal

### User Stories

TBD

### Implementation Details

In order to implement a new kubeadm sub command that allows user to change cluster
attributes in a more controlled way following high level problems should be addressed:

1. How to define target state of the cluster
2. How to define components to be modified

#### How to define target state of the cluster

A Kubernetes cluster created by kubeadm is described by the API Objects stored in
the `kubeadm-config` ConfigMap and by the component config ConfigMaps.

Similarly to what already in place for changing other Kubernetes API Objects (rif.
`kubectl apply -f` or `kubectl patch`), the initial working assumption is that the
requested change of cluster attributes could be be specified by providing one among:

- a new version of the above ConfigMaps.
- a patch to be applied on the current ConfigMaps.

Given requested changes, then kubeadm will use the [strategic merge patch](https://github.com/kubernetes/community/blob/master/contributors/devel/strategic-merge-patch.md)
for computing the target state of API Objects. More specifically:

```
Target ConfigMaps = ([current ConfigMaps] strategic merge patch [new ConfigMaps/Patch])
```

#### How to define components to be modified

Once the changes to be applied are known, it is necessary to determine the minimum set
of components is impacted by the change.

The first step of this activity is to determine the cluster attributes that should change:

```
Changed attributes = diff between [current ConfigMaps] and [target ConfigMaps]
```

Given the cluster attributes that should change, it is necessary to determine the minimum
set of components to be modified accordingly.

This kind of knowledge currently is hard coded in kubeadm, and thus not easily accessible.
It is then necessary to encode the above knowledge into a machine readable format, to be used
by the new kubeadm sub command introduced by this KEP.

The initial working assumption for solving this problem leverage on go struct tags. e.g

```go
type ClusterConfiguration struct {
	....

	APIServerExtraArgs map[string]string  `update:"apiserver"`

	ControllerManagerExtraArgs map[string]string  `update:"controller-manager"`

	SchedulerExtraArgs map[string]string  `update:"scheduler"`

	ImageRepository string  `update:"apiserver,controller-manager,scheduler,etcd"`

	...
}
```

The same approach can be used also for defining what attributes are not allowed to change e.g.

```go
type ClusterConfiguration struct {
	....

	KubernetesVersion string  `update:"KubernetesVersion cannot be changed using kubeadm change cluster (name TBD). Use kubeadm upgrade apply instead "`

	ControlPlaneEndpoint string `update:"blocked,msg=Updates to ControlPlaneEndpoint are not supported in this release."` 

	...
}
```

### Notes

TBD

### Constraints

TBD

### Risks and Mitigations

TBD

- Struct tags not validated at compile time --> mitigation test
- Status of changes on different nodes is not tracked by kubeadm --> take the risk in this initial
  version and see if this generates problems for the user (giving instruction only is better than
  current situation)

## Graduation Criteria

TBD

## Implementation History

TBD

## Drawbacks

TBD

- Core DNS not covered

- Distill all the knowledge/lesson learned for all the existing attributes takes times
  and should generate new requisites --> divide et impera

## Alternatives

TBD

## Infrastructure Needed

TBD