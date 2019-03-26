## DNS programming latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __WIP__ | Latency of programming dns instance, measured from when service spec or list of its `Ready` pods change to when it is reflected in that dns instance, measured as 99th percentile over last 5 minutes aggregated across all dns instances<sup>[1](#footnote1)</sup> | In default Kubernetes installation, 99th percentile per cluster-day <= X |

<a name="footnote1">[1\]</a>Aggregation across all programmers means that all
samples from all programmers go into one large pool, and SLI is percentile
from all of them.

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how quickly in-cluster
DNS will start resolving service name to its newly started backends.
- As a user of vanilla Kubernetes, I want some guarantee how quickly in-cluster
DNS will stop resolving service name to its removed (or unhealthy) backends.
- As a user of vanilla Kubernetes, I wasn some guarantee how quickly newly
create services will be resolvable via in-cluster DNS.

### Other notes
- We are consciously focusing on in-cluster DNS for the purpose of this SLI,
as external DNS resolution clearly depends on cloud provider or environment
in which the cluster is running (it hard to set the SLO for it).

### Caveats
- The SLI is aggregated across all DNS instances, which is what is interesting
for the end-user. It may happen that small percentage of DNS instances are
completely unresponsive (if all others are fast), but that is desired - we need
to allow slower/unresponsive ones because at some scale it will be happening.
The reason for doing it this way is feasibility for efficiently computing that:
  - if we would be doing aggregation at the SLI level (i.e. the SLI would be
    formulated like "... reflected in in-cluster load-balancing mechanism and
    visible from 99% of programmers"), computing that SLI would be extremely
    difficult. It's because in order to decide e.g. whether pod transition to
    Ready state is reflected, we would have to know when exactly it was reflected
    in 99% of programmers (e.g. iptables). That requires tracking metrics on
    per-change base (which we can't do efficiently).

### How to measure the SLI.
There [network programming latency](./network_programming_latency.md) is
formulated in almost exactly the same way. As a result, the methodology for
measuring the SLI here is exactly the same and can be found
[here](./network_programming_latency.md#how-to-measure-the-sli).

### Test scenario

__TODO: Describe test scenario.__
