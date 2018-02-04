
Feature Workflow
----------------

Notes by @jberkus

TSC: Getting something done in Kubernetes is byzantine.  You need to know someone, who to ask, where to go.  If you aren't already involved in the Kubernetes community, it's really hard to get involved.  Vendors don't know where to go.

Jeremy: we had to watch the bug tracker to figure out what sig owned the thing we wanted to change.

TSC: so you create a proposal.  But then what?  Who needs to buy-in for the feature to get approved?

Dhawal: maybe if it's in the right form, SIGs should be required to look at it.

Robert B: are we talking about users or developers? Are we talking about people who will build features or people who want to request features?

???: Routing people to the correct SIG is the first hurdle.  You have to get the attention of a SIG to do anything.  Anybody can speak in a SIG meeting, but ideas do get shot down.

Caleb: we've had some success in the release process onboarding people to the right SIG. Maybe this is a model. The roles on the release team are documented.

Anthony: as a release team, we get scope from the SIGs.  The SIGs could come up with ideas for feature requests/improvement.

Tim: there's a priority sort, different projects have different priorities for developers.  You need a buddy in the sig.

Clayton: review bandwidth is a problem.  Review buddies hasn't really worked. If you have buy-in but no bandwidth, do you really have buy-in?

TSC: The KEP has owners, you could have a reviewer field and designate a reviewer.  But there's still a bandwidth problem.

Dhawal: many SIG meetings aren't really traceable because they're video meetings.  Stuff in Issues/PRs are much more referencable for new contributors.  If the feature is not searchable, then it's not available for anyone to check.  If it is going to a SIG, then you need to update the issue, and summarize the discussions in the SIG.

TSC: Just because a feature is assigned to a SIG doesn't mean they'll actually look at it.  SIGs have their own priorities.  There's so many issues in the backlog, nobody can deal with it.  My search for sig/scheduling is 10 different searches to find all of the sig/scheduling issues.  SIG labels aren't always applied.  And then you have to prioritize the list.

???: Test plans also seem to be late in the game.  This could be part of the KEP process.  And user-facing-documentation.

Robert B: but then there's a thousand comments.  the KEP proposal is better.

???: The KEP process could be way to heavy-weight for new contributors.

???: new contributors should not be starting on major features. The mentoring process should take them through minor contributions.  We have approximately 200 full-time contributors.  We need to make those people more effective.

TSC: even if you're a full timer, it's hard to get things in and get a reviewer.  Every release, just about everything that it's p0 or p1 gets cut, because the person working on it can't get the reviewer all of the stuff lined up.

Caleb: you need to spend some time in the project before you can make things work.

Dhawal: is there a way to measure contributor hours?  Are people not getting to things because people are overcommitting?

Jago: The problem is that the same people who are on the hook for the complicated features are the people who you need to review your complicated feature.  Googlers who work on this are trying to spread out their own projects to that they have more time at the end of the review cycle.

Jaice: If you're talking about a feature, and you can't get anyone to talk about it, either the right people aren't in the room, or there just aren't enough people to make it happen.  If we do "just enough" planning to decide what we can do and not do, then we'll waste a lot less effort.  We need to know what a SIG's "velocity" is.

Connor: the act of acquiring a shepard is itself subject to nepotism.  You have to know the right people.  We need a "hopper" for shepherding.

Tim: not every contributor is equal, some contributors require a lot more effort than others.

Robert: A "hopper" would circumvent the priority process.

Josh: there will always be more submitters than reviewers.  We've had this issue in Postgres forever.  The important thing is to have a written, transparent process so that when things get rejected it's clear why.  Even if it's "sorry, the SIG is super-busy and we just can't pay attention right now."

Dhawal: there needs to be a ladder.  The contributor ladder.

TSC: a lot of folks who work on Kube are a "volunteer army."  A lot of folks aren't full-time.

Caleb: there is a ladder.  People need to work hard on replacing themselves, so that they're not stuck doing the same thing all the time.  How do you scale trust?

???: Kubernetes is a complicated system, and not enough is written down, and a lot of what's there we'd like to change.  It's a lot easier for a googler to help another googler, because they're in the same office, and the priorities alighn.  That's much harder to do across organizations, because maybe my company doesn't care about the VMWare provider.

Jaice: for the ladder, is there any notion that in order to assend the ladder you have to have a certain number of people you shepherded in?  There should be.

TSC: frankly, mentoring people is more important than writing code.  We need to bring more people into Kubernetes in order to scale the community.

Josh: we need the process to be documented, for minor features and major ones.  Maybe the minor feature process belongs to each SIG.

Jaice: the KEP is not feature documentation, it's process documentation for any major change.   It breaks down into multiple features and issues.

???: The KEP needs to include who the shepherds should be.

Clayton: reviewer time is the critical resource.  The prioritization process needs to allocate that earlier to waste less.

Jeremy: the people we sell to are having problems we can't satisfy in Kubernetes.  We have a document for a new feature, but we need every SIG to look at it (multi-network).  This definitely needs a KEP, but is a KEP enough?  We've probably done too much talking.

Clayton: the conceptual load on this is so high that people are afraid of it.  This may be beyond what we can do in the feature horizon.  It's almost like breaking up the monolith.

Robert: even small changes you need buy-in across SIGs.  Big changes are worse.

Connor: working groups are one way to tackle some of these big features.
