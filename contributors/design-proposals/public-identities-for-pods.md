# Public identities for Pods

* Authors: Sebastian Łaskawiec (@slaskawi)
* Last edit: [2017-03-13](#history)
* Status: design

Approvers:
* [ ] smarterclayton
* [ ] bprashanth
* [ ] enisoc
* [ ] foxish
* [ ] janetkuo
* [ ] kargakis
* [ ] kow3ns
* [ ] thockin

**Table of Contents**

* [Goals](#goals)
* [Non-goals](#non-goals)
* [API](#api)
* [Behavior](#behavior)
* [Upgrading](#upgrading)
* [Implementation](#implementation)
* [Alternatives](#alternatives)
* [History](#history)

# Goals

* The main goal of Public Identities for Pods is to expose each Pod as an external IP address.
  This approach allows sophisticated applications to communicate with the outside world using custom protocols.
  A examples which require this approach are:
  * Game Servers (which often use RTP, UDP)
  * Data Grids (with L7 routing, such as [Infinispan](http://infinispan.org/docs/dev/user_guide/user_guide.html#using_hot_rod_server))
  * Media Servers

* The secondary goal is to create an API similar to `LoadBalancer` to manage external IP addresses.

* The implementation is targeted (but not limited to) StatefulSets. Stable hostnames allow to identify the same
  Stateful Pod after crash or restart (it should have a stable hostname). However it also makes sense to expose
  ordinary deployments this way.

# Existing solution

* The above goal is partially possible to achieve with Loadbalanced Services, NodePort Services or
  Ingress with session stickiness (if HTTP is the main protocol used for communication) but very often
  this is not good enough and application clients assume there is only one application instance
  behind an IP address (in other words there 1 to 1 mapping between IP addresses and Application
  instances). Therefore a specialized implementation exposing an IP address per Pod is required.

# API

A new type of Service needs to be created in order to indicate that an external IP address will be exposed per Pod:

```go
const (
    //Existing Service types
	ServiceTypeClusterIP ServiceType = "ClusterIP"
	ServiceTypeNodePort ServiceType = "NodePort"
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeExternalName ServiceType = "ExternalName"

	//new Service type - expose IP per Pod
	ServiceTypeIPPerPod ServiceType = "IPPerPod"
)
```

A new interface in `cloud.go` will be created to allow creating an external IP addresses.

```go
// ExternalIP is an abstract, pluggable interface for external ips.
type ExternalIP interface {
	// GetExternalIP returns whether the specified external IP exists, and
	// if so, what its status is.
	// Implementations must treat the *v1.Service parameter as read-only and not modify it.
	// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
	GetExternalIP(clusterName string, service *v1.Service) (status *v1.ExternalIPStatus, exists bool, err error)
	// EnsureExternalIP creates a new external IP 'name', or updates the existing one.
	// Returns the status of the external IP.
	// Implementations must treat the *v1.Service
	// parameters as read-only and not modify them.
	// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
	EnsureExternalIP(clusterName string, service *v1.Service) (*v1.ExternalIPStatus, error)
	// EnsureExternalIPDeleted deletes the specified external IP if it
	// exists, returning nil if the external IP specified either didn't exist or
	// was successfully deleted.
	// Implementations must treat the *v1.Service parameter as read-only and not modify it.
	// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
	EnsureExternalIPDeleted(clusterName string, service *v1.Service) error
}
```

Finally a new Annotation will be created to allow falling from "IP per Pod" back into "LoadBalancer per Pod" solution
(more abot this in the following paragraphs).

```go
const FallbackToLoadBalancerPerPod = "service.beta.kubernetes.io/fallback-loadbalancer-per-pod"
```

# Behavior

The logic implemented behind assigning and updating external IPs will be similar to Load Balanced Services.

A ServiceController will periodically iterate though Pods governed by a Service and assign an external IP to each of them.
Note that there is no `UpdateExternalIP` function in `ExternalIP` interface. This means that a single IP address will not
be reassigned into another Pod. It will be explicitly deleted (by `EnsureExternalIPDeleted`).

Since `ExternalIP` interface might not be implemented and supported by all Cloud Vendors, we allow creating a Load Balancer
per Pod. This should however be considered as a last-resort approach. In that case, a user will need to add
`FallbackToLoadBalancerPerPod` annotation on a Service.

# Implementation

Checked items had been completed at the time of the [last edit](#history) of
this proposal.

* [ ] Add `ServiceTypeIPPerPod`, `ExternalIP` and `FallbackToLoadBalancerPerPod`
* [ ] Add proper validation (e.g. a nil `ClusterIP` for the Service etc)
* [ ] Implement logic inside `ServiceController`
* [ ] Implement e2e tests `ServiceController`
* [ ] Add `FallbackToLoadBalancerPerPod` fallback mechanism
* [ ] Implement e2e tests for `FallbackToLoadBalancerPerPod`

# Alternatives

One of the alternatives is to use Load Balancer or Node Port per Pod. This way each Pod will have its own external IP address which will be accessible from the outside world. 
Most of the clustered applications (such as Data Grid for the instance) assume that the port will be the same for all cluster members. This requirement favors Load Balancer per Pod approach however the downside is that Load Balancers might be expensive and abusing them might generate additional costs on the user.

# History

Summary of significant revisions to this document:

* 2017-03-13 (slaskawi)
  * Initial version
* 2017-04-26 (slaskawi)  
  * Addressed comments from review
  * Added Alternatives section

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/controller-ref.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
