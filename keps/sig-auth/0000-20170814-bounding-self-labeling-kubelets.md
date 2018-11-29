---
kep-number: 0
title: Bounding Self-Labeling Kubelets
authors:
  - "@mikedanese"
  - "@liggitt"
owning-sig: sig-auth
participating-sigs:
  - sig-node
  - sig-storage
reviewers:
  - "@saad-ali"
  - "@tallclair"
approvers:
  - "@thockin"
  - "@smarterclayton"
creation-date: 2017-08-14
last-updated: 2018-10-31
status: implementable
---

# Bounding Self-Labeling Kubelets

## Motivation

Today the node client has total authority over its own Node labels.
This ability is incredibly useful for the node auto-registration flow.
The kubelet reports a set of well-known labels, as well as additional
labels specified on the command line with `--node-labels`.

While this distributed method of registration is convenient and expedient, it
has two problems that a centralized approach would not have. Minorly, it makes
management difficult. Instead of configuring labels in a centralized
place, we must configure `N` kubelet command lines. More significantly, the
approach greatly compromises security. Below are two straightforward escalations
on an initially compromised node that exhibit the attack vector.

### Capturing Dedicated Workloads

Suppose company `foo` needs to run an application that deals with PII on
dedicated nodes to comply with government regulation. A common mechanism for
implementing dedicated nodes in Kubernetes today is to set a label or taint
(e.g. `foo/dedicated=customer-info-app`) on the node and to select these
dedicated nodes in the workload controller running `customer-info-app`.

Since the nodes self reports labels upon registration, an intruder can easily
register a compromised node with label `foo/dedicated=customer-info-app`. The
scheduler will then bind `customer-info-app` to the compromised node potentially
giving the intruder easy access to the PII.

This attack also extends to secrets. Suppose company `foo` runs their outward
facing nginx on dedicated nodes to reduce exposure to the company's publicly
trusted server certificates. They use the secret mechanism to distribute the
serving certificate key. An intruder captures the dedicated nginx workload in
the same way and can now use the node certificate to read the company's serving
certificate key.

## Proposal

1. Modify the `NodeRestriction` admission plugin to prevent Kubelets from self-setting labels
within the `k8s.io` and `kubernetes.io` namespaces *except for these specifically allowed labels/prefixes*:

    ```
    kubernetes.io/hostname
    kubernetes.io/instance-type
    kubernetes.io/os
    kubernetes.io/arch

    beta.kubernetes.io/instance-type
    beta.kubernetes.io/os
    beta.kubernetes.io/arch

    failure-domain.beta.kubernetes.io/zone
    failure-domain.beta.kubernetes.io/region

    failure-domain.kubernetes.io/zone
    failure-domain.kubernetes.io/region

    [*.]kubelet.kubernetes.io/*
    [*.]node.kubernetes.io/*
    ```

2. Reserve and document the `node-restriction.kubernetes.io/*` label prefix for cluster administrators
that want to label their `Node` objects centrally for isolation purposes.

    > The `node-restriction.kubernetes.io/*` label prefix is reserved for cluster administrators
    > to isolate nodes. These labels cannot be self-set by kubelets when the `NodeRestriction`
    > admission plugin is enabled.

This accomplishes the following goals:

- continues allowing people to use arbitrary labels under their own namespaces any way they wish
- supports legacy labels kubelets are already adding
- provides a place under the `kubernetes.io` label namespace for node isolation labeling
- provide a place under the `kubernetes.io` label namespace for kubelets to self-label with kubelet and node-specific labels

## Implementation Timeline

v1.13:

* Kubelet deprecates setting `kubernetes.io` or `k8s.io` labels via `--node-labels`, 
other than the specifically allowed labels/prefixes described above,
and warns when invoked with `kubernetes.io` or `k8s.io` labels outside that set.
* NodeRestriction admission prevents kubelets from adding/removing/modifying `[*.]node-restriction.kubernetes.io/*` labels on Node *create* and *update*
* NodeRestriction admission prevents kubelets from adding/removing/modifying `kubernetes.io` or `k8s.io`
labels other than the specifically allowed labels/prefixes described above on Node *update* only

v1.15:

* Kubelet removes the ability to set `kubernetes.io` or `k8s.io` labels via `--node-labels`
other than the specifically allowed labels/prefixes described above (deprecation period
of 6 months for CLI elements of admin-facing components is complete)

v1.17:

* NodeRestriction admission prevents kubelets from adding/removing/modifying `kubernetes.io` or `k8s.io`
labels other than the specifically allowed labels/prefixes described above on Node *update* and *create*
(oldest supported kubelet running against a v1.17 apiserver is v1.15)

## Alternatives Considered

### File or flag-based configuration of the apiserver to allow specifying allowed labels

* A fixed set of labels and label prefixes is simpler to reason about, and makes every cluster behave consistently
* File-based config isn't easily inspectable to be able to verify enforced labels
* File-based config isn't easily kept in sync in HA apiserver setups

### API-based configuration of the apiserver to allow specifying allowed labels

* A fixed set of labels and label prefixes is simpler to reason about, and makes every cluster behave consistently
* An API object that controls the allowed labels is a potential escalation path for a compromised node

### Allow kubelets to add any labels they wish, and add NoSchedule taints if disallowed labels are added

* To be robust, this approach would also likely involve a controller to automatically inspect labels and remove the NoSchedule taint. This seemed overly complex. Additionally, it was difficult to come up with a tainting scheme that preserved information about which labels were the cause.

### Forbid all labels regardless of namespace except for a specifically allowed set

* This was much more disruptive to existing usage of `--node-labels`.
* This was much more difficult to integrate with other systems allowing arbitrary topology labels like CSI.
* This placed restrictions on how labels outside the `kubernetes.io` and `k8s.io` label namespaces could be used, which didn't seem proper.
