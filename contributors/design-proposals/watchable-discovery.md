# Watchable API Group Discovery

The Kubernetes API server provides a number of HTTP endpoinds to
discover the available API groups, versions and resources:

- `/apis`: returns a list of groups with available versions (compare [metav1.APIGroupList](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L639))
- `/apis/<group>`: returns a list of versions, **without** resources (compare [metav1.APIVersions](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L623))
- `/apis/<group>/<version>`: returns a list of resources (compare [metav1.APIResourceList](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L718))

This discovery information is used by many API clients like kubectl
and controllers like the namespace controller. The garbage
collection controller is supposed to use discovery as well, [but
doesn't yet](https://github.com/kubernetes/kubernetes/blob/6f5337cff7baffc2b27d85036cb1ecb6a4c1f73b/cmd/kube-controller-manager/app/core.go#L174).

Kubectl will cache the discovery information in
`~/.kube/cache` for ten minutes and queries the API server again after
that time or on a cache miss on RESTMapper lookups. In the namespace controller
a [polling logic is implemented](https://github.com/kubernetes/kubernetes/blob/6f5337cff7baffc2b27d85036cb1ecb6a4c1f73b/cmd/kube-controller-manager/app/core.go#L115).

With ThirdPartyResources and API aggregation the discovery information
is not static anymore within a deployed Kubernetes version. It can
change at any time, either because additional API servers are added to
or removed from the cluster, or because the user creates a
ThirdPartyResource. For consistent handling of the new resources in
kubectl and the controllers, they have to adapt to the changed API
groups, prefereably in real-time or at least seconds.

The discovery API is not tailored for many API groups and versions. To
get a list of all GroupVersionResources in the cluster, a client has
to query the available groups (with their versions through `/apis`)
and then the resources for each GroupVersion. In addition, retry logic
is involved to make this reliable. This can easily lead to
20 and more discovery requests, slowing down especially an interactive
tool like kubectl. This was the reason to add the cache as a
workaround, with the danger that discovery infos is outdates for up to
10 minutes.

The discovery API is not versioned. Incompatible changes will break
old client therefore.

The discovery API does not use the API semantics of API group resource
endpoints. I.e. it is not possible to query with selectors, to watch
for changes or to use protobuf for more efficient transport. Moreover,
discovery info does not carry a `ResourceVersion` such that it's
impossible to compare cache information with the last version in the
API server.

## Goals

- near real-time updates without polling
- better scalability to more groups and versions with less HTTP requests
  - one request which returns all GroupVersionResources
- versioning
- default Kubernetes API semantics including:
  - watchability
  - resource versioning
  - protobuf support
  - compatibility with all existing machinery like clients,
    ListWatches, informers, caching, etc.
- compatibility with kube-aggregator, third-party apiservers and
  ThirdPartyResources

### Non-goals

- different view for different clients: e.g. kubectl 1.5 won't see
  `batch/v1`
- ETag like response without content for even better scalability and
  responsiveness
- readability of certain `Group` objects for unauthenticated users
  (e.g. to show those APIs used by kubelet and kubeadm bootstrapping
  before both-way trust is established).

### Desired API Semantics

The discovery information should be served like any other resource in
the Kubernetes API, i.e. as its own resource within a versioned API
group.

The API should be read-only, i.e. only GET, LIST and WATCH will be
allowed verbs.

The default RBAC rules will make the resources readable for every
authenticated user.

It would be possible to add a flag for unauthenticated readability,
i.e. as a counterpart to the `kube-public` API group that is used
during kubelet bootstrapping. But this is a non-goal for now.

The discovery information includes its own API group.

## New API Types

The proposed API group name is `discovery.k8s.io`, starting with a
`v1alpha1` version.

The only new Kind is `Group`, containing all version information and
all resources. The name of group objects is the API group name. The
Kind is cluster-scoped.

```golang
// +genclient=true
// +nonNamespaced=true

// Group describes an API group with all its versions and resources.
type Group struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Status of the group.
	Status GroupStatus `json:"spec,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// GroupStatus describes the API group with all its versions and resources.
type GroupStatus struct {
	// versions are the versions supported in this group, order by preference. The first one is the preferred
	// version for the group.
	Versions []GroupVersion `json:"versions" protobuf:"bytes,1,rep,name=versions"`
	// a map of client CIDR to server address that is serving this group.
	// This is to help clients reach servers in the most network-efficient way possible.
	// Clients can use the appropriate server address as per the CIDR that they match.
	// In case of multiple matches, clients should use the longest matching CIDR.
	// The server returns only those CIDRs that it thinks that the client can match.
	// For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP.
	// Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
	ServerAddressByClientCIDRs []metav1.ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs" protobuf:"bytes,2,rep,name=serverAddressByClientCIDRs"`
}

// GroupVersion describes one version of an API group, including all resources available in this version.
type GroupVersion struct {
	// the version name
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// resources contains the name of the resources and if they are namespaced.
	Resources []metav1.APIResource `json:"resources" protobuf:"bytes,2,rep,name=resources"`
}

// GroupList is a list of API groups.
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of group.
	Items []Group `json:"items" protobuf:"bytes,2,rep,name=items"`
}
```

## Behavior

Each API server in the cluster serves its own `discovery.k8s.io` API
group. By default, this is backed by normal etcd storage, i.e. the
semantics of the `ResourceVersion` is the one of the underlaying etcd
storage backend.

Each API server has its own `DiscoveryUpdateController` that's
regularly reconciles the storage state of the discovery API group with
its internal registered/announce API group registry. Moreover, when
a new API groups is installed or removed in its HTTP handler, the
corresponding API group object is updated synchronously.

### Aggegration

The aggregator watches the discovery API group of each of its proxied
API servers. It filters the returned discovery information by taking
the `APIService` status into consideration, i.e. by 
- only showing those GroupVersions that are registered via an
  `APIService`
- only showing those GroupVersions that are accepted by the aggregator
  (i.e. which have no naming conflict).

The aggregator will update its own discovery API group resources
triggered by changes in the sub-API-servers.

### Embedded Aggregator in kube-apiserver

In this mode, both the aggregator and the kube-apiserver share the
storage. To avoid conflicts about updating the API group, the `Group`
objects of kube-apiserver are marked as "local" with an
annotation. The aggregator will not touch them and the other way
around for the kube-apiserver for the not-marked `Group` objects.

### HA

API servers in HA mode will share a storage. Hence, they will share
the state of the discovery API group. Their controllers will fight for
the correct state of the discovery information.

During normal operation the desired discovery state is equal in an HA
setup. So no real race will take place.

During rolling updates of the control plane, an old and a new API
server version might be active. Their reconciliation loops will race
to update the stored groups.

The consumers of the discovery information might notice the
flipping. But, this is not a regression of the cluster behaviour as it
already happens now if a load-balancer distributes requests among the
cluster API servers, each returning different discovery information.

A long reconciliation period will reduce the effect of the flipping,
e.g. an API server update shouldn't take longer than a minute.

We could introduce a server version for each API group. Then an older
server could be less aggressive about reconciliation of new
groups. But, this is not part of the API types yet.

### Alternative Architectures

- We could implement the discovery resource in the aggregator only and
  use polling to the sub-API-servers. This would look mostly the same
  for the cluster end-user, but this would have the following down-sides:
  1. it's **less real-time** if no watch is involved to trigger a new re-poll.
	 We could watch TPRs (the old and the new kind) in the aggregator
     and assume that no other third-party API server has dynamic API
	 groups. Having this knowledge in the aggregator is ugly, but might
	 be acceptable if it reduces complexity or implementation effort.
  1. **watchable discovery won't work without the aggregator**. This is
     probably mostly an issue during testing. In a real setup,
     everything will go through the aggregator, especially because it
     does authentication.
  The advantage is that 
  1. we don't have to implement the DiscoveryController in normal
     generic storage based API server (as kube-apiserver,
     service-catalog and probably most other third-party API
     servers). This though is a one-time implementation in
     `k8s.io/apiserver` and then two calls per API server:
	 - `genericApiServer.InstallDiscoveryAPI()`
	 - `genericApiServer.StartDiscoveryController(loopBackClient)`
  1. we don't have to implement the DiscoveryController in the
	 new kube-apiextensions-apiserver.
  The prototype will show whether these two steps actually safe much
  effort or reduce complexity. It is expected that the changes will
  be pretty straight forward.
- We could implement an in-memory variant of the discovery
  resource. While this is considerable more complex (painful
  experience exists in the OpenShift project), we still have to
  implement the always-ingreasing behavior of the `ResourceVersion`
  field. In a distributed environment, this requires some kind of
  state. This state would naturally be provided by etcd. At this
  points, it's not that far from storing the `Group` objects
  persistently in etcd and inherit all the machinery. E.g. we can use
  normal informers to watch existing API groups, versions and
  resources.
- We could extend the current endpoints to be watchable (the main goal
  of this proposal). But this would mean to reinvent a lot of
  infrastructure we already have.
- We could re-use the `meta.k8s.io` group name instead of
  `discovery.k8s.io`. At the end, this is probably a matter of
  taste. Both would work. But because the new Kind is exposed as an
  endpoint in contrast to all the other meta types, the later seems to
  be more appropriate. After all, the `Group` is describes something
  on the meta level, but it is a real, concrete Kind.
