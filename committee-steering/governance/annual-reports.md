# Kubernetes Community Group Annual Reports

This document outlines the process for an annual reporting and
communication structure for Special Interest Groups (SIGs), Working Groups (WGs), and Committees.
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

## Questions for report

### Special Interest Groups

The questions are available as a [go-template](https://pkg.go.dev/text/template) file [here](../../generator/annual-report/sig_report.tmpl).

### Working Groups:

The questions are available as a [go-template](https://pkg.go.dev/text/template) file [here](../../generator/annual-report/wg_report.tmpl).


### Thanks
Thanks to the Apache Software Foundation for their open guidance on PMC
reporting, the many PMCs that have shared their experiences, and the Kubernetes
community for collaboration.
https://www.apache.org/foundation/board/reporting


[SIG]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[WG]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[governance]: https://git.k8s.io/community/governance.md
[questions]: #questions-for-report
