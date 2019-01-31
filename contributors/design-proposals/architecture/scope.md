# Kubernetes scope

Purpose of this doc: Clarify factors affecting decisions regarding
what is and is not in scope for the Kubernetes project.

Related documents:
* [What is Kubernetes?](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)
* [Kubernetes design and architecture](architecture.md)
* [Kubernetes architectural roadmap (2017)](architectural-roadmap.md)
* [Design principles](principles.md)
* [Kubernetes resource management](resource-management.md)

Kubernetes is a portable, extensible open-source platform for managing
containerized workloads and services, that facilitates both
declarative configuration and automation.  Workload portability is an
especially high priority. Kubernetes provides a flexible, easy-to-run,
secure foundation for running containerized applications on any cloud
provider or your own systems.

While not a full distribution in the Linux sense, adoption of
Kubernetes has been facilitated by the fact that the upstream releases
are usable on their own, with minimal dependencies (e.g., etcd, a
container runtime, and a networking implementation). 

The high-level scope and goals are often insufficient for making
decisions about where to draw the line, so this documents where the
line is, the rationale for some past decisions, and some general
criteria that have been applied, including non-technical
considerations. For instance, user adoption and continued operation of
the project itself are also important factors.

## Significant areas

More details can be found below, but a concise list of areas in scope follows:
* Containerized workload execution and management
* Service discovery, load balancing, and routing
* Workload identity propagation and authentication
* Declarative resource management platform
* Command-line tool
* Web dashboard (UI)
* Cluster lifecycle tools
* Extensibility to support execution and management in diverse environments
* Multi-cluster management tools and systems
* Project GitHub automation and other process automation
* Project continuous build and test infrastructure
* Release tooling
* Documentation
* Usage data collection mechanisms

## Scope domains

Most decisions are regarding whether any part of the project should
undertake efforts in a particular area. However, some decisions may
sometimes be necessary for smaller scopes. The term "core" is sometimes
used, but is not well defined. The following are scopes that may be relevant:
* Kubernetes project github orgs
  * All github orgs
  * The kubernetes github org
  * The kubernetes-sigs and kubernetes-incubator github orgs
  * The kubernetes-client github org
  * Other github orgs
* Release artifacts
  * The Kubernetes release bundle
  * Binaries built in kubernetes/kubernetes
    * “core” server components: apiserver, controller manager, scheduler, kube-proxy, kubelet
    * kubectl
    * kubeadm
  * Other images, packages, etc.
* The kubernetes/kubernetes repository (aka k/k)
  * master branch
  * kubernetes/kubernetes/master/pkg
  * kubernetes/kubernetes/master/staging
* [Functionality layers](architectural-roadmap.md)
  * required
  * pluggable
  * optional
  * usable independently of the rest of Kubernetes

## Other inclusion considerations

The Kubernetes project is a large, complex effort.

* Is the functionality consistent with the existing implementation
  conventions, design principles, architecture, and direction?

* Do the subproject owners, approvers, reviewers, and regular contributors
  agree to maintain the functionality?

* Do the contributors to the functionality agree to follow the
  project’s development conventions and requirements, including CLA,
  code of conduct, github and build tooling, testing, documentation,
  and release criteria, etc.?

* Does the functionality improve existing use cases, or mostly enable
  new ones? The project isn't completely blocking new functionality
  (more reducing the rate of expansion), but it is trying to
  limit additions to kubernetes/kubernetes/master, and aims to improve the
  quality of the functionality that already exists.

* Is it needed by project contributors? Example: We need cluster
  creation and upgrade functionality in order to run end-to-end tests.

* Is it necessary in order to enable workload portability?

* Is it needed in order for upstream releases to be usable? For
  example, things without which users otherwise were
  reverse-engineering Kubernetes to figure out, and/or copying code
  out of Kubernetes itself to make work.

* Is it functionality that users expect, such as because other
  container platforms and/or service discovery and routing mechanisms
  provide it? If a capability that relates to Kubernetes's fundamental
  purpose were to become table stakes in the industry, Kubernetes
  would need to support it in order to stay relevant. (Whether it
  would need to be addressed by the core project would depend on the
  other criteria.)

* Is there sufficiently broad user demand and/or sufficient expected
  user benefit for the functionality?

