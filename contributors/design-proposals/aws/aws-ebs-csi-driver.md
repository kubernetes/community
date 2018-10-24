# AWS EBS CSI Driver
## Problems with current in-tree cloud provider
### Cache of used / free device names

On AWS, it&#39;s the client who [must assign device names](https://aws.amazon.com/premiumsupport/knowledge-center/ebs-stuck-attaching/) to volumes when calling AWS.AttachVolume. At the same time, AWS [imposes some restrictions on the device names](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html).

Therefore Kubernetes AWS volume plugin maintains cache of used / free device names for each node. This cache is lost when controller-manager process restarts. We try to populate the cache during startup, however there are some corner cases when this fails. TODO: exact flow how we can get wrong cache.

It would be great if either AWS itself assigned the device names, or there would be robust way how to restore the cache after restart, e.g. using some persistent database. Kubernetes should not care about the device names at all.

### DescribeVolumes quota

In order to attach/detach volume to/from a node, current AWS cloud provider issues AWS.AttachVolume/DetachVolume call and then it polls DescribeVolume until the volume is attached or detached. The frequency of DescribeVolume is quite high to minimize delay between AWS finishing attachment of the volume and Kubernetes discovering that. Sometimes we even hit API quota for these calls.

It would be better if CSI driver could get reliable and fast event from AWS when a volume has become attached / detached.

Or the driver could batch the calls and issue one big DescribeVolume call with every volume that&#39;s being attached/detached in it.

### AWS API weirdness

AWS API is quite different to all other clouds.

- AWS.AttachVolume/DetachVolume can take ages (days, weeks) to complete. For example, when Kubernetes tries to detach a volume that&#39;s still mounted on a node, it will be Detaching until the volume is unmounted or Kubernetes issues force-detach. All other clouds return sane error relatively quickly (e.g. &quot;volume is in use&quot;) or force-detach the volume.
- AWS.DetachVolume with force-detach is useless. Documentation says: Forced detachment of a stuck volume can cause damage to the file system or the data it contains or an inability to attach a new volume using the same device name, unless you reboot the instance.
  - We cannot reboot instance after each force-detach nor we can afford to &quot;loose&quot; a device name. AWS supports only 40 volumes per node and even that is quite low number already.
- AWS.CreateVolume is not idempotent. There is no way to create a volume with ID provided by user. Such call would then fail when such volume already exists.
- AWS.CreateVolume does not return errors when creating an encrypted volume using either non-existing or non-available KMS key (e.g. with wrong permission). It returns success instead and it even returns volumeID of some volume! This volume exists for a short while and it&#39;s deleted in couple of seconds.

### Errors with slow kubelet

Very rarely a node gets too busy and kubelet starves for CPU. It does not unmount a volume when it should and Kubernetes initiates detach of the volume.

## Requirements

### Idempotency

All CSI driver calls should be idempotent. A CSI method call with the same parameters must always return the same result. It&#39;s task of CSI driver to ensure that. Examples:

- CreateVolume call must first check that the requested EBS volume has been already provisioned and return it if so. It should create a new volume only when such volume does not exist.
- ControllerPublish (=i.e. attach) does not do anything and returns &quot;success&quot; when given volume is already attached to requested node.
- DeleteVolume does not do anything and returns success when given volume is already deleted (i.e. it does not exist, we don&#39;t need to check that it had existed and someone really deleted it)

Note that it&#39;s task of the CSI driver to make these calls idempotent if related AWS API call is not.

### Timeouts

gRPC always passes a timeout together with a request. After this timeout, the gRPC client call actually returns. The server (=CSI driver) can continue processing the call and finish the operation, however it has no means how to inform the client about the result.

Kubernetes will retry failed calls, usually after some exponential backoff. Kubernetes heavily relies on idempotency here - i.e. when the driver finished an operation after the client timed out, the driver will get the same call again and it should return success/error based on success/failure of the previous operation.

Example:

1. Kubernetes calls ControllerPublishVolume(vol1, nodeA) ), i.e. &quot;attach vol1 to nodeA&quot;.
2. The CSI driver checks vol1 and sees it&#39;s not attached to nodeA yet. It calls AttachVolume(vol1, nodeA).
3. The attachment takes a long time, Kubernetes times out.
4. Kubernetes sleeps for some time.
5. AWS finishes attaching of the volume.
6. Kubernetes re-issues ControllerPublishVolume(vol1, nodeA) again.
7. The CSI driver checks vol1 and sees it is attached to nodeA and returns success immediately.

Note that there are some issues:

- Kubernetes can change its mind at any time. E.g. a user that wanted to run a pod on the node in the example got impatient so he deleted the pod at step 4. In this case Kubernetes will call ControllerUnpublishVolume(vol1, nodeA) to &quot;cancel&quot; the attachment request. It&#39;s up to the driver to do the right thing - e.g. wait until the volume is attached and then issue detach() and wait until the volume is detached and \*then\* return from
- Note that Kubernetes may time out waiting for ControllerUnpublishVolume too. In this case, it will keep calling it until it gets confirmation from the driver that the volume has been detached (i.e. until the driver returns either success or non-timeout error) or it needs the volume attached to the node again (and it will call ControllerPublishVolume in that case).
- The same applies to NodeStage and NodePublish calls (&quot;mount device, mount volume&quot;). These are typically much faster than attach/detach, still they must be idempotent when it comes to timeouts.

It looks complicated, but it should be actually simple - always check that if the required operation has been already done

### Restarts

The CSI driver should survive its own crashes or reboots of the node where it runs. For the controller service, Kubernetes will either start a new driver on a different node or re-elect a new leader of stand-by drivers. For the node service, Kubernetes will start a new driver shortly.

