---
kep-number: 18
title: Azure Availability Zones
authors:
  - "@feiskyer"
owning-sig: sig-azure
participating-sigs:
  - sig-azure
  - sig-storage
reviewers:
  - name: "@khenidak"
  - name: "@colemickens"
approvers:
  - name: "@brendanburns"
editor:
  - "@feiskyer"
creation-date: 2018-07-11
last-updated: 2018-09-29
status: implementable
---

# Azure Availability Zones

## Table of Contents

- [Azure Availability Zones](#azure-availability-zones)
  - [Summary](#summary)
  - [Scopes and Non-scopes](#scopes-and-non-scopes)
    - [Scopes](#scopes)
    - [Non-scopes](#non-scopes)
  - [AZ label format](#az-label-format)
  - [Cloud provider options](#cloud-provider-options)
  - [Node registration](#node-registration)
    - [Get by instance metadata](#get-by-instance-metadata)
    - [Get by Go SDK](#get-by-go-sdk)
  - [LoadBalancer and PublicIP](#loadbalancer-and-publicip)
  - [AzureDisk](#azuredisk)
    - [PVLabeler](#pvlabeler)
    - [PersistentVolumeLabel](#persistentvolumelabel)
    - [StorageClass](#storageclass)
  - [Appendix](#appendix)

## Summary

This proposal aims to add [Azure Availability Zones (AZ)](https://azure.microsoft.com/en-us/global-infrastructure/availability-zones/) support to Kubernetes.

## Scopes and Non-scopes

### Scopes

The proposal includes required changes to support availability zones for various functions in Azure cloud provider and AzureDisk volumes:

- Detect availability zones automatically when registering new nodes (by kubelet or node controller) and node's label `failure-domain.beta.kubernetes.io/zone` will be replaced with AZ instead of fault domain
- LoadBalancer and PublicIP will be provisioned with zone redundant
- `GetLabelsForVolume` interface will be implemented for Azure managed disks so that PV label controller in cloud-controller-manager can appropriately add `Labels` and `NodeAffinity` to the Azure managed disk PVs. Additionally, `PersistentVolumeLabel` admission controller will be enhanced to achieve the same for Azure managed disks.
- Azure Disk's `Provision()` function will be enhanced to take into account the zone of the node as well as `allowedTopologies` when determining the zone to create a disk in.

> Note that unlike most cases, fault domain and availability zones mean different on Azure:
>
> - A Fault Domain (FD) is essentially a rack of servers. It consumes subsystems like network, power, cooling etc.
> - Availability Zones are unique physical locations within an Azure region. Each zone is made up of one or more data centers equipped with independent power, cooling, and networking.
>
> An Availability Zone in an Azure region is a combination of a fault domain and an update domain (Same like FD, but for updates. When upgrading a deployment, it is carried out one update domain at a time). For example, if you create three or more VMs across three zones in an Azure region, your VMs are effectively distributed across three fault domains and three update domains.

### Non-scopes

Provisioning Kubernetes masters and nodes with availability zone support is not included in this proposal. It should be done in the provisioning tools (e.g. acs-engine). Azure cloud provider will auto-detect the node's availability zone if `availabilityZones` option is configured for the Azure cloud provider.

## AZ label format

Currently, Azure nodes are registered with label `failure-domain.beta.kubernetes.io/zone=faultDomain`.

The format of fault domain is numbers (e.g. `1` or `2`), which is in same format with AZ (e.g. `1` or `3`). If AZ is using same format with faultDomain, then there'll be scheduler issues for clusters with both AZ and non-AZ nodes. So AZ will use a different format in kubernetes: `<region>-<AZ>`, e.g. `centralus-1`.

The AZ label will be applied in multiple Kubernetes resources, e.g.

- Nodes
- AzureDisk PersistentVolumes
- AzureDisk StorageClass

## Cloud provider options

Because only standard load balancer is supported with AZ, it is a prerequisite to enable AZ for the cluster.

Standard load balancer has been added in Kubernetes v1.11, related options include:

| Option                      | Default | **AZ Value**  | Releases | Notes                                 |
| --------------------------- | ------- | ------------- | -------- | ------------------------------------- |
| loadBalancerSku             | basic   | **standard**  | v1.11    | Enable standard LB                    |
| excludeMasterFromStandardLB | true    | true or false | v1.11    | Exclude master nodes from LB backends |

These options should be configured in Azure cloud provider configure file (e.g. `/etc/kubernetes/azure.json`):

```json
{
    ...,
    "loadBalancerSku": "standard",
    "excludeMasterFromStandardLB": true
}
```

Note that with standard SKU LoadBalancer, `primaryAvailabitySetName` and `primaryScaleSetName` is not required because all available nodes (with configurable masters via `excludeMasterFromStandardLB`) are added to LoadBalancer backend pools.

## Node registration

When registering new nodes, kubelet (with build in cloud provider) or node controller (with external cloud provider) automatically adds labels to them with region and zone information:

- Region: `failure-domain.beta.kubernetes.io/region=centralus`
- Zone: `failure-domain.beta.kubernetes.io/zone=centralus-1`

```sh
$ kubectl get nodes --show-labels
NAME                STATUS    AGE   VERSION    LABELS
kubernetes-node12   Ready     6m    v1.11      failure-domain.beta.kubernetes.io/region=centralus,failure-domain.beta.kubernetes.io/zone=centralus-1,...
```

Azure cloud providers sets fault domain for label `failure-domain.beta.kubernetes.io/zone` today. With AZ enabled, we should set the node's availability zone instead. To keep backward compatibility and distinguishing from fault domain, `<region>-<AZ>` is used here.

The node's zone could get by ARM API or instance metadata. This will be added in  `GetZoneByProviderID()` and `GetZoneByNodeName()`.

### Get by instance metadata

This method is used in kube-controller-manager.

```sh
# Instance metadata API should be upgraded to 2017-12-01.
$ curl -H Metadata:true "http://169.254.169.254/metadata/instance/compute/zone?api-version=2017-12-01&format=text"
2
```

### Get by Go SDK

This method is used in cloud-controller-manager.

No `zones` property is included in `VirtualMachineScaleSetVM` yet in Azure Go SDK (including latest 2018-04-01 compute API).

We need to ask Azure Go SDK to add `zones` for `VirtualMachineScaleSetVM`. Opened the issue https://github.com/Azure/azure-sdk-for-go/issues/2183 for tracking it.

> Note: there's already `zones` property in `VirtualMachineScaleSet`, `VirtualMachine` and `Disk`.

## LoadBalancer and PublicIP

LoadBalancer with standard SKU will be created and all available nodes (including VirtualMachines and VirtualMachineScaleSetVms, together with optional masters configured via excludeMasterFromStandardLB) are added to LoadBalancer backend pools.

PublicIPs will also be created with standard SKU, and they are zone redundant by default.

Note that zonal PublicIPs are not supported. We may add this easily if there’re clear use-cases in the future.

## AzureDisk

When Azure managed disks are created, the `PersistentVolumeLabel` admission controller or PV label controller automatically adds zone labels and node affinity to them. The scheduler (via `VolumeZonePredicate` or `PV.NodeAffinity`) will then ensure that pods that claim a given volume are only placed into the same zone as that volume, as volumes cannot be attached across zones. In addition, admission controller

Note that

- Only managed disks are supported. Blob disks don't support availability zones on Azure.
- Node affinity is enabled by feature gate `VolumeScheduling`.

### PVLabeler interface

To setup AzureDisk's zone label correctly (required by cloud-controller-manager's PersistentVolumeLabelController), Azure cloud provider's [PVLabeler](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L212) interface should be implemented:

```go
// PVLabeler is an abstract, pluggable interface for fetching labels for volumes
type PVLabeler interface {
    GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error)
}
```

It should return the region and zone for the AzureDisk, e.g.

- `failure-domain.beta.kubernetes.io/region=centralus`
- `failure-domain.beta.kubernetes.io/zone=centralus-1`

so that the PV will be created with labels:

```sh
$ kubectl get pv --show-labels
NAME           CAPACITY   ACCESSMODES   STATUS    CLAIM            REASON    AGE       LABELS
pv-managed-abc 5Gi        RWO           Bound     default/claim1             46s       failure-domain.beta.kubernetes.io/region=centralus,failure-domain.beta.kubernetes.io/zone=centralus-1
```

### PersistentVolumeLabel admission controller

Cloud provider's `PVLabeler` interface is only applied when cloud-controller-manager is used. For build in Azure cloud provider, [PersistentVolumeLabel](https://github.com/kubernetes/kubernetes/blob/master/plugin/pkg/admission/storage/persistentvolume/label/admission.go) admission controller should also updated with AzureDisk support, so that new PVs could also be applied with above labels.

```go
func (l *persistentVolumeLabel) Admit(a admission.Attributes) (err error) {
    ...
    if volume.Spec.AzureDisk != nil {
        labels, err := l.findAzureDiskLabels(volume)
        if err != nil {
            return admission.NewForbidden(a, fmt.Errorf("error querying AzureDisk volume %s: %v", volume.Spec.AzureDisk.DiskName, err))
        }
        volumeLabels = labels
    }
    ...
}
```

> Note: the PersistentVolumeLabel admission controller will be deprecated, and cloud-controller-manager is preferred after its GA (probably v1.13 or v1.14).

### StorageClass

Note that the above interfaces are only applied to AzureDisk persistent volumes, not StorageClass. For AzureDisk StorageClass, we should add a few new options for zone-aware and [topology-aware](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/volume-topology-scheduling.md) provisioning. The following three new options will be added in AzureDisk StorageClass:

- `zoned`: indicates whether new disks are provisioned with AZ. Default is `true`.
- `zone` and `zones`: indicates which zones should be used to provision new disks (zone-aware provisioning). Only can be set if `zoned` is not false and `allowedTopologies` is not set.
- `allowedTopologies`: indicates which topologies are allowed for [topology-aware](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/volume-topology-scheduling.md) provisioning. Only can be set if `zoned` is not false and `zone`/`zones` are not set.

An example of zone-aware provisioning storage class is:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
  name: managed-premium
parameters:
  kind: Managed
  storageaccounttype: Premium_LRS
  # only one of zone and zones are allowed
  zone: "centralus-1"
  # zones: "centralus-1,centralus-2,centralus-3"
provisioner: kubernetes.io/azure-disk
```

Another example of topology-aware provisioning storage class is:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
  name: managed-premium
parameters:
  kind: Managed
  storageaccounttype: Premium_LRS
provisioner: kubernetes.io/azure-disk
allowedTopologies:
- matchLabelExpressions:
  - key: failure-domain.beta.kubernetes.io/zone
    values:
    - centralus-1
    - centralus-2
```

AzureDisk can only be created with one specific zone, so if multiple zones are specified in the storage class, then new disks will be provisioned with zone chosen by following rules:

- If `DynamicProvisioningScheduling` is enabled and `VolumeBindingMode: WaitForFirstConsumer` is specified in the storage class, zone of the disk should be set to the zone of the node passed to `Provision()`. Specifying zone/zones in storage class should be considered an error in this scenario.
- If `DynamicProvisioningScheduling` is enabled and `VolumeBindingMode: WaitForFirstConsumer` is not specified in StorageClass, zone of disk should be chosen from `allowedTopologies` or zones depending on which is specified. Specifying both `allowedTopologies` and `zones` should lead to error.
- If `DynamicProvisioningScheduling` is disabled and `zones` are specified, then the zone maybe arbitrarily chosen as specified by arbitrarily choosing from the zones specified in the storage class.
- If `DynamicProvisioningScheduling` is disabled and no zones are specified and `zoned` is `true`, then new disks will be provisioned with zone chosen by round-robin across all active zones, which means
  - If there are no zoned nodes, then an `no zoned nodes` error will be reported
  - Zoned AzureDisk will only be provisioned when there are zoned nodes
  - If there are multiple zones, then those zones are chosen by round-robin

Note that

- active zones means there're nodes in that zone.
- there are risks if the cluster is running with both zoned and non-zoned nodes. In such case, zoned AzureDisk can't be attached to non-zoned nodes. So
  - new pods with zoned AzureDisks are always scheduled to zoned nodes
  - old pods using non-zoned AzureDisks can't be scheduled to zoned nodes

So if users are planning to migrate workloads to zoned nodes, old AzureDisks should be recreated (probably backup first and restore to the new one).

## Implementation History

- [kubernetes#66242](https://github.com/kubernetes/kubernetes/pull/66242): Adds initial availability zones support for Azure nodes.
- [kubernetes#66553](https://github.com/kubernetes/kubernetes/pull/66553): Adds avaialability zones support for Azure managed disks.
- [kubernetes#67121](https://github.com/kubernetes/kubernetes/pull/67121): Adds DynamicProvisioningScheduling and VolumeScheduling support for Azure managed disks.
- [cloud-provider-azure#57](https://github.com/kubernetes/cloud-provider-azure/pull/57): Adds documentation for Azure availability zones.

## Appendix

Kubernetes will automatically spread the pods in a replication controller or service across nodes in a single-zone cluster (to reduce the impact of failures).

With multiple-zone clusters, this spreading behavior is extended across zones (to reduce the impact of zone failures.) (This is achieved via `SelectorSpreadPriority`). This is a best-effort placement, and so if the zones in your cluster are heterogeneous (e.g. different numbers of nodes, different types of nodes, or different pod resource requirements), this might prevent perfectly even spreading of your pods across zones. If desired, you can use homogeneous zones (same number and types of nodes) to reduce the probability of unequal spreading.

There's also some [limitations of availability zones of various Kubernetes functions](https://kubernetes.io/docs/setup/multiple-zones/#limitations), e.g.

- No zone-aware network routing
- Volume zone-affinity will only work with a `PersistentVolume`, and will not work if you directly specify an AzureDisk volume in the pod spec.
- Clusters cannot span clouds or regions (this functionality will require full federation support).
- StatefulSet volume zone spreading when using dynamic provisioning is currently not compatible with pod affinity or anti-affinity policies.
- If the name of the StatefulSet contains dashes (“-”), volume zone spreading may not provide a uniform distribution of storage across zones.
- When specifying multiple PVCs in a Deployment or Pod spec, the StorageClass needs to be configured for a specific, single zone, or the PVs need to be statically provisioned in a specific zone. Another workaround is to use a StatefulSet, which will ensure that all the volumes for a replica are provisioned in the same zone.

See more at [running Kubernetes in multiple zones](https://kubernetes.io/docs/setup/multiple-zones/).
