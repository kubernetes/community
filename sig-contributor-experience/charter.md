# Contributor Experience Special Interest Group Charter

## Scope

Being conscientious of our contributor base is critical to scaling the project, growing the ecosystem, and helping the project succeed. The [Contributor Experience Special Interest Group] (SIG) is responsible for improving the experience of those who contribute upstream to the Kubernetes project. We do this by creating, and maintaining policies and processes that promote community health and reduce project friction. 

### In scope

#### Code, Binaries and Services

The inscope bullets listed below have a matching bullet for out of scope

- Help onboard new and current contributors into the culture, workflow, and CI of our contributor experience through the [contributor guide], other related documentation, [contributor summits] and programs tailored to onboarding.
- Strategize, build, and execute on scalable [mentoring programs] for all contributor levels. These may include:
  - [Google Summer of Code]
  - [Outreachy]
  - [Meet Our Contributors]
  - [Group Mentoring]
  - [The 1:1 Hour]
  - Speed Mentoring sessions at selected KubeCon/CloundNativeCon's
- Establish policies, standards and procedures for the following:
  - [discuss.kubernetes.io]
  - [GitHub Project Organizational Management]
  - [Mailing lists] / Google groups for the project as a whole (eg: kubernetes-dev@googlegroups.com) and for individual sigs and wgs where the Chairs have provided us ownership
  - [Slack]
  - [/kubernetescommunity] YouTube channel
  - [Zoom]
  - Perform issue triage on and maintain the [kubernetes/community] repository.  
  - Maintenance of the Kubernetes Gsuite
    - Managing @kubernetes.io aliases for project usage
    - Managing the Kubernetes Community Calendar
  - The weekly [Kubernetes Community meeting]
  - Content and Format of the [Contributor Summit(s)]
  - Execution of the [Steering Committee elections]
  - Retrospective moderation for other SIGs upon request
  - Assist SIGs/WG Chairs and Technical Leads with organizational management operations as laid out in the [sig-governance] doc
  - Distribute contributor related news on various Kubernetes channels, including Cloud Native Computing Foundation ([CNCF]) for posting blogs, social media, and other platforms as needed.
  - Establish and share metrics to measure project health, community health, and general trends, including:
    - Ongoing work with the CNCF to improve [DevStats]
    - The contributor experience survey(s)
- [GitHub Project Organizational Management]
  - Work with other SIGs and interested parties in the project to execute GitHub tasks where required.
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

- External communications platforms that we do not control but might have some influence over
    - [r/kubernetes]
    - [@kubernetesio] twitter account
    - [Kubernetes Blog]

### Out of scope

- Code for the testing and CI infrastructure - that’s SIG Testing
- [kubernetes/community] Members are to follow those folders owners files and SIG leadership for the specific issue/PR in question.
- User community relations and management.
  - Slack or YouTube moderation (This should be a new SIG)
- The contributor experience for repos not included in the Kubernetes associated repositories list found in the [GitHub Project Organizational Management] subproject README.
- Steering committee election policy updates and maintenance. (Steering Comittee handles this)
- We do not create SIGs/WGs but can assist in the various community management needs of their micro communities that would kick off their formation and keep them going.
  - We do not close down Working Groups, that is the responsibility of the sponsoring SIGs of that working group
- We are not the [code of conduct committee] and therefore do not control incident management reporting or decisions; however, our moderation guidelines allow us to act swiftly if there is a clear violation of terms of either our code of conduct or one of our supported platforms terms of service. If there is an action that the committee needs to take that involves one of these platforms (example: the removal of someone from GitHub), we will carry that out if none of the committee members have access.
- Ongoing maintenance of established SIG Resources. We are available to assist SIGs with these properties and provide guidance, but the responsibility of managing these day to day belongs to the individual SIG.
  - SIG Calendars and their settings
  - SIG Mailing List permissions, moderation, and maintenance
  - SIG YouTube playlist videos
  - SIG settings on their Zoom account
- Ownership, staffing, and moderation of Kubernetes communication properties. The CNCF may delegate admin or management positions on these properties to members of Contributor Experience:
  - [discuss.kubernetes.io] - This tool is mostly self-moderating and enough Kubernetes members participate that most of these roles are administrative.
  - [Mailing lists] / Google groups (this belongs to the each SIG)
  - [Slack] (This should be a new SIG)
  - [/kubernetescommunity] YouTube channel (This should be a shared responsibility with a new SIG)
  - [Zoom] (CNCF)
  - Ownership of the Kubernetes G suite (this is Steering Committee)
- Logistical Execution of the [Contributor Summit(s)]
  - The CNCF will provide us with an events manager who will handle the event logistics and provide guidance with the requirements we set forth. 

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
[Group Mentoring]:  https://git.k8s.io/community/mentoring/group-mentoring.md
[The 1:1 Hour]: https://git.k8s.io/community/mentoring/the1-on-1hour.md
[kubernetes/community]: https://git.k8s.io/community/
[Contributor Summit(s)]: https://git.k8s.io/community/events/2018/12-contributor-summit
[contributor summits]: https://git.k8s.io/community/events/2018/12-contributor-summit
[DevStats]: https://k8s.cncf.devstats.io
[kubernetes-sig-contribex@]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[kubernetes blog]: https://www.kubernetes.io/blog
[GitHub Project Organizational Management]: https://git.k8s.io/community/github-management
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