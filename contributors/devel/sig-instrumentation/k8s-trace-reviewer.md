---
name: k8s-trace-reviewer
description: Review a kubernetes/kubernetes pull request that changes distributed tracing — API Server tracing (KEP-647), kubelet tracing (KEP-2831), component-base/tracing, OpenTelemetry (OTel) spans/attributes/propagation, the OTLP exporter, and TracingConfiguration. Applies SIG-Instrumentation approver standards. Use when asked to review a k8s PR/diff touching tracing, OpenTelemetry, spans, trace context/propagation, or otel dependency bumps. Trigger phrases: "review this tracing PR", "sig-instrumentation traces review", "is this span/attribute correct".
---

# SIG-Instrumentation Traces Review

Review a `kubernetes/kubernetes` PR the way a **SIG-Instrumentation approver** would for **tracing**
changes. 

This skill reviews **only** tracing-relevant changes. If the PR doesn't touch tracing, say so and stop.
(Use the metrics/log/events review skills for those pillars.)

## Step 1 — Get the change
- PR number/URL → `gh pr view <n> --repo kubernetes/kubernetes` + `gh pr diff <n> --repo kubernetes/kubernetes`.
- Local branch → `git diff` vs merge base. Pasted diff → use directly. No `gh`/auth → GitHub API or ask.

## Step 2 — Scope to tracing
In-scope signals (any of):
- Touches `staging/src/k8s.io/component-base/tracing` (or older `component-base/traces`),
  `k8s.io/utils/trace`, or `tracer.Start` / `span.*` / `otel*` calls.
- Adds/changes propagators, `TracerProvider` wiring, `WithTracing`/`otelhttp`/`otelgrpc` handlers, or

  `TracingConfiguration`.
- Bumps `go.opentelemetry.io/otel*` or `…/contrib/…` dependencies.
- Edits `test/integration/apiserver/tracing/tracing_test.go`.

If none apply, report "not a tracing change — out of scope" and stop.

## Step 3 — Review checklist
For each finding cite file:line and tag **BLOCKER** / **NIT** / **QUESTION**.

### A. Provider & propagation
- **BLOCKER**: use of the global `TracerProvider`/propagator. Pass them explicitly per client/server,
  reusing the single provider constructed at component start (kubedeps / apiserver config).
- Disabled/unset → `NoopTracerProvider`, not nil; instrumentation is feature-gated and no-op safe.

### B. Security
- Tracing/instrumentation sits **after authentication & authorization**. Unauthenticated/read-only
  endpoints use `WithPublicEndpoint` (ignore incoming context) or omit tracing — callers must not
  control sampling. **BLOCKER** otherwise.

### C. Spans
- `defer span.End()` immediately after `tracer.Start`; no code path leaves a span unclosed.
- Span names follow OTel conventions (no spaces, dotted/namespaced); scope spans to bounded work, not
  long-lived watches/streams.
- Errors via `span.RecordError(err)` + error status, not as attributes.

### D. Attributes
- OTel semantic conventions: dotted, namespaced keys (`api_server.id`, `transformer.provider.name`);
  prefer `semconv` constants over hand-rolled strings. Avoid unbounded-cardinality attribute values.

### E. Config & deps
- Init/config functions take `TracingConfiguration`/`resourceOpts` for forward-compat; rely on OTel
  defaults. otel/contrib bumps keep the dependency graph consistent; semconv-transition shims carry a
  removal plan + tracking issue.

### F. Tests
- New spans/attributes are asserted in `test/integration/apiserver/tracing/tracing_test.go` (or the
  equivalent kubelet tracing test). A span with no test assertion is incomplete.

## Step 4 — Output
```
## SIG-Instrumentation Traces Review — PR #<n>: <title>

**Scope:** <which tracing code/files this touches>
**Verdict:** /lgtm | /approve | changes requested | needs SIG discussion

### Blockers
- [file:line] <issue> — <why> → <fix>

### Nits
- [file:line] <suggestion>

### Questions for author
- <question>

### Required before merge
- assert new spans/attributes in test/integration/apiserver/tracing/tracing_test.go
- confirm no global TracerProvider/propagator use; Noop when disabled; instrumentation after authz
```
Be concrete, cite lines, explain the *why* (sampling-control security, span leaks, OTel semconv
interop, dependency-graph consistency) — not just the rule.

## Reference
- `OWNERS_ALIASES` → `sig-instrumentation-approvers`.
- API Server tracing: KEP-647; kubelet tracing: KEP-2831.
- Tracing helpers: `staging/src/k8s.io/component-base/tracing`; `TracingConfiguration` in
  `component-base/tracing/api/v1`. Slow-trace logging util: `k8s.io/utils/trace`.
- OpenTelemetry semantic conventions (`go.opentelemetry.io/otel/semconv`).
- Integration test: `test/integration/apiserver/tracing/tracing_test.go`.

