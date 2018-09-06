---
kep-number: 28
title: Node Health Extensibility
authors:
  - "@vishh"
owning-sig: sig-node
participating-sigs:
  - sig-architecture
  - sig-scheduling
  - sig-autoscaling
reviewers:
  - derekwaynecarr
  - dawnchen
  - thockin
approvers:
  - TBD
editor: TBD
creation-date: 2018-08-27
status: provisional
---

# Node Health Extensibility

## Summary

As of Kubernetes v1.12, the health of extended node components like storage, network, device plugins or logging, monitoring agents aren't taken into account while managing user pods on kubernetes nodes. This leads to a broken contract between cluster administrators and end users where the latter ends up debugging issues like lost metrics or logs, lack of network or storage connectivity, etc. On top of that lack of a generic abstraction for node readiness leads to hacks across the k8s ecosystem to deal with the asynchronous & dynamic nature of bootsrapping and managing k8s nodes. 
This proposal presents several approaches for solving the problem of extended node health or Node Readiness++.

## Problem Statement

Since the early days of kubernetes, kubelet has been the primary source of readiness of a node where it queried the health of critical system components like the docker daemon, network components, etc.
Kubelet owns the node “Ready” condition and the rest of kubernetes cluster components rely on this (and a few other conditions) to infer that a node is “online”. 

Kubernetes has evolved quite a bit since the early days and now a kubernetes “node” is comprised of many system microservices that are all essential for a node to go “online” and start running user applications.
Examples include device & storage plugins (network plugins in the future?), logging & monitoring agents, networking daemons like kube-proxy, machine management daemons like Node Problem Detector, etc.
Over the years, kubelet and the container runtimes have evolved adequately to become highly reliable and so they become “ready” more consistently.
But the other newer system extensions may not be as mature yet, or may be slow to transition to “ready” for legitimate reasons.

The lack of a single unambiguous signal on when a “node” is “ready”, beyond just the kubelet or a container runtime is causing several issues today.

- System components like Cluster Autoscaler and Scheduler have to combine several Node Conditions to evaluate a node’s actual readiness. It is not straightforward to introduce new conditions into the policy.
- Cluster Autoscaler incorrectly assumes a GPU node to be “ready” prior to device plugins and/or device drivers being installed which leads to spurious node additions and deletions thereby leading to custom logic in the autoscaler.
- Logs & metrics are lost if logging & monitoring agents are not online which can lead to sub-optimal (often unexpected) user behavior.
- A node with offline kube-proxy can result in service outage in the event of an application update (via deployments for example).
- It is difficult for Node Controller and/or node repair systems to detect and remedy issues with extended system components of the node since they lack a set of generic & easy to identify signals from those individual components on the node.

As Kubernetes becomes more and more customizable (or extensible) and gets deployed in mission critical systems, it is imperative to provide fail safe mechanisms that are built into the system.

The remainder of this proposal explores a few approaches for addressing the issue of node health extensibility.

## Solutions

Any solution that is chosen for addressing the set of problems stated above should meet the following design requirements:

- Be backwards compatible
- Be extensible to accommodate variations in kubernetes environments - the set of system microservices can vary across k8s deployments.
- Ideally, provide a single signal (not a policy) to cluster level components. 
- Avoid introducing new API primitives (ideally).
- Support self-bootstrapping - Scheduling system components (including the kubelet) even if the node is not logically ready.

### Approach A - Track state of extensions via APIs (Preferred)

