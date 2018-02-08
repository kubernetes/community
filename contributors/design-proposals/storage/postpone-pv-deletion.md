# Postpone deletion of a Persistent Volume if it is bound by a PVC

Status: Pending

Version: Beta

Implementation Owner: NickrenREN@ 

## Motivation

Admin can delete a Persistent Volume (PV) that is being used by a PVC.  It may result in data loss.

## Proposal

Postpone the PV deletion until the PV is not used by any PVC.


## User Experience
### Use Cases

* Admin deletes a PV that is being used by a PVC and a pod referring that PVC is not aware of this. 
This may result in data loss. As a user, I do not want to experience data loss.

## Implementation

### API Server, PV Admission Controller, PV Create: 

We can rename and reuse the PVC admission controller, let it automatically add finalizer information into newly created PV's and PVC's metadata.

PVC protection proposal: [PVC protection](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/postpone-pvc-deletion-if-used-in-a-pod.md)

#### PV controller:

When we look for best matched PV for the PVC, if the PV has the  deletionTimestamp set, we will not choose it even if the PV satisfies all the PVC’s requirement.

For pre-bound PVC/PV, if the PV has  the  deletionTimestamp set, we will not perform the `bind` operation and keep the PVC `Pending`.

#### PV Protection Controller:

PV protection controller is a new internal controller.

Since we already have PV controller which is responsible for synchronizing PVs and PVCs, here in PV protection controller, 
we can just watch PV events.

PV protection controller watches for PV events that are processed as described below:

* PV add/update/delete events:
    * If deletionTimestamp is nil and finalizer is missing, the PV is added to PV queue.
    * If deletionTimestamp is non-nil and finalizer is present, the PV is added to PV queue.

PV information is kept in a cache that is done inherently for an informer.

The PV queue holds PVs that need to be processed according to the below rules:

* If PV is not found in cache, the PV is skipped.
* If PV is in cache with nil deletionTimestamp and missing finalizer, finalizer is added to the PV. In case the adding finalizer operation fails, the PV is re-queued into the PV queue.
* If PV is in cache with non-nil deletionTimestamp and finalizer is present, we try to get the PV's status, if it is not `bound`(synchronized by PV controller), the finalizer removal is attempted. The PV will be re-queued if the finalizer removal operation fails.

#### CLI:

If a PV’s deletionTimestamp is set, the commands kubectl get pv and kubectl describe pv will display that the PV is in terminating state.


### Client/Server Backwards/Forwards compatibility

N/A

## Alternatives considered

### Add this logic to the existing PV controller instead of creating a new admission and protection controller
When we bind PV to PVC, we add finalizer for PV and remove finalizer when PV is no longer bound to a PVC.
