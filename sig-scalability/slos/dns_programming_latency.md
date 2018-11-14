## Network programming latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __WIP__ | Latency of programming a single in-cluster dns instance, measured from when service spec or list of its `Ready` pods change to when it is reflected in that dns instance, measured as 99th percentile over last 5 minutes | In default Kubernetes installation, 99th percentile of (99th percentiles across all dns instances) per cluster-day <= X |

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
- The SLI is formulated for a single DNS instance, even though that value
itself is not very interesting for the user.
If there are multiple DNS instances in the cluster, the aggregation across
them is done only at the SLO level (and only that gives a value that is
interesting for the user). The reason for doing it this is feasibility for
efficiently computing that:
  - if we would be doing aggregation at the SLI level (i.e. the SLI would be
    formulated like "... reflected in in-cluster DNS and visible from 99%
    of DNS instances"), computing that SLI would be extremely
    difficult. It's because in order to decide e.g. whether pod transition to
    Ready state is reflected, we would have to know when exactly it was reflected
    in 99% of DNS instances. That requires tracking metrics on
    per-change base (which we can't do efficiently).
  - we admit that the SLO is a bit weaker in that form (i.e. it doesn't necessary
    force that a given change is reflected in 99% of programmers with a given
		99th percentile latency), but it's close enough approximation.

### How to measure the SLI.
There [network programming latency](./network_programming_latency.md) is
formulated in almost exactly the same way. As a result, the methodology for
measuring the SLI here is exactly the same and can be found
[here](./network_programming_latency.md#how-to-measure-the-sli).

### Test scenario

__TODO: Describe test scenario.__
