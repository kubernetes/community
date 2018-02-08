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

There is logic within the
[MasterCountEndpointReconciler](https://github.com/kubernetes/kubernetes/blob/68814c0203c4b8abe59812b1093844a1f9bdac05/pkg/master/controller.go#L293)
to attempt to make the Endpoints eventually consistent, but the code relies on
the Endpoints count becoming equal to or greater than masterCount. When the
apiservers become greater than the masterCount the Endpoints tend to flap.

If the number endpoints were scaled down from automation, then the
Endpoints would never become consistent.

## Proposal

### Create New Reconciler

| Kubernetes Release  | Quality | Description |
| ------------- | ------------- | ----------- |
| 1.9           | alpha         | <ul><li>Add a new reconciler</li><li>Add a command-line type `--alpha-apiserver-endpoint-reconciler-type`<ul><li>storage</li><li>default</li></ul></li></ul>
| 1.10          | beta          | <ul><li>Turn on the `storage` type by default</li></ul>
| 1.11          | stable        | <ul><li>Remove code for old reconciler</li><li>Remove --apiserver-count</li></ul>

The MasterCountEndpointReconciler does not meet the current needs for durability
of API Endpoint creation, deletion, or failure cases.

Custom Resource Definitions were proposed, but they do not have clean layering.
Additionally, liveness and locking would be a nice to have feature for a long
term solution.

ConfigMaps were proposed, but since they are watched globally, liveliness
updates could be overly chatty.

By porting OpenShift's
[LeaseEndpointReconciler](https://github.com/openshift/origin/blob/master/pkg/cmd/server/election/lease_endpoint_reconciler.go)
to Kubernetes we can use the Storage API directly to store Endpoints
dynamically within the system.

### Alternate Proposals

#### Custom Resource Definitions and ConfigMaps

CRD's and ConfigMaps were considered for this proposal. They were not adopted
for this proposal by the community due to technical issues explained earlier.

#### Refactor Old Reconciler

| Release | Quality |                         Description                          |
| ------- | ------- | ------------------------------------------------------------ |
| 1.9     | stable  | Change the logic in the current reconciler

We could potentially reuse the old reconciler by changing the reconciler to count
the endpoints and set the `masterCount` (with a RWLock) to the count.
