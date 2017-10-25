# kubeadm upgrades proposal

Authors: Lucas Käldström & the SIG Cluster Lifecycle team
Last updated: October 2017

## Abstract

This proposal describes how kubeadm will support a upgrading clusters in an user-friendly and automated way using different techniques.
Eventually we aim to consolidate and unify upgrading procedures across different Kubernetes installer/deployment tools, but that is out of scope
to address for this initial proposal.

SIG Cluster Lifecycle is the responsible SIG for kubeadm and its functionality, and prioritized the making the upgrading functionality easier and more user-friendly during the v1.8 coding cycle (v1.6 -> v1.7 upgrades with kubeadm were supported, but in alpha and with a non-optimal UX).

## Graduation requirements

 - Beta in v1.8:
   - Support for the `kubeadm upgrade plan` and `kubeadm upgrade apply` commands.
   - Support for dry-running upgrades.
   - Supports upgrading clusters, but not 100% support for downgrades or "runtime reconfiguration"
   - Support for upgrading the API Server, the Controller Manager, the Scheduler, kube-dns & kube-proxy
   - Support for performing necessary post-upgrade steps like upgrading Bootstrap Tokens that were alpha in v1.6 & v1.7 to beta ones in v1.8
   - Automated e2e tests running.
 - GA in v1.10:
   - Support for downgrades and runtime reconfiguration.
   - Support for upgrading etcd.
   - The upgrading functionality is available from the `kubeadm phase` command.
   - Support for upgrading kubelets.
   - Supported Control Plane <-> kubelet minor skew: two minor versions.
   - Support for fully remote-executable upgrade operations against the cluster
   - Support for machine-readable `kubeadm plan` output
   - Support for upgrading HA/multi-master clusters

## Upgrading policies

The initial task that must be performed when considering to add functionality to seamlessly upgrade clusters, is to define policies for which kinds of upgrades
and version skews should be supported.

### Definitions

 - **Minor version upgrade**: An upgrade from a version (`vX.Y.Z1`) with the minor component equal to `Y` to a version (`vX.Y+1.Z2`) with the minor component of `Y+1`. Upgrading from `v1.8.3` to `v1.9.0` is a minor version upgrade.
 - **Patch version upgrade**: An upgrade from a version (`vX.Y.Z`) with the patch component equal to `Z` to a version (`vX.Y.Z+n`) with the patch component of `Z+n`. `n` can be freely chosen by the user. Upgrading from `v1.8.3` to `v1.8.6` is a patch version upgrade.
 - **Required upgrade policy**: An upgrade policy that must be satisfied; a precondition the user can't ignore.
 - **Skippable upgrade policy**: An upgrade policy that should be satisfied, but is skippable by passing the `--force` flag.
 - **Runtime reconfiguration**: An user invokes the `kubeadm upgrade apply` command with the current version, but with a different configuration.

