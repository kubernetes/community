# 2020 WG K8s Infra Annual Report

## You and Your Role

**When did you become a chair and do you enjoy the role?**

- **bartsmykla**: February 2020 and I enjoy the role
- **dims**: along with spiffxp, been there right from the beginning. Lately
  having some conflicts on the meeting time, but definitely enjoying the process
- **spiffxp**: Was a chair (organizer?) at group’s formation. I enjoy the role
  when I have time to dedicate to it.

**What do you find challenging?**

- **bartsmykla**: As our working group’s efforts are related to multiple SIGs,
  and there is multiple places, tools, repositories which are needed to move
  some things forward I sometimes feel overwhelmed and anxiety about not
  understanding some of the tools (Prow for example), what is also hard is I
  don’t feel I have enough access and knowledge to speed up and move things
  fasters in relation to Prow migration.
- **dims**: takes too long :) finding/building coalition is hard. Trying hard to
  avoid doing everything by a small set of folks, but not doing too good on that
  front.
- **spiffxp**: Prioritizing this group’s work, and the work necessary to tend to
  this group’s garden (by that I mean weeding/planting workstreams,
  building/smoothing onramps). Work that usually takes precedence is related to
  company-internal priorities, SIG Testing and kubernetes/kubernetes fires.  I
  often find myself unprepared for meetings unless I have been actively pushing
  a specific item in the interim.  Very rarely am I sufficiently aware of the
  group’s activity as a whole to effectively drive.

**Do you have goals for the group?**

- **bartsmykla**: My goal would be to improve documentation for people to be
  easier to understand what and where is happening and which tools and resources
  are being used for which efforts
- **dims**: breaking things up into small chunks that can be easily farmed out
- **spiffxp**: The thing I care most about is community ownership of
  prow.k8s.io, including on-call. If possible, I would like to see the group’s
  mission through to completion, ensuring that all project infrastructure is
  community-owned and maintained.

**Do you want to continue or find a replacement? If you feel that you aren’t
ready to pass the baton, what would you like to accomplish before you do?**

