
# Out-of-Tree Volume Plugin FAQ for Storage Vendors

Last Updated: 1/18/2018

**What options do I have to implement a Kubernetes volume plugin out-of-tree?**

As of Kubernetes 1.9, there are two methods to implement out-of-tree volume plugins: CSI and FlexVolume.

**What is the Container Storage Interface (CSI)?**

Container Storage Interface (CSI) is a standardized mechanism for Container Orchestration Systems (COs), including Kubernetes, to expose arbitrary storage systems to containerized workloads. CSI is the next evolution of the volume plugin system in Kubernetes: one of the goals is to eventually have a CSI-compatible plugin for most existing in-tree plugins. Kubernetes 1.9 introduces an implementation of the CSI. In 1.9, CSI is alpha, and there are plans to promote it to beta in the 1.10 to 1.11 timeframe.

For more information about CSI, see:



*   http://blog.kubernetes.io/2018/01/introducing-container-storage-interface.html
*   [kubernetes-csi.github.io/docs](http://kubernetes-csi.github.io/docs)

**What are the limitations of CSI?**



*   CSI implementation in Kubernetes is not yet stable/GA (currently in alpha).

**What is FlexVolume?**

FlexVolume is an out-of-tree plugin interface that has existed in Kubernetes since version 1.2. It uses an exec-based model to interface with drivers implementing out-of-tree plugins. FlexVolume drivers must be installed on host machines. Kubernetes performs volume operations by executing pre-defined commands in the FlexVolume API against the driver on the host. FlexVolume is GA as of Kubernetes 1.8.

For more information about Flex, see:



*   https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md

**What are the limitations of FlexVolume?**



*   FlexVolume requires root access on host machine to install FlexVolume driver files.
*   FlexVolume drivers assume all volume mount dependencies, e.g. mount and filesystem tools, are available on the host OS. Installing these dependencies also require root access.

**Should I use CSI or FlexVolume?**

We suggest implementing a CSI driver if possible. CSI overcomes the limitations of FlexVolume listed above, and we plan to focus most of our development efforts on CSI in the long term. However, CSI is currently alpha, and will take several quarters to become GA. So if depending on alpha/beta software is a concern and you have a time constraint, implementing a FlexVolume driver may be a better option.

**If I already have FlexVolume driver implemented, how do I migrate to CSI?**

If Flex Volume satisfies your requirements, there is no need to migrate to CSI. The Kubernetes Storage-SIG plans to continue to support and maintain the Flex Volume API.

For those who would still like to migrate to CSI, there is an effort underway in the storage community to build a CSI adapter for FlexVolume. This will allow existing FlexVolume implementations to easily be containerized and deployed as a CSI plugin. See [this link](https://github.com/kubernetes-csi/drivers/tree/master/pkg/flexadapter) for details. However, the adapter will be a stop-gap solution, and if migration to CSI is the goal, we recommend writing a CSI driver from scratch to take full advantage of the API.
