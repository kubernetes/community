## December 19, 2022 [Private Meeting]

**Bosun**:

**Note taker**:

**Steering Attendees**:

* Ben Elder (@bentheelder)
* Bob Killen (@mrbobbytables)
* Carlos Tadeu Panato Jr. (@cpanato)
* Christoph Blecker (@cblecker)
* Nabarun Pal (@palnabarun
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

**CNCF updates?**

**Votes**

**Topics**

* [Bob] KubeCon EU Session
    * Standard simple K8s SC panel is not appropriate for KubeCon in terms of panel representation, raised in Detroit, now formally part of the CFP and maintainers track panels
    * We will not tokenize anybody by simply asking them to participate in response to the program committee suggest of changing the panel _membership_
    * Could shift the _topic_, ie: not just panel of “SC members speaking” versus panel of “Current governance trends and changes compared to the past across projects”
    * We can’t ad hoc change the structures of our sitting committees, but we do need to foster a more inclusive Steering Committee membership ahead of 2023 election
* [Christoph] Workflow / ContribEx idea to float for maintainers: a label for OWNERS file PRs for easier triage and dashboarding to keep eyes on tracking conversation to completion
* [Bob] Due for a next quarterly sync with K8s CoCC in January
* Cancel Jan. 2, 2023 Steering Committee meeting
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## December 5, 2022 [Public Meeting]

**Bosun**: Bob Killen

**Note taker**: Tim Pepper

**Steering Attendees**:

* Ben Elder (@bentheelder)
* Bob Killen (@mrbobbytables)
* Carlos Tadeu Panato Jr. (@cpanato)
* Christoph Blecker (@cblecker)
* Nabarun Pal (@palnabarun
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Arnaud Meukam
* Noah Abrahams

**CNCF updates?**

**Votes**

**Topics**

* Welcome Carlos!
* [Bob] Annual Reports
    * [Tools updates](https://github.com/kubernetes/community/issues/6357#issuecomment-1336320564) req’d?
        * Overall theme: need to make this easier, more templated for consistency/completeness
        * ContribEx shadow cohort is starting some work here, Christoph Blecker is leading push on this
        * Need issues and template PRs, may need a code change if the generator is to be run in 2022 for 2023
        * AI SC members: review the annual reports [retro](#bookmark=id.yeoum670jlxc) for potential points of update to the template questions: [https://github.com/kubernetes/steering/issues/242](https://github.com/kubernetes/steering/issues/242)
        * Annual Report Retro:
            * [https://docs.google.com/document/d/1X7rQx4rC9FbpdLMI5C8FY9j4TmeH2m1aKy_i4kemHXo/edit](https://docs.google.com/document/d/1X7rQx4rC9FbpdLMI5C8FY9j4TmeH2m1aKy_i4kemHXo/edit)
    * [Liaisons](https://github.com/kubernetes/community/blob/master/liaisons.md):
        * Time for liaisons to ping SIGs/WGs
        * Re-shuffle liaisons or new SC members just take on prior members’?  Let’s keep it simple, [Carlos takes on Paris’ spot](https://github.com/kubernetes/community/commit/c5fa26a3ba44445b2bbbab5c31ca097a02250a27)
    * Timeline:
        * Definitely would like to be done by KubeCon EU (3rd week of April 2023)
        * Need buffer for SC to collate all parts into one report.  End of March-ish target for SIGs/WGs?
        * Pitch to leads that having reports done is part of what they can use in their maintainers track talks
        * If goal is completion in March, work by leads in February/March means tooling changes need done ASAP on tight timeline
* Signing MacOS binaries [Arnaud]
    * Renewing question from [https://github.com/kubernetes/funding/issues/30](https://github.com/kubernetes/funding/issues/30)
    * Paris had been looking at Apple
    * Knative/Mohammed/SIGRelease has some prior work on this, might have been prior to CNCF though, maybe Google helped on it?
    * [https://developer.apple.com/programs/enroll/](https://developer.apple.com/programs/enroll/) has individual vs organization info.  Org account needs to be a “legal entity”.  Kubernetes does not meet the requirement and probably can’t itself
    * As per [https://github.com/cncf/foundation/blob/main/charter.md](https://github.com/cncf/foundation/blob/main/charter.md) point 2.d.i  “The foundation seeks to offer up a fully integrated and qualified build of each of the constituent pieces, on a well-defined cadence across the reference architecture.”  Reasonable next level of maturation for CNCF to have a signing service.
* Minikube is interested in enabling opt-in telemetry
    * Complicated, don’t want to steward data, whoever is must comply with regulations
    * K8s infra has some similar work related to the registry, probably “ii” did a PII review of telemetry, link [https://registry.k8s.io/privacy](https://registry.k8s.io/privacy)
    * LF Policies:
        * Telemetry specific: [https://www.linuxfoundation.org/legal/telemetry-data-policy](https://www.linuxfoundation.org/legal/telemetry-data-policy)
        * Privacy generic: [https://www.linuxfoundation.org/legal/privacy-policy](https://www.linuxfoundation.org/legal/privacy-policy)
    * Additional other (eg: GDPR) compliance requirements
    * If we do this route, where is the data to be stored, what mechanism controls access?  Do we want to go this route as an open source project?  Can of worms, deep rabbit hole…
    * Project has had previous telemetry in [https://github.com/kubernetes-retired/spartakus](https://github.com/kubernetes-retired/spartakus), but it has been archived/retired because it was unused
    * There is precedent in open source for collecting info on usage patterns to inform the project’s development roadmaps.  Goal of stats of versions, drivers, etc. is more easily distinguished from PII.  Goal is not anywhere near tracking humans, versus the installed codebase and useful features.  Still have identifiers of some sort (eg: for validation, deduplication, …)
    * Does CNCF have any mechanism today standardized for projects? AI: open service desk ticket to inquire on status of existing tools
    * What about timing out, purging data?  User request for removal of their data?
    * Is there a pre-approved third party route (eg: Google Analytics)?
    * LFX has a project portal aimed at perhaps this type of project level analytics, and they must already manage PII
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## November 21, 2022 [Private Meeting]

**Bosun**: Bob Killen

**Note taker**:

**Steering Attendees**:

* ~~Ben Elder (@bentheelder)~~
* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* ~~Nabarun Pal (@palnabarun)~~
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**CNCF updates?**

**Votes**

**Topics**

* 2022 year-end annual reports
    * Stay the course?
    * Update/revise process (and tools)?
    * Timeline?
    * Update liaisons list…
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## November 7, 2022 [Public Meeting] (recording)

**Bosun**: Bob Killen

**Note taker**: Tim Pepper

**Steering Attendees**:

* Ben Elder (@bentheelder)
* Bob Killen (@mrbobbytables)
* ~~Christoph Blecker (@cblecker)~~
* Nabarun Pal (@palnabarun)
* ~~Paris Pittman (@parispittman)~~
* ~~Stephen Augustus (@justaugustus)~~
* Tim Pepper (@tpepper)

**Community Attendees**:

* Kendall Nelson

**CNCF updates?**

**Votes**

**Topics**

* No SC quorum
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## October 17, 2022 [Private Meeting]

**Bosun**: Christoph Blecker

**Note taker**:

**Steering Attendees**:

* Ben Elder (@bentheelder)
* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Nabarun Pal (@palnabarun)
* ~~Paris Pittman (@parispittman)~~
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**CNCF updates?**

**Votes**

**Topics**

* [cblecker] Discussion on statement around equality
* [tpepper] Recap of joint meeting with CoCC
* [all] misc onboarding discussion, context sharing, administrivia


## October 3, 2022 [Public Meeting] (recording)

**Bosun**: Christoph Blecker

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables) - continuing
* Christoph Blecker (@cblecker)
* ~~Davanum Srinivas (@dims) - outgoing~~
* Jordan Liggitt (@liggitt) - outgoing
* ~~Paris Pittman (@parispittman)~~
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)
* Ben Elder (@bentheelder) - incoming
* Nabarun Pal (@palnabarun) - incoming

**Community Attendees**:

* Noah Abrahams
* Rey Lejano
* Kaslin Fields (@kaslin)
* Antonio Ojea
* Ben Elder
* Nabarun Pal
* Ian Coldwater

**CNCF updates?**

**Votes**

**Topics**

* **[cblecker] 2022 Steering Committee Election Results**
    * Bob Killen
    * Nabarun Pal
    * Ben Elder
    * Thank you to Kaslin, Noah, Dims for running the election!
    * Election retrospect coming soon (tomorrow?), [doodle poll](https://doodle.com/meeting/participate/id/bkZn335a), discussion [doc](https://docs.google.com/document/d/1q5jdCAWfRIWK08POymFuoqS4BY2bSfRFxAQSk5xpvx8/edit)
* [cblecker] Steering AMAs: main conference session, and contributor summit
    * KubeCon is three weeks away!
    * Would like Bob, Nabarun, Ben at these two sessions if planning already to attend
* [liggitt] need an issue/template for onboarding ben/nabarun, offboarding dims/liggitt
    * template: [https://github.com/kubernetes/steering/blob/main/operations/onboarding.md](https://github.com/kubernetes/steering/blob/main/operations/onboarding.md)
    * example: [https://github.com/kubernetes/steering/issues/219](https://github.com/kubernetes/steering/issues/219)
    * Stephen can create, new folks will drive completion
        * Opened: [https://github.com/kubernetes/steering/issues/256](https://github.com/kubernetes/steering/issues/256)
* [Kendall] OpenInfra Foundation was having a f2f the week before KubeCon, our Steering members were invited.  This has shifted to virtual, we’re still invited, ping Kendall for link to participate
* [Nabarun] Will be stepping down from CoCC as of SC election:
    * SC needs to pull in the next available candidate, Tim will lead
    * Stephen: SC has recently changed its process for its own elections, need to consider if similar makes sense for CoCC.  Consequence would be holding a fresh election instead of pulling from prior elections next in line
* [end of meeting] any topics of interest for the next Community Meeting?


## September 12, 2022 [Public Meeting] (recording)

**Bosun**: Stephen Augustus

**Note taker**: Jordan Liggitt

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Stephen Augustus (@justaugustus)

**Community Attendees**:

* Arnaud Meukam
* Noah Abrahams
* Nabarun Pal

**CNCF updates?**

* KubeCon Maintainer updates
    * Anything you want to highlight on the big stage?
    * reply to [leads@kubernetes.io](mailto:leads@kubernetes.io) email with topics/information ASAP (link to survey in email has September 23rd deadline)
        * k8s.gcr.io -> registry.k8s.io ?
* dims - four issues currently open for health checks of various projects
    * stephen: would it make sense to make sure impacted sig areas are aware? e.g. sig-apps for brigade? sig-api-machinery are already aware of etcd issues.
    * lag in two-way communication took longer than desired for this issue
        * etcd didn't communicate clearly this behavior for single-node clusters (even some etcd maintainers were actually surprised by this)
        * etcd didn't have clear communication from downstreams that this behavior was assumed to be correct
    * current etcd issue has traction to be fixed, agreement in principle on backporting a fix to 3.4, 3.5
    * how can etcd tighten communication for future issues?
    * christoph:
        * we (k8s) are also a foundational building block for many downstreams
        * what lessons can we learn about communicating better/earlier about surprising behaviors or "known" issues to our downstreams
    * stephen: agree, communication around releases, blogposts communicating non-feature changes coming in releases has improved
    * dims:
        * consider well-known / popular / public installer default behavior and target communications
        * suggestion for etcd maintainers: let major distributors/installers indicate their use to stay informed of known issues impacting them
* [https://www.linuxfoundation.org/blog/welcoming-pytorch-to-the-linux-foundation/](https://www.linuxfoundation.org/blog/welcoming-pytorch-to-the-linux-foundation/)
* Two initiatives upcoming related community safety/trust, announcements hopefully incoming soon

**Votes**

**Topics**

* Steering Committee election is underway (voter exception requests end 2022-09-16, voting ends 2022-09-30)! ([readme](https://github.com/kubernetes/community/tree/master/elections/steering/2022), [ballot](https://elections.k8s.io/app/elections/steering---2022))
* Sponsor Developer Programs
    * Developer accounts are required to ensure signatures are produced in the release process of k8s for specific operating systems.
        * e.g. Apple Developer program: [https://github.com/kubernetes/funding/issues/30](https://github.com/kubernetes/funding/issues/30)
        * discussed in [May meeting](#bookmark=id.kdszhzcl9bt3)
        * the issue is not the money / reimbursement, the issue is tying the process to individuals
        * Paris is trying to dig into it for MacOS
        * Stephen will sync with Paris
        * Arnaud will open umbrella issue for Windows and MacOS
        * Stephen: sig-release is doing a proof of concept for getting build artifacts to stop being reliant on individuals
        * Jordan: in addition to not wanting to inject individuals into artifact production critical paths, we don't want to tie historical project artifact signatures to individuals in ways where future accidents on an individual account could result in invalidation/revocation of project artifacts
        * Christoph: in addition to technical aspects, developer agreements are contracts and we need to make sure Kubernetes/CNCF is the party involved in the account being used to sign, not a release manager personal account; would suggest not lumping windows/mac issues into an umbrella, keep them separate since resolution will likely be independent
        * Arnaud: what is the next step?
        * Christoph: macos issue is open, paris/stephen have next step to figure out how a non-personal account can be obtained; need a separate new issue for windows/microsoft if that signing is needed
    * Dims: update on the CDN image registry effort. we did a proof of concept on cloudflare fronting the image registry; on pause for the moment, feel like we should also do a proof of concept on alternate providers
        * arnaud: personal feeling is that we should move on from cloudflare. we do want to use a CDN. we tried to reach out to fastly and weren't able to make contact
        * dims: fastly is on my TODO, via ChrisA
        * Christoph: if we're back to a vendor selection / POC stage, should we bring the list of requirements to CNCF?
        * dims: we sort of know what we want, trying to validate approach will function properly before crafting a formal list of requirements
        * Christoph: once we have the technical requirements, could be worth a call for vendors that can meet those requirements to make sure we don't overlook an option, especially if there's a CNCF member that could meet those requirements more affordably
        * dims: FYI, cloud provider credits program ([https://cncf.io/credits](https://cncf.io/credits)) is not yet active
        * arnaud: needed to consider CDN providers who partner with google ([https://cloud.google.com/network-connectivity/docs/cdn-interconnect](https://cloud.google.com/network-connectivity/docs/cdn-interconnect)) because current artifacts are google-hosted
        * stephen: need to clarify whether the list of providers at ^ is the pool we have to choose from or not
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?
    * TBD on Steering results/schedule
        * jk, it’s Christoph for October!

## August 14, 2022 [Private Meeting]

**Bosun**: Tim Pepper

**Note taker**:

**Steering Attendees**:

* ~~Bob Killen (@mrbobbytables)~~
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* ~~Paris Pittman (@parispittman)~~
* ~~Stephen Augustus (@justaugustus)~~
* Tim Pepper (@tpepper)

**CNCF updates?**

**Votes**

**Topics**

* **[private]** CNCF CoCC WG update [Christoph]
* Steering Election [dims]
    * Schedule merged last week: [https://github.com/kubernetes/community/pull/6760](https://github.com/kubernetes/community/pull/6760)
    * Today due for “Announcement of Election and publication of voters.md”: voter roles pulled last Friday.
    * TBD in next 2 weeks:
        * candidate call so folks can have Q&A around the role
        * Consider scheduling 2x for timezones
* K8s CoCC election update [Tim]
    * Newly elected:
        * Danielle Lancashire
        * Hilliary Lipsig
        * Xander Grzywinski
        * Jeremy Rickard (Vallery stepped down additionally, after election)
    * CoCC is working to onboard folks
    * Joint meeting this week or next
* KubeCon NA Oct. 24-28, 2022:
    * Contrib summit presence? Bob is involved in planning, it’s early days still
    * AMA in main program:
        * need to define speakers/content (beyond the audience asking questions): Dims, Jordan, Bob won't be there as SC members, add additional names in [https://sched.co/182Pc](https://sched.co/182Pc) this week
        * Oct. 3 is our election results announce, hopefully newly elected will be at KubeCon
* GB Updates [Paris? Dims?]
    * Upcoming meetings on legal committee and follow ups after Valencia
* Etcd update?
    * Dims pinged recently for update
    * Marek (Google) is primary person left maintaining
    * Clarifying moves (or have moved?) inactive maintainers to emeritus to show the shortness of the list
    * Drafting position statement on new feature working being ~ on hold
    * Have discussed the cloud credits program with folks (some past jepsen testing was on a personal bill on AWS?)
    * Jordan: stabilizing and communicating, is there more we can do to spread the word on needs?
    * Dims: partly blocked lately on Marek being on holiday in August
    * Jordan: at what point does this need an escalation to CNCF ToC for broader call for help?
    * Dims: has happened, but currently hoping on some face to face in Detroit to convince folks of the need
    * Jordan: will reach out to Marek to see what support he could need
    * …at what point do we abstract crisply to try to enable a potential switch to something well maintained?  That only helps us, not all the other dependents on etcd.
        * Where are the RedHat experiments on crdb?
    * …or pull into k8s proper to fully support?  Still no people for that.


## August 1, 2022 [Public Meeting] (recording)

**Bosun**: Christoph Blecker

**Note taker**: Tim / Jordan

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* ~~Paris Pittman (@parispittman)~~
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* David Eads
* Rey Lejano
* Danielle
* Kendall Nelson
* Seth Jennings
* Lubomir
* Daniel Smith

**CNCF updates?**

**Votes**

**Topics**

* bosun for next times:
    * Aug 15 (private) - TimP
    * Sept. 5 (public) - Stephen Augustus
    * Sept. 19 (private) - Stephen Augustus
* [Bob] Election schedule review: [https://github.com/kubernetes/community/pull/6760](https://github.com/kubernetes/community/pull/6760)
    * targeting merge by 8/5
    * AI: steering review/comment early this week
* [Kendall Nelson] Any Steering Committee members interested in attending the next OpenStack PTG meeting Oct. 17-20 [https://openinfra.dev/ptg/](https://openinfra.dev/ptg/)
    * driving distance from Detroit :)
    * early-bird pricing through 8/15
* [cblecker] CoCC election update – TimP
    * Dates
        * August 4th - Nominations close ([https://forms.gle/4LPrENt1roTfG5Q8A](https://forms.gle/4LPrENt1roTfG5Q8A))
            * can nominate yourself or someone else
            * nominees can be outside the community ([eligibility guidelines](https://github.com/kubernetes/community/blob/master/committee-code-of-conduct/election.md#eligibility-for-candidacy))
        * August 5th - Steering to hold election
        * August 6-7th - Steering to confirm with candidates and announce results
    * Have 3 openings and 3 folks nominated so far; would like more in order to have future backups if needed
* [dims] can Christoph give an update on the CNCF-level CoCC progress?
    * Feedback process underway on defining the bootstrapping and permanent bodies
    * future election for members of the permanent committee
    * Dates (details in thread at [https://lists.cncf.io/g/cncf-kubernetes-maintainers/message/288](https://lists.cncf.io/g/cncf-kubernetes-maintainers/message/288))
        * ~~Self-nomination period - June 23 to July 11~~ - COMPLETE
        * [Voter opt-in period](https://docs.google.com/forms/d/e/1FAIpQLScguua9zdH1drcHGJcu_okGgw5pXofQIcCCdnZ7jJ30je_qrw/viewform) for the approvers list- July 18 to August 1 at 12 PM PT
        * Candidates will be announced on August 2
        * Ballots will be sent on August 2
        * Voting period - August 2 to August 11 at 12 PM PT
        * Election results announced on August 11
* [Paris? GB updates?]
* [end of meeting] any topics of interest for the next Community Meeting?

## July 18, 2022 [Private Meeting]

**Bosun**: Christoph Blecker

**Note taker**: Jordan

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Votes**

**Topics**

* [paris] ask v discussion: liaisons, can you ask your chairs and TLs what mentoring programs they are involved in for more reviewers+? Contributor ladder mentoring is awesome. Example: [https://github.com/kubernetes/community/issues/6665](https://github.com/kubernetes/community/issues/6665) - cli and apps are looking for more reviewers; this is not a program for new contributors
    * Chairs and tech leads should not be mentors; this work should be delegated
    * Annual reports made it clear ~every group is looking for more reviewers/approvers, but that isn't matched by mentoring programs
    * paris: if people wanting to start up a mentoring effort need pointers, paris can be a resource if they want to reach out
    * cblecker: agree with conclusions / need for this, and that mentoring programs are a good way to do this… is this steering or contribx?
        * paris: from contribx perspective, need help
        * cblecker: if chairs/tls don't have capacity to do the mentoring, what are we asking them to do?
        * paris: ask chairs/tls to ask subproject owners to be mentors
    * AI for liaisons: remind chairs/TLs to ask subproject owners about mentoring (not for chairs/TLs to take this on themselves, but to delegate). Highlight contribx/paris as resource for folks wanting to start a mentoring effort.
        * Bob Killen (@mrbobbytables)
        * Christoph Blecker (@cblecker)
        * Davanum Srinivas (@dims)
        * Jordan Liggitt (@liggitt) - done 7/19
        * Paris Pittman (@parispittman)
        * Stephen Augustus (@justaugustus)
        * Tim Pepper (@tpepper)
* [paris] etcd - how do we make sure they get the support they need?
    * paris: they need a senior community engineer direly; no onboarding programs and need more building of contributor documentation and videos
    * paris: published maintainers list isn't accurate, needs cleanup to reflect active maintainers to make it more clear what the current state is; feedback from GB has been confused because the published maintainers list looks like there is plenty of support
    * dims: have a contributor (benjamin) from vmware side, doing good work, helping with release, but not currently a senior contributor; another 2 contributors from AWS expected to show up in 2-3 weeks from now, once they free up from internal work.
    * tpepper: can we coordinate conversation around this at kubecon / contrib summit?
    * bob: has an etcd status update been presented to the GB?
        * paris: no, not yet
    * paris: should we ask CNCF for joint meeting with actionable next steps? has been six months without much movement
        * dims: dev reps should request time to talk specifically about this
* [paris] steering committee election: by the end of july we need to -
    * Select the voter eligibility - 50 contributions? Same as last year?
        * Confirm committee members and named roles will be auto-in - make sure all PRs and changes are in for this
        * Exceptions:  does exception basis voting eligibility lead also to ability to run for role?  Can exceptions be moved to formal voting eligibility requirements?  Is this a problem (ie: how many exceptions are happening)?
    * Confirm officer selections from josh
        * initial selections made, not yet presented to steering?
    * Make any last minute changes to election policy and procedure?
        * Do you need to be an eligible voter to stand for election?
            * same question raised in [https://github.com/kubernetes/steering/issues/227](https://github.com/kubernetes/steering/issues/227), discussed in a public meeting at  / [recording of discussion](https://www.youtube.com/watch?v=wEi47H7_Bxo&t=1306s)
            * stephen and paris and TimP: +1
            * stephen: if we do this, should look at the common exception scenarios and try to fold those into official eligibility rules so that exceptions are not required for most of the people who would be voting / running (+1 from Jordan, mixing judgment call exceptions as a gate to candidate eligibility seems problematic)
            * ran out of time on this topic, will continue discussion async and resolve before end of July
    * Figure out the GB alt situation so we can implement that this round
        * Should it alt every election for primary and alt seat?
        * Should all members including emeritus / no matter the term limit be included?
* [paris] cocc election: by the end of July we need to -
    * Make sure all election policy and procedure reflects the reality we are in and the future we want to go to
    * Start the nomination process from Chairs and Tech Leads first and then open up to the wider community
        * Example of wider community: [https://groups.google.com/a/kubernetes.io/d/msgid/steering/CADPpvp1SzFYJDR%3DZqdyOLNdZqZmYm193vVfEkVzBbvRfv0hJJA%40mail.gmail.com](https://groups.google.com/a/kubernetes.io/d/msgid/steering/CADPpvp1SzFYJDR%3DZqdyOLNdZqZmYm193vVfEkVzBbvRfv0hJJA%40mail.gmail.com?utm_medium=email&utm_source=footer)
        * Chairs and TLs comms: ASAP
        * Dev and Steering: 3-5 days after Chairs and TLs comms
    * [paris] funding: 4 out of 5 issues need responses from us
        * [https://github.com/kubernetes/funding/issues](https://github.com/kubernetes/funding/issues)
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## July 7, 2022 [Public Meeting]

**Bosun**: Christoph Blecker

**Note taker**: dims

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Jeremy Rickard
* Jason DeTiberus
* Bridget Kromhout
* David Eads
* Kirsten Garrison
* Tim Hockin
* Keith Mattix
* Chris Short
* Madhav Jivrajani
* Mike Morris
* Josh Berkus
* Nabarun Pal
* Priyanka Saggu

**CNCF updates?**

* [paris] GB updates [5 minutes]
    * We have short term and long term needs+goals at governing board right now
    * Legal committee and/or legal matters
        * New legal committee spun up
            * Made up of folks on GB (and their legal folks)
            * Proposing the opt-in legal private list next time they meet
                * Similar to ASF’s legal-private mailing list
        * Waiting on answers to questions from d&o insurance ask
            * [https://github.com/cncf/foundation/issues/329](https://github.com/cncf/foundation/issues/329)
        * Legal fund resolution: Need to get details from cra and others; maybe have this?
            * Use general funds?
            * Need to create/define a process
        * [“Real Name” / Pseudonym issue](https://github.com/cncf/foundation/issues/383)
            * Relates to DCO, are real names required, what even is a real name
    * Special issues committee
        * Good news! There is a process (yay!)
            * Everything needs to start at servicedesk before going to board - even the ambiguous/large scale/high level strategy asks and questions
                * Will submit servicedesk tickets for SIGs with descriptive needs for more reviewers, etc
                * Decision making process has been defined
            * Resolution update for community of practice for managers of contributors/maintainers + training for participation in cloud native communities: set servicedesk ticket; amye+chris likes the idea and can fund from general bucket
                * CoC Community of Practice may be staffed this way
            * Working on creative staffing fund with group
                * Fellows, contractors, and arrangements that can move the project forward in areas that are needed
    * Introduce an alternate [https://github.com/kubernetes/steering/issues/218](https://github.com/kubernetes/steering/issues/218) for Kubernetes GB representative
        * Most all reps do already have an alt, Kubernetes is 1 of 2 w/o
        * Important to have alt. for continuity and bringing up next people, also people sometimes do have to miss meetings.
        * Kubernetes does have good representation between Paris and Dims, but Dims will roll off
        * [Stephen] Choose staggered term alternate members
        * The alt. is a non-voting member, but can proxy for the regular
        * Paris is +1 as proposer, dims, bob, jordan, christoph indicated +1 in the meeting
* [paris] FYI only contributor kubecon ticket and/or travel fund service desk ticket filed [2 minutes]
    * CNCF folks will start to work on this post break
    * Includes leads ask for give-away recognition tickets & shwag
    * Many refs but this one is an oldie: [https://github.com/kubernetes/steering/issues/109](https://github.com/kubernetes/steering/issues/109)
    * Not on track for resolution ahead of KCCNCNA22, but we’ll keep pushing on it

**Votes**

* [stephen] [https://github.com/kubernetes/steering/pull/249](https://github.com/kubernetes/steering/pull/249) [10 minutes]
    * No-op PR for reorganization with elections, vacancies, and changes
    * Christoph, Dims, Jordan, Stephen, Bob, Paris, TimP - +1’s

**Topics**

* [paris] elections update [5 minutes]
    * Josh Berkus from ContribEx:
        * creation of election subproject: [https://github.com/kubernetes/community/pull/6720](https://github.com/kubernetes/community/pull/6720)
        * Support elections for non-steering leadership roles across the community
        * Need additional volunteers for election officers
            * Dims is rolling off of steering, will help.
            * Target is chosen election officers with draft election schedule submitted by end of July
        * Feedback from SRC about elekto security. End of the month elekto will be paused to fix issues.
        * No actions from steering needed until officers/schedule are submitted
    * Ours / SC [https://github.com/kubernetes/steering/blob/main/elections.md](https://github.com/kubernetes/steering/blob/main/elections.md#eligibility-for-voting)
    * Election for Code of Conduct Committee
        * TimP to run the CoCC election as [liaison](https://github.com/kubernetes/community/blob/master/liaisons.md)
        * Please ping Josh for using elekto
* [christoph] AMA style Steering Committee session for KubeCon NA 2022 (submissions due Monday, July 11)? [3 minutes]
    * stephen: will be there, in favor, normal topics plus output of annual report
    * paris, dims, tim: +1
    * +1 from several community folks as the sessions being helpful. several liked the AMA style
    * Will submit an AMA style discussion again
    * [Bridget] Talk about behind-the-scenes
        * Path to Steering?
* [Bob] [WG Proposal: Service Mesh](https://github.com/kubernetes/community/pull/6724) [15 minutes]
    * [Charter Proposal](https://docs.google.com/document/d/1Qt3N1RLnRom7jKNhcnNLg_xaPn0TwKGuxxgXLOguNvE/edit)
    * surfacing proposal to steering
    * single kubernetes stakeholder sig
    * thockin:
    * keith mattix: didn't want to overwhelm gateway API folks; subproject seemed a little weird because this wouldn't be owning code
    * stephen: normally a working group would be across sigs and time-bounded (to completion of some project); lots of subprojects don't "own code"
    * CNCF Service Mesh WG is a natural ally? [https://github.com/cncf/tag-network/tree/master/service-mesh-wg](https://github.com/cncf/tag-network/tree/master/service-mesh-wg)
    * Can be an additional stream in an existing subproject and have all the infra (zoom meetings, lists, etc.) without bothering with the full on WG formality
    * Christoph: not an explicit requirement to have multiple stakeholder sponsoring SIGs (see [https://github.com/cncf/tag-network/tree/master/service-mesh-wg](https://github.com/cncf/tag-network/tree/master/service-mesh-wg))
    * Mike Morris: motivation is to not be parallel to gateway api to avoid confusion on code ownership/authority.  This would be focused on channeling proposals to gateway project, not owning implementation.
    * TimH: we can pragmatically NACK proposal and go with simple sub-sub-project, add addl formality & structure if/when needed
* [stephen] [https://github.com/kubernetes/steering/pull/249](https://github.com/kubernetes/steering/pull/249) [10 minutes]
    * No-op PR for reorganization with elections, vacancies, and changes
    * Unanimous +1 vote, “/hold” removed so it can merge.
* [paris] [https://github.com/kubernetes/steering/pull/248](https://github.com/kubernetes/steering/pull/248) [20 minutes]
    * Lifecycle was missing
    * Added voting as a place to house abstention
        * Should we take this out and put in an issue?
    * stephen: summary of current state at [https://github.com/kubernetes/steering/pull/248#issuecomment-1177525893](https://github.com/kubernetes/steering/pull/248#issuecomment-1177525893)
    * How do we communicate context on a no confidence vote?
        * David Eads: Critical to communicate why a removal is being considered, so observers can understand what is happening.  Personally wants to see how each SC member has voted on a removal.  Wants to know as a community member that they have a way to feedback into choices made.  Charter change process via PR is good.
        * Stephen: some kind of reason is important; requiring public reasons is not good
        * David: a private but trust me is not satisfying
        * TimP: committees are chartered allowing private activity.  Transparency and trust versus that are ongoing concerns.  CoCC for example has struggled with this for years.
        * Stephen: leads to bigger, scarier question of lack of trust leading to dissolution (issue link?)
        * Chris Short: a no confidence vote that tells community members "just trust us" isn't likely to fly because community members don't typically know steering members personally
        * Paris: detailed logging creates future harms and biases.  Definitely do have public meetings and discussion (eg: this meeting has agenda, minutes, recording)
        * Dims: my perspective is that the most we can do is summary text of the vote and a small description of why we don't have confidence; more details could be recorded for future steering members to review
        * cblecker interjection: will prioritize folks that haven't spoken
        * Kirsten: observing that by deferring some points of the PR to a later date, a merge of just some points doesn't give any indication that there were deferred or parts of the proposal slotted for later follow-up
            * dims/jordan/stephen: whatever merges is the effective charter/policy
        * David Eads: transparency encourages reflection on author, accountability.  Antithetical to disallow that.  Privacy does happen, but removal is severe and shouldn’t be allowed to skirt accountability
        * TimP (replying to dims' suggestion that more details could be recorded privately): if decision log is kept, and private, what is the document retention policy?  Does the private log create future biases?
        * Paris (with David): What part of the current proposal is private?  Could be anything today as proposed as there is no requirement to show removal is justified.  Is the why about giving somebody satisfaction?  What next, if there is no satisfaction received?  Not necessarily about creating a public referendum gating a SC decision versus stating the reason, to which people may record their objections.
        * David: proposes a reasons template (ie: transgressions list discussed in PR)
        * Stephen: general consensus in SC to stating reasons.  Personally against transgressions.md concept
        * TimH: wants something like 100 words or less statement as described in PR.
        * TODO:
            * open discussion underway in the PR, comment publicly on the PR, share private comments if needed at steering-private list
            * Not ready for a vote; July 27th was 4 weeks from PR creation is earliest that might happen
                * Jordan: [changes.md](https://github.com/kubernetes/steering/blob/main/operations/changes.md) makes it seem like 4 weeks is the *latest* it can happen: "A vote is scheduled no later than 4 weeks after initial introduction of the change."
            * Jordan: it would be helpful to pick a target date for final content of the change being considered for a vote that is at least a few days ahead of the vote


## June 29, 2022 [Emergent Private Meeting]

**Bosun**: Christoph Blecker

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Topics**

* **Private Discussion**


## June 27, 2022 [Private Meeting]

**Bosun**: Christoph Blecker

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper) [2nd half only]

**Topics**

* **Private Discussion**


## June 6, 2022 [Public Meeting] ([recording](https://www.youtube.com/watch?v=qloe5ObSKoM))

**Bosun**: Tim Pepper

**Note taker**: Dims

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* ~~Jordan Liggitt (@liggitt)~~
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Rey Lejano
* Madhav Jivrajani

**CNCF updates?**

* Code of Conduct topics, following on after KubeCon:
    * [https://lists.cncf.io/g/cncf-toc/message/6999](https://lists.cncf.io/g/cncf-toc/message/6999)
    * [https://kccnceu2022.sched.com/event/12PIO](https://kccnceu2022.sched.com/event/12PIO) public, not recorded though
    * Also there was a private meeting between CNCF and K8s, this was streamed to some folks but not recorded
    * A CNCF working group is to be formed, details TBD
    * WG will define how a CoCC will be formed and how it will operate
* CNCF Board Town Hall this week (with K8s and other project leads)
    * Paris will be leading this
    * How GB interacts with the projects and how does it interact with other CNCF entities
* GB Elected seat for TOC is now open for a replacement for Cornelia Davis (Thanks Stephen!)
    * Deadline is tomorrow for nominations
* Paris and Stephen are working on formalizing CNCF Travel Fund especially where our existing Kubernetes Maintainers/Community Members have had need
    * Example release team members will qualify
    * Trying to get done ahead of Detroit
* [add them here]

**Votes**

*  Yay! Nothing needs voting upon :)

**Topics**

* YAY! Annual Report:
    * [https://www.cncf.io/reports/kubernetes-annual-report-2021/](https://www.cncf.io/reports/kubernetes-annual-report-2021/)
    * [https://kubernetes.io/blog/2022/06/01/annual-report-summary-2021/](https://kubernetes.io/blog/2022/06/01/annual-report-summary-2021/)
    * [tim] It is getting eyeballs, conversations are starting
    * PARIS FOR THE WIN!!!
    * [Bob] How can we make this sustainable?
    * [paris] will die on the vine if we don’t automate this and have to have liaisons pull things needed from groups.
* K8s CoCC election process to start in July…
    * [tim] Need to get the ball rolling. Need to cleanup our election process
* KubeCon observations?
    * 121 COVID test positives reported, or 1.7% of attendance
    * [tim] Quarter of staff were sick(?) on friday so there was no lunch?
    * 7500 people, it felt like a KubeCon again
    * Good presentations, good hallway track
    * Small events feel better, how can we do that again?
    * 140 people for contributor summit (possibly 110 attended)
    * Air travel was crappy for a lot of folks in this event
    * Post-event has been really quiet.
        * Sustainability / Reliability discussion was good
            * Track bugs better (triage accepted + kind bug): devstats needs updated for the new “triage accepted” label
                * Need a carrot to get people to use triage accepted
            * Get going on mentoring programs
            * Code coverage reporting on where there are gaps in testing
            * Lots of good discussion as we were not recording
        * Human aspect / turn over / minimal bandwidth common topics across the week in conversations
            * [paris] Mentoring, liaisoning with SIGs isn’t just new contributor, but the ladder for active people moving up in responsibility, growing people into more responsibility
            * [paris] How can we use the dev list better? 1-1 does not scale for folks who need to teach as well. Need to get leads to participate better.
            * [Paris] What is SIG outreach for building folks up to reviewer?
            * [Paris] Tim Hockin email on go workspaces and offer to mentor folks through action
            * [tim] management + project management to make projects sustainable
            * [stephen] flip side of this to drag something along that was not working as well, things need to be sunset as well
    * Steering members did have a dinner, most stuff talked about has already been talked about in private slack or public venues.
* Administrivia: Meeting recordings need linked to past minutes below  vvvv
* Our election needs to get started and kicked up with election officers in contribex [paris]
    * Dims, Jordan, and Bob have terms ending
    * Dims is on last term already
    * Election subproject for ContribEx: got +1, but waiting for Josh
    * Need to check with Josh in case he wants to run, Dims can run it or help now too
    * Backlog of coding needs? Probably resolved…?
* [end of meeting] any topics of interest for the next Community Meeting?
    * Annual Report Summary, retro: [https://github.com/kubernetes/steering/issues/242](https://github.com/kubernetes/steering/issues/242)
        * Need to make it sustainable (as mentioned above)
        * What can we do about sigs that had trouble writing up their PRs.
            * Need folks to throw links
            * How to follow up things
            * Bob had to rewrite paragraphs
            * Automation is needed, prompt them for exact information (mad libs!!)
            * People being burnt out, could we do monthly through the leads meeting? (work on things every month and then stack it up at the year) to reduce cognitive overload
            * Could we do this post-release? Jot down some quick notes. Could we file the templates earlier so folks can fill it in?
            * Quarterly/Monthly sig updates at the Community Meeting was hard too, but it was easier than now? Problem was so many groups canceled and contribex had to scramble to fill in the spot.
            * Opportunity for liaison should be doing, prod sigs/wgs. When the PR was filling up, liaisons helped push filling up the blanks (educators, point at direction etc). Going to sig meetings as well. Thanks Eddie!!
            * Not just DM’ing, engage the SIG, be part of them.
            * Question is how to distribute load. Where is the toolbox. More meetings is hard. How can we get folks to ping us instead of us pinging them.
            * How to encourage durable written stuff (async!) - proposals?
            * It should be just collecting existing info, but it is NOT as there are gaps.
            * Automation can be run post-release can be helpful to rollup at the end of the year
            * KEP website, kepctl how can we use it better
            * Madhav: “Something I was thinking about that echoed with Bob and Stephen’s point - End of every release, each SIG could maybe give details in largely two folds: 1) KEP updates (this will be available readily through data maintained by the enhancements team or even through the release retro) 2) a small and minimal questionnaire, and this could be 3 questions: any new subproject?, any leadership changes?, any non-kep efforts? (Based on what the info was given in this year’s annual report)”
            * So much happening, not happening in KEPs. code content is in KEPs for sure, but so many things are not (policy changes etc).
                * Madhav: “I think the efforts around project boarding on gh that enhancements team is discussing could help substantially with automation as well (in case kepctl isn’t sufficient)”
            * In-tree vs out-of-tree where most stuff is not happening in k/k + k/test-infra + k/release etc. How do we cover everyone.
            * How to reduce work and toil? Just non-kep work will need to be added.
            * Get an intern (LFX) to help with our generator, specifically for annual reports process. Need to collect a wish list. Christoph can help with mentoring (Madhav also interested to help).
            * [christoph] There is value to our report (feedback from teams/people reading it). Need to be sustainable
            * Several areas folks don’t think they are celebrated, we need to call efforts around this in shoutouts etc
            * Tailor questions based on what info we want to put to a Board (CNCF level and our member companies) when we collect the stats. We are burying the lead, that lead will get us the help. Need to re-review the questions. Also have roles for us. Jordan helped prep us (policy docs updates etc) with Paris’ help. Steering member can be “pre-prep”. We need to write down roles and task out each person on the steering to avoid things getting lost. Also shoutouts/celebrations to community members who helped but were not leads. Folks who did multiple SIGs as well. Maintainer tooling helped as well (Thanks Madhav!).
            * [Madhav] pain point for SIGs and Chairs collect data at end is a problem. Spending [bulk, trailing] time on data means other things get dropped. We could do per release thing it will help. What should the SIGs be focusing on. KEP website. How do we minimize what we miss on our data capture then we can look at non-KEP stuff which is usually more important.
            * Need targeted tools for different nuggets of info. Summary is hard.
            * Thanks RAY for jumping in with last minute edits. And sftim as well!
* [end of meeting] bosun?


## May 2, 2022	[Public Meeting] ([recording](https://www.youtube.com/watch?v=bWfeaHYEb3g))

Bosun: Paris

Note taker: Tim Pepper

Steering:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* ~~Davanum Srinivas (@dims)~~
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

Community folks

* Kendall Nelson (@diablo_rojo)
* Eddie Zaneski
* Adolfo Garcia Veytia
* Patrick Ohly
* Lachlan Evenson

**CNCF updates:**

* [paris] starting on GB resolutions this week for
    * Legal / D&O insurance, membership tiers (eg: higher tiers bring maintainers, get a discount for doing so), and more
    * Eddie: wanted to ask about the resolutions; specifically about a requirement for engineering staffing resource requirements as part of membership tiers; Paris: yes, that is one of the items I'm planning to open a draft for
* [paris] Good for kubecon? Loose ends? Group is ready.
* [paris] are we good with what's in here [https://github.com/cncf/foundation/issues/329](https://github.com/cncf/foundation/issues/329)? Where do you want to see this information end up? READMEs on committee pages? Governance.md under committees? We will close this issue once we have an artifact [https://github.com/kubernetes/funding/issues/29](https://github.com/kubernetes/funding/issues/29)
    * Boilerplate present
    * Awaiting word from lawyers (Priyanka’s chasing down info for us) to give some clarity for volunteers in a project vs employees in a company
    * Likely after KubeCon giving business ahead of it, re-ping for update after KubeCon if no update, so revisit ahead of June meeting
* Anything else from CNCF land? n/a

**Votes:**

* [bob] service catalog SIG archive
    * [https://github.com/kubernetes/community/pull/6632 ](https://github.com/kubernetes/community/pull/6632)
    * Bob: FYI CloudFoundry folks may be interested in some subprojects.  After archival, they’d likely take a fork and carry it on.  Details coming to parent issue.

**Funding:**

* **Contributor-Comms would like to spend $120 for a Buffer subscription to do timed tweets.  Which includes tweets during Kubecon if this could be authorized this week.**
* Apple dev account [https://github.com/kubernetes/funding/issues/30](https://github.com/kubernetes/funding/issues/30)
    * Final call: we would have to get CNCF to reimburse an individual as Apple doesn’t seem to have a corporate/headless option, and Apple would not consider us / CNCF a non-profit
    * Risk / reward / etc?
        * Goal would be curl’ing and running one of our binaries doesn’t pop a warning, but today already you can brew install and not have that warning
        * Info to read up at [https://developer.apple.com/support/code-signing/](https://developer.apple.com/support/code-signing/)
        * Stephen: +1 on ability to sign, -1 if only one individual can do the release then, and this means the release cannot run headless, unattended
        * AI: Stephen to get answers to the unattended / multiple user aspect
* Placeholder for SIG Windows
    * [https://github.com/kubernetes/funding/issues/31](https://github.com/kubernetes/funding/issues/31)
    * remove if they want an intern; keep if they want an experienced contractor?
    * [https://docs.google.com/document/d/1ksvxHru2Df6w23jesnpGEdZAalZfRtjCbssMAdn7BrA/edit](https://docs.google.com/document/d/1ksvxHru2Df6w23jesnpGEdZAalZfRtjCbssMAdn7BrA/edit)
    * They have high need for specific Kubernetes skill, unlikely to be satisfied by just a newcomer/intern
    * Ask SIG if they know of potential vendor/candidate to speed up selection
    * One dev for three months is the ask, but shift to kpng and in less than one release cycle, feels like it might be under scoped, underestimate
    * Any ability to decrease something in flight in the SIG to increase attention on the move?  They’ve said no
    * Should an “implementable” phase KEP be a requisite for going for funding like this?
    * AI: Jordan to reach out to networking/windows leads (e.g. Tim/Mark) to clarify where in the design phase this is, scoping/estimates (cited 3 months is less than a single Kube release cycle, which seems too short?) - followed up on the issue
* What do we want to do with this: Alibaba funding for minikube
    * [https://github.com/kubernetes/funding/issues/14](https://github.com/kubernetes/funding/issues/14)
    * Stalled in the past for cloud credits program, which is not active…?
    * Need to bump to another CNCF rep as Ihor’s currently unavailable

**Topics**

* Walk the annual report summary
    * [https://hackmd.io/VBeH9-D_QtGv8xcEydipWA](https://hackmd.io/VBeH9-D_QtGv8xcEydipWA)
    * Same question as last year: how do we handle discrepancies and no reports?
    * Ideas to streamline next year
        * per-group tracking issues early so they can dump ideas there
        * rephrase per-group docs to make it easier to copy/paste/aggregate highlights into project report
* [topics for community]


## April 18, 2022 [private meeting]

Bosun: Bob Killen

Note taker:

Steering:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**CNCF updates?**

* [paris] GB Meeting
    * Special Issues Committee
        * Goal: sustainability resolutions for August meeting
    * Legal
* [paris] kubecon loose ends?
    * AMA:
        * Maintainers track:
            * [https://kccnceu2022.sched.com/event/ytnp](https://kccnceu2022.sched.com/event/ytnp)
            * Wednesday, May 18 • 17:25 - 18:00
            * Need a 2 slide intro “who we are”
        * Contrib Summit:
            * Whoever can make it
            * Timing flexible
            * Not recorded
            * Tim might be only one who can’t attend onsite for this one
    * Travel and/or ticket assistance?
        * Paris and Stephen working on NA-forward for a better requesting process on contributor and our ends
    * Dinner placeholder sent to calendars

**Votes**

* [paris] can we vote on finding UGs another home within CNCF?
    * Bob to open issue

**Topics**

* [bob] service catalog - going to talk to cloud foundry to see if they can take it
    * Making connections
* [paris] annual report summary is still ticking along - take a peek at the progress and make early comments now [https://hackmd.io/VBeH9-D_QtGv8xcEydipWA?edit](https://hackmd.io/VBeH9-D_QtGv8xcEydipWA?edit)
    * Missing reports:
        * Node: is coming, making some last tweaks this/last week
        * Networking: is coming
    * Many of the groups need next steps / help: how do we document/make sure there are no balls dropped?
* [paris] windows wants a contractor and typed up the scope, reason for need, and why a contractor
    *
    * Next step: windows to fill out a funding issue for approval from steering, servicedesk ticket, magic (might need leads on potentially contractor entities with relevant experience)
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun?


## April 4, 2022 [Public Meeting]  ([recording](https://www.youtube.com/watch?v=ctiECKpqHXE))

**Bosun**: Bob Killen

**Note taker**: Christoph

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Arnaud Meukam
* Rey Lejano

**CNCF updates?**

* GB Meeting on the 14th [paris] - need at least 10 minutes on the deck
    * Go over legal
    * Go over sustainability deck

**Votes**

**Topics**

* [Bob] Retire SIG Service Catalog
    * [https://groups.google.com/g/kubernetes-sig-service-catalog/c/i4MqKYWhYx0/m/MO_7sNRqAgAJ](https://groups.google.com/g/kubernetes-sig-service-catalog/c/i4MqKYWhYx0/m/MO_7sNRqAgAJ)
    * [https://github.com/kubernetes-sigs/go-open-service-broker-client/issues/172](https://github.com/kubernetes-sigs/go-open-service-broker-client/issues/172)
    * [https://github.com/kubernetes-sigs/service-catalog/issues/2913](https://github.com/kubernetes-sigs/service-catalog/issues/2913)
    * [https://github.com/kubernetes-sigs/minibroker/issues/256](https://github.com/kubernetes-sigs/minibroker/issues/256)
    * [Bob] Started discussion 2021Q4 to combine with apps or retire sig
    * **AI:** send note to [dev@kubernetes.io](mailto:dev@kubernetes.io) indicating our intention to sunset
* [Bob] Annual Report Status
    * Requires updates / pending merge
        * [Bob] SIG Multicluster
            * [dims] tasha mentioned she would create a PR shortly
        * [Jordan] SIG Architecture
            * [https://github.com/kubernetes/community/pull/6551](https://github.com/kubernetes/community/pull/6551)
        * [Stephen] Reviewing this week
            * SIG UI: [https://github.com/kubernetes/community/pull/6439](https://github.com/kubernetes/community/pull/6439)
            * SIG CLI: [https://github.com/kubernetes/community/pull/6450](https://github.com/kubernetes/community/pull/6450)
            * SIG Cloud Provider: [https://github.com/kubernetes/community/pull/6489](https://github.com/kubernetes/community/pull/6489)
        * [Dims]
            * SIG-CL: [https://github.com/kubernetes/community/pull/6458](https://github.com/kubernetes/community/pull/6458)
    * Not yet open
        * [Jordan] WH Multitenancy
        * [Dims] sig-usability (expecting a PR today from tasha)
        * [Tim] way behind on life, the universe, and everything
        * Won’t land this week
            * [Arnaud] SIG K8s Infra
    * [paris] Will be working to get it done this week. Please reach out if you need help.
* [Bob] Steering AMA @ Contributor Summit EU
    * Monday, May 16th
* [Bob] Steering Session @ KubeCon EU
    * Session still lists only cblecker
        * Christoph on point w/ Cody to get other SC members listed
            * Follow up sent.
* [Bob] OpenStack PTG Gathering on the 8th @ 15 UTC / 8AM PT
    * Notes: [https://etherpad.opendev.org/p/tc-zed-ptg#L15](https://etherpad.opendev.org/p/tc-zed-ptg#L15)
    * Zoom: [https://zoom.us/j/96239912151?pwd=b2NLMTdQeThjdW0vMzZWV3RPRHVidz09#success](https://zoom.us/j/96239912151?pwd=b2NLMTdQeThjdW0vMzZWV3RPRHVidz09#success)
    * **AI:** Christoph to add to steering calendar
        * Done
* [dims] Meeting with Cloudflare later today
* [end of meeting] any topics of interest for the next Community Meeting?
* [end of meeting] bosun for next meeting?


## March 21, 2022 [Private Meeting]

**Bosun**: Christoph

**Note taker**:

**Steering Attendees**:
* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* ~~Davanum Srinivas (@dims)~~
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Votes**

* [SC Elections: COCC members should automatically be voters](https://github.com/kubernetes/steering/issues/226)
    * Follow up to clarify committees added as voters, SC will vote on the eventual PR

**Topics**

* Legal asks for CNCF/GB
    * Put in tickets and file issues
    * Approve brief
    * **AI**: All SC members to review/approve by EOD
* SC at KubeCon
* Annual reports
    * Super important to have this done by GB meeting - mid April
    * Paris prioritizing summary this week
    * Anything else?
* The new community meetings are great and we can help contribex with community wide topic ideas
    * TL separation - next meeting?
    * Terms for chairs
    * Consider moving the meeting out of “Kubernetes” (root) and into ContribEx so more people have the ability to help manually download from zoom cloud and upload to youtube.  Even if the automation is problematic, we could make it easier for humans to do the manual publish the recording closer to when it happened.
* Code dependency staffing
    * Multiple past points of discussion on go modules, json.iterator, go.yaml, yaml.sync
    * Recent need for more maintainers in etcd
    * Do we need to work to shift staffing to specific critical deps?
    * Supply chain security is a growing concern across the past year
    * Need more discussion/process around choosing to add a dep?
    * Jordan: last week we dropped ~12% of deps
    * There is at least in k/k a strong hesitance to adding deps.  Project wide though across ~300 repos, less so.
    * Ask sig-arch code organization folks if any of the tooling developed there could be used by non-k/k projects
    * Etcd staffing
        * Issue raised to steering [insert link to mailing list post]
        * TOC has been made aware of the issue as well
        * Stephen: Should etcd be part of k8s?
            * We are well positioned to better staff it
            * How do we [k8s] feel about such a large dependency that is outside of the project’s control?
        * Jordan: looped in api-machinery regarding the issue, working with api-machinery to tease apart test/bugs and features
        * Paris: forward email to kdev? Try and get more eyes on it, loop in employers


## March 7, 2022 [Public Meeting] ([recording](https://www.youtube.com/watch?v=lJ90r4lWXTU))

**Bosun**: Christoph

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Rey Lejano
* Eddie Zane
* Arnaud Meukam

**CNCF updates?**

* Ihor is a bit busy defending his country 🇺🇦 :)

**Votes**

**Topics**

* [Paris] Next CNCF Governing Board (GB) meeting is April and we need to request time on agenda mid-March
    * Need to email Executive Director Priyanka and GB chair Arun. Paris has a meeting with them end of the month
    * Sustainability is a topic we should get on their calendar
    * We have a couple of docs we need to finesse and get public
* [Dims] Evaluate CDN(s) for Kubernetes releases and other artifacts ([Issue](https://github.com/kubernetes/steering/issues/239))
    * Binaries, system packages and images. Ownership is between google and k8s-infra folks.
    * We would like to reduce cost by offloading some of these to third party folks like Cloudflare, Fastly, etc.
    * Many open source projects do this already, ASF, debian, alpine etc. (and the CDNs have had these programs for years)
    * Need someone to start conversation and figure out what this looks like so we don’t surprise with massive amounts of data
    * Dims: opened [issue](https://github.com/kubernetes/kubernetes/issues/108398) on k/k requesting stats so we have data to provide with the request
    * Arnuad: this is about more than just k/k too
    * Bob: both fastly and cloudflare will list project name as using their services. Do we need CNCF approval for use of the trademarked name in that way?
    * Dims: with stats, will craft a proposal, go through sign-offs in the project SIGs/WGs and CNCF
    * Jordan: +1 on getting stats and ack on plan from sigs (like release) before applying
    * Stephen: would it make sense to leverage existing relationship?
    * Arnaud: there is a pre-existing relationship between Fastly/LF
    * Dims: think in terms of small reversible steps (that can be rolled back if needed… xref DNS / container registry moves from the past)
    * Christoph: are there existing vendor criteria that can/should be applied here (even if the anticipated monetary cost is $0)
    * Dims: since the CDNs tout their support of OSS, they may have a sort of prospectus doc on what they’d look for from the project
    * Christoph: infra credits work from past gives some baseline from how we’d discussed this in the past with potential partners and how we’d scoped potential costs and how we should frame discussion
    * Eddie: what’s the place for requirements gathering and vendor criteria?
    * Arnaud: will start an issue to pull together req’s on trademark and marketing requirements
    * Arnaud: on scale, Alpine is multiple times larger data volume that what we’d expect for our K8s project artifacts
    * Christoph: scope…does it include what is today in community infra, or also what’s still on Google infra?  Suggest starting with community, and improved CDN with lower cost is then incentive to move remaining to community.
    * Arnaud: agree, wants to start with kops
    * Stephen: lots of prior art, plans in issues and KEPs.  kops is good initial candidate, but also need visibility to costs of the google infrastructure bits in order to plan accurately, and getting all the artifacts onto community infra would enable a lot of release work that is desired
    * Christoph: ok then need SIG K8s Infra & Release to coordinate on scope and requirements
    * Stephen: file and image promotion policy/mechanism need resolution.  Deb & RPM creation is effectively a managed service today (with Google running it for signing), so full community delivery could mean incremental work/staffing required.
* [Eddie] Block of kubecon tickets for maintainers
    * Tickets only (don’t need travel / hotel).
    * Past CNCF said apply to Diversity Scholarship, but put a note in it saying it’s not that
    * Tim: The foundation has generally felt that the individuals’ employers should pay for them, that isn’t always relevant, we need to better highlight this
    * Eddie: yes, this was a grad student active in the community but not working in the field, their employer had not reason to support their ticket cost
    * Dims: in the past we’ve successfully gone to Chris A direct and told story of specific needs and got tickets
    * Stephen: need a workflow better than texting Chris…formalize having leads roll up structured requests to steering and steering rolls up to CNCF/Chris en bulk.
    * Paris: or a form we give to lead, results route to steering.
    * Dims: let’s try it for KubeCon EU and try to refine for Kubecon NA
    * Jordan: if only we had a way to use computers to gather data…
    * [ACTION] Stephen: has older ticket still assign, will run with it, figure out surveymonkey or something
* Annual reports [15 minutes]
    * Summary working session: [https://hackmd.io/VBeH9-D_QtGv8xcEydipWA](https://hackmd.io/VBeH9-D_QtGv8xcEydipWA)
        * ID themes from your groups to bubble up
        * Do all help wanted items have issues?
        * Do all actions from the reports have issues assigned?
        * Who is helping Paris finish the draft summary to ship to editors? Stephen & Bob
        * While adding content remember there’s a Terminology end section to help explain our vocabulary to newer readers


## Feb 28, 2022 [Private Meeting] - rescheduled from Feb 21

**Bosun**: Stephen

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Votes**

*  [Christoph] SIG-Instrumentation charter update: [https://github.com/kubernetes/community/pull/6464](https://github.com/kubernetes/community/pull/6464)

**Topics**

* GB discussion
    * AI: Paris to create draft
* Note to community on Ukraine
    * AI: Bob to create draft
        * Done: [[Please Read] An Important Note Regarding Ukraine](https://groups.google.com/a/kubernetes.io/g/dev/c/swqcTp51JsI/m/tQQyqpWSAwAJ)
* SRC Private Google Drive: [https://github.com/kubernetes/steering/issues/236](https://github.com/kubernetes/steering/issues/236)
    * AI: Bob to create
        * Done
* Steering Committee Private Drive
    * AI: Bob to create
        * Done
* [Dims/Stephen] Chair/TL/maintainer bandwidth/time off
* Licensing questions
    * AI: Christoph to update GitHub management documentation to reflect licensing guidance
        * Done: [https://github.com/kubernetes/community/pull/6502](https://github.com/kubernetes/community/pull/6502)

## Feb 7, 2022 [Public Meeting] ([recording](https://www.youtube.com/watch?v=6iq4n_aKhe0))

**Bosun**: Stephen

**Note taker**:

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Alison Dowdney
* Ihor Dvoretskyi
* Kendall Nelson
* Aldo Culquicondor

**CNCF updates?**

* n/a from Ihor

**Votes**

* WG Batch vote: [https://github.com/kubernetes/community/pull/6299](https://github.com/kubernetes/community/pull/6299)
    * Technically awaiting node’s ACK, but looks to be roughly there
    * SIG liaison?  Let’s see after Node’s requested clarification if this ends up closer to Autoscaling, Apps, Scheduling, Node, etc., then decide

**Topics**

* Steering operational changes / discussion (prioritize, bumped twice)
    * Roles
        * ID more: [https://github.com/kubernetes/steering/issues/225](https://github.com/kubernetes/steering/issues/225)
        * Liaison [https://github.com/kubernetes/steering/issues/228](https://github.com/kubernetes/steering/issues/228)
    * Triage [stephen]
        * A few things I’ve noticed while onboarding:
            * We’re not effectively using the project board (while resolved now, many of the open items for Steering were not listed on the project board)
            * Stale milestones
            * Several Steering issues lacking assignees
                * I’ve marked a few of these w/ /help, but if it’s on the board+triaged, we should assign a point person, even if it’s something that eventually needs to be taken to vote
        * [Paris] Do we need an offsite or all day meeting to get through back topics? ([from Slack](https://kubernetes.slack.com/archives/CPNFRNLTS/p1638484593106500?thread_ts=1638466730.104300&cid=CPNFRNLTS))
            * [Stephen] Let's see if we can spin a lightweight triage process before EOY, and then async a bunch of stuff. If we find there are still things stuck in the hopper, then maybe a triage sesh in the new year?
            * [Bob] one-time initial scrub and align, then shift to sustaining async
            * Should align this with election transition, which doesn’t map to normal KubeCon face-to-face.  Future KubeCon’s do make sense for this.  Need a catch up sooner too.
        * [Paris] Do we need a decision making process for Steering?
            * How do we know when something is up for a vote? Or when something has had a decent SLA to activity? How do we know when to put something on the agenda?  Do community members know when to put something on the agenda?  Does the boson of each meeting drive the agenda and prioritization, or do we have a shared sense of it?
            * [Stephen] Add a column to the project board for “needs discussion in meeting”
            * [Christoph] self-organizing has effectively spread the agenda/prioritizing focus
            * Fewer concerns about abuse of bosun role
            * [Bob] Items that might not have a liaison or person assigned yet should be assigned an owner to at least take them to the next step (liaison assigned etc)
* [TimP] [Election “Committee”](https://github.com/kubernetes/community/tree/master/events/elections):
    * [https://github.com/kubernetes/community/pull/6243](https://github.com/kubernetes/community/pull/6243)
    * Streamlining CoCC and Steering Committee conflict of interest docs bumped into additional conflict (steering candidates shouldn’t be on election body) and the non-charted but so-named election committee.  Should Election Committee be a (short term) chartered entity as a “Committee”
    * Currently elections are run out of Contributor Experience, with officers approved by Steering Committee, with docs split across Steering, Community, and ContribEx.  Consolidate to k/community/events/elections
    * Should elections be a “Committee” or a subproject of ContribEx?  [https://github.com/kubernetes/community/issues/6084](https://github.com/kubernetes/community/issues/6084)
    * Is the body holding non-public information, which would mean our rules require them to be a Committee?  No.  Current mechanism is well anonymized.  Fits as a subproject.
    * TimP will update PRs to shift this to k/community, deduplicate, call elections a subproject not committee
* [Christoph] [KCEU22: Steering Committee Session](https://github.com/kubernetes/steering/issues/237)
    * [Dims] Put something on the calendar
    * [Christoph] To submit session, TBD Speakers
* [paris] have you talked to your groups about annual reports?
    * Are there any pivots we need to do right now?
    * Should be doable to get draft PRs this/next week and mergeable by end of month
* March Boson:


## Jan 24, 2022 [Private Meeting]

**Bosun**: Jordan

**Note taker**:

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

* FYI: wg-batch PR progressing
* Annual report next steps
* Owners files and promoting people
* workflow for CNCF legal consulations

**Action items**

* All: RSVP for CoCC sync dates/times
* Annual reports
    * cblecker to run and check in files
    * paris to open follow-up issue for generator work - [https://github.com/kubernetes/community/issues/6357](https://github.com/kubernetes/community/issues/6357)
* All: read liaison responsibility doc (especially w.r.t. annual report responsibilities)
* ?: add password to private zoom meeting
* ?: add location for private meeting notes

## Jan 10, 2022 [Public Meeting] ([recording](https://www.youtube.com/watch?v=D78xlu8z-2k&))

**Bosun**: Jordan

**Note taker**: Dims

**Steering Attendees**:

* Bob Killen (@mrbobbytables)
* Christoph Blecker (@cblecker)
* Davanum Srinivas (@dims)
* Jordan Liggitt (@liggitt)
* Paris Pittman (@parispittman)
* Stephen Augustus (@justaugustus)
* Tim Pepper (@tpepper)

**Community Attendees**:

* Alison Dowdney (@alisondy)
* Kendall Nelson (@diablo_rojo)
* Aldo Culquicondor (@alculquicondor)
* Micah Hausler (@micahhausler)
* Ihor Dvoretskyi (@idvoretskyi)
* Eddie Zaneski (@eddiezane)
* Arnaud Meukam (@ameukam)

**CNCF updates?**

* 0:00:46 [dims] Maintainer Track session email for Kubecon EU is out.
    * Possibly in the leads meeting tomorrow (stephen)
    * Liaisons should reach out their sigs as well
* [add them here]

**Votes**

**Topics**

* [0:02:00](https://www.youtube.com/watch?v=D78xlu8z-2k&t=2m): Steering/CoCC housekeeping
    * AI: Liaison for CoCC Tim Pepper to ping to set a date for recurring quarterly sync.
    * [tpepper] Close on [conflict-of-interest documentation](https://github.com/kubernetes/steering/pull/224#issuecomment-990664183) location: see end of agenda today, broader topic needing discussion beyond CoCC
* [0:03:56](https://www.youtube.com/watch?v=D78xlu8z-2k&t=3m56s): Security Response Committee [micahhausler]
    * FYI, we intend to expand the PSC by 2 or 3 beyond the [minimum 7](https://github.com/kubernetes/committee-security-response/blob/main/security-release-process.md#joining). Key motivations are to add additional folks in the oncall rotation, and to handle the amount of reports coming in.
        * [micah] Increase size to be able add people over time (to rotate folks out)
        * [stephen] is there a promotion process already in place?
            * [micah] 3 month for an associate. They handle things that are already public. (see “[membership](https://github.com/kubernetes/committee-security-response/blob/main/security-release-process.md#security-response-committee-membership)”)
        * [christoph] sustainability, is there a good pipeline for incoming associates. Fairly trusted role in the community.
        * [micah] open up more slots for associates to build the pipeline and promoting those who are already associates to the full member.
        * [micah] will drop a note to steering committee mailing list about PSC accountability etc. (no open meetings etc), we could expose more metrics etc without compromising issues.
        * [jordan] yes please, similar to what we did with CoCC
        * [paris] +1 we could do annual reports
        * [dims] see [https://blogs.apache.org/foundation/entry/apache-software-foundation-security-report2](https://blogs.apache.org/foundation/entry/apache-software-foundation-security-report2) from today
* [0:10:28](https://www.youtube.com/watch?v=D78xlu8z-2k&t=10m28s): Annual report prep
    1. [liggitt] Updated [target draft/publication dates](https://github.com/kubernetes/community/pull/6322)
    2. [liggitt] Update [SIG/WG templates](https://github.com/kubernetes/community/pull/6301) (need lgtm from steering)
        * [liggitt] steering folks, please review, targeting completion by Jan 12
    * [cblecker] Status of template generator
        * [cblecker] Generator is from last year, have a local branch with some changes, need to write tests later. Will open a PR in next few days. Need to merge in the WG template.
        * [https://github.com/kubernetes/community/pull/5514](https://github.com/kubernetes/community/pull/5514)
        * [liggitt] we can announce in leads to build awareness that this is coming. Liaisons too.
    * Next step once templates are generated:
        * Comms - announce to dev@kubernetes.io, etc
            * Paris can help and engage upstream marketing team
            * start in January after templates are generated
        * Outreach/reminders to individual group leads - liaisons
            * start in January after templates are generated
            * Liaison doc: [https://hackmd.io/fu_S0ASnRAGqUDMyaH_-ag](https://hackmd.io/fu_S0ASnRAGqUDMyaH_-ag)
                * **Steering AI: This is first draft; needs a good review **
            * [paris] we can do this offline, please read and leave comments on the doc. Pretty basic to kick us off. Left wild ideas for later! :)
* [0:15:00](https://www.youtube.com/watch?v=D78xlu8z-2k&t=15m): WG Batch charter discussion
    * [https://github.com/kubernetes/community/pull/6299](https://github.com/kubernetes/community/pull/6299)
    * [abdullah/aldo] for quite a while requirements for batch related things were coming to sig-scheduling, looked like we needed a WG as a lot of things cross multiple sigs. Improve batch experience in core kubernetes. Focus on Job API, job queueing, work better with schedule and kube controller manager. Reduces duplication, example everyone is rewriting Job API. provide hooks needed to build on things here that eco system can depend on. scheduling/apps/node etc will be involved. WG is more focused than sig-scheduling.
    * [stephen] how close are we picking organizers
    * [abdullah] pretty close.. Trying to engage apps and node for sure. Need folks from others and not just scheduling. IBM, RedHat, Google are represented in leadership.
    * [stephen] co-chairs don’t need to be already leads.
    * [dims] how is the community participation from other folks in CNCF being thought of. How can we bridge and not end up splitting efforts
        * ahg: we want to build core primitives/features to support external batch integrations
        * jordan: would be good to intentionally seek out feedback from external batch integrators (+1 bob)
        * [tpepper]
            * ensure leaders of this working group share that vision about bridging built-in/external groups and are committed to doing the work to make that happen
            * vision/roadmap in PR is vague, which is normal to start, but would like to see that crisp up early
            * ahg: could go in a few directions, depending on whether code/components are produced by this working group and don't slot neatly under a single sig
                * jordan: would like to see investigating this as an explicit step in the working group (basically answering the "what is the roadmap?" question from the annual report)
                * tpepper: would be useful to sketch out possible outcomes, even if they're not committed directions yet, milestone checkpoint for ~1year in to re-evaluate & update
        * [christoph] current charter is “normal”, concerned about engaging folks outside of core k8s properly. sensitive to enabling ecosystem instead of "king-making" and building in "one true way". ensure working group is engaged with sponsoring sigs to make ownership of artifacts/definitions/code produced clear
            * [tpepper] add to the doc to describe there are N ways of doing things which are quite similar in the abstract…the goal is not to king make but to distill out of these N similar ways, a common way.  This is implied in the “want to unify the way”, but it’s not described so clearly as a “distill from multiple similar into a quality common alternative” words
            * ahg: want to see unification of disparate approaches where it makes sense. batch ecosystem is very fragmented today, to the point where the ecosystem suffers
                * christoph: fragmentation isn’t necessarily a bad thing on its own. Ecosystem enablement is important. Defining what problems fragmentation is causing should be important to getting buy in from stakeholders
            * jordan: vision makes sense, but is good to be explicit in the charter
        * [stephen] layout a rough draft and make the first action to be refining that. Not worried about king making. More concerned about this at the CNCF level.
        * [dims] what do we need to do to re-engage the folks that went off and did different things?
            * ahg: think the working group is the answer
        * [tpepper] re: leads selection; will take real work to engage folks with external batch solutions
        * [paris]: (said in chat) id like to see a kubecon update chartered in so make sure there is a solid bridge of communication with the wider community
            * [dims] +100
            * [https://github.com/kubernetes/community/pull/6299/files/9ad2b608cfd47ac271d734a0afb44845996250b1#r781431751](https://github.com/kubernetes/community/pull/6299/files/9ad2b608cfd47ac271d734a0afb44845996250b1#r781431751)
        * ordered AIs:
            * wg proposers: incorporate feedback into PR, update proposed leads
            * sponsoring sigs: ack final state of PR
            * steering: assign liaison, vote
* [0:43:20](https://www.youtube.com/watch?v=D78xlu8z-2k&t=43m20s): K8s Infra [ameukam]: ProwJobs migration
    * [https://github.com/kubernetes/k8s.io/issues/1469](https://github.com/kubernetes/k8s.io/issues/1469)
    * [ameukam]: Progress is slow. Most SIGs don’t make it a priority. I would like to use the SC liaisons to request help from the different SIG chairs to migrate away the jobs their SIGs are responsible for to the community-owned infrastructure.
    * [stephen]: can the work be sorted by responsible sig?
    * [jordan]: individual issues per sig with more specific actions needed is usually more effective
    * [christoph] for k8s-infra, need to be *very* specific what you are asking sigs to do. don't assume sigs will know what jobs they have or which ones are not running on community infra. sig chair may or may not know which jobs need work. [+1 from bob]
    * [tpepper] coordinating point can be helpful
    * [paris] if having trouble getting in contact with other sigs, that's helpful to know
    * [stephen] k8s-infra is doing a fair amount of documentation, but getting the information in the right place and making it clearer who the owner of each piece is would be good
    * [dims] question: what is the headroom in the budget for the new jobs this is asking to add to community infrastructure?
        * ameukam: we don't really know until we complete the migration
        * dims: we need to migrate gradually so we don't move everything and blow stuff up
* _— topics below here bumped for time_
* Steering operational changes / discussion (prioritize for next public meeting)
    * Roles
        * ID more: [https://github.com/kubernetes/steering/issues/225](https://github.com/kubernetes/steering/issues/225)
        * Liaison [https://github.com/kubernetes/steering/issues/228](https://github.com/kubernetes/steering/issues/228)
    * Triage [stephen]
        * A few things I’ve noticed while onboarding:
            * We’re not effectively using the project board (while resolved now, many of the open items for Steering were not listed on the project board)
            * Stale milestones
            * Several Steering issues lacking assignees
                * I’ve marked a few of these w/ /help, but if it’s on the board+triaged, we should assign a point person, even if it’s something that eventually needs to be taken to vote
        * [Paris] Do we need an offsite or all day meeting to get through back topics? ([from Slack](https://kubernetes.slack.com/archives/CPNFRNLTS/p1638484593106500?thread_ts=1638466730.104300&cid=CPNFRNLTS))
            * [Stephen] Let's see if we can spin a lightweight triage process before EOY, and then async a bunch of stuff. If we find there are still things stuck in the hopper, then maybe a triage sesh in the new year?
        * [Paris] Do we need a decision making process for Steering?
            * How do we know when something is up for a vote? Or when something has had a decent SLA to activity? How do we know when to put something on the agenda?
* [Tim] [Election “Committee”](https://github.com/kubernetes/community/tree/master/events/elections): no time today, will shift to [https://github.com/kubernetes/community/pull/6243](https://github.com/kubernetes/community/pull/6243) async discussion
    * bumped to slack
