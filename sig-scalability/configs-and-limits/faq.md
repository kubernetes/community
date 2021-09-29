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
