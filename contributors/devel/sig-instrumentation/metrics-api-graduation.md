# Metrics APIs graduation guide

This document provides a checklist for graduating the Kubernetes metrics APIs:

- `metrics.k8s.io`
- `custom.metrics.k8s.io`
- `external.metrics.k8s.io`

These APIs are served through the Kubernetes aggregation layer. Graduating one
of them therefore requires coordinating the API definition, serving
implementations, discovery behavior, clients, compatibility, and testing.

[KEP-5207](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/5207-metrics-k8s-io-api-definition),
which graduates `metrics.k8s.io` to `v1`, is used as a worked example.

## Graduation checklist

### API definition and generated clients

Verify that the new API version has the required:

- API types
- registration
- conversions
- generated clients
- OpenAPI definitions
- round-trip and client tests

The stable version should preserve compatibility guarantees defined by the
graduation KEP.

For the `metrics.k8s.io/v1` graduation, the API and generated client work was
implemented in
[kubernetes/kubernetes#139223](https://github.com/kubernetes/kubernetes/pull/139223).

### Serving implementations

Identify implementations that serve the API before migrating consumers.

During a version transition, serving implementations may need to register the
new and previous versions simultaneously. This allows newer and older clients
to operate during version skew.

For each `APIService`, verify:

- the correct API group and version
- the backing Service
- `groupPriorityMinimum`
- `versionPriority`
- availability of every version expected by supported clients

`versionPriority` controls the ordering of versions inside an API group during
discovery. Higher values are preferred. Set it deliberately when multiple
versions are registered. If two Kubernetes-style versions have equal priority,
stable versions sort ahead of beta versions, and beta versions ahead of alpha
versions.

For the `metrics.k8s.io/v1` example, stable API serving and
`v1.metrics.k8s.io` APIService registration are implemented in
[kubernetes-sigs/metrics-server#1855](https://github.com/kubernetes-sigs/metrics-server/pull/1855).

### Aggregation layer behavior

An `APIService` represents one exact group/version. The aggregation layer does
not automatically convert a request for one aggregated API version into
another version.

If no `APIService` is registered for the exact version requested, the request
returns `404 NotFound`.

For example:

| Registered APIServices | `v1` request | `v1beta1` request |
| --- | --- | --- |
| `v1beta1` only | `404 NotFound` | Served |
| `v1` and `v1beta1` | Served | Served |
| `v1` only | Served | `404 NotFound` |

Do not confuse an unregistered version with a registered but unavailable
`APIService`. If the APIService exists but its backing service cannot serve the
request, the request can fail with `503 ServiceUnavailable`.

Clients should distinguish version absence from backend unavailability.
Fallback to an older API version can be appropriate when the requested version
is not served, but automatically falling back on backend failures can hide a
broken serving implementation.

### Consumers

Identify both in-tree and out-of-tree consumers of the API.

Check:

- which API versions each consumer supports
- how each consumer discovers available versions
- which version it selects when several are served
- whether it must work against older clusters or older adapters

The consumer set differs between the three metrics APIs.

The custom metrics client provides useful prior art for version transitions.
[`custom_metrics/discovery.go`](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/metrics/pkg/client/custom_metrics/discovery.go)
selects a supported version through API discovery, while
[`custom_metrics/multi_client.go`](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/metrics/pkg/client/custom_metrics/multi_client.go)
hides the selected version from callers and supports periodically invalidating
the cached selection so discovery can run again.

For resource metrics, important in-tree consumers include:

- `kubectl top`
- the HorizontalPodAutoscaler

Custom and external metrics are primarily consumed by autoscaling components
through their corresponding clients in `k8s.io/metrics`.

#### kubectl top

For the KEP-5207 example, support for both `metrics.k8s.io/v1` and
`v1beta1` in `kubectl top` is implemented in
[kubernetes/kubernetes#139726](https://github.com/kubernetes/kubernetes/pull/139726).

#### HorizontalPodAutoscaler

The HorizontalPodAutoscaler consumes resource, custom, and external metrics.

Migration of its resource metrics client for KEP-5207 is tracked in
[kubernetes/kubernetes#141366](https://github.com/kubernetes/kubernetes/pull/141366).

When graduating `custom.metrics.k8s.io` or `external.metrics.k8s.io`, review
the corresponding HPA client paths independently rather than assuming the
resource metrics migration applies to them.

### Upgrade, downgrade, and version skew

Test the transition from both the client and server side.

Consider:

- newer clients against servers that only expose the previous API version
- older clients against servers that expose both versions
- upgrading a serving implementation before its consumers
- upgrading consumers before the serving implementation
- downgrading the serving implementation
- independent release schedules for out-of-tree adapters

Where possible, consumers should use discovery to determine which versions are
actually served instead of assuming the new version is available.

### Tests

Plan test coverage early. Testing an aggregated API often requires more than
unit tests in the API client.

Cover, as applicable:

- API registration and generated code
- discovery of the new version
- discovery ordering when multiple versions are registered
- requests using the new version
- compatibility with the previous version
- consumer version selection
- `404 NotFound` for versions that are not registered
- serving failures independently from version absence
- upgrade and downgrade scenarios
- end-to-end consumer behavior

Also determine early whether existing Kubernetes end-to-end infrastructure can
serve the API versions needed by the test.

Aggregated metrics implementations may live outside
`kubernetes/kubernetes`. If the real implementation is unavailable in the
relevant e2e environment, a mock aggregated API server or test image may be
needed. Account for the additional review and image-promotion work when
planning the graduation.

The discussion in
[kubernetes/kubernetes#141366](https://github.com/kubernetes/kubernetes/pull/141366)
illustrates this issue: resource metrics e2e coverage cannot assume an in-tree
metrics-server, while Kubernetes already has test infrastructure for external
metrics.

## Differences between the metrics APIs

The checklist applies to all three API groups, but their serving
implementations and consumers differ.

| API group | Typical serving implementation | Main in-tree consumers |
| --- | --- | --- |
| `metrics.k8s.io` | metrics-server | `kubectl top`, HorizontalPodAutoscaler resource metrics |
| `custom.metrics.k8s.io` | custom metrics adapters | HorizontalPodAutoscaler custom metrics client |
| `external.metrics.k8s.io` | external metrics adapters | HorizontalPodAutoscaler external metrics client |

For custom and external metrics in particular, account for the independent
release schedules of out-of-tree serving implementations.

## KEP-5207 worked example

The `metrics.k8s.io/v1` graduation demonstrates how a change to an aggregated
metrics API spans several components.

| Area | Implementation or tracking reference |
| --- | --- |
| Stable API definition and generated code | [kubernetes/kubernetes#139223](https://github.com/kubernetes/kubernetes/pull/139223) |
| metrics-server `v1` support | [kubernetes-sigs/metrics-server#1855](https://github.com/kubernetes-sigs/metrics-server/pull/1855) |
| `kubectl top` consumer | [kubernetes/kubernetes#139726](https://github.com/kubernetes/kubernetes/pull/139726) |
| HorizontalPodAutoscaler consumer | [kubernetes/kubernetes#141366](https://github.com/kubernetes/kubernetes/pull/141366) |
| Aggregation and compatibility design | [KEP-5207](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/5207-metrics-k8s-io-api-definition) |
| Post-GA follow-up work | [kubernetes/kubernetes#141516](https://github.com/kubernetes/kubernetes/issues/141516) |

## References

- [KEP-5207: metrics.k8s.io API definition](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/5207-metrics-k8s-io-api-definition)
- [Metrics API post-GA tracking issue](https://github.com/kubernetes/kubernetes/issues/141516)
- [APIService API reference](https://kubernetes.io/docs/reference/kubernetes-api/apiregistration/api-service-v1/)
