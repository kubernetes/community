---
kep-number: 25
title: Cross resource group nodes
authors:
  - "@feiskyer"
owning-sig: sig-azure
participating-sigs:
  - sig-azure
reviewers:
  - name: "@khenidak"
  - name: "@justaugustus"
approvers:
  - name: "@brendanburns"
editor: TBD
creation-date: 2018-08-09
last-updated: 2018-08-09
status: provisional
---

# Cross resource group nodes

## Table of Contents

<!-- TOC -->

- [Cross resource group nodes](#cross-resource-group-nodes)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
    - [Motivation](#motivation)
        - [Assumptions](#assumptions)
        - [Non-Goals](#non-goals)
    - [Design](#design)
    - [Implementation](#implementation)
        - [Cross-RG nodes](#cross-rg-nodes)
        - [On-prem nodes](#on-prem-nodes)
    - [Alternatives](#alternatives)

<!-- /TOC -->

## Summary

This KEP aims to add support for cross resource group (RG) and on-prem nodes to the Azure cloud provider.

## Motivation

Today, the Azure cloud provider only supports nodes from a specified RG (which is set in the cloud provider configuration file). For nodes in a different RG, Azure cloud provider reports `InstanceNotFound` error and thus they would be removed by controller manager. The same holds true for on-prem nodes.

With managed clusters, like [AKS](https://docs.microsoft.com/en-us/azure/aks/), there is limited access to configure nodes. There are instances where users may need to customize nodes in ways that are not possible in a managed service. This document proposes support for joining arbitrary nodes to a cluster and the required changes to make in both the Azure cloud provider and provisioned setups, which include:

- Provisioning tools should setup kubelet with required labels (e.g. via `--node-labels`)
- Azure cloud provider would fetch RG from those labels and then get node information based on that

### Assumptions

While new nodes (either from different RGs or on-prem) would be supported in this proposal, not all features would be supported for them. For example, AzureDisk will not work for on-prem nodes.

This proposal makes following assumptions for those new nodes:

- Nodes are in same region and set with required labels (as clarified in the following design part)
- Nodes will not be part of the load balancer managed by cloud provider
- Both node and container networking are properly configured
- AzureDisk is supported for Azure cross-RG nodes, but not for on-prem nodes

In addition, feature gate [ServiceNodeExclusion](https://github.com/kubernetes/kubernetes/blob/master/pkg/features/kube_features.go#L174) must also be enabled for Kubernetes cluster.

### Non-Goals

Note that provisioning the Kubernetes cluster, setting up networking and provisioning new nodes are out of this proposal scope. Those could be done by external provisioning tools (e.g. acs-engine).

## Design

Instance metadata is a general way to fetch node information for Azure, but it doesn't work if cloud-controller-manager is used (`kubelet --cloud-provider=external`). So it won't be used in this proposal. Instead, the following labels are proposed for providing required information:

- `alpha.service-controller.kubernetes.io/exclude-balancer=true`, which is used to exclude the node from load balancer. Required.
- `kubernetes.azure.com/resource-group=<rg-name>`, which provides external RG and is used to get node information. Required for cross-RG nodes.
- `kubernetes.azure.com/managed=true|false`, which indicates whether a node is on-prem or not. Required for on-prem nodes with `false` value.

When initializing nodes, these two labels should be set for kubelet by provisioning tools, e.g.

```sh
# For cross-RG nodes
kubelet --node-labels=alpha.service-controller.kubernetes.io/exclude-balancer=true,kubernetes.azure.com/resource-group=<rg-name> ...

# For on-prem nodes
kubelet --node-labels=alpha.service-controller.kubernetes.io/exclude-balancer=true,kubernetes.azure.com/managed=false ...
```

Node label `alpha.service-controller.kubernetes.io/exclude-balancer=true` has already been supported in Kubernetes, and it is controlled by feature gate `ServiceNodeExclusion`. Cluster admins should ensure the feature gate `ServiceNodeExclusion` opened when provisioning the cluster.

Note that

- Azure resource group name supports a [wider range of valid characters](https://docs.microsoft.com/en-us/azure/architecture/best-practices/naming-conventions#naming-rules-and-restrictions) than [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set). **Only kubernetes labels compatible resource group names** are supported, which must be 63 characters or less and must be empty or begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
- If the label `kubernetes.azure.com/managed` is not provided, then Azure cloud provider will assume the node to be managed.

## Implementation

### Cross-RG nodes

Cross-RG nodes should register themselves with required labels together with cloud provider:

- `--cloud-provider=azure` when using kube-controller-manager
- `--cloud-provider=external` when using cloud-controller-manager

For example,

```sh
kubelet ... \
  --cloud-provider=azure \
  --cloud-config=/etc/kubernetes/azure.json \
  --node-labels=alpha.service-controller.kubernetes.io/exclude-balancer=true,kubernetes.azure.com/resource-group=<rg-name>
```

[LoadBalancer](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L92) is not required for cross-RG nodes, hence only following features will be implemented for them:

- [Instances](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L121)
- [Zones](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L194)
- [Routes](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L169)
- [Azure managed disks](https://github.com/kubernetes/kubernetes/tree/master/pkg/volume/azure_dd)

Most operations of those features are similar with existing nodes, except the RG name. The existing nodes are using RG from cloud provider configure, while cross-RG nodes will get RG from node label `kubernetes.azure.com/resource-group=<rg-name>`.

To achieve this， [Informers](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L52-L55) will be used to get node labels and then their RGs will be cached in `nodeResourceGroups map[string]string`.

```go
type Cloud struct {
   ...
   // nodeResourceGroups is a mapping from Node's name to resource group name.
   // It will be updated by the nodeInformer.
   nodeResourceGroups map[string]string
}
```

### On-prem nodes

On-prem nodes are different from Azure nodes, all Azure coupled features (including Instances, LoadBalancer, Zones, Routes and Azure managed disks) are not supported for them. To prevent the node being deleted, Azure cloud provider will always assumes the node existing and use providerID in format `azure://<node-name>`.

On-prem nodes should register themselves with labels `alpha.service-controller.kubernetes.io/exclude-balancer=true` and `kubernetes.azure.com/managed=false`, e.g.

```sh
kubelet --node-labels=alpha.service-controller.kubernetes.io/exclude-balancer=true,kubernetes.azure.com/managed=false ...
```

Because AzureDisk is also not supported, and we don't expect Pods using AzureDisk being scheduled to on-prem nodes, a new taint `kubernetes.azure.com/managed:NoSchedule` will be added for those nodes.

To run workloads on them, nodeSelector and tolerations should be provided. For example,

```yaml
apiVersion: v1
kind: Pod
metadata:
 name: nginx
spec:
 containers:
 - image: nginx
   name: nginx
   ports:
   - containerPort: 80
     name: http
     protocol: TCP
 dnsPolicy: ClusterFirst
 nodeSelector:
   kubernetes.azure.com/resource-group: on-prem
 tolerations:
 - key: kubernetes.azure.com/managed
   effect: NoSchedule
```

## Alternatives

Annotations, additional cloud provider options and querying directly from Azure API are three alternatives ways to provide resource group information. They are not preferred because

- Kubelet doesn't support registering itself with annotations, so it requires admin to annotate the node afterward. The extra steps add complexity for cluster operations.
- Cloud provider options are not flexible compared to labels and annotations. It needs configure file updates and controller manager restarts if unknown resource groups are used for new nodes.
- Querying node information directly from Azure API is also not feasible because that would need list all resource groups, all virtual machine scale sets and all virtual machines. The operation is time consuming and easy to hit rate limits.
