# 2021 Annual Report: SIG Instrumentation

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Formed WG Structured Logging. Successfully migrated multiple components to structured logs and graduated feature to beta.
   - Added tracing support to the Kubernetes API server and began work on Kubelet tracing.
   - Graduated the metrics stability framework.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Bi-weekly triage meeting
   - Subprojects (see below)

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [1209 - Metrics Stability](https://git.k8s.io/enhancements/...) - 1.21
     - [1933 - Prevent logging secrets via static analysis](https://git.k8s.io/community/$link/README.md) - 1.23
   - Beta
     - [1602 - Structured Logging](https://git.k8s.io/community/$link/README.md) - 1.23
     - [1748 - Pod resource requests/limits metrics](https://git.k8s.io/community/$link/README.md) - 1.22
   - Alpha
     - [2305 - Metrics Cardinality Enforcement](https://git.k8s.io/community/$link/README.md) - 1.21
     - [647 - API Server Tracing](https://git.k8s.io/community/$link/README.md) - 1.22
     - [2845 - Deprecate klog-specific flags in k8s components](https://git.k8s.io/community/$link/README.md) - 1.23
   - Pre-alpha
     - [2831 - Kubelet OpenTelemetry Tracing](https://git.k8s.io/community/$link/README.md) - alpha in 1.24

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [kubernetes-sigs/custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver/blob/master/OWNERS) has 1 active approver
   - [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/blob/master/OWNERS) has 1 active approver
   - [kubernetes-sigs/prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/OWNERS_ALIASES) has 1 active approver

2. What metrics/community health stats does your group care about and/or measure?

   - Devstats
     - [Review Load](https://k8s.devstats.cncf.io/d/80/pr-workload-per-sig-and-repository-chart?orgId=1&var-sigs=%22instrumentation%22&var-repo_name=kubernetes%2Fkubernetes&var-repo=kuberneteskubernetes&from=now-1y&to=now) has been relatively level over the year
     - [Time to Approve and Merge](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Instrumentation&var-repo_name=kubernetes%2Fkubernetes&var-apichange=All&var-size_name=All&var-kind_name=All) has been relatively low over the last year.
   - Meeting attendance
       - Meeting attendance is ~10 each week
       - Triage attendance is ~5 each week
   - Enhancement velocity
       - 7 Active enhancements over the year

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   -

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A, we don't have special training or requirements

5. Does the group have contributors from multiple companies/affiliations?

   - Yes: Red Hat, Google, Intel, ...

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Yes
   -

## Membership

- Primary slack channel member count: 1,740
- Primary mailing list member count: 324
- Primary meeting attendee count (estimated, if needed): 8-12
- Primary meeting participant count (estimated, if needed): 4-6
- Unique reviewers for SIG-owned packages: 24 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 22 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

- We track active members in the SIG primarily based on devstats and meeting participation, and maintain an up-to-date roster of members in [kubernetes/org](https://github.com/kubernetes/org/blob/main/config/kubernetes/sig-instrumentation/teams.yaml). These teams also serve as aliases for GitHub pings.

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [instrumentation](https://github.com/kubernetes-sigs/instrumentation)
- [instrumentation-addons](https://github.com/kubernetes-sigs/instrumentation-addons)

Retired in 2021:
- [heapster](https://github.com/kubernetes-retired/heapster)
- [mutating-trace-admission-controller](https://github.com/kubernetes-retired/mutating-trace-admission-controller)

Continuing:
- [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics)
- [metrics](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/metrics)
- [custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver)
- [metrics-server](https://github.com/kubernetes-sigs/metrics-server)
- [prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter)
- [klog](https://github.com/kubernetes/klog)
- [structured-logging](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/component-base/logs)
- [metric-stability-framework](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/component-base/metrics)
- [instrumentation-tools](https://github.com/kubernetes-sigs/instrumentation-tools)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- [WG Structured Logging](https://git.k8s.io/community/$wg-id/) ([$YYYY report](https://git.k8s.io/community/$wg-id/annual-report-$YYYY.md))

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon EU 2021 SIG Update](https://sched.co/iE8c)
      - [KubeCon NA 2021 SIG Update](https://sched.co/lV72)
      - [Structured logging targeting 1.21 beta (kubernetes-dev@)](https://groups.google.com/g/kubernetes-dev/c/vjSqUtPO0hs/m/wF91qunnBQAJ)
      - [WG Creation Request: WG Structured Logging (kubernetes-dev@)](https://groups.google.com/g/kubernetes-dev/c/y4WIw-ntUR8/m/NaQHu1cnAwAJ)
      - [Deprecation: Dynamic log sanitization removal in 1.24 (kubernetes-dev@)](https://groups.google.com/g/kubernetes-dev/c/xhQuwdd2Smw/m/L_EyHKbDAAAJ)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-instrumentation/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-instrumentation/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md