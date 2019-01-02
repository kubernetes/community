## In-cluster dns latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __WIP__ | In-cluster dns latency from a single prober pod, measured as latency of per second DNS lookup<sup>[1](#footnote1)</sup> for "null service" from that pod, measured as 99th percentile over last 5 minutes. | In default Kubernetes installataion with RTT between nodes <= Y, 99th percentile of (99th percentile over all prober pods) per cluster-day <= X |

<a name="footnote1">\[1\]</a> In fact two DNS lookups: (1) to nameserver IP from
/etc/resolv.conf (2) to kube-system/kube-dns service IP and track them as two
separate SLIs.

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how fast my in-cluster
DNS requests are resolved

### Other notes
- We obviously can't give any guarantee in a general case, because cluster
administrators may configure cluster as they want.
- As a result, we define the SLI to be very generic (no matter how your cluster
is set up), but we provide SLO only for default installations with an additional
requirement that low-level RTT between nodes is lower than Y.
- DNS latency is one of the most crucial aspects from the point of view
of application performance, especially in microservices world. As a result, to
meet user expectations, we need to provide some guarantees arount that.
- We are introducing two SLIs (for two IP addresses) to enable measuring the
impact of node-local caching.

### Caveats
- The SLI is formulated for a prober pods, even though users are mostly
interested in the aggregation across all pods (that is done only at the SLO
level). However, that provides very similar guarantees and makes it fairly
easy to measure.
- The RTT between nodes may significantly differ, if nodes are in different
topologies (e.g. GCP zones). However, given that topology-aware service routing
is not natively supported in Kubernetes yet, we explicitly acknowledge that
depending on the pinged endpoint, results may signiifcantly differ if nodes
are spanning multiple topologies.
- The prober reporting that is fairly trivial and itself needs only negligible
amount of resources. Unfortunately there isn't any component to which we can
attach that functionality (e.g. KubeProxy is running in host network), so
**we will create a dedicated set of prober pods**. We will run a set of prober
pods (number proportional to cluster size).
- We don't have any "null service" running in cluster, so an administrator has
to set up one to make the SLI measurable in real cluster. In tests, we will
create a service on top of prober pods.

### TODOs
- DNS Latency is only a part of criticial metrics, the other being "drop rate"
or "timeout rate". Given that is seems harder to measure/sample, we would like
to address that separate to avoid blocking this SLI on the resolution.

### Test scenario

__TODO: Describe test scenario.__
