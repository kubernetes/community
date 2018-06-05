## API call latency SLIs/SLOs details

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly I get the
response from an API call.
- As an administrator of Kubernetes cluster, if I know characteristics of my
external dependencies of apiserver (e.g custom admission plugins, webhooks and
initializers) I want to be able to provide guarantees for API calls latency to
users of my cluster

### Other notes
- We obviously can’t give any guarantee in general, because cluster
administrators are allowed to register custom admission plugins, webhooks
and/or initializers, which we don’t have any control about and they obviously
impact API call latencies.
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

__TODO: Descibe test scenario.__
