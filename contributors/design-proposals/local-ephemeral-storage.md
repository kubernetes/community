# Local Ephemeral Storage Resource Management
Authors: jinxu@ vishh@

Currently Kubernetes does not support storage resource usage guarantee and isolation like compute resources such as CPU and memory. This doc details an effort for improving Storage Resource Management in Kubernetes with focus on
capacity isolation for local ephemeral storage. 


# Goals
* Support ephemeral storage usage isolation for shared root partition among
  containers and pods
* Provide ephemeral storage requirement and garantee for containers and pods
* Add quota management for ephemeral storage for shared root parition
* Add limitrange support for ephemeral storage

# Non Goals
* Provide storage usage isolation for non-shared partitions.
* Support for I/O isolation using CFS & blkio cgroups.
  * IOPS isn't safe to be a schedulable resource. IOPS on rotational media is very limited compared to other resources like CPU and Memory. This leads to severe resource stranding.
  * Blkio cgroup + CFS (Completely Fair Scheduler) based I/O isolation doesn't provide deterministic behavior compared to memory and cpu cgroups. Years of experience at Google with Borg has taught that relying on blkio or I/O scheduler isn't suitable for multi-tenancy.
  * Blkio cgroup based I/O isolation isn't suitable for SSDs. Turning on CFQ on SSDs will hamper performance. Its better to statically partition SSDs and share them instead of using CFS.
  * I/O isolation can be achieved by using a combination of static partitioning and remote storage. This proposal recommends this approach with illustrations below.

# Problems
Today, ephemeral local storage is exposed to pods via the container’s writable layer, logs directory, and EmptyDir volumes.  Pods use ephemeral local storage for scratch space, caching and logs.  There are many issues related to the lack of local storage accounting and isolation, including:

* Pods do not know how much local storage is available to them.
* Pods cannot request “guaranteed” local storage.
* Local storage is a “best-effort” resource.
* Pods can get evicted due to other pods filling up the local storage, after which no new pods will be admitted until sufficient storage has been reclaimed.
* There is no quota control for local ephemeral storage.

# Backgroud
## Local Storage Configuration
A node’s local storage can be broken into primary and secondary partitions.

## Primary Partitions
Primary partitions are shared partitions that can provide ephemeral local storage.  The two supported primary partitions are:

### Root
This partition holds the kubelet’s root directory (`/var/lib/kubelet` by default) and `/var/log` directory. This partition may be shared between user pods, OS and Kubernetes system daemons. This partition can be consumed by pods via EmptyDir volumes, container logs, image layers and container writable layers. Kubelet will manage shared access and isolation of this partition. This partition is “ephemeral” and applications cannot expect any performance SLAs (Disk IOPS for example) from this partition.

### Runtime
This is an optional partition which runtimes can use for overlay filesystems. This feature will not attempt to provide resource garantee and isolation to this partition.

## Secondary Partitions
All other partitions are exposed as local persistent volumes which are not covered in this feature.

## Resource Management Solution for CPU and Memory
Currently kubernetes provides different levels of resource management for CPU and memory resource types. 

* Node-level: Node allocatable resources is for providing isolation between system daemons and user pods to ensure resources for system daemons. kubelet exposes a feature named Node Allocatable that helps to reserve compute resources for system daemons.  Enforcement is performed by evicting pods whenever the overall usage across all pods exceeds Allocatable.
* Container-level: For each resource, containers specify a request, which is the amount of that resource that the system will guarantee to the container, and a limit which is the maximum amount that the system will allow the container to use. Therefore, it can provide resource guarantee and isolation among containers.
* Pod-level: Although currently user cannot specify pod-level resource request/limit directly, controller uses the sum of all containers’ limit as the pod-level constraints. The total usage of the pod could not exceed this pod-level limit. This provides isolation among pods.
* Namespace level: A resource quota, defined by a ResourceQuota object, provides limit on aggregated resource request/limit per namespace. If quota is enabled in a namespace for compute resources like cpu and memory, users must specify requests or limits for their pods; otherwise, the quota system may reject pod creation.

