### 12/14/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

* Avinesh Tripathi, SIG ContribEx marketing, student
* Ayushman Mishra, student

**Updates/reminders**

* This week’s PR wrangler: @bradtopol
* Next week’s PR wrangler: _Holidays until next year!_
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)
    * AI Natali - Next year’s PR wrangler
    * Chris - this week’s docs PR wrangler shadow, may continue next week

**Agenda:**

* Release 1.24
    * **Current Status:**
    * **Updates**:
        * [Celeste] Dockershim removal ([project board](https://github.com/orgs/kubernetes/projects/67)):
            * Introduction to the project **🎉**
                * Meetings on the SIG Docs public calendar at the oh-so-pleasant time of Tuesdays @ 7-7:30am PST starting in January
                * Tuesdays 7am-7:30am PST meeting for special meeting of the dockershim removal project, invite is on SIG Docs public calendar
                * SIGs: Node, Release, Arch, Docs are highly involved
            * **Good first issues:** Please pick up an issue from [https://github.com/kubernetes/website/issues/30771](https://github.com/kubernetes/website/issues/30771) if you have time! Let me demonstrate what you need you tho&lt;3
                * This is time sensitive; if you can’t do it before 2022 please don’t pick it up!
                * It’s also easy though :)
        * [Rey] Nate Waddington is the 1.24 Docs lead, 1.24 will likely start on Jan 10 and release on April 19
        * [Rey] dev-1.24 branch is open
* Issues & PRs
    * [Avinesh] [https://github.com/kubernetes/website/issues/29649](https://github.com/kubernetes/website/issues/29649)
        * What next? Look into notification e.g. api-review label
    * [Tim B] [https://github.com/kubernetes/website/issues/30801](https://github.com/kubernetes/website/issues/30801)
        * (v1.23 release not listed on releases page)
* Discussion
    * [Jim] Auto-close PRs with need’s rebase older than “X” duration?
        * Thoughts on the ideas
        * Thoughts on the duration if the idea is accepted
        * [Jim] We can close PRs that haven’t had action or needs to rebase or has conflicts, author can re-open the PR.
        * [Tim] Everyone (org member or not) can do reviews, tell authors to rebase, everyone is empowered. Continue to write nice messages when we close PRs
        * [Natali] +1 to that, don’t be afraid to close PRs
    * [Jim/Arsh] New contributor role
        * [Jim] As the SIG Docs mentoring program, new contributor ambassador, help onboarding process. Will send out a more formal announcement
    * [Jim/Arsh] Office hours suggestion
        * [Jim] Proposed by Arsh, have a general/informal on how to contribute to documentation
            * Chris M: volunteer new contrib help / OH
            * Whitney: No matter the attendance should bring the group together near kubecon
            * Celeste: Going to events, specifically under represented groups, to encourage contributions.
            * Topic idea: GitHub navigation / help.
                * Maybe create / define pipelines for folks
                    * How to choose an issue
                    * Work with mentor
            * Topic idea: how to use labels
            * Encourage safety around github contributions (no permissions / nothing to break / etc)
    * [Rey for Divya] Under 7-day lazy consensus (started on Dec 7), APAC friendly meeting is proposed to be on the last Wednesday of the month at 5:30am UTC / 11am IST / 9:30pm PST / 12:30am EST. [Slack thread](https://kubernetes.slack.com/archives/C1J0BPD2M/p1638888987356000)
    * [Jesse] Would like to host a 1.23 Release Docs retro w/ SIG Docs, open on timing
    * [Brad] PR Wrangler Schedule and Shadow program update
        * Will be making the wrangler schedule for 2022 for half the year, this week is the first week of the shadow program. Chris Negus is shadowing Brad this week. Whitney scheduled to shadow on Jan 10
        * [Vitthal] Is there any process to volunteer to be a shadow wrangler?
        * [Brad] Find a gap and get in touch with the wrangler (or Brad) via Slack
        * [https://github.com/kubernetes/website/wiki/PR-Wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers)

### 12/7/2021 - IN SLACK!

### 11/30/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @zparnold
* Next week’s PR wrangler: @jimangel
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers) (shout out shadows)
    * [https://kubernetes.io/docs/contribute/participate/pr-wranglers/](https://kubernetes.io/docs/contribute/participate/pr-wranglers/)

**Agenda:**

* Release 1.23
    * **Current Status:** Yellow
    * **Updates**:
        * Tentative release date: Dec 7th
        * Docs complete deadline is today
        * 14 open PRs to dev-1.23
            * Seems mostly to be done, missing SIG thumbs up.
        * k/website is frozen on Monday 12/6
            * Comms will go out before that
    * **Tim:** Question? Does anyone need context?
        * Rey: Enhancements and/or graduations (alpha->beta) and require docs
        * How do you know what needs to be created?
            * Are issues created?
                * SIGs opt-in for enhancements & ask if they need docs
        * Oline: Do the release teams give a blurb, or how do you know what needs to be added.
            * Owner SIGs create the documentation changes
            * Major themes
            * Tim: Lean on SIGs for info / tech and docs help with consumption (technical level).
            * Chris: Surprised me, the writers do the writing, and it’s the KEP owners.
    * [Kubernetes 1.23 Enhancements Tracking](https://docs.google.com/spreadsheets/d/1P1J1QpayRmh2SNjs8T-wBCb6SgEOdWTRQ7MBol7yibk/edit#gid=813297075) spreadsheet
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/30684](https://github.com/kubernetes/website/pull/30684) (awareness)
        * [`content/ko/docs/reference/issues-security/security.md`](https://github.com/kubernetes/website/pull/30684/files#diff-d0ed0a2e9021208a29f126a94abf4c7934934753a799db877f3a1068d612d368)
    * Bring up during localization meeting
    * Ashish: Update that page (which page?)
* Discussion
    * Any interest in a docs sprint?
        * Abigail, Rey, Oline, Natali, Vitthal, Anubhav, Whitney (if working with video content or tutorials)
        * Topics?
            * 2017 austin - list of glossary terms, each took each one (bite sized)
            * Natali: Docs sprint, could be a video content?
                * If tackle on own
            * Redoing the tutorials
                * Well bounded (3rd party content rules)
                * Survey for tutorials
                * Rework / add tutorials
            * Site architecture (more advanced) - Chris N. interested in that (localizations).
            * Hugo clean up (advanced)
            * Some tutorials (WG naming) redis tutorials
    * [Divya]The Google questionnaire for APAC friendly meetings is now closed. We will be announcing the results shortly.
    * [Celeste] Dockershim removal is coming; I (+ the other CNCF tech writers) will be buzzing around doing work on this
        * [https://github.com/cncf/techdocs/issues/88](https://github.com/cncf/techdocs/issues/88)
        * Use OCI spec vs. docker
        * Downstream impact: In 1.24, the code is actually removed (and needs docs removed too).
        * First call to action, host discovery / scope call with co-chairs / tech leads to figure out the staffing required etc.
        * Chris Q? What is the actual impact for running docker containers
            * Celeste: Desire to remove all docker references, but main focus is what’s the MVP? It hasn’t been scoped yet.
    * Chris M.: Closed issue, what’s etiquette for opening

### 11/23/2021 US Holiday Week

### 11/16/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

* Oline Case - attended a couple weeks ago, finally got ramped up
* Grace Gude - kubecon! SRE at Fairwinds

**Updates/reminders**

* This week’s PR wrangler: Tim Bannister
* Next week’s PR wrangler: N/A US Thanksgiving holiday
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.23
    * [Nate W., @nate-double-u]
    * **Current Status: Yellow**
        * **Updates**:
        * Very little change since our previous meeting. Only a few PRs are not yet filed for those that need docs and are not either At Risk or Removed from Milestone.
        * Tracking 54 enhancements originally
            * 9 None required
            * 45 Needs Docs (20 of these are At Risk)
            * 7 Drafts
            * 10 Ready for review
            * 6 Merged
        * Integration branch is healthy.
        * Docs Placeholder PR deadline on Nov 18
        * Docs Ready for Review on Nov 23
        * Docs Ready to Merge on Nov 30
        * v1.23.0 scheduled to be released on Dec 7
* Issues & PRs
    * [Rey Lejano, @rlejano] [https://github.com/kubernetes/contributor-site/issues/202](https://github.com/kubernetes/contributor-site/issues/202)
        * [Oline] Look into the Community site to highlight the Contributor site
        * [Tim] Write up good-first-issue, mention @occase
    * [Rey Lejano, @rlejano]
      [https://github.com/kubernetes/website/pull/28454](https://github.com/kubernetes/website/pull/28454)
    * [Name, Slack]
* Discussion
    * [Divya Mohan, @Divya] APAC friendly SIG Docs meeting [questionnaire](https://docs.google.com/forms/d/e/1FAIpQLScUai5lwL2NOtiWBRas7gHYYMhEowNR0tsjsqnoa9KgG6o76Q/viewform?vc=0&c=0&w=1&flr=0&usp=mail_form_link): Please respond by EOD PST, 23rd November, 2021
        * Email went to spam
        * Jim to boost on slack
        * Not clear changes are for this meeting or APAC.
    * [Jim Angel, @jimangel] Congrats Shannon on [becoming a Reviewer](https://github.com/kubernetes/website/pull/30443)! **🎉**
        * Anyone else interested in becoming a Reviewer?
            * [Requirements](https://github.com/kubernetes/community/blob/master/community-membership.md#reviewer):
                * member for at least 3 months
                * Primary reviewer for at least 5 PRs to the codebase
                * Reviewed or merged at least 20 substantial PRs to the codebase
                * Knowledgeable about the codebase
                * Sponsored by a subproject approver
                * With no objections from other approvers
                * Done through PR to update the OWNERS file
                * May either self-nominate, be nominated by an approver in this subproject, or be nominated by a robot
            * [https://kubernetes.io/docs/contribute/participate/roles-and-responsibilities/#becoming-a-reviewer](https://kubernetes.io/docs/contribute/participate/roles-and-responsibilities/#becoming-a-reviewer)
    * [Name, Slack]
    * @chrismetz09 Mermaid blog for contributor-site actions moving forward. [https://github.com/kubernetes/contributor-site/pull/255](https://github.com/kubernetes/contributor-site/pull/255)
    * [Natali Vlatko, @nvlatko]:
        * Towards the end of the year, and we need to renew PR wrangler - do we need to refresh?
            * Yes. Challenges (approvers / localizations / etc)
            * Should we want to extend to reviewers?
                * Do reviewers want to be involved?
                * Temporary grants or shadowing
            * Who do you send the “hard ones to” ?
            * Nate: Shadowing (if there’s a pool). Nate is not an approver, but could pitch in and help out. Opportunity to help would have helped more.
            * Natali: +1 add in who’s shadowing and other folks can volunteer
            * Brad to create a schedule and add columns for shadow.
                * Nate: whole year? Quarterly?
                    * Brad: 2 half years
            * Abigail: +1 on shadows, wrangling idea is great. How many wranglers do we have?

**11/9 Slack Standup**

### 11/02/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

* Naren - also attended SIG ContribEx meeting, adding diagrams to Service page (first time attending Sig docs meeting 🙂🥰)
* Whitney - first PR
* Amit - looking to contribute
* Adrianna - looking to write a blog post
* Ashish - (mic issues)
* Leon - Linux admin new to k8s

**Updates/reminders**

* This week’s PR wrangler: Jim Angel (@jimangel)
* Next week’s PR wrangler: Qiming Teng (@tengqm)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.23
    * Docs status: Yellow
    * Integration branch is healthy
    * Tracking 56 issues
        * Yellow as we have very few placeholder PRs so far
        * 7 drafts so far (2 ready for review, 1 merged)
    * Code freeze coming up! 2 weeks, Nov 16th
    * Release team info
        * Shadows: [https://github.com/kubernetes/sig-release/blob/master/release-team/shadows.md](https://github.com/kubernetes/sig-release/blob/master/release-team/shadows.md)
        * Release team selection
          [https://github.com/kubernetes/sig-release/blob/master/release-team/release-team-selection.md](https://github.com/kubernetes/sig-release/blob/master/release-team/release-team-selection.md)
        * Role handbooks
          https://github.com/kubernetes/sig-release/tree/master/release-team/role-handbooks
* Issues & PRs
    * [Chris] Spoke with Tim regarding Mermaid-related blog
        * Where to open a PR?
        * [Divya]k8s.dev content can be mirrored to kubernetes blog
        * [Jim] Should open against k8s.dev even though not aware of how the mirroring process works
            * Mermaid.js = Javascript plugin running natively on the browser that build diagrams on fly
            * Relatively new to Kubernetes
            * Chris is boosting the efforts for Mermaid within the Kubernetes project
    * Where should non-code contributions be made?
        * k/website, k/community, k8s.dev (k/contributor-site)
        * Add a getting started link / group invite on k/community page (good to open an issue)
* Discussion
    * [Divya Mohan, @Divya] Around 100+ PRs (including aged) in the English language. [GH query](https://github.com/kubernetes/website/pulls?q=is%3Aopen+is%3Apr+label%3Alanguage%2Fen+-milestone%3A1.23+)
        * Different sizes of work based on appetite. Any volunteers?
        * Link to the thread on #sig-docs channel: https://kubernetes.slack.com/archives/C1J0BPD2M/p1635877576129400?thread_ts=1635877369.127000&cid=C1J0BPD2M
    * [Divya Mohan, @Divya] Working on potential slots/days for APAC-friendly meetings. Poll to be sent out by the end of this week.
    * [Divya Mohan, @Divya] Label command can now be used with refactor to indicate large refactoring changes like removing files & moving content from one page to another. Thank you Rey Lejano for leading the effort 🎉 See slack thread - [https://kubernetes.slack.com/archives/C1J0BPD2M/p1635784507115800?thread_ts=1634607120.424200&cid=C1J0BPD2M](https://kubernetes.slack.com/archives/C1J0BPD2M/p1635784507115800?thread_ts=1634607120.424200&cid=C1J0BPD2M)
        * Would this be for new pages, as well?
    * There are multiple versions, where do diagram changes fit in?
        * [Rey]If diagrams are for current version, submit it to main branch for k/website
        * [Rey] For further releases, role lead will do the cherry-picking & see if it needs to be updated for the release.

### 10/26/2021 Slack standup!

### 10/19/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

* Garima Negi
* Whitney Lee – VMWare
* Atharva Shinde
* Vitthal Kaul
* Federico (Karmavil) part of ES localization

**Updates/reminders**

* This week’s PR wrangler: ???
* Next week’s PR wrangler: ???
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.23
    * Docs Placeholder PR deadline is on Nov 18
    * 7 draft PRs open, 3 complete, 40 need placeholder PRs
    * Email will go out to the community on Thursday to remind folks about Docs Placeholder PR milestone
    * 1.23 release for December 7th
* Issues & PRs
    * [Tim B / @sftim] [Second level menu items are missing on docs](https://github.com/kubernetes/website/issues/29997#) PR #29997
        * Looking at improving this, same bug experienced by Rodolfo on the EN site
        * General feeling from the virtual room that this isn’t intended behavior, should be looking into a PR to address
    * [firstname lastname, email or @slackname] Add your suggested topic here and move this to the line below.
* Discussion
    * [Rey Lejano, @rlejano] New `refactor` label is almost ready, the `refactor` label indicates large refactoring changes e.g. files removed, content moved. Avinesh Tripathi is looking into how to automate notification to localization reviewers &/or SIG Docs leads/reviewers
        * New contributors could possibly help? Rey will update the SIG-Docs Slack channel when the label is ready
    * [Chris Metz, @chrisMetz] Diagrams using Mermaid update – link to Slack thread on a mock-up in the live editor: [https://kubernetes.slack.com/archives/C1J0BPD2M/p1634599198422800](https://kubernetes.slack.com/archives/C1J0BPD2M/p1634599198422800)
        * In the interest of making docs new contributor-friendly and our docs overall accessible
        * Mermaid: Available package within the org, using Markdown, with others able to collaborate on diagrams to make changes, contribute via a REST API as well
        * Kim Schlesinger: Beyond Block Diagrams (Digital Ocean) – talk at KubeCon, recommended to watch
            * [https://kubernetes.slack.com/archives/C1J0BPD2M/p1634599198422800](https://kubernetes.slack.com/archives/C1J0BPD2M/p1634599198422800)
        * Whitney suggests she could help contribute videos! Example: [https://youtu.be/BgrQ16r84pM](https://youtu.be/BgrQ16r84pM) – currently working at VMWare, used to make videos at IBM for Cloud
        * Nate: Separate repo for art/image assets for the site? Store the working files in this separate repo, going back and editing them later would be possible this way
    * [@arsh] Any docs related changes we need to be aware of for the kubectl out of tree KEP.
        * Similar kubeadm KEP - [https://github.com/kubernetes/enhancements/pull/1425](https://github.com/kubernetes/enhancements/pull/1425) [line 511 onwards]
        * More discussion for folks between SIG-CLI and SIG-Release needed to know of next steps around extra docs for here
        * Related? [https://github.com/kubernetes/website/pull/29847](https://github.com/kubernetes/website/pull/29847)
        * The above-mentioned KEP by Arsh will deal with auto generated content primarily, as opposed to hard-coded documentation that the above PR specifies (shared by Chris Negus)
    * [firstname lastname, email or @slackname] Add your suggested topic here and move this to the line below.

### 10/05/2021

### 10:30am Pacific (18:30 UTC)

**🎉 New contributors 🎉**

* Christina (she/her) – Mozilla

**Updates/reminders**

* Divya Mohan as the new SIG Docs co-chair 🎉
* This week’s PR wrangler: Taylor Dolezal (@onlydole)
* Next week’s PR wrangler: KubeCon week - but we don’t have a PR Wrangler for the week after that
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.23
    * Status: Green
        * Tracking 57 issues.
        * No blockers as of now (yaaay!)
    * [firstname lastname, email or @slackname] Add your suggested topic here and move this to the line below.
* Issues & PRs
    * [@arsh] NSA / CISA Kubernetes Hardening Guidance was finally [merged](https://github.com/kubernetes/website/pull/29791).
    * [firstname lastname, email or @slackname] Add your suggested topic here and move this to the line below.
* Discussion
    * [@arsh] Automate messages on PRs updating blog posts older than 2 years - [https://kubernetes.slack.com/archives/CJDHVD54J/p1633441645037800](https://kubernetes.slack.com/archives/CJDHVD54J/p1633441645037800)
        * AI: Create an issue [@arsh]
    * [@arsh] KubeCon next week!!
        * Docs sprint which was planned earlier has now been cancelled
        * You can [meet folks](https://kubernetes.slack.com/archives/C1J0BPD2M/p1632850983208800?thread_ts=1632850226.207400&cid=C1J0BPD2M) from SIG Docs at the [SIG Meet and Greet](https://kccncna2021.sched.com/event/luKT/kubernetes-project-sig-meet-and-greet) on Oct 13 at 1:30pm - 2:30pm PDT
        * [SIG Docs' Deep Dive talk](https://kccncna2021.sched.com/event/lV7N/kubernetes-sig-docs-a-deep-dive-jim-angel-google-chris-negus-aws-chris-metz-brad-topol-ibm) at KubeCon NA on Oct 13 5:25pm PDT
    * SIG docs localization group meeting [@bradtopol]
        * Weekly meetings for individual localizations subprojects.
        * Two new localizations - Hindi and Punjabi
            * Hindi localization tracking issue - [https://github.com/kubernetes/website/issues/29353](https://github.com/kubernetes/website/issues/29353)
    * [firstname lastname, email or @slackname] Add your suggested topic here and move this to the line below.

### 09/21/2021

### 10:30am Pacific (18:30 UTC)

### 09/21/2021

### 6:00pm Pacific (1:00 UTC)

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: Jim Angel (@jimangel)
* Next week’s PR wrangler: Karen Bradshow (@kbhawkey)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda**:

* Release 1.23
    * Status: Green
    * [Integration branch](https://github.com/kubernetes/website/pull/29658) is healthy
    * Docs team have started to reach out to enhancement owners about Docs Placeholder PRs
    * [State of the Release email](https://groups.google.com/g/kubernetes-dev/c/9TAMqOEV-Qk/m/0nn3vR40CAAJ) went out last Friday with a reminder for Docs Placeholder PRs and the Docs related deadline are around holidays so start work early
* Issues & PRs
    * [Tim] [29677](https://github.com/kubernetes/website/issues/29677) (disable Netlify drawer) - _for visibility_
    * [Natali] [29385](https://github.com/kubernetes/website/pull/29385) (big docs changes and how to align with loc-leads)
        * [Shannon]: Discussion issue [29649](https://github.com/kubernetes/website/issues/29649)
        * [Tim]: Can we announce a quota / ceiling “n” # of big changes per release cycle?
            * Shannon +1 but how to enforce?
        * [Jim]: For the discussion: potentially `/hold` for future releases.
        * [Avinesh/Shannon]: Could the bot use the PR size labels to tag loc-leads when over an X threshold? If yes, what threshold?
        * [Qiming] 2 kinds of big changes: new features and have to be localized by each team and re-factoring of existing documentation e.g. context moved from one page to another. Can we use new tag to label for this type of change or ping localization team to warn about huge change of re-factoring of documentation.
        * [Tim] We can add a re-factor label and will need to teach people to use it
        * [Rey] SIG Docs mentorees to add re-factor label
        * [Tim] That label info is not in git (it’s just in GitHub - there is an API though)
    * [Rey] [29706](https://github.com/kubernetes/website/pull/29706) Request for more reviews
    * [Rey] [29754](https://github.com/kubernetes/website/pull/29754) Request for a review, needs an lgtm
* Discussion
    * Divya Mohan, incoming SIG Docs co-chair, is working on the official SIG Docs meeting notes Google doc
    * Hindu L11n efforts

**_09/14/2021_** Async in Slack

**_09/7/2021_**

**_10:30am Pacific (18:30 UTC)_**

🎉 New contributors 🎉

* Natali Vlatko - Helping with SIG Auth / SIG Docs
* Arsh Sharma - Helping with SIG Arch / SIG Docs

Updates/reminders

* This week’s PR wrangler: Tim B (@sftim). - Arsh (to Shadow)
* Next week’s PR wrangler: Jim Angel (@jimangel)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.23
    * Enhancements freeze starts this Thursday at 23:59 PDT, currently at 67 enhancements opted-in, 51 need docs
    * First doc related milestone is Docs Placeholder PR deadline on Nov 18
        * Enhancement Freeze:     Thu September 9
        * Code Freeze:    Tue November 16
        * Docs Placeholder PR:    Thu November 18
        * Docs Ready for Review:    Tue November 23
        * Docs Ready:     Tue November 30
        * v1.23 Release:     Tue December 7
    * [Jim] Can we help with anything leveraging KubeCon?
        * Start communicating that the docs placeholder to be open after KubeCon
* Issues & PRs
    * [PR 29385](https://github.com/kubernetes/website/pull/29385) - refactor of a page with problematic layout, confusing wording, and multiple style/grammar issues.
        * Issue raised by @tengqm: PRs of this scale have substantial impact on L11n teams downstream. Suggestion is better heads up and planning to ensure those teams know what’s coming.
        * Shannon: What’s the process for starting a larger-than-usual change for not minor issues on a page? We should document that in the Contributor’s Guide as well. Also, should I delete the PR, or hold it until we have informed sig-docs-localization?
        * [Tim]: Can we announce a quota / ceiling “n” # of big changes per release cycle?
        * [Jim]: For the discussion: potentially `/hold` for future releases.
* Discussion
    * Divya Mohan Co-Chair nomination

**_08/31/2021_**: Async in slack

**_08/24/2021_**

**_6pm Pacific (01:00 UTC)_**

🎉 New contributors 🎉

* Louis from Singapore (lk_le) - help with French localization
* Ram (DevAdvocate) CloudFoundry foundation

Updates/reminders

* This week’s PR wrangler: @sftim - @reylejano
* Next week’s PR wrangler: @CelesteHorgan
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.23
    * Release 1.23 started this week
    * Jesse Butler is the 1.23 release docs lead
    * Shadow onboarding will start next week
    * 1.22 released a blog on upcoming deprecations, the blog was released post code-freeze. This will be added in the comms handbook and will continue in 1.23
* Issues & PRs
    * Issue when translating, some teams were moving docs around.
        * Moving large amounts of content causing issues
            * Example: [https://github.com/kubernetes/website/pull/28870](https://github.com/kubernetes/website/pull/28870)
        * Labels would help
* Discussion
    * Co-Chair email
    * Meeting times and options
        * Moved weekly meetings to bi-weekly
        * Before, when we had an APAC friendly meeting, we would cancel the regular meeting but changing to bi-weekly and if we skip the regular meeting then it can be 4 weeks between regular SIG Docs meetings. Looking for a better way to schedule the APAC friendly meeting
            * Options brought up: bi-monthly, last Thursday of the month based on UTC time (using first or last day of the month is easier)
            * Poll in Slack

**_08/17/2021_:** Async in slack

**_08/10/2021_**

**_10:30am Pacific (18:30 UTC)_**

🎉 New contributors 🎉

* Anubhav Vardhan
* Arsheel Sheikh
* Rajat Gupta

**Updates/reminders**

* This week’s PR wrangler: Savitha Raghunathan
* Next week’s PR wrangler: Brad Topol (deja vu - OOO?)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda**:

* Release 1.22
    * Complete! 🎉🎉🎉
    * Congrats to the 1.22 Docs Release Team:
        * Victor Palade (@pi-victor)
            * Aashish Nehete ( @Aashish Nehete)
            * Carlisia Thompson (@Carlisia)
            * Chris Negus (@Chris Negus)
            * Ritu Panjwani (@ritpanjw)
    * 1.23 led by Jesse Butler (@Jesse Butler)
        * Shadow application ends on 13th August
        * [https://docs.google.com/forms/d/e/1FAIpQLSf0uW8PoxNqzIypOwFs57ZFgf0-HUbu8-O0RcQeOuVEcNccKg/viewform](https://docs.google.com/forms/d/e/1FAIpQLSf0uW8PoxNqzIypOwFs57ZFgf0-HUbu8-O0RcQeOuVEcNccKg/viewform)
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/29247](https://github.com/kubernetes/website/pull/29247)
* Discussion
    * Chair transitions / rotations
        * Call for interest?
    * Move to Slack meetings (every other meeting?)
        * AI: Jim - to figure out
    * Contributor Summit NA (Los Angeles)
        * CFP and Registration now open!
    * KubeCon talk ideas?
    * [Rey] API reference doc generation during releases
        * 1.20->1.22
            * Issues generating each release.
            * Can the tech owners update the role handbook?
            * Transferable skills
            * AI: Rey to drive / contact
        * Tim: Can we connect golang folks to the process. (Jesse)
            * Containers would be a good solution vs. by hand
            * AI: Tim - introduce new folks to challenges
* Other
    * [Jim] Started uploading YouTube meeting videos
        * Sorry.
        * Limited to 15 per 24/hrs
        * Have 13 more (28 total!!)
    * [Jim] Meeting notes converted to markdown
        * Nate W: hackmd.io - connects to repo / searchability
        * AI: Jim Angel (check contribex)
            * Answer: Contribex isn’t pushing a specific format
    * Chrome warns on netlify preview? (92/windows)
    * Nate: links in chat, zoom goes away, is there a way to catch the chat?

**_08/03/2021_**

**_10:30am Pacific (18:30 UTC_)**

🎉 New contributors 🎉

**Updates/reminders**

* This week’s PR wrangler: Brad Topol
* Next week’s PR wrangler: Savitha Raghunathan
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda**:

* Release 1.22
    * Status: Green [GO]
    * Freezing the k/website repo will be done at 10 AM PDT
    * Need help tomorrow during the release to update the netlify config
    * Don’t understand which process applies for generating the api reference for kubectl & co.
    * Need approvals for release day activity
* Issues & PRs
* Discussion

**_07/28/2021_ - CANCELLED DUE TO CONFLICTS**

**_6pm Pacific (01:00 UTC)_**

~~🎉 New contributors 🎉~~

~~Updates/reminders~~

* ~~This week’s PR wrangler: @celesteHorgan~~
* ~~Next week’s PR wrangler: @zparnold (help needed!)~~
* ~~Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)~~

~~Agenda:~~

* ~~Release 1.22~~
    * ~~Status: Yellow~~
    * ~~4 PRs targeting enhancements left to merge after the 27th of July deadline, PR for integration branch fix and sync is open: [https://github.com/kubernetes/website/pull/29133](https://github.com/kubernetes/website/pull/29133) please approve when you get the chance.~~
    * ~~9 PRs total open against dev-1.22 branch, please see: [https://github.com/kubernetes/website/pulls?q=is%3Apr+is%3Aopen+base%3Adev-1.22](https://github.com/kubernetes/website/pulls?q=is%3Apr+is%3Aopen+base%3Adev-1.22) ~~
* ~~Issues & PRs~~
    * ~~[Tim B] [https://github.com/kubernetes/website/pull/29062](https://github.com/kubernetes/website/pull/29062) seems to be deprecating some uses of `.spec.nodeSelector` for Pods - are we communicating this deprecation effectively?~~
* ~~Discussion~~

**_07/20/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

* Purneswar Prasad

**Updates/reminders**

* This week’s PR wrangler: @tengqm
* Next week’s PR wrangler: ~~@zacharysarah~~ (I believe Celeste is stepping in, will double check [Jim])
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Status: Yellow
    * Integration branch is not healthy due to some conflicts, i will try to address that in a PR today.
    * As we head into the deadline for docs PRs ready for review, we have 6 drafts, 24 ready for review and 22 merged.
    * I’ve sent a friendly reminder to the 6 draft PR owners and will reach out to them tomorrow via slack if they don’t reply.
* Issues & PRs
* Discussion
    * [Chris] How to liaise with other SIGs for tech review (and similar) on PRs
        * Rey: first off, mention @kubernetes/sig-whatever-pr-reviews in a comment
        * Can also ask the SIG via their Slack channel - this is definitely OK
        * If still no response: find specific SIG participants (eg a technical lead) and ask them to help you find a reviewer
        * Rey’s last resort: use the social grapevine / find someone who works at a company with an interest in that SIG, ask them to help get things moving
        * Victor: for release documentation, SIGs can nominate a contact for docs liaison. 3 SIGs participating for v1.22 release.
    * [Tim] New contributor AMA

**_07/13/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: ~~@steveperry-53~~ * help needed (slack)
* Next week’s PR wrangler: @tengqm
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Status: Red
    * 3 enhancements need docs after docs placeholder PR deadline, we have reached out to the owners and are still waiting for feedback
    * 16 drafts, 13 ready for review, 20 completed
* Issues & PRs
    * (Jim) [https://github.com/kubernetes/website/pull/28522](https://github.com/kubernetes/website/pull/28522) - membership clean up / merging
    * (Chris N) [https://github.com/kubernetes/website/issues/28291](https://github.com/kubernetes/website/issues/28291)
    * (Chris M) Uniform way to add punctuation to docs.
    * (Chris M) Reviewing PR (PNG -> SVG), reformatted / punctuation.
        * Fig caption shortcode
* Discussion
    * (Jim) Call out for more help
    * (Chris M) Mermaid
    * Victor: Thank Tim B for releasing / blogs / etc!!!
        * Approvals: SIG docs -> SIG Release
        * Invest more time (SIG Docs)

**_07/06/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: **nobody** (US holiday)
* Next week’s PR wrangler: @steveperry-53 (US Holiday)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Status: Yellow
      Integration branch is unhealthy (feature gates PR introduced conflicts) will be fixed later this week during the branch sync.
* Currently 30 enhancements need docs, 12 drafts, 12 Ready for review and 11 merged.
* Will ping remaining enhancements this week to remind them about the 9th of July docs placeholder deadline.

	

* Issues & PRs
    * (Tim) [https://github.com/kubernetes/website/issues/25645](https://github.com/kubernetes/website/issues/25645)
        * Docs release team happy to endorse the automation
        * Some concerns about using a manually managed file (YAML format) rather than generating this automatically
        * Action: Victor to raise this at SIG Release meeting (2021-07-13)
    * (Chris M) [https://github.com/kubernetes/website/issues/28808](https://github.com/kubernetes/website/issues/28808)
        * (Anna) Just started looking at Mermaid for SIG Release, can share these
        * (Chris M) brainstorming via Google Docs?
        * (Tim) [https://mermaid-js.github.io/mermaid-live-editor/](https://mermaid-js.github.io/mermaid-live-editor/)
* Discussion
    * (Chris N / Tim) ~~Weekly~~ blog meetings
        * Scheduled(?) for 11:30am Pacific Tuesday (18:30 UTC)
    * (Victor) Release branch localizations
        * When do we localize changes for the next release?
        * (Tim) risk of merge conflict blocking _dev-x.yy_ merging into main - and then docs release team might not speak the language to resolve merge conflicts.
        * (Rey) not sure there were many such changes for previous releases

**_06/29/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @sftim
* Next week’s PR wrangler: _vacation_ (US Holiday)
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Status: **Green**
    * Integration branch healthy
    * 40 enhancements need docs (lots ‘At Risk’), 10 drafts, 10 ready for review, 7 merged
    * We are approaching the docs PR deadline with lots of enhancements ‘At Risk’
    * We will remind people of the deadline on 9th of July this week
    * SIG-Release also sent multiple emails to remind enhancement owners about the deadline.
* Issues & PRs
    * (Tim) [https://github.com/kubernetes/website/issues/25645](https://github.com/kubernetes/website/issues/25645)
        * Generate feature gates list from data
        * (Jim) best time to switch is fast-follow for a release
* Discussion
    * (Tim) Deprecations (API removals) blog article
        * [https://github.com/kubernetes/sig-release/discussions/1606](https://github.com/kubernetes/sig-release/discussions/1606)
        * (Chris N) yes - and good to be proactive
        * (Jim) Fundamentally a comms issue
        * (Jim) kubernetes API is happy to serve a GA API version even when apps are interacting with a deprecated API version
        * (Chris) should we have a “prepare for an upgrade (deprecations etc)” task page?
            * (Jim) include lesser-known gnarly upgrade-related details
            * (Jim) [https://github.com/rikatz/kubepug](https://github.com/rikatz/kubepug) reports expected API changes
        * (Victor) Blog article **might** clash with release notes (“_No, really, you must read this before you upgrade_”)
            * Maybe we should try to hyperlink / not duplicate
            * (Tim) can we try to use identical wording in both places?
        * (Chris) Common / clear terminology will really help
            * (Victor) “Deprecated _and removed_”
            * (Shannon) explain differences in style guide and recommend consistent terms to use
    * (Jim) Redis vs. MongoDB conversation
        * [https://github.com/kubernetes/website/pull/28530#issuecomment-865483779](https://github.com/kubernetes/website/pull/28530#issuecomment-865483779)
    * (Chris / Tim) ~~Weekly~~ blog meetings
        * (Jim) Different day, same time could work
        * (Tim) Would fortnightly meetings help attendance?
        * **(Tim) pick this up next meeting**

**_06/22/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @savitharaghunathan
* Next week’s PR wrangler: @sftim
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Chris N: [no formal status update]
    * Pi-victor: Retroactive update: status: GREEN. Integration branch is healthy again, 41 enhancements need docs, 10 drafts, 10 ready for review and 6 merged
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/28522](https://github.com/kubernetes/website/pull/28522)
        * Bring up at weekly meeting: interests for growing reviewer pool (for blog and docs)
        * Chris: what about blog meetings?
        * Tim: blog meetings may be on YouTube
        * Seokho: can we use labels to mark PRs as good for reviewers?
            * Maybe to show which PRs might suit someone doing their first review?
            * Tim: any contributor can review - let’s communicate this more!
        * Tim: time to clear up stale per-page reviewers? (listed in front matter)
        * Seokho: localizations have a smaller reviewer pool than English
            * Also use “PR wrangling” rota (Korean)
            * When assigned reviewers don’t respond, wrangler nudges or switches assigned reviewers
        * Qiming: reviewing guidelines are overlong and complex. Even existing reviewers cannot find details they already know!
            * Chris: how about a reviewing cheat sheet
            * Tim: mention reviewing cheat sheet in PR template?
* Discussion
    * Branch rename 🎉
    * Slack meetings
        * SIG Contribex holds **some** of their meetings in slack via threads in their own meeting channel: [https://kubernetes.slack.com/archives/C01M6BTU9D3/p1623764891025800](https://kubernetes.slack.com/archives/C01M6BTU9D3/p1623764891025800)
        * Do we want to try for a **subset** of our meetings? Reduce meetings? Is there more value in person or something we would miss in slack?
            * Rey: is that suggestion just for APAC?
            * Qiming: APAC monthly cadence works OK
            * Tim: could consider moving Pacific time meeting to monthly cadence as well?
            * Rey: enjoy attending weekly when time allows
            * Tim / Rey: not the same saying Hi to new contributors on Slack rather than Zoom
            * Rey: meetings are also an opportunity to “pull andon cord” / signal for help
    * Seokho: Korean Localization Team Update
        * Korean localization milestone
            * Current (dev-1.21-ko.5)
                * Milestone announcement: [https://github.com/kubernetes/website/issues/28472](https://github.com/kubernetes/website/issues/28472)
                * PR Freeze: 06/28
                * Merge target: 07/01
                * Outdated Korean docs in dev-1.21-ko.5: [https://github.com/kubernetes/website/issues/28473](https://github.com/kubernetes/website/issues/28473)
        * Change period for Korean Team meeting
            * biweekly -> triweekly
            * Korean: UTC 13:00 Thursdays (triweekly)
        * Election for new Korean l10n team leader
            * Seokho’s duty as the team leader will be finished on July (1 year)
            * Schedule (Tentative)
                * Election announcement (2020-06-25)
                * Candidate registration (~2020-07-11)
                * Vote and confirm (~2020-07-15)
    * List of contacts in other SIGs for review (Qiming)
        * Qiming: other communities provide lists of technical specialists / focal points.
        * Qiming: informally we might know who to task for particular SIGs but could we have a more formal list
        * Tim: maybe ask in #kubernetes-contributors on Slack?
    * Blog / release liaison
        * Rey: release blogging gained an opt-in process for v1.21 release where SIGs opt in to having an article tied to the release.
        * Rey: not sure about cadence of those meetings
        * Rey: most articles published after release, but reviewed before release
        * Rey: it’s a **time crunch** especially the main article about the release
        * Rey: There is a dedicated release comms team (for v1.22: Jesse Butler is leading that)
        * Tim: cadence post release typically leaves 1 day gaps between articles and goes out on weekdays (eg mon, weds, fri with tues and thurs as gaps)
        * Tim: can also have pre-release articles eg for v1.22 Ingress beta APIs removal being mentioned both before and after release

**_06/15/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @onlydole
* Next week’s PR wrangler: @savitharaghunathan
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Integration branch: branch not healthy, will submit a PR to fix conflict
    * Tracking 67/68 enhancement, of which 43 should need docs.
    * **43** enhancements need docs, 10 PR drafts, 8 ready for review and 3 merged
* Issues & PRs
    * [Tim] [https://github.com/kubernetes/website/issues/25645](https://github.com/kubernetes/website/issues/25645) (generate feature gates page from data)
        * Action item: Victor / release shadows to nominate someone to lead on getting a positive thumbs-up approval on this change from SIG Releaee
    * “Official Docs Image” ([https://github.com/kubernetes/website/pull/24387](https://github.com/kubernetes/website/pull/24387)) (awareness)
        * Any objections? MAKEFILE / tag changes ok?
    * [Proposed hugo archetype for blog](https://github.com/kubernetes/website/pull/27971) (awareness)
    * Make an exception to back-port this fix to release-1.17? [https://github.com/kubernetes/website/pull/28399](https://github.com/kubernetes/website/pull/28399)
        * Usability impacting? Yes
        * Critical? No
        * (first time contributor)
    * API translations inflight [https://github.com/kubernetes-sigs/reference-docs/pull/232](https://github.com/kubernetes-sigs/reference-docs/pull/232) (awareness)
* Discussion
    * [Shannon]
        * Scope of changes / SMEs / review scope and comments
        * Cross cutting sigs (apps / api machinery)
        * [Tim] identifies liaison (SIG(s)) and should add members for lgtm.
        * Anyone available to help pair for sigs etc?
            * Tim
    * [https://github.com/kubernetes/website/pull/27991](https://github.com/kubernetes/website/pull/27991) Merged!
        * [https://kubernetes.io/docs/home/#contribute-to-the-docs](https://kubernetes.io/docs/home/#contribute-to-the-docs)
        * [https://kubernetes.io/docs/contribute/analytics/](https://kubernetes.io/docs/contribute/analytics/)
    * Maintainer Track session at KubeCon + CloudNativeCon North America 2021 (Due July 6)
        * Do we want to do? Volunteers?
        * **35-minute session** be a combination of an introduction and deep dive
            * **The Intro portion** should include some slides and a demo to explain what your activity does and why people should care.
            * **The Deep Dive portion** is for contributors, maintainers, and other attendees with some pre-existing familiarity with the activity. We recommend that this should be a working part focusing on technical discussions and debates, as opposed to just processes (e.g., should the API be changed, not what's the best time for monthly conference calls). (Recommended: 20-minutes)
            * Volunteers: Brad, Jim, Chris Metz, Chris Negus
                * [Tim] Can help with the actual recording part if needed
            * [Tim] What would you like to see in such a video?
                * [Shannon] Updating reference doc (and read API docs) +100 (everyone ;-)
                * Walk repo [Jim]
            * Double check on recording
    * Kube Doc sprint at Kubecon NA [Brad]
        * Jim and Brad were able to get us a room and on the schedule for KubeCon NA on Monday Oct 11 (This is a Pre-Event programming Day used to host co-located events). We currently have a room for the full day
        * No “summit” but can have space on pre-event schedule. Kube Docs sprint, room for full day.
        * [brad] Coordinate with the video track
        * Who would like to help us lead the session?
            * Chris Negus, Chris Metz Jim, Brad
        * Any thoughts on what should be the focus for the doc sprint? What should be on the agenda? Do we want a full day? Or halfday?
            * How to do git / get up and running
            * First PR / commit
            * On boarding to docs
            * API ref docs?
            * Walk release builds?
            * Walk repo?
        * [Jim] Let’s wrap a little after lunch
        * [Tim] Open to remote folks too.
        * TL;DR: full day to fill with agenda.
    * kubectl docs (updates? [Tim / Chris N])
        * Tim / Chris catching up
        * Chris, talking to Eddie around issues kind lists. We don’t document “`kind: List`”
        * Brad: Operator SDK had issues with it (can help with more details / mark)
    * Housekeeping:
        * [Jim] Update youtube playlists, archive notes, and other clean up.

**_06/08/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @kbhawkey
* Next week’s PR wrangler: @onlydole
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Integration branch: HEALTHY
        * PR due to merge
    * 45 enhancements need docs, 9 drafts PRs, 7 reviewable, 3 merged
        * plus 1 not on tracking sheet
    * No big challenges
* Issues & PRs
* Discussion
    * Stale content (Abigail)
        * There is a somewhat dated mechanism for running CI checks against the YAML (etc) manifests inside the /examples path (Tim).
        * Taylor (@onlydole) interested in helping on this front and refreshing any tooling/workflows.
    * Blog
        * (Tim) Nearly all contributions are welcome! Anything relevant to the audience of Kubernetes end users, contributors and interested others
            * Obvious plugs / advertising is disallowed
        * (Tim) Looking for anyone keen to help organize / run a weekly SIG Docs blog subproject meeting, and similar tasks.

**_06/01/2021_**

**_10:30am Pacific (18:30 UTC)_**

**🎉 New contributors 🎉**

**Updates/reminders**

* This week’s PR wrangler: @kbarnard10
* Next week’s PR wrangler: @kbhawkey
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

**Agenda:**

* Release 1.22
    * Chris - latest from victor: 67 enhancements / 8 drafts / 3 PRs have been merged
    * Friday to sync branch
    * Vaccine shots!
* Issues & PRs
* Discussion
    * kubectl docs [@eddiezane]
        * [https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700](https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700)
            * Next steps:
            * ID owner from each sig
            * Track in quarterly meeting
            * Duplicate / deprecate / track
            * Chris Negus Volunteer w/ Irvi / Tim
        * Reference / FAQs need to be brought over too.

**05/25/2021**

**18:00 Pacific (01:00 UTC)**

🎉 New contributors 🎉

* Jai Govindani - first time hopping on!

Updates/reminders

* This week’s PR wrangler: @jimangel
* Next week’s PR wrangler: @kbarnard10
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.22
    * integration branch is healthy
    * 47 enhancements need docs
    * PRs so far: 9 ready for review, 6 drafts and 1 merged
* Issues & PRs
    * master → main branch rename (reminder)
        * [https://groups.google.com/g/kubernetes-sig-docs/c/s98sr2appP4/m/XleflmLLAgAJ](https://groups.google.com/g/kubernetes-sig-docs/c/s98sr2appP4/m/XleflmLLAgAJ)
        * [https://www.kubernetes.dev/resources/rename/#anytime](https://www.kubernetes.dev/resources/rename/#anytime)
* Discussion

**05/18/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Shatakshi G

Updates/reminders

* This week’s PR wrangler: @irvifa
* Next week’s PR wrangler: @jimangel
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.22
    * Week 4 - Started pinging enhancements owners for docs placeholder PRs
    * Tracking 54 (55th incoming) enhancements out of 66
    * Open already: 1 Draft, 7 ready for review
    * Status Green
* Issues & PRs
    * master → main branch rename (dates) ([https://groups.google.com/g/kubernetes-sig-docs/c/s98sr2appP4/m/XleflmLLAgAJ](https://groups.google.com/g/kubernetes-sig-docs/c/s98sr2appP4/m/XleflmLLAgAJ))
        * Victor: config.toml for dev branch should reference main not master
        * Victor: release scripts will need checking
        * Tim: this rename process will benefit from a co-ordinator / champion
        * Jim: (in Groups) will lead
    * Analytics dashboard PR - [https://github.com/kubernetes/website/pull/27991](https://github.com/kubernetes/website/pull/27991)
        * Abigail: review still needed
* Discussion
    * (from last week) Reference docs toil? [@sftim / @jimangel]
        * Rey: +1 to reducing toil
        * Rey: past few releases have had API docs as a bit of a hurdle
        * Rey: there is tacit knowledge in how to generate these that is not passed on
        * Tim: Docs generation could (should?) be run within a container
        * Victor: interested in helping track work here
    * (from last week) Moving up the contributor ladder [@sftim]
    * (from last week) kubectl docs [@eddiezane]
        * [https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700](https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700)
    * Feature gates docs generation from source code [@pi-victor]
        * Easy to miss
        * Current process is manual
        * Feature gates upstream in `kube_features.go`
        * Tim:
    * Chris: scheduling and pre-emption discussion
        * Tim: approving-with-feedback reviews are great where feasible

**05/11/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Eddie Zaneski
* Carlisia Thompson
* Aldo Culquicondor
* Ritu Panjwani

Updates/reminders

* This week’s PR wrangler: @savitharaghunathan
* Next week’s PR wrangler: @irvifa
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.22
    * K8s org access for shadows complete
    * Two shadow onboarding sessions last week
    * Tracking ~69 enhancements so far with more to come this week (looks like enhancements freeze will be extended to next week)
* Issues & PRs
    * Additional discussion around line lengths [https://github.com/kubernetes/website/issues/27657](https://github.com/kubernetes/website/issues/27657)
    * General awareness around the releases page [https://github.com/kubernetes/website/pull/27929](https://github.com/kubernetes/website/pull/27929)
    * Scheduling / eviction / preemption (@sftim / @shannonxtreme)
        * [https://github.com/kubernetes/website/pull/27739](https://github.com/kubernetes/website/pull/27739)
        * [https://github.com/kubernetes/website/pull/27947](https://github.com/kubernetes/website/pull/27947)
        * We have a new page on eviction (manual using API + kubelet). Also including preemption.
        * How do we classify preemption in respect of eviction (Shannon, Aldo)
        * Brad: can help connect with SIG Scheduling chair Huang
        * Aldo: preemption not kubelet-style eviction. Do we expand “eviction” to mean any kind of (forced) Pod removal?
        * Today, eviction and pre-emption separated. Refactoring aims to bring these concepts into one page.
        * Aldo: node-initiated eviction and API eviction should be separate pages
        * Action: Move to slack convo and bring up next week if not resolved
    * website-milestone-maintainers and sig-docs team cleanup (@pi-victor)
        * Prow k/website config to apply milestone
    * Pretty please review and approve my PR for changing config.toml for dev-1.22 ([https://github.com/kubernetes/website/pull/27869](https://github.com/kubernetes/website/pull/27869))
    * Adding analytics dashboard - [https://github.com/kubernetes/website/issues/27360](https://github.com/kubernetes/website/issues/27360) (Abbie)
        * Add google docs leads group (if not, add individual owners)
        * How to host the dashboard?
            * Abbie - add google group to data studio report, add page with hyperlink and information around the dashboard
            * Jim - sending Abbie google group to add
* KubeCon Sessions
    * [Writing for Developers: Take your Project Docs to the Next Level](https://sched.co/iE3p)
        * Content on managing docs
        * [Kohei] Thanks Celeste!
    * [Panel: Your Path To Non-code Contribution In The Kubernetes Community](https://sched.co/iE1P)
        * Covering all non-code contributions (beginner level)
        * [Kohei] We had around 400 attendees (both real time and on demand)
* Discussion
    * Reference docs toil? [@sftim / @jimangel]
        * Bump for next week
    * Moving up the contributor ladder [@sftim]
    * kubectl docs [@eddiezane]
        * [https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700](https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700)

**05/4/2021**

**10:30am Pacific (18:30 UTC)**

* No Agenda Meeting (KubeCon EU)

**04/27/2021**

**17:00 Pacific (01:00 UTC)**

🎉 New contributors 🎉

* Jihoon-Seo

Updates/reminders

* This week’s PR wrangler: @bradtopol
* Next week’s PR wrangler: _nobody on rota_ (KubeCon)
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.22
* Issues & PRs
    * [sftim] [https://github.com/kubernetes/website/issues/27657](https://github.com/kubernetes/website/issues/27657) (see issue comments about motivation - could be a discussion for the next meeting section)
        * 80 / 120 chars
        * Does it impact hugo?
            * Even if you break lines, the paragraph remains the same.
        * Tengqm: Tested and merged into HTML
        * TODO: test / confirm no impact to widescreen and bring up next week for discussion.
        * +1 korean localization - Seokho
    * [sftim] [https://github.com/kubernetes/website/pull/27637](https://github.com/kubernetes/website/pull/27637) (new wrangler schedule needed)
    * [Tengqm] Reference generation, including kube components and config APIs.
        * 2w ago it was on agenda / complaints about refgen not being smooth
        * Using gomod, to prevent problems.
        * TODO: make sure all videos are up to date / review if inaccurate
* Localization
    * Korean team
        * about to finish dev-1.21-ko.1 (will be merged to master)
        * New org. member Jihoon-Seo.
* Discussion
    * [sftim] Thank you @zacharysarah

**04/20/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @kbhawkey
* Next week’s PR wrangler: @bradtopol
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.22
    * [Victor] Week 0
    * Release cadence for 2021 not 100% settled
    * Victor selecting docs shadows for v1.22 release
    * Week 1 next week
* Blog
    * [@sftim] Blog team looking for reviewers
    * Blog PRs are labelled area/blog
    * Some folks will prefer to work with a Google Doc or similar
    * [https://kubernetes.io/docs/contribute/new-content/blogs-case-studies/#guidelines-and-expectations](https://kubernetes.io/docs/contribute/new-content/blogs-case-studies/#guidelines-and-expectations)
* Localization
* Issues & PRs
    * [chrisnegus] Discuss new proposed Production environment page ([PR 27466](https://github.com/kubernetes/website/pull/27466))
* Discussion
    * [christophermetz] “Control plane” vs. “Control plane node” in documentation
        * (carried over from 2021-04-13)
        * Formerly “master” or “master node”

**04/13/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @zparnold
* Next week’s PR wrangler: @kbhawkey
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * 1.21 Release Docs Retrospective
        * What went well
            * Moving the Docs Placeholder PR deadline after code freeze
            * Getting pre-release tasks for next release done early
            * Reviews were quick from SIG Docs
            * [Anna] Rey did a fantastic job!!! 🎉🎉
        * Can we improve process
            * Reference documentation generation. This release and prior release we needed assistance from SIG Docs
                * More ref docs coming in (Qiming added new docs). All good, but means more toil. Also improvements arrived from Philippe M.
                * [Tim] AI: Next meeting, let’s talk about docs toils (ref) and avoid mistakes
                * [Rey] How do we test? What do we expect? What do we review?
            * [Anna] Do we need to collect primary contact for the enhancement docs?
                * [Rey] to bring up in official retro, maybe an opt-in approach? Might not work for docs if enhancement is opt-in.
                * [Tim] Can we make a “docs liaison” to assist with capturing documentarians as well as tracking docs?
                  With this mechanism, SIGs could opt in to nominating a liaison lead: someone SIG Docs can contact about documentation changes relating to the upcoming release.
            * [Anna] Can we just tag the localization owner github groups and use github discussion to notify?
                * [Rey] +1 to tagging localization owner github group and starting a github discussion at the start of the release cycle to start notifications to localization groups
                * [Jim] Is there still value in this step?
                    * [Rey] good for announcing freezes / timelines
        * Can we improve the playbook or add/remove stuff from the playbook
            * More build errors - require squashing commits or add prow label to squash commits.
                * Add label after a few days
                * [Tim] on dev branch, squash, before
                * [Tim] advantages to rebasing too (against target branch)
                * [Jim] Bring up to SIG Docs as we find them for future releases
            * Creating the release branch right before the repo freeze
            * Doing a branch sync the morning before repo freeze
                * Still confirm that the dev-1.xx future branch and the release-1.xx current branch are in sync with the master branch
            * Remove patch versions in the config.toml?
                * [Tim] We might be able to automate keeping the patch version current.
                * [Jim] Work on releases page might be relevant here.
                * [Jim] Bring into quarterly planning.
            * Linking to the release notes instead
                * [Tim] relnotes that look like a website, ok with resolution short term, but long term bring into site format (relnote maybe too?)
                * [Jim] Bring into quarterly planning.
            * Update on finding the last commit hash of the previous release
                * New, better, process.
                * [Tim] It’s always parent 1 of the merge commit that added the _dev-x.yy_ branch into the live website.
* Release 1.22
    * Kicking off. Week 1 on 4/19
* Blog
    * [@sftim] upcoming articles
        * A lot upcoming needs help (current tim / bob)
    * [@sftim] Blog team looking for reviewers
        * Have a backlog, need help! - see [https://kubernetes.slack.com/messages/sig-docs-blog](https://kubernetes.slack.com/messages/sig-docs-blog) =
        * Need process for helping / sparse
* Localization
    * [Seokho] FYI: Started first Korean localization milestone for v1.21 (dev-1.21-ko.1)
* Issues & PRs
* Discussion
    * [@leilajal] [Server Side Apply]( https://github.com/kubernetes/website/pull/27379): We are curious what the best format is to have these stories (maybe a blog-post series).
        * [Leila] SSA how to have best practices for documents each release. Best practices. Most likely not a single PR for review, but more content for relevant info.
    * [Tim] Control plane vs. Control plane node in documentation
        * Control plane node is valid for some clusters
        * Control plan is valid for all clusters
        * Context is important
        * WG-Naming is a good
        * [Jim] move to slack or next week for further discussion

**04/06/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @zacharysarah
* Next week’s PR wrangler: @zparnold
* Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] Docs
        * **Current Status: Yellow-Red**
        * [Integration branch](https://github.com/kubernetes/website/pull/26153) is not healthy
            * dev-1.21 branch sync has build errors
                * [Slack thread](https://kubernetes.slack.com/archives/C1J0BPD2M/p1617719678096000)
                * Branch sync with failed build on container.md [PR 27432](https://github.com/kubernetes/website/pull/27432)
                    * container.md was removed in [PR 26413](https://github.com/kubernetes/website/pull/26413) because you can find it under pod-v1.md
                * Philippe: 2 conflicting PRs are [26481](https://github.com/kubernetes/website/pull/26481) and [26413](https://github.com/kubernetes/website/pull/26413)
                    * Philippe: [26481](https://github.com/kubernetes/website/pull/26481) should be applied before [26413](https://github.com/kubernetes/website/pull/26413)
                    * [26481](https://github.com/kubernetes/website/pull/26481) was applied to master and [26413](https://github.com/kubernetes/website/pull/26413) was applied to dev-1.21 so when we merge master to dev-1.21 then it's the reverse order on what it should be.
                    * Rey: we can revert [26413](https://github.com/kubernetes/website/pull/26413) in dev-1.21, merge wmaster to dev-1.21 to apply [26481](https://github.com/kubernetes/website/pull/26481) then merge [26413](https://github.com/kubernetes/website/pull/26413) to dev-1.21
                * Tim: To create a revert / sync PR for a fix
        * [PR 27248 Reference Docs](https://github.com/kubernetes/website/pull/27248) has an approve, needs an lgtm
        * This week:
            * Create release branch
            * Update Netlify (assistance from a SIG Docs chair)
            * Inform localization teams about freeze and timelines
              April 7 (day before release)
            * Freeze the k/website repo
            * Merge master branch into the dev-1.21 branch and release-1.20 branch
              April 8 (release day)
            * Merge [integration branch](https://github.com/kubernetes/website/pull/26153)
            * Publish the release blog post
            * Create release with tag
            * Unfreeze the k/website repo
            * Close the 1.21 milestone
* Blog
    * [@sftim] [https://github.com/kubernetes/website/pull/27369](https://github.com/kubernetes/website/pull/27369) (PSP deprecation)
    * [Divya] 1.21 Release blog + Feature blog cadence for review up on Slack thread: [https://kubernetes.slack.com/archives/CJDHVD54J/p1617467226029800](https://kubernetes.slack.com/archives/CJDHVD54J/p1617467226029800)
        * 3 blogs to be published this week - Release Blog, PSP Deprecation Blog, Cron job goes GA.
    * Need clarification on blog review process
    * Tim: Get post out today
* Localization
    * [bradtopol] The Korean localization team would like to use the official zoom account for their weekly meeting (UTC 13:00 Thursdays Biweekly) similar to how the Spanish localization team uses it.
      Can leadership approve this?
    * Eventually we need to worry about meeting times conflicting.
        * How should we manage this?
        * Should we ask for an extra zoom account eventually?
    * [Jim] We might want to discuss how we can automate Zoom meeting setup etc / sharing out meeting links. Current process is very manual / toil-like.
    * No issue providing more meetings. SIG Docs TLs & chairs have access to schedule these and provide links. Also no issue providing a second Zoom account when that proves necessary.
* Issues & PRs
    * [#27248](https://github.com/kubernetes/website/pull/27248) for dev-1.21, refgen, needs lgtm (maybe needs verification)
    * [https://github.com/kubernetes/website/pull/27129](https://github.com/kubernetes/website/pull/27129)
        * AI: (Jim) review
* Discussion
    * [Data Studio](https://github.com/kubernetes/website/issues/27360#issuecomment-814255033)
    * [leilajal] questions ([#wg-api-expression](https://docs.google.com/document/d/1vXqKm_K__ySl1iRxnbdZM5DM6i5iqYF3lduTf9jE0mY/edit#heading=h.lch4dq1pplno)) [SSA](https://kubernetes.io/docs/reference/using-api/server-side-apply/)
        * Request: smaller group to discover best practices
            * Jim: great to bring up here (sig docs weekly)
        * PR collection
        * Guidelines

**03/30/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @tengqm
* Next week’s PR wrangler: @zacharysarah
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] Docs
        * **Current Status: Yellow**
        * [Integration branch](https://github.com/kubernetes/website/pull/26153) is healthy
        * Docs Complete deadline is tomorrow
        * Doc PR stats: we are currently tracking 41 docs out of 49
            * Docs merged: 34
            * Ready for review: 7
            * No docs required: 8
        * Appropriate SIGs have been pinged for tech reviews
        * Next week:
            * Create release branch
            * Update Netlify (assistance from a SIG Docs chair)
            * Inform localization teams about freeze and timelines
              April 7 (day before release)
            * Freeze the k/website repo
            * Merge master branch into the dev-1.21 branch and release-1.20 branch
              April 8 (release day)
            * Merge [integration branch](https://github.com/kubernetes/website/pull/26153)
            * Publish the release blog post
            * Create release with tag
            * Unfreeze the k/website repo
            * Close the 1.21 milestone
* Blog
    * [@sftim] Release imminent / will check in with blog subproject on Slack
* Issues & PRs
    * [Rey] Service Internal Traffic Policy feature [https://github.com/kubernetes/website/pull/27088](https://github.com/kubernetes/website/pull/27088)
    * [Rey] config.toml for 1.21 [https://github.com/kubernetes/website/pull/27223](https://github.com/kubernetes/website/pull/27223)
    * [Rey] Reference docs for 1.21 [https://github.com/kubernetes/website/pull/27248](https://github.com/kubernetes/website/pull/27248)
        * [Jim] Chicken / egg problem (update handbook and do it after release for tag?)
    * [Rey] Prepare 1.21 API ref docs, moves definitions, deep indentations for nested field, [https://github.com/kubernetes/website/pull/26413](https://github.com/kubernetes/website/pull/26413)
        * [Tim] Deprecate in 1.22, remove in 1.23 (and announce / coms)
* Discussion
    * [Abbie] Google Data Studio for localizations: [https://support.google.com/datastudio/answer/6283323?hl=en](https://support.google.com/datastudio/answer/6283323?hl=en)

**03/23/2021**

**17:00 Pacific (01:00 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @jimangel/@tengqm
* Next week’s PR wrangler: @tengqm
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey, Docs]
        * Status: Yellow
        * **Doc Ready for Review deadline is tomorrow**
        * [Integration branch](https://github.com/kubernetes/website/pull/26153) is healthy
        * Doc PR stats: we are currently tracking 39 docs out of 49
            * Docs merged: 18
            * Ready for review: 20
            * In draft: 1
            * No docs required: 10
        * One of the draft Doc PR owners is working on their Doc PR this evening, the other has not yet replied
* Issues & PRs
* Discussion
    * [Jim / Abbie] Google Data Studio for localizations: [https://support.google.com/datastudio/answer/6283323?hl=en](https://support.google.com/datastudio/answer/6283323?hl=en)
    * [Jim] Combine APAC and WG localization?
        * Brad: WG is more focused and weekly is more open.
        * Kohei: Two purposes, keep separate
        * Brad: New time for APAC meeting if participation is dragging.
    * Kohei: Chatting at KubeCon / panel

**03/16/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @CelesteHorgan
* Next week’s PR wrangler: @steveperry-53
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Divya, Comms] Upcoming Release blog review deadline **25th March, 2021**
        * Needs to be published on **8th April, 2021 [Target date for 1.21 Release]**
        * Release Logo + associated blurb will be added by Nabarun towards the end of the cycle.
    * [Divya, Comms] 8 feature blogs tracked so far
        * Tracker sheet: [https://docs.google.com/spreadsheets/d/1-rFkGmpyDN39gY2M_RX6Ugs_AIig9GXw1FkRLWWmWfw/edit#gid=0](https://docs.google.com/spreadsheets/d/1-rFkGmpyDN39gY2M_RX6Ugs_AIig9GXw1FkRLWWmWfw/edit#gid=0)
        * **Will be ready for review by 31st March, 2021**
        * Released in staggered format post end of release cycle on 8th April, 2021
    * [Rey, Docs]
        * Status: Green
        * **Doc Placeholder PR deadline is today**
        * [Integration branch](https://github.com/kubernetes/website/pull/26153) is healthy
        * All enhancements have replied if they need docs or not
        * Doc PR stats: we are currently tracking 43 docs out of 47
            * Docs merged: 8
            * Ready for review: 17
            * In draft: 8
            * No docs required: 10
            * Docs needed with no PR: 3 (owners said they will create PRs today)
            * Unknown: 0
* Issues & PRs
* Third-party & dual-sourced content
* Discussion
    * [irvi] kubectl documentation (currently in 2 places):
        * [Documentation on the kubernetes.io](https://kubernetes.io/docs/reference/kubectl/) and [kubectl handbook](https://kubectl.docs.kubernetes.io/), we can bring up this to get common grounds so we can merge these two.
        * I already sent a discussion in the Slack and Google Groups
            * [https://groups.google.com/g/kubernetes-sig-cli/c/G9wU_xHy7EA](https://groups.google.com/g/kubernetes-sig-cli/c/G9wU_xHy7EA)
            * [https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700](https://kubernetes.slack.com/archives/C2GL57FJ4/p1615910382001700)
    * [Anna] Docs deadline discussion
        * [https://github.com/kubernetes/sig-release/issues/1496](https://github.com/kubernetes/sig-release/issues/1496)

**03/09/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @onlydole
* Next week’s PR wrangler: @CelesteHorgan
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * Current Status: Yellow
    * Code Freeze is today
    * About ~20ish enhancements are at-risk and may be removed from the release at EOD or tomorrow
    * Last week’s _dev-1.21_ branch sync merged in [PR 26915](https://github.com/kubernetes/website/pull/26915)
    * [Integration branch](https://github.com/kubernetes/website/pull/26153) is healthy
    * Docs PR stats: we are currently tracking 40 docs out of 65
        * Docs merged: 7
        * Ready for review: 6
        * In draft: 9
        * No docs required: 5
        * Docs needed with no PR: 26
        * Unknown: 20
* Issues & PRs
    * [Karmvil] Credentials for private repository issue will be removed
    * [Tim] Migration to new API reference [https://github.com/kubernetes/website/issues/25505](https://github.com/kubernetes/website/issues/25505)
        * Discussion this Thursday with SIG Architecture
    * [Irvi] [https://github.com/kubernetes/community/pull/5569](https://github.com/kubernetes/community/pull/5569)
* Third-party & dual-sourced content
* Discussion
    * [Abbie] Are we planning anything for Kubecon EU?
        * Do we have a session in the Kubecon, may be a Docs Sprint
        * Will post in Slack
        * [Chris] Other idea for the Docs Sprint in Kubecon EU
        * [Tim] The folks who know better about the gaps are actually the newest contributors.
    * [Karmvil] Link from THIS agenda document to the archive of meeting videos
        * Done!

**03/02/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @kbhawkey
* Next week’s PR wrangler: @onlydole
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * Current Status: Yellow
    * Last week's dev-1.21 branch sync merged in k/website PR 26741
    * Integration branch is healthy
    * **Code Freeze is on March 9**, Docs team will start to communicate with the majority of enhancement owners on March 10
    * We will start to communicate with several enhancement owners this week. The Enhancements team have finished comms with these enhancement owners and are marked accordingly on their tracking sheet.
    * Docs PR stats: we are currently tracking 27 docs out of 65
        * Docs merged: 4
        * Ready for review: 3
        * In draft: 6
        * No docs required: 3
        * Docs needed with no PR: 23
        * Unknown: 35
    * Action: validate / verify proper branches are targeted (dev-1.2x vs. main/master)
* Issues & PRs
    * [https://github.com/kubernetes/website/pull/26786](https://github.com/kubernetes/website/pull/26786) PR about kubeadm doc to dev-1.21 branch but not related to a tracked enhancement. [Action-item] Ping SIG Cluster Lifecycle for a review
* Third-party & dual-sourced content
    * [Jim] Status check
* Discussion
    * [Tim] Revised API Reference, Tim reached out to SIG Architecture [mailing list](https://groups.google.com/g/kubernetes-sig-architecture/c/a_WQHt22iYo) and per @dims advice we put the concern on SIG Architecture [meeting agenda.](https://docs.google.com/document/d/1HtTI0rJEGP_MSf6eO87aCmx_tzpovPAAg7U2Zxwm8FE/edit?ts=5c9d0a4c#heading=h.8phgvuqrxjsg)
    * [Abigail] Taking a look into the analytics dashboard
        * Finding a way of generating reports and making it available for the folks who would like to know. Specific for the localization sub-group it can be in the form of setting up separate email so folks interested in getting analytics can subscribe. Create automated reports that get set to that group. Need feedback on what the community wants for these reports.
        * Review analytics to see if there were any site or docs improvements we should make.
        * [Irvi] Localization groups can try out reports if they’re available
        * [Jim] Korean localization used Google analytics to prioritize work
        * Historically SIG Docs sent out a top-10 list to the contributor group each week(?), this was seen as “noisy”. Clearly some folks do value this information.
        * [Jim] Japanese localization team (Kohei/@inductor) suggested an update once every 6 months.
        * [Jim] We investigated a public, anonymous means to have readonly access to the GA dashboard - this doesn’t seem easy.
        * [Jim] Suggests a SIG Docs reporting group that interested people can join. Challenge is that this is a dashboard style report, we can’t break out each localization (there are too many). Maybe a PDF per localization?
        * Mailing List info: [https://www.kubernetes.dev/docs/comms/mailing-lists/](https://www.kubernetes.dev/docs/comms/mailing-lists/)
        * [Karmavil] Maybe share via Slack
    * [Irvi] kubectl documentation (currently in 2 places)
        * [Tim] [Documentation on the kubernetes.io](https://kubernetes.io/docs/reference/kubectl/) and [kubectl handbook](https://kubectl.docs.kubernetes.io/), we can bring up this to get common grounds so we can merge these two.
        * [Tim] Suggestion: ping someone on the SIG CLI Slack channel and write a formal discussion on the google groups and possibly join into their meeting.
    * [Rael] Spanish localization weekly meeting kicking off
    * [Brad] Localizations
        * groups planning to share / document tooling
        * Approvals for localization .toml files are a bit frustrating (approvers configuration is managed per directory)

**02/23/2021**

**17:00 Pacific (01:00 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @kbarnard10
* Next week’s PR wrangler: @kbhawkey
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] Release Docs
    * Status: **Yellow** (many enhancements have not indicated if they will need Docs)
    * Docs team will start to communicate with enhancements owners on **March 10** (after code freeze)
    * Enhancements team will communicate with enhancement owners about Code Freeze on March 9 and include a question if they need Docs then make a Placeholder PR and also the Doc Placeholder PR Deadline of March 16
    * Last week’s _dev-1.21_ branch sync merged in [PR 26630](https://github.com/kubernetes/website/pull/26630)
    * [Integration Branch](https://github.com/kubernetes/website/pull/26153) is healthy
    * Docs PR stats: we are currently tracking **22** docs out of 67
        * Merged: 3
        * Ready for review: 2
        * In draft: 2
        * No docs required: 2
        * Docs needed with no PR: 24
        * Unknown: **34**
* Issues & PRs
    * [Rey] [PR 26629](https://github.com/kubernetes/website/pull/26629) Adding a deprecation to API Deprecation Guide on master before release date
        * [Rey] Bring up with sig release
    * [Rey] [PR 26109](https://github.com/kubernetes/website/pull/26109), targeted dev-1.21 branch, CLA has not been signed.
    * [Qiming] Config API generator
        * Issues: [#23889](https://github.com/kubernetes/website/pull/23889)
        * Reference-docs: [#176](https://github.com/kubernetes-sigs/reference-docs/pull/176)
        * Website PRs: [#26590](https://github.com/kubernetes/website/pull/26590),[ #26595](https://github.com/kubernetes/website/pull/26595), [#26596](https://github.com/kubernetes/website/pull/26596), [#26602](https://github.com/kubernetes/website/pull/26602), [#26605](https://github.com/kubernetes/website/pull/26605), [#26606](https://github.com/kubernetes/website/pull/26606), [#26608](https://github.com/kubernetes/website/pull/26608)
        * [Preview example](https://deploy-preview-26606--kubernetes-io-master-staging.netlify.app/docs/reference/config-api/apiserver-audit.v1/)
    * [Qiming] Tune flag type for component reference: [https://github.com/kubernetes-sigs/reference-docs/pull/213](https://github.com/kubernetes-sigs/reference-docs/pull/213)
    * AI: Qiming contact sig multicluster lifecycle for tech lgtm and PR release process before merging PR.
* Discussion
    * [Tim] **Reminder **Q4 review / Q1/2 planning **2/25** outcomes
        * [Agenda Doc](https://docs.google.com/document/d/1JIvPBRBFw47sxv-ijVYHc4852R30dDHdmKqUmKJxcas/edit?usp=sharing)
    * [Seokho] Korean Localization Update
        * Working on Korean localization (current: 1.20-ko.5, base: master, status: GREEN)
        * Korean l10n team would like to use `/milestone` prow command
            * We are using milestone to announce current Ko l10n schedule
                * Issue: [Korean Localization for dev-1.20-ko.5 branch · Issue #26621 · kubernetes/website](https://github.com/kubernetes/website/issues/26621)
                * milestone: [dev-1.20-ko.5 Milestone](https://github.com/kubernetes/website/milestone/60)
            * Website-milestone-maintainers team: [https://github.com/orgs/kubernetes/teams?query=website-milestone-maintainers](https://github.com/orgs/kubernetes/teams?query=website-milestone-maintainers)
            * For all Korean contents reviewers
        * AI: Jim to add seokho son to group
            * Related PR: [https://github.com/kubernetes/org/pull/644/files](https://github.com/kubernetes/org/pull/644/files)
    * [Seokho] Q: Commit squashing of a PR needs to be enforced?
        * [https://kubernetes.io/docs/contribute/new-content/open-a-pr/#squashing-commits](https://kubernetes.io/docs/contribute/new-content/open-a-pr/#squashing-commits)
            * “If your PR has multiple commits, you must squash them into a single commit before merging your PR.”
            * Requesting it to contributor is best practice (+ some exceptional cases)
            * A method to squash from approvers side
                * using tide label: /label tide/merge-method-squash
                * Info: https://kubernetes.io/docs/contribute/participate/pr-wranglers/#helpful-prow-commands-for-wranglers
    * [Jailton] Portuguese Localization
        * 2021 Kickoff Meeting: [https://www.youtube.com/watch?v=gMQ-jYY9zpc](https://www.youtube.com/watch?v=gMQ-jYY9zpc)
        * [https://github.com/kubernetes/community/tree/master/sig-docs#meetings](https://github.com/kubernetes/community/tree/master/sig-docs#meetings)

**02/16/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

Updates/reminders

* This week’s PR wrangler: @jimangel
* Next week’s PR wrangler: @kbarnard10
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] Release Docs
    * Status: **Yellow** (many enhancements have not indicated if they will need Docs)
    * Docs team will start to communicate with enhancements owners on March 10 (after code freeze)
    * Enhancements team will communicate with enhancement owners about Code Freeze on March 9 and include a question if they need Docs then make a Placeholder PR and also the Doc Placeholder PR Deadline of March 16
    * Last week’s _dev-1.21_ branch sync merged in [PR 26404](https://github.com/kubernetes/website/pull/26404)
    * [Integration Branch](https://github.com/kubernetes/website/pull/26153) is healthy
    * Last week, we started to track Enhancements and Doc PRs
    * This cycle is changing how we communicate with enhancement owners
    * Docs PR stats: we are currently tracking **17** docs out of 67
        * Merged: 3
        * Ready for review: 0
        * In draft: 1
        * No docs required: 2
        * Docs needed with no PR: 11
        * Unknown: **50**
    * [Divya] Release Comms 1.21
        * Notification regarding feature blog opt-in deadline to be sent out by RT Lead tomorrow in update email.
        * PR for formalizing the Comms deadlines + process for future releases - [https://github.com/kubernetes/sig-release/pull/1458](https://github.com/kubernetes/sig-release/pull/1458)
* Blog Update
* Issues & PRs
    * [Tim] [https://github.com/kubernetes/website/issues/25505](https://github.com/kubernetes/website/issues/25505) **Updated API reference generator**
* Discussion
    * [Jim] **Reminder **Q4 review / Q1/2 planning **2/18 (Thursday) 5pm Pacific**
        * [Agenda Doc](https://docs.google.com/document/d/1JIvPBRBFw47sxv-ijVYHc4852R30dDHdmKqUmKJxcas/edit?usp=sharing)
        * [Zoom](https://zoom.us/j/97784723884?pwd=OUJWdzZHb2F6THo2T1c1VjRzMFgvUT09)

**02/9/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Mike Tougeron - looking to learn more
* Jailton Lopes - Brazil looking to help with localization

Updates/reminders

* This week’s PR wrangler: @irvifa
* Next week’s PR wrangler: @jimangel
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Divya] SIGs will be opting in for feature blog delivery, starting this cycle. Official email with more details to be sent out by Release lead, later this week. Reference link: [https://kubernetes.slack.com/archives/C2C40FMNF/p1612420578158900?thread_ts=1612420578.158900&cid=C2C40FMNF](https://kubernetes.slack.com/archives/C2C40FMNF/p1612420578158900?thread_ts=1612420578.158900&cid=C2C40FMNF)
    * [Divya] TL;DR version for dates:
        * SIGs to opt-in for writing feature blogs by March 1st
            * Gave access to a google form to fill out to list if feature blog is desired (per enhancement)
        * Release Comms team to submit Release blog for review on March 25th.
        * Feature blogs to be submitted for review by March 31st.
        * Staggered publishing of feature blogs after Release on 8th April, following M-W-F cadence.
    * [Divya] PSP Deprecation blog - [https://github.com/kubernetes/website/pull/26279#issuecomment-774209178](https://github.com/kubernetes/website/pull/26279#issuecomment-774209178) (Also, Tim & Savitha have discussed this in the sig-security-docs call & are in talks about the idea for having a follow-up blog to above PR)
        * Follow up blog w/ volunteers
    * [Rey] Green
        * Healthy / merged
        * Today is enhancement freeze
        * 1 doc PR is already merged
* Blog Update
    * n/a
* Issues & PRs
* Discussion
    * [Abbie] (continued from last week) Tracking stale content in website repo - [https://docs.google.com/document/d/1H5I5kVovfXWd3jpXm-mmy7L2D9uIweo77kZs4-NazQs/edit](https://docs.google.com/document/d/1H5I5kVovfXWd3jpXm-mmy7L2D9uIweo77kZs4-NazQs/edit)
        * Goal to get ideas for requirements MVP and POC
            * Communicate project and get buy in from SIG Release and awareness from SIGs about this project. Get buy in from 1-2 SIGs to participate in an POC
            * Script that checks repo and returns stale content based on last commit date &lt; 1 year old
            * Running (small scale) process in 1.22 release cycle
        * Taylor working on [script](https://github.com/notchairmk/kubernetes-website/blob/git_history/scripts/git-file-history.sh) / draft PR
            * Probably need to scope down criteria
            * [Abbie] Issues could be a criteria too
            * [Tim] Reviewers in front matter (last reviewed that file)
                * Use analytics as a priority
            * [Tim] For front matter, initially validate format but don’t require it
            * [Jim] Focus on opening an issue to track and discuss this work
            * [Abbie] Want to get things more formalized into issue
            * [Abbie] How does web analytics work? [Jim] can grant access, happy to.
    * [Jim] Q4 review / Q1/2 planning **2/18 (Thursday) 5pm Pacific**
        * [Agenda Doc](https://docs.google.com/document/d/1JIvPBRBFw47sxv-ijVYHc4852R30dDHdmKqUmKJxcas/edit?usp=sharing)
        * [Zoom](https://zoom.us/j/97784723884?pwd=OUJWdzZHb2F6THo2T1c1VjRzMFgvUT09)

**02/2/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Karmavil - About me: passionate about programming in general (frontend, backend, sysadmin) but have worked most of my life in HVAC. I don’t speak english very well I do my best and I joined to collaborate in translation (spanish) and some typo maybe. Glad to be here.

Updates/reminders

* This week’s PR wrangler: @CelesteHorgan
* Next week’s PR wrangler: @irvifa
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * Status: Green
    * Last week's dev-1.21 branch sync merged PR 26292
    * Enhancements Freeze is next week on Feb 9; 32 enhancements currently
    * Docs team members are being assigned to enhancements this week to start tracking
    * [Divya] wanting to improve comms process, discussion potentially next week. Would also like to understand the process for reviewing of deprecation-related feature blogs prior to the release, if need be.
        * Setting deadlines around feature blog delivery / merges (on release day)
        * Major deprecation in cycle, needs a better communication process (contribex working on a blog for that)
            * PSP is so important and should be release blocking if can’t get communications out
        * Problem: Outline alternatives after posting deprecation (doesn’t exist as of today)
        * Rolfe: Concerns about a backlash? (yes)
            * Never too soon to send out bad news
            * Stage it, preview and later provide more context (with a followup date)
        * Tim: Comms around deprecation announcement
            * Three dates (when we make announcement, post release date, and when do folks worry that we don’t have the information in place)
            * When is the last date deadline? Right now?
        * Who saw the docker shim fun? What did we get wrong?
            * Rey - comms should come from us (community) vs Twitter / social
            * Rolfe - 1 word: deprecation, if we said “replacing” with docker shim etc. people would care less
            * Irvi - word choice
            * Brad - Review communications etc
            * Rolfe: empathy with audience is important
            * Tim: Including a call to SIG Security would reassure people +1
        * Divya: Post draft PR - [https://github.com/kubernetes/website/pull/26279](https://github.com/kubernetes/website/pull/26279)
* Blog Update
    * n/a
* Issues & PRs
    * n/a
* Discussion
    * [Abbie] Tracking stale content in website repo - [https://docs.google.com/document/d/1H5I5kVovfXWd3jpXm-mmy7L2D9uIweo77kZs4-NazQs/edit](https://docs.google.com/document/d/1H5I5kVovfXWd3jpXm-mmy7L2D9uIweo77kZs4-NazQs/edit)
        * Current plan to go off last commit (good faith that drive-by commits are reviewing the doc)
        * Rey: Release cycle is good (first few weeks, before is good). Add tasks to track stale content
        * Tim: we do have drive-by commits without full reviews
        * Test small section first / MVP to get traction
        * Rolfe: pilot storage
        * Tim: MVP, could identify things that need updating, not sure if removal is the way forward.
            * Abigail: No full removal, evaluate where it’s at & analytics (live disruption). Potentially a long term plan to create a policy to deprecate
        * Tim: we can automate (eg shortcode) comparing now to last reviewed date, and add warning through hugo - (brad: open issue instead).
        * Brad: be careful with messaging, automatic warnings etc. Potential to upset SIGs.
        * Rolfe: script to add notice to page (open issue, then resolve page)
        * Brad: Planning Q4 (top 5) become liaison (bring up)
        * Tim: Google Analytics + topics last updated = right sigs to reach out to first.
    * [Jim] Q4 review / Q1/2 planning: **2/18 (Thursday) 5pm Pacific**
    * [Brad] Localization subgroup request. Add a message to the PR instructions message that if the PR is for one of the localizations to add the localization country code as a prefix in the PR title. For example “[ko], [ja], [zh]”
    * [Jim] - New contributors open discussion

**01/26/2021**

**5pm Pacific (01:00 UTC)**

🎉 New contributors 🎉

* n/a

Updates/reminders

* This week’s PR wrangler: @bradtopol
* Next week’s PR wrangler: @CelesteHorgan
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] GREEN dev-1.21 create (to be synced later this week)
        * Major notes (new api generation) to be promoted
        * .toml config file needed (PR open)
    * [Divya] wanting to improve comms process, discussion potentially next week.
        * Theme: teams that are end of cycle heavy / alleviate
* Korean Localization Update
    * Branching strategy (with milestones every 2/3w)
    * Working on localization branch 1.20 (current: 1.20-ko.3, next: 1.20-ko.4)
    * GREEN - good to go!
    * New reviewer for Korean content (PyungHo Yoon. @yoonian)
* Blog Update
    * n/a
* Issues & PRs
    * [Jim] Fixing calendars
* Discussion
    * [Jim] Creating localization email thread
        * Filtering unique email (sounds better)
        * Statistics are needed ½ year or before presentation
        * Maybe unique email
    * [Seokho] Adding translator’s name in Blog `contents`.
        * Korean localization for blog contents will start including blog translators.
        * Discussed in Korean l10n team meeting.
        * Useful for promoting l10n contributions for Blog contents.
        * Share with other localizations too
    * [Jim] Q4 review / Q1/2 planning (no discussion, just reminder)
    * [qiming] API browser demo
        * Resources / groups versions and can list verbs
            * Details of definition
            * Object metadata
            * Closer to json file
        * Button to toggle entire template resource
        * Button to toggle version of API
        * Button to compare versions
        * Button to view history
            * Added (first)
            * Changed versions
        * Button to edit in browser
        * (?) help buttons how they are mapped
        * Filter definitions based on search bar
        * Mobile responsive
        * Focus on structures vs. ref, working on open sourcing. Where to host

**01/19/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Tony Gosselin - SRE at Adobe (on release team)
* Mike Tougeron - SRE at Adobe (getting involved in upstream)
* Chandani - Mathworks (looking to start contributing and a shadow for 1.21)
* Nivedita - (looking to start contributing)
* Sanchit Balchandani - (looking to start contributing)
* Victor - (1.21 shadow)

Updates/reminders

* This week’s PR wrangler: @onlydole
* Next week’s PR wrangler: @CelesteHorgan
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * [Rey] Week 2
        * Onboarding shaddows
        * Creating integration branch later today including config.toml
        * Light week
* Blog Update
    * n/a
* Issues & PRs
    * [Tim] API reference adoption & migration
        * [https://github.com/kubernetes/website/pull/23294](https://github.com/kubernetes/website/pull/23294)
        * [Zach] Context: google season of docs project to update API ref docs
            * “Guts” have merged
            * PR (above) is generated results of the generation mechanism (in a SIG repo)
            * Currently discussing what’s needed to merge successfully in the PR (above).
        * [Karen] It’s been thoroughly reviewed by a small subset
            * Open questions
            * Current proposal is to merge and have ref side-by-side until comfortable sunsetting
        * [Zach] Karen's feedback on breaking bookmarks, but should not be a blocker. (agree!)
            * Can we frame “why” we are running side-by-side
                * “Testing in production”
                * “Alpha / beta” etc…
        * [Zach] Mark old API reference as deprecated
        * [Mike T.] Impact on Google / SEO of old links
            * [Zach] will have the least SEO impact if we cut at the actual release time.
        * [Tim] Challenge is cut over and deprecation (not merging)
            * Merge the new Reference first and then mark it as deprecated later
            * Reaching out to SIG Architecture because they’re mostly aware about the deprecation warning
            * Communications plan
                * Arch (asking who else needs to know)
                * Mailing Lists
                * Blogs(?)
                * API machinery
* Discussion
    * [Jim] Q4 review / Q1/2 planning
        * Do we want to bundle Q4/1 together and then do the review in the next few weeks
        * AI: schedule meeting for Q4/Q1 combo review and plan for Q2.
    * [Tim] Using GitHub Discussions more
        * Can we use GitHub Discussions to discuss key decisions and record how we decided / should we?
            * [Rey] SIG Release already using this, +1
            * [Rey] Discussion per release
            * [Chris M] Do we run the risk of discussion tool overload (eg also Slack, Discourse, etc)
            * [Jim] If it makes folks’ lives easier I’m all for it.
            * [Jim] Use this when it has a very clear purpose.
            * [Brad] Have plenty of things to check already/
        * Can we somehow turned existing key issues to discussion
    * New contributors
        * Sample pipeline or place to practice PRs.
            * Open a test PR
            * Look at a [good first issues](https://github.com/kubernetes/website/labels/good%20first%20issue)
            * Look at [merged PRs](https://github.com/kubernetes/website/pulls?q=is%3Apr+is%3Amerged+label%3Alanguage%2Fen)
        * What does docs do around blogs and what is the process?
            * k8s.io/blog
            * Published much like SIG Docs (website) contributions
                * However, reviewed by a blog sub project team vs. the standard docs team
            * Kubernetes [blog team](https://app.slack.com/client/T09NY5SBT/CJDHVD54J/thread/C1J0BPD2M-1610737963.018200) on Slack
            * Content requirements differ a little
                * Tech outside of Kubernetes is out of scope
            * Help for writers (and folks not familiar with markdown)
        * [Tim] Join slack, join the sig docs channel, say Hi!

**01/12/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Rolfe DH

Updates/reminders

* This week’s PR wrangler: @sftim
* Next week’s PR wrangler: @annajung
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
* Blog Update
    * If you know of anyone that would like to help with reviewing blogs, please refer them or point them our way!
* Issues & PRs
    * [Tim]: [https://github.com/kubernetes/website/issues/25759](https://github.com/kubernetes/website/issues/25759)
        * [https://github.com/kubernetes/website/pull/25769](https://github.com/kubernetes/website/pull/25769)
        * “Separate the config driving the display and content of the announcement banner on the website to its own dir, with steering approvers”
        * **~~TODO: bring up weekly for review~~**
        * Merged & working!
    * [Tim] API reference work and attention
        * [https://github.com/kubernetes/website/pull/23294](https://github.com/kubernetes/website/pull/23294)
    * [Tim] [https://github.com/kubernetes/website/pull/24648](https://github.com/kubernetes/website/pull/24648) (short link for kubectl cheat sheet)
* Discussion
    * [Taylor] Google Analytics - transparently sharing metrics for the site
        * Research any potential security issues with “view” access to Google Analytics
        * Add a view to popular views/metrics without the need for GSuite
    * [Jim] Q4 review / Q1/2 planning

**01/05/2021**

**10:30am Pacific (18:30 UTC)**

🎉 New contributors 🎉

* Taylor Chaparro - looking to get started
* Clarke Vennerbeck - same

Updates/reminders

* This week’s PR wrangler: @bradtopol
* Next week’s PR wrangler: @sftim
    * Approvers, make sure you know your scheduled [PR wrangler shifts](https://github.com/kubernetes/website/wiki/PR-Wranglers)

Agenda:

* Release 1.21
    * Release cycle starts _next_ week (officially)
    * Improvement around tracking KEP reviewers for docs
* Blog Update
    * Things are great!
* Issues & PRs
    * [Celeste]: [https://github.com/kubernetes/examples/pull/393](https://github.com/kubernetes/examples/pull/393), [https://github.com/kubernetes/website/pull/22934](https://github.com/kubernetes/website/pull/22934)
        * Issue 1: WG Naming is okay with this in principle from a language perspective, but this replacement changes MongoDB to Redis. Is SIG Docs ok with this?
          _No objection to the spirit of these PRs._
        * Issue 2: Paul may not have time to push these over the line – can we take over the PRs?
          _Volunteer required_
    * [Tim]: [https://github.com/kubernetes/website/issues/25759](https://github.com/kubernetes/website/issues/25759)
        * [https://github.com/kubernetes/website/pull/25769](https://github.com/kubernetes/website/pull/25769)
        * “Separate the config driving the display and content of the announcement banner on the website to its own dir, with steering approvers”
            * Banners driven by date (metadata) but needs netlify to rebuild twice a day (just implemented this morning)
        * “Add the ability to specify a start and end date for an announcement, and adjust the display logic to only render the announcement between those dates, so no PR is required to enable or disable the announcement”
        * **TODO:** bring up weekly for review
    * Tim shout out to api-ref work and attention
        * [https://github.com/kubernetes/website/pull/23294](https://github.com/kubernetes/website/pull/23294)
* Discussion
    * [PR wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers)? [Tim] - thanks Brad!
