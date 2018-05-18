---
kep-number: 11
title: Switch CoreDNS to the default DNS
authors:
  - "@johnbelamaric"
  - "@rajansandeep"
owning-sig: sig-network
participating-sigs:
  - sig-cluster-lifecycle
reviewers:
  - "@bowei"
  - "@thockin"
approvers:
  - "@thockin"
editor: "@rajansandeep"
creation-date: 2018-05-18
last-updated: 2018-05-18
status: provisional
---

# Switch CoreDNS to the default DNS

## Table of Contents

* [Summary](#summary)
* [Goals](#goals)
* [Proposal](#proposal)
    * [User Cases](#use-cases)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

CoreDNS is now well-established in Kubernetes as the DNS service, with CoreDNS starting as an alpha feature from Kubernetes v1.9 to now being GA in v1.11.
After successfully implementing the road-map defined [here](https://github.com/kubernetes/features/issues/427), CoreDNS is GA in Kubernetes v1.11, which can be installed as an alternate to kube-dns in tools like kubeadm, kops, minikube and kube-up.
Following the [KEP to graduate CoreDNS to GA](https://github.com/kubernetes/community/pull/1956), the purpose of this proposal is to make CoreDNS as the default DNS for Kubernetes, replacing kube-dns.

## Goals
* Make CoreDNS the default DNS for Kubernetes for all the remaining install tools (kube-up, kops, minikube).
* Make CoreDNS available as an image in a Kubernetes repository (To Be Defined) and ensure a workflow/process to update the CoreDNS versions in the future.
  This goal is carried over from the [previous KEP](https://github.com/kubernetes/community/pull/1956), in case it cannot be completed there.

## Proposal

The proposed solution is to enable CoreDNS as the default cluster service discovery DNS for Kubernetes.
Some of the most used deployment tools will be upgraded by the CoreDNS team, in cooperation with the owners of these tools, to be able to deploy CoreDNS as default:
* kubeadm (already done for Kubernetes v1.11)
* kube-up
* minikube
* kops

For other tools, each maintainer would have to add the upgrade to CoreDNS.

### Use Cases

Use cases for CoreDNS has been well defined in the [previous KEP](https://github.com/kubernetes/community/pull/1956).
The following can be expected when CoreDNS is made the default DNS.

#### Kubeadm

* CoreDNS is already the default DNS from Kubernetes v1.11 and shall continue be the default DNS.
* In case users want to install kube-dns instead of CoreDNS, they have to set the feature-gate of CoreDNS to false. `--feature-gates=CoreDNS=false`

#### Kube-up

* CoreDNS will now become the default DNS.
* To install kube-dns in place of CoreDNS, set the environment variable `CLUSTER_DNS_CORE_DNS` to `false`.

#### Minikube

* CoreDNS to be enabled by default in the add-on manager, with kube-dns disabled by default.

#### Kops

* CoreDNS will now become the default DNS.

## Graduation Criteria

* Add CoreDNS image in a Kubernetes repository (To Be Defined) and ensure a workflow/process to update the CoreDNS versions in the future.
* Have a certain number (To Be Defined) of clusters of significant size (To Be Defined) adopting and running CoreDNS as their default DNS.

## Implementation History

* 20170912 - [Feature proposal](https://github.com/kubernetes/features/issues/427) for CoreDNS to be implemented as the default DNS in Kubernetes.
* 20171108 - Successfully released [CoreDNS as an Alpha feature-gate in Kubernetes v1.9](https://github.com/kubernetes/kubernetes/pull/52501).
* 20180226 - CoreDNS graduation to Incubation in CNCF.
* 20180305 - Support for Kube-dns configmap translation and move up [CoreDNS to Beta](https://github.com/kubernetes/kubernetes/pull/58828) for Kubernetes v1.10.
* 20180515 - CoreDNS was added as [GA and the default DNS in kubeadm](https://github.com/kubernetes/kubernetes/pull/63509) for Kubernetes v1.11.
