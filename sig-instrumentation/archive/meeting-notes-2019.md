## Agenda (2019-12-12) 



    *   [hase1128]Discuss how to proceed KEP(#1348) \
[https://github.com/kubernetes/enhancements/pull/1348](https://github.com/kubernetes/enhancements/pull/1348) \
And I would like to consult about the following slack comment \
[https://kubernetes.slack.com/archives/C20HH14P7/p1574840290078600](https://kubernetes.slack.com/archives/C20HH14P7/p1574840290078600)
    *   [serathius] Discuss high level design of Structured logging and collect feedback [https://github.com/kubernetes/enhancements/pull/1367](https://github.com/kubernetes/enhancements/pull/1367)
    *   [serathius] Ask for contributions to Metrics Server [https://github.com/kubernetes-sigs/metrics-server/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22](https://github.com/kubernetes-sigs/metrics-server/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)
    *   [Tom Kerkhove] Improve Kubernetes events to be CloudEvents 1.0 compatible \
[https://github.com/kubernetes/kubernetes/issues/85544](https://github.com/kubernetes/kubernetes/issues/85544) \



## Agenda (2019-11-28) 



*   Cancelled.


## Agenda (2019-11-18) [KubeCon NA]


    @Contributor Summit / Technical Discussion Room


    Strategy Session Notes



*   Goals
    *   Review Open Keps
    *   Discuss and set next of goals
        *   Road map with releases
        *   
*   Open Keps
    *   [https://github.com/kubernetes/enhancements/pull/650](https://github.com/kubernetes/enhancements/pull/650) 
    *   [https://github.com/kubernetes/enhancements/pull/1343](https://github.com/kubernetes/enhancements/pull/1343) [Metric Stability - Beta]
    *   [https://github.com/kubernetes/enhancements/pull/1367](https://github.com/kubernetes/enhancements/pull/1367) [Structured Logging]
*   Metric Stability Framework
    *   Provides API for specifying stable metrics
    *   Beta:  must use stability metric framework
    *   GA: allow for end users to selectively turn off individual metrics
    *   Future: rules for allowing metrics to graduate
        *   Allows for continuous cleanup of metrics
    *   @logicalhan: to file HelpWanted/GoodFirstIssue for this KEP
*   Structured Logging
    *   Original decision was to use glog because it was simple
    *   Migrate all logging to an API
        *   Json as format
        *   Standard metadata enforced via schema
    *   Migration is a massive job
        *   Double writing but only one will be used (other is a NoOp), user flag decides which to use
    *   Reviews wanted
    *   Timeline:  
        *   KEP reviewed/approved by 1.18 (mid-january 2020)
        *   Initial implementation in 1.18
        *   Alpha auto migration framework by end of 2020
        *   Beta milestone by end of 2020
        *   GA 2021 at some point
*   Using Tracing for Kubernetes Object Lifecycle
    *   Intend to use OpenTelemetry when it's ready, currently PoC’d in OpenCensus
    *   [https://github.com/open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go)
    *   Changes needed (high level)
        *   All component boundaries modified to include context for propagation
            *   This should happen in client-go no matter what
    *   Requires discussion with how it will interact with structured logging
    *   Timeline:
        *   KEP almost approved
        *   Initial implementation in 1.18
*   API graduation:
    *   Metrics, beta -> GA
    *   Custom Metrics, beta -> GA
    *   External Metrics, beta -> GA
        *   All three require coordination with sig-autoscaling
        *   Must implement `watch`, this is not supported by any known current monitoring implementation
*   Metrics Server graduation
    *   Default implementation for Metrics API
    *   Could use developers
*   Kube State Metrics
    *   v2.0
    *   Could use developers
*   Unused tests
    *   Remove unused e2e tests, these are mostly behind feature flags already
*   Moving anything we own in /cluster to bit bucket or kubernetes-sigs org
*   @piotr and @brancz to create document as a 2 year planning document with clear goals
    *   Expected by end meeting on Dec 12th


## Agenda (2019-11-14) [Cancelled]



*   Skipping in preparation for KubeCon


## Agenda (2019-10-31)



*   [RainbowMango] 1.17 plans again
    *   I think following tasks should be in 1.17, list them with dependency
    *   I think we should hide deprecated metrics with stability framework in 1.17
        *   [https://github.com/kubernetes/kubernetes/pull/83836](https://github.com/kubernetes/kubernetes/pull/83836)
        *   [https://github.com/kubernetes/kubernetes/pull/83837](https://github.com/kubernetes/kubernetes/pull/83837)
        *   [https://github.com/kubernetes/kubernetes/pull/83838](https://github.com/kubernetes/kubernetes/pull/83838)
        *   [https://github.com/kubernetes/kubernetes/pull/83839](https://github.com/kubernetes/kubernetes/pull/83839)
        *   [https://github.com/kubernetes/kubernetes/pull/83841](https://github.com/kubernetes/kubernetes/pull/83841)
    *   But they rely on [https://github.com/kubernetes/kubernetes/pull/84135](https://github.com/kubernetes/kubernetes/pull/84135)
    *   Migration task(as well as remove prometheus reference) almost done except custom collector. And this rely on [https://github.com/kubernetes/kubernetes/pull/83062](https://github.com/kubernetes/kubernetes/pull/83062)
    *   The last one I think it should be in 1.17 is the flag for kube-binaries
        *   The first one is: [https://github.com/kubernetes/kubernetes/pull/84292](https://github.com/kubernetes/kubernetes/pull/84292)
*   KubeCon


## Agenda (2019-10-17)



*   [alejandrox1] quick hello from release team 
    *   [https://github.com/kubernetes/sig-release/tree/master/releases/release-1.17#enhancements-freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.17#enhancements-freeze)
    *   All enhancements wishing to be included in 1.17 must have
        *   A KEP in an implementable state
        *   Including Testing Plans
        *   Including Graduation Criteria
        *   An open issue in the 1.17 Milestone
*   [piosz] 1.17 plans
    *   Finish metrics kep - metrics stability implementation (especially wrt hidden metrics) deferred to 1.18
    *   OOM kill metrics
    *   AI(brancz, piosz): keep the stuff trackable
*   Long-term plans (2020) - discuss at Contributor Summit during KubeCon?
    *   Han/Elana’s intro to instrumentation talk scheduled at the same time as sig-inst intro session
    *   Elana to email about fixing the schedule conflict
        *   Done: SIG Instrumentation intro moved forward to 4:25pm
    *   Deep dive session - metrics stability discussion
    *   Tariq to get back to sig-inst on format of Contributor Summit SIG Meet and Greet


## Agenda (2019-10-03)



*   Cancelled


## Agenda (2019-09-05)



*   Status of 1.16 SIG Instrumentation feature implementation (Han, Frederic)
*   Roadmap for 1.17 release
    *   Remove direct dependency on Prometheus (did I get that right?)
*   Improving OOMKill metrics (gauge -> counter?) 


## Agenda (2019-07-25)



*   [Discuss metric validation/verification KEP](https://github.com/kubernetes/enhancements/pull/1169)
*   PR for removing deprecated cadvisor labels is live: [https://github.com/kubernetes/kubernetes/pull/80376](https://github.com/kubernetes/kubernetes/pull/80376)
    *   Need SIG Node + SIG Testing approval (due to test case update)


## Agenda (2019-07-11)



*   Update on feedback from sig-cloud provider (logicalhan) and sig-node (ehashman)
    *   Sig cloud provider
        *   They’re refactoring cloud providers out of tree, into separate binaries
    *   Sig-node
        *   Informed them of the *_name label removal on cadvisor metrics, hoping to get this feature in for 1.16 (see [https://github.com/kubernetes/kubernetes/pull/69099](https://github.com/kubernetes/kubernetes/pull/69099) for label duplication, landed in 1.14)
        *   They said they would get back to us at the sig-node meeting next week with approval
        *   Once we get sig-node’s blessing, we should probably mention this at the next community meeting


## Agenda (2019-06-27)



*   [Discuss metrics migration (control-plane stability) KEP](https://github.com/kubernetes/enhancements/pull/1093/)


## Agenda (2019-06-13) ~~(2019-05-30 _deferred until next meeting_~~)



*   [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/XbElxDtww0Y](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/XbElxDtww0Y)
*   [https://github.com/kubernetes/kubernetes/pull/76496](https://github.com/kubernetes/kubernetes/pull/76496)
    *   Consensus was that we should include component owners as reviewers on KEPs which will affect their binaries and verify that they read and are aware of upcoming changes.
*   Initial discussion for metrics migration (control-plane stability)
    *   Issue for discussion (migration of shared metrics, i.e. client-go), how can we do component based migration (i.e. per metrics endpoint) if we have metrics which are shared between migrated and non-migrated components?
    *   [Link to draft KEP](https://github.com/kubernetes/enhancements/pull/1093/)
*   Initial discussion for metrics conformance (control-plane stability)
    *   [https://github.com/kubernetes/enhancements/pull/1089](https://github.com/kubernetes/enhancements/pull/1089)


## Agenda (2019-05-02)



*   Continued OpenCensus/OpenTracing discussion from last meeting
*   Structured Logging ([https://github.com/kubernetes/kubernetes/issues/69825](https://github.com/kubernetes/kubernetes/issues/69825)) & ([https://groups.google.com/forum/#!topic/kubernetes-sig-architecture/wCWiWf3Juzs](https://groups.google.com/forum/#!topic/kubernetes-sig-architecture/wCWiWf3Juzs))
    *   [https://github.com/go-commons/commons/issues/1](https://github.com/go-commons/commons/issues/1) (standard logger interface discussion after dotGo 2017)


## Agenda (2019-04-18)



*   Watch API
*   Review [control-plane metric stability KEP](https://github.com/kubernetes/enhancements/pull/946).
*   Start discussion around OpenCensus/OpenTracing (especially since [we are introducing some OpenCensus stuff in container-runtime](https://github.com/kubernetes-sigs/controller-runtime/pull/368), [corresponding groups discussion](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/n0Fq2Dg5Ixs))


## Agenda (2019-04-04)



*   Should we standardize success/failure label values? (coming out of: [https://github.com/kubernetes/kubernetes/issues/75839](https://github.com/kubernetes/kubernetes/issues/75839))
*   Discuss metrics stability proposal ([https://docs.google.com/document/d/1CcbfC-M8CHDfq1rMAOtW0-LKHvermyUiV6BMXXYiqoM/edit#heading=h.r5x1ipcsw2c8](https://docs.google.com/document/d/1CcbfC-M8CHDfq1rMAOtW0-LKHvermyUiV6BMXXYiqoM/edit#heading=h.r5x1ipcsw2c8))
*   Watch API


## Agenda (2019-03-21)



*   Mail thread: unbounded metric labels [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/7wbr6eQ58b0 ](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/7wbr6eQ58b0)
*   Discuss possible modifications to existing kubelet probe metrics 
*   Metric stability KEP (Han)
*   ehashman is giving a talk on Kubernetes monitoring at SREcon next week: [https://www.usenix.org/conference/srecon19americas/presentation/hashman](https://www.usenix.org/conference/srecon19americas/presentation/hashman) 


## Agenda (2019-03-07)



*   Mail thread: Metric deprecation: [https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/XbElxDtww0Y](https://groups.google.com/forum/#!topic/kubernetes-sig-instrumentation/XbElxDtww0Y) 
*   Plan to cut Kube-state-metrics v1.6.0
    *   PR to cut release candidate is out: [https://github.com/kubernetes/kube-state-metrics/pull/702](https://github.com/kubernetes/kube-state-metrics/pull/702) 
*   Shoutout to [https://github.com/tariq1890](https://github.com/tariq1890) for the awesome work on kube-state-metrics


## Agenda (2019-02-21)



*   Cancelled due to no agenda


## Agenda (2019-02-07)



*   [Kubelet Resource Metrics Endpoint KEP](https://github.com/kubernetes/enhancements/pull/726) review (dashpole@) [Slides](https://docs.google.com/presentation/d/14zM8S7Ftymo3OabGc208EIjLCXpDheA8yjVV7hWUr2M/edit?usp=sharing)
*   Our dev docs need a review [https://github.com/kubernetes/community/issues/3097](https://github.com/kubernetes/community/issues/3097)
    *   Metrics Overhaul Review
    *   Outstanding PRs:
        *   [https://github.com/kubernetes/kubernetes/pull/69099](https://github.com/kubernetes/kubernetes/pull/69099)
        *   [https://github.com/kubernetes/kubernetes/pull/72470](https://github.com/kubernetes/kubernetes/pull/72470)
        *   [https://github.com/kubernetes/kubernetes/pull/73366](https://github.com/kubernetes/kubernetes/pull/73366) (this one might need a discussion)
*   Fluentd-elasticsearch addon image repository move PR  [https://github.com/kubernetes/kubernetes/pull/73819](https://github.com/kubernetes/kubernetes/pull/73819) (@coffeepac)


## Agenda (2019-01-24)



*   Cancelled due to no agenda


## Agenda (2019-01-10)



*   Demo on Prometheus Adapter replacing Metrics Server for resource metrics
*   [Metrics overhaul KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/0031-kubernetes-metrics-overhaul.md) review
    *   Only one that seems to be contentious is: [https://github.com/kubernetes/kubernetes/pull/67476](https://github.com/kubernetes/kubernetes/pull/67476)
*   Status on kube-state-metrics
    *   V1.5.0 stable release PR is out, please review! [https://github.com/kubernetes/kube-state-metrics/pull/629](https://github.com/kubernetes/kube-state-metrics/pull/629)
*   Update on Pod Termination Reason Counter discussion with sig-node (Brian)
    *   [https://github.com/kubernetes/kubernetes/issues/69676#issuecomment-442391695](https://github.com/kubernetes/kubernetes/issues/69676#issuecomment-442391695) 
