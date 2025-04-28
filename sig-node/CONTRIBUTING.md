# Contributing on SIG Node

Welcome!

Thank you for your interest in contributing to SIG Node. SIG Node is one of the
biggest SIGs in Kubernetes. Reliability and performance of millions of nodes
running critical applications in production rely on the quality of your
contribution(s). The diversity of workloads, environments, and the scale SIG
Node needs to support makes every code change risky as all the side effects need
to be evaluated.

Please make sure you understand and are up to the challenge. The contributing
instructions are designed to help you.

## For Kubernetes Code Contributions

Read the [Kubernetes Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide#contributor-guide).

If you aspire to grow scope in the SIG, please review the [SIG Node contributor ladder](./sig-node-contributor-ladder.md) for SIG specific guidance.

### For Enhancements

Find out if [your thing is an enhancement a.k.a KEP](https://github.com/kubernetes/enhancements/?tab=readme-ov-file#is-my-thing-an-enhancement).
A good way to do that is to open and issue and get feedback from SIG Node
reviewers and approvers. You can present your idea at a weekly
sig-node meeting to get interactive and more immediate feedback.

If you plan to contribute an enhancement, please prepare yourself for at least 1
year of engagement. A KEP takes at least 4 kubernetes releases to move from
alpha to beta to GA. If there are API / kubelet skew considerations, it may take
even longer. SIG Node expects that contributors commit to taking a KEP to GA stage to avoid
[permanent betas](https://kubernetes.io/blog/2020/08/21/moving-forward-from-beta/#avoiding-permanent-beta).
It is always surprising how much work is needed to productize the feature after
it seems complete - addressing edge cases, comprehensive testing, and documentation
take a lot of effort.

If you are not ready for this commitment, you can consider teaming up with other
contributors in the community or contribute to a KEP driven by somebody else.

SIG Node enhancements are available:

- Committed KEPs [directory](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node)
- All open KEPs [tracking issues](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+)

Here are some best practices how to approach KEP development:
It is based on a [talk](https://kcsna2023.sched.com/event/1Sp9i/implementing-a-big-feature-on-an-example-of-a-sidecar-container-kep) 
*"Implementing a big feature on an example of a Sidecar container KEP"*
([Recording](https://www.youtube.com/watch?v=c3iV8E8EDUA), [Slides](https://static.sched.com/hosted_files/kcsna2023/a0/KCS-NA-2023-ppt.pdf)).

#### Before Starting

- **Prove the need**: Clearly articulate the problem the KEP addresses, identify
  the target audience, and demonstrate the community-wide benefits.  
- **Secure sponsorship and reviews**: Find sponsors, reviewers, and approvers
  early in the process to ensure alignment and avoid delays.  
- **Show commitment**: Demonstrate dedication to the KEP's success by actively
  working on its implementation and ensuring code quality.  
- **Manage expectations**: The KEP process takes time, anticipate at least two
  releases for beta and four for GA.

At this stage the expectation is that the proposal is written in general-enough
terms as a Google Doc for easy commenting and fast collaborative editing.
Sharing the design document with `dev@kubernetes.io` for commenting and with SIG
members `sig-node@googlegroups.com` for commenting or in some cases
editing is a good practice.

It is also very helpful to attend SIG Node weekly meeting to present your
proposal. Most of the time meeting agenda is open to discuss any proposals.
During the meeting you can gather initial feedback, find collaborators, and
secure sponsorship.

#### API Design

- **Define use cases and scope**: Enumerate specific use cases and define the
  problem's boundaries to avoid scope creep.  
- **Consider the bigger picture**: Illustrate how the KEP fits into the existing
  Kubernetes design and how it will handle future requests.  
- **Document decisions**: Record all design decisions, including pros, cons, and
  responsible individuals.  
- **Address potential misuse**: Anticipate potential abuse or misuse scenarios
  and design the API to mitigate them.  
- **Engage reviewers**: Utilize SIG experts for API pre-reviews and PRR
  pre-reviews to gain support and streamline the review process.

Kubernetes API is a main interface many users experience Kubernetes. API
approvers are often the most experienced community members, who can help ensure
the feature fit Kubernetes best practices, will not break compatibility, and
will fit nicely with other Kubernetes capabilities. Even if use cases were
approved by SIG Node approvers, API approvers may reject the proposal and ask to
redesign it.

Some KEPs may go back and forth between use cases and API design for many
iterations. This often happens when use cases are not described completely or a
lot of context is lost in written feedback. If the KEP is going in those
circles, the recommendation is to request a meeting between SIG Approvers and
API approvers driven by KEP author. It may be a dedicated meeting or an invite
of API approvers to SIG Node weekly meeting or SIG Node approvers to API
meeting.

Once API approval was received, SIG Node approvers will review it again as SIG
always has the last word in the feature design.

Note, SIG approvers may request sign offs from other SIGs like Security,
Instrumentation, Storage, Networking, Windows, etc. The process of getting
approval is generally the same.

#### Implementation

- **Structure the implementation**: Divide the implementation into
  pre-factoring, minimal complete product, and post-API changes for better
  organization and review.  
- **Isolate feature gate logic**: Ensure the mainline code remains unaffected
  when the feature gate is disabled.  
- **Adapt and adjust**: Be prepared to modify the KEP's scope or features based
  on implementation challenges.  
- **Collaborate effectively**: Maintain communication through group chats or
  meetings and consider using a shared branch for better coordination.  
- **Improve the codebase**: Leave the code in a better state than you found it
  to facilitate future maintenance and enhance your credibility.

By adhering to these best practices, you can increase the chances of your KEP
being successfully implemented and accepted.

Sometimes SIG may over-commit on KEPs for the release. And if two big KEPs
touching the same code path, SIG may decide to only take one KEP for a release.
Even if this is the case, properly structured KEP implementation will ensure
that some progress was made for that release.

### Helpful Links for SIG Node

**Code**:  

For general code organization, read [contributors/devel/README.md](../contributors/devel/README.md) to learn about things like
`vendor/`, `staging`, etc.

* kubelet
  * <https://github.com/kubernetes/kubernetes/tree/master/cmd/kubelet>
  * <https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet>
* Probe: <https://github.com/kubernetes/kubernetes/tree/master/pkg/probe>
* NodeLifecycle: <https://github.com/kubernetes/kubernetes/tree/master/pkg/controller/nodelifecycle>
* Node API: <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/api/node>
* CRI API: <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/cri-api>
* DRA:
  * <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/dynamic-resource-allocation>
  * <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/kubelet/pkg/apis/dra/>
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e/dra/>
* E2E test:
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e/node/>
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e_node/>
* CI (test-infra)
  * [sig-node jobs](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-node/)
  * [e2e_node jobs](<https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/>
  * [sig-node test-grids](https://github.com/kubernetes/test-infra/blob/master/config/testgrids/kubernetes/sig-node/)
  * [containerd test-grids](https://github.com/kubernetes/test-infra/blob/master/config/testgrids/kubernetes/containerd/)

**Development Resources**:  

* <https://github.com/kubernetes/community/tree/master/contributors/devel#table-of-contents>

There are two types of end-to-end tests in Kubernetes:

* [Cluster end-to-end tests](https://git.k8s.io/community/contributors/devel/sig-testing/e2e-tests.md)
* [Node end-to-end tests](https://git.k8s.io/community/contributors/devel/sig-node/e2e-node-tests.md)

**Shared space / Sub projects**:  

* <https://github.com/kubernetes/community/tree/master/contributors/devel/sig-node>
* <https://github.com/kubernetes/community/tree/master/sig-node#subprojects>

**Triage**:

* <https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md>

**Test Grids**:

* SIG Node overview: <https://testgrid.k8s.io/sig-node>
* Release Blocking: <https://testgrid.k8s.io/sig-node-release-blocking>
* Kubelet: <https://testgrid.k8s.io/sig-node-kubelet>
* Containerd: <https://testgrid.k8s.io/sig-node-containerd>

## Getting Started

Task #1 : Compile kubelet
See tips in the root Makefile:

* <https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#building-kubernetes>

Task #2 : Run a single unit test  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#unit-tests>

Task #3 : Explore update/verify scripts  

hack/update-gofmt.sh + hack/verify-gofmt.sh  

* <https://github.com/kubernetes/kubernetes/blob/master/hack/update-gofmt.sh>
* <https://github.com/kubernetes/kubernetes/blob/master/hack/verify-gofmt.sh>

Task #4 : Explore dependencies  

hack/pin-dependency.sh + hack/update-vendor.sh  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/vendor.md>

Task #5 : Using local-up-cluster script  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md#starting-the-cluster>

Running a local cluster  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md>

Note: Task 5 requires Linux OS

## Subgroups

### CI testing

CI testing subgroup of SIG Node provides a great way to get started with the SIG.
CI testing concentrate on  reliability and ongoing issues, see the [charter](sig-node-ci-testing-group-charter.md) for details.

There are many ways to participate:

1. Bring a topic for the discussion at a weekly (see "How to..." section below).
  Topic must be aligned with the group charter.
2. Participate in test issues, PRs, and bugs triage during the meeting.
  Expertise in various areas is needed when triaging.
  So YOUR expertise may be very valuable for some of the topics discussed.
3. Host a section of a meeting. Once you attended a meeting a few times,
  you may be ready to host the section of a meeting. It will include
  presenting your screen and applying needed labels to the issues.

## How to ...

### add agenda item to the meeting

This applies to both, SIG Node weekly meeting and CI testing subgroup weekly meeting.

1. Find the meeting invite and notes document.
2. Join the group sig-node@kubernetes.io for edit permissions.
3. Add you topic to the end of the list, prefixed with your alias, to the agenda of the next meeting.
4. Show up at the meeting to present your agenda item or find a delegate to present it.

There is generally no restrictions on topics, and topics are added on first come, first served basis.
Chair who host the meeting will try to allocate time and fit as many topics as possible.
However declared order is not guaranteed. Also topics may be moved to the next meeting
if other things like KEP planning must be prioritized.

### upload meeting recording to YouTube (for chairs)

1. Download the video file from zoom account `sig-node-leads@kubernetes.io`
2. Log in to YouTube under the `kubernetes-sig-node` account (use switch account)
3. Upload video to YouTube
   - Name should follow the template: `<name of the meeting> yyyymmdd` like `Kubernetes SIG Node 20240618`.
   - Add a link to the meeting notes and agenda document to the meeting description
   - Add video to k8s-owned YouTube playlist:
     - Once video is uploaded, open the video page, click +Save button under the video, select the checkbox `SIG Node Meetings`.
4. Add link to the video to the meeting notes document
5. In zoom account change the title of the recording - add `(uploaded)` suffix
