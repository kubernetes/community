# Kubernetes Community Group Annual Reports

This document outlines the process for an annual reporting and
communication structure for Special Interest Groups (SIGs),
Working Groups (WGs), and Committees.
All policy updates will be in their respective [SIG] or [WG]
 governance docs as well as the general [governance] guidance.

## Goals
- Paint a complete project health picture for all of our community groups
- Create a feedback loop between Chairs, Tech Leads, Subproject Owners, WG
  Organizers, the community groups at large, and the Steering
Committee to move the project forward  
- Encourage dialogue about the wellbeing of the projects contributors and offer
suggested guidance and coaching
- Promote healthy, active, engaged community groups  
- Understand and have context before issues arise and celebrate wins where they  
should be highlighted  
- Help reshape project priorities at a high level

## Reporting Process  

Chairs and Organizers are responsible for compiling a yearly public report but
may be completed with the help of members of that group. Groups are encouraged
to make this an agenda item of topic to discuss together. Chairs and Organizers
ensure that reports are complete, accurate, and submitted to the Steering
Committee.

1. Early January (of the following year)
   * Steering Committee finalizes [questions] and generates draft 
     `annual-report-YYYY.md` templates for each group in the community repo
   * Steering Committee liaison reaches out to group leads to kick off
2. January-February
   * Chairs/Organizers work with their group (see tips below) to open
     a pull request to the group's documentation in the community repo,
     updating the draft to complete the questions
   * draft ready for review by February 14
   * post a call for comments period on the PR with your group's mailing list;
     you can follow your charter communication period if you have one,
     but no less than 72 hours
   * verify that Steering Liaison has reviewed
   * merge a completed report by March 1
3. March
   * Steering Committee produces a project-wide annual report,
     summarizing and highlighting elements from the individual group reports.
   * The Steering Committee liaison will work directly with groups that have 
     follow up items and update Steering during regular monthly meetings. The liaison
     will also coordinate time with the Chairs (as a group). If you'd like to meet 
     1:1 instead, please let your liaison know.
   * Draft summary for tech writer / editor review: March 15
   * Publication date on cncf.io/reports: March 30
   * The March edition of the "Chairs, Tech Leads, and Organizers" meetings will 
     be used as follow up for the community groups that have questions from/to 
     Steering and a retrospective

### Tips for Chairs and Working Group Organizers:      
- Work together with your groups roles and community members to complete;
suggestion: schedule a dedicated meeting or intentional agenda item to go over 
project health with the goal of compiling this report, delegate to subproject
owners or other community members
- All questions require a response. 
- End users and other members of the community will read these. Err on the side
of being more explicit than using our upstream shorthand or abbreviations. 
- The [questions] can be forked into whatever medium you wish to
collaborate with your community: gdocs, hackmd, etc - it will all land back into
a pull request at the end. 

## Questions for report:

### Special Interest Groups:

```
# $sig-name - $YYYY annual report

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - 
   - 
   - 

2. What initiatives are you working on that aren't being tracked in KEPs?

   - 
   - 
   - 

3. KEP work in $YYYY (1.x, 1.y, 1.z):

<!-- 
Generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.stable
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.stable
   - Beta
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.beta
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.beta
   - Alpha
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.alpha
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md) - $milestone.alpha
   - Pre-alpha
      - [$kep-number - $title](https://git.k8s.io/community/$link/README.md)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - 
   - 
   - 

2. What metrics/community health stats does your group care about and/or measure?

   - 
   - 
   - 

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing 
   to activities or programs that provide useful context or allow easy participation?

   - 

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - 

5. Does the group have contributors from multiple companies/affiliations?

   - 

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - 
   - 

## Membership

- Primary slack channel member count: 
- Primary mailing list member count: 
- Primary meeting attendee count (estimated, if needed): 
- Primary meeting participant count (estimated, if needed): 
- Unique reviewers for SIG-owned packages: {generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files}
- Unique approvers for SIG-owned packages: {generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files}

Include any other ways you measure group membership

## Subprojects

<!--
Generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in $YYYY:
- [$subproject-name](https://git.k8s.io/community/$sig-id#$subproject-name)
- 

Retired in $YYYY:
- [$subproject-name](https://git.k8s.io/community/$sig-id#$subproject-name)
- 

Continuing:
- [$subproject-name](https://git.k8s.io/community/$sig-id#$subproject-name)
- 

## Working groups

<!--
Generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
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

[ ] [README.md] reviewed for accuracy and updated if needed
[ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
    (or created if missing and your contributor steps and experience are different or more
    in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
[ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
[ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
[ ] Meeting notes and recordings for $YYYY are linked from [README.md] and updated/uploaded if needed
[ ] Did you have community-wide updates in $YYYY (e.g. kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
    - 
    - 

[CONTRIBUTING.md]: https://git.k8s.io/community/$sig-id/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/$sig-id/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
```

### Working Groups:

```
# $wg-name - $YYYY annual report

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - 
   - 
   - 

2. What initiatives are you working on that aren't being tracked in KEPs?

   - 
   - 
   - 

## Project health

1. What’s the current roadmap until completion of the working group?

   - 
   - 
   - 

2. Does the group have contributors from multiple companies/affiliations?

   - 

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - 
   - 

## Membership

- Primary slack channel member count: 
- Primary mailing list member count: 
- Primary meeting attendee count (estimated, if needed): 
- Primary meeting participant count (estimated, if needed): 

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

[ ] [README.md] reviewed for accuracy and updated if needed
[ ] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
[ ] Meeting notes and recordings for $YYYY are linked from [README.md] and updated/uploaded if needed
[ ] Updates provided to sponsoring SIGs in $YYYY
    - [$sig-name](https://git.k8s.io/community/$sig-id/)
      - links to email, meeting notes, slides, or recordings, etc
    - [$sig-name](https://git.k8s.io/community/$sig-id/)
      - links to email, meeting notes, slides, or recordings, etc
    - 

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/$wg-id/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
```

### Thanks   
Thanks to the Apache Software Foundation for their open guidance on PMC 
reporting, the many PMCs that have shared their experiences, and the Kubernetes
community for collaboration.
https://www.apache.org/foundation/board/reporting


[SIG]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[WG]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[governance]: https://git.k8s.io/community/governance.md
[questions]: #questions-for-report
