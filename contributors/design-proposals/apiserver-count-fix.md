# apiserver-count fix proposal

Authors: @rphillips

## Table of Contents

1. [Overview](#overview)
2. [Known Issues](#known-issues)
3. [Proposal](#proposal)
4. [Prior Art](#prior-art)

## Overview

Proposal to fix Issue [#22609](https://github.com/kubernetes/kubernetes/issues/22609)

`kube-apiserver` currently has a command-line argument `--apiserver-count`
specifying the number of api masters. This masterCount is used in the
MasterCountEndpointReconciler on a 10 second interval to potentially cleanup
stale API Endpoints. The issue is when the number of kube-apiserver instances
gets below masterCount. If this case happens, the stale instances within the
Endpoints does not get cleaned up.

## Known Issues

Each apiserver’s reconciler only cleans up for it's own IP. If a new server
is spun up at a new IP, then the old IP in the Endpoints list is only
reclaimed if the number of apiservers becomes greater-than or equal to the
masterCount. For example:

* If the masterCount = 3, and there are 3 API servers running (named: A, B, and
C) 
* ‘B’ API server is terminated for any reason
* The IP for endpoint ‘B’ is not
removed from the Endpoints list

There is logic within the [MasterCountEndpointReconciler](https://github.com/kubernetes/kubernetes/blob/68814c0203c4b8abe59812b1093844a1f9bdac05/pkg/master/controller.go#L293) to attempt to make
the Endpoints eventually consistent, but the code relies on the Endpoints
count becoming equal to or greater than masterCount. When the apiservers
become greater than the masterCount the Endpoints tend to flap.

If the number endpoints were scaled down from automation, then the Endpoints
would never become consistent. 

## Proposal

### Create New Reconciler

| Kubernetes Release  | Quality | Description |
| ------------- | ------------- | ----------- |
| 1.9           | alpha         | <ul><li>Add a new reconciler</li><li>Add a command-line switch --new-reconciler</li><li>Add a command-line switch --old-reconciler</li></ul>
| 1.10          | beta          | <ul><li>Turn on the new reconciler by default |</li></ul>
| 1.11          | stable        | <ul><li>Remove code for old reconciler</li><li>Remove --old-reconciler</li><li>Remove --new-reconciler</li><li>Remove --apiserver-count</li></ul>

The MasterCountEndpointReconciler does not meet the current needs for durability of API Endpoint creation, deletion, or failure cases.

Create a new `MasterEndpointReconciler` within master/controller.go.

Add a `kube-apiserver-endpoints-config` ConfigMap in the `default` namespace. The duration found within the map would be configurable by admins without a recompile. The ConfigMap would include the following:

```go
ConfigMap{
	"expire-duration": "1m", // golang duration (ns,us,ms,s,m,h)
}
```

Add a standard `kube-apiserver-endpoints` ConfigMap in the `default` namespace. The ConfigMap would be formed such that: 

*Key Format*: ip-[IP String Formatted]-[port]

```go
ConfigMap{
	"ip-2001-4860-4860--8888-443": "serialized JSON ControllerEndpointData",
	"ip-192-168-0-3-443":          "serialized JSON ControllerEndpointData",
}

type ControllerEndpointData struct {
     metav1.TypeMeta   `json:",inline"`
     metav1.ObjectMeta `json:",inline"`
     api.EndpointPort
     CreationTimestamp Time `json:"creationTimestamp,omitempty" protobuf:"bytes,8,opt,name=creationTimestamp"`
     UpdateTimestamp   Time `json:"creationTimestamp,omitempty" protobuf:"bytes,8,opt,name=updateTimestamp"`
}
```

The reconcile loop will expire endpoints that do not meet the duration. On
each reconcile loop (the loop runs every 10 seconds currently):

1. Retrieve `kube-apiserver-endpoints-config` ConfigMap (as configMap)
1. Retrieve `kube-apiserver-endpoints` ConfigMap (as endpointMap)
1. Update the `UpdateTimestamp` for the currently running API server
1. Remove all endpoints where the UpdateTimestamp is greater than `expire-duration` from the configMap.
1. Write endpointMap back to Kubernetes ConfigMap API

configmap.yml:

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-apiserver-endpoints
  namespace: default
```

### Refactor Old Reconciler

| Kubernetes Release | Quality | Description |
| ------------- | ------------- | ----------- |
| 1.9           | stable  | <ul><li>Change the logic in the current reconciler</li></ul>

We could potentially reuse the old reconciler, but ignore the masterCount and change the logic to use the proposal from the previous section.

## Prior Art

[Security Labeller](https://github.com/coreos-inc/security-labeller/issues/18#issuecomment-320791878)
