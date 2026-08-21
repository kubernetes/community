---
name: k8s-metrics-reviewer
description: Review a kubernetes/kubernetes pull request that changes metrics/instrumentation code, applying SIG-Instrumentation approver standards (metric stability framework, naming conventions, cardinality, deprecation policy, component-base/metrics usage, and the verify-* tooling). Use when asked to review a k8s PR/diff that touches metrics, instrumentation, or the stable metrics list. Trigger phrases: "review this metrics PR", "sig-instrumentation review", "is this metric stable/named correctly".
---

# SIG-Instrumentation Metrics Review

Review a `kubernetes/kubernetes` pull request the way a **SIG-Instrumentation approver** would.
This skill reviews **only** metrics/instrumentation-relevant changes. If the PR doesn't touch
metrics, say so and stop.

## Step 1 — Get the change

Determine the target from the user's input:
- A PR number/URL → `gh pr view <n> --repo kubernetes/kubernetes` and `gh pr diff <n> --repo kubernetes/kubernetes`.
- A local checkout / current branch → `git diff` against the merge base.
- A pasted diff → use it directly.

If `gh` is unavailable or unauthenticated, fall back to the GitHub API
(`https://api.github.com/repos/kubernetes/kubernetes/pulls/<n>` and `.../files`) or ask the user to paste the diff.

## Step 2 — Scope to metrics

Confirm the change is in-scope. In-scope signals (any of):
- Touches `k8s.io/component-base/metrics` or imports it.
- Adds/changes a `metrics.NewCounterVec / NewGauge / NewHistogram / NewSummary` (the
  **component-base** wrappers, not raw `prometheus.NewX`).
- Edits `hack/tools/instrumentation/testdata/stable-metrics-list.yaml` (the generated stable-metrics snapshot)
  or `hack/tools/instrumentation/documentation/...`.
- Changes `StabilityLevel`, `DeprecatedVersion`, metric `Name`/`Subsystem`/`Namespace`, `Help`, `Buckets`, or labels.

If none apply, report "not a metrics change — out of SIG-Instrumentation scope" and stop.

## Step 3 — Review checklist

Go through each item. For every finding, cite the file:line and classify severity:
**BLOCKER** (must fix before approval), **NIT** (style/optional), **QUESTION** (need author clarification).

### A. Use the component-base wrappers, not raw Prometheus
- **BLOCKER** if new metrics are created with `prometheus.NewCounter/...` directly instead of
  `k8s.io/component-base/metrics`. The wrappers are what make the stability framework, hidden-metrics,
  and static analysis work. Raw client_golang metrics bypass all of it.
- Metrics must be registered through a `metrics.KubeRegistry` (e.g. `legacyregistry.MustRegister`),
  not `prometheus.MustRegister`.
- **BLOCKER** if a metric is registered inside an `init()` block. Registrations must happen dynamically during component initialization/startup (e.g. via an explicit `Register()` function called after feature gates have been initialized) so that metrics can be gated behind feature flags.

### B. Naming conventions (`hack/verify-metrics-naming.sh`)
- Follow the [Prometheus metric naming best practices](https://prometheus.io/docs/practices/naming/).
- snake_case, all lowercase.
- **Base units only**: `seconds` not `ms`/`milliseconds`, `bytes` not `kilobytes`. Append the unit.
- Counters MUST end in `_total`.
- Duration histograms/summaries SHOULD end in `_seconds` (e.g. `..._duration_seconds`).
- Don't encode label values into the metric name (no `requests_get_total` + `requests_post_total`;
  use a `verb` label instead — but watch cardinality, item C).
- Use a consistent `Namespace`/`Subsystem` prefix matching the component (e.g. `apiserver_`, `kubelet_`).

### C. Cardinality (the most common real-world problem)
- **BLOCKER/QUESTION** on unbounded label values: user IDs, full request paths, pod/namespace names,
  IPs, error strings, resource UIDs, full URLs. These explode series count and can OOM Prometheus.
- Prefer a bounded set. If a label must be bounded, require `metrics.ConstrainedLabels` /
  a [label-value allowlist](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2305-metrics-cardinality-enforcement), or bucket the values.
- Sanity-check: `cardinality ≈ Π(distinct values per label)`. Flag anything that can grow with
  cluster size, request volume, or user input.

### D. Stability level (`StabilityLevel`)
- New metrics default to and should be `ALPHA` unless there's an explicit, justified promotion.
  **BLOCKER if a brand-new metric is introduced at `BETA`/`STABLE`** — a metric's maturity is governed
  by the stability framework, *not* by the maturity of the feature it measures. A GA feature gate does
  **not** justify a BETA/STABLE metric.
- **Promotions must walk every stage** (ALPHA→BETA→STABLE), soaking ≥1 release at BETA — they can't
  skip straight to STABLE just because the feature went GA.
- **Promotion to `STABLE`/`BETA` requires explicit SIG-Instrumentation sign-off** — flag it prominently
  and confirm there's a graduation rationale.
- Guarantees to enforce in review:
  - `ALPHA` / `INTERNAL`: no guarantees, may change/delete anytime.
  - `BETA`: labels may be **added** but **not removed**; ~1 release / 4 months min lifetime.
  - `STABLE`: name, type, labels, buckets are **frozen** — no labels added or removed; only change
    allowed is marking deprecated. **BLOCKER** if a STABLE metric's signature changes.

### E. Stable metrics list must be regenerated
- Any change affecting a STABLE/BETA metric must be reflected in
  `hack/tools/instrumentation/testdata/stable-metrics-list.yaml`.
