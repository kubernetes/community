End User Panel

Josh Mickelson Conde Nast
Federico Hernandez Mellwater, provides kubernetes-as-a-service to their developers.
Andy Snowden, Equityzen

Brian Grant
Peter Norquist & Kevin Fox, PNW National Lab
Josh Berkus
Ryan Bohnham, Granular.

How many developers & clusters?

JoshM: 380 developers across all clusters now, 18 clusters.
  8 clusters in international, 10 clusters in US.
  Separate production and dev clusters.

Federico: dev, test, production clusters.  We might have special purpose clusters in the future, like ML, but we dont' need them yet.

Andy: just getting started, have 8 devs, 1 cluster, very small.

Peter: none of our clusters launched yet.  Going to have 2 clusters, one for dev purely internal, the prod cluster will be managed by the security group because exposed.

Brian: about 200 engineers, have 15 clusters today.  Spread across 4 regions. Run dev/test/prod in each region.

Paris talked about navigating the project, when you don't have contributors on staff.

How are people deloying and upgrading clusters?

JoshM: CoreOS techtonic, are still maintaining a fork of that, investigating EKS, ClusterAPI.  On the US side, rolled everything manually, mainly cloudformation and scripts.  They don't upgrade, mostly.

Federico: Run on AWS, use KOPS w/ Amazon CNI plugin, love/hate relationship with KOPS.  Discovered a lot of bugs with CNI plugin, discovered a lot of bugs and submitted PRs for Kops.  Wrote our own custom upgrade script because Kops was not zero-downtime.  Specifically, zero downtime for applications.

Andy: we use Kops too.  Are just encountering the same problems.  Right now we are allowed downtime because 9 to 5, that will change, and Kops doesn't support that.  The lack of rerouting service requests while upgrading is an issue.  Stay around 2 releases behind.

Kevin: using kubeadm, have a container-based repo do to CentOS kickstarts with custom code. Custom stuff for upgrading the nodes using labels.  Metalkubed looks interesting once that's mature.

Ryan: Kops again.

JoshM: the Techtonic parts don't do the full upgrade, so we have custom scripts.  We run CoreOS, and Docker.

Kevin: some CRIO versions change the image storage, so you have to drain a node. Nothing in the KubeAPI lets you completely drain a node, daemonsets are a problem.

Federico: we also double all the nodes in the cluster, cordon the old ones, and migrade over.  That's worked well so far, better than node-at-a-time via Kops.

Andy: we've had a few times when Istio gets restarted where we lose a request.

JoshM: our cluster serves stuff that can be cached, so having nodes shut down is not as much of a problem.

Andy: Any manual control of the load balancers?  JoshM: No.  We have an ELB in front of everything.

JoshB: is upgrading everyone's biggest problem with Kubernetes?

Federico: yes.   Especially you run into dependencies, and you have to upgrade all of them, it's a puzzle to solve.  We have add-ons to give users a better experience, like DNS, we have the cluster scaler for cost management.  Those need to be maintained when you upgrade the cluster.  Installing them ourselves, not using Kops add-ons, we wrote our own installer.

Andy: we had the same experience with Helm.  For the cluster to be useful you have to install a lot of add-ons.  Like for resourcing, you need to figure out all your add ons.  Like we only have 70% of our cluster available because of all the add-ons.  Ryan: are you tracking the cost?  Andy: manually, we don't have tools.

... missed some discussions on cost accounting ...

Federico: Zalando has something for cloudcost.

JoshM: we don't have to track in detail. I'm lucky.

F: we need to forecast cost, it's not about chargebacks.  We need to know what it'll cost next year.  We have a couple teams who are really cost-aware.

Andy: one of our reasons to bring Kubernetes in was cost efficiency, so we want to know how much we saved.  We were able to downsize using K compared with previous AWS solution.  We compute cost-per-click.

JoshM: cluster upgrades are not the most difficult thing for us.  They're still difficult, but we've worked around a lot of the problems.  Right now our most difficult thing is getting our US infra in parity, so I guess that cluster upgrades are still a problem.

Federico: the other most diffcult thing for us is finding out cluster user changes.  Like which things in the release log will affect my users.  Finding that in the release notes is a challenge.  It has become a lot better, but it's still a big effort. I'd like to have "these are neutral changes, these are things ou need to know about, these are things users need to do something about".  Like when the Kubectl API removed duplicates.

Andy: yeah, that's also a big effort for me, reading through.

Kevin: two things: multitenancy.  After that, security things, like certificates etc.  We end up deploying ingresses for users in namespaces they can't manage, and we need workarounds for that.

Brian: do you expose Kubernetes directly to your app devs, or do you hide it behind a UI?

Federico: they have direct access to command line.  Most teams have some kind of CI/CD, but we don't hide anything.  They're still responsible for writing their own YAML.  A few teams use Teraform, a few use Helm with no Tiller.

JoshM: we optimized for app delivery.  We do expose it, but we put a lot of effort into base Helm charts etc. so that users use those templates (all apps deployed through same chart).  They use parameters to change things.  They can't deploy whatever the want, they have to go through CI/CD, there's several options, but they have to go through that.

Kevin: we try to do a "batteries included" configuration, so that our devs can have a common template where they can just deploy their applications.

Paris: do you feel supported from the Kubernetes community/ecosystem?

Andy: I haven't actually had to try yet.

JoshM: from the support side we haven't had any problems.  More recently I've started wanting to contribute back to the project, so most gripes have to do with initial contributor experience, PRs not getting reviewed, there's so many PRs in the queue.  My biggest question is how to jump in and get started?

Paris talked about doing a contributor orientation at companies.

Federico: we have some devs who are interested in contributions, but they're nervous about doing the contributor summit stuff, something at our office would be really nice.

Brian: is the structure of the project clear?  Like routing the PRs to the right place?

Paris went over the contributor structure.

Kevin: I contributed to Openstack, you got influence by showing up at meetings, but there was no way to get visibilty across the whole project.

Kevin: the problems I run into are typically WG problems, those really help me.

JoshM: one of the things I've pushed for is hackdays, that's easier that getting my company to pay for full time contributors. Are there features we can just knock out?  Folks suggested Prow or Docs.

Kevin: a lot of the docs make the assumption that the developer and the cluster admin are the same person, so we need to separate personas.

Federico: I copied code comments into documentation for Kops, stuff got noticed much faster.
