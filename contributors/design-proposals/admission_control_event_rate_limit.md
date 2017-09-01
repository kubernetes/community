# Admission control plugin: EventRateLimit

## Background

This document proposes a system for using an admission control to enforce a limit
on the number of event requests that the API Server will accept in a given time
slice. In a large cluster with many namespaces managed by disparate administrators,
there may be a small percentage of namespaces that have pods that are always in
some type of error state, for which the kubelets and controllers in the cluster
are producing a steady stream of error event requests. Each individual namespace
may not be causing a large amount of event requests on its own, but taken
collectively the errors from this small percentage of namespaces can have a
significant impact on the performance of the cluster overall. 

## Use cases

1. Ability to protect the API Server from being flooded by event requests.
2. Ability to protect the API Server from being flooded by event requests for
   a particular namespace.
3. Ability to protect the API Server from being flooded by event requests for
   a particular user.
4. Ability to protect the API Server from being flooded by event requests from
   a particular source+object.

## Data Model

### Configuration

```go
// LimitType is the type of the limit (e.g., per-namespace)
type LimitType string

const (
	// ServerLimitType is a type of limit where there is one bucket shared by
	// all of the event queries received by the API Server.
	ServerLimitType LimitType = "server"
	// NamespaceLimitType is a type of limit where there is one bucket used by
	// each namespace
	NamespaceLimitType LimitType = "namespace"
	// UserLimitType is a type of limit where there is one bucket used by each
	// user
	UserLimitType LimitType = "user"
	// SourceAndObjectLimitType is a type of limit where there is one bucket used
	// by each combination of source and involved object of the event.
	SourceAndObjectLimitType LimitType = "sourceAndObject"
)

// Configuration provides configuration for the EventRateLimit admission
// controller.
type Configuration struct {
	metav1.TypeMeta `json:",inline"`

	// limits are the limits to place on event queries received.
	// Limits can be placed on events received server-wide, per namespace,
	// per user, and per source+object.
	// At least one limit is required.
	Limits []Limit `json:"limits"`
}

// Limit is the configuration for a particular limit type
type Limit struct {
	// type is the type of limit to which this configuration applies
	Type LimitType `json:"type"`

	// qps is the number of event queries per second that are allowed for this
	// type of limit. The qps and burst fields are used together to determine if
	// a particular event query is accepted. The qps determines how many queries
	// are accepted once the burst amount of queries has been exhausted.
	QPS int32 `json:"qps"`

	// burst is the burst number of event queries that are allowed for this type
	// of limit. The qps and burst fields are used together to determine if a
	// particular event query is accepted. The burst determines the maximum size
	// of the allowance granted for a particular bucket. For example, if the burst
	// is 10 and the qps is 3, then the admission control will accept 10 queries
	// before blocking any queries. Every second, 3 more queries will be allowed.
	// If some of that allowance is not used, then it will roll over to the next
	// second, until the maximum allowance of 10 is reached.
	Burst int32 `json:"burst"`

	// cacheSize is the size of the LRU cache for this type of limit. If a bucket
	// is evicted from the cache, then the allowance for that bucket is reset. If
	// more queries are later received for an evicted bucket, then that bucket
	// will re-enter the cache with a clean slate, giving that bucket a full
	// allowance of burst queries.
	//
	// The default cache size is 4096.
	//
	// If limitType is 'server', then cacheSize is ignored.
	// +optional
	CacheSize int32 `json:"cacheSize,omitempty"`
}
```

### Validation

Validation of a **Configuration** enforces that the following rules apply:

* There is at least one item in **Limits**.
* Each item in **Limits** has a unique **Type**.

Validation of a **Limit** enforces that the following rules apply:

* **Type** is one of "server", "namespace", "user", and "source+object".
* **QPS** is positive.
* **Burst** is positive.
* **CacheSize** is non-negative.

### Default Value Behavior

If there is no item in **Limits** for a particular limit type, then no limits
will be enforced for that type of limit.

## AdmissionControl plugin: EventRateLimit

The **EventRateLimit** plug-in introspects all incoming event requests and
determines whether the event fits within the rate limits configured.

To enable the plug-in and support for EventRateLimit, the kube-apiserver must
be configured as follows:

```console
$ kube-apiserver --admission-control=EventRateLimit --admission-control-config-file=$ADMISSION_CONTROL_CONFIG_FILE
```

## Example

An example EventRateLimit configuration:

| Type | RequestBurst | RequestRefillRate | CacheSize |
| ---- | ------------ | ----------------- | --------- |
| Server | 1000 | 100 | |
| Namespace | 100 | 10 | 50 |

The API Server starts with an allowance to accept 1000 event requests. Each
event request received counts against that allowance. The API Server refills
the allowance at a rate of 100 per second, up to a maximum allowance of 1000.
If the allowance is exhausted, then the API Server will respond to subsequent
event requests with 429 Too Many Requests, until the API Server adds more to
its allowance.

For example, let us say that at time t the API Server has a full allowance to
accept 1000 event requests. At time t, the API Server receives 1500 event
requests. The first 1000 to be handled are accepted. The last 500 are rejected
with a 429 response. At time t + 1 second, the API Server has refilled its
allowance with 100 tokens. At time t + 1 second, the API Server receives
another 500 event requests. The first 100 to be handled are accepted. The last
400 are rejected.

The API Server also starts with an allowance to accept 100 event requests from
each namespace. This allowance works in parallel with the server-wide
allowance. An accepted event request will count against both the server-side
allowance and the per-namespace allowance. An event request rejected by the
server-side allowance will still count against the per-namespace allowance,
and vice versa. The API Server tracks the allowances for at most 50 namespaces.
The API Server will stop tracking the allowance for the least-recently-used
namespace if event requests from more than 50 namespaces are received. If an
event request for namespace N is received after the API Server has stop
tracking the allowance for namespace N, then a new, full allowance will be
created for namespace N.

In this example, the API Server will track any allowances for neither the user
nor the source+object in an event request because both the user and the
source+object details have been omitted from the configuration. The allowance
mechanisms for per-user and per-source+object rate limiting works identically
to the per-namespace rate limiting, with the exception that the former consider
the user of the event request or source+object of the event and the latter
considers the namespace of the event request.

## Client Behavior

Currently, the Client event recorder treats a 429 response as an http transport
type of error, which warrants retrying the event request. Instead, the event
recorder should abandon the event. Additionally, the event recorder should
abandon all future events for the period of time specified in the
Retry-After header of the 429 response.
