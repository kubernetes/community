# Container Runtime Interface: Container Metrics

[Container runtime interface
(CRI)](/contributors/devel/sig-node/container-runtime-interface.md)
provides an abstraction for container runtimes to integrate with Kubernetes.
CRI expects the runtime to provide resource usage statistics for the
containers.

## Background

Historically Kubelet relied on the [cAdvisor](https://github.com/google/cadvisor)
library, an open-source project hosted in a separate repository, to retrieve
container metrics such as CPU and memory usage. These metrics are then aggregated
and exposed through Kubelet's [Summary
API](https://git.k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1/types.go)
for the monitoring pipeline (and other components) to consume. Any container
runtime (e.g., Docker and Rkt) integrated with Kubernetes needed to add a
corresponding package in cAdvisor to support tracking container and image file
system metrics.

With CRI being the new abstraction for integration, it was a natural
progression to augment CRI to serve container metrics to eliminate a separate
integration point.

*See the [core metrics design
proposal](/contributors/design-proposals/instrumentation/core-metrics-pipeline.md)
for more information on metrics exposed by Kubelet, and [monitoring
architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md)
for the evolving monitoring pipeline in Kubernetes.*

# Container Metrics

Kubelet is responsible for creating pod-level cgroups based on the Quality of
Service class to which the pod belongs, and passes this as a parent cgroup to the
runtime so that it can ensure all resources used by the pod (e.g., pod sandbox,
containers) will be charged to the cgroup. Therefore, Kubelet has the ability
to track resource usage at the pod level (using the built-in cAdvisor), and the
API enhancement focuses on the container-level metrics.


We include the only a set of metrics that are necessary to fulfill the needs of
Kubelet. As the requirements evolve over time, we may extend the API to support
more metrics. Below is the API with the metrics supported today.

```go
// ContainerStats returns stats of the container. If the container does not
// exist, the call returns an error.
rpc ContainerStats(ContainerStatsRequest) returns (ContainerStatsResponse) {}
// ListContainerStats returns stats of all running containers.
rpc ListContainerStats(ListContainerStatsRequest) returns (ListContainerStatsResponse) {}
```

```go
// ContainerStats provides the resource usage statistics for a container.
message ContainerStats {
    // Information of the container.
    ContainerAttributes attributes = 1;
    // CPU usage gathered from the container.
    CpuUsage cpu = 2;
    // Memory usage gathered from the container.
    MemoryUsage memory = 3;
    // Usage of the writable layer.
    FilesystemUsage writable_layer = 4;
}

// CpuUsage provides the CPU usage information.
message CpuUsage {
    // Timestamp in nanoseconds at which the information were collected. Must be > 0.
    int64 timestamp = 1;
    // Cumulative CPU usage (sum across all cores) since object creation.
    UInt64Value usage_core_nano_seconds = 2;
}

// MemoryUsage provides the memory usage information.
message MemoryUsage {
    // Timestamp in nanoseconds at which the information were collected. Must be > 0.
    int64 timestamp = 1;
    // The amount of working set memory in bytes.
    UInt64Value working_set_bytes = 2;
}

// FilesystemUsage provides the filesystem usage information.
message FilesystemUsage {
    // Timestamp in nanoseconds at which the information were collected. Must be > 0.
    int64 timestamp = 1;
    // The underlying storage of the filesystem.
    StorageIdentifier storage_id = 2;
    // UsedBytes represents the bytes used for images on the filesystem.
    // This may differ from the total bytes used on the filesystem and may not 
    // equal CapacityBytes - AvailableBytes.
    UInt64Value used_bytes = 3;
    // InodesUsed represents the inodes used by the images.
    // This may not equal InodesCapacity - InodesAvailable because the underlying
    // filesystem may also be used for purposes other than storing images.
    UInt64Value inodes_used = 4;
}
```

There are three categories or resources: CPU, memory, and filesystem. Each of
the resource usage message includes a timestamp to indicate when the usage
statistics is collected. This is necessary because some resource usage (e.g.,
filesystem) are inherently more expensive to collect and may be updated less
frequently than others. Having the timestamp allows the consumer to know how
stale/fresh the data is, while giving the runtime flexibility to adjust.

Although CRI does not dictate the frequency of the stats update, Kubelet needs
a minimum guarantee of freshness of the stats for certain resources so that it
can reclaim them timely when under pressure. We will formulate the requirements
for any of such resources and include them in CRI in the near future.


*For more details on why we request cached stats with timestamps as opposed to
requesting stats on-demand, here is the [rationale](https://github.com/kubernetes/kubernetes/pull/45614#issuecomment-302258090)
behind it.*

## Status

The container metrics calls were added to CRI in Kubernetes 1.7, but Kubelet did not
yet use it to gather metrics from the runtime. In Kubernetes 1.8, Kubelet was
given the option to [consume the container metrics using CRI
stats](https://github.com/kubernetes/kubernetes/pull/51557). See the
`pkg/kubelet/cadvisor.go#UsingLegacyCadvisorStats`
[function](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/cadvisor/util.go#L73)
for more information on how Kubelet determines the proper metrics source.
