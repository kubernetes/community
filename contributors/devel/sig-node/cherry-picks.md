# SIG Node Cherry-Pick Management Process

Kubernetes supports [multiple release versions] simultaneously. How do new
fixes make it into earlier versions, when we only develop against one version
at a time?

We [cherry-pick] them!

To assist the Release Team, a member of SIG Node can volunteer to oversee
Node's cherry-picks. This guide is intended to document that person's
responsibilities.

[multiple release versions]: https://kubernetes.io/docs/setup/release/version-skew-policy/
[cherry-pick]: /contributors/devel/sig-release/cherry-picks.md

## Release timeline

SIG Release regularly (on a ~monthly cadence) ships [patch releases]. Every
release, the Release Team publishes a list of [target dates] with **cherry-pick
deadlines**.

SIG Node's responsibility is to ensure that all patches that we want
cherry-picked to previous releases must be [triaged], prioritized, LGTM'd, and
approved by the deadline.

[patch releases]: https://github.com/kubernetes/sig-release/blob/master/release-engineering/versioning.md#patch-releases
[target dates]: https://kubernetes.io/releases/patch-releases/#upcoming-monthly-releases
[triaged]: triage.md

## Monthly cherry-pick process

Throughout the release cycle, merged bugfixes should be evaluated for inclusion
as cherry-picks. Any bugs marked **critical-urgent** should be considered. Some
**important-soon** bugs may also be considered if they are sufficiently
contained and meet the criteria above.

**Only** bug fixes should be considered for backports. Patches should be merged
and soaked in CI for at least a week.

To create a cherry-pick PR, you can use the `./hack/cherry_pick_pull.sh`
script. Read through the [cherry pick documentation][cherry-pick] for details
on that process.

## Preparing the SIG's patches

The week of the cherry-pick deadline, create a Slack thread ([example
thread]) in the [#sig-node] channel to track the cherry-picks.

To determine what fixes need to be cherry-picked, you can use the [SIG Node PR
Board with the cherry-pick filter view]. Track each change considered for
inclusion individually.

Release branches are typically kept in sync. This means that bugfixes should be
backported to all applicable supported branches. The [cherry pick
schedule][target dates] will include a list of all supported releases and their
end-of-life (EOL) dates.

In the first comment in the Slack thread, track each patch and the status of
backports to each branch. Some examples of what you might write are available
below for reference. You can use the rest of the Slack thread for discussion.

Typically, we will backport 2-3 patches per monthly cycle.

[example thread]: https://kubernetes.slack.com/archives/C0BP8PW9G/p1617919799137500
[#sig-node]: https://kubernetes.slack.com/messages/sig-node
[SIG Node PR Board with the cherry-pick filter view]: https://github.com/orgs/kubernetes/projects/49?card_filter_query=label%3Ado-not-merge%2Fcherry-pick-not-approved

### Example cherry-pick: all supported branches

Consider the following PR to backport:
https://github.com/kubernetes/kubernetes/pull/99600

This fixes broken accounting in a beta feature that has been beta for the past
three releases.

Since it is a **critical-urgent** fix, it is eligible to be backported to all
affected releases. During the 1.21 cycle, that means 1.18 through 1.20.

Thus, our Slack update should look like the following:

> Count pod overhead as an entity's resource usage https://github.com/kubernetes/kubernetes/pull/99600
> - 1.18 https://github.com/kubernetes/kubernetes/pull/100039
> - 1.19 https://github.com/kubernetes/kubernetes/pull/100038
> - 1.20 https://github.com/kubernetes/kubernetes/pull/100037

You can use this list and the Node Triage board to ensure all PRs have LGTMs
and approvals, so they are ready for the Release Team by the cherry-pick
deadline.

### Example cherry-pick: limited release support

Consider the following PR to backport:
https://github.com/kubernetes/kubernetes/pull/98088

This fixes an issue with an alpha feature introduced in 1.20, so it is only
eligible to backport to 1.20 during the 1.21 cycle.

In this case,

- the fix only affects feature-gated code
- the fix addresses a serious bug (repeatedly acquiring a lock unnecessarily)
- the fix is small

Hence, we included it in our backport.

Thus, our Slack update should look like the following:

> Fix repeatedly acquire the inhibit lock https://github.com/kubernetes/kubernetes/pull/98088
> - 1.18 N/A, only affects 1.20+
> - 1.19 N/A, only affects 1.20+
> - 1.20 https://github.com/kubernetes/kubernetes/pull/99255

### Example cherry-pick: rejected

Consider the following PR to backport:
https://github.com/kubernetes/kubernetes/pull/98376

In this case,

- the fix is a relatively involved refactor
- the fix caused a known test regression
- the regression fix is not being backported
- the priority of the fix was **important-longterm**

Hence, this would **not** be an appropriate patch to cherry-pick. SIG Node
chose to close its cherry-pick PRs.