### Pre-upgrade policies

 - The kubeadm CLI version must be higher or equal to the upgrade target version
   - **Required policy**: The kubeadm CLI minor version must be equal to or higher than the upgrade target minor version.
     - Example: kubeadm CLI version `v1.8.3` can upgrade a cluster to `v1.7.4` and `v1.8.5`, but not `v1.9.0`.
   - **Optional policy**: The kubeadm CLI patch version should be equal to or higher than the upgrade target patch version.
     - Example: kubeadm CLI version `v1.8.3` can upgrade a cluster to `v1.8.3` without an explicit grant, but upgrading to `v1.8.5` requires forcing. 
 - The kubeadm CLI of version `vX.Y.Z` can create clusters with minor versions `Y` and `Y-1`.
   - This is a **required policy**, and is not specific to upgrades. This policy is also is also enforced when creating a cluster.
     - Example: kubeadm CLI version `v1.8.2` can create clusters of (and upgrade clusters to) version `v1.8.x` and `v1.7.x`.
 - The control plane version must always be higher than or equal to the version of all the kubelets, and the maximum allowed skew is one minor versions.
   - **Required policy**: The control plane version must be higher than or equal to the kubelet version. This is a generic Kubernetes policy.
     - Example: For a control plane of version `v1.8.4`; a kubelet version may be `v1.8.2`, but not `v1.9.3`
   - **Skippable policy**: The maximum allowed minor skew between the control plane and the kubelets is one minor version.
     - Example: For a control plane of version `v1.7.2` being upgraded to `v1.8.5`; the minimum allowed kubelet version is `v1.7.0`.
     - Note: We might relax this policy to allow two minor versions before going to GA.
 - The version change must be positive, i.e. only upgrades are supported; not downgrades or runtime reconfiguration.
   - This is a **skippable policy** in v1.8, and will be removed in future versions before going to GA.
     - Example: Downgrading from `v1.8.3` to `v1.8.2` or reconfiguring a cluster on version `v1.8.1` can only be done (initially) along with the `--force` flag.
 - The maximum allowed upgrade minor skew is one minor version.
   - This is a **required policy**. You can upgrade from `vX.Y.Z1` to `vX.Y+1.Z2`, but not to `vX.Y+2.Z3`.
     - Example: Upgrading from `v1.6.2` to `v1.8.4` is not supported.
     - Note: We might relax this policy to allow two minor versions before going to GA.
 - The target upgrade version must not be an alpha, beta version unless the user explicitely allows it with `--allow-experimental-upgrades`
   - This is a **skippable policy**, in order to have an extra safety check but still allow it for users who may want it. Such upgrades are not supported.
     - Example: Upgrading from `v1.7.1` to `v1.8.0-beta.3` can only be done with at least one of the `--allow-experimental-upgrades` or `--force` flags.
 - The target upgrade version must not be a release candidate version unless the user explicitely allows it with `--allow-release-candidate-upgrades` (or `--allow-experimental-upgrades`)
   - This is a **skippable policy**, in order to have an extra safety check but still allow users to easily test our RCs. These upgrades "should work".
     - Example: Upgrading from `v1.7.1` to `v1.8.0-rc.1` can only be done with at least one of the `--allow-release-candidate-upgrades`, `--allow-experimental-upgrades` or `--force` flags.

## Initial scope of `kubeadm upgrade`

The initial scope of kubeadm upgrades in v1.8 is to upgrade the _control plane_ along with kube-proxy and kube-dns that are hosted on top of Kubernetes.
The control plane upgrade is atomic; either all or nocontrol plane components get upgraded. If a post-upgrade step fails though (like upgrading Bootstrap Tokens or applying a new kube-dns), kubeadm upgrade just returns a non-zero code instead of reverting the successful control plane upgrade. We might change the latter behavior in the future.

In the first iteration, upgrading the kubelets and/or etcd in the cluster is a manual task (see the sections below).
We will address these limitations in future versions before graduating this functionality to GA.

## Command structure

 - `kubeadm upgrade`: Placeholder command for other subcommands.
   - Relevant, generic flags:
     - `--config`: A path to a configuration file to use as the desired state for the cluster component configuration. If omitted, configuration is read from a ConfigMap in the cluster.
     - `--kubeconfig`: The path to a configuration file to use when talking to the cluster. Default: `/etc/kubernetes/admin.conf`.
     - `--allow-experimental-upgrades`: Allow the user to upgrade to alpha, beta and release candidate versions.
     - `--allow-release-candidate-upgrades`: Allow the user to upgrade to release candidate versions.
     - `--print-config`: Print out the configuration read from the file or ConfigMap in the cluster.
     - `--skip-preflight-checks`: Skip the preflight checks that are run before execution.
 - `kubeadm upgrade plan`: Serves as a way to show users:
   - If the cluster is in an upgradeable state.
   - What versions the user can upgrade the cluster to.
 - `kubeadm upgrade apply [version]`:
   - Is the command that actually performs the upgrade (of the control plane).
   - Will interactively ask the user for a `[Y/n]` confirmation if the upgrade really should be performed.
   - Relevant flags:
     - `--dry-run`: Only log the changes that would be made to the cluster without changing any state.
     - `--force`, `-f`: Ignore any skippable version policy errors and the interactive yes/no upgrade confirmation.
     - `--image-pull-timeout`: The time to wait for the new control plane images to pull before timing out. Default: 15 minutes.
     - `--yes`, `-y`: Run the command non-interactively, don't ask for an upgrade confirmation.

