# Decoupling kubectl from core k8s

Status: Pending

Version: Alpha

Implementation Owner: jregan@google.com

## Motivation

The [kubectl program] is a command-line client used to configure and control k8s
(kubernetes) clusters.  The program lives in the core [k8s repo].

kubectl depends on various parts of k8s and k8s depends on various parts of
kubectl with varying degrees of intentional design.  If a bug is detected in how
kubectl interacts with k8s, there's no procedure to release a fix that doesn't
involve a full release of k8s.

This document describes how a clear line will be drawn in the code between
kubectl and the remainder of k8s to leave only intentional dependencies.

This is a precondition to being able to release kubectl as a product distinct
from the rest of k8s, or more dramatic changes
like
[moving kubectl out of the kubernetes repository][proposal for kubectl repo] and
into [its own repository][kubectl repo].

## Proposal

1. Identify, via `BUILD` file rules, existing good and bad dependencies.

   _Good dependencies are on public libraries intended for the given usage; bad
   dependencies are otherwise._

   At the same time, inhibit introduction of _new_ bad dependencies.

1. Create. delete, move or merge packages to _remove_ bad dependencies.

   At the same time, adapt CICD tools, miscellaneous shell scripts, etc. to
   continue to function properly.


![insert amazing before and after diagram here][reposCandide]

### Non-goals

#### Extensive changes to testing.

Decoupling kubectl is a necessary but insufficient condition for releasing
kubectl at a different cadence from k8s.

Another necessary condition is an end to end (e2e) test framework to test an
instance of kubectl (previously released or HEAD) against an instance of k8s
(likewise, previously released or HEAD).  This may require its own proposal if
existing framework cannot handle it.

#### Immediately changing policy to release kubectl and k8s at a different cadence.

This can wait.  For purposes of bootstrapping a new style of e2e testing and
release, its convenient to already have a _decoupled_ kubectl in the wild,
released per the _existing_ process.  Any new process will need a rollback state
it understands.

#### Moving kubectl to its [own repo][kubectl repo], outside the [k8s repo].

This decoupling work is a precondition to such a move, but doesn't _require_ a
new repo.

## User Experience

The user should see no change other than, it's hoped, faster feature development
and bug fixes accruing from the kubectl/k8s decoupling.

## Implementation

The identification of bad dependencies, and the eventual prevention of their
re-introduction will be accomplished via _visibility_ rules in `BUILD` files.
Violations in visibility cause `make bazel-build` to fail, which in turn causes
the submit queue to fail - so attempting to introduce a bad dependency is like a
failing unit tests.

### Background Facts

 * A _package_ is any directory that contains a BUILD file.
 
 * A `package_group` is a BUILD file term that defines a set of
   packages for use in other rules.
 
 * A _visibility rule_ takes a list of package groups as its
   argument - or one of the pre-defined groups `//visibility:private` or
   `//visibility:public`.
 
 * If no visibility is explicitly defined, a package is _private_ by default.
 
 * kubernetes has explicitly labeled its nearly 2000 packages as public - a
   _sensible thing to do_ when first introducing bazel.
   
   Now that kubernetes builds under bazel, its possible to use bazel to
   discover, fix and prevent bad deps.


### Step One: Introduce visibility rules and clarify dependence on kubectl.

When this work in this section is complete, all dependencies on kubectl code by
non-kubectl code will either be removed or specifically allowed.

 * A new k8s package will be created called `kubernetes/visible_to`.

   It will have only two files: `BUILD` and `OWNERS`.

 * The only rules in `BUILD` will define package groups.

   These groups will exist only to be used in visibility rules in
   other packages throughout k8s.

 * All kubectl code will switch to private by default.

   This means removal of all package-level policies like
   ```
   package(default_visibility = ["//visibility:public"])
   ```

 * Exceptions to this will be per-target.

   A target will let others depend on it like this:
   ```
   visibility = ["//visible_to:kubectl_client"]
   ```
   Once all these changes are made, all dependencies are
   intentional, albeit both good and bad.
 
 * `OWNERS` file imposes special approval for visibility changes.

   Review will require familiarity with visibility usage, since
   small changes can open large holes.  This is intended to stop
   introduction of new bad depenencies.

 * Bad dependencies labelled for removal.

   Bad deps will be allowed and identified with a special package
   group naming convention to call them out for removal.

 * Code changed to remove bad deps.

   
   
