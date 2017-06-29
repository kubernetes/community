# Container Device Interface

**Author**: Dennis Schridde <dennis.schridde@uni-heidelberg.de> (@urzds)

**Version**: 0.0.2

**Last edit**: [2017-05-05](#history)

**Status**: Draft proposal

## Introduction

### Motivation

We use Kubernetes to schedule scientific 3D visualisation applications onto a cluster. Hence our containers use GPUs with the Linux Direct Rendering Manager (DRM) and other device nodes (e.g. TTYs).  Scheduling of Pods to Nodes with a capacity for such devices can be dealt with using [Opaque Integer Resources (OIR)](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#opaque-integer-resources-alpha-feature), where scripts external to Kubernetes handle device discovery and capacity reporting (see e.g. [kube-gpu-resources](https://github.com/urzds/kube-gpu-resources)).  The [Container Runtime Interface (CRI)](https://github.com/kubernetes/community/blob/master/contributors/devel/container-runtime-interface.md) already supports [binding host device nodes](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/runtime/api.proto#L586) to a container.  What is currently missing is a way for the Kubelet to generically allocate the available devices of a certain type to a container while it is being started.  The Container Device Interface (CDI) described here shall bridge this gap.

We propose out-sourcing the management of device node allocations to external binaries, to easily allow to support different device types and provide the possibility for future extensions e.g. using device partitioning.  (E.g. PCIe SR-IOV could be used to dynamically allocate a certain amount of timeslices or VRAM.)  We propose an API similar to CNI, since its design is already well tested within Kubernetes and capable of this task.  This specification also contains a command that can be used for device discovery, which can be used to implement or supplant OIR for device nodes.

### Overview

CDI is modelled after CNI in that it specifies an API between the Kubelet and CDI binaries, which transports requests through environment variables and configuration and responses through JSON chunks transmitted via stdin/stdout.  As with CNI, which binaries to call and their configuration is described via JSON configuration files located in `/etc/cdi/<resource-type>.d/<binary-name>.conf`.  The binaries themselves are to be located in `/opt/cdi/bin`, similar to CNI.

### Terminology

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

The terms `resource-type`, `resource-subtype`, `resource-spec` and `request-id` shall be defined by the [ABNF rules below](#common-abnf-rules).

We generally distinguish the terms "device" (meaning the physical hardware) and "device node" (meaning the file-like entity in the Linux file-system).

## Configuration

Configuration files contain a JSON struct and are to be placed in `/etc/cdi/<resource-type>.d/<plugin-name>.conf`, i.e. configuration files for plugins serving the same `resource-type` are to be placed in the same folder.

### Required fields

* `cdiVersion` is the version of this specification that the configuration adheres to.
* `name` is a human-readable descriptive name and does not have any defined interpretation.
* `type` is the `resource-type` this plugin handles.  This MUST be the same as the `resource-type` that is part of the folder name and is included to communicate this type to the plugin.
* `plugin` is the name of the binary to call.

### Optional fields
* `args` is a struct -- the interpretation of which is open to the plugin.  CDI does not define anything about its contents.

### Examples

`/etc/cdi/gpu.d/drm.conf` contains:
```json
{
  "cdiVersion": "0.0.1",
  "name": "DRM GPUs from vendor International Amazing Graphics Inc.",
  "type": "gpu",
  "plugin": "drm",
  "args": {
    "want_device_nodes" : [
      "primary",
      "render"
    ],
    "vendorid_whitelist": [
      "0x1234",
      "0x2345"
    ]
  }
}
```

Here `want_device_types` lists the types of device nodes we are interested in, since Linux DRM provides multiple nodes for one physical device.  The `vendorid_whitelist` can be used to limit this plugin to allocate only devices from a certain hardware vendor, i.e. with a certain PCIe vendor ID.

`/etc/cdi/tty.d/generic.conf` contains:
```json
{
  "cdiVersion": "0.0.1",
  "name": "TTYs",
  "type": "tty",
  "plugin": "generic-tty",
  "args": {
    "num_system_reserved": 12
  }
}
```

`num_system_reserved` would be the number of TTYs reserved for the system, i.e. not usable by containers. Or in other words the number of the first TTY that could be used by a container.

## Commands

A command is issued through the `CDI_COMMAND` environment variable, which MUST be set by the caller.  Commands SHALL be interpreted case-insensitive by the plugin.  For all commands, except `VERSION`, the `CDI_VERSION` environment variable MUST be set by the caller.  The latter is to ensure the version used by the Kubelet (environment variable), the CDI config (JSON chunk on stdin) and the binary match.  For each command, certain parameters are required to be passed through environment variables -- these are listed under "required parameters".  Other parameters are optional (hence listed under "optional parameters"): Plugins MAY react to them, if they implement the required functionality, or otherwhise ignore them.

Whether the plugin responds to a command by printing a JSON struct on stdout depends on the command.  If anything is printed to stdout, the JSON struct MUST contain a `cdiVersion` field, specifying the version of the output format.  The exit status of the plugin determines whether the command was successful (0) or not (anything else).

### Errors

If an error occurs, the plugin MUST respond by printing a JSON struct to stdout describing the error, which MUST contain following fields:
* `error` - an integer identifying the error.  Error codes 0-99 are reserved for [well-known error codes](#well-known-error-codes).  Values of 100+ can be freely used for plugin-specific errors.
* `message` - a short, human-readable description of the error

In addition, the response MAY contain:
* `details` - a possibly longer description with additional details suitable for debugging.

In case of an error, the plugin SHALL exit with an exit status not equal to 0.  The exit status and the error code do not necessarily have to match.

#### Well-known error codes

Error code | Description
-----------|------------
1          | CDI version requested via env-var is not supported by plugin
2          | CDI version of the config is not supported by plugin
3          | The command is not supported
4          | Resource spec is not supported
5          | The referenced request ID is unknown

#### Examples

1. Specifying an unknown `CDI_VERSION`
  - `# env CDI_VERSION=0.999 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "code": 1,
  "msg": "Incompatible CDI version",
  "details": "Unsupported version: 0.999"
}
```
  - Exit status: 1

2. Specifying an unknown `CDI_COMMAND`
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=MYCMD /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "code": 3,
  "msg": "Command unsupported",
  "details": "Unsupported command: MYCMD"
}
```
  - Exit status: 1

### VERSION

This allows the Kubelet and the CDI to negotiate a version of the specification to use.  In response to `CDI_COMMAND=VERSION` the plugin SHALL reply with the versions of the specification supported by the plugin in the `supportedVersions` field of the reply, which MUST be of type array-of-strings.  This command MUST NOT fail.

#### Required parameters

None

#### Optional parameters

None

#### Examples

1. Retrieving the versions supported by this plugin
  - `# env CDI_COMMAND=VERSION /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.2",
  "supportedVersions": [ "0.0.1", "0.0.2" ]
}
```
  - Exit status: 0

### INFO

This command can be used by the caller to discover the number of devices of certain types available on this node.  No additional arguments are required.  The plugin SHALL reply to this command with a struct of `resource-spec` to `available-amount` mappings, the latter being an integer.  The plugin MUST mention an `available-amount` for the `resource-type` specified in its configuration, it SHALL NOT advertise any resources that do not start with the specified `resource-type`, but MAY additionaly advertise `resource-subtype`s thereof.

#### Required parameters

None

#### Optional parameters

None

#### Examples

1. Retrieve number of available GPUs and related sub-resources
  - `# env CDI_COMMAND=INFO /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "gpu": 12,
  "gpu-memory": 65536
}
```
  - Exit status: 0

2. Retrieve number of available TTYs
  - `# env CDI_COMMAND=INFO /opt/cdi/bin/generic-tty < /etc/cdi/tty.d/generic.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "tty": 32
}
```
  - Exit status: 0

### ADD

`CDI_COMMAND=ADD` allocates a device to a container.  The Kubelet communicates the requested amount of resources through a `CDI_REQUEST=<resource-spec>:<amount>,...` environment variable.  The `amount` of requested resources MUST be an integer.  Multiple resources MAY be requested in a comma-separated list.  The caller is also REQUIRED to provide a `CDI_REQUEST_ID=<request-id>`, which serves as an identifier for the pod or container these resources are being allocated for and MUST conform to `request-id`.  The caller MUST NOT call the plugin multiple times with the same `request-id` and `resource-spec`.

The CDI binary creates or picks suitable devices, does any potentially necessary setup, stores the association of `host device nodes` to `request-id` and returns a list of paths to the host device nodes in the `devices` field of the response.  The intention of this is that, before creating the container, the Kubelet gathers a list of device nodes by executing CDI plugins in `kubelet.makeDevices()`, which it then hands of as `kubecontainer.DeviceInfo` to CRI, when setting up the pod and container.

The JSON response to a successfully executed command MUST include a `devices` field, even if it would be empty.

The command SHALL fail, if the request cannot be fulfilled in its entirety.  Failure of this command results in a non-zero exit status and a JSON encoded error message on stdout.

The number of `devices` listed in the response does not necessarily correspond to the `amount` of resources requested, since multiple device nodes per device might be necessary to fulfil the request, or multiple devices might be accessed through a common device node.  E.g. DRM GPU devices might consist of a render node and a control node, while on the other hand an array of several FPGAs, connected via a common bus, might be accessible through just one device node.

#### NUMA locality

This command supports an optional environment variable `CDI_REQUEST_NUMA_LOCALITY=<node-ids...>`, with `node-ids` being a comma separated list of integers.  If this variable is set, the CDI plugin SHOULD try to find a device which is close to at least one of the listed NUMA nodes.  The response SHOULD include a `numa_locality` field in the response, which consists of a list of NUMA nodes, identified by integer ID, that are close to the allocated devices.  The response MAY include this field, even if the plugin was not requested to do so by setting `CDI_REQUEST_NUMA_LOCALITY`.

#### Required parameters
* `CDI_REQUEST`
* `CDI_REQUEST_ID`

#### Optional parameters
* `CDI_REQUEST_NUMA_LOCALITY`

#### Examples

1. Allocate 1 GPU for our pod or container with Kubernetes internal ID 1234
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=gpu:1 CDI_REQUEST_ID=1234 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "devices": [
    "/dev/dri/card0",
    "/dev/dri/renderD128"
  ]
}
```
  - Exit status: 0

