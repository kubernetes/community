## Instrumenting Kubernetes with Traces

The following references and outlines general guidelines for trace instrumentation
in Kubernetes components. Components are instrumented using the
[OpenTelemetry Go client library](https://github.com/open-telemetry/opentelemetry-go).
For non-Go components. [Libraries in other languages](https://opentelemetry.io/docs/languages/)
are available.

Traces are exposed via gRPC using the [OpenTelemetry Protocol](https://opentelemetry.io/docs/specs/otel/protocol/)
(OTLP), which is open and well-understood by a wide range of third party
applications and vendors in the cloud-native eco-system.

The [general instrumentation advice](https://opentelemetry.io/docs/concepts/instrumentation/libraries/)
from the OpenTelemetry documentation applies. This document reiterates common pitfalls and some
Kubernetes specific considerations.

### When to instrument

While spans are sampled to avoid high costs, recording too many spans will
force consumers to lower the sampling rate, and will "drown out" important
spans. If your component has more than two or three nested spans, you are
likely over-using trace instrumentation. Most trace instrumentation in
Kubernetes components falls into one of two categories:

1. Spans for incoming or outgoing network calls
2. Spans when initiating new work, such as reconciling an object, which may result in network calls.

For network-based telemetry, Kubernetes components should use OpenTelemetry
instrumentation libraries for
[HTTP](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp) and
[gRPC](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc).

**Note:** When creating spans at the start of reconciling an Object, only
create the span changes are actually required. Avoid creating "empty" spans
which simply compare the desired and actual state of an object without
performing any real work, or making any network requests.

### Configuration and Setup

Kubernetes components should expose a flag, `--tracing-config-file`, which accepts a
[TracingConfiguration](https://github.com/kubernetes/kubernetes/blob/cb989b84c82ab8c695d02c9a1866bc27de37caec/staging/src/k8s.io/component-base/tracing/api/v1/types.go)
object. The `component-base/tracing` library provides a `NewProvider()` helper
to convert a TracingConfiguration to a TracerProvider, which can be used to
record spans. Components should avoid using OpenTelemetry globals, and instead
pass the configured TracerProvider to libraries where they are used. Components
should use the W3C Traceparent and Baggage propagators, as provided by the
`Propagators()` helper.

### Context Propagation

Generally, components should not interact directly with OpenTelemetry
Propagators, other than by passing them to libraries. Context propagation
across network boundaries is handled by the
[HTTP](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp) and
[gRPC](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc)
network client and server instrumentation libraries.

Components need to propagate Golang's `context.Context` from incoming network
calls or spans from the initiation of new work to any outgoing network calls to
ensure spans are properly connected into traces.

### Naming and Style

Follow the OpenTelemetry [guidelines for span naming](https://opentelemetry.io/docs/specs/otel/trace/api/#span), and the OpenTelemetry [guidelines for attributes](https://opentelemetry.io/docs/specs/semconv/general/attribute-naming/).

### Tracing stability

Tracing instrumentation in Kubernetes components does not currently have
stability guarantees. Component owners should avoid breaking context
propagation, or removing trace instrumentation entirely. Component owners may
rename spans, or span attributes.
