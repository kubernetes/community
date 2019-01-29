# Pod in CSI NodePublish request
Author: @jsafrane

## Goal
* Pass Pod information (pod name/namespace/UID + service account) to CSI drivers in `NodePublish` request as CSI volume attributes.

## Motivation
We'd like to move away from exec based Flex to gRPC based CSI volumes. In Flex, kubelet always passes `pod.namespace`, `pod.name`, `pod.uid` and `pod.spec.serviceAccountName` ("pod information") in every `mount` call. In Kubernetes community we've seen some Flex drivers that use pod or service account information to authorize or audit usage of a volume or generate content of the volume tailored to the pod (e.g. https://github.com/Azure/kubernetes-keyvault-flexvol).

CSI is agnostic to container orchestrators (such as Kubernetes, Mesos or CloudFoundry) and as such does not understand  concept of pods and service accounts. [Enhancement of CSI protocol](https://github.com/container-storage-interface/spec/pull/252) to pass "workload" (~pod) information from Kubernetes to CSI driver has met some resistance. 

## High-level design
We decided to pass the pod information as `NodePublishVolumeRequest.volume_attributes`.

* Kubernetes passes pod information only to CSI drivers that explicitly require that information in their [`CSIDriver` instance](https://github.com/kubernetes/community/pull/2523). These drivers are tightly coupled to Kubernetes and may not work or may require reconfiguration on other cloud orchestrators. It is expected (but not limited to) that these drivers will provide ephemeral volumes similar to Secrets or ConfigMap, extending Kubernetes secret or configuration sources.
* Kubernetes will not pass pod information to CSI drivers that don't know or don't care about pods and service accounts. It is expected (but not limited to) that these drivers will provide real persistent storage. Such CSI driver would reject a CSI call with pod information as invalid. This is current behavior of Kubernetes and it will be the default behavior.

## Detailed design

### API changes
No API changes.

### CSI enhancement
We don't need to change CSI protocol in any way. It allows kubelet to pass `pod.name`, `pod.uid` and `pod.spec.serviceAccountName` in [`NodePublish` call as `volume_attributes`]((https://github.com/container-storage-interface/spec/blob/master/spec.md#nodepublishvolume)). `NodePublish` is roughly equivalent to Flex `mount` call.

The only thing we need to do is to **define** names of the `volume_attributes` keys that CSI drivers can expect:
	*	`csi.storage.k8s.io/pod.name`: name of the pod that wants the volume.
	*	`csi.storage.k8s.io/pod.namespace`: namespace of the pod that wants the volume.
	*	`csi.storage.k8s.io/pod.uid`: uid of the pod that wants the volume.
	*	`csi.storage.k8s.io/serviceAccount.name`: name of the service account under which the pod operates. Namespace of the service account is the same as `pod.namespace`.

Note that these attribute names are very similar to [parameters we pass to flex volume plugin](https://github.com/kubernetes/kubernetes/blob/10688257e63e4d778c499ba30cddbc8c6219abe9/pkg/volume/flexvolume/driver-call.go#L55).

### Kubelet
Kubelet needs to create informer to cache `CSIDriver` instances. It passes the informer to CSI volume plugin as a new argument of [`ProbeVolumePlugins`](https://github.com/kubernetes/kubernetes/blob/43f805b7bdda7a5b491d34611f85c249a63d7f97/pkg/volume/csi/csi_plugin.go#L58).

### CSI volume plugin
In `SetUpAt()`, the CSI volume plugin checks the `CSIDriver` informer if `CSIDriver` instance exists for a particular CSI driver that handles the volume. If the instance exists and has `PodInfoRequiredOnMount` set, the volume plugin adds `csi.storage.k8s.io/*` attributes to `volume_attributes` of the CSI volume. It blindly overwrites any existing values there.

Kubelet and the volume plugin must tolerate when CRD for `CSIDriver` is not created (yet). Kubelet and CSI volume plugin falls back to original behavior, i.e. does not pass any pod information to CSI. We expect that CSI drivers will return reasonable error code instead of mounting a wrong volume. 

TODO(jsafrane): check what (shared?) informer does when it's created for non-existing CRD. Will it start working automatically when the CRD is created? Or shall we retry creation of the informer every X seconds until the CRD is created? Alternatively, we may GEt fresh `CSIDriver` from API server in `SetUpAt()`, without any informer.

## Implementation

* Alpha in 1.12 (behind `CSIPodInfo` feature gate)
* Beta in 1.13 (behind `CSIPodInfo` feature gate)
* GA 1.14?
