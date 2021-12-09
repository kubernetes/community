## Triage (2021-12-30)



* CANCELED FOR YEAR END SHUTDOWN


## Agenda (2021-12-23)



* CANCELED FOR YEAR END SHUTDOWN


## Triage (2021-12-16)

Cancelled due to conflict with [Contributor Celebration](https://www.kubernetes.dev/events/kcc2021/)


## [NEXT] Agenda (2021-12-09)

Agenda:



* Announcements and AI follow-up
    * Contributor celebration is 16-18 Dec.
    * Registration & info: [https://www.kubernetes.dev/events/kcc2021/](https://www.kubernetes.dev/events/kcc2021/)
* [Leads] Selecting a new TL
    * Action: written email on Frederic’s departure
* [ehashman] 1.24 KEP discussion
    * [dgrisonnet] custom and external metrics API graduation to GA
    * [dgrisonnet] graduation of the [cardinality enforcement](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2305-metrics-cardinality-enforcement) KEP to Beta
    * [pohly] new KEP from WG structured logging: [contextual logging](https://github.com/kubernetes/enhancements/pull/3078) alpha
* [shuaichen] metrics.k8s.io SLO (performance) and pagination
    * Ping ehashman offline for help with this
* [fromani] (can be postponed, 1.24 or beyond) klog: towards per-flow verbosity	
    * Looking for previous history/attempts (if any) and design yay/nay
    * Probably deserves a full KEP, will write depending on the above bullet point
        * POC/usecase description [here](https://github.com/kubernetes-sigs/scheduler-plugins/pull/289) (caveat: security implications not addressed)


## [NEXT] Triage (2021-12-02)

Attendees:

- ehashman
- Damien Grisonnet
- David Ashpole
- Haoyu Sun
- Jan Fajerski
- Kevin Wiesmuller
- MZ
- Shuai
- yongfeng du


## Agenda (2021-11-25)

Agenda:



* CANCELED FOR US THANKSGIVING
    * Leads won’t be available


## Triage (2021-11-18)

Attendees:



* Ehashman
* Logicalhan
* Dgrisonnet
* Ben Luddy

[https://github.com/prometheus-operator/kube-prometheus/pull/1499](https://github.com/prometheus-operator/kube-prometheus/pull/1499)


## Agenda (2021-11-11)

Agenda:



* AI review, announcements
    * Code freeze upcoming: Nov. 16
    * Docs placeholder deadline is Nov. 23 - a bit earlier than usual
    * Frederic will be stepping down as TL
        * We do not have a successor selected and we will be considering interested maintainers
    * Sustainable staffing requests
        * [https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit#heading=h.849b7ydpl7ip](https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit#heading=h.849b7ydpl7ip)
* [shuaichen] KEP: Promote k8s resource metrics API to GA, link to [doc](https://docs.google.com/document/d/e/2PACX-1vS10SozEaPnTziMJtxLS9fCc0HNMBouqK1JEjLo737F7P7B3aIAaIMBg-kM1KQ7OdeXgqGKA8Ligjbf/pub)
    * May have slipped through the cracks from API Machinery
    * There was an issue where pod resources were not necessarily represented correctly in metrics and that caused issues with the API, since they can only represent resources on a per-container basis, but pods also have pod-level cgroup controls; there is a difference between the per-container and cgroups values
    * Resource metrics API account for container usage instead of pod usage: [https://github.com/kubernetes/kubernetes/issues/102051](https://github.com/kubernetes/kubernetes/issues/102051)
* [ehashman] cadvisor registry after refactor
    * [https://kubernetes.slack.com/archives/C20HH14P7/p1635962175015900](https://kubernetes.slack.com/archives/C20HH14P7/p1635962175015900)
    * [https://github.com/kubernetes/kubernetes/pull/106334/files#diff-c3e7f724d1b0dcee40df80716ed57d90d2649710150fa92bf9822bdad35e0429](https://github.com/kubernetes/kubernetes/pull/106334/files#diff-c3e7f724d1b0dcee40df80716ed57d90d2649710150fa92bf9822bdad35e0429)
* [dashpole] [https://github.com/kubernetes/kubernetes/pull/106007](https://github.com/kubernetes/kubernetes/pull/106007)
    * Seems inline with what we currently have, but the docs are very out of date
    * Feels a bit like vendor advertising, maybe we should just remove this readme?
    * **Action:** ehashman to remove the outdated README info and any vendor references.
        * [https://github.com/kubernetes/kubernetes/pull/106360](https://github.com/kubernetes/kubernetes/pull/106360)
* [fromani] (can be postponed, 1.24 or beyond) klog: towards per-flow verbosity	
    * No time left! Will discuss on slack!
        * Looking for previous history/attempts (if any) and design yay/nay
        * Probably deserves a full KEP, will write depending on the above bullet point
            * POC/usecase description [here](https://github.com/kubernetes-sigs/scheduler-plugins/pull/289) (caveat: security implications not addressed)


## Triage (2021-10-04)


## Agenda (2021-10-28)

Agenda:



* AI review, announcements
    * Code freeze is coming up: November 16
    * AI from previous meeting: no way to turn off PRs in staged repos, could turn off issues but contribex doesn’t support it yet; the most we could do would be to set up a bot to warn people
* [sallyom] Kubelet OpenTelemetry tracing KEP
    * [https://github.com/kubernetes/enhancements/issues/2831](https://github.com/kubernetes/enhancements/issues/2831)
    * [https://github.com/kubernetes/kubernetes/pull/105126](https://github.com/kubernetes/kubernetes/pull/105126)
* [sallyom] OperateFirst environment for testing tracing enabled
    * [https://www.operate-first.cloud/](https://www.operate-first.cloud/)  open source operations
    * [https://github.com/operate-first/support/issues/401](https://github.com/operate-first/support/issues/401)
* [logicalhan] a bunch of metrics are getting proposed to promote to stable
    * [https://github.com/kubernetes/kubernetes/issues/105861](https://github.com/kubernetes/kubernetes/issues/105861)
    * [https://github.com/kubernetes/kubernetes/issues/105862](https://github.com/kubernetes/kubernetes/issues/105862)
    * [https://github.com/kubernetes/kubernetes/issues/105864](https://github.com/kubernetes/kubernetes/issues/105864)


## Triage (2021-10-21)

Cancelled as all four leads were unavailable to start the meeting.


## Agenda (2021-10-14)

Agenda:



* CANCELED DUE TO KUBECON NA


## Triage (2021-10-07)

Attendees:



*


## Agenda (2021-09-30)

Agenda:



* CANCELED DUE TO NO AGENDA


## Triage (2021-09-23)

Attendees:



* Ehashman
* Dashpole
* CatherineF
* Kefan Yang
* dgrissonet

Agenda:



* Closed some issues/PRs in [https://github.com/kubernetes/metrics](https://github.com/kubernetes/metrics)
    * **Action:** ehashman to ask if we can disable PRs/issues in this repo
        * PRs: no
        * Issues: maybe
        * [https://kubernetes.slack.com/archives/C1TU9EB9S/p1632416603223800](https://kubernetes.slack.com/archives/C1TU9EB9S/p1632416603223800)


## Agenda (2021-09-16)

Agenda:



* CANCELED DUE TO NO AGENDA


## Triage (2021-09-09)


## Agenda (2021-09-02)

Attendees:



* Dashpole
* Brancz
* Ehashman
* Erain
* Catherine Fang
* Yashika Badaya

Agenda:



* [logicalhan] Revisit stability classes
* [ehashman] KEP review for 1.23
    *


## Agenda (2021-08-19)

Cancelled?


## Agenda (2021-08-05)

Attendees:



*
*
*
* kakkoyun

Agenda:



* [deads2k] Default metrics cardinality
    * [https://github.com/kubernetes/kubernetes/issues/104008](https://github.com/kubernetes/kubernetes/issues/104008) Cardinality regression in 1.22
    * [https://github.com/kubernetes/kubernetes/pull/102523](https://github.com/kubernetes/kubernetes/pull/102523)
    * [ehashman] In 1.22, we had a metric added that accidentally included a namespace dimension. This caused a cardinality explosion which wasn’t detected until Red Hat performed upgrade testing in downstream OpenShift by running e2e tests and saw large memory regressions for the Prometheus instances, causing nodes to go not ready.
    * **How can we prevent this in the future rather than reacting to this many months later in downstream integration testing?**
    * [aojea] SIG Scalability currently isn’t gathering all the metrics, which means we can’t see a trend in the number of overall metrics.
    * There is a hook in the scalability perf framework that we could potentially use.
    * If you don’t run the correctness tests, you’d not generate the namespaces so we wouldn’t have caught it. So we need to ensure we run full E2Es/conformance.
    * [aojea] Other CI/e2es don’t run a Prometheus so scalability tests will be the easiest as they have one.
    * [deads2k] Rather than a scalability suite, we could add a [Late] annotation like in OpenShift that run the tests in gingko last in the e2e suite and thus could hit a metrics endpoint and block a PR if they cause a large regression.
        * [https://github.com/openshift/origin/blob/ab34e3b4a313d779270c12f9a28f82c23a21c201/pkg/test/ginkgo/cmd_runsuite.go#L249-L251](https://github.com/openshift/origin/blob/ab34e3b4a313d779270c12f9a28f82c23a21c201/pkg/test/ginkgo/cmd_runsuite.go#L249-L251)
    * [ehashman] why not both? Start with scalability tests to get an idea of baseline numbers, perhaps adding e2e blocking tests later once we have the machinery in place.
    * **Action (ehashman):** Investigate adding total metric counts in scalability tests with existing Prom. File an issue describing what we want in the perf-tests repo.
        * [https://github.com/kubernetes/perf-tests/issues/1870](https://github.com/kubernetes/perf-tests/issues/1870)
    * **Action (coffeepac):** Look into adding metrics-grabber tests to e2e suite.
        * [https://github.com/kubernetes/kubernetes/pull/102050](https://github.com/kubernetes/kubernetes/pull/102050)
* [serathius] Klog kep
    * [https://github.com/kubernetes/enhancements/issues/2845](https://github.com/kubernetes/enhancements/issues/2845)
    * Please comment!


## Triage (2021-07-29)

Attendees:



*


## Agenda (2021-07-22)

Attendees:



* Brancz
* Logicalhan
* Dashpole
* ehashman
* Dgrisonnet
* Coffeepac
* Gaurav Tiwari
* Catherine Fang
* Joadavis
* jpbetz

Agenda:



* [logicalhan] Metrics stability (adding beta phase)
    * Han reached out to SIG Arch but did not get a response
        * **Action:** Han to add to next week’s SIG Arch agenda
    * WG Reliability suggested that we should add this but trail graduation
        * Will this delay graduation?
        * No: e.g. beta feature has alpha metrics, feature can go GA with a beta metric
        * Could we move back metric requirements in PRR?
        * [ehashman] No, that shouldn’t be necessary; idea of metrics as part of PRR is that beta features are (in 95% of cases) on by default, and therefore people need a way to debug them, so we require people to define how to measure feature perf with metrics. Those metrics don’t need to be net new for the feature.
        * [logicalhan] Maybe we could use this as criteria for promoting metrics? If other KEPs rely on them?
* [logicalhan, lili, jpbetz] Donating auger to SIG Instrumentation
    * Looks like a useful thing, as it relates to observability it falls under our charter
    * Would increase visibility for the project and also hopefully improve diversity of maintainers
    * **Agreed,** we will accept as SIG Instrumentation project under kubernetes-sigs
    * **Action:** Han to kick off ownership process


## Triage (2021-07-15)

Attendees:



* Dashpole
* Ehashman
* Logicalhan
* Dgrissonet
* Catherine
* Parul


## Agenda (2021-07-08)



* [logicalhan] New triage time to avoid conflict
    * Proposed: same time as SIG Instrumentation but alternating weeks
* [logicalhan/ehashman] k/k Reviewers updates
    * Congrats to Lili on becoming a reviewer!!
    * [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)
    * [https://github.com/kubernetes/community/blob/master/contributors/guide/expectations.md#code-review](https://github.com/kubernetes/community/blob/master/contributors/guide/expectations.md#code-review)
* Ping about structured logging status from RT
    * Might miss on feature parity between json and text due to deprecation work that requires a new KEP
    * Will decide on scope in 1.23
* Congrats to dashpole on landing tracing!!
    * [https://github.com/kubernetes/kubernetes/pull/103216](https://github.com/kubernetes/kubernetes/pull/103216) remaining PR
* CRI stats work
    * Frederic commented after merge: concerned about metrics differences between cgroupsv1 and v2
    * [ehashman] KEP just removes delegation to cadvisor and instead talks to runtime via CRI directly with rpc calls; cadvisor would have the same issue with v1 vs. v2
    * Comments to go on PRs in flight: [https://github.com/kubernetes/kubernetes/pull/102789](https://github.com/kubernetes/kubernetes/pull/102789) [https://github.com/kubernetes/kubernetes/pull/103095](https://github.com/kubernetes/kubernetes/pull/103095)


## Triage (2021-06-30)

Attendees:



* Logicalhan
* Ehashman
* dgrisonnet


## Agenda (2021-06-24)



* Reminder: code freeze is July 8
* [ehashman] KEP review
    * Only one Instrumentation-owned KEP: tracing [API Server Tracing · Issue #647 · kubernetes/enhancements · GitHub](https://github.com/kubernetes/enhancements/issues/647)
    * Do we have reviewers/approvers? On track?
* [lilic] Continue discussion around promoting metrics stability that we started in triage
    * [Get rid of duplicate set of metrics for watch counts · Issue #102545 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/issues/102545)
    * [https://github.com/kubernetes/kubernetes/pull/102595](https://github.com/kubernetes/kubernetes/pull/102595)
        * Issue 1: duplicated metrics, one is a subset of the other
        * Issue 2: the superset metric doesn’t comply with the naming guidelines so we can’t promote it to stable
        * Issue 3: we don’t have a proposal for either of these to promote to stable
        * Issue 4: we don’t own this metric, need this to be driven by API machinery, but we can make recommendations to get something promoted
        * [lilic] We haven’t fully agreed on what makes it stable - should be it something we can alert or dashboard on?
        * Proposal: propose a new metric that meets the criteria to replace both, and then deprecate both
        * **Action:** Han to drive (make it better)
* [logicalhan] Do we need a beta stage for metrics?
    * Since everything is alpha but we don’t have a lot of stable metrics, people are relying on alpha
    * Many metrics will never be promoted but we don’t have a way to incentivize that
    * Metric graduation doesn’t follow the standard feature flag cycle (nothing exists between alpha and GA)
    * Experimental/debug metrics?
    * **Action:** Need an initial proposal. Han and Frederic volunteer to start a draft.
* [serathius] proposal for deprecating klog flags in core k8s components [Json format should support same set of feature flags as klog · Issue #99270 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/issues/99270#issuecomment-864605988)
    * Describe case by case why particular flag is hard to implement and should be deprecated.

Attendees:



* logicalhan
* ehashman
* dashpole
* lili
* dgrisonnet
* catherine
* marek
* brancz


## Triage (2021-06-16)


## Agenda (2021-06-10)

Attendees:



* ehashman
* logicalhan
* dashpole
* Marek
* joadavis
* dgrisonnet
* Pat Christopher
* Gerassimos
* Filip
* Yu Yi
* Scott
* [ehashman] Putting out a call for help with metric documentation
    * We have a lot of metrics. We don’t have any documentation. Do we want to expand the static analysis we have for stable metrics to alpha ones, and work on improving the documentation available? Perhaps we could add docstring annotations similar to the conformance tests to allow for expanded documentation?
    * Problems with static analysis for alpha ones: doesn’t work necessarily for the alpha types. Variable names, concatenated strings, etc. Some metrics are automatically generated (e.g. from kubelet), custom collectors also cause
    * Static analysis is not super resilient; only parses if something is stable
    * Let’s not let the perfect be the enemy of the good; we only have 4 metrics in our docs right now, and we could add a lot more
    * Static analysis would need to be improved before we can do more of this stuff
    * Pat Christopher would be interested in digging into this if we can get bugs filed
    * **Action:** Han to file some bugs detailing the issues with current static analysis to unblock doc generation of metrics
    * We would also need for documentation for how the static analysis works (KEP has a lot of detail) -- good starting point for the developer docs
    * After we fix the static analysis, we would need to parse the data and then we can autogenerate the docs for the website
* [ehashman] Back from SIG Node on [https://github.com/kubernetes/kubernetes/issues/101851#issuecomment-848101063](https://github.com/kubernetes/kubernetes/issues/101851#issuecomment-848101063)
    * Confirmed that we will not add node start time.
* [logicalhan] Recent issue: someone’s trying to remove a metric because it’s alpha, but all metrics are alpha so it isn’t necessarily safe to remove. How do we handle this?
    * Can’t really tell people to follow the stable/alpha policy when there are only 4 stable metrics, most components don’t have _any_ stable metrics
    * Can we make a policy for deprecation? Say, if a metric has been in at least 2 releases, there needs to be a 1-release deprecation period?
        * [han] Maybe even longer: if it’s been in for 4+ releases, need a 1-release deprecation period
        * All metric removals must be accompanied by ACTION REQUIRED release notes, on both deprecation and removal.
        * We could enforce this with the tooling if metrics had a version they were introduced.
        * [dashpole] What if we only allowed metrics to stay in alpha for a set number of releases? Need to prevent “perma-beta”, we make this an explicit decision point by the maintainers of each component
        * Can’t really force a metric to stable, because some metrics aren’t suitable for stable (e.g. constantly changing)
    * Could also introduce a beta metrics phase
        * We would need a KEP for this
        * **Action:** If anyone wants to pick this up and write a proposal, **help is wanted**.
    * We need a policy for deprecation of alpha metrics. We’ve informally had a policy for years but we’ve never written it down.
        * Right now 90%+ of metrics in kube have no guarantees; we can’t just remove things randomly.
        * Suggestion: we need to formalize a policy in our community docs.
        * **Action:** Han to open a PR for discussion.


## Agenda (2021-05-27)



* [serathius/dgrisonnet] Create metrics-api-machinery project that encompasses both core, custom, and external metrics.
    * Name:
        * … &lt;please propose>
        * metrics-api-machinery
    * TODO
        * Create new repo, migrate code and deprecate to don’t break backward compatibility
        * Who wants to do the work?

Attendees



*


## Triage (2021-05-19)

Attendees



*


## Agenda (2021-05-13)



* Announcement: enhancements deadline is TODAY
* Quick KEP review for 1.22
    * Pushed [https://github.com/kubernetes/enhancements/issues/1668](https://github.com/kubernetes/enhancements/issues/1668) to 1.23
    * Dropped [https://github.com/kubernetes/enhancements/issues/2305](https://github.com/kubernetes/enhancements/issues/2305) from 1.22 due to lack of resourcing
* [coffeepac] Reviews for [https://github.com/kubernetes-sigs/instrumentation-addons/pull/1](https://github.com/kubernetes-sigs/instrumentation-addons/pull/1) and 2
    * Action: add coffeepac and monotek to kubernetes-sigs
        * [https://github.com/kubernetes/org/pull/2714](https://github.com/kubernetes/org/pull/2714)
* [directxman12] quick overview of kubebuilder docs tooling
    * [https://github.com/rust-lang/mdBook](https://github.com/rust-lang/mdBook) ← tool that we use
        * [https://rust-lang.github.io/mdBook/](https://rust-lang.github.io/mdBook/) ← docs for that tool
    * [https://github.com/kubernetes-sigs/kubebuilder/tree/master/docs/book/src](https://github.com/kubernetes-sigs/kubebuilder/tree/master/docs/book/src) ← our book
    * [https://github.com/kubernetes-sigs/kubebuilder/blob/master/netlify.toml](https://github.com/kubernetes-sigs/kubebuilder/blob/master/netlify.toml) ← our netlify config
* [serathius] Expose start_time in Kubelet [https://github.com/kubernetes/kubernetes/issues/101851](https://github.com/kubernetes/kubernetes/issues/101851)
    * Problem with start_time  [https://github.com/kubernetes/kubernetes/issues/101902](https://github.com/kubernetes/kubernetes/issues/101902)

Attendees:



* logicalhan
* ehashman
* Andrew Pollack
* dashpole
* Marek
* Yu Yi
* Solly Ross
* Joadavis
* Kristin Barkardottir
* Nikos Fotiou
* Pat Christopher
* John


## [CANCELLED] Triage (2021-05-05)



* Cancelled due to conflict with KubeCon


## Agenda (2021-04-29)



* Announcement: Release schedule is out
* klog ownership: [https://github.com/kubernetes/klog/issues/222](https://github.com/kubernetes/klog/issues/222)
    * OWNERS are outdated
    * Action: Frederic to audit OWNERS file
* [ehashman] Finalizing 1.22 planning
    * [https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#](https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#)
    * Done!!
    * Everything added to 1.22 enhancements tracking sheet: [https://docs.google.com/spreadsheets/d/1mlui0brYypOAsgS2D13fvcs3At1uMq4i1gWfneq-jxY/edit#gid=1954476102](https://docs.google.com/spreadsheets/d/1mlui0brYypOAsgS2D13fvcs3At1uMq4i1gWfneq-jxY/edit#gid=1954476102)
* [lilic] etcd issue on tracing
    * If we don’t get it landed soon, it will take over a year for next release
    * [https://github.com/etcd-io/etcd/issues/12460](https://github.com/etcd-io/etcd/issues/12460)
    * Blocking GRPC dependencies got fixed, there will be a release soon
    * Lili will propose at etcd meeting next week
    * 1 year delay on code = probably 2 years to reach users, given the delay of uptake of new k8s releases
    * dashpole: Go SDK/API are not stable yet in opentelemetry, could be a blocker
* Cancel triage next week for KubeCon?
    * Yes.

Attendees:



* logicalhan
* brancz
* ehashman
* Andrew Pollack
* Damien Grisonnet
* dashpole
* Kemal Akkoyun
* Lili Cosic
* Marek
* Yu Yi
* Matthias Loibl


## Triage (2021-04-22)



* [from liggitt] [https://github.com/kubernetes/kubernetes/issues/75551#issuecomment-823306379](https://github.com/kubernetes/kubernetes/issues/75551#issuecomment-823306379)


## Agenda (2021-04-15)

Issues:



* [scott] Question for Elana: based on [project 53](https://github.com/orgs/kubernetes/projects/53) should we be benchmarking or looking out for certain errors or panics when doing structured log migrations?
    * Issue for benchmarking [https://github.com/kubernetes/kubernetes/issues/99803](https://github.com/kubernetes/kubernetes/issues/99803)
* [serathius] Structured logging WG update
    * Collecting approvals currently
    * Next Monday, approval will go to Steering with dims as Liaison
* [scott] Worth making a documentation site like [kubebuilder](https://book.kubebuilder.io/) to consolidate information in KEPs, READMEs, etc.?
    * [ehashman] Referring to SIG Docs/Leads meeting - probably May meetings
* [ehashman] 1.22 planning
    * [https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#](https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#)
* [ksm maintainers] We cut kube-state-metrics v2.0 - test and report any issues :tada:
    * [https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0](https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0)
* [logicalhan] Update on ownership of Event code

Attendees:



* logicalhan
* dashpole
* ehashman
* marek
* Eddie zaneski
* Scott
* Joadavis
* Yu yi
* Kemal akkoyun


## Triage (2021-04-07)

Note: we have begun removing sig/instrumentation labels from Structured Logging PRs in favour of area/logging.

Attendees:



* ehashman
* dashpole
* scott


## Agenda (2021-04-01)

Issues:



* [voutcn] Make webhook-caused critical request failures more visible: [doc](https://docs.google.com/document/d/1IH3zFw9TMOy1E7hunj-anD_VuyAEOrpo2vKDQ0db6p4/edit)
* [coffeepac] moving fluentd-elasticsearch to kubernetes-sigs/instrumentation-tools from cluster/addons (or sending it someplace else, time for it to go)
* [serathius] wg structured logging formation updates
    * Creation process starts next week

Attendees:



* dashpole
* kakkoyun
* coffeepac
* voutcn
* logicalhan


## Triage (2021-03-24)

Attendees:



* logicalhan
* marek
* brancz


## Agenda (2021-03-18)

Issues:



* Reminder: test freeze is Mar. 24
* 2021 Contributor Survey [https://www.surveymonkey.com/r/k8scommsurvey2021](https://www.surveymonkey.com/r/k8scommsurvey2021)
* [logicalhan] Logging WG
    * Interested List: serathius, kakkoyun, ehashman, erain…
    * **Action:** Han to kick off the process of chartering WG
        * [https://groups.google.com/g/kubernetes-sig-architecture/c/ao1sviI_4rc/m/NG9BoE4GAwAJ](https://groups.google.com/g/kubernetes-sig-architecture/c/ao1sviI_4rc/m/NG9BoE4GAwAJ)
* [ehashman] stackdriver docs: who should review this? Do we want to keep this page? [https://github.com/kubernetes/website/pull/26925](https://github.com/kubernetes/website/pull/26925)
    * **Action:** ehashman to follow up on item saying we are okay with removing this, but we believe SIG Cloud Provider/GCP should probably own this
        * [https://github.com/kubernetes/website/pull/26925#issuecomment-802101121](https://github.com/kubernetes/website/pull/26925#issuecomment-802101121)
* [dgrisonnet] Make [PodMetrics](https://github.com/kubernetes/metrics/blob/v0.20.4/pkg/apis/metrics/v1beta1/types.go#L63-L74) account for everything that is charged into the pod cgroups
* [bboreham] [kspan](https://github.com/bboreham/kspan) demo

Attendees:



* ehashman
* logicalhan
* kakkoyun
* dgrisonnet
* dashpole
* bboreham
* brancz
* marek
* lilic
* Scott Lee
* Yuchen Zhou
* metalmatze
* joadavis


## Triage (2021-03-10)

Notes:



* Helm chart releases are noisy in KSM [https://github.com/kubernetes/kube-state-metrics/issues/1392](https://github.com/kubernetes/kube-state-metrics/issues/1392)
* Helm chart requested for metrics-server [https://github.com/kubernetes-sigs/metrics-server/issues/572](https://github.com/kubernetes-sigs/metrics-server/issues/572)
* **Action:** ehashman to follow up with Instrumentation steering rep about “Official Helm Charts” for KSM which is causing a lot of churn
    * [https://kubernetes.slack.com/archives/C20HH14P7/p1615398698008400](https://kubernetes.slack.com/archives/C20HH14P7/p1615398698008400)
* **Action:** brancz to request creation of a kubernetes-sigs/instrumentation repo
    * Done! [https://github.com/kubernetes-sigs/instrumentation](https://github.com/kubernetes-sigs/instrumentation)

Attendees:



* Ehashman
* Logicalhan
* Brancz
* Dashpole
* Serathius
* lilic


## Agenda (2021-03-04)

Issues:



* Reminder: code freeze is March 9th! Next week!
* [ehashman] Structured logging sync
    * ~~Add new repo for the logging migration tool? [https://kubernetes.slack.com/archives/CHGFYJVAN/p1614606853044600](https://kubernetes.slack.com/archives/CHGFYJVAN/p1614606853044600)~~ (We added this to klog.)
    * **Action:** Han to investigate WG for logging.
* Ownership of events~~ e2e tests~~ [https://github.com/kubernetes/kubernetes/pull/99495](https://github.com/kubernetes/kubernetes/pull/99495)
    * **Action:** Matthias Loibl to review code and volunteer for approver?
* [lili] We just cut the first release candidate for v2.0 kube-state-metrics - v2.0.0-rc.0! Test away and please report any issues. [https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-rc.0](https://github.com/kubernetes/kube-state-metrics/releases/tag/v2.0.0-rc.0)
* [Huang-Wei] sig-scheduling needs help on review of [#99228](https://github.com/kubernetes/kubernetes/pull/99228) and [#99472](https://github.com/kubernetes/kubernetes/pull/99472)
* [logicalhan] Stable metric PRs:
    * [https://github.com/kubernetes/kubernetes/pull/99785](https://github.com/kubernetes/kubernetes/pull/99785)
    * [https://github.com/kubernetes/kubernetes/pull/99788](https://github.com/kubernetes/kubernetes/pull/99788)

AI:



* Update docs to reference structured logging
    * [https://github.com/kubernetes/community/issues/5593](https://github.com/kubernetes/community/issues/5593)

Attendees:



* logicalhan
* ehashman
* dashpole
* brancz
* joadavis
* metalmatze
* Huang-Wei
* Scott
* kakkoyun


## Triage (2021-02-24)

Attendees:



* Ehashman
* Dashpole
* Kakkoyun
* Steve Nguyen
* Joseph A Davis
* Serathius


## Agenda (2021-02-18)

Issues:



* [ehashman] Reminder: code freeze is March 9th
* [ehashman] Annual OWNERS files/org cleanup
    * Inactive members and reviewers have been removed
    * We would love to see more people step up to become reviewers!
    * How to reviewer: [https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md)
* [ehashman] Annual report is coming!
* SIG status report for 1.21 feature dev
    * Han: metrics stability has people assigned for all the 1.21 components (selecting stable metrics, escape hatch), on track
    * Marek: would be helpful to have automatic checking for schemas/conventions for structured logging; need better guidance in docs on using keys
    * **ACTION:** serathius to do some analysis of present labels/keys and work on determining conventions
* [brancz] otel datapoint
    * OpenTelemetry Prometheus WG: Meetings Wednesdays at 8PT ([https://github.com/open-telemetry/community#calendar](https://github.com/open-telemetry/community#calendar)) [https://docs.google.com/document/d/19bnXziPn2MZ9wO6684UoI4D-LCjGL5bTJkGhux29bx8/edit#heading=h.svcors6u1w3t](https://docs.google.com/document/d/19bnXziPn2MZ9wO6684UoI4D-LCjGL5bTJkGhux29bx8/edit#heading=h.svcors6u1w3t)
    * Could we migrate all of Kubernetes otel compat (instead of Prometheus)?
        * **SIG consensus: **Not feasible in the short-term.
        * Would be an enormous undertaking for probably not much value
        * /metrics endpoint of in-process metrics are a stability guarantee, so if ever wanted otel must bring this

Attendees:



* ehashman
* logicalhan
* voutcn
* Metalmatze
* brancz
* Serathius
* lilic
* dgrisonnet
* kakkoyun
* dashpole
* erain


## Triage (2021-02-10)



* triage/unresolved label doesn’t remove the needs-triage label
* What do we do when we don’t want to accept a PR but we also want to mark that we’ve looked at it?

Attendees:



* Ehashman
* Dashpole
* serathius


## Agenda (2021-02-04)

Issues:



* [ehashman] 1.21 KEP finalization!
    * See: [https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#](https://docs.google.com/document/d/1ZxIfu8_ZBMdVanrbXYseJZeSGbeupKfevq1pMDvJGps/edit?ts=6011b13c#)
* [ehashman] OWNERS files/org cleanup (in advance of our annual report)
* …

Attendees:



* Ehashman
* Brancz
* Kakkoyun
* Logicalhan
* Yu yi
* Marek
* Dashpole
* ?


## Triage (2021-01-27)

Attendees:



* Ehashman
* Logicalhan
* Scott
* joadavis


## Agenda (2021-01-21)

Issues:



* [serathius] Demo of kubectl top with custom metrics
    * Implemented in the kubectl client
    * See recording for video
    * Only ‘top pods’ implemented so far, but other objects can be added
    * Code: [https://github.com/kubernetes/kubernetes/pull/98114](https://github.com/kubernetes/kubernetes/pull/98114)
    * Contact serathius@ if your interested in helping
* [ehashman] KEP review
* [brancz] prometheus-k8s-adapter has been accepted into kubernetes-sigs!
    * Got past legal issues
    * [https://github.com/kubernetes/org/issues/2182](https://github.com/kubernetes/org/issues/2182)

Attendees:



* Ehashman
* Brancz
* logicalhan
* damien grisonnet
* erain
* joadavis
* Serathius
* kakkoyun


## Triage (2021-01-13)

Attendees:



* Dashpole
* Logicalhan
* Ehashman
* Akonarde


## Agenda (2021-01-07)


## Issues:



* [ehashman] Broken cadvisor machine metrics in 1.19+ [https://github.com/kubernetes/kubernetes/issues/95204](https://github.com/kubernetes/kubernetes/issues/95204)
    * tl;dr: machine_* metrics all disappeared in 1.19/1.20, bug sat for 3mo
    * How to prevent this sort of thing from happening in the future?
    * cadvisor metrics are sort of a special case, should be handled specially
    * Need to finish up KEP for metrics in kubelet metrics endpoint: [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/727-resource-metrics-endpoint](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/727-resource-metrics-endpoint)
* [ehashman] Stable metrics in 1.21?
    * SIGs can propose them, and they must maintain stability
    * What is the process? Documented in the KEP
        * Set metadata to stable
        * Tag SIG Instrumentation for review
    * We need to send out project-wide comms
        * **Action:** Han to send email after testing and socializing with API Machinery
    * Also need a doc in the k/community repo
    * Kubernetes official slos
        * [https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/slos.md)

Attendees:



* Dashpole
* ehashman
* erain
* Akonarde
* Kakkoyun
* logicalhan
* Metalmatze
* Lilic
* Brancz
* joadavis