For reasons similar to recent [node heartbeats revamp](https://github.com/kubernetes/community/blob/6a2652fdf1f7d662a70aa3202147b85ff2d6408e/keps/sig-node/0009-node-heartbeat.md), individual node extensions will be expected to post their state via dedicated "[Lease](https://github.com/kubernetes/api/blob/master/coordination/v1beta1/types.go#L27)" API objects.

The Node API will include a new field that will track a list of node extensions that are expected on the node.
This list can be initially populated by the node controller when a node object is created or by the kubelet as it’s extensions register, or by individual extensions themselves (if they have sufficient privileges). 

Here are some sample illustrations:

```yaml
type Node struct {
 …
 Spec:
  SystemExtensions []SystemExtension
...
}

type SystemExtension struct {
   Reference ObjectReference
   AffectsReadiness bool // Ignore non-critical node extensions from node readiness. Helps identify the overall set of extensions deterministically.
}
```

The following journey attempts to illustrate this solution a bit more in detail:

1. Kubelet starts up on a node and immediately taints the node with a special “kubernetes.io/MetaReadiness” taint with an effect of “NoSchedule” in addition to flipping it’s “Ready” condition to true (when appropriate). 
2. Kubelet will not re-admit existing pods on the node that do not tolerate the “Readiness” taint. Kubelet will also not evict yet to be admitted pods to avoid causing disruptions. The node is still not fully functional yet.
3. The existing conditions that kubelet manages will continue to behave as-is (until they are transitioned to taints eventually).
4. Every node level extension component is expected to post it’s health continually via a dedicated Lease API that is understood by the the node controller.
5. The node controller will notice that the kubelet has restarted via the presence of the special taint. It will then wait for all expected node extension components to update their status prior to removing the taint.
6. Existing or newly deployed kubelet first class extensions (CSI or device plugins) on the node will register themselves with the kubelet and then update their state via a dedicated Lease object in a special (TBD) namespace.
7. Kubelet if newly registering will update the Node Object to include the device plugin as a required extension for determining node meta readiness (updates Node.Spec.SystemExtensions field as illustrated earlier) or the node controller can be updated to inject standard node extensions as part of Node object registration (possibly using webhooks).
8. Other non-kubelet extensions like kube-proxy, logging and monitoring agents will also register themselves (or get registered during Node object creation) as required addons by updating the Node object and periodically update their state.
9. For each node, the Node controller will keep track of the Lease objects that are included in Node.Spec.SystemExtensions field.
8. Once the node controller notices that all the required extensions have updated their health after the kubelet had placed a taint on the node, it will remove the taint. The node controller will continue to take into account the existing conditions as part of this design.
9. Once the taint is removed, kubelet will re-admit any existing pods and then transition to evicting any inadmissible pods.
10. If any of the system extensions fail to renew their lease (post health updates) for a TBD period (10s) of time, the node controller will taint the node with the same taint that the kubelet used originally. This prevents additional pods from being scheduled on to that node. If the node does not transition out of that state for a TBD duration (5 min?), the pods on the node will be evicted by the node controller.
11. Cluster Autoscaler (CA) will consider a new node to be logically booting until the special taint is removed.

Individual kubernetes vendors can easily customize their deployments and still use upstream Cluster Autoscaler and other ecosystem components that care about node health with this design.


#### Optional Extension - Track state of components along with Health

Each of the individual node extension will most likely want to expose domain specific status for instrospection purposes.
This status can be expressed as plugin specific CRDs which may be a good starting point.
Down the line though, having a generic abstraction(s) to represent the Status of different plugins or extensions will be necessary to empower the wider introspection & remediation ecosystem.

Assuming it is possible to define Status for some of the popular plugins like CSI, device plugins, kube proxy, logging & monitoring agents, the Lease API based approach can be extended as follows to easily track status of extensions in addition to Health.

1. Embed the "Lease" API object in an extension specific Status object. An alternative is to have a local Lease object with the same name as the Status object.
2. Similar to the workflow described above, the node controller, kubelet or extensions will register their status object as an extension via the Node.Spec.SystemExtensions field.
3. The Node Controller will be updated to extract the Lease object (or identifying a separate Lease object by name in the local namespace) and influence health of extensions similar to the workflow described above.
4. Ecosystem components like Monitoring, repair/remediation systems will be able to deterministically identify and expose or act on the status of the extensions.

Here are some API illustations.

##### Node Object extension

```yaml
type ExtensionFoo struct {
 TypeMeta
 ObjectMeta
 Spec ExtensionFooStatus
}

type ExtensionFooStatus {
  v1beta1.LeaseSpec
  OtherFooStatus interface{} 
}

type Node struct {
 …
 Spec:
  SystemExtensions []SystemExtension
...
}

type SystemExtension struct {
   Reference ObjectReference
   AffectsReadiness bool // Ignore non-critical node extensions from node readiness
}
```

##### Sample Extension Status Object

```yaml
type DevicePluginExtension struct {
 TypeMeta
 ObjectMeta
 Status DevicePluginExtensionStatus
}

type DevicePluginExtensionStatus struct {
  LeaseSpec
  ComputeResource struct{ // embedded for brevity
    Name: “foo.com/bar”
    Quantity: x // Sample fields for illustration purposes
    UnHealthy: y 
    DeviceIDs []string
  }
}

```

One of the goals for this proposal is to evaluate the viability and usefullness of combining Status and Health for extensions.

### Approach B - Use Readiness of Node extension pods

This approach assumes that all node level extensions will get deployed as kubernetes pods and that the health of those extension pods can be exposed via existing Readiness probes.

Similar to the previous approach `A`, a special taint will be used to track the overall readiness of the node and the node controller and kubelet will manage the special taint.

The Node Object will be extended to include a list of system extensions that are identified via a label selector and a namespace.

```yaml
type Node struct {
 …
 Spec:
  SystemExtensions []SystemExtension
...
}

type SystemExtension struct {
   Selector LabelSelector
   Namespace Namespace
   AffectsReadiness bool // Ignore non-critical node extensions from node readiness
}
```

Node selector will take into account `Readiness` all the pods that match the selector(s) specified in the Node Spec `SystemExtensions` field to determine Node Readiness.

The individual system extensions are required to detect kubelet crash (or restart) quickly and reflect that in their Readiness.
If the Node Controller observes that all the system extension pods are Ready after the kubelet places the special meta readiness taint, it will remove that taint immediately.
If any of the pods become unready, then the node controller will place the meta readiness taint on the nodes.

This approach has the advantage of not introducing additional API concepts to track state as illustrated in Approach A.
Side-car containers can be used to expose Readiness more easily than in approach A where it involves updating API objects directly.

Handling of kubelet restarts can be tricky though since there is a chance that the extension pods may not reflect their connections with the kubelet soon after the kubelet restarts.

### Approach C - Use conditions

This design continues the existing pattern of expressing health of various node features and components via Conditions.
For example the Kubelet uses conditions heavily and the Node Problem Detector also uses conditions. Conditions allow for recording heartbeats already.

This design is similar to the approach `A` where instead of recording state to a separate object, kubelet and other system components will record their state via Conditions.
The node controller will be extended to support a configurable policy on what Conditions influence the meta readiness of a node.

The kubelet and node controller will use a taint similar to the one illustrated in approach `A` to signal node meta readiness.

The scheduler needs to understand that certain special pods (system pods) can be scheduled to nodes even if the node isn’t logically ready.
There is no canonical construct today to differentiate system pods from regular pods. 

The main drawback of this approach is the explosion on heartbeats (updates to Condition objects) to the Node object from various components which was the reason why a new heartbeat mechanism was introduced for the kubelet in the first place.
Another drawback is that Conditions are not extensible - it is not possible to include extension specific information as part of conditions leading to leaking the state of extensions to other sub-objects in the Node API object eventually.

### Approach D - Support extensible health checks in the Kubelet

This approach assumes kubelet to be the central arbitrator of Node health.

Every node extension needs to register a canonical Readiness endpoint (similar to pod readiness/handlers) with the kubelet via the Node API object.

The kubelet will then poll health of individual extensions periodically and reflect the state of the individual extensions via a taint similar to the solutions mentioned above.
The main difference here is that the node controller doesn’t have to track the state of individual extensions since the kubelet is already doing it.

Optionally, the kubelet can reflect the state of each extension via a dedicated condition.

Upon restart, Kubelet will wait for one successful health check run across all registered extensions prior to removing the special taint.

This approach requires that all extensions are accessible to the kubelet via host network interfaces (barring the option of using exec binaries per addon that will suffer a distribution problem). 

This approach also lacks the ability to include more state information about each extension in the API object - aka it’s limited to only health tracking. The load on the kubelet could be much higher due to the proliferation of system extensions. Every extension is required to build a canonical healthz endpoint (which is a best practice by itself).




