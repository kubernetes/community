# Container Device Interface

**Author**: Dennis Schridde <dennis.schridde@uni-heidelberg.de> (@urzds)

**Version**: 0.0.1

**Last edit**: [2017-04-18](#history)

**Status**: Draft proposal

## Motivation

We use Kubernetes to schedule scientific 3D visualisation applications onto a cluster. Hence our containers use GPUs / DRM and other device nodes (e.g. TTYs).  Scheduling of Pods to Nodes with a capacity for such devices can be dealt with using [Opaque Integer Resources (OIR)](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#opaque-integer-resources-alpha-feature), where scripts external to Kubernetes handle device discovery and capacity reporting (see e.g. [kube-gpu-resources](https://github.com/urzds/kube-gpu-resources)).  The [Container Runtime Interface (CRI)](https://github.com/kubernetes/community/blob/master/contributors/devel/container-runtime-interface.md) already supports [binding host device nodes](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/runtime/api.proto#L586) to a container.  What is currently missing is a way for the Kubelet to generically allocate the available devices of a certain type to a container while it is being started.  The Container Device Interface (CDI) described here shall bridge this gap.

## Description

CDI is modelled after CNI in that it consist of an API between the Kubelet and CDI binaries that transports instructions through environment variables and configuration and data through JSON chunks transmitted through stdin/stdout.  As in CNI, which binaries to call and their configuration is described through JSON configuration files located in `/etc/cdi/<device-type>.d/<binary-name>.conf`.  The principal commands are `ADD` and `DEL`, which allocate a device to a container or return it to the node's pool.  All commands expect a `CDI_VERSION` environment variable, to ensure the version used by the Kubelet (environment variable), the CDI config (JSON chunk on stdin) and the binary match.

### ADD

`CDI_COMMAND=ADD` allocates a device to a container.  The Kubelet communicates the requested amount of resources through a `CDI_REQUEST=<resource-subtype>:<amount>,...` environment variable.  The CDI binary creates or picks suitable devices, does any potentially necessary setup, stores the association `host device node <> container` and returns the host path of the device node.  The Kubelet will then pass this information as a `Device` to CRI.  Failure of this command results in a non-zero exit status and a JSON encoded error message on stdout.

### DEL

`CDI_COMMAND=DEL` returns devices allocated to a container to the pool.  The CDI binary retrieves the list of device nodes allocated to this container from its own storage, runs any potentially necessary tear down steps and clears the device node / container association from its storage.  This command can not fail, i.e. it always returns a zero exit status, but might produce an error message on stdout, suitable for logging the incident.

## Rationale

We propose out-sourcing the management of device allocations to external binaries, to easily allow to support different device types and provide the possibility for future extensions e.g. using virtualisation techniques.  (In our case, e.g. PCIe SR-IOV could be used to dynamically allocate a certain amount of timeslices or VRAM.)  We propose an API similar to CNI, since its design is already well tested within Kubernetes and capable of this task.

## Usage examples

### GPU / DRM

```
# cat /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "name": "drm-gpus",
  "type": "drm",
  "args": {
    "device_node_type": "all"
  }
}

# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=gpu:1 CDI_CONTAINERID=1234 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "devices": [
    "/dev/dri/card0",
    "/dev/dri/renderD128"
  ]
}
#=> exit status 0

# env CDI_VERSION=0.0.1 CDI_COMMAND=DEL CDI_CONTAINERID=1234 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1"
}
#=> exit status 0

# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=gpu:1,gpu-memory:2048Mi CDI_CONTAINERID=3456 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "code": 3,
  "msg": "Resource sub-type unsupported",
  "details: "Unsupported resource sub-type: gpu-memory"
}
#=> exit status 1

# env CDI_VERSION=0.0.1 CDI_COMMAND=DEL CDI_CONTAINERID=3456 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "code": 4,
  "msg": "Unknown container ID"
}
#=> exit status 0

# env CDI_VERSION=0.999 CDI_COMMAND=ADD ... /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "code": 1,
  "msg": "Incompatible CDI version",
  "details: "Unsupported version: 0.999"
}
#=> exit status 1

# env CDI_VERSION=0.0.1 CDI_COMMAND=MYCMD ... /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf
{
  "cdiVersion": "0.0.1",
  "code": 5,
  "msg": "Command unsupported",
  "details: "Unsupported command: MYCMD"
}
#=> exit status 1
```

### TTY

```
# cat /etc/cdi/tty.d/generic.conf
{
  "cdiVersion": "0.0.1",
  "name": "ttys",
  "type": "generic-tty",
  "args": {
    "num_reserved": 12
  }
}

# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=tty:3 CDI_CONTAINERID=7890 /opt/cdi/bin/generic-tty < /etc/cdi/tty.d/generic.conf
{
  "cdiVersion": "0.0.1",
  "devices": [
    "/dev/tty13",
    "/dev/tty14",
    "/dev/tty15"
  ]
}
#=> exit status 0
```

## Implementation

* [ ] CDI binary
  - [ ] Write sample implementation
* [ ] Kubelet
  - [ ] Call CDI binary when starting a container requesting device resources and pass results to CRI

## References

See-Also: https://github.com/kubernetes/kubernetes/issues/5607
See-Also: https://github.com/kubernetes/kubernetes/issues/44107

## History

* 2017-04-18 (@urzds)
  - Initial proposal
