Contributor summit - KubeCon/CloudNativeCon 2017

**@AUTHORS - CONNOR DOYLE**

**@SLIDE AUTHORS ARE THE PRESENTERS**

**@SKETCHNOTES AUTHOR DAN ROMLEIN**

# 2018 features and roadmap update

Presenters: Jaice, Aparna, Ihor, Craig, Caleb

Slidedeck: https://docs.google.com/presentation/d/10AcxtnYFT9Btg_oTV4yWGNRZy41BKK9OjK3rjwtje0g/edit?usp=sharing 

What is SIG PM: the "periscope" of Kubernetes

* They look out for what’s next, translate what’s going on in the community to what that means for Kubernetes

* Responsible for Roadmap and Features Process

Understanding the release notes is difficult sometimes since the docs aren't always done by the time the release team is compiling the notes.

## Retrospective on 2017

* What did we do since Seattle

* Feedback on how we can enhance SIG-PM

* Moving from "product-driven" to "project-driven" where SIGs are defining their roadmaps based on market, end-use input etc.

Major 2017 Features:

* (Apps) Workloads API to Stable

* (Scalability) 5k nodes in a single cluster

* (Networking) NetworkPolicy

* (Node) CRI, GPU support

* (Auth) RBAC stable

* (Cloud providers) Out of core

* (Cluster Lifecycle) kubeadm enhancements - road to GA in 2018?

* (Autoscaler & instrumentation): custom metrics for HPA

* (Storage) CSI (Container Storage Interface)

How did we do on the contributor roadmap?

