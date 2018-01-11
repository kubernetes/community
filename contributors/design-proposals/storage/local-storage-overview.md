# Local Storage Management
Authors: vishh@, msau42@

This document presents a strawman for managing local storage in Kubernetes. We expect to provide a UX and high level design overview for managing most user workflows. More detailed design and implementation will be added once the community agrees with the high level design presented here.

# Goals
* Enable ephemeral & durable access to local storage
* Support storage requirements for all workloads supported by Kubernetes
* Provide flexibility for users/vendors to utilize various types of storage devices
* Define a standard partitioning scheme for storage drives for all Kubernetes nodes
* Provide storage usage isolation for shared partitions
* Support random access storage devices only (e.g., hard disks and SSDs)

# Non Goals
* Provide storage usage isolation for non-shared partitions.
* Support all storage devices natively in upstream Kubernetes. Non standard storage devices are expected to be managed using extension mechanisms.
* Support for I/O isolation using CFS & blkio cgroups.
  * IOPS isn't safe to be a schedulable resource. IOPS on rotational media is very limited compared to other resources like CPU and Memory. This leads to severe resource stranding.
  * Blkio cgroup + CFS based I/O isolation doesn't provide deterministic behavior compared to memory and cpu cgroups. Years of experience at Google with Borg has taught that relying on blkio or I/O scheduler isn't suitable for multi-tenancy.
  * Blkio cgroup based I/O isolation isn't suitable for SSDs. Turning on CFQ on SSDs will hamper performance. Its better to statically partition SSDs and share them instead of using CFS.
  * I/O isolation can be achieved by using a combination of static partitioning and remote storage. This proposal recommends this approach with illustrations below.
  * Pod level resource isolation extensions will be made available in the Kubelet which will let vendors add support for CFQ if necessary for their deployments.

# Use Cases

## Ephemeral Local Storage
Today, ephemeral local storage is exposed to pods via the container’s writable layer, logs directory, and EmptyDir volumes.  Pods use ephemeral local storage for scratch space, caching and logs.  There are many issues related to the lack of local storage accounting and isolation, including:

* Pods do not know how much local storage is available to them.
* Pods cannot request “guaranteed” local storage.
* Local storage is a “best-effort” resource.
* Pods can get evicted due to other pods filling up the local storage, after which no new pods will be admitted until sufficient storage has been reclaimed.

## Persistent Local Storage
Distributed filesystems and databases are the primary use cases for persistent local storage due to the following factors:

* Performance: On cloud providers, local SSDs give better performance than remote disks.
* Cost: On bare metal, in addition to performance, local storage is typically cheaper and using it is a necessity to provision distributed filesystems.

Distributed systems often use replication to provide fault tolerance, and can therefore tolerate node failures. However, data gravity is preferred for reducing replication traffic and cold startup latencies.

# Design Overview

A node’s local storage can be broken into primary and secondary partitions.

## Primary Partitions
Primary partitions are shared partitions that can provide ephemeral local storage.  The two supported primary partitions are:

### Root
This partition holds the kubelet’s root directory (`/var/lib/kubelet` by default) and `/var/log` directory. This partition may be shared between user pods, OS and Kubernetes system daemons. This partition can be consumed by pods via EmptyDir volumes, container logs, image layers and container writable layers. Kubelet will manage shared access and isolation of this partition. This partition is “ephemeral” and applications cannot expect any performance SLAs (Disk IOPS for example) from this partition.

### Runtime
This is an optional partition which runtimes can use for overlay filesystems. Kubelet will attempt to identify and provide shared access along with isolation to this partition. Container image layers and writable later is stored here. If the runtime partition exists, `root` partition will not hold any image layer or writable layers.

## Secondary Partitions
All other partitions are exposed as local persistent volumes. The PV interface allows for varying storage configurations to be supported, while hiding specific configuration details to the pod.  All the local PVs can be queried and viewed from a cluster level using the existing PV object.  Applications can continue to use their existing PVC specifications with minimal changes to request local storage.

The local PVs can be precreated by an addon DaemonSet that discovers all the secondary partitions at well-known directories, and can create new PVs as partitions are added to the node.  A default addon can be provided to handle common configurations.

