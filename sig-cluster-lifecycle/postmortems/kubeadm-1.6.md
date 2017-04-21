Kubernetes Postmortem: kubeadm 1.6.0 release
============================================

**Incident Date:** 2017-03-28

**Owners:** Jacob Beacham (@pipejakob)

**Collaborators:** Joe Beda (@jbeda), Mike Danese (@mikedanese), Robert Bailey
(@roberthbailey)

**Status:** \[draft | pending feedback | **final**\]

**Summary:** kubeadm 1.6.0 consistently hangs while trying to initialize new
clusters. A fix required creating the 1.6.1 patch release six days after 1.6.0.

**Impact:** Initialization of a new 1.6.0 master using kubeadm.

**Root Causes:** kubelet’s behavior was changed to report NotReady instead of
Ready when CNI was unconfigured
([\#43474](https://github.com/kubernetes/kubernetes/pull/43474)), which caused
kubeadm to hang indefinitely on initialization while waiting for the master node
to become Ready (and then schedule a dummy deployment) in order to validate the
control plane’s health, which was intended to happen before a CNI provider gets
added.

**Resolution:** kubeadm initialization now only waits for the master node to
register with the API server, but does not require it to be Ready, and does not
attempt a dummy deployment to validate the control plane’s health
([\#43835](https://github.com/kubernetes/kubernetes/pull/43835)). This behavior
is being revisited for the 1.7 release.

**Detection:** A customer filed an issue against kubeadm after trying to
initialize a new cluster with the 1.6.0 release
([\#212](https://github.com/kubernetes/kubeadm/issues/212)).

**Lessons Learned:**

**What went well**

- The bug was discovered quickly after the release.

- Once the bug was discovered, a solution was ready within a day, and the patch
  release was available five days after that (on a Monday, so the weekend
  accounted for some of the gap).

**What went wrong**

- The 1.6.0 release of kubeadm never passed end-to-end testing.

  - End-to-end tests only existed against the master branch instead of the
    release-1.6 branch.

  - Conformance testing of kubeadm requires a functioning CNI provider, but due
    to changes in 1.6.0 clusters and CNI itself, previous kubeadm-endorsed CNI
    providers required updating to reflect the new master taint, RBAC being
    enabled, the master’s insecure port being disabled, and to tolerate deletion
    of unknown pods. Without a functional 1.6 CNI provider until very late in
    the development cycle, Conformance tests were disabled for kubeadm’s
    end-to-end jobs in favor of only testing initialization and node joining.

  - The kubeadm end-to-end tests were also constantly breaking throughout the
    development cycle due to upstream kubeadm and test-infra changes. There was
    no automated monitoring to notify the SIG of failures, nor was there a
    process defined for fixing them, which led to a single person manually
    watching them and addressing failures as they occurred. As a result, manual
    testing was being used near the release milestone, and was successfully
    passing through the 1.6.0-beta.4 release. Without coordination with the
    Release Czar, 1.6.0-rc.1 and 1.6.0 were released without manual end-to-end
    testing of kubeadm and contained the regression.

- There was a lot of rushing to get the release ready before KubeCon EU, causing
  a shortened timeframe for RC and release, lowered bandwidth for SIG members,
  and the last SIG meeting before the release to be cancelled, which decreased
  communication.

- There was no explicit release-readiness sign-off by the SIG. The SIG had
  checklists for bringing kubeadm to Beta (the goal for 1.6.0), and they
  included end-to-end tests which were knowingly in a bad state, but no one
  escalated to delay the release or to remove kubeadm’s Beta status.

- After the 1.6.0 bug was discovered, there was no public announcement to let
  users know about the flaw or the timeline to expect a fix.

- There were two GitHub issues
  ([kubeadm\#212](https://github.com/kubernetes/kubeadm/issues/212) and
  [kubernetes\#43815](https://github.com/kubernetes/kubernetes/issues/43815))
  both tracking the bug. Both were flooded by users duplicating the bug report
  or their workarounds, with splintered developer discussions for the short-term
  and long-term fixes, which made the issue noisy for anyone who just wanted
  updates on the status of the official fix. Additional communication occurred
  on Slack channels, so there was no single authoritative source to follow for
  updates.

- Older versions of the kubeadm Debian packages were removed when 1.6.0 was
  released, so users could not fall back on older versions of kubeadm unless
  they had cached versions. This was intentional for this release (since prior
  versions were Alpha and insecure), and shouldn’t happen in future releases,
  but it left some users out of luck who were knowingly depending on kubeadm 1.5
  or wanted to fall back after 1.6.0 failed for them.

**Where we got lucky**

- This bug only manifested during cluster initialization, and occurred
  consistently. This meant that it was detected very quickly, was trivial to
  reproduce, and had minimal impact on customers since they could not have been
  relying on the cluster yet. If the bug had been more subtle, it could have
  been triggered at random points during the lifecycle of a cluster, been more
  difficult to reproduce and fix, and caused harm to clusters that were already
  in use by customers.

- Even without full testing, there were no other kubeadm regressions between
  1.6.0-rc.1 and 1.6.0.

**Action Items:**

| **Item**                                                                                                                                                                      | **Type** | **Owner** | **Issue**                                                                |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|--------------------------------------------------------------------------|
| Add end-to-end kubeadm postsubmit tests for release-1.6 branch                                                                                                                | detect   | pipejakob | DONE                                                                     |
| Add end-to-end kubeadm presubmit tests (non-blocking)                                                                                                                         | prevent  |           | [kubeadm\#250](https://github.com/kubernetes/kubeadm/issues/250)         |
| Add end-to-end kubeadm variants that use non-third-party CNI providers, like “bridge”                                                                                         | prevent  |           | [kubeadm\#218](https://github.com/kubernetes/kubeadm/issues/218)         |
| Notify SIG on kubeadm postsubmit end-to-end test failures                                                                                                                     | detect   |           | [test-infra\#2555](https://github.com/kubernetes/test-infra/issues/2555) |
| Define process of who should triage and/or fix kubeadm end-to-end test failures, and how                                                                                      | prevent  |           | [kubeadm\#251](https://github.com/kubernetes/kubeadm/issues/251)         |
| Do not remove old versions from distribution repositories during release                                                                                                      | mitigate |           | [kubeadm\#252](https://github.com/kubernetes/kubeadm/issues/252)         |
| Define kubeadm release process that blocks future releases on its completion (e.g. setup end-to-end tests for new release branch, when and how to make the go/no-go decision) | prevent  |           | [kubeadm\#252](https://github.com/kubernetes/kubeadm/issues/252)         |
| Document incident response process for critically flawed Kubernetes releases, including how to notify the community and track progress to conclusion                          | mitigate |           | [community\#564](https://github.com/kubernetes/community/issues/564)     |

**Timeline**

All times are in 24-hour PST8PDT.

**2017/02/24**

> 06:00 fejta changes e2e-runner.sh
> ([test-infra\#1657](https://github.com/kubernetes/test-infra/pull/1657)),
> inadvertently regresses kubeadm e2e test

**2017/03/08**

> 13:44 pipejakob fixes regression
> ([test-infra\#2179](https://github.com/kubernetes/test-infra/pull/2179)), but
> the e2e test is still failing because of recent kubeadm CLI changes

**2017/03/08**

> 13:22 spxtr refactors prow config
> ([test-infra\#2192](https://github.com/kubernetes/test-infra/pull/2192)),
> which later breaks kubeadm e2e job configuration when it gets pushed (this
> timestamp is for the merge, but actual activation of config is unknown since
> it is done manually)

**2017/03/09**

> 21:43 pipejakob merges commit to accommodate recent kubeadm CLI changes to
> attempt to fix e2e jobs
> ([kubernetes-anywhere\#352](https://github.com/kubernetes/kubernetes-anywhere/pull/352))

**2017/03/13**

> 11:27 pipejakob temporarily disables kubeadm e2e Conformance testing
> ([test-infra\#2184](https://github.com/kubernetes/test-infra/pull/2184)) to
> get a better signal; test runs are back to green but only exercise
> initializing the cluster and verifying that nodes join correctly
>
> 12:01 while still trying to fix CNI providers on kubeadm e2e test, pipejakob
> finds that even after accounting for expected changes (master taint renaming,
> RBAC being enabled, unauthenticated access being turned off), CNI providers
> still aren’t working
> ([kubeadm\#190](https://github.com/kubernetes/kubeadm/issues/190#issuecomment-286209644))

**2017/03/14**

> 13:11 pipejakob fixes kubeadm e2e job configuration (which was pushed at some
> point after spxtr’s prow configuration refactoring)
> ([test-infra\#2246](https://github.com/kubernetes/test-infra/pull/2246))

**2017/03/16**

> 11:23 bboreham fixes the weave-net CNI provider
> ([weave\#2850](https://github.com/weaveworks/weave/pull/2850)) to account for
> “CNI unknown pod deletion” change
>
> 14:52 krzyzacy migrates kubeadm e2e job to be scenario/json based
> ([test-infra\#2141](https://github.com/kubernetes/test-infra/pull/2141)),
> which breaks the job.
>
> Over the next few days, krzyzacy tries to fix the above regression, but the
> job is ultimately left failing because Conformance testing has been
> erroneously re-enabled, which is known to be broken due to CNI issues
> ([test-infra\#2280](https://github.com/kubernetes/test-infra/pull/2280),
> [test-infra\#2284](https://github.com/kubernetes/test-infra/pull/2284),
> [test-infra\#2285](https://github.com/kubernetes/test-infra/pull/2285),
> [test-infra\#2286](https://github.com/kubernetes/test-infra/pull/2286),
> [test-infra\#2288](https://github.com/kubernetes/test-infra/pull/2288))

**2017/03/17**

> 10:41 enisoc releases
> [1.6.0-beta.4](https://github.com/kubernetes/kubernetes/commit/b202120be3a97e5f8a5e20da51d0b6f5a1eebd31).
> Since e2e tests are broken, pipejakob manually tests cluster initialization
> locally (kubeadm still works), as well as the updated weave-net manifest

**2017/03/22**

> 12:35 dcbw merges change to make kubelet report NotReady when CNI is
> unconfigured
> ([kubernetes\#43474](https://github.com/kubernetes/kubernetes/pull/43474)),
> but e2e tests are already failing so no one notices the kubeadm regression

**2017/03/24**

> 12:06 enisoc releases
> [1.6.0-rc.1](https://github.com/kubernetes/kubernetes/commit/8ea07d1fd277de8ab5ea7f281766760bcb7d0fe5).
> This was the first release to regress kubeadm, but it goes untested.

**2017/03/28**

> 09:23 enisoc releases Kubernetes
> [1.6.0](https://github.com/kubernetes/kubernetes/releases/tag/v1.6.0).
>
> 09:27 pipejakob updates kubeadm e2e job to use weave-net plugin so that
> Conformance testing can be re-enabled
> ([test-infra\#2347](https://github.com/kubernetes/test-infra/pull/2347)), but
> due to the subtle [gcloud ssh
> bug](https://github.com/kubernetes/kubeadm/issues/219), the job is still
> broken after the update, so it masks the new regression in kubeadm init
>
> 22:40 jimmycuadra reports kubeadm 1.6.0 being broken ([kubeadm issue
> 212](https://github.com/kubernetes/kubeadm/issues/212))

**2017/03/29**

> 13:04 kensimon opens PR to fix kubeadm master: “Tolerate node network not
> being ready“
> ([kubernetes\#43824](https://github.com/kubernetes/kubernetes/pull/43824))
>
> 18:29 mikedanese opens second PR to fix kubeadm master in different way:
> “don't wait for first kubelet to be ready and drop dummy deploy.”
> ([kubernetes\#43835](https://github.com/kubernetes/kubernetes/pull/43835))
> pipejakob helps manually test it for QA purposes.
>
> 18:51 mikedanese opens PR for cherry-pick of above fix to release-1.6 branch
> ([kubernetes\#43837](https://github.com/kubernetes/kubernetes/pull/43837))

**2017/03/30**

> 16:57 mikedanese’s kubeadm fix
> ([kubernetes\#43835](https://github.com/kubernetes/kubernetes/pull/43835)) is
> merged to master (kensimon’s gets discarded)
>
> 21:57 mikedanese adds new build of .deb to kubernetes-xenial-unstable channel
> for users to test ([kubernetes issue
> 43815](https://github.com/kubernetes/kubernetes/issues/43815#issuecomment-290616036))

**2017/03/31**

> 00:26 mikedanese’s cherry-pick merged to release-1.6 branch
> ([kubernetes\#43837](https://github.com/kubernetes/kubernetes/pull/43837))
>
> 11:34 pipejakob merges CI job for release-1.6 branch
> ([test-infra\#2352](https://github.com/kubernetes/test-infra/pull/2352))
>
> 16:30 pipejakob merges quick fix
> ([test-infra\#2380](https://github.com/kubernetes/test-infra/pull/2380)) for
> the “gcloud ssh issue,” which fixes Conformance testing

**2017/04/03**

> 13:32 enisoc releases Kubernetes
> [1.6.1](https://github.com/kubernetes/kubernetes/releases/tag/v1.6.1)
