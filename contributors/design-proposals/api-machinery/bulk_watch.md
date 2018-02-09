# Background

 As part of increasing security of a cluster, we are planning to limit the
ability of a given Kubelet (in general: node), to be able to read only
resources associated with it. Those resources, in particular means: secrets,
configmaps & persistentvolumeclaims. This is needed to avoid situation when
compromising node de facto means compromising a cluster. For more details &
discussions see https://github.com/kubernetes/kubernetes/issues/40476.
 
 However, by some extension to this effort, we would like to improve scalability
of the system, by significantly reducing amount of api calls coming from
kubelets. As of now, to avoid situation that kubelet is watching all secrets/
configmaps/... in the system, it is not using watch for this purpose. Instead of
that, it is retrieving individual objects, by sending individual GET requests.
However, to enable automatic updates of mounted secrets/configmaps/..., Kubelet
is sending those GET requests periodically. In large clusters, this is
generating huge unnecessary load, as this load in principle should be
watch-based. We would like to address this together with solving the
authorization issue.


# Proposal

 In this proposal, we're not focusing on how exactly security should be done.
We're just sketching very high level approach and exact authorization mechanism
should be discussed separately.

 At the high level, what we would like to achieve is to enable LIST and WATCH
requests to support more sophisticated filtering (e.g. we would like to be able
to somehow ask for all Secrets attached to all pods bound to a given node in a
single LIST or WATCH request). However, the design has to be consistent with
authorization of other types of requests (in particular GETs).

 To solve this problem, we propose to introduce the idea of `bulk watch ` and
`bulk get ` (and in the future other bulk operations). This idea was already
appearing in the community and now we have good usecase to proceed with it.

 Once a bulk watch is set up, we also need to periodically verify its ACLs.
Whenever a user (in our case Kubelet) loses access to a given resource, the
watch should be closed within some bounded time. The rationale behind this
requirement is that by using (bulk) watch we still want to enforce ACLs
similarly to how we enforce them with get operations (and that Kubelet would
eventually be unable to access a secret no matter if it is watching or
polling it).

 That said, periodic verification of ACLs isn't specific to bulk watch and
needs to be solved also in `regular` watch (e.g. user watching just a single
secret may also lose access to it and such watch should also be closed in this
case). So this requirement is common for both regular and bulk watch. We
just need to solve this problem on low enough level that would allow us to
reuse the same mechanism in both cases - we will solve it by sending an error
event to the watcher and then just closing this particular watch.

 At the high level, we would like the API to be generic enough so that it will
be useful for many different usecases, not just this particular one (another
example may be a controller that needs to deal with just subset of namespaces).
As a result, below we are describing requirements that the end solution has to
meet to satisfy our needs
- a single bulk requests has to support multiple resource types (e.g. get a
node and all pods associated with it)
- the wrappers for aggregating multiple objects (in case of list we can return
a number of objects of different kinds) should be `similar` to lists in core
API (by lists I mean e.g. `PodList` object)
- the API has to be implemented also in aggregator so that bulk operations
are supported also if different resource types are served by different
apiservers
- clients has to be able to alter their watch subscriptions incrementally (it
may not be implemented in the initial version though, but has to be designed)


# Detailed design

 As stated in above requirements, we need to make bulk operations work across
different resource types (e.g. watch pod P and secret S within a single watch
call). Spanning multiple resources, resource types or conditions will be more
and more important for large number of watches. As an example, federation will
be adding watches for every type it federates. With that in mind, bypassing
aggregation at the resource type level and going to aggregation over objects
with different resource types will allow us to more aggressively optimize in the
future (it doesn't mean you have to watch resources of different types in a
single watch, but we would like to make it possible).

 That means, that we need to implement the bulk operation at the aggregation
level. The implications of it are discussed below.

 Moreover, our current REST API doesn't even offer an easy way to handle
"multiple watches of a given type" within a single request. As a result, instead
of inventing new top level pattern per type, we can introduce a new resource
type that follows normal RESTful rules and solves even more generic problem
of spanning multiple different resource types.

 We will start with introducing a new dedicated API group:
 ```
 /apis/bulk.k8s.io/
 ```
 that underneath will have a completely separate implementation.

 In all text below, we are assuming v1 version of the API, but it will obviously
go through alpha and beta stages before (it will start as v1alpha1).

 In this design, we will focus only on bulk get (list) and watch operations.
Later, we would like to introduce new resources to support bulk create, update
and delete operations, but that's not part of this design.

 We will start with introducing `bulkgetoperations` resource and supporting the
