# Shared PID Namespace

Pods share namespaces where possible, but a requirement for sharing the PID
namespace has not been defined due to lack of support in Docker. Docker began
supporting a shared PID namespace in 1.12, and other Kubernetes runtimes (rkt,
cri-o, hyper) have already implemented a shared PID namespace.

This proposal defines a shared PID namespace as a requirement of the Container
Runtime Interface and links its rollout in Docker to that of the CRI.

## Motivation

Sharing a PID namespace between containers in a pod is discussed in
[#1615](https://issues.k8s.io/1615), and enables:

  1. signaling between containers, which is useful for side cars (e.g. for
     signaling a daemon process after rotating logs).
  2. easier troubleshooting of pods.
  3. addressing [Docker's zombie problem][1] by reaping orphaned zombies in the
     infra container.

## Goals and Non-Goals

Goals include:
  - Changing default behavior in the Docker runtime as implemented by the CRI
  - Making Docker behavior compatible with the other Kubernetes runtimes

Non-goals include:
  - Creating an init solution that works for all runtimes
  - Supporting isolated PID namespace indefinitely

## Modification to the Docker Runtime

We will modify the Docker implementation of the CRI to use a shared PID
namespace when running with a version of Docker >= 1.12. The legacy
`dockertools` implementation will not be changed.

Linking this change to the CRI means that Kubernetes users who care to test such
changes can test the combined changes at once. Users who do not care to test
such changes will be insulated by Kubernetes not recommending Docker >= 1.12
until after switching to the CRI.

Other changes that must be made to support this change:

1. Add a test to verify all containers restart if the infra container
   responsible for the PodSandbox dies. (Note: With Docker 1.12 if the source
   of the PID namespace dies all containers sharing that namespace are killed
   as well.)
2. Modify the Infra container used by the Docker runtime to reap orphaned
   zombies ([#36853](https://pr.k8s.io/36853)).

## Rollout Plan

SIG Node is planning to switch to the CRI as a default in 1.6, at which point
users with Docker >= 1.12 will receive a shared PID namespace by default.
Cluster administrators will be able to disable this behavior by providing a flag
to the kubelet which will cause the dockershim to revert to previous behavior.

The ability to disable shared PID namespaces is intended as a way to roll back
to prior behavior in the event of unforeseen problems. It won't be possible to
configure the behavior per-pod. We believe this is acceptable because:

* We have not identified a concrete use case requiring isolated PID namespaces.
* Making PID namespace configurable requires changing the CRI, which we would
  like to avoid since there are no use cases.

In a future release, SIG Node will recommend docker >= 1.12. Unless a compelling
use case for isolated PID namespaces is discovered, we will remove the ability
to disable the shared PID namespace in the subsequent release.


[1]: https://blog.phusion.nl/2015/01/20/docker-and-the-pid-1-zombie-reaping-problem/
