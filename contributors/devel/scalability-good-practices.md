# Scalability good practices for Kubernetes contributors

*This document is oriented at contributors who don't want their code to be reverted for performance reasons*

**Table of Contents**

- [Who should read this document and what's in it?](#who-should-read-this-document-and-whats-in-it)
- [TL;DR:](#tldr)
- [What does it mean to "break scalability"?](#what-does-it-mean-to-break-scalability)
- [Examples](#examples)
  - [Inefficient use of memory](#inefficient-use-of-memory)
  - [Explicit lists from the API server](#explicit-lists-from-the-api-server)
  - [Superfluous API calls](#superfluous-api-calls)
  - [Complex and expensive computations on a critical path](#complex-and-expensive-computations-on-a-critical-path)
  - [Big dependency changes](#big-dependency-changes)
- [Summary](#summary)
- [Closing remarks](#closing-remarks)

## Who should read this document and what's in it?
This document is targeted at developers of "vanilla Kubernetes" who don't want their changes rolled-back or blocked because of performance regressions they cause. It contains some of the knowledge and experience gathered by the scalability team in over two years.
 
It is presented as a set of real, or close-to-real examples from the past that broke scalability tests followed by some explanations and general suggestions on how to avoid causing similar problems.

## TL;DR:
1. You should understand how go memory management work to write code that is fast and doesn't use too much memory - pretty obvious, but we spent noticeable amount of time removing various bad patterns. In particular:
pass arguments by pointers wherever it's correct
avoid unnecessary copying of data (especially slices and maps)
Avoid unnecessary allocations (pre-size slices, reuse buffers, be aware of anonymous function definitions with variable captures, etc)

2. Wherever you want to write client.Sth().List(..) try using Informer (client-go/tools/cache/shared_informer.go). Be very, very sure you need to LIST resources if you're doing it <sup>[1](#1)</sup>.
 
3. Be careful when adding new API calls. Estimate how many QPS your change will add to the API server. If it's a big number, you need to find a way of doing it in a different way.
In particular be aware of “runs on every node” components, and the fan-in that can cause on the API server
 
4. From time to time you need to add a very complex logic to the system, one computation of which on big clusters may take a lot of time. You need to be sure that you won't add this logic on a critical path of any of standard flows.
 
5. When updating big dependencies (especially etcd or golang version) you need to run large scale test to verify that bump doesn't break Kubernetes performance.

## What does it mean to "break scalability"?
By "breaking scalability" we mean causing performance SLO violation in one of our performance tests. Performance SLOs for Kubernetes are <sup>[2](#2)</sup>:
- 99th percentile of API call latencies <= 1s
- 99th percentile of e2e Pod startup, excluding image pulling, latencies <= 5s

Tests that we run are Density and Load, we invite everyone interested in details to read the code.
 
We run those tests on big clusters (100+ Nodes), which means tests are somewhat resistant to limited concurrency in Kubelet (e.g. they're routinely failing on very small clusters, when scheduler can't spread Pod creations broadly enough).

## Examples
### Inefficient use of memory
*TL;DR:* You should understand how go memory management work to write code that is fast and doesn't use too much memory - pretty obvious, but we spent noticeable amount of time removing various bad patterns. In particular:
pass arguments by pointers wherever it's correct
avoid unnecessary copying of data (especially slices and maps)
 
Consider following (artificial) code snippet:
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

It contains a number of problems that were present in Kubernetes codebase always, and still keep coming. We actively fight them in most important places, but it's a neverending work.
 
First problem is that `func (s Scheduler) ScheduleOne…` means that each call of `ScheduleOne` will happen on new copy of the Scheduler object. This in turn means that go will need to copy whole `Scheduler` struct every time something calls `ScheduleOne` function and the copy will be discarded just after function returns. It's usually just a waste of resources, and even may be incorrect in some cases.
 
Secondly `(pod v1.Pod, nodes []v1.Nodes)` has a lot in common to the first problem. By default go passes arguments as values, i.e. copies them when they're passed to the function - *note that this is very different than Java or Python*. Note that there are things that are fine to pass directly. Slices, maps, strings and interfaces are actually pointers (in general interfaces might not be pointers, but in our code they are - see first point), so only a pointer value is copied when they're passed as an argument. For flat structures copying is sometimes necessary (e.g. when doing asynchronous modifications), but most often it's not - in such cases use pointers.
 
As there are no constant references in go this is the only option of passing objects without copying them (except creating read-only interfaces for all types, but this is infeasible). Note that it is, and should be, scary to pass a pointer to your object to strangers. Before you do so please make sure that the code you're passing the pointer to won't modify the object. Races are bad as well. Note that all `Informers` (see next paragraph) caches are expected to be immutable.
 
I could go on, and on, but I think the point is clear - when writing code that will be executed often you need to think about memory management. From time to time every one of us wants to be lazy and not think about it at all, but this pretty much always bites us back when we look at performance. General rules are: using heap is very expensive (garbage collection) - you should avoid unnecessary heap operations altogether, copying objects a lot is noticeable - it should be minimized. Learn how go manages memory. This is especially important in components running on the master. Otherwise we may end up in the situation when API server is starved on CPU and can't respond quickly to requests.

### Explicit lists from the API server
*TL;DR:* Wherever you want to write `client.Sth().List(..)` try using [`Informer`](client-go/tools/cache/shared_informer.go). Be very, very sure you need to LIST resources if you're doing it. 
 
Some time ago most of our controllers looked in the following way:

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

This looks fine-ish if you don't know that LIST are very expensive calls. Object can have size of few kilobytes, and we can have 150k of those. This means that LIST would need to send those hundreds of megabytes through the network, not to mention the fact that API server would need to do few conversions along the way of all this data. It's not the end of the world, but it needs to be minimized. Especially that the solution is simple (quoting Clayton):

>As a rule, use informer. If using informer, use shared informers. If your use case doesn't look like an informer, look harder. If at the very end of that, and it still doesn't look like an informer, maybe use something else after talking to someone. But probably use informer.

`Informer` is our library that provides a read interface of the store - it's a read-only cache that provides you a local copy of the store that will contain only object that you're interested in (matching given selector). From it you can GET, LIST, or do whatever read operations you want. `Informer` also allows you to register functions that will be called when an object is created, modified or deleted, which is what most people want.
 
The magic behind `Informers` is that they are populated by the WATCH, so they don't stress API server too much. Code for Informer is [here](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/tools/cache/shared_informer.go).
 
In general: use `Informers` - if we were able to rewrite most vanilla controllers to them, you'll be able to do it as well. If you don't you may dramatically increase CPU requirements of the API server which will starve it and make it too slow to meet our SLOs.

### Superfluous API calls
*TL;DR:* Be careful when adding new API calls. Estimate how many QPS your change will add to the API server. If it's a big number, you need to find a way of doing it in a different way.
 
One past regression was caused by `Secret` refreshing logic in Kubelet. By contract we want to update values of `Secrets` (update env variables, contents of `Secret` volume) when the contents of `Secret` are updated in the API server. Normally we'd use `Informer` (see above), but there's additional security constraint, that Kubelet should know only `Secrets` that are attached to `Pods` scheduled on the corresponding `Node`. So no watching all `Secret` updates (which is how `Informers` work). We already know that LISTs are also bad (not to mention that they have the same security problem as WATCH), so only way we can read `Secrets` is through GET.
 
So for each `Secret` periodically we were GETting it's value and updating underlying variables/volumes if necessary. We have the same logic for `ConfigMaps`. Everything was great until we turned on `ServiceAccount` admission controller in our performance tests. Then everything went to hell for a very simple reason. `ServiceAccount` admission controller creates a `Secret` that it attached to every `Pod` (different one in every Namespace, but this doesn't change anything). Multiply above behavior by 150k. Given refresh period of 60s it meant that additional 2.5k QPS went to the API server, which of course blew up.
 
What we did to mitigate this issue was, in a way, reimplementing Informers using GETs instead of WATCHes. Current solution consists of a `Secret` cache shared between all `Pod`s. When a `Pod` wants to check if `Secret` changed to looks into the cache. If the `Secret` stored in the cache is too old, cache issues a GET request to API server to get current value. As `Pods` within a single `Namespace` share the `Secret` for `ServiceAccount` it means that Kubelet will need to refresh the `Secret` only once a while per `Namespace`, not per `Pod`, as it was before. This of course is a stopgap, not a final solution, which is currently (as of early May 2017) being designed as a ["Bulk Watch"](https://github.com/kubernetes/community/pull/443).
 
This example should show that you need to treat API calls as a rather expensive shared resource. It's fine to use some of it, but there should be a good reason for increasing number of calls drastically, as you share this resource with everything else. This is especially important on the Node side, as every change is multiplied by 5000. In controllers, especially when writing some disaster recovery logic it's perfectly fine to add a new call - first of all there are not a lot of controllers, secondly disaster recovery shouldn't happen too often. That being said whenever you add a new API server request you should do quick estimation of QPS that will be added to the API server, and if you get a noticeable number you probably should think about a way to reduce it.
 
If you don't do that a lot of things may break. One thing is the same as previously - starve API server on CPU. But this particular pattern, if particularly bad, can also drain `max-inflight-request` in API server which will make it respond with 429s (Too Many Requests) on other requests and thus make the whole system slower. At best it will only cause draining of local client rate limiter for API calls in your component (default value is 5 QPS, controllers normally have 20). This will result in your component being very, very, slow.

### Complex and expensive computations on a critical path
*TL;DR:* From time to time you need to add a very complex logic to the system, one computation of which on big clusters may take a lot of time. You need to be sure that you won't add this logic on a critical path of any of standard flows.
 
As an example we'll give a `PodAntiAffinity` scheduling feature. The goal of this feature is to allow users to prevent co-scheduling of `Pods`, for very broad definition of co-scheduling. When defining `PodAntiAffinity` you pass two things: `Node` grouping and `Pod` selector. Semantics is following: for each group of `Nodes` you check if any `Node` in the group runs a `Pod` matching the selector. If it does all `Nodes` from the group are discarded. This of course needs to be symmetric, as if you prevent pods from set A to be co-scheduled with `Pods` from set B, but not other way around, then when adding new `Pod` to set B, you'll end up with `Pods` from A and B running in the same group, which you wanted to avoid.
 
This means that even when scheduling `Pods` that do not explicitly use `PodAntiAffinity` feature you need to check `PodAntiAffinities` of all `Pods` running in the cluster. This means that scheduling of every `Pod` gets additional check with `O(#Pods * #Nodes)` complexity - if implemented naively. Given the fact that we can have 150k `Pods` in the cluster it's generally not a good idea to have quadratic algorithms on a critical path for Pods - even for ones that don't use PodAntiAffinity feature!
 
First time it was implemented it was actually implemented in a very simple way making the scheduler unusable pretty quickly, and `Pod` startup time went through the roof. We had to block this feature at that point and it didn't made into the release it was aimed at. Later we fixed it slightly by improving the algorithm to `O(#(scheduled Pods with PodAntiAffinity) * #Nodes)`, which was enough to allow this feature to get in as beta with a huge asterisk by it.
 
This should serve as an example for the fact that we're working in a very complex area, where a lot of problems are actually very complex - not only non-linear, but some of them are NP-complete. Because of that it's understandable that you need to write something complex. But when you do that you need to take care to protect the rest of the system from that complexity, and add it only in the place where it's absolutely necessary.

### Big dependency changes
*TL;DR:* When updating big dependencies (especially etcd or golang version) you need to run large scale test to verify that bump doesn't break Kubernetes performance.
 
Kubernetes depends on pretty much the whole universe. From time to time we need to update some dependencies (Godeps, etcd, go version). This can break us in many ways, and we saw it doing it a couple of times now. We skipped one version of `go` (1.5) exactly because of breaking our performance, and in the moment of writing this doc we're working with go team on understanding why `go1.8` also makes Kubernetes performance noticeably worse.
 
If you're changing some big and important dependency, there's no way of knowing what performance impact it'll have. The only thing we can do is to actually run test and check.
Where to look to get data?
If you want to check the impact of your changes there are a number of places you can look at:
Density/Load tests output quite a lot of data either to test logs, or files inside 'ReportDir' - both of them include API call latencies and Density also includes pod e2e startup latency information,
for resource usage you can either use monitoring tools (heapster + graphana, but at the time of writing it stops working around 100 Nodes), or just plain old top on master (which scales as much as you want:),
more data is available on `/metrics` endpoint on all our components (e.g. one for the API server contains API call latencies),
to profile a component is to create an ssh tunnel to the machine running it, and run `go tool pprof localhost:<your_tunnel_port>` locally

## Summary
Summing it up, when writing code you should:
- understand how go manages memory and use it wisely,
- not LIST from API server,
- run performance tests in case of huge system-wide changes (e.g. updating big dependencies),

When designing new features/thinking about refactoring you should:
- when adding new API calls estimate number of additional QPS you'll be sending to the API server,
- when adding complex logic make sure that you don't add it on a critical path of any basic workflow

## Closing remarks
We know that thinking about performance impact of changes that all of us write is hard, this is exactly why we want you to help us cater for it, by always having problems that we mentioned here in the back of your mind. In return we'll answer all your question and doubts about possible impact of your changes if you post them either to #sig-scale Slack channel, or cc @kubernetes/sig-scalability-pr-reviews in your PR/proposal.

* * *

<a name="1">1</a>: If you are using List in tight loops, it's common to be doing that on a subset of a list (field, label, or namespace). Most informers have indices on namespaces, but you may end up needing another index (but only if profile shows it)

<a name="2">2</a>: As of now. We're currently working on adding new SLOs and improving the system to meet them.
