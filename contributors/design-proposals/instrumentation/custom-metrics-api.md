Custom Metrics API
==================

The new [metrics monitoring vision](monitoring_architecture.md) proposes
an API that the Horizontal Pod Autoscaler can use to access arbitrary
metrics.

Similarly to the [master metrics API](resource-metrics-api.md), the new
API should be structured around accessing metrics by referring to
Kubernetes objects (or groups thereof) and a metric name.  For this
reason, the API could be useful for other consumers (most likely
controllers) that want to consume custom metrics (similarly to how the
master metrics API is generally useful to multiple cluster components).

The HPA can refer to metrics describing all pods matching a label
selector, as well as an arbitrary named object.

API Paths
---------

The root API path will look like `/apis/custom-metrics/v1alpha1`.  For
brevity, this will be left off below.

- `/{object-type}/{object-name}/{metric-name...}`:  retrieve the given
  metric for the given non-namespaced object (e.g. Node, PersistentVolume)

- `/{object-type}/*/{metric-name...}`: retrieve the given metric for all
  non-namespaced objects of the given type

- `/{object-type}/*/{metric-name...}?labelSelector=foo`: retrieve the
  given metric for all non-namespaced objects of the given type matching
  the given label selector

- `/namespaces/{namespace-name}/{object-type}/{object-name}/{metric-name...}`:
  retrieve the given metric for the given namespaced object

- `/namespaces/{namespace-name}/{object-type}/*/{metric-name...}`: retrieve the given metric for all
  namespaced objects of the given type

- `/namespaces/{namespace-name}/{object-type}/*/{metric-name...}?labelSelector=foo`: retrieve the given
  metric for all namespaced objects of the given type matching the
  given label selector

- `/namespaces/{namespace-name}/metrics/{metric-name}`: retrieve the given
  metric which describes the given namespace.

For example, to retrieve the custom metric "hits-per-second" for all
ingress objects matching "app=frontend` in the namespaces "webapp", the
request might look like:

```
GET /apis/custom-metrics/v1alpha1/namespaces/webapp/ingress.extensions/*/hits-per-second?labelSelector=app%3Dfrontend`

---

Verb: GET
Namespace: webapp
APIGroup: custom-metrics
APIVersion: v1alpha1
Resource: ingress.extensions
Subresource: hits-per-second
Name: ResourceAll(*)
```

Notice that getting metrics which describe a namespace follows a slightly
different pattern from other resources; Since namespaces cannot feasibly
have unbounded subresource names (due to collision with resource names,
etc), we introduce a pseudo-resource named "metrics", which represents
metrics describing namespaces, where the resource name is the metric name:

```
GET /apis/custom-metrics/v1alpha1/namespaces/webapp/metrics/queue-length

---

Verb: GET
Namespace: webapp
APIGroup: custom-metrics
APIVersion: v1alpha1
Resource: metrics
Name: queue-length
```

NB: the branch-node LIST operations (e.g. `LIST
/apis/custom-metrics/v1alpha1/namespaces/webapp/pods/`) are unsupported in
v1alpha1. They may be defined in a later version of the API.

API Path Design, Discovery, and Authorization
---------------------------------------------

The API paths in this proposal are designed to a) resemble normal
Kubernetes APIs, b) facilitate writing authorization rules, and c)
allow for discovery.

Since the API structure follows the same structure as other Kubernetes
APIs, it allows for fine grained control over access to metrics.  Access
can be controlled on a per-metric basic (each metric is a subresource, so
metrics may be whitelisted by allowing access to a particular
resource-subresource pair), or granted in general for a namespace (by
allowing access to any resource in the `custom-metrics` API group).

Similarly, since metrics are simply subresources, a normal Kubernetes API
discovery document can be published by the adapter's API server, allowing
clients to discover the available metrics.

Note that we introduce the syntax of having a name of ` * ` here since
there is no current syntax for getting the output of a subresource on
multiple objects.

API Objects
-----------

The request URLs listed above will return the `MetricValueList` type described
below (when a name is given that is not ` * `, the API should simply return a
list with a single element):

```go

// a list of values for a given metric for some set of objects
type MetricValueList struct {
    metav1.TypeMeta`json:",inline"`
    metav1.ListMeta`json:"metadata,omitempty"`

    // the value of the metric across the described objects
    Items []MetricValue `json:"items"`
}

