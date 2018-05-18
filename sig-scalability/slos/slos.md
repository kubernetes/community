# Kubernetes scalability and performance SLIs/SLOs

## What Kubernetes guarantees?

One of the important aspects of Kubernetes is its scalability and performance
characteristic. As Kubernetes user or operator/administrator of a cluster
you would expect to have some guarantees in those areas.

The goal of this doc is to organize the guarantees that Kubernetes provides
in these areas.

## What do we require from SLIs/SLOs?

We are going to define more SLIs and SLOs based on the most important indicators
in the system.

Our SLOs need to have the following properties:
- <b> They need to be testable </b> <br/>
  That means that we need to have a benchmark to measure if it's met.
- <b> They need to be understandable for users </b> <br/>
  In particular, they need to be understandable for people not familiar
  with the system internals, i.e. their formulation can't depend on some
  arcane knowledge.

However, we may introduce some internal (for developers only) SLIs, that
may be useful for understanding performance characterstic of the system,
but for which we don't provide any guarantees for users and thus may not
be fully understandable for users.

On the other hand, we do NOT require that our SLOs:
- are measurable in a running cluster (though that's desired if possible) <br/>
  In other words, not SLOs need to be easily translatable to SLAs.
  Being able to benchmark is enough for us.

## Types of SLOs

While SLIs are very generic and don't really depend on anything (they just
define what and how we measure), it's not the case for SLOs.
SLOs provide guarantees, and satisfying them may depend on meeting some
specific requirements.

As a result, we build our SLOs in "you promise, we promise" format.
That means, that we provide you a guarantee only if you satisfy the requirement
that we put on you.

As a consequence we introduce the two types of SLOs.

### Steady state SLOs

With steady state SLOs, we provide guarantees about system's behavior during
normal operations. We are able to provide much more guarantees in that situation.

```Definition
We define system to be in steady state when the cluster churn per second is <= 20, where

churn = #(Pod spec creations/updates/deletions) + #(user originated requests) in a given second
```

### Burst SLO

With burst SLOs, we provide guarantees on how system behaves under the heavy load
(when user wants the system to do something as quickly as possible not caring too
much about response time).

## Environment

In order to meet the SLOs, system must run in the environment satisfying
the following criteria:
- Runs a single or more appropriate sized master machines
- Main etcd running on master machine(s)
- Events are stored in a separate etcd running on the master machine(s)
- Kubernetes version is at least X.Y.Z
- ...

__TODO: Document other necessary configuration.__

## Thresholds

To make the cluster eligible for SLO, users also can't have too many objects in
their clusters. More concretely, the number of different objects in the cluster
MUST satisfy thresholds defined in [thresholds file][].

[thresholds file]: https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md


## Kubernetes SLIs/SLOs

The currently existing SLIs/SLOs are enough to guarantee that cluster isn't
completely dead. However, the are not enough to satisfy user's needs in most
of the cases.

We are looking into extending the set of SLIs/SLOs to cover more parts of
Kubernetes.

### Steady state SLIs/SLOs

| Status | SLI | SLO | User stories, test scenarios, ... |
| --- | --- | --- | --- |

__TODO: Migrate existing SLIs/SLOs here:__
- __API-machinery ones__
- __Pod startup time__

### Burst SLIs/SLOs

| Status | SLI | SLO | User stories, test scenarios, ... |
| --- | --- | --- | --- |
| WIP | Time to start 30\*#nodes pods, measured from test scenario start until observing last Pod as ready | Benchmark: when all images present on all Nodes, 99th percentile <= X minutes | [Details](./system_throughput.md) |
