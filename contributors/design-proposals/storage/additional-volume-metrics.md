# Volume operation metrics

## Goal

Capture additional metrics that expose internal state of Storage Sub System in Kubernetes.
These metrics can be useful for debugging as well as for gauging health of the cluster.

## Motivation

In Kubernetes 1.8 and 1.9 we added metrics for CloudProvider API calls and Volume operations.
While those metrics are useful, this proposal specifically adds metrics that expose
internal state of Kubernetes's Storage Sub system.

These metrics will be very useful for debugging and determining health of the cluster.

## List of metrics

### Total mount and umount time metrics

Total mount and unmount time taken by the volume with volume_plugin as dimension. This time measurement starts from the time when volume manager gets the mount request (when the pod add event is observed by the volume manager) to the time when the mount is succeeded. This is different from currently emitted volume operation metrics because that metric only includes time for “mount” or “umount” operation itself to happen.

### Total Attach and Detach Time metrics

Similar to above metric. This is total time take by volume to get attached or detached once Attach Detach Controllers sees the request.  In a nutshell - this is total time taken for volume from entering desired state of world (when attach/detach request is received) to entering actual state of world (when attach/detach operation finishes) .

This is different from currently emitted operation metrics because that metric only includes time taken by attach or detach operation once reconciler has already decided that operation can proceed.

Again volume_plugin will be emitted as dimension/label.

### Total Provision Time metrics

Similar to above metric with volume_plugin and storageClass as possible dimensions/labels.

### Total PV Deletion Time metrics

Similar to above metric with volume_plugin and storageClass as possible dimensions/labels.


### Number of volumes in ActualStateofWorld of A/D Controller

This is number of volumes that Attach Detach controller believes to be successfully attached.

### Number of Volumes in desiredStateOfWorld of A/D Controller

This is number of volumes that Attach Detach Controller believes that require attaching. This number will count unique volumes but if a volume needs to be mounted on more than one node, it will count them as different volumes.

### Number of bound PVCs
Number of known bound PVCs with namespace as dimension.

### Number of unbound PVCs
Number of unbound PVCs with namespace as dimension.

### Number of bound PVs
Number of bound PVs

### Number of unbound PVs
Number of unbound PVs in cluster

### Number of PVCs in use by pods
This metric will report number of PVCs that are in use by pods. It will include both pending and running pods but exclude terminated pods. It will report this number with volume plugin name as dimension.

### Number of times A/D Controller performs force detach
This metric will report number of times Attach Detach controller performs force detaches because 6 minute time period has expired after pod is removed from desired state of world.

### Number of Volume in ActualStateofWorld of VolumeManager
This metric will report number of volumes that kubelet believes is successfully attached and mounted to the node. Again metric will report volume plugin as a dimension/label.

### Number of Volume in DesiredStateofWorld of VolumeManager
This metric will report number of volumes that should be attached and mounted to the node. Metric will report volume plugin as dimension/label.

### Number of times ReconstructVolumeSpec on kubelet failed
This metric will report number of times kubelet failed to reconstruct volume spec by reading from disc. It will report volume plugin as dimension.

### Number of times PV Provisioning failed
As evident from name with plugin name and storageclass as dimension.

### Number of times PV Deletion failed
As evident from name with plugin name and storageclass as dimension.

### Metric of Volumes that fail to reach “attached” state after seemingly successful attach call
Sometimes even if “attach” call of a volume plugin succeeds, the volume may fail to reach desired state within expiry time window. This metric reports such
numbers, but implmentation will be plugin dependent.

### Metric of Volumes that fail to reach “detached” state after seemingly successful detach call

Same as above.

### Volume formatting time (mkfs)
Time it takes to format a new volume with volume-plugin name and fstype as dimension.

### Number of Orphaned Pod Directories

Number of observed orphaned pod directories on node. A high number usually means, more work for kubelet and people have even reported timeouts while starting pods.

Will possibly emit plugin name as dimension if available.

### Number of time adding/deleting pod to A/D Controller fails because node is unknown
Attaching a volume to pod may fail in time because a known race condition in pod and node addition (but don’t worry we have a fallback mechanism).

This metric tracks number of times a volume doesn’t get attached in time because pod addition happened before node could be processed.


### Number of times attach failed because it was still attached to another node
As evident from name. With plugin name as possible dimension.

### Number of times PVC/PV binding API update failed
Sometimes pv controller has to retry binding because PVC/PV binding API update may fail on first try. This metric will keep track of failed attempts.

### Number of times VerifyVolumesAreAttached mechanism fixes volumes
Sometimes a volume may attach/detach for seemingly reasons outside k8s control. This check makes sure such volumes are attached or detached as expected.

This metric keeps track of such events with plugin name as possible dimension.

### Number of times API requests is getting throttled

It will be useful to know when Storage API requests to external cloudprovider are being throttled and preferably some kind of timing information for it.
