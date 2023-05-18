# End-to-End Testing in Kubernetes

**Table of Contents**

- [End-to-End Testing in Kubernetes](#end-to-end-testing-in-kubernetes)
  - [Overview](#overview)
  - [Setting up kubetest2](#setting-up-kubetest2)
  - [Running the Tests](#running-the-tests)
    - [Cleaning Up](#cleaning-up)
  - [Advanced Testing](#advanced-testing)
    - [Extracting a specific version of Kubernetes](#extracting-a-specific-version-of-kubernetes)
    - [Testing against an existing cluster](#testing-against-an-existing-cluster)
    - [Debugging clusters](#debugging-clusters)

## Overview

End-to-end (e2e) tests for Kubernetes provide a mechanism to test end-to-end
behavior of the system. It is the last signal to ensure end user operations
match developer specifications. Although unit and integration tests provide a
good signal, in a distributed system like Kubernetes it is not uncommon that a
minor change may pass all unit and integration tests, but cause unforeseen
changes at the system level.

The primary objectives of the e2e tests are to ensure a consistent and reliable
behavior of the Kubernetes code base, and to catch hard-to-test bugs that were not caught by unit and integration tests.

The e2e tests in Kubernetes are built on top of
[Ginkgo](http://onsi.github.io/ginkgo/) and
[Gomega](http://onsi.github.io/gomega/). There are a host of features that this
Behavior-Driven Development (BDD) testing framework provides, and it is
recommended that the developer read the documentation prior to diving into the
tests.

The purpose of this document is to serve as a primer for developers who are
looking to execute or add tests using a local development environment.

Before writing new tests or making substantive changes to existing tests, you
should also read [Writing Good e2e Tests](writing-good-e2e-tests.md)

## Setting up kubetest2

The e2e tests in Kubernetes are managed and run with a helper
application. On July 14, 2020, it was announced that the venerable
[kubetest](https://github.com/kubernetes/test-infra/tree/master/kubetest#installation)
would be deprecated in favor of
[kubetest2](https://github.com/kubernetes-sigs/kubetest2). These
instructions walk you through the installation and configuration of
`kubetest2` and then describe some usage.

Before proceeding to these instructions, you should have a working
Kubernetes development environment. Please follow the directions in
the [Developer Guide](../development.md) to get started. You should
also have a working and configured `gcloud` CLI tool for use with
Google Cloud 
Platform. [Check out the `gcloud` cheatsheet for help getting started.](https://cloud.google.com/sdk/docs/cheatsheet)

The first step is to install the Ginkgo testing framework and
Gomega. Run these commands:

```sh
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...
```

In order to successfully build Kubernetes tests, you will need to
install Bazel.
[Follow the instructions for your development environment.](https://docs.bazel.build/versions/3.4.0/install.html)

Next, you need to install the `kubetest2` binary and plugins. Run
these commands to do that:

```sh
cd
GO111MODULE=on go get sigs.k8s.io/kubetest2/...@latest
```

The currently available plugins for cloud providers are:

- Google Cloud Compute Engine (kubetest2-gce)
- Google Cloud Kubernetes Engine (kubetest2-gke)
- [KIND](https://kind.sigs.k8s.io/) (kubetest2-kind)

## Building Kubernetes

To reliably and repeatedly test changes locally, it is important to
minimize the number of steps the developer needs to take. To
facilitate this, you can use `kubetest2` to build Kubernetes using the
following command: 

```sh
kubetest2 gce --build --legacy-mode
```

To build without the `--legacy-mode` option, you will need to clone
the
[`gcp-cloud-provider` repository](https://github.com/kubernetes/cloud-provider-gcp). 
This repository is part of the
[cloud provider extraction effort](https://github.com/kubernetes/community/tree/master/sig-cloud-provider#cloud-provider-extraction-migration)
and the eventual home of some of the build scripts in the Kubernetes
source tree. Use this command to clone:

```sh
git clone https://github.com/kubernetes/cloud-provider-gcp.git
```

You can now use the following command to build using the
`gcp-cloud-provider` repository. Replace `<path to
cloud-provider-gcp>` with the path to your cloned repository:

```sh
kubetest2 gce --build --repo-root <path to cloud-provider-gcp>
```

## Running the Tests

The examples below all use the Google Cloud Compute Engine plugin
(kubetest2-gce). Different cloud provider plugins use different
command-line switches.

In these examples, replace `<project>` with the name of the
Google Cloud project in which you are running e2e tests.

To provision and start a Kubernetes test cluster, run this command:

```sh
kubetest2 gce --gcp-project <project> --up
```

You can also shut down and destroy a testing cluster with this
command:

```sh
kubetest2 gce --gcp-project <project> --down
```

You do not need the `--gcp-project` argument to run tests. While your
test cluster is up and running, you can run all tests with this
command:

```sh
kubetest2 gce --test ginkgo
```

You can also pass options to Ginkgo to select which tests you'd like
to run. If you add a `--` at the end of the command line, everything
after that is passed as arguments to Ginkgo.

For example, to run all tests that match the regular expression
`\[Feature:Performance\]`, use this command:

```sh
kubetest2 gce --test ginkgo -- --focus-regex "\[Feature:Performance\]"
```

Conversely, to exclude tests that match the regular expression
`Pods.*env`, use this command:

```sh
kubetest2 gce --test ginkgo -- --skip-regex "Pods.*env"
```

In this example, we instruct `kubetest2` to run two tests in parallel
at once, while skipping any that must be run serially:

```sh
kubetest2 gce --test ginkgo -- --skip-regex "\[Serial\]" --parallel 2
```

See `kubetest2 gce --help` for more options. Note that the different
deployer and tester plugins have their own additional options, which
can be seen in their help listings.

### Cleaning Up

During a run, pressing **Control-C** should result in an orderly
shutdown. However, if something goes wrong and you still have
containers or VMs provisioned, you can run a cleanup with a command
like this:

```sh
kubetest2 gce --gcp-project <project> --down
```

## Advanced Testing

### Extracting a specific version of Kubernetes

It is possible for `kubetest2` to download and extract a specific
version of Kubernetes for testing. This can be accomplished by passing
the `--test-package-version` flag to the tester plugin. For example:

```sh
kubetest2 gce --test ginkgo -- --test-package-version v1.18.0
```

The argument to `--test-package-version` can be changed to specify
some other test package version. To see available release names for
this option, visit https://github.com/kubernetes/kubernetes/releases.

Examples: v1.26.0-alpha.2, v1.26.0-beta.0, v1.26.1-rc.0, v1.26.5


### Testing against an existing cluster

You can run tests against an existing Kubernetes cluster by setting
your `KUBECONFIG` environment variable to point to that cluster's
configuration file.

```sh
KUBECONFIG=<path to config file> kubetest2 gce --test
```

It doesn't matter which deployment plugin you specify on the command
line.

### Debugging clusters

The configuration file for your `kubetest2` testing cluster is kept in
`_artifacts/kubetest2-kubeconfig` off of your Kubernetes working
directory. If you would like to get a better understanding of your
cluster's state in the event of failed tests, you can use the
following command to collect information.

```sh
KUBECONFIG=./_artifacts/kubetest2-kubeconfig kubectl cluster-info dump
```

If you are using the `gce` deployment plugin, you can also use the
`cluster/log-dump/log-dump.sh` script to copy the logs and
configuration files from your test cluster to your local machine. 

Use the following command to dump all cluster logs:

```sh
KUBE_GCE_INSTANCE_PREFIX=kubetest2 KUBECONFIG=$PWD/_artifacts/kubetest2-kubeconfig ./cluster/log-dump/log-dump.sh
```

Note that `kubetest2` also automatically dumps all logs to
`_artifacts/cluster-logs/` when you bring up your test cluster with
the `--up` command.