following operation:
```
POST /apis/bulk.k8s.io/v1/bulkgetoperations <body defines filtering>
```
 We can't simply make this an http GET request, due to limitations of GET for
the size (length) of the url (in which we would have to pass filter options).

 We could consider adding `watch` operation using the same pattern with just
`?watch=1` parameter. However, the main drawback of this approach is that it
won't allow for dynamic altering of watch subscriptions (which we definitely
also need to support).
As a result, we need another API for watch that will also support incremental
subscriptions - it will look as following:
```
websocket /apis/bulk.k8s.io/v1/bulkgetoperations?watch=1
```

*Note: For consistency, we also considered introducing websocket API for
handling LIST requests, where first client sends a filter definition over the
channel and then server sends back the response, but we dropped this for now.*

*Note: We also considered implementing the POST-based watch handler that doesn't
allow for altering subscriptions, which should be very simple once we have list
implemented. But since websocket API is needed anyway, we also dropped it.*


### Filtering definition

 We need to make our filtering mechanism to support different resource types at
the same time. On the other hand, we would like it to be as consistent with all
other Kubernetes APIs as possible. So we define the selector for bulk operations
as following:

```
type BulkGetOperation struct {
	Operations []GetOperation
}

type GetOperation struct {
	Resource GroupVersionResource

	// We would like to reuse the same ListOptions definition as we are using
	// in regular APIs.
	// TODO: We may consider supporting multiple ListOptions for a single
	// GetOperation.
	Options ListOptions
}
```

 We need to be able to detect whether a given user is allowed to get/list
objects requested by a given "GetOperation". For that purpose, we will create
some dedicated admission plugin (or potentially reuse already existing one).
That one will not support partial rejections and will simply allow or reject
the whole request.

For watch operations, as described in requirements, we also need to periodically
(or maybe lazily) verify whether their ACLs didn't change (and user is still
allowed to watch requested objects). However, as also mentioned above, this
periodic checking isn't specific to bulk operations and we need to support the
same mechanism for regular watch too. We will just ensure that this mechanism
tracking and periodically verifying ACLs is implemented low enough in apiserver
machinery so that we will be able to reuse exactly the same one for the purpose
of bulk watch operations.
For watch request, we will support partial rejections. The exact details of it
will be described together with dynamic watch description below.


### Dynamic watch

 As mentioned in the Proposal section, we will implement bulk watch that will
allow for dynamic subscription/unsubscription for (sets of) objects on top of
websockets protocol.

 Note that we already support websockets in the regular Kubernetes API for
watch requests (in addition to regular http requests), so for the purpose of
bulk watch we will be extending websocket support.

 The high level, the protocol will look:
1. client opens a new websocket connection to a bulk watch endpoint to the
server via ghttp GET
1. this results in creating a single channel that is used only to handle
communication for subscribing/unsubscribing for watches; no watch events are
delivered via this particular channel.
*TODO: Consider switching to two channels, one for incoming and one for
outgoing communication*
1. to subscribe for a watch of a given (set of) objects, user sends `Watch`
object over the channel; in response a new channel is created and the message
with the channel identifier is send back to the user (we will be using integers
as channel identifiers).
*TODO: Check if integers are mandatory or if we can switch to something like
ch1, ch2 ... .*
1. once subscribed, all objects matching a given selector will be send over
the newly created channel
1. to stop watching for a given (set of) objects, user sends `CloseWatch`
object over the channel; in response the corresponding watch is broken and
corresponding channel within websocket is closed
1. once done, user should close the whole websocket connection (this results in
breaking all still opened channels and corresponding watches).

 With that high level protocol, there are still multiple details that needs
to be figured out. First, we need to define `Watch` and `CloseWatch`
message. We will solve it with the single `Request` object:
```
type Request struct {
	// Only one of those is set.
	Watch      *Watch
	CloseWatch *CloseWatch
}

type Identifier int64

// Watch request for objects matching a given selector.
type Watch struct {
	Selector GetOperation
}

// Request to stop watching objects from a watch identified by the channel.
type CloseWatch struct {
	Channel Identifier
}

// Depending on the request, channel that was created or deleted.
type Response struct {
	Channel Identifier
}
```
With the above structure we can guarantee that we only send and receive
objects of a single type over the channel.