Local PVs can only provide semi-persistence, and are only suitable for specific use cases that need performance, data gravity and can tolerate data loss.  If the node or PV fails, then either the pod cannot run, or the pod has to give up on the local PV and find a new one.  Failure scenarios can be handled by unbinding the PVC from the local PV, and forcing the pod to reschedule and find a new PV.

Since local PVs are only accessible from specific nodes, the scheduler needs to take into account a PV's node constraint when placing pods.  This can be generalized to a storage topology constraint, which can also work with zones, and in the future: racks, clusters, etc.

The term `Partitions` are used here to describe the main use cases for local storage. However, the proposal doesn't require a local volume to be an entire disk or a partition - it supports arbitrary directory.  This implies that cluster administrator can create multiple local volumes in one partition, each has the capacity of the partition, or even create local volume under primary partitions. Unless strictly required, e.g. if you have only one partition in your host, this is strongly discouraged.  For this reason, following description will use `partition` or `mount point` exclusively.

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
        storage.kubernetes.io/overlay: 100Gi
        storage.kubernetes.io/scratch: 100Gi
      allocatable:
        storage.kubernetes.io/overlay: 100Gi
        storage.kubernetes.io/scratch: 90Gi
    ```

2. Alice adds new storage resource requirements to her pod, specifying limits for the container's writeable and overlay layers, and emptyDir volumes.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       resources:
         limits:
           storage.kubernetes.io/logs: 500Mi
           storage.kubernetes.io/overlay: 1Gi
       volumeMounts:
       - name: myEmptyDir
         mountPath: /mnt/data
     volumes:
     - name: myEmptyDir
       emptyDir:
         sizeLimit: 20Gi
    ```

3. Alice’s pod “foo” is Guaranteed a total of “21.5Gi” of local storage. The container “fooc” in her pod cannot consume more than 1Gi for writable layer and 500Mi for logs, and “myEmptyDir” volume cannot consume more than 20Gi.
4. For the pod resources, `storage.kubernetes.io/logs` resource is meant for logs. `storage.kubernetes.io/overlay` is meant for writable layer.
5. `storage.kubernetes.io/logs` is satisfied by `storage.kubernetes.io/scratch`.
6. `storage.kubernetes.io/overlay` resource can be satisfied by `storage.kubernetes.io/overlay` if exposed by nodes or by `storage.kubernetes.io/scratch` otherwise. The scheduler follows this policy to find an appropriate node which can satisfy the storage resource requirements of the pod.
7. EmptyDir.size is both a request and limit that is satisfied by `storage.kubernetes.io/scratch`.
8. Kubelet will rotate logs to keep scratch space usage of “fooc” under 500Mi
9. Kubelet will track the usage of pods across logs and overlay filesystem and restart the container if it's total usage exceeds it's storage limits. If usage on `EmptyDir` volume exceeds its `limit`, then the pod will be evicted by the kubelet. By performing soft limiting, users will be able to easily identify pods that run out of storage.
10. Primary partition health is monitored by an external entity like the “Node Problem Detector” which is expected to place appropriate taints.
11. If a primary partition becomes unhealthy, the node is tainted and all pods running in it will be evicted by default, unless they tolerate that taint. Kubelet’s behavior on a node with unhealthy primary partition is undefined. Cluster administrators are expected to fix unhealthy primary partitions on nodes.

### Bob runs batch workloads and is unsure of “storage” requirements

1. Bob can create pods without any “storage” resource requirements.

    ```yaml
    apiVersion: v1
    kind: Pod
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
         storage.kubernetes.io/logs: 200Mi
         storage.kubernetes.io/overlay: 200Mi
         type: Container
       - default:
         sizeLimit: 1Gi
         type: EmptyDir
    ```

