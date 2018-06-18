# User Capabilities

**Authors**:

1. Filipe Brandenburger (@filbranden)

**Last Updated**: 2018-06-18

**Status**: Draft

This document proposes the introduction of a new "userCapabilities" setting in
Kubernetes and the CRI, in order to support augmented capabilities for non-root
users in containers.

## Introduction

Capabilities have been used in OCI containers to weaken root inside the
container, preventing operations (such as creating device nodes, etc.) deemed to
counter other security measures, or have other security implications.

Capabilities can also be useful when applied to non-root users, in order to
empower them to perform selective actions usually reserved to root. A great
example of where this is useful is having a non-root user running nginx in a
container listening on port 80 (which is typically reserved for root and
requires a non-root user to have the `CAP_NET_BIND_SERVICE` capability to be
able to perform bind on that port.) (See also kubernetes/kubernetes#56374.)

This document proposes introducing a new "userCapabilities" setting to be able
to add capabilities to a non-root user running inside the container.

## Why not reuse "capabilities"?

The main problem with reusing the existing setting (as suggested or implied in
kubernetes/kubernetes#56374), is that it might be unclear at the time of
configuration whether the workload in the container will run as root or
non-root.

For instance, the user to run an application inside the container might be set
in an [`USER` directive of a
Dockerfile](https://docs.docker.com/v17.09/engine/reference/builder/#user), in
which case the operator writing the Kubernetes config might think they are
configuring capabilities for root, when in fact they are configuring them for
non-root, and might inadvertently give them more access than required,
effectively making that non-root user behave the same as root inside the
container.

Such cases might arise when the full set of capabilities is redefined in the
config, for example:

```yaml
securityContext:
    capabilities:
        drop:
        - all
        add:
        - CAP_CHOWN
        - CAP_DAC_OVERRIDE
        - CAP_FOWNER
        - CAP_FSETID
        - CAP_KILL
        - CAP_SETGID
        - CAP_SETUID
        - CAP_SETPCAP
        - CAP_NET_BIND_SERVICE
        - CAP_NET_RAW
        - CAP_SYS_CHROOT
        - CAP_MKNOD
        - CAP_AUDIT_WRITE
        - CAP_SETFCAP
```

This defines the capability set to the OCI default for root (for example, to
avoid any changes if the OCI default set is redefined, or perhaps to work around
differences in how CRI implementations handle capabilities.)

If such a setting were to be reused for non-root, suddenly the non-root user
would become as powerful as root inside the container.

Having two separate settings (one for root and one for non-root) also help
clarify what should happen when the container is entered more than once using
different users (for example, using `kubelet exec` into an existing container.)

It also helps clarify what happens if a non-root user should execute a setuid
binary inside the container (in which case, the existing "capabilities" setting
can still be used to bound how root inside the container should behave.)

For all these reasons, I propose that introducing a new, separate
"userCapabilities" setting is the best approach.

## Background

### Ambient Capabitilies

Until recently, Linux did not have a good API for setting capabilities and
preserving them as non-root users. When running as non-root, all capabilities
are by default dropped whenever a command is executed using the `execve()`
family of syscalls.

While the API includes a set of "inheritable" capabilities, those will only be
preserved when the binaries executed are marked with specific corresponding
capabilities bits, which, in practice, never actually happens (much less so in
the context of OCI containers, where there is not even a mechanism to apply or
preserve those bits in binaries inside the containers.) As a result,
"inheritable" capabilities do not work in practice, and there was no good way to
set capabilities for non-root users.

The newly introduced "ambient" capabilities solve that problem, by copying those
to the "permitted" and "effective" sets across `execve()` by non-root users,
regardless of the binaries being executed having any special attributes set.

This feature that makes it possible to implement "userCapabilities" is available
starting with Linux kernel 4.3.

### Capabilities support in runc and libcontainer, history in Docker

While runc tries to preserve capabilities across changing uids (so, in theory
making it possible to set capabilities for non-root users), the history lack of
support for "ambient" capabilities made it such that somewhat akin to
"userCapabilities" was never really supported.

libcontainer has code to [preserve capabilities while changing
users](https://github.com/opencontainers/runc/blob/v1.0.0-rc5/libcontainer/init_linux.go#L141).
Historically, this only included "inheritable" capabilities (but not "ambient"
capabilities), therefore all those settings would essentially just go away
whenever `execve()` was called inside the container. The end effect was that,
even though libcontainer took steps to preserve the capabilities, they were
effectively dropped by the time a binary was executed as a non-root user inside
the container.

libcontainer also takes [separate capability masks for each capability
set](https://github.com/opencontainers/runc/blob/v1.0.0-rc5/libcontainer/configs/config.go#L208),
including "ambient" capabilities.  We can now leverage that existing support to
implement "userCapabilities" as described here.

## Design

For sake of clarity, let's use an example to describe how each capability mask
should be set in order to implement root and non-root capabilities. In this
example, we'll use the [standard set of
capabilities](https://github.com/opencontainers/runc/blob/v1.0.0-rc5/libcontainer/SPEC.md#security)
for the root capabilities mask, which encodes to `00000000a80425fb` in hex, and
we'll use a capabilities mask that only sets `CAP_NET_BIND_SERVICE` for the
non-root user, which encodes to `0000000000000400` in hex.

### root capabilities

In order to get the capabilities masks correctly set for the root user, we need
to set them this way:

```
CapInh: xxxxxxxxxxxxxxxx  <- unimportant, in general
CapPrm: 00000000a80425fb
CapEff: 00000000a80425fb  <- the essential setting
CapBnd: 00000000a80425fb
CapAmb: xxxxxxxxxxxxxxxx  <- unimportant for root
```

The key here is to set the "effective" mask to include the capabilities root
inside the container should have.

Furthermore, the "permitted" and "bounding" sets should use the same
capabilities, to prevent root inside the container from gaining more
capabilities by executing a setuid binary or a binary with file capabilities.

The "inheritable" capabilities are unimportant, since they only work with file
capabilities (which we can effectively ignore in general, and especially inside
containers.)

The "ambient" capabilities are mostly ignored when running as root, so also not
very important in this context.

### user capabilities

To get non-root capabilities set correctly, we should set them this way:

```
CapInh: xxxxxxxxxxxxxxxx  <- unimportant, in general
CapPrm: 00000000a80425fb  <- important for setuid binaries
CapEff: 0000000000000400  <- makes a difference before execve(), but not essential
CapBnd: 00000000a80425fb
CapAmb: 0000000000000400  <- the essential setting
```

Here, the key is to set the "ambient" capabilities to the ones the non-root user
wants. Whenever `execve()` is called, the "ambient" capabilities will be copied
to the "effective" set, making those capabilities take effect.

The "permitted" set is also important here, since it is used whenever a setuid
binary is executed by the non-root user inside the container. We still want a
limited root in container in that case, so setting the "permitted" capabilities
correctly ensures this will be the case.

The "effective" set should ideally be set to the more restrictive permissions of
the non-root user. If runc/libcontainer sets them to the root capabilities
(`00000000a80425fb`), then in effect runc/libcontainer will be running as
non-root but with the *same capabilities as root has*.

That is not really that big of a problem, since as soon as `execve()` is called
(and one will be called to execute a binary inside the container), the "ambient"
capabilities will prevail. So this only happens during the time when runc is
running after switching to the non-root user, but before it executes a file
inside the container, which is fairly short.

It might be possible to update runc/libcontainer to fix that, by checking
whether it switched id to non-root and then masking the "effective" capabilities
to be applied by ANDing it with the "ambient" capabilities.

### Putting both together

Given that root capabilities control what the "effective", "permitted" and
"bounding" sets should be set to and that the user capabilities control what the
"ambient" capabilities should be set to, we can converge to a single setting
that will work for both cases.

The "inheritable" capabilities don't really matter much in this context, but
setting them to match the root capabilities is probably fine, since they're only
going in effect when executing a file with file capabilities and that's akin to
a setuid binary, thus using a similar upper bound is OK in that case.

So the full setting should be:

```
CapInh: 00000000a80425fb
CapPrm: 00000000a80425fb  <- determines how setuid binaries act
CapEff: 00000000a80425fb  <- capabilities for root inside the container
CapBnd: 00000000a80425fb
CapAmb: 0000000000000400  <- capabilities for non-root inside the container
```

As explained above, setting the "effective" capabilities to the root
capabilities will make runc/libcontainer still run with privileges even after
switching to a non-root user, but that situation will be fixed as soon as it
executes a binary inside the container.

## Implementation

We will extend the CRI protocol to include two sets of capabilities, the current
`Capabilities` and a new field `UserCapabilities`. Both of these will be
available when configuring the "securityContext:" of a pod or container.

The fields will be passed on to the Runtimes through the CRI, where they will
eventually be decoded into the sets of capabilities (including "ambient"
capabilities) to be passed to runc/libcontainer.

Most of the code is plumbing to expose this new field in the configs, then pass
it to the Runtime through the CRI, then in the Runtime implementations, using
it to populate the "ambient" capabilities.

### Kubernetes

An early WIP for adding the field to Kubernetes can be found
[here](https://github.com/filbranden/kubernetes/commit/e8561087343c81478221a4cd6f8a9cc7e17cf502).
It still needs more checking around the values that can be set here (as it's
currently done for the root capabilities.)

### containerd/cri

For containerd/cri, a first step is to update to latest
opencontainers/runtime-tools which gives more granular access to setting each
capability set individually. PR containerd/cri#820 covers that.

A follow up is to take the newly passed user capabilities and use them to
populate the "ambient" capabilities, which can be done
[here](https://github.com/containerd/cri/blob/v1.0.3/pkg/server/container_create.go#L375).

### CRI-O

Similar to containerd, "ambient" capabilities are currently being cleared in
CRI-O, so it's also clear [where in the
code](https://github.com/kubernetes-incubator/cri-o/blob/v1.10.3/server/container_create.go#L502)
we should populate those using the new "userCapabilities".

### User Interface

The user will have access to this new feature through a new field in their
"securityContext:" config for a container.

As an example of a container running netcat as user "nobody" with
`CAP_NET_BIND_SERVICE` to be able to bind to port 80, while also adding
`CAP_SYS_NICE` to root inside container (in case someone uses `kubectl exec` on
it to renice a process):

```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: netcat
    image: alpine
    args:
    - /bin/sh
    - -c
    - "nc -lk -p 80 -e echo hello"
    securityContext:
      runAsUser: nobody
      userCapabilities:
        add:
        - NET_BIND_SERVICE
      capabilities:
        add:
        - CAP_SYS_NICE
```

Note that adding "all" or "ALL" to "userCapabilities:" will have no effect,
since doing otherwise would just amount to turning non-root inside container
into real root. That's too dangerous, so let's just avoid it. Technically, it's
possible to get a non-root user get all capabilities by listing them all
explicitly, but at least that's explicit, so it's easier to catch.

### Compatibility

With the fields being passed as protobufs using gRPC, the absence of support for
the field on either side just reverts to the default behavior, which is ignoring
that user capabilities exist and keeping the "ambient" capabilities set to none.

In effect, if either side lacks support, user capabilities will simply be
silently dropped.

This is the most secure setting and should not be too hard to troubleshoot given
the failure scenario is likely to make the container fail quickly with a message
that is likely to point to the lack of specific capabilities.
