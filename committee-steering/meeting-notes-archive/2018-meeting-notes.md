## Dec 5, 2018

* **Bosun:** spiffxp
* **Note Taker:** philips
* **Video:** [https://youtu.be/jV9s1bt45rE](https://youtu.be/jV9s1bt45rE)
* **Attendees:**
    * Aaron Crickenberger, Brian Grant, Phil Wittrock (Google)
    * Derek Carr, Clayton Coleman, Brandon Philips (Red Hat)
    * Michelle Noorali, Brendan Burns (Microsoft)
* **Topics:**
    * Contributor summit discussion (spiffxp)
    * Any AI followups?
        * We were going to try to push on charters, let’s see where we are on that
        * Push on all charters reviewed - action items about pinging these
        * AI About CLA
        * AI Infra wg charter - make it “less sig like” - who created and owned subprojects (wg will find SIGs)
            * Will have a couple SC folks for approval instead of a full re-review
            * Brian
            * Brandon can if reviewable by Wed of next week
    * Going through the board
        * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
        * Contributing Structure
            * AI Phil - follow up with Nikita
        * Charters
            * Contributor Experience
                * Need to handoff to another SC member
                * AI to Phil
            * PM
            * Big Data
                * Don’t have consensus as to if this should be a SIG vs something else
                * OK with being a SIG for now, can sort out whether this is the correct organizational structure next year
                * Not worth changing just before kubecon and end of year
                * Issue: Just because we have charters, doesn’t mean the state if finalized or the SIG will continue to exist as such
                * Will leave charter with ‘/hold’ to avoid confusion w.r.t. charter merge status
            * Docs
                * Joe + Brian on the hook to be reviewers
                * Seems lite
                * Aaron can help review
            * Cloud Charters
                * Azure - before had template
                * GCP - considering rolling into SIG cloud provider
        * CNCF dev rep
            * Continue to be Michelle
            * Would be good to have more attendance at meeting - will bring up at SC meeting more with more advance - has been done before
            * In future will go through agenda for GB at SC meeting if it is available (usually isn’t)
        * K8s Infra to CNCF
            * Need to add fields to working groups to reference the SIGs that are collaborating
        * WGs
            * Should be ephemeral, but some WG that seem to be permanent
            * Should revisit these - should they be sub projects?
        * Dependencies
            * Need to be able to do a better job
            * Have been vetting based off Licensing
            * Need to vet based of confidence that the libraries don’t contain any vulnerabilities
            * Dependencies keep getting reintroduced
            * Need to review items on issue
            * Need to make sure we are able to maintain a good state
            * Have good people as gatekeepers, but impossible task for humans to audit all libraries
            * _Need tooling to do CVE scanning_
            * Mostly driving changes that have been done in k/k to also be done in other repositories
            * Can turn on github scanner that will give us a badge to put on repos -  this is close
            * Talked with legal rep from CNCF, he promised to scan all repos 1 more time at the start of the year
            * Concerned that this is an attack vector
            * If there is a good service, CNCF would pay for it
            * RedHat might have something similar that could do this
                * RH VS code plugin might have a CVE to dependency translator
                * A good place to start
            * Not focussed on subtle bugs - focussed on injection attacks
            * Go dependencies makes this a good target for attack
            * NodeJS
                * Security Audit performed
            * High priority item
            * No one has a DB of go package CVEs
            * Other languages have better support for metadata for CVEs
            * Brendan to put together a proposal for how to deal with this problem (bring to f2f at kubecon)
            * [Go CVE’s](https://www.cvedetails.com/vulnerability-list/vendor_id-14185/product_id-29205/Golang-GO.html)
        * SIG Release Charter Observations
            * 2 things:
                * How and when artifacts are released
                * Declare a release for SC
            * Lot of things in the charter that make SC the escalation point for everything in the project
            * Who do we need involved in discussion as we move things out of tree - SIG Arc + SIG Release?  Hopefully don’t need to bring SC into the conversation.
            * SIGs should be able to collaborate if there is a policy change - Arc, Release, Docs, Testing, etc
            * Need to do some clean up anyway
                * Kubeup has been deprecated for 1+ years, but still in the release bundle and used for testing
        * Contributor Summit planning
            * Need to talk about what we have accomplished
            * Worthwhile to talk about what we are planning to do for the coming year
            * Security thing is top of mind
            * Lot of governance bits are not completely done, but there is lite at the end of the tunnel
            * Find a way to empower as many people as possible to move forward on these issues
                * Move away from “SC” will do it
                * SC plates are already super full
                * Take advantage of role board - populate candidates
            * Rationalization to make scope and size and organization of projects manageable
                * Flat organizational structure makes it hard for new contributors to navigate
            * What do we need to do to keep Kubernetes boring?
                * Impart a desire to have the project focussed on health - de-flaking tests, putting gates in place, define alpha -> beta -> stable
            * Google doc to discuss what we want to talk about?
        * Smaller conferences instead of massive Kubecon
            * Starting to see conferences focussed on ecosystem components
            * More local events for people who can’t travel to big conferences
            * At the conference there will be a session about making the conference better
                * Thursday afternoon
            * Will have Kubernetes day in Bangalore
            * Super-Meetups style events could be popular
        * Contributor summit
            * Has been invite only in past to keep it a manageable in size
            * Have used dev stats numbers in the past
            * Seen a lot of people submitting random-typo PRs - trying to game the stats?
            * One thing fixed this year - all of the sessions will be recorded (yay - thanks Paris)
                * Some of the most valuable content - will be able to link to this
            * Raw numbers to quantify contributor impact - hasn’t worked well
            * In past was invite only - 100 folks
            * Recently was 400 - seem like just a conference intro session - shouldn’t have been on the main page at all?
            * Need to be careful of “who you know network” - going back to pure invite only not great either
            * Being more clear about what the contributor summit is about would allow folks to self select
            * Only want people who will be serious contributors to be coming
            * How much of an issue is this
                * It is definitely happening, but is it impacting decisions we are making - e.g. if we pull the top 100 are we going to see folks who have gotten there through gaming the system
        * Kubernetes Steering Committee Thursday dinner/Friday meeting planning
            * Please add to the [issue](https://github.com/kubernetes/steering/issues/87) if you have items to discuss
        * Ratify Code of Conduct Committee Charter
            * Private mailing list has some stuff about this - this is important to review

## Nov 21, 2018

Canceled due to Thanksgiving

## Nov 7, 2018

* **Bosun:** Brendan Burns
* **Note Taker:** Aaron Crickenberger
* **Video:** [https://youtu.be/I9qRE9DM1Tw](https://youtu.be/I9qRE9DM1Tw)
* **Attendees:**
    * Aaron Crickenberger, Brian Grant, Phil Wittrock, Sarah Novotny (Google)
    * Dims (Huawei)
    * Joe Beda, Timothy St Clair (Heptio)
    * Brendan Burns (Microsoft)
* **Topics:**
    * AI Followup:
        * AI Aaron: Will send out some next steps off the call to steering (re: subprojects)
        * AI Aaron: Move the SIG role stuff into the other governance doc
        * AI Brendan Burns: Ping SIG Windows
            * [https://github.com/kubernetes/community/pull/2875](https://github.com/kubernetes/community/pull/2875) - went in without SC reviewers
        * AI Brendan Burns: Poke SIG Azure
            * Charters Merged but needs another person to walk back and review
        * AI Joe: Ping SIG AWS [https://github.com/kubernetes/community/pull/2733](https://github.com/kubernetes/community/pull/2733)
            * Close, but pending lazy consensus.
        * AI Aaron: Ping SIG PM
            * Have not done: will do this week
        * AI Dims: SIG Network
            * PR is ready - [https://github.com/kubernetes/community/pull/2744](https://github.com/kubernetes/community/pull/2744)
        * AI Dims: Chase down SIG Cloudprovider and SIG OpenStack merging
            * Need to follow the new process to consolidating the OpenStack SIG
        * AI Dims: understand how many PRs are blocked on CLA signing
            * [15 open PRs](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3A%22cncf-cla%3A+no%22)
        * AI Dims: SIG Architecture meeting discussion on KEP for OpenStack extraction
            * Done
        * AI Dims/Aaron Review the Infra WG PR and see if we can make it feel less SIG like
            * Not Done yet.
    * [Tim] QQ - Do we want to approve WGs?  Yes/No **[YES, but we reserve the right to delegate]**
    * Going through the board
        * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
    * Philips: Recommendations for CNCF TOC nominations?
        * Nominations close Nov 26th
        * We could suggest people on steering-private@
        * If we’ve got objections we could decide not to nominate (effective lazy consensus)
        * What criteria are we looking for?
        * Brian can send notes about what TOC actually does
        * It is a somewhat substantial time commitment, make sure we’re nominating people who can handle that
        * Keep well qualified diverse candidates in mind
        * Think broadly, beyond just Kubernetes, the TOC is a different kind of work than Kubernetes
    * [Charters](https://github.com/kubernetes/steering/issues/31): [SIG Apps has a charter](https://github.com/kubernetes/community/pull/2881/files) up that Brian wants to review
        * Done! The charter LGTM.
    * The KubeCons are eating up our time
        * Upcoming meetings are going to be interleaved with them
        * What kind of prep do we want to do with them
        * Contributor summit stuff looks reasonable for Seattle
        * We are going to cancel the next meeting because it is so close to thanksgiving
        * But OK we’ll keep the Dec 5th meeting
        * Maybe a good time to review our backlog and consider whether it’s worthwhile?
        * AI: Joe to draft a document to describe the process by which we will avoid undue process

## Oct 24, 2018

* **Bosun:** Aaron Crickenberger
* **Note Taker:** Brandon
* **Video:** [https://youtu.be/PbLQhLP9F_4](https://youtu.be/PbLQhLP9F_4)
* **Attendees:**
    * Tim St. Clair, Joe Beda - Heptio
    * Aaron Crickenberger, Brian Grant - Google
    * Brandon Philips, Derek Carr, Clayton Coleman, Red Hat
    * Dims - Huawei
* **Topics:**
    * Going through the board
        * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
        * SIG Lifecycle
            * [https://github.com/kubernetes/community/issues/2029](https://github.com/kubernetes/community/issues/2029)
            * Needs to be a lifecycle only doc
            * AI Aaron: Move the SIG role stuff into the other governance doc
        * Sub projects
            * [https://github.com/kubernetes/steering/issues/36](https://github.com/kubernetes/steering/issues/36)
            * Brian Grant: we need to have fewer OWNERs files
            * Aaron: we could build some automation around the OWNERs file and building READMEs for them
            * AI Aaron: Will send out some next steps off the call to steering
        * Charters
            * [https://github.com/kubernetes/steering/issues/31](https://github.com/kubernetes/steering/issues/31)
            * Notifying various SIGs
            * Has anyone notified SIGs big data, SIG windows
            * Brian Grant SIG Big Data should probably be a working group
            * AI Brendan Burns: Ping SIG Windows
            * AI Brendan Burns: Poke SIG Azure
            * AI Joe: Ping SIG AWS [https://github.com/kubernetes/community/pull/2733](https://github.com/kubernetes/community/pull/2733)
            * AI Aaron: Ping SIG pm
            * AI Dims: Chase down SIG Cloudprovider and SIG OpenStack merging
        * DCO v. CLA
            * [https://github.com/kubernetes/steering/issues/74](https://github.com/kubernetes/steering/issues/74)
            * Brian Grant: there is some disagreement on the tradeoffs of upfront complexity vs ongoing complexity
            * Discussion around occasional contributors
            * Punting this discussion to another time
            * AI Dims: understand how many PRs are blocked on CLA signing
    * Sig cloud and parentless work around provider.
        * Discussion around carrots and sticks
        * Brian Grant: SIG Architecture should be the place to have the next escalation
        * AI Dims: SIG Architecture meeting discussion on KEP for OpenStack extraction
    * Removal of out of date OWNERS
        * Dims trying to identify who is and isn’t active today
        * Aaron probably this should be part of SIG Contrib Exp putting a policy in place
        * Brian Grant we created the contrib ladder to help with this but we never worked to ensure people are pulling people up the ladder. It wasn’t well socialized to owners and approvers.
        * Joe Beda is underfunded as far as tooling is concerned, do we need to involve the CNCF
        * Dims should feel empowered to go get funding from CNCF to make this tooling :)
        * AI Dims: Take this to Contrib Exp?
    * [https://github.com/kubernetes/community/pull/2830](https://github.com/kubernetes/community/pull/2830) - wg-k8s-infra majority vote needed to proceed
        * Discussion around whether it should be a SIG or a working group
        * Aaron pointed out that it is unclear whether this needs to be long lived vs getting SIGs to take ownership of things they rely on. E.g. GCP stuff primiarly SIG Release
        * AI Dims/Aaron Review the Infra WG PR and see if we can make it feel less SIG like
        * Do we need something separate governance of important secrets?

## Oct 10, 2018

* **Bosun:**
* **Note Taker:** Sarah Novotny
* **Video:** [https://youtu.be/mdbJs_ApdoU](https://youtu.be/mdbJs_ApdoU)
* **Attendees:**
    * Derek Carr, Clayton Coleman - Red Hat
    * Tim St. Clair - Heptio
    * Sarah Novotny - Google
    * Brendan Burns, Michelle Noorali - Microsoft
    * Dims - Huawei
* **Topics:**
    * Election has ended. Are there next steps for the steering committee?
        * Issues filed to be added to the project board.
            * Subgroup to address feedback
                * Those who didn’t join today are that subgroup.
            * Ask Jaice to add the new issues to the project board.
                * https://github.com/kubernetes/steering/issues/81
            * Do analysis of voter turnout (compare to devstats threshold?)
            * Mapping for Github email ID?
            * Need a goal for turnout in elections -- want to be sure we are thought of as the legit leaders.   Does this election represent this?
        * How do we onboard new members? (michelle)
            * Michelle is doing this right now with Dims.
            * We should have a more formal process.
                * Introduction at first meeting
                * Add to all the mailing lists
                * Mentor
                * Process for onboarding needs plan for next election.  More people who will leave the committee with fewer joining.
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
        * Blocked election issue -- not blocked
        * SIG/WG sunset -- AI Michelle to ping Paris/Jorge
        * WG creation --  AI Brendan to revisit.
        * Steering Committee Liason
            * Meta problem with charters not in place.
            * AI Aaron to look at issue
            * Carrot vs stick convo re: charter
                * Dims suggest that we don’t serve (i.e. create new repo etc) SIGs who don’t have charters.
                * Freezing zoom account?
                * Blocking merge privileges?
                * We need to do this. . but, the penalty is rational -- and, need to manage enforcement.
                * Reach out to each SIG.  before a heavy hand.
                * Tie to a release cadence.
                * GA a feature without a SIG  charter?
                * Add to milestone?
            * EVERYONE -- go through [the spreadsheet](https://docs.google.com/spreadsheets/d/1i_8oM1KlmXo90zIIgNTdOg-M6xesTggFLnTkgmscMlI/edit#gid=0) by EOW. and nag/facilitate charters.
            * Contributing.md
            * New Issue -- DCO vs CLA -- try to bring Max Sills to meeting 10/24.  Find a CNCF lawyer… get presentations on positions pros/cons.
                * [Discussion in Apache](https://mail-archives.apache.org/mod_mbox/www-legal-discuss/201810.mbox/browser)
    * In person meeting in December -- Steering committee at Kubecon
        * Dinner Thursday
        * Friday meeting.

## Sep 26, 2018

* **Bosun:** Michelle Noorali
* **Note Taker:**  Brandon
* **Video:** [https://youtu.be/OgaZZytEphs](https://youtu.be/OgaZZytEphs)
* **Attendees:**
    * Aaron Crickenberger - Google
    * Brandon Philips - red hat
    * Michelle, Brendan - Microsoft
    * Philip, Brian - Google
    * Joe, Tim - Heptio
    * Quinton - Huawei
* **Topics:**
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
    * [https://github.com/kubernetes/community/issues/2715](https://github.com/kubernetes/community/issues/2715) - what do we call the effort to move k8s-infra to CNCF?
    * CoCC Knowledge Transfer Session happening Monday. Anyone from Steering Committee who wants to join can contact Michelle for an invite.
    * [Documentation requirement](https://github.com/kubernetes/steering/issues/53)
        * AI: Michelle will follow up with Jaice to see if steering backup is needed and let them know about their authority
        * Brian Grant context: whenever someone in the project puts a gate on changes there is often push back from the community where people walk around the gate, reject the gate, or otherwise ignore it
        * Quinton: hopefully we can delegate to SIG docs
        * Brian Grant: yes, that is the idea
        * Aaron: it would be helpful to have links about what policies SIG docs is trying to implement
        * Part of the problem is that repetition is needed to ensure everyone is even aware of the policies
        * Aaron: part of this should probably be ensuring that all SIGs have a Liason from the steering committee. But, we are still sort of blocked on this until after the elections.
    * [Elections](https://github.com/kubernetes/steering/issues/63)
        * They are in progress
        * AI: Aaron see if swapping meet contributors timing to after the elections is possible with Paris
        * AI: Quinton talk with Jorge and Paris to ensure we are getting a reasonable number of votes
            * The steering commitee will make the announcement, election commitee can prepare a draft
        * AI: Aaron figure out how to ensure steering committee members running don’t see detailed election results
    * [Document working group formations](https://github.com/kubernetes/steering/issues/27)
        * BB [walking through his comments](https://github.com/kubernetes/community/pull/2702)
        * BG: I feel like many things that are working groups should be subprojects
        * BB: mostly concerned about endorsement of SIGs
        * Discussion around code getting merged into Kubernetes via SIG ownership vs. a working group developing code to prove out an idea concept
        * AI: Phil clarify code ownership in Kubernetes v. building code in a working group to prove out an idea to build consensus - add suggest where to put that code (associated repo)
        * AI: Phil consider striking the SIG approval lines BB commented on
    * Discussion about Infra Team
        * Aaron: is it a SIG, Team, Committee, WG, etc
        * AI: Aaron propose a new working group to push the infra move forward

## Sep 12, 2018

* **Bosun:** Michelle
* **Note Taker:** Tim St. Clair
* **Video:** [https://youtu.be/2FyTzm5REUU](https://youtu.be/2FyTzm5REUU)
* **Attendees:**
    * Derek, Brendan, Michelle, Aaron, Joe, Clayton, Sarah, Phil, Tim
* **Topics:**
    * [AI]: Need to review the CoC charter on a fast track
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
        * Sunsetting SIGs
            * [AI]: Right now we are pending on PR from contribex on policy and procedures for “How” we sunset a SIG.
        * Formalize Subprojects
            * Still in flight, but still pending on better documentation.
            * Nothing is currently blocking
        * Working Group Formalization
            * [AI:Phil] Still pending on converting doc to markdown to file the PRs.
                * [Link](https://docs.google.com/document/d/1AlI89KijzO9_KAqUX_pbumgld1LN0tte2FMdjLw83wY/edit)
                * [PR Link](https://github.com/kubernetes/community/pull/2702)
        * SIG Charters
            * Edicate
            * There are currently a small subset of charters that are pending.

## August 29, 2018

Canceled due to lack of quorum

## August 15, 2018

* **Bosun:** Aaron Crickenberger
* **Note Taker:** Brian Grant
* **Video:** [https://youtu.be/PhXlFhbwqWg](https://youtu.be/PhXlFhbwqWg)
* **Attendees: **
    * Michelle Noorali
    * Aaron Crickenberger, Sarah Novotny, Brian Grant
    * Joe Beda
    * Quinton Hoole
    * Brandon Philips
* **Topics:**
    * Private Code of Conduct discussion
    * Elections
        * Member of standing criteria ([https://github.com/kubernetes/community/pull/2532](https://github.com/kubernetes/community/pull/2532)) and other election guide details were merged without SC review. We want to move the authoritative policies to the Steering Committee org.
    * Charters
        * What to do about areas not actively being pursued by SIGs or that doesn’t fit with K8s current scope
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
    * [aaron] Chasing after repo ownership and consistent automation
        * I’m not sure which steering issue(s) this falls under, more of a heads up
        * Would like all managed github orgs to be using our automation
        * I’ve been getting pushy on repo/subproject ownership (in sigs.yaml, and with OWNERS files)
        * [https://github.com/kubernetes/test-infra/issues/6227](https://github.com/kubernetes/test-infra/issues/6227) - tracks most, but probably missing some (eg: /retest and /lifecycle commenter jobs that run as @fejta-bot)

## August 1, 2018

* **Bosun:** Aaron Crickenberger
* **Note Taker:** Brian Grant
* **Video:** [https://youtu.be/_jukicl1yK4](https://youtu.be/_jukicl1yK4)
* **Attendees: **
    * Tim St. Clair, Joe Beda - Heptio
    * Aaron Crickenberger, Sarah Novotny, Brian Grant - Google
    * Derek Carr - Red Hat
    * Quinton Hoole - Huawei
    * Brendan Burns - Microsoft
* **Topics:**
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
    * CoCC vote completed. Electees will be notified by Michelle.
        * Aim to announce at community meeting next Thursday
    * [quinton] Steering committee election is going ahead: [https://github.com/kubernetes/steering/issues/63](https://github.com/kubernetes/steering/issues/63)
        * Please object this week if you have any concerns
    * Working group documentation: [https://github.com/kubernetes/steering/issues/27](https://github.com/kubernetes/steering/issues/27)
        * Brendan will convert doc to a PR
    * SC liaiasons: [https://github.com/kubernetes/steering/issues/64](https://github.com/kubernetes/steering/issues/64)
        * Underway
        * Aaron will solicit folks to staff gaps as identified
    * Charters: [https://github.com/kubernetes/steering/issues/31](https://github.com/kubernetes/steering/issues/31)
        * Concerns about owners of subprojects falling under SIGs that don’t attend SIG meetings
        * Some SIG should own removing the cluster directory
            * Brian will follow up regarding whether anyone is working on this
        * We need to move charters forward and get them merged
            * SIG Auth and Service Catalog merged
            * SIG Node is ready
            * Liaisons should attend SIG meetings to discuss
            * We need a replacement for Phil on SIG CLI
            * We should try to get the charters into mergeable states, and expect to iterate on them
        * Checkpoint prior to next week’s community meeting
        * We should provide an update at the community meeting next week and emphasize that it will be an iterative process
    * Meeting logistical issues
        * Office hours
        * Non-SC invitees, participants, delegates
            * Broader participation would help get more stuff done and grow the next set of leaders
            * Concern about potential to influence the direction if just appointed
            * We will make a broader call for participation
            * The Github Administration proposal should be called out as a good example of how to productively engage -- clear, well bounded, fully formed proposal for approval

## July 18, 2018

* **Bosun:** Aaron Crickenberger
* **Note Taker:** Tim St. Clair
* **Video:** [https://youtu.be/I6BwkOA9dn4](https://youtu.be/I6BwkOA9dn4)
* **Attendees:**
    * Brandon Philips, Derek Carr - Red Hat
    * Phillip Wittrock, Aaron Crickenberger - Google
    * Joe Beda, Tim St. Clair - Heptio
    * Brendan, Michelle - MSFT
    * Quinton Hoole - Huawei
* **Topics:** tbd
    * [https://github.com/kubernetes/steering/projects/1](https://github.com/kubernetes/steering/projects/1)
        * New Contributor Site - pending a vote on creating a repo
        * Documentation requirements - should be punted back to docs/arch
            * AI: Aaron will ask Jaice to reach out.
        * [phillip, derek, tim] - TL;DR on charter template + feedback.
            * [https://github.com/kubernetes/steering/issues/31](https://github.com/kubernetes/steering/issues/31)
            * Template has been updated and **node & auth** are the canonical examples to reference as we start to give feedback on charters.
            * [Michelle] Has also created a separate doc [https://github.com/kubernetes/steering/issues/62](https://github.com/kubernetes/steering/issues/62)
                * How do we evaluate the charter(s).
                * Aaron - I’d really like a checklist on evaluation.
            * [Aaron] Should we expect iterations in the charters
                * [Brian] If we change the scope/leads then steering needs to evaluate
                * [Brian] We could version it w/ policy.
                * [Michelle] We can just open an issue and adjust.
            * [Aaron & Michelle] Sig-cloud-provider and separate providers will need to be treated with care.
            * [Quinton] Should we provide a longer term liason per-sig who does periodic reviews.
                * [Aaron] AI: Then we should take the spreadsheet and make it more formal.
            * [Aaron] How should we proceed
                * [Michelle] Deal with the current PRs and loop back and discuss during the next meeting.
                * … details on certain “in flight” PRs …
                * [Tim] just comment on the PRs, and email if there are issues.
                * [Derek] will update the node charter as the better-better example that we should be referencing.
                * [Michelle & Brendan ] We should send out an update on the PRs and email(s).
        * Document how working groups should be formed, [Jaice’s Doc ](https://docs.google.com/document/d/1AlI89KijzO9_KAqUX_pbumgld1LN0tte2FMdjLw83wY/edit)
            * [Aaron] what do we need to move this forwards?
                * [Brian] look at the proposal, summarize a TL;DR ->PR and get it merged.
                * AI: Michelle, Aaron, Joe, and Derek will work on this.
        * License Scanning
            * [tstclair] Isn’t this sig-release’s ownership?
            * [Brian] This overlaps with testing and…
        * [Michelle] Code of Conduct
            * Opened a PR which contains a [outline](https://github.com/kubernetes/community/pull/2384) on committee bootstrapping, etc.
            * Deadline to vote is tomorrow.
            * Please **GO VOTE** so the committee can be formed.
            * There is still a lot of work todo on policies, there should be PRs coming.
            * There are more candidates that came in and there are some issues currently with any one company monopolizing the committee.
        * [Quinton] We should start putting together the vote details.
            * [Aaron] AI: please log an issue with the details.

## June 20, 2018

* **Bosun:** Tim St. Clair
* **Note Taker:** Joe Beda
* **Video:** [https://youtu.be/UKJYy9Oiuv0](https://youtu.be/UKJYy9Oiuv0)
* **Attendees:**
    * Tim, Joe - Heptio
    * Phil, Tim, Sarah - Google
    * Derek, Clayton- Red Hat
* **Topics: **
    * Getting things done as a steering committee
        * Joe -- need to find ways to get things done async
        * Tim -- we had a sprint thing with a backlog.  Jaice was going to help groom that but it seems to have stalled.
        * Sarah -- Example: I haven’t had time to get the steering committee office hours.
    * [TimSt] Charter review backlog and state.
        * [TimSt] We were going to find commonality and update the template.  Has anyone reviewed these?
        * [TimH] I’ve read several charters but haven’t found commonality.
        * [Derek] Last time didn’t we decide to start with 3?  Perhaps do that next meeting? Node can be one
        * [TimSt] Start with small group together? (yes) Who wants to sign up: TimH, Phil, Derek, TimS
            * By next meeting come armed with results/action items around commonality and modifications to template.
            * Phil is going to schedule and send out charters to review.
            * Then divvy out backlog of outstanding charters.
    * [TimSt] Documentation and helping define overall guidelines and policies
        * TimSt has done a lot of work around sig-docs this cycle. Lots of stale docs. Docs folks don’t know what to do with them. Lead to confusion.
        * Owners and feature policy (project wide)
            * There is a lot of rot and no one owns some of them (Cloud-provider)
            * Alpha features need to have policy.
            * Who do docs folks poke if there is stuff that is out of lifecycle.
        * [Joe] I’m okay with stuff without owners being dropped and moved to an “attic” someplace.
        * [Sarah] This should be delegated to SIG-Docs
        * 2 goals:
            * Write and update docs
            * Kick out docs that are unmaintained
        * What are the sticks?
            * Remove docs for things that aren’t maintained
            * Is this part of the charter? Is the chair on the hook?
            * Don’t let things go in without docs
        * [TimH] the fact we allow code checkins without docs is a problem. Put docs folks in the loop without a PR?
        * [Joe] Perhaps gate docs on promotion of APIs to beta and feature gates.
        * [Derek] Who are the docs folks?  Are there these people?
            * [TimH] Perhaps let Docs inject into the process
        * [Phil] Perhaps the KEP is the right place to have this happen.
            * [TimSt] Sometimes code drifts after KEP. That may be too early.
        * [TimH] Just APIs? [TimSt] Feature gates?
        * [Phil] What level of docs? What is the quality bar? Reference vs. tutorial vs task?
        * [TimH] Have SIG-docs define the rubrick -- but they may not have bandwidth.
        * [TimH] Push SIG-docs to give us a proposal.
        * [TimH] What can we do to make sure docs don’t rot?
            * TimSt: part of the charter?
            * TimH: won’t force people
            * Phil: Get a signal.  Dashboard?  Based on every doc they own.
            * TimSt: part of release process w/ automation.  Just poking the chairs would go a long way
            * Phil: clearly defined ownership?
            * TimSt: they are looking for an owning sig and a release owner?
            * TimH: are they expected to review every doc every cycle?
            * TimSt: that is what commercial projects do.
            * Phil: we need to get a place where owners can’t hold up a release.
        * Joe: What about ownership for docs -- how do we know what docs are owned by which SIG?
        * Joe: Volunteer army -- what sticks do we have?
            * TimH: We have the power to say no to PRs. Example: Docker has docs folks sign off on every PR.
        * 4 options:
            * KEP time
            * Every check in
            * API and feature gate promotion
            * At release time
        * Derek: can we deputize people inside SIGs to act on behalf of SIG-docs to call if docs are required or not?
            * TimSt: responsibility of chair?
            * TimH: people are lazy.  Could happen now. But not in your face so gets forgotten (I’m guilty too).  Maybe not every PR -- perhaps only things that hit APIs/feature gates/flags/configs?
            * Derek: perhaps a conscious action to say docs not required.  How happy are with release notes?
                * Joe: better than nothing?
                * TimSt: problem where people don’t respond when pulling in release notes.
            * Phil: wouldn’t cost us much to put in another label and a docs LGTM. Even if that LGTM doesn’t required.
            * TimSt: Something along the lines of release notes required.
            * TimH: we need a cultural change that you can promise to do docs next week
            * TimH: fan of release note like process
            * Phil: we need guidance and motivation
            * TimSt: I do like automation.  Key reviewers will see it.
        * TimH: different types of LGTMs around code, docs, relnote
            * TimH: Docker doesn’t even look at code until docs are written
            * Joe: perhaps a generic set of LGTMs that need more over time
        * Action: TimSt -- point docs folks at this recording. TimH says “they are empowered to ask for more.”
    * [Quinton] Steering committee elections:
        * Milestones (working backwards):
            * Election results announced - early Oct 2018? (‘t’)
            * Open for votes - t-2 weeks?
            * Announce who may vote (and allow exception applications) - t-4 weeks?
            * Candidate bio distribution deadline - t-4 weeks?
            * Candidate solicitation - t-6 weeks?
            * Sounds like we need to kickstart the process ~early Aug (t-8 weeks).

## June 6, 2018

* **Bosun:** Joe Beda
* **Note Taker:** Joe Beda
* **Video:** [https://youtu.be/peqx1GWW7qQ](https://youtu.be/peqx1GWW7qQ)
* **Attendees:**
    * Joe, Tim - Heptio
    * Brendan Burns - Microsoft
    * Sarah, Tim, Brian, Phil - Google
    * Clayton - Red Hat
* **Topics:**
    * CoCC nominations
        * Sarah -- wait for end date. Nominations: 4-5 people -- may be enough.  Don’t need to be too big a committee with clear leader.
        * How open are nominations?
            * Keep it scoped wrt membership
            * Hand pick from SC
            * What about sharing nomination form more widely?
    * Charters
        * What is the plan now?
        * Look at samples -- make sure template is working. Make a round of feedback, update template, go wider.
        * Have we delegated to reviewers to make the call.
        * New SIGs being proposed -- discussions on if/how and the charter.  Specifically SIG-CloudProvider
        * Some folks are unsure about losing the status as sig lead/chair.
        * Separable issue: Approval for new SIGs?  SIG-SDK/PDK…?
            * Lots of technical discussions but we decided we didn’t have the right people in the room and pulled back.
            * Steering committee office hour…
        * We need to be clear about what powers are dedicated to SIG-Architecture (etc). Need a list of explicit delegations.
            * Brian: that will end up in SIG charters
            * Tim/Brendan: is that sufficient?  We should be explicit vs. having stuff enumerated with charter.
            * Brian: Agree
    * QQ on release cadence of subprojects
        * SIG-architecture owns policy on when and who releases and where they live.
        * SIG-release owns the mechanics of how and where we host/sign.
            * SIG-release is underfunded to take on some of this. But perhaps we can motivate this.
    * CoreDNS and third party components
        * Quick update -- don’t fork into our space. Took their container images (to docker hub) and mirror into our GCR buckets. Checksums are the same.

## May 23, 2018

* **Bosun: Michelle Noorali**
* **Note Taker: Tim St. Clair**
* **Video:** [https://youtu.be/pMaRzcRI9HU](https://youtu.be/pMaRzcRI9HU)
* **Attendees:**
    * Derek, Clayton - Red Hat
    * Tim, Joe - Heptio
    * Michelle, Brendan - MSFT
    * Brian, Sarah, Phil - Google
    * Quinton - Huawei
* **Topics**
    * Code of Conduct Committee
        * Please nominate someone for the Code of Conduct Committee using the form sent to you on steering-private
        * Jaice has been working on docs and has outlined some details on “abuse of power”.
        * Q/A/Notes
            * Q: Can steering committee members be on the Code of Conduct Committee
                * A: Not a requirement, but:
                    * may want an upper bound,
                    * or possibly an advisory position
                    * … still not fully defined
            * Note: We should nominate Jaice
            * Q: There are some questions about organizational vs. contributor violation
                * A: The violations for leads is an individual contributor
            * Note: Follow ups should refer back to the documents.
            * Note: Folks would like more specific examples where possible.
            * Q: Do we agree with the framework that Michelle has started in the creation of the CoCC and should we proceed?
                * A:  YES!  But we would like it to remain Kubernetes specific, and may need to revisit some of the logistics.  Please see (TODO:link)
                * Note: Contention on escalation towards CNCF, and division of responsibilities.
            * Note: It may be beneficial for the CoCC to publish stats.
            * Note: Keeping a historical record would be beneficial with an option on whether to purge or not purge.  Folks will review other communities to try and get a larger understanding of the policy space that exists. The goal is create a workable, transparent, system.
            * Note: Decisions should be binding and final.
        * ~ 3:40 Brian: The steering committee may need a “office-hours” option.  Sarah is volunteering to send out an invite.
        * Tim: Do we have a process for retiring projects?  If there is no activity or no maintainers, there exists a kubernetes-retired org.
            * We can send an email to the list w/timeout prior to moving.
            * AI: Brendan will PR the original
        * Tim: What about repos being transferred into kubernetes-sigs, are there any CLA constraints.
            * There should not be, but folks can discuss with CNCF Legal.
        * Michelle: Our backlog needs some grooming, state = ?
            * Part of that was moving to github projects,
            * Jaice offered to move those items over and Michelle will help organize them.
            * AI: Michelle will sync with Jaice about having the backlog in order for us to work on regularly.
        * Michelle: Follow up on reviewing charters
            * Brian has started on a number of the charters and may be updating the main charter(s).
        * Derek: Do we have a policy on sub-projects having their own slack channels and mailing lists
            * Brian: That’s fine, that should not be a blocking issue.
        * Joe: As we look to fold sigs in, we may want to have more custom structures.
        * AI (everyone) : Review existing charter PRs.
        * AI: Michelle will follow up on how we should review those charters.

## May 9, 2018

* **Bosun: ** Joe Beda
* **Note Taker: **Aaron Crickenberger
* **Video: **[https://youtu.be/xtg3sWWRf4c](https://youtu.be/xtg3sWWRf4c)
* **Attendees:**
    * Tim, Joe - Heptio
    * Brian, Tim, Phil - Google
    * Michelle, Brendan - MSFT
    * Clayton - RHAT
    * Aaron - Samsung SDS
* **Topics**
    * Upcoming CNCF governance board meeting items [Michelle]
        * There is a licensing whitelist/blacklist proposal. Will send out slides.
        * That got voted on
        * The CNCF has to approve all licenses for projects in the CNCF as well as their dependencies (ie: compliance of all licenses, blanket approval via an allow / disallow list)
        * Yay all around, no objections and that got passed
        * Isn’t to say if you’re not on whitelist you’re not approved, you just have to go through a process
        * They run projects through (fossoligy?) on a per-request basis rather than continually
        * [brian] sig testing folks are working on automating the license verification
            * [https://github.com/kubernetes/kubernetes/pull/62088](https://github.com/kubernetes/kubernetes/pull/62088) PR from rmmh
            * Thockin will take a look at this PR
            * [https://github.com/kubernetes/kubernetes/issues/44505](https://github.com/kubernetes/kubernetes/issues/44505) contribex issue
            * Maybe this actually going through contributor experience?
        * Automating it turns out to be hard?
        * We also have a bunch of third party code that doesn’t necessarily import licnese files
        * There is an [issue filed on steering](https://github.com/kubernetes/steering/issues/21) with the list of things that are believed not to meet requirements
        * This seems like a CNCF-wide thing not a kubernetes/kubernetes thing
        * Is this a thing we can ask the CNCF to take for its projects?
        * Make following CNCF licensing guidelines requirement to become part of the project?
        * [Aaron] This is something I think belongs with contribex and we could have them raise in visibility, I haven’t heard about it lately. Will raise in priority with contribex
        * [Aaron] This was something that was talked about at LF OSLS repeatedly, hope we’re not reinventing the wheel here? Might be worth reaching out to some of the folks who presented on their efforts
        * [Michelle] going forward is this sort of stuff something I could send on to the steering committee? This was an example of cherry picking something I knew might be relevant to our interests
        * [Michelle] Will sending out the relevant information and start sending out summaries
        * “The marking stuff is a huge step forward from last year” brian
        * [Michelle] Will double check with CNCF if they have any automated compliance checking tools
        * [in chat] - Quinton provided this link [https://compliance.linuxfoundation.org/references/tools](https://compliance.linuxfoundation.org/references/tools)
    * [Various governance clarifications](https://github.com/kubernetes/community/pull/1994) PR
        * [tim+phil] imo we need a process doc, there appears to be a fair amount of confusion, e.g. - governance&lt;>charter is lost in translation.
        * Questions/comments/concerns for the process re: filing their charter
        * We need to start evangelizing and making clear what the expectations are and how the process should work
        * [joe] We need to close this amongst ourselves before getting it through
        * [brian] We had discussed using other sample charters to cherry pick what good examples are to come up with the ideal charter.  Put effort into sig arch, sig apps, sig scalability.
        * Contributor experience already merged theirs
        * [joe] Questions around the meta process
        * [Aaron] suggest we go to steering-private, use the existing charter drafts submitted as examples
        * Send things that work things that don’t through mailing list
        * Close loop on consensus by next meeting
        * Strawmen of meta process to be discussed by next
        * Answer: what constitutes a charter, how do we close on it, and then what is the process we broadcast to the community
        * [Aaron] Will assign people based on specific sig charters
        * [Michelle] will handle drafting meta process, brendan will act as reviewer, call last call for yay/nay at next meeting
    * [tim, quinton, joe] Voting rights proposal
        * Went back and forth on a couple of idea
        * Original simplest idea: just use devstas
        * Can roll up devstas across repos for github events etc
        * Draw a cut bar line across that
        * There will be some potential bad actors within that whole state space
        * But on average the signal to noise ratio is higher there
        * About on the order of 600 people
        * Separates this from the idea of sig membership
        * Is there room for “hey I’m not on the list but I think I should be”
            * [quinton] get a SIG to sponsor you
        * Is there room for blacklist people who are gaming the system
        * Specifics: All events, threshold of 100
        * So yeah you can emoji your way into voting
        * Quinton will take on calling last call
    * [tim] CNCF conformance group
        * [https://github.com/kubernetes/kubernetes/issues/62912](https://github.com/kubernetes/kubernetes/issues/62912) and conversation on sig-testing shows that they are free running with no proper feedback loop.  The goal was to have them work with sig-arch 1st.
        * The steering committee had OK’ed that we would hire a bunch of folks
        * But what was not clear was the mandate by which they were given to execute
        * What happened was prviate communication between Dan Kohn and folks at Google, unclear how they were executing
        * Bunch of random PR’s showed up to sig testing
        * The handoff of expectations was clearly not articulated
        * We as a group were OK with it, but there appeared to be a drop in handoff
        * The point of order here is about process or the lack thereof
        * … (I missed a bunch because I was talking - ed) ...
        * The leadup was clean but then I don’t know what happened once we had the money
        * That request for money, should that be talked about in the relevant sigs as well? Or should it be talked about after the money is allocated?
        * As long as there’s a plan
        * In this specific case, Brian will chase down Mithra to see what’s going on, and convey some of the things around process that we discussed here
        * Do we want to require any sort of checkpointing? How formal of a process should we get here?
        * [Aaron] Seems like long term yes? But not needed today? We reached more of a consensus
        * We really don’t want to be in the business of PM’ing these sorts of things, the idea is to have a SIG run it
    * [michelle] Code of Conduct Committee doc

## April 25, 2018

* **Bosun: **
* **Note Taker: **
* **Video: **n/a
* **Attendees:**
* **Topics**
    * Skipped due to lack of critical mass

## April 11, 2018

* **Bosun: **Aaron Crickenberger
* **Note Taker: **Michelle Noorali
* **Video: **[https://youtu.be/EcSQ4BF4LO4](https://youtu.be/EcSQ4BF4LO4)
* **Attendees:**
    * Aaron Crickenberger, Samsung SDS
    * Sarah Novotny, Google
    * Brian Grant, Google
    * Michelle Noorali, Microsoft
    * Quinton Hoole, Huawei
    * Tim St Clair, Heptio
    * Derek Carr, Redhat
    * Tim Hockin, Google
    * First Last, Company
* **Topics**
    * Follow Up - charter approvals [aaron]
        * Tracking [spreadsheet](https://docs.google.com/spreadsheets/d/1i_8oM1KlmXo90zIIgNTdOg-M6xesTggFLnTkgmscMlI/edit) - missing sig-ui, and some secondaries
        * Keep using spreadsheet to track progress, columns for discussed, drafted, done
        * What deadlines do we want to impose? Anything pre-kubecon eu?
        * Conversation around how much scope a SIG should define in their charter
        * Brian: would be useful to define scope which would inform new sub projects being in or out of scope
        * Derek: Is there a charter we should model against? One of the tensions seen is finding boundary of scope is tough. In SIG Node, it’s enumerated as a set of goals
        * Tim S (in chat): goals and non-goals = scope
        * Brian: [SIG Architecture charter](https://github.com/kubernetes/community/blob/master/sig-architecture/charter.md) as example charter with definition of scope
        * Chicken and egg problem. Should we iterate with a few SIGs and then show as shining example to the rest?
            * Brian disagrees here. Iterate as we go.
            * Tim: Should we put a timebox to create a charter?
        * Derek: Codifying process of a long standing SIG without outside effort as a first pass is useful.
        * Aaron: Let’s iterate on the charters that currently exist (including PRs) and come back to next steering committee meeting with feedback on those.
        * **AI: Making a commitment to push forward on the 6 charters that are in now and let’s come back to the table next meeting with what we like and don’t like **
    * Follow Up - project membership / contributor ladder replacement proposal
        * Have we progressed on this since last meeting?
        * 3 proposals are in and now we need to re-sync on those.
        * Sarah, Michelle, Aaron still interested in helping out here.
        * The current proposal represents a wide spectrum.
        * Discussion around what the constraints have been in the past for voting and running for steering committee.
        * Derek: The bar for voting / membership is higher than the one running for steering committee. Maybe there should be some minimum bar for both.
        * Quinton: How do we subvert hostile actors
        * Derek: Let’s control the probability of bad actors by also limiting nominees
        * Tim: Let’s limit the bar for one or the other (nominees or voters) and voting can be taken advantage of by popular folks on Twitter
        * Quinton: There are rules in place around how much of steering committee is changed at a time
    * Asks from sig-contributor-experience
        * Have mostly de-SPOF’ed, want to use a mailing list to handle requests for administriva such as sig creation, sig sunsetting
            * In lieu of a help desk
        * [community@kubernetes.io](mailto:community@kubernetes.io)? [kubernetes-community-admins@googlegroups.com](mailto:kubernetes-community-admins@googlegroups.com)?
            * Tim: may have ability to create [community@kubernetes.io](mailto:community@kubernetes.io)
            * **AI: Tim volunteers as tribute**
            * **Tim needs a list of people to add as owners on the list to add before removing himself.**
            * Are we setting precedent for others groups?
            * Gsuite accounts have features that public groups do not.
                * Audit logs are an example of a useful feature in some cases
                * Nesting mailing lists is another example
        * [Need docs on why / how to dissolve / convert sigs and wg’s](https://github.com/kubernetes/community/issues/2029https://github.com/kubernetes/community/issues/2029), can steering committee help with the “why”?
            * [onprem converting to a wg](https://github.com/kubernetes/community/issues/2030)
            * [cluster-ops moving to a wg or subproject of cluster-lifecycle?](https://github.com/kubernetes/community/issues/2031)
    * Governance
        * [sigs, wgs, committees… bofs?](https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/o4u0SvPUyyU)
            * Confusion in past about SIGs, WGs
            * Example: SIG Big-data, resurrected
            * There is a proliferation of use cases for Kubernetes and different groups of people wanted working groups.
            * Aaron: How many working groups have gone away?
            * Brian enumerates examples
            * Brian: WGs should have concrete outputs
                * There are certain kinds of working groups that are coming about
                * Use cases and personas are helpful output from working groups
                * We need more structure around working group cases
                * There needs to be some kind of process around creating a working group
                * BoF concept proposal. Let’s have a standard way that we let people know that conversations are happening
                * Useful to help people connect and coordinate
                * SIG Apps + developer WG conversation as example
                * SIG cluster lifecycle - some conversations happening there around subprojects. We push binaries in random places
            * Aaron: Everyone wants discoverability. Steering committee doesn’t want to endorse. Push everything as a sub project under a SIG to create some sense of heirarachy
            * **AI: Brian will give a tour of sub projects**
            * **AI: Brian, Michelle will talk about putting some structure around WG**
            * Quinton: maybe we’re over complicating
            * Brian: Norms should be documented
            * Let’s continue to discuss this on the mailing list
            * Kubernetes- mailing list infringes on the trademark. We
        * [Various governance clarifications](https://github.com/kubernetes/community/pull/1994) PR

## March 28, 2018

* **Bosun: **Joe Beda
* **Note Taker: **Aaron Crickenberger
* **Video: **[https://youtu.be/7yfKRJ59WUA](https://youtu.be/7yfKRJ59WUA)
* **Attendees:**
    * Aaron Crickenberger, Samsung SDS
    * Joe Beda, Heptio
    * Sarah Novotny, Brian Grant, Philip Wittrock, Google
    * Clayton Coleman, Red Hat
    * Quinton Hoole, Huawei
    * Brendan Burns
    * Michelle Noorali, Microsoft
* **Topics**
    * Followup from previous meetings
        * Charter approvals?
            * Have a [spreadsheet](https://docs.google.com/spreadsheets/d/1i_8oM1KlmXo90zIIgNTdOg-M6xesTggFLnTkgmscMlI/edit)
            * Err on the side of having someone with familiarity of the sig
            * Look at: scope, deviation from template
            * Do we have anything like a deadline at the moment?
            * Going around to sigs individually to explain what’s going on
            * Some sigs still don’t understand subprojects, they’re doing working groups instead of subprojects
            * Eg: sig cluster lifecycle didn’t look at governance or charter, a lot of this was news to them
            * Suggestion: why don’t we use the spreadsheet to sign up for outreach / Q & A
            * A number of sigs have created repos in kubernetes-sigs, not all of those sigs met requirements for being able to do that
                * Eg: they have to identify their subprojects, some of them didn’t do that and they got repos
                * So make sure we hold the line
            * Idea: have like a primary ambassador per sig? Monthly or quarterly?
            * Suggestion: contributor experience is already going around to SIGs on a quarterly basis, the more we keep them in the loop regarding governance decisions, the less legwork we would have to do to communicate that
            * AI: take the spreadhseet, and make sure the steering committee is signing up for at least X of these sigs (Aaron will be whip for this)
        * Project membership / contributor ladder replacement proposal? (Joe, Quinton, Tim)
            * Have had good discussions in terms of what we think is acceptable or not
            * What matters most: who and how will we determine how we do voting in the next steering committee election
            * Tensions between delegating voter roles to sigs so they can define things but also trying to make sure companies/sigs can’t game the system
            * One proposal paraphrased:
                * Say that it’s more like a representative democracy, where people inside of a sig vote for members, who in the sig gets to represent that sig, and then those members get to vote who’s in the steering committee
                * There may be sigs who are too populated by a vendor in leadership or membership
                * There are some sigs that ineffective
                * Rough notes: [https://docs.google.com/document/d/1PtNDBBdkDCyNXVau9v4oWGgBGgdFfL6IAtWuhM_HNzY/edit#heading=h.vlo3cwkmsazx](https://docs.google.com/document/d/1PtNDBBdkDCyNXVau9v4oWGgBGgdFfL6IAtWuhM_HNzY/edit#heading=h.vlo3cwkmsazx)
            * How to reconcile the democratic aspects with the demonstrated contribution model?
            * Alternate proposal
                * Let sigs maintain a list of roles / membership types
                * Those sigs have their own criteria for those things
                * Some number of months before the election we could take a snapshot of all the members, and review in that context the number of people in any given role for that sig, and determine whether that group is qualifying or non-qualifying
                * Eg: whole bunch of people who are chart maintainers, should being a chart maintainer qualify you for voting for steering committee
                * Will a sig try to pack members or game the system
                * Not just malicious intent: if each sig doesn’t get guidance on how many votes to dish out, they may weight inappropriately
                * If you’re a member of multiple sigs, do you get multiple votes?
                * We never intended for it to be sig representation, prove that you have skin in the game
                * But to be fair all this was done before the elections
                * Bburns: I don’t think we necessarily want proportional representation
                * Fairly happy with one person one vote, but if you had one person who made very large contributions in multiple sigs should they count for more, and I think the answer is they shouldn’t
                * Steering committee is beholding to the community not the sigs
                * Eg: we have cloud specific sigs, and we find that the leads there tend to be dominated by one company and then they add a ton of members, the value they add are disproportionate for that company
                * How much of this do we need to work out in advance vs. attempting to correct for obvious abuse?
                * Maybe we apply the same representation to guidelines to the voting population that we do to a given company?
                * But then how we decide that a given company is oversubscribed?
                * Do we need to cap votes or is capping representation on the committee enough?
                * We’ll keep noodling on this, appreciate more feedback on the doc or people to join the group and participate
                * AI: Sarah & Aaron interested if it doesn’t conflict (don’t need to move to accommodate), Michelle interested
        * KubeCon EU F2F or other F2F?
            * We already have a lot of meetings during kubecon and doing f2f things
            * Maybe it would be better to spend that time speaking with others
            * Sounds like no strong desire to have Yet Another Meeting
            * Option Fun Social Event Including Trustfalls And Molecular Gastronomy?
        * Values PR [https://github.com/kubernetes/steering/pull/25](https://github.com/kubernetes/steering/pull/25) (Sarah)
            * Go +1 sarah’s PR
        * Deprecating incubator, spelling out repository process (Aaron)
            * [https://github.com/kubernetes/community/issues/1922](https://github.com/kubernetes/community/issues/1922)
            * Down payment from Phil: [https://github.com/kubernetes/community/pull/1979](https://github.com/kubernetes/community/pull/1979)
                * LGTM’d (sarah)
        * Subprojects (Aaron)
            * [https://github.com/kubernetes/community/issues/1673](https://github.com/kubernetes/community/issues/1673)
            * Need to see what progress we’ve made here
            * How are we doing on communication
                * [https://github.com/kubernetes/community/blob/abeab354ca37f73453779921130364c20f3a5cfe/governance.md#subprojects](https://github.com/kubernetes/community/blob/abeab354ca37f73453779921130364c20f3a5cfe/governance.md#subprojects)
            * Next steps?
                * Label repos with topics based on subprojects?
                * Chase after repos that lack OWNERS? [https://github.com/kubernetes/community/issues/1721](https://github.com/kubernetes/community/issues/1721)
                * Add sig/owners fields to OWNERS files?
            * Feeling like we need to make them feel more like a thing?
                * Find orphaned code that nobody claims ownership of and delete it, see who yells?
                * We need the next level of tooling
                * Right now sigs.yaml creates readmes
                    * Do we create subdirectories for each subproject instead of a bulleted list?
                    * Sounds like we may as well try it
        * Code of Conduct Committee… TBD
        * Definition of a Working Group is rather thing, maybe this is something we should hash out
            * Need in writing what the difference is between sig/wg/committee
            * AI: Joe is signing up to do this, Michelle raised the issue

## March 14, 2018

* **Bosun:**
* **Note Taker:** Joe Beda
* **Video: **[https://youtu.be/Pp3gQw-iIAI](https://youtu.be/Pp3gQw-iIAI)
* **Attendees:**
    * Joe Beda - Heptio
    * Quinton Hoole - Huawei
    * Brandon Phillips - CoreOS/Redhat
    * Brian Grant - Google
    * Phillip Wittrock - Google
    * Tim St Clair - Heptio
    * Brendan Burns - Microsoft
    * Sarah Novotny - Google
    * Michelle Noorali - Microsoft
* **Topics**
    * Repo naming for kubernetes-sigs
        * Came up in a bunch of threads.  Started with a sig name prefix to make association clear and for metric automation and such.
        * ~6 examples now
        * people don’t like the pattern:
            * too verbose.
            * What happens as SIGs mutate over time
            * Abbreviate over time
            * Dashes or underscores
        * Examples of topics: [https://github.com/kubernetes-sigs](https://github.com/kubernetes-sigs)
        * API: [https://developer.github.com/v3/repos/#list-all-topics-for-a-repository](https://developer.github.com/v3/repos/#list-all-topics-for-a-repository)
        * General support -- approved
    * Charter approval requirements
        * What is the process?
        * Brian: need to have steering committee play a role for 2 reasons:
            * Keep consistent and well formed -- all new so we need to establish patterns
            * Spell out how scope -- need to validate that and approve
        * We need to be responsive -- SLO?
            * Load balancing responsibility?
            * For first chunk -- [create spreadsheet with owners?](https://docs.google.com/spreadsheets/d/1i_8oM1KlmXo90zIIgNTdOg-M6xesTggFLnTkgmscMlI/edit?usp=sharing)
        * Need to define OARP
            * O: sig has to drive it themselves
            * A: SIG asks steering@ for an approver
            * R: SIG members + rest of the steering
            * P: Community at large
    * Project Membership tracking
        * Brian proposal (joe would have said the same thing):
            * SIGs should identify additional roles (beyond coding)
            * Put these in OWNERS files (or if no place else, then contributors or sigs.yaml)
                * These should only be people with “skin in the game”
            * Also a place to do job postings and such
            * Project voting/”membership” is the transitive closure of all SIG/committee roles
        * Questions:
            * Does project voting == has a role?
            * What about people that write a lot of code but don’t have a formal role
                * Make them be reviewers?
                * Some people are sponsored by companies and don’t have bandwidth to do reviews
                * Automation precludes new people from doing reviews? How do people dip their toe in.
                * Perhaps this is a corner case.
                * Need to find a way to bring them into the org and have move impact.
                * Need to write this down.  Show how you can demonstrate skin in the game.
                    * Tragedy of the commons without common roles
                    * Is there a “prolific coder” role
                    * Samples of what is “above the bar?”
                * Are all SIGs equal? Senate vs house? Some SIGs will have more members and have more influence.
                    * We should curate SIGs at charter level
                    * Members are members -- no senate.
            * Renewal/expiration?
            * SIG stacking?
        * AI: smaller group of people put together a proposal for how this works
            * This will replace/augment the contributor ladder
            * Joe, Quinton, Tim
    * Licensing standards for dependencies
        * CNCF is looking at tools for continuous monitoring
            * Push back on CNCF that requirements need to come with automation
        * Tim Hockin has been driving most of the cleanup
        * Not everyone understands all the rules -- nested dependencies
        * AI: Brian remembers what he has already chased down?
        * AI: Brian to give feedback to CNCF about license tools again
        * AI: Michelle/Brandon give feedback to CNCF GB about license requirements and needing tooling
    * F2F KubeCon EU and Agenda? [Tim]
        * Is there critical mass?
            * Brendan may not be there
        * Use mailing list to see who is going.
        * Even partial group is worthwhile
            * Counterpoint -- use time to meet with folks that we don’t meet regularly
        * Some non-conf f2f?  Might be valuable in lieu of or in addition to kubecon f2f.
    * Code of Conduct Committee [Michelle]
        * Asked around at OSLS -- not super common
        * What’s the point? What do we want to get out of it?
        * CNCF wide CoC vs. just k8s?
            * Can put more resources into it -- CNCF lawyer
            * Learn across projects
        * Brian: need to rewrite CoC regardless.  Events in the past makes the CoC less than ideal. Will have to break new ground.
            * LF doesn’t have a great track record
            * Also increases likelihood that we are held accountable
            * Joint committee would be less familiar with our norms/players
        * Brendan -- on board -- no huge concerns
        * Quinton: seems difficult to get right in the first pass.  Hesitant to be too general until we have success at smaller scope.  May be sensible to start with k8s only and grow.
        * AI: move discussion to steering-private

## February 28, 2018

* **Bosun:**
* **Note Taker:** Philips
* **Video: **[https://youtu.be/bWXVBsyRppg](https://youtu.be/bWXVBsyRppg)
* **Attendees (10/13):**
    * Michelle Noorali - Microsoft
    * Tim St. Clair - Heptio
    * Joe Beda - Heptio
    * Brandon Philips - Red Hat
    * Derek Carr - Red Hat
    * Phillip Wittrock, Brian Grant, Tim Hockin, Sarah Novotny - Google
    * Quinton Hoole - Huawei
* **Topics:**
    * Review of progress on[ current sprint](https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-feb18) / outstanding efforts
        * Update project values
            * Sarah needs to do a PR about it
            * User voice from Chen is interesting, but need to put it onto the next sprint
            * Brian Grant said we have to define who our users are
                * Multi-tenant, conformance, etc
        * Intra-SIG governance / SIG charter template
            * Philipp got the [requirements doc merged](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance-requirements.md)
            * There is a PR out for updating specification for having subprojects define technical requirements as well
            * Short template [out for review](https://github.com/kubernetes/community/pull/1830)
                * What to call the non-technical lead - vote is out
            * Timeline for when SIGs should have their charter?
                * No. Don’t want to set it up until we’ve tested a few.
                * Idea: you can’t sponsor a new project until you have a charter in place
            * .
        * Repository proposal
            * Brendan owns this. Aaron finished setting up automation.
            * As far as we know we are ready to go, blocked on SIGs doing homework (see above) / open for business
        * Formalize concept of subprojects
            * Brian Grant: we took a first stab at categorizing all of the subprojects, added owners files, and updating owners files
            * Phil has a PR out to merge owners and maintainers
            * Structure: effectively two types of orgs. github.com/kubernetes/kubernetes also known as core; and then subprojects that are in a SIG Github Org
    * Any CNCF GB issues Michelle and Brandon need to be aware of?
        * Brandon's notes:
            * Via Dan Kohn: Approving up to $600 K for improving K8s conformance tests, now that SIG Architecture has implemented a policy of requiring conformance tweets of all stable features, as the GB requested in December.
                * $600K spending will be overseen by SIG Architecture
                * Approved spending from Steering Committee
            *  $60-170k Bug bounty
                * [https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/S71W5sBcdVc](https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/S71W5sBcdVc)
                * Discussion about the kubernetes security committee: [https://kubernetes.io/security/](https://kubernetes.io/security/)
                *
    * CNCF graduation
        * We passed!
    * Next Sprint Teams/Items?
* Actions
    * Michelle: Need to explain what subprojects are, announcements at community meeting are not sufficient
    * Michelle: README in community repo for subprojects
    * Phill: Agree on new Name for SIG Lead
    * Phil: Update short template
        * If nothing contentious, merge
        * If not merged by Friday we won’t block creation of new projects based on lack of SIG charter
    * ???: We could get feedback from SIG contribx on experience of writing a charter
    * ???: Send email to sig leads mailing list. Phil to write doc framing message
    * Michelle: find/send Jago’s proposal to ???
    * Philips: tell Maya that we are going to raise the bug bounty program budget to the CNCF GB
    * Philips: rename the product security team to product security committee

## February 14, 2018

* **Bosun:** Aaron
* **Note Taker:** Joe Beda
* **Video:** [https://youtu.be/bDVFHvq1EmM](https://youtu.be/bDVFHvq1EmM)
* **Attendees:**
    * Joe Beda, Tim St. Clair- Heptio
    * Phillip Wittrock, Tim Hockin, Brian Grant, Sarah Novotny - Google
    * Aaron Crickenberger (@spiffxp) - Samsung SDS
    * Derek Carr - Red Hat
    * Michelle Noorali, Brendan Burns - Microsoft
* **Topics:**
    * Review of progress on[ current sprint](https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-feb18) / outstanding efforts
        * Update project values
            * Sarah: really close. Need to look for address last comments.
            * Tim: I have one PR I want to get in there. Put in something around sustainability.
            * Brian: automation over toil plays in. Also continuity is about turning over leadership.
            * Sarah: I’ll see if I can put anti-heroism in here some place.
            * [Sarah types]. Thumbs up.
            * Done! Sarah will get it done.
        * Intra-SIG governance/SIG charter template
            * Phil: [Churning on PR](https://github.com/kubernetes/community/pull/1650), restructured into 4 docs
                * Requirements on SIGs
                    * Brian: Include the scope of the SIG
                * Templates are starting places for SIGs (requirements, short, long)
            * Things to resolve
                * Subprojects -- how to delegate and such
                * Name for non TL (Phil thought “SIG Lead” but Joe and Tim aren’t fans).
                * Do SIGs need to get charter approval from steering committee?
                    * Consensus: yes
                * Suggested way to track membership
                    * Why do we need this? When the SIG has to vote on stuff. Also plays into steering committee voting
                    * Sarah: if we do have this we want to have *some* level of proscriptiveness to minimize politics.
            * Decision: Let’s get something out even if there are some TBDs
                * [New reduced scope PR](https://github.com/kubernetes/community/pull/1800)
        * Repository proposal/“Incubator process” (what gaps are left after repository proposal)
            * [spiffxp] AFAIK automation is setup
            * Brendan -- addressed a lot of comments and about to push up new version with comments addressed.
            * What process do we use? Have everyone on committee LGTM?
            * Aaron: open questions with implementing this. What about teams? [https://github.com/kubernetes/community/pull/1799](https://github.com/kubernetes/community/pull/1799)
            * Is this an org per SIG or one org? We are starting with single org (kubernetes-sigs) with naming convention.
            * We will probably have ~10-50 repos in the first month.
            * Need to spell out logistics someplace? Separate out governance vs. logistics.
                * Brendan will send a follow PR with process
        * Formalize concept of subprojects
            * Identify all OWNERS files that don’t exist and pre-populate
            * Brian: start asking SIGs/repo owners to do this
            * Brian: merge owners and maintainers in ladder
            * Need automation to create bi-directional links.
                * Perhaps have someone from ContribX build this out
    * Discuss items we want to bring up at the next CNCF GB meeting [Michelle]
        * Michelle had technical issues and had to drop :(
        * $600K for contractors to backfill tests. Contractors have been identified.  Clock is ticking.
        * Brian has context: Initial conformance test suite has low coverage.  Want to up the coverage by writing more tests.  Proposals within the conformance efforts to get automated broad coverage.
        * CNCF Need steering committee to approve this.  Brian hasn’t been happy with the proposals so far.  Would go through SIG-Architecture and steering committee would approve.  Can’t have this level of funding forever.  SIG-Arch will need to have some level of policy around getting to GA perhaps.
            * Test framework to have same test run both in unit test and e2e test.
        * Next action on SIG-architecture and CNCF conformance WG.
        * (Some discussions around the fact that this is a CNCF level WG -- accident of history).
    * Discuss communication infra (google groups, groups.io, k8s.io) and tracking SIG membership (email list, OWNERS). What do we expect from sub-projects wrt consistency in this area? [Joe]
        * Opt to share docs as public by default so as to avoid excluding folks who can’t login to google (ie: China)
        * Joe will work with contribx to explore other options (discourse?)
            * Accessibility from china
            * Groups/forums that can be driven via API from OWNERS files so we have one source of truth of who is what. We want to have one source of record that is tracked.
            * Better discoverability experience — all of the groups in one place vs. having it spread throughout groups.
            * Email connections
        * We will use OWNERS-ish files for tracking SIG membership.  Syncing to mailing lists isn’t possible with current google groups and will be TBD until we have bandwidth to do tooling and make a transition off google groups.
    * Code of conduct committee
        * [spiffxp] does steering need to be involved in all moderation decisions? sig-contributor-experience is taking these duties on in their charter as part of the community-management subproject
        * should steering be listed as point of contact for code of conduct? Currently it’s [Sarah or Dan on the CNCF CoC](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
        * We need to come up with concrete proposal and CNCF is up to adjusting their CoC.
    * [sarahnovotny] [project graduation at CNCF](https://github.com/kubernetes/steering/issues/12)

## January 31, 2018

* **Bosun:** Aaron Crickenberger
* **Note Taker:** Phillip Wittrock
* **Video:** [https://youtu.be/j8tj13gFYZY](https://youtu.be/j8tj13gFYZY)
* **Attendees:**
    * Tim St. Clair, Joe Beda - Heptio
    * Aaron Crickenberger - Samsung SDS
    * Derek Carr, Clayton Coleman - Red Hat
    * Phillip Wittrock, Brian Grant, Tim Hockin, Sarah Novotny - Google
* **Topics:**
    * [jbeda] Heads up -- rename from Master to something else? [kubernetes/website#6525](https://github.com/kubernetes/website/issues/6525)
        * Already had this conversation before, and made a decision not to change the name
        * Master / Worker common in literature to reference
        * If we want to make this change, need to figure out the cost of the change
            * How much code needs to be changed - maybe not that much, but some?
        * SIG Architecture decision
        * Probably will continue to come up both inside and outside the community
    * [briangrant] Review of progress on [current sprint](https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-jan18) / outstanding efforts
        * [tstclair] update on [brendans proposal](https://docs.google.com/document/d/1yauW9zMtWgXN8xh4q6144B2xBTPuIOLGc0L6aQPMj1I/edit?ts=5a3c8c55)?  Community is asking for direction, and I think we owe some formal statement during community meeting.
            * Open issue - What is the GA terminology for these?  Can SIG projects publish “V1” images?
            * What is the feedback and issues from org-structure?
                * Multi-org has issue(s), for the time being we’ll create a kubernetes-sigs org.
            * Brian’s feedback
                * In November we agreed to unblock SIGs that needed repos
                * Decouple goals and directional aspect, provide an interim solution, and defer finalizing the details
            * Goal to get this out to the community ASAP and allow SIGs to start creating repos
            * Stop gap proposals
                * Plan to continue to use incubator as a stop gap for SIGs now
                    * Some concern about doing this
                    * CLA bot is easy to set up for new orgs
                * Create kubernetes-sigs org and SIGs can create repos there
                    * General agreement to this solution
                    * Will defer the GA question since it is controversial
                    * Should have a top levels OWNERs file with the SIG that owns
                    * SIGs accept through lazy consensus
            * Will cut the proposal down and send tomorrow
        * [briangrant] Resurrecting the interim SIG repo agreement from Nov 8 (see notes below)
        * [sarahnovotny] [Kubernetes Values](https://docs.google.com/document/d/12J0YLKBOfBTMd8ZdPpFMIAhMmGIfapMg-nkYPJYY7FY/edit#) -- several seem concise and uncontested.  Two still need some refinement.
            * Tightened up some things, reduce content
            * Keep it high level and strategic
            * &lt;Reviewed during the meeting>
            * Tiny wordsmithing left
        * [spiffxp] Resolve the disconnect between code organization (OWNERS, maintainers) and people organization
            * [Initial implementation](https://github.com/kubernetes/community/pull/1674)
            * How to roll out subprojects without making people feel voluntold?
                * e-mail to kubernetes-dev requesting PR review?
                * Merge and e-mail to kubernetes-dev / present at community that this is first pass, list of subprojects isn’t set in stone, feedback appreciated?
                * Set a deadline for sigs to finalize their subprojects?
                * Set a deadline for brian to get the list of subprojects out of his head?
            * Next steps?
                * Add sig, subproject fields to OWNERS files (would like brian’s help on scoping this)
                * Start using automation to audit:
                    * Are these OWNERS files even present?
                    * End goal (next sprint?) is to treat owners files as source of truth rather than k/community
                * Merge maintainer into owner role (would like brian to take lead on this)
            * Can send email to SIG leads to ask to communicate to relevant people in their SIGs to take a look
                * Need to communicate adoption of subprojects
            * Why: Aligning code and organization
            * Focussing on strawman human readable implementation before starting to shard
            * Taking best guess as single repo things - need help with multi repo projects
            * Will present about this at the community meeting tomorrow
            * Start talking about subprojects in OWNERs files
            * Will be updating the membership ladder
                * Want to separate ladder from GH roles
        * [pwittrock] SIG Governance
            * Proposal out for review [here](https://github.com/kubernetes/community/pull/1650)
            * Need a companion doc with the interface for what SIGs have to fulfill
            * Not ready to solicit feedback from community - needs to be reworked
            * Want a “checklist” for SIGs to check off, and a reference implementation
            * Steering committee involvement should be exception process, not oversight process
    * [jbeda] Get an answer to Jorge/Paris on what we’d like to see out of contributor summit. [Email thread.](https://groups.google.com/a/kubernetes.io/forum/?utm_medium=email&utm_source=footer#!msg/steering/KpEZ8KNtgOY/iuNvQVLwAAAJ)
    * [sarahnovotny] [project graduation at CNCF](https://github.com/kubernetes/steering/issues/12)
        * We will be the first project to graduate
        * Will set a precedent for other projects - want to make sure that we meet the criteria
        * Size of the project makes it hard to make big changes
        * We meet the basic criteria, and other projects are mirroring what we are doing - setting a precedent whether we intend to or not
        * No downside to graduating - governance has light touch
        * Toc does voting through email - has benefits
            * Can be answered asynchronously
            * Has permanent record that is less easily changed than through GitHub
            * Maybe we should add this process to the charter
    * [sarahnovotny] mentoring program feedback.  [Email thread](https://groups.google.com/d/msg/kubernetes-wg-contribex/kisf4lrqUG0/MLk8LNePAwAJ)
    * [pwittrock] Should we have a weekly update at the community meeting?
        * Answer: Let’s do a bi-weekly meeting
    * Paris and Jorge requested input in process for contributor summit invites
        * Would prefer to get input early instead of afterward
        * Speak now or hold your peace

## January 17, 2018

* **Bosun: **bburns
* **Note Taker: **jbeda
* **Video:** [https://youtu.be/PyjUlxxbJ9Y](https://youtu.be/PyjUlxxbJ9Y)
* **Attendees**
    * Aaron Crickenberger [Samsung SDS]
    * Michelle Noorali [Microsoft]
    * Quinton Hoole [Huawei]
    * Joe Beda, Tim St. Clair [Heptio]
    * Phillip Wittrock [Google]
    * Tim Hockin [Google]
    * Derek Carr [Red Hat]
    * Brian Grant [Google]
* **Topics**
    * [spiffxp] [Where are we on our sprint items?](https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-jan18)
        * Project values
            * Brendan and sarah haven’t sync’d so there are no updates. Brendan to follow up.
        * Intra SIG governance
            * Sent survey to SIG leads
                * Long form answers
                * Some good takeaways -- big diversity of needs.  Some sigs have the structure they needs, some need some guidance.
            * Have survey for community - plan to send out tomorrow
                * Multiple choice questions -- easier to aggregate
            * Started document to outline intended organization structure + provide a charter template that follows the structure - [link](https://docs.google.com/document/d/1JU4LLgF2OFQQJyTec-MK5ykVQjY5TzE2WlS5ob-TD6A/edit#)
        * Check governance docs into community repo
            * [tstclair] PRs are merged or /lgtm’d move to ratify.
            * [decarr] PR for changes to charter [https://github.com/kubernetes/steering/pull/22](https://github.com/kubernetes/steering/pull/22)
                * All members of steering committee will `/lgtm` that PR.
        * Resolve disconnect between code org and people org
            * [Brian’s proposal](https://docs.google.com/document/d/1FHauGII5LNVM-dZcNfzYZ-6WRs9RoPctQ4bw5dczrkk/edit#heading=h.aptlokk78hrh)
            * Suggest starting with formalizing subprojects
            * [Strawman of adding subprojects to sigs](https://github.com/kubernetes/community/compare/master...spiffxp:sig-ownership-strawman)
            * Get list of subprojects beyond repos
            * Merge [maintainer](https://github.com/kubernetes/community/blob/master/community-membership.md#maintainer) role into [owner](https://github.com/kubernetes/community/blob/master/community-membership.md#owner) role
            * Discussion:
                * Lots of discussion around what this the sub-project proposal is used for
                * Worries about too much bureaucracy.
                * Goals: Who owns the code in larger sense (SIG), accountability, aligning people structures with code structure.
                * Example -- kube-proxy.  Subproject of SIG-network. Carving it off enables breaking out into separate repo.
                * Questions around what is authoritative -- the stuff in community repo or the references from OWNERS.  Need toolability here. Aaron starting with community repo stuff.  Don’t wait on tooling to get started this.
                * Also find a way to have additional metadata around sub-project -- meetings, notes, channels, etc.
                * owner vs. maintainer -- just merge them together.  Need to socialize these.
                * What does making a group an owner mean? Not part of brian’s proposal. Owner is effectively TL of a sub-project.  Not uniformly recognized/specified/implemented.
                * Brian proposal: explicit role for sub-project.
                * Michelle -- need to have a better way to reach out to repo owners/maintainers/etc.  Do we want to add email addresses to these lists?
                    * We get nasty mail from github when we put email addresses into repo
                    * Ideally SIG leads communicate down
                * One of the thing that Aaron did was start creating consistent OWNERS_ALIASES around SIG/subprojects/git leads. Then listed this into approvers.  This caused confusion and came across as a power grab.  Just starting with PR came across as aggressive.  Not having good contact info meant that it was more difficult to get this stuff done.
    * [tstclair] Piloting policy changes and being mindful and deliberate.
        * I'd like to delay the pushing of policy until we've settled on the proposal and have a sunset timeframe on what is in and out of k-org.  Once that is done then k-org should be a well defined set of repos where I agree that we should define and enforce policy.  Even then, we should pilot the program and iterate with feedback to ensure that policies are not onerous.  E.g. - we should not be rolling test-infra out on any repos until we’ve settled on where they should live, and even then we should pilot and gain feedback before we try to implement other policies.
        * [spiffxp] attempting to document list of current kubernetes github orgs and the policies that should be applied [https://github.com/kubernetes/community/pull/1569](https://github.com/kubernetes/community/pull/1569)
        * [spiffxp] attempting to make kubernetes-template-project a living example of the guidelines that repos should follow [https://github.com/kubernetes/kubernetes-template-project/pull/14](https://github.com/kubernetes/kubernetes-template-project/pull/14)
        * [spiffxp] would like to deprecate / remove the outside collaborator role [https://github.com/kubernetes/community/issues/1565](https://github.com/kubernetes/community/issues/1565)
        * Discussion:
            * Aaron -- just trying to canary stuff out and was difficult to communicate. Didn’t view it as top down but it was percieved that way.
            * Brendan -- we should make sure that we categorize repos and top-down based on which repo we are talking about. Stuff in k/* will be top down.
            * Michelle -- communication snafu.  Did this go out to mailing list? Need to find ways to reach out before big changes land.
            * Aaron -- steering committee shouldn’t be involved in every automation rollout like this.
            * Brian -- no way to reach everyone. Need to look at worked before -- example is the roll out of 2FA.  The roll out of github teams/aliases didn’t work so well.
            * Brendan -- need to watch out when we give power to SIGs they don’t overstep. Need to make sure we give folks a way to opt out and are clear about rights/responsibilities.
    * [tstclair] Follow up on Brendan’s doc and outline feedback
        * [Docs ](https://docs.google.com/document/d/1yauW9zMtWgXN8xh4q6144B2xBTPuIOLGc0L6aQPMj1I/edit?ts=5a3c8c55)
        * Brendan going to do another pass to resolve things
        * Some agreement to roll share this widely to get agreement
        * Brian: the CLA/CNCF issues need to be vetted.  Can we do associated projects?
        * Intent is understood -- questions about mechanism.

## January 3, 2018

* **Bosun**: Brandon Philips [CoreOS] (@brandonphilips)
* **Note Taker**: Brandon Philips [CoreOS] (@brandonphilips)
* **Attendees**
    * Aaron Crickenberger [Samsung SDS]
    * Brandon Philips [CoreOS]
    * Timothy St Clair [Heptio]
    * Brendan Burns [Microsoft]
    * Derek Carr [Red Hat]
    * Tim Hockin [Google]
    * Phillip Wittrock [Google]
    * Clayton Coleman [Red Hat]
    * Joe Beda [Heptio]
* **Absent**
    * Brian Grant [Google]
    * Michelle Noorali [Microsoft]
    * Sarah Novotny [Google]
    * Quinton Hoole	[Huawei]
* **Topics**
    * [tstclair] Split up charter and elections. [https://github.com/kubernetes/steering/pull/18](https://github.com/kubernetes/steering/pull/18)
    * [philips] Reminder on sprint [https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-jan18](https://github.com/kubernetes/steering/blob/master/backlog.md#current-sprint-end-of-jan18)
    * [aaron] administrivia
        * Permissions [https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/xXG-cUqoBbw](https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/xXG-cUqoBbw)
        * List of orgs we “own” ? [https://github.com/kubernetes/community/issues/1407](https://github.com/kubernetes/community/issues/1407)
        * [https://github.com/kubernetes/community/issues/1527](https://github.com/kubernetes/community/issues/1527)
        * Need a bot or something to manage all of the different orgs
        * TimH is going to give Aaron owner on github.com/kubernetes to start setting up management
        * TimH: Everyone needs to be safe with this responsibility
        * Aaron is going to spec out a tool to manage the owners access via some authoritative file that we can have the CNCF hire a contractor to implement
            * Needs to eventually be the only authoritative source
        * TimH: Write a
    * [bburns] Proposing repo-structuring
        * [Docs ](https://docs.google.com/document/d/1yauW9zMtWgXN8xh4q6144B2xBTPuIOLGc0L6aQPMj1I/edit?ts=5a3c8c55)
        * Open questions:
            * What happens to vanity URLs for Go packages
            * What happens to addons that don't have to go into core: Service Catalog, Helm, etc. Can they remain SIG projects forever?
            * What is the difference between a SIG repo, a kubernetes/ repo, and an associated repo?
            * Associated repo user story: it is about frictionless progression into the Kubernetes community
            * End of life of a project and stability expectation setting
    * [joe] Intra-SIG governance survey
        * Phil put it together: [https://docs.google.com/forms/d/1h41mwLlcQwQDP6c4VDZzjiJepaQ1irHjP4mXYQFQAIo/edit](https://docs.google.com/forms/d/1h41mwLlcQwQDP6c4VDZzjiJepaQ1irHjP4mXYQFQAIo/edit)
        * Goal is to source ideas for a recommended solution
