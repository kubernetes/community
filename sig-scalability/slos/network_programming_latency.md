## Network programming latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __WIP__ | Latency of programming in-cluster load balancing mechanism (e.g. iptables), measured from when service spec or list of its `Ready` pods change to when it is reflected in load balancing mechanism, measured as 99th percentile over last 5 minutes aggregated across all programmers<sup>[1](#footnote1)</sup> | In default Kubernetes installation, 99th percentile per cluster-day <= X |

<a name="footnote1">[1\]</a>Aggregation across all programmers means that all
samples from all programmers go into one large pool, and SLI is percentile
from all of them.

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly new backends
of my service will be targets of in-cluster load-balancing
- As a user of vanilla Kubernetes, I want some guarantee how quickly deleted
(or unhealthy) backends of my service will be removed from in-cluster
load-balancing
- As a user of vanilla Kubernetes, I want some guarantee how quickly changes
to service specification (including creation) will be reflected in in-cluster
load-balancing

### Other notes
- We are consciously focusing on in-cluster load-balancing for the purpose of
this SLI, as external load-balancing is clearly provider specific (which makes
it hard to set the SLO for it).
- However, in the future it should be possible to formulate the SLI for external
load-balancing in pretty much the same way for consistency.
- The SLI measuring end-to-end time from pod creation was also considered,
but rejected due to being application specific, and thus introducing SLO would
be impossible.

### Caveats
- The SLI is aggregated across all "programmers", which is what is interesting
for the end-user. It may happen that small percentage of programmers are
completely unresponsive (if all others are fast), but that is desired - we need
to allow slower/unresponsive nodes because at some scale it will be happening.
The reason for doing it this way is feasibility for efficiently computing that:
  - if we would be doing aggregation at the SLI level (i.e. the SLI would be
    formulated like "... reflected in in-cluster load-balancing mechanism and
    visible from 99% of programmers"), computing that SLI would be extremely
    difficult. It's because in order to decide e.g. whether pod transition to
    Ready state is reflected, we would have to know when exactly it was reflected
    in 99% of programmers (e.g. iptables). That requires tracking metrics on
    per-change base (which we can't do efficiently).

### How to measure the SLI.
The method of measuring this SLI is not obvious, so for completeness we describe
it here how it will be implemented with all caveats.
1. We assume that for the in-cluster load-balancing programming we are using
Kubernetes `Endpoints` objects.
1. We will introduce a dedicated annotation for `Endpoints` object (name TBD).
1. Endpoints controller (while updating a given `Endpoints` object) will be
setting value of that annotation to the timestamp of the change that triggered
this update:
- for pod transition between `Ready` and `NotReady` states, its timestamp is
  simply part of pod condition
- TBD for service updates (ideally we will add `LastUpdateTimestamp` field in
  object metadata next to already existing `CreationTimestamp`. The data is
  already present at storage layer, so it won't be hard to propagate that.
1. The in-cluster load-balancing programmer will export a prometheus metric
once done with programming. The latency of the operation is defined as
difference between timestamp of then whe operation is done and timestamp
recorded in the newly introduced annotation.

#### Caveats
There are a couple of caveats to that measurement method:
1. Single `Endpoints` object may batch multiple pod state transition. <br/>
In that case, we simply choose the oldest one (and not expose all timestamps
to avoid theoretically unbounded growth of the object). That makes the metric
imprecise, but the batching period should be relatively small comparing
to whole end-to-end flow.
1. A single pod may transition its state multiple times within batching
period. <br/>
For that case, we will add additional cache in Endpoints controller caching
the first observed transition timestamp for each pod. The cache will be
cleared when controller picks up a pod into Endpoints object update. This is
consistent with choosing the oldest update in the above point. <br/>
Initially, we may consider simply ignoring this fact.
1. Components may fall out of watch window history and thus miss some watch
events. <br/>
This may be the case for both Endpoints controller or kube-proxy (or other
network programmers if used instead). That becomes a problem when a single
object changed multiple times in the meantime (otherwise informers will
deliver handlers on relisting). Additionally, this can happen only when
components are too slow in processing events (that would already be reflected
in metrics) or (sometimes) after kube-apiserver restart. Given that, we are
going to neglect this problem to avoid unnecessary complications for little
or no gain.

### Test scenario

__TODO: Describe test scenario.__