We should also introduce some way of correlating responses with requests
when a client is sending multiple of them at the same time. To achieve this
we will add a `request identified` field to the `Request` that user can set
and that will then be returned as part of `Response`. With this mechanism
user can set the identifier to increasing integers and then will be able
to correlate responses with requests he sent before. So the final structure
will be as following:
```
type Request struct {
	ID Identifier
	// Only one of those is set.
	Watch      *Watch
	CloseWatch *CloseWatch
}

// Depending on the request, channel that was created or deleted.
type Response struct {
	// Propagated from the Request.
	RequestID Identifier
	Channel   Identifier
}
```

Another detail we need to point at is about semantic. If there are multiple
selectors selecting exactly the same object, the object will be send multiple
times (once for every channel interested in that object).
If we decide to also have http-POST-based watch, since there will be basically
a single channel there (as it is in watch in regular API), such an object will
be send once. This semantic difference needs to be explicitly described.

*TODO: Since those channels for individual watches are kind of independent,
we need to decide whether they are supposed to be synchronized with each other.
In particular, if we have two channels watching for objects of the same type
are events guaranteed to be send in the increasing order of resource versions.
I think this isn't necessary, but it needs to be explicit.*

Yet another thing to consider is what if the server wants to close the watch
even though user didn't request it (some potential reasons of it may be failed
periodic ACL check or some kind of timeout).
We will solve it by saying that in such situation we will send an error via
the corresponding channel and every error automatically closes the channel.
It is responsibility of a user to re-subscribe if that's possible.

We will also reuse this mechanism for partial rejections. We will be able to
reject (or close failed periodic ACL check) any given channel separately from
all other existing channels.


### Watch semantics

 There are a lot of places in the code (including all our list/watch-related
frameworks like reflector) that rely on two crucial watch invariants:
1. watch events are delivered in increasing order of resource version
1. there is at most one watch event delivered for any resource version

 However, we have no guarantee that resource version series is shared between
