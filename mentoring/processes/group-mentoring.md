# Running a Contributor Ladder Growth Program 

There are differences in curriculum for reviewer and Chair and/or Tech Lead 
groups but the foundation of each are mostly the same unless otherwise noted 
below in the set up instructions. This difference is also reflected in our 
[community membership guidelines].

## Discovery of Needs

Cadence:  
ContribEx sets an issue in k/community requesting SIGs to participate preceding 
the next release by one month.

Rolling:  
SIGs can approach ContribEx at anytime to start a group; file an issue in 
kubernetes/community or talk with us in #sig-contribex.  

## Curriculum Building
While not a fully structured program, some light planning is needed to make sure
the curriculum being taught is relevant for the group. Here are our base 
suggestions:
  
*Reviewer*  
- Enhancements  
- Overview of Release Cycle  
- Development
  - Triage
  - Reviews
  - API Changes
  - Testing
- Code Freeze
- Cherry-Picks/Branch Management

*Chair*
- Enhancement Planning
- Meetings and Organization
- Responsibilities and governance overview
- Consensus-building
- Cross-project communication 

*Chair/TL combo*
Chair + Reviewer
- All of the above

One of the best strategies is to have everyone work on something together, 
Examples:
- reviewing guidelines for the SIG if there are none/updating guide
- CONTRIBUTING.md for the SIG if there isn't one/updating guide
- mid-size project
- a release cycle of features for the SIG, establishing better processes

## Setup and Outreach  
0. Need is established and SIG provides a mentor*
1. ContribEx kicks off two issues, sets a target start and end date, and selects
a coordinator
  - issue #1: outreach issue to collect folks who are interested, issue will 
  close when there are the max participants needed
  - issue #2 setting up infra - template here: [TODO link to example]
2. ContribEx or SIG member posts the issue to the following places:
  - SIG and k-dev mailing lists
  - #diversity, #kubernetes-contributors, #sig-[name], and other slack channels
3. Coordinator arranges a doodle for the kick off meeting and days/times that 
would be best for the mentor and group for slack standups and touch points. Once conscensus is reached, coordinator schedules it.
4. Mentor creates template that serves as a reference guide for the group. See
a sig-node example here: https://hackmd.io/8i8prErgSOamUmEZDDZPWA
5. Program begins at the kick off meeting
  - [TODO add kick off template]
6. Feedback is continuous as you stand up weekly; however, the ContribEx 
Coorindator will set a mid program check-in. In this check-in, members of the 
group will start their application process for their roles and any group members
 who have not had the time. 
  - [TODO add community membership template]
7. By this point, folks should know where they stand. At the end of the program,
 mentors and/or individuals will start the PR process in OWNERS files or 
 sigs.yaml for leadership roles for those who are ready  


[community membership guidelines]: community-membership.md 
