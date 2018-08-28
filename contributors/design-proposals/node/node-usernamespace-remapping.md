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
- To let user facing API request specific mapping between host usernamespace and container usernamespace
- Remote volumes support eg. NFS

## User Stories
- As a cluster admin, I want to protect the node from the rogue container process(es) running inside pod containers with root privileges. If such a process is able to break out into the node, it could be a security issue.
- As a cluster admin, I want to support all the images irrespective of what user/group that image is using.
- As a cluster admin, I want to allow some pods to disable user namespaces if they require elevated privileges.

## Proposal
Proposal is to leverage usernamespace remapping security feature for the Kubernetes pods, if enabled at container runtime. 

There are three possible ways in which a remapped usernamespace range can be consumed by the pods:
- Node-level: All pods on a node share a common usernamespace mapping. This scheme provides an additional security layer between the host and the pods running on it. 
- Namespace-Level: In this scheme, all pods on a node, belonging to the same kubernetes namespace, use a subset of the entire usernamespace mapping range. Each subset range is non-overlapping. This scheme provides additional security layer between hosts and also between pod-to-pod, if pods belong to different namesapces.
- Pod-Level/Container-Level: Similar to namespace-level, usernamesapce remapping range is divided into non-overlapping subsets, where each pod gets its own subset or range.

As we increase isolation granularity from node-level to pod-level, security improves at the cost of implementation complexity of the subset range allocation algorithm. This feature will have implications at UI as well. Security context has a number of fields like `RunAsUser`, `FsGroup` etc which are going to be affected. We propose to first implement at Node-level which does not involve any complex allocation logic and focus more on making backward compatible API changes. In next phases, we will go on to increase isolation granularity.

Now, lets discuss Node-level user-namespace support in detail:
1. All the pods on a (node)runtime which are running with "remapped" usernamespace, will share a common user-namespace, common UID(and GID) range (which is a subset of node’s total UIDs(and GIDs)). This common user-namespace is runtime’s default user-namespace mapping range which is remapped to containers’ UIDs(and GID), starting with the first UID as container’s ‘root’.

In general Linux convention, UID(or GID) mapping consists of three parts:
    - Host (U/G)ID: First (U/G)ID of the range on the host that is being remapped to the (U/G)IDs in the container user-namespace
    - Container (U/G)ID: First (U/G)ID of the range in the container namespace and this is mapped to the first (U/G)ID on the host(mentioned in previous point).
    - Count/Size: Total number of consecutive mapping between host and container user-namespaces, starting from the first one (including) mentioned above.

As an example, `host_id 1000, container_id 0, size 10`
In this case, 1000 to 1009 on host will be mapped to 0 to 9 inside the container.

2. Enabling usernamespace remapping at runtime will not require any changes in the existing pod yamls. Kubelet will make sure that older pod yamls will continue to work even in usernamespace remapped configuration. There are valid pod configurations which cannot be supported in usernamespace remapped environment, ex:
  - host namespaces (pid, ipc, net)
  - non-namespaced capabilities (mknod, sys_time, sys_module)
  - The pod contains a privileged container or using host path volumes.
To maintain **backward compatibility**, if such a pod configuration is detected by Kubelet, usernamespace remapping will be disabled, implicitly, for each of such pod in isolation by passing case appropriate `NamespaceOption` to the runtime. For more details see the CRI section below. Other pods will be run with remapped usernamespace.

### Pod API Change
`HostUserNamespace` field will be introduced in the pod spec. Its meaning is similar to other `Host*` fields the pod spec. If set to `true`, pod will share usernamespace from underlying host.
```
        // Use host's user namespace for the pod.
        // This field is alpha-level and can be set to true/false only if HostUserNamespace feature-gate is enabled.
        // Optional: Defaults to nil which means behavior will be container runtime defined.
        // +k8s:conversion-gen=false
        // +optional
        HostUserNamespace *bool `json:"hostUserNamespace,omitempty" protobuf:"varint,31,opt,name=hostUserNamespace"`
```
Today end users are NOT provided any guaranteed behavior by the Kubernetes related to usernamespace remapping support at runtime. If runtime has usernamespace remapping disabled, pods will run otherwise pods will fail. In other words, behavior is container runtime defined. `nil` value of `HostUserNamespace` will show the same behavior i.e runtime defined usernamespace behavior. However, with an added improvement of running pods with remapped usernamespace "when possible". "when possible" means runtime is running with usernamespace remapping enabled and pod does not specify any configuration which is not possible to provide in remapped usernamespace. 

