# Contributor Experience Special Interest Group Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

The [Contributor Experience Special Interest Group] (SIG) is responsible for improving the experience of those who upstream contribute to the Kubernetes project. We do this by creating, and maintaining programs and processes that promote community health and reduce project friction, while retiring those programs and processes that don't. Being conscientious of our contributor base is critical to scaling the project, growing the ecosystem, and helping the project succeed.

We do this by listening - whether it’s through our roadshows to SIG meetings, surveys, data, or [GitHub issues], we take in the feedback and turn it into our [project list]. We build a welcoming and inclusive community of contributors by giving them places to be heard and productive.

### In scope

#### Code, Binaries and Services

- Establish policies, standards and procedures for the use, [moderation], and management of all public platforms officially used by the project, including but not limited to:
  - [discuss.kubernetes.io]
  - [GitHub Management]
  - [Mailing lists] / Google groups for the project as a whole (eg: kubernetes-dev@googlegroups.com) and for individual sigs and wgs where the Chairs have provided us ownership
  - [Slack]
  - [/kubernetescommunity] YouTube channel
  - [Zoom]
- Establish and staff teams responsible for the administration and moderation of these platforms
  - Teams must be staffed by trusted contributors spanning time zones, see [moderation] for more detail
  - They are authorized to take immediate action when dealing with code of conduct issues, see [moderators] for the full list
  - They are expected to escalate to the project's code of conduct committee when issues rise above the level of simple moderation
- Work with other SIGs and interested parties in the project to execute GitHub tasks where required, see [GitHub Management] for more detail
- Own and execute events that are targeted to the Kubernetes contributor community, including:
  - The weekly [Kubernetes Community meeting]
  - [Contributor Summit(s)]
  - [Steering Committee elections] (though we do not own policy creation, see 'out of scope' below)
  - Retrospective moderation for other SIGs upon request
  - Other events, like other SIG face to face events, upon request and consideration
- Strategize, build, and execute on scalable [mentoring programs] for all contributor levels. These may include:
  - [Google Summer of Code]
  - [Outreachy]
  - [Meet Our Contributors]
  - [Group Mentoring - WIP]
  - [The 1:1 Hour - WIP]
  - Speed Mentoring sessions at selected KubeCon/CloundNativeCon's
- Help onboard new and current contributors into the culture, workflow, and CI of our contributor experience through the [contributor guide], other related documentation, [contributor summits] and programs tailored to onboarding.
- Perform issue triage on and maintain the [kubernetes/community] repository.  
- Help SIGs with being as transparent and open as possible through creating best practices, guidelines, and general administration of YouTube, Zoom, and our mailing lists where applicable
- Assist SIGs/WG Chairs and Technical Leads with organizational management operations as laid out in the [sig-governance] doc
- Distribute contributor related news on various Kubernetes channels, including Cloud Native Compute Foundation ([CNCF]) for posting blogs, social media, and other platforms as needed.
- Establish and share metrics to measure project health, community health, and general trends, including:
  - ongoing work with the CNCF to improve [DevStats]
  - the contributor experience survey(s)
  - engagement on project platforms that we manage
- Research other OSS projects and consult with their leaders about contributor experience topics to make sure we are always delivering value to our contributors

#### Cross-cutting and Externally Facing Processes

We cross-cut all SIGs and WGs to deliver the following processes:

- Deploying Changes:
  When implementing policy changes we strive to balance responding quickly to the needs of the community and ensuring a disruption-free experience for project contributors. As such, the amount of notice we provide and the amount of consensus we seek is driven by our estimation of risk. We don't measure risk objectively at this time, but estimate it based on these parameters:
  - Low-risk changes impact a small number (<4) of SIGs, WGs, or repos, do not break existing contributor workflows, and are easy to roll back. When implementing low-risk changes we:
    - Socialize on kubernetes-sig-contribex@googlegroups.com and our weekly update calls
    - We will go to each lead, their mailing lists, slack channel, and/or their update meetings and ask for feedback and a [lazy consensus] process. We will follow up with a post to [kubernetes-dev@]googlegroups.com mailing list
  - High-risk changes impact a large number (>4) of SIGs, WGs, or repos, break existing contributor workflows, and are not easy to roll back. When implementing high-risk changes we:
    - Socialize on kubernetes-sig-contribex@googlegroups.com and our weekly update calls
    - Seek [lazy consensus] with a time box of at least 72 business hours with a GitHub issue link (or proposal if not applicable) to the following mailing lists:
        - [kubernetes-sig-contribex@]googlegroups.com
        - sig-leads@googlegroups.com
        - [kubernetes-dev@]googlegroups.com with the GitHub issue link including the subject [NOTICE]: $announcement
        - We will also announce it at the weekly Kubernetes Community Meeting on Thursdays
