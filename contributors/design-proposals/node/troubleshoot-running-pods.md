# Troubleshoot Running Pods

*   Status: Implementing
*   Version: Alpha
*   Implementation Owner: @verb

This proposal seeks to add first class support for troubleshooting by creating a
mechanism to execute a shell or other troubleshooting tools inside a running pod
without requiring that the associated container images include such tools.

## Motivation

### Development

Many developers of native Kubernetes applications wish to treat Kubernetes as an
execution platform for custom binaries produced by a build system. These users
can forgo the scripted OS install of traditional Dockerfiles and instead `COPY`
the output of their build system into a container image built `FROM scratch` or
a
[distroless container image](https://github.com/GoogleCloudPlatform/distroless).
This confers several advantages:

1.  **Minimal images** lower operational burden and reduce attack vectors.
1.  **Immutable images** improve correctness and reliability.
1.  **Smaller image size** reduces resource usage and speeds deployments.

The disadvantage of using containers built `FROM scratch` is the lack of system
binaries provided by an Operating System image makes it difficult to
troubleshoot running containers. Kubernetes should enable one to troubleshoot
pods regardless of the contents of the container images.

### Operations and Support

As Kubernetes gains in popularity, it's becoming the case that a person
troubleshooting an application is not necessarily the person who built it.
Operations staff and Support organizations want the ability to attach a "known
good" or automated debugging environment to a pod.

## Requirements

A solution to troubleshoot arbitrary container images MUST:

*   troubleshoot arbitrary running containers with minimal prior configuration
*   allow access to namespaces and the file systems of individual containers
*   fetch troubleshooting utilities at debug time rather than at the time of pod
    creation
*   be compatible with admission controllers and audit logging
*   allow discovery of current debugging status
*   support arbitrary runtimes via the CRI (possibly with reduced feature set)
*   require no administrative access to the node
*   have an excellent user experience (i.e. should be a feature of the platform
    rather than config-time trickery)
*   have no _inherent_ side effects to the running container image
*   v1.Container must be available for inspection by admission controllers

## Feature Summary

Any new debugging functionality will require training users. We can ease the
transition by building on an existing usage pattern. We will create a new
command, `kubectl debug`, which parallels an existing command, `kubectl exec`.
Whereas `kubectl exec` runs a _process_ in a _container_, `kubectl debug` will
be similar but run a _container_ in a _pod_.

A container created by `kubectl debug` is a _Debug Container_. Unlike `kubectl
exec`, Debug Containers have status that is reported in `PodStatus` and
displayed by `kubectl describe pod`.

For example, the following command would attach to a newly created container in
a pod:

```
kubectl debug -c debug-shell --image=debian target-pod -- bash
```

It would be reasonable for Kubernetes to provide a default container name and
image, making the minimal possible debug command:

```
kubectl debug target-pod
```

This creates an interactive shell in a pod which can examine and signal other
processes in the pod. It has access to the same network and IPC as processes in
the pod. When [process namespace sharing](https://features.k8s.io/495) is
enabled, it can access the filesystem of other processes by `/proc/$PID/root`.
Debug Containers can enter arbitrary namespaces of another visible container via
`nsenter` when run with `CAP_SYS_ADMIN`.

_Please see the User Stories section for additional examples and Alternatives
Considered for the considerable list of other solutions we considered._

## Implementation Details

From the perspective of the user, there's a new command, `kubectl debug`, that
creates a Debug Container and attaches to its console. We believe a new command
will be less confusing for users than overloading `kubectl exec` with a new
concept. Users give Debug Containers a name (e.g. "debug" or "shell") which can
subsequently be used to reattach and is reported by `kubectl describe`.

### Kubernetes API Changes

This will be implemented in the Core API to avoid new dependencies in the
kubelet. The user-level concept of a _Debug Container_ implemented with the
API-level concept of an _Ephemeral Container_. The API doesn't require an
Ephemeral Container to be used as a Debug Container. It's intended as a general
purpose construct for running a short-lived process in a pod.

#### Pod Changes

Ephemeral Containers are represented in `PodSpec` and `PodStatus`:

```
type PodSpec struct {
  ...
  // List of user-initiated ephemeral containers to run in this pod.
  // This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.
  // +optional
  EphemeralContainers []EphemeralContainer `json:"ephemeralContainers,omitempty" protobuf:"bytes,29,opt,name=ephemeralContainers"`
}

type PodStatus struct {
  ...
  // Status for any Ephemeral Containers that running in this pod.
  // This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.
  // +optional
  EphemeralContainerStatuses []ContainerStatus `json:"ephemeralContainerStatuses,omitempty" protobuf:"bytes,12,rep,name=ephemeralContainerStatuses"`
}
```

`EphemeralContainerStatuses` resembles the existing `ContainerStatuses` and
`InitContainerStatuses`, but `EphemeralContainers` introduces a new type:

```
// An EphemeralContainer is a container which runs temporarily in a pod for human-initiated actions
// such as troubleshooting. This is an alpha feature enabled by the EphemeralContainers feature flag.
type EphemeralContainer struct {
  // Spec describes the Ephemeral Container to be created.
  Spec Container `json:"spec,omitempty" protobuf:"bytes,1,opt,name=spec"`

  // If set, the name of the container from PodSpec that this ephemeral container targets.
  // The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container.
  // If not set then the ephemeral container is run in whatever namespaces are shared
  // for the pod.
  // +optional
  TargetContainerName string `json:"targetContainerName,omitempty" protobuf:"bytes,2,opt,name=targetContainerName"`
}
```

Much of the utility of Ephemeral Containers comes from the ability to run a
container within the PID namespace of another container. `TargetContainerName`
allows targeting a container that doesn't share its PID namespace with the rest
of the pod. We must modify the CRI to enable this functionality (see below).

##### Alternative Considered: Omitting TargetContainerName

It would be simpler for the API, kubelet and kubectl if `EphemeralContainers`
was a `[]Container`, but as isolated PID namespaces will be the default for some
time, being able to target a container will provide a better user experience.

#### Updates

Most fields of `Pod.Spec` are immutable once created. There is a short whitelist
of fields which may be updated, and we could extend this to include
`EphemeralContainers`. The ability to add new containers is a large change for
Pod, however, and we'd like to begin conservatively by enforcing the following
best practices:

1.  Ephemeral Containers lack guarantees for resources or execution, and they
    will never be automatically restarted. To avoid pods that depend on
    Ephemeral Containers, we allow their addition only in pod updates and
    disallow them during pod create.
1.  Some fields of `v1.Container` imply a fundamental role in a pod. We will
    disallow the following fields in Ephemeral Containers: `resources`, `ports`,
    `livenessProbe`, `readinessProbe`, and `lifecycle.`
1.  Cluster administrators may want to restrict access to Ephemeral Containers
    independent of other pod updates.

To enforce these restrictions and new permissions, we will introduce a new Pod
subresource, `/ephemeralcontainers`. `EphemeralContainers` can only be modified
via this subresource. `EphemeralContainerStatuses` is updated with everything
else in `Pod.Status` via `/status`.

To create a new Ephemeral Container, one appends a new `EphemeralContainer` with
the desired `v1.Container` as `Spec` in `Pod.Spec.EphemeralContainers` and
`PUT`s the pod to `/ephemeralcontainers`.

The subresources `attach`, `exec`, `log`, and `portforward` are available for
Ephemeral Containers and will be forwarded by the apiserver. This means `kubectl
attach`, `kubelet exec`, `kubectl log`, and `kubectl port-forward` will work for
Ephemeral Containers.

Once the pod is updated, the kubelet worker watching this pod will launch the
Ephemeral Container and update its status. The client is expected to watch for
the creation of the container status and then attach to the console of a debug
container using the existing attach endpoint,
`/api/v1/namespaces/$NS/pods/$POD_NAME/attach`. Note that any output of the new
container occurring between its creation and attach will not be replayed, but it
can be viewed using `kubectl log`.

##### Alternative Considered: Standard Pod Updates

It would simplify initial implementation if we updated the pod spec via the
normal means, and switched to a new update subresource if required at a future
date. It's easier to begin with a too-restrictive policy than a too-permissive
one on which users come to rely, and we expect to be able to remove the
`/ephemeralcontainers` subresource prior to exiting alpha should it prove
unnecessary.

### Container Runtime Interface (CRI) changes

The CRI requires no changes for basic functionality, but it will need to be
updated to support container namespace targeting, as described in the
[Shared PID Namespace Proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/pod-pid-namespace.md#targeting-a-specific-containers-namespace).

### Creating Debug Containers

To create a debug container, kubectl will take the following steps:

1.  `kubectl` constructs an `EphemeralContainer` based on command line arguments
    and appends it to `Pod.Spec.EphemeralContainers`. It `PUT`s the modified pod
    to the pod's `/ephemeralcontainers`.
1.  The apiserver discards changes other than additions to
    `Pod.Spec.EphemeralContainers` and validates the pod update.
    1.  Pod validation fails if container spec contains fields disallowed for
        Ephemeral Containers or the same name as a container in the spec or
        `EphemeralContainers`.
    1.  API resource versioning resolves update races.
1.  The kubelet's pod watcher notices the update and triggers a `syncPod()`.
    During the sync, the kubelet calls `kuberuntime.StartEphemeralContainer()`
    for any new Ephemeral Container.
    1.  `StartEphemeralContainer()` uses the existing `startContainer()` to
        start the Ephemeral Container.
    1.  After initial creation, future invocations of `syncPod()` will publish
        its ContainerStatus but otherwise ignore the Ephemeral Container. It
        will exist for the life of the pod sandbox or it exits. In no event will
        it be restarted.
1.  `syncPod()` finishes a regular sync, publishing an updated PodStatus (which
    includes the new `EphemeralContainer`) by its normal, existing means.
1.  The client performs an attach to the debug container's console.

There are no limits on the number of Debug Containers that can be created in a
pod, but exceeding a pod's resource allocation may cause the pod to be evicted.

### Restarting and Reattaching Debug Containers

Debug Containers will not be restarted.

We want to be more user friendly by allowing re-use of the name of an exited
debug container, but this will be left for a future improvement.

One can reattach to a Debug Container using `kubectl attach`. When supported by
a runtime, multiple clients can attach to a single debug container and share the
terminal. This is supported by Docker.

### Killing Debug Containers

Debug containers will not be killed automatically unless the pod is destroyed.
Debug Containers will stop when their command exits, such as exiting a shell.
Unlike `kubectl exec`, processes in Debug Containers will not receive an EOF if
their connection is interrupted.

A future improvement to Ephemeral Containers could allow killing Debug
Containers when they're removed the `EphemeralContainers`, but it's not clear
that we want to allow this. Removing an Ephemeral Container spec makes it
unavailable for future authorization decisions (e.g. whether to authorize exec
in a pod that had a privileged Ephemeral Container).

### Security Considerations

Debug Containers have no additional privileges above what is available to any
`v1.Container`. It's the equivalent of configuring an shell container in a pod
spec except that it is created on demand.

Admission plugins must be updated to guard `/ephemeralcontainers`. They should
apply the same container image and security policy as for regular containers.

### Additional Consideration

1.  Debug Containers are intended for interactive use and always have TTY and
    Stdin enabled.
1.  There are no guaranteed resources for ad-hoc troubleshooting. If
    troubleshooting causes a pod to exceed its resource limit it may be evicted.
1.  There's an output stream race inherent to creating then attaching a
    container which causes output generated between the start and attach to go
    to the log rather than the client. This is not specific to Ephemeral
    Containers and exists because Kubernetes has no mechanism to attach a
    container prior to starting it. This larger issue will not be addressed by
    Ephemeral Containers, but Ephemeral Containers would benefit from future
    improvements or work arounds.
1.  Ephemeral Containers should not be used to build services, which we've
    attempted to reflect in the API.

## Implementation Plan

### 1.12: Initial Alpha Release

We're targeting an alpha release in Kubernetes 1.12 that includes the following
basic functionality:

1.  Approval for basic core API changes to Pod
1.  Basic support in the kubelet for creating Ephemeral Containers

Functionality out of scope for 1.12:

*   Killing running Ephemeral Containers by removing them from the Pod Spec.
*   Updating `pod.Spec.EphemeralContainers` when containers are garbage
    collected.
*   `kubectl` commands for creating Ephemeral Containers

Functionality will be hidden behind an alpha feature flag and disabled by
default.

## Appendices

We've researched many options over the life of this proposal. These Appendices
are included as optional reference material. It's not necessary to read this
material in order to understand the proposal in its current form.

### Appendix 1: User Stories

These user stories are intended to give examples how this proposal addresses the
above requirements.

#### Operations

Jonas runs a service "neato" that consists of a statically compiled Go binary
running in a minimal container image. One of the its pods is suddenly having
trouble connecting to an internal service. Being in operations, Jonas wants to
be able to inspect the running pod without restarting it, but he doesn't
necessarily need to enter the container itself. He wants to:

1.  Inspect the filesystem of target container
1.  Execute debugging utilities not included in the container image
1.  Initiate network requests from the pod network namespace

This is achieved by running a new "debug" container in the pod namespaces. His
troubleshooting session might resemble:

```
% kubectl debug -it -m debian neato-5thn0 -- bash
root@debug-image:~# ps x
  PID TTY      STAT   TIME COMMAND
    1 ?        Ss     0:00 /pause
   13 ?        Ss     0:00 bash
   26 ?        Ss+    0:00 /neato
  107 ?        R+     0:00 ps x
root@debug-image:~# cat /proc/26/root/etc/resolv.conf
search default.svc.cluster.local svc.cluster.local cluster.local
nameserver 10.155.240.10
options ndots:5
root@debug-image:~# dig @10.155.240.10 neato.svc.cluster.local.

; <<>> DiG 9.9.5-9+deb8u6-Debian <<>> @10.155.240.10 neato.svc.cluster.local.
; (1 server found)
;; global options: +cmd
;; connection timed out; no servers could be reached
```

Thus Jonas discovers that the cluster's DNS service isn't responding.

#### Debugging

Thurston is debugging a tricky issue that's difficult to reproduce. He can't
reproduce the issue with the debug build, so he attaches a debug container to
one of the pods exhibiting the problem:

```
% kubectl debug -it --image=gcr.io/neato/debugger neato-5x9k3 -- sh
Defaulting container name to debug.
/ # ps x
PID   USER     TIME   COMMAND
    1 root       0:00 /pause
   13 root       0:00 /neato
   26 root       0:00 sh
   32 root       0:00 ps x
/ # gdb -p 13
...
```

He discovers that he needs access to the actual container, which he can achieve
by installing busybox into the target container:

```
root@debug-image:~# cp /bin/busybox /proc/13/root
root@debug-image:~# nsenter -t 13 -m -u -p -n -r /busybox sh


BusyBox v1.22.1 (Debian 1:1.22.0-9+deb8u1) built-in shell (ash)
Enter 'help' for a list of built-in commands.

/ # ls -l /neato
-rwxr-xr-x    2 0        0           746888 May  4  2016 /neato
```

Note that running the commands referenced above require `CAP_SYS_ADMIN` and
`CAP_SYS_PTRACE`.

#### Automation

Ginger is a security engineer tasked with running security audits across all of
her company's running containers. Even though his company has no standard base
image, she's able to audit all containers using:

```
% for pod in $(kubectl get -o name pod); do
    kubectl debug -m gcr.io/neato/security-audit -p $pod /security-audit.sh
  done
```

#### Technical Support

Roy's team provides support for his company's multi-tenant cluster. He can
access the Kubernetes API (as a viewer) on behalf of the users he's supporting,
but he does not have administrative access to nodes or a say in how the
application image is constructed. When someone asks for help, Roy's first step
is to run his team's autodiagnose script:

```
% kubectl debug --image=k8s.gcr.io/autodiagnose nginx-pod-1234
```

### Appendix 2: Requirements Analysis

Many people have proposed alternate solutions to this problem. This section
discusses how the proposed solution meets all of the stated requirements and is
intended to contrast the alternatives listed below.

**Troubleshoot arbitrary running containers with minimal prior configuration.**
This solution requires no prior configuration.

**Access to namespaces and the file systems of individual containers.** This
solution runs a container in the shared pod namespaces (e.g. network) and will
attach to the PID namespace of a target container when not shared with the
entire pod. It relies on the behavior of `/proc/<pid>/root` to provide access to
filesystems of individual containers.

**Fetch troubleshooting utilities at debug time**. This solution uses normal
container image distribution mechanisms to fetch images when the debug command
is run.

**Respect admission restrictions.** Requests from kubectl are proxied through
the apiserver and so are available to existing
[admission controllers](https://kubernetes.io/docs/admin/admission-controllers/).
Plugins already exist to intercept `exec` and `attach` calls, but extending this
to support `debug` has not yet been scoped.

**Allow introspection of pod state using existing tools**. The list of
`EphemeralContainerStatuses` is never truncated. If a debug container has run in
this pod it will appear here.

**Support arbitrary runtimes via the CRI**. This proposal is implemented
entirely in the kubelet runtime manager and requires no changes in the
individual runtimes.

**Have an excellent user experience**. This solution is conceptually
straightforward and surfaced in a single `kubectl` command that "runs a thing in
a pod". Debug tools are distributed by container image, which is already well
understood by users. There is no automatic copying of files or hidden paths.

By using container images, users are empowered to create custom debug images.
Available images can be restricted by admission policy. Some examples of
possible debug images:

*   A script that automatically gathers a debugging snapshot and uploads it to a
    cloud storage bucket before killing the pod.
*   An image with a shell modified to log every statement to an audit API.

**Require no direct access to the node.** This solution uses the standard
streaming API.

**Have no inherent side effects to the running container image.** The target pod
is not modified by default, but resources used by the debug container will be
billed to the pod's cgroup, which means it could be evicted. A future
improvement could be to decrease the likelihood of eviction when there's an
active debug container.

### Appendix 3: Alternatives Considered

#### Container Spec in PodStatus

Originally there was a desire to keep the pod spec immutable, so we explored
modifying only the pod status. An `EphemeralContainer` would contain a Spec, a
Status and a Target:

```
// EphemeralContainer describes a container to attach to a running pod for troubleshooting.
type EphemeralContainer struct {
        metav1.TypeMeta `json:",inline"`

        // Spec describes the Ephemeral Container to be created.
        Spec *Container `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

        // Most recently observed status of the container.
        // This data may not be up to date.
        // Populated by the system.
        // Read-only.
        // +optional
        Status *ContainerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`

        // If set, the name of the container from PodSpec that this ephemeral container targets.
        // If not set then the ephemeral container is run in whatever namespaces are shared
        // for the pod.
        TargetContainerName string `json:"targetContainerName,omitempty" protobuf:"bytes,4,opt,name=targetContainerName"`
}
```

Ephemeral Containers for a pod would be listed in the pod's status:

```
type PodStatus struct {
        ...
        // List of user-initiated ephemeral containers that have been run in this pod.
        // +optional
        EphemeralContainers []EphemeralContainer `json:"ephemeralContainers,omitempty" protobuf:"bytes,11,rep,name=ephemeralContainers"`

}
```

To create a new Ephemeral Container, one would append a new `EphemeralContainer`
with the desired `v1.Container` as `Spec` in `Pod.Status` and updates the `Pod`
in the API. Users cannot normally modify the pod status, so we'd create a new
subresource `/ephemeralcontainers` that allows an update of solely
`EphemeralContainers` and enforces append-only semantics.

Since we have a requirement to describe the Ephemeral Container with a
`v1.Container`, this lead to a "spec in status" that seemed to violate API best
practices. It was confusing, and it required added complexity in the kubelet to
persist and publish user intent, which is rightfully the job of the apiserver.

#### Extend the Existing Exec API ("exec++")

A simpler change is to extend `v1.Pod`'s `/exec` subresource to support
"executing" container images. The current `/exec` endpoint must implement `GET`
to support streaming for all clients. We don't want to encode a (potentially
large) `v1.Container` into a query string, so we must extend `v1.PodExecOptions`
with the specific fields required for creating a Debug Container:

```
// PodExecOptions is the query options to a Pod's remote exec call
type PodExecOptions struct {
        ...
        // EphemeralContainerName is the name of an ephemeral container in which the
        // command ought to be run. Either both EphemeralContainerName and
        // EphemeralContainerImage fields must be set, or neither.
        EphemeralContainerName *string `json:"ephemeralContainerName,omitempty" ...`

        // EphemeralContainerImage is the image of an ephemeral container in which the command
        // ought to be run. Either both EphemeralContainerName and EphemeralContainerImage
        // fields must be set, or neither.
        EphemeralContainerImage *string `json:"ephemeralContainerImage,omitempty" ...`
}
```

After creating the Ephemeral Container, the kubelet would upgrade the connection
to streaming and perform an attach to the container's console. If disconnected,
the Ephemeral Container could be reattached using the pod's `/attach` endpoint
with `EphemeralContainerName`.

Ephemeral Containers could not be removed via the API and instead the process
must terminate. While not ideal, this parallels existing behavior of `kubectl
exec`. To kill an Ephemeral Container one would `attach` and exit the process
interactively or create a new Ephemeral Container to send a signal with
`kill(1)` to the original process.

Since the user cannot specify the `v1.Container`, this approach sacrifices a
great deal of flexibility. This solution still requires the kubelet to publish a
`Container` spec in the `PodStatus` that can be examined for future admission
decisions and so retains many of the downsides of the Container Spec in
PodStatus approach.

#### Ephemeral Container Controller

Kubernetes prefers declarative APIs where the client declares a state for
Kubernetes to enact. We could implement this in a declarative manner by creating
a new `EphemeralContainer` type:

```
type EphemeralContainer struct {
        metav1.TypeMeta
        metav1.ObjectMeta

        Spec v1.Container
        Status v1.ContainerStatus
}
```

A new controller in the kubelet would watch for EphemeralContainers and
create/delete debug containers. `EphemeralContainer.Status` would be updated by
the kubelet at the same time it updates `ContainerStatus` for regular and init
containers. Clients would create a new `EphemeralContainer` object, wait for it
to be started and then attach using the pod's attach subresource and the name of
the `EphemeralContainer`.

A new controller is a significant amount of complexity to add to the kubelet,
especially considering that the kubelet is already watching for changes to pods.
The kubelet would have to be modified to create containers in a pod from
multiple config sources. SIG Node strongly prefers to minimize kubelet
complexity.

#### Mutable Pod Spec Containers

Rather than adding to the pod API, we could instead make the pod spec mutable so
the client can generate an update adding a container. `SyncPod()` has no issues
adding the container to the pod at that point, but an immutable pod spec has
been a basic assumption and best practice in Kubernetes. Changing this
assumption complicates the requirements of the kubelet state machine. Since the
kubelet was not written with this in mind, we should expect such a change would
create bugs we cannot predict.

#### Image Exec

An earlier version of this proposal suggested simply adding `Image` parameter to
the exec API. This would run an ephemeral container in the pod namespaces
without adding it to the pod spec or status. This container would exist only as
long as the process it ran. This parallels the current kubectl exec, including
its lack of transparency. We could add constructs to track and report on both
traditional exec process and exec containers. In the end this failed to meet our
transparency requirements.

#### Attaching Container Type Volume

Combining container volumes ([#831](https://issues.k8s.io/831)) with the ability
to add volumes to the pod spec would get us most of the way there. One could
mount a volume of debug utilities at debug time. Docker does not allow adding a
volume to a running container, however, so this would require a container
restart. A restart doesn't meet our requirements for troubleshooting.

Rather than attaching the container at debug time, kubernetes could always
attach a volume at a random path at run time, just in case it's needed. Though
this simplifies the solution by working within the existing constraints of
`kubectl exec`, it has a sufficient list of minor limitations (detailed in
[#10834](https://issues.k8s.io/10834)) to result in a poor user experience.

#### Inactive container

If Kubernetes supported the concept of an "inactive" container, we could
configure it as part of a pod and activate it at debug time. In order to avoid
coupling the debug tool versions with those of the running containers, we would
want to ensure the debug image was pulled at debug time. The container could
then be run with a TTY and attached using kubectl.

The downside of this approach is that it requires prior configuration. In
addition to requiring prior consideration, it would increase boilerplate config.
A requirement for prior configuration makes it feel like a workaround rather
than a feature of the platform.

#### Implicit Empty Volume

Kubernetes could implicitly create an EmptyDir volume for every pod which would
then be available as a target for either the kubelet or a sidecar to extract a
package of binaries.

Users would have to be responsible for hosting a package build and distribution
infrastructure or rely on a public one. The complexity of this solution makes it
undesirable.

#### Standalone Pod in Shared Namespace ("Debug Pod")

Rather than inserting a new container into a pod namespace, Kubernetes could
instead support creating a new pod with container namespaces shared with
another, target pod. This would be a simpler change to the Kubernetes API, which
would only need a new field in the pod spec to specify the target pod. To be
useful, the containers in this "Debug Pod" should be run inside the namespaces
(network, pid, etc) of the target pod but remain in a separate resource group
(e.g. cgroup for container-based runtimes).

This would be a rather large change for pod, which is currently treated as an
atomic unit. The Container Runtime Interface has no provisions for sharing
outside of a pod sandbox and would need a refactor. This could be a complicated
change for non-container runtimes (e.g. hypervisor runtimes) which have more
rigid boundaries between pods.

This is pushing the complexity of the solution from the kubelet to the runtimes.
Minimizing change to the Kubernetes API is not worth the increased complexity
for the kubelet and runtimes.

It could also be possible to implement a Debug Pod as a privileged pod that runs
in the host namespace and interacts with the runtime directly to run a new
container in the appropriate namespace. This solution would be runtime-specific
and pushes the complexity of debugging to the user. Additionally, requiring
node-level access to debug a pod does not meet our requirements.

#### Exec from Node

The kubelet could support executing a troubleshooting binary from the node in
the namespaces of the container. Once executed this binary would lose access to
other binaries from the node, making it of limited utility and a confusing user
experience.

This couples the debug tools with the lifecycle of the node, which is worse than
coupling it with container images.

## Reference

*   [Pod Troubleshooting Tracking Issue](https://issues.k8s.io/27140)
*   [CRI Tracking Issue](https://issues.k8s.io/28789)
*   [CRI: expose optional runtime features](https://issues.k8s.io/32803)
*   [Resource QoS in Kubernetes](resource-qos.md)
*   Related Features
    *   [#1615](https://issues.k8s.io/1615) - Shared PID Namespace across
        containers in a pod
    *   [#26751](https://issues.k8s.io/26751) - Pod-Level cgroup
    *   [#10782](https://issues.k8s.io/10782) - Vertical pod autoscaling
