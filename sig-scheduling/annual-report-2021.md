# 2020 SIG-Scheduling Annual Report

## Operational

**How are you doing with operational tasks in [sig-governance.md]?**
    
- **Is your README accurate? have a CONTRIBUTING.md file?**  
    Yes. Developer-oriented guides are under [contributors/devel/sig-scheduling].

- **All subprojects correctly mapped and listed in [sigs.yaml]?**  
    Yes.

- **What's your meeting culture? Large/small, active/quiet, learnings? 
Meeting notes up to date? Are you keeping recordings up to date/trends 
in community members watching recordings?**  
    Our meetings tends to be small and quiet. The agenda usually consists of items suggested
    by users, and debatable items that need a consensus during issues/PR reviews.
    The meeting notes and recordings are up to date.

**How does the group get updates, reports, or feedback from subprojects? Are there any 
springing up or being retired? Are OWNERS files up to date in these areas?**

Owners from (active) subprojects introduce the latest development, and sometime demonstrate
cool features.
OWNER files in k/k are not that up to date. We may need a cleanup.

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

Aug 20, 2020. [Slides] & [Recording].

[Slides]: https://docs.google.com/presentation/d/1H27SDMqkzq8zCRveWWtK5g9hCAomKbrzTTVZ5r4h6Xo/edit
[Recording]: https://www.youtube.com/watch?v=oDL3Kp5-9eM&feature=youtu.be

## Membership

**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

Yes, except for inactive subprojects.

**How do you measure membership? By mailing list members, OWNERs, or something else?**

We don't have an official way of measuring membership, there is some churn in the number of active 
members and so membership is hard to keep track of.

**How does the group measure reviewer and approver bandwidth? Do you need help in any area now?
 What are you doing about it?**

PRs are usually directed to the reviewer most familiar with the code base the PR is modifying.
We don't quite measure bandwidth, but one way of doing that is by looking at pending PRs broken
by assignment.

**Is there a healthy onboarding and growth path for contributors in your SIG? What are some 
activities that the group does to encourage this? What programs are you participating in to 
grow contributors throughout the contributor ladder?**

There is no official onboarding process. One thing we try to do frequently is breaking up 
larger features into smaller enough tasks for new members to contribute.

**What programs do you participate in for new contributors?**

Currently we don't.

**Does the group have contributors from multiple companies/affiliations? Can end users/companies 
contribute in some way that they currently are not?**

Yes.

## Current initiatives and project health

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of?
 Currently underway? What are some of the longer tail projects that your group is working on?**

- Initiatives:
    - Focusing on turning the scheduler into a pluggable framework to allow developing new custom
     features outside the main repo.
     [k-sigs/scheduler-plugins] initiated this to build a scheduler plugin ecosystem. It has a wide
     participation from different companies, by either direct contributions from Alibaba, Tencent,
     Apple, etc., as well as adoption by companies like [OpenAI].
    - Improving scheduler performance. We're working on some items like optimizing internal queues,
     tailored preemption logic, as well as exposing meaningful metrics to help define SLA/SLOs.
- Longer tail projects:
    - Continuing to refactor the core code around the scheduling framework.
    - Graduating the scheduler's ComponentConfig to GA.

**Year to date KEP work review: What's now stable? Beta? Alpha? Road to alpha?**

- Alpha
    - [Prefer Nominated Node]
    - [Node Resource Strategy]
    - [Pod Affinity Namespace Selector]
    - [Volume Capacity Priority] (co-owned by sig-storage)
- Beta
    - [Default topology spread] (will graduate with CC)
    - [Multi Scheduling Profiles]
    - [Non-preempting priority class]
    - [Component Config]

**What areas and/or subprojects does the group need the most help with?**

- **Docs improvement**:
    - developer oriented docs to understand more details of scheduler internals, so that they can come
    up with k8s-scheduler-native extensions to fit their business needs
    - user or cluster-admin oriented docs to make the most of scheduler, like best practices and tips
    that are not documented well
- **Standardize issue triage process**:

**What's the average open days of a PR and Issue in your group? What metrics does your group care about and/or measure?**

We haven't started leveraging devstat data or Github board to get a high-level picture of PR/Issue.

[sig-governance.md]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[contributors/devel/sig-scheduling]: https://github.com/kubernetes/community/tree/master/contributors/devel/sig-scheduling
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sig-list.md
[k-sigs/scheduler-plugins]: https://github.com/kubernetes-sigs/scheduler-plugins

[Prefer Nominated Node]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1923-prefer-nominated-node
[Node Resource Strategy]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2458-node-resource-score-strategy
[Pod Affinity Namespace Selector]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/2249-pod-affinity-namespace-selector
[Volume Capacity Priority]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1845-prioritization-on-volume-capacity
[Default topology spread]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1258-default-pod-topology-spread
[Multi Scheduling Profiles]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/1451-multi-scheduling-profiles
[Non-preempting priority class]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/902-non-preempting-priorityclass
[Component Config]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/785-scheduler-component-config-api

[OpenAI]: https://openai.com/blog/scaling-kubernetes-to-7500-nodes/#gangscheduling
