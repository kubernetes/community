# Admission control plugin: EventRateLimit

## Background

This document proposes a system for using admission control to enforce a limit
on the number of event requests that the API Server will accept in a given time
slice. In a large multi-tenant cluster, there may be a small percentage of
tenants that are always in some type of error state, in which each tenant is
producing a steady stream of error event requests. Each individual tenant may not
be producing an absurd amount of event requests on its own, but taken collectively
the errors from this small percentage of tenants can have a significant impact on
the performance of the cluster overall. 

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
// Type of limit (e.g., per-namespace)
type LimitType string

const (
	// ServerLimitType limits are maintained against all events received by the server
	ServerLimitType LimitType = "server"
	// NamespaceLimitType limits are maintained against events from each namespace
	NamespaceLimitType LimitType = "namespace"
	// UserLimitType limits are maintained against events from each user
	UserLimitType LimitType = "user"
	// SourceObjectLimitType limits are maintained against events from each source+object
	SourceObjectLimitType LimitType = "source+object"
)

// Configuration provides configuration for the EventRateLimit admission controller.
type Configuration struct {
	metav1.TypeMeta `json:",inline"`

	// Limits to place on events received.
	// Limits can be placed on events received server-wide, per namespace,
	// per user, and per source+object.
	// At least one limit is required.
	Limits []Limit `json:"limits"`
}

type Limit struct {
	// Type of limit
	Type LimitType `json:"type"`

	// Maximum QPS of events for this limit
	QPS float32 `json:"qps"`

	// Maximum burst for throttle of events for this limit
	Burst int64 ` json:"burst"`

	// Size of the LRU cache for this limit. If a bucket is evicted from the cache,
	// then the stats for that bucket are reset. If more events are later received
	// for that bucket, then that bucket will re-enter the cache with a clean slate,
	// giving that bucket a full Burst number of tokens to use.
	//
	// The default cache size is 4096.
	//
	// If LimitType is ServerLimitType, then CacheSize is ignored.
	// +optional
	CacheSize int64 `json:"cacheSize,omitempty"`
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
each namespace. This allowance works in conjunction with the server-wide
allowance. An accepted event request will count against both the server-side
allowance and the per-namespace allowance. The API Server keeps an allowance
for at most 50 namespaces. The API Server will evict the allowance for the
least-recently used nameapce if more than event requests from more than 50
namespaces are received. If an event request for namespace N is received after
the API Server has evicted the allowance for namespace N, then a new, full
allowance will be created for namespace N.

In this example, the API Server will not maintain any allowances for the
user nor source+object of the event in an event request because both the
user and the source+object details have been omitted from the configuration.
The allowance and eviction for per-user and per-source+object rate limiting
works identically to the per-namespace rate limiting, with the exception that
the former consider the user or source+object of the event and the latter
considers the namespace of the event request.

## Client Behavior

Currently, the Client event recorder treats a 429 response as http transport
type of error, that warrants retrying the event request. Instead, the event
recorder should abandon the event. Additionally, the event recorder should
abandon all future events for the period of time specified in the
retryAfterSeconds field of the 429 response.