3. The limit range will update the pod specification as follows:

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       resources:
         limits:
           storage.kubernetes.io/logs: 200Mi
           storage.kubernetes.io/overlay: 200Mi
       volumeMounts:
       - name: myEmptyDir
         mountPath: /mnt/data
     volumes:
     - name: myEmptyDir
       emptyDir:
         sizeLimit: 1Gi
    ```

4. Bob’s “foo” pod can use upto “200Mi” for its containers logs and writable layer each, and “1Gi” for its “myEmptyDir” volume.
5. If Bob’s pod “foo” exceeds the “default” storage limits and gets evicted, then Bob can set a minimum storage requirement for his containers and a higher `sizeLimit` for his EmptyDir volumes.

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
   name: foo
  spec:
   containers:
   - name: fooc
     resources:
       requests:
         storage.kubernetes.io/logs: 500Mi
         storage.kubernetes.io/overlay: 500Mi
     volumeMounts:
     - name: myEmptyDir
       mountPath: /mnt/data
   volumes:
   - name: myEmptyDir
     emptyDir:
       sizeLimit: 2Gi
  ```

6. It is recommended to require `limits` to be specified for `storage` in all pods. `storage` will not affect the `QoS` Class of a pod since no SLA is intended to be provided for storage capacity isolation. It is recommended to use Persistent Volumes as much as possible and avoid primary partitions.

### Alice manages a Database which needs access to “durable” and fast scratch space

1. Cluster administrator provisions machines with local SSDs and brings up the cluster
2. When a new node instance starts up, an addon DaemonSet discovers local “secondary” partitions which are mounted at a well known location and creates Local PVs for them if one doesn’t exist already. The PVs will include a path to the secondary device mount points, and a node affinity ties the volume to a specific node.  The node affinity specification tells the scheduler to filter PVs with the same affinity key/value on the node.  For the local storage case, the key is `kubernetes.io/hostname`, but the same mechanism could be used for zone constraints as well.

    ```yaml
    kind: StorageClass
    apiVersion: storage.k8s.io/v1
    metadata:
      name: local-fast
    topologyKey: kubernetes.io/hostname
    ```
    ```yaml
    kind: PersistentVolume
    apiVersion: v1
    metadata:
      name: local-pv-1
    spec:
      nodeAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - key: kubernetes.io/hostname
              operator: In
              values:
              - node-1
      capacity:
        storage: 100Gi
      local:
        path: /var/lib/kubelet/storage-partitions/local-pv-1
      accessModes:
        - ReadWriteOnce
      persistentVolumeReclaimPolicy: Delete
      storageClassName: local-fast
    ```
    ```
    $ kubectl get pv
    NAME       CAPACITY ACCESSMODES RECLAIMPOLICY STATUS    CLAIM … NODE
    local-pv-1 100Gi    RWO         Delete        Available         node-1
    local-pv-2 10Gi     RWO         Delete        Available         node-1
    local-pv-1 100Gi    RWO         Delete        Available         node-2
    local-pv-2 10Gi     RWO         Delete        Available         node-2
    local-pv-1 100Gi    RWO         Delete        Available         node-3
    local-pv-2 10Gi     RWO         Delete        Available         node-3
    ```
3. Alice creates a StatefulSet that requests local storage from StorageClass "local-fast".  The PVC will only be bound to PVs that match the StorageClass name.

    ```yaml
    apiVersion: apps/v1beta1
    kind: StatefulSet
    metadata:
      name: web
    spec:
      serviceName: "nginx"
      replicas: 3
      template:
        metadata:
          labels:
            app: nginx
        spec:
          terminationGracePeriodSeconds: 10
          containers:
          - name: nginx
            image: k8s.gcr.io/nginx-slim:0.8
            ports:
            - containerPort: 80
              name: web
            volumeMounts:
            - name: www
              mountPath: /usr/share/nginx/html
            - name: log
              mountPath: /var/log/nginx
      volumeClaimTemplates:
      - metadata:
          name: www
        spec:
          accessModes: [ "ReadWriteOnce" ]
          storageClassName: local-fast
          resources:
            requests:
              storage: 100Gi
      - metadata:
          name: log
        spec:
          accessModes: [ "ReadWriteOnce" ]
          storageClassName: local-slow
          resources:
            requests:
              storage: 1Gi
    ```

