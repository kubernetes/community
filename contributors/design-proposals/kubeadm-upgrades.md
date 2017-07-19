kubeadm upgrades proposal

Lucas Käldström, Luke Marsden & the SIG Cluster Lifecycle team
July 2017

See also: [sig-cluster-lifecycle 1.7 Planning](https://docs.google.com/document/d/1xpjz9373KQ56gr7812eo85wB7G6n39LtUIn9Lq74t0Q/edit) and [sig-cluster-lifecycle 1.8 Planning](https://docs.google.com/document/d/17C2D1Ghgv40pWByevo0ih8T-18Q1GlxbR4gu1_ZKlPE/edit).

# Abstract

This proposal describes how kubeadm will support upgrades via self-hosting, and includes enough design for HA to ensure that we're not painting ourselves into an architectural corner.

This document is somewhat based on the [notes](https://docs.google.com/document/d/19GzaaqqW7Q6IdfJa6gYy2JsCgnmChTmol4V0kHI50pc/edit) from the face-to-face SIG-cluster-lifecycle meeting during KubeCon 2017 in Berlin and the "self-hosting, upgrades & HA"-special effort SIG Cluster Lifecycle meetings late June and early July 2017.

Timescale for implementation of this proposal (working self-hosting & upgrades) is Kubernetes 1.8. This means that [features issues must be in by Aug 1](https://github.com/kubernetes/features/pull/305) and [code must be in by Sept 1](https://github.com/kubernetes/features/pull/305).

# Self-hosting

Self-hosting will be the default way of deploying the control plane, but it will be optional. There will be a `kubeadm init`-time feature gate to disable it.

Only *control plane components* will be in-scope for self-hosting for v1.8.

Notably, not etcd or kubelet for the moment.

Self-hosting will be a [phase](https://github.com/kubernetes/kubeadm/blob/master/docs/design/design.md) in the kubeadm perspective. As there are Static Pods on disk from earlier in the `kubeadm init` process, kubeadm can pretty easily parse that file and extract the PodSpec from the file. This PodSpec will be modified a little to fit the self-hosting purposes and be injected into a DaemonSet which will be created for all control plane components (API Server, Scheduler and Controller Manager). kubeadm will wait for the self-hosted control plane Pods to be running and then destroy the Static Pod-hosted control plane.

What actually happens in the Static Pod -> Self-hosted transition?

First the PodSpec is modified. For instance, instead of hard-coding an IP for the API server to listen on it is dynamically fetched from the Downward API. When it’s ok to proceed, the self-hosted control plane DaemonSets are posted to the Static Pod API Server. The Static Pod API Server starts the self-hosted Pods normally. The Pods get in the Running state, but fails to bind to the port on the host and backs off internally. During that time, kubeadm notices the Running state of the self-hosted Pod, and deletes the Static Pod file which leads to that the kubelet stops the Static Pod-hosted component immediately. On the next run, the self-hosted component will try to bind to the port, which now is free, and will succeed. A self-hosted control plane has been created!

Status as of v1.7: self hosted config file option exists; not split out into phase. The code is fragile and uses Deployments for the scheduler and controller-manager, which unfortunately [leads to a deadlock](https://github.com/kubernetes/kubernetes/issues/45717) at `kubeadm init` time. It does not leverage checkpointing so the entire cluster burns to the ground if you reboot your computer(s). Long story short; self-hosting is not production-ready in 1.7.

The rest of the document assumes a self-hosted cluster. It will not be possible to use 'kubeadm upgrade' a cluster unless it's self-hosted. The sig-cluster-lifecycle team will still provide some manual, documented steps that an user can follow to upgrade a non-self-hosted cluster.

Self-hosting implementation for v1.8

The API Server, controller manager and scheduler will all be deployed in DaemonSets with node affinity (using `nodeSelector`) to masters. The current way DaemonSet upgrades currently operate is "remove an old Pod, then add a new Pod". Since supporting “single masters” is a definite requirement, we have to either workaround removing the only replica of the scheduler/controller-manager by duplicating the existing DaemonSets (e.g. a `temp-self-hosted-kube-apiserver` DS will be created as a copy of the normal `self-hosted-kube-apiserver` DS during the upgrade) or ask for a [new upgrade strategy](https://github.com/kubernetes/kubernetes/issues/48841) (“add first, then delete”, which is generally useful) from sig-apps.

Checkpointing

In order to be able to reboot a self-hosted cluster (e.g. a single, self-hosted master), there has to be some kind of checkpointing mechanism. Basically the kubelet will write some state to disk for Pods that opt-in to checkpointing ("I’m running an API Server, let’s write that down so I remember it", etc.). Then if the kubelet reboots (in the single-master case for example), it will check the state store for Pods it was running before the reboot. It discovers that it ran an API Server, Controller Manager and Scheduler and starts those Pods now as well.

This solves the chicken-and-egg problem that otherwise would occur when the kubelet comes back up, where the kubelet tries to connect to the API server, but the API server hasn’t been started yet, it should be running as a Pod on that kubelet, although it isn’t aware of that.

We’re aiming to build this functionality into the kubelet behind an alpha feature flag.

# Upgrading policies

Defining decent upgrading and version skew policies is important before implementing.

Definitions

Minor version upgrade = This is an upgrade from a version (vX.Y.Z1) with the minor component equal to `Y` to a version (vX.Y+1.Z2) with the minor component of `Y+1`. For instance, a minor version upgrade can look like `v1.8.3` to `v1.9.0`.

Patch version upgrade = This is an upgrade from a version (vX.Y.Z) with the patch component equal to `Z` to a version (vX.Y.Z+n) with the patch component of `Z+n`. `n` can be freely chosen by the user. For instance, a patch version upgrade can look like `v1.8.3` to `v1.8.6`.

1. The user must upgrade the kubeadm CLI before the control plane

    1. Upgrading to a higher patch release than the kubeadm CLI version will
*be possible*, but yield a warning and require the `--force` flag to `kubeadm upgrade apply` to continue

    2. Upgrading to a higher minor release than the kubeadm CLI version will **not** be supported.

    3. Example: kubeadm v1.8.3 can upgrade your v1.8.3 cluster to v1.8.6 if you specify `--force` at the time of the upgrade, but kubeadm can never upgrade your v1.8.3 cluster to v1.9.0

2. The kubeadm CLI can create clusters with the latest two minor releases.

    4. Example: kubeadm v1.8.2 supports *creating a new cluster* for Kubernetes versions v1.7.x and v1.8.x

3. The control plane must be upgraded before the kubelets in the cluster

    5. For kubeadm; the maximum amount of skew between the control plane and the kubelets is *one minor release*.

    6. Example: running `kubeadm upgrade apply --version v1.9.0` against a v1.8.2 control plane will error out if the nodes are still on v1.7.x

4. This means that there are possibly two kinds of upgrades kubeadm can do:

    7. A patch release upgrade (from v1.8.1 to v1.8.3 for instance)

        1. This is very much about bumping the version number. It is very uncommon that things change in a breaking manner between these versions

    8. A minor release upgrade (from v1.7.5 to v1.8.3 for instance)

        2. This kind of upgrade is much harder as lots of things tend to change with every minor version. Common minor release upgrade (from v1.7.5 to v1.8.3 for instance) tasks kubeadm must handle:

            1. Old/deprecated flags that should be removed

                1. Example: the --insecure-experimental-approve-all-node-csrs flag for the controller-manager

            2. Arguments that have to be changed

                2. Example: the list of admission controllers grew from v1.6 to v1.7; two new admission controllers (Initializers and NodeRestriction) had to be inserted in the right place of the chain

            3. New flags that should be added

                3. In order to enable new and required features, kubeadm should automatically append these config options

            4. Resources that have to be modified

                4. Example: The `system:nodes` ClusterRoleBinding had to lose its binding to the `system:nodes` Group when upgrading to v1.7; otherwise the Node Authorizer wouldn’t have had any effect.

5. Using kubeadm, you must upgrade the control plane atomically.

    9. It’s out of scope for now to upgrade only a part of the control plane. All  control plane components must be upgraded at the same time for it to be a supported operation from kubeadm’s point of view.

    10. If you upgrade your cluster, the apiserver, controller-manager, scheduler, proxy and dns will get upgraded at the same time, and necessary resources (like RBAC rules) are modified in-cluster.

        3. Kube-dns will be upgraded to the latest available, branched, manifest inside of the kubeadm source code.

6. `kubeadm upgrade plan` will fetch the latest stable versions from `dl.k8s.io/release/stable-1.X.txt` where X is the minor version for the given release branch.

    11. If the user can’t connect to the internet, `kubeadm upgrade plan` won’t show the user which versions he/she can upgrade to, but `kubeadm upgrade apply --version` will work just normally. 

    12. The only thing that changes is that the user has to know which version to upgrade to, but that is a good requirement, since the user is the one that knows which images has been pre-pulled anyway.

7. Upgrading etcd might or might not be supported, depending on the complexity of the upgrade. 

    13. Major version bumps with schema changes require a proposal on its own

    14. If the user opts-in to it, and kubeadm knows that a higher Kubernetes version than what the user have now requires or recommends a minor/patch version bump of etcd, kubeadm may be able to upgrade etcd via a Job (see the related section below)

8. An upgrade *could* imply other things than a version number bump, for example 

# The `kubeadm upgrade` command

There will exist a `kubeadm upgrade` command that takes care of all the hard parts involved in upgrading a Kubernetes cluster.

It will have two subcommands:

* `plan`: Usually the first command the user will run (the recommended).

    * Validates that the cluster is upgradable

    * Outputs available versions to upgrade to (if any)

    * In an air-gapped environment, it won’t show the available versions to upgrade to.

* `apply`: Will do the actual upgrade

    * `--dry-run`: Computes the changes/actions that would have to be taken and outputs the manifests that will be applied, resources that will be modified, etc.

UX when upgrading a patch version:

$ kubeadm upgrade plan

Inspecting current cluster state:

 -> Self-hosting active ✓

 -> Configuration exists in-cluster ✓

 --> Validating configuration ✓

 -> Fetching cluster versions ✓

 --> Control plane version: v1.8.1

 --> Latest stable patch version: v1.8.3

 --> Latest stable minor version: v1.8.3

 --> Kubeadm CLI version: v1.8.3

 ---> Good! Kubeadm is up-to-date!

 Upgrade to the latest patch version of your release branch:

			Current version	Upgrade available

 API server:		v1.8.1		v1.8.3

 Controller-manager:	v1.8.1		v1.8.3

 Scheduler:		v1.8.1		v1.8.3

 Kube-proxy:		v1.8.1		v1.8.3

 Kube-dns:			v1.15.4		v1.15.5

 Etcd				v3.1.10		v3.1.10

 Upgrades available!

 You can now apply the upgrade by executing the following command:

	kubeadm upgrade apply --version v1.8.3

 Components that must be upgraded manually:

 Kubelet			2 x v1.8.0		v1.8.3

				3 x v1.8.1		v1.8.3

To upgrade the kubelets in the cluster, please upgrade them from the source you installed them from (most often using your package manager).

$ kubeadm upgrade apply --version v1.8.3

Pulling the required images... ✓

* gcr.io/google_containers/kube-apiserver: 2/2 pulled

* gcr.io/google_containers/kube-controller-manager: 2/2 pulled

* gcr.io/google_containers/kube-scheduler: 2/2 pulled

* gcr.io/google_containers/kube-proxy: 5/5 pulled

Upgrading the API Server DaemonSet... 

* Applying the new version of the DaemonSet... ✓

* Verifying that it comes up cleanly... ✓

Upgrading the Controller-manager DaemonSet...

* Applying the new version of the DaemonSet... ✓

* Verifying that it comes up cleanly... ✓

Upgrading the Scheduler DaemonSet...

* Applying the new version of the DaemonSet... ✓

* Verifying that it comes up cleanly... ✓

Upgrading the kube-proxy DaemonSet... ✓

Upgrading the kube-dns Deployment... ✓

Success! Your cluster was upgraded to v1.8.3

UX when upgrading a minor version:

$ kubeadm upgrade plan

Inspecting current cluster state:

 -> Self-hosting active ✓

 -> Configuration exists in-cluster ✓

 --> Validating configuration ✓

 -> Fetching cluster versions ✓

 --> Control plane version: v1.7.5

 --> Latest stable patch version: v1.7.8

 --> Latest stable minor version: v1.8.3

 --> Kubeadm CLI version: v1.8.1

 ---> Note! The kubeadm CLI should be upgraded
      to v1.8.3 if you want to upgrade your
      Cluster to v1.8.3

 Both minor and patch release upgrades available!

 Upgrade to the latest patch version of your branch:

			Current version	Upgrade available

 API server:		v1.7.5		v1.7.8

 Controller-manager:	v1.7.5		v1.7.8

 Scheduler:		v1.7.5		v1.7.8

 Kube-proxy:		v1.7.5		v1.7.8

 Kube-dns:			v1.14.2		v1.14.4

 Etcd:			v3.0.17		v3.1.10

 You can now apply the upgrade by executing the following command:

	kubeadm upgrade apply --version v1.7.8

Upgrade to the latest stable minor release:

			Current version	Upgrade available

 API server:		v1.7.5		v1.8.3

 Controller-manager:	v1.7.5		v1.8.3

 Scheduler:		v1.7.5		v1.8.3

 Kube-proxy:		v1.7.5		v1.8.3

 Kube-dns:			v1.14.2		v1.15.5

 You can now apply the upgrade by executing the following command:

	kubeadm upgrade apply --version v1.8.3

 Components that must be upgraded manually:

 Kubelet			2 x v1.8.0		v1.8.3

				3 x v1.8.1		v1.8.3

# Implementation options of the upgrade

High-level approaches:

1. Build the upgrade logic into kubeadm

2. Build an upgrading Operator that does the upgrading for us

    * Would consume a TPR with details how to do the upgrade

The historical preference was to build a reusable Operator that does the upgrading for us. *Either way we would deliver the same 'kubeadm upgrade' UX.* In the latter case, that would involve kubeadm deploying an upgrading controller on behalf of the user, and tracking its lifecycle through the kubernetes API.

There are pros and cons for each method:

1. Build it into the `kubeadm` CLI command

    1. Pros:

        1. *Simplicity*; Very straightforward to implement; basically doing the same steps as `kubeadm init` would, only slightly modified to fit the upgrading purpose. We will construct a `kubeadm upgrade` command in any case

        2. The kubeadm v1.6 -> v1.7 upgrade used this method this successfully

        3. Being able to easily have an interactive flow where the user can inspect the changes that are about to be applied and accept/deny for instance.

        4. Doesn’t have to deal with any version skew between the Operator image tag/version and the CLI version

        5. Doesn’t have to define an API (CRD or the like) between the client and the Operator

    2. Cons:

        6. Maybe not as reusable as an Operator inside the cluster would be

2. Build an Operator

    3. Pros:

        7. The idea is that it would be generic and could be re-used by others

        8. Other operators are already handling complex upgrades successfully, and we could even standardise patterns

    4. Cons:

        9. Has to deal with any version skew between the controller and the CLI

        10. Has to deal with pre-pulling the image to a specific node in the case of no internet connection

        11. Has to define an API (CRD or the like) between the client and the Operator

**Decision**: Keep the logic inside of the kubeadm CLI (option 1) for the implementation in v1.8.0.

We *might* build an upgrade operator in future kubeadm versions.

Upgrade a Static Pod-hosted cluster to a self-hosted one

The first thing `kubeadm upgrade` would have to do when upgrading from v1.7 to v1.8 would be to make the static-pod-hosted v1.7 cluster self-hosted. This can easily be done thanks to the recent refactoring of self-hosting into an atomically invoked phase. 
`kubeadm upgrade plan` will inform the user about the Static Pod => self-hosted switch in the case of an v1.7.x to v1.8.x upgrade. It will also be possible to convert a Static Pod-hosted control plane to a self-hosted one with another command like `kubeadm phase selfhosting apply`.

Keeping user customizations between releases

One of the hardest parts with implementing the upgrade will be to respect customizations made by the user at `kubeadm init` time. The proposed solution would be to store the kubeadm configuration given at `init`-time in the API as a ConfigMap, and then retrieve that configuration at upgrade time, parse it using the API machinery and use it for the upgrade.

This highlights a very important point: **We have to get the kubeadm configuration API group to Beta (v1beta1) in time for v1.8.**

By declaring Beta status on kubeadm’s configuration, we can be sure that kubeadm CLI v1.9 can read it and be able to upgrade from v1.8 to v1.9.

### Alternative considered

Instead of generating new manifests from a versioned configuration object, we could try to add "filters" to the existing manifests and apply different filters depending on what the upgrade looks like. This approach, to modify existing manifests in various ways depending on the version bump has some pros (modularity, simple mutating functions), but the matrix of functions and different policies would grow just too big for this kind of system, so we voted against this alternative in favor for the solution above.

"Upgrading" without bumping the version number

One interesting use-case for `kubeadm upgrade` will be to modify other configuration in the cluster. If we assume the configuration file options from `kubeadm init` will be fetchable from a ConfigMap inside of the cluster at any time, it would be possible to fetch the current state, compare it to the desired state (from `kubeadm apply --config`) and build a patch of it. Kubeadm would validate the fields changed in that patch and allow/deny "upgrading" those values.

This is a stretch goal for v1.8, but not a strictly necessary feature.

Pre-pulling of images/ensuring the upgrade doesn’t take too long

One tricky part of the upgrade flow is that it will require new images to be run on a given node; and that image might not exist locally. When upgrading, the kubelet will start to pull the image (which is fine), but it might take a while. We don’t want the scenario of a poor internet connection when everything is proceeding fine but slow and `kubeadm upgrade` times out before the image is pulled.

We plan to solve this by having a separate `prepull` task in the beginning of the upgrade to minimize service outage. What it will basically do is to create a DaemonSet for all the control plane components, but instead of running the component, it executes `sleep 3600` or a similar command.

Upgrading etcd

In v1.7 and v1.8, etcd runs in a Static Pod on the master. In v1.7, the default etcd version for Kubernetes is v3.0.17 and in k8s v1.8 the recommended version will be something like v3.1.10. In the v1.7->v1.8 upgrade path, we could offer upgrading etcd as well as an opt-in*. *This only applies to minor versions and 100% backwards-compatible upgrades like v3.0.x->v3.1.y, not backwards-incompatible upgrades like etcdv2 -> etcdv3.

The easiest way of achieving an etcd upgrade is probably to create a Job (with `.replicas=<masters>`, Pod Anti-Affinity and `.spec.parallellism=1`) of some kind that would upgrade the etcd Static Pod manifest on disk by writing a new manifest, waiting for it to restart cleanly or rollback, etc.

Verbose upgrading mode

Currently it is not possible to specify a `--v` flag to `kubeadm init`. In time for v1.8 that should be possible, as it is needed for `kubeadm upgrade` as well. Upgrading a cluster is a task that has to be handled carefully, and often the user wants to know what happens under the hood in order to trust the component that automates that work away.

Network Plugins

As CNI network providers are crucial to the health of the cluster, we have to somehow ensure that the network plugin continues to work after a minor release upgrade.

Since we can’t assume that the same network provider manifest that worked in v1.7 will continue to work in v1.8 we might have to treat it in a special manner.

The "real" solution to this problem would be to implement an “[Addon Manager](https://github.com/kubernetes/features/issues/18)” in core Kubernetes as an API Group. Then CNI providers would just hook up themselves to that API and the addon manager would take care of using the right manifest for the right version.

One proposed method of dealing with this is to have a `--cni-network-manifests` flag to `kubeadm upgrade apply`. The flag would be optional for patch upgrades but semi-required with minor upgrades (skippable with `--force` or similar)

Interactive vs non-interactive mode

By default, kubeadm should tell the user what it is doing in a way that is verbose enough to satisfy the most users (but doesn’t print out a wall of text). Also see the verbosity section.

`kubeadm upgrade apply` implementation flow

1. Validate the version given and configuration. Fail fast if the configuration is missing or invalid. Fail fast if the user is trying to upgrade to a version that is not supported.

2. Make sure the cluster is healthy

    1. Make sure the API Server’s `/healthz` endpoint returns `ok`

    2. Makes sure all Nodes return `Ready` status

    3. Makes sure all desired self-hosting Pods for the control plane are Running and Ready

3. (Upgrade the cluster to use self-hosting if static pods are used currently)

    4. Might require some extra flags, etc

4. Parse the control plane arguments into `map[string]string`, traverse them and create a new alternative.

5. Upgrade the API Server, Controller-manager and Scheduler in turn

    5. This will result in a very short API Server outage (will post here with information about how many milliseconds) if you have a single master (which is the case with v1.7)

    6. If any of them don’t come up cleanly, kubeadm will rollback.

6. Modify resources in the API as needed (e.g. RBAC rules)

7. Upgrade kube-proxy and kube-dns

8. Apply the network manifest if specified

9. Update the ConfigMap with the new config file that was used

Possible failures, timeouts/deadlocks and rollback policy

By default, kubeadm should have a (configurable) timeout value; if a given task takes longer than; kubeadm will automatically try to rollback.

# HA design

Note: HA being fully implemented and on by default for v1.8 is a *non-goal*.

Design docs are available:

* [Kubeadm (HA) - @timothysc](https://docs.google.com/document/d/1lH9OKkFZMSqXCApmSXemEDuy9qlINdm5MfWWGrK3JYc/edit#)

* [kubeadm HA implementation plan - @luxas](https://docs.google.com/document/d/1ff70as-CXWeRov8MCUO7UwT-MwIQ_2A0gNNDKpF39U4/edit#)