# Design Overview
This feature adds only one storage API to represent ephemeral storage from root partition and manage its isolation. The management of local ephemeral storage is consistent with memory management.
```
// Local ephemeral storage for root partition
ResourceEphemeralStorage ResourceName = "ephemeral-storage"
```

## Resource Request/Limit
* Container-level resource requirement: Similar to CPU and memory, container spec can specify the request/limit for local ephemeral storage. All the validation related to the storage resource requirements will be the same as memory resource. As mentioned above in  Storage types and configuration, kubernetes support optional imagefs partition. In the case of two partition, only the storage usage in root partition will be managed and isolated across different containers. (In another word, container writable layer usage will not be counted for container-level eviction management)
* Pod-level resource constraint: Since we haven’t supported pod-level resource API, the sum of the resource request/limit is considered as pod-level resource requirement. Similar to memory, all local ephemeral storage resource usage is subject to this requirement. For example, emptyDir disk usage (also secrets, configuMap, downwardAPI, gitRepo since they are wrapped emptyDir volume) plus containers disk usage should not exceed the the pod-level resource requirement (the sum of all container’s limit).
* EmptyDir SizeLimit: Because empryDir volume is a pod-level resource which is managed separately, we also add a sizeLimit for emptyDir Volume for additional storage isolation. If this limit is set for emptyDir volume (default medium), eviction manager will validate this limit with emptyDir usage too in additional to the above pod-level resource constraints. 

## Eviction Policy and Scheduler Predicates
CPU and memory use cgroup for limiting the resource usage. However, cgroup for disk usage is not available. Our current design is to evict pods when exceeding the limit set for local storage. The eviction policy is listed as follows.
* If the container writable layer (overlay) usage exceeds its limit, pod gets evicted.
* If the emptyDir volume usage exceeds the sizeLimit, pod gets evicted.
* If the sum of the usage from emptyDir and all contains exceeds the sum of the container’s local storage limits (pod-level limit), pod gets evicted.

When scheduler admits pods, it sums up the storage requests from each container and pod can be scheduled only if the sum is smaller than the allocatable of the local storage space.

## Resource Quota
Add two more resource quotas for storage. The request and limit set constraints on the total requests/limits of all containers’ in a namespace 
* ResourceRequestsEphemeralStorage
* ResourceLimitsEphemeralStorage

## LimitRange
Similar to CPU and memory, admin could use LimitRange to set default container’s local storage request/limit, and/or minimum/maximum resource constraints for a namespace. 


# User Workflows

### Alice manages a deployment and requires “Guaranteed” ephemeral storage

1. Kubelet running across all nodes will identify primary partition and expose capacity and allocatable for the primary partitions.  This allows primary partitions' storage capacity to be considered as a first class resource when scheduling.

    ```yaml
    apiVersion: v1
    kind: Node
    metadata:
      name: foo
    status:
      capacity:
        ephemeral-storage: 100Gi
      allocatable:
        ephemeral-storage: 90Gi
    ```

2. Alice adds new storage resource requirements to her pod.

    ```yaml
    apiVersion: v1
    kind: pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooa
       resources:
        requests:
           ephemeral-storage: 10Gi
         limits:
           ephemeral-storage: 10Gi
     - name: foob
       resources:
        requests:
           ephemeral-storage: 20Gi
         limits:
           ephemeral-storage: 20Gi
       volumeMounts:
       - name: myEmptyDir
         mountPath: /mnt/data
     volumes:
     - name: myEmptyDir
       emptyDir:
         sizeLimit: 5Gi
    ```

3. Alice’s pod “foo” is Guaranteed a total of “30Gi” of local ephemeral storage. The container “fooa” in her pod cannot consume more than 10Gi for local ephemeral storage (could be used by container writable layer and logs), and container “foob” cannot consume more than 20Gi for local ephemeral storag. 
4. EmptyDir.sizeLimit is both a request and limit. So “myEmptyDir” volume is garanteed to have 5Gi storage and at the same time it cannot consume more than 5Gi.
5. The total ephemeral storage is limited by 10+20=30Gi for pod "foo". Kubelet will track the usage of pods across containers and emptyDir volumes if it's total usage exceeds it's storage limits. If usage exceeds its `limit`, then the pod will be evicted by the kubelet. By performing soft limiting, users will be able to easily identify pods that run out of storage.
6. Primary partition health is monitored by an external entity like the “Node Problem Detector” which is expected to place appropriate taints.
7. If a primary partition becomes unhealthy, the node is tainted and all pods running in it will be evicted by default, unless they tolerate that taint. Kubelet’s behavior on a node with unhealthy primary partition is undefined. Cluster administrators are expected to fix unhealthy primary partitions on nodes.

