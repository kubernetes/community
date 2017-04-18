# Defer Containers

dhilip.kumar.s@huawei.com

March 2017

## Prerequisite
This is a continuation of [initContainers](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/container-init.md)

## Motivation
Since the introduction to Statefulset and Daemonset PODs that consist of stateful containers are now fully supported by
Kubernetes.  Such stateful pods benefit from constructor and destructor semantics. InitContainers provide constructor
semantics. The proposed DeferContainers will provide destructor semantics for such stateful PODs

## Abstract
This introduces the concept of ‘deferContainers’, inspired by golang’s keyword ‘defer’. This will allow one to define a
set of Containers to be executed at the time of POD termination. The defined containers will be executed sequentially and
in the specified order. It will bring destructor() capability to Pods.

## Use Cases

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
* In future when kuberenetes supports Virtual Machine runtime (eg: hyperContainer) for better isolation, we should shutdown
  the VMs instead of killing them abruptly.  ‘deferContainers’ could help us run such shutdwon commands.

## Limitation with the current system
Container pre-stop hooks are not sufficient for all termination cases:
* They cannot easily coordinate complex conditions across containers in multi-container pods
* They can only function with code in the image or code in a shared volume, which would have to be statically linked 
  (not a common pattern in wide use)

## Design Requirements
Most of the requirements are very similar to initContainers.  They are replicated and modified as necessary.
* deferContainers should be able to:
  * Use the same volume (PV) as appContainers such that it can
    * Perform cleanup of shared volume, such as delete several temp directories.
    * Delete unwanted files before a final sync is initiated.
    * Update Configuration files about the changes in the distributed system so that the next pod getting attached to this PV 
	  will benefit from it. (like a new leader/master etc).
    * Deleted secrets or security related files before the pod de-couples from the PV.
  * Delay the termination of application containers until operations are complete
  * De-Register the pod with other components of the system

* Reduce coupling:
  * Between application images, eliminating the need to customize those images for Kubernetes generally or specific roles
  * Inside of images, by specializing which containers perform which tasks (install git into init container, use filesystem 
    contents in web container)
  * Between termination steps, by supporting multiple sequential cleanup containers
* Pre-Exit and Post-Exit workflow
  * Post - specify that the controller can continue to delete the appContainers but wait for deferContainer’s to complete 
    its execution before marking as the Pod is deleted.
  * Pre - specify that the controller should wait until deferContainers completes its execution first and then proceed to 
    delete the appContainers if needed
* deferContainers should allow us to specify
  * if certain containers in deferContainer list need to be re-started on failures.
* GracePeriod
  * It should be possible to mention overall terminationGracePeriod for defercontainers.
* Run-once and run-forever pods should be able to use deferContainers
* Reduce Complexity, 
  * It should be possible to use a generic container as a deferContainer,
  * A deferContainer should be independently invokable, ie:- should not require code in the same image as the appContainers.

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
Most of the design elements are similar to initContainers such as below
* Will have 0...N containers and will be executed in sequence (specified order)
* Restart policy could be either ‘restartAlways’ or ‘restartNever’
* Adding a new phase ‘Terminating’ in the Pod lifecycle

### Terminating Phase / Defer Phase
* A POD reaches terminating phase when it is about to be removed from Kubernetes Cluster
* During this phase, the appContainers are not restarted if they get terminated/killed
* deferContainers will be executed one-after-the-other in the same sequence as they were specified in the POD Spec. From the above
  pod spec example execution sequence will be defer-Container1, then defer-Container2, …, etc.,
* we will either move-on from a failed deferContainer or restart it depending upon its restart policy, the default behavior will be 
  restartNever
* if the user specifies `kubectl delete pod foo --grace-period=0 --force` to delete a pod deferContainers will not be executed.
```
Example status output when a pod is being terminated.
   NAME      READY     STATUS     RESTARTS   AGE
    foo-0     2/2       Defer:0/4   0          7m
```
* Failure of one or all deferContainers will not trigger a POD restart.

### TerminationGracePeriod
* It takes default value (30 seconds as of today), explicitly mentioning this flag overrides the default value.
* To retain backward compatibility PreStopHooks will be started (if configured) for all the containers. 
* Then deferContainer will start to execute one after the other (without waiting for preStop hooks to complete)
* Currently when a Pod (PreStopHook) did not finish within the given GracePeriod, then kubelet will provide an extension of 2 seconds, we should retain that property for deferContainers too.
* When the configured GracePeriod expires (with the additional 2 second graceperiod), it will kill all the running app containers (if not deleted already)
* It will kill currently executing deferContainer and no further deferContainer will be executed (if there are any). 
* deferContainers are timebound by TerminationGracePeriod 

For those termination scenarios where running time of a deferContainers is not easy to predict ahead of time, Such as filecopy or fileupload which depends on the disk speed and internet bandwidth respectively.  In the future depending on the communities feedback we can consider below possible approaches. This will provide slightly elastic terminationGracePeriod mechanism.

#### Solution 1 `deferContainerGracePeriodExtension: True` (a possible future improvement to deferContainer)

