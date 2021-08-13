## 2017-12-28 (Cancelled - Christmas week) 


## 2017-12-14



*   Kubernetes Contributors Summit report by Solly
*   1.9 release notes:  due EOD today


## 2017-11-30



*   Ways of exporting Counter metrics from the subprocess in Prometheus. \
https://github.com/prometheus/snmp_exporter


## 2017-11-16



*   Update for community meeting
    *   Core and Custom Metrics APIs promoted to beta
    *   Multiple kube-state-metrics releases (many new metrics, stability, features)
    *   Prometheus backed metrics API gateway: [https://github.com/DirectXMan12/k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter)
    *   Current work: removing heapster dependencies


## 2017-10-05

Agenda:



*   Clayton would like to talk about securing instrumentation endpoints

Notes:



*   How to expose etcd metrics if etcd is generally locked down
    *   Proxy metrics?
    *   Read status/alert endpoints and transform those metrics?


## 2017-09-21

Agenda:



*   Sematext demo (20 minutes)
    *   https://github.com/sematext/sematext-agent-docker
*   Clayton would like to talk about securing instrumentation endpoints
    *   More components adding instrumentation
    *   Some are sensitive (raised in sig-auth)
    *   Would like to identify how we can endorse / suggest instrumentation best practices going forward


## 2017-09-07

Notes:



*   Summary of what has been delivered for 1.8
    *   Metrics api graduation
    *   Metrics server as recommended for serving metrics for cluster
        *   As replacement for heapster
*   Where to host kube-state-metrics containers?
    *   No way to give permissions to maintainers on gcr.io
    *   Keep it in quay.io/coreos for now and provide gcr.io in a best-effort manner
*   


## 2017-08-24

Notes:



*   Want to to enable metric server by default in 1.8


## 2017-08-10

Agenda:



*   Kube-state-metrics in GA
*   Custom Metrics API adapters
*   Master metrics API going beta
*   Kubecon CFP
*   Prometheus output is growing

Notes:



*   Kube-state-metrics 1.0 was released
    *   Compatibility aligned with the client-go version it uses
    *   Load testing was performed and it scales really well
        *   Well below 200MB even for 1000 nodes clusters with 30 pods/node
    *   Default deployment manifest comes with addon resizer configuration
*   Custom metrics API adapters
    *   Solly just tagged first release of Prometheus adapter: [https://github.com/DirectXMan12/k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter)
    *   Stackdriver work in progress
    *   Potentially move to beta for 1.8
*   Historic metrics API not target for 1.8 as it’s a stabilization release
*   Master Metrics API going beta [#50148](https://github.com/kubernetes/kubernetes/issues/50148)
*   Kubecon coming up
    *   CFP closing soon
*   Metrics exposition with Prometheus client library
    *   Google ran into problems where they are exposing metrics about different k8s objects over time. This linearly grows the metric registry and the number of metrics linearly increase
    *   Solution: we have to distinguish between metrics about the running application itself and logical objects like k8s resources. Using custom collectors allows to determine exposed metrics at collection time: [http://godoc.org/github.com/prometheus/client_golang/prometheus#hdr-Custom_Collectors_and_constant_Metrics](http://godoc.org/github.com/prometheus/client_golang/prometheus#hdr-Custom_Collectors_and_constant_Metrics)
    *   This is how exporters like kube-state-metrics usually handle this. Example of the methods that need to be implemented for a custom collector: [https://github.com/kubernetes/kube-state-metrics/blob/master/collectors/service.go#L79-L94](https://github.com/kubernetes/kube-state-metrics/blob/master/collectors/service.go#L79-L94)
    *   It’s advisable to serve metrics on two different ports in this case. One port with metrics about the process itself (e.g. requests it received, open FDs), and another one for the objects its extracting metrics for. That’s because one generally may want to apply different rules for extending the metrics’ label set with external information at ingestion time. (Not doing this has been a long running problem with the kubelet which mixed metrics about itself with cAdvisor metrics in pre-1.7)


## 2017-06-29

Agenda:



*   Limited-scope API for retrieval of historical metrics (i.e. can we build an API that suits VPA and things like idling w/o re-inventing PrometheusQL)
*   Public API for retrieving events from long-term storage.

Notes:



*   Proposal: add API to retrieve historical metrics, e.g. for basic dashboarding data that’s currently collected from Heapster, VPA, idling, …
    *   Can this be folded into the custom metrics API?
    *   Piotr: let’s first start defining idling as a feature before proposing an API required to implement it
    *   SIG autoscaling needs historical API for VPA
*   Idea: Drop-in replacement for event data that can merge local events and events from long-term storage
    *   Solly: should be reasonably doable with existing features, similar to metrics API
    *   Have replacement with higher priority in API aggregator than default handler
    *   Might need to extend existing events API to deal with lack of time bounds in current API
*   Kube State Metrics
    *   Goal in have release for the 1.7 release
    *   Haven’t gotten bugs in a while, seems stable
    *   1 more metric that needs to be reviewed for consistency, but otherwise consistent
    *   Piotr: need scalability test, Google has capacity to test, should have some free cycles next week
        *   Might need sharding, might not (Heapster needs 20GiB for 5k node cluster, which is doable on a cluster that size)
    *   


## 2017-06-22 (cancelled - lack of agenda)


## 2017-06-15


## 2017-06-08 (cancelled - GKE Summit)


## 2017-06-01 (cancelled - OSS Leadership summit)


## 2017-05-25

Agenda:



*   One week to code freeze - checkpoint
*   Metrics server [design doc](https://docs.google.com/a/google.com/document/d/1w6-ZfnA18aKYLJ8DCLBFKlv_1umm74x0I2X4V188kvU/edit?usp=sharing) + implementation [kubernetes-incubator/metrics-server](https://github.com/kubernetes-incubator/metrics-server)

Notes:



*   Status:
*   [WILL NOT BE DONE] Move master metrics API to beta
*   Kube-state-metrics stable release [kube-state-metrics#124](https://github.com/kubernetes/kube-state-metrics/issues/124)
*   [IN PROGRESS] Metrics-server
*   Evaluate custom metrics API state
    *   Hawkular
    *   [IN PROGRESS] Prometheus
    *   [WILL BE DELAYED] Stackdriver
*   [POSTPONED TO 1.8] discuss/propose historical metrics API


## 2017-05-18

Agenda:



*   Metrics server + [master metrics API](https://github.com/kubernetes/metrics/blob/master/pkg/apis/metrics/v1alpha1/types.go) to beta ([original proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md))

Notes:



*   Master Metrics API
    *   Lives in `k8s.io/metrics` (lives in staging now, will get sycned to github.com/kubernetes/metrics), as well as Heapster
    *   Provides basic resource usage metrics for Pods and Nodes, only one data point, no sophisticated query language
    *   Uses:
        *   HPA for the `Resource` source type,
        *   `kubectl top`
        *   scheduler (in future)
        *   Dashboard
    *   Need to decide which approach to use when graduating to beta:
        *   Currently, there is a direct mapping between structure of pod resource requests and resource metrics API
        *   Alternative is to have arbitrary names like “cpu_usage_average_5m”.
    *   [piosz] to create issue
    *   [sross] comment: probably should replace `kind: Pod, metadata.name: &lt;podname>` with `kind: MetricValue, target: &lt;object reference to pod>`.
*   Metrics Server
    *   Minimal server (similar to Heapster), but no storage, limited aggregation, no history, all data in memory
    *   Scrapes from summary API on Kubelets
        *   (summary API may be revamped and moved to beta)
    *   Will be available through aggregator
    *   [piosz] Design doc coming soon (early next week, hopefully)
*   Custom Metrics Prometheus Adapter
    *   Initial rough form at https://github.com/directxman12/k8s-prometheus-adapter


## 2017-05-11 (Logging)

Agenda:



*   Samsung CNCT presenting (non-standard log collection)
    *   Slides: https://docs.google.com/presentation/d/13LUq6TyaWSZTmYKQKPCxVffX-v3_nRJSKYrbas5Jt4M/edit?usp=sharing
*   Owner for ES setup
*   Logging vision (@piosz)
    *   High level vision
    *   Sources
    *   Format


## 2017-05-04

Agenda:



*   Historical API design for VPA, etc

Notes:



*   Kube-state-metrics 0.5 released with new metrics for more resource kinds [https://github.com/kubernetes/kube-state-metrics/releases/tag/v0.5.0](https://github.com/kubernetes/kube-state-metrics/releases/tag/v0.5.0)
*   


## 2017-04-27

Agenda:



*   ~~Metrics server + master metrics API to beta~~
*   [tstclair] Discuss auditing, history, and even offloading… 

Notes:



*   Plumbing to get auditing and event data out of Kubernetes
    *   Some of that could be part of infrastore proposal
    *   System external to Kubernetes
    *   Need to define an idea of what auditing means
    *   Needs standard set of APIs that define how that data can be collected
*   cAdvisor
    *   Breaking changes to metrics
    *   Metrics following wrong format
    *   Need to discuss with sig-node to overhaul the metrics and move cAdvisor out of /metrics of the kubelet
*   [https://github.com/heptio/eventrouter](https://github.com/heptio/eventrouter)
*   


## [Cancelled] 2017-04-20 (Logging)


## 2017-04-13

Agenda



*   1.7 Feature Planning Update
*   [sross] Update on Custom Metrics API boilerplate
*   Kubecon report
*   Prometheus Operator Code

Meeting notes:



*   1.7 planning (copied from 2017-03-02)
    *   Move master metrics API to beta
    *   Kube-state-metrics stable release [kube-state-metrics#124](https://github.com/kubernetes/kube-state-metrics/issues/124)
    *   Metrics-server
    *   Evaluate custom metrics API state
        *   Implementations for testing server, Hawkular, Prometheus, Stackdriver
        *   probably won’t move to beta until 1.8 (we want a release where at least two implementations exist)
    *   discuss/propose historical metrics API
*   Updates on Custom Metrics API boilerplate
    *   PR in progress to switch away from custom patched version of apiserver repository (should make it easier to consume)
*   Prometheus Operator code
    *   Interest in moving towards incubator
    *   Probably also move towards aggregated API server from TPR
        *   [sross] can be pinged for some “getting started tips” on making aggregated API servers


## [Cancelled] 2017-04-06 no meetings week


## [Cancelled] 2017-03-30 Kubecon


## 2017-03-23

Agenda



*   [Solly] Custom metrics API server building
    *   Boilerplate repository: [https://github.com/directxman12/custom-metrics-boilerplate](https://github.com/directxman12/custom-metrics-boilerplate)
    *   Need to implement “pkg/provider”.CustomMetricsProvider and wrap command as in “sample-main.go” and “pkg/sample-cmd”
    *   You can either vendor or fork the repository
    *   Due to some issues, currently there’s no vendor directory in the repo (I’m going to try and fix this soon).  You’ll need most of the same vendor directories as kubernetes, but with the kubernetes from [https://github.com/directxman12/kubernetes/tree/feature/dynamic-resource-routes](https://github.com/directxman12/kubernetes/tree/feature/dynamic-resource-routes) for the mean time (specifically, the k8s.io/apiserver code there)
    *   Feel free to ping @directxman12 on Slack with questions


## 2017-03-16 (Logging)

Agenda



*   Integration with logging on the node ([kubernetes/kubernetes#42718](https://github.com/kubernetes/kubernetes/issues/42718))


## [Cancelled] 2017-03-09


## 2017-03-02

Agenda:



*   Heapster 1.3 release
    *   Code freeze at Friday, March 10th, 6pm PST
        *   [sross] to send out email
*   kube-state-metrics status
    *   Has most of important features
    *   Need to find good balance between adding new metrics and not caching entirety of the API server
    *   Not rushed to release anything in 1.6
    *   Plan to release stable version for 1.7
        *   Need performance test on huge cluster first
*   1.7 planning
    *   Move master metrics API to beta
    *   Kube-state-metrics stable release
    *   Metrics-server
    *   Evaluate custom metrics API state (start moving towards beta?)
        *   Implementations for testing server, Hawkular, Prometheus, Stackdriver
    *   discuss/propose historical metrics API


## 2017-02-23 (Logging)

Agenda:



*   Fluent Bit: Intro & status update by Eduardo Silva ([eduardo@treasure-data.com](mailto:eduardo@treasure-data.com)) \
[https://docs.google.com/presentation/d/1Ovbvk5TsOzVy7wLcyJiBonv6EmET36XSba9zvm_l7tM/edit?usp=sharing](https://docs.google.com/presentation/d/1Ovbvk5TsOzVy7wLcyJiBonv6EmET36XSba9zvm_l7tM/edit?usp=sharing)
*   releases planed for march 1st and may 17th

questions:

- what happens when to many buffers?  

  - no limit.  feature requested

  - buffers can be check pointed to disk for reliability

- can lose messages if log in parse/filter during a bad restart

- metric integration?

  - metrics already being written to a file and a webservice for core systems

  - going to expand this mechanism to be available for all plugins

- journald as input not supported yet

  - feature requested

- memory consumption numbers

  - do not have comparison to old fluentd tracking

  - approximate number show improvement over fluentd

- only talks to kubernetes API right now

Log rotation issues (by vmik@)



*   logrotate acts independently of log aggregators.  this can cause log loss.  change for 1.6 is to move from logrotate to native docker mechanism for rotating.  currently only enabled for GCE.  will be a flag and will be shared widely for how others can enable it. Releated PR [#40634](https://github.com/kubernetes/kubernetes/pull/40634)


## 2017-02-16

Agenda:



*   [Monasca](https://wiki.openstack.org/wiki/Monasca) demo: Quick architecture overview, demo of running in kubernetes environment by Michael Hoppal ([hoppalm@gmail.com](mailto:hoppalm@gmail.com))
*   Announcement: k8s.io/metrics repository (for metrics API type definitions)
*   Graduating master metrics API to beta soon (hopefully in Q2) [https://github.com/kubernetes/heapster/blob/master/metrics/apis/metrics/v1alpha1/types.go](https://github.com/kubernetes/heapster/blob/master/metrics/apis/metrics/v1alpha1/types.go) \
[https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metAnnouncement: k8s.io/metrics repository (for metrics API type definitions)rics-api.md](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-metrics-api.md) 


## [Cancelled] 2017-02-09


## 2017-02-02 (Logging)

Agenda:



*   Log rotation problem [#40634](https://github.com/kubernetes/kubernetes/pull/40634), [#38495](https://github.com/kubernetes/kubernetes/issues/38495)


## 2017-01-26

No agenda, meeting was shortened


## 2017-01-19

No agenda, meeting was shortened


## 2017-01-12 (Logging)

Topics:



*   Logging to files inside containers, possible solution
    *   [Proposed solution](https://goo.gl/IgCSjI)
    *   [PR with the proposed solution](https://github.com/kubernetes/kubernetes.github.io/pull/2141)

Meeting Notes: 



*   [https://docs.google.com/document/d/1kDnQphHJogwGR6U5oX-tpHM6k3mC49TNnFgQV_uBxSk/edit?usp=sharing](https://docs.google.com/document/d/1kDnQphHJogwGR6U5oX-tpHM6k3mC49TNnFgQV_uBxSk/edit?usp=sharing)


## 2017-01-05 (Monitoring)

Notes:



*   Maintenance model of Heapster:
    *   Rotation for maintainers for responding to issues; exact people TBD
    *   Sink owners to discuss pull requests in meetings, give ok-to-merge
        *   Rieman sink potentially deprecated
    *   Full document specifying to be submitted to Heapster, will be under docs/ directory
*   Custom metrics API proposal: [https://github.com/kubernetes/community/pull/152](https://github.com/kubernetes/community/pull/152)
*   Instrumentation guidelines: [https://github.com/kubernetes/community/pull/195](https://github.com/kubernetes/community/pull/195)
    *   What are deprecation rules? Potentially carry new and old metric for one release
*   Deprecating builtin-cAdvisor from kubelet
    *   Kubelet exposes a lot of (all) cAdvisor metrics but only needs a small subset itself
    *   Moving towards a slimmer API only exposing data cAdvisor needs
    *   Users will have to run their own cAdvisor for monitoring purposes
    *   [https://github.com/kubernetes/community/pull/252](https://github.com/kubernetes/community/pull/252)