- Depending on how wide of an ecosystem change this is, we may also slack, blog, tweet, and use other channels to get the word out.
- Our standard time box is 72 business hours; however, there may be situations where we need to act quickly but the time period will always be clear and upfront.
- Encourage automation to improve productivity for contributors where it makes sense and consult with SIG Testing if automation is covering GitHub workflows.

If we need funding for any reason, we approach [CNCF].
CNCF in many of the noted cases above, contributes funding to our platforms, processes, and/or programs. They do not play a day-to-day operations role and have bestowed that to our community to run as we see fit. Since they do fund some of our initiatives, this means that they hold owner account privileges on certain platforms like Zoom and Slack. In these cases, such as Slack, there are at least two Contributor Experience [communication] subproject OWNERs listed as admins.

### Out of scope

- Code for the testing and CI infrastructure - that’s SIG Testing
- [kubernetes/community]  ownership of folders for KEPs and Design Proposals. Members are to follow those folders owners files and SIG leadership for the specific issue/PR in question.
- User community management. We hold office hours because contributors are a large portion of the volunteers that run that program.
- The contributor experience for repos not included in the Kubernetes associated repositories list found in the [GitHub Management] subproject README.
- Steering committee election policy updates and maintenance.
- We do not create SIGs/WGs but can assist in the various community management needs of their micro communities that would kick off their formation and keep them going.
- We are not the [code of conduct committee] and therefore do not control incident management reporting or decisions; however, our moderation guidelines allow us to act swiftly if there is a clear violation of terms of either our code of conduct or one of our supported platforms terms of service. If there is an action that the committee needs to take that involves one of these platforms (example: the removal of someone from GitHub), we will carry that out if none of the committee members have access.
- Communication platforms that are out of our scope for maintenance and support but we may still have some influence:
    - [r/kubernetes]
    - [@kubernetesio] twitter account
    - [kubernetes blog]

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].


### Additional responsibilities of Chairs

Chairs SHOULD plan at least one face to face Contributor Experience meeting per year

### Additional responsibilities of Tech Leads

Ensuring that technical changes from subprojects follow the process change communication guidelines listed above.

### Deviations from sig-governance
Six months after this charter is first ratified, it MUST be reviewed and re-approved by the SIG in order to evaluate the assumptions made in its initial drafting.

### Subproject Creation
Chairs and Technical Leads

[sig-governance]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://git.k8s.io/community/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus
[Contributor Experience Special Interest Group]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[kubernetes-dev@]: https://groups.google.com/forum/#!forum/kubernetes-dev
[@kubernetesio]: https://www.twitter.com/kubernetesio
[r/kubernetes]: https://kubernetes.reddit.com
[Google Summer of Code]: https://git.k8s.io/community/mentoring/google-summer-of-code.md
[Outreachy]: https://git.k8s.io/community/mentoring/outreachy.md
[Meet Our Contributors]:  https://git.k8s.io/community/mentoring/meet-our-contributors.md
[Group Mentoring - WIP]:  https://git.k8s.io/community/mentoring/group-mentoring.md
[The 1:1 Hour - WIP]: https://git.k8s.io/community/mentoring/the1-on-1hour.md
[kubernetes/community]: https://git.k8s.io/community/
[Contributor Summit(s)]: https://git.k8s.io/community/events/2018/12-contributor-summit
[contributor summits]: https://git.k8s.io/community/events/2018/12-contributor-summit
[DevStats]: https://k8s.cncf.devstats.io
[kubernetes-sig-contribex@]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[kubernetes blog]: https://www.kubernetes.io/blog
[GitHub Management]: https://git.k8s.io/community/github-management
[communication]: https://git.k8s.io/community/communication
[CNCF]: https://cncf.io
[GitHub issues]: https://github.com/kubernetes/community/issues
[project list]: https://github.com/orgs/kubernetes/projects/1
[Kubernetes Community meeting]: https://git.k8s.io/community/communication#weekly-meeting
[mentoring programs]: https://git.k8s.io/community/mentoring
[Steering Committee elections]: https://git.k8s.io/community/events/elections
[Slack]: https://git.k8s.io/community/communication/slack-guidelines.md
[Zoom]: https://git.k8s.io/community/communication/zoom-guidelines.md
[/kubernetescommunity]: https://www.youtube.com/kubernetescommunity
[discuss.kubernetes.io]: https://discuss.kubernetes.io
[contributor guide]: https://git.k8s.io/community/contributor
[moderation]: https://git.k8s.io/community/communication/moderation.md
[code of conduct committee]: https://git.k8s.io/community/committee-code-of-conduct
[Mailing lists]: https://git.k8s.io/community/communication/moderation.md#specific-guidelines
[moderators]: https://git.k8s.io/community/communication/moderators.md
