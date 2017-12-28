# SIG Governance and Check-in

**Identify the following before beginning your session. Do not move forward until these are decided / assigned:**

- **Session Topic: SIG Governance and check-in**
- **Topic Facilitator(s):** brendandburns@
- **Note-taker(s) (Collaborating on this doc)**:  jbeda@
- **Person responsible for converting to Markdown &amp; uploading to Github:** philips@

### Session Notes

- Core discussions about SIG leadership
    - Is the SIG lead role more like an on-call position or a team lead?
    - There are 2 parts of the role: facilitation and management of the owners files
    - authority comes from commits?

#### Conclusions

#### Key Takeaways / Analysis of Situation

Recommendations &amp; Decisions Moving Forward (High-Level)

Specific Action Items (&amp; Owners)

| **Action Item** | **Owner(s)** |
| --- | --- |
| SIG starter kit: Gather constructive feedback from SIGs; Create a summary of that feedback | Bob Wise |
| How should SIG lead/governance work? Figure out a lot of this stuff.  Everyone should share with the steering committee and get that ready for the full committee. | Steering Committee |
| Schedule SIG lead office hours to coordinate around PRs and getting stuff done. | Jason DuMars |
| SIG-architecture? | Steering Committee |
| SIG consolidation? | Steering Committee |

#### Purpose

-  Michelle
    -  What is your format?
    -  How is it going?
    -  Are there things that you are struggling with?
    -  Do we need to formalize about forming/running/governing a SIG?
-  Joe
    -  How do we pick SIG leads?  What type of leadership? What are the responsibilities?  Is this the same across the project or per SIG?

#### Discussion

-  Rob: Cross cutting sig concerns?  Venn diagram? Common concerns get lost.  SIG-cloud is similar.
-  What other types of horizontal/project?
    -  Ops, testing, etc.
-  Bob: Would it be useful to draw connection map and formalize communication channels
-  Eric: Staff and line position?  Line = line of business.  Staff = centralized responsibility.
    -  Q: Where would you put SIG-testing? A: staff?  Testing guidance for all the SIGs.
    -  Technical area (LoB) vs. staff (shared service and expertise)
-  Brendan -- there are technical staff groups vs. non-technical staff groups.  Example testing vs. contrib-x.
-  Joe: without some sort of automated mechanism that forces people to communicate it won&#39;t happen. Perhaps having the SIGs own code will force discussions at time of PR
-  If we have a forcing function there requires more planning up front.  Example is the rush to a release.
-  Scrum of scrums?  Have the leads get together and hash over that.
    -  Focus on visibility with accountability.
-  Luke: perhaps have sig leads do offline sync ups?  Bring in others as necessary.
    -  Vish: Sig leads may not know what is going on.
    -  Quinton: have identified people in the SIG that owns relationship with specific other SIGs.
-  Bob: Getting better definition of what the sig lead job is would help.
    -  Make SIG-PM be made up of SIG-leads
    -  Brendan: SIG lead should be a thankless position. Very little power.  Should be obligation.  Somewhat like an on-call rotation vs. TL.
-  Quinton: Is there a TL position?  Who is driving the overall technical direction of an area.  Separate from the lead position.
    -  Brendan: SIGs own directories.  Approvers in those directories are TLs.
    -  Bob: current model SIG leads are volunteer thing and so it is optimistic to expect too much.  Do we want to have elections.
    -  Brendan: worry about gaming with the openstack model wrt elections.
    -  Rob: in OpenStack -- SIGs were project groups.  The usage of SIG to mean project is confusing.
    -  Eric: If people are going to be SIG leads -- too much power for too much responsibility.  Should we solve by adding responsibility.
    -  Do sig leads have power now?  It is mostly influence?
      - **¦¦** Code vs. talk.
    -  Bob: You can say &quot;code talks&quot; but getting code in and earning that position can be super frustrating.  &quot;Code == influence&quot; just isn&#39;t true right now.
    -  Adding code is easy.  Reviewing code, curating tests, etc.  That is the hard thing.
-  What are the expectations to put on a SIG.  Keep track of issue response times?
-  Jason:         What is the job of SIGs?  How do we identify that and make sure it gets done?  There isn&#39;t that one thing to test against.
    -  Leadership follows from this mission.
    -  Eric: As a governing board do we care how a sig does it as long as it does it well?  I think the answer is we don&#39;t care.  Doubts around template.
    -  Luke: That seems fine -- but we need to share what works and what goes poorly.
    -  Brendan: others said cross sig stuff is hard.  Also poor communication across dependencies.
    -  What is success?  Depends on the SIG.  Lay out a plan and then accomplish it.
    -  SIG have obligations to release and own thing.  Example is tests for that thing. Need to define interfaces
-  Michelle: what are the best processes for governing and monitoring a SIG.  What are the best processes?
-  Aaron: Qs: Do SIG updates and recording help with communication.
-  Joe: What we discussed through bootstrap is that we should have a starter kit for the SIGs; we don&#39;t really have strong opinions on how it actually works in practice though.
-  Dalton: Clear lines of ownership for SIGs would be appreciated.  Would welcome more guidance. SIG on-prem struggles with what they are responsible for today
-  Challenge for horizontal SIGs include membership (especially technical) and influence.
-  Brendan: focus on fewer larger SIGs.  Not enough people with smaller SIGs.  Consolidation may make things more efficient.
-  Aaron: SIG updates at community seem to help tell what the SIGs are focusing on but it doesn&#39;t help on a week to week or release wide process.
    -  Does SIG-PM coordinate a snapshot of what is going on.  But it isn&#39;t enough communication.
    -  Ihor: Almost meeting have reports from SIGs.
-  How do we push people to write code in the right direction.  What tools do we have to guide things.
    -  Hope from the steering committee is that if SIGs aren&#39;t doing something they should be doing there is an escalation.  Example is SIG-testing is empowered to stop queue until things change.
-  Joe: Eng reviews/SIG-architecture
    -  Advise to who and where to get sign off.
    -  Bob: alleviate frustrations that we had

Lots of discussions on how to drive communication and when things get escalated. Help SIGs coordinate between themselves.

Up to steering committee to help define what the requirements are on a SIG.  Still unsure what that&#39;ll look like.

Jason: rename SIG-lead to SIG-facilitator

Good Practices

Kind of like a starter kit for SIGs?  Needed and recommended?

-  Notes and videos help
-  Focusing bi-weekly for PRs vs. design discussion.
