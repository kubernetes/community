# ProcMount/ProcMountType Option

## Background

Currently the way docker and most other container runtimes work is by masking
and setting as read-only certain paths in `/proc`. This is to prevent data
from being exposed into a container that should not be. However, there are
certain use-cases where it is necessary to turn this off.

## Motivation

For end-users who would like to run unprivileged containers using user namespaces
_nested inside_ CRI containers, we need an option to have a `ProcMount`. That is,
we need an option to designate explicitly turn off masking and setting 
read-only of paths so that we can
mount `/proc` in the nested container as an unprivileged user.

Please see the following filed issues for more information:
- [opencontainers/runc#1658](https://github.com/opencontainers/runc/issues/1658#issuecomment-373122073)
- [moby/moby#36597](https://github.com/moby/moby/issues/36597)
- [moby/moby#36644](https://github.com/moby/moby/pull/36644)

Please also see the [use case for building images securely in kubernetes](https://github.com/jessfraz/blog/blob/master/content/post/building-container-images-securely-on-kubernetes.md).

Unmasking the paths in `/proc` option really only makes sense for when a user 
is nesting
unprivileged containers with user namespaces as it will allow more information
than is necessary to the program running in the container spawned by
kubernetes.

The main use case for this option is to run
[genuinetools/img](https://github.com/genuinetools/img) inside a kubernetes
container. That program then launches sub-containers that take advantage of
user namespaces and re-mask /proc and set /proc as read-only. So therefore
there is no concern with having an unmasked proc open in the top level container.

It should be noted that this is different that the host /proc. It is still
a newly mounted /proc just the container runtimes will not mask the paths.

Since the only use case for this option is to run unprivileged nested
containers,
this option should only be allowed or used if the user in the container is not `root`.
This can be easily enforced with `MustRunAs`.
Since the user inside is still unprivileged,
doing things to `/proc` would be off limits regardless, since linux user
support already prevents this.

## Existing SecurityContext objects

Kubernetes defines `SecurityContext` for `Container` and `PodSecurityContext`
for `PodSpec`. `SecurityContext` objects define the related security options
for Kubernetes containers, e.g. selinux options.

To support "ProcMount" options in Kubernetes, it is proposed to make
the following changes:

## Changes of SecurityContext objects

Add a new `string` type field named `ProcMountType` will hold the viable
options for `procMount` to the `SecurityContext`
definition.

By default,`procMount` is `default`, aka the same behavior as today and the
paths are masked.

This will look like the following in the spec:

```go
type ProcMountType string

const (
    // DefaultProcMount uses the container runtime default ProcType.  Most 
    // container runtimes mask certain paths in /proc to avoid accidental security
    // exposure of special devices or information.
    DefaultProcMount ProcMountType = "Default"

    // UnmaskedProcMount bypasses the default masking behavior of the container
    // runtime and ensures the newly created /proc the container stays in tact with
    // no modifications.  
    UnmaskedProcMount ProcMountType = "Unmasked"
)

procMount *ProcMountType
```

This requires changes to the CRI runtime integrations so that
kubelet will add the specific `unmasked` or `whatever_it_is_named` option.

## Pod Security Policy changes

A new `[]ProcMountType{}` field named `allowedProcMounts` will be added to the Pod
Security Policy as well to gate the allowed ProcMountTypes a user is allowed to
set. This field will default to `[]ProcMountType{ DefaultProcMount }`.