2. Allocate 3 TTYs
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=tty:3 CDI_REQUEST_ID=7890 /opt/cdi/bin/generic-tty < /etc/cdi/tty.d/generic.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "devices": [
    "/dev/tty13",
    "/dev/tty14",
    "/dev/tty15"
  ]
}
```
  - Exit status: 0

3. Error when trying to allocate with unknown sub-resource
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=ADD CDI_REQUEST=gpu:1,gpu-memory:2048Mi CDI_REQUEST_ID=3456 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "code": 4,
  "msg": "resource-spec unsupported",
  "details": "Unsupported resource-spec: gpu-memory"
}
```
  - Exit status: 1

### DEL

`CDI_COMMAND=DEL` returns devices allocated to a container to the pool.  The CDI binary retrieves the list of device nodes allocated to this container from its own storage, runs any potentially necessary tear down steps and clears the association of `host device node` to `request-id` from its internal storage.  The caller MUST set `CDI_REQUEST_ID=<request-id>` as with the `ADD` command.  This command MUST NOT fail, i.e. it always returns a zero exit status, but might produce an error message on stdout, suitable for logging the incident.

#### Required parameters
* `CDI_REQUEST_ID`

#### Examples

1. Successfully delete the association of device nodes to previous allocation with request-id 1234
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=DEL CDI_REQUEST_ID=1234 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output: None
  - Exit status: 0

