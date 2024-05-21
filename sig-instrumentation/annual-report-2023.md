# 2023 Annual Report: SIG Instrumentation

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- Started off a [SIG Instrumentation Mentorship Program](https://docs.google.com/document/d/1Qa0KflaD2H1AbBtC--c4UD_xaxVkB0zb1UOBz7Vbu14/edit#heading=h.dxk9okw7st8s) to attract and retain new contributors.
- Introduced a new subproject, [usage-metrics-collector](https://github.com/kubernetes-sigs/usage-metrics-collector), to collect kube usage and capacity metrics.
- `APIServer Tracing` graduated to [stable](https://github.com/kubernetes/enhancements/commit/97713189b3107b41c4c19505d04aa7ef22df063b).
- `Dynamic Cardinality Enforcement` graduated to [stable](https://github.com/kubernetes/enhancements/commit/ab798d7c2f9a75c770dc4369be702036b984b40e).
- `Extending Metrics Stability` graduated to [beta](https://github.com/kubernetes/enhancements/commit/5c9771693c9820176ff37854f4727bad0889b492).
- `Kubernetes Components Health SLIs` graduated to [stable](https://github.com/kubernetes/enhancements/commit/6ede6f1f2163957f5f10033a38fe83804d0df671).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

- [kubernetes-sigs/custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver/blob/master/OWNERS)
- [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/blob/master/OWNERS)
- [kubernetes-sigs/prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/OWNERS_ALIASES)

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- [SIG Instrumentation Introduction and Deep Dive (KubeCon NA '23)](https://youtu.be/Lf5h8bPrSBM?si=HOv63HqVRIAB2mjw)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Beta
    - [2305 - Dynamic Cardinality Enforcement](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2305-metrics-cardinality-enforcement) - v1.28
    - [647 - APIServer Tracing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/647-apiserver-tracing) - v1.27

  - Stable
    - [1748 - Expose Pod Resource Request Metrics](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1748-pod-resource-metrics) - v1.27
    - [2831 - Kubelet OpenTelemetry Tracing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2831-kubelet-tracing) - v1.28
    - [3466 - Kubernetes Component Health SLIs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3466-kubernetes-component-health-slis) - v1.29
    - [3498 - Extending Metrics Stability](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3498-extending-stability) - v1.28

## [Subprojects](https://git.k8s.io/community/sig-instrumentation#subprojects)

**New in 2023:**
  - [usage-metrics-collector](https://git.k8s.io/community/<no value>#usage-metrics-collector)

**Continuing:**
  - custom-metrics-apiserver
  - instrumentation
  - instrumentation-addons
  - instrumentation-tools
  - klog
  - kube-state-metrics
  - metric-stability-framework
  - metrics
  - metrics-server
  - prometheus-adapter
  - structured-logging

## [Working groups](https://git.k8s.io/community/sig-instrumentation#working-groups)

**Continuing:**
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-instrumentation/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-instrumentation/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
