# Defer Containers

dhilip.kumar.s@huawei.com

March 2017

## Prerequisite
Understanding of  [initContainers](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/container-init.md)

## Motivation
Since the introduction to Statefulset and Daemonset PODs that consist of stateful containers are now fully supported by
Kubernetes.  Such stateful pods benefit from constructor and destructor semantics. InitContainers provide constructor
semantics. The proposed DeferContainers will provide destructor semantics for such stateful PODs

## Abstract
This introduces the concept of ‘deferContainers’, inspired by golang’s keyword ‘defer’. This will allow one to define a
set of Containers to be executed at the time of POD termination. The defined containers will be executed sequentially and
in the specified order. It will bring destructor() capability to Pods.

## Use Cases
[5mins Uses Presentation Video during recent Sig-App Meeting](https://youtu.be/s1oZ00_JA00?t=2692)

### Cleanup:
Most stateful workloads require certain Cleanup activity before or after the appContainers are terminated such as
* Sync/Flush the Disk or logs before the POD is evicted off a node.
* Delete or Update Global Configuration / Release Application level Lock.
* Some Legacy applications may not respond to SIGTERM signals sent by docker daemon, they might require special command/procedure
  to initiate graceful shutdown.

### Shard rebalancing.
In a sharded workload configured as a statefulset, if the application is scaled down, then the exiting Pod should trigger
a re-balance of keys from the current shard to the rest of them. This needs to be done befoe the shard goes down.  
‘deferContainers’ could handle such scenarios with ease.

### Statefulset Upgrading
If in future, statefulsets support rolling updates, an update request might attempt to replace Pods one after the other. 
With deferContainers each Pod could be removed gracefully without having to implement complicated sidecar logic to prevent data loss.

### Master-Slave or Leader-Follower Statefulset down-size
Consider a Stateful set consisting a Master-Slave or Leader-follower type application running N replicas. Ordinal index 
ranging from 0 to N-1.  If you scale it down (reduce the number of replicas), the statefulset controller would attempt 
to bring down the last replica (N-1th pod in this case).  If that pod is the elected Leader or Master, then this would 
disrupt the service intermittently. With deferContainers, this handover or re-election could be programmed gracefully/elegantly.

### SHUTDOWN Sequence
* Traditional relational  databases typically support a reasonable shutdown sequence, for instance, Oracle has 4 types of 
shutdown such as NORMAL, IMMEDIATE, TRANSACTIONAL and ABORT. ‘deferContainers’ will allow us to program and wait for such
complex shutdown scenarios. 
* In future when kubernetes supports Virtual Machine runtime (eg: hyperContainer) for better isolation, we should shutdown
  the VMs instead of killing them abruptly.  ‘deferContainers’ could help us run such shutdown commands.

## Limitation with the current system
Container pre-stop lifecycle hook is not sufficient for all termination cases:
* It is container Specific and not pod specific
* They cannot easily coordinate complex termination conditions across containers in multi-container pods
* They can only function with code in the image or code in a shared volume, which would have to be statically linked 
  (not a common pattern in wide use)
* Does not work across kubelet restart
* Waits for the entire graceperiod even after the pre-stop hook finished earlier.
* Wont restart on failed termination steps.
* Cannot contain complex termination scripts as no logging support.

## Design Requirements
* deferContainers should be able to:
  * Use the same volume (PV) as appContainers such that it can
    * Perform cleanup of shared volume, such as delete several temp directories.
    * Delete unwanted files before a final sync is initiated.
    * Update Configuration files about the changes in the distributed system so that the next pod getting attached to this PV  will benefit from it. (like a new leader/master etc).
    * Deleted secrets or security related files before the pod de-couples from the PV.
  * Delay the termination of application containers until operations are complete
  * De-Register the pod with other components of the system
  * Program termination sequence for cases where TerminationGracePeriod will be hard to predict before hand.

* Reduce coupling:
  * Between application images, eliminating the need to customize those images for Kubernetes generally or specific roles
  * Inside of images, by specializing which containers perform which tasks (install git into init container, use filesystem 
    contents in web container)
  * Between termination steps, by supporting multiple sequential cleanup containers
* Pre-Exit 
  * Should act as pre-exit trigger, should be called when the application is about to be deleted.
* restart on Failure
  * If a certain deferContainer failed while execution it should be automatically restarted
* GracePeriod behaviour
  * It should be possible to mention overall terminationGracePeriod for defercontainers, if the termination sequences completed before the overall graceperiod then the pod should be deleted without waiting further.
* Reduce Complexity, 
  * It should be possible to use a generic container as a deferContainer,
  * A deferContainer should be independently invokable, ie:- should not require code in the same image as the appContainers.
  * deferContainer Images that are not already in the node will be pre-populated while the application is being executed.

## Design
This proposed pod spec would look like below.
```yaml
pod:
  spec:
    initContainers: ...
    containers: ...
    deferContainers:
    - name: defer-container1
      image: ...
      ...
    - name: defer-container2
    ...
  status:
    initContainersStatuses: ...
    containerStatuses: ...
    deferContainerStatuses:
    - name: defer-container1
      ...
    - name: defer-container2
      ...
```
The api will look like below
```
// PodSpec is a description of a pod.
type PodSpec struct {
.
.
.

	InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Cannot be updated.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	//List of termination Containers, those will be executed when during the TerminationGracePeriod of the pod
	// +patchMergeKey=name
	// +patchStrategy=merge
	DeferContainers []Container `json:"deferContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,26,rep,name=deferContainers"`
.
.
.
}
```
* Will have 0...N containers and will be executed in sequence (specified order)
* Restart policy for deferConatiners are ‘onFailure’
* Adding a new phase ‘Terminating’ in the Pod lifecycle

### Terminating Phase / Defer Phase
* A POD reaches terminating phase when it is about to be removed from Kubernetes Cluster
* During this phase, the appContainers are not restarted if they get terminated/killed
* deferContainers will be executed one-after-the-other in the same sequence as they were specified in the POD Spec. From the above
  pod spec example execution sequence will be defer-Container1, then defer-Container2, …, etc.,
* if a particular deferContainer failed, it will be restarted until it succeeds.
* if the user specifies `kubectl delete pod foo --grace-period=0 --force` to delete a pod deferContainers will not be executed.
```
Example status output when a pod is being terminated.
   NAME      READY     STATUS     RESTARTS   AGE
    foo-0     2/2       Defer:0/4   0          7m
```
* Failure of one or all deferContainers will not trigger a POD restart.
* If deferContainers are configured Pre-Stop hooks will not be executed. 

### TerminationGracePeriod
* It takes default value (30 seconds as of today), explicitly mentioning this flag overrides the default value.
* Then deferContainer will start to execute one after the other in the specified order
* If a particular deferContainer failed it will be restart until it succeeds or graceperiod is exhausted. 
* When the configured graceperiod expires then all the containers (AppContainers) including the current deferContainer will be terminated. 
* It will kill currently executing deferContainer and no further deferContainer will be executed (if there are any). 
* deferContainers are time bound by TerminationGracePeriod 
* If all the deferContainers completed execution well ahead of TerminationGracePeriod, then we should 

### PrePopulate  deferContainers Images
By default, all the deferContainer images will be pulled (if not available) when the POD reaches ‘running’ stage.

## Implementation Plan
Development and release lifecycle of this feature will follow other kubernetes experimental feature. This will be originally
rolled out as alpha as annotations based on its usefulness and community feedback it will graduate to PodSpec.

## Examples 
### Cleanup
```yaml
pod:
  spec:
    terminationGracePeriod: 60
    initContainers: ...
    containers: ...
    deferContainers:
#pre exit operations
    - name: fsync
      image: my-utils
#Something like this https://docs.mongodb.com/manual/reference/command/fsync/
      comamnd: ["/bin/sh", "-C", "./syncDisk.sh"] 
      name: killApp
      image: my-utils
      command: ["/bin/sh", "-C", "./killAndWait.sh", "--name=${POD_NAME}"]  
#post exit operation
    - name: rm-tmpdir
      image: my-utils
      command: ["/bin/sh", "-C", "./disk_cleanup.sh"]
#contact and inform a thirdparty system about this pods termination, such as reducing a reference counter (if one is maintained)
      name: ref-counter
      image: my-utils
      comamnd: ["/bin/sh", "-C", "./decrementRefCount.sh"]
      ...
```

### Master-slave / Leader-follower statefulset down-size / scale down the replicas
Below scripts 'selectaMaster.sh and reConfSlaves.sh should be designed in such a way that even if a terminating pod is a slave it
should not affect the cluster. This will fit controllers such as Statefulset because they guarantee only one pod goes down at once.
```yaml
pod:
  spec:
    terminationGracePeriod: 60
    initContainers: ...
    containers: ...
    deferContainers:
    - name: electMaster
      image: my-utils
      #Select a mastter among the available slaves
      comamnd: ["/bin/sh", "-C", "./selectAMaster.sh"] 
      name: reconfigureSlaves
      image: my-utils
      #Re-configure all the slaves to the new master
      command: ["/bin/sh", "-C", "./reconfSlaves.sh", "--MasterSlave=${MasterEP}"]  
      name: killApp
      image: my-utils
      command: ["/bin/sh", "-C", "./killAndWait.sh", "--name=${POD_NAME}"] 
      ...
```
### Shutdown Sequence
If we attempt to shutdown a runing Oracle instance which has four stages https://docs.oracle.com/cd/B28359_01/backup.111/b28273/rcmsynta045.htm#RCMRF155
```yaml
pod:
  spec:
    initContainers: ...
    containers: ...
    deferContainers:
#Pre-Exit
    - name: shutdown-db
      image: db-utils
      #Select the shutdown sequnce 
      comamnd: ["/bin/sh", "-C", "./shutDown.sh", "--shutdowntype=${SH_TYPE}"] 
      name: waitforDB
      image: db-utils
      #run a script that will wait until the DB is down.
      command: ["/bin/sh", "-C", "./waitForDB.sh"]  
      ...
```

## Kubelet Changes
* The images are pre-pulled in SyncPod() when pod phase is ‘Running’ as a step 7
* killPodWithSyncResult() is blocking, deferContainers execution is implemented inside this function.  
* We needed killPod or killPodWithSyncResults to get access to pullSecrets and podStatus, these two have been propagated
* A new method for ContainerManager interface added WaitForContainer (containerID string) error  so that we could start a container and block on it during termination.

A simple pseudo code implementation
```go
func killContainersWithSyncResult() {
	runDeferContainers()

	for _, container := range runningPod.Containers {

		//if deferContainers are configured skip preStopHook()
		go killContainer(pod, container.ID, container.Name)

	}
	//Wait for all the container to be killed
}
```
And runDeferContainers will be implemented as 
```go
func runDeferContainers(){
	for _, container := range pod.Spec.DeferContainers {

		m.startContainerAndWait(podSandboxID,podSandboxConfig, container)
	 
//Wait for container to finish or time.After(GracePeriod)

      }
}
```
Sync POD has a new phase 7 to pre-pull deferContainer images if not available
```go
func SyncPod() {

	// Step 1: Compute sandbox and container changes.
	
	// Step 2: Kill the pod if the sandbox has changed.

	// Step 3: kill any running containers in this pod which are not to keep.

	// Step 4: Create a sandbox for the pod if necessary.

	// Step 5: start init containers.

	// Step 6: start containers in podContainerChanges.ContainersToStart.

	//Step 7: If the Pods is in running phase pre-populate deferContainer images

	if pod.Status.Phase == v1.PodRunning {
	
		pre-pullDeferContainerImage()
	
	}
}
```

## Caviate
This Design primarily focuses on handling graceful termination cases.  If the Node running a deferContainer configured pod
crashes abruptly, then this design does not guarantee that cleanup was performed gracefully.  This still requires community
feedback on how such scenarios are handled and how important it is for deferContainers to handle that situation. 

## Reference
[Community Request](https://github.com/kubernetes/kubernetes/issues/35183)
[WIP PR](https://github.com/kubernetes/kubernetes/pull/47422)
[UseCase Sides](https://docs.google.com/presentation/d/12WEEWQh8ffiLyqh8F60PgRvQn3mfdC2rx3E8biZm3oM/edit?usp=sharing)
