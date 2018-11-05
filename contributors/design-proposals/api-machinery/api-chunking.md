# Allow clients to retrieve consistent API lists in chunks

On large clusters, performing API queries that return all of the objects of a given resource type (GET /api/v1/pods, GET /api/v1/secrets) can lead to significant variations in peak memory use on the server and contribute substantially to long tail request latency.

When loading very large sets of objects -- some clusters are now reaching 100k pods or equivalent numbers of supporting resources -- the system must:

* Construct the full range description in etcd in memory and serialize it as protobuf in the client
  * Some clusters have reported over 500MB being stored in a single object type
  * This data is read from the underlying datastore and converted to a protobuf response
  * Large reads to etcd can block writes to the same range (https://github.com/coreos/etcd/issues/7719)
* The data from etcd has to be transferred to the apiserver in one large chunk
* The `kube-apiserver` also has to deserialize that response into a single object, and then re-serialize it back to the client
  * Much of the decoded etcd memory is copied into the struct used to serialize to the client
* An API client like `kubectl get` will then decode the response from JSON or protobuf
  * An API client with a slow connection may not be able to receive the entire response body within the default 60s timeout
    * This may cause other failures downstream of that API client with their own timeouts
  * The recently introduced client compression feature can assist
  * The large response will also be loaded entirely into memory

The standard solution for reducing the impact of large reads is to allow them to be broken into smaller reads via a technique commonly referred to as paging or chunking. By efficiently splitting large list ranges from etcd to clients into many smaller list ranges, we can reduce the peak memory allocation on etcd and the apiserver, without losing the consistent read invariant our clients depend on.

This proposal does not cover general purpose ranging or paging for arbitrary clients, such as allowing web user interfaces to offer paged output, but does define some parameters for future extension. To that end, this proposal uses the phrase "chunking" to describe retrieving a consistent snapshot range read from the API server in distinct pieces.

Our primary consistent store etcd3 offers support for efficient chunking with minimal overhead, and mechanisms exist for other potential future stores such as SQL databases or Consul to also implement a simple form of consistent chunking.

Relevant issues:

* https://github.com/kubernetes/kubernetes/issues/2349

## Terminology

**Consistent list** - A snapshot of all resources at a particular moment in time that has a single `resourceVersion` that clients can begin watching from to receive updates. All Kubernetes controllers depend on this semantic. Allows a controller to refresh its internal state, and then receive a stream of changes from the initial state.

**API paging** - API parameters designed to allow a human to view results in a series of "pages".

**API chunking** - API parameters designed to allow a client to break one large request into multiple smaller requests without changing the semantics of the original request.


## Proposed change:

Expose a simple chunking mechanism to allow large API responses to be broken into consistent partial responses. Clients would indicate a tolerance for chunking (opt-in) by specifying a desired maximum number of results to return in a `LIST` call. The server would return up to that amount of objects, and if more exist it would return a `continue` parameter that the client could pass to receive the next set of results.  The server would be allowed to ignore the limit if it does not implement limiting (backward compatible), but it is not allowed to support limiting without supporting a way to continue the query past the limit (may not implement `limit` without `continue`).

```
GET /api/v1/pods?limit=500
{
  "metadata": {"continue": "ABC...", "resourceVersion": "147"},
  "items": [
     // no more than 500 items
   ]
}
GET /api/v1/pods?limit=500&continue=ABC...
{
  "metadata": {"continue": "DEF...", "resourceVersion": "147"},
  "items": [
     // no more than 500 items
   ]
}
GET /api/v1/pods?limit=500&continue=DEF...
{
  "metadata": {"resourceVersion": "147"},
  "items": [
     // no more than 500 items
   ]
}
```

The token returned by the server for `continue` would be an opaque serialized string that would contain a simple serialization of a version identifier (to allow future extension), and any additional data needed by the server storage to identify where to start the next range. 

The continue token is not required to encode other filtering parameters present on the initial request, and clients may alter their filter parameters on subsequent chunk reads. However, the server implementation **may** reject such changes with a `400 Bad Request` error, and clients should consider this behavior undefined and left to future clarification. Chunking is intended to return consistent lists, and clients **should not** alter their filter parameters on subsequent chunk reads.

If the resource version parameter specified on the request is inconsistent with the `continue` token, the server **must** reject the request with a `400 Bad Request` error.

The schema of the continue token is chosen by the storage layer and is not guaranteed to remain consistent for clients - clients **must** consider the continue token as opaque. Server implementations **should** ensure that continue tokens can persist across server restarts and across upgrades.

Servers **may** return fewer results than `limit` if server side filtering returns no results such as when a `label` or `field` selector is used. If the entire result set is filtered, the server **may** return zero results with a valid `continue` token. A client **must** use the presence of a `continue` token in the response to determine whether more results are available, regardless of the number of results returned. A server that supports limits **must not** return more results than `limit` if a `continue` token is also returned. If the server does not return a `continue` token, the server **must** return all remaining results. The server **may** return zero results with no `continue` token on the last call.

The server **may** limit the amount of time a continue token is valid for. Clients **should** assume continue tokens last only a few minutes.

The server **must** support `continue` tokens that are valid across multiple API servers. The server **must** support a mechanism for rolling restart such that continue tokens are valid after one or all API servers have been restarted.


### Proposed Implementations

etcd3 is the primary Kubernetes store and has been designed to support consistent range reads in chunks for this use case. The etcd3 store is an ordered map of keys to values, and Kubernetes places all keys within a resource type under a common prefix, with namespaces being a further prefix of those keys. A read of all keys within a resource type is an in-order scan of the etcd3 map, and therefore we can retrieve in chunks by defining a start key for the next chunk that skips the last key read.

etcd2 will not be supported as it has no option to perform a consistent read and is on track to be deprecated in Kubernetes.  Other databases that might back Kubernetes could either choose to not implement limiting, or leverage their own transactional characteristics to return a consistent list. In the near term our primary store remains etcd3 which can provide this capability at low complexity.

Implementations that cannot offer consistent ranging (returning a set of results that are logically equivalent to receiving all results in one response) must not allow continuation, because consistent listing is a requirement of the Kubernetes API list and watch pattern.

#### etcd3

For etcd3 the continue token would contain a resource version (the snapshot that we are reading that is consistent across the entire LIST) and the start key for the next set of results. Upon receiving a valid continue token the apiserver would instruct etcd3 to retrieve the set of results at a given resource version, beginning at the provided start key, limited by the maximum number of requests provided by the continue token (or optionally, by a different limit specified by the client). If more results remain after reading up to the limit, the storage should calculate a continue token that would begin at the next possible key, and the continue token set on the returned list.

The storage layer in the apiserver must apply consistency checking to the provided continue token to ensure that malicious users cannot trick the server into serving results outside of its range. The storage layer must perform defensive checking on the provided value, check for path traversal attacks, and have stable versioning for the continue token.

#### Possible SQL database implementation

A SQL database backing a Kubernetes server would need to implement a consistent snapshot read of an entire resource type, plus support changefeed style updates in order to implement the WATCH primitive. A likely implementation in SQL would be a table that stores multiple versions of each object, ordered by key and version, and filters out all historical versions of an object. A consistent paged list over such a table might be similar to:

    SELECT * FROM resource_type WHERE resourceVersion < ? AND deleted = false AND namespace > ? AND name > ? LIMIT ? ORDER BY namespace, name ASC

where `namespace` and `name` are part of the continuation token and an index exists over `(namespace, name, resourceVersion, deleted)` that makes the range query performant. The highest returned resource version row for each `(namespace, name)` tuple would be returned.


### Security implications of returning last or next key in the continue token

If the continue token encodes the next key in the range, that key may expose info that is considered security sensitive, whether simply the name or namespace of resources not under the current tenant's control, or more seriously the name of a resource which is also a shared secret (for example, an access token stored as a kubernetes resource). There are a number of approaches to mitigating this impact:

1. Disable chunking on specific resources
2. Disable chunking when the user does not have permission to view all resources within a range
3. Encrypt the next key or the continue token using a shared secret across all API servers
4. When chunking, continue reading until the next visible start key is located after filtering, so that start keys are always keys the user has access to.

In the short term we have no supported subset filtering (i.e. a user who can LIST can also LIST ?fields= and vice versa), so 1 is sufficient to address the sensitive key name issue. Because clients are required to proceed as if limiting is not possible, the server is always free to ignore a chunked request for other reasons. In the future, 4 may be the best option because we assume that most users starting a consistent read intend to finish it, unlike more general user interface paging where only a small fraction of requests continue to the next page.


### Handling expired resource versions

If the required data to perform a consistent list is no longer available in the storage backend (by default, old versions of objects in etcd3 are removed after 5 minutes), the server **must** return a `410 Gone ResourceExpired` status response (the same as for watch), which means clients must start from the beginning.

```
# resourceVersion is expired
GET /api/v1/pods?limit=500&continue=DEF...
{
  "kind": "Status",
  "code": 410,
  "reason": "ResourceExpired"
}
```

Some clients may wish to follow a failed paged list with a full list attempt.

The 5 minute default compaction interval for etcd3 bounds how long a list can run.  Since clients may wish to perform processing over very large sets, increasing that timeout may make sense for large clusters. It should be possible to alter the interval at which compaction runs to accommodate larger clusters.


#### Types of clients and impact

Some clients such as controllers, receiving a 410 error, may instead wish to perform a full LIST without chunking.

* Controllers with full caches
  * Any controller with a full in-memory cache of one or more resources almost certainly depends on having a consistent view of resources, and so will either need to perform a full list or a paged list, without dropping results
* `kubectl get`
  * Most administrators would probably prefer to see a very large set with some inconsistency rather than no results (due to a timeout under load).  They would likely be ok with handling `410 ResourceExpired` as "continue from the last key I processed"
* Migration style commands
  * Assuming a migration command has to run on the full data set (to upgrade a resource from json to protobuf, or to check a large set of resources for errors) and is performing some expensive calculation on each, very large sets may not complete over the server expiration window.

For clients that do not care about consistency, the server **may** return a `continue` value on the `ResourceExpired` error that allows the client to restart from the same prefix key, but using the latest resource version.  This would allow clients that do not require a fully consistent LIST to opt in to partially consistent LISTs but still be able to scan the entire working set. It is likely this could be a sub field (opaque data) of the `Status` response under `statusDetails`.


### Rate limiting

Since the goal is to reduce spikiness of load, the standard API rate limiter might prefer to rate limit page requests differently from global lists, allowing full LISTs only slowly while smaller pages can proceed more quickly.


### Chunk by default?

On a very large data set, chunking trades total memory allocated in etcd, the apiserver, and the client for higher overhead per request (request/response processing, authentication, authorization).  Picking a sufficiently high chunk value like 500 or 1000 would not impact smaller clusters, but would reduce the peak memory load of a very large cluster (10k resources and up).  In testing, no significant overhead was shown in etcd3 for a paged historical query which is expected since the etcd3 store is an MVCC store and must always filter some values to serve a list.

For clients that must perform sequential processing of lists (kubectl get, migration commands) this change dramatically improves initial latency - clients got their first chunk of data in milliseconds, rather than seconds for the full set. It also improves user experience for web consoles that may be accessed by administrators with access to large parts of the system.

It is recommended that most clients attempt to page by default at a large page size (500 or 1000) and gracefully degrade to not chunking.


### Other solutions

Compression from the apiserver and between the apiserver and etcd can reduce total network bandwidth, but cannot reduce the peak CPU and memory used inside the client, apiserver, or etcd processes. 

Various optimizations exist that can and should be applied to minimizing the amount of data that is transferred from etcd to the client or number of allocations made in each location, but do not how response size scales with number of entries.


## Plan

The initial chunking implementation would focus on consistent listing on server and client as well as measuring the impact of chunking on total system load, since chunking will slightly increase the cost to view large data sets because of the additional per page processing. The initial implementation should make the fewest assumptions possible in constraining future backend storage.

For the initial alpha release, chunking would be behind a feature flag and attempts to provide the `continue` or `limit` flags should be ignored. While disabled, a `continue` token should never be returned by the server as part of a list.

Future work might offer more options for clients to page in an inconsistent fashion, or allow clients to directly specify the parts of the namespace / name keyspace they wish to range over (paging).
