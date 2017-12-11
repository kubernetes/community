# PV monitoring proposal

Status: Pending

Version: Alpha

Implementation Owner: NickrenREN@ 

## Motivation

For now, kubernetes has no way to monitor the PVs, which may cause serious problems. 
For example: if volumes are unhealthy, and pods do not know that and still try to get and write data, 
which will lead to data loss and unavailability of services. 
So it is necessary to have a mechanism for monitoring PVs and react when PVs have problems.

## Proposal

We can separate the proposal into two parts:

* monitoring PVs and marking them if they have problems
* reacting to the unhealthy PVs

For monitoring, we may create a controller for it, and each volume plugin should have its own function to check volume health. 
Controller can call them periodically. The controller also needs to watch node events because local PVs will be unreachable if nodes break down.

For reacting, different kinds of apps may have different methods,we can also create a controller for it. 

At first phase, we can focus on local storage PVs monitoring.

## User Experience
### Use Cases

* If the local PV path is deleted, users should know that and the local PV should be marked and deleted;
* If the local PV path is not a mountpoint any more, the local PV should be marked and deleted;
* If nodes which have local PVs are breaking down, the local PVs should be marked and deleted (the application has data backup and can restore it or can tolerate data loss and the PV protection feature may help);
* For local PVs, we need to make sure that PV capacity must not be greater than device capacity and PV used bytes must not be greater than PV capacity;
* For network storage, if the storage driver volume is deleted, the PV object in kubernetes should be marked and deleted too;
* If we can not get access to the PV volume for a certain time (network or some other problems), we need to mark and delete the PV;
* PV fsType checking ? bad blocks checking ?

## Implementation

As mentioned above, we can split this into two parts and put them in the external repo at first.

### Monitoring controller: 

Like PV controller, monitoring controller should check PVs’ health condition periodically and taint them if PVs are unhealthy.

Health checking implementation should be per plugin. Each volume plugin needs to have its own methods to check its volumes.

At the first stage, we can focus on local storage PVs, and then extend to other network storage PVs.
#### For local storage:

The local storage PV monitor consists of two parts

* create a daemonset on every node, which is responsible for monitoring local PVs in that specific node, no matter the PVs are created manually or by provisioner;
* create a monitor controller, which is responsible for watching PVs and Nodes events. PVs may be updated if they are unhealthy and we also need to react to node failure event.

At the first phase, we can support local storage monitoring first.

Take local storage as an example, detailed checking method may be like this:

```
// checkStatus checks local pv health condition
func (monitor *LocalPVMonitor) checkStatus(pv *v1.PersistentVolume) {
    // check if PV is local storage
	if pv.Spec.Local == nil {
		glog.Infof("PV: %s is not local storage", pv.Name)
		return
	}
	// check node and pv affinity
	fit, err := CheckNodeAffinity(pv, monitor.Node.Labels)
	if err != nil {
		glog.Errorf("check node affinity error: %v", err)
		return
	}
	if !fit {
		glog.Errorf("pv: %s does not belong to this node: %s", pv.Name, monitor.Node.Name)
		return
	}

	// check if host dir still exists
	mountPath, continueThisCheck := monitor.checkHostDir(pv)
	if !continueThisCheck {
		glog.Errorf("Host dir is modified, PV should be marked")
		return
	}

	// check if it is still a mount point
	continueThisCheck = monitor.checkMountPoint(mountPath, pv)
	if !continueThisCheck {
		glog.Errorf("Retrieving mount points error or %s is not a mount point any more", mountPath)
		return
	}

	// check PV size: PV capacity must not be greater than device capacity and PV used bytes must not be greater that PV capacity
	if pv.Spec.VolumeMode != nil && *pv.Spec.VolumeMode == v1.PersistentVolumeBlock {
		monitor.checkPVAndBlockSize(mountPath, pv)
	} else {
		monitor.checkPVAndFSSize(mountPath, pv)
	}

    // other checks ...
}
```
If monitor finds that one PV is unhealthy, it will mark the PV by adding annotations including timestamp. 
The reaction controller then can react to this PV depending on the annotations and timestamp.

When we first mark a PV, we will add another annotation which key is `FirstMarkTime`. 
And if local PV is unhealthy, annotation keys may be like: `HostPathNotExist`, `MisMatchedVolSize`, and `NotMountPoint`...

A marked local PV looks like:
```
Name:            example-local-pv-1
Labels:          <none>
Annotations:     FirstMarkTime=2018-04-17 07:31:02.388570492 +0000 UTC m=+600.033905921
                 HostPathNotExist=yes
                 NotMountPoint=yes
                 volume.alpha.kubernetes.io/node-affinity={ "requiredDuringSchedulingIgnoredDuringExecution": { "nodeSelectorTerms": [ { "matchExpressions": [ { "key": "kubernetes.io/hostname", "operator": "In", "valu...
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    local-disks
Status:          Available
Claim:
Reclaim Policy:  Retain
Access Modes:    RWO
Capacity:        200Mi
Node Affinity:   <none>
Message:
Source:
    Type:  LocalVolume (a persistent volume backed by local storage on a node)
    Path:  /mnt/disks/vol/vol1
Events:
  Type    Reason           Age   From                                                                 Message
  ----    ------           ----  ----                                                                 -------
  Normal  MarkPVSucceeded  1m    local-volume-monitor-127.0.0.1-40a8fb4d-4206-11e8-8e52-080027765304  Mark PV successfully with annotation key: NotMountPoint
  Normal  MarkPVSucceeded  22s   local-volume-monitor-127.0.0.1-40a8fb4d-4206-11e8-8e52-080027765304  Mark PV successfully with annotation key: HostPathNotExist
```

#### For out-tree volume plugins(except local storage):

We can implement the monitor at external-repo at first. So for networked storage monitor, 
we can create a new controller called MonitorController like ProvisionController, 
which is responsible for creating informers, watching Node and PV events and calling each plugin’s monitor functions and watch . 
And each volume plugin will create its own monitor to check its volumes’ status.

#### For in-tree volume plugins(except local storage):

We can add a new volume plugin interface: PVHealthCheckingVolumePlugin.
```
  type PVHealthCheckingVolumePlugin interface {
       VolumePlugin

      CheckHealthCondition(spec *Spec) (string, error)
  }
```
And each volume plugin will implement it. The entire monitoring controller workflow is:

* Fill PV cache with initial data from etcd
* Resync and check volumes status periodically
* Taint PV if the volume status is abnormal

### PV controller changes:
For unbound PVCs/PVs,  PVCs will not be bound to PVs which have taints.

### Reaction controller:
Reaction part can be implemented at the second stage, and can focus on statefulset reaction at first.
Reaction controller will react to the PV update event (PVs tainted/marked by monitoring controller). 
Different kinds of apps should have different reactions.

statefulset reaction: check the annotation timestamp, if the PV can recover within the predefined time interval, 
we will do nothing, otherwise we need to delete the PVC bound to the unhealthy volume(PV) as well as pods referencing it. 
Notice the statefulset apps must  have data backup and can restore it or can tolerate data loss. 
The PV protection feature may help.

Reaction controller’s workflow is:

* Fill PV cache from etcd;
* Watch for PV update events;
* Resync and populate periodically;
* Delete related PVC and pods if needed ;



## Roadmap to support PV monitoring

* support local storage PV monitoring(marking PVs);
* out-tree networked volume plugins monitor and statefulset reaction and add PV taint API support;
* support in-tree volume plugins and react to other kinds of applications if needed.

## Alternatives considered
