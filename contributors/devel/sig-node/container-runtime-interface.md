# CRI: the Container Runtime Interface

## What is CRI?

CRI (_Container Runtime Interface_) consists of a
specifications/requirements (to-be-added),
[protobuf API](https://git.k8s.io/kubernetes/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1/api.proto),
and [libraries](https://git.k8s.io/kubernetes/pkg/kubelet/cri/streaming)
for container runtimes to integrate with kubelet on a node. The CRI API is
currently in Alpha, and the CRI-Docker integration is used by default as of
Kubernetes 1.7+.

In the future, we plan to add more developer tools such as the CRI validation
tests.

## Why develop CRI?

Prior to the existence of CRI, container runtimes (e.g., `docker`, `rkt`) were
integrated with kubelet through implementing an internal, high-level interface
in kubelet. The entrance barrier for runtimes was high because the integration
required understanding the internals of kubelet and contributing to the main
Kubernetes repository. More importantly, this would not scale because every new
addition incurs a significant maintenance overhead in the main Kubernetes
repository.

Kubernetes aims to be extensible. CRI is one small, yet important step to enable
pluggable container runtimes and build a healthier ecosystem.

## How to use CRI?

See the [CRI installation documentation](https://kubernetes.io/docs/setup/cri/).

CRI is still young and we are actively incorporating feedback from developers
to improve the API. Although we strive to maintain backward compatibility,
developers should expect occasional API breaking changes.

*For Kubernetes 1.5, additional flags are required:*
 - Set apiserver flag `--feature-gates=StreamingProxyRedirects=true`.
 - Set kubelet flag `--experimental-cri=true`.

## Does Kubelet use CRI today?

Yes, Kubelet always uses CRI except for using the rktnetes integration.

The old, pre-CRI Docker integration was removed in 1.7.

## Specifications, design documents and proposals

The Kubernetes 1.5 [blog post on CRI](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/)
serves as a general introduction.

Below is a mixed list of CRI specifications/requirements, design docs and
proposals. We are working on adding more documentation for the API.

  - [Original proposal](https://github.com/kubernetes/kubernetes/blob/release-1.5/docs/proposals/container-runtime-interface-v1.md)
  - [Networking](kubelet-cri-networking.md)
  - [Container metrics](cri-container-stats.md)
  - [Exec/attach/port-forward streaming requests](https://docs.google.com/document/d/1OE_QoInPlVCK9rMAx9aybRmgFiVjHpJCHI9LrfdNM_s/edit?usp=sharing)
  - [Container stdout/stderr logs](https://github.com/kubernetes/kubernetes/blob/release-1.5/docs/proposals/kubelet-cri-logging.md)

## CRI runtimes

 - [cri-o](https://github.com/cri-o/cri-o)
 - [rktlet](https://github.com/kubernetes-retired/rktlet)
 - [frakti](https://github.com/kubernetes/frakti)
 - [cri-containerd](https://github.com/containerd/cri)
 - [singularity-cri](https://github.com/sylabs/singularity-cri)

## [Status update](#status-update)
### Kubernetes v1.7 release (Docker-CRI integration GA, container metrics API)
  - The Docker CRI integration has been promoted to GA.
  - The legacy, non-CRI Docker integration has been completely removed from
    Kubelet. The deprecated `--enable-cri` flag has been removed.
  - CRI has been extended to support collecting container metrics from the
    runtime.

### Kubernetes v1.6 release (Docker-CRI integration Beta)
 **The Docker CRI integration has been promoted to Beta, and been enabled by
default in Kubelet**.
  - **Upgrade**: It is recommended to drain your node before upgrading the
    Kubelet. If you choose to perform in-place upgrade, the Kubelet will
    restart all Kubernetes-managed containers on the node.
  - **Resource usage and performance**: There is no performance regression
    in our measurement. The memory usage of Kubelet increases slightly
    (~0.27MB per pod) due to the additional gRPC serialization for CRI.
  - **Disable**: To disable the Docker CRI integration and fall back to the
    old implementation, set `--enable-cri=false`. Note that the old
    implementation has been *deprecated* and is scheduled to be removed in
    the next release. You are encouraged to migrate to CRI as early as
    possible.
  - **Others**: The Docker container naming/labeling scheme has changed
    significantly in 1.6. This is perceived as implementation detail and
    should not be relied upon by any external tools or scripts.

### Kubernetes v1.5 release (CRI v1alpha1)

  - [v1alpha1 version](https://github.com/kubernetes/kubernetes/blob/release-1.5/pkg/kubelet/api/v1alpha1/runtime/api.proto) of CRI is released.

#### [CRI known issues](#cri-1.5-known-issues):

Note, these are known issues as of the 1.5 release. They may or may not have
been fixed since.

  - [#27097](https://github.com/kubernetes/kubernetes/issues/27097): Container
    metrics are not yet defined in CRI.
  - [#36401](https://github.com/kubernetes/kubernetes/issues/36401): The new
     container log path/format is not yet supported by the logging pipeline
    (e.g., fluentd, GCL).
  - CRI may not be compatible with other experimental features (e.g., Seccomp).
  - Streaming server needs to be hardened.
     - [#36666](https://github.com/kubernetes/kubernetes/issues/36666):
       Authentication.
     - [#36187](https://github.com/kubernetes/kubernetes/issues/36187): Avoid
       including user data in the redirect URL.

#### [Docker CRI integration known issues](#docker-cri-1.5-known-issues)

Note, these are known issues as of the 1.5 release. They may or may not have
been fixed since.

  - Docker compatibility: Support only Docker v1.11 and v1.12.
  - Network:
     - [#35457](https://github.com/kubernetes/kubernetes/issues/35457): Does
       not support host ports.
     - [#37315](https://github.com/kubernetes/kubernetes/issues/37315): Does
       not support bandwidth shaping.
  - Exec/attach/port-forward (streaming requests):
     - [#35747](https://github.com/kubernetes/kubernetes/issues/35747): Does
       not support `nsenter` as the exec handler (`--exec-handler=nsenter`).
     - Also see [CRI 1.5 known issues](#cri-1.5-known-issues) for limitations
       on CRI streaming.

## Contacts

  - Email: sig-node (sig-node@kubernetes.io)
  - Slack: https://kubernetes.slack.com/messages/sig-node