4. The scheduler identifies nodes for each pod that can satisfy all the existing predicates.
5. The nodes list is further filtered by looking at the PVC's StorageClass, and checking if there is available PV of the same StorageClass on a node.
6. The scheduler chooses a node for the pod based on a ranking algorithm.
7. Once the pod is assigned to a node, then the pod’s local PVCs get bound to specific local PVs on the node.

    ```
    $ kubectl get pvc
    NAME            STATUS VOLUME     CAPACITY ACCESSMODES … NODE
    www-local-pvc-1 Bound  local-pv-1 100Gi    RWO           node-1
    www-local-pvc-2 Bound  local-pv-1 100Gi    RWO           node-2
    www-local-pvc-3 Bound  local-pv-1 100Gi    RWO           node-3
    log-local-pvc-1 Bound  local-pv-2 10Gi     RWO           node-1
    log-local-pvc-2 Bound  local-pv-2 10Gi     RWO           node-2
    log-local-pvc-3 Bound  local-pv-2 10Gi     RWO           node-3
    ```
    ```
    $ kubectl get pv
    NAME       CAPACITY … STATUS    CLAIM           NODE
    local-pv-1 100Gi      Bound     www-local-pvc-1 node-1
    local-pv-2 10Gi       Bound     log-local-pvc-1 node-1
    local-pv-1 100Gi      Bound     www-local-pvc-2 node-2
    local-pv-2 10Gi       Bound     log-local-pvc-2 node-2
    local-pv-1 100Gi      Bound     www-local-pvc-3 node-3
    local-pv-2 10Gi       Bound     log-local-pvc-3 node-3
    ```

8. If a pod dies and is replaced by a new one that reuses existing PVCs, the pod will be placed on the same node where the corresponding PVs exist. Stateful Pods are expected to have a high enough priority which will result in such pods preempting other low priority pods if necessary to run on a specific node.
9. Forgiveness policies can be specified as tolerations in the pod spec for each failure scenario.  No toleration specified means that the failure is not tolerated.  In that case, the PVC will immediately be unbound, and the pod will be rescheduled to obtain a new PV.  If a toleration is set, by default, it will be tolerated forever.  `tolerationSeconds` can be specified to allow for a timeout period before the PVC gets unbound.

  Node taints already exist today.  Pod scheduling failures are specified separately as a timeout.
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: foo
  spec:
    <snip>
    nodeTolerations:
      - key: node.alpha.kubernetes.io/not-ready
        operator: TolerationOpExists
        tolerationSeconds: 600
      - key: node.alpha.kubernetes.io/unreachable
        operator: TolerationOpExists
        tolerationSeconds: 1200
    schedulingFailureTimeoutSeconds: 600
  ```

  A new PV taint will be introduced to handle unhealthy volumes.  The addon or another external entity can monitor the volumes and add a taint when it detects that it is unhealthy.
  ```yaml
  apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    name: foo
  spec:
    <snip>
    pvTolerations:
      - key: storage.kubernetes.io/pvUnhealthy
        operator: TolerationOpExists
  ```
10. Once Alice decides to delete the database, she destroys the StatefulSet, and then destroys the PVCs.  The PVs will then get deleted and cleaned up according to the reclaim policy, and the addon adds it back to the cluster.

### Bob manages a distributed filesystem which needs access to all available storage on each node

1. The cluster that Bob is using is provisioned with nodes that contain one or more secondary partitions
2. The cluster administrator runs a DaemonSet addon that discovers secondary partitions across all nodes and creates corresponding PVs for them.
3. The addon will monitor the health of secondary partitions and mark PVs as unhealthy whenever the backing local storage devices have failed.
4. Bob creates a specialized controller (Operator) for his distributed filesystem and deploys it.
5. The operator will identify all the nodes that it can schedule pods onto and discovers the PVs available on each of those nodes. The operator has a label selector that identifies the specific PVs that it can use (this helps preserve fast PVs for Databases for example).
6. The operator will then create PVCs and manually bind to individual local PVs across all its nodes.
7. It will then create pods, manually place them on specific nodes (similar to a DaemonSet) with high enough priority and have them use all the PVCs created by the Operator on those nodes.
8. If a pod dies, it will get replaced with a new pod that uses the same set of PVCs that the old pod had used.
9. If a PV gets marked as unhealthy, the Operator is expected to delete pods if they cannot tolerate device failures

### Phippy manages a cluster and intends to mitigate storage I/O abuse

1. Phippy creates a dedicated partition with a separate device for her system daemons. She achieves this by making `/var/log/containers`, `/var/lib/kubelet`, `/var/lib/docker` (with the docker runtime) all reside on a separate partition.
2. Phippy is aware that pods can cause abuse to each other.
3. Whenever a pod experiences I/O issues with it's EmptyDir volume, Phippy reconfigures those pods to use an inline Persistent Volume, whose lifetime is tied to the pod.
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       resources:
       limits:
         storage.kubernetes.io/logs: 500Mi
         storage.kubernetes.io/overlay: 1Gi
       volumeMounts:
       - name: myEphemeralPersistentVolume
         mountPath: /mnt/tmpdata
     volumes:
     - name: myEphemeralPersistentVolume
       inline:
         spec:
           accessModes: [ "ReadWriteOnce" ]
           storageClassName: local-fast
           resources:
             limits:
               size: 1Gi
    ```