- The author must run `hack/update-generated-stable-metrics.sh`; CI runs
  `hack/verify-generated-stable-metrics.sh`. **BLOCKER** if the metric changed but the snapshot didn't
  (or vice-versa) — the diff should be internally consistent.

### F. Deprecation / removal policy
- You cannot just delete or rename a STABLE/BETA metric. Lifecycle is
  **Stable → Deprecated (`DeprecatedVersion: "1.x"`) → Hidden → Deleted**. Refer to the [Kubernetes metric deprecation policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-a-metric) for full details.
- Minimum bake times before hiding/removal:
  - STABLE: ≥ 3 releases or 9 months after deprecation.
  - BETA: ≥ 1 release or 4 months.
  - ALPHA: same release is fine.
- Confirm `DeprecatedVersion` matches the actual target release.

### G. Instrument type (gauge vs counter vs histogram)
- A value that only ever **accumulates / counts events** must be a **counter** (`_total`), not a gauge.
- A **duration/size distribution** must be a **histogram**, not a gauge — a gauge only shows the last
  value as a flat line and can't give you rate or percentiles.
- A point-in-time level (queue depth, in-flight) is a **gauge**. Flag mismatches as a BLOCKER.

### H. Histogram buckets
- Buckets should match the realistic value range and be powered/spaced sensibly (often
  `prometheus.ExponentialBuckets` / `DefBuckets`). Flag too-few buckets, buckets in the wrong unit,
  or buckets that don't cover the expected range.
- **Limit the number of buckets:** A classic histogram should have a reasonable number of buckets (typically ≤ 30-50 buckets). If a metric defines too many manual buckets, flag it as a **BLOCKER/QUESTION** to protect against memory explosion.
- **Bound the top bucket by a real limit.** If the measured operation has a known timeout/cap, use that
  variable as the max bucket rather than an arbitrary number.
- **If a unit changed (e.g. ms→s), the buckets MUST be re-scaled** (ms→s is ×1000). Leaving buckets
  unchanged collapses values into the left-most bucket and destroys resolution — treat as a BLOCKER and
  ask "do these buckets still make sense at the new unit?"

### I. Documentation
- New/changed documented metrics must update the metrics documentation list, e.g.
  `hack/tools/instrumentation/documentation/documentation-list.yaml`, via
  `hack/update-metrics-documentation-list.sh` (CI: `hack/verify-metrics-documentation-list.sh`).
- `Help` text must be present, accurate, and describe units. Diff `Help`/doc wording against the prior
  text — reviewers catch spelling/regressions here.

### J. Tests (quality, not coverage theater)
- **BLOCKER: reject trivial metric tests.** A test that calls `Inc()`/`Observe()` directly and asserts
  the value moved tests the Prometheus library, not the code. The metric must be exercised by the
  behavior test for the surrounding feature (ref "Adding a new metric" point #5 in
  `metric-instrumentation.md`). If existing behavior tests already cover it, the new metric-only test
  is redundant — drop it.
- Use `k8s.io/component-base/metrics/testutil` (`GatherAndCompare` / `CollectAndCompare`) to assert the
  **whole** rendered metric (name, labels, `# HELP`, `# TYPE`) rather than many `assert.Contains` calls.
- For verbose histograms, compare only `_count`/`_sum` and ignore individual buckets.
- Keep expected golden text **inline in the test**, not in separate fixture files.

### K. PR hygiene the approvers enforce
- **One concern per PR**: don't mix unrelated fixes (e.g. flaky-test fix) with a metrics change; squash
  unrelated commits. Flag and ask to split.
- **Description & release notes must match the actual change** (e.g. list exactly which metrics are
  promoted). Flag mismatches.
- **New metrics package** must add an `OWNERS` assigning SIG-Instrumentation, and be placed to avoid odd
  cross-component imports.
- Watch for **data races** introduced in hot-path label/metric construction.

### L. Metric Renaming (including Subsystem/Namespace changes)
- **BLOCKER** if a metric is renamed in-place (i.e. changing the metric name, subsystem, or namespace) without going through the deprecation cycle.
- Renaming a metric is a breaking change and requires:
  1. Deprecating the old metric signature (keeping it registered but marking it as deprecated with `DeprecatedVersion`).
  2. Adding the new metric signature as a separate metric.
  3. Maintaining both metrics during the required transition period (based on stability level guarantees) to avoid breaking user dashboards and alerts.


## Step 4 — Output

Produce a review in this format:

```
## SIG-Instrumentation Metrics Review — PR #<n>: <title>

**Scope:** <which metrics/files this touches>
**Verdict:** /lgtm | /approve | changes requested | needs SIG discussion

### Blockers
- [file:line] <issue> — <why it matters> → <fix>

### Nits
- [file:line] <suggestion>

### Questions for author
- <question>

### Required local commands before merge
- hack/verify-metrics-naming.sh
- hack/verify-generated-stable-metrics.sh   (run update-* first if it fails)
- hack/verify-metrics-documentation-list.sh
```

Be concrete, cite lines, and explain the *why* (stability guarantee, cardinality blast radius, prom
convention) — not just the rule. Default to `ALPHA` skepticism: question any promotion to STABLE/BETA.

## Reference
- `OWNERS_ALIASES` → `sig-instrumentation-approvers` / `-reviewers` define who can `/approve` and `/lgtm`.
- Metric stability framework & deprecation policy:
  `kubernetes/community` → `contributors/devel/sig-instrumentation/metric-stability.md`.
- Instrumentation guidance (incl. test guidance): `metric-instrumentation.md` in the same directory.
- Stability is implemented in `k8s.io/component-base/metrics`.

