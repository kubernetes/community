# 2025 Annual Report: SIG Instrumentation

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->
  - Han Kang (logicalhan) stepped down as chair, Richa Banker (richabanker) added as new chair
  - 2025 Contributor Awards:
    - Yongrui Lin (yongruilin)
  - KEPs that moved forward
    - `API Server Tracing` graduated to [GA](https://github.com/kubernetes/kubernetes/pull/132340)
    - `Kubelet Tracing` graduated to [GA](https://github.com/kubernetes/kubernetes/pull/132341)
  - Announced deprecations
    - `Prometheus Adapter` deprecation [announced](https://github.com/kubernetes-sigs/prometheus-adapter/issues/701)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?
  - [kubernetes/kube-state-metrics](https://github.com/kubernetes/kube-state-metrics)
  - [kubernetes-sigs/custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver/blob/master/OWNERS)
  - [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/blob/master/OWNERS)

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->
  - [SIG Instrumentation Introduction and Deep Dive (KubeCon EU '25)](https://www.youtube.com/watch?v=RwcC44BWDvA&list=PLj6h78yzYM2MP0QhYFK8HOb8UqgbIkLMc&index=206&pp=iAQB)
  - [SIG Instrumentation Introduction and Deep Dive (KubeCon NA '25)](https://www.youtube.com/watch?v=cr-_IRp5qX0)

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->
  - Stable
    - [2831 - Kubelet OpenTelemetry Tracing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2831-kubelet-tracing) - v1.34
    - [647 - APIServer Tracing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/647-apiserver-tracing) - v1.34
  - Beta
    - [3077 - Contextual Logging](https://github.com/kubernetes/enhancements/pull/5621) more packages (e.g. kubelet components and some client-go packages) were migrated to contextual logging
    - [5270 - Graduate custom.metrics.k8s.io to stable](https://github.com/kubernetes/enhancements/pull/5269) 
    - [5271 - Graduate external.metrics.k8s.io to stable](https://github.com/kubernetes/enhancements/pull/5268)
  - Alpha
    - [4827 - ComponentStatusz](https://github.com/kubernetes/kubernetes/pull/134313) converted to a versioned structured API
    - [4828 - ComponentFlagz](https://github.com/kubernetes/kubernetes/pull/134995) converted to a versioned structured API


## [Subprojects](https://git.k8s.io/community/sig-instrumentation#subprojects)


**New in 2025:**
  - resource-state-metrics
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
  - usage-metrics-collector

## [Working groups](https://git.k8s.io/community/sig-instrumentation#working-groups)

**Continuing:**
 - Serving
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-instrumentation/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-instrumentation/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
