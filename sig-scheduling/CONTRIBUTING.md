# Contributing to SIG Scheduling

Welcome to contributing to SIG Scheduling. We are excited about the prospect of you
joining our [community](https://github.com/kubernetes/community/tree/master/sig-scheduling)!

SIG Scheduling is responsible for the components that make Pod placement decisions.
You can read the SIG mission outlined in the [charter](https://git.k8s.io/community/sig-cluster-lifecycle/charter.md).

We have multi areas you can contribute to, PRs, issues, documentations, new ideas,
helping to answer end-user's questions, attending meetings, all kinds of contributions are welcomed.

## Before you begin

We strongly recommend you to check out the [Kubernetes Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide)
and [Contributor Cheat Sheet](https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet).

## Getting Started

* If you're a newcomer and have no idea where to start, we have a non-stale pool of issues that are
available for you:
  * [first-good-issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22+label%3Asig%2Fscheduling+)
  * [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22+label%3Asig%2Fscheduling+)

* If you want to understand the architectural details of scheduler, you can refer to the [design docs](https://github.com/kubernetes/community/tree/master/contributors/devel/sig-scheduling)
  or [KEPs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling).

* If you have any questions, you can find us in the [slack](https://app.slack.com/client/T09NY5SBT/C09TP78DV).

* If you find a bug, please open an issue under kubernetes/kubernetes with labels `/sig scheduling` and `/kind bug`,
also follow the requirements of `Bug Report`, it will help us a lot when analyzing the errors.

* If you have a feature request:
  * Open an issue under kubernetes/kubernetes with labels `/sig scheduling` and `/kind feature`, please focus on
  the user stories
  * If there's any debate on the rationalities, you can bring it to the [SIG meeting](https://github.com/kubernetes/community/tree/master/sig-scheduling#meetings).
  * If there're multiple implementation options, consider to create a doc with Pros and Cons, and gather some feedbacks.
  * If necessary, open an issue in kubernetes/enhancements and write a KEP for it.

* If you find any out-of-date documentations, please help us to revise that, the readers come latter
will be grateful for your kindness.

* We also maintain a list of [sub projects](https://github.com/kubernetes/community/tree/master/sig-scheduling#subprojects) here,
if you're interest, you can contribute to them at your wish.

## Best Practices

From years contributions to the Kubernetes, we summarized several practices here:

* In general, PR should be reviewed by the Reviewers firstly, once it got a `/lgtm`,
  approvers will get onboard. This helps to reduce the burden on the Approvers.

* Leave the Reviewers assigned by the bot automatically if possible, unless there is a need
for one specific contributor's expertise.

* Critical bug fixes can be assigned to approvers directly.

* Always add a new commit to address review comments, instead of amending, this helps to
review the new changes unless you're asked by the Reviewer or it's ready to merge.

* Squash the commits when PR is ready to merge, this does a great favor for the git history.

* Code contributions should be relatively small, simple, well documented and well tested.

* Whatever discussed offline or at the community meeting should be recorded back
to the issue/PR, which helps to preserve the context.

* Always open an issue for a TODO or a follow-up just in case you forget it.

## Use of @mentions

* @kubernetes/sig-scheduling-api-reviews - API Changes and Reviews
* @kubernetes/sig-scheduling-bugs - Bug Triage and Troubleshooting
* @kubernetes/sig-scheduling-feature-requests - Feature Requests
* @kubernetes/sig-scheduling-misc - General Discussion
* @kubernetes/sig-scheduling-pr-reviews - PR Reviews
* @kubernetes/sig-scheduling-proposals - Design Proposals
* @kubernetes/sig-scheduling-test-failures - Test Failures and Triage
