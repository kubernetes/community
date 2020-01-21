**12/17/2019**

**APAC - 02:00 UTC (6 Pacific)**

New contributors

*   

Updates/reminders



*   This week's PR wrangler is @kbarnard10
*   No PR wranglers the weeks of December 22 and 29
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Review ZH release-branch issue / path forward
*   Wrapping up the year
    *   We passed the 18000 - (combined issues and PRs) mark.
    *   Vacations?
        *   Zach C on vacation Dec 14-Jan 6
        *   Jared Bhatti on vacation: Jan 1 - Jan 13
        *   Brad Topol out Dec 16-Jan 4
*   Monthly Korean l10n update (June Yi)
    *   Membership update: No
    *   One team milestone branch is merged (dev-1.16-ko.6)
    *   Current team milestone branches: dev-1.17-ko.1, dev-1.16-ko.7(bug fix only, final for 1.16)
    *   @seokho-son and @ianychoi made a [presentation](https://youtu.be/1TLBSmeaVJc) about participating SIG Docs and Korean l10n at Kubernetes Forum Seoul 2019 (Dec 9)

**12/10/2019**

**10:30am Pacific**

New contributors



*   none

Updates/reminders



*   This week's PR wrangler is @jaredbhatti
*   Next week's PR wrangler is @kbarnard10
*   No PR wranglers the weeks of December 22 and 29
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Release 1.17 update
    *   Last week was quite hairy, many thanks to Bob Killen, Guin Saenger, Jim Angel
    *   Website repo is FROZEN until thaw post-release
    *   Official release PR is now: [https://github.com/kubernetes/website/pull/18011](https://github.com/kubernetes/website/pull/18011)
*   Chair business
    *   Jennifer stepping down as chair, not going away
    *   Kaitlyn stepping up as chair
    *   Dates / Comms
        *   Coming to a theatre near you.
*   Better approver guidelines
    *   [https://docs.google.com/document/d/1JCWVEn5h-Wfa-yEirvfRU0FjxHYIMT0nTvq0dBvK2gA/edit?pli=1](https://docs.google.com/document/d/1JCWVEn5h-Wfa-yEirvfRU0FjxHYIMT0nTvq0dBvK2gA/edit?pli=1)
        *   Kaitlyn and Brad requesting feedback for v1 (and location)
    *   David: Can we look further into the New Contributor Ambassador role
        *   Jennifer: Maybe we need to schedule / announce regularly.
        *   Brad to own in January.
            *   Maintain / list / onboard / announce
*   Docsy Status? (Zach C)
    *   Work in progress
    *   Jared: Contractor is lined up w/ tech writers @ Google. Snags during migration will be surfaced. Still trying to get resources / no timeline yet.
    *   David: Is there any visual uplift / styling being considered?
*   KEP Update (Zach C)
*   Wrapping up the year
    *   We passed the 18000 mark.
    *   Vacations?
        *   Zach C on vacation Dec 14-Jan 6
        *   Jared Bhatti on vacation: Jan 1 - Jan 13
        *   Jennifer Rondeau out week of Dec 9 (evidently not today)
        *   Brad Topol out Dec 16-Jan 4
*   PR Wrangler schedule: needs building for Q1/Q2 2020
    *   Jennifer will create first draft, circulate for anyone who needs to amend
    *   DONE: Zach C will freshen OWNERS_ALIASES list
        *   [https://github.com/kubernetes/org/pull/1459/](https://github.com/kubernetes/org/pull/1459/files)
        *   [https://github.com/kubernetes/website/pull/17929/files](https://github.com/kubernetes/website/pull/17929/files)
*   Overlapping initiatives across sigs (docs, contribex, more): JimA suggests opening up meetings, slack conversations
*   Blog subproject: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject)
    *   Aging out docs: don’t fix older than a year, but can be lenient 
    *   tutorials / how-tos / walkthroughs / examples
*   No recommendation around use of sudo in doc. There should be a suggestion around this in ‘[Documentation Style Guide](https://kubernetes.io/docs/contribute/style/style-guide/)’. (Vanou Ishii)
    *   Related Issue \
[https://github.com/kubernetes/website/issues/17181](https://github.com/kubernetes/website/issues/17181)
    *   Needs consistency
    *   Style Guide update is needed - Vanou
*   Rob: I need some Hugo / how-is-this-table-generated help ? :-D
    *   on[ https://github.com/kubernetes/website/issues/16281](https://github.com/kubernetes/website/issues/16281)
    *   website/content/en/docs/reference/command-line-tools-reference/feature-gates.md
*   Next week is APAC time

**12/3/2019**

**10:30am Pacific**

New contributors



*   Simon Forster - joining after a long hiatus :)
*   Jaime Duncan - @VMWare, looking to contribute to docs :)

Updates/reminders



*   This week's PR wrangler is @rajakavitha1
*   Next week's PR wrangler is @jaredbhatti
*   No PR wranglers the weeks of December 22 and 29
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Release 1.17 update
    *   [https://github.com/kubernetes/website/pull/17065#issuecomment-561029393](https://github.com/kubernetes/website/pull/17065#issuecomment-561029393) 
    *   Looks like there’s a CLA problem to resolve before release on Monday, December 9 - Can’t merge otherwise. 
    *   Tracking down how this happened
*   KubeCon update
    *   Thank you Jim for putting together the contributor day experience for docs
    *   Deep dive on Hugo - it’s good to highlight the tooling we support for content. Most people don’t realize this. 
    *   Got lucky with timing, had a lot of experienced contributors
*   4th Tuesday meetings 
    *   Move APAC meeting up 1 hour so we’re not interfering with lunch time with folks in Korea
    *   One hour earlier is easier for east coast folks too.
    *   Overall consensus - yes :)
    *   AI: Jim to send out email comms around this conversation
        *   Lazy consensus on hour change (update)
        *   Move 4th tuesday -> 3rd for Dec
        *   Cancel holiday week (last week of Dec)
        *   Cancel New Years (first week of Jan)
    *   AI: Zach will update the calendar invite. 
*   Configuring a Kubernetes Microservice katacoda tutorial (Brad)
    *   Hosting this with other tutorials? 
    *   Under: [Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-intro/) or [Kubernetes Bootcamp](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/)?
    *   History of poor maintenance, but we can remove if issues are not addressed within ~30 days 
    *   Decision: Place under Kubernetes Basics with agreement on maintenance.
*   Docsy Status? 
    *   Cat herding in progress, still possible by EOY
*   Wrapping up the year
    *   On track for 18,000 issues/PRs by EOY
    *   Vacations?
        *   Zach C on vacation Dec 14-Jan 6
        *   Jared Bhatti on vacation: Jan 1 - Jan 13
        *   Jennifer Rondeau out week of Dec 9 (will miss Dec 10 meeting)
        *   Brad Topol out Dec 16-Jan 4
*   PR Wrangler schedule: needs building for Q1/Q2 2020
    *   Jennifer, Jim, Brad, and Kaitlyn will create
    *   AI: Zach C will freshen OWNERS_ALIASES list
*   Overlapping initiatives across sigs (docs, contribex, more): JimA suggests opening up meetings, slack conversations
*   Blog subproject: [https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject)

**11/26/2019**

**APAC 7pm Pacific**

New contributor:



*   

Updates/reminders:



*   This week's PR wrangler is @tengqm
*   Next week's PR wrangler is @rajakavitha1
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   AI: Change APAC meeting time, 1 hour earlier, Jim to propose in SIG-Docs and implement with lazy consensus.
*   Post-kubecon / holidays
*   Korean l10n team update (June Yi)
    *   Membership Update: No
    *   Three team milestone branches are merged onto master
        *   dev-1.16-ko.3~5
    *   Current team milestone
        *   dev-1.16-ko.6 (~ Dec 5)
    *   Next team milestone
        *   Freeze until 1.17 release
        *   dev-1.17-ko.1 / dev-1.16-ko.7
    *   Preparing for Kubernetes Forum Seoul 2019 (Dec 9-10) https://sched.co/WIRH

**11/19/2019**

**10:30am Pacific**

New contributor:



*   Luis Figueroa (Red Hat)

Updates/reminders:



*   This week's PR wrangler is @ryanmcginnis
*   Next week's PR wrangler is @tengqm
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

No agenda



*   Minimal attendance, presumably bc KubeCon: Steve Perry, Jennifer Rondeau, Luis. Steve went over Q4 planning a bit, we discussed Luis's possible first contributions. Light discussion, no recording.

**11/12/2019**

**10:30am Pacific**

New contributors



*   David K (formerly Marz)

Updates/reminders



*   This week’s PR wrangler is @tfogo
*   Next week’s PR wrangler is @ryanmcginnis 
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Docs: 
    *   Number of PRs looking good: [https://github.com/kubernetes/website/pulls?q=is%3Aopen+is%3Apr+label%3Alanguage%2Fen](https://github.com/kubernetes/website/pulls?q=is%3Aopen+is%3Apr+label%3Alanguage%2Fen) 
    *   Issue health also looks good: repo stats are healthy (yay!)
*   Repeat visibility:
    *   [Zach] Docsy update? 
        *   Zach is working with contractors; hope is by EOY
        *   Zach is less confident that we’ll have testing available by KubeCon NA, good to make a backup plan
    *   Lists and code blocks
        *   Black Friday may be replaced by Goldmark: [https://github.com/gohugoio/hugo/issues/5963](https://github.com/gohugoio/hugo/issues/5963) 
*   Process:
    *   Jim, Zach, and Cody meeting tomorrow (Wednesday) to finalize sprint topics
    *   PR wranglers through EOY:
        *   Brad Topol will wrangle Dec 8, Jan 5
    *   Need to update OWNERS_ALIASES to remove approvers who are no longer current
        *   Everyone can play along! :) 
    *   Quarterly Meeting date: Thursday, Nov 14, 5pm-8pm Pacific
        *   [https://zoom.us/j/191972276](https://zoom.us/j/191972276) 
    *   What is needed to drive 3rd party content KEP forward?
        *   Zach update: responded to initial feedback, pushing back on the idea that feature developers can handwave inclusion
        *   Did an initial review to clarify wording/intent
        *   BUT we need to make some specific policy recommendations to gain traction
            *   Zach will link to specific policy introduced in content guidelines TODAY (11/12)
            *   Brad Topol: Beware bikeshedding!
*   KubeCon meetup:
    *   Pacific Ballroom 23 all day

**11/5/2019**

**10:30am Pacific**

**If you're reading this document today, please join zoom at https://zoom.us/j/299255558?pwd=VHc1S0VjQVJ4OXJJZEJKUkxRaDNoZz09**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @makoscafee
*   Next week’s PR wrangler is @tfogo 
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Docs: 
    *   [Zach] Docsy update? 
        *   Zach is working with contractors; hope is by EOY
        *   Zach is less confident that we’ll have testing available by KubeCon NA, good to make a backup plan
    *   Lists and code blocks
        *   Black Friday may be replaced by Goldmark: [https://github.com/gohugoio/hugo/issues/5963](https://github.com/gohugoio/hugo/issues/5963) 
*   Process:
    *   PR wranglers through EOY:
        *   Need to revise through EOY due to Google withdrawals
        *   Need to update OWNERS_ALIASES to remove approvers who are no longer current
    *   Finalize Quarterly Meeting date: Thursday, Nov 14, 5pm-8pm Pacific
        *   Lazy consensus
        *   Can we change to Wednesday Nov 13?
    *   What is needed to drive 3rd party content KEP forward?
        *   Zach update: responded to initial feedback, pushing back on the idea that feature developers can handwave inclusion
        *   Did an initial review to clarify wording/intent
        *   BUT we need to make some specific policy recommendations to gain traction

**10/29/2019**

**10:30am Pacific**

New contributors



*   Jamie Coleman
*   ykchang
*   Mbroz

Updates/reminders



*   This week’s PR wrangler is @bradtopol (replacing @chenopis)
*   Next week’s PR wrangler is @makoscafee
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Docs: 
    *   [Zach] Docsy update? 
        *   Zach is working with contractors; hope is by EOY
        *   It may be possible to align testing with KubeCon NA
    *   [Brad] Java microservices tutorial?
        *   Brad introduced Michal, who’s giving a demo
            *   Link to demo tutorial: [https://katacoda.com/jamiecoleman/scenarios/kubeconfig](https://katacoda.com/jamiecoleman/scenarios/kubeconfig)
            *   Zach C: Looks good, but raises questions about where tutorial content lives. SIG Docs needs to discuss best practices about where tutorial content lives
            *   Jim: Who owns tutorial content?
    *   New localizations:
        *   Active: Vietnamese ([https://github.com/kubernetes/website/pull/16965](https://github.com/kubernetes/website/pull/16965))
        *   WIP: Russian ([https://github.com/kubernetes/website/pull/16404](https://github.com/kubernetes/website/pull/16404))
        *   ?: Arabic ([https://github.com/kubernetes/community/pull/4172](https://github.com/kubernetes/community/pull/4172))
*   Process:
    *   Finalize Quarterly Meeting date: Thursday, Nov 14, 5pm-8pm Pacific
        *   Lazy consensus
    *   What is needed to drive 3rd party content KEP forward?
        *   [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327) 
        *   [zach] Clear path for action forward (lazy consensus) and expectations.
            *   Avoid opinions encourage actions
            *   Clear set of deadlines for feedback
            *   Criteria needed
        *   [marz] Framework (decision making). How it’s approved / vetted vs. specialized process.
        *   [brad] document certain areas that are not in scope if sensitive
        *   AI: zach to drive kep

**10/22/2019**

**7pm Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @bradtopol
*   Next week’s PR wrangler is also @bradtopol (replacing @chenopis)

Agenda



*   Docs
    *   Reference doc generation: @aimeeu is working on an automation script: [https://github.com/kubernetes/website/pull/16681/](https://github.com/kubernetes/website/pull/16681/)
    *   Better UX: Zach C is working with contractors to apply the Docsy Hugo template to the website. [https://github.com/google/docsy](https://github.com/google/docsy) 
        *   Will probably do a section at a time, starting with docs
        *   Hopefully on track before the end of the year
        *   [Brad] If we can line this up for sprint testing at KubeCon NA, all the better!
    *   [Brad] We may get a preview next week of a Java microservices tutorial put together by IBM team on Katacoda
*   Process
    *   Quarterly SIG meeting--we’re overdue what date shall we choose?
        *   Prefer dates in November of December
        *   Tentatively planning Thursday, Nov 14, 5pm-9pm Pacific
        *   Confirm at weekly meeting on 10/29
    *   Should we create a KEP to finalize policy regarding [dual-hosted content (issue #16091)](https://github.com/kubernetes/website/issues/16091)?
        *   We did, courtesy of @sftim: [https://github.com/kubernetes/enhancements/pull/1327](https://github.com/kubernetes/enhancements/pull/1327) 
        *   [Kubernetes Repository Guidelines](https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md) does not contain guidelines on content that should be in the project repo vs k8s docs
        *   Hopeful for KEP passing by end of year
        *   [Brad] Uncomfortable removing k8s.io/setup/ because it’s high value
            *   [Zach C] Two different conversations: how does 3rd party content get in VS. what should be removed?
            *   Make sure we focus on what/how 3rd party content gets in; what/how to remove will follow sequentially in follow-up action for conformity
*   Korean l10n team update
    *   Membership Update: None
    *   Final team milestone branch for 1.15 is merged
        *   Dev-1.15-ko.6
    *   Two team milestone branches are merged onto master
        *   dev-1.16-ko.1, dev-1.16-ko.2
    *   Current team milestone
        *   dev-1.16-ko.3 (~ Oct 24)
    *   Proposal of @seokho-son and @ianychoi have accepted for Kubernetes Forum Seoul 2019 (Dec 9-10) [https://sched.co/WIRH](https://sched.co/WIRH)
        *   request for contributor patches for the speakers
            *   [Zach C] Contact @paris (Paris Pittman) for patches
            *   Paris Pittman <parispittman@google.com>

**10/15/2019**

**10:30am Pacific**

New contributors



*   Tamao Nakahara (Weaveworks), finally introduced!

Updates/reminders



*   This week’s PR wrangler is @zparnold
*   Next week’s PR wrangler is @bradtopol
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   Docs: 
    *   Third party content KEP
        *   [https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing](https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing)
            *   @sftim: Work in progress, can become a PR - [done](https://github.com/kubernetes/enhancements/pull/1327)
            *   Review / grooming prior to Sunday 10/20 - @sftim will open PR 
    *   
*   Process:
    *   Approver responsibilities
        *   [https://github.com/kubernetes/website/pull/16750](https://github.com/kubernetes/website/pull/16750)
            *   AI: Zach updates PR to move to emeritus approvers and slack (jim and jenifer) with contact information.
        *   [https://github.com/kubernetes/community/blob/master/community-membership.md#approver](https://github.com/kubernetes/community/blob/master/community-membership.md#approver)
    *   Nagging bot: [https://github.com/kubernetes/community/issues/3999](https://github.com/kubernetes/community/issues/3999)

**10/8/2019**

**10:30am Pacific**

New contributors



*   Rin Oliver

Updates/reminders



*   This week’s PR wrangler is @zacharysarah
*   Next week’s PR wrangler is @zparnold
*   Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
    *   Zach C will ask in Slack for a replacement for @chenopis in week of October 27

Agenda



*   Docs: 
    *   Third party content KEP
        *   [https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing](https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing)
        *   Source issue: [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748) 
        *   Content guide: [https://kubernetes.io/docs/contribute/style/content-guide/](https://kubernetes.io/docs/contribute/style/content-guide/)
        *   PR repo: [https://github.com/kubernetes/enhancements/tree/master/keps](https://github.com/kubernetes/enhancements/tree/master/keps)
        *   @sftim and @bradtopol will port umbrella issue and content guide into the KEP template [Complete in 1 week, 10/14]
    *   Release 1.17
        *   Jeefy is a release shadow!
        *   Docs freeze date: EOD Tuesday 15 October
        *   Code freeze date: 14 November
        *   Docs review complete: 19 November (in the middle of KubeCon)
        *   Release date: 9 December
        *   Ping Damini Satya for release date changes/awareness
            *   @sftim: Let’s move up the dates for placeholder PRs to give more time for docs awareness
*   Process:
    *   OWNERS cleanup
        *   Zach will send Bob’s email to SIG Docs googlegroup
        *   Use the link in the email to check your OWNERS file membership
*   Contributor Summit
    *   Docs Sprint - What can we accomplish? 
        *   Cover the basics of starting (CLA, repo)
        *   Create **theme** / agenda prior to (what do we want to cover)
            *   Preso overview about contrib (docs 101)
            *   Then workshop (get in groups)
            *   Have approvers on hand to approve PRs
            *   Maybe split the day, new intros in beginning and more experienced later on.
            *   Note: Themes have been “needs” driven
            *   Advertise the date/time of the sprint
        *   Glossary and low-hanging fruit is a good starting spot
            *   Compared to previous KubeCons, glossary has fewer gaps in it
        *   “good-first-issues” are good too
            *   Maybe curate a list to label on the day of the summit (@sftim)
        *   Target: get folks into docs, get folks into code (git)
            *   Demo as an entry path
        *   Start the day with docs sprint (run to lunch) and break off. Socialize afternoon (last 2 hours)
        *   Sample sprint issue: [https://github.com/kubernetes/website/issues/7452](https://github.com/kubernetes/website/issues/7452) 
        *   Swag?
            *   Zach C: Please no swag, let’s focus on inclusion instead
        *   Any opportunities to promote / reference inclusion?
            *   Zach C: Yes, let’s do it! Open to ideas
            *   @sftim will try to think of something
    *   Lightning talks / subtopics?

**10/1/2019**

**10:30 AM Pacific**

New contributors



*   Matt Boersma (MS, AKS the service/open source)

Updates/reminders



*   This week’s PR wrangler is @stewart-yu
*   Next week’s PR wrangler is @zacharysarah
*   Welcome back from vacation Zach and Jen!
*   The OFFICIAL official announcement that Jim Angel is a SIG Docs chair

Agenda



*   [aimee] For 10/1 meeting, please review for discussion:
    *   [How should we handle pages that contain “allowed project” content that is similar but not exactly a duplicate of the project’s README, especially when the README appears to be more up-to-date](https://github.com/kubernetes/website/issues/16527)
    *   [Content Guide: Revisit wording around adding content relating to CNCF, kubernetes, kubernetes-sigs projects with their own documentation](https://github.com/kubernetes/website/issues/16143)
        *   Defer until KEP resolves from #15748
*   [Zach C] Larger issue of third party content 
    *   Feedback on [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748)
    *   /hold if Zach returns 
    *   Putting together a KEP:
        *   Draft in a google doc
        *   Put draft in issue for review
            *   Tim, Jay, Brad, Aimee, Zach C
        *   [https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing](https://docs.google.com/document/d/1vtA6iNCCzaX-6uUvWA6WtwanBd5EDiFtz9CI0Ek1-lM/edit?usp=sharing)
    *   When working with a steering committee, be patient and keep breathing :)
    *   
*   [Savitha R] Should production setup examples be a part of[ Docs ](https://kubernetes.io/)or[ K8s examples](https://github.com/kubernetes/examples) site? Ref: [https://github.com/kubernetes/website/issues/16635](https://github.com/kubernetes/website/issues/16635) . Can we have examples in K8s examples site and reference the link in [https://kubernetes.io/](https://kubernetes.io/) pages?
    *   Recommend PVs to be a set of tasks - Steve
    *   Savitha to take point on issue alongside engaging sig-storage for assistance with issue 16635
    *   Larger question of how to deal with example content in tutorials
        *   Revisit on 10/8 at top of agenda
        *   Not that many tutorials: we could test them ourselves
*   [Jim A]: Quick note: reference-docs have been “officially” relocated from [https://github.com/kubernetes-incubator/](https://github.com/kubernetes-incubator/) to [https://github.com/kubernetes-sigs](https://github.com/kubernetes-sigs). PR’s have been merged to reflect this change and links to the old repo redirect to the new one.
    *   This is mainly for awareness / docs release team.
*   [Zach C] Technical discussion for what happens if a chair “vanishes”
    *   If a host drops, whoever becomes host, please resume recording
*   Erick Carty
    *   Can we provide specific guidance for which version where localizers start?

**9/24/19**

**FOURTH TUESDAY - APAC **

**7 PM PACIFIC**

New contributors



*   n/a

Updates/reminders



*   This week’s PR wrangler is @steveperry-53
*   Next week’s PR wrangler is @stewart-yu
*   Zach is OOO this week; returns 9/30

Agenda



*   1.16 released! Thanks everyone, especially @simplytunde
*   [aimee] Could an approver please merge these two LGTM PRs?
    *   [Enhance Triage and categorize issues section](https://github.com/kubernetes/website/pull/16295)
    *   [Add references to Content guide alongside Style guide](https://github.com/kubernetes/website/pull/16101)
*   [aimee] For 10/1 meeting, please review for discussion:
    *   [Content Guide: Revisit wording around adding content relating to CNCF, kubernetes, kubernetes-sigs projects with their own documentation](https://github.com/kubernetes/website/issues/16143)
    *   [How should we handle pages that contain “allowed project” content that is similar but not exactly a duplicate of the project’s README, especially when the README appears to be more up-to-date](https://github.com/kubernetes/website/issues/16527)
    *   Note: Moved to next week
*   Korean team update
    *   Membership Update
        *   New K8s Members from Korean L10n: @ysyukr 
    *   One team milestone branch is merged
        *   dev-1.15-ko.5
    *   Current and next team milestone
        *   Current: 
            *   dev-1.15-ko.6 (~ TBD)
            *   dev-1.16-ko.1 (~ Sep 26)
        *   Next: dev-1.16-ko.2 (Sep 27 ~ )
    *   @seokho-son and @ianychoi have submitted a proposal for Kubernetes Forum Seoul 2019 (Dec 9-10)

**9/17/2019**

**10:30 AM Pacific**

New contributors



*   n/a

Updates/reminders



*   This week’s PR wrangler is @mistyhacks (acknowledged)
*   Next week’s PR wrangler is @steveperry-53
*   Zach and Jennifer both OOO this week at Write the Docs Prague

Agenda



*   1.16 Release update
    *   Going through milestones / needing /lgtm and /approve
    *   Nothing currently blocking atm
    *   Overall good
    *   4pm PDT (branch manager, docs going out)
*   Jim A: Triage Issues and assign owners for beginning of sig-docs meeting?
    *   Could also manage stale issues.
    *   [aimee] Is there a “Triage Wrangler” role?
    *   Idea: Looking at stale (closing / removing label).
    *   Time box issues / bring up and look at stale issues.
    *   Enhance Triage and Categorize Issues PR [https://github.com/kubernetes/website/pull/16295](https://github.com/kubernetes/website/pull/16295)
*   Jim A: [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748) - 3rd content
    *   Google dev email thread: [https://groups.google.com/a/kubernetes.io/forum/#!msg/steering/8v8_IkHFX8M/2MXV0z6PAgAJ](https://groups.google.com/a/kubernetes.io/forum/#!msg/steering/8v8_IkHFX8M/2MXV0z6PAgAJ)
*   Aimee: Survey Ended, 99 responses. Nothing to share at the moment. Across the board, users would like more detailed content and advanced tutorials. [https://www.surveymonkey.com/results/SM-2TLG8VVJ7/](https://www.surveymonkey.com/results/SM-2TLG8VVJ7/)
    *   Next time, use google analytics for question origins

**	**

**09/10/2019**

**10:30 AM Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @cody-clark
*   Next week’s PR wrangler is @mistyhacks, Jennifer will confirm
*   Zach and Jennifer both OOO next week at Write the Docs Prague

Agenda



*   Docs
    *   [aimee] Triaging issues - sometimes hard to tell which issues have been looked at by SIG Docs team member; should we add [priority labels](https://github.com/kubernetes/website/labels?page=3&sort=name-asc) to all issues? priority/backlog to low issues, perhaps create a new label for medium priority? Thoughts?
        *   Aimee: Suggest priority/medium
        *   Jared: [Defined prioirties here](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md#define-priority) - might be useful to include examples of each of these in [our triage documentation for SIG-Docs](https://kubernetes.io/docs/contribute/intermediate/#triage-and-categorize-issues). Should make tagging issues with a priority mandatory for triaging
        *   Jay: Sorting by date/staleness ("last touched")
            *   Diff between priority/backlog and priority/important-longterm?
                *   backlog is stuff we accept may never be done
                *   important-longterm is stuff we agree should be done at some point?
        *   Zach: Don’t think we need more granularity - will we work on this now, later, or “never”? (👍 from Tim B, +1 from Jennifer)
            *   
        *   AI: [Aimee] - Will update contributing guide with the following: 
        *   What labels to use
            *   priority/backlog - could defer this forever
            *   priority/important-longterm - within 6 months
            *   priority/important-soon - within 3 months
            *   priority/critical-urgent - need this yesterday
            *   [Labels](https://github.com/kubernetes/website/labels?page=3&sort=name-asc)
        *   AI: Zach C will remove old triage/priority labels from k/website. 
    *   [Vineeth] Docs Release Team update.
    *   [Zach C] Location for upstream fixes to generated reference docs? [https://github.com/kubernetes/website/pull/16263](https://github.com/kubernetes/website/pull/16263) 
*   Layout(?)
    *   [Rob K / Tim B] [Sortable, filterable tables](https://github.com/kubernetes/website/issues/16281)
        *   Rob will update the issue with a summary
        *   Zach suggested that @kbarnard10 might be able / willing to help
*   Community
    *    [Aimee] Survey - where to store questions and results? [https://github.com/kubernetes/community/tree/master/sig-docs](https://github.com/kubernetes/community/tree/master/sig-docs) seems like the correct repo, if we follow contrib-ex’s lead - lazy consensus is YES
*   Tooling
    *   [Zach C] Script for localization diffs by June Yi (@gochist): [https://github.com/kubernetes/website/pull/15789](https://github.com/kubernetes/website/pull/15789) 

**9/03/2019**

**10:30 AM Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @jimangel
*   Next week’s PR wrangler is @chenopis (but will almost certainly reschedule)
    *   @codyclark will wrangle instead

Agenda



*   Docs
    *   Aimee: Consensus on how to handle PRs that add instructional content about a CNCF, kubernetes, or kubernetes-sig project that has its own docs; see [Issue 16143](https://github.com/kubernetes/website/issues/16143) 
        *   Revisit 9/10: Take the week to review and comment on the issue
        *   My worry is if we end up with neither of 2 projects willing to document how they work together - Tim B
        *   (no real strong opinions BTW - Tim B)
    *   Jim A - Deprecate kubernetes-incubator in favor of kubernetes-sigs (ref gen)
        *   Issues:
            *   Deprecation announce: [https://github.com/kubernetes/community/issues/1922](https://github.com/kubernetes/community/issues/1922)
            *   Docs issue to move from release 1.14: [https://github.com/kubernetes/website/issues/13904](https://github.com/kubernetes/website/issues/13904)
        *   I will open the issue to migrate and comment on the community issue, need alignment with sig-docs on timing.
            *   [https://github.com/kubernetes/org/issues/new/choose](https://github.com/kubernetes/org/issues/new/choose)
            *   pick: Repository migration
    *   Release updates: 9/9 all docs need to be reviewed and merged.
    *   Aimee - removing third-party content - running into an issue in which the localizations (fr, it) point to English docs that I want to remove… so just want to make everyone aware when reviewing PRs to ensure localized docs aren’t broken as a result
        *   I’m going to push up a WIP and CC relevant localization leads
        *   Link checker issue: [https://github.com/kubernetes/website/issues/15893](https://github.com/kubernetes/website/issues/15893)
        *   Resolution: share when you come across this, but it’s not realistic to expect approvers to dig into localizations when approving a PR
*   Community
    *   Survey - the Kubernetes Documentation survey test site is live: [https://www.surveymonkey.com/r/8SWG2W3](https://www.surveymonkey.com/r/8SWG2W3)    Please take a look today and direct feedback to @aimeeu on k8s slack
        *   Intro paragraph: The Kubernetes SIG Docs team would like Kubernetes users to evaluate the Kubernetes documentation ([https://kubernetes.io/docs/home/](https://kubernetes.io/docs/home/)). Please take a moment to provide your feedback so we can plan how to improve the Kubernetes docs.
        *   Do we include the survey in this week’s KubeWeekly, which is published tomorrow?
*   [Tunde] 1.16 release is going well, working with shadows to get docs PRs merged by 9/7 (deadline of 9/9)
*   [Zach C] Meeting coverage
    *   Jennifer: will run 9/10
    *   Jim will run 9/17, 9/24

**8/27/2019**

**Fourth Tuesday**

**7pm Pacific**

New contributors



*   Jesang

Updates/reminders



*   This week’s PR wrangler is @zparnold (thanks for stepping in!)
*   Next week’s PR wrangler is @jimangel

Agenda



*   Docs
    *   Removing duplicate content: [https://github.com/kubernetes/website/issues/16091](https://github.com/kubernetes/website/issues/16091) 
    *   Removing third party content: [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748) 
        *   Definition of third party content: [https://kubernetes.io/docs/contribute/style/content-guide/](https://kubernetes.io/docs/contribute/style/content-guide/) 
    *   Aimee is creating a site survey for pain points
        *   Thanks to Brad Topol for his Twitter survey research!
*   Community
    *   Jared Bhatti is stepping down as co-chair in mid September
    *   Chairs have nominated Jim Angel to replace Jared as co-chair
    *   [https://github.com/kubernetes/website/issues/16106](https://github.com/kubernetes/website/issues/16106) 
*   Site engineering
    *   Fixing the deprecation warning: [https://github.com/kubernetes/website/issues/16066](https://github.com/kubernetes/website/issues/16066)
        *   Resolved :-) 
    *   Zach C is searching for a contractor to apply the Docsy template to the site
*   Korean team update
    *   Membership Update
        *   New Korean content approver: @seokho-son
        *   New K8s Members from Korean L10n: @yoonian, @lapee79
    *   Three team milestone branches are merged
        *   [dev-1.15-ko.2](https://github.com/kubernetes/website/pull/15533), [ko.3](https://github.com/kubernetes/website/pull/15744), [ko.4](https://github.com/kubernetes/website/pull/16063)
    *   Current and next team milestone
        *   Current: dev-1.15-ko.5 (~ Sep 5)
        *   Freeze: Sep 5 ~ Sep 9 (waiting for 1.16 release)
        *   Next: dev-1.15-ko.6, dev-1.16-ko.1 (Sep 9 ~ )
    *   Kubernetes Forum Seoul 2019 (Dec 9-10)
        *   @seokho-son is preparing for a proposal (~ Aug 31)
*   Can deprecation warnings (set in config.toml) be overridden at the localization level? 
    *   Luc Perkins has strong Hugo expertise
    *   May not be possible to toggle deprecation per localization, but it is totally possible to localize deprecation warnings (layouts/shortcodes/deprecationwarning.html) to be accurate about maintenance per each localization

**8/20/2019**

**10:30am Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @kbarnard (thanks for stepping in at the last minute!)
*   Next week’s PR wrangler is @zparnold (replacing @jaredbhatti)

Agenda



*   Docs
    *   JD & deprecated flag
        *   Zach C: I realize I hand-waved you off to test-infra without much guidance on adding a new label
        *   Check out [https://github.com/kubernetes/test-infra/pull/9835](https://github.com/kubernetes/test-infra/pull/9835) 
*   Release notes formatting?
    *   [Jennifer will investigate]
*   Kubecon 2019 SIG Workshop(Barnie)
    *   we're looking for Session and Workshop submissions for pre-programmed sessions([https://forms.gle/6tgMmJ4sYHJ4CDxy9](https://forms.gle/6tgMmJ4sYHJ4CDxy9)) - Team will not have/participate or prepare the workshop but will do Meet and Greet.
    *   SIG Meet and greet ([https://forms.gle/hxx1qz8XtwtXEBMm8](https://forms.gle/hxx1qz8XtwtXEBMm8))
        *    Cody will lead, Zach will attend
*   [aimee] Pain Points Survey [https://github.com/kubernetes/website/issues/15828](https://github.com/kubernetes/website/issues/15828)
*   Third-party content: are we truly neutral? [cody] [https://github.com/kubernetes/website/pull/15919](https://github.com/kubernetes/website/pull/15919)
*   Style Guide PR - guidelines for removing third-party content: [https://github.com/kubernetes/website/pull/15892](https://github.com/kubernetes/website/pull/15892)
*   Umbrella issue for removing third-party content: [https://github.com/kubernetes/website/issues/15748](https://github.com/kubernetes/website/issues/15748)
*   1.16 release update (Vineeth): 
    *   Aug 23 placeholder PR deadline
        *   On track
*   Etiquette for requesting PR reviews
    *   Once every 7 days drop into #sig-docs
    *   and/or once every 7 days ping existing reviewers
    *   Would be great to add this to the style guide :-) 
    *   Savitha will add a PR to the style guide

**8/13/2019**

**10:30am Pacific**

New contributors



*   Savitha Raghunathan

Updates/reminders



*   This week’s [PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) is @xiangpengzhao (Peter)
*   Next week’s PR wrangler is @zhangxiaoyu-zidif (Tim)

Agenda



*   Docs
    *   (JD) Proposal for handling deprecated content
        *   [https://docs.google.com/document/d/1pMkSQuiFAbxIP9FxJJRiTPwrpbOqWbOvE26YbWPiQbE/edit?usp=sharing](https://docs.google.com/document/d/1pMkSQuiFAbxIP9FxJJRiTPwrpbOqWbOvE26YbWPiQbE/edit?usp=sharing) 
        *   Can we resolve this with a new repo label/Prow command? (/deprecated)
        *   Flag into @mentions for SIGs and content authors
        *   Create an issue template for deprecated content
        *   Different kinds of obsolescence:
            *   Completely taken down
            *   Outdated code sample
            *   No longer best practice
    *   Pain point discovery
        *   Brad: results on Twitter?
            *   (Brad) [https://stackoverflow.com/search?tab=newest&q=kubernetes%20answers%3a3%20views%3a1000](https://stackoverflow.com/search?tab=newest&q=kubernetes%20answers%3a3%20views%3a1000)
            *   Revisit on 8/20
        *   Seth, Simon?
    *   Survey (Aimee U.)
        *   Proposal [https://github.com/kubernetes/website/issues/15828](https://github.com/kubernetes/website/issues/15828)
        *   Revisit on 8/20
    *   Style guide & third party content
        *   Context: [https://github.com/kubernetes/website/issues/15576](https://github.com/kubernetes/website/issues/15576)
        *   No dual-sourced content from anyone
        *   No links to content outside kubernetes/kubernetes-sigs or the CNCF
*   [PR wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers) & approvers (Zach C)
    *   Some folks have been unresponsive for a while, and Google’s involvement has scaled back considerably
    *   Will Google folks be filling their PR wrangler shifts?
    *   PR wrangling is the one real responsibility we ask of approvers
        *   Audit at end of September
        *   Slackbot to remind folks
*   SIG Docs Chair changes (Jared B)
    *   Planning on stepping down as SIG-Docs Chair End of September
*   Bad certificates
    *   (Luc) More to say about proposal/DNS? (Revisit 8/20)
    *   k/k8s.io
*   Issue triage: how’s it going?
    *   REVIEW: Cody will write issue for how to deal with dual-hosted tutorial content, Brad/Steve/Zach will review (8/20)
*   Slow page serving from East Asia
    *   Flagged to CNCF, waiting on internal response

**8/6/2019**

**10:30am Pacific**

New contributors



*   Aimee Ukasick (@aimeeu)
*   Abhinav.Korpal
*   Nono
*   Savitha
*   Welcome back, Jared!

Updates/reminders



*   This week’s [PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) is @zacharysarah
*   Next week’s PR wrangler is @xiangpengzhao (Peter)

Agenda



*   Docs
    *   (Zach C) Pain point discovery
        *   1. Find pain points
        *   2. Open an issue with a link to specific pain point
    *   1 week volunteers: Seth, Brad, Simon Forster (check in 8/13 on results)
*   Bad certificates
    *   (Luc) More to say about proposal/DNS? (Revisit 8/13)
*   Issue triage: how’s it going?
    *   (Cody C) Issues are lots of requests for support, Katacoda bugs
    *   Zach C will provide support request boilerplate
    *   Katacoda: Remove Katacoda text and link to it instead
    *   Assign remaining Katacoda issues to @BenHall
    *   Cody will write issue for how to deal with dual-hosted tutorial content, Brad/Steve/Zach will review (8/13)
*   Slow page serving from East Asia
*   (JD) Proposal for handling deprecated content
    *   [https://docs.google.com/document/d/1pMkSQuiFAbxIP9FxJJRiTPwrpbOqWbOvE26YbWPiQbE/edit?usp=sharing](https://docs.google.com/document/d/1pMkSQuiFAbxIP9FxJJRiTPwrpbOqWbOvE26YbWPiQbE/edit?usp=sharing) 
    *   Review on 8/13, place early on agenda

**7/30/2019**

**10:30am Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @ryanmcginnis
*   Next week’s PR wrangler is @zacharysarah

Agenda



*   Highlights from [Q2 review/Q3 planning](https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#)
    *   We’re focusing on documentation pain points in Q3.
    *   Call to action:
        *   Find pain points on Stack Overflow, Disqus, and other sources
        *   Open issues in k/website to report them with links
    *   Chair goals through end of year: completely document the chair role and make it possible for new folks to step into the role
*   Site is serving bad certificates for deprecated/deleted branches: [https://github.com/kubernetes/website/issues/12303](https://github.com/kubernetes/website/issues/12303)
    *   (Zach C) Will reach out to netlify and find what it would take to return a 410 instead of a bad certificate
*   Issue triage leads
    *   We need issue triage!
    *   We decided in April to actively pursue issue triage, but life happened and all the folks who agreed to do it have stepped away.
    *   Jdpalomino, jaypipes, codyc
*   Slow page serves from China
    *   Zach C will talk to CNCF about site registration with China
    *   (Tim B) Might there be firewall issues since this was last fixed?
*   (Ryan McG) Good open source citizenship
    *   Thanks everyone for being conscientious <3 
*   (Luc) Issue 14332: how to resolve legacy content for 3rd party content
    *   How should we remove that content?
    *   JD: Will write a design doc proposing how to flag content for deprecation, to be removed during a community day/doc sprint/someone with extra time & energy :-)
*   (Jay Pipes) 

**7/23/2019**

**7pm Pacific - Fourth Tuesday**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @tengqm [Confirmed]
*   Next week’s PR wrangler is @ryanmcginnis
*   [Q2 quarterly review & Q3 planning meeting](https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#): Thursday, 25 July at 5pm Pacific

Agenda



*   Docs
    *   Controller guide update needs review! Work with @sftim: [https://github.com/kubernetes/website/pull/15373](https://github.com/kubernetes/website/pull/15373) 
*   (Seth) Link checker updates, if any?
*   Korean team update
    *   No membership update but meaningful number of contributors are getting interest in docs l10n after the local event, Open Infrastructure & Cloud Native Days Korea 2019.
    *   Two team milestone branches are merged
        *   dev-1.15-ko.1 (merged to master, Jul 10)
        *   dev-1.14-ko.6 (merged to release-1.14, Jul 23)
    *   Current and next team milestones
        *   Current: dev-1.15-ko.2 (Jul 9 ~ Jul 25 KST)
        *   Next: dev-1.15-ko.3 (Jul 26 ~ Aug 8 KST)
    *   Changed Ko l10n team online meeting schedule from weekly Tuesday 22:00 to biweekly Thursday 22:00 (KST)
        *   Currently the team is posting meeting records to [personal playlist](https://www.youtube.com/playlist?list=PLAOP7m08QDCWZ7RwGca6cU4vzrOMw3ht7). Is there better place to post them?
            *   Can you confirm that you’re using Zoom to host meetings?

**7/16/2019**

**10:30am Pacific**

New contributors



*   Sam Marder

Updates/reminders



*   This week’s PR wrangler is @bradamant3, who's likely to get to the queue on the late side
*   Next week’s PR wrangler is @tengqm
*   Q2 review/Q3 planning is on Thursday, July 25, 5-8pm Pacific
    *   Make [agenda proposals](https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#heading=h.ghlki2544yp0)

Agenda



*   Docs
    *   Think about which sections to work on next for kubeadm
        *   Look at Stack Overflow for “desire lines”, actual exhibited behavior about folks’ docs needs and pain points -- any volunteers?
        *   Jennifer AI find other volunteers
*   (Seth) Link checker [update](https://kubernetes.slack.com/archives/C1J0BPD2M/p1562689035218300)
    *   Maybe add a “broken links” tag, to expedite issues/PRs with broken links
    *   Sharing the load, @lucperkins (and anyone else) are welcome to help
    *   Investigating -  [https://github.com/raviqqe/liche](https://github.com/raviqqe/liche) - seems to work decent, may fall over under load, still investigating
    *   Update: “No update” :( 
*   (Jay Pipes) LFG update?
    *   From 7/2/2019: Jay Pipes: Looking for folks wanting to pair on PR for removing cluster directory in k/k
    *   (jaypipes): Sorry, no update from me. Still working on this and will try to have a curated list of items for others to take on at next meeting.
*   (Kaitlyn) [Blog process overview ](https://github.com/kubernetes/community/tree/master/sig-docs/blog-subproject)
    *   See team list in doc for assigning blog post PRs
    *   Greater sig-docs team can review and approve minor updates; only new content needs review by editorial team
    *   See also dedicated Slack channel (#kubernetes-docs-blog)
*   [PR 15373](https://github.com/kubernetes/website/pull/15373)
    *   Comments about direction welcome
    *   Also ideas for how to break down into manageable chunks (106 files!)

**7/9/2019**

**10:30am Pacific**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is @zparnold
*   Next week’s PR wrangler is @bradamant3
*   Jennifer will run next week’s meeting
*   Q2 review/Q3 planning is on Thursday, July 25, 5-8pm Pacific
    *   Make [agenda proposals](https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#heading=h.ghlki2544yp0)

Agenda



*   Docs
    *   (Jennifer) No kubeadm update
    *   Think about which sections to work on next
        *   Look at Stack Overflow for “desire lines”, actual exhibited behavior about folks’ docs needs and pain points
*   (Seth) Link checker [update](https://kubernetes.slack.com/archives/C1J0BPD2M/p1562689035218300)
    *   Maybe add a “broken links” tag, to expedite issues/PRs with broken links
*   PR queue is much better this week. Thanks for closing out assigned PRs!
*   (Jay Pipes) LFG update?
    *   From 7/2/2019: Jay Pipes: Looking for folks wanting to pair on PR for removing cluster directory in k/k
    *   (jaypipes): Sorry, no update from me. Still working on this and will try to have a curated list of items for others to take on at next meeting.

**7/2/2019**

**10:30am Pacific**

New contributors



*   Jay Pipes
*   Vinay Amaranathan
*   Abhinav Korpal

Updates/reminders



*   Q3/Q4 PR wrangler schedule is up: [https://github.com/kubernetes/website/wiki/PR-Wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers) 
*   This week’s wrangler is @cody-clark (June 30)
*   Next week’s wrangler is @zparnold (July 7)
*   Happy to announce that Jim Angel is a shadow for SIG Docs chair
*   Q2 review/Q3 planning is on Thursday, July 25, 5-8pm Pacific

Agenda



*   Docs
    *   Start thinking about what we want to accomplish in Q3
    *   [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3) 
    *   (Jennifer) Any input on kubeadm setup?
    *   [https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#](https://docs.google.com/document/d/10EFtv-DEIaMStLjWNiGgmLTJMG6UGY0iEOcPTfJLplY/edit#)
        *   Jay Pipes can reach out to SIG Cluster lifecycle about current state of cluster API vs. kubeadm setup
        *   Jim, Brad, Jay, Jennifer: Cluster API still very alpha-state, kubeadm docs still needed
        *   Jennifer: Will create a plan to move forward by Q3 planning
*   Open PRs
    *   (Zach C) A lot of fresh ones in English, but a lot of older ones too--let’s resolve what we can, close the others, and get the queue down to size.
    *   AI: add tips on closing PRs to guide somewhere?
        *   Zach C: We need to add better guidelines for closing to [https://kubernetes.io/docs/contribute/advanced/#be-the-pr-wrangler-for-a-week](https://kubernetes.io/docs/contribute/advanced/#be-the-pr-wrangler-for-a-week) 
*   Brad Topol on KubeCon Shanghai:
    *   “Death or glory”! 
    *   Rui Chen translated slides: awesome! 
    *   New contrib workshop content, plus more for deep dive
    *   Indonesian l10n folks, Chinese l10n folks showed up to help
    *   ~18-20 people
*   (Jay Pipes): What’s the etiquette for asking for PR reviews?
    *   @mention the team’s PR review
    *   Ask them in Slack
    *   Mailing lists may not get a quick response
*   (Simon Forster) Update on English review pool for i18n questions?
*   (Seth McCombs) Link checker update?
    *   Found 3 different tools, running them each to see what they find
    *   Scans take time
    *   Poor/unstable internet while traveling is slowing efforts
    *   Manual step to one of “Contributing to the docs” pages?
    *   Jim: Link checkers are more for validity; need an external resource checker as well, if we care about running the site offline
*   Running Hugo locally? 
    *   Anyone able to run locally with success? [https://github.com/kubernetes/website/issues/15245](https://github.com/kubernetes/website/issues/15245)
*   Jay Pipes: Looking for folks wanting to pair on PR for removing cluster directory in k/k 

**6/25/2019**

**7pm Pacific - Fourth Tuesday APAC meeting**

New contributors



*   

Updates/reminders



*   This week’s PR wrangler is Kaitlyn Barnard (@kbarnard10)
*   We’re still making the schedule for PR wranglers in Q3/Q4, hopefully an update by next week.
*   Happy to announce that Jim Angel is a shadow for SIG Docs chair 
*   Q2 review/Q3 planning meeting is July 25th, 5pm Pacific 

Agenda



*   (Paris) San Diego contributor summit
    *   Defer to next week
*   Docs updates:
    *   [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3) 
    *   Follow up with Jennifer about kubeadm docs from project card
        *   Review during Q2/Q3 meeting
*   (June Yi, Claudia Kang) Korean team update?
    *   No membership update
    *   Running two team milestone branches temporarily 
        *   dev-1.15-ko.1 -> master
        *   dev-1.14-ko.6 -> release-1.14
    *   Called for a new single point of contact: no volunteer
        *   June Yi remains as a contact point
    *   PTAL: [#14787](https://github.com/kubernetes/website/pull/14787)
    *   Kaitlyn & Zach: Can we boost visibility for Korean and other localization teams?
*   Japanese team update?
    *   No attendees
*   Chinese team update?
    *   Happy KubeCon! Zach is sorry to miss you this week.
*   (Seth McCombs) Link checker update?
    *   No progress this week, update next week
*   (Qiming) PR queue update--looks better now?
    *   Seeing some old PRs for English
*   (Simon Forster) Update on English review pool for i18n questions?

    	Defer to next week


**6/18/2019**

**10:30am Pacific**

New contributors



*   Tunde Oladipupo
*   Josiah (@josiahbjorgaard)

Updates/reminders



*   PR wrangler this week: Steve Perry (@steveperry-53)
*   PR wrangler next week: Kaitlyn Barnard (@kbarnard10)
*   Need PR wranglers schedule for Q3/Q4
    *   Automate? Do it by Thursday :-) 
*   (Barnie) 1.15 updates?
    *   We are trying figuring out how to resolve conflicts in dev1.15 [https://github.com/kubernetes/test-infra/issues/13065](https://github.com/kubernetes/test-infra/issues/13065)
    *   I have proposed Tunde to be docs Lead for next release. He has been a docs Shadow for the past 2 releases.
    *   Zach C: will add Barnie & Jim as ICs with write permissions so they can work directly with branches to resolve merge conflicts

Agenda



*   Docs progress: [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3) 
*   Q2 review, Q3 planning (Zach C)
    *   When should we meet? Mid-July? Late July?
        *   Meeting on Thursday, July 25th, same time slot but 1 hour earlier
        *   [Zach C] Agenda doc for Q2/Q3 and make it available for input at least 2 weeks in advance
*   jimangel - I can't recall, do we have anyone looking into broken links / adding in redirects? With the updated layout, some links that return from google are broken - see: https://kubernetes.slack.com/archives/C0SJ4AFB7/p1560874030038300. Does anyone know if we can re-index the site for google? 
    *   [Seth] Looking into a link checking tool and re-indexing the site from Google
    *   Will be ongoing effort - Jim
    *   PR that shook a few things up: [https://github.com/kubernetes/website/pull/14826](https://github.com/kubernetes/website/pull/14826)

**6/11/2019**

**10:30am Pacific**

New contributors 



*   

Updates/reminders



*   PR wrangler this week: Cody Clark @cody-clark
*   PR wrangler next week: Steve Perry
*   (Barnie) 1.15 updates?
    *   All PR associated with enhancements for this release are merged.
    *   

Agenda



*   (cody-clark) 
    *   Analytics Reports 
        *   Have access to K8s analytics and can create reports
        *   Need explicit permission from CNCF to proceed
        *   Does each localization effort already have a mailing list? 
            *   To not spam kubernetes-sig-docs@googlegroups.com
    *   Page Satisfaction Survey ([#11037](https://github.com/kubernetes/website/pull/11037))
        *   Highlight: [/docs/tutorials/kubernetes-basics/](https://kubernetes.io/docs/tutorials/kubernetes-basics/) (83% found helpful)
        *   Lowlight: [/zh/docs/home/](https://kubernetes.io/zh/docs/home/) (49% found helpful)
    *   Cloud Provider-specific info on the [Service Concept page](https://kubernetes.io/docs/concepts/services-networking/service/) 

**6/4/2019**

**10:30am Pacific**

New contributors 



*   Chris Simpson (Open Bank, 6 months into K8s)

Updates/reminders



*   PR wrangler this week: Cody Clark @cody-clark
*   PR wrangler next week: Ryan McGinnis
*   (Barnie) 1.15 updates?
    *   Jennifer AI to find right resources to help with branch PR that merges from master

Agenda



*   Simon Forster is wondering how native English speakers can help to the localization efforts. After some brainstorming, we maybe can have a pool of native speakers to assist on complex localizations, doubts, helping to rephrase english, clarify concepts, etcetera.
    *   Bilingual contributors are important to bridge those gaps
    *   Also willing native-English only contributors to help l10n teams figure out original meaning
        *   Contacts on other sigs too
        *   Volunteers: Steve Perry, Simon Forster
            *   Include useful changes in contribution/style guidelines
            *   Develop a more formal role for this work?
    *   What can we do after the first onboarding / commit?
        *   Ongoing support of their journey (releases, branching, etc)
            *   How to keep contributors interested, involved
        *   Which concepts to translate -- prioritize, but also which need to stay in English. Glossary/word list of terms. Gender problems
        *   Coordinate with effort to rework Concepts topics -- look at GH Projects for priorities -- Jennifer AI to reach out to l10n leads with this information, also look for external l10n expertise, resources
        *   Working group?
        *   Planning l10n retrospective
    *   **Bring this topic up during the next 2 sig-docs meetings (APAC + regular post kubecon)**
*   Onboarding new contributors: l10n a specific case, how to improve ongoing commitment
*   Analytics data for specific translations? AI: make sure we have access all the time. Cody will also see what he has access to.

**5/28/2019**

**(APAC) 7pm Pacific**

New contributors 



*   Seth

Updates/reminders



*   PR wrangler this week: Tim Zhang @zhangxiaoyu-zidif
*   PR wrangler next week: Cody Clark @cody-clark
*   (Barnie) 1.15 updates?
    *   We are pushing people to continue creating placeholder PR by 30th here are the one opened [https://github.com/kubernetes/website/milestone/32](https://github.com/kubernetes/website/milestone/32)

Agenda



*   Simon Forster is wondering how native English speakers can help to the localization efforts. After some brainstorming, we maybe can have a pool of native speakers to assist on complex localizations, doubts, helping to rephrase english, clarify concepts, etcetera.
    *   Bilingual contributors are important to bridge those gaps 
    *   What can we do after the first onboarding / commit?
        *   Ongoing support of their journey (releases, branching, etc)
    *   **Bring this topic up during the next 2 sig-docs meetings (APAC + regular post kubecon)**
*   (June Yi) Ko l10n update:
    *   New Ko reviewer(Seokho), New K8s Member mainly contribute to Ko l10n(nowjean)
    *   June 11 ~ 18: Pause Ko l10n PR merge, waiting for release-1.15(June 17)
    *   June 19: Elect new Ko l10n leader who will be a single point of contact with upstream. (currently June Yi)

    (@Qiming): Suggestion for PR wrangler: 200 PRs accumulated during the past 2 weeks, not good. Should contact the scheduled wrangler next time and see if he/she would be available. Someone else can step in if the schedule needs some fixes.


    	AI: Jim: check with cody clark


    (@Seth McCombs) PR merged causing dead links, what can we do to stop this? Seth offered to 

*    [https://github.com/kubernetes/website/issues/14220](https://github.com/kubernetes/website/issues/14220) 
*    [https://github.com/kubernetes/website/pull/14442](https://github.com/kubernetes/website/pull/14442) 
*   [https://github.com/kubernetes/website/issues/14466](https://github.com/kubernetes/website/issues/14466)
*   (@qiming) in general we should not have a lot links to the source code, right?
*   Suggestions:
    *   (Jim) Could try: Sig Docs Tools Slack Channel [https://kubernetes.slack.com/messages/sig-docs-tools/](https://kubernetes.slack.com/messages/sig-docs-tools/)
    *   (Qiming) Raja started efforts to clean up dead links
        *   **AI:** Seth to pair up to coordinate efforts

**5/21/2019**

**10:30am Pacific**

New contributors 



*   Simon Forster

Updates/reminders



*   PR wrangler this week: Peter Zhao @xiangpengzhao
*   PR wrangler next week: Tim Zhang @zhangxiaoyu-zidif
*   (Barnie) 1.15 updates?
    *   Nothing much we are still getting folks to open their draft PR as we are approaching the deadline for that.

Agenda



*   Simon Forster is wondering how native English speakers can help to the localization efforts. After some brainstorming, we maybe can have a pool of native speakers to assist on complex localizations, doubts, helping to rephrase english, clarify concepts, etcetera.
    *   Bilingual contributors are important to bridge those gaps 
    *   What can we do after the first onboarding / commit?
        *   Ongoing support of their journey (releases, branching, etc)
    *   **Bring this topic up during the next 2 sig-docs meetings (APAC + regular post kubecon)**
*   Cody Crudgington: examples (tech preview)
    *   Stricter guidelines when PRing tech examples to prevent a simple image/tag switch (vs full deployment test)
    *   Dani_C gap in testing?
        *   Could add testing - is it too heavy compared to # of issues? -Jim 
*   Dani_C Naming node schemes in email
    *   More folks with history jump in on thread on naming nodes
    *   Link: [https://groups.google.com/forum/#!topic/kubernetes-sig-docs/tBtFGieqX7k](https://groups.google.com/forum/#!topic/kubernetes-sig-docs/tBtFGieqX7k) 

**5/14/2019**

**10:30am Pacific**

New contributors 



*   Erick Carty

Updates/reminders



*   PR wrangler this week: @tfogo (?)
*   PR wrangler next week: @xiangpengzhao
*   (Barnie) 1.15 updates?
    *   Zach can approve a PR to this file: [https://github.com/kubernetes/org/blob/master/config/kubernetes/sig-docs/teams.yaml#L258](https://github.com/kubernetes/org/blob/master/config/kubernetes/sig-docs/teams.yaml#L258)
    *   Adding release leads/shadows to milestone team: great addition for handbook!

Agenda



*   Content
    *   (Steve Perry, JD Palomino) Issues/PRs for every topic in Concepts/Overview: how’s that going?
        *   How to add issues to repo project?
    *   What is Kubernetes page: Shavi’s working on it
    *   Adding “CRD” definitions, better definition for addons (they’re a form of operators)
    *   Dani: Spoke with SIG Cluster Lifecycle about replacing “masters/workers” with “control plane/peer compute”
*   Availability
    *   Conference season is starting, so is summer
    *   Jared B is out on sabbatical until 1 August,  \
Zach C will be at conferences/on vacation until 22 Jul
    *   Jennifer Rondeau and Jim Angel will lead weekly meetings
        *   Jennifer is out at conferences/work offsites 21 May, 4 June
*   Version Skew / Docs [Jim A.]
    *   [https://github.com/kubernetes/website/issues/14307](https://github.com/kubernetes/website/issues/14307)
    *   [https://github.com/kubernetes/sig-release/issues/626](https://github.com/kubernetes/sig-release/issues/626)
    *   Example: [https://www.kernel.org/](https://www.kernel.org/) 
    *   Track this for Q3 in the Q2 planning doc: [https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/edit#heading=h.ghlki2544yp0](https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/edit#heading=h.ghlki2544yp0) 
*   (Rael) Are subprojects/hosted subdomains based out of SIG Docs?
    *   Jennifer/Zach C: Nope, we can consult but we don’t own it

**5/7/2019**

**10:30am Pacific**

New contributors



*   

Updates/reminders



*   PR wrangler this week: @rajakavitha1 
*   PR wrangler next week: @tfogo (Zach C will reach out)
*   (Barnie) 1.15 updates?
    *   We met with team of last friday May 3rd may and update the team on the important dates for this release and what is expected from the team 
    *   We have divided the enhancement for each member to help followup with owners of the enhancements to have placeholder PR by 30th so these coming 3 weeks is to make sure we have all enhancements that needs docs have placeholder PR ready.
    *   Team members 
        *   Tunde Oladipupo(simplytunde@gmail.com) Slack/ simplytunde Github: simplytunde
        *   Damini Satya Kammakomati(daminisatya@gmail.com) Slack/@Damini satya Github/ daminisatya
        *   Christian Hernandez(christian.hernandez@yahoo.com) Slack/ christianh814 Github: christianh814 
    *   We have opened the official 1.15 docs PR that will be merged on release date([https://github.com/kubernetes/website/pull/14176](https://github.com/kubernetes/website/pull/14176))

Agenda



*   Content
    *   Review all existing content: 
        *   Work with SIGs to determine whether info is current
        *   Refactor content, apply style guide
    *   Need issues/PRs for every topic in Concepts/Overview: [https://kubernetes.io/docs/concepts/overview/](https://kubernetes.io/docs/concepts/overview/)
    *   Who’s willing to open issues and assign them to a project? [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3) 
        *   JD Palomino & Steve Perry will open concepts, Steve will start in Object Management, JD will start with topics at the top
*   New contributor ambassadors
    *   Zach Arnold (@zparnold), Shawn (@zhangqx2010), Brad Topol (@bradtopol)
    *   How’s that going so far?
*   Issue triage leads
    *   @chenrui333, @codyc, @zhangqx2010, @dani_c
    *   How’s that going so far?
*   Accessibility: Shavi has slides to share (Cody Clark and Rajie)
*   Kube scheduler docs: (JD Palomino) 
*   Zach C: follow up with Paris/contribex on docs presence at community day

**4/30/2019**

**10:30am Pacific**

New contributors



*   

Updates/reminders



*   PR wrangler this week: @bradtopol 
*   PR wrangler next week: @rajakavitha1
*   (Barnie) 1.15 updates?

Agenda



*   Blog subproject:
    *   PR open to make it official:
        *   [https://github.com/kubernetes/community/pull/3640](https://github.com/kubernetes/community/pull/3640) 
        *   It’s merged! Yay!
    *   Still needs to be done:
        *   (Kaitlyn Barnard) Set up OWNERS and optional team(s):
            *   [content/en/blog/OWNERS](https://github.com/kubernetes/website/blob/master/content/en/blog/OWNERS)
            *   If needed: [kubernetes/org/config/kubernetes/sig-docs/teams.yaml](https://github.com/kubernetes/org/blob/d272559f45b76b3bb4b0304cd7b0a9517e97bfc3/config/kubernetes/sig-docs/teams.yaml)
*   New contributor ambassadors
    *   Almost there!
*   Issue triage captains
    *   @chenrui333, @codyc, @zhangqx2010, @dani_c
    *   Goal of issue triage captains is a list of open issues where everything is actionable
*   Running meetings
    *   We need folks to run meetings, both 10:30am slots and 7pm slots. SIG Docs chairs will be absent/partially available for most of the summer.
        *   Willing to provide guidance, opportunities to shadow before leading
    *   Reach out to Zach C on Slack
*   Demographics survey
    *   Thinking about doing a SIG Docs member survey: who we are, job titles
    *   Goal is to gain data about what 
    *   (Steve Perry) We need both technical writers and people to apply the style guides
        *   But we need tech writers dedicated to K8s/open source, at least 50%
        *   Zach C: Support CommunityBridge for “doc bounty”
*   Content
    *   (Jim A) Boost visibility on this issue: [https://github.com/kubernetes/website/issues/13904](https://github.com/kubernetes/website/issues/13904) 
        *   Inquire about best org for refdocs: kubernetes or kubernetes-sigs
            *   @cblecker can point in the right direction
        *   Feedback from Qiming 
    *   This issue needs attention (currently STALE): [https://github.com/kubernetes/website/issues/12333](https://github.com/kubernetes/website/issues/12333) 
*   Localization teams: approval for specific files at root level/outside content path?
    *   June Yi will open issue; Zach C will review
*   Sig Docs security is going well, and this week we are going to focus our efforts on the CVE management process (building out a place for people to put CVE’s and have them visible on Kubernetes.io)

Need feedback from a security/Kubernetes expert on

**4/23/2019**

**7pm Pacific - 4th Tuesday APAC**

New contributors



*   Dan Roscigno (Elastic)

Updates/reminders



*   PR wrangler this week: @jaredbhatti (maybe @cody-clark)
*   PR wrangler next week: @bradtopol
*   (Barnie) 1.15 updates?

Agenda



*   Blog subproject:
    *   PR open to make it official:
        *   [https://github.com/kubernetes/community/pull/3640](https://github.com/kubernetes/community/pull/3640) 
    *   Still needs to be done:
        *   (Kaitlyn Barnard) Set up OWNERS and optional team(s):
            *   [content/en/blog/OWNERS](https://github.com/kubernetes/website/blob/master/content/en/blog/OWNERS)
            *   If needed: [kubernetes/org/config/kubernetes/sig-docs/teams.yaml](https://github.com/kubernetes/org/blob/d272559f45b76b3bb4b0304cd7b0a9517e97bfc3/config/kubernetes/sig-docs/teams.yaml)
*   New contributor ambassadors
    *   Almost there!
*   Issue triage
    *   We need folks willing to work through the issue backlog and identify high-priority issues (Cody C, Rui, Shawn)
    *   We also need folks to identify good first issues and apply the label. (/good-first-issue)
        *   Good first issues are issues that describe all of the work that needs to be done in a PR: no guessing, investigation, or research required. 
*   Running meetings
    *   We need folks to run meetings, both 10:30am slots and 7pm slots. SIG Docs chairs will be absent/partially available for most of the summer.
        *   Willing to provide guidance, opportunities to shadow before leading
*   Demographics survey
    *   
*   Content
    *   Contributor cheatsheet localization [https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet](https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet) (Rui will open issue for English, need separate issues for each localization)
    *   We need help with this issue (probably 2 weeks-1` month of work) [https://github.com/kubernetes/website/issues/12333](https://github.com/kubernetes/website/issues/12333)
    *   Pick the Right Solution 
*   Localization teams: approval for specific files at root level/outside content path?
    *   June Yi will open issue; Zach C will review
*   Jim A: Boost visibility on this issue: [https://github.com/kubernetes/website/issues/13904](https://github.com/kubernetes/website/issues/13904) 
*   Jim A: EOL for Fedv1 refdocs [https://github.com/kubernetes/website/issues/13905](https://github.com/kubernetes/website/issues/13905) 

**4/16/2019**

**10:30am Pacific **

New contributors



*   Cody - Professional services @ Sysdig, Austin k8s meetup

Updates/reminders



*   PR Wranglers: 
    *   This week: Jennifer (@bradamant3)
    *   Next week: Jared (@jaredbhatti)

Agenda



*   New Contributor Ambassador Role (Jared, Jennifer, Zach C)
    *   [Draft doc - Feedback welcome](https://docs.google.com/document/d/10snlue28FL-hjmCuw1UEFubcWk5PohH3gXIOCONqSBw/edit#)
    *   Cody: I think this is a good idea - I’m struggling with the same thing
    *   Any volunteers to be members of this group? 
        *   Brad T, Zach A, Tim
*   Integrating Blog process with SIG-Docs process - any updates? (Kaitlin, Zach C)
    *   AI: Jared: Kaitlin and Zach need to update
    *   Brad: Please document the process
*   Sig-docs-es- Updates? @raelga
*   Sig Docs Security Updates? - Zach A

**4/9/2019**

**10:30am Pacific **

New contributors



*   Barnabas (barnie): Leading the Docs release this cycle - shadow the past two releases

Updates/reminders



*   Qiming (@tengqm) is PR wrangler this week
*   Jennifer (@bradamant3) is the PR wrangler for next week
    *   Would like to switch weeks - please reach out to Jennifer on Slack. 

Agenda



*   (Paris, Kaitlyn) Blog becoming a sub-project of SIG-Docs: Create an editorial team for blog review with an owner and lead as a sub-project of SIG-Docs
    *   Kaitlyn Barnard would be the lead
    *   TODOs: 
        *   Zach C will add a label for blog issues, make sure it can be manually assigned
        *   Zach C will adjust OWNERS file for content/en/blog/ in k/website
        *   Kaitlyn: Write a README.md for content/en/blog outlining blog process file, link to blog rolebook in k/community
        *   Kaitlyn: Set standards around when SIG-Docs maintainers should approve changes. 
        *   Kaitlyn: Write a Meta-Blog Post on how the new blog process will work. 
    *   [Notes for the project planning meeting are here](https://docs.google.com/document/d/1-XWPHvd-Y141JnOurLD6pyk50rJ_fwPQgjmYomB8_9Q/edit#heading=h.kjkuct5152l0)
*   sig-docs-es
    *   Spanish localization of the website is live.
    *   New contributors are starting to help on the localization of the documentation.
    *   The Spanish team is working on a style guide for the Documentation.
    *   The localization progress is being tracked with http[://bit.ly/k8s-sig-docs-es-track](http://bit.ly/k8s-sig-docs-es-track) at this point.
    *   @raelga is working with Ben Hall from Katacoda to help in the localization the labs of the kubernetes.io tutorials.
*   Sig Docs Security (Zach A.) (Can’t attend, double booked sorry)
    *   We have analyzed current docs and created a framework for thought going forward
    *   I am experiencing the “unknown unknown” issues for which we may need broader community insight. There are some topics in Kubernetes security I’m not sure about, so we may need some expert consults, and would love suggestions from the team here for how to proceed. 
    *   Jennifer: Has some folks in mind who can help.
    *   Jared: Maybe Steve Perry would be a good choice? (I know he’s currently swamped, though). 
        *   More information needed on how we can help and what all the open issues are. 

**4/2/2019**

**10:30am Pacific **

New contributors



*   Rael Garcia (@raelga)
*   Angel Guarisma (@angelg)
*   Leslie Salazar (@lsalazar)	
*   Nathan Melehan (@nmelehan)
*   Andy Stevens (@astevens)

Updates/reminders



*   Ryan McGinnis (@ryanmcginnis) is PR wrangler this week
*   Q2 planning video is up: [https://www.youtube.com/watch?v=uIlU2vFJyqw](https://www.youtube.com/watch?v=uIlU2vFJyqw)
    *   Wins to highlight: 
        *   1.2 million page views per week! Up from 1 mil last quarter!
        *   1.14 Release (thanks Jim Angel and all the shadows!)
        *   2x number of localizations (French, Italian, German have been added!)
    *   [Notes from the planning session are here](https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/edit)
        *   Q1 goals graded, Q2 goals proposed
*   Spanish localization of the website almost ready
    *   New slack channel #kubernetes-docs-es created to start the Spanish localization of the Documentation
    *   Init PR started and channel started
        *   https://github.com/kubernetes/website/pull/13543

Agenda



*   (Paris Pittman) Editorial teams for blog review
    *   Who owns the content of the blog? 
    *   SIG-Docs didn’t call out the content, only the tooling in our charter
    *   What is the blog process in general? Doesn’t seem transparent
    *   PROPOSAL: Create an editorial team for blog review with an owner and lead as a sub-project of SIG-Docs
        *   Kaitlyn Barnard would be the lead
        *   Vision:
            *   Put together a team of editors with clear blog guidelines
            *   Rotations for blog editors
            *   AI: Put out a call to folks in SIG-Docs and SIG-Contribx to help
            *   AI: Create proposal for managing blog content, Jaredb, Katlyn, Paris, Cody, Zach  - Jared will set up initial meeting

**3/26/2019**

**7pm Pacific - Fourth Tuesday**

New contributors



*   

Updates/reminders



*   1.14 docs are live! Thanks and congratulations, Jim Angel (@jimangel), Barney Makonda (@MAKOSCAFEE), Cody Clark (@cody-clark), Jared Bhatti (@jaredbhatti), Naomi Pentrel (@npentrel), and Tunde Oladipupo (@simplytunde)!
*   Misty Linville (@mistyhacks) is PR wrangler this week.
*   Q2 planning meeting is Thursday, March 28, 6-9pm Pacific. [https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/](https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/) 

Agenda



*   (Zach C) New community role coming soon: New Contributor Ambassador
    *   Do we need issue triage captains? Context: [https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md) 
*   KubeCon Shanghai (June)
    *   How many people planning to attend? (5)
*   Doc Sprint 
    *   (Rui) depends on the backlog tasks
    *   (Zach) Will work with Julia to set questionnaire and find deadline for questionnaire
    *   (Julia) Will the CNCF provide the same incentive for attendance as Nov 2018? Zach will confirm.
    *   Zach: Will follow up with CNCF to make sure we’re funded for dev headcount for better refdoc generation: DEFER discussion for two weeks (April 9)
        *   Loop in Qiming Teng (K8s community is great; challenge is some libraries, not a simple problem to solve)

**3/19/2019**

**10:30am Pacific**

New contributors



*   David Watson (@dwat)
*   Barnie (shadows releases 1.13, 1.14)
*   Suresh Palemoni (joining from Dubai)

Updates/reminders



*   Jared Bhatti (@jaredbhatti) is PR wrangler this week
    *   Cody (@cody-clark) will cover first two days
*   New API review process: [https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md](https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md) 
*   Feedback request for K8s community etiquette cheat sheet: [https://github.com/kubernetes/community/pull/3213](https://github.com/kubernetes/community/pull/3213) 

Agenda



*   (Jim A) Update on 1.14
    *   Help needed?
*   (Craig/Zach C) Where will Windows docs live?
    *   What’s the best path? /Setup/
*   (Zach C) KubeCon Barcelona: We’ve got an intro/deep dive scheduled; do we also want time at the community day? 
    *   Deadline to decide is April 1
*   (Zach C) New community role coming soon: new contributor ambassador
    *   Do we need issue triage captains? Context: [https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md) 
*   (Chen Rui) Relative links in the website repo \
	[https://github.com/kubernetes/website/issues/13231](https://github.com/kubernetes/website/issues/13231)
*   (Dominik/Andrew) RBAC
    *   Blog post [Inside Kubernetes RBAC](https://medium.com/@dominik.tornow/inside-kubernetes-rbac-9988b08a738a)
*   (Dani Comnea) discuss the path for [https://github.com/kubernetes/kubernetes/tree/release-1.3/examples/celery-rabbitmq](https://github.com/kubernetes/kubernetes/tree/release-1.3/examples/celery-rabbitmq) example
*   Zach C/Suresh/Dani: Link to Q2 planning doc: [https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/edit?usp=sharing](https://docs.google.com/document/d/16po48F83oaQGC_BxeB7_k_RlhMzCdJD6wKYq4kpE-P8/edit?usp=sharing) 

**3/12/2019**

**10:30am Pacific**

New contributors



*   Lots of familiar faces :-) 

Updates/reminders



*   @cody-clark is PR wrangler this week
    *   ...but check your assigned PRs! https://github.com/kubernetes/website/pulls
*   Quarterly planning meeting on Thursday, 3/28, 6-9pm Pacific

Agenda



*   Q2 planning: co-chairs meeting this week to set agenda.
    *   What should we talk about?
    *   New stuff: new K8s features, docs coverage for new features, how new features fit in the big picture of docs (thanks, Brad and Jennifer!)
*   Content review
    *   SIG-level review, let’s talk about implementation details offline
*   Use `feature-state` shortcodes in future releases
    *   Context: [https://github.com/kubernetes/website/pull/12915#issuecomment-470768767](https://github.com/kubernetes/website/pull/12915#issuecomment-470768767)
    *   Valid feature states: alpha, beta, deprecated, stable \
[https://github.com/kubernetes/website/blob/master/layouts/shortcodes/feature-state.html#L1](https://github.com/kubernetes/website/blob/master/layouts/shortcodes/feature-state.html#L1) 
    *   Where’s the best place to document that developers should use `feature-state`?
        *   k/website PR template
        *   k/enhancements repo PR template?
        *   PR guideline checklist/cheatsheet
    *   Add “introduced” to feature set
*   Season of Docs (jaredb) - unfortunately, I can’t be at the meeting - but I want folks to know about this. 
    *   Information launched here: [https://developers.google.com/season-of-docs/](https://developers.google.com/season-of-docs/)
    *   Would be good for K8s to participate
    *   2 Admins and 2 mentors are needed
    *   Applications open in early april, close in late april
*   SigDocsSec (Zach A.)
    *   Our first meeting will be this week or next to get feedback on goals and set timelines...I am very excited!!!
    *   Possibly get a channel in Slack maybe? I don’t know who to ping on this or if it’s even worth it
    *   Starting with a survey of security concerns
    *   Who’s the audience, how do you target them effectively?
        *   (Jim Angel) Suggest adding actionable items (or links to them) to the end of each conceptual section
*   No kube-scheduler documentation (#12609)? (JD)
    *   Needed to run out early, but pls feel free to reach out to me if someone has some ideas. It’d be great just to be able to sit down with a sig-scheduling member to hammer something out
    *   Steve: Kube-scheduler docs are autogenerated; if content is obsolete, place to fix it upstream in k/k
*   Improving auto generation (Zach A)
    *   Can we smooth out this process overall?
    *   Build on Qiming’s existing work?
*   We’re using the free version of Travis?!
    *   Move to prow?
    *   Dani will talk to sig-test-infra about migrating to Prow

**3/5/2019**

**10:30 am Pacific**

New contributors



*   Dani Comnea
*   Suresh Kumar Palemoni

Updates/reminders



*   Zach C (@zacharysarah) is listed as PR wrangler this week -- Jennifer will check in when he's back (OOO today and tomorrow) and cover in the meantime
*   [Write the Docs speakers announced](https://www.writethedocs.org/conf/portland/2019/news/announcing-speakers/). Time to buy your tickets if you plan to attend (note conflict with KubeCon Barcelona)
*   U.S. is on Daylight time starting next week, so meeting times elsewhere around the globe that don't adjust (or don't adjust at the same time) will be an hour earlier
*   (Zach A. can't attend) SIG-Docs security working group is still going on the sign up link is here: [https://goo.gl/forms/RE0bwxw3axH20y2P2](https://goo.gl/forms/RE0bwxw3axH20y2P2)
    *   The goals of the group can be found in Slack, first meeting to hopefully occur this week

Agenda



*   (Dominik) Is Kubernetes really a declarative system? [[slides](https://docs.google.com/presentation/d/1ne0z2oH_x-ojQphLlxxT5EdjIx_pGp66ZaLR51rJZUg/edit?usp=sharing)]
    *   Fundamental Modeling Working Group -- continue discussion of how we disrupt the terminology of "declarative" Kubernetes effectively in the docs
    *   Working group meets Fridays @ 10 am Pacific, [Hangouts Meet](https://meet.google.com/spm-jota-twn) [[notes](https://docs.google.com/document/d/1VCMdsRGyCViWta7l1c7Hji77ha1_vSYqihrb4MrS96A/edit?usp=sharing)]
*   (Jim Angel) 1.14 and docs
    *   [http://bit.ly/k8s114-enhancements](http://bit.ly/k8s114-enhancements)
*   (Steve) for next week, reminder/report on issues in [GH Project for the quarter](https://github.com/kubernetes/website/projects/3)

**2/26/2019**

**7pm Pacific**

New contributors



*   None yet :-( 

Updates/reminders



*   Zach A (@zparnold) is PR wrangler this week.

Agenda



*   (Zach C) Q2 planning meeting: 28 March, 6-9pm Pacific
    *   I’ll send out invitations this week
*   (Ben Elder) Update on kind docs subdomain: “the kind experiment is going excellently so far and that hopefully we can scale this out pretty easily”
    *   Context: [https://github.com/kubernetes/k8s.io/issues/189#issuecomment-465884240](https://github.com/kubernetes/k8s.io/issues/189#issuecomment-465884240) 
*   Chen Rui: How’s 1.14 going?
    *   Cody gave an update, things are going well and according to plan.
*   Chen Rui: Will there be a doc sprint at Shanghai?
    *   Zach C: Yes, in the main track. But we can do a longer community day too.
*   Brad: Getting started on [https://github.com/kubernetes/website/issues/12333](https://github.com/kubernetes/website/issues/12333)
    *   Working on a PR, wants lots of feedback, will solicit feedback in Slack
    *   Seeking help/collaboration for another PR
*   June Yi: Local meetup next week
    *   Working on 1.13 branching
    *   Final iteration will be 8th (currently on 7th)
*   Tamao: Interested in accessibility
    *   Zach C: Will follow up with Rajie
    *   Tamao: Will follow up with Rajie’s imminent draft

        [Rajie] - Shared the draft with Cody Clark on Feb 27, 2019. Cody is helping in adding relevant examples to the guidelines for Level A compliance.


**2/19/2019**

**10:30am Pacific**

New contributors



*   Soheila (Munich)

Updates/reminders



*   Rajie is wrangling this week 

Agenda



*   Zach C: Set a date for our 2Q planning meeting: 3/28, 6-9pm
    *   Check in with Brad
*   Zach C: SIG Cloud Provider needs someone with good info architecture and organization skills to review their template for cloud provider information: [https://github.com/kubernetes/enhancements/pull/827](https://github.com/kubernetes/enhancements/pull/827)
*   Rui: Interested in talking about Transifex (translation software) with other localization teams [https://docs.google.com/document/d/12vUq51yeeAPlldM4P_aEY3VUSi37YVeVfD_TfzzkGGQ/edit?usp=sharing](https://docs.google.com/document/d/12vUq51yeeAPlldM4P_aEY3VUSi37YVeVfD_TfzzkGGQ/edit?usp=sharing) 
*   Zach A: Discuss potential for a sig docs security working group (I recall that I am carrying the flag for documenting security in more detail...and aggregating information related to KubeSec)
    *   There is a lot of information here, and I wanted to open it up to other contributors who are interested in participation.
    *   ng in security stuff. I have a rough implementation plan, but security is a team sport.
    *   +1 from Luc: adding support for security is a high demand
    *   +1 from Steve: there’s good security documentation, but it may be hard to find
    *   +1 from Jennifer: We need a better story for securing a cluster, and for security as a whole
*   Brad T: KubeCon Shanghai - Is anyone else besides me going? Are we doing a doc sprint or a main track intro/deep dive session, or a session at Contributor Day?
    *   Action Item:: Zach will write the main track request for Shanghai by Feb 26.
*   Jim Angel: Update on 1.14?
    *   Docs release playbook will be ready for review soon, merge-ready PR due by end of next week (Mar 1)
    *   Handbook WIP (to be “ready to merge” in the next few days): https://github.com/kubernetes/sig-release/pull/498
    *   Release notes: Dave Strebel, Jeffrey Sica

**2/12/2019**

**10:30am Pacific**

New contributors 



*   Remy Leone (French!)

Updates/reminders



*   Zach Arnold (@zparnold) is PR wrangler this week, and Rajie (@rajakavitha1) is shadowing

Agenda



*   Craig Peters: Guidance for [Windows Container docs](https://drive.google.com/open?id=1a2bRd7PZXygIEm4cEcCeLXpEqJ7opakP_j4Pc6AJVYA) - best method to convert a GDoc to properly formatted md
    *   Zach: Etherpad? Maybe, but GooDoc UX is really handy
    *   Jennifer: There’s a useful plugin for GooDoc to Markdown ([GD2M](https://chrome.google.com/webstore/detail/export-as-markdown/hbojhdcnbcondcdfpfocpkjkfkbnbdad?hl=en-US))
*   Juan Diego: Could we get some updated docs for the kube-scheduler?
    *   JD will write up an issue
    *   Zach will help boost signal to get 1:1 time with SIG Scheduling if needed
*   [Zach C] Match PRs to Issues
*   [Zach C] Issues in To Do project: https://github.com/kubernetes/website/projects/3

**2/5/2019**

New contributors



*   

Agenda



*   Announcements: We are losing Andrew (not altogether, but still :-() But he is going to awesome OSS docs work.
*   Mentorship (from last week)
*   Updates to content projects: [Steve] I'll be late, around 11:00. This is the last time I need to be late.
    *   @juandiegopalomino is updating the [ReplicaSet concept topic](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/). The [PR](https://github.com/kubernetes/website/pull/12409) is almost ready to be merged. This update brings several improvements to the topic, including an explanation of how owner references link Pods to a ReplicaSet. @codyclark has a [PR](https://github.com/kubernetes/website/pull/12409) that creates two new Ingress topics and updates one existing Ingress topic. @dtornow and I will work on an update to the [Deployment concept topic](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/). In the Deployment topic, we will include a prototype of a controller fact sheet. The idea is that each controller topic would have a fact sheet that gives the essence of the controller's behavior in a concise form.
    *   Controller fact sheets [Dominik]
*   Netlify [Jennifer] -- looks as though the Korean dev branch is now set up
*   [Andrew, Dominik] WIP: "Kubernetes in 15 minutes" presentation. Kubernetes is a robot. Or a genie. Or both!

**1/29/2019**

**10:30am PST**

New contributors



*   Alex Ball (tech writer from Canada)
*   Craig Peters, Microsoft PM (interested in docs, support for Windows containers/new node type)

Updates/Reminders/Housekeeping



*   Approvers, take a look at the [PR wrangler schedule](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda



*   [Steve] I'll be late to the meeting, around 11:00. If the meeting is still going when I connect, I'd like to ask about [this issue](https://github.com/kubernetes-incubator/reference-docs/issues/49). Also, I'd like to put in a request for people to take ownership of the issues in [this project](https://github.com/kubernetes/website/projects/3). Several of the issues have owners; thank you! We still need owners for these:
*   [Need top-level Controllers concept guide](https://github.com/kubernetes/website/projects/3)
*   [Update Deployment concept topic](https://github.com/kubernetes/website/issues/12082)
*   Let’s consider mentorship--actively mentoring first and new contributors; let’s revisit this next week (2/5)
*   [Zach] SIG chairs are looking at how to mentor our own replacements (for some eventual day). If you’re an approver who’s possibly interested in being a chair for SIG Docs, let us know!
*   [Zach] Netlify: Still working on site setup for Korean dev branch
*   [Craig] 

**1/22/2019**

**7pm PST - Fourth Tuesday **

New contributors



*   JD Palomino (joined us from KubeCon)
*   Seokho Son (welcome back!)

Updates/Reminders/Housekeeping



*   [Zach C] I’m back! 
    *   Creating new PR wrangler schedule this week (draft by Friday)
*   [Jim A] 1.14 Release team update
    *   Shadows selected (slack)
        *   codyclark
        *   jaredb
        *   barnie
        *   naomi
        *   simplytunde
    *   Main goal of creating playbook

Agenda



*   [Steve] Please take a look at these issues in [this project](https://github.com/kubernetes/website/projects/3). Several of the issues still need owners. (Brad will take #12081)
*   [Jim] @tfogo recommended getting in touch with the l18n owners early as part of the release cycle - do we have a preferred / recommended contact list?
    *   Owners file
    *   Sig-docs channel
    *   TODO [jim]: open issue, cc other localization members in advance.
*   [Jim] @bentheelder is looking for direction on sig-docs involvement with sub projects that are still “k8s community projects.” The main outcome is for Ben to create his website in a sig-docs approved manor. This approach can be used for other sub projects in the future. Depending on availability this can be postponed until next week.
    *   [Zach C] Loop in Alex Contini on theme development/portability
    *   Talk about sharing netlify in sig-docs-tools
    *   Can we create standards for standing up sig-specific project sites?
*   [Andrew] @Rui reported some broken translation links. Wanted to follow up on this since he's at the meeting.
*   [Seokho] Netlify Setup issue for dev-1.13-ko.*
    *   Zach C will set up a Korean branch 
    *   Look into granular permissions
*   [Rajie] Persistence of language selection? 
    *   Talk with sig-docs-tools about language persistence
    *   Side issue: Who owns this site? [https://kubernetes-csi.github.io/docs/Home.html](https://kubernetes-csi.github.io/docs/Home.html)
    *   [https://github.com/containernetworking/cni](https://github.com/containernetworking/cni) 
    *   @chenopis will contact Saad Ali
        *   Saad Ali said the content is still under development and needs to be rewritten, so it may make sense to leave them there for now. They need some help w/ content review and edits, so I suggested he come to one of the SIG Docs meetings.
        *   FYI, Michelle Au ([msau@google.com](mailto:msau@google.com)) is also involved in the CSI effort.

**1/15/2019**

**10:30am PST**

New contributors



*   Seokho Son 
*   Dan Roscigno 

Updates / Reminders / Housekeeping



*   [Jared] [PR Wranglers](https://kubernetes.io/docs/contribute/advanced/#be-the-pr-wrangler-for-a-week) needed - need to be an approver
    *   1/15: Brad 
    *   1/22: Tim
    *   1/29: Based on rotations

Agenda

    New items:



*   [Jim/Jared] 1.14 Launch kickoff
    *   Finalizing the 1.14 list in the next week
    *   Shadows - survey sent out - please respond!
        *   [https://www.surveymonkey.com/r/k8s-114-rt-shadows](https://www.surveymonkey.com/r/k8s-114-rt-shadows) 
*   Items from Q1 planning
    *   [Steve] I will miss today's meeting. But I have listed the writing tasks that I think should be our focus for Q1. See [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3). So far, two of the tasks have owners. Cody Clark will own updates to the Ingress topics. Jennifer will own updates to the kubeadm setup topic(s). The other four tasks are up for grabs. Of course, we can discuss and adjust the list.
    *   [Jim] localization issues with predefined english shortcodes
        *   [https://github.com/kubernetes/website/issues/12160](https://github.com/kubernetes/website/issues/12160)
        *   AI: talk to Andrew re: option 0 (consensus) - Jared
    *   [Jim] Can we merge the bump to the new hugo version? (.49 - .52)
        *   [https://github.com/kubernetes/website/pull/11552](https://github.com/kubernetes/website/pull/11552)

**1/8/2019**

**10:30am PST**

New contributors



*   

Updates and Reminders



*   Welcome back and Happy New Year!

Agenda

    New items:



*   [Jared] Meeting notes
*   [Shavi] whitespace on user journeys pages
    *   What's the status of entry pages redesign? (Andrew, Cody, ??) [PR 10795](https://github.com/kubernetes/website/pull/10795)
    *   AI: Shavi - will sync with Andrew on CSS location. 
    *   AI: Jared will sync with Andrew on User Journey pages, but will be looking at content 
*   [Jennifer] OWNERS for l10n: how do we want to handle the need for translation teams to submit PRs for their languages to files at root of repo? See [PR #11581](https://github.com/kubernetes/website/pull/11581) for an example issue.
*   [Dominik] ReplicaSet Controller model update
    *   OwnerReferences!
*   [Steve] I will miss today's meeting. But I have listed the writing tasks that I think should be our focus for Q1. See [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3). So far, two of the tasks have owners. Cody Clark will own updates to the Ingress topics. Jennifer will own updates to the kubeadm setup topic(s). The other four tasks are up for grabs. Of course, we can discuss and adjust the list.
*   [Jared] Who are the shadows for the upcoming launch? (Jim?)

 

   Ongoing/previous items:



*   [Misty] PR Wrangling, Configure stale issues bot
    *   Bot’s current configuration needs to be improved.
    *   Why do we stale-out issues after 90 days? Organization’s defaults
    *   In GH: A PR is a type of issue, stale-bot shouldn’t close these
    *   AI: We should create a doc that guides issue triage, response time, priorities
*   [Jim] 1.14 release meister / shadow
    *   Steven Augustus sent out last call for shadows: Cody Clark, Jared Bhatti are current shadows
    *   Shadows: Working to write up a guide for future release meisters
*   Community redesign: last call for feedback?  [https://cjyabraham.gitlab.io/community](https://cjyabraham.gitlab.io/community)
*   [Andrew] Kubecon Planning and Feedback 
    *   [Planning Document](https://docs.google.com/document/d/12lgT9usnk2vEKb9Fvnz2stXBZ35xvx8yEPTJDhwicic/edit#)
        *   Q1 planning - let’s check in on these regularly during meetings starting next year - Steve volunteers to work on this.
        *   AI: Jennifer will get an agenda up with the Q1 goals
        *   Overall feedback: Felt more concrete, let’s do a short one again for Q2 and a longer one for Q3-4. 
    *   Need to make a decision (in Jan) about pwittrock’s kubectl content. 
        *   Make sure it aligns with our overall content strategy
        *   It’s likely SIG-Docs will have some level of ownership in the future, let’s figure that out early. 
    *   Paris/Misty: We welcome more help at the Diversity lunch next year
        *   Jared: I’m interested in helping out at an upcoming conference
*   AIs:
    *   AI: We should create a doc that guides issue triage, response time, priorities
    *   AI: Jennifer will get an agenda up with the Q1 goals for our first meeting of the year.
