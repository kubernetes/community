**12/15/2020**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉



* Joseph Sandoval
* Rolfe DH (Red Hat)
* Abubakar Siddiq (GitLab; met some SIG Docs folks at a KubeCon)

Updates/reminders



* This week’s PR wrangler: @vineethreddy02
* Next week’s PR wrangler: _Holiday_
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda: 



* Release 1.20 retrospective
    * Anna: release went well
        * PRs piled up (avoid)
        * Good job with reviews and approvals
        * Comms comms (communication with SIG Release comms team especially for the blogpost release)
    * Open floor
    * Rey: thumbs up as a shadow / comms internal
    * Rey: getting things done early is good
    * Rey: some trouble around communicating important changes (eg deprecations)
    * Joseph: Anna killed it! From comms, no handbook mention of relationship
        * Unaware of SIG docs’ ability to assist
        * Knowing of blog before day of would be good
        * Knowing of docs deadlines in comms would be good
        * Themes could improve (deadlines might help)
    * Kaitlyn: Historically the blog team has removed themselves from the review process. Helped if needed, but it would be good to make a decision if we want blog team review or not. Happy to help update the role handbook.
    * Anna: PR timing / who is responsible?
    * Jim: more than one way to do it. Docs release lead (or shadow) could meet with the release comms team, maybe halfway through the cycle, and pick an approach. _Doesn’t have to be the same approach every release?_ Could be more formal if needed.
    * Jim: slimming down the release handbook is a positive thing - it’s a living (resurrected?) document!
    * Tim: We did a good job for being an “enabling” SIG. 🎉🎉
    * Celeste: Was kubecon timing a problem?
