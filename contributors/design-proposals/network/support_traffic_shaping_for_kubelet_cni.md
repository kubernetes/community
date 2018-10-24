# Support traffic shaping for CNI network plugin

Version: Alpha

Authors: @m1093782566

## Motivation and background

Currently the kubenet code supports applying basic traffic shaping during pod setup. This will happen if bandwidth-related annotations have been added to the pod's metadata, for example:

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

Our current implementation uses the `linux tc` to add an download(ingress) and upload(egress) rate limiter using 1 root `qdisc`, 2 `class `(one for ingress and one for egress) and 2 `filter`(one for ingress and one for egress attached to the ingress and egress classes respectively).

Kubelet CNI code doesn't support it yet, though CNI has already added a [traffic sharping plugin](https://github.com/containernetworking/plugins/tree/master/plugins/meta/bandwidth). We can replicate the behavior we have today in kubenet for kubelet CNI network plugin if we feel this is an important feature.

## Goal

Support traffic shaping for CNI network plugin in Kubernetes.

## Non-goal

CNI plugins to implement this sort of traffic shaping guarantee.

## Proposal

If kubelet starts up with `network-plugin = cni` and user enabled traffic shaping via the network plugin configuration, it would then populate the `runtimeConfig` section of the config when calling the `bandwidth` plugin.

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

Kubelet would firstly parse the ingress and egress bandwidth values and transform them to Kbps because both `ingressRate` and `egressRate` in cni bandwidth plugin are in Kbps. A user would add something like this to their CNI config list if they want to enable traffic shaping via the plugin:

```json
{
  "type": "bandwidth",
  "capabilities": {"trafficShaping": true}
}
```

Kubelet would then populate the `runtimeConfig` section of the config when calling the `bandwidth` plugin:

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

When we delete a pod, kubelet will build the runtime config for calling cni plugin `DelNetwork/DelNetworkList` API, which will remove this pod's bandwidth configuration.

## Next step

* Support ingress and egress burst bandwidth in Pod.
* Graduate annotations to Pod Spec.
