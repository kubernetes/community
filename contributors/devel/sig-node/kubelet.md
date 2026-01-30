# Kubelet Component

This document provides a high-level walkthrough of the Kubelet's code structure,
specifically focusing on the Pod lifecycle and the "Sync Loop". It tracks the
flow of a Pod from the moment it is assigned to a node to its execution and
eventual termination.

## Overview & Scope

The Kubelet deals primarily in **Pods**. While it handles some supporting
resources (Volumes, etc.), its main unit of work is the Pod. It does not
natively understand Deployments, StatefulSets, or DaemonSets, or any other
workload abstractions -- those are handled by upstream controllers.

1.  **Creation**: Workload controllers create Pods.
2.  **Scheduling**: The Scheduler picks a node and **binds** the Pod to it
    (setting `spec.nodeName`).
3.  **Kubelet Observation**: Only after a Pod is bound to a node does the
    Kubelet on that node see it and take ownership. _**<==== WE ARE HERE!!!**_ 

The Kubelet **orchestrates** the Pod spec into running processes on the host via
the **Container Runtime Interface (CRI)**. The runtime (e.g., containerd)
translates these requests for an OCI runtime (e.g. `runc`), which handles the
low-level operating system setup for the container.

## Startup Flow: From API to Pod Worker

When a new Pod is assigned to a node, it follows this path through the Kubelet
internals:

1.  **Config Source**: An
    [Informer](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/config/apiserver.go#L38-L39)
    watches for Pods bound to this node.
2.  **PodConfig**: The
    [PodConfig](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/config/config.go#L75)
    aggregates updates from multiple sources (API server, file static pods, HTTP
    static pods) and merges them into a single stream of updates.
3.  **SyncLoopIteration**: The main control loop
    [syncLoopIteration](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L2574)
    receives these updates.
4.  **HandlePodAdditions**: The
    [HandlePodAdditions](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L2713)
    function handles the initial setup:
    *   **Allocation**: Calls
        [allocationManager.AddPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/allocation/allocation_manager.go#L527)
        to check internal admission checks (e.g., topology resources).
5.  **Pod Workers**: The work is queued to a dedicated [Pod
    Worker](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/pod_workers.go#L755).
    Each Pod has its own worker goroutine to ensure operations for a single Pod
    are serialized.

## SyncPod: The Core Logic

The
[SyncPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L1941)
function is the primary workhorse. It orchestrates the following steps:

1.  **Status Generation**: [Generates the
    status](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L1996)
    representing the state *before* work begins.
2.  **Cgroups**: Calls
    [pcm.EnsureExists](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/cm/pod_container_manager_linux.go#L74)
    to configure Pod-level cgroups.
    *   *Note*: While the container runtime manages container cgroups, the
        Kubelet manages the parent Pod cgroups (a holdover from the Docker shim
        era).
3.  **Volumes**: Calls
    [WaitForAttachAndMount](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L2120)
    to ensure volumes are ready.
4.  **Runtime Sync**: Hand off to `kuberuntime_manager.SyncPod` for container
    operations.

### Runtime Manager Sync

The [kuberuntime
SyncPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L1394)
function bridges the declarative Kubernetes API to the imperative CRI API.

1.  **Compute Actions**:
    [computePodActions](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L1397)
    compares the desired spec with the current runtime state. It outputs a
    [podActions](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L563)
    struct listing what to kill, create, or update.
2.  **Actuation**:
    *   **Sandboxes**: Calls
        [createPodSandbox](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L1543)
        if needed.
    *   **Containers**: For each container to start, calls
        [startContainer](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_container.go#L200):
        1.  Checks CrashLoopBackOff.
        2.  [Ensures Image
            Exists](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_container.go#L215)
            (pulling if necessary).
        3.  [Generates Container
            Config](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_container.go#L343)
            (this is when it translates K8s API to [CRI
            spec](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1/api.proto)).
        4.  Calls CRI `CreateContainer` and `StartContainer`.
        5.  Executes the **PostStart hook** synchronously (blocking other
            containers in the Pod!).

## Steady State: PLEG and Probes

Once running, the Kubelet maintains the Pod via:

*   **PLEG (Pod Lifecycle Event Generator)**: [Polls the
    runtime](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/pleg/generic.go#L128)
    (default every ~2s) to detect state changes. If a [change is
    detected](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/pleg/generic.go#L397),
    it generates an event to wake up the Sync Loop.
*   **Probes**: The
    [ProbeManager](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/prober/prober_manager.go#L185)
    runs workers for Liveness, Readiness, and Startup probes. Results trigger
    updates via
    [resultsManager](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/prober/worker.go#L365).

## Termination

Termination is handled by
[SyncTerminatingPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L2182):

1.  Stops probes.
2.  Calls the kuberuntime's
    [KillPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L1860)
    which itself calls the container runtime using CRI to stop containers and
    the sandbox.
3.  Generates final status.

After termination,
[SyncTerminatedPod](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet.go#L2339)
performs final cleanup (unmounting volumes, releasing resources).

## Garbage Collection

*   **Container GC**: Periodic loop in the runtime manager to remove exited
    containers.
*   **Image GC**: Periodic check to remove unused images based on disk pressure.
*   **"Housekeeping"**:
    [HandlePodCleanups](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/kubelet_pods.go#L1192)
    cleans up orphaned pod directories and internal state.

## Advanced Concepts

### Static & Mirror Pods

*   **Static Pod**: A Pod sourced from a local file or HTTP endpoint, not the
    API server. The API server initially knows nothing about it.
*   **Mirror Pod**: A read-only Pod object created by the Kubelet in the API
    server to represent a Static Pod. This allows the Scheduler to see the
    resource usage and users to see the status. Identified by the
    `kubernetes.io/config.mirror` annotation.

### In-Place Resize

Resize involves reconciling four resource states:
1.  **Desired**: From Pod Spec.
2.  **Allocated**: Admitted by Kubelet (persisted in checkpoints).
3.  **Actuated**: Successfully applied to the runtime.
4.  **Actual**: Read back from cgroups.

The
[AllocationManager](https://github.com/kubernetes/kubernetes/blob/03e14cc9432975dec161de1e52d7010f9711a913/pkg/kubelet/allocation/allocation_manager.go#L527)
mediates these transitions.

## Testing

Kubelet testing is split into:
*   **Unit Tests**: Heavily used for logic verification.
*   **Node E2E**:
    [test/e2e_node](https://github.com/kubernetes/kubernetes/tree/master/test/e2e_node).
    Runs a single-node cluster (Kubelet + API Server typically) to test Kubelet
    in isolation on various OS/Runtime combinations.
*   **Cluster E2E**:
    [test/e2e/node](https://github.com/kubernetes/kubernetes/tree/master/test/e2e/node).
    Tests that require a full control plane.
*   **Common**:
    [test/e2e/common](https://github.com/kubernetes/kubernetes/tree/master/test/e2e/common).
    Tests compliant with both environments.
