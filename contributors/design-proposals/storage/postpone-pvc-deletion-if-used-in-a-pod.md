# Postpone Deletion of a Persistent Volume Claim in case It Is Used by a Pod

Status: Proposal

Version: GA

Implementation Owner: @pospispa

## Motivation

User can delete a Persistent Volume Claim (PVC) that is being used by a pod. This may have negative impact on the pod and it may result in data loss.

For more details see issue https://github.com/kubernetes/kubernetes/issues/45143

## Proposal

Postpone the PVC deletion until the PVC is not used by any pod.

## User Experience

### Use Cases

1. User deletes a PVC that is being used by a pod. This may have negative impact on the pod and may result in data loss. As a user, I want that any PVC deletion does not have any negative impact on any pod. As a user, I do not want to experience data loss.

#### Scenarios for data loss
Depending on the storage type the data loss occurs in one of the below scenarios:
- in case the dynamic provisioning is used and reclaim policy is `Delete` the PVC deletion triggers deletion of the associated storage asset and PV.
- the same as above applies for the static provisioning and `Delete` reclaim policy.

## Implementation

### API Server, PVC Admission Controller, PVC Create
A new plugin for PVC admission controller will be created. The plugin will automatically add finalizer information into newly created PVC's metadata.

### Scheduler
Scheduler will check if a pod uses a PVC and if any of the PVCs has `deletionTimestamp` set. In case this is true an error will be logged: "PVC (%pvcName) is in scheduled for deletion state" and scheduler will behave as if PVC was not found.

### Kubelet
Kubelet does currently live lookup of PVC(s) that are used by a pod.

In case any of the PVC(s) used by the pod has the `deletionTimestamp` set kubelet won't start the pod but will report and error: "can't start pod (%pod) because it's using PVC (%pvcName) that is being deleted". Kubelet will follow the same code path as if PVC(s) do not exist.

### PVC Finalizing Controller
PVC finalizing controller is a new internal controller.

PVC finalizing controller watches for both PVC and pod events that are processed as described below:
1. PVC add/update/delete events:
  - If `deletionTimestamp` is `nil` and finalizer is missing, the PVC is added to PVC queue.
  - If `deletionTimestamp` is `non-nil` and finalizer is present, the PVC is added to PVC queue.
2. Pod add events:
  - If pod is terminated, all referenced PVCs are added to PVC queue.
3. Pod update events:
  - If pod is changing from non-terminated to terminated state, all referenced PVCs are added to PVC queue.
4. Pod delete events:
  - All referenced PVCs are added to PVC queue.

PVC and pod information are kept in a cache that is done inherently for an informer.

The PVC queue holds PVCs that need to be processed according to the below rules:
- If PVC is not found in cache, the PVC is skipped.
- If PVC is in cache with `nil` `deletionTimestamp` and missing finalizer, finalizer is added to the PVC. In case the adding finalizer operation fails, the PVC is re-queued into the PVC queue.
- If PVC is in cache with `non-nil` `deletionTimestamp` and finalizer is present, live pod list is done for the PVC namespace. If all pods referencing the PVC are not yet bound to a node or are terminated, the finalizer removal is attempted. In case the finalizer removal operation fails the PVC is re-queued.

### CLI
In case a PVC has the `deletionTimestamp` set the commands `kubectl get pvc` and `kubectl describe pvc` will display that the PVC is in terminating state.

### Client/Server Backwards/Forwards compatibility

N/A

## Alternatives considered

1. Check in admission controller whether PVC can be deleted by listing all pods and checking if the PVC is used by a pod. This was discussed and rejected in PR https://github.com/kubernetes/kubernetes/pull/46573

There were alternatives discussed in issue https://github.com/kubernetes/kubernetes/issues/45143

### Scheduler Live Lookups PVC(s) Instead of Kubelet
The implementation proposes that kubelet live updates PVC(s) used by a pod before it starts the pod in order not to start a pod that uses a PVC that has the `deletionTimestamp` set.

An alternative is that scheduler will live update PVC(s) used by a pod in order not to schedule a pod that uses a PVC that has the `deletionTimestamp` set.

But live update represents a performance penalty. As the live update performance penalty is already present in the kubelet it's better to do the live update in kubelet.

### Scheduler Maintains PVCUsedByPod Information in PVC
Scheduler will maintain information on both pods and PVCs from API server.

In case a pod is being scheduled and is using PVCs that do not have condition PVCUsedByPod set it will set this condition for these PVCs.

In case a pod is terminated and was using PVCs the scheduler will update PVCUsedByPod condition for these PVCs accordingly.

PVC finalizing controller won't watch pods because the information whether a PVC is used by a pod or not is now maintained by the scheduler.

In case PVC finalizing controller gets an update of a PVC and this PVC has `deletionTimestamp` set it will do live PVC update for this PVC in order to get up-to-date value of its PVCUsedByPod field. In case the PVCUsedByPod is not true it will remove the finalizer information from this PVC.

### Scheduler In the Role of PVC Finalizing Controller
Scheduler will be responsible for removing the finalizer information from PVCs that are being deleted.

So scheduler will watch pods and PVCs and will maintain internal cache of pods and PVCs.

In case a PVC is deleted scheduler will do one of the below:
- In case the PVC is used by a pod it will add the PVC into its internal set of PVCs that are waiting for deletion.
- In case the PVC is not used by a pod it will remove the finalizer information from the PVC metadata.

Note: scheduler is the source of truth of pods that are being started. The information on active pods may be a little bit outdated that causes that deletion of a PVC may be postponed (pod status in schedular is active while the pod is terminated in API server), but this does not cause any harm.

The disadvantage is that scheduler will become responsible for PVC deletion postponing that will make scheduler bigger.
