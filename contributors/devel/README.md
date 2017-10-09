# Kubernetes Developer Guide

The developer guide is for anyone wanting to either write code which directly accesses the
Kubernetes API, or to contribute directly to the Kubernetes project.
It assumes some familiarity with concepts in the [User Guide](http://kubernetes.io/docs/user-guide/) and the [Cluster Admin
Guide](http://kubernetes.io/docs/admin/).


## The process of developing and contributing code to the Kubernetes project

* **Welcome to Kubernetes (New Developer Guide)**
  ([welcome-to-kubernetes-new-developer-guide.md](welcome-to-kubernetes-new-developer-guide.md)):
  An introductory guide to contributing to K8s.

* **On Collaborative Development** ([collab.md](collab.md)): Info on pull requests and code reviews.

* **GitHub Issues** ([issues.md](issues.md)): How incoming issues are triaged.

* **Pull Request Process** ([pull-requests.md](pull-requests.md)): When and why pull requests are closed.

* **Kubernetes On-Call Rotations** ([on-call-rotations.md](on-call-rotations.md)): Descriptions of on-call rotations for build and end-user support.

* **Faster PR reviews** ([faster_reviews.md](faster_reviews.md)): How to get faster PR reviews.

* **Getting Recent Builds** ([getting-builds.md](getting-builds.md)): How to get recent builds including the latest builds that pass CI.

* **Automated Tools** ([automation.md](automation.md)): Descriptions of the automation that is running on our github repository.


## Setting up your dev environment, coding, and debugging

* **Development Guide** ([development.md](development.md)): Setting up your development environment.

* **Testing** ([testing.md](testing.md)): How to run unit, integration, and end-to-end tests in your development sandbox.

* **Hunting flaky tests** ([flaky-tests.md](flaky-tests.md)): We have a goal of 99.9% flake free tests.
  Here's how to run your tests many times.

* **Logging Conventions** ([logging.md](logging.md)): Glog levels.

* **Profiling Kubernetes** ([profiling.md](profiling.md)): How to plug in go pprof profiler to Kubernetes.

* **Instrumenting Kubernetes with a new metric**
  ([instrumentation.md](instrumentation.md)): How to add a new metrics to the
  Kubernetes code base.

* **Coding Conventions** ([coding-conventions.md](coding-conventions.md)):
  Coding style advice for contributors.

* **Document Conventions** ([how-to-doc.md](how-to-doc.md))
  Document style advice for contributors.

* **Running a cluster locally** ([running-locally.md](running-locally.md)):
  A fast and lightweight local cluster deployment for development.

## Developing against the Kubernetes API

* The [REST API documentation](http://kubernetes.io/docs/reference/) explains the REST
  API exposed by apiserver.

* **Annotations** ([Annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)): are for attaching arbitrary non-identifying metadata to objects.
  Programs that automate Kubernetes objects may use annotations to store small amounts of their state.

* **API Conventions** ([api-conventions.md](api-conventions.md)):
  Defining the verbs and resources used in the Kubernetes API.

* **API Client Libraries** ([client-libraries.md](client-libraries.md)):
  A list of existing client libraries, both supported and user-contributed.


## Writing plugins

* **Authentication** ([Authentication](http://kubernetes.io/docs/admin/authentication/)):
  The current and planned states of authentication tokens.

* **Authorization Plugins** ([Authorization](http://kubernetes.github.io/docs/admin/authorization/)):
  Authorization applies to all HTTP requests on the main apiserver port.
  This doc explains the available authorization implementations.

* **Admission Control Plugins** ([admission_control](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/admission_control.md))


## Building releases

See the [kubernetes/release](https://github.com/kubernetes/release) repository for details on creating releases and related tools and helper scripts.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/README.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
