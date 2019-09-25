# Kubernetes Release Versioning

Reference: [Semantic Versioning](http://semver.org)

Legend:

* **Kube X.Y.Z** refers to the version (git tag) of Kubernetes that is released.
This versions all components: apiserver, kubelet, kubectl, etc.  (**X** is the
major version, **Y** is the minor version, and **Z** is the patch version.)

## Release versioning

### Minor version scheme and timeline

* Kube X.Y.0-alpha.W, W > 0 (Branch: master)
  * Alpha releases are released roughly every two weeks directly from the master
branch.
  * No cherrypick releases. If there is a critical bugfix, a new release from
master can be created ahead of schedule.
* Kube X.Y.Z-beta.W (Branch: release-X.Y)
  * When master is feature-complete for Kube X.Y, we will cut the release-X.Y
branch 2 weeks prior to the desired X.Y.0 date and cherrypick only PRs essential
to X.Y.
  * This cut will be marked as X.Y.0-beta.0, and master will be revved to X.Y+1.0-alpha.0.
  * If we're not satisfied with X.Y.0-beta.0, we'll release other beta releases,
(X.Y.0-beta.W | W > 0) as necessary.
* Kube X.Y.0 (Branch: release-X.Y)
  * Final release, cut from the release-X.Y branch cut two weeks prior.
  * X.Y.1-beta.0 will be tagged at the same commit on the same branch.
  * X.Y.0 occur 3 to 4 months after X.(Y-1).0.
* Kube X.Y.Z, Z > 0 (Branch: release-X.Y)
  * [Patch releases](#patch-releases) are released as we cherrypick commits into
the release-X.Y branch, (which is at X.Y.Z-beta.W,) as needed.
  * X.Y.Z is cut straight from the release-X.Y branch, and X.Y.Z+1-beta.0 is
tagged on the followup commit that updates pkg/version/base.go with the beta
version.
* Kube X.Y.Z, Z > 0 (Branch: release-X.Y.Z)
  * These are special and different in that the X.Y.Z tag is branched to isolate
the emergency/critical fix from all other changes that have landed on the
release branch since the previous tag
  * Cut release-X.Y.Z branch to hold the isolated patch release
  * Tag release-X.Y.Z branch + fixes with X.Y.(Z+1)
  * Branched [patch releases](#patch-releases) are rarely needed but used for
emergency/critical fixes to the latest release
  * See [#19849](https://issues.k8s.io/19849) tracking the work that is needed
for this kind of release to be possible.

### Major version timeline

There is no mandated timeline for major versions and there are currently no criteria
for shipping 2.0.0. We haven't so far applied a rigorous interpretation of semantic 
versioning with respect to incompatible changes of any kind (e.g., component flag changes).
We previously discussed releasing 2.0.0 when removing the monolithic `v1` API
group/version, but there are no current plans to do so.

### CI and dev version scheme

* Continuous integration versions also exist, and are versioned off of alpha and
beta releases. X.Y.Z-alpha.W.C+aaaa is C commits after X.Y.Z-alpha.W, with an
additional +aaaa build suffix added; X.Y.Z-beta.W.C+bbbb is C commits after
X.Y.Z-beta.W, with an additional +bbbb build suffix added. Furthermore, builds
that are built off of a dirty build tree, (during development, with things in
the tree that are not checked it,) it will be appended with -dirty.

### Supported releases and component skew

We expect users to stay reasonably up-to-date with the versions of Kubernetes
they use in production, but understand that it may take time to upgrade,
especially for production-critical components.

We expect users to be running approximately the latest patch release of a given
minor release; we often include critical bug fixes in
[patch releases](#patch-releases), and so encourage users to upgrade as soon as
possible.

Different components are expected to be compatible across different amounts of
skew, all relative to the master version.  Nodes may lag masters components by
up to two minor versions but should be at a version no newer than the master; a
client should be skewed no more than one minor version from the master, but may
lead the master by up to one minor version.  For example, a v1.3 master should
work with v1.1, v1.2, and v1.3 nodes, and should work with v1.2, v1.3, and v1.4
clients.

Furthermore, we expect to "support" three minor releases at a time.  "Support"
means we expect users to be running that version in production, though we may
not port fixes back before the latest minor version. For example, when v1.3
comes out, v1.0 will no longer be supported: basically, that means that the
reasonable response to the question "my v1.0 cluster isn't working," is, "you
should probably upgrade it, (and probably should have some time ago)". With
minor releases happening approximately every three months, that means a minor
release is supported for approximately nine months.

## Patch releases

Patch releases are intended for critical bug fixes to the latest minor version,
such as addressing security vulnerabilities, fixes to problems affecting a large
number of users, severe problems with no workaround, and blockers for products
based on Kubernetes.

They should not contain miscellaneous feature additions or improvements, and
especially no incompatibilities should be introduced between patch versions of
the same minor version (or even major version).

Dependencies, such as Docker or Etcd, should also not be changed unless
absolutely necessary, and also just to fix critical bugs (so, at most patch
version changes, not new major nor minor versions).

## Upgrades

* Users can upgrade from any Kube 1.x release to any other Kube 1.x release as a
rolling upgrade across their cluster. (Rolling upgrade means being able to
upgrade the master first, then one node at a time. See [#4855](https://issues.k8s.io/4855) for details.)
  * However, we do not recommend upgrading more than two minor releases at a
time (see [Supported releases and component skew](#Supported-releases-and-component-skew)), and do not recommend
running non-latest patch releases of a given minor release.
* No hard breaking changes over version boundaries.
  * For example, if a user is at Kube 1.x, we may require them to upgrade to
Kube 1.x+y before upgrading to Kube 2.x. In others words, an upgrade across
major versions (e.g. Kube 1.x to Kube 2.x) should effectively be a no-op and as
graceful as an upgrade from Kube 1.x to Kube 1.x+1. But you can require someone
to go from 1.x to 1.x+y before they go to 2.x.

There is a separate question of how to track the capabilities of a kubelet to
facilitate rolling upgrades. That is not addressed here.

