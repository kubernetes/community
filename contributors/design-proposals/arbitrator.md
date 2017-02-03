# Policy based resource sharing of Kubernetes

@k82cn

January, 2017

## User Scenario and Requirements:

As a cluster admin, I’d like to build an environment to run different workload together, e.g. long running service, bigdata.
 As those applications are managed by different department, I have to provide resource guarantee to each applications, 
 demonstrated as following:

* Long running service (app area) and bigdata (bigdata area) can share resources:
  - Define resource usage of each area, e.g. 40% resources to app area, 60% to bigdata area.
  - Borrow/lending protocol: if the resources is idle in one area, it can be lend out and be preempted back
* Run multiple cluster in bigdata area:
  - Define resources usage of each cluster within bigdata area, e.g. Spark, Hadoop
  - Sharing resources between those big data clusters, e.g. borrow/lending protocol

The detail of requirements for the "bigdata" are

* run a set of applications
* provided each application guaranteed access to some quantity of resources
* provided all applications best-effort access to all unused resources according to some target weight (one weight assigned
  to each application, i.e. if all applications wanted to use all free resources, then they would be allowed to do so 
  in some relative proportion)
* If some application A is using less than its guarantee, and then if it decides to use its guarantee and there aren't 
  enough free resources to do so, it should be able to evict tasks from some other application or applications
  (that is/are using more than their guarantee) in order to obtain its guarantee

Further, group "bigdata" apps and "service" apps into two buckets, providing each bucket (in aggregate) guaranteed 
access to some fraction of the cluster, and best-effort access to the entire cluster with the understanding that usage 
above the guarantee can be revoked at any time. 

