# Breaking Up The Monolith

Note Takers: Josh Berkus(@jberkus), Aaron Crickenberger (@spiffxp)

Here's what we're covering: build process, issue tracking, logistics.

## Question: are we committed emotionally 100% to splitting the monolith

Opinion: committed, but need to define what that means (past provides the impetus)

We've been trying [the monolith] for so long and it's too expensive.

- erictune: as the project matures the core can't keep growing at the same rate
- mattfarina: reason to break it up - developer experience,difficulty to contribute pushes people away from contributing (hard to step into issue queue, etc)
- clayton: I'm scared by this because kubectl is never going to be simpler, it may be easier to just write a new one
- erictune: ecosystem yes, more contributors... narrowly defined core no
- bburns: we don't do a good job of measuring the cost of this, we love breaking things up because it's "cleaner" but we're not estimating the cost in terms of the human effort and complexity
- lavalamp: respectfully disagree, we _have_ estimated the cost, knew it was going to be painfu but we haven't had a choice
- timstclair: nobody has ever answered the question to be of how we're going to deliver a set of binaries as part of a release, there isn't a cohesive plan for doing this
- timstclair: how are you actually going to get test infra to test all the things
- bburns: how do you actually find the commit that causes a failure in e2e
- lavalamp: have you see the doc I have posted everywhere this comes up? (TODO: link) main repo becomes integration point, binary artifacts produced by other repos
- clayton: kubectl is a very "big" thing, if we said the problem is that kubectl itself is too big, it might be better to focus on something new and novel, it might be better to build a community around that
- dims/matt: can we have a list of four or five items we want to move out in priority
  - Kubectl
  - Client-Go?  (but that's already published separately) (this is way too involved for this session)
  - Try to tackle "how do we deal with client-go" elsewhere
  - kubeadm is in main repo to catch release train, wants to move out, but it's fairly hard.  We tried to move it out, but couldn't and follow the release train.
- clayton: government approach: why build one thing when we can build two for twice the cost? an (extracted) client-go that's all high-perf, a completely rewritten client-go that's human focused and more reusable

## Question: Do we understand the problem we're trying to solve with the repo split

Assumption: contributor velocity is correlated with velocity, new contributor experince
Assumption: "big tangled ball of pasta" is hard to contribute to

- thockin: our codebase is not organized well (need to make it easier to actually *find* what you want to contribute to) one of the things we could do is "move code around" until it looks effectively like there are multiple repos in our repo? eg if we centralized all of the kubelet logic in one directory except for one util directory, maybe that would be an improvement
- jberkus: people who work on other large projects say that modularity is the way to go.  Not necessarily separate repos, but a modular architecture.
- thockin: what we've done with github bots alleviates a lot of the pain we've had with github
- dims: two problems
  - vendoring in all kinds of sdk's that we don't care about, they get dragged in (e.g. AWS SDK) and if you're only working on eg: openstack related stuff, it's hard to get signoff on getting other stuff in if it's in your own repo it's easier to do that than in the main repo
  - (guess we missed #2)
- erictune: we've already promised people cloud providers you can write in different ways, extensibility of apis, operators, if we deliver on all this we will have enough modularity to improve user/developer experience
- mattfarina: I think you're right about that
- First thing: we're breaking up cloud providers, it helps let the long tail of cloud providers go out
  - what about storage providers, if we break them up there's a clear interface for storage providers
- second thing: here are api/components that are required, required but swappable, optional... how can you do that swapping around of stuff in the monorepo
- robertbailey: when I proposed this session I thought we had all agreed that breaking the repo was the way to go, but maybe we don't have that conensus could we try and figure out who is the decision maker / set of decision makers there and try to come to consensus
- dchen1107: I heard a lot of benefits of the split.  So there's a huge cost for release management, as the release manager, I don't know what's in the other repository.  Clear APIs and interfaces could be built without needing separate repositories. Gave example of Docker and CRI, which got API without moving repos.
- Example of Cloud Foundry, build process which took two weeks.  How do we ensure that we're delivering security updates quickly?
- thockin: We can put many things in APIs and cloud providers are a good example of that.  Splitting stuff out into multiple repos comes down to the question of: are we a piece of software or a distribution?  You'll have to come to my talk to see the rest.
- spiffxp: Increasing our dependency on integration tools will add overhead and process we don't have now.  But I understand that most OSS people expect multiple repos and it's easier for them.  Github notifications are awful, having multiple repos would make this better.  The bots have improved the automation situation, but not really triaging notifications.  How many issues do we have that are based on Github.  Maybe we should improve the routing of GH notifications instead?
- solly: if we split, we really need to make it easier for tiny repos to plug into our automation and other tools. I shouldn't have to figure out who to ask about getting plugged in, we should have a doc.  We need to make sure it's not possible for repos to fall through the cracks like Heapster, which I work on.
- bgrant: We're already in the land of 90 repos.  We don't need to debate splitting, we're alread split.  We have incubator, and kubernates-client.  I think client has helped a lot, we have 7 languages and that'll grow.
- bgrant: the velocity of things in the monorepo are static
- thockin: if issues in the main repo are more than 6 weeks old, nobody looks at it.  There's like 400-500 abandoned "network" issues.
- bgrant: we tested gitlab and gerrit, and importing kubernetes failed.  We didn't find them that much better.
- spiffxp: we have hundreds of directories with no owners files, which is one of the reasons for excessive notifications.
- bgrant: kubernetes, as a API-driven system, you have to touch the api to do almost anything.  we've added mechanisms like CRD to extend the API. We need SDKs to build kube-style APIs.
- dims: I want to focus on the cost we're paying right now.  First we had the google KMS provider.  Then we had to kick out the KMS provider, and there was a PR to add the gRPC interface, but it didn't go into 1.9.
- thockin: the cloud provider stuff is an obvious breeding ground for new functionality, how and if we should add a whole separate grpc plugin interface is a separate question
- jdumars: the vault provider thing was one of the ebtter things that happened, it pushed us at MS to thing about genercizing the solution, it pushed us to think about what's better for the community vs. what's better for the provider
- jdumars: flipside is we need to have a process where people can up with a well accepted / adopted solution, the vault provider thing was one way of doing that
- lavalamp: I tend to think that most extension points are special snowflakes and you can't have a generic process for adding a new extension point
- thockin: wandering back to kubernetes/kubernetes "main point", looking at staging as "already broken out", are there other ones that we want to break out?
- dims: kubeadm could move out if needed, could move it to staging for sure
- thockin: so what about the rest? eg: kubelet, kube-proxy... do we think that people will concretely get benefits from that? or will that cause more pain
- thockin: we recognize this will slow down things
- lavalamp: there are utility functions that people commonly use and there's no good common place
- lavalamp: for kubectl at least it's sphaghetti code that pulls in lots of packages and makes it difficult to do technically
- thockin: do we think that life would be better at the end of that tunnel, would things be better if kubectl was a different repository, etc.
- timallclair: I'm worried about dependency management, godeps is already a nightmare and with multiple repos it would be worse.
- luxas: in the kubeadm attack plan, we need to get a release for multiple repos.  We need the kubeadm repo to be authoritative, and be able to include it in a build.
- pwittrock: how has "staging" improved development? can we see any of the perceived or hoped-for benefits by looking at staging repos as example use cases?
- lavalamp: getting to the "staging" state and then stopping is because api-machinery was unblocked once we got there
- thockin: the reason I consider "staging" solved, is you have to untangle a lot of the stuff already
- erictune: I would make a plea that we finish some of our started-but-not-finished breakaparts
