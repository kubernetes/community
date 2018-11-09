Kubernetes PVC Namespace Transfer Proposal
==========================================

**Authors:** [John Griffith](https://github.com/j-griffith)

## Background

There are a number of use cases where having the ability to transfer a PVC from one namespace to another is a valuable use case.  The goal would be to allow Volumes to be populated with data in one namespace (this could be a clone, a volume from snapshot or any populated data set) and then transferred to a user in another namespace.  While in some cases it might be possible to recreate the data on new volume, this isn't always feasible, and in cases where the data may be large, or controlled in some way the ability to transfer is a valuable feature.  It's also useful in multi namespace Cluster to be able to duplicate and migrate data between namespaces like development and test or in some cases blue-green deployments.

## Objectives

Provide a mechansim to allow users to "give" a PVC to a user in another namespace based on mutual acceptance.  A transfer should be initiated by updating an existing PVC with a transfer request annotation, and accepted by the intended destination namespace by the recepient requesting a PVC using the standard mechanism and specifying the correct PVC Name and transfer annotation as per the transfer request.

* Provide a Cluster wide mechanism to enable or disable the transfer of PVCs
* Expose standardized process for transferring a PVC from one Namespace to another
* Require explicit acceptance of a PVC into another users Namespace, ensuring a malitious user can't just dump PVCs into another users Namespace
* Utilize existing PVC creation semantics for creating a PVC from a Transfer
* Honor existing configuration options including quota checks, Storage Class access etc

## Terms
* Originator - is the original owner/namespace of a PVC that is being transferred
* Recepient - is the owner/namespace that is allowed to receive the PVC being transferred

## Solution Overview

The proposed solution implements changes to the PV Controller to detect and handle PVC Transfer requests.  The proces is outlined below, and further details regarding the design are included in the DESIGN DETAILS:

1. Originator has a PVC that they've already populated with some data (Clone, Snapshot etc), this PVC (pvc-source), is bound and MAY or MAY NOT be ``in-use``.

    Originator wants to allow pvc-source to be transferred to a Recepient (namespace=development) and the Recepient has indicated they'd like to use a PVC Name of pvc-from-foo.
2. Originator adds the following annotation to pvc-source via a ``kubectl apply...``:
3. The PVC and PV are updated with a pending transfer annotation, this process can be aborted by the Originator by removing the transfer request annotations and perorming another update
4. The Originator deletes pvc-source (if it's desired that the process is aborted on the Oringiator side at this point, an admin must delete the PV), transfer status is still pending
5. The transfer is still pending until the Recepient requests a new PVC with the proper naming and transfer acceptance annotations, if/when a Claim request with the proper annotations is received and matches the transfer annotations, the Claim is fulfilled by the PV from the Originators Namespace

The process, including all checks against quotas, storage class etc are handled as they currently are with any PVC request.  If the Recepient never requests a Claim with the proper annotations, or they don't have access to the StorageClass, the transfer request will remain in a pending state.


## Design Details

The design of this proposal attempts to leverage the existing PVC create semantics, and leverage the existing policy, quota and other checks that are already in place.  The proposed design adds a check to the existing PV controller, that when an update to a PVC is detected that includes a transfer request, an annotation is added by the controller to the respective PVC and PV.  The result is that when the Originator deletes the PVC, the PV controller checks for this annotation and if it's present, rather than perform the clean up and delete the PV it instead will modify the claim ref for the PV to reflect the namespace/pvc-name in the transfer request.  The annotations added by the user are requests, while the annotations added by the controller are used as directives to fulfill the request.

* Example update to a PVC from an Originator in namespace ``bar`` to request a transfer of a PVC (pvc-1), to namespace ``foo`` and a pvc named ``pvc-from-pat``

  ``
 .....
kind: PersistenetVolumeClaim
metadata:
    name: pvc-1
    annotations:
        transferRequest: foo/pvc-from-pat
 .....
 ``

This action will only initiate the transfer request, initiation will add an annotation to the PVC and it's backing PV to indicate that a transfer has been requested, and will be marked as pending.  The original user can cancel the request at any time (up until the final step) by simply updating the PVC with a removal of the ``TransferDestination`` annotation.

* Annotations added to the PVC

  ``
  pvc.kubernetes.io/transfer-status: pending
  pvc.kubernetes.io/transfer-destination: foo/pvc-from-pat
  ....
  ``

* Annotations added to the PV

  ``
  pv.kubernetes.io/transfer-status: pending
  pv.kubernetes.io/transfer-destination: foo/pvc-from-pat
  pv.kubernetes.io/transfer-source: bar/pvc-1
  ....
  ``

* Receiving user adds the TransferAccept annotation to a PVC with the source Namespace/PVC-Name, this provides an explicit acknowledgement that the user wants a specific PVC from an outside Namespace

  ``
 .....
kind: PersistenetVolumeClaim
metadata:
    name: pvc-1
    annotations:
        TransferAccept: bar/pvc-1
.....
  ``
## Proof Of Concept

A simple POC implementation to test the basics of the concept is available on this [GitHub branch](https://github.com/j-griffith/kubernetes/tree/pvc-transfer-poc), and a [screencast demo](https://asciinema.org/a/210854) is availabe on asciinema.

## Outstanding Questions

1. The use of annotations versus adding a new field to the PVC Object

Rather than annotations it's worth considering a new ``Transfer`` struct for the PVC.

``
// Transfer represents a transfer request for a Kubernetes resource
type Transfer struct {
    // Kind is the resource being referenced
	// +optional (provided if we want to implement transfer of other resource types)
    Kind string
	// Source is the Originators namespace/pvc-name
	// +optionsal (only required for the acceptance of a transfer on the Receiving sides PVC request)
    Source string
	// Desitnation is the Recepient namespace/pvc-name
	// Source is the Originators namespace/pvc-name
	// +optionsal (only required for the origination of a transfer request on the Origintating side on an existing PVC)
	Destination string
}
``

2. Leveraging the existing PV controller, vs using a dedicated Transfer controller
   The current design implements the functionality by modifying the existing PV Controller.  This makes sense as long as there's no interest in extending the concept of transfering resources to other resources.

3. On the Recepient side, it might be worth considering using the existing DataSource field rather than introducing a new annotation model
   The PVC Object includes a DataSource field that is currently used for creating a Volume from a Snapshot.  There's currently a note that this is for Snapshots only, but it could be extended to allow other types of sources for PVC creation as well, including Transfer.  This would replace Recepient annotations, or Transfer: Source fields in the previously mentioned designs.  This would result in a PVC request that looks something like this:

``
 .....
kind: PersistentVolumeClaim
metadata:
    name: pvc-from-pat
dataSource:
    kind: PersistentVolumeClaim
	name: bar/pvc-1
 .....
 ``