4. Phippy notices some of her pods are experiencing spurious downtimes. With the help of monitoring (`iostat`), she notices that the nodes pods are running on are overloaded with I/O operations. She then updates her pods to use Logging Volumes which are backed by persistent storage. If a logging volumeMount is associated with a container, Kubelet will place log data from stdout & stderr of the container under the volume mount path within the container. Kubelet will continue to expose stdout/stderr log data to external logging agents using symlinks as it does already.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       volumeMounts:
       - name: myLoggingVolume
         mountPath: /var/log/
         policy:
           logDir:
             subDir: foo
             glob: *.log
     - name: barc
       volumeMounts:
       - name: myInMemoryLoggVolume
         mountPath: /var/log/
         policy:
           logDir:
             subDir: bar
             glob: *.log
    volumes:
    - name: myLoggingVolume
      inline:
        spec:
          accessModes: [ "ReadWriteOnce" ]
          storageClassName: local-slow
          resources:
            requests:
              storage: 1Gi
    - name: myInMemoryLogVolume
      emptyDir:
        medium: memory
        resources:
          limits:
            size: 100Mi
    ```

5. Phippy notices some of her pods are suffering hangs by while writing to their writable layer. Phippy again notices that I/O contention is the root cause and then updates her Pod Spec to use memory backed or persistent volumes for her pods writable layer. Kubelet will instruct the runtimes to overlay the volume with `overlay` policy over the writable layer of the container.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
     name: foo
    spec:
     containers:
     - name: fooc
       volumeMounts:
       - name: myWritableLayer
         policy:
           overlay:
             subDir: foo
     - name: barc
       volumeMounts:
       - name: myDurableWritableLayer
         policy:
           overlay:
               subDir: bar
     volumes:
     - name: myWritableLayer
       emptyDir:
         medium: memory
         resources:
           limits:
             storage: 100Mi
     - name: myDurableWritableLayer
       inline:
         spec:
           accessModes: [ "ReadWriteOnce" ]
           storageClassName: local-fast
           resources:
             requests:
               storage: 1Gi
    ```

### Bob manages a specialized application that needs access to Block level storage
Note: Block access will be considered as a separate feature because it can work for both remote and local storage.  The examples here are a suggestion on how such a feature can be applied to this local storage model, but is subject to change based on the final design for block access.

1. The cluster that Bob uses has nodes that contain raw block devices that have not been formatted yet.
2. The same addon DaemonSet can also discover block devices and creates corresponding PVs for them with the `volumeType: block` spec. `path` is overloaded here to mean both fs path and block device path.

    ```yaml
    kind: PersistentVolume
    apiVersion: v1
    metadata:
      name: foo
      labels:
        kubernetes.io/hostname: node-1
    spec:
      capacity:
        storage: 100Gi
      volumeType: block
      local:
        path: /var/lib/kubelet/storage-raw-devices/foo
      accessModes:
        - ReadWriteOnce
      persistentVolumeReclaimPolicy: Delete
      storageClassName: local-fast
    ```

3. Bob creates a pod with a PVC that requests for block level access and similar to a Stateful Set scenario the scheduler will identify nodes that can satisfy the pods request.  The block devices will not be formatted to allow the application to handle the device using their own methods.

    ```yaml
    kind: PersistentVolumeClaim
    apiVersion: v1
    metadata:
      name: myclaim
    spec:
      volumeType: block
      storageClassName: local-fast
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 80Gi
    ```

