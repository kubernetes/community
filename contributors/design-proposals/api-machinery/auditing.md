# Auditing

Maciej Szulik (@soltysh)
Dr. Stefan Schimanski (@sttts)
Tim St. Clair (@timstclair)

## Abstract

This proposal aims at extending the auditing log capabilities of the apiserver.

## Motivation and Goals

With [#27087](https://github.com/kubernetes/kubernetes/pull/27087) basic audit logging was added to Kubernetes. It basically implements `access.log` like http handler based logging of all requests in the apiserver API. It does not do deeper inspection of the API calls or of their payloads. Moreover, it has no specific knowledge of the API objects which are modified. Hence, the log output does not answer the question how API objects actually change.

The log output format in [#27087](https://github.com/kubernetes/kubernetes/pull/27087) is fixed. It is text based, unstructured (e.g. non-JSON) data which must be parsed to be usable in any advanced external system used to analyze audit logs.

The log output format does not follow any public standard like e.g. https://www.dmtf.org/standards/cadf.

With this proposal we describe how the auditing functionality can be extended in order:

- to allow multiple output formats, e.g. access.log style or structured JSON output
- to allow deep payload inspection to allow
  - either real differential JSON output (which field of an object have changed from which value to which value)
  - or full object output of the new state (and optionally the old state)
- to be extensible in the future to fully comply with the Cloud Auditing Data Federation standard (https://www.dmtf.org/standards/cadf)
- to allow filtering of the output
  - by kind, e.g. don't log endpoint objects
  - by object path (JSON path), e.g. to ignore all `*.status` changes
  - by user, e.g. to only log end user action, not those of the controller-manager and scheduler
  - by level (request headers, request object, storage object)

while

- not degrading apiserver performance when auditing is enabled.

## Constraints and Assumptions

* it is not the goal to implement all output formats one can imagine. The main goal is to be extensible with a clear golang interface. Implementations of e.g. CADF must be possible, but won't be discussed here.
* dynamic loading of backends for new output formats are out of scope.

## Use Cases

1. As a cluster operator I want to enable audit logging of requests to the apiserver in order **to comply with given business regulations** regarding a subset of the 7 Ws of auditing:

    - **what** happened?
    - **when** did it happen?
    - **who** initiated it?
    - **on what** did it happen (e.g. pod foo/bar)?
    - **where** was it observed (e.g. apiserver hostname)?
    - from **where** was it initiated? (e.g. kubectl IP)
    - to **where** was it going? (e.g. node 1.2.3.4 for kubectl proxy, apiserver when logged at aggregator).

1. Depending on the environment, as a cluster operator I want to **define the amount of audit logging**, balancing computational overhead for the apiserver with the detail and completeness of the log.

1. As a cluster operator I want to **integrate with external systems**, which will have different requirements for the log format, network protocols and communication modes (e.g. pull vs. push).

1. As a cluster operator I must be able to provide a **complete trace of changes to an object** to API objects.

1. As a cluster operator I must be able to create a trace for **all accesses to a secret**.

1. As a cluster operator I must be able to log non-CRUD access like **kubectl exec**, when it started, when it finished and with which initial parameters.

### Out of scope use-cases

1. As a cluster operator I must be able to get a trace of non-REST calls executed against components other than kube-apiserver, kube-aggregator and their counterparts in federation. This includes operations requiring HTTP upgrade requests to support multiplexed bidirectional streams (HTTP/2, SPDY), direct calls to kubelet endpoints, port forwarding, etc.

## Community Work

- Kubernetes basic audit log PR: https://github.com/kubernetes/kubernetes/pull/27087/
- OpenStack's implementation of the CADF standard: https://www.dmtf.org/sites/default/files/standards/documents/DSP2038_1.1.0.pdf
- Cloud Auditing Data Federation standard: https://www.dmtf.org/standards/cadf
- Ceilometer audit blueprint: https://wiki.openstack.org/wiki/Ceilometer/blueprints/support-standard-audit-formats
- Talk from IBM: An Introduction to DMTF Cloud Auditing using
the CADF Event Model and Taxonomies https://wiki.openstack.org/w/images/e/e1/Introduction_to_Cloud_Auditing_using_CADF_Event_Model_and_Taxonomy_2013-10-22.pdf

## Architecture

When implementing audit logging there are basically two options:

1. put a logging proxy in front of the apiserver
2. integrate audit logging into the apiserver itself

Both approaches have advantages and disadvantages:
- **pro proxy**:
  + keeps complexity out of the apiserver
  + reuses existing solutions
- **contra proxy**:
  + has no deeper insight into the Kubernetes api
  + has no knowledge of authn, authz, admission
  + has no access to the storage level for differential output
  + has to terminate SSL and complicates client certificates based auth

In the following, the second approach is described without a proxy.  At which point there are a few possible places, inside the apiserver, where auditing could happen, namely:
1. as one of the REST handlers (as in [#27087](https://github.com/kubernetes/kubernetes/pull/27087)),
2. as an admission controller.

The former approach (currently implemented) was picked over the other one, due to the need to be able to get information about both the user submitting the request and the impersonated user (and group), which is being overridden inside the [impersonation filter](https://git.k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/endpoints/filters/impersonation.go).  Additionally admission controller does not have access to the response and runs after authorization which will prevent logging failed authorization.  All of that resulted in continuing the solution started in [#27087](https://github.com/kubernetes/kubernetes/pull/27087), which implements auditing as one of the REST handlers
after authentication, but before impersonation and authorization.

## Proposed Design

The main concepts are those of

- an audit *event*,
- an audit *level*,
- an audit *filters*,
- an audit *output backend*.

An audit event holds all the data necessary for an *output backend* to produce an audit log entry. The *event* is independent of the *output backend*.

The audit event struct is passed through the apiserver layers as an `*audit.Event` pointer inside the apiserver's `Context` object. It is `nil` when auditing is disabled.

If auditing is enabled, the http handler will attach an `audit.Event` to the context:

```go
func WithAuditEvent(parent Context, e *audit.Event) Context
func AuditEventFrom(ctx Context) (*audit.Event, bool)
```

Depending on the audit level (see [below](#levels)), different layers of the apiserver (e.g. http handler, storage) will fill the `audit.Event` struct. Certain fields might stay empty or `nil` if given level does not require that field. E.g. in the case when only http headers are supposed to be audit logged, no `OldObject` or `NewObject` is to be retrieved on the storage layer.

### Levels

Proposed audit levels are:

- `None` - don't audit the request.
- `Metadata` - reflects the current level of auditing, iow. provides following information about each request: timestamp, source IP, HTTP method, user info (including group as user and as group), namespace, URI and response code.
- `RequestBody` - additionally provides the unstructured request body.
- `ResponseBody` - additionally provides the unstructured response body. Equivalent to `RequestBody` for streaming requests.
- `StorageObject` - provides the object before and after modification.

```go
package audit

// AuditLevel defines the amount of information logged during auditing
type AuditLevel string

// Valid audit levels
const (
    // AuditNone disables auditing
    AuditNone AuditLevel = "None"
    // AuditMetadata provides basic level of auditing, logging data at HTTP level
    AuditMetadata AuditLevel   = "Metadata"
    // AuditRequestBody provides Header level of auditing, and additionally
    // logs unstructured request body
    AuditRequestBody AuditLevel = "RequestBody"
    // AuditResponseBody provides Request level of auditing, and additionally
    // logs unstructured response body
    AuditResponseBody AuditLevel = "ResponseBody"
    // AuditStorageobject provides Response level, and additionally
    // logs object before and after saving in storage
    AuditStorageObject AuditLevel = "StorageObject"
)
```

The audit level is determined by the policy, which maps a combination of requesting user, namespace, verb, API group, and resource. The policy is described in detail [below](#policy).

In an [aggregated](aggregated-api-servers.md) deployment, the `kube-aggregator` is able to fill in
`Metadata` level audit events, but not above. For the higher audit levels, an audit event is
generated _both_ in the `kube-aggregator` and in the end-user apiserver. The events can be
de-duplicated in the audit backend based on the audit ID, which is generated from the `Audit-ID`
header. The event generated by the end-user apiserver may not have the full authentication information.

**Note:** for service creation and deletion there is special REST code in the apiserver which takes care of service/node port (de)allocation and removal of endpoints on service deletion. Hence, these operations are not visible on the API layer and cannot be audit logged therefore. **No other resources** (with the exception of componentstatus which is not of interest here) **implement this kind of custom CRUD operations.**

### Events

The `Event` object contains the following data:

```go
package audit

type Event struct {
    // AuditLevel at which event was generated
    Level AuditLevel

    // below fields are filled at Metadata level and higher:

    // Unique ID of the request being audited, and able to de-dupe audit events.
    // Set from the `Audit-ID` header.
    ID string
    // Time the event reached the apiserver
    Timestamp Timestamp
    // Source IPs, from where the request originates, with intermediate proxy IPs.
    SourceIPs []string
    // HTTP method sent by the client
    HttpMethod string
    // Verb is the kube verb associated with the request for API requests.
    // For non-resource requests, this is identical to HttpMethod.
    Verb string
    // Authentication method used to allow users access the cluster
    AuthMethod string
    // RequestURI is the Request-Line as sent by the client to a server
    RequestURI string
    // User information
    User UserInfo
    // Impersonation information
    Impersonate UserInfo
    // Object reference this request is targeted at
    Object ObjectReference
    // Response status code are returned by the server
    ResponseStatusCode int
    // Error response, if ResponseStatusCode >= 400
    ResponseErrorMessage string

    // below fields are filled at RequestObject level and higher:

    // RequestObject logged before admission (json format)
    RequestObject runtime.Unstructured
    // Response object in json format
    ResponseObject runtime.Unstructured

    // below fields are filled at StorageObject level and higher:

    // Object value before modification (will be empty when creating new object)
    OldObject runtime.Object
    // Object value after modification (will be empty when removing object)
    NewObject runtime.Object
}
```

### Policy

The audit policy determines what audit event is generated for a given request. The policy is configured by the cluster administrator. Here is a sketch of the policy API:

```go
type Policy struct {
    // Rules specify the audit Level a request should be recorded at.
    // A request may match multiple rules, in which case the FIRST matching rule is used.
    // The default audit level is None, but can be overridden by a catch-all rule at the end of the list.
    Rules []PolicyRule

    // Discussed under Filters section.
    Filters []Filter
}

// Based off the RBAC PolicyRule
type PolicyRule struct {
    // Required. The Level that requests matching this rule are recorded at.
    Level Level

    // The users (by authenticated user name) this rule applies to.
    // An empty list implies every user.
    Users []string
    // The user groups this rule applies to. If a user is considered matching
    // if they are a member of any of these groups
    // An empty list implies every user group.
    UserGroups []string

    // The verbs that match this rule.
    // An empty list implies every verb.
    Verbs []string

    // Rules can apply to API resources (such as "pods" or "secrets"),
    // non-resource URL paths (such as "/api"), or neither, but not both.
    // If neither is specified, the rule is treated as a default for all URLs.

    // APIGroups is the name of the APIGroup that contains the resources ("" for core).
    // If multiple API groups are specified, any action requested against one of the
    // enumerated resources in any API group will be allowed.
    // Any empty list implies every group.
    APIGroups []string
    // Namespaces that this rule matches.
    // This field should be left empty if specifying non-namespaced resources.
    // Any empty list implies every namespace.
    Namespaces []string
    // GroupResources is a list of GroupResource types this rule applies to.
    // Any empty list implies every resource type.
    GroupResources []GroupResource
    // ResourceNames is an optional white list of names that the rule applies to.
    // Any empty list implies everything.
    ResourceNames []string

    // NonResourceURLs is a set of partial urls that should be audited.
    // *s are allowed, but only as the full, final step in the path.
    // If an action is not a resource API request, then the URL is split on '/' and
    // is checked against the NonResourceURLs to look for a match.
    NonResourceURLs []string
}
```

As an example, the administrator may decide that by default requests should be audited with the
response, except for get and list requests. On top of that, they wish to completely ignore the noisy
`kube-proxy` endpoint requests. This looks like:

```yaml
rules:
    - level: None
      users: ["system:kube-proxy"]
      apiGroups: [""] # The core API group
      resources: ["endpoints"]
    - level: RequestBody
      verbs: ["get", "list"]
    # The default for non-resource URLs
    - level: Metadata
      nonResourceURLs: ["*"]
    # The default for everything else
    - level: ResponseBody
```

The policy is checked immediately after authentication in the request handling, and determines how
the `audit.Event` is formed.

In an [aggregated](aggregated-api-servers.md) deployment, each apiserver must be independently
configured for audit logging (including the aggregator).

### Filters

In addition to the high-level policy rules, auditing can be controlled at a more fine-grained level
with `Filters`. Unlike the policy, filters are applied _after_ the `audit.Event` is constructed, but
before it's passed to the output backend.

TODO: Define how filters work. They should enable dropping sensitive fields from the
request/response/storage objects.

### Output Backend Interface

```go
package audit

type OutputBackend interface {
  // Thread-safe, blocking.
  Log(e *Event) error
}
```

It is the responsibility of the OutputBackend to manage concurrency, but it is acceptable for the
method to block. Errors will be handled by recording a count in prometheus, and attempting to write
to the standard debug log.

### Apiserver Command Line Flags

Deprecate flags currently used for configuring audit:
* `--audit-log-path` - specifies the file where requests are logged,
* `--audit-log-maxage` - specifies the maximum number of days to retain old log files,
* `--audit-log-maxbackup` - specifies maximum number of old files to retain,
* `--audit-log-maxsize` - specifies maximum size in megabytes of the log file.

Following new flags should be introduced in the apiserver:
* `--audit-output` - which specifies the backend and its configuration, example:
  `--audit-output file:path=/var/log/apiserver-audit.log,rotate=1d,max=1024MB,format=json`
  which will log to `/var/log/apiserver-audit.log`, and additionally defines rotation arguments (analogically to the deprecated ones) and output format.
* `--audit-policy` - which specifies a file with policy configuration, see [policy](#policy) for a sample file contents.

### Audit Security

Several parts of the audit system could be exposed to spoofing or tampering threats. This section
lists the threats, and how we will mitigate them.

**Audit ID.** The audit ID is set from the "front door" server, which could be a federation apiserver,
a kube-aggregator, or end-user apiserver. Since the server can't currently know where in the serving
chain it falls, it is possible for the client to set a `Audit-ID` header that is
non-unique. For this reason, any aggregation that happens based on the audit ID must also sanity
check the known fields (e.g. URL, source IP, time window, etc.). With this additional check, an
attacker could generate a bit more noise in the logs, but no information would be lost.

**Source IP.** Kubernetes requests may go through multiple hops before a response is generated
(e.g. federation apiserver -> kube-aggregator -> end-user apiserver). Each hop must append the
previous sender's IP address to the `X-Forwarded-For` header IP chain. If we simply audited the
original sender's IP, an attacker could send there request with a bogus IP at the front of the
`X-Forwarded-For` chain. To mitigate this, we will log the entire IP chain. This has the additional
benefit of supporting external proxies.

## Sensible (not necessarily sequential) Milestones of Implementation

1. Add `audit.Event` and `audit.OutputBackend` and implement [#27087](https://github.com/kubernetes/kubernetes/pull/27087)'s basic auditing using them, using a single global audit Level, up to `ResponseBody`.
1. Implement the full `audit.Policy` rule specification.
1. Add deep inspection on the storage level to the old and the new object.
1. Add filter support (after finishing the Filters section of this proposal).

## Future evolution

Below are the possible future extensions to the auditing mechanism:
* Define how filters work. They should enable dropping sensitive fields from the request/response/storage objects.
* Allow setting a unique identifier which allows matching audit events across apiserver and federated servers.

