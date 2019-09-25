## Instrumenting Kubernetes

The following references and outlines general guidelines for metric instrumentation
in Kubernetes components. Components are instrumented using the
[Prometheus Go client library](https://github.com/prometheus/client_golang). For non-Go
components. [Libraries in other languages](https://prometheus.io/docs/instrumenting/clientlibs/)
are available.

The metrics are exposed via HTTP in the
[Prometheus metric format](https://prometheus.io/docs/instrumenting/exposition_formats/),
which is open and well-understood by a wide range of third party applications and vendors
outside of the Prometheus eco-system.

The [general instrumentation advice](https://prometheus.io/docs/practices/instrumentation/)
from the Prometheus documentation applies. This document reiterates common pitfalls and some
Kubernetes specific considerations.

Prometheus metrics are cheap as they have minimal internal memory state. Set and increment
operations are thread safe and take 10-25 nanoseconds (Go &amp; Java).
Thus, instrumentation can and should cover all operationally relevant aspects of an application,
internal and external.

## Quick Start

The following describes the basic steps required to add a new metric (in Go).

1. Import "github.com/prometheus/client_golang/prometheus".

2. Create a top-level var to define the metric. For this, you have to:

    1. Pick the type of metric. Use a Gauge for things you want to set to a
particular value, a Counter for things you want to increment, or a Histogram or
Summary for histograms/distributions of values (typically for latency).
Histograms are better if you're going to aggregate the values across jobs, while
summaries are better if you just want the job to give you a useful summary of
the values.
    2. Give the metric a name and description.
    3. Pick whether you want to distinguish different categories of things using
labels on the metric. If so, add "Vec" to the name of the type of metric you
want and add a slice of the label names to the definition.

   [Example](https://github.com/kubernetes/kubernetes/blob/cd3299307d44665564e1a5c77d0daa0286603ff5/pkg/apiserver/apiserver.go#L53)
   ```go
    requestCounter = prometheus.NewCounterVec(
      prometheus.CounterOpts{
        Name: "apiserver_request_count",
        Help: "Counter of apiserver requests broken out for each verb, API resource, client, and HTTP response code.",
      },
      []string{"verb", "resource", "client", "code"},
    )
   ```

3. Register the metric so that prometheus will know to export it.

   [Example](https://github.com/kubernetes/kubernetes/blob/cd3299307d44665564e1a5c77d0daa0286603ff5/pkg/apiserver/apiserver.go#L78)
   ```go
    func init() {
      prometheus.MustRegister(requestCounter)
      prometheus.MustRegister(requestLatencies)
      prometheus.MustRegister(requestLatenciesSummary)
    }
   ```

4. Use the metric by calling the appropriate method for your metric type (Set,
Inc/Add, or Observe, respectively for Gauge, Counter, or Histogram/Summary),
first calling WithLabelValues if your metric has any labels

   [Example](https://github.com/kubernetes/kubernetes/blob/cd3299307d44665564e1a5c77d0daa0286603ff5/pkg/apiserver/apiserver.go#L87)
   ```go
  	requestCounter.WithLabelValues(*verb, *resource, client, strconv.Itoa(*httpCode)).Inc()
   ```


## Instrumentation types

Components have metrics capturing events and states that are inherent to their
application logic. Examples are request and error counters, request latency
histograms, or internal garbage collection cycles. Those metrics are instrumented
directly in the application code.

Secondly, there are business logic metrics. Those are not about observed application
behavior but abstract system state, such as desired replicas for a deployment.
They are not directly instrumented but collected from otherwise exposed data.

In Kubernetes they are generally captured in the [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics)
component, which reads them from the API server.
For this types of metric exposition, the
[exporter guidelines](https://prometheus.io/docs/instrumenting/writing_exporters/)
apply additionally.

## Naming 

General [metric and label naming best practices](https://prometheus.io/docs/practices/naming/) apply.
Beyond that, metrics added directly by application or package code should have a unique name. 
This avoids collisions of metrics added via dependencies. They also clearly
distinguish metrics collected with different semantics. This is solved through
prefixes:

```
<component_name>_<metric>
```

For example, suppose the kubelet instrumented its HTTP requests but also uses
an HTTP router providing its own implementation. Both expose metrics on total
http requests. They should be distinguishable as in:

```
kubelet_http_requests_total{path=”/some/path”,status=”200”}
routerpkg_http_requests_total{path=”/some/path”,status=”200”,method=”GET”}
```

As we can see they expose different labels and thus a naming collision would
not have been possible to resolve even if both metrics counted the exact same
requests.

Resource objects that occur in names should inherit the spelling that is used
in kubectl, i.e. daemon sets are `daemonset` rather than `daemon_set`.

## Dimensionality & Cardinality

Metrics can often replace more expensive logging as they are time-aggregated
over a sampling interval. The [multidimensional data model](https://prometheus.io/docs/concepts/data_model/)
enables deep insights and all metrics should use those label dimensions
where appropriate.

A common error that often causes performance issues in the ingesting metric
system is considering dimensions that inhibit or eliminate time aggregation
by being too specific. Typically those are user IDs or error messages.
More generally: one should know a comprehensive list of all possible values
for a label at instrumentation time.

Notable exceptions are exporters like kube-state-metrics, which expose per-pod
or per-deployment metrics, which are theoretically unbound over time as one could
constantly create new ones, with new names. However, they have
a reasonable upper bound for a given size of infrastructure they refer to and
its typical frequency of changes.

In general, “external” labels like pod or node name do not belong in the
instrumentation itself. They are to be attached to metrics by the collecting
system that has the external knowledge ([blog post](https://www.robustperception.io/target-labels-are-for-life-not-just-for-christmas/)).

## Normalization

Metrics should be normalized with respect to their dimensions. They should
expose the minimal set of labels, each of which provides additional information.
Labels that are composed from values of different labels are not desirable.
For example:

```
example_metric{pod=”abc”,container=”proxy”,container_long=”abc/proxy”}
```

It often seems feasible to add additional meta information about an object
to all metrics about that object, e.g.:

```
kube_pod_container_restarts{namespace=...,pod=...,container=...}
```

A common use case is wanting to look at such metrics w.r.t to the node the
pod is scheduled on. So it seems convenient to add a “node” label.

```
kube_pod_container_restarts{namespace=...,pod=...,container=...,node=...}
```

This however only caters to one specific query use case. There are many more
pieces of metadata that could be added, effectively blowing up the instrumentation.
They are also not guaranteed to be stable over time. What if pods at some
point can be live migrated?
Those pieces of information should be normalized into an info-level metric
([blog post](https://www.robustperception.io/exposing-the-software-version-to-prometheus/)),
which is always set to 1. For example:

```
kube_pod_info{pod=...,namespace=...,pod_ip=...,host_ip=..,node=..., ...}
```

The metric system can later denormalize those along the identifying labels
“pod” and “namespace” labels. This leads to...

## Resource Referencing

It is often desirable to correlate different metrics about a common object,
such as a pod. Label dimensions can be used to match up different metrics.
This is most easy if label names and values are following a common pattern.
For metrics exposed by the same application, that often happens naturally.

For a system composed of several independent, and also pluggable components,
it makes sense to set cross-component standards to allow easy querying in
metric systems without extensive post-processing of data.
In Kubernetes, those are the resource objects such as deployments,
pods, or services and the namespace they belong to.

The following should be consistently used:

```
example_metric_ccc{pod=”example-app-5378923”, namespace=”default”}
```

An object is referenced by its unique name in a label named after the resource
itself (i.e. `pod`/`deployment`/... and not `pod_name`/`deployment_name`)
and the namespace it belongs to in the `namespace` label.

Note: namespace/name combinations are only unique at a certain point in time.
For time series this is given by the timestamp associated with any data point.
UUIDs are truly unique but not convenient to use in user-facing time series
queries.
They can still be incorporated using an info level metric as described above for
`kube_pod_info`. A query to a metric system selecting by UUID via a the info level
metric could look as follows:

```
kube_pod_restarts and on(namespace, pod) kube_pod_info{uuid=”ABC”}
```

