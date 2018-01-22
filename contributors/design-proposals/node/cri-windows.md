# CRI: Windows Container Configuration

**Authors**: Jiangtian Li (@JiangtianLi), Pengfei Ni (@feiskyer)

**Status**: Proposed

## Background
Container Runtime Interface (CRI) defines [APIs and configuration types](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/apis/cri/v1alpha1/runtime/api.proto) for kubelet to integrate various container runtimes. The Open Container Initiative (OCI) Runtime Specification defines [platform specific configuration](https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration), including Linux, Windows, and Solaris. Currently CRI only suppports Linux container configuration.  This proposal is to bring the Memory & CPU resource restrictions already specified in OCI for Windows to CRI.

The Linux & Windows schedulers differ in design and the units used, but can accomplish the same goal of limiting resource consumption of individual containers.

For example, on Linux platform, cpu quota and cpu period represent CPU resource allocation to tasks in a cgroup and cgroup by [Linux kernel CFS scheduler](https://www.kernel.org/doc/Documentation/scheduler/sched-design-CFS.txt). Container created in the cgroup are subject to those limitations, and additional processes forked or created will inherit the same cgroup.

On the Windows platform, processes may be assigned to a job object, which can have [CPU rate control information](https://msdn.microsoft.com/en-us/library/windows/desktop/hh448384(v=vs.85).aspx), memory, and storage resource constraints enforced by the Windows kernel scheduler. A job object is created by Windows to at container creation time so all processes in the container will be aggregated and bound to the resource constraint.

## Umbrella Issue
[#56734](https://github.com/kubernetes/kubernetes/issues/56734)

## Motivation
The goal is to start filling the gap of platform support in CRI, specifically for Windows platform. For example, currrently in dockershim Windows containers are scheduled using the default resource constraints and does not respect the resource requests and limits specified in POD. With this proposal, Windows containers will be able to leverage POD spec and CRI to allocate compute resource and respect restriction.

## Proposed design

The design is faily straightforward and to align CRI container configuration for Windows with [OCI runtime specification](https://github.com/opencontainers/runtime-spec/blob/master/specs-go/config.go):
```
// WindowsResources has container runtime resource constraints for containers running on Windows.
type WindowsResources struct {
	// Memory restriction configuration.
	Memory *WindowsMemoryResources `json:"memory,omitempty"`
	// CPU resource restriction configuration.
	CPU *WindowsCPUResources `json:"cpu,omitempty"`
}
```

Since Storage and Iops for Windows containers is optional, it can be postponed to align with Linux container configuration in CRI. Therefore we propose to add the following to CRI for Windows container (PR [here](https://github.com/kubernetes/kubernetes/pull/57076)).

### API definition
```
// WindowsContainerConfig contains platform-specific configuration for
// Windows-based containers.
message WindowsContainerConfig {
    // Resources specification for the container.
    WindowsContainerResources resources = 1;
}

// WindowsContainerResources specifies Windows specific configuration for
// resources.
message WindowsContainerResources {
    // CPU shares (relative weight vs. other containers). Default: 0 (not specified).
    int64 cpu_shares = 1;
    // Number of CPUs available to the container. Default: 0 (not specified).
    int64 cpu_count = 2;
    // Specifies the portion of processor cycles that this container can use as a percentage times 100.
    int64 cpu_maximum = 3;
    // Memory limit in bytes. Default: 0 (not specified).
    int64 memory_limit_in_bytes = 4;
}
```

## Implementation
The implementation will mainly be in two parts:
* In kuberuntime, where configuration is generated from POD spec.
* In container runtime, where configuration is passed to container configuration. For example, in dockershim, passed to [HostConfig](https://github.com/moby/moby/blob/master/api/types/container/host_config.go).

In both parts, we need to implement:
* Fork code for Windows from Linux
* Convert from Resources.Requests and Resources.Limits to Windows configuration in CRI, and convert from Windows configration in CRI to container configuration.

To implement resource controls for Windows containers, refer to [this MSDN documentation](https://docs.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/resource-controls) and [Docker's conversion to OCI spec](https://github.com/moby/moby/blob/master/daemon/oci_windows.go).


## Future work

Windows [storage resource controls](https://github.com/opencontainers/runtime-spec/blob/master/config-windows.md#storage), security context (analog to SELinux, Apparmor, readOnlyRootFilesystem, etc.) and pod resource controls (analog to LinuxPodSandboxConfig.cgroup_parent already in CRI) are under investigation and would be handled in separate propsals. They will supplement and not replace the fields in `WindowsContainerResources` from this proposal.