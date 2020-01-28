**12/18/2018**

**10:30am PST**

New contributors

*   Suresh Kumar Palemoni

Updates and Reminders

*   Kubecon last week!

Agenda

*   [Misty] PR Wrangling, Configure stale issues bot
    *   Bot’s current configuration needs to be improved.
    *   Why do we stale-out issues after 90 days? Organization’s defaults
    *   In GH: A PR is a type of issue, stale-bot shouldn’t close these
    *   AI: We should create a doc that guides issue triage, response time, priorities
*   [Jim] 1.14 release meister / shadow
    *   Steven Augustus sent out last call for shadows: Cody Clark, Jared Bhatti are current shadows
    *   Shadows: Working to write up a guide for future release meisters
*   [Jennifer] Covering for Zach during vacation
    *   AI: Jennifer: Will cover (loosely) over vacation 
*   [Jared] Holiday meetings?
    *   Next two weeks meetings are cancelled. 
    *   AI: Jared: Will announce on SIG-Docs and Slack
*   [paris] kubernetes.io/community redesign 
    *   [https://cjyabraham.gitlab.io/community](https://cjyabraham.gitlab.io/community)
    *   Will circulate in the contribex meeting tomorrow, too. 
    *   Final comments?
        *   Goals of the site: surface our resources/communication channels and the why, modern design, mobile friendly, general community voice vs one/two personas. 
    *   Need to clean up final copy, populate events, and we are ready to roll. Need to figure out a launch date that works. 
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
    *   AI: Jared: Will announce on SIG-Docs and Slack that the next two meetings are cancelled
    *   AI: Jennifer: Will cover PR Wrangling (loosely) over vacation 
    *   AI: Jennifer will get an agenda up with the Q1 goals for our first meeting of the year. 

**12/04/2018**

**10:30am Pacific**

New contributors



*   None today

Updates and Reminders



*   PR wrangler this week: @bradtopol
*   KubeCon Seattle: SIG Docs meeting is part of contributor day on Monday, 10 December.

Agenda



*   Congratulations to @tfogo and @jimangel on the 1.13 release! 
    *   Big thank you to @jrondeau for helping thoroughly and often, especially during crunch time
    *   [Zach] Need to add netlify site for 1.13, clean up old sites (1.8, 1.7)
*   [Zach] For the good of the order
    *   PR wranglers: 
        *   For week of 12/10, @mistyhacks is listed. Is she still available?
        *   Expectations for approvers to keep their status
    *   Meeting next week/KubeCon?
        *   NO regularly scheduled meeting
    *   Meetings through end of the year?
        *   @jrondeau will lead 12/18
        *   @jaredb will take first meeting 1/8/19
    *   Zach on vacation through mid-January: need someone to run/record
*   [Zach] Still working on planning session agenda, will finish by 5pm on Thursday, December 6th. 
    *   Placeholder issue: [https://github.com/kubernetes/website/issues/11331](https://github.com/kubernetes/website/issues/11331)
        *   Zach will find out the specific locations 
        *   We will meet from 10-2
        *   Please add our meeting to the [community meeting calendar](http://bit.ly/kubernetes-summit)
    *   [Jared] Can we please make this a google doc for collaborative editing and note-taking during the meeting?  
        *   Yes! Hoorays. 
*   [Jennifer, Jared, others?] Phil Wittrock’s [proposal for kubectl docs](https://groups.google.com/forum/#!topic/kubernetes-sig-docs/TXiYNIAB5hA) [[demo site](https://pwittrock-kubectl.firebaseapp.com/)]
    *   Can we (iteratively) incorporate into k8s.io?
            *   Current content: [https://kubernetes.io/docs/reference/kubectl/kubectl/](https://kubernetes.io/docs/reference/kubectl/kubectl/)
            *   Interlinking between K8s.io and this site
            *   Define contribution process
    *   What’s the contribution process?
        *   Would have to be defined. Owned by CLI group
    *   Is this a model we’d like to develop for other areas of the docs?
        *   [pwittrock] How do we feel about comprehensive linear “book” structure for specific topics - Tooling, Workload APIs, Cluster Management, etc?
*   ~~[Jennifer, Cody] Need to add information about separate PRs for different language directories to docs contributor guide ~~Thanks to @zachorsarah!
*   
*   [Jim--DEFER until 12/18, review Jared’s additions] Continued discussion of limiting simple / typo PRs on X-old blog posts
    *   [https://github.com/kubernetes/website/pull/11456#pullrequestreview-180862330](https://github.com/kubernetes/website/pull/11456#pullrequestreview-180862330)
    *   [https://github.com/kubernetes/website/pull/11451#pullrequestreview-180889637](https://github.com/kubernetes/website/pull/11451#pullrequestreview-180889637)
    *   Similar issue in k/k [https://github.com/kubernetes/community/issues/2953](https://github.com/kubernetes/community/issues/2953) focused on people “gamifying” PRs (I don’t know if that really is what we’re seeing or people just trying to submit a simple first PR).
    *   OpenFaaS handles it by having users open issues and then later creating a bulk PR ( more ownership on us) [https://github.com/openfaas/faas/blob/master/CONTRIBUTING.md#ive-found-a-typo](https://github.com/openfaas/faas/blob/master/CONTRIBUTING.md#ive-found-a-typo)
    *   Outcome: Do we want to define / document a policy on minor PR’s that address old or unused docs (If yes, what)?
    *   AI: Jared, will take a look at
    *   Tangentially relevant/helpful: [“what to leave out”](http://www.writethedocs.org/guide/contributing/#what-to-leave-out) from the WTD guide to contributing to the WTD documentation guide 
*   [Andrew, Dominik] Kubernetes Modeling
    *   Project update at SIG Docs planning meeting during contributor summit
        *   What time?
    *   Latest blog post [The Mechanics of Kubernetes](https://medium.com/@dominik.tornow/the-mechanics-of-kubernetes-ac8112eaa302)
        *   Feedback welcome!
    *   Upcoming: The Scheduler
        *   blog post
        *   will propose unconference session @ contributor summit


## Past Meetings {#past-meetings}


### 2018 Meetings {#2018-meetings}

**11/27/2018**

**7pm Pacific**

New contributors



*   

Updates and Reminders



*   PR wrangler this week: @jrondeau (@bradamant3 on GH)
*   1.13 updates?
*   KubeCon Seattle is in two weeks: SIG Docs meeting is part of contributor day on Monday, 10 December. Anyone from APAC planning to attend?

Agenda



*   Netlify configuration for release 1.12 branch should have build-per-PR turned on
*   Chinese contribution/membership: How is it going?
    *   Conversion from Google Doc to GH issues is going great
    *   Netlify builds for 1.12 PRs are working
    *   Progress is great!
*   Korean contribution/membership: How is it going?
    *   Meetup last week: new members joining project
    *   Working on dev-1.13 branch, coordinating with @tfogo for final PR on December 3
    *   About 10 pages of updates for 1.13

**11/20/2018**

**10:30am Pacific**

New contributors



*   Dennis Salama - Microsoft - Azure services and support team

Updates and Reminders



*   1.13 updates?
    *   Update from Tim
    *   Big crunch yesterday, down to 3 PRs that need to be reviewed
        *   Review deadline: Nov 27
        *   Release: Dec 3
    *   All PRs in review except 3. Tracking [here](https://docs.google.com/spreadsheets/d/1umeZ-AHjjD6ntbFv1J2dJOjWd0Ft0WKAV2Wtd8EQ8FM/edit#gid=0).
        *   Kubeadm
*   PR wrangler this week is @tengqm, with thanks for swapping because of Thanksgiving in the U.S. Thank you!

Agenda



*   Kubectl docs generation -- inspired by [this PR](https://github.com/kubernetes/website/pull/10778). PR might could be fixed with a PR to upstream code to fix string literal. But another option is to change the generator that we use for the kubectl docs. What says the docs community?
    *   (late addition to notes) AI: Jennifer will find right resource to get PR to upstream code merged, and submit said PR
*   [chenopis] Simplifying the solution buckets on the [Setup landing page](https://kubernetes.io/docs/setup/).
    *   Managed Solutions, e.g. GKE, EKS, etc. -- simpler billing, not paying for master
    *   Hosted Solutions, e.g. K8s on GCE or EC2 -- more control of the # of machines, version of K8s, monitoring code, identity management
    *   Custom Solutions - bare metal, on-premises
    *   Local Machine - Minikube, etc.
    *   Notes from meeting: 
        *   Looks good to the SIG-Docs group
        *   Codyclark can send an email to chenopis
*   [jrondeau/Bradamant3] PRs against old blog posts. Inspired by [this PR](https://github.com/kubernetes/website/pull/11009), which fixes a 404 but points uselessly to an old release of a project no longer actively under development. Link checking is a good thing. How do we deal with the old blog posts part, though?
    *   Document in contributor guide
        *   Past one year don’t update :)
        *   AI: Jared will submit a PR
    *   Compassionate responses to PRs that come in from outside sig-docs
*   [jrondeau/Bradamant3] [Redirects to current page for earlier version](https://github.com/kubernetes/website/pull/10985) -- PR raises an excellent issue, but seems non-trivial to address properly. (Will remove from agenda before meeting if assignees come up with solution.)
*   [Jaredbhatti] Docfixit at Kubecon Seattle?
    *   Someone leading this? Specific efforts we’d like to do?
    *   Main Focus: Doc Summit instead of doc sprint
    *   AI: Jared and Jennifer will discuss in Slack, consider scheduling options

**11/13/2018**

**10:30am Pacific**

New contributors



*   Shavi Dissanayake -- new CNCF intern!

Updates and Reminders



*   1.13 updates? (no updates this week)
*   PR wrangler this week is @Bradamant3 (jrondeau on Slack)

Agenda



*   N/A -- Jennifer, Shavi, Steve only attendees

**11/6/2018**

**10:30am Pacific**

New contributors



*   Robin Rakowski 

Updates and Reminders



*   PR wrangler this week is @cody-clark
*   Reminder: test changes locally, even for trivial PRs
    *   Verify that Netlify builds succeed before merging
        *   Add guidance to contributor section

Agenda



*   Content updates:
    *   Fun with Modeling (@chenopis, @dtornow) ~15 min
        *   Who is our audience exactly?
        *   We believe these questions should be able to be answered by reading the documentation:
            *   [https://github.com/kubernetes/kubernetes/issues/62415](https://github.com/kubernetes/kubernetes/issues/62415)
            *   [https://www.reddit.com/r/kubernetes/comments/9u4b3a/is_there_a_way_to_specify_order_of_containers_to/](https://www.reddit.com/r/kubernetes/comments/9u4b3a/is_there_a_way_to_specify_order_of_containers_to/)
    *   Refactoring home page: [https://github.com/kubernetes/website/pull/10795](https://github.com/kubernetes/website/pull/10795) (@chenopis) [[proposal](https://docs.google.com/document/d/1PxuMaXm7hFjO3MuzWmog-IW62ZnrTQ1S3GLcaAsOiIY/edit?usp=sharing)]
        *   Content structure: google doc
        *   Wait until WIP is dropped from title to add comments to PR
    *   CNCF devoting more resources to initial onboarding
*   Tooling updates:
    *   Working group? (Luc)
        *   (Zach C) https://discourse.gohugo.io/t/ignorefiles-and-i18n-expected-behavior/15097
*   KubeCon/CloudNativeCon Seattle
    *   Quarterly planning
        *   Contributor day info (draft): [https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit](https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit)
        *   We’re currently scheduled for 2 hours but can have more. I’ve asked that SIG Docs have 4 hours. 
        *   It looks like we have 2 or 3 things going on in the same space. Let’s make sure we understand what we’re doing so we can explain to others what we’re doing :-) 
*   Intro to contributing to k/website video (@Steve, @Jen)

**10/30/2018**

**10:30am Pacific**

New contributors



*   Sam @ Heptio

Updates and Reminders



*   PR wrangler this week (and next!) is @cody-clark

Agenda



*   Content updates:
    *   Fun with Modeling (@chenopis, @dtornow) ~15 min
        *   Pod [[slides](https://docs.google.com/presentation/d/13vOH98rTGT4x-9cQznJLbsWtW472TIdawLxsSFI7R5s/edit?usp=sharing)]
        *   [What the hell is a Pod anyways?](https://medium.com/@dominik.tornow/what-the-hell-is-a-pod-anyways-72e5534b892c) blog post
    *   GH project cleanup? (Steve Perry -- update from APAC meeting last week)
*   KubeCon/CloudNativeCon Seattle
    *   December 10-13, who’s going?
        *   8-9
    *   Yearly retrospective
        *   SIG Docs awards
    *   Quarterly planning
        *   Zach: get details by 11/6

**10/23/2018**

**Fourth Tuesday**

**7pm Pacific**

New contributors



*   

Updates and Reminders



*   

Agenda



*   Fun with Modeling (@chenopis, @dtornow) ~15 min
    *   [High Availability](https://docs.google.com/presentation/d/1stVHEnEf1tr83gq4Am2MCTqhmJvF5XdXm3mJbGZzlNk/edit?usp=sharing)
*   Content updates: GH project cleanup? (Steve Perry -- defer to next week, Jennifer can also help coordinate, brainstorm) (Zach C: Defer to 10/30)
*   Logistics surrounding #sig-docs-tools WG (Luc)
    *   [Slack channel](https://kubernetes.slack.com/messages/CCTBSCVB9/details/) for docs tooling proposals/questions 
*   Internationalization workflows
    *   Language labels
    *   Write permissions for branches/labels: [https://github.com/kubernetes/website/issues/10682](https://github.com/kubernetes/website/issues/10682)

**10/16/2018**

**10:30am Pacific**

New contributors



*   

Updates and Reminders



*   Congrats and thanks to new reviewers/approvers:
*   New English reviewer: Jim Angel (GH: jimangel)
*   New English approver: Stewart Yu (GH: stewart-yu)
*   New Chinese reviewer: Yang Li (GH: idealhack)

Agenda



*   Fun with Modeling (@chenopis, @dtornow) ~15 min
    *   [High Availability](https://docs.google.com/presentation/d/1stVHEnEf1tr83gq4Am2MCTqhmJvF5XdXm3mJbGZzlNk/edit?usp=sharing)
*   Content updates: GH project cleanup? (Steve Perry -- defer to next week, Jennifer can also help coordinate, brainstorm)
*   Logistics surrounding #sig-docs-tools WG (Luc)
    *   Luc will consider labels for issues to help track
    *   Generated docs
    *   Need a list (everything that’s not content)
*   NCW SIG Doc Presentation for KubeCon Shanghai (@bradtopol -- deferred, Brad unable to attend today)

**10/9/2018**

**10:30am Pacific**

New contributors



*   

Updates and reminders



*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 10/8: Qiming Teng (@tengqm)
*   SIG Docs Leads: Jared Bhatti stepping down, Jennifer Rondeau stepping up

Agenda



*   Renewed focus on content (Zach)
    *   Let’s make sure our onboarding content rocks
    *   Focus on content first
    *   Defer tooling/automation to the #sig-docs-tools WG, get updates
    *   Need to improve project tracking for content work (see next item)
*   Longer-term project planning (paris)
    *   Some suggestions from Paris Pittman
    *   [Project board from Contrib-x](https://github.com/orgs/kubernetes/projects/1)
    *   Each project has an owner, each work-item has an issue
    *   Issues that aren’t owned are given a “help wanted” label
    *   New contributors are directed at “help wanted” issues
    *   Projects have standups during the meeting - makes the meeting run a bit faster.
    *   Sub-projects, or projects that need more time can use the remaining time in the meeting. 
    *   Zach: [Current progress in repo](https://github.com/kubernetes/website/projects)
    *   Steve: AI: I can take on cleaning up 
    *   Jennifer: Think about division between website projects and larger kubernetes (umbrella) projects
        *   Paris
            *   Those umbrella issues are key and labels for areas/foo 
            *   areas/foo being your subprojects
*   Quarterly/yearly goals (Paris Pittman, Zach C)
    *   Start setting goals/priorities on a quarterly basis ongoing
        *   Paris: plan projects better using the process above, which leads to better quarterly planning. It’s not a science, so we’re not ready for OKRs
    *   Zach and Paris will meet and discuss
    *   Meet at KubeCon Seattle and chat further
*   Fun with Modeling (@chenopis, @dtornow) ~15 min
    *   [The great lie](https://docs.google.com/presentation/d/1UuinyQSyuL0bsp4lYcfl4O-4GP5bSkdEdH4GRPflnuE/edit?usp=sharing)
*   Accessibility tools (@rajie)
    *   NOTE: Deferred to SIG Docs tooling working group
    *   Move to docs tools WG for consideration
    *   http://pa11y.org
    *   http://khan.github.io/tota11y/

**10/2/2018**

**10:30am Pacific**

New contributors



*   Sam - TW @ Hepito
*   Alasdair - IBM - Java Runtime
*   YK - IBM - Java Runtime
*   Michal - IBM - Java Runtime

Updates and reminders



*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 10/1: Zach Arnold (@zparnold)

Agenda



*   Update on 1.12 (Zach A)
    *   DONE! 
    *   Moar automation!
*   Accessibility tools (@rajie)
    *   DEFER for 1 more week and revisit 10/9
    *   http://pa11y.org
    *   http://khan.github.io/tota11y/
*   I18n repo branching strategy (Zach A)
    *   Move i18n repos back into k/website BUT need to:
        *   add labels for languages so folks can scope PRs (Zach A. demo)
        *   Need to add OWNERS files in various content/** folders (Zach C)
        *   Need to update localization guidelines (Zach A)
*   Quarterly/yearly goals (Zach C)
    *   Meet at KubeCon Seattle
    *   Start setting goals/priorities on a quarterly basis ongoing
    *   Jared: Check-in on this 10/9
*   (10/2, Ben Hall) SIG Docs and Katacoda tutorial content
*   (10/2, Michal Broz, Brad Topol) Desire to add a new interactive tutorial for deploying Java Applications to Kubernetes 
    *   Interactive guides are very valuable
    *   Currently use a Node.js container
    *   Want to add another flow with the Java runtime
    *   Where will it sit in the overall k8s architecture?
        *   Current flow is 6 pages telling a story
*   Is there an ongoing Spanish effort? else I could work on/coordinate that. How about a German one?[Silvia]
    *   There are currently no Spanish or German translations in progress, but you're welcome to start one! We're updating the localization guidelines to reflect an updated workflow--please feel free to follow along in the next two weeks!

**9/25/2018**

**Fourth Tuesday: 7pm Pacific**

New contributors



*   

Updates and reminders



*   Thanks to @jrondeau for wrangling PRs this week--down a net of 23 PRs!
*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 9/24: Zach Corleissen (@zacharysarah)

Agenda



*   Update on 1.12 (Zach A)
    *   Going well (rumor has it!)
*   Accessibility tools (@rajie)
    *   DEFER for 1 week and revisit 10/1
    *   http://pa11y.org
    *   http://khan.github.io/tota11y/
*   I18n repo branching strategy (Zach C)
    *   After more research, subtrees look unlikely to work
    *   WG is meeting on Thursday 9/27 to confirm and recommend an official strategy (passing commits between repos)
    *   (June Yi) Korean docs branching strategy
        *   Open PR to move OWNERS to content/ko/
    *   Jim Angel, Svetlana: please message me in Slack (@zacharysarah) with your email so I can invite you to the WG meeting
*   Quarterly/yearly goals
    *   Meet at KubeCon Seattle
    *   Start setting goals/priorities on a quarterly basis ongoing
*   (10/2, Ben Hall) SIG Docs and Katacoda tutorial content
*   (10/2, Michal Broz, Brad Topol) Desire to add a new interactive tutorial for deploying Java Applications to Kubernetes 

    [Silvi]


**9/18/2018**

New contributors



*   AGV - Google customer engineer
*   Dominik returned

Updates and reminders



*   Thanks to @bradtopol for wrangling PRs this week (and surviving!)
*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 9/17: Jennifer Rondeau (@jrondeau on Slack, @bradamant3 on GH)

Agenda



*   Update on 1.12 (Zach A)
    *   All docs are staged, Jen Rondeau is working on kubeadm docs, and I will hopefully have final merge done this week
    *   Had to back out CoreDNS
    *   Hit the tripwire that restarted the branching discussion with maintainers :)
*   Fun with Modeling (@chenopis, @dtornow) ~25 min
    *   What is a model? [[slides](https://docs.google.com/presentation/d/1gkT67nnS0STPMZR2rdA1Jj9fHjFp291PU3M2_e9dd_c/edit?usp=sharing)]
    *   Imperative vs Declarative [[slides](https://docs.google.com/presentation/d/1-FjTUxuJo3nvGS8jnFdsDMS9B8jniEuU5FgQ0rOFk0I/edit?usp=sharing)]
    *   [Request for Comment (RFC)](https://docs.google.com/document/d/1eolKdH_LaT1dfMn8zry9CJ5gbs6hCm7prctkgmH3HZQ/edit?usp=sharing)
*   Accessibility tools (@rajie)
    *   DEFER for 1 week and revisit
    *   http://pa11y.org
    *   http://khan.github.io/tota11y/
*   I18n repo branching strategy (Zach C)
    *   After some discussion, commits back and forth could work but subtrees have the potential to be awesome and easier (after some initial setup)
    *   Zach C is working on a prototype
    *   Add [AGV](mailto:agv@google.com) and Brad Topol to the WG meeting
    *   Add Jim Angel and Svetlana Karslioglu
*   Quarterly/yearly goals
    *   Meet at KubeCon Seattle
    *   Start setting goals/priorities on a quarterly basis ongoing

**9/11/2018**

New contributors



*   Welcome to Silvi TSM

Updates and reminders



*   Thanks @tengqm for wrangling last week
*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 9/10: Brad Topol (@bradtopol)
*   Thanks, Brad for stepping in so Zach A can get 1.12 docs out the door
    *   Update: Brad is facing down Hurricane Florence

Agenda



*   (Zach A) 1.12 update
    *   Ability to create 1.13 milestone? (Zach C can help create)
    *   Teams to mention for docs: 
        *   @kubernetes-website-admins 
        *   @sig-docs-pr-reviews 
    *   Where do SIG Docs tools live in k/website?
*   (Vlad) K8s chatbot: gives answers to user questions based on docs, stack overflow, Slack convos
    *   Working group: Q&A was yesterday
    *   (Zach) Should SIG Docs implement answerbot for #sig-docs? 
    *   (Jennifer) Talk about data from #kubernetes-novice and #kubernetes-users
    *   Let’s take a look at the raw data from Vlad/Foqal for Foqal use in existing channels
    *   Follow up with Jennifer in week after release for advice about how/whether to proceed (9/??)
        *   Steve: As we examine research, can we examine that data may be mixed/in cheatsheet form? Consider that URLs may have mixed info.
*   (Luc) Tools/infra WG update
    *   Link to: [https://docs.google.com/presentation/d/1W_M3jAirn5tl8kN2B8pd4OEradZIJONJf988aIQiLhE/edit?ts=5b968ea9#slide=id.p](https://docs.google.com/presentation/d/1W_M3jAirn5tl8kN2B8pd4OEradZIJONJf988aIQiLhE/edit?ts=5b968ea9#slide=id.p)
    *   Rajie: How can we improve site accessibility? Will initially present recommendations to the WG regarding tooling, i.e. how we can measure site accessibility and point out instances of non-compliance. Plan(s) of action to be discussed later as a group.
*   (Zach A./Zach C./Zachs/Zachii?) Branching strategy WG update
    *   We met on Monday, meeting again at 11:30am for a follow-up.
*   (Misty/Tim F) best practices for using `for_k8s_version`?
    *   DEFER until Misty/TFogo attend
    *   Zach C will contact Misty in advance of next meeting: devote time for sufficient discussion, as it’s a large issue

**9/4/2018**

New contributors



*   Vlad Shlosberg

Updates and reminders



*   [Weekly PR wrangler](https://github.com/kubernetes/website/wiki/PR-Wranglers) for 9/3: Qiming Teng (@tengqm) 

Agenda



*   (Vlad) K8s chatbot: gives answers to user questions based on docs, stack overflow, Slack convos
    *   Working group: Zach C will set up a time for Q&A with Vlad
*   (Zach A.) Code Freeze/Docs Deadline is today, so expect an exciting update from me. I’ve intentionally not merged any PR’s into the 1.12 branch to I could have time to ensure that everything was ship shape from a Git standpoint
    *   43 features need docs updates
    *   29 features have permission to pull from milestone
    *   [Tracking spreadsheet](https://docs.google.com/spreadsheets/d/177LIKnO3yUmE0ryIg9OBek54Y-abw8OE8pq-9QgnGM4/edit)
    *   Jennifer will help with kubeadm docs
    *   Generated ref docs: where are those at?
        *   Jim and/or Jim are fixing bugs in refdoc generation process
        *   Steve: Which doc tools are being revised?
        *   Jim: [Link to PR](https://github.com/kubernetes/website/pull/10123)
        *   Steve: role is changing, but happy to help with refdocs in release PRs
    *   PR reviewers:
        *   Make sure 1.12 releases base on release-1.12, not master
        *   Add 1.12 feature PRs to 1.12 Milestone
*   (Luc) Update on tooling and infra working group. Proposal document: [https://docs.google.com/document/d/12RTw9s7zGncKuaY6zXiTJ4f7cWf_7iAWlaWgoJJyQ3g/edit](https://docs.google.com/document/d/12RTw9s7zGncKuaY6zXiTJ4f7cWf_7iAWlaWgoJJyQ3g/edit)
*   (Zach A.) Proposal/Idea from the Korean Translation team on how to do some automated branch syncing for long running (i.e. release style) branches 
    *   (Zach C.) Branching strategy for l10n repos: working group to figure out addition to l10n guide?
    *   (Misty / Zach) Release management branching strategy [proposal](https://docs.google.com/document/d/159xLCtr_0SffB28qXPlWGVefn6ihMM3z9DAU_V-oJMQ/edit#)
    *   Further discussion will take place in the Tooling and Infra working group
*   Misty role change
    *   Tech lead for Google GKE docs, focusing less on K8s docs

**8/28/2018**

**Fourth Tuesday (7pm Pacific)**

New contributors



*   None :-(

Updates and reminders



*   Reminder: 8/28 is a Fourth Tuesday, so this meeting happens at 7pm Pacific. Next week’s meeting resumes at 10:30am Pacific.

Agenda



*   Follow up from last APAC meeting (7/24):
    *   Is the localization guideline document up to date?
        *   Yes: [#9192](https://github.com/kubernetes/website/pull/9192)
*   Proposal: tooling and infrastructure working group (Luc)
    *   Dedicate a smaller group to ongoing, more immediate work
    *   Andrew: good idea!
    *   Luc will communicate and create a Slack channel if need be
*   I18n repo status: OWNERS and OWNERS_ALIASES files
    *   kubernetes-docs-ja: updated
    *   kubernetes-docs-ko: updated
    *   kubernetes-docs-zh: need updates
        *   Follow prow workflow
        *   Update their OWNERS* files
    *   Korean team: Put June Yi in touch with Zach Arnold
    *   Set up Netlify for i18n teams for preview builds
*   Branching strategy for i18n repos
    *   What do teams need?
*   Follow up on [https://kubernetes.io/docs/setup/pick-right-solution/](https://kubernetes.io/docs/setup/pick-right-solution/) and left navi panel inconsistencies;  Are we still updating [https://kubernetes.io/docs/setup/pick-right-solution/](https://kubernetes.io/docs/setup/pick-right-solution/) to ensure it is up to date? Also see discussion in https://github.com/kubernetes/website/pull/10028
*   KubeCon @ Shanghai
    *   Who is planning to attend?
    *   What do people want to do for the doc sprint?

**8/21/2018**

New contributors



*   

Updates & reminders



*   Zach Arnold Joining Late

Agenda



*   [Andrew] WTD conversation about where we are, where to go
    *   Treading water, not really advancing
    *   What 1-2 things do we focus on through end of Q4?
    *   Zach & Luc will talk with Chris A at CNCF about how to approach partners (AWS, Azure) to fulfill writer hiring commitments, especially in the context of Getting Started refactors

**8/14/2018**

New contributors



*   

Updates & reminders



*   PR wrangler this week is Brad Topol
*   Zach is still semi-unavailable due to health reasons, DM him directly on Slack for response

Agenda



*   HOLD [Zach Arnold] Update on 1.12
    *   Will be joining late
*   [Jennifer/Andrew] Kubelet generated docs [k/k PR #66034](https://github.com/kubernetes/kubernetes/pull/66034) -- we need someone or a small working group focused on generated docs
    *   Specific to kubelet, but representative of issues specific to generated docs as a whole
    *   Needs to be addressed by k/k and sig-docs
    *   Andrew is tapped, Jennifer can’t work alone, need folks to step up
    *   Another option may be to contact SPF folks to turn off third-party flags
    *   Jim Angel is available! Yay!
*   [Jennifer reporting only] should this sig be taking the lead on containerd docs?
    *   Link to diagram for underlying runtime diagrams:  [https://docs.google.com/drawings/d/1vb31vSh3qwzRgjevjJayVZsnh7tRsAszx__EzXoddZA/edit](https://docs.google.com/drawings/d/1vb31vSh3qwzRgjevjJayVZsnh7tRsAszx__EzXoddZA/edit)
*   [Jennifer] add to style guide: avoid patronizing language, specifically the major “weasel” words “just”, “only”, “simply”, “simple”, “easy”
*   [Zach] We need someone to lead next week’s meeting (8/21), many folks will be at Write the Docs Cincinnati
    *   Use meeting time for SIG Docs planning
*   [Andrew] WtD @ Cincinnati -- what do people want to do there?
    *   Mon - coordinated PR bash vs doc sprint?
        *   Jennifer will run doc sprint, Andrew will run coordinated PR bash
    *   Tue - SIG Docs planning
        *   Andrew/Zach will help plan/retrospective like SIG Docs Summit
*   [Zach] What should we do with OWNERS_ALIASES files in i18n repos?
    *   I propose stripping them/providing empty files, but letting SIGs make subsequent additions for members with fluency
    *   Empty files, SIGs can add to empty files as preferred/needed
*   [Andrew] K8s.io Search Outage post-mortem
    *   Published in public GitHub [Issue #9758](https://github.com/kubernetes/website/issues/9758#issuecomment-413698701)
*   [Zach A.] Message Dan Kohn sent to Sig Docs
    *   CDN speed for kubernetes.cn [https://github.com/kubernetes/website/issues/9634](https://github.com/kubernetes/website/issues/9634) 
        *   Prefer to use a mirror in China (Alibaba) for assets that are coming from elsewhere
    *   Search engine selection for kubernetes.cn [https://github.com/kubernetes/website/pull/9845](https://github.com/kubernetes/website/pull/9845) 
        *   Want to adjust the PR to include a way to adjust the search engine by drop-down if possible
        *   Additionally use Google Analytics to back up decision

**8/7/2018**

New contributors



*   Svetlana Karslioglu (Rackspace)
*   Pooja Gadige

Updates & reminders



*   Zach C is semi-unavailable due to ongoing health stuff, message him directly on Slack for attention

Agenda



*   1.12 Update (Zach A) [[status spreadsheet](https://airtable.com/shrqQbRD8oSDQizgY)]
    *   Things are going well, just managing the branch.
    *   There are 7 PR’s at present in k/k that will need docs and we’re beginning to bother them
*   PR Wrangling update
    *   82 open PR’s, and things seem to be operating as normal, but will reach out if conditions worsen
    *   354 open issues, will be going through those as well
*   [Google Search problem](https://github.com/kubernetes/website/pull/9767) discussion (Misty added on behalf of Andrew)
*   What to do about docs that reference maintenance [releases by number that have to be bumped manually, such as 1.11.1](https://github.com/kubernetes/website/pull/9613/)? (Misty)
    *   Follow up from 7/31
    *   Fix with a contractor (Zach)
    *   Release meister: update point release variables? 
    *   Update release notes for point releases as well? (Nick Chase)
*   Tactics for importing content from localization repos into k/website
    *   Subtree: [https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another](https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another)
    *   Follow up from 7/31
    *   ...Or just open a giant PR for the first import, then set up webhooks for subsequent changes?
    *   Zach will start a WG for automation
*   [Fundamental Modeling Concepts (FMC) for Kubernetes](https://docs.google.com/presentation/d/1vUAkRP-MjNqusqDHBptycdSbC_HSTBKyEHAFZ5OdbQA/edit?usp=sharing) (Andrew, Dominik)
    *   Dominik will reach out to Zach to initiate a contract
*   DSathe: Talk about cluster operator docs next week (8/7) [[slides](https://docs.google.com/presentation/d/1ebkJBOE2r-UeEGe7Ua5TUAHadd4Xm9Wri6FacUY8G9U/edit#slide=id.g3e00e93648_0_25)]
*   [Docs Contributor guide](https://kubernetes.io/docs/contribute/) update (MIsty)
*   WtD @ Cincinnati -- PR bash, SIG Docs planning (Andrew)
    *   Propose doing PR bash on Tue, SIG Docs planning on Wed
    *   Planning would be a continuation of the discussions from the SIG Docs Summit @ last WtD Portland
*   kubernetes.io/community redesign (paris)
    *   Next steps: copy decisions and calendar mechanisms 
    *   Will update this group when we have more to show
    *   [Issue is still collecting suggestions and updates](https://github.com/kubernetes/website/issues/7795)
*   Refactoring _[Picking the Right Solution](https://kubernetes.io/docs/setup/pick-right-solution/)_ page (Cody)

**7/31/2018**

New contributors



*   Dhananjay Sathe - cloud robotics startup (+1)

Updates & reminders



*   Maintainers, check the wiki for the PR wrangler schedule through the end of the year: [https://github.com/kubernetes/website/wiki/PR-Wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers)
*   Maintainers, please check this list for accuracy/completeness: [https://github.com/kubernetes/website/pull/9510#issuecomment-409056524](https://github.com/kubernetes/website/pull/9510#issuecomment-409056524) 

Agenda



*   What to do about Black Friday (Luc)
    *   Talked to BEP--Black Friday 2 is on the way, but developmentally costly
    *   Porting Kramdown to Go may be more effort than it’s worth
    *   Follow up conversation about next steps with Luc/BEP
*   1.12 Update
    *   Zach A. Has a meeting conflict, but all is going well so far! 
    *   We’ve automated a chunk of our workflow: [https://airtable.com/shrqQbRD8oSDQizgY](https://airtable.com/shrqQbRD8oSDQizgY) and the release team is now going to be borrowing from our Airtable for task coordination. Go docs!
    *   Docs PR Deadline is 3 weeks away!
*   What to do about docs that reference maintenance [releases by number that have to be bumped manually, such as 1.11.1](https://github.com/kubernetes/website/pull/9613/)? (Misty)
*   Tactics for importing content from localization repos into k/website
    *   Subtree: [https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another](https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another)
    *   Other possibilities?
    *   Revisit this on 8/7
*   Improvements to site UX (Neha, Andrew)
    *   Link to presentation:
    *   How to contact Neha with questions:
*   [Fundamental Modeling Concepts (FMC) for Kubernetes](https://docs.google.com/presentation/d/1vUAkRP-MjNqusqDHBptycdSbC_HSTBKyEHAFZ5OdbQA/edit?usp=sharing) (Andrew, Dominik)
    *   Deliverables: what would they be? (SVGs in docs)
    *   Replace/augment: “What is a Pod” section in Docs
    *   Pilot: proposed to be Pod, but could it be Kubernetes as a whole?
        *   It could be; but
        *   Pods are fundamental concepts to Kubernetes
    *   Dominik will come back on 8/7 to talk more
*   DSathe: Talk about cluster operator docs next week (8/7)

**7/24/2018**

REMINDER: Fourth Tuesdays are APAC meeting times (7pm Pacific)

Agenda



*   
*   New localization repositories:
    *   Korean: [https://github.com/kubernetes/kubernetes-docs-ko](https://github.com/kubernetes/kubernetes-docs-ko)
    *   Japanese: [https://github.com/kubernetes/kubernetes-docs-ja](https://github.com/kubernetes/kubernetes-docs-ja) 
    *   NEXT STEPS:
        *   Get familiar with the Kubernetes bot workflow commands: [https://prow.k8s.io/command-help](https://prow.k8s.io/command-help)
            *   Specifically: /lgtm, /approve, and /hold
        *   Migrate the existing repositories into the Kubernetes organization
            *   Some guidance: 
                *   [https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another](https://stackoverflow.com/questions/1683531/how-to-import-existing-git-repository-into-another) 
        *   Open a PR to create your own OWNERS files
*   Kubernetes.cn is live!
    *   Some issues: Dan will contact Kaitlyn/Zach to resolve
*   Is the localization guideline document still up to date?
    *   Luc Perkins is working on a PR (1 week deadline, 7/31)

**7/17/2018**

New contributors



*   

Updates & reminders



*   

Agenda



*   Zach Corleissen is out sick :(
*   PR Wrangler (Zach C.)
    *   Queue is growing longer; need to manage it more intentionally
        *   For PRs with no CLA or no response from contributors, it’s OK to close more aggressively
            *   15 days of no response warning, 30 days to close 
            *   Close after 15 days to close if it’s a CLA based issue
            *   Things on hold or in a milestone should probably not be automatically closed
            *   For typo (or simple fixes) with an unsigned CLA, recreating the commit and closing the original is usually best after a period of silence (like 2 weeks)
        *   PR queue bash: Sprint days at WtD Cincinnati (August 20-21)
            *   Will there be remote participation? Zoom meeting? Slack?
            *   Tuesday August 21st is the proposed day for a bash, and whenever people can participate is great. Thanks!
                *   Potentially check in at the sig-docs meeting that day for the status of the bash.
    *   PR queue bash: Mid- to late October
        *   Potentially last opportunity for big, coordinated push for PR queue bashing before the holidays and both KubeCons
*   Docs 1.12 Release Update
    *   Meet the team (Tim, Jim, and Sam)
    *   Meet our strategy (do the same as always and ask lots of questions)
    *   Meet the program pulling PR’s from the milestone for our review
*   Ref docs format
    *   [https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/)
    *   [https://deploy-preview-9536--kubernetes-io-master-staging.netlify.com/docs/reference/generated/kubernetes-api/v1.11/](https://deploy-preview-9536--kubernetes-io-master-staging.netlify.com/docs/reference/generated/kubernetes-api/v1.11/)
    *   Suggestions:
        *   Docs working with engineers to add examples
        *   Docs to add examples directly
        *   “PEFixit day” - taking time to work with engineers to add examples with docs together

**7/10/2018**

New contributors



*   

Updates & reminders



*   New SIG Docs approvers:
    *   Kaitlyn Barnard (CNCF), @kbarnard10 on Slack/GitHub
    *   Zach Arnold (Ygrene), @zparnold on Slack/GitHub
*   Congratulations to Misty Linville, Zach Arnold, and Nick Chase on the 1.11 docs!
*   PR [wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers) (Zach)
    *   Feel free to flesh out the wrangler description
    *   Best guess for shift assignments based on what I know of folks’ schedules through the end of 2018 (KubeCon Shanghai, KubeCon Seattle, WtD Cincinnati, OSCon)
*   Translation repos (Zach)
    *   On agenda for discussion with SIG Arch this week (Thursday, 12pm Pacific)
*   Big thank you to @tengqm for moving all YAML configuration files to a common directory. This involved fifteen PRs. [Issue 9283](https://github.com/kubernetes/website/issues/9283).

Agenda



*   kubernetes.io/community: CNCF (Kaitlyn) will be doing a mock of a site redesign, includes more modern calendar view, other community topics, communication platforms
    *   Future presentation of mock when it’s ready for review
    *   Feedback will be in a related [GH issue](https://github.com/kubernetes/website/issues/7795) for a single source of truth
*   Netlify auth tokens (Misty)
    *   Check back on progress in 1 week 7/17
*   Zoom maintenance (Zach C)
    *   Paris will do an AMA for Zoom moderation
*   REVISIT from 6/26: Docs accuracy discussion with SIG Arch
    *   POSTPONE until Andrew can join--revisit 7/17
*   [Issue with generated docs for kubelet](https://github.com/kubernetes/website/pull/9344) (link is to PR for temporary fix; PR links to well-documented issue) (Jennifer)
    *   Jennifer will open an issue in k/k for the underlying issue -- [PR already exists](https://github.com/kubernetes/kubernetes/pull/66034). But see discussion there and in [this issue in the docs repo](https://github.com/kubernetes/website/issues/9413). Summary: there’s disagreement about whether the fix should be to the kubelet help code or to the gen-docs script. If the latter, we’d need a separate script and process to generate kubelet docs. Right now the script takes care of multiple tools at once.
    *   Merge the PR to fix the immediate problem
    *   Follow up on 7/24 for issue in k/k
    *   It's live: [https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/)
*   Docs contribution landing page is weird. :( (Jared)
    *   [https://kubernetes.io/docs/home/contribute/](https://kubernetes.io/docs/home/contribute/) which is the primary docs contribution page. It’s ironic, but not useful. 
    *   Seems like a new landing page for this content would be the best path forward. Is someone working on this? 
    *   Communty guide: [https://kubernetes.io/docs/imported/community/guide/](https://kubernetes.io/docs/imported/community/guide/)
    *   Misty is working on this; timeline 2 weeks (7/24)
        *   June Harton will share her own observations/work with Misty

**7/3/2018**

NO MEETING--canceled due to the US holiday.

**6/26/2018 (APAC)**

Introductions



*   Zach Corleissen, SIG Docs Lead

Updates & Reminders



*   This meeting follows the [CNCF code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
*   Please speak one at a time
*   Please mute unless you are speaking

Agenda



*   First meeting!
*   Julia: Hugo migration
    *   Ready by June 28th
*   Rajie: How to add new contributors to this meeting time?
    *   Zach: Feel free to add participants directly to the meeting invitation
*   Ian (Korean l10n):
    *   Team meets weekly
    *   New leader: June Yi
    *   How to add reviewers to repository
        *   Zach: I can help you add an OWNERS file
    *   Zach: If June Yi sends GH ID and email, I can help set up repo permissions
*   Zach: contact Taylor Waggoner with request to open Korean translation mailing list
*   Beijing birthday for Kubernetes
    *   21st and 22nd of July
    *   Photos on Twitter! Hooray!
*   Zach will work with Jared to manage agenda invitations

**6/26/2018 (Pacific)**

New contributors



*   Tim, SF, EmLab

Updates & Reminders



*   APAC meeting time starts today at 7pm
    *   (Zach) Proposal to have only one SIG meeting on Tuesdays going forward. For example, next fourth Tuesday (July 24) would have one meeting at 7pm PST.
    *   Consensus: YES
*   Next week (July 3): US holiday 
    *   Zach and Andrew out on vacation
    *   Who can lead? Steve Perry can.
*   New maintainers:
    *   Qiming Teng (@tengqm)
    *   Misty Linville (@mistyhacks)

Agenda



*   (Andrew) [@timothysc](https://github.com/timothysc) brought up in a recent K8s steering committee the question about how to ensure docs are created and maintained when code changes happen. How do we enforce this? They want to empower us to take this on. Do we need a policy/process for the lifecycle of docs (e.g. when a feature is introduced and moves from Alpha to Beta to GA)?
    *   video of discussion: [https://www.youtube.com/watch?v=UKJYy9Oiuv0&t=5m59s](https://www.youtube.com/watch?v=UKJYy9Oiuv0&t=5m59s)
    *   Discussion:
        *   Gate checks in feature development repos would likely require more bandwidth/headcount than we have
        *   SIG Release also wants to track feature release status: we don’t have a formal list of feature status
            *   What about features with multiple moving parts that advance at different stages?
        *   SIG Release spreadsheet: late notification of features that require docs
        *   If we have a hard line against documenting alpha features, that makes it harder to document beta features
            *   If we do document alpha features, we need to require specific guidelines about what docs are required
        *   How to improve feature/area ownership?
            *   Add a SIG to front matter of each document
            *   Work with SIGs for annual review
            *   Surface the last-updated date on the page (Hugo feature)
            *   Use Hugo shortcodes to indicate feature state, instead of having feature state info in prose
            *   Steve: Include a dedicated field in front matter (“sig”)
            *   Work with SIG release to make sure docs are identified/included at earlier phases
        *   Carrots and sticks
            *   Carrot: For SIGs that clean up their docs for a release, highlight them in the release notes
            *   Sticks: If a particular doc stales out/no review within <X> time (1 year?), remove it from docs
                *   Add a label for Technical LGTM, require it annually or doc rots
        *   Revisit in two weeks (7/10)
            *   Andrew will put together a doc of action items
            *   Approach SIG Release, SIG Arch
            *   One immediate approach: ask ourselves, “Does it make sense to assign a single SIG to a topic?”
            *   Ask other SIGs to tag themselves in topics (provide a sample PR)
*   (Misty) 1.11 updates
*   BUMP to 7/10 - (Misty) Retro of branching strategy for 1.11, proposal for 1.12 based on discussions with lots of Git-experienced folks
*   BUMP to 7/10 - (Luc) Questions “Kubernetes Day Two” documentation.
*   (Steve) Moving YAML files to common directory. Several PRs. Example: [https://github.com/kubernetes/website/pull/9236](https://github.com/kubernetes/website/pull/9236) \


**6/19/2018**

New contributors



*   Jim Angel (Austin)

Updates & Reminders



*   (Misty) 1.11 docs update
    *   Things are on track
    *   Ping Misty if you are blocked on a PR, first release candidate (RC) may be cut on Wednesday 6/20
*   (Zach) PR wrangler update (6/19)
    *   Partial fix: repo project to enable monitoring at a glance, make it easier to hand off week-to-week - VERDICT: Don’t do it, it’s a bad idea
    *   Exclude release meister
    *   Given current maintainer numbers, we’re looking at 2 wrangler shifts per maintainer per quarter. 

Agenda



*   Style guide: Should we capitalize Containers? ([#9040](https://github.com/kubernetes/website/pull/9040))
    *   Jennifer’s answer: NO (and yes I have Further Thoughts About Capitalization Generally In The Docs. No surprise there ...) Wait, no, maybe … (like all things style-wise …)
    *   Zach’s answer: I’m strongly ambivalent! It’s useful to distinguish between objects vs. instances, but we apply the distinction inconsistently at best.
    *   Steve's answer: This needs to be part of a larger discussion about capitalizing API objects. Note that Container is an API object even though users don't create Container objects directly. [https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#container-v1-core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#container-v1-core)
    *   Zach: Create clear style guide guidance for capitallzation
        *   Open an issue, allow multiple PRs against it for style guide fixes
        *   Create a checklist with some guidance for different levels of editorial difficulty/passthrough
        *   (Steve) Prioritize markdown fixes \

*   APAC Meeting time
    *   Proposal: start with monthly meetings on Tuesdays at 7pm Pacific:
        *   Wed 7:30am Bangalore
        *   Wed 10am Beijing
        *   Wed 11am Tokyo
        *   Tuesday 10pm Eastern
    *   Advertise to make time visible
    *   Revisit after 3 months
*   Internationalization
    *   Issue open for Korean and Japanese translations: [https://github.com/kubernetes/community/issues/2221](https://github.com/kubernetes/community/issues/2221)
    *   In conversation with Bjorn Erik Pedersen about optimizing k/website for internationalization
*   How best to generate kubectl docs
    *   [https://groups.google.com/forum/?pli=1#!topic/kubernetes-sig-docs/ZMVVNBu2TD8](https://groups.google.com/forum/?pli=1#!topic/kubernetes-sig-docs/ZMVVNBu2TD8)
    *   Whatever we do must be different from what we do right now; it’s difficult, cumbersome, and always changing
    *   Moving targets for correctly documented processes are a time sink
    *   Not just kubectl, but all generated API/component docs
    *   One idea may be to break up single (huge) HTML file into separate files for each command (several hundred HTML files)
        *   Be sure before action; would be difficult to redirect in the future
    *   (Misty) CSS for different refdocs are all slightly different; also look unlike the rest of the site
    *   (Misty) Can we abstract the content from the presentation layer?
        *   (Steve) k/k may have some options for format generation in the generation scripts; we’re currently using the generated Swagger spec
            *   (Steve) Start by investigating Swagger spec
        *   How Docker does it for Swagger API docs ([source](https://github.com/docker/docker.github.io/tree/master/engine/api/v1.37)) ([output](https://docs.docker.com/engine/api/v1.37/)) using [redoc](https://swagger.io/blog/api-development/redoc-openapi-powered-documentation/)
        *   How Docker does it for CLI refs ([YAML](https://github.com/docker/docker.github.io/blob/master/_data/engine-cli/docker_checkpoint_create.yaml) and [stub file](https://github.com/docker/docker.github.io/blob/master/engine/reference/commandline/checkpoint_create.md)) ([output](https://docs.docker.com/engine/reference/commandline/checkpoint_create/))
    *   Involve Qiming Teng in this discussion--Zach will create an offline meeting
*   Anyone else notice that some choices from [https://kubernetes.io/docs/setup/turnkey/aws/](https://kubernetes.io/docs/setup/turnkey/aws/)  disappeared compared to previous  versions of the docs?** **
    *   Short term fixes are possible (Steve will open a PR for orphaned topics)
    *   BUT remember that larger Setup changes are in the works
    *   New SIG Cloud Provider meeting: [https://zoom.us/my/sigcloudprovider](https://zoom.us/my/sigcloudprovider)
        *   https://github.com/kubernetes/community/tree/master/sig-cloud-provider#meetings
    *   Tangentially, Jennifer will be reorganizing kubeadm content post-1.11
*   Anyone besides Brad going to KubeCon Shanghai to help run a Doc Sprint?
    *   Andrew (chenopis) is going too. Current thinking is to craft the Doc Sprint around the Chinese translation efforts to make sure their translation workflow is solid.
*   PRs that fix the double bullet issue via css hacks?

**6/12/2018**

New contributors



*   Takuya Takuda (Japanese translation)

Updates & Reminders



*   Rotations for the PR queue?
    *   last week
        *   Zach/Misty
    *   this week
        *   Brad Topol (intermittent, starting on Friday)
    *   next week
        *   ?
    *   PR assignments: ad hoc no longer working, need to come up with regular shifts
        *   Check in with other SIGs
        *   Report back on 6/19 (Jared will send an email)
*   (Misty) 1.11 docs updates
    *   Docs are going well, PRs assigned out to SIG Docs reviewers
    *   For release-specific PRs, only release meisters should be approving/merging
        *   Function of /lgtm is configurable

Agenda



*   (Andrew) Consider alternating SIG Docs weekly meeting times for APAC participation
    *   7pm Pacific meeting once per month on Mondays
*   (Zach) Review how the prow (k8s-ci-bot) command `/hold` works
*   (Zach) Looking for another sponsor for K8s org membership for @cstoku: [https://github.com/kubernetes/website/pulls/cstoku](https://github.com/kubernetes/website/pulls/cstoku)
    *   (Jennifer) I’m happy to sponsor if he still needs another sponsor.
*   (Steve) Prototype by Qiming in [PR 8965](https://github.com/kubernetes/website/pull/8965). Move all YAML files to a dedicated directory. Our current practice is to put a YAML file in the same directory as the topic that uses it. Two problems: 1) When we move a topic, we have to move all the associated YAML files. 2) When we move a YAML file, we have to update [examples_test.go](https://github.com/kubernetes/website/blob/master/test/examples_test.go).
*   (Andrew, Zach) Status of Hugo migration
    *   If you see doubled bullet points in the TOC for a page, make sure to apply a template shortcode
    *   Steve: Do we need another template type for reference material? (Concepts are fine for now)
    *   We may need a new template for generated refdoc (Steve will follow up in 1 week, 6/19 and two weeks 6/26)
*   (Jennifer, per discussion today in #sig-docs) Guidelines for documenting alpha features (or docs at any specific feature state)
    *   We need guidelines for how to document feature state (Jennifer will follow up in 2-3 weeks)
    *   Tim St. Clair: Don’t document alpha features in the wild; creates a supportability nightmare
        *   Alpha features are disabled on purpose
    *   Another possibility is a standard disclaimer warning
    *   How about having a feature_state variable in the front matter that would trigger an include of a disclaimer at the top of the page for alpha features?
*   (Zach A) Have made a prototype link checker in Go [https://github.com/zparnold/k8s-docs-link-checker](https://github.com/zparnold/k8s-docs-link-checker) what would we like to do with it? Email broken links? Have it create issues in k/website?
    *   Sounds lovely! Limit scope to files in PR
*   Tim St Clair wants to restructure some docs
    *   Wants to reorganize kubeadm

**6/5/2018**

New contributors



*   Yasmary Diaz: Alternate Software
*   Ian Choi: Translation efforts 

Updates & Reminders



*   Rotations for the PR queue?
    *   last week
        *   Zach
    *   this week
        *   Misty
        *   Zach
    *   next week
        *   ?
*   (Misty) 1.11 docs updates

Agenda



*   (Zach) Approvers for individual files: [https://github.com/kubernetes/website/pull/8506](https://github.com/kubernetes/website/pull/8506)
    *   Change “approvers” to “reviewers”
    *   Misty: These should go in the contributors guide!
    *   Steve: will take on remaining legacy approvers and change to reviewers
*   (Misty) Related to the above PR, should we have something in the PR template and maybe in official docs about how other SIGs interact with us, that before a PR is raised which tries to change sig-docs processes, that change should be discussed and approved in a sig-docs meeting? This was also Steve’s comment in the PR.
*   (Misty) Korean translation -- had another query on Slack from someone interested in contributing. Send them to Ian?
    *   Direct them to Ian (#kubernetes-docs-ko)
    *   Zach will create a K8s org repo for Korean translation
    *   Add guidelines for repository structure
*   (Jennifer, Steve) Update docs generation instructions
    *   Steve will update doc generation instructions
    *   New files to generate are now at https://github.com/kubernetes-incubator/reference-docs
    *   Jennifer and Steve will examine and report back 6/12
    *   Steve will generate ref docs for 1.11
*   (Zach) Followup on fixing broken links. Originally (Steve): Would someone have the time/interest to do a big comprehensive fixing of broken links? This would save us the time it takes to handle a bunch of PRs, each of which fixes a few broken links.
    *   Misty and Zach Arnold will take a look at automating checks and performing regular maintenance
    *   Please include @masaya_aoyama
    *   DEFER 6/12 - 6/19
*   (Zach) Doc sprint at KubeCon Shanghai
    *   Brad is going 
        *   Brad: Please at least one other person go too. Zach?
        *   Jared: Will try, but gotta find out from others. 
    *   Jared/Andrew: Google still evaluating how many people it will send. Coordinating with Paris. 
    *   Jberkus: Contributor summit organizer
*   (Andrew) Update Zach on post-Hugo bug bash. 
    *   Do we need another one?
    *   We need to identify ones that Bjorn (@bep) has to address and assign them to him.
    *   Zach C & Andrew will review how to flag issues for BEP, how others can do likewise
    *   REPORT 6/12
*   (Jennifer) Followup on including generated docs in release PRs
    *   1.10: Generating docs happened at release cycle. 
    *   1.11: New commands need to have placeholders or breakage happens 
        *   Found this out with kubeadm
    *   (Misty) Should we ask the testing team to develop a per-PR test that tries to detect PRs raised by those who are not sig-docs maintainers but modify generated content like reference docs? Another potential test is for a PR which attempts to modify both manual and generated files in the same PR
    *   AI: Jennifer and Misty will discuss options going forward. 
*   (Misty) Git class
    *   Is this still a thing?
    *   Git Office hours? 
    *   Decision: Overview of git should be in contributor guide written by Misty. Refer to other resources for in depth knowledge. Misty can answer in-depth questions. 
*   (Brad) New Hugo code base review/walk through?
    *   Would like a walkthrough from BEP on design concerns
    *   +1 from Misty
    *   Office hours or presentation
    *   Andrew: Did @bep implement no-index in netlify.toml, which used to be implemented in the Netlify build commands?
    *   AI: Sync with Andrew offline, follow 
*   (Brad) Best way for folks to provide feedback on blogs (PRs are of course rejected) 
*   (Ian) Korean translation status
    *   Initial pull request: [https://github.com/kubernetes/website/pull/8636](https://github.com/kubernetes/website/pull/8636)
        *   Can Korean team get review from Docs SIG members?
    *   AI: Ask Andrew to review - Jared
*   (Josh Berkus) Looking at reviewing all PR’s since about March on k/k to see if there is a need for Docs
*   (Zach Arnold) Potentially look at working on the weekends doing docs sprints (for those contributors who cannot do a lot during the week because of jobs.
    *   Jared: Work-life balance warning :) - If folks from google want to work on a doc sprint outside of normal work hours, let me know and we’ll figure out time management. 
*   HOLD for 6/19 or 6/26 (Misty) Recommendation for workflows in 1.12
*   HOLD for 6/19 or 6/26 (Zach) Best practices for SIG Docs workflows

**5/29/2018**

New contributors



*   Masaya Aoyama
*   Neha
*   June Harton

Updates & Reminders



*   Zach C is back
*   Rotations for the PR queue?
    *   this week
        *   @zcorleissen
    *   next week
        *   
*   Misty is release meister for 1.11 release

Agenda



*   Steve: Fixing broken links. Would someone have the time/interest to do a big comprehensive fixing of broken links? This would save us the time it takes to handle a bunch of PRs, each of which fixes a few broken links.
    *   Misty and Zach Arnold will take a look at automating checks and performing regular maintenance
*   Steve / Jennifer (via Misty): Generated reference docs update
    *   Several PRs in flight about this and several issues coming in each week
    *   There is a big PR that needs review. Jennifer / Steve to talk about it offline this week.
    *   Kubeadm docs updates are problematic as per @jrondeau
        *   Let’s follow up on June 5 about including generated docs in release PRs
        *   Zach & Jennifer will chat
    *   Big docs PRs coming from outside of sig-docs need to be discussed either in a sig-docs meeting (person can add to agenda and not be at the meeting) or in a GH issue before dropping a huge PR
*   Misty: 1.11 docs updates
    *   Still 7 1.11 features with unknown docs state
    *   Deadline for placeholder PRs for features is past (5/25/2018). Pinged all feature PRs to remind.
    *   Placeholder PRs (or already merged PRs) open for lots of other features
    *   Rebased release-1.11 last week
    *   Working with sig-PM to improve feature tracking spreadsheet and add some stats
*   Misty: `kubernetes/website` security contacts ([#8724](https://github.com/kubernetes/website/issues/8724#issuecomment-392319946))
    *   Zach C will review by June 5
*   Misty: GCE content in open source docs
    *   Style guide guidance for vendor-specific commands in K8s docs
*   Andrew: update Zach on post-Hugo bug bash. 
    *   Do we need another one?
    *   We need to identify ones that Bjorn (@bep) has to address and assign them to him.
    *   Zach C & Andrew will review how to flag issues for BEP, how others can do likewise
*   Brad: New Hugo code base  review/walk through?
    *   Would like a walkthrough from BEP on design concerns
    *   +1 from Misty
    *   Office hours or presentation
    *   Andrew: Did @bep implement no-index in netlify.toml, which used to be implemented in the Netlify build commands?
*   Brad: Best way for folks to provide feedback on blogs (PRs are of course rejected) 
*   Andrew: Going to WtD Cincinnati (Aug 18-22, 2018)
    *   Wants to run a doc sprint for K8s

**5/22/2018**

New contributors



*   June
*   Riona - Google, open source

Updates & Reminders



*   Rotations for the PR queue?
    *   this week
        *   @steveperry-53
        *   @bradtopol
    *   next week
        *   Jennifer Rondeau?
*   Misty is release meister for 1.11 release
*   Zach C returning next week

Agenda



*   post-Hugo bug bash
    *   [Spreadsheet](https://docs.google.com/spreadsheets/d/1MBwcTaDo9tNUOxfemH065v1leghO-pfYM0b2rIiHotM/edit#gid=0)
    *   any issues that require HTML workaround or direct fix to Blackfriday should be assigned to Bjorn (@bep)
    *   things that look like a list in a code block are not rendered correctly
    *   use GitHub Issues, tag w/ "Needs Tech Review", assign/cc @bep
    *   can use Hugo server for local dev and preview
        *   need to update docs for how to run local copy of site w/ Hugo server
*   k/website + prow
    *   go to k8s community to get support for feature branches and rebasing workflow
    *   AI: @chenopis to review test-infra [issue](https://github.com/kubernetes/test-infra/issues/8058) and get more sponsors
*   [DEFERRED 5/29] Reviews on WIP pull requests [Zach]
*   (Chris Negus) 
    *   How to fix left nav issues? [Example PR](https://github.com/kubernetes/website/pull/8635)
    *   Ok to add in changes re: Fedora content -- create PR against `master`
*   (Jennifer Rondeau)
    *   Issues discussion upstream
    *   Generated kubeadm docs
*   (Andrew Chen)
    *   break out onboarding to different tiers
        *   use existing env, like Play w/ Kubernetes
        *   Minikube
        *   dev/prod/hosted environment
    *   Neha to revamp app dev foundational user journey, including Minikube docs
        *   

**5/15/2018**

New contributors



*   Neha Dhawan - Google intern with us through July :)
*   Melissa Anderson - TW @ Digital Ocean

Updates & Reminders



*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   **We need wranglers!**
        *   Jennifer can help w/ some PR reviews
        *   
*   Misty is release meister for 1.11 release
    *   [transition gist](https://gist.github.com/zacharysarah/c7dfb90fe1d48762f48d8d6cbb0040b2)
    *   [previous playbook](https://github.com/kubernetes/sig-release/blob/master/release-process-documentation/documentation-guides/update-release-docs-new.md)
*   Zach C is on vacation for the next two weeks, returning 5/29
*   Steve will miss today's meeting. Has been working on Hugo migration issues. For example, moving topics to new directories so they appear in the proper place in the left nav.

Agenda



*   SIG Docs Summit Recap [Jared, Tom]
*   Hugo migration bug bash
    *   Fri - Jennifer, Andrew, Misty (needs admin priv), Brad
*   [DEFERRED 5/29] Reviews on WIP pull requests [Zach]
    *   
*   PR wrangling
    *   Large percentage of PRs on hold
        *   We may not want to merge many of these
        *   Need style guidance
        *   Need a way to address them so they don't bloat the queue
    *   Missing CLAs
        *   Misty working on improving contributor docs, which should improve this.
        *   Some may not be sure if they can sign on behalf of company or whether they should sign individually instead.
        *   If they explicitly say they can't CLA then we should close the PR.
    *   **Need a policy to close stale PRs**
        *   Ping them, if we don't hear back within **one month**, close the PR.
*   Kubernetes governance [Andrew]
    *   automation overreach
    *   how do we escalate when we run into issues?
        *   e.g. non-SIG Docs maintainer was able to merge because of Kubernetes org admin?
        *   prow policies applied w/o our knowledge
    *   Andrew to:
        *   talk to Paris
        *   ask at Community meeting
        *   ask Steering Committee
*   (May 15): How did a PodSecurityPolicy change break our test build last week (Jennifer)
    *   (THANK YOU TO QIMING for fixing it!!: [https://github.com/kubernetes/website/pull/8086](https://github.com/kubernetes/website/pull/8086))
    *   Revisit documentation options that don't appear in generated docs. See [https://github.com/kubernetes/website/pull/7887](https://github.com/kubernetes/website/pull/7887) and [https://github.com/kubernetes/website/issues/8280](https://github.com/kubernetes/website/issues/8280)
    *   DEFERRED to 15 May - Docs tests: no consistent ownership
    *   Automation governance discussion
    *   **Need to add task to release meister process: update release we are testing against to be current release not master**
        *   File to update is [https://github.com/kubernetes/website/blob/master/test/examples_test.go](https://github.com/kubernetes/website/blob/master/test/examples_test.go)
        *   High-level additional requirements: check release/release notes for changes to API resources that the tests call.

**5/1/2018**

New contributors



*   nope

Updates & Reminders



*   **No weekly meeting next week**, 5/8 (Write the Docs)
*   There WILL BE a Kubernetes session during the Write the Docs Writing Day (May 6, Portland). Andrew and Jennifer will run it, with Misty and Steve to provide their special expertises.
*   Rotations for the PR queue (Andrew) - Update
    *   Jennifer Rondeau - 4/30 - 5/6
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Who’s up next? We’ll need coverage during WTD!
*   **SIG Docs Summit: Wednesday, May 9, Portland**
    *   [https://github.com/kubernetes/website/issues/8116](https://github.com/kubernetes/website/issues/8116) 

Agenda



*   5/1: Chenopis will run this meeting.
    *   Jared and Zach will send update on Docs Sprint / Contrib Summit. 
*   Write the Docs @ Portland
    *   Sprint instructions: [https://bit.ly/2rcZ3YU](https://bit.ly/2rcZ3YU)
    *   [Hugo migration issues](https://github.com/kubernetes/website/milestone/20) from KubeCon docs sprint need to be triaged.
    *   Hugo preview: [https://hugo-migration.docs.kubernetes.io](https://hugo-migration.docs.kubernetes.io)
    *   Migration script will be run again before final migration, so we can continue working and merging content like normal.
    *   **Decision: avoid Hugo migration issues for doc sprint.**
    *   **Jennifer will get the donuts!**
    *   **Issues for doc sprint: [https://github.com/kubernetes/website/projects/8](https://github.com/kubernetes/website/projects/8)**
    *   **Andrew will dump issues into assignment [spreadsheet](https://bit.ly/2KrwPT7).**
*   PR wrangling
    *   Large percentage of PRs on hold
    *   Missing CLAs
        *   Misty working on improving contributor docs, which should improve this.
        *   Some may not be sure if they can sign on behalf of company or whether they should sign individually instead.
    *   **Need a policy to close stale PRs**
*   [Jennifer] Do we want to standardize on how to document options that don’t appear in generated docs? (Related discussion: should they do so if they’re deprecated). See [https://github.com/kubernetes/website/pull/7887](https://github.com/kubernetes/website/pull/7887) and [https://github.com/kubernetes/website/issues/8280](https://github.com/kubernetes/website/issues/8280)
    *   Is a bug and should be fixed in 1.10. Need to verify.
    *   Need to understand why this is happening.
        *   Followup on 8280: see [https://github.com/kubernetes/kubernetes/blob/master/pkg/kubeapiserver/options/admission.go#L67ff](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubeapiserver/options/admission.go#L67ff). New flag is not properly documented. See especially l75. (Still need to dive after where flag is defined)
    *   Need to coordinate w/ upstream efforts to fix it.
*   **Kubernetes version needs to be updated in the test code for each release as part of Docs Lead process.**
*   Kubernetes governance
    *   automation overreach
    *   Changing automatically protecting branches: [https://github.com/kubernetes/test-infra/pull/7900](https://github.com/kubernetes/test-infra/pull/7900)
    *   chenopis@
        *   Figure out chain of escalation
        *   Find champion on steering committee
        *   Work w/ governance to establish policy
*   Refactor Minikube docs [[issue #6339](https://github.com/kubernetes/website/issues/6339)]
    *   Jennifer and Misty team up!
*   DEFERRED (May 15): How did a PodSecurityPolicy change break our test build last week (Jennifer)
    *   (THANK YOU TO QIMING for fixing it!!: [https://github.com/kubernetes/website/pull/8086](https://github.com/kubernetes/website/pull/8086))
    *   Revisit documentation options that don't appear in generated docs. See [https://github.com/kubernetes/website/pull/7887](https://github.com/kubernetes/website/pull/7887) and [https://github.com/kubernetes/website/issues/8280](https://github.com/kubernetes/website/issues/8280)
    *   DEFERRED to 15 May - Docs tests: no consistent ownership
    *   Automation governance discussion

**4/24/2018**

New contributors



*   Jason Van Brackel - Rancher - Member of Sig Windows
*   Zach Arnold - here to help out with release for 1.11 

Updates & Reminders



*   Rotations for the PR queue (Andrew) - Update
    *   Andrew’s turn this week!
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Andrew Chen - 4/23 - 4/29 (Andrew taking this since Joe covered for him) 
    *   Jennifer Rondeau - 4/30 - 5/6
*   SIG Docs Summit: Wednesday, May 9, Portland
    *   [https://github.com/kubernetes/website/issues/8116](https://github.com/kubernetes/website/issues/8116) 

Agenda



*   Jennifer and Joe both send apologies for today’s meeting
*   Next two weekly meetings: what to do?
    *   5/1: Jared, Brad, and Zach will both be at KubeCon
        *   Chenopis will run this meeting.
        *   Jared and Zach will send update on Docs Sprint / Contrib Summit. 
    *   5/8: Many folks will be at Write the Docs
        *   Cancel this meeting since the Docs Meetup is the same week.
*   ANNOUNCEMENT: There WILL BE a Kubernetes session during the Write the Docs Writing Day (May 6, Portland). Andrew and Jennifer will run it, with Misty and Steve to provide their special expertises.
*   Paris agenda slot 
    *   Contributor Experience Deep Dive for Contrib Summit @ Kubecon
    *   Want to add new facet to the contrib / developer guide on api conventions
    *   Want to do a mini doc sprint 
    *   35 minutes
    *   Suggestions:
        *   Do a demo
        *   Brainstorming session?
        *   Breakout session afterwards?
*   Git tech talk (Misty)
    *   Let’s do a git tech talk for best practices! (June)
    *   Considering doing office hours.
    *   Show people how to set up their environment (TBD)
    *   Available for one-off Git questions in Slack as time permits
*   Zach gives an update on Hugo?
    *   Very close to “flipping the switch” in Hugo. 
    *   Syncing with chinese translation team to give them the heads up.
    *   Aiming for Thursday - 4/26
*   DEFERRED (June 1): How did a PodSecurityPolicy change break our test build last week (Jennifer)
    *   (THANK YOU TO QIMING for fixing it!!: [https://github.com/kubernetes/website/pull/8086](https://github.com/kubernetes/website/pull/8086))
    *   Andrew will approach test-infra for notification about upstream changes
    *   DEFERRED to 1 June - Docs tests: no consistent ownership

**4/17/2018**

New contributors



*   

Updates & Reminders



*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Brad Topol - 4/16 - 4/22
    *   Andrew Chen - 4/23 - 4/29 (Andrew taking this since Joe covered for him) 
    *   Jennifer Rondeau - 4/30 - 5/6

Agenda



*   SIG Docs Summit: Wednesday, May 9, Portland
    *   [https://github.com/kubernetes/website/issues/8116](https://github.com/kubernetes/website/issues/8116) 
    *   Attendees: Brad, Zach, 
*   A few notes from PR wrangling (heckj)
    *   Updated the PR wrangler gist
    *   More changes starting to appear for generated content - closing politely and trying to point people to the relevant source to update upstream
        *   Example: [https://github.com/kubernetes/website/pull/8108](https://github.com/kubernetes/website/pull/8108)
*   How did a PodSecurityPolicy change break our test build last week (Jennifer)
    *   (THANK YOU TO QIMING for fixing it!!: [https://github.com/kubernetes/website/pull/8086](https://github.com/kubernetes/website/pull/8086))
    *   Andrew will approach test-infra for notification about upstream changes
    *   DEFERRED to 1 June - Docs tests: no consistent ownership
*   Release-1.10 branch and merge conflicts (Zach)
    *   Make sure you aren’t committing conflict markers
    *   Zach will work with Misty to hard reset release-1.10 to master
*   Git tech talk (Misty)
    *   Let’s do a git tech talk for best practices! (June)
*   External content (Rajie)
    *   What should we do with useful external content?
        *   [https://eyskens.me/building-a-kubernetes-ingress-controller/](https://eyskens.me/building-a-kubernetes-ingress-controller/)
        *   Link to it but don’t host
*   Summit: Define K8s content strategy very clearly

Reference:



*   Hugo migration meeting notes
    *   Early discussion summary: [https://gist.github.com/heckj/2129560e68c6104690dfa6bf9099c3e9](https://gist.github.com/heckj/2129560e68c6104690dfa6bf9099c3e9)
    *   16 apr 2018 progress/sync meeting [https://gist.github.com/heckj/ab20589283e2029f0a6d928173ce1767](https://gist.github.com/heckj/ab20589283e2029f0a6d928173ce1767)

**4/10/2018**

New contributors



*   Everything old is new again. 

Updates & Reminders



*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Joe Heck - 4/10 - 4/15
    *   Brad Topol - 4/16 - 4/22
    *   Joe Heck - 4/23 - 4/29 (Andrew might take this one) 
    *   Jennifer Rondeau - 4/30 - 5/6
*   Reminder: no cookie licking!

Agenda



*   PR/Style Guide suggestions (rajie)
    *   Accessibility guidelines: [https://github.com/kubernetes/website/pull/7991](https://github.com/kubernetes/website/pull/7991) Reference: [https://www.w3.org/WAI/intro/wcag](https://www.w3.org/WAI/intro/wcag)
    *   Capitalization/usage additions: [https://github.com/kubernetes/website/pull/8017](https://github.com/kubernetes/website/pull/8017)
    *   Chatbot?
        *   Do some user acceptance testing for search at Write the Docs?
*   Hugo Migration plan/update (zach)
    *   Page weighting: how will it be done in Hugo?
*   Need some cross-browser CSS help! (heckj)
    *   Pull-down nav rendering issue (safari only): [https://github.com/kubernetes/website/issues/7801](https://github.com/kubernetes/website/issues/7801)

**4/3/2018**

New contributors



*   Marcus Heese (@mheese) 

Updates & Reminders



*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   Rescheduled for 4/4
*   Master branch strategy (Andrew, Steve, Jennifer) [[slides](https://docs.google.com/presentation/d/1OIs-Hdq7ZaVyuwSzzCw-ww3o9rSn5EMEpy5QvpF3Oik/edit?usp=sharing)]
    *   Today (4/3) at 1pm
        *   Add Nick Chase, Tall Tom
*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Steve Perry - 4/2 - 4/8
    *   Andrew Chen - 4/9 - 4/15
    *   Brad Topol - 4/16 - 4/22
    *   Joe Heck - 4/23 - 4/29

Agenda



*   New Provider Hosting (Chris Hoge) (with impact on docs for providers)
    *   [https://github.com/kubernetes/community/pull/1942](https://github.com/kubernetes/community/pull/1942)
    *   Zach will set up a working meeting, invite: Joe Heck, Andrew, Steve
*   API group naming (@mheese)
    *   Steve will republish with long names 
    *   Nobody present understood why there currently are short names in the reference docs
*   Issues: No cookie licking Good First Issues! (Zach): [https://github.com/kubernetes/website/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22](https://github.com/kubernetes/website/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22)
    *   If you don’t have a PR after 3 days on a Good First Issue then it is no longer yours. 
    *   Mark work in progress PRs as “WIP” in the title of the PR. 
*   Serving Kubernetes fonts from a server other than Google
    *   Google font serving slows down China web load times
    *   Need an alternative font server
        *   For example: https://landscape.cncf.io/ has fonts served from elsewhere
    *   Suggestion: 
        *   Serve fonts directly from kubernetes.io
        *   Fallback font that users have locally on their machine. 
        *   Zach will file an issue. (Arial for ALL!) 
*   Considering improvements to CI/CD (Zach)
    *   Link checker
    *   Linter
    *   Proposal:
        *   Check after build is done, like the netlify check or travis. 
        *   Options: Xenu link checker, linklint, w3c checker (very slow)
*   Proposing Korean documentation translation (@ianychoi): \
[https://groups.google.com/forum/#!topic/kubernetes-sig-docs/akgMNB_pHXU](https://groups.google.com/forum/#!topic/kubernetes-sig-docs/akgMNB_pHXU) 
    *   Get the Chinese translation working first before we give guidance to Korean team
*   SIG Docs contributor process (Andrew / Jared) 
    *   Writing out how to become a member of the SIG Docs org and a site maintainer. 
    *   Reorging the docs contributor content / style guide / doc processes content.
    *   This page needs love/improvements/context ([https://kubernetes.io/docs/home/contribute/participating/](https://kubernetes.io/docs/home/contribute/participating/))
    *   Joe, Steve, Rajie, Misty, Tall Tom, Jennifer

Add results as a single point of truth to community site/resources

**3/27/2018**

New contributors



*   Rajie (hi officially!)
*   Misty Stanley-Jones (hello!)

Agenda



*   1.10 update (jrondeau, nickchase)
    *   Posted yesterday!
    *   Standing items:
        *   Steve is taking care of generated docs
    *   Some maintenance still required:
        *   1.9 deprecation banner
    *   Documentation for release process needs more love
    *   Release notes: Andrew will help import 
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   Check in on initial assessment tomorrow
    *   Zach will track overall process (in a GH issue)
    *   Paris: community guidelines for rolling out big changes: SIG contribex charter
*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Jennifer Rondeau - 3/26 - 4/1
        *   PRs frozen for 1.10 release day, feel free to remove /holds
    *   Brad Topol - 4/1
    *   Steve Perry - 4/2 - 4/8
    *   Andrew Chen - 4/9 - 4/15
    *   Brad Topol - 4/16 - 4/22
    *   Joe Heck - 4/23 - 4/29
*   Master branch strategy (Andrew, Steve, Jennifer) [[slides](https://docs.google.com/presentation/d/1OIs-Hdq7ZaVyuwSzzCw-ww3o9rSn5EMEpy5QvpF3Oik/edit?usp=sharing)]
    *   Changing SIG Docs branch workflow to match the rest of K8s orgs/SIG workflows
    *   This will be a greatly complex project; deserves full, dedicated treatment in Q3
    *   Misty: let’s consider alternate workflows to optimize how we do things
    *   Work with Paris to present migration proposal to community (meeting)
    *   Meeting to talk about branch strategy
        *   Andrew, Steve, Misty, Jennifer, Zach, Jared, Joe, Brad
*   New Provider Hosting (Chris) (with impact on docs for providers)
    *   [https://github.com/kubernetes/community/pull/1942](https://github.com/kubernetes/community/pull/1942)
    *   Follow up on 3/27 after reviewing PR
    *   Migrated current cloud providers to new repositories.
    *   Followup next week. 
*   Linter? (Zach)
*   Offline docs (Zach)

**3/20/2018**

New contributors



*   Jmosco - worked on a few PRs already

Agenda



*   1.10 update (jrondeau)
    *   Reverting an Alpha feature
    *   Content looking good, go/no go signal on Thursday for Monday release
*   Front loading docs in the release process (calebmiles/chenopis) 
    *   KEP process: how can we use it to front load early quality drafts of release notes?
    *   Pilot example of KEP for docs: [https://github.com/kubernetes/community/blob/master/keps/sig-cluster-lifecycle/0003-cluster-api.md](https://github.com/kubernetes/community/blob/master/keps/sig-cluster-lifecycle/0003-cluster-api.md)
    *   Question from Zach: Does this fix early contribution in the way we hope?
    *   Question from Steve: is this across all sig-groups for 1.11? Does this replace PRs for release notes?
    *   Andrew/Caleb: it’s a test-drive
*   New/ongoing project to improve release note content at the source (Jennifer). See [new issue for discussion ](https://github.com/kubernetes/website/issues/7791)
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   Bjorn Erik Pedersen: [http://bep.is/](http://bep.is/)
    *   PR Jubilee (what else to call it?): 
        *   PRs still open at merge may cause merge conflicts
        *   Need to merge/close all PRs in progress prior to merge
    *   Initial stage: two “days” of work for initial assessment, resulting in a PR
    *   SIG Docs maintainers: input/review needed on upcoming migration PRs.
        *   Look for changes to site tooling
    *   Zach will track overall process (in a GH issue)
    *   Paris: community guidelines for rolling out big changes: SIG contribex charter
*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute.
    *   Steve Perry (last week) Lots of Tech Review Needed and Docs Open Issues.
    *   Zach Corleissen - 3/19 - 3/25 (Currently at 65 PRs, check your queues)
    *   Jennifer Rondeau - 3/26 - 4/1
    *   Andrew Chen - 4/2 - 4/8
*   Next meeting 3/27: Let’s revisit the master branch strategy (Andrew, Steve, Jennifer)
*   New Provider Hosting (Chris) (with impact on docs for providers)
    *   [https://github.com/kubernetes/community/pull/1942](https://github.com/kubernetes/community/pull/1942)
    *   Follow up on 3/27 after reviewing PR
*   Steve will reopen issues related to the migration from user-guides to tasks, tutorials, and concepts.
*   Homonym pun (Jennifer)

**3/13/2018**

New contributors



*   Tom van Waardhuizen (Google)
*   Rajie Kodhandapani

Agenda



*   1.10 update (jrondeau)
    *   According to past docs wranglers, we’re in unusually good shape. A few stragglers, will be merged today or tomorrow
    *   Jennifer might need help with generated docs; already alerted Steve
*   [Spell-checking PR](https://github.com/kubernetes/website/pull/7691): take a different approach? (Steve, Jennifer)
    *   There are already some excellent spell checkers out there. 
    *   Long explanation and rationale for rejection in the PR comment 
    *   Make sure note about local spellchecker gets added to contrib guidelines
*   DEFERRED (1 week, 3/20, depends on state of release) Reviewing release note copy: possibly rotate reviewers from sig-docs throughout release cycle, the way we’ve started doing for rotating docs repo PRs? (Jennifer)
*   Future of [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) (Ben Hall) - Update? 
    *   UPDATE from Ben: “we need to do it via a HTML or JavaScript tag instead of any github based approach”, “Should have it tested and a PR ready for next week”
    *   Related issue: [https://github.com/kubernetes/website/issues/7506](https://github.com/kubernetes/website/issues/7506)
    *   Yes, redirect plz. 
    *   AI: Ben will submit a PR, Steve will review. 
*   Brad Topol for SIG Docs maintainer (Zach)
    *   Approved! Hoorays! 
*   New team: kubernetes-blog-maintainers (Zach)
    *   Part of getting ready for blog migration
    *   Sarah Conway, Bob Hrdinsky, others who had been updating the blog. 
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   We are proceeding. Hoorays!
    *   Working with a contractor. Bjorn Eric Pederson
    *   Aiming to schedule start for first week of April (after 1.10 goes out the door). 
    *   Estimated to take 3-4 weeks
        *   Expect a doc freeze at some point in the migration process
    *   Migration work should be invisible to contributors/users
    *   What you have to do differently: nothing (hopefully)
*   Rotations for the PR queue (Andrew) - Update
    *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
    *   Every week, someone volunteers as tribute. 
    *   Steve Perry - 3/12 - 3/18
    *   Zach Corleissen - 3/19 - 3/25
    *   Jennifer Rondeau - 3/26 - 4/1
    *   Andrew Chen - 4/2 - 4/8
*   Docs proposal: Move away from hosting Operating System/platform-specific and link to it instead [Zach]
    *   Highly out of date: [https://kubernetes.io/docs/setup/pick-right-solution/](https://kubernetes.io/docs/setup/pick-right-solution/)
        *   Cloud Provider working group: all in favor
        *   Impact assessment: [https://github.com/kubernetes/website/pull/7501#issuecomment-371342046](https://github.com/kubernetes/website/pull/7501#issuecomment-371342046) (thanks Philip Mallory!)
        *   PR in progress: [https://github.com/kubernetes/website/pull/7501](https://github.com/kubernetes/website/pull/7501)
*   Twitter handles for Kube Doc folks [Brad]
    *   Go for it: offline discussion about ownership and best practices
        *   Jennifer Rondeau (@Bradamante)
*   K8s site metrics (Andrew, Joe)
    *   Process for determining/adding metrics is just starting

**3/6/2018**

New contributors



*   

Agenda



*   1.10 update (heckj, jrondeau)
    *   
*   Future of [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) (Ben Hall) - Update? [DEFERRED 1 week]
    *   Related issue: [https://github.com/kubernetes/website/issues/7506](https://github.com/kubernetes/website/issues/7506)
    *   Yes, redirect plz. 
    *   AI: Ben will submit a PR, Steve will review. 
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   
*   Rotations for the PR queue (Andrew) - Update
    *   Week 1: Joe, what did you learn?
    *   Suggested process:
        *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
        *   Every week, someone volunteers as tribute. 
        *   Joe Heck - This week 
        *   Brad Topol - 3/5 - 3/11
        *   Steve Perry - 3/12 - 3/18
        *   Zach Corleissen - 3/19 - 3/25
        *   Jennifer Rondeau - 3/26 - 4/1
        *   Andrew Chen - 4/2 - 4/8
*   Docs proposal: Move away from hosting Operating System/platform-specific and link to it instead [Zach]
    *   Highly out of date: [https://kubernetes.io/docs/setup/pick-right-solution/](https://kubernetes.io/docs/setup/pick-right-solution/)
        *   Indicative of larger problems with freshness/staleness and maintaining dual source
        *   Better to link to canonical docs from an OS/platform’s own community
        *   sig/open-stack lead agrees: externalize docs
        *   Discuss this at the cloud provider working meeting (when is that? Zach will find out!)
        *   ZACH: Action item to assess externalizing impact
    *   PR in progress: [https://github.com/kubernetes/website/pull/7501](https://github.com/kubernetes/website/pull/7501)
*   First annual SIG Docs Summit: May 9, Portland, CENTRL Office
    *   Day after Write the Docs in Portland
    *   Attendees?
        *   Zach Corleissen
        *   Jared Bhatti
        *   Jennifer Rondeau
        *   Chris Hoge
        *   Andrew Chen
        *   Steve Perry
        *   Brad Topol
        *   Stephen Augustus
*   Twitter handles for Kube Doc folks [Brad]
    *   ZACH: Find out CNCF/LF policy (if any) about Twitter accounts for individual K8s SIGs

How many folks on the call have a twitter handle that they use with some frequency? When I tweet doc related information tweets I would like to include those who are active to help amplify our messages. For example, last night I tweeted out Joe Heck’s review document.  I got lots of interest. Please share your twitter handles!

@bradamante

@stephenaugustus

@bradtopol 

@zachorsarah

For next meeting: 



*   Update on I18n support (Chenopis) (next week, 3/13)

**3/6/2018**

New contributors



*   None this week :-(

Agenda



*   1.10 update (heckj, jrondeau)
    *   4 features without doc information; looks as though 1 may not make it into the release, and 1 possibly doesn’t need docs. Tracking 2X/day at this point
    *   Some PRs that should have been submitted to master were submitted to 1.10 branch. This has made things a bit more challenging to track. Ideas for solutions? Add to review guidelines? Promote branch PR practices better?
*   Future of [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) (Ben Hall) - Update? [DEFERRED]
    *   Related issue: [https://github.com/kubernetes/website/issues/7506](https://github.com/kubernetes/website/issues/7506)
    *   Yes, redirect plz. 
    *   AI: Ben will submit a PR, Steve will review. 
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)) (Jared B, Zach)
    *   
*   Rotations for the PR queue (Andrew) - Update
    *   Week 1: Joe, what did you learn?
    *   Suggested process:
        *   PR Wrangler guidelines: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
        *   Every week, someone volunteers as tribute. 
        *   Joe Heck - This week 
        *   Brad Topol - 3/5 - 3/11
        *   Steve Perry - 3/12 - 3/18
        *   Zach Corleissen - 3/19 - 3/25
        *   Jennifer Rondeau - 3/26 - 4/1
        *   Andrew Chen - 4/2 - 4/8
*   Docs proposal: Move away from hosting Operating System/platform-specific and link to it instead [Zach]
    *   Highly out of date: [https://kubernetes.io/docs/setup/pick-right-solution/](https://kubernetes.io/docs/setup/pick-right-solution/)
        *   Indicative of larger problems with freshness/staleness and maintaining dual source
        *   Better to link to canonical docs from an OS/platform’s own community
        *   sig/open-stack lead agrees: externalize docs
        *   Discuss this at the cloud provider working meeting (when is that? Zach will find out!)
        *   ZACH: Action item to assess externalizing impact
    *   PR in progress: [https://github.com/kubernetes/website/pull/7501](https://github.com/kubernetes/website/pull/7501)
    *   Related to Update CoreOS install docs (Stephen Augustus, following)
*    [Tentative] Update CoreOS install docs [Stephen Augustus]
    *   [https://github.com/kubernetes/website/issues/6859](https://github.com/kubernetes/website/issues/6859)
*   First annual SIG Docs Summit: May 9, Portland, CENTRL Office
    *   Day after Write the Docs in Portland
    *   Attendees?
        *   Zach Corleissen
        *   Jared Bhatti
        *   Jennifer Rondeau
        *   Chris Hoge
        *   Andrew Chen
        *   Steve Perry
        *   Brad Topol
        *   Stephen Augustus
*   Twitter handles for Kube Doc folks [Brad]
    *   ZACH: Find out CNCF/LF policy (if any) about Twitter accounts for individual K8s SIGs

How many folks on the call have a twitter handle that they use with some frequency? When I tweet doc related information tweets I would like to include those who are active to help amplify our messages. For example, last night I tweeted out Joe Heck’s review document.  I got lots of interest. Please share your twitter handles!

@bradamante

@stephenaugustus

@bradtopol 

@zachorsarah

For next meeting: 



*   Update on I18n support (Chenopis) (next week, 3/13)

**2/27/2018**

**Note: **Jared Bhatti leads this meeting (Zach is traveling)

New contributors



*   Philip Mallory - Google, Kubernetes Engine
*   Abraham - IBM, lurking :)

Agenda



*   1.10 update (heckj, jrondeau)
    *   Tagging PRs related to the upcoming release
    *   Most have been merged into the 1.10 branch
*   Updating interactive tutorials [Ben Hall] - Update? (heckj, stevepe)
    *   All done. Hooray!
    *   All using the new minikube. Much improved. Thank you!
*   Future of [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) (Ben Hall)
    *   Related issue: [https://github.com/kubernetes/website/issues/7506](https://github.com/kubernetes/website/issues/7506)
    *   Yes, redirect plz. 
    *   AI: Ben will submit a PR, Steve will review. 
*   Migrate kubernetes.io from Jekyll to Hugo ([Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#))
    *   Will have update next week. 
*   Rotations for the PR queue (Andrew)
    *   PR are going up. Suggest assigning a “PR Wrangler” for the week
    *   Keeping it <50 PRs
    *   Time to first response/closing - actually doing pretty good. 
    *   Suggested process:
        *   Every week, someone volunteers as tribute. 
        *   Joe Heck - This week 
        *   Brad Topol  - Next week
        *   Steve Perry - Third week
    *   Heckj wrote up a "how I do this.." for Brad: [https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17](https://gist.github.com/heckj/50df12da789ee8402a937be83a88ef17)
*   First annual SIG Docs Summit: May 9, Portland, CENTRL Office
    *   Day after Write the Docs in Portland
    *   Attendees?
        *   Zach Corleissen
        *   Jared Bhatti
        *   Jennifer Rondeau
        *   Chris Hoge
        *   Andrew Chen
        *   Steve Perry
        *   Brad Topol
*   [Tentative] Update CoreOS install docs [Stephen Augustus]
    *   [https://github.com/kubernetes/website/issues/6859](https://github.com/kubernetes/website/issues/6859)

For next meeting: 



*   Update on Hugo (jaredb, zach)
*   Update on I18n support (Chenopis) (in 2 weeks)

**2/20/2018**

**Note: **Jared Bhatti leads this meeting (Zach is running a doc sprint at IBM Index)

New contributors



*   Chris Hoge - Sig Open-Stack

Agenda



*   Updating interactive tutorials [Ben Hall]
    *   [https://github.com/kubernetes/website/issues/7240#issuecomment-365327820](https://github.com/kubernetes/website/issues/7240#issuecomment-365327820)
    *   Steve: Let's talk about both locations:
        *   [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html)
        *   [https://kubernetes.io/docs/tutorials/kubernetes-basics/](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
        *   (Heckj: there’s a 3rd location for the hosted content: [https://github.com/katacoda-scenarios/kubernetes-bootcamp-scenarios](https://github.com/katacoda-scenarios/kubernetes-bootcamp-scenarios))
*   [Ilya] status update on getting started guides (~1min) ([kubernetes/website#7278](https://github.com/kubernetes/website/issues/7278))
*   Migrate kubernetes.io from Jekyll to Hugo
    *   [Discussion Doc](https://docs.google.com/document/d/1XAmCSueuhU6YOvO2ensdiJX63jVapr9XijwY3HlENl4/edit#)
*   First annual SIG Docs Summit: May 9, Portland, CENTRL Office
    *   Day after Write the Docs in Portland
    *   Attendees?
        *   Zach Corleissen
        *   Jared Bhatti
        *   Jennifer Rondeau
        *   Chris Hoge
        *   Andrew Chen
        *   Steve Perry
        *   
*   [Tentative] Update CoreOS install docs [Stephen Augustus]
    *   [https://github.com/kubernetes/website/issues/6859](https://github.com/kubernetes/website/issues/6859)

For next meeting: 



*   Rotations for the PR queue (Andrew)
*   Future of [https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) (Ben Hall)

**2/13/2018**

New attendees?



*   None

Agenda items



*   Consensus/Guidance for issue #7232([https://github.com/kubernetes/website/pull/7232](https://github.com/kubernetes/website/pull/7232)) - how to represent API changes with different versions (heckj, Qiming)
    *   related PR: [https://github.com/kubernetes/website/pull/7385](https://github.com/kubernetes/website/pull/7385)
    *   a) changing the front-matter of the markdown file per Steve
    *   b) using the {% feature state%} jekyll add-in - [https://kubernetes.io/docs/home/contribute/includes/#feature-state-code](https://kubernetes.io/docs/home/contribute/includes/#feature-state-code)
    *   c) comments in YAML about the min version required (getting common, but not consistent)
        *   What do we do about referencing earlier versions? Do we include a comment with what k8s.io v.prior used?
    *   Joe: Send an email to approvers: check version information, make sure that include statements are included, that samples are updated, that version comments in samples are removed
    *   Steve: No blind updates to YAML files; carefully review for versioning
*   Follow up on responsiveness, roles & responsibilities for reviewers and maintainers (heckj)
    *   PR: [https://github.com/kubernetes/website/pull/7133](https://github.com/kubernetes/website/pull/7133)
        *   Comments due by Thursday, 2/15
    *   For example, what expectations do we have & should we have of a reviewer or approver in handling PRs, issues, etc.
    *   Clearer reviewer guidelines
        *   Set an SLO for when we’ll respond to PRs: 1 week
        *   Do we need to codify standards for reviewers: for example, don’t merge a PR assigned to someone else? Is our current method sustainable?
    *   As we're talking about SLO for getting attention on PRs, Paris shared this view of how we're doing today:
        *   [https://k8s.devstats.cncf.io/d/000000013/first-non-author-activity?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&from=1454786880740&to=1516735680740](https://k8s.devstats.cncf.io/d/000000013/first-non-author-activity?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&from=1454786880740&to=1516735680740)
    *   

**2/6/2018**

New attendees?



*   Kris Nova
*   PMallory
*   Ivan Font
*   Matt Dorn

Follow-ups from last meeting (10 minutes max): 



*   CNCF is hiring! (Zach)
*   Update/revisit style guide. [Issue 7030](https://github.com/kubernetes/website/issues/7030) (Steve)
    *   Revisit this next week (2/6) to talk about whether it would be better to adapt an existing style guide (for example, Google Cloud if they have one) or continue creating from scratch
    *   Keep the existing style guide, reorganize its contents, promote its visibility better
*   Repo team cleanup (Zach)
    *   Cleanup for individual contributor permissions (Steve)
    *   Cleanup for teams: add reviewers and approvers teams (Zach)
        *   Nominating Brad Topol (@bradtopol) as a reviewer
    *   When changing OWNERS file, change OWNERS_ALIASES as well and vice versa

Agenda items



*   Approval workflow
    *   We may need to adjust our review commands from /lgtm to /lgtm+/hold
    *   Issue: [https://github.com/kubernetes/test-infra/issues/6589](https://github.com/kubernetes/test-infra/issues/6589)
    *   SIG Test update meeting at 1pm PST today
*   User journeys update (Jennifer & Andrew)
    *   Victory dance!
    *   What next? 
        *   Improve quality of linked docs
        *   Improve navigation/UX
        *   Test it!
        *   Analytics and monitoring: are we seeing improvements?

                Some good material on measuring: [Bob Watson’s articles](http://docsbydesign.com/2017/08/29/measuring-your-technical-content-part-3/) (and see linked items, plus more throughout his site)

*   Expectations and how to volunteer for reviewing duties (heckj): DEFER for one week 2/6
    *   PR: [https://github.com/kubernetes/website/pull/7133](https://github.com/kubernetes/website/pull/7133)
    *   For example, what expectations do we have & should we have of a reviewer or approver in handling PRs, issues, etc.
    *   I have been asked “how do I volunteer” informally a couple of times, with this being the topic behind the question
    *   Clearer reviewer guidelines
        *   Set an SLA for when we’ll respond to PRs: What’s a reasonable standard? 48 hours? 4 business days?
        *   Do we need to codify standards for reviewers: for example, don’t merge a PR assigned to someone else? Is our current method sustainable?
    *   As we're talking about SLO for getting attention on PRs, Paris shared this view of how we're doing today:
        *   [https://k8s.devstats.cncf.io/d/000000013/first-non-author-activity?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&from=1454786880740&to=1516735680740](https://k8s.devstats.cncf.io/d/000000013/first-non-author-activity?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&from=1454786880740&to=1516735680740)
*   Cluster-registry looking for "how/where to host" advice for their docs (heckj)
    *   @font (Ivan Font) of sig-multi-cluster requesting...
    *   Do we advise on best practices for projects within sigs?
    *   Hosting binary usage (aka like kubeadm, kubectl), API reference, etc.
    *   Current docs are markdown-in-github: [https://github.com/kubernetes/cluster-registry/blob/master/docs/development.md](https://github.com/kubernetes/cluster-registry/blob/master/docs/development.md)

**1/30/2018**

New attendees?



*   

Follow-ups from last meeting (10 minutes max): 



*   Bot automation: it’s alive!
*   Steve: Update/revisit style guide. [Issue 7030](https://github.com/kubernetes/website/issues/7030)
    *   Revisiting to make sure we recommend what we actually enforce (including usage of “we”)
    *   Add a style guide template to the [PR template](https://github.com/kubernetes/website/blob/master/.github/PULL_REQUEST_TEMPLATE.md) in kubernetes/website
    *   Revisit this next week (2/6) to talk about whether it would be better to adapt an existing style guide (for example, Google Cloud if they have one) or continue creating from scratch

Agenda items



*   Zach will be out Weds-Sunday, returning on Monday 2/5
*   What to do with repo teams/collaborators now that approvals are automated (Zach)
    *   [https://github.com/kubernetes/website/settings/collaboration](https://github.com/kubernetes/website/settings/collaboration)
    *   Zach: Add teams for reviewers/approvers (https://github.com/orgs/kubernetes/teams/sig-docs-pr-reviews/members)
    *   Steve will take on individual contributor admin cleanup
*   Expectations and how to volunteer for reviewing duties (heckj): DEFER for one week 2/6
    *   For example, what expectations do we have & should we have of a reviewer or approver in handling PRs, issues, etc.
    *   I have been asked “how do I volunteer” informally a couple of times, with this being the topic behind the question
*   User journeys update (Jennifer & Andrew)
    *   Prioritize content improvements to content inside user journey paths
*   Questions about static generation of the [community repo docs](https://kubernetes.io/docs/imported/community/guide/) (Jorge)
    *   Script is self-service: https://github.com/kubernetes/website/blob/master/update-imported-docs/community.yml
*   CNCF is hiring! (Zach)

**1/23/2018**

New attendees?



*   

Follow-ups from last meeting (10 minutes max): 



*   [https://github.com/kubernetes/website/issues/6906](https://github.com/kubernetes/website/issues/6906) [[PR 6378](https://github.com/kubernetes/test-infra/pull/6378)] (Aaron Crickenberger)
    *   We are proceeding with this!
    *   /lgtm: serves as Tech LGTM
    *   /approve: serves as Docs LGTM, will flag for automatic merging
    *   Repo maintenance teams: can we streamline?
        *   [https://github.com/kubernetes/website/settings/collaboration](https://github.com/kubernetes/website/settings/collaboration)

Agenda items



*   Prioritizing 2018
    *   In flight projects: [https://github.com/kubernetes/website/issues/6935](https://github.com/kubernetes/website/issues/6935)
*   Expectations and how to “volunteer” for reviewing duties (heckj)
    *   For example, what expectations do we have & should we have of a reviewer or approver in handling PRs, issues, etc.
    *   I have been asked “how do I volunteer” informally a couple of times, with this being the topic behind the question
*   Steve: Update/revisit style guide. [Issue 7030](https://github.com/kubernetes/website/issues/7030)

Follow-ups for 1/30: 



*   User journeys update (Jennifer & Andrew)
*   Automation meeting: Need to set an hour-long meeting for early Feb to discuss doc automation needs. (Zach)  (interested folks: Andrew, Steve, Chris, Jessica, Jennifer, Zach)

**1/16/2018**

New attendees?



*   None :-(

Followups from last meeting: 



*   SIG-Docs summit 2018.  (Zach)
    *   Portland, May 9-10 (immediately following Write the Docs)
    *   Working with LF event team to start logistics in motion
        *   Jaredb: We can host in the Google Portland office if that helps. Or, for something more fun, I recommend picking a McMennamin’s property like Edgefield :)
*   Community site followup (Jessica) 
    *   Jorge is taking the lead on this ([See netlify](https://deploy-preview-6863--kubernetes-io-master-staging.netlify.com/docs/imported/community/keps/))
    *   Import is live
*   [https://github.com/kubernetes/website/issues/6906](https://github.com/kubernetes/website/issues/6906) (Aaron Crickenberger)
    *   Need to up/down the proposal
    *   Jennifer: need to have approver/reviewer roles well defined and in place
        *   Community discussion about normalizing expectations across project
        *   SIG Docs could benefit from rigor of roles defined in kubernetes/kubernetes
    *   Steve: Do we copy what they do in kubernetes/kubernetes?
        *   /lgtm & /approve
        *   /lgtm can mean tech approval
        *   /approve can mean doc approval
    *   Zach: Take an action item for better PR review guidelines, specifically: clean commit history in PRs
    *   End result: +1, yes to Aaron to implement bot proposal

Agenda items?



*   Steve: Removal of getting started guides that result in a non-conformant cluster.
    *   Question from Ilya Dmitrichenko to sig-docs, sig-cluster-lifecycle: \
 \
Over time, a number of Getting Started Guides had been contributed to Kubernetes website. General quality and levels of maintenance of all of these guides vary. Some achieve similarly end goals with different tools. At sig-cluster-lifecycle, we believe that the best way to ensure all documentation related to cluster provisioning and bootstrap that is published on the website results in a conformant cluster configuration. \
 \
We would like to enforce conformance rules on getting started guides, and remove any guides that result in non-conformant cluster configuration by end of March (i.e. after 1.10 release). \
 \
Before we go ahead and implement this, we would like to hear from sig-docs folks and any maintainers or users of getting started guides.
    *   SIG-Docs: All in favor
        *   SIG Cluster Lifecycle will open a PR of GSGs to remove (Ilya)
        *   Reach out to owners for state of guide conformance?
        *   Add conformance/GSG removal to in-flight guide (Zach)
*   Better review habits (Zach)
    *   Get reviews before merging
    *   Reject PRs for static branches (currently anything prior to 1.9)
    *   Review for quality, not just test-passing
    *   (Steve) Establish /lgtm as tech approval, /approve as doc approval?
    *   Mature style guide (Zach, add to in-flight guide)
        *   What does it mean when you /approve?
    *   How to deal with editorial debt:
        *   Open an issue against editorial need
        *   Add to a dedicated project?
*   User journeys update (Jennifer)
    *   Deferred one week (1/23)
*   Writing mentorship proposal (Jennifer)
    *   Embedding writers in SIGs (Zach, add to in-flight guide)
    *   Need a list of SIGs in need:
        *   SIG cluster lifecycle (Jennifer)
        *   SIG apps (writer?)
        *   SIG community (Steve Perry)
*   Automation meeting: Need to set an hour-long meeting for early Feb to discuss doc automation needs. (Zach)  (interested folks: Andrew, Steve, Chris, Jessica, Jennifer, Zach)

**1/9/2018**

**NOTE: Zach on vacation.**

New attendees?



*   Aaron Crickenerger

Followups from last meeting: 



*   Move to the release repo and link to it from the docs repo (Jennifer)
    *   Move docs playbook for releases; content audit of current docs in release repo also needed; WIP (no PR yet, will announce to SIG when ready)
*   SIG-Docs summit 2018.  Eyeing May/June in 2018. (Zach)
*   Supported documentation versions: Need to follow up on https://github.com/kubernetes/website/pull/6442 (Andrew and Steve)
    *   [https://github.com/kubernetes/website/pull/6882](https://github.com/kubernetes/website/pull/6882)
    *   [Preview](https://deploy-preview-6882--kubernetes-io-master-staging.netlify.com/docs/home/supported-doc-versions/)
*   Community site followup (Jessica) 
    *   Jorge is taking the lead on this ([See netlify](https://deploy-preview-6863--kubernetes-io-master-staging.netlify.com/docs/imported/community/keps/))
*   1.9 Retro (Zach) 
    *   Deferred

Agenda items?



*   [https://github.com/kubernetes/website/issues/6906](https://github.com/kubernetes/website/issues/6906) (Aaron Crickenberger)
    *   Thoughts on implementation?
        *   Probably need to keep Docs LGTM/Tech LGTM separate
        *   We can probably implement the existing system for /hold, /lgtm, /approve--just need to be clear and specific about approval chain
        *   Discussion in issue; Aaron will not pull trigger on merge until we (SIG docs maintainers) signal that we’re ready
*   Request to do a doc sprint at IBM Index on 2/20, SF (Zach)
    *   Export a planning kit for doc sprints at other meetups?
*   We should define our big site/docs goals for 2018. What’s our process? I suggest OKRs/KPIs. (Jared)
    *   OKRs are OK!
    *   Set them up at the Summit in May
    *   SIG Maintainers meet separately beforehand
        *   List projects currently in flight
            *   Zach will set up contrib doc by 1/12
            *   Proposed meeting on 1/23
    *   Zach: set up some goals/OKRs
    *   Summit: talk about how to implement
    *   Possible top level objectives:
        *   Improve Kubernetes onboarding
        *   Findability
        *   One-stop-shop for broader community docs
        *   Improve (make more efficient) PR approval process
        *   Improve docs release process for quarterly k8s releases (including release notes) 
        *   Create tighter loop w/ community feedback on docs
        *   Increase the number docs contributors 
        *   ?
*   Importing docs from other repos, creating guidelines [PR #6863](https://github.com/kubernetes/website/pull/6863) [Preview](https://deploy-preview-6863--kubernetes-io-master-staging.netlify.com/docs/imported/) (Andrew)
    *   Follow up in three weeks: 1/30 (Andrew)

Next week: 



*   Reviewer conventions in docs file. Automation and names in front matter of file? (Steve & Andrew) 
*   Automation meeting: Need to set an hour-long meeting for early Feb to discuss doc automation needs. (Zach)  (interested folks: Andrew, Steve, Chris, Jessica, Jennifer, Zach)
*   We have a staleness/rot bot now (Zach) 

**1/2/2018**

**NOTE: Zach on vacation.**

New attendees?



*   Nope :(

Followups from last meeting: 



*   Future docs release meisters:
    *   1.10: Jennifer Rondeau
    *   1.11: Nick Chase
    *   Chris Short (@chris-short) wants to shadow a release
    *   Move to the release repo and link to from the docs repo
        *   Jennifer is on it :)
*   SIG-Docs summit 2018.  Eyeing May/June in 2018. Details to follow from Zach.
*   Supported documentation versions: Need to follow up on https://github.com/kubernetes/website/pull/6442 re:current+4 
    *   Andrew and Steve
*   Reviewer conventions in docs files (Steve & Andrew)
    *   Follow-up: automation and names in front matter of file?
    *   Steve: Check in on this in 2 weeks. 
    *   Andrew will follow up with tooling people to modify bots to specify category
*   Automation meeting: Need to set an hour-long meeting for early Feb to discuss doc automation needs. 
    *   Andrew, Steve, Chris, Jessica, Jennifer, Zach
*   Community site (community.k8s.io) for developer contributions (from Joe Beda). 
    *   Any Updates? 
        *   Jessica will do a follow up
    *   Any thoughts?
        *   Start with Keps. Keps.kubernetes.io
        *   ContribEx should own this. (via Paris)
        *   Onboarding should fall under the contributor site
    *   “Lightweight process”
        *   Questions of ownership for KEPs and Community
        *   Developer perception? If the dev perception is entirely launch based, we should emphasize the process outside of launches
*   SIG Docs representation on other sigs?
    *   What are the highest priority SIGs?
        *   Reach out to previous release leads and features leads - ex. Jaice, Ihor
        *   Reach out to SIG-Arch
*   1.9 Retro is happening next week
*   Chinese translation work:
    *   Researching best practices for i18n
    *   [Refactoring k8s.io for i18n](https://docs.google.com/document/d/10teK1YmLBcRVc6B6GjgJxRP1KcuGjL-KpHwSaANr4J4/edit?usp=sharing) doc
*   Upcoming Kubecons for 2018:
    *   [Europe: May 2-4, Copenhagen](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/)
    *   [China: June 25-27, Bejing](https://www.lfasiallc.com/linuxcon-containercon-cloudopen-china)
    *   [North America: Dec 11-13, Seattle](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/)
