---
kep-number: 14
title: Runtime Class
authors:
  - "@tallclair"
owning-sig: sig-node
participating-sigs:
  - sig-architecture
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-06-19
status: provisional
---

# Runtime Class

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non\-Goals](#non-goals)
  * [User Stories](#user-stories)
* [Proposal](#proposal)
  * [API](#api)
    * [API Specification](#api-specification)
    * [CRI API Specification](#cri-api-specification)
  * [Versioning, Updates, and Rollouts](#versioning-updates-and-rollouts)
  * [Implementation Details](#implementation-details)
  * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Appendix](#appendix)
  * [Examples of runtime variation](#examples-of-runtime-variation)

## Summary

`RuntimeClass` is a new cluster-scoped resource that surfaces container runtime properties to the
control plane. RuntimeClasses are assigned to pods through a `runtimeClass` field on the
`PodSpec`. This provides a new mechanism for supporting multiple runtimes in the cluster and on the
nodes.

## Motivation

RuntimeClasses address two distinct problems: surfacing support and configuration of optional
features, and supporting multiple runtimes.

Today's runtimes have varying support for features that are expressible through the PodSpec (see
[examples of runtime variation](#examples-of-runtime-variation) below). In some cases such as
SELinux, features are silently ignored when they aren't supported. This is problematic from a
security perspective since it makes it difficult to audit actual runtime options. Other features
such as AppArmor are checked for support, but on failure the pod is just held in a "Blocked" state
to avoid the churn of recreating the rejected pod. Still other features maybe be applied in slightly
different configurations, as is the case with capabilities today (the default set is not
well-defined), and `privileged` containers. By raising these properties to the control plane errors
can be surfaced much earlier, and configuration can be made explicit.

The second problem that RuntimeClass addresses is supporting multiple runtimes in a cluster, or even
on the same node. [Sandboxes][] are the primary motivator for this right now, with both Kata
containers and gVisor looking to integrate with Kubernetes. Other runtime models such as Windows
containers or even remote runtimes will also require support in the future. RuntimeClass provides a
way to configure different runtimes in the cluster, surface their properties (both to the cluster &
the user), and provides a mechanism for pods to select a runtime.

[Sandboxes]: https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit

### Goals

- Provide a mechanism for surfacing container runtime properties to the control plane
- Support multiple runtimes per-cluster, and provide a mechanism for users to select the desired
  runtime

### Non-Goals

- RuntimeClass is NOT a general policy mechanism
- RuntimeClass is NOT "NodeClass". Although different nodes may run different runtimes, in general
  RuntimeClass should not be a cross product of runtime properties and node properties.

The following goals are out-of-scope for the initial implementation, but may be explored in a future
iteration:

- Automatic runtime or feature discovery - initially RuntimeClasses are manually defined (by the
  cluster admin or provider), and are asserted to be an accurate representation of the runtime.
- Scheduling in heterogeneous clusters - it is possible to operate a heterogeneous cluster
  (different runtime configurations on different nodes) through scheduling primitives like
  `NodeAffinity` and `Taints+Tolerations`, but the user is responsible for setting these up and
  automatic runtime-aware scheduling is out-of-scope.
- Define standardized or conformant runtime classes - although I would like to declare some
  predefined RuntimeClasses with specific properties, doing so is out-of-scope for this initial KEP.
- [Pod Overhead][] - Although RuntimeClass is likely to be the configuration mechanism of choice,
  the details of how pod resource overhead will be implemented is out of scope for this KEP.
- Provide a mechanism to dynamically register or provision additional runtimes.
- Requiring specific RuntimeClasses according to policy. This should be addressed by other
  cluster-level policy mechanisms, such as PodSecurityPolicy.
- "Fitting" a RuntimeClass to pod requirements - In other words, specifying runtime properties and
  letting the system match an appropriate RuntimeClass, rather than explicitly assigning a
  RuntimeClass by name. This approach can increase portability, but can be added seamlessly in a
  future iteration.

[Pod Overhead]: https://docs.google.com/document/d/1EJKT4gyl58-kzt2bnwkv08MIUZ6lkDpXcxkHqCvvAp4/edit

### User Stories

- As a cluster operator, I want to provide multiple runtime options to support a wide variety of
  workloads. Examples include native linux containers, "sandboxed" containers, and windows
  containers.
- As a cluster operator, I want to provide stable rolling upgrades of nodes.
- As an application developer, I want to select the runtime that best fits my workload.
- As an application developer, I don't want to study the nitty-gritty details of different runtime
  implementations.
- As an application developer, I want my application to be portable across clusters that use similar
  but different variants of a "class" of runtimes.
- As an application operator, I want to know precisely how the application is running, and which
  features are being used.
- As an application operator, I want pods to fail fast when trying to use an optional feature that
  isn't supported by the chosen runtime.
- As a security operator, I want to ensure that pods are running with the security features they
  claim.

## Proposal

The initial design includes:

- `RuntimeClass` API resource definition
- `RuntimeClass` pod field for specifying the RuntimeClass the pod should be run with
- Kubelet implementation for fetching & interpreting the RuntimeClass
- CRI API & implementation for passing along the RuntimeClass definition

A subsequent update will include the API for describing supported features and an admission
controller for performing validation & defaulting.

### API

`RuntimeClass` is a new cluster-scoped resource in the `node.k8s.io` API group.

> _The `node.k8s.io` API group would eventually hold the Node resource when `core` is retired.
> Alternatives considered: `runtime.k8s.io`, `cluster.k8s.io`_

There are 3 categories of fields that may be included in the `RuntimeClassSpec`:

1. **Runtime parameters** - These fields control the behavior of the runtime in some way,
   including configuring arbitrary options in the runtime (`Parameters`), and eventually specifying
   resource overhead. This is the focus of this initial proposal.
2. **Compatibility validation** - These fields declare support for various features, and pods are
   validated against them in admission. This category, along with the admission controller, will be
   addressed in a subsequent update to this proposal.
3. **Informational** - These fields surface semantic properties of the runtime to guide runtime
   selection. Examples could include hardware virtualization, stability, and I/O performance
   class. These fields are out-of-scope for the initial proposal.

#### API Specification

_(This is a simplified declaration, syntactic details will be covered in the API PR review)_

```go
type RuntimeClass struct {
    metav1.TypeMeta
    // ObjectMeta minimally includes the RuntimeClass name, which is used to reference the class.
    // Namespace should be left blank.
    metav1.ObjectMeta

    Spec RuntimeClassSpec
}

type RuntimeClassSpec struct {
    // Parameters holds opaque parameters that are interpreted by the runtime (CRI implementation).
    Parameters map[string]string
}
```

The runtime is selected by the pod by specifying the RuntimeClass in the PodSpec:

```go
type PodSpec struct {
    ...
    // RuntimeClassName refers to a RuntimeClass object with the same name,
    // which should be used to run this pod.
    // +optional
    RuntimeClassName string
    ...
}
```

The `legacy` RuntimeClass name is reserved. The legacy RuntimeClass is defined to be fully backwards
compatible with current Kubernetes. This means that the legacy runtime does not specify any runtime
parameters or perform any feature validation (all features are "supported").

```go
const (
    // RuntimeClassNameLegacy is a reserved RuntimeClass name. The legacy
    // RuntimeClass does not specify any runtime parameters or perform any
    // feature validation.
    RuntimeClassNameLegacy = "legacy"
)
```

An unspecified RuntimeClassName `""` is equivalent to the `legacy` RuntimeClass, though the field is
not defaulted to `legacy` (to leave room for configurable defaults in a future update).

#### CRI API Specification

The runtime parameters are passed to the CRI as part of the `RunPodSandboxRequest`:

```proto
message RunPodSandboxRequest {
    // Configuration for creating a PodSandbox.
    PodSandboxConfig config = 1;
    // Configuration for the runtime.
    map<string, string> RuntimeParameters = 2;
}
```

### Versioning, Updates, and Rollouts

_Axiom: Rerunning a pod on a different node in the same cluster with the same RuntimeClassName should
not change its behavior._

This means that if a runtime is updated in a fully backwards-compatible way, then the RuntimeClass
does not _need_ to change. It also means that if the updated runtime breaks backwards compatibility,
then it **MUST** be renamed. RuntimeClass does not provide a predefined versioning or naming scheme,
so versioned upgrades are left to the cluster operator to handle.

Since breaking runtime changes require a new RuntimeClass definition, it means that pods must be
updated to reference the new RuntimeClass. This comes with the advantage of native support for
rolling updates through the same mechanisms as any other application update.

### Implementation Details

The Kubelet uses an Informer to keep a local cache of all RuntimeClass objects. When a new pod is
added, the Kubelet resolves the Pod's RuntimeClass against the local RuntimeClass cache.  Once
resolved, a subset of its fields are passed to the CRI as part of the [`RunPodSandboxRequest`][]. At
that point, the interpretation of the parameters is left to the CRI implementation, but they should
be cached if needed for subsequent calls.

[RunPodSandboxRequest]: https://github.com/kubernetes/kubernetes/blob/b05a61e299777c2030fbcf27a396aff21b35f01b/pkg/kubelet/apis/cri/runtime/v1alpha2/api.proto#L344

### Risks and Mitigations

**Scope creep.** RuntimeClass has a fairly broad charter, but it should not become a default
dumping ground for every new feature exposed by the node. For each feature, careful consideration
should be made about whether it belongs on the Pod, Node, RuntimeClass, or some other resource. The
[non-goals](#non-goals) should be kept in mind when considering RuntimeClass features. Features that
are too specific to a single runtime implementation should use `Parameters` rather than expanding
the `RuntimeClassSpec`.

**Becoming a general policy mechanism.** RuntimeClass should not be used a replacement for
PodSecurityPolicy. The use cases for defining multiple RuntimeClasses for the same underlying
runtime implementation should be extremely limited (generally only around updates & rollouts). To
enforce this, no authorization or restrictions are placed directly on RuntimeClass use; in order to
restrict a user to a specific RuntimeClass, you must use another policy mechanism such as
PodSecurityPolicy.

**Pushing complexity to the user.** RuntimeClass is a new resource in order to hide the complexity
of runtime configuration from most users (aside from the cluster admin or provisioner). However, we
are still side-stepping the issue of precisely defining specific types of runtimes like
"Sandboxed". However, it is still up for debate whether precisely defining such runtime categories
is even possible. RuntimeClass allows us to decouple this specification from the implementation, but
it is still something I hope we can address in a future iteration through the concept of pre-defined
or "conformant" RuntimeClasses.

**Non-portability.** We are already in a world of non-portability for many features (see [examples
of runtime variation](#examples-of-runtime-variation). RuntimeClasses actually address this issue by
formally declaring supported features. Another issue is that pods need to refer to a RuntimeClass by
name, which may not be defined in every cluster. This is something that can be addressed through
pre-defined runtime classes (see previous risk), and/or by "fitting" pod requirements to compatible
RuntimeClasses.

## Graduation Criteria

Alpha:

- Everything described in the current proposal
- Supported features and associated admission controller

Beta:

- Major runtimes support RuntimeClass
- RuntimeClasses are configured in the E2E environment with test coverage of a non-legacy RuntimeClass
- TBD

## Implementation History

- 2018-06-11: SIG-Node decision to move forward with proposal
- 2018-06-19: Initial KEP published.

## Appendix

### Examples of runtime variation

- Linux Security Module (LSM) choice - Kubernetes supports both AppArmor & SELinux options on pods,
  but those are mutually exclusive, and support of either is not required by the runtime. The
  default configuration is also not well defined.
- Seccomp-bpf - Kubernetes has alpha support for specifying a seccomp profile, but the default is
  defined by the runtime, and support is not guaranteed.
- Windows containers - isolation features are very OS-specific, and most of the current features are
  limited to linux. As we build out Windows container support, we'll need to add windows-specific
  features as well.
- Host namespaces (Network,PID,IPC) may not be supported by virtualization-based runtimes
  (e.g. Kata-containers & visor).
- Per-pod and Per-container resource overhead varies by runtime.
- Device support (e.g. GPUs) varies wildly by runtime & nodes.
- Supported volume types varies by node - it remains TBD whether this information belongs in
  RuntimeClass.
- The list of default capabilities is defined in Docker, but not Kubernetes. Future runtimes may
  have differing defaults, or support a subset of capabilities.
- `Privileged` mode is not well defined, and thus may have differing implementations.
- Support for resource over-commit and dynamic resource sizing (e.g. Burstable vs Guaranteed
  workloads)
