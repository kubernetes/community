# Node End-To-End (e2e) tests

Node e2e tests are component tests meant for testing the Kubelet code on a custom host environment. Tests can be run either locally or against a host running on Google Compute Engine (GCE). Node e2e tests are run as both pre- and post- submit tests by the Kubernetes project.

*Note: Linux only. Mac and Windows unsupported.*

*Note: There is no scheduler running. The e2e tests have to do manual scheduling, e.g. by using `framework.PodClient`.*

# Running tests
To ensure node e2e test build can locate the Kubernetes source code, please take one of the following actions:
- *Set the `KUBE_ROOT` environment variable* to the absolute path of the directory where the Kubernetes repository has been cloned.
- *Clone the Kubernetes repository* into the default location: `~/k8s.io/kubernetes`.

## Locally

Why run tests *locally*? It is much faster than running tests remotely.

Prerequisites:
- [Install etcd](https://github.com/coreos/etcd/releases) and include the path to the installation in your PATH
  - Verify etcd is installed correctly by running `which etcd`
  - Or make etcd binary available and executable at `/tmp/etcd`
- [containerd](https://github.com/containerd/containerd) configured with the cgroupfs driver
- Working CNI
  - Ensure that you have a valid CNI configuration in /etc/cni/net.d/. For testing purposes, a [bridge](https://www.cni.dev/plugins/current/main/bridge/) configuration should work.

From the Kubernetes base directory, run:

```sh
make test-e2e-node
```

This will run the *ginkgo* binary against the subdirectory *test/e2e_node*, which will in turn:
- Ask for sudo access (needed for running some of the processes)
- Build the Kubernetes source code
- Pre-pull docker images used by the tests
- Start a local instance of *etcd*
- Start a local instance of *kube-apiserver*
- Start a local instance of *kubelet*
- Run the test using the locally started processes
- Output the test results to STDOUT
- Stop *kubelet*, *kube-apiserver*, and *etcd*

To view the settings and print help, run:

```sh
make test-e2e-node PRINT_HELP=y
```

## Remotely

Why run tests *remotely*? Tests will be run in a customized testing environment. This environment closely mimics the pre- and post- submit testing performed by the project.

Prerequisites:
- [Join the googlegroup](https://groups.google.com/a/kubernetes.io/group/dev) `dev@kubernetes.io`
  - *This provides read access to the node test images.*
- Setup a [Google Cloud Platform](https://cloud.google.com/) account and project with Google Compute Engine enabled
- Install and setup the [gcloud sdk](https://cloud.google.com/sdk/downloads)
  - Set your project and a zone by running `gcloud config set project $PROJECT` and `gcloud config set compute/zone $zone`
  - Verify the sdk is setup correctly by running `gcloud compute instances list` and `gcloud compute images list --project cos-cloud`
  - Configure credentials for the same project that you configured above for "application defaults". This can be done with `gcloud auth application-default login` or by setting the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to a path to a credentials file.

Run:

```sh
make test-e2e-node REMOTE=true
```

This will:
- Build the Kubernetes source code
- Create a new GCE instance using the default test image
  - Instance will be named something like **test-cos-beta-81-12871-44-0**
- Lookup the instance public ip address
- Copy a compressed archive file to the host containing the following binaries:
  - ginkgo
  - kubelet
  - kube-apiserver
  - e2e_node.test (this binary contains the actual tests to be run)
- Unzip the archive to a directory under **/tmp/gcloud**
- Run the tests using the `ginkgo` command
  - Starts etcd, kube-apiserver, kubelet
  - The ginkgo command is used because this supports more features than running the test binary directly
- Output the remote test results to STDOUT
- `scp` the log files back to the local host under /tmp/_artifacts/e2e-node-containervm-v20160321-image
- Stop the processes on the remote host
- **Leave the GCE instance running**

**Note: Subsequent tests run using the same image will *reuse the existing host* instead of deleting it and
provisioning a new one. To delete the GCE instance after each test see
*[DELETE_INSTANCE](#delete-instance-after-tests-run)*.**


## Additional Remote Options

## Run tests using different images

This is useful if you want to run tests against a host using a different OS distro or container runtime than
provided by the default image.

List the available test images using gcloud.

```sh
gcloud compute images list --project="cos-cloud" --no-standard-images --filter="name ~ 'cos-beta.*'"
```

This will output a list of the available images for the default image project.

Then run:

```sh
make test-e2e-node REMOTE=true IMAGES="<comma-separated-list-images>"
```

## Run tests against a running GCE instance (not an image)

This is useful if you have an host instance running already and want to run the tests there instead of on a new instance.

```sh
make test-e2e-node REMOTE=true HOSTS="<comma-separated-list-of-hostnames>"
```

## Run tests against a different network and subnet (not the default)

This is useful if you want to run tests on a non-default network and subnet.

```sh
make test-e2e-node REMOTE=true NETWORK="<network> SUBNET="<subnet>"
```

## Delete instance after tests run

This is useful if you want recreate the instance for each test run to trigger flakes related to starting the instance.

```sh
make test-e2e-node REMOTE=true DELETE_INSTANCES=true
```

## Keep instance, test binaries, and processes around after tests run

This is useful if you want to manually inspect or debug the kubelet process run as part of the tests.

```sh
make test-e2e-node REMOTE=true CLEANUP=false
```

## Run tests using an image in another project

This is useful if you want to create your own host image in another project and use it for testing.

```sh
make test-e2e-node REMOTE=true IMAGE_PROJECT="<name-of-project-with-images>" IMAGES="<image-name>"
```

Setting up your own host image may require additional steps such as installing etcd or docker.  See
[setup_host.sh](https://git.k8s.io/kubernetes/test/e2e_node/environment/setup_host.sh) for common steps to setup hosts to run node tests.

## Create instances using a different instance name prefix

This is useful if you want to create instances using a different name so that you can run multiple copies of the
test in parallel against different instances of the same image.

```sh
make test-e2e-node REMOTE=true INSTANCE_PREFIX="my-prefix"
```

## Run tests using a custom image configuration

This is useful if you want to test out different runtime configurations. First, make a local
(temporary) copy of the base image config from the test-infra repo:
https://github.com/kubernetes/test-infra/tree/master/jobs/e2e_node

Make your desired modifications to the config, and update data paths to be absolute paths to the
relevant files on your local machine (e.g. prepend your home directory path to each). For example:

```diff
 images:
   cos-stable:
     image_regex: cos-stable-60-9592-84-0
     project: cos-cloud
-    metadata: "user-data</go/src/github.com/containerd/cri/test/e2e_node/init.yaml,containerd-configure-sh</go/src/github.com/containerd/cri/cluster/gce/configure.sh,containerd-extra-init-sh</go/src/github.com/containerd/cri/test/e2e_node/gci-init.sh,containerd-env</workspace/test-infra/jobs/e2e_node/containerd/cri-master/env,gci-update-strategy=update_disabled"
+    metadata: "user-data</home/tallclair/go/src/github.com/containerd/cri/test/e2e_node/init.yaml,containerd-configure-sh</home/tallclair/go/src/github.com/containerd/cri/cluster/gce/configure.sh,containerd-extra-init-sh</home/tallclair/go/src/github.com/containerd/cri/test/e2e_node/gci-init.sh,containerd-env</home/tallclair/workspace/test-infra/jobs/e2e_node/containerd/cri-master/env,gci-update-strategy=update_disabled"
```

Finally, run the tests with your custom config:

```sh
make test-e2e-node REMOTE=true IMAGE_CONFIG_FILE="<local file>" [...]
```

Image configuration files can further influence how FOCUS and SKIP match test cases.

For example:

```sh
--focus="\[Feature:.+\]" --skip="\[Flaky\]|\[Serial\]"
```

runs all e2e tests labeled `"[Feature:*]"` and will skip any tests labeled `"[Flaky]"` or `"[Serial]"`.

Two example e2e tests that match this expression are:
  * https://github.com/kubernetes/kubernetes/blob/0e2220b4462130ae8a22ed657e8979f7844e22c1/test/e2e_node/security_context_test.go#L155
  * https://github.com/kubernetes/kubernetes/blob/0e2220b4462130ae8a22ed657e8979f7844e22c1/test/e2e_node/security_context_test.go#L175

However, image configuration files select test cases based on the `tests` field.

See https://github.com/kubernetes/test-infra/blob/4572dc3bf92e70f572e55e7ac1be643bdf6b2566/jobs/e2e_node/benchmark-config.yaml#L22-23 for an example configuration. 

If the [Prow e2e job configuration](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/containerd/image-config.yaml) does **not** specify the `tests` field, FOCUS and SKIP will run as expected.

# Additional test options for both remote and local execution

## Only run a subset of the tests

To run tests matching a regex:

```sh
make test-e2e-node REMOTE=true FOCUS="<regex-to-match>"
```

To run tests NOT matching a regex:

```sh
make test-e2e-node REMOTE=true SKIP="<regex-to-match>"
```

These are often configured in the CI environment.
For example, the [`ci-kubernetes-node-kubelet`](https://github.com/kubernetes/test-infra/blob/05eeaff67cc936181c18a63fdc9d5847c55ef258/config/jobs/kubernetes/sig-node/node-kubelet.yaml#L31) uses `--focus="\[NodeConformance\]" --skip="\[Flaky\]|\[Serial\]"`, this can be specified to the make target as:

```sh
make test-e2e-node REMOTE=true FOCUS="\[NodeConformance\]" SKIP="\[Flaky\]|\[Serial\]"
```

See http://onsi.github.io/ginkgo/#focused-specs in the Grinkgo documentation to learn more about how FOCUS and SKIP work.

## Run a single test

To run a particular e2e test, simply pass the Grinkgo `It` string to the `--focus` argument.

For example, suppose we have the following test case: https://github.com/kubernetes/kubernetes/blob/0e2220b4462130ae8a22ed657e8979f7844e22c1/test/e2e_node/security_context_test.go#L175. We could select this test case by adding the argument:

```sh
--focus="should not show its pid in the non-hostpid containers \[Feature:HostAccess\]"
```

## Run all tests related to a feature

In contrast, to run all node e2e tests related to the "HostAccess" feature one could run:

```sh
--focus="\[Feature:HostAccess\]"
```

## Run tests continually until they fail

This is useful if you are trying to debug a flaky test failure.  This will cause ginkgo to continually
run the tests until they fail.  **Note: this will only perform test setup once (e.g. creating the instance) and is
less useful for catching flakes related creating the instance from an image.**

```sh
make test-e2e-node REMOTE=true RUN_UNTIL_FAILURE=true
```

## Run tests in parallel

Running test in parallel can usually shorten the test duration. By default node
e2e test runs with`--nodes=8` (see ginkgo flag
[--nodes](https://onsi.github.io/ginkgo/#parallel-specs)). You can use the
`PARALLELISM` option to change the parallelism.

```sh
make test-e2e-node PARALLELISM=4 # run test with 4 parallel nodes
make test-e2e-node PARALLELISM=1 # run test sequentially
```

## Run tests with kubenet network plugin

[kubenet](http://kubernetes.io/docs/admin/network-plugins/#kubenet) is
the default network plugin used by kubelet since Kubernetes 1.3.  The
plugin requires [CNI](https://github.com/containernetworking/cni) and
[nsenter](http://man7.org/linux/man-pages/man1/nsenter.1.html).

Currently, kubenet is enabled by default for Remote execution `REMOTE=true`,
but disabled for Local execution.  **Note: kubenet is not supported for
local execution currently. This may cause network related test result to be
different for Local and Remote execution. So if you want to run network
related test, Remote execution is recommended.**

To enable/disable kubenet:

```sh
# enable kubenet
make test-e2e-node TEST_ARGS='--kubelet-flags="--network-plugin=kubenet --network-plugin-dir=/opt/cni/bin"'
# disable kubenet
make test-e2e-node TEST_ARGS='--kubelet-flags="--network-plugin= --network-plugin-dir="'
```

## Additional QoS Cgroups Hierarchy level testing

For testing with the QoS Cgroup Hierarchy enabled, you can pass --cgroups-per-qos flag as an argument into Ginkgo using TEST_ARGS

```sh
make test_e2e_node TEST_ARGS="--cgroups-per-qos=true"
```
## Testgrid

TestGrid (https://testgrid.k8s.io) is a publicly hosted and configured automated testing framework developed by Google.

Here (https://testgrid.k8s.io/sig-node-containerd#containerd-node-features) we see an example of an e2e Prow job running e2e tests. Each gray row in the grid corresponds
to an e2e test or a stage of the job (i.e. created a VM, downloaded some files). Passed tests are colored green and failed tests are colored red.

# Notes on tests run by the Kubernetes project during pre-, post- submit.

The node e2e tests are run by the [Prow](https://prow.k8s.io/) for each Pull Request and the results published 
in the status checks box at the bottom of the Pull Request below all comments. To have Prow re-run the node e2e tests against a PR add the comment
`/test pull-kubernetes-node-e2e` and **include a link to the test failure logs if caused by a flake.**
Note that [commands to Prow](https://prow.k8s.io/command-help#test) must be on separate lines from any commentary.

For example, 

    /test pull-kubernetes-node-e2e
    flake due to #12345

The PR builder runs tests against the images listed in [image-config.yaml](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/containerd/image-config.yaml).

Other [node e2e Prow jobs](https://github.com/kubernetes/test-infra/tree/master/config/jobs/kubernetes/sig-node)
run against different images depending on the configuration chosen in the
[test-infra repo](https://github.com/kubernetes/test-infra/tree/master/jobs/e2e_node).
The source code for these tests comes from the [kubernetes/kubernetes repo](https://github.com/kubernetes/kubernetes/tree/master/test/e2e_node).

# Notes on the Topology Manager tests

The Topology Manager tests require a multi-numa node box (two or more nodes) with at least one SRIOV device installed to run.
The tests automatically skip if the conditions aren't met.

The test code statically includes the manifests needed to configure the SRIOV device plugin.
However, is not possible to anticipate all the possible configuration, hence the included configuration is intentionally minimal.

It is recommended you supply a ConfigMap describing the cards installed in the machine running the tests using TEST_ARGS.
[Here's the upstream reference](https://github.com/intel/sriov-network-device-plugin/blob/master/deployments/configMap.yaml)
```sh
make test-e2e-node TEST_ARGS='--sriovdp-configmap-file="/path/to/sriovdp-config-map.yaml"'
```

You must have the Virtual Functions (VFs) already created in the node you want to run the test on.
Example command to create the VFs - please note the PCI address of the SRIOV device depends on the host
system hardware.
```bash
cat /sys/devices/pci0000:00/0000:00:01.0/0000:01:00.1/sriov_numvfs
echo 7 > /sys/devices/pci0000:00/0000:00:01.0/0000:01:00.1/sriov_numvfs
```

Some topology manager tests require minimal knowledge of the host topology in order to be performed.
The required information is to which NUMA node in the system are the SRIOV device attached to.
The test code tries to autodetect the information it needs, skipping the relevant tests if the autodetection fails.

You can override the autodetection adding annotations to the config map like this example:
```yaml
metadata:
  annotations:
    pcidevice_node0: "1"
    pcidevice_node1: "0"
    pcidevice_node2: "0"
    pcidevice_node3: "0"
```

Please note that if you add the annotations, then you must provide the full information:
you must should specify the number of SRIOV devices attached to each NUMA node in the system,
even if the number is zero.

# Debugging E2E Tests Locally

1. Install kubectl on the node
2. Set your KUBCONFIG environment variable to reference the kubeconfig created by the e2e node tests
   `export KUBECONFIG=./_output/local/go/bin/kubeconfig`
3. Inspect the node and pods as needed while the tests are running
   ```
   $ kubectl get pod -A
   $ kubectl describe node
   ```
