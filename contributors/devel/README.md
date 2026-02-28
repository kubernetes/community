# Table of Contents

The developer guide is for anyone wanting to either write code which directly accesses the
Kubernetes API, or to contribute directly to the Kubernetes project.
It assumes some familiarity with concepts in the [User Guide](https://kubernetes.io/docs/concepts/) and the [Cluster Admin
Guide](https://kubernetes.io/docs/concepts/cluster-administration/).


## The process of developing and contributing code to the Kubernetes project

* **Contributor Guide**
  ([Please start here](/contributors/guide/README.md)) to learn about how to contribute to Kubernetes.

* **GitHub Issues** ([/contributors/guide/issue-triage.md](/contributors/guide/issue-triage.md)): How incoming issues are triaged.

* **Pull Request Process** ([/contributors/guide/pull-requests.md](/contributors/guide/pull-requests.md)): When and why pull requests are closed.

* **Getting Recent Builds** ([getting-builds.md](sig-release/getting-builds.md)): How to get recent builds including the latest builds that pass CI.

* **Automated Tools** ([automation.md](automation.md)): Descriptions of the automation that is running on our github repository.


## Setting up your dev environment, coding, and debugging

* **Development Guide** ([development.md](development.md)): Setting up your development environment.

* **Testing** ([testing.md](sig-testing/testing.md)): How to run unit, integration, and end-to-end tests in your development sandbox.

* **Conformance Testing** ([conformance-tests.md](sig-architecture/conformance-tests.md))
  What is conformance testing and how to create/manage them.

* **Hunting flaky tests** ([flaky-tests.md](sig-testing/flaky-tests.md)): We have a goal of 99.9% flake free tests.
  Here's how to run your tests many times.

* **Logging Conventions** ([logging.md](sig-instrumentation/logging.md)): klog levels.

* **Profiling Kubernetes** ([profiling.md](sig-scalability/profiling.md)): How to plug in go pprof profiler to Kubernetes.

* **Instrumenting Kubernetes with a new metric**
  ([instrumentation.md](sig-instrumentation/metric-instrumentation.md)): How to add a new metrics to the
  Kubernetes code base.

* **Coding Conventions** ([coding-conventions.md](../guide/coding-conventions.md)):
  Coding style advice for contributors.

* **Document Conventions** ([The Kubernetes documentation](https://github.com/kubernetes/website))
  Document style advice for contributors.

* **Running a cluster locally** ([running-locally.md](running-locally.md)):
  A fast and lightweight local cluster deployment for development.

## Developing against the Kubernetes API

* The [REST API documentation](https://kubernetes.io/docs/reference/) explains the REST
  API exposed by apiserver.

* **Annotations** ([Annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)): are for attaching arbitrary non-identifying metadata to objects.
  Programs that automate Kubernetes objects may use annotations to store small amounts of their state.

* **API Conventions** ([api-conventions.md](sig-architecture/api-conventions.md)):
  Defining the verbs and resources used in the Kubernetes API.

* **API Client Libraries** ([Client Libraries](https://kubernetes.io/docs/reference/using-api/client-libraries/)):
  A list of existing client libraries, both supported and user-contributed.


## Writing plugins

* **Authentication** ([Authentication](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)):
  The current and planned states of authentication tokens.

* **Authorization Plugins** ([Authorization](https://kubernetes.io/docs/reference/access-authn-authz/authorization/)):
  Authorization applies to all HTTP requests on the main apiserver port.
  This doc explains the available authorization implementations.

* **Admission Control Plugins** ([admission_control](https://git.k8s.io/design-proposals-archive/api-machinery/admission_control.md))


## Building releases

See the [kubernetes/release](https://github.com/kubernetes/release) repository for details on creating releases and related tools and helper scripts.

## SIG Developer Guide Contributions

### SIG Release
* **Cherry Picks** [cherry-picks.md](sig-release/cherry-picks.md)
  How cherry picks are managed on release branches within the `kubernetes/kubernetes` repository.

* **Getting Kubernetes Builds** [getting-builds.md](sig-release/getting-builds.md)

* **Targeting enhancements, Issues and PRs to Release Milestones** [release.md](sig-release/release.md)

### SIG Instrumentation
* **Logging Conventions** [logging.md](sig-instrumentation/logging.md)

* **Event style guide** [event-style-guide.md](sig-instrumentation/event-style-guide.md)

* **Instrumenting Kubernetes with a new metric** [instrumentation.md](sig-instrumentation/metric-instrumentation.md)

* **Structured Logging migration instructions** [migration-to-structured-logging.md](sig-instrumentation/migration-to-structured-logging.md)

### SIG Storage
* **NOTE** Flexvolume is deprecated. Out-of-tree CSI driver is the recommended way to write volume drivers in Kubernetes. See this doc [here]( https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md) for more information.

* **CSI Drivers Doc** [CSI drivers doc](https://kubernetes-csi.github.io/docs/)
  This site documents how to develop, deploy, and test a [Container Storage Interface](https://github.com/container-storage-interface/spec/blob/master/spec.md) (CSI) driver on Kubernetes.

* **Flexvolume** [flexvolume.md](sig-storage/flexvolume.md)
  Flexvolume enables users to write their own drivers and add support for their volumes in Kubernetes.

### SIG Scalability
* **Kubemark User Guide** [kubemark-guide.md](sig-scalability/kubemark-guide.md)

* **How to Set Up a Kubemark Cluster** [kubemark-setup-guide.md](sig-scalability/kubemark-setup-guide.md)

* **Profiling Kubernetes** [profiling.md](sig-scalability/profiling.md)

### SIG Scheduling

* **Understanding the Kubernetes Scheduler** [scheduling_code_hierarchy_overview.md](sig-scheduling/scheduling_code_hierarchy_overview.md)

* **Customizing the Kubernetes Scheduler** [scheduler configuration](https://kubernetes.io/docs/reference/scheduling/config/)

* **Understanding how Pods are queued in Kubernetes Scheduler** [scheduler_queues.md](sig-scheduling/scheduler_queues.md)

* **Scheduler Benchmarking** [scheduler_benchmarking.md](sig-scheduling/scheduler_benchmarking.md)

### SIG Architecture

* **API Conventions** [api-conventions.md](sig-architecture/api-conventions.md)

* **Component Configuration Conventions** [component-config-conventions.md](sig-architecture/component-config-conventions.md)

* **Changing the API** [api_changes.md](sig-architecture/api_changes.md)

* **Staging Directory and Publishing** [staging.md](sig-architecture/staging.md)

* **Using Go Modules to Manage Dependencies** [vendor.md](sig-architecture/vendor.md)
  This document only applies to Kubernetes development after 1.14.x. See [previous godep documentation for working with dependencies](sig-architecture/godep.md) for Kubernetes 1.14.x and earlier.

* **Using Go Modules to Manage Dependencies (for Kubernetes 1.14.x and earlier)** [godep.md](sig-architecture/godep.md)
  See [current documentation for working with dependencies](sig-architecture/vendor.md) for master branch development.

* **Conformance Testing in Kubernetes** [conformance-tests.md](sig-architecture/conformance-tests.md)

### SIG API Machinery
* **Strategic Merge Patch** [strategic-merge-patch.md](sig-api-machinery/strategic-merge-patch.md)
* **Writing Controllers** [controllers.md](sig-api-machinery/controllers.md)
* **Generation and release cycle of clientset** [generating-clientset.md](sig-api-machinery/generating-clientset.md)

### SIG Testing
* **Testing guide** [testing.md](sig-testing/testing.md)

* **Writing good e2e tests for Kubernetes** [writing-good-e2e-tests.md](sig-testing/writing-good-e2e-tests.md)

* **Writing Good Conformance Tests for Kubernetes** [writing-good-conformance-tests.md](sig-testing/writing-good-conformance-tests.md)

* **Integration Testing in Kubernetes** [integration-tests.md](sig-testing/integration-tests.md)

* **End-to-End Testing in Kubernetes** [e2e-tests.md](sig-testing/e2e-tests.md) and [e2e-tests-kubetest2.md](sig-testing/e2e-tests-kubetest2.md)

* **Flaky tests** [flaky-tests.md](sig-testing/flaky-tests.md)

### SIG Node

* **CRI: the Container Runtime Interface** [container-runtime-interface.md](sig-node/container-runtime-interface.md)

* **Container Runtime Interface (CRI) Networking Specifications** [kubelet-cri-networking.md](sig-node/kubelet-cri-networking.md)

* **Measuring Node Performance** [node-performance-testing.md](sig-node/node-performance-testing.md)

* **Container Runtime Interface: Container Metrics** [cri-container-stats.md](sig-node/cri-container-stats.md)

* **Node End-To-End tests** [e2e-node-tests.md](sig-node/e2e-node-tests.md)

* **Container Runtime Interface: Testing Policy** [cri-testing-policy.md](sig-node/cri-testing-policy.md)

### SIG CLI
* **Kubectl Conventions** [kubectl-conventions.md](sig-cli/kubectl-conventions.md)
