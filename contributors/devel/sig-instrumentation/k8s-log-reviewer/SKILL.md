---
name: k8s-log-reviewer
description: Review a kubernetes/kubernetes pull request that changes logging code, applying SIG-Instrumentation approver standards for structured logging (klog InfoS/ErrorS), contextual logging (logr via context, klog.FromContext, WithLogger variants), logcheck enforcement, key/value and verbosity conventions, and ktesting. Use when asked to review a k8s PR/diff that touches logging, klog, structured/contextual logging, or logcheck.conf. Trigger phrases: "review this logging PR", "sig-instrumentation log review", "is this structured/contextual logging correct".
---

# SIG-Instrumentation Logging Review

Review a `kubernetes/kubernetes` PR the way a **SIG-Instrumentation approver** would for **logging**
changes. 

This skill reviews **only** logging-relevant changes. If the PR doesn't touch logging, say so and stop.
(For metrics, use `sig-instrumentation-metrics-review` instead.)

## Step 1 — Get the change
- A PR number/URL → `gh pr view <n> --repo kubernetes/kubernetes` + `gh pr diff <n> --repo kubernetes/kubernetes`.
- A local branch → `git diff` against the merge base. A pasted diff → use directly.
- No `gh`/auth → GitHub API (`/repos/kubernetes/kubernetes/pulls/<n>/files`) or ask for the diff.

## Step 2 — Scope to logging
In-scope signals (any of):
- Adds/changes `klog.*` calls (`InfoS`, `ErrorS`, `Info`, `Infof`, `V(n)`, `KObj`, `KRef`,
  `FromContext`, `Background`, `TODO`), or `logr.Logger` usage.
- Adds a `logger`/`ctx` parameter or a `…WithLogger` variant for contextual logging.
- Edits `hack/logcheck.conf`, `hack/tools` logcheck/logtools versions, or golangci-lint logging linters.
- Touches `k8s.io/component-base/logs` (incl. JSON logging) or `ktesting`.

If none apply, report "not a logging change — out of scope" and stop.

## Step 3 — Review checklist
For each finding cite file:line and tag **BLOCKER** / **NIT** / **QUESTION**.

### A. Structured form
- **BLOCKER**: new unstructured `klog.Info/Infof/Errorf/Warningf` in a migrated package — use
  `InfoS`/`ErrorS`. `ErrorS(err, msg, kv...)` takes the error **first**. Use it for log output that needs admin attention,  otherwise prefer `InfoS`. `ErrorS` may be called with nil error. `InfoS` may be called with `"err": err` to log an error.
  genuinely no error).
- Message must be a **constant string** (no `fmt.Sprintf`), capitalized, no trailing punctuation/newline.
- Keys are **lowerCamelCase**, stable, and reused (don't invent a new key for an existing concept).
- Object values use `klog.KObj(obj)` / `klog.KRef(ns, name)`, not `%s`/`%v` formatting.

### B. Contextual logging (KEP-3077)
- Obtain the logger from context: `logger := klog.FromContext(ctx)`; libraries receive a logger via
  ctx rather than calling `klog.Background()`.
- New public functions that log should accept `ctx`/`logger`; prefer extending an unreleased signature
  over adding a `…WithLogger` twin. Carry the convention comment.
- `WithValues` for repeated key/value pairs; `WithName` to scope a component's logger.
- **QUESTION/BLOCKER** on `context.Background()`/`context.TODO()` in new code without a justifying
  comment + tracking issue.

### C. logcheck / tooling
- If the PR migrates a package, it must **add that package to `hack/logcheck.conf`** so enforcement
  sticks; conversely, don't enable a directory whose prerequisites aren't met.
- logcheck runs via golangci-lint; logtools/logcheck version bumps live in `hack/tools`. Confirm CI
  (`pull-kubernetes-verify` / golangci-lint) covers the change.

### D. Verbosity
- Right `V()` level: `V(0)` important/always, `V(2)` useful steady-state, `V(4)` debug, `V(5)` trace.
  Flag high-cost values logged at low verbosity, and secrets/PII logged at any level.

### E. Events
- Contextual event migration is **all-or-nothing** per area; grep for stragglers. Plugins should be
  able to attach key/values via `EventRecorderLogger`.

### F. Tests
- Log-output assertions use `ktesting` with per-test isolation, not the global logger. Behavior tests
  should still exercise the real code path.

## Step 4 — Output
```
## SIG-Instrumentation Logging Review — PR #<n>: <title>

**Scope:** <which logging/files this touches>
**Verdict:** /lgtm | /approve | changes requested | needs SIG discussion

### Blockers
- [file:line] <issue> — <why> → <fix>

### Nits
- [file:line] <suggestion>

### Questions for author
- <question>

### Required before merge
- make verify WHAT=verify-golangci-lint   (logcheck runs here)
- ensure migrated packages are added to hack/logcheck.conf
```
Be concrete, cite lines, explain the *why* (JSON searchability, contextual-logging correctness,
cancellation propagation) — not just the rule.

## Reference
- `OWNERS_ALIASES` → `sig-instrumentation-approvers`.
- Structured logging: `kubernetes/community` → `contributors/devel/sig-instrumentation/migration-to-structured-logging.md`
  (KEP-1602); contextual logging: `…/contextual-logging.md` (KEP-3077).
- logcheck lives in `sigs.k8s.io/logtools`; config at `hack/logcheck.conf`.
- klog API & `ktesting`: `k8s.io/klog/v2`.


