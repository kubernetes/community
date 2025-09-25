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

Introduced in 1.23, *Product Readiness Freeze* happens a week before *Enhancements 
Freeze Date*. All KEPs must be opted in by this date to help PRR reviewers provision 
their workload. KEPs opted in after the *Product Readiness Freeze* are at risk of not 
being reviewed by the PRR team, depending on bandwidth. KEP owners can file an 
[Exception request](https://github.com/kubernetes/sig-release/blob/master/releases/EXCEPTIONS.md) if the PRR has not been completed post Enhancements Freeze.

Note that some of the questions in the [KEP template] should be answered in both
the KEP's README.md and the `kep.yaml`, in order to support automated checks on
the PRR. The template points out these as needed.

## Submitting a KEP for production readiness approval

The KEP template production readiness questionnaire should be filled out by the
KEP authors, and reviewed by the SIG leads. Once the leads are satisfied with
both the overall KEP (i.e., it is ready to move to `implementable` state) and
the PRR answers, the authors may request PRR approval:

* Make sure the Enhancement KEP is labeled `lead-opted-in` before PRR Freeze. This is required so that the Enhancements 
  release team and PRR team are aware the KEP is targeting the release.
* Assign a PRR approver from the `prod-readiness-approvers` list in the
  [OWNERS_ALIASES] file. This may be done earlier as well, to get early feedback
  or just to let the approver know. Reach out on the `#prod-readiness` Slack
  channel or just pick someone from the list. The team may rebalance the
  assignees if necessary.
* Update the `kep.yaml`, setting the `stage`, `latest-milestone`, and the
  `milestone` struct (which captures per-stage release versions). See the 
  [KEP template](https://github.com/kubernetes/enhancements/tree/master/keps/NNNN-kep-template)
  for required sections.
* Create a `prod-readiness/<sig>/<KEP number>.yaml` file, with the PRR
  approver's GitHub handle for the specific stage
* See this [example PRR approval request PR].

At the beginning of each Kubernetes development cycle, a project board is created
for that release at the [GitHub Project Page](https://github.com/orgs/kubernetes/projects).
This project board contains a "PRR" tab that lists all the current KEPs and
adds PRR metadata such as PRR assignee, shadow, status and PRR-specific notes.

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

### Becoming an approver

After serving as reviewer/shadow for at least one release and showing good judgement and quality reviews,
you can propose yourself as an approver by submitting a PR to add your GitHub
handle to the `prod-readiness-approvers` alias in [OWNERS_ALIASES].

When submitting the PR, you should include links to your KEP review comments that demonstrate a good variety
of different situations.
Here is a good starting point (remember that one PR can cover multiple categories):

* Transitions from new to alpha
* Transitions from alpha to beta
* Transitions from beta to GA
* Must have successfully reviewed at least three enhancements that require coordination between multiple components.
* Must have successfully reviewed at least three enhancements that require version skew consideration (both HA and component skew):
  does behavior fail safely and eventually reconcile.
* Must have successfully reviewed at least three enhancements that are outside your primary domain.
* Examples where the feature requires considering the case of administering thousands of clusters.
  This comes up frequently for host-based features in storage, node, or networking.
* Examples where the feature requires considering the case of very large clusters.  This is commonly covered by metrics.

The success in all of the above can be claimed if the final PRR approver perceived their review
as non-essential after the primary PRR review done by the shadow reviewer.

*Watchout*: The fact that final approver didn't have any substantial comments may mean one of two things:

* the reviewer did a great job in finding all the issues
* the KEP author did a great job in answering all the questions up front

When promoting shadow reviewer to approver we want to ensure that not all KEPs belong to the second
category to ensure that all PRR approvers can push for changes if required. We definitely don't
require all reviews to be in that category, but we require at least three to belong there.

## Finding KEPs needing prod readiness review

The status of prod readiness progress on KEPs is tracked on a project board
specific to the current Kubernetes release. The board typically follows
naming scheme `[1.XX] Enhancements Tracking`, where `1.XX` represents the current
Kubernetes release version.

* \[PRR KEP\]: https://git.k8s.io/enhancements/keps/sig-architecture/1194-prod-readiness
* \[KEP template\]: https://git.k8s.io/enhancements/keps/NNNN-kep-template
* \[OWNERS_ALIASES\]: https://git.k8s.io/enhancements/OWNERS_ALIASES
* \[example PRR approval request PR\]: https://github.com/kubernetes/enhancements/pull/2179/files
* \[example metrics PR\]: https://github.com/kubernetes/kubernetes/pull/97814
