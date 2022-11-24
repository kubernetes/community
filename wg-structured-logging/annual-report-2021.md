# 2021 Annual Report: WG Structured Logging

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - Graduated [Structured Logging](https://github.com/kubernetes/enhancements/issues/1602) to Beta
   - Graduated [Deprecation of klog specific flags](https://github.com/kubernetes/enhancements/issues/2845) to Beta
   - Created prototype and KEP for [Contextual logging](https://github.com/kubernetes/enhancements/pull/3078)

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Continuously migrating kubernetes/kubernetes repository to structured logging. In 2021 kubelet, kube-scheduler, kube-proxy were migrated.


## Project health

1. What's the current roadmap until completion of the working group?

   - Graduate [Contextual Logging](https://github.com/kubernetes/enhancements/issues/3077) to Beta and GA
   - Graduate [Deprecation of klog specific flags](https://github.com/kubernetes/enhancements/issues/2845) to GA
   - Graduated [Structured Logging](https://github.com/kubernetes/enhancements/issues/1602) to GA
   - All code in kubernetes/kubernetes repository is migrated to Structured Logging API

2. Does the group have contributors from multiple companies/affiliations?

   - Yes, current members are from 1 ZTE, 1 AppDynamics, 1 Google, 3 VMware, 1 Intel and 2 unaffiliated.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - End users can easily contribute by helping with [structured logging migration](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/migration-to-structured-logging.md)
   -

## Membership

- Primary slack channel member count: 104
- Primary mailing list member count: 46
- Primary meeting attendee count (estimated, if needed): 4-5
- Primary meeting participant count (estimated, if needed): 3-4

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [ ] Updates provided to sponsoring SIGs in 2021
      - [$sig-name](https://git.k8s.io/community/$sig-id/)
        - links to email, meeting notes, slides, or recordings, etc
      - [$sig-name](https://git.k8s.io/community/$sig-id/)
        - links to email, meeting notes, slides, or recordings, etc
      -

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-structured-logging/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml

