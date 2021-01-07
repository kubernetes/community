## _Triage (2020-12-16)_

Attendees:



*   Ehashman
*   Logicalhan
*   Brancz
*   Dashpole
*   lilic
*   Serathius
*   Metalmatze
*   Afrouz Mashayekhi


## _Agenda (2020-12-10)_

Issues:



*   Last regular SIG meeting of 2020!
*   Continue discussion on metrics naming policy
    *   Dashpole: PR to update the metric naming guideline.
*   What’s going on with stable metrics? Lots of bugs against website for the list being empty e.g. [https://github.com/kubernetes/website/issues/24915](https://github.com/kubernetes/website/issues/24915) 
    *   In freeze now, nothing we can do
    *   SLO support is a concern, reference [https://github.com/kubernetes/community/blob/master/sig-scalability/slos/api_call_latency.md](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/api_call_latency.md)
    *   Etcd metrics on large numbers of entries is a concern. What is “large” is different per customer
*   [lili]: Adding new subproject/or part of kube-state-metrics itself but on another /metrics endpoint, that exposes metrics that do cannot be added to kube-state-metrics, e.g. things that do not map 1:1 to the k8s API [https://github.com/kubernetes/kube-state-metrics/issues/1302#issuecomment-734237720](https://github.com/kubernetes/kube-state-metrics/issues/1302#issuecomment-734237720) 
*   [lili] kube-state-metrics v2.0.0-beta was cut [https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-beta](https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-beta)

Attendees:



*   


## _Triage (2020-12-02)_

Issues:



*   Test out SIG Node triage tool: [https://docs.google.com/document/d/1JOXKBDgXmQzz8YQSYa7XYcfVteM79iMtvId1aQXC1e8/edit](https://docs.google.com/document/d/1JOXKBDgXmQzz8YQSYa7XYcfVteM79iMtvId1aQXC1e8/edit) 
    *   Kick this to the next triage meeting, probably in the new year

Attendees:



*   Ehashman
*   Logicalhan
*   Joadavis


## _(CANCELLED) Agenda (2020-11-25)_

Cancelled due to US Thanksgiving.


## _(CANCELLED) Triage (2020-11-18)_

Cancelled due to KubeCon.


## _Agenda (2020-11-12)_

Issues:



*   [dashpole] FYI Node proxy conformance testing vs removal: [https://github.com/kubernetes/kubernetes/issues/95930](https://github.com/kubernetes/kubernetes/issues/95930).  I think prometheus integration depends on this.
    *   Please add any feedback for groups that may be using these in the issue linked here
*   [ehashman] Developer documentation updates [https://github.com/kubernetes/community/issues/5233](https://github.com/kubernetes/community/issues/5233) 
    *   Wrappers for metrics framework - needs to documented [https://github.com/kubernetes/community/issues/5080](https://github.com/kubernetes/community/issues/5080) 
    *   Global metrics registry - Lili
    *   Tracing - dashpole
    *   More detailed naming policies for metrics
    *   Any needed updates for existing docs - logicalhan
*   [serathius] Strategy and direction for metrics server [https://github.com/kubernetes-sigs/metrics-server/issues/627](https://github.com/kubernetes-sigs/metrics-server/issues/627) 
    *   Reach out to contribex to run a user survey?
*   Discussion about metric naming
    *   Scheduler: [https://github.com/kubernetes/kubernetes/pull/94866](https://github.com/kubernetes/kubernetes/pull/94866) 
    *   Kubelet: [https://github.com/kubernetes/kubernetes/pull/95839](https://github.com/kubernetes/kubernetes/pull/95839) 
    *   [brancz] Suggestion: there are two types of metrics:
        *   For k8s resources: node/pod/container…
        *   For k8s components: apiserver/…
        *   For resource related metrics, it would be better to use the recommended naming e.g. `kube_pod_…`
    *   [ehashman] worry that the proposal would confuse the people about where to find the metrics definition. 
        *   kube_* used for KSM only currently
        *   kube_* doesn’t say where a metric is coming from
        *   We don’t have any documentation for metrics other than KSM docs
        *   Operators won’t easily be able to understand these metrics
*   [ehashman] Announcements/Reminders:
    *   Code freeze: TODAY Nov. 12 for 1.20
    *   Doc freeze: Nov. 30
    *   Upcoming meetings: Nov/Dec cancellations

Attendees:



*   Ehashman
*   Logicalhan
*   Dashpole
*   Brancz
*   Serathius
*   lilic
*   Leszek
*   Erain
*   Joadavis
*   Kakkoyun
*   Damien Grisonnet


## _Triage (2020-11-04)_

Attendees:



*   Dashpole
*   Logicalhan
*   Ehashman
*   Joadavis
*   erain


## _Agenda (2020-10-29)_

No agenda items, cancelling.


## _Triage (2020-10-21)_

Attendees:



*   Dashpole
*   Kakkoyun
*   erain


## _Agenda (2020-10-15)_

Issues:



*   [dashpole] Metadata standards: [See this comment](https://github.com/kubernetes/enhancements/pull/1458#issuecomment-705842455)
    *   We have two competing metadata standards: [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/resource/semantic_conventions/k8s.md) and [Kubernetes](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/instrumentation.md).
    *   Should we change?  Ask them to change?
    *   Why did they choose ‘.’ as the namespace delineator? 
*   [logicalhan] adding context to metrics interfaces so that we can absorb label values via context, rather than tedious, error-prone argument plumbing 
    *   conclusion: we achieved community consensus on adding context to metrics interfaces in component-base, will be backwards compatible with existing usages so we are free to move forward with this w/o a KEP
*   [lili] [https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-alpha.1](https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-alpha.1) new pre-release towards 2.0 release

Attendees:



*   Brancz
*   Dashpole
*   Lilic
*   Logicalhan
*   kakkoyun
*   erain


## _Triage (2020-10-07)_

Attendees:



*   Ehashman
*   Dashpole
*   Brancz
*   Utkarsh
*   erain


## _Agenda (2020-10-01)_

Issues:



*   [ehashman] Revisiting KEPs targeted for 1.20 before enhancements freeze
    *   KEP freeze after this weekend
    *   7 KEPs currently open: [https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Finstrumentation](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Finstrumentation) 
    *   Will track 5 into 1.20, assuming they are merged as implementable by freeze (Oct. 6)
*   [lili] v1 klog compatibility in v1.19
    *   Was not announced well enough?
    *   Due for a community update this month - should include that
*   [SergeyKanzhelev] Help with some PRs review: [https://docs.google.com/document/d/13TZ2CVXIp9BRUbkyUZMn-GDFRTJL-EC2hnQv0lqh_Vs/edit](https://docs.google.com/document/d/13TZ2CVXIp9BRUbkyUZMn-GDFRTJL-EC2hnQv0lqh_Vs/edit) 
    *   -> triage meeting
    *   Event deduplication PRs - not clear what’s the rules on what’s goes into the event and what not
*   [Added container name in the back-off event log · Issue #93636 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/pull/93636)
*   [fix duplicate FailedMount events when using multiple volumes · Issue #88369 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/pull/88369)
*   [dashpole] Demo of [systemd_exporter](https://github.com/dashpole/systemd_exporter)
    *   Povilas’ version of systemd_exporter: [https://github.com/povilasv/systemd_exporter](https://github.com/povilasv/systemd_exporter) 

Attendees:



*   ehashman
*   joadavis
*   Dashpole
*   brancz
*   Leszek Jakubowski
*   Serathius
*   SergeyKanzhelev
*   Metalmatze
*   Steph Rifai
*   Damien Grissonet
*   Clayton


## _Triage (2020-09-23)_

Attendees:



*   ehashman
*   dashpole
*   erain


## _Agenda (2020-09-17)_

Issues:



*   [ehashman] KEPs targeted for 1.20
    *   Enhancements freeze Oct. 6, code freeze Nov. 12
    *   [https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Finstrumentation](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Finstrumentation) 
    *   Reviewed and updated all open KEPs
*   [ehashman] Follow-up on e2e test flakiness/state of logging addons
    *   Metrics grabber test had a race condition, fixed by dashpole with a retry [https://github.com/kubernetes/kubernetes/issues/93688](https://github.com/kubernetes/kubernetes/issues/93688) 
    *   [https://github.com/kubernetes/kubernetes/issues/92731](https://github.com/kubernetes/kubernetes/issues/92731) might be because of rename of pod; erain@ to (learn to) run the e2e test and confirm. 
    *   [https://github.com/kubernetes/kubernetes/issues/93480](https://github.com/kubernetes/kubernetes/issues/93480) Reassign elasticsearch flake to @brancz
*   [lili] Migration of [https://github.com/DirectXMan12/k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter) 
    *   [https://github.com/kubernetes/org/issues/2182](https://github.com/kubernetes/org/issues/2182) 
    *   In progress
    *   Anyone want to join the OWNERS file? @brancz offers to help with reviews
*   [lili] kube-state-metrics 2.0.0-alpha pre release -** breaking changes** - please help testing! [https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-alpha](https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-alpha)

Attendees:



*   ehashman
*   brancz
*   dashpole
*   lilic
*   metalmatze
*   Marek
*   Sergey Kranzhelev
*   erain
*   44past4


## _Triage (2020-09-09)_

Attendees:



*   ehashman
*   brancz
*   dashpole
*   kakkoyun
*   erain


## _Agenda (2020-09-03)_

Issues:



*   [chelseychen] - Overview for [New Event API GA Graduation](https://github.com/kubernetes/enhancements/pull/1662)
*   [ehashman] - Follow up on last meeting’s topic (e2e tests)
    *   Dashpole - looked a little, saw a few isolated instances but still nothing big
    *   Serathius - will have an update by last week
*   [ehashman] Does SIG instrumentation want to use feature tracking similar to SIG Node? (from Lauri Apple)
    *   [https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit?ts=5f3da0c9#](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit?ts=5f3da0c9#) 
    *   -> defer to triage meeting next week

Attendees:



*   ehashman
*   dashpole
*   joadavis
*   Keikhara
*   Leszek (makdaam)
*   erain
*   ...


## _Agenda (2020-08-20)_

Issues:



*   [logicalhan] Chairs met last week and have started a weekly triage session on Wednesdays
    *   On the SIG calendar, anyone can attend
*   [ehashman] instrumentation e2e test roundup
    *   See e.g. [https://github.com/kubernetes/kubernetes/issues/93688](https://github.com/kubernetes/kubernetes/issues/93688) (Metrics grabber)
        *   ~~dashpole said there’s only one flake because the controller didn’t come up; probably not much to fix~~
    *   [https://github.com/kubernetes/kubernetes/issues/92731](https://github.com/kubernetes/kubernetes/issues/92731) (Stackdriver logging tests)
    *   [https://github.com/kubernetes/kubernetes/issues/93480](https://github.com/kubernetes/kubernetes/issues/93480) (Elasticsearch logging tests)
    *   Need to determine some sort of strategy for unmaintained e2e tests
    *   Elasticsearch addon does not seem to be in great shape, similar Stackdriver addon was removed
    *   Marek should have the best context on this
        *   **Action:** find someone to assist with Stackdriver tests
    *   ehashman suggests we should just remove or find a maintainer for the logging tests
*   Ended early [:46] - short meeting because of KubeCon

Attendees:



*   ehashman
*   dashpole
*   logicalhan
*   Serathius
*   erain
*   ...


## _Agenda (2020-08-06)_

Issues:



*   Lauri Apple - triage workflow discussion
    *   [https://github.com/google/triage-party](https://github.com/google/triage-party)
*   [Solly Ross, Yuchen Zhou] - demo a new instrumentation tool (promq) 
    *   Temporary link: [https://github.com/logicalhan/instrumentation-tools](https://github.com/logicalhan/instrumentation-tools)
*   Frederic - k8s-prometheus-adapter consensus on adoption

Attendees:



*   dashpole
*   brancz
*   kakkoyun
*   lilic
*   metalmatze
*   ehashman
*   logicalhan
*   joadavis
*   Leszek


## _Agenda (2020-07-09)_


## Issues:



*   [Frederic] Adopt [https://github.com/DirectXMan12/k8s-prometheus-adapter](https://github.com/DirectXMan12/k8s-prometheus-adapter) as sig-instrumentation project
    *   Need to include a clear positioning statement, how it relates to metrics-server
    *   Makes more sense to be in SIG Instrumentation than to try to push to Prometheus community

[Matthias] Update on  [https://github.com/kubernetes/kubernetes/issues/91536](https://github.com/kubernetes/kubernetes/issues/91536)?

Attendees:



*   lilic
*   brancz
*   Leszek
*   Kemal Akkoyun
*   David Ashpole
*   joadavis
*   Damien Grisonnet
*   Mark Siarkowicz
*   Matthias Loibl


## Agenda (2020-06-25)


## Issues:



*   [immutableT] [KEP-1753](https://github.com/kubernetes/enhancements/pull/1754): Kubernetes system components logs sanitization
*   [serathius] Code freeze moved back by 2 weeks to July 9

Attendees:



*   Alextc
*   Brancz
*   dashpole
*   ehashman
*   marek
*   pawel
*   damien grisonnet
*   joadavis
*   lili
*   matthias loibl
*   patrick rhomberg
*   logicalhan


## Agenda (2020-06-11)

Issues:



*   [brancz] [https://github.com/kubernetes/kubernetes/issues/58638](https://github.com/kubernetes/kubernetes/issues/58638)
*   [logicalhan] [API call latencies SLO](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/api_call_latency.md) (as we discussed in the previous meeting)
*   Should probably also do a sweep through PR backlog
*   [immutableT] KEP-1753: Kubernetes system components logs sanitization \


Attendees:


## _ \
_Agenda (2020-05-28)

Attendees:



*   logicalhan
*   ehashman
*   brancz
*   dashpole
*   rainbowmango
*   serathius
*   metalmatze
*   (this was all i could recollect from memory) 

Agenda Issues:



*   [logicalhan] setting up SIG repos
*   [brancz] Revisit Metric Stability KEP graduation criteria - [https://github.com/kubernetes/kubernetes/issues/91536](https://github.com/kubernetes/kubernetes/issues/91536)
    *   Kubernetes-mixin already heavily depends on that metric being stable: \
[https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/config.libsonnet#L3-L18](https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/config.libsonnet#L3-L18) \
[https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/alerts/kube_apiserver.libsonnet](https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/alerts/kube_apiserver.libsonnet) \
[https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/rules/kube_apiserver.libsonnet](https://github.com/kubernetes-monitoring/kubernetes-mixin/blob/master/rules/kube_apiserver.libsonnet)
*   [ehashman] (housekeeping) wrap-up of Metrics Overhaul KEP? [https://github.com/kubernetes/enhancements/issues/1206](https://github.com/kubernetes/enhancements/issues/1206) 
    *   Circle back with sig-testing/sig-release to close the loop on this
    *   **Action:** ehashman to finish KEP
*   [logicalhan] 


## Agenda (2020-05-14)

Attendees:



*   ehashman
*   logicalhan
*   dashpole
*   brancz
*   Kakkoyun
*   dgrisonnet
*   Joadavis
*   serathius

Agenda Issues:



*   ~~[logicalhan] setting up SIG repos ~~(deferred)
*   [44past4] ([New KEP](https://github.com/kubernetes/enhancements/pull/1754)) [Dynamic logs sanitization](https://github.com/kubernetes/enhancements/issues/1753).
    *   Sensitive data like tokens, keys, passwords can leak in logs for controller-manager, other components
*   [logicalhan] [Dynamic cardinality enforcement](https://github.com/kubernetes/enhancements/pull/1692)
*   [logicalhan] fixing api request total metric before returning to stable
    *   Han to file an issue detailing the problems - [https://github.com/kubernetes/kubernetes/issues/91536](https://github.com/kubernetes/kubernetes/issues/91536)


## Agenda (2020-04-30)

Attendees:



*   dashpole
*   brancz
*   ehashman
*   logicalhan
*   serathius
*   liggitt
*   joadavis
*   lilic
*   smarterclayton
*   44past4
*   dgrisonnet

Agenda issues:



*   [serathius] FYI Enhancement Freeze moved back 2 weeks [kubernetes/sig-release#1065](https://github.com/kubernetes/sig-release/pull/1065)
*   [liggitt] [KEP-1693](https://github.com/kubernetes/enhancements/pull/1694): Mechanism for warning API clients about deprecated API use
    *   proposed deprecated API use metric is relevant to sig-instrumentation
    *   AI: Jordan follow up with logicalhan/brancz about new metric vs adding to existing request count metric - [https://groups.google.com/g/kubernetes-sig-instrumentation/c/BACdXH4LscY](https://groups.google.com/g/kubernetes-sig-instrumentation/c/BACdXH4LscY) 
*   [brancz/smarterclayton/lilic] (New KEP) Accurate Pod resource request/limit metric reporting as seen by scheduler. ([Email](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!topic/kubernetes-sig-instrumentation/Klg_DJxFd54), [Google Doc](https://docs.google.com/document/d/13mjommm_wbwqKtisThmIa33TRa-gfIHJicxQ1S-TOlw/edit#heading=h.scvi1ournpue), [KSM issue](https://github.com/kubernetes/kube-state-metrics/issues/1095))
    *   AI: smarterclayton to submit a draft KEP for more comment
*   [44past4] (New KEP) [Dynamic logs sanitization](https://github.com/kubernetes/enhancements/issues/1753).
    *   Sensitive data like tokens, keys, passwords can leak in logs for controller-manager, other components


## Agenda (2020-04-16)

Attendees:



*   ehashman
*   logicalhan
*   dashpole
*   brancz
*   Joadavis
*   Serathius
*   lilic

Agenda Issues:



*   [logicalhan] WIP Cardinality KEP ([https://github.com/kubernetes/enhancements/pull/1692](https://github.com/kubernetes/enhancements/pull/1692))
*   Should we plan on bringing stability to GA this release? I would like to start marking metrics as STABLE, especially the ones that we can’t modify without really breaking people.
    *   Few remaining things from metrics stability KEP: CI checks, dynamic disabling of metrics
    *   Need to identify metrics proposed to be made STABLE
    *   Need to contact subproject owners to propose metrics to be made STABLE - someone should draft a communication?
    *   Is there already some thinking about having a maturity index per metric?  Not much other than what is in source control.  Metric stability framework should provide some of that, though not GA yet. KEP - [https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/20190404-kubernetes-control-plane-metrics-stability.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/20190404-kubernetes-control-plane-metrics-stability.md)  Framework will/could have alpha/beta/stable per metric.
*   Bug scrub? 
    *   [logicalhan] (related - [https://github.com/kubernetes/kubernetes/issues/89788](https://github.com/kubernetes/kubernetes/issues/89788))
*   Fixing the calendar invite [logicalhan]
    *   Fix meeting notes link
    *   Invite was created by a coreos account that doesn’t exist any more


## Agenda (2020-04-02)

Attendees:



*   Logicalhan
*   dashpole
*   Akonarde
*   kakkoyun
*   joadavis
*   ehashman
*   Lili
*   Serathius
*   hase1128
*   44past4

Agenda Issues:



*   [logicalhan] Cardinality discussion
    *   Issues: 
        *   [https://github.com/kubernetes/kubernetes/issues/76302](https://github.com/kubernetes/kubernetes/issues/76302)
        *   [https://github.com/kubernetes/kubernetes/issues/89378](https://github.com/kubernetes/kubernetes/issues/89378) �

        *   [https://github.com/kubernetes/kubernetes/issues/89377 s
](https://github.com/kubernetes/kubernetes/issues/89377)
    *   Can/should we introduce any mechanisms to protect against unbounded cardinality?
    *   Decision: logicalhan will work on a KEP to address this.  Should be small
*   Bug scrub


## Agenda (2020-03-19) 	

Attendees:



*   logicalhan
*   ehashman
*   Brancz
*   olivierlemasle
*   dashpole
*   Jpepin
*   Metalmatze
*   kakkoyun
*   Joadavis
*   Leszek Jakubowski
*    lilic
*   Maximillianbrain1
*   serathius

Agenda Issues:



*   Welcome new tech lead
    *   David Ashpole
*   [ehashman] Fixed terms and/or term limits for SIG leads?
    *   Proposed: 1 year terms
    *   ACTION: Han to inquire about setting up lead term limits
*   [ehashman] Community update 2020-03-19 done! [Slides](https://docs.google.com/presentation/d/1tWWWsnZZPcoMAYj60jL2541cNA1GyLCR4i-EQzaI9y8/edit#)
*   [ehashman] Community meeting -> SIG repos need to be moved out of kubernetes/*
    *   KSM is affected, may break pretty badly (uses k8s.io imports)
    *   ACTION: Elana to look for info on repo migration [done]
    *   Should be an issue in k/org for migration (have done this for metrics-server from k-incubating -> k-sigs)
        *   Issue: [https://github.com/kubernetes/steering/issues/136](https://github.com/kubernetes/steering/issues/136) 
*   [metalmatze] I created a PR with an example APIServer SLO implementation (following [SIG Scalability guidelines](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/api_call_latency.md)). Might be interesting for folks and comments welcome (not strictly for this SIG): [https://github.com/kubernetes-monitoring/kubernetes-mixin/pull/382](https://github.com/kubernetes-monitoring/kubernetes-mixin/pull/382) 


## Agenda (2020-03-05) 



*   Welcome new chairs
    *   logicalhan
    *   ehashman
*   New chairs to set up the vote for Tech Leads


## Agenda (2020-02-20) 



    *   [pszczesniak] Proposal - @serathius to replace @piosz as SIG Chair (see [details](https://groups.google.com/g/kubernetes-sig-instrumentation/c/Rod5jxVluT0))
    *   Suggestion (ehashman): governance structure change from 2 chairs to up to 2 chairs, up to 2 TLs.
    *   Suggestion (logicalhan): terms (mandatory re-election on a regular interval) for chairs/TLs.
    *   Action items:
        *   Please ensure you are in the [member tracking yml](https://github.com/kubernetes/org/blob/master/config/kubernetes/sig-instrumentation/teams.yaml), and send out email to mailing list
        *   Call for nominations for new chair/TL positions on mailing list (nominees must confirm their acceptance)
        *   Secret vote (CIVS?)


## Agenda (2020-02-06) 



    *   [serathius] Update on structured logging


## Agenda (2020-01-23) 



    *   [mark peek] Cloud events introduction
    *   [hase1128]  [I considered how to proceed Request-ID again, so I want feedback.](https://kubernetes.slack.com/archives/C20HH14P7/p1579742164003900) 


## Agenda (2020-01-09) 



    *   [dashpole] Discuss how [tracing](https://github.com/kubernetes/enhancements/pull/650), [structured logging](https://github.com/kubernetes/enhancements/pull/1367), and [RequestID](https://github.com/kubernetes/enhancements/pull/1348) proposals interact. \
[Slides with potential solutions](https://docs.google.com/presentation/d/1NmMKaZwIhdgiktSt1OyAktWPdO5D83br0ueVWkDO8uQ/edit?usp=sharing). \
Background: Kubecon tracing [presentation](https://www.youtube.com/watch?v=lEACvRW6T_U)
    *   [hase1128]  [Update Request-ID KEP and add basic interaction idea](https://github.com/hase1128/enhancements/blob/kep-request-id/keps/sig-auth/20191101-add-request-id-to-k8s-logs.md#proposal)
    *   [brancz] resource metrics API working set defined but metrics-server exposes RSS because VPA [https://github.com/kubernetes-sigs/metrics-server/pull/391](https://github.com/kubernetes-sigs/metrics-server/pull/391)
