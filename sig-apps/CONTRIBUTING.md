# Contributing to SIG Apps

Welcome to contributing to SIG Apps. We are excited about the prospect of you
joining our [community](https://github.com/kubernetes/community/tree/master/sig-apps)!

SIG Apps has multiple areas you can contribute to. Those contributions can be in the form of code, documentation, support, being involved in mailing list
discussions, attending meetings, and more. This guide describes different
major functional areas SIG Apps is involved in, provides an overview of the
areas, and gives pointers on getting more involved in each area.

## Before you begin

We strongly recommend you to check out the [Kubernetes Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide)
and [Contributor Cheat Sheet](https://github.com/kubernetes/community/tree/master/contributors/guide/contributor-cheatsheet) first.

## Getting Started

* If you're a newcomer and have no idea where to start, we have a non-stale pool of issues that are
available for you:
  * [first-good-issue](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22+label%3Asig%2Fapps+)
  * [help wanted](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22+label%3Asig%2Fapps+)

* The _Workloads API_ including controllers and jobs is part of SIG Apps.The
controllers are Deployments, DaemonSets, StatefulSets, and ReplicaSets. The jobs
are part of the batch API and include Jobs and CronJobs. These are part of core
Kubernetes and developed within the [monorepo](https://github.com/kubernetes/kubernetes).Each controller has its own package that can be found in the [pkg/controllers directory](https://github.com/kubernetes/kubernetes/tree/master/pkg/controller).
* The [Application CRD and controller](https://github.com/kubernetes-sigs/application) provides a method to group application resources and describe the application.
This SIG sponsored project has a fairly minimal and straight forward [contributing
process](https://github.com/kubernetes-sigs/application/blob/master/CONTRIBUTING.md)

* [Kompose](https://github.com/kubernetes/kompose) stands for Kubernetes + Compose
(the docker tooling). This project translates Docker compose configuration into
Kubernetes configuration. Kompose has its own [contributing guide](https://github.com/kubernetes/kompose/blob/master/CONTRIBUTING.md).

* Documentation is a vital resource for those using Kubernetes. The documentation teaches concepts, usage of resources, how to accomplish tasks, and more. When it comes to running applications, such as when leveraging workloads, the documentation has room for improvement. If you know how to run applications in Kubernetes, consider [contributing to better documentation](https://kubernetes.io/docs/contribute/).
