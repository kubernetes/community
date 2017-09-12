# Shared PID Namespace

*   Status: Pending
*   Version: Alpha
*   Implementation Owner: [@verb](https://github.com/verb)

## Motivation

Pods share namespaces where possible, but support for sharing the PID namespace
had not been defined due to lack of support in Docker. This created an implicit
API on which certain container images now rely. This document proposes adding
support for sharing a process namespace between containers in a pod while
maintaining backwards compatibility with the existing implicit API.

## Proposal

### Goals and Non-Goals

Goals include:

*   Backwards compatibility with container images expecting `pid == 1` semantics
*   Per-pod configuration of PID namespace sharing
*   Ability to change default sharing behavior in `v2.Pod`

Non-goals include:

*   Creating a general purpose container init solution
*   Multiple shared PID namespaces per pod
*   Per-container configuration of PID namespace sharing

### Summary

We will add support for configuring pod-shared process namespaces by adding a
new boolean field `ShareProcessNamespace` to the pod spec. The default to false
means that each container will have a separate process namespace. When set to
true, all containers in the pod will share a single process namespace.

The Container Runtime Interface (CRI) will be updated to support three namespace
modes: Container, Pod & Node. The Runtime Manager will translate the pod spec
into one of these modes as follows:

Pod `shareProcessNamespace` | Pod `hostPID` | CRI PID Mode
--------------------------- | ------------- | ------------
false                       | false         | Container
false                       | true          | Node
true                        | false         | Pod
true                        | true          | *Error*

If a runtime does not implement a particular PID mode, it must return an error.
For reference, Docker will support all three modes when using version >= 1.13.1.

The shared PID functionality will be hidden behind a new feature gate in both
the API server and the kubelet, and the existing `--docker-disable-shared-pid`
flag will be removed from the kubelet, subject to [deprecation
policy](https://kubernetes.io/docs/reference/deprecation-policy/).

## User Experience

### Use Cases

Sharing a PID namespace between containers in a pod is discussed in
[#1615](https://issues.k8s.io/1615) and enables:

1.  signaling between containers, which is useful for side cars (e.g. for
    signaling a daemon process after rotating logs).
1.  easier troubleshooting of pods.
1.  addressing [Docker's zombie
    problem](https://blog.phusion.nl/2015/01/20/docker-and-the-pid-1-zombie-reaping-problem/)
    by reaping orphaned zombies in the infra container.

### Behavioral Changes

Sharing a process namespace fits well with Kubernetes' pod abstraction, but it's
a significant departure from the traditional behavior of Docker. This may break
container images and development patterns that have come to rely on process
isolation. Notably:

1.  **The main container process no longer has PID 1**. It cannot be signalled
    using `kill 1`, and attempting to do so will instead signal the
    infrastructure container and potentially restart the pod. Containers
    shipping an init system like systemd may [require additional
    flags](https://github.com/kubernetes/kubernetes/issues/48937#issuecomment-321243669).
1.  **Processes are visible to other containers in the pod**. This includes all
    information visible in `/proc`, such as passwords as arguments or
    environment variables, and process signalling. This can be somewhat
    mitigated by running processes as separate, non-root users.
1.  **Container filesystems are visible to other containers in the pod through
    the <code>/proc/$pid/root</code> magic symlink**. This makes debugging
    easier, but it also means that secrets are protected only by standard
    filesystem permissions.

## Implementation

### Kubernetes API Changes

`v1.PodSpec` gains a new field named `ShareProcessNamespace`:

```
// PodSpec is a description of a pod.
type PodSpec struct {
    ...
       // Use the host's pid namespace.
       // Note that HostPID and ShareProcessNamespace cannot both be set.
       // Optional: Default to false.
       // +k8s:conversion-gen=false
       // +optional
       HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
       // Share a single process namespace between all of the containers in a pod.
       // Note that HostPID and ShareProcessNamespace cannot both be set.
       // Optional: Default to false.
       // +k8s:conversion-gen=false
       // +optional
       ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty" protobuf:"varint,XX,opt,name=shareProcessNamespace"`
       ...
```

The field name deviates from that of HostPID in an attempt to [better signal the
consequences](https://github.com/kubernetes/community/pull/1048/files#r159146536)
of setting the option. Setting both `ShareProcessNamespace` and `HostPID` will
cause a validation error.

### Container Runtime Interface Changes

Namespace options in the CRI are currently specified for both `PodSandbox` and
`Container` creation requests via booleans in `NamespaceOption`:

```
message NamespaceOption {
    // If set, use the host's network namespace.
     bool host_network = 1;
     // If set, use the host's PID namespace.
     bool host_pid = 2;
     // If set, use the host's IPC namespace.
     bool host_ipc = 3;
}
```

We will change `NamespaceOption` to use a `NamespaceMode` enumeration for the
existing namespace options:

```
enum NamespaceMode {
    POD       = 0;
    CONTAINER = 1;
    NODE      = 2;
}

// NamespaceOption provides options for Linux namespaces.
message NamespaceOption {
    // Network namespace for this container/sandbox.
    // Runtimes must support: POD, NODE
    NamespaceMode network = 1;
    // PID namespace for this container/sandbox.
    // Note: The CRI default is POD, but the v1.PodSpec default is CONTAINER.
    // The kubelet's runtime manager will set this to CONTAINER explicitly for v1 pods.
    // Runtimes must support: POD, CONTAINER, NODE
    NamespaceMode pid = 2;
    // IPC namespace for this container/sandbox.
    // Runtimes must support: POD, NODE
    NamespaceMode ipc = 3;
}
```

Note that this breaks backwards compatibility in the CRI, which is still in
alpha.

The protocol default for a namespace is `POD` because that's the default for
network and IPC, and we will consider making it the default for PID in `v2.Pod`.
The kubelet will explicitly set `pid` to `CONTAINER` for `v1.Pod` by default so
that the default behavior of `v1.Pod` does not change.

This CRI design allows different namespace configuration for each of the
containers in the pod and the sandbox, but currently we have no plans to support
this in the Kubernetes API. The kubelet will translate namespace booleans from
v1.PodSpec into a single `NamespaceMode` to be used for the sandbox and all
regular and init containers in a pod.

#### Targeting a Specific Container's Namespace

Though we don't intend to support this in general pod configuration, there is a
use case for mixed process namespaces within a single pod. [Troubleshooting
Running Pods](troubleshooting-running-pods.md) allows inserting an ephemeral
Debug Container in an existing, running pod. In order for this to be useful we
want to share, within the pod, a process namespace between the new container
performing the debugging and its existing target container.

This is done with the additional `NamespaceMode` `TARGET` and field `target_id`:

```
enum NamespaceMode {
    POD       = 0;
    CONTAINER = 1;
    NODE      = 2;
    TARGET    = 3;
}

// NamespaceOption provides options for Linux namespaces.
message NamespaceOption {
    // Network namespace for this container/sandbox.
    // Runtimes must support: POD, NODE
    NamespaceMode network = 1;
    // PID namespace for this container/sandbox.
    // Note: The CRI default is POD, but the v1.PodSpec default is CONTAINER.
    // The kubelet's runtime manager will set this to CONTAINER explicitly for v1 pods.
    // Runtimes must support: POD, CONTAINER, NODE, TARGET
    NamespaceMode pid = 2;
    // IPC namespace for this container/sandbox.
    // Runtimes must support: POD, NODE
    NamespaceMode ipc = 3;
    // Target Container ID for NamespaceMode of TARGET. This container must be in the
    // same pod as the target container.
    string target_id = 4;
}
```

When `NamespaceOption.pid` is set to `TARGET`, a runtime must create the new
container in the namespace used by the container ID in `target_id`. If the
target container has `NamespaceOption.pid` set to `POD`, then the new container
should also use the pod namespace. If the target container has an isolated
process namespace, then the new container will join only that container's
namespace. Examples are provided for dockershim below.

There is no mechanism in the Kubernetes API for an end-user to set `TARGET`. It
exists for the kubelet to run automation or debugging from a container image in
the namespace of an existing pod and container. Additionally, we choose to
explicitly not support sharing namespaces between different pods. The kubelet
must not generate such a reference, and the runtime should not accept it. That
is, for pod{Container `A`, Container `B`, Sandbox `S}` and any other unrelated
Container `C`:

valid `target_id` | invalid `target_id`
----------------- | -------------------
containerID(A)    | sandboxID(S)
containerID(B)    | containerID(C)

### dockershim Changes

The Docker runtime implements the pod sandbox as a container running the pause
container image. When configured for `POD` namespace sharing, the PID namespace
of the sandbox will become the single PID namespace for the pod. This means a
namespace of `POD` and `CONTAINER` are equivalent for the sandbox. The mapping
of the _sandbox's_ PID mode to docker's `HostConfig.PidMode` is (`v1.Pod`
settings provided as reference):

ShareProcessNamespace | HostPID | Sandbox PID Mode | HostConfig.PidMode
--------------------- | ------- | ---------------- | ------------------
false                 | false   | CONTAINER        | *unset*
true                  | false   | POD              | *unset*
false                 | true    | NODE             | "host"
\-                    | \-      | TARGET           | *Error*

For _containers_, `HostConfig.PidMode` will be set as follows:

ShareProcessNamespace | HostPID | Container PID Mode | HostConfig.PidMode
--------------------- | ------- | ------------------ | ------------------
false                 | false   | CONTAINER          | *unset*
true                  | false   | POD                | "container:[sandbox-container-id]"
false                 | true    | NODE               | "host"
false                 | false   | TARGET             | "container:[target-container-id]"
true                  | false   | TARGET             | "container:[sandbox-container-id]"
false                 | true    | TARGET             | "host"

If the Docker runtime version does not support sharing pid namespaces, a
`CreateContainerRequest` with `namespace_options.pid` set to `POD` will return
an error.

### Deprecation of existing kubelet flag

SIG Node did not anticipate the strong objections to migrating from isolated to
shared process namespaces for Docker. The previous (now abandoned) migration
plan introduced a kubelet flag to toggle the shared namespace behavior, but
objections did not materialize until the flag had moved from experimental to GA.

The `--docker-disable-shared-pid` (default: true) kubelet flag disables the use
of shared process namespaces for the Docker runtime. We will immediately mark it
as deprecated, but according to the [deprecation
policy](https://kubernetes.io/docs/reference/deprecation-policy/) we must
support it for 6 months.

We must provide a transition path for users setting this kubelet flag to false.
Setting this flag asserts a desire to override the default Kubernetes behavior
for all pods. Until the flag is removed, the kubelet will honor this assertion
by ignoring the value of `ShareProcessNamespace` and logging a warning to the
event log.

## Alternatives Considered

### Explicit Container/Sandbox ID Targeting

Rather than using a `NamespaceMode`, `NamespaceOption.pid` could be a string
that explicitly targets a container or sandbox ID:

```
// NamespaceOption provides options for Linux namespaces.
message NamespaceOption {
    ...
    // ID of Sandbox or Container to use for PID namespace, or "host"
    string pid = 2;
    ...
}
```

This removes the need for a separate `TARGET` mode, but a mode enumeration
better captures the intent of the option.

### Defaulting to PID Namespace Sharing

Other Kubernetes runtimes already share a single PID namespace between
containers in a pod. We could easily change the Docker runtime to always share a
PID namespace when supported by the installed Docker version, but this would
cause problems for container images that assume they will always be PID 1.

### Migration to Shared-only Namespaces

Rather than adding support to the API for configuring namespaces we could allow
changing the default behavior with pod annotations with the intention of
removing support for isolated PID namespaces in v2.Pod. Many members of the
community want to use the isolated namespaces as security boundary between
containers in a pod, however.
