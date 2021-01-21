# Production Readiness Review Process

Production readiness reviews are intended to ensure that features merging into
Kubernetes are observable, scalable and supportable, can be safely operated in
production environments, and can be disabled or rolled back in the event they
cause increased failures in production.

Production readiness reviews are done by a separate team, apart from the SIG
leads (although SIG lead approval is needed as well, of course). It is useful to
have the viewpoint of a team that is not as familiar with the intimate details
of the SIG, but *is* familiar with Kubernetes and with operating Kubernetes in
production. Experience through our dry runs in 1.17-1.20 have shown that
this slightly "outsider" view helps identify otherwise missed items.

More background may be found in the [PRR KEP].

## Status

As of 1.21, PRRs are now blocking. PRR _approval_ is required for the enhancement
to be part of the release. This means that any KEPs targeting the release for any
stage will require production readiness approval by the *Enhancements Freeze Date*.

Note that some of the questions in the [KEP template] should be answered in both
the KEP's README.md and the `kep.yaml`, in order to support automated checks on
the PRR. The template points out these as needed.

## Submitting a KEP for production readiness approval

The KEP template production readiness questionnaire should be filled out by the
KEP authors, and reviewed by the SIG leads. Once the leads are satisfied with
both the overall KEP (i.e., it is ready to move to `implementable` state) and
the PRR answers, the authors may request PRR approval:

* Assign a PRR approver from the `prod-readiness-approvers` list in the
  [OWNERS_ALIASES] file. This may be done earlier as well, to get early feedback
  or just to let the approver know. Reach out on the `#prod-readiness` Slack
  channel or just pick someone from the list. The team may rebalance the
  assignees if necessary.
* Update the `kep.yaml`, setting the `stage`, `latest-milestone`, and the
  `milestone` struct (which captures per-stage release versions).
* Create a `prod-readiness/<sig>/<KEP number>.yaml` file, with the PRR
  approver's GitHub handle for the specific stage
* See this [example PRR approval request PR].

The PRR approvers use the `kepctl` tool to identify
all outstanding requests for PRR approval, and are responsible for providing
timely feedback so that KEPs are not delayed. If your need is urgent it doesn't
hurt to ping the person on Slack, as well, but it is not necessary.

## Common feedback from reviewers

Some common issues we see:
* Missing metrics. Often metrics are overlooked, or sometimes they are defined
  but not wired up (implemented). Please see this [example metrics PR] to better
  understand how to add a metric.
* Be sure to differentiate between when to use a *metric* vs an *event*. Metrics
  are for the cluster operators and should be primarily about system status and issues,
  whereas events can be used to surface user errors. That is, in general the audience
  for metrics is cluster operator whereas the audience for events is the user (creator
  of that resource). Users won't usually see metrics, and cluster operators
  don't generally need to know about user mistakes.

## Becoming a prod readiness reviewer or approver

The prod readiness team is open and eager to add new members. The ideal
production readiness approver is someone with deep knowledge of Kubernetes overall
and the ability to think from the point of view of a cluster operator. An
excellent background would, for example, be someone that is an SRE for a fleet
of Kubernetes clusters.

To become a reviewer:
 * Inform the PRR team on Slack (`#prod-readiness`) or by attending the
   bi-weekly meeting.
 * Read/study previous PRR comments and production readiness responses in existing KEPs.
 * Choose some KEPs requiring PRR and perform a review. Put "shadow prod readiness review"
   in your review comments so that the assigned PRR approver knows your intent.
 * After at least one release cycle, if you have shown good judgement and quality reviews,
   you can propose yourself as approver by submitting a PRR to add your GitHub
   handle to the `prod-readiness-approvers` alias in [OWNERS_ALIASES].

## Finding KEPs needing prod readiness review

The prod readiness team uses the [kepctl query] command line tool to identify KEP PRs
that need review. For example:

`./cmd/kepctl/kepctl query --sig '.*' --prr '@johnbelamaric' --include-prs
--gh-token-path ~/gh-token --status implementable,provisional`

[PRR KEP]: https://git.k8s.io/enhancements/keps/sig-architecture/1194-prod-readiness
[KEP template]: https://git.k8s.io/enhancements/keps/NNNN-kep-template
[OWNERS_ALIASES]: https://git.k8s.io/enhancements/OWNERS_ALIASES
[example PRR approval request PR]: https://github.com/kubernetes/enhancements/pull/2179/files
[example metrics PR]: https://github.com/kubernetes/kubernetes/pull/97814
[kepctl query]: https://git.k8s.io/enhancements/cmd/kepctl
