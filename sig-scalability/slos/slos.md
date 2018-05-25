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

```
Prerequisite: Kubernetes cluster is available and serving.
```

### Steady state SLIs/SLOs

| Status | SLI | SLO | User stories, test scenarios, ... |
| --- | --- | --- | --- |
| __Official__ | Latency<sup>[1](#footnote1)</sup> of mutating<sup>[2](#footnote2)</sup> API calls for single objects for every (resource, verb) pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, verb) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day<sup>[3](#footnote3)</sup> <= 1s | [Details](./api_call_latency.md) |
| __Official__ | Latency<sup>[1](#footnote1)</sup> of non-streaming read-only<sup>[4](#footnote3)</sup> API calls for every (resource, scope<sup>[5](#footnote4)</sup>) pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, scope) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day (a) <= 1s if `scope=resource` (b) <= 5s if `scope=namespace` (c) <= 30s if `scope=cluster` | [Details](./api_call_latency.md) |

<a name="footnote1">\[1\]</a>By latency of API call in this doc we mean time
from the moment when apiserver gets the request to last byte of response sent
to the user.

<a name="footnote2">\[2\]</a>By mutating API calls we mean POST, PUT, DELETE
and PATCH.

<a name="footnote3">\[3\]</a> For the purpose of visualization it will be a
sliding window. However, for the purpose of reporting the SLO, it means one
point per day (whether SLO was satisfied on a given day or not).

<a name="footnote4">\[4\]</a>By non-streaming read-only API calls we mean GET
requests without `watch=true` option set. (Note that in Kubernetes internally
it translates to both GET and LIST calls).

<a name="footnote5">\[5\]</a>A scope of a request can be either (a) `resource`
if the request is about a single object, (b) `namespace` if it is about objects
from a single namespace or (c) `cluster` if it spawns objects from multiple
namespaces.


__TODO: Migrate existing SLIs/SLOs here:__
- __Pod startup time__

### Burst SLIs/SLOs

| Status | SLI | SLO | User stories, test scenarios, ... |
| --- | --- | --- | --- |
| WIP | Time to start 30\*#nodes pods, measured from test scenario start until observing last Pod as ready | Benchmark: when all images present on all Nodes, 99th percentile <= X minutes | [Details](./system_throughput.md) |

### Other SLIs

| Status | SLI | User stories, ... |
| --- | --- | --- |
| WIP | Watch latency for every resource, (from the moment when object is stored in database to when it's ready to be sent to all watchers), measured as 99th percentile over last 5 minutes | TODO |
| WIP | Admission latency for each admission plugin type, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |
| WIP | Webhook call latency for each webhook type, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |
| WIP | Initializer latency for each initializer, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |

