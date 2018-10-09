---
kep-number: 29
title: Resource configuration in device plugins upon container start
authors:
  - "@rojkov"
  - "@bart0sh"
  - "@kad"
owning-sig: sig-node
reviewers:
  - "@derekwaynecarr"
  - "@dashpole"
  - "@jiayingz"
  - "@vikaschoudhary16"
  - "@vishh"
  - "@RenaudWasTaken"
approvers:
  - "@sig-node-leads"
editors:
  - "@rojkov"
  - "@bart0sh"
  - "@kad"
creation-date: 2018-10-08
last-updated: 2018-10-08
status: provisional
---

# Resource configuration in device plugins upon container start

## Table of contents

* [Summary](#summary)
* [Motivation](#motivation)
   * [Goals](#goals)
* [Proposal](#proposal)
   * [User Stories](#user-stories)
      * [Cluster orchestrated FPGA configuration](#cluster-orchestrated-fpga-configuration)
   * [Risks and Mitigations](#risks-and-mitigations)
* [Implementation History](#implementation-history)
* [Alternatives](#alternatives)

## Summary

Provide container level resource configuration to Device Plugins before starting
containers in order to configure the devices consumed by the containers.

## Motivation

Devices like FPGAs can be configured dynamically to offload CPUs with various
types of acceleration specific to pod workloads. At the same time FPGAs
impose big security risks in multi-tenant configurations of Kubernetes because
they can be configured to access memory used by other tenants' processes. That is
why it's better to disallow direct access to programming facilities of FPGAs for
pod workloads and to provide the workloads with devices already pre-configured
by the device plugin.

### Goals

* Provide a dynamically pre-configured device to a pod workload.

## Proposal

Extend the struct `PreStartContainerRequest` of the Device Plugin API with
a new field called `Env` of the type `map[string]string`. This field is
populated by kubelet with effective environment variables defined in Pod
spec.

### User Stories

#### Cluster orchestrated FPGA configuration

A user creates the following pod:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: intel-fpga-demo-compress
spec:
  containers:
  - name: intel-fpga-demo-compress
    image: intel-fpga-demo-compress:devel
    imagePullPolicy: IfNotPresent
    command: ["sh", "/opt/intel/intel-fpga-demo-compress/demo/demo.sh"]
    securityContext:
      capabilities:
        add:
          [IPC_LOCK]
    resources:
      limits:
        fpga.intel.com/region-ce48969398f05f33946d560708be108a: 1
        cpu: 1
        hugepages-2Mi: 20Mi
    env:
      FPGA_BITSTREAM_ID_1: 18b79ffa2ee54aa096ef4230dafacb5f
      FPGA_REGION_1: ce48969398f05f33946d560708be108a
      FPGA_DRIVER_PARAMS_1: clock=1000,debug=1
  restartPolicy: Never
```

Just before starting the `intel-fpga-demo-compress` container kubelet makes
a `PreStartContainer` call to the device plugin and passes the content of the
`env` dictionary in the field `Env` of the `PreStartContainerRequest` parameter.

The device plugin checks if the value `FPGA_REGION_1` corresponds to the
interface ID of the allocated device, programs the device with the
bitstream whose ID is given in `FPGA_BITSTREAM_ID_1` and overrides default
driver settings with the content of `FPGA_DRIVER_PARAMS_1`.

Then kubelet starts the container whose workload can use the programmed device
right away.

### Risks and Mitigations

No known risks.

## Implementation History

- Initial version of this KEP.

## Alternatives

Environment variables defined in Pod spec are meant to be used in containers.
This KEP proposes their use outside of containers, but the alternative
would be introducing a new container level field in the Core API for `Pod`
objects specifically targeted to Device Plugins. Such limited use
doesn't justify the change of the Core API.

A special programming agent can run on FPGA-enabled nodes which drives FPGA
configuration in controlled manner upon requests from workloads. This agent
has to be smart enough to control access to bitstreams available for a given
tenant though. We may end up with a situation when a workload is scheduled
on a node, but it fails because of insufficient permissions. A more
natural approach would be to check permissions in an admission webhook
before admitting the workload for scheduling, but it returns us to the
proposed change in the Device Plugin API.

Also the KEP ["Kubelet endpoint for device assignment observation details"](https://github.com/kubernetes/community/pull/2454)
can be used to find Pod's name from kubelet's endpoint (and then to fetch
the Pod's specs from the API server) upon a `PreStartContainer` call in the
device plugin. This approach puts additional load on the API servers and
introduces additional delays. Another concern is that in the current
[POC](https://github.com/dashpole/kubernetes/tree/device_id) of the KEP the
code which updates the "/pods" endpoint is run as an
[asynchronous task](https://github.com/dashpole/kubernetes/blob/9ff717ee9c87d5b3248a3d28b8893e21028ea42d/pkg/kubelet/kubelet.go#L1982)
and this creates the possibility of race conditions in real production
environment.
