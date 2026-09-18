---
name: k8s-events-reviewer
description: Review a kubernetes/kubernetes pull request that changes Kubernetes Events — the events.k8s.io / core v1 Event API, EventRecorder / EventBroadcaster (client-go tools/events and tools/record), event reason/message/type conventions, the spam filter / aggregation cache, and event recorder lifecycle. Applies SIG-Instrumentation approver standards. Use when asked to review a k8s PR/diff touching events, EventRecorder, EventBroadcaster, event spam/aggregation, or events.k8s.io. Trigger phrases: "review this events PR", "sig-instrumentation events review".
---

# SIG-Instrumentation Events Review

Review a `kubernetes/kubernetes` PR the way a **SIG-Instrumentation approver** would for **Events**
changes.

This skill reviews **only** events-relevant changes. If the PR doesn't touch events, say so and stop.
(Use `sig-instrumentation-metrics-review` for metrics, `sig-instrumentation-log-review` for logging.)

## Step 1 — Get the change
- PR number/URL → `gh pr view <n> --repo kubernetes/kubernetes` + `gh pr diff <n> --repo kubernetes/kubernetes`.
- Local branch → `git diff` vs merge base. Pasted diff → use directly. No `gh`/auth → GitHub API or ask.

## Step 2 — Scope to events
In-scope signals (any of):
- Touches `staging/src/k8s.io/client-go/tools/events` or `…/tools/record` (EventRecorder,
  EventBroadcaster, `events_cache.go`, the spam/aggregation filter).
- Changes `Eventf` / `Event` call sites, `WithLogger`/`EventRecorderLogger`, or event recorder wiring
  in a controller's `Run`.
- Touches the `events.k8s.io` API, core `v1.Event` validation, or event quota/retention.
- Edits `test/e2e/instrumentation/core_events.go`.

If none apply, report "not an events change — out of scope" and stop.

## Step 3 — Review checklist
For each finding cite file:line and tag **BLOCKER** / **NIT** / **QUESTION**.

### A. Event content conventions
- **`Type`** is exactly `Normal` or `Warning` — nothing else (BLOCKER otherwise).
- **`Reason`** is a short, machine-readable UpperCamelCase token from a small, stable set (meant for
  switch/alerting) — not a sentence, not user-supplied free text.
- **`Note`/message** is a human-readable complete sentence; **don't** embed high-cardinality/unbounded
  values that defeat aggregation, and never log secrets/PII.
- **`action`** (events/v1) describes the operation; `regarding` is the primary object, `related` the
  secondary. Verify both are set correctly.

### B. Spam / aggregation
- Don't make the aggregation/spam key more granular without justification — it must still throttle a
  faulty object effectively. Use the broadcaster's canonical event key for dedup, not ad-hoc keys.
- Flag new code paths that emit events in tight loops / per-reconcile without relying on aggregation.

### C. Recorder/Broadcaster API & lifecycle
- Prefer `events.k8s.io/v1` (`client-go/tools/events`) for new recorders.
- Don't break published interfaces; add `*WithContext` / `EventRecorderLogger` variants instead, and
  don't deprecate `Eventf`. Use the `//logcheck:context` directive to steer callers.
- Construct/own the recorder where there's a proper cancellable context (typically `Run`); avoid races
  where informer handlers emit events before the recorder exists.

### D. Immutability / validation
- Respect immutable Event fields (`message`, etc.); BLOCKER on update paths that mutate them or tests
  that assume mutability.

### E. Tests
- Behavior tests should exercise real emission paths; core events e2e lives in
  `test/e2e/instrumentation/core_events.go` (SIG-Instrumentation owned).

## Step 4 — Output
```
## SIG-Instrumentation Events Review — PR #<n>: <title>

**Scope:** <which events code/files this touches>
**Verdict:** /lgtm | /approve | changes requested | needs SIG discussion

### Blockers
- [file:line] <issue> — <why> → <fix>

### Nits
- [file:line] <suggestion>

### Questions for author
- <question>

### Required before merge
- consider client-go consumer blast radius for any tools/events|record API change
- verify events.k8s.io/v1 path; keep Eventf non-breaking
```
Be concrete, cite lines, explain the *why* (immutability contract, spam-throttle effectiveness,
goroutine cancellation, client-go API blast radius) — not just the rule.

## Reference
- `OWNERS_ALIASES` → `sig-instrumentation-approvers`.
- EventRecorder/Broadcaster: `staging/src/k8s.io/client-go/tools/events` (events.k8s.io/v1) and
  `…/tools/record` (legacy core/v1); spam filter in `…/tools/record/events_cache.go`.
- Event field semantics: Kubernetes API conventions (`Type`/`Reason`/`Note`/`action`/`regarding`).
- Core events e2e: `test/e2e/instrumentation/core_events.go`.

