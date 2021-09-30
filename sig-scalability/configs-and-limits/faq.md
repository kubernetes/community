# FAQ

### How many QPS kube-apiserver can support?

There is NO single answer to this question. The reason for it is that,
every Kubernetes request is potentially very different. As a result,
kube-apiserver may be able to process N in a given usage pattern, but
may not be able to process even N/5 in a different usage pattern.

To be more specific, the difference is fairly obvious if we compare
GET and LIST requests (e.g. "GET my-pod" vs "LIST all pods").
It's also not counter-intuitive that there is a difference between
ready-only and mutating requests (e.g. "GET my-pod" vs "POST my-pod").

However, even if we focus only on a single type of request (let's say just
POST or just PUT), the differences can be in the orders of magnitude.
The two main reasons to think about are:
- size of the object - there is a huge difference between a small Lease
  object (say ~300B) and large custom CRD of e.g. 1MB
- fan-out - if a given object is being watched by N watchers, you
  suddenly get a multiplication factor for the cost of your request,
  that as a sender of a single API call you don't fully control

As a result, we are consciously not providing any values, because
one can imagine a cluster that handles thousands of QPS and a cluster
that is set up the same way may have troubles handling small tens of
QPS if the requests are fairly expensive.


### How do you setup clusters for scalability testing?

We are testing Kubernetes on two levels of scale: 100 nodes and 5000 nodes.
Given that Kubernetes scalability is a multidimensional problem, we are
obviously trying to exercise many other dimensions in our tests.

However, all our scalability tests are performed on the cluster with just
a single large control-plane machine (VM) with all control-plane components
running that machine.
The single large machine provides sufficient scalability to let us run
our load tests on a cluster with 5000 nodes.

The control-plane VM used for testing 5000-node clusters has 64 cores,
256 GB of RAM and is backed by 200GB SSD persistent disk. Public cloud
providers typically support even large machines, so as a user you still
have some slack here.

We are also running a `simulated` clusters ([Kubemark]) where the control
plane VM is set up the same was as in regular cluster, but many simulated
nodes are running on a single VM (to simmulate 5000 node clusters, we run
~80 machines with 60-ish hollow-nodes per machine). However, those are
currently somewhat auxiliary and release validation is done using a regular
clusters setup on GCE cloud provider.

[Kubemark]: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-scalability/kubemark-guide.md


### Why you are not testing HA cluster (with multiple control-plane instances)?

Honestly, we would like to. But we are lacking capacity to do that.
The main reason is the tooling to setup cluster - we are still relying on
kube-up and a bunch of knobs we have there. And kube-up doesn't really
support HA clusters. However, migration to anything else is a lot of work.
At the same time we would like to do that consistently for all other jobs
running on GCE (which all of still rely on kube-up).

We are aware that this is a gap in our testing for several reasons, including:
* large production clusters generally run at least 3 control plane instances to
  ensure tolerance to outages of individual instances, zero-downtime upgrade etc.
  we should be testing what users are doing in real life.
* etcd is RAFT-based so a cluster of etcd instance have slightly different
  performance characteristics that a single etcd server
* distributed systems often show different performance characteristics when the
  components are separate by a network
* distributed system may be affected by inconsistency of caches

Noting difference between various configuration is also important as this
can give indications which optimizations would have the best return-on-investments.
