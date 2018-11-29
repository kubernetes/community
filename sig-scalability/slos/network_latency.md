## In-cluster network latency SLIs/SLOs details

### Definition

| Status | SLI | SLO |
| --- | --- | --- |
| __WIP__ | In-cluster network latency from a single prober pod, measured as latency of per second ping from that pod to "null service", measured as 99th percentile over last 5 minutes. | In default Kubernetes installataion with RTT between nodes <= Y, 99th percentile of (99th percentile over all prober pods) per cluster-day <= X |

### User stories
- As a user of vanilla Kubernetes, I want some guarantee how fast my http
request to some Kubernetes service reaches its endpoint

### Other notes
- We obviously can't give any guarantee in a general case, because cluster
administrators may configure cluster as they want.
- As a result, we define the SLI to be very generic (no matter how your cluster
is set up), but we provide SLO only for default installations with an additional
requirement that low-level RTT between nodes is lower than Y.
- Network latency is one of the most crucial aspects from the point of view
of application performance, especially in microservices world. As a result, to
meet user expectations, we need to provide some guarantees arount that.
- We decided for the SLI definition as formulated above, because:
  - it represents a user oriented end-to-end flow - it involves among others
    latency of in-cluster network programming mechanism (e.g. iptables). <br/>
    __TODO:__ We considered making DNS resolution part of it, but decided not
    to mix them. However, longer term we should consider joining them.
  - it is easily measurable in all running clusters in which we can run probers
    (e.g. measuring request latencies coming from all pods on a given
    node would require some additional instrumentation, such as a side car for
    each of them, and that overhead may be not acceptable in many cases)
  - it is not application-specific

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

### Test scenario

__TODO: Describe test scenario.__
