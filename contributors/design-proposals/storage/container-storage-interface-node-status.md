# Node.Status usage for CSI

***Version:*** Alpha

***Author:*** @jsafrane

## Terminology

Term | Definition
---|---
Container Storage Interface (CSI) | A specification attempting to establish an industry standard interface that Container Orchestration Systems (COs) can use to expose arbitrary storage systems to their containerized workloads.
CSI Volume Plugin | A in-tree volume plugin that acts as an adapter and enables out-of-tree, third-party CSI volume drivers to be used in Kubernetes.
CSI Volume Driver | An out-of-tree CSI compatible implementation of a volume plugin that can be used in Kubernetes through the Kubernetes CSI Volume Plugin.

## Background & Motivations

Storage backends may use different node IDs than Kubernetes - they may run on a distinct network with their own IP addresses or use separate DNS domains. Therefore when Kubernetes talks to a CSI Volume Driver, it must use node ID that is understandable to the CSI Volume Driver. A CSI Volume Driver reports ID of a node where it runs during registration of the driver to kubelet. Kubelet then collects node IDs of all CSI drivers installed on the node and currently publishes it as JSON dictionary in `csi.volume.kubernetes.io/nodeid`. In the end, each node has node IDs for all CSI drivers that run on the node.

Example of a 1.11 node running hypothetical CSI drivers for AWS and Gluster:

```yaml
apiVersion: v1
kind: Node
metadata:
  name: node1.example.com
  annotations:
    csi.volume.kubernetes.io/nodeid: { \"csi.amazon.com/ebs\": \"ip-172-31-60-59.ec2.internal\", \"csi.gluster.org\": \"172.31.60.59\" }
```

This annotation cannot be extended with additional information as CSI specification evolves. In this proposal, we want to:

  * Move this annotation into proper Node field.
  * Add topology information introduced in CSI v0.3.0.

This is part of

* https://github.com/kubernetes/features/issues/557
* https://github.com/kubernetes/community/pull/2034

This fixes:

* https://github.com/kubernetes/kubernetes/issues/66497


## API

```go
// NodeStatus is information about the current status of a node.
type NodeStatus struct {
    // ...

    // List of CSI drivers running on the node and their properties.
    CSIDrivers []CSIDriverStatus
}

type CSIDriverStatus struct {
    // CSI driver name.
    Name string

    // ID of the node from the driver point of view.
    NodeID string

    // Topology keys reported by the driver on the node.
    TopologyKeys []string
}
```

`Node.Status.CSIDrivers` will be reported by kubelet and updated when as CSI Volume Drivers register/deregister in kubelet.

`Node.Status.CSIDrivers` will be consumed by:

* [CSI attacher](https://github.com/kubernetes-csi/external-attacher/) when attaching a CSI volume to a node. CSI attacher must know node ID reported by a particular CSI driver when talking to the driver.
* [CSI provisioner](https://github.com/kubernetes-csi/external-provisioner/) when provisioning a topology-aware CSI volume, i.e. a volume that is available on given node. Details of the provisioning is available in https://github.com/kubernetes/community/pull/2034.
  In short, CSI provisioner needs to know which topology keys a CSI driver uses on a particular node in order to provision a volume that is available to the node. Again, topology keys will be reported during CSI driver registration in kubelet and should be available in `Node.Status.CSIDrivers` to the provisioner.
  CSI does not require that all nodes use the same topology keys, so we must be prepared for situation when different nodes use different set of keys and `Node.Status` is the best place where to put them.

`Node.Status.CSIDrivers` will be behind alpha feature gate "CSINodeStatus".

### Removal of `csi.volume.kubernetes.io/nodeid` annotation

The annotation was never intended as long-term supported API, we used it as temporary solution to speed up prototyping the CSI implementation. In the middle of our development were all non-alpha annotations declared as stable and supported API.

Still, we'd like to remove the annotation when "CSINodeStatus" alpha features is declared stable + 2 relases. "2" is just suggestion, it can be longer. But we'd like to get rid of the annotation eventually.

The annotation is used only by the components listed below.

## Required changes

### CSI attacher
It must look both at `Node.Status.CSIDrivers` and annotation `csi.volume.kubernetes.io/nodeid` when looking up a node ID for a CSI Volume Driver. It will use `Node.Status.CSIDrivers` when both are set.
It will keep using the annotation until `CSINodeStatus` feature is declared stable + 2 Kubernetes releases.

### CSI provisioner
All provisioner code changes are part of https://github.com/kubernetes/community/pull/2034.

CSI provisioner can't use annotation `csi.volume.kubernetes.io/nodeid` because there is no topology information there.

### kubelet

* CSI Volume Plugin has currently global (!) variable with list of CSI drivers. We need to change this into a CSIDriverManager that keeps list of CSI drivers.
* New setter callback will be added to [Kubelet.defaultNodeStatusFuncs](https://github.com/kubernetes/kubernetes/blob/8e2a444b6d81d245952cae51f293ff97843636b8/pkg/kubelet/kubelet_node_status.go#L476). This callback fills both the annotation and `Node.Status.CSIDrivers` from CSIDriverManager every time kubelet updates Node API object. The callback overwrites any changes in `Node.Status.CSIDrivers` or the annotation done by user or any other component in the cluster.
* (currently open) Remove [CSI label manager](https://github.com/kubernetes/kubernetes/blob/f2e92776bccf51f4f20e7221a1cec2c64b27badb/pkg/volume/csi/labelmanager/labelmanager.go), i.e. the code that currently sets `csi.volume.kubernetes.io/nodeid` annotation.
* (currently open) Change the node status reporting to report new/deleted CSI drivers immediately instead of every 10 seconds.

Kubelet will keep reporting the annotation until `CSINodeStatus` feature is declared stable + 2 Kubernetes releases.
