# 2024 Annual Report: WG Structured Logging

## Current initiatives and Project Health


1. What work did the WG do this year that should be highlighted?

Conversion of some additional components to structured, contextual logging
proceeded, most notably in the kubelet. [Contextual
logging](https://github.com/kubernetes/enhancements/issues/3077) got promoted
to beta in Kubernetes 1.30.

Lack of support for contextual logging in client-go leads to a significant
amount of log entries without context or not even using unstructured logging,
[sometimes lacking information](https://github.com/kubernetes/klog/issues/414)
that would identify which component the log message is about.

The focus end of 2024 was on [addressing that gap in
client-go](https://github.com/kubernetes/kubernetes/pull/129125), with
hopefully all changes landing in Kubernetes 1.34 in 2025.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Once client-go is updated, we need to revisit components that were already
converted and make them use the new client-go APIs.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed:
      regular meetings continue to be on hold, we discuss topics ad-hoc on Slack
- [x] Updates provided to sponsoring SIGs in 2024
      - KubeCon EU 2024 maintainer track presentation: ["Leverage Contextual and Structured Logging in K8s for Enhanced Monitoring"](https://kccnceu2024.sched.com/event/1YhhM/leverage-contextual-and-structured-logging-in-k8s-for-enhanced-monitoring-patrick-ohly-intel-gmbh-shivanshu-raj-shrivastava-independent-mengjiao-liu-daocloud)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-structured-logging/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
