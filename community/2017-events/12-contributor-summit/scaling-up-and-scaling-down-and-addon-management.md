# Scaling Up & Scaling Down & Addon Management Session…

Notes by @justinsb 

## Scaling Up & Scaling Down

Lots of users that want to run on  a single node cluster - one node, one core
thockin’s position: 1 node 1 core should continue to work (with 3.75GB, but ideally less than 2GB)
Works today but only just - e.g. you only have 0.2 cores left today
Addons / core components written with a different target in mind
Some of the choices not optimal for single node (e.g. fluentd / ES)

Cluster Proportional Autoscaler tries to solve this
    Targets a daemonset / deployment and changes number of pods

CPVPA Cluster Proportional Vertical Pod Autoscaler
    Applies a step function based on the number of cluster cores to change the resources needed by a target deployment / daemonset
Prototype being used today in GKE for calico
    Needs a pod per target which leads to higher resource usage
    Idea: to have a single pod that uses a CRD to amortize the costs across multiple targets
    Problems we want to address:
        We want more granularity at the low end, both because it matters more proportionately
        We want to avoid constant rescaling - changing resources requires restarting pods

Q (solly): What is missing from kubernetes today - what is different vs generic HPA/VPA?  Why can’t I reuse the existing HPA/VPA?  How do I tell the difference from a UI perspective?
A (thockin): We could wrap the autoscaler; or we could have a ref pointing back up; or we could integrate back into HPA/VPA.  The third option is mostly what we’ve encouraged - where things are proved out of tree and then merged back in.

Some changes needed in HPA/VPA - pulling in metrics from a different source e.g. number of cores in cluster.  (Do we need a “count” API)

Q (clayton): Does this indicate that requests / limits are “broken”?  Should we instead be thinking about e.g. overcommit to deal with smaller sizes.
A (thockin): No - the step function is defined by the author of e.g. kube-dns based on their knowledge of the requirements.   And this is typically for “fundamental” components, not user applications, and so the broader question for users is separate.

Q (gus): We need to consider architectures (e.g. 64 bit might use more memory).  Are we wasting more memory running a pod to right-size the other pods?
A (thockin): That’s why we want one pod with a CRD, but yes … we want to fold it back in. but we have to prove it first.

Q: Why can’t this be done in e.g. kubeadm
A: Because clusters change size - e.g. cluster autoscaler

Q (jago): Focusing on small clusters - we should come up with clear guidance for resource requirements.  Maybe we need T-shirt sizes which helps us figure out what we can run in various clusters.
A (thockin): We need to start collecting data - that’s part of this effort.

Q (vish): Is kubernetes the right approach for small clusters
A: Yes!  Yes!  Yes! Lots of resounding yeses - devs want to run on their laptops, people happily run personal clusters on single nodes etc.

Q (dawn): It is hard to come up with a single scale function - e.g. with docker on the node team the kernel options have a big impact.  But we should still provide a scaling function as a guide, and to allow for operators to tweak it

Q (chrislovecnm): Key point - we must collect data!

Q: There is a lack of documentation on resource requirements e.g. memory requirements of apiserver at various sizes.


## Add-on Management

It’s been the can that’s been kicked down the road
Questions:
    What is an addon?
        Kubeadm says kube-dns, kube-proxy
        GKE adds dashboard & heapster

Thockin: Addons are things we manage vs things you manage.  We = administrator, you = User

Luxas: kubeadm installs minimal addons to get conformance to pass

solly: Important not to conflate addons with discussion of what is core vs not

Clayton: Is Istio an addon?  What if it becomes super-complicated... 

Mrubin: 3 categories of addons:
    Istio, Service Catalog, Service Broker
    Network policy
    Kube-proxy and kube-dns
Sniff test: Does it have to be tied to a version of kubernetes?

Robertbailey: Maybe addons should be the minimal set of things required to pass conformance

Chrislovecnm: What’s our goal - we want something to install a manifest on the cluster?

Clayton: the challenge isn’t installation - it’s upgrades / updates.  Are we really ready to manage this?  Addons should be things that we can actually manage, including addons.

Thockin: we want a simple solution, not a full solution for arbitrarily complicated applications.

Jago: We want a set of components that as a platform operator we install and upgrade/version together and validated together, and cluster-level.  This is distinct from user applications or even application etcd.

Bboreham: cross-grade between alternatives is a related aspect - e.g. swap out implementation of network policy.  In practice today you have to reinstall the cluster, because otherwise it requires the network policies to cooperate.

Clayton: comparison to runlevels - where it’s very possible to brick your cluster by doing things out of sequence.

Gus: The installation tools do all differ in what addons they are installing (e.g. AWS and GCE have very different requirements).  So the installation tool will have opinions here - what value are we adding?

Thockin: There is an opportunity here to define something cross-cutting - “apt or yum for kubernetes”.  Having each installer do something different is not great.

Clayton: apt or yum doesn’t get you all the way - you still need to tweak configurations.

Mrubin: Kubernetes is a bunch of small pods cooperating.  The mechanism is the same for installing DNS as it is for network policy - which is good in that we have a harmonized approach.  We want to figure out the mechanism by which components can be packaged by authors, discovered, and integrated by installers.

Mrubin: There is a belief that some pods in kube-system are special

Dawn: The kube-system namespace is known to be system-managed / updated and critical pods - even if it isn’t actually special in a technical sense.

Jess: We want to avoid decision paralysis.  And there’s odd incentives where people want to monetize their code or their approach.  We have to limit the scope as much as possible.

Robertbailey: Agreed - we have to scope this just to the minimal components needed to install stuff.  And then users can install more opinionated things like helm.

Thockin wrapping up: Still an open question of what we should do here, how the various distributions should avoid duplicating effort.
