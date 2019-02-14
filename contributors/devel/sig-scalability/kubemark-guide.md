# Kubemark User Guide

## Introduction

Kubemark is a performance testing tool which allows users to run experiments on
simulated clusters. The primary use case is scalability testing, as simulated
clusters can be much bigger than the real ones. The objective is to expose
problems with the master components (API server, controller manager or
scheduler) that appear only on bigger clusters (e.g. small memory leaks).

This document serves as a primer to understand what Kubemark is, what it is not,
and how to use it.

## Architecture

On a very high level, Kubemark cluster consists of two parts: a real master
and a set of “Hollow” Nodes. The prefix “Hollow” to any component means an
implementation/instantiation of the actual component with all “moving”
parts mocked out. The best example is HollowKubelet, which pretends to be an
ordinary Kubelet, but does not start anything, nor mount any volumes - it just
lies it does. More detailed design and implementation details are at the end
of this document.

Currently, master components run on a dedicated machine as pods that are
created/managed by kubelet, which itself runs as either a systemd or a supervisord
service on the master VM depending on the VM distro (though currently it is
only systemd as we use a GCI image). Having a dedicated machine for the master
has a slight advantage over running the master components on an external cluster,
which is being able to completely isolate master resources from everything else.
The HollowNodes on the other hand are run on an ‘external’ Kubernetes cluster
as pods in an isolated namespace (named kubemark). This idea of using pods on a
real cluster behave (or act) as nodes on the kubemark cluster lies at the heart of
kubemark's design.

## Requirements

To run Kubemark you need a Kubernetes cluster (called `external cluster`)
for running all your HollowNodes and a dedicated machine for a master.
Master machine has to be directly routable from HollowNodes. You also need
access to a Docker repository (which is gcr.io in the case of GCE) that has the
container images for etcd, hollow-node and node-problem-detector.

Currently, scripts are written to be easily usable by GCE, but it should be
relatively straightforward to port them to different providers or bare metal.
There is an ongoing effort to refactor kubemark code into provider-specific (gce)
and provider-independent code, which should make it relatively simple to run
kubemark clusters on other cloud providers as well.

## Common use cases and helper scripts

Common workflow for Kubemark is:
- starting a Kubemark cluster (on GCE)
- running e2e tests on Kubemark cluster
- monitoring test execution and debugging problems
- turning down Kubemark cluster

(For now) Included in descriptions there will be comments helpful for anyone who’ll
want to port Kubemark to different providers.
(Later) When the refactoring mentioned in the above section finishes, we would replace
these comments with a clean API that would allow kubemark to run on top of any provider.

### Starting a Kubemark cluster

To start a Kubemark cluster on GCE you need to create an external kubernetes
cluster (it can be GCE, GKE or anything else) by yourself, make sure that kubeconfig
points to it by default, build a kubernetes release (e.g. by running
`make quick-release`) and run `test/kubemark/start-kubemark.sh` script.
This script will create a VM for master (along with mounted PD and firewall rules set),
then start kubelet and run the pods for the master components. Following this, it
sets up the HollowNodes as Pods on the external cluster and do all the setup necessary
to let them talk to the kubemark apiserver. It will use the configuration stored in
`cluster/kubemark/config-default.sh` - you can tweak it however you want, but note that
some features may not be implemented yet, as implementation of Hollow components/mocks
will probably be lagging behind ‘real’ one. For performance tests interesting variables
are `NUM_NODES` and `KUBEMARK_MASTER_SIZE`. After start-kubemark script is finished,
you’ll have a ready Kubemark cluster, and a kubeconfig file for talking to the Kubemark
cluster is stored in `test/kubemark/resources/kubeconfig.kubemark`.

Currently we're running HollowNode with a limit of 0.09 CPU core/pod and 220MB of memory.
However, if we also take into account the resources absorbed by default cluster addons
and fluentD running on the 'external' cluster, this limit becomes ~0.1 CPU core/pod,
thus allowing ~10 HollowNodes to run per core (on an "n1-standard-8" VM node).

#### Behind the scene details:

start-kubemark.sh script does quite a lot of things:

