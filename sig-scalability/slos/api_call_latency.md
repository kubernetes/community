## API call latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __Official__ | Latency<sup>[1](#footnote1)</sup> of mutating<sup>[2](#footnote2)</sup> API calls for single objects for every (resource, verb) pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, verb) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day <= 1s |
| __Official__ | Latency<sup>[1](#footnote1)</sup> of non-streaming read-only<sup>[3](#footnote3)</sup> API calls for every (resource, scope<sup>[4](#footnote4)</sup>) pair, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, for every (resource, scope) pair, excluding virtual and aggregated resources and Custom Resource Definitions, 99th percentile per cluster-day: (a) <= 1s if `scope=resource` (b) <= 30s<sup>[5](#footnote5)</sup> otherwise (if `scope=namespace` or `scope=cluster`) |

<a name="footnote1">\[1\]</a> By latency of API call in this doc we mean time
from the moment when apiserver gets the request to last byte of response sent
to the user.

<a name="footnote2">\[2\]</a> By mutating API calls we mean POST, PUT, DELETE
and PATCH.

<a name="footnote3">\[3\]</a> By non-streaming read-only API calls we mean GET
requests without `watch=true` option set. (Note that in Kubernetes internally
it translates to both GET and LIST calls).

<a name="footnote4">\[4\]</a> A scope of a request can be either
- `resource` - if the request is about a single object
- `namespace` - if it is about objects from a single namespace
- `cluster` - if it spawns objects from multiple namespaces

<a name="footnote5">\[5\]</a> Historically, the threshold for LISTs with
`scope=namespace` was set to 5 seconds. However, the threshold was chosen when
Kubernetes didn't support the scale it supports today and when individual
namespace didn't contain tens of thousands (if not more) objects of a given
type. We adjusted the limits to accommodate the usage patterns change, given
that users are fine with listing tens of thousands of objects taking more than
5 seconds.

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly I get the
response from an API call.
- As an administrator of Kubernetes cluster, if I know characteristics of my
external dependencies of apiserver (e.g custom admission plugins and webhooks)
I want to be able to provide guarantees for API calls latency to users of my
cluster.

### Other notes
- We obviously can’t give any guarantee in general, because cluster
administrators are allowed to register custom admission plugins or webhooks,
which we don’t have any control about and they obviously impact API call
latencies.
- As a result, we define the SLIs to be very generic (no matter how your
cluster is set up), but we provide SLO only for default installations (where we
have control over what apiserver is doing). This doesn’t provide a false
impression, that we provide guarantee no matter how the cluster is setup and
what is installed on top of it.
- At the same time, API calls are part of pretty much every non-trivial workflow
in Kubernetes, so this metric is a building block for less trivial SLIs and
SLOs.
- The SLO for latency for read-only API calls of a given type may have significant
buffer in threshold. In fact, the latency of the request should be proportional to
the amount of work to do (which is number of objects of a given type in a given
scope) plus some constant overhead. For better tracking of performance, we
may want to define purely internal SLI of "latency per object". But that
isn't in near term plans.
- To recall, SLOs are guaranteed only if thresholds defined in [thresholds file][]
are satisfied. This is particularly important for this SLO, because it limits
the number of objects that are returned by LIST calls.

[thresholds file]: https://github.com/kubernetes/community/blob/master/sig-scalability/configs-and-limits/thresholds.md

### Caveats
- The SLO has to be satisfied independently from used encoding in user-originated
requests. This makes mix of client important while testing. However, we assume
that all `core` components communicate with apiserver using protocol buffers.
- In case of GET requests, user has an option opt-in for accepting potentially
stale data (being served from cache) and the SLO again has to be satisfied
independently of that. This makes the careful choice of requests in tests
important.

### TODOs
- We may consider treating `non-namespaced` resources as a separate bucket in
the future. However, it may not make sense if the number of those may be
comparable with `namespaced` ones.

### Test scenario

__TODO: Describe test scenario.__
