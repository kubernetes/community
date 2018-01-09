OBSOLETE

Cluster lifecycle includes deployment (infrastructure provisioning and bootstrapping Kubernetes), scaling, upgrades, and turndown.

Owner: @kubernetes/sig-cluster-lifecycle (kubernetes-sig-cluster-lifecycle at googlegroups.com, sig-cluster-lifecycle on slack)

There is no one-size-fits-all solution for cluster deployment and management (e.g., upgrades). There's a spectrum of possible solutions, each with different tradeoffs:
* opinionated solution (easier to use for a narrower solution space) vs. toolkit (easier to adapt and extend)
* understandability (easier to modify) vs. configurability (addresses a broader solution space without coding)

Some useful points in the spectrum are described below. 

There are a number of tasks/features/changes that would be useful to multiple points in the spectrum. We should prioritize them, since they would enable multiple solutions.

See also:
* [Feature-tracking issue](https://github.com/kubernetes/features/issues/11)
* [Cluster deployment v2 issue](https://github.com/kubernetes/kubernetes/issues/23174)
* [Shared infrastructure issue](https://github.com/kubernetes/kube-deploy/issues/123)
* [Kube-up replacement that works for everyone](https://github.com/kubernetes/kubernetes/issues/23478)
* https://github.com/kubernetes/kubernetes/labels/area%2Fcluster-lifecycle
* [Deployment DX Notes](https://docs.google.com/document/d/1jAQ5H222pJzpv-m0OsezZgdhdkf299E7H8Q9yzNrAjg/edit#)

We need to improve support/behavior for cluster updates/upgrades, as well. TODO: make a list of open issues. Examples include [feature gating](https://github.com/kubernetes/kubernetes/issues/4855), [node upgrades](https://github.com/kubernetes/kubernetes/issues/6079), and [node draining](https://github.com/kubernetes/kubernetes/issues/6080).

## Single-node laptop/development cluster

Should be sufficient to kick the tires for most examples and for local development. Should be dead simple to use and highly opinionated rather than configurable.

Owner: @dlorenc

See also:
* https://github.com/kubernetes/minikube

To do:
* Replace single-node getting-started guides
  * [Docker single-node](http://kubernetes.io/docs/getting-started-guides/docker/)

## Portable multi-node cluster understandable reference implementation

For people who want to get Kubernetes running painlessly on an arbitrary set of machines -- any cloud provider (or bare metal), any OS distro, any networking infrastructure. Porting work should be minimized via separation of concerns (composition) and ease of modification rather than automated configuration transformation. Not intended to be highly optimized by default, but the cluster should be reliable.

Also a reference implementation for people who want to understand how to build Kubernetes clusters from scratch.

Ideally cluster scaling and upgrades would be supported by this implementation.

Replace [Docker multi-node guide](http://kubernetes.io/docs/getting-started-guides/docker-multinode/).

To facilitate this, we aim to provide an understandable, declarative, decoupled infrastructure provisioning implementation and a portable cluster bootstrapping implementation. Networking setup needs to be decoupled, so it can be swapped out with alternative implementations.

For portability, all components need to be containerized (though Kubelet may use an alternative to Docker, so long as it is portable and meets other requirements) and we need a default network overlay solution.

Eventually, we'd like to entirely eliminate the need for Chef/Puppet/Ansible/Salt. We shouldn't need to copy files around to host filesystems.

For simplicity, users shouldn't need to install/launch more than one component or execute more than one command per node. This could be achieved a variety of ways: monolithic binaries, monolithic containers, a launcher/controller container that spawns other containers, etc. 

Once we have this, we should delete out-of-date, untested "getting-started guides" ([example broken cluster debugging thread](https://github.com/kubernetes/dashboard/issues/971)).

See also:
* [Summary proposal](https://git.k8s.io/kubernetes-anywhere/PROPOSAL.md)
* [kubernetes-anywhere umbrella issue](https://github.com/kubernetes/kubernetes-anywhere/issues/127)
* https://git.k8s.io/kubernetes/docs/proposals/cluster-deployment.md
* [Bootstrap API](https://github.com/kubernetes/kubernetes/issues/5754)
* [jbeda's simple setup sketch](https://gist.github.com/jbeda/7e66965a23c40a91521cf6bbc3ebf007)

To do:
* [Containerize](https://github.com/kubernetes/kubernetes/issues/246) all the components in order to achieve OS-distro independence
* [Self-host](https://github.com/kubernetes/kubernetes/issues/18735) as much as possible, such as using Deployment, DaemonSet, [ConfigMap](https://github.com/kubernetes/kubernetes/issues/1627)
  * Eliminate the need to read configurations from local disk
  * [Dynamic Kubelet configuration](https://github.com/kubernetes/kubernetes/issues/27980)
  * [Kubelet checkpointing](https://github.com/kubernetes/kubernetes/issues/489) in order to reliably run control-plane components without static pods
  * [Run Kubelet in a container](https://github.com/kubernetes/kubernetes/issues/4869)
  * [DaemonSet updates](https://github.com/kubernetes/community/wiki/Roadmap:-DaemonSet)
  * [DaemonSet for bootstrapping](https://github.com/kubernetes/kubernetes/issues/15324)
  * [Bootkube](https://docs.google.com/document/d/1VNp4CMjPPHevh2_JQGMl-hpz9JSLq3s7HlI87CTjl-8/edit)
* Adopt a default network overlay, but enable others to be swapped in via composition
  * Need to make it clear that this is for simplicity and portability, but isn't the only option
* [Link etcd and the master components](https://github.com/kubernetes/kubernetes/issues/5755) into a monolithic [monokube-like](https://github.com/polvi/monokube) binary and/or [finish hyperkube](https://github.com/kubernetes/kubernetes/issues/16508)
* [Replace multi-node Docker getting-started guide](https://github.com/kubernetes/kubernetes/issues/24114)
* Make it easy to get the right version of Docker on major Linux distros (e.g., apt-get install...)
  * It's easy to get the wrong version: docker, docker.io, docker-engine, ...

Starting points:
* https://github.com/kubernetes/kubernetes-anywhere
* https://github.com/Capgemini/kubeform

## Building a cluster from scratch

For people starting from scratch:
* http://kubernetes.io/docs/getting-started-guides/scratch/
* https://github.com/kelseyhightower/kubernetes-the-hard-way
* https://news.ycombinator.com/item?id=12022215

We should simplify this as much as possible, and clearly document it.

This is probably the only viable way to support people who want to do significant customization:
* cloud provider (including bare metal)
* OS distro
* cluster size
* master and worker node configurations
* networking solution and parameters (e.g., CIDR)
* container runtime (Docker or rkt) and its configuration
* monitoring solutions
* logging solutions
* ingress controller
* image registry
* IAM
* HA
* multi-zone
* K8s component configuration

To do:
* Simplify release packaging and installation
  * Finding and installing the right version of Docker itself can be hard (`apt-get install docker/docker.io/docker-engine` isn't the right thing)
  * Build rpms, debs?
* Verify that system requirements have been satisfied (docker version, kernel configuration, etc.)
  * And ideally degrade gracefully and warn if they are not
* Documentation
  * What is the latest release, how can I find it, how do I install it, what version of Docker/rkt/etc. is required?
  * An architectural diagram (like the one we use in our presentations) would help, too.
  * Explain the architecture
    * Link to instructions about how to manage etcd
    * Link to [Chubby paper](http://static.googleusercontent.com/media/research.google.com/en//archive/chubby-osdi06.pdf)
  * Document system requirements ("Node Spec")
    * OS distro versions
    * kernel configuration
    * resources
    * IP forwarding
  * [Document how to set up a cluster](https://docs.google.com/document/d/1c4DMomZgS1i6AlKbb_8CiTcuimotkcJZAqJtQTN1iqc/edit#heading=h.58fpuhrw9g2o)
  * Adequately document how to configure our components. 
    * Improve/simplify/organize command help
    * Hide/remove/deemphasize test-only options
  * Document how to integrate IAM
  * [Create guides to help with key decisions for production clusters](https://github.com/kubernetes/kubernetes/issues/10100)
    * Selecting a networking model
    * Managing a CA
    * Managing user authentication and authorization
    * Initial deployment requirements (memory, cpu, networking, storage)
    * Upgrading best practices
* Code changes
  * Finish [converting components to use configuration files rather than command-line flags](https://github.com/kubernetes/kubernetes/issues/12245)
  * Facilitate [managing that configuration using ConfigMap](https://github.com/kubernetes/kubernetes/issues/1627)
  * [Cluster config](https://github.com/kubernetes/kubernetes/issues/19831)
  * Reduce external dependencies
    * APIs for reusable building blocks, such as [TLS bootstrap](https://github.com/kubernetes/kubernetes/issues/18112), [certificate signing for addons](https://github.com/kubernetes/kubernetes/issues/11725), [teardown](https://github.com/kubernetes/kubernetes/issues/4630)
  * Need key/cert rotation ([master](https://github.com/kubernetes/kubernetes/issues/4672), [service accounts](https://github.com/kubernetes/kubernetes/issues/20165))
  * Finish generalization of [component registration](https://github.com/kubernetes/kubernetes/pull/13216)
  * [Improve addon management](https://github.com/kubernetes/features/issues/18)

## Production-grade, easy-to-use cluster management tools/services

Easy to use and opinionated. Potentially highly optimized. Acceptable for production use. Not necessarily easily portable nor easy to extend/adapt/change.

Examples:
* [Kube-AWS](https://github.com/coreos/coreos-kubernetes/tree/master/multi-node/aws)
* [kops](https://github.com/kubernetes/kops)
* [Kargo](https://github.com/kubespray/kargo)
  * (is https://git.k8s.io/contrib/ansible still needed?)
* [kompose8](https://github.com/digitalrebar/kompos8)
* [Tectonic](https://tectonic.com/)
* [Kraken](https://github.com/samsung-cnct/kraken)
* [NavOps Launch](https://www.navops.io/launch.html)
* [Photon Cluster Manager](https://github.com/vmware/photon-controller/tree/master/java/cluster-manager)
* [Platform 9](https://platform9.com/blog/containers-as-a-service-kubernetes-docker/)
* [GKE](https://cloud.google.com/container-engine/)
* [Stackpoint.io](https://stackpoint.io/)
* [Juju](https://jujucharms.com/canonical-kubernetes)
