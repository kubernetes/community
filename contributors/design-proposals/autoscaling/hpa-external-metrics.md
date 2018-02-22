# **HPA v2 API extension proposal**

# Objective

[Horizontal Pod Autoscaler v2 API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa-v2.md) allows users to autoscale based on custom metrics. However, there are some use-cases that are not well supported by the current API. The goal of this document is to propose the following changes to the API:

*   Allow autoscaling based on metrics coming from outside of Kubernetes. Example use-case is autoscaling based on a hosted cloud service used by a pod.
*   Allow specifying per-pod target for global metrics. This makes more sense for many metrics than a global target (ex. 200 QPS / pod makes sense while a global target for QPS doesn't).

# Overview

A new External metric source will be added. It will identify a specific metric to autoscale on based on metric name and a label selector. The assumed model is that specific time series in monitoring systems are identified with metric name and a set of key-value labels or tags. The details vary in different systems (ref: [Prometheus](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels), [Stackdriver](https://cloud.google.com/monitoring/api/v3/metrics#time_series), [Datadog](https://docs.datadoghq.com/agent/tagging/), [Sysdig](https://www.sysdig.org/wiki/sysdig-user-guide/#user-content-filtering)), however, in general the adapter should be able to use metric name and a set of labels to construct a query sufficient to identify specific time series in underlying system.

External and Object metrics will specify the desired target by setting exactly one of two fields: TargetValue (global target) or TargetAverageValue (per-pod target).


# Multiple metric values

Label selector specified by user can match multiple time series, resulting in multiple values provided to HPA. In such case the sum of all those values will be used for autoscaling. This is meant to allow autoscaling using a metric that is drilled down by some criteria not relevant for autoscaling a particular workload (for example to allow autoscaling based on a total number of HTTP requests, regardless of which HTTP method is used).

If the need arises we can easily add other simple aggregations and allow user to choose one of them.

# Example

This is an example HPA configuration autoscaling based on number of pending messages in a message queue (RabbtMQ) running outside of cluster.

```yaml
kind: HorizontalPodAutoscaler
apiVersion: autoscaling/v2beta2
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: Worker
  minReplicas: 2
  maxReplicas: 10
  metrics:
   - type: External
     external:
       metricName: queue_messages_ready
       metricSelector:
         matchLabels:
           queue: worker_tasks
       targetAverageValue: 30
```

# API

This is the part of autoscaling/v2beta2 API that includes changes from v2beta1.

Some parts containing obvious changes (MetricSpec and MetricStatus) have been omitted for clarity.

```go
// MetricSourceType indicates the type of metric.
type MetricSourceType string

var (
        // ObjectMetricSourceType is a metric describing a Kubernetes object
        // (for example, hits-per-second on an Ingress object).
        ObjectMetricSourceType MetricSourceType = "Object"
        // PodsMetricSourceType is a metric describing each pod in the current scale
        // target (for example, transactions-processed-per-second).  The values
        // will be averaged together before being compared to the target value.
        PodsMetricSourceType MetricSourceType = "Pods"
        // ResourceMetricSourceType is a resource metric known to Kubernetes, as
        // specified in requests and limits, describing each pod in the current
        // scale target (e.g. CPU or memory).  Such metrics are built in to
        // Kubernetes, and have special scaling options on top of those available
        // to normal per-pod metrics (the "pods" source).
        ResourceMetricSourceType MetricSourceType = "Resource"
        // ExternalMetricSourceType is a global metric that is not associated
        // with any Kubernetes object. It allows autoscaling based on information
        // coming from components running outside of cluster
        // (for example length of queue in cloud messaging service, or
        // QPS from loadbalancer running outside of cluster).
        ExternalMetricSourceType MetricSourceType = "External"
)

// ObjectMetricSource indicates how to scale on a metric describing a
// Kubernetes object (for example, hits-per-second on an Ingress object).
type ObjectMetricSource struct {
        // target is the described Kubernetes object.
        Target CrossVersionObjectReference `json:"target" protobuf:"bytes,1,name=target"`

        // metricName is the name of the metric in question.
        MetricName string `json:"metricName" protobuf:"bytes,2,name=metricName"`
        // TargetValue is the target value of the metric (as a quantity).
        // Mutually exclusive with TargetAverageValue.
        TargetValue *resource.Quantity `json:"targetValue,omitempty" protobuf:"bytes,3,opt,name=targetValue"`
        // TargetAverageValue is the target per-pod value of global metric.
        // Mutually exclusive with TargetValue.
        TargetAverageValue *resource.Quantity `json:"targetAverageValue,omitempty" protobuf="bytes,4,opt,name=targetAverageValue"`
}

// ExternalMetricSource indicates how to scale on a metric not associated with
// any Kubernetes object (for example length of queue in cloud
// messaging service, or QPS from loadbalancer running outside of cluster).
type ExternalMetricSource struct {
        // MetricName is the name of a metric used for autoscaling in
        // metric system.
        MetricName string `json:"metricName" protobuf:"bytes,1,name=metricName"`

        // MetricSelector is used to identify a specific time series
        // within a given metric.
        MetricSelector metav1.LabelSelector `json:"metricSelector" protobuf:"bytes,2,name=metricSelector"`

        // TargetValue is the target value of the metric (as a quantity).
        // Mutually exclusive with TargetAverageValue.
        TargetValue *resource.Quantity `json:"targetValue,omitempty" protobuf:"bytes,3,opt,name=targetValue"`
        // TargetAverageValue is the target per-pod value (as a quantity) of global metric.
        // Mutually exclusive with TargetValue.
        TargetAverageValue *resource.Quantity `json:"targetAverageValue,omitempty" protobuf="bytes,4,opt,name=targetAverageValue"`
}

// ExternalMetricStatus indicates the current value of a global metric
// not associated with any Kubernetes object.
type ExternalMetricStatus struct {
        // MetricName is the name of a metric used for autoscaling in
        // metric system.
        MetricName string `json:"metricName" protobuf:"bytes,1,name=metricName"`

        // MetricSelector is used to identify a specific time series
        // within a given metric.
        MetricSelector metav1.LabelSelector `json:"metricSelector" protobuf:"bytes,2,name=metricSelector"`

        // CurrentValue is the current value of the metric (as a quantity)
        CurrentValue resource.Quantity `json:"currentValue" protobuf:"bytes,3,name=currentValue"`

        // CurrentAverageValue is the current value of metric averaged over
        // autoscaled pods.
        CurrentAverageValue *resource.Quantity `json:"currentAverageValue,omitempty" protobuf:"bytes,4,opt,name=currentAverageValue"`
}
```

# Implementation

A new [External Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/external-metrics-api.md) will be used to obtain values of External metrics.

As a result of proposed changes TargetValue field becomes optional. Kubernetes convention is to make optional fields pointers. Changing value to pointer makes this proposal non backward compatible, requiring moving to v2beta2 as opposed to extending v2beta1 API.

[Kubernetes deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/) requires that autoscaling/v2beta1 is still supported for 1.10 and 1.11 and roundtrip between different API versions can be made. This will be implemented the same way as current conversion between v2 and v1 - MetricSpec specifying TargetAverageValue will be serialized into json and stored in annotation when converting to v2beta1 representation.

# Future considerations

### Add LabelSelector to Object and Pods metric

Object and Pods metrics rely on implicit assumption that there is just a single time series per metric. This can be pretty limiting when using applications that drill down the metrics by some additional criteria (ex. method in case of metrics related to HTTP requests). We could consider adding LabelSelector to Object and Pods metrics working similarly to how it works for External metrics.

# Alternatives considered

### Identifying external metrics

The main argument for choosing metric name and label selector over other ways of identifying external metric was access control. Any query specified by user in External metric spec will be executed by HPA controller (ie. system account). Per-metric access control can be applied using standard kubernetes mechanisms. It is up to adapter implementing [External Metrics API to ensure access control at labels level](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/external-metrics-api.md#access-control).

An important consideration is that query languages of monitoring systems can be very powerful (ex. [Prometheus docs](https://prometheus.io/docs/prometheus/latest/querying/examples/#using-functions-operators-etc) and [example queries](https://github.com/infinityworks/prometheus-example-queries/blob/master/README.md#promql-examples)), making it extremely difficult (if not impossible) to implement access control in adapter. The idea behind the proposed approach is to limit expressive power available to user to identifying an existing time series in monitoring system.

In order to use a more advanced query (including aggregations of multiple time series) for autoscaling the user will have to re-export the result of such query as a new metric. Such advanced use-cases are likely to be rare and re-exporting is a viable workaround. On the other hand, allowing user to perform arbitrary query in underlying metric system using system account would have been a blocker for enabling External Metrics API for a large number of users.

### Using Object metric source for external metrics
An alternative to adding External metric source would be to reuse existing Object metric source. Both External metrics proposed in this document and Object metrics represent a global custom metric and there is no conceptual difference between them, except for how the metric is identified in underlying monitoring system. Using Object metrics for both use-cases can help keep API simple.

The immediate problem with this approach is that there is no inherent relationship between any Kubernetes object and an arbitrary external metric. It's not clear how Kubernetes object reference used to specify Object metrics could be used to identify such metric. This section discusses different solutions to this problem along with their pros and cons.

#### Attach external metric to a Kubernetes object
One possible solution would be to allow user to explicitly specify a relationship between an external metric and some Kubernetes object (in particular attaching metric to a Namespace seems logical). Custom Metrics Adapter could use this additional information to translate object reference provided by HPA into query to monitoring system. This could be either left to adapter (i.e. adapter specific config) or introduced to custom metrics API. Below table details pros and cons of each approach.

<table>
  <tbody>
    <tr>
      <th>Option</th>
      <th>Pros</th>
      <th>Cons</th>
    </tr>
    <tr>
      <td>Adapter specific config</td>
      <td><ul>
        <li>No need to change HPA API.
        <li>Adapter-specific way of configuring metrics will better match underlying metrics system than any generic solution. It could be both more logical and offer better validation.
      </ul></td>
      <td><ul>
        <li>This just pushes the problem to adapter. Different adapters are likely to solve it differently (or not at all), resulting in widespread incompatibility.
        <li>Hard to use. Instead of just understanding HPA API user would need to know about Custom Metric Adapter and learn it’s configuration syntax and best practices.
        <li>Potential access control problems if shared config is used for attaching multiple metrics in different namespaces.
      </ul></td>
    </tr>
      <td>Add an object containing mapping to Custom Metrics API</td>
      <td><ul>
        <li>No need to change HPA API.
      </ul></td>
      <td><ul>
        <li>This is just moving External metric from HPA into a separate object, for no obvious benefit (in most cases the new object will map 1-1 to HPA anyway).
        <li>Harder to use (need to create an additional object, than reference it from HPA.
      </ul></td>
    </tr>
  </tbody>
</table>

Overall it seems that the same information provided to External metric spec would have to be provided by user anyway. Storing it elsewhere makes the feature more complex to use, for no clear benefit.

#### Implicitly attach external metrics to namespace
After [extending Object metrics with LabelSelector](#add-labelselector-to-object-and-pods-metric) Object metric will contain enough information to identify external metric. Theoretically Custom Metrics Adapter could implicitly attach every available metric to some arbitrary object (ex. `default` namespace or every namespace). However, this is equivalent to just making the object reference optional using the fact that the set of fields in Object metric would be superset of fields in External metric. Technically it would work, but it feels like a hack and it's likely to be confusing to users. Also it makes it easy to rely on access control on referenced object, which could be easily circumvented if every metric is available via every object.

#### Relabel external metrics to add metadata attaching them to a chosen Kubernetes object
After [extending Object metrics with LabelSelector](#add-labelselector-to-object-and-pods-metric) user could just [relabel their metrics](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config) to add metadata attaching them to Kubernetes object of their choice. However, this approach assumes user has sufficient access to relabel metrics. This is not always the case (for example when autoscaling on metrics from a hosted service).

A variant of this approach would be to ask user to create a pod that reads a metric and reexports it. This will work even without any changes in HPA, however, it requires complex setup, wastes resources and may introduce additional latency.
