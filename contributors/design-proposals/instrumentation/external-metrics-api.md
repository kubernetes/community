# **External Metrics API**

# Overview

[HPA v2 API extension proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/autoscaling/hpa-external-metrics.md) introduces new External metric type for autoscaling based on metrics coming from outside of Kubernetes cluster. This document proposes a new External Metrics API that will be used by HPA controller to get those metrics.

This API performs a similar role to and is based on existing [Custom Metrics API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md). Unless explicitly specified otherwise all sections related to semantics, implementation and design decisions in [Custom Metrics API design](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md) apply to External Metrics API as well. It is generally expected that a Custom Metrics Adapter will provide both Custom Metrics API and External Metrics API, however, this is not a requirement and both APIs can be implemented and used separately.


# API

The API will consist of a single path:


```
/apis/external.metrics.k8s.io/v1beta1/namespaces/<namespace_name>/<metric_name>?labelSelector=<selector>
```

Similar to endpoints in Custom Metrics API it would only support GET requests.

The query would return the `ExternalMetricValueList` type described below:

```go
// a list of values for a given metric for some set labels
type ExternalMetricValueList struct {
       metav1.TypeMeta `json:",inline"`
       metav1.ListMeta `json:"metadata,omitempty"`

       // value of the metric matching a given set of labels
       Items []ExternalMetricValue `json:"items"`
}

// a metric value for external metric
type ExternalMetricValue struct {
    metav1.TypeMeta`json:",inline"`

    // the name of the metric
    MetricName string `json:"metricName"`

    // label set identifying the value within metric
    MetricLabels map[string]string `json:"metricLabels"`

    // indicates the time at which the metrics were produced
    Timestamp unversioned.Time `json:"timestamp"`

    // indicates the window ([Timestamp-Window, Timestamp]) from
    // which these metrics were calculated, when returning rate
    // metrics calculated from cumulative metrics (or zero for
    // non-calculated instantaneous metrics).
    WindowSeconds *int64 `json:"window,omitempty"`

    // the value of the metric
    Value resource.Quantity
}
```

# Semantics

## Namespaces

Kubernetes namespaces don't have a natural 1-1 mapping to metrics coming from outside of Kubernetes. It is up to adapter implementing the API to decide which metric is available in which namespace. In particular a single metric may be available through many different namespaces.

## Metric Values

A request for a given metric may return multiple values if MetricSelector matches multiple time series. Each value should include a complete set of labels, which is sufficient to uniquely identify a timeseries.

A single value should always be returned if MetricSelector specifies a single value for every label defined for a given metric.

## Metric names

Custom Metrics API [doesn't allow using certain characters in metric names](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md#metric-names). The reason for that is a technical limitation in GO libraries. This list of forbidden characters includes slash (`/`). This is problematic as many systems use slashes in their metric naming convention.

Rather than expect metric adapters to come up with their custom ways of handling that this document proposes introducing `\|` as a custom escape sequence for slash. HPA controller will automatically replace any slashes in MetricName field for External metric with this escape sequence.

Otherwise the allowed metric names are the same as in Custom Metrics API.

## Access Control

Access can be controlled with per-metric granularity, same as in Custom Metrics API. The API has been designed to allow adapters to implement more granular access control if required. Possible future extension of API supporting label level access control is described in [ExternalMetricsPolicy](#externalmetricspolicy) section.

# Future considerations

## ExternalMetricsPolicy

If a more granular access control turns out to be a common requirement an ExternalMetricPolicy object could be added to API. This object could be defined at cluster level, per namespace or per user and would consist of a list of rules. Each rule would consist of a mandatory regexp and either a label selector or a 'deny' statement. For each metric the rules would be applied top to bottom, with the first matching rule being used. A query that hit a deny rule or specified a selector that is not a subset of selector specified by policy would be rejected with 403 error.

Additionally an admission controller could be used to check the policy when creating HPA object.