`HostUserNamespace` field will be gated behind `HostUserNamespace` feature-gate. To set it to `true` or `false`, feature-gate will have to be enabled. If feature-gate not enabled or feature-gate enabled but any value is not set, `nil` will be the default value. 

### Effect on existing Security Context fields

* **RunAsUser**: UID to run the entrypoint of the container process. After usernamespace support in kubelet, if Kubelet discovers that usernamespace remapping is enabled at runtime, this will be the UID in the container usernamespace. In such a case, if no mapping(corresponding host UID) is found in the mapping range, pod will be failed admission at Kubelet. If kubelet discovers that usernamespace remapping is not supported at runtime, no validation check will be performed. 
* **RunAsGroup**: same as **RunAsUser** but for gid mapping ranges.
* **RunAsNonRoot**: Today, kubelet fails to start the container if image is found to be having UID 0. After usernamespace support in kubelet, kubelet should ignore this check if runtime is found to be running with usernamespace remapping enabled. Kubelet should verify image UID only when it is learned that usernamespace remapping is either not supported or not enabled at runtime.
* **FsGroup and SupplementalGroups**: FsGroup and SupplementalGroups will continue to work after usernamespace support is Kubelet if runtime either does not support usernamespaces or support is not enabled. In case usernamespace support is discovered by kubelet as enabled at runtime, FsGroups/SupplementalGroups will work only if supported by the runtime (CRI implementation). One implementation, for example, could be where runtime(CRI implementation) will add a triplet in `/proc/<container-pid>/gid_map` for the FsGroup/Supplemental GIDs. This will be a 1:1 mapping between host usernamespace group id and container usernamespace group id i.e container ID and host ID will be same. Example:

    If pod has `FsGroup: 2000` and runtime's default gid mappings are `containerid 0,  hostid 100000, size 1000`, runtime should be adding two entries in the `/proc/<container-pid>/gid_map`:
    - [0 100000 1000]  // From GID 0 to 1000 in the container are mapped with GID range 100k to 100k+1000 on host
    - [2000 2000 1]    // GID 2000 in container usernamespace is mapped to GID 2000 is host usernamespace

    Once this 1:1 mapping is set for the container process, `runc` already has the logic to call `setgroups()` to set the FsGroupID as the supplemental ID in the container usernamespace.

    Setting the permissions on volume is resposibility of volume plugin. Once GID on the volume matches the supplemental GID in the container, access will work.

    NOTE: Pod will be failed admission at kubelet if FsGroup/SupplementalGroups are in the default mapping range and mappings are not 1:1 i.e host id for a container id is different.


### Volume Permissions  
Kubelet will change the file permissions, i.e chown, at `/var/lib/kubelet/pods/<pod-uid>/volumes` after volume plugin is done with provisioning to get updated ownership of the volume directory according to remapped UID and GID.
This proposal will work only for local volumes and not with remote volumes which have client-server kind of complex architecture where client is not allowed to chown volumes as root such as NFS.


### CRI API Changes
A new CRI API, `GetRuntimeConfigInfo` will be added. At the initialization, Kubelet will use this API to retrieve default {U/G}ID mappings from the container runtime. 

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

// GetRuntimeConfigInfoRequest is the message sent for requesting runtime configuration details
message GetRuntimeConfigInfoRequest {}

// GetRuntimeConfigInfoResponse is the response message from runtime that includes configuration details
message GetRuntimeConfigInfoResponse {
    ActiveRuntimeConfig runtime_config = 1
}

// ActiveRuntimeConfig contains the configuration details from the runtime.
message ActiveRuntimeConfig {
    LinuxUserNamespaceConfig user_namespace_config = 1;
}