### Bob runs batch workloads and is unsure of “storage” requirements

1. Bob can create pods without any “storage” resource requirements.

    ```yaml
    apiVersion: v1
    kind: pod
    metadata:
     name: foo
     namespace: myns
    spec:
     containers:
     - name: fooc
       volumeMounts:
       - name: myEmptyDir
         mountPath: /mnt/data
     volumes:
     - name: myEmptyDir
       emptyDir:
    ```

2. His cluster administrator being aware of the issues with disk reclamation latencies has intelligently decided not to allow overcommitting primary partitions. The cluster administrator has installed a [LimitRange](https://kubernetes.io/docs/user-guide/compute-resources/) to “myns” namespace that will set a default storage size. Note: A cluster administrator can also specify burst ranges and a host of other features supported by LimitRange for local storage.

    ```yaml
    apiVersion: v1
    kind: LimitRange
    metadata:
      name: mylimits
    spec:
       - default:
         ephemeral-storage: 2Gi
         type: Container
    ```

3. The limit range will update the pod specification as follows:

    ```yaml
    apiVersion: v1
    kind: pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       resources:
         limits:
           ephemeral-storage: 2Gi
    ```

4. Bob’s “foo” pod can use upto “2Gi” for its containers. 
5. If Bob’s pod “foo” exceeds the “default” storage limits and gets evicted, then Bob can set a minimum storage requirement for his containers and a higher `sizeLimit` for his EmptyDir volumes.

  ```yaml
  apiVersion: v1
  kind: pod
  metadata:
   name: foo
  spec:
   containers:
   - name: fooc
     resources:
       requests:
         ephemeral-storage: 10Gi
       limits:
         ephemeral-storage: 10Gi
     volumeMounts:
     - name: myEmptyDir
       mountPath: /mnt/data
   volumes:
   - name: myEmptyDir
     emptyDir:
       sizeLimit: 5Gi
  ```

6. It is recommended to require `limits` to be specified for `storage` in all pods. `storage` will not affect the `QoS` Class of a pod since no SLA is intended to be provided for storage capacity isolation. It is recommended to use Persistent Volumes as much as possible and avoid primary partitions.


# FAQ

### Why is the kubelet managing logs?

Kubelet is managing access to shared storage on the node. Container logs outputted via it's stdout and stderr ends up on the shared storage that kubelet is managing. So, kubelet needs direct control over the log data to keep the containers running (by rotating logs), store them long enough for break glass situations and apply different storage policies in a multi-tenent cluster. All of these features are not easily expressible through external logging agents like journald for example.


### Master are upgraded prior to nodes. How should storage as a new compute resource be rolled out on to existing clusters?

Capacity isolation of shared partitions (ephemeral storage) will be controlled using a feature gate. Do not enable this feature gate until all the nodes in a cluster are running a kubelet version that supports capacity isolation.
Since older kubelets will not surface capacity of shared partitions, the scheduler will ignore those nodes when attempting to schedule pods that request storage capacity explicitly.


### What happens if storage usage is unavailable for writable layer?

Kubelet will attempt to enforce capacity limits on a best effort basis. If the underlying container runtime cannot surface usage metrics for the writable layer, then kubelet will not provide capacity isolation for the writable layer.


### Are LocalStorage PVs required to be a whole partition?

No, but it is the recommended way to ensure capacity and performance isolation.  For HDDs, a whole disk is recommended for performance isolation.  In some environments, multiple storage partitions are not available, so the only option is to share the same filesystem.  In that case, directories in the same filesystem can be specified, and the adminstrator could configure group quota to provide capacity isolation.

