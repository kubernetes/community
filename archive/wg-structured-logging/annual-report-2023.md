# 2023 Annual Report: WG Structured Logging

## Current initiatives and Project Health


1. What work did the WG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - artifacts
   - reports
   - white papers
   - work not tracked in KEPs
-->
- Graduated [contextual logging](https://github.com/kubernetes/enhancements/issues/3077) as Beta.
  - All of kube-controller-manager and some parts of kube-scheduler converted (in-tree).
- klog package updates that support using slog as a backend.
- Enhance the logcheck tool for better detection of contextual logging and use newer APIs in Kubernetes.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?
- Our working group is looking for more contributors to help us in [the migration of contextual logging](https://github.com/kubernetes/enhancements/issues/3077). 
If you are interested or experienced in the area, we would very much welcome your participation.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2023
  - [SIG Instrumentation](https://git.k8s.io/community/sig-instrumentation/)
    - [KubeCon China 2023](https://sched.co/1PTJw)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-structured-logging/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