// LinuxUserNamespaceConfig represents runtime's user-namespace configuration on a linux host.
message LinuxUserNamespaceConfig {
   // uid_mappings is an array of user id mappings.
   repeated LinuxIDMapping uid_mappings = 1;
   // gid_mappings is an array of group id mappings.
   repeated LinuxIDMapping gid_mappings = 2;
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
    // The container runtime should IGNORE this if user namespace remapping is NOT supported.
    // Kubelet will set it to NODE_WIDE_REMAPPED if the following is true:
    // - pod spec field `HostUserNamespace` == nil OR false
    // - AND host usernamespace requiring configuration like hostpath volumes not there in pod
    // - AND userns remapping is enabled at runtime  
    // Kubelet will set it to NODE if the following is true:
    // - userns remapping is NOT enabled at runtime
    // - OR remapping enabled at runtime AND pod spec field `HostUserNamespace` == false
    // - OR remapping enabled at runtime AND host usernamespace requiring configuration like hostpath volumes present in pod definition
    // Namespaces currently set by the kubelet: NODE, NODE_WIDE_REMAPPED
	NamespaceMode user = 4;
}

// A NamespaceMode describes the intended namespace configuration for each
// of the namespaces (Network, PID, IPC) in NamespaceOption. Runtimes should
// map these modes as appropriate for the technology underlying the runtime.
enum NamespaceMode {
    // A POD namespace is common to all containers in a pod.
    // For example, a container with a PID namespace of POD expects to view
    // all of the processes in all of the containers in the pod.
    POD       = 0;
    // A CONTAINER namespace is restricted to a single container.
    // For example, a container with a PID namespace of CONTAINER expects to
    // view only the processes in that container.
    CONTAINER = 1; 
    // A NODE namespace is the namespace of the Kubernetes node.
    // For example, a container with a PID namespace of NODE expects to view
    // all of the processes on the host running the kubelet.
    NODE      = 2; 
    // A NODE_WIDE_REMAPPED namespace applies to all pods on a given kubernetes node.
    // The uid/gids of the pods/containers on a given node are mapped to a range of uids/gids
    // from that Node's namespace. For e.g. starting with 200000 id on the nodes namespace,
    // 10000 ids are allocated for a pod. This means uid/gid 0 inside the pod would map to 200000 id on the nodes namespace.
    // For example, starting from uid/gid 0, 10000 uids/gids in this namespace are mapped to
    // 10000 ids on node namespace,starting with id 200000. i.e uid/gid 0 in this namespace is
    // mapped to uid/gid 200000 on node namespace.
    NODE_WIDE_REMAPPED = 3; 
}

