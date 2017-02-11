# Cluster Lifecycle Deployment & Upgrade Roadmap

Moderator: Mike Danese

Note taker: Robert Bailey

Date: 2016-11-10

## Goals

Discuss HA, upgrades, and config management beyond kubeadm/kops & try to identify things that are currently underserved (upgrade testing, version skew policy, security upgrades)

## Discussion

kubeadm - not destined for production?
* Doing resource provisioning (cloud VMs) is out of scope
* Should be a toolbox that does the common parts of cluster lifecycle
   * And should be able to break out just the pieces that you want
* Found a bunch of common parts of existing cluster deployment and want to build more of it into the core

AI (luke): Create an intro guide to cluster lifecyle

If anyone wants to work on Upgrades the Hard Way with Rob this afternoon.

Wishlist
* HA
* Upgrades
* Config Management
* Toolbox vs. guided flow
* Documentation
* Conformance Testing
* PKI

### HA

* Story hasn't changed in a long time:
   * People set up the cluster, run them in production
   * Lack of documentation
* What haven't we been focused on?
   * Some day there may be apiservers that may move, but there aren't today
   * If you've misconfigure it, it's really hard to debug
      * E.g. If you just launch 2 apiservers they fight over the endpoint. It requires insider knowledge to recognize this and know how to fix it
   * Forces an ip address on the endpoint, which isn't compatible with AWS
      * AI (claytonc): Fix the flag to take a host instead of just an ip address
* There are things that are command line flags that are going to be a pain to synchronize
   * Move more configuration into etcd

### Upgrades

* Minor and patch upgrades, look at them separately
* Do we have skew requirements that are different for minor vs patch upgrades?
   * E.g. Can you upgrade nodes before master for a patch version?
* AI: Socialize the existing version skew documentation
* AI: Clarify the documentation skew
* Do we want to support 4-part version numbers?
   * Chris love: please don't do this
* Mike Rubin: Patch releases shouldn't create surprises
* Distribution of Kubernetes?
* Jordan Liggitt: Would like to see upgrade documentation on every install guide
   * At least for patch releases
   * AI (?): File issues against owners for the current getting started guides to add a section on upgrades
* Luke Marsden: I would like to lead an effort to support self hosting in some of the user flows (in particular a kubeadm flow) in an attempt to make it really easy to deploy a patch upgrade
   * Assume single master, external etcd
   * Joe Beda: The nice thing about self hosting upgrades is that there isn't anything cloud platform specific which allows us to build more general tooling

### Distribution

* In 1.5 we've begun experimenting with OS package management tooling
* Should we push further on this?
* The release tarball is getting bigger
   * Should be ameliorated in 1.5 by breaking it into arch-specific bundles
* Jordan Liggitt: We can't tell people to script their install against the tarball because the structure of the tarball becomes an api

### Config Management

* Config is currently command line flags. Work is being done to convert into structured api types (outstanding PRs from ncdc@ and mtoffen@).
* How much should we use config maps vs flags vs something else?
* Joe Beda: Definitely an issue for the kubelet - specifically setting the DNS IP flag
   * Vish: Unless/until the kubelet has local checkpointing it's dangerous to use the apiserver for checkpointing configuration since the apiserver may become unavailable
* Need to figure out how we deal with config that points to other files (e.g. certs)
* Rob: We may want to split the discussion about configuring the kubelet vs the control plane
* Mike Rubin: The kubelet will eventually need to be able to run standalone. Need to think about packaging and configuration as distinct.
   * The Kubelet has a lot of value if it can work without an apiserver
   * The node effort takes Docker + Linux and productizes it
* Mike Danese: The same type of config could also benefit other components
* Jordan Liggitt: Have client cert bootstrap stuff in Kubelet

### Toolbox vs guided flow

* [mostly skipped for time]
* Chris Love: Need to document our compartmentelizing each thing
* Luke: This isn't a "vs" it's an "and"

### Documentation

* What is lacking?
* Joe Beda: The fact that Kelsey had to write "Kubernetes the hard way" shows that we don't have documentation
* Chris Love: HA Upgrades
* Jordan Liggitt: Docs should look like a tree
   * Start at the high level, if you want more detail, then you can drill down into each piece
   * If you expand to all of the leaves, then you end up back at k8s the hard way
* Mike Rubin: Questions from support/users are less about setting up and more about tearing down
   * What will still be around after destroying a cluster
   * E.g. Deleting a namespace, deleting a cluster from a federation, deleting a node from the cluster
* Need an introduction to the SIG (Luke already volunteered to write one)
* Mike Rubin: Rollbacks and rollback documentation
   * When you add a new feature (say in 1.4) and we roll back to an earlier version what happens to those resources
   * Chris Love: elephant is what happens when you roll back from etcd3 → etcd2

### Conformance Testing

* What can we do in 2017 to make progress on this?
* Jordan Liggitt: Need to categorize conformance tests into ones that you could run against a production cluster vs those that you shouldn't
* Lucas: Three levels of validation: node, k8s standard base, deep/destructive testing. Want to make these all easy through kubeadm
* Is performance testing out of scope?
   * Clayton: Misconfiguration is often caught through performance testing so we shouldn't remove it from scope

### PKI

* Jordan Liggitt: Have client cert bootstrap stuff in Kubelet
* Chris Love: Need to loop in sig-auth. Need to use TLS certs for etcd clusters.
* Aaron Levy: Plan to add CSR into the etcd operation similar to what is going into the k8s api. 
* Joe Beda: Two modes right now: can have the apiserver act as a CA; many serious users will want to use their own CA
* Jordan Liggitt: Things that need certs should be able to take them or components should be able to generate (if appropriate)
   * Rotation depends on whether we are using the built-in CA or an external CA
* Mike Rubin: Why not do both rotation and revocation
   * Rob: Many applications don't respect revocation so it's generally considered weaker
* Clayton: If you can rotate then you may not need to revoke
* Jordan Liggitt: Tied to config management
* Joe Beda: Part of the discovery info is the root CA and many people don't realize that it can be a bundle instead of a single CA — this enables rotation
* Clayton: In a secured cluster, etcd is the core. Have to think about it as the inner circle of security that the apiserver is outside of. If you are extremely cautious then you should use client certs in the apiserver. You can collapse the rings if you want.

