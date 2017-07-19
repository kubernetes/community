# Patch Release Manager Playbook

This is a playbook intended to guide new patch release managers.
It consists of opinions and recommendations from former patch release managers.

Note that patch release managers are ultimately responsible for carrying out
their [duties](README.md#patch-release-manager) in whatever manner they deem
best for the project.
The playbook is more what you call "guidelines" than actual rules.

## Getting started

* Add yourself to the [Release Manager table](https://github.com/kubernetes/community/wiki)
  so the community knows you're the point of contact.
* Ask a maintainer to add you to the [kubernetes-release-managers](https://github.com/orgs/kubernetes/teams/kubernetes-release-managers/members)
  team so you have write access to the main repository.
* Ask to be added to the [kubernetes-security](https://groups.google.com/forum/#!forum/kubernetes-security)
  mailing list.
* Ask to be given access to post to the [kubernetes-announce](https://groups.google.com/forum/#!forum/kubernetes-announce)
  and [kubernetes-dev-announce](https://groups.google.com/forum/#!forum/kubernetes-dev-announce)
  mailing lists.
* Sync up with the outgoing release branch manager to take ownership of any
  lingering issues on the branch.
* Run [anago](https://github.com/kubernetes/release) in mock mode to get prompts
  for setting up your environment, and familiarize yourself with the tool.

## Cherrypick requests

As a patch release manager, you are responsible for reviewing
[cherrypicks](../cherry-picks.md) on your release branch.

You can find candidate PRs in the [cherrypick queue dashboard](http://cherrypick.k8s.io/#/queue).
Once a cherrypick PR is created and ready for your review, it should show up in
a GitHub search such as [`is:pr is:open base:release-1.6`](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr%20is%3Aopen%20base%3Arelease-1.6).

As an example of the kind of load to expect, there were about 150 cherrypick PRs
against the `release-1.6` branch in the 3 months between v1.6.0 and v1.7.0.

For each cherrypick request:

1.  **Decide if it meets the criteria for a cherrypick**

    Make sure the PR author has supplied enough information to answer:

    * What bug does this fix?
      (e.g. *feature X was already launched but doesn't work as intended*)
    * What is the scope of users affected?
      (e.g. *anyone who uses feature X*)
    * How big is the impact on affected users?
      (e.g. *pods using X fail to start*)
    * How have you verified the fix works and is safe?
      (e.g. *added new regression test*)

    Ask the PR author for details if these are missing and not obvious.
    If you aren't sure what to do, escalate to the relevant SIGs.

    **Notes**

    * Version bumps (e.g. v0.5.1 -> v0.5.2) for dependencies with their own
      release cycles (e.g. kube-dns, autoscaler, ingress controllers, etc.)
      deserve special attention because it's hard to see what's changing.
      In the past, such bumps have been a significant source of regressions in
      the stable release branch.

      Check the release notes for the dependency to make sure there are no new
      behaviors that could destabilize the release branch.
      Ideally you should only accept version bumps whose release deltas contain
      only changes that you would have approved individually, if they had been
      part of the Kubernetes release cycle.

      However, this gets tricky when there are fixes you need for your branch
      that are tied up with other changes. Ask the cherrypick requester for
      context on the other changes and use your best judgment.

    * Historically (up through at least 1.6), patch release managers have
      occasionally granted exceptions to the "no new features" rule for
      cherrypicks that are confined to plugins like cloudproviders
      (e.g. vSphere, Azure) and volumes (e.g. Portworx).

      However, we required that these exceptions be approved by the plugin
      owners, who were asked to `/approve` through the normal `OWNERS` process
      (despite it being a cherrypick PR).

1.  **Make sure it has an appropriate release note**

    [Good release notes](https://github.com/kubernetes/community/issues/484)
    are particularly important for patch releases because cluster admins expect
    the release branch to remain stable and need to know exactly what changed.
    Take care to ensure every cherrypick that deserves a release note has one
    *before you approve it* or else the change may fall through the cracks at
    release cut time.

    Also make sure the release note expresses the change from a user's
    perspective, not from the perspective of someone contributing to Kubernetes.
    Think about what the user would experience when hitting the problem,
    not the implementation details of the root cause.

    For example:

    User perspective (good) | Code perspective (bad)
    ----------------------- | ----------------------
    *"Fix kubelet crash when Node detaches old volumes after restart."* | *"Call initStuff() before startLoop() to prevent race condition."*

    Ask the PR author for context if it's not clear to you what the release note
    should say.

    Lastly, make sure the release note is located where the [relnotes](https://github.com/kubernetes/release/blob/master/relnotes)
    script will find it:

    * If the cherrypick PR comes from a branch called `automated-cherry-pick-of-*`,
      then the release notes are taken from each parent PR (possibly more than one)
      and the cherrypick PR itself is ignored.

      Make sure the cherrypick PR and parent PRs have the `release-note` label.

    * Otherwise, the release note is taken from the cherrypick PR.

      Make sure the cherrypick PR has the `release-note` label.

    **Notes**

    * Almost all changes that are important enough to cherrypick are important
      enough that we should inform users about them when they upgrade.

      Rare exceptions include test-only changes or follow-ups to a previous
      cherrypick whose release note already explains all the intended changes.

1.  **Approve for cherrypick**

    PRs on release branches follow a different review process than those on the
    `master` branch.
    Patch release managers review every PR on the release branch,
    but the focus is just on ensuring the above criteria are met.
    The code itself was already reviewed, assuming it's copied from `master`.

    * For an *automated cherrypick* (created with `hack/cherry_pick_pull.sh`),
      you can directly apply the `approved` label as long as the parent PR was
      approved and merged into `master`.
      If the parent PR hasn't merged yet, leave a comment explaining that you
      will wait for it before approving the cherrypick.
      We don't want the release branch to get out of sync if the parent PR changes.

      Then comment `/lgtm` to apply the `lgtm` label and notify the author
      you've reviewed the cherrypick request.

    * For a *manual patch or cherrypick* (not a direct copy of a PR already merged
      on `master`), leave a comment explaining that it needs to get
      LGTM+Approval through the usual review process.

      You don't need to do anything special to fall back to this process.
      The bot will suggest reviewers and approvers just like on `master`.

    Finally, apply the `cherrypick-approved` label and remove the `do-not-merge`
    label to tell the bot that this PR is allowed to merge into a release
    branch.

    Note that the PR will not actually merge until it meets the usual criteria
    enforced by the merge bot (`lgtm` + `approved` labels, required presubmits,
    etc.) and makes its way through the submit queue.
    To give cherrypick PRs priority over other PRs in the submit queue,
    make sure the PR is in the `vX.Y` release milestone, and that the milestone
    has a due date.

## Branch health

Keep an eye on approved cherrypick PRs to make sure they aren't getting blocked
on presubmits that are failing across the whole branch.
Also periodically check the [testgrid](https://k8s-testgrid.appspot.com)
dashboard for your release branch to make sure the continuous jobs are healthy.

Escalate to test owners or [sig-testing](https://github.com/kubernetes/community/tree/master/sig-testing)/[test-infra](https://github.com/kubernetes/test-infra)
as needed to diagnose failures.

## Release timing

The general guideline is to leave about 2 to 4 weeks between patch releases on
a given minor release branch.
The lower bound is intended to avoid upgrade churn for cluster administrators,
and to allow patches time to undergo testing on `master` and on the release
branch.
The upper bound is intended to avoid making users wait too long for fixes that
are ready to go.

The actual timing is up to the patch release manager, who should take into
account input from cherrypick PR authors and SIGs.
For example, some bugs may be serious enough, and have a clear enough fix,
to trigger a new patch release immediately.

You should attend the [sig-release](https://github.com/kubernetes/community/tree/master/sig-release)
meetings whenever possible to give updates on activity in your release branch
(bugs, tests, cherrypicks, etc.) and discuss release timing.

When you have a plan for the next patch release, send an announcement
([example](https://groups.google.com/forum/#!topic/kubernetes-dev-announce/HGYsjOFtcdU))
to [kubernetes-dev@googlegroups.com](https://groups.google.com/forum/#!forum/kubernetes-dev)
(and *BCC* [kubernetes-dev-announce@googlegroups.com](https://groups.google.com/forum/#!forum/kubernetes-dev-announce))
several working days in advance.
You can generate a preview of the release notes with the [relnotes](https://github.com/kubernetes/release/blob/master/relnotes)
script ([example usage](https://gist.github.com/enisoc/058bf0feddf6bffd8e25aa72f9dc38d6)).

## Release cut

A few days before you plan to cut a patch release, put a temporary freeze on
cherrypick requests by removing the `cherrypick-approved` label from any PR that
isn't ready to merge.
Leave a comment explaining that a freeze is in effect until after the release.

The freeze serves several purposes:

1.  It ensures a minimum time period during which problems with the accepted
    patches may be discovered by people testing on `master`, or by continuous
    test jobs on the release branch.

1.  It allows the continuous jobs to catch up with `HEAD` on the release branch.
    Note that you cannot cut a patch release from any point other than `HEAD`
    on the release branch; for example, you can't cut at the last green build.

1.  It allows slow test jobs like "serial", which has a period of many hours,
    to run several times at `HEAD` to ensure they pass consistently.

On the day before the planned release, run a mock build with `anago` to make
sure the tooling is ready.
If the mock goes well and the tests are healthy, run the real cut the next day.

After the release cut, reapply the `cherrypick-approved` label to any PRs that
had it before the freeze, and go through the backlog of new cherrypicks.

### Hotfix release

A normal patch release rolls up everything that merged into the release branch
since the last patch release.
Sometimes it's necessary to cut an emergency hotfix release that contains only
one specific change relative to the last past release.
For example, we may need to fix a severe bug quickly without taking on the added
risk of allowing other changes in.

In this case, you would create a new, three-part branch of the form
`release-X.Y.Z`, which [branches from a tag](https://github.com/kubernetes/release/blob/master/docs/branching.md#branching-from-a-tag)
called `vX.Y.Z`.
You would then use the normal cherrypick PR flow, except that you target PRs at
the `release-X.Y.Z` branch instead of `release-X.Y`.
This lets you exclude the rest of the changes that already went into
`release-X.Y` since the `vX.Y.Z` tag was cut.

Make sure you communicate clearly in your release plan announcement that some
changes on the release branch will be excluded, and will have to wait until the
next patch release.

### Security release

The Product Security Team (PST) will contact you if a security release is needed
on your branch.
In contrast to a normal release, you should not make any public announcements
or push tags or release artifacts to public repositories until the PST tells you to.

See the [Security Release Process](../security-release-process.md) doc for more
details.

