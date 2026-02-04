# Metric Stability

[KEP-1209] introduced the concept of metric stability to Kubernetes. When metrics graduate to the `STABLE` `StabilityLevel`, we provide guarantees to consumers of those metrics so they can confidently build alerting and monitoring platforms.

## Stability Classes

There are currently four stability classes for metrics: (1) Alpha, (2) Beta, (3) Stable, and (4) Internal. These classes are intended to make the API contract between the control-plane and the consumer of control-plane metrics explicit.

### Alpha

__Alpha__ metrics have __*no*__ stability guarantees; as such they can be modified or deleted at any time. All Kubernetes metrics begin as alpha metrics.

Alpha metrics are new metrics with minimal or no production experience. There is no expectation for users to monitor them yet, as they may belong to Alpha features or haven't yet proven their long-term value. Due to the lack of production feedback, there is a high probability these metrics will change or be removed entirely.

An example of an alpha metric is as follows:

```go
var alphaMetricDefinition = kubemetrics.CounterOpts{
    Name: "some_alpha_metric",
    Help: "some description",
    StabilityLevel: kubemetrics.ALPHA, // this is also a custom metadata field
	DeprecatedVersion: "1.15", // this can optionally be included on alpha metrics, although there is no change to contractual stability guarantees
}
```

### Beta

__Beta__ metrics have a looser stability contract than their stable counterparts. No labels can be removed from beta metrics during their lifetime, however, labels can be added while the metric is in the beta stage.
Beta metrics can also be marked as __deprecated__ for a future Kubernetes version.

Beta metrics are metrics that have been validated in production and deemed useful. While not yet critical, it is recommended that users begin monitoring them. The expectation is that metrics for GA features and critical metrics for Beta features eventually graduate to this level. While less likely than Alpha, these metrics might still change if they are tightly coupled to internal implementation details.

An example of a beta metric follows:

```go
var betaMetricDefinition = kubemetrics.CounterOpts{
    Name: "some_beta_metric",
    Help: "some description",
    StabilityLevel: kubemetrics.BETA,
    DeprecatedVersion: "1.15", // this is a custom metadata field
}
```

By the beta stability contract, we mean:

1. the metric will not be deleted without graduating to stable first or being deprecated for a minimum of 1 release (or 4 months, whichever is longer)
2. the type of metric will not be modified
3. no labels can be removed from this metric
4. labels **can** be added to this metric while in beta

### Stable

__Stable__ metrics can be guaranteed to *not change*, except that the metric may become marked deprecated for a future Kubernetes version.

Stable metrics are crucial metrics covering core Kubernetes features that are expected to change very rarely, if ever. Users should prioritize monitoring these metrics to understand the fundamental health and performance of their clusters.

An example of a stable metric follows:

```go
var deprecatedMetricDefinition = kubemetrics.CounterOpts{
    Name: "some_deprecated_metric",
    Help: "some description",
    StabilityLevel: kubemetrics.STABLE, // this is also a custom metadata field
    DeprecatedVersion: "1.15", // this is a custom metadata field
}
```

By *not change*, we mean three things:

