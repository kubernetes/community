# Policy: Adopting New Go Features in Kubernetes

## Background

When Kubernetes updates to a new Go version, new language features, standard
library additions, linter checks, and tooling capabilities become available. 
Encouraging or requiring use of them immediately on the main branch can create friction for contributors 
backporting changes to older release branches still on an earlier Go version.

## Policy

**Do not actively encourage or require use of a new Go feature until at least
one active older release branch has adopted the Go version that introduced it.**

Until the threshold is met, open tracking issues to revisit once the condition
is satisfied.

## Example: `ptr.To` → `new` (Go 1.26)

Go 1.26 allows `new(x)` in place of `ptr.To(x)`. The `modernize` linter
surfaces this as a hint (`newexpr`). We should hint at changes like this in
master only when at least one older stable branch is also on Go 1.26.