## Dec 13, 2021 [Private Meeting]

**Bosun**: Paris

**Note taker**: Jordan

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**CNCF updates?**

**Votes**

**Topics**

* Quick votes/put your lgtm in/add emoji reaction:
    1. [https://github.com/kubernetes/community/issues/6274](https://github.com/kubernetes/community/issues/6274) - kdev email migration rename
    2. [https://github.com/kubernetes/steering/pull/231](https://github.com/kubernetes/steering/pull/231) - bosun.md modifications
    3. [https://github.com/kubernetes/steering/issues/223](https://github.com/kubernetes/steering/issues/223)  - annual report deadlines inside of issue body
* Quick discussion/actions needed?
    * Charter update for cocc [https://github.com/kubernetes/community/pull/6247](https://github.com/kubernetes/community/pull/6247)
    * [https://github.com/kubernetes/steering/pull/224](https://github.com/kubernetes/steering/pull/224) - steering / cocc conflict of interest resolution
        * **AI: We need a sync with cocc and someone to take an action there **
* Deep dive into [sustainability plan](https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit#heading=h.koj07ciun0s8) [15 mins]
    * [Gap analysis ](https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit#heading=h.sv6elf11kc6h)
    * paris: **looking for liaisons to help gather this info; need executive summary, very concrete asks (in terms of specific staffing/budget/projects), very concrete reasons (risks, features slipping, etc)**
    * cblecker: would this make sense to use a question in the annual report to gather?
    * dims: we have sufficient information to make a request to GB, not sure adding more details/info is necessary
    * cblecker: would it make sense to pick "pilot" sigs/areas to improve staffing/support?
    * tpepper: ties to KubeCon keynote on bringing employees to key functional areas in projects, TOC/GB reps are people who can influence those business choices
    * stephen: +1 on gathering info as part of annual reports, +1 on prioritizing most urgent sig needs
    * paris: sort of a chicken/egg issue of budget/staffing and concrete plans to address issues;
    * dims: it's going to be hard to try to address all needs simultaneously; agree with prioritizing and doing a trial with top needs
    * Paris: what would we define as a “Kubernetes Fellow” role?
        * Dims can help draft stuff for k8s/fellow.
* Annual reports [10 minutes]
    * Any updates? Past items or current
    * Quick question review
        * inline in [annual report doc](https://github.com/kubernetes/community/blob/master/committee-steering/governance/annual-reports.md#2021-and-on)
    * Roles for annual reports:
        * Prep - ensure [questions](https://github.com/kubernetes/community/blob/master/committee-steering/governance/annual-reports.md#questions-for-report) and process is up to date
            * Jordan leading this, targeting complete/ready for review 12/16 (process depends on generator, generator depends on questions)
            * process cleanup at [https://github.com/kubernetes/community/pull/6295](https://github.com/kubernetes/community/pull/6295)
            * question cleanup at [https://github.com/kubernetes/community/pull/6301](https://github.com/kubernetes/community/pull/6301)
        * Tools
            * Generator for tracking issues, template md files
                * Christoph is building one, need folks who can run it later (depends on questions)
                * Include/generate links to devstats queries with year/sig plugged in in generated template
            * Devstats/queries
                * Accuracy is a concern especially around file-paths for PRs
                * ​​Add wg/sig label filter to Contributions chart: [https://github.com/cncf/devstats/issues/299](https://github.com/cncf/devstats/issues/299)
        * Comms - announce to dev@kubernetes.io, etc
            * Paris can help and engage upstream marketing team
            * start in January after templates are generated
        * Outreach/reminders to individual group leads - liaisons
            * start in January after templates are generated
        * Summary
            * ~March after individual group reports are submitted
    * [https://github.com/kubernetes/steering/issues/228](https://github.com/kubernetes/steering/issues/228)
        * Prioritize sub-items related to annual reports are a dependency for Jan 2
        * Incrementally PR more into liaison play book after that
* Owners files and promoting people [10 mins]
        * [community-membership.md](https://git.k8s.io/community/community-membership.md) follow-up from your chair asks
            * Are they using it to promote folks? Do they have their own?
            * dims: ask in annual report
            * bob: these are general guidelines; push on groups that have additional processes/requirements to document those and how new folks can engage and get promoted into those roles
            * [Stephen] to add color with SIG Release/Release Engineering examples
                * help make sure automation we set up plays nicely with the types of subproject or subpackage memberships sig-release sets up for things like changelog reviewers
        * [Funding] Do we need a contractor on owners file work? Scripts to help ID folks to promote are blocked by this work and many other initiatives
            * [Invalid owners files in sigs.yaml · Issue #4125 · kubernetes/community · GitHub](https://github.com/kubernetes/community/issues/4125)
            * [Subprojects in sigs.yaml will be instantly out of date · Issue #1913 · kubernetes/community · GitHub](https://github.com/kubernetes/community/issues/1913)
            * dims: first instinct is to use as an LFX mentorship project, will need iteration which doesn't lend itself to a crisp scope-of-work for a contractor
            * stephen: have been using the tool to cull inactive folks, works well for that; less sure about using it for promotion/nomination
            * dims: cleanup is higher priority (avoid routing issues/PRs to inactive folks), easier to do since we have signal for activity, makes it more obvious where we need folks
            * bob: more than just adding/removing reviewers; for example, reverse mappings of packages to sigs/subprojects is needed
* Chairs and TLs policy and procedure [10 mins]
        * We need to see PRs for the following, who wants to take a go at them? PRs will help us with further discussion:
            * Selection process: [https://github.com/kubernetes/community/issues/5855](https://github.com/kubernetes/community/issues/5855)
            * Terms and Term Limits for CHAIRS (not TL): [https://github.com/kubernetes/community/issues/5886](https://github.com/kubernetes/community/issues/5886)
            * Should TLs be mandatory/explicit: [https://github.com/kubernetes/community/issues/5890](https://github.com/kubernetes/community/issues/5890)
                * [Bob] +1 for explicit. The roles are frequently conflated with each other and they need to be treated as separate responsibilities.
                * [Christoph] Do we have sufficient community consensus?
                    * recollection was there were people in favor, people opposed, and people unsure on all of these topics.
                * Is this a purely mechanical change, or is it proposing that different people must hold the two roles?
                    * Jordan: my understanding was it was a mechanical change to make explicit the responsibilities in chair and TL roles, and who was fulfilling those responsibilities (even if it's the same person)
                    * Bob: there was a lot of confusion about the chair/TL split, which reinforces my perspective that they should be split and clarified
                    * Dims: think we should start bottom up with defining membership, who are a SIG’s decision makers?
                    * Paris: think we can do explicit TL/chair in parallel
                    * Stephen: +1 on explicit TL/chair, and also asking if there are other roles not listed in common governance docs
                    * Tim: can Steering define baseline operations of governance for all SIG/WGs, and delegate/defer finer grained operational definitions to each SIG/WG as they see fit?
                    * Jordan: common ask from TL/chairs was “what problem is this solving”?


## Dec 6, 2021 [Public Meeting] (recording)

**Bosun**: Paris

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables) - will be late
* Christoph Blecker (@cblecker)
* ~~Davanum Srinivas (@dims)~~
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Josh Berkus
* Arnaud Meukam
* Ihor Dvoretskyi
* Kendall Nelson
* Sergey Kanzhelev

**CNCF updates?**

**Votes**

*  [1 min] DONE - Liaisons for committees: [https://github.com/kubernetes/community/pull/6260](https://github.com/kubernetes/community/pull/6260)

**Topics**

* Is this done? What else needs to be done here? [1 mins]
    * Artifact Hub issue: [https://github.com/kubernetes/steering/issues/188](https://github.com/kubernetes/steering/issues/188)
        * AI for bob this week
    * Chair responsiveness:[https://github.com/kubernetes/community/issues/4289](https://github.com/kubernetes/community/issues/4289)
        * Stephen: annual reports were not a thing at the time of this, we also have the leads meeting
            * Bob: this was also pre-steering-liaison as well, which has helped some with community group communication
            * Bob: to go through the existing comments to affirm none remain unaddressed (if so shift to separate issue?)
* Needs an owner/assignment [1 min]
    * to run down with CNCF for contributor travel funding: [Issue](https://github.com/kubernetes/steering/issues/109); [last comment](https://github.com/kubernetes/steering/issues/109#issuecomment-984905164)
        * Stephen
        * Ihor is going to check with CNCF on their current take
    * [https://github.com/kubernetes/community/issues/1913](https://github.com/kubernetes/community/issues/1913) - sigs yaml / subproject stuff
        * Jordan
    * To get a quote and then start funding process for lead training: [https://github.com/kubernetes/community/issues/5913](https://github.com/kubernetes/community/issues/5913)
        * Tim and paris will pair in 2022
* [Sergey Kanzhelev] Dockershim removal user feedback, comms, and sponsoring work to support migration off dockershim ([SIG Node discussion 11/30/2021](https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit#bookmark=id.r77y11bgzid)). [2 mins]
    * [Bob] For initial blog post, message should come from sig-arch and sig-node. This is more of a technical/architectural decision and they would be the right groups to talk about the issue.
    * [christoph] we support sig release, sig arch, and sig node for making the decisions needed; for comms - steering doesn’t take on technical decisions
    * [paris] lets work with the end user community too and make sure this is on their agenda
    * [Stephen] Request from Celeste + other coordinators for SIG Release to assist w/ messaging as well

        Notes:

1. Dockershim deprecation announcement was made a year ago
2. Migration numbers are not ideal based on for example DataDog survey, ie: ~10% have migrated off dockershim, ~90% remain on it
3. Survey was conducted and shown:

        **_TL;DR; The number of people who don't feel prepared is over 40%. The most common reason for not migrating is a dependency on docker, in most cases those dependencies are from both - self-authored and third-party tools. 40% of respondents who do NOT feel ready for dockershim removal planned to adopt 1.24 in 2022. There are a list of tasks I compiled from free text comments. The need to re-educate personnel, change processes and tooling, without the direct business benefit (basically the cost of operation) is a big complaint that we also need to address as a community._**


            [https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#heading=h.wzgwyg229djr](https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#heading=h.wzgwyg229djr)

4. Some comments from the survey:

        Many comments about specific scenarios that doesn't have any docs coverage


        _To me, this is a lot of work for a change that doesn't seem to offer much benefit. Kubernetes keeps changing, and staying current is almost a full time job. The frequent upgrade cycle seems to be for the benefit of the Kubernetes engineers, not the end users._

5. SIG Node members spoke up ([SIG Node discussion 11/30/2021](https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit#bookmark=id.r77y11bgzid)):
    1. End users are getting enough time. There is an extra full year (till December 2022) while kubernetes 1.23 (with dockershim) will be supported by upstream. And likely longer for managed platforms.
    2. The lack of CRI support by third party tools vendors is a chicken-and-egg problem. Vendors are not migrating as end users keep using dockershim. And end users keep using dockershim because of dependency to these vendors. And the only way to resolve the chicken-and-egg problem in this case is to set a strict deadline.
    3. Deadline extension will keep draining contributors' energy and unlikely will change the situation drastically in terms of adoption numbers.
    4. Some problems listed by end users in the feedback form comments section can be addressed by documentation.
6. Things that look bad:
    5. Vendors and third party tool authors are very slow to react for this change. Effectively leaving end users with no viable alternative till very recently.
    6. Many documentation tasks for end users, especially self-hosted, are still pending.
    7. CRI was just promoted to v1 (from v1alpha2) in 1.23.
    8. Kubeadm didn’t even update docs to indicate the runtime deprecation. Still uses Docker as a default. From the end users perspective using Kubeadm, it is hard to find out that it was deprecated. There is no migration story as well.
    9. Promised backup [story](https://www.mirantis.com/blog/mirantis-to-take-over-support-of-kubernetes-dockershim-2/) from Mirantis is in limbo now.

        Embracing the CNCF value #1:


            [(a) Fast is better than slow. The foundation enables projects to progress at high velocity to support aggressive adoption by users.](https://github.com/cncf/foundation/blob/master/charter.md#:~:text=(a)%20fast%20is%20better%20than%20slow.%20the%20foundation%20enables%20projects%20to%20progress%20at%20high%20velocity%20to%20support%20aggressive%20adoption%20by%20users.)

7. We had an initial session with Celeste Horgan, Bob Killen and others on comms/tasks we may need to complete.
8. Asks:
    10. Secure commitment from CNCF to work on the migration tasks and provide educational materials for end users. Ideally announce this commitment as steps to mitigate the user feedback.
                1. Bob: CNCF has committed, docs folks lined up
                2. Stephen: also, sig-release can track delivery of these docs/artifacts as a prereq for release of 1.24 when dockershim is removed
    11. Before dockershim will be removed, publish the blog post from Steering, indicating acknowledgement of the user survey results, possibly commitment for docs updates and support, and reasoning why we proceed (empathy blog post).
                3. stephen: comms should come from sig-arch/node (with some input from sig-release on the release specifics)
                4. Christoph: agrees the technical authority is delegated here to the SIGs, is not Steering’s place to do such a post.  Do agree though with the technical decisions made
    12. Consider extending EOL time for 1.23 as the last version with dockershim.
                5. stephen: minor version release lifecycle [already includes a ~2 month maintenance period](https://kubernetes.io/releases/patch-releases/#support-period) beyond 1 year to accommodate situations like this
    13. Help define, document, and get necessary commitments for this issue: [https://github.com/kubernetes/community/issues/5344](https://github.com/kubernetes/community/issues/5344) Some ideas:
        1. Get CNCF help talking to vendors, projects, and end users.
        2. Deprecation timeline of 1 year, given many third party tools dependencies, looks aggressive. Need to make sure there is a clear message about what vendors and end users need to accomplish in that year.
* [Election Retro](https://docs.google.com/document/d/1edqfsSNk_p746PcXlbHTJa8ZHNlnw3QEqNPzlVDJmUc/edit#) Item: [10 mins]
    * [https://github.com/kubernetes/community/issues/5092](https://github.com/kubernetes/community/issues/5092) - Refine rules on exceeding maximal representation for SC election
        *
    * [SC Candidate criteria should be revisited](https://github.com/kubernetes/steering/issues/227). At this time there are no restrictions on who can run. Should there be a minimum requirement such as org member? Valid voter etc?
        * Stephen: feel like SC eligibility should include being an eligible voter
        * Christoph: current requirements were intentionally set up as open by the bootstrap committee (not that they couldn't be changed), but the reasoning was that restricting who can vote was a clearer line to draw, and was sufficient, while leaving the candidacy aspect open (to allow the community to draw leaders from outside the community if desired)
        * Bob: gut reaction is candidates should be an org member (ensures they are on mailing lists, etc). If we don't do that, and keep current stance, we should at document why the current requirement is the way it is (since the question has come up multiple times)
        * AI for steering members: comment/vote on issue
    * Comments from retro:
        * Consider who gets excluded, are there Steering level stakeholders who would not be GitHub active (eg: user group folks?)
        * Beware the trap of SC getting to decide who can and cannot be on SC
    * [Automatically add anyone in Sigs.yaml to voters?](https://github.com/kubernetes/steering/issues/226)
        * groups that are active but work in private are frequent sources of voter exception requests
        * Stephen: in favor of private committees, in favor of sigs.yaml IF we ensure people listed there are current (some user groups, etc, might )
        * AI for steering members: comment on issue
    * Finally: ContribEx forming Elections subproject
* Annual Reports [15 mins]
    * Wrap up from last time [https://github.com/kubernetes/steering/issues/207](https://github.com/kubernetes/steering/issues/207)
        * Need to set an issue for the wg reporting into sponsoring sigs
            * Stephen: sponsorship seems not a thing otherwise
            * to create issue
            * Jordan: Should they tag/email their sponsoring SIG to review their annual report? (annual cadence isn't ideal, but is better than current?)
        * Generator update: [https://github.com/kubernetes/community/pull/5514](https://github.com/kubernetes/community/pull/5514)
            * Christoph can put the time in the next month; contribex mentees are looking into this in the meantime
        * Jordan and Stephen update on automatable questions ([ref](https://github.com/kubernetes/steering/issues/207#issuecomment-984871009))
            * Jordan: PR to update annual report template to prioritize questions we know will need human answers, and mark questions that can/should be automatable
            * Tim: unsure how much can be automated in all cases for some of these  (positive case is easy to verify, negative case requires human assessment, e.g. meeting was cancelled so recording wasn't uploaded)
            *
        * [Bob] Need a tool to crawl shared repos (e.g. k/k, k/enhancements, etc) to update subproject OWNERS files + feed to devstats for accurate reporting.
            * to set issue
        *
    * [Dates and Deadlines](https://github.com/kubernetes/steering/issues/223)
        * Questions complete: target 1/15/22, stretch goal 1/1/22
        * Community group drafts ready for review: discuss
        * Community group drafts merged: discuss
        * Annual report merged: maybe 3/31/22?
    * Who is doing what
        * Prep
            * are there ways we can tweak questions so they are easier to drop directly into the summary?
            * Google Form? Add structure to the ask and recording of answer
        * Outreach
            *
        * Summary
            * this was by far the hardest part last year
            * definitely need to find ways to distribute this
                * liaisons pull in content from their groups
                * other ways to shard?
* [Sustainability Plan](https://docs.google.com/document/u/0/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit) [15 mins]
    * Needs analysis
        * [Paris] Liaisons should work w/ groups to determine needs
    * What we need that still doesn’t have a solution
        * “Is it meeting the annual report?” / do we have a baseline to determine what's good
        * are maintainers spending their time/energy training new folks, putting out fires, keeping the lights on, improving/developing scalable approaches to maintaining their components, new work, support of existing work?
* Topics from the community
    * [your name here]


## Nov 15, 2021 [Private Meeting]

* Bosun: Dims
* Note taker: Jordan
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Paris Pittman (@parispittman)
    * Stephen Augustus (@justaugustus)
    * Tim Pepper (@tpepper)
* Agenda
    * Promotion process for sigs
        * Are folks using to promote people (e.g. to Reviewer, Approver, Subproject owner)?
        * Possible sig specific notes on what it would take for a person to be promoted as reviewer/approver/chair (sig-node was doing something?)
        * Chartering involves calling out where your area differs from project norm, adding linkage to area-specific membership processes or highlighting in the doc where they differ from  [community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md) would clarify
        * Bob: mechanical suggestion is to include link to description of requirements in OWNERS file as a comment
        * Tim: "With no objections from other subproject owners" language gives pause… worries can block healthy change and turnover, risks opaque black-balling of a candidate
        * AI: review [community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md), consider updates on if/how a candidate might be blocked
        * Christoph: would area-specific changes to reviewer/approver requirements require steering review?
            * Dims: envision a gradual process (ask in annual report process, work with individual groups currently facing difficulties here, eventually end up with language or links in charter)
            * Bob: want to find a balance between making steering a bottleneck for updates while giving a recourse to someone encountering requirements that don't seem possible to meet
        * Paris: suggest liaisons ask their community groups to get a baseline of current state:
            * Question 1 - do you deviate from this community membership doc
            * Question 2 - do you have it written down
            * Question 3 - can you put that written down thing somewhere public
    * SIG / WG Liaisons
        * [https://github.com/kubernetes/community/blob/master/liaisons.md](https://github.com/kubernetes/community/blob/master/liaisons.md)
        * P0: Transitioning groups from outgoing steering members to incoming members
        * P1: Also consider any rebalancing that might be helpful
        * AI: start doc to transition liaisons
        * Add liaisons for committees?
            * General agreement that having someone paying attention to needs from committees is good
            * Especially useful for committees that work in private (CoCC, PSC, etc) and easy to lose track of how they are doing
    * What’s important to us this year?
        * We used to do this in person when things were non-pandemic state, getting to know each others' priorities is important to this crew
        * AI: Paris to schedule quality time w/ Steering members :)
    * Annual Reports
        * [https://github.com/kubernetes/steering/issues/207](https://github.com/kubernetes/steering/issues/207) - things we need to do before next reporting cycle
            * theme: less toil for report authors
            * Jordan+Stephen: look at last year’s reports, specifically note the things for which there was not a tool and that meant human toil, give concrete ask instead of broad help to devstats dev
            *
        * [https://github.com/kubernetes/steering/issues/223](https://github.com/kubernetes/steering/issues/223) - dates and deadlines for next reporting cycle
            *
    * Bosun : Paris to run the next 2 steering meetings
        * Private December meeting is currently 12/20, will we have quorum or do we need to move earlier?


## Nov 8, 2021 [Public Meeting]

* Bosun: Dims
* Note taker: Jordan
* Steering Attendees:
    * ~~Bob Killen (@mrbobbytables)~~
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
    * Stephen Augustus (@justaugustus)
    * Tim Pepper (@tpepper)
* Topics
    * [election officers] 2021 Steering Committee Election Results
        * Re-elected Members
            * Paris Pittman
            * Christoph Blecker
        * Newly elected Members
            * Stephen Augustus
            * Tim Pepper
        * Emeritus Members
            * Nikhita Raghunath
            * Derek Carr
    * Our thanks to the election officers : Alison Dowdney, Josh Berkus, Noah Kantrowitz
    * Thanks to all who ran
    * Election retro scheduled? Initial election officer observations
        * election mechanics went smoothly
        * most candidates for any election
        * downward trend on absolute numbers of voters
        * AI: date for retro
    * Transition planning & [onboarding](https://github.com/kubernetes/steering/blob/main/onboarding.md)
        * AI: [issue](https://github.com/kubernetes/steering/issues/219) copying [onboarding checklist](https://github.com/kubernetes/steering/blob/main/onboarding.md) (by tradition assigned to incoming members)
        * handoff of owned issues from derek/nikhita
        * handoff of community group liaisons from derek/nikhita
        * Tim Code of Conduct committee transition
            * (precedent with Paris vacating CoCC seat when joining steering)
            * AI: formalize/document Steering/CoCC mutual exclusivity
            * AI: promote runner-up from last CoCC election? Paris to look up results
        * AI: set date for steering context handoff
            * tracked in [https://github.com/kubernetes/steering/issues/219](https://github.com/kubernetes/steering/issues/219)


## Oct 18, 2021 [Private Meeting]

* Bosun: Davanum Srinivas
* Note taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * ~~Jordan Liggitt (@liggitt)~~
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
* Topics
    * Kubecon debrief
        * Positive feedback, encourage corporate support (cblecker)
        * Wish they had heard this earlier (paris)
    * GB Meeting debrief (wednesday / paris)
        * Need to be more clear with asks
        * SIG breakdown(s), what’s missing? Help middle managers make a case to get full time folks
        * sig-cloud provider needs help too
        * OpenStack attempted this once before with mixed results. They were good for single one off tasks, but not for maintainers - people that stick around.
        * There was a big push for maintainers, not contributors - sustained contributions.
        * Need more senior contributors, not just jrs that get trained up and eventually leave the project.
    * Election progress
        * All bios are in
        * There are 2 slots available for steering candidates to ask questions of steering, and what its like to be on steering.
    * Openstack PTG meeting this Friday
    * [Issue Triage](https://github.com/kubernetes/steering/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc) if time permits in k/steering
    * [Issue Triage](https://github.com/kubernetes/community/labels/committee%2Fsteering) if time permits in k/community
        * Maximal Representation issue (no changes)
        * Write down rules for selecting GB rep and their lifecycle (open issue, thanks Paris!)
    * Thanks a ton to both Derek and Nikhita!
    * Would be good to have a retro. (when we bring in the new crew) also do a board walk for the first meeting with new crew.

## Oct 4, 2021 [Public Meeting]

* Bosun: @parispittman
* Note taker: @liggitt
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * ~~Derek Carr (@derekwaynecarr)~~
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
* Topics
    * [Kendall] OpenStack Technical Committee Invite to PTG
        * Dates/Times Available:
            * Monday, Oct 18 15 UTC  - 17 UTC
            * Thursday, Oct 21, 13 UTC  - 17 UTC
            * Friday, Oct 22, 13 UTC  - 17 UTC - Dims +1, Bob +1, Paris +1
        * [Free Registration](https://www.eventbrite.com/e/project-teams-gathering-october-2021-tickets-161235669227)
        * AI: reach out to kendall with specific time in the above slots that works for most steering members to sync up (16-17UTC on 10/22 looks best currently)
    * [paris] Working on sustainable staffing ideas doc
        * [Please add yours and comment to the doc](https://docs.google.com/document/d/1qeoP6i7GBTVJuJE1AGY5iU9dqmAOxrjqkfNQ2-rBeyI/edit?usp=sharing )
        * Part of Paris' work on special projects committee of GB
        * k8s-infra, testing, docs, contribx staffing needs identified
        * AI: steering liaisons sync with your sigs/wgs this week to get feedback and ask for specific needs and include them in the doc
            * Bob Killen (@mrbobbytables)
            * Christoph Blecker (@cblecker)
            * Derek Carr (@derekwaynecarr)
            * Davanum Srinivas (@dims)
            * Jordan Liggitt (@liggitt) - done 2021-10-08
            * Nikhita Raghunath (@nikhita)
    * [https://github.com/kubernetes/steering/issues/215](https://github.com/kubernetes/steering/issues/215) - this is important; wg-> sig transition docs
        * We also need clarity on turning down a WG (how many steering members need to weigh in?)
        * [https://github.com/kubernetes/steering/issues/214](https://github.com/kubernetes/steering/issues/214)
        * Specific questions:
            * what is the process of converting a wg to a sig?
                * cblecker: review/comment timeframe? does the proposal need to be brought to a public meeting first? who weighs in? what's the minimum comment period?
                * cblecker: explicit ack vs lazy consensus
                * cblecker: establish timeframes so the proposer and the commenters understand expectations to get / provide feedback
                * Jordan: sponsorings sigs need to ack on the conversion and there needs to be a clear assignment of each aspect of wg artifacts/work to a sig (either one of the wg's sponsoring sigs or the proposed new sig)
                * Dims: There should be discussion around creation/deletion in a public meeting
                * cblecker: we need to provide better visibility to public steering agendas ahead of time so people know when something will be discussed and what the right forum is to provide feedback
                    * dims: agree, though there's a tension around wanting to work async to avoid blocking/waiting for meetings, but balanced with giving time and allowing for interaction during meetings on topics
            * how many steering members required (majority >= ½, or super majority, >= ⅔)?
                * steering voting requirement for wg creation is a [simple majority](https://github.com/kubernetes/community/blob/master/sig-wg-lifecycle.md#prerequisites-for-a-wg) (e.g. 4/7)
                * steering voting requirement for sig creation is unclear?
                    * jordan: wg → sig conversion should match that (once clarified)
                    * dims: gut feel is super majority (e.g. 5/7) (+1 paris, +1 jordan)
                    * cblecker: don't feel strongly about simple vs super majority for sig creation; think the timelines, feedback processes are more important
                    * dims: even if it is done async, need to be clear on time allowed for feedback, number of steering members input
                    * cblecker: should it be possible to form or promote a wg→ sig without it appearing in a public meeting?
                    * dims: in the spirit of lettings sigs make decisions and self-organize, should a meeting be required?
                    * jordan: establishing sigs is rare and ~one of the most important things steering does… doing async work ahead of time in threads/docs/issues/comments is good, but if anything deserved discussion in a well-publicized synchronous meeting, this seems like it does
            * AI: cblecker: propose PR update to sig establishment/wg promotion process
    * Speaking of, any comments for wg-naming and wg-component-standard? Both are turning down / done sunsetting
        * dims: many thanks to leads of both groups
        * contribx channel/list turndowns in progress
        * some owner file cleanups remaining, acked by wg leads
    * election status
        * nominations open
    * Note: – quorum lost, bob had to drop –
    * Bosun and other roles of steering committee
        * paris: following stated guidelines on bosuning meetings ([https://github.com/kubernetes/steering/blob/main/bosun.md](https://github.com/kubernetes/steering/blob/main/bosun.md)) important to ensure distribution of work and engagement of community
            * dims +1 (also, dims will be bosun for november)
            * liggitt +1
        * paris: think we need to rotate steering work (issue wrangler, pr wrangler, communications lead) to ensure SLOs on work in progress, clarify who is responsible for moving work forward
            * jordan: in favor of clarifying responsibilities and establishing a rotation for the year so all seven of us are not chasing the ball all at once (or all seven of us assuming one of the other six is). strawman: ~4 areas, set up monthly rotation for the year once new steering members are in place, if you can't make a rotation swap with someone
            *  Order a bosun whistle?: *

## Sep 20, 2021 [Private Meeting]

* Bosun: Bob Killen
* Note taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * ~~Christoph Blecker (@cblecker)~~
    * ~~Derek Carr (@derekwaynecarr)~~
    * ~~Davanum Srinivas (@dims)~~
    * ~~Jordan Liggitt (@liggitt)~~
    * ~~Nikhita Raghunath (@nikhita)~~
    * ~~Paris Pittman (@parispittman)~~
* Topics

## Sep 13, 2021 [Public Meeting]

* Bosun: Bob Killen
* Note taker: Christoph Blecker
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Community attendees:
    * Alison Dowdney
    * Arnaud Meukam
    * Ihor Dvoretskyi
    * Lubomir Ivanov
    * Kendall Nelson
* Topics
    * [Bob] Default branch rename of steering owned repositories (k/steering, k/funding)
        * AI: Bob will take care of it later today
    * [paris] Do we need extensions re: election?
        * election committee set? ([thread](https://groups.google.com/a/kubernetes.io/g/steering/c/nB1K0XZeYRQ/m/DXf9uj-4CAAJ), [PR](https://github.com/kubernetes/community/pull/5986))
        * [Update community repo for the 2021 elections](https://github.com/kubernetes/community/pull/5986)
            * AI for steering, please review!
        * AI: Alison to talk to other election committee members.
        * Follow up item, kubecon in Nov may be a permanent change
    * [paris] Started asking chairs for sustainability suggestions including areas that need staffing/contract support
        * Paris will send a draft doc to chairs for comment; goal of presenting to GB and having a plan for testing, docs, contribex and other support
        * Current convos around staffing json-iterator and go-yaml; lucas brought up that this is a problem across cncf projects
            * Jordan: stdlib/third-party serialization libraries specifically are ~all in a questionable state (gogo-protobuf, yaml, json-iterator). sig-api-machinery owns most of the code around these, sig-architecture is probably the right place for a policy around stdlib implementations vs third-party implementations vs Kubernetes forks of third-party implementations. AI: add to sig-architecture / sig-api-machinery meeting agendas

## Aug 16, 2021 [Private Meeting]

* Bosun: Christoph Blecker
* Note taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * ~~Davanum Srinivas (@dims)~~
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * CoCC issue
    * Cocc election - please vote
        * Pop quiz - do we need a majority?
    * Conversion of wg-k8s-infra to a SIG
        * [https://github.com/kubernetes/community/pull/5928](https://github.com/kubernetes/community/pull/5928)
    * Follow up on AIs from last meeting
    * I think we are ready to merge!! [https://github.com/kubernetes/community/pull/5736](https://github.com/kubernetes/community/pull/5736)
        * All feedback from previous meetings and comments addressed

## Aug 2, 2021 [Public Meeting]

* Bosun: Christoph Blecker
* Note taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * ~~Paris Pittman (@parispittman) ~~
* Community attendees:

* Topics
    * WG Naming Dissolution request
        * [https://groups.google.com/a/kubernetes.io/g/steering/c/8fy8omuKdpM](https://groups.google.com/a/kubernetes.io/g/steering/c/8fy8omuKdpM)
        * Bob had talked to Chairs (Celeste/Stephen) on moving things out (dangling artifacts, maybe some of them can stay where it is now)
            * Evaluation framework docs
                * AI for  (link to from API review/conventions docs)
            * Who drives evolution of the framework in the future? (maybe we don't imagine the framework evolving much… may not need active owners?)
            * Who owns work identified? Looks like component owner sigs, sig-release/k8s-infra/contribex etc own the actions (make sure there are issues open/tracking)
            * AI for Bob to talk to the Chairs again and make sure AI’s are followed up.
    * "kubernetes and kubernetes-sigs org membership should be equivalent"
        * [https://github.com/kubernetes/org/issues/966](https://github.com/kubernetes/org/issues/966)
        * Now that we don’t have incubator and have a handle on the stale folks (and we prune folks), seems like a good idea now. (we pruned 400 people in the recent past).
        * [https://github.com/kubernetes/org/issues/966#issuecomment-822045622](https://github.com/kubernetes/org/issues/966#issuecomment-822045622) has stats on "in kubernetes-sigs, not in kubernetes" org
        * 5 steering members present today all voted in favor of making the membership equivalent. The ball is back in sig-contribex court.
    * Can someone help drive this stuff?
        * Annual Reports Actions and Wrap Up
            * [https://github.com/kubernetes/steering/issues/207](https://github.com/kubernetes/steering/issues/207)
            * Also set issues for each theme and assign someone - the membership theme has one already
            * Individual tasks needs owners to drive.
            * Christoph to help shepherd the AI(s) / tasks.
        * Should TechLeads be mandatory?
            * [https://github.com/kubernetes/community/issues/5890](https://github.com/kubernetes/community/issues/5890)
            * Explicitly list out who is the TL for each SIG
            * Get feedback on changing Governance doc first! (email to chairs/leads and discussion in their mailing list)
            * AI: for Bob to gather feedback.
        * Terms and Term Limits for Chairs
            * [https://github.com/kubernetes/community/issues/5886](https://github.com/kubernetes/community/issues/5886)
            * AI: Same as above, we need to gather feedback as well. This is assigned to Paris already. Jordan has volunteered to help with this as well, will prep request for feedback for chairs/leads meeting next week.
            * Elana has requested to split the “Term” vs “Term Limit” separate. +1 from various folks, but will take discussion to issue.
            * Derek: how do minimum requirements for number of folks in leadership roles interact with terms/term limits? Ideally we want overlap to get successors up to speed, etc. Will add comments to issue.
            * Derek: did openstack add terms/term limits? Dims: they do have voting on a periodic basis, unsure about term length or limits.
            * Project energy needs to be balanced with the frequency of elections as it takes up time/effort.
    * This is an FYI only that a funding request is going to come through: [https://github.com/kubernetes/community/issues/5913](https://github.com/kubernetes/community/issues/5913)
        * Leadership training options!
        * Yay!
    * GB meeting end result: we should meet with every GB member individually to see what they can help out with from their org (time, money, etc). We asked for more full time engineers and headcount for testing, infra, docs, community.
        * Thanks a ton for doing this Paris!
    * GB meeting kubecon discussion summary (dims)
        * mostly focused on safety, requirements, plans for venue
        * event is a go at this point unless venue or governmental entities intervene
        * oscon is running before kubecon, intended to have same vaccine requirements, dry-run for kubecon procedures
        * bob: also factors into contributor summit planning
        * dims: another smaller event around the same time [https://events.linuxfoundation.org/lf-member-summit/](https://events.linuxfoundation.org/lf-member-summit/)
    * WG k8s infra convert to SIG
        * [https://groups.google.com/a/kubernetes.io/g/steering/c/udk1gQn0q0w/m/zf_uac2DBQAJ](https://groups.google.com/a/kubernetes.io/g/steering/c/udk1gQn0q0w/m/zf_uac2DBQAJ)
        * AI(Jordan): add comment to get review/feedback/ack from stakeholder sigs before steering [https://github.com/kubernetes/community/tree/master/wg-k8s-infra#stakeholder-sigs](https://github.com/kubernetes/community/tree/master/wg-k8s-infra#stakeholder-sigs)
        * AI: Nikhita as WG liaison to help with progress
    * SC election update
        * Bob and Josh are on point. Update in a week or so.
    * Call for community member questions/topics

## July 19, 2021 [Private Meeting]

* Bosun: Paris Pittman
* Topics:
    * Annual Reports
        * Next phase of work
            * [https://github.com/kubernetes/steering/issues/207](https://github.com/kubernetes/steering/issues/207)
        * How to handle governance issues
            * [https://github.com/kubernetes/steering/issues/212](https://github.com/kubernetes/steering/issues/212)
                * Steering member takes 1 theme’d issue
                    * Paris will look to see what themes we don’t have issues for
                * If it’s a community like or contribex owned, then contribex owns the issue
    * Steering Committee Election guide and election information
        * [https://github.com/kubernetes/steering/issues/190](https://github.com/kubernetes/steering/issues/190)
            * Contribex will recommend for EC
            * 50+ contributions a year, github member status - async vote
            * Who can run? - change policy to say must have a valid voter before they can run? no
            * Happen in the next two weeks
    * Naming Working Group
        * Wait until PR is filed
        * Sig-arch can be back stop for things that don’t have owners
        * Sig-release is going to do the switch
    * CNCF updates
        * Next GB meeting at the end of the month
        * Do we have any open ended funding matters?
        * [paris] we should revive the contributor travel / kubecon
            * Let chairs and tech leads know if folks need funding then they should contact us

## July 12, 2021 [Public Meeting]

* Bosun: Paris Pittman
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * ~~Derek Carr (@derekwaynecarr)~~
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Community attendees:
    * Aaron Crickenberger (@spiffxp)
    * Konstantin Semenov (@jhvhs)
    * Scott Andrews (@scothis)
    * Ben Hale (@nebhale)
    * Rey Lejano (@reylejano)
    * Nabarun Pal (@palnabarun)
    * Aeva Black (@AevaOnline)
    * Baiju Muthukadan (@baijum)
* Note Taker(s)
    * spiffxp
* Topics:
    * [https://github.com/kubernetes/community/issues/5855](https://github.com/kubernetes/community/issues/5855) - Better Selection of Chairs/TLs
        * Context:
            * There have been ongoing discussions around how SIG Chairs are selected. Idea of term limits came up
            * There is no standard way across the community of handling leadership changes.  Some have leadership nominate + SIG consensus, some have mini-elections
            * Are there things we should make mandatory? E.g. selection process, role lifecycle, TL role
            * We’ve had lots of healthy discussion on the issue, trying to discern if there are any conclusions that can be drawn
        * Notes
            * Seems like three things: mechanisms for selecting leads, terms/term-limits, TL role optional-or-mandatory.  Maybe don’t try to handle all three at once? Which is most important, can we drive to conclusion
            * TL optional or not.  Maybe make it explicit so that those chairs that do double duty can be listed in both places, make it clearer that there are two buckets of work.
            * Turn these three distinct things back over to the community vs. continuing to engage on attempting to talk about all the things?
            * Terms/term-limits: could start with the concept of term, helps clarify what someone is stepping into, and separate from term-limit
            * Everyone generally in favor of “terms” but want consensus from community leads
            * Selection process - we didn’t really discuss, but will need to be a part of the conversation as we look at “terms mean you regularly have the opportunity to step-down/step-up”
        * Decision/Action/Follow up:
            * Break issue into 3 prongs above [dims/paris]
                * Include data
                * Include the why we are doing this for each one / what we are trying to solve
            * Chair and tech leads to weigh in on each; timebox
            * Get on the community meeting agenda
            * Put mechanism first to decide - probably need the most feedback on this; could lead into the term discussion
    * [https://github.com/kubernetes/community/pull/5736](https://github.com/kubernetes/community/pull/5736) Roles visibility PR
        * Context:
            * Looking to establish uniformity across SIGs
            * Other roles for other SIGs
            * Additional role definitions don’t need sign off from Steering
            * Roles that peel off responsibilities from Chairs/TLs need steering approval, go through usual charter review process
            * Roles that don’t take responsibilities off of Chair/TLs, more refine/define how the SIG executes, they don’t require steering approval
            * Aligns with laissez-faire
            * Eg. PRR / API Reviewer roles don’t blow the scope of SIG Architecture, so even if they choose to organize execution of those responsibilities “differently”, that doesn’t need an extra layer of steering review from a “core governance” perspective
        * Decision/Action/Follow up:
            * Take off tech lead line - open a new PR
            * Put musts at the top
            * Kdev
    * [https://github.com/kubernetes/community/pull/5746](https://github.com/kubernetes/community/pull/5746) - Extend the scope of sig-service-catalog
        * Context:
            * The project itself has fallen into disuse, various stakeholders moved on to operator-based solutions, so it’s seen a lot less use these days
            * There is still a need for the idea of bindings, some programmatic way of attaching provisioned services to running apps.  Take the service-binding idea out of service catalog in a way that it can be used more generically
            * Considering adopting a binding spec more specific to kubernetes vs. the generic spec used by e.g. cloudfoundry, heroku.  Go beyond OSBAPI implementation, and be more k8s-native
        * Notes
            * Might make more sense for service binding to go to SIG Apps as a subproject instead?
            * Does this still need to be its own SIG?
            * The spec itself is designed to act as a bridge between service world and apps so it doesn’t really sit in either
            * Sig-apps wasn’t initially responsive re: taking on the subproject but service-catalog did ‘
            * Started having trouble whether the scope expansion would fall in SIG SC vs. SIG Apps, which led to idea that maybe homing as a subproject in SIG Apps would make more sense
            * Having a single SIG as point-of-contact for questions related to service binding and how it hooks up to apps seems more beneficial than have distinct technical boundaries fall at the SIG level vs. subproject level
            * The thing that this was enabling was off-cluster services, but the move to “in-cluster” is what causes it to have overlap with SIG Apps / SIG Network
            * Where is the maturity curve for this? The service binding spec is quite mature.  Looking for a home in kubernetes to land.  This is just another variant of how do I get credentials to an application.  So if we were to land in SIG Service Catalog, it requires an expansion of their charter.  If that doesn’t work, we definitely want to find someplace to land.
            * Is this just a spec or is there a reference impl?  Not a reference impl yet, but is broadly developed in the open.
            * Could possibly be a CNCF thing? The entire spec is built basically on kubernetes, so it seemed to make sense to start here.
            * SIG Service Catalog has been inactive in general, so seeing an expansion was a little surprising when looking at e.g. current status of meeting recordings, annual report, etc.  SIG Apps has been more active than this, so if all you’re looking for is a home, maybe that makes more sense (but should poll Apps/Network for clarity)
        * Decision/Action/Follow up:
            * Talk to sig-apps [who?] - can this be a subproject of sig-apps instead of charter change for service catalog
            * Talk to sig-network [who?] - get feedback from the sig-apps / sig-network leads to see if they see a clear distinction between their responsibilities and this proposal?
                * Set up a meeting to force a time
                * Liaison - bob, christoph, derek
            * Can this be a working group discussion
            * Should service catalog be a sig?
            * Could this be housed in cncf?
        * [https://github.com/kubernetes/steering/issues/213](https://github.com/kubernetes/steering/issues/213) - Google Workspace Automation
            * Context: mailing lists, and stuff. Gcp org kubermetes.io that hosts project infra; tied to this workspace. Has five super admin (ihor, 3 SC, +1 more); spiffxp used to be a super admin and helped dims set up mailing list management but is no longer on steering so has lost super admin privs. Lots of back and forth with access and people who have keys.
            * Notes: tldr - aaron needs super admin privs
                * Jordan asked if super admins can read private groups - bob answered: not unless they add themselves to the groups directly, in that case there is an audit trail
            * Decision/Action/Follow up:
                * Dims to take action
    * Discuss/Assign: - BUMP THESE TOPICS TO NEXT MEETING
        * Annual Reports
            * Next phase of work
                * [https://github.com/kubernetes/steering/issues/207](https://github.com/kubernetes/steering/issues/207)
            * How to handle governance issues
                * [https://github.com/kubernetes/steering/issues/212](https://github.com/kubernetes/steering/issues/212)
        * Steering Committee Election guide
            * [https://github.com/kubernetes/steering/issues/190](https://github.com/kubernetes/steering/issues/190)
    * CNCF updates
        * Next GB meeting at the end of the month
        * Do we have any open ended funding matters?
        * [paris] we should revive the contributor travel / kubecon

## June 21, 2021 [Private Meeting]

* Bosun: Nikhita/Paris
* Note Taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * **[private]** CoCC issue

## June 7, 2021 [Public Meeting]

* Bosun: Nikhita
* Note Taker:
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    *  Derek Carr (@derekwaynecarr)
    * ~~Davanum Srinivas (@dims)~~
    * ~~Jordan Liggitt (@liggitt) - OOO~~
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * Finish last comments on annual report summary and publish this!! Woohooo
        * https://github.com/kubernetes/steering/pull/209
    * Working on [https://github.com/kubernetes/community/pull/5736](https://github.com/kubernetes/community/pull/5736)
        * What other roles should be linked?
        * What is missing from context of text?
    * Open Infra Live Episode [Kendall Nelson - diablo_rojo] (won’t be joining till close to the end of the meeting)

## May 17, 2021 [Private Meeting]

* Bosun: Christoph
* Note Taker: Jordan
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * [Funding Request for streamyard](https://github.com/kubernetes/funding/issues/19) [bob]
        * individual account seems usable for our purposes (no org account option currently)
        * enables streaming without relying on individual's hardware
    * wg-policy update [cblecker]
        * got response from one existing lead, who indicated they did not have time to continue as lead
        * was able to get control of WG channels/resources from existing lead
        * attended last meeting (5/12) and spoke to attendees
        * got a sense from attendees on history of leadership
            * robert (proposed as lead, paperwork was not completed)
            * jim (proposed as lead, also hadn't completed the process)
            * plan is to communicate robert+jim as new wg-leads (when), wait for lazy consensus
            * AI: follow-up if this doesn't surface this week
        * first act of new leads will be to complete annual report questions and determine if the WG is accomplishing goals and should continue, or should shift the work elsewhere; will involve sponsoring sigs as well (primarily sig-auth)
    * Outstanding PR:
        * [Update LF course for inclusive training](https://github.com/kubernetes/community/pull/5328)
    * Annual reports
        * paris: could use help with liaisons filling in summary items for their groups, and reviewing themes
        * AI for liaisons this week:
            * merge reports from [outstanding groups](https://github.com/kubernetes/community/issues?q=is%3Aissue+is%3Aopen+%222021+annual+report%22+in%3Atitle+)
                * dims: usability
                * paris: testing
                * cblecker: service-catalog
                * derek: network, autoscaling
                * ~~liggitt: architecture~~
                * bob: apps
            * fill in summary items and review themes
            * gather and follow up with outstanding items/requests from groups

## May 3, 2021 [Public Meeting]

* Bosun:
* Note Taker: [need one]
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
* Community Members

* Topics
    * [2021 Annual Report Retro](https://docs.google.com/document/d/1X7rQx4rC9FbpdLMI5C8FY9j4TmeH2m1aKy_i4kemHXo/edit)


## April 26, 2021 [Private Special Meeting]

* Bosun: Bob
* Note Taker: Christoph
* Steering Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * [Add Program Manager governance docs](https://github.com/kubernetes/community/pull/5566)
        * **Decision:** Consensus is to close this PR and not introduce a project-wide PgM role
    * [Recognizing roles defined in SIGs and recognizing leaders in SIGs](https://github.com/kubernetes/community/issues/5722)
        * **Decision:** New roles should not be added to sigs.yaml
        * **Decision:** SIGs should be able to introduce new roles
        * **Decision:** New roles should be sent to sig/k-dev with a 1 week lazy consensus period
        * **Decision:** If a role includes delegation of responsibilities from a sig-governance role, that delegation should be documented in your charter

## April 19th, 2021 [Private Meeting]

* Bosun: Bob
* Note Taker:
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * ~~Derek Carr (@derekwaynecarr)~~
    * ~~Davanum Srinivas (@dims)~~
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
* Topics
    * [Additional Community Group Roles](https://github.com/kubernetes/community/issues/5722)
        * waiting for @dims to discuss
    * Revisit Kubernetes [CNCF project maintainers](https://docs.google.com/spreadsheets/d/1Pr8cyp8RLrNGx9WBAgQvBzUUmqyOv69R7QAFKhacJEM/edit#gid=262035321) [Bob]
        * Proposal: Add all community group leads as maintainers
        * CNCF listed maintainers are the only members **consistently** sent comms regarding CNCF initiatives, announcements, and opportunities (mentoring, KubeCon maintainer talks etc). These are inconsistently forwarded on to the leads mailing list.
        * Maintainers do get to participate in GB/TOC elections* (rules are slightly different for k8s)
            * Kubernetes Steering elects a GB seat
            * Kubernetes maintainers participate in the election of a TOC member using a [fractional voting method](https://github.com/cncf/foundation/blob/master/maintainers-election-policy.md#maintainer-voting) ([Example](https://docs.google.com/spreadsheets/d/1wrscxSVgtq8sp7vi59eVOhbFn_vrtdF49_xJqJ0yLRY/edit#gid=0)).
        * Maintainers are usually granted access to the CNCF service desk, however a separate list of members can be provided
    * [1Password approval (sc1/2/3 as root owners for org)](https://groups.google.com/g/kubernetes-sig-contribex/c/9eyzjOpT4mg/m/05CRuGGoBQAJ) [Bob]
        * Maybe just [steering-private@kubernetes.io](mailto:steering-private@kubernetes.io) instead. It’d be easier to get notifications / act on issues as that account is actively being monitored.
    * wg-policy leadership followup [cblecker]

## April 5, 2021 [Public Meeting]

* Bosun: Paris
* Note Taker: [need one]
* Steering Attendees:
    * ~~Bob Killen (@mrbobbytables)~~
    * ~~Christoph Blecker (@cblecker)~~
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Community Members
    * Kendall Nelson (@diablo_rojo)
* Topics
    * Annual Reports [https://github.com/kubernetes/steering/issues/205](https://github.com/kubernetes/steering/issues/205)
        * Due April 8th for next Chair and Tech Lead Meeting
        * “Executive Summary” should be due then, too. / Maybe ‘State of K8s Groups’ etc
            * [https://hackmd.io/R7hFEhc4SrOvaBl_vJDhEg](https://hackmd.io/R7hFEhc4SrOvaBl_vJDhEg)
        * Discuss outreach plans
            * How do we distro the summary? Blog? What else?
            * Kdev mailing list
        * Themes from reports so far
            * Diversity
            * Membership (slack? Mailing list? nothing?)
            * Automation (zoom->youtube!)
            * Autonomous subprojects (some repos started by SIGs seem to be run independently)
            * SIGs that don’t meet
                * May be review their charters?
                * 30 mins per month mandatory?
    * OpenStack Invites you to the PTG!
        * The Technical Committee will be meeting from 13-17 UTC on Friday April 23rd (Also on Thursday April 22 from 13-14 UTC) We can reserve time for you if you know when would work best :)
        * Week after kubecon
        * A few of us think Friday is good - will convene and then slack Kendall with time
    * Other business(?)
        * Governing Board meeting on Apr 14, 2021 08:00 AM Pacific Time - let Paris know
        * Governance requirements for meeting - look at the guidance

## March 15, 2021 [Private Meeting]

* Bosun: Bob
* Note Taker: @liggitt
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * WG Naming // INI OSS workstream collaboration
    * Program Manager Role
        * dims:
            * idea about recognizing SIG "coordinators" - [https://groups.google.com/a/kubernetes.io/g/steering/c/_QNg-xsawOg](https://groups.google.com/a/kubernetes.io/g/steering/c/_QNg-xsawOg)
            * Thinking of ways for people to advocate for allocating time from their day jobs to the community
            * Thinking of ways to identify people who are in SIG leadership positions
            * Would adding leaders for subprojects work for this?
            * When a SIG has a particular role, giving them a way to surface that
            * nikhita: does "coordinator" need a uniform definition
            * dims: no strong opinion, could use chairs/TL definitions, or just give SIGs a way to define a local role and set of responsibilities
            * cblecker: sounds like my understanding of what subproject owner already is
            * relevant subproject definition issues:
                * [https://github.com/kubernetes/steering/issues/200](https://github.com/kubernetes/steering/issues/200)
            * bob: several subprojects span many packages/folders/OWNERS files... makes it confusing to understand who runs the subproject
            * dims: yes, that was the intent here, to clarify and raise visibility to subproject leadership
            * cblecker: clarifying subproject owners/leadership seems like an easier next step than defining a new global term like "coordinator"; sigs can already define local roles / permissions / recognition (xref new member coordinator in contribx)
            * dims: fine with any technical solution that helps solve: surfacing the people doing the work, making sure SIGs can enumerate those people's responsibilities
            * derek: trying to map to initiatives sig-node did recently... needed a CI focus group; created a subproject and added folks to lead that; is that what we're looking for?
            * dims: that is a good example of work being done; what was missing was visibility, recognition, ability for them to call that out to their employer; ability to clearly define/sustain those positions to let other people take them over in the future
            * will continue discussion in doc/thread
    * WG Policy Leadership
        * cblecker: reached out to leads again, have not heard back; pursuing offline contact, but should start considering next steps, appointing interim leadership, etc
        * AI: cblecker to give until end of week, then raise interim leadership topic
    * [Needs Review] PSC rename to “Security Response Committee” - [https://github.com/kubernetes/community/pull/5597](https://github.com/kubernetes/community/pull/5597)
        * No objection, needs +1s, also need to sweep references to make sure name is updated
    * Annual Reports - “exec review”
        * [https://github.com/kubernetes/steering/issues/202](https://github.com/kubernetes/steering/issues/202)
        * Quick link to draft template: [https://hackmd.io/R7hFEhc4SrOvaBl_vJDhEg](https://hackmd.io/R7hFEhc4SrOvaBl_vJDhEg)
        * Target date: before kubecon, ideally 4/26, definitely by end of April
        * AI:
            * all: review template
            * liaisons: ping trailing sigs (draft PRs due 3/8, final PRs due 4/8)
                * derek/dims: sig-arch
                * others...?

## March 1, 2021 [Public Meeting]

* Bosun: dims
* Note Taker: Christoph Blecker
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * ~~Nikhita Raghunath (@nikhita)~~
    * Paris Pittman (@parispittman)
* Awesome Community Members in Attendance:
    * Kendall Nelson
    * Alison Dowdney
    * Eric Duen
    * Lauri Apple
    * Emilien Macchi
* Topics
    * [Lauri Apple] Driving adoption of process/workflow guide, program manager role, and other related improvements ([see Slack thread](https://kubernetes.slack.com/archives/CPNFRNLTS/p1613581175028300))
        * draft of [process/workflow resource](https://github.com/LappleApple/community/blob/master/contributors/chairs-and-techleads/process-workflow-guide.md)
        * [dims] where is this coming from? How does the program manager role add value?
            * [lauri] guide is meant to be tools for other folks to start a program manager type role
            * [lauri] allows for continuity in leadership (folks able to move on); has had benefits in sig-release... stephen has been able to articulate vision doc and delegate nuts and bolts implementation of that vision
        * [paris] Can we talk about the differences between this and a chair
            * [lauri] in sig-release, chairs are defining vision, and this frees them up to do more of that
            * [jordan] perhaps be more explicit about that the tasks of this doc are important to sigs. Not all sigs may need an explicit PM role; whether these responsibilities are made more explicit in the chair role or split into a separate optional role, these tools and examples are good to give sigs
            * [christoph] +1 to giving sigs examples/toolkit;
            * [jordan] describing chair tasks as "running meetings and administrivia" makes it sound like the job consists of unimportant or busywork items. I think it's important to recognize the importance of what chairs do in organizing the sig and helping drive processes to carry out the technical vision and direction, and either clarify specific responsibilities of that role or split out an optional program manager role that can be done by the same person as a chair in smaller sigs.
        * [bob] We do go over SHOULD, MAY, MUST etc in the chair description
        * [derek] PM role might be tied to subprojects based on size of subprojects
        * [lauri] Proposal is for SIG PMs to be optional role, based on their circumstances and unique needs. It’s generally a way to gain visibility for the activities, recruit new contributors to load-balance, and ensure PM work isn’t second-class citizen. Also want to encourage supportive time management for current SIG leads—ie how do we want to help leads spend their time doing the most valuable work for them, given their skillset?
        * Broad support from steering members for Lauri as sig-release chair as roles are currently defined
        * AI: clarify chair role definition to include these responsibilities, or the PM type role should be split out
            * Decide whether the chair role captures the specific responsibilities
                * If so, are current chairs meeting those responsibilities?
                * If not, should those be made more explicit in the chair role or captured in a designated role?
            * Consider applicability to sigs of different sizes/types
    * [christoph] Request from wg-policy around leadership inactivity
        * Christoph to reach out to leads between now and next steering meeting, get their input on whether the working group should continue or dissolve
    * OpenStack TC joining us :)
        * [https://etherpad.opendev.org/p/kubernetes-cross-community-topics](https://etherpad.opendev.org/p/kubernetes-cross-community-topics)
        * [2020-08-03 Distributed Project Leadership](https://governance.openstack.org/tc/resolutions/20200803-distributed-project-leadership.html) (relevant to previous program manager role discussion)

## February 22, 2021 [Private Meeting]

* Bosun: Jordan Liggitt
* Note Taker: Bob Killen
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * ~~Christoph Blecker (@cblecker)~~
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics
    * Community meeting next steps from Governance perspective [bob]
        * Overall consensus from leads (from chair and TL meeting) is the monthly meeting is not useful, and would prefer to make it more discussion focused. Feedback from ContribEx discussion has largely been the same.
        * If we take off the community meeting, we need a mechanism to ask them to do some community update and what can count for it (do one report at a KubeCon) [Paris]
        * We should have some mechanism so people can ask questions [liggitt]
        * [Draft doc of inputs and outputs and what are desired from them](https://docs.google.com/spreadsheets/d/1UP4zOZOFxiwE5J98y0f-LorrPztadPvg-wfYrS0DjRA/edit#gid=0) [Paris]
        * Regarding making the community meeting more discussion focused - the agenda would need to be curated and topics should be applicable to the broad community [liggitt]
    * Reviews on wg annual reports [liggitt]
        * lgtm from two steering members? Majority?
        * paris: inconsistent involvement of sig (some chairs copied sigs on the PR, some involved in editing, some just wrote and opened the PR). opinion: liaison has lgtm, would be good to be consistent in involving the community.
            * Decision: steering liaison has lgtm+approve, once lgtm, send annual report to sig/wg list and cc steering for comment and feedback before merge
    * Follow up items:
        * Date for next CoCC sync, transparency item to agenda
            * thread started to pick date/time for next meeting, item on agenda
        * Need to think about how we engage and do our outreach for better representation
        * Engage with CNCF on DI consultant
        * codify and write up a punch list for projects that have been moved to maintenance mode
        * Slack-admins should write up when / what action will be taken for violations better. ([Tracking Issue](https://github.com/kubernetes/community/issues/5541))
        *  Derek to follow up with the sig-networking folks about revitalizing one of the sub projects, what the choices are etc.

## February 1, 2021 [Public Meeting]

* Bosun: Paris Pittman
* Note Taker:
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * Derek Carr (@derekwaynecarr)
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Awesome Community Members in Attendance:
    * Kendall Nelson (@kendallnelson / diablo_rojo)
    * Lachie
    * Ihor
* Topics:
    * Meeting Cadence
        * Community Meeting
            * There isn’t a happy here with 37 groups and a monthly meeting; at least 4 groups update + release + other sections. Groups are giving 2-3 community meeting udpates a year.
            * KubeCon updates - at least two a year
        * SIG / Community group meeting governance
            * Once every 3 weeks
            * Folks also have office hours and other ad hoc collaboration
        * Supporting async comms
            * Do we need face to face meetings?
            * What does an alternate look like?
            * Using the mailing list as an option instead of face to face
            * More folks from india and other areas/TZ that are not PT :)
            * [jordan] doesn't want groups to turn it into a broadcast-out-only mechanism
            * [derek] Does not having an agenda a sign that there are other things going on?
                * Example: working groups - have they hit their time?
            * [kendall] Openstack has meeting bots, agenda, etc in IRC - could we do this for slack?
            * [christoph] doesn’t feel comfortable dropping them completely but wants some kind of time cadence guardrail like monthly, quarterly
            * [dims] wants to experiment
                * Dims is signing up for this :)
        * [Discussion from ContribEx](https://groups.google.com/g/kubernetes-sig-contribex/c/4yzdXsRRZGA/m/ol5F6RkhDwAJ)
    * Annual Report Check in [5 mins max]
        * Everyone meet / discuss the process with your groups? Should we extend the deadline?
            * Christoph - need one more month; april meeting
            * Dims, nikhita, bob
            * Jordan - can we do what docs does for release notes draft PRs; need to have a draft PR up just to keep it flowing
            * Follow up: ^^ that is the end result
        * Christoph - did you still want to automate the creation of the issue for each annual report
        * Other outstanding business, issues, or comments here?
        * Paris to do a draft email to kdev and leads@ for steering review TODAY
    * Follow up on ingress nginx maintainer [10 mins max]
        * Christoph needs to follow up; Derek needs to connect
    * Follow up on demographic questions for Annual Contributor Survey [5]
        * Cocc - we should engage with cncf about a dei consultant or work more at that level
        * Privacy guarantees came up as well in the discussion
        * Nikhita - what would it involve to engage with cncf on this?
        * Ihor - start with servicedesk ticket to get the convo going
        * Dims - i think we should lead the way
        * Dims - we should still think of activities to help improve while we look at survey/getting data options
            * Paris - +1
        * Christoph - engaging with cncf is worthwhile
            * Thinks a prefer not to answer should be good
            * We already do some anon scrubbing before the data goes live anyway
            * If we delay on this, we are waiting another year to make moves on these efforts
        * Jordan - formulating what we are wanting to know and why
            * Ball is in our court there
            * Describe what we want to be able to compare year over year
        * Nikhita - not sure why this would take a year
            * Definitely start the process
        * Bob - maybe we do the survey mid year
            * +1
        * Derek - +1 to Jordan
    * CNCF / TOC / GB updates [2]
        * Next GB is April
        * Voting today for TOC seat
        * Gsuite update for CoCC secure shared drive
            * Bob filed; anyone of the sc1, sc2, sc3
            * Bob button
        * [LFX Mentoring open for 2021](https://github.com/cncf/mentoring/issues/317)
    * KubeCon Maintainer Track session? [2 mins]
        * Christoph - due soon, do we want to do anything?
        * Last year we did a virtual panel and then a project update in Q3?4?
        * We also have a project update
        * Bob - maybe ama?
        * Nikhita - no opinion
        * Jordan - no opinion
        * Derek - no opinion
        * Dims - no opinion
        * Christoph - in person amas are good, but we have a monthly opportunity for the community to ask steering questions (this meeting), so virtual AMA is maybe of questionable benefit?
        * [community] - agree with folks about amas
        * Paris - will take the action
    * https://github.com/kubernetes/steering/issues

## January 18, 2021 [Private Meeting]

* Bosun: Jordan Liggitt
* Note Taker: Jordan Liggitt
* Steering Committee Attendees:
    * Bob Killen (@mrbobbytables)
    * Christoph Blecker (@cblecker)
    * ~~Derek Carr (@derekwaynecarr)~~
    * Davanum Srinivas (@dims)
    * Jordan Liggitt (@liggitt)
    * Nikhita Raghunath (@nikhita)
    * Paris Pittman (@parispittman)
* Topics:
    * 2020 WG report wrap-ups
        * [Jordan] Please remember to check in with your wg to have them PR the report into the repo
    * 2021 SIG/WG report process [updates](https://github.com/kubernetes/community/pull/5388)
        * Bob/Jordan comments mostly addressed
        * AI:
        * Open questions:
            * "About you" liaison/1-on-1: per group? chairs individually? chairs and TLs individually?
            * Paris: meeting individually is a large number, but is only yearly... individual meetings allow folks to raise topics that might be difficult in a group setting. Any thoughts on how to scale this to make sense?
                * Dims: maybe start with group meetings, liaison follow-up with individual meetings if it seems warranted
                * Jordan: I like that, would also explicitly let chairs/TLs know they can reach out to their liaison for a 1-on-1 if they would like
                * Nikhita: is this throughout the year or trying to do it before annual report is submitted?
                    * Paris: not necessarily accompanying reports, but would like to see it happen early in the year
                    * Jordan: +1, not blocking annual report on it, but doing it early
                    * Bob: +1, any issues might also be fresh on their mind after compiling the annual report
                * AI: Paris will update PR to clarify
            * Paris: how to raise community awareness of the annual reports? sent an email to k-dev last year, but certain it was lost/ignored
                * Bob: as a data point, in the leads meeting, many were not aware this was a thing and was imminent. Suggest opening an issue per group and assigning to the leads, their PR with the report can close the issue.
                * Paris: individual issues or umbrella with checklist?
                * Bob: suggest individual so the report PRs can close them
                * AI (Christoph): easily scripted if we want to automate report issues assigned to leads.
                * Christoph: are we asking WGs that just filed a 2020 report to report again?
                    * Paris: no, 2020 WGs shouldn't have to repeat themselves. could document or link to their 2020 reports (and WGs submitting reports now could title them 2021)
                * Nikhita: can we move liaison assignments to sigs.yaml? will keep it up to date
                    * Bob: +1
                    * Jordan: +1
                    * Paris: +1
                    * Christoph: +1, let's remove the data from liaisons doc and link to the canonical location
                    * ~~AI (nikhita): add to sigs.yaml, link from liaisons.md~~
                        * [https://github.com/kubernetes/community/pull/5415](https://github.com/kubernetes/community/pull/5415)
    * Website banner approval mechanism [implemented](https://github.com/kubernetes/website/pull/25769)
        *
    * BLM banner expiring
        * steering committee approved the BLM banner staying up through the end of 2020 when it was restored after kubecon NA
            * (paris was -1)
        * dims: can we keep the announcements text and the dates it was up in the announcements source, so it is browseable/searchable, even though it is no longer rendered on the website because it expired?
            * yes, several +1s on adding history to the source
            * ~~AI (jordan): set up announcements source containing announcements/dates posted in 2020 -~~ [https://github.com/kubernetes/website/pull/26149](https://github.com/kubernetes/website/pull/26149)
        * AI (bob): transcribe Christoph statement on community values
        * What are the areas where we want to see continued progress in underrepresented groups?
            * Paris: OWNERS files
            * Dims: go back to funnel perspective. To increase representation of underrepresented groups at the leadership level, we need to work from the top. Increase the number of folks interested in the technology, increase the number of folks attending kubecon, increase the number of folks getting into the community (communitybridge/gsoc/outreachy), get a bunch of them into reviews, approvers and finally sig TL’s or chairs.
            * Bob: Annual survey could give us insight into where people are in the ladder
        * AI (bob): Looking into adding background to annual community survey. Language will require sign off from CNCF.
            * CNCF has given the go ahead and will supply language they use from Kubecon registration
    * [SIG Release has a charter update](https://github.com/kubernetes/community/pull/5409)
    * Ingress-nginx maintainer update
        * primary maintainer stepping down
        * AI (christoph): networking liaison (derek) check with networking leads on next steps, if transfer is being handled
    * Wg-component standard - follow up in slack
    * [nikhita] FYI need to follow up on election retro items
        * ~~AI (nikhita): will post tracking issues for AIs in slack~~