* 1.20 Release Retro (idea for 1.21)
    * Discussion for the release comms for deprecation [https://github.com/kubernetes/community/issues/5344#issuecomment-740218218](https://github.com/kubernetes/community/issues/5344#issuecomment-740218218) 
    * Anna: Move placeholder deadline to match code freeze (currently a week before)
    * Tim: easier for review / feedback over longer. However this may affect the suggestions on what kind of feature gate needs to be included on the docs. Opening a placeholder PR late could also be a signal if the feature is ready to be graduated since there’s also still a long list that needs to be done in the code /comms about the feature itself. 
* Blog Update
    * Nothing new and exciting
        * need to raise awareness on the process
        * Looking for more reviewers (please attend the blog team meeting on thursdays)
* Issues & PRs
    * (for next meeting) [Celeste]: [https://github.com/kubernetes/examples/pull/393](https://github.com/kubernetes/examples/pull/393), [https://github.com/kubernetes/website/pull/22934](https://github.com/kubernetes/website/pull/22934) 
        * Issue 1: WG Naming is okay with this in principle from a language perspective, but this replacement changes MongoDB to Redis. Is SIG Docs ok with this?
        * Issue 2: Paul may not have time to push these over the line – can we take over the PRs?
* Discussion
    * Release 1.21 upcoming work
    * 

**12/8/2020**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉



* Walid Shaari

Updates/reminders



* This week’s PR wrangler: @daminisatya
* Next week’s PR wrangler: @vineethreddy02
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * On course for release - Anna
    * Including blogs - PRs (only) done yesterday, ready to go post-release
        * Also good to cover in retrospective
* Blog Update
    * Kaitlyn - some articles not getting assigned to right people
    * Blog subproject meeting on Thursday
    * Might need help reviewing blogs w/ release
        * Rolfe DH / Chris Metz / Tim interested
* [Brad] Update from Monthly SIG Docs Localization subgroup meeting on 12/7/2020
    * Several issues discussed
    * Multiple styles for contributing localizations
        * 1st - most things up to date, tracking when the English pages are more recent than the last change to a localized page. “Typo fixes” mask whether or not the localized page is current with respect to English. Within this model there is also diff-based and timestamp-based approaches to work out what to sync.
        * 2nd - batched, behind the master branch of English. “Typo fixes” fine.
    * Accidental approvals by English-localization approvers (SIG Docs leads & chairs) are frustrating sometimes for localizations.
    * Some people don’t look at GitHub labels for messages, they look just at PR title. Some localization teams are going to recommend porefixing.
    * **AI:**  Brad to open issue to describe the issue faced by localization teams regarding timely update of .toml and .nd files that they don’t have authority to approve PRs on
* Issues & PR
    * Merging GSoDs API ref (??)
        * [https://github.com/kubernetes/website/pull/23294#issuecomment-736300275](https://github.com/kubernetes/website/pull/23294#issuecomment-736300275) 
        * Tooling change merged, update for reference docs to follow v1.20 release
        * Next version for reference docs close, changes to preserve existing hyperlinks are important - Karen (very important) - Tim
        * Tim to log issue about preserving those existing hyperlinks
    * First Sig-Security-Docs Recommendation
        * PR: [https://github.com/kubernetes/website/pull/25498](https://github.com/kubernetes/website/pull/25498)
        * Related issue: [https://github.com/kubernetes/website/issues/25497](https://github.com/kubernetes/website/issues/25497)
    * Style guide revision (Geoff Cline)
        * [https://github.com/kubernetes/website/pull/24645](https://github.com/kubernetes/website/pull/24645)
* Other items 
    * Anna - Docs release retrospective
    * + comms team

**12/1/2020**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉



* Rolfe - tech writer at RedHat
* Bredam - tech writer at IBM 

Updates/reminders



* This week’s PR wrangler: @sftim
* Next week’s PR wrangler: @daminisatya
    * AI: Jim to follow up
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * Kristin - Yellow / Red
        * 3 open PRs each needing action from owner (DM’d waiting)
            * Retro (how to get a hold of owners)
        * Docs merge deadline EOD tomorrow
        * Potentially merge as-is
* Blog Update
    * Move forward PR
    * Might need help reviewing blogs w/ release
        * Rolfe DH / Chris Metz / Tim interested
* Issues & PRs
    * Merging GSoDs API ref
        * [https://github.com/kubernetes/website/pull/23294#issuecomment-736300275](https://github.com/kubernetes/website/pull/23294#issuecomment-736300275) 
        * Maybe we can merge the tooling now, and then target the website change for the dev-1.21 branch?
* [Celeste] Updates from WG Naming
    * [https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md](https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md) is in!
        * L18ns still use master:
            * [https://github.com/kubernetes/website/issues/22580](https://github.com/kubernetes/website/issues/22580) 
            * [https://github.com/kubernetes/community/issues/5043](https://github.com/kubernetes/community/issues/5043) 
    * [https://github.com/kubernetes/website/issues/21749](https://github.com/kubernetes/website/issues/21749) – Why ‘live’? 
            * It’s level with the live content - Tim B	
            * Doesn’t feel strongly about it, go with main otherwise
* [Brad]  Kube Localization Subgoup meeting is not showing up on calendar at appropriate UTC time (at least on my calendar)
* [Jim] localization partial updates
    * Came up last week, might be better for localization subgroup
    * If doing a typo fix, update entire content
        * Causes issues with timestamp checks for updates.
    * Celeste: Trying to track change / importance
    * AI: add to sub project agenda
    * AI: create an issue to add a “last big update” frontmatter
        * Need a robot to check / validate
    * AI: User stories for localization (at next meeting)
* [Kohei] Please do not approve localization PRs before confirming each language owners
    * [https://github.com/kubernetes/website/pull/25280](https://github.com/kubernetes/website/pull/25280)
    * Irvi -  I was wondering if there’s a way to separate owner for the admin and each of localization. So if the owner of the localisation haven’t approved it then it won’t be merged.
* [Savitha] Sig-security-docs sub project has planned to meet on Dec 3rd @ 2 PM EST for the first time. We will be focusing on improving docs related to K8s security, creating a security hardening guide, CKS materials, and so on. If you are interested, please come join us by subscribing to[ sig-security mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-security). Once you subscribe to the mailing list, you will get the calendar invite with meeting details. Looking forward to see y'all at the meeting 

<p id="gdcalert1" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/image1.png). Store image on your image server and adjust path/filename/extension if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert2">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/image1.png "image_tooltip")
 Happy Honking!!  \
 \
You can also reach out at [https://kubernetes.slack.com/archives/C01D8R7ACQ2](https://kubernetes.slack.com/archives/C01D8R7ACQ2)
* [Geoffrey] first draft of the first new contributor video
    * [https://groups.google.com/u/1/g/kubernetes-sig-docs/c/28xtYr7PUPM](https://groups.google.com/u/1/g/kubernetes-sig-docs/c/28xtYr7PUPM) 

**11/24/2020**

**5pm Pacific (01:00 UTC)**

🎉 New contributors 🎉



* Youqing - user interested in Chinese
* Christopher Metz - from KubeCon, writer for years, working on technical documentation
* Yue Ma - what’s going on, never contributed before, start at docs.

Updates/reminders



* This week’s PR wrangler: **US Thanksgiving **- @qiming
* Next week’s PR wrangler: @sftim
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * YELLOW
    * PR “ready for review” deadline was yesterday
        * 3 PRs extended until tomorrow (11/25 EOD)
    * We have 5 PRs that’s ready to be reviewed and 2 PRs that’s in draft
    * Next deadline: Tuesday 12/2 all PRs “to be merged.”
    * More info is slack
    * [https://github.com/kubernetes/sig-release/tree/master/release-team/role-handbooks/docs](https://github.com/kubernetes/sig-release/tree/master/release-team/role-handbooks/docs) 
* Blog Update
    * No updates
* Issues & PRs
    * (release things) 
        * [https://github.com/kubernetes/website/pull/25226/files](https://github.com/kubernetes/website/pull/25226/files) 
        * [https://github.com/kubernetes/website/pull/24921](https://github.com/kubernetes/website/pull/24921)
    * Team awareness
        * [https://github.com/kubernetes/website/pull/23294](https://github.com/kubernetes/website/pull/23294) 
            * Feature gaps in solution
            * Localization adoption?
            * Pressure to merge?
            * **AI: **Qiming - will boost visibility
* Localization teams might have issues similar issue:
    * “If you want to fix small change, validate entire page is synced with entire site”
    * No partial updates
        * What do other localization teams do?
    * Track upstream (zh)
    * If typo changes, timestamps no longer work for tracking
    * Can share with the localization sub project
    * **AI:** Jim will bring up next meeting and for localization working group
* Kubelet config file not documenting
    * [https://github.com/kubernetes/website/pull/25094](https://github.com/kubernetes/website/pull/25094)
    * [https://github.com/kubernetes-sigs/reference-docs/pull/176](https://github.com/kubernetes-sigs/reference-docs/pull/176)
    * Many parts can be generated in markdown and copied to website
    * Jim: GSoD work - different - uses swagger not source code
    * Multiclusterlifecycle review?
        * Until API is published - deprecate after
    * **Need: reviews / opinions**
* [Celeste] Updates from WG Naming
    * [https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md](https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md) is in!
        * L18ns still use master:
            * [https://github.com/kubernetes/website/issues/22580](https://github.com/kubernetes/website/issues/22580) 
            * [https://github.com/kubernetes/community/issues/5043](https://github.com/kubernetes/community/issues/5043) 
    * [https://github.com/kubernetes/website/issues/21749](https://github.com/kubernetes/website/issues/21749) – Why ‘live’? 
            * It’s level with the live content - Tim B

**11/17/2020**

**10:30am Pacific (18:30 UTC)**

Agenda



* KubeCon Social Hour - No Agenda, feel free to promote on Twitter / Slack / etc.

**11/10/2020**

**10:30am Pacific (17:30 UTC) **(Long time no meeting!)

🎉 New contributors 🎉



* 

Updates/reminders



* This week’s PR wrangler: @jimangel
* Next week’s PR wrangler: **KubeCon!! **@sftim might chip in
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* KubeCon
    * social hour
    *  k8s.io boost out tweets account (and CNCF twitter)
* Release 1.20
    * Docs deadline is not extended (Nov 23d - ready to review)
        * Beyond Dec 8th should not be moved
        * Tight timeline
        * 58 enhancements (code freeze this Thursday)
            * All docs are on hold bc k/k PRs waiting to merge
        * Need help reviewing 50+ doc
            * Half might need tech review (reach out as needed)
        * Celeste - could provide LGTMs
            * Not ready for review
    * (reminder: tentative release Dec 8th)
    * Final deadline is Nov 30th
    * Draft labels / WIP
        * Limit feedback to obvious mistakes (vs. content)
            * Addressing like a feature gate, etc.
        * Nags
        * Anna - Docs team will ping content review
    * Celeste - Retro: assign point for tech review on docs at opening.
    * Integration branch is not healthy after merge (failed build)
    * Auto labels (issue opened for tracking) - two topics need to be discussed.
    * SIG Release - KEPs receipt process, start creating new process (no longer using enhancement tracking sheet - moved to git), [https://docs.google.com/document/d/1qnfXjQCBrikbbu9F38hdzWEo5ZDBd57Cqi0wG0V6aGc/edit#heading=h.j4vyn8gci8l6](https://docs.google.com/document/d/1qnfXjQCBrikbbu9F38hdzWEo5ZDBd57Cqi0wG0V6aGc/edit#heading=h.j4vyn8gci8l6) 
        * KEP of KEPs
            * Arrange a discussion with folks who will be responsible for the KEP
* Blog Update
    * Business as usual, keep these coming! 
* Issues & PRs
    * (awareness) [https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20180731-cloud-provider-docs.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20180731-cloud-provider-docs.md) 
* [Jim] AI: talk about background and demo `krel` if prepared
* [Irvi] Automation some manual process [merged]
    * Issue in k/website: [https://github.com/kubernetes/website/issues/24960](https://github.com/kubernetes/website/issues/24960):
        * Applying hold label automatically against the future release branch
            * ref:[ #24784 (comment)](https://github.com/kubernetes/website/pull/24784#issuecomment-719566601)
        * Change to k/website label strategy (prefix with v) to match other Kubernetes repos
            * ref:[ kubernetes/test-infra#19877 (comment)](https://github.com/kubernetes/test-infra/pull/19877#issuecomment-723462034)
    * PR to automatically apply milestones: [https://github.com/kubernetes/test-infra/pull/19877](https://github.com/kubernetes/test-infra/pull/19877)
    * PR to update docs handbook: [https://github.com/kubernetes/sig-release/pull/1325](https://github.com/kubernetes/sig-release/pull/1325)
* [Brad] Time zone changes US
    * Daylight savings time is bad :(((( 😭😭😭
    * GMT impacts Tim (BST used in Summer)
        * Europe is generally also have this summer time
        * US will end DST a week after Europe :(( even more confusing :(((😭😭
    * AI: Jim will switch the calendar to UTC for APAC events.
    * AI: Brad to announce
* (moved to next meeting) [Celeste] Updates from WG Naming
    * [https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md](https://github.com/kubernetes/community/blob/master/wg-naming/recommendations/master-control-plane.md) is in!
        * L18ns still use master:
            * [https://github.com/kubernetes/website/issues/22580](https://github.com/kubernetes/website/issues/22580) 
            * [https://github.com/kubernetes/community/issues/5043](https://github.com/kubernetes/community/issues/5043) 
    * [https://github.com/kubernetes/website/issues/21749](https://github.com/kubernetes/website/issues/21749) – Why ‘live’? 
            * It’s level with the live content - Tim B
* [Kohei] Japanese first work of 1.18 has finally been merged (phew)
    * Facing difficulties in finding active and regular contributors

**~~10/27/2020~~ Cancelled **

**~~6pm Pacific (01:00 UTC)~~**

~🎉 New contributors 🎉~~



* 

~Updates/reminders~~



* ~~This week’s PR wrangler: @kbarnard10~~
* ~~Next week’s PR wrangler: @onlydole~~
    * ~~Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)~~

~Agenda~~



* ~~Release 1.20~~
    * 
* ~~Blog Update~~
    * 
* ~~Issues & PRs~~
    * ~~(awareness) [https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20180731-cloud-provider-docs.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cloud-provider/20180731-cloud-provider-docs.md) ~~
* ~~KubeCon Talk (introduction to SIG Docs) status~~
    * 
* ~~Cancelled 11/3 weekly meeting for the US election reminder (next week!)~~
* ~~[Jim] AI: talk about background and demo `krel` if prepared~~

**10/20/2020**

**10:30am Pacific (17:30 UTC)**

🎉 New contributors 🎉



* n/a

Updates/reminders



* This week’s PR wrangler: @kbhawkey
* Next week’s PR wrangler: @kbarnard10
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * [Rey] Green
        * Branches are synced
        * Dev-1.20 is healthy
        * Tracking 7 docs
            * 3 merged / 3 ready for review / 1 WIP
        * Placeholders 
* Blog Update
    * n/a
* Localization WG Update
    * 
* Issues & PRs
    * [https://github.com/kubernetes/sig-release/pull/1289/files](https://github.com/kubernetes/sig-release/pull/1289/files) 
* KubeCon Talk (introduction to SIG Docs) status
    * Irvi / Tim / Celeste - recorded a talk
    * 1 week less than planned (oops)
    * October 22nd submission deadline
* Cancelled 11/3 weekly meeting for the US election reminder
* Quarterly (Q4) planning dates
    * Date change: ****Wednesday 10/21**** at 5:00pm Pacific
    * [Agenda Doc](https://docs.google.com/document/d/1RCb8mglnr1MQTy4yAabHaKOIrHf6QnSftCEVPSnqhTI/edit#)
        * Please add anything you’d like to discuss.
* Sig-security-docs sub-project update - Savitha Raghunathan
    * SIG Security met for second time (breakout group)
        * Address certain guidelines / hardening / vulns
        * Goals
    * Savitha Co-leading project, looking for volunteers
        * Ways to connect: sig-security slack channel if interested
        * Future state will be a sig-security-docs channel
    * TODO: archive sig-docs-security slack channel (mention and archive)
    * TODO: share out old sig docs security (jim)

**10/13/2020**

**10:30am Pacific (17:30 UTC)**

🎉 New contributors 🎉



* 

Updates/reminders



* This week’s PR wrangler: @zparnold
* Next week’s PR wrangler: @kbhawkey
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * Somtochi - GREEN!
        * PR merged master (dev-1.20)
            * # issue resolved!
        * Tracking enhancements
* Blog Update
    * n/a
* Localization WG Update
    * Need meetings (AI: Jim)
    * Slack namings / group bengali
    * 
* Issues & PRs
    * [Zach] [https://github.com/kubernetes/website/pull/24149](https://github.com/kubernetes/website/pull/24149) 
    * New triage process - all new issues get tagged “need triage”
        * [https://github.com/kubernetes/enhancements/blob/master/keps/sig-contributor-experience/1553-issue-triage/README.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-contributor-experience/1553-issue-triage/README.md) 
* KubeCon Talk (introduction to SIG Docs) status
    * Doing ok - round one slides are good
    * Need consensus - on recording date / rehearsal times
* Cancelling 11/3 weekly meeting for the US election reminder
* Quarterly (Q4) planning dates
    * (need to update)
    * Thursday 10/22 at 5:00pm Pacific
    * [Q3 doc](https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#heading=h.ghlki2544yp0)
* YouTube playlist fixed
    * (need to update)
    * [Jim] Need to contact someone in the ContribEx about the plugin
    * [https://www.youtube.com/playlist?list=PL69nYSiGNLP3b5hlx0YV7Lo7DtckM84y8](https://www.youtube.com/playlist?list=PL69nYSiGNLP3b5hlx0YV7Lo7DtckM84y8) 

**10/6/2020**

**10:30am Pacific (17:30 UTC)**

🎉 New contributors 🎉



* 

Updates/reminders



* This week’s PR wrangler: @bradtopol
* Next week’s PR wrangler: @Rajakavitha1 - Zach Arnold 
    * Will reach out to confirm
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * Synced dev-1.20 with master
    * [Anna] Extra PR template about the release.
        * Action item:
            * Create a PR template
    * Enhancement freeze (today - tracking docs)
    * Green
* Blog Update
    * n/a
* Localization WG Update
    * [https://docs.google.com/document/d/1NwO1AN8Ea2zlK8uAdaDAKf1-LZDAFvSewIfrKqfl5No/edit](https://docs.google.com/document/d/1NwO1AN8Ea2zlK8uAdaDAKf1-LZDAFvSewIfrKqfl5No/edit)
    * [Irvi] We just performed our first kick-off meeting yesterday, and we learn a lot from the folks. However we still not really sure if the automation to upload to youtube channel works, do we have a mechanism to check this?
        * [Jim] Upload it manually to playlist
* Issues & PRs
    * Hacktoberfest & spam
        * Prow changes
        * [https://hacktoberfest.digitalocean.com/hacktoberfest-update](https://hacktoberfest.digitalocean.com/hacktoberfest-update) 
        * Many thanks to the whole Kubernetes community for helping with triage
        * Tim meeting Digital Ocean on Friday (open source maintainer round table)
    * Need reviewer [https://github.com/kubernetes/website/pull/24035](https://github.com/kubernetes/website/pull/24035) 
* KubeCon Talk (introduction to SIG Docs)
    * Open for suggestions
        * What do you wish you knew before starting?
        * **Process** (who reviews, what are these labels, where do I start)
        * How do I ask for a review?
        * Getting up to speed for non-native speakers
        * Localization subgroup activity
        * **Deadline: need slides uploaded by last week of october**
* [Jim] Quarterly (Q4) planning dates
    * Thursday 10/22 at 5:00 Pacific?
    * [Q3 doc](https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#heading=h.ghlki2544yp0)
* [Jim] YouTube playlist fixed
    * [https://www.youtube.com/playlist?list=PL69nYSiGNLP3b5hlx0YV7Lo7DtckM84y8](https://www.youtube.com/playlist?list=PL69nYSiGNLP3b5hlx0YV7Lo7DtckM84y8) 

**9/29/2020**

**10:30am Pacific (17:30 UTC)**

Meeting mixup



* _Sorted - _use the meeting link at the top of the agenda

🎉 New contributors 🎉



* Somtochi - release docs shadow
* Kristin Martin - release docs shadow
* Nate W. - new CNCF worker / dev advocate for docs
* Ashley Newton - tech writer at salt stack

Updates/reminders



* This week’s PR wrangler: @sftim
* Next week’s PR wrangler: @bradtopol
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * ~~Some progress~~ (Somtochi)
    * Everything’s on track (Anna)
* Blog Update
    * 1 blog article mirroring from k8s.dev (new thing!)
        * [https://k8s.dev/](https://k8s.dev/) publishes first then copy out - more to come
    * Article from CockroachDB
* Localization WG Update
    * Initial group or Subgroup participants identified and Monthly meeting time selected (Brad)
    * Need to get ready for Kick off next monday 11am EST.
    * Need to get access to our zoom account.
    * Need a meeting invite?
        * SIG Docs chairs can assist
    * Brad to create google doc similar to this doc and add meeting and doc details to localization page in Kube Docs
    * Discuss expectations for deliverables of the Subgroup
    * Localization is a large part of SIG Docs effort (Jim)
    * Synchronising to avoid duplicated work / rework will pay off (Jim)
    * (Tim) is keen to find out about localization teams’ pain points
* Kubecon N/A / contributor summit alternative activities:
    * [Irvi] Do we need to somehow polish the slides that we have?
    * Brad has related slides, currently IBM-branded, that can be the basis for a SIG Docs deck with Kubernetes branding.
    * What outcome or call-to-action do we want from this presentation? (Brad)
        * Good practise for technical writing(?)
        * How to start contributing(?)
    * Historically SIG Docs has not been part of the traditional contributor tracks (deep dives etc) (Jim)
    * Questions about who’s producing content, who’s presenting, etc (Jim)
    * Separate meeting? (Jim)
        * When will we do this? (Irvi)
    * Available time to present? (Brad)
    * Pre-recorded contributor summit (Savitha)
        * Savitha & Tim are planning to prepare & record a video on docs contributing (“new contributor workshop”)
        * Let’s co-ordinate the 2 activities (Savitha)
    * Let’s summarise again in Slack so we’re clear on who’s proposing what
    * [Zach] CNCF tech writers also presenting at CloudNativeCon
* Issues & PRs
    * [Jim] [https://github.com/kubernetes/website/issues/20232](https://github.com/kubernetes/website/issues/20232)
        * 3rd party content removal, planned for 1.19 completion
        * What else is needed? (Review open issues)
        * Thoughts on turnkey sections (does any sig own?): [https://github.com/kubernetes/website/issues/21636](https://github.com/kubernetes/website/issues/21636)
    * [Nate W.] [https://github.com/kubernetes/website/issues/23354](https://github.com/kubernetes/website/issues/23354)
        * Should we remove the [minikube installation instructions page](https://kubernetes.io/docs/tasks/tools/install-minikube/) as dual sourced content, and replace it with a link to [minikube’s own installation instructions](https://minikube.sigs.k8s.io/docs/)?
            * [Tim] Redirect the page
            * [Tim] Possibly identify any text worth copying and pasting into Minikube (as a PR to the Minikube docs)
            * [Karen B] tstromberg opened a PR to trim that page, looks relevant [https://github.com/kubernetes/website/pull/20717](https://github.com/kubernetes/website/pull/20717)
    * [Jim]
        * [https://github.com/kubernetes/website/issues/24170](https://github.com/kubernetes/website/issues/24170) 

Process



* [Jim] Quarterly (Q4) planning coming up in October / November
    * TBD
    * [Q3 doc](https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#heading=h.ghlki2544yp0)
* [Jim] Need to fix our YouTube playlist
    * Jim will add link to playlist for convenience

**9/22/2020**

**6pm Pacific (01:00 UTC)**

New contributors



* New new

Updates/reminders



* This week’s PR wrangler: @vineethreddy02
* Next week’s PR wrangler: @xiangpengzhao (Peter Zhao)
    * AI: Jim to confirm / reach out to peter
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20
    * Shadow selection is complete, Anna is working on onboarding them to the release team
    * If you have any questions/concerns about 1.20 release, please feel free to Anna or reach out at #release-docs channel
* Blog Update
    * Caught up on content
    * Opening blog shadowing
        * If interested reach out to Kaitlyn ot sig-docs-blog
* Co-chair transition
    * [Co-chair transition · Issue #23797 · kubernetes/website · GitHub](https://github.com/kubernetes/website/issues/23797)  
    * Lazy consensus / vote
    * Zach: In future - open from non transitioning chair so transitioning chair can “approve” out
        * AI: Irvi or someone else open PRs listed in issue
* Issues & PRs
    * [https://github.com/kubernetes/website/issues/20862](https://github.com/kubernetes/website/issues/20862) 
        * [https://github.com/kubernetes/website/pull/23497](https://github.com/kubernetes/website/pull/23497)
        * “Issue conversation” vs “PR” conversation
        * Lazy consensus to merge in and iterate forward
        * Add label and merge
* [Tim B] (missing the meeting) Zoom passcode update; see [k/community#5146](https://github.com/kubernetes/community/issues/5146)
    * AI: Jim to update meeting links 
* [Celeste] Update about [Kubernetes Contributors](https://www.kubernetes.dev/)
    * [Celeste to update] [https://github.com/kubernetes/website/issues/23417](https://github.com/kubernetes/website/issues/23417) 
        * Make sure that the meetup content and CoCC are on kubernetes.dev and assign it to Rin Oliver 
    * Chris A. defers to us
    * K8s.dev exists (community page conflicts)
        * Code of conduct
        * Meetups (feed) - might need to find a home for that
    * CNCF employees are the only folks maintaining
    * Ray: Make sure there are not “load bearing” release instructions
* [Seokho] Korean l10n update
    * Membership update
        * New reviewer for Korean contents: @pjhwa (total 6 members)
    * Team Milestones
        * Finished Ko l10n for release-1.18 (dev-1.18-ko.1 - dev-1.18-ko.11)
            * dev-1.18-ko.11 is a branch to update outdated contents (merged to release-1.18 branch)
        * Working on 1.19 master (dev-1.19-ko.2 will be merged soon!)
        * Looking for a method to notify team milestone to GitHub-only-contributors
            * Testing the GitHub Milestone with a help-wanted Issue
                * [https://github.com/kubernetes/website/issues/23835](https://github.com/kubernetes/website/issues/23835) 
                * [https://github.com/kubernetes/website/milestone/47](https://github.com/kubernetes/website/milestone/47) 
* [Zach] Community norms for approving our own PRs
    * Because PRs to k/website have a high impact, it’s good to have an additional/explicit approval review. 
    * TLDR: don’t approve your own PRs
    * Zach: There may be a need to recruit more technical website reviewers
    * Need to get functionality review / check for breaking changes.
    * Raise to community meeting for visibility (SIG Contribex monthly update)
        * AI: Celeste someone to raise at next meeting
    * Qiming: Owner Alias file, team owners, only half team is active. Random assignment to inactive reviewers.
        * Comment some names out - but what is the criteria?
            * Stated expectation: reviewers must remain active (as short as 30 days or as long as 90 days)

**9/15/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* welcome back Adam Kaplan
* Welcome @rghav

Updates/reminders



* This week’s PR wrangler: @tengqm (Qiming) 
* Next week’s PR wrangler: @vineethreddy02 (Vineeth)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20 updates
    * Slack channels
        * [https://kubernetes.slack.com/messages/sig-docs-release](https://kubernetes.slack.com/messages/sig-docs-release)
        * [https://kubernetes.slack.com/messages/release-docs](https://kubernetes.slack.com/messages/release-docs) (_selected_)
            * Deprecated unselected channel(?)
        * Use public channel by preference
    * Access control [https://github.com/kubernetes/website/pull/23905](https://github.com/kubernetes/website/pull/23905) need approval
* Chair transition
    * Zach becomes emeritus on September 15th, nominates @irvifa to replace him
    * [https://github.com/kubernetes/website/issues/23797](https://github.com/kubernetes/website/issues/23797) 
* PRs and issues 
    * [Celeste/Tim] [https://github.com/kubernetes/website/pull/23210](https://github.com/kubernetes/website/pull/23210) – needs an /approve, this is an implementation change
    * [Celeste] [https://github.com/kubernetes/website/pull/23708](https://github.com/kubernetes/website/pull/23708) – I need help reviewing
* Localization sub-group:
    * SIG Docs Localization update
        * [https://github.com/kubernetes/website/issues/22960](https://github.com/kubernetes/website/issues/22960) 
    * [Irvi] several challenges identified
        * Can we reuse scripts & tools (probably!)
            * Some tooling won’t apply to all localizations
        * Reducing workload and toil for localization contributors
    * [Brad] Mission statement 
    * _The mission of the SIG Docs localization subgroup is to work across the SIG Docs localization teams to collaborate on defining and documenting the processes for creating localized contribution guides. In addition, the SIG Docs localization subgroup will look for opportunities for the creation and sharing of common tools across localization teams and also serve to identify new requirements to the SIG Docs Leadership team._
    * **Action on Brad** to announce the subgroup
        * After announcing the subgroup we can define periodic meetings? 
* Process
    * On September 27th, Zoom will be rolling out a change that will force all of our meetings to use either a waiting room or a passcode for people to enter. If you are prompted for a passcode - use five sevens: **77777**
        * Spread the word, and if you happen to see someone confused about the new passcode please help them out!
* [Tim] KubeCon talks etc
* [Tim] [https://www.kubernetes.dev/](https://www.kubernetes.dev/) - how to integrate?
    * [Celeste] follow-on: SIG Contribex sees [https://kubernetes.io/community/](https://kubernetes.io/community/) as superfluous now
    * [Tim] Are we maintaining the existing community page well?
    * [Tim] Does [https://kubernetes.io/community/](https://kubernetes.io/community/) serve more than developers?
        * check Google Analytics?
        * **Action on Celeste** to reach out to @oicheryl for another opinion
    * [Celeste] How do we hyperlink to the code of conduct?
    * [Irvi] How is the new [https://www.kubernetes.dev/](https://www.kubernetes.dev/) maintained?
    * [Anna] Join SIG ContribEx weekly meeting to liaise? [Irvi] via Slack
        * **Action on Celeste **to liaise with SIG ContribEx
    * [Tim] How does this fit in with [https://kubernetes.io/docs/contribute/](https://kubernetes.io/docs/contribute/) ?
* [Anna] Where are the weekly meeting video recordings?

**9/8/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* (returning) Matt Boersma

Updates/reminders



* This week’s PR wrangler: @onlydole (Taylor)
* Next week’s PR wrangler: @tengqm (Qiming) 
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.20 updates (if any)
    * Milestone maintainers access for shadows - Anna/Savitha
    * Next week to discuss retro for Docs / improvements
* Chair transition
    * Zach becomes emeritus on September 15th, nominates @irvifa to replace him
        * Zach will open a PR like [https://github.com/kubernetes/website/issues/18117](https://github.com/kubernetes/website/issues/18117) and do a hound search for himself: [https://cs.k8s.io/?q=zacharysarah&i=nope&files=OWNERS*&repos=](https://cs.k8s.io/?q=zacharysarah&i=nope&files=OWNERS*&repos=)
* PRs and issues
    * [Tim] Serve a HTTP 410 instead of a TLS error: [https://github.com/kubernetes/website/issues/12303](https://github.com/kubernetes/website/issues/12303) 
    * [Tim/ Philippe Martin] Redesign of API reference docs

        WIP PR [https://github.com/kubernetes/website/pull/23294](https://github.com/kubernetes/website/pull/23294)


**9/1/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* 

Updates/reminders



* This week’s PR wrangler: @CelesteHorgan (subbing for Savitha)
* Next week’s PR wrangler: @Onlydole (Taylor) 
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.19/1.20 updates (if any)
* Blog
* Issues & PRs
    * [Zach/Taylor] Client libraries: [https://github.com/kubernetes/website/pull/21509#issuecomment-679718381](https://github.com/kubernetes/website/pull/21509#issuecomment-679718381)
        * Waiting on, 3rd Party Warning: [https://github.com/kubernetes/website/pull/23107](https://github.com/kubernetes/website/pull/23107) 
        * Then I (Taylor) will add the warning and rebase the branch
        * Celeste is going to be wrapping up this PR and closing mine (thank you!) 
    * [Zach/Jim] Better CKA experience: [https://github.com/kubernetes/website/pull/22973/](https://github.com/kubernetes/website/pull/22973/files)
        * How do we want to proceed here in the context of this issue: [https://github.com/kubernetes/website/issues/22961](https://github.com/kubernetes/website/issues/22961)
* [Zach] Google Season of Docs
    * [https://developers.google.com/season-of-docs](https://developers.google.com/season-of-docs) 
    * Action Item: Let’s add to our SIG Docs calendar
* [Taylor] Remove restriction on wiki
    * Currently is: Restrict editing to users in teams with push access only Public wikis will still be readable by everyone.
    * Proposal is to uncheck this to make updates to wiki a bit easier for our team
    * @Brad - please contact for any wiki changes
* [Irvi] Last KubeCon some of people asking if the SIG Docs will be having session in KubeCon NA
    * Small group needs to be willing to create a talk or workshop (might’ve slipped through the cracks with all going on lately)
    * Concern on time zone and best time(s) to meet
    * Who will be able to jump in: Celeste, Brad, Irvi, other folks will be very welcome(?)
        * Write 101 tech writer
    * Action item: talk to the program management committee/organize (@Taylor to file GitHub issue/Slack Zach/Jim)
* [Irvi] For this issue regarding link checker, does anyone knows on how we can put the result directly on the PR (I don’t know how we can integrate this with Prow somehow and I don’t think it’s also supported):
    * [https://github.com/kubernetes/website/issues/22489](https://github.com/kubernetes/website/issues/22489) 
    * [https://github.com/kubernetes/website/pull/22923](https://github.com/kubernetes/website/pull/22923) 
    * Action item: @Taylor to help collaborate with these tickets
* [Brad] Kicking off sub-group for localization
    * 

**8/25/2020**

**6pm Pacific (01:00 UTC)**

New contributors



* Sujay Pillai (Malaysia) 

Updates/reminders



* This week’s PR wrangler: @makoscafee
    * [Zach: I haven’t seen any activity from Barnie in a while; I can sort-of wrangle this week, but let’s follow up with him and see what Barnie’s thoughts are about continuing as an approver]
* Next week’s PR wrangler: Scheduled to be @savitharaghunathan, but she needs a break after the release!
    * Celeste will take the shift!
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.19
* Blog
* Issues & PRs
    * [Zach/Taylor] Client libraries: [https://github.com/kubernetes/website/pull/21509#issuecomment-679718381](https://github.com/kubernetes/website/pull/21509#issuecomment-679718381) 
    * [Zach/Jim] Better CKA experience: [https://github.com/kubernetes/website/pull/22973/](https://github.com/kubernetes/website/pull/22973/files)
        * How do we want to proceed here in the context of this issue: [https://github.com/kubernetes/website/issues/22961](https://github.com/kubernetes/website/issues/22961)
    * [Zach/Karen] Moving search: [https://github.com/kubernetes/website/pull/23083](https://github.com/kubernetes/website/pull/23083)
    * [Zach/Jim] Submodule update: [https://github.com/kubernetes/website/pull/23434](https://github.com/kubernetes/website/pull/23434)
* [Zach] Redesign of API reference docs starts on 14 September with Google Season of Docs writer, Philippe Martin (@feloy on GitHub)
* Persian language localization init in progress: [https://github.com/kubernetes/website/pull/22565](https://github.com/kubernetes/website/pull/22565) 
* [Zach] Style guide additions: avoid colloquial language; define abbreviations on first use
    * Shortcode / html
* [Jim] Localization analytics and improvements still in progress
    * Only measure things prepared to act on. Remove feedback (yes / no) or measure and act.
* [Celeste] Editing w/ Divya

**8/18/2020**

**10:30am Pacific (17:30 UTC)**

_Skipped for Kubecon 2020 EU_

**8/11/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* 

Updates/reminders



* This week’s PR wrangler: @kbarnard10
* Next week’s PR wrangler @zacharysarah
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.19
    * Also see issue 23031 below
    * Branch sync in progress - [https://github.com/kubernetes/website/pull/23075](https://github.com/kubernetes/website/pull/23075)
    * Update API generation in progress - [https://github.com/kubernetes/website/pull/23018](https://github.com/kubernetes/website/pull/23018)
    * Dev-1.19 branch health (https://github.com/kubernetes/website/pull/20785) - Red
* Blog
* _(from previous week)_ Guidelines for promotion of localization thresholds
    * [Jim] I thought this would be good to bring up from the APAC meeting
    * When / how to promote localization contributors
    * [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)
    * Create localization subgroup
        * Invite stakeholders
        * Publicize (also mention “diff script”?)
            * If the diff is large probably need to do that periodically and see the analytics for prioritization
            * AI: Zach will do admin pieces for subproject formation
* _(from previous week)_ CKS~~S~~ Qualification [Tim]
    * Adding CKS~~S~~ to the training section on docs
    * Potentially create good first issue (clearly defined for a new contributor)
    * Related CKA
* Issues & PRs
    * [https://github.com/kubernetes/website/issues/23031](https://github.com/kubernetes/website/issues/23031) [Tim] ([Celeste] can speak to this a bit, i think)
    * [https://github.com/kubernetes/website/issues/23032](https://github.com/kubernetes/website/issues/23032) (deprecation warning fix)
        * [Celeste] Need to clarify testing criteria before testing
        * Add comment on issue
    * [https://github.com/kubernetes/website/issues/21797](https://github.com/kubernetes/website/issues/21797) - can we pick an approach? [Tim]
        * [Celeste] How about creating a poll in the slack
        * Tim will create a poll (48 hours) in slack and see folks preference
    * [https://github.com/kubernetes/website/issues/22962](https://github.com/kubernetes/website/issues/22962) [Celeste] – [let’s talk wording asynchornously](https://github.com/kubernetes/website/issues/22962#issuecomment-672131589)
        * Zach will review by Friday EOD 14 August
    * [https://github.com/kubernetes/website/pull/22923](https://github.com/kubernetes/website/pull/22923) [Irvi] - This is tested for push and PR trigger. Need review from the others.
        * Sample report [https://github.com/irvifa/website/runs/972373005?check_suite_focus=true](https://github.com/irvifa/website/runs/972373005?check_suite_focus=true) 
        * Improve discoverability / review
* Hugo & Docsy in a container for docs [Jim / Celeste / Joel]
    * npm / yarn [https://github.com/kubernetes/website/pull/22995](https://github.com/kubernetes/website/pull/22995)
        * docsy wants npm
        * +1 - Jim will drive with inductor 
* [Jim] Analytics - assuming I can, any reason not to give anonymous view / read access?
* [Brad] discuss feedback I got from Poland, France, and Italy
    * [Zach] Are they aware of the diff script in main repo
    * [Irvi] Remy already uses a diff script, but sometimes the diff is huge. How to pick and choose what is enforced / to prioritize for updates.
    * Toil in correlating changed documents with analytics.
        * Any possibility of integrating analytics and the diff(?)
* (Add to agenda for 8/18) SIG Security liaison [Tim]

**8/4/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* Abby McCarthy - returning after a while

Updates/reminders



* This week’s PR wrangler: @kbhawkey
* Next week’s PR wrangler @kbarnard10
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.19
    * (not able to attend): Status is green. Dev-1.19 is healthy
* Blog
    * Tim - has one open article
    * 1 other needs review (Kevin?) from blog team - maybe turn into case study?
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/22799](https://github.com/kubernetes/website/pull/22799)
        * Add to tracking issue?
            * Issue to clean up how the CNI information is presented (and inlign with the KEP)
                * Does everything need a GitHub repo / be OSS?
                    * No [Tim]: SIG-Windows
        * Is this CNI provider up to date and installable a cloud native way?
        * Does this product align?
            * Review - does it align as a CNI? (consensus: NO, not CNI)
        * Outcome: open issue to clean up AI: Jim
            * Tasks / UI improvement (tag Tim)
        * Outcome (part2): Publicly define 3rd party content and call out page and requirements. AI: Celeste
    * CKA CNI Kubeadm docs [Jim]
        * [https://github.com/kubernetes/website/pull/22600#event-3587871540](https://github.com/kubernetes/website/pull/22600#event-3587871540)
* Hugo & Docsy in a container for docs [Jim / Celeste / Joel]
    * Updated makefile and continuing work
    * [Kohei] README suggests yarn while the repo has package-lock.json but not yarn.lock. What’s the goal for the npm package management(inside or outside of the container)?
        * [Tim] We should use what Netlify uses (package.json)
        * [Kohei] had issues with npm ci, yarn worked, will double check
        * [Celeste] why is yarn here?
        * [Kohei]: opening issue and troubleshooting / tagging folks interested
* Localization struggles to keep up localization pages when things change.
    * [Jim] I thought this would be good to bring up from the APAC meeting
    * No great answer
    * Recommended reaching out to ko localization
    * AI: bring up at tuesday meeting
    * [Kohei]: not standardized, each localization has their own standard / team. Independent. (Can’t really force / make work a certain way). 
        * KO is doing things very well (milestones)
        * JA is automating tracking the diff (git commands)
        * Tips can be shared
    * [Brad] ZH / KO have tools (discussed at KubeCon shanghai)
        * Circle back to see if there’s ability for reuse
    * [Jim] Can we create a universal tool that shows diffs?
    * [Kohei] localization docs lack information on maintaining a repo (after getting started). KO forked this page and created their own with better information (branch strategy / maintain / and how-to translate [whole vs. half]).
        * Needs improvement for making own pages
    * [Tim] Normal flow is English -> Translation (good rule). For a localization page, maybe that rule bends.
    * [Tim] Anyone can jump into a localization channel
    * AI: Jim will open issue (scope of improving localization ongoing guides) but needs other folks to drive

Moved to the following week:



* Guidelines for promotion of localization thresholds
    * [Jim] I thought this would be good to bring up from the APAC meeting
    * When / how to promote localization contributors
    * [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)
* CKS~~S~~ qualification [Tim]

**7/28/2020**

**0100 UTC on 2020-07-29 (6pm Pacific on 7/28/2020)**



* Ping APAC channels before starts
* Boost in mailing list

New contributors (returning)



* Ippei Suzuki - Creation Line (CNCF)
* Giri Kubcoro - ID translation
* Aris Risdianto - ID translation

Updates/reminders



* This week’s PR wrangler: @jimangel
* Next week’s PR wrangler @kbhawkey
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.19
    * all PR's have been merged and the status is green. In the latest branch synch performed by [@sraghunathan](https://kubernetes.slack.com/team/UC8U2V3BM), there were no merge conflicts. Cc: [@sraghunathan](https://kubernetes.slack.com/team/UC8U2V3BM) [@annajung](https://kubernetes.slack.com/team/U8SLB1P2Q) [@mikjoh](https://kubernetes.slack.com/team/UFNRWH15H) [@zestrella](https://kubernetes.slack.com/team/ULTV3SWR2)
* Korean l10n update
    * n/a
* Blog
    * n/a
* Issues & PRs
    * PR health is good
* Hugo & Docsy in a container for docs [Jim / Celeste / Joel]
    * Updated makefile and continuing work
* Localization struggles to keep up localization pages when things change.
    * No great answer
    * Recommended reaching out to ko localization
    * AI: bring up at tuesday meeting
* Japanese localization
    * Catching up with things is ja channel
    * (localizing training at the moment)
* Guidelines for promotion of localization thresholds
    * [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)

**7/21/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* 

Updates/reminders



* This week’s PR wrangler: @CelesteHorgan
* Next week’s PR wrangler @zacharysarah
    * Zach on vacation next week, needs to trade
    * Jim will wrangle!
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* 7/28 meeting: Need someone to lead the APAC meeting
    * Jim will lead :)
        * Ping APAC channels before starts
        * Boost in mailing list

Agenda



* Release 1.19
    * Green
    * Release team rocks!
* Blog
    * Business as usual
    * Removed form, PRs come strait through GitHub
    * [https://kubernetes.io/docs/contribute/new-content/blogs-case-studies/](https://kubernetes.io/docs/contribute/new-content/blogs-case-studies/)
* Issues & PRs
    * Persian docs: [https://github.com/kubernetes/website/pull/22565](https://github.com/kubernetes/website/pull/22565)
    * Periodic Link Checker : [https://github.com/kubernetes/test-infra/pull/18399](https://github.com/kubernetes/test-infra/pull/18399)
        * How to clone the repo in periodic jobs. I notice that in presubmit jobs we only need to mention the orgz and repo’s name. Not sure if we need another auth for cloning the repo before doing the link check. On the other hand I notice that for each release we’re doing periodic jobs to detect flaky tests etc. 
            * Follow up with @spiffxp, @alejandrox1
        * Need help on someone familiar with testgrid. Most periodic jobs are reported on testgrids. 
            * Follow up with @cblecker
* Hugo & Docsy in a container for docs [Jim / Celeste / Joel]
    * It’s messy
    * [https://github.com/kubernetes/website/issues/22515](https://github.com/kubernetes/website/issues/22515) 
    * [https://github.com/kubernetes/website/issues/22586](https://github.com/kubernetes/website/issues/22586) 
    * [https://github.com/kubernetes/website/pull/22518](https://github.com/kubernetes/website/pull/22518) 

**7/14/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* Joel Barker

Updates/reminders



* This week’s PR wrangler: @CelesteHorgan
* Next week’s PR wrangler @CelesteHorgan
    * Jim, Kaitlyn, let’s open a PR to remove @daminisatya and @rajakavitha1 as approvers
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* 7/28 meeting: Need someone to lead the APAC meeting
    * Jim will lead :)

Agenda



* Release 1.19
    * Pool of reviewers for docs
    * Status: RED :(
* Quarterly meeting last week: Review the notes: [https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#heading=h.ghlki2544yp0](https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#heading=h.ghlki2544yp0)
* Blog
    * 
* Issues & PRs
    * [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748)
        *  +Zach pinged SIG API Machinery in Slack, no response
    * [https://github.com/kubernetes/website/issues/22489](https://github.com/kubernetes/website/issues/22489)
        * +Irvi Testing Prow command and add PR for the config
    * Localization friendly link:
        * Is there any possibility of using Hugo config as possible improvement since this will mean that we do the conversion in Hugo. This is discussed in [https://github.com/kubernetes/website/pull/19526](https://github.com/kubernetes/website/pull/19526)
        * A PR for a script to do conversion is already included in [https://github.com/kubernetes/website/pull/21996#pullrequestreview-448313571](https://github.com/kubernetes/website/pull/21996#pullrequestreview-448313571) 
* Hugo & Docsy in a container for docs [Jim / Celeste / Joel]

**7/7/2020**

**10:30am Pacific (17:30 UTC)**

New contributors



* 

Updates/reminders



* No scheduled PR wrangler this week (Thanks @sftim!)
* Next week’s PR wrangler @daminisatya
    * Zach C: I suspect we’ll need to find replacements for next week and the week following
    * UPDATE: @CelesteHorgan will wrangle next week and the week following, thanks Celeste!
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* Quarterly review: Thursday, July 9th at 5pm Pacific (00:00 UTC Friday, July 10th)
    * Zach set up a calendar invite
    * [https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#](https://docs.google.com/document/d/1VZ2pgZKQ-v4SsjVh-IqXz6nNLhxcCK45Mp_cdwrFxa8/edit#)

Agenda



* Over 1M unique site visitors last month!
* Release 1.19
    * Docs deadline: July 16th
    * Status: yellow
    * 6 PRs in progress from feature developers, contacting authors to work with them
    * Merged two docs PRs!
    * Issues with the last branch sync merge PR: all addressed, needs approval
* Quarterly review upcoming
    * Thursday, July 9th at 5:01pm Pacific (0001 UTC Friday 10 July)
    * Propose topics for discussion: 
* Blog
    * A bit of an increase in PRs in progress, but should be cleared up after this week.
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/22218](https://github.com/kubernetes/website/pull/22218)
    * [https://github.com/kubernetes/website/issues/21797](https://github.com/kubernetes/website/issues/21797)
        * @sftim to clarify
    * [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748)
        *  +Zach will talk to SIG API Machinery and ask what they’d like to do

**6/30/2020**

**17:30 UTC on 2020-06-30 **

New contributors

Updates/reminders



* This week's PR wrangler is @zparnold
* Next week’s PR wrangler is US Holiday Break (July 4th), with best-effort help from @sftim
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* Quarterly review: Thursday, July 9th at 5pm Pacific (00:00 UTC Friday, July 10th)
    * Zach set up a calendar invite
* Next deadline is July 16th

Agenda



* Release 1.19
    * Green (2 prs merged)
    * Merge issue with PR (invalid commit)
    * Get netlify access to Savitha/release lead
* Quarterly review upcoming
    * Thursday, July 9th at 5:01pm Pacific (0001 UTC Friday 10 July)
    * Proposal: [https://github.com/kubernetes/website/issues/22024](https://github.com/kubernetes/website/issues/22024) [Celeste]
    * Proposal #2: Moderation of some sort for CSS/JS/etc. [Celeste]
* Blog
    * All good on this front, looking for more engagement around PRs
    * Jorge Castro has a PR in for revised Blog guidelines (take a peek if you can)
        * [https://github.com/kubernetes/website/pull/22176](https://github.com/kubernetes/website/pull/22176)
* WG Naming
    * [https://github.com/kubernetes/community/pull/4884](https://github.com/kubernetes/community/pull/4884) [Celeste]
* Issues & PRs
    * [https://github.com/kubernetes/website/issues/21563](https://github.com/kubernetes/website/issues/21563) (@sftim / 3rd party content)
* Future chair transition
    * 

**6/23/2020**

**0100 UTC on 2020-06-24 (6pm Pacific on 6/23/2020)**

New contributors

Updates/reminders



* This week's PR wrangler is @sftim
* Next week’s PR wrangler is @zparnold
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* Quarterly review: Thursday, July 9th at 5pm Pacific (00:00 UTC Friday, July 10th)
    * Zach will set up a calendar invite

Agenda



* Release 1.19
    * Some commits with failed tests but everything still looks healthy
* Quarterly review upcoming
    * Thursday, July 9th at 5:01pm Pacific (0001 UTC Friday 10 July)
        * Client libraries
* Blog
    * No updates
* Ko l10n update.
    * Membership update: No
    * Team Milestones
        * One team milestone branch has been merged into master (dev-1.18-ko.4)
        * Another milestone branch merge planned for tomorrow (which is big to stay up with docsy theme)
* WG Naming
    * Draft complete, official PR pending
* Issues & PRs
    * [https://github.com/kubernetes/website/issues/21687](https://github.com/kubernetes/website/issues/21687) (@zacharysarah / Reorganize Advanced Contributing, Participation in SIG Docs to reduce duplication)
    * [https://github.com/kubernetes/website/issues/21563](https://github.com/kubernetes/website/issues/21563) (@sftim / 3rd party content)
* Future chair transition
    * Zach is transitioning away, goal is to become chair emeritus by Labor Day (September 7th 2020)
    * Not going _away _away, will still be a resource for advice, legacy info, and mentorship
    * Recommend keeping at least three chairs, happy to train my replacement but need help identifying one--preferably a technical writer, but not a hard requirement

**6/16/2020**

**17:30 UTC (10:30am Pacific)**

New contributors



* Lots of friendly familiar faces :-)

Updates/reminders



* This week's PR wrangler is @tengqm
* Next week’s PR wrangler is @sftim
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
* Quarterly review: Thursday, July 9th at 6pm Pacific (01:00 UTC Friday, July 10th)
    * Zach will set up a calendar invite

Agenda



* Release 1.19
    * Milestones delayed
    * Docs status; yellow: not all placeholder PRs open
        * Follow up with SIG owners to get visibility on unopened PRs
* Docsy theme:  🎉merged! 🎉
* Blog: 
    * New template soon-ish
* WG Naming
    * [https://groups.google.com/g/kubernetes-dev/c/kry8QbIpxRs/m/xaWKZYUGBgAJ?pli=1](https://groups.google.com/g/kubernetes-dev/c/kry8QbIpxRs/m/xaWKZYUGBgAJ?pli=1)
    * Related issue: [https://github.com/kubernetes/website/issues/21621](https://github.com/kubernetes/website/issues/21621)
* Improved contributor guide:
    * Celeste
    * [https://github.com/kubernetes/website/issues/21687](https://github.com/kubernetes/website/issues/21687) 
* Content guide
    * Blog post later in 2020-06 / early 2020-07 (@sftim)

**6/09/2020**

**17:30 UTC (10:30am Pacific)**

New contributors



* Anne Ulrich - wants to apply to Google Season of Docs!

Updates/reminders



* This week's PR wrangler is @jimangel
* Next week’s PR wrangler is @tengqm
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * [Savitha] - Status: Yellow
        * Potentially pushing the release back by 2 weeks. If so Docs deadline is moved by a week. 
            * Needs help with an invalid commit message – [https://github.com/kubernetes/website/pull/21604](https://github.com/kubernetes/website/pull/21604) (Change the Prow rules long term?)
                * Savitha following up with Test Infra
    * Docsy theme updates
        * [https://github.com/kubernetes/website/pull/20874](https://github.com/kubernetes/website/pull/20874) 
            * What happens when this is ready to merge?
    * Blog
        * 
    * Any PRs needing review / attention?
        * [Remove API client-libraries section](https://github.com/kubernetes/website/pull/21509)
            * What’s the criteria for entry/removal from this list? Does this fall under our third-party KEP?
            * Celeste: Potentially add in a warning label of sorts instead of remove. (list them in alphabetical order).
        * [Replace blacklist to blocklist](https://github.com/kubernetes/website/pull/21595)

            [Changed whitelist to ignorelist and blacklist to denylist](https://github.com/kubernetes/website/pull/21591)  


            [replace whitelist to allowlist](https://github.com/kubernetes/kubernetes/pull/91927)

* Preferred wording? 
        * Scheduling & Eviction [https://github.com/kubernetes/website/issues/19081](https://github.com/kubernetes/website/issues/19081)
            * Adam: Separate Sched/Eviction & make more conceptually clear. Discuss @ quarterly review? 
        * Architectural diagrams: - David
            * [https://github.com/kubernetes/website/pull/20767](https://github.com/kubernetes/website/pull/20767)
            * [https://kubernetes.io/docs/concepts/overview/components/](https://kubernetes.io/docs/concepts/overview/components/)
            * [https://deploy-preview-20767--kubernetes-io-master-staging.netlify.app/docs/concepts/overview/components/](https://deploy-preview-20767--kubernetes-io-master-staging.netlify.app/docs/concepts/overview/components/)
    * It’s time for a quarterly review!
* ???
* Community

**5/26/2020**

**6pm Pacific**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @onlydole
* Next week’s PR wrangler is @makoscafee
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * Week into enhancement freeze look for incoming docs
        * Savitha merging master into dev-1.19 regularly
    * Docsy theme updates
        * Theme merge PR: [https://github.com/kubernetes/website/pull/20874](https://github.com/kubernetes/website/pull/20874)
        * updated the "actionable bugs" list for the Docsy conversion:[ https://github.com/kubernetes/website/pull/20874#issuecomment-628411032](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fwebsite%2Fpull%2F20874%23issuecomment-628411032&v=3)
        * Open PR to fix / remove capture issues (Karen)
            * [https://github.com/kubernetes/website/pull/20977](https://github.com/kubernetes/website/pull/20977)
    * Link Checker [Celeste]
        * [https://github.com/kubernetes/website/issues/20607](https://github.com/kubernetes/website/issues/20607)
        * [https://github.com/kubernetes/website/pull/20606](https://github.com/kubernetes/website/pull/20606) 
            * After completion, Celeste will update Slack channel with next steps.
    * Blog [Kaitlyn]
        * In search of more technical and editorial blog editors. Technical blog editors should have deep knowledge of the Kubernetes ecosystem. Editorial blog editors do not need to be as experience with Kubernetes, but should still have experience with git/GitHub
        * Needing help in the editorial role!
        * Shadow application can be filled out here: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows)
    * Any PRs needing review / attention?
        *      
    * Monthly Korean l10n update
        * Membership update: No
        * Team Milestones
            * One team milestone branch has been merged into master (dev-1.18-ko.3)
            * Another milestone branch merged planned for tomorrow                  
* Community
    * Note for chairs / tech leads: governance update to review (in mailing list )
    * 

**5/19/2020**

**10:30am Pacific**

New contributors



* Subod - Cisco 
* Zachary - returning (release shadow)

Updates/reminders



* This week's PR wrangler is @kbarnard10
* Next week’s PR wrangler is @zacharysarah
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
    * Updated for the rest of the year (thanks Brad!)

Agenda



* Docs
    * Release update 1.19
        * _release-1.18_ branch was deleted, to be created closer to release
        * Divya - overall status = green
            * Starting tracking on docs
            * Enhancement freeze is today!
    * Docsy theme updates
        * Theme merge PR: [https://github.com/kubernetes/website/pull/20874](https://github.com/kubernetes/website/pull/20874)
        * Feedback should go in the comment section in this PR
            * Scope guidance: [https://github.com/kubernetes/website/pull/20874#issuecomment-628411032](https://github.com/kubernetes/website/pull/20874#issuecomment-628411032) 
        * Related impact (hugo upgrade) - tracking issues from that
    * Link Checker [Celeste]
        * [https://github.com/kubernetes/website/issues/20607](https://github.com/kubernetes/website/issues/20607)
            * Addressing remaining comments
            * After completion, Celeste will update Slack channel with next steps.
    * Diagrams [David]
        * [https://github.com/kubernetes/website/pull/20767](https://github.com/kubernetes/website/pull/20767)
            * Looking for additional help if interested, reach out to david.kypuros / Slack SIG-Docs channel.
            * Zachary Estrella interested
    * Blog [Kaitlyn]
        * In search of more technical and editorial blog editors. Technical blog editors should have deep knowledge of the Kubernetes ecosystem. Editorial blog editors do not need to be as experience with Kubernetes, but should still have experience with git/GitHub
        * Needing help in the editorial role!
        * Shadow application can be filled out here: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows)
    * Any PRs needing review / attention?
        *                                   
* Community

**5/12/2020**

**10:30am Pacific**

New contributors



* Welcome back release shadows!

Updates/reminders



* This week's PR wrangler is @kbhawkey
* Next week’s PR wrangler is @kbarnard10
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * Updates from lead / shadows
            * Anna: Everything is green! 
        * Removal of the release branch is a “go”
    * Docsy theme updates
        * Theme merge PR: [https://github.com/kubernetes/website/pull/20874](https://github.com/kubernetes/website/pull/20874)
        * Umbrella issue: [https://github.com/kubernetes/website/issues/20344](https://github.com/kubernetes/website/issues/20344)
        * Close 20344 and open a new tracking issue for post-20874
    * Link Checker [Celeste]
        * [https://github.com/kubernetes/website/issues/20607](https://github.com/kubernetes/website/issues/20607)
        * Revisit on TODAY
    * Diagrams [David]
        * PR: [https://github.com/kubernetes/website/pull/20767](https://github.com/kubernetes/website/pull/20767)
        * Other feedback welcome
    * Blog [Kaitlyn]
        * In search of more technical and editorial blog editors. Technical blog editors should have deep knowledge of the Kubernetes ecosystem. Editorial blog editors do not need to be as experience with Kubernetes, but should still have experience with git/GitHub
        * Shadow application can be filled out here: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows)
    * Any PRs needing review / attention?
        * Issues/PRs related to [upgrading Hugo to 0.70.0](https://github.com/kubernetes/website/pull/19907)                                         
* Community

**5/5/2020**

**10:30am Pacific**

New contributors



* Anna Jung (docs shadows)
* Mikael Johansson (docs shadow)
* Divya Mohan (docs shadow)

Updates/reminders



* This week's PR wrangler is @zacharysarah
* Next week’s PR wrangler is @kbhawkey
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * Deletion of release-1.xx (current) branch
            * **Context:** A release branch is cut at the end of a release that mirrors master so Netlify can be set up and configured. This branch is never used except when the NEXT release is cut (9+ months).
            * **Proposal:** delete the release-1.18 (current) branch and [Jim] update the handbook with instructions about cutting it ~1 week or days before the release + Netlify setup.
            * **Pros: **
                * Reduces ability for branch to be used incorrectly (causing conflicts)
                * Stops release team from pointless merge-syncs
            * **Cons:**
                * Adds a little more complexity to an already busy “end of cycle” for docs release team.
                * Prevents future setting up of Netlify … kinda ¯\_(ツ)_/¯
            * **Lazy consensus?**
                * **No one’s against it :-) **
    * Docsy theme updates
        * Umbrella issue: [https://github.com/kubernetes/website/issues/20344](https://github.com/kubernetes/website/issues/20344)
    * Link Checker [Celeste]
        * [https://github.com/kubernetes/website/issues/20607](https://github.com/kubernetes/website/issues/20607)
        * Tim will review
        * Revisit on May 12
    * Diagrams [David]
        * PR: [https://github.com/kubernetes/website/pull/20767](https://github.com/kubernetes/website/pull/20767)
        * Tim will review
        * Other feedback welcome
    * Blog [Kaitlyn]
        * In search of more technical and editorial blog editors. Technical blog editors should have deep knowledge of the Kubernetes ecosystem. Editorial blog editors do not need to be as experience with Kubernetes, but should still have experience with git/GitHub
        * Shadow application can be filled out here: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows)
    * Any PRs needing review / attention?
        * [https://github.com/kubernetes/website/pull/20659](https://github.com/kubernetes/website/pull/20659) needs lgtm
        * [https://github.com/kubernetes/website/issues/20607](https://github.com/kubernetes/website/issues/20607) (Celeste)                                              
* Community
    * [Zach] Google Season of Docs intern: [https://github.com/cncf/mentoring/tree/master/seasonofdocs](https://github.com/cncf/mentoring/tree/master/seasonofdocs)
    * [Brad] investigating CLA issues with CNCF (role field)

**4/28/2020**

**6pm Pacific**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @bradtopol
* Next week’s PR wrangler is @rajakavitha3
    * [Zach will reach out to confirm wrangling and say hello]
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * Follow up on proposal to extend release by a month (3 weeks).
        * [https://groups.google.com/forum/#!topic/kubernetes-dev/IVpiIOZ4WcM](https://groups.google.com/forum/#!topic/kubernetes-dev/IVpiIOZ4WcM)
    * Docsy template updates
        * [Zach] Confirmed with contractor, EOD April 30 available for testing
        * [review at 5/5 weekly meeting]
    * Tracking issue for third party content: [https://github.com/kubernetes/website/issues/20232](https://github.com/kubernetes/website/issues/20232)
        * 
    * Diagrams [David]
        * Does the mermaidjs addition impact your work?
            * [https://github.com/kubernetes/website/pull/20434](https://github.com/kubernetes/website/pull/20434)
            * [https://github.com/kubernetes/website/pull/20621](https://github.com/kubernetes/website/pull/20621)
            * Mermaid & K8s architecture diagrams: 
                * Review in 90 days: overlap/cruft? 
                * Jim will write a blog post about mermaidjs by May 29th about these cool new toys
                * David K. will write about architectural diagrams, possibly in same blog post but possibly its own post
    * Blog [Kaitlyn]
        * In search of more technical blog editors. Technical blog editors should have deep knowledge of the Kubernetes ecosystem. 
        * Shadow application can be filled out here: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject#shadows)
    * Any PRs needing review / attention?
        * 
    * [Jim] Goldmark changes
    * Monthly Korean l10n update
        * Membership update:
            * - New K8s community member from Korean l10n: @pjhwa
            * - New approver for Korean contents: @ysyukr (total 5 members)
            * - web-site-maintainer from Korean l10n will be updated (now 4 members, 3 members will be selected from ko-owners)
        * Team Milestones
            * - Four team milestone branches have been merged into master (dev-1.17-ko.7, dev-1.17-ko.8, dev-1.18-ko.1 and dev-1.18-ko.2)
            * - dev-1.17-ko.8 was the final milestone branch for dev-1.17                                                   
* Community
    * [Zach] Google Season of Docs intern: [https://github.com/cncf/mentoring/tree/master/seasonofdocs](https://github.com/cncf/mentoring/tree/master/seasonofdocs)
    * [Brad] investigating CLA issues with CNCF (role field)

**4/21/2020**

**10:30am Pacific**

New contributors



* Tina Oberoi - first steps getting involved - get started help

Updates/reminders



* This week's PR wrangler is @jimangel
* Next week’s PR wrangler is @bradtopol
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release update 1.19
        * Proposal to extend release by a month (3 weeks).https://github.com/kubernetes/sig-release/pull/1058
        * No deadline changing for Docs
        * Proposed: August 18 and July (email out for lazy consensus)
    * Docsy template updates
        * EOM April 30
    * Diagrams [David]
        * Part 1
            * Coming up with mini scope of work
            * (alter 9 icons) - need improvement from graphic artist
            * At a point where we are ready to bring in a contractor
        * Part 2
            * Create a pull request to merge new content
            * Done screencast recordings 
            * ID things are missing / gaps
            * Reaching out to sigs / clouds for feedback on better apps diagrams
        * AI: open issue with feedback
    * Blog [Kaitlyn]
        * Reverted blog post yesterday due to community feedback
            * First time we’ve had to do this so process will be smoother in the future
            * Reminder: we are looking for additional technical reviewers 
                * Public meeting every 2 weeks to review posts 
        * Updating the blog guidelines to provide more guidance on submitting posts and what is acceptable content: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject)
        * Shadow application: [https://docs.google.com/forms/d/e/1FAIpQLScg9fHsyW-LlsBF8rc9J0sR8u3O3g17lwFUKIE-qrjL6Z-AyA/viewform](https://docs.google.com/forms/d/e/1FAIpQLScg9fHsyW-LlsBF8rc9J0sR8u3O3g17lwFUKIE-qrjL6Z-AyA/viewform)
        * PRwrangler - make sure someone is assigned / reach out to assigned person
    * Any PRs needing review / attention?
    * [Jim] Goldmark changes
    * New features [Jim] - [https://mermaidjs.github.io/](https://mermaidjs.github.io/#/)
        * [https://github.com/kubernetes/website/pull/20434](https://github.com/kubernetes/website/pull/20434)
* Community
    * Google Season of Docs intern

**4/14/2020**

**10:30am Pacific**

New contributors



* Kevin Chen
    * Heard about us from CNCF, GitLab, Kong
    * At Kong, works with Kbar! Yay!
* Lorenzo Paris
    * At VMWare

Updates/reminders



* This week's PR wrangler is @sftim
* Next week’s PR wrangler is @zparnold
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Docsy template updates
        * No update this week
    * Diagrams [David]
        * No update this week
    * Blog [Kaitlyn]
        * Questions about impact of third-party content on blog submissions
        * Blog team is reviewing guidelines (which are already stringent) to make sure they’re clear to folks submitting content
    * Any PRs needing review / attention?
    * [Jim] Goldmark changes
        * PR to v0.69.0: [https://github.com/kubernetes/website/pull/19907](https://github.com/kubernetes/website/pull/19907)
            * and issue https://github.com/kubernetes/website/issues/20335
        * comments:
            * Ref gen convert to markdown (k/k engineers)
            * We cannot break the reference page (HTML).
        * Fixes:
            * Remove “feature state” widget, swap for a static image? [SFTim] 
* Community
    * Google Season of Docs intern
        * 
    * Jared Bhatti working on Fuchsia, stepping back from approvals & PR wrangling
    * New features [Jim] - [https://mermaidjs.github.io/](https://mermaidjs.github.io/#/)
        * Would be handy for a release page
        * [https://github.com/kubernetes/website/issues/20293](https://github.com/kubernetes/website/issues/20293)

**4/7/2020**

**10:30am Pacific**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @zacharysarah
* Next week’s PR wrangler is @sftim
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Docsy template updates
        * No update this week; Zach will ping contractor for status
    * Diagrams [David]
        * Incorporating official icons from Celeste
        * Revisit scope of work: much smaller and more focused (Zach: great!)
    * Any PRs needing review / attention?
        * [https://github.com/kubernetes/website/pull/20020](https://github.com/kubernetes/website/pull/20020) [Tim B]
        * [https://github.com/kubernetes/website/pull/19576](https://github.com/kubernetes/website/pull/19576) [Celeste]
    * Blog [Kaitlyn]
        * No update today
    * [Jim] Goldmark changes
        * “unsafe” (allows &lt;scripts> embedded, disallows HTML): [https://github.com/kubernetes/website/pull/19907](https://github.com/kubernetes/website/pull/19907)
        * “safe” (disallows &lt;scripts>, disallows HTML): [https://github.com/kubernetes/website/pull/19905](https://github.com/kubernetes/website/pull/19905)
        * comments:
            * Ref gen convert to markdown (k/k engineers)
            * We cannot break the reference page (HTML).
            * Pain now or later
* Administration

**3/31/2020**

**10:30am Pacific**

New contributors



* Heena
* Davi Garcia

Updates/reminders



* This week's PR wrangler is Tim B.
* Next week’s PR wrangler is @xiangpengzhao (Peter Zhao)
    * Posted for help on #sig-docs-maintainers
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 released! 3/25
        * Comment about 1.19 docs improvement
    * Docsy template updates
        * n/a
    * Diagrams [David]
        * Issue: [https://github.com/kubernetes/website/issues/19811](https://github.com/kubernetes/website/issues/19811)
        * Early stages / draft w/ example
        * Looking for more feedback / help / edits
    * Any PRs needing review / attention?
        * n/a
    * Blog [Kaitlyn]
        * Help wanted if interested!
        * Blog subproject: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject)
    * Merge Strategy
        * Tracking Issue: [https://github.com/kubernetes/release/issues/1204](https://github.com/kubernetes/release/issues/1204)
        * **TODO:** Celeste volunteered to create the content / review for how to squash commits
            * Not started - who is this doc intended for (Author)
                * New / UI
                    * Press to merge via CLI
                    * Label if needed (write access) 
                * Fork / clone
            * Tag Stephen August on PR
            * Update contributors guide and link out (discussed earlier today)
            * AI: update issue w/ info
        * Question: Do we have git CLI from start to finish anywhere for docs
            * Black laptop > setup env + extra (coffee shop version)
            * AI: discuss in planning (+ [https://github.com/kubernetes/release/issues/1204#issuecomment-606774424](https://github.com/kubernetes/release/issues/1204#issuecomment-606774424) )
* Administration
    * Tech Lead
        * Kaitlyn B. nominated Taylor Dolezal
            * [https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#tech-lead](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#tech-lead)
            * Alignment / Vote
            * Next steps would be
                * PR [https://github.com/kubernetes/community/blob/master/sigs.yaml#L1252](https://github.com/kubernetes/community/blob/master/sigs.yaml#L1252)
    * Quarterly review **April 1st 11am Pacific**
        * Agenda: [https://docs.google.com/document/d/1hRTGcegeeUubYwiOjDupYWmXhFlvS-rp6sEk752sG6I/edit?usp=sharing](https://docs.google.com/document/d/1hRTGcegeeUubYwiOjDupYWmXhFlvS-rp6sEk752sG6I/edit?usp=sharing)
        * Zoom: [https://zoom.us/j/957359264](https://zoom.us/j/957359264)

**3/24/2020**

**6pm* Pacific APAC**

*SIG Docs meeting will be held during the google calendar meeting’s scheduled time, **1:00 UTC**. To be later adjusted correctly to **2:00 UTC**. This ends up being 6pm Pacific. This error was caused by daylight savings time (US).

New contributors



* 

Updates/reminders



* This week's PR wrangler is @tengqm
* Next week’s PR wrangler is @ryanmcginnis
    * Need replacement
    * **AI:** Check with Peter and Jenifer (April 5 / 12) for PR rotation
    * Add Tim and Karen into the rotation
    * Zach open to another rotation (19th?)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18
        * Not quite… out the door!
        * Issues / retro
    * Docsy template updates
        * Update: Aim to be done by mid April
        * Feedback: desktop looks great, mobile not so much
    * Diagrams (David)
        * Combing through feedback, opened issue: [https://github.com/kubernetes/website/issues/19811](https://github.com/kubernetes/website/issues/19811) 
        * Copied SVG and is reviewing ([https://docs.google.com/drawings/d/11P-yXzso7zapqyqP3ulnIsQ1UE7IbzEi-7kp5jNebLM/edit?usp=sharing](https://docs.google.com/drawings/d/11P-yXzso7zapqyqP3ulnIsQ1UE7IbzEi-7kp5jNebLM/edit?usp=sharing) )
        * Goal: Give developers a framework to start with to improve
        * Finish draft this week for review soon
        * Part 2: If launched, where do we put this?
            * Istio has guidelines where we have “start contributing”
            * Zach: under or close to style guide.
    * Any PRs needing review / attention?
        * [Qiming] [https://github.com/kubernetes/website/pull/19526](https://github.com/kubernetes/website/pull/19526)
            * Causing errors and issues, can be fixed in english version which then will benefit localization teams.
            * Will be a BIG change to the website.
            * **AI:** Discuss at quarterly planning meeting
    * KEP [Zach C]  [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327)
        * **AI:** close [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748) and open new umbrella issue to track 3rd party content for removal.
        * No changes until after the 1.19 cycle (held for review)
        * Please don’t spend bandwidth justifying the change, focus on action and point towards KEP when needed.
    * Blog [Kaitlyn]
        * Few new posts out last week
        * A couple in the pipeline after the 1.18 release
        * 2-3 in the “5 days of k8s series”
        * Help wanted if interested!
    * Merge Strategy
        * [https://github.com/kubernetes/test-infra/pull/16892](https://github.com/kubernetes/test-infra/pull/16892)
        * Related Issue (the “why”): [https://github.com/kubernetes/release/issues/1204](https://github.com/kubernetes/release/issues/1204)
        * **AI:** As approvers, make sure PRs against master are squashed.
        * Celeste volunteered to create the content / review for how to squash commits
    * Video / Steve Perry status (from doc) [David]
        * [https://docs.google.com/document/d/1UDSv2i3kHz180WQCokZAQRf1G-bcbO-gxn_YzE9SOyA/edit?ts=5e6190ed](https://docs.google.com/document/d/1UDSv2i3kHz180WQCokZAQRf1G-bcbO-gxn_YzE9SOyA/edit?ts=5e6190ed)
        * **AI:** discuss at quarterly review about how to move forward, while avoiding tech debt.
    * Korean l10n update
        * Membership update: No
        * Team Milestones
            * Two team milestone branches(dev-1.17-ko.5, dev-1.17-ko.6) have been merged into master
            * dev-1.17-ko.7 will be merged into release-1.17 on Mar 26 KST
            * dev-1.18-ko.1 will be started on Mar 26 KST
* Administration
    * Quarterly review April 1st 10am Pacific
        * 4/1 at 10am is not ideal… Try 4/1 **11am** or 4/3 at **10am** Pacific
        * **AI: **Zach to surface discussion.
        * **AI: **[Zach, Kaitlyn, Jim] Need to put out an Agenda doc by Friday 3/27
            * Blast to slack and ask for folks’ feedback
        * Fixing Calendar woes (focus on UTC)

**3/17/2020**

**10:30am Pacific**

New contributors



* Tedodor - bare metal clusters interest

Updates/reminders



* This week's PR wrangler is @jimangel
* Next week’s PR wrangler is @tengqm
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 update
        * Rey - rel notes
        * Docs are reviewed waiting for merge
    * Docsy template
        * No update
    * Diagrams
        * David - back from spring break, picking it back up. Update next week
        * Would like help from other folks
            * Tim will assist
    * Good first issues (live review)
        * [https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22](https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
            * Lifecycle removals
            * Remove pending assignments after 1 week of no activity?
    * Anyone have PRs needing review / attention?
        * [https://github.com/kubernetes/website/pull/19530](https://github.com/kubernetes/website/pull/19530) options:
            * Remove taint keys/values all together (as neolit suggests) 
            * merge PR as-is based on previous examples.
                * [https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/admin/cloud/ccm-example.yaml](https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/admin/cloud/ccm-example.yaml)
        * [https://github.com/kubernetes/website/pull/19198](https://github.com/kubernetes/website/pull/19198)
            * Can we automate version vars?
            * Merge as is?
        * [https://github.com/kubernetes/website/pull/19086](https://github.com/kubernetes/website/pull/19086)
            * Multiple languages anchor links
            * Communication to localization teams to open fixes.
        * [https://github.com/kubernetes/website/pull/19640](https://github.com/kubernetes/website/pull/19640)
            * Tim B will open an issue
            * Untranslated article slug modifies original blog post for Kubernetes 1.17 release
* Administration
    * KEP [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327)
        * Has consensus lgtm, should merge soon and we can move forward
        * Merged!
    * Quarterly review
        * April 2nd 10am Pacific
        * Tentatively changing (1st preferred or 3rd)
* Automation proposal (Jim)
    * [https://docs.google.com/document/d/1Y1u8IBfs5p9bQuUMxarltHoioMmuKybgtGL9U8tlO0E/edit#](https://docs.google.com/document/d/1Y1u8IBfs5p9bQuUMxarltHoioMmuKybgtGL9U8tlO0E/edit#)
    * Comment on merge method defaults
        * [https://github.com/kubernetes/test-infra/blob/master/config/prow/config.yaml#L503](https://github.com/kubernetes/test-infra/blob/master/config/prow/config.yaml#L503)
        * [https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/config.md#general-configuration](https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/config.md#general-configuration)
        * Also relevant to Docs releases & merge conflicts:[ https://github.com/kubernetes/sig-release/pull/1014](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fsig-release%2Fpull%2F1014)
    * AI: can we switch to using the default merge method?
        * Lazy consensus to start convo officially with PR
* Video / Steve Perry status (from doc)
    * David will follow up next week with a link.

**3/10/2020**

**10:30am Pacific**

New contributors



* Sujay Pillai - k8s meetup

Updates/reminders



* This week's PR wrangler is @jaredbhatti
* Next week’s PR wrangler is @jimangel
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 update
        * Dev-1.8 branch is not healthy
        * Enhancements have been reached out to rebase
    * Docsy template
        * No updates
    * Training tab is live! k8s.io/training
        * woo!
    * Diagrams
        * Next steps: open issue / add to docs (istio guidelines / k8s logos)
        * Needs help opening.
    * Good first issues (live review)
        * [https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22](https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
    * Anyone have PRs needing review / attention?
* Administration
    * Tech leads
    * KEP [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327)
        * Meeting with SIG Architecture this week (thursday)
            * [https://github.com/kubernetes/community/tree/master/sig-architecture](https://github.com/kubernetes/community/tree/master/sig-architecture)
            * Window for public comment open
    * Quarterly review
        * Date/time?
        * April 2nd 10am Pacific
* Automation (Jim)
    * Need time to review input [https://docs.google.com/document/d/1Y1u8IBfs5p9bQuUMxarltHoioMmuKybgtGL9U8tlO0E/edit#](https://docs.google.com/document/d/1Y1u8IBfs5p9bQuUMxarltHoioMmuKybgtGL9U8tlO0E/edit#)
    * Questions?
    * Comment on merge method defaults
        * [https://github.com/kubernetes/test-infra/blob/master/config/prow/config.yaml#L503](https://github.com/kubernetes/test-infra/blob/master/config/prow/config.yaml#L503)
        * [https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/config.md#general-configuration](https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/config.md#general-configuration)
        * Also relevant to Docs releases & merge conflicts:[ https://github.com/kubernetes/sig-release/pull/1014](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fsig-release%2Fpull%2F1014)
    * Kicked to next week

**3/3/2020**

**10:30am Pacific**

New contributors



* Eric Knauer (Seattle, Netapp)

Updates/reminders



* This week's PR wrangler is @onlydole (thanks!)
* Next week’s PR wrangler is @jaredbhatti
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 update
        * AI: Zach will follow up on merge commit
        * [https://github.com/kubernetes/website/pull/19424](https://github.com/kubernetes/website/pull/19424) 
    * Docsy template
        * Blog and Partners pages are ready for testing: [https://kubernetes-hugo-staging.netlify.com/](https://kubernetes-hugo-staging.netlify.com/)
        * 3-4 weeks to completion
    * Training tab
        * Add a training tab in the top nav: [https://github.com/kubernetes/website/pull/19214](https://github.com/kubernetes/website/pull/19214) 
    * Diagrams (David K)
        * Open an issue describing what to add to docs:
            * Istio diagrams: [https://istio.io/about/contribute/diagrams/](https://istio.io/about/contribute/diagrams/)
            * [https://github.com/kubernetes/community/tree/master/icons](https://github.com/kubernetes/community/tree/master/icons)
    * PRs for review
    * Good first issues
        * [https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22](https://github.com/kubernetes/website/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
* Administration
    * Technical leads
        * Tim / Karen
        * Need emails
    * KEP
        * [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327)
* (Jim) Automation Discussion (strategy)
    * History / branching: [https://jimsurls.com/git-branching/](https://jimsurls.com/git-branching/)
        * [https://github.com/kubernetes/release/issues/956](https://github.com/kubernetes/release/issues/956) 
        * Admin (repo admin) vs Fork (hand tide label)
        * Merge vs rebase
            * [https://medium.com/datadriveninvestor/git-rebase-vs-merge-cc5199edd77c](https://medium.com/datadriveninvestor/git-rebase-vs-merge-cc5199edd77c)
            * Resolving merge conflicts
        * I can automate a strategy “action” but dealing with edge cases still needs human intervention?
    * AI: Jim to summarize the problem / proposed solution?

**2/25/2020**

**6pm Pacific APAC**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @daminisatya
    * Need new wrangler / modify @daminisatya’s approver status
* Next week’s PR wrangler is @gochist
    * June Yi comment: concern about reviewing english PRs (is an approver for KO localization)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 update
        * Some question over the health of dev-1.18 last week
        * It’s resolved now and merging master in normally
    * Docsy template
        * Moving slowly but it’s moving
    * Diagrams (David K)
        * Istio diagrams: [https://istio.io/about/contribute/diagrams/](https://istio.io/about/contribute/diagrams/)
        * Architectural icons: [https://github.com/kubernetes/community/tree/master/icons](https://github.com/kubernetes/community/tree/master/icons)
        * Zach: potential working group for PoC example (1-2 examples) then ask in SIG Docs if there’s interest in working on diagrams.
        * Request a slack channel: [https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md#requesting-a-channel](https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md#requesting-a-channel)
    * PRs for review
        * Refactor contributor guide: [https://github.com/kubernetes/website/pull/19109](https://github.com/kubernetes/website/pull/19109)
        * Add Training tab: [https://github.com/kubernetes/website/pull/19214](https://github.com/kubernetes/website/pull/19214)
    * Korean l10n update
        * Membership update: No
        * Team Milestones
            * two team milestone branches(dev-1.17-ko.3, dev-1.17-ko.4) are merged into master
            * local offline meetup which was planned on Feb 13 had been cancelled by COVID19
* Administration
    * Technical leads - share your interest by 3/3/2020 (more details in SIG Docs slack)

**2/18/2020**

**10:30am Pacific**

Previous agenda was deleted, possibly maliciously. Going forward, editing access is strictly limited to co-chairs

New contributors



* 

Updates/reminders



* This week's PR wrangler is @makoscafee
* Next week’s PR wrangler is @daminisatya
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * Release 1.18 update
        * [https://github.com/kubernetes/website/pull/19117](https://github.com/kubernetes/website/pull/19117) broke dev-1.18 because it didn’t have the --merge-method-merge label on it
            * How to resolve?
                * One possibility: Rebase on [https://github.com/kubernetes/website/pull/19107](https://github.com/kubernetes/website/pull/19107), forcibly push the branch, then do a merge PR with the correct merge method label applied
                * Better ideas?
    * KEP for third-party & dual-sourced content
        * [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327)
        * We’re on the [Steering Committee agenda](https://docs.google.com/document/d/1qazwMIHGeF3iUh5xMJIJ6PDr-S3bNkT8tNLRkSiOkOU/edit#) for the open meeting on Monday, March 2
        * When this KEP passes, many backlogged PRs will need review/approval
    * Docsy template
        * Zach has pinged contractor for an update
        * Preview link is here: [https://kubernetes-hugo-staging.netlify.com/docs/home/](https://kubernetes-hugo-staging.netlify.com/docs/home/) 
    * [https://github.com/kubernetes/website/issues/19154](https://github.com/kubernetes/website/issues/19154)
        * Requires netlify permissions to diagnose/fix
        * Outcome depends on discussion about technical leads
    * Hugo {{ .Summary }} attribute and effect on SEO
        * [https://kubernetes.slack.com/archives/C1J0BPD2M/p1581881172196600](https://kubernetes.slack.com/archives/C1J0BPD2M/p1581881172196600) 
        * Also relevant to technical leads discussion
        * @RemyLeone on Slack
    * Diagrams (David K)
        * [https://istio.io/about/contribute/diagrams/](https://istio.io/about/contribute/diagrams/)
        * If you know a graphic design resource, reach out to @david.kypuros on K8s Slack
    * [https://github.com/kubernetes/website/issues/19081](https://github.com/kubernetes/website/issues/19081)
        * Not a good first issue but you can see it from there
    * How can we deepen the pool of good first issues? (Discuss 2/18)
        * Take an “efficiency hit” in order to make obvious/trivial/easy tasks available for first contributors
        * Hackathon 
            * Section by section: review in weekly meetings, note what’s wrong and convert the feedback into issues (several +1s!)
            * Do this at KubeCon
            * Do this at local meetups
        * Break issues into good-first-issue pieces?
            * [https://github.com/kubernetes/website/issues/19139](https://github.com/kubernetes/website/issues/19139)
        * Revive new contributor ambassadors?
            * Target specific mentors for specific issues
            * How do we make this idea visible, and how do we normalize it?
        * Add Prow label “mentor-available” (Brad will open issue in SIG contribex, community repo)
* Blog
    * Technical lead interest for blog as well--blog encountering netlify issues and would benefit from a technical lead
* Administration
    * Update calendar invite with latest agenda URL
* Process
    * Technical leads

**2/11/2020**

**10:30am Pacific**

Previous agenda was deleted, possibly maliciously. Going forward, editing access is strictly limited to co-chairs

New contributors



* Adam Kaplan, Red Hat (contributed years ago, welcome back!)

Updates/reminders



* This week's PR wrangler is @kbhawkey
* Next week’s PR wrangler is @makoscafee
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Process
    * Chair transition
        * Vote: Jennifer Rondeau becomes co-chair emeritus
        * Vote: Nominate and affirm Kaitlyn Barnard as co-chair
    * Vacation
        * Jim 2/13 -> 2/18
* Docs
    * Release 1.18 update
        * Release-1.17 force rebased (DONE per @vineeth)
        * Merging master into dev-1.18 is happening: [https://github.com/kubernetes/website/pull/19055](https://github.com/kubernetes/website/pull/19055) 
        * From Savitha: “Docs release team started tracking enhancements working on placeholder PR's for Feb 28 deadline. Overall status: green.”
    * Using url vs slug vs alias: [https://github.com/kubernetes/website/issues/19068](https://github.com/kubernetes/website/issues/19068)
        * Use an alias instead of a url attribute unless you’re really sure you know what you’re doing
        * Best practices for aliases: discuss on 2/18
    * Test changes locally!
        * #19068 was on Zach as an approver, and that’s the expectation of approvers: you approve a breaking change, you help fix it
        * Contributors to test your work before review
    * [Adam Kaplan] Docs proposal for scheduling and eviction
        * Open an issue with proposed refactor plan
        * Open iterative PRs
        * Make sure existing content gets redirected properly
* Architecture
    * Remove Hugo version check: [https://github.com/kubernetes/website/pull/18817](https://github.com/kubernetes/website/pull/18817) 
    * Do we want to review release notes? [https://github.com/kubernetes/kubernetes/pull/87879#issuecomment-583082473](https://github.com/kubernetes/kubernetes/pull/87879#issuecomment-583082473) 
* Cleanup
    * [https://github.com/kubernetes/website/pull/18817](https://github.com/kubernetes/website/pull/18817) (Tim Bannister, _but I might miss the meeting_). Can someone talk about this for me? (Zach says: Sure!)
    * Docs cleaned up complete (jim)
* Diagrams
    * [https://istio.io/about/contribute/diagrams/](https://istio.io/about/contribute/diagrams/)
* How can we deepen the pool of good first issues? (Discuss 2/18)

**2/4/2020**

**10:30am Pacific**

New contributors



* Tegan Broderick (recent grad, junior dev)
* Brandon Willmott (VMWare)
* Rey Lejano (RX-M)

Updates/reminders



* This week's PR wrangler is @kbarnard10
* Next week’s PR wrangler is @kbhawkey
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Release 1.18 update
    * Vineeth to have a working session with Jim to force rebase release-1.18 (still pending)
    * Any updates on the dev-1.18 branch and merging master? (If a release member is present) 
    * Update on enhancement docs tracking: Each docs team member has been assigned with 8-11 enhancements. Progress has been made and few PR owners have gotten back with docs status. Overall status is green. If interested, here is the link to docs tracking - [https://docs.google.com/spreadsheets/d/1RtCvByYdcqWc6I_A1cKgeXT2tBS7SyHGvSt_DWXz270/edit#gid=813297075](https://docs.google.com/spreadsheets/d/1RtCvByYdcqWc6I_A1cKgeXT2tBS7SyHGvSt_DWXz270/edit#gid=813297075)  (Savitha)
    * [https://github.com/kubernetes/website/tree/release-1.17](https://github.com/kubernetes/website/tree/release-1.17)
* Security docs update (Seth)
    * None this week
* Blog update (Kaitlyn)
    * None this week
* Diagrams (David K/Zach)
    * Looking for graphic designers for contract work
    * Goal: Create and maintain diagrams
    * David is searching for someone who can actually do the work
    * [https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit?pli=1](https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit?pli=1)
* Docsy Migration Status? (Jared)
    * Gearbox (contractor) is currently doing the work
    * CNCF is funding Gearbox to apply the [Docsy](https://github.com/google/docsy) template to the site.
    * What is better? Seeing the template applied live to the site piecemeal or the site collectively is done: 
        * Other options: Ensure that it’s possible to switch in previews (feature flag) but not switch. Generate two netlify sites per build? Yes. 
    * Considerations: Allowing people to test the site, debugging, communicating out
* Doc Sprint for KubeCon Europe? (Brad)
    * Who’s going?
        * Brad, Zach, Celeste
    * Debugging sprint if template timing works well
    * Celeste could lead an intro to refactored contributor guide
* Chair transition process (Zach)
    * Zach continuing past end of March
    * Jennifer unresponsive for 1+ month, need to remove
    * Kaitlyn unresponsive
    * Brad: thoughts?
* Next week: discuss how to deepen pool of good first issues (2/11)

**1/28/2020**

**6pm Pacific**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @steveperry-53
* Next week’s PR wrangler is @kbarnard10
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
    * Brad will open a PR to disable approver permissions for @simplytunde, hopefully he returns soon!

Agenda



* Docs
    * 1.18 docs update
        * Vineeth: Force push branch? [https://github.com/kubernetes/website/tree/release-1.17](https://github.com/kubernetes/website/tree/release-1.17)
            * AI: Jim following up
        * Krel [WIP]
    * Security docs update
    * New time update (kaitlyn)
        * Doodle: [https://doodle.com/poll/grara8u2g92kn7cs](https://doodle.com/poll/grara8u2g92kn7cs)
        * AI: Zach will talk with Kaitlyn about proposing a specific day/time
    * PR template change—see PR [#18744](https://github.com/kubernetes/website/pull/18744) (Tim B)
    * Reference links and localization (Zach)
        * [https://github.com/kubernetes/website/pull/18895](https://github.com/kubernetes/website/pull/18895)
        * k8s.io/security redirs to zh localization, this [issue ](https://github.com/kubernetes/website/issues/18896)describes the problem.
    * Related conversation around security links (Jim):
        * [https://github.com/kubernetes/website/pull/18834](https://github.com/kubernetes/website/pull/18834)
        * Keep _redirects if the content is gone, not for new redirects
        * Front matter would be better
        * How does this impact search ability?
        * AI: Jim to open issue for when alias are used (in turn for sig-docs-security)
* Korean l10n update.
    * Membership update
        * new reviewer: @ysyukr
    * Team Milestones
        * finished 1.16 (merged dev-1.16-ko.7 onto release-1.16)
        * two team milestone branches(dev-1.17-ko.1, dev-1.17-ko.2) are merged into master
        * Planning a local offline meetup for docs localization at Feb 13
    * Good work, Korean team!
* Process
    * Meeting notes officially moved (Jim): [https://github.com/kubernetes/community/tree/master/sig-docs/meeting-notes-archive](https://github.com/kubernetes/community/tree/master/sig-docs/meeting-notes-archive)
        * Jim A will work on updating links / refs
            * Ex: [https://github.com/kubernetes/community/tree/master/sig-docs#meetings](https://github.com/kubernetes/community/tree/master/sig-docs#meetings)
            * And any old “archive doc”
            * Latest doc will point to repo.
* Diagrams
    * Zach has this AI first on my TODO
        * Message sent to CNCF requesting vendor contacts

**1/21/2020**

**10:30am Pacific**

New contributors



* 

Updates/reminders



* This week's PR wrangler is @simplytunde
* Next week’s PR wrangler is @steveperry-53
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * 1.18 docs update (Savitha)
        * Shadow orientation is done.
        * Dev 1.18 branch sync-up is done
        * End of the freeze is Jan 28th for Docs tracking. 
        * All green for now. 
        * Shadows: Seth McCombs, Irvi Aini, Savitha Raghunathan , Chima Iheanyichukwu 
    * Diagrams (Steve Perry)
        * Statement of work from David K: [https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit?pli=1](https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit?pli=1) 
        * Steve: Has reviewed, needs more clarification on what “Icons” mean. 
        * Zach review: Great start—will forward to CNCF for feedback and hopefully contractor assignment
        * AI: Anyone who is interested should take a look at this document and offer comments, but not required.
    * Video intro for doc sprints (Steve)
        * Next steps: 
            * Create a proposal with a set of goals
            * Propose a budget for creating the video
            * Define a script and goals
            * Timeline?
                * Q1 description of work
                * Q2 demo / rollout, possibly at OSSCon
    * Security docs update (Jim?)
        * AI: Jared: Follow up with Jim on status. 
        * From Jim: Kicked off two weeks ago, second installment will be this Thursday (meets biweekly). They are just getting started so not much to update yet!
    * Blog update (Kaitliyn)
        * 3 posts
            * KubeInvaders - Gamified Chaos Engineering 
            * CSI Support for inline volumes
            * Kubeadm / Openstack
        * Blog meeting meets after this meeting:
            * @ 11:30, [Agenda](https://docs.google.com/document/d/1W5MKkaQGd3YKKZINzj1tJAQbql5R_Y4KAHlFNsJ44Bc/edit#)
        * Doodle: [https://doodle.com/poll/grara8u2g92kn7cs](https://doodle.com/poll/grara8u2g92kn7cs)
    * PR template change—see PR [#18744](https://github.com/kubernetes/website/pull/18744) (Tim B)
        * Save for next week. 

**1/14/2020**

**10:30am Pacific**

New contributors



* Maciej Filocha (Polish localization)
* Celeste Horgan (CNCF)

Updates/reminders



* This week's PR wrangler is @sftim
* Next week’s PR wrangler is @simplytunde
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
    * Thanks, @bradtopol for putting together the PR wrangler schedule!

Agenda



* Docs
    * What to do with third-party docs
        * Zach is still catching up
        * Backstory: Lots of vendor content in the docs already, and a lot of vendors submitting PRs to promote their technology, which doesn’t actually help new users understand and use kubernetes itself. 
        * We support “what helps run default kubernetes in-tree”. - Seems like folks agree on this :)
    * API Docs
        * More of an issue than we thought
        * [https://github.com/kubernetes/kubernetes/pull/84654](https://github.com/kubernetes/kubernetes/pull/84654)
        * [https://github.com/kubernetes/website/issues/18675](https://github.com/kubernetes/website/issues/18675)
    * Docsy conversion proceeds
* Process
    * (Jim A) cleanup of docs** in progress**. Trying to clean up circular references / links. Ultimate state:
        * 1 active doc (this one)
        * Yearly archived docs linked at [https://github.com/kubernetes/community/tree/master/sig-docs](https://github.com/kubernetes/community/tree/master/sig-docs)
        * Any reference to archives will point to the community repo
        * 2019 backup (if needed ASAP): [https://docs.google.com/document/d/1E0oCkvNys1vTjyGucKOWv1b52ld-eqw5XLNT1SvHeNY/edit?usp=sharing](https://docs.google.com/document/d/1E0oCkvNys1vTjyGucKOWv1b52ld-eqw5XLNT1SvHeNY/edit?usp=sharing) 
* SIG Docs Security
    * 
* Blog WG
    * Blog post on bug bounty ([check it out](https://kubernetes.io/blog/2020/01/14/kubernetes-bug-bounty-announcement/)!)
* Proposed different meeting time
    * Kaitlyn will assemble a Doodle :) 
    * Doodle: [https://doodle.com/poll/grara8u2g92kn7cs](https://doodle.com/poll/grara8u2g92kn7cs)
* What’s happening with diagrams?
    * See quarterly release notes (p5): [https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/edit#](https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/edit#)
    * David Kypuros has a draft, will send it for review this week (by Jan 10) to #sig-docs Slack, David will send to Zach for CNCF funding
    * [https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit](https://docs.google.com/document/d/19XY2PkkAB489sM-56lJ9auNcLOd3RV63t_dql8XB_PQ/edit)
* 1.18 Release docs update (Savitha)
    * Savitha will be giving weekly updates
    * Jim: working with the release engineering team on K-Rel
* Congrats to the Polish localization!
    * Zach will tweet about it (DONE)

**1/7/2020**

**10:30am Pacific**

New contributors



* Rey Lejano (RX-M)

Updates/reminders



* This week's PR wrangler is ?
* Next week’s PR wrangler is ?
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



* Docs
    * What to do with third-party docs
        * Resuscitate the KEP?
        * Pull back its scope significantly and proceed without a KEP?
            * What it takes to have 3rd party content in the docs
            * 90 days (?) to find home for content that doesn’t have a home on the k8s website
            * Redirects acceptable
            * Suggestion for blog articles if there’s a piece that shows of Kubernetes features and uses a particular 3rd party component
    * API Docs
        * Would love to get some user metrics specific to satisfaction with API docs
* Process
    * @sftim nomination for approver
        * Approved enthusiastically!
    * We need PR wrangler shifts for 2020
        * Brad Topol will put these together this week (by Jan 10)
    * Currently 160 open PRs for English: [https://github.com/kubernetes/website/pulls?q=is%3Aopen+is%3Apr+label%3Alanguage%2Fen](https://github.com/kubernetes/website/pulls?q=is%3Aopen+is%3Apr+label%3Alanguage%2Fen)
    * Archiving this doc’s 2019 and earlier contents (Jim)
* SIG Docs Security
    * Seth McCombs and Peter Benjamin Co-leading
    * Starting Jan 9th - Meeting Bi-weekly - 9:30AM PST
* Quarterly meeting notes for review
    * [https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/](https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/)
* Proposed different meeting time
    * Kaitlyn will assemble a Doodle :) 
    * Doodle: [https://doodle.com/poll/grara8u2g92kn7cs](https://doodle.com/poll/grara8u2g92kn7cs)
* What’s happening with diagrams?
    * See quarterly release notes (p5): [https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/edit#](https://docs.google.com/document/d/1dqfCjp0XEScBwLNB8NMO0RKE2GfEo5A81s35YjmOssk/edit#)
    * David Kypuros has a draft, will send it for review this week (by Jan 10) to #sig-docs Slack, David will send to Zach for CNCF funding
