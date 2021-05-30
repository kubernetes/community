# Kubernetes SIG Release Community Group Annual Reports 2021

This report reflects back on CY 2020 and was written in April 2021.

## Operational

- How are you doing with operational tasks in [SIG]-governance.md?
  - Is your README accurate? have a CONTRIBUTING.md file?
    - The README is up to date and accurate, we do not have a CONTRIBUTING.md
      file.
  - All subprojects correctly mapped and listed in sigs.yaml?
    - Yes
  - What’s your meeting culture? Large/small, active/quiet, learnings? Meeting
    notes up to date? Are you keeping recordings up to date/trends in community
    members watching recordings?
    - We have weekly zoom meetings with a duration of 45min: one for the overall
      SIG and one for the Release Engineering subproject. Both meetings have ~20
      people, where sometimes new contributors or members of other SIGs join as
      well.
    - We refined the meeting structure to have a fixed 20 minute block to walk
      our project boards.
    - There is a timebox of 20minutes for informal status updates, too.
    - The last 5 minutes are free for open discussion.
    - Meeting notes are up to date and all meetings are recorded.
- How does the group get updates, reports, or feedback from subprojects? Are
  there any springing up or being retired? Are OWNERS.md files up to date in
  these areas?
  - The SIG Release meeting serves updates for the release cycle as well as the
    Engineering subproject.
  - The release Engineering meeting give more in-depth details about the
    technical topics within the SIG.
  - We have an additional agreement how to work together in terms of [Release
    Engineering](https://github.com/kubernetes/sig-release/tree/master/release-engineering#release-engineering)
- Same question as above but for working groups.
  - We do not have any working groups within the SIG.
- When was your last public community-wide update? (provide link to deck and/or
  recording)
  - KubeCon NA 2019: https://www.youtube.com/watch?v=sQuxWeVlrJQ

#### Membership

- Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  - Yes
- How do you measure membership? By mailing list members, OWNERs, or something
  else?
  - Those that are contributing and actively engaged with the meetings.
- How does the group measure reviewer and approver bandwidth? Do you need help
  in any area now? What are you doing about it?
  - We do not actively monitor reviewer/approver bandwidth in a formalized way.
    We walk the project boards during our meetings and triage incoming issues
    and PRs in dedicated sessions. We also have to do the cherry-pick reviews
    for patches, which is done by the release managers in a continuous process.
- Is there a healthy onboarding and growth path for contributors in your SIG?
  What are some activities that the group does to encourage this? What programs
  are you participating in to grow contributors throughout the contributor ladder?
  - We have the SIG Release Shadowing Program for new members to join the
    release cycle.
  - There is the opportunity to become a Release Manager (associate) within the
    Release Engineering subproject, too.
  - The SIG has Technical Lead and a Program Manager roles to further support
    the SIG.
  - We are actively applying the `good-first-issue` and `help-wanted` labels to issues
    and are closely mentoring new contributors who pick up these issues.
- What programs do you participate in for new contributors?
  - The SIG Release Shadowing program
- Does the group have contributors from multiple companies/affiliations? Can end
  users/companies contribute in some way that they currently are not?
  - Yes, we have multiple avenues for contributing for both code and non-code
    projects alike.

#### Current initiatives and project health

- What are initiatives that should be highlighted, lauded, shout outs, that
  your group is proud of? Currently underway? What are some of the longer tail
  projects that your group is working on?
  - The introduced Program Manager role is our highlight to keep the SIG making
    continuously progress
  - Dedicated issue triage session
  - Building a North Star Vision Roadmap for long term planning
  - Continuously enhancing the release cycle timings and tooling around it
- Year to date KEP work:
  - [KEP-2572: Release Cadence][kep]
- What initiatives are you working on that aren't being tracked in KEPs?
  - Formalize supported release platforms:
    https://github.com/kubernetes/sig-release/issues/1337
  - Implement a Bill of Materials (BOM) for release artifacts
    https://github.com/kubernetes/release/issues/1837
  - Enhance Kubernetes binary artifact management
  - Simplify CVE process for release management (Secure)
    https://github.com/kubernetes/sig-release/issues/896
    https://github.com/kubernetes/release/issues/1354
- What areas and/or subprojects does the group need the most help with?
  - Nothing to mention right now, we're always looking for help in Release
    Engineering related topics. Beside that we have to assemble a Release Team
    each cycle which follows its own process.
- What metrics/community health stats does your group care about and/or measure?
  - We mainly stick to our project boards

[kep]: https://github.com/kubernetes/enhancements/issues/2572
