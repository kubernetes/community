Horizontal Pod Autoscaler with Arbitrary Metrics
===============================================

The current Horizontal Pod Autoscaler object only has support for CPU as
a percentage of requested CPU.  While this is certainly a common case, one
of the most frequently sought-after features for the HPA is the ability to
scale on different metrics (be they custom metrics, memory, etc).

The current HPA controller supports targeting "custom" metrics (metrics
with a name prefixed with "custom/") via an annotation, but this is
suboptimal for a number of reasons: it does not allow for arbitrary
"non-custom" metrics (e.g. memory), it does not allow for metrics
describing other objects (e.g. scaling based on metrics on services), and
carries the various downsides of annotations (not be typed/validated,
being hard for a user to hand-construct, etc).

Object Design
-------------

### Requirements ###

This proposal describes a new version of the Horizontal Pod Autoscaler
object with the following requirements kept in mind:

1. The HPA should continue to support scaling based on percentage of CPU
   request

2. The HPA should support scaling on arbitrary metrics associated with
   pods

3. The HPA should support scaling on arbitrary metrics associated with
   other Kubernetes objects in the same namespace as the HPA (and the
   namespace itself)

4. The HPA should make scaling on multiple metrics in a single HPA
   possible and explicit (splitting metrics across multiple HPAs leads to
   the possibility of fighting between HPAs)

### Specification ###

