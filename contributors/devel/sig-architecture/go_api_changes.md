# Definition of Go API

The term "Kubernetes API" usually refers to the externally visible behavior of
components in a Kubernetes release. For a full definition what this includes see
the [API review process](./../../sig-architecture/api-review-process.md#what-apis-need-to-be-reviewed).
The [deprecation policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/)
explains under which circumstances and how it is possible to break that API.

The Go API of source code packages *is not* part of that Kubernetes API:
- It does not fall under the API review process.
- Strong stability guarantees only apply to modules with a version >= 1,
  as usual for Go package.
  Staging repositories are intentionally published with
  `0.<k/k minor version>.<k/k patch version>` as version because
  their Go API may change from one release to the next.

Having said that, the Kubernetes project nonetheless strives to minimize the
impact on downstream consumers and the wider ecosystem when changing Go source
code. This document describes in more details how the Kubernetes Go APIs are
maintained.

# Stability goals

## Staging repositories

All staging repositories get published separately and may be used by other
projects. Changing their Go API must consider the impact on the ecosystem. In
particular, `k8s.io/client-go` and `k8s.io/apimachinery` are so widely used
that breaking Go API changes must be avoided as much as practical.

## k8s.io/kubernetes

Importing `k8s.io/kubernetes` is discouraged. Projects doing it cannot expect
any kind of Go API stability.

## Other

Some modules like `k8s.io/klog/v2` have committed to maintain a certain Go API
without breaking changes. Others can still be changed or might not even have
tagged releases (for example, `k8s.io/utils`). Nonetheless the same caution as for
staging repositories applies.

# Pull requests

The "bug", "deprecation" and "api-change" kind labels refer to the user-visible
Kubernetes behavior. They are not applicable to pull requests which only change Go
APIs unless those changes also have those user-visible effect. Depending on the scope,
"cleanup" and "feature" can be suitable kind labels for source code changes.

The Kubernetes release notes are targeting users and administrators, not Go
developers consuming Kubernetes packages, therefore a release note is only
required for those user-visible effects. "Action required" in a release note
*must not* be used for Go API changes.

# Changes

## Breaking changes

Breaking Go API changes can be hard to spot. Several repositories have been
configured to use the [apidiff tool](https://pkg.go.dev/golang.org/x/exp/apidiff#section-readme).

In kubernetes/kubernetes, "pull-kubernetes-apidiff-client-go" gets triggered
automatically for changes to client-go and then only checks that.
"pull-kubernetes-apidiff" can be invoked manually. It covers the entire
repository and includes trial builds of downstream consumers (currently
controller-runtime) with the modified code. These jobs fail if breaking
changes are found. But because those changes are allowed, such a failure
does not block merging the pull request. Instead, developers and reviewers
need to check the result manually:

```
## ./staging/src/k8s.io/client-go
Incompatible changes:
- ./kubernetes.(*Clientset).Discovery: changed from func() k8s.io/client-go/discovery.DiscoveryInterface to func() k8s.io/client-go/discovery.DiscoveryInterfaces
- ./kubernetes.Interface.Discovery: changed from func() k8s.io/client-go/discovery.DiscoveryInterface to func() k8s.io/client-go/discovery.DiscoveryInterfaces
- ./kubernetes/fake.(*Clientset).Discovery: changed from func() k8s.io/client-go/discovery.DiscoveryInterface to func() k8s.io/client-go/discovery.DiscoveryInterfaces
Compatible changes:
- ./discovery.(*DiscoveryClient).GroupsAndMaybeResourcesWithContext: added
- ./discovery.(*DiscoveryClient).OpenAPISchemaWithContext: added
...
```

The goal is to add only "compatible" changes and to avoid "incompatible"
changes. For incompatible changes, consider the impact on the ecosystem.
Breaking changes in widely consumed APIs need a very strong justification.

It's often better to avoid them, even if that makes the source code more complicated:
- Add new methods instead of changing an existing one. The old one often can
  call the new one to avoid code duplication. `//go:fix inline` may be useful
  in such a case to support automated replacement of the old call with the new
  one. But be sure that such a mechanical replacement really leads to the desired
  result. For example, a function `Foo(..) { FooWithContext(context.Background(), ... }`
  should *not* use it because callers of `FooWithContext` should provide
  a real context.
- Instead of adding methods to an existing interface, create a new one.  In
  code receiving the old interface, use type assertions to check for the new
  methods and only use them if implemented. See
  https://github.com/kubernetes/kubernetes/pull/129109 for an example using
  this approach.

## Interfaces

When defining an interface as part of the Go API of a package, consider the
pros and cons:
- Pro: when an API accepts an interface instead of a concrete type,
  consumers can provide their own implementation.
- Con: any change of the interface, including adding more methods, breaks
  consumers which implement it (the reason why this gets called out here).
- Con: documentation in `go doc` or pkg.go.dev for an interface is
  less useful (fewer options for formatting, cannot be referenced);
  comments attached to an unexported implementation of the interface are
  not visible to consumers.
- Con: some compiler optimizations or linter checks may not be possible.

If the advantage of an interface is not needed, then it might be better to
return a concrete type which only exports fields and methods that consumers of
the package are meant to use. An interface can still be introduced later.

When an interface is needed because there are multiple implementations inside
the package, but consumers are not expected to implement it, then a `private()`
method can be included in the interface, for example as in
[testing.TB](https://cs.opensource.google/go/go/+/refs/tags/go1.26.3:src/testing/testing.go;l=914-917;drc=2b8dbb35b0d6a5601ae9b6f1d1de106774251214).
Then additions to the interface are not breaking changes because there
cannot be any implementation outside of the package which would be
broken.

## Deprecations

A formal `// Deprecated: use <something> instead` comment in the last line of a
doc comment gets picked up by Go tools (IDEs, linters). Such a comment can mean
"you must take action" (hard deprecation), for example because some exported
API will be removed in the future, or it can mean "you might want to use
something else because it's better, but it's okay to keep using it" (soft
deprecation). Tools cannot tell the difference, therefore downstream consumers
may lean towards playing it safe with a "no usage of deprecated code" policy.

For hard deprecations, the formal comment is appropriate. How long to wait
before taking next steps depends on how long it takes for the ecosystem to
adapt. Use a tracking issue to make sure that future actions are not forgotten
and to provide further information about the deprecation plan.

For soft deprecations, consider whether the impact on downstream consumers
is really worth it: assume that consumers will treat it like a hard deprecation
and eventually adapt their code to get rid of linter warnings. The litmus test
is whether we would accept all PRs which remove usage of the deprecated API. If
the answer is "no", then a free-form comment with information on what to use
instead is better.

## Documenting changes

In key Go modules, breaking Go API changes must be documented in a
CHANGELOG.md file. This documentation must be part
of the PR which makes the change. This serves two purposes:
- Make the API change more visible during code review.
- Inform consumers of the module, including guidance on how to deal
  with the change.

This will be enforced by a `hack/verify-go-apidocs.sh` check in
https://github.com/kubernetes/kubernetes/pull/138351.