// a metric value for some object
type MetricValue struct {
    metav1.TypeMeta`json:",inline"`

    // a reference to the described object
    DescribedObject ObjectReference `json:"describedObject"`

    // the name of the metric
    MetricName string `json:"metricName"`

    // indicates the time at which the metrics were produced
    Timestamp unversioned.Time `json:"timestamp"`

    // indicates the window ([Timestamp-Window, Timestamp]) from
    // which these metrics were calculated, when returning rate
    // metrics calculated from cumulative metrics (or zero for
    // non-calculated instantaneous metrics).
    WindowSeconds *int64 `json:"window,omitempty"`

    // the value of the metric for this
    Value resource.Quantity
}
```

For instance, the example request above would yield the following object:

```json
{
    "kind": "MetricValueList",
    "apiVersion": "custom-metrics/v1alpha1",
    "items": [
        {
            "metricName": "hits-per-second",
            "describedObject": {
                "kind": "Ingress",
                "apiVersion": "extensions",
                "name": "server1",
                "namespace": "webapp"
            },
            "timestamp": SOME_TIMESTAMP_HERE,
            "windowSeconds": "10",
            "value": "10"
        },
        {
            "metricName": "hits-per-second",
            "describedObject": {
                "kind": "Ingress",
                "apiVersion": "extensions",
                "name": "server2",
                "namespace": "webapp"
            },
            "timestamp": ANOTHER_TIMESTAMP_HERE,
            "windowSeconds": "10",
            "value": "15"
        }
    ]
}
```

Semantics
---------

### Object Types ###

In order to properly identify resources, we must use resource names
qualified with group names (since the group for the requests will always
be `custom-metrics`).

The `object-type` parameter should be the string form of
`unversioned.GroupResource`.  Note that we do not include version in this;
we simply wish to uniquely identify all the different types of objects in
Kubernetes.  For example, the pods resource (which exists in the un-named
legacy API group) would be represented simply as `pods`, while the jobs
resource (which exists in the `batch` API group) would be represented as
`jobs.batch`.

In the case of cross-group object renames, the adapter should maintain
a list of "equivalent versions" that the monitoring system uses. This is
monitoring-system dependent (for instance, the monitoring system might
record all HorizontalPodAutoscalers as in `autoscaling`, but should be
aware that HorizontalPodAutoscaler also exist in `extensions`).

Note that for namespace metrics, we use a pseudo-resource called
`metrics`.  Since there is no resource in the legacy API group, this will
not clash with any existing resources.

### Metric Names ###

Metric names must be able to appear as a single subresource.  In particular,
metric names, *as passed to the API*, may not contain the characters '%', '/',
or '?', and may not be named '.' or '..' (but may contain these sequences).
Note, specifically, that URL encoding is not acceptable to escape the forbidden
characters, due to issues in the Go URL handling libraries. Otherwise, metric
names are open-ended.

### Metric Values and Timing ###

There should be only one metric value per object requested.  The returned
metrics should be the most recently available metrics, as with the resource
metrics API.  Implementers *should* attempt to return all metrics with roughly
identical timestamps and windows (when appropriate), but consumers should also
verify that any differences in timestamps are within tolerances for
a particular application (e.g. a dashboard might simply display the older
metric with a note, while the horizontal pod autoscaler controller might choose
to pretend it did not receive that metric value).

### Labeled Metrics (or lack thereof) ###

For metrics systems that support differentiating metrics beyond the
Kubernetes object hierarchy (such as using additional labels), the metrics
systems should have a metric which represents all such series aggregated
together. Additionally, implementors may choose to identify the individual
"sub-metrics" via the metric name, but this is expected to be fairly rare,
since it most likely requires specific knowledge of individual metrics.
For instance, suppose we record filesystem usage by filesystem inside the
container. There should then be a metric `filesystem/usage`, and the
implementors of the API may choose to expose more detailed metrics like
`filesystem/usage/my-first-filesystem`.

### Resource Versions ###

API implementors should set the `resourceVersion` field based on the
scrape time of the metric.  The resource version is expected to increment
when the scrape/collection time of the returned metric changes.  While the
API does not support writes, and does not currently support watches,
populating resource version preserves the normal expected Kubernetes API
semantics.

Relationship to HPA v2
----------------------

The URL paths in this API are designed to correspond to different source
types in the [HPA v2](../autoscaling/hpa-v2.md).  Specifically, the `pods` source type
corresponds to a URL of the form
`/namespaces/$NS/pods/*/$METRIC_NAME?labelSelector=foo`, while the
`object` source type corresponds to a URL of the form
`/namespaces/$NS/$RESOURCE.$GROUP/$OBJECT_NAME/$METRIC_NAME`.

The HPA then takes the results, aggregates them together (in the case of
the former source type), and uses the resulting value to produce a usage
ratio.

The resource source type is taken from the API provided by the
"metrics" API group (the master/resource metrics API).

The HPA will consume the API as a federated API server.

Relationship to Resource Metrics API
------------------------------------

The metrics presented by this API may be a superset of those present in the
resource metrics API, but this is not guaranteed.  Clients that need the
information in the resource metrics API should use that to retrieve those
metrics, and supplement those metrics with this API.

Mechanical Concerns
-------------------

This API is intended to be implemented by monitoring pipelines (e.g.
inside Heapster, or as an adapter on top of a solution like Prometheus).
It shares many mechanical requirements with normal Kubernetes APIs, such
as the need to support encoding different versions of objects in both JSON
and protobuf, as well as acting as a discoverable API server.  For these
reasons, it is expected that implemenators will make use of the Kubernetes
genericapiserver code.  If implementors choose not to use this, they must
still follow all of the Kubernetes API server conventions in order to work
properly with consumers of the API.

Specifically, they must support the semantics of the GET verb in
Kubernetes, including outputting in different API versions and formats as
requested by the client.  They must support integrating with API discovery
(including publishing a discovery document, etc).

Location
--------

The types and clients for this API will live in a separate repository
under the Kubernetes organization (e.g. `kubernetes/metrics`).  This
repository will most likely also house other metrics-related APIs for
Kubernetes (e.g. historical metrics API definitions, the resource metrics
API definitions, etc).

Note that there will not be a canonical implementation of the custom
metrics API under Kubernetes, just the types and clients.  Implementations
will be left up to the monitoring pipelines.

Alternative Considerations
--------------------------

### Quantity vs Float ###

In the past, custom metrics were represented as floats.  In general,
however, Kubernetes APIs are not supposed to use floats. The API proposed
above thus uses `resource.Quantity`.  This adds a bit of encoding
overhead, but makes the API line up nicely with other Kubernetes APIs.

### Labeled Metrics ###

Many metric systems support labeled metrics, allowing for dimensionality
beyond the Kubernetes object hierarchy.  Since the HPA currently doesn't
support specifying metric labels, this is not supported via this API.  We
may wish to explore this in the future.
