# Seccomp

- [Abstract](#abstract)
- [Motivation](#motivation)
- [Constraints and Assumptions](#constraints-and-assumptions)
- [Use Cases](#use-cases)
  * [Use Case: Administrator Access Control](#use-case--administrator-access-control)
  * [Use Case: Seccomp profiles similar to container runtime defaults](#use-case--seccomp-profiles-similar-to-container-runtime-defaults)
  * [Use Case: Applications that link to libseccomp](#use-case--applications-that-link-to-libseccomp)
  * [Use Case: Custom profiles](#use-case--custom-profiles)
- [Community Work](#community-work)
  * [Docker / OCI](#docker---oci)
  * [rkt / appcontainers](#rkt---appcontainers)
  * [HyperContainer](#hypercontainer)
  * [lxd](#lxd)
  * [Other platforms and seccomp-like capabilities](#other-platforms-and-seccomp-like-capabilities)
- [Proposed Design](#proposed-design)
  * [Seccomp API Resource?](#seccomp-api-resource-)
  * [Pod Security Policy annotation](#pod-security-policy-annotation)
  * [Spec](#spec)
  * [Default Profile](#default-profile)
    + [Various Syscalls Not Allowed](#various-syscalls-not-allowed)
    + [Default Behavior](#default-behavior)
- [Examples](#examples)
  * [Unconfined profile](#unconfined-profile)
  * [Custom profile](#custom-profile)

## Abstract

A proposal for adding **alpha** support for
[seccomp](https://github.com/seccomp/libseccomp) to Kubernetes.  Seccomp is a
system call filtering facility in the Linux kernel which lets applications
define limits on system calls they may make, and what should happen when
system calls are made.  Seccomp is used to reduce the attack surface available
to applications.

## Motivation

Applications use seccomp to restrict the set of system calls they can make.
Recently, container runtimes have begun adding features to allow the runtime
to interact with seccomp on behalf of the application, which eliminates the
need for applications to link against libseccomp directly.  Adding support in
the Kubernetes API for describing seccomp profiles will allow administrators
greater control over the security of workloads running in Kubernetes.

Goals of this design:

1.  Describe how to reference seccomp profiles in containers that use them

## Constraints and Assumptions

This design should:

*  build upon previous security context work
*  be container-runtime agnostic
*  allow use of custom profiles
*  facilitate containerized applications that link directly to libseccomp
*  enable a default seccomp profile for containers

## Use Cases

1.  As an administrator, I want to be able to grant access to a seccomp profile
    to a class of users
2.  As a user, I want to run an application with a seccomp profile similar to
    the default one provided by my container runtime
3.  As a user, I want to run an application which is already libseccomp-aware
    in a container, and for my application to manage interacting with seccomp
    unmediated by Kubernetes
4.  As a user, I want to be able to use a custom seccomp profile and use
    it with my containers
5.  As a user and administrator I want kubernetes to apply a sane default
    seccomp profile to containers unless I otherwise specify.

### Use Case: Administrator Access Control

Controlling access to seccomp profiles is a cluster administrator
concern. It should be possible for an administrator to control which users
have access to which profiles.

The [Pod Security Policy](https://github.com/kubernetes/kubernetes/pull/7893)
API extension governs the ability of users to make requests that affect pod
and container security contexts.  The proposed design should deal with
required changes to control access to new functionality.

### Use Case: Seccomp profiles similar to container runtime defaults

Many users will want to use images that make assumptions about running in the
context of their chosen container runtime.  Such images are likely to
frequently assume that they are running in the context of the container
runtime's default seccomp settings.  Therefore, it should be possible to
express a seccomp profile similar to a container runtime's defaults.

As an example, all dockerhub 'official' images are compatible with the Docker
default seccomp profile.  So, any user who wanted to run one of these images
with seccomp would want the default profile to be accessible.

### Use Case: Applications that link to libseccomp

Some applications already link to libseccomp and control seccomp directly.  It
should be possible to run these applications unmodified in Kubernetes; this
implies there should be a way to disable seccomp control in Kubernetes for
certain containers, or to run with a "no-op" or "unconfined" profile.

Sometimes, applications that link to seccomp can use the default profile for a
container runtime, and restrict further on top of that.  It is important to
note here that in this case, applications can only place _further_
restrictions on themselves.  It is not possible to re-grant the ability of a
process to make a system call once it has been removed with seccomp.

As an example, elasticsearch manages its own seccomp filters in its code.
Currently, elasticsearch is capable of running in the context of the default
Docker profile, but if in the future, elasticsearch needed to be able to call
`ioperm` or `iopr` (both of which are disallowed in the default profile), it
should be possible to run elasticsearch by delegating the seccomp controls to
the pod.

### Use Case: Custom profiles

Different applications have different requirements for seccomp profiles; it
should be possible to specify an arbitrary seccomp profile and use it in a
container.  This is more of a concern for applications which need a higher
level of privilege than what is granted by the default profile for a cluster,
since applications that want to restrict privileges further can always make
additional calls in their own code.

An example of an application that requires the use of a syscall disallowed in
the Docker default profile is Chrome, which needs `clone` to create a new user
namespace.  Another example would be a program which uses `ptrace` to
implement a sandbox for user-provided code, such as
[eval.in](https://eval.in/).

## Community Work

### Docker / OCI

Docker supports the open container initiative's API for
seccomp, which is very close to the libseccomp API.  It allows full
specification of seccomp filters, with arguments, operators, and actions.

Docker allows the specification of a single seccomp filter.  There are
community requests for:

* [docker/22109](https://github.com/docker/docker/issues/22109): composable
  seccomp filters
* [docker/21105](https://github.com/docker/docker/issues/22105): custom
  seccomp filters for builds

Implementation details:

* [docker/17989](https://github.com/moby/moby/pull/17989): initial
  implementation
* [docker/18780](https://github.com/moby/moby/pull/18780): default blacklist
  profile
* [docker/18979](https://github.com/moby/moby/pull/18979): default whitelist
  profile

### rkt / appcontainers

The `rkt` runtime delegates to systemd for seccomp support; there is an open
issue to add support once `appc` supports it.  The `appc` project has an open
issue to be able to describe seccomp as an isolator in an appc pod.

The systemd seccomp facility is based on a whitelist of system calls that can
be made, rather than a full filter specification.

Issues:

* [appc/529](https://github.com/appc/spec/issues/529)
* [rkt/1614](https://github.com/coreos/rkt/issues/1614)

### HyperContainer

[HyperContainer](https://hypercontainer.io) does not support seccomp.

### lxd

[`lxd`](http://www.ubuntu.com/cloud/lxd) constrains containers using a default profile.

Issues:

* [lxd/1084](https://github.com/lxc/lxd/issues/1084): add knobs for seccomp

### Other platforms and seccomp-like capabilities

FreeBSD has a seccomp/capability-like facility called
[Capsicum](https://www.freebsd.org/cgi/man.cgi?query=capsicum&sektion=4).


## Proposed Design

### Seccomp API Resource?

An earlier draft of this proposal described a new global API resource that
could be used to describe seccomp profiles.  After some discussion, it was
determined that without a feedback signal from users indicating a need to
describe new profiles in the Kubernetes API, it is not possible to know
whether a new API resource is warranted.

That being the case, we will not propose a new API resource at this time.  If
there is strong community desire for such a resource, we may consider it in
the future.

Instead of implementing a new API resource, we propose that pods be able to
reference seccomp profiles by name.  Since this is an alpha feature, we will
use annotations instead of extending the API with new fields.

In the alpha version of this feature we will use annotations to store the
names of seccomp profiles.  The keys will be:

`container.seccomp.security.alpha.kubernetes.io/<container name>`

which will be used to set the seccomp profile of a container, and:

`seccomp.security.alpha.kubernetes.io/pod`

which will set the seccomp profile for the containers of an entire pod.  If a
pod-level annotation is present, and a container-level annotation present for
a container, then the container-level profile takes precedence.

The value of these keys should be container-runtime agnostic. We will
establish a format that expresses the conventions for distinguishing between
an unconfined profile, the container runtime's default, or a custom profile.
Since format of profile is likely to be runtime dependent, we will consider
profiles to be opaque to kubernetes for now.

The following format is scoped as follows:

1.  `docker/default` - the default profile for the container runtime
2.  `unconfined` - unconfined profile, ie, no seccomp sandboxing
3.  `localhost/<profile-name>` - the profile installed to the node's local seccomp profile root

Since seccomp profile schemes may vary between container runtimes, we will
treat the contents of profiles as opaque for now and avoid attempting to find
a common way to describe them.  It is up to the container runtime to be
sensitive to the annotations proposed here and to interpret instructions about
local profiles.

A new area on disk (which we will call the seccomp profile root) must be
established to hold seccomp profiles.  A field will be added to the Kubelet
for the seccomp profile root and a knob (`--seccomp-profile-root`) exposed to
allow admins to set it. If unset, it should default to the `seccomp`
subdirectory of the kubelet root directory.

### Pod Security Policy annotation

The `PodSecurityPolicy` type should be annotated with the allowed seccomp
profiles using the key
`seccomp.security.alpha.kubernetes.io/allowedProfileNames`.  The value of this
key should be a comma delimited list.

### Spec

We will start from the OCI specification. This API resource will be added to
`settings.k8s.io` as an `alpha` resource.

```
// Seccomp represents syscall restrictions
type Seccomp struct {
    unversioned.TypeMeta
    ObjectMeta

    // +optional
    Spec SeccompSpec
}

// SeccompSpec represents the spec for syscall restrictions
type SeccompSpec struct {
	DefaultAction Action        `json:"defaultAction"`
	Architectures []Arch        `json:"architectures,omitempty"`
	Syscalls      []Syscall     `json:"syscalls,omitempty"`
}

// Arch used for additional architectures
type Arch string

// Additional architectures permitted to be used for system calls
// By default only the native architecture of the kernel is permitted
const (
	ArchX86         Arch = "SCMP_ARCH_X86"
	ArchX86_64      Arch = "SCMP_ARCH_X86_64"
	ArchX32         Arch = "SCMP_ARCH_X32"
	ArchARM         Arch = "SCMP_ARCH_ARM"
	ArchAARCH64     Arch = "SCMP_ARCH_AARCH64"
	ArchMIPS        Arch = "SCMP_ARCH_MIPS"
	ArchMIPS64      Arch = "SCMP_ARCH_MIPS64"
	ArchMIPS64N32   Arch = "SCMP_ARCH_MIPS64N32"
	ArchMIPSEL      Arch = "SCMP_ARCH_MIPSEL"
	ArchMIPSEL64    Arch = "SCMP_ARCH_MIPSEL64"
	ArchMIPSEL64N32 Arch = "SCMP_ARCH_MIPSEL64N32"
	ArchPPC         Arch = "SCMP_ARCH_PPC"
	ArchPPC64       Arch = "SCMP_ARCH_PPC64"
	ArchPPC64LE     Arch = "SCMP_ARCH_PPC64LE"
	ArchS390        Arch = "SCMP_ARCH_S390"
	ArchS390X       Arch = "SCMP_ARCH_S390X"
	ArchPARISC      Arch = "SCMP_ARCH_PARISC"
	ArchPARISC64    Arch = "SCMP_ARCH_PARISC64"
)

// SeccompAction taken upon Seccomp rule match
type SeccompAction string

// Define actions for Seccomp rules
const (
	ActKill  SeccompAction = "SCMP_ACT_KILL"
	ActTrap  SeccompAction = "SCMP_ACT_TRAP"
	ActErrno SeccompAction = "SCMP_ACT_ERRNO"
	ActTrace SeccompAction = "SCMP_ACT_TRACE"
	ActAllow SeccompAction = "SCMP_ACT_ALLOW"
)

// SeccompOperator used to match syscall arguments in Seccomp
type SeccompOperator string

// Define operators for syscall arguments in Seccomp
const (
	OpNotEqual     SeccompOperator = "SCMP_CMP_NE"
	OpLessThan     SeccompOperator = "SCMP_CMP_LT"
	OpLessEqual    SeccompOperator = "SCMP_CMP_LE"
	OpEqualTo      SeccompOperator = "SCMP_CMP_EQ"
	OpGreaterEqual SeccompOperator = "SCMP_CMP_GE"
	OpGreaterThan  SeccompOperator = "SCMP_CMP_GT"
	OpMaskedEqual  SeccompOperator = "SCMP_CMP_MASKED_EQ"
)

// SeccompArg used for matching specific syscall arguments in Seccomp
type SeccompArg struct {
	Index    uint                 `json:"index"`
	Value    uint64               `json:"value"`
	ValueTwo uint64               `json:"valueTwo"`
	Op       SeccompOperator      `json:"op"`
}

// Syscall is used to match a syscall in Seccomp
type Syscall struct {
	Names  []string      `json:"names"`
	Action SeccompAction `json:"action"`
	Args   []SeccompArg  `json:"args,omitempty"`
}
```

### Default Profile

We will create our own default seccomp profile  that uses the above spec
for containers and use the set of syscalls from the docker default profile
as the initial base. Having our own will allow us to control and
restrict different syscalls in the future.

#### Various Syscalls Not Allowed

Below includes a table of some of the syscalls we will not allow in our
whitelist and why. It does not include all the syscalls but merely some
important ones. Most of this was taken from the
[original pull request](https://github.com/moby/moby/pull/19059) to Docker
for the default profile.

| Syscall             | Description                                                                                                                           |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| `acct`              | Accounting syscall which could let containers disable their own resource limits or process accounting. Also gated by `CAP_SYS_PACCT`. |
| `add_key`           | Prevent containers from using the kernel keyring, which is not namespaced.                                   |
| `adjtimex`          | Similar to `clock_settime` and `settimeofday`, time/date is not namespaced.  Also gated by `CAP_SYS_TIME`.   |
| `bpf`               | Deny loading potentially persistent bpf programs into kernel, already gated by `CAP_SYS_ADMIN`.              |
| `clock_adjtime`     | Time/date is not namespaced. Also gated by `CAP_SYS_TIME`.                                                   |
| `clock_settime`     | Time/date is not namespaced. Also gated by `CAP_SYS_TIME`.                                                   |
| `clone`             | Deny cloning new namespaces. Also gated by `CAP_SYS_ADMIN` for CLONE_* flags, except `CLONE_USERNS`.         |
| `create_module`     | Deny manipulation and functions on kernel modules. Obsolete. Also gated by `CAP_SYS_MODULE`.                 |
| `delete_module`     | Deny manipulation and functions on kernel modules. Also gated by `CAP_SYS_MODULE`.                           |
| `finit_module`      | Deny manipulation and functions on kernel modules. Also gated by `CAP_SYS_MODULE`.                           |
| `get_kernel_syms`   | Deny retrieval of exported kernel and module symbols. Obsolete.                                              |
| `get_mempolicy`     | Syscall that modifies kernel memory and NUMA settings. Already gated by `CAP_SYS_NICE`.                      |
| `init_module`       | Deny manipulation and functions on kernel modules. Also gated by `CAP_SYS_MODULE`.                           |
| `ioperm`            | Prevent containers from modifying kernel I/O privilege levels. Already gated by `CAP_SYS_RAWIO`.             |
| `iopl`              | Prevent containers from modifying kernel I/O privilege levels. Already gated by `CAP_SYS_RAWIO`.             |
| `kcmp`              | Restrict process inspection capabilities, already blocked by dropping `CAP_PTRACE`.                          |
| `kexec_file_load`   | Sister syscall of `kexec_load` that does the same thing, slightly different arguments. Also gated by `CAP_SYS_BOOT`. |
| `kexec_load`        | Deny loading a new kernel for later execution. Also gated by `CAP_SYS_BOOT`.                                 |
| `keyctl`            | Prevent containers from using the kernel keyring, which is not namespaced.                                   |
| `lookup_dcookie`    | Tracing/profiling syscall, which could leak a lot of information on the host. Also gated by `CAP_SYS_ADMIN`. |
| `mbind`             | Syscall that modifies kernel memory and NUMA settings. Already gated by `CAP_SYS_NICE`.                      |
| `mount`             | Deny mounting, already gated by `CAP_SYS_ADMIN`.                                                             |
| `move_pages`        | Syscall that modifies kernel memory and NUMA settings.                                                       |
| `name_to_handle_at` | Sister syscall to `open_by_handle_at`. Already gated by `CAP_SYS_NICE`.                                      |
| `nfsservctl`        | Deny interaction with the kernel nfs daemon. Obsolete since Linux 3.1.                                       |
| `open_by_handle_at` | Cause of an old container breakout. Also gated by `CAP_DAC_READ_SEARCH`.                                     |
| `perf_event_open`   | Tracing/profiling syscall, which could leak a lot of information on the host.                                |
| `personality`       | Prevent container from enabling BSD emulation. Not inherently dangerous, but poorly tested, potential for a lot of kernel vulns. |
| `pivot_root`        | Deny `pivot_root`, should be privileged operation.                                                           |
| `process_vm_readv`  | Restrict process inspection capabilities, already blocked by dropping `CAP_PTRACE`.                          |
| `process_vm_writev` | Restrict process inspection capabilities, already blocked by dropping `CAP_PTRACE`.                          |
| `ptrace`            | Tracing/profiling syscall, which could leak a lot of information on the host. Already blocked by dropping `CAP_PTRACE`. |
| `query_module`      | Deny manipulation and functions on kernel modules. Obsolete.                                                  |
| `quotactl`          | Quota syscall which could let containers disable their own resource limits or process accounting. Also gated by `CAP_SYS_ADMIN`. |
| `reboot`            | Don't let containers reboot the host. Also gated by `CAP_SYS_BOOT`.                                           |
| `request_key`       | Prevent containers from using the kernel keyring, which is not namespaced.                                    |
| `set_mempolicy`     | Syscall that modifies kernel memory and NUMA settings. Already gated by `CAP_SYS_NICE`.                       |
| `setns`             | Deny associating a thread with a namespace. Also gated by `CAP_SYS_ADMIN`.                                    |
| `settimeofday`      | Time/date is not namespaced. Also gated by `CAP_SYS_TIME`.
| `socket`, `socketcall` | Used to send or receive packets and for other socket operations. All `socket` and `socketcall` calls are blocked except communication domains `AF_UNIX`, `AF_INET`, `AF_INET6`, `AF_NETLINK`, and `AF_PACKET`. |
| `stime`             | Time/date is not namespaced. Also gated by `CAP_SYS_TIME`.                                                    |
| `swapon`            | Deny start/stop swapping to file/device. Also gated by `CAP_SYS_ADMIN`.                                       |
| `swapoff`           | Deny start/stop swapping to file/device. Also gated by `CAP_SYS_ADMIN`.                                       |
| `sysfs`             | Obsolete syscall.                                                                                             |
| `_sysctl`           | Obsolete, replaced by /proc/sys.                                                                              |
| `umount`            | Should be a privileged operation. Also gated by `CAP_SYS_ADMIN`.                                              |
| `umount2`           | Should be a privileged operation. Also gated by `CAP_SYS_ADMIN`.                                              |
| `unshare`           | Deny cloning new namespaces for processes. Also gated by `CAP_SYS_ADMIN`, with the exception of `unshare --user`. |
| `uselib`            | Older syscall related to shared libraries, unused for a long time.                                            |
| `userfaultfd`       | Userspace page fault handling, largely needed for process migration.                                          |
| `ustat`             | Obsolete syscall.                                                                                             |
| `vm86`              | In kernel x86 real mode virtual machine. Also gated by `CAP_SYS_ADMIN`.                                       |
| `vm86old`           | In kernel x86 real mode virtual machine. Also gated by `CAP_SYS_ADMIN`.                                       |

#### Default Behavior

For `privileged` containers, no default seccomp profile will be used unless
explicitly requested by the user via annotations.

If `capAdd` is used on a Container, the default profile will be adjusted to
interact accordingly with the capability added. These are documented below in
a table by the cap being added:

| Capability           | Syscalls Allowed                                             |
|----------------------|--------------------------------------------------------------|
| `CAP_CHOWN`          | chown, chown32, fchown, fchown32, fchownat, lchown, lchown32 |
| `CAP_DAC_READ_SEARCH`| open_by_handle_at                                            |
| `CAP_IPC_LOCK`       | mlock, mlock2, mlockall                                      |
| `CAP_SYS_ADMIN`      | name_to_handle_at, bpf, clone, fanotify_init, lookup_dcookie, mount, perf_event_open, setdomainname, sethostname, setns, umount, umount2, unshare |
| `CAP_SYS_BOOT`       | reboot                                                       |
| `CAP_SYS_CHROOT`     | chroot                                                       |
| `CAP_SYS_MODULE`     | delete_module, init_module, finit_module, query_module       |
| `CAP_SYS_PACCT`      | acct                                                         |
| `CAP_SYS_PTRACE`     | kcmp, process_vm_readv, process_vm_writev, ptrace            |
| `CAP_SYS_RAWIO`      | iopl, ioperm                                                 |
| `CAP_SYS_TIME`       | settimeofday, stime, adjtimex, clock_settime                 |
| `CAP_SYS_TTY_CONFIG` | vhangup                                                      |


## Examples

### Unconfined profile

Here's an example of a pod that uses the unconfined profile:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: trustworthy-pod
  annotations:
    seccomp.security.alpha.kubernetes.io/pod: unconfined
spec:
  containers:
    - name: trustworthy-container
      image: sotrustworthy:latest
```

### Custom profile

Here's an example of a pod that uses a profile called `example-explorer-
profile` using the container-level annotation:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: explorer
  annotations:
    container.seccomp.security.alpha.kubernetes.io/explorer: localhost/example-explorer-profile
spec:
  containers:
    - name: explorer
      image: gcr.io/google_containers/explorer:1.0
      args: ["-port=8080"]
      ports:
        - containerPort: 8080
          protocol: TCP
      volumeMounts:
        - mountPath: "/mount/test-volume"
          name: test-volume
  volumes:
    - name: test-volume
      emptyDir: {}
```
