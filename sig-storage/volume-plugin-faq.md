
# Kubernetes Volume Plugin FAQ for Storage Vendors

Last Updated: 02/08/2018

**What is Kubernetes volume plugin?**

A *Kubernetes Volume plugin* extends the Kubernetes volume interface to support a block and/or file storage system.

## In-tree vs Out-of-Tree Volume Plugins

**How do I implement a Volume plugin?**

There are three methods to implement a volume plugin:
1. In-tree volume plugin
2. Out-of-tree FlexVolume driver
3. Out-of-tree CSI driver

The Kubernetes Storage SIG, which is responsible for all volume code in the Kubernetes core repository, is no longer accepting new in-tree volume plugins. Instead, the SIG recommends storage vendors develop plugins as either Flex or CSI drivers.

**What is an in-tree vs out-of-tree volume plugin?**

Before the introduction of the Container Storage Interface (CSI) and FlexVolume, all volume plugins were *in-tree* meaning they were built, linked, compiled, and shipped with the core Kubernetes binaries and extend the core Kubernetes API. This meant that adding a new storage system to Kubernetes (a volume plugin) required checking code into the core Kubernetes code repository.

*Out-of-tree* volume plugins are developed independently of the Kubernetes code base, and are deployed (installed) on Kubernetes clusters as extensions.

**Why are new in-tree volume plugins not allowed?**

In-tree volume plugins require checking code in to the core Kubernetes repository which is undesirable for many reasons, including:
1. In-tree volume plugins make volume plugin development tightly coupled and dependent on Kubernetes releases.
2. In-tree volume plugins make Kubernetes developers/community responsible for testing and maintaining all volume plugins (which is nearly impossible).
3. In-tree volume plugins allow bugs in volume plugins to crash critical Kubernetes components, instead of just the plugin.
4. In-tree volume plugins grant volume plugins the same privileges as kubernetes components (like kubelet and kube-controller-manager).
5. In-tree volume plugins force volume plugin developers to make plugin source code public.

For these reasons, the Kubernetes Storage SIG choose to stop accepting new in-tree volume plugins, and instead requires new volume plugins be developed out-of-tree.

**What options do I now have to implement a Kubernetes volume plugin?**

As of Kubernetes 1.9, there are two out-of-tree methods to implement volume plugins: CSI and FlexVolume.

**What happens to existing in-tree volume plugins?**

One of the goals of SIG Storage is to eventually have a CSI-compatible plugin for most existing in-tree plugins and migrate the in-tree plugins to CSI.

## Container Storage Interface (CSI)

**What is the Container Storage Interface (CSI)?**

Container Storage Interface (CSI) is a standardized mechanism for Container Orchestration Systems (COs), including Kubernetes, to expose arbitrary storage systems to containerized workloads. CSI is planned to become the primary volume plugin system for Kubernetes. It was introduced in Kubernetes 1.9 as alpha. There are plans to promote it to beta in the 1.10 to 1.11 timeframe.

For more information about CSI, see:

*   http://blog.kubernetes.io/2018/01/introducing-container-storage-interface.html
*   [kubernetes-csi.github.io/docs](http://kubernetes-csi.github.io/docs)

**What are the limitations of CSI?**
*   CSI implementation in Kubernetes is not yet stable/GA (currently in alpha).

**How do I write a CSI Driver?**

For more information on how to write and deploy a CSI Driver on Kubernetes, see https://kubernetes-csi.github.io/docs/CSI-Driver.html

## FlexVolume

**What is FlexVolume?**

FlexVolume is an out-of-tree plugin interface that has existed in Kubernetes since version 1.2 (before CSI). It uses an exec-based model to interface with drivers. FlexVolume driver binaries must be installed on host machines. Kubernetes performs volume operations by executing pre-defined commands in the FlexVolume API against the driver on the host. FlexVolume is GA as of Kubernetes 1.8.

For more information about Flex, see:
*   https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md

**What are the limitations of FlexVolume?**

*   FlexVolume requires root access on host machine to install FlexVolume driver files.
*   FlexVolume drivers assume all volume mount dependencies, e.g. mount and filesystem tools, are available on the host OS. Installing these dependencies also require root access.

**Should I use CSI or FlexVolume?**

The Storage SIG suggests implementing a CSI driver if possible. CSI overcomes the limitations of FlexVolume listed above, and the SIG plans to focus most of its development efforts on CSI in the long term. However, CSI is currently alpha, and will take several quarters to become GA. So if depending on alpha/beta software is a concern and you have a time constraint, implementing a FlexVolume driver may be a better option.

**If I already have FlexVolume driver implemented, how do I migrate to CSI?**

If Flex Volume satisfies your requirements, there is no need to migrate to CSI. The Kubernetes Storage-SIG plans to continue to support and maintain the Flex Volume API.

For those who would still like to migrate to CSI, there is an effort underway in the storage community to build a CSI adapter for FlexVolume. This will allow existing FlexVolume implementations to easily be containerized and deployed as a CSI plugin. See [this link](https://github.com/kubernetes-csi/drivers/tree/master/pkg/flexadapter) for details. However, the adapter will be a stop-gap solution, and if migration to CSI is the goal, we recommend writing a CSI driver from scratch to take full advantage of the API.
