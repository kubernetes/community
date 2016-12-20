# Proposal: Self-hosted Control Plane

Author: Brandon Philips <brandon.philips@coreos.com>
Last Updated: 2016-12-20

## Motivations

> Running in our components in pods would solve many problems, which we'll otherwise need to implement other, less portable, more brittle solutions to, and doesn't require much that we don't need to do for other reasons. Full self-hosting is the eventual goal.

- Brian Grant ([ref](https://github.com/kubernetes/kubernetes/issues/4090#issuecomment-74890508))

### What is self-hosted?

Self-hosted Kubernetes runs all required and optional components of a Kubernetes cluster on top of Kubernetes itself.

The advantages of a self-hosted Kubernetes cluster are:

1. **Small Dependencies:** self-hosted should reduce the number of components required, on host, for a Kubernetes cluster to be deployed to a Kubelet (ideally running in a container). This should greatly simplify the perceived complexity of Kubernetes installation.
2. **Deployment consistency:** self-hosted reduces the number of files that are written to disk or managed via configuration management or manual installation via SSH. Our hope is to reduce the number of moving parts relying on the host OS to make deployments consistent in all environments.
3. **Introspection:** internal components can be debugged and inspected by users using existing Kubernetes APIs like `kubectl logs`
4. **Cluster Upgrades:** Related to introspection the components of a Kubernetes cluster are now subject to control via Kubernetes APIs. Upgrades of Kubelet's are possible via new daemon sets, API servers can be upgraded using daemon sets and potentially deployments in the future, and flags of add-ons can be changed by updating deployments, etc. (An example script is in progress.)

However, there is a spectrum of ways that a cluster can be self-hosted. To do this we are going to divide the Kubernetes cluster into a variety of layers beginning with the Kubelet (level 0) and going up to the add-ons (Level 4). A cluster can self-host all of these levels 0-4 or only partially self-host.

![](self-hosted-layers.png)

For example, a 0-4 self-hosted cluster means that the kubelet is a daemon set, the API server runs as a pod and is exposed as a service, and so on. While a 1-4 self-hosted cluster would have a system installed Kubelet.

## Practical Implementation Overview

This document outlines the current implementation of "self-hosted Kubernetes" installation and upgrade of Kubernetes clusters based on the work that the teams at CoreOS and Google have been doing. The work is motivated by the upstream "Self-hosted Proposal".

The entire system is working today and is used by Bootkube, a Kubernetes Incubator project, and all Tectonic clusters created since July 2016. This document outlines the implementation, not the experience. The experience goal is that users not know all of these details and instead get a working Kubernetes cluster out the other end that can be upgraded using the Kubernetes APIs.

The target audience of this document are others, like [kubeadm](https://github.com/kubernetes/kubernetes/pull/38407), thinking about and building the way forward for install and upgrade of Kubernetes. If you want a higher level demonstration of "Self-Hosted" and the value see this [video and blog](https://coreos.com/blog/self-hosted-kubernetes.html).


### Bootkube

Today, the first component of the installation of a self-hosted cluster is [`bootkube`](https://github.com/kubernetes-incubator/bootkube). A kubelet connects to the temporary Kubernetes API server provided by bootkube and is told to deploy the required Kubernetes components, as pods. This diagram shows all of the moving parts:

![](self-hosted-moving-parts.png)

At the end of this process the bootkube can be shut down and the system kubelet will coordinate, through a POSIX lock, to let the self-hosted kubelet take over lifecycle and management of the control plane components. The final cluster state looks like this:

![](self-hosted-final-cluster.png)


### Bootkube Challenges

This process has a number of moving parts. Most notably the hand off of control from the "host system" to the Kubernetes self-hosted system. And things can go wrong:

1) The self-hosted Kubelet is in a precarious position as there is no one around to restart the process if it crashes. The high level is that the system init system will watch for the Kubelet POSIX lock and start the system Kubelet if the lock is missing. Once the system Kubelet starts it will launch the self-hosted Kubelet.

2) Recovering from reboots of single-master installations is a challenge as the Kubelet won't have an API server to talk to to restart the self-hosted components. We are solving this today with "[user space checkpointing](https://github.com/kubernetes-incubator/bootkube/tree/master/cmd/checkpoint#checkpoint)" container in the Kubelet pod that will periodically check the pod manifests and persist them to the static pod manifest directory. Longer term we would like for the kubelet to be able to checkpoint itself without external code.

## Long Term Goals

Ideally bootkube disappears over time and is replaced by a [Kubelet pod API](https://github.com/kubernetes/kubernetes/issues/28138). The write API would enable an external installation program to setup the control plane of a self-hosted Kubernetes cluster without requiring an existing API server.

[Checkpointing](https://github.com/kubernetes/kubernetes/issues/489) is also required to make for a reliable system that can survive a number of normal operations like full down scenarios of the control plane. Today, we can sufficiently do checkpointing external of the Kubelet process, but checkpointing inside of the Kubelet would be ideal.

A simple updater can take care of helping users update from v1.3.0 to v1.3.1, etc over time.

### Self-hosted Cluster Upgrades

#### Kubelet upgrades

The kubelet could be upgraded in a very similar process to that outlined in the self-hosted proposal.

However, because of the challenges around the self-hosted Kubelet (see above) Tectonic has implemented an alternative scheme that side-steps the self-hosted Kubelet challenges. First, a kubelet system service is launched that uses the [chrooted kubelet](https://github.com/kubernetes/community/pull/131) implemented by the [kubelet-wrapper](https://coreos.com/kubernetes/docs/latest/kubelet-wrapper.html) then when an update is required a daemonset updates the kubelet-wrapper configuration based on version annotations and kills the kubelet PID triggering a restart.

#### API Server, Scheduler, and Controller Manager

Upgrading these components is fairly straightforward. They are stateless, easily run in containers, and can be modeled as pods and services. Upgrades are simply a matter of deploying new versions, health checking them, and changing the service label selectors.

#### etcd self-hosted

As the primary data store of Kubernetes etcd plays an important role. Today, etcd does not run on top of the self-hosted cluster. However, progress is being made with the introduction of the [etcd Operator](https://coreos.com/blog/introducing-the-etcd-operator.html) and integration into [bootkube](https://github.com/kubernetes-incubator/bootkube/blob/848cf581451425293031647b5754b528ec5bf2a0/cmd/bootkube/start.go#L37). 

### Conclusions

Kubernetes self-hosted is working today. Bootkube is an implementation of the "temporary control plane" and this entire process has been used by [`bootkube`](https://github.com/kubernetes-incubator/bootkube) users and Tectonic since the Kubernetes v1.4 release. We are excited to give users a simpler installation flow and sustainable cluster lifecycle upgrade/management.

## Known Issues

- [Health check endpoints for components don't work correctly](https://github.com/kubernetes-incubator/bootkube/issues/64#issuecomment-228144345)
- [kubeadm doesn't do self-hosted yet](https://github.com/kubernetes/kubernetes/pull/38407)
