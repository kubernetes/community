# apiserver-count fix proposal

Authors: @rphillips

## Table of Contents

1. [Overview](#overview)
2. [Known Issues](#known-issues)
3. [Proposal](#proposal)
4. [Alternate Proposals](#alternate-proposals)
    1. [Custom Resource Definitions](#custom-resource-definitions)
    2. [Refactor Old Reconciler](#refactor-old-reconciler)

## Overview

Proposal to fix Issue [#22609](https://github.com/kubernetes/kubernetes/issues/22609)

`kube-apiserver` currently has a command-line argument `--apiserver-count`
specifying the number of api servers. This masterCount is used in the
MasterCountEndpointReconciler on a 10 second interval to potentially cleanup
stale API Endpoints. The issue is when the number of kube-apiserver instances
gets below or above the masterCount. If the below case happens, the stale
instances within the Endpoints does not get cleaned up, or in the latter case
the endpoints start to flap.
## Known Issues

Each apiserver’s reconciler only cleans up for it's own IP. If a new
server is spun up at a new IP, then the old IP in the Endpoints list is
only reclaimed if the number of apiservers becomes greater-than or equal
to the masterCount. For example:

* If the masterCount = 3, and there are 3 API servers running (named: A, B, and C) 
* ‘B’ API server is terminated for any reason
* The IP for endpoint ‘B’ is not
removed from the Endpoints list

There is logic within the [MasterCountEndpointReconciler](https://github.com/kubernetes/kubernetes/blob/68814c0203c4b8abe59812b1093844a1f9bdac05/pkg/master/controller.go#L293) to attempt to make
the Endpoints eventually consistent, but the code relies on the
Endpoints count becoming equal to or greater than masterCount. When the
apiservers become greater than the masterCount the Endpoints tend to
flap.

If the number endpoints were scaled down from automation, then the
Endpoints would never become consistent.

## Proposal

### Create New Reconciler

| Kubernetes Release  | Quality | Description |
| ------------- | ------------- | ----------- |
| 1.9           | alpha         | <ul><li>Add a new reconciler</li><li>Add a command-line type `--alpha-apiserver-endpoint-reconciler-type`<ul><li>configmap</li><li>default</li></ul></li></ul>
| 1.10          | beta          | <ul><li>Turn on the `configmap` type by default</li></ul>
| 1.11          | stable        | <ul><li>Remove code for old reconciler</li><li>Remove --apiserver-count</li></ul>

The MasterCountEndpointReconciler does not meet the current needs for
durability of API Endpoint creation, deletion, or failure cases.

Create a new `MasterEndpointConfigMapReconciler` within
master/controller.go.

Add a `kube-apiserver-config` ConfigMap in the `kube-system`
namespace. The duration found within the Coordination map would be configurable by
admins without a recompile. The ConfigMap would include the following:

```go
KubeApiServerConfigMap{
  "expiration-duration":        "1m",
  ... [other flags potentially]
}
```

The Coordination ConfigMap would be formed such that:

```go
CoordinationConfigMap{
  "[namespace]/[name]": "[KubeAPIServerEndpoint]",
}
```

*Key Format*: ip-[IP String Formatted]-[port]

```go
KubeAPIServerEndpointConfigMap{
  "master-count":                "0",
	"ip-2001-4860-4860--8888-443": "time.Time",
	"ip-192-168-0-3-443":          "time.Time",
}
```

***TODO: should it a serialized struct, or the raw time value?***

The reconcile loop will expire endpoints that do not meet the duration.
On each reconcile loop (the loop runs every 10 seconds currently,
but interval will be changed to 80% of the `expiration-duration`):

1. GET `kube-apiserver-endpoints` ConfigMap (as endpointMap)
1. Update the timestamp for the currently running API server
1. Remove all endpoints where the timestamp is greater than `expiration-duration` from the configMap.
1. Do a GET on the `kube-apiserver-endpoints`
1. Construct a PATCH operation with the deleted endpoints

configmap.yml:

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-apiserver-endpoints
  namespace: default
```

### Alternate Proposals

#### Custom Resource Definitions

CRD's were considered for this proposal, but were not proposed due to
constraints of having CRDs within core has not been clearly defined.
Layering could also be an issue, so the proposal defined ConfigMaps as
the current path forward.

#### Refactor Old Reconciler

| Release | Quality |                         Description                          |
| ------- | ------- | ------------------------------------------------------------ |
| 1.9     | stable  | Change the logic in the current reconciler

We could potentially reuse the old reconciler by changing the reconciler to count
the endpoints and set the `masterCount` (with a RWLock) to the count.
