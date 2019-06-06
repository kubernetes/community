# Kubernetes scalability and performance SLIs/SLOs

## What Kubernetes guarantees?

One of the important aspects of Kubernetes is its scalability and performance
characteristic. As Kubernetes user or operator/administrator of a cluster
you would expect to have some guarantees in those areas.

The goal of this doc is to organize the guarantees that Kubernetes provides
in these areas.

## How we define scalability?

Our scalability definition is built on two concepts:
- [Service Level Indicators]
- [Service Level Objectives]

We require our SLIs/SLOs to have the following properties:
- <b> They are precise and well-defined </b> <br/>
  It's extremely important to ensure that both users and us have exactly the
  same understanding of what we guarantee.
- <b> They are consistent with each other </b> <br/>
  This is mostly about using the same terminology, same concepts, etc.
- <b> They are user-oriented </b> <br/>
  First, the SLOs we provide need to be things users really care about.
  Second, they need to be understandable for people not familiar with the system
  internals (e.g. their formulation can't depend on some arcane knowledge or
  implementation details of the system).
- <b> They are testable </b> <br/>
  Ideally, SLIs/SLOs should be measurable in all running clusters, but if measuring
  some metrics isn't possible or would be extremely expensive (e.g. in terms
  of resource overhead for the system), benchmarks sometimes may be enough.
  That means that not every SLO may be translatable to SLA ([Service Level
  Agreement]).

While SLIs are generic (they just define what and how we measure), SLOs provide
specific guarantees and satisfying them may depend on meeting some specific
requirements. Specific examples that may visibly affect ability to satisfy them
are:
- cluster configuration
- user of Kubernetes extensibility features
- load on the cluster.

As a result, we define Kubernetes scalability using "you promise, we promise"
framework, as following:

<b> If you promise to:
- correctly configure your cluster
- use extensibility features "reasonably"
- keep the load in the cluster within recommended limits

then we promise that your cluster scales, i.e.:
- all the SLOs are satisfied. </b>

We are in the process of extending coverage of the system with SLIs and SLOs
to better reflect user expectations.

Note that may also introduce internal (for developers only) SLIs, that may be
useful for understanding performance characteristic of the system, but for which
we will not provide any guarantees for users.

[Service Level Indicators]: https://en.wikipedia.org/wiki/Service_level_indicator
[Service Level Objectives]: https://en.wikipedia.org/wiki/Service_level_objective
[Service Level Agreement]: https://en.wikipedia.org/wiki/Service-level_agreement

### Environment (cluster configuration)

In order to meet SLOs, system must run in the environment satisfying
the following criteria:
- Runs a single or more appropriately sized master machines
- Events are stored in a separate etcd instance (or cluster)
- All etcd instances are running on master machine(s)
- Kubernetes version is at least X.Y.Z
- ...

__TODO: Document other necessary configuration.__

### Scalability thresholds

To make the cluster eligible for SLO, users also can't have too many objects in
their clusters. More concretely, the number of different objects in the cluster
MUST satisfy thresholds defined in [thresholds file][].

[thresholds file]: https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md

### Kubernetes extensibility

In order to meet SLOs, you have to use extensibility features "wisely".
The more precise formulation is to-be-defined, but this includes things like:
- webhooks have to provide high availability and low latency
- CRDs and CRs have to be kept within thresholds
- ...

## Kubernetes SLIs/SLOs

The currently existing SLIs/SLOs are enough to guarantee that cluster isn't
completely dead. However, they are not meeting user expectations in many areas of
the system and we are actively working on extending their coverage.

We are also introducing two more prerequisites which have to be met to ensure that
SLOs can be satisfied:

```
Prerequisites:
   1. Kubernetes cluster is available and serving.
   2. Cluster churn is <= 20, where churn is defined as:
     churn = #(Pod spec creations/updates/deletions) + #(user originated requests) in a given second
```

__TODO: Cluster churn should be moved to scalability thresholds.__


### Steady state SLIs/SLOs

| Status | SLI | SLO | User stories, test scenarios, ... |
| --- | --- | --- | --- |
| __Official__ | Latency of mutating API calls for single objects for every (resource, verb) pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, verb) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> <= 1s | [Details](./api_call_latency.md) |
| __Official__ | Latency of non-streaming read-only API calls for every (resource, scope pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, scope) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> (a) <= 1s if `scope=resource` (b) <= 5s if `scope=namespace` (c) <= 30s if `scope=cluster` | [Details](./api_call_latency.md) |
| __Official__ | Startup latency of schedulable stateless pods, excluding time to pull images and run init containers, measured from pod creation timestamp to when all its containers are reported as started and observed via watch, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> <= 5s | [Details](./pod_startup_latency.md) |
| __WIP__ | Startup latency of schedulable stateful pods, excluding time to pull images, run init containers, provision volumes (in delayed binding mode) and unmount/detach volumes (from previous pod if needed), measured from pod creation timestamp to when all its containers are reported as started and observed via watch, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> <= X where X depends on storage provider | [Details](./pod_startup_latency.md) |
| __WIP__ | Latency of programming in-cluster load balancing mechanism (e.g. iptables), measured from when service spec or list of its `Ready` pods change to when it is reflected in load balancing mechanism, measured as 99th percentile over last 5 minutes aggregated across all programmers | In default Kubernetes installation, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> <= X | [Details](./network_programming_latency.md) |
| __WIP__ | Latency of programming dns instance, measured from when service spec or list of its `Ready` pods change to when it is reflected in that dns instance, measured as 99th percentile over last 5 minutes aggregated across all dns instances | In default Kubernetes installation, 99th percentile per cluster-day<sup>[1](#footnote1)</sup> <= X | [Details](./dns_programming_latency.md) |
| __WIP__ | In-cluster network latency from a single prober pod, measured as latency of per second ping from that pod to "null service", measured as 99th percentile over last 5 minutes. | In default Kubernetes installataion with RTT between nodes <= Y, 99th percentile of (99th percentile over all prober pods) per cluster-day<sup>[1](#footnote1)</sup> <= X | [Details](./network_latency.md) |
| __WIP__ | In-cluster dns latency from a single prober pod, measured as latency of per second DNS lookup for "null service" from that pod, measured as 99th percentile over last 5 minutes. | In default Kubernetes installataion with RTT between nodes <= Y, 99th percentile of (99th percentile over all prober pods) per cluster-day<sup>[1](#footnote1)</sup> <= X | [Details](./dns_latency.md) |

<a name="footnote1">\[1\]</a> For the purpose of visualization it will be a
sliding window. However, for the purpose of SLO itself, it basically means
"fraction of good minutes per day" being within threshold.


### Other SLIs

| Status | SLI | User stories, ... |
| --- | --- | --- |
| WIP | Watch latency for every resource, (from the moment when object is stored in database to when it's ready to be sent to all watchers), measured as 99th percentile over last 5 minutes | [Details](./watch_latency.md) |
| WIP | Admission latency for each admission plugin type, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |
| WIP | Webhook call latency for each webhook type, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |
| WIP | Initializer latency for each initializer, measured as 99th percentile over last 5 minutes | [Details](./api_extensions_latency.md) |

