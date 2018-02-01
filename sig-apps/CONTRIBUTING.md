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

### Helm

[Helm](https://helm.sh) is the package manager for Kubernetes. The packages for
Helm are called Charts. Helm version 2 is the current stable release series and
is currently focused on stability, minor non-breaking feature additions, and
better documentation.

Helm is a sub-project on a much smaller scale from Kubernetes core, the most
active repository on GitHub. Contributing to Helm and navigating the project is
similar to many other open source projects.

The [Helm Contributing Guide](https://github.com/kubernetes/helm/blob/master/CONTRIBUTING.md)
contains much of what you need to get started including:

* Support and conversation channels
* Filing issues and the lifecycle of issues
* Project milestones and what [Semantic Versioning](http://semver.org) means to Helm
* Details on contributing a patch, via a pull request, and what to expect

The Helm maintainers have a weekly public meeting that's open to anyone to attend.
The meeting is recorded and available on YouTube for those unable to attend.
Details on the meeting are in the [Readme](README.md) for SIG Apps.

In addition to Helm itself, there are a number of sub-projects of Helm. These
projects can be found on the [kubernetes-helm](https://github.com/kubernetes-helm).
Each of these projects has its own contribution guide but is part of the broader
Helm project.

### Community Charts

[Charts](https://github.com/kubernetes/charts) is a community curated set of Helm
packages for Kubernetes. While organizations and individuals are encouraged to
create their own charts – something Helm provides tools for – the community
charts are a place for people to use, share, and contribute to.

There are a few ways to contribute to charts:

1. Create or improve charts
1. Contribute to continuous testing
1. Write down best practices or document other details

Charts has its own [contributing guide](https://github.com/kubernetes/charts/blob/master/CONTRIBUTING.md)
and [review guidelines](https://github.com/kubernetes/charts/blob/master/REVIEW_GUIDELINES.md)
that can act as a launching off point for involvement.

The charts maintainers have a weekly meeting in addition to the normal SIG Apps
meeting. Details on the meeting are in the [Readme](README.md) for SIG Apps. 

### Kompose

[Kompose](https://github.com/kubernetes/kompose) stands for Kubernetes + Compose
(the docker tooling). This project translates Docker compose configuration into
Kubernetes configuration.

Kompose, like Helm and the community charts, has its own [contributing guide](https://github.com/kubernetes/kompose/blob/master/CONTRIBUTING.md).

## Contributing Beyond Code

There are a couple ways to contribute beyond code and the sub-projects. They include:

* In almost all SIG Apps meetings there are **demos**. Those can include tools to
  help with running applications, new ways of building workloads, methods for
  tying different tools together, and more. If you are interested in
  contributing a demo please contact the SIG leads via the
  [mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-apps)
* The **App Def Working Group** is a cross collaboration with SIG CLI, SIG API
  Machinery, and others to look at how we can improve the core tooling and
  process for operating applications. To produce real world solutions we want
  input from people who operate applications. If you are interested in learning
  more please contact the App Def Working Group mailing list or attend a
  meeting. You can find out more on the
  [working groups organization page](https://github.com/kubernetes/community/tree/master/wg-app-def)