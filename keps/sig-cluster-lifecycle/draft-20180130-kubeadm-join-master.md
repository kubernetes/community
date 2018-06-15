# kubeadm join --master workflow

## Metadata

```yaml
---
kep-number: draft-20180130
title: kubeadm join --master workflow
status: accepted
authors:
  - "@fabriziopandini"
owning-sig: sig-cluster-lifecycle
reviewers:
  - "@errordeveloper"
  - "@jamiehannaford"
approvers:
  - "@luxas"
  - "@timothysc"
  - "@roberthbailey"
editor:
  - "@fabriziopandini"
creation-date: 2018-01-28
last-updated: 2018-01-28
see-also:
  - KEP 0004

```

## Table of Contents

  * [kubeadm join --master workflow](#kubeadm-join---master-workflow)
    * [Metadata](#metadata)
    * [Table of Contents](#table-of-contents)
    * [Summary](#summary)
    * [Motivation](#motivation)
      * [Goals](#goals)
      * [Non-goals](#non-goals)
      * [Challenges and Open Questions](#challenges-and-open-questions)
    * [Proposal](#proposal)
      * [User Stories](#user-stories)
        * [Add a new master node](#add-a-new-master-node)
      * [Implementation Details](#implementation-details)
        * [advertise-address = IP/DNS of the external load balancer](#advertise-address--ipdns-of-the-external-load-balancer)
        * [kubeadm init --feature-gates=HighAvailability=true](#kubeadm-init---feature-gateshighavailabilitytrue)
        * [kubeadm join --master workflow](#kubeadm-join---master-workflow-1)
        * [Strategies for deploying control plane components](#strategies-for-deploying-control-plane-components)
        * [Strategies for distributing cluster certificates](#strategies-for-distributing-cluster-certificates)
    * [Graduation Criteria](#graduation-criteria)
    * [Implementation History](#implementation-history)
    * [Drawbacks](#drawbacks)
    * [Alternatives](#alternatives)

## Summary

We are extending the kubeadm distinctive `init` and `join` workflow, introducing the
capability to add more than one master node to an existing cluster by means of the
new `kubeadm join --master` option.

As a consequence, kubeadm will provide a best-practice, “fast path” for creating a
minimum viable, conformant Kubernetes cluster with one or more master nodes and
zero or more worker nodes.

## Motivation

Support for high availability is one of the most requested features for kubeadm.

Even if, as of today, there is already the possibility to create an HA cluster
using kubeadm in combination with some scripts and/or automation tools (e.g.
[this](https://kubernetes.io/docs/setup/independent/high-availability/)), this KEP was
designed with the objective to introduce an upstream simple and reliable solution for
achieving the same goal.

### Goals

* "Divide and conquer”

  This proposal - at least in its initial release - does not address all the possible
  user stories for creating an highly available Kubernetes cluster, but instead
  focuses on:

  * Defining a generic and extensible flow for bootstrapping an HA cluster, the
    `kubeadm join --master` workflow.
  * Providing a solution *only* for one, well defined user story. see
    [User Stories](#user-stories) and [Non-goals](#non-goals).

* Provide support for a dynamic bootstrap flow

  At the time a user is running `kubeadm init`, s/he/an operator might not know what
  the cluster setup will look like eventually. For instance, the user may start with
  only one master + n nodes,  and then add further master nodes with `kubeadm join --master`
  or add more worker nodes with `kubeadm join` (in any order).

* Enable higher-level tools integration

  We expect higher-level and more tailored tooling to be built on top of kubeadm,
  and ideally, using kubeadm as the basis of all deployments will make it easier
  to create conformant cluster.

  Accordingly, the `kubeadm join --master` workflow should provide support for
  the following operational practices used by higher level tools:

  * Parallel node creation

    Higher-level tools could create nodes in parallel (both masters and workers)
    for reducing the overall cluster startup time.
    `kubeadm join --master` should support natively this practice without requiring
    the implementation of any synchronization mechanics by higher-level tools.

  * Replace reconciliation strategies

    Especially in case of cloud deployments, higher-level automation tools could
    decide for any reason to replace existing nodes with new ones (instead of apply
    changes in-place to existing nodes).
    `kubeadm join --master` will support this practice by making easier to replace
    existing master nodes with new ones.

### Non-goals

* By design, kubeadm cares only about bootstrapping, not about provisioning machines.
  Likewise, installing various nice-to-have addons, like the Kubernetes Dashboard,
  monitoring solutions, and cloud-specific addons, is not in scope.

* This proposal doesn't include a solution for etcd cluster management\*.

* Nothing in this proposal should prevent users to run master nodes components
  and etcd on the same machines; however users should be aware that this will
  introduce limitations for strategies like parallel node creation and In-place
  vs. Replace reconciliation.

* Nothing in this proposal should prevent kubeadm to implement in future a
  solution for provisioning an etcd cluster based on static pods/pods.

* This proposal doesn't include a solution for API server load balancing.

* Nothing in this proposal should prevent users from choosing their preferred
  solution for API server load balancing.

* Nothing in this proposal should prevent practices that exist today.

* Nothing in this proposal should prevent user from pre-provision TLS assets
  before running `kubeadm init` or `kubeadm join --master`.

\* At the time of writing, the CoreOS recommended approach for etcd is to run 
the etcd cluster outside kubernetes (see discussion in [kubeadm office hours](https://goo.gl/fjyeqo)).

### Challenges and Open Questions

* Keep the UX simple.

  * _What are the acceptable trade-offs between the need to have a clean and simple
    UX and the complexity of the following challenges and open questions?_

* Create a cluster without knowing its final layout

  Supporting a dynamic workflow implies that some information about the cluster are
  not available at init time, like e.g. the number of master nodes, the ip of
  master nodes etc. etc.

  * _How to configure a Kubernetes cluster in order to easily adapt to future change
    of its own layout like e.g. add a master node, remove a master node?_

  * _What are the "pivotal" cluster settings that must be defined before initialising
    the cluster?_

  * _What are the mandatory conditions to be verified when executing `kubeadm init`
    to allow/not allow the execution of `kubeadm join --master` in future?_

* Kubeadm limited scope of action

  * Kubeadm binary can execute actions _only_ on the machine where it is running
    e.g. it is not possible to execute actions on other nodes, to copy files across
    nodes etc.
  * During the join workflow, kubeadm can access the cluster _only_ using identities
    with limited grants, `system:unauthenticated` or `system:node-bootstrapper`.

* Dependencies graduation

  The solution for `kubeadm join --master` will rely on a set of dependencies/other
  features which are still in the process for graduating to GA like e.g. dynamic
  kubelet configuration, self-hosting, component config.

  * _When `kubeadm join --master` should rely entirely on dependencies/features
    still under graduation vs provide compatibility with older/less convenient but
    more consolidated approaches?_

  * _Should we support  `kubeadm join --master` for cluster with a control plane
    deployed as static pods? What about cluster with a self-hosted
    control plane?_
  * _Should we support `kubeadm join --master` only  for cluster storing
    cluster-certificates on file system? What about cluster
    storing certificates in secrets?_

* Upgradability

  * How to setup an high available cluster in order to simplify the execution
    of cluster version upgrades, both manually or with the support of `kubeadm upgrade`?_

## Proposal

### User Stories

#### Add a new master node

As a kubernetes administrator, I want to run  `kubeadm join --master`  for adding
a new master node* to an existing Kubernetes cluster**, so that the cluster become
more resilient to failures of the existing master nodes (high availability).

\* A new "master node" is a new  kubernetes node with
`node-role.kubernetes.io/master=""` label and
`node-role.kubernetes.io/master:NoSchedule` taint; a new instance of control plane
components will be deployed on the new master node

> NB. In this first release of the proposal creating a new master node doesn't
trigger the creation of a new etcd member on the same machine.

\*\* In this first release of the proposal, `kubeadm join --master` could be
executed _only_ on Kubernetes cluster compliant with following conditions:

* The cluster was initialized with `kubeadm init`.
* The cluster was initialized with `--feature-gates=HighAvailability=true`.
* The cluster uses an external etcd.
* An external load balancer was provisioned and the IP/DNS of the external
  load balancer is used as advertise-address for the kube-api server.

### Implementation Details

#### advertise-address = IP/DNS of the external load balancer

There are many ways to configure an highly available cluster.

After prototyping and various discussions in
[kubeadm office hours](https://youtu.be/HcvVi8O_ZGY), it was agreed to implement
the approach that sets the `--advertise-address` equal to the IP/DNS of the
external load balancer, without assigning dedicated `--advertise-address` IPs 
for each master nodes.

By excluding the IP of master nodes, kubeadm can create a unique API server
serving certificate, and share this certificate across many masters nodes;
no changes will be required to this certificate when adding/removing master nodes.

Such properties make this approach best suited for the initial set up of
the desired `kubeadm join --master` dynamic workflow.

> Please note that in this scenario the kubernetes service will always resolve
to the IP/DNS of the external load balancer, instead of resolving to the list
of master IPs, but this fact was considered an acceptable trade-off at this stage.

> It is expected to add support also for different HA configurations in future releases
of this KEP.

#### kubeadm init --feature-gates=HighAvailability=true

When executing  `kubeadm join --master`, due to current kubeadm limitations, only
few information about the cluster/about other master nodes are available.

As a consequence, this proposal delegates to the initial `kubeadm init` - when
executed with `--feature-gates=HighAvailability=true` - all the controls about
the compliance of the cluster with the supported user story:

* The cluster uses an external etcd.
* An external load balancer is provisioned and the IP/DNS of the external load balancer is used as advertise-address.

#### kubeadm join --master workflow

The `kubeadm join --master` target workflow is an extension of the
existing `kubeadm join` flow:

1. Discovery cluster info [No changes to this step]

   Access the `cluster-info` configMap in `kube-public` namespace (or read
   the same information provided in a file).

   > This step waits for a first instance of the kube-apiserver to become ready;
   such wait cycle acts as embedded mechanism for handling the sequence
   `kubeadm init` and `kubeadm join` in case of parallel node creation.

2. In case of `join --master` [New step]

   1. Using the bootstrap token as identity, read the `kubeadm-config` configMap
      in `kube-system` namespace.

      > This requires to grant access to the above configMap for
      `system:bootstrappers` group (or to provide the same information
      provided in a file like in 1.).

   2. Check if the cluster is ready for joining a new master node:

      a. Check if the cluster was created with  `--feature-gates=HighAvailability=true`.

      > We assume that all the necessary conditions where already checked
      during `kubeadm init`:
      > * The cluster uses an external etcd.
      > * An external load balancer is provisioned and the IP/DNS of the external
    load balancer is used as advertise-address.

      b. In case of cluster certificates stored on file system, check if the
      expected certificates exists.

      > see "Strategies for distributing cluster certificates" paragraph for
      additional info about this step.

   3. Prepare the node for joining as a master node:

      a. In case of control plane deployed as static pods, create kubeconfig files
      and static pod manifests for control plane components.

      >  see "Strategies for deploying control plane components" paragraph
      for additional info about this step.

   4. Create the admin.conf kubeconfig file

3. Executes the TLS bootstrap process, including [No changes to this step]:

   1. Start kubelet using the bootstrap token as identity
   2. Request a certificate for the node - with the node identity - and retrieves
      it after it is automatically approved
   3. Restart kubelet with the node identity
   4. Eventually, apply the kubelet dynamic configuration

4. In case of `join --master` [New step]

   1. Apply master taint and label to the node.

      > This action is executed using the admin.conf identity created above;
      >
      > This action triggers the deployment of master components in case of
        self-hosted control plane

#### Strategies for deploying control plane components

As of today kubeadm supports two solutions for deploying control plane components:

1. Control plane deployed as static pods (current kubeadm default)
2. Self-hosted control plane in case of `--feature-gates=SelfHosting=true`

"Self-hosted control plane" is a solution that we expect - *in the long term* -
will become mainstream, because it simplifies both deployment and upgrade of control
plane components due to the fact that Kubernetes itself will take care of deploying
corresponding pods on nodes.

Unfortunately, at the time of writing it is unknown when this feature will graduate
to beta/GA or when this feature will become the new kubeadm default; as a consequence,
this proposal assumes that is still required to provide a solution both for case 1.
and case 2.

The proposed solution for case 1. "Control plane deployed as static pods", assumes
that the `kubeadm join --master` flow will take care of creating required kubeconfig
files and required static pod manifests.

Case 2. "Self-hosted control plane," as described above, does not requires any
additional steps to be implemented in the  `kubeadm join --master` flow.

#### Strategies for distributing cluster certificates

As of today kubeadm supports two solutions for storing cluster certificates:

1. Cluster certificates stored on file system in case of:
   * Control plane deployed as static pods (current kubeadm default)
   * Self-hosted control plane in case of `--feature-gates=SelfHosting=true`
2. Cluster certificates stored in secrets in case of:
   * Self-hosted control plane + secrets in certs in case of
     `--feature-gates=SelfHosting=true,StoreCertsInSecrets=true`

"Storing cluster certificates in secrets" is a solution that we expect - *in the
long term* - will become mainstream, because it simplifies certificates distribution
and also certificate rotation, due to the fact that Kubernetes itself will take
care of distributing certs on nodes.

Unfortunately, at the time of writing it is unknown when this feature will graduate
to beta/GA or when this feature will become the new kubeadm default; as a
consequence, this proposal assumes it is required to provide a solution for both
for case 1 and case 2.

The proposed solution for case 1. "Cluster certificates stored on file system",
requires the user/the higher level tools to execute an additional action _before_
invoking `kubeadm join --master` (NB. kubeadm is limited  to execute actions *only*
in the machine where it is running, so it is not possible to copy automatically
certificates from remote locations).

More specifically, in case of cluster with "cluster certificates stored on file
system", before invoking `kubeadm join --master`, the user/higher level tools
should copy control plane certificates from an existing node, e.g. the node
where `kubeadm init` was run, to the joining node.

Then, the  `kubeadm join --master` flow  will take care of checking certificates
existence and conformance.

Case 2. "Cluster certificates stored in secrets", as described above, does not
requires any additional steps to be implemented in the  `kubeadm join --master`
flow .

## Graduation Criteria

* To create a periodic E2E test that bootstraps an HA cluster with kubeadm
  and exercise the dynamic bootstrap workflow
* To ensure upgradability of HA clusters (possibly with another E2E test)
* To document the kubeadm support for HA in kubernetes.io

## Implementation History

* original HA proposals [#1](https://goo.gl/QNtj5T) and [#2](https://goo.gl/C8V8PV)
* merged [Kubeadm HA design doc](https://goo.gl/QpD5h8)
* HA prototype [demo](https://goo.gl/2WLUUc) and [notes](https://goo.gl/NmTahy)
* [PR #58261](https://github.com/kubernetes/kubernetes/pull/58261)

## Drawbacks

This proposal provides support for a single, well defined HA scenario.
While this is considered a sustainable approach to the complexity of HA in Kubernetes,
the limited scope of this proposal could be negatively perceived by final users.

## Alternatives

1)  Execute `kubeadm init` on many nodes

The approach based on execution of `kubeadm init` on each master was considered as well,
but not chosen because it seems to have several draw backs:

*  There is no real control on parameters passed to `kubeadm init` executed on secondary masters,
    and this can lead to unpredictable inconsistent configurations.
*  The init sequence for secondary master won't go through the TLS boostrap process,
    and this can be perceived security concern.
*  The init sequence executes a lot of steps which are un-necessary on a secondary master;
    now those steps are mostly idempotent, so basically now no harm is done by executing
    them two or three times. Nevertheless to mantain this contract in future could be complex.

2) Allow HA configurations with `--advertise-address` equal to the master ip address
(and adding the IP/DNS of the external load balancer as an additional apiServerCertSANs).

After some testing, this option was considered too complex/not
adequate for the initial set up of the desired `kubeadm join --master` dynamic workflow;
this can be better explained by looking at two implementation based on this option:

* [kubernetes the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way)
   uses the IP address of all master nodes for creating a new API server
   serving certificate before bootstrapping the cluster, but this approach
   can't be used if considering the desired dynamic workflow.

* [Creating HA cluster with kubeadm](https://kubernetes.io/docs/setup/independent/high-availability/)
  uses a different API server serving certificate for each master, and this
  could increases the complexity of the first implementation because:
  * the `kubeadm join --master` flow has to generate different certificates for
    each master node.
  * self-hosting control plane, should be adapted to mount different certificates
    for each master.
  * bootstrap check pointing should be designed to checkpoint a different
    set of certificates for each master.
  * upgrades should be adapted to consider master specific settings
