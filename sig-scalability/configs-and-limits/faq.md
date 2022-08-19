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


### What is the ideal size for API objects we should target?

Technically, the only hard limit that we have is the one of 1.5MB for
the size of individual objects. That said, approaching that limit is
definitely not recommended unless absolutely necessary.

In typical usecases, huge majority of objects doesn't exceed ~20KB of
size and this is the usecase that is best tested and many optimizations
assume (which are done based on existing tests) silently assume that.

If we look into individual objects larger than 20kB, significant majority
of cases that we've seen were representing a single pattern of grouping
multiple "subobjects" into a single object. The best example of that
from the core Kubernetes is `Endpoints` API, which is effectively grouping
all endpoints backing a given Kubernetes Service into a single object.
Those kind of APIs proved to be problematic for multiple different reasons,
including:
- the objects become large and even small change (of a single subobject)
  becomes expensive from the system perspective
- they become a contention point if different agents are updating different
  subobjects
- they become wasteful as we are able to get/watch only full objects and
  many agents may not need information about all subobjects
As a result, this pattern should really be avoided.
In case of `Endpoints` API, we introduce the `EndpointSlice` API and if
singular objects are problematic for your usecase, this is the pattern
you should explore.

So from scalability/performance perspective, the rule of thumb can be
summarized as:
- try to keep your object size below ~20kB
- if really needed, you can get to 100kB if it's not changing frequently
- if you can't keep your object size below 100kB, reach out to SIG
  Scalability and discuss the usecase to see how we can make it performant

### How should we code client applications to improve scalability?

As noted above, LIST requests can be particularly expensive. So when working with lists
that may have more than a few thousand elements, consider these guidelines:

1. When defining a new resource type (new CRD) consider expected numbers 
of objects that will exist (numbers of CRs).  See guidelines 
[here](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/95-custom-resource-definitions#scale-targets-for-ga).
1. If your code needs to hold an up-to-date list of objects in memory,
avoid repeated LIST calls if possible.  Instead consider using the
`Informer` classes that are provided in most Kubernetes client 
libraries.   Informers automatically combine LIST and WATCH functionality
to efficiently maintain an in-memory collection.
1. If `Informer`s don't suit your needs, try to use the API Server cache 
when LISTing.  To use the cache you must supply a `ResourceVersion`. 
Read the [documentation about ResourceVersions](https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions) carefully to understand how it will affect the 
freshness of the data you receive.
1. If you can't use `Informer`s AND you can't use the API Server cache,
 then be sure to [read large lists in chunks](https://kubernetes.io/docs/reference/using-api/api-concepts/#retrieving-large-results-sets-in-chunks).
1. Consider the number of instances of your client application which will be running. For instance,
there is a big difference between having 
just one controller listing objects, versus having demonsets on every node 
doing the same thing.  If there will be many instances of your client application
(either in daemonsets or some other form) you should be particularly careful
about LIST-related load.

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
