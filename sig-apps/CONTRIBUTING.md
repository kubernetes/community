# Contributing to SIG Apps

Welcome to contributing to SIG Apps.

SIG Apps has multiple areas you can contribute to. Those contributions can be in
the form of code, documentation, support, being involved in mailing list
discussions, attending meetings, and more. This guide describes different
major functional areas SIG Apps is involved in, provides an overview of the
areas, and gives pointers on getting more involved in each area. Consider this
a launching point or the start of a [choose your own adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure)
for SIG Apps.

## Major Areas

The major launching off point for SIG Apps comes in the form of the different
areas we're involved in. Each of the areas has a different scope, a slightly
different feel to the active developers, and small differences to the way
things are done.

To get an idea of what's going on, here's an overview of each of the areas.

### Workloads API

The _Workloads API_ including controllers and jobs is part of SIG Apps.The
controllers are Deployments, DaemonSets, StatefulSets, and ReplicaSets. The jobs
are part of the batch API and include Jobs and CronJobs. These are part of core
Kubernetes and developed within the [monorepo](https://github.com/kubernetes/kubernetes).

Don't let the fact that these are part of core Kubernetes intimidate you. Each
controller is its own package that can be found in the [pkg/controllers directory](https://github.com/kubernetes/kubernetes/tree/master/pkg/controller). Each
controller is small.

Like other parts of core Kubernetes, you can learn about contributing code or
reviewing pull requests via the [Contributors Guide](https://github.com/kubernetes/community/tree/master/contributors/guide).

### Application CRD and Controller

The [Application CRD and controller](https://github.com/kubernetes-sigs/application)
provides a method to group application resources and describe the application.
This SIG sponsored project has a fairly minimal and straight forward [contributing
process](https://github.com/kubernetes-sigs/application/blob/master/CONTRIBUTING.md)

### Kompose

[Kompose](https://github.com/kubernetes/kompose) stands for Kubernetes + Compose
(the docker tooling). This project translates Docker compose configuration into
Kubernetes configuration.

Kompose has its own [contributing guide](https://github.com/kubernetes/kompose/blob/master/CONTRIBUTING.md).

## Contributing Beyond Code

There are a couple ways to contribute beyond code and the sub-projects. They include:

* In almost all SIG Apps meetings there are **demos**. Those can include tools to
  help with running applications, new ways of building workloads, methods for
  tying different tools together, and more. If you are interested in
  contributing a demo please contact the SIG leads via the
  [mailing list](https://groups.google.com/a/kubernetes.io/g/sig-apps)
* Documentation is a vital resource for those using Kubernetes. The documentation
  teaches concepts, usage of resources, how to accomplish tasks, and more. When
  it comes to running applications, such as when leveraging workloads, the
  documentation has room for improvement. If you know how to run applications
  in Kubernetes, consider [contributing to better documentation](https://kubernetes.io/docs/contribute/).
