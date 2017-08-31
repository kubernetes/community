# Deploying a default StorageClass during installation

## Goal

Usual Kubernetes installation tools should deploy a default StorageClass
where it makes sense.

"*Usual installation tools*" are:

* cluster/kube-up.sh
* kops
* kubeadm

Other "installation tools" can (and should) deploy default StorageClass
following easy steps described in this document, however we won't touch them
during implementation of this proposal.

"*Where it makes sense*" are:

* AWS
* Azure
* GCE
* Photon
* OpenStack
* vSphere

Explicitly, there is no default storage class on bare metal.

## Motivation

In Kubernetes 1.5, we had "alpha" dynamic provisioning on aforementioned cloud
platforms. In 1.6 we want to deprecate this alpha provisioning. In order to keep
the same user experience, we need a default StorageClass instance that would
provision volumes for PVCs that do not request any special class. As
consequence, this default StorageClass would provision volumes for PVCs with
"alpha" provisioning annotation - this annotation would be ignored in 1.6 and
default storage class would be assumed.

## Design

1. Kubernetes will ship yaml files for default StorageClasses for each platform
   as `cluster/addons/storage-class/<platform>/default.yaml` and all these
   default classes will distributed together with all other addons in
   `kubernetes.tar.gz`.

2. An installation tool will discover on which platform it runs and installs
   appropriate yaml file into usual directory for addon manager (typically
   `/etc/kubernetes/addons/storage-class/default.yaml`).

3. Addon manager will deploy the storage class into installed cluster in usual
   way. We need to update addon manager not to overwrite any existing object
   in case cluster admin has manually disabled this default storage class!

## Implementation

* AWS, GCE and OpenStack has a default StorageClass in
  `cluster/addons/storage-class/<platform>/` - already done in 1.5

* We need a default StorageClass for vSphere, Azure and Photon in `cluster/addons/storage-class/<platform>`

* cluster/kube-up.sh scripts need to be updated to install the storage class on appropriate platforms
  * Already done on GCE, AWS and OpenStack.

* kops needs to be updated to install the storage class on appropriate platforms
  * already done for kops on AWS and kops does not support other platforms yet.

* kubeadm needs to be updated to install the storage class on appropriate platforms (if it is cloud-provider aware)

* addon manager fix: https://github.com/kubernetes/kubernetes/issues/39561
