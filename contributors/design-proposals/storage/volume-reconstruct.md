# New Design Volume Reconstruction

**Author**: Jing Xu (@jingxu97)

**Last Updated**: 1/11/2018

**Status**: Proposal

This document describes the new design of volume reconstruction in Kubelet
volume manager running on cluster nodes.

## Background and Motivation

Today, kubelet volume manager uses desired/actual states to store some volume
and pod information for reconciler to work upon. When kubelet restarts, all
states become empty. Desired state populator will recover the desired states by
getting the existing pods information from pod manager. Reconciler can then
recover actual state through reconcile loop. 

The problem is that if pod is deleted during kubelet restarts, desired state
no longer has this pod and volume information so as the actual state. But the
actual mount points and direcotries still exist on the disk. In order to
clean up those left over mounts and directories, we propose a volume
reconstruction process. The basic idea is to scan pod directories and recover
volume information (such as volume spec) from the mount path and device mount
path which will be updated into the actual state. With this reconstructure
process, actual state can match the real world and reconciler will tear down
volumes accordingly. 

However, this reconstruction design has issues for some plugins such as 
iscsi, ceph rbd and fc, because there is not enough information to reconstruct volume spec correctly.
For those volume plugins, volume tear down will fail.

## Possible solution

In searching of the solutions, we first need to understand why we need to reconstruct volume spec. 
Volume spec is an internal representation of a volume.  All API volume types translate to Spec.
```
type Spec struct {
	Volume           *v1.Volume                // for inline volume use case
	PersistentVolume *v1.PersistentVolume      // for pvc use case
	ReadOnly         bool
}
```
Both Volume and PersistentVolume has a name and volume source. Different volume plugins have different representation of their sources. 
Volume spec is constructed when adding pod volume information into desired state.
The main use of this spec is to get the volume source name to construct unique volume name (which typically includes plugin name and volume source name). 
The unique volume name is used as a unique identification/key in actual/desired state of volume controller. The volume source name is also used on construct global mount path (device mount path)
for volume mounter. When volume is torn down, unmounter uses the same function to construct the mount path from the spec again.
Different from global mount path, the bind mount path uses plugin name and spec name (not the volume source name). 

Use a GCE PD as an example,
```
Volume{
	Name: "my-volume",
	VolumeSource: core.VolumeSource{
		GCEPersistentDisk: &core.GCEPersistentDiskVolumeSource{
					PDName:    "my-pd",
					FSType:    "ext4",
					Partition: 1,
					ReadOnly:  false,
		}
	}
}
```

* volume source name: my-pd
* volume spec name: my-volume
* Unique volume name: kubernetes.io/gce-pd/my-pd
* device mount path: kubernetes.io/gce-pd/mounts/my-pd
* bind mount path (also pod volume dir): pods/podUID/volumes/kubernetes.io~gce-pd/my-volume

So look more closely, you can see for only the purpose of volume tear down, if we have a record of those mount paths (bind and global) in actual state, there is no need to 
construct those paths from the volume spec again. In case of kubelet restarts, those mount paths could be scanned from the mount lists left on the disk. (bind mount path has
a path structure which can be easily idenetified and global mount path could be obtained from the mount reference of the bind mount path). This observation leads to our first
possible solution.

### Solution1: Remove the dependency of volume spec. 
Modify the data structure in actual state to keep a record of the mount paths so that volume tear down process (unmount) can use those records instead of reconstructing them
directly from volume spec. In case the kubelet restarts, reconstruction process just need to find out the left mounting paths and mark them into the actual state. Even if the
volume spec cannot be reconstructed correctly, volume tear down process can still work. 

Issue: Because some plugins cannot reconstruct volume spec correctly, which means the generated unique volume name might not match with the real one.
In case another new pod is added to use the same volume, two volume operations on the same device might happen concurrently and cause race condition because they have different
unique volume names.

### Solution2: Cleanup mounts directly without reconstruction.
This approach will try to clean up mounts that are no longer needed before reconciler starts to work. It scans the pod volume directories when kubelet restarts to find the mount
paths (bind and global). If the pod volume no longer exist in desired state, clean up the mounts. To avoid long delay or mount hanging problem, we can use operation executor to
start a new routine. 

Issue: The potential issue of this approach is that cleanup process might fail and
reconciler will not try to do this process again.

## Proposal

The current proposal is trying to take the advantages from the two possible
solutions mentioned above. 
1. If volume plugins can support reconstruction well (i.e., volume spec can
   be fully recovered), reconciler will update this information into the states
   and reconciler will take care of volume cleanup. Reconciler can retry if
   cleanup fails.

Currently, the following volume plugins can support reconstruction. These volume
source names are just the device names which can be retrived from the mount paths.

AWS EBS, Azure_dd, Cinder, CSI, GCE_PD, FC, iScsi, Photon, Rbd, Vsphere 

2. If volume plugins cannot support reconstruction, reconciler clean up the
   mounts directly with a go routine (operation executor) to avoid delay. If it
   fails, reconciler will not retry.

The following volume plugins cannot support reconstruction,
Azure_file, Cephfs, ConfigMap, DownwardAPI, EmptyDir, flexvolume, flocker,Git_repo, GlusterFs, Host_path, nfs, portwox, projected, quobyte, scaleio, secrete, storageOs

## Implementation Timeline:
The work is targeted for kubernetes v1.10

