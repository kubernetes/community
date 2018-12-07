# Support Node-Level User Namespaces Remapping

- [Summary](#summary)
- [Motivation](#motivation)
- [Goals](#goals)
- [Non-Goals](#non-goals)
- [Use Stories](#user-stories)
- [Proposal](#proposal)
- [Future Work](#future-work)
- [Risks and Mitigations](risks-and-mitigations)
- [Graduation Criteria](graduation-criteria)
- [Alternatives](alternatives)


_Authors:_

* Mrunal Patel &lt;mpatel@redhat.com&gt;
* Jan Pazdziora &lt;jpazdziora@redhat.com&gt;
* Vikas Choudhary &lt;vichoudh@redhat.com&gt;

## Summary
Container security consists of many different kernel features that work together to make containers secure. User namespaces is one such feature that enables interesting possibilities for containers by allowing them to be root inside the container while not being root on the host. This gives more capabilities to the containers while protecting the host from the container being root and adds one more layer to container security. 
In this proposal we discuss:
- use-cases/user-stories that benefit from this enhancement
- implementation design and scope for alpha release
- long-term roadmap to fully support this feature beyond alpha

## Motivation
From user_namespaces(7):
> User namespaces isolate security-related identifiers and attributes, in particular, user IDs and group IDs, the root directory, keys, and capabilities. A process's user and group IDs can be different inside and outside a user namespace. In particular, a process can have a normal unprivileged user ID outside a user namespace while at the same time having a user ID of 0 inside the namespace; in other words, the process has full privileges for operations inside the user namespace, but is unprivileged for operations outside the namespace.

In order to run Pods with software which expects to run as root or with elevated privileges while still containing the processes and protecting both the Nodes and other Pods, Linux kernel mechanism of user namespaces can be used make the processes in the Pods view their environment as having the privileges, while on the host (Node) level these processes appear as without privileges or with privileges only affecting processes in the same Pods

The purpose of using user namespaces in Kubernetes is to let the processes in Pods think they run as one uid set when in fact they run as different “real” uids on the Nodes.

In this text, most everything said about uids can also be applied to gids.

## Goals
Enable user namespace support in a kubernetes cluster so that workloads that work today also work with user namespaces enabled at runtime. Furthermore, make workloads that require root/privileged user inside the container, safer for the node using the additional security of user namespaces. Containers will run in a user namespace different from user-namespace of the underlying host.

## Non-Goals
- Non-goal is to support pod/container level user namespace isolation. There can be images using different users but on the node, pods/containers running with these images will share common user namespace remapping configuration. In other words, all containers on a node share a common user-namespace range.
- Remote volumes support eg. NFS

## User Stories
- As a cluster admin, I want to protect the node from the rogue container process(es) running inside pod containers with root privileges. If such a process is able to break out into the node, it could be a security issue.
- As a cluster admin, I want to support all the images irrespective of what user/group that image is using.
- As a cluster admin, I want to allow some pods to disable user namespaces if they require elevated privileges.

## Proposal
Proposal is to support user-namespaces for the pod containers. This can be done at two levels:
- Node-level : This proposal explains this part in detail.
- Namespace-Level/Pod-level: Plan is to target this in future due to missing support in the low level system components such as runtimes and kernel. More on this in the `Future Work` section.

Node-level user-namespace support means that, if feature is enabled, all pods on a node will share a common user-namespace, common UID(and GID) range (which is a subset of node’s total UIDs(and GIDs)). This common user-namespace is runtime’s default user-namespace range which is remapped to containers’ UIDs(and GID), starting with the first UID as container’s ‘root’.
In general Linux convention, UID(or GID) mapping consists of three parts:
1. Host (U/G)ID: First (U/G)ID of the range on the host that is being remapped to the (U/G)IDs in the container user-namespace
2. Container (U/G)ID: First (U/G)ID of the range in the container namespace and this is mapped to the first (U/G)ID on the host(mentioned in previous point).
3. Count/Size: Total number of consecutive mapping between host and container user-namespaces, starting from the first one (including) mentioned above.

As an example, `host_id 1000, container_id 0, size 10`
In this case, 1000 to 1009 on host will be mapped to 0 to 9 inside the container.

User-namespace support should be enabled only when container runtime on the node supports user-namespace remapping and is enabled in its configuration. To enable user-namespaces, feature-gate flag will need to be passed to Kubelet like this `--feature-gates=”NodeUserNamespace=true”`

A new CRI API, `GetRuntimeConfigInfo` will be added. Kubelet will use this API:
- To verify if user-namespace remapping is enabled at runtime. If found disabled, kubelet will fail to start
- To determine the default user-namespace range at the runtime, starting UID of which is mapped to the UID '0' of the container.

### Volume Permissions
Kubelet will change the file permissions, i.e chown, at `/var/lib/kubelet/pods` prior to any container start to get file permissions updated according to remapped UID and GID.
This proposal will work only for local volumes and not with remote volumes such as NFS.

### How to disable `NodeUserNamespace` for a specific pod
This can be done in two ways:
- **Alpha:** Implicitly using host namespace for the pod containers
This support is already present (currently it seems broken, will be fixed) in Kubernetes as an experimental functionality, which can be enabled using `feature-gates=”ExperimentalHostUserNamespaceDefaulting=true”`.
If Pod-Security-Policy is configured to allow the following to be requested by a  pod, host user-namespace will be enabled for the container:
  - host namespaces (pid, ipc, net)
  - non-namespaced capabilities (mknod, sys_time, sys_module)
  - the pod contains a privileged container or using host path volumes.
  - https://github.com/kubernetes/kubernetes/commit/d0d78f478ce0fb9d5e121db3b7c6993b482af82c#diff-a53fa76e941e0bdaee26dcbc435ad2ffR437 introduced via https://github.com/kubernetes/kubernetes/commit/d0d78f478ce0fb9d5e121db3b7c6993b482af82c.

- **Beta:** Explicit API to request host user-namespace in pod spec
          This is being targeted under Beta graduation plans.

### CRI API Changes
Proposed CRI API changes:

```golang
// Runtime service defines the public APIs for remote container runtimes
service RuntimeService {
    // Version returns the runtime name, runtime version, and runtime API version.
    rpc Version(VersionRequest) returns (VersionResponse) {}
    …….
    …….
    //  GetRuntimeConfigInfo returns the configuration details of the runtime.
    rpc GetRuntimeConfigInfo(GetRuntimeConfigInfoRequest) returns (GetRuntimeConfigInfoResponse) {}
}
// LinuxIDMapping represents a single user namespace mapping in Linux.
message LinuxIDMapping {
   // container_id is the starting id for the mapping inside the container.
   uint32 container_id = 1;
   // host_id is the starting id for the mapping on the host.
   uint32 host_id = 2;
   // size is the length of the mapping.
   uint32 size = 3;
}

message LinuxUserNamespaceConfig {
   // is_enabled, if true indicates that user-namespaces are supported and enabled in the container runtime
   bool is_enabled = 1;
   // uid_mappings is an array of user id mappings.
   repeated LinuxIDMapping uid_mappings = 1;
   // gid_mappings is an array of group id mappings.
   repeated LinuxIDMapping gid_mappings = 2;
}
message GetRuntimeConfig {
    LinuxUserNamespaceConfig user_namespace_config = 1;
}

message GetRuntimeConfigInfoRequest {}

message GetRuntimeConfigInfoResponse {
    GetRuntimeConfig runtime_config = 1
}

...

// NamespaceOption provides options for Linux namespaces.
message NamespaceOption {
	// Network namespace for this container/sandbox.
	// Note: There is currently no way to set CONTAINER scoped network in the Kubernetes API.
	// Namespaces currently set by the kubelet: POD, NODE
	NamespaceMode network = 1;
	// PID namespace for this container/sandbox.
	// Note: The CRI default is POD, but the v1.PodSpec default is CONTAINER.
	// The kubelet's runtime manager will set this to CONTAINER explicitly for v1 pods.
	// Namespaces currently set by the kubelet: POD, CONTAINER, NODE
	NamespaceMode pid = 2;
	// IPC namespace for this container/sandbox.
	// Note: There is currently no way to set CONTAINER scoped IPC in the Kubernetes API.
	// Namespaces currently set by the kubelet: POD, NODE
	NamespaceMode ipc = 3;
	// User namespace for this container/sandbox.
	// Note: There is currently no way to set CONTAINER scoped user namespace in the Kubernetes API.
	// The container runtime should ignore this if user namespace is NOT enabled.
	// POD is the default value. Kubelet will set it to NODE when trying to use host user-namespace
	// Namespaces currently set by the kubelet: POD, NODE
	NamespaceMode user = 4;
}

```

### Runtime Support
- Docker: Here is the [user-namespace documentation](https://docs.docker.com/engine/security/userns-remap/) and this is the [implementation PR](https://github.com/moby/moby/pull/12648)
    - Concerns:
Docker API does not provide user-namespace mapping. Therefore to handle `GetRuntimeConfigInfo` API, changes will be done in `dockershim` to read system files, `/etc/subuid` and `/etc/subgid`, for figuring out default user-namespace mapping. `/info` api will be used to figure out if user-namespace is enabled and `Docker Root Dir` will be used to figure out host uid mapped to the uid `0` in container. eg. `Docker Root Dir: /var/lib/docker/2131616.2131616` this shows host uid `2131616` will be mapped to uid `0`
- CRI-O: https://github.com/kubernetes-incubator/cri-o/pull/1519
- Containerd: https://github.com/containerd/containerd/blob/129167132c5e0dbd1b031badae201a432d1bd681/container_opts_unix.go#L149

### Implementation Roadmap
#### Phase 1: Support in Kubelet, Alpha, [Target: Kubernetes v1.11]
- Add feature gate `NodeUserNamespace`, disabled by default
- Add new CRI API, `GetRuntimeConfigInfo()`
- Add logic in Kubelet to handle pod creation which includes parsing GetRuntimeConfigInfo response and changing file-permissions in /var/lib/kubelet with learned userns mapping.
- Add changes in dockershim to implement GetRuntimeConfigInfo() for docker runtime
- Add changes in CRI-O to implement userns support and GetRuntimeConfigInfo() support
- Unit test cases
- e2e tests

#### Phase 2: Beta Support [Target: Kubernetes v1.12]
- PSP integration
- To grow ExperimentalHostUserNamespaceDefaulting from experimental feature gate to a Kubelet flag
- API changes to allow pod able to request HostUserNamespace in pod spec
- e2e tests

### References
- Default host user namespace via experimental flag
  - https://github.com/kubernetes/kubernetes/pull/31169
- Enable userns support for containers launched by kubelet
  - https://github.com/kubernetes/features/issues/127
- Track Linux User Namespaces in the Pod Security Policy
  - https://github.com/kubernetes/kubernetes/issues/59152
- Add support for experimental-userns-remap-root-uid and experimental-userns-remap-root-gid options to match the remapping used by the container runtime.
  - https://github.com/kubernetes/kubernetes/pull/55707
- rkt User Namespaces Background
  - https://coreos.com/rkt/docs/latest/devel/user-namespaces.html

## Future Work
### Namespace-Level/Pod-Level user-namespace support
There is no runtime today which supports creating containers with a specified user namespace configuration. For example here is the discussion related to this support in Docker https://github.com/moby/moby/issues/28593
Once user-namespace feature in the runtimes has evolved to support container’s request for a specific user-namespace mapping(UID and GID range), we can extend current Node-Level user-namespace support in Kubernetes to support Namespace-level isolation(or if desired even pod-level isolation) by dividing and allocating learned mapping from runtime among Kubernetes namespaces (or pods, if desired). From end-user UI perspective, we don't expect any change in the UI related to user namespaces support.
### Remote Volumes
Remote Volumes support should be investigated and should be targeted in future once support is there at lower infra layers.


## Risks and Mitigations
The main risk with this change stems from the fact that processes in Pods will run with different “real” uids than they used to, while expecting the original uids to make operations on the Nodes or consistently access shared persistent storage.
- This can be mitigated by turning the feature on gradually, per-Pod or per Kubernetes namespace.
- For the Kubernetes' cluster Pods (that provide the Kubernetes functionality), testing of their behaviour and ability to run in user namespaced setups is crucial.

## Graduation Criteria
- PSP integration
- API changes to allow pod able to request host user namespace using for example, `HostUserNamespace: True`, in pod spec
- e2e tests

## Alternatives
User Namespace mappings can be passed explicitly through kubelet flags similar to https://github.com/kubernetes/kubernetes/pull/55707 but we do not prefer this option because this is very much prone to mis-configuration.
