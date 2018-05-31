---
kep-number: TBD
title: Support traffic shaping for Kubelet CNI network plugin
authors:
    - "@Lion-Wei"
    - "@m1093782566"
owning-sig: sig-network
reviewers:
  - "@thockin"
  - "@m1093782566"
approvers:
  - "@thockin"
  - "@m1093782566"
editor:
  - "@thockin"
  - "@m1093782566"
creation-date: 2018-05-31
---

# Support traffic shaping for Kubelet CNI network plugin

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non\-goals](#non-goals)
* [Proposal](#proposal)
  * [Pod Setup](#pod-setup)
  * [Pod Teardown](#pod-teardown)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
  * [CNI plugin part](#cni-plugin-part)
  * [Kubernetes part](#kubernetes-part)
* [Drawbacks](#drawbacks)
* [Alternatives](#alternatives)

## Summary

Make kubelet cni network plugin support basic traffic shapping capability `bandwidth`.

## Motivation

Currently the kubenet code supports applying basic traffic shaping during pod setup.
This will happen if bandwidth-related annotations have been added to the pod's metadata.

Kubelet CNI code doesn't support it yet, though CNI has already added a [traffic sharping plugin](https://github.com/containernetworking/plugins/tree/master/plugins/meta/bandwidth).
We can replicate the behavior we have today in kubenet for kubelet CNI network plugin if we feel this is an important feature.

### Goals

* Support traffic shaping for CNI network plugin in Kubernetes.

### Non-goals

* CNI plugins to implement this sort of traffic shaping guarantee.


## Proposal

If kubelet starts up with `network-plugin = cni` and user enabled traffic shaping via the network plugin configuration,
it would then populate the runtimeConfig section of the config when calling the bandwidth plugin.

Traffic shaping in Kubelet CNI network plugin can work with ptp and bridge network plugins.

### Pod Setup

When we create a pod with bandwidth configuration in its metadata, for example,

```json
{
    "kind": "Pod",
    "metadata": {
        "name": "iperf-slow",
        "annotations": {
            "kubernetes.io/ingress-bandwidth": "10M",
            "kubernetes.io/egress-bandwidth": "10M"
        }
    }
}
```

Kubelet would firstly parse the ingress and egress bandwidth values and transform them to ingressRate and egressRate for cni bandwidth plugin.
Kubelet would then detect whether user has enabled the traffic shaping plugin by checking the following CNI config file:

```json
{
  "type": "bandwidth",
  "capabilities": {"trafficShaping": true}
}
```

If traffic shaping plugin is enabled, kubelet would populate the runtimeConfig section of the config when call the bandwidth plugin:

```json
{
  "type": "bandwidth",
  "runtimeConfig": {
    "trafficShaping": {
      "ingressRate": "X",
      "egressRate": "Y"
    }
  }
}
```

### Pod Teardown

When we delete a pod, kubelet will bulid the runtime config call cni plugin DelNetworkList API, which will remove this pod's bandwidth configuration.

## Graduation Criteria

* Add traffic shaping as part of the Kubernetes e2e runs and ensure tests are not failing.

## Implementation History

### CNI plugin part

* [traffic shaping plugin](https://github.com/containernetworking/plugins/pull/96)
* [support runtime config](https://github.com/containernetworking/plugins/pull/138)

### Kubernetes part

* [add traffic shaping support](https://github.com/kubernetes/kubernetes/pull/63194)

## Drawbacks [optional]

None

## Alternatives [optional]

None