* Is there an adequate mechanism to discover, deploy, express a
  dependency on, and upgrade the functionality if implemented using an
  extension mechanism? Are there consistent notions of releases, maturity,
  quality, version skew, conformance, etc. for extensions?

* Is it needed as a reference implementation exercising extension
  points or other APIs?

* Is the functionality sufficiently general-purpose?

* Is it an area where we want to provide an opinionated solution
  and/or where fragmentation would be problematic for users, or are
  there many reasonable alternative approaches and solutions to the
  problem?

* Is it an area where we want to foster exploration and innovation in
  the ecosystem?

* Has the ecosystem produced adequate solutions on its own? For
  instance, have ecosystem projects taken on requirements of the
  Kubernetes project, if needed? Example: etcd3 added a number of features
  and other improvements to benefit Kubernetes, so the project didn't
  need to launch a separate storage effort.

* Is there an acceptable home for the recommended ecosystem solution(s)?
  Example: the [CNCF Sandbox](https://github.com/cncf/toc/blob/master/process/sandbox.md) is one possible home

* Has the functionality been provided by the project/release/component
  historically?

## Technical scope details and rationale

### Containerized workload execution and management

Including:
* common general categories of workloads, such as stateless, stateful, batch, and cluster services
* provisioning, allocation, accessing, and managing compute, storage, and network resources on behalf of the workloads, and enforcement of security policies on those resources
* workload prioritization, capacity assessment, placement, and relocation (aka scheduling)
* graceful workload eviction
* local container image caching
* configuration and secret distribution
* manual and automatic horizontal and vertical scaling
* deployment, progressive (aka rolling) upgrades, and downgrades
* self-healing
* exposing container logs, status, health, and resource usage metrics for collection

### Service discovery, load balancing, and routing

Including:
* endpoint tracking and discovery, including pod and non-pod endpoints
* the most common L4 and L7 Internet protocols (TCP, UDP, SCTP, HTTP, HTTPS)
* intra-cluster DNS configuration and serving
* external DNS configuration
* accessing external services (e.g., imported services, Open Service Broker)
* exposing traffic latency, throughput, and status metrics for collection
* access authorization

### Workload identity propagation and authentication

Including:
* internal identity (e.g., SPIFFE support)
* external identity (e.g., TLS certificate management)

### Declarative resource management platform

Including:
* CRUD API operations and behaviors, diff, patch, dry run, watch
* declarative updates (apply)
* resource type definition, registration, discovery, documentation, and validation mechanisms
* pluggable authentication, authorization, admission (API-level policy enforcement), and audit-logging mechanisms
* Namespace (resource scoping primitive) lifecycle
* resource instance persistence and garbage collection
* asynchronous event reporting
* API producer SDK
* API client SDK / libraries in widely used languages
* dynamic, resource-oriented CLI, as a reference implementation for interacting with the API and basic tool for declarative and imperative management
  * simplifies getting started and avoids complexities of documenting the system with just, for instance, curl

### Command-line tool

Since some Kubernetes primitives are fairly low-level, in addition to
general-purpose resource-oriented operations, the CLI also supports
“porcelain” for common simple, domain-specific operational operations (both
status/progress extraction and mutations) that don’t have discrete API
implementations, such as run, expose, rollout, cp, top, cordon, and
drain. And there should be support for non-resource-oriented APIs,
such as exec, logs, attach, port-forward, and proxy.

### Web dashboard (UI)

The project supported a dashboard, initially built into the apiserver,
almost from the beginning. Other projects in the space had UIs and
users expected one. There wasn’t a vendor-neutral one in the
ecosystem, however, and a solution was needed for the project's local
cluster environment, minikube. The dashboard has also served as a UI
reference implementation and a vehicle to drive conventions (e.g.,
around resource category terminology). The dashboard has also been
useful as a tool to demonstrate and to learn about Kubernetes
concepts, features, and behaviors.

### Cluster lifecycle tools

Cluster lifecycle includes provisioning, bootstrapping,
upgrade/downgrade, and teardown. The project develops several such tools.
Tools are needed for the following scenarios/purposes:
* usability of upstream releases: at least one solution that can be used to bootstrap the upstream release (e.g., kubeadm)
* testing: solutions that can be used to run multi-node end-to-end tests (e.g., kind), integration tests, upgrade/downgrade tests, version-skew tests, scalability tests, and other types of tests the projects deems necessary to ensure adequate release quality
* portable, low-dependency local environment: at least one local environment (e.g., minikube), in order to simplify documentation tutorials that require a cluster to exist

### Extensibility to support execution and management in diverse environments

Including:
* CRI
* CNI
* CSI
* external cloud providers
* KMS providers
* OSB brokers
* Cluster APIs

### Multi-cluster management tools and systems

Many users desire to operate in and deploy applications to multiple
geographic locations and environments, even across multiple providers.
This generally requires managing multiple Kubernetes clusters.  While
general deployment pipeline tools and continuous deployment systems
are not in scope, the project has explored multiple mechanisms to
simplify management of resources across multiple clusters, including
Federation v1, Federation v2, and the Cluster Registry API.

### Project GitHub automation and other process automation

As one of the largest, most active projects on Github, Kubernetes has
some extreme needs.

Including:
* prow
* gubernator
* velodrome and kettle
* website infrastructure
* k8s.io

### Project continuous build and test infrastructure

Including:
* prow
* tide
* triage dashboard

### Release tooling

Including:
* anago

### Documentation

Documentation of project-provided functionality and components, for
multiple audiences, including:
* application developers
* application operators
* cluster operators
* ecosystem developers
* distribution providers, and others who want to port Kubernetes to new environments
* project contributors

### Usage data collection mechanisms

Including:
* Spartakus

## Examples of projects and areas not in scope

Some of these are obvious, but many have been seriously deliberated in the
past.
* The resource instance store (etcd)
* Container runtimes, other than current grandfathered ones
* Network and storage plugins, other than current grandfathered ones
* CoreDNS
  * Since intra-cluster DNS is in scope, we need to ensure we have
    some solution, which has been kubedns, but now that there is an
    adequate alternative outside the project, we are adopting it.
* Service load balancers (e.g., Envoy, Linkerd), other than kube-proxy
* Cloud provider implementations, other than current grandfathered ones
* Container image build tools
* Image registries and distribution mechanisms
* Identity (user/group) sources of truth (e.g., LDAP)
* Key management systems (e.g., Vault)
* CI, CD, and GitOps (push to deploy) systems, other than
  infrastructure used to build and test the Kubernetes project itself
* Application-level services, such as middleware (e.g., message
  buses), data-processing frameworks (e.g., Spark), machine-learning
  frameworks (e.g., Kubeflow), databases (e.g., Mysql), caches, nor
  cluster storage systems (e.g., Ceph) as built-in services. Such
  components can run on Kubernetes, and/or can be accessed by
  applications running on Kubernetes through portable mechanisms, such
  as the Open Service Broker. Application-specific Operators (e.g.,
  Cassandra Operator) are also not in scope.
* Application and cluster log aggregation and searching, application
  and cluster monitoring aggregation and dashboarding (other than
  heapster, which is grandfathered), alerting, application performance
  management, tracing, and debugging tools
* General-purpose machine configuration (e.g., Chef, Puppet, Ansible,
  Salt), maintenance, automation (e.g., Rundeck), and management systems
* Templating and configuration languages (e.g., jinja, jsonnet,
  starlark, hcl, dhall, hocon)
* File packaging tools (e.g., helm, kpm, kubepack, duffle)
* Managing non-containerized applications in VMs, and other general
  IaaS functionality
* Full Platform as a Service functionality
* Full Functions as a Service functionality
* [Workflow
 orchestration](https://github.com/kubernetes/kubernetes/pull/24781#issuecomment-215914822):
 "Workflow" is a very broad, diverse area, with solutions typically
 tailored to specific use cases (e.g., data-flow graphs, data-driven
 processing, deployment pipelines, event-driven automation,
 business-process execution, iPaaS) and specific input and event
 sources, and often requires arbitrary code to evaluate conditions,
 actions, and/or failure handling.
* Other forms of human-oriented and programmatic interfaces over the
  Kubernetes API other than “basic” CLIs (e.g., kubectl) and UI
  (dashboard), such as mobile dashboards, IDEs, chat bots, SQL,
  interactive shells, etc.
