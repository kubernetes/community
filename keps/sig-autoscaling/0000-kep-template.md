---
kep-number: 0
title: My First KEP
authors:
  - "@janedoe"
owning-sig: sig-xxx
participating-sigs:
  - sig-aaa
  - sig-bbb
reviewers:
  - TBD
  - "@alicedoe"
approvers:
  - TBD
  - "@oscardoe"
editor: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: provisional
see-also:
  - KEP-1
  - KEP-2
replaces:
  - KEP-3
superseded-by:
  - KEP-100
---

# Integrating cluster autoscaler (CA) with cluster-api

The *filename* for the KEP should include the KEP number along with the title.
The title should be lowercased and spaces/punctuation should be replaced with `-`.
As the KEP is approved and an official KEP number is allocated, the file should be renamed.

To get started with this template:
1. **Pick a hosting SIG.**
  Make sure that the problem space is something the SIG is interested in taking up.
  KEPs should not be checked in without a sponsoring SIG.
1. **Allocate a KEP number.**
  Do this by (a) taking the next number in the `NEXT_KEP_NUMBER` file and (b) incrementing that number.
  Include the updated `NEXT_KEP_NUMBER` file in your PR.
1. **Make a copy of this template.**
  Name it `NNNN-YYYYMMDD-my-title.md` where `NNNN` is the KEP number that was allocated.
1. **Fill out the "overview" sections.**
  This includes the Summary and Motivation sections.
  These should be easy if you've preflighted the idea of the KEP with the appropriate SIG.
1. **Create a PR.**
  Assign it to folks in the SIG that are sponsoring this process.
1. **Merge early.**
  Avoid getting hung up on specific details and instead aim to get the goal of the KEP merged quickly.
  The best way to do this is to just start with the "Overview" sections and fill out details incrementally in follow on PRs.
  View anything marked as a `provisional` as a working document and subject to change.
  Aim for single topic PRs to keep discussions focused.
  If you disagree with what is already in a document, open a new PR with suggested changes.

The canonical place for the latest set of instructions (and the likely source of this file) is [here](/keps/0000-kep-template.md).

The `Metadata` section above is intended to support the creation of tooling around the KEP process.
This will be a YAML section that is fenced as a code block.
See the KEP process for details on each of these items.

## Table of Contents

TPB

## Summary

TPB

## Motivation

- delegate the responsibility of managing cloud providers out of the cluster autoscaler
- in general, every tool has its own implementation of the cloud provider layer to communicate with cloud provider API (logic duplicated) => cluster-api project



This section is for explicitly listing the motivation, goals and non-goals of this KEP.
Describe why the change is important and the benefits to users.
The motivation section can optionally provide links to [experience reports][] to demonstrate the interest in a KEP within the wider Kubernetes community.

[experience reports]: https://github.com/golang/go/wiki/ExperienceReports

### Goals

* cluster autoscaler is capable of scaling the cluster through machine API

### Non-Goals

* cluster autoscaler is able to autoprovision new node groups through machine API
* cluster autoscaler is able to estimate node resource requirements through machine API
* cluster autoscaler is able to use at least one pricing model on top of the machine API

## Proposal

