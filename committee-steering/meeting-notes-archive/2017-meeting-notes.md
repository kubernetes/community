## December 7, 2017

* **Bosun**: Brandon Philips [CoreOS] (@brandonphilips)
* **Note Taker**: Brandon Philips [CoreOS] (@brandonphilips)
* **Attendees**
    * Aaron Crickenberger [Samsung SDS]
    * Brandon Philips [CoreOS]
    * Brendan Burns [Microsoft]
    * Brian Grant [Google]
    * Clayton Coleman [Redhat]
    * Derek Carr [Red Hat]
    * Joe Beda [Heptio]
    * Phillip Wittrock [Google]
    * Michelle Noorali [Microsoft]
    * Sarah Novotny [Google]
    * Timothy St Clair [Heptio]
    * Tim Hockin [Google]
* **Absent**
    * Quinton Hoole	[Huawei]
* **Topics**
    * Discussion around steering committee velocity and breaking into smaller teams
    * **Proposal**: Do weekly standups and break into 2-3 teams that tackle 2-3 issues a month
    * [Draft values discussion](https://docs.google.com/document/d/12J0YLKBOfBTMd8ZdPpFMIAhMmGIfapMg-nkYPJYY7FY/edit) - everyone says it looks good

**Iteration Teams**

* **Values**: Sarah, Brendan (everyone LGTM)
* **Owners**: Brian, Aaron
* **SIG Interfaces: **Phillip, Joe, Michelle
* **Charter Cleanup: **Clayton, Tim, Derek (amendment language)

**Actions**

* Philips: update the backlog.md to show the sprint [https://github.com/kubernetes/steering/pull/13](https://github.com/kubernetes/steering/pull/13)

## November 8, 2017

* **Bosun**: Aaron Crickenberger [Samsung SDS] (@spiffxp)
* **Note Taker**: Sarah Novotny [Google] (@sarahnovotny)
* **Attendees**
    * Aaron Crickenberger [Samsung SDS]
    * Sarah Novotny [Google]
    * Tim Hockin {Google}
    * Brian Grant [Google]
    * Joe Beda [Heptio]
    * Michelle Noorali [Microsoft]
    * Derek Carr [Red Hat]
    * Brendan Burns [Microsoft]
    * Timothy St Clair [Heptio]
    * Clayton Coleman [Redhat]
* **Topics**
    * Administrivia
        * Public meeting notes: this doc? Separate doc? .md in k/steering repo?
            * Read only publishing of the steering committee agenda
            * AI -- Aaron to read before pushing publish to kubernetes-dev@
        * Public meeting recording: where do we want to post these?
            * Kubernetes YouTube Channel -- @jorge can help create a playlist
            * AI -- Joe to reach out to Jorge
        * [Do we need to ratify Michelle as CNCF GB rep with a vote?](https://github.com/kubernetes/steering/issues/9)
            * Nope -- announced on twitter so it must be true.
        * Next Bosun: Any volunteers? If we’re still going alphabetically by first name (thanks for that), it’s Brandon
            * Brandon next Bosun in either December or January.
        * December meetings
            * Still an open question.
    * Action Item followup
        * [steering@k8s.io](mailto:steering@k8s.io), [steering-private@k8s.io](mailto:steering-private@k8s.io) have been created, thanks Tim
            * Need final clarity on subscribe-ability of steering@
                * Yes, subscribable.  But, committee reserves the right to move conversations to other lists.
            * AI -- Tim steering-private Publicly postable (done)
        * Do we have a document that describes when/how to use them? Consensus?
        * k/steering has a CONTRIBUTING.md, thanks Joe
            * AI -- michelle make PR
        * Discuss the charter: are we ready to ratify?
            * Sounded like Tim SC had some thoughts
            * Strip the existing doc into mission, scope, outline
            * Move the election content into another doc.
            * Must improve the MoS definition.
            * Propose a PR and get LGTMs from each committee member?
            * Committee’s responsibility for a charter.  Comments from the community will be considered, but not necessarily addressed.
            * Scope: Streamline charter
                * Remove bootstrap language
                * Defer MoS update
                * **Add Amendment process**
                    * ⅔ of the steering committee and one week notice
                    * Lazy consensus for everyday biz
                * AI -- Derek -- Amendment language
                * AI -- Tim St Clair -- streamline charter.
    * Incubation
        * Where do we stand on this?
            * Review of incubator projects -- in/out?
            * What is the architectural vision of the kubernetes project?
            * Endorsement as a byproduct of graduation
            * Ecosystem
            * Brendan
                * Sig sponsored --  scratch space - free form
                * Incubator bestows “officialdom” / cla -- get rid of that
                * Started on the outside do we want to bring it in? .
            * Joe
                * Another category
                    * Graduated already and don’t fit the new devs.
                * \
            * Aaron
                * 61 repos in k/
                * Brian doesn’t think we have enough
                * Derek doesnt’ think they’re well enough bounded
            * Tim St Clair
                * Sunset period for projects which should exit the k/k space
                * Potential org or set of orgs for these adjacent projects
            * Brendan
                * Sunsetting isn’t bad, but hard to make progress.
            * Clayton --
                * If a sig decides to adopt something, who is a SIG? membership/vote.
                * If sig node wants to put another container runtime into k/k  are they able to?
                    * New project, yes.
                    * Existing project, no
            * Brendan
                * Saad is doing this with snapshopts… not in core.
            * Clayton
                * Prototype repos under SIG-CLI
            * Brian
                * No clear notion of decider for sig
            * Derek
                * Tied to owners in kube
            * Joe
                * That’s not the discussion we are having.. .but agree should be related to owners.
            * Brendan
                * SIG as owners… not k/k
                * k/k can be draw in pieces from sigs
            * Aaron
                * Is Incubator the same as SIG scratch space?
            * Brendan
                * Contrib was a BAD
                * Each SIG breaks the connection of kubernetes
            * Clayton
                * Sigs should be more independent
                * Loose model works well for prototyping
            * Derek
                * What is the expectation of testing services
            * Aaron
                * Consistent tooling from kubernetes? Orthogonal to our incubation process
                * Need to unblock people who are trying to put projects in k/k and or incubator.
            * Brendan
                * Don’t want to start in a new org == blessing
                * Kubernetes-sig-cli org etc.
                * Playground
            * Aaron
                * What is the architectural vision?
                * What is the k8s project?
                * What repos fit that vision?
                * What is the mapping from sig &lt;-> repo?
                * A few people from this group need to work offline and discover what the boundaries might be.
            * Brendan
                * Not trying to retro organize
                * Trying to give space to explore without blessing everyhing
            * Joe
                * Do need to go through the exiting repos
                * Fairness for incubators
            * Brian
                * We’re rehashing
                * Review the docs.  We’re mostly in agreement
            * Derek
                * Merit looking at incubator repos
                * Sig vs playground?
            * Brian
                * Eventually want every piece of code to be assigned to a sig.
                * Other than k/k k/docs …. Are cross project mono repos
            * Michelle
                * Org for each sig?
            * Brian
                * Api-machinery will need multiple orgss
                * Main driver for more orgs == permissions
                * Overhead of admin is messy either way
            * AI -- Brendan write a document to walk through incubator list and find patterns of in/out/ bubble
            * AI -- Brian write a document for what to bring to the architecture for new repos
            * Brian
                * Unblock organic growth
                * SIG org veto was around no decision process inside the sigs.
        * Matt’s doc: [https://docs.google.com/document/d/1DjPSKe01a4BLi2cxUENTbQZGmANhumt8LQVz03wiiAw/edit#](https://docs.google.com/document/d/1DjPSKe01a4BLi2cxUENTbQZGmANhumt8LQVz03wiiAw/edit#)
        * Phil’s doc: [https://docs.google.com/document/d/1_FBm9zR4vU2X3-j-jSGlDPIkXweQbc_x6Fc6Ph0OupQ/edit](https://docs.google.com/document/d/1_FBm9zR4vU2X3-j-jSGlDPIkXweQbc_x6Fc6Ph0OupQ/edit)
    * Kubecon
        * We have one more f2f meeting prior to Kubecon, is there anything we need to prioritize?
* **Action Items**

## October 25, 2017 -

* **Bosun**: Aaron Crickenberger [Samsung SDS] (@spiffxp)
* **Note Taker**: Brandon Philips [CoreOS] (@philips)
* **Attendees**
    * Joe Beda, Tim St. Clair [Heptio]
    * Aaron Crickenberger, [Samsung SDS]
    * Derek Carr, Clayton Coleman [Red Hat]
    * Phillip Wittrock, Sarah Novotny, Brian Grant, Tim Hockin [Google]
    * Brandon Philips [CoreOS] NOTE TAKER!
    * Michelle Noorali, Brendan Burns [Microsoft]
    * Quinton Hoole [Huawei]
* **Topics**
    * Administrivia
        * How do people contact us?
            * Today: p2p or e-mail [steering@k8s.io](mailto:steering@k8s.io)
            * Ideas: e-mail [steering-public@k8s.io](mailto:steering-public@k8s.io), #steering-committee slack, issues on kubernetes/steering
            * **Consensus**: See AIs below
        * How do we respond?
            * Today: “we’re very busy, it’s in our backlog”
            * **Consensus**: use the list
        * How do we keep track of our ongoing backlog / priorities?
            * Today: we have backlog.md in kubernetes/steering
            * Ideas: this working doc, a spreadsheet, issues in kubernetes/steering, regular updates to backlog.md
            * **Consensus:** using a backlog.md file and steering folks should just merge their own stuff. PR workflow reserved for CoC or something. Joe will write-up something.
        * Do we record meetings?
            * Today: we aren’t
            * Ideas: what does our charter say? tick-tock a public recording, and a private meeting with no recording?
            * **Consensus:** start meetings as not recording, confirm with the group before starting the recording early in the meeting
        * What do we make public?
            * Today: the contents of kubernetes/steering
            * Ideas: what does our charter say? Public mailing list (separate from kubernetes-dev), public meeting notes (follow same tick-tock?), github issues etc
            * **Consensus: **Default to public on recording and mailing list
        * What is quorum for us?
            * Meeting quorum?
            * Decision / binding quorum?
            * **Consensus: **Half + 1 on meetings to start, 2/3rd on votes for binding stuff. Consensus is the goal but not requirement. Future: modified [lazy consensus](https://docs.google.com/document/d/1fRYZYQlCGv4ebMqwZE8BRzPbuGl8ElyaNr5CuEmaRyk/edit).
    * Top 3 priorities ([spreadsheet](https://docs.google.com/spreadsheets/d/1Ym3gXY7I1gs_RSqCfzeg7BlCZk0THAYc-jvVrLe0vDo/edit#gid=0))
        * Incubator
            * What does being a Kubernetes project mean?
            * Policy of accepting new code -> means it is a CNCF project
            * Management of the GitHub orgs
            * Decision process for how and when a project qualifies
                * Problem: no clear way to reject
* What is motivating people to be an incubator project?
    * Legal ownership of code
    * Process support (bots, release, etc)
    * Official blessing / endorsement
    * "Neutral ownership"
    * Visibility
    * Bundling into releases
* What are we afraid of?
    * Transferring legal ownership of code after it is produced
* Why do we feel this is urgent?
    * Incubation projects coming up for being kicked out after 1yr
    * No process is blocking to grandparent in old code
        * Code of conduct
        * SIG Charters
        * Honorable mention: charter ratification
            * Do we have an amendment process for the doc?
            * [Charter](https://github.com/kubernetes/steering/blob/master/charter.md)
* Discussion around ownership and SIGs
    * Misc questions ([spiffxp] we didn’t have time to get to these)
        * [quinton] [Measuring the value of Kubernetes contributions]( https://groups.google.com/forum/#!msg/kubernetes-wg-contribex/UtfHQCvaKw4/qVu0yt4cBQAJ) - there’s a thread about this on sig-contribex, and I’ve offered to draft a proposal for review.  Does the steering committee have any early input/opinions?
        * [spiffxp] clarifying community-membership ladder
            * These are coming from “we want a contributor mentorship program”
            * Is the reviewers / approvers / maintainers distinction useful across all repos?
            * are the numbers for each of these roles set in stone?
* **Action Items**
    * **Tim Hockin: **own [steering@k8s.io](mailto:steering@k8s.io) and [steering-private@k8s.io](mailto:steering-private@k8s.io), rename existing list to -private, create a new steering that is world writeable and world readable but cannot be subscribed to
    * ???: Document when we use each list: regular business happens on steering list and the steering-private is postable only by the team and how to engage with the steering committee
    * **Joe Beda:** write a steering.git editing guideline -- done
    * **Brandon Philips: **Send charter discussion to mailing list
    * **All: **Discuss the charter on the public mailing list and get to consensus on ratification

## October 18, 2017 -

* **Bosun**:  404
* **Note Taker**: Joe Beda - Heptio
* **Attendees**
    * Joe Beda, Tim St. Clair - Heptio
    * Sarah Novotny - Google
    * Aaron Crickenberger - Samsung SDS
    * Brian Grant - Google
    * Phillip Wittrock - Google
    * Tim Hockin - Google
    * Michelle Noorali - Microsoft
    * Brandon Philips - CoreOS
* **Topics**
    * [sarahnovotny] Contributor Summit invitation selection
        * Thread on selection.
        * Need to pick something now.
        * Options:
            * 1. SIG leads and WG leads + SIG nominations + lottery
            * 2. SIG leads and WG leads + lottery
            * 3. Lottery
                * a) from those that voted in election (300 people)
                * b) members of standing
            * **4. SIG leads WG leads + editorial quota + MoS lottery**
        * Is it deliberately not on conf agenda? Yes -- because it is invite only.
        * 150 people
        * Contributor summit last time:
            * SIG leads + approvers + 5 top companies to nominate people + lottery
        * Allow transfers? No -- go to the next picked person, etc
        * What are purposes? Can we make decisions?
            * Mostly about getting on the same page and motivation
            * Direct effort to the right areas
            * Review what happened since leadership summit
        * Joe is a broken record about his concerns around SIG leadership rights and responsibilities
        * Use owners file? Already used for “members of standing”
            * Quinton: We need to be careful here that we don’t use things created for one purpose for another purpose and offend folks.
            * Tim: We were super lenient with member of standing list. May want to curate more for contrib summit.
        * Quinton: do the least offensive thing and fix this going forward?
        * Sarah: all of this is great but we need to get invites out ASAP
        * Brian: get interest and then lottery?
            * Tim/Tim: not enough time
        * Brian: Pure lottery seems less than useful.  We have
    * [sarahnovotny] CNCF Project representation to the board
        * This is the Governing Board -- business/marketing/serving
        * Who from steering committee doesn’t want to be on board?
            * Tim StC, Joe, Clayton, Derek, Quinton
            * Brian, Sarah - have to deal with it already
            * TimH - will do it if I have to
            * Brandon - already there
            * Left: Aaron, **Michelle**, Phil
    * [jbeda] Curation of top things to address?
        * Perhaps have everyone list their top three priorities and we can see what bubbles to the top?
            * Project incubation
                * Graduation
            * CoC group
        * [Backlog](https://github.com/kubernetes/steering/blob/master/backlog.md) for reference.
        * Offline discussion
            * Items in spreadsheet by fri
            * Pull together agenda by tuesday
    * [spiffxp] test-infra cluster funded by cncf?
        * Current cluster under a google.com project, can’t add outside collaborators
        * Want to run just the core test-infra stack (eg: prow, submit-queue, etc) on a cluster that non-googlers can help support
        * Curious if this is the right place to handle this (or falls under top priority / backlog above)
        * Notes:
            * Aaron wants to see all tests to be run by non-googlers
            * Tim: there is a GCP project but there is question on who/how it gets paid.
            * Things that directly benefit cloud providers (GKE testing, GPU) should be funded by those providers
            * Tim and Aaron to take this offline.
    * [pwittroc] where do subprojects store container images & bits?
        * Is this something we talk about today or later w/ backlog.  Related to the test-infra item above
        * GCR - more vanity name -- k8s.gcr.io will be a thing
        * Can put any subdirectories we want in there.  Can categorize.
        * Easy transition outside of google
        * [Brandon] Circling around some central tenants
            * Domain names, hosting and billing needs to be owned by k8s project
    * Questions to answer:
        * Do we record meetings?
            * Quinton/Aaron: I’m in favor of public but there is a lot of sensitive discussion that we need a safe space
            * Brian: Public office hours?
        * Do we have a “chair” or work via consensus?
            * Monthly or quarterly rotation of speaker/whip
            * Name:** Boatswain**
            * Frequency:** Monthly**
            * First Boatswain: **Aaron**
        * What is open, what is closed?
            * Brian: we should have an open email list for public business
* **Action items**
    * Joe to pull together spreadsheet for folks to list their 3 top things from backlog and send mail to steering committee
