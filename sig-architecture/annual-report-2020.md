# 2020 SIG Architecture Annual Report

## Operational

**How are you doing with operational tasks in sig-governance.md?**

**Is your README accurate? have a CONTRIBUTING.md file?**

 - Yes, our README is accurate. We do not have a CONTRIBUTING.md file, and so will add one.

**All subprojects correctly mapped and listed in sigs.yaml?**

 - Yes.

**What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?**

 - Our bi-weekly meeting is normally well attended, with approximately 20 individuals on a typical day, though it varies from a handful of folks to 40 or 50 depending on topics.

 - Subproject meetings are smaller, generally in the half dozen range.

**How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**

 - We have a rotating schedule of subproject updates in the bi-weekly community meeting. Our OWNERS files are up-to-date as well.

**How does the group get updates, reports, or feedback from working groups? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**

 - We have not been getting updates from the working groups on any regular cadence.
 - We will make it a priority to get updates in the next upcoming meetings.
 - Currently sponsored workgroups are:
    - wg-api-expression
    - wg-component-standard
    - wg-k8s-infra
    - wg-naming
    - wg-policy
    - wg-reliability

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

 - [June, 2020](https://youtu.be/ObqQxRRl9RQ?t=2277) at Kubernetes Community Meeting. [Slides](https://docs.google.com/presentation/d/1NytMrpVYKzFo7rLcEEHnFl8zOx05fnjs3xBSZXVE0nI/edit?usp=sharing).
 - [December, 2020](https://youtu.be/rnNqcUeCD8E) at KubeCon NA, published on YouTube. We do not have anything lined up for Kubecon EU. We will prioritize giving our next update in Kubecon NA.

## Membership

**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

- All SIG chairs are active and attend most bi-weekly full SIG meetings. We do not have separate tech leads.

- For the conformance subproject, all owners are active in PR review and meeting attendance.

- For the production readiness subproject, all owners are active in both KEP reviews and meetings.

- For the enhancements subproject, all owners are all active.

- For the API review subproject, there are no standing meetings. The subproject coordinates via a [project board](https://github.com/orgs/kubernetes/projects/169) and [mailing list](https://groups.google.com/g/kubernetes-api-reviewers).

- For the code organization subproject, @dims and @liggitt run regular meetings and review PRs and issues, pulling in other owners as needed.

**How do you measure membership? By mailing list members, OWNERs, or something else?**

- Mailing list members, though based on bounces during calendar updates, there are quite a few out-of-date memberships.

**How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?**

- PRR assigns specific approvers to each KEP that is targeting a release. Therefore, we have a clear picture of how much each
  approver must do each cycle. In 1.21, each approver reviewed 22 KEPs. for the 1.22 cycle, we have a new approver (@ehashman) to
  help reduce the load on each approver. Go Elana!
- Code Organization is coming along. Though we haven't updated dep-approvers in a while. Currently the root OWNERS are a bottleneck as their time is limited for dependency updates. We have a new approval plugin that will help.

    - [kubernetes/test-infra#7690](https://github.com/kubernetes/test-infra/issues/7690) would allow routing dep approval to a distinct group
    - Started dep-reviewers alias in [kubernetes/kubernetes#101670](https://github.com/kubernetes/kubernetes/pull/101670), can use to get additional help on reviews

**Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?**

- SIG Architecture owns very little code. The API review and Production Readiness Review projects do have reviewers and approvers
  for those specialized purposes. Each of those programs has a documented shadow program to add new reviewers and approvers.
- The main meeting itself is losing steam a bit. We could use more across the board attendance. We have to find ways to do this with better timely topics.
- The sub projects are in sustainable shape for the most part. We could use more people in conformance, code-organization, and PRR for sure.
- We want to rework the API reviews as well to make sure SIGs can for the most part do their own API reviews.
- The Enhancements subproject has some tooling code, and is a good place for newer folks to contribute. It has been quite active recently, and as a process-focused subproject is a place where program managers can collaborate with other contributors on optimising processes regarding feature development, deprecations, policy rollouts, and other substantial changes.

**What programs do you participate in for new contributors?**
- Mainly through code-organization, we have a new mentee from LFX for example. 
- We do need to grow more folks who can work across SIGs somehow. We might need to put together a mentoring program to help with this.
- We are reviving the idea of a reading group for KEPs.

**Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?**

- Yes. Our chairs are from three different companies. We have many participants and subproject owners from those same companies and others.

## Current initiatives and project health

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?**

- We are proud of the work of the conformance subproject in burning down the [list of API endpoints that were untested](https://apisnoop.cncf.io/conformance-progress).
- the code-organization has come a long way as well, helping coordinate updates to dependencies across various projects to help make our dependencies cleaner (example grpc update across etcd/containerd etc)
- PRR ([policy](https://github.com/kubernetes/community/blob/master/sig-architecture/production-readiness.md), [KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/1194-prod-readiness)) was made mandatory in 1.21 and the PRR team reviewed 66 KEPs, improving the scalability and supportability, and ensuring monitoring and feature enablement is properly conducted to make features production ready.
- Policies to move features across the project to GA, including:
    - [Conformance Without Beta](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/1333-conformance-without-beta)
    - [Preventing Permabeta](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/1635-prevent-permabeta)
- The Enhancements subproject has a number of initiatives, including:
    - Process changes: the team has been working with the Release Team to assist SIGs in taking greater ownership of their KEPs during the release cycle, reducing reliance on the RT and allowing them to better organize their work and increase communication between SIGs and authors.
    - The Receipts Project (ongoing) is working on automating some of the mechanics of KEP process to streamline and better track KEPs in a release cycle.
    - KEPs have been transitioned to the new format which includes kep.yamls and a new directory structure (special thanks to @wojtek-t and new contributor @shekhar-rajak for all of their hard work on this!).

**Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**

- We are working on updates to enhancements process to make it easier for everyone involved.

**What areas and/or subprojects does the group need the most help with?**

- Setting up a mentoring program to increase the number of folks who can work across SIGs.
- Increasing the number of API reviewers and PRR approvers.

**What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?**

- We mostly care about landing changes for the conformance subproject as it directly affects every end user. That has been going well the last few cycles.
