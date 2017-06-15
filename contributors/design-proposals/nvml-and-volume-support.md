# Proposal for NVML support and nvidia driver volume injection
We propose leveraging the following to enhance the experimental GPU work:
*  [NVML library](https://developer.nvidia.com/nvidia-management-library-nvml) to conduct GPU discovery
* Nvidia driver volume injection

## NVML support and dynamic loading
Currently, GPU discovery is implemented by querying nvidia GPU devices under
`/dev/nvidia*` directly, which may bring several problems. For example, it
cannot detect GPU failures and  it only collects the entry of the device.  We
think NVML is a more reliable and powerful method for GPU discovery, and advanced  GPU features (driver volume injection, GPU failure detection, heterogeneous GPU, topology-aware scheduling).

We have found that dynamic loading is best when including the NVML library,
because we can check for its existence at runtime. This allows a single kubelet
binary to work on nodes that do not include NVML since we can skip all GPU logic.


## Nvidia driver volume injection
 Container-based applications  need Nvidia GPU drivers to interface with the
 GPU. Installing the  GPU driver in the container is problematic  because the
 GPU driver version must align with the host GPU driver version exactly for
 proper functionality. Both host and container libraries need to be updated
 simultaneously in this deployment.

To simplify, we can borrow an idea from
the [nvidia-docker-plugin](https://github.com/NVIDIA/nvidia-docker). We can
extend the kubelet to collect and manage Nvidia driver binary and library files
on a single volume, and automatically inject this volume into a container job
requesting GPUs. This simplifies our deployment and allows the host
and all containers to share the same GPU drivers.

## Architecture of GPU Manger Module
The current `NvidiaGPUManager` handles all GPU allocation related logic. This
proposal aims at introducing two submodules under
`NvidiaGPUManager`;`NvidiaGPUAllocator` and
`NvidiaGPUVolumeManager`. `NvidiaGPUAllocator` is the original
`NvidiaGPUManager`, which is responsible for GPU allocation
logics. `NvidiaGPUVolumeManager` takes responsibility for driver volume injection
logic. We believe that other accelerators, such as AMD/Intel GPU and FPGA, will
benefit from the volume support in future releases.

### Architecture Example
![Architecture](images/nvml-and-volume-support.png?raw=true "Architecture")
* nvmlWrapper:  Go wrapper for nvml C library that will dynamically link to NVML lib (libnvidia-ml.so.1) during runtime
* fakeNvmlWrapper: This is a fake nvml go wrapper for testing propose only.  We
  can fake low-level GPU related calls so that the unit tests can run without GPU and NVML.
* NvidiaGPUAllocator: The original NvidiaGPUManager that handles GPU allocation related logic
* NvidiaGPUVolumeManger: Manages nvidia GPU volume injection
* Volume: This module borrows from the  nvidia-docker-plugin and handles low-level bin/lib collection and management.
* NvidiaGPUManager: The overall abstract to manage nvidia GPU allocation and volume management

## Structure and interface
This is a brief list for structure/interface of important modules. We will
ignore detailed input/output parameters in the interface because it is still under design now.
* NvidiaGPUManager structure:
```go
type NvidiaGPUManager struct {
        allocator     NvidiaGPUAllocator
        volumeManager NvidiaGPUVolumeManager
}
```
* GPUAllocator interface, implemented by NvidiaGPUAllocator:
```go
type GPUAllocator interface {
        // Start logically initializes GPUAllocator.
        Start() error
        // Capacity returns the total number of GPUs on the node.
        Capacity() v1.ResourceList
        // AllocateGPU attempts to allocate GPUs for input container.
        // Returns paths to allocated GPUs and nil on success.
        // Returns an error on failure.
        AllocateGPU(*v1.Pod, *v1.Container) ([]string, error)
}
```
* GPUVolumeManager interface, implemented by NvidiaGPUVolumeManager:
```go
type GPUVolumeManager interface {
        // Create a volume for GPU devices use.
        Prepare() error
        // Inject volume to the container with GPU binded.
        Inject() error
        // Clean volume on local disk.
        Remove() error
}
```
