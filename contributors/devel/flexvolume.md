# Flexvolume

Flexvolume enables users to write their own drivers and add support for their volumes in Kubernetes. Vendor drivers should be installed in the volume plugin path on every node, and on master if the driver requires attach capability (unless `--enable-controller-attach-detach` Kubelet option is set to false, but this is highly discouraged because it is a legacy mode of operation).

Flexvolume is a GA feature from Kubernetes 1.8 release onwards.

## Prerequisites

Install the vendor driver on all nodes (also on master nodes if "--enable-controller-attach-detach" Kubelet option is enabled) in the plugin path. Path for installing the plugin: `<plugindir>/<vendor~driver>/<driver>`. The default plugin directory is `/usr/libexec/kubernetes/kubelet-plugins/volume/exec/`. It can be changed in kubelet via the `--volume-plugin-dir` flag, and in controller manager via the `--flex-volume-plugin-dir` flag.

For example to add a `cifs` driver, by vendor `foo` install the driver at: `/usr/libexec/kubernetes/kubelet-plugins/volume/exec/foo~cifs/cifs`

The vendor and driver names must match flexVolume.driver in the volume spec, with '~' replaced with '/'. For example, if `flexVolume.driver` is set to `foo/cifs`, then the vendor is `foo`, and driver is `cifs`.

## Dynamic Plugin Discovery
Beginning in v1.8, Flexvolume supports the ability to detect drivers on the fly. Instead of requiring drivers to exist at system initialization time or having to restart kubelet or controller manager, drivers can be installed, upgraded/downgraded, and uninstalled while the system is running.
For more information, please refer to the [design document](/contributors/design-proposals/storage/flexvolume-deployment.md).

## Automated Plugin Installation/Upgrade
One possible way to install and upgrade your Flexvolume drivers is by using a DaemonSet. See [Recommended Driver Deployment Method](/contributors/design-proposals/storage/flexvolume-deployment.md#recommended-driver-deployment-method) for details, and see [here](https://git.k8s.io/examples/staging/volumes/flexvolume/deploy/) for an example.

## Plugin details
The plugin expects the following call-outs are implemented for the backend drivers. Some call-outs are optional. Call-outs are invoked from Kubelet and Controller Manager.

### Driver invocation model:

#### Init:
Initializes the driver. Called during Kubelet & Controller manager initialization. On success, the function returns a capabilities map showing whether each Flexvolume capability is supported by the driver.
Current capabilities:
* `attach` - a boolean field indicating whether the driver requires attach and detach operations. This field is *required*, although for backward-compatibility the default value is set to `true`, i.e. requires attach and detach.
See [Driver output](#driver-output) for the capabilities map format.
```
<driver executable> init
```

#### Attach:
Attach the volume specified by the given spec on the given node. On success, returns the device path where the device is attached on the node. Called from Controller Manager.

This call-out does not pass "secrets" specified in Flexvolume spec. If your driver requires secrets, do not implement this call-out and instead use "mount" call-out and implement attach and mount in that call-out.

```
<driver executable> attach <json options> <node name>
```

#### Detach:
Detach the volume from the node. Called from Controller Manager.
```
<driver executable> detach <mount device> <node name>
```

#### Wait for attach:
Wait for the volume to be attached on the remote node. On success, the path to the device is returned. Called from Controller Manager. The timeout should be 10m (based on https://git.k8s.io/kubernetes/pkg/kubelet/volumemanager/volume_manager.go#L88 )

```
<driver executable> waitforattach <mount device> <json options>
```

#### Volume is Attached:
Check the volume is attached on the node. Called from Controller Manager.

```
<driver executable> isattached <json options> <node name>
```

#### Mount device:
Mount device mounts the device to a global path which individual pods can then bind mount. Called only from Kubelet.

This call-out does not pass "secrets" specified in Flexvolume spec. If your driver requires secrets, do not implement this call-out and instead use "mount" call-out and implement attach and mount in that call-out.

```
<driver executable> mountdevice <mount dir> <mount device> <json options>
```

#### Unmount device:
Unmounts the global mount for the device. This is called once all bind mounts have been unmounted. Called only from Kubelet.

```
<driver executable> unmountdevice <mount device>
```
In addition to the user-specified options and [default JSON options](#default-json-options), the following options capturing information about the pod are passed through and generated automatically.

```
kubernetes.io/pod.name
kubernetes.io/pod.namespace
kubernetes.io/pod.uid
kubernetes.io/serviceAccount.name
```

#### Mount:
Mount the volume at the mount dir. This call-out defaults to bind mount for drivers which implement attach & mount-device call-outs. Called only from Kubelet.

```
<driver executable> mount <mount dir> <json options>
```

#### Unmount:
Unmount the volume. This call-out defaults to bind mount for drivers which implement attach & mount-device call-outs. Called only from Kubelet.

```
<driver executable> unmount <mount dir>
```

See [lvm] & [nfs] for a quick example on how to write a simple flexvolume driver.

### Driver output:

Flexvolume expects the driver to reply with the status of the operation in the
following format.

```
{
	"status": "<Success/Failure/Not supported>",
	"message": "<Reason for success/failure>",
	"device": "<Path to the device attached. This field is valid only for attach & waitforattach call-outs>"
	"volumeName": "<Cluster wide unique name of the volume. Valid only for getvolumename call-out>"
	"attached": <True/False (Return true if volume is attached on the node. Valid only for isattached call-out)>
    "capabilities": <Only included as part of the Init response>
    {
        "attach": <True/False (Return true if the driver implements attach and detach)>
    }
}
```

### Default Json options

In addition to the flags specified by the user in the Options field of the FlexVolumeSource, the following flags (set through their corresponding FlexVolumeSource fields) are also passed to the executable.
Note: Secrets are passed only to "mount/unmount" call-outs.

```
"kubernetes.io/fsType":"<FS type>",
"kubernetes.io/readwrite":"<rw>",
"kubernetes.io/fsGroup":"<FS group>",
"kubernetes.io/mountsDir":"<string>",
"kubernetes.io/pvOrVolumeName":"<Volume name if the volume is in-line in the pod spec; PV name if the volume is a PV>"

"kubernetes.io/pod.name":"<string>",
"kubernetes.io/pod.namespace":"<string>",
"kubernetes.io/pod.uid":"<string>",
"kubernetes.io/serviceAccount.name":"<string>",

"kubernetes.io/secret/key1":"<secret1>"
...
"kubernetes.io/secret/keyN":"<secretN>"
```

### Example of Flexvolume

Please refer to the [Flexvolume example directory]. See [nginx-lvm.yaml] & [nginx-nfs.yaml] for a quick example on how to use Flexvolume in a pod.


[lvm]: https://git.k8s.io/examples/staging/volumes/flexvolume/lvm
[nfs]: https://git.k8s.io/examples/staging/volumes/flexvolume/nfs
[nginx-lvm.yaml]: https://git.k8s.io/examples/staging/volumes/flexvolume/nginx-lvm.yaml
[nginx-nfs.yaml]: https://git.k8s.io/examples/staging/volumes/flexvolume/nginx-nfs.yaml
[Flexvolume example directory]: https://git.k8s.io/examples/staging/volumes/flexvolume/
