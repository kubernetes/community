# 2024 Annual Report: SIG Instrumentation

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- `Contextual Logging` graduated to [beta](https://github.com/kubernetes/enhancements/pull/4219)
- Introduced `Component Statusz` [alpha](https://github.com/kubernetes/enhancements/pull/4830)
- Introduced `Component Flagz` [alpha](https://github.com/kubernetes/enhancements/pull/4831)
- Final draft ready for `Resource State Metrics` [alpha](https://github.com/kubernetes/enhancements/pull/4811)
- [Usage metrics collector](https://github.com/kubernetes-sigs/usage-metrics-collector) now supports [cgroupsv2](https://github.com/kubernetes-sigs/usage-metrics-collector/pull/140) in metrics sampler
- [Usage metrics collector](https://github.com/kubernetes-sigs/usage-metrics-collector) had 6 contributors from 4 different companies

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- [kubernetes-sigs/custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver/blob/master/OWNERS)
- [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/blob/master/OWNERS)
- [kubernetes-sigs/prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/OWNERS_ALIASES)

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- [SIG Instrumentation Introduction and Deep Dive (KubeCon NA '24)](https://youtu.be/IAGjj4s3F_M?si=dWCcua8XROpgqUZZ&t=0)

- [SIG Instrumentation Introduction and Deep Dive (KubeCon EU '24)](https://www.youtube.com/watch?v=Sx1jmIJhfyA&list=PLj6h78yzYM2N8nw1YcqqKveySH6_0VnI0&index=106)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

- Alpha
  - [4827 - Component Statusz](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/4827-component-statusz) - v1.32
  - [4828 - Component Flagz](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/4828-component-flagz) - v1.32

- Beta
  - [3077 - Contextual logging](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3077-contextual-logging) - v1.30

- Stable
  - [2305 - Dynamic Cardinality Enforcement](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2305-metrics-cardinality-enforcement) - v1.30

## [Subprojects](https://git.k8s.io/community/sig-instrumentation#subprojects)

**Continuing:**

- custom-metrics-apiserver
- instrumentation-addons
- instrumentation-tools
- klog
- kube-state-metrics
- metrics
- metrics-server
- prometheus-adapter
- usage-metrics-collector

## [Working groups](https://git.k8s.io/community/sig-instrumentation#working-groups)

**Continuing:**

- Serving
- Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in [devel] dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed

[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-instrumentation/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
