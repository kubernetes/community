---
title: "Platforms Guide"
weight: 5
description: |
  Outlines the necessary steps to either add or remove supported platform builds
  in Kubernetes.
---

## Adding supported platforms

The default Kubernetes platform is `linux/amd64`. This platform has always been
fully tested, and the build and release systems initially supported only this
platform. SIG Release started an [effort to support multiple architectures][0].
As part of this effort, they added support in our build and release pipelines
for the architectures `arm`, `arm64`, `ppc64le` and `s390x` on different
operating systems like Linux, Windows and macOS.

[0]: https://github.com/kubernetes/kubernetes/issues/38067

The main focus was to have binaries and container images to be available for
these architectures/operating systems. Contributors should be able to to take
these artifacts and set up CI jobs to adequately test these platforms.
Specifically to call out the ability to run conformance tests on these
platforms.

Target of this document is to provide a starting point for adding new platforms
to Kubernetes from a SIG Architecture and SIG Release perspective. This does not
include release mechanics or supportability in terms of functionality.

### Step 0: Engage with the community

The first step is to express the interest in adding a new platform or promoting
one into a different Tier. This can be done by opening a GitHub issue in the
[SIG Release repository](https://github.com/kubernetes/sig-release/issues),
attending the weekly SIG Release meetings or writing a message to the [SIG
Release mailing list](https://groups.google.com/g/kubernetes-sig-release).

The ultimate decision to approve or reject a platform will be done in the
community as part of the standard lazy consensus. Even if all mentioned
requirements for a platform are being met, it's possible that external
dependencies, infrastructure costs or anything else have influence on the
maintainability of the platform.

### Step 1: Building

The container image based build infrastructure should support this architecture.
This implicitly requires the following:

- golang should support the platform
- All dependencies, whether vendored or run separately, should support this
  platform

In other words, anyone in the community should be able to use the SIG Release
provided build tooling to generate all artifacts required to build Kubernetes.

More information about how to build Kubernetes can be found in [the build
documentation][1].

[1]: https://github.com/kubernetes/kubernetes/tree/3f7c09e/build#building-kubernetes

### Step 2: Testing

It is not enough for builds to work as it gets bit-rotted quickly when vendoring
in in new changes, update versions of things to be used etc. This means the
project need a good set of tests that exercise a wide battery of jobs in this
new architecture.

A good starting point from a testing perspective are:

- unit tests
- e2e tests
- node e2e tests

This will ensure that community members can rely on these architectures on a
consistent basis, which will give folks who are making changes a signal when
they break things in a specific architecture.

This implies a set of folks who stand up and maintain both post-submit and
periodic tests, watch them closely and raise the flag when things break. They
will also have to help debug and fix any platform specific issues as well.

Creating custom [testgrid][4] dashboards can help to monitor platform specific
tests.

[4]: https://testgrid.k8s.io

### Step 3: Releasing

The first 2 steps provide a reasonable expectation that there are people taking
care of a supported platform and it works in a reproducible environment.

Getting to the next level is another big jump, because it needs to be ensured
that real users can rely on the shipped artifacts.

This means specifically to add a set of CI jobs to the release-informing and
release-blocking tabs of testgrid. The Kubernetes release team has a "CI signal"
group that relies on the status(es) of these jobs to either ship or hold a
release. Essentially, if things are mostly red with occasional green, it would
be prudent to not even bother making this architecture as part of the release.
CI jobs get added to release-informing first and when these get to a point where
they work really well, then they get promoted to release-blocking.

The main problem here is once the project starts to ship something, users will
start to rely on it. While it's straight forward to setup a CI job as a one time
thing, it's a totally one to consistently maintain them over time. What SIG
Release is looking for is a strong green CI signal for release managers to cut a
release and for folks to be able to report problems and them getting addressed.
This also includes [conformance testing][2] to ensure that the supported
platform behaves as intended. This can be done by working with SIG Architecture
as part of the [conformance sub project][3] in addition to testing and release.

[2]: https://github.com/cncf/k8s-conformance
[3]: https://github.com/kubernetes/community/tree/master/sig-architecture#conformance-definition

### Step 4: Finishing

If you got this far, you really have made it! You have a clear engagement with
the community, you are working seamlessly with all the relevant SIGs, you have
your content in the Kubernetes release and get end users to adopt your
architecture. Having achieved conformance, you will gain conditional use of the
Kubernetes trademark relative to your offerings.

## Deprecating and removing supported platforms

Supported platforms may be considered as deprecated for various reasons, for
example if they are being replaced by new ones, are not actively used or
maintained any more. Deprecating an already supported platform has to follow a
couple of steps:

1. The platform deprecation has been announced on the [Kubernetes Announce
   mailing list](https://groups.google.com/g/kubernetes-announce)
   and links to an Kubernetes GitHub issue for further discussions and consensus.

1. The deprecation will be active immediately after consensus has been reached
   at a set deadline. This incorporates approval from SIG Release and
   Architecture.

1. Removing the supported platform will be done in the beginning of the next
   minor (v1.N+1.0) release cycle, which means to:
   - Update the Kubernetes build scripts to exclude the platform from all targets
   - Update the [kubernetes/sig-release](https://github.com/kubernetes/sig-release)
     repository to reflect the current set of supported platforms.

Please note that actively supported release branches are not affected by the
removal. This ensures compatibility with existing artifact consumers on a best
effort basis.
