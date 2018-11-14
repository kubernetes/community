# Monitoring support for hardware accelerators

Version: Alpha

Owner: @mindprince (agarwalrohit@google.com)

## Motivation

We have had alpha support for running containers with GPUs attached in Kubernetes for a while. To take this to beta and GA, we need to provide GPU monitoring, so that users can get insights into how their GPU jobs are performing.

## Detailed Design

The current metrics pipeline for Kubernetes is:
- Container level metrics are collected by [cAdvisor](https://github.com/google/cadvisor).
- Kubelet embeds cAdvisor as a library. It uses its knowledge of pod-to-container mappings and the metrics from cAdvisor to expose pod level metrics as the summary API.
- [Heapster](https://github.com/kubernetes/heapster) uses kubelet’s summary API and pushes metrics to the some sink.
There are plans to change this pipeline but the details for that are still not finalized.

To expose GPU metrics to Kubernetes users, we would need to make changes to all these components.

First up is cAdvisor: we need to make cAdvisor collect metrics for GPUs that are attached to a container.

The source for getting metrics for NVIDIA GPUs is [NVIDIA Management Library (NVML)](https://developer.nvidia.com/nvidia-management-library-nvml). NVML is a closed source C library [with a documented API](http://docs.nvidia.com/deploy/nvml-api/index.html). Because we want to use NVML from cAdvisor (which is written in Go), we need to [write a cgo wrapper for NVML](https://github.com/mindprince/gonvml).

The cAdvisor binary is statically linked currently. Because we can’t statically link the closed source NVML code, we would need to make cAdvisor a dynamically linked binary. We would use `dlopen` in the cgo wrapper to dynamically load NVML. Because kubelet embeds cAdvisor, kubelet will also need to be a dynamically linked binary. In my testing, kubelet running on GCE 1.7.x clusters was found to be a dynamically linked binary already but now being dynamically linked will become a requirement.

When cAdvisor starts up, it would read the vendor files in `/sys/bus/pci/devices/*` to see if any NVIDIA devices (vendor ID: `0x10de`) are attached to the node.
- If no NVIDIA devices are found, this code path would become dormant for the rest of cAdvisor/kubelet lifetime.
- If NVIDIA devices are found, we would start a goroutine that would check for the presence of NVML by trying to dynamically load it at regular intervals (say every minute or every 5 minutes). We need to do this regular checking instead of doing it just once because it may happen that cAdvisor is started before the nvidia drivers and nvml are installed. Once the NVML dynamic loading succeeds, we would use NVML’s query methods to find out how many devices exist on the node and create a map from their minor numbers to their handles and cache that map. The goroutine would exit at this point.

If we detected the presence of NVML in the previous step, whenever a new container is detected by cAdvisor, cAdvisor would read the `devices.list` file from the container [devices cgroup](https://www.kernel.org/doc/Documentation/cgroup-v1/devices.txt). The `devices.list` file lists the major:minor number of all the devices that the container is allowed to access. If we find any device with major number `195` ([which is the major number assigned to NVIDIA devices](https://github.com/torvalds/linux/blob/v4.13/Documentation/admin-guide/devices.txt#L2583)), we would cache the list of corresponding minor numbers for that container.

During every housekeeping operation, in addition to collecting all the existing metrics, we will use the cached nvidia device minor numbers and the map from minor numbers to device handles to get metrics for GPU devices attached to the container.

The following new metrics would be exposed per container from cAdvisor:

```
type ContainerStats struct {
...
        // Metrics for Accelerators.
        // Each Accelerator corresponds to one element in the array.
        Accelerators []AcceleratorStats `json:"accelerators,omitempty"`
...
}

type AcceleratorStats struct {
        // Make of the accelerator (nvidia, amd, google etc.)
        Make string `json:"make"`

        // Model of the accelerator (tesla-p100, tesla-k80)
        Model string `json:"model"`

        // ID of the accelerator. device minor number? Or UUID?
        ID string `json:"id"`

        // Total accelerator memory.
        // unit: bytes
        MemoryTotal uint64 `json:"memory_total"`

        // Total accelerator memory allocated.
        // unit: bytes
        MemoryUsed uint64 `json:"memory_used"`

        // Percent of time over the past sample period during which
        // the accelerator was actively processing.
        DutyCycle uint64 `json:"duty_cycle"`
}
```

The API is generic to add support for different types of accelerators in the future even though we will only add support for NVIDIA GPUs initially. The API is inspired by what Google has in borg.

We will update kubelet’s summary API to also add these metrics.

From the summary API, they will flow to heapster and stackdriver.

## Caveats
- As mentioned before, this would add a requirement that cAdvisor and kubelet are dynamically linked.
- We would need to make sure that kubelet is able to access the nvml libraries. Some existing container based nvidia driver installers install drivers in a special directory. We would need to make sure that directory is in kubelet’s `LD_LIBRARY_PATH`.

## Testing Plan
- Adding unit tests and e2e tests to cAdvisor for this code.
- Manually testing various scenarios with nvml installed and not installed; containers running with nvidia devices attached and not attached.
- Performance/Utilization testing: impact on cAdvisor/kubelet resource usage. Impact on GPU performance when we collect metrics.

## Alternatives Rejected
Why collect GPU metrics in cAdvisor? Why not collect them in [device plugins](/contributors/design-proposals/resource-management/device-plugin.md)? The path forward if we collected GPU metrics in device plugin is not clear and may take a lot of time to get finalized.

Here’s a rough sketch of how things could work:

(1) device plugin -> kubelet summary API -> heapster -> ...
- device plugin collects GPU metrics using the cgo wrapper. This is straightforward, in fact, this may even be easier because we don’t have to worry about making kubelet dynamically linked.
- device plugin exposes a new container-level metrics API. This is complicated. There's no good way to have a device plugin metrics API. All we can have is a device plugin metrics endpoint. We can't really define how the metrics inside that will look like because different device types can have wildly different metrics. We can't have a metrics structure that will work well both for GPUs and NICs for example.
- We would have to make the kubelet understand whatever metrics are exposed in the device plugin metrics endpoint and expose it though the summary API. This is not ideal because device plugins are out-of-tree and controlled by vendors, so there can’t a mapping between the metrics exposed by the device plugins and what’s exposed in the kubelet’s summary API. If we try to define such a mapping, it becomes an implicit API that new device plugins have to follow to get their metrics exposed by the kubelet or they would have to update the mapping.

(2) device plugin -> heapster -> ...
- If we don’t go through the kubelet, we can make heapster directly talk to the metrics endpoint exposed by the device plugin. This has the same problem as the last bullet point: how would heapster understand the metrics exposed by the device plugins so that it [can expose them to its backends](https://github.com/kubernetes/heapster/blob/v1.4.3/docs/storage-schema.md). In addition, we would have to solve the issue of how to map containers to their pods.

(3) device plugin -> …
- If we don’t go through kubelet or heapster. We can have the device plugins directly expose metrics to the monitoring agent. For example, device plugins can expose a /metrics endpoint in prometheus format and prometheus can scrape it directly or a prom-to-sd container can send metrics from that endpoint directly to stackdriver. This becomes a more DIY solution, where there’s no real monitoring support provided by kubernetes and device plugin vendors are expected to add metrics to their plugins and users/operators are expected to plumb those metrics to their metrics storage backends. This approach also requires a way to map containers to their pods.

Once the new monitoring architecture plans are implemented, we can revisit this and maybe collect GPU metrics in device plugins instead of cAdvisor.
