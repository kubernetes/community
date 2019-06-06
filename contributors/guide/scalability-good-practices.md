---
title: "Scalability Good Practices"
weight: 1
slug: "scalability-good-practices" 
---

*This document is written for contributors who would like to avoid their code being reverted for performance reasons*

**Table of Contents**

- [Who should read this document and what is in it?](#who-should-read-this-document-and-what-is-in-it)
- [What does it mean to "break scalability"?](#what-does-it-mean-to-break-scalability)
- [Examples](#examples)
  - [Inefficient use of memory](#inefficient-use-of-memory)
  - [Explicit lists from the API server](#explicit-lists-from-the-api-server)
  - [Superfluous API calls](#superfluous-api-calls)
  - [Complex and expensive computations on a critical path](#complex-and-expensive-computations-on-a-critical-path)
  - [Big dependency changes](#big-dependency-changes)
- [Summary](#summary)
- [Closing remarks](#closing-remarks)

## Who should read this document and what is in it?
This document is targeted at developers of "vanilla Kubernetes" who do not want their changes rolled-back or blocked because they cause performance regressions. It contains some of the knowledge and experience gathered by the scalability team over more than two years.
 
It is presented as a set of examples from the past which broke scalability tests, followed by some explanations and general suggestions on how to avoid causing similar problems.

## What does it mean to "break scalability"?
"Breaking scalability" means causing performance SLO violations in one of our performance tests. Performance SLOs for Kubernetes are <sup>[2](#2)</sup>:
- 99th percentile of API call latencies <= 1s
- 99th percentile of e2e Pod startup, excluding image pulling, latencies <= 5s

We run density and load tests, and we invite anyone interested in the details to read the code.
 
We run those tests on large clusters (100+ Nodes). This means tests are somewhat resistant to limited concurrency in Kubelet (e.g. they are routinely failing on very small clusters, when the Scheduler cannot spread Pod creations broadly enough).

## Examples
### Inefficient use of memory
Consider the following sample code snippet:
```golang
func (s Scheduler) ScheduleOne(pod v1.Pod, nodes []v1.Nodes) v1.Node {
  for _, node := range nodes {
    if s.FitsNode(pod, node) {
      …
    }
  }
}
 
func (s Scheduler) DoSchedule(podsChan chan v1.Pod) {
  for {
    …
    node := s.ScheduleOne(pod, s.nodes)
    …
  }
}
```

This snippet contains a number of problems that were always present in the Kubernetes codebase, and continue to appear. We try to address them in the most important places, but the work never ends.
 
The first problem is that `func (s Scheduler) ScheduleOne…` means each call of `ScheduleOne` will run on a new copy of the Scheduler object. This in turn means Golang will need to copy the entire `Scheduler` struct every time the `ScheduleOne` function is called. The copy will then be discarded when the function returns. Clearly, this is a waste of resources, and in some cases may be incorrect.
 
Next, `(pod v1.Pod, nodes []v1.Nodes)` has much in common with the first problem. By default, Golang passes arguments as values, i.e. copies them when they are passed to the function. *Note that this is very different from Java or Python*. Of course, some things are fine to pass directly. Slices, maps, strings and interfaces are actually pointers (in general interfaces might not be pointers, but in our code they are - see first point), so only a pointer value is copied when they are passed as an argument. For flat structures, copying is sometimes necessary (e.g. when doing asynchronous modifications), but most often it is not. In such cases, use pointers.
 
As there are no constant references in Golang, this is the only option for passing objects without copying them (except creating read-only interfaces for all types, but that is not feasible). Note that it is (and should be) scary to pass a pointer to your object to strangers. Before you do so, make sure the code to which you are passing the pointer will not modify the object. Races are bad as well. Note that all `Informers` (see next paragraph) caches are expected to be immutable.
 
We could go on and on, but the point is clear -- when writing code that will be executed often, you need to think about memory management. From time to time we all occasionally forget to keep this in mind, but we are reminded of it when we look at performance. General rules are:
- Using heap is very expensive (garbage collection)
- Avoid unnecessary heap operations altogether
- Repeatedly copying objects is noticeable and should be minimized.
- Learn how Golang manages memory. This is especially important in components running on the control plane. Otherwise we may end up in the situation where the API server is starved on CPU and cannot respond quickly to requests.

### Explicit lists from the API server
Some time ago most of our controllers looked like this:

```golang
func (c *ControllerX) Reconcile() {
  items, err := c.kubeClient.X(v1.NamespaceAll).List(&v1.ListOptions{})
  if err != nil {
    ...
  }
  for _, item := range items {
    ...
  }
}
 
func (c *ControllerX) Run() {
  wait.Until(c.Reconcile, c.Period, wait.NeverStop)
  ...
}
```

This may look OK, but List() calls are expensive. Objects can have sizes of a few kilobytes, and there can be 150,000 of those. This means List() would need to send hundreds of megabytes through the network, not to mention the API server would need to do conversions of all this data along the way. It is not the end of the world, but it needs to be minimized. The solution is simple (quoting Clayton):

>As a rule, use Informer. If using Informer, use shared Informers. If your use case does not look like an Informer, look harder. If at the very end of that it still does not look like an Informer, consider using something else after talking to someone. But probably use Informer.

`Informer` is our library which provides a read interface to the store - it is a read-only cache that provides you with a local copy of the store that contains only the object you are interested in (matching given selector). From it you can Get(), List() or whatever read operations you desire. `Informer` also allows you to register functions that will be called when an object is created, modified or deleted.
 
The magic behind `Informers` is that they are populated by the WATCH, so they create minimal stress on the API server. Code for Informer is [here](https://git.k8s.io/kubernetes/staging/src/k8s.io/client-go/tools/cache/shared_informer.go).
 
In general: use `Informers` - if we were able to rewrite most vanilla controllers to use them, you should be able to do so as well. Otherwise, you may dramatically increase the CPU requirements of the API server which will starve it and make it too slow to meet our SLOs.

### Superfluous API calls
One past regression was caused by `Secret` refreshing logic in Kubelet. By contract we want to update values of `Secrets` (update env variables, contents of `Secret` volume) when the contents of `Secret` are updated in the API server. Normally we would use `Informer` (see above), but there is an additional security constraint; Kubelet should know only `Secrets` that are attached to `Pods` scheduled on the corresponding `Node`, so there should be no watching of all `Secret` updates (which is how `Informers` work). We already know that List() calls are also bad (not to mention that they have the same security problem as WATCH), so the only way we can read `Secrets` is through GET.
 
For each `Secret` we were periodically GETting its value and updating underlying variables/volumes as necessary. We have the same logic for `ConfigMaps`. Everything was great until we turned on the `ServiceAccount` admission controller in our performance tests. Then everything went wrong for a very simple reason; the `ServiceAccount` admission controller creates a `Secret` that it attaches to every `Pod` (a different one in every Namespace, but this does not change anything). Multiply this behavior by 150,000 and, given a refresh period of 60 seconds, an additional 2.5k QPS were being sent to the API server, which of course caused it to fail.
 
To mitigate this issue we had to reimplement Informers using GETs instead of WATCHes. The current solution consists of a `Secret` cache shared between all `Pod`s. When a `Pod` wants to check if the `Secret` has changed it looks in the cache. If the `Secret` stored in the cache is too old, the cache issues a GET request to the API server to refresh the value. As `Pods` within a single `Namespace` share the `Secret` for `ServiceAccount`, it means Kubelet will need to refresh the `Secret` only once in a while per `Namespace`, not per `Pod`, as it was before. This of course is a stopgap and not a final solution, which is currently (as of early May 2017) being designed as a ["Bulk Watch"](https://github.com/kubernetes/community/pull/443).
 
This example demonstrates why you need to treat API calls as a rather expensive shared resource. This is especially important on the Node side, as every change is multiplied by 5,000. In controllers, especially when writing some disaster recovery logic, it is perfectly fine to add a new call. There are not a lot of controllers, and disaster recovery should not happen too often. That being said, whenever you add a new API server request you should do quick estimation of QPS that will be added to the API server, and if the result is a noticeable number you probably should think about a way to reduce it.
 
One obvious consequence of not reducing API calls is that you will starve the API server on CPU. This particular pattern can also drain `max-inflight-request` in the API server, which will make it respond with 429's (Too Many Requests) and thus slow down the system. At best it will only cause draining of the local client rate limiter for API calls in your component (default value is 5 QPS, controllers normally have 20). This will result in your component being very, very slow.

### Complex and expensive computations on a critical path
Let us use the `PodAntiAffinity` scheduling feature as an example. The goal of this feature is to allow users to prevent co-scheduling of `Pods` (using a very broad definition of co-scheduling). When defining `PodAntiAffinity` you pass two things: `Node` grouping and `Pod` selector. The semantics is that for each group of `Nodes` you check if any `Node` in the group runs a `Pod` matching the selector. If it does, all `Nodes` from the group are discarded. This of course needs to be symmetric, as if you prevent pods from set A to be co-scheduled with `Pods` from set B, but not the other way around. When adding new `Pod` to set B, you'll end up with `Pods` from A and B running in the same group, which you wanted to avoid.
 
This means that even when scheduling `Pods` that do not explicitly use the `PodAntiAffinity` feature you need to check `PodAntiAffinities` of all `Pods` running in the cluster. It also means that scheduling of every `Pod` gets an additional check of `O(#Pods * #Nodes)` complexity, if naively implemented. Given the fact that we can have 150.000 `Pods` in the cluster, it becomes obvious it is not a good idea to have quadratic algorithms on a critical path for Pods - even for ones that do not use the PodAntiAffinity feature!
 
This was initially implemented in a very simple way, rapidly making the scheduler unusable, and `Pod` startup times went through the roof. We were forced to block this feature, and it did not make into the target release. Later, we slightly improved the algorithm to `O(#(scheduled Pods with PodAntiAffinity) * #Nodes)`, which was enough to allow the feature to get in as beta, with a huge asterisk next to it.
 
This example illustrates how many problems in this area can be much more complex than they seem. Not only that, they are non-linear, and some of them are NP-complete. Understandably, sometimes you need to write something complex, but when you do, you must protect the rest of the system from that complexity, and add it only where it is absolutely necessary.

### Big dependency changes
Kubernetes depends on pretty much the whole universe. From time to time we need to update some dependencies (Godeps, etcd, go version). This can break us in many ways, as has already happened a couple of times. We skipped one version of Golang (1.5) precisely because it broke our performance. As this is being written, we are working with the Golang team to try to understand why Golang version 1.8 negatively affects Kubernetes performance.

If you are changing a large and important dependency, the only way to know what performance impact it will have is to run test and check.

#### Where to look to get data?

If you want to check the impact of your changes there are a number of places to look.
- Density and load tests output quite a lot of data either to test logs, or files inside 'ReportDir' - both of them include API call latencies, and density tests also include pod e2e startup latency information.
- For resource usage you can either use monitoring tools (heapster + Grafana, but note that at the time of writing, this stops working at around 100 Nodes), or just plain 'top' on the control plane (which scales as much as you want),
- More data is available on the `/metrics` endpoint of all our components (e.g. the one for the API server contains API call latencies),
to profile a component create an ssh tunnel to the machine running it, and run `go tool pprof localhost:<your_tunnel_port>` locally

## Summary
To summarize, when writing code you should:
- understand how Golang manages memory and use it wisely,
- not List() from the API server,
- run performance tests when making large systemwide changes (e.g. updating big dependencies),

When designing new features or thinking about refactoring you should:
- Estimate the number of additional QPS you will be sending to the API server when adding new API calls 
- Make sure to not add any complex logic on a critical path of any basic workflow

## Closing remarks
We know that thinking about the performance impact of changes is hard. This is exactly why we want you to help us cater for it, by keeping all the knowledge we have given you here in the back of your mind as you write your code. In return, we will answer all your question and doubts about possible impact of your changes if you post them either to #sig-scalability Slack channel, or cc @kubernetes/sig-scalability-pr-reviews in your PR/proposal.

* * *

<a name="1">1</a>: If you are using List() in tight loops, it is common to do so on a subset of a list (field, label, or namespace). Most Informers have indices on namespaces, but you may end up needing another index if profile shows the need.

<a name="2">2</a>: We are working on adding new SLOs and improving the system to meet them.
