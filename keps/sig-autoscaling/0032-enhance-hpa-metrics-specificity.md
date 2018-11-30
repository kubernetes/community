---
kep-number: 0032
title: Enhance HPA Metrics Specificity
authors:
  - "@directxman12"
owning-sig: sig-autoscaling
participating-sigs:
  - sig-instrumentation
reviewers:
  - "@brancz"
  - "@maciekpytel"
approvers:
  - "@brancz"
  - "@maciekpytel"
  - "@directxman12"
editor: "@directxman12"
creation-date: 2018-04-19
status: implemented
---

# Enhance HPA Metrics Specificity

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Alternatives](#alternatives)

## Summary

The `External` metric source type in the HPA currently supports passing
a metric label selectors, which is passed to the custom metrics API
(custom.metrics.k8s.io) to select a more specific metrics series.  This
allows users to more easily make use of existing metrics structure,
without need to manipulate their metrics labeling and ingestion
externally.

Additionally, it supports the `targetAverageValue` field, which allows
artificially dividing an external metric by the number of replicas in the
target scalable.

This proposal brings both of those fields to the `Object` metric source
type, and further brings the selector field to the `Pods` metric source
type, making both types more flexible and bringing them in line with the
`External` metrics types.

## Motivation

With custom-metrics-based autoscaling, users frequently ask how to select
more specific metrics in their metric storage.  For instance, a user might
have message queue statefulset with several queues producing metrics:

```
queue_length{statefulset="foo",pod="foo-1",queue="some-jobs"}
queue_length{statefulset="foo",pod="foo-1",queue="other-jobs"}
```

Suppose they have a pool of works that they wish to scale for each queue.
In the current-day HPA, it's non-trivial to allow selecting the metric for
a specific queue.  Current suggestions are metric-backend-specific (for
instance, you could create a Prometheus recording rule to relabel or
rename the metric), and often involve making external changes to the
metrics pipeline.

With the addition of the metrics label selector, users could simply select
the queue using the label selector:

```yaml
- type: Object
  object:
    describedObject:
      kind: StatefulSet
      apiVersion: apps/v1
      name: foo
    target:
      type: Value
      value: 2
    metric:
      name: queue_length
      selector: {matchLabels: {queue: some-jobs}}
```

Similarly, in discussions of scaling on queues, being able to divide
a target backlog length by the number of available pods is often useful --
for instance, a backlog length of 3 might be acceptable if there are three
pods processing items, but not if there is only one.

### Goals

- The autoscaling/v2 API is updated with the additional fields
  described below.
- A corresponding change is made to the custom metrics API to support the
  additional label selector.
- The testing adapter is updated to support these changes (for e2e
  purposes).

### Non-Goals

It is outside of the purview of the KEP to ensure that current custom
metrics adapters support the new changes -- this is up to those adapters
maintainers.

## Proposal

The autoscaling/v2 API will be updated the following way:

```go
type ObjectMetricSource struct {
	DescribedObject CrossVersionObjectReference
	Target MetricTarget
	Metric MetricIdentifier
}

type PodsMetricSource struct {
    Target MetricTarget
	Metric MetricIdentifier
}

type ExternalMetricSouce struct {
	Metric MetricIdentifier
	Target MetricTarget
}

type ResourceMetricSource struct {
	Name v1.ResourceName
	Target MetricTarget
}

type MetricIdentifier struct {
	// name is the name of the given metric
	Name string
	// selector is the selector for the given metric
	// +optional
	Selector *metav1.LabelSelector
}

type MetricTarget struct {
	Type MetricTargetType // Utilization, Value, AverageValue
    // value is the raw value of the single metric (valid for object metrics)
	Value *resource.Quantity

    // averageValue is the raw value or values averaged across the number
    // of pods targeted by the HPA (valid for all metric types).
	AverageValue *resource.Quantity

    // averageUtilization is the average value (as defined above) as
    // a percentage of the corresponding average pod request (valid
    // for resource metrics).
	AverageUtilization *int32
}

// and similarly for the statuses:

type MetricValueStatus struct {
    // value is the current value of the metric (as a quantity).
    // +optional
    Value *resource.Quantity
    // averageValue is the current value of the average of the
    // metric across all relevant pods (as a quantity)
    // (always reported for resource metrics)
    // +optional
    AverageValue *resource.Quantity
    // currentAverageUtilization is the current value of the average of the
    // resource metric across all relevant pods, represented as a percentage of
    // the requested value of the resource for the pods.
    // +optional
    AverageUtilization *int32
}

```

Notice that the `metricName` field is replaced with a new `metric` field,
which encapsulates both the metric name, and an optional label selector,
which takes the form of a standard kubernetes label selector.

The `targetXXX` fields are replaced by a unified `Target` field that
contains the different target types. The `target` field in the Object
metric source type is renamed to `describedObject`, since the `target`
field is now taken, and to more accurately describe its purpose.

The `External` source is updated slightly to match the new form of the
`Pods` and `Object` sources.

These changes necessitate a second beta of `autoscaling/v2`:
`autoscaling/v2beta2`.

Similarly, corresponding changes need to be made to the custom metrics
API:

```go
type MetricValue struct {
	metav1.TypeMeta
	DescribedObject ObjectReference

	Metric MetricIdentifier

	Timestamp metav1.Time
	WindowSeconds *int64
	Value resource.Quantity
}

type MetricIdentifier struct {
	// name is the name of the given metric
	Name string
	// selector represents the label selector that could be used to select
	// this metric, and will generally just be the selector passed in to
	// the query used to fetch this metric.
	// +optional
	Selector *metav1.LabelSelector
}
```

This will also require bumping the custom metrics API to
`custom.metrics.k8s.io/v1beta2`.

**Note that if a metrics pipeline works in such a way that multiple series
are matched by a label selector, it's the metrics adapter's job to deal
with it, similarly to the way things current work with the custom metrics
API.**

### Risks and Mitigations

The main risk around this proposal revolves around metric backend support.
When crafting the initial API, there were two constraints: a) limit
ourselves to an API surface that could be limited in adapters without any
additional processing of metrics, and b) avoid creating a new query
language.

