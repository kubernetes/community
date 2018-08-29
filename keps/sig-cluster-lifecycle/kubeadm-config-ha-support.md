---
kep-number: TBD
title: Augment Kubeadm Config to Enable Upgrades of HA Clusters
authors:
  - "@mattkelly"
owning-sig: sig-cluster-lifecycle
reviewers:
  - "@timothysc"
  - "@fabriziopandini"
approvers:
  - TBD
editor:
  - "@mattkelly"
creation-date: 2018-03-08
last-updated: 2018-04-09
status: provisional
see-also:
  - [Creating HA clusters with kubeadm](https://kubernetes.io/docs/setup/independent/high-availability/)
  - [KEP kubeadm join --master workflow](https://github.com/kubernetes/community/pull/1707) (in progress)
  - [Upgrading kubeadm HA clusters from 1.9.x to 1.9.y](https://kubernetes.io/docs/tasks/administer-cluster/upgrade-downgrade/kubeadm-upgrade-ha/)
---

# Augment Kubeadm Config to Enable Upgrades of HA Clusters

## Table of Contents

* [Augment Kubeadm Config to Enable Upgrades of HA Clusters](#augment-kubeadm-config-to-enable-upgrades-of-ha-clusters)
* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
   * [Goals](#goals)
   * [Non-Goals](#non-goals)
   * [Challenges and Open Questions](#challenges-and-open-questions)
* [Proposal](#proposal)
   * [Implementation Details](#implementation-details)
      * [Background](#background)
      * [Adding Additional Master-Specific ConfigMaps](#adding-additional-master-specific-configmaps)
         * [Key Design Considerations and Benefits](#key-design-considerations-and-benefits)
            * [Parallel Node Creation](#parallel-node-creation)
            * [Guaranteed Consistent kubeadm-config](#guaranteed-consistent-kubeadm-config)
   * [Risks and Mitigations](#risks-and-mitigations)
      * [Migrating Existing Clusters](#migrating-existing-clusters)
      * [More Complex User Experience for Overriding Configuration](#more-complex-user-experience-for-overriding-configuration)
* [Implementation History](#implementation-history)
* [Drawbacks](#drawbacks)
* [Alternatives](#alternatives)

## Summary

One of the first steps of the upgrade process is to retrieve the `kubeadm-config` ConfigMap that is stored at `kubeadm init` time in order to get any relevant cluster configuration.
The current `kubeadm-config` ConfigMap was designed to support upgrades of clusters with a single master.
As we move towards supporting High Availability (HA) natively in kubeadm, the persistent configuration must be enhanced in order to store information unique to each master node.
For example, master node-specific information is required in order for kubeadm to identify which control plane static pods belong to the current node during an upgrade (the static pods are identifiable by the `nodeName` field embedded in the pod name).
If a kubeadm HA cluster is created today using the [workarounds recommended by the community](https://github.com/kubernetes/kubeadm/issues/546), then each subsequent `kubeadm init` on a master node will overwrite the master node-specific information from the previous `init`, which is not desired.

This KEP outlines a possible solution for adding and retrieving master node-specific information through the use of additional kubeadm ConfigMaps.

## Motivation

Kubeadm is driving towards natively supporting highly available clusters.
As part of HA support, a clean upgrade path is required.
The purpose of this KEP is to introduce support for multiple masters in the kubeadm configuration that is stored in-cluster in order to enable that clean upgrade path.

### Goals

Enable `kubeadm upgrade` of highly available clusters by augmenting the existing persistent kubeadm configuration.

### Non-Goals

This proposal does not aim to solve the entire problem of upgrading HA clusters.
This KEP specifically tackles the persistent configuration problem so that the information required at upgrade time is available.

### Challenges and Open Questions

The final implementation of this KEP will require deciding exactly what "master node-specific information" means.
Currently, the `nodeName` and `advertiseAddress` of the master are two entries known to be node-specific.
It may be possible that additional config entries could be split out into the node-specific area(s) of the config.
This could result in asymmetric configuration across the masters, which may or may not be something that we wish to support.

## Proposal

### Implementation Details

#### Background

Currently, the `kubeadm-config` ConfigMap in the `kube-system` namespace serves as the single source of truth for how kubeadm has been used to create and modify a cluster.
Because kubeadm is not a process that runs on the cluster (it is only run to perform operations, e.g. `init` and `upgrade`), this config is not modified during normal operation.
In the non-HA case today, it is guaranteed to be an accurate representation of the kubeadm configuration.

If kubeadm is used to create an HA cluster today, e.g. using the workarounds described in [kubeadm #546](https://github.com/kubernetes/kubeadm/issues/546) and/or @mbert's [document](https://docs.google.com/document/d/1rEMFuHo3rBJfFapKBInjCqm2d7xGkXzh0FpFO0cRuqg), then the `kubeadm-config` ConfigMap will be an accurate representation except for any master node-specific information.
As explained in [Challenges and Open Questions](#challenges-and-open-questions), such node-specific information is not yet well-defined but minimally consists of the master's `nodeName` and `advertiseAddress`.
The `nodeName` in `kubeadm-config` will correspond to the last master that happened to write to the ConfigMap.
In the case of parallel node creation, this may not be well-defined.
When `kubeadm upgrade` is run on a master and this `nodeName` is fetched, it may be incorrect and the upgrade process will fail.

#### Adding Additional Master-Specific ConfigMaps

The proposed solution is to add additional kubeadm ConfigMaps that are specific to each master (one ConfigMap for each master).
Each master-specific ConfigMap will be created as part of the `kubeadm init` process for the initial master and as part of the to-be-implemented [`kubeadm join --master` process](https://github.com/kubernetes/community/pull/1707) for additional masters.
Any master-specific information in the main `kubeadm-config` ConfigMap will be removed.
Each master-specific ConfigMap can be automatically deleted via garbage collection (see [below](#guaranteed-consistent-kubeadm-config) for details).

The names of these new ConfigMaps will be `kubeadm-config-<machine_UID>` where `machine_UID` is an identifier that is guaranteed to be unique for each node in the cluster.
There is a precedent for using such a `machine_UID`, and in fact kubeadm already has a [prerequisite](https://kubernetes.io/docs/setup/independent/install-kubeadm/#verify-the-mac-address-and-product_uuid-are-unique-for-every-node) that such machine identifiers be unique for every node.
For the purpose of this KEP, let us assume that `machine_UID` is the full `product_uuid` of the machine that the master node is running on.

Kubeadm operations such as upgrade that require master-specific information should now also retrieve the corresponding ConfigMap for their node.
This master-specific configuration will be explicitly provided to any functions that require it instead of e.g. merging everything into one configuration.

##### Key Design Considerations and Benefits

There are a few key benefits to the approach of adding additional ConfigMaps over an approach which would augment the existing `kubeadm-config` with master-specific information:

###### Parallel Node Creation

Node creation in parallel is a valid use-case that works today.
By adding additional ConfigMaps instead of requiring each master to modify the existing `kubeadm-config`, we avoid the need to lock on that ConfigMap.

###### Guaranteed Consistent kubeadm-config

This approach allows us to continue to guarantee that the main `kubeadm-config` is consistent with the actual cluster configuration.
If we put master-specific information into `kubeadm-config` itself, then we would require either a yet-to-be-defined `kubeadm leave` workflow or active reconciliation of `kubeadm-config` in order to ensure accurateness.
This may not be critical, but it is a consideration.

With this proposal, if a node unexpectedly leaves a cluster, then at worst a dangling ConfigMap will be left in the cluster.
For the case where a node is explicitly deleted, we can leverage garbage collection to automatically delete the master-specific ConfigMap by listing the node as an `ownerReference` when the ConfigMap is created.

### Risks and Mitigations

#### Migrating Existing Clusters

There will be situations in which a kubeadm operation (e.g. upgrade) that requires the new master-specific ConfigMap is run and finds that the expected ConfigMap does not exist.
For example, this will happen for users who are upgrading HA clusters that were created using the aforementioned workarounds required before kubeadm HA support is available.
We can automate the creation of missing node-specific ConfigMaps in the following manner when a `kubeadm upgrade` (or other operation requiring it) is performed:

1. Determine which node kubeadm is currently running on by listing all master nodes in the cluster and looking at the existing `product_uuid` field
2. Get the `nodeName` from this node's metadata
3. Compare the current `nodeName` to the `nodeName` in `kubeadm-config`
4. If they are equal, update `kubeadm-config` to remove node-specific information (to match the new `kubeadm-config` specification) 
5. Create the `kubeadm-config-<machine_UID>` ConfigMap for this node
6. Continue the upgrade process

#### More Complex User Experience for Overriding Configuration

Currently, users may override configuration items by providing a configuration file when running kubeadm.
The existence of additional, disjoint ConfigMaps may make the user experience more complex for overriding configuration.
One possibility for mitigating this would be to keep the `kubeadm-config` specification the same as it is today instead of removing fields.
This would allow a user to specify any node-specific information in the same configuration file instead of having to provide multiple files.
Placing this information in the appropriate node-specific ConfigMap would be an implementation detail not requiring any impact to user experience.

## Implementation History

- [Issue #546: Workarounds for the time before kubeadm HA becomes available](https://github.com/kubernetes/kubeadm/issues/546)
- [Adding HA to kubeadm-deployed clusters](https://docs.google.com/document/d/1rEMFuHo3rBJfFapKBInjCqm2d7xGkXzh0FpFO0cRuqg)
- [Issue #706: Make kubeadm upgrade HA ready](https://github.com/kubernetes/kubeadm/issues/706)

## Drawbacks

This KEP introduces additional ConfigMaps for kubeadm to use (one for each master).

## Alternatives

An alternative approach would be to augment the existing `kubeadm-config` ConfigMap with master-specific information.
The advantages over this approach are detailed in the [Proposal](#proposal) section.
