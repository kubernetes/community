LTS Session

13 attendees

Opening slides shared:

 * https://docs.google.com/presentation/d/1Q0ZkKP_6jAZezWRF3aiDESflVm-1oGXz8lX9LD2FbFQ/edit?usp=sharing
 * https://static.sched.com/hosted_files/kcsna2019/d6/2019%20Contrib%20Summit%20-%20WG%20LTS.pdf

We've shied away from talking long term support because we don't want to predefine the mission but we had to call the WG something.

This is a summary of what's happened in 2019.

We took a survey in Q1 of 2019. 600 people started filling it out, 324 completed it.  Survey was very long with a lot of details.  People are upgrading continually, so we're cloud.  But things are moving more slowly in infrastructure.  A lot of people are falling behind.  We got 45% users so we're not just talking to each other.

Put your Q1 hat on.  At that time, 1.13, 1.12, 1.11 were under support.  Even 1.11 will be out of support in 2 months.  specifically 1.9 and 1.10 were a big chunk of people who are just out of support.

Why are they falling behind?  Well, some don't care.  Many want to be up to date, but there are lots of business reasons to not upgrade.

The other thing we discussed is what does "support" really mean?  Critical bug fixes, upgrade path so that users can get to newer versions.  ALso user stability & API compatibility.  We're relatively "normal" as a project relatively to the general industry.

Patch releases, we maintain 3 branches, each branch gets 9+ months of support; around the eol edge there's a fuzzy period where we don't necessarily stop patching, depending.  Lots of people said "why 9 months", which is a weird timespan.  Also we only support upgrading to the next version, but that's standard for industry.

API discussion: rest APIs, config data, provider abstractions.  We have robust promotion process, better than industry.

Proposals: suggested faster releases, like monthly.  Or maybe slower releases (6 months?). Or do a Kernel.org and pick a single release per year and have it be an "LTS".

We need to separate out patch releass, upgrades, stability.  Distinct although related.

API stability options:  this is all the SIG-arch stuff.  KEPs, conformance, pushing for more key APIs to hit stable.  Only 1 or 2 APIs out of 60 still not v1.  What about stale APIs?  Should we do a cleanup?

Upgrades:  this is hard.  maybe preflight checks?

Patch Release: we have a draft KEP for 4 branches of patch release support, which is 1 year of support.  We can do something impactful -- 30% of userbase is in that window of 3mo out of support.  Cost is known.  Google runs most things, but k8s-infra can quantify.  Because of reoganization of patch release team it's not as much effort.  We could stand to streamline how we generate patches though.

The WG should continue for one more year.  Maybe another survey, more concrete action items, and get contributors around those.

Brian Grant:  every version of Kubernetes we have has schema changes.  We don't have a way to unbundle those from API changes, which would be required to skip releases.  Releases a year old are just stabilizing now because they've been used.  We don't want to support 1.13 for 2 years, so we need to make releases more stable faster.  So more test coverage.   The reason we're patching the same thing into 4 different branches is that we find problems very late.  If we can get people using newer releases sooner we'll find problems sooner.
How do we fix this?  Better test coverage.  Not letting features going in until they're more mature, but that could mean finding issues slower for those.  Maybe we could not merge things without good test coverage.  We experimented with feature branches.  And with multiple repos, we should decide maybe we shouuldn't integrate a repo.  We have better test coverage for compatibility, but still happens with one thing every release.
Anyway, that's the whole philosophy of faster releases.
People are on the versions they're on, because that's where they've found stability.  Just doing things today the same wont lead to things being more stable.

Jago: The three letters "LTS" are easy to say, made people nervous last year, but the WG has been very admirable for keeping discussion open minded toward any possibility.  Need to also look at support programs for external repos.  This becomes a combinatorial explosion.  Lots of good work over the last year.

Tim: we can also make ourselve more consumable in the aggregate.  We need more distro experience.  Where's the debian of Kubernetes?

Jago: I support the 1-year extension.  Nick: the 9-month window means you'll have to upgrade at a bad time.  Jago: make sure we know it is 4 this year and not 5 next year.  12 months is the important part.  Don't enable a slippery slope.

Tim: everyone in the conversation has not wanted to go to extremely long support.  Just asking for a little bit more.

Josh: vendors don't have all the answers either.

Josh: even if stability is perfect, people have other business reasons not to upgrade.  Regulation, certification, management approval, time required to do the upgrade.  Nick: regulatory environment works with yearly upgrades.

Quinton: companies batch upgrades.  Tim: What's an example of software with great compat?  Even if there is one, businesses build risk management processes which make upgrade friction.

Josh: Kops stalled during the survey and their users weren't upgrading.  Could happen on any component.

Everyone good with 12 months of patching?  Everyone was good.

Noah: Likes the 12 month idea, but also sees no path technically viable toward "traditional LTS" of pick 1 release and support it for years.

BrianG: can technical changes make it cheaper/faster to move to new releases?  Improve tools, audit ability. Nick: even machine parseable changelogs is a notable improvement.

Mark Johnson: Q4 can actually be a period where ops has free cycles to do beta experimentation. Could be about to get more beta feedback now, since we're making beta releases again.

Quinton: even if people only upgrade once a year, people will upgrade at different times, so we'll get feedback around the year.

Tim: one user said that they actually do skip version updates, not sure how they do it.  For Nodes you can do it, but for control planes it's known to be unsafe.  Creating new clusters and migrating is one of the things that people do.  SIG-MC says that the idea of MC for migration is popular.  These users may be missing subtle compat issues in their clusters today.  Things likely better the more stateless you are.

Jago: upgrade test and tooling needs work, too few people contributing.  They're not covering everything and they break all the time.  Josh: our upgrade tests don't test an upgrade path that anyone uses (GCP & kube-up.sh versus other providers and upgrade tools).

Tim: what about more common tools instead of kube-up and 12 different installers.

Noah: time for another survey?  Tim has started discussion of combining forces on a new survey with the "Production Readiness" folks who are considering a survey soon.  Josh says we need to digest the existing data first, figure out what we want to ask.  Quinn: can we consolidate surveys across several groups?  Maybe consolidate them?  So have a multi-SIG survey?  Maybe do something like Linux Registry?  Give more thought to role focuses, maybe do multiple small surveys, still in a coordinated way across SIGs/WGs, but try to avoid a long survey.

BrianG: overarching theme here is shortening the release feedback loop.
