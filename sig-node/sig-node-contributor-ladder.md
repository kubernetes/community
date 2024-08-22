# SIG Node Contributor Ladder

**Note:**
This document is seeded with content based on initial review from existing SIG Node technical and sub-project leads.

## Goal

This document intends to define a standard for supporting community members to grow responsibility in the SIG.  The SIG has a goal to scale individual participation using transparent criteria.  It is aspirational, and should be re-evaluated periodically to ensure we can meet the needs of the project with the resources available to participate.  It is expected that it provides the SIG a log for future evolution of decision making criteria.

If you are a community member that wants to suggest changes, please do!

## New Contributors

SIG Node welcomes new contributors to the community, help is always desired.  Not all contributors are able to provide sustained contribution, but each contribution is always welcome.  This document intends to capture a leadership path for contributors that intend to provide a sustained contribution to the SIG and its associated code base by taking on reviewer and approver responsibilities at various levels.

Please review the list of [sub-projects](./README.md#subprojects) for areas to engage!

## SIG Node reviewers and approvers

SIG Node maintains Kubernetes components that are running large amounts of workloads across clouds, and standalone deployments. Changes in these components have a drastic effect on workload availability. This sets a high bar for SIG Node contributions and puts a lot of emphasis on reliability and security of these changes. Practices and policies of earlier stages of Kubernetes need to be adjusted to the current maturity level of Kubernetes.

This document explains SIG Node guidelines and requirements to become a reviewer or approver. These requirements are written in the spirit of established Kubernetes [membership document](../community-membership.md) and attempt to help contributors, who want to get reviewer and approver status, to formalize the implicit requirements that the SIG aspires to use when making decisions.

Ultimately, membership in the Kubernetes community, as well as reviewer and approver status is a way to establish trust. Starting with permission to run tests, all the way to deciding on the SIG and Kubernetes overall roadmap for key APIs. No amount of documentation can formalize how this trust is established. This document outlines ways to demonstrate expertise of the code base, sound judgement on decision tradeoffs, end user advocacy, care for community, and ability to work as a distributed team.

As a high level guidance, SIG Node reviewers and approvers must have a bias to say “no” in reviews and design discussions, and an initial bias towards coding over reviews to demonstrate a baseline for knowledge of code base.  This high level guidance is dictated by the maturity stage of Kubernetes. It is better to have inefficiencies or known bugs than break existing workloads. It is also important to keep the codebase healthy and maintainable.

We anticipate providing opportunities for emerging contributors to make these contributions where they express their desire for long term maintenance.  This is facilitated during release planning when pairing approvers with contributors for new feature work and/or specific contributions focused on testing and e2e health as we maintain our CI. Helping maintain our CI and test suite is a great way of demonstrating competency!

### Reviewer
Reviewer status is an indication that the person is committed to the SIG activities, demonstrates technical depth, has accumulated enough context, and is overall trustworthy to help inform approvers and aid contributors by applying an lgtm. Anyone is welcome to review a PR with a comment or feedback even if they do not have rights to apply an lgtm.  The requirements listed in the [membership document](../community-membership.md#reviewer) highlight this as well.

The following is the breakdown of requirements listed in the membership document broken down to these categories:

- Committed - proof of sustained contributions
- Kubernetes member for at least 3 months
- Demonstrates technical depth
- Primary reviewer for at least 5 PRs to the codebase
- Reviewed or merged at least 20 substantial PRs to the codebase
- Has enough context
- Knowledgeable about the codebase
- Trustworthy - established trust with the community
- Sponsored by a subproject approver
- With no objections from other approvers

SIG Node establishes the following clarifications and additional requirements.

**Committed**

The 3 month of activities should be established looking at PRs reviews history.
Active participating in SIG Node or SIG Node CI weekly meetings or other transient meetings that arise when tackling specific problems in future (exceptions are allowed for cases when timezone or other personal limitations are not allowing for the meeting participation)

SIG Node recognizes that 3 months may not be enough to establish a deep understanding of the codebase. Reviewer status for a subfolder is a step towards SIG-wide reviewer. See the areas section of this document below.

**Technically sound**

Proof of primary reviewership and significant contributions must be provided. Nominees must provide the list of PRs (at least 5 for primary reviewer and 20 substantial PRs authored or reviewed) as suggested in the membership document. Here are additional comments for this list of PRs:

- Reviewed PRs must be merged.
- Since the purpose is to demonstrate the nominee's technical depth, PRs like analyzer warnings fixes, mechanical “find/replace”-type PRs, minor improvements of kubelet logging and insignificant bug fixes are valued, but not counted towards the reviewer status nomination. Lack of reviews of those PRs may be a red flag for nomination approval.
- Making the most of reviews in a single area is a good indication that a subdirectory reviewer role may be preferred or top-level node reviewer.  Finding the right balance helps balance the nominee's future workload.
- A primary reviewer should drive the review of the PR without significant input / guidance from the approver or other reviewers. 

It is hard to assess codebase knowledge and it always will be a judgement call. SIG Node will rely on the listed PR to ensure the person reviewed PRs from different areas of SIG Node codebase and on the comments made during SIG Node meetings.

Additional ways to establish the knowledge of context are:
- Contributions to k8s documentation
- Blog posts - k8s-hosted and external
- Talks at conferences and meetups
- Contributions to other SIGs

**Trustworthy**

Reviewer nominations are accepted by SIG Node approvers. SIG Node approvers take nominations seriously and are invested in building a healthy community.  Nominees should help approvers understand their future goals in the community so we can help continue to build trust and mutual relationships and nurture new opportunities if and when a contributor wants to become an approver!

### Approver

SIG Node approvers have a lot of responsibilities and it is a very demanding role as the project keeps expanding in capability. It is expected that the SIG Node approver keeps the codebase quality high by giving feedback, thoroughly reviewing code, and giving recommendations to SIG members and reviewers.  SIG Node approvers are essentially gatekeepers to keep the code base at high quality. SIG Node maintains a rigidly high bar for becoming a SIG Node approver by developing trust in a community and demonstrating expertise with a bias towards initial code contributions over reviewing PRs.

SIG Node approver role is not based on a volume of ongoing work, there is no absolute quota. Approver role is based on accumulated trust and knowledge and it is not expected to require 100% full time commitment from any individual member. The expectation is that an active (non emeritus) approver is responsive to a direct ask to review a complex PR or KEP in their area of expertise when solicited.  It is encouraged for a top-level approver to opt out of a review where another approver is more appropriate.

We expect at this stage of SIG Node maturity for approvers to have a strong bias to say “no” to unneeded changes or improvements that don't clearly articulate and demonstrate broad benefits. It also means that the velocity of new features may be affected by this bias. Our continuous work to improve the reliability of the codebase will help to maintain feature velocity going forward.  

While evaluating a nomination for top level approval, nominees may be asked to provide examples of strict scrutiny.  Strict scrutiny refers to instances where a performance regression, vulnerability, or complex unintended interaction could have occurred.  We do not expect existing approvers or nominees to be perfect (no one is!) but as a maintainer community we have had instances of pull requests that we want to learn from and spot to mitigate potential risks given our trust to users and existing project maturity level.  Where specific examples are not present for a nominee (which is fine), we may privately share examples from our past experience for warning signs.

SIG Node differentiates code approvers and roadmap approvers by keeping the bar of enhancements approver status higher than code approver status.  In turn, achieving approver status for SIG Node is not an all or nothing experience. We require approvers to develop trust by gradually getting this status for subareas or by maintenance in adjacent projects (e.g. container runtimes).

In addition to the formal requirements for the [approver role](../community-membership.md#approver), SIG Node makes these recommendations for nominees for the SIG Node top level approver status on how to demonstrate expertise and develop trust.

A top-level approver MUST have intermediate approver rights in a subfolder OR adjacent vendored project or client-server flow (device plugin, container runtime, etc.).  Ideally approver rights in more than one subfolder before top-level kubelet approver is solicited is desired but not required.  This is a means of earning trust to existing approvers. 

There are different ways to earn trust:

**Deep expertise across multiple sub-systems**

- Demonstrated impact across multiple sub-systems.
 - Troubleshooting complex issues that cross-subsystems with deep understanding of tooling and code base.
 - Code contribution informed from low-level analysis or optimization of bottlnecks in kubelet.  Examples would be optimization of pod startup time, materially improving resource allocations, optimizations from pprof analysis, improvements of kubelet to kube-api traffic at scale (lease optimization, etc.).
 - Create and merge major code simplification and/or optimization PRs indicating deep understanding of tradeoffs taken and validation of potential side effects.

**Proficient in features development**
- Drive a few major KEPs at all three stages: 
 - “alpha” - design proposal and discussions
 - “beta” - initial customer feedback collection
 - “GA/deprecation” - stabilizing feature, following PRRs, or managing deprecation. 
 - Demonstrate ability to stage changes and pass PRRs keeping the end user experience and Kubernetes reliability as top priorities.
- Be a reviewer for a few major KEPs and demonstrate meaningful participation in the review process. 
- Drive features in connected projects which have the direct effect on SIG Node (like Runtimes - Containerd or CRI-O, cAdvisor, runc, etc.).
- Give actionable feedback for the KEPs and initial proposals during the SIG node meetings.

**Active community support**
- Be a primary PR reviewer for numerous PRs in multiple areas listed as a requirement for a reviewer.
- Actively triage issues and PRs, provide support to contributors to drive their PRs to completion.

**Be present**
Participate in SIG Node meetings by speaking about KEPs or improvements driven, or find some other way to prove the identity behind GitHub handle.

## Emeritus Approvers

The [emeritus_approvers](https://www.kubernetes.dev/docs/guide/owners/#emeritus) section is used
to list approvers who may no longer have time to spend on the project regularly. Keeping the list of
_active_ approvers up to date helps contributors find the approver for their work easier.
However, the emeritus section as important, as people listed in it are recognized
for their domain knowledge and expertise.

As the process of becoming a SIG Node approver is a multi-year journey, the SIG Node recognizes
that the change of jobs and focus is only natural. So when considering the contributions to become SIG Node approver,
the SIG Node recognizes contributions over years, no matter how long ago.
It also makes the process of returning the emeritus approver back to regular approver easy.

It is important for approvers and reviewers to be actively engaged in the SIG Node community
both to maintain the health of the community, and to maintain up-to-date technical knowledge and state.
We encourage reviewers and approvers who anticipate an extended absence (6+ months) from SIG Node
to move themselves to emeritus.

On return to SIG Node, emeritus members can be fast-tracked into previous roles.
To do so, we expect the following:

- return to the SIG Node community, and demonstrate familiarity with the current state
- committed to sustained future SIG Node presence (for at least the next 3 months)
- no objections from other approvers
- request the previous role, providing the proof of satisfied requirements.