we could introduce a new flag in the podSpec `deferContainerGracePeriodExtension`.  

This will allow deferContainers to run beyond the TerminationGracePeriod if liveliness probe is  configured for those containers.
After TerminationGraceperiod expires every time liveliness probe succeeds the termination grace period will extend to ‘Probe.PeriodSeconds`.

```
  60s            60s              1       kubelet, kube-node-3    spec.initContainers{initialization}     Normal          Pulling                  pulling image "initConteiner"
  55s            55s              1       kubelet, kube-node-3    spec.initContainers{initialization}     Normal          Started                  Started container with id 8627e1fa6df3be2b2ff976e5ef46bb06dd768fb06e938b27c3171cf1e0c79932
  50s            50s              1       kubelet, kube-node-3    spec.containers{AppContainer}           Normal          Started                  Started container with id a3ff2f84a7ee9a1d8d0aebae9b69096eb189d3390f509506351c1e9a987a0674
  45s            45s              1       kubelet, kube-node-3    spec.deferContainers{Termination}       Normal          Pulling                  pulling image "cleanup-container"
  40s            40s              1       kubelet, kube-node-3    spec.deferContainers{Termination}       Normal          Started                  Started container with id 8627e1fa6df3be2b2ff976e5ef46bb06dd768fb06e938b27c3171cf1e0c79933
  10s            10s              1       kubelet, kube-node-3    spec.deferContainers{Termination}       Normal          extendGracepriod         deferContainer still running gracePeriodExtend for 2 more seconds
  8s              8s              1       kubelet, kube-node-3    spec.deferContainers{Termination}       Normal          extendGracepriod         deferContainer still running gracePeriodExtend for 2 more seconds
  6s              6s              1       kubelet, kube-node-3    spec.deferContainers{Termination}       Normal          extendGracepriod         deferContainer still running gracePeriodExtend for 2 more seconds
```

If this flag is set then its is recommended that all the deferContainers are configured with livelinessProbe for predictable behaviour.

#### Solution 2 (a possible future improvement to deferContainers)
We could introduce two new flags `deferContainerGracePeriodExtensionIntervel: S seconds` and `deferContainerGracePeriodExtensionCount: N Integer` in the pod spec, such that post TerminationGracePeriod we could keep extending the graceperiod every 'S' seconds and retry that for 'N' times.  This will be a common for all deferContainers and will be simpler to implement for Pod Designers.

Either cases to delete the pod instantly `--force` && `--grace-period=0` should be supplied. 

### PerPopulate Heavy deferContainers
Eventhough it is extreemly rare for someone to actually use a heavy image for a deferContainer, there might be some user scenarios for this.  Pulling such heavy images might introduced unexpteced delay in the termination sequence.  A new flag will be introduced `prePullDeferImages: true` in podspec that will instruct kubelet to pull all the deferContainers images once the Pod reaches 'running' state.
### Pre/Post Termination triggers
All the deferContainers will behave like a preExit trigger, it should be easier program deferContainers in such a way that it does both 
pre and post Termination tasks.
```yaml
pod:
  spec:
    terminationGracePeriod: 60
    initContainers: ...
    containers: ...
    deferContainers:
#pre exit operations
    - name: remove-shard
      image: dbUtils
      command: ["/bin/sh", "-C", "remove-shard", "--name=${POD_NAME}"]
    - name: wait-for-rebalance
      image: dbUtils
      command: ["/bin/sh", "-C", "while[[ 1 ]];do sleep 1; if key_rebalance_complete.sh; then exit 0; fi; done"]
    - name: kill-db
      image: dbUtils
      command: ["/bin/sh", "-C", "shutdown-db", "--name=${POD_NAME}"]
#post exit operation
    - name: clean-update
      image: dbUtils
      command: ["/bin/sh", "-C", "./disk_cleanup.sh"]
      ...
```
### Short running Job / Pod
For PODs which run and exit gracefully themselves if deferContainers are configured they will act as PostExit triggers, 
an internal flag should indicate if  deferContainers have already been called or not.

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
#contact and inform a thirdparty system about this pods termination, such as reducing a refernce counter (if one is maintained)
      name: ref-counter
      image: my-utils
      comamnd: ["/bin/sh", "-C", "./decrementRefCount.sh"]
      ...
```

### Master-slave / Leader-follower statefulset down-size / scale down the replicas
Below scripts 'selectaMaster.sh and reConfSlaves.sh should be designed in such a way that even if a terminating pod is a slave it
shoudnt affect the cluster. This will fit controllers such as Statefulset because they gaurentee only one pod goes down at once.
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

## Caviate
This Design primarily focuses on handling graceful termination cases.  If the Node running a deferContainer configured pod
crashes abruptly, then this design does not guarantee that cleanup was performed gracefully.  This still requires community
feedback on how such scenarios are handled and how important it is for deferContainers to handle that situation. 

## Reference
[Community Request](https://github.com/kubernetes/kubernetes/issues/35183)

[Places for hooks](https://github.com/kubernetes/kubernetes/issues/35183)
