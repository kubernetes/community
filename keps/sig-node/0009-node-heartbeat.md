---
kep-number: 8
title: Efficient Node Heartbeat
authors:
  - "@wojtek-t"
  - "with input from @bgrant0607, @dchen1107, @yujuhong, @lavalamp"
owning-sig: sig-node
participating-sigs:
  - sig-scalability
  - sig-apimachinery
  - sig-scheduling
reviewers:
  - "@deads2k"
  - "@lavalamp"
approvers:
  - "@dchen1107"
  - "@derekwaynecarr"
editor: TBD
creation-date: 2018-04-27
last-updated: 2018-04-27
status: implementable
see-also:
  - https://github.com/kubernetes/kubernetes/issues/14733
  - https://github.com/kubernetes/kubernetes/pull/14735
replaces:
  - n/a
superseded-by:
  - n/a
---

# Efficient Node Heartbeats

## Table of Contents

Table of Contents
=================

* [Efficient Node Heartbeats](#efficient-node-heartbeats)
   * [Table of Contents](#table-of-contents)
   * [Summary](#summary)
   * [Motivation](#motivation)
      * [Goals](#goals)
      * [Non-Goals](#non-goals)
   * [Proposal](#proposal)
      * [Risks and Mitigations](#risks-and-mitigations)
   * [Graduation Criteria](#graduation-criteria)
   * [Implementation History](#implementation-history)
   * [Alternatives](#alternatives)
      * [Dedicated “heartbeat” object instead of “leader election” one](#dedicated-heartbeat-object-instead-of-leader-election-one)
      * [Events instead of dedicated heartbeat object](#events-instead-of-dedicated-heartbeat-object)
      * [Reuse the Component Registration mechanisms](#reuse-the-component-registration-mechanisms)
      * [Split Node object into two parts at etcd level](#split-node-object-into-two-parts-at-etcd-level)
      * [Delta compression in etcd](#delta-compression-in-etcd)
      * [Replace etcd with other database](#replace-etcd-with-other-database)

## Summary

Node heartbeats are necessary for correct functioning of Kubernetes cluster.
This proposal makes them significantly cheaper from both scalability and
performance perspective.

## Motivation

While running different scalability tests we observed that in big enough clusters
(more than 2000 nodes) with non-trivial number of images used by pods on all
nodes (10-15), we were hitting etcd limits for its database size. That effectively
means that etcd enters "alert mode" and stops accepting all write requests.

The underlying root cause is combination of:

- etcd keeping both current state and transaction log with copy-on-write
- node heartbeats being pontetially very large objects (note that images
  are only one potential problem, the second are volumes and customers
  want to mount 100+ volumes to a single node) - they may easily exceed 15kB;
  even though the patch send over network is small, in etcd we store the
	whole Node object
- Kubelet sending heartbeats every 10s

This proposal presents a proper solution for that problem.


Note that currently (by default):

- Lack of NodeStatus update for `<node-monitor-grace-period>` (default: 40s)
  results in NodeController marking node as NotReady (pods are no longer
  scheduled on that node)
- Lack of NodeStatus updates for `<pod-eviction-timeout>` (default: 5m)
  results in NodeController starting pod evictions from that node

We would like to preserve that behavior.


### Goals

- Reduce size of etcd by making node heartbeats cheaper

### Non-Goals

The following are nice-to-haves, but not primary goals:

- Reduce resource usage (cpu/memory) of control plane (e.g. due to processing
  less and/or smaller objects)
- Reduce watch-related load on Node objects

## Proposal

We propose introducing a new `Lease` built-in API in the newly create API group
`coordination.k8s.io`. To make it easily reusable for other purposes it will
be namespaced. Its schema will be as following:

```
type Lease struct {
  metav1.TypeMeta `json:",inline"`
  // Standard object's metadata.
  // More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
  // +optional
  ObjectMeta metav1.ObjectMeta `json:"metadata,omitempty"`

  // Specification of the Lease.
  // More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
  // +optional
  Spec LeaseSpec `json:"spec,omitempty"`
}

type LeaseSpec struct {
  HolderIdentity       string           `json:"holderIdentity"`
  LeaseDurationSeconds int32            `json:"leaseDurationSeconds"`
  AcquireTime          metav1.MicroTime `json:"acquireTime"`
  RenewTime            metav1.MicroTime `json:"renewTime"`
  LeaseTransitions     int32            `json:"leaseTransitions"`
}
```

The Spec is effectively of already existing (and thus proved) [LeaderElectionRecord][].
The only difference is using `MicroTime` instead of `Time` for better precision.
That would hopefully allow us go get directly to Beta.

We will use that object to represent node heartbeat - for each Node there will
be a corresponding `Lease` object with Name equal to Node name in a newly
created dedicated namespace (we considered using `kube-system` namespace but
decided that it's already too overloaded).
That namespace should be created automatically (similarly to "default" and
"kube-system", probably by NodeController) and never be deleted (so that nodes
don't require permission for it).

We considered using CRD instead of built-in API. However, even though CRDs are
`the new way` for creating new APIs, they don't yet have versioning support
and are significantly less performant (due to lack of protobuf support yet).
We also don't know whether we could seamlessly transition storage from a CRD
to a built-in API if we ran into a performance or any other problems.
As a result, we decided to proceed with built-in API.


With this new API in place, we will change Kubelet so that:

1. Kubelet is periodically computing NodeStatus every 10s (at it is now), but that will
   be independent from reporting status
1. Kubelet is reporting NodeStatus if:
   - there was a meaningful change in it (initially we can probably assume that every
     change is meaningful, including e.g. images on the node)
   - or it didn’t report it over last `node-status-update-period` seconds
1. Kubelet creates and periodically updates its own Lease object and frequency
   of those updates is independent from NodeStatus update frequency.

In the meantime, we will change `NodeController` to treat both updates of NodeStatus
object as well as updates of the new `Lease` object corresponding to a given
node as healthiness signal from a given Kubelet. This will make it work for both old
and new Kubelets.

We should also:

1. audit all other existing core controllers to verify if they also don’t require
   similar changes in their logic ([ttl controller][] being one of the examples)
1. change controller manager to auto-register that `Lease` CRD
1. ensure that `Lease` resource is deleted when corresponding node is
   deleted (probably via owner references)
1. [out-of-scope] migrate all LeaderElection code to use that CRD

Once all the code changes are done, we will:

1. start updating `Lease` object every 10s by default, at the same time
   reducing frequency of NodeStatus updates initially to 40s by default.
   We will reduce it further later.
   Note that it doesn't reduce frequency by which Kubelet sends "meaningful"
   changes - it only impacts the frequency of "lastHeartbeatTime" changes.
   <br> TODO: That still results in higher average QPS. It should be acceptable but
   needs to be verified.
1. announce that we are going to reduce frequency of NodeStatus updates further
   and give people 1-2 releases to switch their code to use `Lease`
   object (if they relied on frequent NodeStatus changes)
1. further reduce NodeStatus updates frequency to not less often than once per
   1 minute.
   We can’t stop periodically updating NodeStatus as it would be API breaking change,
   but it’s fine to reduce its frequency (though we should continue writing it at
   least once per eviction period).


To be considered:

1. We may consider reducing frequency of NodeStatus updates to once every 5 minutes
   (instead of 1 minute). That would help with performance/scalability even more.
   Caveats:
   - NodeProblemDetector is currently updating (some) node conditions every 1 minute
     (unconditionally, because lastHeartbeatTime always changes). To make reduction
     of NodeStatus updates frequency really useful, we should also change NPD to
     work in a similar mode (check periodically if condition changes, but report only
     when something changed or no status was reported for a given time) and decrease
     its reporting frequency too.
   - In general, we recommend to keep frequencies of NodeStatus reporting in both
     Kubelet and NodeProblemDetector in sync (once all changes will be done) and
     that should be reflected in [NPD documentation][].
   - Note that reducing frequency to 1 minute already gives us almost 6x improvement.
     It seems more than enough for any foreseeable future assuming we won’t
     significantly increase the size of object Node.
     Note that if we keep adding node conditions owned by other components, the
     number of writes of Node object will go up. But that issue is separate from
     that proposal.

Other notes:

1. Additional advantage of using Lease for that purpose would be the
   ability to exclude it from audit profile and thus reduce the audit logs footprint.

[LeaderElectionRecord]: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/tools/leaderelection/resourcelock/interface.go#L37
[ttl controller]: https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/ttl/ttl_controller.go#L155
[NPD documentation]: https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
[kubernetes/kubernetes#63667]: https://github.com/kubernetes/kubernetes/issues/63677

### Risks and Mitigations

Increasing default frequency of NodeStatus updates may potentially break clients
relying on frequent Node object updates. However, in non-managed solutions, customers
will still be able to restore previous behavior by setting appropriate flag values.
Thus, changing defaults to what we recommend is the path to go with.

## Graduation Criteria

The API can be immediately promoted to Beta, as the API is effectively a copy of
already existing LeaderElectionRecord. It will be promoted to GA once it's gone
a sufficient amount of time as Beta with no changes.

The changes in components logic (Kubelet, NodeController) should be done behind
a feature gate. We suggest making that enabled by default once the feature is
implemented.

## Implementation History

- RRRR-MM-DD: KEP Summary, Motivation and Proposal merged

## Alternatives

We considered a number of alternatives, most important mentioned below.

### Dedicated “heartbeat” object instead of “leader election” one

Instead of introducing and using “lease” object, we considered
introducing a dedicated “heartbeat” object for that purpose. Apart from that,
all the details about the solution remain pretty much the same.

Pros:

- Conceptually easier to understand what the object is for

Cons:

- Introduces a new, narrow-purpose API. Lease is already used by other
  components, implemented using annotations on Endpoints and ConfigMaps.

### Events instead of dedicated heartbeat object

Instead of introducing a dedicated object, we considered using “Event” object
for that purpose. At the high-level the solution looks very similar. 
The differences from the initial proposal are:

- we use existing “Event” api instead of introducing a new API
- we create a dedicated namespace; events that should be treated as healthiness
  signal by NodeController will be written by Kubelets (unconditionally) to that
  namespace
- NodeController will be watching only Events from that namespace to avoid
  processing all events in the system (the volume of all events will be huge)
- dedicated namespace also helps with security - we can give access to write to
  that namespace only to Kubelets

Pros:

- No need to introduce new API
   - We can use that approach much earlier due to that.
- We already need to optimize event throughput - separate etcd instance we have
  for them may help with tuning
- Low-risk roll-forward/roll-back: no new objects is involved (node controller
  starts watching events, kubelet just reduces the frequency of heartbeats)

Cons:

- Events are conceptually “best-effort” in the system:
   - they may be silently dropped in case of problems in the system (the event recorder
     library doesn’t retry on errors, e.g. to not make things worse when control-plane
     is starved)
   - currently, components reporting events don’t even know if it succeeded or not (the
     library is built in a way that you throw the event into it and are not notified if
     that was successfully submitted or not).
     Kubelet sending any other update has full control on how/if retry errors.
   - lack of fairness mechanisms means that even when some events are being successfully
     send, there is no guarantee that any event from  a given Kubelet will be submitted
     over a given time period
	So this would require a different mechanism of reporting those “heartbeat” events.
- Once we have “request priority” concept, I think events should have the lowest one.
  Even though no particular heartbeat is important, guarantee that some heartbeats will
  be successfully send it crucial (not delivering any of them will result in unnecessary
  evictions or not-scheduling to a given node). So heartbeats should be of the highest
  priority. OTOH, node heartbeats are one of the most important things in the system
  (not delivering them may result in unnecessary evictions), so they should have the
  highest priority.
- No core component in the system is currently watching events
   - it would make system’s operation harder to explain
- Users watch Node objects for heartbeats (even though we didn’t recommend it).
  Introducing a new object for the purpose of heartbeat will allow those users to
  migrate, while using events for that purpose breaks that ability. (Watching events
  may put us in tough situation also from performance reasons.)
- Deleting all events (e.g. event etcd failure + playbook response) should continue to
  not cause a catastrophic failure and the design will need to account for this.

### Reuse the Component Registration mechanisms

Kubelet is one of control-place components (shared controller). Some time ago, Component
Registration proposal converged into three parts:

- Introducing an API for registering non-pod endpoints, including readiness information: #18610
- Changing endpoints controller to also watch those endpoints
- Identifying some of those endpoints as “components”

We could reuse that mechanism to represent Kubelets as non-pod endpoint API.

Pros:

- Utilizes desired API

Cons:

- Requires introducing that new API
- Stabilizing the API would take some time
- Implementing that API requires multiple changes in different components

### Split Node object into two parts at etcd level

We may stick to existing Node API and solve the problem at storage layer. At the
high level, this means splitting the Node object into two parts in etcd (frequently
modified one and the rest).

Pros:

- No need to introduce new API
- No need to change any components other than kube-apiserver

Cons:

- Very complicated to support watch
- Not very generic (e.g. splitting Spec and Status doesn’t help, it needs to be just
  heartbeat part)
- [minor] Doesn’t reduce amount of data that should be processed in the system (writes,
  reads, watches, …)

### Delta compression in etcd

An alternative for the above can be solving this completely at the etcd layer. To
achieve that, instead of storing full updates in etcd transaction log, we will just
store “deltas” and snapshot the whole object only every X seconds/minutes.

Pros:

- Doesn’t require any changes to any Kubernetes components

Cons:

- Computing delta is tricky (etcd doesn’t understand Kubernetes data model, and
  delta between two protobuf-encoded objects is not necessary small)
- May require a major rewrite of etcd code and not even be accepted by its maintainers
- More expensive computationally to get an object in a given resource version (which
  is what e.g. watch is doing)

### Replace etcd with other database

Instead of using etcd, we may also consider using some other open-source solution.

Pros:

- Doesn’t require new API

Cons:

- We don’t even know if there exists solution that solves our problems and can be used.
- Migration will take us years.
