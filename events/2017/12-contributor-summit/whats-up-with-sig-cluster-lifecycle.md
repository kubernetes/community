
What's Up with SIG-Cluster-Lifecycle
------------------------------------

Notes by @jberkus

Luxas talking: I can go over what we did last year, but I'd like to see your ideas about what we should be doing for the future, especially around hosting etc.

How can we make Kubeadm beta for next year?
Opinions:
* HA
  * etcd-multihost
  * some solution for apiserver, controller
* Self-hosted Kubeadm

Q: Can someone write a statement on purpose & scope of Kubeadm?

To install a minimum viable, best-practice cluster for kubernetes. You have to install your own CNI provider.  Kubeadm isn't to endorse any providers at any level of the stack.

Joe: sub-goal (not required for GA), would be to break out components so that you can just install specific things.
Would also like documentation for what kubeadm does under the covers.

Josh: requested documentation on "how do I modify my kubeadm install".  Feels this is needed for GA.  Another attendee felt the same thing.

One of the other goals is to be a building block for higher-level installers.  Talking to Kubespray people, etc.  Enabling webhooks used as example.

There was some additional discussion of various things people might want.  One user wanted UI integration with Dashboard.  The team says they want to keep the scope really narrow in order to be successful.  UI would be a different project.  GKE team may be working on some combination of Kubeadm + Cluster API.  Amazon is not using Kubeadm, but Docker is.  Docker for Mac and Windows will ship with embedded kubernetes in beta later this week.

Kubeadm dilemma: we want humans to be able to run kubeadm and for it to be a good experience, and for automation to be able to run it.  I don't think that can be the same tool.  They've been targeting Kubeadm at people, we might want to make a slightly different UI for machines.  Josh says that automating it works pretty well, it's just that error output is annoying to capture.

HA definition:
* etcd should be in a quorum HA standard (3+ nodes)
* more than one master
* all core components: apiserver, scheduler, kcm, need to be on each master
* have to be able to add a master
* upgrades

Or: we want to be able to survive a loss of one host/node, including a master node.  This is different, if we want to survive the loss of any one master, we only need two then. Argument ensued.  Also, what about recovery or replacement case? A new master needs to be able to join (manual command).

What about HA upgrades?  Are we going to support going from one master to three?  Yes, we have to support that.

Revised 4 requirements:
* 3+ etcd replicas
* all master components running in each master
* all TLS secured
* Upgrades for HA clusters

Everyone says that we want a production environment, but it's hard to define what "production grade" means.  We need to stop saying that.  Over time, what matters is, "is it maintained".  If it's still being worked on, over time it'll get better and better.

CoreOS guy: trying to do self-hosted etcd.  There's a lot of unexpected fragile moments.  Just HA etcd isn't well tested upstream, there's not enough E2E tests.  Self-hosting makes this worse.  The etcd operator needs work.  There needs to be a lot of work by various teams.  Self-hosted control planes work really well, they host all of their customers that way.  It's etcd that's special.

There's some problems with how Kubernetes uses HA Etcd in general.  Even if the etcd operator was perfect, and it worked, we couldn't necessarily convince people to use it.

Should Kubeadm focus on stable installs, or should it focus on the most cutting-edge features?  To date, it's been focused on the edge, but going to GA means slowing down.  Does this mean that someone else will need to be forward-looking?  Or do we do feature flags?

SIG-Cluster-Lifecycle should also document recommendations on things like "how much memory do I need."  But these requirements change all the time.  We need more data, and testing by sig-scalability.

For self-hosting, the single-master situation is different from multi-master.  We can't require HA.  Do we need to support non-self-hosted? We can't test all the paths, there's a cost of maintaining it.  One case for non-self-hosted is security, in order to prevent subversion of nodes.

Also, we need to support CA-signed Kubelet certs, but that's mostly done.

So, is HA a necessity for GA? There are a bunch of automation things that already work really well. Maybe we should leave that to external controllers (like Kops etc) to use Kubeadm as a primitive.  Now we're providing documentation for how you can provide HA by setting things up.  But how would kubeadm upgrade work, then?