```

### Determining state of the Usernamespace remapping feature at runtime from the received ID mappings
- NOT Supported: If there is only one gid and one uid mapping with all the values as `0`,for example `host_id 0, container_id 0, size 0`
- Supported but disabled: If there is only one gid and one uid mapping. But the size is `4294967295`(max value for uint32) i.e `host_id 0, container_id 0, size 4294967295`
- Supported and enabled: Host {U/G}ID is not `0` and size is anything between `0` and `4294967295`                                

### Pod admission at Kubelet
A pod admit handler will be introduced which will fail the pod admission at Kubelet if any of the following conditions is met:
- If pod spec has non-nil `hostUserNamespace` value AND the container runtime is sandbox type (does not support usernamespace remapping)
- If pod spec has `hostuserNamespace: false` AND usernamespace remapping is NOT enabled at container runtime
- If pod has **RunAsUser** or **RunAsGroup** AND  usernamespace remapping is enabled at runtime AND no mapping is found in idmappings corresponding to the **RunAsUser** or **RunAsGroup**
- If pod has **FsGroup** or **SupplementalGroups** AND  usernamespace remapping is enabled at runtime AND mapping is found in idmappings corresponding to the **FsGroup** or **SupplementalGroups** where container id is not equal to host id i.e mapping is not 1:1


## Runtime(runc based) Support

- Docker: Here is the [user-namespace documentation](https://docs.docker.com/engine/security/userns-remap/) and this is the [implementation PR](https://github.com/moby/moby/pull/12648)
    - Concerns:
Docker API does not provide user-namespace mapping. Therefore to handle `GetRuntimeConfigInfo` API, changes will be done in `dockershim` to read system files, `/etc/subuid` and `/etc/subgid`, for figuring out default user-namespace mapping. `/info` api will be used to figure out if user-namespace is enabled and `Docker Root Dir` will be used to figure out host uid mapped to the uid `0` in container. eg. `Docker Root Dir: /var/lib/docker/2131616.2131616` this shows host uid `2131616` will be mapped to uid `0`
- CRI-O: CRI-O supports usernamespace remapping. [Here is the PR](https://github.com/kubernetes-sigs/cri-o/pull/1941) for getting supplemental groups(fsgroups) working in namespaced environnment.
- Containerd: https://github.com/containerd/containerd/blob/129167132c5e0dbd1b031badae201a432d1bd681/container_opts_unix.go#L149

## Sandbox type runtimes
`HostUserNamespace` will work only with runc based runtimes. Techniques such as node selectors will have to be used in pods(that are using `HostuserNamespace`) to correctly assign runtimes/nodes. If such a pod unexpectedly lands on a sandbox type runtime node, pod admit handler will reject it.


## Implementation Roadmap
### Phase 1: Support in Kubelet, Alpha, [Target: Kubernetes v1.14]
- Add `hostUserNamespace` in pod API
- Add feature gate `HostUserNamespace`, disabled by default
- Add new CRI API, `GetRuntimeConfigInfo()`
- Add logic in Kubelet to handle pod creation which includes parsing GetRuntimeConfigInfo response and changing file-permissions in /var/lib/kubelet with learned userns mapping.
- Add changes in dockershim to implement GetRuntimeConfigInfo() for docker runtime
- Add changes in CRI-O to implement userns support and GetRuntimeConfigInfo() support
- Unit test cases
- e2e tests

### Phase 2: Beta Support [Target: Kubernetes v1.15]
- PSP integration
- Namespace-Level usernamespace remapping support

## Rollout Strategy
Upgrading a cluster to 1.14 will not impact any existing workloads (Assuming that existing cluster nodes were not running runtimes with usernamespace remapping enabled). To be able to use the security benefit of usernamespace remapping on existing node, usernamespace remapping will have to enabled at the runtime. To perform this runtime configuration update, first node will be tainted to make it unschedulable, then it will be drained and runtime will be updated. Next, the taint will removed to make it available to the schedular. Now Kubelet will discover this configuration change and starts creating pods with remapped usernamespace.
As we discussed in above sections, Values of `RunAsUser`, `RunAsGroup`, `SupplementalGroups` and `FsGroup` in the older/existing templates may cause pod admission failure if the values are not consistent with the default id mappings of the runtime as learned by Kubelet.

`RunAsUser` and `RunAsGroup` values are treated as translated IDs in the container usernamespace and if the runtime's default mappings does not have a mapping for these containerIds, pod will fail admission at Kubelet.

`FsGroup` and `SupplementalGroups` values are are not treated as translated IDs. For a FsGroup/SupplementalGroup ID, container usernamespace ID is mapped to the same host usernamespace ID. Therefore for these GIDs in the pod objects in the older templates, either there should be no mapping in runtime's default id mappings or if mapping is there, container id and host id should be same fopr the mapping range i.e 1:1 mapping.

### Recommended default id mappings configuration for the runtime
- Map `0` {U/G}ID in the container namespace with max {u/g}id i.e `4294967295`
- Map `1` ID to `4294967294`(max id -1) ID in the container namespace with the exactly same IDs on host user namespace

In triplet format, something like this:  
**[0 4294967295 1] [1 1 4294967294]**

With these mappings purpose of Node-Level usernamespace will be solved and it will also satisfy any values of `RunAsUser`, `RunAsGroup`, `SupplementalGroups` and `FsGroup` in the older/existing templates.


## References
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
We can extend current Node-Level user-namespace support in Kubernetes to support Namespace-level isolation(or if desired even pod-level isolation) by dividing and allocating learned mapping from runtime among Kubernetes namespaces (or pods, if desired). From end-user UI perspective, we don't expect any change in the UI related to user namespaces support.


### Remote Volumes
Remote Volumes support should be investigated and should be targeted in future once support is there at lower infra layers.


## Risks and Mitigations
The main risk with this change stems from the fact that processes in Pods will run with different “real” uids than they used to, while expecting the original uids to make operations on the Nodes or consistently access shared persistent storage.
- This can be mitigated by turning the feature on gradually, per-Pod or per Kubernetes namespace.
- For the Kubernetes' cluster Pods (that provide the Kubernetes functionality), testing of their behaviour and ability to run in user namespaced setups is crucial.

## Graduation Criteria
- PSP integration
- e2e tests

## Alternatives
User Namespace mappings can be passed explicitly through kubelet flags similar to https://github.com/kubernetes/kubernetes/pull/55707 but we do not prefer this option because this is very much prone to mis-configuration.