2. Fail to delete the association for an unknown request-id
  - `# env CDI_VERSION=0.0.1 CDI_COMMAND=DEL CDI_REQUEST_ID=3456 /opt/cdi/bin/drm < /etc/cdi/gpu.d/drm.conf`
  - Output:
```json
{
  "cdiVersion": "0.0.1",
  "code": 5,
  "msg": "Unknown request ID"
}
```
  - Exit status: 0


## Kubernetes Pod Spec examples

### Device node exclusively allocated to container

The simplest use-case is for a container to exclusively allocate a device.  This works just like any other resource request.  The difference would be that for every OIR requested the Kubelet would call the corresponding CDI, stripping `pod.alpha.kubernetes.io/opaque-int-resource-` from the request and then handing it over as described for the `ADD` command.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: exclusive
spec:
  containers:
  - name: compute
    image: compute-app:v1
    resources:
      limits:
        cpu: 1000m
        memory: 12Gi
      requests:
        cpu: 1000m
        memory: 12Gi
        pod.alpha.kubernetes.io/opaque-int-resource-gpu: 1
        pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 8Gi
```

### Device node shared among mulitple containers in a Pod

If a device shall be shared, it has to be requested by the pod.  Containers can then piggy-back on the pod-level resource request by mentioning the names of the resources in a special `shared` sub-section of their resource section.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: sharing
spec:
  resources:
    requests:
      pod.alpha.kubernetes.io/opaque-int-resource-gpu: 1
      pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 8Gi
  containers:
  - name: preprocessor
    image: preprocess-app:v1
    resources:
      requests:
        pod.alpha.kubernetes.io/opaque-int-resource-gpu: 2
        pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 16Gi
      shared:
      - pod.alpha.kubernetes.io/opaque-int-resource-gpu
  - name: analyst
    image: compute-app:v1
    resources:
      requests:
        pod.alpha.kubernetes.io/opaque-int-resource-gpu: 1
        pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 4Gi
      shared:
      - pod.alpha.kubernetes.io/opaque-int-resource-gpu
```