- Prepare a master machine named MASTER_NAME (this variable's value should be set by this point):
  (*the steps below use gcloud, and should be easy to do outside of GCE*)
  1. Creates a Persistent Disk for use by the master (one more for etcd-events, if flagged)
  2. Creates a static IP address for the master in the cluster and assign it to variable MASTER_IP
  3. Creates a VM instance for the master, configured with the PD and IP created above.
  4. Set firewall rule in the master to open port 443\* for all TCP traffic by default.

<sub>\* Port 443 is a secured port on the master machine which is used for all
external communication with the API server. In the last sentence *external*
means all traffic coming from other machines, including all the Nodes, not only
from outside of the cluster. Currently local components, i.e. ControllerManager
and Scheduler talk with API server using insecure port 8080.</sub>

- [Optional to read] Establish necessary certs/keys required for setting up the PKI for kubemark cluster:
  (*the steps below are independent of GCE and work for all providers*)
  1. Generate a randomly named temporary directory for storing PKI certs/keys which is delete-trapped on EXIT.
  2. Create a bearer token for 'admin' in master.
  3. Generate certificate for CA and (certificate + private-key) pair for each of master, kubelet and kubecfg.
  4. Generate kubelet and kubeproxy tokens for master.
  5. Write a kubeconfig locally to `test/kubemark/resources/kubeconfig.kubemark` for enabling local kubectl use.

- Set up environment and start master components (through `start-kubemark-master.sh` script):
  (*the steps below use gcloud for SSH and SCP to master, and should be easy to do outside of GCE*)
  1. SSH to the master machine and create a new directory (`/etc/srv/kubernetes`) and write all the
     certs/keys/tokens/passwords to it.
  2. SCP all the master pod manifests, shell scripts (`start-kubemark-master.sh`, `configure-kubectl.sh`, etc),
     config files for passing env variables (`kubemark-master-env.sh`) from the local machine to the master.
  3. SSH to the master machine and run the startup script `start-kubemark-master.sh` (and possibly others).

  Note: The directory structure and the functions performed by the startup script(s) can vary based on master distro.
        We currently support the GCI image `gci-dev-56-8977-0-0` in GCE.

- Set up and start HollowNodes (as pods) on the external cluster:
  (*the steps below (except 2nd and 3rd) are independent of GCE and work for all providers*)
  1. Identify the right kubemark binary from the current kubernetes repo for the platform linux/amd64.
  2. Create a Docker image for HollowNode using this binary and upload it to a remote Docker repository.
     (We use gcr.io/ as our remote docker repository in GCE, should be different for other providers)
  3. [One-off] Create and upload a Docker image for NodeProblemDetector (see kubernetes/node-problem-detector repo),
     which is one of the containers in the HollowNode pod, besides HollowKubelet and HollowProxy. However we
     use it with a hollow config that essentially has an empty set of rules and conditions to be detected.
     This step is required only for other cloud providers, as the docker image for GCE already exists on GCR.
  4. Create secret which stores kubeconfig for use by HollowKubelet/HollowProxy, addons, and configMaps
     for the HollowNode and the HollowNodeProblemDetector.
  5. Create a ReplicationController for HollowNodes that starts them up, after replacing all variables in
     the hollow-node_template.json resource.
  6. Wait until all HollowNodes are in the Running phase.

### Running e2e tests on Kubemark cluster

To run standard e2e test on your Kubemark cluster created in the previous step
you execute `test/kubemark/run-e2e-tests.sh` script. It will configure ginkgo to
use Kubemark cluster instead of something else and start an e2e test. This
script should not need any changes to work on other cloud providers.

By default (if nothing will be passed to it) the script will run a Density '30
test. If you want to run a different e2e test you just need to provide flags you want to be
passed to `hack/ginkgo-e2e.sh` script, e.g. `--ginkgo.focus="Load"` to run the
Load test.

By default, at the end of each test, it will delete namespaces and everything
under it (e.g. events, replication controllers) on Kubemark master, which takes
a lot of time. Such work aren't needed in most cases: if you delete your
Kubemark cluster after running `run-e2e-tests.sh`; you don't care about
namespace deletion performance, specifically related to etcd; etc. There is a
flag that enables you to avoid namespace deletion: `--delete-namespace=false`.
Adding the flag should let you see in logs: `Found DeleteNamespace=false,
skipping namespace deletion!`

### Monitoring test execution and debugging problems

Run-e2e-tests prints the same output on Kubemark as on ordinary e2e cluster, but
if you need to dig deeper you need to learn how to debug HollowNodes and how
Master machine (currently) differs from the ordinary one.

If you need to debug master machine you can do similar things as you do on your
ordinary master. The difference between Kubemark setup and ordinary setup is
that in Kubemark etcd is run as a plain docker container, and all master
components are run as normal processes. There's no Kubelet overseeing them. Logs
are stored in exactly the same place, i.e. `/var/logs/` directory. Because
binaries are not supervised by anything they won't be restarted in the case of a
crash.

To help you with debugging from inside the cluster startup script puts a
`~/configure-kubectl.sh` script on the master. It downloads `gcloud` and
`kubectl` tool and configures kubectl to work on unsecured master port (useful
if there are problems with security). After the script is run you can use
kubectl command from the master machine to play with the cluster.

Debugging HollowNodes is a bit more tricky, as if you experience a problem on
one of them you need to learn which hollow-node pod corresponds to a given
HollowNode known by the Master. During self-registeration HollowNodes provide
their cluster IPs as Names, which means that if you need to find a HollowNode
named `10.2.4.5` you just need to find a Pod in external cluster with this
cluster IP. There's a helper script
`test/kubemark/get-real-pod-for-hollow-node.sh` that does this for you.

When you have a Pod name you can use `kubectl logs` on external cluster to get
logs, or use a `kubectl describe pod` call to find an external Node on which
this particular HollowNode is running so you can ssh to it.

E.g. you want to see the logs of HollowKubelet on which pod `my-pod` is running.
To do so you can execute:

```
$ kubectl kubernetes/test/kubemark/resources/kubeconfig.kubemark describe pod my-pod
```

Which outputs pod description and among it a line:

```
Node:				1.2.3.4/1.2.3.4
```

To learn the `hollow-node` pod corresponding to node `1.2.3.4` you use
aforementioned script:

```
$ kubernetes/test/kubemark/get-real-pod-for-hollow-node.sh 1.2.3.4
```

which will output the line:

```
hollow-node-1234
```

Now you just use ordinary kubectl command to get the logs:

```
kubectl --namespace=kubemark logs hollow-node-1234
```

All those things should work exactly the same on all cloud providers.

### Turning down Kubemark cluster

On GCE you just need to execute `test/kubemark/stop-kubemark.sh` script, which
will delete HollowNode ReplicationController and all the resources for you. On
other providers you’ll need to delete all this stuff by yourself. As part of
the effort mentioned above to refactor kubemark into provider-independent and
provider-specific parts, the resource deletion logic specific to the provider
would move out into a clean API.

## Some current implementation details and future roadmap

Kubemark master uses exactly the same binaries as ordinary Kubernetes does. This
means that it will never be out of date. On the other hand HollowNodes use
existing fake for Kubelet (called SimpleKubelet), which mocks its runtime
manager with `pkg/kubelet/dockertools/fake_manager.go`, where most logic sits.
Because there's no easy way of mocking other managers (e.g. VolumeManager), they
are not supported in Kubemark (e.g. we can't schedule Pods with volumes in them
yet).

We currently plan to extend kubemark along the following directions:
- As you would have noticed at places above, we aim to make kubemark more structured
  and easy to run across various providers without having to tweak the setup scripts,
  using a well-defined kubemark-provider API.
- Allow kubemark to run on various distros (GCI, debian, redhat, etc) for any
  given provider.
- Make Kubemark performance on ci-tests mimic real cluster ci-tests on metrics such as
  CPU, memory and network bandwidth usage and realizing this goal through measurable
  objectives (like the kubemark metric should vary no more than X% with the real
  cluster metric). We could also use metrics reported by Prometheus.
- Improve logging of CI-test metrics (such as aggregated API call latencies, scheduling
  call latencies, %ile for CPU/mem usage of different master components in density/load
  tests) by packing them into well-structured artifacts instead of the (current) dumping
  to logs.
- Create a Dashboard that lets easy viewing and comparison of these metrics across tests.

