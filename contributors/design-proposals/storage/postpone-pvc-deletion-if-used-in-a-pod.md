# Postpone Deletion of a Persistent Volume Claim in case It Is Used by a Pod

Status: Proposal

Version: GA

Implementation Owner: @pospispa

## Motivation

User can delete a Persistent Volume Claim (PVC) that is being used by a pod. This may have negative impact on the pod.

For more details see issue https://github.com/kubernetes/kubernetes/issues/45143

## Proposal

Postpone the PVC deletion until the PVC is not used by any pod.

## User Experience

### Use Cases

1. User deletes a PVC that is being used by a pod. This may have negative impact on the pod. As a user, I want that any PVC deletion does not have any negative impact on any pod.

## Implementation

### PVC Create
When a PVC is created finalizer is added into its metadata.

### Finalizer
Finalizer can list pods and watch for pods. It caches status of pods that use a PVC. It also watches for PVC deletes.

Finalizer reacts to PVC delete in one of the below ways:
- In case the PVC is not used by a pod: removes the finalizer information from the PVC metadata.
- In case the PVC is used by a pod: sets condition Terminating to true for the PVC.

### Scheduler
Scheduler will check if a pod uses a PVC and if any of the PVCs has condition Terminating set to true. In case this is true the scheduling of the pod will end up with an error: "PVC (%pvcName) is in Terminating state".

### CLI
Commands `kubectl get pvc` and `kubectl describe pvc` will display the Terminating condition.

### Client/Server Backwards/Forwards compatibility

N/A

## Alternatives considered

1. Check in admission controller whether PVC can be deleted by listing all pods and checking if the PVC is used by a pod. This was discussed and rejected in PR https://github.com/kubernetes/kubernetes/pull/46573

There were alternatives discussed in issue https://github.com/kubernetes/kubernetes/issues/45143