- **bartsmykla**: I would like to continue
- **dims**: Happy to if folks show up who can take over. Always on the look out.
- **spiffxp**:  I personally want to continue, but sometimes wonder if my
  best-effort availability is doing the group a disservice. If there’s a
  replacement and I’m the impediment, I’m happy to step down.  The ideal
  replacement or a dedicated TL would have: 
  - ability to craft and build consensus on operational policies, lead implementation
  - ability to identify cost hotspots, lead or implement cost-reduction solutions
  - ability to identify security vulnerabilities or operational sharp edges
    (e.g. no backups, easy accidents), lead or implement mitigations
  - familiarity with GCP and Kubernetes
  - ability to document/understand how existing project infra is wired together
    (e.g. could fix https://github.com/kubernetes/test-infra/issues/13063 )

**Is there something we can provide that would better support you?**

- **bartsmykla**: I can’t think about anything right now
- **dims**: what spiffxp said!
- **spiffxp**: TBH I feel like a lot of what we need to make this group as
  active/healthy as I would like needs to come from us. For example I don’t
  think a dedicated PM would help without a dedicated TL. I’m not sure how to
  more effectively motivate our contributing companies to prioritize this work.
  I have pined in the past for dedicated contractors paid by the CNCF for this,
  but I think that could just as easily be fulfilled by contributing companies
  agreeing to staff this.

**Do you have feedback for Steering? Suggestions for what we should work on?**

- **bartsmykla**: I can’t think about anything right now
- **dims**: yep, talking to CNCF proactively and formulating a plan. 
- **spiffxp**:  I think there are three things Steering could help with:
  - Policy guidance from Steering on what is in-scope / out-of-scope for
    Kubernetes’ project-infrastructure budget (e.g. mirroring
    dependency/ecosystem projects like cert-manager [1], ci jobs).  It might
    better drive billing requirements, and make it easier/quicker to decide what
    is appropriate to pursue.  At the moment we’re using our best judgement, and
    I trust it, but I sometimes feel like we’re flying blind or making stuff up.
    As far as existing spend and budgeting, we don’t have
    quotas/forecasts/alerts; we’re mostly hoping everyone is on their best
    behavior until something seems outsized, at which point it’s case-by-case on
    what to do.
  - I think it would be helpful to get spend on platforms other than Google
    above-the-table, and driven through this group. I know how much money Google
    has provided, and I know where it’s being spent (though not to the
    granularity of per-SIG).  I lack the equivalent for other companies “helping
    out” (e.g. AWS, Microsoft, DigitalOcean)
  - This is not a concrete request that can be acted upon now, but I anticipate
    we will want to reduce costs by ensuring that other clouds or large entities
    participate in mirroring Kubernetes artifacts.

## Working Group

**What was the initial mission of the group and if it's changed, how?**

Initial mission was to migrate Kubernetes project infrastructure to the CNCF,
creation of teams and processes to support ongoing maintenance.

There has been a slight growth in scope in that new infrastructure that
previously didn't exist is proposed and managed under this group. Examples
include:
- binary-artifact-promotion (project only had image-promotion internally, now
  externally, now attempting to expand to binary artifacts)
- [running triage-party for SIG Release](https://github.com/kubernetes/k8s.io/issues/906)
  (didn't exist until this year)
- [build infrastructure for windows-based images](https://docs.google.com/document/d/16VBfsFMynA7tObzuZGPpw-sKDKfFc_T5W_E4IeEIaOQ/edit#bookmark=id.3w0g7fo9cp7m)
- [image vulnerability dashboard](https://docs.google.com/document/d/16VBfsFMynA7tObzuZGPpw-sKDKfFc_T5W_E4IeEIaOQ/edit#bookmark=id.s3by3vki8jer)
  (it's not clear to me whether even google had this internally before)
- [sharding out / scaling up gitops-based Google Group management](https://docs.google.com/document/d/16VBfsFMynA7tObzuZGPpw-sKDKfFc_T5W_E4IeEIaOQ/edit#bookmark=id.ou5hk544r70m)

**What’s the current roadmap until completion?**

What has been migrated:
- DNS for kubernetes.io, k8s.io
- Container images hosted on k8s.gcr.io
- node-perf-dash.k8s.io
- perf-dash.k8s.io
- publishing-bot
- slack-infra
- 288 / 1780 prow jobs
- GCB projects used to create kubernetes/kubernetes releases
  (exception .deb/.rpm packages)

What remains (TODO: we need to update our issues to reflect this)
- migrate .deb/.rpm package building/hosting to community
  (this would be owned by SIG Release)
  - stop using google-internal tool "rapture"
  - come up with signing keys community agrees to host/trust
  - migrate apt.kubernetes.io to community
- stop using google-containers GCP project (this would be owned by SIG Release)
  - gs://kubernetes-release, dl.k8s.io
  - [gs://kubernetes-release-dev](https://github.com/kubernetes/k8s.io/issues/846)
- stop using k8s-prow GCP project (this would be owned by SIG Testing)
  - Prow.k8s.io
  - Ensure community-staffed on-call can support
- stop using k8s-prow-build GCP project (this would be owned by SIG Testing)
  - 288/1780 jobs migrated out thus far
  - Ensure community-staffed on-call can support
- [stop using k8s-gubernator GCP project](https://github.com/kubernetes/k8s.io/issues/1308)
  (this would be owned by SIG Testing)
  - migrate/replace gubernator.k8s.io/pr (triage-party?), drop gubernator.k8s.io
  - [migrate kette](https://github.com/kubernetes/k8s.io/issues/787)
  - [migrate k8s-gubernator:builds dataset](https://github.com/kubernetes/k8s.io/issues/1307)
  - [migrate triage.k8s.io](https://github.com/kubernetes/k8s.io/issues/1305)
  - [migrate gs://k8s-metrics](https://github.com/kubernetes/k8s.io/issues/1306)
- stop using kubernetes-jenkins GCP project (this would be owned by SIG Testing)
  - gs://kubernetes-jenkins (all CI artifacts/logs for prow.k8s.io jobs)
  - sundry other GCS buckets (gs://k8s-kops-gce, gs://kubernetes-staging*)
- [stop using k8s-federated-conformance GCP project](https://github.com/kubernetes/k8s.io/issues/1311)
  (this would be owned by SIG Testing)
  - Migrate to CNCF-owned k8s-conform (rename/copy sundry GCS buckets, distribute new service account keys)
-   [stop using k8s-testimages GCP project](https://github.com/kubernetes/k8s.io/issues/1312)
    (this could be owned either by SIG Testing or SIG Release)
  - Migrate images used by CI jobs (kubekins, bazel-krte, gcloud, etc.)
  - Migrate test-infra components (kettle, greenhouse, etc.)
  - (This may push us toward [limited/lifecycle-based retention of images, which
    GCR does not natively have](https://github.com/kubernetes/k8s.io/issues/525)?)
- stop using kubernetes-site GCP project (unsure, maybe SIG ContribEx or SIG Docs depending)
  - ???
- Ensure SIG ownership of all infra and services
  - Must be supportable by non-google community members
  - Ensure critical contributor user journeys are well documented for each service

**Have you produced any artifacts, reports, white papers to date?**

We provide a [publicly viewable billing report](https://datastudio.google.com/u/0/reporting/14UWSuqD5ef9E4LnsCD9uJWTPv8MHOA3e)
accessible to members of kubernetes-wg-k8s-infra@googlegroups.com.
The project was given $3M/yr for 3 years, and our third year started ~August 2020.
Our spend over the past 28 days has been ~$109K, which works out to ~$1.42M/yr.
A very rough breakdown of the $109k:
- $74k - k8s-artifacts-prod* (~ k8s.gcr.io)
- $34k - k8s-infra-prow*, k8s-infra-e2e*, k8s-staging* (~ project CI thus far, follows kubernetes/kubernetes traffic)
- $0.7k - kubernetes-public (~ everything else)

**Is everything in your readme accurate? posting meetings on youtube?**

Our community
[readme](https://github.com/kubernetes/community/tree/master/wg-k8s-infra) is
accurate if sparse. The
[readme](https://github.com/kubernetes/k8s.io/blob/main/README.md) in k8s.io,
which houses most of the actual infrastructure, is terse and slightly out of
date (missing triage party)

[We are having problems with our zoom automation](https://github.com/kubernetes/community/issues/5199),
causing [our youtube playlist](https://www.youtube.com/playlist?list=PL69nYSiGNLP2Ghq7VW8rFbMFoHwvORuDL)
to fall out of date; I noticed while writing this report and have gotten help
backfilling. We're currently missing 2020-10-14.

**Do you have regular check-ins with your sponsoring SIGs?**

No formal reporting in either direction. Meetings/slack/issues see active
participation from @spiffxp (SIG Testing chair), and occasional participation
from @justaugustus (SIG Release) and @nikhita (SIG Contributor Experience). We
also see participation on slack/issues/PRs from @dims (SIG Architecture) who has
a schedule conflict.