* [See 2017 roadmap slides](https://docs.google.com/presentation/d/1GkDV6suQ3IhDUIMLtOown5DidD5uZuec4w6H8ve7XVo/edit)

Audience Feedback:

* Liked that it felt like there was a common vision or theme last year.

* Liked that there was a PM rep saying "do docs", "your tests are failing" etc

* Leadership summit 6 months ago: shot heard called was "stability releases, stability releases".  Not sure that 1.8 was really a stability release, not sure 1.9 is either.  Will 2018 be the year of stability (on the desktop)

    * (Brian Grant) Come to the feature branch session!

    * Notion of stability needs to be captured as a feature or roadmap item to talk and brag about.  Quantify, talk about as roadmap item

    * Idea for 2018: how do you measure and gamify stability?  See a number in the red, people will want to fix.  Testing, code quality, other metrics - might improve stability organically

    * Context of stability and testing: achievement was conformance program. 30+ conformant distros!

    * Want to see project continue to be useful: within your SIG, invest in conformance, extending suite. Going back to what is and is not k8s - define core, extensions: don't compromise the stability of the core.

    * Please come see Eric Tune and define stability : 

        * "cluster crashed"

        * "too many features"

        * "an API I was using stopped working"

        * "community is chaotic, how do I navigate that"

* There are many new alpha features in each new release. Please prioritize graduating and stabilizing the existing features.  (More than 50 APIs already)

* Looking for volunteers on writing a stability proposal?

* Jaice has one already!

    * *May* have broken the comment limit on Google Docs 

    * Need to define lens:

        * architecture, community etc; look at a proposal under each. 

        * Brian is working on arch stability.

        * Contribex is looking at mentorship and ladder.

        * Myriad of ways to approach problem set.  How do we mature the processes of the Kubernetes ecosystem?

* Looking for co-authors?

## Proposals and KEPs

(Caleb)

Please hang out in SIG-Arch, SIG-Contribex, SIG-PM to drive this process forward

Looking at "is this feature ready to move to the next stability milestone?" (Alpha to Beta, Beta to GA etc)

* Proposals are now **K**ubernetes **E**nhancement **P**roposals

    * Piece-by-piece over on more multiple releases (living documents)

    * Looked at a lot of other open source projects, e.g. the Rust community (RFC Process)

    * designed in the era of GitHub; decided on a lightweight process that works well with the VCS. 

* Talk about what we want to do without a long spec doc, about what we agreed to ahead of time, but* don't want to diverge 2 years later*.

* Helps tracking individual features (easier to read release notes)

    * Release note writing take tracking down a lot of docs, GitHUb issues, design docs, Slack and Google Group comments; combine from a bunch of places

    * Hard to tell from the release notes what's important and what's a minor detail.

* Every SIG should set their own roadmap - the KEP proposal enables that.

* Template that asks you to think ahead for the lifecycle of the feature; let people know what you're planning to do.  

    * It's a medium for comms; not saying "It has to be done this way" but saying why this is important. 

    * Inspired by "[Toward Go 2](https://blog.golang.org/toward-go2)" blog post by rsc

* Has been tested - [draft KEP for Cloud Providers](https://github.com/kubernetes/community/pull/1455), Tim St. Clair has tested.

* Want to make easier for new contributors to write KEPs.

* Starting with "what is a unit of work, how do people care"

Questions: 

* Are KEPS google docs, or pull requests, etc?  How do you submit one?

    * Original intend: something that lives in source control.  Discoverable like any part of the code.  Attempt to combine design proposals and umbrella GitHub issues, link to 10s of other issues.  They will live as long as we're a project; doesn't depend on hosting providers.

    * Vision is that writing KEPs, know from them what the roadmap is; can write the blog post based on the value articulated in the KEPs.

    * Right now they are [buried in community repo](/contributors/design-proposals/architecture/0000-kep-template.md) 3 levels deep: a lot of great feedback so far. Joe wouldn't object to making it more discoverable.

        * Kep.k8s.io for rendered versions?

        * Rust community has a sub-repo just for this (rust-lang/rfcs)

        * More than one person has said that KEPs weren't known about - move to somewhere discoverable sooner rather than later. Who can own that action? Matt Farina from Samsung is keen to help but doesn't have resources to lead.

        * *Can only have its own repo if we get rid of features repo!*

        * Don't want anyone to do anything not adding value to work; hope is that KEP is worthwhile and adds value. Caleb will help drive, and create repos as needed. 

## Features

(Jaice)

Features repo and process: I will say what I'm not happy about and what I've heard, but I do want to say there has been so much work from Ihor and SIG-PM to get where we are. If you haven't seen the machinations of feature/release product you won't know!


At velocity we have outgrown the notion of a central body leading.  Seeing increasing cross-SIG concern where some SIGs rely on others to get their work done.

We need to find ways to detangle those dependencies earlier.

KEPs are focusing on enhancements, changes in ecosystem.  A feature has a connotation of being user facing or interactive.  KEP can be for something contributor facing, doc facing transparent.  Get out of mindset we're delivering a "thing" - we're delivering value to user, contributor community, or both.

Currently the process is cumbersome.  We want to follow agile principles about local control, local understanding, sharing amongst teams.

### The steel thread

KEP provides opportunity for "steel thread" of custody. A KEP, as a living Markdown doc with metadata, lets you see high level idea at any group of time.  A SIG can break this down into meltable ideas for each release.  A KEP can last multiple milestones.  In terms of features/issues, defined by SIGs, issue should be actionable.  If a SIG writes an issue for that release milestone it should be approachable by any contributor who is a SIG member.

PRs roll up to the issue: links to this issue, which links to this KEP. For the first time we would be able to see at the PR level, what the grand scheme behind everything we do, is. Not heavyweight - just linking of issues.  Limit administrivia, paperwork to delivery value.

Q:

* For CloudProvider - discovery, had an idea, people started digging in; how do you keep updating the KEP?  Do you update as you go, you odn' tknow what you need?

    * Yes, it's a living doc first and foremost.  Has metadata (alpha, beta, GA). The issues worked on per milestone - say you pick 3, you document those issues; if during the course of the PRs to complete those issues you realise it's a misstep, you close the issue, you update the KEP, say you're modifying this; look at prior versions to see what it looked like before/what it looks like now. As you move to next release milestone the issues reflect that change: in terms of KEP, issues, planning.

* Give an analysis on why issues in feature repo didn't solve this problem and why KEPs will?

    * From features standpoint: SIGs interacted with features only when they had to. By trying to keep all that info separate from PRs, issues in each milestone: no way to tie that work back to the feature issue.  No way to easily understand where in this features repo issue, long if multi-milestone: where was work, and the link to the issue?  Eliminate the work

    * KEP is not solving that problem.  KEP is saying "how do we best define and discuss the value we're adding to the ecosystem".  Learns from patterns and antipatterns of other communities.  

    * Features process is central body of people not doing the work.

    * Some friction with the GitHub mechanics of issues, relabeling quarterly, etc; would be nice to keep pace with the number of issues. Moving to files to provide value will help make that clearer.

    * Managing issues in the features repo: hundreds of issues from the last year, created spreadsheets etc.  Bringing value but consuming extra resources. Not synced with features repo. Good we have a spreadsheet, but difficult to manage.

    * Started going through this process in SIG-Cluster LIfecycle: replacing GitHub thread (no one updates top) with a living document (use git history for mutations; see concise summary) - think it will be a huge improvement.

        * KEPs are in Google Docs now, haven't converted to PRs

* Does this mean the features repo is going away?

    * "Yes, eventually". 

* We need clear communication about where we are in this process. Are KEPs required, encouraged, etc?  Trying with one project, be useful to know what the granularity is.	

    * "It's in Alpha now" - trying to validate some assumptions.

        * Has momentum with SIG-Arch and SIG-PM

        * Assumption we will try and see if it works. 

        * Want to steward people into trying it to see if it's valuable.  If not, how do we make it so?

    * Community size and velocity now makes it difficult to radiate information to get what you need to know.

    * Kubernetes 1.10 would be a great place.

* Will it still be true in 1.10 that you need features? Can you have KEP instead of a Feature issue? Make it very obvious for 1.10 please.

    * One thing we need to sort out 

    * 1.10 will have existing features process

    * Cluster Lifecycle are experimenting with how to do differently

    * Would love to see SIGs try the KEP process but it's not a full surrogate for the features process yet.

* Existing features can be sorted, filtered etc - will there be tooling around KEPs? Until that's ready, I wouldn't want to switch

    * Proposals are different to features: issues will be associated with a KEP. In some ways it will be better as it's in the repo itself.

* How are the issues tied to the KEP?

    * Ties to this markdown file, in /kep/____

    * We need to work out the details.

    * When I create issue template, it will have a link to that KEP.

    * *Make sure it's searchable!*

* (Jago Macleod) Be explicit of scope of a KEP.  A trend to more SIG ownership, code, roadmap etc; context is a SIG plans their own work.  Certain faster moving SIGs, more capacity, to solve problems in the wrong component.  Some KEPs will span SIG boundaries and should be explicit about this in comms. 

* (Tim Hockin) GitHub UX is not going to be great for this.  It's not built for what we're trying to do.  All of this predicates on having an actual site that manages, a web app?

If that's true, can we start to collect the expertise?

Is there a SaaS tool for this?

    * See other projects using tracking tooling, used on files in source control

    * Won't be pretty or searchable in first iteration.

    * Want to see built out as they're consumed. Don't build tooling before people are using the process

    * *GH search is less than usable*

    * Do you have a place to host this?  We will find a place to host if someone is willing to write.  Matt will help write.

    * As someone involved with Maven repositories (trello) - it's a total pain in the neck.  More in favour of PRs, discoverability, etc.  Keep PRs for process. Push a site for discoverability.

* (Joe Beda) Current processes are discussion, argument, write-only: comment threads no-one reads. Make it more discoverable, more readable: check stuff into GitHub, use a static site generator, publishes it somehow, crosslinks between documents. Just like when you go to an RFC, reference one there's a link; that's the ecosystem of discoverability we want.

    * If someone comes to project and we have no way of telling what's happening.

* (Eric Tune) I created features process so I'm a little attached.  This discussion is a refactor vs rewrite debate.  The feature repo template is there; in GitHub; change it. See how many incremental changes we can make to solve proposals.  Put fields there and yell at people who don't fill it in.

* (Angus, Bitnami) As someone involved in Rust community:

    *  fairly well-read Rust weekly news summary, including a section with a summary on outstanding proposals, broken down into early and late discussion phases.  Read late phase to see what's a done deal vs. what's up in the air. There's also a RFC podcast, where they get a couple of people, have a chat show about what's involved in that proposal.  Lots of ways as a community member to stay up to date.

    * Fixed timeline: trying to get approval. If nothing happens by the end it's approved by default.  

    * May not want to copy but good to know.

* (Daniel Smith, Google) summarising: problems are mostly discoverability: we'll write tooling. Why can't we write tooling against that existing process? Both current and hypothetical new process suffer 

    * Interacting with objects in GitHub is hard at our scale, we have a limited number of tokens.

    * Procedural aspects: if I look at a PR, there's a link to an issue: that will tie up to a KEP.

    * Discoverability of relationships: exposed in the API. Links on GitHub are implemented by full-text search, not hard to do this.

* (Chen Goldberg) Schedule follow-up sometime this week - communicating is hard, we're all here!

* (Henning, Zalando) Structured format with some metadata: all in one repo, or distributed amongst incubator projects etc?  Are KEPs with a unique identity going to live in other project repos? Was this an idea?

    * Idea is to have as consolidated as possible: given committee questions of "what is Kubernetes, what follows our PM norms".  Problem is visibility for people who trust workloads in the software we create.  Want to provide this so someone can determine why a thing was done a certain way.  Projects outside the core repo: would follow a similar process in a perfect world.

If you're not excited about features process, try a KEP!

Reach out to SIG-PM - calebamiles@ 

### The long view

(Jaice) Have planning sessions about what SIGs are hoping to deliver for any milestone. Want to facility planning meetings in a more structured way (Planning as a  Service).  I did this for a proposal SIG-Cluster Lifecycle are working on: to have the discussion early and untangle dependencies, see when things could go off the rails, etc. Will talk to SIG leaders.

This planning activity will be one of the key success factors for the project moving forward.

Roadmap for 2018 (30min summary)
--------------------------

Notes by @jberkus

Speakers (check spelling): Apprena Singha, Igor, Jaice DuMars, Caleb Miles, someone I didn't get.  SIG-PM

We have the roadmap, and we have this thing called the features process, which some of you may (not) love.  And then we write a blog post, because the release notes are long and most of the world doesn't understand them.

Went over SIG-PM mission.  We had several changes in how the community behave over 2017.  We are moving to a model where SIGs decide what they're going to do instead of overall product decisions.

2017 Major Features listed (workloads API, scalability, networkpolicy, CRI, GPU support, etc.).  See slides.  The question is, how did we do following the 2017 roadmap?

Last year, we got together and each SIG put together a roadmap.  In your SIG, you can put together an evaluation of how close we came to what was planned.

Q: Last year we kept hearing about stability releases.  But I'm not sure that either 1.8 or 1.9 was a "stability release".  Will 2018 be the "year of the stability release?"

Q: Somehow the idea of stability needs to be captured as a feature or roadmap item.

Q: More clearly defining what is in/out of Kubernetes will help stability.

Q: What do we mean by stability?  Crashing, API churn, too many new features to track, community chaos?

Q: Maybe the idea for 2018 is to just measure stability.  Maybe we should gamify it a bit.

Q: The idea is to make existing interfaces and features easy to use for our users and stable.  In SIG-Apps we decided to limit new features to focus everything on the workloads API.

Proposals are now KEPs (Kubernetes Enhancement Proposals) are a way to catalog major initiatives.  KEPs are big picture items that get implemented in stages.  This idea is based partly on how the Rust project organizes changes.  Every SIG needs to set their own roadmap, so the KEP is just a template so that SIGs can plan ahead to the completion of the feature and SIG-PM and coordinate with other SIGs.

Q: How do you submit a KEP?
A: It should live in source control.  Each KEP will releate to dozens or hundreds of issues, we need to preserve that as history.

If you look at the community repo, there's a draft KEP template in process.  We need to make it a discoverable doc.
