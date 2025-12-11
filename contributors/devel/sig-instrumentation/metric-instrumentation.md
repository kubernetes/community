## Instrumenting Kubernetes with Metrics

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

### Adding a New Metric

The following describes the basic steps required to add a new metric (in Go).

1. Import "k8s.io/component-base/metrics" for metrics and "k8s.io/component-base/metrics/legacyregistry" to register your declared metrics.

2. Create a top-level var to define the metric. For this, you have to:

   1. Pick the type of metric. Use a Gauge for things you want to set to a
particular value, a Counter for things you want to increment, or a Histogram or
Summary for histograms/distributions of values (typically for latency).
Histograms are better if you're going to aggregate the values across jobs, while
summaries are better if you just want the job to give you a useful summary of
the values.
   2. Give the metric a name and description (follow the [Prometheus best practices](https://prometheus.io/docs/practices/naming/) for this step).   
   3. Pick whether you want to distinguish different categories of things using
labels on the metric. If so, add "Vec" to the name of the type of metric you
want and add a slice of the label names to the definition.

   [Example](https://github.com/kubernetes/kubernetes/blob/v1.21.1-rc.0/staging/src/k8s.io/apiserver/pkg/endpoints/metrics/metrics.go#L75-L82)
   ```go
	requestCounter = compbasemetrics.NewCounterVec(
		&compbasemetrics.CounterOpts{
			Name:           "apiserver_request_total",
			Help:           "Counter of apiserver requests broken out for each verb, dry run value, group, version, resource, scope, component, and HTTP response code.",
			StabilityLevel: compbasemetrics.STABLE,
		},
		[]string{"verb", "dry_run", "group", "version", "resource", "subresource", "scope", "component", "code"},
	)
   ```

3. Register the metric so that prometheus will know to export it. This can be done in manually or through an init function.

   [Example](https://github.com/kubernetes/kubernetes/blob/v1.21.1-rc.0/staging/src/k8s.io/apiserver/pkg/endpoints/metrics/metrics.go#L280)
   ```go
	legacyregistry.MustRegister(metric)
   ```

4. Use the metric by calling the appropriate method for your metric type (Set,
Inc/Add, or Observe, respectively for Gauge, Counter, or Histogram/Summary),
first calling WithLabelValues if your metric has any labels

   [Example](https://github.com/kubernetes/kubernetes/blob/cd3299307d44665564e1a5c77d0daa0286603ff5/pkg/apiserver/apiserver.go#L87)
   ```go
  	requestCounter.WithLabelValues(*verb, *resource, client, strconv.Itoa(*httpCode)).Inc()
   ```

1. Add tests for the metric that validate its behavior under known conditions.

   [Example test](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/endpoints/metrics/metrics_test.go)
   ```go
   func TestRequestCounter(t *testing.T) {
       // Reset metrics (global registry may have state from other tests)
       metrics.Register()
       defer legacyregistry.Reset()
       
       // Call the underlying code that instruments the metric
       recordRequest("GET", "v1", "pods", "", "default", "", "", "200")
       
       // Use testutil.GatherAndCompare to verify expected values
       expected := `
           # HELP apiserver_request_total Counter of apiserver requests broken out for each verb, dry run value, group, version, resource, scope, component, and HTTP response code.
           # TYPE apiserver_request_total counter
           apiserver_request_total{code="200",component="",dry_run="",group="",resource="pods",scope="default",subresource="",verb="GET",version="v1"} 1
       `
       if err := testutil.GatherAndCompare(legacyregistry.DefaultGatherer, strings.NewReader(expected), "apiserver_request_total"); err != nil {
           t.Fatal(err)
       }
   }
   
   // Alternative: Create a fresh isolated registry only when you need complete isolation
   // or can't use the global registry (e.g., feature-gate dependent metrics)
   func TestIsolatedMetric(t *testing.T) {
       testRegistry := metrics.NewKubeRegistry()
       
       myCounter := metrics.NewCounter(&metrics.CounterOpts{
           Name:           "my_test_counter",
           Help:           "test counter help",
           StabilityLevel: metrics.ALPHA,
       })
       
       testRegistry.MustRegister(myCounter)
       // ... rest of test
   }
   ```

### Graduating an Existing Metric

When graduating a metric from Alpha to Beta or from Beta to Stable, the following requirements must be met. For more information on stability levels and their guarantees, see [Metrics Stability](#metrics-stability).

#### Graduating to Beta

1. **Use-case validation**: Ensure the metric has clear, well-defined use-cases that justify its continued existence. The graduation process is an ideal time to identify metrics that may no longer be relevant or useful and should be considered for cleanup rather than promotion. For eg, see this [issue](https://github.com/kubernetes/kubernetes/pull/136196/changes/BASE..f2ebddae6078062a335481623393f509dc9f53f4#r2747940804) where a metric graduation was questioned due to unclear use-cases.

2. **Naming**: Wherever possible, ensure that metrics graduating to Beta follow [Prometheus metric naming best practices](https://prometheus.io/docs/practices/naming/).

3. **Cardinality**: To the best of ability, only graduate metrics that have bounded cardinality. Pay special attention to the cardinality aspect of a metric to ensure it doesn't create performance issues in monitoring systems.
    
4. **Testing requirement**: The metric must have a corresponding test that validates:
   - The metric exhibits the correct behavior under known conditions (e.g., increments, observations, or value changes as expected for the code path being tested).
   - Avoid tests that only check if a metric is registered or emitted, or that simply call `.Set()`/`.Inc()`/`.Observe()` and expect the value to change since these do not provide meaningful coverage. Focus on testing the metric's behavior as part of the component's logic.
   - For a concrete example of proper metric testing, see the [test example](#adding-a-new-metric) in the "Adding a New Metric" section.

5. **Documentation**: Ensure the metric has a clear and accurate help text description and the metric must be included in the [stable metrics list](https://github.com/kubernetes/kubernetes/blob/master/test/instrumentation/testdata/stable-metrics-list.yaml)
   - See the [instrumentation test README](https://github.com/kubernetes/kubernetes/tree/master/test/instrumentation/README.md) for steps on how to generate this file correctly

6. **API Review**: Graduating a metric to Beta requires an API review by SIG Instrumentation, as it represents a contractual API agreement. See the [API Review](/contributors/devel/sig-instrumentation/metric-stability.md#api-review) section in the metrics stability documentation.

#### Graduating to Stable

The metric must meet all of the requirements for Beta graduation.

1. **Stability validation**: The metric should have been at Beta stability for at least two releases to ensure it has been sufficiently validated in production environments.

2. **Documentation**: The metric must be included in the [stable metrics list](https://github.com/kubernetes/kubernetes/blob/master/test/instrumentation/testdata/stable-metrics-list.yaml)
   - See the [instrumentation test README](https://github.com/kubernetes/kubernetes/tree/master/test/instrumentation/README.md) for steps on how to generate this file correctly

3. **API Review**: Marking a metric as stable is a commitment by the owning SIG to maintain stability guarantees. The owning SIG leads must review and approve the graduation first. Additionally, approval from SIG Instrumentation is required as it represents a contractual API agreement. See the [API Review](/contributors/devel/sig-instrumentation/metric-stability.md#api-review) section in the metrics stability documentation.

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

## Metrics Stability

Please see our documentation on Kubernetes [metrics stability](/contributors/devel/sig-instrumentation/metric-stability.md).

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

### Exception for object state metrics

One exception to the component prefix rule is for metrics derived from
the state of Kubernetes objects.  From the users' perspective, controllers are an
implementation detail of object reconciliation.  The collection of controllers
which comprise a working Kubernetes cluster is viewed as a single system which
drives objects towards their specified desired state.  Metrics concerning a
given object should be easily discoverable and comparable even when they are
produced by different controllers.  Metrics describing the state of a built-in
Kubernetes object take the form:

```
kube_<kind>_<metric>
```

Metrics describing the state of a custom resource avoids collisions by adding a
group.  Metrics take the form:

```
kube_[<group>](https://kubernetes.io/docs/reference/using-api/#api-groups)_<kind>_metric
```

The [Kube-State-Metrics](https://github.com/kubernetes/kube-state-metrics) 
project introduced the original kube_* prefixed metrics.  For examples of
kube_* prefixed metrics, refer to the list of 
[Exposed Metrics](https://github.com/kubernetes/kube-state-metrics/tree/master/docs#exposed-metrics)
in the Kube-State-Metrics documentation.

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

In general, “external” labels like pod name, node name (any object name), & namespace do not belong in the
instrumentation itself (the exception being kube-state-metrics). They are to be attached to metrics by the collecting
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
kube_pod_info{pod=...,namespace=...,pod_ip=...,host_ip=..,node=..., ...} 1
```

The metric system can later denormalize those along the identifying labels
“pod” and “namespace” labels.

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

## Deprecating Metrics

The process of metric deprecation is outlined in the official [Kubernetes Deprecation Policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/). When deprecating a metric, one must set the deprecated version for a version which is in the future from which point that metric will be considered deprecated. If there is a replacement metric, please note that in the help text of the deprecated metric as well as in the corresponding release note of the relevant pull request. 
