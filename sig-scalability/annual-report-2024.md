# 2024 Annual Report: SIG Scalability

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

Notable work we produced in 2024 includes:
- Helping validate scalability/performance impact and reliability for many features across the year as well as preventing regressions
- Enhancements to our scalability test framework around better instrumentation and reduction in flakiness
- Additional test coverage spanning across the Kubernetes project
- Driving/influencing a bunch of KEPs and fixes, mainly within Kubernetes core (API server, API machinery, etcd and kube-controller-manager)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Overall in 2024 we have had healthy contributions from multiple companies (see these [devstats](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Scalability&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1672560000000&to=1704095999000).
We continue to seek help from various SIGs to force-multiply scalability test coverage and regression hunting for features/components they own. We also encourage them to proactively identify and document SLIs/SLOs/limits for APIs and workflows they own. This allows each SIG to set a scalability bar for their systems (just like up-time/availability) and thereby make scalability a first class citizen of Kubernetes and related CNCF projects. As always, SIG scalability is eager to assist/guide with this process.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

We presented SIG Scalability Intro + Deep-Dive updates at EU KubeCon:
- [KubeCon Paris](https://kccnceu2024.sched.com/event/1Yhgs/intro-deep-dive-kubernetes-sig-scalability-wojciech-tyczynski-google-shyam-jeedigunta-amazon-web-services)
  - [Recording](https://www.youtube.com/watch?v=g75sjSmdneE)

4. KEP work in 2024 (v1.30, v1.31, v1.32):

Notable KEPs, mostly co-owned with SIG API machinery and SIG etcd:
- [2340 - Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache)
  - Beta in 1.31
- [3157 - Streaming List](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list)
  - Beta in 1.32
- [4222 - CBOR Serializer](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4222-cbor-serializer)
  - Alpha in 1.32
- [4568 - Resilient watchcache initialization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4568-resilient-watchcache-initialization)
  - Beta in 1.31


## [Subprojects](https://git.k8s.io/community/sig-scalability#subprojects)


**Continuing:**
  - kubernetes-scalability-and-performance-tests-and-validation
  - kubernetes-scalability-bottlenecks-detection
  - kubernetes-scalability-definition
  - kubernetes-scalability-governance
  - kubernetes-scalability-test-frameworks

## [Working groups](https://git.k8s.io/community/sig-scalability#working-groups)


## Operational

Operational tasks in [sig-governance.md]:
- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [X] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-scalability/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-scalability/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
