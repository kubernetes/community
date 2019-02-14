# CRI: Dockershim PodSandbox Checkpoint 

## Umbrella Issue
[#34672](https://github.com/kubernetes/kubernetes/issues/34672)

## Background
[Container Runtime Interface (CRI)](/contributors/devel/sig-node/container-runtime-interface.md) 
is an ongoing project to allow container runtimes to integrate with 
kubernetes via a newly-defined API. 
[Dockershim](https://github.com/kubernetes/kubernetes/blob/release-1.5/pkg/kubelet/dockershim) 
is the Docker CRI implementation. This proposal aims to introduce 
checkpoint mechanism in dockershim.

## Motivation
### Why do we need checkpoint?


With CRI, Kubelet only passes configurations (SandboxConfig, 
ContainerConfig and ImageSpec) when creating sandbox, container and 
image, and only use the reference id to manage them after creation. 
However, information in configuration is not only needed during creation. 

In the case of dockershim with CNI network plugin, CNI plugins needs 
the same information from PodSandboxConfig at creation and deletion. 

```
Kubelet ---------------------------------
     | RunPodSandbox(PodSandboxConfig)
     | StopPodSandbox(PodSandboxID)
     V
Dockershim-------------------------------
     | SetUpPod
     | TearDownPod
     V
Network Plugin---------------------------
     | ADD
     | DEL
     V
CNI plugin-------------------------------
```


In addition, checkpoint helps to improve the reliability of dockershim. 
With checkpoints, critical information for disaster recovery could be 
preserved. Kubelet makes decisions based on the reported pod states 
from runtime shims. Dockershim currently gathers states from docker 
engine. However, in case of disaster, docker engine may lose all 
container information, including the reference ids. Without necessary
information, kubelet and dockershim could not conduct proper clean up. 
For example, if docker containers are removed underneath kubelet, reference 
to the allocated IPs and iptables setup for the pods are also lost. 
This leads to resource leak and potential iptables rule conflict. 

### Why checkpoint in dockershim?
- CNI specification does not require CNI plugins to be stateful. And CNI 
specification does not provide interface to retrieve states from CNI plugins. 
- Currently there is no uniform checkpoint requirements across existing runtime shims.
- Need to preserve backward compatibility for kubelet.
- Easier to maintain backward compatibility by checkpointing at a lower level.

## PodSandbox Checkpoint 
Checkpoint file will be created for each PodSandbox. Files will be 
placed under `/var/lib/dockershim/sandbox/`. File name will be the 
corresponding `PodSandboxID`. File content will be json encoded. 
Data structure is as follows:

```go
const schemaVersion = "v1"

type Protocol string

// PortMapping is the port mapping configurations of a sandbox.
type PortMapping struct {
	// Protocol of the port mapping.
	Protocol *Protocol `json:"protocol,omitempty"`
	// Port number within the container.
	ContainerPort *int32 `json:"container_port,omitempty"`
	// Port number on the host.
	HostPort *int32 `json:"host_port,omitempty"`
}

// CheckpointData contains all types of data that can be stored in the checkpoint.
type CheckpointData struct {
	PortMappings []*PortMapping `json:"port_mappings,omitempty"`
}

// PodSandboxCheckpoint is the checkpoint structure for a sandbox
type PodSandboxCheckpoint struct {
	// Version of the pod sandbox checkpoint schema.
	Version string `json:"version"`
	// Pod name of the sandbox. Same as the pod name in the PodSpec.
	Name string `json:"name"`
	// Pod namespace of the sandbox. Same as the pod namespace in the PodSpec.
	Namespace string `json:"namespace"`
	// Data to checkpoint for pod sandbox.
	Data *CheckpointData `json:"data,omitempty"`
}
```


## Workflow Changes


`RunPodSandbox` creates checkpoint: 
```
() --> Pull Image --> Create Sandbox Container --> (Create Sandbox Checkpoint) --> Start Sandbox Container --> Set Up Network --> ()
```

`RemovePodSandbox` removes checkpoint:
```
() --> Remove Sandbox --> (Remove Sandbox Checkpoint) --> ()
```

`ListPodSandbox` need to include all PodSandboxes as long as their 
checkpoint files exist. If sandbox checkpoint exists but sandbox 
container could not be found, the PodSandbox object will include 
PodSandboxID, namespace and name. PodSandbox state will be `PodSandboxState_SANDBOX_NOTREADY`.

`StopPodSandbox` and `RemovePodSandbox` need to conduct proper error handling to ensure idempotency. 



## Future extensions
This proposal is mainly driven by networking use cases. More could be added into checkpoint. 



