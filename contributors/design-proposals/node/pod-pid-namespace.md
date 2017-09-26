# Shared PID Namespace

*   Status: Pending
*   Version: Alpha
*   Implementation Owner: verb

## Motivation

Pods share namespaces where possible, but a requirement for sharing the PID
namespace has not been defined due to lack of support in Docker. This created an
implicit API on which certain container images now rely. This proposal
explicitly defines behavior for sharing a PID namespace.

## Proposal

### Goals and Non-Goals

Goals include:

*   Backwards compatibility with container images expecting `pid == 1` semantics
*   Per-pod configuration of PID namespace sharing
*   Ability to change default sharing behavior in `v2.Pod`

Non-goals include:

*   Creating a general purpose init solution
*   Multiple shared PID namespaces per pod
*   Per-container configuration of PID namespace sharing

### Configurable PID Namespace Sharing

We will add first class support for configuring isolated and shared PID
namespaces for pods by adding a new boolean field `sharedPID` to the pod spec,
which will default to false. When set to true, all containers in the pod will
share a single PID namespace when supported by runtime.

The Container Runtime Interface (CRI) will be updated to require runtimes to
support three PID Namespace modes: Isolated, Shared & Host. The Runtime Manager
will translate the pod spec into one of these modes as follows:

Pod `sharedPID` | Pod `hostPID` | CRI PID Mode
--------------- | ------------- | ------------
false           | false         | Isolated
false           | true          | Host
true            | false         | Shared
true            | true          | *Error*

When a runtime does not implement a particular PID mode, it must return an
error. Docker will support all three modes when using version >= 1.13.1.

The shared PID functionality will be hidden behind a new feature gate in both
the API server and the kubelet, and the `--docker-disable-shared-pid` flag will
be removed from the kubelet.

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

## Implementation

### Kubernetes API Changes

`v1.PodSpec` gains a new field resembling the existing `HostPID` field:

```
// PodSpec is a description of a pod.
type PodSpec struct {
    ...
       // Use the host's pid namespace.
       // Optional: Default to false.
       // +k8s:conversion-gen=false
       // +optional
       HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
    // Use a single shared PID namespace for the pod.
       // Optional: Default to false.
       // +k8s:conversion-gen=false
       // +optional
       SharedPID bool `json:"sharedPID,omitempty" protobuf:"varint,XX,opt,name=sharedPID"`
       ...
```

Setting both `SharedPID` and `HostPID` will cause a validation error.

### Container Runtime Interface Changes

Namespace options in the CRI are specified for both PodSandbox and Container
creation requests. We will add a new field to NamespaceOption:

```
message NamespaceOption {
    ...
    // If set, use the host's PID namespace.
    bool host_pid = 2;
    ...
    // If set, use the pod's shared PID namespace.
    bool shared_pid = 4;
}
```

The kubelet runtime manager will set `shared_pid` to
`PodSecurityContext.SharedPID` for both the `LinuxSandboxSecurityContext` and
`LinuxContainerSecurityContext`.

### dockershim Changes

The Docker runtime implements the pod sandbox as a container running the pause
container image. The `shared_pid` of `LinuxSandboxSecurityContext` will be
ignored by this runtime as the sandbox container will become the namespace
that's shared with the containers in the pod.

The dockershim will translate a `CreateContainerRequest` with `shared_pid` set
to true in `LinuxContainerSecurityContext` to a call to Docker's create
container request with `PidMode` set to the container ID of the sandbox
container. All containers in the pod will then start in the same PID namespace
as the sandbox container, with the `/pause` binary as PID 1.

If the Docker runtime version does not support sharing pid namespaces, a
`CreateContainerRequest` with `shared_pid` set to true will return an error.

## Alternatives Considered

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

### Enumerated PID Mode

Rather than a boolean `SharedPID` we could create a single PID mode field that
could be set to `SHARED`, `HOST`, or `ISOLATED`. `HostPID` already exists in
`v1.PodSpec`, however, and we should prefer consistency within an API version to
attempting to define backwards-compatible semantics with an existing field.
