# Priority in Kubernetes API

@bsalamat

May 2017
  * [Objective](#objective)
    * [Non-Goals](#non-goals)
  * [Background](#background)
  * [Overview](#overview)
  * [Detailed Design](#detailed-design)
    * [Effect of priority on scheduling](#effect-of-priority-on-scheduling)
    * [Effect of priority on preemption](#effect-of-priority-on-preemption)
    * [Priority in PodSpec](#priority-in-podspec)
    * [Priority Classes](#priority-classes)
    * [Resolving priority class names](#resolving-priority-class-names)
    * [Ordering of priorities](#ordering-of-priorities)
    * [System Priority Class Names](#system-priority-class-names)
    * [Modifying Priority Classes](#modifying-priority-classes)
    * [Drawbacks of changing priority names](#drawbacks-of-changing-priority-classes)
    * [Priority and QoS classes](#priority-and-qos-classes)


## Objective



*   How to specify priority for workloads in Kubernetes API.
*   Define how the order of these priorities are specified. 
*   Define how new priority levels are added.
*   Effect of priority on scheduling and preemption.

### Non-Goals 



*   How preemption works in Kubernetes.
*   How quota allocation and accounting works for each priority.

## Background 

It is fairly common in clusters to have more tasks than what the cluster
resources can handle. Often times the workload is a mix of high priority
critical tasks, and non-urgent tasks that can wait. Cluster management should be
able to distinguish these workloads in order to decide which ones should acquire
the resources sooner and which ones can wait. Priority of the workload is one of
the key metrics that provides the information to the cluster. This document is a
more detailed design proposal for part of the high-level architecture described
in [Resource sharing architecture for batch and serving workloads in Kubernetes](https://docs.google.com/document/d/1-H2hnZap7gQivcSU-9j4ZrJ8wE_WwcfOkTeAGjzUyLA).

## Overview 

This design doc introduces the concept of priorities for pods in Kubernetes and
how the priority impacts scheduling and preemption of pods when the cluster
runs out of resources. A pod can specify a priority at the creation time. The
priority must be one of the valid values and there is a total order on the
values. The priority of a pod is independent of its workload type. The priority
is global and not specific to a particular namespace.

## Detailed Design 

### Effect of priority on scheduling 

One could generally expect a pod with higher priority has a higher chance of
getting scheduled than the same pod with lower priority. However, there are
many other parameters that affect scheduling decisions. So, a high priority pod
may or may not be scheduled before lower priority pods. The details of
what determines the order at which pods are scheduled are beyond the scope of
this document.

### Effect of priority on preemption 

Generally, lower priority pods are more likely to get preempted by higher
priority pods when cluster has reached a threshold. In such a case, scheduler
may decide to preempt lower priority pods to release enough resources for higher
priority pending pods. As mentioned before, there are many other parameters
that affect scheduling decisions, such as affinity and anti-affinity. If
scheduler determines that a high priority pod cannot be scheduled even if lower
priority pods are preempted, it will not preempt lower priority pods. Scheduler
may have other restrictions on preempting pods, for example, it may refuse to
preempt a pod if PodDisruptionBudget is violated. The details of scheduling and
preemption decisions are beyond the scope of this document.

### Priority in PodSpec 

Pods may have priority in their pod spec. PodSpec will have two new fields
called "PriorityClassName" which is specified by user, and "Priority" which will
be populated by Kubernetes. User-specified priority (PriorityClassName) is a 
string and all of the valid priority classes are defined by a system wide
mapping that maps each string to an integer. The PriorityClassName specified in
a pod spec must be found in this map or the pod creation request will be
rejected. If PriorityClassName is empty, it will resolve to the default
priority (See below for more info on name resolution). Once the 
PriorityClassName is resolved to an integer, it is placed in "Priority" field of
PodSpec.


```
type PodSpec struct {
  ...
  PriorityClassName string
  Priority          *int32  // Populated by Admission Controller. Users are not allowed to set it directly.
}
```

### Priority Classes 

The cluster may have many user defined priority classes for
various use cases. The following list is an example of how the priorities and
their values may look like.
Kubernetes will also have special priority class names reserved for critical system
pods. Please see [System Priority Class Names](#system-priority-class-names) for
more information. Any priority value above 1 billion is reserved for system use.
Aside from those system priority classes, Kubernetes is not shipped with predefined
priority classes usable by user pods. The main goal of having no built-in
priority classes for user pods is to avoid creating defacto standard names which
may be hard to change in the future.

```
system  2147483647 (int_max)
tier1   4000
tier2   2000
tier3   1000
```

The following shows a list of example workloads in a Kubernetes cluster in decreasing order of priority:

* Kubernetes system daemons (per-node like fluentd, and cluster-level like 
  Heapster)
* Critical user infrastructure (e.g. storage servers, monitoring system like
  Prometheus, etc.)
* Components that are in the user-facing request serving path and must be able 
  to scale up arbitrarily in response to load spikes (web servers, middleware,
  etc.)
* Important interruptible workloads that need strong guarantee of
  schedulability and of not being interrupted
* Less important interruptible workloads that need a less strong guarantee of
  schedulability and of not being interrupted
* Best effort / opportunistic

### Resolving priority class names 

User requests sent to Kubernetes may have `PriorityClassName` in their PodSpec.
Admission controller resolves a PriorityClassName to its corresponding number
and populates the "Priority" field of the pod spec. The rest of Kubernetes
components look at the "Priority" field of pod status and work with the integer
value. In other words, `PriorityClassName` will be ignored by the rest of the
system.

We are going to add a new API object called PriorityClass. The priority class
defines the mapping between the priority name and its value. It can have an
optional description. It is an arbitrary string and is provided
only as a guideline for users.

A priority class can be marked as "Global Default" by setting its
`GlobalDefault` field to true. If a pod does not specify any `PriorityClassName`,
the system resolves it to the value of the global default priority class if
exists. If there is no global default, the pod's priority will be resolved to
zero. Priority admission controller ensures that there is only one global
default priority class.

```
type PriorityClass struct {
  metav1.TypeMeta
  // +optional
  metav1.ObjectMeta
  
  // The value of this priority class. This is the actual priority that pods
  // receive when they have the above name in their pod spec.
  Value        int32
  GlobalDefault     bool
  Description       string
}
```

### Ordering of priorities 

As mentioned earlier, a PriorityClassName is resolved by the admission controller to
its integral value and Kubernetes components use the integral value. The higher
the value, the higher the priority.

### System Priority Class Names
There will be special priority class names reserved for system use only. These
classes have a value larger than one billion. 
Priority admission controller ensures that new priority classes will be not
created with those names. They are used for critical system pods that must not
be preempted. We set default policies that deny creation of pods with
PriorityClassNames corresponding to these priorities. Cluster admins can 
authorize users or service accounts to create pods with these priorities. When
non-authorized users set PriorityClassName to one of these priority classes in
their pod spec, their pod creation request will be rejected. For pods created by
controllers, the service account must be authorized by cluster admins.

### Modifying priority classes 

Priority classes can be added or removed, but their name and value cannot be
updated. We allow updating `GlobalDefault` and `Description` as long as there is
a maximum of one global default. While
Kubernetes can work fine if priority classes are changed at run-time, the change
can be confusing to users as pods with a priority class which were created
before the change will have a different priority value than those created after
the change. Deletion of priority classes is allowed, despite the fact that there
may be existing pods that have specified such priority class names in their pod
spec. In other words, there will be no referential integrity for priority
classes. This is another reason that all system components should only work with
the integer value of the priority and not with the `PriorityClassName`.

One could delete an existing priority class and create another one with the same
name and a different value. By doing so, they can achieve the same effect as
updating a priority class, but we still do not allow updating priority classes
to prevent accidental changes.

Newly added priority classes cannot have a value higher than what is reserved
for "system". The reason for this restriction
is that Kubernetes critical system pods will have one of the "system" priorities
and no pod should be able to preempt them.

#### Drawbacks of changing priority classes 

While Kubernetes effectively allows changing priority classes (by deleting and
adding them with a different value), it should be done only when
absolutely needed. Changing priority classes has the following disadvantages:


*   May remove config portability: pod specs written for one cluster are no
    longer guaranteed to work on a different cluster if the same priority classes
    do not exist in the second cluster. 
*   If quota is specified for existing priority classes (at the time of this writing,
    we don't have this feature in Kubernetes), adding or deleting priority classes
    will require reconfiguration of quota allocations.
*   An existing pods may have an integer value of priority that does not reflect
    the current value of its PriorityClass.

### Priority and QoS classes 

Kubernetes has [three QoS
classes](/contributors/design-proposals/node/resource-qos.md#qos-classes)
which are derived from request and limit of pods. Priority is introduced as an
independent concept; meaning that any QoS class may have any valid priority.
When a node is out of resources and pods needs to be preempted, we give
priority a higher weight over QoS classes. In other words, we preempt the lowest
priority pod and break ties with some other metrics, such as, QoS class, usage
above request, etc. This is not finalized yet. We will discuss and finalize
preemption in a separate doc.