## Example user experience

UX when upgrading a patch version:

```console
$ kubeadm upgrade plan
[kubeadm upgrade plan output]
```

To upgrade the kubelets in the cluster, please upgrade them from the source you installed them from (most often using your package manager).

```console
$ kubeadm upgrade apply --version v1.8.3
[kubeadm upgrade apply output]
```

## Upgrade planning

In order for the user to know whether the cluster may be upgraded or not; the `kubeadm upgrade plan` command was written.
The command:

1. Checks if the cluster is healthy
  - Does the API server `/healthz` endpoint return `ok`?
  - Are all nodes in the cluster in the `Ready` state?
  - (If Static Pods are used) Do all required Static Pod manifests exist on disk?
  - (If cluster is self-hosted) Are all control plane DaemonSets healthy?
1. Builds an internal configuration object to apply to the cluster
  - The `--config` file takes precedence; otherwise the current configuration stored in a ConfigMap in the cluster is used.
  - The `KubernetesVersion` field is set to the desired version.
  - The built configuration object is optionally printed to STDOUT if `--print-config` is `true`.
1. Information about what releases are available are downloaded from the Kubernetes Release & CI GCS buckets.
  - The `dl.k8s.io/release/stable.txt` endpoint (and optionally `.../stable-1.X.txt`) is queried for the latest stable version to display.
  - The API Server, kubeadm CLI, and all kubelets' versions are taken into account when computing the possible upgrades.
  - The command may show zero to four upgrade possibilities.
  - The information is nicely printed to the user.


## Upgrade implementation

On a high level, this is what `kubeadm upgrade apply [version]` does:

1. Checks if the cluster is healthy
  - Does the API server `/healthz` endpoint return `ok`?
  - Are all nodes in the cluster in the `Ready` state?
  - (If Static Pods are used) Do all required Static Pod manifests exist on disk?
  - (If cluster is self-hosted) Are all control plane DaemonSets healthy?
1. Builds an internal configuration object to apply to the cluster
  - The `--config` file takes precedence; otherwise the current configuration stored in a ConfigMap in the cluster is used.
  - The `KubernetesVersion` field is set to the desired version.
  - The built configuration object is optionally printed to STDOUT if `--print-config` is `true`.
