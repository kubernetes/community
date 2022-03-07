# 2021 Annual Report: SIG Contributor Experience

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - [Migrating K-Dev to a managed kubernetes.io account](https://github.com/kubernetes/community/issues/5877)
   - Steering Election with Elekto
      - Elekto was created by Manish Sahani as part of an LFX internship. It allows us to run Elections on our own infrastructure.
      - Voters can login using their github credentials and vote from the election site.
      - TODO : Write about how implementing elekto helped us overcome challenges with CiVS
   - Community meeting overhaul and reboot
   - EasyCLA migration
      - TODO : Describe impact, Value
   - Contributor Summit North America
   - Contributor Celebration
   - Restaffed up teams
      - Moderation
      - Youtube Admin
   - Ran three group mentoring cohorts to grow contributors in named roles (50% graduation rate) for contribex, docs, and a 'chair' cohort for multiple SIGs
   - Grew @k8scontributors twitter account to 5,700 followers
   - Created documentation geared towards Chairs and Tech Leads to help grow and support their roles and groups operations
   - Curated and ran 8 months of Leads meetings (16 meetings total)

2. What initiatives are you working on that aren't being tracked in KEPs?

Contribex is a service and program orientated SIG. Most of our initiatives cover long term services for the Kubernetes project.


  |         Subproject        |          Initiative / Program          |
  |:-------------------------:|:--------------------------------------:|
  | Community                 | Community Repo Stewardship             |
  | Community Management      | Annual Contributor Survey              |
  | Community Management      | Calendar Admin                         |
  | Community Management      | Leadership Operations                  |
  | Community Management      | discuss.k8s.io Admin                   |
  | Community Management      | Mailing List Admin                     |
  | Community Management      | Slack Admin                            |
  | Community Management      | Zoom / YouTube Admin                   |
  | Contributor Documentation | Contributor Guide Stewardship          |
  | Contributor Documentation | Contributor Site                       |
  | Contributor Documentation | Developer Guide Audit                  |
  | Contributor Documentation | Developer Guide Stewardship            |
  | Contributor Comms         | Contributor / SIG Profiling            |
  | Contributor Comms         | SIG Outreach and Support               |
  | Contributor Comms         | Contributor Events Outreach            |
  | Contributor Comms         | Stewardship of k8scontributors twitter |
  | Devstats                  | Devstats Dashboard Update              |
  | Events                    | Monthly Community Meeting              |
  | Events                    | Office Hours                           |
  | Events                    | Elections                              |
  | Events                    | Contributor Summits                    |
  | GitHub Management         | GitHub Admin / Moderation              |
  | GitHub Management         | GitHub Master -> Main rename           |
  | GitHub Management         | GitHub New Membership Coordinator      |
  | Mentoring                 | Google Summer of Code                  |
  | Mentoring                 | Group Mentoring                        |
  | Mentoring                 | LFX Mentor Program                     |
  | Mentoring                 | Meet our Contributors                  |
  | Mentoring                 | Outreachy                              |
  | Mentoring                 | New Contributor Workshop               |
  | Slack Infra               | slack-infra                            |

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)
   
   - GitHub Admin Subproject
      - We need to bring on more membership co-ordinators.
         - These membership co-ordinators will be current contributors.
         <!-- TODO : Link to requirements -->
         - A full list of requirements can be found here 
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
      - We are a service orientated sig, Issues are the main form of how we track work.
      - [Dashboard for SIG Contributor Experience issue velocity, Devstats](https://k8s.devstats.cncf.io/d/15/issues-age-by-sig-and-repository-groups?orgId=1&var-period=d7&var-repogroup_name=SIG%20Contributor%20Experience&var-repo_name=kubernetes%2Fkubernetes&var-repo=kuberneteskubernetes&var-sig_name=All&var-kind_name=All&var-prio_name=All&from=now-1y%2Fy&to=now-1y%2Fy_)
      

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
    - This work will reduce toil for not only contributor experience, but all groups in the project.

## Membership

Statistics were retrived on 2022-01-30

- Primary slack channel member count: 1895
- Primary mailing list member count: 366 
- Primary meeting attendee count (estimated, if needed): 10-15 
- Primary meeting participant count (estimated, if needed): 5
- Unique reviewers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
<!-- use dims tool, look over all the links to owners files, manual consolidation -->
- Unique approvers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
<!-- use dims tool, look over all links to owners files manual -->

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

New in $YYYY:
- [$wg-name](https://git.k8s.io/community/$wg-id/) ([$YYYY report](https://git.k8s.io/community/$wg-id/annual-report-$YYYY.md))
-

Retired in $YYYY:
- [$wg-name](https://git.k8s.io/community/$wg-id/) ([$YYYY report](https://git.k8s.io/community/$wg-id/annual-report-$YYYY.md))
-

Continuing:
- [$wg-name](https://git.k8s.io/community/$wg-id/) ([$YYYY report](https://git.k8s.io/community/$wg-id/annual-report-$YYYY.md))
-

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
   - TODO : Update the README to account for slack meetings
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
   - TODO : Update sigs.yaml
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
   - TODO : Review this
- [X] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
   - 
- [ ] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -
      - TODO : Gather links for all these

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-contributor-experience/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-contributor-experience/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
