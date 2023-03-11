# 2022 Annual Report: SIG Instrumentation

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Tracing support for the Kubernetes api-server and kubelet have graduated to beta.
   - We've introduced a KEP for Kubernetes control-plane component SLIs.
   - The metrics stability framework is being extended and now support `BETA` stability levels and auto-generated documentation.
   - Contextual logging has graduated to beta.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Bi-weekly triage meeting
   - Subprojects (see below)
   - Mentorship Program

3. KEP work in 2022 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->
   - Stable
     - [2845 - Deprecate klog specific flags in Kubernetes components](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/2845-deprecate-klog-specific-flags-in-k8s-components) - v1.26
     - [1748 - Expose metrics about resource requests and limits that represent the pod model](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1748-pod-resource-metrics) - v1.27

   - Beta
     - 1.27 Metrics stability enhancement
     - 1.27 Kubernetes component health SLIs
     - 1.27 Apiserver tracing 
     - 1.27 Kubelet tracing
     - 1.27 Kubelet Resource Metrics Endpoint
   - Alpha
     - 1.26 Metrics stability enhancement
     - 1.26 Kubernetes component health SLIs
     - 1.25 Kubelet tracing

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - [kubernetes-sigs/prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/OWNERS_ALIASES) has 1 active approver
   - [kubernetes-sigs/custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver/blob/master/OWNERS) has 1 active approver
   - [kubernetes-sigs/metrics-server](https://github.com/kubernetes-sigs/metrics-server/blob/master/OWNERS) has 2 approvers but are both outdated

2. What metrics/community health stats does your group care about and/or measure?

   - Devstats
     - [Review Load](https://k8s.devstats.cncf.io/d/80/pr-workload-per-sig-and-repository-chart?orgId=1&var-sigs=%22instrumentation%22&var-repo_name=kubernetes%2Fkubernetes&var-repo=kuberneteskubernetes&from=now-1y&to=now) 
     - [Time to Approve and Merge](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Instrumentation&var-repo_name=kubernetes%2Fkubernetes&var-apichange=All&var-size_name=All&var-kind_name=All) 
   - Meeting attendance
       - Meeting attendance is ~12 each week
       - Triage attendance is ~8 each week
   - Enhancement velocity
       - We have varying levels of velocity depending on our KEPs. Some graduate in successive releases while other may take more than 1 release to promote in level.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?
   - We certainly hope so.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We have a mentorship program and now have enrolled mentees/mentors.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes: Google, Red Hat, Sony, VMware, Intel, independent contributors, and more

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

    - SIG leads performed a [staffing and gap analysis](https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit#heading=h.849b7ydpl7ip) for the SIG's projects. We definitely can use more help, and need more part-time/full-time contributors.

### Summary

- **KTLO:** 4 leads at 5%, 2 maintainers at 10% for core, 9-10 maintainers at 5% for 5 subprojects (can have overlap between roles, but need a minimum of 3-4 part-time contributors at 25%)
- **Feature work:** needs significant ongoing additional investment, minimum of 2-3 FT devs or features will continue to slip

### Details

- **KTLO:** Requires two experienced part-time maintainers at 10% (e.g. 2x4h = 8h/wk) in addition to the SIG leadership (chairs/TLs @ minimum of 2h/wk)
  - Chairs/TLs currently perform the bulk of this work but even amongst the four of them, **do not** have 8h total weekly allocated
- **Feature work:** requires significantly more investment from development and review time. 
  - E.g. Structured logging initiative requires a minimum of 2 FT staff for the duration of feature development from beta -> GA
  - SIG currently owns a number of KEPs stuck in alpha/beta due to lack of dev resources:
      - [metric cardinality enforcement](https://github.com/kubernetes/enhancements/issues/2305)

### Subprojects

- Subprojects are currently mature/stable and mainly have KTLO needs
- kube-state-metrics: 2-3 experienced maintainers at 5%, currently staffed
- Metrics-server: 1 experienced maintainers at 5%, **under staffed**
- Custom-metrics-apiserver: 1 experienced maintainer at 5%, currently staffed
- Usage-metrics-collector: 4 experienced maintainers at 5%, currently staffed
- Klog: 2 experienced maintainers at 5%, has **no current staffing**
- Prometheus-adapter: 1 experienced maintainers at 5%, has **no current staffing**


## Membership

- Primary slack channel member count: 1,985
- Primary mailing list member count: 358
- Primary meeting attendee count (estimated, if needed): 10-12
- Primary meeting participant count (estimated, if needed): 7-8
- Unique reviewers for SIG-owned packages: 18
- Unique approvers for SIG-owned packages: 17

Include any other ways you measure group membership

- We track active members in the SIG primarily based on devstats and meeting participation, and maintain an up-to-date roster of members in [kubernetes/org](https://github.com/kubernetes/org/blob/main/config/kubernetes/sig-instrumentation/teams.yaml). These teams also serve as aliases for GitHub pings.

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2022:
- [usage-metrics-collector](https://github.com/kubernetes-sigs/usage-metrics-collector)

Continuing:
- [instrumentation](https://github.com/kubernetes-sigs/instrumentation)
- [instrumentation-addons](https://github.com/kubernetes-sigs/instrumentation-addons)
- [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics)
- [metrics](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/metrics)
- [custom-metrics-apiserver](https://github.com/kubernetes-sigs/custom-metrics-apiserver)
- [metrics-server](https://github.com/kubernetes-sigs/metrics-server)
- [prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter)
- [klog](https://github.com/kubernetes/klog)
- [structured-logging](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/component-base/logs)
- [metric-stability-framework](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/component-base/metrics)
- [instrumentation-tools](https://github.com/kubernetes-sigs/instrumentation-tools)
- [log-tools](https://github.com/kubernetes-sigs/logtools)

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing in 2022:
- WG Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [KubeCon EU 2022 SIG Update](https://www.youtube.com/watch?v=xxG0-Ex6bjM)
      - [KubeCon NA 2022 SIG Update](https://www.youtube.com/watch?v=JIzrlWtAA8Y)
      - [Youtube Playlist for recurring meetings](https://www.youtube.com/playlist?list=PL69nYSiGNLP1tue6RXLncPTGjfnBVHP-f)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-instrumentation/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-instrumentation/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
