# SIG Node Annual Report 2021

## Operational

* How are you doing with operational tasks in
[sig-governance.md](https://git.k8s.io/community/committee-steering/governance/sig-governance.md)?
  * Is your README accurate? have a CONTRIBUTING.md file?
    * Yes, [README.md](https://github.com/kubernetes/community/blob/master/sig-node/README.md) is accurate
    * Opened [issue](https://github.com/kubernetes/community/issues/5600) to track completion CONTRIBUTING.md
  * All subprojects correctly mapped and listed in [sigs.yaml](https://git.k8s.io/community/sig-list.md)?
    * [Yes](https://github.com/kubernetes/community/tree/master/sig-node#subprojects)
  * What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping
  recordings up to date/trends in community members watching recordings?
    * The SIG operates two recurring meetings.  The Regular SIG Meeting occurs weekly.  It focuses on top-level trends for the SIG, and provides a forum to discuss enhancements, blocked PRs, and release plans.  Each meeting has approximately 15-30 participants depending on the topic.  The meetings are relatively active.  The Chairs host the SIG meeting.  A member of the SIG reports on velocity metrics, and tries to draw attention to overall trends.  The agenda is open for members of the community to discuss an idea or enhancement.  Discussion may ensue for each topic.  SIG members that have participated in the SIG for an extended period will share historical context.  Recordings are up-to-date.  The Weekly CI meeting focuses on improving CI health and improving the number of SIG participants that have sufficient understanding of our test apparatus.  Each meeting has ~7 participants, but participation may improve with new scheduled time.  Each meeting focuses on supporting active triage and review of the project board.  [Recordings are kept up to date with regular SIG meetings](https://www.youtube.com/playlist?list=PL69nYSiGNLP1wJPj5DYWXjiArF-MJ5fNG).
* How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being
retired? Are OWNERS.md files up to date in these areas?
  * The regular SIG meetings provide a forum to discuss sub-projects closely tied to the Kubernetes release.  This includes cri-api, cri-tools, kubelet, and node-api.  A major focus in the SIG this year has been the graduation of existing features.  Other sub-projects get less frequent updates as development velocity has slowed as the project reached sufficient maturity.
* Same question as above but for working groups.
  * There is no regular interlock between working groups and the SIG.  SIG members based on interest and/ or desire may attend particular working group meetings.  The SIG did create a dedicated set of meetings to talk through particular areas like topology management for a limited period of time in order to provide room for discussion outside of the regular SIG meeting.  Once consensus was reached, those meetings ceased.  
* When was your last monthly community-wide update? (provide link to deck and/or recording)
  * Last presentation was on April 16, 2020:
    * [Slides](https://docs.google.com/document/d/1VQDIAB0OqiSjIHI8AWMvSdceWhnz56jNpZrLs6o7NJY/edit#heading=h.di6sf3cdf3yr)

## Membership

* Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  * For the SIG leads yes.
  * For subproject owners, likely there are some inactive maintainers, but it’s up to the subproject to prune their OWNERs files.
  * For 2021, the SIG will look to audit the set of sponsored sub-projects to ensure OWNERs files are updated.
* How do you measure membership? By mailing list members, OWNERs, or something else?
  * Anyone can be considered a SIG member if they join the Zoom calls or general discussions regularly.
* How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
  * The SIG identifies the set of items they wish to focus on each release.  A primary engineer and approver is identified up-front.  This provides a baseline understanding of responsibility and capacity going into the release.  As we near release milestones, we measure where we are relative to our planned goals and ensure attention is redirected where appropriate.
  * The SIG suffered from atrophy in CI health and awareness.  Members of the SIG have developed a dedicated CI group to focus on health and train up new community members to assist.  The group started meeting in May 2020 with ~30 individuals volunteered to assist.  The group has created a large amount of learning materials, improved the CI health, and instituted project boards and triage processes.  The sub-group meets each week with ~7 regular attendees as the CI health has stabilized.
  * The SIG has elevated 3 members to approver status for the kubelet sub-project in the past year.  It has shifted 2 members to emeritus status.  The SIG has members who have expressed interest and desire to grow their scope in particular parts, and we try to encourage that over a set of releases by focusing on particular sub-components.  We anticipate growing more approvers over the calendar year.
* Is there a healthy onboarding and growth path for contributors in your SIG?
What are some activities that the group does to encourage this? What programs are you participating in to grow contributors
throughout the contributor ladder?
  * The SIG identifies the set of items that we would like to tackle each release and clearly identifies the primary engineering owner and primary approver.  We try to pair new engineers with senior approvers to grow scope and comfort.  In the CI health meetings, we have established a project board to try and provide directional guidance on where and how to best contribute.
* What programs do you participate in for new contributors?
  * KubeCon updates with “how to get involved”, beginning at KubeCon EU 2021
  * Improving documentation and process so new individuals can join, see e.g. [Node Project Board](https://github.com/orgs/kubernetes/projects/49), [Node PR Triage Guidelines](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md)
  * Mentoring group for underrepresented individuals to become reviewers, first cohort (July-Dec. 2020) run by @dims and second cohort (Jan-Mar. 2021) run by @ehashman
  * Next cohort leader/participants TBD  
* Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that
they currently are not?
  * The group has contributors from multiple companies/affiliations.
  * [39 companies made 1+](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Node&var-companies=All) contributions over the last year (70 if all activity is counted).
  * Deprecating of dockershim is a big project that will require many end user to make a step to migrate and some companies who are not active on SIG Node, namely monitoring vendors, to help end users with this migration.

## Current initiatives and project health

* [x] Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members.
* What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway?
What are some of the longer tail projects that your group is working on?
  * Better tracking of [PRs setup](https://github.com/orgs/kubernetes/projects/49) and [Node PR Triage Guide](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md) by @ehashaman
  * cgroups v2
  * [Graduating features](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#)
  * Topology manager / device alignment
* Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
  * The enumerated links list graduated KEPs for each of the related releases.
  * 1.18: [GA-ed KEPs](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+milestone%3Av1.18+label%3Asig%2Fnode)
  * 1.19: [GA-ed KEPs](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+milestone%3Av1.19+label%3Asig%2Fnode+)
  * 1.20: [GA-ed KEPs](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+milestone%3Av1.20+label%3Asig%2Fnode+)
* What areas and/or subprojects does the group need the most help with?
  * The SIG could always use more help in sustaining CI.
* What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
  * Weekly sig meeting starts with an overview of how PRs were opened/reviewed/approved.
  * The reporting is generally applicable to other SIGs and is based on following [template](https://docs.google.com/document/d/1JOXKBDgXmQzz8YQSYa7XYcfVteM79iMtvId1aQXC1e8/edit)
  * The goal is to keep the reviews going so the numbers don’t spin out of control.  It also provides an overview of the week
  as the scope of the SIG is large and it may be hard to keep up with all changes.
  * [DevStats report](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=y&var-repogroup_name=SIG%20Node&var-apichange=All&var-size_name=All&var-kind_name=All)