According to this [mesos-style.md](https://github.com/kubernetes/kubernetes/blob/master/docs/devel/mesos-style.md), 
it definitely can build a customized components to do so; but it’s better for Kubernetes to provide such a general feature
 about resource planning, resource management and resource sharing.

## Terminology

* Arbitrator: A new components/binary that allocate resources based on policy (resource allocation).
  The default scheduler still assign Pods to resources (resource assignment) 
* Deserved (Resource): The total number of resources that the arbitrator allocated to the namespace
* Overused: The namespace is overused if it used more resources than deserved

## Context

As Kubernetes growing up, there’s several features on resource QoS and sharing between namespaces.

### Preemption and Re-scheduler:

A pod can be evicted because some other pod needs the resources it is using (preemption). There’s a priority-based preemption 
scheme (each pod has a priority, and pods with higher and possibly equal priority can preempt it; who makes the decision of which
pods to preempt is TBD but could be the default scheduler, the rescheduler, and/or application-specific controllers 
with integrated scheduling functionality). Preemption always uses graceful termination. A priority scheme generally implies that
quota is allocated on a per-priority-level basis, so that applications can be given a limited amount of quota at the highest priority level
and a much larger amount of quota (perhaps even infinite, i.e. the cluster's entire capacity) at lower priority level(s).
And rescheduler evicts pods to enforce cluster-level policies (currently there’s a proto-rescheduler that enforces
the policy "critical pods like Heapster, DNS, etc. should never be blocked from running due to insufficient free resources
in the cluster" but there are many other policies it could enforce). It works by evicting one or more pods to
allow some pending pod(s) to schedule.

The preemption is required to shuffle resources between namespaces; arbitrator will be one of the decision maker who defined “priority”,
 e.g. under-used namespace’s priority is higher than overused namespace’s. The arbitrator will leverage `Eviction` feature for preemption.
The rescheduler will help to avoid blocking critical pods due to insufficient free resources and re-schedule pods for better placement. 
With arbitrator, the `kube-system` has infinite deserved resources: deserved always equals to it's request, the other namespaces 
share the rest resources; no additional impact to "better placement".  

### Workload-specific Controller and ThirdPartyResource
ThirdPartyResource objects are a way to extend the Kubernetes API with a new API object type. 
The new API object type will be given an API endpoint URL and support CRUD operations, and watch API.
You can then create custom objects using this API endpoint. Thanks to mesos-style.md and ThirdPartyResource,
the developer can build workload specific controller with customized objects.

There’s an example at [k82cn/kube-arbitrator](https://github.com/k82cn/kube-arbitrator) which provides resource sharing 
and preemption feature by ThirdPartyResource features.

### Horizontal\/Vertical scaling and Node-level QoS

Node level utilization improvement, no contribution on cluster-level resource plan. But regarding Node-level QoS,
the Pod’s request and limit should be also considered.

## Proposal:
### Overview
To achieve the above requirements, a new binary (k8s-arbitrator) and two ThirdPartyResource (Consumer\/Allocation) is introduced.

The following yaml file demonstrates the definition of `Consumer`, the ThirdPartyResource for arbitrator:

```yaml
apiVersion: kuabe-arbitrator.incubator.k8s.io/v1
kind: Consumer
metadata:
  name: defaults
spec:
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
  reserved:
    limits.memory: 1Gi
```

For each `Consumer` object, it has two fields, `hard` and `reserved`:
* reserved: The reserved section defines the resources that reserved for the namespace. It uses the same resources type with
 "Compute Resource Quota" and "Storage Resource Quota"; it can not exceed resources in hard section.
 If ResourceQuota Admission enabled, it’ll also check whether total “reserved” exceed resources in cluster.
* hard: The hard section defines the max resources that a namespace can use; it can not exceed `Quota.hard` of namespace if any. 

The `Consumer` is created by arbitrator for each namespace and updated by cluster admin if necessary; the arbitrator creates
 the `Customer` with infinite `hard` and "zero" `reserved`, so the namespace shares cluster resource equally by default. 

And the arbitrator will create or update `Allocation` with additional field `deserved`: 
* deserved: Similar to “Used” of Quota, it’s not defined in yaml file but updated by arbitrator. It defines the total resources
that arbitrator allocates to a namespace. It’s not more than hard of Quota, and maybe changed because of workload in namespaces.
* hard/reserved: Copy from `Consumer`; if `Consumer` was updated, it's also updated in next arbitration cycle 

```yaml
apiVersion: kuabe-arbitrator.incubator.k8s.io/v1
kind: Allocation
metadata:
  name: defaults
spec:
  deserved:
    cpu: "1.5"
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
  reserved:
    limits.memory: 1Gi
```

The following figure demonstrate the relationship of hard, reserved and deserved in `Consumer`/`Allocation`.

> NOTES: only __Compute Resource__ and __Storage Resource__ are available for reserved & deserved.

```text
    -------------  <-- hard
    |           |
    |           |
    - - - - - - -  <-- deserved
    |           |
    |           |
    |           |
    -------------  <-- reserved
    |           |
    -------------
```

The `k8s-arbitrator` is a new binary that created/updated `Consumer` and `Allocation`; it'll

1. Calculate deserved resource (`Allocation.deserved`) based on arbitrator’s policy, e.g. DRF, and
  namespace’s request (pending pods in PoC)
1. Evict Pods of overused namespace
1. Provide predicates that “namespace will not exceed Allocation.deserved” as scheduler's HTTPExtender

Meanwhile, `k8s-scheduler` dispatch tasks to host based on its policy, e.g. PodAffinity:
`k8s-arbitrator` own resource allocation, `k8s-scheduler` own resource assignment.

The arbitrator takes DRF as default policy. It will list pods/nodes from `k8s-apiserver`, and calculate 
the deserved resource of each namespace based on DRF algorithm; and then, update the deserved to `Allocation` accordingly.
The default interval is 1s which is configurable. The arbitrator will NOT assign hostname to the deserved resources;
it dependent on scheduler to dispatch tasks on the suitable hosts.

The arbitrator also meets the following requirements:

* Total deserved resource of namespace can not exceed resources in cluster
* The deserved resources can not exceed hard resource in `Consumer`
* The deserved resources can not less than reserved resources if enough resource in cluster

### Preemption:
When workload/quota changed, the deserved resources of each namespace maybe also changed. The “higher” priority pod will
trigger eviction, the following figure demonstrate the case of eviction because of deserved resources.

```text
T1:                 T2:                               T3:
 --------------     -------------- --------------     -------------- --------------
 | Consumer-1 |     | Consumer-1 | | Consumer-2 |     | Consumer-1 | | Consumer-2 |
 |   cpu:2    | ==> |   cpu:1    | |   cpu:0    | ==> |   cpu:1    | |   cpu:1    | 
 |   mem:2    |     |   mem:1    | |   mem:0    |     |   mem:1    | |   mem:1    |
 --------------     -------------- --------------     -------------- --------------
```

* T1: there's only one namespace in the cluster: `Consumer-1`; all resource(cpu:2,mem:2) are allocated to it
* T2: a new namespace `Consuemr-2` is created; the arbitrator re-calculate resource allocation for each namespace,
  shrink the overused namespace
* T3: the controller, who manages overused namespace, MUST select a Pod to evict, or arbitrator will evict randomly.
  After eviction, assign resources to underused namespace


The arbitrator uses “/evict” subresource of pod to reclaim resources. But when arbitrator choose the pods to evict,
here’re at least two requirements:
* After eviction, the of pods can not less than `PodDisruptionBudget`
* After eviction, the resources of namespace can not less than reserved

The namespace maybe underused after eviction; the arbitrator will try to evict pods from most overused namespace.
For the resources fragment issue, it’s out scope of this doc; will design detail in preemption implementation doc.

## Feature Interaction:

### Scheduler
Use scheduler’s HTTPExtender to access arbitrator for predicates based on `Allocation.Deserved`.
Enhance HTTPExtender’s interface to only send `v1.Pod` for performance; arbitrator only predicates the # of resources 
instead of specify host.

### Workload-specific controller
The arbitrator will also evict overused namespace in workload-specific controller. The workload-specific controller can
 not use more resource than `Allocation.deserved`. If `Allocation.deserved` updated, it selects Pods to evict; 
 otherwise, arbitrator will evict pods (e.g. FCFS) after grace period

### Multiple-scheduler
If enable multiple-scheduler, enable only one arbitrator to avoid race condition

### Admission Controller
Arbitrator check `Consumer` definition against __Compute Resource__ and __Storage Resource__ of Quota if any:
the `Consumer.hard` can not exceed `Quota.hard`. The other metrics of ResourceQuotaAdmission will follow current behaviors.

No impact to other admission plugins

### Node level-QoS
In the prototype, only request is considered; limits is considered in backlog

### ReplicaController/ReplicaSet
As `Consumer`/`Allocation` is namespace level instead of Job/RC level, k8s-controller-manager can not create pods 
according to RC’s proportion; it need end user to balance replicas between RCs or request reserved resource in Consumer.

### kubelet

Does not involve _kubelet_ for now, although `Allocation.deserved` of namespace is also a factor for eviction in _kubelet_,
 for example:
* Reject overused scheduler’s request
* Evict most overused scheduler’s request if node is exhaust 

### Current

* Basic DRF arbitration policy (done)
* `Consumer`/`Allocation` creation by `k8s-arbitrator` (on-going)
* Arbitration policy with `Consumer.reserved` (on-going)
* Misc, e.g. doc, test (on-going)

## Roadmap and Future

* Enhance scheduler HTTPExtender to send `v1.Pod` only
* Arbitration Eviction policy, e.g. not break `PodDisruptionBudget`
* “resourceRequest” for controllers to avoid pending Pods
* Arbitration policy for storage, e.g. "PV/PVC"
* HA of Arbitrator
* Hierarchical Consumer
* Infinite `Allocation.deserved` for `kube-system` 
* Arbitration policy for limits (Node/Resource QoS)
* `Consumer`/`Allocation` client for workload-specified controller

## Reference:

1. [Kubernetes] ResourceQuota: http://kubernetes.io/docs/admin/resourcequota/
1. [Kubernetes] Admission Control: http://kubernetes.io/docs/admin/admission-controllers/
1. [Kubernetes] Preemption and Re-scheduler: https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/rescheduling.md
1. [Kubernetes] Node-level QoS: http://kubernetes.io/docs/user-guide/compute-resources/
1. Kubernetes on EGO: http://sched.co/8K3n
1. Kubernetes on Mesos: https://github.com/kubernetes-incubator/kube-mesos-framework/
1. IBM Spectrum Conductor for Container: http://ibm.biz/ConductorForContainers
1. Support Spark natively in Kubernetes: https://github.com/kubernetes/kubernetes/issues/34377
1. Multi-Scheduler in Kubernetes: https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/multiple-schedulers.md
