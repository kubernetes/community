## Agenda (2018-12-13)



*   Metrics overhaul KEP discussion - in person in Seattle at KubeCon
    *   Discussed what needs to be done, priority and what is already in-flight
    *   Decided to keep any non-conformant metric labels for v1.14 but clearly state they are deprecated and will be removed in v1.15 (or v1.16 if we get any pushback)
    *   Add histograms wherever there are summaries
    *   Make summary metrics opt-in with a kubelet flag
        *   Not a breaking change, can be done after v1.14 target
    *   Update KEP status to implementable
        *   Thanks @ehashman
    *   Create plan to add dev, operator and user docs to metrics
        *   I don’t remember all of the context on this, @directmanx12 this was something you brought up, can you fill it in a bit?
    *   Discussed how to change a single global metrics registry to something that gets passed in and can be replaced with a no-op registry if desired
        *   This pattern has been implemented in client-go as part of the controller runtime implementation with the logger object


## Agenda (2018-11-29)



*   Demo on tracing Sam Naser
    *   KEP here: [https://github.com/kubernetes/enhancements/pull/650](https://github.com/kubernetes/enhancements/pull/650)
    *   Next steps:
        *   create tracing feature proposal
        *   house mutating webhook for adding trace to an object in kubernetes-sigs
        *   use annotations for not to not go through an immediate API review


## Agenda 2018-11-15



*   [https://github.com/kubernetes/community/pull/2909/](https://github.com/kubernetes/community/pull/2909/) 
*   Current state of tracing in Kubernetes
    *   [https://docs.google.com/document/d/1cqdw7JfHSovl1E-FoH4rTpI32Xt0saZvdKv6q6-v4uc/edit?usp=sharing](https://docs.google.com/document/d/1cqdw7JfHSovl1E-FoH4rTpI32Xt0saZvdKv6q6-v4uc/edit?usp=sharing) &lt;- link to public design document
    *   [https://github.com/Monkeyanator/kubernetes/pulls](https://github.com/Monkeyanator/kubernetes/pulls)


## Agenda 2018-11-1



*   Elasticsearch logging addon - @coffeepac
    *   Additional OWNER
    *   New image repo
*   Metrics overhaul KEP opened and targeted for 1.14


## Agenda 2018-10-18



*   Review initial KEP draft: [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/TMUTDP4cLQw](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/TMUTDP4cLQw)
    *   Introduce promtool in order to check for metric best practices
    *   Open pull request to add KEP to repository
*   Bug [https://github.com/kubernetes/kubernetes/issues/68918](https://github.com/kubernetes/kubernetes/issues/68918)
    *   Introduce heuristic for detecting cardinality explosions in releases
*   Community demo: Filebeat hints based autodiscover (exekias / [carlos@elastic.co](mailto:carlos@elastic.co))
*   Kube-state-metrics performance optimization update


## Agenda 2018-10-04



*   Canceled due to having no agenda points to discuss.


## Agenda 2018-09-06



*   Charter merged
*   We need to write a KEP (Kubernetes Enhancement Proposal) for metrics overhaul, because it affects lots of users
    *   Will there be a draft and feedback? - Yes, just like design proposals
    *   Follow up: setup google doc to flesh out initial proposal for this KEP and start collaborating on it and review it together in the next meeting
        *   Done: [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/TMUTDP4cLQw](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/TMUTDP4cLQw)
*   SIG Instrumentation has to use the Kubernetes organizations for now
*   Kube-state-metrics performance optimization
    *   Second PR up for early feedback, refactoring collectors logic to cache metrics instead of Kubernetes objects

        [https://github.com/kubernetes/kube-state-metrics/pull/534](https://github.com/kubernetes/kube-state-metrics/pull/534)

    *   Can there be a docker image be provided with these changes? - Yes, mxinden will provide a personal one


## Agenda 2018-08-23:



*   Charter document [https://github.com/kubernetes/community/pull/2266](https://github.com/kubernetes/community/pull/2266)
*   Kube-state-metrics performance optimization
    *   [https://github.com/kubernetes/kube-state-metrics/issues/498](https://github.com/kubernetes/kube-state-metrics/issues/498)
*   Kubernetes metrics overhaul
    *   [https://github.com/kubernetes/kubernetes/pull/67476#issuecomment-413785762](https://github.com/kubernetes/kubernetes/pull/67476#issuecomment-413785762)
    *   Consider renaming cAdvisor labels [https://github.com/kubernetes/kubernetes/issues/66790](https://github.com/kubernetes/kubernetes/issues/66790)
    *   General consensus is: yes we should do this at once, probably aiming for 1.13
    *   We need to figure out whether we need a KEP or feature.
        *   Researched answer: Asked a couple of people and unanimously was told a KEP would be more appropriate and give this the appropriate visibility. 
*   [sross] metrics-server status/release prep
    *   Preparing a new release of a rather major cleanup of metrics-server
    *   Soon alpha version
        *   Probably a stable version soon afterwards
*   [sross] Moving stuff to kubernetes-sigs
    *   Can we have our own org?
        *   Researched answer: Orgs per sig is currently not manageable so currently everything goes into kubernetes-sigs.


## Agenda 2018-07-26:



*   [Proposed] - Review of [feature idea](https://docs.google.com/document/d/1PjbaImDrSs3qj1oqu46lSChGgJ6ka_N5AuQv0HVkBbI/edit#heading=h.te3fbxigdo0t) - CRD for “Draining” namespaces to a `syslog:// `endpoint
*   Charter: [https://github.com/kubernetes/community/pull/2266](https://github.com/kubernetes/community/pull/2266)
    *   Needs more review
*   Sig update in community meeting
    *   Heapster deprecated
        *   Deprecation timeline ([https://github.com/kubernetes/heapster/blob/master/docs/deprecation.md](https://github.com/kubernetes/heapster/blob/master/docs/deprecation.md)) -- next step is setup removal in 1.12, completely deprecated as of 1.13
    *   Node metrics reworking
    *   Metrics-server refactoring (not yet merged, calling for feedback) - [https://github.com/kubernetes-incubator/metrics-server/pull/65](https://github.com/kubernetes-incubator/metrics-server/pull/65)
    *   k8s-prometheus-adapter advanced config merged
    *   A number of third party service involving e2e tests have been put behind a feature flag in the test infrastructure (to improve flaking tests from sig-instrumentation)


## Agenda 2018-06-28:



*   Charter: [https://github.com/kubernetes/community/pull/2266](https://github.com/kubernetes/community/pull/2266)
    *   Needs more review
*   Non googlers to push images to gcr.io
*   Third party e2e test results:  [https://github.com/kubernetes/test-infra/blob/master/docs/contributing-test-results.md](https://github.com/kubernetes/test-infra/blob/master/docs/contributing-test-results.md)
    *   This is how we will recommend that third party tools submit their test results for inclusion in testgrid


## 2018-06-14:



*   Charter: [https://github.com/kubernetes/community/pull/2266](https://github.com/kubernetes/community/pull/2266)
    *   Needs more review
*   How to enforce instrumentation guidelines, when there are existing violations? [https://github.com/kubernetes/kubernetes/pull/64481#discussion_r192527282](https://github.com/kubernetes/kubernetes/pull/64481#discussion_r192527282)
    *   Do a review of all metrics in a certain release, make public in release notes
    *   Then introduce stricter workflow for introducing metrics
    *   No metric stability currently, but we also shouldn’t frustrate users by breaking often
*   
*   Testing PRs, need review from @piosz
    *   [https://github.com/kubernetes/test-infra/pull/8451](https://github.com/kubernetes/test-infra/pull/8451)
    *   [https://github.com/kubernetes/kubernetes/pull/64564](https://github.com/kubernetes/kubernetes/pull/64564)
    *   None needed for log interface, [already exists](https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/log_path_test.go).


## 2018-05-31:



*   Sig-instrumentation charter
*   Testing notes
    *   Sig-instrumentation breaking e2e owned tests
        *   [https://docs.google.com/spreadsheets/d/1OirZorG4bbwlEkxAW-2qdp0dXDZrVKtDBDFC0Nq226s/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1OirZorG4bbwlEkxAW-2qdp0dXDZrVKtDBDFC0Nq226s/edit?usp=sharing)
    *   Check if SIg-node has any logging interface tests, if not write one
    *   @piosz move the top level testgrid google-gke-stackdriver somewhere else


## 2018-06-14



*   How to submit test results as a third party
    *   Prefer to find sig-testing doc, will try and prepare a minimal sig-inst doc if needed


## 2018-05-31



*   Charter PR or doc should be coming tomorrow (6/1)
    *   Charter defaults align with what we already do


## 2018-05-17



*   KubeCon recap
    *   Medium well attended and lots of good questions
    *   Very good audience
    *   Lengthen one session to include a compressed intro and the entire deep dive and not one shorter topic on each
    *   Energetic custom metric adapter interest from vendors (at least 3 new)
    *   Public link for videos forthcoming
*   Heapster is now deprecated
    *   Thanks @directxman12
    *   This is official, feature requests closed
    *   Make sure this makes it to the v1.11 release notes
    *   What are the next steps to graduate kube-state-metrics out of alpha
    *   Action item: @piosz to find current dashboard maintainers and determine what the current state of the dashboard is, 
        *   Historical API, does dashboard want to access data directly
*   Sig-instrumentation-kubernetes group
    *   What is the policy for allowing projects
    *   Need a charter
        *   Includes official processes for a sig, structure of sig, etc.
        *   @brancz to fill out template prior to next meeting ~~@coffeepac to add template to this~~
            *   [README](https://github.com/kubernetes/community/blob/9565401b5702a3deffb0e5d9f2999e8d12bbc9a2/committee-steering/governance/README.md) for what the process is, includes link to template
*   3rd party/vendor test comments
    *   What should be marked as ‘e2e’
        *   @coffeepac to generate list of e2e tests we own, if a reasonable number share a spreadsheet to #sig-instrumentation slack
    *   How to label 3rd party/vendor tests for viewing 
        *   @coffeepac to write up how to do this


## 2018-04-19



*   “Ignoring flakes: sig-instrumentation” [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/cbbzkMXSMaw](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/cbbzkMXSMaw)
    *   If it is not kube code, then we should not have tests on them - Solly
    *   Given we have one kind of e2e tests we are not fixing in time, we shouldn’t add more (Regarding last meetings discussion) - Frederic
    *   What is the Kubernetes code being tested here (it looks like “can Stackdriver scrape Kube logs”)?  If it’s “can thing X connect to Kubernetes”, then it probably shouldn’t be in Kubernetes e2e tests - Solly
        *   Can we have a way for external projects to test integrations with Kube?  Might want to reach out to SIG testing - Frederic
    *   @coffeepac to ask sig-instrumentation about what is the desired way to handle 3rd party/vendor integrations for e2e testing
*   Prometheus cluster-monitoring addon [https://github.com/kubernetes/kubernetes/pull/62195#issuecomment-382778622](https://github.com/kubernetes/kubernetes/pull/62195#issuecomment-382778622)
    *   Addons should not belong in the Kubernetes repository - Frederic/Solly
    *   Cluster-monitoring seems like a lot larger scope than discussed e2e setup from last meeting - Frederic
    *   Should have gone into a sig-instrumentation specific repo - @coffeepac
    *   Contrib repo recommends Prometheus Operator - Frederic
*   Kubernetes Node Monitoring - Solly
    *   Draft: [https://docs.google.com/document/d/1_CdNWIjPBqVDMvu82aJICQsSCbh2BR-y9a8uXjQm4TI/edit?usp=sharing](https://docs.google.com/document/d/1_CdNWIjPBqVDMvu82aJICQsSCbh2BR-y9a8uXjQm4TI/edit?usp=sharing)
*   Kube-pod-exporter POC demo


## 2018-04-05



*   [piosz] kube-up is in a bit of shaky position
    *   Deprecate InfluxDB kube-up in 1.11, remove in 1.12
    *   [sross] deprecate Influx e2e tests as well
    *   [piosz] deploy Prometheus as well
        *   [sross] it’s not needed for e2e tests, so I’d lean against
        *   [piosz] want a “real” test for custom metrics, with an actual monitoring solution, Prometheus would be good for that, non-blocking
        *   [sross] just need to be careful to avoid maintenance issues with Influx in the future
*   [brancz] have PoC for pod exporter, blocked on getting crio up with supports for stats endpoint, share it hopefully next meeting


## 2018-03-22



*   Aligning cAdvisor labels with official Kubernetes instrumentation guidelines (possibly related to [https://github.com/kubernetes/kubernetes/issues/45043](https://github.com/kubernetes/kubernetes/issues/45043))
    *   TODO(brancz): Share POC of pod-exporter once CRI implementation with stats endpoints is available
    *   Further: brancz and directxman12 will take lead on stable metrics for pods in Kubernetes
        *   Need to figure out pod-level cgroups, other data endpoints (device metrics, etc)
*   Road to heapster deprecation/phase out? Should we put a deprecation note at the top of the heapster readme?
    *   Mark Heapster as being in maintenance mode
        *   No new features
        *   No new sinks
        *   Only bugfixes
    *   Come up with timeline for deprecation
        *   No support
        *   No new bugfixes
    *   Need better docs on metrics-server setup
    *   Docs missing?
*   Metrics Server Cleanup
    *   Backport fixes from Heapster (IPV6, etc)
    *   Remove unneeded code
    *   Abstract out serving interface to serve resource metrics API from other sources (e.g. directly from monitoring pipeline), implement testing tools, etc
    *   [directxman12] to publish a bunch of the refactor code
*   Proxying counter metrics in Prometheus client
    *   Pain point of prometheus client library when writing exporters, where counter semantics cannot necessarily applied with available abstractions by the golang Prometheus library
        *   Interim solution: Implement necessary semantics with “lower level” Prometheus “const” metrics
        *   Long term: Learn from the interim solution in order to provide re-usable abstraction to Prometheus client-library


## 2018-02-22



*   Kubecon sig-instrumentation deep dives sessions
*   Best practices for exposing kubelet health checks?
    *   Probably health checks has to be exposed on different endpoint (not a _/metrics_).
    *   AI(Solly): Include details in issue [https://github.com/kubernetes/kubernetes/issues/58235](https://github.com/kubernetes/kubernetes/issues/58235)
        *   Commented on [https://github.com/kubernetes/kubernetes/pull/58827](https://github.com/kubernetes/kubernetes/pull/58827)
    *   We will need to write our own exporter of metrics
*   External Metrics API/HPA changes
    *   [https://github.com/kubernetes/community/pull/1801](https://github.com/kubernetes/community/pull/1801)
    *   [https://github.com/kubernetes/community/pull/1802](https://github.com/kubernetes/community/pull/1802)


## 2018-02-08



*   Metrics-server cleanup continued - needs to be taken care of
    *   [https://github.com/kubernetes-incubator/metrics-server/issues/37](https://github.com/kubernetes-incubator/metrics-server/issues/37)
*   External Metrics API - a proposal will be written up
*   cAdvisor, core/resource metrics and CRI? What’s our stand, everything consumed via CRI? (RE: [https://github.com/kubernetes/kubernetes/issues/55905](https://github.com/kubernetes/kubernetes/issues/55905)) - Solly will revise his proposal and then share
*   Log file separation? [https://github.com/kubernetes/kubernetes/issues/58638#issuecomment-359979485](https://github.com/kubernetes/kubernetes/issues/58638#issuecomment-359979485)
*   Kubernetes workload benchmarker
    *   [https://docs.google.com/document/d/1hYOzX8jBHceuXgDVzlasveMqetpKtnq433aNMj1_x0o/edit](https://docs.google.com/document/d/1hYOzX8jBHceuXgDVzlasveMqetpKtnq433aNMj1_x0o/edit)
    *   [https://github.com/ZJU-SEL/capstan/tree/prometheus](https://github.com/ZJU-SEL/capstan/tree/prometheus) 

     -     Failing e2e test:  https://github.com/kubernetes/kubernetes/issues/58837


## 2018-01-25



*   Intro and Deep Dive Sessions in Copenhagen
*   The road to heapster deprecation?
*   State of metrics-server
    *   Are we intending to keep sinks?
    *   Cleanups necessary (many heapster things still lurking around)
    *   PVC stats? [https://github.com/kubernetes/features/issues/497](https://github.com/kubernetes/features/issues/497)
*   Prometheus-k8s-adapter

Notes:



*   brancz@ is interested in making Intro for KubeCon (and DeepDive as well). Piotr  can also prepare something for Intro.
*   Heapster deprecation:
    *   kubectl top switched to metric-server in 1.10.
    *   Google is need heapster for exporting metrics to Stackdriver. Their team is going to support it.
    *   We can remove Metrics API from the Heapster. Dashboard may still rely on Model API of heapster.
*   Metric-server:
    *   We don’t want to keep sinks in the codebase
    *   Need well defined interface between metric-server and kubelet. Summary API is not ideal right now.
    *   It’s not clear if PVC should be represented as separate entity or as a part of Pod stats.


## 2018-01-11



*   2018 Vision
    *   Move all sig-instrumentation projects to new home (cluster addons, contrib, standlone apps, etc) - @coffeepac to start planning
    *   Make build/release of projects be publically viewable/triggerable
        *   Find out where kubernetes/kubernetes is and start moving sig-inst work to mainline process - @coffeepac to find starting issue
    *   Historical metrics API - @brancz follow up on VPA design doc to find out involvement needed from sig-instrumentation
    *   Kubernetes Pod exporter - @brancz share prototype and figure out what the plan of CRI stats is going to be going forward
*   kube-state-metrics release