```go
type HorizontalPodAutoscalerSpec struct {
    // the target scalable object to autoscale
    ScaleTargetRef CrossVersionObjectReference `json:"scaleTargetRef"`

    // the minimum number of replicas to which the autoscaler may scale
    // +optional
    MinReplicas *int32 `json:"minReplicas,omitempty"`
    // the maximum number of replicas to which the autoscaler may scale
    MaxReplicas int32 `json:"maxReplicas"`

    // the metrics to use to calculate the desired replica count (the
    // maximum replica count across all metrics will be used).  The
    // desired replica count is calculated multiplying the ratio between
    // the target value and the current value by the current number of
    // pods.  Ergo, metrics used must decrease as the pod count is
    // increased, and vice-versa.  See the individual metric source
    // types for more information about how each type of metric
    // must respond.
    // +optional
    Metrics []MetricSpec `json:"metrics,omitempty"`
}

// a type of metric source
type MetricSourceType string
var (
    // a metric describing a kubernetes object (for example, hits-per-second on an Ingress object)
    ObjectSourceType MetricSourceType = "Object"
    // a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
    // The values will be averaged together before being compared to the target value
    PodsSourceType MetricSourceType = "Pods"
    // a resource metric known to Kubernetes, as specified in requests and limits, describing each pod
    // in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes,
    // and have special scaling options on top of those available to normal per-pod metrics (the "pods" source)
    ResourceSourceType MetricSourceType = "Resource"
)

// a specification for how to scale based on a single metric
// (only `type` and one other matching field should be set at once)
type MetricSpec struct {
    // the type of metric source (should match one of the fields below)
    Type MetricSourceType `json:"type"`

    // a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object)
    Object *ObjectMetricSource `json:"object,omitempty"`
    // a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
    // The values will be averaged together before being compared to the target value
    Pods *PodsMetricSource `json:"pods,omitemtpy"`
    // a resource metric (such as those specified in requests and limits) known to Kubernetes
    // describing each pod in the current scale target (e.g. CPU or memory). Such metrics are
    // built in to Kubernetes, and have special scaling options on top of those available to
    // normal per-pod metrics using the "pods" source.
    Resource *ResourceMetricSource `json:"resource,omitempty"`
}

// a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object)
type ObjectMetricSource struct {
    // the described Kubernetes object
    Target CrossVersionObjectReference `json:"target"`

    // the name of the metric in question
    MetricName string `json:"metricName"`
    // the target value of the metric (as a quantity)
    TargetValue resource.Quantity `json:"targetValue"`
}

// a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
// The values will be averaged together before being compared to the target value
type PodsMetricSource struct {
    // the name of the metric in question
    MetricName string `json:"metricName"`
    // the target value of the metric (as a quantity)
    TargetAverageValue resource.Quantity `json:"targetAverageValue"`
}

// a resource metric known to Kubernetes, as specified in requests and limits, describing each pod
// in the current scale target (e.g. CPU or memory).  The values will be averaged together before
// being compared to the target.  Such metrics are built in to Kubernetes, and have special
// scaling options on top of those available to normal per-pod metrics using the "pods" source.
// Only one "target" type should be set.
type ResourceMetricSource struct {
    // the name of the resource in question
    Name api.ResourceName `json:"name"`
    // the target value of the resource metric, represented as
    // a percentage of the requested value of the resource on the pods.
    // +optional
    TargetAverageUtilization *int32 `json:"targetAverageUtilization,omitempty"`
    // the target value of the resource metric as a raw value, similarly
    // to the "pods" metric source type.
    // +optional
    TargetAverageValue *resource.Quantity `json:"targetAverageValue,omitempty"`
}

type HorizontalPodAutoscalerStatus struct {
    // most recent generation observed by this autoscaler.
    ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
    // last time the autoscaler scaled the number of pods;
    // used by the autoscaler to control how often the number of pods is changed.
    LastScaleTime *unversioned.Time `json:"lastScaleTime,omitempty"`

    // the last observed number of replicas from the target object.
    CurrentReplicas int32 `json:"currentReplicas"`
    // the desired number of replicas as last computed by the autoscaler
    DesiredReplicas int32 `json:"desiredReplicas"`

    // the last read state of the metrics used by this autoscaler
    CurrentMetrics []MetricStatus `json:"currentMetrics" protobuf:"bytes,5,rep,name=currentMetrics"`
}

// the status of a single metric
type MetricStatus struct {
    // the type of metric source
    Type MetricSourceType `json:"type"`

    // a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object)
    Object *ObjectMetricStatus `json:"object,omitemtpy"`
    // a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
    // The values will be averaged together before being compared to the target value
    Pods *PodsMetricStatus `json:"pods,omitemtpy"`
    // a resource metric known to Kubernetes, as specified in requests and limits, describing each pod
    // in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes,
    // and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    Resource *ResourceMetricStatus `json:"resource,omitempty"`
}

// a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object)
type ObjectMetricStatus struct {
    // the described Kubernetes object
    Target CrossVersionObjectReference `json:"target"`

    // the name of the metric in question
    MetricName string `json:"metricName"`
    // the current value of the metric (as a quantity)
    CurrentValue resource.Quantity `json:"currentValue"`
}

// a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
// The values will be averaged together before being compared to the target value
type PodsMetricStatus struct {
    // the name of the metric in question
    MetricName string `json:"metricName"`
    // the current value of the metric (as a quantity)
    CurrentAverageValue resource.Quantity `json:"currentAverageValue"`
}

// a resource metric known to Kubernetes, as specified in requests and limits, describing each pod
// in the current scale target (e.g. CPU or memory).  The values will be averaged together before
// being compared to the target.  Such metrics are built in to Kubernetes, and have special
// scaling options on top of those available to normal per-pod metrics using the "pods" source.
// Only one "target" type should be set.  Note that the current raw value is always displayed
// (even when the current values as request utilization is also displayed).
type ResourceMetricStatus struct {
    // the name of the resource in question
    Name api.ResourceName `json:"name"`
    // the target value of the resource metric, represented as
    // a percentage of the requested value of the resource on the pods
    // (only populated if the corresponding request target was set)
    // +optional
    CurrentAverageUtilization *int32 `json:"currentAverageUtilization,omitempty"`
    // the current value of the resource metric as a raw value
    CurrentAverageValue resource.Quantity `json:"currentAverageValue"`
}
```

### Example ###

In this example, we scale based on the `hits-per-second` value recorded as
describing a service in our namespace, plus the CPU usage of the pods in
the ReplicationController being autoscaled.

```yaml
kind: HorizontalPodAutoscaler
apiVersion: autoscaling/v2alpha1
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: WebFrontend
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80
  - type: Object
    object:
      target:
        kind: Service
        name: Frontend
      metricName: hits-per-second
      targetValue: 1k
```

### Alternatives and Future Considerations ###

Since the new design mirrors volume plugins (and similar APIs), it makes
it relatively easy to introduce new fields in a backwards-compatible way:
we simply introduce a new field in `MetricSpec` as a new "metric type".

