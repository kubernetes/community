# No New Privileges

- [Description](#description)
  * [Interactions with other Linux primitives](#interactions-with-other-linux-primitives)
- [Current Implementations](#current-implementations)
  * [Support in Docker](#support-in-docker)
  * [Support in rkt](#support-in-rkt)
  * [Support in OCI runtimes](#support-in-oci-runtimes)
- [Existing SecurityContext objects](#existing-securitycontext-objects)
- [Changes of SecurityContext objects](#changes-of-securitycontext-objects)
- [Pod Security Policy changes](#pod-security-policy-changes)


## Description

In Linux, the `execve` system call can grant more privileges to a newly-created
process than its parent process. Considering security issues, since Linux kernel
v3.5, there is a new flag named `no_new_privs` added to prevent those new
privileges from being granted to the processes.

[`no_new_privs`](https://www.kernel.org/doc/Documentation/prctl/no_new_privs.txt)
is inherited across `fork`, `clone` and `execve` and can not be unset. With
`no_new_privs` set, `execve` promises not to grant the privilege to do anything
that could not have been done without the `execve` call.

For more details about `no_new_privs`, please check the
[Linux kernel documentation](https://www.kernel.org/doc/Documentation/prctl/no_new_privs.txt).

This is different from `NOSUID` in that `no_new_privs`can give permission to
the container process to further restrict child processes with seccomp. This
permission goes only one-way in that the container process can not grant more
permissions, only further restrict.

### Interactions with other Linux primitives

- suid binaries: will break when `no_new_privs` is enabled
- seccomp2 as a non root user: requires `no_new_privs`
- seccomp2 with dropped `CAP_SYS_ADMIN`: requires `no_new_privs`
- ambient capabilities: requires `no_new_privs`
- selinux transitions: bugs that were fixed documented [here](https://github.com/moby/moby/issues/23981#issuecomment-233121969)


## Current Implementations

### Support in Docker

Since Docker 1.11, a user can specify `--security-opt` to enable `no_new_privs`
while creating containers, for example
`docker run --security-opt=no_new_privs busybox`.

Docker provides via their Go api an object named `ContainerCreateConfig` to
configure container creation parameters. In this object, there is a string
array `HostConfig.SecurityOpt` to specify the security options. Client can
utilize this field to specify the arguments for security options while
creating new containers.

This field did not scale well for the Docker client, so it's suggested that
Kubernetes does not follow that design.

This is not on by default in Docker.

More details of the Docker implementation can be read
[here](https://github.com/moby/moby/pull/20727) as well as the original
discussion [here](https://github.com/moby/moby/issues/20329).

### Support in rkt

Since rkt v1.26.0, the `NoNewPrivileges` option has been enabled in rkt.

More details of the rkt implementation can be read
[here](https://github.com/rkt/rkt/pull/2677).

### Support in OCI runtimes

Since version 0.3.0 of the OCI runtime specification, a user can specify the
`noNewPrivs` boolean flag in the configuration file.

More details of the OCI implementation can be read
[here](https://github.com/opencontainers/runtime-spec/pull/290).

## Existing SecurityContext objects

Kubernetes defines `SecurityContext` for `Container` and `PodSecurityContext`
for `PodSpec`. `SecurityContext` objects define the related security options
for Kubernetes containers, e.g. selinux options.

To support "no new privileges" options in Kubernetes, it is proposed to make
the following changes:

## Changes of SecurityContext objects

Add a new `*bool` type field named `allowPrivilegeEscalation` to the `SecurityContext`
definition.

By default, ie when `allowPrivilegeEscalation=nil`, we will set `no_new_privs=true`
with the following exceptions:

- when a container is `privileged`
- when `CAP_SYS_ADMIN` is added to a container
- when a container is not run as root, uid `0` (to prevent breaking suid
  binaries)

The API will reject as invalid `privileged=true` and
`allowPrivilegeEscalation=false`, as well as `capAdd=CAP_SYS_ADMIN` and
`allowPrivilegeEscalation=false.`

When `allowPrivilegeEscalation` is set to `false` it will enable `no_new_privs`
for that container.

`allowPrivilegeEscalation` in `SecurityContext` provides container level
control of the `no_new_privs` flag and can override the default in both directions
of the `allowPrivilegeEscalation` setting.

This requires changes to the Docker, rkt, and CRI runtime integrations so that
kubelet will add the specific `no_new_privs` option.

## Pod Security Policy changes

The default can be set via a new `*bool` type field named `defaultAllowPrivilegeEscalation`
in a Pod Security Policy.
This would allow users to set `defaultAllowPrivilegeEscalation=false`, overriding the
default `nil` behavior of `no_new_privs=false` for containers
whose uids are not 0.

This would also keep the behavior of setting the security context as
`allowPrivilegeEscalation=true`
for privileged containers and those with `capAdd=CAP_SYS_ADMIN`.

To recap, below is a table defining the default behavior at the pod security
policy level and what can be set as a default with a pod security policy.

| allowPrivilegeEscalation setting | uid = 0 or unset   | uid != 0           | privileged/CAP_SYS_ADMIN |
|----------------------------------|--------------------|--------------------|--------------------------|
|  nil                             | no_new_privs=true  | no_new_privs=false | no_new_privs=false       |
|  false                           | no_new_privs=true  | no_new_privs=true  | no_new_privs=false       |
|  true                            | no_new_privs=false | no_new_privs=false | no_new_privs=false       |

A new `bool` field named `allowPrivilegeEscalation` will be added to the Pod
Security Policy as well to gate whether or not a user is allowed to set the
security context to `allowPrivilegeEscalation=true`. This field will default to
false.
