Metrics Server
==============

Resource Metrics API is an effort to provide a first-class Kubernetes API
(stable, versioned, discoverable, available through apiserver and with client support)
that serves resource usage metrics for pods and nodes. The use cases were discussed
and the API was proposed a while ago in
[another proposal](/contributors/design-proposals/instrumentation/resource-metrics-api.md).
This document describes the architecture and the design of the second part of this effort:
making the mentioned API available in the same way as the other Kubernetes APIs.

### Scalability limitations ###
We want to collect up to 10 metrics from each pod and node running in a cluster.
Starting with Kubernetes 1.6 we support 5000 nodes clusters with 30 pods per node.
Assuming we want to collect metrics with 1 minute granularity this means:
```
10 x 5000 x 30 / 60 = 25000 metrics per second by average
```
 
Kubernetes apiserver persists all Kubernetes resources in its key-value store [etcd](https://coreos.com/etcd/).
It’s not able to handle such load. On the other hand metrics tend to change frequently,
are temporary and in case of loss of them we can collect them during the next housekeeping operation.
We will store them in memory then. This means that we can’t reuse the main apiserver
and instead we will introduce a new one - metrics server.

### Current status ###
The API has been already implemented in Heapster, but users and Kubernetes components
can only access it through master proxy mechanism and have to decode it on their own.
Heapster serves the API using go http library which doesn’t offer a number of functionality
that is offered by Kubernetes API server like authorization/authentication or client generation.
There is also a prototype of Heapster using [generic apiserver](https://github.com/kubernetes/apiserver) library.
 
The API is in alpha and there is a plan to graduate it to beta (and later to GA),
but it’s out of the scope of this document.

### Dependencies ###
In order to make metrics server available for users in exactly the same way
as the regular Kubernetes API we need a mechanism that redirects requests to `/apis/metrics`
endpoint from the apiserver to metrics server. The solution for this problem is
[kube-aggregator](https://github.com/kubernetes/kube-aggregator).
The effort is on track to be completed for Kubernetes 1.7 release.
Previously metrics server was blocked on this dependency.

### Design ###
Metrics server will be implemented in line with
[Kubernetes monitoring architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md)
and inspired by [Heapster](https://github.com/kubernetes/heapster).
It will be a cluster level component which periodically scrapes metrics from all Kubernetes nodes
served by Kubelet through Summary API. Then metrics will be aggregated, 
stored in memory (see Scalability limitations) and served in
[Metrics API](https://git.k8s.io/metrics/pkg/apis/metrics/v1alpha1/types.go) format.
 
Metrics server will use apiserver library to implement http server functionality.
The library offers common Kubernetes functionality like authorization/authentication,
versioning, support for auto-generated client. To store data in memory we will replace
the default storage layer (etcd) by introducing in-memory store which will implement
[Storage interface](https://git.k8s.io/apiserver/pkg/registry/rest/rest.go).
 
Only the most recent value of each metric will be remembered. If a user needs an access
to historical data they should either use 3rd party monitoring solution or
archive the metrics on their own (more details in the mentioned vision).
 
Since the metrics are stored in memory, once the component is restarted, all data are lost.
This is an acceptable behavior because shortly after the restart the newest metrics will be collected,
though we will try to minimize the priority of this (see also Deployment).

### Deployment ###
Since metrics server is prerequisite for a number of Kubernetes components (HPA, scheduler, kubectl top)
it will run by default in all Kubernetes clusters. Metrics server initiates connections to nodes,
due to security reasons (our policy allows only connection in the opposite direction) so it has to run on user’s node.
 
There will be only one instance of metrics server running in each cluster. In order to handle
high metrics volume, metrics server will be vertically autoscaled by
[addon-resizer](https://git.k8s.io/contrib/addon-resizer).
We will measure its resource usage characteristic. Our experience from profiling Heapster shows
that it scales vertically effectively. If we hit performance limits we will consider scaling it 
horizontally, though it’s rather complicated and is out of the scope of this doc.
 
Metrics server will be Kubernetes addon, create by kube-up script and managed by
[addon-manager](https://git.k8s.io/kubernetes/cluster/addons/addon-manager).
Since there are a number of dependent components, it will be marked as a critical addon.
In the future when the priority/preemption feature is introduced we will migrate to use this
proper mechanism for marking it as a high-priority, system component.

### Users migration ###
In order to make the API usable we will provide auto-generated set of clients.
Currently the API is being used by a number of components and after we will introduce
the metrics server we will migrate all of them to use the new path.

