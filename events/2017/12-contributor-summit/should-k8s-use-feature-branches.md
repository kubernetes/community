# Should Kubernetes use feature branches?

Lead: Brian Grant

Note Taker(s): Aaron Crickenberger

Focusing on our current release cadence:

- We spend several weeks putting features in
- We spend several weeks stabilizing the week kinda/sortof

The people who are managing the release process don’t have the ability to push back and say this can’t go in and we don’t want to slip the release so they’re left powerless

In earlier session today over breaking up monolith:

- We have over 90 repos
- The releases are getting bigger and more complex
- Not everything we’re moving out of the monorepo, not everything makes sense

What do we do about those things
I think the only way we can get out of the release fire drills is to have an earlier gate

Q: Aaron: Why not use the tests we have today and push back saying the tests have to be green?

- A lot of the new features have inadequate testing or inadequate features
- We want to continue to release on a regular on cadence
- Use some model like gitflow that puts feature work into feature branches before they land into master
- I think we need to allow features to be developed in feature branches instead of landing in kubernetes/kubernetes master
- But we also can’t have large monolithic feature branches
- Because one thing going in might cause rebase hell for other things going in
- Test machinery is getting to a point where it can use branches other than master to spin up tests
- You can spin up a feature branch and get tests running on it without too many issues
- I think we have to do within either 1.10 or 1.11 release

Q. I heard two 2 concerns:

- How do you protect mainline development from these features going in?
- How do you know when they’re ready to go in?
- So nodejs claims to have something like 90-something% coverage, we don’t even have coverage measurement
- With the interdependence of everything in the system it’s basically impossible to back changes out

Q: (thockin) 
How do we think we can do this across repos?

- Can we hear from people who have actually tried this for what the pain points are?
- There are lots of details to work out and people are just going to have to try it at small scale and then at large scale
- A lot of the times people try to get their foot in the door with tiny feature PR’s, and then docs, and then feature completion
- Jan commenting on experience with workload api’s
- Tried bumping workload api’s from v1beta2 to v1, tried it on feature branch
- Main blocker was tests, tests weren’t running against the feature branch, only the master branch automatically
- We decided we were just going to do it directly against the master branch

Q (spiffxp) How do we prevent merge chicken, aka racing to get in so you don’t have to avoid rebase hell

- With api machinery, one thing we’ve done is wait until after the release cycle to minimize the amount of time that people have to do rebasing
- We scheduled that massive change to land right after code freeze
- Re: how do we prevent once you’re in you’re in… that’s basically the way it goes right now

Thockin: 

- the pattern that I see frequently is somebody sends me a 10k line PR and I see please break this up and they do, but then they forget the docs and somewhere along the way we miss it
- Bgrant: it’d be great if we could ask people to write the docs first (maybe this is a criteria for feature branches)
- Venezia: what about having a branch per KEP? Also, facebook does something like feature flags, is that something we could do or is that just impossible?

Bgrant: 

- That can be used, but it requires a huge amount of discipline to ensure those flags are used correctly and consistently throughout the code
- Some features we can easily do that eg: if it’s an API, but it’s much harder to do with a feature that’s threaded through components
- A branch per KEP is roughly the right level of granularity
- What’s the difference between a feature branch and a massive PR?
- Review flow isn’t well suited to massive PR’s
(thockin) It’s easier on maintainers if we get things in and lock them in, but not on the main tree
(thockin) one thing I like about the linux kernel model is the lieutenant model

Q: (maru) you’re still going to have to rebase feature branches the way that you rebase PR’s right?

- A (jago) the difference is that a whole team could work on a feature branch
- I like the idea of trying a few projects and then reporting back
- I think roughly the level of granularity of having a branch per KEP sounds about right
- Multiple feature branches affecting the same area of code would be a very useful place to get some information

Q (???) if someone needs to rebase on a feature branch, could it instead be resolved by a merge commit instead?

- (thockin) I think we could solve that with technology and it could be better than building a better review tool
- Rebasing sucks, part of the reason we have that is generated code, part of it is poor modularity
- We could stand to expend some energy to improve modularity to reduce rebases
(bgrant) either feature branches or moving code to other repos are forcing functions to help us clean some of that up

Q: how do we make sure that when issues are filed people know which commit/feature branch they should be filing against?
- (bgrant) I suspect most issues on kubernetes/kubernetes are filed by contributors
- (jdumars) imagine you have a KEP that represents three issues, three icecubes you’ve chipped off the iceberg, the feature branch would hold code relevant to all of those issues
- (luxas) feature branches should be protected

Yes we’ll definitely want a bot that automatically
Q: (alex) do we want to have feature branches
Q: (thockin) why not just use a fork instead of branches?

- I think it’s harder to get tests spun up for other users
- (jgrafton) worry about adding yet another dimension of feature branches to our already large matrix of kubernetes and release branches

- Not so much for PR’s, but for CI/periodic jobs
- You need to run more than just the PR tests
- The experience with gated features is that tests do run less often, so racey bugs that pop up less frequently get exposed less often; so there is a concern that we wouldn’t get enough testing coverage on the code
- We’ll figure out what the right granularity is for feature branches is, probably shorter than a year, but probably shorter than a year
- One idea is instead of support forks to users, what if we forked to orgs owned by sigs, and just made sure the CI was well setup for those orgs?
(ed. I’m nodding my head vigorously that this is possible)
- Just another comment on branches on main repo vs. in other repos
- The branches in the main repo could be another potential for mistakes and fat fingers that we could be better solve by making the tooling work against arbitrary repos
- (jgrafton) related to the ownership of e2e tests and the explosion of stuff.. Do we expect sigs to own their tests across all feature branches? - Or should they only care about master

Q: worry about tying feature ownership too closely to a single sig, because features could involve multiple sigs

If you want to get started with feature branches, reach out to kubernetes-dev first and we’ll get you in touch with the right people to get this process started
