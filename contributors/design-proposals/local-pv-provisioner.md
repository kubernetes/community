# Local PV Distributed Static Provisioner
Authors: @verult, @msau42

The current preference is [design 5](#design-5-worker-as-jobs-watcher-in-daemonset).
        
## Design Goals
1. Limit all PV object interactions to a single master pod in order to minimize node access to the Kubernetes system. The original discussion is [here](https://github.com/kubernetes/community/pull/989#discussion_r135397619).
1. Assume disk deletion is asynchronous, as required by block volumes.

In all designs mentioned below, the provisioner will be run in a master-worker configuration. ProvisionerWorkers are deployed on every node in a DaemonSet (just as before) and ProvisionerMaster is deployed in a Deployment on a trusted node.

## Related Works
* Original provisioner design, [here](https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume/provisioner#design).
* Disk cleaner (deleter) implementation for block volumes, [here](https://github.com/kubernetes-incubator/external-storage/pull/312).

## Design 1: Master-Heavy

### High Level Overview
The static provisioner will be split into a single master (ideally placed on the master node) and many workers (one per node). The master contains the bulk of provisioner logic, including syncing PV objects with the API server and triggering discovery and deletion. Workers are left to perform the physical discovery and deletion of the local volumes on its node. gRPC is used to communicate between master and workers, with the master as the client and each worker as a separate server.

The master keeps track of `WorkerInfo`, a collection containing one entry for each worker keeping track of relevant information. The master also runs a group of goroutines, one for each worker, which controls and interacts with the workers.

### Worker Provisioner
The worker provisioner exposes the following RPCs:

| RPC                           | Description |
| --- | --- |
| `GetWorkerName()`             | Returns the worker provisioner's name. |
| `GetProvisionerStatus()`      | Returns the worker status, as described in the next section. |
| `Delete(PersistentVolume)`    | Starts the asynchronous delete of the given PV. |

In addition, the Discoverer and Deleter from the previous design are kept here, but they only perform filesystem operations. All other logic is moved to the master.

### ProvisionerStatus
The following fields are retrieved as part of the `GetWorkerStatus()` call:

| Field               | Description |
| --- | --- |
| `LocalVolumes`          | A collection of PVs mounted in the local storage directory on the node. The discoverer is triggered to acquire this information. |
| `DeletesInProgress` | A collection of PVs currently being deleted asynchronously. This status is managed by the cleaner. |

Every worker must maintain the invariant that `DeletesInProgress` is correct even after a worker restart. In the case that a user manually deletes the PV currently being deleted by a worker, master can only depend on the WorkerStatus to ensure a replacement PV isn't created.

### WorkerInfo
Each WorkerInfo entry contains the following fields:

| Field                 | Description |
| --- | --- |
| `Name`                | Name of the worker provisioner. |
| `Endpoint`            | The worker's host address. |
| `PVCache`             | Cache of API PV objects managed by this worker. |
| `StopCh`              | Used to terminate a goroutine. |

### Goroutine
The goroutine will start by populating the `apiCache` using the populator logic that already exists in the previous design. Then, it will run the following loop with several seconds of delay in between executions:
1. Call the RPC `GetWorkerStatus()`.
1. Using `LocalVolumes` from the worker status, **reconcile** with the `PVCache`. Create an additional PV if it exists in `LocalVolumes` but not in `PVCache` *and no delete is in progress for this PV, i.e. it's not in `DeletesInProgress`*.
1. For each PV in `VolumeReleased` state (ready for cleanup), if it's not in `DeletesInProgress`, call the `Delete()` RPC.

Master must delete the PV object after the disk delete operation is complete. In order for master to know when this occurs, the worker must persist a "complete" state, because in the wake of worker failure, master has to differentiate between an error and a completion. Master also needs to notify the worker that the PV is deleted, so the persisted state can be removed.

**TODO**: gRPC backoff strategy in the case of network partition or worker failure.

### Worker Discovery
The master needs to have the host address of every worker in order to communicate with them directly. To do this, workers expose a single Service, which is associated with an Endpoints object containing all worker addresses.

To listen to worker membership changes, master uses an Endpoints informer. An informer update is triggered for every endpoint addition and removal. For each informer update, master sends a `GetWorkerName()` call to every endpoint to obtain their identity. If master is aware of the worker, then it updates the worker's `Endpoint`. Otherwise, master adds a new entry to `WorkerInfo` and spawns a new goroutine. If a worker no longer exists, master sends a signal through the `StopCh` of the goroutine handling the worker associated with this endpoint, then deletes the associated entry. **TODO**: Master must wait until goroutine finishes before `WorkerInfo` entry is removed, but this wait could delay the update of other Endpoints. This is also related to gRPC backoff strategy. Need to handle this.

At initialization, master iterates through fetched endpoints and adds them appropriately to the system as described above. 

It's possible for a new pod to reuse the IP of a deleted pod, so even if the endpoint stays the same over time, the associated worker might be different. To fix this issue, master can pass the target worker's name in every RPC call, and have each worker verify the name. Eventually, worker discovery will assign the Endpoint to the correct worker.

In the case of network partition or worker failure during the `GetWorkerStatus()` call, retry a few times, then skip this worker. If the Endpoint is unreachable, the Informer will eventually be triggered, which will correct this scenario.

## Design 2: Worker-Heavy

### High Level Overview
Another possible design is to leave most of provisioner logic inside workers, only moving PV API interactions to the master. The worker still maintains a PV cache, which is populated by the master, but the difference in this design is it's solely updated by the master. Without this difference, managing the correctness of cache data becomes more difficult, especially when the master fails. When a worker needs to update a PV object, it puts the request in a queue, which is periodically polled by the master.

### Worker Provisioner
The worker provisioner retains most key components from the previous design, including the Cache, Discoverer, and Deleter. The Populator is moved to master because it interacts with the API server. In order to interact with the master, it exposes the following gRPC calls:

| RPC                      | Description |
| --- | --- |
| `Add/Update/RemovePV(PersistentVolume)` | Updates the Cache with the given PV. Called by master when PV informer is triggered. |
| `PushCurrentPVs([]PersistentVolume)`    | Merges given list of PVs with the Cache. Used by the Populator. |
| `GetPVRequest()`                        | Peeks at the next unhandled request in the queue. |
| `CompletePVRequest(PVRequest)`          | Signals that the last request is completed and can be removed from the queue. |

In the case of master failure while a request is being processed, the worker needs to keep record of the request so that a new master can process it again. Thus it's necessary for master to signal request completion.

### Master Provisioner
The populator is initialized so that it calls the appropriate gRPC for every PV Informer update. Then it populates each worker's cache by calling `PushCurrentPVs()`, passing the collection of PVs to be controlled by this worker.

The master then performs worker discovery as described in the [Worker Discovery](#worker-discovery) section of design 1, and spawns a goroutine for each worker. Each goroutine periodically calls `GetPVRequest()`, performing that request if it receives one. Then it calls `CompletePVRequest()` to signal the request completion.

## Design 3: Event-Driven with Custom Resources

### High Level Overview
The previous two designs are complicated by the asynchronous nature of delete and the presence of failures. 
As recommended by @dhirajh, another design is to persist ProvisionerStatus as a Custom Resource, and to drive all operations with PV and ProvisionerStatus Informer events.

To ensure no malicious pods can affect provisioner execution, the CustomResourceDefinition and all pods must be in a trusted admin namespace.

In the following description, we assume all Informer updates are cached and events buffered. On initialization, all caches are populated, and every Informer update event handler is executed once to ensure the entire provisioner system is in a consistent state.

### ProvisionerStatus CRD
The CRD contains the following fields (similar to design 1):

* `LocalVolumes` - A collection of volumes mounted in the local storage directory on the node. The discoverer is triggered to acquire this information.
* `DeleteOperations` - A list keeping track of information related to each delete operation. This status is managed by the cleaner. It contains the following fields:
  * PV - The PV to be deleted on disk.
  * DeleteStatus - takes one of the following values:
    * Requested
    * InProgress
    * Complete
    * Error
    
There exists one ProvisionerStatus object per node.

To ensure correctness, master and worker never write to the same field. Specifically:
* After a `DeleteOperation` is created, its `PV` field is never modified.
* Master can only create and delete `DeleteOperation`s. It never modifies those objects.
* Worker can only modify `DeleteStatus`. It can never create or delete `DeleteOperation`s.
* Only a worker can write to `LocalVolumes`.

### Worker Provisioner
Create the ProvisionerStatus object on initialization.

The Discoverer periodically probes the filesystem and updates ProvisionerStatus `LocalVolumes`. This can be further improved by using a filesystem watch.

The Deleter uses ProvisionerStatus to keep track of states for each PV being deleted. Deletion is started when the ProvisionerStatus is updated (by the master) and some PVs are in the Requested state. Details will be included in the async volume cleaner design.

### Master Provisioner
A single goroutine processes events from all informers.

On ProvisionerStatus Informer update, perform the following actions:
1. For deletion: If a volume is in the Completed state, delete the corresponding PV from the API server and remove the entry from `DeleteOperations`.
1. For discovery: compare `LocalVolumes` with existing PVs in the API server. If a volume exists in `LocalVolumes` but has not been persisted, and if the disk is not about to be deleted (i.e. in Requested or InProgress state), then create the PV.

On PV Informer update: if the updated PV is in VolumeReleased state, add the PV to `DeleteOperations` with the Requested state.

On Node deletion: delete the associated ProvisionerStatus object.

**TODO**: Consider marking cache dirty when its corresponding API object is modified within master and ignoring Informer events if cache is dirty. This ensures master is always operating on the most up-to-date data. 

## Design 4: Hybrid RPC-CRD Approach
Keep the same CRD object described in Design 3, but only master has write access to them. Workers must persist its delete state locally and report its status to master through gRPC. Master makes requests to workers by writing to CRD, and fetches worker status by polling (similar to Design 1).

Several modifications to Design 1:
* For discovery, do not rely on worker to report its identity as it can be compromised. Identify workers using the Endpoints object (pod name and node name).
* In future versions, master goroutines should be event-driven, as constant polling is not scalable to large number of nodes. A worker pings the master when it's ready to report status, and master makes a separate call to fetch status from the worker. Worker must continuously ping master until it receives a status request to ensure master recognizes the status update. As a safety check, master should also poll workers periodically, but with very low frequency.
  * It's not as secure to have workers push status directly because a worker can't be accurately identified from a gRPC request (the request IP can be spoofed).
  * It's more convenient to use gRPC server streams, but it's hard to scale this to thousands of nodes as streams require TCP connections to be kept open.

**TODO**: Use network policy to limit pod communication to only within the provisioner system?

## Design 5: Worker as Jobs, Watcher in DaemonSet
One of the major issues with previous approaches is worker identity - when a worker tries to communicate with master, it has to prove it's the worker it claims to be. To fix these issues, each worker must maintain different credentials.

Alternatively, master can actively create worker Jobs to perform discovery and delete rather than having long-running worker pods in a DaemonSet sending data to the master. For discovery, workers can store the list of local volumes in its container STDOUT (**TODO**: decide on a form of structured output rather than plain text). For deletion, the status of the Job acts as an indication of deletion state. Master has perfect knowledge of which Job runs on which node, so if any worker is compromised, the damage is contained within the node the worker is scheduled on.

To trigger discovery, we have local volume directory watcher pod running in a DaemonSet. It keeps an in-memory cache of local volumes and periodically checks the volume directory for changes. If there is a change, it notifies master through either gRPC or a CRD notification object, and master can then spawn a discovery Job.

To ensure the provisioner system has up-to-date state on initialization, the following needs to be done:
* When the master initializes, it spawns discovery jobs on every node to make sure it knows about all local volumes in the cluster.
* When a watcher initializes, it should notify master so a discovery Job can be spawned to fetch the latest list of local volumes.

Master can be made completely event driven using Job informers.

## Conclusion
One main downside of design 2 is the possibility of duplicate discover and delete operations, as a result of the cache not being updated immediately. Also, because the logic of PV updates is divided over the network, it's easy for PVs to end up in an inconsistent state.

Design 3 has the following advantages:
* Master has perfect knowledge of when to issue PV delete (based on the Complete DeleteStatus) without requiring multiple acknowledgment messages between master and worker.
* Because the API server persists state reliably, we can simplify our reasoning with failure modes.
* Provisioner discovery is no longer necessary.
* Retries and backoffs are no longer necessary.
* Provisioner state is more visible to admins.
* Almost entirely event-driven (the only exception being the filesystem probe).
But one major concern with this approach is each worker has access to all ProvisionerStatus objects. If a worker is compromised, it could disrupt the entire provisioner system instead of just its node.

Design 4 is a compromise, taking advantage of benefits from other designs. However, it would not work if the network can be damaged to the point that a compromised worker can intercept all traffic to the IP of another normal worker.

Design 5 removes the need to maintain worker identity, and thus is the simplest design that protects against our security concerns. Each pod in the DaemonSet is also smaller in size and actual workers are spawned only as needed.
