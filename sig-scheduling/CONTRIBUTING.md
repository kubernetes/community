# Contributing to SIG Scheduling

Welcome to contributing to SIG Scheduling. We are excited about the prospect of you
joining our [community](https://github.com/kubernetes/community/tree/master/sig-scheduling)!

SIG Scheduling is responsible for the components that make Pod placement decisions.
You can read the SIG mission outlined in the [charter](https://git.k8s.io/community/sig-scheduling/charter.md).

There are multiple ways you can participate, including PRs, issues, documentation, new proposals,
helping to answer end-user's questions, attending meetings. All kinds of contributions are welcomed.

## Before you begin

We strongly recommend you to check out the [Kubernetes Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide)
and [Contributor Cheat Sheet](https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet) first.

## Getting Started

* If you're a newcomer and have no idea where to start, we have a non-stale pool of issues that are
available for you:
  * [first-good-issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22+label%3Asig%2Fscheduling+)
  * [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22+label%3Asig%2Fscheduling+)

* If you want to know how kube-scheduler generally works, you can read the [scheduling-eviction](https://kubernetes.io/docs/concepts/scheduling-eviction/).

* If you want to understand the architectural details of kube-scheduler, you can refer to the [design docs](https://github.com/kubernetes/community/tree/master/contributors/devel/sig-scheduling)
  or [KEPs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling).

* If you have any questions, you can find us in the Kubernetes [slack](https://app.slack.com/client/T09NY5SBT/C09TP78DV).

* If you find a bug, please open an issue under kubernetes/kubernetes with labels `/sig scheduling` and `/kind bug`,
also follow the requirements of `Bug Report`, it will help us a lot when analyzing the errors.

* If you have a feature request:
  * Open an issue under kubernetes/kubernetes with labels `/sig scheduling` and `/kind feature`, please focus on
  the user stories
  * If there's any debate on the rationalities, you can bring it to the [SIG meeting](https://github.com/kubernetes/community/tree/master/sig-scheduling#meetings).
  * If there are multiple implementation options, consider creating a doc with Pros and Cons, and gather some feedback.
  You can do this by sharing the doc with the SIG's [mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-scheduling).
  * Any feature that requires an API change or significant refactoring
    should be preceded by a [Kubernetes Enhancement Proposal (KEP)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling).

* If you find any out-of-date documentation, please help the community correct that by either sending a PR to
update the docs or open an issue if you are not sure about how to fix it.
  * For website documentation, you can open issues at [kubernetes/website](https://github.com/kubernetes/website).
  * For developer documentation, you can open issues at [kubernetes/community](https://github.com/kubernetes/community).

* We also maintain a list of [sub projects](https://github.com/kubernetes/community/tree/master/sig-scheduling#subprojects) here.
If you're interested, you can contribute to them as well.

## Best Practices

The community has been following some practices to help ensure maintainable and quality code:

* It is best if a PR is first reviewed by the [Reviewers](https://github.com/kubernetes/community/blob/master/community-membership.md#reviewer). Once the PR gets a `/lgtm` from a Reviewer,
  [Approvers](https://github.com/kubernetes/community/blob/master/community-membership.md#approver) will review the PR to approve it.

* Leave the Reviewers assigned by the bot automatically if possible, unless there is a need
for a specific contributor's expertise.

* Critical bug fixes can be assigned to approvers directly.

* Always add a new commit to address review comments instead of amending. This helps to
review the new changes. You might be asked by the Reviewer to squash at a certain point.

* Squash the commits when the PR is ready to merge, this does a great favor for the git history.

* Code contributions should be relatively small, simple, well documented and well tested.
Try to split your changes into incremental PRs if the feature is big.

* Whatever discussed offline or at the community meeting should be recorded back
to the issue/PR, which helps to preserve the context.

* Always open an issue for a TODO or a follow-up just in case you forget it.

### Technical and style guidelines

The following guidelines apply primarily to kube-scheduler, but some subprojects
might also adhere to them.

When designing a feature, think about components that depend on kube-scheduler
code, such as [cluster-autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)
or [scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins).
Also consider interactions with other core components such as [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/).

When coding:
- Follow [effective go](https://go.dev/doc/effective_go) guidelines.
- Use [contextual logging](https://git.k8s.io/community/contributors/devel/sig-instrumentation/migration-to-structured-logging.md#contextual-logging-in-kubernetes).
  Some packages might still be using [structured logging](https://git.k8s.io/community/contributors/devel/sig-instrumentation/logging.md).
- When writing APIs, follow [k8s API conventions](https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md)
- Naming:
  - Length: As a rule-of-thumb, the length of a variable name should be
    proportional to the size of the scope where it is used and inversely
    proportional to the number of times that it is used.

Testing:
- Unit tests: every change should have high coverage by unit tests.
- Integration tests: should cover interactions between the different components
  of kube-scheduler (event handlers, queue, cache, scheduling cycles) and
  kube-apiserver.
- E2E tests: should cover interactions with other components, such as kubelet,
  kube-controller-manager, etc.
- [Perf tests](https://github.com/kubernetes/kubernetes/tree/master/test/integration/scheduler_perf):
  should be considered for critical and/or CPU intensive operations.
- General guidelines:
  - Follow a [DAMP principle](https://stackoverflow.com/a/11837973).
  - Use `cmp.Diff` instead of `reflect.DeepEqual`, to provide useful comparisons.
  - Compare errors using `errors.Is` (`cmpopts.EquateErrors` when using
    `cmp.Diff`) instead of comparing the error strings.
  - Leverage existing utility functions from `pkg/scheduler/testing`.
  - Avoid creating or using assertion libraries.
    Use standard `t.Error` or `t.Fatal`, as necessary.
  - `gomega` and `ginkgo` should only be used in E2E tests.

Note that some existing code might be in violation of these guidelines, as it
might have been written before these guidelines were established. Feel free to
open PRs to get the code up to the standard.

## Use of @mentions

* @kubernetes/sig-scheduling-api-reviews - API Changes and Reviews
* @kubernetes/sig-scheduling-bugs - Bug Triage and Troubleshooting
* @kubernetes/sig-scheduling-feature-requests - Feature Requests
* @kubernetes/sig-scheduling-misc - General Discussion from Approvers and Reviewers
* @kubernetes/sig-scheduling-pr-reviews - PR Reviews
* @kubernetes/sig-scheduling-proposals - Design Proposals
* @kubernetes/sig-scheduling-test-failures - Test Failures and Triage