4. It is also possible for a PVC that requests `volumeType: block` to also use file-based volume.  In this situation, the block device would get formatted with the filesystem type specified in the PVC spec.  And when the PVC gets destroyed, then the filesystem also gets destroyed to return back to the original block state.

    ```yaml
    kind: PersistentVolumeClaim
    apiVersion: v1
    metadata:
      name: myclaim
    spec:
      volumeType: block
      fsType: ext4
      storageClassName: local-fast
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 80Gi
    ```

*The lifecycle of the block level PV will be similar to that of the scenarios explained earlier.*

# Open Questions & Discussion points
* Single vs split “limit” for storage across writable layer and logs
    * Split allows for enforcement of hard quotas
    * Single is a simpler UI
* Local Persistent Volume bindings happening in the scheduler vs in PV controller
    * Should the PV controller fold into the scheduler
	* This will help spread PVs and pods across matching zones.
* Repair/replace scenarios.
    * What are the implications of removing a disk and replacing it with a new one?
    * We may not do anything in the system, but may need a special workflow
* Volume-level replication use cases where there is no pod associated with a volume.  How could forgiveness/data gravity be handled there?

# Related Features
* Support for encrypted secondary partitions in order to make wiping more secure and reduce latency
* Co-locating PVs and pods across zones. Binding PVCs in the scheduler will help with this feature.

# Recommended Storage best practices
* Have the primary partition on a reliable storage device
* Have a dedicated storage device for system daemons.
* Consider using RAID and SSDs (for performance)
* Partition the rest of the storage devices based on the application needs
    * SSDs can be statically partitioned and they might still meet IO requirements of apps.
    * TODO: Identify common durable storage requirements for most databases
* Avoid having multiple logical partitions on hard drives to avoid IO isolation issues
* Run a reliable cluster level logging service to drain logs from the nodes before they get rotated or deleted
* The runtime partition for overlayfs is optional. You do not **need** one.
* Alert on primary partition failures and act on it immediately. Primary partition failures will render your node unusable.
* Use EmptyDir for all scratch space requirements of your apps when IOPS isolation is not of concern.
* Make the container’s writable layer `readonly` if possible.
* Another option is to keep the writable layer on tmpfs. Such a setup will allow you to eventually migrate from using local storage for anything but super fast caching purposes or distributed databases leading to higher reliability & uptime for nodes.

# FAQ

### Why is the kubelet managing logs?

Kubelet is managing access to shared storage on the node. Container logs outputted via it's stdout and stderr ends up on the shared storage that kubelet is managing. So, kubelet needs direct control over the log data to keep the containers running (by rotating logs), store them long enough for break glass situations and apply different storage policies in a multi-tenant cluster. All of these features are not easily expressible through external logging agents like journald for example.


### Master are upgraded prior to nodes. How should storage as a new compute resource be rolled out on to existing clusters?

Capacity isolation of shared partitions (ephemeral storage) will be controlled using a feature gate. Do not enable this feature gate until all the nodes in a cluster are running a kubelet version that supports capacity isolation.
Since older kubelets will not surface capacity of shared partitions, the scheduler will ignore those nodes when attempting to schedule pods that request storage capacity explicitly.


### What happens if storage usage is unavailable for writable layer?

Kubelet will attempt to enforce capacity limits on a best effort basis. If the underlying container runtime cannot surface usage metrics for the writable layer, then kubelet will not provide capacity isolation for the writable layer.


### Are LocalStorage PVs required to be a whole partition?

No, but it is the recommended way to ensure capacity and performance isolation.  For HDDs, a whole disk is recommended for performance isolation.  In some environments, multiple storage partitions are not available, so the only option is to share the same filesystem.  In that case, directories in the same filesystem can be specified, and the administrator could configure group quota to provide capacity isolation.

# Features & Milestones

#### Features with owners

1. Support for durable Local PVs
2. Support for capacity isolation

Alpha support for these two features are targeted for v1.7. Beta and GA timelines are TBD.
Currently, msau42@, jinxu@ and vishh@ will be developing these features.

#### Features needing owners

1. Support for persistent volumes tied to the lifetime of a pod (`inline PV`)
2. Support for Logging Volumes
3. Support for changing the writable layer type of containers
4. Support for Block Level Storage