This does not allow for multiple containers to share different devices (say containers A and B share one GPU, while B and C share another), which would require that resource requests can be identified by some sort of ID, similar to PersistentVolumeClaims:

```yaml
apiVersion: v1
kind: ResourceRequest
metadata:
  name: shared-gpus-8g
spec:
  resources:
    requests:
      pod.alpha.kubernetes.io/opaque-int-resource-gpu: 1
      pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 8Gi

---

apiVersion: v1
kind: ResourceRequest
metadata:
  name: shared-gpus-2x4g
spec:
  resources:
    requests:
      pod.alpha.kubernetes.io/opaque-int-resource-gpu: 2
      pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 4Gi

---

apiVersion: v1
kind: Pod
metadata:
  name: sharing
spec:
  containers:
  - name: preprocessor
    image: preprocess-app:v1
    resources:
      requests:
        pod.alpha.kubernetes.io/opaque-int-resource-gpu: 2
        pod.alpha.kubernetes.io/opaque-int-resource-gpu-memory: 16Gi
      shared:
      - resourceRequestRef: shared-gpus-8g
  - name: analyst
    image: compute-app:v1
    resources:
      shared:
      - resourceRequestRef: shared-gpus-8g
      - resourceRequestRef: shared-gpus-2x4g
  - name: postprocessor
    image: postprocess-app:v1
    resources:
      shared:
      - resourceRequestRef: shared-gpus-2x4g
```

### Fractions of a device allocated to container

Fractions of a device can be allocated, iff the device itself supports it, e.g. through PCIe SR-IOV.  Partitioning the device could be done through an SR-IOV CDI plugin, which chain-loads a more generic CDI plugin that deals with managing the actual device node.  This mechanism is not yet defined in this specification.

## Implementation

* [ ] CDI binary
  - [ ] Write sample implementation
* [ ] Kubelet
  - [ ] Call CDI binary when starting a container requesting device resources and pass results to CRI

## References

See-Also: https://github.com/kubernetes/kubernetes/issues/5607
See-Also: https://github.com/kubernetes/kubernetes/issues/44107

## Common ABNF rules

This specification uses the ABNF core rules as defined by [RFC 5234](https://tools.ietf.org/html/rfc5234).

```abnf
resource-spec = resource-type [ hyphen resource-subtype ]

resource-subtype = 1*15alnum

resource-type = ALPHA *15alnum

request-id = alnum *63alnum-hyp

alnum = ALPHA / DIGIT

alnum-hyp = alnum / hyphen

hyphen = %x2D
```

## History

* 2017-05-05 (@urzds)
  - Rewrite in the style of an actual specification, go more into detail and add more explanations
  - Add VERSION command
  - Add INFO command for device discovery
  - Add NUMA extension to ADD command
  - Add Kubernetes Pod Spec examples
  - Mention plugin chain-loading, but leave specification for later
  - Mention an alternative to Kubernetes OIRs for device discovery

* 2017-04-18 (@urzds)
  - Initial proposal
