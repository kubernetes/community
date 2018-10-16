# Contributor Experience Special Interest Group Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

The [Contributor Experience Special Interest Group] (SIG) is responsible for improving the experience of those who upstream contribute to the Kubernetes project. We do this by creating, and maintaining programs and processes that promote community health and reduce project friction, while retiring those programs and processes that don't. Being conscientious of our contributor base is critical to scaling the project, growing the ecosystem, and helping the project succeed.

We do this by listening - whether it’s through our roadshows to SIG meetings, surveys, data, or [GitHub issues], we take in the feedback and turn it into our [project list]. We build a welcoming and inclusive community of contributors by giving them places to be heard and productive.

### In scope

#### Code, Binaries and Services

- Establish policies, standards, and procedures:
  - routine GitHub management tasks, including but not limited to: org membership, org permissions, repo creation/administration.
  - "Org Owner" GitHub permissions, and grant/limit these privileges accordingly
  - bot accounts, service accounts, webhooks, and third-party integrations for all communication platforms that we support but most importantly, GitHub.
- Establishing a "GitHub Administration team" that will oversee the execution of GitHub management tasks: inviting new members to the org, creating repos, executing moderation decisions, auditing permissions.
- Work with other SIGs and interested parties in the project to execute GitHub tasks where required
-Own and execute events that are targeted to the Kubernetes contributor community, including:
  - The weekly [Kubernetes Community meeting]
  - [Contributor Summit(s)]
  - SIG Contributor Experience face to face meetings
  - [Steering Committee elections] (though we do not own policy creation, see 'out of scope' below)
  - Retrospective moderation for other SIGs upon request
  - Other events, like other SIG face to face events, upon request and consideration
- Strategize, build, and execute on scalable [mentoring programs] for all contributor levels. These may include:
  - [Google Summer of Code]
  - [Outreachy]
  - [Meet Our Contributors]
  - [Group Mentoring - WIP]
  - [The 1:1 Hour - WIP]
  - Speed Mentoring at KubeCon
- Create best practices, trainings, policies, and [moderation] of Kubernetes communication platforms, the transparency vehicles of information throughout the project. This includes user guidelines when the contributor is the user of the platform. Moderation is an always on, never off service that we supply with trusted contributors spanning time zones. This includes immediate action when dealing with code of conduct issues in a public space that are defined below:
    - All mailing lists/googlegroups where we have ownership, including [kubernetes-dev@]. At least one Owner of the mailing list must be a [communication]s subproject OWNER.
    - [Zoom] via CNCF funding
    - YouTube via the [/kubernetescommunity] channel
    - [Slack] via CNCF funding
    - [discuss.kubernetes.io] via CNCF funding
- Help onboard new contributors and current into the culture, workflow, and CI of our contributor experience through the [contributor guide], other related documentation, [contributor summits] and programs tailored to onboarding.
- Perform issue triage on and maintain the [kubernetes/community] repository.  
- Help SIGs with being as transparent and open as possible through creating best practices, guidelines, and general administration of YouTube, Zoom, and our mailing lists where applicable
- Assist SIGs/WG Chairs and Technical Leads with organizational management operations as laid out in the [sig-governance] doc
- Distribute contributor related news on various Kubernetes channels, including Cloud Native Compute Foundation [CNCF] for posting blogs, social media, and other platforms as needed.
- Establish and share metrics to measure project health, community health, and general trends, including:
  - ongoing work with the CNCF to improve [DevStats]
  - the contributor experience survey(s)
  - engagement on project platforms that we manage
- Encourage automation to improve productivity for contributors where it makes sense and consult with SIG Testing if automation is covering GitHub workflows.
- Research other OSS projects and consult with their leaders about contributor experience topics to make sure we are always delivering value to our contributors.
- Provide retrospective hosting services on request to SIGs

#### Cross-cutting and Externally Facing Processes

We cross-cut all SIGs and WGs to deliver the following processes:

Communicating (proposed or executing on approved) process change:
    - Socialize on kubernetes-sig-contribex@googlegroups.com and on our weekly update calls
    - If the change impacts a small number (<4) of SIGs, WGs, or repos, we will go to each lead, their mailing lists, slack channel, and/or their update meetings and ask for feedback and a [lazy consensus] process. We will follow up with a post to [kubernetes-dev@]googlegroups.com mailing list if it affects more than us.
    - In cases of impacts across a large number of areas and/or project wide, we will:
        - [Lazy consensus] with a time box of at least 72 hours to the following mailing lists with the GitHub issue link including the subject [NOTICE]: foobar or [ANNOUNCE]: to the following mailing lists in order:
            - [kubernetes-sig-contribex@]googlegroups.com
            - sig-leads@googlegroups.com
            - [kubernetes-dev@]googlegroups.com
        - Announce at weekly Kubernetes Community Meeting on Thursdays
        - Depending on how wide of an ecosystem change this is, we may also slack, blog, tweet, and use other channels to get the word out.
        - Our standard time box is three (3) days (72 business hours); however, there may be situations where we need to act quickly but the time period will always be clear and upfront.

If we need funding for any reason, we approach [CNCF].
CNCF in many of the noted cases above, contributes funding to our platforms, processes, and/or programs. They do not play a day-to-day operations role and have bestowed that to our community to run as we see fit. Since they do fund some of our initatives, this means that they hold owner account privileges on certain platforms like Zoom and Slack. In these cases, such as Slack, there are at least two Contributor Experience [communication] subproject OWNERs listed as admins.

### Out of scope

- Code for the testing and CI infrastructure - that’s SIG Testing
- [kubernetes/community]  ownership of folders for KEPs and Design Proposals. Members are to follow those folders owners files and SIG leadership for the specific issue/PR in question.
- User community management. We hold office hours because contributors are a large portion of the volunteers that run that program.
- The contributor experience for repos not included in the Kubernetes [associated repositories]list.
- Steering committee election policy updates and maintenance.
- We do not create SIGs/WGs but can assist in the various community management needs of their micro communities that would kick off their formation and keep them going.
- We are not the [code of conduct committee] and therefore do not control incident management reporting or decisions; however, our moderation guidelines allow us to act swiftly if there is a clear violation of terms on one of our supported platforms. If there is an action that the committee needs to take that involves one of these platforms (example: the removal of someone from GitHub), we will carry that out if none of the committee members have access.
- Communication platforms that are out of our scope for maintenance and support but we may still have some influence:
    - [r/kubernetes]
    - [kubernetesio@] twitter account
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

### Subproject Creation - TODO
Pick one:
SIG Technical Leads
Federation of Subprojects

[sig-governance]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://git.k8s.io/community/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus
[Contributor Experience Special Interest Group]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[kubernetes-dev@]: https://groups.google.com/forum/#!forum/kubernetes-dev
[kubernetesio@]: https://www.twitter.com/kubernetesio
[r/kubernetes]: https://kubernetes.reddit.com
[Google Summer of Code]: https://git.k8s.io/community/mentoring/google-summer-of-code.md
[Outreachy]: https://git.k8s.io/community/mentoring/outreachy.md
[Meet Our Contributors]:  https://git.k8s.io/community/mentoring/meet-our-contributors.md
[Group Mentoring - WIP]:  https://git.k8s.io/community/mentoring/group-mentoring.md
[The 1:1 Hour - WIP]: https://git.k8s.io/community/mentoring/the1-on-1hour.md
[kubernetes/community]: https://git.k8s.io/community/
[Contributor Summit(s)]: https://git.k8s.io/community/events/2018/12-contributor-summit
[DevStats]: https://k8s.cncf.devstats.io
[kubernetes-sig-contribex@]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[kubernetes blog]: https://www.kubernetes.io/blog
[associated repositories]: https://git.k8s.io/community/github-management
[communication]: https://git.k8s.io/community/communication
[CNCF]: https://cncf.io
[GitHub issues]: https://github.com/kubernetes/community/issues
[project list]: https://github.com/orgs/kubernetes/projects/1
[Kubernetes Community meeting]: https://git.k8s.io/community/
[mentoring programs]: https://git.k8s.io/community/mentoring
[Steering Committee elections]: https://git.k8s.io/community/
[Slack]: https://git.k8s.io/community/communication/slack-guidelines.md
[Zoom]: https://git.k8s.io/community/communication/zoom-guidelines.md
[/kubernetescommunity]: https://www.youtube.com/kubernetescommunity
[discuss.kubernetes.io]: https://discuss.kubernetes.io
[moderation]: https://git.k8s.io/community/communication/moderation.md
[contributor guide]: https://git.k8s.io/community/contributors/guide