There are currently three adapter implementations (known to SIG
Autoscaling): Prometheus, Stackdriver, and Sysdig.  Of those three, both
Prometheus and Stackdriver map nicely to the `name+labels` abstraction,
while Sysdig does not seem to natively have a concept of labels.  However,
this simply means that users of sysdig metrics will not make use of labels
-- there should be no need for the sysdig adapter to do anything special
with the labels besides ignore them.  The "name+label" paradigm also seems
to match nicely with other metric solutions (InfluxDB, DataDog, etc) used
with Kubernetes.

As for moving closer to a query language, this change is still very
structured and very limitted.  It requires no additional parsing logic
(since it uses standard kubernetes label selectors), and translation to
underlying APIs and query languages should be relatively simple.

## Graduation Criteria

In general, we'll want to graduate the autoscaling/v2 and
custom.metrics.k8s.io APIs to GA once we have a release with at least one
adapter up to date, and positive user feedback that does not suggest
urgent need for further changes.

## Implementation History

- (2018/4/19) Proposal proposed
- (2018/8/27) Implementation (kubernetes/kubernetes#64097) merged for Kubernetes 1.12

## Alternatives

- Continuing to require out-of-band changes to support more complex metric
  environments: this induces a lot of friction with traditional
  Prometheus-style monitoring setups, which favor selecting on labels.
  Furthermore, the changes required often involve admin intervention,
  which is not always simple or scalable in larger environments.

- Allow passing full queries instead of metric names: this would make the
  custom metrics API significantly more scalable, at the cost of adapter
  complexity, security issues, and lesser portability.  Effectively,
  adapters would have to implement query rewriting to inject extra labels
  in to scope metrics down to their target objects, which could in turn
  cause security issues. Additionally, it makes it a lot hard to port the
  HPAs between different metrics solutions.
