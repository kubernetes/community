# 2023 Annual Report: SIG Scalability

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

Notable work we produced in 2023 includes:
- Helping validate scalability/performance impact and reliability for many features across the year as well as preventing regressions
- Enhancements to our scalability test framework around better instrumentation and reduction in flakiness
- Additional test coverage spanning across the Kubernetes project (e.g list calls with AP&F, in-cluster network/DNS programming SLIs)
- Driving/influencing a bunch of KEPs and fixes, mainly within Kubernetes core (API server, API machinery and etcd) 

A key good news we now have recurring [5k-node scalability CI tests on AWS](https://github.com/kubernetes/test-infra/issues/29139) (using kops). These are currently release-informing and plan is to graduate them to release-blocking in 2024. This effort helps decentralize scalability testing and costs for the SIG that were previously heavily borne by GCP.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Overall in 2023 we have had healthy contributions from multiple companies (see these [devstats](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Scalability&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1672560000000&to=1704095999000). Compared to 2022, we have seen an  uptick in contributions around test framework improvements and test suite coverage for scalability/performance. Some of those were driven by other SIGs, but most continue to come from within the SIG. The increased contributions from AWS towards scale testing and debugging performance issues comes as good news. But overall, we continue to seek help from various SIGs to force-multiply scalability test coverage and regression hunting for features/components they own. We also encourage them to proactively identify and document SLIs/SLOs/limits for APIs and workflows they own. This allows each SIG to set a scalability bar for their systems (just like up-time/availability) and thereby make scalability a first class citizen of Kubernetes and related CNCF projects. As always, SIG scalability is eager to assist/guide with this process.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

We presented SIG Scalability Intro + Deep-Dive updates at both EU and NA KubeCons:
- [KubeCon Chicago](https://kccncna2023.sched.com/event/1R2pT/intro-deep-dive-kubernetes-sig-scalability-wojciech-tyczynski-google-marcel-zieba-isovalent)
  - [Slides](https://static.sched.com/hosted_files/kccncna2023/85/SIG%20Scalability%20-%20Kubecon%20NA%202023.pdf)
  - [Recording](https://youtu.be/nrAskCyG_Xk)
- [KubeCon Amsterdam](https://kccnceu2023.sched.com/event/1HyTL/intro-deep-dive-kubernetes-sig-scalability-wojciech-tyczynski-google)
  - [Recording](https://youtu.be/bxY5Q5Eoj0s)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

Notable KEPs, mostly co-owned with SIG API machinery and SIG etcd:
- [1040 - API Priority and Fairness](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1040-priority-and-fairness)
  - GA in 1.29
- [2340 - Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache)
  - Second alpha in 1.30 (blocked by resolution of [this etcd issue](https://github.com/kubernetes/kubernetes/issues/123072))
- [3157 - API Streaming Lists](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list)
  - Third alpha in 1.30 (blocked by resolution of [this etcd issue](https://github.com/kubernetes/kubernetes/issues/123072))
- [4222 - Binary Encoding for CRDs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4222-cbor-serializer)
  - Pre-alpha in 1.30

Some notable non-KEP improvements:
- [Graceful termination of watches during API server shutdown](https://github.com/kubernetes/kubernetes/pull/114925)
  - Landed in 1.27
- [Addressed monitoring gaps in API server extension mechanisms](https://github.com/kubernetes/kubernetes/issues/117167)
  - Landed in 1.28
- [Cache JSON-encoded watch events to reduce redundant work with multiple watches](https://github.com/kubernetes/kubernetes/pull/120300)
  - Landed in 1.29
- [Memory-efficient handling of watch requests preflight](https://github.com/kubernetes/kubernetes/pull/120902)
  - Landed in 1.30

## [Subprojects](https://git.k8s.io/community/sig-scalability#subprojects)

**New in 2023:**
- None

**Retired in 2023:**
- None

**Continuing:**
- [kubernetes-scalability-and-performance-tests-and-validation](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-and-performance-tests-and-validation)
- [kubernetes-scalability-bottlenecks-detection](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-bottlenecks-detection)
- [kubernetes-scalability-definition](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-definition)
- [kubernetes-scalability-governance](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-governance)
- [kubernetes-scalability-test-frameworks](https://git.k8s.io/community/sig-scalability#kubernetes-scalability-test-frameworks)

## [Working groups](https://git.k8s.io/community/sig-scalability#working-groups)

**New in 2023:**
- None

**Retired in 2023:**
- Reliability

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed
  - The meeting notes are typically kept up-to-date and comprehensive. For meeting recordings though we have been a bit sloppy admittedly (trying to improve in 2024)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scalability/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scalability/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
