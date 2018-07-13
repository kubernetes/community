# kubeadm join --master workflow

## Metadata

```yaml
---
kep-number: 15
title: kubeadm join --master workflow
status: accepted
authors:
  - "@fabriziopandini"
owning-sig: sig-cluster-lifecycle
reviewers:
  - "@chuckha”
  - "@detiber"
  - "@luxas" 
approvers:
  - "@luxas"
  - "@timothysc"
editor:
  - "@fabriziopandini"
creation-date: 2018-01-28
last-updated: 2018-06-29
see-also:
  - KEP 0004
```

## Table of Contents

<!-- TOC -->

- [kubeadm join --master workflow](#kubeadm-join---master-workflow)
    - [Metadata](#metadata)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
    - [Motivation](#motivation)
        - [Goals](#goals)
        - [Non-goals](#non-goals)
        - [Challenges and Open Questions](#challenges-and-open-questions)
    - [Proposal](#proposal)
        - [User Stories](#user-stories)
            - [Create a cluster with more than one master nodes (static workflow)](#create-a-cluster-with-more-than-one-master-nodes-static-workflow)
            - [Add a new master node (dynamic workflow)](#add-a-new-master-node-dynamic-workflow)
        - [Implementation Details](#implementation-details)
            - [Initialize the Kubernetes cluster](#initialize-the-kubernetes-cluster)
            - [Preparing for execution of kubeadm join --master](#preparing-for-execution-of-kubeadm-join---master)
            - [The kubeadm join --master workflow](#the-kubeadm-join---master-workflow)
            - [dynamic workflow (advertise-address == `controlplaneAddress`)](#dynamic-workflow-advertise-address--controlplaneaddress)
            - [Static workflow (advertise-address != `controlplaneAddress`)](#static-workflow-advertise-address--controlplaneaddress)
            - [Strategies for deploying control plane components](#strategies-for-deploying-control-plane-components)
            - [Strategies for distributing cluster certificates](#strategies-for-distributing-cluster-certificates)
            - [`kubeadm upgrade` for HA clusters](#kubeadm-upgrade-for-ha-clusters)
    - [Graduation Criteria](#graduation-criteria)
    - [Implementation History](#implementation-history)
    - [Drawbacks](#drawbacks)
    - [Alternatives](#alternatives)

<!-- /TOC -->

## Summary

We are extending the kubeadm distinctive `init` and `join` workflow, introducing the
capability to add more than one master node to an existing cluster by means of the
new `kubeadm join --master` option (in alpha release the flag will be named --experimental-master)

As a consequence, kubeadm will provide a best-practice, “fast path” for creating a
minimum viable, conformant Kubernetes cluster with one or more master nodes and
zero or more worker nodes; as better detailed in following paragraphs, please note that
this proposal doesn't solve every possible use case or even the full end-to-end flow automatically.

## Motivation

Support for high availability is one of the most requested features for kubeadm.

Even if, as of today, there is already the possibility to create an HA cluster
using kubeadm in combination with some scripts and/or automation tools (e.g.
[this](https://kubernetes.io/docs/setup/independent/high-availability/)), this KEP was
designed with the objective to introduce an upstream simple and reliable solution for
achieving the same goal.

Such solution will provide a consistent and repeatable base for implementing additional
capabilities like e.g. kubeadm upgrade for HA clusters.

### Goals

- "Divide and conquer”

  This proposal - at least in its initial release - does not address all the possible
  user stories for creating an highly available Kubernetes cluster, but instead
  focuses on:

  - Defining a generic and extensible flow for bootstrapping a cluster with multiple masters,
    the `kubeadm join --master` workflow.
  - Providing a solution *only* for well defined user stories. see
    [User Stories](#user-stories) and [Non-goals](#non-goals).

- Enable higher-level tools integration

  We expect higher-level and tooling will leverage on kubeadm for creating HA clusters;
  accordingly, the `kubeadm join --master` workflow should provide support for
  the following operational practices used by higher level tools:

  - Parallel node creation

    Higher-level tools could create nodes in parallel (both masters and workers)
    for reducing the overall cluster startup time.
    `kubeadm join --master` should support natively this practice without requiring
    the implementation of any synchronization mechanics by higher-level tools.

- Provide support both for dynamic and static bootstrap flow

  At the time a user is running `kubeadm init`, they might not know what
  the cluster setup will look like eventually. For instance, the user may start with
  only one master + n nodes, and then add further master nodes with `kubeadm join --master`
  or add more worker nodes with `kubeadm join` (in any order). This kind of workflow, where the
  user doesn’t know in advance the final layout of the control plane instances, into this
  document is referred as “dynamic bootstrap workflow”.

  Nevertheless, kubeadm should support also more “static bootstrap flow”, where a user knows
  in advance the target layout of the controlplane instances (the number, the name and the IP
  of master nodes).

- Support different etcd deployment scenarios, and more specifically run master nodes components
  and the etcd cluster on the same machines (stacked control plane nodes) or run the etcd
  cluster on dedicated machines.

### Non-goals

- Graduating an existing node to master.
  The nodes must be created as a master or as workers and then are supposed to stick to the assigned role
  for their entire life cycle.

- This proposal doesn't include a solution for etcd cluster management (but nothing in this proposal should
  prevent to address this in future).

- This proposal doesn't include a solution for API server load balancing (Nothing in this proposal
  should prevent users from choosing their preferred solution for API server load balancing).

- This proposal doesn't address the ongoing discussion about kubeadm self-hosting; in light of
  divide and conquer goal stated before, it is not planned to provide support for self-hosted clusters
  neither in the initial proposal nor in the foreseeable future (but nothing in this proposal should
  explicitly prevent to reconsider this in future as well).

- This proposal doesn't provide an automated solution for transferring the CA key and other required
  certs from one master to the other. More specifically, this proposal doesn't address the ongoing
  discussion about storage of kubeadm TLS assets in secrets and it it is not planned
  to provide support for clusters with TLS stored in secrets (but nothing in this
  proposal should explicitly prevent to reconsider this in future).

- Nothing in this proposal should prevent practices that exist today.

### Challenges and Open Questions

- Keep the UX simple.

  - _What are the acceptable trade-offs between the need to have a clean and simple
    UX and the variety/complexity of possible kubernetes HA deployments?_

- Create a cluster without knowing its final layout

  Supporting a dynamic workflow implies that some information about the cluster are
  not available at init time, like e.g. the number of master nodes, the IP of
  master nodes etc. etc.

  - _How to configure a Kubernetes cluster in order to easily adapt to future change
    of its own controlplane layout like e.g. add a master node, remove a master node?_

  - _What are the "pivotal" cluster settings that must be defined before initialising
    the cluster?_

  - _How to combine into a single UX support for both static and dynamic bootstrap
    workflows?_

- Kubeadm limited scope of action

  - Kubeadm binary can execute actions _only_ on the machine where it is running
    e.g. it is not possible to execute actions on other nodes, to copy files across
    nodes etc.
  - During the join workflow, kubeadm can access the cluster _only_ using identities
    with limited grants, namely `system:unauthenticated` or `system:node-bootstrapper`.

- Upgradability

  - How to setup an high available cluster in order to simplify the execution
    of cluster version upgrades, both manually or with the support of `kubeadm upgrade`?_

## Proposal

### User Stories

#### Create a cluster with more than one master nodes (static workflow)

As a kubernetes administrator, I want to create a Kubernetes cluster with more than one
master nodes*, of which I know in advance the name and the IP.

\* A new "master node" is a new kubernetes node with
`node-role.kubernetes.io/master=""` label and
`node-role.kubernetes.io/master:NoSchedule` taint; a new instance of control plane
components will be deployed on the new master node.
As described in goals/non goals, in this first release of the proposal
creating a new master node doesn't trigger the creation of a new etcd member on the
same machine.

#### Add a new master node (dynamic workflow)

As a kubernetes administrator, (_at any time_) I want to add a new master node* to an existing
Kubernetes cluster.

### Implementation Details

#### Initialize the Kubernetes cluster

As of today, a Kubernetes cluster should be initialized by running `kubeadm init` on a
first master, afterward referred as the bootstrap master.

in order to support the `kubeadm join --master` workflow a new Kubernetes cluster is
expected to satisfy following conditions :

- The cluster must have a stable `controlplaneAddress` endpoint (aka the IP/DNS of the
  external load balancer)
- The cluster must use an external etcd.

All the above conditions/settings could be set by passing a configuration file to `kubeadm init`.

#### Preparing for execution of kubeadm join --master

Before invoking `kubeadm join --master`, the user/higher level tools
should copy control plane certificates from an existing master node, e.g. bootstrap master

> NB. kubeadm is limited to execute actions *only*
> in the machine where it is running, so it is not possible to copy automatically
> certificates from remote locations.

Please note that strictly speaking only ca, front-proxy-ca certificate and and service account key pair
are required to be equal among all masters. Accordingly:

- `kubeadm join --master` will check for the mandatory certificates and fail fast if
  they are missing
- given the required certificates exists, if some/all of the other certificates are provided
  by the user as well, `kubeadm join --master` will use them without further checks.
- If any other certificates are missing, `kubeadm join --master` will create them.

> see "Strategies for distributing cluster certificates" paragraph for
> additional info about this step.

#### The kubeadm join --master workflow

The `kubeadm join --master` workflow will be implemented as an extension of the
existing `kubeadm join` flow.

`kubeadm join --master` will accept an additional parameter, that is the apiserver advertise
address of the joining node; as details in following paragraphs, the value assigned to
this parameter depends on the user choice between a dynamic bootstrap workflow or a static
bootstrap workflow.

The updated join workflow will be the following:

1. Discovery cluster info [No changes to this step]

   > NB This step waits for a first instance of the kube-apiserver to become ready
   > (the bootstrap master); And thus it acts as embedded mechanism for handling the sequence
   > `kubeadm init` and `kubeadm join` actions in case of parallel node creation.

2. Executes the kubelet TLS bootstrap process [No changes to this step]:

3. In case of `join --master` [New step]

   1. Using the bootstrap token as identity, read the `kubeadm-config` configMap
      in `kube-system` namespace.

      > This requires to grant access to the above configMap for
      > `system:bootstrappers` group.

   2. Check if the cluster is ready for joining a new master node:

      a. Check if the cluster has a stable `controlplaneAddress`
      a. Check if the cluster uses an external etcd
      a. Checks if the mandatory certificates exists on the file system

   3. Prepare the node for joining as a master node:

      a. Create missing certificates (in any).
         > please note that by creating missing certificates kubeadm can adapt seamlessly
         > to a dynamic workflow or to a static workflow (and to apiserver advertise address
         > of the joining node). see following paragraphs for more details for additional info.

      a. In case of control plane deployed as static pods, create related kubeconfig files
      and static pod manifests.

      > see "Strategies for deploying control plane components" paragraph
      > for additional info about this step.

   4. Create the admin.conf kubeconfig file

      > This operation creates an additional root certificate that enables management of the cluster
      > from the joining node and allows a simple and clean UX for the final steps of this workflow
      > (similar to the what happen for `kubeadm init`).
      > However, it is important to notice that this certificate should be treated securely
      > for avoiding to compromise the cluster.

   5. Apply master taint and label to the node.

   6. Update the `kubeadm-config` configMap with the information about the new master node

#### dynamic workflow (advertise-address == `controlplaneAddress`)

There are many ways to configure an highly available cluster.

Among them, the approach best suited for a dynamic bootstrap workflow requires the
user to set the `--apiserver-advertise-address` of each master, including the bootstrap master
itself, equal to the `controlplaneAddress` endpoint provided during kubeadm init
(the IP/DNS of the external load balancer).

By using the same advertise address for all the IP masters, `kubeadm init` can create
a unique API server serving certificate that could be shared across many masters nodes;
no changes will be required to this certificate when adding/removing master nodes.

Please note that:

- if the user is not planning to distribute the apiserver serving certificate among masters,
  kubeadm will generate a new apiserver serving certificate “almost equal” to the certificate
  created on the bootstrap master (it differs only for the domain name of the joining master)

#### Static workflow (advertise-address != `controlplaneAddress`)

In case of a static bootstrap workflow the final layout of the controlplane - the number, the
name and the IP of master nodes - is know in advance.

Given such information, the user can choose a different approach where each master has a
specific apiserver advertise address different from the `controlplaneAddress`.

Please note that:

- if the user is not planning to distribute the apiserver certificate among masters, kubeadm
  will generate a new apiserver serving certificate with the required SANS
- if the user is planning to distribute the apiserver certificate among masters, the
  operator is required to provide during `kubeadm init` the list of masters/the list of IP
  addresses for all the masters as alternative names for the API servers certificate, thus
  allowing the proper functioning of all the API server instances that will join

#### Strategies for deploying control plane components

As of today kubeadm supports two solutions for deploying control plane components:

1. Control plane deployed as static pods (current kubeadm default)
2. Self-hosted control plane (currently alpha)

The proposed solution for case 1. "Control plane deployed as static pods", assumes
that the `kubeadm join --master` flow will take care of creating required kubeconfig
files and required static pod manifests.

As stated above, supporting for Self-hosted control plane is non goal for this
proposal.

#### Strategies for distributing cluster certificates

As of today kubeadm supports two solutions for storing cluster certificates:

1. Cluster certificates stored on file system (current kubeadm default)
2. Cluster certificates stored in secrets (currently alpha)

The proposed solution for case 1. "Cluster certificates stored on file system",
requires the user/the higher level tools to execute an additional action _before_
invoking `kubeadm join --master`.

More specifically, in case of cluster with "cluster certificates stored on file
system", before invoking `kubeadm join --master`, the user/higher level tools
should copy control plane certificates from an existing master node, e.g. bootstrap master

> NB. kubeadm is limited to execute actions *only*
in the machine where it is running, so it is not possible to copy automatically
certificates from remote locations.

Then, the `kubeadm join --master` flow will take care of checking certificates
existence and conformance.

As stated above, supporting for Cluster certificates stored in secrets is a non goal
for this proposal.

#### `kubeadm upgrade` for HA clusters

Nothing in this proposal prevents implementation of `kubeadm upgrade` for HA cluster.

Further detail will be provided in a subsequent release of this KEP when all the detail
of the `v1beta1` release of kubeadm api will be available (including a proper modeling
of a multi master cluster).

## Graduation Criteria

- To create a periodic E2E test that bootstraps an HA cluster with kubeadm
  and exercise the static bootstrap workflow
- To create a periodic E2E test that bootstraps an HA cluster with kubeadm
  and exercise the dynamic bootstrap workflow
- To ensure upgradability of HA clusters (possibly with another E2E test)
- To document the kubeadm support for HA in kubernetes.io

## Implementation History

- original HA proposals [#1](https://goo.gl/QNtj5T) and [#2](https://goo.gl/C8V8PV)
- merged [Kubeadm HA design doc](https://goo.gl/QpD5h8)
- HA prototype [demo](https://goo.gl/2WLUUc) and [notes](https://goo.gl/NmTahy)
- [PR #58261](https://github.com/kubernetes/kubernetes/pull/58261) with the showcase implementation of the first release of this KEP

## Drawbacks

The kubeadm join --master workflow requires that some condition are satisfied at `kubeadm init` time,
that is use a `controlplaneAddress` and use an external etcd.

Strictly speaking, that's mean that the `kubeadm join --master` defined in this proposal supports 
a dynamic workflow _only_ in some cases.

## Alternatives

1) Execute `kubeadm init` on many nodes

The approach based on execution of `kubeadm init` on each master was considered as well,
but not chosen because it seems to have several drawbacks:

- There is no real control on parameters passed to `kubeadm init` executed on secondary masters,
  and this might lead to unpredictable inconsistent configurations.
- The init sequence for secondary master won't go through the TLS bootstrap process,
  and this might be perceived as a security concern.
- The init sequence executes a lot of steps which are un-necessary on a secondary master;
  now those steps are mostly idempotent, so basically now no harm is done by executing
  them two or three times. Nevertheless to maintain this contract in future could be complex.

Additionally, by having a separated `kubeadm join --master` workflow instead of a single `kubeadm init`
workflow we can provide better support for:

- Steps that should be done in a slightly different way on a secondary master with respect
  to the bootstrap master (e.g. updating the kubeadm-config map adding info about the new master instead
  of creating a new configMap from scratch).
- Checking that the cluster/the kubeadm-config is properly configured for multi masters
- Blocking users trying to create multi masters with configurations we don't want to support as a sig
  (e.g. HA with self-hosted control plane)