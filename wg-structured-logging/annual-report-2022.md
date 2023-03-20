# 2022 Annual Report: WG Structured Logging

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - Graduated [Deprecation of klog specific flags](https://github.com/kubernetes/enhancements/issues/2845) to GA
   - Added [Contextual logging](https://github.com/kubernetes/enhancements/issues/3077) as Alpha

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Continuously migrating kubernetes/kubernetes repository to structured, contextual logging.
     In 2022 and early 2023, 82% of kube-controller-manager were converted, over 1500 log calls.
   - Maintaining and enhancing klog. In 2022, [performance improved](https://github.com/kubernetes/kubernetes/pull/115277)
     between up to 45% depending on the component and output format.
   - Maintaining [logcheck](https://github.com/kubernetes-sigs/logtools/tree/main/logcheck).
     In 2022, it was enhanced and moved into a new repository.

## Project health

1. What's the current roadmap until completion of the working group?

   - In 2023, we need to decide if and how to integrate with
     [slog](https://pkg.go.dev/golang.org/x/exp/slog) before graduating
     contextual logging to Beta.
   - In the meantime, conversion of existing code continue.

2. Does the group have contributors from multiple companies/affiliations?

   - Yes, current members are from 1 ZTE, 1 AppDynamics, 1 Google, 3 VMware, 1 Intel and 2 unaffiliated.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - End users can easily contribute by helping with [structured, contextual logging migration](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/migration-to-structured-logging.md)

## Membership

- Primary [slack channel](https://kubernetes.slack.com/archives/C020CCMUEAX) member count: 152 (up 46% year-over-year)
- Primary [mailing list](https://groups.google.com/g/kubernetes-wg-structured-logging) member count: 56 (up 21%)
- Primary meeting attendee count (estimated, if needed): 3-4 (down by one)
- Primary meeting participant count (estimated, if needed): 2-3 (down by one)

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [ ] Updates provided to sponsoring SIGs in 2022

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-structured-logging/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