different resource types (in fact in default GCE setup events are not sharing
the same series as they are stored in a separate etcd instance). That said,
to avoid introducing too many assumptions (that already aren't really met)
we can't guarantee exactly the same.

 With the above description of "dynamic watch", within a single channel you
are allowed to only watch for objects of a single type. So it is enough to
introduce only the following assumption:
1. within a single resource type, all objects are sharing the same resource
version series.

 This means, we can still shard etcd by resource types, but we can't really
shard by e.g. by namespaces. Note that this doesn't introduce significant
limitations compared to what we already have, because even now you can watch
all objects of a single type and there is no mechanism to inject multiple
resource versions into it. So this assumption is not making things worse.

 To support multi-resource-type watches, we can build another framework on
top of frameworks we already have as following:
- we will have a single reflector/informer/... per resource type
- we will create an aggregator/dispatcher in front of them that will be
responsible for aggregating requests from underlying frameworks into a single
one and then dispatching incoming watch events to correct reflect/informer.
This will obviously require changes to existing frameworks, but those should
be local changes.

 One more thing to mention is detecting resource type of object being send via
watch. With "dynamic watch" proposal, we already know it based on the channel
from which it came (only objects of a single type can be send over the single
channel).

 Note that this won't be true if we would decide for regular http-based watch
and as a result we would have to introduce a dedicated type for bulk watch
event containing object type. This is yet another argument to avoid implementing
http-based bulk watch at all.


### Implementation details

 As already mentioned above, we need to support API aggregation. The API
(endpoint) has to be implemented in kube-aggregator. For the implementation
we have two alternatives:
1. aggregator forwards the request to all apiservers and aggregates results
2. based on discovery information, aggregator knows which type is supported
by which apiserver so it is forwarding requests with just appropriate
resource types to corresponding apiservers and then aggregates results.

Neither of those is difficult, so we should proceed with the second, which
has an advantage for watch, because for a given subrequest only a single
apiserver will be returning some events. However, no matter which one we
choose, client will not see any difference between contacting apiserver or
aggregator, which is crucial requirement here.

NOTE: For watch requests, as an initial step we can consider implementing
this API only in aggregator and simply start an individual watch for any
subrequest. With http2 we shouldn't get rid of descriptors and it can be
enough as a proof of concept. However, with such approach there will be
difference between sending a given request to aggregator and apiserver
so we need to implement it properly in apiserver before entering alpha
anyway. This would just give us early results faster.

 The implementation of bulk get and bulk watch in a single apiserver will
also work as kind of aggregator. Whenever a request is coming, it will:
- check what resource type(s) are requested in this request
- for every resource type, combine only parts of the filter that are about
this particular resource type and send the request down the stack
- gather responses for all those resource types and combine them into
single response to the user.

 The only non-trivial operation above is sending the request for a single
resource type down the stack. In order to implement it, we will need to
slightly modify the interface of "Registry" in apiserver. The modification
will have to allow passing both what we are passing now and BulkListOptions
(in some format) (this may e.g. changing signature to accept BulkListOptions
and translating ListOptions to BulkListOptions in the current code).

 With those changes in place, there is a question of how to call this
code. There are two main options:
1. Make each registered resource type, register also in BulkAggregator
(or whatever we call it) and call those methods directly
1. Expose also per-resource bulk operation in the main API, e.g.:
```
POST /api/v1/namespace/default/pods/_bulkwatch <body defines filtering>
```
and use the official apiserver API for delegating requests. However, this
may collide with resource named `_bulkwatch ` and detecting whether
this is bulk operation or regular api operation doesn't seem to be worth
pursuing.

As a result, we will proceed with option 1.


## Considered alternatives

 We considered introducing a different way of filtering that would basically be
"filter only objects that the node is allowed to see". However, to make this
kind of watch work correctly with our list/watch-related frameworks, it would
need to preserve the crucial invariants of watch. This in particular means:

1. There is at most one watch event for a given resource version.
2. Watch events are delivered in the increasing order of resource versions.

 Ideally, whenever a new pod referencing object X is bound to a node, we send
"add" event for object X to the watcher. However, that would break the above
assumptions because:

1. There can be more objects referenced by a given pod (so we can't send all
of them with the rv corresponding to that pod add/update/delete)
2. If we decide for sending those events with their original resource version,
then we could potentially go back in time.

 As a result, we considered the following alternatives to solve this problems:

1. Don't set the event being result of pod creation/update/deletion.
It would be responsibility of a user to grab current version of all object that
are being referenced by this new pod. And only from that point, events for all
objects being referenced would be delivered to the watcher as long as the pod
existing on the node.

This approach in a way it leaking the watch logic to the watcher, that needs
to duplicate the logic of tracking what objects are referenced.

2. Whenever a new pod is bound to a node (or existing is deleted) we send
all add/delete events of attached object to the watcher with the same resource
version being the one of the modified pod.

In this approach, we violate the first assumption (this is a problem in case
of breaking watch, as we don't really know where to resume) as well as we
send events with fake resource versions, which might be misleading to watchers.

3. Another potential option is to change the watch api so that, instead of
sending a single object in a watch event, we would be sending a list of objects
as part of single watch event.

This would solve all the problems from previous two solutions, but this is
change in the api (we would need to introduce a new API for it), and would also
require changes in our list/watch-related frameworks.

 In all above proposals, the tricky part is determining whether an object X
is referenced by any pods bound to a given node is to avoid race conditions and
do it in deterministic way. The crucial requirements are:

1. Whenever "list" request returns a list of objects and a resource version "rv",
starting a watch from the returned "rv" will never drop any events.
2. For a given watch request (with resource version "rv"), the returned stream
of events is always the same (e.g. very slow lagging watch may not cause dropped
events).

We can't really satisfy these conditions using the existing machinery. To solve
this problem reliably we need to be able to serialize events between different
object types in a deterministic way.
This could be done via resource version, but would require assumption that:

1. all object types necessary to determine the in-memory mapping share the same
resource version series.

With that assumption, we can have a "multi-object-type" watch that will serialize
events for different object types for us. Having exactly one watch responsible
for delivering all objects (pods, secrets, ...) will guarantee that if we are
currently at resource version "rv", we processed objects of all types up to rv
and nothing with resource version greater than rv. Which is exactly what we need.


## Other Notes

 We were seriously considering implementing http-POST-based approach as an
additional, simpler watch to implement watch (not supporting altering
subscriptions). In this approach whenever a user want to start watching for
another object (or set of objects) or drop one (or a set), he needs to break the
watch and initiated a new one with a different filter. This approach isn't
perfect, but can solve different usecase of naive clients and is much simpler
to implement.

 However, this has multiple drawbacks, including:
- this is a second mechanism for doing the same thing (we need "dynamic watch"
no matter if we implement it or not)
- there would be some semantic differences between those approaches (e.g. if
there are few selectors selecting the same object, in dynamic approach it will
be send multiple times, once over each channel, here it would be send once)
- we would have to introduce a dedicate "BulkWatchEvent" type to incorporate
resource type. This would make those two incompatible even at the output format.

 With all of those in mind, even though the implementation would be much
simpler (and could potentially be a first step and would probably solve the
original "kubelet watching secrets" problem good enough), we decided not to
proceed with it at all.
