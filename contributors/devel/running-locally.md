Getting started locally
-----------------------

**Table of Contents**

- [Requirements](#requirements)
    - [Linux](#linux)
    - [Container Runtime](#container-runtime)
    - [etcd](#etcd)
    - [go](#go)
    - [OpenSSL](#openssl)
    - [CFSSL](#cfssl)
- [Clone the repository](#clone-the-repository)
- [Starting the cluster](#starting-the-cluster)
- [Running a container](#running-a-container)
- [Running a user defined pod](#running-a-user-defined-pod)
- [Troubleshooting](#troubleshooting)
    - [I cannot reach service IPs on the network.](#i-cannot-reach-service-ips-on-the-network)
    - [I cannot create a replication controller with replica size greater than 1!  What gives?](#i-cannot-create-a-replication-controller-with-replica-size-greater-than-1--what-gives)
    - [I changed Kubernetes code, how do I run it?](#i-changed-kubernetes-code-how-do-i-run-it)
    - [kubectl claims to start a container but `get pods` and `docker ps` don't show it.](#kubectl-claims-to-start-a-container-but-get-pods-and-docker-ps-dont-show-it)
    - [The pods fail to connect to the services by host names](#the-pods-fail-to-connect-to-the-services-by-host-names)

## Requirements

### Linux

Not running Linux? Consider running [Minikube](https://kubernetes.io/docs/tasks/tools/#minikube), or on a cloud provider like [Google Compute Engine](https://cloud.google.com/compute/).

### Container Runtime

You will need a [Container Runtime](https://kubernetes.io/docs/setup/production-environment/container-runtimes/) like [Containerd](https://github.com/containerd/containerd) or [CRI-O](https://github.com/cri-o/cri-o) installed and running.

### etcd

You need [etcd](https://github.com/coreos/etcd/releases) installed and in your `$PATH`. [Check here](https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#etcd) for instructions on installing a local copy.

### go

You need [go](https://golang.org/doc/install) in your path (see [Development Guide](development.md#go) for supported versions), please make sure it is installed and in your ``$PATH``.

### OpenSSL

You need [OpenSSL](https://www.openssl.org/) installed.  If you do not have the `openssl` command available, the script will print an appropriate error.

### CFSSL

The [CFSSL](https://cfssl.org/) binaries (cfssl, cfssljson) must be installed and available on your ``$PATH``.

The easiest way to get it is to run these shell commands:

```sh
go install github.com/cloudflare/cfssl/cmd/...@latest
PATH=$PATH:$GOPATH/bin
```

## Clone the repository

In order to run kubernetes you must have the kubernetes code on the local machine. Cloning this repository is sufficient.

```sh
git clone --filter=blob:none https://github.com/kubernetes/kubernetes.git
```

The `--filter=blob:none` parameter is optional and will ensure a smaller download.

## Starting the cluster

Set the endpoint for container runtime e.g. containerd
```sh
export CONTAINER_RUNTIME_ENDPOINT="unix:///run/containerd/containerd.sock"
```

Optionally set any other environment variables as needed to change the cluster configuration. The possible options are listed at the top of the `./hack/local-up-cluster.sh` script.

In a separate tab of your terminal, run the following:

```sh
cd kubernetes
./hack/local-up-cluster.sh
```

Since root access is sometimes needed to start/stop Kubernetes daemons, `./hack/local-up-cluster.sh` may need to be run as root. If it reports failures, try this instead:

```sh
sudo -E PATH=$PATH ./hack/local-up-cluster.sh
```

This will build and start a lightweight local cluster, consisting of a master and a single node. Press Control+C to shut it down.

**Note:** If you've already compiled the Kubernetes components, you can avoid rebuilding them with the `-O` flag.

```sh
./hack/local-up-cluster.sh -O
```

You can use the `./cluster/kubectl.sh` script to interact with the local cluster. `./hack/local-up-cluster.sh` will
print the commands to run to point kubectl at the local cluster.


## Running a container

Your cluster is running, and you want to start running containers!

You can now use any of the cluster/kubectl.sh commands to interact with your local setup after setting your KUBECONFIG

```sh
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
```

```sh
./cluster/kubectl.sh get pods
./cluster/kubectl.sh get services
./cluster/kubectl.sh get replicationcontrollers
./cluster/kubectl.sh run my-nginx --image=nginx --port=80
```

While waiting for the provisioning to complete, you can monitor progress in another terminal with these commands.

```sh
# containerd
# To list images
ctr --namespace k8s.io image ls
# To list containers
ctr --namespace k8s.io containers ls
```

Once provisioning is complete, you can use the following commands for Kubernetes introspection.

```sh
./cluster/kubectl.sh get pods
./cluster/kubectl.sh get services
./cluster/kubectl.sh get replicationcontrollers
```

## Running a user defined pod

Note the difference between a [container](https://kubernetes.io/docs/concepts/containers/)
and a [pod](https://kubernetes.io/docs/concepts/workloads/pods/). Since you only asked for the former, Kubernetes will create a wrapper pod for you.
However, you cannot view the nginx start page on localhost. To verify that nginx is running, you need to run `curl` within the Docker container (try `docker exec`).

You can control the specifications of a pod via a user defined manifest, and reach nginx through your browser on the port specified therein:

```sh
./cluster/kubectl.sh create -f test/fixtures/doc-yaml/user-guide/pod.yaml
```

Congratulations!

## Troubleshooting

### I cannot reach service IPs on the network.

Some firewall software that uses iptables may not interact well with
kubernetes.  If you have trouble around networking, try disabling any
firewall or other iptables-using systems, first.  Also, you can check
if SELinux is blocking anything by running a command such as `journalctl --since yesterday | grep avc`.

By default the IP range for service cluster IPs is 10.0.*.* - depending on your
docker installation, this may conflict with IPs for containers.  If you find
containers running with IPs in this range, edit hack/local-cluster-up.sh and
change the service-cluster-ip-range flag to something else.

### I cannot create a replication controller with replica size greater than 1!  What gives?

You are running a single node setup.  This has the limitation of only supporting a single replica of a given pod.  If you are interested in running with larger replica sizes, we encourage you to try Kind or one of the cloud providers.

### I changed Kubernetes code, how do I run it?

```sh
cd kubernetes
make
./hack/local-up-cluster.sh
```

### kubectl claims to start a container but `get pods` and `docker ps` don't show it.

One or more of the Kubernetes daemons might've crashed. Tail the logs of each in /tmp.

### The pods fail to connect to the services by host names

To start the DNS service, you need to set the following variables:

```sh
KUBE_ENABLE_CLUSTER_DNS=true
KUBE_DNS_SERVER_IP="10.0.0.10"
KUBE_DNS_NAME="cluster.local"
```

To know more on DNS service you can check out the [docs](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/) and see [Debugging DNS Resolution](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) for diagnosing DNS problems.

### All pod fail to start with a cgroups error of `expected cgroupsPath to be of format "slice:prefix:name" for systemd cgroups`

Your container runtime is using the systemd cgroup driver, but by default `./hack/local-up-cluster.sh` uses cgroupfs.  To correct the mismatch, set the `CGROUP_DRIVER` environment variable to systemd as well.

```sh
export CGROUP_DRIVER=systemd
```