#### External ####

It was discussed adding a source type of `External` which has a single
opaque metric field and target value.  This would indicate that the HPA
was under control of an external autoscaler, which would allow external
autoscalers to be present in the cluster while still indicating to tooling
that autoscaling is taking place.

However, since this raises a number of questions and complications about
interaction with the existing autoscaler, it was decided to exclude this
feature.  We may reconsider in the future.

#### Limit Percentages ####

In cluster environments where request is automatically set for scheduling
purposes, it is advantageous to be able to autoscale on percentage of
limit for resource metrics.  We may wish to consider adding
a `targetPercentageOfLimit` to the `ResourceMetricSource` type.

#### Referring to the current Namespace ####

It is beneficial to be able to refer to a metric on the current namespace,
similarly to the `ObjectMetricSource` source type, but without an explicit
name.  Because of the similarity to `ObjectMetricSource`, it may simply be
sufficient to allow specificying a `kind` of "Namespace" without a name.
Alternatively, a similar source type to `PodsMetricSource` could be used.

#### Calculating Final Desired Replica Count ####

Since we have multiple replica counts (one from each metric), we must have
a way to aggregated them into a final replica count.  In this iteration of
the proposal, we simply take the maximum of all the computed replica
counts.  However, in certain cases, it could be useful to allow the user
to specify that they wanted the minimum or average instead.

In the general case, maximum should be sufficient, but if the need arises,
it should be fairly easy to add such a field in.

Mechanical Concerns
-------------------

The HPA will derive metrics from two sources: resource metrics (i.e. CPU
request percentage) will come from the
[master metrics API](resource-metrics-api.md), while other metrics will
come from the [custom metrics API](custom-metrics-api.md), which is
an adapter API which sources metrics directly from the monitoring
pipeline.

Conversion and Defaulting
-------------------------

### Conversion ###

According the deprecation policy, objects must be round-trippable through
all supported versions.  Therefore, two alpha annotations will be
introduced: `autoscaling.alpha.kubernetes.io/metrics` and
`autoscaling.alpha.kubernetes.io/current-metrics`.  These will correspond
roughly to the `Metrics` and `CurrentMetrics` fields in `v2alpha1`.

In the case of conversion from `v1` to `v2alpha`, the value of the
appropriate CPU field will be read, converted into a `MetricSpec` of type
`Resource`, and then appended to the list of existing `MetricSpec`s to
form the slice for the `Metrics` field in `v2alpha1` (and similarly for
`MetricStatus`es and `CurrentMetrics`).

In the opposite direction, the list of metrics will be searched for
a field of type `Resource` with a name of `api.ResourceCPU` and a set
utilization percentage.  This will be removed from the list and used to
populate the CPU field in `v1`, while the remainder of the `MetricSpec`s
(and `MetricStatus`es) will be placed into the appropriate annotation.

#### Example ####

Consider the example object described above.  It would appear in
`autoscaling/v1` as

```yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/metrics: "[{\"type\": \"Object\", \"object\": {\"target\": {\"kind\": \"Service\", \"name\": \"Frontend\"}, \"metricName\": \"hits-per-second\", \"targetValue\": \"1k\"}}]"
    object:
      target:
        kind: Service
        name: Frontend
      metricName: hits-per-second
      targetValue: 1k
spec:
  scaleTargetRef:
    kind: ReplicationController
    name: WebFrontend
  minReplicas: 2
  maxReplicas: 10
  TargetCPUUtilizationPercentage: 80
```

### Defaulting ###

Currently, in `autoscaling/v1`, no API-level defaulting is done -- if the
user creates an HPA without a set CPU autoscaling level, the controller
will look at whether or not the custom metrics annotation is set.  If it
is, the controller will scale on custom metrics, and if not, it will use
a "default" value built in to the controller.

Since `v2alpha1` has field-level "custom" metrics, we no longer need to
rely on this "implicit" defaulting.  If a user creates an empty HPA, it
will be populated with an explicit default CPU policy.  This makes it
clearer to the end user how their HPA is scaling, and makes the jump from
"default HPA" to "customized HPA" easier.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/hpa-v2.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
