# Proposal: local cluster tool based on kubeadm and Docker-in-Docker

## Abstract

The `kubeadm-local-cluster` tool creates local multinode Kubernetes
clusters using `kubeadm` tool and the Docker-in-Docker technique. Most
of its functionality is already implemented in Mirantis'
[kubeadm-dind-cluster](https://github.com/Mirantis/kubeadm-dind-cluster)
project. The plan is to rewrite the shell scripts that comprise
`kubeadm-dind-cluster` in Go language.

The `kubeadm-local-cluster` tool supports building Kubernetes from
source, but also can utilize prebuilt binaries. The tool can work
against both remote and local docker engines. It's well suited for
constrained environments such as CI jobs running inside VMs
(e.g. Travis CI) because it doesn't create any VMs for the cluster,
thus avoiding problems with nested
virtualization. `kubeadm-local-cluster` works on Linux and Mac.

Unlike `hack/local-up-cluster.sh` and `minikube`,
`kubeadm-local-cluster` supports starting _multinode_ clusters that
can pass conformance e2e tests. Unlike `minikube`,
`kubeadm-local-cluster` supports building Kubernetes from source
easily and doesn't create VMs.

## Background and Motivation

There's a number of tools that can be used to start local clusters,
most well known being
[minikube](https://github.com/kubernetes/minikube) and
`hack/local-up-cluster.sh`. `minikube` covers the needs of many
application developers and `hack/local-up-cluster.sh` is quite helpful
if you're working on Kubernetes itself or a project that extends
Kubernetes functionality by adding Third Party Resources, controllers,
implementing CRI and so on. There's also `vagrant` provider in
`cluster/` directory of Kubernetes source, but it's very slow, becomes
broken at times and on top of that the whole `cluster/` subtree is
considered deprecated.

Despite the availability of said tools there's a number of use cases
that unfortunately isn't covered well enough by them:

* Multiple nodes are required to run some of the e2e tests. Multiple
  nodes may also be necessary if the user wants to try out some of the
  Kubernetes features like node affinity, pod (anti)affinity and so on.
* Running e2e tests locally with ease. It should be easy for
  developers _of_ Kubernetes to invoke e2e tests against local
  cluster.
* Realistic clusters. `hack/local-up-cluster.sh` has little
  resemblance to production clusters in how it runs Kubernetes
  components. `minikube` utilizes `localkube` wrapper which is also
  something specific to local cluster setups.
* Using different CNI implementations. In some cases it may be useful
  to debug CNI implementations themselves, e.g. use local cluster
  setup to investigate a network issue specific to particular CNI
  plugin such as `weave`, `flannel` or `calico`.
* Using locally-built Kubernetes binaries, with ability to quickly
  relaunch the cluster with new binaries. This can be used for quick
  local testing of Kubernetes functionality you work on or for using
  `git bisect` to track down a problem in k8s. This is covered by
  `hack/local-up-cluster.sh` but not `minikube`. So for instance if
  your project became broken by newer k8s version and it can't run in
  the environment provided by `hack/local-up-cluster.sh`, `git bisect`
  will not help you much.
* Running in a constrained CI environment such as Travis
  CI. `minikube` is not well-suited for this because this use case may
  entail the need for nested virtualization which is disabled by cloud
  operators most of the time. `hack/local-up-cluster.sh` requires
  building Kubernetes from source which adds to the build time and
  complexity of CI setup.
* Local Kubernetes clusters that don't hog machine's resources. Both
  building Kubernetes from source and running VMs can place
  substantial load on one's system. In case of Mac, most k8s users are
  likely to run Mac Docker anyway, and on 8 Gb machine having another
  VM for `minikube` may cause noticeable slowdown.
* Local development environments and easy to use single-script demos
  for Kubernetes-related projects. This use case may be covered to an
  extent using `minikube` yet the previous points about an ability to
  run such scripts on CI and resource usage are also valid for this
  use case. DIND solutions provides the possibility of a more
  lightweight demo for Docker users, see
  [Virtlet demo](https://github.com/Mirantis/virtlet#virtlet-usage-demo)
  as an example.
* Support for other platforms besides Linux. You can run `minikube` on
  your Mac (at the cost of having to run a VM for it) but you can't
  run `hack/local-up-cluster.sh` there unless you make an effort to
  run it in a container.
* Support for remote docker engines. If you prefer lightweight laptop
  for your work it would be nice to be able to utilize some powerful
  remote box or a cloud instance for your Kubernetes work. You can
  [build Kubernetes](https://github.com/kubernetes/kubernetes/tree/master/build#really-remote-docker-engine)
  this way, but what about test clusters? Moreover, if you have a
  habit of working from places with limited network speed, while
  current k8s build container rsync functionality will save you some
  traffic, you still have to pull back bulky binaries like
  `hyperkube`.

These use cases are actually already covered by
[kubeadm-dind-cluster](https://github.com/Mirantis/kubeadm-dind-cluster). The
plan is to adapt this project for inclusion in the `kubeadm` repository.

## Proposal

The `kubeadm-local-cluster` tool consists of a docker image that
contains `systemd`, `docker` and necessary tools for the cluster
startup. There are several image versions for supported versions of
Kubernetes that include corresponding version of `hyperkube` and
`kubeadm` binaries. Also there's a 'bare' image that doesn't include
`hyperkube` and `kubeadm` binaries and can only be used for building
Kubernetes from source.

The cluster itself consists of several docker containers used as
Kubernetes nodes. The containers are created with `systemd` and
`docker` running inside them, then `kubeadm` is used to install
Kubernetes onto them.

![kubeadm-local-cluster architecture](local-cluster.png)

## Usage

`local-cluster` defaults to start the most recent stable Kubernetes
version based on binaries downloaded from
`https://storage.googleapis.com/kubernetes-release/release/...` It's
behavior can be customized using configuration variables specified in
`~/.kubeadm-local-cluster/config.yaml` or in another configuration
file passed via `LOCAL_CLUSTER_CONF` environment variable. For each
configuration variable there's corresponding command line flag and
environment variable that can be passed to `./local-cluster up` and
`./local-cluster reup` commands.

Below is an example shell session
```
$ wget https://...../local-cluster
$ chmod +x local-cluster

$ # start the cluster
$ ./local-cluster

$ # make it easier to access the cluster
$ export KUBECONFIG=$HOME/.kubeadm-local-cluster/admin.conf
$ alias kubectl=$HOME/.kubeadm-local-cluster/kubectl

$ kubectl get nodes
NAME          STATUS    AGE       VERSION
kube-master   Ready     2m        v1.6.1
kube-node-1   Ready     2m        v1.6.1
kube-node-2   Ready     2m        v1.6.1

$ kubectl get pods --all-namespaces
NAMESPACE     NAME                                    READY     STATUS    RESTARTS   AGE
kube-system   etcd-kube-master                        1/1       Running   1          1m
kube-system   kube-apiserver-kube-master              1/1       Running   1          2m
kube-system   kube-controller-manager-kube-master     1/1       Running   1          2m
kube-system   kube-dns-3946503078-6mh7l               3/3       Running   0          1m
kube-system   kube-proxy-b3rs2                        1/1       Running   1          2m
kube-system   kube-proxy-pqb48                        1/1       Running   1          2m
kube-system   kube-proxy-zdt8f                        1/1       Running   1          2m
kube-system   kube-scheduler-kube-master              1/1       Running   1          2m
kube-system   kubernetes-dashboard-2396447444-8wz7p   1/1       Running   0          1m

$ kubectl proxy &
$ # k8s dashboard available at http://localhost:8001/ui

$ # restart the cluster, this should happen much quicker than initial startup
$ ./local-cluster up

$ # stop the cluster
$ ./local-cluster down

$ # remove DIND containers, network and volumes
$ ./local-cluster clean

```

After the cluster is set up using `kubeadm`, it's state is snapshotted
to make cluster restarts several times faster. The underlying
mechanism will be described later in this document. Only single
cluster snapshot per docker engine is supported at the moment.

The `local-cluster` command supports the following subcommands:

```shell
./local-cluster up [flags...]
```
Brings up a local cluster. If cluster config didn't change since the
last cluster startup, reuse the previous snapshot.  If a snapshot
didn't exist, a new one is created for current cluster config. If
there's a cluster snapshot which was made for different cluster
config, the snapshot is deleted and a new one is made.
After this command the user can access the cluster conveniently
by issuing the following additional commands:
```shell
export KUBECONFIG=$HOME/.kubeadm-local-cluster/admin.conf
alias kubectl=$HOME/.kubeadm-local-cluster/kubectl
```
For this command to work with `BUILD_*` options, it must be invoked
from Kubernetes source directory.

```shell
./local-cluster down
```
If the cluster is up, brings it down, deleting all the node
containers, but keeping the current cluster snapshot. This command
does nothing if the cluster is not running.

```shell
./local-cluster clean
```
Same as `./local-cluster down`, but also deletes the current
cluster snapshot if it exists, along with `k8s-kubeadm-local-cluster` docker
network.

```shell
./local-cluster e2e
./local-cluster e2e "test name substring"
```
When invoked without extra arguments, this runs all the non-serial e2e
conformance tests against the current cluster. If an argument is
given, the command will run e2e tests that have the specified
substring in their name. When using this command, `GINKGO_PARALLEL` is
set to a non-empty string so the tests run in parallel.

```shell
./local-cluster e2e-serial
./local-cluster e2e-serial "test name substring"
```
When invoked without extra arguments, this runs all the serial e2e
conformance tests against the current cluster. If an argument is
given, the command will run e2e tests that have the specified
substring in their name. When using this command the tests are not run
in parallel.

```shell
./local-cluster reup [flags...]
```
Same as `./local-cluster up`, but also rebuilds hyperkube & kubectl
binaries and updates them in case if Kubernetes is built from source.
This command can still reuse the last snapshot.

## Configuration variables

The variables are listed with their default values and corresponding
command line flags. They can be specified in
`~/.kubeadm-local-cluster/config.yaml` or a file under the path passed
via `LOCAL_CLUSTER_CONF`, or using environment variables. The
environment variables which are names are based on the names of
command line options, but uppercased, with leading dashes replaced by
`K8S_DIND_` and with remaining dashes replaced with underscore (`_`),
e.g. `K8S_DIND_BUILD_KUBEADM` or `K8S_DIND_NO_DEDICATED`.


| Configuration setting | Command line option | Type | Default | Description |
| ----------------------|---------------------|------|---------|-------------|
| dind.subnet | `--dind-subnet` | `CIDR` | `10.192.0.0/16` | an IP range for use by node containers |
| verbose | `--v` | `int` | 1 | verbosity level for `kubeadm-local-cluster`. Set to at least `2` to see `kubeadm` output |
| source.kubernetesVersion | `--kubernetes-version` | `string` | latest stable | the version of k8s to use |
| source.buildKubeadm | `--build-kubeadm` | `bool` | `false` | build `kubeadm` locally from Kubernetes sources |
| source.buildHyperkube | `--build-hyperkube` | `bool` | `false` | build `hyperkube` locally from Kubernetes sources |
| source.buildKubectl | `--build-kubectl` | `bool` | `false` | build `kubectl` locally from Kubernetes sources |
| source.k8sDir | `--k8s-dir` | current directory| Kubernetes source directory (only when at least one of `source.build*` options is enabled) |
| source.kubeadmUrl | `--kubeadm-url` | `URL` | | an URL to download `kubeadm` binary from |
| source.kubeadmSha1 | `--kubeadm-sha1` | `SHA1` | | sha1sum of the kubeadm binary downloaded from `kubeadmUrl` |
| source.hyperkubeUrl | `--hyperkube-url` | `URL` | | an URL to download `hyperkube` binary from |
| source.hyperkubeSha1 | `--hyperkube-sha1` | `SHA1` | | sha1sum of the hyperkube binary downloaded from `hyperkubeUrl` |
| source.kubectlUrl | `--kubectl-url` | `URL` | | an URL to download `kubectl` binary from |
| source.kubectlSha1 | `--kubectl-sha1` | `SHA1` | | sha1sum of the kubectl binary downloaded from `kubectlUrl` |
| cluster.numNodes | `--num-nodes` | `int` | `2` | the number of worker nodes to create. If set to `0`, only master node is set up and `dedicated` taint is removed from it. |
| cluster.cniPlugin | `--cni-plugin` | `weave`, `flannel`, `calico`, `bridge` | `bridge` | CNI plugin to use. When `bridge` plugin is used, `kubeadm-local-cluster` joins the bridges from the nodes together with another bridge |
| cluster.noDedicated | `--no-dedicated` | `bool` | `false` | remove `dedicated` taint from master node even if `NUM_NODES` is greater than zero |
| cluster.apiServerPort | `--apiserver-port` | `bool` | `false` | apiserver port to use. It'll be exposed on the Docker host |
| cluster.addons | `--adons` | comma-separated list | `dashboard` | specifies a list of non-critical cluster addons to install (see below) |

The `cluster.addons` configuration variable is a comma-separated lists
of optional cluster addons to install. Each item is either a path to a
`.yaml` file containing addon object definition(s), an URL or a name
of predefined addon. Currently the only predefined addon supported is
`dashboard` which is the same as specifying the URL of current
Kubernetes dashboard installation yaml file.

Below is an example of configuration file.

```yaml
verbose: 2
source:
  # build kubeadm from source
  buildKubeadm: true
  # download hyperkube
  hyperkubeUrl: "https://storage.googleapis.com/kubernetes-release/release/v1.6.1/bin/linux/amd64/hyperkube"
  hyperkubeSha1: 4e375e27a5aa10a8d1a0d10da5c3c2f2d0ee4770
cluster:
  # set number of nodes to 6 for faster e2e
  numNodes: 6
  # use Weave Net
  cniPlugin: weave
  # use some addons besides the default (dashboard)
  addons: dashboard,https://example.com/my-favorite-addon.yaml,/path/to/local/addon.yaml
```

As an additional configuration example, `gce.sh` script bundled with
`kubeadm-local-cluster` can be used to set up a cluster on a GCE
instance that's prepared using `docker-machine`.

## Implementation notes

### Base image

The `kubeadm-local-cluster` tool by default uses a prebuilt image to
start the cluster.  There's an image per each supported Kubernetes
version and supported architecture (`amd64`, `arm`, `arm64`, `ppc64le`
and `s390x`). There's also an image for each architecture that doesn't
contain prebuilt Kubernetes binaries.  The images are based on
`gcr.io/google-containers/debian-base-$(ARCH):0.1`

It's possible to customize the behavior of the tool by creating a
config file named `~/.kubeadm-local-cluster/config.yaml`.  It's
possible to override default config path via `LOCAL_CLUSTER_CONF`
environment variable.

When `source.build*` options are enabled, `local-cluster`
invokes appropriate `build/run.sh make ...` commands in Kubernetes
source tree, preventing the build container from copying back the
binaries using `KUBE_RUN_COPY_OUTPUT=n`. The control container then
pulls the binaries directly from build container.

### Networking

The `kubeadm-local-cluster` tool utilizes named docker network
`k8s-kubeadm-local-cluster`. This network is always created or replaced upon
the cluster startup using address range specified via `DIND_SUBNET`
configuration variable. `kubeadm-local-cluster` makes docker's internal DNS
server available to all the nodes in the cluster, so nodes can be
referred by their hostnames: `kube-master`, `kube-node-1`,
`kube-node-2` and so on.

### Cluster snapshot mechanism

In order to reduce the amount of time needed to restart the cluster,
the `kubeadm-local-cluster` tool takes the snapshot of the cluster
after setting it up.  The snapshot is made by means of preserving each
node's `/var/lib/docker` directory as a docker volume and saving each
node container's filesystem in an image. The volumes and images
created by snapshotting mechanism are removed during
`./local-cluster clean`.

### Future directions

The following features may be added later:

* automatic PV creation like it's done in `minikube`, so most Helm
  charts from
  [kubernetes/charts](https://github.com/kubernetes/charts) repository
  work out of the box
* enabling feature gates for Kubernetes components
* automatic installation of `helm` and `tiller`
* support using `delve` to debug k8s components.  Should be an option
  for building k8s with `GOGCFLAGS="-N -l"` plus `delve` binary in the
  images. Possible problem: as of Go 1.7, trying to build e.g.
  `kubectl` with `GOGCFLAGS="-N -l"` may result in stack overflow
  in Go compiler.

### CI

A test matrix needs to be set up so each combination of supported
Kubernetes version and supported CNI plugin is tested. In addition to
this, there should be a test that builds Kubernetes binaries based on
`master` branch of
[kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
repository. k8s `master` test should be executed daily so possible
issues that break DIND can be tracked down in timely manner and with
minimal use of `git bisect`.

### Limitations

There are problems with `btrfs` support because of a
[docker bug](https://github.com/docker/docker/issues/9939) and
possible problems with running `kubelet` on `btrfs`. To be
investigated further.

## Differences from kubeadm-dind-cluster

Listed below are the differences between `kubeadm-local-cluster` and
[kubeadm-dind-cluster](https://github.com/Mirantis/kubeadm-dind-cluster)
project.

* `kubeadm-local-cluster` is written in Go, while
  `kubeadm-dind-cluster` consists mostly of shell scripts
* `kubeadm-dind-cluster` uses Ubuntu images while `kubeadm-local-cluster` uses
  smaller `debian` images
* `kubeadm-local-cluster` supports more platforms, e.g. ARM
* `kubeadm-dind-cluster` doesn't support command line flags
* `kubeadm-dind-cluster` uses a combination of `docker diff` and `tar` for snapshots
  instead of images
* `kubeadm-local-cluster` by default hides `kubeadm` output
* `kubeadm-local-cluster` provides `BUILD_KUBECTL` option (`kubectl` binary
  management in case of source build is incomplete in `kubeadm-dind-cluster`)
* `kubeadm-local-cluster` makes it possible to specify URLs for
  `hyperkube`/`kubeadm`/`kubectl` binaries as part of its configuration
  (in `kubeadm-dind-cluster` the URLs are saved in the image, there's
  possibility to override them but it's not well tested)
* `kubeadm-local-cluster` automatically detects k8s version mismatch in snapshots
* `kubeadm-local-cluster` can start single-node clusters and supports
  `NO_DEDICATED` option
* `kubeadm-dind-cluster` uses `docker save` to save the images used by
  the cluster inside the main image, so they don't have to be pulled
  afterwards. It needs to be checked how much it really helps before
  including this functionality in `kubeadm-local-cluster`.

## Prior work

* `kubeadm-dind-cluster` was initially derived from
  [kubernetes-dind-cluster](https://github.com/sttts/kubernetes-dind-cluster),
  although as of now the code was completely rewritten.
  kubernetes-dind-cluster is somewhat faster but uses less standard
  way of k8s deployment. It also doesn't include support for consuming
  binaries from remote dockerized builds.
* [kubeadm-ci-dind](https://github.com/errordeveloper/kubeadm-ci-dind),
  [kubeadm-ci-packager](https://github.com/errordeveloper/kubeadm-ci-packager) and
  [kubeadm-ci-tester](https://github.com/errordeveloper/kubeadm-ci-tester).
  These projects are similar to kubeadm-dind-cluster but are intended primarily for CI.
  They include packaging step which is too slow for the purpose of having
  convenient k8s "playground". kubeadm-dind-cluster uses Docker images
  from `kubeadm-ci-dind`.
* [nkube](https://github.com/marun/nkube) starts
  Kubernetes-in-Kubernetes clusters.
