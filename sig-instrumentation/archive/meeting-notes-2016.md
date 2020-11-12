## 2016-12-15

Agenda:



*   Demo by Datadog (rescheduled)
*   Kubernetes Metric Conventions: [https://docs.google.com/document/d/1YVs02Li6QFCg8Th2Wa4z1u2NBlQHDp2dj3EdAt6uskE/edit#](https://docs.google.com/document/d/1YVs02Li6QFCg8Th2Wa4z1u2NBlQHDp2dj3EdAt6uskE/edit#)
*   Resource metrics API: looking towards beta
    *   [https://docs.google.com/document/d/1t0G7OS6OP9qPndkkNROCu0pF3-vkDmzonmT-6gEWcx0/edit?ts=5852bda8](https://docs.google.com/document/d/1t0G7OS6OP9qPndkkNROCu0pF3-vkDmzonmT-6gEWcx0/edit?ts=5852bda8)

Notes:



*   Put metric convention document somewhere visible for reference
    *   [https://github.com/kubernetes/community/tree/master/contributors/devel](https://github.com/kubernetes/community/tree/master/contributors/devel)
*   Resource metrics API should be moved towards beta
    *   To be finalized after holiday break
    *   Working towards beta in 1.7
*   Custom metrics API:
    *   [https://github.com/kubernetes/community/pull/152/files](https://github.com/kubernetes/community/pull/152/files)


## 2016-12-08

**Warning: This meeting will be about logging. If you are not interested please skip.**

Agenda



*   Restart LogDir proposal ([https://github.com/kubernetes/kubernetes/pull/13010](https://github.com/kubernetes/kubernetes/pull/13010))
*   Alternative [https://github.com/kubernetes/kubernetes/pull/33111](https://github.com/kubernetes/kubernetes/pull/33111)

Meeting notes:  [https://gist.github.com/leahnp/463501f6dfe39f6f21ea5d3ebcb787d7](https://gist.github.com/leahnp/463501f6dfe39f6f21ea5d3ebcb787d7)


## 2016-12-01


### Agenda



*   Heapster needs your help
    *   [sross] Need to come up with map of sinks to maintainers
        *   Maybe consider dropping sinks without mainters
    *   [sross] need statement of plans for Heapster
        *   [sross] putting into maintenance mode, what does maintenance mode entail, should we continue accepting sinks?
        *   [piosz] to write something up and send out
*   [mwringe] what is plan for timeline for monitoring pipeline work
    *   [piosz] plan is starting work Q2 2017, unless anyone else can help
        *   [piosz] major missing component is discovery summarizer
        *   [sross] we (Red Hat) are willing to help out in this area


## [Cancelled] 2016-11-24: Thanksgiving in US


## [Cancelled] 2016-11-17: no meeting week


## [Cancelled] 2016-11-10: Kubecon


## [Cancelled] 2016-11-03


## 2016-10-27


### Agenda



*   F2f meeting about monitoring in Seattle during KubeCon (on Monday Nov 7th)


## 2016-10-20

**Warning: This meeting will be about logging. If you are not interested please skip.**


### Agenda



*   f2f meeting about logging in Seattle during KubeCon (probably on Monday Nov 7th)
    *   There is going to be a kubernetes dev summit (Nov 10th) meeting for logging
*   Group administrivia:  frequency?  Length? Topics?
*   Current state of logging in Kubernetes
*   What’s going on with logging?

Notes

Developers Summit - 45 minute unconference topic on the future of logging

 - moderated by Vishnu and Patrick

 - open to anyone who is attending the Kubernetes Developers Conference

Discussion of Face to Face meeting - Piotr and Patrick to sync up offline

Frequency:  every three weeks, going to skip next week/push back one week next meeting is during KubeCon Developers Summit.  

 - There will be an announcement for exactly when the next meeting is

Logging Discussion Topics:

  - logging volumes (proposal started by David Cowden -[ https://docs.google.com/document/d/1K2hh7nQ9glYzGE-5J7oKBB7oK3S_MKqwCISXZK-sB2Q/edit#](https://docs.google.com/document/d/1K2hh7nQ9glYzGE-5J7oKBB7oK3S_MKqwCISXZK-sB2Q/edit#))

  - hot loop logging and verbosity for scalability issues. 

     - how to detect spammy instances

     - how to not let this wreck the cluster

  - general dissatisfaction with the logging facility

  - structured logging kubernetes wide for consistent consumption

  - application log type detection

    - what metadata do we need to carry through a logging pipeline to id a source system (e.g. mysql, user application)

    - what do logging vendors need supplied to aid in this

Current logging pipelines

  - fluentd direct to GCP or ES

  - fluentd to kafka to fluentd to ES

Action Items

 - Piotr & Patrick to determine f2f details

 - Try and get logging vendors to join the SIG


## [Cancelled] 2016-10-13


## 2016-10-06


### Agenda



*   No response from sig api machinery (moving to next meeting)
*   Continue discussion on monitoring architecture
    *   Agreed to versioned, well-defined API
    *   Rest API vs. Query Language
    *   A webhook model was suggested for the APIs (like Auth in Kube today)
        *   [sross] has concerns over discoverability of webhooks
        *   Webhook vs API server is largely an implementation question
        *   will decide on discovery vs webhook for consumption once we get the API design in place
    *   [sross] will propose an API design for the custom metrics API and historical metrics API
*   Discuss [roadmap](https://docs.google.com/document/d/1j6uHkU8m6GvElNKCJdBN8KrejkUzVbp2l0zTyeSxrl8/edit)
    *   Discussed briefly, please go read afterwards
    *   [sross] to lead push on custom metrics design/implementation for 1.5
    *   1.5 API features will be mainly implemented in terms of Heapster
*   looking forward for one-click install of 3rd party monitoring (possibly Prometheus, but as an out of the box, one command setup; possible choices for deployment: helm, kpm)
*   Logging discussion feasibility conversation (ie: is this a reasonable location for having discussions about logging)
    *   This may be a reasonable place for logging discussions, if we explicitly note which meetings will discuss logging (and/or when logging will be discussed)
    *   May also just want to create a separate SIG
    *   [decarr] mentioned CRI discussion on logging and metrics
        *   Outcome was that we should sync with SIG node on that, but it should probably stay more in SIG node


## 2016-09-29


### Agenda



*   Discuss [Kubernetes monitoring architecture proposal ](https://docs.google.com/document/d/1z7R44MUz_5gRLwsVH0S9rOy8W5naM9XE5NrbeGIqO2k/edit#)
    *   


### Notes



*   Main metrics pipeline used by Kubernetes components
*   Separate operator-defined monitoring pipeline for user-exposed monitoring
    *   Generally collects core metrics redundantly/independently
*   Should it be possible to implement the core metrics pipeline on top of the custom monitoring system
    *   As long as one implements the core metrics API, one could swap it out for scheduler etc.
*   Upstream Kubernetes would test against the stable core pipeline
*   Replaceable != Pluggable – the entire thing gets replaced in a custom scenario
*   Master Metrics API part of main Kubernetes API
    *   Should further APIs like for historic metrics also be in that group?
    *   Discussion for sig-apimachinery
*   Should Infrastore be part of core Kubernetes
    *   Provides historic time series data about the system
    *   Would require implementing a subset of a TSDB
    *   Not an implemented component, just an API

     

*   What are core metrics exactly?
    *   CPU, memory, disk
    *   What about network and ingress?
    *   Resource estimator would not read from master metrics API but collect information itself (e.g. from kubelet)


## 2016-09-22


### Agenda



*   Mission statement: [https://docs.google.com/document/d/15Q47xbYTGHEZ-wVULGSgOSD5Kq-OehJj-MEChVH1kqk/edit?usp=sharing](https://docs.google.com/document/d/15Q47xbYTGHEZ-wVULGSgOSD5Kq-OehJj-MEChVH1kqk/edit?usp=sharing)
*   Kubesnap demo


### Notes



*   Kubesnap demo by Andrzej Kuriata, Intel ([slides](https://docs.google.com/presentation/d/1fgGik1nq-yEN7Y2dRIQWTjb7r5HEWaG9paDCdvzE_IA/edit?usp=sharing)):
    *   Daemon set in k8s
    *   Integration with Heapster
*   Mission Statement:
    *   Enough people to coordinate, but small enough to be focused
    *   List of people actually doing development/design in the scope of this sig
    *   Scratchpad before a meeting to set discussions of features before meeting
    *   Sig autoscaling discussed and committed to features/metrics in previous meetings
    *   A plan for an api for 1.5?


## 2016-09-15


### Agenda



*   Presentation by Eric Lemoine (Mirantis): monitoring Kubernetes with [Snap](http://snap-telemetry.io/) and [Hindsight](https://github.com/trink/hindsight). [Slides](https://docs.google.com/presentation/d/1XWM0UmuYdcP_VsbKg6yiSDb6TR1JmouHdZAnLelBWXg/edit?usp=sharing)
*   Meeting frequency
*   Ownership SIG instrumentation vs SIG autoscaling
*   [Discuss how to export pod labels for cAdvisor metrics (see kubernetes/kubernetes#32326)](https://github.com/trink/hindsight)


### Notes



*   Meeting frequency - defer until ownership clarified
*   Ownership SIG autoscaling vs instrumentation
    *   Triggering issue: [https://github.com/kubernetes/kubernetes/issues/31784](https://github.com/kubernetes/kubernetes/issues/31784)
    *   HPA is consumer of Master Metrics API (also kubectl top, scheduler, UI)
    *   Could potentially be relevant to monitoring as well
    *   Make distinction between metrics used by the cluster and metrics about the cluster
    *   One SIG lead cares about system level metrics, one about the external/monitoring side. Good setup for the SIG to handle both areas?
    *   Follow up with mission statement on the mailing list taking these things into account
*   Kube-state-metrics v0.2.0 was released with many more metrics:
    *   [https://github.com/kubernetes/kube-state-metrics#metrics](https://github.com/kubernetes/kube-state-metrics#metrics)


## 2016-09-08


### Agenda



*   Sylvain Boily showing their monitoring solution


### Notes



*   Demo by Sylvain on their monitoring setup using InfluxDB+Grafana+Kapacitor
    *   Scraping metrics from Heapster, Eventer, and apiserver
*   Separation apiserver vs kube-state-metrics
    *   The apiserver exposes metrics on /metrics about the running state of the apiserver process
        *   How man requests came in from clients? What was their latency?
        *   Outbound latency to the etcd cluster?
    *   Kube-state-metrics aims to provide metrics on logical state of the entire Kubernetes cluster
        *   How many deployments exist?
        *   How many restarts did pod X have?
        *   How many available/desired pods does a deployment have?
        *   How much capacity does node X have?
*   Separation Heapster vs [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/commits/master)
    *   Heapster holds metrics about characteristics about things running on Kubernetes, used by other system components.
    *   Currently Heapster asks the Kubelet for cAdvisor metrics vs. kube-state-metrics collecting information from the apiserver
*   Should eventer information be consolidated with kube-state-metrics?
*   Should we look into the creation of a monitoring namespace / service for all other namespace to use? 
*   Should monitoring be available out of the box with a k8s installation when done in a private datacenter ?


## 2016-09-01


### Agenda



*   State of [Kubernetes monitoring at Soundcloud](https://drive.google.com/file/d/0B_br6xk3Iws3aGZ5NkFMMDRqRjhvM1p1RWZXbVF2aVhiWGZz/view?usp=sharing) (Matthias Rampke)
*   Future of [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics)
*   Application metric separation in cAdvisor ([https://github.com/google/cadvisor/issues/1420](https://github.com/google/cadvisor/issues/1420))
*   ...


### Notes



*   Matthias Rampke giving an intro to their Kubernetes monitoring setup
    *   Currently running Prometheus generally outside of Kubernetes
        *   Easy migration path from previous infrastructure
    *   Still using DNS as service discovery instead of Kubernetes API
    *   Sharded Prometheus servers by team for application monitoring
    *   Severe lack of metrics around Kubernetes cluster state itself
    *   Long-term vision (1yr): all services and their dependencies running inside of Kubernetes
        *   Prometheus part of that via a standard configuration
        *   Easy to spin up monitoring new components
*   People using Heapster as it gives them all metrics in one component
*   Something as easy to deploy as Heapster would be useful
*   Three sets of metrics
    *   Those useful only for monitoring (e.g. number of pods)
    *   Metrics for auto-scaling (CPU, custom app metrics)
    *   Those that fit both
*   Make Prometheus a first-class citizen/best practice for exposing custom auto-scaling metrics?
*   Overlap between auto-scaling and monitoring metrics seems generally fine
    *   storing them twice is okay, auto-scaling metrics are way fewer
*   Kube-state-metrics
    *   Keep it as a playground or fold it into controller manager?
    *   


## 2016-08-25


### Notes



*   CoreOS would like to see
    *   more instrumentation as insight into cluster
    *   Remove orthogonal features in for example cadvisor
*   RedHat
    *   Good out-of-the-box solution for cluster observability, component interaction
    *   Collaboration with sig-autoscaling
*   SoundCloud:
    *   Prometheus originated at SoundCloud
    *   Bare metal kubernetes setup: separation of monitoring
    *   Separation of heapster and overall kubernetes architecture
    *   How are people instrumenting around kubernetes
*   Mirantis:
    *   Scalability of monitoring solutions
    *   More metadata from kubelet “stats” API: labels are missing for example
    *   Also interested in “Separation of heapster and overall kubernetes architecture” (from SoundCloud)
    *   Extended insight into OpenStack & Kubernetes
    *   During our scalability tests we want to measure k8s behaviour in some set of defined metrics
*   Intel:
    *   Integration of snap into kubernetes
    *   Help deliver monitoring goals

Where should guides for flavors of monitoring live?

→ ad hoc currently, not all the same

→ best practices in the community

Where are we and where do we want to do? → Google doc will be setup

Next meeting: Discuss google doc & Matthias from SoundCloud will give an insight of how they are using Prometheus to monitor Kubernetes and its pain points.

Next time will use Zoom as hangout limit is 10 participants.

Kubernetes monitoring architecture (~~requires joining [https://groups.google.com/forum/#!forum/kubernetes-sig-node](https://groups.google.com/forum/#!forum/kubernetes-sig-node)~~): [https://docs.google.com/document/d/1HMvhhtV3Xow85iZdowJ7GMsryU6pvjOzruqcJYY9MMI/edit?ts=57b0eec1#heading=h.gav7ymlujqys](https://docs.google.com/document/d/1HMvhhtV3Xow85iZdowJ7GMsryU6pvjOzruqcJYY9MMI/edit?ts=57b0eec1#heading=h.gav7ymlujqys)

