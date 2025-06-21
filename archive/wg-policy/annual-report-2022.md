# 2022 Annual Report: WG Policy

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - CR for PolicyReport being used more widely in other projects and by end users
   - 2 whitepapers released
   - 2 KubeCon talks NA + EU

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We are discussing a KEP for the PolicyReport CR but still pending
   - Feedback from some of the sig leadership recommend NOT doing a KEP but just hosting the code in sig-auth or sig-security namespace
   - Outside of that there has been a lot of community interest, and workgroup effort spent, on control mapping
     and control-as-code implementation, eg OSCAL, that might be better served moved into its own workgroup or a 
     sandbox project

## Project health

1. What's the current roadmap until completion of the working group?

   - Once the CR KEP is submitted or the sig decides yea or nay, we anticipate winding down the WG unless the community asks for new prototypes
   - There seems limited/no interest in a corresponding CR for policy inputs/profiles
   - One option is that many of the attendees are interested in compliance, so maybe a sig-security compliance WG is a follow on 
   - Also several of the concrete policy implementations can be carried over to SLEDGEHammer (which will be submitting a Sandbox application)

2. Does the group have contributors from multiple companies/affiliations?
   - Yes (RedHat, IBM. Kyverno, Google, Fairwinds, Defense Unicorns, others)

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?
   - Maintaining the PolicyReport API code
   - Building out more PolicyReport API client code and examples
   - Contributing more concrete policy library content (SLEDGEHammer will be committed to this)
   - There is considerable interest in continuing the governance and assessment and lifecycle of policy and controls,
     however as these necessarily cross boundaries, it seems like something that should either be re-homed to sig-security,
     and/or hosted in a CNCF-level workgroup and/or moved into a relevant sandbox CNCF project

## Membership

- Primary slack channel member count: 360
- Primary mailing list member count: 139
- Primary meeting attendee count (estimated, if needed): ~8
- Primary meeting participant count (estimated, if needed): ~6

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [X] Updates provided to sponsoring SIGs in 2022
      - [sig-auth](https://git.k8s.io/community/sig-auth/)
        - TODO: JIM: links to email, meeting notes, slides, or recordings, etc

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-policy/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
