# 2022 Annual Report: WG Batch

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - We're gathering feedback from multiple users running various batch-related frameworks and projects on kubernetes, see [our agenda](https://docs.google.com/document/d/1XOeUN-K0aKmJJNq7H07r74n-mGgSFyiEDQ3ecwsGhec/edit) for both past and upcoming topics.
   - Closely collaborated with SIG Apps, SIG Node and SIG Scheduling to improve batch workloads on kubernetes:
      - [2307 - Job tracking without lingering Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/2307-job-tracking-without-lingering-pods)
      - [3329 - Retriable and non-retriable Pod failures for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3329-retriable-and-non-retriable-failures)
   - [Introduction to Batch WG during KubeCon EU 2022](https://www.youtube.com/watch?v=XeX2zBOykC4)
   - Create [Kueue: A Kubernetes-native Job Queueing](https://www.youtube.com/watch?v=YwSZUdU3iRY) project and release [2 major versions in 2022](https://github.com/kubernetes-sigs/kueue/releases).
   - Organize Batch Days during [KubeCon EU 2022](https://kubernetesbatchdayeu22.sched.com/) and [KubeCon NA 2022](https://kubernetesbatchdayna22.sched.com/)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Continuing work on [Kueue](https://github.com/kubernetes-sigs/kueue/releases).
   - Continuing work on Topology-aware Scheduling by enhancing [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)
     and [Scheduler plugins](https://github.com/kubernetes-sigs/scheduler-plugins) to enable support for specialized hardware in Kubernetes.

## Project health

1. What's the current roadmap until completion of the working group?

   - We're working on a document, that will be shared with the entire work group, about the exit criteria for the group.

2. Does the group have contributors from multiple companies/affiliations?

   - Yes, see [our agenda](https://docs.google.com/document/d/1XOeUN-K0aKmJJNq7H07r74n-mGgSFyiEDQ3ecwsGhec/edit) for variety of presentations we've had over the past year.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Everyone are welcome to present any topic related to batch, so at this point in time there's no need for additional support.

## Membership

- Primary slack channel member count: 217
- Primary mailing list member count: 152
- Primary meeting attendee count (estimated, if needed): 20
- Primary meeting participant count (estimated, if needed): 10

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2022
   - [Introduction to Kubernetes WG Batch @ KubeCon EU 2022](https://youtu.be/XeX2zBOykC4)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-batch/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
