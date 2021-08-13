# SIG Node Meeting Notes (2020)

## Future

* `regular resource usage tracking resource tracking for 100 pods per node` =&gt; this e2e test has been failing (flaking?) quite consistently in release-blocking dashboards. Should we block the 1.14 release on this? If not, could you help resolve it? [Issue:75039](https://github.com/kubernetes/kubernetes/issues/75039) (@mariantalla)
* Issue to discuss: [Hardware topology awareness at node level (including NUMA)](https://github.com/kubernetes/kubernetes/issues/49964)
* OCI Hooks PreStart, PostStop - @alban
  * Suggestion to add Kubernetes labels in the OCI State ([slide 15](https://docs.google.com/presentation/d/1i8csKAf15j3ZDeHxuUlHBDvHiBI44_UxATdaaAy-pjE/edit#slide=id.g5da9f7933e_0_51))
  * Use case: [Inspektor Gadget](https://github.com/kinvolk/inspektor-gadget)
  * Previous presentation at the OCI meeting, July 24th: [notes](https://hackmd.io/El8Dd2xrTlCaCG59ns5cwg?view) ; [slides](https://docs.google.com/presentation/d/1i8csKAf15j3ZDeHxuUlHBDvHiBI44_UxATdaaAy-pjE/edit)
* Issues to discuss:
  * [Handling timeouts on calls to CreateContainer](https://github.com/kubernetes/kubernetes/issues/94085)
  * [Redirect container stdout / stderr to file](https://github.com/kubernetes/kubernetes/issues/94892)
* Can hugepages be handled like other namespaced sysctls?
* Can [this project](https://github.com/openebs/node-disk-manager) be considered under sig-node?

## Dec 22 & 29, 2020 Cancelled

Cancelled.  Merry holiday season!

## Dec 15, 2020 Cancelled

No agenda. Cancelled. Merry holiday season!

## Dec 8, 2020

Total active pull requests: [196](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+6 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**14**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-12-01T17%3A00%3A00%2B0000..2020-12-08T17%3A50%3A44%2B0000)                                           |Closed:      |[5](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-12-01T17%3A00%3A00%2B0000..2020-12-08T17%3A50%3A44%2B0000)|
|Updated:    |[**60**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-12-01T17%3A00%3A00%2B0000..2020-12-08T17%3A50%3A44%2B0000+created%3A%3C2020-12-01T17%3A00%3A00%2B0000)|Merged:      |[3](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-12-01T17%3A00%3A00%2B0000..2020-12-08T17%3A50%3A44%2B0000)              |

* [@rata]: Friendly ping to review [sidecar pre-proposal from Tim Hockin](http://bit.ly/thockin-pod-lifecycle-phases-proposal)
  * Next steps?
  * Discussing issue with the init containers and where the pod takes it’s state from. Maybe we need more check points and instead of runtime telling the status we need to keep the status in the kubelet. Kubelet knows way more about the status. So kubelet will not mirror runtime’s understanding of containers state into the API.
  * Dawn - yes, with the checkpoint side cars will be a reality. Dependency are needed.
  * Derek: I want to emphasys there are people working on making what we have working and reliable that needs to be appreciated.
  * Dawn: GC behavior is also not very well defined and needs to be formalized or even redesigned.
  * Derek: what’s missing in proposals is a confident implementation proposal that will make things better rather than more unreliable.
  * Sergey: we need to try to agree on the north star and split it into the road map later. Mrunal: have a feeling that the north star is a DAG. Rodrigo: maybe more phases and custom phases would be a better north star.
  * Derek: also need to make sure we are addressing the keystone container issue where sidecar container failure must stop the main container as it provides security or something.
  * Dawn: next time lets invite Tim. Maybe not the next week, let people time to read it.
  * Dawn: very important to make sure we listed scenarios to support and scenarios that will not be supported.
* [@SergeyKanzhelev] dockershim deprecation - plan review
  * We may need to revisit the plan of not compiling dockershim in 1.22 in end of Jan when we will have first customers on windows and some commitments from telemetry vendors.
  * Derek, Dawn: yes, we need to make sure we are not breaking things. Pushing a few releases should be feasible.
  * Dawn: maybe even limit it to windows-only as an alternative
  * Dawn: important to keep dockershim testing in OSS so we are not breaking anybody.
  * Mike Brown: we may need to keep the tests running even when dockershim is not released with k8s, but is still in-tree (1 year after 1.22 as the current plan states). Additionally, there will be a second dockershim that is external to kubelet, see cri-dockershim project. This external shim should be able to be tested in a similar fashion to containerd/cri-o after adding proper infra impl.
* [@ehashman] SIG Node Triage meeting?
  * Have previously discussed in the past
  * Sync on current test health efforts first before adding another meeting?

## Dec 1, 2020

Total active pull requests: [188](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+5 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**18**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-11-24T17%3A00%3A00%2B0000..2020-12-01T17%3A54%3A08%2B0000)                                           |Closed:      |[5](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-11-24T17%3A00%3A00%2B0000..2020-12-01T17%3A54%3A08%2B0000)|
|Updated:    |[**42**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-11-24T17%3A00%3A00%2B0000..2020-12-01T17%3A54%3A08%2B0000+created%3A%3C2020-11-24T17%3A00%3A00%2B0000)|Merged:      |[8](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-11-24T17%3A00%3A00%2B0000..2020-12-01T17%3A54%3A08%2B0000)              |

* [mauriciovasquezbernal]: [User ns KEP](https://github.com/kubernetes/enhancements/pull/2101)
  * Reminder asking for review.
* [SergeyKanzhelev] ContainerD tests: looking for volunteers (TODO: link)
* [Mrunalp] seccomp enabled by default? Provide a flag to disable
  * [Dawn] - this change will need a wide and loud notifications way ahead of the release as this default can break some vendor
  * Seccomp enabled by default already prevented some security CVE for container runtime - example is gAdvisor
  * Tim Alclair may also be interested in this topic
  * [Action] Mrunal will start writing proposal.

## Nov 24th, 2020

Total active pull requests: [182](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-6 from the last meeting) (note: discrepancy in numbers is due to label sig/node was applied or removed from PRs).

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**45**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-11-10T17%3A00%3A00%2B0000..2020-11-24T17%3A53%3A02%2B0000)                                            |Closed:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-11-10T17%3A00%3A00%2B0000..2020-11-24T17%3A53%3A02%2B0000)|
|Updated:    |[**115**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-11-10T17%3A00%3A00%2B0000..2020-11-24T17%3A53%3A02%2B0000+created%3A%3C2020-11-10T17%3A00%3A00%2B0000)|Merged:      |[35](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-11-10T17%3A00%3A00%2B0000..2020-11-24T17%3A53%3A02%2B0000)              |

* [SergeyKanzhelev] Conformance testing for RuntimeClass [https://github.com/kubernetes/kubernetes/issues/96524#issuecomment-731443843](https://github.com/kubernetes/kubernetes/issues/96524#issuecomment-731443843)
  * Will include more tests for conformance in 1.21
  * Will review other features if they needs to be added to conformance list
* [Hippie Hacker]
  * [https://apisnoop.cncf.io/](https://apisnoop.cncf.io/)
  * [https://kubernetes.slack.com/archives/C0BP8PW9G/p1606242452276000](https://kubernetes.slack.com/archives/C0BP8PW9G/p1606242452276000)
    * I would like a chance to pair with someone from [#sig-node](https://kubernetes.slack.com/archives/C0BP8PW9G) to go through the conformance process for the technical debt for ‘not tested’ Node related endpoints if anyone is keen.
    * It would probably take about 30-45 minutes and the resulting k/k tickets would fully flesh out what needs to be done, and our team would likely be able to help with the test writing itself.
    * From [https://apisnoop.cncf.io/conformance-progress/endpoints/1.20.0?filter=untested](https://apisnoop.cncf.io/conformance-progress/endpoints/1.20.0?filter=untested)
      * deleteCoreV1CollectionNode
      * readCoreV1NodeStatu
      * replaceCoreV1NodeStatu
    * Though there may be others.

## Nov 17, 2020

Cancelled because of kubecon (<https://kubernetes.slack.com/archives/C0BP8PW9G/p1605631458154700>).

## Nov 10, 2020

Total active pull requests: [185](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-11 from the two weeks)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**59**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-10-27T17%3A00%3A00%2B0000..2020-11-10T17%3A44%3A51%2B0000)                                            |Closed:      |[24](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-10-27T17%3A00%3A00%2B0000..2020-11-10T17%3A44%3A51%2B0000)|
|Updated:    |[**129**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-10-27T17%3A00%3A00%2B0000..2020-11-10T17%3A44%3A51%2B0000+created%3A%3C2020-10-27T17%3A00%3A00%2B0000)|Merged:      |[46](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-10-27T17%3A00%3A00%2B0000..2020-11-10T17%3A44%3A51%2B0000)              |

29 lgtm, but not approved PRs: [https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved)

* 9 lgtm, but not approved test PRs: [https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved+label%3Aarea%2Ftest](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved+label%3Aarea%2Ftest)

**AI**: Sergey - can we exclude all PRs that are not sig/node specific. Release manager, API review blocked, etc.

* [derekwaynecarr] should we revisit [swap](https://github.com/kubernetes/kubernetes/issues/53533#issuecomment-334947625) in 1.21, volunteers to work together?
  * **AI**: Derek to start google doc
  * Alexander @kad interested to join
  * Karan (@karan) also interested
  * Peter (@haircommander) is interested
  * [added after meeting] @ehashman is also interested
* [dashpole/egernst] naming for pod metrics in [https://github.com/kubernetes/kubernetes/pull/95839](https://github.com/kubernetes/kubernetes/pull/95839)
  * kube_pod_resource_cpu_seconds_total (in-line with [clayton’s KEP](https://github.com/kubernetes/enhancements/pull/1916))
  * pod_cpu_usage_seconds_total (in-line with other metrics on the endpoint; [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/kubelet-resource-metrics-endpoint.md))
  * **AI**: David Ashpole to start slack thread
* [mrunalp] CRI beta next step
  * [https://github.com/kubernetes/kubernetes/pull/96387](https://github.com/kubernetes/kubernetes/pull/96387)
* [SergeyKanzhelev] RuntimeClass GA: [https://github.com/kubernetes/kubernetes/pull/95718](https://github.com/kubernetes/kubernetes/pull/95718)
  * Should v1alpa1 be removed?
    * no objections - will send a separate PR
* [@jpepin] [cadvisor metrics deprecation](https://github.com/kubernetes/kubernetes/issues/68522) in the kubelet and [timeline for proposed KEP](https://github.com/kubernetes/enhancements/pull/2130/files?short_path=58871be#diff-58871be1cadaf7855efac07908b9fa9b589e4aac74ce6badc2c0f5cd05af8faa)
* ~~[fromani] quick question: podresources API PRs: suitable for code freeze exception? Any concern? (already answered by the following bullet point)~~
* PRs for 1.20 to be landed this week!!!
  * [dims] [Deprecate Dockershim](https://github.com/kubernetes/kubernetes/pull/94624)
  * ~~[AlexeyPerevalov]: [Implement TopologyInfo and cpu_ids in podresources interface](https://github.com/kubernetes/kubernetes/pull/93243)~~
  * [fromani]: [podresources APIs: concrete resources apis: implement GetAllocatableResources](https://github.com/kubernetes/kubernetes/pull/95734)
  * [alukiano, cez, krzwiatrzyk] [Memory manager](https://github.com/kubernetes/kubernetes/pull/95479)
  * ~~[andrewsykim] kubelet credential provider plugin [https://github.com/kubernetes/kubernetes/pull/94196](https://github.com/kubernetes/kubernetes/pull/94196)~~
  * ~~[SergeyKanzhelev] RuntimeClass [https://github.com/kubernetes/kubernetes/pull/95718](https://github.com/kubernetes/kubernetes/pull/95718)~~
  * [David Porter] Graceful Node shutdown -- [https://github.com/kubernetes/kubernetes/pull/96129](https://github.com/kubernetes/kubernetes/pull/96129)
  * ~~[Mrunal Patel] Add CRI v1beta1 proto[https://github.com/kubernetes/kubernetes/pull/96387](https://github.com/kubernetes/kubernetes/pull/96387)~~
* [karan] NPD dev is blocked due to failing e2e presubmit test: [https://github.com/kubernetes/kubernetes/issues/95955](https://github.com/kubernetes/kubernetes/issues/95955)
  * ~~PR 1: [https://github.com/kubernetes/kubernetes/pull/96381](https://github.com/kubernetes/kubernetes/pull/96381) (need approval)~~
  * ~~PR 2: [https://github.com/kubernetes/kubernetes/pull/96262](https://github.com/kubernetes/kubernetes/pull/96262) (need rebase + approval)~~
* [David Porter] Call out for graceful node shutdown PR - [https://github.com/kubernetes/kubernetes/pull/96129](https://github.com/kubernetes/kubernetes/pull/96129)

## Nov 3rd, 2020

Today’s SIG Node was cancelled due to the election.

* [@egernst] still need help with review:  adding pod usage to kubelet metrics/resource endpoint: <https://github.com/kubernetes/kubernetes/pull/95839>

## October 27th, 2020

Total active pull requests: [192](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**24**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)                                            |Closed:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)|
|Updated:    |[**100**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000+created%3A%3C2020-10-20T17%3A00%3A00%2B0000)|Merged:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)              |

* [@dchen1107/Dims] cadvisor ownership and updates (5 mins)
  * Meet David Porter! [https://github.com/bobbypage](https://github.com/bobbypage)
  * Thank you David Ashpole for being a long time maintainer!
  * [dims] main issue - cAdvisor is both - binary and library. Only library being vendored in k8s. Worked over time to reduce the amount of vendored code. The problem is whether we want to keep going like this or we can split library out of the binary. Need more exploration. Having two different owners for cAdvisor pieces will be challenging
* [@dchen1107] [NPD](https://github.com/kubernetes/node-problem-detector) ownership (@vteratipally, @ForestCold) and updates (5 mins)
  * Welcome, Varsha and Hanfei!
  * Made NPD kube-apiserver independent.
  * Working on adding more metrics related to system stats on a linux based OS.
  * Working on expanding the problems detected related to IO errors, memory and file system errors.
  * Exploring the idea of working NPD as a library.
* [@egernst] pod level usage accounting. See [PR](https://github.com/kubernetes/kubernetes/pull/95839) - needs eyes/feedback
* [Riaan ii] Conformance testing of Node Proxy endpoint
  * [https://hackmd.io/nz9DHbSlQgaQh7CqxFhBzQ?view](https://hackmd.io/nz9DHbSlQgaQh7CqxFhBzQ?view)
* [@SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/95641](https://github.com/kubernetes/kubernetes/pull/95641) Should the +11 bytes to every http probe request out of the box be a big deal?
* [@SergeyKanzhelev] Review [SIG Node - Feature Health Check](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#heading=h.2g5gkv44vir8)
  * Enhancements [approved for 1.20](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.20+label%3Asig%2Fnode).
  * Enhancements status tracking: [Kubernetes 1.20 Enhancements Tracking](https://docs.google.com/spreadsheets/d/1Ch7PIapJhwdl83HEnYvwFNz96z3vk_eHnf8heLxxgyw/edit#gid=936265414)
  * Dates:
    * Friday, Nov 6th: Week 8 - [Docs Placeholder PR deadline](https://kubernetes.io/docs/contribute/new-content/new-features/#open-a-placeholder-pr)
    * Thursday, Nov 12th: Week 9 - [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#code-freeze)
  * Should we mark some features as deprecated? Or this is not possible at this stage?
  * What’s the status of [CRI alpha to beta](https://docs.google.com/document/d/1xMfiVUX4lF0G6Y_yQ4l-5-wrOneeJWcXY2NGVGuTdH8/edit)? ([Kubelet CRI support · Issue #2040 · kubernetes/enhancements](https://github.com/kubernetes/enhancements/issues/2040))
    * PR open for the seccomp changes <https://github.com/kubernetes/kubernetes/pull/95876/file>
* [@swsehgal/@aperevalov] Topology Aware Scheduling
  * Location of NodeResourceTopology CRD API definition
* [@k-wiatrzyk, @klueska, @alukiano, @cezaryzukowski, @p.rapacz, @bg-chun]
  * should Topology Manager Scope have a separate kubernetes/enhancement issue tracker?
    * Topology Manager tracker: [https://github.com/kubernetes/enhancements/issues/693](https://github.com/kubernetes/enhancements/issues/693)
  * just a reminder for review
    * Topology Manager Scope PR awaiting review: [https://github.com/kubernetes/kubernetes/pull/92967](https://github.com/kubernetes/kubernetes/pull/92967)
    * PR for Memory Manager awaiting review:
            [https://github.com/kubernetes/kubernetes/pull/95479](https://github.com/kubernetes/kubernetes/pull/95479)
            ( blocks the official documentation PR: [https://github.com/kubernetes/website/pull/24642](https://github.com/kubernetes/website/pull/24642) )
* [@fromani] podresources concrete assignments API - quick process question
  * The implementation of one of the proposed extensions (Watch() endpoint) is very unlikely to make it (capacity). If we need to postpone it to 1.21, do we need to fix anything (e.g. amend the KEP)?
    * [fromani] - sorry, I had an unexpected last-minute issue and could not join the meeting. Will make sure the docs reflect the fact the Watch() is postponed to 1.21. Need to update the KEP (per slack conversation).

## October 20th, 2020

Total active pull requests: [196](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+22 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**28**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-10-13T17%3A00%3A00%2B0000..2020-10-20T16%3A36%3A42%2B0000)                                           |Closed:      |[4](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-10-13T17%3A00%3A00%2B0000..2020-10-20T16%3A36%3A42%2B0000)|
|Updated:    |[**51**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-10-13T17%3A00%3A00%2B0000..2020-10-20T16%3A36%3A42%2B0000+created%3A%3C2020-10-13T17%3A00%3A00%2B0000)|Merged:      |[2](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-10-13T17%3A00%3A00%2B0000..2020-10-20T16%3A36%3A42%2B0000)              |

* [@mauriciovasquezbernal] User namespace
  * PR with new KEP [https://github.com/kubernetes/enhancements/pull/2101](https://github.com/kubernetes/enhancements/pull/2101)
  * Action Items:
    * klueska (nvidia) to help shephard review
    * derekwaynecarr to help approval (target 1.21)
* [@jeremyje] Node Problem Detector for Window
  * [https://docs.google.com/document/d/1eiK6KAp_TFR0PgBMu2WCf49fMZcg-HHnBHMc9fALquU/edit](https://docs.google.com/document/d/1eiK6KAp_TFR0PgBMu2WCf49fMZcg-HHnBHMc9fALquU/edit)
  * Action Items:
    * desire to get sub-project readout on npd to see if it needs supplemental support from community
    * will review doc and leave comments there
* [@k-wiatrzyk, @klueska, @alukiano, @cezaryzukowski, @p.rapacz, @bg-chun]
  * Topology Manager Scope PR awaiting review: [https://github.com/kubernetes/kubernetes/pull/92967](https://github.com/kubernetes/kubernetes/pull/92967)
  * PR for Memory Manager awaiting review:
        [https://github.com/kubernetes/kubernetes/pull/95479](https://github.com/kubernetes/kubernetes/pull/95479)
        ( blocks the official documentation PR: [https://github.com/kubernetes/website/pull/24642](https://github.com/kubernetes/website/pull/24642) )
* [@fromani - Francesco] Just operational Q: how to serialize the podresources PRs?
  * is separate PRs OK? - also, need reviewers :)
  * [https://github.com/kubernetes/kubernetes/pull/93243/](https://github.com/kubernetes/kubernetes/pull/93243/) extend current API
  * [https://github.com/kubernetes/kubernetes/pull/95734](https://github.com/kubernetes/kubernetes/pull/95734)  add new APIs (ready for review ETA 20201021)
  * Action Item
    * recommend following up with @renaud
* [@SergeyKanzhelev] (low pri) Deprecating PodUnknown podPhase [https://github.com/kubernetes/kubernetes/pull/95286](https://github.com/kubernetes/kubernetes/pull/95286)
* [@SergeyKanzhelev] since we are on time: [https://github.com/kubernetes/kubernetes/pull/95718/](https://github.com/kubernetes/kubernetes/pull/95718/) : Promote RuntimeClass to GA, including promoting the API to v1. Note, PodOverhead is part of the API and being promoted to v1 as it cannot be separated. We can consider it a beta feature of v1 API. Please advice if this is not recommended and whether there is a workaround.
* Do we want to start using triage tag?
  * Dims will take a look

## October 13th, 2020

Total active pull requests:

* [@rata, @jirving] sidecar KEP and next step
  * Rodrigo (rata) will add more use cases (“fatalToPod”, clarify this adds shutdown order guarantees too)
  * Critical container feature KEP: [https://github.com/kubernetes/enhancements/pull/912](https://github.com/kubernetes/enhancements/pull/912)
  * Rodrigo will start drafting ideas for the use cases, please write to coordinate if you want to do it too!
* [@gergely.csatari] Topology Aware Scheduling [Use cases](https://docs.google.com/document/d/1OJAejm4kdRDUqgWdlYQj9b9DBcLmZY71PrLH7eSvX6I/edit#)
  * Kubernetes in telecomunicaitons survey results: [https://docs.google.com/presentation/d/1moCT0Q6dsyUHYaGWe6sJY8_X543Wbft0HgBZ7X7xJeM/edit#slide=id.p1](https://docs.google.com/presentation/d/1moCT0Q6dsyUHYaGWe6sJY8_X543Wbft0HgBZ7X7xJeM/edit#slide=id.p1)

Total active pull requests: [172](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-2 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**18**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-10-06T17%3A00%3A00%2B0000..2020-10-13T17%3A00%3A00%2B0000)                                           |Closed:      |[6](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-10-06T17%3A00%3A00%2B0000..2020-10-13T17%3A00%3A00%2B0000)|
|Updated:    |[**52**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-10-06T17%3A00%3A00%2B0000..2020-10-13T17%3A00%3A00%2B0000+created%3A%3C2020-10-06T17%3A00%3A00%2B0000)|Merged:      |[14](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-10-06T17%3A00%3A00%2B0000..2020-10-13T17%3A00%3A00%2B0000)             |

## October 6th, 2020

Total active pull requests: [173](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-3 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**16**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-09-29T17%3A00%3A00%2B0000..2020-10-06T16%3A59%3A38%2B0000)                                           |Closed:      |[11](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-09-29T17%3A00%3A00%2B0000..2020-10-06T16%3A59%3A38%2B0000)|
|Updated:    |[**56**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-09-29T17%3A00%3A00%2B0000..2020-10-06T16%3A59%3A38%2B0000+created%3A%3C2020-09-29T17%3A00%3A00%2B0000)|Merged:      |[8](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-09-29T17%3A00%3A00%2B0000..2020-10-06T16%3A59%3A38%2B0000)               |

* [@dlipovetsky] Document requirement to drain before upgrading kubelet minor version ([k/website issue #12326](https://github.com/kubernetes/website/issues/12326))
  * It would also help to decide on + document the requirement for upgrading kubelet *patch* version
* [@hasheddan, @pjbgf, @saschagrunet]
  * seccomp-operator demo
  * seccomp-operator changes in scope &amp; rename to security-profiles-operator [PR](https://github.com/kubernetes-sigs/seccomp-operator/issues/127)
  * [https://github.com/kubernetes/enhancements/pull/1444](https://github.com/kubernetes/enhancements/pull/1444) KEP AppArmor GA
* [@swsehgal] PodResource API KEP [#1884](https://github.com/kubernetes/enhancements/pull/1884) Ready for final review
* [@SergeyKanzhelev/@harche] [RuntimeClass to GA](https://github.com/kubernetes/enhancements/pull/2071) - this is merged today
  * [https://github.com/kubernetes/kubernetes/pull/95046](https://github.com/kubernetes/kubernetes/pull/95046) and [https://github.com/kubernetes/kubernetes/pull/94796](https://github.com/kubernetes/kubernetes/pull/94796) would be good to approve
* [@mitar] [https://github.com/kubernetes/enhancements/pull/2013/](https://github.com/kubernetes/enhancements/pull/2013/) Expose container imageID through Downward API
  * Issue: [https://github.com/kubernetes/enhancements/issues/2012](https://github.com/kubernetes/enhancements/issues/2012)
  * Pending review

## September 29th, 2020

Total active pull requests: [174](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**8**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-09-22T17%3A00%3A00%2B0000..2020-09-29T16%3A51%3A20%2B0000)                                            |Closed:      |[7](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-09-22T17%3A00%3A00%2B0000..2020-09-29T16%3A51%3A20%2B0000)|
|Updated:    |[**64**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-09-22T17%3A00%3A00%2B0000..2020-09-29T16%3A51%3A20%2B0000+created%3A%3C2020-09-22T17%3A00%3A00%2B0000)|Merged:      |[9](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-09-22T17%3A00%3A00%2B0000..2020-09-29T16%3A51%3A20%2B0000)              |

* [@bobbypage] Node Shutdown KEP: [https://github.com/](https://github.com/kubernetes/enhancements/pull/2001)[kubernetes](https://github.com/kubernetes/enhancements/pull/2001)[/enhancements/pull/2001](https://github.com/kubernetes/enhancements/pull/2001)
* [@ambguo] Windows privileged container KEP ([KEP](https://docs.google.com/document/d/12EUtMdWFxhTCfFrqhlBGWV70MkZZPOgxw0X-LTR0VAo/edit?usp=sharing))
  * [enhancement issue](https://github.com/kubernetes/enhancements/issues/1981) for tracking
  * Looking for sig-node reviewer
  * reviewing with sig-auth
* [@alukiano,@cezaryzukowski,@krzwiatrzyk,@bg.chun,@klueska,@p.rapacz] The Memory Manager KEP - final review addressed and merge/approve request  [\[KEP\]](https://github.com/kubernetes/enhancements/blob/43884638e460e2724ff4439a16684cfd68591988/keps/sig-node/1769-memory-manager/README.md) [\[discussion\]](https://github.com/kubernetes/enhancements/pull/1203)
* [@mitar] [https://github.com/kubernetes/enhancements/pull/2013/](https://github.com/kubernetes/enhancements/pull/2013/) Expose container imageID through Downward API
  * Issue: [https://github.com/kubernetes/enhancements/issues/2012](https://github.com/kubernetes/enhancements/issues/2012)
  * Pending review
* [@SergeyKanzhelev / @andrewsykim] Kubelet Exec Timeout KEP: [https://github.com/kubernetes/enhancements/pull/1973](https://github.com/kubernetes/enhancements/pull/1973) let’s approve
* [@SergeyKanzhelev] RuntimeClass GA plan [https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw](https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw)
* [@SergeyKanzhelev] Reminder
  * Sidecar meeting time: [https://groups.google.com/g/kubernetes-sig-node/c/w019G3R5VsQ/m/bbRDZTv5CAAJ](https://groups.google.com/g/kubernetes-sig-node/c/w019G3R5VsQ/m/bbRDZTv5CAAJ)
  * CRI alpha-&gt; beta 2nd meeting is tomorrow 11-12 PT [https://groups.google.com/g/kubernetes-sig-node/c/w019G3R5VsQ/m/bbRDZTv5CAAJ](https://groups.google.com/g/kubernetes-sig-node/c/w019G3R5VsQ/m/bbRDZTv5CAAJ) ([https://docs.google.com/document/d/1xMfiVUX4lF0G6Y_yQ4l-5-wrOneeJWcXY2NGVGuTdH8/edit#heading=h.4isgyltbe2x6](https://docs.google.com/document/d/1xMfiVUX4lF0G6Y_yQ4l-5-wrOneeJWcXY2NGVGuTdH8/edit#heading=h.4isgyltbe2x6) for notes)

## September 22nd, 2020

Total active pull requests: [179](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-7 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**16**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-09-15T17%3A00%3A00%2B0000..2020-09-22T16%3A49%3A14%2B0000)                                                     |Closed:      |[10](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-09-15T17%3A00%3A00%2B0000..2020-09-22T16%3A49%3A14%2B0000)|
|Updated:    |[**41**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A2020-09-15T17%3A00%3A00%2B0000..2020-09-22T16%3A49%3A14%2B0000+created%3A%3C2020-09-15T17%3A00%3A00%2B0000)|Merged:      |[13](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-09-15T17%3A00%3A00%2B0000..2020-09-22T16%3A49%3A14%2B0000)              |

potentially needs to be fished out of rotten:
    [https://github.com/kubernetes/kubernetes/pull/86071](https://github.com/kubernetes/kubernetes/pull/86071)
    [https://github.com/kubernetes/kubernetes/pull/88741](https://github.com/kubernetes/kubernetes/pull/88741)

* [@derekwayne] KEPs to track for enhancement freeze (10/6)
  * lets get everything enumerated this week [see items today]
  * review [SIG Node - Feature Health Check](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit)
* [@rata, @jirving] sidecar KEP
* [@mauriciovasquezbernal] user namespaces KEP
  * Minimal [KEP](https://github.com/kubernetes/enhancements/pull/1903) PR with summary and motivation (To be updated)
  * [Slides](https://docs.google.com/presentation/d/1Op6g0mw2sQ_PopT5O54WSxpgirPm6vBLtvKegx-1qwQ/edit?usp=sharing)
  * Discussion:
    * Move userNamespaceMode in PodSecurityContext?
    * image garbage collection in the scenario of a sidecar used in all pods when all pods use userNamespaceMode=Pod
    * Get feedback from gVisor and Katacontainer
* [@mitar] [https://github.com/kubernetes/kubernetes/issues/80346](https://github.com/kubernetes/kubernetes/issues/80346) Expose container imageID through Downward API
* [@renaudwastaken] Graduate Pod resources API to G.A in 1.20
  * [Reformat to KEP template + G.A plan Pull request](https://github.com/kubernetes/enhancements/pull/1865)
  * Feature is alpha since 1.13 and beta since 1.15, used in prod by many user
  * Work remaining:
    * Add a metric (pod_resources_requests_total)
    * Add some gRPC config options to ensure rate/resource limiting
    * Copy the beta gRPC API to a “v1 path” (and serve it on the unix socket)
* [@cez,@alukiano,@krzwiatrzyk,@bg.chun,@klueska] The Memory Manager KEP is ready for final review
* [@bobbypage] Node Shutdown KEP: [https://github.com/kubernetes/enhancements/pull/2001](https://github.com/kubernetes/enhancements/pull/2001)
* Container Notifier KEP [https://github.com/kubernetes/enhancements/pull/1995](https://github.com/kubernetes/enhancements/pull/1995)
* [@SergeyKanzhelev / @andrewsykim] Kubelet Exec Timeout KEP: [https://github.com/kubernetes/enhancements/pull/1973](https://github.com/kubernetes/enhancements/pull/1973) - should it include preStop “timeout” i.e. graceful termination? and let’s approve
* [@SergeyKanzhelev] RuntimeClass GA plan [https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw](https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw)
* [@SergeyKanzhelev] QQ: what’s the difference between node-kubelet-master and node-kubelet-conformance [https://github.com/kubernetes/test-infra/issues/18973](https://github.com/kubernetes/test-infra/issues/18973)
  * what is the NodeConformance tag for? What’s the difference from the SIG Architecture’s conformance tests?

## September 15th, 2020

Total active pull requests: [186](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-4 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**24**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-09-08T17%3A00%3A00%2B0000..2020-09-15T16%3A44%3A14%2B0000)                                                     |Closed:      |[12](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-09-08T17%3A00%3A00%2B0000..2020-09-15T16%3A44%3A14%2B0000)|
|Updated:    |[**41**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A2020-09-08T17%3A00%3A00%2B0000..2020-09-15T16%3A44%3A14%2B0000+created%3A%3C2020-09-08T17%3A00%3A00%2B0000)|Merged:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-09-08T17%3A00%3A00%2B0000..2020-09-15T16%3A44%3A14%2B0000)              |

* [@[ruiwen-zhao](https://github.com/ruiwen-zhao)] [https://github.com/kubernetes/k8s.io/pull/1194](https://github.com/kubernetes/k8s.io/pull/1194) (Nvidia GPU device plugin)
  * GoogleCloudPlatform/container-engine-accelerators has been releasing device plugins through [Kubernetes’ addons](https://github.com/kubernetes/kubernetes/blob/master/cluster/addons/device-plugins/nvidia-gpu/daemonset.yaml#L39)
  * After the [Vanity Domain Flip](https://github.com/kubernetes/k8s.io/blob/main/k8s.gcr.io/Vanity-Domain-Flip.md), the vanity domain (k8s.gcr.io) now points to {asia,eu,us}.gcr.io/k8s-artifacts-prod. The PR above is to create a repo in OSS community so that we can keep releasing the device plugin to the new domain.
* [@xing-yang, @yuxiangqian] ContainerNotifier
  * [enhancement issue](https://github.com/kubernetes/enhancements/issues/1977)
  * working on a KEP based on the doc [here](https://docs.google.com/document/d/1SWSlZoxY5zFjBKFKaATP07s3q02UenSp8R9-yRkCcwg/edit#)
* [@pjbgf]
  * [https://github.com/kubernetes/org/issues/2177](https://github.com/kubernetes/org/issues/2177) Renaming seccomp-operator to security-operator
  * [https://github.com/kubernetes/enhancements/pull/1444](https://github.com/kubernetes/enhancements/pull/1444) KEP AppArmor GA
* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/6928d9f33261d0a72f639b3eddc39a32b4055fad/keps/sig-node/0753-sidecarcontainers.md))
  * [open PR](https://github.com/kubernetes/enhancements/pull/1980/files) addressing all concerns, using suggestions stated previously
* [@ambguo] Windows privileged container support ([KEP](https://docs.google.com/document/d/12EUtMdWFxhTCfFrqhlBGWV70MkZZPOgxw0X-LTR0VAo/edit?usp=sharing))
  * [enhancement issue](https://github.com/kubernetes/enhancements/issues/1981) for tracking
  * Looking for sig-node reviewer
* [@dims, @derekwaynecarr] Removing dockershim from kubelet ([KEP](https://github.com/kubernetes/enhancements/pull/1985))
* [@derekwaynecarr, @mrunalp] CRI alpha-&gt;beta transition discussion
  * propose to have a few meetings week of 9/21 to identify gap/plan, will send note to mailing list with proposed times.
* [@SergeyKanzhelev] QQ: what’s the difference between node-kubelet-master and node-kubelet-conformance [https://github.com/kubernetes/test-infra/issues/18973](https://github.com/kubernetes/test-infra/issues/18973)
  * what is the NodeConformance tag for? What’s the difference from the SIG Architecture’s conformance tests?
* [@SergeyKanzhelev] RuntimeClass GA plan [https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw](https://docs.google.com/document/d/17nROj6ayPsUpx09mhLrzOkRLgo-8sn6Vgfyy-gmyuo4/edit#heading=h.3qoppsm7jdvw)
* [@SergeyKanzhelev / @andrewsykim] [https://github.com/kubernetes/enhancements/pull/1973](https://github.com/kubernetes/enhancements/pull/1973) - should it include preStop “timeout” i.e. graceful termination? and let’s approve
* [@klueska] The Memory Manager KEP is ready for final review

## September 8th, 2020

Total active pull requests: [188](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-22 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**23**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-09-01T17%3A00%3A00%2B0000..2020-09-08T16%3A40%3A45%2B0000)                                                     |Closed:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-09-01T17%3A00%3A00%2B0000..2020-09-08T16%3A40%3A45%2B0000)|
|Updated:    |[**50**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A2020-09-01T17%3A00%3A00%2B0000..2020-09-08T16%3A40%3A45%2B0000+created%3A%3C2020-09-01T17%3A00%3A00%2B0000)|Merged:      |[29](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-09-01T17%3A00%3A00%2B0000..2020-09-08T16%3A40%3A45%2B0000)              |

potential to fish out of rotten:

* [Removed redundant hardcoded default for cpu-manager-policy flag · Issue #88825 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/pull/88825),
* [Add unittest for Refactor kubelet convertStatusToAPIStatus · Issue #87557 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/pull/87557)

* [@andrewsykim] on timeout
  * [https://docs.google.com/document/d/1wHoZ514Ji9qhzovDQ9u8m1T-UinAYiOBNMSWG7akZxg/edit?usp=sharing](https://docs.google.com/document/d/1wHoZ514Ji9qhzovDQ9u8m1T-UinAYiOBNMSWG7akZxg/edit?usp=sharing)
  * **Action item**: short KEP on introducing the feature flag. Overall proposal looks OK and everybody want to move it forward.
  * Also discussed: expedite deprecation of dockershim. Try not to invest too much there with this work.
* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/6928d9f33261d0a72f639b3eddc39a32b4055fad/keps/sig-node/0753-sidecarcontainers.md))
  * Any remaining major concern?
  * To simplify, as was suggested in the comments, I was planning to open a PR:
    * Removing the proposed TerminationHook and FatalToPod field
    * Update proposal to not change the pod phase (as suggested [here](https://github.com/kubernetes/enhancements/blob/6928d9f33261d0a72f639b3eddc39a32b4055fad/keps/sig-node/0753-sidecarcontainers.md#revisit-if-we-want-to-modify-the-podphase))
  * How to continue the discussion on some specific issues?
    * [This callout](https://github.com/kubernetes/enhancements/blob/6928d9f33261d0a72f639b3eddc39a32b4055fad/keps/sig-node/0753-sidecarcontainers.md#killing-pods-take-3x-the-time) on time to kill a pod. No suggestion alternative here
      * David Porter suggested: Maybe similar to the timeout issue raised be andrew? This was not respected and make a release note saying it is respected.
    * [This callout](https://github.com/kubernetes/enhancements/blob/6928d9f33261d0a72f639b3eddc39a32b4055fad/keps/sig-node/0753-sidecarcontainers.md#how-to-split-the-shutdown-time-to-kill-different-types-of-containers) on how to split the time. Any concerns on moving forward with the suggestion?
  * I’d like to aim this for Kubernetes 1.20, if we all feel comfortable with it
* [@derekwaynecarr] Sizable memory backed volumes ([PR](https://github.com/kubernetes/kubernetes/pull/94444))
* [@fromani] just requesting another round of review for podresources API extensions t
  * [https://github.com/kubernetes/enhancements/pull/1884](https://github.com/kubernetes/enhancements/pull/1884) (added GetAvailableResources API)
  * [https://github.com/kubernetes/enhancements/pull/1926](https://github.com/kubernetes/enhancements/pull/1926) (add Watch API)
* [@kad] on resource related API improvements: [link to WIP notes](https://github.com/container-orchestrated-devices/resource-management-improvements-wg)

## September 1, 2020

Total active pull requests: [209](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-33 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**11**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-08-25T17%3A00%3A00%2B0000..2020-09-01T16%3A48%3A02%2B0000)                                                     |Closed:      |[6](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-08-25T17%3A00%3A00%2B0000..2020-09-01T16%3A48%3A02%2B0000)|
|Updated:    |[**65**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A2020-08-25T17%3A00%3A00%2B0000..2020-09-01T16%3A48%3A02%2B0000+created%3A%3C2020-08-25T17%3A00%3A00%2B0000)|Merged:      |[38](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-08-25T17%3A00%3A00%2B0000..2020-09-01T16%3A48%3A02%2B0000)             |

* [@SergeyKanzhelev/@andrewsykim] Timeouts in exec probe
  * Latest PR (@andrewsykim): [https://github.com/kubernetes/kubernetes/pull/94115](https://github.com/kubernetes/kubernetes/pull/94115/)
    * [andrewsykim] summary of existing regressions w/ exec timeouts: [https://github.com/kubernetes/kubernetes/pull/94115/#issuecomment-678681860](https://github.com/kubernetes/kubernetes/pull/94115/#issuecomment-678681860)
  * @louyihua (Jan 28, 2018): [https://github.com/kubernetes/kubernetes/pull/58925/](https://github.com/kubernetes/kubernetes/pull/58925/)
        (replaces [https://github.com/kubernetes/kubernetes/pull/58510](https://github.com/kubernetes/kubernetes/pull/58510))
  * @tnqn (Jun 24, 2020): [https://github.com/kubernetes/kubernetes/pull/92465/](https://github.com/kubernetes/kubernetes/pull/92465/)
  * tedyu (Jan 16, 2020) [https://github.com/kubernetes/kubernetes/pull/87281/](https://github.com/kubernetes/kubernetes/pull/87281/)

        From Alexander Kanevskiy to Everyone:  10:08 AM
        Let’s deprecate dockershim :) it is long time overdue
        From Me to Everyone:  10:09 AM
        =) containerD has the similar issue
        but I think I agree in principle
        From Alexander Kanevskiy to Everyone:  10:12 AM
        actually…. “runc exec” doesn’t have way to specify timeout
        so, OCI compatible runtime will execute something until it return
        From michael crosby to Everyone:  10:15 AM
        If you want runc exec to timeout, you need to use context.Context with a timeout then containerd should handle it when that context is canceled
        We use exec.CommandContext for all calls to external binarie
        From Alexander Kanevskiy to Everyone:  10:19 AM
        so containerd will be the one who kills “runc exec” process…. in theory that is ok, but sometimes might be not always reliable. would it make sense to integrate “timeout” functionality on the lower level (OCI runtime spec?) so it will be reliable cleaning up processes inside container
        From Me to Everyone:  10:20 AM
        Yes, I think this is the desire. Basically first question we wanted to answer whether we need to support timeouts on exec at all. Now - whatever mechanism we will have - how to introduce it in new versioin without affecting payloads. And finally we can discuss whether to start with Andrew;s PR
        From michael crosby to Everyone:  10:24 AM
        I wouldn’t think timeout in the OCI runtime spec would make sense.  The cancel of a context should unroll things correctly as exec.CommandContext does a SIGKILL
       
       
        [Dawn] maybe timeout would not lead to container restart/kill? Ideally timeout value shouldn’t be 1 second, it should be like a catch all bigger value.
       
        [michael Crosdby] let’s not have explicit default timeout at all?
       
        [Andrew] is changing default a breaking change? [Dawn] Yes, definitely
        When readiness probe relying on default of 1 second - extending it might affect user payload

* [@bmcfall] Status of exposing hugepages in pod
  * [add support for hugepages in downward API #86102](https://github.com/kubernetes/kubernetes/pull/86102) - Closed due to inactivity
  * @kad: as a workaround for hugepages specifically, inside container it is possible to read /sys/fs/cgroup/hugetlb/\*limit_in_byte
    * @bmcfall: Thanks @kad - Followed up on slack.

* [@bmcfall]] Code freezed for 1.20?
  * Release deadlines [https://github.com/kubernetes/sig-release/tree/master/releases/release-1.20#tldr](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.20#tldr)
  * Nov 12th for bug
  * Oct 6th for feature

## August 25, 2020

Total active pull requests: [238](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+20 from the last meeting (2 weeks))

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**35**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-08-11T17%3A00%3A00%2B0000..2020-08-25T16%3A46%3A18%2B0000)                                                     |Closed:      |[11](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-08-11T17%3A00%3A00%2B0000..2020-08-25T16%3A46%3A18%2B0000)|
|Updated:    |[**78**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A2020-08-11T17%3A00%3A00%2B0000..2020-08-25T16%3A46%3A18%2B0000+created%3A%3C2020-08-11T17%3A00%3A00%2B0000)|Merged:      |[4](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-08-11T17%3A00%3A00%2B0000..2020-08-25T16%3A46%3A18%2B0000)               |

* [@dchen1107] Followup: SIG Windows &amp; Microsoft representative at SIG Node: Mark Rossetti
* [@derekwaynecarr] [Feature Health Check](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#)
  * state of features sig has in flight, see if we can discuss path fwd.
* [@derekwaynecarr] Promote \*PidLimits to GA
  * see: [https://github.com/kubernetes/kubernetes/pull/94140](https://github.com/kubernetes/kubernetes/pull/94140)
* [@sergeykanzhelev] Promote RuntimeClass to GA. What extra feedback do we need to get?
  * follow up with Mark Rossetti about RuntimeClass on windows - the question is what changes are needed and whether this is blocking GA
  * Derek: what does it mean to test it in conformance? Let’s make a review of test
  * Mrunal: [https://github.com/openshift/kata-operator](https://github.com/openshift/kata-operator) uses RuntimeClass to install kata on OpenShift with CRI-O.
  * Dawn: decoupling from PodOverhead since kata usecase is not blocked by the feature.
* [@sergeykanzhelev] Plan for reopening merges for v1.20: [read the plan](https://docs.google.com/document/d/1yI7o7R2bOB7QtkHUFMHMb5XObiJmzIAnJPuxhdAX-CA/edit) (membership in kubernetes-dev is needed to access)
* [@sergeykanzhelev] SIG Node CI group:
    Group [charter document](https://docs.google.com/document/d/1yS-XoUl6GjZdjrwxInEZVHhxxLXlTIX2CeWOARmD8tY/edit#heading=h.te6sgum6s8uf).

    Join the group: [kubernetes-sig-node-test-failures](https://groups.google.com/forum/#!forum/kubernetes-sig-node-test-failures) for the meeting invite and updates.
    E-mail: kubernetes-sig-node-test-failures@googlegroups.com
* [@dchen1107] Previous SIG Node E2E test re-category proposal done by yjhong@:
  * [https://docs.google.com/document/d/1BdNVUGtYO6NDx10x_fueRh_DLT-SVdlPC_SsXjYCHOE/edit?usp=sharing](https://docs.google.com/document/d/1BdNVUGtYO6NDx10x_fueRh_DLT-SVdlPC_SsXjYCHOE/edit?usp=sharing)
  * Some are done, some are still pending. Have some background contexts.
* [@derekwaynecarr] to follow up on ServiceMesh in RedHat and schedule a meeting to discuss it with interested parties.Still has some reservations regarding scope and target customer.

## August 18, 2020

Cancelled due to the conflict with KubeCon.

## August 11, 2020

Total active pull requests: [219](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+0 from last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**11**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A%3E%3D2020-08-04T17%3A00%3A00%2B0000)                                                     |Closed:      |[3](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A%3E%3D2020-08-04T17%3A00%3A00%2B0000)|
|Updated:    |[**47**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+updated%3A%3E%3D2020-08-04T17%3A00%3A00%2B0000+created%3A%3C2020-08-04T17%3A00%3A00%2B0000)|Merged:      |[8](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A%3E%3D2020-08-04T17%3A00%3A00%2B0000)              |

* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/0753-sidecarcontainers.md))
  * Friendly ping on [open PR](https://github.com/kubernetes/enhancements/pull/1913)
* [@swsehgal, @aperevalov @fromani] Topology Aware Scheduler
  * [Status Update](https://docs.google.com/presentation/d/18-ZQWK7P4SL3niUxmNr2Y5Ir6VMmNxsyik23am76Kgs/edit#slide=id.g8fa4b3ed81_0_1490) and [Demo](https://asciinema.org/a/3F2c4wZRat3ybWLoS8ifCE6HG)
* [@sjenning]
  * discuss solution for [https://github.com/kubernetes/kubernetes/issues/72881#issuecomment-671389165](https://github.com/kubernetes/kubernetes/issues/72881#issuecomment-671389165)
* [@stewartbutler] Recap on node readiness gateway: [https://github.com/kubernetes/enhancements/pull/1003](https://github.com/kubernetes/enhancements/pull/1003)
  * Istio race condition description and original mitigation: <https://tinyurl.com/istio-cni-race>
* [@SergeyKanzhelev] Test failure group - meeting scheduling question. Would it be ok to keep using meet instead of zoom so Dawn and Derek are not needed to start the meeting? -- Done!

## August 4, 2020

* [@dashpole] PR health from devstat
  * Summary: There was a drop in reviewers, and a spike in neglected PRs early this year, but that has since been reduced to mostly normal levels.
  * [Awaiting](https://k8s.devstats.cncf.io/d/70/awaiting-prs-by-sig?orgId=1&var-sigs=%22node%22):
    * Spike in Jan - March, but back down to the same as last year.
    * Increasing numbers of PRs open for more than a year
  * [Inactive](https://k8s.devstats.cncf.io/d/72/inactive-prs-by-sig?orgId=1&var-sigs=%22node%22&from=now-1y&to=now): Spike in Jan - March, but down slightly in last year.
  * [Throughput and number of reviewers](https://k8s.devstats.cncf.io/d/33/pr-workload-per-sig-chart?orgId=1&var-sigs=%22node%22&from=now-1y&to=now):
    * Drop in number of reviewers in Jan - Feb, otherwise constant
    * Constant throughput
  * [PRs Currently Open](https://k8s.devstats.cncf.io/d/71/prs-labels-by-sig?orgId=1&var-sig_name=node&var-label_name=All%20labels%20combined&from=1548540596095&to=1611785396095):
    * Increasing slowly to 335 in Jan, then slow decrease to 242 today
  * [Rate of PRs Opened](https://k8s.devstats.cncf.io/d/69/prs-opened-by-sig?orgId=1&from=now-1y&to=now&var-period=w&var-sigs=%22node%22): Relatively constant
  * [PRs open, actionable, needing approval](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+is%3Aopen+label%3Asig%2Fnode+label%3Algtm+-label%3Aapproved+-label%3Ado-not-merge%2Fhold+-label%3Ado-not-merge%2Fwork-in-progress+-label%3Aneeds-rebase+-label%3Ado-not-merge%2Frelease-note-label-needed+) in k/k (\~25)
  * [PRs open, needing lgtm](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+is%3Aopen+label%3Asig%2Fnode+-label%3Algtm+-label%3Ado-not-merge%2Fhold+-label%3Ado-not-merge%2Fwork-in-progress+-label%3Aneeds-rebase) (\~100)
  * Unactionable PRs (hold, rebase-required, WIP, release-note-label-needed):
    * 242 - 125 = 117
* [Sergey / Victor] Request to form SIG Node CI Subproject
  * [SIG Node CI Group Charter](https://docs.google.com/document/d/1yS-XoUl6GjZdjrwxInEZVHhxxLXlTIX2CeWOARmD8tY/edit#)
  * [7/31 Meeting Notes](https://docs.google.com/document/d/1I0Exz-VBvs1nDu5yQynhFbEQzqCs8chedgL8YC-yVhE/edit)
  * Revive SIG-Node performance dashboard
    * [https://github.com/kubernetes-retired/contrib/tree/master/node-perf-dash](https://github.com/kubernetes-retired/contrib/tree/master/node-perf-dash)
    * [https://node-perf-dash.k8s.io/#/builds](https://node-perf-dash.k8s.io/#/builds)
  * Doodle to figure out best meeting time/cadence
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md)
  * Discuss [node checkpointing approach](https://drive.google.com/file/d/12wKXl6_y0tDqOhTEXPIK7gN_qsfJzjJ9/view?usp=sharing) instead of API change to PodSpec

## July 28, 2020

* [@mauriciovasquezbernal, @alban, @rata] User namespaces support
  * Old [[design-proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/node-usernamespace-remapping.md)], old [[issue](https://github.com/kubernetes/enhancements/issues/127)] and old [[PR](https://github.com/kubernetes/kubernetes/pull/64005)]
  * New minimal [KEP](https://github.com/kubernetes/enhancements/pull/1903) PR with summary and motivation
* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/0753-sidecarcontainers.md))
  * Callouts PR coming. It’s quite long, err on the side of being very clear.
    * Regarding the length added, is only needed temporarily. Once we agree on alternatives, we can remove the callouts and it will shrink quite a lot :)
    * Unless it is a problem, will open the PR like thi
  * Rodrigo(@rata) will be away 2-3 days next week
* [@jameslaverack] “Major Themes” for Kubernetes 1.19 release.
* [https://docs.google.com/spreadsheets/d/1VW5_Eq8MzswfDi9xEvfYyP8edF_Ny7MBANIsJXT3VGw/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1VW5_Eq8MzswfDi9xEvfYyP8edF_Ny7MBANIsJXT3VGw/edit?usp=sharing)
        [approval from any sig to merge change in .golint_failures · Issue #92617 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/pull/92617)
* [@dchen1107] [https://github.com/orgs/kubernetes/projects/43](https://github.com/orgs/kubernetes/projects/43) From Victor
* [@nliao] SIG Node Test/CI subproject/group initial discussion meeting scheduled from 9-10am Friday (PST).
* [@dchen1107] From: [**alejandrox1**](https://app.slack.com/team/U6AS37R50)cross posting for visibility of anyone interested in CI  [https://kubernetes.slack.com/archives/C2C40FMNF/p1595884044265400](https://kubernetes.slack.com/archives/C2C40FMNF/p1595884044265400)

## July 21, 2020

* [@crosbymichael] Node Resource Interface:  Building an extensible API for managing node resources. See: [https://github.com/containerd/containerd/pull/4411](https://github.com/containerd/containerd/pull/4411)
* [@hh] API endpoint deleteCoreV1CollectionNamespacedPod needs a conformance test
  * [https://apisnoop.cncf.io/1.19.0/stable/core/deleteCoreV1CollectionNamespacedPod](https://apisnoop.cncf.io/1.19.0/stable/core/deleteCoreV1CollectionNamespacedPod) (e2e test exists, no conformance test yet)
  * Current test slightly flakey: [https://github.com/kubernetes/kubernetes/pull/93086](https://github.com/kubernetes/kubernetes/pull/93086)
  * Could use some help understanding why so we can increase conformance coverage (required for GA endpoints)
* [@renaudwastaken] Disable Accelerator Metrics exception
  * Marking the [KEP implementable is the last step that needs approval before the exception can move forward](https://github.com/kubernetes/enhancements/pull/1896)
  * Kubernetes PR is here: [https://github.com/kubernetes/kubernetes/pull/91930](https://github.com/kubernetes/kubernetes/pull/91930)
  * Exception request is here: [https://groups.google.com/forum/#!topic/kubernetes-sig-node/_4Eb5I7CY18](https://groups.google.com/forum/#!topic/kubernetes-sig-node/_4Eb5I7CY18)
  * This topic just needs an lgtm/approve from sig-node lead
* [alejandrox1] CI subproject [https://groups.google.com/forum/?oldui=1#!topic/kubernetes-sig-node/Ur8RnPrR4V0](https://groups.google.com/forum/?oldui=1#!topic/kubernetes-sig-node/Ur8RnPrR4V0)
* [@AlexeyPerevalov @swatisehgal] Topology Aware Scheduling
  * [KEP](https://github.com/kubernetes/enhancements/pull/1884) for extending podresource
* [@sjenning] New e2e test that validates terminationGracePeriodSeconds exposed bug in kubelet causing \~5% flake rate in e2e-node.  Would appreciate review.
  * [https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/mirror_pod_grace_period_test.go](https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/mirror_pod_grace_period_test.go)
  * [https://github.com/kubernetes/kubernetes/issues/93210](https://github.com/kubernetes/kubernetes/issues/93210)
  * proposed PR [https://github.com/kubernetes/kubernetes/pull/93261](https://github.com/kubernetes/kubernetes/pull/93261)
  * more complex, but more granular cache invalidation and thus more performant [https://github.com/liggitt/kubernetes/commit/da1f636514251fa84d89471d1f66c9fb20909d11](https://github.com/liggitt/kubernetes/commit/da1f636514251fa84d89471d1f66c9fb20909d11)

## July 14, 2020

* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/0753-sidecarcontainers.md))
  * Just a quick update
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) - how do we move forward?
  * emptyDir.memory: Abort attempts to set limit &lt; usage reasonable?
    * Ref: [discussion](https://github.com/kubernetes/enhancements/pull/1883/files#r454113287) with @thockin
  * As discussed in today’s meeting - [this conversation](https://github.com/kubernetes/enhancements/pull/1883/files#r453339451) has link to previous version of KEP where we had Status.ResorurcesAllocated, and its challenges with scheduler functioning vs Spec.ResourcesAllocated (or ResourcesToAllocate for better naming)
* [@yash97][Sending events Distributed manner](https://github.com/kubernetes/enhancements/pull/1891)--KEP discussion.
* [@mauriciovasquezbernal, @alban, @rata] User namespaces support
  * Old [[design-proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/node-usernamespace-remapping.md)], old [[issue](https://github.com/kubernetes/enhancements/issues/127)] and old [[PR](https://github.com/kubernetes/kubernetes/pull/64005)]
  * Starting to write KEP, looking for potential reviewers.
* [kad] SIG Node Resource Management Forum: new meeting time? calendar invites?
* [@bobbypage] Cherrypick for kubelet reporting incorrect status: [https://github.com/kubernetes/kubernetes/pull/93041](https://github.com/kubernetes/kubernetes/pull/93041)

## July 07, 2020

* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - Update &amp; next step
  * Late-stage API review and questions about KEP - not making 1.19
  * @thockin is revisiting completed API review decisions:
    * Is ResourcesAllocated really needed? Why not do local checkpointing?
    * Subresource to set resources &amp; resourcesAllocated
    * How do we handle runtimes that *may* sometimes require container restart in order to resize? How does that work with restartPolicy=Never
    * Should Restart be the default resize policy?
    * Should RuntimeClass [should be allowed to disable in-place update](https://github.com/kubernetes/enhancements/pull/1883/files#r448754705) so users get synchronous error. sounds reasonable to me if runtime doesn’t support this - thoughts?
    * Are RestartOnGrow, RestartOnShrink resize policies a good idea? (Doesn’t have to be added today, but quite easy to add if it feels useful)
* [mrunalp, nalin] /dev/fuse in pods [https://github.com/kubernetes/kubernetes/pull/79925](https://github.com/kubernetes/kubernetes/pull/79925)
  * Nalin and Mrunal to come up with use cases and design
* [mrunalp, zvonkok] Rlimits support - [https://github.com/kubernetes/kubernetes/issues/3595](https://github.com/kubernetes/kubernetes/issues/3595)
* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * Switch KEP to provisionable state: [https://github.com/kubernetes/enhancements/pull/1874](https://github.com/kubernetes/enhancements/pull/1874)
* [@AlexeyPerevalov @swatisehgal] Topology Aware Scheduling
  * Topology Exporter Daemon in kubernetes-sig
* [harche/mrunalp] Fedora images for node testing
* [@yash97]Informal Kep:Sending Events to user directly to kubelet in distributed ways instead to receive it from api server.
* [alejandrox1] [https://github.com/kubernetes/kubernetes/pull/80917](https://github.com/kubernetes/kubernetes/pull/80917) needs review from SIG Node approver

## June 30, 2020

* [Balaji] Add support for disabling /logs endpoint
  * [https://github.com/kubernetes/kubernetes/pull/87273#issuecomment-649652855](https://github.com/kubernetes/kubernetes/pull/87273#issuecomment-649652855) would like to get some opinion on which option to pick. Handling backward compatibility complicated some of the things, so I’m thinking of going with enableSystemLogHandler requiring enableDebuggingHandler to be set to true as well. Looking for some suggestions and feedback here.
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - Update &amp; next step
  * Feature code complete with unit tests, reviewed by @dashpole.
    * PR: [https://github.com/kubernetes/kubernetes/pull/92127](https://github.com/kubernetes/kubernetes/pull/92127)
  * Basic e2e test framework done, adding test-cases promised for alpha.
  * Need reviews @liggitt, @thockin, @ahg-g, and Derek, Dawn, sig-testing.
  * Identify process for informing CRI changes to runtime folks.
* [@mukesh-dua/@guptavishal7982] CPU Reservation for Infrastructure Services Deployed in Kubernete
  * Would like to introduce the enhancement and its requirements.
* [@rata/Kinvolk, @alban/Kinvolk, @jirving], sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * PR to update to provisional state and depend on kubelet node shutdown [in place](https://github.com/kubernetes/enhancements/pull/1874)
  * Reviewers/approvers?
    * SergeyKanzhelev
  * PR coming \~1 or 2 weeks (hopefully sooner) with design “callout” discussed in the previous meeting, alternatives and history of all decisions made in the past (WIP)
    * Plan to update summary, motivation, etc. too.
* [@hasheddan] seccomp-operator: [https://github.com/kubernetes/org/issues/1873](https://github.com/kubernetes/org/issues/1873)

## June 23, 2020

* [@harche]: GCP accounts for hosting custom images for node test
* [dawnchen] Introducing PoC from google for node images and critool/critest
  * Karan [GKE] and Sergey Kanzhelev
* [mrunal]: CRI missing detach for exec. See [https://github.com/kubernetes/kubernetes/issues/](https://github.com/kubernetes/kubernetes/issues/92057)[92057](https://github.com/kubernetes/kubernetes/issues/92057).
* [@rata, @alban](Kinvolk) sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * How to raise concerns about KEP and have review/feedback from sig-node? To address all concerns at the KEP level
    * Yes, PR to KEP calling them out
  * Add KEP use cases?
    * Yes, PR to KEP
  * Should we make a PR to the KEP to depend on the kubelet graceful shutdown?
    * Yes, PR to KEP to set clear expectations to everyone
* [@alban, @mauriciovasquezbernal, @rata](Kinvolk) User namespaces [[design-proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/node-usernamespace-remapping.md)] [[issue](https://github.com/kubernetes/enhancements/issues/127)] [[PR](https://github.com/kubernetes/kubernetes/pull/64005)]
  * Reusing the upstream PR above, we’re working on a PoC based on Kubernetes 1.17 + containerd/cri 1.3
    * [Kubernetes PoC branch](https://github.com/kinvolk/kubernetes/pull/3)
    * [containerd/cri PoC branch](https://github.com/kinvolk/containerd-cri/pull/1) ([issue](https://github.com/containerd/cri/issues/790))
    * runc bug ([runc#2484](https://github.com/opencontainers/runc/issues/2484))
  * We will make a KEP update PR once we have a better understanding
    * History of Node wide mapping?
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - Update &amp; follow-up
  * Feature code complete with unit test
  * Upstream k/k pr created: [https://github.com/kubernetes/kubernetes/pull/92127](https://github.com/kubernetes/kubernetes/pull/92127)
    * Review items addressed.
    * Working E2E tests with @wangchen615
* [@AlexeyPerevalov] Topology aware scheduling
  * Current status: [KEP](https://github.com/kubernetes/enhancements/pull/1870) for resource with topology provisioning daemon and [KEP](https://github.com/kubernetes/enhancements/pull/1858) for scheduler plugin
  * Request of feedback for [New ideas for topology aware scheduling](https://docs.google.com/presentation/d/1SC6SYkK1YvfyW0HXkic3xA75C6s9B3PCwTACTa_tl90/edit#slide=id.p)
* [@mythi] [non-root device usage, follow-up](https://docs.google.com/document/d/1SX4o71AIIrJAzbGJEIfhT2NxQpPvLlrHiZj5uUawWxk/edit)

## June 16, 2020

* [@renaudwastaken] [Disabling GPU metrics provided by the Kubelet/cadvisor](https://github.com/kubernetes/kubernetes/pull/91930)
  * Awaiting review, specific agenda items were discussed in the previous meeting
* [@pjbgf, @hasheddan] seccomp-operator [https://docs.google.com/presentation/d/1tkz8Zgd4nzrTPaR1jDBkJgBu3jcle6jy-4Vd9F1AyCyKE/edit#slide=id.g88d31c8bbc_0_63](https://docs.google.com/presentation/d/1tkz8Zgd4nzrTPaR1jDBkJgBu3jc6jy-4Vd9F1AyCyKE/edit#slide=id.g88d31c8bbc_0_63)
* [@dashpole] NamespaceUID in CRI log path: [https://github.com/kubernetes/kubernetes/issues/58638](https://github.com/kubernetes/kubernetes/issues/58638)
* [@renaudwastaken] [Make the pod metrics API G.A](https://github.com/kubernetes/kubernetes/pull/92165)
* [@AlexeyPerevalov, @swatisehgal] Repository for daemon responsible for exposing node resources with NUMA topology
  * [https://github.com/kubernetes/enhancements/pull/1858](https://github.com/kubernetes/enhancements/pull/1858)
* [@mythi, @kad]: [non-root user containers and devices access](https://docs.google.com/document/d/1SX4o71AIIrJAzbGJEIfhT2NxQpPvLlrHiZj5uUawWxk)
* [@vpickard, @harche]: fedora 31 images for crio testing in gcp
  * TODO (dawnchen@): figure out POC from Google side.
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - Status Update
  * All planned dev tasks for alpha complete
  * Upstream k/k pr created: [https://github.com/kubernetes/kubernetes/pull/92127](https://github.com/kubernetes/kubernetes/pull/92127)
  * Started working on E2E tests with @wangchen615

## June 9, 2020

* [@rata, @alban](Kinvolk) sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * Synced with kubelet graceful shutdown “working group”
  * Found some edge cases on termination (pod=nil, overall time spent to kill a pod) and we have some ideas to improve it
  * pod.DeletionGracePeriodSeconds is not mandatory? [Code example](https://github.com/kubernetes/kubernetes/blob/e63fb9a597bfbf6f3d454489e4fb49b40ad8c48f/pkg/kubelet/kuberuntime/kuberuntime_container.go#L604-L610)
    * Added in [https://github.com/kubernetes/kubernetes/pull/31322](https://github.com/kubernetes/kubernetes/pull/31322)
* [@alban](Kinvolk) User namespaces [[design-proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/node-usernamespace-remapping.md)] [[issue](https://github.com/kubernetes/enhancements/issues/127)] [[PR](https://github.com/kubernetes/kubernetes/pull/64005)]
  * I am planning to work on this in the near future
  * There were some challenges on upgrades and some areas under-defined in the KEP
  * Set of folks (Kinvolk and other interested parties) can maybe update the KEP? It predates ephemeral containers, PID ns sharing, many clarify what it means to different containers types? (what does it mean for ephemeral containers or other cases?)
  * see vikas pr from 2018 that got far:
    * [https://github.com/kubernetes/kubernetes/pull/64005](https://github.com/kubernetes/kubernetes/pull/64005)
    * testing this is important and was never sufficiently captured.
* [kmala] [https://github.com/kubernetes/kubernetes/pull/89667](https://github.com/kubernetes/kubernetes/pull/89667)
  * [https://github.com/kubernetes/kubernetes/issues/91316](https://github.com/kubernetes/kubernetes/issues/91316)
* [vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - status update
  * Core implementation initial review done - thanks @dashpole!
  * First-cut implementation [ready](https://github.com/vinaykul/kubernetes/pull/1).
  * Resource-quota, limit-ranger, e2e tests are next.
* [@renaudwastaken] [Disabling GPU metrics provided by the Kubelet/cadvisor](https://github.com/kubernetes/kubernetes/pull/91930)
  * Discussion about enabling this by default in the future (note: k8s deprecation policy is 1+ year)
  * TLDR: Consensus on deprecating the GPU metrics, as a sig we have agreed on a way to collect metrics from out of tree. Need to figure out deprecation.
  * Dawn: Make the announcement through release notes and other channel
  * Derek: Can we tie this with cadvisor vendoring in some way?
  * David: We should tie the deprecation of GPU metrics with the summary API
  * Dawn: Deprecation of the summary API started 3+ years ago (even before the CRI), could this be a baby step for the cleanup?
  * Note: Do not add this as another CLI flag but only a config flag
* [@renaudwastaken] Cherry pick metrics bug fix to [1.16](https://github.com/kubernetes/kubernetes/pull/90070) and [1.17](https://github.com/kubernetes/kubernetes/pull/90071)
* [vpickard] [sig-node testing enhancements](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit?usp=sharing) update
  * Great efforts from team to understand and fix failing test
    * hey, we made [an entire testgrid tab blue](https://testgrid.k8s.io/sig-node-cri#Summary) with [test-infra#17831](https://github.com/kubernetes/test-infra/pull/17831)!
  * Created github project board to help manage issues/PR
    * experimenting with board, not public yet
  * 10 PRs complete/merged - doc updates, COS image fixes, test config, fail on missing image
  * 13 PRs in progress - more doc updates, more image cleanup, benchmark test
  * benchmark tests failing - OOM - # pods lowered from 105 - 90
    * containerd mem usage seems to have increased
    * opened issue [https://github.com/kubernetes/test-infra/issues/17853](https://github.com/kubernetes/test-infra/issues/17853) to track root cause
    * wishlist and todos for density tests (existing issues may exist, but rotten and closed automatically)
  * Considering adding additional [images](https://cloud.google.com/compute/docs/images/os-details) once tests are stable with current images (COS and ubuntu)
    * one-monthish failing test policy for supporting image
    * RHEL uploading results separately
* [vpickard] [sig-node-resource management](https://docs.google.com/document/d/15G9f0yoP8Gy5Agx-uCqo0CUl67EjZ-FgAk-KbI-lakw/edit?usp=sharing) forum
  * Is it possible to move this mtg to 8 am PDT so folks from US west coast can attend?
  * [https://github.com/kubernetes/enhancements/pull/1121](https://github.com/kubernetes/enhancements/pull/1121) approved, will likely not be implemented until 1.20
  * [https://github.com/kubernetes/enhancements/pull/1752](https://github.com/kubernetes/enhancements/pull/1752) pod level alignment for resource
    * Alex to take a final review of latest comment
  * no topics for Thursday, June 11 (Holiday in Poland, topic moved to 6/18/2020). Will cancel mtg tomorrow if no topic
  * Next scheduled mtg Thursday, June 18.
    * 5G deployment scenarios for pod level resource alignment

## June 2, 2020

Meeting is [canceled](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-node/Uxv-4_1n_Jo/giLeOtwlAwAJ).

## May 26, 2020

* [@rata, @alban](Kinvolk) sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * terminationGracePeriodSecond
  * multi-level dependency sidecar
  * Retro before in the year (is there a video recording?)
  * Concerns:
    * Moving to alpha without understanding termination sequence
    * Losing data during a node shutdown sequence (Kubernetes running in a train)
    * Device access, CPU, Memory policy
    * More complexity: debugContainer
* [tedyu] [https://github.com/kubernetes/kubernetes/pull/91211](https://github.com/kubernetes/kubernetes/pull/91211)
    Remove excess log
    [https://github.com/kubernetes/kubernetes/issues/90999](https://github.com/kubernetes/kubernetes/issues/90999)
    Give static pod deletion grace period
    [https://github.com/kubernetes/kubernetes/pull/91453](https://github.com/kubernetes/kubernetes/pull/91453)
* [vinaykul] In-Place Pod Vertical Scaling for v1.19 - update
  * [Kubelet-CRI KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191025-kubelet-container-resources-cri-api-changes.md) initial code ready for [review](https://github.com/vinaykul/kubernetes/pull/1/commits/5879787bbfee6951afbe39f0cc02dff409a2339a).
  * First-cut implementation in progress, \~ETA end of 1st week of June.
  * API-only code changes with review feedback: [5126b9e1](https://github.com/vinaykul/kubernetes/pull/1/commits/5126b9e1e93d2530ac820981c93986d64e46faa3)
* [mrunalp] CRI errors - [https://github.com/kubernetes/kubernetes/pull/91273](https://github.com/kubernetes/kubernetes/pull/91273/files)
  * Ready for review
* [vpickard] sig-node testing enhancements update ([doc](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit?usp=sharing))
  * rescheduled mtg from Monday, May 25th to Tuesday May 26th at 11 am EDT due to Monday being US holiday
  * [test spreadsheet](https://docs.google.com/spreadsheets/d/1mEU8B2_PmMwwgp-_xnyp7QYMBwcLoA9NNlHwDyMvO0Y/edit?usp=sharing) updated with all sig-node tests. Signup if interested! Few slots open
  * Priority - merge blocking, release blocking, release informing
  * conformance-node-rhel test - no result
    * Import tests results from AWS E2E testing to testgrid
    * data is in bucket, but results are not in testgrid
      * gs://kubernetes-github-redhat/logs/ci-kubernetes-conformance-node-e2e-containerized-rhel/10653/
    * KETTLE issue? Any pointers?
      * [https://github.com/kubernetes/test-infra/tree/master/kettle](https://github.com/kubernetes/test-infra/tree/master/kettle)
* [vpickard, bart0sh, mhb]
  * [bart0sh] PR to update cos-stable images [https://github.com/kubernetes/test-infra/pull/17617](https://github.com/kubernetes/test-infra/pull/17617) updated cos image
  * [https://github.com/kubernetes/kubernetes/issues/91292](https://github.com/kubernetes/kubernetes/issues/91292) sig-node release-blocking failure on 5/20/2020. COS images had been updated, intermittent failures. Learned that COS image testing was broken (not being tested, silently failing) for the last \~4 weeks. Debugged, replaced broken COS image with newer one. Thanks @bart0sh, @mhb for great debugging!
    * Victor to send COS image PRs to Ning Liao and Roy Yang
      * COS image policy may need updating per Dawn ( I don’t understand the image policy, maybe Ning and Roy can share?)
    * Create email list and update jobs with email alias to alert folks that are monitoring jobs to avoid fire-drills when release-blocking/merge-blocking tests fail

## May 19, 2020

* [Javier Diaz-Montes] Discuss new feature to set FQDN as hostname of pods, issue [#91036](https://github.com/kubernetes/kubernetes/issues/91036), initial draft for PR in [#91035](https://github.com/kubernetes/kubernetes/pull/91035). KEP PR in [kubernetes/enhancements/1792](https://github.com/kubernetes/enhancements/pull/1792)
* [vinaykul] In-Place Pod Vertical Scaling v1.19 update &amp; [CRI-API design](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191025-kubelet-container-resources-cri-api-changes.md#proposal) question
  * Updates:
    * API code changes initial review done by Tim Hockin, David Ashpole.
    * First-cut implementation \~3 weeks to PR-ready.
      * David Ashpole is primary reviewer for Kubelet &amp; CRI changes.
  * Concern:
    * CRI clients may return partial or no CPU/memory limit info in ContainerStatus response. What’s the best way to handle this?
      * Option 1: Assume zero means no information returned?
      * Option 2: Add a flag that CRI client can set?
      * Dawn: Prefer to Option 1 since ContainerStatus is setting with the value reading from the host directly, and 0 is invalid the value for the kernel.
* [vpickard] sig-node testing enhancements update ([doc](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit?usp=sharing))
  * rescheduled mtg from Monday, May 25th to Tuesday May 26th at 11 am EDT due to Monday being US holiday
  * [test spreadsheet](https://docs.google.com/spreadsheets/d/1mEU8B2_PmMwwgp-_xnyp7QYMBwcLoA9NNlHwDyMvO0Y/edit?usp=sharing) updated with all sig-node tests. Signup if interested!
  * Priority - merge blocking, release blocking, release informing
  * [bart0sh] PR to update cos-stable images - needs /approve [https://github.com/kubernetes/test-infra/pull/17617](https://github.com/kubernetes/test-infra/pull/17617)
    * follow-up PR will move to latest LTS cos image
  * conformance-node-rhel test - no result
    * Import tests results from AWS E2E testing to testgrid
    * data is in bucket, but results are not in testgrid
      * gs://kubernetes-github-redhat/logs/ci-kubernetes-conformance-node-e2e-containerized-rhel/10653/
    * KETTLE issue? Any pointers?
      * [https://github.com/kubernetes/test-infra/tree/master/kettle](https://github.com/kubernetes/test-infra/tree/master/kettle)
* [cezaryzukowski, cynepco3hahue, bg.chun, krzwiatrzyk] How can we move forward with the Memory Manager KEP? Could we get any feedback (approval, change-request, partial-merge of KEP (prologue sections of KEP, [Summary](https://github.com/kubernetes/enhancements/blob/7bf4d0253f915677a50702ee12cef46e0df7758e/keps/sig-node/20200203-memory-manager.md#summary) till [Story 2 : Databases](https://github.com/kubernetes/enhancements/blob/7bf4d0253f915677a50702ee12cef46e0df7758e/keps/sig-node/20200203-memory-manager.md#story-2--databases)), etc.)? Today is the day of Enhancement Freeze.
* [vpickard] sig-node resource management forum ([doc](https://docs.google.com/document/d/15G9f0yoP8Gy5Agx-uCqo0CUl67EjZ-FgAk-KbI-lakw/edit?usp=sharing))
  * Reviewed Memory Manager presentation
  * Some concern about new kubelet flags and impact to user experience (more flags!)
  * For multi-numa hint generation, some discussion about preferred flag, great explanation from @klueska [here](https://www.google.com/url?q=https://kubernetes.slack.com/archives/C012XSGFZQE/p1589547856109400?thread_ts%3D1589542354.108700%26cid%3DC012XSGFZQE&sa=D&ust=1589902108688000&usg=AFQjCNGxuwACiCF5fLVWIn5mK4eAlbxzQA)
  * [bg.chun/review request] Update Topology Manager to support pod-level resource alignment
    * [https://github.com/kubernetes/enhancements/pull/1752](https://github.com/kubernetes/enhancements/pull/1752)
* [@rata, @alban] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))

## May 12, 2020

* [kad] Our experience of advanced resource management (CPU, Memory, etc.) 10 mins [demo](https://asciinema.org/a/327044) + \~20 mins for other slides.
* [cezaryzukowski, cynepco3hahue, bg.chun, krzwiatrzyk] Memory Manager KEP:
  * [\[ presentation \]](https://docs.google.com/presentation/d/1WLHbEdm3vO94eu5QyQGkutXfbfWcci4-kEDmp8ty8Ho/edit?usp=sharing)  \~25 mins + Q&amp;A
  * KEP: [https://github.com/kubernetes/enhancements/pull/1203](https://github.com/kubernetes/enhancements/pull/1203) ( [\[file preview\]](https://github.com/kubernetes/enhancements/blob/83fed369a22aff23edd2a7962fbab224eeefc990/keps/sig-node/20200203-memory-manager.md) )
* [joe conway (Crunchy), mrunalp (Red Hat)] - Challenges with running Postgres on Kubernetes [https://github.com/kubernetes/kubernetes/issues/90973](https://github.com/kubernetes/kubernetes/issues/90973)
* [vinaykul] In-Place Pod Vertical Scaling v1.19 update &amp; [CRI-API design](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20191025-kubelet-container-resources-cri-api-changes.md#proposal) question
  * API changes reviewed by Tim Hockin - one more change coming.
  * First-cut implementation \~3 weeks to PR-ready, identify primary reviewer.
  * Concern: CRI client may return partial on no CPU/memory limit info in ContainerStatus response. What’s the best way to handle this?
    * Option 1: Assume zero means no information returned?
    * Option 2: Add a flag that CRI client can set?
* [krzwiatrzyk, bgchun] Request for review - Topology Manager *pod-level-single-numa-node* policy [KEP’s update PR](https://github.com/kubernetes/enhancements/pull/1752)
* [vpickard,jaypipes] sig-node test kickoff update
  * great attendance
  * thanks to @dims for sharing how to navigate around and debug! Will incorporate into documentation
  * sig-node test [document](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit#)
  * sig-node test [spreadsheet](https://docs.google.com/spreadsheets/d/1mEU8B2_PmMwwgp-_xnyp7QYMBwcLoA9NNlHwDyMvO0Y/edit?usp=sharing)
  * mtg [recording](https://bluejeans.com/s/8OsOP)
  * Next: complete spreadsheet with remainder of tests, some cleanup of columns. Volunteers sign up for tests to investigate.
  * Weekly meeting Monday at 1 pm EDT until we get this under control
* [vpickard] sig-node working group for Topology Aware Scheduling
  * #topology-aware-scheduling channel for discussion
  * meeting logistics - propose Thursday 9 am EDT, will survey if needed
  * meetings will be recorded for those not able to attend
  * Send invite to sig-node group with mtg link
* seccomp-operator: proposal to move to k-sig
    [https://github.com/saschagrunert/seccomp-operator/blob/master/RFC.md](https://github.com/saschagrunert/seccomp-operator/blob/master/RFC.md)

## May 5, 2020

* Making Seccomp GA: [PR: 1148](https://github.com/kubernetes/enhancements/pull/1148) current state / changes proposed. (@pjbgf)
* seccomp-operator: [out of tree project to extend seccomp](https://github.com/saschagrunert/seccomp-operator/blob/master/RFC.md). (@sascha)
* Needs reviewers for CRI API updates: [PR: 90061](https://github.com/kubernetes/kubernetes/pull/90061) (@marosset)
* [krzwiatrzyk, aperavelov] NUMA-aware scheduling, new approach
  * presentation: [https://docs.google.com/presentation/d/1y2ObdZUtMFp2Z0EeyUu1cm469EQNMBjJbhjKiHwSeQo/edit?usp=sharing](https://docs.google.com/presentation/d/1y2ObdZUtMFp2Z0EeyUu1cm469EQNMBjJbhjKiHwSeQo/edit?usp=sharing)
  * PoC: [https://github.com/kubernetes/kubernetes/pull/90708](https://github.com/kubernetes/kubernetes/pull/90708)
  * Proposal to create a smaller interest group that can discuss this in a higher-bandwidth fashion
    * Interested parties:
      * jaypipe
      * vpickard
      * krzwiatrzyk
      * aperavelov
      * derekwaynecarr
      * kad
      * dchen1107 (dawnchen)
      * nolancon
      * Ed
      * swsehgal
      * bg.chun
      * zhiyu
      * fromani
      * cynepco3hahue(alukiano)
* [vpickard, jaypipes] Sig-node testing infrastructure update
  * Met with Aaron (spiffxp), notes in 1st link below
  * Plan to create sig-node-test email group to notify volunteers of test failures and update all job
  * sig-node-kubelet jobs are running on a new project, thanks @spiffxp. These jobs run now, but lots of red (failures). Jobs with no color are passing consistently. Volunteers can sign up to look at jobs in this test group. For reference, 2 jobs have been filled in with info. Open to feedback/input as to format/content.
  * Volunteers, please update your contact info at the top [https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit?usp=sharing](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit?usp=sharing)
  * Spreadsheet for volunteers to add info on existing jobs [https://docs.google.com/spreadsheets/d/1mEU8B2_PmMwwgp-_xnyp7QYMBwcLoA9NNlHwDyMvO0Y/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1mEU8B2_PmMwwgp-_xnyp7QYMBwcLoA9NNlHwDyMvO0Y/edit?usp=sharing)
* [SaranBalaji90] looking for initial feedback on [https://github.com/kubernetes/enhancements/pull/1461](https://github.com/kubernetes/enhancements/pull/1461)

## April 28, 2020

* [SaranBalaji90] Get feedback on adding node-local plugin support for pod admission handler [https://github.com/kubernetes/enhancements/pull/1461](https://github.com/kubernetes/enhancements/pull/1461)  [xing-yang/yuxiangqian] ContainerNotifier
  * [https://groups.google.com/forum/#!topic/kubernetes-sig-node/vDqbgwQeh4g](https://groups.google.com/forum/#!topic/kubernetes-sig-node/vDqbgwQeh4g)
  * [https://docs.google.com/document/d/1Q4UMrx9r58LUNOHdn_4nDhqHN21iOOYQeViMKSPa_Ks/edit?usp=sharing](https://docs.google.com/document/d/1Q4UMrx9r58LUNOHdn_4nDhqHN21iOOYQeViMKSPa_Ks/edit?usp=sharing)
  * KEP uses CRD and external controller approach:
    [https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/20190120-execution-hook-design.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/20190120-execution-hook-design.md#alternative-option-1a)
    * Implementation: [https://github.com/kubernetes-sigs/execution-hook/pull/5](https://github.com/kubernetes-sigs/execution-hook/pull/5), [https://github.com/kubernetes-sigs/execution-hook/pull/4](https://github.com/kubernetes-sigs/execution-hook/pull/4)
    * **AI**: Compare pros and cons of ContainerNotifier vs CRD approach (application operator)

* [krzwiatrzyk, c.zukowski, bg.chun, +@] Topology Manager Enhancement Proposal, slide([Discussion about Topology Manager](https://docs.google.com/presentation/d/1SR6XSIFsHkiWTws66LABpTiZaRwYk8IoRts4HQWE8Bc/edit#slide=id.g74d9a6aad0_49_0))
  * Support binding multiple topology policies in a node
  * scheduler enhancement
* [vpickard, jaypipes] Sig-node testing infrastructure update
  * Met with Aaron (spiffxp), notes in 1st link below
  * Create sig-node-test email group to notify volunteers of test failures and update all job
  * Volunteers, please update your contact info at the top [https://docs.google.com/document/d/1edtGmngdjQStAaSqGdndaJqi47SQN535cH9k18YRG-U/edit?usp=sharing](https://docs.google.com/document/d/1edtGmngdjQStAaSqGdndaJqi47SQN535cH9k18YRG-U/edit?usp=sharing)
  * Spreadsheet for volunteers to add info on existing jobs [https://docs.google.com/spreadsheets/d/1nc2TffhSNaJi1NLf3QdJH5eRUs8ar-ju2EjD9GwWcLQ/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1nc2TffhSNaJi1NLf3QdJH5eRUs8ar-ju2EjD9GwWcLQ/edit?usp=sharing)

## Apr 21, 2020

* [derekwaynecarr] SIG Health Check
  * retro: release blocking test was RED for 10d [4/6-4/16]
    * [https://testgrid.k8s.io/sig-release-master-blocking#node-kubelet-master](https://testgrid.k8s.io/sig-release-master-blocking#node-kubelet-master)
    * network access blocked
      * [https://github.com/kubernetes/kubernetes/issues/89847#issuecomment-610109930](https://github.com/kubernetes/kubernetes/issues/89847#issuecomment-610109930)
    * Actions Taken and owners required for next step
      * [https://github.com/kubernetes/kubernetes/issues/89892#issuecomment-614998826](https://github.com/kubernetes/kubernetes/issues/89892#issuecomment-614998826)
  * Q: Are we able to sustain the number of test suites we run? Is pruning required?
    * [https://testgrid.k8s.io/sig-node-kubelet](https://testgrid.k8s.io/sig-node-kubelet)
  * Q: Carrot and sticks to improve sustainability?
  * Q: Impact of covid situation on members?  Holiday week?
  * Q: Volunteers for mentorship?
  * Notes:
    * [dawn] Communication issues around internal changes, and growing new engineers to backfill.
    * [dawn] Stale old image caused the network blocking.  Viewed this as an opportunity to grow new members to backfill that role and had a hand-off problem.
    * [mikebrown] GCP issue hit containerd community pre-submits.
    * [dawn] sig built first level conformance test, node e2e expanded to cluster level, but still want to hold more node e2e as release blocker as its harder to grow to cluster level and keep deterministic, particularly on tests for different node profiles.
    * [victor] testgrid alert email when there is a test failure.  not working as he was not getting notifications on his related tests.
    * [derek] ask for volunteers to audit state of test-infrastructure and provide some recommendations back to sig, jay/victor to coordinate a small sub-group in sig over mailing list to determine next steps.
      * [jaypipes, saran balaji] volunteer from aw
      * [victor] volunteer from red hat
      * [morgan] volunteer from ibm
      * [ning liao, david porter] volunteer from google
      * [bart0sh] volunteer from intel
      * [daniel mangum] ci-signal lead for 1.19, aaron is working on transition to community owned infrastructure.
* [SaranBalaji90] Add node-local plugin support for pod admission handler [https://github.com/kubernetes/kubernetes/pull/87273](https://github.com/kubernetes/kubernetes/pull/87273) and disabling logging handler in kubelet [https://github.com/kubernetes/enhancements/pull/1461](https://github.com/kubernetes/enhancements/pull/1461)
* [howardjohn] What is needed to get moving on [Sidecar Containers](https://github.com/kubernetes/enhancements/issues/753)?

## Apr 14, 2020

* [tedyu] Remove unhealthy symlink only for dead containers [https://github.com/kubernetes/kubernetes/pull/89160](https://github.com/kubernetes/kubernetes/pull/89160).Updated description after two iterations of review comment
* [dashpole] Allocatable eviction for PIDs.  [https://github.com/kubernetes/kubernetes/issues/89807](https://github.com/kubernetes/kubernetes/issues/89807)Should we add “all of these” metrics to the summary API, or do we need an internal version of the summary API so we don’t keep expanding the summary API?
  * Follow-up with derek.
* [krzwiatrzyk, c.zukowski, bg.chun] Topology Manager Enhancement-Pod Spec extension to describe topology policy-Improve TM to allow binding Pods with multiple topology policy to the same node.[https://github.com/kubernetes/enhancements/pull/1674](https://github.com/kubernetes/enhancements/pull/1674)- presentation slides: [https://docs.google.com/presentation/d/1LTfgP_n51zxr57Gw7UKqTL39JdqsoTnTfnK5vVHq8Hc/edit?usp=sharing](https://docs.google.com/presentation/d/1LTfgP_n51zxr57Gw7UKqTL39JdqsoTnTfnK5vVHq8Hc/edit?usp=sharing)

## Apr 7, 2020

* [tedyu] Protect log rotation against concurrent symlink removal
    [https://github.com/kubernetes/kubernetes/pull/89160](https://github.com/kubernetes/kubernetes/pull/89160)
    Prototype for checking container status before symlink removal is attached to the PR.
* [k.wiatrzyk (krzwiatrzyk), c.zukowski, bg.chun] Enhancements proposals for 5G packet processing in Kubernetes.
  * Towards high-performance 5G packet processing, [slide](https://docs.google.com/presentation/d/1DueZ65OEPHvEMBf_nck7YOxWIXhOpYTIiPNix7rGIik/edit#slide=id.p)
  * Overall Proposals in detail, [docs](https://docs.google.com/document/d/1ASbIJNfeWmYrBIEMQl1MxHBPgvoztb_Fx5GvFhoPXF0/edit)

## March 31, 2020

* [tedyu] Protect log rotation against concurrent symlink removal
    [https://github.com/kubernetes/kubernetes/pull/89160](https://github.com/kubernetes/kubernetes/pull/89160)
* [harche] Add CRIO tests for Kubernes Prow (CI)
  * [https://github.com/kubernetes/test-infra/pull/17014](https://github.com/kubernetes/test-infra/pull/17014)
* [c.zukowski, cynepco3hahue, bg.chun] a warm invitation to leave your comments for a new component proposed in: [https://docs.google.com/document/d/1XvOUwI1DiZf0vTlIa2vPOT_QJYh-g3vpOa1D1pajqEs/edit?usp=sharing](https://docs.google.com/document/d/1XvOUwI1DiZf0vTlIa2vPOT_QJYh-g3vpOa1D1pajqEs/edit?usp=sharing)
    Memory Manager opens a series of further enablers of NUMA-aware resource management in Kubernetes. Do not hesitate to ask &quot;Why?&quot; questions, and point out whenever the document should expand on or stress something else (we will elaborate).
* [mrunalp] KEP for custom cri errors - [https://github.com/kubernetes/enhancements/pull/1654](https://github.com/kubernetes/enhancements/pull/1654)
* [jberkus] Discuss [year-of-support KEP](https://github.com/kubernetes/enhancements/pull/1497), particularly what the effect would be on required kubelet version skew
  * will just keep skew tests running for one more version
  * don't see this as serious problem
* [mattjmcnaughton/dims] KEP- Build Kubelet without Docker
  * [https://github.com/kubernetes/enhancements/pull/1546](https://github.com/kubernetes/enhancements/pull/1546)
    * KEP adds a “dockerless” tag to avoid docker/docker imports, builds kubelet without dockershim
  * PR [https://github.com/kubernetes/enhancements/pull/1546](https://github.com/kubernetes/enhancements/pull/1546)
* Need a way to mock network related testing for example to test issue:
  * [https://github.com/kubernetes/kubernetes/issues/88543](https://github.com/kubernetes/kubernetes/issues/88543)

## March 24, 2020

* [mrunalp] CRI custom errors update[https://docs.google.com/presentation/d/1DdmlEmK4TkfJWW19wXPReAw8MA0oRdzPwDDhjtDKJo8/edit#slide=id.g8190ad4d27_0_38](https://docs.google.com/presentation/d/1DdmlEmK4TkfJWW19wXPReAw8MA0oRdzPwDDhjtDKJo8/edit#slide=id.g8190ad4d27_0_38)
* [mikebrow] Ensure Secret Pulled Images (without pull always)
  * [https://github.com/kubernetes/enhancements/pull/1608](https://github.com/kubernetes/enhancements/pull/1608)
* [decarr] Identifying areas to improve kubelet reliability across contributor
  * A number of subtle prs have been opened recently around edge cases on pod status reporting, wanted us to have a plan of record as a group on what we tackle so we can have full context when reviewing.  They are non-trivial to review given their subtle updates, so just having a dialog on what the rationale will help all.
  * Examples:
    * [https://github.com/kubernetes/kubernetes/pull/89155/files](https://github.com/kubernetes/kubernetes/pull/89155/files) which targets issue #80968
    * Other
* [decarr] flaky node conformance [help desired]
  * [https://github.com/kubernetes/kubernetes/issues/85762#issuecomment-601273307](https://github.com/kubernetes/kubernetes/issues/85762#issuecomment-601273307)
  * [https://github.com/kubernetes/kubernetes/pull/89379/files](https://github.com/kubernetes/kubernetes/pull/89379/files)
* [keerthan] Issue: [https://github.com/kubernetes/kubernetes/issues/88543](https://github.com/kubernetes/kubernetes/issues/88543)
  * [dashpole] test do create delete pod through kubelet and interact with cri to ensure its gone (similar to checking inside cgroups for kubelet settings)
* [dashpole] Out-of-Pid eviction update [https://github.com/kubernetes/kubernetes/pull/89359](https://github.com/kubernetes/kubernetes/pull/89359).  Pod-level PID stats in the summary API?  KEP update: [https://github.com/kubernetes/enhancements/pull/1631](https://github.com/kubernetes/enhancements/pull/1631)
* [tallclair] RuntimeClass to GA
* [smarterclayton] Doc pending on Kubelet status suggestions (ran out of time last week), will email sig-node with a shared google doc soon
* [smarterclayton] Discuss exposing calculated resource limits as metrics briefly - too hard today to emulate the complex calculations scheduler and node do for resource limit in external monitoring tool

## March 17, 2020

* [vpickard] Please upload latest video recording
* [liorokman] Enable defining Pod-level resource limit
  * [https://github.com/kubernetes/enhancements/pull/1592](https://github.com/kubernetes/enhancements/pull/1592)
  * Way forward is to implement this using a daemonset that manipulates the cgroups externally to Kubernetes, and report back once there is some real-world experience around this feature.

## March 10, 2020

* [roycaihw] When a kubelet goroutine (e.g. a pod worker) panicked, kubelet should crash and restart, instead of keeping running with a non-functioning pod worker
  * [https://github.com/kubernetes/kubernetes/pull/88915](https://github.com/kubernetes/kubernetes/pull/88915)
  * Kubelet has a [flag](https://github.com/kubernetes/kubernetes/blob/900143c6d490616604d2f91e979d591eb78c20ae/cmd/kubelet/app/options/options.go#L360) controlling the crash behavior, which is [defaulted](https://github.com/kubernetes/kubernetes/blob/900143c6d490616604d2f91e979d591eb78c20ae/cmd/kubelet/app/options/options.go#L195) to not crash
    * Background:
      * The flag was introduced in [2015](https://github.com/kubernetes/kubernetes/pull/3577)
      * We thought we made kubelet crash in [2016](https://github.com/kubernetes/kubernetes/pull/28800#issuecomment-231911897), but we didn’t
  * Alternatively we could create a new goroutine when the old one died
  * AI [roycaihw] Does systemd have an easy way to check how many times kubelet restarted?
  * systemd service restart counter [follow-up to see]
    * [https://github.com/systemd/systemd/pull/6495](https://github.com/systemd/systemd/pull/6495)
    * <https://github.com/prometheus/node_exporter/blob/master/collector/systemd_linux.go#L103>
* [liorokman] Enable defining Pod-level resource limit
  * [https://github.com/kubernetes/enhancements/pull/1592](https://github.com/kubernetes/enhancements/pull/1592)
  * [https://github.com/kubernetes/kubernetes/pull/88899](https://github.com/kubernetes/kubernetes/pull/88899)
    * Figure out how this impacts the Init container
    * Verify if resources are released for reuse between sibling cgroups in the pod
    * What is the expected effect on hugetlb?
    * What is the interaction with the ResourceOverhead feature?
    * How would this work in NUMA environments?
    * Does it make more sense to not tie this to the QoS functionality, or make this available for BestEffort pods on the pod level?
    * Does this also apply to ephemeral storage (tmpfs) or possibly pids (if we ever made that container level rather than a global default)
* [vpickard] Topology Manager documentation PRs need review/merge to finish Beta
  * [https://github.com/kubernetes/website/pull/19050](https://github.com/kubernetes/website/pull/19050)
  * [https://github.com/kubernetes/community/pull/4503](https://github.com/kubernetes/community/pull/4503)
  * Windows platform affected by recent merge. Is there a CI job that should be run on PRs to help prevent this in the future? [https://github.com/kubernetes/kubernetes/pull/88917](https://github.com/kubernetes/kubernetes/pull/88917)
* [mattjmcnaughton/dims] KEP- Build Kubelet without Docker
  * [https://github.com/kubernetes/enhancements/pull/1546](https://github.com/kubernetes/enhancements/pull/1546)
    * KEP adds a “dockerless” tag to avoid docker/docker imports, builds kubelet without dockershim
* [smarterclayton] Looking at pod end to end latency in the kubelet with an eye on the status loop
  * The status loop is very simple and reliable (good!) but fairly slow
  * No KEP yet, gathered some feedback from ashpole and investigating
    * In an e2e run, we take about \~800s overall from time we detect a status change to the time we successfully write it to apiserver across all tests in sum
    * With some simple changes, i was able to get that down to less than 200s (which means certain types of pod operations complete a lot faster)
    * Found a few correctness bugs already in our core sync loop
      * Internal kubelet state not in sync with apiserver
      * We check the wrong cache for certain pod
      * We are depending on a live lookup to create the patch, but there is no guarantee that is up to date either
      * Kubelet is sending new data when feature flags are off (bad!)
      * Some internal safety checks in kubelet are just log messages, but potentially should be crashes (we ignored them as logs)
  * Going to move to a KEP for some improvements soon
    * “Improve pod status reporting latency and prioritize important transitions”

## March 3, 2020

* [vpickard,nolancon] Topology Manager -&gt; Beta for 1.18. Need review/approval/merge following PRs by Thursday, March 5
  * [https://](https://github.com/kubernetes/kubernetes/pull/87759)[github](https://github.com/kubernetes/kubernetes/pull/87759)[.com/kubernetes/kubernetes/pull/87759](https://github.com/kubernetes/kubernetes/pull/87759)
  * [https://github.com/kubernetes/kubernetes/pull/87650](https://github.com/kubernetes/kubernetes/pull/87650)
  * [https://github.com/kubernetes/test-infra/pull/16547](https://github.com/kubernetes/test-infra/pull/16547)
  * [https://github.com/kubernetes/kubernetes/pull/88566](https://github.com/kubernetes/kubernetes/pull/88566)
  * [https://github.com/kubernetes/kubernetes/pull/88721](https://github.com/kubernetes/kubernetes/pull/88721)
* [vpickard,nolancon] Topology Manager documentation PRs, merge by March 16
  * [https://github.com/kubernetes/website/pull/19050](https://github.com/kubernetes/website/pull/19050)
  * [https://github.com/kubernetes/community/pull/4503](https://github.com/kubernetes/community/pull/4503)
  * [https://github.com/kubernetes/enhancements/pull/1548](https://github.com/kubernetes/enhancements/pull/1548)
* [mrunalp] Custom errors in CRI
* [bart0sh] HugePage
  * [https://github.com/kubernetes/website/pull/19008](https://github.com/kubernetes/website/pull/19008)
  * [kubelet: sync node allocatable cgroups upon status update](https://github.com/kubernetes/kubernetes/pull/81774) bugfix
* [joelsmith] Missing Memory statistics on exited cri-o init containers breaks HPA
  * For runtimes that don't use cAdvisor stats provider in kubelet (i.e. CRI stats provider), we populate CPU and memory zeroes to fix the same issue: [https://github.com/kubernetes/kubernetes/pull/74336](https://github.com/kubernetes/kubernetes/pull/74336)
  * Proposal: do basically the same thing for the cAdvisor stats provider in kubelet, a la [https://github.com/kubernetes/kubernetes/pull/88734](https://github.com/kubernetes/kubernetes/pull/88734)

## February 25,2020

* [giuseppe,mrunalp] Cgroups v2
* [mrunalp] Termination grace period at shutdown [https://github.com/kubernetes/kubernetes/pull/88495](https://github.com/kubernetes/kubernetes/pull/88495)
* [spikecurtis] Issue: [https://github.com/kubernetes/kubernetes/issues/85966](https://github.com/kubernetes/kubernetes/issues/85966)
* [bart0sh] 2 hugepages PRs need review/approval:
  * [Implement support for multiple sizes huge pages](https://github.com/kubernetes/kubernetes/pull/84051) was already reviewed and passed API review. Last signode review is required. This PR is covered by [granted sig-release exception](https://github.com/kubernetes/enhancements/pull/1540#issuecomment-583072382) and is going to be included into 1.18 if merged.
  * [kubelet: sync node allocatable cgroups upon status update](https://github.com/kubernetes/kubernetes/pull/81774) bugfix got lgtm in Oct 2019. It would be great to have it in 1.18.
* [cynepco3hahue] The KEP [Propose Memory Manager for NUMA awareness](https://github.com/kubernetes/enhancements/pull/1203) needs review, it can be great to get feedback from someone under the the sig-node group.
* [vpickard] ~~Topology Manager CI job PR needs /approve from sig-node owner [https://github.com/kubernetes/test-infra/pull/16062](https://github.com/kubernetes/test-infra/pull/16062)~~
* [fromani] Topology Manager Upgrade test PR [https://github.com/kubernetes/kubernetes/pull/88107](https://github.com/kubernetes/kubernetes/pull/88107)
* [dashpole] Distributed Tracing in Kubernete
  * [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/0034-distributed-tracing-kep.md), Kubecon [talk](https://www.youtube.com/watch?v=lEACvRW6T_U&feature=youtu.be&t=909), [Slides](https://docs.google.com/presentation/d/15XiScaHskf1CqT-EGC8zemSdZl-vLtUBJhzW38InQh4/edit?usp=sharing)
* [vinaykul] In-Place Vertical Scaling update
  * API code changes addressing Tim &amp; Jordan’s feedback ready for review
  * [https://github.com/vinaykul/kubernetes/pull/1](https://github.com/vinaykul/kubernetes/pull/1)
* [vpickard] Topology Manager feature enablement causes scaling failure [https://github.com/kubernetes/kubernetes/pull/87650](https://github.com/kubernetes/kubernetes/pull/87650)

## February 18,2020

* [vpickard,fromani] E2E Topology Manager PRs needing review/approval
  * Upgrade tests needs review/approval [https://github.com/kubernetes/kubernetes/pull/88107](https://github.com/kubernetes/kubernetes/pull/88107)
  * Multiple container NUMA alignment tests needs review/approval [https://github.com/kubernetes/kubernetes/pull/88234](https://github.com/kubernetes/kubernetes/pull/88234)
  * CPU Manager E2E tests Numa awareness needs /approve from sig-node owner [https://github.com/kubernetes/kubernetes/pull/87921](https://github.com/kubernetes/kubernetes/pull/87921)
  * Topology Manager CI job PR needs /approve from sig-node owner
  * [https://github.com/kubernetes/test-infra/pull/16062](https://github.com/kubernetes/test-infra/pull/16062)
  * e2e and e2e_node tests code refactor needs review/approve [https://github.com/kubernetes/kubernetes/pull/88110](https://github.com/kubernetes/kubernetes/pull/88110)
    * @patricklang will check re: Windows. Plan is to try running SIG-Windows tests after rebasing #[86101](https://github.com/kubernetes/kubernetes/pull/86101) which was also broken by changes in test/e2e/framework
  * Sync owners for test-infra with present sig-node owner
    * [https://github.com/kubernetes/test-infra/pull/16037](https://github.com/kubernetes/test-infra/pull/16037) - Done!
* [vpickard] Enable Topology Manager Feature PR [https://github.com/kubernetes/kubernetes/pull/87650](https://github.com/kubernetes/kubernetes/pull/87650)
  * Consistently fails pull-kubernetes-kubemark-e2e-gce-big
  * Kubelet is killing pods/containers, suspect additional memory consumed by Topology Manager is causing kubelet eviction and/or oom killer
  * Reached out to #sig-scalability and #sig-testing for support, and working to reproduce on local setup. Is it possible to interactively debug this failing job?
  * Debug PR ~~needs ok-to-test~~ [https://github.com/kubernetes/kubernetes/pull/88275](https://github.com/kubernetes/kubernetes/pull/88275)
* [nolancon] Topology Manager PRs - Related issue [https://github.com/kubernetes/kubernetes/issues/83476](https://github.com/kubernetes/kubernetes/issues/83476)
  * Guarantee aligned resources across multiple containers - needs review/approval [https://github.com/kubernetes/kubernetes/pull/87759](https://github.com/kubernetes/kubernetes/pull/87759)
  * [https://github.com/kubernetes/kubernetes/pull/87983](https://github.com/kubernetes/kubernetes/pull/87983) will need rebase after 87759 is merged

## February 4, 2020

* [spikecurtis] Issue: [https://github.com/kubernetes/kubernetes/issues/85966](https://github.com/kubernetes/kubernetes/issues/85966)
* [bart0sh] Hugepages PRs need review and/or approval:
  * [Implement support for multiple sizes huge pages](https://github.com/kubernetes/kubernetes/pull/84051)
  * [kubelet: sync node allocatable cgroups upon status update](https://github.com/kubernetes/kubernetes/pull/81774)
* [vinaykul] In-Place Vertical Scaling weekly update
  * API coding work in progress, a dev resource short
* [giuseppe,mrunalp] cgroupsv2 update [https://asciinema.org/a/UsX5p60wNaMZQJrDsDWrS8rcZ](https://asciinema.org/a/UsX5p60wNaMZQJrDsDWrS8rcZ)
  * kep [https://github.com/kubernetes/enhancements/pull/1370](https://github.com/kubernetes/enhancements/pull/1370)

## January 28, 2020

* [Rootless mode KEP](https://github.com/kubernetes/enhancements/pull/1371) needs reviews! (Added by @dims on behalf of @AkihiroSuda)
  * Related [PR #78635](https://github.com/kubernetes/kubernetes/pull/78635) is ready to merge (needs hold to be removed as well)
  * Related KEP [cgroups v2](https://github.com/kubernetes/enhancements/pull/1370) needs approval
  * NOTE: the PR above is tiny enough (and already [used in say k3s](https://github.com/rancher/kubernetes/commit/b4da65ff558aa944c4d339296151c6b72f549a1d)), so it may be worth merging it.
* [vinaykul] In-Place Vertical Scaling follow-up
  * KEP [updated](https://github.com/kubernetes/enhancements/pull/1342) based on feedback from @liggitt
  * Is it close enough to merge before Jan 28 3.00 pm deadline for 1.18?
* [patricklang, kkmsft] RuntimeClass / pull annotations [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/windows-runtimeclass-support.md) (#[1448](https://github.com/kubernetes/enhancements/pull/1448)) - very close to merge, need OK from mikebrown &amp; random-liu on final version
  * propose keeping it narrow for first phase - don’t require CRI to store annotations. Just pass them down to Pull &amp; CheckIfExists.
  * second phase would allow containerd to store annotations, and pass them back. This would be needed if you want to support multiple ways of pulling and enumerating what could be the same image, eg for experimenting with pluggable snapshotters or other CRI-specific features later. This would allow the annotations to be returned from crtctl for example &lt;comments holder/&gt;
* ~~[vpickard, klueska] Need approvals on TopologyManager KEP update PRs~~
  * [~~https://github.com/kubernetes/enhancements/pull/1477~~](https://github.com/kubernetes/enhancements/pull/1477)
  * [~~https://github.com/kubernetes/enhancements/pull/1519~~](https://github.com/kubernetes/enhancements/pull/1519)
* [vpickard] Need approvals on TopologyManager E2E testing PR
  * [https://github.com/kubernetes/test-infra/pull/15609](https://github.com/kubernetes/test-infra/pull/15609)
  * [https://github.com/kubernetes/test-infra/pull/16037](https://github.com/kubernetes/test-infra/pull/16037)
* [vpickard, vladikr] Introduce PR for new kubelet option for system reserved cpus. Would appreciate input/comments on PR. [https://github.com/kubernetes/kubernetes/pull/87532](https://github.com/kubernetes/kubernetes/pull/87532)
  * Suggestion to convert this to KEP. Deadline is today for 1.18
* [tallclair] Deprecate StreamingProxyRedirects [kubernetes/enhancements#1395](https://github.com/kubernetes/enhancements/pull/1395)

## January 21, 2020

* [@SaranBalaji90] Get opinions from others on [#87252](https://github.com/kubernetes/kubernetes/issues/87252). Adding support for disabling /logs endpoint in kubelet
  * agreement that PR should only add ComponentConfig fields (no CLI args for kubelet)
  * general approval of exposing finer-grained endpoint handlers as long as existing behaviour does not break (w.r.t enableDebugHandlers config)
* [vinaykul] In-Place Vertical Scaling follow-up
  * Moved KEP to sig-node directory
  * Awaiting review of API changes, test plan, GA criteria
  * Targeting to [merge KEPs as implementable](https://github.com/kubernetes/enhancements/pull/1342) before Jan 28 deadline for 1.18
* [@SaranBalaji90] Discuss about KEP - [#1461](https://github.com/kubernetes/enhancements/pull/1461/files) To implement “out-of-tree” plugin for pod admit handler. This will benefit from not needing to update k8s source code whenever we have new feature in docker/kernel versions and preventing pods that even tolerate any taints from running on custom worker nodes.

## January 14, 2020

* Update to RuntimeClass / pull annotations [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/windows-runtimeclass-support.md) (#[1448](https://github.com/kubernetes/enhancements/pull/1448)) - [@kkmsft, @patricklang]
  * Mike &amp; Lantao to review since they had feedback and suggestions on last proposal - if ok, will proceed. Dawn will /approve if they’re ok
* Container namespace targeting [#84731](https://pr.k8s.io/84731) for 1.18? [@verb]
  * Three concerns were raised:
    * How will zombies from the ephemeral container be cleaned up, if at all? For docker will they be configured automatically?
    * What's the behavior for Init Containers?
    * Is there any coordination to do with the [user namespace work](http://features.k8s.io/127)
  * @verb will send a PR updating the ShareProcessNamespace KEP to address these concerns.
* [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) - preparing to merge as ‘implementable’
  * Final changes from API review
    * New admission controller instead of subresource
    * Using list of named subresource
  * Test-plan, graduation criteria section added
  * Should we move [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-autoscaling/20181106-in-place-update-of-pod-resources.md) to sig-node directory?
    * Yes

## January 7, 2020

* [bg.chun] Asking a review for PR [#84154](https://github.com/kubernetes/kubernetes/pull/84154) of container isolation of hugepage
  * This change enables kubelet to set hugepage limit on container level cgroup sandbox when kubelet creates a container. and it is the last piece for container isolation of hugepages on kubernetes side except docker shim. (see next step for details)
  * Writing e2e-test on top of this change is almost done. My coworker will open PR to add test suit for container isolation soon. But it has a dependency on the PR#84154 so it will be tagged as WIP.
  * next step
    * In kubernetes side, next step of this work will be adding e2e-test and updating online [document](https://kubernetes.io/docs/tasks/manage-hugepages/scheduling-hugepages/).
    * In container runtime side, 1) changes for cri-o are merged. 2) changes for c8d are waiting merging now. 3) For docker, it requires more effort, 1. adding hugepages field on docker configuration fields. 2. update docker vendoring in kubernetes to have new field 3. update dockershim to set hugepage limit to docker configuration.
* [vinaykul, quinton-hoole] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) update
  * @thockin approved!!
  * Next steps:
    * Update [PR 1342](https://github.com/kubernetes/enhancements/pull/1342) with API review change
    * CRI changes review &amp; implementation plan
    * Formal SIG-Node approval, merge as implementable, and start writing code