1. Enforces the version policies highlighted above in [Pre-upgrade policies](#pre-upgrade-policies)
1. Asks the user for a `[Y|n]` answer whether to proceed with the upgrade or not (if the session is interactive)
1. Performs a method of ensuring that all necessary control plane images are available locally to the master(s), aka. "pre-pulling"
  - Due to lack of an API to the kubelet that would make it possible to say "make sure this image X is available locally"; we have to
    implement the prepulling in a slightly worse, but still acceptable and pretty straightforward way.
  - One dummy DaemonSet per control plane component is created; and when all control plane components are running a dummy "sleep 3600" command
    kubeadm can for sure know that the new container image was successfully pulled or already cached locally; and the upgrade can then proceed.
1. Performs the actual upgrade. The implementation here depends on how the control plane was pulled, see the chapters below for more details.
1. Performs post-upgrade tasks like applying the new kube-dns and kube-proxy manifests and version-specific upgrade steps like upgrading Bootstrap Tokens from alpha in v1.7 to beta in v1.8. 

### Architectural upgrade implementation options

**Upgrade code in the CLI tool vs controller running in-cluster:**

On a high level, the upgrade could be implemented in two primary ways:

1. As synchronous code executing client-side in the kubeadm CLI binary.
1. As an asynchronous controller that handles the upgrade implementation; with the kubeadm CLI just invoking the upgrade

The historical preference was to build a reusable Operator that does the upgrading for us.
An important note here is that in *either way we would deliver the exact same 'kubeadm upgrade' UX.*

A head-to-head breakdown of pros/cons for the two methods outlined:

1. Build the upgrading code into the `kubeadm` CLI command
 -  Pros:
   - *Simplicity*; Easier and simpler to implement. A `kubeadm upgrade` command will be built in any case.
   - The [kubeadm v1.6 -> v1.7 alpha upgrade method](#TODO) used this option successfully
   - Being able to easily have an interactive flow where the user can inspect the changes that are about to be applied and accept/deny for instance.
   - No need to deal with version skew between the Operator image tag/version and the kubeadm CLI version
   - No need to define an API (CRD or the like) between the client and the Operator
 - Cons:
   - Maybe not as reusable as an Operator inside the cluster would be

1. Build an Operator
 - Pros:
   - The idea is that it would be generic and could be re-used by others
   - Other operators are already handling complex upgrades successfully, and we could even standardise patterns
 - Cons:
   - Would have to deal with version skew between the controller and the CLI
   - Would have to deal with pre-pulling the image to a specific node in the case of no internet connection
   - Would have to define an API (CRD or the like) between the client and the Operator (more time-consuming to get right)

**Decision**: Keep the logic inside of the kubeadm CLI (option 1) for the implementation in v1.8.0.

We *might* revisit this and build an upgrade operator in future, either before GA of this functionality or after, in v2.

**Rewrite manifests completely from config object or mutate existing manifest:**

Instead of generating new manifests from a versioned configuration object, we could try to add "filters" to the existing manifests and apply different filters depending on what the upgrade looks like. This approach, to modify existing manifests in various ways depending on the version bump has some pros (modularity, simple mutating functions), but the matrix of functions and different policies would grow just too big for this kind of system, so we voted against this alternative in favor for the solution above.


### Static Pod-hosted or Self-hosted control plane?

A control plane can be set up in a couple of ways:

 - Running the control plane components standalone; as systemd services for instance.
   - kubeadm **does not** support handling any such cluster
 - Running the control plane components in containers with kubelet as the babysitter; using Static Pods
   - [Static Pods](#TODO) is a way to run standalone Pods on a node by dropping Pod manifests in a specific path.
   - In kubeadm v1.4 to v1.8, this is the default way of setting up the control plane
 - Running the control plane in Kubernetes-hosted containers as DaemonSets; aka. [Self-Hosting](#TODO)
   - When creating a self-hosted cluster; kubeadm first creates a Static Pod-hosted cluster and then pivots to the self-hosted control plane.
   - This is the default way to deploy the control plane since v1.9; but the user can opt-out if it and stick with the Static Pod-hosted cluster. 

#### Static Pod-hosted control plane

When the upgrade functionality is introduced in v1.8; the default way to set up the control plane is by using Static Pods.

When the control plane is hosted using Static Pods, the actual upgrade flow looks like this:
 - Gets the Mirror Pod manifests from the API Server (`kubectl -n kube-system get pod ${component}-${nodename} -ojson`), and sha256-hashes the JSON content.
 - Writes new Static Pod manifests to a **temporary directory** (with the prefix `/etc/kubernetes/tmp/kubeadm-upgraded-manifests*`) based on the new configuration object built internally earlier.
 - Creates a **backup directory** with the prefix `/etc/kubernetes/tmp/kubeadm-backup-manifests*`.
   - Q: Why `/etc/kubernetes/tmp`?
   - A: Possibly not very likely, but we concluded that there may be an attack area for computers where `/tmp` is shared and writable by all users.
     We wouldn't want anyone to mock with the new Static Pod manifests being applied to the clusters. Hence we chose `/etc/kubernetes/tmp`, which is
     root-only owned.
 - In a loop for all control plane components:
   - Moves the **current manifest** (in `/etc/kubernetes/manifests`) to the **backup directory** (`/etc/kubernetes/tmp/kubeadm-backup-manifests*/`)
   - Moves the **upgraded manifest** (in `/etc/kubernetes/tmp/kubeadm-upgraded-manifests*`) to the **Static Pod manifest folder** (`/etc/kubernetes/manifests`)
     - => kubelet notices that the Static Pod on disk changed, kills the old Pod and creates the new one.
   - kubeadm waits for the API server to come up, executes `kubectl -n kube-system get pod ${component}-${nodename} -ojson`, sha256-hashes the response,
     and compares it to the pre-upgrade hash. If the hash has changed, kubeadm proceeds with the process.
     - _Race condition warning_: If the Pod manifest hashing wouldn't be performed, there is a very common case where kubeadm upgrades the Static Pod manifest
       on disk, but proceeds too quickly (before the kubelet has actually restarted the Pod), and falsely checks that the old component is running.
   - kubeadm waits for the Static Pod's Mirror Pod to register itself with the API Server; just as an extra security measure
 - In order to not leak directories on the host, kubeadm deletes the **backup** and **temporary directory**.
   - If there are user requests, we might make it possible to opt-out of the "always delete old manifests"-policy.

##### Rollbacks

In case any of these operations fail, kubeadm will rollback the earlier manifests to the previous state. Hence the backup directory is important.
For instance, if the scheduler doesn't come up cleanly; kubeadm will rollback the previously (successfully upgraded) API server, controller manager
manifests as well as the scheduler manifest.

#### Self-hosted control plane

You can read more about how kubeadm creates self-hosted clusters in detail in the [kubeadm self-hosting](kubeadm-selfhosting.md) document.

On a high level, each master component (kube-apiserver, kube-controller-manager and kube-scheduler) is deployed as a DaemonSet.
When a self-hosted cluster should be upgraded, the key advantage is that only access to the API server should technically be required.

In practice, access to the master's filesystem for disaster recovery via the Static Pod mechanism doesn't hurt.

Although we're aiming and planning to start supporting to create and upgrade clusters with multiple masters (HA clusters), that functionality didn't
make it into the v1.8 timeframe. Supporting to upgrade self-hosted, "single masters" is a definite requirement that would be a prereq in any case.

Upgrading self-hosted "single masters" is a bit tricky.
When modifying the DaemonSet specification, the `daemonset` controller in the controller-manager will kick in and remove the currently running Pod and
later create an upgraded Pod. We call the current behavior "delete first, then add".

This becomes problematic in the one-master case, as the `daemonset` controller will delete the only running replica of the component :(

There are three main ways to approach this problem:
 - Create a new DaemonSet upgrade strategy that would "add first, then delete"
   - This was SIG Cluster Lifecycle's preferred approach during the v1.8 timeframe, but the request got declined in the last minute, didn't make v1.8,
     and eventually was punted on in favor for graduating the current DaemonSet API to GA in v1.9 (which is fair anyway).
   - Pros:
     - foo
   - Cons:
     - bar
 - Create a temporary, duplicated DaemonSet of the component in question that will kick in while the "real" DaemonSet Pod is deleted and recreated.
   - This approach was implemented in the meantime for the `kubeadm upgrade` command in v1.8
   - Pros:
     - foo
   - Cons:
     - bar
 - Use Static Pods as the recovery mechanism while the DaemonSet is being upgraded for this "corner case"
   - Pros:
     - foo
   - Cons:
     - bar

## Various other notes

### Keeping user customizations between releases

One of the most important parts with implementing an upgrade is to respect customizations made by the user at `kubeadm init` time.

The proposed and implemented solution is to store the kubeadm configuration given at `init`-time in the API as a ConfigMap (can be seen with
`kubectl -n kube-system get configmap kubeadm-config -oyaml`), and later retrieve that configuration when upgrading. While that is the default
solution, the user may still explicitely pass a configuration file to use for the new desired state of the cluster.

### "Upgrading" without bumping the version number (runtime reconfiguration)

One interesting use-case for `kubeadm upgrade` is to modify the cluster configuration without bumping the version number.

As we have access to the current (in kubeadm's PoV) state of the cluster via the `kubeadm-config` ConfigMap; it would be possible to build a patch of the new changes and validate that patch... TODO


assume the configuration file options from `kubeadm init` will be fetchable from a ConfigMap inside of the cluster at any time, it would be possible to fetch the current state, compare it to the desired state (from `kubeadm apply --config`) and build a patch of it. Kubeadm would validate the fields changed in that patch and allow/deny "upgrading" those values.

This is a stretch goal for v1.8, but not a strictly necessary feature.

### Pre-pulling of images/ensuring the upgrade doesn’t take too long

One tricky part of the upgrade flow is that it will require new images to be run on a given node; and that image might not exist locally. When upgrading, the kubelet will start to pull the image (which is fine), but it might take a while. We don’t want the scenario of a poor internet connection when everything is proceeding fine but slow and `kubeadm upgrade` times out before the image is pulled.

We plan to solve this by having a separate `prepull` task in the beginning of the upgrade to minimize service outage. What it will basically do is to create a DaemonSet for all the control plane components, but instead of running the component, it executes `sleep 3600` or a similar command.

### Upgrading etcd

In v1.7 and v1.8, etcd runs in a Static Pod on the master. In v1.7, the default etcd version for Kubernetes is v3.0.17 and in k8s v1.8 the recommended version will be something like v3.1.10. In the v1.7->v1.8 upgrade path, we could offer upgrading etcd as well as an opt-in*. *This only applies to minor versions and 100% backwards-compatible upgrades like v3.0.x->v3.1.y, not backwards-incompatible upgrades like etcdv2 -> etcdv3.*


The easiest way of achieving an etcd upgrade is probably to create a Job (with `.replicas=<masters>`, Pod Anti-Affinity and `.spec.parallellism=1`) of some kind that would upgrade the etcd Static Pod manifest on disk by writing a new manifest, waiting for it to restart cleanly or rollback, etc.

### Verbose upgrading mode

Currently it is not possible to specify a `--v` flag to `kubeadm init`. In time for v1.8 that should be possible, as it is needed for `kubeadm upgrade` as well. Upgrading a cluster is a task that has to be handled carefully, and often the user wants to know what happens under the hood in order to trust the component that automates that work away.

### Network Plugins

As CNI network providers are crucial to the health of the cluster, we have to somehow ensure that the network plugin continues to work after a minor release upgrade.

Since we can’t assume that the same network provider manifest that worked in v1.7 will continue to work in v1.8 we might have to treat it in a special manner.

The "real" solution to this problem would be to implement an “[Addon Manager](https://github.com/kubernetes/features/issues/18)” in core Kubernetes as an API Group. Then CNI providers would just hook up themselves to that API and the addon manager would take care of using the right manifest for the right version.

One proposed method of dealing with this is to have a `--cni-network-manifests` flag to `kubeadm upgrade apply`. The flag would be optional for patch upgrades but semi-required with minor upgrades (skippable with `--force` or similar)

### Interactive vs non-interactive mode

By default, kubeadm should tell the user what it is doing in a way that is verbose enough to satisfy the most users (but doesn’t print out a wall of text). Also see the verbosity section.


### Possible failures, timeouts/deadlocks and rollback policy

By default, kubeadm should have a (configurable) timeout value; if a given task takes longer than; kubeadm will automatically try to rollback.


### Upgrading the kubelets

Foo