### Step Two: Clarify dependence on k8s.

When this work in this section is complete, any dependence kubectl has
on k8s will be specifically allowed.

_This requires removing all public exposure from the approximately 2000
package in k8s because_ __a package cannot be both public and declare its
visibility to a specific package.__ This task will be done via a one-off 
script.

To avoid turning this into the much more difficult task of identifying good and
bad dependencies throughout kubernetes, _the focus will remain on kubectl_.

Specifically, every package in ks8 will define a visibility rule that allows
_one or both_ of the following package groups to see it:

 * `kubectl` (i.e. a package group containing relevant parts of kubectl)

 * `everything_else` (i.e. all off k8s, minus kubectl)

A package that kubectl should have no interest in will exclude kubectl from its
visibility rules.

A package intended for public consumption will now make itself visible to both
`kubectl` and `everything_else`.

As other projects (e.g. kubelet, scheduler, kubeadm, ...) seek more isolation,
they can remove themselves in the same way from `everything_else`.


#### k8s code will have three destinations

The goal is to have kubectl so thoroughly decoupled from k8s that it can move to
its own repo and get code from kubernetes only via libraries k8s supports
specifically for its clients (e.g. the generated API code).

This means that anything [kubectl depends on] from non-kubectl paths _must move to
one of three places:_

1. Move to a kubectl path.

   I.e. the code becomes _kubectl code_ rather than _k8s code_.  This would be a
   reasonable place to move code used by kubectl, kubeadm and other clients.  If
   kubectl became its own repo, it could vendor code to kubeadm.
   
   Example:
   ```
   k8s.io/kubernetes/pkg/printers
   ```

2. Move to the [client-go repo].

   This is where kubectl should find all code for apis groups like
   authorization, autoscaling etc., and all code that faciliates access to these
   apis.

   Where kubectl now imports lines like

   ```
   k8s.io/kubernetes/pkg/apis/authorization
   k8s.io/kubernetes/pkg/apis/autoscaling
   ...
   k8s.io/kubernetes/pkg/util/crlf
   k8s.io/kubernetes/pkg/util/exec
   ```

   it will instead import from _client-go_.  The client-go repo itself will
   likely morph into distinct vendoring paths, or even distict repos (see
   comments in this discussion of the [federation API]).
   
3. Move logic out of kubectl and into the api-server.
   
   This is likely the fate of kubectl's current dependence on
   ```
   k8s.io/kubernetes/pkg/controller
   k8s.io/kubernetes/pkg/controller/deployment/util
   ```

## Client/Server Backwards/Forwards compatibility

These changes have no effect on backward/forward compatibility.

## Alternatives considered

None.  Decoupling is good, and a precondition to other good things
like shipping kubectl on its own without shipping all of k8s.

[k8s repo]: https://github.com/kubernetes/kubernetes
[original proposal]: https://docs.google.com/document/d/1i1vISyLhVcc6skMEgu8idxeL4LrieCCEO77ImZQwNjA/edit#
[kubectl program]: https://kubernetes.io/docs/user-guide/kubectl-overview/
[reposBefore]: reposBefore.svg
[reposAfter]: reposAfter.svg
[reposCandide]: reposCandide.svg
[kubectl repo]: https://github.com/kubernetes/kubectl
[client-go repo]: https://github.com/kubernetes/client-go
[federation API]: https://github.com/kubernetes/kubernetes/issues/41302
[proposal for kubectl repo]: https://docs.google.com/document/d/1ELBwwBstbNEK8906Qz5G-A9oqJWluFZIIWModYTHe6w/edit#heading=h.1nt8emg22nd9
[kubectl depends on]: https://docs.google.com/document/d/1ZtbE6xZwSbz2sqQVW-chST2atLSZBbpXWxwauPjMfUk/edit
