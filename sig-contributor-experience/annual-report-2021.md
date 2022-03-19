# 2021 Annual Report: SIG Contributor Experience

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Continuing to provide critical services to enable the work of over 75,000+ contributors
   - [Migrating K-Dev to a managed kubernetes.io account](https://github.com/kubernetes/community/issues/5877)
      -  We encountered a problem with the k-dev mailing list where we couldn't send out community-wide meeting invites, due to restrictions on public google groups.
      - In order to overcome this issue, we migrated to Google Workspace.
      - It took over 80+ hours across the Contribex team to complete this migration.
   - Steering Election with Elekto
      - Elekto was created by Manish Sahani as part of an LFX internship. It allows us to run Elections on our own infrastructure.
      - Voters can login using their github credentials and vote from the election site.
      - The previous method of holding elections was prone to issues when it came to the following:
         - Obtaining the list of elegible voters from 75,000+ contributors 
         - Sending ballots out via email
      - Elekto automated the proccess of obtaining the list of elegible voters
      - Ballots are no longer sent out via email, and instead contributors log in to the elections site to submit their ballots.
      - More information can be found in [community issue #5096](https://github.com/kubernetes/community/issues/5096)
   - Community meeting overhaul and reboot  
   - [EasyCLA migration](https://github.com/kubernetes/org/issues/2778)
      - The CLA service provided to us by the Linux Foundation was being deprecated.
      - We needed to migrate to EasyCLA2 to ensure continuity of service for our 75,000+ active committers.
   - Contributor Summit North America
   - Contributor Celebration
   - Restaffed teams
      - Moderation
      - Youtube Admin
   - Ran three group mentoring cohorts to grow contributors in named roles (50% graduation rate) for contribex, docs, and a 'chair' cohort for multiple SIGs
   - Grew @k8scontributors twitter account to 5,700 followers
   - Created documentation geared towards Chairs and Tech Leads to help grow and support their roles and groups operations
   - Curated and ran 8 months of Leads meetings (16 meetings total)

2. What initiatives are you working on that aren't being tracked in KEPs?

Contribex is a service and program orientated SIG. Most of our initiatives cover long term services for the Kubernetes project.

|                                                  **Subproject**                                                 |                                                **Initiative / Program**                                               |
|:---------------------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------:|
| [Community](https://git.k8s.io/community/sig-contributor-experience#community)                                  | [Community Repo Stewardship](https://git.k8s.io/community)                                                            |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | Calendar Admin                                                                                                    |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | Leadership Operations                                                                                             |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | [discuss.k8s.io End User Forum Admin](https://discuss.k8s.io)                                                         |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | [Mailing List Admin](https://k8s.dev/docs/comms/moderation/)                                                          |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | [Slack Admin](https://k8s.dev/docs/comms/slack/)                                                                      |
| [Community Management](https://git.k8s.io/community/sig-contributor-experience#community-management)            | [Zoom](https://k8s.dev/docs/comms/zoom) / [YouTube Admin](https://k8s.dev/docs/comms/youtube/#admin-responsibilities) |
| [Contributor Documentation](https://git.k8s.io/community/sig-contributor-experience#contributors-documentation) | [Contributor Guide Stewardship](https://k8s.dev/guide)                                                                |
| [Contributor Documentation](https://git.k8s.io/community/sig-contributor-experience#contributors-documentation) | [Contributor Site](https://git.k8s.io/contributor-site)                                                               |
| [Contributor Documentation](https://git.k8s.io/community/sig-contributor-experience#contributors-documentation) | [Developer Guide Audit](https://github.com/kubernetes/community/issues/5229)                                          |
| [Contributor Documentation](https://git.k8s.io/community/sig-contributor-experience#contributors-documentation) | [Developer Guide Stewardship](https://github.com/kubernetes/community/tree/master/contributors/devel)                 |
| [Contributor Comms](https://git.k8s.io/community/sig-contributor-experience#contributor-comms)                  | Contributor / SIG Profiling                                                                                           |
| [Contributor Comms](https://git.k8s.io/community/sig-contributor-experience#contributor-comms)                  | SIG Outreach and Support                                                                                              |
| [Contributor Comms](https://git.k8s.io/community/sig-contributor-experience#contributor-comms)                  | Contributor Events Outreach                                                                                           |
| [Contributor Comms](https://git.k8s.io/community/sig-contributor-experience#contributor-comms)                  | [Stewardship of k8scontributors twitter](https://twitter.com/k8scontributors)                                         |
| [Devstats](https://git.k8s.io/community/sig-contributor-experience#devstats)                                    | [Devstats Dashboard Update](https://github.com/cncf/devstats/issues/289)                                              |
| [Events](https://git.k8s.io/community/sig-contributor-experience#events)                                        | Monthly Community Meeting                                                                                             |
| [Events](https://git.k8s.io/community/sig-contributor-experience#events)                                        | Office Hours                                                                                                          |
| [Events](https://git.k8s.io/community/sig-contributor-experience#events)                                        | [Elections](git.k8s.io/community/events/elections)                                                                    |
| [Events](https://git.k8s.io/community/sig-contributor-experience#events)                                        | [Contributor Summits](https://k8s.dev/events/past-events/2021/)                                                       |
| [GitHub Management](https://git.k8s.io/community/sig-contributor-experience#github-management)                  | [GitHub Admin / Moderation](https://git.k8s.io/community/github-management#github-management)                         |
| [GitHub Management](https://git.k8s.io/community/sig-contributor-experience#github-management)                  | [GitHub Master -> Main rename](https://github.com/kubernetes/org/issues/2222)                                         |
| [GitHub Management](https://git.k8s.io/community/sig-contributor-experience#github-management)                  | [GitHub New Membership Coordinator](https://git.k8s.io/community/github-management/README.md#other-roles)             |
| [Mentoring](https://git.k8s.io/community/sig-contributor-experience#mentoring)                                  | [Group Mentoring](https://git.k8s.io/community/mentoring/programs/group-mentoring.md)                                 |
| [Mentoring](https://git.k8s.io/community/sig-contributor-experience#mentoring)                                  | [LFX Mentor Program](https://git.k8s.io/community/mentoring/programs/lfx-mentoring.md)                                |
| [Slack Infra](https://git.k8s.io/community/sig-contributor-experience#slack-infra)                              | [slack-infra](https://sigs.k8s.io/slack-infra)                                                                        |

  We put the following initiatives on hold, and will be revisiting these at a later date.

  |         Subproject        |          Initiative / Program          |
  |:-------------------------:|:--------------------------------------:|
  | Mentoring                 | Meet our Contributors                  |
  | Mentoring                 | Outreachy                              |
  | Mentoring                 | New Contributor Workshop               |
  | Mentoring                 | Google Summer of Code                  |

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)
   
   - GitHub Admin Subproject
      - We need to bring on more [new membership co-ordinators](https://cs.github.com/kubernetes/community/blob/cf77df8935eb5460d8d2856c40009d981691ec2f/github-management/README.md#L49).
         - New Membership Coordinators are current contributors to the Kubernetes project that help serve as a friendly face to newer, prospective community members, guiding them through the process to request membership to a Kubernetes GitHub organization.
   - Community Management Automation
      - [Zoom to Youtube](https://github.com/kubernetes/community/issues/5201)
      - [Workspace Automation](https://github.com/kubernetes/steering/issues/213)
         - Mailing list management
         - Calendar management
         - Shared drives
   - Mentoring Program Management and new Roles
      - We need a [Group Mentoring Coordinator](https://github.com/kubernetes/community/issues/6517)
      - We need a [3rd Party Mentoring Coordinator](https://github.com/kubernetes/community/issues/6471)

2. What metrics/community health stats does your group care about and/or measure?

   - Issue Velocity
      - [We are a service orientated sig, Issues are the main form of how we track work.](https://k8s.devstats.cncf.io/d/73/inactive-issues-by-sig?orgId=1&var-sigs=%22contributor-experience%22&from=1609459200000&to=1640995199000)
      

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - We don't have a CONTRIBUTING.md as our SIG's contributing proccess does not differ from the proccess in the Contributor Guide.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A

5. Does the group have contributors from multiple companies/affiliations?

   - Yes. [25+ Different groups](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Contributor%20Experience&var-repo_name=kubernetes%2Fkubernetes&var-companies=All&from=1577854800000&to=1640926800000&viewPanel=1)

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - A full time community manager who can maintain several subprojects and help with the overall direction of the group
   - More full time support for the items under bullet point 1.
    - This work will reduce toil for contributor experience and all groups in the project.

## Membership

Statistics were retrived on 2022-01-30

- Primary slack channel member count: 1895
- Primary mailing list member count: 366 
- Primary meeting attendee count (estimated, if needed): 10-15 
- Primary meeting participant count (estimated, if needed): 5
- Unique reviewers for SIG-owned packages: 20
- Unique approvers for SIG-owned packages: 28

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
   N/A
Retired in 2021:
   N/A

Continuing:

 - community
Owns and manages overall community repo, including community group documentation and operations.
 - community-management
Manages operations and policy for upstream community group communication platforms.
 - contributor-comms
Contributor Communications focuses on amplifying the success of Kubernetes contributors through marketing.
 - contributors-documentation
writes and maintains documentation around contributing to Kubernetes, including the Contributor's Guide, Developer's Guide, and contributor website.
 - devstats
Maintains and updates https://k8s.devstats.cncf.io, including taking requests for new charts.
 - events
Creates and runs contributor-focused events, such as the Contributor Summit.  Event Teams are part of this subproject.
 - github-management
Manages and controls Github permissions, repos, and groups, including Org Membership.
 - mentoring
Oversees and develops programs for helping contributors ascend the contributor ladder, including the New Contributor Workshops, Meet Our Contributors, and other programs.
 - slack-infra
Creates and maintains tools and automation for Kubernetes Slack.
<!-- BEGIN CUSTOM CONTENT -->

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
N/A

Retired in 2021:
- [wg-naming](https://git.k8s.io/community/archive/wg-naming)

Continuing:
N/A

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
   - TODO : Update the README to account for slack meetings
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
   - [SIG Contributor Experience Deep Dive - KubeCon EU 2021](https://www.youtube.com/watch?v=vPK3QmVOE4Y)
   - [SIG Contributor Experience Deep Dive - KubeCon NA 2021](https://www.youtube.com/watch?v=QOiyWWFjG5Q)


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-contributor-experience/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-contributor-experience/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
