# Local PV Distributed Static Provisioner
Authors: @verult, @msau42

The current preference is [design 5](#design-5-worker-as-jobs-watcher-in-daemonset). In all the designs below, assume all relevant API resources are cached.
        
## Design Goals
1. Limit all PV object interactions to a single master pod in order to minimize node access to the Kubernetes system. The original discussion is [here](https://github.com/kubernetes/community/pull/989#discussion_r135397619).
1. Assume disk deletion is asynchronous, as required by block volumes.
1. If a node is compromised, damage is contained to that particular node.

## Related Works
* Original provisioner design, [here](https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume/provisioner#design).
* Disk cleaner (deleter) implementation for block volumes, [here](https://github.com/kubernetes-incubator/external-storage/pull/312).

## Design 1: Master-Heavy

### High Level Overview
The static provisioner will be split into a single master (ideally placed on the master node) and many workers (one per node). The master contains the bulk of provisioner logic, including syncing PV objects with the API server and triggering discovery and deletion. Workers are left to perform the physical discovery and deletion of the local volumes on its node. gRPC is used to communicate between master and workers. This design uses an event-driven model: workers notify master when a status change is available, and master then pulls the status from worker. Master also performs a long poll.

### Worker
The worker provisioner exposes the following RPCs:

| RPC                           | Description |
| --- | --- |
| `GetWorkerStatus()`           | Returns the worker status, as described in the next section. |
| `Delete(LocalVolume)`    | Starts the asynchronous delete of the given PV. |

In addition, the Discoverer and Deleter from the previous design are kept here, but they only perform filesystem operations. All other logic is moved to the master.

### ProvisionerStatus
The following fields are retrieved as part of the `GetWorkerStatus()` call:

* `LocalVolumes` - A collection of volumes mounted in the local storage directory on the node. The discoverer is triggered to acquire this information.
* `DeleteOperations` - A list keeping track of information related to each delete operation. This status is managed by the cleaner. It contains the following fields:
  * PV - The PV to be deleted on disk.
  * DeleteStatus - takes one of the following values:
    * Requested
    * InProgress
    * Complete
    * Error

`DeleteOperations` does not have to be correct after worker restarts. See [PV Finalizer](#pv-finalizer) for more info.

### Master
The master provisioner exposes the following RPC:
| RPC                           | Description |
| --- | --- |
| `ProvisionerStatusReady()`             | Notifies master that there is a status update. Also called on worker initialization. |

When it receives a notification from a worker, it
1. Calls the RPC `GetWorkerStatus()`.
1. Using `LocalVolumes` from the worker status, **reconciles** with the API server, i.e. creates an additional PV if it exists in `LocalVolumes` but not in API server *and no delete is in progress for this PV*.

When it receives an API server update that a PVC bound to a local volume is released, and if a deletion is not already in progress, call the `Delete()` RPC.

**TODO**: gRPC backoff strategy in the case of network partition or worker failure.

### PV Finalizer
If a user is allowed to manually delete a PV while the local volume provisioner is performing a deletion, workers must guarantee that DeletionStatus is correct at all times, even in the wake up worker failure. This is to prevent the discovery component from creating a PV while a deletion is underway, which causes a corrupted volume to become available.

To solve this problem, master adds a finalizer to the PV it's about to delete, preventing users from accidentally deleting the PV while the provisioner is performing a delete. This removes the requirement of consistent DeletionStatus, and thus the need for workers to persist state. If a PV has the finalizer but is not in its corresponding worker's DeletionStatus, master needs to trigger another delete.

### Worker Discovery using HostPorts
The master needs to determine which worker to talk to given a node name (tracked in the PV managed by the provisioner). To do this, master deploys the worker DaemonSet and assigns them a random HostPort. The cluster admin specifies a range of possible HostPorts in order to minimize port collisions. To communicate with a worker on a particular node, master fetches the IP of the node and uses the HostPort to reach the worker directly.

### Considerations

* Client authentication? Must be able to establish master identity in order to protect against malicious delete operations.

* Upgrading the local volume provisioner requires all Deployments and DaemonSets to be recreated. Rolling update is not possible.

* Should `Delete()` return when deletion starts or completes?

## Design 2: Worker-Heavy

### High Level Overview
This design is similar to Design 1, with the key difference being workers listen to PVC changes instead of master. As a result, deletion is no longer triggered by master.

### Worker
The bulk of deletion logic is moved to workers. When a worker detects a PVC it manages is deleted, it triggers the deletion process, updating its internal deletion state accordingly. The discoverer now watches the deletion state such that if a deletion is underway for a volume it detects, it will not add it in its list of discovered local volumes. This logic is similar to Design 1 except it's moved from master to worker.

The worker provisioner exposes the following RPCs:

| RPC                           | Description |
| --- | --- |
| `GetWorkerStatus()`           | Returns the worker status, as described in the next section. |
| `DeleteACK(LocalVolume)` | Acknowledges the completion of this PV. Removes the DeleteStatus of the corresponding volume. |


### ProvisionerStatus
The following fields are retrieved as part of the `GetWorkerStatus()` call:

* `LocalVolumes` - A collection of volumes mounted in the local storage directory on the node. The discoverer is triggered to acquire this information.
* `DeleteOperations` - A list keeping track of information related to each delete operation. This status is managed by the cleaner. It contains the following fields:
  * PV - The PV to be deleted on disk.
  * DeleteStatus - takes one of the following values:
    * InProgress
    * Complete
    * Error

Note that the Requested status has been removed. Deletion start, completion, and errors are also reported through pod events.

### Master
Comparing to Design 1, master retains the same RPC calls. It no longer has the responsibility to trigger volume deletion, and instead only updates PVs. When it detects that a worker's `LocalVolumes` is different from its PV cache, it creates a PV for each volume. When a worker reports that a volume deletion is complete, master removes the finalizer and deletes the PV.

### Considerations

* How should worker failures be handled when master ACKs a delete?

* Client authentication no longer necessary because master doesn't trigger deletion.

## Design 3: Event-Driven with Custom Resources

### High Level Overview
The previous two designs are complicated by the asynchronous nature of delete and the presence of failures. 
As recommended by @dhirajh, another design is to persist ProvisionerStatus as a Custom Resource, and to drive all operations with PV and ProvisionerStatus Informer events.

To ensure no malicious pods can affect provisioner execution, the CustomResourceDefinition and all pods must be in a trusted admin namespace.

In the following description, we assume all Informer updates are cached and events buffered. On initialization, all caches are populated, and every Informer update event handler is executed once to ensure the entire provisioner system is in a consistent state.

### ProvisionerStatus CRD
The CRD fields are identical to those of the ProvisionerStatus in Design 1. There exists one ProvisionerStatus object per node.

To ensure correctness, master and worker never write to the same field. Specifically:
* After a `DeleteOperation` is created, its `PV` field is never modified.
* Master can only create and delete `DeleteOperation`s. It never modifies those objects.
* Worker can only modify `DeleteStatus`. It can never create or delete `DeleteOperation`s.
* Only a worker can write to `LocalVolumes`.

### Worker
Create the ProvisionerStatus object on initialization.

The Discoverer periodically probes the filesystem and updates ProvisionerStatus `LocalVolumes`. This can be further improved by using a filesystem watch.

The Deleter uses ProvisionerStatus to keep track of states for each PV being deleted. Deletion is started when the ProvisionerStatus is updated (by the master) and the DeleteStatus of some volumes are in the Requested state. Details will be included in the async volume cleaner design.

### Master
A single goroutine processes events from all informers.

On ProvisionerStatus Informer update, perform the following actions:
1. For deletion: If a volume is in the Completed state, delete the corresponding PV from the API server and remove the entry from `DeleteOperations`.
1. For discovery: compare `LocalVolumes` with existing PVs in the API server. If a volume exists in `LocalVolumes` but has not been persisted, and if the disk is not about to be deleted (i.e. in Requested or InProgress state), then create the PV.

On PV Informer update: if the updated PV is in VolumeReleased state, add the PV to `DeleteOperations` with the Requested state.

On Node deletion: delete the associated ProvisionerStatus object.

**TODO**: Consider marking cache dirty when its corresponding API object is modified within master and ignoring Informer events if cache is dirty. This ensures master is always operating on the most up-to-date data. 

### Considerations
* This design has the following advantages:
  * Because the API server persists state reliably, we can simplify our reasoning with failure modes.
  * Retries and backoffs are no longer necessary.
  * Provisioner state is more visible to admins.

* If a node is compromised, it has access to modify CRDs representing other nodes.
  * What if ProvisionerStatus is non-namespaced? Only admins should have access. If a compromised node has access to all non-namespaced resources, there is a bigger security issue.

## Design 4: Hybrid RPC-CRD Approach
Keep the same CRD object described in Design 3, but only master has write access to them. Workers must persist its delete state locally and report its status to master through gRPC. Master makes requests to workers by writing to CRD, and fetches worker status by polling (similar to Design 1).

**TODO**: Use network policy to limit pod communication to only within the provisioner system.

### Considerations
* It's not safe against man-in-the-middle attacks as is. Requires encryption.

## Design 5: Worker as Jobs, Watcher in DaemonSet

### Overview
One of the major issues with previous approaches is worker identity - when a worker tries to communicate with master, it has to prove it's the worker it claims to be. To fix these issues, each worker must maintain different credentials.

Alternatively, master can actively create worker Jobs to perform discovery and delete rather than having long-running worker pods in a DaemonSet sending data to the master. To trigger discovery, we also have local volume directory watcher pods running in a DaemonSet.

### Worker
Workers are spawned as Jobs by master on demand. In discovery mode, a worker writes the list of local volumes in JSON format to a file inside an emptyDir. A sidecar container then reads the file and writes to its STDOUT. In deletion mode, the worker executes async delete and uses the status of the Job as an indication of deletion state.

### Watcher
Watchers are deployed as a DaemonSet. Each watcher keeps an in-memory cache of local volumes and periodically checks the volume directory for changes. If there is a change, it notifies master through gRPC.

If there is a network failure, the gRPC call will fail and the watcher retries with exponential backoff. If master fails after receiving the notification, once it comes back up it will spawn a discovery worker on every node (as explained in the Master section).

### Master
Master exposes the same RPC calls as Design 1 to enable discovery notifications.
The master provisioner exposes the following RPC:
| RPC                           | Description |
| --- | --- |
| `LocalVolumeReady()`             | Notifies master that there are new local volumes. Also called on watcher initialization. |

Master can be made completely event driven using Job informers and watcher notification. Master deploys a discovery worker when it gets a notification from a watcher. When the Job is complete, master parses the output of the sidecar container as the list of discovered volumes. If the output is empty (may occur if container logs have been rotated) or invalid, it must re-deploy the worker. Master deploys a deletion worker when a PVC is released and uses the Job's status as deletion state. When the Job is complete, it removes the corresponding PV's finalizer and the PV resource itself.

To ensure the provisioner system has up-to-date state on initialization, the following needs to be done:
* When the master initializes, it spawns discovery jobs on every node to make sure it knows about all local volumes in the cluster.
* When a watcher initializes, it should notify master so a discovery worker can be spawned to fetch the latest list of local volumes.

### Considerations
* The API server and kubelet ensure all communications are secure. There's no additional security overhead.

* Master has perfect knowledge of which Job runs on which node, so if any worker is compromised, the damage is contained within the node the worker is scheduled on.

* Instead of writing local volumes to a file, a worker can send them back to master through gRPC. However, this requires encryption and retry + exponential backoff.

## Conclusion
Design 5 provides the security guarantees the local volume provisioner needs without additional overhead. It provides visibility of discovery and deletion statuses with API server resources and container logs. Each pod in the DaemonSet is also smaller in size and actual workers are spawned only as needed. Thus Design 5 is the preferred solution.
