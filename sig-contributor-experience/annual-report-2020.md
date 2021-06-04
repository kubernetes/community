# Special Interest Groups: Contributor Experience

#### Operational 

- How are you doing with operational tasks in [SIG]-governance.md?
  - Is your README accurate? have a CONTRIBUTING.md file?
      - The README is up to date and accurate, we do not have a
        CONTRIBUTING.md file.
  - All subprojects correctly mapped and listed in sigs.yaml?
      - Yes
  - What’s your meeting culture? Large/small, active/quiet, learnings? Meeting 
    notes up to date? Are you keeping recordings up to date/trends in community 
    members watching recordings?
      - We have biweekly zoom meetings. It is small, but active.
      - The APAC-friendly zoom meeting was poorly attended so we are
      discontinuing it.
      - We are also hosting biweekly asynchronous meetings on slack, which has
      made it easier for members from APAC regions to participate.
      - Meeting notes are up to date, recordings of the meetings are published.
        The YouTube recordings often have <10 views with jumps to 30~ when they
        are featured on the front page of the channel. Collectively during 2020
        the SIG ContribEx playlist had 167 views with an average watch time of
        00:03:55. Compared to the previous year, this was an increase in views,
        but a decrease in average watch time.
- How does the group get updates, reports, or feedback from subprojects? Are
  there any springing up or being retired? Are OWNERS.md files up to date in
  these areas?
  - The weekly meeting (zoom or in slack) serves as a subproject status update.
    The notes are then sent to the mailing list.
  - A [subproject owner audit][audit] was completed in the first half of 2020
  - Some processes/meetings have been retired or retooled to reduce administrative
    overhead and become more async friendly.
- Same question as above but for working groups.
  - There is no formal reporting mechanism from the working groups. The updates
    largely come from issues or discussions that bubble up into the SIG.
  - SIG Contributor Experience TLs also participate in WG Infra directly, so we haven't
    found the need for a formal reporting process.
- When was your last public community-wide update? (provide link to deck and/or
  recording)
  - August 2020 at KubeCon Europe Virtual - [link] to the session.
    A similar Contribex session is also planned for KubeCon Europe 2021.
  - July 2020 in the community meeting. Links to [updates] are added to the SIG ContribEx README.

[link]: https://kccnceu20.sched.com/event/c9yh/intro-contributor-experience-sig-jorge-castro-vmware-bob-killen-university-of-michigan
[audit]: https://github.com/kubernetes/community/issues/4585
[updates]: https://github.com/kubernetes/community/tree/master/sig-contributor-experience#current-status


#### Membership

- Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  - No, there are several subproject owners that have become inactive and should
    be moved to emeritus status. An audit is planned for 2021. 
- How do you measure membership? By mailing list members, OWNERs, or something
  else?
  - Those that are actively engaged with meetings
- How does the group measure reviewer and approver bandwidth? Do you need help
  in any area now? What are you doing about it?
  - We do not actively monitor reviewer/approver bandwidth in a formalized way.
    However, our [relative PR velocity][vel] is in a good state. We could use more
    reviewers/approvers to help reduce risk and improve long-term sustainability.
- Is there a healthy onboarding and growth path for contributors in your SIG?
  What are some activities that the group does to encourage this? What programs
  are you participating in to grow contributors throughout the contributor ladder?
  - We have lost many contributors over the past year. General calls for help did
    not attract many new contributors, but with some additional outreach and
    raising the problem areas in an [easier to see way][pr] has attracted new folk.
  - We are actively applying the `good-first-issue` and `help-wanted` labels to issues
    and are closely mentoring new contributors who pick up these issues.
- What programs do you participate in for new contributors?
  - We have previously participated with Outreachy and Season of Docs.
- Does the group have contributors from multiple companies/affiliations? Can end
  users/companies contribute in some way that they currently are not?
  - Yes, we have multiple avenues for contributing for both code and non-code
    projects alike.
 

#### Current initiatives and project health

- What are initiatives that should be highlighted, lauded, shout outs, that
  your group is proud of? Currently underway? What are some of the longer tail
  projects that your group is working on?
  - The contributor website (https://k8s.dev)
  - Triage process improvements
  - The 2020 Contributor Celebration
  - Retired the Kubernetes-Incubator Org
  - In Progress: Migrating the default branch on GitHub from `master` to `main`
  - In Progress: Revamping the prow approval plugin to support granual approvals
- Year to date KEP work:
  - [KEP-1553: Issue Triage Workflow and Automation][kep]
  - In Progress: KEP for revamping the prow approval plugin
- What initiatives are you working on that aren't being tracked in KEPs?

  |         Subproject        |          Initiative / Program          |
  |:-------------------------:|:--------------------------------------:|
  | Community                 | Community Repo Stewardship             |
  | Community Management      | Annual Contributor Survey              |
  | Community Management      | Calendar Admin                         |
  | Community Management      | Chair and TL Meetings + Docs           |
  | Community Management      | Discuss Admin                          |
  | Community Management      | Mailing List Admin                     |
  | Community Management      | Slack Admin                            |
  | Community Management      | Zoom / YouTube Admin                   |
  | Community Management      | Zoom / YouTube Automation (zapier)     |
  | Contributor Documentation | Contributor Guide Stewardship          |
  | Contributor Documentation | Contributor Site                       |
  | Contributor Documentation | Developer Guide Audit                  |
  | Contributor Documentation | Developer Guide Stewardship            |
  | Contributor Documentation | Season of Docs                         |
  | Contributors Comms        | Contributor / SIG Profiling            |
  | Contributors Comms        | Stewardship of k8scontributors twitter |
  | Devstats                  | Devstats Dashboard Update              |
  | Events                    | Monthly Community Meeting              |
  | Events                    | Office Hours                           |
  | GitHub Management         | GitHub Admin / Moderation              |
  | GitHub Management         | GitHub Master -> Main rename           |
  | GitHub Management         | GitHub New Membership Coordinator      |
  | Mentoring                 | 1:1 Hour                               |
  | Mentoring                 | Google Summer of Code                  |
  | Mentoring                 | Group Mentoring                        |
  | Mentoring                 | LFX Mentor Program                     |
  | Mentoring                 | Meet our Contributors                  |
  | Mentoring                 | New Contributor Workshop               |
  | Mentoring                 | Outreachy                              |
  | Slack Infra               | slack-infra                            |
  
  **source:** [ContribEx Initatives and Prioritization spreadsheet][pr]

- What areas and/or subprojects does the group need the most help with?
  - Many need help. These are being tracked separately in the 
    [ContribEx Initatives and Prioritization spreadsheet][pr].
- What metrics/community health stats does your group care about and/or measure?
  - The only dashboard we use at this point in time is the
    [PR workloads table][vel] in devstats.



[vel]: https://k8s.devstats.cncf.io/d/34/pr-workload-per-sig-table?orgId=1&var-period_name=Last%20year
[kep]: https://git.k8s.io/enhancements/keps/sig-contributor-experience/1553-issue-triage
[pr]: https://docs.google.com/spreadsheets/d/1glhdFcUdqYAByW16hujxK1X_0k9mt_nrkCO4POeDNbs/edit#gid=0