1. the metric itself will not be deleted ([or renamed](#metric-renaming)) without being deprecated for a minimum of 3 releases (or 9 months, whichever is longer)
2. the type of metric will not be modified
3. no labels can be added **or** removed from this metric

From an ingestion point of view, it is backwards-compatible to add or remove possible __values__ for labels which already do exist (but __not__ labels themselves). Therefore, adding or removing __values__ from an existing label is permissible. Stable metrics can also be marked as __deprecated__ for a future Kubernetes version, since this is a metadata field and does not actually change the metric itself.

**Removing or adding labels from stable metrics is not permissible.** In order to add/remove a label to an existing stable metric, one would have to introduce a new metric and deprecate the stable one; otherwise this would violate compatibility agreements.

### Internal

__Internal__ metrics are intended for metrics that are used for internal purposes only and are not meant to be consumed by end users. These metrics have __*no*__ stability guarantees and can be modified or deleted at any time, similar to alpha metrics. However, they are explicitly marked as internal to signal that they are not part of the public API.

An example of an internal metric follows:

```go
var internalMetricDefinition = kubemetrics.CounterOpts{
    Name: "some_internal_metric",
    Help: "some description",
    StabilityLevel: kubemetrics.INTERNAL,
}
```

Internal metrics:

1. have no stability guarantees
2. can be modified or deleted at any time
3. should not be relied upon by external monitoring or alerting systems
4. are typically used for debugging, testing, or internal component communication

## API Review

Graduating a metric to beta or stable is a contractual API agreement (in line with current Kubernetes [api-review processes](https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md)).

For **Beta** metrics, approval from SIG Instrumentation is required.

For **Stable** metrics, marking a metric as stable is a commitment by the owning SIG to maintain stability guarantees. The owning SIG leads must review and approve the graduation first. Additionally, approval from SIG Instrumentation is required.

We use a verification script to flag stable metric changes for review by SIG Instrumentation approvers.

## Metric Renaming

Metric renaming is tantamount to deleting a metric and introducing a new one. Accordingly, metric renaming will also be disallowed for stable metrics.

## Deprecation Lifecycle

Metrics can be annotated with a Kubernetes version, from which point that metric will be considered deprecated. This allows us to indicate that a metric is slated for future removal and provides the consumer a reasonable window in which they can make changes to their monitoring infrastructure which depends on this metric.

While a deprecation period is only required for __stable__ metrics, an alpha or beta metric may still be deprecated prior to removal to help component owners inform users of future intent, and to ease the transition to the replacement metrics.

When a stable metric undergoes the deprecation process, we are signaling that the metric will eventually be deleted. The lifecyle looks roughly like this (each stage represents a Kubernetes release):

__Stable metric__ -> __Deprecated metric__ -> __Hidden metric__ -> __Deletion__

__Deprecated__ metrics have the same stability guarantees of their counterparts. If a stable metric is deprecated, then a deprecated stable metric is guaranteed to *not change*. When deprecating a stable metric, a future Kubernetes release is specified as the point from which the metric will be considered deprecated.

```go
var someCounter = kubemetrics.CounterOpts{
    Name: "some_counter",
    Help: "this counts things",
    StabilityLevel: kubemetrics.STABLE,
    DeprecatedVersion: "1.15", // this metric is deprecated when the Kubernetes version == 1.15
}
````

__Deprecated__ metrics will have their description text prefixed with a deprecation notice string '(Deprecated since x.y)' and a warning log will be emitted during metric registration (in the spirit of the official [Kubernetes deprecation policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-a-flag-or-cli)).

Before deprecation:

```text
# HELP some_counter this counts things
# TYPE some_counter counter
some_counter 0
```

During deprecation:

```text
# HELP some_counter (Deprecated since 1.15) this counts things
# TYPE some_counter counter
some_counter 0
```
Like their stable metric counterparts, deprecated metrics will be automatically registered to the metrics endpoint.

After a period of time, a deprecated metric will become a __hidden metric__. The timeline for this transition depends on the metric's stability level:

| Stability Level | Time until hidden                                           |
| --------------- | ----------------------------------------------------------- |
| **STABLE**      | Minimum of 3 releases or 9 months, whichever is longer      |
| **BETA**        | Minimum of 1 release or 4 months, whichever is longer       |
| **ALPHA**       | Can be hidden or removed in the same release as deprecation |

_Unlike_ their deprecated counterparts, hidden metrics will __*no longer be automatically registered*__ to the metrics endpoint (hence hidden). However, they can be explicitly enabled through a command line flag on the binary (i.e. '--show-hidden-metrics-for-version=<previous minor release>'). This is to provide cluster admins an escape hatch to properly migrate off of a deprecated metric, if they were not able to react to the earlier deprecation warnings. Hidden metrics will be deleted after one release.


[KEP-1209]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1209-metrics-stability