Generalize the concept of node group autoscaling on the level of cloud provider
by integrating with machine API layer of [sigs.k8s.io/cluster-api](https://github.com/kubernetes-sigs/cluster-api) project to
build cloud provider-free implementation of cluster autoscaling mechanism
([#1050](https://github.com/kubernetes/autoscaler/issues/1050)).

Also suggest how to:
* perform node autoprovision on the level of the machine API ([node autoprovisioning
  upstream proposal](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/proposals/node_autoprovisioning.md)).
* perform estimation of cluster autoscaling on the level of the machine API.
* build various pricing models on top of the machine API.

Cluster autoscaling builds on top of concept of node groups which effectively
translates to cloud provider specific autoscaling groups, resp. autoscaling sets.
In case of the machine API, the node group translates into machine set (living
as native Kubernetes object).

### Node group (scaling group)

The autoscaler currently supports `AWS`, `Azure`, `GCE`, `GKE` providers.
With `kubemark` as other cloud provider implementation that allows to create
hollow kubernetes cluster with huge number of nodes over a small subset of real
machines (for scaling testing purposes).
Each of the currently supported cloud providers has its own concept of scaling groups:
- AWS has autoscaling groups
- Azure has autoscaling sets
- GCE/GKE have autoscaling groups

Autoscaling group allows to flexibly and natively increase/decrease number of instances
within cloud provider based on actual cluster workload.

TODO(jchaloup): describe advantages of auto scaling groups, why they are beneficial, etc.

### Machine API

Machine API (as part of the [cluster-api project](https://github.com/kubernetes-sigs/cluster-api/))
allows to declaratively specify a machine without saying how to provision underlying instance.
Machines can be grouped into a machine set that allows to specify a machine template
and by setting a number replicas create required number of machines (followed
by provisioning of underlying instances).
With some mama's marmalade, one can represent an autoscaling group with a machine set.
Though, the current implementation of the machineset controller needs to be improved
to support autoscaling use cases. E.g. define delete strategy to pick appropriate machine
to delete from the machine set ([#75](https://github.com/kubernetes-sigs/cluster-api/issues/75)).

`Important`: future implementation of the autoscaling concept through machineset will not
replace the current cloud specific implementation. The goal is to lay down a solution
that will cover the most common use cases currently implemented by the autoscaler.
It's important to draw a line into the sand so only necessary portion of the cloud
autoscaling logic is re-implemented instead of re-inventing entire wheel.

### Cloud provider interface
The cluster autoscaler defines two interfaces (`NodeGroup` and `CloudProvider`
from `k8s.io/autoscaler/cluster-autoscaler/cloudprovider` package)
that a cloud provider needs to implemented in order to allow the cluster
autoscaler to properly perform scaling operations.

The `CloudProvider` interface allows to:
- operate on top of node groups
- work with pricing models (to estimate reasonable expenses)
- check limit resources (check if maximum resources per node group are exceeded)

The `NodeGroup` interface allows to operate with node groups themselves:
- change size of a node group
- list nodes belonging to a node group
- autoprovision new node group
- generate node template for a node group (to determine resource capacity of
  new node in case a node group has zero size)

In case of `AWS` the `NodeGroup` then corresponds to autoscaling groups,
in case of `Azure` to autoscaling sets, etc.

### Node template

In case a node group has no node, the cluster autoscaler needs a way to generate
a node template for a given node group before it can predict how much new resources
will get allocated and thus how many new nodes are needed to scale up.

Each of the mentioned cloud providers has its own implementation of constructing
the node template with its own list of information pulled from cloud provider API:

* Azure:
  - instance type to determine resource capacity (CPU, GPU, Memory)
  - instance tags to construct node labels
  - generic node labels: Arch, OS, instance type, region, failure domain, hostname
* AWS:
  - instance type to determine resource capacity (CPU, GPU, Memory)
  - instance tags to construct node labels and taints
  - generic node labels: Arch, OS, instance type, region, failure domain, hostname
* GCE:
  - instance type to determine resource capacity (CPU, Memory)
  - instance tags to construct node labels (e.g. gpu.ResourceNvidiaGPU)
  - maximum number of pods: apiv1.ResourcePods
  - taints (from instance metadata)
  - kubelet system and reserved allocatable

## Code status

At this point the cluster autoscaler does not support node autoprovisioning (TODO(jchaloup): link to the proposal).
Every implementation assumes every node group already exists in the cloud provider.
Both `Create` and `Delete` operations return an error `ErrAlreadyExist` and the autoprovisioning is off.

The autoscaler allows to either enumerate all node groups (via `--nodes` option)
or automatically discover all node groups on-the-fly (via `--node-group-auto-discovery`).

### Integration with machine API

Given the scaling up/down operation corresponds to increasing/decreasing the number
of nodes in a node group, it's sufficient to "just" change the number of replicas
of corresponding machine set.

#### Node template

The `machine` object holds all cloud provider specific information under
`spec.providerConfig` field. Given only the actuator can interpret the provider
configuration and there can be multiple actuator implementations for the same cloud provider,
all information necessary to render the template must be contained outside of the machine
's provider config. At the same time the `machineset` object works over a set of
`machines` and has no way of communicating with the machine controller.

The node template corresponding to a node group has to be either rendered
on the cluster autoscaler side or inside the machine controller.
In the former case the machine (outside of the provider config) needs to contain
all the information in a generic form to render the template (e.g. through labels)
or import the actuator code that can interpret the provider config. In the latter,
the machine controller has to store the rendered node template (that is free
of cloud provider specifics) into machine object.

* **Labels**: Label use case may lead to data duplication as one needs to provide a cpu, gpu
  and memory requirements (based on instance type) to specify node's allocatable
  resources. Or other information such as region, availability zones (which may
  have different representation through different cloud providers).
  Implemented through labels the `machine` spec can look like:

  ```yaml
  ---
  apiVersion: cluster.k8s.io/v1alpha1
  kind: Machine
  metadata:
    name: node-group-machine
    namespace: application
    labels:
      sigs.k8s.io/cluster-autoscaler-resource-capacity: "cpu:200m,gpu:1,memory:4GB"
      sigs.k8s.io/cluster-autoscaler-region: us-west-1
  ...
  ```

* **Actuator code imports**: Importing the actuator code and pulling all
  the information from the `spec.providerConfig` forces the cluster autoscaler
  to choose a single implementation of actuator for each cloud provider.
  That can be to restrictive.

* **Rendering inside the machine controller**: From the "I only know what I need to know"
  point of view, this approach encapsulates all the knowledge about provider
  configuration within the actuator itself. The actuator knows the best how to
  properly construct the node template. Obviously, the node template is stored
  in the machine object. Thus, any client querying the machine can then read
  and consume the template free of any cloud provider. Given the node template
  can reflect the current state of the underlying machine's instance
  (e.g. instance tags) the node template can be periodically rendered
  and published in machine's status.

#### Scaling from 0 

In all cases, if it can be assumed all machines/nodes in the same node group have the same
system requirements, one can use generic kubelet's configuration discovery to get
the kubelet's reserved and system requirements from the first machine in a given
node group (given there is at least one node in the group).
The same holds for node taints.


### User Stories [optional]

TPB

### Implementation Details/Notes/Constraints [optional]

TPB

### Risks and Mitigations

TPB

## Graduation Criteria

TPB

## Implementation History

TPB

## Drawbacks [optional]

TPB

## Alternatives [optional]

TPB

## Infrastructure Needed [optional]

TPB