The perfect CSI driver should be stateless. After start, it should recover its state by observing the actual status of AWS (i.e. describe instances / volumes). Current cloud provider follows this approach, however there are some corner cases around restarts when Kubernetes can try to attach two volumes to the same device on a node.

When the stateless driver is not possible, it can use some persistent storage outside of the driver. Since the driver should support multiple Container Orchestrators (like Mesos), it must not use Kubernetes APIs. It should use AWS APIs instead to persist its state if needed (like AWS DynamoDB). We assume that costs of using such db will be negligible compared to rest of Kubernetes.

### No credentials on nodes

General security requirements we follow in Kubernetes is &quot;if a node gets compromised then the damage is limited to the node&quot;. Paranoid people typically dedicate handful of nodes in Kubernetes cluster as &quot;infrastructure nodes&quot; and dedicate these nodes to run &quot;infrastructure pods&quot; only. Regular users can&#39;t run their pods there. CSI attacher and provisioner is an example of such &quot;infrastructure pod&quot; - it need permission to create/delete any PV in Kubernetes and CSI driver running there needs credentials to create/delete volumes in AWS.

There should be a way how to run the CSI driver (=container) in &quot;node mode&quot; only. Such driver would then respond only to node service RPCs and it would not have any credentials to AWS (or very limited credentials, e.g. only to Describe things). Paranoid people would deploy CSI driver in &quot;node only&quot; mode on all nodes where Kubernetes runs user containers.

## High level overview of CSI calls

### Identity Service RPC

#### GetPluginInfo

Blindly return:

```
  Name: ebs.csi.aws.com 
  VendorVersion: 0.1.0-alpha 
```

#### GetPluginCapabilities

Blindly return:

```
   Capabilities:
     - CONTROLLER_SERVICE
     - ACCESSIBILITY_CONSTRAINTS
```

#### Probe

- Check that the driver is configured and it can do simple AWS operations, e.g. describe volumes or so.
- This call is used by Kubernetes liveness probe to check that the driver is healthy. It&#39;s called every ~10 seconds, so it should not do anything &quot;expensive&quot; or time consuming.  (10 seconds are configurable, we can recommend higher values).

### Controller Service RPC

#### CreateVolume

Checks that the requested volume was not created yet and creates it.

- Idempotency: several calls with the same name parameter must return the same volume. We can store this name in volume tags in case the driver crashes after CreateVolume call and before returning a response. In other words:
  - The driver first looks for an existing volume with tag CSIVolumeName=<name>. It returns it if it&#39;s found.
  - When such volume is not found, it calls CreateVolume() to create the required volume with tag CSIVolumeName=<name>
  - _Is this robust enough? Can this happen on AWS?_
    1. A driver calls CreateVolume() and dies before the new volume is created.
    2. New driver quickly starts, gets the same CreateVolume call, checks that there is no volume with given tag (previous CreateVolume() from step 1. has not finished yet) and issues a new CreateVolume().
    3. Both AWS.CreateVolume() calls succeed -> the driver has provisioned 2 volumes for one driver.CreateVolume call.
- Snapshot: if creating volume from snapshot, read the snapshot ID from request.

#### DeleteVolume

Checks if the required volume exists and is &quot;available&quot; (not attached anywhere) and deletes it if so. Returns success if the volume can&#39;t be found. Returns error if the volume is attached anywhere.

#### ControllerPublishVolume

- Checks that given volume is already attached to given node. Returns success if so.
- Checks that given volume is available (i.e. not attached to any other node) and returns error if it is attached.
- Chooses the right device name for the volume on the node (more on that below) and issues AttachVolume. TODO: this has complicated idempotency expectations. It cancels previously called ControllerUnpublishVolume that may be still in progress (i.e. AWS is still detaching the volume and Kubernetes now wants the volume to be attached back).

#### ControllerUnpublishVolume

Checks that given volume is not attached to given node. Returns success if so. Issues AWS.DetachVolume and marks the detached device name as free (more on that below). TODO: this has complicated idempotency expectations. It cancels previously called ControllerPublishVolume (i.e.AWS is still attaching the volume and Kubernetes now wants the volume to be detached).

#### ValidateVolumeCapabilities

Check whether access mode is supported for each capability

#### ListVolumes

Not implemented in the initial release, Kubernetes does not need it.

#### GetCapacity

Not implemented in the initial release, Kubernetes does not need it.

#### ControllerGetCapabilities

Blindly return:

```
  rpc:
    - CREATE\_DELETE\_VOLUME
    - PUBLISH\_UNPUBLISH\_VOLUME
```

#### CreateSnapshot

Not implemented yet.

#### DeleteSnapshot

Not implemented yet.

#### ListSnapshots

Not implemented yet.

### Node Service RPC

#### NodeStageVolume

1. Find the device.
2. Check if it&#39;s unformatted (lsblk or blkid).
3. Format it if it&#39;s needed.
4. fsck it if it&#39;s already formatted + refuse to mount it on errors.
5. Mount it to given directory (with given mount options).

Steps 3 and 4 can take some time, so the driver must ensure idempotency somehow.

#### NodeUnstageVolume

Just unmount the volume.

#### NodePublishVolume

Just bind-mount the volume.

#### NodeUnpublishVolume

Just unmount the volume.

#### NodeGetInfo

Blindly return:

```
    NodeId: AWS InstanceID.
    AccessibleTopology: {"topology.ebs.csi.aws.com/zone": [availablility zone]}
```

#### NodeGetId

Return AWS InstanceID.

#### NodeGetCapabilities

Blindly return:

```
  rpc:
    - STAGE\_UNSTAGE\_VOLUME
```

